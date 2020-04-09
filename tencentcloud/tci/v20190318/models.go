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

package v20190318

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AIAssistantRequest struct {
	*tchttp.BaseRequest

	// 输入分析对象内容，输入数据格式参考FileType参数释义
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 输入分析对象类型，picture_url:图片地址，vod_url:视频地址，live_url：直播地址，audio_url: 音频文件，picture：图片二进制数据的BASE64编码
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 音频源的语言，默认0为英文，1为中文
	Lang *int64 `json:"Lang,omitempty" name:"Lang"`

	// 查询人员库列表
	LibrarySet []*string `json:"LibrarySet,omitempty" name:"LibrarySet" list`

	// 视频评估时间，单位秒，点播场景默认值为2小时（无法探测长度时）或完整视频，直播场景默认值为10分钟或直播提前结束
	MaxVideoDuration *int64 `json:"MaxVideoDuration,omitempty" name:"MaxVideoDuration"`

	// 标准化模板选择：0：AI助教基础版本，1：AI评教基础版本，2：AI评教标准版本。AI 助教基础版本功能包括：人脸检索、人脸检测、人脸表情识别、学生动作选项，音频信息分析，微笑识别。AI 评教基础版本功能包括：人脸检索、人脸检测、人脸表情识别、音频信息分析。AI 评教标准版功能包括人脸检索、人脸检测、人脸表情识别、手势识别、音频信息分析、音频关键词分析、视频精彩集锦分析。
	Template *int64 `json:"Template,omitempty" name:"Template"`

	// 识别词库名列表，评估过程使用这些词汇库中的词汇进行词汇使用行为分析
	VocabLibNameList []*string `json:"VocabLibNameList,omitempty" name:"VocabLibNameList" list`

	// 语音编码类型 1:pcm
	VoiceEncodeType *int64 `json:"VoiceEncodeType,omitempty" name:"VoiceEncodeType"`

	// 语音文件类型 1:raw, 2:wav, 3:mp3，10:视频（三种音频格式目前仅支持16k采样率16bit）
	VoiceFileType *int64 `json:"VoiceFileType,omitempty" name:"VoiceFileType"`
}

func (r *AIAssistantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AIAssistantRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AIAssistantResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图像任务直接返回结果
		ImageResults []*ImageTaskResult `json:"ImageResults,omitempty" name:"ImageResults" list`

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AIAssistantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AIAssistantResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ASRStat struct {

	// 当前音频的平均语速
	AvgSpeed *float64 `json:"AvgSpeed,omitempty" name:"AvgSpeed"`

	// Vad的平均音量
	AvgVolume *float64 `json:"AvgVolume,omitempty" name:"AvgVolume"`

	// Vad的最大音量
	MaxVolume *float64 `json:"MaxVolume,omitempty" name:"MaxVolume"`

	// Vad的最小音量
	MinVolume *float64 `json:"MinVolume,omitempty" name:"MinVolume"`

	// 当前音频的非发音时长
	MuteDuration *int64 `json:"MuteDuration,omitempty" name:"MuteDuration"`

	// 当前音频的发音时长
	SoundDuration *int64 `json:"SoundDuration,omitempty" name:"SoundDuration"`

	// 当前音频的总时长
	TotalDuration *int64 `json:"TotalDuration,omitempty" name:"TotalDuration"`

	// 当前音频的句子总数
	VadNum *int64 `json:"VadNum,omitempty" name:"VadNum"`

	// 当前音频的单词总数
	WordNum *int64 `json:"WordNum,omitempty" name:"WordNum"`
}

type AbsenceInfo struct {

	// 识别到的人员所在的库id
	LibraryIds *string `json:"LibraryIds,omitempty" name:"LibraryIds"`

	// 识别到的人员id
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

type ActionCountStatistic struct {

	// 数量
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ActionDurationRatioStatistic struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 比例
	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`
}

type ActionDurationStatistic struct {

	// 时长
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ActionInfo struct {

	// 躯体动作识别结果，包含坐着（sit）、站立（stand）和趴睡（sleep）
	BodyPosture *ActionType `json:"BodyPosture,omitempty" name:"BodyPosture"`

	// 举手识别结果，包含举手（hand）和未检测到举手（nothand）
	Handup *ActionType `json:"Handup,omitempty" name:"Handup"`

	// 是否低头识别结果，包含抬头（lookingahead）和未检测到抬头（notlookingahead）
	LookHead *ActionType `json:"LookHead,omitempty" name:"LookHead"`

	// 是否写字识别结果，包含写字（write）和未检测到写字（notlookingahead）
	Writing *ActionType `json:"Writing,omitempty" name:"Writing"`

	// 动作图像高度
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 动作出现图像的左侧起始坐标位置
	Left *int64 `json:"Left,omitempty" name:"Left"`

	// 动作出现图像的上侧起始侧坐标位置
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// 动作图像宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`
}

type ActionStatistic struct {

	// 数量统计
	ActionCount []*ActionCountStatistic `json:"ActionCount,omitempty" name:"ActionCount" list`

	// 时长统计
	ActionDuration []*ActionDurationStatistic `json:"ActionDuration,omitempty" name:"ActionDuration" list`

	// 时长比例统计
	ActionDurationRatio []*ActionDurationRatioStatistic `json:"ActionDurationRatio,omitempty" name:"ActionDurationRatio" list`
}

type ActionType struct {

	// 置信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 动作类别
	Type *string `json:"Type,omitempty" name:"Type"`
}

type AllMuteSlice struct {

	// 所有静音片段。
	MuteSlice []*MuteSlice `json:"MuteSlice,omitempty" name:"MuteSlice" list`

	// 静音时长占比。
	MuteRatio *float64 `json:"MuteRatio,omitempty" name:"MuteRatio"`

	// 静音总时长。
	TotalMuteDuration *int64 `json:"TotalMuteDuration,omitempty" name:"TotalMuteDuration"`
}

type AttendanceInfo struct {

	// 识别到的人员信息
	Face *FrameInfo `json:"Face,omitempty" name:"Face"`

	// 识别到的人员id
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

type BodyMovementResult struct {

	// 置信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 识别结果高度
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 识别结果左坐标
	Left *int64 `json:"Left,omitempty" name:"Left"`

	// 老师动作识别结果，包含
	// 1、teach_on_positive_attitude 正面讲解
	// 2、point_to_the_blackboard 指黑板
	// 3、writing_blackboard 写板书
	// 4、other 其他
	Movements *string `json:"Movements,omitempty" name:"Movements"`

	// 识别结果顶坐标
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// 识别结果宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`
}

type CancelTaskRequest struct {
	*tchttp.BaseRequest

	// 待取消任务标志符。
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`
}

func (r *CancelTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 取消任务标志符。
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckFacePhotoRequest struct {
	*tchttp.BaseRequest

	// 输入分析对象内容
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 输入分析对象类型，picture_url:图片地址
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

func (r *CheckFacePhotoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckFacePhotoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckFacePhotoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人脸检查结果，0：通过检查，1：图片模糊
		CheckResult *int64 `json:"CheckResult,omitempty" name:"CheckResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckFacePhotoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckFacePhotoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFaceRequest struct {
	*tchttp.BaseRequest

	// 人员唯一标识符
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 图片数据 base64 字符串，与 Urls 参数选择一个输入
	Images []*string `json:"Images,omitempty" name:"Images" list`

	// 人员库唯一标识符
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 图片下载地址，与 Images 参数选择一个输入
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *CreateFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人脸操作结果信息
		FaceInfoSet []*FaceInfo `json:"FaceInfoSet,omitempty" name:"FaceInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLibraryRequest struct {
	*tchttp.BaseRequest

	// 人员库名称
	LibraryName *string `json:"LibraryName,omitempty" name:"LibraryName"`

	// 人员库唯一标志符，为空则系统自动生成。
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`
}

func (r *CreateLibraryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLibraryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLibraryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人员库唯一标识符
		LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

		// 人员库名称
		LibraryName *string `json:"LibraryName,omitempty" name:"LibraryName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLibraryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLibraryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePersonRequest struct {
	*tchttp.BaseRequest

	// 人员库唯一标识符
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 人员名称
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 图片数据 base64 字符串，与 Urls 参数选择一个输入
	Images []*string `json:"Images,omitempty" name:"Images" list`

	// 人员工作号码
	JobNumber *string `json:"JobNumber,omitempty" name:"JobNumber"`

	// 人员邮箱
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 人员性别，0：未知 1：男性，2：女性
	Male *int64 `json:"Male,omitempty" name:"Male"`

	// 自定义人员 ID，注意不能使用 tci_person_ 前缀
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人员电话号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 人员学生号码
	StudentNumber *string `json:"StudentNumber,omitempty" name:"StudentNumber"`

	// 图片下载地址，与 Images 参数选择一个输入
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *CreatePersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人脸操作结果信息
		FaceInfoSet []*FaceInfo `json:"FaceInfoSet,omitempty" name:"FaceInfoSet" list`

		// 人员库唯一标识符
		LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

		// 人员唯一标识符
		PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

		// 人员名称
		PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVocabLibRequest struct {
	*tchttp.BaseRequest

	// 词汇库名称
	VocabLibName *string `json:"VocabLibName,omitempty" name:"VocabLibName"`
}

func (r *CreateVocabLibRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVocabLibRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVocabLibResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVocabLibResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVocabLibResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVocabRequest struct {
	*tchttp.BaseRequest

	// 要添加词汇的词汇库名
	VocabLibName *string `json:"VocabLibName,omitempty" name:"VocabLibName"`

	// 要添加的词汇列表
	VocabList []*string `json:"VocabList,omitempty" name:"VocabList" list`
}

func (r *CreateVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVocabRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVocabResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVocabResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFaceRequest struct {
	*tchttp.BaseRequest

	// 人脸标识符数组
	FaceIdSet []*string `json:"FaceIdSet,omitempty" name:"FaceIdSet" list`

	// 人员唯一标识符
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人员库唯一标识符
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`
}

func (r *DeleteFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人脸操作结果
		FaceInfoSet []*FaceInfo `json:"FaceInfoSet,omitempty" name:"FaceInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLibraryRequest struct {
	*tchttp.BaseRequest

	// 人员库唯一标识符
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`
}

func (r *DeleteLibraryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLibraryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLibraryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人员库唯一标识符
		LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

		// 人员库名称
		LibraryName *string `json:"LibraryName,omitempty" name:"LibraryName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLibraryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLibraryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePersonRequest struct {
	*tchttp.BaseRequest

	// 人员库唯一标识符
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 人员唯一标识符
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

func (r *DeletePersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人脸信息
		FaceInfoSet []*FaceInfo `json:"FaceInfoSet,omitempty" name:"FaceInfoSet" list`

		// 人员库唯一标识符
		LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

		// 人员唯一标识符
		PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

		// 人员名称
		PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVocabLibRequest struct {
	*tchttp.BaseRequest

	// 词汇库名称
	VocabLibName *string `json:"VocabLibName,omitempty" name:"VocabLibName"`
}

func (r *DeleteVocabLibRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVocabLibRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVocabLibResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVocabLibResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVocabLibResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVocabRequest struct {
	*tchttp.BaseRequest

	// 要删除词汇的词汇库名
	VocabLibName *string `json:"VocabLibName,omitempty" name:"VocabLibName"`

	// 要删除的词汇列表
	VocabList []*string `json:"VocabList,omitempty" name:"VocabList" list`
}

func (r *DeleteVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVocabRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVocabResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVocabResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAITaskResultRequest struct {
	*tchttp.BaseRequest

	// 任务唯一标识符。在URL方式时提交请求后会返回一个任务标识符，后续查询该url的结果时使用这个标识符进行查询。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 限制数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAITaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAITaskResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAITaskResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 音频分析结果
		AudioResult *StandardAudioResult `json:"AudioResult,omitempty" name:"AudioResult"`

		// 图像分析结果
		ImageResult *StandardImageResult `json:"ImageResult,omitempty" name:"ImageResult"`

		// 视频分析结果
		VideoResult *StandardVideoResult `json:"VideoResult,omitempty" name:"VideoResult"`

		// 任务状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 任务唯一id。在URL方式时提交请求后会返回一个jobid，后续查询该url的结果时使用这个jobid进行查询。
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAITaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAITaskResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttendanceResultRequest struct {
	*tchttp.BaseRequest

	// 任务唯一标识符
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeAttendanceResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttendanceResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAttendanceResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 缺失人员的ID列表(只针对请求中的libids字段)
		AbsenceSetInLibs []*AbsenceInfo `json:"AbsenceSetInLibs,omitempty" name:"AbsenceSetInLibs" list`

		// 确定出勤人员列表
		AttendanceSet []*AttendanceInfo `json:"AttendanceSet,omitempty" name:"AttendanceSet" list`

		// 疑似出勤人员列表
		SuspectedSet []*SuspectedInfo `json:"SuspectedSet,omitempty" name:"SuspectedSet" list`

		// 缺失人员的ID列表(只针对请求中的personids字段)
		AbsenceSet []*string `json:"AbsenceSet,omitempty" name:"AbsenceSet" list`

		// 请求处理进度
		Progress *int64 `json:"Progress,omitempty" name:"Progress"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttendanceResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAttendanceResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAudioTaskRequest struct {
	*tchttp.BaseRequest

	// 音频任务唯一id。在URL方式时提交请求后会返回一个jobid，后续查询该url的结果时使用这个jobid进行查询。
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`

	// 限制数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAudioTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAudioTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAudioTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 如果请求中开启了静音检测开关，则会返回所有的静音片段（静音时长超过阈值的片段）。
		AllMuteSlice *AllMuteSlice `json:"AllMuteSlice,omitempty" name:"AllMuteSlice"`

		// 返回的当前音频的统计信息。当进度为100时返回。
		AsrStat *ASRStat `json:"AsrStat,omitempty" name:"AsrStat"`

		// 返回当前音频流的详细信息，如果是流模式，返回的是对应流的详细信息，如果是 URL模式，返回的是查询的那一段seq对应的音频的详细信息。
		Texts []*WholeTextItem `json:"Texts,omitempty" name:"Texts" list`

		// 返回词汇库中的单词出现的详细时间信息。
		VocabAnalysisDetailInfo []*VocabDetailInfomation `json:"VocabAnalysisDetailInfo,omitempty" name:"VocabAnalysisDetailInfo" list`

		// 返回词汇库中的单词出现的次数信息。
		VocabAnalysisStatInfo []*VocabStatInfomation `json:"VocabAnalysisStatInfo,omitempty" name:"VocabAnalysisStatInfo" list`

		// 返回音频全部文本。
		AllTexts *string `json:"AllTexts,omitempty" name:"AllTexts"`

		// 音频任务唯一id。在URL方式时提交请求后会返回一个jobid，后续查询该url的结果时使用这个jobid进行查询。
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 返回的当前处理进度。
		Progress *float64 `json:"Progress,omitempty" name:"Progress"`

		// 结果总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAudioTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAudioTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConversationTaskRequest struct {
	*tchttp.BaseRequest

	// 音频任务唯一id。在URL方式时提交请求后会返回一个jobid，后续查询该url的结果时使用这个jobid进行查询。
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`

	// 要查询明细的流的身份，1 老师 2 学生
	Identity *int64 `json:"Identity,omitempty" name:"Identity"`

	// 限制数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConversationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConversationTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConversationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的当前音频的统计信息。当进度为100时返回。
		AsrStat *ASRStat `json:"AsrStat,omitempty" name:"AsrStat"`

		// 返回当前音频流的详细信息，如果是流模式，返回的是对应流的详细信息，如果是 URL模式，返回的是查询的那一段seq对应的音频的详细信息。
		Texts []*WholeTextItem `json:"Texts,omitempty" name:"Texts" list`

		// 返回词汇库中的单词出现的详细时间信息。
		VocabAnalysisDetailInfo []*VocabDetailInfomation `json:"VocabAnalysisDetailInfo,omitempty" name:"VocabAnalysisDetailInfo" list`

		// 返回词汇库中的单词出现的次数信息。
		VocabAnalysisStatInfo []*VocabStatInfomation `json:"VocabAnalysisStatInfo,omitempty" name:"VocabAnalysisStatInfo" list`

		// 整个音频流的全部文本
		AllTexts *string `json:"AllTexts,omitempty" name:"AllTexts"`

		// 音频任务唯一id。在URL方式时提交请求后会返回一个jobid，后续查询该url的结果时使用这个jobid进行查询。
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 返回的当前处理进度。
		Progress *float64 `json:"Progress,omitempty" name:"Progress"`

		// 结果总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConversationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConversationTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHighlightResultRequest struct {
	*tchttp.BaseRequest

	// 精彩集锦任务唯一id。在URL方式时提交请求后会返回一个JobId，后续查询该url的结果时使用这个JobId进行查询。
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeHighlightResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHighlightResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHighlightResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 精彩集锦详细信息。
		HighlightsInfo []*HighlightsInfomation `json:"HighlightsInfo,omitempty" name:"HighlightsInfo" list`

		// 精彩集锦任务唯一id。在URL方式时提交请求后会返回一个JobId，后续查询该url的结果时使用这个JobId进行查询。
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 任务的进度百分比，100表示任务已完成。
		Progress *float64 `json:"Progress,omitempty" name:"Progress"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHighlightResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHighlightResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageTaskRequest struct {
	*tchttp.BaseRequest

	// 任务标识符
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`

	// 限制数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeImageTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务处理结果
		ResultSet []*ImageTaskResult `json:"ResultSet,omitempty" name:"ResultSet" list`

		// 任务唯一标识
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 任务执行进度
		Progress *int64 `json:"Progress,omitempty" name:"Progress"`

		// 任务结果数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageTaskStatisticRequest struct {
	*tchttp.BaseRequest

	// 图像任务标识符
	JobId *int64 `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeImageTaskStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageTaskStatisticRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageTaskStatisticResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务统计信息
		Statistic *ImageTaskStatistic `json:"Statistic,omitempty" name:"Statistic"`

		// 图像任务唯一标识符
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageTaskStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageTaskStatisticResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLibrariesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLibrariesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLibrariesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLibrariesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人员库列表
		LibrarySet []*Library `json:"LibrarySet,omitempty" name:"LibrarySet" list`

		// 人员库总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLibrariesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLibrariesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePersonRequest struct {
	*tchttp.BaseRequest

	// 人员库唯一标识符
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 人员唯一标识符
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

func (r *DescribePersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人员人脸列表
		FaceSet []*Face `json:"FaceSet,omitempty" name:"FaceSet" list`

		// 创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 工作号码
		JobNumber *string `json:"JobNumber,omitempty" name:"JobNumber"`

		// 人员库唯一标识符
		LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

		// 邮箱
		Mail *string `json:"Mail,omitempty" name:"Mail"`

		// 性别
		Male *int64 `json:"Male,omitempty" name:"Male"`

		// 人员唯一标识符
		PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

		// 人员名称
		PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

		// 电话号码
		PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

		// 学生号码
		StudentNumber *string `json:"StudentNumber,omitempty" name:"StudentNumber"`

		// 修改时间
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePersonsRequest struct {
	*tchttp.BaseRequest

	// 人员库唯一标识符
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePersonsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePersonsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePersonsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人员列表
		PersonSet []*Person `json:"PersonSet,omitempty" name:"PersonSet" list`

		// 人员总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePersonsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePersonsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVocabLibRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVocabLibRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVocabLibRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVocabLibResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回该appid下的所有词汇库名
		VocabLibNameSet []*string `json:"VocabLibNameSet,omitempty" name:"VocabLibNameSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVocabLibResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVocabLibResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVocabRequest struct {
	*tchttp.BaseRequest

	// 要查询词汇的词汇库名
	VocabLibName *string `json:"VocabLibName,omitempty" name:"VocabLibName"`
}

func (r *DescribeVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVocabRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVocabResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 词汇列表
		VocabNameSet []*string `json:"VocabNameSet,omitempty" name:"VocabNameSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVocabResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetailInfo struct {

	// 单词出现在该音频中的那个句子的时间戳，出现了几次， 就返回对应次数的起始和结束时间戳
	Value []*WordTimePair `json:"Value,omitempty" name:"Value" list`

	// 词汇库中的单词
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type DoubleVideoFunction struct {

	// 片头片尾增加图片开关
	EnableCoverPictures *bool `json:"EnableCoverPictures,omitempty" name:"EnableCoverPictures"`
}

type ExpressRatioStatistic struct {

	// 出现次数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 表情
	Express *string `json:"Express,omitempty" name:"Express"`

	// 该表情时长占所有表情时长的比例
	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`

	// 该表情时长占视频总时长的比例
	RatioUseDuration *float64 `json:"RatioUseDuration,omitempty" name:"RatioUseDuration"`
}

type Face struct {

	// 人脸唯一标识符
	FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

	// 人脸图片 URL
	FaceUrl *string `json:"FaceUrl,omitempty" name:"FaceUrl"`

	// 人员唯一标识符
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

type FaceAttrResult struct {

	// 年龄
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 性别
	Sex *string `json:"Sex,omitempty" name:"Sex"`
}

type FaceDetectStatistic struct {

	// 人脸大小占画面平均占比
	FaceSizeRatio *float64 `json:"FaceSizeRatio,omitempty" name:"FaceSizeRatio"`

	// 检测到正脸次数
	FrontalFaceCount *int64 `json:"FrontalFaceCount,omitempty" name:"FrontalFaceCount"`

	// 正脸时长占比
	FrontalFaceRatio *float64 `json:"FrontalFaceRatio,omitempty" name:"FrontalFaceRatio"`

	// 正脸时长在总出现时常占比
	FrontalFaceRealRatio *float64 `json:"FrontalFaceRealRatio,omitempty" name:"FrontalFaceRealRatio"`

	// 人员唯一标识符
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 检测到侧脸次数
	SideFaceCount *int64 `json:"SideFaceCount,omitempty" name:"SideFaceCount"`

	// 侧脸时长占比
	SideFaceRatio *float64 `json:"SideFaceRatio,omitempty" name:"SideFaceRatio"`

	// 侧脸时长在总出现时常占比
	SideFaceRealRatio *float64 `json:"SideFaceRealRatio,omitempty" name:"SideFaceRealRatio"`
}

type FaceExpressStatistic struct {

	// 人员唯一标识符
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 表情统计结果
	ExpressRatio []*ExpressRatioStatistic `json:"ExpressRatio,omitempty" name:"ExpressRatio" list`
}

type FaceExpressionResult struct {

	// 表情置信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 表情识别结果，包括"neutral":中性,"happiness":开心，"angry":"生气"，"disgust":厌恶，"fear":"恐惧"，"sadness":"悲伤"，"surprise":"惊讶"，"contempt":"蔑视"
	Expression *string `json:"Expression,omitempty" name:"Expression"`
}

type FaceIdentifyResult struct {

	// 人脸标识符
	FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

	// 人员库标识符
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 人员标识符
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 相似度
	Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`
}

type FaceIdentifyStatistic struct {

	// 持续时间
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 结束时间
	EndTs *int64 `json:"EndTs,omitempty" name:"EndTs"`

	// 人员唯一标识符
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 相似度
	Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`

	// 开始时间
	StartTs *int64 `json:"StartTs,omitempty" name:"StartTs"`
}

type FaceInfo struct {

	// 人脸操作错误码
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 人脸操作结果信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 人脸唯一标识符
	FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

	// 人脸保存地址
	FaceUrl *string `json:"FaceUrl,omitempty" name:"FaceUrl"`

	// 人员唯一标识
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

type FaceInfoResult struct {

	// 人脸尺寸的占比
	FaceRatio *float64 `json:"FaceRatio,omitempty" name:"FaceRatio"`

	// 帧高度
	FrameHeight *int64 `json:"FrameHeight,omitempty" name:"FrameHeight"`

	// 帧宽度
	FrameWidth *int64 `json:"FrameWidth,omitempty" name:"FrameWidth"`

	// 人脸高度
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 人脸左坐标
	Left *int64 `json:"Left,omitempty" name:"Left"`

	// 人脸顶坐标
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// 人脸宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`
}

type FacePoseResult struct {

	// 正脸或侧脸的消息
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// 围绕Z轴旋转角度，俯仰角
	Pitch *float64 `json:"Pitch,omitempty" name:"Pitch"`

	// 围绕X轴旋转角度，翻滚角
	Roll *float64 `json:"Roll,omitempty" name:"Roll"`

	// 围绕Y轴旋转角度，偏航角
	Yaw *float64 `json:"Yaw,omitempty" name:"Yaw"`
}

type FrameInfo struct {

	// 相似度
	Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`

	// 截图的存储地址
	SnapshotUrl *string `json:"SnapshotUrl,omitempty" name:"SnapshotUrl"`

	// 相对于视频起始时间的时间戳，单位秒
	Ts *int64 `json:"Ts,omitempty" name:"Ts"`
}

type Function struct {

	// 输出全部文本标识，当该值设置为true时，会输出当前音频的全部文本
	EnableAllText *bool `json:"EnableAllText,omitempty" name:"EnableAllText"`

	// 输出关键词信息标识，当该值设置为true时，会输出当前音频的关键词信息。
	EnableKeyword *bool `json:"EnableKeyword,omitempty" name:"EnableKeyword"`

	// 静音检测标识，当设置为 true 时，需要设置静音时间阈值字段mute_threshold，统计结果中会返回静音片段。
	EnableMuteDetect *bool `json:"EnableMuteDetect,omitempty" name:"EnableMuteDetect"`

	// 输出音频统计信息标识，当设置为 true 时，任务查询结果会输出音频的统计信息（AsrStat）
	EnableVadInfo *bool `json:"EnableVadInfo,omitempty" name:"EnableVadInfo"`

	// 输出音频音量信息标识，当设置为 true 时，会输出当前音频音量信息。
	EnableVolume *bool `json:"EnableVolume,omitempty" name:"EnableVolume"`
}

type GestureResult struct {

	// 识别结果，包含"USPEAK":听你说，"LISTEN":听我说，"GOOD":GOOD，"TOOLS":拿教具，"OTHERS":其他
	Class *string `json:"Class,omitempty" name:"Class"`

	// 置信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 识别结果高度
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 识别结果左坐标
	Left *int64 `json:"Left,omitempty" name:"Left"`

	// 识别结果顶坐标
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// 识别结果宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`
}

type HLFunction struct {

	// 是否开启人脸检测
	EnableFaceDetect *bool `json:"EnableFaceDetect,omitempty" name:"EnableFaceDetect"`

	// 是否开启表情识别
	EnableFaceExpression *bool `json:"EnableFaceExpression,omitempty" name:"EnableFaceExpression"`

	// 是否开启人脸检索
	EnableFaceIdent *bool `json:"EnableFaceIdent,omitempty" name:"EnableFaceIdent"`

	// 是否开启视频集锦-老师关键字识别
	EnableKeywordWonderfulTime *bool `json:"EnableKeywordWonderfulTime,omitempty" name:"EnableKeywordWonderfulTime"`

	// 是否开启视频集锦-微笑识别
	EnableSmileWonderfulTime *bool `json:"EnableSmileWonderfulTime,omitempty" name:"EnableSmileWonderfulTime"`
}

type HandTrackingResult struct {

	// 识别结果
	Class *string `json:"Class,omitempty" name:"Class"`

	// 置信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 识别结果高度
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 识别结果左坐标
	Left *int64 `json:"Left,omitempty" name:"Left"`

	// 识别结果顶坐标
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// 识别结果宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`
}

type HighlightsInfomation struct {

	// 专注的起始与终止时间信息。
	Concentration []*TimeType `json:"Concentration,omitempty" name:"Concentration" list`

	// 微笑的起始与终止时间信息。
	Smile []*TimeType `json:"Smile,omitempty" name:"Smile" list`

	// 高光集锦视频地址，保存剪辑好的视频地址。
	HighlightsUrl *string `json:"HighlightsUrl,omitempty" name:"HighlightsUrl"`

	// 片段中识别出来的人脸ID。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

type ImageTaskFunction struct {

	// 大教室场景学生肢体动作识别选项
	EnableActionClass *bool `json:"EnableActionClass,omitempty" name:"EnableActionClass"`

	// 人脸检测选项（默认为true，目前不可编辑）
	EnableFaceDetect *bool `json:"EnableFaceDetect,omitempty" name:"EnableFaceDetect"`

	// 人脸表情识别选项
	EnableFaceExpression *bool `json:"EnableFaceExpression,omitempty" name:"EnableFaceExpression"`

	// 人脸检索选项（默认为true，目前不可编辑）
	EnableFaceIdentify *bool `json:"EnableFaceIdentify,omitempty" name:"EnableFaceIdentify"`

	// 手势选项
	EnableGesture *bool `json:"EnableGesture,omitempty" name:"EnableGesture"`

	// 优图手势选项（该功能尚未支持）
	EnableHandTracking *bool `json:"EnableHandTracking,omitempty" name:"EnableHandTracking"`

	// 光照选项
	EnableLightJudge *bool `json:"EnableLightJudge,omitempty" name:"EnableLightJudge"`

	// 小班课场景学生肢体动作识别选项
	EnableStudentBodyMovements *bool `json:"EnableStudentBodyMovements,omitempty" name:"EnableStudentBodyMovements"`

	// 教师动作选项（该功能尚未支持）
	EnableTeacherBodyMovements *bool `json:"EnableTeacherBodyMovements,omitempty" name:"EnableTeacherBodyMovements"`

	// 判断老师是否在屏幕中（该功能尚未支持）
	EnableTeacherOutScreen *bool `json:"EnableTeacherOutScreen,omitempty" name:"EnableTeacherOutScreen"`
}

type ImageTaskResult struct {

	// 大教室场景学生肢体动作识别信息
	ActionInfo *ActionInfo `json:"ActionInfo,omitempty" name:"ActionInfo"`

	// 属性识别结果
	FaceAttr *FaceAttrResult `json:"FaceAttr,omitempty" name:"FaceAttr"`

	// 表情识别结果
	FaceExpression *FaceExpressionResult `json:"FaceExpression,omitempty" name:"FaceExpression"`

	// 人脸检索结果
	FaceIdentify *FaceIdentifyResult `json:"FaceIdentify,omitempty" name:"FaceIdentify"`

	// 人脸检测结果
	FaceInfo *FaceInfoResult `json:"FaceInfo,omitempty" name:"FaceInfo"`

	// 姿势识别结果
	FacePose *FacePoseResult `json:"FacePose,omitempty" name:"FacePose"`

	// 动作分类结果
	Gesture *GestureResult `json:"Gesture,omitempty" name:"Gesture"`

	// 手势分类结果
	HandTracking *HandTrackingResult `json:"HandTracking,omitempty" name:"HandTracking"`

	// 光照识别结果
	Light *LightResult `json:"Light,omitempty" name:"Light"`

	// 学生肢体动作识别结果
	StudentBodyMovement *StudentBodyMovementResult `json:"StudentBodyMovement,omitempty" name:"StudentBodyMovement"`

	// 老师肢体动作识别结果
	TeacherBodyMovement *BodyMovementResult `json:"TeacherBodyMovement,omitempty" name:"TeacherBodyMovement"`

	// 教师是否在屏幕内判断结果
	TeacherOutScreen *TeacherOutScreenResult `json:"TeacherOutScreen,omitempty" name:"TeacherOutScreen"`

	// 时间统计结果
	TimeInfo *TimeInfoResult `json:"TimeInfo,omitempty" name:"TimeInfo"`
}

type ImageTaskStatistic struct {

	// 人员检测统计信息
	FaceDetect []*FaceDetectStatistic `json:"FaceDetect,omitempty" name:"FaceDetect" list`

	// 人脸表情统计信息
	FaceExpression []*FaceExpressStatistic `json:"FaceExpression,omitempty" name:"FaceExpression" list`

	// 人脸检索统计信息
	FaceIdentify []*FaceIdentifyStatistic `json:"FaceIdentify,omitempty" name:"FaceIdentify" list`

	// 姿势识别统计信息
	Gesture *ActionStatistic `json:"Gesture,omitempty" name:"Gesture"`

	// 手势识别统计信息
	Handtracking *ActionStatistic `json:"Handtracking,omitempty" name:"Handtracking"`

	// 光照统计信息
	Light *LightStatistic `json:"Light,omitempty" name:"Light"`

	// 学生动作统计信息
	StudentMovement *ActionStatistic `json:"StudentMovement,omitempty" name:"StudentMovement"`

	// 教师动作统计信息
	TeacherMovement *ActionStatistic `json:"TeacherMovement,omitempty" name:"TeacherMovement"`
}

type Library struct {

	// 人员库创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 人员库唯一标识符
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 人员库名称
	LibraryName *string `json:"LibraryName,omitempty" name:"LibraryName"`

	// 人员库人员数量
	PersonCount *int64 `json:"PersonCount,omitempty" name:"PersonCount"`

	// 人员库修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type LightDistributionStatistic struct {

	// 时间点
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 光线值
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type LightLevelRatioStatistic struct {

	// 名称
	Level *string `json:"Level,omitempty" name:"Level"`

	// 比例
	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`
}

type LightResult struct {

	// 光照程度，参考提交任务时的LightStandard指定的Name参数
	LightLevel *string `json:"LightLevel,omitempty" name:"LightLevel"`

	// 光照亮度
	LightValue *float64 `json:"LightValue,omitempty" name:"LightValue"`
}

type LightStandard struct {

	// 光照名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 范围
	Range []*float64 `json:"Range,omitempty" name:"Range" list`
}

type LightStatistic struct {

	// 各个时间点的光线值
	LightDistribution []*LightDistributionStatistic `json:"LightDistribution,omitempty" name:"LightDistribution" list`

	// 光照程度比例统计结果
	LightLevelRatio []*LightLevelRatioStatistic `json:"LightLevelRatio,omitempty" name:"LightLevelRatio" list`
}

type ModifyLibraryRequest struct {
	*tchttp.BaseRequest

	// 人员库唯一标识符
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 人员库名称
	LibraryName *string `json:"LibraryName,omitempty" name:"LibraryName"`
}

func (r *ModifyLibraryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLibraryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLibraryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人员库唯一标识符
		LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

		// 人员库名称
		LibraryName *string `json:"LibraryName,omitempty" name:"LibraryName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLibraryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLibraryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonRequest struct {
	*tchttp.BaseRequest

	// 人员库唯一标识符
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 人员唯一标识符
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人员工作号码
	JobNumber *string `json:"JobNumber,omitempty" name:"JobNumber"`

	// 人员邮箱
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 人员性别
	Male *int64 `json:"Male,omitempty" name:"Male"`

	// 人员名称
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 人员电话号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 人员学生号码
	StudentNumber *string `json:"StudentNumber,omitempty" name:"StudentNumber"`
}

func (r *ModifyPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人脸信息
		FaceInfoSet []*FaceInfo `json:"FaceInfoSet,omitempty" name:"FaceInfoSet" list`

		// 人员所属人员库标识符
		LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

		// 人员唯一标识符
		PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

		// 人员名称
		PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MuteSlice struct {

	// 起始时间。
	MuteBtm *int64 `json:"MuteBtm,omitempty" name:"MuteBtm"`

	// 终止时间。
	MuteEtm *int64 `json:"MuteEtm,omitempty" name:"MuteEtm"`
}

type Person struct {

	// 人员库唯一标识符
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 人员唯一标识符
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人员名称
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 工作号码
	JobNumber *string `json:"JobNumber,omitempty" name:"JobNumber"`

	// 邮箱
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 性别
	Male *int64 `json:"Male,omitempty" name:"Male"`

	// 电话号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 学生号码
	StudentNumber *string `json:"StudentNumber,omitempty" name:"StudentNumber"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type PersonInfo struct {

	// 需要匹配的人员的ID列表。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 视频集锦开始封面照片。
	CoverBeginUrl *string `json:"CoverBeginUrl,omitempty" name:"CoverBeginUrl"`

	// 视频集锦结束封面照片。
	CoverEndUrl *string `json:"CoverEndUrl,omitempty" name:"CoverEndUrl"`
}

type StandardAudioResult struct {

	// 返回的当前音频的统计信息。当进度为100时返回。
	AsrStat *ASRStat `json:"AsrStat,omitempty" name:"AsrStat"`

	// 返回当前音频流的详细信息，如果是流模式，返回的是对应流的详细信息，如果是 URL模式，返回的是查询的那一段seq对应的音频的详细信息。
	Texts []*WholeTextItem `json:"Texts,omitempty" name:"Texts" list`

	// 返回词汇库中的单词出现的详细时间信息。
	VocabAnalysisDetailInfo []*VocabDetailInfomation `json:"VocabAnalysisDetailInfo,omitempty" name:"VocabAnalysisDetailInfo" list`

	// 返回词汇库中的单词出现的次数信息。
	VocabAnalysisStatInfo []*VocabStatInfomation `json:"VocabAnalysisStatInfo,omitempty" name:"VocabAnalysisStatInfo" list`

	// 状态描述
	Message *string `json:"Message,omitempty" name:"Message"`

	// 任务状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 结果数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type StandardImageResult struct {

	// 详细结果
	ResultSet []*ImageTaskResult `json:"ResultSet,omitempty" name:"ResultSet" list`

	// 分析完成后的统计结果
	Statistic *ImageTaskStatistic `json:"Statistic,omitempty" name:"Statistic"`

	// 状态描述
	Message *string `json:"Message,omitempty" name:"Message"`

	// 任务状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 结果总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type StandardVideoResult struct {

	// 分析完成后的统计结果
	HighlightsInfo []*HighlightsInfomation `json:"HighlightsInfo,omitempty" name:"HighlightsInfo" list`

	// 状态描述
	Message *string `json:"Message,omitempty" name:"Message"`

	// 任务状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

type StatInfo struct {

	// 词汇库中的单词
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 单词出现在该音频中总次数
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type StudentBodyMovementResult struct {

	// 置信度（已废弃）
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 举手识别结果置信度
	HandupConfidence *float64 `json:"HandupConfidence,omitempty" name:"HandupConfidence"`

	// 举手识别结果，包含举手（handup）和未举手（nothandup）
	HandupStatus *string `json:"HandupStatus,omitempty" name:"HandupStatus"`

	// 识别结果高度
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 识别结果左坐标
	Left *int64 `json:"Left,omitempty" name:"Left"`

	// 动作识别结果（已废弃）
	Movements *string `json:"Movements,omitempty" name:"Movements"`

	// 站立识别结果置信度
	StandConfidence *float64 `json:"StandConfidence,omitempty" name:"StandConfidence"`

	// 站立识别结果，包含站立（stand）和坐着（sit）
	StandStatus *string `json:"StandStatus,omitempty" name:"StandStatus"`

	// 识别结果顶坐标
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// 识别结果宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`
}

type SubmitAudioTaskRequest struct {
	*tchttp.BaseRequest

	// 音频源的语言，默认0为英文，1为中文
	Lang *int64 `json:"Lang,omitempty" name:"Lang"`

	// 音频URL。客户请求为URL方式时必须带此字段指名音频的url。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 语音编码类型 1:pcm
	VoiceEncodeType *int64 `json:"VoiceEncodeType,omitempty" name:"VoiceEncodeType"`

	// 语音文件类型 1:raw, 2:wav, 3:mp3，10:视频（三种音频格式目前仅支持16k采样率16bit）
	VoiceFileType *int64 `json:"VoiceFileType,omitempty" name:"VoiceFileType"`

	// 功能开关列表，表示是否需要打开相应的功能，返回相应的信息
	Functions *Function `json:"Functions,omitempty" name:"Functions"`

	// 视频文件类型，默认点播，直播填 live_url
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 静音阈值设置，如果静音检测开关开启，则静音时间超过这个阈值认为是静音片段，在结果中会返回, 没给的话默认值为3s
	MuteThreshold *int64 `json:"MuteThreshold,omitempty" name:"MuteThreshold"`

	// 识别词库名列表，评估过程使用这些词汇库中的词汇进行词汇使用行为分析
	VocabLibNameList []*string `json:"VocabLibNameList,omitempty" name:"VocabLibNameList" list`
}

func (r *SubmitAudioTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitAudioTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitAudioTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 	查询结果时指名的jobid。在URL方式时提交请求后会返回一个jobid，后续查询该url的结果时使用这个jobid进行查询。
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitAudioTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitAudioTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitCheckAttendanceTaskPlusRequest struct {
	*tchttp.BaseRequest

	// 输入数据
	FileContent []*string `json:"FileContent,omitempty" name:"FileContent" list`

	// 视频流类型，vod_url表示点播URL，live_url表示直播URL，默认vod_url
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 人员库 ID列表
	LibraryIds []*string `json:"LibraryIds,omitempty" name:"LibraryIds" list`

	// 确定出勤阈值；默认为0.92
	AttendanceThreshold *float64 `json:"AttendanceThreshold,omitempty" name:"AttendanceThreshold"`

	// 是否开启陌生人模式，陌生人模式是指在任务中发现的非注册人脸库中的人脸也返回相关统计信息，默认不开启
	EnableStranger *bool `json:"EnableStranger,omitempty" name:"EnableStranger"`

	// 考勤结束时间（到视频的第几秒结束考勤），单位秒；默认为900 
	// 对于直播场景，使用绝对时间戳，单位秒，默认当前时间往后12小时
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 通知回调地址，要求方法为post，application/json格式
	NoticeUrl *string `json:"NoticeUrl,omitempty" name:"NoticeUrl"`

	// 考勤开始时间（从视频的第几秒开始考勤），单位秒；默认为0 
	// 对于直播场景，使用绝对时间戳，单位秒，默认当前时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 识别阈值；默认为0.8
	Threshold *float64 `json:"Threshold,omitempty" name:"Threshold"`
}

func (r *SubmitCheckAttendanceTaskPlusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitCheckAttendanceTaskPlusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitCheckAttendanceTaskPlusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务标识符
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 没有注册的人的ID列表
		NotRegisteredSet *string `json:"NotRegisteredSet,omitempty" name:"NotRegisteredSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitCheckAttendanceTaskPlusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitCheckAttendanceTaskPlusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitCheckAttendanceTaskRequest struct {
	*tchttp.BaseRequest

	// 输入数据
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 视频流类型，vod_url表示点播URL，live_url表示直播URL，默认vod_url
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 人员库 ID列表
	LibraryIds []*string `json:"LibraryIds,omitempty" name:"LibraryIds" list`

	// 确定出勤阈值；默认为0.92
	AttendanceThreshold *float64 `json:"AttendanceThreshold,omitempty" name:"AttendanceThreshold"`

	// 是否开启陌生人模式，陌生人模式是指在任务中发现的非注册人脸库中的人脸也返回相关统计信息，默认不开启
	EnableStranger *bool `json:"EnableStranger,omitempty" name:"EnableStranger"`

	// 考勤结束时间（到视频的第几秒结束考勤），单位秒；默认为900 
	// 对于直播场景，使用绝对时间戳，单位秒，默认当前时间往后12小时
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 通知回调地址，要求方法为post，application/json格式
	NoticeUrl *string `json:"NoticeUrl,omitempty" name:"NoticeUrl"`

	// 考勤开始时间（从视频的第几秒开始考勤），单位秒；默认为0 
	// 对于直播场景，使用绝对时间戳，单位秒，默认当前时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 识别阈值；默认为0.8
	Threshold *float64 `json:"Threshold,omitempty" name:"Threshold"`
}

func (r *SubmitCheckAttendanceTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitCheckAttendanceTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitCheckAttendanceTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务标识符
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 没有注册的人的ID列表
		NotRegisteredSet []*string `json:"NotRegisteredSet,omitempty" name:"NotRegisteredSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitCheckAttendanceTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitCheckAttendanceTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitConversationTaskRequest struct {
	*tchttp.BaseRequest

	// 音频源的语言，默认0为英文，1为中文
	Lang *int64 `json:"Lang,omitempty" name:"Lang"`

	// 学生音频流
	StudentUrl *string `json:"StudentUrl,omitempty" name:"StudentUrl"`

	// 教师音频流
	TeacherUrl *string `json:"TeacherUrl,omitempty" name:"TeacherUrl"`

	// 语音编码类型 1:pcm
	VoiceEncodeType *int64 `json:"VoiceEncodeType,omitempty" name:"VoiceEncodeType"`

	// 语音文件类型 1:raw, 2:wav, 3:mp3（三种格式目前仅支持16k采样率16bit）
	VoiceFileType *int64 `json:"VoiceFileType,omitempty" name:"VoiceFileType"`

	// 功能开关列表，表示是否需要打开相应的功能，返回相应的信息
	Functions *Function `json:"Functions,omitempty" name:"Functions"`

	// 识别词库名列表，评估过程使用这些词汇库中的词汇进行词汇使用行为分析
	VocabLibNameList []*string `json:"VocabLibNameList,omitempty" name:"VocabLibNameList" list`
}

func (r *SubmitConversationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitConversationTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitConversationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 	查询结果时指名的jobid。在URL方式时提交请求后会返回一个jobid，后续查询该url的结果时使用这个jobid进行查询。
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitConversationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitConversationTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitDoubleVideoHighlightsRequest struct {
	*tchttp.BaseRequest

	// 学生视频url
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 需要检索的人脸合集库，不在库中的人脸将不参与精彩集锦；目前仅支持输入一个人脸库。
	LibIds []*string `json:"LibIds,omitempty" name:"LibIds" list`

	// 详细功能开关配置项
	Functions *DoubleVideoFunction `json:"Functions,omitempty" name:"Functions"`

	// 需要匹配的人员信息列表。
	PersonInfoList []*PersonInfo `json:"PersonInfoList,omitempty" name:"PersonInfoList" list`

	// 视频处理的抽帧间隔，单位毫秒。建议留空。
	FrameInterval *int64 `json:"FrameInterval,omitempty" name:"FrameInterval"`

	// 旧版本需要匹配的人员信息列表。
	PersonIds []*string `json:"PersonIds,omitempty" name:"PersonIds" list`

	// 人脸检索的相似度阈值，默认值0.89。建议留空。
	SimThreshold *float64 `json:"SimThreshold,omitempty" name:"SimThreshold"`

	// 老师视频url
	TeacherFileContent *string `json:"TeacherFileContent,omitempty" name:"TeacherFileContent"`
}

func (r *SubmitDoubleVideoHighlightsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitDoubleVideoHighlightsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitDoubleVideoHighlightsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 视频拆条任务ID，用来唯一标识视频拆条任务。
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 未注册的人员ID列表。若出现此项，代表评估出现了问题，输入的PersonId中有不在库中的人员ID。
		NotRegistered []*string `json:"NotRegistered,omitempty" name:"NotRegistered" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitDoubleVideoHighlightsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitDoubleVideoHighlightsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitFullBodyClassTaskRequest struct {
	*tchttp.BaseRequest

	// 输入分析对象内容，输入数据格式参考FileType参数释义
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 输入分析对象类型，picture_url:图片地址，vod_url:视频地址，live_url：直播地址，picture: 图片二进制数据的BASE64编码
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 音频源的语言，默认0为英文，1为中文
	Lang *int64 `json:"Lang,omitempty" name:"Lang"`

	// 查询人员库列表，可填写老师的注册照所在人员库
	LibrarySet []*string `json:"LibrarySet,omitempty" name:"LibrarySet" list`

	// 视频评估时间，单位秒，点播场景默认值为2小时（无法探测长度时）或完整视频，直播场景默认值为10分钟或直播提前结束
	MaxVideoDuration *int64 `json:"MaxVideoDuration,omitempty" name:"MaxVideoDuration"`

	// 识别词库名列表，这些词汇库用来维护关键词，评估老师授课过程中，对这些关键词的使用情况
	VocabLibNameList []*string `json:"VocabLibNameList,omitempty" name:"VocabLibNameList" list`

	// 语音编码类型 1:pcm，当FileType为vod_url或live_url时为必填
	VoiceEncodeType *int64 `json:"VoiceEncodeType,omitempty" name:"VoiceEncodeType"`

	// 语音文件类型 10:视频（三种音频格式目前仅支持16k采样率16bit），当FileType为vod_url或live_url时为必填
	VoiceFileType *int64 `json:"VoiceFileType,omitempty" name:"VoiceFileType"`
}

func (r *SubmitFullBodyClassTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitFullBodyClassTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitFullBodyClassTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图像任务直接返回结果，包括： FaceAttr、 FaceExpression、 FaceIdentify、 FaceInfo、 FacePose、 TeacherBodyMovement、TimeInfo
		ImageResults []*ImageTaskResult `json:"ImageResults,omitempty" name:"ImageResults" list`

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitFullBodyClassTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitFullBodyClassTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitHighlightsRequest struct {
	*tchttp.BaseRequest

	// 表情配置开关项。
	Functions *HLFunction `json:"Functions,omitempty" name:"Functions"`

	// 视频url。
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 视频类型及来源，目前只支持点播类型："vod_url"。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 需要检索的人脸合集库，不在库中的人脸将不参与精彩集锦。
	LibIds []*string `json:"LibIds,omitempty" name:"LibIds" list`

	// 视频处理的抽帧间隔，单位毫秒。建议留空。
	FrameInterval *int64 `json:"FrameInterval,omitempty" name:"FrameInterval"`

	// 关键词语言类型，0为英文，1为中文。
	KeywordsLanguage *int64 `json:"KeywordsLanguage,omitempty" name:"KeywordsLanguage"`

	// 关键词数组，当且仅当Funtions中的EnableKeywordWonderfulTime为true时有意义，匹配相应的关键字。
	KeywordsStrings []*string `json:"KeywordsStrings,omitempty" name:"KeywordsStrings" list`

	// 处理视频的总时长，单位毫秒。该值为0或未设置时，默认值两小时生效；当该值大于视频实际时长时，视频实际时长生效；当该值小于视频实际时长时，该值生效；当获取视频实际时长失败时，若该值设置则生效，否则默认值生效。建议留空。
	MaxVideoDuration *int64 `json:"MaxVideoDuration,omitempty" name:"MaxVideoDuration"`

	// 人脸检索的相似度阈值，默认值0.89。建议留空。
	SimThreshold *float64 `json:"SimThreshold,omitempty" name:"SimThreshold"`
}

func (r *SubmitHighlightsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitHighlightsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitHighlightsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 视频拆条任务ID，用来唯一标识视频拆条任务。
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitHighlightsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitHighlightsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitImageTaskPlusRequest struct {
	*tchttp.BaseRequest

	// 输入分析对象内容，输入数据格式参考FileType参数释义
	FileContent []*string `json:"FileContent,omitempty" name:"FileContent" list`

	// 输入分析对象类型，picture：二进制图片的 base64 编码字符串，picture_url:图片地址，vod_url：视频地址，live_url：直播地址
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 任务控制选项
	Functions *ImageTaskFunction `json:"Functions,omitempty" name:"Functions"`

	// 光照标准列表
	LightStandardSet []*LightStandard `json:"LightStandardSet,omitempty" name:"LightStandardSet" list`

	// 抽帧的时间间隔，单位毫秒，默认值1000，保留字段，当前不支持填写。
	FrameInterval *int64 `json:"FrameInterval,omitempty" name:"FrameInterval"`

	// 查询人员库列表
	LibrarySet []*string `json:"LibrarySet,omitempty" name:"LibrarySet" list`

	// 视频评估时间，单位秒，点播场景默认值为2小时（无法探测长度时）或完整视频，直播场景默认值为10分钟或直播提前结束
	MaxVideoDuration *int64 `json:"MaxVideoDuration,omitempty" name:"MaxVideoDuration"`

	// 人脸识别中的相似度阈值，默认值为0.89，保留字段，当前不支持填写。
	SimThreshold *float64 `json:"SimThreshold,omitempty" name:"SimThreshold"`
}

func (r *SubmitImageTaskPlusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitImageTaskPlusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitImageTaskPlusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别结果
		ResultSet []*ImageTaskResult `json:"ResultSet,omitempty" name:"ResultSet" list`

		// 任务标识符
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 任务进度
		Progress *int64 `json:"Progress,omitempty" name:"Progress"`

		// 结果总数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitImageTaskPlusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitImageTaskPlusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitImageTaskRequest struct {
	*tchttp.BaseRequest

	// 输入分析对象内容，输入数据格式参考FileType参数释义
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 输入分析对象类型，picture：二进制图片的 base64 编码字符串，picture_url:图片地址，vod_url：视频地址，live_url：直播地址
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 任务控制选项
	Functions *ImageTaskFunction `json:"Functions,omitempty" name:"Functions"`

	// 光照标准列表
	LightStandardSet []*LightStandard `json:"LightStandardSet,omitempty" name:"LightStandardSet" list`

	// 结果更新回调地址。
	EventsCallBack *string `json:"EventsCallBack,omitempty" name:"EventsCallBack"`

	// 抽帧的时间间隔，单位毫秒，默认值1000，保留字段，当前不支持填写。
	FrameInterval *int64 `json:"FrameInterval,omitempty" name:"FrameInterval"`

	// 查询人员库列表
	LibrarySet []*string `json:"LibrarySet,omitempty" name:"LibrarySet" list`

	// 视频评估时间，单位秒，点播场景默认值为2小时（无法探测长度时）或完整视频，直播场景默认值为10分钟或直播提前结束
	MaxVideoDuration *int64 `json:"MaxVideoDuration,omitempty" name:"MaxVideoDuration"`

	// 人脸识别中的相似度阈值，默认值为0.89，保留字段，当前不支持填写。
	SimThreshold *float64 `json:"SimThreshold,omitempty" name:"SimThreshold"`
}

func (r *SubmitImageTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitImageTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitImageTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别结果
		ResultSet []*ImageTaskResult `json:"ResultSet,omitempty" name:"ResultSet" list`

		// 任务标识符
		JobId *int64 `json:"JobId,omitempty" name:"JobId"`

		// 任务进度
		Progress *int64 `json:"Progress,omitempty" name:"Progress"`

		// 结果总数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitImageTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitImageTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitOneByOneClassTaskRequest struct {
	*tchttp.BaseRequest

	// 输入分析对象内容，输入数据格式参考FileType参数释义
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 输入分析对象类型，picture_url:图片地址，vod_url:视频地址，live_url：直播地址，picture: 图片二进制数据的BASE64编码
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 音频源的语言，默认0为英文，1为中文 
	Lang *int64 `json:"Lang,omitempty" name:"Lang"`

	// 查询人员库列表，可填写学生的注册照所在人员库
	LibrarySet []*string `json:"LibrarySet,omitempty" name:"LibrarySet" list`

	// 视频评估时间，单位秒，点播场景默认值为2小时（无法探测长度时）或完整视频，直播场景默认值为10分钟或直播提前结束
	MaxVideoDuration *int64 `json:"MaxVideoDuration,omitempty" name:"MaxVideoDuration"`

	// 识别词库名列表，这些词汇库用来维护关键词，评估学生对这些关键词的使用情况
	VocabLibNameList []*string `json:"VocabLibNameList,omitempty" name:"VocabLibNameList" list`

	// 语音编码类型 1:pcm，当FileType为vod_url或live_url时为必填
	VoiceEncodeType *int64 `json:"VoiceEncodeType,omitempty" name:"VoiceEncodeType"`

	// 语音文件类型10:视频（三种音频格式目前仅支持16k采样率16bit），当FileType为vod_url或live_url时为必填
	VoiceFileType *int64 `json:"VoiceFileType,omitempty" name:"VoiceFileType"`
}

func (r *SubmitOneByOneClassTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitOneByOneClassTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitOneByOneClassTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图像任务直接返回结果，包括：FaceAttr、 FaceExpression、 FaceIdentify、 FaceInfo、 FacePose、TimeInfo
		ImageResults []*ImageTaskResult `json:"ImageResults,omitempty" name:"ImageResults" list`

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitOneByOneClassTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitOneByOneClassTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitOpenClassTaskRequest struct {
	*tchttp.BaseRequest

	// 输入分析对象内容，输入数据格式参考FileType参数释义
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 输入分析对象类型，picture_url:图片地址，vod_url:视频地址，live_url：直播地址,picture: 图片二进制数据的BASE64编码
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 查询人员库列表，可填写学生们的注册照所在人员库
	LibrarySet []*string `json:"LibrarySet,omitempty" name:"LibrarySet" list`

	// 视频评估时间，单位秒，点播场景默认值为2小时（无法探测长度时）或完整视频，直播场景默认值为10分钟或直播提前结束
	MaxVideoDuration *int64 `json:"MaxVideoDuration,omitempty" name:"MaxVideoDuration"`
}

func (r *SubmitOpenClassTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitOpenClassTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitOpenClassTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图像任务直接返回结果，包括：FaceAttr、 FaceExpression、 FaceIdentify、 FaceInfo、 FacePose、 StudentBodyMovement、TimeInfo
		ImageResults []*ImageTaskResult `json:"ImageResults,omitempty" name:"ImageResults" list`

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitOpenClassTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitOpenClassTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitPartialBodyClassTaskRequest struct {
	*tchttp.BaseRequest

	// 输入分析对象内容，输入数据格式参考FileType参数释义
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 输入分析对象类型，picture_url:图片地址，vod_url:视频地址，live_url：直播地址，picture: 图片二进制数据的BASE64编码
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 音频源的语言，默认0为英文，1为中文
	Lang *int64 `json:"Lang,omitempty" name:"Lang"`

	// 查询人员库列表，可填写老师的注册照所在人员库
	LibrarySet []*string `json:"LibrarySet,omitempty" name:"LibrarySet" list`

	// 视频评估时间，单位秒，点播场景默认值为2小时（无法探测长度时）或完整视频，直播场景默认值为10分钟或直播提前结束
	MaxVideoDuration *int64 `json:"MaxVideoDuration,omitempty" name:"MaxVideoDuration"`

	// 识别词库名列表，这些词汇库用来维护关键词，评估老师授课过程中，对这些关键词的使用情况
	VocabLibNameList []*string `json:"VocabLibNameList,omitempty" name:"VocabLibNameList" list`

	// 语音编码类型 1:pcm，当FileType为vod_url或live_url时为必填
	VoiceEncodeType *int64 `json:"VoiceEncodeType,omitempty" name:"VoiceEncodeType"`

	// 语音文件类型 10:视频（三种音频格式目前仅支持16k采样率16bit），当FileType为vod_url或live_url时为必填
	VoiceFileType *int64 `json:"VoiceFileType,omitempty" name:"VoiceFileType"`
}

func (r *SubmitPartialBodyClassTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitPartialBodyClassTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitPartialBodyClassTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图像任务直接返回结果，包括： FaceAttr、 FaceExpression、 FaceIdentify、 FaceInfo、 FacePose、 Gesture 、 Light、 TimeInfo
		ImageResults []*ImageTaskResult `json:"ImageResults,omitempty" name:"ImageResults" list`

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitPartialBodyClassTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitPartialBodyClassTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitTraditionalClassTaskRequest struct {
	*tchttp.BaseRequest

	// 输入分析对象内容，输入数据格式参考FileType参数释义
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 输入分析对象类型，picture_url:图片地址，vod_url:视频地址，live_url：直播地址，picture：图片二进制数据的BASE64编码
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 查询人员库列表，可填写学生们的注册照所在人员库
	LibrarySet []*string `json:"LibrarySet,omitempty" name:"LibrarySet" list`

	// 视频评估时间，单位秒，点播场景默认值为2小时（无法探测长度时）或完整视频，直播场景默认值为10分钟或直播提前结束
	MaxVideoDuration *int64 `json:"MaxVideoDuration,omitempty" name:"MaxVideoDuration"`
}

func (r *SubmitTraditionalClassTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitTraditionalClassTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitTraditionalClassTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图像任务直接返回结果，包括： ActionInfo、FaceAttr、 FaceExpression、 FaceIdentify、 FaceInfo、 FacePose、 TimeInfo
		ImageResults []*ImageTaskResult `json:"ImageResults,omitempty" name:"ImageResults" list`

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitTraditionalClassTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SubmitTraditionalClassTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SuspectedInfo struct {

	// TopN匹配信息列表
	FaceSet []*FrameInfo `json:"FaceSet,omitempty" name:"FaceSet" list`

	// 识别到的人员id
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

type TeacherOutScreenResult struct {

	// 动作识别结果，InScreen：在屏幕内
	// OutScreen：不在屏幕内
	Class *string `json:"Class,omitempty" name:"Class"`

	// 识别结果高度
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 识别结果左坐标
	Left *int64 `json:"Left,omitempty" name:"Left"`

	// 识别结果顶坐标
	Top *int64 `json:"Top,omitempty" name:"Top"`

	// 识别结果宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`
}

type TextItem struct {

	// 当前句子包含的所有单词信息
	Words []*Word `json:"Words,omitempty" name:"Words" list`

	// 当前句子的置信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 当前句子语音的起始时间点，单位为ms
	Mbtm *int64 `json:"Mbtm,omitempty" name:"Mbtm"`

	// 当前句子语音的终止时间点，单位为ms
	Metm *int64 `json:"Metm,omitempty" name:"Metm"`

	// 保留参数，暂无意义
	Tag *int64 `json:"Tag,omitempty" name:"Tag"`

	// 当前句子
	Text *string `json:"Text,omitempty" name:"Text"`

	// 当前句子的字节数
	TextSize *int64 `json:"TextSize,omitempty" name:"TextSize"`
}

type TimeInfoResult struct {

	// 持续时间，单位毫秒
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 结束时间戳，单位毫秒
	EndTs *int64 `json:"EndTs,omitempty" name:"EndTs"`

	// 开始时间戳，单位毫秒
	StartTs *int64 `json:"StartTs,omitempty" name:"StartTs"`
}

type TimeType struct {

	// 结束时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 起始时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

type TransmitAudioStreamRequest struct {
	*tchttp.BaseRequest

	// 功能开关列表，表示是否需要打开相应的功能，返回相应的信息
	Functions *Function `json:"Functions,omitempty" name:"Functions"`

	// 流式数据包的序号，从1开始，当IsEnd字段为1后后续序号无意义。
	SeqId *int64 `json:"SeqId,omitempty" name:"SeqId"`

	// 语音段唯一标识，一个完整语音一个SessionId。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 当前数据包数据, 流式模式下数据包大小可以按需设置，在网络良好的情况下，建议设置为0.5k，且必须保证分片帧完整（16bit的数据必须保证音频长度为偶数），编码格式要求为BASE64。
	UserVoiceData *string `json:"UserVoiceData,omitempty" name:"UserVoiceData"`

	// 语音编码类型 1:pcm。
	VoiceEncodeType *int64 `json:"VoiceEncodeType,omitempty" name:"VoiceEncodeType"`

	// 语音文件类型 	1: raw, 2: wav, 3: mp3 (语言文件格式目前仅支持 16k 采样率 16bit 编码单声道，如有不一致可能导致评估不准确或失败)。
	VoiceFileType *int64 `json:"VoiceFileType,omitempty" name:"VoiceFileType"`

	// 是否传输完毕标志，若为0表示未完毕，若为1则传输完毕开始评估，非流式模式下无意义。
	IsEnd *int64 `json:"IsEnd,omitempty" name:"IsEnd"`

	// 音频源的语言，默认0为英文，1为中文
	Lang *int64 `json:"Lang,omitempty" name:"Lang"`

	// 是否临时保存 音频链接
	StorageMode *int64 `json:"StorageMode,omitempty" name:"StorageMode"`

	// 识别词库名列表，评估过程使用这些词汇库中的词汇进行词汇使用行为分析
	VocabLibNameList []*string `json:"VocabLibNameList,omitempty" name:"VocabLibNameList" list`
}

func (r *TransmitAudioStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransmitAudioStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TransmitAudioStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的当前音频的统计信息。当进度为100时返回。
		AsrStat *ASRStat `json:"AsrStat,omitempty" name:"AsrStat"`

		// 返回当前音频流的详细信息，如果是流模式，返回的是对应流的详细信息，如果是 URL模式，返回的是查询的那一段seq对应的音频的详细信息。
		Texts []*WholeTextItem `json:"Texts,omitempty" name:"Texts" list`

		// 返回词汇库中的单词出现的详细时间信息。
		VocabAnalysisDetailInfo []*VocabDetailInfomation `json:"VocabAnalysisDetailInfo,omitempty" name:"VocabAnalysisDetailInfo" list`

		// 返回词汇库中的单词出现的次数信息。
		VocabAnalysisStatInfo []*VocabStatInfomation `json:"VocabAnalysisStatInfo,omitempty" name:"VocabAnalysisStatInfo" list`

		// 音频全部文本。
		AllTexts *string `json:"AllTexts,omitempty" name:"AllTexts"`

		// 临时保存的音频链接
		AudioUrl *string `json:"AudioUrl,omitempty" name:"AudioUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransmitAudioStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransmitAudioStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VocabDetailInfomation struct {

	// 词汇库中的单词出现在该音频中的那个句子的时间戳，出现了几次，就返回对应次数的起始和结束时间戳
	VocabDetailInfo []*DetailInfo `json:"VocabDetailInfo,omitempty" name:"VocabDetailInfo" list`

	// 词汇库名
	VocabLibName *string `json:"VocabLibName,omitempty" name:"VocabLibName"`
}

type VocabStatInfomation struct {

	// 单词出现在该音频中总次数
	VocabDetailInfo []*StatInfo `json:"VocabDetailInfo,omitempty" name:"VocabDetailInfo" list`

	// 词汇库名称
	VocabLibName *string `json:"VocabLibName,omitempty" name:"VocabLibName"`
}

type WholeTextItem struct {

	// 当前句子的信息
	TextItem *TextItem `json:"TextItem,omitempty" name:"TextItem"`

	// Vad的平均音量
	AvgVolume *float64 `json:"AvgVolume,omitempty" name:"AvgVolume"`

	// Vad的最大音量
	MaxVolume *float64 `json:"MaxVolume,omitempty" name:"MaxVolume"`

	// Vad的最小音量
	MinVolume *float64 `json:"MinVolume,omitempty" name:"MinVolume"`

	// 当前句子的语速
	Speed *float64 `json:"Speed,omitempty" name:"Speed"`
}

type Word struct {

	// 当前词的置信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 当前单词语音的起始时间点，单位为ms
	Mbtm *int64 `json:"Mbtm,omitempty" name:"Mbtm"`

	// 当前单词语音的终止时间点，单位为ms
	Metm *int64 `json:"Metm,omitempty" name:"Metm"`

	// 当前词
	Text *string `json:"Text,omitempty" name:"Text"`

	// 当前词的字节数
	Wsize *int64 `json:"Wsize,omitempty" name:"Wsize"`
}

type WordTimePair struct {

	// 单词出现的那个句子的起始时间
	Mbtm *int64 `json:"Mbtm,omitempty" name:"Mbtm"`

	// 	单词出现的那个句子的结束时间
	Metm *int64 `json:"Metm,omitempty" name:"Metm"`
}
