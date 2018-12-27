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

package v20180717

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ApplyUploadRequest struct {
	*tchttp.BaseRequest

	// 媒体类型，可选值请参考[上传能力综述](https://cloud.tencent.com/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B)。
	MediaType *string `json:"MediaType" name:"MediaType"`

	// 媒体名称。
	MediaName *string `json:"MediaName" name:"MediaName"`

	// 封面类型，可选值请参考[上传能力综述](https://cloud.tencent.com/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B)。
	CoverType *string `json:"CoverType" name:"CoverType"`

	// 媒体后续任务操作，详见[上传指定任务流](https://cloud.tencent.com/document/product/266/9759)。
	Procedure *string `json:"Procedure" name:"Procedure"`

	// 媒体文件过期时间，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	ExpireTime *string `json:"ExpireTime" name:"ExpireTime"`

	// 指定上传园区，仅适用于对上传地域有特殊需求的用户。
	StorageRegion *string `json:"StorageRegion" name:"StorageRegion"`

	// 分类ID，用于对媒体进行分类管理，可通过[创建分类](https://cloud.tencent.com/document/product/266/7812)接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId" name:"ClassId"`

	// 来源上下文，用于透传用户请求信息，上传回调接口将返回该字段值，最长 250 个字符。
	SourceContext *string `json:"SourceContext" name:"SourceContext"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId" name:"SubAppId"`
}

func (r *ApplyUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyUploadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApplyUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 存储桶，用于上传接口 URL 的 bucket_name。
		StorageBucket *string `json:"StorageBucket" name:"StorageBucket"`

		// 存储园区，用于上传接口 Host 的 Region。
		StorageRegion *string `json:"StorageRegion" name:"StorageRegion"`

		// 点播会话，用于确认上传接口的参数 VodSessionKey。
		VodSessionKey *string `json:"VodSessionKey" name:"VodSessionKey"`

		// 媒体存储路径，用于上传接口存储媒体的对象键（Key）。
		MediaStoragePath *string `json:"MediaStoragePath" name:"MediaStoragePath"`

		// 封面存储路径，用于上传接口存储封面的对象键（Key）。
		CoverStoragePath *string `json:"CoverStoragePath" name:"CoverStoragePath"`

		// 临时凭证，用于上传接口的权限验证。
		TempCertificate *TempCertificate `json:"TempCertificate" name:"TempCertificate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyUploadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommitUploadRequest struct {
	*tchttp.BaseRequest

	// 点播会话，取申请上传接口的返回值 VodSessionKey。
	VodSessionKey *string `json:"VodSessionKey" name:"VodSessionKey"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId" name:"SubAppId"`
}

func (r *CommitUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommitUploadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommitUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 媒体文件的唯一标识。
		FileId *string `json:"FileId" name:"FileId"`

		// 媒体播放地址。
		MediaUrl *string `json:"MediaUrl" name:"MediaUrl"`

		// 媒体封面地址。
		CoverUrl *string `json:"CoverUrl" name:"CoverUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CommitUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommitUploadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClassRequest struct {
	*tchttp.BaseRequest

	// 父类 ID，一级分类填写 -1。
	ParentId *int64 `json:"ParentId" name:"ParentId"`

	// 分类名称，长度限制：1-64 个字符。
	ClassName *string `json:"ClassName" name:"ClassName"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId" name:"SubAppId"`
}

func (r *CreateClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分类 ID
		ClassId *uint64 `json:"ClassId" name:"ClassId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClassRequest struct {
	*tchttp.BaseRequest

	// 分类 ID
	ClassId *int64 `json:"ClassId" name:"ClassId"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId" name:"SubAppId"`
}

func (r *DeleteClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaRequest struct {
	*tchttp.BaseRequest

	// 媒体文件的唯一标识。
	FileId *string `json:"FileId" name:"FileId"`

	// 指定本次需要删除的部分。默认值为 "[]", 表示删除媒体及其对应的全部视频处理文件。
	DeleteParts []*MediaDeleteItem `json:"DeleteParts" name:"DeleteParts" list`

	// 点播[子应用](/document/product/266/14574) ID 。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId" name:"SubAppId"`
}

func (r *DeleteMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAllClassRequest struct {
	*tchttp.BaseRequest

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId" name:"SubAppId"`
}

func (r *DescribeAllClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAllClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAllClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分类信息集合
		ClassInfoSet []*MediaClassInfo `json:"ClassInfoSet" name:"ClassInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAllClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaInfosRequest struct {
	*tchttp.BaseRequest

	// 媒体文件 ID 列表，N 从 0 开始取值，最大 19。
	FileIds []*string `json:"FileIds" name:"FileIds" list`

	// 指定所有媒体文件需要返回的信息，可同时指定多个信息，N 从 0 开始递增。如果未填写该字段，默认返回所有信息。选项有：
	// <li>basicInfo（视频基础信息）。</li>
	// <li>metaData（视频元信息）。</li>
	// <li>transcodeInfo（视频转码结果信息）。</li>
	// <li>animatedGraphicsInfo（视频转动图结果信息）。</li>
	// <li>imageSpriteInfo（视频雪碧图信息）。</li>
	// <li>snapshotByTimeOffsetInfo（视频指定时间点截图信息）。</li>
	// <li>sampleSnapshotInfo（采样截图信息）。</li>
	// <li>keyFrameDescInfo（打点信息）。</li>
	Filters []*string `json:"Filters" name:"Filters" list`
}

func (r *DescribeMediaInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 媒体文件信息列表。
		MediaInfoSet []*MediaInfo `json:"MediaInfoSet" name:"MediaInfoSet" list`

		// 不存在的文件 ID 列表。
		NotExistFileIdSet []*string `json:"NotExistFileIdSet" name:"NotExistFileIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MediaAnimatedGraphicsInfo struct {

	// 视频转动图结果信息
	AnimatedGraphicsSet []*MediaAnimatedGraphicsItem `json:"AnimatedGraphicsSet" name:"AnimatedGraphicsSet" list`
}

type MediaAnimatedGraphicsItem struct {

	// 转动图的文件地址。
	Url *string `json:"Url" name:"Url"`

	// 转动图模板 ID，参见[转动图参数模板](https://cloud.tencent.com/document/product/266/11701#.E9.A2.84.E7.BD.AE.E8.BD.AC.E5.8A.A8.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition" name:"Definition"`

	// 动图格式，如 gif。
	Container *string `json:"Container" name:"Container"`

	// 动图的高度，单位：px。
	Height *int64 `json:"Height" name:"Height"`

	// 动图的宽度，单位：px。
	Width *int64 `json:"Width" name:"Width"`

	// 动图码率，单位：bps。
	Bitrate *int64 `json:"Bitrate" name:"Bitrate"`

	// 动图大小，单位：字节。
	Size *int64 `json:"Size" name:"Size"`

	// 动图的md5值。
	Md5 *string `json:"Md5" name:"Md5"`

	// 动图在视频中的起始时间偏移，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset" name:"StartTimeOffset"`

	// 动图在视频中的结束时间偏移，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset" name:"EndTimeOffset"`
}

type MediaAudioStreamItem struct {

	// 音频流的码率，单位：bps。
	Bitrate *int64 `json:"Bitrate" name:"Bitrate"`

	// 音频流的采样率，单位：hz。
	SamplingRate *int64 `json:"SamplingRate" name:"SamplingRate"`

	// 音频流的编码格式，例如 aac。
	Codec *string `json:"Codec" name:"Codec"`
}

type MediaBasicInfo struct {

	// 媒体文件名称。
	Name *string `json:"Name" name:"Name"`

	// 媒体文件描述。
	Description *string `json:"Description" name:"Description"`

	// 媒体文件的创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime" name:"CreateTime"`

	// 媒体文件的最近更新时间（如修改视频属性、发起视频处理等会触发更新媒体文件信息的操作），使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`

	// 媒体文件的过期时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。“9999-12-31T23:59:59Z”表示永不过期。
	ExpireTime *string `json:"ExpireTime" name:"ExpireTime"`

	// 媒体文件的分类 ID。
	ClassId *int64 `json:"ClassId" name:"ClassId"`

	// 媒体文件的分类名称。
	ClassName *string `json:"ClassName" name:"ClassName"`

	// 媒体文件的分类路径，分类间以“-”分隔，如“新的一级分类 - 新的二级分类”。
	ClassPath *string `json:"ClassPath" name:"ClassPath"`

	// 媒体文件的封面图片地址。
	CoverUrl *string `json:"CoverUrl" name:"CoverUrl"`

	// 媒体文件的封装格式，例如 mp4、flv 等。
	Type *string `json:"Type" name:"Type"`

	// 原始媒体文件的 URL 地址。
	MediaUrl *string `json:"MediaUrl" name:"MediaUrl"`

	// 该媒体文件的来源信息。
	SourceInfo *MediaSourceData `json:"SourceInfo" name:"SourceInfo"`

	// 媒体文件存储地区，如 ap-guangzhou，参见[地域列表](https://cloud.tencent.com/document/api/213/15692#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)。
	StorageRegion *string `json:"StorageRegion" name:"StorageRegion"`

	// 媒体文件的标签信息。
	TagSet []*string `json:"TagSet" name:"TagSet" list`

	// 直播录制文件的唯一标识
	Vid *string `json:"Vid" name:"Vid"`
}

type MediaClassInfo struct {

	// 分类 ID
	ClassId *int64 `json:"ClassId" name:"ClassId"`

	// 父类 ID，一级分类的父类 ID 为 -1。
	ParentId *int64 `json:"ParentId" name:"ParentId"`

	// 分类名称
	ClassName *string `json:"ClassName" name:"ClassName"`

	// 分类级别，一级分类为 0，最大值为 3，即最多允许 4 级分类层次。
	Level *uint64 `json:"Level" name:"Level"`

	// 当前分类的第一级子类 ID 集合
	SubClassIdSet []*int64 `json:"SubClassIdSet" name:"SubClassIdSet" list`
}

type MediaDeleteItem struct {

	// 所指定的删除部分。如果未填写该字段则参数无效。可选值有：
	// <li>TranscodeFiles（删除转码文件）。</li>
	// <li>WechatPublishFiles（删除微信发布文件）。</li>
	Type *string `json:"Type" name:"Type"`

	// 删除由Type参数指定的种类下的视频模板号，模板定义参见[转码模板](https://cloud.tencent.com/document/product/266/11701#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
	// 默认值为0，表示删除参数Type指定种类下所有的视频。
	Definition *int64 `json:"Definition" name:"Definition"`
}

type MediaImageSpriteInfo struct {

	// 特定规格的雪碧图信息集合，每个元素代表一套相同规格的雪碧图。
	ImageSpriteSet []*MediaImageSpriteItem `json:"ImageSpriteSet" name:"ImageSpriteSet" list`
}

type MediaImageSpriteItem struct {

	// 雪碧图规格，参见[雪碧图参数模板](https://cloud.tencent.com/document/product/266/11702#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition" name:"Definition"`

	// 雪碧图小图的高度。
	Height *int64 `json:"Height" name:"Height"`

	// 雪碧图小图的宽度。
	Width *int64 `json:"Width" name:"Width"`

	// 每一张雪碧图大图里小图的数量。
	TotalCount *int64 `json:"TotalCount" name:"TotalCount"`

	// 每一张雪碧图大图的地址。
	ImageUrlSet []*string `json:"ImageUrlSet" name:"ImageUrlSet" list`

	// 雪碧图子图位置与时间关系的 WebVtt 文件地址。WebVtt 文件表明了各个雪碧图小图对应的时间点，以及在在雪碧大图里的坐标位置，一般被播放器用于实现预览。
	WebVttUrl *string `json:"WebVttUrl" name:"WebVttUrl"`
}

type MediaInfo struct {

	// 基础信息。包括视频名称、大小、时长、封面图片等。
	BasicInfo *MediaBasicInfo `json:"BasicInfo" name:"BasicInfo"`

	// 元信息。包括视频流信息、音频流信息等。
	MetaData *MediaMetaData `json:"MetaData" name:"MetaData"`

	// 转码结果信息。包括该视频转码生成的各种码率的视频的地址、规格、码率、分辨率等。
	TranscodeInfo *MediaTranscodeInfo `json:"TranscodeInfo" name:"TranscodeInfo"`

	// 转动图结果信息。对视频转动图（如 gif）后，动图相关信息。
	AnimatedGraphicsInfo *MediaAnimatedGraphicsInfo `json:"AnimatedGraphicsInfo" name:"AnimatedGraphicsInfo"`

	// 采样截图信息。对视频采样截图后，相关截图信息。
	SampleSnapshotInfo *MediaSampleSnapshotInfo `json:"SampleSnapshotInfo" name:"SampleSnapshotInfo"`

	// 雪碧图信息。对视频截取雪碧图之后，雪碧的相关信息。
	ImageSpriteInfo *MediaImageSpriteInfo `json:"ImageSpriteInfo" name:"ImageSpriteInfo"`

	// 指定时间点截图信息。对视频依照指定时间点截图后，各个截图的信息。
	SnapshotByTimeOffsetInfo *MediaSnapshotByTimeOffsetInfo `json:"SnapshotByTimeOffsetInfo" name:"SnapshotByTimeOffsetInfo"`

	// 视频打点信息。对视频设置的各个打点信息。
	KeyFrameDescInfo *MediaKeyFrameDescInfo `json:"KeyFrameDescInfo" name:"KeyFrameDescInfo"`

	// 媒体文件唯一标识 ID。
	FileId *string `json:"FileId" name:"FileId"`
}

type MediaKeyFrameDescInfo struct {

	// 视频打点信息数组。
	KeyFrameDescSet []*MediaKeyFrameDescItem `json:"KeyFrameDescSet" name:"KeyFrameDescSet" list`
}

type MediaKeyFrameDescItem struct {

	// 打点的视频偏移时间，单位：秒。
	TimeOffset *float64 `json:"TimeOffset" name:"TimeOffset"`

	// 打点的内容字符串，限制 1-128 个字符。
	Content *string `json:"Content" name:"Content"`
}

type MediaMetaData struct {

	// 上传的媒体文件大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	Size *int64 `json:"Size" name:"Size"`

	// 容器类型，例如 m4a，mp4 等。
	Container *string `json:"Container" name:"Container"`

	// 视频流码率平均值与音频流码率平均值之和，单位：bps。
	Bitrate *int64 `json:"Bitrate" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width" name:"Width"`

	// 视频时长，单位：秒。
	Duration *float64 `json:"Duration" name:"Duration"`

	// 视频拍摄时的选择角度，单位：度。
	Rotate *int64 `json:"Rotate" name:"Rotate"`

	// 视频流信息。
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet" name:"VideoStreamSet" list`

	// 音频流信息。
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet" name:"AudioStreamSet" list`

	// 视频时长，单位：秒。
	VideoDuration *float64 `json:"VideoDuration" name:"VideoDuration"`

	// 音频时长，单位：秒。
	AudioDuration *float64 `json:"AudioDuration" name:"AudioDuration"`
}

type MediaSampleSnapshotInfo struct {

	// 特定规格的采样截图信息集合，每个元素代表一套相同规格的采样截图。
	SampleSnapshotSet []*MediaSampleSnapshotItem `json:"SampleSnapshotSet" name:"SampleSnapshotSet" list`
}

type MediaSampleSnapshotItem struct {

	// 采样截图规格 ID，参见[采样截图参数模板](https://cloud.tencent.com/document/product/266/11702#.E9.87.87.E6.A0.B7.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition" name:"Definition"`

	// 采样方式，取值范围：
	// <li>Percent：根据百分比间隔采样。</li>
	// <li>Time：根据时间间隔采样。</li>
	SampleType *string `json:"SampleType" name:"SampleType"`

	// 采样间隔
	// <li>当 SampleType 为 Percent 时，该值表示多少百分比一张图。</li>
	// <li>当 SampleType 为 Time 时，该值表示多少时间间隔一张图，单位秒， 第一张图均为视频首帧。</li>
	Interval *int64 `json:"Interval" name:"Interval"`

	// 生成的截图 url 列表。
	ImageUrlSet []*string `json:"ImageUrlSet" name:"ImageUrlSet" list`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition" name:"WaterMarkDefinition" list`
}

type MediaSnapshotByTimeOffsetInfo struct {

	// 特定规格的指定时间点截图信息集合。目前每种规格只能有一套截图。
	SnapshotByTimeOffsetSet []*MediaSnapshotByTimeOffsetItem `json:"SnapshotByTimeOffsetSet" name:"SnapshotByTimeOffsetSet" list`
}

type MediaSnapshotByTimeOffsetItem struct {

	// 指定时间点截图规格，参见[指定时间点截图参数模板](https://cloud.tencent.com/document/product/266/11702#.E6.8C.87.E5.AE.9A.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition" name:"Definition"`

	// 同一规格的截图信息集合，每个元素代表一张截图。
	PicInfoSet []*MediaSnapshotByTimePicInfoItem `json:"PicInfoSet" name:"PicInfoSet" list`
}

type MediaSnapshotByTimePicInfoItem struct {

	// 该张截图对应视频文件中的时间偏移，单位：秒。
	TimeOffset *float64 `json:"TimeOffset" name:"TimeOffset"`

	// 该张截图的 URL 地址。
	Url *string `json:"Url" name:"Url"`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition" name:"WaterMarkDefinition" list`
}

type MediaSourceData struct {

	// 媒体文件的来源类别：
	// <li>Record：来自录制。如直播录制、直播时移录制等。</li>
	// <li>Upload：来自上传。如拉取上传、服务端上传、客户端 UGC 上传等。</li>
	// <li>VideoProcessing：来自视频处理。如视频拼接、视频剪辑等。</li>
	// <li>Unknown：未知来源。</li>
	SourceType *string `json:"SourceType" name:"SourceType"`

	// 用户创建文件时透传的字段
	SourceContext *string `json:"SourceContext" name:"SourceContext"`
}

type MediaTranscodeInfo struct {

	// 各规格的转码信息集合，每个元素代表一个规格的转码结果。
	TranscodeSet []*MediaTranscodeItem `json:"TranscodeSet" name:"TranscodeSet" list`
}

type MediaTranscodeItem struct {

	// 转码后的视频文件地址。
	Url *string `json:"Url" name:"Url"`

	// 转码规格 ID，参见[转码参数模板](https://cloud.tencent.com/document/product/266/11701#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition" name:"Definition"`

	// 视频流码率平均值与音频流码率平均值之和， 单位：bps。
	Bitrate *int64 `json:"Bitrate" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width" name:"Width"`

	// 媒体文件总大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	Size *int64 `json:"Size" name:"Size"`

	// 视频时长，单位：秒。
	Duration *float64 `json:"Duration" name:"Duration"`

	// 容器类型，例如 m4a，mp4 等。
	Container *string `json:"Container" name:"Container"`

	// 视频的 md5 值。
	Md5 *string `json:"Md5" name:"Md5"`

	// 音频流信息。
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet" name:"AudioStreamSet" list`

	// 视频流信息。
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet" name:"VideoStreamSet" list`
}

type MediaVideoStreamItem struct {

	// 视频流的码率，单位：bps。
	Bitrate *int64 `json:"Bitrate" name:"Bitrate"`

	// 视频流的高度，单位：px。
	Height *int64 `json:"Height" name:"Height"`

	// 视频流的宽度，单位：px。
	Width *int64 `json:"Width" name:"Width"`

	// 视频流的编码格式，例如 h264。
	Codec *string `json:"Codec" name:"Codec"`

	// 帧率，单位：hz。
	Fps *int64 `json:"Fps" name:"Fps"`
}

type ModifyClassRequest struct {
	*tchttp.BaseRequest

	// 分类 ID
	ClassId *uint64 `json:"ClassId" name:"ClassId"`

	// 分类名称。长度限制：1-64 个字符。
	ClassName *string `json:"ClassName" name:"ClassName"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId" name:"SubAppId"`
}

func (r *ModifyClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaInfoRequest struct {
	*tchttp.BaseRequest

	// 媒体文件唯一标识。
	FileId *string `json:"FileId" name:"FileId"`

	// 媒体文件名称，最长 64 个字符。
	Name *string `json:"Name" name:"Name"`

	// 媒体文件描述，最长 128 个字符。
	Description *string `json:"Description" name:"Description"`

	// 媒体文件分类 ID。
	ClassId *int64 `json:"ClassId" name:"ClassId"`

	// 媒体文件过期时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。填“9999-12-31T23:59:59Z”表示永不过期。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。
	ExpireTime *string `json:"ExpireTime" name:"ExpireTime"`

	// 视频封面图片文件（如 jpeg, png 等）进行 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串，仅支持 gif、jpeg、png 三种图片格式。
	CoverData *string `json:"CoverData" name:"CoverData"`

	// 新增的一组视频打点信息，如果某个偏移时间已存在打点，则会进行覆盖操作，单个媒体文件最多 100 个打点信息。同一个请求里，AddKeyFrameDescs 的时间偏移参数必须与 DeleteKeyFrameDescs 都不同。
	AddKeyFrameDescs []*MediaKeyFrameDescItem `json:"AddKeyFrameDescs" name:"AddKeyFrameDescs" list`

	// 要删除的一组视频打点信息的时间偏移，单位：秒。同一个请求里，AddKeyFrameDescs 的时间偏移参数必须与 DeleteKeyFrameDescs 都不同。
	DeleteKeyFrameDescs []*float64 `json:"DeleteKeyFrameDescs" name:"DeleteKeyFrameDescs" list`

	// 取值 1 表示清空视频打点信息，其他值无意义。
	// 同一个请求里，ClearKeyFrameDescs 与 AddKeyFrameDescs 不能同时出现。
	ClearKeyFrameDescs *int64 `json:"ClearKeyFrameDescs" name:"ClearKeyFrameDescs"`

	// 新增的一组标签，单个媒体文件最多 16 个标签，单个标签最多 16 个字符。同一个请求里，AddTags 参数必须与 DeleteTags 都不同。
	AddTags []*string `json:"AddTags" name:"AddTags" list`

	// 要删除的一组标签。同一个请求里，AddTags 参数必须与 DeleteTags 都不同。
	DeleteTags []*string `json:"DeleteTags" name:"DeleteTags" list`

	// 取值 1 表示清空媒体文件所有标签，其他值无意义。
	// 同一个请求里，ClearTags 与 AddTags 不能同时出现。
	ClearTags *int64 `json:"ClearTags" name:"ClearTags"`

	// 点播[子应用](/document/product/266/14574) ID 。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId" name:"SubAppId"`
}

func (r *ModifyMediaInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新的视频封面 URL。
	// * 注意：仅当请求携带 CoverData 时此返回值有效。 *
		CoverUrl *string `json:"CoverUrl" name:"CoverUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMediaInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchMediaRequest struct {
	*tchttp.BaseRequest

	// 搜索文本，模糊匹配媒体文件名称或描述信息，匹配项越多，匹配度越高，排序越优先。长度限制：64 个字符。
	Text *string `json:"Text" name:"Text"`

	// 标签集合，匹配集合中任意元素。
	// <li>单个标签长度限制：8 个字符。</li>
	// <li>数组长度限制：10。</li>
	Tags []*string `json:"Tags" name:"Tags" list`

	// 分类 ID 集合，匹配集合指定 ID 的分类及其所有子类。数组长度限制：10。
	ClassIds []*int64 `json:"ClassIds" name:"ClassIds" list`

	// 创建时间的开始时间。
	// <li>大于等于开始时间。</li>
	// <li>格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。</li>
	StartTime *string `json:"StartTime" name:"StartTime"`

	// 创建时间的结束时间。
	// <li>小于结束时间。</li>
	// <li>格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。</li>
	EndTime *string `json:"EndTime" name:"EndTime"`

	// 媒体文件来源，来源取值参见 [SourceType](https://cloud.tencent.com/document/product/266/31773#MediaSourceData)。
	SourceType *string `json:"SourceType" name:"SourceType"`

	// 推流[直播码](https://cloud.tencent.com/document/product/267/5959)。
	StreamId *string `json:"StreamId" name:"StreamId"`

	// 直播录制文件的唯一标识。
	Vid *string `json:"Vid" name:"Vid"`

	// 排序方式。
	// <li>Sort.Field 可选值：CreateTime</li>
	// <li>指定 Text 搜索时，将根据匹配度排序，该字段无效</li>
	Sort *SortBy `json:"Sort" name:"Sort"`

	// 偏移量。
	// <li>默认值：0。</li>
	// <li>取值范围：Offset + Limit 不超过 5000。</li>
	Offset *uint64 `json:"Offset" name:"Offset"`

	// 返回记录条数，默认值：10。
	Limit *uint64 `json:"Limit" name:"Limit"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId" name:"SubAppId"`
}

func (r *SearchMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合搜索条件的记录总数
	// <li>最大值：5000，即，当命中记录数超过 5000，该字段将返回 5000，而非实际命中总数。</li>
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`

		// 媒体文件信息列表，只包含基础信息（BasicInfo）
		MediaInfoSet []*MediaInfo `json:"MediaInfoSet" name:"MediaInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SortBy struct {

	// 排序字段
	Field *string `json:"Field" name:"Field"`

	// 排序方式，可选值：Asc（升序）、Desc（降序）
	Order *string `json:"Order" name:"Order"`
}

type TempCertificate struct {

	// 临时安全证书 Id。
	SecretId *string `json:"SecretId" name:"SecretId"`

	// 临时安全证书 Key。
	SecretKey *string `json:"SecretKey" name:"SecretKey"`

	// Token 值。
	Token *string `json:"Token" name:"Token"`

	// 证书无效的时间，返回 Unix 时间戳，精确到秒。
	ExpiredTime *uint64 `json:"ExpiredTime" name:"ExpiredTime"`
}
