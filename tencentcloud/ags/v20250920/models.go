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

// Predefined struct for user
type AcquireSandboxInstanceTokenRequestParams struct {
	// 沙箱实例ID，生成的访问Token将仅可用于访问此沙箱实例
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type AcquireSandboxInstanceTokenRequest struct {
	*tchttp.BaseRequest
	
	// 沙箱实例ID，生成的访问Token将仅可用于访问此沙箱实例
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
	// 访问Token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 过期时间
	ExpiresAt *string `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

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
type CreatePreCacheImageTaskRequestParams struct {
	// 镜像地址
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 镜像仓库类型：`enterprise`、`personal`。
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`
}

type CreatePreCacheImageTaskRequest struct {
	*tchttp.BaseRequest
	
	// 镜像地址
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 镜像仓库类型：`enterprise`、`personal`。
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
	// 镜像地址
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 镜像 Digest
	ImageDigest *string `json:"ImageDigest,omitnil,omitempty" name:"ImageDigest"`

	// 镜像仓库类型：`enterprise`、`personal`。
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
type CreateSandboxToolRequestParams struct {
	// 沙箱工具名称，长度 1-50 字符，支持英文、数字、下划线和连接线。同一 AppId 下沙箱工具名称必须唯一
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// 沙箱工具类型，目前支持：browser、code-interpreter、custom
	ToolType *string `json:"ToolType,omitnil,omitempty" name:"ToolType"`

	// 网络配置
	NetworkConfiguration *NetworkConfiguration `json:"NetworkConfiguration,omitnil,omitempty" name:"NetworkConfiguration"`

	// 沙箱工具描述，最大长度 200 字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 默认超时时间，支持格式：5m、300s、1h 等，不指定则使用系统默认值（5 分钟）。最大 24 小时
	DefaultTimeout *string `json:"DefaultTimeout,omitnil,omitempty" name:"DefaultTimeout"`

	// 标签规格，为沙箱工具绑定标签，支持多种资源类型的标签绑定
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 幂等性 Token，长度不超过 64 字符
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 角色ARN
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 沙箱工具存储配置
	StorageMounts []*StorageMount `json:"StorageMounts,omitnil,omitempty" name:"StorageMounts"`

	// 沙箱工具自定义配置
	CustomConfiguration *CustomConfiguration `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`
}

type CreateSandboxToolRequest struct {
	*tchttp.BaseRequest
	
	// 沙箱工具名称，长度 1-50 字符，支持英文、数字、下划线和连接线。同一 AppId 下沙箱工具名称必须唯一
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// 沙箱工具类型，目前支持：browser、code-interpreter、custom
	ToolType *string `json:"ToolType,omitnil,omitempty" name:"ToolType"`

	// 网络配置
	NetworkConfiguration *NetworkConfiguration `json:"NetworkConfiguration,omitnil,omitempty" name:"NetworkConfiguration"`

	// 沙箱工具描述，最大长度 200 字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 默认超时时间，支持格式：5m、300s、1h 等，不指定则使用系统默认值（5 分钟）。最大 24 小时
	DefaultTimeout *string `json:"DefaultTimeout,omitnil,omitempty" name:"DefaultTimeout"`

	// 标签规格，为沙箱工具绑定标签，支持多种资源类型的标签绑定
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 幂等性 Token，长度不超过 64 字符
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 角色ARN
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 沙箱工具存储配置
	StorageMounts []*StorageMount `json:"StorageMounts,omitnil,omitempty" name:"StorageMounts"`

	// 沙箱工具自定义配置
	CustomConfiguration *CustomConfiguration `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSandboxToolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSandboxToolResponseParams struct {
	// 创建的沙箱工具 ID
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

type CustomConfiguration struct {
	// 镜像地址
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 镜像仓库类型：`enterprise`、`personal`。
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`

	// 启动命令
	Command []*string `json:"Command,omitnil,omitempty" name:"Command"`

	// 启动参数
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`

	// 环境变量
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// 端口配置
	Ports []*PortConfiguration `json:"Ports,omitnil,omitempty" name:"Ports"`

	// 资源配置
	Resources *ResourceConfiguration `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 探针配置
	Probe *ProbeConfiguration `json:"Probe,omitnil,omitempty" name:"Probe"`
}

type CustomConfigurationDetail struct {
	// 镜像地址
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 镜像仓库类型：`TCR`、`CCR`。
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`

	// 镜像 Digest
	ImageDigest *string `json:"ImageDigest,omitnil,omitempty" name:"ImageDigest"`

	// 启动命令
	Command []*string `json:"Command,omitnil,omitempty" name:"Command"`

	// 启动参数
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`

	// 环境变量
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// 端口配置
	Ports []*PortConfiguration `json:"Ports,omitnil,omitempty" name:"Ports"`

	// 资源配置
	Resources *ResourceConfiguration `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 探针配置
	Probe *ProbeConfiguration `json:"Probe,omitnil,omitempty" name:"Probe"`
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
type DescribePreCacheImageTaskRequestParams struct {
	// 镜像地址
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 镜像 Digest
	ImageDigest *string `json:"ImageDigest,omitnil,omitempty" name:"ImageDigest"`

	// 镜像仓库类型：`enterprise`、`personal`。
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`
}

type DescribePreCacheImageTaskRequest struct {
	*tchttp.BaseRequest
	
	// 镜像地址
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 镜像 Digest
	ImageDigest *string `json:"ImageDigest,omitnil,omitempty" name:"ImageDigest"`

	// 镜像仓库类型：`enterprise`、`personal`。
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
	// 镜像地址
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 镜像 Digest
	ImageDigest *string `json:"ImageDigest,omitnil,omitempty" name:"ImageDigest"`

	// 镜像仓库类型：`enterprise`、`personal`。
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`

	// 镜像预热状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 镜像预热状态描述
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
type DescribeSandboxInstanceListRequestParams struct {
	// 沙箱实例ID列表，指定要查询的实例。如果为空则查询所有实例。最大支持100个ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 沙箱工具ID，指定时查询该沙箱模板下的实例，为空则查询所有沙箱模板的实例
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSandboxInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 沙箱实例ID列表，指定要查询的实例。如果为空则查询所有实例。最大支持100个ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 沙箱工具ID，指定时查询该沙箱模板下的实例，为空则查询所有沙箱模板的实例
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSandboxInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSandboxInstanceListResponseParams struct {
	// 沙箱实例列表
	InstanceSet []*SandboxInstance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 符合条件的实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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

type EnvVar struct {
	// 环境变量名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境变量值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
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
	// 镜像地址
	Reference *string `json:"Reference,omitnil,omitempty" name:"Reference"`

	// 镜像仓库类型：`enterprise`、`personal`。
	ImageRegistryType *string `json:"ImageRegistryType,omitnil,omitempty" name:"ImageRegistryType"`

	// 镜像内部的路径
	SubPath *string `json:"SubPath,omitnil,omitempty" name:"SubPath"`

	// 镜像 Digest，请求时无需传入
	Digest *string `json:"Digest,omitnil,omitempty" name:"Digest"`
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

type PortConfiguration struct {
	// 端口名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
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

type ResourceConfiguration struct {
	// cpu 资源量
	CPU *string `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 内存资源量
	Memory *string `json:"Memory,omitnil,omitempty" name:"Memory"`
}

type SandboxInstance struct {
	// 沙箱实例唯一标识符
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 所属沙箱工具 ID
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// 所属沙箱工具名称
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// 实例状态：STARTING（启动中）、RUNNING（运行中）、STOPPING（停止中）、STOPPED（已停止）、STOP_FAILED（停止失败）、FAILED（失败状态）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 超时时间（秒），null 表示无超时设置
	TimeoutSeconds *uint64 `json:"TimeoutSeconds,omitnil,omitempty" name:"TimeoutSeconds"`

	// 过期时间（ISO 8601 格式），null 表示无过期时间
	ExpiresAt *string `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// 停止原因：manual（手动）、timeout（超时）、error（错误）、system（系统），仅在状态为 STOPPED、STOP_FAILED 或 FAILED 时有值。当 provider 停止失败时，状态为 STOP_FAILED，原因为 error
	StopReason *string `json:"StopReason,omitnil,omitempty" name:"StopReason"`

	// 创建时间（ISO 8601 格式）
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间（ISO 8601 格式）
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 存储挂载选项
	MountOptions []*MountOption `json:"MountOptions,omitnil,omitempty" name:"MountOptions"`

	// 沙箱实例自定义配置
	CustomConfiguration *CustomConfigurationDetail `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`
}

type SandboxTool struct {
	// 沙箱工具唯一标识符
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// 沙箱工具名称，长度 1-50 字符，支持中英文、数字、下划线。同一 AppId 下沙箱工具名称必须唯一
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// 沙箱工具类型，取值：browser（浏览器工具）、code-interpreter（代码解释器工具）、computer（计算机控制工具）、mobile（移动设备工具）
	ToolType *string `json:"ToolType,omitnil,omitempty" name:"ToolType"`

	// 沙箱工具状态，取值：CREATING（创建中）、ACTIVE（可用）、DELETING（删除中）、FAILED（失败）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 沙箱工具描述信息，最大长度 200 字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 默认超时时间，支持格式：5m、300s、1h 等，不指定则使用系统默认值（5 分钟）。最大 24 小时
	DefaultTimeoutSeconds *uint64 `json:"DefaultTimeoutSeconds,omitnil,omitempty" name:"DefaultTimeoutSeconds"`

	// 网络配置
	NetworkConfiguration *NetworkConfiguration `json:"NetworkConfiguration,omitnil,omitempty" name:"NetworkConfiguration"`

	// 标签规格，包含资源标签绑定关系。用于为沙箱工具绑定标签，支持多种资源类型的标签绑定
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 沙箱工具创建时间，格式：ISO8601
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 沙箱工具更新时间，格式：ISO8601
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 沙箱工具绑定角色ARN
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 沙箱工具中实例存储挂载配置
	StorageMounts []*StorageMount `json:"StorageMounts,omitnil,omitempty" name:"StorageMounts"`

	// 沙箱工具自定义配置
	CustomConfiguration *CustomConfigurationDetail `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`
}

// Predefined struct for user
type StartSandboxInstanceRequestParams struct {
	// 沙箱工具 ID，与 ToolName 至少有一个要填
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// 沙箱工具名称，与 ToolId 至少有一个要填
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// 超时时间，超过这个时间就自动回收实例。支持格式：5m、300s、1h 等，默认 5m。最小 30s，最大 24h
	Timeout *string `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 幂等性 Token，长度不超过 64 字符
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 沙箱实例存储挂载配置
	MountOptions []*MountOption `json:"MountOptions,omitnil,omitempty" name:"MountOptions"`

	// 沙箱实例自定义配置
	CustomConfiguration *CustomConfiguration `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`
}

type StartSandboxInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 沙箱工具 ID，与 ToolName 至少有一个要填
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// 沙箱工具名称，与 ToolId 至少有一个要填
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// 超时时间，超过这个时间就自动回收实例。支持格式：5m、300s、1h 等，默认 5m。最小 30s，最大 24h
	Timeout *string `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 幂等性 Token，长度不超过 64 字符
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 沙箱实例存储挂载配置
	MountOptions []*MountOption `json:"MountOptions,omitnil,omitempty" name:"MountOptions"`

	// 沙箱实例自定义配置
	CustomConfiguration *CustomConfiguration `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartSandboxInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartSandboxInstanceResponseParams struct {
	// 创建的沙箱实例完整信息
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
	// 存储挂载配置名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 存储配置
	StorageSource *StorageSource `json:"StorageSource,omitnil,omitempty" name:"StorageSource"`

	// 沙箱实例本地挂载路径
	MountPath *string `json:"MountPath,omitnil,omitempty" name:"MountPath"`

	// 存储挂载读写权限配置，默认为false
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`
}

type StorageSource struct {
	// 对象存储桶配置
	Cos *CosStorageSource `json:"Cos,omitnil,omitempty" name:"Cos"`

	// 镜像卷配置
	Image *ImageStorageSource `json:"Image,omitnil,omitempty" name:"Image"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type UpdateSandboxInstanceRequestParams struct {
	// 沙箱实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新的超时时间（从设置时开始重新计算超时），支持格式：5m、300s、1h等。最小30s，最大24h。如果不指定则保持原有超时设置
	Timeout *string `json:"Timeout,omitnil,omitempty" name:"Timeout"`
}

type UpdateSandboxInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 沙箱实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新的超时时间（从设置时开始重新计算超时），支持格式：5m、300s、1h等。最小30s，最大24h。如果不指定则保持原有超时设置
	Timeout *string `json:"Timeout,omitnil,omitempty" name:"Timeout"`
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
	// 沙箱工具ID
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// 沙箱工具描述，最大长度200字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 网络配置
	NetworkConfiguration *NetworkConfiguration `json:"NetworkConfiguration,omitnil,omitempty" name:"NetworkConfiguration"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 沙箱工具自定义配置
	CustomConfiguration *CustomConfiguration `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`
}

type UpdateSandboxToolRequest struct {
	*tchttp.BaseRequest
	
	// 沙箱工具ID
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// 沙箱工具描述，最大长度200字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 网络配置
	NetworkConfiguration *NetworkConfiguration `json:"NetworkConfiguration,omitnil,omitempty" name:"NetworkConfiguration"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 沙箱工具自定义配置
	CustomConfiguration *CustomConfiguration `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`
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
	// VPC子网ID列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 安全组ID列表
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}