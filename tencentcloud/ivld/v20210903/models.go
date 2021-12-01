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

package v20210903

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AppearIndexPair struct {

	// 出现信息，取值范围为[1，3]
	AppearIndex *int64 `json:"AppearIndex,omitempty" name:"AppearIndex"`

	// AppearInfo中AppearIndex对应元素的第Index元素，从0开始技术
	Index *int64 `json:"Index,omitempty" name:"Index"`
}

type AppearInfo struct {

	// 关键词在音频文本结果中的出现位置数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioAppearSet []*TextAppearInfo `json:"AudioAppearSet,omitempty" name:"AudioAppearSet"`

	// 关键词在可视文本结果中的出现位置数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextAppearSet []*TextAppearInfo `json:"TextAppearSet,omitempty" name:"TextAppearSet"`

	// 关键词在视频信息中的出现位置数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoAppearSet []*VideoAppearInfo `json:"VideoAppearSet,omitempty" name:"VideoAppearSet"`
}

type AudioInfo struct {

	// ASR提取的文字信息
	Content *string `json:"Content,omitempty" name:"Content"`

	// ASR起始时间戳，从0开始
	StartTimeStamp *float64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// ASR结束时间戳，从0开始
	EndTimeStamp *float64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// ASR提取的音频标签
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest

	// 媒资文件ID，最长32B
	MediaId *string `json:"MediaId,omitempty" name:"MediaId"`

	// 媒资素材先验知识，相关限制参考MediaPreknownInfo
	MediaPreknownInfo *MediaPreknownInfo `json:"MediaPreknownInfo,omitempty" name:"MediaPreknownInfo"`

	// 任务名称，最长100个中文字符
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 是否上传转码后的视频，仅设置true时上传，默认为false
	UploadVideo *bool `json:"UploadVideo,omitempty" name:"UploadVideo"`
}

func (r *CreateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MediaId")
	delete(f, "MediaPreknownInfo")
	delete(f, "TaskName")
	delete(f, "UploadVideo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 智能标签视频分析任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Data struct {

	// 节目粒度结构化结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowInfo *ShowInfo `json:"ShowInfo,omitempty" name:"ShowInfo"`
}

type DeleteMediaRequest struct {
	*tchttp.BaseRequest

	// 媒资文件在系统中的ID
	MediaId *string `json:"MediaId,omitempty" name:"MediaId"`
}

func (r *DeleteMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MediaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaRequest struct {
	*tchttp.BaseRequest

	// 导入媒资返回的媒资ID，最长32B
	MediaId *string `json:"MediaId,omitempty" name:"MediaId"`
}

func (r *DescribeMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MediaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 媒资信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		MediaInfo *MediaInfo `json:"MediaInfo,omitempty" name:"MediaInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMediasRequest struct {
	*tchttp.BaseRequest

	// 分页序号，从1开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每个分页所包含的元素数量，最大为50
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 列举过滤条件，相关限制相见MediaFilter
	MediaFilter *MediaFilter `json:"MediaFilter,omitempty" name:"MediaFilter"`

	// 返回结果排序信息，By字段只支持CreateTime
	SortBy *SortBy `json:"SortBy,omitempty" name:"SortBy"`
}

func (r *DescribeMediasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "MediaFilter")
	delete(f, "SortBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMediasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 满足过滤条件的媒资视频总数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 满足过滤条件的媒资信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		MediaInfoSet []*MediaInfo `json:"MediaInfoSet,omitempty" name:"MediaInfoSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 创建任务返回的TaskId
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务信息，不包含任务结果
		TaskInfo *TaskInfo `json:"TaskInfo,omitempty" name:"TaskInfo"`

		// 任务结果数据，只在任务结束时返回
		TaskData *Data `json:"TaskData,omitempty" name:"TaskData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest

	// CreateTask返回的任务ID，最长32B
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务信息，详情参见TaskInfo的定义
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskInfo *TaskInfo `json:"TaskInfo,omitempty" name:"TaskInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 分页序号，从1开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每个分页所包含的元素数量，最大为50
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 任务过滤条件，相关限制参见TaskFilter
	TaskFilter *TaskFilter `json:"TaskFilter,omitempty" name:"TaskFilter"`

	// 返回结果排序信息，By字段只支持CreateTimeStamp
	SortBy *SortBy `json:"SortBy,omitempty" name:"SortBy"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "TaskFilter")
	delete(f, "SortBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 满足过滤条件的任务总数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 满足过滤条件的任务数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskInfoSet []*TaskInfo `json:"TaskInfoSet,omitempty" name:"TaskInfoSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportMediaRequest struct {
	*tchttp.BaseRequest

	// 待分析视频的URL，目前只支持*不带签名的*COS地址，长度最长1KB
	URL *string `json:"URL,omitempty" name:"URL"`

	// 待分析视频的MD5，为空时不做校验，否则会做MD5校验，长度必须为32B
	MD5 *string `json:"MD5,omitempty" name:"MD5"`

	// 待分析视频的名称，指定后可支持筛选，最多100个中文字符
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ImportMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportMediaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "URL")
	delete(f, "MD5")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ImportMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 媒资文件在系统中的ID
		MediaId *string `json:"MediaId,omitempty" name:"MediaId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportMediaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type L1Tag struct {

	// 一级标签名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 二级标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	L2TagSet []*L2Tag `json:"L2TagSet,omitempty" name:"L2TagSet"`

	// 一级标签出现信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppearIndexPairSet []*AppearIndexPair `json:"AppearIndexPairSet,omitempty" name:"AppearIndexPairSet"`

	// 一级标签首次出现信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstAppear *int64 `json:"FirstAppear,omitempty" name:"FirstAppear"`
}

type L2Tag struct {

	// 二级标签名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 从属于此二级标签的三级标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	L3TagSet []*L3Tag `json:"L3TagSet,omitempty" name:"L3TagSet"`

	// 二级标签出现信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppearIndexPairSet []*AppearIndexPair `json:"AppearIndexPairSet,omitempty" name:"AppearIndexPairSet"`

	// 二级标签首次出现信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstAppear *int64 `json:"FirstAppear,omitempty" name:"FirstAppear"`
}

type L3Tag struct {

	// 三级标签名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 三级标签出现信息索引数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppearIndexPairSet []*AppearIndexPair `json:"AppearIndexPairSet,omitempty" name:"AppearIndexPairSet"`

	// 三级标签首次出现信息，可选值为[1,3]
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstAppear *int64 `json:"FirstAppear,omitempty" name:"FirstAppear"`
}

type MediaFilter struct {

	// 媒资名称过滤条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaNameSet []*string `json:"MediaNameSet,omitempty" name:"MediaNameSet"`

	// 媒资状态数组，媒资状态可选值参见MediaInfo
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet"`

	// 媒资ID数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaIdSet []*string `json:"MediaIdSet,omitempty" name:"MediaIdSet"`
}

type MediaInfo struct {

	// 媒资ID
	MediaId *string `json:"MediaId,omitempty" name:"MediaId"`

	// 媒资名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 媒资下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownLoadURL *string `json:"DownLoadURL,omitempty" name:"DownLoadURL"`

	// 媒资状态，取值参看上方表格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 若状态为失败，表示失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedReason *string `json:"FailedReason,omitempty" name:"FailedReason"`

	// 媒资视频元信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metadata *MediaMetadata `json:"Metadata,omitempty" name:"Metadata"`

	// 导入视频进度，取值范围为[0,100]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *float64 `json:"Progress,omitempty" name:"Progress"`
}

type MediaMetadata struct {

	// 媒资视频文件大小
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 媒资视频文件MD5
	MD5 *string `json:"MD5,omitempty" name:"MD5"`

	// 媒资视频时长，单位为秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 媒资视频总帧数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumFrames *int64 `json:"NumFrames,omitempty" name:"NumFrames"`

	// 媒资视频宽度，单位为像素
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 媒资视频高度，单位为像素
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 媒资视频帧率，单位为Hz
	// 注意：此字段可能返回 null，表示取不到有效值。
	FPS *float64 `json:"FPS,omitempty" name:"FPS"`

	// 媒资视频比特率，单位为kbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	BitRate *int64 `json:"BitRate,omitempty" name:"BitRate"`
}

type MediaPreknownInfo struct {

	// 媒资文件类型，参见MediaPreknownInfo结构体定义
	MediaType *int64 `json:"MediaType,omitempty" name:"MediaType"`

	// 媒资素材一级类型，参见MediaPreknownInfo结构体定义
	MediaLabel *int64 `json:"MediaLabel,omitempty" name:"MediaLabel"`

	// 媒资素材二级类型，参见MediaPreknownInfo结构体定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaSecondLabel *int64 `json:"MediaSecondLabel,omitempty" name:"MediaSecondLabel"`

	// 媒资音频类型，参见MediaPreknownInfo结构体定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaLang *int64 `json:"MediaLang,omitempty" name:"MediaLang"`
}

type MultiLevelTag struct {

	// 树状标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*L1Tag `json:"TagSet,omitempty" name:"TagSet"`

	// 标签在识别结果中的定位信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppearInfo *AppearInfo `json:"AppearInfo,omitempty" name:"AppearInfo"`
}

type ShowInfo struct {

	// 节目日期(只在新闻有效)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 台标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logo *string `json:"Logo,omitempty" name:"Logo"`

	// 节目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Column *string `json:"Column,omitempty" name:"Column"`

	// 来源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 节目封面
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverImageURL *string `json:"CoverImageURL,omitempty" name:"CoverImageURL"`

	// 节目内容概要列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummarySet []*string `json:"SummarySet,omitempty" name:"SummarySet"`

	// 节目片段标题列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TitleSet []*string `json:"TitleSet,omitempty" name:"TitleSet"`

	// 音频识别结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioInfoSet []*AudioInfo `json:"AudioInfoSet,omitempty" name:"AudioInfoSet"`

	// 可视文字识别结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextInfoSet []*TextInfo `json:"TextInfoSet,omitempty" name:"TextInfoSet"`

	// 文本标签列表，包含标签内容和出现信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTagSet *MultiLevelTag `json:"TextTagSet,omitempty" name:"TextTagSet"`

	// 帧标签列表，包括人物信息，场景信息等
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameTagSet *MultiLevelTag `json:"FrameTagSet,omitempty" name:"FrameTagSet"`

	// 视频下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebMediaURL *string `json:"WebMediaURL,omitempty" name:"WebMediaURL"`

	// 媒资分类信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaClassifierSet []*string `json:"MediaClassifierSet,omitempty" name:"MediaClassifierSet"`

	// 概要标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryTagSet []*string `json:"SummaryTagSet,omitempty" name:"SummaryTagSet"`
}

type SortBy struct {

	// 排序字段，默认为CreateTime
	By *string `json:"By,omitempty" name:"By"`

	// true表示降序，false表示升序
	Descend *bool `json:"Descend,omitempty" name:"Descend"`
}

type TaskFilter struct {

	// 媒资文件类型
	MediaTypeSet []*int64 `json:"MediaTypeSet,omitempty" name:"MediaTypeSet"`

	// 待筛选的任务状态列表
	TaskStatusSet []*int64 `json:"TaskStatusSet,omitempty" name:"TaskStatusSet"`

	// 待筛选的任务名称数组
	TaskNameSet []*string `json:"TaskNameSet,omitempty" name:"TaskNameSet"`

	// TaskId数组
	TaskIdSet []*string `json:"TaskIdSet,omitempty" name:"TaskIdSet"`

	// 媒资文件名数组
	MediaNameSet []*string `json:"MediaNameSet,omitempty" name:"MediaNameSet"`

	// 媒资语言类型
	MediaLangSet []*int64 `json:"MediaLangSet,omitempty" name:"MediaLangSet"`

	// 媒资素材一级类型
	MediaLabelSet []*int64 `json:"MediaLabelSet,omitempty" name:"MediaLabelSet"`
}

type TaskInfo struct {

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 描述任务名称，指定后可根据名称筛选
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 媒资文件ID
	MediaId *string `json:"MediaId,omitempty" name:"MediaId"`

	// 任务执行状态
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 任务进度，范围为[0，100]
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskProgress *float64 `json:"TaskProgress,omitempty" name:"TaskProgress"`

	// 任务执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTimeCost *int64 `json:"TaskTimeCost,omitempty" name:"TaskTimeCost"`

	// 任务创建时间
	TaskCreateTime *string `json:"TaskCreateTime,omitempty" name:"TaskCreateTime"`

	// 任务开始执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStartTime *string `json:"TaskStartTime,omitempty" name:"TaskStartTime"`

	// 任务失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedReason *string `json:"FailedReason,omitempty" name:"FailedReason"`

	// 任务执行时指定的先验知识
	MediaPreknownInfo *MediaPreknownInfo `json:"MediaPreknownInfo,omitempty" name:"MediaPreknownInfo"`

	// 媒资文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`
}

type TextAppearInfo struct {

	// 文本结果数组中的下标
	Index *int64 `json:"Index,omitempty" name:"Index"`

	// 关键词在文本中出现的起始偏移量(包含)
	StartPosition *int64 `json:"StartPosition,omitempty" name:"StartPosition"`

	// 关键词在文本中出现的结束偏移量(不包含)
	EndPosition *int64 `json:"EndPosition,omitempty" name:"EndPosition"`
}

type TextInfo struct {

	// OCR提取的内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// OCR起始时间戳，从0开始
	StartTimeStamp *float64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// OCR结束时间戳，从0开始
	EndTimeStamp *float64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// OCR标签信息
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

type VideoAppearInfo struct {

	// 视觉信息起始时间戳，从0开始
	StartTimeStamp *float64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 视觉信息终止时间戳，从0开始
	// 关键词在视觉信息中的区间为[StartTimeStamp, EndTimeStamp)
	EndTimeStamp *float64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 关键词在视觉信息中的封面图片
	ImageURL *string `json:"ImageURL,omitempty" name:"ImageURL"`
}
