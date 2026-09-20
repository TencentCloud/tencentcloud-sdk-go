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

package v20250920

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type APIKeyInfo struct {
	// API密钥名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// API密钥ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 密钥状态。可以为API_KEY_STATUS_ACTIVE，或API_KEY_STATUS_INACTIVE
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 隐藏部分字符的API密钥，方便用户辨认
	MaskedKey *string `json:"MaskedKey,omitnil,omitempty" name:"MaskedKey"`

	// API密钥创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`
}

type AccountQuotaOverview struct {
	// <p>主账号各资源维度的配额上限</p>
	Quota *QuotaResourceInfo `json:"Quota,omitnil,omitempty" name:"Quota"`

	// <p>主账号各资源维度的当前用量</p>
	Usage *QuotaResourceInfo `json:"Usage,omitnil,omitempty" name:"Usage"`
}

// Predefined struct for user
type AcquireDeploymentTokenRequestParams struct {
	// <p>目标 ACTIVE Deployment 的稳定 ID。</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`
}

type AcquireDeploymentTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>目标 ACTIVE Deployment 的稳定 ID。</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`
}

func (r *AcquireDeploymentTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcquireDeploymentTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeploymentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcquireDeploymentTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcquireDeploymentTokenResponseParams struct {
	// <p>只用于目标 Deployment 数据面入口的短期 bearer Token，格式为 dpt_ 加非空、无 padding 的 Base64URL opaque 后缀。</p>
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// <p>Token 的绝对过期时间，UTC、秒精度 RFC3339 格式。</p>
	ExpiresAt *string `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AcquireDeploymentTokenResponse struct {
	*tchttp.BaseResponse
	Response *AcquireDeploymentTokenResponseParams `json:"Response"`
}

func (r *AcquireDeploymentTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcquireDeploymentTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcquireSandboxInstanceTokenRequestParams struct {
	// <p>沙箱实例ID，生成的访问Token将仅可用于访问此沙箱实例</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type AcquireSandboxInstanceTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>沙箱实例ID，生成的访问Token将仅可用于访问此沙箱实例</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *AcquireSandboxInstanceTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcquireSandboxInstanceTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcquireSandboxInstanceTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcquireSandboxInstanceTokenResponseParams struct {
	// <p>访问Token</p>
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// <p>过期时间</p>
	ExpiresAt *string `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// <p>除管控面envd端口(49983)以外端口的访问Token</p>
	TrafficToken *string `json:"TrafficToken,omitnil,omitempty" name:"TrafficToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AcquireSandboxInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *AcquireSandboxInstanceTokenResponseParams `json:"Response"`
}

func (r *AcquireSandboxInstanceTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcquireSandboxInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AffinityConfiguration struct {
	// <p>Affinity 模式。</p><p>枚举值：</p><ul><li>BEST_EFFORT：优先复用原 Instance，不可用时允许改选。</li><li>STRICT：只复用原 Instance，不可用时失败且不改选。</li><li>EXCLUSIVE：一个 Affinity ID 独占一个 Instance，不能迁移。</li></ul><p>缺失或空字符串表示关闭 Affinity。</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>请求和响应使用的 Affinity Header 名称。必须符合 HTTP field-name token 语法，长度为 1..128 个 ASCII 字节，且不能使用平台保留 Header。</p>
	HeaderName *string `json:"HeaderName,omitnil,omitempty" name:"HeaderName"`
}

type AgentBucketStorageSource struct {
	// <p>用于传入 AgentBucket 的 LibraryID</p>
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// <p>用于传入 AgentBucket 的 spaceId</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>用于传入 AgentBucket 的 AccessDomain</p>
	AccessDomain *string `json:"AccessDomain,omitnil,omitempty" name:"AccessDomain"`
}

// Predefined struct for user
type AppendEventRequestParams struct {
	// <p>会话所属空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>用户 ID。可通过调用方业务系统接口获取。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>会话 ID。可通过 CreateSession 或 DescribeSessions 接口获取。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>事件内容。</p>
	Event *EventInfo `json:"Event,omitnil,omitempty" name:"Event"`

	// <p>Agent ID。可选。</p>
	//
	// Deprecated: AgentId is deprecated.
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type AppendEventRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话所属空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>用户 ID。可通过调用方业务系统接口获取。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>会话 ID。可通过 CreateSession 或 DescribeSessions 接口获取。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>事件内容。</p>
	Event *EventInfo `json:"Event,omitnil,omitempty" name:"Event"`

	// <p>Agent ID。可选。</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *AppendEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppendEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "UserId")
	delete(f, "SessionId")
	delete(f, "Event")
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AppendEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AppendEventResponseParams struct {
	// <p>事件信息。</p>
	Event *EventInfo `json:"Event,omitnil,omitempty" name:"Event"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AppendEventResponse struct {
	*tchttp.BaseResponse
	Response *AppendEventResponseParams `json:"Response"`
}

func (r *AppendEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AppendEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApproveRegistryRecordRequestParams struct {

}

type ApproveRegistryRecordRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ApproveRegistryRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApproveRegistryRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApproveRegistryRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApproveRegistryRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApproveRegistryRecordResponse struct {
	*tchttp.BaseResponse
	Response *ApproveRegistryRecordResponseParams `json:"Response"`
}

func (r *ApproveRegistryRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApproveRegistryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CLSConfig struct {
	// 沙箱工具日志推送所使用的CLS日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

// Predefined struct for user
type CancelRegistryRecordRequestParams struct {

}

type CancelRegistryRecordRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CancelRegistryRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelRegistryRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelRegistryRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelRegistryRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelRegistryRecordResponse struct {
	*tchttp.BaseResponse
	Response *CancelRegistryRecordResponseParams `json:"Response"`
}

func (r *CancelRegistryRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelRegistryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CfsStorageSource struct {
	// CFS资源ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// CFS挂载路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type ComputerConfiguration struct {
	// <p>waa沙箱工具配置</p>
	WAAConfiguration *WAAConfiguration `json:"WAAConfiguration,omitnil,omitempty" name:"WAAConfiguration"`

	// <p>配置内置 OSWorld</p>
	OSWorldConfiguration *OSWorldConfiguration `json:"OSWorldConfiguration,omitnil,omitempty" name:"OSWorldConfiguration"`
}

type CosStorageSource struct {
	// 对象存储访问域名
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 对象存储桶名称
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 对象存储桶路径，必须为以/起始的绝对路径
	BucketPath *string `json:"BucketPath,omitnil,omitempty" name:"BucketPath"`
}

// Predefined struct for user
type CreateAPIKeyRequestParams struct {
	// API密钥名称，方便用户记忆
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CreateAPIKeyRequest struct {
	*tchttp.BaseRequest
	
	// API密钥名称，方便用户记忆
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *CreateAPIKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAPIKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAPIKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAPIKeyResponseParams struct {
	// 用户传入的API密钥名称，方便用户记忆
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 生成的API密钥，仅返回此一次，后续无法获取
	APIKey *string `json:"APIKey,omitnil,omitempty" name:"APIKey"`

	// API密钥ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAPIKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAPIKeyResponseParams `json:"Response"`
}

func (r *CreateAPIKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAPIKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeploymentRequestParams struct {
	// <p>唯一的 Deployment 名称，必须符合 DNS-1123 命名规范，创建后不可修改。</p>
	DeploymentName *string `json:"DeploymentName,omitnil,omitempty" name:"DeploymentName"`

	// <p>用于关联 Sandbox Tool 的标识，格式为 sdt- 加 8 位小写 base36 字符。</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>伸缩配置；省略的成员由服务端补全默认值。</p>
	ScalingConfiguration *ScalingConfiguration `json:"ScalingConfiguration,omitnil,omitempty" name:"ScalingConfiguration"`

	// <p>空闲生命周期配置；省略的成员由服务端补全默认值。</p>
	LifecycleConfiguration *LifecycleConfiguration `json:"LifecycleConfiguration,omitnil,omitempty" name:"LifecycleConfiguration"`

	// <p>Affinity 配置；省略或空 Mode 表示不启用。</p>
	AffinityConfiguration *AffinityConfiguration `json:"AffinityConfiguration,omitnil,omitempty" name:"AffinityConfiguration"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateDeploymentRequest struct {
	*tchttp.BaseRequest
	
	// <p>唯一的 Deployment 名称，必须符合 DNS-1123 命名规范，创建后不可修改。</p>
	DeploymentName *string `json:"DeploymentName,omitnil,omitempty" name:"DeploymentName"`

	// <p>用于关联 Sandbox Tool 的标识，格式为 sdt- 加 8 位小写 base36 字符。</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>伸缩配置；省略的成员由服务端补全默认值。</p>
	ScalingConfiguration *ScalingConfiguration `json:"ScalingConfiguration,omitnil,omitempty" name:"ScalingConfiguration"`

	// <p>空闲生命周期配置；省略的成员由服务端补全默认值。</p>
	LifecycleConfiguration *LifecycleConfiguration `json:"LifecycleConfiguration,omitnil,omitempty" name:"LifecycleConfiguration"`

	// <p>Affinity 配置；省略或空 Mode 表示不启用。</p>
	AffinityConfiguration *AffinityConfiguration `json:"AffinityConfiguration,omitnil,omitempty" name:"AffinityConfiguration"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateDeploymentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeploymentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeploymentName")
	delete(f, "ToolId")
	delete(f, "ScalingConfiguration")
	delete(f, "LifecycleConfiguration")
	delete(f, "AffinityConfiguration")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeploymentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeploymentResponseParams struct {
	// <p>已创建并完成默认值物化的 Deployment。</p>
	Deployment *Deployment `json:"Deployment,omitnil,omitempty" name:"Deployment"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDeploymentResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeploymentResponseParams `json:"Response"`
}

func (r *CreateDeploymentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePreCacheImageTaskRequestParams struct {
	// <p>镜像地址</p>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>镜像仓库类型：<code>enterprise</code>、<code>personal</code>、<code>custom</code></p><p>枚举值：</p><ul><li>enterprise： tcr 企业容器镜像服务</li><li>personal： ccr 个人容器镜像服务</li></ul>
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`
}

type CreatePreCacheImageTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>镜像地址</p>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>镜像仓库类型：<code>enterprise</code>、<code>personal</code>、<code>custom</code></p><p>枚举值：</p><ul><li>enterprise： tcr 企业容器镜像服务</li><li>personal： ccr 个人容器镜像服务</li></ul>
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`
}

func (r *CreatePreCacheImageTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePreCacheImageTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "ImageRegistryType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePreCacheImageTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePreCacheImageTaskResponseParams struct {
	// <p>镜像地址</p>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>镜像 Digest</p>
	ImageDigest *string `json:"ImageDigest,omitnil,omitempty" name:"ImageDigest"`

	// <p>镜像仓库类型：<code>enterprise</code>、<code>personal</code>。</p>
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePreCacheImageTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreatePreCacheImageTaskResponseParams `json:"Response"`
}

func (r *CreatePreCacheImageTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePreCacheImageTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRegistryRecordRequestParams struct {

}

type CreateRegistryRecordRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateRegistryRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRegistryRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRegistryRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRegistryRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRegistryRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreateRegistryRecordResponseParams `json:"Response"`
}

func (r *CreateRegistryRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRegistryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRegistryRequestParams struct {

}

type CreateRegistryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateRegistryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRegistryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRegistryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRegistryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRegistryResponse struct {
	*tchttp.BaseResponse
	Response *CreateRegistryResponseParams `json:"Response"`
}

func (r *CreateRegistryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRegistryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSandboxToolRequestParams struct {
	// <p>沙箱工具名称，长度 1-50 字符，支持英文、数字、下划线和连接线。同一 AppId 下沙箱工具名称必须唯一</p>
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// <p>沙箱工具类型，目前支持：browser、code-interpreter、custom等</p><p>枚举值：</p><ul><li>browser： browser</li><li>code-interpreter： code-interpreter</li><li>mobile： mobile</li><li>osworld： osworld</li><li>custom： custom</li><li>swebench： swebench</li><li>aio： aio</li><li>android-world： android-world</li><li>waa： waa</li></ul>
	ToolType *string `json:"ToolType,omitnil,omitempty" name:"ToolType"`

	// <p>网络配置</p>
	NetworkConfiguration *NetworkConfiguration `json:"NetworkConfiguration,omitnil,omitempty" name:"NetworkConfiguration"`

	// <p>沙箱工具描述，最大长度 200 字符</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>默认超时时间，支持格式：5m、300s、1h 等，不指定则使用系统默认值（5 分钟）。最大 24 小时</p>
	DefaultTimeout *string `json:"DefaultTimeout,omitnil,omitempty" name:"DefaultTimeout"`

	// <p>标签规格，为沙箱工具绑定标签，支持多种资源类型的标签绑定</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>幂等性 Token，长度不超过 64 字符</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// <p>角色ARN</p>
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// <p>沙箱工具存储配置</p>
	StorageMounts []*StorageMount `json:"StorageMounts,omitnil,omitempty" name:"StorageMounts"`

	// <p>沙箱工具自定义配置</p>
	CustomConfiguration *CustomConfiguration `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`

	// <p>桌面电脑环境类沙箱配置</p>
	ComputerConfiguration *ComputerConfiguration `json:"ComputerConfiguration,omitnil,omitempty" name:"ComputerConfiguration"`

	// <p>沙箱工具日志推送相关配置</p>
	LogConfiguration *LogConfiguration `json:"LogConfiguration,omitnil,omitempty" name:"LogConfiguration"`

	// <p>常驻沙箱标识</p>
	Persistent *bool `json:"Persistent,omitnil,omitempty" name:"Persistent"`
}

type CreateSandboxToolRequest struct {
	*tchttp.BaseRequest
	
	// <p>沙箱工具名称，长度 1-50 字符，支持英文、数字、下划线和连接线。同一 AppId 下沙箱工具名称必须唯一</p>
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// <p>沙箱工具类型，目前支持：browser、code-interpreter、custom等</p><p>枚举值：</p><ul><li>browser： browser</li><li>code-interpreter： code-interpreter</li><li>mobile： mobile</li><li>osworld： osworld</li><li>custom： custom</li><li>swebench： swebench</li><li>aio： aio</li><li>android-world： android-world</li><li>waa： waa</li></ul>
	ToolType *string `json:"ToolType,omitnil,omitempty" name:"ToolType"`

	// <p>网络配置</p>
	NetworkConfiguration *NetworkConfiguration `json:"NetworkConfiguration,omitnil,omitempty" name:"NetworkConfiguration"`

	// <p>沙箱工具描述，最大长度 200 字符</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>默认超时时间，支持格式：5m、300s、1h 等，不指定则使用系统默认值（5 分钟）。最大 24 小时</p>
	DefaultTimeout *string `json:"DefaultTimeout,omitnil,omitempty" name:"DefaultTimeout"`

	// <p>标签规格，为沙箱工具绑定标签，支持多种资源类型的标签绑定</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>幂等性 Token，长度不超过 64 字符</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// <p>角色ARN</p>
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// <p>沙箱工具存储配置</p>
	StorageMounts []*StorageMount `json:"StorageMounts,omitnil,omitempty" name:"StorageMounts"`

	// <p>沙箱工具自定义配置</p>
	CustomConfiguration *CustomConfiguration `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`

	// <p>桌面电脑环境类沙箱配置</p>
	ComputerConfiguration *ComputerConfiguration `json:"ComputerConfiguration,omitnil,omitempty" name:"ComputerConfiguration"`

	// <p>沙箱工具日志推送相关配置</p>
	LogConfiguration *LogConfiguration `json:"LogConfiguration,omitnil,omitempty" name:"LogConfiguration"`

	// <p>常驻沙箱标识</p>
	Persistent *bool `json:"Persistent,omitnil,omitempty" name:"Persistent"`
}

func (r *CreateSandboxToolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSandboxToolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ToolName")
	delete(f, "ToolType")
	delete(f, "NetworkConfiguration")
	delete(f, "Description")
	delete(f, "DefaultTimeout")
	delete(f, "Tags")
	delete(f, "ClientToken")
	delete(f, "RoleArn")
	delete(f, "StorageMounts")
	delete(f, "CustomConfiguration")
	delete(f, "ComputerConfiguration")
	delete(f, "LogConfiguration")
	delete(f, "Persistent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSandboxToolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSandboxToolResponseParams struct {
	// <p>创建的沙箱工具 ID</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSandboxToolResponse struct {
	*tchttp.BaseResponse
	Response *CreateSandboxToolResponseParams `json:"Response"`
}

func (r *CreateSandboxToolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSandboxToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSessionRequestParams struct {
	// <p>会话所属空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>用户 ID。可通过调用方业务系统接口获取。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Agent ID。可选。</p>
	//
	// Deprecated: AgentId is deprecated.
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>会话 ID。可通过 CreateSession 或 DescribeSessions 接口获取。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>会话标题，最大长度 256 字符。</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>初始会话状态。</p>
	State *SessionState `json:"State,omitnil,omitempty" name:"State"`

	// <p>创建会话时设置的初始元数据，以键值对数组形式表示。每个元素包含 Metadata 名称和对应值。</p><p>入参限制：本参数可选，最多支持 64 项。Name 不能为空或重复，最大长度为 253 字节；Value 最大长度为 1024 字节，允许为空字符串。Metadata 序列化后的总大小不能超过 64 KiB。</p>
	Metadata []*MetadataVar `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

type CreateSessionRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话所属空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>用户 ID。可通过调用方业务系统接口获取。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Agent ID。可选。</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>会话 ID。可通过 CreateSession 或 DescribeSessions 接口获取。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>会话标题，最大长度 256 字符。</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>初始会话状态。</p>
	State *SessionState `json:"State,omitnil,omitempty" name:"State"`

	// <p>创建会话时设置的初始元数据，以键值对数组形式表示。每个元素包含 Metadata 名称和对应值。</p><p>入参限制：本参数可选，最多支持 64 项。Name 不能为空或重复，最大长度为 253 字节；Value 最大长度为 1024 字节，允许为空字符串。Metadata 序列化后的总大小不能超过 64 KiB。</p>
	Metadata []*MetadataVar `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

func (r *CreateSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "UserId")
	delete(f, "AgentId")
	delete(f, "SessionId")
	delete(f, "Title")
	delete(f, "State")
	delete(f, "Metadata")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSessionResponseParams struct {
	// <p>会话信息。</p>
	Session *SessionInfo `json:"Session,omitnil,omitempty" name:"Session"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateSessionResponseParams `json:"Response"`
}

func (r *CreateSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSessionSpaceRequestParams struct {
	// <p>会话空间名称，用于标识会话空间的业务用途。</p><p>入参限制：必填；去除首尾空白后不能为空；最大长度为 128 个字符。</p><p>建议名称包含业务和环境信息，便于识别和管理。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>会话空间描述，用于补充说明会话空间的业务用途。</p><p>入参限制：选填；最大长度为 512 个字符。</p><p>未传入时创建为空描述。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>创建 SessionSpace 时为资源绑定标签。</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateSessionSpaceRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话空间名称，用于标识会话空间的业务用途。</p><p>入参限制：必填；去除首尾空白后不能为空；最大长度为 128 个字符。</p><p>建议名称包含业务和环境信息，便于识别和管理。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>会话空间描述，用于补充说明会话空间的业务用途。</p><p>入参限制：选填；最大长度为 512 个字符。</p><p>未传入时创建为空描述。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>创建 SessionSpace 时为资源绑定标签。</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateSessionSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSessionSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSessionSpaceResponseParams struct {
	// <p>创建成功后的会话空间完整信息。</p><p>接口成功时一定返回；接口失败时返回 Error，不会返回该字段。</p>
	SessionSpace *SessionSpaceInfo `json:"SessionSpace,omitnil,omitempty" name:"SessionSpace"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSessionSpaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateSessionSpaceResponseParams `json:"Response"`
}

func (r *CreateSessionSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomConfiguration struct {
	// <p>镜像地址</p>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>镜像仓库类型：<code>enterprise</code>、<code>personal</code>、<code>custom</code></p><p>枚举值：</p><ul><li>enterprise： tcr 企业容器镜像服务</li><li>personal： ccr 个人容器镜像服务</li></ul>
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`

	// <p>启动命令</p>
	Command []*string `json:"Command,omitnil,omitempty" name:"Command"`

	// <p>启动参数</p>
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`

	// <p>环境变量</p>
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>端口配置</p>
	Ports []*PortConfiguration `json:"Ports,omitnil,omitempty" name:"Ports"`

	// <p>资源配置</p>
	Resources *ResourceConfiguration `json:"Resources,omitnil,omitempty" name:"Resources"`

	// <p>探针配置</p>
	Probe *ProbeConfiguration `json:"Probe,omitnil,omitempty" name:"Probe"`

	// <p>沙箱 DNS 配置</p>
	DNSConfig *DNSConfig `json:"DNSConfig,omitnil,omitempty" name:"DNSConfig"`
}

type CustomConfigurationDetail struct {
	// <p>镜像地址</p>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>镜像仓库类型：<code>enterprise</code>、<code>personal</code>、<code>custom</code>。</p><p>枚举值：</p><ul><li>enterprise： TCR 企业容器镜像服务</li><li>personal： CCR 个人容器镜像服务</li></ul>
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`

	// <p>镜像 Digest</p>
	ImageDigest *string `json:"ImageDigest,omitnil,omitempty" name:"ImageDigest"`

	// <p>启动命令</p>
	Command []*string `json:"Command,omitnil,omitempty" name:"Command"`

	// <p>启动参数</p>
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`

	// <p>环境变量</p>
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>端口配置</p>
	Ports []*PortConfiguration `json:"Ports,omitnil,omitempty" name:"Ports"`

	// <p>资源配置</p>
	Resources *ResourceConfiguration `json:"Resources,omitnil,omitempty" name:"Resources"`

	// <p>探针配置</p>
	Probe *ProbeConfiguration `json:"Probe,omitnil,omitempty" name:"Probe"`

	// <p>沙箱 DNS 配置</p>
	DNSConfig *DNSConfig `json:"DNSConfig,omitnil,omitempty" name:"DNSConfig"`
}

type DNSConfig struct {
	// <p>DNS 服务器地址</p><p>参数格式：需要有效 IP 地址</p><p>默认值：10.0.0.1</p>
	Servers []*string `json:"Servers,omitnil,omitempty" name:"Servers"`

	// <p>搜索域(对应 resolv.conf 的 search 指令)</p>
	Searches []*string `json:"Searches,omitnil,omitempty" name:"Searches"`

	// <p>配置项(对应  resolv.conf 选项)</p>
	Options []*string `json:"Options,omitnil,omitempty" name:"Options"`
}

// Predefined struct for user
type DeleteAPIKeyRequestParams struct {
	// 需要删除的API密钥ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type DeleteAPIKeyRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的API密钥ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *DeleteAPIKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAPIKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAPIKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAPIKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAPIKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAPIKeyResponseParams `json:"Response"`
}

func (r *DeleteAPIKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAPIKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeploymentRequestParams struct {
	// <p>待删除的 Deployment ID。</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`
}

type DeleteDeploymentRequest struct {
	*tchttp.BaseRequest
	
	// <p>待删除的 Deployment ID。</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`
}

func (r *DeleteDeploymentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeploymentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeploymentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeploymentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeploymentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDeploymentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeploymentResponseParams `json:"Response"`
}

func (r *DeleteDeploymentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRegistryRecordRequestParams struct {

}

type DeleteRegistryRecordRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteRegistryRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRegistryRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRegistryRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRegistryRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRegistryRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRegistryRecordResponseParams `json:"Response"`
}

func (r *DeleteRegistryRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRegistryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRegistryRequestParams struct {

}

type DeleteRegistryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteRegistryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRegistryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRegistryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRegistryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRegistryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRegistryResponseParams `json:"Response"`
}

func (r *DeleteRegistryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRegistryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSandboxToolRequestParams struct {
	// 沙箱工具ID
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`
}

type DeleteSandboxToolRequest struct {
	*tchttp.BaseRequest
	
	// 沙箱工具ID
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`
}

func (r *DeleteSandboxToolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSandboxToolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ToolId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSandboxToolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSandboxToolResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSandboxToolResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSandboxToolResponseParams `json:"Response"`
}

func (r *DeleteSandboxToolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSandboxToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSessionRequestParams struct {
	// <p>会话所属空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>用户 ID。可通过调用方业务系统接口获取。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>会话 ID。可通过 CreateSession 或 DescribeSessions 接口获取。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Agent ID。可选。</p>
	//
	// Deprecated: AgentId is deprecated.
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type DeleteSessionRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话所属空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>用户 ID。可通过调用方业务系统接口获取。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>会话 ID。可通过 CreateSession 或 DescribeSessions 接口获取。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Agent ID。可选。</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *DeleteSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "UserId")
	delete(f, "SessionId")
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSessionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSessionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSessionResponseParams `json:"Response"`
}

func (r *DeleteSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSessionSpaceRequestParams struct {
	// <p>需要删除的会话空间唯一标识。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DeleteSessionSpaceRequest struct {
	*tchttp.BaseRequest
	
	// <p>需要删除的会话空间唯一标识。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DeleteSessionSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSessionSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSessionSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSessionSpaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSessionSpaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSessionSpaceResponseParams `json:"Response"`
}

func (r *DeleteSessionSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSessionSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Deployment struct {
	// <p>Deployment 稳定 ID，格式为 dpl- 加 8 位小写 base36 字符。</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`

	// <p>唯一且创建后不可修改的名称，必须符合 DNS-1123 命名规范。</p>
	DeploymentName *string `json:"DeploymentName,omitnil,omitempty" name:"DeploymentName"`

	// <p>用于关联 Sandbox Tool 的标识，格式为 sdt- 加 8 位小写 base36 字符。</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>完整的活跃容量配置。</p>
	ScalingConfiguration *ScalingConfiguration `json:"ScalingConfiguration,omitnil,omitempty" name:"ScalingConfiguration"`

	// <p>完整的空闲生命周期配置。</p>
	LifecycleConfiguration *LifecycleConfiguration `json:"LifecycleConfiguration,omitnil,omitempty" name:"LifecycleConfiguration"`

	// <p>可选 Affinity 配置；未启用时省略。</p>
	AffinityConfiguration *AffinityConfiguration `json:"AffinityConfiguration,omitnil,omitempty" name:"AffinityConfiguration"`

	// <p>Deployment 控制面状态。</p><p>枚举值：</p><ul><li>ACTIVE：入口可用。</li><li>DELETING：入口已关闭并正在异步删除。</li><li>DELETE_FAILED：最近一次异步删除失败，可再次调用 DeleteDeployment。</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>DELETE_FAILED 状态下 1..1024 个 UTF-8 字节的安全失败摘要，格式为 {Code}[.{SubCode}]: {Message}；其他状态省略。</p>
	StatusReason *string `json:"StatusReason,omitnil,omitempty" name:"StatusReason"`

	// <p>创建时间，UTC、秒精度 RFC3339 格式。</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>最近一次成功公共配置写入或 Deployment 状态迁移时间，UTC、秒精度 RFC3339 格式。</p>
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

// Predefined struct for user
type DescribeAPIKeyListRequestParams struct {

}

type DescribeAPIKeyListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAPIKeyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIKeyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAPIKeyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPIKeyListResponseParams struct {
	// API密钥简略信息列表。
	APIKeySet []*APIKeyInfo `json:"APIKeySet,omitnil,omitempty" name:"APIKeySet"`

	// 列表中API密钥数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAPIKeyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAPIKeyListResponseParams `json:"Response"`
}

func (r *DescribeAPIKeyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIKeyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeploymentListRequestParams struct {
	// <p>分页偏移量，默认 0，必须大于等于 0。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>分页返回数量，默认 20，范围 1..200。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>查询过滤条件。</p><p>Filter.Name 枚举值：</p><ul><li>deployment-id：按 DeploymentId 精确匹配</li><li>deployment-name：按 DeploymentName 精确匹配</li><li>deployment-name-like：按 DeploymentName 进行普通文本包含匹配，%、_ 等字符没有通配语义</li><li>tool-id：按 ToolId 精确匹配</li><li>status：按 Deployment 状态精确匹配，支持 ACTIVE、DELETING、DELETE_FAILED</li></ul><p>所有匹配均区分大小写。不同 Filter 之间为 AND，同一 Filter 的 Values 之间为 OR。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDeploymentListRequest struct {
	*tchttp.BaseRequest
	
	// <p>分页偏移量，默认 0，必须大于等于 0。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>分页返回数量，默认 20，范围 1..200。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>查询过滤条件。</p><p>Filter.Name 枚举值：</p><ul><li>deployment-id：按 DeploymentId 精确匹配</li><li>deployment-name：按 DeploymentName 精确匹配</li><li>deployment-name-like：按 DeploymentName 进行普通文本包含匹配，%、_ 等字符没有通配语义</li><li>tool-id：按 ToolId 精确匹配</li><li>status：按 Deployment 状态精确匹配，支持 ACTIVE、DELETING、DELETE_FAILED</li></ul><p>所有匹配均区分大小写。不同 Filter 之间为 AND，同一 Filter 的 Values 之间为 OR。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeDeploymentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeploymentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeploymentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeploymentListResponseParams struct {
	// <p>当前页完整 Deployment；无匹配时为空数组。</p>
	DeploymentSet []*Deployment `json:"DeploymentSet,omitnil,omitempty" name:"DeploymentSet"`

	// <p>应用 Filters 后、分页前的结果总数。</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeploymentListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeploymentListResponseParams `json:"Response"`
}

func (r *DescribeDeploymentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeploymentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeploymentRequestParams struct {
	// <p>待查询的 Deployment ID。</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`
}

type DescribeDeploymentRequest struct {
	*tchttp.BaseRequest
	
	// <p>待查询的 Deployment ID。</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`
}

func (r *DescribeDeploymentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeploymentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeploymentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeploymentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeploymentResponseParams struct {
	// <p>完整 Deployment。</p>
	Deployment *Deployment `json:"Deployment,omitnil,omitempty" name:"Deployment"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeploymentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeploymentResponseParams `json:"Response"`
}

func (r *DescribeDeploymentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsRequestParams struct {
	// <p>会话所属空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>用户 ID。可通过调用方业务系统接口获取。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>会话 ID。可通过 CreateSession 或 DescribeSessions 接口获取。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Agent ID。可选。</p>
	//
	// Deprecated: AgentId is deprecated.
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>事件作者。取值示例：user、assistant、tool。</p>
	Author *string `json:"Author,omitnil,omitempty" name:"Author"`

	// <p>起始时间，仅返回该时间之后的事件，使用 RFC3339 格式，最大长度 64 字符。</p>
	AfterTimestamp *string `json:"AfterTimestamp,omitnil,omitempty" name:"AfterTimestamp"`

	// <p>分页偏移量，默认为 0。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认为 50，最大值为 200。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeEventsRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话所属空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>用户 ID。可通过调用方业务系统接口获取。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>会话 ID。可通过 CreateSession 或 DescribeSessions 接口获取。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Agent ID。可选。</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>事件作者。取值示例：user、assistant、tool。</p>
	Author *string `json:"Author,omitnil,omitempty" name:"Author"`

	// <p>起始时间，仅返回该时间之后的事件，使用 RFC3339 格式，最大长度 64 字符。</p>
	AfterTimestamp *string `json:"AfterTimestamp,omitnil,omitempty" name:"AfterTimestamp"`

	// <p>分页偏移量，默认为 0。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认为 50，最大值为 200。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "UserId")
	delete(f, "SessionId")
	delete(f, "AgentId")
	delete(f, "Author")
	delete(f, "AfterTimestamp")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsResponseParams struct {
	// <p>事件列表。</p>
	Events []*EventInfo `json:"Events,omitnil,omitempty" name:"Events"`

	// <p>符合条件的事件总数。</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventsResponseParams `json:"Response"`
}

func (r *DescribeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePreCacheImageTaskRequestParams struct {
	// <p>镜像地址</p>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>镜像 Digest</p>
	ImageDigest *string `json:"ImageDigest,omitnil,omitempty" name:"ImageDigest"`

	// <p>镜像仓库类型：<code>enterprise</code>、<code>personal</code>、<code>custom</code> 。</p><p>枚举值：</p><ul><li>enterprise： tcr 企业容器镜像服务</li><li>personal： ccr 个人容器镜像服务</li></ul>
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`
}

type DescribePreCacheImageTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>镜像地址</p>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>镜像 Digest</p>
	ImageDigest *string `json:"ImageDigest,omitnil,omitempty" name:"ImageDigest"`

	// <p>镜像仓库类型：<code>enterprise</code>、<code>personal</code>、<code>custom</code> 。</p><p>枚举值：</p><ul><li>enterprise： tcr 企业容器镜像服务</li><li>personal： ccr 个人容器镜像服务</li></ul>
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`
}

func (r *DescribePreCacheImageTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePreCacheImageTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "ImageDigest")
	delete(f, "ImageRegistryType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePreCacheImageTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePreCacheImageTaskResponseParams struct {
	// <p>镜像地址</p>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>镜像 Digest</p>
	ImageDigest *string `json:"ImageDigest,omitnil,omitempty" name:"ImageDigest"`

	// <p>镜像仓库类型：<code>enterprise</code>、<code>personal</code>。</p>
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`

	// <p>镜像预热状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>镜像预热状态描述</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePreCacheImageTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribePreCacheImageTaskResponseParams `json:"Response"`
}

func (r *DescribePreCacheImageTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePreCacheImageTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaOverviewRequestParams struct {
	// <p>分页偏移量，从 0 开始，默认值为 0，必须大于等于 0。</p><p>单位：偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页返回的配额组数量</p><p>单位：个</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>配额组过滤条件</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeQuotaOverviewRequest struct {
	*tchttp.BaseRequest
	
	// <p>分页偏移量，从 0 开始，默认值为 0，必须大于等于 0。</p><p>单位：偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页返回的配额组数量</p><p>单位：个</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>配额组过滤条件</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeQuotaOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQuotaOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaOverviewResponseParams struct {
	// <p>主账号配额上限及全账号当前用量</p>
	AccountQuotaOverview *AccountQuotaOverview `json:"AccountQuotaOverview,omitnil,omitempty" name:"AccountQuotaOverview"`

	// <p>当前分页下的配额组配额与用量列表。没有数据时返回空数组。</p>
	QuotaGroupSet []*QuotaGroupOverview `json:"QuotaGroupSet,omitnil,omitempty" name:"QuotaGroupSet"`

	// <p>满足过滤条件的配额组总数，不受当前分页大小影响。</p><p>单位：个</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>本次查询完成时间，格式为 RFC3339</p>
	DataTime *string `json:"DataTime,omitnil,omitempty" name:"DataTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQuotaOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQuotaOverviewResponseParams `json:"Response"`
}

func (r *DescribeQuotaOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistryAuditLogListRequestParams struct {

}

type DescribeRegistryAuditLogListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegistryAuditLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistryAuditLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegistryAuditLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistryAuditLogListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegistryAuditLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegistryAuditLogListResponseParams `json:"Response"`
}

func (r *DescribeRegistryAuditLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistryAuditLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistryListRequestParams struct {

}

type DescribeRegistryListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegistryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegistryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistryListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegistryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegistryListResponseParams `json:"Response"`
}

func (r *DescribeRegistryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistryRecordListRequestParams struct {

}

type DescribeRegistryRecordListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegistryRecordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistryRecordListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegistryRecordListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistryRecordListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegistryRecordListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegistryRecordListResponseParams `json:"Response"`
}

func (r *DescribeRegistryRecordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistryRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistryRecordRequestParams struct {

}

type DescribeRegistryRecordRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegistryRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistryRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegistryRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistryRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegistryRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegistryRecordResponseParams `json:"Response"`
}

func (r *DescribeRegistryRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistryRecordVersionListRequestParams struct {

}

type DescribeRegistryRecordVersionListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegistryRecordVersionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistryRecordVersionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegistryRecordVersionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistryRecordVersionListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegistryRecordVersionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegistryRecordVersionListResponseParams `json:"Response"`
}

func (r *DescribeRegistryRecordVersionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistryRecordVersionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistryRequestParams struct {

}

type DescribeRegistryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegistryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegistryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegistryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegistryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegistryResponseParams `json:"Response"`
}

func (r *DescribeRegistryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegistryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSandboxInstanceListRequestParams struct {
	// <p>沙箱实例ID列表，指定要查询的实例。如果为空则查询所有实例。最大支持100个ID</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>沙箱工具ID，指定时查询该沙箱模板下的实例，为空则查询所有沙箱模板的实例</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>偏移量，默认为0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认为20，最大值为100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>每次调用返回的最大结果数。如果查询返回的时候有NextToken返回，您可以使用NextToken值获取更多页结果， 当NextToke返回空或者返回的结果数量小于MaxResults时，表示没有更多数据了。允许的最大页面大小为 100。</p>
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>如果NextToken返回非空字符串 ，表示还有更多可用结果。 NextToken是每个页面唯一的分页令牌。使用返回的令牌再次调用以检索下一页。需要保持所有其他参数不变。每个分页令牌在 24 小时后过期。</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// <p>是否返回符合当前查询条件的沙箱实例总数，仅在使用 MaxResults/NextToken 分页时生效。设置为 true 时，首次请求（NextToken 为空）计算并返回精确的 TotalCount；后续使用 NextToken 翻页时返回首次请求计算的 TotalCount，分页期间该值保持不变。重新发起不带 NextToken 的请求时将重新计算。使用 NextToken 翻页时，本参数及其他查询参数必须与首次请求保持一致。默认值为 false，此时 TotalCount 返回 0。</p>
	NeedTotalCount *bool `json:"NeedTotalCount,omitnil,omitempty" name:"NeedTotalCount"`
}

type DescribeSandboxInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// <p>沙箱实例ID列表，指定要查询的实例。如果为空则查询所有实例。最大支持100个ID</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>沙箱工具ID，指定时查询该沙箱模板下的实例，为空则查询所有沙箱模板的实例</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>偏移量，默认为0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认为20，最大值为100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>每次调用返回的最大结果数。如果查询返回的时候有NextToken返回，您可以使用NextToken值获取更多页结果， 当NextToke返回空或者返回的结果数量小于MaxResults时，表示没有更多数据了。允许的最大页面大小为 100。</p>
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>如果NextToken返回非空字符串 ，表示还有更多可用结果。 NextToken是每个页面唯一的分页令牌。使用返回的令牌再次调用以检索下一页。需要保持所有其他参数不变。每个分页令牌在 24 小时后过期。</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// <p>是否返回符合当前查询条件的沙箱实例总数，仅在使用 MaxResults/NextToken 分页时生效。设置为 true 时，首次请求（NextToken 为空）计算并返回精确的 TotalCount；后续使用 NextToken 翻页时返回首次请求计算的 TotalCount，分页期间该值保持不变。重新发起不带 NextToken 的请求时将重新计算。使用 NextToken 翻页时，本参数及其他查询参数必须与首次请求保持一致。默认值为 false，此时 TotalCount 返回 0。</p>
	NeedTotalCount *bool `json:"NeedTotalCount,omitnil,omitempty" name:"NeedTotalCount"`
}

func (r *DescribeSandboxInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSandboxInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ToolId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	delete(f, "NeedTotalCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSandboxInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSandboxInstanceListResponseParams struct {
	// <p>沙箱实例列表</p>
	InstanceSet []*SandboxInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// <p>符合条件的实例总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>如果NextToken返回非空字符串 ，表示还有更多可用结果。 NextToken是每个页面唯一的分页令牌。使用返回的令牌再次调用以检索下一页。需要保持所有其他参数不变。每个分页令牌在 24 小时后过期。</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSandboxInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSandboxInstanceListResponseParams `json:"Response"`
}

func (r *DescribeSandboxInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSandboxInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSandboxToolListRequestParams struct {
	// 沙箱工具ID列表，指定要查询的工具。如果为空则查询所有工具。最大支持100个ID
	ToolIds []*string `json:"ToolIds,omitnil,omitempty" name:"ToolIds"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSandboxToolListRequest struct {
	*tchttp.BaseRequest
	
	// 沙箱工具ID列表，指定要查询的工具。如果为空则查询所有工具。最大支持100个ID
	ToolIds []*string `json:"ToolIds,omitnil,omitempty" name:"ToolIds"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeSandboxToolListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSandboxToolListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ToolIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSandboxToolListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSandboxToolListResponseParams struct {
	// 沙箱工具列表
	SandboxToolSet []*SandboxTool `json:"SandboxToolSet,omitnil,omitempty" name:"SandboxToolSet"`

	// 符合条件的沙箱工具总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSandboxToolListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSandboxToolListResponseParams `json:"Response"`
}

func (r *DescribeSandboxToolListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSandboxToolListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionRequestParams struct {
	// <p>会话所属空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>用户 ID。可通过调用方业务系统接口获取。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>会话 ID。可通过 CreateSession 或 DescribeSessions 接口获取。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Agent ID。可选。</p>
	//
	// Deprecated: AgentId is deprecated.
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>返回最近事件数量，默认为 0，最大值为 200。</p>
	NumRecentEvents *int64 `json:"NumRecentEvents,omitnil,omitempty" name:"NumRecentEvents"`

	// <p>事件起始时间，RFC3339 格式，最大长度 64 字符。</p>
	AfterTimestamp *string `json:"AfterTimestamp,omitnil,omitempty" name:"AfterTimestamp"`
}

type DescribeSessionRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话所属空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>用户 ID。可通过调用方业务系统接口获取。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>会话 ID。可通过 CreateSession 或 DescribeSessions 接口获取。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Agent ID。可选。</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>返回最近事件数量，默认为 0，最大值为 200。</p>
	NumRecentEvents *int64 `json:"NumRecentEvents,omitnil,omitempty" name:"NumRecentEvents"`

	// <p>事件起始时间，RFC3339 格式，最大长度 64 字符。</p>
	AfterTimestamp *string `json:"AfterTimestamp,omitnil,omitempty" name:"AfterTimestamp"`
}

func (r *DescribeSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "UserId")
	delete(f, "SessionId")
	delete(f, "AgentId")
	delete(f, "NumRecentEvents")
	delete(f, "AfterTimestamp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionResponseParams struct {
	// <p>会话信息。</p>
	Session *SessionInfo `json:"Session,omitnil,omitempty" name:"Session"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSessionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSessionResponseParams `json:"Response"`
}

func (r *DescribeSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionSpaceRequestParams struct {
	// <p>需要查询的会话空间唯一标识。</p><p>入参限制：必填，不能为空。</p><p>可通过 CreateSessionSpace 或 DescribeSessionSpaces 获取，不应自行构造。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DescribeSessionSpaceRequest struct {
	*tchttp.BaseRequest
	
	// <p>需要查询的会话空间唯一标识。</p><p>入参限制：必填，不能为空。</p><p>可通过 CreateSessionSpace 或 DescribeSessionSpaces 获取，不应自行构造。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DescribeSessionSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSessionSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionSpaceResponseParams struct {
	// <p>查询到的会话空间信息。</p>
	SessionSpace *SessionSpaceInfo `json:"SessionSpace,omitnil,omitempty" name:"SessionSpace"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSessionSpaceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSessionSpaceResponseParams `json:"Response"`
}

func (r *DescribeSessionSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionSpacesRequestParams struct {
	// <p>分页查询的起始偏移量。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>单次分页查询返回的会话空间数量。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>会话空间筛选条件列表，支持按空间 ID 精确匹配、名称精确或模糊匹配、描述模糊匹配。同一 Filter 内多个 Values 之间为 OR，不同 Filter 之间为 AND。不传或传空数组时不增加筛选限制。</p><p>入参限制：Filter.Name 支持 space-id、name、name-like、description-like，不可重复。name 与 name-like 不可同时提供。Values 不可为空数组，筛选值不可为空或纯空白。匹配区分大小写，包含匹配中的 %、_ 按普通字符处理，不具有通配含义。</p><p>例如 Name 为 name-like，Values 为 [&quot;客服&quot;,&quot;测试&quot;]，表示查询名称包含“客服”或“测试”的会话空间。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSessionSpacesRequest struct {
	*tchttp.BaseRequest
	
	// <p>分页查询的起始偏移量。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>单次分页查询返回的会话空间数量。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>会话空间筛选条件列表，支持按空间 ID 精确匹配、名称精确或模糊匹配、描述模糊匹配。同一 Filter 内多个 Values 之间为 OR，不同 Filter 之间为 AND。不传或传空数组时不增加筛选限制。</p><p>入参限制：Filter.Name 支持 space-id、name、name-like、description-like，不可重复。name 与 name-like 不可同时提供。Values 不可为空数组，筛选值不可为空或纯空白。匹配区分大小写，包含匹配中的 %、_ 按普通字符处理，不具有通配含义。</p><p>例如 Name 为 name-like，Values 为 [&quot;客服&quot;,&quot;测试&quot;]，表示查询名称包含“客服”或“测试”的会话空间。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeSessionSpacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionSpacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSessionSpacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionSpacesResponseParams struct {
	// <p>会话空间列表。</p>
	SessionSpaces []*SessionSpaceInfo `json:"SessionSpaces,omitnil,omitempty" name:"SessionSpaces"`

	// <p>满足查询条件的会话空间总数。</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSessionSpacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSessionSpacesResponseParams `json:"Response"`
}

func (r *DescribeSessionSpacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionSpacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionsRequestParams struct {
	// <p>查询的会话空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Agent ID 列表，最多支持 100 个。</p>
	//
	// Deprecated: AgentIds is deprecated.
	AgentIds []*string `json:"AgentIds,omitnil,omitempty" name:"AgentIds"`

	// <p>用户 ID 列表，最多支持 100 个。</p>
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// <p>分页偏移量，默认为 0。</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认为 20，最大值为 100。</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>会话 ID 列表，最多支持 100 个。</p>
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`

	// <p>会话筛选条件列表，支持 Metadata 精确匹配、标题精确匹配和标题模糊匹配。同一 Filter 内多个 Values 之间为 OR，不同 Filter 之间为 AND。不传或传空数组时不增加筛选限制。</p><p>入参限制：最多传入 10 个 Filter，每个 Filter 最多支持 100 个 Values。Filter.Name 不可重复，支持 metadata:MetadataKey、title、title-like；title 与 title-like 不可同时提供。标题筛选值不可为空或纯空白。匹配区分大小写，标题包含匹配中的 %、_ 按普通字符处理，不具有通配含义。</p><p>例如 Name 为 title-like，Values 为 [&quot;客服&quot;,&quot;测试&quot;]，表示查询标题包含“客服”或“测试”的会话。Name 为 metadata:env，Values 为 [&quot;dev&quot;,&quot;test&quot;]，表示按 Metadata env 的值精确筛选。标题条件与 Metadata、SessionIds、UserIds 筛选条件可组合使用，条件之间为 AND。筛选在分页前执行，TotalCount 为符合条件的会话总数。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSessionsRequest struct {
	*tchttp.BaseRequest
	
	// <p>查询的会话空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Agent ID 列表，最多支持 100 个。</p>
	AgentIds []*string `json:"AgentIds,omitnil,omitempty" name:"AgentIds"`

	// <p>用户 ID 列表，最多支持 100 个。</p>
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// <p>分页偏移量，默认为 0。</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认为 20，最大值为 100。</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>会话 ID 列表，最多支持 100 个。</p>
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`

	// <p>会话筛选条件列表，支持 Metadata 精确匹配、标题精确匹配和标题模糊匹配。同一 Filter 内多个 Values 之间为 OR，不同 Filter 之间为 AND。不传或传空数组时不增加筛选限制。</p><p>入参限制：最多传入 10 个 Filter，每个 Filter 最多支持 100 个 Values。Filter.Name 不可重复，支持 metadata:MetadataKey、title、title-like；title 与 title-like 不可同时提供。标题筛选值不可为空或纯空白。匹配区分大小写，标题包含匹配中的 %、_ 按普通字符处理，不具有通配含义。</p><p>例如 Name 为 title-like，Values 为 [&quot;客服&quot;,&quot;测试&quot;]，表示查询标题包含“客服”或“测试”的会话。Name 为 metadata:env，Values 为 [&quot;dev&quot;,&quot;test&quot;]，表示按 Metadata env 的值精确筛选。标题条件与 Metadata、SessionIds、UserIds 筛选条件可组合使用，条件之间为 AND。筛选在分页前执行，TotalCount 为符合条件的会话总数。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "AgentIds")
	delete(f, "UserIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SessionIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionsResponseParams struct {
	// <p>符合条件的会话总数。</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>会话列表。</p>
	Sessions []*SessionInfo `json:"Sessions,omitnil,omitempty" name:"Sessions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSessionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSessionsResponseParams `json:"Response"`
}

func (r *DescribeSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnvVar struct {
	// 环境变量名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境变量值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type EventActionsInfo struct {
	// 状态增量，JSON 字符串，最大长度 8192 字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StateDelta *string `json:"StateDelta,omitnil,omitempty" name:"StateDelta"`
}

type EventContentInfo struct {
	// 角色，最大长度 64 字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 内容片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parts []*EventPartInfo `json:"Parts,omitnil,omitempty" name:"Parts"`
}

type EventInfo struct {
	// <p>事件 ID。为空时由服务生成。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// <p>调用 ID，最大长度 128 字符。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationId *string `json:"InvocationId,omitnil,omitempty" name:"InvocationId"`

	// <p>事件作者，最大长度 128 字符。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Author *string `json:"Author,omitnil,omitempty" name:"Author"`

	// <p>事件内容。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *EventContentInfo `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>事件动作信息。StateDelta 为 JSON 对象字符串</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Actions *EventActionsInfo `json:"Actions,omitnil,omitempty" name:"Actions"`

	// <p>事件元数据。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metadata *string `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// <p>事件扩展信息 JSON 对象字符串，最大长度 8192 字符。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extensions *string `json:"Extensions,omitnil,omitempty" name:"Extensions"`

	// <p>错误码，最大长度 128 字符。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// <p>错误信息，最大长度 2048 字符。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// <p>事件时间。</p>
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type EventPartInfo struct {
	// 文本内容，最大长度 8192 字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 是否为思考内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Thought *bool `json:"Thought,omitnil,omitempty" name:"Thought"`

	// 工具调用信息，JSON 字符串，最大长度 8192 字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionCall *string `json:"FunctionCall,omitnil,omitempty" name:"FunctionCall"`

	// 工具返回信息，JSON 字符串，最大长度 8192 字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionResponse *string `json:"FunctionResponse,omitnil,omitempty" name:"FunctionResponse"`

	// 内联数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InlineData *InlineDataInfo `json:"InlineData,omitnil,omitempty" name:"InlineData"`
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type GetSkillPackageDownloadURLRequestParams struct {

}

type GetSkillPackageDownloadURLRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetSkillPackageDownloadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSkillPackageDownloadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSkillPackageDownloadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSkillPackageDownloadURLResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSkillPackageDownloadURLResponse struct {
	*tchttp.BaseResponse
	Response *GetSkillPackageDownloadURLResponseParams `json:"Response"`
}

func (r *GetSkillPackageDownloadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSkillPackageDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSkillPackageUploadURLRequestParams struct {

}

type GetSkillPackageUploadURLRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetSkillPackageUploadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSkillPackageUploadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSkillPackageUploadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSkillPackageUploadURLResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSkillPackageUploadURLResponse struct {
	*tchttp.BaseResponse
	Response *GetSkillPackageUploadURLResponseParams `json:"Response"`
}

func (r *GetSkillPackageUploadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSkillPackageUploadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HttpGetAction struct {
	// 路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 协议
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`
}

type ImageStorageSource struct {
	// <p>镜像地址</p>
	Reference *string `json:"Reference,omitnil,omitempty" name:"Reference"`

	// <p>镜像仓库类型：<code>enterprise</code>、<code>personal</code>。</p>
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`

	// <p>镜像内部的路径</p>
	SubPath *string `json:"SubPath,omitnil,omitempty" name:"SubPath"`

	// <p>镜像 Digest，请求时无需传入</p>
	Digest *string `json:"Digest,omitnil,omitempty" name:"Digest"`
}

type InlineDataInfo struct {
	// 媒体类型，最大长度 128 字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MimeType *string `json:"MimeType,omitnil,omitempty" name:"MimeType"`

	// Base64 编码数据，最大长度 8192 字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

type LifecycleConfiguration struct {
	// <p>Sandbox Instance 没有活跃 Deployment 请求或连接后进入 IdleAction 的秒数，必须大于等于 30。</p>
	IdleTimeoutSeconds *int64 `json:"IdleTimeoutSeconds,omitnil,omitempty" name:"IdleTimeoutSeconds"`

	// <p>空闲处理动作。</p><p>枚举值：</p><ul><li>STOP：停止并释放 Sandbox Instance。</li><li>PAUSE：暂停并保留 Sandbox Instance 状态。</li></ul>
	IdleAction *string `json:"IdleAction,omitnil,omitempty" name:"IdleAction"`
}

type LogConfiguration struct {
	// <p>日志推送CLS的配置。</p>
	CLSConfig *CLSConfig `json:"CLSConfig,omitnil,omitempty" name:"CLSConfig"`

	// <p>日志源配置</p>
	LogSources *LogSources `json:"LogSources,omitnil,omitempty" name:"LogSources"`
}

type LogSources struct {
	// <p>需要采集的日志文件路径，必须是 /logs/ 目录下的文件，不支持子目录，最大支持 10 个文件。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Files []*string `json:"Files,omitnil,omitempty" name:"Files"`
}

type MetadataVar struct {
	// <p>元数据名</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>元数据值</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ModifyDeploymentRequestParams struct {
	// <p>待修改的 Deployment ID。</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`

	// <p>完整替换伸缩配置；提供时必须包含全部三个成员。</p>
	ScalingConfiguration *ScalingConfiguration `json:"ScalingConfiguration,omitnil,omitempty" name:"ScalingConfiguration"`

	// <p>完整替换生命周期配置；提供时必须包含全部两个成员。</p>
	LifecycleConfiguration *LifecycleConfiguration `json:"LifecycleConfiguration,omitnil,omitempty" name:"LifecycleConfiguration"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyDeploymentRequest struct {
	*tchttp.BaseRequest
	
	// <p>待修改的 Deployment ID。</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`

	// <p>完整替换伸缩配置；提供时必须包含全部三个成员。</p>
	ScalingConfiguration *ScalingConfiguration `json:"ScalingConfiguration,omitnil,omitempty" name:"ScalingConfiguration"`

	// <p>完整替换生命周期配置；提供时必须包含全部两个成员。</p>
	LifecycleConfiguration *LifecycleConfiguration `json:"LifecycleConfiguration,omitnil,omitempty" name:"LifecycleConfiguration"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *ModifyDeploymentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeploymentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeploymentId")
	delete(f, "ScalingConfiguration")
	delete(f, "LifecycleConfiguration")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeploymentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeploymentResponseParams struct {
	// <p>修改后的完整 Deployment。</p>
	Deployment *Deployment `json:"Deployment,omitnil,omitempty" name:"Deployment"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDeploymentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeploymentResponseParams `json:"Response"`
}

func (r *ModifyDeploymentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySessionRequestParams struct {
	// <p>会话所属的 SessionSpace ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>会话所属的用户 ID。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>待修改的会话 ID。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>修改后的会话标题。</p><p>入参限制：本参数可选，最大长度为 255 个字符。</p><p>不传表示保持原会话标题不变，传空字符串表示清空会话标题。Title 与 Metadata 至少传入一项。</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>修改后的完整会话元数据，以键值对数组形式表示。</p><p>入参限制：本参数可选，最多支持 64 项。Name 不能为空或重复，最大长度为 253 字节；Value 最大长度为 1024 字节，允许为空字符串。Metadata 序列化后的总大小不能超过 64 KiB。</p><p>不传表示保持原 Metadata 不变；传空数组表示清空全部 Metadata；传非空数组表示使用传入内容全量覆盖原 Metadata。Metadata 与 Title 至少传入一项。</p>
	Metadata []*MetadataVar `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

type ModifySessionRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话所属的 SessionSpace ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>会话所属的用户 ID。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>待修改的会话 ID。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>修改后的会话标题。</p><p>入参限制：本参数可选，最大长度为 255 个字符。</p><p>不传表示保持原会话标题不变，传空字符串表示清空会话标题。Title 与 Metadata 至少传入一项。</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>修改后的完整会话元数据，以键值对数组形式表示。</p><p>入参限制：本参数可选，最多支持 64 项。Name 不能为空或重复，最大长度为 253 字节；Value 最大长度为 1024 字节，允许为空字符串。Metadata 序列化后的总大小不能超过 64 KiB。</p><p>不传表示保持原 Metadata 不变；传空数组表示清空全部 Metadata；传非空数组表示使用传入内容全量覆盖原 Metadata。Metadata 与 Title 至少传入一项。</p>
	Metadata []*MetadataVar `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

func (r *ModifySessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "UserId")
	delete(f, "SessionId")
	delete(f, "Title")
	delete(f, "Metadata")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySessionResponseParams struct {
	// <p>修改后的完整会话信息。</p>
	Session *SessionInfo `json:"Session,omitnil,omitempty" name:"Session"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySessionResponse struct {
	*tchttp.BaseResponse
	Response *ModifySessionResponseParams `json:"Response"`
}

func (r *ModifySessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySessionSpaceRequestParams struct {
	// <p>需要修改的会话空间唯一标识。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>修改后的会话空间名称。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>修改后的会话空间描述。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifySessionSpaceRequest struct {
	*tchttp.BaseRequest
	
	// <p>需要修改的会话空间唯一标识。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>修改后的会话空间名称。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>修改后的会话空间描述。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifySessionSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySessionSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySessionSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySessionSpaceResponseParams struct {
	// <p>修改后的会话空间信息。</p>
	SessionSpace *SessionSpaceInfo `json:"SessionSpace,omitnil,omitempty" name:"SessionSpace"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySessionSpaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifySessionSpaceResponseParams `json:"Response"`
}

func (r *ModifySessionSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySessionSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountOption struct {
	// 指定沙箱工具中的存储配置名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 沙箱实例本地挂载路径（可选），默认继承工具中的存储配置
	MountPath *string `json:"MountPath,omitnil,omitempty" name:"MountPath"`

	// 沙箱实例存储挂载子路径（可选）
	SubPath *string `json:"SubPath,omitnil,omitempty" name:"SubPath"`

	// 沙箱实例存储挂载读写权限（可选），默认继承工具存储配置
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`
}

type NetworkConfiguration struct {
	// 网络模式（当前支持 PUBLIC, VPC, SANDBOX）
	NetworkMode *string `json:"NetworkMode,omitnil,omitempty" name:"NetworkMode"`

	// VPC网络相关配置
	VpcConfig *VPCConfig `json:"VpcConfig,omitnil,omitempty" name:"VpcConfig"`
}

type OSWorldConfiguration struct {
	// <p>指定内置 OSWorld 版本</p><p>枚举值：</p><ul><li>osworld1： osworld v1</li><li>osworld2： osworld v2</li></ul><p>默认值：osworld1</p>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

// Predefined struct for user
type PauseSandboxInstanceRequestParams struct {
	// <p>沙箱实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>可选。带内存暂停，恢复后保留进程和内存状态。true=带内存；false=仅磁盘；不传=系统默认（当前默认 true，带内存）。</p>
	Memory *bool `json:"Memory,omitnil,omitempty" name:"Memory"`
}

type PauseSandboxInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>沙箱实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>可选。带内存暂停，恢复后保留进程和内存状态。true=带内存；false=仅磁盘；不传=系统默认（当前默认 true，带内存）。</p>
	Memory *bool `json:"Memory,omitnil,omitempty" name:"Memory"`
}

func (r *PauseSandboxInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseSandboxInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseSandboxInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseSandboxInstanceResponseParams struct {
	// <p>目标沙箱实例当前的状态</p><p>枚举值：</p><ul><li>PAUSING： 正在暂停中</li><li>PAUSED： 已暂停</li><li>PAUSE_FAILED： 暂停失败</li></ul>
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseSandboxInstanceResponse struct {
	*tchttp.BaseResponse
	Response *PauseSandboxInstanceResponseParams `json:"Response"`
}

func (r *PauseSandboxInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseSandboxInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PortConfiguration struct {
	// 端口名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

// Predefined struct for user
type PreviewRegistryRecordRequestParams struct {

}

type PreviewRegistryRecordRequest struct {
	*tchttp.BaseRequest
	
}

func (r *PreviewRegistryRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewRegistryRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PreviewRegistryRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PreviewRegistryRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PreviewRegistryRecordResponse struct {
	*tchttp.BaseResponse
	Response *PreviewRegistryRecordResponseParams `json:"Response"`
}

func (r *PreviewRegistryRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewRegistryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProbeConfiguration struct {
	// HTTP GET 探测配置
	HttpGet *HttpGetAction `json:"HttpGet,omitnil,omitempty" name:"HttpGet"`

	// 健康检查就绪超时
	ReadyTimeoutMs *int64 `json:"ReadyTimeoutMs,omitnil,omitempty" name:"ReadyTimeoutMs"`

	// 健康检查单次探测超时
	ProbeTimeoutMs *int64 `json:"ProbeTimeoutMs,omitnil,omitempty" name:"ProbeTimeoutMs"`

	// 健康检查间隔
	ProbePeriodMs *int64 `json:"ProbePeriodMs,omitnil,omitempty" name:"ProbePeriodMs"`

	// 健康检查成功阈值
	SuccessThreshold *int64 `json:"SuccessThreshold,omitnil,omitempty" name:"SuccessThreshold"`

	// 健康检查失败阈值
	FailureThreshold *int64 `json:"FailureThreshold,omitnil,omitempty" name:"FailureThreshold"`
}

type QuotaGroupOverview struct {
	// <p>配额组关联的标签键值</p>
	Tag *Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// <p>配额组名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>配额组各资源维度的配额上限</p>
	Quota *QuotaResourceInfo `json:"Quota,omitnil,omitempty" name:"Quota"`

	// <p>配额组各资源维度的当前用量</p>
	Usage *QuotaResourceInfo `json:"Usage,omitnil,omitempty" name:"Usage"`

	// <p>创建时间</p><p>参数格式：RFC3339 格式</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>最后更新时间</p><p>参数格式：RFC3339 格式</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type QuotaResourceInfo struct {
	// <p>沙箱工具配额或当前用量</p><p>单位：个</p>
	SandboxTools *int64 `json:"SandboxTools,omitnil,omitempty" name:"SandboxTools"`

	// <p>沙箱实例配额或当前用量</p><p>单位：个</p>
	SandboxInstances *int64 `json:"SandboxInstances,omitnil,omitempty" name:"SandboxInstances"`

	// <p>暂停实例配额或当前用量</p><p>单位：个</p>
	PausedInstances *int64 `json:"PausedInstances,omitnil,omitempty" name:"PausedInstances"`

	// <p>暂停实例配额或当前用量。目前只在主账号中返回</p><p>单位：核</p>
	CPUCores *float64 `json:"CPUCores,omitnil,omitempty" name:"CPUCores"`

	// <p>内存配额或当前用量</p><p>单位：GiB</p>
	MemoryGiB *float64 `json:"MemoryGiB,omitnil,omitempty" name:"MemoryGiB"`
}

// Predefined struct for user
type RejectRegistryRecordRequestParams struct {

}

type RejectRegistryRecordRequest struct {
	*tchttp.BaseRequest
	
}

func (r *RejectRegistryRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectRegistryRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RejectRegistryRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectRegistryRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RejectRegistryRecordResponse struct {
	*tchttp.BaseResponse
	Response *RejectRegistryRecordResponseParams `json:"Response"`
}

func (r *RejectRegistryRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectRegistryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceConfiguration struct {
	// <p>cpu 资源量</p>
	CPU *string `json:"CPU,omitnil,omitempty" name:"CPU"`

	// <p>内存资源量</p>
	Memory *string `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>自定义磁盘大小</p><p>枚举值：</p><ul><li>1Gi： 1Gi</li><li>5Gi： 5Gi</li><li>10Gi： 10Gi</li><li>20Gi： 20Gi</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Storage *string `json:"Storage,omitnil,omitempty" name:"Storage"`
}

// Predefined struct for user
type ResumeSandboxInstanceRequestParams struct {
	// <p>沙箱实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>超时时间，超过这个时间就自动回收实例。支持格式：5m、300s、1h 等，默认 5m。最小 30s，最大 24h</p>
	Timeout *string `json:"Timeout,omitnil,omitempty" name:"Timeout"`
}

type ResumeSandboxInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>沙箱实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>超时时间，超过这个时间就自动回收实例。支持格式：5m、300s、1h 等，默认 5m。最小 30s，最大 24h</p>
	Timeout *string `json:"Timeout,omitnil,omitempty" name:"Timeout"`
}

func (r *ResumeSandboxInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeSandboxInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Timeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeSandboxInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeSandboxInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeSandboxInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ResumeSandboxInstanceResponseParams `json:"Response"`
}

func (r *ResumeSandboxInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeSandboxInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SandboxInstance struct {
	// <p>沙箱实例唯一标识符</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>所属沙箱工具 ID</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>所属沙箱工具名称</p>
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// <p>实例状态：STARTING（启动中）、RUNNING（运行中）、STOPPING（停止中）、STOPPED（已停止）、STOP_FAILED（停止失败）、FAILED（失败状态）</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>是否常驻实例</p>
	Persistent *bool `json:"Persistent,omitnil,omitempty" name:"Persistent"`

	// <p>超时时间（秒），null 表示无超时设置</p>
	TimeoutSeconds *uint64 `json:"TimeoutSeconds,omitnil,omitempty" name:"TimeoutSeconds"`

	// <p>过期时间（ISO 8601 格式），null 表示无过期时间</p>
	ExpiresAt *string `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// <p>停止原因：manual（手动）、timeout（超时）、error（错误）、system（系统），仅在状态为 STOPPED、STOP_FAILED 或 FAILED 时有值。当 provider 停止失败时，状态为 STOP_FAILED，原因为 error</p>
	StopReason *string `json:"StopReason,omitnil,omitempty" name:"StopReason"`

	// <p>创建时间（ISO 8601 格式）</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间（ISO 8601 格式）</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>存储挂载选项</p>
	MountOptions []*MountOption `json:"MountOptions,omitnil,omitempty" name:"MountOptions"`

	// <p>沙箱实例自定义配置</p>
	CustomConfiguration *CustomConfigurationDetail `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`

	// <p>桌面电脑环境类沙箱配置</p>
	ComputerConfiguration *ComputerConfiguration `json:"ComputerConfiguration,omitnil,omitempty" name:"ComputerConfiguration"`

	// <p>网络模式</p><p>枚举值：</p><ul><li>PUBLIC： 公网访问</li><li>SANDBOX： 无网络</li><li>INTERNAL_SERVICE： 腾讯云内部公共服务</li></ul><p>可以覆盖工具级别的网络配置。但如果一个工具本身就不支持 VPC 网络，那么即便在实例设置里选了 VPC 模式，也是无效的</p>
	NetworkMode *string `json:"NetworkMode,omitnil,omitempty" name:"NetworkMode"`

	// <p>沙箱实例元数据</p>
	Metadata []*MetadataVar `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// <p>沙箱访问认证模式</p><p>枚举值：</p><ul><li>DEFAULT： 默认，即 TOKEN 认证</li><li>TOKEN： Token认证，即所有端口访问都需携带TOKEN</li><li>NONE： 免认证，即所有端口访问无需携带TOKEN</li><li>PUBLIC： 公开模式，即ENVD管理端口（49983）访问需携带TOKEN，其他端口无需携带TOKEN</li></ul><p>默认值：DEFAULT</p>
	AuthMode *string `json:"AuthMode,omitnil,omitempty" name:"AuthMode"`
}

type SandboxTool struct {
	// <p>沙箱工具唯一标识符</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>沙箱工具名称，长度 1-50 字符，支持中英文、数字、下划线。同一 AppId 下沙箱工具名称必须唯一</p>
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// <p>沙箱工具类型，取值：browser（浏览器工具）、code-interpreter（代码解释器工具）、computer（计算机控制工具）、mobile（移动设备工具）</p>
	ToolType *string `json:"ToolType,omitnil,omitempty" name:"ToolType"`

	// <p>沙箱工具状态，取值：CREATING（创建中）、ACTIVE（可用）、DELETING（删除中）、FAILED（失败）</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>沙箱工具描述信息，最大长度 200 字符</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>是否常驻沙箱</p>
	Persistent *bool `json:"Persistent,omitnil,omitempty" name:"Persistent"`

	// <p>默认超时时间，支持格式：5m、300s、1h 等，不指定则使用系统默认值（5 分钟）。最大 24 小时</p>
	DefaultTimeoutSeconds *uint64 `json:"DefaultTimeoutSeconds,omitnil,omitempty" name:"DefaultTimeoutSeconds"`

	// <p>网络配置</p>
	NetworkConfiguration *NetworkConfiguration `json:"NetworkConfiguration,omitnil,omitempty" name:"NetworkConfiguration"`

	// <p>标签规格，包含资源标签绑定关系。用于为沙箱工具绑定标签，支持多种资源类型的标签绑定</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>沙箱工具创建时间，格式：ISO8601</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>沙箱工具更新时间，格式：ISO8601</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>沙箱工具绑定角色ARN</p>
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// <p>沙箱工具中实例存储挂载配置</p>
	StorageMounts []*StorageMount `json:"StorageMounts,omitnil,omitempty" name:"StorageMounts"`

	// <p>沙箱工具自定义配置</p>
	CustomConfiguration *CustomConfigurationDetail `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`

	// <p>沙箱工具日志推送相关配置</p>
	LogConfiguration *LogConfiguration `json:"LogConfiguration,omitnil,omitempty" name:"LogConfiguration"`

	// <p>桌面电脑环境类沙箱配置</p>
	ComputerConfiguration *ComputerConfiguration `json:"ComputerConfiguration,omitnil,omitempty" name:"ComputerConfiguration"`

	// <p>用于说明沙箱工具处于该状态的原因</p>
	StatusReason *string `json:"StatusReason,omitnil,omitempty" name:"StatusReason"`
}

type ScalingConfiguration struct {
	// <p>活跃 Sandbox Instance 下限，必须大于等于 0。</p>
	MinInstanceCount *int64 `json:"MinInstanceCount,omitnil,omitempty" name:"MinInstanceCount"`

	// <p>活跃 Sandbox Instance 上限，必须大于等于 1，并且不小于 MinInstanceCount。</p>
	MaxInstanceCount *int64 `json:"MaxInstanceCount,omitnil,omitempty" name:"MaxInstanceCount"`

	// <p>每个活跃 Sandbox Instance 同时持有的 Deployment 请求或连接 Lease 上限，必须大于等于 1。</p>
	MaxInstanceRequestConcurrency *int64 `json:"MaxInstanceRequestConcurrency,omitnil,omitempty" name:"MaxInstanceRequestConcurrency"`
}

type SessionInfo struct {
	// <p>会话 ID。</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>会话所属空间 ID。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Session 快照状态</p>
	State *SessionState `json:"State,omitnil,omitempty" name:"State"`

	// <p>会话元数据，以键值对数组形式表示。每个元素包含 Metadata 名称和对应值，最多支持 64 项。</p>
	Metadata []*MetadataVar `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// <p>Agent ID。</p>
	//
	// Deprecated: AgentId is deprecated.
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>用户 ID。</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>会话标题。</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>事件数量。</p>
	EventCount *int64 `json:"EventCount,omitnil,omitempty" name:"EventCount"`

	// <p>创建时间。</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间。</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type SessionSpaceInfo struct {
	// <p>会话空间唯一标识，由服务端生成，最大长度为 128 个字符。调用方不应自行构造或解析。</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>会话空间名称，用于标识会话空间的业务用途，最大长度为 128 个字符。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>会话空间描述，用于说明业务用途和使用范围，最大长度为 512 个字符。为空时该字段可能不返回</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>会话空间当前状态。</p><p>枚举值：</p><ul><li>Active： 正常可用</li><li>Deleting： 正在删除</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>是否为系统默认会话空间。true 表示默认会话空间，false 表示普通会话空间。默认会话空间不允许删除。</p>
	Default *bool `json:"Default,omitnil,omitempty" name:"Default"`

	// <p>会话空间创建时间，采用 ISO 8601/RFC 3339 格式。</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>会话空间最后更新时间，采用 ISO 8601/RFC 3339 格式。</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type SessionState struct {
	// <p>自定义状态 JSON 对象字符串</p>
	CustomState *string `json:"CustomState,omitnil,omitempty" name:"CustomState"`
}

// Predefined struct for user
type StartSandboxInstanceRequestParams struct {
	// <p>沙箱工具 ID，与 ToolName 至少有一个要填</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>沙箱工具名称，与 ToolId 至少有一个要填</p>
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// <p>超时时间，超过这个时间就自动回收实例。支持格式：5m、300s、1h 等，默认 5m。最小 30s，最大 24h</p>
	Timeout *string `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>幂等性 Token，长度不超过 64 字符</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// <p>沙箱实例存储挂载配置</p>
	MountOptions []*MountOption `json:"MountOptions,omitnil,omitempty" name:"MountOptions"`

	// <p>沙箱实例自定义配置</p>
	CustomConfiguration *CustomConfiguration `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`

	// <p>沙箱访问认证模式</p><p>枚举值：</p><ul><li>DEFAULT： 默认，即TOKEN认证</li><li>TOKEN： Token认证，即所有端口访问都需携带Token</li><li>NONE： 免认证，即所有端口访问无需携带Token</li><li>PUBLIC： 公开模式，即ENVD管理端口（49983）访问需携带Token，其他端口无需携带Token</li></ul><p>默认值：DEFAULT</p>
	AuthMode *string `json:"AuthMode,omitnil,omitempty" name:"AuthMode"`

	// <p>沙箱元数据</p>
	Metadata []*MetadataVar `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

type StartSandboxInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>沙箱工具 ID，与 ToolName 至少有一个要填</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>沙箱工具名称，与 ToolId 至少有一个要填</p>
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// <p>超时时间，超过这个时间就自动回收实例。支持格式：5m、300s、1h 等，默认 5m。最小 30s，最大 24h</p>
	Timeout *string `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>幂等性 Token，长度不超过 64 字符</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// <p>沙箱实例存储挂载配置</p>
	MountOptions []*MountOption `json:"MountOptions,omitnil,omitempty" name:"MountOptions"`

	// <p>沙箱实例自定义配置</p>
	CustomConfiguration *CustomConfiguration `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`

	// <p>沙箱访问认证模式</p><p>枚举值：</p><ul><li>DEFAULT： 默认，即TOKEN认证</li><li>TOKEN： Token认证，即所有端口访问都需携带Token</li><li>NONE： 免认证，即所有端口访问无需携带Token</li><li>PUBLIC： 公开模式，即ENVD管理端口（49983）访问需携带Token，其他端口无需携带Token</li></ul><p>默认值：DEFAULT</p>
	AuthMode *string `json:"AuthMode,omitnil,omitempty" name:"AuthMode"`

	// <p>沙箱元数据</p>
	Metadata []*MetadataVar `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

func (r *StartSandboxInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartSandboxInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ToolId")
	delete(f, "ToolName")
	delete(f, "Timeout")
	delete(f, "ClientToken")
	delete(f, "MountOptions")
	delete(f, "CustomConfiguration")
	delete(f, "AuthMode")
	delete(f, "Metadata")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartSandboxInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartSandboxInstanceResponseParams struct {
	// <p>创建的沙箱实例完整信息</p>
	Instance *SandboxInstance `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartSandboxInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StartSandboxInstanceResponseParams `json:"Response"`
}

func (r *StartSandboxInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartSandboxInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSandboxInstanceRequestParams struct {
	// 沙箱实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StopSandboxInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 沙箱实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *StopSandboxInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSandboxInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopSandboxInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSandboxInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopSandboxInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StopSandboxInstanceResponseParams `json:"Response"`
}

func (r *StopSandboxInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSandboxInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageMount struct {
	// <p>存储挂载配置名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>存储配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSource *StorageSource `json:"StorageSource,omitnil,omitempty" name:"StorageSource"`

	// <p>沙箱实例本地挂载路径</p>
	MountPath *string `json:"MountPath,omitnil,omitempty" name:"MountPath"`

	// <p>存储挂载读写权限配置，默认为false</p>
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`
}

type StorageSource struct {
	// <p>对象存储桶配置</p>
	Cos *CosStorageSource `json:"Cos,omitnil,omitempty" name:"Cos"`

	// <p>镜像卷配置</p>
	Image *ImageStorageSource `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>文件存储配置</p>
	Cfs *CfsStorageSource `json:"Cfs,omitnil,omitempty" name:"Cfs"`

	// <p>AgentBucket 存储配置</p>
	AgentBucket *AgentBucketStorageSource `json:"AgentBucket,omitnil,omitempty" name:"AgentBucket"`
}

// Predefined struct for user
type SyncRegistryRecordRequestParams struct {

}

type SyncRegistryRecordRequest struct {
	*tchttp.BaseRequest
	
}

func (r *SyncRegistryRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncRegistryRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncRegistryRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncRegistryRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncRegistryRecordResponse struct {
	*tchttp.BaseResponse
	Response *SyncRegistryRecordResponseParams `json:"Response"`
}

func (r *SyncRegistryRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncRegistryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type UpdateRegistryRecordRequestParams struct {

}

type UpdateRegistryRecordRequest struct {
	*tchttp.BaseRequest
	
}

func (r *UpdateRegistryRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRegistryRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRegistryRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRegistryRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRegistryRecordResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRegistryRecordResponseParams `json:"Response"`
}

func (r *UpdateRegistryRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRegistryRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRegistryRequestParams struct {

}

type UpdateRegistryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *UpdateRegistryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRegistryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRegistryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRegistryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRegistryResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRegistryResponseParams `json:"Response"`
}

func (r *UpdateRegistryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRegistryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSandboxInstanceRequestParams struct {
	// <p>沙箱实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>新的超时时间（从设置时开始重新计算超时），支持格式：5m、300s、1h等。最小30s，最大24h。如果不指定则保持原有超时设置</p>
	Timeout *string `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>沙箱实例元数据</p>
	Metadata []*MetadataVar `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

type UpdateSandboxInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>沙箱实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>新的超时时间（从设置时开始重新计算超时），支持格式：5m、300s、1h等。最小30s，最大24h。如果不指定则保持原有超时设置</p>
	Timeout *string `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>沙箱实例元数据</p>
	Metadata []*MetadataVar `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

func (r *UpdateSandboxInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSandboxInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Timeout")
	delete(f, "Metadata")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSandboxInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSandboxInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSandboxInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSandboxInstanceResponseParams `json:"Response"`
}

func (r *UpdateSandboxInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSandboxInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSandboxToolRequestParams struct {
	// <p>沙箱工具ID</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>沙箱工具描述，最大长度200字符</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>网络配置</p>
	NetworkConfiguration *NetworkConfiguration `json:"NetworkConfiguration,omitnil,omitempty" name:"NetworkConfiguration"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>沙箱工具自定义配置</p>
	CustomConfiguration *CustomConfiguration `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`

	// <p>桌面电脑环境类沙箱配置</p>
	ComputerConfiguration *ComputerConfiguration `json:"ComputerConfiguration,omitnil,omitempty" name:"ComputerConfiguration"`
}

type UpdateSandboxToolRequest struct {
	*tchttp.BaseRequest
	
	// <p>沙箱工具ID</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>沙箱工具描述，最大长度200字符</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>网络配置</p>
	NetworkConfiguration *NetworkConfiguration `json:"NetworkConfiguration,omitnil,omitempty" name:"NetworkConfiguration"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>沙箱工具自定义配置</p>
	CustomConfiguration *CustomConfiguration `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`

	// <p>桌面电脑环境类沙箱配置</p>
	ComputerConfiguration *ComputerConfiguration `json:"ComputerConfiguration,omitnil,omitempty" name:"ComputerConfiguration"`
}

func (r *UpdateSandboxToolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSandboxToolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ToolId")
	delete(f, "Description")
	delete(f, "NetworkConfiguration")
	delete(f, "Tags")
	delete(f, "CustomConfiguration")
	delete(f, "ComputerConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSandboxToolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSandboxToolResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSandboxToolResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSandboxToolResponseParams `json:"Response"`
}

func (r *UpdateSandboxToolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSandboxToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VPCConfig struct {
	// <p>VPC子网ID列表</p>
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// <p>安全组ID列表</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type WAAConfiguration struct {
	// <p>自定义waa镜像ID</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}