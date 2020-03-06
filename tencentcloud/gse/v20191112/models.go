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

package v20191112

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateGameServerSessionRequest struct {
	*tchttp.BaseRequest

	// 最大玩家数量
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 创建者ID
	CreatorId *string `json:"CreatorId,omitempty" name:"CreatorId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏属性
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties" list`

	// 游戏服务器会话属性详情
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 游戏服务器会话自定义ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 幂等token
	IdempotencyToken *string `json:"IdempotencyToken,omitempty" name:"IdempotencyToken"`

	// 游戏服务器会话名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGameServerSessionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话
	// 注意：此字段可能返回 null，表示取不到有效值。
		GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGameServerSessionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// 服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScalingPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScalingPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetCapacityRequest struct {
	*tchttp.BaseRequest

	// 服务部署 Id列表
	FleetIds []*string `json:"FleetIds,omitempty" name:"FleetIds" list`

	// 结果返回最大数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFleetCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetCapacityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetCapacityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务部署容量配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetCapacity []*FleetCapacity `json:"FleetCapacity,omitempty" name:"FleetCapacity" list`

		// 结果返回最大数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFleetCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetCapacityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionDetailsRequest struct {
	*tchttp.BaseRequest

	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏服务器会话ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 游戏服务器会话状态(ACTIVE,ACTIVATING,TERMINATED,TERMINATING,ERROR)
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`
}

func (r *DescribeGameServerSessionDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionDetailsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		GameServerSessionDetails []*GameServerSessionDetail `json:"GameServerSessionDetails,omitempty" name:"GameServerSessionDetails" list`

		// 页偏移，用于查询下一页
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionDetailsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话放置的唯一标识符
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`
}

func (r *DescribeGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionPlacementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话放置
		GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionPlacementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionsRequest struct {
	*tchttp.BaseRequest

	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏服务器会话ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 游戏服务器会话状态
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`
}

func (r *DescribeGameServerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		GameServerSessions []*GameServerSession `json:"GameServerSessions,omitempty" name:"GameServerSessions" list`

		// 页便宜，用于查询下一页
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePlayerSessionsRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 玩家ID
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家会话ID
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`

	// 玩家会话状态（RESERVED,ACTIVE,COMPLETED,TIMEDOUT）
	PlayerSessionStatusFilter *string `json:"PlayerSessionStatusFilter,omitempty" name:"PlayerSessionStatusFilter"`
}

func (r *DescribePlayerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePlayerSessionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePlayerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 玩家会话列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		PlayerSessions []*PlayerSession `json:"PlayerSessions,omitempty" name:"PlayerSessions" list`

		// 页偏移
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePlayerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePlayerSessionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingPoliciesRequest struct {
	*tchttp.BaseRequest

	// 服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 状态过滤条件
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`

	// 结果返回最大数量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果偏移
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeScalingPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScalingPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 动态扩缩容配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		ScalingPolicies []*ScalingPolicy `json:"ScalingPolicies,omitempty" name:"ScalingPolicies" list`

		// 返回总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScalingPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScalingPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DesiredPlayerSession struct {

	// 与玩家会话关联的唯一玩家标识
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 开发人员定义的玩家数据
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`
}

type FleetCapacity struct {

	// 服务部署 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 服务器实例统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCounts *InstanceCounts `json:"InstanceCounts,omitempty" name:"InstanceCounts"`
}

type GameProperty struct {

	// 属性名称
	Key *string `json:"Key,omitempty" name:"Key"`

	// 属性值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type GameServerSession struct {

	// 游戏服务器会话创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 创建者ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorId *string `json:"CreatorId,omitempty" name:"CreatorId"`

	// 当前玩家数量
	CurrentPlayerSessionCount *uint64 `json:"CurrentPlayerSessionCount,omitempty" name:"CurrentPlayerSessionCount"`

	// CVM的DNS标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties" list`

	// 游戏服务器会话属性详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 游戏服务器会话ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// CVM IP地址
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 对战进程详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchmakerData *string `json:"MatchmakerData,omitempty" name:"MatchmakerData"`

	// 最大玩家数量
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 游戏服务器会话名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 玩家会话创建策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerSessionCreationPolicy *string `json:"PlayerSessionCreationPolicy,omitempty" name:"PlayerSessionCreationPolicy"`

	// 端口号
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 游戏服务器会话状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 游戏服务器会话状态附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusReason *string `json:"StatusReason,omitempty" name:"StatusReason"`

	// 终止的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`

	// 实例类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type GameServerSessionDetail struct {

	// 游戏服务器会话
	GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

	// 保护策略，可选（NoProtection,FullProtection）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectionPolicy *string `json:"ProtectionPolicy,omitempty" name:"ProtectionPolicy"`
}

type GameServerSessionPlacement struct {

	// 部署Id
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`

	// 服务部署组名称
	GameServerSessionQueueName *string `json:"GameServerSessionQueueName,omitempty" name:"GameServerSessionQueueName"`

	// 玩家延迟
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerLatencies []*PlayerLatency `json:"PlayerLatencies,omitempty" name:"PlayerLatencies" list`

	// 服务部署状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 分配给正在运行游戏会话的实例的DNS标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// 游戏会话Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 游戏会话名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionName *string `json:"GameServerSessionName,omitempty" name:"GameServerSessionName"`

	// 服务部署区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionRegion *string `json:"GameServerSessionRegion,omitempty" name:"GameServerSessionRegion"`

	// 游戏属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties" list`

	// 最大玩家数量
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 游戏会话数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 运行游戏会话的实例的IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 运行游戏会话的实例的端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 游戏匹配数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchmakerData *string `json:"MatchmakerData,omitempty" name:"MatchmakerData"`

	// 部署的玩家游戏数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlacedPlayerSessions []*PlacedPlayerSession `json:"PlacedPlayerSessions,omitempty" name:"PlacedPlayerSessions" list`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type GetGameServerSessionLogUrlRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`
}

func (r *GetGameServerSessionLogUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGameServerSessionLogUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGameServerSessionLogUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志下载URL
	// 注意：此字段可能返回 null，表示取不到有效值。
		PreSignedUrl *string `json:"PreSignedUrl,omitempty" name:"PreSignedUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGameServerSessionLogUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGameServerSessionLogUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceCounts struct {

	// 活跃的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Active *uint64 `json:"Active,omitempty" name:"Active"`

	// 期望的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desired *uint64 `json:"Desired,omitempty" name:"Desired"`

	// 空闲的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Idle *uint64 `json:"Idle,omitempty" name:"Idle"`

	// 服务器实例数最大限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxiNum *uint64 `json:"MaxiNum,omitempty" name:"MaxiNum"`

	// 服务器实例数最小限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	MiniNum *uint64 `json:"MiniNum,omitempty" name:"MiniNum"`

	// 已开始创建，但未激活的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pending *uint64 `json:"Pending,omitempty" name:"Pending"`

	// 结束中的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Terminating *uint64 `json:"Terminating,omitempty" name:"Terminating"`
}

type JoinGameServerSessionRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 玩家ID
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家自定义信息
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`
}

func (r *JoinGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *JoinGameServerSessionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type JoinGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 玩家会话
	// 注意：此字段可能返回 null，表示取不到有效值。
		PlayerSession *PlayerSession `json:"PlayerSession,omitempty" name:"PlayerSession"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *JoinGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *JoinGameServerSessionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PlacedPlayerSession struct {

	// 玩家Id
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家会话Id
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`
}

type PlayerLatency struct {

	// 玩家Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 延迟对应的区域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionIdentifier *string `json:"RegionIdentifier,omitempty" name:"RegionIdentifier"`

	// 毫秒级延迟
	LatencyInMilliseconds *uint64 `json:"LatencyInMilliseconds,omitempty" name:"LatencyInMilliseconds"`
}

type PlayerSession struct {

	// 玩家会话创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 游戏服务器会话运行的DNS标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏服务器会话ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 游戏服务器会话运行的CVM地址
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 玩家相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`

	// 玩家ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家会话ID
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`

	// 端口号
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 玩家会话的状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 玩家会话终止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`
}

type PutScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// 服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 基于规则的扩缩容配置
	TargetConfiguration *TargetConfiguration `json:"TargetConfiguration,omitempty" name:"TargetConfiguration"`
}

func (r *PutScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutScalingPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PutScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		Name *string `json:"Name,omitempty" name:"Name"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutScalingPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ScalingPolicy struct {

	// 服务部署ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalingAdjustment *string `json:"ScalingAdjustment,omitempty" name:"ScalingAdjustment"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalingAdjustmentType *string `json:"ScalingAdjustmentType,omitempty" name:"ScalingAdjustmentType"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" name:"ComparisonOperator"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Threshold *string `json:"Threshold,omitempty" name:"Threshold"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	EvaluationPeriods *string `json:"EvaluationPeriods,omitempty" name:"EvaluationPeriods"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 基于规则的配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetConfiguration *TargetConfiguration `json:"TargetConfiguration,omitempty" name:"TargetConfiguration"`
}

type StartGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest

	// 开始部署游戏服务器会话的唯一标识符
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`

	// 游戏服务器会话队列名称
	GameServerSessionQueueName *string `json:"GameServerSessionQueueName,omitempty" name:"GameServerSessionQueueName"`

	// 游戏服务器允许同时连接到游戏会话的最大玩家数量
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 玩家游戏会话信息
	DesiredPlayerSessions []*DesiredPlayerSession `json:"DesiredPlayerSessions,omitempty" name:"DesiredPlayerSessions" list`

	// 玩家游戏会话属性
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties" list`

	// 游戏服务器会话数据
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 游戏服务器会话名称
	GameServerSessionName *string `json:"GameServerSessionName,omitempty" name:"GameServerSessionName"`

	// 玩家延迟
	PlayerLatencies []*PlayerLatency `json:"PlayerLatencies,omitempty" name:"PlayerLatencies" list`
}

func (r *StartGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartGameServerSessionPlacementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话放置
		GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartGameServerSessionPlacementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话放置的唯一标识符
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`
}

func (r *StopGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopGameServerSessionPlacementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话放置
		GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopGameServerSessionPlacementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TargetConfiguration struct {

	// 预留存率
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetValue *uint64 `json:"TargetValue,omitempty" name:"TargetValue"`
}

type UpdateFleetCapacityRequest struct {
	*tchttp.BaseRequest

	// 服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 期望的服务器实例数
	DesiredInstances *uint64 `json:"DesiredInstances,omitempty" name:"DesiredInstances"`

	// 服务器实例数最小限制
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// 服务器实例数最大限制
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

func (r *UpdateFleetCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFleetCapacityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFleetCapacityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务部署ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFleetCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFleetCapacityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGameServerSessionRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话ID
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 最大玩家数量
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 游戏服务器会话名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 玩家会话创建策略（ACCEPT_ALL,DENY_ALL）
	PlayerSessionCreationPolicy *string `json:"PlayerSessionCreationPolicy,omitempty" name:"PlayerSessionCreationPolicy"`

	// 保护策略(NoProtection,TimeLimitProtection,FullProtection)
	ProtectionPolicy *string `json:"ProtectionPolicy,omitempty" name:"ProtectionPolicy"`
}

func (r *UpdateGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGameServerSessionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 更新后的游戏会话
		GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGameServerSessionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
