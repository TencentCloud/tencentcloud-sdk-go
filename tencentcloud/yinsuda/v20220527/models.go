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

package v20220527

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AMEMusicBaseInfo struct {
	// 歌曲 Id。
	MusicId *string `json:"MusicId,omitnil,omitempty" name:"MusicId"`

	// 歌曲名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 歌手列表。
	SingerSet []*string `json:"SingerSet,omitnil,omitempty" name:"SingerSet"`
}

// Predefined struct for user
type ApplyChorusRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 歌曲 Id。
	MusicId *string `json:"MusicId,omitnil,omitempty" name:"MusicId"`

	// 最大合唱人数，默认值为 8，最大值为 20。
	MaxChorusNum *uint64 `json:"MaxChorusNum,omitnil,omitempty" name:"MaxChorusNum"`

	// 合唱用户标识列表。
	ChorusUserIds []*string `json:"ChorusUserIds,omitnil,omitempty" name:"ChorusUserIds"`
}

type ApplyChorusRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 歌曲 Id。
	MusicId *string `json:"MusicId,omitnil,omitempty" name:"MusicId"`

	// 最大合唱人数，默认值为 8，最大值为 20。
	MaxChorusNum *uint64 `json:"MaxChorusNum,omitnil,omitempty" name:"MaxChorusNum"`

	// 合唱用户标识列表。
	ChorusUserIds []*string `json:"ChorusUserIds,omitnil,omitempty" name:"ChorusUserIds"`
}

func (r *ApplyChorusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyChorusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "RoomId")
	delete(f, "MusicId")
	delete(f, "MaxChorusNum")
	delete(f, "ChorusUserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyChorusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyChorusResponseParams struct {
	// 合唱 Token。
	ChorusToken *string `json:"ChorusToken,omitnil,omitempty" name:"ChorusToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyChorusResponse struct {
	*tchttp.BaseResponse
	Response *ApplyChorusResponseParams `json:"Response"`
}

func (r *ApplyChorusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyChorusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDescribeKTVMusicDetailsRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 歌曲 Id 列表。
	MusicIds []*string `json:"MusicIds,omitnil,omitempty" name:"MusicIds"`

	// 播放场景。默认为Chat
	// <li>Live：直播</li><li>Chat：语聊</li>
	PlayScene *string `json:"PlayScene,omitnil,omitempty" name:"PlayScene"`

	// 玩家用户标识
	GuestUserId *string `json:"GuestUserId,omitnil,omitempty" name:"GuestUserId"`

	// 房间Id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type BatchDescribeKTVMusicDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 歌曲 Id 列表。
	MusicIds []*string `json:"MusicIds,omitnil,omitempty" name:"MusicIds"`

	// 播放场景。默认为Chat
	// <li>Live：直播</li><li>Chat：语聊</li>
	PlayScene *string `json:"PlayScene,omitnil,omitempty" name:"PlayScene"`

	// 玩家用户标识
	GuestUserId *string `json:"GuestUserId,omitnil,omitempty" name:"GuestUserId"`

	// 房间Id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *BatchDescribeKTVMusicDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDescribeKTVMusicDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "MusicIds")
	delete(f, "PlayScene")
	delete(f, "GuestUserId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDescribeKTVMusicDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDescribeKTVMusicDetailsResponseParams struct {
	// 歌曲详细信息列表。
	KTVMusicDetailInfoSet []*KTVMusicDetailInfo `json:"KTVMusicDetailInfoSet,omitnil,omitempty" name:"KTVMusicDetailInfoSet"`

	// 不存在歌曲Id列表。
	NotExistMusicIdSet []*string `json:"NotExistMusicIdSet,omitnil,omitempty" name:"NotExistMusicIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchDescribeKTVMusicDetailsResponse struct {
	*tchttp.BaseResponse
	Response *BatchDescribeKTVMusicDetailsResponseParams `json:"Response"`
}

func (r *BatchDescribeKTVMusicDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDescribeKTVMusicDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChorusClip struct {
	// 开始时间，单位：毫秒。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，单位：毫秒。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

// Predefined struct for user
type CreateKTVRobotRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// RTC厂商类型，取值有：
	// <li>TRTC</li>
	RTCSystem *string `json:"RTCSystem,omitnil,omitempty" name:"RTCSystem"`

	// 进房参数。
	JoinRoomInput *JoinRoomInput `json:"JoinRoomInput,omitnil,omitempty" name:"JoinRoomInput"`

	// 创建机器人时初始化参数。
	SyncRobotCommands []*SyncRobotCommand `json:"SyncRobotCommands,omitnil,omitempty" name:"SyncRobotCommands"`
}

type CreateKTVRobotRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// RTC厂商类型，取值有：
	// <li>TRTC</li>
	RTCSystem *string `json:"RTCSystem,omitnil,omitempty" name:"RTCSystem"`

	// 进房参数。
	JoinRoomInput *JoinRoomInput `json:"JoinRoomInput,omitnil,omitempty" name:"JoinRoomInput"`

	// 创建机器人时初始化参数。
	SyncRobotCommands []*SyncRobotCommand `json:"SyncRobotCommands,omitnil,omitempty" name:"SyncRobotCommands"`
}

func (r *CreateKTVRobotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKTVRobotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "RTCSystem")
	delete(f, "JoinRoomInput")
	delete(f, "SyncRobotCommands")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKTVRobotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKTVRobotResponseParams struct {
	// 机器人Id。
	RobotId *string `json:"RobotId,omitnil,omitempty" name:"RobotId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateKTVRobotResponse struct {
	*tchttp.BaseResponse
	Response *CreateKTVRobotResponseParams `json:"Response"`
}

func (r *CreateKTVRobotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKTVRobotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVMatchMusicsRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 匹配规则列表。
	Rules []*KTVMatchRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type DescribeKTVMatchMusicsRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 匹配规则列表。
	Rules []*KTVMatchRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *DescribeKTVMatchMusicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVMatchMusicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVMatchMusicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVMatchMusicsResponseParams struct {
	// 匹配到的歌曲列表。
	MatchMusicSet []*KTVMatchMusic `json:"MatchMusicSet,omitnil,omitempty" name:"MatchMusicSet"`

	// 未匹配的规则列表。
	NotMatchRuleSet []*KTVMatchRule `json:"NotMatchRuleSet,omitnil,omitempty" name:"NotMatchRuleSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKTVMatchMusicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVMatchMusicsResponseParams `json:"Response"`
}

func (r *DescribeKTVMatchMusicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVMatchMusicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVMusicAccompanySegmentUrlRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 歌曲 Id 。
	MusicId *string `json:"MusicId,omitnil,omitempty" name:"MusicId"`

	// 播放场景。默认为Chat
	// <li>Live：直播</li><li>Chat：语聊</li>
	PlayScene *string `json:"PlayScene,omitnil,omitempty" name:"PlayScene"`

	// 房间Id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeKTVMusicAccompanySegmentUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 歌曲 Id 。
	MusicId *string `json:"MusicId,omitnil,omitempty" name:"MusicId"`

	// 播放场景。默认为Chat
	// <li>Live：直播</li><li>Chat：语聊</li>
	PlayScene *string `json:"PlayScene,omitnil,omitempty" name:"PlayScene"`

	// 房间Id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DescribeKTVMusicAccompanySegmentUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVMusicAccompanySegmentUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "MusicId")
	delete(f, "PlayScene")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVMusicAccompanySegmentUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVMusicAccompanySegmentUrlResponseParams struct {
	// 歌曲状态。
	// 0：可用
	// 1：下线
	// 2：没权限
	// 3：没伴奏
	// 当返回2时，其他参数有可能全部为空
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 伴奏链接
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 伴奏类型，如mkv，mp3等
	ExtName *string `json:"ExtName,omitnil,omitempty" name:"ExtName"`

	// 高潮开始时间
	SegmentBegin *int64 `json:"SegmentBegin,omitnil,omitempty" name:"SegmentBegin"`

	// 高潮结束时间
	SegmentEnd *int64 `json:"SegmentEnd,omitnil,omitempty" name:"SegmentEnd"`

	// 链接文件大小 单位 字节
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 其它片段时间（可用于抢唱）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherSegments []*KTVOtherSegments `json:"OtherSegments,omitnil,omitempty" name:"OtherSegments"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKTVMusicAccompanySegmentUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVMusicAccompanySegmentUrlResponseParams `json:"Response"`
}

func (r *DescribeKTVMusicAccompanySegmentUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVMusicAccompanySegmentUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVMusicAccompanySegmentUrlVipRequestParams struct {
	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 歌曲 Id 
	MusicId *string `json:"MusicId,omitnil,omitempty" name:"MusicId"`
}

type DescribeKTVMusicAccompanySegmentUrlVipRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 歌曲 Id 
	MusicId *string `json:"MusicId,omitnil,omitempty" name:"MusicId"`
}

func (r *DescribeKTVMusicAccompanySegmentUrlVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVMusicAccompanySegmentUrlVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "MusicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVMusicAccompanySegmentUrlVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVMusicAccompanySegmentUrlVipResponseParams struct {
	// 0:成功获取 1:歌曲下架 2:无权限 3: 非包月会员 4:没有对应的链接
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 伴奏链接
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 伴奏类型，如mkv，mp3等
	ExtName *string `json:"ExtName,omitnil,omitempty" name:"ExtName"`

	// 高潮开始时间
	SegmentBegin *int64 `json:"SegmentBegin,omitnil,omitempty" name:"SegmentBegin"`

	// 高潮结束时间
	SegmentEnd *int64 `json:"SegmentEnd,omitnil,omitempty" name:"SegmentEnd"`

	// 链接文件大小 (单位:字节)
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKTVMusicAccompanySegmentUrlVipResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVMusicAccompanySegmentUrlVipResponseParams `json:"Response"`
}

func (r *DescribeKTVMusicAccompanySegmentUrlVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVMusicAccompanySegmentUrlVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVMusicsByTagRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 标签 Id。
	TagId *string `json:"TagId,omitnil,omitempty" name:"TagId"`

	// 滚动标记。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 返回条数限制，默认 20，最大 50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 权益过滤，取值有：
	// <li>Play：可播；</li>
	// <li>Sing：可唱。</li>
	RightFilters []*string `json:"RightFilters,omitnil,omitempty" name:"RightFilters"`

	// 物料过滤，取值有：
	// <li>Lyrics：含有歌词；</li>
	// <li>Midi：含有音高线。</li>
	MaterialFilters []*string `json:"MaterialFilters,omitnil,omitempty" name:"MaterialFilters"`
}

type DescribeKTVMusicsByTagRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 标签 Id。
	TagId *string `json:"TagId,omitnil,omitempty" name:"TagId"`

	// 滚动标记。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 返回条数限制，默认 20，最大 50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 权益过滤，取值有：
	// <li>Play：可播；</li>
	// <li>Sing：可唱。</li>
	RightFilters []*string `json:"RightFilters,omitnil,omitempty" name:"RightFilters"`

	// 物料过滤，取值有：
	// <li>Lyrics：含有歌词；</li>
	// <li>Midi：含有音高线。</li>
	MaterialFilters []*string `json:"MaterialFilters,omitnil,omitempty" name:"MaterialFilters"`
}

func (r *DescribeKTVMusicsByTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVMusicsByTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "TagId")
	delete(f, "ScrollToken")
	delete(f, "Limit")
	delete(f, "RightFilters")
	delete(f, "MaterialFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVMusicsByTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVMusicsByTagResponseParams struct {
	// 歌曲信息列表。
	KTVMusicInfoSet []*KTVMusicBaseInfo `json:"KTVMusicInfoSet,omitnil,omitempty" name:"KTVMusicInfoSet"`

	// 滚动标记，用于设置下次请求的 ScrollToken 参数。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKTVMusicsByTagResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVMusicsByTagResponseParams `json:"Response"`
}

func (r *DescribeKTVMusicsByTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVMusicsByTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVPlaylistDetailRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 歌单 Id。
	PlaylistId *string `json:"PlaylistId,omitnil,omitempty" name:"PlaylistId"`

	// 滚动标记。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 返回条数，默认：20，最大：50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 权益过滤，取值有：
	// <li>Play：可播；</li>
	// <li>Sing：可唱。</li>
	RightFilters []*string `json:"RightFilters,omitnil,omitempty" name:"RightFilters"`

	// 播放场景。默认为Chat
	// <li>Live：直播</li><li>Chat：语聊</li>
	PlayScene *string `json:"PlayScene,omitnil,omitempty" name:"PlayScene"`

	// 物料过滤，取值有：
	// <li>Lyrics：含有歌词；</li>
	// <li>Midi：含有音高线。</li>
	MaterialFilters []*string `json:"MaterialFilters,omitnil,omitempty" name:"MaterialFilters"`
}

type DescribeKTVPlaylistDetailRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 歌单 Id。
	PlaylistId *string `json:"PlaylistId,omitnil,omitempty" name:"PlaylistId"`

	// 滚动标记。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 返回条数，默认：20，最大：50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 权益过滤，取值有：
	// <li>Play：可播；</li>
	// <li>Sing：可唱。</li>
	RightFilters []*string `json:"RightFilters,omitnil,omitempty" name:"RightFilters"`

	// 播放场景。默认为Chat
	// <li>Live：直播</li><li>Chat：语聊</li>
	PlayScene *string `json:"PlayScene,omitnil,omitempty" name:"PlayScene"`

	// 物料过滤，取值有：
	// <li>Lyrics：含有歌词；</li>
	// <li>Midi：含有音高线。</li>
	MaterialFilters []*string `json:"MaterialFilters,omitnil,omitempty" name:"MaterialFilters"`
}

func (r *DescribeKTVPlaylistDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVPlaylistDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "PlaylistId")
	delete(f, "ScrollToken")
	delete(f, "Limit")
	delete(f, "RightFilters")
	delete(f, "PlayScene")
	delete(f, "MaterialFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVPlaylistDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVPlaylistDetailResponseParams struct {
	// 歌曲信息列表。
	KTVMusicInfoSet []*KTVMusicBaseInfo `json:"KTVMusicInfoSet,omitnil,omitempty" name:"KTVMusicInfoSet"`

	// 滚动标记，用于设置下次请求的 ScrollToken 参数。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKTVPlaylistDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVPlaylistDetailResponseParams `json:"Response"`
}

func (r *DescribeKTVPlaylistDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVPlaylistDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVPlaylistsRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 类型列表，取值有：
	// <li>OfficialRec：官方推荐；</li>
	// <li>Customize：自定义。</li>
	// 默认值为 OfficialRec。
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：20，最大值：50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeKTVPlaylistsRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 类型列表，取值有：
	// <li>OfficialRec：官方推荐；</li>
	// <li>Customize：自定义。</li>
	// 默认值为 OfficialRec。
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：20，最大值：50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeKTVPlaylistsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVPlaylistsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "Types")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVPlaylistsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVPlaylistsResponseParams struct {
	// 歌单基础信息。
	PlaylistBaseInfoSet []*KTVPlaylistBaseInfo `json:"PlaylistBaseInfoSet,omitnil,omitempty" name:"PlaylistBaseInfoSet"`

	// 歌单总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKTVPlaylistsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVPlaylistsResponseParams `json:"Response"`
}

func (r *DescribeKTVPlaylistsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVPlaylistsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVRobotsRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 机器人Id列表。
	RobotIds []*string `json:"RobotIds,omitnil,omitempty" name:"RobotIds"`

	// 机器人状态，取值有：
	// <li>Play：播放</li>
	// <li>Pause：暂停</li>
	// <li>Destroy：销毁</li>
	Statuses []*string `json:"Statuses,omitnil,omitempty" name:"Statuses"`

	// 匹配创建时间在此时间段内的机器人。
	// <li>包含所指定的头尾时间点。</li>
	CreateTime *TimeRange `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页返回的起始偏移量，默认值：10。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeKTVRobotsRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 机器人Id列表。
	RobotIds []*string `json:"RobotIds,omitnil,omitempty" name:"RobotIds"`

	// 机器人状态，取值有：
	// <li>Play：播放</li>
	// <li>Pause：暂停</li>
	// <li>Destroy：销毁</li>
	Statuses []*string `json:"Statuses,omitnil,omitempty" name:"Statuses"`

	// 匹配创建时间在此时间段内的机器人。
	// <li>包含所指定的头尾时间点。</li>
	CreateTime *TimeRange `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页返回的起始偏移量，默认值：10。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeKTVRobotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVRobotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "RobotIds")
	delete(f, "Statuses")
	delete(f, "CreateTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVRobotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVRobotsResponseParams struct {
	// 机器人总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 机器人信息集合。
	KTVRobotInfoSet []*KTVRobotInfo `json:"KTVRobotInfoSet,omitnil,omitempty" name:"KTVRobotInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKTVRobotsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVRobotsResponseParams `json:"Response"`
}

func (r *DescribeKTVRobotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVRobotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVSuggestionsRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 搜索词。
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`
}

type DescribeKTVSuggestionsRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 搜索词。
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`
}

func (r *DescribeKTVSuggestionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVSuggestionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "KeyWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVSuggestionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVSuggestionsResponseParams struct {
	// 联想词信息列表。
	KTVSuggestionInfoSet []*KTVSuggestionInfo `json:"KTVSuggestionInfoSet,omitnil,omitempty" name:"KTVSuggestionInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKTVSuggestionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVSuggestionsResponseParams `json:"Response"`
}

func (r *DescribeKTVSuggestionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVSuggestionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVTagsRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeKTVTagsRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeKTVTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVTagsResponseParams struct {
	// 标签分组列表。
	TagGroupInfoSet []*KTVTagGroupInfo `json:"TagGroupInfoSet,omitnil,omitempty" name:"TagGroupInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKTVTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVTagsResponseParams `json:"Response"`
}

func (r *DescribeKTVTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveVipTradeInfosRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 直播会员充值下单起始时间，格式为 ISO。默认为当前时间前一天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 直播会员充值下单截止时间，格式为 ISO。默认为当前时间。 EndTime不能小于StartTime
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 交易流水号集合，匹配集合指定所有流水号 。
	// <li>数组长度限制：10。</li>
	TradeSerialNos []*string `json:"TradeSerialNos,omitnil,omitempty" name:"TradeSerialNos"`

	// 用户标识集合，匹配集合指定所有用户标识 。
	// <li>数组长度限制：10。</li>
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：20，最大值：50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeLiveVipTradeInfosRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 直播会员充值下单起始时间，格式为 ISO。默认为当前时间前一天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 直播会员充值下单截止时间，格式为 ISO。默认为当前时间。 EndTime不能小于StartTime
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 交易流水号集合，匹配集合指定所有流水号 。
	// <li>数组长度限制：10。</li>
	TradeSerialNos []*string `json:"TradeSerialNos,omitnil,omitempty" name:"TradeSerialNos"`

	// 用户标识集合，匹配集合指定所有用户标识 。
	// <li>数组长度限制：10。</li>
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：20，最大值：50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeLiveVipTradeInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveVipTradeInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TradeSerialNos")
	delete(f, "UserIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveVipTradeInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveVipTradeInfosResponseParams struct {
	// 直播会员充值流水信息列表
	LiveVipTradeInfoSet []*LiveVipTradeInfo `json:"LiveVipTradeInfoSet,omitnil,omitempty" name:"LiveVipTradeInfoSet"`

	// 直播会员充值流水总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveVipTradeInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveVipTradeInfosResponseParams `json:"Response"`
}

func (r *DescribeLiveVipTradeInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveVipTradeInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInfoRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInfoResponseParams struct {
	// 用户信息。
	UserInfo *UserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserInfoResponseParams `json:"Response"`
}

func (r *DescribeUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVipUserInfoRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeVipUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeVipUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVipUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVipUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVipUserInfoResponseParams struct {
	// 是否是会员。（0:不是会员 1:是会员）
	IsVip *int64 `json:"IsVip,omitnil,omitempty" name:"IsVip"`

	// 主播id
	AnchorId *string `json:"AnchorId,omitnil,omitempty" name:"AnchorId"`

	// 房间id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 会员过期时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 会员状态。（-1:未开通过；1:已开通，未过期；2:已开通，已过期）
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVipUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVipUserInfoResponseParams `json:"Response"`
}

func (r *DescribeVipUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVipUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyKTVRobotRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 机器人Id。
	RobotId *string `json:"RobotId,omitnil,omitempty" name:"RobotId"`
}

type DestroyKTVRobotRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 机器人Id。
	RobotId *string `json:"RobotId,omitnil,omitempty" name:"RobotId"`
}

func (r *DestroyKTVRobotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyKTVRobotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "RobotId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyKTVRobotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyKTVRobotResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyKTVRobotResponse struct {
	*tchttp.BaseResponse
	Response *DestroyKTVRobotResponseParams `json:"Response"`
}

func (r *DestroyKTVRobotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyKTVRobotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JoinRoomInput struct {
	// TRTC进房参数
	TRTCJoinRoomInput *TRTCJoinRoomInput `json:"TRTCJoinRoomInput,omitnil,omitempty" name:"TRTCJoinRoomInput"`
}

type KTVBPMInfo struct {
	// 节拍类型，取值有：
	// <li>Slow：慢；</li>
	// <li>Middle：中等；</li>
	// <li>Fast：快；</li>
	// <li>Unknown：未知。</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// BPM 值。
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type KTVMatchMusic struct {
	// 匹配到的歌曲基础信息。
	KTVMusicBaseInfo *KTVMusicBaseInfo `json:"KTVMusicBaseInfo,omitnil,omitempty" name:"KTVMusicBaseInfo"`

	// 命中规则。
	MatchRule *KTVMatchRule `json:"MatchRule,omitnil,omitempty" name:"MatchRule"`

	// AME 歌曲基础信息，仅在使用音速达歌曲 Id 匹配 AME 曲库时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AMEMusicBaseInfo *AMEMusicBaseInfo `json:"AMEMusicBaseInfo,omitnil,omitempty" name:"AMEMusicBaseInfo"`
}

type KTVMatchRule struct {
	// AME 曲库 Id。
	AMEMusicId *string `json:"AMEMusicId,omitnil,omitempty" name:"AMEMusicId"`

	// 歌曲匹配信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MusicInfo *KTVMatchRuleMusicInfo `json:"MusicInfo,omitnil,omitempty" name:"MusicInfo"`

	// 音速达歌曲 Id，用于匹配 AME 曲库歌曲。
	MusicIdToMatchAME *string `json:"MusicIdToMatchAME,omitnil,omitempty" name:"MusicIdToMatchAME"`
}

type KTVMatchRuleMusicInfo struct {
	// 歌曲名称。
	MusicName *string `json:"MusicName,omitnil,omitempty" name:"MusicName"`

	// 歌手列表。
	SingerSet []*string `json:"SingerSet,omitnil,omitempty" name:"SingerSet"`
}

type KTVMusicBaseInfo struct {
	// 歌曲Id。
	MusicId *string `json:"MusicId,omitnil,omitempty" name:"MusicId"`

	// 歌曲名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 歌手名称。
	SingerSet []*string `json:"SingerSet,omitnil,omitempty" name:"SingerSet"`

	// 播放时长。
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 歌手图片链接。
	SingerImageUrl *string `json:"SingerImageUrl,omitnil,omitempty" name:"SingerImageUrl"`

	// 专辑信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbumInfo *MusicAlbumInfo `json:"AlbumInfo,omitnil,omitempty" name:"AlbumInfo"`

	// 权益列表，取值有：
	// <li>Play：可播；</li>
	// <li>Sing：可唱。</li>
	RightSet []*string `json:"RightSet,omitnil,omitempty" name:"RightSet"`

	// 推荐类型，取值有：
	// <li>Featured：精选；</li>
	// <li>Other：其他。</li>
	RecommendType *string `json:"RecommendType,omitnil,omitempty" name:"RecommendType"`
}

type KTVMusicDetailInfo struct {
	// 歌曲基础信息。
	KTVMusicBaseInfo *KTVMusicBaseInfo `json:"KTVMusicBaseInfo,omitnil,omitempty" name:"KTVMusicBaseInfo"`

	// 播放凭证。
	PlayToken *string `json:"PlayToken,omitnil,omitempty" name:"PlayToken"`

	// 歌词下载链接。
	LyricsUrl *string `json:"LyricsUrl,omitnil,omitempty" name:"LyricsUrl"`

	// 音高数据下载链接。
	MidiUrl *string `json:"MidiUrl,omitnil,omitempty" name:"MidiUrl"`

	// 副歌片段信息。
	ChorusClipSet []*ChorusClip `json:"ChorusClipSet,omitnil,omitempty" name:"ChorusClipSet"`

	// 前奏间隔。
	PreludeInterval *int64 `json:"PreludeInterval,omitnil,omitempty" name:"PreludeInterval"`

	// 歌曲流派列表。
	GenreSet []*string `json:"GenreSet,omitnil,omitempty" name:"GenreSet"`

	// 节拍信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BPMInfo *KTVBPMInfo `json:"BPMInfo,omitnil,omitempty" name:"BPMInfo"`
}

type KTVOtherSegments struct {
	// 片段开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentBegin *int64 `json:"SegmentBegin,omitnil,omitempty" name:"SegmentBegin"`

	// 片段结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentEnd *int64 `json:"SegmentEnd,omitnil,omitempty" name:"SegmentEnd"`
}

type KTVPlaylistBaseInfo struct {
	// 歌单Id。
	PlaylistId *string `json:"PlaylistId,omitnil,omitempty" name:"PlaylistId"`

	// 歌单标题。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

type KTVRobotInfo struct {
	// 机器人Id。
	RobotId *string `json:"RobotId,omitnil,omitempty" name:"RobotId"`

	// 状态，取值有：
	// <li>Play：播放</li>
	// <li>Pause：暂停</li>
	// <li>Destroy：销毁</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 播放列表。
	Playlists []*string `json:"Playlists,omitnil,omitempty" name:"Playlists"`

	// 当前歌单索引位置。
	CurIndex *int64 `json:"CurIndex,omitnil,omitempty" name:"CurIndex"`

	// 播放进度，单位：毫秒。
	Position *uint64 `json:"Position,omitnil,omitempty" name:"Position"`

	// 音频参数。
	SetAudioParamInput *SetAudioParamCommandInput `json:"SetAudioParamInput,omitnil,omitempty" name:"SetAudioParamInput"`

	// 进房信息。
	JoinRoomInput *JoinRoomInput `json:"JoinRoomInput,omitnil,omitempty" name:"JoinRoomInput"`

	// RTC厂商类型，取值有：
	// <li>TRTC</li>
	RTCSystem *string `json:"RTCSystem,omitnil,omitempty" name:"RTCSystem"`

	// 播放模式，PlayMode取值有：
	// <li>RepeatPlaylist：列表循环</li>
	// <li>Order：顺序播放</li>
	// <li>RepeatSingle：单曲循环</li>
	// <li>Shuffle：随机播放</li>
	SetPlayModeInput *SetPlayModeCommandInput `json:"SetPlayModeInput,omitnil,omitempty" name:"SetPlayModeInput"`
}

type KTVSuggestionInfo struct {
	// 联想词。
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`
}

type KTVTagGroupInfo struct {
	// 分组 Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分组名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签列表。
	TagInfoSet []*KTVTagInfo `json:"TagInfoSet,omitnil,omitempty" name:"TagInfoSet"`
}

type KTVTagInfo struct {
	// 标签 Id。
	TagId *string `json:"TagId,omitnil,omitempty" name:"TagId"`

	// 标签名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type LiveVipTradeInfo struct {
	// 交易流水号。
	TradeSerialNo *string `json:"TradeSerialNo,omitnil,omitempty" name:"TradeSerialNo"`

	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 房间标识。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 充值会员天数。
	// 取值有： 
	// <li>31</li> <li>93</li><li>186</li> <li>372</li>
	VipDays *int64 `json:"VipDays,omitnil,omitempty" name:"VipDays"`

	// 订单状态。 
	// 取值有： 
	// <li>Success：成功</li><li>Fail：失败</li><li>Processing：订单处理中</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type LiveVipUserInfo struct {
	// 房间标识。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 直播会员结束时间。
	LiveVipEndTime *string `json:"LiveVipEndTime,omitnil,omitempty" name:"LiveVipEndTime"`

	// 会员生效状态
	// <li>Valid：生效</li><li>Invalid：无效</li>
	LiveVipStatus *string `json:"LiveVipStatus,omitnil,omitempty" name:"LiveVipStatus"`
}

type MusicAlbumCoverInfo struct {
	// 尺寸规格，取值有：
	// <li>Mini：150 x 150 尺寸；</li>
	// <li>Small：240 x 240 尺寸；</li>
	// <li>Medium：480 x 480 尺寸。</li>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// 下载链接。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type MusicAlbumInfo struct {
	// 专辑名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 封面列表。
	CoverInfoSet []*MusicAlbumCoverInfo `json:"CoverInfoSet,omitnil,omitempty" name:"CoverInfoSet"`
}

type PlayCommandInput struct {
	// 歌曲位置索引。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}

// Predefined struct for user
type RechargeLiveVipRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 交易流水号，用于标记此次充值记录，多次充值记录传入相同的 TradeSerialNo 会判断为失败，可用于防止重提提交造成重复计费。
	TradeSerialNo *string `json:"TradeSerialNo,omitnil,omitempty" name:"TradeSerialNo"`

	// 房间标识。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 充值会员天数。
	// 取值有：
	// <li>31</li>
	// <li>93</li>
	// <li>186</li>
	// <li>372</li>
	VipDays *int64 `json:"VipDays,omitnil,omitempty" name:"VipDays"`

	// 充值分类。取值有：room_card-包月房卡; 其他-保留。
	GiveType *string `json:"GiveType,omitnil,omitempty" name:"GiveType"`

	// 播放场景。默认为Live
	// <li>Live：直播</li><li>Chat：语聊</li>
	PlayScene *string `json:"PlayScene,omitnil,omitempty" name:"PlayScene"`
}

type RechargeLiveVipRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 交易流水号，用于标记此次充值记录，多次充值记录传入相同的 TradeSerialNo 会判断为失败，可用于防止重提提交造成重复计费。
	TradeSerialNo *string `json:"TradeSerialNo,omitnil,omitempty" name:"TradeSerialNo"`

	// 房间标识。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 充值会员天数。
	// 取值有：
	// <li>31</li>
	// <li>93</li>
	// <li>186</li>
	// <li>372</li>
	VipDays *int64 `json:"VipDays,omitnil,omitempty" name:"VipDays"`

	// 充值分类。取值有：room_card-包月房卡; 其他-保留。
	GiveType *string `json:"GiveType,omitnil,omitempty" name:"GiveType"`

	// 播放场景。默认为Live
	// <li>Live：直播</li><li>Chat：语聊</li>
	PlayScene *string `json:"PlayScene,omitnil,omitempty" name:"PlayScene"`
}

func (r *RechargeLiveVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RechargeLiveVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "TradeSerialNo")
	delete(f, "RoomId")
	delete(f, "VipDays")
	delete(f, "GiveType")
	delete(f, "PlayScene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RechargeLiveVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RechargeLiveVipResponseParams struct {
	// 直播会员信息。
	LiveVipUserInfo *LiveVipUserInfo `json:"LiveVipUserInfo,omitnil,omitempty" name:"LiveVipUserInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RechargeLiveVipResponse struct {
	*tchttp.BaseResponse
	Response *RechargeLiveVipResponseParams `json:"Response"`
}

func (r *RechargeLiveVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RechargeLiveVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RechargeVipRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 房间Id。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 充值会员天数。(取值有：31、93、186、372)
	VipDays *int64 `json:"VipDays,omitnil,omitempty" name:"VipDays"`

	// 主播id。
	AnchorId *string `json:"AnchorId,omitnil,omitempty" name:"AnchorId"`
}

type RechargeVipRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 房间Id。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 充值会员天数。(取值有：31、93、186、372)
	VipDays *int64 `json:"VipDays,omitnil,omitempty" name:"VipDays"`

	// 主播id。
	AnchorId *string `json:"AnchorId,omitnil,omitempty" name:"AnchorId"`
}

func (r *RechargeVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RechargeVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "RoomId")
	delete(f, "VipDays")
	delete(f, "AnchorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RechargeVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RechargeVipResponseParams struct {
	// 厂商订单号。
	PartnerNo *string `json:"PartnerNo,omitnil,omitempty" name:"PartnerNo"`

	// TME订单号。
	OrderNo *string `json:"OrderNo,omitnil,omitempty" name:"OrderNo"`

	// 订单创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RechargeVipResponse struct {
	*tchttp.BaseResponse
	Response *RechargeVipResponseParams `json:"Response"`
}

func (r *RechargeVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RechargeVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchKTVMusicsRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 关键词。
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// 滚动标记。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 返回条数限制，默认 20，最大 50.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 权益过滤，取值有：
	// <li>Play：可播；</li>
	// <li>Sing：可唱。</li>
	RightFilters []*string `json:"RightFilters,omitnil,omitempty" name:"RightFilters"`

	// 播放场景。默认为Chat
	// <li>Live：直播</li><li>Chat：语聊</li>
	PlayScene *string `json:"PlayScene,omitnil,omitempty" name:"PlayScene"`

	// 物料过滤，取值有：
	// <li>Lyrics：含有歌词；</li>
	// <li>Midi：含有音高线。</li>
	MaterialFilters []*string `json:"MaterialFilters,omitnil,omitempty" name:"MaterialFilters"`
}

type SearchKTVMusicsRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 关键词。
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// 滚动标记。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 返回条数限制，默认 20，最大 50.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 权益过滤，取值有：
	// <li>Play：可播；</li>
	// <li>Sing：可唱。</li>
	RightFilters []*string `json:"RightFilters,omitnil,omitempty" name:"RightFilters"`

	// 播放场景。默认为Chat
	// <li>Live：直播</li><li>Chat：语聊</li>
	PlayScene *string `json:"PlayScene,omitnil,omitempty" name:"PlayScene"`

	// 物料过滤，取值有：
	// <li>Lyrics：含有歌词；</li>
	// <li>Midi：含有音高线。</li>
	MaterialFilters []*string `json:"MaterialFilters,omitnil,omitempty" name:"MaterialFilters"`
}

func (r *SearchKTVMusicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchKTVMusicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "KeyWord")
	delete(f, "ScrollToken")
	delete(f, "Limit")
	delete(f, "RightFilters")
	delete(f, "PlayScene")
	delete(f, "MaterialFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchKTVMusicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchKTVMusicsResponseParams struct {
	// 歌曲信息列表。
	KTVMusicInfoSet []*KTVMusicBaseInfo `json:"KTVMusicInfoSet,omitnil,omitempty" name:"KTVMusicInfoSet"`

	// 滚动标记，用于设置下次请求的 ScrollToken 参数。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchKTVMusicsResponse struct {
	*tchttp.BaseResponse
	Response *SearchKTVMusicsResponseParams `json:"Response"`
}

func (r *SearchKTVMusicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchKTVMusicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SeekCommandInput struct {
	// 播放位置，单位：毫秒。
	Position *uint64 `json:"Position,omitnil,omitempty" name:"Position"`
}

type SendMessageCommandInput struct {
	// 自定义消息，json格式字符串。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 消息重复次数，默认为 1。
	Repeat *uint64 `json:"Repeat,omitnil,omitempty" name:"Repeat"`
}

type SetAudioParamCommandInput struct {
	// 音频类型，取值有：
	// <li>Original：原唱</li>
	// <li>Accompaniment：伴奏</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type SetDestroyModeCommandInput struct {
	// 销毁模式，取值有：
	// <li>Auto：房间没人时自动销毁</li>
	// <li>Expire：房间没人时过期自动销毁</li>
	// <li>Never：不自动销毁，需手动销毁</li>默认为：Auto。
	DestroyMode *string `json:"DestroyMode,omitnil,omitempty" name:"DestroyMode"`

	// 过期销毁时间，单位：秒，当DestroyMode取Expire时必填。
	DestroyExpireTime *int64 `json:"DestroyExpireTime,omitnil,omitempty" name:"DestroyExpireTime"`
}

type SetPlayModeCommandInput struct {
	// 播放模式，取值有：
	// <li>RepeatPlaylist：列表循环</li>
	// <li>Order：顺序播放</li>
	// <li>RepeatSingle：单曲循环</li>
	// <li>Shuffle：随机播放</li>
	PlayMode *string `json:"PlayMode,omitnil,omitempty" name:"PlayMode"`
}

type SetPlaylistCommandInput struct {
	// 变更类型，取值有：
	// <li>Add：添加</li>
	// <li>Delete：删除</li>
	// <li>ClearList：清空歌曲列表</li>
	// <li>Move：移动歌曲</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 歌单索引位置，
	// 当 Type 取 Add 时，-1表示添加在列表最后位置，大于-1表示要添加的位置；
	// 当 Type 取 Delete 时，表示待删除歌曲的位置；
	// 当 Type 取 Move 时，表示待调整歌曲的位置。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 当 Type 取 Move 时，必填，表示移动歌曲的目标位置。
	ChangedIndex *int64 `json:"ChangedIndex,omitnil,omitempty" name:"ChangedIndex"`

	// 歌曲 ID 列表，当 Type 取 Add 时，必填。
	MusicIds []*string `json:"MusicIds,omitnil,omitempty" name:"MusicIds"`
}

// Predefined struct for user
type SyncKTVRobotCommandRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 机器人Id。
	RobotId *string `json:"RobotId,omitnil,omitempty" name:"RobotId"`

	// 指令及指令参数数组。
	SyncRobotCommands []*SyncRobotCommand `json:"SyncRobotCommands,omitnil,omitempty" name:"SyncRobotCommands"`
}

type SyncKTVRobotCommandRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 机器人Id。
	RobotId *string `json:"RobotId,omitnil,omitempty" name:"RobotId"`

	// 指令及指令参数数组。
	SyncRobotCommands []*SyncRobotCommand `json:"SyncRobotCommands,omitnil,omitempty" name:"SyncRobotCommands"`
}

func (r *SyncKTVRobotCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncKTVRobotCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "UserId")
	delete(f, "RobotId")
	delete(f, "SyncRobotCommands")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncKTVRobotCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncKTVRobotCommandResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncKTVRobotCommandResponse struct {
	*tchttp.BaseResponse
	Response *SyncKTVRobotCommandResponseParams `json:"Response"`
}

func (r *SyncKTVRobotCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncKTVRobotCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncRobotCommand struct {
	// 可同时传入多个指令，顺序执行。取值有：
	// <li>Play：播放</li>
	// <li>Pause：暂停</li>
	// <li>SwitchPrevious：上一首</li>
	// <li>SwitchNext：下一首</li>
	// <li>SetPlayMode：设置播放模式</li>
	// <li>Seek：调整播放进度</li>
	// <li>SetPlaylist：歌单变更</li>
	// <li>SetAudioParam：音频参数变更</li>
	// <li>SendMessage：发送自定义消息</li>
	// <li>SetDestroyMode：设置销毁模式</li>
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 播放参数。
	PlayCommandInput *PlayCommandInput `json:"PlayCommandInput,omitnil,omitempty" name:"PlayCommandInput"`

	// 播放列表变更信息，当Command取SetPlaylist时，必填。
	SetPlaylistCommandInput *SetPlaylistCommandInput `json:"SetPlaylistCommandInput,omitnil,omitempty" name:"SetPlaylistCommandInput"`

	// 播放进度，当Command取Seek时，必填。
	SeekCommandInput *SeekCommandInput `json:"SeekCommandInput,omitnil,omitempty" name:"SeekCommandInput"`

	// 音频参数，当Command取SetAudioParam时，必填。
	SetAudioParamCommandInput *SetAudioParamCommandInput `json:"SetAudioParamCommandInput,omitnil,omitempty" name:"SetAudioParamCommandInput"`

	// 自定义消息，当Command取SendMessage时，必填。
	SendMessageCommandInput *SendMessageCommandInput `json:"SendMessageCommandInput,omitnil,omitempty" name:"SendMessageCommandInput"`

	// 播放模式，当Command取SetPlayMode时，必填。
	SetPlayModeCommandInput *SetPlayModeCommandInput `json:"SetPlayModeCommandInput,omitnil,omitempty" name:"SetPlayModeCommandInput"`

	// 销毁模式，当Command取SetDestroyMode时，必填。
	SetDestroyModeCommandInput *SetDestroyModeCommandInput `json:"SetDestroyModeCommandInput,omitnil,omitempty" name:"SetDestroyModeCommandInput"`
}

type TRTCJoinRoomInput struct {
	// 签名。
	Sign *string `json:"Sign,omitnil,omitempty" name:"Sign"`

	// 房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 推流应用ID。
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户唯一标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// TRTC房间号的类型：
	// 
	// Integer：数字类型
	// String：字符串类型
	// 默认为：Integer 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomIdType *string `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`
}

type TimeRange struct {
	// <li>大于等于此时间（起始时间）。</li>
	// <li>格式按照 ISO 8601标准表示，详见 <a href="https://cloud.tencent.com/document/product/266/11732#I" target="_blank">ISO 日期格式说明</a>。</li>
	Before *string `json:"Before,omitnil,omitempty" name:"Before"`

	// <li>小于此时间（结束时间）。</li>
	// <li>格式按照 ISO 8601标准表示，详见 <a href="https://cloud.tencent.com/document/product/266/11732#I" target="_blank">ISO 日期格式说明</a>。</li>
	After *string `json:"After,omitnil,omitempty" name:"After"`
}

type UserInfo struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 直播会员详细信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveVipUserInfo *LiveVipUserInfo `json:"LiveVipUserInfo,omitnil,omitempty" name:"LiveVipUserInfo"`

	// 用户类型
	// <li>Normal：普通用户</li>
	// <li>LiveVip：直播会员用户</li>
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`
}