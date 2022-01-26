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

package v20190916

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Album struct {

	// 专辑名
	AlbumName *string `json:"AlbumName,omitempty" name:"AlbumName"`

	// 专辑图片大小及类别
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImagePathMap []*ImagePath `json:"ImagePathMap,omitempty" name:"ImagePathMap"`
}

type Artist struct {

	// 歌手名
	ArtistName *string `json:"ArtistName,omitempty" name:"ArtistName"`
}

type AuthInfo struct {

	// 主体名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectName *string `json:"SubjectName,omitempty" name:"SubjectName"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 应用场景
	AppScene *int64 `json:"AppScene,omitempty" name:"AppScene"`

	// 应用地域
	AppRegion *int64 `json:"AppRegion,omitempty" name:"AppRegion"`

	// 授权时间
	AuthPeriod *int64 `json:"AuthPeriod,omitempty" name:"AuthPeriod"`

	// 是否可商业化
	Commercialization *int64 `json:"Commercialization,omitempty" name:"Commercialization"`

	// 是否可跨平台
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`

	// 加密后Id
	Id *string `json:"Id,omitempty" name:"Id"`
}

type CreateKTVRobotRequest struct {
	*tchttp.BaseRequest

	// RTC厂商类型，取值有：
	// <li>TRTC</li>
	RTCSystem *string `json:"RTCSystem,omitempty" name:"RTCSystem"`

	// 进房参数。
	JoinRoomInput *JoinRoomInput `json:"JoinRoomInput,omitempty" name:"JoinRoomInput"`
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
	delete(f, "RTCSystem")
	delete(f, "JoinRoomInput")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKTVRobotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateKTVRobotResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机器人Id。
		RobotId *string `json:"RobotId,omitempty" name:"RobotId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DataInfo struct {

	// Song Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// 歌曲版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 歌曲总时长（非试听时长）
	Duration *string `json:"Duration,omitempty" name:"Duration"`

	// 试听开始时间
	AuditionBegin *uint64 `json:"AuditionBegin,omitempty" name:"AuditionBegin"`

	// 试听结束时间
	AuditionEnd *uint64 `json:"AuditionEnd,omitempty" name:"AuditionEnd"`

	// 标签名称
	TagNames []*string `json:"TagNames,omitempty" name:"TagNames"`
}

type DescribeAuthInfoRequest struct {
	*tchttp.BaseRequest

	// 偏移量：Offset=Offset+Limit
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *DescribeAuthInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Key")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuthInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 授权项目列表
		AuthInfo []*AuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`

		// 总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudMusicPurchasedRequest struct {
	*tchttp.BaseRequest

	// 授权项目Id
	AuthInfoId *string `json:"AuthInfoId,omitempty" name:"AuthInfoId"`
}

func (r *DescribeCloudMusicPurchasedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudMusicPurchasedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuthInfoId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudMusicPurchasedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudMusicPurchasedResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云音乐列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		MusicOpenDetail []*MusicOpenDetail `json:"MusicOpenDetail,omitempty" name:"MusicOpenDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudMusicPurchasedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudMusicPurchasedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudMusicRequest struct {
	*tchttp.BaseRequest

	// 歌曲Id
	MusicId *string `json:"MusicId,omitempty" name:"MusicId"`

	// 歌曲类型，可选值有：
	// <li>MP3-128K-FTW：含有水印的试听资源；</li>
	// <li>MP3-320K-FTD-P：320kbps歌曲热门片段；</li>
	// <li>MP3-320K-FTD：320kbps已核验歌曲完整资源。</li>
	// 默认为：MP3-128K-FTW
	MusicType *string `json:"MusicType,omitempty" name:"MusicType"`
}

func (r *DescribeCloudMusicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudMusicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MusicId")
	delete(f, "MusicType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudMusicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudMusicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 歌曲Id
		MusicId *string `json:"MusicId,omitempty" name:"MusicId"`

		// 歌曲名称
		MusicName *string `json:"MusicName,omitempty" name:"MusicName"`

		// 歌曲时长
	// 注意：此字段可能返回 null，表示取不到有效值。
		Duration *int64 `json:"Duration,omitempty" name:"Duration"`

		// 歌曲链接
		MusicUrl *string `json:"MusicUrl,omitempty" name:"MusicUrl"`

		// 歌曲图片
	// 注意：此字段可能返回 null，表示取不到有效值。
		MusicImageUrl *string `json:"MusicImageUrl,omitempty" name:"MusicImageUrl"`

		// 歌手列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Singers []*string `json:"Singers,omitempty" name:"Singers"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudMusicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudMusicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeItemByIdRequest struct {
	*tchttp.BaseRequest

	// 歌曲ID，目前暂不支持批量查询
	ItemIDs *string `json:"ItemIDs,omitempty" name:"ItemIDs"`
}

func (r *DescribeItemByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeItemByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ItemIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeItemByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeItemByIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 歌曲信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Items []*Item `json:"Items,omitempty" name:"Items"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeItemByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeItemByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeItemsRequest struct {
	*tchttp.BaseRequest

	// offset (Default = 0)，(当前页-1) * Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 条数，必须大于0，最大值为30
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// （电台/歌单）ID，CategoryId和CategoryCode两个必传1个，可以从<a href="https://cloud.tencent.com/document/product/1155/40109">获取分类内容（Station）列表接口</a>中获取。
	CategoryId *string `json:"CategoryId,omitempty" name:"CategoryId"`

	// （电台/歌单）ID，CategoryId和CategoryCode两个必传1个，可以从<a href="https://cloud.tencent.com/document/product/1155/40109">获取分类内容（Station）列表接口</a>中获取。
	CategoryCode *string `json:"CategoryCode,omitempty" name:"CategoryCode"`
}

func (r *DescribeItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CategoryId")
	delete(f, "CategoryCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeItemsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页偏移量
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// 当前页歌曲数量
		Size *uint64 `json:"Size,omitempty" name:"Size"`

		// 总数据条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 剩余数量（total-offset-size），通过这个值判断是否
	// 还有下一页
		HaveMore *uint64 `json:"HaveMore,omitempty" name:"HaveMore"`

		// Items 歌曲列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Items []*Item `json:"Items,omitempty" name:"Items"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKTVMusicDetailRequest struct {
	*tchttp.BaseRequest

	// 曲目 Id
	MusicId *string `json:"MusicId,omitempty" name:"MusicId"`
}

func (r *DescribeKTVMusicDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVMusicDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MusicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVMusicDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKTVMusicDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 歌曲基础信息
		KTVMusicBaseInfo *KTVMusicBaseInfo `json:"KTVMusicBaseInfo,omitempty" name:"KTVMusicBaseInfo"`

		// 播放凭证
		PlayToken *string `json:"PlayToken,omitempty" name:"PlayToken"`

		// 歌词下载地址
		LyricsUrl *string `json:"LyricsUrl,omitempty" name:"LyricsUrl"`

		// 歌曲规格信息列表
		DefinitionInfoSet []*KTVMusicDefinitionInfo `json:"DefinitionInfoSet,omitempty" name:"DefinitionInfoSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKTVMusicDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVMusicDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKTVPlaylistDetailRequest struct {
	*tchttp.BaseRequest

	// 歌单Id
	PlaylistId *string `json:"PlaylistId,omitempty" name:"PlaylistId"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	delete(f, "PlaylistId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVPlaylistDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKTVPlaylistDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 歌曲基础信息列表
		KTVMusicInfoSet []*KTVMusicBaseInfo `json:"KTVMusicInfoSet,omitempty" name:"KTVMusicInfoSet"`

		// 歌单基础信息
		PlaylistBaseInfo *KTVPlaylistBaseInfo `json:"PlaylistBaseInfo,omitempty" name:"PlaylistBaseInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeKTVPlaylistsRequest struct {
	*tchttp.BaseRequest

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVPlaylistsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKTVPlaylistsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 推荐歌单列表
		PlaylistBaseInfoSet []*KTVPlaylistBaseInfo `json:"PlaylistBaseInfoSet,omitempty" name:"PlaylistBaseInfoSet"`

		// 推荐歌单列表总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeKTVRobotsRequest struct {
	*tchttp.BaseRequest

	// 机器人Id列表。
	RobotIds []*string `json:"RobotIds,omitempty" name:"RobotIds"`

	// 机器人状态，取值有：
	// <li>Play：播放</li>
	// <li>Pause：暂停</li>
	// <li>Destroy：销毁</li>
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// 匹配创建时间在此时间段内的机器人。
	// <li>包含所指定的头尾时间点。</li>
	CreateTime *TimeRange `json:"CreateTime,omitempty" name:"CreateTime"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的起始偏移量，默认值：10。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

type DescribeKTVRobotsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机器人总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 机器人信息集合。
		KTVRobotInfoSet []*KTVRobotInfo `json:"KTVRobotInfoSet,omitempty" name:"KTVRobotInfoSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeLyricRequest struct {
	*tchttp.BaseRequest

	// 歌曲ID
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 格式，可选项，可不填写，默认值为：LRC-LRC。
	// <li>LRC-LRC：歌词；</li>
	// <li>JSON-ST：波形图。</li>
	SubItemType *string `json:"SubItemType,omitempty" name:"SubItemType"`
}

func (r *DescribeLyricRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLyricRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ItemId")
	delete(f, "SubItemType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLyricRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLyricResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 歌词或者波形图详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		Lyric *Lyric `json:"Lyric,omitempty" name:"Lyric"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLyricResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLyricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMusicRequest struct {
	*tchttp.BaseRequest

	// 歌曲ID
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 在应用前端播放音乐C端用户的唯一标识。无需是账户信息，用户唯一标识即可。
	IdentityId *string `json:"IdentityId,omitempty" name:"IdentityId"`

	// MP3-320K-FTD-P  为获取320kbps歌曲热门片段。
	// MP3-320K-FTD 为获取320kbps已核验歌曲完整资源。
	SubItemType *string `json:"SubItemType,omitempty" name:"SubItemType"`

	// CDN URL Protocol:HTTP or HTTPS/SSL
	// Values:Y , N(default)
	Ssl *string `json:"Ssl,omitempty" name:"Ssl"`
}

func (r *DescribeMusicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMusicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ItemId")
	delete(f, "IdentityId")
	delete(f, "SubItemType")
	delete(f, "Ssl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMusicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMusicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 音乐相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Music *Music `json:"Music,omitempty" name:"Music"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMusicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMusicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMusicSaleStatusRequest struct {
	*tchttp.BaseRequest

	// 歌曲Id集合，可传单个，也可传多个，上线查询单次50个
	MusicIds []*string `json:"MusicIds,omitempty" name:"MusicIds"`

	// 查询哪个渠道的数据，1为曲库包，2为单曲
	PurchaseType *int64 `json:"PurchaseType,omitempty" name:"PurchaseType"`
}

func (r *DescribeMusicSaleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMusicSaleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MusicIds")
	delete(f, "PurchaseType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMusicSaleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMusicSaleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// musicId对应歌曲状态
		MusicStatusSet []*MusicStatus `json:"MusicStatusSet,omitempty" name:"MusicStatusSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMusicSaleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMusicSaleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackageItemsRequest struct {
	*tchttp.BaseRequest

	// 订单id，从获取已购曲库包列表中获取
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 默认0，Offset=Offset+Length
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 默认20
	Length *uint64 `json:"Length,omitempty" name:"Length"`
}

func (r *DescribePackageItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackageItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	delete(f, "Offset")
	delete(f, "Length")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePackageItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackageItemsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 已核销歌曲信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		PackageItems []*PackageItem `json:"PackageItems,omitempty" name:"PackageItems"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePackageItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackageItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackagesRequest struct {
	*tchttp.BaseRequest

	// 默认0，Offset=Offset+Length
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 默认20
	Length *uint64 `json:"Length,omitempty" name:"Length"`
}

func (r *DescribePackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Length")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePackagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 已购曲库包列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Packages []*Package `json:"Packages,omitempty" name:"Packages"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePkgOfflineMusicRequest struct {
	*tchttp.BaseRequest

	// 订单id
	PackageOrderId *string `json:"PackageOrderId,omitempty" name:"PackageOrderId"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条(注：单次上限为100)。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePkgOfflineMusicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePkgOfflineMusicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageOrderId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePkgOfflineMusicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePkgOfflineMusicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 曲库包中不可用歌曲信息
		OfflineMusicSet []*OfflineMusicDetail `json:"OfflineMusicSet,omitempty" name:"OfflineMusicSet"`

		// 返回总量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePkgOfflineMusicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePkgOfflineMusicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStationsRequest struct {
	*tchttp.BaseRequest

	// 条数，必须大于0
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// offset (Default = 0)，Offset=Offset+Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeStationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 分页偏移量
		Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

		// 当前页station数量
		Size *uint64 `json:"Size,omitempty" name:"Size"`

		// 剩余数量（total-offset-size），通过这个值判断是否还有下一页
		HaveMore *uint64 `json:"HaveMore,omitempty" name:"HaveMore"`

		// Stations 素材库列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Stations []*Station `json:"Stations,omitempty" name:"Stations"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyKTVRobotRequest struct {
	*tchttp.BaseRequest

	// 机器人Id。
	RobotId *string `json:"RobotId,omitempty" name:"RobotId"`
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
	delete(f, "RobotId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyKTVRobotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DestroyKTVRobotResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ImagePath struct {

	// station图片大小及类别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// station图片地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Item struct {

	// Song ID
	ItemID *string `json:"ItemID,omitempty" name:"ItemID"`

	// Song info
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataInfo *DataInfo `json:"DataInfo,omitempty" name:"DataInfo"`

	// 专辑信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Album *Album `json:"Album,omitempty" name:"Album"`

	// 多个歌手集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Artists []*Artist `json:"Artists,omitempty" name:"Artists"`

	// 歌曲状态，1:添加进购物车；2:核销进曲库包
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type JoinRoomInput struct {

	// TRTC进房参数
	TRTCJoinRoomInput *TRTCJoinRoomInput `json:"TRTCJoinRoomInput,omitempty" name:"TRTCJoinRoomInput"`
}

type KTVMusicBaseInfo struct {

	// 歌曲 Id
	MusicId *string `json:"MusicId,omitempty" name:"MusicId"`

	// 歌曲名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 演唱者列表
	SingerSet []*string `json:"SingerSet,omitempty" name:"SingerSet"`

	// 作词者列表
	LyricistSet []*string `json:"LyricistSet,omitempty" name:"LyricistSet"`

	// 作曲者列表
	ComposerSet []*string `json:"ComposerSet,omitempty" name:"ComposerSet"`

	// 标签列表
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet"`

	// 歌曲时长
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
}

type KTVMusicDefinitionInfo struct {

	// 规格，取值有：
	// <li>audio/mi：低规格；</li>
	// <li>audio/lo：中规格；</li>
	// <li>audio/hi：高规格。</li>
	Definition *string `json:"Definition,omitempty" name:"Definition"`

	// 码率，单位为 bps。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 文件大小，单位为字节。
	Size *int64 `json:"Size,omitempty" name:"Size"`
}

type KTVPlaylistBaseInfo struct {

	// 歌单Id
	PlaylistId *string `json:"PlaylistId,omitempty" name:"PlaylistId"`

	// 歌单标题
	Title *string `json:"Title,omitempty" name:"Title"`

	// 歌单介绍
	Description *string `json:"Description,omitempty" name:"Description"`

	// 歌曲数量
	MusicNum *int64 `json:"MusicNum,omitempty" name:"MusicNum"`
}

type KTVRobotInfo struct {

	// 机器人Id。
	RobotId *string `json:"RobotId,omitempty" name:"RobotId"`

	// 状态，取值有：
	// <li>Play：播放</li>
	// <li>Pause：暂停</li>
	// <li>Destroy：销毁</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 播放列表。
	Playlists []*string `json:"Playlists,omitempty" name:"Playlists"`

	// 当前歌单索引位置。
	CurIndex *int64 `json:"CurIndex,omitempty" name:"CurIndex"`

	// 播放进度，单位：毫秒。
	Position *uint64 `json:"Position,omitempty" name:"Position"`

	// 音频参数
	SetAudioParamInput *SetAudioParamCommandInput `json:"SetAudioParamInput,omitempty" name:"SetAudioParamInput"`

	// 进房信息
	JoinRoomInput *JoinRoomInput `json:"JoinRoomInput,omitempty" name:"JoinRoomInput"`

	// RTC厂商类型，取值有：
	// <li>TRTC</li>
	RTCSystem *string `json:"RTCSystem,omitempty" name:"RTCSystem"`

	// 播放模式，PlayMode取值有：
	// <li>RepeatPlaylist：列表循环</li>
	// <li>Order：顺序播放</li>
	// <li>RepeatSingle：单曲循环</li>
	// <li>Shuffle：随机播放</li>
	SetPlayModeInput *SetPlayModeCommandInput `json:"SetPlayModeInput,omitempty" name:"SetPlayModeInput"`
}

type Lyric struct {

	// 歌词cdn地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 歌词后缀名
	FileNameExt *string `json:"FileNameExt,omitempty" name:"FileNameExt"`

	// 歌词类型
	SubItemType *string `json:"SubItemType,omitempty" name:"SubItemType"`
}

type ModifyMusicOnShelvesRequest struct {
	*tchttp.BaseRequest

	// 歌曲变更信息
	MusicDetailInfos *MusicDetailInfo `json:"MusicDetailInfos,omitempty" name:"MusicDetailInfos"`

	// ame对接资源方密钥
	AmeKey *string `json:"AmeKey,omitempty" name:"AmeKey"`
}

func (r *ModifyMusicOnShelvesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMusicOnShelvesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MusicDetailInfos")
	delete(f, "AmeKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMusicOnShelvesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMusicOnShelvesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMusicOnShelvesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMusicOnShelvesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Music struct {

	// 音乐播放链接相对路径，必须通过在正版曲库直通车控制台上登记的域名进行拼接。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 音频文件大小
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 音频文件类型
	FileExtension *string `json:"FileExtension,omitempty" name:"FileExtension"`

	// Song fragment start.试听片段开始时间，试听时长为auditionEnd-auditionBegin
	// Unit :ms
	AuditionBegin *uint64 `json:"AuditionBegin,omitempty" name:"AuditionBegin"`

	// Song fragment end.试听片段结束时间, 试听时长为auditionEnd-auditionBegin
	// Unit :ms
	AuditionEnd *uint64 `json:"AuditionEnd,omitempty" name:"AuditionEnd"`

	// 音乐播放链接全路径，前提是在正版曲库直通车控制台添加过域名，否则返回空字符。
	// 如果添加过多个域名只返回第一个添加域名的播放全路径。
	FullUrl *string `json:"FullUrl,omitempty" name:"FullUrl"`
}

type MusicDetailInfo struct {

	// 资源方音乐Id
	MusicId *string `json:"MusicId,omitempty" name:"MusicId"`

	// 资源方识别信息
	AmeId *string `json:"AmeId,omitempty" name:"AmeId"`

	// 分类标签
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 关键词
	HitWords []*string `json:"HitWords,omitempty" name:"HitWords"`

	// 节奏信息
	Bpm *int64 `json:"Bpm,omitempty" name:"Bpm"`

	// 商业化权益
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 应用歌曲信息,1.图文/短视频,2.网络直播,3.网络电台FM,4.免费游戏,5.商业游戏,6.网店网站设计,7.广告营销,8.网络长视频
	Scene []*string `json:"Scene,omitempty" name:"Scene"`

	// 应用地域,1. 中国大陆,2. 中国含港澳台,3. 全球
	Region []*string `json:"Region,omitempty" name:"Region"`

	// 授权时间,1. 1年, 5. 随片永久
	AuthPeriod *string `json:"AuthPeriod,omitempty" name:"AuthPeriod"`

	// 商业化授权，1. 支持商业化 ,2. 不支持商业化
	Commercialization *string `json:"Commercialization,omitempty" name:"Commercialization"`

	// 跨平台传播，1. 支持跨平台传播 ,2. 不支持跨平台传播
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 传播渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

type MusicOpenDetail struct {

	// 音乐Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MusicId *string `json:"MusicId,omitempty" name:"MusicId"`

	// 专辑名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbumName *string `json:"AlbumName,omitempty" name:"AlbumName"`

	// 专辑图片路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbumImageUrl *string `json:"AlbumImageUrl,omitempty" name:"AlbumImageUrl"`

	// 音乐名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MusicName *string `json:"MusicName,omitempty" name:"MusicName"`

	// 音乐图片路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	MusicImageUrl *string `json:"MusicImageUrl,omitempty" name:"MusicImageUrl"`

	// 歌手
	// 注意：此字段可能返回 null，表示取不到有效值。
	Singers []*string `json:"Singers,omitempty" name:"Singers"`

	// 播放时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 歌词url
	// 注意：此字段可能返回 null，表示取不到有效值。
	LyricUrl *string `json:"LyricUrl,omitempty" name:"LyricUrl"`

	// 波形图url
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaveformUrl *string `json:"WaveformUrl,omitempty" name:"WaveformUrl"`
}

type MusicStatus struct {

	// 歌曲Id
	MusicId *string `json:"MusicId,omitempty" name:"MusicId"`

	// 在售状态,0为在售，1为临时下架，2为永久下架
	SaleStatus *int64 `json:"SaleStatus,omitempty" name:"SaleStatus"`
}

type OfflineMusicDetail struct {

	// 歌曲Id
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 歌曲名称
	MusicName *string `json:"MusicName,omitempty" name:"MusicName"`

	// 不可用原因
	OffRemark *string `json:"OffRemark,omitempty" name:"OffRemark"`

	// 不可用时间
	OffTime *string `json:"OffTime,omitempty" name:"OffTime"`
}

type Package struct {

	// 订单id
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 曲库包名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 授权地区-global: 全球  CN: 中国
	AuthorizedArea *string `json:"AuthorizedArea,omitempty" name:"AuthorizedArea"`

	// 授权次数
	AuthorizedLimit *int64 `json:"AuthorizedLimit,omitempty" name:"AuthorizedLimit"`

	// 套餐有效期，单位:天
	TermOfValidity *int64 `json:"TermOfValidity,omitempty" name:"TermOfValidity"`

	// 0:不可商业化；1:可商业化
	Commercial *int64 `json:"Commercial,omitempty" name:"Commercial"`

	// 套餐价格，单位：元
	PackagePrice *float64 `json:"PackagePrice,omitempty" name:"PackagePrice"`

	// 生效开始时间,格式yyyy-MM-dd HH:mm:ss
	EffectTime *string `json:"EffectTime,omitempty" name:"EffectTime"`

	// 生效结束时间,格式yyyy-MM-dd HH:mm:ss
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 剩余授权次数
	UsedCount *int64 `json:"UsedCount,omitempty" name:"UsedCount"`

	// 曲库包用途信息
	UseRanges []*UseRange `json:"UseRanges,omitempty" name:"UseRanges"`
}

type PackageItem struct {

	// 订单id
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 歌曲名
	TrackName *string `json:"TrackName,omitempty" name:"TrackName"`

	// 歌曲ID
	ItemID *string `json:"ItemID,omitempty" name:"ItemID"`

	// 专辑图片
	Img *string `json:"Img,omitempty" name:"Img"`

	// 歌手名
	ArtistName *string `json:"ArtistName,omitempty" name:"ArtistName"`

	// 歌曲时长
	Duration *string `json:"Duration,omitempty" name:"Duration"`

	// 授权区域，global: 全球 CN: 中国
	AuthorizedArea *string `json:"AuthorizedArea,omitempty" name:"AuthorizedArea"`

	// 标签数组
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type PlayCommandInput struct {

	// 歌曲位置索引。
	Index *int64 `json:"Index,omitempty" name:"Index"`
}

type PutMusicOnTheShelvesRequest struct {
	*tchttp.BaseRequest

	// 资源方歌曲Id
	MusicIds []*string `json:"MusicIds,omitempty" name:"MusicIds"`
}

func (r *PutMusicOnTheShelvesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutMusicOnTheShelvesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MusicIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutMusicOnTheShelvesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PutMusicOnTheShelvesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作成功数量
		SuccessNum *int64 `json:"SuccessNum,omitempty" name:"SuccessNum"`

		// 操作失败数量
		FailedNum *int64 `json:"FailedNum,omitempty" name:"FailedNum"`

		// 失败歌曲Id
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailedMusicIds []*string `json:"FailedMusicIds,omitempty" name:"FailedMusicIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutMusicOnTheShelvesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutMusicOnTheShelvesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReportDataRequest struct {
	*tchttp.BaseRequest

	// 上报数据
	// 注:reportData为客户端压缩后的上报数据进行16进制转换的字符串数据
	// 压缩说明：
	// a) 上报的json格式字符串通过流的转换（ByteArrayInputStream, java.util.zip.GZIPOutputStream），获取到压缩后的字节数组。
	// b) 将压缩后的字节数组转成16进制字符串。
	// 
	// reportData由两部分数据组成：
	// 1）report_type（上报类型）
	// 2）data（歌曲上报数据）
	// 不同的report_type对应的data数据结构不一样。
	// 
	// 详细说明请参考文档reportdata.docx：
	// https://github.com/tencentyun/ame-documents
	ReportData *string `json:"ReportData,omitempty" name:"ReportData"`
}

func (r *ReportDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReportData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ReportDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReportDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchKTVMusicsRequest struct {
	*tchttp.BaseRequest

	// 搜索关键词
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	// 取值范围：Offset + Limit 不超过5000。取值范围：小于5000
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的起始偏移量，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	// 取值范围：Offset + Limit 不超过5000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	delete(f, "KeyWord")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchKTVMusicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SearchKTVMusicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// KTV 曲目列表
		KTVMusicInfoSet []*KTVMusicBaseInfo `json:"KTVMusicInfoSet,omitempty" name:"KTVMusicInfoSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	Position *uint64 `json:"Position,omitempty" name:"Position"`
}

type SendMessageCommandInput struct {

	// 自定义消息，json格式字符串。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 消息重复次数，默认为 1。
	Repeat *uint64 `json:"Repeat,omitempty" name:"Repeat"`
}

type SetAudioParamCommandInput struct {

	// 规格，取值有：
	// <li>audio/mi：低规格</li>
	// <li>audio/lo：中规格</li>
	// <li>audio/hi：高规格</li>
	Definition *string `json:"Definition,omitempty" name:"Definition"`

	// 音频类型，取值有：
	// <li>Original：原唱</li>
	// <li>Accompaniment：伴奏</li>
	Type *string `json:"Type,omitempty" name:"Type"`
}

type SetPlayModeCommandInput struct {

	// 播放模式，取值有：
	// <li>RepeatPlaylist：列表循环</li>
	// <li>Order：顺序播放</li>
	// <li>RepeatSingle：单曲循环</li>
	// <li>Shuffle：随机播放</li>
	PlayMode *string `json:"PlayMode,omitempty" name:"PlayMode"`
}

type SetPlaylistCommandInput struct {

	// 变更类型，取值有：
	// <li>Add：添加</li>
	// <li>Delete：删除</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 歌单索引位置，
	// 当 Type 取 Add 时，-1表示添加在列表最后位置，大于-1表示要添加的位置；
	// 当 Type 取 Delete 时，表示要删除的位置。
	Index *int64 `json:"Index,omitempty" name:"Index"`

	// 歌曲 ID 列表，当 Type 取 Add 时，必填。
	MusicIds []*string `json:"MusicIds,omitempty" name:"MusicIds"`
}

type Station struct {

	// StationID
	CategoryID *string `json:"CategoryID,omitempty" name:"CategoryID"`

	// Station MCCode
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryCode *string `json:"CategoryCode,omitempty" name:"CategoryCode"`

	// Category Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// Station的排序值，供参考（返回结果已按其升序）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rank *int64 `json:"Rank,omitempty" name:"Rank"`

	// station图片集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImagePathMap []*ImagePath `json:"ImagePathMap,omitempty" name:"ImagePathMap"`
}

type SyncKTVRobotCommandRequest struct {
	*tchttp.BaseRequest

	// 机器人Id。
	RobotId *string `json:"RobotId,omitempty" name:"RobotId"`

	// 指令，取值有：
	// <li>Play：播放</li>
	// <li>Pause：暂停</li>
	// <li>SwitchPrevious：上一首</li>
	// <li>SwitchNext：下一首</li>
	// <li>SetPlayMode：设置播放模式</li>
	// <li>Seek：调整播放进度</li>
	// <li>SetPlaylist：歌单变更</li>
	// <li>SetAudioParam：音频参数变更</li>
	// <li>SendMessage：发送自定义消息</li>
	Command *string `json:"Command,omitempty" name:"Command"`

	// 播放参数。
	PlayCommandInput *PlayCommandInput `json:"PlayCommandInput,omitempty" name:"PlayCommandInput"`

	// 播放列表变更信息，当Command取SetPlaylist时，必填。
	SetPlaylistCommandInput *SetPlaylistCommandInput `json:"SetPlaylistCommandInput,omitempty" name:"SetPlaylistCommandInput"`

	// 播放进度，当Command取Seek时，必填。
	SeekCommandInput *SeekCommandInput `json:"SeekCommandInput,omitempty" name:"SeekCommandInput"`

	// 音频参数，当Command取SetAudioParam时，必填。
	SetAudioParamCommandInput *SetAudioParamCommandInput `json:"SetAudioParamCommandInput,omitempty" name:"SetAudioParamCommandInput"`

	// 自定义消息，当Command取SendMessage时，必填。
	SendMessageCommandInput *SendMessageCommandInput `json:"SendMessageCommandInput,omitempty" name:"SendMessageCommandInput"`

	// 播放模式，当Command取SetPlayMode时，必填。
	SetPlayModeCommandInput *SetPlayModeCommandInput `json:"SetPlayModeCommandInput,omitempty" name:"SetPlayModeCommandInput"`
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
	delete(f, "RobotId")
	delete(f, "Command")
	delete(f, "PlayCommandInput")
	delete(f, "SetPlaylistCommandInput")
	delete(f, "SeekCommandInput")
	delete(f, "SetAudioParamCommandInput")
	delete(f, "SendMessageCommandInput")
	delete(f, "SetPlayModeCommandInput")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncKTVRobotCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SyncKTVRobotCommandResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type TRTCJoinRoomInput struct {

	// 签名。
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// 房间号。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 推流应用ID。
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户唯一标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type TakeMusicOffShelves struct {

	// 资源方对应音乐Id
	MusicIds *string `json:"MusicIds,omitempty" name:"MusicIds"`

	// 当曲目临时下架时：已订购客户无影响，无需消息通知。当曲目封杀下架后，推送消息至已订购老客户，枚举值，判断是否上/下架
	// 在售状态，0在售，1临时下架，2永久下架
	SaleStatus *string `json:"SaleStatus,omitempty" name:"SaleStatus"`
}

type TakeMusicOffShelvesRequest struct {
	*tchttp.BaseRequest

	// 资源方下架必传结构
	TakeMusicOffShelves []*TakeMusicOffShelves `json:"TakeMusicOffShelves,omitempty" name:"TakeMusicOffShelves"`
}

func (r *TakeMusicOffShelvesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TakeMusicOffShelvesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TakeMusicOffShelves")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TakeMusicOffShelvesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TakeMusicOffShelvesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回成功数量
		SuccessNum *int64 `json:"SuccessNum,omitempty" name:"SuccessNum"`

		// 返回失败数量
		FailedNum *int64 `json:"FailedNum,omitempty" name:"FailedNum"`

		// 返回失败歌曲musicId
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailedMusicIds []*string `json:"FailedMusicIds,omitempty" name:"FailedMusicIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TakeMusicOffShelvesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TakeMusicOffShelvesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TimeRange struct {

	// <li>大于等于此时间（起始时间）。</li>
	// <li>格式按照 ISO 8601标准表示，详见 <a href="https://cloud.tencent.com/document/product/266/11732#I" target="_blank">ISO 日期格式说明</a>。</li>
	Before *string `json:"Before,omitempty" name:"Before"`

	// <li>小于此时间（结束时间）。</li>
	// <li>格式按照 ISO 8601标准表示，详见 <a href="https://cloud.tencent.com/document/product/266/11732#I" target="_blank">ISO 日期格式说明</a>。</li>
	After *string `json:"After,omitempty" name:"After"`
}

type UseRange struct {

	// 用途id
	UseRangeId *int64 `json:"UseRangeId,omitempty" name:"UseRangeId"`

	// 用途范围名称
	Name *string `json:"Name,omitempty" name:"Name"`
}
