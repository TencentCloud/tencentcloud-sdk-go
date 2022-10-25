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

package v20200304

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ArtifactReduction struct {
	// 去毛刺方式：weak,,strong
	Type *string `json:"Type,omitempty" name:"Type"`

	// 去毛刺算法，可选项：
	// edaf,
	// wdaf，
	// 默认edaf。
	// 注意：此参数已经弃用
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
}

type AudioEnhance struct {
	// 音效增强种类，可选项：normal
	Type *string `json:"Type,omitempty" name:"Type"`
}

type AudioInfo struct {
	// 音频码率，取值范围：0 和 [26, 256]，单位：kbps。
	// 注意：当取值为 0，表示音频码率和原始音频保持一致。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频编码器，可选项：aac,mp3,ac3,flac,mp2。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 声道数，可选项：
	// 1：单声道，
	// 2：双声道，
	// 6：立体声。
	Channel *int64 `json:"Channel,omitempty" name:"Channel"`

	// 采样率，单位：Hz。可选项：32000，44100,48000
	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 音频降噪信息
	Denoise *Denoise `json:"Denoise,omitempty" name:"Denoise"`

	// 开启添加静音，可选项：
	// 0：不开启，
	// 1：开启，
	// 默认不开启
	EnableMuteAudio *int64 `json:"EnableMuteAudio,omitempty" name:"EnableMuteAudio"`

	// 音频响度信息
	LoudnessInfo *LoudnessInfo `json:"LoudnessInfo,omitempty" name:"LoudnessInfo"`

	// 音频音效增强
	AudioEnhance *AudioEnhance `json:"AudioEnhance,omitempty" name:"AudioEnhance"`

	// 去除混音
	RemoveReverb *RemoveReverb `json:"RemoveReverb,omitempty" name:"RemoveReverb"`
}

type AudioInfoResultItem struct {
	// 音频流的流id。
	Stream *int64 `json:"Stream,omitempty" name:"Stream"`

	// 音频采样率 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sample *int64 `json:"Sample,omitempty" name:"Sample"`

	// 音频声道数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Channel *int64 `json:"Channel,omitempty" name:"Channel"`

	// 编码格式，如aac, mp3等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 码率，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频时长，单位：ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
}

type CallbackInfo struct {
	// 回调URL。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type ClassificationEditingInfo struct {
	// 是否开启视频分类识别。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
}

type ClassificationTaskResult struct {
	// 编辑任务状态。 
	// 1：执行中；2：成功；3：失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 编辑任务失败错误码。 
	// 0：成功；其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 编辑任务失败错误描述。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 视频分类识别结果集。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSet []*ClassificationTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet"`
}

type ClassificationTaskResultItem struct {
	// 分类名称。
	Classification *string `json:"Classification,omitempty" name:"Classification"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type ColorEnhance struct {
	// 颜色增强类型，可选项：
	// 1.  tra；
	// 2.  weak；
	// 3.  normal;
	// 4.  strong;
	// 注意：tra不支持自适应调整，处理速度更快；weak,normal,strong支持基于画面颜色自适应，处理速度更慢。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type CosAuthMode struct {
	// 授权类型，可选值： 
	// 0：bucket授权，需要将对应bucket授权给本服务帐号（3020447271和100012301793），否则会读写cos失败； 
	// 1：key托管，把cos的账号id和key托管于本服务，本服务会提供一个托管id； 
	// 3：临时key授权。
	// 注意：目前智能编辑还不支持临时key授权；画质重生目前只支持bucket授权
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// cos账号托管id，Type等于1时必选。
	HostedId *string `json:"HostedId,omitempty" name:"HostedId"`

	// cos身份识别id，Type等于3时必选。
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// cos身份秘钥，Type等于3时必选。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// 临时授权 token，Type等于3时必选。
	Token *string `json:"Token,omitempty" name:"Token"`
}

type CosInfo struct {
	// cos 区域值。例如：ap-beijing。
	Region *string `json:"Region,omitempty" name:"Region"`

	// cos 存储桶，格式为BuketName-AppId。例如：test-123456。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// cos 路径。 
	// 对于写表示目录，例如：/test； 
	// 对于读表示文件路径，例如：/test/test.mp4。
	Path *string `json:"Path,omitempty" name:"Path"`

	// cos 授权信息，不填默认为公有权限。
	CosAuthMode *CosAuthMode `json:"CosAuthMode,omitempty" name:"CosAuthMode"`
}

type CoverEditingInfo struct {
	// 是否开启智能封面。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
}

type CoverTaskResult struct {
	// 编辑任务状态。 
	// 1：执行中；2：成功；3：失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 编辑任务失败错误码。 
	// 0：成功；其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 编辑任务失败错误描述。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 智能封面结果集。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSet []*CoverTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet"`
}

type CoverTaskResultItem struct {
	// 智能封面地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

// Predefined struct for user
type CreateEditingTaskRequestParams struct {
	// 智能编辑任务参数。
	EditingInfo *EditingInfo `json:"EditingInfo,omitempty" name:"EditingInfo"`

	// 视频源信息。
	DownInfo *DownInfo `json:"DownInfo,omitempty" name:"DownInfo"`

	// 结果存储信息。对于包含智能拆条、智能集锦或者智能封面的任务必选。
	SaveInfo *SaveInfo `json:"SaveInfo,omitempty" name:"SaveInfo"`

	// 任务结果回调地址信息。
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitempty" name:"CallbackInfo"`
}

type CreateEditingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 智能编辑任务参数。
	EditingInfo *EditingInfo `json:"EditingInfo,omitempty" name:"EditingInfo"`

	// 视频源信息。
	DownInfo *DownInfo `json:"DownInfo,omitempty" name:"DownInfo"`

	// 结果存储信息。对于包含智能拆条、智能集锦或者智能封面的任务必选。
	SaveInfo *SaveInfo `json:"SaveInfo,omitempty" name:"SaveInfo"`

	// 任务结果回调地址信息。
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitempty" name:"CallbackInfo"`
}

func (r *CreateEditingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEditingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EditingInfo")
	delete(f, "DownInfo")
	delete(f, "SaveInfo")
	delete(f, "CallbackInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEditingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEditingTaskResponseParams struct {
	// 编辑任务 ID，可以通过该 ID 查询任务状态。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEditingTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateEditingTaskResponseParams `json:"Response"`
}

func (r *CreateEditingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEditingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMediaProcessTaskRequestParams struct {
	// 编辑处理任务参数。
	MediaProcessInfo *MediaProcessInfo `json:"MediaProcessInfo,omitempty" name:"MediaProcessInfo"`

	// 编辑处理任务输入源列表。
	SourceInfoSet []*MediaSourceInfo `json:"SourceInfoSet,omitempty" name:"SourceInfoSet"`

	// 结果存储信息，对于涉及存储的请求必选。部子任务支持数组备份写，具体以对应任务文档为准。
	SaveInfoSet []*SaveInfo `json:"SaveInfoSet,omitempty" name:"SaveInfoSet"`

	// 任务结果回调地址信息。部子任务支持数组备份回调，具体以对应任务文档为准。
	CallbackInfoSet []*CallbackInfo `json:"CallbackInfoSet,omitempty" name:"CallbackInfoSet"`
}

type CreateMediaProcessTaskRequest struct {
	*tchttp.BaseRequest
	
	// 编辑处理任务参数。
	MediaProcessInfo *MediaProcessInfo `json:"MediaProcessInfo,omitempty" name:"MediaProcessInfo"`

	// 编辑处理任务输入源列表。
	SourceInfoSet []*MediaSourceInfo `json:"SourceInfoSet,omitempty" name:"SourceInfoSet"`

	// 结果存储信息，对于涉及存储的请求必选。部子任务支持数组备份写，具体以对应任务文档为准。
	SaveInfoSet []*SaveInfo `json:"SaveInfoSet,omitempty" name:"SaveInfoSet"`

	// 任务结果回调地址信息。部子任务支持数组备份回调，具体以对应任务文档为准。
	CallbackInfoSet []*CallbackInfo `json:"CallbackInfoSet,omitempty" name:"CallbackInfoSet"`
}

func (r *CreateMediaProcessTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMediaProcessTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MediaProcessInfo")
	delete(f, "SourceInfoSet")
	delete(f, "SaveInfoSet")
	delete(f, "CallbackInfoSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMediaProcessTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMediaProcessTaskResponseParams struct {
	// 编辑任务 ID，可以通过该 ID 查询任务状态和结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMediaProcessTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateMediaProcessTaskResponseParams `json:"Response"`
}

func (r *CreateMediaProcessTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMediaProcessTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMediaQualityRestorationTaskRequestParams struct {
	// 源文件地址。
	DownInfo *DownInfo `json:"DownInfo,omitempty" name:"DownInfo"`

	// 画质重生任务参数信息。
	TransInfo []*SubTaskTranscodeInfo `json:"TransInfo,omitempty" name:"TransInfo"`

	// 任务结束后文件存储信息。
	SaveInfo *SaveInfo `json:"SaveInfo,omitempty" name:"SaveInfo"`

	// 任务结果回调地址信息。
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitempty" name:"CallbackInfo"`

	// 极速高清体验馆渠道标志。
	TopSpeedCodecChannel *uint64 `json:"TopSpeedCodecChannel,omitempty" name:"TopSpeedCodecChannel"`
}

type CreateMediaQualityRestorationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 源文件地址。
	DownInfo *DownInfo `json:"DownInfo,omitempty" name:"DownInfo"`

	// 画质重生任务参数信息。
	TransInfo []*SubTaskTranscodeInfo `json:"TransInfo,omitempty" name:"TransInfo"`

	// 任务结束后文件存储信息。
	SaveInfo *SaveInfo `json:"SaveInfo,omitempty" name:"SaveInfo"`

	// 任务结果回调地址信息。
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitempty" name:"CallbackInfo"`

	// 极速高清体验馆渠道标志。
	TopSpeedCodecChannel *uint64 `json:"TopSpeedCodecChannel,omitempty" name:"TopSpeedCodecChannel"`
}

func (r *CreateMediaQualityRestorationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMediaQualityRestorationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DownInfo")
	delete(f, "TransInfo")
	delete(f, "SaveInfo")
	delete(f, "CallbackInfo")
	delete(f, "TopSpeedCodecChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMediaQualityRestorationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMediaQualityRestorationTaskResponseParams struct {
	// 画质重生任务ID，可以通过该ID查询任务状态。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMediaQualityRestorationTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateMediaQualityRestorationTaskResponseParams `json:"Response"`
}

func (r *CreateMediaQualityRestorationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMediaQualityRestorationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQualityControlTaskRequestParams struct {
	// 质检任务参数
	QualityControlInfo *QualityControlInfo `json:"QualityControlInfo,omitempty" name:"QualityControlInfo"`

	// 视频源信息
	DownInfo *DownInfo `json:"DownInfo,omitempty" name:"DownInfo"`

	// 任务结果回调地址信息
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitempty" name:"CallbackInfo"`
}

type CreateQualityControlTaskRequest struct {
	*tchttp.BaseRequest
	
	// 质检任务参数
	QualityControlInfo *QualityControlInfo `json:"QualityControlInfo,omitempty" name:"QualityControlInfo"`

	// 视频源信息
	DownInfo *DownInfo `json:"DownInfo,omitempty" name:"DownInfo"`

	// 任务结果回调地址信息
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitempty" name:"CallbackInfo"`
}

func (r *CreateQualityControlTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQualityControlTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QualityControlInfo")
	delete(f, "DownInfo")
	delete(f, "CallbackInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQualityControlTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQualityControlTaskResponseParams struct {
	// 质检任务 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateQualityControlTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateQualityControlTaskResponseParams `json:"Response"`
}

func (r *CreateQualityControlTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQualityControlTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DarInfo struct {
	// 填充模式，可选值：
	// 1：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。
	// 2：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“。
	// 默认为2。
	FillMode *uint64 `json:"FillMode,omitempty" name:"FillMode"`
}

type Denoise struct {
	// 音频降噪强度，可选项：
	// 1. weak
	// 2.normal，
	// 3.strong
	// 默认为weak
	Type *string `json:"Type,omitempty" name:"Type"`
}

type Denoising struct {
	// 去噪方式，可选项：
	// templ：时域降噪；
	// spatial：空域降噪,
	// fast-spatial：快速空域降噪。
	// 注意：可选择组合方式：
	// 1.type:"templ,spatial" ;
	// 2.type:"templ,fast-spatial"。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时域去噪强度，可选值：0.0-1.0 。小于0.0的默认为0.0，大于1.0的默认为1.0。
	TemplStrength *float64 `json:"TemplStrength,omitempty" name:"TemplStrength"`

	// 空域去噪强度，可选值：0.0-1.0 。小于0.0的默认为0.0，大于1.0的默认为1.0。
	SpatialStrength *float64 `json:"SpatialStrength,omitempty" name:"SpatialStrength"`
}

// Predefined struct for user
type DescribeEditingTaskResultRequestParams struct {
	// 编辑任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeEditingTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 编辑任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeEditingTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEditingTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEditingTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEditingTaskResultResponseParams struct {
	// 编辑任务结果信息。
	TaskResult *EditingTaskResult `json:"TaskResult,omitempty" name:"TaskResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEditingTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEditingTaskResultResponseParams `json:"Response"`
}

func (r *DescribeEditingTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEditingTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaProcessTaskResultRequestParams struct {
	// 编辑处理任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeMediaProcessTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 编辑处理任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeMediaProcessTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaProcessTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaProcessTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaProcessTaskResultResponseParams struct {
	// 任务处理结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskResult *MediaProcessTaskResult `json:"TaskResult,omitempty" name:"TaskResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMediaProcessTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediaProcessTaskResultResponseParams `json:"Response"`
}

func (r *DescribeMediaProcessTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaProcessTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaQualityRestorationTaskRusultRequestParams struct {
	// 画质重生任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeMediaQualityRestorationTaskRusultRequest struct {
	*tchttp.BaseRequest
	
	// 画质重生任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeMediaQualityRestorationTaskRusultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaQualityRestorationTaskRusultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaQualityRestorationTaskRusultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMediaQualityRestorationTaskRusultResponseParams struct {
	// 画质重生任务结果信息
	TaskResult *MediaQualityRestorationTaskResult `json:"TaskResult,omitempty" name:"TaskResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMediaQualityRestorationTaskRusultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMediaQualityRestorationTaskRusultResponseParams `json:"Response"`
}

func (r *DescribeMediaQualityRestorationTaskRusultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaQualityRestorationTaskRusultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityControlTaskResultRequestParams struct {
	// 质检任务 ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeQualityControlTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 质检任务 ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeQualityControlTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityControlTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQualityControlTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityControlTaskResultResponseParams struct {
	// 质检任务结果信息
	TaskResult *QualityControlInfoTaskResult `json:"TaskResult,omitempty" name:"TaskResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeQualityControlTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQualityControlTaskResultResponseParams `json:"Response"`
}

func (r *DescribeQualityControlTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityControlTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownInfo struct {
	// 下载类型，可选值： 
	// 0：UrlInfo； 
	// 1：CosInfo。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Url形式下载信息，当Type等于0时必选。
	UrlInfo *UrlInfo `json:"UrlInfo,omitempty" name:"UrlInfo"`

	// Cos形式下载信息，当Type等于1时必选。
	CosInfo *CosInfo `json:"CosInfo,omitempty" name:"CosInfo"`
}

type DynamicImageInfo struct {
	// 画面质量，范围：1~100。
	// <li>对于webp格式，默认：75</li>
	// <li>对于gif格式，小于10为低质量，大于50为高质量，其它为普通。默认：低质量。</li>
	Quality *uint64 `json:"Quality,omitempty" name:"Quality"`
}

type EditInfo struct {
	// 剪辑开始时间，单位：ms。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 剪辑结束时间，单位：ms
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type EditingInfo struct {
	// 视频标签识别任务参数，不填则不开启。
	TagEditingInfo *TagEditingInfo `json:"TagEditingInfo,omitempty" name:"TagEditingInfo"`

	// 视频分类识别任务参数，不填则不开启。
	ClassificationEditingInfo *ClassificationEditingInfo `json:"ClassificationEditingInfo,omitempty" name:"ClassificationEditingInfo"`

	// 智能拆条任务参数，不填则不开启。
	StripEditingInfo *StripEditingInfo `json:"StripEditingInfo,omitempty" name:"StripEditingInfo"`

	// 智能集锦任务参数，不填则不开启。
	HighlightsEditingInfo *HighlightsEditingInfo `json:"HighlightsEditingInfo,omitempty" name:"HighlightsEditingInfo"`

	// 智能封面任务参数，不填则不开启。
	CoverEditingInfo *CoverEditingInfo `json:"CoverEditingInfo,omitempty" name:"CoverEditingInfo"`

	// 片头片尾识别任务参数，不填则不开启。
	OpeningEndingEditingInfo *OpeningEndingEditingInfo `json:"OpeningEndingEditingInfo,omitempty" name:"OpeningEndingEditingInfo"`
}

type EditingTaskResult struct {
	// 编辑任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 编辑任务状态。 
	// 1：执行中；2：已完成。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 视频标签识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagTaskResult *TagTaskResult `json:"TagTaskResult,omitempty" name:"TagTaskResult"`

	// 视频分类识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationTaskResult *ClassificationTaskResult `json:"ClassificationTaskResult,omitempty" name:"ClassificationTaskResult"`

	// 智能拆条结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StripTaskResult *StripTaskResult `json:"StripTaskResult,omitempty" name:"StripTaskResult"`

	// 智能集锦结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighlightsTaskResult *HighlightsTaskResult `json:"HighlightsTaskResult,omitempty" name:"HighlightsTaskResult"`

	// 智能封面结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverTaskResult *CoverTaskResult `json:"CoverTaskResult,omitempty" name:"CoverTaskResult"`

	// 片头片尾识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpeningEndingTaskResult *OpeningEndingTaskResult `json:"OpeningEndingTaskResult,omitempty" name:"OpeningEndingTaskResult"`
}

type FaceProtect struct {
	// 人脸区域增强强度，可选项：0.0-1.0。小于0.0的默认为0.0，大于1.0的默认为1.0。
	FaceUsmRatio *float64 `json:"FaceUsmRatio,omitempty" name:"FaceUsmRatio"`
}

type FileInfo struct {
	// 任务结束后生成的文件大小。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 任务结束后生成的文件格式，例如：mp4,flv等等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 任务结束后生成的文件整体码率，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 任务结束后生成的文件时长，单位：ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 任务结束后生成的文件视频信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoInfoResult []*VideoInfoResultItem `json:"VideoInfoResult,omitempty" name:"VideoInfoResult"`

	// 任务结束后生成的文件音频信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioInfoResult []*AudioInfoResultItem `json:"AudioInfoResult,omitempty" name:"AudioInfoResult"`
}

type FrameTagItem struct {
	// 标签起始时间戳PTS(ms)
	StartPts *uint64 `json:"StartPts,omitempty" name:"StartPts"`

	// 语句结束时间戳PTS(ms)
	EndPts *uint64 `json:"EndPts,omitempty" name:"EndPts"`

	// 字符串形式的起始结束时间
	Period *string `json:"Period,omitempty" name:"Period"`

	// 标签数组
	TagItems []*TagItem `json:"TagItems,omitempty" name:"TagItems"`
}

type FrameTagRec struct {
	// 标签类型：
	// "Common": 通用类型
	// "Game":游戏类型
	TagType *string `json:"TagType,omitempty" name:"TagType"`

	// 游戏具体类型:
	// "HonorOfKings_AnchorViews":王者荣耀主播视角
	// "HonorOfKings_GameViews":王者荣耀比赛视角
	// "LOL_AnchorViews":英雄联盟主播视角
	// "LOL_GameViews":英雄联盟比赛视角
	// "PUBG_AnchorViews":和平精英主播视角
	// "PUBG_GameViews":和平精英比赛视角
	GameExtendType *string `json:"GameExtendType,omitempty" name:"GameExtendType"`
}

type FrameTagResult struct {
	// 帧标签结果数组
	FrameTagItems []*FrameTagItem `json:"FrameTagItems,omitempty" name:"FrameTagItems"`
}

type HiddenMarkInfo struct {
	// 数字水印路径,，如果不从Cos拉取水印，则必填
	Path *string `json:"Path,omitempty" name:"Path"`

	// 数字水印频率，可选值：[1,256]，默认值为30
	Frequency *int64 `json:"Frequency,omitempty" name:"Frequency"`

	// 数字水印强度，可选值：[32,128]，默认值为64
	Strength *int64 `json:"Strength,omitempty" name:"Strength"`

	// 数字水印的Cos 信息，从Cos上拉取图片水印时必填。
	CosInfo *CosInfo `json:"CosInfo,omitempty" name:"CosInfo"`
}

type HighlightsEditingInfo struct {
	// 是否开启智能集锦。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
}

type HighlightsTaskResult struct {
	// 编辑任务状态。 
	// 1：执行中；2：成功；3：失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 编辑任务失败错误码。 
	// 0：成功；其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 编辑任务失败错误描述。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 智能集锦结果集。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSet []*HighlightsTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet"`
}

type HighlightsTaskResultItem struct {
	// 智能集锦地址。
	HighlightUrl *string `json:"HighlightUrl,omitempty" name:"HighlightUrl"`

	// 智能集锦封面地址。
	CovImgUrl *string `json:"CovImgUrl,omitempty" name:"CovImgUrl"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 智能集锦持续时间，单位：秒。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 智能集锦子片段结果集，集锦片段由这些子片段拼接生成。
	SegmentSet []*HighlightsTaskResultItemSegment `json:"SegmentSet,omitempty" name:"SegmentSet"`
}

type HighlightsTaskResultItemSegment struct {
	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 集锦片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 集锦片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type IntervalTime struct {
	// 间隔周期，单位ms
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// 开始时间点，单位ms
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

type LoudnessInfo struct {
	// 音频整体响度
	Loudness *float64 `json:"Loudness,omitempty" name:"Loudness"`

	// 音频响度范围
	LoudnessRange *float64 `json:"LoudnessRange,omitempty" name:"LoudnessRange"`
}

type LowLightEnhance struct {
	// 低光照增强类型，可选项：normal。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type MediaCuttingInfo struct {
	// 截取时间信息。
	TimeInfo *MediaCuttingTimeInfo `json:"TimeInfo,omitempty" name:"TimeInfo"`

	// 输出结果信息。
	TargetInfo *MediaTargetInfo `json:"TargetInfo,omitempty" name:"TargetInfo"`

	// 截取结果形式信息。
	OutForm *MediaCuttingOutForm `json:"OutForm,omitempty" name:"OutForm"`

	// 列表文件形式，存储到用户存储服务中，可选值：
	// <li>NoListFile：不存储结果列表; </li>
	// <li>UseSaveInfo：默认，结果列表和结果存储同一位置（即SaveInfoSet 的第一个存储位置）；</li>
	// <li>SaveInfoSet 存储的Id：存储在指定的存储位置。</li>
	ResultListSaveType *string `json:"ResultListSaveType,omitempty" name:"ResultListSaveType"`

	// 水印信息，最多支持 10 个水印。
	WatermarkInfoSet []*MediaCuttingWatermark `json:"WatermarkInfoSet,omitempty" name:"WatermarkInfoSet"`

	// 是否去除纯色截图，如果值为 True ，对应时间点的截图如果是纯色，将略过。
	DropPureColor *string `json:"DropPureColor,omitempty" name:"DropPureColor"`
}

type MediaCuttingOutForm struct {
	// 输出类型，可选值：
	// Static：静态图；
	// Dynamic：动态图；
	// Sprite：雪碧图；
	// Video：视频。
	// 
	// 注1：不同类型时，对应的 TargetInfo.Format 格式支持如下：
	// Static：jpg、png；
	// Dynamic：gif；
	// Sprite：jpg、png；
	// Video：mp4。
	// 
	// 注2：当 Type=Sprite时，TargetInfo指定的尺寸表示小图的大小，最终结果尺寸以输出为准。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 背景填充方式，可选值：
	// White：白色填充；
	// Black：黑色填充；
	// Stretch：拉伸；
	// Gaussian：高斯模糊；
	// 默认White。
	FillType *string `json:"FillType,omitempty" name:"FillType"`

	// 【废弃】参考SpriteInfo
	SpriteRowCount *int64 `json:"SpriteRowCount,omitempty" name:"SpriteRowCount"`

	// 【废弃】参考SpriteInfo
	SpriteColumnCount *int64 `json:"SpriteColumnCount,omitempty" name:"SpriteColumnCount"`

	// Type=Sprite时有效，表示雪碧图参数信息。
	SpriteInfo *SpriteImageInfo `json:"SpriteInfo,omitempty" name:"SpriteInfo"`

	// Type=Dynamic时有效，表示动图参数信息。
	DynamicInfo *DynamicImageInfo `json:"DynamicInfo,omitempty" name:"DynamicInfo"`
}

type MediaCuttingTaskResult struct {
	// 如果ResultListType不为NoListFile时，结果（TaskResultFile）列表文件的存储位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListFile *TaskResultFile `json:"ListFile,omitempty" name:"ListFile"`

	// 结果个数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultCount *int64 `json:"ResultCount,omitempty" name:"ResultCount"`

	// 第一个结果文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstFile *TaskResultFile `json:"FirstFile,omitempty" name:"FirstFile"`

	// 最后一个结果文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastFile *TaskResultFile `json:"LastFile,omitempty" name:"LastFile"`

	// 任务结果包含的图片总数。
	// 静态图：总数即为文件数；
	// 雪碧图：所有小图总数；
	// 动图、视频：不计算图片数，为 0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageCount *int64 `json:"ImageCount,omitempty" name:"ImageCount"`
}

type MediaCuttingTimeInfo struct {
	// 时间类型，可选值：
	// PointSet：时间点集合；
	// IntervalPoint：周期采样点；
	// SectionSet：时间片段集合。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 截取时间点集合，单位毫秒，Type=PointSet时必选。
	PointSet []*int64 `json:"PointSet,omitempty" name:"PointSet"`

	// 周期采样点信息，Type=IntervalPoint时必选。
	IntervalPoint *IntervalTime `json:"IntervalPoint,omitempty" name:"IntervalPoint"`

	// 时间区间集合信息，Type=SectionSet时必选。
	SectionSet []*SectionTime `json:"SectionSet,omitempty" name:"SectionSet"`
}

type MediaCuttingWatermark struct {
	// 水印类型，可选值：
	// <li>Image：图像水印；</li>
	// <li>Text：文字水印。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 图像水印信息，当 Type=Image 时必选。
	Image *MediaCuttingWatermarkImage `json:"Image,omitempty" name:"Image"`

	// 文字水印信息，当 Type=Text 时必选。
	Text *MediaCuttingWatermarkText `json:"Text,omitempty" name:"Text"`
}

type MediaCuttingWatermarkImage struct {
	// 水印源的ID，对应SourceInfoSet内的源。
	// 注意1：对应的 MediaSourceInfo.Type需要为Image。
	// 注意2：对于动图，只取第一帧图像作为水印源。
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// 水印水平坐标，单位像素，默认：0。
	PosX *uint64 `json:"PosX,omitempty" name:"PosX"`

	// 水印垂直坐标，单位像素，默认：0。
	PosY *uint64 `json:"PosY,omitempty" name:"PosY"`

	// 水印宽度，单位像素，默认：0。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 水印高度，单位像素，默认：0。
	// 注意：对于宽高符合以下规则：
	// 1、Width>0 且 Height>0，按指定宽高拉伸；
	// 2、Width=0 且 Height>0，以Height为基准等比缩放；
	// 3、Width>0 且 Height=0，以Width为基准等比缩放；
	// 4、Width=0 且 Height=0，采用源的宽高。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 指定坐标原点，可选值：
	// <li>LeftTop：PosXY 表示水印左上点到图片左上点的相对位置</li>
	// <li>RightTop：PosXY 表示水印右上点到图片右上点的相对位置</li>
	// <li>LeftBottom：PosXY 表示水印左下点到图片左下点的相对位置</li>
	// <li>RightBottom：PosXY 表示水印右下点到图片右下点的相对位置</li>
	// <li>Center：PosXY 表示水印中心点到图片中心点的相对位置</li>
	// 默认：LeftTop。
	PosOriginType *string `json:"PosOriginType,omitempty" name:"PosOriginType"`
}

type MediaCuttingWatermarkText struct {
	// 水印文字。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 文字大小
	FontSize *uint64 `json:"FontSize,omitempty" name:"FontSize"`

	// 水印水平坐标，单位像素，默认：0。
	PosX *uint64 `json:"PosX,omitempty" name:"PosX"`

	// 水印垂直坐标，单位像素，默认：0。
	PosY *uint64 `json:"PosY,omitempty" name:"PosY"`

	// 文字颜色，格式为：#RRGGBBAA，默认值：#000000。
	FontColor *string `json:"FontColor,omitempty" name:"FontColor"`

	// 文字透明度，范围：0~100，默认值：100。
	FontAlpha *uint64 `json:"FontAlpha,omitempty" name:"FontAlpha"`

	// 指定坐标原点，可选值：
	// <li>LeftTop：PosXY 表示水印左上点到图片左上点的相对位置</li>
	// <li>RightTop：PosXY 表示水印右上点到图片右上点的相对位置</li>
	// <li>LeftBottom：PosXY 表示水印左下点到图片左下点的相对位置</li>
	// <li>RightBottom：PosXY 表示水印右下点到图片右下点的相对位置</li>
	// <li>Center：PosXY 表示水印中心点到图片中心点的相对位置</li>
	// 默认：LeftTop。
	PosOriginType *string `json:"PosOriginType,omitempty" name:"PosOriginType"`

	// 字体，可选值：
	// <li>SimHei</li>
	// <li>SimKai</li>
	// <li>Arial</li>
	// 默认 SimHei。
	Font *string `json:"Font,omitempty" name:"Font"`
}

type MediaJoiningInfo struct {
	// 输出目标信息，拼接只采用FileName和Format，用于指定目标文件名和格式。
	// 其中Format只支持mp4.
	TargetInfo *MediaTargetInfo `json:"TargetInfo,omitempty" name:"TargetInfo"`

	// 拼接模式：
	// Fast：快速；
	// Normal：正常；
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type MediaJoiningTaskResult struct {
	// 拼接结果文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	File *TaskResultFile `json:"File,omitempty" name:"File"`
}

type MediaProcessInfo struct {
	// 编辑处理任务类型，可选值：
	// MediaEditing：媒体编辑（待上线）；
	// MediaCutting：媒体剪切；
	// MediaJoining：媒体拼接。
	// MediaRecognition: 媒体识别。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 视频剪切任务参数，Type=MediaCutting时必选。
	MediaCuttingInfo *MediaCuttingInfo `json:"MediaCuttingInfo,omitempty" name:"MediaCuttingInfo"`

	// 视频拼接任务参数，Type=MediaJoining时必选。
	MediaJoiningInfo *MediaJoiningInfo `json:"MediaJoiningInfo,omitempty" name:"MediaJoiningInfo"`

	// 媒体识别任务参数，Type=MediaRecognition时必选
	MediaRecognitionInfo *MediaRecognitionInfo `json:"MediaRecognitionInfo,omitempty" name:"MediaRecognitionInfo"`
}

type MediaProcessTaskResult struct {
	// 编辑处理任务ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 编辑处理任务类型，取值：
	// MediaEditing：视频编辑（待上线）；
	// MediaCutting：视频剪切；
	// MediaJoining：视频拼接。
	// MediaRecognition：媒体识别；
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 处理进度，范围：[0,100]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// 任务状态：
	// 1100：等待中；
	// 1200：执行中；
	// 2000：成功；
	// 5000：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 任务错误码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 任务错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 剪切任务处理结果，当Type=MediaCutting时才有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaCuttingTaskResult *MediaCuttingTaskResult `json:"MediaCuttingTaskResult,omitempty" name:"MediaCuttingTaskResult"`

	// 拼接任务处理结果，当Type=MediaJoining时才有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaJoiningTaskResult *MediaJoiningTaskResult `json:"MediaJoiningTaskResult,omitempty" name:"MediaJoiningTaskResult"`

	// 媒体识别任务处理结果，当Type=MediaRecognition时才有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaRecognitionTaskResult *MediaRecognitionTaskResult `json:"MediaRecognitionTaskResult,omitempty" name:"MediaRecognitionTaskResult"`
}

type MediaQualityRestorationTaskResult struct {
	// 画质重生任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 画质重生处理后文件的详细信息。
	SubTaskResult []*SubTaskResultItem `json:"SubTaskResult,omitempty" name:"SubTaskResult"`
}

type MediaRecognitionInfo struct {
	// 帧标签识别
	FrameTagRec *FrameTagRec `json:"FrameTagRec,omitempty" name:"FrameTagRec"`

	// 语音字幕识别
	SubtitleRec *SubtitleRec `json:"SubtitleRec,omitempty" name:"SubtitleRec"`
}

type MediaRecognitionTaskResult struct {
	// 帧标签识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameTagResults *FrameTagResult `json:"FrameTagResults,omitempty" name:"FrameTagResults"`

	// 语音字幕识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubtitleResults *SubtitleResult `json:"SubtitleResults,omitempty" name:"SubtitleResults"`
}

type MediaResultInfo struct {
	// 媒体时长，单位：毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 视频流信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultVideoInfoSet []*ResultVideoInfo `json:"ResultVideoInfoSet,omitempty" name:"ResultVideoInfoSet"`

	// 音频流信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultAudioInfoSet []*ResultAudioInfo `json:"ResultAudioInfoSet,omitempty" name:"ResultAudioInfoSet"`
}

type MediaSourceInfo struct {
	// 媒体源资源下载信息。
	DownInfo *DownInfo `json:"DownInfo,omitempty" name:"DownInfo"`

	// 媒体源ID标记，用于多个输入源时，请内媒体源的定位，对于多输入的任务，一般要求必选。
	// ID只能包含字母、数字、下划线、中划线，长读不能超过128。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 媒体源类型，具体类型如下：
	// Video：视频
	// Image：图片
	// Audio：音频
	Type *string `json:"Type,omitempty" name:"Type"`
}

type MediaTargetInfo struct {
	// 目标文件名，不能带特殊字符（如/等），无需后缀名，最长200字符。
	// 
	// 注1：部分子服务支持占位符，形式为： {parameter}
	// 预设parameter有：
	// index：序号；
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 媒体封装格式，最长5字符，具体格式支持根据子任务确定。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 视频流信息。
	TargetVideoInfo *TargetVideoInfo `json:"TargetVideoInfo,omitempty" name:"TargetVideoInfo"`

	// 【不再使用】
	ResultListSaveType *string `json:"ResultListSaveType,omitempty" name:"ResultListSaveType"`
}

type MuxInfo struct {
	// 删除流，可选项：video,audio。
	DeleteStream *string `json:"DeleteStream,omitempty" name:"DeleteStream"`

	// Flv 参数，目前支持add_keyframe_index
	FlvFlags *string `json:"FlvFlags,omitempty" name:"FlvFlags"`
}

type OpeningEndingEditingInfo struct {
	// 是否开启片头片尾识别。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
}

type OpeningEndingTaskResult struct {
	// 编辑任务状态。 
	// 1：执行中；2：成功；3：失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 编辑任务失败错误码。 
	// 0：成功；其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 编辑任务失败错误描述。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 片头片尾识别结果项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item *OpeningEndingTaskResultItem `json:"Item,omitempty" name:"Item"`
}

type OpeningEndingTaskResultItem struct {
	// 视频片头的结束时间点，单位：秒。
	OpeningTimeOffset *float64 `json:"OpeningTimeOffset,omitempty" name:"OpeningTimeOffset"`

	// 片头识别置信度，取值范围是 0 到 100。
	OpeningConfidence *float64 `json:"OpeningConfidence,omitempty" name:"OpeningConfidence"`

	// 视频片尾的开始时间点，单位：秒。
	EndingTimeOffset *float64 `json:"EndingTimeOffset,omitempty" name:"EndingTimeOffset"`

	// 片尾识别置信度，取值范围是 0 到 100。
	EndingConfidence *float64 `json:"EndingConfidence,omitempty" name:"EndingConfidence"`
}

type PicMarkInfoItem struct {
	// 图片水印的X坐标。
	PosX *int64 `json:"PosX,omitempty" name:"PosX"`

	// 图片水印的Y坐标 。
	PosY *int64 `json:"PosY,omitempty" name:"PosY"`

	// 图片水印路径,，如果不从Cos拉取水印，则必填
	Path *string `json:"Path,omitempty" name:"Path"`

	// 图片水印的Cos 信息，从Cos上拉取图片水印时必填。
	CosInfo *CosInfo `json:"CosInfo,omitempty" name:"CosInfo"`

	// 图片水印宽度，不填为图片原始宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 图片水印高度，不填为图片原始高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 添加图片水印的开始时间,单位：ms。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 添加图片水印的结束时间,单位：ms。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type QualityControlInfo struct {
	// 对流进行截图的间隔ms，默认1000ms
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// 是否保存截图
	VideoShot *bool `json:"VideoShot,omitempty" name:"VideoShot"`

	// 是否检测抖动重影
	Jitter *bool `json:"Jitter,omitempty" name:"Jitter"`

	// 是否检测模糊
	Blur *bool `json:"Blur,omitempty" name:"Blur"`

	// 是否检测低光照、过曝
	AbnormalLighting *bool `json:"AbnormalLighting,omitempty" name:"AbnormalLighting"`

	// 是否检测花屏
	CrashScreen *bool `json:"CrashScreen,omitempty" name:"CrashScreen"`

	// 是否检测黑边、白边、黑屏、白屏、绿屏
	BlackWhiteEdge *bool `json:"BlackWhiteEdge,omitempty" name:"BlackWhiteEdge"`

	// 是否检测噪点
	Noise *bool `json:"Noise,omitempty" name:"Noise"`

	// 是否检测马赛克
	Mosaic *bool `json:"Mosaic,omitempty" name:"Mosaic"`

	// 是否检测二维码，包括小程序码、条形码
	QRCode *bool `json:"QRCode,omitempty" name:"QRCode"`

	// 是否开启画面质量评价
	QualityEvaluation *bool `json:"QualityEvaluation,omitempty" name:"QualityEvaluation"`

	// 画面质量评价过滤阈值，结果只返回低于阈值的时间段，默认60
	QualityEvalScore *uint64 `json:"QualityEvalScore,omitempty" name:"QualityEvalScore"`

	// 是否检测视频音频，包含静音、低音、爆音
	Voice *bool `json:"Voice,omitempty" name:"Voice"`
}

type QualityControlInfoTaskResult struct {
	// 质检任务 ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 质检任务状态。
	// 1：执行中；2：成功；3：失败
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 表示处理进度百分比
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 处理时长(s)
	UsedTime *uint64 `json:"UsedTime,omitempty" name:"UsedTime"`

	// 计费时长(s)
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 为true时表示视频无音频轨
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoAudio *bool `json:"NoAudio,omitempty" name:"NoAudio"`

	// 为true时表示视频无视频轨
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoVideo *bool `json:"NoVideo,omitempty" name:"NoVideo"`

	// 视频无参考质量打分，百分制
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityEvaluationScore *uint64 `json:"QualityEvaluationScore,omitempty" name:"QualityEvaluationScore"`

	// 视频画面无参考评分低于阈值的时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityEvaluationResults []*QualityControlResultItems `json:"QualityEvaluationResults,omitempty" name:"QualityEvaluationResults"`

	// 视频画面抖动时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	JitterResults []*QualityControlResultItems `json:"JitterResults,omitempty" name:"JitterResults"`

	// 视频画面模糊时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlurResults []*QualityControlResultItems `json:"BlurResults,omitempty" name:"BlurResults"`

	// 视频画面低光、过曝时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	AbnormalLightingResults []*QualityControlResultItems `json:"AbnormalLightingResults,omitempty" name:"AbnormalLightingResults"`

	// 视频画面花屏时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrashScreenResults []*QualityControlResultItems `json:"CrashScreenResults,omitempty" name:"CrashScreenResults"`

	// 视频画面黑边、白边、黑屏、白屏、纯色屏时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlackWhiteEdgeResults []*QualityControlResultItems `json:"BlackWhiteEdgeResults,omitempty" name:"BlackWhiteEdgeResults"`

	// 视频画面有噪点时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoiseResults []*QualityControlResultItems `json:"NoiseResults,omitempty" name:"NoiseResults"`

	// 视频画面有马赛克时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	MosaicResults []*QualityControlResultItems `json:"MosaicResults,omitempty" name:"MosaicResults"`

	// 视频画面有二维码的时间段，包括小程序码、条形码
	// 注意：此字段可能返回 null，表示取不到有效值。
	QRCodeResults []*QualityControlResultItems `json:"QRCodeResults,omitempty" name:"QRCodeResults"`

	// 视频音频异常时间段，包括静音、低音、爆音
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceResults []*QualityControlResultItems `json:"VoiceResults,omitempty" name:"VoiceResults"`

	// 任务错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 任务错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type QualityControlItem struct {
	// 置信度，取值范围是 0 到 100
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *uint64 `json:"Confidence,omitempty" name:"Confidence"`

	// 出现的起始时间戳，秒
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 出现的结束时间戳，秒
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 区域坐标(px)，即左上角坐标、右下角坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCoordsSet []*uint64 `json:"AreaCoordsSet,omitempty" name:"AreaCoordsSet"`
}

type QualityControlResultItems struct {
	// 异常类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 质检结果项
	QualityControlItems []*QualityControlItem `json:"QualityControlItems,omitempty" name:"QualityControlItems"`
}

type RemoveReverb struct {
	// 去混响类型，可选项：normal
	Type *string `json:"Type,omitempty" name:"Type"`
}

type ResultAudioInfo struct {
	// 流在媒体文件中的流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamId *int64 `json:"StreamId,omitempty" name:"StreamId"`

	// 流的时长，单位：毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
}

type ResultVideoInfo struct {
	// 流在媒体文件中的流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamId *int64 `json:"StreamId,omitempty" name:"StreamId"`

	// 流的时长，单位：毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 画面宽度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 画面高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频帧率，如果高于原始帧率，部分服务将无效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`
}

type SaveInfo struct {
	// 存储类型，可选值： 
	// 1：CosInfo。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Cos形式存储信息，当Type等于1时必选。
	CosInfo *CosInfo `json:"CosInfo,omitempty" name:"CosInfo"`

	// 存储信息ID标记，用于多个输出场景。部分任务支持多输出时，一般要求必选。
	// ID只能包含字母、数字、下划线、中划线，长读不能超过128。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type ScratchRepair struct {
	// 去划痕方式，取值：normal。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 去划痕强度， 可选项：0.0-1.0。小于0.0的默认为0.0，大于1.0的默认为1.0。
	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`
}

type SectionTime struct {
	// 开始时间点，单位ms
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 时间区间时长，单位ms
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
}

type SegmentInfo struct {
	// 每个切片平均时长，默认10s。
	FragmentTime *int64 `json:"FragmentTime,omitempty" name:"FragmentTime"`

	// 切片类型，可选项：hls，不填时默认hls。
	SegmentType *string `json:"SegmentType,omitempty" name:"SegmentType"`

	// 切片文件名字。注意：
	// 1.不填切片文件名时，默认按照按照如下格式命名：m3u8文件名{order}。
	// 2.若填了切片文件名字，则会按照如下格式命名：用户指定文件名{order}。
	FragmentName *string `json:"FragmentName,omitempty" name:"FragmentName"`
}

type Sharp struct {
	// 细节增强方式,取值：normal。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 细节增强强度，可选项：0.0-1.0。小于0.0的默认为0.0，大于1.0的默认为1.0。
	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`
}

type SpriteImageInfo struct {
	// 表示雪碧图行数，默认：10。
	RowCount *uint64 `json:"RowCount,omitempty" name:"RowCount"`

	// 表示雪碧图列数，默认：10。
	ColumnCount *uint64 `json:"ColumnCount,omitempty" name:"ColumnCount"`

	// 第一行元素与顶部像素距离，默认：0。
	MarginTop *int64 `json:"MarginTop,omitempty" name:"MarginTop"`

	// 最后一行元素与底部像素距离，默认：0。
	MarginBottom *int64 `json:"MarginBottom,omitempty" name:"MarginBottom"`

	// 最左一行元素与左边像素距离，默认：0。
	MarginLeft *int64 `json:"MarginLeft,omitempty" name:"MarginLeft"`

	// 最右一行元素与右边像素距离，默认：0。
	MarginRight *int64 `json:"MarginRight,omitempty" name:"MarginRight"`

	// 小图与元素顶部像素距离，默认：0。
	PaddingTop *int64 `json:"PaddingTop,omitempty" name:"PaddingTop"`

	// 小图与元素底部像素距离，默认：0。
	PaddingBottom *int64 `json:"PaddingBottom,omitempty" name:"PaddingBottom"`

	// 小图与元素左边像素距离，默认：0。
	PaddingLeft *int64 `json:"PaddingLeft,omitempty" name:"PaddingLeft"`

	// 小图与元素右边像素距离，默认：0。
	PaddingRight *int64 `json:"PaddingRight,omitempty" name:"PaddingRight"`

	// 背景颜色，格式：#RRGGBB，默认：#FFFFFF。
	BackgroundColor *string `json:"BackgroundColor,omitempty" name:"BackgroundColor"`
}

// Predefined struct for user
type StopMediaProcessTaskRequestParams struct {
	// 编辑处理任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type StopMediaProcessTaskRequest struct {
	*tchttp.BaseRequest
	
	// 编辑处理任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopMediaProcessTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMediaProcessTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMediaProcessTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMediaProcessTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopMediaProcessTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopMediaProcessTaskResponseParams `json:"Response"`
}

func (r *StopMediaProcessTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMediaProcessTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMediaQualityRestorationTaskRequestParams struct {
	// 要删除的画质重生任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type StopMediaQualityRestorationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的画质重生任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopMediaQualityRestorationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMediaQualityRestorationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMediaQualityRestorationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMediaQualityRestorationTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopMediaQualityRestorationTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopMediaQualityRestorationTaskResponseParams `json:"Response"`
}

func (r *StopMediaQualityRestorationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMediaQualityRestorationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StripEditingInfo struct {
	// 是否开启智能拆条。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
}

type StripTaskResult struct {
	// 编辑任务状态。 
	// 1：执行中；2：成功；3：失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 编辑任务失败错误码。 
	// 0：成功；其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 编辑任务失败错误描述。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 智能拆条结果集。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSet []*StripTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet"`
}

type StripTaskResultItem struct {
	// 视频拆条片段地址。
	SegmentUrl *string `json:"SegmentUrl,omitempty" name:"SegmentUrl"`

	// 拆条封面图片地址。
	CovImgUrl *string `json:"CovImgUrl,omitempty" name:"CovImgUrl"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 拆条片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 拆条片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type SubTaskResultItem struct {
	// 子任务名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 子任务状态。
	// 0：成功；
	// 1：执行中；
	// 其他值：失败。
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 子任务状态描述。
	StatusMsg *string `json:"StatusMsg,omitempty" name:"StatusMsg"`

	// 子任务进度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgressRate *int64 `json:"ProgressRate,omitempty" name:"ProgressRate"`

	// 画质重生处理后文件的下载地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 画质重生处理后文件的MD5。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 画质重生处理后文件的详细信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileInfo *FileInfo `json:"FileInfo,omitempty" name:"FileInfo"`
}

type SubTaskTranscodeInfo struct {
	// 子任务名称。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 目标文件信息。
	TargetInfo *TargetInfo `json:"TargetInfo,omitempty" name:"TargetInfo"`

	// 视频剪辑信息。注意：如果填写了EditInfo，则VideoInfo和AudioInfo必填
	EditInfo *EditInfo `json:"EditInfo,omitempty" name:"EditInfo"`

	// 视频转码信息，不填保持和源文件一致。
	VideoInfo *VideoInfo `json:"VideoInfo,omitempty" name:"VideoInfo"`

	// 音频转码信息，不填保持和源文件一致。
	AudioInfo *AudioInfo `json:"AudioInfo,omitempty" name:"AudioInfo"`

	// 指定封装信息。
	MuxInfo *MuxInfo `json:"MuxInfo,omitempty" name:"MuxInfo"`
}

type SubtitleItem struct {
	// 语音识别结果
	Id *string `json:"Id,omitempty" name:"Id"`

	// 中文翻译结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zh *string `json:"Zh,omitempty" name:"Zh"`

	// 英文翻译结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	En *string `json:"En,omitempty" name:"En"`

	// 语句起始时间戳PTS(ms)
	StartPts *uint64 `json:"StartPts,omitempty" name:"StartPts"`

	// 语句结束时间戳PTS(ms)
	EndPts *uint64 `json:"EndPts,omitempty" name:"EndPts"`

	// 字符串形式的起始结束时间
	Period *string `json:"Period,omitempty" name:"Period"`

	// 结果的置信度（百分制）
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 当前语句是否结束
	EndFlag *bool `json:"EndFlag,omitempty" name:"EndFlag"`

	// 语句分割时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	PuncEndTs *string `json:"PuncEndTs,omitempty" name:"PuncEndTs"`
}

type SubtitleRec struct {
	// 语音识别：
	// zh：中文
	// en：英文
	AsrDst *string `json:"AsrDst,omitempty" name:"AsrDst"`

	// 翻译识别：
	// zh：中文
	// en：英文
	TransDst *string `json:"TransDst,omitempty" name:"TransDst"`
}

type SubtitleResult struct {
	// 语音字幕数组
	SubtitleItems []*SubtitleItem `json:"SubtitleItems,omitempty" name:"SubtitleItems"`
}

type TagEditingInfo struct {
	// 是否开启视频标签识别。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
}

type TagItem struct {
	// 标签内容
	Id *string `json:"Id,omitempty" name:"Id"`

	// 结果的置信度（百分制）
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 分级数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Categorys []*string `json:"Categorys,omitempty" name:"Categorys"`

	// 标签备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ext *string `json:"Ext,omitempty" name:"Ext"`
}

type TagTaskResult struct {
	// 编辑任务状态。 
	// 1：执行中；2：成功；3：失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 编辑任务失败错误码。 
	// 0：成功；其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 编辑任务失败错误描述。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 视频标签识别结果集。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSet []*TagTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet"`
}

type TagTaskResultItem struct {
	// 标签名称。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type TargetInfo struct {
	// 目标文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 目标文件切片信息
	SegmentInfo *SegmentInfo `json:"SegmentInfo,omitempty" name:"SegmentInfo"`
}

type TargetVideoInfo struct {
	// 视频宽度，单位像素，一般要求是偶数，否则会向下对齐。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 视频高度，单位像素，一般要求是偶数，否则会向下对齐。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频帧率，范围在1到120之间
	FrameRate *int64 `json:"FrameRate,omitempty" name:"FrameRate"`
}

type TaskResultFile struct {
	// 文件链接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 文件大小，部分任务支持，单位：字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 媒体信息，对于媒体文件，部分任务支持返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaInfo *MediaResultInfo `json:"MediaInfo,omitempty" name:"MediaInfo"`

	// 文件对应的md5。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type TextMarkInfoItem struct {
	// 文字内容。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 文字水印X坐标。
	PosX *int64 `json:"PosX,omitempty" name:"PosX"`

	// 文字水印Y坐标。
	PosY *int64 `json:"PosY,omitempty" name:"PosY"`

	// 文字大小
	FontSize *int64 `json:"FontSize,omitempty" name:"FontSize"`

	// 字体，可选项：hei,song，simkai,arial；默认hei(黑体）。
	FontFile *string `json:"FontFile,omitempty" name:"FontFile"`

	// 字体颜色，颜色见附录，不填默认black。
	FontColor *string `json:"FontColor,omitempty" name:"FontColor"`

	// 文字透明度，可选值0-1。0：不透明，1：全透明。默认为0
	FontAlpha *float64 `json:"FontAlpha,omitempty" name:"FontAlpha"`
}

type UrlInfo struct {
	// 视频 URL。
	// 注意：编辑理解仅支持mp4、flv等格式的点播文件，不支持hls；
	Url *string `json:"Url,omitempty" name:"Url"`

	// 视频地址格式，可选值： 
	// 0：音视频 ;
	// 1：直播流。 
	// 默认为0。其他非0非1值默认为0。画质重生任务只支持0。
	Format *int64 `json:"Format,omitempty" name:"Format"`

	// 【不再支持】指定请求资源时，HTTP头部host的值。
	Host *string `json:"Host,omitempty" name:"Host"`
}

type VideoEnhance struct {
	// 去编码毛刺、伪影参数。
	ArtifactReduction *ArtifactReduction `json:"ArtifactReduction,omitempty" name:"ArtifactReduction"`

	// 去噪声参数。
	Denoising *Denoising `json:"Denoising,omitempty" name:"Denoising"`

	// 颜色增强参数。
	ColorEnhance *ColorEnhance `json:"ColorEnhance,omitempty" name:"ColorEnhance"`

	// 细节增强参数。
	Sharp *Sharp `json:"Sharp,omitempty" name:"Sharp"`

	// 超分参数，可选项：2，目前仅支持2倍超分。
	// 注意：此参数已经弃用，超分可以使用VideoSuperResolution参数
	WdSuperResolution *int64 `json:"WdSuperResolution,omitempty" name:"WdSuperResolution"`

	// 人脸保护信息。
	FaceProtect *FaceProtect `json:"FaceProtect,omitempty" name:"FaceProtect"`

	// 插帧，取值范围：[0, 60]，单位：Hz。
	// 注意：当取值为 0，表示帧率和原始视频保持一致。
	WdFps *int64 `json:"WdFps,omitempty" name:"WdFps"`

	// 去划痕参数
	ScratchRepair *ScratchRepair `json:"ScratchRepair,omitempty" name:"ScratchRepair"`

	// 低光照增强参数
	LowLightEnhance *LowLightEnhance `json:"LowLightEnhance,omitempty" name:"LowLightEnhance"`

	// 视频超分参数
	VideoSuperResolution *VideoSuperResolution `json:"VideoSuperResolution,omitempty" name:"VideoSuperResolution"`

	// 视频画质修复参数
	VideoRepair *VideoRepair `json:"VideoRepair,omitempty" name:"VideoRepair"`
}

type VideoInfo struct {
	// 视频帧率，取值范围：[0, 60]，单位：Hz。
	// 注意：当取值为 0，表示帧率和原始视频保持一致。
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 宽度，取值范围：0 和 [128, 4096]
	// 注意：
	// 当 Width、Height 均为 0，则分辨率同源；
	// 当 Width 为 0，Height 非 0，则 Width 按比例缩放；
	// 当 Width 非 0，Height 为 0，则 Height 按比例缩放；
	// 当 Width、Height 均非 0，则分辨率按用户指定。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高度，取值范围：0 和 [128, 4096]
	// 注意：
	// 当 Width、Height 均为 0，则分辨率同源；
	// 当 Width 为 0，Height 非 0，则 Width 按比例缩放；
	// 当 Width 非 0，Height 为 0，则 Height 按比例缩放；
	// 当 Width、Height 均非 0，则分辨率按用户指定。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 长边分辨率，取值范围：0 和 [128, 4096]
	// 注意：
	// 当 LongSide、ShortSide 均为 0，则分辨率按照Width，Height；
	// 当 LongSide 为 0，ShortSide 非 0，则 LongSide 按比例缩放；
	// 当 LongSide非 0，ShortSide为 0，则 ShortSide 按比例缩放；
	// 当 LongSide、ShortSide 均非 0，则分辨率按用户指定。
	// 长短边优先级高于Weight,Height,设置长短边则忽略宽高。
	LongSide *int64 `json:"LongSide,omitempty" name:"LongSide"`

	// 短边分辨率，取值范围：0 和 [128, 4096]
	// 注意：
	// 当 LongSide、ShortSide 均为 0，则分辨率按照Width，Height；
	// 当 LongSide 为 0，ShortSide 非 0，则 LongSide 按比例缩放；
	// 当 LongSide非 0，ShortSide为 0，则 ShortSide 按比例缩放；
	// 当 LongSide、ShortSide 均非 0，则分辨率按用户指定。
	// 长短边优先级高于Weight,Height,设置长短边则忽略宽高。
	ShortSide *int64 `json:"ShortSide,omitempty" name:"ShortSide"`

	// 视频流的码率，取值范围：0 和 [128, 35000]，单位：kbps。当取值为 0，表示视频码率和原始视频保持一致。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 固定I帧之间，视频帧数量，取值范围： [25, 2500]，如果不填，使用编码默认最优序列。
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// 编码器支持选项，可选值：
	// h264,
	// h265,
	// av1。
	// 不填默认h264。
	VideoCodec *string `json:"VideoCodec,omitempty" name:"VideoCodec"`

	// 图片水印。
	PicMarkInfo []*PicMarkInfoItem `json:"PicMarkInfo,omitempty" name:"PicMarkInfo"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。
	DarInfo *DarInfo `json:"DarInfo,omitempty" name:"DarInfo"`

	// 支持hdr,可选项：
	// hdr10,
	// hlg。
	// 此时，VideoCodec会强制设置为h265, 编码位深为10
	Hdr *string `json:"Hdr,omitempty" name:"Hdr"`

	// 画质增强参数信息。
	VideoEnhance *VideoEnhance `json:"VideoEnhance,omitempty" name:"VideoEnhance"`

	// 数字水印参数信息。
	HiddenMarkInfo *HiddenMarkInfo `json:"HiddenMarkInfo,omitempty" name:"HiddenMarkInfo"`

	// 文本水印参数信息。
	TextMarkInfo []*TextMarkInfoItem `json:"TextMarkInfo,omitempty" name:"TextMarkInfo"`
}

type VideoInfoResultItem struct {
	// 视频流的流id。
	Stream *int64 `json:"Stream,omitempty" name:"Stream"`

	// 视频宽度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 视频高度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频码率，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频帧率，用分数格式表示，如：25/1, 99/32等等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fps *string `json:"Fps,omitempty" name:"Fps"`

	// 编码格式，如h264,h265等等 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 播放旋转角度，可选值0-360。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 视频时长，单位：ms 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 颜色空间，如yuv420p，yuv444p等等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PixFormat *string `json:"PixFormat,omitempty" name:"PixFormat"`
}

type VideoRepair struct {
	// 画质修复类型，可选值：weak，normal，strong;
	// 默认值: weak
	Type *string `json:"Type,omitempty" name:"Type"`
}

type VideoSuperResolution struct {
	// 超分视频类型：可选值：lq,hq
	// lq: 针对低清晰度有较多噪声视频的超分;
	// hq: 针对高清晰度视频超分;
	// 默认取值：lq。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 超分倍数，可选值：2。
	// 注意：当前只支持两倍超分。
	Size *int64 `json:"Size,omitempty" name:"Size"`
}