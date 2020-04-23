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

type Credentials struct {

	// ssh私钥
	Secret *string `json:"Secret,omitempty" name:"Secret"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`
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

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 结果返回最大数量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果偏移
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Instances []*Instance `json:"Instances,omitempty" name:"Instances" list`

		// 结果返回最大数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
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

type DesiredPlayerSession struct {

	// 与玩家会话关联的唯一玩家标识
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 开发人员定义的玩家数据
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`
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

	// 当前自定义数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentCustomCount *int64 `json:"CurrentCustomCount,omitempty" name:"CurrentCustomCount"`

	// 最大自定义数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxCustomCount *int64 `json:"MaxCustomCount,omitempty" name:"MaxCustomCount"`

	// 权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 会话可用性状态，是否被屏蔽
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailabilityStatus *string `json:"AvailabilityStatus,omitempty" name:"AvailabilityStatus"`
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

type GetInstanceAccessRequest struct {
	*tchttp.BaseRequest

	// 服务部署Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetInstanceAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetInstanceAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetInstanceAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例登录所需要的凭据
		InstanceAccess *InstanceAccess `json:"InstanceAccess,omitempty" name:"InstanceAccess"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInstanceAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetInstanceAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 服务部署ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// dns
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// 操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type InstanceAccess struct {

	// 访问实例所需要的凭据
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// 服务部署Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例公网IP
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 操作系统
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
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

type SearchGameServerSessionsRequest struct {
	*tchttp.BaseRequest

	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 搜索条件表达式，支持如下变量
	// gameServerSessionName 游戏会话名称 String
	// gameServerSessionId 游戏会话ID String
	// maximumSessions 最大的玩家会话数 Number
	// creationTimeMillis 创建时间，单位：毫秒 Number
	// playerSessionCount 当前玩家会话数 Number
	// hasAvailablePlayerSessions 是否有可用玩家数 String 取值true或false
	// gameServerSessionProperties 游戏会话属性 String
	// 
	// 表达式String类型 等于=，不等于<>判断
	// 表示Number类型支持 =,<>,>,>=,<,<=
	FilterExpression *string `json:"FilterExpression,omitempty" name:"FilterExpression"`

	// 排序条件关键字
	// 支持排序字段
	// gameServerSessionName 游戏会话名称 String
	// gameServerSessionId 游戏会话ID String
	// maximumSessions 最大的玩家会话数 Number
	// creationTimeMillis 创建时间，单位：毫秒 Number
	// playerSessionCount 当前玩家会话数 Number
	SortExpression *string `json:"SortExpression,omitempty" name:"SortExpression"`
}

func (r *SearchGameServerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchGameServerSessionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchGameServerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		GameServerSessions []*GameServerSession `json:"GameServerSessions,omitempty" name:"GameServerSessions" list`

		// 页偏移，用于查询下一页
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchGameServerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchGameServerSessionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
