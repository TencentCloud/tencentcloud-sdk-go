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

package v20260709

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type A2AConfig struct {
	// Agent 级唯一 A2A 开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2AEnabled *bool `json:"A2AEnabled,omitnil,omitempty" name:"A2AEnabled"`

	// 对外 A2A handle（已注册时；仅 DescribeAgent / ModifyAgentA2AConfig 填充）
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2APublicRef *string `json:"A2APublicRef,omitnil,omitempty" name:"A2APublicRef"`

	// 对外 A2A card 发现地址（已注册时）
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2AEndpoint *string `json:"A2AEndpoint,omitnil,omitempty" name:"A2AEndpoint"`

	// 注册状态：DRAFT / REGISTERED / DISABLED / NONE / UNKNOWN
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2AStatus *string `json:"A2AStatus,omitnil,omitempty" name:"A2AStatus"`
}

type A2ASkillInput struct {
	// <p>A2A skill ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2ASkillId *string `json:"A2ASkillId,omitnil,omitempty" name:"A2ASkillId"`

	// <p>skill 名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>skill 描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>标签</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>示例</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Examples []*string `json:"Examples,omitnil,omitempty" name:"Examples"`
}

type A2ASkillItem struct {
	// A2A skill ID（加 A2A 前缀与内部 SkillId 概念区分）
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2ASkillId *string `json:"A2ASkillId,omitnil,omitempty" name:"A2ASkillId"`

	// skill 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// skill 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type AgentItem struct {
	// Agent 业务 ID（全局唯一，数字字符串形态）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Agent 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// Agent 描述；未填写时缺省
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 头像 URL；未设置时缺省
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 创建时间，RFC3339 UTC 格式（如 2026-06-01T09:00:00Z）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间，RFC3339 UTC 格式（如 2026-09-10T15:20:00Z）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// Agent 级 A2A 开关。false 恒输出（未开启不等于字段缺失）；A2AEndpoint / A2AStatus 由本接口在 A2A 开启时直接下发
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2AEnabled *bool `json:"A2AEnabled,omitnil,omitempty" name:"A2AEnabled"`

	// 历史会话总数（t_managed_agent_sessions 未软删计数，含全部状态）。注意与 DescribeAgent.ActiveSessionCount（活跃会话数）口径不同
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionCount *int64 `json:"SessionCount,omitnil,omitempty" name:"SessionCount"`

	// 最新版本的模型标识，取 latest_version_id 指向版本的 model；Agent 尚无版本时缺省
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 最新版本 ID（latest_version_id 转字符串，19 位雪花数字形态）；Agent 尚无版本时缺省
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestVersionId *string `json:"LatestVersionId,omitnil,omitempty" name:"LatestVersionId"`

	// 最新版本名（可能为 default / test-N / prod-N 任意类型）；Agent 尚无版本时缺省
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestVersionName *string `json:"LatestVersionName,omitnil,omitempty" name:"LatestVersionName"`

	// 对外 A2A card 发现地址（Agent Card JSON 地址），仅 A2AEnabled=true 的行下发；未注册 / registry 读失败时缺省
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2AEndpoint *string `json:"A2AEndpoint,omitnil,omitempty" name:"A2AEndpoint"`

	// A2A 注册态：DRAFT / REGISTERED / DISABLED / NONE / UNKNOWN，仅 A2AEnabled=true 的行下发，与 DescribeAgent.A2AConfig.A2AStatus 同枚举；用于「开关已开但地址尚未生成」的空态文案
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2AStatus *string `json:"A2AStatus,omitnil,omitempty" name:"A2AStatus"`

	// 公网链接访问开关。false 恒输出（未开启不等于字段缺失）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicApiEnabled *bool `json:"PublicApiEnabled,omitnil,omitempty" name:"PublicApiEnabled"`

	// 公网访问地址，仅 PublicApiEnabled=true 的行下发。固定拼法 https://{AgentId}-{region}.{endpoint_suffix}，与 DescribeAgentPublicAccess.Url 同规则；endpoint_suffix 未配置时为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicApiUrl *string `json:"PublicApiUrl,omitnil,omitempty" name:"PublicApiUrl"`

	// 创建人 UIN（建号时落库的 sub_account_uin；主账号自建时为主账号 uin）。注意语义为「实际操作建号的账号」
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 绑定的 OneID 企业账号 ID（数字字符串形态，如 1438693592234206274）；空=未绑定（缺省）。与 DescribeAgent.AgentInfo.AccountId 同源同语义；创建时可选传入，之后不可变
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`
}

type AgentVersionItem struct {
	// 版本 ID（雪花算法生成的数字字符串，唯一标识）
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 版本名称，形如 default / test-N / prod-N（N 为同类型版本的自增序号）
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 版本类型（服务端按 VersionName 派生）：DEFAULT（默认版本，可编辑）/ TEST（测试版本，可编辑）/ PROD（生产版本，内容冻结）
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionType *string `json:"VersionType,omitnil,omitempty" name:"VersionType"`

	// 版本绑定的模型标识；未设置时缺省
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 版本运行时使用的沙箱模板业务 ID；空字符串表示使用默认沙箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	SandboxTemplateId *string `json:"SandboxTemplateId,omitnil,omitempty" name:"SandboxTemplateId"`

	// 版本状态：DRAFT（草稿）/ ENABLED（已启用）/ DISABLED（已停用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>该版本累计承接的会话总数（历史累计值，只增不减）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionCount *int64 `json:"SessionCount,omitnil,omitempty" name:"SessionCount"`

	// 创建时间，RFC3339 UTC 格式（如 2026-08-01T10:00:00Z）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间，RFC3339 UTC 格式（如 2026-08-10T15:30:00Z）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`
}

// Predefined struct for user
type BindExternalAgentRequestParams struct {
	// TMA managed agent 业务 ID（CloudAgentID）
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 已注册的外部 A2A agent ID
	A2AAgentId *string `json:"A2AAgentId,omitnil,omitempty" name:"A2AAgentId"`

	// 版本 ID
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type BindExternalAgentRequest struct {
	*tchttp.BaseRequest
	
	// TMA managed agent 业务 ID（CloudAgentID）
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 已注册的外部 A2A agent ID
	A2AAgentId *string `json:"A2AAgentId,omitnil,omitempty" name:"A2AAgentId"`

	// 版本 ID
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *BindExternalAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindExternalAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "A2AAgentId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindExternalAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindExternalAgentResponseParams struct {
	// 操作结果状态（大写枚举）：BOUND=已绑定 / UNBOUND=已解绑
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 绑定记录 ID（自增 ID 字符串）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindingId *string `json:"BindingId,omitnil,omitempty" name:"BindingId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindExternalAgentResponse struct {
	*tchttp.BaseResponse
	Response *BindExternalAgentResponseParams `json:"Response"`
}

func (r *BindExternalAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindExternalAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BuiltinModel struct {
	// 模型唯一标识
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 供应商，如 TENCENT、OPENAI、ANTHROPIC、DEEPSEEK 等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vendor *string `json:"Vendor,omitnil,omitempty" name:"Vendor"`

	// 最大输出 Token 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxOutputTokens *int64 `json:"MaxOutputTokens,omitnil,omitempty" name:"MaxOutputTokens"`

	// 最大输入 Token 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxInputTokens *int64 `json:"MaxInputTokens,omitnil,omitempty" name:"MaxInputTokens"`

	// 是否支持函数调用（Tool Call）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportsToolCall *bool `json:"SupportsToolCall,omitnil,omitempty" name:"SupportsToolCall"`

	// 是否支持视觉（图片输入）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportsImages *bool `json:"SupportsImages,omitnil,omitempty" name:"SupportsImages"`

	// 模型中文描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescriptionZh *string `json:"DescriptionZh,omitnil,omitempty" name:"DescriptionZh"`

	// 模型英文描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescriptionEn *string `json:"DescriptionEn,omitnil,omitempty" name:"DescriptionEn"`

	// 模型标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 支持的客户端列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Clients []*string `json:"Clients,omitnil,omitempty" name:"Clients"`

	// 服务接入地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceEndpoint *string `json:"ServiceEndpoint,omitnil,omitempty" name:"ServiceEndpoint"`

	// 状态：ENABLED（已启用）/ DISABLED（已停用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 本企业内绑定该模型的 Agent 数（过滤软删除 Agent/版本与调试 Agent）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentCount *int64 `json:"AgentCount,omitnil,omitempty" name:"AgentCount"`
}

type ChatEndpoint struct {
	// 接入点类型：PUBLIC（公网）/ PRIVATE（私网，预留）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndpointType *string `json:"EndpointType,omitnil,omitempty" name:"EndpointType"`

	// 接入点地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type ConnectorInfo struct {
	// 连接器 ID
	ConnectorId *string `json:"ConnectorId,omitnil,omitempty" name:"ConnectorId"`

	// 连接器短标识（终身不变，跨版本稳定）
	ConnectorSlug *string `json:"ConnectorSlug,omitnil,omitempty" name:"ConnectorSlug"`

	// 版本级连接器密钥
	ConnectorKey *string `json:"ConnectorKey,omitnil,omitempty" name:"ConnectorKey"`

	// 连接器名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 连接器描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 头像 URL
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 连接器来源：ENTERPRISE_AGENT / ASSISTANT
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 归属企业 ID
	EnterpriseId *string `json:"EnterpriseId,omitnil,omitempty" name:"EnterpriseId"`

	// 连接器类型：MCP_SERVER / A2A / API_SERVICE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 上游服务地址
	ServiceUrl *string `json:"ServiceUrl,omitnil,omitempty" name:"ServiceUrl"`

	// 授权方式列表：NONE / ONEID / OAUTH2_IDP
	AuthModes []*string `json:"AuthModes,omitnil,omitempty" name:"AuthModes"`

	// 最新版本号
	LatestVersionNo *int64 `json:"LatestVersionNo,omitnil,omitempty" name:"LatestVersionNo"`

	// 连接器状态：ACTIVE / DISABLED
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建人 ID
	CreatorId *string `json:"CreatorId,omitnil,omitempty" name:"CreatorId"`

	// 创建时间（ISO8601，UTC）
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 最后修改时间（ISO8601，UTC）
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`
}

type ConnectorRefInput struct {
	// connector 主表 ID（雪花 ID 数字串）
	ConnectorId *string `json:"ConnectorId,omitnil,omitempty" name:"ConnectorId"`
}

// Predefined struct for user
type CreateAgentRequestParams struct {
	// Agent 名称
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// Agent 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 头像 URL
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 模型标识
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// Manifest v2.0 原文（JSON 字符串），作为 default 版本初始内容。ConnectorSet 非空时 Manifest 不可为空，否则返回 InvalidParameter
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// 该 Agent 最终绑定的连接器集合（全量覆盖语义）：缺省 = 不绑定连接器；非空 = 物化为 manifest v2 mcp_servers 网关条目。ConnectorSet 非空时 Manifest 不可为空，否则返回 InvalidParameter
	ConnectorSet []*ConnectorRefInput `json:"ConnectorSet,omitnil,omitempty" name:"ConnectorSet"`

	// 绑定的 OneID 企业账号 ID。非空时必须是当前主账号已在企业授权表（t_managed_agent_enterprise_authorization）中授权的租户，否则返回 UnauthorizedOperation.AccountNotAuthorized。绑定后不可修改。TrimSpace 后长度 1~64 字符
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`
}

type CreateAgentRequest struct {
	*tchttp.BaseRequest
	
	// Agent 名称
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// Agent 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 头像 URL
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 模型标识
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// Manifest v2.0 原文（JSON 字符串），作为 default 版本初始内容。ConnectorSet 非空时 Manifest 不可为空，否则返回 InvalidParameter
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// 该 Agent 最终绑定的连接器集合（全量覆盖语义）：缺省 = 不绑定连接器；非空 = 物化为 manifest v2 mcp_servers 网关条目。ConnectorSet 非空时 Manifest 不可为空，否则返回 InvalidParameter
	ConnectorSet []*ConnectorRefInput `json:"ConnectorSet,omitnil,omitempty" name:"ConnectorSet"`

	// 绑定的 OneID 企业账号 ID。非空时必须是当前主账号已在企业授权表（t_managed_agent_enterprise_authorization）中授权的租户，否则返回 UnauthorizedOperation.AccountNotAuthorized。绑定后不可修改。TrimSpace 后长度 1~64 字符
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`
}

func (r *CreateAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentName")
	delete(f, "Description")
	delete(f, "AvatarUrl")
	delete(f, "Model")
	delete(f, "Manifest")
	delete(f, "ConnectorSet")
	delete(f, "AccountId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentResponseParams struct {
	// Agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Agent 名称
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// Agent 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 头像 URL
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 是否调试 Agent
	IsDebug *bool `json:"IsDebug,omitnil,omitempty" name:"IsDebug"`

	// 创建时间（RFC3339）
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间（RFC3339）
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 流量路由配置（VersionId 恒为字符串，防 JS 精度丢失）
	RoutingSet []*RoutingItem `json:"RoutingSet,omitnil,omitempty" name:"RoutingSet"`

	// A2A 对外互通配置与注册态（只读回显；原四个平铺字段收进结构）
	A2AConfig *A2AConfig `json:"A2AConfig,omitnil,omitempty" name:"A2AConfig"`

	// 绑定的 OneID 企业账号 ID。允许为空：未绑定的存量与新建 Agent 该字段缺省，绑定后回显绑定值
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentResponseParams `json:"Response"`
}

func (r *CreateAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentSessionRequestParams struct {
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>指定版本 ID（可选）。非空且合法时固定使用该版本，跳过 routing_config 权重挑选；指定版本需归属同一 Agent 且未被废弃</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type CreateAgentSessionRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>指定版本 ID（可选）。非空且合法时固定使用该版本，跳过 routing_config 权重挑选；指定版本需归属同一 Agent 且未被废弃</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *CreateAgentSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentSessionResponseParams struct {
	// <p>会话 ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>可用的聊天接入点列表（当前仅含一个 PUBLIC 公网接入点；空数组 = 无可用接入点）</p>
	EndpointSet []*ChatEndpoint `json:"EndpointSet,omitnil,omitempty" name:"EndpointSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgentSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentSessionResponseParams `json:"Response"`
}

func (r *CreateAgentSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentVersionFromSourceRequestParams struct {
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>源版本 ID，同 Agent 下未 DISABLED 的任意版本</p>
	SourceVersionId *string `json:"SourceVersionId,omitnil,omitempty" name:"SourceVersionId"`

	// <p>可选，覆盖源版本的 Model</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>可选，覆盖源版本的 Description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>可选，完整 v2.0 manifest JSON 字符串；传入则整体覆盖源版本 manifest</p>
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// <p>沙箱模板 ID。可选，patch 语义：null 沿用源版本绑定的模板；空串解绑（恢复系统默认模板）；非空时模板须属于当前企业且可用（未删除、状态正常）。</p>
	SandboxTemplateId *string `json:"SandboxTemplateId,omitnil,omitempty" name:"SandboxTemplateId"`
}

type CreateAgentVersionFromSourceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>源版本 ID，同 Agent 下未 DISABLED 的任意版本</p>
	SourceVersionId *string `json:"SourceVersionId,omitnil,omitempty" name:"SourceVersionId"`

	// <p>可选，覆盖源版本的 Model</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>可选，覆盖源版本的 Description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>可选，完整 v2.0 manifest JSON 字符串；传入则整体覆盖源版本 manifest</p>
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// <p>沙箱模板 ID。可选，patch 语义：null 沿用源版本绑定的模板；空串解绑（恢复系统默认模板）；非空时模板须属于当前企业且可用（未删除、状态正常）。</p>
	SandboxTemplateId *string `json:"SandboxTemplateId,omitnil,omitempty" name:"SandboxTemplateId"`
}

func (r *CreateAgentVersionFromSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentVersionFromSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "SourceVersionId")
	delete(f, "Model")
	delete(f, "Description")
	delete(f, "Manifest")
	delete(f, "SandboxTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentVersionFromSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentVersionFromSourceResponseParams struct {
	// <p>版本 ID</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>版本名称</p>
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// <p>版本类型：DEFAULT / TEST / PROD</p>
	VersionType *string `json:"VersionType,omitnil,omitempty" name:"VersionType"`

	// <p>版本变更说明</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>模型标识</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>Manifest v2.0 精简 manifest 原文（JSON 字符串）</p>
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// <p>版本状态：DRAFT / ENABLED / DISABLED</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>创建时间</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>更新时间</p>
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// <p>绑定的沙箱模板 ID；未绑定时为空，创建会话沙箱使用系统默认模板。</p>
	SandboxTemplateId *string `json:"SandboxTemplateId,omitnil,omitempty" name:"SandboxTemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgentVersionFromSourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentVersionFromSourceResponseParams `json:"Response"`
}

func (r *CreateAgentVersionFromSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentVersionFromSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentVersionRequestParams struct {
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Manifest v2.0 精简 manifest 原文（JSON 对象序列化后的字符串）</p>
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// <p>模型标识</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>版本变更说明</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>沙箱模板 ID。可选；传入时模板须属于当前企业且可用（未删除、状态正常），绑定到新建的 test/prod 版本。</p>
	SandboxTemplateId *string `json:"SandboxTemplateId,omitnil,omitempty" name:"SandboxTemplateId"`
}

type CreateAgentVersionRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Manifest v2.0 精简 manifest 原文（JSON 对象序列化后的字符串）</p>
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// <p>模型标识</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>版本变更说明</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>沙箱模板 ID。可选；传入时模板须属于当前企业且可用（未删除、状态正常），绑定到新建的 test/prod 版本。</p>
	SandboxTemplateId *string `json:"SandboxTemplateId,omitnil,omitempty" name:"SandboxTemplateId"`
}

func (r *CreateAgentVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "Manifest")
	delete(f, "Model")
	delete(f, "Description")
	delete(f, "SandboxTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentVersionResponseParams struct {
	// <p>版本 ID</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>版本名称</p>
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// <p>版本类型：DEFAULT / TEST / PROD</p>
	VersionType *string `json:"VersionType,omitnil,omitempty" name:"VersionType"`

	// <p>版本变更说明</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>模型标识</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>Manifest v2.0 精简 manifest 原文（JSON 字符串）</p>
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// <p>版本状态：DRAFT / ENABLED / DISABLED</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>创建时间</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>更新时间</p>
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// <p>绑定的沙箱模板 ID；未绑定时为空，创建会话沙箱使用系统默认模板。</p>
	SandboxTemplateId *string `json:"SandboxTemplateId,omitnil,omitempty" name:"SandboxTemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgentVersionResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentVersionResponseParams `json:"Response"`
}

func (r *CreateAgentVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentRequestParams struct {
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type DeleteAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *DeleteAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAgentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAgentResponseParams `json:"Response"`
}

func (r *DeleteAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentListRequestParams struct {
	// 偏移量，从 0 开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，缺省为 20，最大 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件数组，多个 Filter 之间为 AND 关系，同一 Filter 内多个 Values 为 OR 关系
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方向：ASC / DESC
	SortDirection *string `json:"SortDirection,omitnil,omitempty" name:"SortDirection"`
}

type DescribeAgentListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，从 0 开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，缺省为 20，最大 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件数组，多个 Filter 之间为 AND 关系，同一 Filter 内多个 Values 为 OR 关系
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方向：ASC / DESC
	SortDirection *string `json:"SortDirection,omitnil,omitempty" name:"SortDirection"`
}

func (r *DescribeAgentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "SortDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentListResponseParams struct {
	// 符合条件的 Agent 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Agent 列表（分页后）；元素含 A2A / 公网 API 访问开关与地址、创建人 UIN、绑定的企业账号 ID
	AgentSet []*AgentItem `json:"AgentSet,omitnil,omitempty" name:"AgentSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentListResponseParams `json:"Response"`
}

func (r *DescribeAgentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentRequestParams struct {
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type DescribeAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *DescribeAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentResponseParams struct {
	// Agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Agent 名称
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// Agent 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 头像 URL
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 是否调试 Agent
	IsDebug *bool `json:"IsDebug,omitnil,omitempty" name:"IsDebug"`

	// 创建时间（RFC3339）
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间（RFC3339）
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 当前活跃 session 数（ACTIVE/CREATING/MIGRATING，未软删）；仅 DescribeAgent 读路径填充，写路径回显不下发
	ActiveSessionCount *int64 `json:"ActiveSessionCount,omitnil,omitempty" name:"ActiveSessionCount"`

	// 流量路由配置（VersionId 恒为字符串，防 JS 精度丢失）
	RoutingSet []*RoutingItem `json:"RoutingSet,omitnil,omitempty" name:"RoutingSet"`

	// A2A 对外互通配置与注册态（只读回显；原四个平铺字段收进结构）
	A2AConfig *A2AConfig `json:"A2AConfig,omitnil,omitempty" name:"A2AConfig"`

	// 绑定的 OneID 企业账号 ID。允许为空：未绑定的存量与新建 Agent 该字段缺省，绑定后回显绑定值
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentResponseParams `json:"Response"`
}

func (r *DescribeAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSessionListRequestParams struct {
	// 偏移量，从 0 开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，缺省为 20，最大 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方向：ASC / DESC
	SortDirection *string `json:"SortDirection,omitnil,omitempty" name:"SortDirection"`

	// 过滤条件数组，多个 Filter 之间为 AND 关系，同一 Filter 内多个 Values 为 OR 关系
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAgentSessionListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，从 0 开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，缺省为 20，最大 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方向：ASC / DESC
	SortDirection *string `json:"SortDirection,omitnil,omitempty" name:"SortDirection"`

	// 过滤条件数组，多个 Filter 之间为 AND 关系，同一 Filter 内多个 Values 为 OR 关系
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAgentSessionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentSessionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "SortDirection")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentSessionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSessionListResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 会话列表
	SessionSet []*SessionItem `json:"SessionSet,omitnil,omitempty" name:"SessionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentSessionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentSessionListResponseParams `json:"Response"`
}

func (r *DescribeAgentSessionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentSessionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSessionRequestParams struct {
	// 会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeAgentSessionRequest struct {
	*tchttp.BaseRequest
	
	// 会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeAgentSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSessionResponseParams struct {
	// 会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 会话名称（AgentOS 侧生成的 AI 标题 / 用户改名）；缺失时为空，调用方可兜底展示 SessionId 后缀
	SessionName *string `json:"SessionName,omitnil,omitempty" name:"SessionName"`

	// Agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Agent 名称
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// 版本 ID
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 会话使用的版本名称（与 VersionId 区分：此为版本名，非 ID）
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 版本状态：DRAFT / ENABLED / DISABLED
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建者 Uin
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 连接器来源：ENTERPRISE_AGENT / ASSISTANT
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 可用的聊天接入点列表（详情独有）
	EndpointSet []*ChatEndpoint `json:"EndpointSet,omitnil,omitempty" name:"EndpointSet"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间（RFC3339）
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentSessionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentSessionResponseParams `json:"Response"`
}

func (r *DescribeAgentSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentVersionListRequestParams struct {
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>偏移量，从 0 开始</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，缺省为 20，最大 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件数组，多个 Filter 之间为 AND 关系，同一 Filter 内多个 Values 为 OR 关系</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAgentVersionListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>偏移量，从 0 开始</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，缺省为 20，最大 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件数组，多个 Filter 之间为 AND 关系，同一 Filter 内多个 Values 为 OR 关系</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAgentVersionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentVersionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentVersionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentVersionListResponseParams struct {
	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>版本列表（原 VersionSet；集合名带实体前缀以区分 Skill 版本接口的同名字段）</p>
	AgentVersionSet []*AgentVersionItem `json:"AgentVersionSet,omitnil,omitempty" name:"AgentVersionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentVersionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentVersionListResponseParams `json:"Response"`
}

func (r *DescribeAgentVersionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentVersionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentVersionRequestParams struct {
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>版本 ID</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type DescribeAgentVersionRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>版本 ID</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *DescribeAgentVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentVersionResponseParams struct {
	// <p>版本 ID</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>版本名称</p>
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// <p>版本类型：DEFAULT / TEST / PROD</p>
	VersionType *string `json:"VersionType,omitnil,omitempty" name:"VersionType"`

	// <p>版本变更说明</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>模型标识</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>Manifest v2.0 精简 manifest 原文（JSON 字符串）</p>
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// <p>版本状态：DRAFT / ENABLED / DISABLED</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>创建时间</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>更新时间</p>
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// <p>绑定的沙箱模板 ID；未绑定时为空，创建会话沙箱使用系统默认模板。</p>
	SandboxTemplateId *string `json:"SandboxTemplateId,omitnil,omitempty" name:"SandboxTemplateId"`

	// <p>该版本累计承接的会话总数（历史累计值，只增不减）</p>
	SessionCount *int64 `json:"SessionCount,omitnil,omitempty" name:"SessionCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentVersionResponseParams `json:"Response"`
}

func (r *DescribeAgentVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuiltinModelListRequestParams struct {
	// 偏移量，从 0 开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，缺省为 20，最大 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件数组，多个 Filter 之间为 AND 关系，同一 Filter 内多个 Values 为 OR 关系
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// OneID 企业账号 ID，可选。传入时拉取该账号对应企业的模型（要求当前主账号 UIN 已授权该账号），不传时使用服务配置的企业 ID
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`
}

type DescribeBuiltinModelListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，从 0 开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，缺省为 20，最大 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件数组，多个 Filter 之间为 AND 关系，同一 Filter 内多个 Values 为 OR 关系
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// OneID 企业账号 ID，可选。传入时拉取该账号对应企业的模型（要求当前主账号 UIN 已授权该账号），不传时使用服务配置的企业 ID
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`
}

func (r *DescribeBuiltinModelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuiltinModelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "AccountId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBuiltinModelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuiltinModelListResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 内置模型列表（分页后）
	BuiltinModelSet []*BuiltinModel `json:"BuiltinModelSet,omitnil,omitempty" name:"BuiltinModelSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBuiltinModelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBuiltinModelListResponseParams `json:"Response"`
}

func (r *DescribeBuiltinModelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuiltinModelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConnectorListRequestParams struct {
	// 过滤条件数组，多个 Filter 之间为 AND 关系。支持 Name：Name（名称模糊匹配）/ Status（ACTIVE / DISABLED）/ Source（ENTERPRISE_AGENT / ASSISTANT）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已废弃：服务端不再读取，请使用 Offset/Limit。字段保留仅为过渡兼容，后续下线
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 已废弃：服务端不再读取，请使用 Offset/Limit。字段保留仅为过渡兼容，后续下线
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 偏移量，0 基准，缺省 0（标准 CAPI 分页参数）
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，取值 1-100，缺省 20（标准 CAPI 分页参数）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeConnectorListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件数组，多个 Filter 之间为 AND 关系。支持 Name：Name（名称模糊匹配）/ Status（ACTIVE / DISABLED）/ Source（ENTERPRISE_AGENT / ASSISTANT）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已废弃：服务端不再读取，请使用 Offset/Limit。字段保留仅为过渡兼容，后续下线
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 已废弃：服务端不再读取，请使用 Offset/Limit。字段保留仅为过渡兼容，后续下线
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 偏移量，0 基准，缺省 0（标准 CAPI 分页参数）
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，取值 1-100，缺省 20（标准 CAPI 分页参数）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeConnectorListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConnectorListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConnectorListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConnectorListResponseParams struct {
	// 符合条件的连接器总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 连接器列表（分页后）；连接器挂调用方主账号 UIN 下，不挂 OneID 企业
	ConnectorSet []*ConnectorInfo `json:"ConnectorSet,omitnil,omitempty" name:"ConnectorSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConnectorListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConnectorListResponseParams `json:"Response"`
}

func (r *DescribeConnectorListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConnectorListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExpertListRequestParams struct {
	// <p>专家来源，必填：BUILTIN（内置）/ CUSTOM（自建）</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>标准过滤条件：ExpertId（精确，多值 OR，携带即按 ID 批量查询，忽略分页）/ Keyword（模糊）</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>偏移量，从 0 开始，默认 0（按 ID 批量查询时忽略）</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量，默认 20，最大 200（按 ID 批量查询时忽略）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeExpertListRequest struct {
	*tchttp.BaseRequest
	
	// <p>专家来源，必填：BUILTIN（内置）/ CUSTOM（自建）</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>标准过滤条件：ExpertId（精确，多值 OR，携带即按 ID 批量查询，忽略分页）/ Keyword（模糊）</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>偏移量，从 0 开始，默认 0（按 ID 批量查询时忽略）</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量，默认 20，最大 200（按 ID 批量查询时忽略）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeExpertListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExpertListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Source")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExpertListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExpertListResponseParams struct {
	// <p>符合条件的专家总数（按 ID 批量时为实际命中数）</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>专家列表</p>
	ExpertSet []*ExpertItem `json:"ExpertSet,omitnil,omitempty" name:"ExpertSet"`

	// <p>全局计数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Counts *ExpertCounts `json:"Counts,omitnil,omitempty" name:"Counts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExpertListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExpertListResponseParams `json:"Response"`
}

func (r *DescribeExpertListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExpertListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExternalAgentListRequestParams struct {
	// Agent 业务 ID（必填：绑定状态的归属主体）
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 标准过滤条件，支持的 Name：Bound（BOUND=仅已绑定 / UNBOUND=仅未绑定 / ALL=全部，缺省 ALL）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，从 0 开始，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认 20，最大 200
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 外部 Agent 列表查询关键字
	DescribeExternalAgentList *string `json:"DescribeExternalAgentList,omitnil,omitempty" name:"DescribeExternalAgentList"`

	// 版本 ID
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type DescribeExternalAgentListRequest struct {
	*tchttp.BaseRequest
	
	// Agent 业务 ID（必填：绑定状态的归属主体）
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 标准过滤条件，支持的 Name：Bound（BOUND=仅已绑定 / UNBOUND=仅未绑定 / ALL=全部，缺省 ALL）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，从 0 开始，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认 20，最大 200
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 外部 Agent 列表查询关键字
	DescribeExternalAgentList *string `json:"DescribeExternalAgentList,omitnil,omitempty" name:"DescribeExternalAgentList"`

	// 版本 ID
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *DescribeExternalAgentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExternalAgentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DescribeExternalAgentList")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExternalAgentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExternalAgentListResponseParams struct {
	// 符合条件的外部 Agent 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 外部 Agent 集合（可见卡片全集 ∪ URL 直连型存量 binding 的合并视图）
	ExternalAgentSet []*ExternalAgentInfo `json:"ExternalAgentSet,omitnil,omitempty" name:"ExternalAgentSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExternalAgentListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExternalAgentListResponseParams `json:"Response"`
}

func (r *DescribeExternalAgentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExternalAgentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExternalAgentRequestParams struct {
	// <p>TMA managed agent 业务 ID（CloudAgentID）</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>已绑定的外部 A2A agent ID</p>
	A2AAgentId *string `json:"A2AAgentId,omitnil,omitempty" name:"A2AAgentId"`
}

type DescribeExternalAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>TMA managed agent 业务 ID（CloudAgentID）</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>已绑定的外部 A2A agent ID</p>
	A2AAgentId *string `json:"A2AAgentId,omitnil,omitempty" name:"A2AAgentId"`
}

func (r *DescribeExternalAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExternalAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "A2AAgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExternalAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExternalAgentResponseParams struct {
	// <p>外部 A2A agent ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2AAgentId *string `json:"A2AAgentId,omitnil,omitempty" name:"A2AAgentId"`

	// <p>外部 Agent 名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>外部 A2A Server URL</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// <p>绑定记录 ID（已绑定时返回）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindingId *string `json:"BindingId,omitnil,omitempty" name:"BindingId"`

	// <p>是否已绑定到当前 Agent</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bound *bool `json:"Bound,omitnil,omitempty" name:"Bound"`

	// <p>头像地址（取自 provider card 的 iconUrl；为空时前端回落首字母头像）</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>外部 agent card 声明的版本号</p>
	A2AVersion *string `json:"A2AVersion,omitnil,omitempty" name:"A2AVersion"`

	// <p>A2A card skills 集合</p>
	A2ASkillSet []*A2ASkillItem `json:"A2ASkillSet,omitnil,omitempty" name:"A2ASkillSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExternalAgentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExternalAgentResponseParams `json:"Response"`
}

func (r *DescribeExternalAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExternalAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageEventListRequestParams struct {
	// <p>Session ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认 100，最大 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件数组，多个 Filter 之间为 AND 关系，同一 Filter 内多个 Values 为 OR 关系</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeMessageEventListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认 100，最大 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件数组，多个 Filter 之间为 AND 关系，同一 Filter 内多个 Values 为 OR 关系</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeMessageEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "AgentId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMessageEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageEventListResponseParams struct {
	// <p>消息事件列表</p>
	MessageEventSet []*MessageEvent `json:"MessageEventSet,omitnil,omitempty" name:"MessageEventSet"`

	// <p>符合过滤条件的事件总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMessageEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMessageEventListResponseParams `json:"Response"`
}

func (r *DescribeMessageEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillListRequestParams struct {
	// 技能来源，必填：BUILTIN（内置）/ CUSTOM（自建）/ AUTHORIZED（企业授权）。数据通路判别，非筛选条件
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 标准过滤条件：SkillId（精确，多值 OR ≤100，携带即按 ID 批量查询）/ Keyword（模糊）/ PublishStatus（DRAFT/PUBLISHED/ALL）/ Status（ENABLED/DISABLED/ALL）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认 0（按 ID 批量查询时忽略）
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认 20，最大 200（按 ID 批量查询时忽略）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 授权方企业账号标识；仅 Source=AUTHORIZED 时生效。不传则由后端用 Uin 推导全部已授权范围；未携带 SkillId 的分页查询必传
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 仅 Source=AUTHORIZED 时生效。Agent 绑定了 OneID 租户时，授权集合强制收窄到绑定租户；显式传入的 AccountId 必须等于绑定值，否则请求被拒绝。绑定 Agent 的分页查询可不传 AccountId（服务端按绑定值收窄到单一授权方）
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type DescribeSkillListRequest struct {
	*tchttp.BaseRequest
	
	// 技能来源，必填：BUILTIN（内置）/ CUSTOM（自建）/ AUTHORIZED（企业授权）。数据通路判别，非筛选条件
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 标准过滤条件：SkillId（精确，多值 OR ≤100，携带即按 ID 批量查询）/ Keyword（模糊）/ PublishStatus（DRAFT/PUBLISHED/ALL）/ Status（ENABLED/DISABLED/ALL）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认 0（按 ID 批量查询时忽略）
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认 20，最大 200（按 ID 批量查询时忽略）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 授权方企业账号标识；仅 Source=AUTHORIZED 时生效。不传则由后端用 Uin 推导全部已授权范围；未携带 SkillId 的分页查询必传
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 仅 Source=AUTHORIZED 时生效。Agent 绑定了 OneID 租户时，授权集合强制收窄到绑定租户；显式传入的 AccountId 必须等于绑定值，否则请求被拒绝。绑定 Agent 的分页查询可不传 AccountId（服务端按绑定值收窄到单一授权方）
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *DescribeSkillListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Source")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AccountId")
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillListResponseParams struct {
	// 符合条件的技能总数（按 ID 批量时为实际命中数）
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 技能列表（仅列表展示所需字段，完整信息走 DescribeSkill）
	SkillSet []*SkillItem `json:"SkillSet,omitnil,omitempty" name:"SkillSet"`

	// 全局计数（不受 keyword 影响）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Counts *SkillCounts `json:"Counts,omitnil,omitempty" name:"Counts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillListResponseParams `json:"Response"`
}

func (r *DescribeSkillListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserAccessTokenRequestParams struct {

}

type DescribeUserAccessTokenRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserAccessTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserAccessTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserAccessTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserAccessTokenResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserAccessTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserAccessTokenResponseParams `json:"Response"`
}

func (r *DescribeUserAccessTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserAccessTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpertCounts struct {
	// <p>内置专家数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Builtin *int64 `json:"Builtin,omitnil,omitempty" name:"Builtin"`

	// <p>自建专家数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Custom *int64 `json:"Custom,omitnil,omitempty" name:"Custom"`

	// <p>总数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type ExpertItem struct {
	// <p>专家来源：builtin、custom</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>展示名</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>图标 URL</p>
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// <p>是否启用</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>下载 URL</p>
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// <p>更新时间</p>
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// <p>启停状态：enabled、disabled</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>专家标识</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpertId *string `json:"ExpertId,omitnil,omitempty" name:"ExpertId"`

	// <p>当前生效版本号</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpertVersion *string `json:"ExpertVersion,omitnil,omitempty" name:"ExpertVersion"`
}

type ExternalAgentInfo struct {
	// 外部 A2A agent ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2AAgentId *string `json:"A2AAgentId,omitnil,omitempty" name:"A2AAgentId"`

	// 外部 Agent 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 外部 A2A Server URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 绑定记录 ID（已绑定时返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindingId *string `json:"BindingId,omitnil,omitempty" name:"BindingId"`

	// 是否已绑定到当前 Agent
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bound *bool `json:"Bound,omitnil,omitempty" name:"Bound"`

	// 头像地址（取自 provider card 的 iconUrl）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// 外部 agent card 声明的版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2AVersion *string `json:"A2AVersion,omitnil,omitempty" name:"A2AVersion"`

	// A2A card skills 集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	A2ASkillSet []*A2ASkillItem `json:"A2ASkillSet,omitnil,omitempty" name:"A2ASkillSet"`
}

type Filter struct {
	// 过滤属性名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤值列表（同一 Filter 内多个值为 OR 关系）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否精确匹配，默认 false（模糊匹配）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

type MessageEvent struct {
	// <p>序号</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sequence *int64 `json:"Sequence,omitnil,omitempty" name:"Sequence"`

	// <p>类型 USER/TOOL/ASSISTANT</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>发生时间 ISO8601</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OccurredAt *string `json:"OccurredAt,omitnil,omitempty" name:"OccurredAt"`

	// <p>消息内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *MessageEventMessage `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>工具调用</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolCall *MessageEventToolCall `json:"ToolCall,omitnil,omitempty" name:"ToolCall"`
}

type MessageEventMessage struct {
	// <p>文本内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>Token用量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenUsage *TokenUsage `json:"TokenUsage,omitnil,omitempty" name:"TokenUsage"`
}

type MessageEventToolCall struct {
	// <p>调用ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolCallId *string `json:"ToolCallId,omitnil,omitempty" name:"ToolCallId"`

	// <p>工具名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// <p>状态 PENDING/IN_PROGRESS/SUCCEEDED/FAILED</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>工具调用Input（已递归脱敏，JSON 字符串）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// <p>工具调用Output（已递归脱敏，JSON 字符串）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// <p>结束时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndedAt *string `json:"EndedAt,omitnil,omitempty" name:"EndedAt"`

	// <p>耗时毫秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DurationMs *int64 `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// <p>调用开始时间（RFC3339 格式）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartedAt *string `json:"StartedAt,omitnil,omitempty" name:"StartedAt"`
}

// Predefined struct for user
type MigrateAgentSessionRequestParams struct {
	// <p>待迁移的会话 ID（必填）</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>目标 Agent 业务 ID（必填），必须与 Session 原 Agent 相同</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>目标版本 ID（必填，字符串形式）。需归属同一 Agent 且未被废弃</p>
	TargetVersionId *string `json:"TargetVersionId,omitnil,omitempty" name:"TargetVersionId"`
}

type MigrateAgentSessionRequest struct {
	*tchttp.BaseRequest
	
	// <p>待迁移的会话 ID（必填）</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>目标 Agent 业务 ID（必填），必须与 Session 原 Agent 相同</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>目标版本 ID（必填，字符串形式）。需归属同一 Agent 且未被废弃</p>
	TargetVersionId *string `json:"TargetVersionId,omitnil,omitempty" name:"TargetVersionId"`
}

func (r *MigrateAgentSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateAgentSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "AgentId")
	delete(f, "TargetVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MigrateAgentSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigrateAgentSessionResponseParams struct {
	// <p>会话 ID（回显原值，保持不变）</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>迁移后的会话状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MigrateAgentSessionResponse struct {
	*tchttp.BaseResponse
	Response *MigrateAgentSessionResponseParams `json:"Response"`
}

func (r *MigrateAgentSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateAgentSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentA2AConfigRequestParams struct {
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent 级唯一 A2A 开关</p>
	A2AEnabled *bool `json:"A2AEnabled,omitnil,omitempty" name:"A2AEnabled"`

	// <p>A2A 技能集合（原 A2ASkills）</p>
	A2ASkillSet []*A2ASkillInput `json:"A2ASkillSet,omitnil,omitempty" name:"A2ASkillSet"`
}

type ModifyAgentA2AConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent 级唯一 A2A 开关</p>
	A2AEnabled *bool `json:"A2AEnabled,omitnil,omitempty" name:"A2AEnabled"`

	// <p>A2A 技能集合（原 A2ASkills）</p>
	A2ASkillSet []*A2ASkillInput `json:"A2ASkillSet,omitnil,omitempty" name:"A2ASkillSet"`
}

func (r *ModifyAgentA2AConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentA2AConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "A2AEnabled")
	delete(f, "A2ASkillSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAgentA2AConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentA2AConfigResponseParams struct {
	// <p>Agent 业务 ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent 名称</p>
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// <p>Agent 描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>头像 URL</p>
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// <p>是否调试 Agent</p>
	IsDebug *bool `json:"IsDebug,omitnil,omitempty" name:"IsDebug"`

	// <p>创建时间（RFC3339）</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>更新时间（RFC3339）</p>
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// <p>流量路由配置（VersionId 恒为字符串，防 JS 精度丢失）</p>
	RoutingSet []*RoutingItem `json:"RoutingSet,omitnil,omitempty" name:"RoutingSet"`

	// <p>A2A 对外互通配置与注册态（只读回显；原四个平铺字段收进结构）</p>
	A2AConfig *A2AConfig `json:"A2AConfig,omitnil,omitempty" name:"A2AConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAgentA2AConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAgentA2AConfigResponseParams `json:"Response"`
}

func (r *ModifyAgentA2AConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentA2AConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentRequestParams struct {
	// Agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Agent 名称（可选，仅传递需要更新的字段）
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// Agent 描述（可选）
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 头像 URL（可选）
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`
}

type ModifyAgentRequest struct {
	*tchttp.BaseRequest
	
	// Agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Agent 名称（可选，仅传递需要更新的字段）
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// Agent 描述（可选）
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 头像 URL（可选）
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`
}

func (r *ModifyAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "AgentName")
	delete(f, "Description")
	delete(f, "AvatarUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentResponseParams struct {
	// Agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Agent 名称
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// Agent 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 头像 URL
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 是否调试 Agent
	IsDebug *bool `json:"IsDebug,omitnil,omitempty" name:"IsDebug"`

	// 创建时间（RFC3339）
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间（RFC3339）
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 流量路由配置（VersionId 恒为字符串，防 JS 精度丢失）
	RoutingSet []*RoutingItem `json:"RoutingSet,omitnil,omitempty" name:"RoutingSet"`

	// A2A 对外互通配置与注册态（只读回显；原四个平铺字段收进结构）
	A2AConfig *A2AConfig `json:"A2AConfig,omitnil,omitempty" name:"A2AConfig"`

	// 绑定的 OneID 企业账号 ID。允许为空：未绑定的存量与新建 Agent 该字段缺省，绑定后回显绑定值
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAgentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAgentResponseParams `json:"Response"`
}

func (r *ModifyAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentRoutingRequestParams struct {
	// Agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 路由配置，覆盖式写入（与出参 AgentInfo.RoutingSet 命名对齐）
	RoutingSet []*RoutingItem `json:"RoutingSet,omitnil,omitempty" name:"RoutingSet"`
}

type ModifyAgentRoutingRequest struct {
	*tchttp.BaseRequest
	
	// Agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 路由配置，覆盖式写入（与出参 AgentInfo.RoutingSet 命名对齐）
	RoutingSet []*RoutingItem `json:"RoutingSet,omitnil,omitempty" name:"RoutingSet"`
}

func (r *ModifyAgentRoutingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentRoutingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "RoutingSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAgentRoutingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentRoutingResponseParams struct {
	// Agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Agent 名称
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// Agent 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 头像 URL
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 是否调试 Agent
	IsDebug *bool `json:"IsDebug,omitnil,omitempty" name:"IsDebug"`

	// 创建时间（RFC3339）
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间（RFC3339）
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 流量路由配置（VersionId 恒为字符串，防 JS 精度丢失）
	RoutingSet []*RoutingItem `json:"RoutingSet,omitnil,omitempty" name:"RoutingSet"`

	// A2A 对外互通配置与注册态（只读回显；原四个平铺字段收进结构）
	A2AConfig *A2AConfig `json:"A2AConfig,omitnil,omitempty" name:"A2AConfig"`

	// 绑定的 OneID 企业账号 ID。允许为空：未绑定的存量与新建 Agent 该字段缺省，绑定后回显绑定值
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAgentRoutingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAgentRoutingResponseParams `json:"Response"`
}

func (r *ModifyAgentRoutingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentRoutingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentVersionRequestParams struct {
	// Agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 版本 ID（仅 default 或 test 版本可原地更新，prod 拒绝）
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Manifest v2.0 原文（可选；Manifest / Model / Description / SandboxTemplateId / ConnectorSet 五个可选字段至少提供一个）
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// 模型标识（可选）
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 版本变更说明（可选）
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 沙箱模板 ID。可选，patch 语义：null 不修改；空串解绑（恢复系统默认模板）；非空时模板须属于当前企业且可用（未删除、状态正常）。
	SandboxTemplateId *string `json:"SandboxTemplateId,omitnil,omitempty" name:"SandboxTemplateId"`

	// 该版本最终绑定的连接器集合（全量覆盖语义）：缺省 = 本次不改动连接器绑定；空数组 = 解绑全部连接器；非空 = 物化为 manifest v2 mcp_servers 网关条目，manifest 中不在本集合内的连接器条目会被移除（解绑在服务端闭环，无需调用方改写 Manifest）
	ConnectorSet []*ConnectorRefInput `json:"ConnectorSet,omitnil,omitempty" name:"ConnectorSet"`
}

type ModifyAgentVersionRequest struct {
	*tchttp.BaseRequest
	
	// Agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 版本 ID（仅 default 或 test 版本可原地更新，prod 拒绝）
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Manifest v2.0 原文（可选；Manifest / Model / Description / SandboxTemplateId / ConnectorSet 五个可选字段至少提供一个）
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// 模型标识（可选）
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 版本变更说明（可选）
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 沙箱模板 ID。可选，patch 语义：null 不修改；空串解绑（恢复系统默认模板）；非空时模板须属于当前企业且可用（未删除、状态正常）。
	SandboxTemplateId *string `json:"SandboxTemplateId,omitnil,omitempty" name:"SandboxTemplateId"`

	// 该版本最终绑定的连接器集合（全量覆盖语义）：缺省 = 本次不改动连接器绑定；空数组 = 解绑全部连接器；非空 = 物化为 manifest v2 mcp_servers 网关条目，manifest 中不在本集合内的连接器条目会被移除（解绑在服务端闭环，无需调用方改写 Manifest）
	ConnectorSet []*ConnectorRefInput `json:"ConnectorSet,omitnil,omitempty" name:"ConnectorSet"`
}

func (r *ModifyAgentVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "VersionId")
	delete(f, "Manifest")
	delete(f, "Model")
	delete(f, "Description")
	delete(f, "SandboxTemplateId")
	delete(f, "ConnectorSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAgentVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentVersionResponseParams struct {
	// 版本 ID
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 版本名称
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 版本类型：DEFAULT / TEST / PROD
	VersionType *string `json:"VersionType,omitnil,omitempty" name:"VersionType"`

	// 版本变更说明
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 模型标识
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// Manifest v2.0 精简 manifest 原文（JSON 字符串）
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// 版本状态：DRAFT / ENABLED / DISABLED
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 绑定的沙箱模板 ID；未绑定时为空，创建会话沙箱使用系统默认模板。
	SandboxTemplateId *string `json:"SandboxTemplateId,omitnil,omitempty" name:"SandboxTemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAgentVersionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAgentVersionResponseParams `json:"Response"`
}

func (r *ModifyAgentVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoutingItem struct {
	// 版本 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 权重，(0, 1] 之间的浮点百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *float64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type SessionItem struct {
	// 会话 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 会话名称（AI 生成标题或用户改名；缺失时为空）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionName *string `json:"SessionName,omitnil,omitempty" name:"SessionName"`

	// Agent 业务 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Agent 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// 会话使用的版本名称（与 VersionId 区分：此为版本名，非 ID；原 AgentVersion）
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 会话使用的 Agent 版本 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 会话状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建者 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 会话来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 创建时间（RFC3339）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间（RFC3339）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`
}

type SkillCounts struct {
	// 内置技能数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Builtin *int64 `json:"Builtin,omitnil,omitempty" name:"Builtin"`

	// 自建技能数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Custom *int64 `json:"Custom,omitnil,omitempty" name:"Custom"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type SkillItem struct {
	// 技能来源：BUILTIN（内置）/ CUSTOM（自建）/ AUTHORIZED（企业授权）
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>slug（仅 custom 返回）</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>展示名</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>图标 URL</p>
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// <p>是否启用</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>下载 URL</p>
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// <p>技能标识</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>当前生效版本号</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillVersion *string `json:"SkillVersion,omitnil,omitempty" name:"SkillVersion"`

	// 创建时间，RFC3339 UTC 格式（如 2026-08-11T09:23:10Z）
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，RFC3339 UTC 格式（如 2026-09-15T06:51:26Z）
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type TokenUsage struct {
	// <p>输入Token</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputTokens *int64 `json:"InputTokens,omitnil,omitempty" name:"InputTokens"`

	// <p>输出Token</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputTokens *int64 `json:"OutputTokens,omitnil,omitempty" name:"OutputTokens"`

	// <p>总Token</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`

	// <p>统计口径</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`
}

// Predefined struct for user
type UnbindExternalAgentRequestParams struct {
	// TMA managed agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 已绑定的外部 A2A agent ID
	A2AAgentId *string `json:"A2AAgentId,omitnil,omitempty" name:"A2AAgentId"`

	// 绑定记录 ID（自增 ID 字符串）
	BindingId *string `json:"BindingId,omitnil,omitempty" name:"BindingId"`

	// 版本 ID
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type UnbindExternalAgentRequest struct {
	*tchttp.BaseRequest
	
	// TMA managed agent 业务 ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 已绑定的外部 A2A agent ID
	A2AAgentId *string `json:"A2AAgentId,omitnil,omitempty" name:"A2AAgentId"`

	// 绑定记录 ID（自增 ID 字符串）
	BindingId *string `json:"BindingId,omitnil,omitempty" name:"BindingId"`

	// 版本 ID
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *UnbindExternalAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindExternalAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "A2AAgentId")
	delete(f, "BindingId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindExternalAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindExternalAgentResponseParams struct {
	// 操作结果状态（大写枚举）：BOUND=已绑定 / UNBOUND=已解绑
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindExternalAgentResponse struct {
	*tchttp.BaseResponse
	Response *UnbindExternalAgentResponseParams `json:"Response"`
}

func (r *UnbindExternalAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindExternalAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}