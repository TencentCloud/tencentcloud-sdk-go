// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AsyncRecognitionTaskInfo struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 音频流Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type AsyncRecognitionTasks struct {
	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*AsyncRecognitionTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`
}

// Predefined struct for user
type CloseAsyncRecognitionTaskRequestParams struct {
	// 语音流异步识别任务的唯一标识，在创建任务时会返回
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type CloseAsyncRecognitionTaskRequest struct {
	*tchttp.BaseRequest
	
	// 语音流异步识别任务的唯一标识，在创建任务时会返回
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *CloseAsyncRecognitionTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAsyncRecognitionTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseAsyncRecognitionTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseAsyncRecognitionTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseAsyncRecognitionTaskResponse struct {
	*tchttp.BaseResponse
	Response *CloseAsyncRecognitionTaskResponseParams `json:"Response"`
}

func (r *CloseAsyncRecognitionTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAsyncRecognitionTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAsrKeyWordLibRequestParams struct {
	// 词表名称，长度在1-20之间
	// 仅限中英文数字-_
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 词文件（纯文本文件）的二进制base64编码，以行分隔
	// 格式要求：TXT
	// 每行只有一个词，不满足格式则报错无法上传
	// 每个词限制**5个汉字，15个字符**，单个词库最多不超过100个词
	// 注意不要有空行，尤其是最后一行
	KeyWordFile *string `json:"KeyWordFile,omitnil,omitempty" name:"KeyWordFile"`
}

type CreateAsrKeyWordLibRequest struct {
	*tchttp.BaseRequest
	
	// 词表名称，长度在1-20之间
	// 仅限中英文数字-_
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 词文件（纯文本文件）的二进制base64编码，以行分隔
	// 格式要求：TXT
	// 每行只有一个词，不满足格式则报错无法上传
	// 每个词限制**5个汉字，15个字符**，单个词库最多不超过100个词
	// 注意不要有空行，尤其是最后一行
	KeyWordFile *string `json:"KeyWordFile,omitnil,omitempty" name:"KeyWordFile"`
}

func (r *CreateAsrKeyWordLibRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAsrKeyWordLibRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "KeyWordFile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAsrKeyWordLibRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAsrKeyWordLibResponseParams struct {
	// 词表ID数据
	Data *KeyWordLibIdData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAsrKeyWordLibResponse struct {
	*tchttp.BaseResponse
	Response *CreateAsrKeyWordLibResponseParams `json:"Response"`
}

func (r *CreateAsrKeyWordLibResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAsrKeyWordLibResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAsrVocabRequestParams struct {
	// 热词表名称，长度在1-255之间
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 热词表描述，长度在0-1000之间
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 词权重数组，包含全部的热词和对应的权重。每个热词的长度不大于10个汉字或30个英文字符，权重为[1,11]之间整数或者100，数组长度不大于1000
	// 注意: 
	// - 热词权重设置为11时，当前热词将升级为超级热词，建议仅将重要且必须生效的热词设置到11，设置过多权重为11的热词将影响整体字准率。
	// - 热词权重设置为100时，当前热词开启热词增强同音替换功能（仅支持8k_zh,16k_zh），举例：热词配置“蜜制|100”时，与“蜜制”同拼音（mizhi）的“秘制”的识别结果会被强制替换成“蜜制”。因此建议客户根据自己的实际情况开启该功能。建议仅将重要且必须生效的热词设置到100，设置过多权重为100的热词将影响整体字准率。
	WordWeights []*HotWord `json:"WordWeights,omitnil,omitempty" name:"WordWeights"`

	// 词权重文件（纯文本文件）的二进制base64编码，以行分隔，每行的格式为word|weight，即以英文符号|为分割，左边为词，右边为权重，如：你好|5。
	// 当用户传此参数（参数长度大于0），即以此参数解析词权重，WordWeights会被忽略
	WordWeightStr *string `json:"WordWeightStr,omitnil,omitempty" name:"WordWeightStr"`
}

type CreateAsrVocabRequest struct {
	*tchttp.BaseRequest
	
	// 热词表名称，长度在1-255之间
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 热词表描述，长度在0-1000之间
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 词权重数组，包含全部的热词和对应的权重。每个热词的长度不大于10个汉字或30个英文字符，权重为[1,11]之间整数或者100，数组长度不大于1000
	// 注意: 
	// - 热词权重设置为11时，当前热词将升级为超级热词，建议仅将重要且必须生效的热词设置到11，设置过多权重为11的热词将影响整体字准率。
	// - 热词权重设置为100时，当前热词开启热词增强同音替换功能（仅支持8k_zh,16k_zh），举例：热词配置“蜜制|100”时，与“蜜制”同拼音（mizhi）的“秘制”的识别结果会被强制替换成“蜜制”。因此建议客户根据自己的实际情况开启该功能。建议仅将重要且必须生效的热词设置到100，设置过多权重为100的热词将影响整体字准率。
	WordWeights []*HotWord `json:"WordWeights,omitnil,omitempty" name:"WordWeights"`

	// 词权重文件（纯文本文件）的二进制base64编码，以行分隔，每行的格式为word|weight，即以英文符号|为分割，左边为词，右边为权重，如：你好|5。
	// 当用户传此参数（参数长度大于0），即以此参数解析词权重，WordWeights会被忽略
	WordWeightStr *string `json:"WordWeightStr,omitnil,omitempty" name:"WordWeightStr"`
}

func (r *CreateAsrVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAsrVocabRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "WordWeights")
	delete(f, "WordWeightStr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAsrVocabRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAsrVocabResponseParams struct {
	// 词表ID，可用于获取词表信息
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAsrVocabResponse struct {
	*tchttp.BaseResponse
	Response *CreateAsrVocabResponseParams `json:"Response"`
}

func (r *CreateAsrVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAsrVocabResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAsyncRecognitionTaskRequestParams struct {
	// 引擎模型类型。
	// • 16k_zh：中文普通话通用；
	// • 16k_en：英语；
	// • 16k_yue：粤语；
	// • 16k_id：印度尼西亚语；
	// • 16k_fil：菲律宾语；
	// • 16k_th：泰语；
	// • 16k_pt：葡萄牙语；
	// • 16k_tr：土耳其语；
	// • 16k_ar：阿拉伯语；
	// • 16k_es：西班牙语；
	// • 16k_hi：印地语；
	// • 16k_fr：法语；
	// • 16k_de：德语；
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 语音流地址，支持rtmp、rtsp等流媒体协议，以及各类基于http协议的直播流(不支持hls, m3u8)
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 支持HTTP和HTTPS协议，用于接收识别结果，您需要自行搭建公网可调用的服务。回调格式&内容详见：[语音流异步识别回调说明](https://cloud.tencent.com/document/product/1093/52633)
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 用于生成回调通知中的签名
	SignToken *string `json:"SignToken,omitnil,omitempty" name:"SignToken"`

	// 是否过滤脏词（目前支持中文普通话引擎）。0：不过滤脏词；1：过滤脏词；2：将脏词替换为 * 。默认值为 0
	FilterDirty *int64 `json:"FilterDirty,omitnil,omitempty" name:"FilterDirty"`

	// 是否过滤语气词（目前支持中文普通话引擎）。0：不过滤语气词；1：部分过滤；2：严格过滤 。默认值为 0
	FilterModal *int64 `json:"FilterModal,omitnil,omitempty" name:"FilterModal"`

	// 是否过滤标点符号（目前支持中文普通话引擎）。 0：不过滤，1：过滤句末标点，2：过滤所有标点。默认为0
	FilterPunc *int64 `json:"FilterPunc,omitnil,omitempty" name:"FilterPunc"`

	// 是否进行阿拉伯数字智能转换。0：不转换，直接输出中文数字，1：根据场景智能转换为阿拉伯数字。默认值为1
	ConvertNumMode *int64 `json:"ConvertNumMode,omitnil,omitempty" name:"ConvertNumMode"`

	// 是否显示词级别时间戳。0：不显示；1：显示，不包含标点时间戳，2：显示，包含标点时间戳。默认为0
	WordInfo *int64 `json:"WordInfo,omitnil,omitempty" name:"WordInfo"`

	// 热词id。用于调用对应的热词表，如果在调用语音识别服务时，不进行单独的热词id设置，自动生效默认热词；如果进行了单独的热词id设置，那么将生效单独设置的热词id。
	HotwordId *string `json:"HotwordId,omitnil,omitempty" name:"HotwordId"`

	// 回调数据中，是否需要对应音频数据。
	AudioData *bool `json:"AudioData,omitnil,omitempty" name:"AudioData"`
}

type CreateAsyncRecognitionTaskRequest struct {
	*tchttp.BaseRequest
	
	// 引擎模型类型。
	// • 16k_zh：中文普通话通用；
	// • 16k_en：英语；
	// • 16k_yue：粤语；
	// • 16k_id：印度尼西亚语；
	// • 16k_fil：菲律宾语；
	// • 16k_th：泰语；
	// • 16k_pt：葡萄牙语；
	// • 16k_tr：土耳其语；
	// • 16k_ar：阿拉伯语；
	// • 16k_es：西班牙语；
	// • 16k_hi：印地语；
	// • 16k_fr：法语；
	// • 16k_de：德语；
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 语音流地址，支持rtmp、rtsp等流媒体协议，以及各类基于http协议的直播流(不支持hls, m3u8)
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 支持HTTP和HTTPS协议，用于接收识别结果，您需要自行搭建公网可调用的服务。回调格式&内容详见：[语音流异步识别回调说明](https://cloud.tencent.com/document/product/1093/52633)
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 用于生成回调通知中的签名
	SignToken *string `json:"SignToken,omitnil,omitempty" name:"SignToken"`

	// 是否过滤脏词（目前支持中文普通话引擎）。0：不过滤脏词；1：过滤脏词；2：将脏词替换为 * 。默认值为 0
	FilterDirty *int64 `json:"FilterDirty,omitnil,omitempty" name:"FilterDirty"`

	// 是否过滤语气词（目前支持中文普通话引擎）。0：不过滤语气词；1：部分过滤；2：严格过滤 。默认值为 0
	FilterModal *int64 `json:"FilterModal,omitnil,omitempty" name:"FilterModal"`

	// 是否过滤标点符号（目前支持中文普通话引擎）。 0：不过滤，1：过滤句末标点，2：过滤所有标点。默认为0
	FilterPunc *int64 `json:"FilterPunc,omitnil,omitempty" name:"FilterPunc"`

	// 是否进行阿拉伯数字智能转换。0：不转换，直接输出中文数字，1：根据场景智能转换为阿拉伯数字。默认值为1
	ConvertNumMode *int64 `json:"ConvertNumMode,omitnil,omitempty" name:"ConvertNumMode"`

	// 是否显示词级别时间戳。0：不显示；1：显示，不包含标点时间戳，2：显示，包含标点时间戳。默认为0
	WordInfo *int64 `json:"WordInfo,omitnil,omitempty" name:"WordInfo"`

	// 热词id。用于调用对应的热词表，如果在调用语音识别服务时，不进行单独的热词id设置，自动生效默认热词；如果进行了单独的热词id设置，那么将生效单独设置的热词id。
	HotwordId *string `json:"HotwordId,omitnil,omitempty" name:"HotwordId"`

	// 回调数据中，是否需要对应音频数据。
	AudioData *bool `json:"AudioData,omitnil,omitempty" name:"AudioData"`
}

func (r *CreateAsyncRecognitionTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAsyncRecognitionTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineType")
	delete(f, "Url")
	delete(f, "CallbackUrl")
	delete(f, "SignToken")
	delete(f, "FilterDirty")
	delete(f, "FilterModal")
	delete(f, "FilterPunc")
	delete(f, "ConvertNumMode")
	delete(f, "WordInfo")
	delete(f, "HotwordId")
	delete(f, "AudioData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAsyncRecognitionTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAsyncRecognitionTaskResponseParams struct {
	// 请求返回结果，包含本次的任务ID(TaskId)
	Data *Task `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAsyncRecognitionTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAsyncRecognitionTaskResponseParams `json:"Response"`
}

func (r *CreateAsyncRecognitionTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAsyncRecognitionTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomizationRequestParams struct {
	// 自学习模型名称，需在1-20字符之间
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 文本文件的下载地址，服务会从该地址下载文件，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitnil,omitempty" name:"TextUrl"`

	// 自学习模型类型，填写8k或者16k
	ModelType *string `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// 标签信息
	//
	// Deprecated: TagInfos is deprecated.
	TagInfos []*string `json:"TagInfos,omitnil,omitempty" name:"TagInfos"`
}

type CreateCustomizationRequest struct {
	*tchttp.BaseRequest
	
	// 自学习模型名称，需在1-20字符之间
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 文本文件的下载地址，服务会从该地址下载文件，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitnil,omitempty" name:"TextUrl"`

	// 自学习模型类型，填写8k或者16k
	ModelType *string `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// 标签信息
	TagInfos []*string `json:"TagInfos,omitnil,omitempty" name:"TagInfos"`
}

func (r *CreateCustomizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelName")
	delete(f, "TextUrl")
	delete(f, "ModelType")
	delete(f, "TagInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomizationResponseParams struct {
	// 模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomizationResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomizationResponseParams `json:"Response"`
}

func (r *CreateCustomizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecTaskRequestParams struct {
	// 引擎模型类型
	// 识别引擎采用分级计费方案，标记为“大模型版”的引擎适用大模型计费方案，[点击这里](https://cloud.tencent.com/document/product/1093/35686) 查看产品计费说明
	// 
	// 电话通讯场景引擎：
	// **注意：电话通讯场景，请务必使用以下8k引擎**
	// • 8k_zh：中文电话通讯；
	// • 8k_en：英文电话通讯；
	// • 8k_zh_large：中文电话场景专用大模型引擎【大模型版】。当前模型同时支持中文、上海话、四川话、武汉话、贵阳话、昆明话、西安话、郑州话、太原话、兰州话、银川话、西宁话、南京话、合肥话、南昌话、长沙话、苏州话、杭州话、济南话、天津话、石家庄话、黑龙江话、吉林话、辽宁话、闽南语、客家话、粤语、南宁话方言识别，通过显著提升模型参数规模与语言建模能力，实现对电话音频中复杂场景（如口音干扰、背景噪声）的高精度识别，识别准确率较常规版本大幅提升。[点击这里](https://console.cloud.tencent.com/asr/demonstrate) 对比常规版本与普方大模型效果；
	// 
	// 注意：如您有电话通讯场景识别需求，但发现需求语种仅支持16k，可将8k音频传入下方16k引擎，亦能获取识别结果。但**16k引擎并非基于电话通讯数据训练，无法承诺此种调用方式的识别效果，需由您自行验证识别结果是否可用**
	// 
	// 通用场景引擎：
	// **注意：除电话通讯场景以外的其它识别场景，请务必使用以下16k引擎**
	// • **16k_zh_en**：中英粤+9种方言大模型引擎【大模型版】。当前模型同时支持中文、英语、粤语、四川、陕西、河南、上海、湖南、湖北、安徽、闽南和潮汕方言识别，模型参数量极大，语言模型性能增强，针对噪声大、回音大、人声小、人声远等低质量音频的识别准确率极大提升;
	// • **16k_zh_large：**普方英大模型引擎【大模型版】。当前模型同时支持中文、英文、[多种中文方言](https://cloud.tencent.com/document/product/1093/35682)等语言的识别，模型参数量极大，语言模型性能增强，针对噪声大、回音大、人声小、人声远等低质量音频的识别准确率极大提升，[点击这里](https://console.cloud.tencent.com/asr/demonstrate) 对比中文普通话常规版本与普方英大模型版本的识别效果；
	// • **16k_multi_lang：**多语种大模型引擎【大模型版】。当前模型同时支持英语、日语、韩语、阿拉伯语、菲律宾语、法语、印地语、印尼语、马来语、葡萄牙语、西班牙语、泰语、土耳其语、越南语、德语的识别，可实现15个语种的自动识别(句子/段落级别)；
	// • **16k_zh：**中文普通话通用引擎，支持中文普通话和少量英语，使用丰富的中文普通话语料训练，覆盖场景广泛，适用于除电话通讯外的所有中文普通话识别场景；
	// • **16k_en：**英语；
	// • **16k_yue：**粤语；
	// • **16k_zh-PY：**中英粤混合引擎，使用一个引擎同时识别中文普通话、英语、粤语三个语言;
	// • **16k_zh-TW**：中文繁体；
	// • **16k_ja：**日语；
	// • **16k_ko：**韩语；
	// • **16k_vi：**越南语；
	// • **16k_ms：**马来语；
	// • **16k_id：**印度尼西亚语；
	// • **16k_fil：**菲律宾语；
	// • **16k_th：**泰语；
	// • **16k_pt：**葡萄牙语；
	// • **16k_tr：**土耳其语；
	// • **16k_ar：**阿拉伯语；
	// • **16k_es：**西班牙语；
	// • **16k_hi：**印地语；
	// • **16k_fr：**法语；
	// • **16k_zh_medical：**中文医疗引擎；
	// • **16k_de：**德语；
	EngineModelType *string `json:"EngineModelType,omitnil,omitempty" name:"EngineModelType"`

	// 识别声道数
	// 1：单声道（16k音频仅支持单声道，**请勿**设置为双声道）；
	// 2：双声道（仅支持8k电话音频，且双声道应分别为通话双方）
	// 
	// 注意：
	// • 16k音频：仅支持单声道识别，**需设置ChannelNum=1**；
	// • 8k电话音频：支持单声道、双声道识别，**建议设置ChannelNum=2，即双声道**。双声道能够物理区分说话人、避免说话双方重叠产生的识别错误，能达到最好的说话人分离效果和识别效果。设置双声道后，将自动区分说话人，因此**无需再开启说话人分离功能**，相关参数（**SpeakerDiarization、SpeakerNumber**）使用默认值即可，返回的ResultDetail中的speakerId的值为0代表左声道，值为1代表右声道。
	ChannelNum *uint64 `json:"ChannelNum,omitnil,omitempty" name:"ChannelNum"`

	// 识别结果返回样式
	// 0：基础识别结果（仅包含有效人声时间戳，无词粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)）；
	// 1：基础识别结果之上，增加词粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)（包含词级别时间戳、语速值，**不含标点**）；
	// 2：基础识别结果之上，增加词粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)（包含词级别时间戳、语速值和标点）；
	// 3：基础识别结果之上，增加词粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)（包含词级别时间戳、语速值和标点），且识别结果按标点符号分段，**适用字幕场景**；
	// 4：**【增值付费功能】**基础识别结果之上，增加词粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)（包含词级别时间戳、语速值和标点），且识别结果按nlp语义分段，**适用会议、庭审记录转写等场景**，仅支持8k_zh/16k_zh引擎
	// 5：**【增值付费功能】**基础识别结果之上，增加词粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)（包含词级别时间戳、语速值和标点），并输出口语转书面语转写结果，该结果去除语气词、重复词、精简冗余表达，并修正发言人口误，实现口语转书面语的效果，**适用于线上、线下会议直接总结为书面会议纪要的场景**，仅支持8k_zh/16k_zh引擎
	// 
	// 注意：
	// 以下引擎仅支持参数值设置为0：16k_multi_lang、16k_ja、16k_ko、16k_vi、16k_ms、16k_id、16k_fil、16k_th、16k_pt、16k_tr、16k_ar、16k_es、16k_hi、16k_fr、16k_zh_medical、16k_de
	// 如果传入参数值4，需确保账号已购买[语义分段资源包](https://cloud.tencent.com/document/product/1093/35686#97ae4aa0-29a0-4066-9f07-ccaf8856a16b)，或账号开启后付费；**若当前账号已开启后付费功能，并传入参数值4，将[自动计费](https://cloud.tencent.com/document/product/1093/35686#d912167d-ffd5-41a9-8b1c-2e89845a6852)**
	// 如果传入参数值5，需确保账号已购买[口语转书面语资源包](https://cloud.tencent.com/document/product/1093/35686#97ae4aa0-29a0-4066-9f07-ccaf8856a16b)，或账号开启后付费；**若当前账号已开启后付费功能，并传入参数值5，将自动计费[自动计费](https://cloud.tencent.com/document/product/1093/35686#d912167d-ffd5-41a9-8b1c-2e89845a6852)**
	ResTextFormat *uint64 `json:"ResTextFormat,omitnil,omitempty" name:"ResTextFormat"`

	// 音频数据来源
	// 0：音频URL；
	// 1：音频数据（post body）
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 音频数据base64编码
	// **当 SourceType 值为 1 时须填写该字段，为 0 时不需要填写**
	// 
	// 注意：音频数据要小于5MB（含）
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 数据长度（此数据长度为数据未进行base64编码时的长度）
	DataLen *uint64 `json:"DataLen,omitnil,omitempty" name:"DataLen"`

	// 音频URL的地址（需要公网环境浏览器可下载）
	// **当 SourceType 值为 0 时须填写该字段，为 1 时不需要填写**
	// 
	// 注意：
	// 1. 请确保录音文件时长在5个小时（含）之内，否则可能识别失败；
	// 2. 请保证文件的下载速度，否则可能下载失败
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 回调 URL
	// 用户自行搭建的用于接收识别结果的服务URL
	// 回调格式和内容详见：[录音识别回调说明](https://cloud.tencent.com/document/product/1093/52632)
	// 
	// 注意：
	// 
	// - 如果用户使用轮询方式获取识别结果，则无需提交该参数
	// - 建议在回调URL中带上您的业务ID等信息，以便处理业务逻辑
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 是否开启说话人分离
	// 0：不开启；
	// 1：开启（仅支持以下引擎：8k_zh/8k_zh_large/16k_zh/16k_ms/16k_en/16k_id/16k_zh_large/16k_zh_dialect/16k_zh_en，且ChannelNum=1时可用）；
	// 3: 开启角色分离，需配合SpeakerRoles参数使用（增值服务，仅支持16k_zh_en引擎，可支持传入声纹对录音文件内的说话人进行角色认证）
	// 默认值为 0
	// 
	// 注意：
	// 8k双声道电话音频请按 **ChannelNum 识别声道数** 的参数描述使用默认值
	SpeakerDiarization *int64 `json:"SpeakerDiarization,omitnil,omitempty" name:"SpeakerDiarization"`

	// 说话人分离人数
	// **需配合开启说话人分离使用，不开启无效**，取值范围：0-10
	// 0：自动分离（最多分离出20个人）；
	// 1-10：指定人数分离；
	// 默认值为 0
	// 注:16k引擎不支持指定人数分离
	SpeakerNumber *int64 `json:"SpeakerNumber,omitnil,omitempty" name:"SpeakerNumber"`

	// 热词表id
	// 如不设置该参数，将自动生效默认热词表；
	// 如设置该参数，将生效对应id的热词表；
	// 点击这里查看[热词表配置方法](https://cloud.tencent.com/document/product/1093/40996)
	HotwordId *string `json:"HotwordId,omitnil,omitempty" name:"HotwordId"`

	// 热词增强功能（目前仅支持8k_zh/16k_zh引擎）
	// 1：开启热词增强功能
	// 
	// 注意：热词增强功能开启后，将对传入的热词表id开启同音替换功能，可以在这里查看[热词表配置方法](https://cloud.tencent.com/document/product/1093/40996)。效果举例：在热词表中配置“蜜制”一词，并开启增强功能，与“蜜制”（mìzhì）同音同调的“秘制”（mìzhì）的识别结果会被强制替换成“蜜制”。**建议客户根据实际的业务需求开启该功能**
	//
	// Deprecated: ReinforceHotword is deprecated.
	ReinforceHotword *int64 `json:"ReinforceHotword,omitnil,omitempty" name:"ReinforceHotword"`

	// 自学习定制模型 id
	// 如设置了该参数，将生效对应id的自学习定制模型；
	// 点击这里查看[自学习定制模型配置方法](https://cloud.tencent.com/document/product/1093/38416)
	CustomizationId *string `json:"CustomizationId,omitnil,omitempty" name:"CustomizationId"`

	// **【增值付费功能】**情绪识别能力（目前仅支持16k_zh,8k_zh）
	// 0：不开启；
	// 1：开启情绪识别，但不在文本展示情绪标签；
	// 2：开启情绪识别，并且在文本展示情绪标签（**该功能需要设置ResTextFormat 大于0**）
	// 默认值为0
	// 支持的情绪分类为：高兴、伤心、愤怒
	// 
	// 注意：
	// 1. **本功能为增值服务**，需将参数设置为1或2时方可按对应方式生效；
	// 2. 如果传入参数值1或2，需确保账号已购买[情绪识别资源包](https://cloud.tencent.com/document/product/1093/35686#97ae4aa0-29a0-4066-9f07-ccaf8856a16b)，或账号开启后付费；**若当前账号已开启后付费功能，并传入参数值1或2，将[自动计费](https://cloud.tencent.com/document/product/1093/35686#d912167d-ffd5-41a9-8b1c-2e89845a6852)）**；
	// 3. 参数设置为0时，无需购买资源包，也不会消耗情绪识别对应资源
	EmotionRecognition *int64 `json:"EmotionRecognition,omitnil,omitempty" name:"EmotionRecognition"`

	// 情绪能量值
	// 取值为音量分贝值/10，取值范围：[1,10]，值越高情绪越强烈
	// 0：不开启；
	// 1：开启；
	// 默认值为0
	EmotionalEnergy *int64 `json:"EmotionalEnergy,omitnil,omitempty" name:"EmotionalEnergy"`

	// 阿拉伯数字智能转换（目前支持中文普通话引擎）
	// 0：不转换，直接输出中文数字；
	// 1：根据场景智能转换为阿拉伯数字；
	// 3：打开数学相关数字转换（如：阿尔法转写为α）；
	// 默认值为 1
	ConvertNumMode *int64 `json:"ConvertNumMode,omitnil,omitempty" name:"ConvertNumMode"`

	// 脏词过滤（目前支持中文普通话引擎）
	// 0：不过滤脏词；
	// 1：过滤脏词；
	// 2：将脏词替换为 * ；
	// 默认值为 0
	FilterDirty *int64 `json:"FilterDirty,omitnil,omitempty" name:"FilterDirty"`

	// 标点符号过滤（目前支持中文普通话引擎）
	// 0：不过滤标点；
	// 1：过滤句末标点；
	// 2：过滤所有标点；
	// 默认值为 0
	FilterPunc *int64 `json:"FilterPunc,omitnil,omitempty" name:"FilterPunc"`

	// 语气词过滤（目前支持中文普通话引擎）
	// 0：不过滤语气词；
	// 1：过滤部分语气词；
	// 2：严格过滤语气词；
	// 默认值为 0
	FilterModal *int64 `json:"FilterModal,omitnil,omitempty" name:"FilterModal"`

	// 单标点最多字数（目前支持中文普通话引擎）
	// **可控制单行字幕最大字数，适用于字幕生成场景**，取值范围：[6，40]
	// 0：不开启该功能；
	// 默认值为0
	// 
	// 注意：需设置ResTextFormat为3，解析返回的ResultDetail列表，通过结构中FinalSentence获取单个标点断句结果
	SentenceMaxLength *int64 `json:"SentenceMaxLength,omitnil,omitempty" name:"SentenceMaxLength"`

	// 附加参数**（该参数无意义，忽略即可）**
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 临时热词表：该参数用于提升识别准确率。
	// 
	// - 单个热词限制："热词|权重"，单个热词不超过30个字符（最多10个汉字），权重[1-11]或者100，如：“腾讯云|5” 或“ASR|11”；
	// 
	// - 临时热词表限制：多个热词用英文逗号分割，最多支持128个热词，如：“腾讯云|10,语音识别|5,ASR|11”；
	// 
	// - 参数 hotword_id（热词表） 与 hotword_list（临时热词表） 区别：
	// 
	//     - hotword_id：热词表。需要先在控制台或接口创建热词表，获得对应hotword_id传入参数来使用热词功能；
	// 
	//     - hotword_list：临时热词表。每次请求时直接传入临时热词表来使用热词功能，云端不保留临时热词表。适用于有极大量热词需求的用户；
	// 
	// 注意：
	// 
	// - 如果同时传入了 hotword_id 和 hotword_list，只有hotword_list 生效；
	// 
	// - 热词权重设置为11时，当前热词将升级为超级热词，建议仅将重要且必须生效的热词设置到11，设置过多权重为11的热词将影响整体字准率。
	// 
	// - 热词权重设置为100时，当前热词开启热词增强同音替换功能（仅支持8k_zh,16k_zh），举例：热词配置“蜜制|100”时，与“蜜制”同拼音（mizhi）的“秘制”的识别结果会被强制替换成“蜜制”。因此建议客户根据自己的实际情况开启该功能。建议仅将重要且必须生效的热词设置到100，设置过多权重为100的热词将影响整体字准率。
	HotwordList *string `json:"HotwordList,omitnil,omitempty" name:"HotwordList"`

	// 关键词识别ID列表，默认空为不进行识别，最多10个
	KeyWordLibIdList []*string `json:"KeyWordLibIdList,omitnil,omitempty" name:"KeyWordLibIdList"`

	// 替换词汇表id,  适用于热词和自学习场景也无法解决的极端case词组,  会对识别结果强制替换。具体可参考[配置控制台](https://console.cloud.tencent.com/asr/replaceword);强制替换功能可能会影响正常识别结果，请谨慎使用
	// 
	// 注意：
	// 1. 本功能配置完成后，预计在10分钟后生效
	ReplaceTextId *string `json:"ReplaceTextId,omitnil,omitempty" name:"ReplaceTextId"`

	// 开启角色分离能力
	// 配合SpeakerDiarization: 3 使用，ASR增值服务，可传入一组声纹信息进行角色认证，仅支持16k_zh_en引擎。
	// 需传入SpeakerRoleInfo数据组，确定说话人的角色信息，涉及RoleAudioUrl和RoleName两个参数。 
	// RoleAudioUrl：需要认证角色的声纹音频地址，建议30s内的纯净人声，最长不能超过45s。 
	// RoleName：需要认证角色的名称，若匹配成功，会替换话者分离中的SpeakerID。 
	// 示例： 
	// "{\"EngineModelType\":\"16k_zh_en\",\"ChannelNum\":1,\"ResTextFormat\":1,\"SourceType\":0,\"Url\":\"需要进行ASR识别的音频链接\",\"SpeakerDiarization\":3,\"SpeakerRoles\":[{\"RoleAudioUrl\":\"需要认证角色的声纹音频地址\",\"RoleName\":\"需要认证角色的名称\"}]}"
	SpeakerRoles []*SpeakerRoleInfo `json:"SpeakerRoles,omitnil,omitempty" name:"SpeakerRoles"`
}

type CreateRecTaskRequest struct {
	*tchttp.BaseRequest
	
	// 引擎模型类型
	// 识别引擎采用分级计费方案，标记为“大模型版”的引擎适用大模型计费方案，[点击这里](https://cloud.tencent.com/document/product/1093/35686) 查看产品计费说明
	// 
	// 电话通讯场景引擎：
	// **注意：电话通讯场景，请务必使用以下8k引擎**
	// • 8k_zh：中文电话通讯；
	// • 8k_en：英文电话通讯；
	// • 8k_zh_large：中文电话场景专用大模型引擎【大模型版】。当前模型同时支持中文、上海话、四川话、武汉话、贵阳话、昆明话、西安话、郑州话、太原话、兰州话、银川话、西宁话、南京话、合肥话、南昌话、长沙话、苏州话、杭州话、济南话、天津话、石家庄话、黑龙江话、吉林话、辽宁话、闽南语、客家话、粤语、南宁话方言识别，通过显著提升模型参数规模与语言建模能力，实现对电话音频中复杂场景（如口音干扰、背景噪声）的高精度识别，识别准确率较常规版本大幅提升。[点击这里](https://console.cloud.tencent.com/asr/demonstrate) 对比常规版本与普方大模型效果；
	// 
	// 注意：如您有电话通讯场景识别需求，但发现需求语种仅支持16k，可将8k音频传入下方16k引擎，亦能获取识别结果。但**16k引擎并非基于电话通讯数据训练，无法承诺此种调用方式的识别效果，需由您自行验证识别结果是否可用**
	// 
	// 通用场景引擎：
	// **注意：除电话通讯场景以外的其它识别场景，请务必使用以下16k引擎**
	// • **16k_zh_en**：中英粤+9种方言大模型引擎【大模型版】。当前模型同时支持中文、英语、粤语、四川、陕西、河南、上海、湖南、湖北、安徽、闽南和潮汕方言识别，模型参数量极大，语言模型性能增强，针对噪声大、回音大、人声小、人声远等低质量音频的识别准确率极大提升;
	// • **16k_zh_large：**普方英大模型引擎【大模型版】。当前模型同时支持中文、英文、[多种中文方言](https://cloud.tencent.com/document/product/1093/35682)等语言的识别，模型参数量极大，语言模型性能增强，针对噪声大、回音大、人声小、人声远等低质量音频的识别准确率极大提升，[点击这里](https://console.cloud.tencent.com/asr/demonstrate) 对比中文普通话常规版本与普方英大模型版本的识别效果；
	// • **16k_multi_lang：**多语种大模型引擎【大模型版】。当前模型同时支持英语、日语、韩语、阿拉伯语、菲律宾语、法语、印地语、印尼语、马来语、葡萄牙语、西班牙语、泰语、土耳其语、越南语、德语的识别，可实现15个语种的自动识别(句子/段落级别)；
	// • **16k_zh：**中文普通话通用引擎，支持中文普通话和少量英语，使用丰富的中文普通话语料训练，覆盖场景广泛，适用于除电话通讯外的所有中文普通话识别场景；
	// • **16k_en：**英语；
	// • **16k_yue：**粤语；
	// • **16k_zh-PY：**中英粤混合引擎，使用一个引擎同时识别中文普通话、英语、粤语三个语言;
	// • **16k_zh-TW**：中文繁体；
	// • **16k_ja：**日语；
	// • **16k_ko：**韩语；
	// • **16k_vi：**越南语；
	// • **16k_ms：**马来语；
	// • **16k_id：**印度尼西亚语；
	// • **16k_fil：**菲律宾语；
	// • **16k_th：**泰语；
	// • **16k_pt：**葡萄牙语；
	// • **16k_tr：**土耳其语；
	// • **16k_ar：**阿拉伯语；
	// • **16k_es：**西班牙语；
	// • **16k_hi：**印地语；
	// • **16k_fr：**法语；
	// • **16k_zh_medical：**中文医疗引擎；
	// • **16k_de：**德语；
	EngineModelType *string `json:"EngineModelType,omitnil,omitempty" name:"EngineModelType"`

	// 识别声道数
	// 1：单声道（16k音频仅支持单声道，**请勿**设置为双声道）；
	// 2：双声道（仅支持8k电话音频，且双声道应分别为通话双方）
	// 
	// 注意：
	// • 16k音频：仅支持单声道识别，**需设置ChannelNum=1**；
	// • 8k电话音频：支持单声道、双声道识别，**建议设置ChannelNum=2，即双声道**。双声道能够物理区分说话人、避免说话双方重叠产生的识别错误，能达到最好的说话人分离效果和识别效果。设置双声道后，将自动区分说话人，因此**无需再开启说话人分离功能**，相关参数（**SpeakerDiarization、SpeakerNumber**）使用默认值即可，返回的ResultDetail中的speakerId的值为0代表左声道，值为1代表右声道。
	ChannelNum *uint64 `json:"ChannelNum,omitnil,omitempty" name:"ChannelNum"`

	// 识别结果返回样式
	// 0：基础识别结果（仅包含有效人声时间戳，无词粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)）；
	// 1：基础识别结果之上，增加词粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)（包含词级别时间戳、语速值，**不含标点**）；
	// 2：基础识别结果之上，增加词粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)（包含词级别时间戳、语速值和标点）；
	// 3：基础识别结果之上，增加词粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)（包含词级别时间戳、语速值和标点），且识别结果按标点符号分段，**适用字幕场景**；
	// 4：**【增值付费功能】**基础识别结果之上，增加词粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)（包含词级别时间戳、语速值和标点），且识别结果按nlp语义分段，**适用会议、庭审记录转写等场景**，仅支持8k_zh/16k_zh引擎
	// 5：**【增值付费功能】**基础识别结果之上，增加词粒度的[详细识别结果](https://cloud.tencent.com/document/api/1093/37824#SentenceDetail)（包含词级别时间戳、语速值和标点），并输出口语转书面语转写结果，该结果去除语气词、重复词、精简冗余表达，并修正发言人口误，实现口语转书面语的效果，**适用于线上、线下会议直接总结为书面会议纪要的场景**，仅支持8k_zh/16k_zh引擎
	// 
	// 注意：
	// 以下引擎仅支持参数值设置为0：16k_multi_lang、16k_ja、16k_ko、16k_vi、16k_ms、16k_id、16k_fil、16k_th、16k_pt、16k_tr、16k_ar、16k_es、16k_hi、16k_fr、16k_zh_medical、16k_de
	// 如果传入参数值4，需确保账号已购买[语义分段资源包](https://cloud.tencent.com/document/product/1093/35686#97ae4aa0-29a0-4066-9f07-ccaf8856a16b)，或账号开启后付费；**若当前账号已开启后付费功能，并传入参数值4，将[自动计费](https://cloud.tencent.com/document/product/1093/35686#d912167d-ffd5-41a9-8b1c-2e89845a6852)**
	// 如果传入参数值5，需确保账号已购买[口语转书面语资源包](https://cloud.tencent.com/document/product/1093/35686#97ae4aa0-29a0-4066-9f07-ccaf8856a16b)，或账号开启后付费；**若当前账号已开启后付费功能，并传入参数值5，将自动计费[自动计费](https://cloud.tencent.com/document/product/1093/35686#d912167d-ffd5-41a9-8b1c-2e89845a6852)**
	ResTextFormat *uint64 `json:"ResTextFormat,omitnil,omitempty" name:"ResTextFormat"`

	// 音频数据来源
	// 0：音频URL；
	// 1：音频数据（post body）
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 音频数据base64编码
	// **当 SourceType 值为 1 时须填写该字段，为 0 时不需要填写**
	// 
	// 注意：音频数据要小于5MB（含）
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 数据长度（此数据长度为数据未进行base64编码时的长度）
	DataLen *uint64 `json:"DataLen,omitnil,omitempty" name:"DataLen"`

	// 音频URL的地址（需要公网环境浏览器可下载）
	// **当 SourceType 值为 0 时须填写该字段，为 1 时不需要填写**
	// 
	// 注意：
	// 1. 请确保录音文件时长在5个小时（含）之内，否则可能识别失败；
	// 2. 请保证文件的下载速度，否则可能下载失败
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 回调 URL
	// 用户自行搭建的用于接收识别结果的服务URL
	// 回调格式和内容详见：[录音识别回调说明](https://cloud.tencent.com/document/product/1093/52632)
	// 
	// 注意：
	// 
	// - 如果用户使用轮询方式获取识别结果，则无需提交该参数
	// - 建议在回调URL中带上您的业务ID等信息，以便处理业务逻辑
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 是否开启说话人分离
	// 0：不开启；
	// 1：开启（仅支持以下引擎：8k_zh/8k_zh_large/16k_zh/16k_ms/16k_en/16k_id/16k_zh_large/16k_zh_dialect/16k_zh_en，且ChannelNum=1时可用）；
	// 3: 开启角色分离，需配合SpeakerRoles参数使用（增值服务，仅支持16k_zh_en引擎，可支持传入声纹对录音文件内的说话人进行角色认证）
	// 默认值为 0
	// 
	// 注意：
	// 8k双声道电话音频请按 **ChannelNum 识别声道数** 的参数描述使用默认值
	SpeakerDiarization *int64 `json:"SpeakerDiarization,omitnil,omitempty" name:"SpeakerDiarization"`

	// 说话人分离人数
	// **需配合开启说话人分离使用，不开启无效**，取值范围：0-10
	// 0：自动分离（最多分离出20个人）；
	// 1-10：指定人数分离；
	// 默认值为 0
	// 注:16k引擎不支持指定人数分离
	SpeakerNumber *int64 `json:"SpeakerNumber,omitnil,omitempty" name:"SpeakerNumber"`

	// 热词表id
	// 如不设置该参数，将自动生效默认热词表；
	// 如设置该参数，将生效对应id的热词表；
	// 点击这里查看[热词表配置方法](https://cloud.tencent.com/document/product/1093/40996)
	HotwordId *string `json:"HotwordId,omitnil,omitempty" name:"HotwordId"`

	// 热词增强功能（目前仅支持8k_zh/16k_zh引擎）
	// 1：开启热词增强功能
	// 
	// 注意：热词增强功能开启后，将对传入的热词表id开启同音替换功能，可以在这里查看[热词表配置方法](https://cloud.tencent.com/document/product/1093/40996)。效果举例：在热词表中配置“蜜制”一词，并开启增强功能，与“蜜制”（mìzhì）同音同调的“秘制”（mìzhì）的识别结果会被强制替换成“蜜制”。**建议客户根据实际的业务需求开启该功能**
	ReinforceHotword *int64 `json:"ReinforceHotword,omitnil,omitempty" name:"ReinforceHotword"`

	// 自学习定制模型 id
	// 如设置了该参数，将生效对应id的自学习定制模型；
	// 点击这里查看[自学习定制模型配置方法](https://cloud.tencent.com/document/product/1093/38416)
	CustomizationId *string `json:"CustomizationId,omitnil,omitempty" name:"CustomizationId"`

	// **【增值付费功能】**情绪识别能力（目前仅支持16k_zh,8k_zh）
	// 0：不开启；
	// 1：开启情绪识别，但不在文本展示情绪标签；
	// 2：开启情绪识别，并且在文本展示情绪标签（**该功能需要设置ResTextFormat 大于0**）
	// 默认值为0
	// 支持的情绪分类为：高兴、伤心、愤怒
	// 
	// 注意：
	// 1. **本功能为增值服务**，需将参数设置为1或2时方可按对应方式生效；
	// 2. 如果传入参数值1或2，需确保账号已购买[情绪识别资源包](https://cloud.tencent.com/document/product/1093/35686#97ae4aa0-29a0-4066-9f07-ccaf8856a16b)，或账号开启后付费；**若当前账号已开启后付费功能，并传入参数值1或2，将[自动计费](https://cloud.tencent.com/document/product/1093/35686#d912167d-ffd5-41a9-8b1c-2e89845a6852)）**；
	// 3. 参数设置为0时，无需购买资源包，也不会消耗情绪识别对应资源
	EmotionRecognition *int64 `json:"EmotionRecognition,omitnil,omitempty" name:"EmotionRecognition"`

	// 情绪能量值
	// 取值为音量分贝值/10，取值范围：[1,10]，值越高情绪越强烈
	// 0：不开启；
	// 1：开启；
	// 默认值为0
	EmotionalEnergy *int64 `json:"EmotionalEnergy,omitnil,omitempty" name:"EmotionalEnergy"`

	// 阿拉伯数字智能转换（目前支持中文普通话引擎）
	// 0：不转换，直接输出中文数字；
	// 1：根据场景智能转换为阿拉伯数字；
	// 3：打开数学相关数字转换（如：阿尔法转写为α）；
	// 默认值为 1
	ConvertNumMode *int64 `json:"ConvertNumMode,omitnil,omitempty" name:"ConvertNumMode"`

	// 脏词过滤（目前支持中文普通话引擎）
	// 0：不过滤脏词；
	// 1：过滤脏词；
	// 2：将脏词替换为 * ；
	// 默认值为 0
	FilterDirty *int64 `json:"FilterDirty,omitnil,omitempty" name:"FilterDirty"`

	// 标点符号过滤（目前支持中文普通话引擎）
	// 0：不过滤标点；
	// 1：过滤句末标点；
	// 2：过滤所有标点；
	// 默认值为 0
	FilterPunc *int64 `json:"FilterPunc,omitnil,omitempty" name:"FilterPunc"`

	// 语气词过滤（目前支持中文普通话引擎）
	// 0：不过滤语气词；
	// 1：过滤部分语气词；
	// 2：严格过滤语气词；
	// 默认值为 0
	FilterModal *int64 `json:"FilterModal,omitnil,omitempty" name:"FilterModal"`

	// 单标点最多字数（目前支持中文普通话引擎）
	// **可控制单行字幕最大字数，适用于字幕生成场景**，取值范围：[6，40]
	// 0：不开启该功能；
	// 默认值为0
	// 
	// 注意：需设置ResTextFormat为3，解析返回的ResultDetail列表，通过结构中FinalSentence获取单个标点断句结果
	SentenceMaxLength *int64 `json:"SentenceMaxLength,omitnil,omitempty" name:"SentenceMaxLength"`

	// 附加参数**（该参数无意义，忽略即可）**
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 临时热词表：该参数用于提升识别准确率。
	// 
	// - 单个热词限制："热词|权重"，单个热词不超过30个字符（最多10个汉字），权重[1-11]或者100，如：“腾讯云|5” 或“ASR|11”；
	// 
	// - 临时热词表限制：多个热词用英文逗号分割，最多支持128个热词，如：“腾讯云|10,语音识别|5,ASR|11”；
	// 
	// - 参数 hotword_id（热词表） 与 hotword_list（临时热词表） 区别：
	// 
	//     - hotword_id：热词表。需要先在控制台或接口创建热词表，获得对应hotword_id传入参数来使用热词功能；
	// 
	//     - hotword_list：临时热词表。每次请求时直接传入临时热词表来使用热词功能，云端不保留临时热词表。适用于有极大量热词需求的用户；
	// 
	// 注意：
	// 
	// - 如果同时传入了 hotword_id 和 hotword_list，只有hotword_list 生效；
	// 
	// - 热词权重设置为11时，当前热词将升级为超级热词，建议仅将重要且必须生效的热词设置到11，设置过多权重为11的热词将影响整体字准率。
	// 
	// - 热词权重设置为100时，当前热词开启热词增强同音替换功能（仅支持8k_zh,16k_zh），举例：热词配置“蜜制|100”时，与“蜜制”同拼音（mizhi）的“秘制”的识别结果会被强制替换成“蜜制”。因此建议客户根据自己的实际情况开启该功能。建议仅将重要且必须生效的热词设置到100，设置过多权重为100的热词将影响整体字准率。
	HotwordList *string `json:"HotwordList,omitnil,omitempty" name:"HotwordList"`

	// 关键词识别ID列表，默认空为不进行识别，最多10个
	KeyWordLibIdList []*string `json:"KeyWordLibIdList,omitnil,omitempty" name:"KeyWordLibIdList"`

	// 替换词汇表id,  适用于热词和自学习场景也无法解决的极端case词组,  会对识别结果强制替换。具体可参考[配置控制台](https://console.cloud.tencent.com/asr/replaceword);强制替换功能可能会影响正常识别结果，请谨慎使用
	// 
	// 注意：
	// 1. 本功能配置完成后，预计在10分钟后生效
	ReplaceTextId *string `json:"ReplaceTextId,omitnil,omitempty" name:"ReplaceTextId"`

	// 开启角色分离能力
	// 配合SpeakerDiarization: 3 使用，ASR增值服务，可传入一组声纹信息进行角色认证，仅支持16k_zh_en引擎。
	// 需传入SpeakerRoleInfo数据组，确定说话人的角色信息，涉及RoleAudioUrl和RoleName两个参数。 
	// RoleAudioUrl：需要认证角色的声纹音频地址，建议30s内的纯净人声，最长不能超过45s。 
	// RoleName：需要认证角色的名称，若匹配成功，会替换话者分离中的SpeakerID。 
	// 示例： 
	// "{\"EngineModelType\":\"16k_zh_en\",\"ChannelNum\":1,\"ResTextFormat\":1,\"SourceType\":0,\"Url\":\"需要进行ASR识别的音频链接\",\"SpeakerDiarization\":3,\"SpeakerRoles\":[{\"RoleAudioUrl\":\"需要认证角色的声纹音频地址\",\"RoleName\":\"需要认证角色的名称\"}]}"
	SpeakerRoles []*SpeakerRoleInfo `json:"SpeakerRoles,omitnil,omitempty" name:"SpeakerRoles"`
}

func (r *CreateRecTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineModelType")
	delete(f, "ChannelNum")
	delete(f, "ResTextFormat")
	delete(f, "SourceType")
	delete(f, "Data")
	delete(f, "DataLen")
	delete(f, "Url")
	delete(f, "CallbackUrl")
	delete(f, "SpeakerDiarization")
	delete(f, "SpeakerNumber")
	delete(f, "HotwordId")
	delete(f, "ReinforceHotword")
	delete(f, "CustomizationId")
	delete(f, "EmotionRecognition")
	delete(f, "EmotionalEnergy")
	delete(f, "ConvertNumMode")
	delete(f, "FilterDirty")
	delete(f, "FilterPunc")
	delete(f, "FilterModal")
	delete(f, "SentenceMaxLength")
	delete(f, "Extra")
	delete(f, "HotwordList")
	delete(f, "KeyWordLibIdList")
	delete(f, "ReplaceTextId")
	delete(f, "SpeakerRoles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecTaskResponseParams struct {
	// 录音文件识别的请求返回结果，包含结果查询需要的TaskId。
	// **注意：TaskId有效期为24小时，不同日期可能出现重复TaskId，请不要依赖TaskId作为您业务系统里的唯一ID。**
	Data *Task `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRecTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateRecTaskResponseParams `json:"Response"`
}

func (r *CreateRecTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAsrKeyWordLibRequestParams struct {
	// 关键词表ID
	KeyWordLibId *string `json:"KeyWordLibId,omitnil,omitempty" name:"KeyWordLibId"`
}

type DeleteAsrKeyWordLibRequest struct {
	*tchttp.BaseRequest
	
	// 关键词表ID
	KeyWordLibId *string `json:"KeyWordLibId,omitnil,omitempty" name:"KeyWordLibId"`
}

func (r *DeleteAsrKeyWordLibRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAsrKeyWordLibRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyWordLibId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAsrKeyWordLibRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAsrKeyWordLibResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAsrKeyWordLibResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAsrKeyWordLibResponseParams `json:"Response"`
}

func (r *DeleteAsrKeyWordLibResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAsrKeyWordLibResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAsrVocabRequestParams struct {
	// 热词表Id
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`
}

type DeleteAsrVocabRequest struct {
	*tchttp.BaseRequest
	
	// 热词表Id
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`
}

func (r *DeleteAsrVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAsrVocabRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VocabId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAsrVocabRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAsrVocabResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAsrVocabResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAsrVocabResponseParams `json:"Response"`
}

func (r *DeleteAsrVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAsrVocabResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomizationRequestParams struct {
	// 要删除的模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

type DeleteCustomizationRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

func (r *DeleteCustomizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomizationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCustomizationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomizationResponseParams `json:"Response"`
}

func (r *DeleteCustomizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRecognitionTasksRequestParams struct {

}

type DescribeAsyncRecognitionTasksRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAsyncRecognitionTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRecognitionTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncRecognitionTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRecognitionTasksResponseParams struct {
	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AsyncRecognitionTasks `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAsyncRecognitionTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAsyncRecognitionTasksResponseParams `json:"Response"`
}

func (r *DescribeAsyncRecognitionTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRecognitionTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// 从CreateRecTask接口获取的TaskId，用于获取任务状态与结果。
	// **注意：TaskId有效期为24小时，超过24小时的TaskId请不要再查询。**
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 从CreateRecTask接口获取的TaskId，用于获取任务状态与结果。
	// **注意：TaskId有效期为24小时，超过24小时的TaskId请不要再查询。**
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusResponseParams struct {
	// 录音文件识别的请求返回结果。
	Data *TaskStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadAsrVocabRequestParams struct {
	// 词表ID。
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`
}

type DownloadAsrVocabRequest struct {
	*tchttp.BaseRequest
	
	// 词表ID。
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`
}

func (r *DownloadAsrVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadAsrVocabRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VocabId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadAsrVocabRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadAsrVocabResponseParams struct {
	// 词表ID。
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`

	// 词表权重文件形式的base64值。
	WordWeightStr *string `json:"WordWeightStr,omitnil,omitempty" name:"WordWeightStr"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DownloadAsrVocabResponse struct {
	*tchttp.BaseResponse
	Response *DownloadAsrVocabResponseParams `json:"Response"`
}

func (r *DownloadAsrVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadAsrVocabResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCustomizationRequestParams struct {
	// 自学习模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

type DownloadCustomizationRequest struct {
	*tchttp.BaseRequest
	
	// 自学习模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

func (r *DownloadCustomizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCustomizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadCustomizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCustomizationResponseParams struct {
	// 下载地址
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DownloadCustomizationResponse struct {
	*tchttp.BaseResponse
	Response *DownloadCustomizationResponseParams `json:"Response"`
}

func (r *DownloadCustomizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCustomizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAsrKeyWordLibListRequestParams struct {
	// 分页Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 词库名称或者UIN检索
	SpecifyNames []*string `json:"SpecifyNames,omitnil,omitempty" name:"SpecifyNames"`

	// 只看用户自己创建的
	OnlySelf *bool `json:"OnlySelf,omitnil,omitempty" name:"OnlySelf"`
}

type GetAsrKeyWordLibListRequest struct {
	*tchttp.BaseRequest
	
	// 分页Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 词库名称或者UIN检索
	SpecifyNames []*string `json:"SpecifyNames,omitnil,omitempty" name:"SpecifyNames"`

	// 只看用户自己创建的
	OnlySelf *bool `json:"OnlySelf,omitnil,omitempty" name:"OnlySelf"`
}

func (r *GetAsrKeyWordLibListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAsrKeyWordLibListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SpecifyNames")
	delete(f, "OnlySelf")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAsrKeyWordLibListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAsrKeyWordLibListResponseParams struct {
	// 关键词列表返回数据
	Data *KeyWordLibListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAsrKeyWordLibListResponse struct {
	*tchttp.BaseResponse
	Response *GetAsrKeyWordLibListResponseParams `json:"Response"`
}

func (r *GetAsrKeyWordLibListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAsrKeyWordLibListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAsrVocabListRequestParams struct {
	// 标签信息，格式为“$TagKey : $TagValue ”，中间分隔符为“空格”+“:”+“空格”
	//
	// Deprecated: TagInfos is deprecated.
	TagInfos []*string `json:"TagInfos,omitnil,omitempty" name:"TagInfos"`

	// 分页Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetAsrVocabListRequest struct {
	*tchttp.BaseRequest
	
	// 标签信息，格式为“$TagKey : $TagValue ”，中间分隔符为“空格”+“:”+“空格”
	TagInfos []*string `json:"TagInfos,omitnil,omitempty" name:"TagInfos"`

	// 分页Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *GetAsrVocabListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAsrVocabListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagInfos")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAsrVocabListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAsrVocabListResponseParams struct {
	// 热词表列表
	VocabList []*Vocab `json:"VocabList,omitnil,omitempty" name:"VocabList"`

	// 热词列表总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAsrVocabListResponse struct {
	*tchttp.BaseResponse
	Response *GetAsrVocabListResponseParams `json:"Response"`
}

func (r *GetAsrVocabListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAsrVocabListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAsrVocabRequestParams struct {
	// 热词表ID
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`
}

type GetAsrVocabRequest struct {
	*tchttp.BaseRequest
	
	// 热词表ID
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`
}

func (r *GetAsrVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAsrVocabRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VocabId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAsrVocabRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAsrVocabResponseParams struct {
	// 热词表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 热词表描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 热词表ID
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`

	// 词权重列表
	WordWeights []*HotWord `json:"WordWeights,omitnil,omitempty" name:"WordWeights"`

	// 词表创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 词表更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 热词表状态，1为默认状态即在识别时默认加载该热词表进行识别，0为初始状态
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAsrVocabResponse struct {
	*tchttp.BaseResponse
	Response *GetAsrVocabResponseParams `json:"Response"`
}

func (r *GetAsrVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAsrVocabResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCustomizationListRequestParams struct {
	// 标签信息，格式为“$TagKey : $TagValue ”，中间分隔符为“空格”+“:”+“空格”
	//
	// Deprecated: TagInfos is deprecated.
	TagInfos []*string `json:"TagInfos,omitnil,omitempty" name:"TagInfos"`

	// 分页大小，默认1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页offset，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type GetCustomizationListRequest struct {
	*tchttp.BaseRequest
	
	// 标签信息，格式为“$TagKey : $TagValue ”，中间分隔符为“空格”+“:”+“空格”
	TagInfos []*string `json:"TagInfos,omitnil,omitempty" name:"TagInfos"`

	// 分页大小，默认1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页offset，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *GetCustomizationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCustomizationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagInfos")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCustomizationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCustomizationListResponseParams struct {
	// 自学习模型数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*Model `json:"Data,omitnil,omitempty" name:"Data"`

	// 自学习模型总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCustomizationListResponse struct {
	*tchttp.BaseResponse
	Response *GetCustomizationListResponseParams `json:"Response"`
}

func (r *GetCustomizationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCustomizationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetModelInfoRequestParams struct {
	// 模型id
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

type GetModelInfoRequest struct {
	*tchttp.BaseRequest
	
	// 模型id
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

func (r *GetModelInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetModelInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetModelInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetModelInfoResponseParams struct {
	// 模型信息
	Data *Model `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetModelInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetModelInfoResponseParams `json:"Response"`
}

func (r *GetModelInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetModelInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUsageByDateRequestParams struct {
	// 需要查询的业务类型名字列表
	// - asr_rt 实时识别
	// - asr_rec 录音文件识别
	BizNameList []*string `json:"BizNameList,omitnil,omitempty" name:"BizNameList"`

	// 查询开始时间
	// 开始时间包含当天，支持 YYYY-MM-DD 日期以国内时区为准
	// 开始时间到结束时间需要在3个月以内
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 查询结束时间
	// 结束时间包含当天，，支持 YYYY-MM-DD 日期以国内时区为准
	// 开始时间到结束时间需要在3个月以内
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

type GetUsageByDateRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的业务类型名字列表
	// - asr_rt 实时识别
	// - asr_rec 录音文件识别
	BizNameList []*string `json:"BizNameList,omitnil,omitempty" name:"BizNameList"`

	// 查询开始时间
	// 开始时间包含当天，支持 YYYY-MM-DD 日期以国内时区为准
	// 开始时间到结束时间需要在3个月以内
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 查询结束时间
	// 结束时间包含当天，，支持 YYYY-MM-DD 日期以国内时区为准
	// 开始时间到结束时间需要在3个月以内
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

func (r *GetUsageByDateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUsageByDateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizNameList")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUsageByDateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUsageByDateResponseParams struct {
	// 用量次数
	Data *UsageByDateInfoData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetUsageByDateResponse struct {
	*tchttp.BaseResponse
	Response *GetUsageByDateResponseParams `json:"Response"`
}

func (r *GetUsageByDateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUsageByDateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HotWord struct {
	// 热词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type KeyWordLib struct {
	// 关键词表ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyWordLibId *string `json:"KeyWordLibId,omitnil,omitempty" name:"KeyWordLibId"`

	// 关键词表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 关键词列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyWordList []*string `json:"KeyWordList,omitnil,omitempty" name:"KeyWordList"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type KeyWordLibIdData struct {
	// 关键词ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyWordLibId *string `json:"KeyWordLibId,omitnil,omitempty" name:"KeyWordLibId"`
}

type KeyWordLibListData struct {
	// 关键词表列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyWordLibList []*KeyWordLib `json:"KeyWordLibList,omitnil,omitempty" name:"KeyWordLibList"`

	// 关键词列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type KeyWordResult struct {
	// 关键词库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyWordLibID *string `json:"KeyWordLibID,omitnil,omitempty" name:"KeyWordLibID"`

	// 关键词库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyWordLibName *string `json:"KeyWordLibName,omitnil,omitempty" name:"KeyWordLibName"`

	// 匹配到的关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`
}

type Model struct {
	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 模型文件名称
	DictName *string `json:"DictName,omitnil,omitempty" name:"DictName"`

	// 模型Id
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型类型，“8k”或者”16k“
	ModelType *string `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// 服务类型
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 模型状态：
	// -2：模型训练失败；
	// -1：已下线；
	// 0：训练中；
	// 1：已上线；
	// 3：上线中；
	// 4：下线中；
	ModelState *int64 `json:"ModelState,omitnil,omitempty" name:"ModelState"`

	// 最后更新时间
	AtUpdated *string `json:"AtUpdated,omitnil,omitempty" name:"AtUpdated"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: TagInfos is deprecated.
	TagInfos []*string `json:"TagInfos,omitnil,omitempty" name:"TagInfos"`
}

// Predefined struct for user
type ModifyCustomizationRequestParams struct {
	// 要修改的模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 要修改的模型名称，长度需在1-20个字符之间
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 要修改的模型类型，为8k或者16k
	ModelType *string `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// 要修改的模型语料的下载地址，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitnil,omitempty" name:"TextUrl"`
}

type ModifyCustomizationRequest struct {
	*tchttp.BaseRequest
	
	// 要修改的模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 要修改的模型名称，长度需在1-20个字符之间
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 要修改的模型类型，为8k或者16k
	ModelType *string `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// 要修改的模型语料的下载地址，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitnil,omitempty" name:"TextUrl"`
}

func (r *ModifyCustomizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "ModelName")
	delete(f, "ModelType")
	delete(f, "TextUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCustomizationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomizationResponseParams `json:"Response"`
}

func (r *ModifyCustomizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizationStateRequestParams struct {
	// 自学习模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 想要变换的模型状态，-1代表下线，1代表上线
	ToState *int64 `json:"ToState,omitnil,omitempty" name:"ToState"`
}

type ModifyCustomizationStateRequest struct {
	*tchttp.BaseRequest
	
	// 自学习模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 想要变换的模型状态，-1代表下线，1代表上线
	ToState *int64 `json:"ToState,omitnil,omitempty" name:"ToState"`
}

func (r *ModifyCustomizationStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizationStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "ToState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomizationStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizationStateResponseParams struct {
	// 自学习模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCustomizationStateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomizationStateResponseParams `json:"Response"`
}

func (r *ModifyCustomizationStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizationStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SentenceDetail struct {
	// 单句最终识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalSentence *string `json:"FinalSentence,omitnil,omitempty" name:"FinalSentence"`

	// 单句中间识别结果，使用空格拆分为多个词
	// 注意：此字段可能返回 null，表示取不到有效值。
	SliceSentence *string `json:"SliceSentence,omitnil,omitempty" name:"SliceSentence"`

	// 口语转书面语结果，开启改功能才有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	WrittenText *string `json:"WrittenText,omitnil,omitempty" name:"WrittenText"`

	// 单句开始时间（毫秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartMs *int64 `json:"StartMs,omitnil,omitempty" name:"StartMs"`

	// 单句结束时间（毫秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndMs *int64 `json:"EndMs,omitnil,omitempty" name:"EndMs"`

	// 单句中词个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordsNum *int64 `json:"WordsNum,omitnil,omitempty" name:"WordsNum"`

	// 单句中词详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Words []*SentenceWords `json:"Words,omitnil,omitempty" name:"Words"`

	// 单句语速，单位：字数/秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpeechSpeed *float64 `json:"SpeechSpeed,omitnil,omitempty" name:"SpeechSpeed"`

	// 声道或说话人 Id（请求中如果设置了 speaker_diarization或者ChannelNum为双声道，可区分说话人或声道）
	// 单声道话者分离时不同的值代表不同的说话人； 8k双声道话者分离时speakerId的值为0代表左声道，值为1代表右声道。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpeakerId *int64 `json:"SpeakerId,omitnil,omitempty" name:"SpeakerId"`

	// 情绪能量值，取值为音量分贝值/10。取值范围：[1,10]。值越高情绪越强烈。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmotionalEnergy *float64 `json:"EmotionalEnergy,omitnil,omitempty" name:"EmotionalEnergy"`

	// 本句与上一句之间的静音时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	SilenceTime *int64 `json:"SilenceTime,omitnil,omitempty" name:"SilenceTime"`

	// 情绪类型（可能为空，有2种情况 1、没有对应资源包；2、情绪跟语音效果相关，如果情绪不够强烈时可能无法识别）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmotionType []*string `json:"EmotionType,omitnil,omitempty" name:"EmotionType"`

	// 关键词识别结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyWordResults []*KeyWordResult `json:"KeyWordResults,omitnil,omitempty" name:"KeyWordResults"`
}

// Predefined struct for user
type SentenceRecognitionRequestParams struct {
	// 引擎模型类型。
	// 电话场景：
	// • 8k_zh：中文电话通用；
	// • 8k_en：英文电话通用；
	// 
	// 非电话场景：
	// • 16k_zh：中文通用；
	// • 16k_zh-PY：中英粤;
	// • 16k_zh_medical：中文医疗；
	// • 16k_en：英语；
	// • 16k_yue：粤语；
	// • 16k_ja：日语；
	// • 16k_ko：韩语；
	// • 16k_vi：越南语；
	// • 16k_ms：马来语；
	// • 16k_id：印度尼西亚语；
	// • 16k_fil：菲律宾语；
	// • 16k_th：泰语；
	// • 16k_pt：葡萄牙语；
	// • 16k_tr：土耳其语；
	// • 16k_ar：阿拉伯语；
	// • 16k_es：西班牙语；
	// • 16k_hi：印地语；
	// • 16k_fr：法语；
	// • 16k_de：德语；
	// • 16k_zh_dialect：多方言，支持23种方言（上海话、四川话、武汉话、贵阳话、昆明话、西安话、郑州话、太原话、兰州话、银川话、西宁话、南京话、合肥话、南昌话、长沙话、苏州话、杭州话、济南话、天津话、石家庄话、黑龙江话、吉林话、辽宁话）；
	EngSerViceType *string `json:"EngSerViceType,omitnil,omitempty" name:"EngSerViceType"`

	// 语音数据来源。0：语音 URL；1：语音数据（post body）。
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 识别音频的音频格式，支持wav、pcm、ogg-opus、speex、silk、mp3、m4a、aac、amr。
	VoiceFormat *string `json:"VoiceFormat,omitnil,omitempty" name:"VoiceFormat"`

	// 腾讯云项目 ID，废弃参数，填写0即可。
	//
	// Deprecated: ProjectId is deprecated.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 子服务类型。2： 一句话识别。
	//
	// Deprecated: SubServiceType is deprecated.
	SubServiceType *uint64 `json:"SubServiceType,omitnil,omitempty" name:"SubServiceType"`

	// 语音的URL地址，需要公网环境浏览器可下载。当 SourceType 值为 0时须填写该字段，为 1 时不填。音频时长不能超过60s，音频文件大小不能超过3MB。
	// 
	// 注意：推荐使用 [腾讯云COS](https://cloud.tencent.com/document/product/436/38484) 来存储音频、生成URL并提交请求，此种方式会走内网下载音频，极大降低整体请求时延；并且不会产生外网和流量下行费用，可节约成本（COS桶权限需要设置公有读私有写，或URL设置时效访问签名）
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 废弃参数，填写任意字符串即可。
	//
	// Deprecated: UsrAudioKey is deprecated.
	UsrAudioKey *string `json:"UsrAudioKey,omitnil,omitempty" name:"UsrAudioKey"`

	// 语音数据，当SourceType 值为1（本地语音数据上传）时必须填写，当SourceType 值为0（语音 URL上传）可不写。要使用base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。音频时长不能超过60s，音频文件大小不能超过3MB（Base64后）。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 数据长度，单位为字节。当 SourceType 值为1（本地语音数据上传）时必须填写，当 SourceType 值为0（语音 URL上传）可不写（此数据长度为数据未进行base64编码时的数据长度）。
	DataLen *int64 `json:"DataLen,omitnil,omitempty" name:"DataLen"`

	// 是否显示词级别时间戳。0：不显示；1：显示，不包含标点时间戳，2：显示，包含标点时间戳。默认值为 0。
	WordInfo *int64 `json:"WordInfo,omitnil,omitempty" name:"WordInfo"`

	// 是否过滤脏词（目前支持中文普通话引擎）。0：不过滤脏词；1：过滤脏词；2：将脏词替换为 * 。默认值为 0。
	FilterDirty *int64 `json:"FilterDirty,omitnil,omitempty" name:"FilterDirty"`

	// 是否过滤语气词（目前支持中文普通话引擎）。0：不过滤语气词；1：部分过滤；2：严格过滤 。默认值为 0。
	FilterModal *int64 `json:"FilterModal,omitnil,omitempty" name:"FilterModal"`

	// 是否过滤标点符号（目前支持中文普通话引擎）。 0：不过滤，1：过滤句末标点，2：过滤所有标点。默认值为 0。
	FilterPunc *int64 `json:"FilterPunc,omitnil,omitempty" name:"FilterPunc"`

	// 是否进行阿拉伯数字智能转换。0：不转换，直接输出中文数字，1：根据场景智能转换为阿拉伯数字。默认值为1。
	ConvertNumMode *int64 `json:"ConvertNumMode,omitnil,omitempty" name:"ConvertNumMode"`

	// 热词id。用于调用对应的热词表，如果在调用语音识别服务时，不进行单独的热词id设置，自动生效默认热词；如果进行了单独的热词id设置，那么将生效单独设置的热词id。
	HotwordId *string `json:"HotwordId,omitnil,omitempty" name:"HotwordId"`

	// 自学习模型 id。如设置了该参数，将生效对应的自学习模型。
	CustomizationId *string `json:"CustomizationId,omitnil,omitempty" name:"CustomizationId"`

	// 热词增强功能。1:开启后（仅支持8k_zh,16k_zh），将开启同音替换功能，同音字、词在热词中配置。举例：热词配置“蜜制”并开启增强功能后，与“蜜制”同拼音（mizhi）的“秘制”的识别结果会被强制替换成“蜜制”。因此建议客户根据自己的实际情况开启该功能。
	//
	// Deprecated: ReinforceHotword is deprecated.
	ReinforceHotword *int64 `json:"ReinforceHotword,omitnil,omitempty" name:"ReinforceHotword"`

	// 临时热词表：该参数用于提升识别准确率。
	// 单个热词限制："热词|权重"，单个热词不超过30个字符（最多10个汉字），权重1-11或者100，如：“腾讯云|5” 或 “ASR|11”；
	// 临时热词表限制：多个热词用英文逗号分割，最多支持128个热词，如：“腾讯云|10,语音识别|5,ASR|11”；
	// 参数 hotword_list（临时热词表） 与 hotword_id（热词表） 区别：
	// hotword_id：热词表。需要先在控制台或接口创建热词表，获得对应hotword_id传入参数来使用热词功能；
	// hotword_list：临时热词表。每次请求时直接传入临时热词表来使用热词功能，云端不保留临时热词表。适用于有极大量热词需求的用户；
	// 注意：
	// • 如果同时传入了 hotword_id 和 hotword_list，会优先使用 hotword_list；
	// • 热词权重设置为11时，当前热词将升级为超级热词，建议仅将重要且必须生效的热词设置到11，设置过多权重为11的热词将影响整体字准率。
	// • 热词权重设置为100时，当前热词开启热词增强同音替换功能（仅支持8k_zh,16k_zh），举例：热词配置“蜜制|100”时，与“蜜制”同拼音（mizhi）的“秘制”的识别结果会被强制替换成“蜜制”。因此建议客户根据自己的实际情况开启该功能。建议仅将重要且必须生效的热词设置到100，设置过多权重为100的热词将影响整体字准率。
	HotwordList *string `json:"HotwordList,omitnil,omitempty" name:"HotwordList"`

	// 支持pcm格式的8k音频在与引擎采样率不匹配的情况下升采样到16k后识别，能有效提升识别准确率。仅支持：8000。如：传入 8000 ，则pcm音频采样率为8k，当引擎选用16k_zh， 那么该8k采样率的pcm音频可以在16k_zh引擎下正常识别。 注：此参数仅适用于pcm格式音频，不传入值将维持默认状态，即默认调用的引擎采样率等于pcm音频采样率。
	InputSampleRate *int64 `json:"InputSampleRate,omitnil,omitempty" name:"InputSampleRate"`
}

type SentenceRecognitionRequest struct {
	*tchttp.BaseRequest
	
	// 引擎模型类型。
	// 电话场景：
	// • 8k_zh：中文电话通用；
	// • 8k_en：英文电话通用；
	// 
	// 非电话场景：
	// • 16k_zh：中文通用；
	// • 16k_zh-PY：中英粤;
	// • 16k_zh_medical：中文医疗；
	// • 16k_en：英语；
	// • 16k_yue：粤语；
	// • 16k_ja：日语；
	// • 16k_ko：韩语；
	// • 16k_vi：越南语；
	// • 16k_ms：马来语；
	// • 16k_id：印度尼西亚语；
	// • 16k_fil：菲律宾语；
	// • 16k_th：泰语；
	// • 16k_pt：葡萄牙语；
	// • 16k_tr：土耳其语；
	// • 16k_ar：阿拉伯语；
	// • 16k_es：西班牙语；
	// • 16k_hi：印地语；
	// • 16k_fr：法语；
	// • 16k_de：德语；
	// • 16k_zh_dialect：多方言，支持23种方言（上海话、四川话、武汉话、贵阳话、昆明话、西安话、郑州话、太原话、兰州话、银川话、西宁话、南京话、合肥话、南昌话、长沙话、苏州话、杭州话、济南话、天津话、石家庄话、黑龙江话、吉林话、辽宁话）；
	EngSerViceType *string `json:"EngSerViceType,omitnil,omitempty" name:"EngSerViceType"`

	// 语音数据来源。0：语音 URL；1：语音数据（post body）。
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 识别音频的音频格式，支持wav、pcm、ogg-opus、speex、silk、mp3、m4a、aac、amr。
	VoiceFormat *string `json:"VoiceFormat,omitnil,omitempty" name:"VoiceFormat"`

	// 腾讯云项目 ID，废弃参数，填写0即可。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 子服务类型。2： 一句话识别。
	SubServiceType *uint64 `json:"SubServiceType,omitnil,omitempty" name:"SubServiceType"`

	// 语音的URL地址，需要公网环境浏览器可下载。当 SourceType 值为 0时须填写该字段，为 1 时不填。音频时长不能超过60s，音频文件大小不能超过3MB。
	// 
	// 注意：推荐使用 [腾讯云COS](https://cloud.tencent.com/document/product/436/38484) 来存储音频、生成URL并提交请求，此种方式会走内网下载音频，极大降低整体请求时延；并且不会产生外网和流量下行费用，可节约成本（COS桶权限需要设置公有读私有写，或URL设置时效访问签名）
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 废弃参数，填写任意字符串即可。
	UsrAudioKey *string `json:"UsrAudioKey,omitnil,omitempty" name:"UsrAudioKey"`

	// 语音数据，当SourceType 值为1（本地语音数据上传）时必须填写，当SourceType 值为0（语音 URL上传）可不写。要使用base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。音频时长不能超过60s，音频文件大小不能超过3MB（Base64后）。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 数据长度，单位为字节。当 SourceType 值为1（本地语音数据上传）时必须填写，当 SourceType 值为0（语音 URL上传）可不写（此数据长度为数据未进行base64编码时的数据长度）。
	DataLen *int64 `json:"DataLen,omitnil,omitempty" name:"DataLen"`

	// 是否显示词级别时间戳。0：不显示；1：显示，不包含标点时间戳，2：显示，包含标点时间戳。默认值为 0。
	WordInfo *int64 `json:"WordInfo,omitnil,omitempty" name:"WordInfo"`

	// 是否过滤脏词（目前支持中文普通话引擎）。0：不过滤脏词；1：过滤脏词；2：将脏词替换为 * 。默认值为 0。
	FilterDirty *int64 `json:"FilterDirty,omitnil,omitempty" name:"FilterDirty"`

	// 是否过滤语气词（目前支持中文普通话引擎）。0：不过滤语气词；1：部分过滤；2：严格过滤 。默认值为 0。
	FilterModal *int64 `json:"FilterModal,omitnil,omitempty" name:"FilterModal"`

	// 是否过滤标点符号（目前支持中文普通话引擎）。 0：不过滤，1：过滤句末标点，2：过滤所有标点。默认值为 0。
	FilterPunc *int64 `json:"FilterPunc,omitnil,omitempty" name:"FilterPunc"`

	// 是否进行阿拉伯数字智能转换。0：不转换，直接输出中文数字，1：根据场景智能转换为阿拉伯数字。默认值为1。
	ConvertNumMode *int64 `json:"ConvertNumMode,omitnil,omitempty" name:"ConvertNumMode"`

	// 热词id。用于调用对应的热词表，如果在调用语音识别服务时，不进行单独的热词id设置，自动生效默认热词；如果进行了单独的热词id设置，那么将生效单独设置的热词id。
	HotwordId *string `json:"HotwordId,omitnil,omitempty" name:"HotwordId"`

	// 自学习模型 id。如设置了该参数，将生效对应的自学习模型。
	CustomizationId *string `json:"CustomizationId,omitnil,omitempty" name:"CustomizationId"`

	// 热词增强功能。1:开启后（仅支持8k_zh,16k_zh），将开启同音替换功能，同音字、词在热词中配置。举例：热词配置“蜜制”并开启增强功能后，与“蜜制”同拼音（mizhi）的“秘制”的识别结果会被强制替换成“蜜制”。因此建议客户根据自己的实际情况开启该功能。
	ReinforceHotword *int64 `json:"ReinforceHotword,omitnil,omitempty" name:"ReinforceHotword"`

	// 临时热词表：该参数用于提升识别准确率。
	// 单个热词限制："热词|权重"，单个热词不超过30个字符（最多10个汉字），权重1-11或者100，如：“腾讯云|5” 或 “ASR|11”；
	// 临时热词表限制：多个热词用英文逗号分割，最多支持128个热词，如：“腾讯云|10,语音识别|5,ASR|11”；
	// 参数 hotword_list（临时热词表） 与 hotword_id（热词表） 区别：
	// hotword_id：热词表。需要先在控制台或接口创建热词表，获得对应hotword_id传入参数来使用热词功能；
	// hotword_list：临时热词表。每次请求时直接传入临时热词表来使用热词功能，云端不保留临时热词表。适用于有极大量热词需求的用户；
	// 注意：
	// • 如果同时传入了 hotword_id 和 hotword_list，会优先使用 hotword_list；
	// • 热词权重设置为11时，当前热词将升级为超级热词，建议仅将重要且必须生效的热词设置到11，设置过多权重为11的热词将影响整体字准率。
	// • 热词权重设置为100时，当前热词开启热词增强同音替换功能（仅支持8k_zh,16k_zh），举例：热词配置“蜜制|100”时，与“蜜制”同拼音（mizhi）的“秘制”的识别结果会被强制替换成“蜜制”。因此建议客户根据自己的实际情况开启该功能。建议仅将重要且必须生效的热词设置到100，设置过多权重为100的热词将影响整体字准率。
	HotwordList *string `json:"HotwordList,omitnil,omitempty" name:"HotwordList"`

	// 支持pcm格式的8k音频在与引擎采样率不匹配的情况下升采样到16k后识别，能有效提升识别准确率。仅支持：8000。如：传入 8000 ，则pcm音频采样率为8k，当引擎选用16k_zh， 那么该8k采样率的pcm音频可以在16k_zh引擎下正常识别。 注：此参数仅适用于pcm格式音频，不传入值将维持默认状态，即默认调用的引擎采样率等于pcm音频采样率。
	InputSampleRate *int64 `json:"InputSampleRate,omitnil,omitempty" name:"InputSampleRate"`
}

func (r *SentenceRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SentenceRecognitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngSerViceType")
	delete(f, "SourceType")
	delete(f, "VoiceFormat")
	delete(f, "ProjectId")
	delete(f, "SubServiceType")
	delete(f, "Url")
	delete(f, "UsrAudioKey")
	delete(f, "Data")
	delete(f, "DataLen")
	delete(f, "WordInfo")
	delete(f, "FilterDirty")
	delete(f, "FilterModal")
	delete(f, "FilterPunc")
	delete(f, "ConvertNumMode")
	delete(f, "HotwordId")
	delete(f, "CustomizationId")
	delete(f, "ReinforceHotword")
	delete(f, "HotwordList")
	delete(f, "InputSampleRate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SentenceRecognitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SentenceRecognitionResponseParams struct {
	// 识别结果。
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 请求的音频时长，单位为ms
	AudioDuration *int64 `json:"AudioDuration,omitnil,omitempty" name:"AudioDuration"`

	// 词时间戳列表的长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordSize *int64 `json:"WordSize,omitnil,omitempty" name:"WordSize"`

	// 词时间戳列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordList []*SentenceWord `json:"WordList,omitnil,omitempty" name:"WordList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SentenceRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *SentenceRecognitionResponseParams `json:"Response"`
}

func (r *SentenceRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SentenceRecognitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SentenceWord struct {
	// 词结果
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 词在音频中的开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 词在音频中的结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type SentenceWords struct {
	// 词文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 在句子中的开始时间偏移量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetStartMs *int64 `json:"OffsetStartMs,omitnil,omitempty" name:"OffsetStartMs"`

	// 在句子中的结束时间偏移量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetEndMs *int64 `json:"OffsetEndMs,omitnil,omitempty" name:"OffsetEndMs"`
}

// Predefined struct for user
type SetVocabStateRequestParams struct {
	// 热词表ID。
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`

	// 热词表状态，1：设为默认状态；0：设为非默认状态。
	State *int64 `json:"State,omitnil,omitempty" name:"State"`
}

type SetVocabStateRequest struct {
	*tchttp.BaseRequest
	
	// 热词表ID。
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`

	// 热词表状态，1：设为默认状态；0：设为非默认状态。
	State *int64 `json:"State,omitnil,omitempty" name:"State"`
}

func (r *SetVocabStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVocabStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VocabId")
	delete(f, "State")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetVocabStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVocabStateResponseParams struct {
	// 热词表ID
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetVocabStateResponse struct {
	*tchttp.BaseResponse
	Response *SetVocabStateResponseParams `json:"Response"`
}

func (r *SetVocabStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVocabStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpeakerRoleInfo struct {
	// 音频url地址，建议不超过30秒，最大45秒
	RoleAudioUrl *string `json:"RoleAudioUrl,omitnil,omitempty" name:"RoleAudioUrl"`

	// 不超过30字节
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type Task struct {
	// 任务ID，可通过此ID在轮询接口获取识别状态与结果。TaskId数据类型为**uint64**。
	// **注意：TaskId有效期为24小时，不同日期可能出现重复TaskId，请不要依赖TaskId作为您业务系统里的唯一ID。**
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type TaskStatus struct {
	// 任务标识。注意：TaskId数据类型为uint64。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态码，0：任务等待，1：任务执行中，2：任务成功，3：任务失败。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务状态，waiting：任务等待，doing：任务执行中，success：任务成功，failed：任务失败。
	StatusStr *string `json:"StatusStr,omitnil,omitempty" name:"StatusStr"`

	// 识别结果。
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 失败原因说明。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 识别结果详情，包含每个句子中的词时间偏移，一般用于生成字幕的场景。(录音识别请求中ResTextFormat=1时该字段不为空)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultDetail []*SentenceDetail `json:"ResultDetail,omitnil,omitempty" name:"ResultDetail"`

	// 音频时长(秒)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioDuration *float64 `json:"AudioDuration,omitnil,omitempty" name:"AudioDuration"`
}

// Predefined struct for user
type UpdateAsrKeyWordLibRequestParams struct {
	// 关键词表ID
	KeyWordLibId *string `json:"KeyWordLibId,omitnil,omitempty" name:"KeyWordLibId"`

	// 词表名称，长度在1-20之间
	// 仅限中英文数字-_
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// - 词文件（纯文本文件）以行分隔 ，进行二进制base64编码
	// - 格式要求：TXT 每行只有一个词，不满足格式则报错无法上传 
	// - 每个词最多5个汉字或15个字符，单个词库最多不超过100个词
	// - 此参数为空则只更新词表名称
	KeyWordFile *string `json:"KeyWordFile,omitnil,omitempty" name:"KeyWordFile"`
}

type UpdateAsrKeyWordLibRequest struct {
	*tchttp.BaseRequest
	
	// 关键词表ID
	KeyWordLibId *string `json:"KeyWordLibId,omitnil,omitempty" name:"KeyWordLibId"`

	// 词表名称，长度在1-20之间
	// 仅限中英文数字-_
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// - 词文件（纯文本文件）以行分隔 ，进行二进制base64编码
	// - 格式要求：TXT 每行只有一个词，不满足格式则报错无法上传 
	// - 每个词最多5个汉字或15个字符，单个词库最多不超过100个词
	// - 此参数为空则只更新词表名称
	KeyWordFile *string `json:"KeyWordFile,omitnil,omitempty" name:"KeyWordFile"`
}

func (r *UpdateAsrKeyWordLibRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAsrKeyWordLibRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyWordLibId")
	delete(f, "Name")
	delete(f, "KeyWordFile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAsrKeyWordLibRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAsrKeyWordLibResponseParams struct {
	// 关键词表ID数据
	Data *KeyWordLibIdData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAsrKeyWordLibResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAsrKeyWordLibResponseParams `json:"Response"`
}

func (r *UpdateAsrKeyWordLibResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAsrKeyWordLibResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAsrVocabRequestParams struct {
	// 热词表ID
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`

	// 热词表名称，长度在1-255之间
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 词权重数组，包含全部的热词和对应的权重。每个热词的长度不大于10个汉字或30个英文字符，权重为[1,11]之间整数或100，数组长度不大于1000 (注意：如果仅更新热词表名称或者描述字段，请求不用带本参数）
	WordWeights []*HotWord `json:"WordWeights,omitnil,omitempty" name:"WordWeights"`

	// 词权重文件（纯文本文件）的二进制base64编码，以行分隔，每行的格式为word|weight，即以英文符号|为分割，左边为词，右边为权重，如：你好|5。
	// 当用户传此参数（参数长度大于0），即以此参数解析词权重，WordWeights会被忽略
	//  (注意：如果仅更新热词表名称或者描述字段，请求不用带本参数）
	WordWeightStr *string `json:"WordWeightStr,omitnil,omitempty" name:"WordWeightStr"`

	// 热词表描述，长度在0-1000之间
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateAsrVocabRequest struct {
	*tchttp.BaseRequest
	
	// 热词表ID
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`

	// 热词表名称，长度在1-255之间
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 词权重数组，包含全部的热词和对应的权重。每个热词的长度不大于10个汉字或30个英文字符，权重为[1,11]之间整数或100，数组长度不大于1000 (注意：如果仅更新热词表名称或者描述字段，请求不用带本参数）
	WordWeights []*HotWord `json:"WordWeights,omitnil,omitempty" name:"WordWeights"`

	// 词权重文件（纯文本文件）的二进制base64编码，以行分隔，每行的格式为word|weight，即以英文符号|为分割，左边为词，右边为权重，如：你好|5。
	// 当用户传此参数（参数长度大于0），即以此参数解析词权重，WordWeights会被忽略
	//  (注意：如果仅更新热词表名称或者描述字段，请求不用带本参数）
	WordWeightStr *string `json:"WordWeightStr,omitnil,omitempty" name:"WordWeightStr"`

	// 热词表描述，长度在0-1000之间
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateAsrVocabRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAsrVocabRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VocabId")
	delete(f, "Name")
	delete(f, "WordWeights")
	delete(f, "WordWeightStr")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAsrVocabRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAsrVocabResponseParams struct {
	// 热词表ID
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAsrVocabResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAsrVocabResponseParams `json:"Response"`
}

func (r *UpdateAsrVocabResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAsrVocabResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsageByDateInfo struct {
	// 业务类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizName *string `json:"BizName,omitnil,omitempty" name:"BizName"`

	// 识别次数
	// 单位：次
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 识别时长
	// 单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

type UsageByDateInfoData struct {
	// 用量信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsageByDateInfoList []*UsageByDateInfo `json:"UsageByDateInfoList,omitnil,omitempty" name:"UsageByDateInfoList"`
}

type VerifyTop struct {
	// 相似度打分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *string `json:"Score,omitnil,omitempty" name:"Score"`

	// 说话人id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`

	// 说话人昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpeakerId *string `json:"SpeakerId,omitnil,omitempty" name:"SpeakerId"`
}

type VerifyTopResult struct {
	// 对比打分结果，按照打分降序排列返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyTops []*VerifyTop `json:"VerifyTops,omitnil,omitempty" name:"VerifyTops"`
}

type Vocab struct {
	// 热词表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 热词表描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 热词表ID
	VocabId *string `json:"VocabId,omitnil,omitempty" name:"VocabId"`

	// 词权重列表
	WordWeights []*HotWord `json:"WordWeights,omitnil,omitempty" name:"WordWeights"`

	// 词表创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 词表更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 热词表状态，1为默认状态即在识别时默认加载该热词表进行识别，0为初始状态
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagInfos []*string `json:"TagInfos,omitnil,omitempty" name:"TagInfos"`
}

type VoicePrintBaseData struct {
	// 说话人id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`

	// 说话人昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpeakerNick *string `json:"SpeakerNick,omitnil,omitempty" name:"SpeakerNick"`
}

type VoicePrintCompareData struct {
	// 匹配度 取值范围(0.0 - 100.0)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *string `json:"Score,omitnil,omitempty" name:"Score"`

	// 验证结果 0: 未通过 1: 通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Decision *int64 `json:"Decision,omitnil,omitempty" name:"Decision"`
}

// Predefined struct for user
type VoicePrintCompareRequestParams struct {
	// 音频格式 0: pcm, 1: wav；pcm和wav音频无损压缩，识别准确度更高
	VoiceFormat *int64 `json:"VoiceFormat,omitnil,omitempty" name:"VoiceFormat"`

	// 音频采样率，目前仅支持16k，请填写16000
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 对比源音频数据, 音频要求：base64 编码,16k采样率， 16bit位深，pcm或者wav格式， 单声道，音频时长不超过30秒的音频，base64编码数据大小不超过2M
	SrcAudioData *string `json:"SrcAudioData,omitnil,omitempty" name:"SrcAudioData"`

	// 对比目标音频数据, 音频要求：base64 编码,16k采样率， 16bit位深，pcm或者wav格式， 单声道，音频时长不超过30秒的音频，base64编码数据大小不超过2M
	DestAudioData *string `json:"DestAudioData,omitnil,omitempty" name:"DestAudioData"`
}

type VoicePrintCompareRequest struct {
	*tchttp.BaseRequest
	
	// 音频格式 0: pcm, 1: wav；pcm和wav音频无损压缩，识别准确度更高
	VoiceFormat *int64 `json:"VoiceFormat,omitnil,omitempty" name:"VoiceFormat"`

	// 音频采样率，目前仅支持16k，请填写16000
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 对比源音频数据, 音频要求：base64 编码,16k采样率， 16bit位深，pcm或者wav格式， 单声道，音频时长不超过30秒的音频，base64编码数据大小不超过2M
	SrcAudioData *string `json:"SrcAudioData,omitnil,omitempty" name:"SrcAudioData"`

	// 对比目标音频数据, 音频要求：base64 编码,16k采样率， 16bit位深，pcm或者wav格式， 单声道，音频时长不超过30秒的音频，base64编码数据大小不超过2M
	DestAudioData *string `json:"DestAudioData,omitnil,omitempty" name:"DestAudioData"`
}

func (r *VoicePrintCompareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintCompareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VoiceFormat")
	delete(f, "SampleRate")
	delete(f, "SrcAudioData")
	delete(f, "DestAudioData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VoicePrintCompareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoicePrintCompareResponseParams struct {
	// 音频声纹比对结果，包含相似度打分
	Data *VoicePrintCompareData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VoicePrintCompareResponse struct {
	*tchttp.BaseResponse
	Response *VoicePrintCompareResponseParams `json:"Response"`
}

func (r *VoicePrintCompareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintCompareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VoicePrintCountData struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 说话人id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoicePrintList []*VoicePrintBaseData `json:"VoicePrintList,omitnil,omitempty" name:"VoicePrintList"`
}

// Predefined struct for user
type VoicePrintCountRequestParams struct {
	// 分组ID,仅支持大小写字母和下划线的组合，不超过128个字符
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 统计模式
	// 0: 统计所有声纹数量
	// 1: 统计指定分组下的声纹数量
	CountMod *int64 `json:"CountMod,omitnil,omitempty" name:"CountMod"`
}

type VoicePrintCountRequest struct {
	*tchttp.BaseRequest
	
	// 分组ID,仅支持大小写字母和下划线的组合，不超过128个字符
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 统计模式
	// 0: 统计所有声纹数量
	// 1: 统计指定分组下的声纹数量
	CountMod *int64 `json:"CountMod,omitnil,omitempty" name:"CountMod"`
}

func (r *VoicePrintCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "CountMod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VoicePrintCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoicePrintCountResponseParams struct {
	// 统计数据
	Data *VoicePrintCountData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VoicePrintCountResponse struct {
	*tchttp.BaseResponse
	Response *VoicePrintCountResponseParams `json:"Response"`
}

func (r *VoicePrintCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoicePrintDeleteRequestParams struct {
	// 说话人id，说话人唯一标识
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`

	// 说话人分组ID,仅支持大小写字母和下划线的组合，不超过128个字符
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 删除模式: 
	// 0.默认值，删除该条声纹
	// 1.从分组中删除该条声纹，声纹本身不删除
	// 2.从声纹库中删除分组，仅删除分组信息，不会真正删除分组中的声纹
	DelMod *int64 `json:"DelMod,omitnil,omitempty" name:"DelMod"`
}

type VoicePrintDeleteRequest struct {
	*tchttp.BaseRequest
	
	// 说话人id，说话人唯一标识
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`

	// 说话人分组ID,仅支持大小写字母和下划线的组合，不超过128个字符
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 删除模式: 
	// 0.默认值，删除该条声纹
	// 1.从分组中删除该条声纹，声纹本身不删除
	// 2.从声纹库中删除分组，仅删除分组信息，不会真正删除分组中的声纹
	DelMod *int64 `json:"DelMod,omitnil,omitempty" name:"DelMod"`
}

func (r *VoicePrintDeleteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintDeleteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VoicePrintId")
	delete(f, "GroupId")
	delete(f, "DelMod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VoicePrintDeleteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoicePrintDeleteResponseParams struct {
	// 说话人基本信息
	Data *VoicePrintBaseData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VoicePrintDeleteResponse struct {
	*tchttp.BaseResponse
	Response *VoicePrintDeleteResponseParams `json:"Response"`
}

func (r *VoicePrintDeleteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintDeleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoicePrintEnrollRequestParams struct {
	// 音频格式 0: pcm, 1: wav
	VoiceFormat *int64 `json:"VoiceFormat,omitnil,omitempty" name:"VoiceFormat"`

	// 音频采样率，目前支持16000，单位：Hz，必填
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频数据, base64 编码, 音频时长不能超过30s，数据大小不超过2M
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 说话人昵称  不超过32字节
	SpeakerNick *string `json:"SpeakerNick,omitnil,omitempty" name:"SpeakerNick"`

	// 分组id, 仅支持大小写字母和下划线的组合，不超过128个字符
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type VoicePrintEnrollRequest struct {
	*tchttp.BaseRequest
	
	// 音频格式 0: pcm, 1: wav
	VoiceFormat *int64 `json:"VoiceFormat,omitnil,omitempty" name:"VoiceFormat"`

	// 音频采样率，目前支持16000，单位：Hz，必填
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频数据, base64 编码, 音频时长不能超过30s，数据大小不超过2M
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 说话人昵称  不超过32字节
	SpeakerNick *string `json:"SpeakerNick,omitnil,omitempty" name:"SpeakerNick"`

	// 分组id, 仅支持大小写字母和下划线的组合，不超过128个字符
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *VoicePrintEnrollRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintEnrollRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VoiceFormat")
	delete(f, "SampleRate")
	delete(f, "Data")
	delete(f, "SpeakerNick")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VoicePrintEnrollRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoicePrintEnrollResponseParams struct {
	// 说话人基本数据
	Data *VoicePrintBaseData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VoicePrintEnrollResponse struct {
	*tchttp.BaseResponse
	Response *VoicePrintEnrollResponseParams `json:"Response"`
}

func (r *VoicePrintEnrollResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintEnrollResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoicePrintGroupVerifyRequestParams struct {
	// 音频格式 0: pcm, 1: wav
	VoiceFormat *int64 `json:"VoiceFormat,omitnil,omitempty" name:"VoiceFormat"`

	// 音频采样率，目前支持16000，单位：Hz，必填
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频数据, base64 编码, 音频时长不能超过30s，数据大小不超过2M
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 分组id, 支持数字，字母，下划线，长度不超过128
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 返回打分结果降序排列topN, TopN大于0， 小于可创建声纹最大数量
	TopN *uint64 `json:"TopN,omitnil,omitempty" name:"TopN"`
}

type VoicePrintGroupVerifyRequest struct {
	*tchttp.BaseRequest
	
	// 音频格式 0: pcm, 1: wav
	VoiceFormat *int64 `json:"VoiceFormat,omitnil,omitempty" name:"VoiceFormat"`

	// 音频采样率，目前支持16000，单位：Hz，必填
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频数据, base64 编码, 音频时长不能超过30s，数据大小不超过2M
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 分组id, 支持数字，字母，下划线，长度不超过128
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 返回打分结果降序排列topN, TopN大于0， 小于可创建声纹最大数量
	TopN *uint64 `json:"TopN,omitnil,omitempty" name:"TopN"`
}

func (r *VoicePrintGroupVerifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintGroupVerifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VoiceFormat")
	delete(f, "SampleRate")
	delete(f, "Data")
	delete(f, "GroupId")
	delete(f, "TopN")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VoicePrintGroupVerifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoicePrintGroupVerifyResponseParams struct {
	// TopN 返回结果;系统建议打分70分以上为同一个人音色，评分也取决于音频质量、长度等其他原因影响，您可以按照业务需求适当提高或降低分数要求
	Data *VerifyTopResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VoicePrintGroupVerifyResponse struct {
	*tchttp.BaseResponse
	Response *VoicePrintGroupVerifyResponseParams `json:"Response"`
}

func (r *VoicePrintGroupVerifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintGroupVerifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoicePrintUpdateRequestParams struct {
	// 音频格式 0: pcm, 1: wav
	VoiceFormat *int64 `json:"VoiceFormat,omitnil,omitempty" name:"VoiceFormat"`

	// 音频采样率 目前仅支持16000 单位Hz
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 说话人id， 说话人唯一标识
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`

	// 音频数据, base64 编码, 音频时长不能超过30s，数据大小不超过2M	
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 说话人昵称  不超过32字节
	SpeakerNick *string `json:"SpeakerNick,omitnil,omitempty" name:"SpeakerNick"`
}

type VoicePrintUpdateRequest struct {
	*tchttp.BaseRequest
	
	// 音频格式 0: pcm, 1: wav
	VoiceFormat *int64 `json:"VoiceFormat,omitnil,omitempty" name:"VoiceFormat"`

	// 音频采样率 目前仅支持16000 单位Hz
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 说话人id， 说话人唯一标识
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`

	// 音频数据, base64 编码, 音频时长不能超过30s，数据大小不超过2M	
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 说话人昵称  不超过32字节
	SpeakerNick *string `json:"SpeakerNick,omitnil,omitempty" name:"SpeakerNick"`
}

func (r *VoicePrintUpdateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintUpdateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VoiceFormat")
	delete(f, "SampleRate")
	delete(f, "VoicePrintId")
	delete(f, "Data")
	delete(f, "SpeakerNick")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VoicePrintUpdateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoicePrintUpdateResponseParams struct {
	// 说话人基础数据
	Data *VoicePrintBaseData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VoicePrintUpdateResponse struct {
	*tchttp.BaseResponse
	Response *VoicePrintUpdateResponseParams `json:"Response"`
}

func (r *VoicePrintUpdateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintUpdateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VoicePrintVerifyData struct {
	// 说话人id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`

	// 匹配度 取值范围(0.0 - 100.0)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *string `json:"Score,omitnil,omitempty" name:"Score"`

	// 验证结果 0: 未通过 1: 通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Decision *int64 `json:"Decision,omitnil,omitempty" name:"Decision"`
}

// Predefined struct for user
type VoicePrintVerifyRequestParams struct {
	// 音频格式 0: pcm, 1: wav
	VoiceFormat *int64 `json:"VoiceFormat,omitnil,omitempty" name:"VoiceFormat"`

	// 音频采样率，目前支持16000，单位：Hz，必填
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频数据, base64 编码, 音频时长不能超过30s，数据大小不超过2M	
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 说话人id, 说话人唯一标识
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`
}

type VoicePrintVerifyRequest struct {
	*tchttp.BaseRequest
	
	// 音频格式 0: pcm, 1: wav
	VoiceFormat *int64 `json:"VoiceFormat,omitnil,omitempty" name:"VoiceFormat"`

	// 音频采样率，目前支持16000，单位：Hz，必填
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频数据, base64 编码, 音频时长不能超过30s，数据大小不超过2M	
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 说话人id, 说话人唯一标识
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`
}

func (r *VoicePrintVerifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintVerifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VoiceFormat")
	delete(f, "SampleRate")
	delete(f, "Data")
	delete(f, "VoicePrintId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VoicePrintVerifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VoicePrintVerifyResponseParams struct {
	// 说话人验证数据
	Data *VoicePrintVerifyData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VoicePrintVerifyResponse struct {
	*tchttp.BaseResponse
	Response *VoicePrintVerifyResponseParams `json:"Response"`
}

func (r *VoicePrintVerifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VoicePrintVerifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}