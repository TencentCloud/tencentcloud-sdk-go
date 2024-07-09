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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddCustomPersonImageRequestParams struct {
	// 自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 自定义人物图片地址
	ImageURL *string `json:"ImageURL,omitnil,omitempty" name:"ImageURL"`

	// 图片数据base64之后的结果
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`
}

type AddCustomPersonImageRequest struct {
	*tchttp.BaseRequest
	
	// 自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 自定义人物图片地址
	ImageURL *string `json:"ImageURL,omitnil,omitempty" name:"ImageURL"`

	// 图片数据base64之后的结果
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`
}

func (r *AddCustomPersonImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCustomPersonImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "ImageURL")
	delete(f, "Image")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCustomPersonImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCustomPersonImageResponseParams struct {
	// 自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 自定义人脸图片信息
	ImageInfo *PersonImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddCustomPersonImageResponse struct {
	*tchttp.BaseResponse
	Response *AddCustomPersonImageResponseParams `json:"Response"`
}

func (r *AddCustomPersonImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCustomPersonImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppearIndexPair struct {
	// 出现信息，取值范围为[1，3]
	AppearIndex *int64 `json:"AppearIndex,omitnil,omitempty" name:"AppearIndex"`

	// AppearInfo中AppearIndex对应元素的第Index元素，从0开始计数
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}

type AppearInfo struct {
	// 关键词在音频文本结果中的出现位置数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioAppearSet []*TextAppearInfo `json:"AudioAppearSet,omitnil,omitempty" name:"AudioAppearSet"`

	// 关键词在可视文本结果中的出现位置数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextAppearSet []*TextAppearInfo `json:"TextAppearSet,omitnil,omitempty" name:"TextAppearSet"`

	// 关键词在视频信息中的出现位置数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoAppearSet []*VideoAppearInfo `json:"VideoAppearSet,omitnil,omitempty" name:"VideoAppearSet"`
}

type AsrResult struct {
	// ASR提取的文字信息
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// ASR起始时间戳，从0开始
	StartTimeStamp *float64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// ASR结束时间戳，从0开始
	EndTimeStamp *float64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`
}

type AudioData struct {
	// 音频识别文本结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioInfoSet []*AudioInfo `json:"AudioInfoSet,omitnil,omitempty" name:"AudioInfoSet"`

	// 音频识别标签数据
	TextTagSet *MultiLevelTag `json:"TextTagSet,omitnil,omitempty" name:"TextTagSet"`
}

type AudioInfo struct {
	// ASR提取的文字信息
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// ASR起始时间戳，从0开始
	StartTimeStamp *float64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// ASR结束时间戳，从0开始
	EndTimeStamp *float64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// ASR提取的音频标签
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

type AudioMetadata struct {
	// 媒资音频文件大小，单位为Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 媒资音频文件MD5
	// 注意：此字段可能返回 null，表示取不到有效值。
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`

	// 媒资音频时长，单位为秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 媒资音频采样率，单位为khz
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleRate *float64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 媒资音频码率，单位为kbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	BitRate *int64 `json:"BitRate,omitnil,omitempty" name:"BitRate"`

	// 媒资音频文件格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type ClassifiedPersonInfo struct {
	// 人物分类名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassifyName *string `json:"ClassifyName,omitnil,omitempty" name:"ClassifyName"`

	// 符合特定分类的人物信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersonInfoSet []*PersonInfo `json:"PersonInfoSet,omitnil,omitempty" name:"PersonInfoSet"`
}

// Predefined struct for user
type CreateCustomCategoryRequestParams struct {
	// 自定义一级类型
	L1Category *string `json:"L1Category,omitnil,omitempty" name:"L1Category"`

	// 自定义二级类型
	L2Category *string `json:"L2Category,omitnil,omitempty" name:"L2Category"`
}

type CreateCustomCategoryRequest struct {
	*tchttp.BaseRequest
	
	// 自定义一级类型
	L1Category *string `json:"L1Category,omitnil,omitempty" name:"L1Category"`

	// 自定义二级类型
	L2Category *string `json:"L2Category,omitnil,omitempty" name:"L2Category"`
}

func (r *CreateCustomCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "L1Category")
	delete(f, "L2Category")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomCategoryResponseParams struct {
	// 自定义分类信息ID
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomCategoryResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomCategoryResponseParams `json:"Response"`
}

func (r *CreateCustomCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomGroupRequestParams struct {
	// 人脸图片COS存储桶Host地址
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`
}

type CreateCustomGroupRequest struct {
	*tchttp.BaseRequest
	
	// 人脸图片COS存储桶Host地址
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`
}

func (r *CreateCustomGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Bucket")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomGroupResponseParams `json:"Response"`
}

func (r *CreateCustomGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomPersonRequestParams struct {
	// 自定义人物姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义人物简要信息(仅用于标记，不支持检索)
	BasicInfo *string `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// 自定义分类ID，如不存在接口会报错
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 自定义人物图片URL，可支持任意地址，推荐使用COS
	ImageURL *string `json:"ImageURL,omitnil,omitempty" name:"ImageURL"`

	// 原始图片base64编码后的数据
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`
}

type CreateCustomPersonRequest struct {
	*tchttp.BaseRequest
	
	// 自定义人物姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义人物简要信息(仅用于标记，不支持检索)
	BasicInfo *string `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// 自定义分类ID，如不存在接口会报错
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 自定义人物图片URL，可支持任意地址，推荐使用COS
	ImageURL *string `json:"ImageURL,omitnil,omitempty" name:"ImageURL"`

	// 原始图片base64编码后的数据
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`
}

func (r *CreateCustomPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomPersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "BasicInfo")
	delete(f, "CategoryId")
	delete(f, "ImageURL")
	delete(f, "Image")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomPersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomPersonResponseParams struct {
	// 自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 自定义人脸信息
	ImageInfo *PersonImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomPersonResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomPersonResponseParams `json:"Response"`
}

func (r *CreateCustomPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomPersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefaultCategoriesRequestParams struct {

}

type CreateDefaultCategoriesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateDefaultCategoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultCategoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDefaultCategoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefaultCategoriesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDefaultCategoriesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDefaultCategoriesResponseParams `json:"Response"`
}

func (r *CreateDefaultCategoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultCategoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskRequestParams struct {
	// 媒资文件ID，最长32B
	MediaId *string `json:"MediaId,omitnil,omitempty" name:"MediaId"`

	// 媒资素材先验知识，相关限制参考MediaPreknownInfo
	MediaPreknownInfo *MediaPreknownInfo `json:"MediaPreknownInfo,omitnil,omitempty" name:"MediaPreknownInfo"`

	// 任务名称，最长100个中文字符
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 是否上传转码后的视频，仅设置true时上传，默认为false
	UploadVideo *bool `json:"UploadVideo,omitnil,omitempty" name:"UploadVideo"`

	// 自定义标签，可用于查询
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 任务分析完成的回调地址，该设置优先级高于控制台全局的设置；
	CallbackURL *string `json:"CallbackURL,omitnil,omitempty" name:"CallbackURL"`
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 媒资文件ID，最长32B
	MediaId *string `json:"MediaId,omitnil,omitempty" name:"MediaId"`

	// 媒资素材先验知识，相关限制参考MediaPreknownInfo
	MediaPreknownInfo *MediaPreknownInfo `json:"MediaPreknownInfo,omitnil,omitempty" name:"MediaPreknownInfo"`

	// 任务名称，最长100个中文字符
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 是否上传转码后的视频，仅设置true时上传，默认为false
	UploadVideo *bool `json:"UploadVideo,omitnil,omitempty" name:"UploadVideo"`

	// 自定义标签，可用于查询
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 任务分析完成的回调地址，该设置优先级高于控制台全局的设置；
	CallbackURL *string `json:"CallbackURL,omitnil,omitempty" name:"CallbackURL"`
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
	delete(f, "Label")
	delete(f, "CallbackURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskResponseParams struct {
	// 智能标签视频分析任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateVideoSummaryTaskRequestParams struct {
	// 目前只支持 1，表示新闻缩编。
	SummaryType *int64 `json:"SummaryType,omitnil,omitempty" name:"SummaryType"`

	// 待处理的视频的URL，目前只支持*不带签名的*COS地址，长度最长1KB
	VideoURL *string `json:"VideoURL,omitnil,omitempty" name:"VideoURL"`

	// 任务处理完成的回调地址。
	CallbackURL *string `json:"CallbackURL,omitnil,omitempty" name:"CallbackURL"`

	// 如果需要你输出 TTS 或者视频，该字段为转存的cos桶地址且不可为空; 示例：https://${Bucket}-${AppId}.cos.${Region}.myqcloud.com/${PathPrefix}/  (注意，cos路径需要以/分隔符结尾)。
	WriteBackCosPath *string `json:"WriteBackCosPath,omitnil,omitempty" name:"WriteBackCosPath"`

	// 是否开启结果视频生成功能，如果开启，需要指定WriteBackCosPath 参数
	ActiveVideoGenerate *bool `json:"ActiveVideoGenerate,omitnil,omitempty" name:"ActiveVideoGenerate"`

	// 生成结果视频的时候，控制生成的结果视频的横转竖参数。如果 ActiveVideoGenerate 为 false, 该参数无效。
	VideoRotationMode *VideoRotationMode `json:"VideoRotationMode,omitnil,omitempty" name:"VideoRotationMode"`

	// 语音合成相关的控制参数
	TTSMode *TTSMode `json:"TTSMode,omitnil,omitempty" name:"TTSMode"`

	// 是否输出合成好的语音列表。
	ActiveTTSOutput *bool `json:"ActiveTTSOutput,omitnil,omitempty" name:"ActiveTTSOutput"`

	// 用户指定的精确的 asr 结果列表 
	ExactAsrSet []*AsrResult `json:"ExactAsrSet,omitnil,omitempty" name:"ExactAsrSet"`

	// 用户指定的精确的文本摘要
	ExactTextSummary *string `json:"ExactTextSummary,omitnil,omitempty" name:"ExactTextSummary"`

	// 用户指定的精确的文本摘要分割结果
	ExactTextSegSet []*string `json:"ExactTextSegSet,omitnil,omitempty" name:"ExactTextSegSet"`

	// 用户指定的精确的镜头分割结果
	ExactShotSegSet []*ShotInfo `json:"ExactShotSegSet,omitnil,omitempty" name:"ExactShotSegSet"`
}

type CreateVideoSummaryTaskRequest struct {
	*tchttp.BaseRequest
	
	// 目前只支持 1，表示新闻缩编。
	SummaryType *int64 `json:"SummaryType,omitnil,omitempty" name:"SummaryType"`

	// 待处理的视频的URL，目前只支持*不带签名的*COS地址，长度最长1KB
	VideoURL *string `json:"VideoURL,omitnil,omitempty" name:"VideoURL"`

	// 任务处理完成的回调地址。
	CallbackURL *string `json:"CallbackURL,omitnil,omitempty" name:"CallbackURL"`

	// 如果需要你输出 TTS 或者视频，该字段为转存的cos桶地址且不可为空; 示例：https://${Bucket}-${AppId}.cos.${Region}.myqcloud.com/${PathPrefix}/  (注意，cos路径需要以/分隔符结尾)。
	WriteBackCosPath *string `json:"WriteBackCosPath,omitnil,omitempty" name:"WriteBackCosPath"`

	// 是否开启结果视频生成功能，如果开启，需要指定WriteBackCosPath 参数
	ActiveVideoGenerate *bool `json:"ActiveVideoGenerate,omitnil,omitempty" name:"ActiveVideoGenerate"`

	// 生成结果视频的时候，控制生成的结果视频的横转竖参数。如果 ActiveVideoGenerate 为 false, 该参数无效。
	VideoRotationMode *VideoRotationMode `json:"VideoRotationMode,omitnil,omitempty" name:"VideoRotationMode"`

	// 语音合成相关的控制参数
	TTSMode *TTSMode `json:"TTSMode,omitnil,omitempty" name:"TTSMode"`

	// 是否输出合成好的语音列表。
	ActiveTTSOutput *bool `json:"ActiveTTSOutput,omitnil,omitempty" name:"ActiveTTSOutput"`

	// 用户指定的精确的 asr 结果列表 
	ExactAsrSet []*AsrResult `json:"ExactAsrSet,omitnil,omitempty" name:"ExactAsrSet"`

	// 用户指定的精确的文本摘要
	ExactTextSummary *string `json:"ExactTextSummary,omitnil,omitempty" name:"ExactTextSummary"`

	// 用户指定的精确的文本摘要分割结果
	ExactTextSegSet []*string `json:"ExactTextSegSet,omitnil,omitempty" name:"ExactTextSegSet"`

	// 用户指定的精确的镜头分割结果
	ExactShotSegSet []*ShotInfo `json:"ExactShotSegSet,omitnil,omitempty" name:"ExactShotSegSet"`
}

func (r *CreateVideoSummaryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoSummaryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SummaryType")
	delete(f, "VideoURL")
	delete(f, "CallbackURL")
	delete(f, "WriteBackCosPath")
	delete(f, "ActiveVideoGenerate")
	delete(f, "VideoRotationMode")
	delete(f, "TTSMode")
	delete(f, "ActiveTTSOutput")
	delete(f, "ExactAsrSet")
	delete(f, "ExactTextSummary")
	delete(f, "ExactTextSegSet")
	delete(f, "ExactShotSegSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVideoSummaryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoSummaryTaskResponseParams struct {
	// 返回的任务 id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVideoSummaryTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateVideoSummaryTaskResponseParams `json:"Response"`
}

func (r *CreateVideoSummaryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoSummaryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomCategory struct {
	// 自定义分类ID
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 一级自定义类型
	L1Category *string `json:"L1Category,omitnil,omitempty" name:"L1Category"`

	// 二级自定义类型
	L2Category *string `json:"L2Category,omitnil,omitempty" name:"L2Category"`
}

type CustomPersonFilter struct {
	// 待查询的人物姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 待过滤的自定义类型Id数组
	CategoryIdSet []*string `json:"CategoryIdSet,omitnil,omitempty" name:"CategoryIdSet"`

	// 待过滤的自定义人物Id数组
	PersonIdSet []*string `json:"PersonIdSet,omitnil,omitempty" name:"PersonIdSet"`

	// 一级自定义人物类型数组
	L1CategorySet []*string `json:"L1CategorySet,omitnil,omitempty" name:"L1CategorySet"`
}

type CustomPersonInfo struct {
	// 自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 自定义人物姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义人物简介信息
	BasicInfo *string `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// 一级自定义人物类型
	L1Category *string `json:"L1Category,omitnil,omitempty" name:"L1Category"`

	// 二级自定义人物类型
	L2Category *string `json:"L2Category,omitnil,omitempty" name:"L2Category"`

	// 自定义人物图片信息
	ImageInfoSet []*PersonImageInfo `json:"ImageInfoSet,omitnil,omitempty" name:"ImageInfoSet"`

	// 自定义人物创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type Data struct {
	// 节目粒度结构化结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowInfo *ShowInfo `json:"ShowInfo,omitnil,omitempty" name:"ShowInfo"`
}

// Predefined struct for user
type DeleteCustomCategoryRequestParams struct {
	// 自定义分类ID
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`
}

type DeleteCustomCategoryRequest struct {
	*tchttp.BaseRequest
	
	// 自定义分类ID
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`
}

func (r *DeleteCustomCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CategoryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomCategoryResponseParams struct {
	// 123
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCustomCategoryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomCategoryResponseParams `json:"Response"`
}

func (r *DeleteCustomCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomPersonImageRequestParams struct {
	// 自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 自定义人脸图片Id
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

type DeleteCustomPersonImageRequest struct {
	*tchttp.BaseRequest
	
	// 自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 自定义人脸图片Id
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

func (r *DeleteCustomPersonImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomPersonImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "ImageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomPersonImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomPersonImageResponseParams struct {
	// 自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 已删除的人物图片Id
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCustomPersonImageResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomPersonImageResponseParams `json:"Response"`
}

func (r *DeleteCustomPersonImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomPersonImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomPersonRequestParams struct {
	// 待删除的自定义人物ID
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
}

type DeleteCustomPersonRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的自定义人物ID
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
}

func (r *DeleteCustomPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomPersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomPersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomPersonResponseParams struct {
	// 已删除的自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCustomPersonResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomPersonResponseParams `json:"Response"`
}

func (r *DeleteCustomPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomPersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMediaRequestParams struct {
	// 媒资文件在系统中的ID
	MediaId *string `json:"MediaId,omitnil,omitempty" name:"MediaId"`
}

type DeleteMediaRequest struct {
	*tchttp.BaseRequest
	
	// 媒资文件在系统中的ID
	MediaId *string `json:"MediaId,omitnil,omitempty" name:"MediaId"`
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

// Predefined struct for user
type DeleteMediaResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMediaResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMediaResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteTaskRequestParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTaskResponseParams `json:"Response"`
}

func (r *DeleteTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomCategoriesRequestParams struct {

}

type DescribeCustomCategoriesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCustomCategoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomCategoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomCategoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomCategoriesResponseParams struct {
	// 自定义人物类型数组
	CategorySet []*CustomCategory `json:"CategorySet,omitnil,omitempty" name:"CategorySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomCategoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomCategoriesResponseParams `json:"Response"`
}

func (r *DescribeCustomCategoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomCategoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomGroupRequestParams struct {

}

type DescribeCustomGroupRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCustomGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomGroupResponseParams struct {
	// 自定义人物库所包含的人物个数
	GroupSize *int64 `json:"GroupSize,omitnil,omitempty" name:"GroupSize"`

	// 自定义人物库图片后续所在的存储桶
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomGroupResponseParams `json:"Response"`
}

func (r *DescribeCustomGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomPersonDetailRequestParams struct {
	// 自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
}

type DescribeCustomPersonDetailRequest struct {
	*tchttp.BaseRequest
	
	// 自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
}

func (r *DescribeCustomPersonDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomPersonDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomPersonDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomPersonDetailResponseParams struct {
	// 自定义人物信息
	PersonInfo *CustomPersonInfo `json:"PersonInfo,omitnil,omitempty" name:"PersonInfo"`

	// 出现该自定义人物的所有分析人物Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIdSet []*string `json:"TaskIdSet,omitnil,omitempty" name:"TaskIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomPersonDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomPersonDetailResponseParams `json:"Response"`
}

func (r *DescribeCustomPersonDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomPersonDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomPersonsRequestParams struct {
	// 分页序号，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页数据行数，最多50
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序信息，默认倒序
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 自定义人物过滤条件
	Filter *CustomPersonFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCustomPersonsRequest struct {
	*tchttp.BaseRequest
	
	// 分页序号，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页数据行数，最多50
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序信息，默认倒序
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 自定义人物过滤条件
	Filter *CustomPersonFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCustomPersonsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomPersonsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "SortBy")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomPersonsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomPersonsResponseParams struct {
	// 满足过滤条件的自定义人物数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 自定义人物信息
	PersonInfoSet []*CustomPersonInfo `json:"PersonInfoSet,omitnil,omitempty" name:"PersonInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomPersonsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomPersonsResponseParams `json:"Response"`
}

func (r *DescribeCustomPersonsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomPersonsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaRequestParams struct {
	// 导入媒资返回的媒资ID，最长32B
	MediaId *string `json:"MediaId,omitnil,omitempty" name:"MediaId"`
}

type DescribeMediaRequest struct {
	*tchttp.BaseRequest
	
	// 导入媒资返回的媒资ID，最长32B
	MediaId *string `json:"MediaId,omitnil,omitempty" name:"MediaId"`
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

// Predefined struct for user
type DescribeMediaResponseParams struct {
	// 媒资信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaInfo *MediaInfo `json:"MediaInfo,omitnil,omitempty" name:"MediaInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMediaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediaResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeMediasRequestParams struct {
	// 分页序号，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每个分页所包含的元素数量，最大为50
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 列举过滤条件，相关限制相见MediaFilter
	MediaFilter *MediaFilter `json:"MediaFilter,omitnil,omitempty" name:"MediaFilter"`

	// 返回结果排序信息，By字段只支持CreateTime
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`
}

type DescribeMediasRequest struct {
	*tchttp.BaseRequest
	
	// 分页序号，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每个分页所包含的元素数量，最大为50
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 列举过滤条件，相关限制相见MediaFilter
	MediaFilter *MediaFilter `json:"MediaFilter,omitnil,omitempty" name:"MediaFilter"`

	// 返回结果排序信息，By字段只支持CreateTime
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`
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

// Predefined struct for user
type DescribeMediasResponseParams struct {
	// 满足过滤条件的媒资视频总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 满足过滤条件的媒资信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaInfoSet []*MediaInfo `json:"MediaInfoSet,omitnil,omitempty" name:"MediaInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMediasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediasResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTaskDetailRequestParams struct {
	// 创建任务返回的TaskId
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 创建任务返回的TaskId
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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

// Predefined struct for user
type DescribeTaskDetailResponseParams struct {
	// 任务信息，不包含任务结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskInfo *TaskInfo `json:"TaskInfo,omitnil,omitempty" name:"TaskInfo"`

	// 视频任务结果数据，只在视频任务结束时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskData *Data `json:"TaskData,omitnil,omitempty" name:"TaskData"`

	// 图片任务结果数据，只在图片任务结束时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageTaskData *ImageData `json:"ImageTaskData,omitnil,omitempty" name:"ImageTaskData"`

	// 音频任务结果数据，只在音频任务结束时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioTaskData *AudioData `json:"AudioTaskData,omitnil,omitempty" name:"AudioTaskData"`

	// 文本任务结果数据，只在文本任务结束时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTaskData *TextData `json:"TextTaskData,omitnil,omitempty" name:"TextTaskData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTaskRequestParams struct {
	// CreateTask返回的任务ID，最长32B
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest
	
	// CreateTask返回的任务ID，最长32B
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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

// Predefined struct for user
type DescribeTaskResponseParams struct {
	// 任务信息，详情参见TaskInfo的定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskInfo *TaskInfo `json:"TaskInfo,omitnil,omitempty" name:"TaskInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// 分页序号，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每个分页所包含的元素数量，最大为50
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 任务过滤条件，相关限制参见TaskFilter
	TaskFilter *TaskFilter `json:"TaskFilter,omitnil,omitempty" name:"TaskFilter"`

	// 返回结果排序信息，By字段只支持CreateTimeStamp
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 分页序号，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每个分页所包含的元素数量，最大为50
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 任务过滤条件，相关限制参见TaskFilter
	TaskFilter *TaskFilter `json:"TaskFilter,omitnil,omitempty" name:"TaskFilter"`

	// 返回结果排序信息，By字段只支持CreateTimeStamp
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`
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

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 满足过滤条件的任务总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 满足过滤条件的任务数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskInfoSet []*TaskInfo `json:"TaskInfoSet,omitnil,omitempty" name:"TaskInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeUsageAmountRequestParams struct {

}

type DescribeUsageAmountRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUsageAmountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageAmountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsageAmountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageAmountResponseParams struct {
	// 资源使用小时数
	UsedHours *float64 `json:"UsedHours,omitnil,omitempty" name:"UsedHours"`

	// 资源包总量小时数
	TotalHours *float64 `json:"TotalHours,omitnil,omitempty" name:"TotalHours"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsageAmountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsageAmountResponseParams `json:"Response"`
}

func (r *DescribeUsageAmountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageAmountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoSummaryDetailRequestParams struct {
	// 要查询的任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeVideoSummaryDetailRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeVideoSummaryDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoSummaryDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoSummaryDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoSummaryDetailResponseParams struct {
	// 任务的状态
	// 1: 等待处理中
	// 2: 处理中
	// 3: 处理成功
	// 4: 处理失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 如果处理失败，返回失败的原因
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`

	// 提取出的视频的 Asr 结果
	AsrSet []*AsrResult `json:"AsrSet,omitnil,omitempty" name:"AsrSet"`

	// 文本摘要结果
	TextSummary *string `json:"TextSummary,omitnil,omitempty" name:"TextSummary"`

	// 文本摘要分割结果
	TextSegSet []*string `json:"TextSegSet,omitnil,omitempty" name:"TextSegSet"`

	// 镜头分割结果
	ShotSegSet []*ShotInfo `json:"ShotSegSet,omitnil,omitempty" name:"ShotSegSet"`

	// 数组第 i 个结构 TextSegMatchShotConfidenceSet[i] 表示第 i 个文本摘要分割结果和所有镜头的匹配度。
	TextSegMatchShotScoreSet []*TextSegMatchShotScore `json:"TextSegMatchShotScoreSet,omitnil,omitempty" name:"TextSegMatchShotScoreSet"`

	// TTS 输出音频下载地址列表
	TTSResultURLSet []*string `json:"TTSResultURLSet,omitnil,omitempty" name:"TTSResultURLSet"`

	// 合成视频输出下载地址
	VideoResultURL *string `json:"VideoResultURL,omitnil,omitempty" name:"VideoResultURL"`

	// 合成后的视频横竖屏转换后的视频下载地址
	VideoRotateResultURL *string `json:"VideoRotateResultURL,omitnil,omitempty" name:"VideoRotateResultURL"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoSummaryDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoSummaryDetailResponseParams `json:"Response"`
}

func (r *DescribeVideoSummaryDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoSummaryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageData struct {
	// 图片中出现的可视文本识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrSet []*ImageOcr `json:"OcrSet,omitnil,omitempty" name:"OcrSet"`

	// 图片中出现的帧标签识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameTagSet *MultiLevelTag `json:"FrameTagSet,omitnil,omitempty" name:"FrameTagSet"`

	// 图片中出现的层级人物识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiLevelPersonInfoSet []*MultiLevelPersonInfo `json:"MultiLevelPersonInfoSet,omitnil,omitempty" name:"MultiLevelPersonInfoSet"`

	// 图片中出现的台标识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TvLogo *ImageLogo `json:"TvLogo,omitnil,omitempty" name:"TvLogo"`

	// 图片中出现的来源信息识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceLogo *ImageLogo `json:"SourceLogo,omitnil,omitempty" name:"SourceLogo"`
}

type ImageLogo struct {
	// 图片中出现的Logo识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// Logo在图片中出现的位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppearRect *Rectf `json:"AppearRect,omitnil,omitempty" name:"AppearRect"`
}

type ImageMetadata struct {
	// 媒资图片文件大小，单位为Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 媒资图片文件MD5
	// 注意：此字段可能返回 null，表示取不到有效值。
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`

	// 媒资图片文件宽度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 媒资图片文件高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 媒资图片文件格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type ImageOcr struct {
	// 图片中可视文本识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 可视文本在图片中的位置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppearRect *Rectf `json:"AppearRect,omitnil,omitempty" name:"AppearRect"`
}

// Predefined struct for user
type ImportMediaRequestParams struct {
	// 待分析视频的URL，目前只支持*不带签名的*COS地址，长度最长1KB
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 待分析视频的MD5，为空时不做校验，否则会做MD5校验，长度必须为32B
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`

	// 待分析视频的名称，指定后可支持筛选，最多64B
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 当非本人外部视频地址导入时，该字段为转存的cos桶地址且不可为空; 示例：https://${Bucket}-${AppId}.cos.${Region}.myqcloud.com/${PathPrefix}/  (注意，cos路径需要以/分隔符结尾)。
	// 推荐采用本主帐号COS桶，如果使用其他帐号COS桶，请确保COS桶可写，否则可导致分析失败
	WriteBackCosPath *string `json:"WriteBackCosPath,omitnil,omitempty" name:"WriteBackCosPath"`

	// 自定义标签，可用于查询
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 媒资导入完成的回调地址，该设置优先级高于控制台全局的设置；
	CallbackURL *string `json:"CallbackURL,omitnil,omitempty" name:"CallbackURL"`

	// 媒资文件类型，详细定义参见[MediaPreknownInfo.MediaType](https://cloud.tencent.com/document/product/1509/65063#MediaPreknownInfo)
	// 默认为2(视频)
	MediaType *int64 `json:"MediaType,omitnil,omitempty" name:"MediaType"`
}

type ImportMediaRequest struct {
	*tchttp.BaseRequest
	
	// 待分析视频的URL，目前只支持*不带签名的*COS地址，长度最长1KB
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 待分析视频的MD5，为空时不做校验，否则会做MD5校验，长度必须为32B
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`

	// 待分析视频的名称，指定后可支持筛选，最多64B
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 当非本人外部视频地址导入时，该字段为转存的cos桶地址且不可为空; 示例：https://${Bucket}-${AppId}.cos.${Region}.myqcloud.com/${PathPrefix}/  (注意，cos路径需要以/分隔符结尾)。
	// 推荐采用本主帐号COS桶，如果使用其他帐号COS桶，请确保COS桶可写，否则可导致分析失败
	WriteBackCosPath *string `json:"WriteBackCosPath,omitnil,omitempty" name:"WriteBackCosPath"`

	// 自定义标签，可用于查询
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 媒资导入完成的回调地址，该设置优先级高于控制台全局的设置；
	CallbackURL *string `json:"CallbackURL,omitnil,omitempty" name:"CallbackURL"`

	// 媒资文件类型，详细定义参见[MediaPreknownInfo.MediaType](https://cloud.tencent.com/document/product/1509/65063#MediaPreknownInfo)
	// 默认为2(视频)
	MediaType *int64 `json:"MediaType,omitnil,omitempty" name:"MediaType"`
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
	delete(f, "WriteBackCosPath")
	delete(f, "Label")
	delete(f, "CallbackURL")
	delete(f, "MediaType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportMediaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportMediaResponseParams struct {
	// 媒资文件在系统中的ID
	MediaId *string `json:"MediaId,omitnil,omitempty" name:"MediaId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportMediaResponse struct {
	*tchttp.BaseResponse
	Response *ImportMediaResponseParams `json:"Response"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 二级标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	L2TagSet []*L2Tag `json:"L2TagSet,omitnil,omitempty" name:"L2TagSet"`

	// 一级标签出现信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppearIndexPairSet []*AppearIndexPair `json:"AppearIndexPairSet,omitnil,omitempty" name:"AppearIndexPairSet"`

	// 一级标签首次出现信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstAppear *int64 `json:"FirstAppear,omitnil,omitempty" name:"FirstAppear"`
}

type L2Tag struct {
	// 二级标签名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 从属于此二级标签的三级标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	L3TagSet []*L3Tag `json:"L3TagSet,omitnil,omitempty" name:"L3TagSet"`

	// 二级标签出现信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppearIndexPairSet []*AppearIndexPair `json:"AppearIndexPairSet,omitnil,omitempty" name:"AppearIndexPairSet"`

	// 二级标签首次出现信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstAppear *int64 `json:"FirstAppear,omitnil,omitempty" name:"FirstAppear"`
}

type L3Tag struct {
	// 三级标签名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 三级标签出现信息索引数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppearIndexPairSet []*AppearIndexPair `json:"AppearIndexPairSet,omitnil,omitempty" name:"AppearIndexPairSet"`

	// 三级标签首次出现信息，可选值为[1,3]
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstAppear *int64 `json:"FirstAppear,omitnil,omitempty" name:"FirstAppear"`
}

type MediaFilter struct {
	// 媒资名称过滤条件
	MediaNameSet []*string `json:"MediaNameSet,omitnil,omitempty" name:"MediaNameSet"`

	// 媒资状态数组，媒资状态可选值参见MediaInfo
	StatusSet []*int64 `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// 媒资ID数组
	MediaIdSet []*string `json:"MediaIdSet,omitnil,omitempty" name:"MediaIdSet"`

	// 媒资自定义标签数组
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`

	// 媒资文件类型，定义参见[MediaPreknownInfo.MediaType](https://cloud.tencent.com/document/product/1509/65063#MediaPreknownInfo)
	MediaType *int64 `json:"MediaType,omitnil,omitempty" name:"MediaType"`
}

type MediaInfo struct {
	// 媒资ID
	MediaId *string `json:"MediaId,omitnil,omitempty" name:"MediaId"`

	// 媒资名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 媒资下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownLoadURL *string `json:"DownLoadURL,omitnil,omitempty" name:"DownLoadURL"`

	// 媒资状态，取值参看上方表格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 若状态为失败，表示失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`

	// 媒资视频元信息，仅在MediaType=VIDEO时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metadata *MediaMetadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 导入视频进度，取值范围为[0,100]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 媒资自定义标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 媒资导入完成后的回调地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallbackURL *string `json:"CallbackURL,omitnil,omitempty" name:"CallbackURL"`

	// 媒资文件类型，具体参看[MediaPreknownInfo](https://cloud.tencent.com/document/product/1509/65063#MediaPreknownInfo)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaType *int64 `json:"MediaType,omitnil,omitempty" name:"MediaType"`

	// 媒资音频元信息，仅在MediaType=Audio时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioMetadata *AudioMetadata `json:"AudioMetadata,omitnil,omitempty" name:"AudioMetadata"`

	// 媒资图片文件元信息，仅在MediaType=Image时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageMetadata *ImageMetadata `json:"ImageMetadata,omitnil,omitempty" name:"ImageMetadata"`

	// 媒资文本文件元信息，仅在MediaType=Text时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextMetadata *TextMetadata `json:"TextMetadata,omitnil,omitempty" name:"TextMetadata"`
}

type MediaMetadata struct {
	// 媒资视频文件大小，单位为字节
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 媒资视频文件MD5
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`

	// 媒资视频时长，单位为秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 媒资视频总帧数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumFrames *int64 `json:"NumFrames,omitnil,omitempty" name:"NumFrames"`

	// 媒资视频宽度，单位为像素
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 媒资视频高度，单位为像素
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 媒资视频帧率，单位为Hz
	// 注意：此字段可能返回 null，表示取不到有效值。
	FPS *float64 `json:"FPS,omitnil,omitempty" name:"FPS"`

	// 媒资视频比特率，单位为kbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	BitRate *int64 `json:"BitRate,omitnil,omitempty" name:"BitRate"`
}

type MediaPreknownInfo struct {
	// 媒资文件类型，参见MediaPreknownInfo结构体定义
	MediaType *int64 `json:"MediaType,omitnil,omitempty" name:"MediaType"`

	// 媒资素材一级类型，参见MediaPreknownInfo结构体定义
	MediaLabel *int64 `json:"MediaLabel,omitnil,omitempty" name:"MediaLabel"`

	// 媒资素材二级类型，参见MediaPreknownInfo结构体定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaSecondLabel *int64 `json:"MediaSecondLabel,omitnil,omitempty" name:"MediaSecondLabel"`

	// 媒资音频类型，参见MediaPreknownInfo结构体定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaLang *int64 `json:"MediaLang,omitnil,omitempty" name:"MediaLang"`
}

// Predefined struct for user
type ModifyCallbackRequestParams struct {
	// 任务分析完成后回调地址
	TaskFinishNotifyURL *string `json:"TaskFinishNotifyURL,omitnil,omitempty" name:"TaskFinishNotifyURL"`

	// 媒体导入完成后回调地址
	MediaFinishNotifyURL *string `json:"MediaFinishNotifyURL,omitnil,omitempty" name:"MediaFinishNotifyURL"`
}

type ModifyCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 任务分析完成后回调地址
	TaskFinishNotifyURL *string `json:"TaskFinishNotifyURL,omitnil,omitempty" name:"TaskFinishNotifyURL"`

	// 媒体导入完成后回调地址
	MediaFinishNotifyURL *string `json:"MediaFinishNotifyURL,omitnil,omitempty" name:"MediaFinishNotifyURL"`
}

func (r *ModifyCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskFinishNotifyURL")
	delete(f, "MediaFinishNotifyURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCallbackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCallbackResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCallbackResponseParams `json:"Response"`
}

func (r *ModifyCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MultiLevelPersonInfo struct {
	// 一级分类名称(分类信息参见自定义人物类型)
	// 注意：此字段可能返回 null，表示取不到有效值。
	L1ClassifyName *string `json:"L1ClassifyName,omitnil,omitempty" name:"L1ClassifyName"`

	// 已分类人物信息数组(所有分类类型为二级分类)
	// 注意：此字段可能返回 null，表示取不到有效值。
	L2ClassifiedPersonInfoSet []*ClassifiedPersonInfo `json:"L2ClassifiedPersonInfoSet,omitnil,omitempty" name:"L2ClassifiedPersonInfoSet"`

	// 检测结果来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`
}

type MultiLevelTag struct {
	// 树状标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*L1Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// 标签在识别结果中的定位信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppearInfo *AppearInfo `json:"AppearInfo,omitnil,omitempty" name:"AppearInfo"`
}

type PersonImageInfo struct {
	// 人脸图片ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 自定义人脸图片的URL，存储在IVLDCustomPreson存储桶内
	ImageURL *string `json:"ImageURL,omitnil,omitempty" name:"ImageURL"`

	// 自定义人脸图片处理错误码
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 自定义人脸图片处理错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`
}

type PersonInfo struct {
	// 公众人物姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 公众人物职务
	Job *string `json:"Job,omitnil,omitempty" name:"Job"`

	// 首次出现模态，可选值为[1,3]，详细参见AppearIndex定义
	FirstAppear *int64 `json:"FirstAppear,omitnil,omitempty" name:"FirstAppear"`

	// 人物出现信息
	AppearInfo *AppearInfo `json:"AppearInfo,omitnil,omitempty" name:"AppearInfo"`

	// 人脸在图片中的位置，仅在图片标签任务有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppearRect *Rectf `json:"AppearRect,omitnil,omitempty" name:"AppearRect"`
}

// Predefined struct for user
type QueryCallbackRequestParams struct {

}

type QueryCallbackRequest struct {
	*tchttp.BaseRequest
	
}

func (r *QueryCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCallbackResponseParams struct {
	// 任务分析完成后回调地址
	TaskFinishNotifyURL *string `json:"TaskFinishNotifyURL,omitnil,omitempty" name:"TaskFinishNotifyURL"`

	// 媒体导入完成后回调地址
	MediaFinishNotifyURL *string `json:"MediaFinishNotifyURL,omitnil,omitempty" name:"MediaFinishNotifyURL"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryCallbackResponse struct {
	*tchttp.BaseResponse
	Response *QueryCallbackResponseParams `json:"Response"`
}

func (r *QueryCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Rectf struct {
	// 矩形框左上角水平座标
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *float64 `json:"X,omitnil,omitempty" name:"X"`

	// 矩形框左上角竖直座标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 矩形框宽度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *float64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 矩形框长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *float64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type ShotInfo struct {
	// 镜头开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeStamp *float64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// 镜头结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTimeStamp *float64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`
}

type ShowInfo struct {
	// 节目日期(只在新闻有效)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 台标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 节目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Column *string `json:"Column,omitnil,omitempty" name:"Column"`

	// 来源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 节目封面
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverImageURL *string `json:"CoverImageURL,omitnil,omitempty" name:"CoverImageURL"`

	// 节目内容概要列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummarySet []*string `json:"SummarySet,omitnil,omitempty" name:"SummarySet"`

	// 节目片段标题列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TitleSet []*string `json:"TitleSet,omitnil,omitempty" name:"TitleSet"`

	// 音频识别结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioInfoSet []*AudioInfo `json:"AudioInfoSet,omitnil,omitempty" name:"AudioInfoSet"`

	// 可视文字识别结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextInfoSet []*TextInfo `json:"TextInfoSet,omitnil,omitempty" name:"TextInfoSet"`

	// 已分类人物信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassifiedPersonInfoSet []*ClassifiedPersonInfo `json:"ClassifiedPersonInfoSet,omitnil,omitempty" name:"ClassifiedPersonInfoSet"`

	// 文本标签列表，包含标签内容和出现信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTagSet *MultiLevelTag `json:"TextTagSet,omitnil,omitempty" name:"TextTagSet"`

	// 帧标签列表，包括人物信息，场景信息等
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameTagSet *MultiLevelTag `json:"FrameTagSet,omitnil,omitempty" name:"FrameTagSet"`

	// 视频下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebMediaURL *string `json:"WebMediaURL,omitnil,omitempty" name:"WebMediaURL"`

	// 媒资分类信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaClassifierSet []*string `json:"MediaClassifierSet,omitnil,omitempty" name:"MediaClassifierSet"`

	// 概要标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryTagSet []*string `json:"SummaryTagSet,omitnil,omitempty" name:"SummaryTagSet"`

	// 未知人物信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnknownPersonSet []*UnknownPerson `json:"UnknownPersonSet,omitnil,omitempty" name:"UnknownPersonSet"`

	// 树状已分类人物信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiLevelPersonInfoSet []*MultiLevelPersonInfo `json:"MultiLevelPersonInfoSet,omitnil,omitempty" name:"MultiLevelPersonInfoSet"`
}

type SortBy struct {
	// 排序字段，默认为CreateTime
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// true表示降序，false表示升序
	Descend *bool `json:"Descend,omitnil,omitempty" name:"Descend"`
}

type TTSMode struct {
	// 语速，范围：[-2，2]，分别对应不同语速：
	// -2代表0.6倍
	// -1代表0.8倍
	// 0代表1.0倍（默认）
	// 1代表1.2倍
	// 2代表1.5倍
	// 如果需要更细化的语速，可以保留小数点后 2 位，例如0.5/1.25/2.81等。
	Speed *float64 `json:"Speed,omitnil,omitempty" name:"Speed"`

	// 音色 ID，[音色体验地址](https://cloud.tencent.com/product/tts)。
	// 
	// 
	// |音乐ID|音色名称|推荐场景|
	// |--|--|--|
	// |1001|智瑜|情感女声|
	// |1002|智聆|通用女声|
	// |1003|智美|客服女声|
	// |1004|智云|通用男声|
	// |1005|智莉|通用女声|
	// |1007|智娜|客服女声|
	// |1008|智琪|客服女声|
	// |1009|智芸|知性女声|
	// |1010|智华|通用男声|
	// |1017|智蓉|情感女声|
	// |1018|智靖|情感男声|
	// 
	// 
	VoiceType *int64 `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`
}

type TaskFilter struct {
	// 媒资文件类型
	MediaTypeSet []*int64 `json:"MediaTypeSet,omitnil,omitempty" name:"MediaTypeSet"`

	// 待筛选的任务状态列表
	TaskStatusSet []*int64 `json:"TaskStatusSet,omitnil,omitempty" name:"TaskStatusSet"`

	// 待筛选的任务名称数组
	TaskNameSet []*string `json:"TaskNameSet,omitnil,omitempty" name:"TaskNameSet"`

	// TaskId数组
	TaskIdSet []*string `json:"TaskIdSet,omitnil,omitempty" name:"TaskIdSet"`

	// 媒资文件名数组
	MediaNameSet []*string `json:"MediaNameSet,omitnil,omitempty" name:"MediaNameSet"`

	// 媒资语言类型
	MediaLangSet []*int64 `json:"MediaLangSet,omitnil,omitempty" name:"MediaLangSet"`

	// 媒资素材一级类型
	MediaLabelSet []*int64 `json:"MediaLabelSet,omitnil,omitempty" name:"MediaLabelSet"`

	// 媒资自定义标签数组
	LabelSet []*string `json:"LabelSet,omitnil,omitempty" name:"LabelSet"`
}

type TaskInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 描述任务名称，指定后可根据名称筛选
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 媒资文件ID
	MediaId *string `json:"MediaId,omitnil,omitempty" name:"MediaId"`

	// 任务执行状态
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 任务进度，范围为[0，100]
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskProgress *float64 `json:"TaskProgress,omitnil,omitempty" name:"TaskProgress"`

	// 任务执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTimeCost *int64 `json:"TaskTimeCost,omitnil,omitempty" name:"TaskTimeCost"`

	// 任务创建时间
	TaskCreateTime *string `json:"TaskCreateTime,omitnil,omitempty" name:"TaskCreateTime"`

	// 任务开始执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStartTime *string `json:"TaskStartTime,omitnil,omitempty" name:"TaskStartTime"`

	// 任务失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`

	// 任务执行时指定的先验知识
	MediaPreknownInfo *MediaPreknownInfo `json:"MediaPreknownInfo,omitnil,omitempty" name:"MediaPreknownInfo"`

	// 媒资文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaName *string `json:"MediaName,omitnil,omitempty" name:"MediaName"`

	// 媒资自定义标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 任务分析完成后的后调地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallbackURL *string `json:"CallbackURL,omitnil,omitempty" name:"CallbackURL"`

	// 任务对应的媒资文件元信息，仅在MediaType为Audio时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioMetadata *AudioMetadata `json:"AudioMetadata,omitnil,omitempty" name:"AudioMetadata"`

	// 任务对应的媒资文件元信息，仅在MediaType为Audio时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageMetadata *ImageMetadata `json:"ImageMetadata,omitnil,omitempty" name:"ImageMetadata"`

	// 任务对应的媒资文件元信息，仅在MediaType为Text时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextMetadata *TextMetadata `json:"TextMetadata,omitnil,omitempty" name:"TextMetadata"`

	// 任务对应的媒资文件元信息，仅在MediaType为Video时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metadata *MediaMetadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

type TextAppearInfo struct {
	// 文本结果数组中的下标
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 关键词在文本中出现的起始偏移量(包含)
	StartPosition *int64 `json:"StartPosition,omitnil,omitempty" name:"StartPosition"`

	// 关键词在文本中出现的结束偏移量(不包含)
	EndPosition *int64 `json:"EndPosition,omitnil,omitempty" name:"EndPosition"`
}

type TextData struct {
	// 文本内容信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 文本概要信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 文本标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTagSet *MultiLevelTag `json:"TextTagSet,omitnil,omitempty" name:"TextTagSet"`
}

type TextInfo struct {
	// OCR提取的内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// OCR起始时间戳，从0开始
	StartTimeStamp *float64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// OCR结束时间戳，从0开始
	EndTimeStamp *float64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// OCR标签信息
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

type TextMetadata struct {
	// 媒资文本文件大小，单位为字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 媒资文本文件MD5
	// 注意：此字段可能返回 null，表示取不到有效值。
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`

	// 媒资文本文件字符数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 媒资文本文件格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type TextSegMatchShotScore struct {
	// 数组第 i 个值表示该文本摘要和第 i 个镜头的匹配度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScoreSet []*float64 `json:"ScoreSet,omitnil,omitempty" name:"ScoreSet"`
}

type UnknownPerson struct {
	// 视觉出现信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoAppearSet []*VideoAppearInfo `json:"VideoAppearSet,omitnil,omitempty" name:"VideoAppearSet"`

	// 未知人物是否可以入库(只有当未知人物人脸小图质量分符合要求时才可入库)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PutLibraryAllowed *bool `json:"PutLibraryAllowed,omitnil,omitempty" name:"PutLibraryAllowed"`
}

// Predefined struct for user
type UpdateCustomCategoryRequestParams struct {
	// 自定义人物类型Id
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 一级自定义人物类型
	L1Category *string `json:"L1Category,omitnil,omitempty" name:"L1Category"`

	// 二级自定义人物类型
	L2Category *string `json:"L2Category,omitnil,omitempty" name:"L2Category"`
}

type UpdateCustomCategoryRequest struct {
	*tchttp.BaseRequest
	
	// 自定义人物类型Id
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 一级自定义人物类型
	L1Category *string `json:"L1Category,omitnil,omitempty" name:"L1Category"`

	// 二级自定义人物类型
	L2Category *string `json:"L2Category,omitnil,omitempty" name:"L2Category"`
}

func (r *UpdateCustomCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCustomCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CategoryId")
	delete(f, "L1Category")
	delete(f, "L2Category")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCustomCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCustomCategoryResponseParams struct {
	// 成功更新的自定义人物类型Id
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCustomCategoryResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCustomCategoryResponseParams `json:"Response"`
}

func (r *UpdateCustomCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCustomCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCustomPersonRequestParams struct {
	// 待更新的自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 更新后的自定义人物名称，如为空则不更新
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 更新后的自定义人物简介，如为空则不更新
	BasicInfo *string `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// 更新后的分类信息，如为空则不更新
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`
}

type UpdateCustomPersonRequest struct {
	*tchttp.BaseRequest
	
	// 待更新的自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 更新后的自定义人物名称，如为空则不更新
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 更新后的自定义人物简介，如为空则不更新
	BasicInfo *string `json:"BasicInfo,omitnil,omitempty" name:"BasicInfo"`

	// 更新后的分类信息，如为空则不更新
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`
}

func (r *UpdateCustomPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCustomPersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Name")
	delete(f, "BasicInfo")
	delete(f, "CategoryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCustomPersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCustomPersonResponseParams struct {
	// 成功更新的自定义人物Id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCustomPersonResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCustomPersonResponseParams `json:"Response"`
}

func (r *UpdateCustomPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCustomPersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VideoAppearInfo struct {
	// 视觉信息起始时间戳，从0开始
	StartTimeStamp *float64 `json:"StartTimeStamp,omitnil,omitempty" name:"StartTimeStamp"`

	// 视觉信息终止时间戳，从0开始
	// 关键词在视觉信息中的区间为[StartTimeStamp, EndTimeStamp)
	EndTimeStamp *float64 `json:"EndTimeStamp,omitnil,omitempty" name:"EndTimeStamp"`

	// 关键词在视觉信息中的封面图片
	ImageURL *string `json:"ImageURL,omitnil,omitempty" name:"ImageURL"`
}

type VideoRotationMode struct {
	// 生成的视频是否需要横屏转竖屏。
	ActiveVideoRotation *bool `json:"ActiveVideoRotation,omitnil,omitempty" name:"ActiveVideoRotation"`
}