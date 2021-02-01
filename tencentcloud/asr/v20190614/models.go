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

package v20190614

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AsyncRecognitionTaskInfo struct {

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 音频流Url
	Url *string `json:"Url,omitempty" name:"Url"`
}

type AsyncRecognitionTasks struct {

	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*AsyncRecognitionTaskInfo `json:"Tasks,omitempty" name:"Tasks" list`
}

type CloseAsyncRecognitionTaskRequest struct {
	*tchttp.BaseRequest

	// 语音流异步识别任务的唯一标识，在创建任务时会返回
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *CloseAsyncRecognitionTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseAsyncRecognitionTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloseAsyncRecognitionTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseAsyncRecognitionTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseAsyncRecognitionTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAsrVocabRequest struct {
	*tchttp.BaseRequest

	// 热词表名称，长度在1-255之间
	Name *string `json:"Name,omitempty" name:"Name"`

	// 热词表描述，长度在0-1000之间
	Description *string `json:"Description,omitempty" name:"Description"`

	// 词权重数组，包含全部的热词和对应的权重。每个热词的长度不大于10，权重为[1,10]之间整数，数组长度不大于128
	WordWeights []*HotWord `json:"WordWeights,omitempty" name:"WordWeights" list`

	// 词权重文件（纯文本文件）的二进制base64编码，以行分隔，每行的格式为word|weight，即以英文符号|为分割，左边为词，右边为权重，如：你好|5。
	// 当用户传此参数（参数长度大于0），即以此参数解析词权重，WordWeights会被忽略
	WordWeightStr *string `json:"WordWeightStr,omitempty" name:"WordWeightStr"`
}

func (r *CreateAsrVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAsrVocabRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAsrVocabResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 词表ID，可用于获取词表信息
		VocabId *string `json:"VocabId,omitempty" name:"VocabId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAsrVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAsrVocabResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAsyncRecognitionTaskRequest struct {
	*tchttp.BaseRequest

	// 引擎模型类型。
	// • 16k_zh：16k 中文普通话通用；
	// • 16k_zh_video：16k 音视频领域；
	// • 16k_en：16k 英语；
	// • 16k_ca：16k 粤语；
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 语音流地址，支持rtmp、hls、rtsp等流媒体协议，以及各类基于http协议的直播流
	Url *string `json:"Url,omitempty" name:"Url"`

	// 支持HTTP和HTTPS协议，用于接收识别结果，您需要自行搭建公网可调用的服务。回调格式&内容详见：[语音流异步识别回调说明](https://github.com/yunjianfei/qcloud-documents/blob/master/product/%E5%A4%A7%E6%95%B0%E6%8D%AE%E4%B8%8EAI/%E8%AF%AD%E9%9F%B3%E8%AF%86%E5%88%AB/%E8%AF%AD%E9%9F%B3%E8%AF%86%E5%88%AB%20API%202017/%E8%AF%AD%E9%9F%B3%E6%B5%81%E5%BC%82%E6%AD%A5%E8%AF%86%E5%88%AB%E5%9B%9E%E8%B0%83%E8%AF%B4%E6%98%8E.md)
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 用于生成回调通知中的签名
	SignToken *string `json:"SignToken,omitempty" name:"SignToken"`

	// 是否过滤脏词（目前支持中文普通话引擎）。0：不过滤脏词；1：过滤脏词；2：将脏词替换为 * 。默认值为 0
	FilterDirty *int64 `json:"FilterDirty,omitempty" name:"FilterDirty"`

	// 是否过语气词（目前支持中文普通话引擎）。0：不过滤语气词；1：部分过滤；2：严格过滤 。默认值为 0
	FilterModal *int64 `json:"FilterModal,omitempty" name:"FilterModal"`

	// 是否过滤标点符号（目前支持中文普通话引擎）。 0：不过滤，1：过滤句末标点，2：过滤所有标点。默认为0
	FilterPunc *int64 `json:"FilterPunc,omitempty" name:"FilterPunc"`

	// 是否进行阿拉伯数字智能转换。0：不转换，直接输出中文数字，1：根据场景智能转换为阿拉伯数字。默认值为1
	ConvertNumMode *int64 `json:"ConvertNumMode,omitempty" name:"ConvertNumMode"`

	// 是否显示词级别时间戳。0：不显示；1：显示，不包含标点时间戳，2：显示，包含标点时间戳。默认为0
	WordInfo *int64 `json:"WordInfo,omitempty" name:"WordInfo"`

	// 热词id。用于调用对应的热词表，如果在调用语音识别服务时，不进行单独的热词id设置，自动生效默认热词；如果进行了单独的热词id设置，那么将生效单独设置的热词id。
	HotwordId *string `json:"HotwordId,omitempty" name:"HotwordId"`
}

func (r *CreateAsyncRecognitionTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAsyncRecognitionTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAsyncRecognitionTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求返回结果，包含本次的任务ID(TaskId)
		Data *Task `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAsyncRecognitionTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAsyncRecognitionTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCustomizationRequest struct {
	*tchttp.BaseRequest

	// 自学习模型名称，需在1-20字符之间
	ModelName *string `json:"ModelName,omitempty" name:"ModelName"`

	// 文本文件的下载地址，服务会从该地址下载文件， 以训练模型，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitempty" name:"TextUrl"`

	// 自学习模型类型，填写8k或者16k
	ModelType *string `json:"ModelType,omitempty" name:"ModelType"`

	// 标签信息
	TagInfos []*string `json:"TagInfos,omitempty" name:"TagInfos" list`
}

func (r *CreateCustomizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCustomizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCustomizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模型ID
		ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCustomizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRecTaskRequest struct {
	*tchttp.BaseRequest

	// 引擎模型类型。
	// 电话场景：
	// • 8k_en：电话 8k 英语；
	// • 8k_zh：电话 8k 中文普通话通用；
	// 非电话场景：
	// • 16k_zh：16k 中文普通话通用；
	// • 16k_zh_video：16k 音视频领域；
	// • 16k_en：16k 英语；
	// • 16k_ca：16k 粤语；
	// • 16k_ja：16k 日语；
	// • 16k_wuu-SH：16k 上海话方言；
	// • 16k_zh_edu 中文教育；
	// • 16k_en_edu 英文教育；
	// • 16k_zh_medical  医疗；
	// • 16k_th 泰语；
	EngineModelType *string `json:"EngineModelType,omitempty" name:"EngineModelType"`

	// 识别声道数。1：单声道；2：双声道（仅支持 8k_zh 引擎模）。注意：录音识别会自动将音频转码为填写的识别声道数
	ChannelNum *uint64 `json:"ChannelNum,omitempty" name:"ChannelNum"`

	// 识别结果返回形式。0： 识别结果文本(含分段时间戳)； 1：词级别粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)(不含标点，含语速值)；2：词级别粒度的详细识别结果（包含标点、语速值）
	ResTextFormat *uint64 `json:"ResTextFormat,omitempty" name:"ResTextFormat"`

	// 语音数据来源。0：语音 URL；1：语音数据（post body）。
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 是否开启说话人分离，0：不开启，1：开启(仅支持8k_zh，16k_zh，16k_zh_video引擎模型，单声道音频)
	SpeakerDiarization *int64 `json:"SpeakerDiarization,omitempty" name:"SpeakerDiarization"`

	// 说话人分离人数（需配合开启说话人分离使用），取值范围：0-10，0代表自动分离（目前仅支持≤6个人），1-10代表指定说话人数分离。
	// 注：话者分离目前是beta版本，请根据您的需要谨慎使用
	SpeakerNumber *int64 `json:"SpeakerNumber,omitempty" name:"SpeakerNumber"`

	// 回调 URL，用户自行搭建的用于接收识别结果的服务URL。如果用户使用轮询方式获取识别结果，则无需提交该参数。回调格式&内容详见：[录音识别回调说明](https://cloud.tencent.com/document/product/1093/52632)
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 语音的URL地址，需要公网可下载。长度小于2048字节，当 SourceType 值为 0 时须填写该字段，为 1 时不需要填写。注意：请确保录音文件时长在5个小时之内，否则可能识别失败。请保证文件的下载速度，否则可能下载失败。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 语音数据，当SourceType 值为1时必须填写，为0可不写。要base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。音频数据要小于5MB。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 数据长度，非必填（此数据长度为数据未进行base64编码时的数据长度）。
	DataLen *uint64 `json:"DataLen,omitempty" name:"DataLen"`

	// 热词id。用于调用对应的热词表，如果在调用语音识别服务时，不进行单独的热词id设置，自动生效默认热词；如果进行了单独的热词id设置，那么将生效单独设置的热词id。
	HotwordId *string `json:"HotwordId,omitempty" name:"HotwordId"`

	// 是否过滤脏词（目前支持中文普通话引擎）。0：不过滤脏词；1：过滤脏词；2：将脏词替换为 * 。默认值为 0。
	FilterDirty *int64 `json:"FilterDirty,omitempty" name:"FilterDirty"`

	// 是否过语气词（目前支持中文普通话引擎）。0：不过滤语气词；1：部分过滤；2：严格过滤 。默认值为 0。
	FilterModal *int64 `json:"FilterModal,omitempty" name:"FilterModal"`

	// 是否进行阿拉伯数字智能转换（目前支持中文普通话引擎）。0：不转换，直接输出中文数字，1：根据场景智能转换为阿拉伯数字，3: 打开数学相关数字转换。默认值为 1。
	ConvertNumMode *int64 `json:"ConvertNumMode,omitempty" name:"ConvertNumMode"`

	// 附加参数
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 是否过滤标点符号（目前支持中文普通话引擎）。 0：不过滤，1：过滤句末标点，2：过滤所有标点。默认为0。
	FilterPunc *int64 `json:"FilterPunc,omitempty" name:"FilterPunc"`
}

func (r *CreateRecTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRecTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRecTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录音文件识别的请求返回结果，包含结果查询需要的TaskId
		Data *Task `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRecTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRecTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAsrVocabRequest struct {
	*tchttp.BaseRequest

	// 热词表Id
	VocabId *string `json:"VocabId,omitempty" name:"VocabId"`
}

func (r *DeleteAsrVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAsrVocabRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAsrVocabResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAsrVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAsrVocabResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomizationRequest struct {
	*tchttp.BaseRequest

	// 要删除的模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`
}

func (r *DeleteCustomizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCustomizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCustomizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncRecognitionTasksRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAsyncRecognitionTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAsyncRecognitionTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncRecognitionTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *AsyncRecognitionTasks `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncRecognitionTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAsyncRecognitionTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 从CreateRecTask接口获取的TaskId，用于获取任务状态与结果。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录音文件识别的请求返回结果。
		Data *TaskStatus `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadAsrVocabRequest struct {
	*tchttp.BaseRequest

	// 词表ID。
	VocabId *string `json:"VocabId,omitempty" name:"VocabId"`
}

func (r *DownloadAsrVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadAsrVocabRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadAsrVocabResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 词表ID。
		VocabId *string `json:"VocabId,omitempty" name:"VocabId"`

		// 词表权重文件形式的base64值。
		WordWeightStr *string `json:"WordWeightStr,omitempty" name:"WordWeightStr"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadAsrVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadAsrVocabResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadCustomizationRequest struct {
	*tchttp.BaseRequest

	// 自学习模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`
}

func (r *DownloadCustomizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadCustomizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadCustomizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 下载地址
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadCustomizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadCustomizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAsrVocabListRequest struct {
	*tchttp.BaseRequest

	// 标签信息，格式为“$TagKey : $TagValue ”，中间分隔符为“空格”+“:”+“空格”
	TagInfos []*string `json:"TagInfos,omitempty" name:"TagInfos" list`

	// 分页Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetAsrVocabListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAsrVocabListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAsrVocabListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 热词表列表
		VocabList []*Vocab `json:"VocabList,omitempty" name:"VocabList" list`

		// 热词列表总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAsrVocabListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAsrVocabListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAsrVocabRequest struct {
	*tchttp.BaseRequest

	// 热词表ID
	VocabId *string `json:"VocabId,omitempty" name:"VocabId"`
}

func (r *GetAsrVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAsrVocabRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAsrVocabResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 热词表名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 热词表描述
		Description *string `json:"Description,omitempty" name:"Description"`

		// 热词表ID
		VocabId *string `json:"VocabId,omitempty" name:"VocabId"`

		// 词权重列表
		WordWeights []*HotWord `json:"WordWeights,omitempty" name:"WordWeights" list`

		// 词表创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 词表更新时间
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 热词表状态，1为默认状态即在识别时默认加载该热词表进行识别，0为初始状态
		State *int64 `json:"State,omitempty" name:"State"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAsrVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAsrVocabResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCustomizationListRequest struct {
	*tchttp.BaseRequest

	// 标签信息，格式为“$TagKey : $TagValue ”，中间分隔符为“空格”+“:”+“空格”
	TagInfos []*string `json:"TagInfos,omitempty" name:"TagInfos" list`

	// 分页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetCustomizationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCustomizationListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCustomizationListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 自学习模型数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*Model `json:"Data,omitempty" name:"Data" list`

		// 自学习模型总量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomizationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCustomizationListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HotWord struct {

	// 热词
	Word *string `json:"Word,omitempty" name:"Word"`

	// 权重
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type Model struct {

	// 模型名称
	ModelName *string `json:"ModelName,omitempty" name:"ModelName"`

	// 模型文件名称
	DictName *string `json:"DictName,omitempty" name:"DictName"`

	// 模型Id
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 模型类型，“8k”或者”16k“
	ModelType *string `json:"ModelType,omitempty" name:"ModelType"`

	// 服务类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 模型状态，-1下线状态，1上线状态, 0训练中, -2 训练失败
	ModelState *int64 `json:"ModelState,omitempty" name:"ModelState"`

	// 最后更新时间
	AtUpdated *string `json:"AtUpdated,omitempty" name:"AtUpdated"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagInfos []*string `json:"TagInfos,omitempty" name:"TagInfos" list`
}

type ModifyCustomizationRequest struct {
	*tchttp.BaseRequest

	// 要修改的模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 要修改的模型名称，长度需在1-20个字符之间
	ModelName *string `json:"ModelName,omitempty" name:"ModelName"`

	// 要修改的模型类型，为8k或者16k
	ModelType *string `json:"ModelType,omitempty" name:"ModelType"`

	// 要修改的模型语料的下载地址，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitempty" name:"TextUrl"`
}

func (r *ModifyCustomizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCustomizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCustomizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomizationStateRequest struct {
	*tchttp.BaseRequest

	// 自学习模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 想要变换的模型状态，-1代表下线，1代表上线
	ToState *int64 `json:"ToState,omitempty" name:"ToState"`
}

func (r *ModifyCustomizationStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCustomizationStateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomizationStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 自学习模型ID
		ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomizationStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCustomizationStateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SentenceDetail struct {

	// 单句最终识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalSentence *string `json:"FinalSentence,omitempty" name:"FinalSentence"`

	// 单句中间识别结果，使用空格拆分为多个词
	// 注意：此字段可能返回 null，表示取不到有效值。
	SliceSentence *string `json:"SliceSentence,omitempty" name:"SliceSentence"`

	// 单句开始时间（毫秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartMs *int64 `json:"StartMs,omitempty" name:"StartMs"`

	// 单句结束时间（毫秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndMs *int64 `json:"EndMs,omitempty" name:"EndMs"`

	// 单句中词个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordsNum *int64 `json:"WordsNum,omitempty" name:"WordsNum"`

	// 单句中词详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Words []*SentenceWords `json:"Words,omitempty" name:"Words" list`

	// 单句语速，单位：字数/秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpeechSpeed *float64 `json:"SpeechSpeed,omitempty" name:"SpeechSpeed"`
}

type SentenceRecognitionRequest struct {
	*tchttp.BaseRequest

	// 腾讯云项目 ID，可填 0，总长度不超过 1024 字节。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 子服务类型。2： 一句话识别。
	SubServiceType *uint64 `json:"SubServiceType,omitempty" name:"SubServiceType"`

	// 引擎模型类型。
	// 电话场景：
	// • 8k_en：电话 8k 英语；
	// • 8k_zh：电话 8k 中文普通话通用；
	// 非电话场景：
	// • 16k_zh：16k 中文普通话通用；
	// • 16k_en：16k 英语；
	// • 16k_ca：16k 粤语；
	// • 16k_ja：16k 日语；
	// •16k_wuu-SH：16k 上海话方言；
	// •16k_zh_medical：16k 医疗。
	EngSerViceType *string `json:"EngSerViceType,omitempty" name:"EngSerViceType"`

	// 语音数据来源。0：语音 URL；1：语音数据（post body）。
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 识别音频的音频格式。mp3、wav。
	VoiceFormat *string `json:"VoiceFormat,omitempty" name:"VoiceFormat"`

	// 用户端对此任务的唯一标识，用户自助生成，用于用户查找识别结果。
	UsrAudioKey *string `json:"UsrAudioKey,omitempty" name:"UsrAudioKey"`

	// 语音 URL，公网可下载。当 SourceType 值为 0（语音 URL上传） 时须填写该字段，为 1 时不填；URL 的长度大于 0，小于 2048，需进行urlencode编码。音频时间长度要小于60s。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 语音数据，当SourceType 值为1（本地语音数据上传）时必须填写，当SourceType 值为0（语音 URL上传）可不写。要使用base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。数据长度要小于3MB（Base64后）。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 数据长度，单位为字节。当 SourceType 值为1（本地语音数据上传）时必须填写，当 SourceType 值为0（语音 URL上传）可不写（此数据长度为数据未进行base64编码时的数据长度）。
	DataLen *int64 `json:"DataLen,omitempty" name:"DataLen"`

	// 热词id。用于调用对应的热词表，如果在调用语音识别服务时，不进行单独的热词id设置，自动生效默认热词；如果进行了单独的热词id设置，那么将生效单独设置的热词id。
	HotwordId *string `json:"HotwordId,omitempty" name:"HotwordId"`

	// 是否过滤脏词（目前支持中文普通话引擎）。0：不过滤脏词；1：过滤脏词；2：将脏词替换为 * 。
	FilterDirty *int64 `json:"FilterDirty,omitempty" name:"FilterDirty"`

	// 是否过语气词（目前支持中文普通话引擎）。0：不过滤语气词；1：部分过滤；2：严格过滤 。
	FilterModal *int64 `json:"FilterModal,omitempty" name:"FilterModal"`

	// 是否过滤标点符号（目前支持中文普通话引擎）。 0：不过滤，1：过滤句末标点，2：过滤所有标点。默认为0。
	FilterPunc *int64 `json:"FilterPunc,omitempty" name:"FilterPunc"`

	// 是否进行阿拉伯数字智能转换。0：不转换，直接输出中文数字，1：根据场景智能转换为阿拉伯数字。默认值为1
	ConvertNumMode *int64 `json:"ConvertNumMode,omitempty" name:"ConvertNumMode"`

	// 是否显示词级别时间戳。0：不显示；1：显示，不包含标点时间戳，2：显示，包含标点时间戳。支持引擎8k_zh，16k_zh，16k_en，16k_ca，16k_ja，16k_wuu-SH
	WordInfo *int64 `json:"WordInfo,omitempty" name:"WordInfo"`
}

func (r *SentenceRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SentenceRecognitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SentenceRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别结果。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 请求的音频时长，单位为ms
		AudioDuration *int64 `json:"AudioDuration,omitempty" name:"AudioDuration"`

		// 词时间戳列表的长度
	// 注意：此字段可能返回 null，表示取不到有效值。
		WordSize *int64 `json:"WordSize,omitempty" name:"WordSize"`

		// 词时间戳列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		WordList []*SentenceWord `json:"WordList,omitempty" name:"WordList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SentenceRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SentenceRecognitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SentenceWord struct {

	// 词结果
	Word *string `json:"Word,omitempty" name:"Word"`

	// 词在音频中的开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 词在音频中的结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type SentenceWords struct {

	// 词文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Word *string `json:"Word,omitempty" name:"Word"`

	// 在句子中的开始时间偏移量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetStartMs *int64 `json:"OffsetStartMs,omitempty" name:"OffsetStartMs"`

	// 在句子中的结束时间偏移量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetEndMs *int64 `json:"OffsetEndMs,omitempty" name:"OffsetEndMs"`
}

type SetVocabStateRequest struct {
	*tchttp.BaseRequest

	// 热词表ID。
	VocabId *string `json:"VocabId,omitempty" name:"VocabId"`

	// 热词表状态，1：设为默认状态；0：设为非默认状态。
	State *int64 `json:"State,omitempty" name:"State"`
}

func (r *SetVocabStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetVocabStateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetVocabStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 热词表ID
		VocabId *string `json:"VocabId,omitempty" name:"VocabId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetVocabStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetVocabStateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Task struct {

	// 任务ID，可通过此ID在轮询接口获取识别状态与结果。注意：TaskId数据类型为uint64
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type TaskStatus struct {

	// 任务标识。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态码，0：任务等待，1：任务执行中，2：任务成功，3：任务失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 任务状态，waiting：任务等待，doing：任务执行中，success：任务成功，failed：任务失败。
	StatusStr *string `json:"StatusStr,omitempty" name:"StatusStr"`

	// 识别结果。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 失败原因说明。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 识别结果详情，包含每个句子中的词时间偏移，一般用于生成字幕的场景。(录音识别请求中ResTextFormat=1时该字段不为空)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultDetail []*SentenceDetail `json:"ResultDetail,omitempty" name:"ResultDetail" list`
}

type UpdateAsrVocabRequest struct {
	*tchttp.BaseRequest

	// 热词表ID
	VocabId *string `json:"VocabId,omitempty" name:"VocabId"`

	// 热词表名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 词权重数组，包含全部的热词和对应的权重。每个热词的长度不大于10，权重为[1,10]之间整数，数组长度不大于128
	WordWeights []*HotWord `json:"WordWeights,omitempty" name:"WordWeights" list`

	// 词权重文件（纯文本文件）的二进制base64编码，以行分隔，每行的格式为word|weight，即以英文符号|为分割，左边为词，右边为权重，如：你好|5。
	// 当用户传此参数（参数长度大于0），即以此参数解析词权重，WordWeights会被忽略
	WordWeightStr *string `json:"WordWeightStr,omitempty" name:"WordWeightStr"`

	// 热词表描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateAsrVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAsrVocabRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAsrVocabResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 热词表ID
		VocabId *string `json:"VocabId,omitempty" name:"VocabId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAsrVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAsrVocabResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Vocab struct {

	// 热词表名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 热词表描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 热词表ID
	VocabId *string `json:"VocabId,omitempty" name:"VocabId"`

	// 词权重列表
	WordWeights []*HotWord `json:"WordWeights,omitempty" name:"WordWeights" list`

	// 词表创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 词表更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 热词表状态，1为默认状态即在识别时默认加载该热词表进行识别，0为初始状态
	State *int64 `json:"State,omitempty" name:"State"`

	// 标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagInfos []*string `json:"TagInfos,omitempty" name:"TagInfos" list`
}
