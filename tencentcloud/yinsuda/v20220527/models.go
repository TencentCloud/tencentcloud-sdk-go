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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type BatchDescribeKTVMusicDetailsRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 歌曲 Id 列表。
	MusicIds []*string `json:"MusicIds,omitempty" name:"MusicIds"`
}

type BatchDescribeKTVMusicDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 歌曲 Id 列表。
	MusicIds []*string `json:"MusicIds,omitempty" name:"MusicIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDescribeKTVMusicDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDescribeKTVMusicDetailsResponseParams struct {
	// 歌曲详细信息列表。
	KTVMusicDetailInfoSet []*KTVMusicDetailInfo `json:"KTVMusicDetailInfoSet,omitempty" name:"KTVMusicDetailInfoSet"`

	// 不存在歌曲Id列表。
	NotExistMusicIdSet []*string `json:"NotExistMusicIdSet,omitempty" name:"NotExistMusicIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，单位：毫秒。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

// Predefined struct for user
type DescribeKTVPlaylistDetailRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 歌单 Id。
	PlaylistId *string `json:"PlaylistId,omitempty" name:"PlaylistId"`

	// 滚动标记。
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

	// 返回条数，默认：20，最大：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 权益过滤，取值有：
	// <li>Play：可播；</li>
	// <li>Sing：可唱。</li>
	RightFilters []*string `json:"RightFilters,omitempty" name:"RightFilters"`
}

type DescribeKTVPlaylistDetailRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 歌单 Id。
	PlaylistId *string `json:"PlaylistId,omitempty" name:"PlaylistId"`

	// 滚动标记。
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

	// 返回条数，默认：20，最大：50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 权益过滤，取值有：
	// <li>Play：可播；</li>
	// <li>Sing：可唱。</li>
	RightFilters []*string `json:"RightFilters,omitempty" name:"RightFilters"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVPlaylistDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVPlaylistDetailResponseParams struct {
	// 歌曲信息列表。
	KTVMusicInfoSet []*KTVMusicBaseInfo `json:"KTVMusicInfoSet,omitempty" name:"KTVMusicInfoSet"`

	// 滚动标记，用于设置下次请求的 ScrollToken 参数。
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeKTVPlaylistsRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVPlaylistsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVPlaylistsResponseParams struct {
	// 歌单基础信息。
	PlaylistBaseInfoSet []*KTVPlaylistBaseInfo `json:"PlaylistBaseInfoSet,omitempty" name:"PlaylistBaseInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeKTVSuggestionsRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 搜索词。
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
}

type DescribeKTVSuggestionsRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 搜索词。
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
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
	KTVSuggestionInfoSet []*KTVSuggestionInfo `json:"KTVSuggestionInfoSet,omitempty" name:"KTVSuggestionInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type KTVMusicBaseInfo struct {
	// 歌曲Id。
	MusicId *string `json:"MusicId,omitempty" name:"MusicId"`

	// 歌曲名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 歌手名称。
	SingerSet []*string `json:"SingerSet,omitempty" name:"SingerSet"`

	// 播放时长。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 歌手图片链接。
	SingerImageUrl *string `json:"SingerImageUrl,omitempty" name:"SingerImageUrl"`

	// 专辑信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbumInfo *MusicAlbumInfo `json:"AlbumInfo,omitempty" name:"AlbumInfo"`

	// 权益列表，取值有：
	// <li>Play：可播；</li>
	// <li>Sing：可唱。</li>
	RightSet []*string `json:"RightSet,omitempty" name:"RightSet"`
}

type KTVMusicDetailInfo struct {
	// 歌曲基础信息。
	KTVMusicBaseInfo *KTVMusicBaseInfo `json:"KTVMusicBaseInfo,omitempty" name:"KTVMusicBaseInfo"`

	// 播放凭证。
	PlayToken *string `json:"PlayToken,omitempty" name:"PlayToken"`

	// 歌词下载链接。
	LyricsUrl *string `json:"LyricsUrl,omitempty" name:"LyricsUrl"`

	// 音高数据下载链接。
	MidiUrl *string `json:"MidiUrl,omitempty" name:"MidiUrl"`

	// 副歌片段信息。
	ChorusClipSet []*ChorusClip `json:"ChorusClipSet,omitempty" name:"ChorusClipSet"`
}

type KTVPlaylistBaseInfo struct {
	// 歌单Id。
	PlaylistId *string `json:"PlaylistId,omitempty" name:"PlaylistId"`

	// 歌单标题。
	Title *string `json:"Title,omitempty" name:"Title"`
}

type KTVSuggestionInfo struct {
	// 联想词。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
}

type MusicAlbumCoverInfo struct {
	// 尺寸规格，取值有：
	// <li>Mini：150 x 150 尺寸；</li>
	// <li>Small：240 x 240 尺寸；</li>
	// <li>Medium：480 x 480 尺寸。</li>
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 下载链接。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type MusicAlbumInfo struct {
	// 专辑名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 封面列表。
	CoverInfoSet []*MusicAlbumCoverInfo `json:"CoverInfoSet,omitempty" name:"CoverInfoSet"`
}

// Predefined struct for user
type SearchKTVMusicsRequestParams struct {
	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 关键词。
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 滚动标记。
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

	// 返回条数限制，默认 20，最大 50.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 权益过滤，取值有：
	// <li>Play：可播；</li>
	// <li>Sing：可唱。</li>
	RightFilters []*string `json:"RightFilters,omitempty" name:"RightFilters"`
}

type SearchKTVMusicsRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 用户标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 关键词。
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 滚动标记。
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

	// 返回条数限制，默认 20，最大 50.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 权益过滤，取值有：
	// <li>Play：可播；</li>
	// <li>Sing：可唱。</li>
	RightFilters []*string `json:"RightFilters,omitempty" name:"RightFilters"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchKTVMusicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchKTVMusicsResponseParams struct {
	// 歌曲信息列表。
	KTVMusicInfoSet []*KTVMusicBaseInfo `json:"KTVMusicInfoSet,omitempty" name:"KTVMusicInfoSet"`

	// 滚动标记，用于设置下次请求的 ScrollToken 参数。
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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