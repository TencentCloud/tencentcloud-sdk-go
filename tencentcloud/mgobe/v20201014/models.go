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

package v20201014

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ChangeRoomPlayerProfileRequestParams struct {
	// 游戏资源Id。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 发起修改的玩家Id。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 需要修改的玩家自定义属性。
	CustomProfile *string `json:"CustomProfile,omitempty" name:"CustomProfile"`
}

type ChangeRoomPlayerProfileRequest struct {
	*tchttp.BaseRequest
	
	// 游戏资源Id。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 发起修改的玩家Id。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 需要修改的玩家自定义属性。
	CustomProfile *string `json:"CustomProfile,omitempty" name:"CustomProfile"`
}

func (r *ChangeRoomPlayerProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeRoomPlayerProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameId")
	delete(f, "PlayerId")
	delete(f, "CustomProfile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeRoomPlayerProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeRoomPlayerProfileResponseParams struct {
	// 房间信息。
	Room *Room `json:"Room,omitempty" name:"Room"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChangeRoomPlayerProfileResponse struct {
	*tchttp.BaseResponse
	Response *ChangeRoomPlayerProfileResponseParams `json:"Response"`
}

func (r *ChangeRoomPlayerProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeRoomPlayerProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeRoomPlayerStatusRequestParams struct {
	// 游戏资源Id。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 玩家自定义状态。
	CustomStatus *uint64 `json:"CustomStatus,omitempty" name:"CustomStatus"`

	// 玩家id。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`
}

type ChangeRoomPlayerStatusRequest struct {
	*tchttp.BaseRequest
	
	// 游戏资源Id。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 玩家自定义状态。
	CustomStatus *uint64 `json:"CustomStatus,omitempty" name:"CustomStatus"`

	// 玩家id。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`
}

func (r *ChangeRoomPlayerStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeRoomPlayerStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameId")
	delete(f, "CustomStatus")
	delete(f, "PlayerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeRoomPlayerStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeRoomPlayerStatusResponseParams struct {
	// 房间信息
	Room *Room `json:"Room,omitempty" name:"Room"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChangeRoomPlayerStatusResponse struct {
	*tchttp.BaseResponse
	Response *ChangeRoomPlayerStatusResponseParams `json:"Response"`
}

func (r *ChangeRoomPlayerStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeRoomPlayerStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlayerRequestParams struct {
	// 游戏资源Id。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 玩家OpenId。
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 玩家PlayerId，由后台分配，当OpenId不传的时候，PlayerId必传，传入PlayerId可以查询当前PlayerId的玩家信息，当OpenId传入的时候，PlayerId可不传，按照OpenId查询玩家信息。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`
}

type DescribePlayerRequest struct {
	*tchttp.BaseRequest
	
	// 游戏资源Id。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 玩家OpenId。
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 玩家PlayerId，由后台分配，当OpenId不传的时候，PlayerId必传，传入PlayerId可以查询当前PlayerId的玩家信息，当OpenId传入的时候，PlayerId可不传，按照OpenId查询玩家信息。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`
}

func (r *DescribePlayerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlayerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameId")
	delete(f, "OpenId")
	delete(f, "PlayerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePlayerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlayerResponseParams struct {
	// 玩家信息。
	Player *Player `json:"Player,omitempty" name:"Player"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePlayerResponse struct {
	*tchttp.BaseResponse
	Response *DescribePlayerResponseParams `json:"Response"`
}

func (r *DescribePlayerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlayerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomRequestParams struct {
	// 游戏资源Id。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 玩家Id。当房间Id不传的时候，玩家Id必传，传入玩家Id可以查询当前玩家所在的房间信息，当房间Id传入的时候，优先按照房间Id查询房间信息。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 房间Id。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

type DescribeRoomRequest struct {
	*tchttp.BaseRequest
	
	// 游戏资源Id。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 玩家Id。当房间Id不传的时候，玩家Id必传，传入玩家Id可以查询当前玩家所在的房间信息，当房间Id传入的时候，优先按照房间Id查询房间信息。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 房间Id。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DescribeRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameId")
	delete(f, "PlayerId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomResponseParams struct {
	// 房间信息。
	Room *Room `json:"Room,omitempty" name:"Room"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRoomResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoomResponseParams `json:"Response"`
}

func (r *DescribeRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismissRoomRequestParams struct {
	// 表示游戏资源唯一 ID, 由后台自动分配, 无法修改。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 表示游戏房间唯一ID。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

type DismissRoomRequest struct {
	*tchttp.BaseRequest
	
	// 表示游戏资源唯一 ID, 由后台自动分配, 无法修改。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 表示游戏房间唯一ID。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DismissRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DismissRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismissRoomResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DismissRoomResponse struct {
	*tchttp.BaseResponse
	Response *DismissRoomResponseParams `json:"Response"`
}

func (r *DismissRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoomRequestParams struct {
	// 游戏资源Id。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 发起者的PlayerId。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 需要修改的房间选项，0表示房间名称，1表示房主，2表示是否允许观战，3表示是否支持邀请码/密码，4表示是否私有，5表示是否自定义房间属性，6表示是否禁止加人。
	ChangeRoomOptionList []*int64 `json:"ChangeRoomOptionList,omitempty" name:"ChangeRoomOptionList"`

	// 房间名称。
	RoomName *string `json:"RoomName,omitempty" name:"RoomName"`

	// 变更房主。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 是否支持观战。
	IsViewed *bool `json:"IsViewed,omitempty" name:"IsViewed"`

	// 是否支持邀请码/密码。
	IsInvited *bool `json:"IsInvited,omitempty" name:"IsInvited"`

	// 是否私有。
	IsPrivate *bool `json:"IsPrivate,omitempty" name:"IsPrivate"`

	// 自定义房间属性。
	CustomProperties *string `json:"CustomProperties,omitempty" name:"CustomProperties"`

	// 房间是否禁止加人。
	IsForbidJoin *bool `json:"IsForbidJoin,omitempty" name:"IsForbidJoin"`
}

type ModifyRoomRequest struct {
	*tchttp.BaseRequest
	
	// 游戏资源Id。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 发起者的PlayerId。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 需要修改的房间选项，0表示房间名称，1表示房主，2表示是否允许观战，3表示是否支持邀请码/密码，4表示是否私有，5表示是否自定义房间属性，6表示是否禁止加人。
	ChangeRoomOptionList []*int64 `json:"ChangeRoomOptionList,omitempty" name:"ChangeRoomOptionList"`

	// 房间名称。
	RoomName *string `json:"RoomName,omitempty" name:"RoomName"`

	// 变更房主。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 是否支持观战。
	IsViewed *bool `json:"IsViewed,omitempty" name:"IsViewed"`

	// 是否支持邀请码/密码。
	IsInvited *bool `json:"IsInvited,omitempty" name:"IsInvited"`

	// 是否私有。
	IsPrivate *bool `json:"IsPrivate,omitempty" name:"IsPrivate"`

	// 自定义房间属性。
	CustomProperties *string `json:"CustomProperties,omitempty" name:"CustomProperties"`

	// 房间是否禁止加人。
	IsForbidJoin *bool `json:"IsForbidJoin,omitempty" name:"IsForbidJoin"`
}

func (r *ModifyRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameId")
	delete(f, "RoomId")
	delete(f, "PlayerId")
	delete(f, "ChangeRoomOptionList")
	delete(f, "RoomName")
	delete(f, "Owner")
	delete(f, "IsViewed")
	delete(f, "IsInvited")
	delete(f, "IsPrivate")
	delete(f, "CustomProperties")
	delete(f, "IsForbidJoin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoomResponseParams struct {
	// 房间信息
	Room *Room `json:"Room,omitempty" name:"Room"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRoomResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoomResponseParams `json:"Response"`
}

func (r *ModifyRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Player struct {
	// 玩家 OpenId。最长不超过64个字符。
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 玩家昵称。最长不超过32个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 队伍 ID。最长不超过16个字符。
	TeamId *string `json:"TeamId,omitempty" name:"TeamId"`

	// 是否为机器人。
	IsRobot *bool `json:"IsRobot,omitempty" name:"IsRobot"`

	// 玩家 PlayerId。出参使用，由后端返回。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 自定义玩家状态。非负数，最大不超过4294967295。默认为0。
	CustomPlayerStatus *uint64 `json:"CustomPlayerStatus,omitempty" name:"CustomPlayerStatus"`

	// 自定义玩家属性。最长不超过256个字符。默认为空字符串。
	CustomProfile *string `json:"CustomProfile,omitempty" name:"CustomProfile"`
}

// Predefined struct for user
type RemoveRoomPlayerRequestParams struct {
	// 游戏资源Id。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 被踢出房间的玩家Id。
	RemovePlayerId *string `json:"RemovePlayerId,omitempty" name:"RemovePlayerId"`
}

type RemoveRoomPlayerRequest struct {
	*tchttp.BaseRequest
	
	// 游戏资源Id。
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 被踢出房间的玩家Id。
	RemovePlayerId *string `json:"RemovePlayerId,omitempty" name:"RemovePlayerId"`
}

func (r *RemoveRoomPlayerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveRoomPlayerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameId")
	delete(f, "RemovePlayerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveRoomPlayerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveRoomPlayerResponseParams struct {
	// 房间信息
	Room *Room `json:"Room,omitempty" name:"Room"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveRoomPlayerResponse struct {
	*tchttp.BaseResponse
	Response *RemoveRoomPlayerResponseParams `json:"Response"`
}

func (r *RemoveRoomPlayerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveRoomPlayerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Room struct {
	// 表示房间名称。最长不超过32个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 表示房间最大玩家数量。最大不超过100人。
	MaxPlayers *int64 `json:"MaxPlayers,omitempty" name:"MaxPlayers"`

	// 表示房主OpenId。最长不超过16个字符。
	OwnerOpenId *string `json:"OwnerOpenId,omitempty" name:"OwnerOpenId"`

	// 表示是否私有，私有指的是不允许其他玩家通过匹配加入房间。
	IsPrivate *bool `json:"IsPrivate,omitempty" name:"IsPrivate"`

	// 表示玩家详情列表。
	Players []*Player `json:"Players,omitempty" name:"Players"`

	// 表示团队属性列表。
	Teams []*Team `json:"Teams,omitempty" name:"Teams"`

	// 表示房间 ID。出参用，由后端返回。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 表示房间类型。最长不超过32个字符。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 表示创建方式：0.单人主动发起创建房间请求；1.多人在线匹配请求分配房间；2. 直接创建满员房间。调用云API的创房请求默认为3，目前通过云API调用只支持第3种方式。
	CreateType *int64 `json:"CreateType,omitempty" name:"CreateType"`

	// 表示自定义房间属性，不传为空字符串。最长不超过1024个字符。
	CustomProperties *string `json:"CustomProperties,omitempty" name:"CustomProperties"`

	// 表示房间帧同步状态。0表示未开始帧同步，1表示已开始帧同步，用于出参。
	FrameSyncState *int64 `json:"FrameSyncState,omitempty" name:"FrameSyncState"`

	// 表示帧率。由控制台设置，用于出参。
	FrameRate *int64 `json:"FrameRate,omitempty" name:"FrameRate"`

	// 表示路由ID。用于出参。
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`

	// 表示房间创建的时间戳（单位：秒）。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 表示开始帧同步时的时间戳（单位：秒）,未开始帧同步时返回为0。
	StartGameTime *int64 `json:"StartGameTime,omitempty" name:"StartGameTime"`

	// 表示是否禁止加入房间。出参使用，默认为False，通过SDK的ChangeRoom接口可以修改。
	IsForbidJoin *bool `json:"IsForbidJoin,omitempty" name:"IsForbidJoin"`

	// 表示房主PlayerId。
	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

type Team struct {
	// 队伍ID。最长不超过16个字符。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 队伍名称。最长不超过32个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 队伍最小人数。最大不超过100人。
	MinPlayers *int64 `json:"MinPlayers,omitempty" name:"MinPlayers"`

	// 队伍最大人数。最大不超过100人。
	MaxPlayers *int64 `json:"MaxPlayers,omitempty" name:"MaxPlayers"`
}