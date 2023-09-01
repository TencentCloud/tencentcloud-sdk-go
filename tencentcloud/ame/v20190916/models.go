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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Album struct {
	// 专辑名
	AlbumName *string `json:"AlbumName,omitnil" name:"AlbumName"`

	// 专辑图片大小及类别
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImagePathMap []*ImagePath `json:"ImagePathMap,omitnil" name:"ImagePathMap"`
}

type ApplicationLicenseInput struct {
	// 应用名称，注：后面三个字段AndroidPackageName、IOSBundleId、PcIdentifier，三者选填一个
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// app的安卓包名
	AndroidPackageName *string `json:"AndroidPackageName,omitnil" name:"AndroidPackageName"`

	// app的IOS的BundleId名
	IOSBundleId *string `json:"IOSBundleId,omitnil" name:"IOSBundleId"`

	// PC标识名
	PcIdentifier *string `json:"PcIdentifier,omitnil" name:"PcIdentifier"`
}

type Artist struct {
	// 歌手名
	ArtistName *string `json:"ArtistName,omitnil" name:"ArtistName"`
}

type AuthInfo struct {
	// 主体名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectName *string `json:"SubjectName,omitnil" name:"SubjectName"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// 应用场景
	AppScene *int64 `json:"AppScene,omitnil" name:"AppScene"`

	// 应用地域
	AppRegion *int64 `json:"AppRegion,omitnil" name:"AppRegion"`

	// 授权时间
	AuthPeriod *int64 `json:"AuthPeriod,omitnil" name:"AuthPeriod"`

	// 是否可商业化
	Commercialization *int64 `json:"Commercialization,omitnil" name:"Commercialization"`

	// 是否可跨平台
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 加密后Id
	Id *string `json:"Id,omitnil" name:"Id"`
}

// Predefined struct for user
type BatchDescribeKTVMusicDetailsRequestParams struct {
	// 歌曲Id列表，注：列表最大长度为50
	MusicIds []*string `json:"MusicIds,omitnil" name:"MusicIds"`
}

type BatchDescribeKTVMusicDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 歌曲Id列表，注：列表最大长度为50
	MusicIds []*string `json:"MusicIds,omitnil" name:"MusicIds"`
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
	delete(f, "MusicIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDescribeKTVMusicDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDescribeKTVMusicDetailsResponseParams struct {
	// 歌曲详情列表信息
	KTVMusicDetailInfoSet []*KTVMusicDetailInfo `json:"KTVMusicDetailInfoSet,omitnil" name:"KTVMusicDetailInfoSet"`

	// 不存在的歌曲 ID 列表。
	NotExistMusicIdSet []*string `json:"NotExistMusicIdSet,omitnil" name:"NotExistMusicIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 副歌时间，单位：毫秒
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 副歌结束时间，单位：毫秒
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`
}

// Predefined struct for user
type CreateKTVRobotRequestParams struct {
	// RTC厂商类型，取值有：
	// <li>TRTC</li>
	RTCSystem *string `json:"RTCSystem,omitnil" name:"RTCSystem"`

	// 进房参数。
	JoinRoomInput *JoinRoomInput `json:"JoinRoomInput,omitnil" name:"JoinRoomInput"`

	// license基础信息
	ApplicationLicenseInput *ApplicationLicenseInput `json:"ApplicationLicenseInput,omitnil" name:"ApplicationLicenseInput"`

	// 创建机器人时初始化参数。
	SyncRobotCommands []*SyncRobotCommand `json:"SyncRobotCommands,omitnil" name:"SyncRobotCommands"`
}

type CreateKTVRobotRequest struct {
	*tchttp.BaseRequest
	
	// RTC厂商类型，取值有：
	// <li>TRTC</li>
	RTCSystem *string `json:"RTCSystem,omitnil" name:"RTCSystem"`

	// 进房参数。
	JoinRoomInput *JoinRoomInput `json:"JoinRoomInput,omitnil" name:"JoinRoomInput"`

	// license基础信息
	ApplicationLicenseInput *ApplicationLicenseInput `json:"ApplicationLicenseInput,omitnil" name:"ApplicationLicenseInput"`

	// 创建机器人时初始化参数。
	SyncRobotCommands []*SyncRobotCommand `json:"SyncRobotCommands,omitnil" name:"SyncRobotCommands"`
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
	delete(f, "ApplicationLicenseInput")
	delete(f, "SyncRobotCommands")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKTVRobotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKTVRobotResponseParams struct {
	// 机器人Id。
	RobotId *string `json:"RobotId,omitnil" name:"RobotId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type DataInfo struct {
	// Song Name
	Name *string `json:"Name,omitnil" name:"Name"`

	// 歌曲版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 歌曲总时长（非试听时长）
	Duration *string `json:"Duration,omitnil" name:"Duration"`

	// 试听开始时间
	AuditionBegin *uint64 `json:"AuditionBegin,omitnil" name:"AuditionBegin"`

	// 试听结束时间
	AuditionEnd *uint64 `json:"AuditionEnd,omitnil" name:"AuditionEnd"`

	// 标签名称
	TagNames []*string `json:"TagNames,omitnil" name:"TagNames"`
}

// Predefined struct for user
type DescribeAuthInfoRequestParams struct {
	// 偏移量：Offset=Offset+Limit
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 数据条数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 搜索关键字
	Key *string `json:"Key,omitnil" name:"Key"`
}

type DescribeAuthInfoRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量：Offset=Offset+Limit
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 数据条数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 搜索关键字
	Key *string `json:"Key,omitnil" name:"Key"`
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

// Predefined struct for user
type DescribeAuthInfoResponseParams struct {
	// 授权项目列表
	AuthInfo []*AuthInfo `json:"AuthInfo,omitnil" name:"AuthInfo"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAuthInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuthInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCloudMusicPurchasedRequestParams struct {
	// 授权项目Id
	AuthInfoId *string `json:"AuthInfoId,omitnil" name:"AuthInfoId"`
}

type DescribeCloudMusicPurchasedRequest struct {
	*tchttp.BaseRequest
	
	// 授权项目Id
	AuthInfoId *string `json:"AuthInfoId,omitnil" name:"AuthInfoId"`
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

// Predefined struct for user
type DescribeCloudMusicPurchasedResponseParams struct {
	// 云音乐列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MusicOpenDetail []*MusicOpenDetail `json:"MusicOpenDetail,omitnil" name:"MusicOpenDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudMusicPurchasedResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudMusicPurchasedResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCloudMusicRequestParams struct {
	// 歌曲Id
	MusicId *string `json:"MusicId,omitnil" name:"MusicId"`

	// 歌曲类型，可选值有：
	// <li>MP3-128K-FTW：含有水印的试听资源；</li>
	// <li>MP3-320K-FTD-P：320kbps歌曲热门片段；</li>
	// <li>MP3-320K-FTD：320kbps已核验歌曲完整资源。</li>
	// 默认为：MP3-128K-FTW
	MusicType *string `json:"MusicType,omitnil" name:"MusicType"`
}

type DescribeCloudMusicRequest struct {
	*tchttp.BaseRequest
	
	// 歌曲Id
	MusicId *string `json:"MusicId,omitnil" name:"MusicId"`

	// 歌曲类型，可选值有：
	// <li>MP3-128K-FTW：含有水印的试听资源；</li>
	// <li>MP3-320K-FTD-P：320kbps歌曲热门片段；</li>
	// <li>MP3-320K-FTD：320kbps已核验歌曲完整资源。</li>
	// 默认为：MP3-128K-FTW
	MusicType *string `json:"MusicType,omitnil" name:"MusicType"`
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

// Predefined struct for user
type DescribeCloudMusicResponseParams struct {
	// 歌曲Id
	MusicId *string `json:"MusicId,omitnil" name:"MusicId"`

	// 歌曲名称
	MusicName *string `json:"MusicName,omitnil" name:"MusicName"`

	// 歌曲时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *int64 `json:"Duration,omitnil" name:"Duration"`

	// 歌曲链接
	MusicUrl *string `json:"MusicUrl,omitnil" name:"MusicUrl"`

	// 歌曲图片
	// 注意：此字段可能返回 null，表示取不到有效值。
	MusicImageUrl *string `json:"MusicImageUrl,omitnil" name:"MusicImageUrl"`

	// 歌手列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Singers []*string `json:"Singers,omitnil" name:"Singers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudMusicResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudMusicResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeItemByIdRequestParams struct {
	// 歌曲ID，目前暂不支持批量查询
	ItemIDs *string `json:"ItemIDs,omitnil" name:"ItemIDs"`
}

type DescribeItemByIdRequest struct {
	*tchttp.BaseRequest
	
	// 歌曲ID，目前暂不支持批量查询
	ItemIDs *string `json:"ItemIDs,omitnil" name:"ItemIDs"`
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

// Predefined struct for user
type DescribeItemByIdResponseParams struct {
	// 歌曲信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*Item `json:"Items,omitnil" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeItemByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeItemByIdResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeItemsRequestParams struct {
	// offset (Default = 0)，(当前页-1) * Limit
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 条数，必须大于0，最大值为30
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// （电台/歌单）ID，CategoryId和CategoryCode两个必传1个，可以从<a href="https://cloud.tencent.com/document/product/1155/40109">获取分类内容（Station）列表接口</a>中获取。
	CategoryId *string `json:"CategoryId,omitnil" name:"CategoryId"`

	// （电台/歌单）ID，CategoryId和CategoryCode两个必传1个，可以从<a href="https://cloud.tencent.com/document/product/1155/40109">获取分类内容（Station）列表接口</a>中获取。
	CategoryCode *string `json:"CategoryCode,omitnil" name:"CategoryCode"`
}

type DescribeItemsRequest struct {
	*tchttp.BaseRequest
	
	// offset (Default = 0)，(当前页-1) * Limit
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 条数，必须大于0，最大值为30
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// （电台/歌单）ID，CategoryId和CategoryCode两个必传1个，可以从<a href="https://cloud.tencent.com/document/product/1155/40109">获取分类内容（Station）列表接口</a>中获取。
	CategoryId *string `json:"CategoryId,omitnil" name:"CategoryId"`

	// （电台/歌单）ID，CategoryId和CategoryCode两个必传1个，可以从<a href="https://cloud.tencent.com/document/product/1155/40109">获取分类内容（Station）列表接口</a>中获取。
	CategoryCode *string `json:"CategoryCode,omitnil" name:"CategoryCode"`
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

// Predefined struct for user
type DescribeItemsResponseParams struct {
	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 当前页歌曲数量
	Size *uint64 `json:"Size,omitnil" name:"Size"`

	// 总数据条数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 剩余数量（total-offset-size），通过这个值判断是否
	// 还有下一页
	HaveMore *uint64 `json:"HaveMore,omitnil" name:"HaveMore"`

	// Items 歌曲列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*Item `json:"Items,omitnil" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeItemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeItemsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeKTVMusicDetailRequestParams struct {
	// 曲目 Id
	MusicId *string `json:"MusicId,omitnil" name:"MusicId"`
}

type DescribeKTVMusicDetailRequest struct {
	*tchttp.BaseRequest
	
	// 曲目 Id
	MusicId *string `json:"MusicId,omitnil" name:"MusicId"`
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

// Predefined struct for user
type DescribeKTVMusicDetailResponseParams struct {
	// 歌曲基础信息
	KTVMusicBaseInfo *KTVMusicBaseInfo `json:"KTVMusicBaseInfo,omitnil" name:"KTVMusicBaseInfo"`

	// 播放凭证
	PlayToken *string `json:"PlayToken,omitnil" name:"PlayToken"`

	// 歌词下载地址
	LyricsUrl *string `json:"LyricsUrl,omitnil" name:"LyricsUrl"`

	// 歌曲规格信息列表
	DefinitionInfoSet []*KTVMusicDefinitionInfo `json:"DefinitionInfoSet,omitnil" name:"DefinitionInfoSet"`

	// 音高数据文件下载地址
	MidiJsonUrl *string `json:"MidiJsonUrl,omitnil" name:"MidiJsonUrl"`

	// 副歌片段数据列表
	ChorusClipSet []*ChorusClip `json:"ChorusClipSet,omitnil" name:"ChorusClipSet"`

	// 前奏间隔，单位：毫秒；注：若参数返回为0则无人声部分
	PreludeInterval *int64 `json:"PreludeInterval,omitnil" name:"PreludeInterval"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeKTVMusicDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVMusicDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeKTVMusicTagsRequestParams struct {

}

type DescribeKTVMusicTagsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeKTVMusicTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVMusicTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVMusicTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVMusicTagsResponseParams struct {
	// 标签分组列表
	TagGroupSet []*KTVMusicTagGroup `json:"TagGroupSet,omitnil" name:"TagGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeKTVMusicTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVMusicTagsResponseParams `json:"Response"`
}

func (r *DescribeKTVMusicTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVMusicTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVPlaylistDetailRequestParams struct {
	// 歌单Id
	PlaylistId *string `json:"PlaylistId,omitnil" name:"PlaylistId"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeKTVPlaylistDetailRequest struct {
	*tchttp.BaseRequest
	
	// 歌单Id
	PlaylistId *string `json:"PlaylistId,omitnil" name:"PlaylistId"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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

// Predefined struct for user
type DescribeKTVPlaylistDetailResponseParams struct {
	// 歌曲基础信息列表
	KTVMusicInfoSet []*KTVMusicBaseInfo `json:"KTVMusicInfoSet,omitnil" name:"KTVMusicInfoSet"`

	// 歌单基础信息
	PlaylistBaseInfo *KTVPlaylistBaseInfo `json:"PlaylistBaseInfo,omitnil" name:"PlaylistBaseInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 歌单类型，取值有：
	// ·OfficialRec：官方推荐
	// ·Normal：自定义
	// 当该字段未填时，默认为取OfficialRec
	Type *string `json:"Type,omitnil" name:"Type"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	// 取值范围：Offset + Limit 不超过5000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	// 取值范围：Offset + Limit 不超过5000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeKTVPlaylistsRequest struct {
	*tchttp.BaseRequest
	
	// 歌单类型，取值有：
	// ·OfficialRec：官方推荐
	// ·Normal：自定义
	// 当该字段未填时，默认为取OfficialRec
	Type *string `json:"Type,omitnil" name:"Type"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	// 取值范围：Offset + Limit 不超过5000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	// 取值范围：Offset + Limit 不超过5000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVPlaylistsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVPlaylistsResponseParams struct {
	// 推荐歌单列表
	PlaylistBaseInfoSet []*KTVPlaylistBaseInfo `json:"PlaylistBaseInfoSet,omitnil" name:"PlaylistBaseInfoSet"`

	// 推荐歌单列表总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 机器人Id列表。
	RobotIds []*string `json:"RobotIds,omitnil" name:"RobotIds"`

	// 机器人状态，取值有：
	// <li>Play：播放</li>
	// <li>Pause：暂停</li>
	// <li>Destroy：销毁</li>
	Statuses []*string `json:"Statuses,omitnil" name:"Statuses"`

	// 匹配创建时间在此时间段内的机器人。
	// <li>包含所指定的头尾时间点。</li>
	CreateTime *TimeRange `json:"CreateTime,omitnil" name:"CreateTime"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页返回的起始偏移量，默认值：10。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeKTVRobotsRequest struct {
	*tchttp.BaseRequest
	
	// 机器人Id列表。
	RobotIds []*string `json:"RobotIds,omitnil" name:"RobotIds"`

	// 机器人状态，取值有：
	// <li>Play：播放</li>
	// <li>Pause：暂停</li>
	// <li>Destroy：销毁</li>
	Statuses []*string `json:"Statuses,omitnil" name:"Statuses"`

	// 匹配创建时间在此时间段内的机器人。
	// <li>包含所指定的头尾时间点。</li>
	CreateTime *TimeRange `json:"CreateTime,omitnil" name:"CreateTime"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页返回的起始偏移量，默认值：10。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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

// Predefined struct for user
type DescribeKTVRobotsResponseParams struct {
	// 机器人总数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 机器人信息集合。
	KTVRobotInfoSet []*KTVRobotInfo `json:"KTVRobotInfoSet,omitnil" name:"KTVRobotInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeKTVSingerCategoriesRequestParams struct {

}

type DescribeKTVSingerCategoriesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeKTVSingerCategoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVSingerCategoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVSingerCategoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVSingerCategoriesResponseParams struct {
	// 歌手性别分类列表
	GenderSet []*KTVSingerCategoryInfo `json:"GenderSet,omitnil" name:"GenderSet"`

	// 歌手区域分类列表
	AreaSet []*KTVSingerCategoryInfo `json:"AreaSet,omitnil" name:"AreaSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeKTVSingerCategoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVSingerCategoriesResponseParams `json:"Response"`
}

func (r *DescribeKTVSingerCategoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVSingerCategoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVSingerMusicsRequestParams struct {
	// 歌手id
	SingerId *string `json:"SingerId,omitnil" name:"SingerId"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeKTVSingerMusicsRequest struct {
	*tchttp.BaseRequest
	
	// 歌手id
	SingerId *string `json:"SingerId,omitnil" name:"SingerId"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeKTVSingerMusicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVSingerMusicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SingerId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVSingerMusicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVSingerMusicsResponseParams struct {
	// 总曲目数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// KTV 曲目列表
	KTVMusicInfoSet []*KTVMusicBaseInfo `json:"KTVMusicInfoSet,omitnil" name:"KTVMusicInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeKTVSingerMusicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVSingerMusicsResponseParams `json:"Response"`
}

func (r *DescribeKTVSingerMusicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVSingerMusicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVSingersRequestParams struct {
	// 歌手id集合，精确匹配歌手id
	// <li> 数组长度限制10</li>
	SingerIds []*string `json:"SingerIds,omitnil" name:"SingerIds"`

	// 歌手性别集合，不传为全部，精确匹配歌手性别类型，
	// <li>数组长度限制1</li>
	// <li>取值范围：直播互动曲库歌手分类信息接口，返回性别分类信息列表中，分类英文名</li>
	Genders []*string `json:"Genders,omitnil" name:"Genders"`

	// 歌手区域集合，不传为全部，精确匹配歌手区域
	// <li>数组长度限制10</li>
	// <li>取值范围：直播互动曲库歌手分类信息接口，返回的区域分类信息列表中，分类英文名</li>
	Areas []*string `json:"Areas,omitnil" name:"Areas"`

	// 排序方式。默认按照播放数倒序
	// <li> Sort.Field 可选 PlayCount。</li>
	Sort *SortBy `json:"Sort,omitnil" name:"Sort"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeKTVSingersRequest struct {
	*tchttp.BaseRequest
	
	// 歌手id集合，精确匹配歌手id
	// <li> 数组长度限制10</li>
	SingerIds []*string `json:"SingerIds,omitnil" name:"SingerIds"`

	// 歌手性别集合，不传为全部，精确匹配歌手性别类型，
	// <li>数组长度限制1</li>
	// <li>取值范围：直播互动曲库歌手分类信息接口，返回性别分类信息列表中，分类英文名</li>
	Genders []*string `json:"Genders,omitnil" name:"Genders"`

	// 歌手区域集合，不传为全部，精确匹配歌手区域
	// <li>数组长度限制10</li>
	// <li>取值范围：直播互动曲库歌手分类信息接口，返回的区域分类信息列表中，分类英文名</li>
	Areas []*string `json:"Areas,omitnil" name:"Areas"`

	// 排序方式。默认按照播放数倒序
	// <li> Sort.Field 可选 PlayCount。</li>
	Sort *SortBy `json:"Sort,omitnil" name:"Sort"`

	// 分页偏移量，默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeKTVSingersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVSingersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SingerIds")
	delete(f, "Genders")
	delete(f, "Areas")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVSingersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVSingersResponseParams struct {
	// 总歌手数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// KTV歌手列表
	KTVSingerInfoSet []*KTVSingerInfo `json:"KTVSingerInfoSet,omitnil" name:"KTVSingerInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeKTVSingersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVSingersResponseParams `json:"Response"`
}

func (r *DescribeKTVSingersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVSingersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVSuggestionsRequestParams struct {
	// 联想关键词
	KeyWord *string `json:"KeyWord,omitnil" name:"KeyWord"`
}

type DescribeKTVSuggestionsRequest struct {
	*tchttp.BaseRequest
	
	// 联想关键词
	KeyWord *string `json:"KeyWord,omitnil" name:"KeyWord"`
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
	delete(f, "KeyWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVSuggestionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVSuggestionsResponseParams struct {
	// 联想词信息列表。返回总数最大为10。
	KTVSuggestionInfoSet []*KTVSuggestionInfo `json:"KTVSuggestionInfoSet,omitnil" name:"KTVSuggestionInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeKTVTopListRequestParams struct {
	// 榜单类型。默认Hot
	// <li> Hot, 热歌榜。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 榜单周期 默认为Week
	// <li> Week, 周榜。</li>
	// <li> Month, 月榜。</li>
	Period *string `json:"Period,omitnil" name:"Period"`
}

type DescribeKTVTopListRequest struct {
	*tchttp.BaseRequest
	
	// 榜单类型。默认Hot
	// <li> Hot, 热歌榜。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 榜单周期 默认为Week
	// <li> Week, 周榜。</li>
	// <li> Month, 月榜。</li>
	Period *string `json:"Period,omitnil" name:"Period"`
}

func (r *DescribeKTVTopListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVTopListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKTVTopListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKTVTopListResponseParams struct {
	// 歌曲基础信息列表
	KTVMusicTopInfoSet []*KTVMusicTopInfo `json:"KTVMusicTopInfoSet,omitnil" name:"KTVMusicTopInfoSet"`

	// 返回总条数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeKTVTopListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKTVTopListResponseParams `json:"Response"`
}

func (r *DescribeKTVTopListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKTVTopListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLyricRequestParams struct {
	// 歌曲ID
	ItemId *string `json:"ItemId,omitnil" name:"ItemId"`

	// 格式，可选项，可不填写，默认值为：LRC-LRC。
	// <li>LRC-LRC：歌词；</li>
	// <li>JSON-ST：波形图。</li>
	SubItemType *string `json:"SubItemType,omitnil" name:"SubItemType"`
}

type DescribeLyricRequest struct {
	*tchttp.BaseRequest
	
	// 歌曲ID
	ItemId *string `json:"ItemId,omitnil" name:"ItemId"`

	// 格式，可选项，可不填写，默认值为：LRC-LRC。
	// <li>LRC-LRC：歌词；</li>
	// <li>JSON-ST：波形图。</li>
	SubItemType *string `json:"SubItemType,omitnil" name:"SubItemType"`
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

// Predefined struct for user
type DescribeLyricResponseParams struct {
	// 歌词或者波形图详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lyric *Lyric `json:"Lyric,omitnil" name:"Lyric"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLyricResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLyricResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeMusicRequestParams struct {
	// 歌曲ID
	ItemId *string `json:"ItemId,omitnil" name:"ItemId"`

	// 在应用前端播放音乐C端用户的唯一标识。无需是账户信息，用户唯一标识即可。
	IdentityId *string `json:"IdentityId,omitnil" name:"IdentityId"`

	// MP3-320K-FTD-P  为获取320kbps歌曲热门片段。
	// MP3-320K-FTD 为获取320kbps已核验歌曲完整资源。
	SubItemType *string `json:"SubItemType,omitnil" name:"SubItemType"`

	// CDN URL Protocol:HTTP or HTTPS/SSL
	// Values:Y , N(default)
	Ssl *string `json:"Ssl,omitnil" name:"Ssl"`
}

type DescribeMusicRequest struct {
	*tchttp.BaseRequest
	
	// 歌曲ID
	ItemId *string `json:"ItemId,omitnil" name:"ItemId"`

	// 在应用前端播放音乐C端用户的唯一标识。无需是账户信息，用户唯一标识即可。
	IdentityId *string `json:"IdentityId,omitnil" name:"IdentityId"`

	// MP3-320K-FTD-P  为获取320kbps歌曲热门片段。
	// MP3-320K-FTD 为获取320kbps已核验歌曲完整资源。
	SubItemType *string `json:"SubItemType,omitnil" name:"SubItemType"`

	// CDN URL Protocol:HTTP or HTTPS/SSL
	// Values:Y , N(default)
	Ssl *string `json:"Ssl,omitnil" name:"Ssl"`
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

// Predefined struct for user
type DescribeMusicResponseParams struct {
	// 音乐相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Music *Music `json:"Music,omitnil" name:"Music"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMusicResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMusicResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeMusicSaleStatusRequestParams struct {
	// 歌曲Id集合，可传单个，也可传多个，上线查询单次50个
	MusicIds []*string `json:"MusicIds,omitnil" name:"MusicIds"`

	// 查询哪个渠道的数据，1为曲库包，2为单曲
	PurchaseType *int64 `json:"PurchaseType,omitnil" name:"PurchaseType"`
}

type DescribeMusicSaleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 歌曲Id集合，可传单个，也可传多个，上线查询单次50个
	MusicIds []*string `json:"MusicIds,omitnil" name:"MusicIds"`

	// 查询哪个渠道的数据，1为曲库包，2为单曲
	PurchaseType *int64 `json:"PurchaseType,omitnil" name:"PurchaseType"`
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

// Predefined struct for user
type DescribeMusicSaleStatusResponseParams struct {
	// musicId对应歌曲状态
	MusicStatusSet []*MusicStatus `json:"MusicStatusSet,omitnil" name:"MusicStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMusicSaleStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMusicSaleStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePackageItemsRequestParams struct {
	// 订单id，从获取已购曲库包列表中获取
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 默认0，Offset=Offset+Length
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 默认20
	Length *uint64 `json:"Length,omitnil" name:"Length"`
}

type DescribePackageItemsRequest struct {
	*tchttp.BaseRequest
	
	// 订单id，从获取已购曲库包列表中获取
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 默认0，Offset=Offset+Length
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 默认20
	Length *uint64 `json:"Length,omitnil" name:"Length"`
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

// Predefined struct for user
type DescribePackageItemsResponseParams struct {
	// 已核销歌曲信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageItems []*PackageItem `json:"PackageItems,omitnil" name:"PackageItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePackageItemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePackageItemsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePackagesRequestParams struct {
	// 默认0，Offset=Offset+Length
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 默认20
	Length *uint64 `json:"Length,omitnil" name:"Length"`
}

type DescribePackagesRequest struct {
	*tchttp.BaseRequest
	
	// 默认0，Offset=Offset+Length
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 默认20
	Length *uint64 `json:"Length,omitnil" name:"Length"`
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

// Predefined struct for user
type DescribePackagesResponseParams struct {
	// 已购曲库包列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Packages []*Package `json:"Packages,omitnil" name:"Packages"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePackagesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePkgOfflineMusicRequestParams struct {
	// 订单id
	PackageOrderId *string `json:"PackageOrderId,omitnil" name:"PackageOrderId"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条(注：单次上限为100)。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribePkgOfflineMusicRequest struct {
	*tchttp.BaseRequest
	
	// 订单id
	PackageOrderId *string `json:"PackageOrderId,omitnil" name:"PackageOrderId"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条(注：单次上限为100)。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页返回的记录条数，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
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

// Predefined struct for user
type DescribePkgOfflineMusicResponseParams struct {
	// 曲库包中不可用歌曲信息
	OfflineMusicSet []*OfflineMusicDetail `json:"OfflineMusicSet,omitnil" name:"OfflineMusicSet"`

	// 返回总量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePkgOfflineMusicResponse struct {
	*tchttp.BaseResponse
	Response *DescribePkgOfflineMusicResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeStationsRequestParams struct {
	// 条数，必须大于0
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// offset (Default = 0)，Offset=Offset+Limit
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeStationsRequest struct {
	*tchttp.BaseRequest
	
	// 条数，必须大于0
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// offset (Default = 0)，Offset=Offset+Limit
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
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

// Predefined struct for user
type DescribeStationsResponseParams struct {
	// 总数量
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 当前页station数量
	Size *uint64 `json:"Size,omitnil" name:"Size"`

	// 剩余数量（total-offset-size），通过这个值判断是否还有下一页
	HaveMore *uint64 `json:"HaveMore,omitnil" name:"HaveMore"`

	// Stations 素材库列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stations []*Station `json:"Stations,omitnil" name:"Stations"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStationsResponseParams `json:"Response"`
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

// Predefined struct for user
type DestroyKTVRobotRequestParams struct {
	// 机器人Id。
	RobotId *string `json:"RobotId,omitnil" name:"RobotId"`
}

type DestroyKTVRobotRequest struct {
	*tchttp.BaseRequest
	
	// 机器人Id。
	RobotId *string `json:"RobotId,omitnil" name:"RobotId"`
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

// Predefined struct for user
type DestroyKTVRobotResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type ImagePath struct {
	// station图片大小及类别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// station图片地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Item struct {
	// Song ID
	ItemID *string `json:"ItemID,omitnil" name:"ItemID"`

	// Song info
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataInfo *DataInfo `json:"DataInfo,omitnil" name:"DataInfo"`

	// 专辑信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Album *Album `json:"Album,omitnil" name:"Album"`

	// 多个歌手集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Artists []*Artist `json:"Artists,omitnil" name:"Artists"`

	// 歌曲状态，1:添加进购物车；2:核销进曲库包
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type JoinRoomInput struct {
	// TRTC进房参数
	TRTCJoinRoomInput *TRTCJoinRoomInput `json:"TRTCJoinRoomInput,omitnil" name:"TRTCJoinRoomInput"`
}

type KTVMusicBaseInfo struct {
	// 歌曲 Id
	MusicId *string `json:"MusicId,omitnil" name:"MusicId"`

	// 歌曲名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 演唱者基础信息列表
	SingerInfoSet []*KTVSingerBaseInfo `json:"SingerInfoSet,omitnil" name:"SingerInfoSet"`

	// 已弃用，请使用SingerInfoSet
	SingerSet []*string `json:"SingerSet,omitnil" name:"SingerSet"`

	// 作词者列表
	LyricistSet []*string `json:"LyricistSet,omitnil" name:"LyricistSet"`

	// 作曲者列表
	ComposerSet []*string `json:"ComposerSet,omitnil" name:"ComposerSet"`

	// 标签列表
	TagSet []*string `json:"TagSet,omitnil" name:"TagSet"`

	// 歌曲时长
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`
}

type KTVMusicDefinitionInfo struct {
	// 规格，取值有：
	// <li>audio/mi：低规格；</li>
	// <li>audio/lo：中规格；</li>
	// <li>audio/hi：高规格。</li>
	Definition *string `json:"Definition,omitnil" name:"Definition"`

	// 码率，单位为 bps。
	Bitrate *int64 `json:"Bitrate,omitnil" name:"Bitrate"`

	// 文件大小，单位为字节。
	Size *int64 `json:"Size,omitnil" name:"Size"`
}

type KTVMusicDetailInfo struct {
	// 即使广播曲库歌曲基础信息
	KTVMusicBaseInfo *KTVMusicBaseInfo `json:"KTVMusicBaseInfo,omitnil" name:"KTVMusicBaseInfo"`

	// 播放凭证
	PlayToken *string `json:"PlayToken,omitnil" name:"PlayToken"`

	// 歌词下载地址
	LyricsUrl *string `json:"LyricsUrl,omitnil" name:"LyricsUrl"`

	// 歌曲规格信息列表
	DefinitionInfoSet []*KTVMusicDefinitionInfo `json:"DefinitionInfoSet,omitnil" name:"DefinitionInfoSet"`

	// 音高数据文件下载地址
	MidiJsonUrl *string `json:"MidiJsonUrl,omitnil" name:"MidiJsonUrl"`

	// 副歌片段数据列表
	ChorusClipSet []*ChorusClip `json:"ChorusClipSet,omitnil" name:"ChorusClipSet"`

	// 前奏间隔，单位：毫秒；注：若参数返回为0则无人声部分
	PreludeInterval *int64 `json:"PreludeInterval,omitnil" name:"PreludeInterval"`
}

type KTVMusicTagGroup struct {
	// 标签分组英文名
	EnglishGroupName *string `json:"EnglishGroupName,omitnil" name:"EnglishGroupName"`

	// 标签分组中文名
	ChineseGroupName *string `json:"ChineseGroupName,omitnil" name:"ChineseGroupName"`

	// 标签分类下标签列表
	TagSet []*KTVMusicTagInfo `json:"TagSet,omitnil" name:"TagSet"`
}

type KTVMusicTagInfo struct {
	// 标签Id
	TagId *string `json:"TagId,omitnil" name:"TagId"`

	// 标签
	TagName *string `json:"TagName,omitnil" name:"TagName"`
}

type KTVMusicTopInfo struct {
	// 歌曲Id
	MusicId *string `json:"MusicId,omitnil" name:"MusicId"`

	// 歌曲名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 歌手名称列表
	SingerInfoSet []*KTVSingerBaseInfo `json:"SingerInfoSet,omitnil" name:"SingerInfoSet"`

	// 歌词名称列表
	LyricistSet []*string `json:"LyricistSet,omitnil" name:"LyricistSet"`

	// 作曲列表
	ComposerSet []*string `json:"ComposerSet,omitnil" name:"ComposerSet"`

	// 标签列表
	TagSet []*string `json:"TagSet,omitnil" name:"TagSet"`

	// 播放时长
	Duration *int64 `json:"Duration,omitnil" name:"Duration"`
}

type KTVPlaylistBaseInfo struct {
	// 歌单Id
	PlaylistId *string `json:"PlaylistId,omitnil" name:"PlaylistId"`

	// 歌单标题
	Title *string `json:"Title,omitnil" name:"Title"`

	// 歌单介绍
	Description *string `json:"Description,omitnil" name:"Description"`

	// 歌曲数量
	MusicNum *int64 `json:"MusicNum,omitnil" name:"MusicNum"`
}

type KTVRobotInfo struct {
	// 机器人Id。
	RobotId *string `json:"RobotId,omitnil" name:"RobotId"`

	// 状态，取值有：
	// <li>Play：播放</li>
	// <li>Pause：暂停</li>
	// <li>Destroy：销毁</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 播放列表。
	Playlists []*string `json:"Playlists,omitnil" name:"Playlists"`

	// 当前歌单索引位置。
	CurIndex *int64 `json:"CurIndex,omitnil" name:"CurIndex"`

	// 播放进度，单位：毫秒。
	Position *uint64 `json:"Position,omitnil" name:"Position"`

	// 音频参数。
	SetAudioParamInput *SetAudioParamCommandInput `json:"SetAudioParamInput,omitnil" name:"SetAudioParamInput"`

	// 进房信息。
	JoinRoomInput *JoinRoomInput `json:"JoinRoomInput,omitnil" name:"JoinRoomInput"`

	// RTC厂商类型，取值有：
	// <li>TRTC</li>
	RTCSystem *string `json:"RTCSystem,omitnil" name:"RTCSystem"`

	// 播放模式，PlayMode取值有：
	// <li>RepeatPlaylist：列表循环</li>
	// <li>Order：顺序播放</li>
	// <li>RepeatSingle：单曲循环</li>
	// <li>Shuffle：随机播放</li>
	SetPlayModeInput *SetPlayModeCommandInput `json:"SetPlayModeInput,omitnil" name:"SetPlayModeInput"`

	// <del>音量，范围 0~100，默认为 50。</del>（已废弃，请采用 SetRealVolumeInput ）
	SetVolumeInput *SetVolumeCommandInput `json:"SetVolumeInput,omitnil" name:"SetVolumeInput"`

	// 真实音量，范围 0~100，默认为 50。
	SetRealVolumeInput *SetRealVolumeCommandInput `json:"SetRealVolumeInput,omitnil" name:"SetRealVolumeInput"`
}

type KTVSingerBaseInfo struct {
	// 歌手id
	SingerId *string `json:"SingerId,omitnil" name:"SingerId"`

	// 歌手名
	Name *string `json:"Name,omitnil" name:"Name"`
}

type KTVSingerCategoryInfo struct {
	// 分类中文名
	ChineseName *string `json:"ChineseName,omitnil" name:"ChineseName"`

	// 分类英文名
	EnglishName *string `json:"EnglishName,omitnil" name:"EnglishName"`
}

type KTVSingerInfo struct {
	// 歌手id
	SingerId *string `json:"SingerId,omitnil" name:"SingerId"`

	// 歌手名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 歌手性别: 男，女，组合
	Gender *string `json:"Gender,omitnil" name:"Gender"`

	// 地区: 大陆，港台，欧美，日本
	Area *string `json:"Area,omitnil" name:"Area"`

	// 歌曲数
	MusicCount *int64 `json:"MusicCount,omitnil" name:"MusicCount"`

	// 歌曲总播放次数
	PlayCount *int64 `json:"PlayCount,omitnil" name:"PlayCount"`
}

type KTVSuggestionInfo struct {
	// 联想词
	Suggestion *string `json:"Suggestion,omitnil" name:"Suggestion"`
}

type Lyric struct {
	// 歌词cdn地址
	Url *string `json:"Url,omitnil" name:"Url"`

	// 歌词后缀名
	FileNameExt *string `json:"FileNameExt,omitnil" name:"FileNameExt"`

	// 歌词类型
	SubItemType *string `json:"SubItemType,omitnil" name:"SubItemType"`
}

// Predefined struct for user
type ModifyMusicOnShelvesRequestParams struct {
	// 歌曲变更信息
	MusicDetailInfos *MusicDetailInfo `json:"MusicDetailInfos,omitnil" name:"MusicDetailInfos"`

	// ame对接资源方密钥
	AmeKey *string `json:"AmeKey,omitnil" name:"AmeKey"`
}

type ModifyMusicOnShelvesRequest struct {
	*tchttp.BaseRequest
	
	// 歌曲变更信息
	MusicDetailInfos *MusicDetailInfo `json:"MusicDetailInfos,omitnil" name:"MusicDetailInfos"`

	// ame对接资源方密钥
	AmeKey *string `json:"AmeKey,omitnil" name:"AmeKey"`
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

// Predefined struct for user
type ModifyMusicOnShelvesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyMusicOnShelvesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMusicOnShelvesResponseParams `json:"Response"`
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
	Url *string `json:"Url,omitnil" name:"Url"`

	// 音频文件大小
	FileSize *uint64 `json:"FileSize,omitnil" name:"FileSize"`

	// 音频文件类型
	FileExtension *string `json:"FileExtension,omitnil" name:"FileExtension"`

	// Song fragment start.试听片段开始时间，试听时长为auditionEnd-auditionBegin
	// Unit :ms
	AuditionBegin *uint64 `json:"AuditionBegin,omitnil" name:"AuditionBegin"`

	// Song fragment end.试听片段结束时间, 试听时长为auditionEnd-auditionBegin
	// Unit :ms
	AuditionEnd *uint64 `json:"AuditionEnd,omitnil" name:"AuditionEnd"`

	// 音乐播放链接全路径，前提是在正版曲库直通车控制台添加过域名，否则返回空字符。
	// 如果添加过多个域名只返回第一个添加域名的播放全路径。
	FullUrl *string `json:"FullUrl,omitnil" name:"FullUrl"`
}

type MusicDetailInfo struct {
	// 资源方音乐Id
	MusicId *string `json:"MusicId,omitnil" name:"MusicId"`

	// 资源方识别信息
	AmeId *string `json:"AmeId,omitnil" name:"AmeId"`

	// 分类标签
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 关键词
	HitWords []*string `json:"HitWords,omitnil" name:"HitWords"`

	// 节奏信息
	Bpm *int64 `json:"Bpm,omitnil" name:"Bpm"`

	// 商业化权益
	Score *float64 `json:"Score,omitnil" name:"Score"`

	// 应用歌曲信息,1.图文/短视频,2.网络直播,3.网络电台FM,4.免费游戏,5.商业游戏,6.网店网站设计,7.广告营销,8.网络长视频
	Scene []*string `json:"Scene,omitnil" name:"Scene"`

	// 应用地域,1. 中国大陆,2. 中国含港澳台,3. 全球
	Region []*string `json:"Region,omitnil" name:"Region"`

	// 授权时间,1. 1年, 5. 随片永久
	AuthPeriod *string `json:"AuthPeriod,omitnil" name:"AuthPeriod"`

	// 商业化授权，1. 支持商业化 ,2. 不支持商业化
	Commercialization *string `json:"Commercialization,omitnil" name:"Commercialization"`

	// 跨平台传播，1. 支持跨平台传播 ,2. 不支持跨平台传播
	Platform *string `json:"Platform,omitnil" name:"Platform"`

	// 传播渠道
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

type MusicOpenDetail struct {
	// 音乐Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MusicId *string `json:"MusicId,omitnil" name:"MusicId"`

	// 专辑名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbumName *string `json:"AlbumName,omitnil" name:"AlbumName"`

	// 专辑图片路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbumImageUrl *string `json:"AlbumImageUrl,omitnil" name:"AlbumImageUrl"`

	// 音乐名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MusicName *string `json:"MusicName,omitnil" name:"MusicName"`

	// 音乐图片路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	MusicImageUrl *string `json:"MusicImageUrl,omitnil" name:"MusicImageUrl"`

	// 歌手
	// 注意：此字段可能返回 null，表示取不到有效值。
	Singers []*string `json:"Singers,omitnil" name:"Singers"`

	// 播放时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 歌词url
	// 注意：此字段可能返回 null，表示取不到有效值。
	LyricUrl *string `json:"LyricUrl,omitnil" name:"LyricUrl"`

	// 波形图url
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaveformUrl *string `json:"WaveformUrl,omitnil" name:"WaveformUrl"`
}

type MusicStatus struct {
	// 歌曲Id
	MusicId *string `json:"MusicId,omitnil" name:"MusicId"`

	// 在售状态,0为在售，1为临时下架，2为永久下架
	SaleStatus *int64 `json:"SaleStatus,omitnil" name:"SaleStatus"`
}

type OfflineMusicDetail struct {
	// 歌曲Id
	ItemId *string `json:"ItemId,omitnil" name:"ItemId"`

	// 歌曲名称
	MusicName *string `json:"MusicName,omitnil" name:"MusicName"`

	// 不可用原因
	OffRemark *string `json:"OffRemark,omitnil" name:"OffRemark"`

	// 不可用时间
	OffTime *string `json:"OffTime,omitnil" name:"OffTime"`
}

type Package struct {
	// 订单id
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 曲库包名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 授权地区-global: 全球  CN: 中国
	AuthorizedArea *string `json:"AuthorizedArea,omitnil" name:"AuthorizedArea"`

	// 授权次数
	AuthorizedLimit *int64 `json:"AuthorizedLimit,omitnil" name:"AuthorizedLimit"`

	// 套餐有效期，单位:天
	TermOfValidity *int64 `json:"TermOfValidity,omitnil" name:"TermOfValidity"`

	// 0:不可商业化；1:可商业化
	Commercial *int64 `json:"Commercial,omitnil" name:"Commercial"`

	// 套餐价格，单位：元
	PackagePrice *float64 `json:"PackagePrice,omitnil" name:"PackagePrice"`

	// 生效开始时间,格式yyyy-MM-dd HH:mm:ss
	EffectTime *string `json:"EffectTime,omitnil" name:"EffectTime"`

	// 生效结束时间,格式yyyy-MM-dd HH:mm:ss
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 剩余授权次数
	UsedCount *int64 `json:"UsedCount,omitnil" name:"UsedCount"`

	// 曲库包用途信息
	UseRanges []*UseRange `json:"UseRanges,omitnil" name:"UseRanges"`
}

type PackageItem struct {
	// 订单id
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 歌曲名
	TrackName *string `json:"TrackName,omitnil" name:"TrackName"`

	// 歌曲ID
	ItemID *string `json:"ItemID,omitnil" name:"ItemID"`

	// 专辑图片
	Img *string `json:"Img,omitnil" name:"Img"`

	// 歌手名
	ArtistName *string `json:"ArtistName,omitnil" name:"ArtistName"`

	// 歌曲时长
	Duration *string `json:"Duration,omitnil" name:"Duration"`

	// 授权区域，global: 全球 CN: 中国
	AuthorizedArea *string `json:"AuthorizedArea,omitnil" name:"AuthorizedArea"`

	// 标签数组
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

type PlayCommandInput struct {
	// 歌曲位置索引。
	Index *int64 `json:"Index,omitnil" name:"Index"`
}

// Predefined struct for user
type PutMusicOnTheShelvesRequestParams struct {
	// 资源方歌曲Id
	MusicIds []*string `json:"MusicIds,omitnil" name:"MusicIds"`
}

type PutMusicOnTheShelvesRequest struct {
	*tchttp.BaseRequest
	
	// 资源方歌曲Id
	MusicIds []*string `json:"MusicIds,omitnil" name:"MusicIds"`
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

// Predefined struct for user
type PutMusicOnTheShelvesResponseParams struct {
	// 操作成功数量
	SuccessNum *int64 `json:"SuccessNum,omitnil" name:"SuccessNum"`

	// 操作失败数量
	FailedNum *int64 `json:"FailedNum,omitnil" name:"FailedNum"`

	// 失败歌曲Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedMusicIds []*string `json:"FailedMusicIds,omitnil" name:"FailedMusicIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PutMusicOnTheShelvesResponse struct {
	*tchttp.BaseResponse
	Response *PutMusicOnTheShelvesResponseParams `json:"Response"`
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

// Predefined struct for user
type ReportDataRequestParams struct {
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
	ReportData *string `json:"ReportData,omitnil" name:"ReportData"`
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
	ReportData *string `json:"ReportData,omitnil" name:"ReportData"`
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

// Predefined struct for user
type ReportDataResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReportDataResponse struct {
	*tchttp.BaseResponse
	Response *ReportDataResponseParams `json:"Response"`
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

// Predefined struct for user
type SearchKTVMusicsRequestParams struct {
	// 搜索关键词
	KeyWord *string `json:"KeyWord,omitnil" name:"KeyWord"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	// 取值范围：Offset + Limit 不超过5000。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页返回的起始偏移量，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式。默认按照匹配度排序
	// <li> Sort.Field 可选 CreateTime</li>
	// <li> Sort.Order 可选 Desc </li>
	// <li> 当 KeyWord 不为空时，Sort.Field 字段无效， 搜索结果将以匹配度排序。</li>
	Sort *SortBy `json:"Sort,omitnil" name:"Sort"`

	// 标签 ID 集合，匹配集合指定所有 ID 。
	// <li>数组长度限制：10。</li>
	TagIds []*string `json:"TagIds,omitnil" name:"TagIds"`
}

type SearchKTVMusicsRequest struct {
	*tchttp.BaseRequest
	
	// 搜索关键词
	KeyWord *string `json:"KeyWord,omitnil" name:"KeyWord"`

	// 分页返回的起始偏移量，默认值：0。将返回第 Offset 到第 Offset+Limit-1 条。
	// 取值范围：Offset + Limit 不超过5000。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页返回的起始偏移量，默认值：50。将返回第 Offset 到第 Offset+Limit-1 条。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式。默认按照匹配度排序
	// <li> Sort.Field 可选 CreateTime</li>
	// <li> Sort.Order 可选 Desc </li>
	// <li> 当 KeyWord 不为空时，Sort.Field 字段无效， 搜索结果将以匹配度排序。</li>
	Sort *SortBy `json:"Sort,omitnil" name:"Sort"`

	// 标签 ID 集合，匹配集合指定所有 ID 。
	// <li>数组长度限制：10。</li>
	TagIds []*string `json:"TagIds,omitnil" name:"TagIds"`
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
	delete(f, "Sort")
	delete(f, "TagIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchKTVMusicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchKTVMusicsResponseParams struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// KTV 曲目列表
	KTVMusicInfoSet []*KTVMusicBaseInfo `json:"KTVMusicInfoSet,omitnil" name:"KTVMusicInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Position *uint64 `json:"Position,omitnil" name:"Position"`
}

type SendMessageCommandInput struct {
	// 自定义消息，json格式字符串。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 消息重复次数，默认为 1。
	Repeat *uint64 `json:"Repeat,omitnil" name:"Repeat"`
}

type SetAudioParamCommandInput struct {
	// 规格，取值有：
	// <li>audio/mi：低规格</li>
	// <li>audio/lo：中规格</li>
	// <li>audio/hi：高规格</li>
	Definition *string `json:"Definition,omitnil" name:"Definition"`

	// 音频类型，取值有：
	// <li>Original：原唱</li>
	// <li>Accompaniment：伴奏</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

type SetDestroyModeCommandInput struct {
	// 销毁模式，取值有：
	// <li>Auto：房间没人时自动销毁</li>
	// <li>Expire：房间没人时过期自动销毁</li>
	// <li>Never：不自动销毁，需手动销毁</li>默认为：Auto。
	DestroyMode *string `json:"DestroyMode,omitnil" name:"DestroyMode"`

	// 过期销毁时间，单位：秒，当DestroyMode取Expire时必填。
	DestroyExpireTime *int64 `json:"DestroyExpireTime,omitnil" name:"DestroyExpireTime"`
}

type SetPlayModeCommandInput struct {
	// 播放模式，取值有：
	// <li>RepeatPlaylist：列表循环</li>
	// <li>Order：顺序播放</li>
	// <li>RepeatSingle：单曲循环</li>
	// <li>Shuffle：随机播放</li>
	PlayMode *string `json:"PlayMode,omitnil" name:"PlayMode"`
}

type SetPlaylistCommandInput struct {
	// 变更类型，取值有：
	// <li>Add：添加</li>
	// <li>Delete：删除</li>
	// <li>ClearList：清空歌曲列表</li>
	// <li>Move：移动歌曲</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 歌单索引位置，
	// 当 Type 取 Add 时，-1表示添加在列表最后位置，大于-1表示要添加的位置；
	// 当 Type 取 Delete 时，表示待删除歌曲的位置；
	// 当 Type 取 Move 时，表示待调整歌曲的位置。
	Index *int64 `json:"Index,omitnil" name:"Index"`

	// 当 Type 取 Move 时，必填，表示移动歌曲的目标位置。
	ChangedIndex *int64 `json:"ChangedIndex,omitnil" name:"ChangedIndex"`

	// 歌曲 ID 列表，当 Type 取 Add 时，与MusicURLs必填其中一项。
	MusicIds []*string `json:"MusicIds,omitnil" name:"MusicIds"`

	// 歌曲 URL 列表，当 Type 取 Add 时，与MusicIds必填其中一项。
	// 注：URL必须以.mp3结尾且必须是mp3编码文件。
	MusicURLs []*string `json:"MusicURLs,omitnil" name:"MusicURLs"`
}

type SetRealVolumeCommandInput struct {
	// 真实音量大小，取值范围为 0~100，默认值为 50。
	RealVolume *int64 `json:"RealVolume,omitnil" name:"RealVolume"`
}

type SetVolumeCommandInput struct {
	// 音量大小，取值范围为 0~100，默认值为 50。
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`
}

type SortBy struct {
	// 排序字段
	Field *string `json:"Field,omitnil" name:"Field"`

	// 排序方式，可选值：Asc（升序）、Desc（降序）
	Order *string `json:"Order,omitnil" name:"Order"`
}

type Station struct {
	// StationID
	CategoryID *string `json:"CategoryID,omitnil" name:"CategoryID"`

	// Station MCCode
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryCode *string `json:"CategoryCode,omitnil" name:"CategoryCode"`

	// Category Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// Station的排序值，供参考（返回结果已按其升序）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rank *int64 `json:"Rank,omitnil" name:"Rank"`

	// station图片集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImagePathMap []*ImagePath `json:"ImagePathMap,omitnil" name:"ImagePathMap"`
}

// Predefined struct for user
type SyncKTVRobotCommandRequestParams struct {
	// 机器人Id。
	RobotId *string `json:"RobotId,omitnil" name:"RobotId"`

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
	// <li>SetDestroyMode：设置销毁模式</li>
	// <li><del>SetVolume：设置音量</del>（已废弃，请采用 SetRealVolume）</li>
	// <li>SetRealVolume：设置真实音量</li>
	Command *string `json:"Command,omitnil" name:"Command"`

	// 播放参数。
	PlayCommandInput *PlayCommandInput `json:"PlayCommandInput,omitnil" name:"PlayCommandInput"`

	// 播放列表变更信息，当Command取SetPlaylist时，必填。
	SetPlaylistCommandInput *SetPlaylistCommandInput `json:"SetPlaylistCommandInput,omitnil" name:"SetPlaylistCommandInput"`

	// 播放进度，当Command取Seek时，必填。
	SeekCommandInput *SeekCommandInput `json:"SeekCommandInput,omitnil" name:"SeekCommandInput"`

	// 音频参数，当Command取SetAudioParam时，必填。
	SetAudioParamCommandInput *SetAudioParamCommandInput `json:"SetAudioParamCommandInput,omitnil" name:"SetAudioParamCommandInput"`

	// 自定义消息，当Command取SendMessage时，必填。
	SendMessageCommandInput *SendMessageCommandInput `json:"SendMessageCommandInput,omitnil" name:"SendMessageCommandInput"`

	// 播放模式，当Command取SetPlayMode时，必填。
	SetPlayModeCommandInput *SetPlayModeCommandInput `json:"SetPlayModeCommandInput,omitnil" name:"SetPlayModeCommandInput"`

	// 销毁模式，当Command取SetDestroyMode时，必填。
	SetDestroyModeCommandInput *SetDestroyModeCommandInput `json:"SetDestroyModeCommandInput,omitnil" name:"SetDestroyModeCommandInput"`

	// <del>音量，当Command取SetVolume时，必填。</del>
	// （已废弃，请采用 SetRealVolumeCommandInput ）
	SetVolumeCommandInput *SetVolumeCommandInput `json:"SetVolumeCommandInput,omitnil" name:"SetVolumeCommandInput"`

	// 真实音量，当Command取SetRealVolume时，必填。
	SetRealVolumeCommandInput *SetRealVolumeCommandInput `json:"SetRealVolumeCommandInput,omitnil" name:"SetRealVolumeCommandInput"`
}

type SyncKTVRobotCommandRequest struct {
	*tchttp.BaseRequest
	
	// 机器人Id。
	RobotId *string `json:"RobotId,omitnil" name:"RobotId"`

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
	// <li>SetDestroyMode：设置销毁模式</li>
	// <li><del>SetVolume：设置音量</del>（已废弃，请采用 SetRealVolume）</li>
	// <li>SetRealVolume：设置真实音量</li>
	Command *string `json:"Command,omitnil" name:"Command"`

	// 播放参数。
	PlayCommandInput *PlayCommandInput `json:"PlayCommandInput,omitnil" name:"PlayCommandInput"`

	// 播放列表变更信息，当Command取SetPlaylist时，必填。
	SetPlaylistCommandInput *SetPlaylistCommandInput `json:"SetPlaylistCommandInput,omitnil" name:"SetPlaylistCommandInput"`

	// 播放进度，当Command取Seek时，必填。
	SeekCommandInput *SeekCommandInput `json:"SeekCommandInput,omitnil" name:"SeekCommandInput"`

	// 音频参数，当Command取SetAudioParam时，必填。
	SetAudioParamCommandInput *SetAudioParamCommandInput `json:"SetAudioParamCommandInput,omitnil" name:"SetAudioParamCommandInput"`

	// 自定义消息，当Command取SendMessage时，必填。
	SendMessageCommandInput *SendMessageCommandInput `json:"SendMessageCommandInput,omitnil" name:"SendMessageCommandInput"`

	// 播放模式，当Command取SetPlayMode时，必填。
	SetPlayModeCommandInput *SetPlayModeCommandInput `json:"SetPlayModeCommandInput,omitnil" name:"SetPlayModeCommandInput"`

	// 销毁模式，当Command取SetDestroyMode时，必填。
	SetDestroyModeCommandInput *SetDestroyModeCommandInput `json:"SetDestroyModeCommandInput,omitnil" name:"SetDestroyModeCommandInput"`

	// <del>音量，当Command取SetVolume时，必填。</del>
	// （已废弃，请采用 SetRealVolumeCommandInput ）
	SetVolumeCommandInput *SetVolumeCommandInput `json:"SetVolumeCommandInput,omitnil" name:"SetVolumeCommandInput"`

	// 真实音量，当Command取SetRealVolume时，必填。
	SetRealVolumeCommandInput *SetRealVolumeCommandInput `json:"SetRealVolumeCommandInput,omitnil" name:"SetRealVolumeCommandInput"`
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
	delete(f, "SetDestroyModeCommandInput")
	delete(f, "SetVolumeCommandInput")
	delete(f, "SetRealVolumeCommandInput")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncKTVRobotCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncKTVRobotCommandResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// <li><del>SetVolume：设置音量</del>（已废弃，请采用 SetRealVolume）</li>
	// <li>SetRealVolume：设置真实音量</li>
	Command *string `json:"Command,omitnil" name:"Command"`

	// 播放参数。
	PlayCommandInput *PlayCommandInput `json:"PlayCommandInput,omitnil" name:"PlayCommandInput"`

	// 播放列表变更信息，当Command取SetPlaylist时，必填。
	SetPlaylistCommandInput *SetPlaylistCommandInput `json:"SetPlaylistCommandInput,omitnil" name:"SetPlaylistCommandInput"`

	// 播放进度，当Command取Seek时，必填。
	SeekCommandInput *SeekCommandInput `json:"SeekCommandInput,omitnil" name:"SeekCommandInput"`

	// 音频参数，当Command取SetAudioParam时，必填。
	SetAudioParamCommandInput *SetAudioParamCommandInput `json:"SetAudioParamCommandInput,omitnil" name:"SetAudioParamCommandInput"`

	// 自定义消息，当Command取SendMessage时，必填。
	SendMessageCommandInput *SendMessageCommandInput `json:"SendMessageCommandInput,omitnil" name:"SendMessageCommandInput"`

	// 播放模式，当Command取SetPlayMode时，必填。
	SetPlayModeCommandInput *SetPlayModeCommandInput `json:"SetPlayModeCommandInput,omitnil" name:"SetPlayModeCommandInput"`

	// 销毁模式，当Command取SetDestroyMode时，必填。
	SetDestroyModeCommandInput *SetDestroyModeCommandInput `json:"SetDestroyModeCommandInput,omitnil" name:"SetDestroyModeCommandInput"`

	// <del>音量，当Command取SetVolume时，必填。</del>
	// （已废弃，请采用 SetRealVolumeCommandInput）
	SetVolumeCommandInput *SetVolumeCommandInput `json:"SetVolumeCommandInput,omitnil" name:"SetVolumeCommandInput"`

	// 真实音量，当Command取SetRealVolume时，必填。
	SetRealVolumeCommandInput *SetRealVolumeCommandInput `json:"SetRealVolumeCommandInput,omitnil" name:"SetRealVolumeCommandInput"`
}

type TRTCJoinRoomInput struct {
	// 签名。
	Sign *string `json:"Sign,omitnil" name:"Sign"`

	// 房间号。
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// 推流应用ID。
	SdkAppId *string `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// 用户唯一标识。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 进房钥匙，若需要权限控制请携带该参数。
	//  [privateMapKey 权限设置](/document/product/647/32240) 
	PrivateMapKey *string `json:"PrivateMapKey,omitnil" name:"PrivateMapKey"`

	// 用户角色，目前支持两种角色：
	// <li>anchor：主播</li>
	// <li>audience：观众</li>
	Role *string `json:"Role,omitnil" name:"Role"`

	// TRTC房间号的类型：
	// <li>Integer：数字类型</li>
	// <li> String：字符串类型</li>
	// 默认为：Integer 。
	RoomIdType *string `json:"RoomIdType,omitnil" name:"RoomIdType"`
}

type TakeMusicOffShelves struct {
	// 资源方对应音乐Id
	MusicIds *string `json:"MusicIds,omitnil" name:"MusicIds"`

	// 当曲目临时下架时：已订购客户无影响，无需消息通知。当曲目封杀下架后，推送消息至已订购老客户，枚举值，判断是否上/下架
	// 在售状态，0在售，1临时下架，2永久下架
	SaleStatus *string `json:"SaleStatus,omitnil" name:"SaleStatus"`
}

// Predefined struct for user
type TakeMusicOffShelvesRequestParams struct {
	// 资源方下架必传结构
	TakeMusicOffShelves []*TakeMusicOffShelves `json:"TakeMusicOffShelves,omitnil" name:"TakeMusicOffShelves"`
}

type TakeMusicOffShelvesRequest struct {
	*tchttp.BaseRequest
	
	// 资源方下架必传结构
	TakeMusicOffShelves []*TakeMusicOffShelves `json:"TakeMusicOffShelves,omitnil" name:"TakeMusicOffShelves"`
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

// Predefined struct for user
type TakeMusicOffShelvesResponseParams struct {
	// 返回成功数量
	SuccessNum *int64 `json:"SuccessNum,omitnil" name:"SuccessNum"`

	// 返回失败数量
	FailedNum *int64 `json:"FailedNum,omitnil" name:"FailedNum"`

	// 返回失败歌曲musicId
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedMusicIds []*string `json:"FailedMusicIds,omitnil" name:"FailedMusicIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TakeMusicOffShelvesResponse struct {
	*tchttp.BaseResponse
	Response *TakeMusicOffShelvesResponseParams `json:"Response"`
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
	Before *string `json:"Before,omitnil" name:"Before"`

	// <li>小于此时间（结束时间）。</li>
	// <li>格式按照 ISO 8601标准表示，详见 <a href="https://cloud.tencent.com/document/product/266/11732#I" target="_blank">ISO 日期格式说明</a>。</li>
	After *string `json:"After,omitnil" name:"After"`
}

type UseRange struct {
	// 用途id
	UseRangeId *int64 `json:"UseRangeId,omitnil" name:"UseRangeId"`

	// 用途范围名称
	Name *string `json:"Name,omitnil" name:"Name"`
}