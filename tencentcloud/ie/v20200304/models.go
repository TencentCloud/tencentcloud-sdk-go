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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

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
	ItemSet []*ClassificationTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet" list`
}

type ClassificationTaskResultItem struct {

	// 分类名称。
	Classification *string `json:"Classification,omitempty" name:"Classification"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type CosAuthMode struct {

	// 授权类型，可选值： 
	// 0：bucket授权，需要将对应bucket授权给本服务帐号（3020447271），否则会读写cos失败； 
	// 1：key托管，把cos的账号id和key托管于本服务，本服务会提供一个托管id； 
	// 3：临时key授权。
	// 注意：目前智能编辑还不支持临时key授权。
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
	ItemSet []*CoverTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet" list`
}

type CoverTaskResultItem struct {

	// 智能封面地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
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

func (r *CreateEditingTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateEditingTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 编辑任务 ID，可以通过该 ID 查询任务状态。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEditingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEditingTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeEditingTaskResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEditingTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 编辑任务结果信息。
		TaskResult *EditingTaskResult `json:"TaskResult,omitempty" name:"TaskResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEditingTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEditingTaskResultResponse) FromJsonString(s string) error {
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
	ItemSet []*HighlightsTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet" list`
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
	SegmentSet []*HighlightsTaskResultItemSegment `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type HighlightsTaskResultItemSegment struct {

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 集锦片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 集锦片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
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

type SaveInfo struct {

	// 存储类型，可选值： 
	// 1：CosInfo。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Cos形式存储信息，当Type等于1时必选。
	CosInfo *CosInfo `json:"CosInfo,omitempty" name:"CosInfo"`
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
	ItemSet []*StripTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet" list`
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

type TagEditingInfo struct {

	// 是否开启视频标签识别。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
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
	ItemSet []*TagTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet" list`
}

type TagTaskResultItem struct {

	// 标签名称。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type UrlInfo struct {

	// 视频 URL。音视频支持mp4、ts等格式；直播流支持flv、rtmp格式。
	// 注意：目前智能编辑还不支持直播流场景。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 视频地址格式，可选值： 
	// 0：音视频 ;
	// 1：直播流。 
	// 默认为0。其他非0非1值默认为0。
	Format *int64 `json:"Format,omitempty" name:"Format"`

	// 指定请求资源时，HTTP头部host的值。
	Host *string `json:"Host,omitempty" name:"Host"`
}
