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

package v20191029

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AudioStreamInfo struct {

	// 码率，单位：bps。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 采样率，单位：hz。
	SamplingRate *uint64 `json:"SamplingRate,omitempty" name:"SamplingRate"`

	// 编码格式。
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

type CMEExportInfo struct {

	// 导出的归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 导出的素材名称，不得超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 导出的素材信息，不得超过50个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 导出的素材分类路径，长度不能超过15字符。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 导出的素材标签，单个标签不得超过10个字符。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet" list`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目类别，取值有：
	// <li>VIDEO_EDIT：视频编辑。</li>
	Category *string `json:"Category,omitempty" name:"Category"`

	// 项目名称，不可超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 画布宽高比，取值有：
	// <li>16:9；</li>
	// <li>9:16。</li>
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目 Id。
		ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLoginStatusRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 用户 Id 列表，N 从 0 开始取值，最大 19。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds" list`
}

func (r *DeleteLoginStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLoginStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLoginStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoginStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLoginStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginStatusRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 用户 Id 列表，N 从 0 开始取值，最大 19。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds" list`
}

func (r *DescribeLoginStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoginStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户登录状态列表。
		LoginStatusInfoSet []*LoginStatusInfo `json:"LoginStatusInfoSet,omitempty" name:"LoginStatusInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoginStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id 列表，N 从 0 开始取值，最大 19。
	ProjectIds []*string `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// 画布宽高比集合。
	AspectRatioSet []*string `json:"AspectRatioSet,omitempty" name:"AspectRatioSet" list`

	// 项目类别集合。
	CategorySet []*string `json:"CategorySet,omitempty" name:"CategorySet" list`

	// 项目归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 分页返回的起始偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 项目信息列表。
		ProjectInfoSet []*ProjectInfo `json:"ProjectInfoSet,omitempty" name:"ProjectInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务状态，取值有：
	// <li>PROCESSING：处理中：</li>
	// <li>SUCCESS：成功；</li>
	// <li>FAIL：失败。</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// 任务进度，取值为：0~100。
		Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

		// 错误码。
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
		ErrCode *uint64 `json:"ErrCode,omitempty" name:"ErrCode"`

		// 错误信息。
		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

		// 任务类型，取值有：
	// <li>VIDEO_EDIT_PROJECT_EXPORT：视频编辑项目导出。</li>
		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

		// 导出项目输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		VideoEditProjectOutput *VideoEditProjectOutput `json:"VideoEditProjectOutput,omitempty" name:"VideoEditProjectOutput"`

		// 创建时间，格式按照 ISO 8601 标准表示。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型集合，取值有：
	// <li>VIDEO_EDIT_PROJECT_EXPORT：视频编辑项目导出。</li>
	TaskTypeSet []*string `json:"TaskTypeSet,omitempty" name:"TaskTypeSet" list`

	// 任务状态集合，取值有：
	// <li>PROCESSING：处理中；</li>
	// <li>SUCCESS：成功；</li>
	// <li>FAIL：失败。</li>
	StatusSet []*string `json:"StatusSet,omitempty" name:"StatusSet" list`

	// 分页返回的起始偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值：10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合搜索条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务基础信息列表。
		TaskBaseInfoSet []*TaskBaseInfo `json:"TaskBaseInfoSet,omitempty" name:"TaskBaseInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Entity struct {

	// 类型，取值有：
	// <li>PERSON：个人。</li>
	// <li>TEAM：团队。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// Id，当 Type=PERSON，取值为用户 Id，当 Type=TEAM，取值为团队 Id。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type ExportVideoEditProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 导出模板 Id，目前不支持自定义创建，只支持下面的预置模板 Id。
	// <li>10：分辨率为 480P，输出视频格式为 MP4；</li>
	// <li>11：分辨率为 720P，输出视频格式为 MP4；</li>
	// <li>12：分辨率为 1080P，输出视频格式为 MP4。</li>
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 导出目标。
	// <li>CME：云剪，即导出为云剪素材；</li>
	// <li>VOD：云点播，即导出为云点播媒资。</li>
	ExportDestination *string `json:"ExportDestination,omitempty" name:"ExportDestination"`

	// 导出的云剪素材信息。指定 ExportDestination = CME 时有效。
	CMEExportInfo *CMEExportInfo `json:"CMEExportInfo,omitempty" name:"CMEExportInfo"`

	// 导出的云点播媒资信息。指定 ExportDestination = VOD 时有效。
	VODExportInfo *VODExportInfo `json:"VODExportInfo,omitempty" name:"VODExportInfo"`
}

func (r *ExportVideoEditProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportVideoEditProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExportVideoEditProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 Id。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVideoEditProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExportVideoEditProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportMediaToProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 云点播媒资 FileId。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 素材名称，不能超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 素材预处理任务模板 ID，取值：
	// <li>10：进行编辑预处理。</li>
	// 注意：如果填0则不进行处理。
	PreProcessDefinition *int64 `json:"PreProcessDefinition,omitempty" name:"PreProcessDefinition"`
}

func (r *ImportMediaToProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportMediaToProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportMediaToProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 素材 Id。
		MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

		// 素材预处理任务 ID，如果未指定发起预处理任务则为空。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportMediaToProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportMediaToProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LoginStatusInfo struct {

	// 用户 Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户登录状态。
	// <li>Online：在线；</li>
	// <li>Offline：离线。</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type MaterialBaseInfo struct {

	// 素材名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 分类路径。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 标签集合。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet" list`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 素材类型。
	MaterialType *string `json:"MaterialType,omitempty" name:"MaterialType"`

	// 素材 URL。
	MaterialUrl *string `json:"MaterialUrl,omitempty" name:"MaterialUrl"`

	// 云点播媒资 FileId。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 创建时间，格式按照 ISO 8601 标准表示。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type MediaMetaData struct {

	// 大小。
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 容器类型。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 视频流码率平均值与音频流码率平均值之和，单位：bps。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 时长，单位：秒。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 视频拍摄时的选择角度，单位：度
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 视频流信息。
	VideoStreamInfoSet []*VideoStreamInfo `json:"VideoStreamInfoSet,omitempty" name:"VideoStreamInfoSet" list`

	// 音频流信息。
	AudioStreamInfoSet []*AudioStreamInfo `json:"AudioStreamInfoSet,omitempty" name:"AudioStreamInfoSet" list`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest

	// 平台名称，指定访问的平台。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称，不可超过30个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProjectInfo struct {

	// 项目 Id。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 画布宽高比。
	AspectRatio *string `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 项目类别。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 归属者。
	Owner *Entity `json:"Owner,omitempty" name:"Owner"`

	// 项目封面图片地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 项目创建时间，格式按照 ISO 8601 标准表示。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目更新时间，格式按照 ISO 8601 标准表示。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TaskBaseInfo struct {

	// 任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务类型，取值有：
	// <li>VIDEO_EDIT_PROJECT_EXPORT：项目导出。</li>
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务状态，取值有：
	// <li>PROCESSING：处理中：</li>
	// <li>SUCCESS：成功；</li>
	// <li>FAIL：失败。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务进度，取值为：0~100。
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 错误码。
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 创建时间，格式按照 ISO 8601 标准表示。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type VODExportInfo struct {

	// 导出的媒资名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 导出的媒资分类 Id。
	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`
}

type VideoEditProjectOutput struct {

	// 云点播媒资 FileId。
	VodFileId *string `json:"VodFileId,omitempty" name:"VodFileId"`

	// 导出的媒资 URL。
	URL *string `json:"URL,omitempty" name:"URL"`

	// 元信息。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// 素材基础信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaterialBaseInfo *MaterialBaseInfo `json:"MaterialBaseInfo,omitempty" name:"MaterialBaseInfo"`
}

type VideoStreamInfo struct {

	// 码率，单位：bps。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 高度，单位：px。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 宽度，单位：px。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 编码格式。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 帧率，单位：hz。
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`
}
