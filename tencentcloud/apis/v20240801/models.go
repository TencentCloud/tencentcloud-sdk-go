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

package v20240801

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AgentAppMcpServerDTO struct {
	// mcp server id
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 是否需要鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeedAuth *bool `json:"NeedAuth,omitnil,omitempty" name:"NeedAuth"`

	// 凭据代填的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentCredentialID *string `json:"AgentCredentialID,omitnil,omitempty" name:"AgentCredentialID"`

	// 应用为OAuth2认证时，sse模式请求mcp时的资源标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	SSEResourceIdentifier *string `json:"SSEResourceIdentifier,omitnil,omitempty" name:"SSEResourceIdentifier"`

	// 应用为OAuth2认证时，streamable模式请求mcp时的资源标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamableResourceIdentifier *string `json:"StreamableResourceIdentifier,omitnil,omitempty" name:"StreamableResourceIdentifier"`
}

type AgentAppMcpServerVO struct {
	// 绑定ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 需要认证
	NeedAuth *bool `json:"NeedAuth,omitnil,omitempty" name:"NeedAuth"`

	// 凭据ID
	AgentCredentialID *string `json:"AgentCredentialID,omitnil,omitempty" name:"AgentCredentialID"`

	// 凭据详情
	AgentCredentialVO *DescribeAgentCredentialResp `json:"AgentCredentialVO,omitnil,omitempty" name:"AgentCredentialVO"`

	// mcp详情
	McpServerVO *DescribeMcpServerResponseVO `json:"McpServerVO,omitnil,omitempty" name:"McpServerVO"`

	// 关联时间
	RelateTime *string `json:"RelateTime,omitnil,omitempty" name:"RelateTime"`

	// 应用为OAuth2认证时，sse模式请求mcp时的资源标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	SSEResourceIdentifier *string `json:"SSEResourceIdentifier,omitnil,omitempty" name:"SSEResourceIdentifier"`

	// 应用为OAuth2认证时，streamable模式请求mcp时的资源标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamableResourceIdentifier *string `json:"StreamableResourceIdentifier,omitnil,omitempty" name:"StreamableResourceIdentifier"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentAppID *string `json:"AgentAppID,omitnil,omitempty" name:"AgentAppID"`

	// mcp ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	McpServerID *string `json:"McpServerID,omitnil,omitempty" name:"McpServerID"`
}

type AgentAppModelServiceDTO struct {
	// 模型服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 是否开启流量控制
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 流量控制
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// 是否开启token控制
	TokenLimitStatus *bool `json:"TokenLimitStatus,omitnil,omitempty" name:"TokenLimitStatus"`

	// token控制
	TokenLimitConfig *TokenLimitConfigDTO `json:"TokenLimitConfig,omitnil,omitempty" name:"TokenLimitConfig"`
}

type AgentAppSecretKeyVO struct {
	// secret id
	SecretID *string `json:"SecretID,omitnil,omitempty" name:"SecretID"`

	// secret key
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`
}

type AgentCredentialContentDTO struct {
	// 如果认证类型为sts时，该项必填
	STSSystem *string `json:"STSSystem,omitnil,omitempty" name:"STSSystem"`

	// 如果认证类型为sts时，该项必填
	STSService *string `json:"STSService,omitnil,omitempty" name:"STSService"`

	// 如果认证类型为reqKey时，该项必填
	Headers []*AgentCredentialContentHeaderDTO `json:"Headers,omitnil,omitempty" name:"Headers"`
}

type AgentCredentialContentHeaderDTO struct {
	// 凭据header key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 凭据header value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type BindMcpSecurityRuleDTO struct {
	// 规则ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 处置动作
	Act *string `json:"Act,omitnil,omitempty" name:"Act"`
}

type BindMcpSecurityRuleVO struct {
	// 规则ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 规则类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 风险等级
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 版本号
	VersionNumber *string `json:"VersionNumber,omitnil,omitempty" name:"VersionNumber"`

	// 当前选择的处置动作
	Act *string `json:"Act,omitnil,omitempty" name:"Act"`

	// 支持的处置动作
	SupportActs []*string `json:"SupportActs,omitnil,omitempty" name:"SupportActs"`

	// 内容类型
	BodyType *string `json:"BodyType,omitnil,omitempty" name:"BodyType"`

	// icon类型
	IconType *string `json:"IconType,omitnil,omitempty" name:"IconType"`
}

// Predefined struct for user
type CreateAgentAppMcpServersRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 关联的mcp server
	McpServers []*AgentAppMcpServerDTO `json:"McpServers,omitnil,omitempty" name:"McpServers"`
}

type CreateAgentAppMcpServersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 关联的mcp server
	McpServers []*AgentAppMcpServerDTO `json:"McpServers,omitnil,omitempty" name:"McpServers"`
}

func (r *CreateAgentAppMcpServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentAppMcpServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	delete(f, "McpServers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentAppMcpServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentAppMcpServersResponseParams struct {
	// app id
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgentAppMcpServersResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentAppMcpServersResponseParams `json:"Response"`
}

func (r *CreateAgentAppMcpServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentAppMcpServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentAppModelServicesRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 关联的model service
	ModelServices []*AgentAppModelServiceDTO `json:"ModelServices,omitnil,omitempty" name:"ModelServices"`
}

type CreateAgentAppModelServicesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 关联的model service
	ModelServices []*AgentAppModelServiceDTO `json:"ModelServices,omitnil,omitempty" name:"ModelServices"`
}

func (r *CreateAgentAppModelServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentAppModelServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	delete(f, "ModelServices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentAppModelServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentAppModelServicesResponseParams struct {
	// app id
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgentAppModelServicesResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentAppModelServicesResponseParams `json:"Response"`
}

func (r *CreateAgentAppModelServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentAppModelServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentAppRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 认证类型
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// OAuth2资源服务器ID
	OAuth2ResourceServerID *string `json:"OAuth2ResourceServerID,omitnil,omitempty" name:"OAuth2ResourceServerID"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateAgentAppRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 认证类型
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// OAuth2资源服务器ID
	OAuth2ResourceServerID *string `json:"OAuth2ResourceServerID,omitnil,omitempty" name:"OAuth2ResourceServerID"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateAgentAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "Name")
	delete(f, "AuthType")
	delete(f, "OAuth2ResourceServerID")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAgentAppResp struct {
	// app id 
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 如果authType为apiKey时，返回该字段
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// 如果authType为secretKey时，返回该字段
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 如果authType为secretKey时，返回该字段
	SecretID *string `json:"SecretID,omitnil,omitempty" name:"SecretID"`
}

// Predefined struct for user
type CreateAgentAppResponseParams struct {
	// app id
	Data *CreateAgentAppResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgentAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentAppResponseParams `json:"Response"`
}

func (r *CreateAgentAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentCredentialRequestParams struct {

}

type CreateAgentCredentialRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateAgentCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentCredentialResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgentCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentCredentialResponseParams `json:"Response"`
}

func (r *CreateAgentCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMcpServerRequestParams struct {
	// 模式：proxy代理模式； wrap封装模式；
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 版本号：2024-11-05 2025-03-26
	McpVersion *string `json:"McpVersion,omitnil,omitempty" name:"McpVersion"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 封装服务列表
	WrapServices []*string `json:"WrapServices,omitnil,omitempty" name:"WrapServices"`

	// 负载方式，robin random consistentHash
	TargetSelect *string `json:"TargetSelect,omitnil,omitempty" name:"TargetSelect"`

	// 目标服务器
	TargetHosts []*TargetHostDTO `json:"TargetHosts,omitnil,omitempty" name:"TargetHosts"`

	// 后端协议：http https
	HttpProtocolType *string `json:"HttpProtocolType,omitnil,omitempty" name:"HttpProtocolType"`

	// 证书检查
	CheckTargetCertsError *bool `json:"CheckTargetCertsError,omitnil,omitempty" name:"CheckTargetCertsError"`

	// 目标路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// 流量控制开启状态
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 流量控制配置
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// IP白名单开启状态
	IpWhiteStatus *bool `json:"IpWhiteStatus,omitnil,omitempty" name:"IpWhiteStatus"`

	// IP白名单配置
	IpWhiteConfig *IpConfig `json:"IpWhiteConfig,omitnil,omitempty" name:"IpWhiteConfig"`

	// IP黑名单开启状态
	IpBlackStatus *bool `json:"IpBlackStatus,omitnil,omitempty" name:"IpBlackStatus"`

	// IP黑名单配置
	IpBlackConfig *IpConfig `json:"IpBlackConfig,omitnil,omitempty" name:"IpBlackConfig"`

	// 自定义host
	CustomHttpHost *string `json:"CustomHttpHost,omitnil,omitempty" name:"CustomHttpHost"`

	// Http 请求host类型：useRequestHost 保持源请求；host targetHost 修正为源站host；  customHost 自定义host
	HttpHostType *string `json:"HttpHostType,omitnil,omitempty" name:"HttpHostType"`

	// 请求的超时时间
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 安全规则集
	McpSecurityRules []*McpSecurityRule `json:"McpSecurityRules,omitnil,omitempty" name:"McpSecurityRules"`

	// 工具集配置（openapi时或许用的是）
	ToolConfigs []*ToolConfigDTO `json:"ToolConfigs,omitnil,omitempty" name:"ToolConfigs"`

	// 封装的API分组ID
	WrapPaasID *string `json:"WrapPaasID,omitnil,omitempty" name:"WrapPaasID"`

	// 插件配置
	PluginConfigs []*PluginConfigDTO `json:"PluginConfigs,omitnil,omitempty" name:"PluginConfigs"`
}

type CreateMcpServerRequest struct {
	*tchttp.BaseRequest
	
	// 模式：proxy代理模式； wrap封装模式；
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 版本号：2024-11-05 2025-03-26
	McpVersion *string `json:"McpVersion,omitnil,omitempty" name:"McpVersion"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 封装服务列表
	WrapServices []*string `json:"WrapServices,omitnil,omitempty" name:"WrapServices"`

	// 负载方式，robin random consistentHash
	TargetSelect *string `json:"TargetSelect,omitnil,omitempty" name:"TargetSelect"`

	// 目标服务器
	TargetHosts []*TargetHostDTO `json:"TargetHosts,omitnil,omitempty" name:"TargetHosts"`

	// 后端协议：http https
	HttpProtocolType *string `json:"HttpProtocolType,omitnil,omitempty" name:"HttpProtocolType"`

	// 证书检查
	CheckTargetCertsError *bool `json:"CheckTargetCertsError,omitnil,omitempty" name:"CheckTargetCertsError"`

	// 目标路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// 流量控制开启状态
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 流量控制配置
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// IP白名单开启状态
	IpWhiteStatus *bool `json:"IpWhiteStatus,omitnil,omitempty" name:"IpWhiteStatus"`

	// IP白名单配置
	IpWhiteConfig *IpConfig `json:"IpWhiteConfig,omitnil,omitempty" name:"IpWhiteConfig"`

	// IP黑名单开启状态
	IpBlackStatus *bool `json:"IpBlackStatus,omitnil,omitempty" name:"IpBlackStatus"`

	// IP黑名单配置
	IpBlackConfig *IpConfig `json:"IpBlackConfig,omitnil,omitempty" name:"IpBlackConfig"`

	// 自定义host
	CustomHttpHost *string `json:"CustomHttpHost,omitnil,omitempty" name:"CustomHttpHost"`

	// Http 请求host类型：useRequestHost 保持源请求；host targetHost 修正为源站host；  customHost 自定义host
	HttpHostType *string `json:"HttpHostType,omitnil,omitempty" name:"HttpHostType"`

	// 请求的超时时间
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 安全规则集
	McpSecurityRules []*McpSecurityRule `json:"McpSecurityRules,omitnil,omitempty" name:"McpSecurityRules"`

	// 工具集配置（openapi时或许用的是）
	ToolConfigs []*ToolConfigDTO `json:"ToolConfigs,omitnil,omitempty" name:"ToolConfigs"`

	// 封装的API分组ID
	WrapPaasID *string `json:"WrapPaasID,omitnil,omitempty" name:"WrapPaasID"`

	// 插件配置
	PluginConfigs []*PluginConfigDTO `json:"PluginConfigs,omitnil,omitempty" name:"PluginConfigs"`
}

func (r *CreateMcpServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMcpServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "McpVersion")
	delete(f, "InstanceID")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "WrapServices")
	delete(f, "TargetSelect")
	delete(f, "TargetHosts")
	delete(f, "HttpProtocolType")
	delete(f, "CheckTargetCertsError")
	delete(f, "TargetPath")
	delete(f, "InvokeLimitConfigStatus")
	delete(f, "InvokeLimitConfig")
	delete(f, "IpWhiteStatus")
	delete(f, "IpWhiteConfig")
	delete(f, "IpBlackStatus")
	delete(f, "IpBlackConfig")
	delete(f, "CustomHttpHost")
	delete(f, "HttpHostType")
	delete(f, "Timeout")
	delete(f, "McpSecurityRules")
	delete(f, "ToolConfigs")
	delete(f, "WrapPaasID")
	delete(f, "PluginConfigs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMcpServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMcpServerResponseParams struct {
	// mcp server ID
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMcpServerResponse struct {
	*tchttp.BaseResponse
	Response *CreateMcpServerResponseParams `json:"Response"`
}

func (r *CreateMcpServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMcpServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelRequestParams struct {
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 协议类型：http/https
	HttpProtocolType *string `json:"HttpProtocolType,omitnil,omitempty" name:"HttpProtocolType"`

	// 目标路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// 目标服务器
	TargetHosts []*TargetHostDTO `json:"TargetHosts,omitnil,omitempty" name:"TargetHosts"`

	// 凭据ID
	CredentialID *string `json:"CredentialID,omitnil,omitempty" name:"CredentialID"`

	// https时，是否检查证书合法
	CheckTargetCertsError *bool `json:"CheckTargetCertsError,omitnil,omitempty" name:"CheckTargetCertsError"`

	// http协议版本：1.1/2.0
	HttpProtocolVersion *string `json:"HttpProtocolVersion,omitnil,omitempty" name:"HttpProtocolVersion"`
}

type CreateModelRequest struct {
	*tchttp.BaseRequest
	
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 协议类型：http/https
	HttpProtocolType *string `json:"HttpProtocolType,omitnil,omitempty" name:"HttpProtocolType"`

	// 目标路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// 目标服务器
	TargetHosts []*TargetHostDTO `json:"TargetHosts,omitnil,omitempty" name:"TargetHosts"`

	// 凭据ID
	CredentialID *string `json:"CredentialID,omitnil,omitempty" name:"CredentialID"`

	// https时，是否检查证书合法
	CheckTargetCertsError *bool `json:"CheckTargetCertsError,omitnil,omitempty" name:"CheckTargetCertsError"`

	// http协议版本：1.1/2.0
	HttpProtocolVersion *string `json:"HttpProtocolVersion,omitnil,omitempty" name:"HttpProtocolVersion"`
}

func (r *CreateModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "Name")
	delete(f, "HttpProtocolType")
	delete(f, "TargetPath")
	delete(f, "TargetHosts")
	delete(f, "CredentialID")
	delete(f, "CheckTargetCertsError")
	delete(f, "HttpProtocolVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelResponseParams struct {
	// 结果集
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateModelResponse struct {
	*tchttp.BaseResponse
	Response *CreateModelResponseParams `json:"Response"`
}

func (r *CreateModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelServiceRequestParams struct {
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 访问路径
	PubPath *string `json:"PubPath,omitnil,omitempty" name:"PubPath"`

	// 模型ID列表
	TargetModels []*TargetModelDTO `json:"TargetModels,omitnil,omitempty" name:"TargetModels"`

	// 路径匹配类型: prefix 前缀匹配(不送默认); absolute 绝对匹配; regex正则匹配;
	PathMatchType *string `json:"PathMatchType,omitnil,omitempty" name:"PathMatchType"`

	// 是否开启限流
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 限流配置
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// 是否开启token控制
	TokenLimitStatus *bool `json:"TokenLimitStatus,omitnil,omitempty" name:"TokenLimitStatus"`

	// token控制
	TokenLimitConfig *TokenLimitConfigDTO `json:"TokenLimitConfig,omitnil,omitempty" name:"TokenLimitConfig"`

	// 是否开启内容安全
	TmsStatus *bool `json:"TmsStatus,omitnil,omitempty" name:"TmsStatus"`

	// 内容安全配置
	TmsConfig *TmsConfigDTO `json:"TmsConfig,omitnil,omitempty" name:"TmsConfig"`

	// 是否开启IP白名单
	IpWhiteStatus *bool `json:"IpWhiteStatus,omitnil,omitempty" name:"IpWhiteStatus"`

	// IP白名单
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// IP黑名单
	IpBlackList []*string `json:"IpBlackList,omitnil,omitempty" name:"IpBlackList"`

	// 插件配置
	PluginConfigs []*PluginConfigDTO `json:"PluginConfigs,omitnil,omitempty" name:"PluginConfigs"`

	// 超时配置，秒
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 是否开启提示词安全检测
	PromptModerateStatus *bool `json:"PromptModerateStatus,omitnil,omitempty" name:"PromptModerateStatus"`

	// 提示词安全检测配置
	PromptModerateConfig *PromptModerateConfigDTO `json:"PromptModerateConfig,omitnil,omitempty" name:"PromptModerateConfig"`
}

type CreateModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 访问路径
	PubPath *string `json:"PubPath,omitnil,omitempty" name:"PubPath"`

	// 模型ID列表
	TargetModels []*TargetModelDTO `json:"TargetModels,omitnil,omitempty" name:"TargetModels"`

	// 路径匹配类型: prefix 前缀匹配(不送默认); absolute 绝对匹配; regex正则匹配;
	PathMatchType *string `json:"PathMatchType,omitnil,omitempty" name:"PathMatchType"`

	// 是否开启限流
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 限流配置
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// 是否开启token控制
	TokenLimitStatus *bool `json:"TokenLimitStatus,omitnil,omitempty" name:"TokenLimitStatus"`

	// token控制
	TokenLimitConfig *TokenLimitConfigDTO `json:"TokenLimitConfig,omitnil,omitempty" name:"TokenLimitConfig"`

	// 是否开启内容安全
	TmsStatus *bool `json:"TmsStatus,omitnil,omitempty" name:"TmsStatus"`

	// 内容安全配置
	TmsConfig *TmsConfigDTO `json:"TmsConfig,omitnil,omitempty" name:"TmsConfig"`

	// 是否开启IP白名单
	IpWhiteStatus *bool `json:"IpWhiteStatus,omitnil,omitempty" name:"IpWhiteStatus"`

	// IP白名单
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// IP黑名单
	IpBlackList []*string `json:"IpBlackList,omitnil,omitempty" name:"IpBlackList"`

	// 插件配置
	PluginConfigs []*PluginConfigDTO `json:"PluginConfigs,omitnil,omitempty" name:"PluginConfigs"`

	// 超时配置，秒
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 是否开启提示词安全检测
	PromptModerateStatus *bool `json:"PromptModerateStatus,omitnil,omitempty" name:"PromptModerateStatus"`

	// 提示词安全检测配置
	PromptModerateConfig *PromptModerateConfigDTO `json:"PromptModerateConfig,omitnil,omitempty" name:"PromptModerateConfig"`
}

func (r *CreateModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "PubPath")
	delete(f, "TargetModels")
	delete(f, "PathMatchType")
	delete(f, "InvokeLimitConfigStatus")
	delete(f, "InvokeLimitConfig")
	delete(f, "TokenLimitStatus")
	delete(f, "TokenLimitConfig")
	delete(f, "TmsStatus")
	delete(f, "TmsConfig")
	delete(f, "IpWhiteStatus")
	delete(f, "IpWhiteList")
	delete(f, "IpBlackList")
	delete(f, "PluginConfigs")
	delete(f, "Timeout")
	delete(f, "PromptModerateStatus")
	delete(f, "PromptModerateConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelServiceResponseParams struct {
	// 结果集
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateModelServiceResponseParams `json:"Response"`
}

func (r *CreateModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentAppMcpServersRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// mcp server id数组
	McpServerIDs []*string `json:"McpServerIDs,omitnil,omitempty" name:"McpServerIDs"`
}

type DeleteAgentAppMcpServersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// mcp server id数组
	McpServerIDs []*string `json:"McpServerIDs,omitnil,omitempty" name:"McpServerIDs"`
}

func (r *DeleteAgentAppMcpServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentAppMcpServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	delete(f, "McpServerIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAgentAppMcpServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentAppMcpServersResponseParams struct {
	// app id
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAgentAppMcpServersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAgentAppMcpServersResponseParams `json:"Response"`
}

func (r *DeleteAgentAppMcpServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentAppMcpServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentAppModelServicesRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 关联的model service id
	ModelServiceIDs []*string `json:"ModelServiceIDs,omitnil,omitempty" name:"ModelServiceIDs"`
}

type DeleteAgentAppModelServicesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 关联的model service id
	ModelServiceIDs []*string `json:"ModelServiceIDs,omitnil,omitempty" name:"ModelServiceIDs"`
}

func (r *DeleteAgentAppModelServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentAppModelServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	delete(f, "ModelServiceIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAgentAppModelServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentAppModelServicesResponseParams struct {
	// app id
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAgentAppModelServicesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAgentAppModelServicesResponseParams `json:"Response"`
}

func (r *DeleteAgentAppModelServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentAppModelServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentAppRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteAgentAppRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteAgentAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAgentAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentAppResponseParams struct {
	// app id
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAgentAppResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAgentAppResponseParams `json:"Response"`
}

func (r *DeleteAgentAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentCredentialRequestParams struct {

}

type DeleteAgentCredentialRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteAgentCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAgentCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentCredentialResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAgentCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAgentCredentialResponseParams `json:"Response"`
}

func (r *DeleteAgentCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMcpServerRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// mcp server ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteMcpServerRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// mcp server ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteMcpServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMcpServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMcpServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMcpServerResponseParams struct {
	// mcp server ID
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMcpServerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMcpServerResponseParams `json:"Response"`
}

func (r *DeleteMcpServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMcpServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelRequestParams struct {
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型ID数组
	IDs []*string `json:"IDs,omitnil,omitempty" name:"IDs"`
}

type DeleteModelRequest struct {
	*tchttp.BaseRequest
	
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型ID数组
	IDs []*string `json:"IDs,omitnil,omitempty" name:"IDs"`
}

func (r *DeleteModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "IDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelResponseParams struct {
	// 结果集
	Data *ResultIDsVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteModelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModelResponseParams `json:"Response"`
}

func (r *DeleteModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelServiceRequestParams struct {
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelServiceResponseParams struct {
	// 结果集
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModelServiceResponseParams `json:"Response"`
}

func (r *DeleteModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentAppMcpServersRequestParams struct {
	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	AgentAppIDs []*string `json:"AgentAppIDs,omitnil,omitempty" name:"AgentAppIDs"`

	// 关联的mcp
	McpServerIDs []*string `json:"McpServerIDs,omitnil,omitempty" name:"McpServerIDs"`

	// 凭据ID
	AgentCredentialIDs []*string `json:"AgentCredentialIDs,omitnil,omitempty" name:"AgentCredentialIDs"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeAgentAppMcpServersRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	AgentAppIDs []*string `json:"AgentAppIDs,omitnil,omitempty" name:"AgentAppIDs"`

	// 关联的mcp
	McpServerIDs []*string `json:"McpServerIDs,omitnil,omitempty" name:"McpServerIDs"`

	// 凭据ID
	AgentCredentialIDs []*string `json:"AgentCredentialIDs,omitnil,omitempty" name:"AgentCredentialIDs"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeAgentAppMcpServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentAppMcpServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceID")
	delete(f, "AgentAppIDs")
	delete(f, "McpServerIDs")
	delete(f, "AgentCredentialIDs")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentAppMcpServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentAppMcpServersResp struct {
	// 条目总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 具体条目
	Items []*AgentAppMcpServerVO `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeAgentAppMcpServersResponseParams struct {
	// 列表
	Data *DescribeAgentAppMcpServersResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentAppMcpServersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentAppMcpServersResponseParams `json:"Response"`
}

func (r *DescribeAgentAppMcpServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentAppMcpServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentAppModelServicesRequestParams struct {

}

type DescribeAgentAppModelServicesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAgentAppModelServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentAppModelServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentAppModelServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentAppModelServicesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentAppModelServicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentAppModelServicesResponseParams `json:"Response"`
}

func (r *DescribeAgentAppModelServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentAppModelServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentAppRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// app id
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeAgentAppRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// app id
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeAgentAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentAppResp struct {
	// 租户appID
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 租户ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 认证类型
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// apiKeys列表，脱敏
	ApiKeys []*string `json:"ApiKeys,omitnil,omitempty" name:"ApiKeys"`

	// secretKey列表，脱敏
	SecretKeys []*AgentAppSecretKeyVO `json:"SecretKeys,omitnil,omitempty" name:"SecretKeys"`

	// OAuth2 Resource Server ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OAuth2ResourceServerID *string `json:"OAuth2ResourceServerID,omitnil,omitempty" name:"OAuth2ResourceServerID"`

	// 绑定mcpServer数量
	McpServersNum *int64 `json:"McpServersNum,omitnil,omitempty" name:"McpServersNum"`

	// 绑定的模型服务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelServicesNum *int64 `json:"ModelServicesNum,omitnil,omitempty" name:"ModelServicesNum"`
}

// Predefined struct for user
type DescribeAgentAppResponseParams struct {
	// app详情
	Data *DescribeAgentAppResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentAppResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentAppResponseParams `json:"Response"`
}

func (r *DescribeAgentAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentAppsRequestParams struct {
	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 服务ID数组
	IDs []*string `json:"IDs,omitnil,omitempty" name:"IDs"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// notID列表
	NotIDs []*string `json:"NotIDs,omitnil,omitempty" name:"NotIDs"`

	// 状态：normal正常状态，disabled下线状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 认证类型：apiKey，secretKey
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 排序
	Sort *DescribeAgentAppsSortDTO `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 凭据ID
	AgentCredentialID *string `json:"AgentCredentialID,omitnil,omitempty" name:"AgentCredentialID"`
}

type DescribeAgentAppsRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 服务ID数组
	IDs []*string `json:"IDs,omitnil,omitempty" name:"IDs"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// notID列表
	NotIDs []*string `json:"NotIDs,omitnil,omitempty" name:"NotIDs"`

	// 状态：normal正常状态，disabled下线状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 认证类型：apiKey，secretKey
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 排序
	Sort *DescribeAgentAppsSortDTO `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 凭据ID
	AgentCredentialID *string `json:"AgentCredentialID,omitnil,omitempty" name:"AgentCredentialID"`
}

func (r *DescribeAgentAppsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentAppsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "IDs")
	delete(f, "InstanceID")
	delete(f, "NotIDs")
	delete(f, "Status")
	delete(f, "Keyword")
	delete(f, "AuthType")
	delete(f, "Sort")
	delete(f, "AgentCredentialID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentAppsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentAppsResp struct {
	// 条目总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 具体条目
	Items []*DescribeAgentAppResp `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeAgentAppsResponseParams struct {
	// app列表
	Data *DescribeAgentAppsResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentAppsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentAppsResponseParams `json:"Response"`
}

func (r *DescribeAgentAppsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentAppsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentAppsSortDTO struct {
	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 名称
	Name *int64 `json:"Name,omitnil,omitempty" name:"Name"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DescribeAgentCredentialRequestParams struct {

}

type DescribeAgentCredentialRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAgentCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentCredentialResp struct {
	// 租户应用ID
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 租户ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 凭据ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 凭据名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 关联应用数
	RelateAgentAppNum *int64 `json:"RelateAgentAppNum,omitnil,omitempty" name:"RelateAgentAppNum"`

	// 关联mcp数
	RelateMcpServerNum *int64 `json:"RelateMcpServerNum,omitnil,omitempty" name:"RelateMcpServerNum"`

	// 关联模型数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelateModelNum *int64 `json:"RelateModelNum,omitnil,omitempty" name:"RelateModelNum"`

	// 凭据内容
	Content *AgentCredentialContentDTO `json:"Content,omitnil,omitempty" name:"Content"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type DescribeAgentCredentialResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentCredentialResponseParams `json:"Response"`
}

func (r *DescribeAgentCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentCredentialsRequestParams struct {

}

type DescribeAgentCredentialsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAgentCredentialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentCredentialsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentCredentialsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentCredentialsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentCredentialsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentCredentialsResponseParams `json:"Response"`
}

func (r *DescribeAgentCredentialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMcpServerRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// mcp server ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeMcpServerRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// mcp server ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeMcpServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMcpServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMcpServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMcpServerResponseParams struct {
	// mcp server详情
	Data *DescribeMcpServerResponseVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMcpServerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMcpServerResponseParams `json:"Response"`
}

func (r *DescribeMcpServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMcpServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMcpServerResponseVO struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 标签ID数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelIDs []*string `json:"LabelIDs,omitnil,omitempty" name:"LabelIDs"`

	// 目录ID数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryIDs []*string `json:"CategoryIDs,omitnil,omitempty" name:"CategoryIDs"`

	// 负载方式，robin random consistentHash
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetSelect *string `json:"TargetSelect,omitnil,omitempty" name:"TargetSelect"`

	// 目标服务器
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetHosts []*TargetHostDTO `json:"TargetHosts,omitnil,omitempty" name:"TargetHosts"`

	// 后端协议：http https
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpProtocolType *string `json:"HttpProtocolType,omitnil,omitempty" name:"HttpProtocolType"`

	// 证书检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckTargetCertsError *bool `json:"CheckTargetCertsError,omitnil,omitempty" name:"CheckTargetCertsError"`

	// 目标路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// 流量控制状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 流量控制配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// IP白名单开启状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpWhiteStatus *bool `json:"IpWhiteStatus,omitnil,omitempty" name:"IpWhiteStatus"`

	// IP白名单配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpWhiteConfig *IpConfig `json:"IpWhiteConfig,omitnil,omitempty" name:"IpWhiteConfig"`

	// IP黑名单开启状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpBlackStatus *bool `json:"IpBlackStatus,omitnil,omitempty" name:"IpBlackStatus"`

	// IP黑名单配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpBlackConfig *IpConfig `json:"IpBlackConfig,omitnil,omitempty" name:"IpBlackConfig"`

	// mcp server ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 预览地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 应用
	// 注意：此字段可能返回 null，表示取不到有效值。
	App *IDNameVO `json:"App,omitnil,omitempty" name:"App"`

	// 目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalogs []*IDNameVO `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*IDNameVO `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 用户appID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 自定义host
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomHttpHost *string `json:"CustomHttpHost,omitnil,omitempty" name:"CustomHttpHost"`

	//  Http 请求host类型 useRequestHost 保持源请求host targetHost 修正为源站host  customHost 自定义host
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpHostType *string `json:"HttpHostType,omitnil,omitempty" name:"HttpHostType"`

	// 请求的超时时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// mcp server模式
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// mcp version
	McpVersion *string `json:"McpVersion,omitnil,omitempty" name:"McpVersion"`

	// 封装模式下绑定的服务ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	WrapServices []*string `json:"WrapServices,omitnil,omitempty" name:"WrapServices"`

	// 工具数量
	ToolNum *int64 `json:"ToolNum,omitnil,omitempty" name:"ToolNum"`

	// 安全规则集响应
	McpSecurityRulesVO []*McpSecurityRulesVO `json:"McpSecurityRulesVO,omitnil,omitempty" name:"McpSecurityRulesVO"`

	// 真实工具级别配置，实时拉取了tool/list做渲染的，如果tool/list不通，就拉不到。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolConfigs []*ToolConfigVO `json:"ToolConfigs,omitnil,omitempty" name:"ToolConfigs"`

	// 访问URL
	UrlObj *McpUrlObj `json:"UrlObj,omitnil,omitempty" name:"UrlObj"`

	// 后端mcp服务是否正常
	ToolMessage *string `json:"ToolMessage,omitnil,omitempty" name:"ToolMessage"`

	// 后端mcp服务的工具列表
	Tools []*McpTool `json:"Tools,omitnil,omitempty" name:"Tools"`

	// 封装的API分组ID
	WrapPaasID *string `json:"WrapPaasID,omitnil,omitempty" name:"WrapPaasID"`

	// 关联的agentApp数量
	RelateAgentAppNum *int64 `json:"RelateAgentAppNum,omitnil,omitempty" name:"RelateAgentAppNum"`

	// 插件配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PluginConfigs []*PluginConfigDTO `json:"PluginConfigs,omitnil,omitempty" name:"PluginConfigs"`
}

// Predefined struct for user
type DescribeMcpServersRequestParams struct {
	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 状态数组：normal正常状态，disabled下线状态
	Statuses []*string `json:"Statuses,omitnil,omitempty" name:"Statuses"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 服务ID数组
	IDs []*string `json:"IDs,omitnil,omitempty" name:"IDs"`

	// 模式：proxy代理模式； wrap封装模式；
	Modes []*string `json:"Modes,omitnil,omitempty" name:"Modes"`

	// 绑定的安全规则ID
	McpSecurityRuleID *string `json:"McpSecurityRuleID,omitnil,omitempty" name:"McpSecurityRuleID"`

	// 绑定安全规则的处置动作（填写时McpSecurityRuleID得必填）
	McpSecurityRuleAct *string `json:"McpSecurityRuleAct,omitnil,omitempty" name:"McpSecurityRuleAct"`

	// 已绑定插件id
	PluginID *string `json:"PluginID,omitnil,omitempty" name:"PluginID"`
}

type DescribeMcpServersRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 状态数组：normal正常状态，disabled下线状态
	Statuses []*string `json:"Statuses,omitnil,omitempty" name:"Statuses"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 服务ID数组
	IDs []*string `json:"IDs,omitnil,omitempty" name:"IDs"`

	// 模式：proxy代理模式； wrap封装模式；
	Modes []*string `json:"Modes,omitnil,omitempty" name:"Modes"`

	// 绑定的安全规则ID
	McpSecurityRuleID *string `json:"McpSecurityRuleID,omitnil,omitempty" name:"McpSecurityRuleID"`

	// 绑定安全规则的处置动作（填写时McpSecurityRuleID得必填）
	McpSecurityRuleAct *string `json:"McpSecurityRuleAct,omitnil,omitempty" name:"McpSecurityRuleAct"`

	// 已绑定插件id
	PluginID *string `json:"PluginID,omitnil,omitempty" name:"PluginID"`
}

func (r *DescribeMcpServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMcpServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceID")
	delete(f, "Statuses")
	delete(f, "Keyword")
	delete(f, "IDs")
	delete(f, "Modes")
	delete(f, "McpSecurityRuleID")
	delete(f, "McpSecurityRuleAct")
	delete(f, "PluginID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMcpServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMcpServersResponseParams struct {
	// mcp server列表
	Data *DescribeMcpServersResponseVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMcpServersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMcpServersResponseParams `json:"Response"`
}

func (r *DescribeMcpServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMcpServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMcpServersResponseVO struct {
	// 条目
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DescribeMcpServerResponseVO `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DescribeModelRequestParams struct {
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeModelRequest struct {
	*tchttp.BaseRequest
	
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelResponseParams struct {
	// 结果集
	Data *DescribeModelResponseVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelResponseParams `json:"Response"`
}

func (r *DescribeModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModelResponseVO struct {
	// 腾讯云AppID
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 腾讯云Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 模型名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 凭据ID
	CredentialID *string `json:"CredentialID,omitnil,omitempty" name:"CredentialID"`

	// 凭据名称
	CredentialName *string `json:"CredentialName,omitnil,omitempty" name:"CredentialName"`

	// http协议类型
	HttpProtocolType *string `json:"HttpProtocolType,omitnil,omitempty" name:"HttpProtocolType"`

	// https时，是否校验目标证书
	CheckTargetCertsError *bool `json:"CheckTargetCertsError,omitnil,omitempty" name:"CheckTargetCertsError"`

	// http协议版本：1.1/2.0
	HttpProtocolVersion *string `json:"HttpProtocolVersion,omitnil,omitempty" name:"HttpProtocolVersion"`

	// 目标路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// 目标器列表
	TargetHosts []*TargetHostDTO `json:"TargetHosts,omitnil,omitempty" name:"TargetHosts"`

	// 被模型服务使用的个数
	ModelServiceCount *int64 `json:"ModelServiceCount,omitnil,omitempty" name:"ModelServiceCount"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后修改时间
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`
}

// Predefined struct for user
type DescribeModelServiceRequestParams struct {
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceResponseParams struct {
	// 结果集
	Data *DescribeModelServiceResponseVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceResponseParams `json:"Response"`
}

func (r *DescribeModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModelServiceResponseVO struct {
	// 腾讯云AppID
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 腾讯云Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 模型名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 访问路径
	PubPath *string `json:"PubPath,omitnil,omitempty" name:"PubPath"`

	// 路径匹配方式：absolute，prefix，regex
	PathMatchType *string `json:"PathMatchType,omitnil,omitempty" name:"PathMatchType"`

	// 目标模型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetModels []*TargetModelDTO `json:"TargetModels,omitnil,omitempty" name:"TargetModels"`

	// 模板模型的名称列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelNames []*string `json:"ModelNames,omitnil,omitempty" name:"ModelNames"`

	// 是否开启限流
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 限流配置
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后修改时间
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 是否开启token控制
	TokenLimitStatus *bool `json:"TokenLimitStatus,omitnil,omitempty" name:"TokenLimitStatus"`

	// token控制
	TokenLimitConfig *TokenLimitConfigDTO `json:"TokenLimitConfig,omitnil,omitempty" name:"TokenLimitConfig"`

	// 是否开启tms配置
	TmsStatus *bool `json:"TmsStatus,omitnil,omitempty" name:"TmsStatus"`

	// tms配置
	TmsConfig *TmsConfigDTO `json:"TmsConfig,omitnil,omitempty" name:"TmsConfig"`

	// 是否开启IP白名单
	IpWhiteStatus *bool `json:"IpWhiteStatus,omitnil,omitempty" name:"IpWhiteStatus"`

	// IP白名单列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// 是否开启IP黑名单
	IpBlackStatus *bool `json:"IpBlackStatus,omitnil,omitempty" name:"IpBlackStatus"`

	// IP黑名单列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpBlackList []*string `json:"IpBlackList,omitnil,omitempty" name:"IpBlackList"`

	// 插件配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PluginConfigs []*PluginConfigDTO `json:"PluginConfigs,omitnil,omitempty" name:"PluginConfigs"`

	// 超时配置，单位秒
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 状态：normal，disabled
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 关联应用数
	RelateAgentAppNum *int64 `json:"RelateAgentAppNum,omitnil,omitempty" name:"RelateAgentAppNum"`

	// 请求路径
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 是否开启提示词安全检测
	// 注意：此字段可能返回 null，表示取不到有效值。
	PromptModerateStatus *bool `json:"PromptModerateStatus,omitnil,omitempty" name:"PromptModerateStatus"`

	// 提示词安全检测配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PromptModerateConfig *PromptModerateConfigDTO `json:"PromptModerateConfig,omitnil,omitempty" name:"PromptModerateConfig"`
}

// Predefined struct for user
type DescribeModelServicesRequestParams struct {
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 分页参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// ID列表
	IDs []*string `json:"IDs,omitnil,omitempty" name:"IDs"`

	// 排除的ID列表
	NotIDs []*string `json:"NotIDs,omitnil,omitempty" name:"NotIDs"`

	// 状态：normal，disabled
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 模型ID
	ModelID *string `json:"ModelID,omitnil,omitempty" name:"ModelID"`

	// 排序
	Sort *DescribeModelServicesSort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

type DescribeModelServicesRequest struct {
	*tchttp.BaseRequest
	
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 分页参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// ID列表
	IDs []*string `json:"IDs,omitnil,omitempty" name:"IDs"`

	// 排除的ID列表
	NotIDs []*string `json:"NotIDs,omitnil,omitempty" name:"NotIDs"`

	// 状态：normal，disabled
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 模型ID
	ModelID *string `json:"ModelID,omitnil,omitempty" name:"ModelID"`

	// 排序
	Sort *DescribeModelServicesSort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

func (r *DescribeModelServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IDs")
	delete(f, "NotIDs")
	delete(f, "Status")
	delete(f, "Keyword")
	delete(f, "ModelID")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServicesResponseParams struct {
	// 结果集
	Data *DescribeModelServicesResponseVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServicesResponseParams `json:"Response"`
}

func (r *DescribeModelServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModelServicesResponseVO struct {
	// 条目
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 列表
	Items []*DescribeModelServiceResponseVO `json:"Items,omitnil,omitempty" name:"Items"`
}

type DescribeModelServicesSort struct {
	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后修改时间
	LastUpdateTime *int64 `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 模型名称
	Name *int64 `json:"Name,omitnil,omitempty" name:"Name"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DescribeModelsRequestParams struct {
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 分页参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// ID列表
	IDs []*string `json:"IDs,omitnil,omitempty" name:"IDs"`

	// 排除的ID列表
	NotIDs []*string `json:"NotIDs,omitnil,omitempty" name:"NotIDs"`

	// 凭据ID
	CredentialID *string `json:"CredentialID,omitnil,omitempty" name:"CredentialID"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 排序
	Sort *DescribeModelsSort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

type DescribeModelsRequest struct {
	*tchttp.BaseRequest
	
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 分页参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// ID列表
	IDs []*string `json:"IDs,omitnil,omitempty" name:"IDs"`

	// 排除的ID列表
	NotIDs []*string `json:"NotIDs,omitnil,omitempty" name:"NotIDs"`

	// 凭据ID
	CredentialID *string `json:"CredentialID,omitnil,omitempty" name:"CredentialID"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 排序
	Sort *DescribeModelsSort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

func (r *DescribeModelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IDs")
	delete(f, "NotIDs")
	delete(f, "CredentialID")
	delete(f, "Keyword")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelsResponseParams struct {
	// 结果集
	Data *DescribeModelsResponseVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelsResponseParams `json:"Response"`
}

func (r *DescribeModelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModelsResponseVO struct {
	// 条目
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 列表
	Items []*DescribeModelResponseVO `json:"Items,omitnil,omitempty" name:"Items"`
}

type DescribeModelsSort struct {
	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后修改时间
	LastUpdateTime *int64 `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 模型名称
	Name *int64 `json:"Name,omitnil,omitempty" name:"Name"`
}

type IDNameVO struct {
	// 业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type InvokeLimitConfigDTO struct {
	// Type类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 令牌桶最大容量
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenBucketMaxNum *int64 `json:"TokenBucketMaxNum,omitnil,omitempty" name:"TokenBucketMaxNum"`

	// 令牌生成速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenBucketRate *int64 `json:"TokenBucketRate,omitnil,omitempty" name:"TokenBucketRate"`

	// 漏斗容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunnelMaxNum *int64 `json:"FunnelMaxNum,omitnil,omitempty" name:"FunnelMaxNum"`

	// 漏嘴流速
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunnelRate *int64 `json:"FunnelRate,omitnil,omitempty" name:"FunnelRate"`

	// 滑动窗口最大请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlidingWindowMaxNum *int64 `json:"SlidingWindowMaxNum,omitnil,omitempty" name:"SlidingWindowMaxNum"`

	// 滑动窗口长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlidingWindowSize *int64 `json:"SlidingWindowSize,omitnil,omitempty" name:"SlidingWindowSize"`

	// 时间窗口最大请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeWindow *int64 `json:"TimeWindow,omitnil,omitempty" name:"TimeWindow"`

	// 时间窗口长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeWindowInterval *int64 `json:"TimeWindowInterval,omitnil,omitempty" name:"TimeWindowInterval"`

	// 请求的超时时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`
}

type IpConfig struct {
	// ip数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ips []*string `json:"Ips,omitnil,omitempty" name:"Ips"`

	// 生效类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EffectType *string `json:"EffectType,omitnil,omitempty" name:"EffectType"`

	// 生效时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EffectTimes []*StartEndTime `json:"EffectTimes,omitnil,omitempty" name:"EffectTimes"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type LimitWindowsDTO struct {
	// 时间窗口，分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 累计上限，k
	// 注意：此字段可能返回 null，表示取不到有效值。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type McpInputOutSchema struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 必填字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Required []*string `json:"Required,omitnil,omitempty" name:"Required"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type McpSecurityRule struct {
	// 规则ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 处置动作
	Act *string `json:"Act,omitnil,omitempty" name:"Act"`
}

type McpSecurityRulesVO struct {
	// 规则ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 规则类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 风险等级
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 版本号
	VersionNumber *string `json:"VersionNumber,omitnil,omitempty" name:"VersionNumber"`

	// 状态开关
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 当前选择的处置动作
	Act *string `json:"Act,omitnil,omitempty" name:"Act"`

	// 支持的处置动作
	SupportActs []*string `json:"SupportActs,omitnil,omitempty" name:"SupportActs"`

	// 内容类型
	BodyType *string `json:"BodyType,omitnil,omitempty" name:"BodyType"`

	// icon类型
	IconType *string `json:"IconType,omitnil,omitempty" name:"IconType"`
}

type McpTool struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 入参参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputSchema *McpInputOutSchema `json:"InputSchema,omitnil,omitempty" name:"InputSchema"`

	// 注释
	// 注意：此字段可能返回 null，表示取不到有效值。
	Annotations *McpToolAnnotation `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// 出参参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputSchema *McpInputOutSchema `json:"OutputSchema,omitnil,omitempty" name:"OutputSchema"`
}

type McpToolAnnotation struct {
	// title for the tool
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// If true, the tool does not modify its environment
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnlyHint *bool `json:"ReadOnlyHint,omitnil,omitempty" name:"ReadOnlyHint"`

	// If true, the tool may perform destructive updates
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestructiveHint *bool `json:"DestructiveHint,omitnil,omitempty" name:"DestructiveHint"`

	// If true, repeated calls with same args have no additional effect
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdempotentHint *bool `json:"IdempotentHint,omitnil,omitempty" name:"IdempotentHint"`

	// If true, tool interacts with external entities
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenWorldHint *bool `json:"OpenWorldHint,omitnil,omitempty" name:"OpenWorldHint"`
}

type McpUrlObj struct {
	// sse访问URL
	SSEUrl *string `json:"SSEUrl,omitnil,omitempty" name:"SSEUrl"`

	// streamable访问URL
	StreamAbleUrl *string `json:"StreamAbleUrl,omitnil,omitempty" name:"StreamAbleUrl"`
}

// Predefined struct for user
type ModifyAgentAppModelServicesRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 关联的model service
	ModelServices []*AgentAppModelServiceDTO `json:"ModelServices,omitnil,omitempty" name:"ModelServices"`
}

type ModifyAgentAppModelServicesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 关联的model service
	ModelServices []*AgentAppModelServiceDTO `json:"ModelServices,omitnil,omitempty" name:"ModelServices"`
}

func (r *ModifyAgentAppModelServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentAppModelServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	delete(f, "ModelServices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAgentAppModelServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentAppModelServicesResponseParams struct {
	// app id
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAgentAppModelServicesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAgentAppModelServicesResponseParams `json:"Response"`
}

func (r *ModifyAgentAppModelServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentAppModelServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentAppRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// OAuth2资源服务器ID
	OAuth2ResourceServerID *string `json:"OAuth2ResourceServerID,omitnil,omitempty" name:"OAuth2ResourceServerID"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyAgentAppRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// OAuth2资源服务器ID
	OAuth2ResourceServerID *string `json:"OAuth2ResourceServerID,omitnil,omitempty" name:"OAuth2ResourceServerID"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyAgentAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	delete(f, "Name")
	delete(f, "OAuth2ResourceServerID")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAgentAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentAppResponseParams struct {
	// app id
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAgentAppResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAgentAppResponseParams `json:"Response"`
}

func (r *ModifyAgentAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentCredentialRequestParams struct {

}

type ModifyAgentCredentialRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyAgentCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAgentCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentCredentialResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAgentCredentialResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAgentCredentialResponseParams `json:"Response"`
}

func (r *ModifyAgentCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMcpServerRequestParams struct {
	// mcp server ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 模式：proxy代理模式； wrap封装模式；
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 版本号：2024-11-05 2025-03-26
	McpVersion *string `json:"McpVersion,omitnil,omitempty" name:"McpVersion"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 封装服务列表
	WrapServices []*string `json:"WrapServices,omitnil,omitempty" name:"WrapServices"`

	// 负载方式，robin random consistentHash
	TargetSelect *string `json:"TargetSelect,omitnil,omitempty" name:"TargetSelect"`

	// 目标服务器
	TargetHosts []*TargetHostDTO `json:"TargetHosts,omitnil,omitempty" name:"TargetHosts"`

	// 后端协议：http https
	HttpProtocolType *string `json:"HttpProtocolType,omitnil,omitempty" name:"HttpProtocolType"`

	// 证书检查
	CheckTargetCertsError *bool `json:"CheckTargetCertsError,omitnil,omitempty" name:"CheckTargetCertsError"`

	// 目标路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// 流量控制开启状态
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 流量控制配置
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// IP白名单开启状态
	IpWhiteStatus *bool `json:"IpWhiteStatus,omitnil,omitempty" name:"IpWhiteStatus"`

	// IP白名单配置
	IpWhiteConfig *IpConfig `json:"IpWhiteConfig,omitnil,omitempty" name:"IpWhiteConfig"`

	// IP黑名单开启状态
	IpBlackStatus *bool `json:"IpBlackStatus,omitnil,omitempty" name:"IpBlackStatus"`

	// IP黑名单配置
	IpBlackConfig *IpConfig `json:"IpBlackConfig,omitnil,omitempty" name:"IpBlackConfig"`

	// 目标Host类型 0 默认 1 vpc
	TargetHostType *int64 `json:"TargetHostType,omitnil,omitempty" name:"TargetHostType"`

	// 自定义host
	CustomHttpHost *string `json:"CustomHttpHost,omitnil,omitempty" name:"CustomHttpHost"`

	// Http 请求host类型：useRequestHost 保持源请求；host targetHost 修正为源站host； customHost 自定义host
	HttpHostType *string `json:"HttpHostType,omitnil,omitempty" name:"HttpHostType"`

	// 请求的超时时间
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 安全规则集
	McpSecurityRules []*McpSecurityRule `json:"McpSecurityRules,omitnil,omitempty" name:"McpSecurityRules"`

	// 工具集配置（openapi可能会用到）
	ToolConfigs []*ToolConfigDTO `json:"ToolConfigs,omitnil,omitempty" name:"ToolConfigs"`

	// 封装的API分组ID
	WrapPaasID *string `json:"WrapPaasID,omitnil,omitempty" name:"WrapPaasID"`

	// 插件配置
	PluginConfigs []*PluginConfigDTO `json:"PluginConfigs,omitnil,omitempty" name:"PluginConfigs"`
}

type ModifyMcpServerRequest struct {
	*tchttp.BaseRequest
	
	// mcp server ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 模式：proxy代理模式； wrap封装模式；
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 版本号：2024-11-05 2025-03-26
	McpVersion *string `json:"McpVersion,omitnil,omitempty" name:"McpVersion"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 封装服务列表
	WrapServices []*string `json:"WrapServices,omitnil,omitempty" name:"WrapServices"`

	// 负载方式，robin random consistentHash
	TargetSelect *string `json:"TargetSelect,omitnil,omitempty" name:"TargetSelect"`

	// 目标服务器
	TargetHosts []*TargetHostDTO `json:"TargetHosts,omitnil,omitempty" name:"TargetHosts"`

	// 后端协议：http https
	HttpProtocolType *string `json:"HttpProtocolType,omitnil,omitempty" name:"HttpProtocolType"`

	// 证书检查
	CheckTargetCertsError *bool `json:"CheckTargetCertsError,omitnil,omitempty" name:"CheckTargetCertsError"`

	// 目标路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// 流量控制开启状态
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 流量控制配置
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// IP白名单开启状态
	IpWhiteStatus *bool `json:"IpWhiteStatus,omitnil,omitempty" name:"IpWhiteStatus"`

	// IP白名单配置
	IpWhiteConfig *IpConfig `json:"IpWhiteConfig,omitnil,omitempty" name:"IpWhiteConfig"`

	// IP黑名单开启状态
	IpBlackStatus *bool `json:"IpBlackStatus,omitnil,omitempty" name:"IpBlackStatus"`

	// IP黑名单配置
	IpBlackConfig *IpConfig `json:"IpBlackConfig,omitnil,omitempty" name:"IpBlackConfig"`

	// 目标Host类型 0 默认 1 vpc
	TargetHostType *int64 `json:"TargetHostType,omitnil,omitempty" name:"TargetHostType"`

	// 自定义host
	CustomHttpHost *string `json:"CustomHttpHost,omitnil,omitempty" name:"CustomHttpHost"`

	// Http 请求host类型：useRequestHost 保持源请求；host targetHost 修正为源站host； customHost 自定义host
	HttpHostType *string `json:"HttpHostType,omitnil,omitempty" name:"HttpHostType"`

	// 请求的超时时间
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 安全规则集
	McpSecurityRules []*McpSecurityRule `json:"McpSecurityRules,omitnil,omitempty" name:"McpSecurityRules"`

	// 工具集配置（openapi可能会用到）
	ToolConfigs []*ToolConfigDTO `json:"ToolConfigs,omitnil,omitempty" name:"ToolConfigs"`

	// 封装的API分组ID
	WrapPaasID *string `json:"WrapPaasID,omitnil,omitempty" name:"WrapPaasID"`

	// 插件配置
	PluginConfigs []*PluginConfigDTO `json:"PluginConfigs,omitnil,omitempty" name:"PluginConfigs"`
}

func (r *ModifyMcpServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMcpServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "Mode")
	delete(f, "McpVersion")
	delete(f, "InstanceID")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "WrapServices")
	delete(f, "TargetSelect")
	delete(f, "TargetHosts")
	delete(f, "HttpProtocolType")
	delete(f, "CheckTargetCertsError")
	delete(f, "TargetPath")
	delete(f, "InvokeLimitConfigStatus")
	delete(f, "InvokeLimitConfig")
	delete(f, "IpWhiteStatus")
	delete(f, "IpWhiteConfig")
	delete(f, "IpBlackStatus")
	delete(f, "IpBlackConfig")
	delete(f, "TargetHostType")
	delete(f, "CustomHttpHost")
	delete(f, "HttpHostType")
	delete(f, "Timeout")
	delete(f, "McpSecurityRules")
	delete(f, "ToolConfigs")
	delete(f, "WrapPaasID")
	delete(f, "PluginConfigs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMcpServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMcpServerResponseParams struct {
	// mcp server ID
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMcpServerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMcpServerResponseParams `json:"Response"`
}

func (r *ModifyMcpServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMcpServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelRequestParams struct {
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 模型名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 协议类型：http/https
	HttpProtocolType *string `json:"HttpProtocolType,omitnil,omitempty" name:"HttpProtocolType"`

	// 目标路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// 目标服务器
	TargetHosts []*TargetHostDTO `json:"TargetHosts,omitnil,omitempty" name:"TargetHosts"`

	// 凭据ID
	CredentialID *string `json:"CredentialID,omitnil,omitempty" name:"CredentialID"`

	// https时，是否检查证书合法
	CheckTargetCertsError *bool `json:"CheckTargetCertsError,omitnil,omitempty" name:"CheckTargetCertsError"`

	// http协议版本：1.1/2.0
	HttpProtocolVersion *string `json:"HttpProtocolVersion,omitnil,omitempty" name:"HttpProtocolVersion"`
}

type ModifyModelRequest struct {
	*tchttp.BaseRequest
	
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 模型名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 协议类型：http/https
	HttpProtocolType *string `json:"HttpProtocolType,omitnil,omitempty" name:"HttpProtocolType"`

	// 目标路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// 目标服务器
	TargetHosts []*TargetHostDTO `json:"TargetHosts,omitnil,omitempty" name:"TargetHosts"`

	// 凭据ID
	CredentialID *string `json:"CredentialID,omitnil,omitempty" name:"CredentialID"`

	// https时，是否检查证书合法
	CheckTargetCertsError *bool `json:"CheckTargetCertsError,omitnil,omitempty" name:"CheckTargetCertsError"`

	// http协议版本：1.1/2.0
	HttpProtocolVersion *string `json:"HttpProtocolVersion,omitnil,omitempty" name:"HttpProtocolVersion"`
}

func (r *ModifyModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	delete(f, "Name")
	delete(f, "HttpProtocolType")
	delete(f, "TargetPath")
	delete(f, "TargetHosts")
	delete(f, "CredentialID")
	delete(f, "CheckTargetCertsError")
	delete(f, "HttpProtocolVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelResponseParams struct {
	// 结果集
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelResponseParams `json:"Response"`
}

func (r *ModifyModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelServiceRequestParams struct {
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 模型服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 模板模型列表
	TargetModels []*TargetModelDTO `json:"TargetModels,omitnil,omitempty" name:"TargetModels"`

	// 是否开启限流
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 限流配置
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// 是否开启token控制
	TokenLimitStatus *bool `json:"TokenLimitStatus,omitnil,omitempty" name:"TokenLimitStatus"`

	// token控制
	TokenLimitConfig *TokenLimitConfigDTO `json:"TokenLimitConfig,omitnil,omitempty" name:"TokenLimitConfig"`

	// 是否开启内容安全
	TmsStatus *bool `json:"TmsStatus,omitnil,omitempty" name:"TmsStatus"`

	// 内容安全配置
	TmsConfig *TmsConfigDTO `json:"TmsConfig,omitnil,omitempty" name:"TmsConfig"`

	// 是否开启IP白名单
	IpWhiteStatus *bool `json:"IpWhiteStatus,omitnil,omitempty" name:"IpWhiteStatus"`

	// IP白名单
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// 是否开启IP黑名单
	IpBlackStatus *bool `json:"IpBlackStatus,omitnil,omitempty" name:"IpBlackStatus"`

	// IP黑名单
	IpBlackList []*string `json:"IpBlackList,omitnil,omitempty" name:"IpBlackList"`

	// 插件配置
	PluginConfigs []*PluginConfigDTO `json:"PluginConfigs,omitnil,omitempty" name:"PluginConfigs"`

	// 超时配置，秒
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 是否开启提示词安全检测配置
	PromptModerateStatus *bool `json:"PromptModerateStatus,omitnil,omitempty" name:"PromptModerateStatus"`

	// 提示词安全检测配置
	PromptModerateConfig *PromptModerateConfigDTO `json:"PromptModerateConfig,omitnil,omitempty" name:"PromptModerateConfig"`
}

type ModifyModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 模型服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 模型服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 模板模型列表
	TargetModels []*TargetModelDTO `json:"TargetModels,omitnil,omitempty" name:"TargetModels"`

	// 是否开启限流
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 限流配置
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// 是否开启token控制
	TokenLimitStatus *bool `json:"TokenLimitStatus,omitnil,omitempty" name:"TokenLimitStatus"`

	// token控制
	TokenLimitConfig *TokenLimitConfigDTO `json:"TokenLimitConfig,omitnil,omitempty" name:"TokenLimitConfig"`

	// 是否开启内容安全
	TmsStatus *bool `json:"TmsStatus,omitnil,omitempty" name:"TmsStatus"`

	// 内容安全配置
	TmsConfig *TmsConfigDTO `json:"TmsConfig,omitnil,omitempty" name:"TmsConfig"`

	// 是否开启IP白名单
	IpWhiteStatus *bool `json:"IpWhiteStatus,omitnil,omitempty" name:"IpWhiteStatus"`

	// IP白名单
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// 是否开启IP黑名单
	IpBlackStatus *bool `json:"IpBlackStatus,omitnil,omitempty" name:"IpBlackStatus"`

	// IP黑名单
	IpBlackList []*string `json:"IpBlackList,omitnil,omitempty" name:"IpBlackList"`

	// 插件配置
	PluginConfigs []*PluginConfigDTO `json:"PluginConfigs,omitnil,omitempty" name:"PluginConfigs"`

	// 超时配置，秒
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 是否开启提示词安全检测配置
	PromptModerateStatus *bool `json:"PromptModerateStatus,omitnil,omitempty" name:"PromptModerateStatus"`

	// 提示词安全检测配置
	PromptModerateConfig *PromptModerateConfigDTO `json:"PromptModerateConfig,omitnil,omitempty" name:"PromptModerateConfig"`
}

func (r *ModifyModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "TargetModels")
	delete(f, "InvokeLimitConfigStatus")
	delete(f, "InvokeLimitConfig")
	delete(f, "TokenLimitStatus")
	delete(f, "TokenLimitConfig")
	delete(f, "TmsStatus")
	delete(f, "TmsConfig")
	delete(f, "IpWhiteStatus")
	delete(f, "IpWhiteList")
	delete(f, "IpBlackStatus")
	delete(f, "IpBlackList")
	delete(f, "PluginConfigs")
	delete(f, "Timeout")
	delete(f, "PromptModerateStatus")
	delete(f, "PromptModerateConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelServiceResponseParams struct {
	// 结果集
	Data *ResultIDVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelServiceResponseParams `json:"Response"`
}

func (r *ModifyModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PluginConfigDTO struct {
	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PluginName *string `json:"PluginName,omitnil,omitempty" name:"PluginName"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PluginID *string `json:"PluginID,omitnil,omitempty" name:"PluginID"`

	// 优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 表单配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PluginFormValues []*PluginFormValueDTO `json:"PluginFormValues,omitnil,omitempty" name:"PluginFormValues"`
}

type PluginFormValueDTO struct {
	// 字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type PromptModerateConfigDTO struct {
	// 执行动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 响应拦截内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	InterceptMessage *string `json:"InterceptMessage,omitnil,omitempty" name:"InterceptMessage"`
}

type ResultIDVO struct {
	// 对象ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

type ResultIDsVO struct {
	// 结果ID数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	IDs []*string `json:"IDs,omitnil,omitempty" name:"IDs"`
}

type StartEndTime struct {
	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type TargetHostDTO struct {
	// 服务器
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rank *int64 `json:"Rank,omitnil,omitempty" name:"Rank"`
}

type TargetModelDTO struct {
	// 模型ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 匹配名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 权重
	Rank *int64 `json:"Rank,omitnil,omitempty" name:"Rank"`
}

type TmsConfigDTO struct {
	// 检测范围,请求/响应
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 检测形式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 执行动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 合并请求检测event数，联动Mode字段sync
	// 注意：此字段可能返回 null，表示取不到有效值。
	MergeCount *uint64 `json:"MergeCount,omitnil,omitempty" name:"MergeCount"`

	// 风控策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 响应拦截内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	InterceptMessage *string `json:"InterceptMessage,omitnil,omitempty" name:"InterceptMessage"`
}

type TokenLimitConfigDTO struct {
	// 单次请求上限，k
	// 注意：此字段可能返回 null，表示取不到有效值。
	LimitRequestBody *uint64 `json:"LimitRequestBody,omitnil,omitempty" name:"LimitRequestBody"`

	// 累次token总量消耗上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	LimitWindows []*LimitWindowsDTO `json:"LimitWindows,omitnil,omitempty" name:"LimitWindows"`
}

type ToolConfigDTO struct {
	// 工具名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// 是否开启限流配置
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 限流配置
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// 绑定安全规则
	McpSecurityRules []*BindMcpSecurityRuleDTO `json:"McpSecurityRules,omitnil,omitempty" name:"McpSecurityRules"`
}

type ToolConfigVO struct {
	// 工具名称
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// 是否开启限流配置
	InvokeLimitConfigStatus *bool `json:"InvokeLimitConfigStatus,omitnil,omitempty" name:"InvokeLimitConfigStatus"`

	// 限流配置
	InvokeLimitConfig *InvokeLimitConfigDTO `json:"InvokeLimitConfig,omitnil,omitempty" name:"InvokeLimitConfig"`

	// 绑定安全规则（响应）
	McpSecurityRules []*BindMcpSecurityRuleVO `json:"McpSecurityRules,omitnil,omitempty" name:"McpSecurityRules"`
}