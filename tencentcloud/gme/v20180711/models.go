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

package v20180711

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AgeDetectTask struct {
	// 数据唯一ID
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 数据文件的url，为 urlencode 编码,音频文件格式支持的类型：.wav、.m4a、.amr、.mp3、.aac、.wma、.ogg
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type AgeDetectTaskResult struct {
	// 数据唯一ID
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 数据文件的url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 任务状态，0: 已创建，1:运行中，2:正常结束，3:异常结束，4:运行超时
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务结果：0: 成年，1:未成年，100:未知
	Age *uint64 `json:"Age,omitnil,omitempty" name:"Age"`
}

type AgentConfig struct {
	// 机器人的UserId，用于进房发起任务。【注意】这个UserId不能与当前房间内的主播观众UserId重复。如果一个房间发起多个任务时，机器人的UserId也不能相互重复，否则会中断前一个任务。需要保证机器人UserId在房间内唯一。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 机器人UserId对应的校验签名，即UserId和UserSig相当于机器人进房的登录密码。
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`

	// 机器人拉流的UserId, 填写后，机器人会拉取该UserId的流进行实时处理
	TargetUserId *string `json:"TargetUserId,omitnil,omitempty" name:"TargetUserId"`

	// 房间内超过MaxIdleTime 没有推流，后台自动关闭任务，默认值是60s。
	MaxIdleTime *uint64 `json:"MaxIdleTime,omitnil,omitempty" name:"MaxIdleTime"`

	// 机器人的欢迎语
	WelcomeMessage *string `json:"WelcomeMessage,omitnil,omitempty" name:"WelcomeMessage"`

	// 智能打断模式，默认为0，0表示服务端自动打断，1表示服务端不打断，由端上发送打断信令进行打断
	InterruptMode *uint64 `json:"InterruptMode,omitnil,omitempty" name:"InterruptMode"`

	// InterruptMode为0时使用，单位为毫秒，默认为500ms。表示服务端检测到持续InterruptSpeechDuration毫秒的人声则进行打断。
	InterruptSpeechDuration *uint64 `json:"InterruptSpeechDuration,omitnil,omitempty" name:"InterruptSpeechDuration"`

	// 控制新一轮对话的触发方式，默认为0。
	// - 0表示当服务端语音识别检测出的完整一句话后，自动触发一轮新的对话。
	// - 1表示客户端在收到字幕消息后，自行决定是否手动发送聊天信令触发一轮新的对话。
	TurnDetectionMode *uint64 `json:"TurnDetectionMode,omitnil,omitempty" name:"TurnDetectionMode"`

	// 是否过滤掉用户只说了一个字的句子，true表示过滤，false表示不过滤，默认值为true
	FilterOneWord *bool `json:"FilterOneWord,omitnil,omitempty" name:"FilterOneWord"`

	// 欢迎消息优先级，0默认，1高优，高优不能被打断。
	WelcomeMessagePriority *uint64 `json:"WelcomeMessagePriority,omitnil,omitempty" name:"WelcomeMessagePriority"`

	// 用于过滤LLM返回内容，不播放括号中的内容。
	// 1：中文括号（）
	// 2：英文括号()
	// 3：中文方括号【】
	// 4：英文方括号[]
	// 5：英文花括号{}
	// 默认值为空，表示不进行过滤。
	FilterBracketsContent *uint64 `json:"FilterBracketsContent,omitnil,omitempty" name:"FilterBracketsContent"`

	// 环境音设置	
	AmbientSound *AmbientSound `json:"AmbientSound,omitnil,omitempty" name:"AmbientSound"`

	// 声纹配置	
	VoicePrint *VoicePrint `json:"VoicePrint,omitnil,omitempty" name:"VoicePrint"`

	// 与WelcomeMessage参数互斥，当该参数有值时，WelcomeMessage将失效。\n在对话开始后把该消息送到大模型来获取欢迎语。	
	InitLLMMessage *string `json:"InitLLMMessage,omitnil,omitempty" name:"InitLLMMessage"`

	// 语义断句检测	
	TurnDetection *TurnDetection `json:"TurnDetection,omitnil,omitempty" name:"TurnDetection"`

	// 机器人字幕显示模式。 - 0表示尽快显示，不会和音频播放进行同步。此时字幕全量下发，后面的字幕会包含前面的字幕。 - 1表示句子级别的实时显示，会和音频播放进行同步，只有当前句子对应的音频播放完后，下一条字幕才会下发。此时字幕增量下发，端上需要把前后的字幕进行拼接才是完整字幕。	
	SubtitleMode *uint64 `json:"SubtitleMode,omitnil,omitempty" name:"SubtitleMode"`
}

type AmbientSound struct {
	// 环境场景选择
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 控制环境音的音量。取值的范围是 [0,2]。值越低，环境音越小；值越高，环境音越响亮。如果未设置，则使用默认值 1。
	Volume *float64 `json:"Volume,omitnil,omitempty" name:"Volume"`
}

type AppStatisticsItem struct {
	// 实时语音统计数据
	RealtimeSpeechStatisticsItem *RealTimeSpeechStatisticsItem `json:"RealtimeSpeechStatisticsItem,omitnil,omitempty" name:"RealtimeSpeechStatisticsItem"`

	// 语音消息统计数据
	VoiceMessageStatisticsItem *VoiceMessageStatisticsItem `json:"VoiceMessageStatisticsItem,omitnil,omitempty" name:"VoiceMessageStatisticsItem"`

	// 语音过滤统计数据
	VoiceFilterStatisticsItem *VoiceFilterStatisticsItem `json:"VoiceFilterStatisticsItem,omitnil,omitempty" name:"VoiceFilterStatisticsItem"`

	// 统计时间
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 录音转文本用量统计数据
	AudioTextStatisticsItem *AudioTextStatisticsItem `json:"AudioTextStatisticsItem,omitnil,omitempty" name:"AudioTextStatisticsItem"`

	// 流式转文本用量数据
	StreamTextStatisticsItem *StreamTextStatisticsItem `json:"StreamTextStatisticsItem,omitnil,omitempty" name:"StreamTextStatisticsItem"`

	// 海外转文本用量数据
	OverseaTextStatisticsItem *OverseaTextStatisticsItem `json:"OverseaTextStatisticsItem,omitnil,omitempty" name:"OverseaTextStatisticsItem"`

	// 实时语音转文本用量数据
	RealtimeTextStatisticsItem *RealtimeTextStatisticsItem `json:"RealtimeTextStatisticsItem,omitnil,omitempty" name:"RealtimeTextStatisticsItem"`
}

type ApplicationDataStatistics struct {
	// 应用ID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// Dau统计项数目
	DauDataNum *uint64 `json:"DauDataNum,omitnil,omitempty" name:"DauDataNum"`

	// 大陆地区Dau统计数据，单位人
	DauDataMainland []*StatisticsItem `json:"DauDataMainland,omitnil,omitempty" name:"DauDataMainland"`

	// 海外地区Dau统计数据，单位人
	DauDataOversea []*StatisticsItem `json:"DauDataOversea,omitnil,omitempty" name:"DauDataOversea"`

	// 大陆和海外地区Dau统计数据汇总，单位人
	DauDataSum []*StatisticsItem `json:"DauDataSum,omitnil,omitempty" name:"DauDataSum"`

	// 实时语音时长统计项数目
	DurationDataNum *uint64 `json:"DurationDataNum,omitnil,omitempty" name:"DurationDataNum"`

	// 大陆地区实时语音时长统计数据，单位分钟
	DurationDataMainland []*StatisticsItem `json:"DurationDataMainland,omitnil,omitempty" name:"DurationDataMainland"`

	// 海外地区实时语音时长统计数据，单位分钟
	DurationDataOversea []*StatisticsItem `json:"DurationDataOversea,omitnil,omitempty" name:"DurationDataOversea"`

	// 大陆和海外地区实时语音时长统计数据汇总，单位分钟
	DurationDataSum []*StatisticsItem `json:"DurationDataSum,omitnil,omitempty" name:"DurationDataSum"`

	// Pcu统计项数目
	PcuDataNum *uint64 `json:"PcuDataNum,omitnil,omitempty" name:"PcuDataNum"`

	// 大陆地区Pcu统计数据，单位人
	PcuDataMainland []*StatisticsItem `json:"PcuDataMainland,omitnil,omitempty" name:"PcuDataMainland"`

	// 海外地区Pcu统计数据，单位人
	PcuDataOversea []*StatisticsItem `json:"PcuDataOversea,omitnil,omitempty" name:"PcuDataOversea"`

	// 大陆和海外地区Pcu统计数据汇总，单位人
	PcuDataSum []*StatisticsItem `json:"PcuDataSum,omitnil,omitempty" name:"PcuDataSum"`

	// 小游戏时长统计项数目
	MiniGameDataNum *uint64 `json:"MiniGameDataNum,omitnil,omitempty" name:"MiniGameDataNum"`

	// 大陆地区小游戏时长统计数据，单位分钟
	MiniGameDataMainland []*StatisticsItem `json:"MiniGameDataMainland,omitnil,omitempty" name:"MiniGameDataMainland"`

	// 海外地区小游戏时长统计数据，单位分钟
	MiniGameDataOversea []*StatisticsItem `json:"MiniGameDataOversea,omitnil,omitempty" name:"MiniGameDataOversea"`

	// 大陆和海外地区小游戏时长统计数据汇总，单位分钟
	MiniGameDataSum []*StatisticsItem `json:"MiniGameDataSum,omitnil,omitempty" name:"MiniGameDataSum"`
}

type ApplicationList struct {
	// 服务开关状态
	ServiceConf *ServiceStatus `json:"ServiceConf,omitnil,omitempty" name:"ServiceConf"`

	// 应用ID(AppID)
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 项目ID，默认为0
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 应用状态，返回0表示正常，1表示关闭，2表示欠费停服，3表示欠费回收
	AppStatus *uint64 `json:"AppStatus,omitnil,omitempty" name:"AppStatus"`

	// 创建时间，Unix时间戳格式
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 应用类型，无需关注此数值
	AppType *uint64 `json:"AppType,omitnil,omitempty" name:"AppType"`
}

type AsrConf struct {
	// 语音转文本服务开关，取值：open/close
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type AudioTextStatisticsItem struct {
	// 统计值，单位：秒
	Data *float64 `json:"Data,omitnil,omitempty" name:"Data"`
}

// Predefined struct for user
type ControlAIConversationRequestParams struct {
	// 任务唯一标识
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 控制命令，目前支持命令如下：- ServerPushText，服务端发送文本给AI机器人，AI机器人会播报该文本. - InvokeLLM，服务端发送文本给大模型，触发对话
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 服务端发送播报文本命令，当Command为ServerPushText时必填
	ServerPushText *ServerPushText `json:"ServerPushText,omitnil,omitempty" name:"ServerPushText"`

	// 服务端发送命令主动请求大模型,当Command为InvokeLLM时会把content请求到大模型,头部增加X-Invoke-LLM="1"
	InvokeLLM *InvokeLLM `json:"InvokeLLM,omitnil,omitempty" name:"InvokeLLM"`
}

type ControlAIConversationRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一标识
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 控制命令，目前支持命令如下：- ServerPushText，服务端发送文本给AI机器人，AI机器人会播报该文本. - InvokeLLM，服务端发送文本给大模型，触发对话
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 服务端发送播报文本命令，当Command为ServerPushText时必填
	ServerPushText *ServerPushText `json:"ServerPushText,omitnil,omitempty" name:"ServerPushText"`

	// 服务端发送命令主动请求大模型,当Command为InvokeLLM时会把content请求到大模型,头部增加X-Invoke-LLM="1"
	InvokeLLM *InvokeLLM `json:"InvokeLLM,omitnil,omitempty" name:"InvokeLLM"`
}

func (r *ControlAIConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlAIConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Command")
	delete(f, "ServerPushText")
	delete(f, "InvokeLLM")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlAIConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlAIConversationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ControlAIConversationResponse struct {
	*tchttp.BaseResponse
	Response *ControlAIConversationResponseParams `json:"Response"`
}

func (r *ControlAIConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlAIConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgeDetectTaskRequestParams struct {
	// 应用id
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 语音检测子任务列表，列表最多支持100个检测子任务。结构体中包含：
	// <li>DataId：数据的唯一ID</li>
	// <li>Url：数据文件的url，为 urlencode 编码，流式则为拉流地址</li>
	Tasks []*AgeDetectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 任务结束时gme后台会自动触发回调
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`
}

type CreateAgeDetectTaskRequest struct {
	*tchttp.BaseRequest
	
	// 应用id
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 语音检测子任务列表，列表最多支持100个检测子任务。结构体中包含：
	// <li>DataId：数据的唯一ID</li>
	// <li>Url：数据文件的url，为 urlencode 编码，流式则为拉流地址</li>
	Tasks []*AgeDetectTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 任务结束时gme后台会自动触发回调
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`
}

func (r *CreateAgeDetectTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgeDetectTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "Tasks")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgeDetectTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgeDetectTaskResponseParams struct {
	// 本次任务提交后唯一id，用于获取任务运行结果
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgeDetectTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgeDetectTaskResponseParams `json:"Response"`
}

func (r *CreateAgeDetectTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgeDetectTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppRequestParams struct {
	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 腾讯云项目ID，默认为0，表示默认项目
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 需要支持的引擎列表，默认全选。
	// 取值：android/ios/unity/cocos/unreal/windows
	EngineList []*string `json:"EngineList,omitnil,omitempty" name:"EngineList"`

	// 服务区域列表，默认全选。
	// 取值：mainland-大陆地区，hmt-港澳台，sea-东南亚，na-北美，eu-欧洲，jpkr-日韩亚太，sa-南美，oc-澳洲，me-中东
	RegionList []*string `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 实时语音服务配置数据
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitnil,omitempty" name:"RealtimeSpeechConf"`

	// 语音消息服务配置数据
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitnil,omitempty" name:"VoiceMessageConf"`

	// 语音分析服务配置数据
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitnil,omitempty" name:"VoiceFilterConf"`

	// 语音转文本配置数据
	AsrConf *AsrConf `json:"AsrConf,omitnil,omitempty" name:"AsrConf"`

	// 需要添加的标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 腾讯云项目ID，默认为0，表示默认项目
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 需要支持的引擎列表，默认全选。
	// 取值：android/ios/unity/cocos/unreal/windows
	EngineList []*string `json:"EngineList,omitnil,omitempty" name:"EngineList"`

	// 服务区域列表，默认全选。
	// 取值：mainland-大陆地区，hmt-港澳台，sea-东南亚，na-北美，eu-欧洲，jpkr-日韩亚太，sa-南美，oc-澳洲，me-中东
	RegionList []*string `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 实时语音服务配置数据
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitnil,omitempty" name:"RealtimeSpeechConf"`

	// 语音消息服务配置数据
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitnil,omitempty" name:"VoiceMessageConf"`

	// 语音分析服务配置数据
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitnil,omitempty" name:"VoiceFilterConf"`

	// 语音转文本配置数据
	AsrConf *AsrConf `json:"AsrConf,omitnil,omitempty" name:"AsrConf"`

	// 需要添加的标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "ProjectId")
	delete(f, "EngineList")
	delete(f, "RegionList")
	delete(f, "RealtimeSpeechConf")
	delete(f, "VoiceMessageConf")
	delete(f, "VoiceFilterConf")
	delete(f, "AsrConf")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppResp struct {
	// 应用ID，由后台自动生成。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 应用名称，透传输入参数的AppName
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 项目ID，透传输入的ProjectId
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 应用密钥，GME SDK初始化时使用
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 服务创建时间戳
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实时语音服务配置数据
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitnil,omitempty" name:"RealtimeSpeechConf"`

	// 语音消息服务配置数据
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitnil,omitempty" name:"VoiceMessageConf"`

	// 语音分析服务配置数据
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitnil,omitempty" name:"VoiceFilterConf"`

	// 语音转文本服务配置数据
	AsrConf *AsrConf `json:"AsrConf,omitnil,omitempty" name:"AsrConf"`
}

// Predefined struct for user
type CreateAppResponseParams struct {
	// 创建应用返回数据
	Data *CreateAppResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppResponseParams `json:"Response"`
}

func (r *CreateAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomizationRequestParams struct {
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 文本文件的下载地址，服务会从该地址下载文件，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitnil,omitempty" name:"TextUrl"`

	// 模型名称，名称长度不超过36，默认为BizId。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

type CreateCustomizationRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 文本文件的下载地址，服务会从该地址下载文件，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitnil,omitempty" name:"TextUrl"`

	// 模型名称，名称长度不超过36，默认为BizId。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
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
	delete(f, "BizId")
	delete(f, "TextUrl")
	delete(f, "ModelName")
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
type CreateScanUserRequestParams struct {
	// 应用ID，登录控制台 - 服务管理创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 需要新增送检的用户号。示例：1234
	// (若UserId不填，则UserIdString必填；两者选其一；两者都填以UserIdString为准)
	UserId *uint64 `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 需要新增送检的用户号，长度不超过1024字符。示例："1234"(若UserIdString不填，则UserId必填；两者选其一；两者都填以UserIdString为准)
	UserIdString *string `json:"UserIdString,omitnil,omitempty" name:"UserIdString"`

	// 当前用户送检过期时间，单位：秒。
	// 若参数不为0，则在过期时间之后，用户不会被送检。
	// 若参数为0，则送检配置不会自动失效。 
	ExpirationTime *uint64 `json:"ExpirationTime,omitnil,omitempty" name:"ExpirationTime"`
}

type CreateScanUserRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，登录控制台 - 服务管理创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 需要新增送检的用户号。示例：1234
	// (若UserId不填，则UserIdString必填；两者选其一；两者都填以UserIdString为准)
	UserId *uint64 `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 需要新增送检的用户号，长度不超过1024字符。示例："1234"(若UserIdString不填，则UserId必填；两者选其一；两者都填以UserIdString为准)
	UserIdString *string `json:"UserIdString,omitnil,omitempty" name:"UserIdString"`

	// 当前用户送检过期时间，单位：秒。
	// 若参数不为0，则在过期时间之后，用户不会被送检。
	// 若参数为0，则送检配置不会自动失效。 
	ExpirationTime *uint64 `json:"ExpirationTime,omitnil,omitempty" name:"ExpirationTime"`
}

func (r *CreateScanUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScanUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "UserId")
	delete(f, "UserIdString")
	delete(f, "ExpirationTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScanUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScanUserResponseParams struct {
	// 返回结果码
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateScanUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateScanUserResponseParams `json:"Response"`
}

func (r *CreateScanUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScanUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomizationConfigs struct {
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型状态，-1下线状态，1上线状态, 0训练中, -2训练失败, 3上线中, 4下线中
	ModelState *int64 `json:"ModelState,omitnil,omitempty" name:"ModelState"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 文本文件的下载地址，服务会从该地址下载文件，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitnil,omitempty" name:"TextUrl"`

	// 更新时间，11位时间戳
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type DeleteCustomizationRequestParams struct {
	// 删除的模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`
}

type DeleteCustomizationRequest struct {
	*tchttp.BaseRequest
	
	// 删除的模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`
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
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomizationResponseParams struct {
	// 返回值。0为成功，非0为失败。
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

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

type DeleteResult struct {
	// 错误码，0-剔除成功 其他-剔除失败
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误描述
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`
}

// Predefined struct for user
type DeleteRoomMemberRequestParams struct {
	// 要操作的房间id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 剔除类型 1-删除房间 2-剔除用户
	DeleteType *uint64 `json:"DeleteType,omitnil,omitempty" name:"DeleteType"`

	// 应用id
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 要剔除的用户列表（整型）
	Uids []*string `json:"Uids,omitnil,omitempty" name:"Uids"`

	// 要剔除的用户列表（字符串类型）
	StrUids []*string `json:"StrUids,omitnil,omitempty" name:"StrUids"`
}

type DeleteRoomMemberRequest struct {
	*tchttp.BaseRequest
	
	// 要操作的房间id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 剔除类型 1-删除房间 2-剔除用户
	DeleteType *uint64 `json:"DeleteType,omitnil,omitempty" name:"DeleteType"`

	// 应用id
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 要剔除的用户列表（整型）
	Uids []*string `json:"Uids,omitnil,omitempty" name:"Uids"`

	// 要剔除的用户列表（字符串类型）
	StrUids []*string `json:"StrUids,omitnil,omitempty" name:"StrUids"`
}

func (r *DeleteRoomMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoomMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "DeleteType")
	delete(f, "BizId")
	delete(f, "Uids")
	delete(f, "StrUids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoomMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoomMemberResponseParams struct {
	// 剔除房间或成员的操作结果
	DeleteResult *DeleteResult `json:"DeleteResult,omitnil,omitempty" name:"DeleteResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRoomMemberResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoomMemberResponseParams `json:"Response"`
}

func (r *DeleteRoomMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoomMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScanUserRequestParams struct {
	// 应用ID，登录控制台 - 服务管理创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 需要删除送检的用户号。示例：1234
	// (若UserId不填，则UserIdString必填；两者选其一；两者都填以UserIdString为准)
	UserId *uint64 `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 需要删除送检的用户号，长度不超过1024字符。示例："1234"(若UserIdString不填，则UserId必填；两者选其一；两者都填以UserIdString为准)
	UserIdString *string `json:"UserIdString,omitnil,omitempty" name:"UserIdString"`
}

type DeleteScanUserRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，登录控制台 - 服务管理创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 需要删除送检的用户号。示例：1234
	// (若UserId不填，则UserIdString必填；两者选其一；两者都填以UserIdString为准)
	UserId *uint64 `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 需要删除送检的用户号，长度不超过1024字符。示例："1234"(若UserIdString不填，则UserId必填；两者选其一；两者都填以UserIdString为准)
	UserIdString *string `json:"UserIdString,omitnil,omitempty" name:"UserIdString"`
}

func (r *DeleteScanUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScanUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "UserId")
	delete(f, "UserIdString")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScanUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScanUserResponseParams struct {
	// 返回结果码
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteScanUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScanUserResponseParams `json:"Response"`
}

func (r *DeleteScanUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScanUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVoicePrintRequestParams struct {
	// 声纹信息ID
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`
}

type DeleteVoicePrintRequest struct {
	*tchttp.BaseRequest
	
	// 声纹信息ID
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`
}

func (r *DeleteVoicePrintRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVoicePrintRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VoicePrintId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVoicePrintRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVoicePrintResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVoicePrintResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVoicePrintResponseParams `json:"Response"`
}

func (r *DeleteVoicePrintResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVoicePrintResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIConversationRequestParams struct {
	// GME的SdkAppId，和开启转录任务的房间使用的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 唯一标识一次任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeAIConversationRequest struct {
	*tchttp.BaseRequest
	
	// GME的SdkAppId，和开启转录任务的房间使用的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 唯一标识一次任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeAIConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIConversationResponseParams struct {
	// 任务开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务状态。有4个值：1、Idle表示任务未开始2、Preparing表示任务准备中3、InProgress表示任务正在运行4、Stopped表示任务已停止，正在清理资源中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一标识一次任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 开启对话任务时填写的SessionId，如果没写则不返回。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIConversationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIConversationResponseParams `json:"Response"`
}

func (r *DescribeAIConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgeDetectTaskRequestParams struct {
	// 应用id
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// [创建年龄语音识别任务](https://cloud.tencent.com/document/product/607/60620)时返回的taskid
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeAgeDetectTaskRequest struct {
	*tchttp.BaseRequest
	
	// 应用id
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// [创建年龄语音识别任务](https://cloud.tencent.com/document/product/607/60620)时返回的taskid
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeAgeDetectTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgeDetectTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgeDetectTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgeDetectTaskResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 语音检测返回。Results 字段是 JSON 数组，每一个元素包含：
	// DataId： 请求中对应的 DataId。
	// Url ：该请求中对应的 Url。
	// Status ：子任务状态，0:已创建，1:运行中，2:已完成，3:任务异常，4:任务超时。
	// Age ：子任务完成后的结果，0:成年人，1:未成年人，100:未知结果。
	Results []*AgeDetectTaskResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgeDetectTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgeDetectTaskResponseParams `json:"Response"`
}

func (r *DescribeAgeDetectTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgeDetectTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppStatisticsRequestParams struct {
	// GME应用ID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 数据开始时间，东八区时间，格式: 年-月-日，如: 2018-07-13
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 数据结束时间，东八区时间，格式: 年-月-日，如: 2018-07-13
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 要查询的服务列表，取值：RealTimeSpeech/VoiceMessage/VoiceFilter/SpeechToText
	Services []*string `json:"Services,omitnil,omitempty" name:"Services"`
}

type DescribeAppStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// GME应用ID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 数据开始时间，东八区时间，格式: 年-月-日，如: 2018-07-13
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 数据结束时间，东八区时间，格式: 年-月-日，如: 2018-07-13
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 要查询的服务列表，取值：RealTimeSpeech/VoiceMessage/VoiceFilter/SpeechToText
	Services []*string `json:"Services,omitnil,omitempty" name:"Services"`
}

func (r *DescribeAppStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Services")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppStatisticsResp struct {
	// 应用用量统计数据
	AppStatistics []*AppStatisticsItem `json:"AppStatistics,omitnil,omitempty" name:"AppStatistics"`
}

// Predefined struct for user
type DescribeAppStatisticsResponseParams struct {
	// 应用用量统计数据
	Data *DescribeAppStatisticsResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppStatisticsResponseParams `json:"Response"`
}

func (r *DescribeAppStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationDataRequestParams struct {
	// 应用ID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 数据开始时间，格式为 年-月-日，如: 2018-07-13
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 数据结束时间，格式为 年-月-日，如: 2018-07-13
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

type DescribeApplicationDataRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 数据开始时间，格式为 年-月-日，如: 2018-07-13
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 数据结束时间，格式为 年-月-日，如: 2018-07-13
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

func (r *DescribeApplicationDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationDataResponseParams struct {
	// 应用统计数据
	Data *ApplicationDataStatistics `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationDataResponseParams `json:"Response"`
}

func (r *DescribeApplicationDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationListRequestParams struct {
	// 项目ID，0表示默认项目，-1表示所有项目，如果需要查找具体项目下的应用列表，请填入具体项目ID，项目ID在项目管理中查看 https://console.cloud.tencent.com/project
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码ID，0表示第一页，以此后推。默认填0
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 每页展示应用数量。默认填200
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 所查找应用名称的关键字，支持模糊匹配查找。空串表示查询所有应用
	SearchText *string `json:"SearchText,omitnil,omitempty" name:"SearchText"`

	// 标签列表
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// 查找过滤关键字列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeApplicationListRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID，0表示默认项目，-1表示所有项目，如果需要查找具体项目下的应用列表，请填入具体项目ID，项目ID在项目管理中查看 https://console.cloud.tencent.com/project
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码ID，0表示第一页，以此后推。默认填0
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 每页展示应用数量。默认填200
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 所查找应用名称的关键字，支持模糊匹配查找。空串表示查询所有应用
	SearchText *string `json:"SearchText,omitnil,omitempty" name:"SearchText"`

	// 标签列表
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// 查找过滤关键字列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeApplicationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "SearchText")
	delete(f, "TagSet")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationListResponseParams struct {
	// 获取应用列表返回
	ApplicationList []*ApplicationList `json:"ApplicationList,omitnil,omitempty" name:"ApplicationList"`

	// 应用总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationListResponseParams `json:"Response"`
}

func (r *DescribeApplicationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealtimeScanConfigRequestParams struct {
	// 应用ID。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`
}

type DescribeRealtimeScanConfigRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`
}

func (r *DescribeRealtimeScanConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeScanConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRealtimeScanConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealtimeScanConfigResponseParams struct {
	// 返回结果码，0正常，非0失败
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 应用ID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 送检类型，0: 全量送审，1: 自定义送审
	AuditType *int64 `json:"AuditType,omitnil,omitempty" name:"AuditType"`

	// 用户号正则表达式。
	// 符合此正则表达式规则的用户号将被送检。示例：^6.*（表示所有以6开头的用户号将被送检）
	UserIdRegex []*string `json:"UserIdRegex,omitnil,omitempty" name:"UserIdRegex"`

	// 房间号正则表达式。
	// 符合此正则表达式规则的房间号将被送检。示例：^6.*（表示所有以6开头的房间号将被送检）
	RoomIdRegex []*string `json:"RoomIdRegex,omitnil,omitempty" name:"RoomIdRegex"`

	// 用户号字符串，逗号分隔，示例："0001,0002,0003"
	UserIdString *string `json:"UserIdString,omitnil,omitempty" name:"UserIdString"`

	// 房间号字符串，逗号分隔，示例："0001,0002,0003"
	RoomIdString *string `json:"RoomIdString,omitnil,omitempty" name:"RoomIdString"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRealtimeScanConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRealtimeScanConfigResponseParams `json:"Response"`
}

func (r *DescribeRealtimeScanConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeScanConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordInfoRequestParams struct {
	// 进行中的任务taskid（StartRecord接口返回）。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 应用ID。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`
}

type DescribeRecordInfoRequest struct {
	*tchttp.BaseRequest
	
	// 进行中的任务taskid（StartRecord接口返回）。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 应用ID。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`
}

func (r *DescribeRecordInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordInfoResponseParams struct {
	// 录制信息。
	RecordInfo []*RecordInfo `json:"RecordInfo,omitnil,omitempty" name:"RecordInfo"`

	// 录制类型：1代表单流 2代表混流 3代表单流和混流。
	RecordMode *uint64 `json:"RecordMode,omitnil,omitempty" name:"RecordMode"`

	// 房间ID。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordInfoResponseParams `json:"Response"`
}

func (r *DescribeRecordInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomInfoRequestParams struct {
	// 应用ID，登录[控制台 - 服务管理](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号列表，最大不能超过10个（RoomIds、StrRoomIds必须填一个）
	RoomIds []*uint64 `json:"RoomIds,omitnil,omitempty" name:"RoomIds"`

	// 字符串类型房间号列表，最大不能超过10个（RoomIds、StrRoomIds必须填一个）
	StrRoomIds []*string `json:"StrRoomIds,omitnil,omitempty" name:"StrRoomIds"`
}

type DescribeRoomInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，登录[控制台 - 服务管理](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号列表，最大不能超过10个（RoomIds、StrRoomIds必须填一个）
	RoomIds []*uint64 `json:"RoomIds,omitnil,omitempty" name:"RoomIds"`

	// 字符串类型房间号列表，最大不能超过10个（RoomIds、StrRoomIds必须填一个）
	StrRoomIds []*string `json:"StrRoomIds,omitnil,omitempty" name:"StrRoomIds"`
}

func (r *DescribeRoomInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomIds")
	delete(f, "StrRoomIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomInfoResponseParams struct {
	// 操作结果, 0成功, 非0失败
	Result *uint64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 房间用户信息
	RoomUsers []*RoomUser `json:"RoomUsers,omitnil,omitempty" name:"RoomUsers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRoomInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoomInfoResponseParams `json:"Response"`
}

func (r *DescribeRoomInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanResult struct {
	// 业务返回码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 数据唯一 ID
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 检测完成的时间戳
	ScanFinishTime *uint64 `json:"ScanFinishTime,omitnil,omitempty" name:"ScanFinishTime"`

	// 是否违规
	HitFlag *bool `json:"HitFlag,omitnil,omitempty" name:"HitFlag"`

	// 是否为流
	Live *bool `json:"Live,omitnil,omitempty" name:"Live"`

	// 业务返回描述
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 检测结果，Code 为 0 时返回
	ScanPiece []*ScanPiece `json:"ScanPiece,omitnil,omitempty" name:"ScanPiece"`

	// 提交检测的时间戳
	ScanStartTime *uint64 `json:"ScanStartTime,omitnil,omitempty" name:"ScanStartTime"`

	// 语音检测场景，对应请求时的 Scene
	Scenes []*string `json:"Scenes,omitnil,omitempty" name:"Scenes"`

	// 语音检测任务 ID，由后台分配
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 文件或接流地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 检测任务执行结果状态，分别为：
	// <li>Start: 任务开始</li>
	// <li>Success: 成功结束</li>
	// <li>Error: 异常</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 提交检测的应用 ID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`
}

// Predefined struct for user
type DescribeScanResultListRequestParams struct {
	// 应用 ID，登录[控制台](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 查询的任务 ID 列表，任务 ID 列表最多支持 100 个。
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 任务返回结果数量，默认10，上限500。大文件任务忽略此参数，返回全量结果
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeScanResultListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID，登录[控制台](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 查询的任务 ID 列表，任务 ID 列表最多支持 100 个。
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 任务返回结果数量，默认10，上限500。大文件任务忽略此参数，返回全量结果
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeScanResultListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanResultListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "TaskIdList")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanResultListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanResultListResponseParams struct {
	// 要查询的语音检测任务的结果
	Data []*DescribeScanResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScanResultListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanResultListResponseParams `json:"Response"`
}

func (r *DescribeScanResultListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanResultListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoRequestParams struct {
	// 应用ID。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 房间ID。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 房间ID。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DescribeTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoResponseParams struct {
	// 进行中的任务taskid（StartRecord接口返回）。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 录制类型：1代表单流 2代表混流 3代表单流和混流。
	RecordMode *uint64 `json:"RecordMode,omitnil,omitempty" name:"RecordMode"`

	// 指定订阅流白名单或者黑名单。
	SubscribeRecordUserIds *SubscribeRecordUserIds `json:"SubscribeRecordUserIds,omitnil,omitempty" name:"SubscribeRecordUserIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskInfoResponseParams `json:"Response"`
}

func (r *DescribeTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInAndOutTimeRequestParams struct {
	// 应用ID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 房间ID
	RoomId *int64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 用户ID
	UserId *int64 `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 字符串类型用户ID
	UserIdStr *string `json:"UserIdStr,omitnil,omitempty" name:"UserIdStr"`

	// 字符串类型房间ID
	RoomIdStr *string `json:"RoomIdStr,omitnil,omitempty" name:"RoomIdStr"`
}

type DescribeUserInAndOutTimeRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 房间ID
	RoomId *int64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 用户ID
	UserId *int64 `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 字符串类型用户ID
	UserIdStr *string `json:"UserIdStr,omitnil,omitempty" name:"UserIdStr"`

	// 字符串类型房间ID
	RoomIdStr *string `json:"RoomIdStr,omitnil,omitempty" name:"RoomIdStr"`
}

func (r *DescribeUserInAndOutTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInAndOutTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "RoomId")
	delete(f, "UserId")
	delete(f, "UserIdStr")
	delete(f, "RoomIdStr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserInAndOutTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInAndOutTimeResponseParams struct {
	// 用户在房间得进出时间列表
	InOutList []*InOutTimeInfo `json:"InOutList,omitnil,omitempty" name:"InOutList"`

	// 用户在房间中总时长
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserInAndOutTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserInAndOutTimeResponseParams `json:"Response"`
}

func (r *DescribeUserInAndOutTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInAndOutTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoicePrintRequestParams struct {
	// 查询方式，0表示查询特定VoicePrintId，1表示分页查询
	DescribeMode *uint64 `json:"DescribeMode,omitnil,omitempty" name:"DescribeMode"`

	// 声纹ID
	VoicePrintIdList []*string `json:"VoicePrintIdList,omitnil,omitempty" name:"VoicePrintIdList"`

	// 当前页码,从1开始,DescribeMode为1时填写
	PageIndex *uint64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 每页条数 最少20,DescribeMode为1时填写
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeVoicePrintRequest struct {
	*tchttp.BaseRequest
	
	// 查询方式，0表示查询特定VoicePrintId，1表示分页查询
	DescribeMode *uint64 `json:"DescribeMode,omitnil,omitempty" name:"DescribeMode"`

	// 声纹ID
	VoicePrintIdList []*string `json:"VoicePrintIdList,omitnil,omitempty" name:"VoicePrintIdList"`

	// 当前页码,从1开始,DescribeMode为1时填写
	PageIndex *uint64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 每页条数 最少20,DescribeMode为1时填写
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeVoicePrintRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVoicePrintRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DescribeMode")
	delete(f, "VoicePrintIdList")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVoicePrintRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoicePrintResponseParams struct {
	// 总的条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 声纹信息
	Data []*VoicePrintInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVoicePrintResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVoicePrintResponseParams `json:"Response"`
}

func (r *DescribeVoicePrintResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVoicePrintResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 要过滤的字段名, 比如"AppName"
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 多个关键字
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type GetCustomizationListRequestParams struct {
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`
}

type GetCustomizationListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`
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
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCustomizationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCustomizationListResponseParams struct {
	// 语音消息转文本热句模型配置
	CustomizationConfigs []*CustomizationConfigs `json:"CustomizationConfigs,omitnil,omitempty" name:"CustomizationConfigs"`

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

type InOutTimeInfo struct {
	// 进入房间时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 退出房间时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type InvokeLLM struct {
	// 请求LLM的内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 是否允许该文本打断机器人说话
	Interrupt *bool `json:"Interrupt,omitnil,omitempty" name:"Interrupt"`
}

// Predefined struct for user
type ModifyAppStatusRequestParams struct {
	// 应用ID，创建应用后由后台生成并返回。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 应用状态，取值：open/close
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyAppStatusRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，创建应用后由后台生成并返回。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 应用状态，取值：open/close
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyAppStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppStatusResp struct {
	// GME应用ID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 应用状态，取值：open/close
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type ModifyAppStatusResponseParams struct {
	// 修改应用开关状态返回数据
	Data *ModifyAppStatusResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAppStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAppStatusResponseParams `json:"Response"`
}

func (r *ModifyAppStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizationRequestParams struct {
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 文本文件的下载地址，服务会从该地址下载文件，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitnil,omitempty" name:"TextUrl"`

	// 修改的模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

type ModifyCustomizationRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 文本文件的下载地址，服务会从该地址下载文件，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitnil,omitempty" name:"TextUrl"`

	// 修改的模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
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
	delete(f, "BizId")
	delete(f, "TextUrl")
	delete(f, "ModelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizationResponseParams struct {
	// 返回值。0为成功，非0为失败。
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

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
	// 模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 想要变换的模型状态，-1代表下线，1代表上线
	ToState *int64 `json:"ToState,omitnil,omitempty" name:"ToState"`

	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`
}

type ModifyCustomizationStateRequest struct {
	*tchttp.BaseRequest
	
	// 模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 想要变换的模型状态，-1代表下线，1代表上线
	ToState *int64 `json:"ToState,omitnil,omitempty" name:"ToState"`

	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`
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
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomizationStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizationStateResponseParams struct {
	// 模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 返回值。0为成功，非0为失败。
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

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

// Predefined struct for user
type ModifyRecordInfoRequestParams struct {
	// 进行中的任务taskid（StartRecord接口返回）。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 录制类型：1代表单流 2代表混流 3代表单流和混流。
	RecordMode *uint64 `json:"RecordMode,omitnil,omitempty" name:"RecordMode"`

	// 应用ID。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 指定订阅流白名单或者黑名单。
	SubscribeRecordUserIds *SubscribeRecordUserIds `json:"SubscribeRecordUserIds,omitnil,omitempty" name:"SubscribeRecordUserIds"`
}

type ModifyRecordInfoRequest struct {
	*tchttp.BaseRequest
	
	// 进行中的任务taskid（StartRecord接口返回）。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 录制类型：1代表单流 2代表混流 3代表单流和混流。
	RecordMode *uint64 `json:"RecordMode,omitnil,omitempty" name:"RecordMode"`

	// 应用ID。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 指定订阅流白名单或者黑名单。
	SubscribeRecordUserIds *SubscribeRecordUserIds `json:"SubscribeRecordUserIds,omitnil,omitempty" name:"SubscribeRecordUserIds"`
}

func (r *ModifyRecordInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "RecordMode")
	delete(f, "BizId")
	delete(f, "SubscribeRecordUserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRecordInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRecordInfoResponseParams `json:"Response"`
}

func (r *ModifyRecordInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserMicStatusRequestParams struct {
	// 来自 [腾讯云控制台](https://console.cloud.tencent.com/gamegme) 的 GME 服务提供的 AppID，获取请参考 [语音服务开通指引](https://cloud.tencent.com/document/product/607/10782#.E9.87.8D.E7.82.B9.E5.8F.82.E6.95.B0)。
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 实时语音房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 需要操作的房间内用户以及该用户的目标麦克风状态。
	Users []*UserMicStatus `json:"Users,omitnil,omitempty" name:"Users"`
}

type ModifyUserMicStatusRequest struct {
	*tchttp.BaseRequest
	
	// 来自 [腾讯云控制台](https://console.cloud.tencent.com/gamegme) 的 GME 服务提供的 AppID，获取请参考 [语音服务开通指引](https://cloud.tencent.com/document/product/607/10782#.E9.87.8D.E7.82.B9.E5.8F.82.E6.95.B0)。
	BizId *int64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 实时语音房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 需要操作的房间内用户以及该用户的目标麦克风状态。
	Users []*UserMicStatus `json:"Users,omitnil,omitempty" name:"Users"`
}

func (r *ModifyUserMicStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserMicStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "RoomId")
	delete(f, "Users")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserMicStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserMicStatusResponseParams struct {
	// 返回结果：0为成功，非0为失败。
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 错误信息。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserMicStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserMicStatusResponseParams `json:"Response"`
}

func (r *ModifyUserMicStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserMicStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverseaTextStatisticsItem struct {
	// 统计值，单位：秒
	Data *float64 `json:"Data,omitnil,omitempty" name:"Data"`
}

type RealTimeSpeechStatisticsItem struct {
	// 大陆地区DAU
	MainLandDau *uint64 `json:"MainLandDau,omitnil,omitempty" name:"MainLandDau"`

	// 大陆地区PCU
	MainLandPcu *uint64 `json:"MainLandPcu,omitnil,omitempty" name:"MainLandPcu"`

	// 大陆地区总使用时长，单位为min
	MainLandDuration *uint64 `json:"MainLandDuration,omitnil,omitempty" name:"MainLandDuration"`

	// 海外地区DAU
	OverseaDau *uint64 `json:"OverseaDau,omitnil,omitempty" name:"OverseaDau"`

	// 海外地区PCU
	OverseaPcu *uint64 `json:"OverseaPcu,omitnil,omitempty" name:"OverseaPcu"`

	// 海外地区总使用时长，单位为min
	OverseaDuration *uint64 `json:"OverseaDuration,omitnil,omitempty" name:"OverseaDuration"`
}

type RealtimeSpeechConf struct {
	// 实时语音服务开关，取值：open/close
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实时语音音质类型，取值：high-高音质 ordinary-普通音质
	Quality *string `json:"Quality,omitnil,omitempty" name:"Quality"`
}

type RealtimeTextStatisticsItem struct {
	// 统计值，单位：秒
	Data *float64 `json:"Data,omitnil,omitempty" name:"Data"`
}

type RecordInfo struct {
	// 用户ID（当混流模式时，取值为0）。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 录制文件名。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 录制开始时间（unix时间戳如：1234567868）。
	RecordBeginTime *uint64 `json:"RecordBeginTime,omitnil,omitempty" name:"RecordBeginTime"`

	// 录制状态：2代表正在录制  10代表等待转码  11代表正在转码  12正在上传  13代表上传完成  14代表通知用户完成。
	RecordStatus *uint64 `json:"RecordStatus,omitnil,omitempty" name:"RecordStatus"`
}

// Predefined struct for user
type RegisterVoicePrintRequestParams struct {
	// 整个wav音频文件的base64字符串,其中wav文件限定为16k采样率, 16bit位深, 单声道, 4到18秒音频时长,有效音频不小于3秒(不能有太多静音段), 编码数据大小不超过2M, 为了识别准确率，建议音频长度为8秒
	Audio *string `json:"Audio,omitnil,omitempty" name:"Audio"`

	// 毫秒时间戳
	ReqTimestamp *uint64 `json:"ReqTimestamp,omitnil,omitempty" name:"ReqTimestamp"`

	// 音频格式,目前只支持0,代表wav
	AudioFormat *uint64 `json:"AudioFormat,omitnil,omitempty" name:"AudioFormat"`

	// 音频名称,长度不要超过32
	AudioName *string `json:"AudioName,omitnil,omitempty" name:"AudioName"`

	// 和声纹绑定的MetaInfo，长度最大不超过512
	AudioMetaInfo *string `json:"AudioMetaInfo,omitnil,omitempty" name:"AudioMetaInfo"`
}

type RegisterVoicePrintRequest struct {
	*tchttp.BaseRequest
	
	// 整个wav音频文件的base64字符串,其中wav文件限定为16k采样率, 16bit位深, 单声道, 4到18秒音频时长,有效音频不小于3秒(不能有太多静音段), 编码数据大小不超过2M, 为了识别准确率，建议音频长度为8秒
	Audio *string `json:"Audio,omitnil,omitempty" name:"Audio"`

	// 毫秒时间戳
	ReqTimestamp *uint64 `json:"ReqTimestamp,omitnil,omitempty" name:"ReqTimestamp"`

	// 音频格式,目前只支持0,代表wav
	AudioFormat *uint64 `json:"AudioFormat,omitnil,omitempty" name:"AudioFormat"`

	// 音频名称,长度不要超过32
	AudioName *string `json:"AudioName,omitnil,omitempty" name:"AudioName"`

	// 和声纹绑定的MetaInfo，长度最大不超过512
	AudioMetaInfo *string `json:"AudioMetaInfo,omitnil,omitempty" name:"AudioMetaInfo"`
}

func (r *RegisterVoicePrintRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterVoicePrintRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Audio")
	delete(f, "ReqTimestamp")
	delete(f, "AudioFormat")
	delete(f, "AudioName")
	delete(f, "AudioMetaInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterVoicePrintRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterVoicePrintResponseParams struct {
	// 声纹信息ID
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterVoicePrintResponse struct {
	*tchttp.BaseResponse
	Response *RegisterVoicePrintResponseParams `json:"Response"`
}

func (r *RegisterVoicePrintResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterVoicePrintResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoomUser struct {
	// 房间id
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 房间里用户uin列表
	Uins []*uint64 `json:"Uins,omitnil,omitempty" name:"Uins"`

	// 字符串房间id
	StrRoomId *string `json:"StrRoomId,omitnil,omitempty" name:"StrRoomId"`

	// 房间里用户字符串uin列表
	StrUins []*string `json:"StrUins,omitnil,omitempty" name:"StrUins"`
}

type STTConfig struct {
	// 
	// 语音转文字支持识别的语言，默认是"zh" 中文
	// 
	// 可通过购买「AI智能识别时长包」解锁或领取包月套餐体验版解锁不同语言. 
	// 
	// 语音转文本不同套餐版本支持的语言如下：
	// 
	// **基础版**：
	// - "zh": 中文（简体）
	// - "zh-TW": 中文（繁体）
	// - "en": 英语
	// 
	// **标准版：**
	// - "8k_zh_large": 普方大模型引擎. 当前模型同时支持中文等语言的识别，模型参数量极大，语言模型性能增强，针对电话音频中各类场景、各类中文方言的识别准确率极大提升.
	// - "16k_zh_large": 普方英大模型引擎. 当前模型同时支持中文、英文、多种中文方言等语言的识别，模型参数量极大，语言模型性能增强，针对噪声大、回音大、人声小、人声远等低质量音频的识别准确率极大提升.
	// - "16k_multi_lang": 多语种大模型引擎. 当前模型同时支持英语、日语、韩语、阿拉伯语、菲律宾语、法语、印地语、印尼语、马来语、葡萄牙语、西班牙语、泰语、土耳其语、越南语、德语的识别，可实现15个语种的自动识别(句子/段落级别).
	// - "16k_zh_en": 中英大模型引擎. 当前模型同时支持中文、英语识别，模型参数量极大，语言模型性能增强，针对噪声大、回音大、人声小、人声远等低质量音频的识别准确率极大提升.
	// 
	// **高级版：**
	// - "zh-dialect": 中国方言
	// - "zh-yue": 中国粤语
	// - "vi": 越南语
	// - "ja": 日语
	// - "ko": 韩语
	// - "id": 印度尼西亚语
	// - "th": 泰语
	// - "pt": 葡萄牙语
	// - "tr": 土耳其语
	// - "ar": 阿拉伯语
	// - "es": 西班牙语
	// - "hi": 印地语
	// - "fr": 法语
	// - "ms": 马来语
	// - "fil": 菲律宾语
	// - "de": 德语
	// - "it": 意大利语
	// - "ru": 俄语
	// - "sv": 瑞典语
	// - "da": 丹麦语
	// - "no": 挪威语
	// 
	// **注意：**
	// 如果缺少满足您需求的语言，请联系我们技术人员。
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// **发起模糊识别为高级版能力,默认按照高级版收费,仅支持填写基础版和高级版语言.**
	// 注意：不支持填写"zh-dialect"
	AlternativeLanguage []*string `json:"AlternativeLanguage,omitnil,omitempty" name:"AlternativeLanguage"`

	// 自定义参数，联系后台使用
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// 语音识别vad的时间，范围为240-2000，默认为1000，单位为ms。更小的值会让语音识别分句更快。
	VadSilenceTime *uint64 `json:"VadSilenceTime,omitnil,omitempty" name:"VadSilenceTime"`

	// 热词表：该参数用于提升识别准确率。 单个热词限制："热词|权重"，单个热词不超过30个字符（最多10个汉字），权重[1-11]或者100，如：“腾讯云|5” 或 “ASR|11”； 热词表限制：多个热词用英文逗号分割，最多支持128个热词，如：“腾讯云|10,语音识别|5,ASR|11”；
	HotWordList *string `json:"HotWordList,omitnil,omitempty" name:"HotWordList"`

	// vad的远场人声抑制能力（不会对asr识别效果造成影响），范围为[0, 3]，默认为0。推荐设置为2，有较好的远场人声抑制能力。	
	VadLevel *uint64 `json:"VadLevel,omitnil,omitempty" name:"VadLevel"`
}

type ScanDetail struct {
	// 违规场景，参照<a href="https://cloud.tencent.com/document/product/607/37622#Label_Value">Label</a>定义
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该场景下概率[0.00,100.00],分值越大违规概率越高
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 违规关键字
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// 关键字在音频的开始时间，从0开始的偏移量，单位为毫秒，Label=moan时有效
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 关键字在音频的结束时间，从0开始的偏移量,，单位为毫秒，Label=moan时有效
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ScanPiece struct {
	// 流检测时返回，音频转存地址，保留30min
	DumpUrl *string `json:"DumpUrl,omitnil,omitempty" name:"DumpUrl"`

	// 是否违规
	HitFlag *bool `json:"HitFlag,omitnil,omitempty" name:"HitFlag"`

	// 违规主要类型
	MainType *string `json:"MainType,omitnil,omitempty" name:"MainType"`

	// 语音检测详情
	ScanDetail []*ScanDetail `json:"ScanDetail,omitnil,omitempty" name:"ScanDetail"`

	// gme实时语音房间ID，透传任务传入时的RoomId
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// gme实时语音用户ID，透传任务传入时的OpenId
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 备注
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 流检测时分片在流中的偏移时间，单位毫秒
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 流检测时分片时长
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 分片开始检测时间
	PieceStartTime *uint64 `json:"PieceStartTime,omitnil,omitempty" name:"PieceStartTime"`
}

// Predefined struct for user
type ScanVoiceRequestParams struct {
	// 应用ID，登录[控制台 - 服务管理](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 语音检测场景，参数值目前要求为 default。 预留场景设置： 谩骂、色情、广告、违禁等场景，<a href="#Label_Value">具体取值见上述 Label 说明。</a>
	Scenes []*string `json:"Scenes,omitnil,omitempty" name:"Scenes"`

	// 是否为直播流。值为 false 时表示普通语音文件检测；为 true 时表示语音流检测。
	Live *bool `json:"Live,omitnil,omitempty" name:"Live"`

	// 语音检测任务列表，列表最多支持100个检测任务。结构体中包含：
	// <li>DataId：数据的唯一ID</li>
	// <li>Url：数据文件的url，为 urlencode 编码，流式则为拉流地址</li>
	Tasks []*Task `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 异步检测结果回调地址，具体见上述<a href="#Callback_Declare">回调相关说明</a>。（说明：该字段为空时，必须通过接口(查询语音检测结果)获取检测结果）。
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 语种，不传默认中文
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`
}

type ScanVoiceRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，登录[控制台 - 服务管理](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 语音检测场景，参数值目前要求为 default。 预留场景设置： 谩骂、色情、广告、违禁等场景，<a href="#Label_Value">具体取值见上述 Label 说明。</a>
	Scenes []*string `json:"Scenes,omitnil,omitempty" name:"Scenes"`

	// 是否为直播流。值为 false 时表示普通语音文件检测；为 true 时表示语音流检测。
	Live *bool `json:"Live,omitnil,omitempty" name:"Live"`

	// 语音检测任务列表，列表最多支持100个检测任务。结构体中包含：
	// <li>DataId：数据的唯一ID</li>
	// <li>Url：数据文件的url，为 urlencode 编码，流式则为拉流地址</li>
	Tasks []*Task `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 异步检测结果回调地址，具体见上述<a href="#Callback_Declare">回调相关说明</a>。（说明：该字段为空时，必须通过接口(查询语音检测结果)获取检测结果）。
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 语种，不传默认中文
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`
}

func (r *ScanVoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVoiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "Scenes")
	delete(f, "Live")
	delete(f, "Tasks")
	delete(f, "Callback")
	delete(f, "Lang")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanVoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanVoiceResponseParams struct {
	// 语音检测返回。Data 字段是 JSON 数组，每一个元素包含：<li>DataId： 请求中对应的 DataId。</li>
	// <li>TaskID ：该检测任务的 ID，用于轮询语音检测结果。</li>
	Data []*ScanVoiceResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScanVoiceResponse struct {
	*tchttp.BaseResponse
	Response *ScanVoiceResponseParams `json:"Response"`
}

func (r *ScanVoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanVoiceResult struct {
	// 数据ID
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type SceneInfo struct {
	// 'RealTime','实时语音分析',
	// 'VoiceMessage','语音消息',
	// 'GMECloudApi':'GME云API接口'
	SceneId *string `json:"SceneId,omitnil,omitempty" name:"SceneId"`

	// 开关状态，true开启/false关闭
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 用户回调地址
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`
}

type ServerPushText struct {
	// 服务端推送播报文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 是否允许该文本打断机器人说话
	Interrupt *bool `json:"Interrupt,omitnil,omitempty" name:"Interrupt"`

	// 播报完文本后，是否自动关闭对话任务
	StopAfterPlay *bool `json:"StopAfterPlay,omitnil,omitempty" name:"StopAfterPlay"`

	// 服务端推送播报音频
	//     格式说明：音频必须为单声道，采样率必须跟对应TTS的采样率保持一致，编码为Base64字符串。
	//     输入规则：当提供Audio字段时，将不接受Text字段的输入。系统将直接播放Audio字段中的音频内容。
	Audio *string `json:"Audio,omitnil,omitempty" name:"Audio"`

	// 默认为0，仅在Interrupt为false时有效
	// - 0表示当前有交互发生时，会丢弃Interrupt为false的消息
	// - 1表示当前有交互发生时，不会丢弃Interrupt为false的消息，而是缓存下来，等待当前交互结束后，再去处理
	// 
	// 注意：DropMode为1时，允许缓存多个消息，如果后续出现了打断，缓存的消息会被清空
	DropMode *uint64 `json:"DropMode,omitnil,omitempty" name:"DropMode"`

	// ServerPushText消息的优先级，0表示可被打断，1表示不会被打断。**目前仅支持传入0，如果需要传入1，请提工单联系我们添加权限。**
	// 注意：在接收到Priority=1的消息后，后续其他任何消息都会被忽略（包括Priority=1的消息），直到Priority=1的消息处理结束。该字段可与Interrupt、DropMode字段配合使用。
	// 例子：
	// - Priority=1、Interrupt=true，会打断现有交互，立刻播报，播报过程中不会被打断
	// - Priority=1、Interrupt=false、DropMode=1，会等待当前交互结束，再进行播报，播报过程中不会被打断
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type ServiceStatus struct {
	// 实时语音服务开关状态
	RealTimeSpeech *StatusInfo `json:"RealTimeSpeech,omitnil,omitempty" name:"RealTimeSpeech"`

	// 语音消息服务开关状态
	VoiceMessage *StatusInfo `json:"VoiceMessage,omitnil,omitempty" name:"VoiceMessage"`

	// 语音内容安全服务开关状态
	Porn *StatusInfo `json:"Porn,omitnil,omitempty" name:"Porn"`

	// 语音录制服务开关状态
	Live *StatusInfo `json:"Live,omitnil,omitempty" name:"Live"`

	// 语音转文本服务开关状态
	RealTimeAsr *StatusInfo `json:"RealTimeAsr,omitnil,omitempty" name:"RealTimeAsr"`

	// 文本翻译服务开关状态
	TextTranslate *StatusInfo `json:"TextTranslate,omitnil,omitempty" name:"TextTranslate"`
}

// Predefined struct for user
type StartAIConversationRequestParams struct {
	// GME的SdkAppId和开启转录任务的房间使用的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// GME的RoomId表示开启对话任务的房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 机器人参数
	AgentConfig *AgentConfig `json:"AgentConfig,omitnil,omitempty" name:"AgentConfig"`

	// 语音识别配置。
	STTConfig *STTConfig `json:"STTConfig,omitnil,omitempty" name:"STTConfig"`

	// LLM配置。需符合openai规范，为JSON字符串，示例如下：
	// <pre> { <br> &emsp;  "LLMType": "大模型类型",  // String 必填，如："openai" <br> &emsp;  "Model": "您的模型名称", // String 必填，指定使用的模型<br>    "APIKey": "您的LLM API密钥", // String 必填 <br> &emsp;  "APIUrl": "https://api.xxx.com/chat/completions", // String 必填，LLM API访问的URL<br> &emsp;  "Streaming": true // Boolean 非必填，指定是否使用流式传输<br> &emsp;} </pre>
	LLMConfig *string `json:"LLMConfig,omitnil,omitempty" name:"LLMConfig"`

	//                                         "description": "TTS配置，为JSON字符串，腾讯云TTS示例如下： <pre>{ <br> &emsp; \"AppId\": 您的应用ID, // Integer 必填<br> &emsp; \"TTSType\": \"TTS类型\", // String TTS类型, 固定为\"tencent\"<br> &emsp; \"SecretId\": \"您的密钥ID\", // String 必填<br> &emsp; \"SecretKey\":  \"您的密钥Key\", // String 必填<br> &emsp; \"VoiceType\": 101001, // Integer  必填，音色 ID，包括标准音色与精品音色，精品音色拟真度更高，价格不同于标准音色。<br> &emsp; \"Speed\": 1.25, // Integer 非必填，语速，范围：[-2，6]，分别对应不同语速： -2: 代表0.6倍 -1: 代表0.8倍 0: 代表1.0倍（默认） 1: 代表1.2倍 2: 代表1.5倍  6: 代表2.5倍  如果需要更细化的语速，可以保留小数点后 2 位，例如0.5/1.25/2.81等。 参数值与实际语速转换\"Volume\": 5, // Integer 非必填，音量大小，范围：[0，10]，分别对应11个等级的音量，默认值为0，代表正常音量。<br> &emsp; \"EmotionCategory\":  \"angry\", // String 非必填 控制合成音频的情感，仅支持多情感音色使用。取值: neutral(中性)、sad(悲伤)、happy(高兴)、angry(生气)、fear(恐惧)、news(新闻)、story(故事)、radio(广播)、poetry(诗歌)、call(客服)、sajiao(撒娇)、disgusted(厌恶)、amaze(震惊)、peaceful(平静)、exciting(兴奋)、aojiao(傲娇)、jieshuo(解说)。<br> &emsp; \"EmotionIntensity\":  150 // Integer 非必填 控制合成音频情感程度，取值范围为 [50,200]，默认为 100；只有 EmotionCategory 不为空时生效。<br> &emsp; }</pre>",
	TTSConfig *string `json:"TTSConfig,omitnil,omitempty" name:"TTSConfig"`

	// 数字人配置，为JSON字符串。**数字人配置需要提工单加白后才能使用**
	AvatarConfig *string `json:"AvatarConfig,omitnil,omitempty" name:"AvatarConfig"`

	// 实验性参数,联系后台使用
	ExperimentalParams *string `json:"ExperimentalParams,omitnil,omitempty" name:"ExperimentalParams"`
}

type StartAIConversationRequest struct {
	*tchttp.BaseRequest
	
	// GME的SdkAppId和开启转录任务的房间使用的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// GME的RoomId表示开启对话任务的房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 机器人参数
	AgentConfig *AgentConfig `json:"AgentConfig,omitnil,omitempty" name:"AgentConfig"`

	// 语音识别配置。
	STTConfig *STTConfig `json:"STTConfig,omitnil,omitempty" name:"STTConfig"`

	// LLM配置。需符合openai规范，为JSON字符串，示例如下：
	// <pre> { <br> &emsp;  "LLMType": "大模型类型",  // String 必填，如："openai" <br> &emsp;  "Model": "您的模型名称", // String 必填，指定使用的模型<br>    "APIKey": "您的LLM API密钥", // String 必填 <br> &emsp;  "APIUrl": "https://api.xxx.com/chat/completions", // String 必填，LLM API访问的URL<br> &emsp;  "Streaming": true // Boolean 非必填，指定是否使用流式传输<br> &emsp;} </pre>
	LLMConfig *string `json:"LLMConfig,omitnil,omitempty" name:"LLMConfig"`

	//                                         "description": "TTS配置，为JSON字符串，腾讯云TTS示例如下： <pre>{ <br> &emsp; \"AppId\": 您的应用ID, // Integer 必填<br> &emsp; \"TTSType\": \"TTS类型\", // String TTS类型, 固定为\"tencent\"<br> &emsp; \"SecretId\": \"您的密钥ID\", // String 必填<br> &emsp; \"SecretKey\":  \"您的密钥Key\", // String 必填<br> &emsp; \"VoiceType\": 101001, // Integer  必填，音色 ID，包括标准音色与精品音色，精品音色拟真度更高，价格不同于标准音色。<br> &emsp; \"Speed\": 1.25, // Integer 非必填，语速，范围：[-2，6]，分别对应不同语速： -2: 代表0.6倍 -1: 代表0.8倍 0: 代表1.0倍（默认） 1: 代表1.2倍 2: 代表1.5倍  6: 代表2.5倍  如果需要更细化的语速，可以保留小数点后 2 位，例如0.5/1.25/2.81等。 参数值与实际语速转换\"Volume\": 5, // Integer 非必填，音量大小，范围：[0，10]，分别对应11个等级的音量，默认值为0，代表正常音量。<br> &emsp; \"EmotionCategory\":  \"angry\", // String 非必填 控制合成音频的情感，仅支持多情感音色使用。取值: neutral(中性)、sad(悲伤)、happy(高兴)、angry(生气)、fear(恐惧)、news(新闻)、story(故事)、radio(广播)、poetry(诗歌)、call(客服)、sajiao(撒娇)、disgusted(厌恶)、amaze(震惊)、peaceful(平静)、exciting(兴奋)、aojiao(傲娇)、jieshuo(解说)。<br> &emsp; \"EmotionIntensity\":  150 // Integer 非必填 控制合成音频情感程度，取值范围为 [50,200]，默认为 100；只有 EmotionCategory 不为空时生效。<br> &emsp; }</pre>",
	TTSConfig *string `json:"TTSConfig,omitnil,omitempty" name:"TTSConfig"`

	// 数字人配置，为JSON字符串。**数字人配置需要提工单加白后才能使用**
	AvatarConfig *string `json:"AvatarConfig,omitnil,omitempty" name:"AvatarConfig"`

	// 实验性参数,联系后台使用
	ExperimentalParams *string `json:"ExperimentalParams,omitnil,omitempty" name:"ExperimentalParams"`
}

func (r *StartAIConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAIConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "AgentConfig")
	delete(f, "STTConfig")
	delete(f, "LLMConfig")
	delete(f, "TTSConfig")
	delete(f, "AvatarConfig")
	delete(f, "ExperimentalParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartAIConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartAIConversationResponseParams struct {
	// 用于唯一标识对话任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartAIConversationResponse struct {
	*tchttp.BaseResponse
	Response *StartAIConversationResponseParams `json:"Response"`
}

func (r *StartAIConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAIConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartRecordRequestParams struct {
	// 应用ID。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 房间ID。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 录制类型：1代表单流 2代表混流 3代表单流和混流。
	RecordMode *uint64 `json:"RecordMode,omitnil,omitempty" name:"RecordMode"`

	// 指定订阅流白名单或者黑名单（不传默认订阅房间内所有音频流）。
	SubscribeRecordUserIds *SubscribeRecordUserIds `json:"SubscribeRecordUserIds,omitnil,omitempty" name:"SubscribeRecordUserIds"`
}

type StartRecordRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 房间ID。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 录制类型：1代表单流 2代表混流 3代表单流和混流。
	RecordMode *uint64 `json:"RecordMode,omitnil,omitempty" name:"RecordMode"`

	// 指定订阅流白名单或者黑名单（不传默认订阅房间内所有音频流）。
	SubscribeRecordUserIds *SubscribeRecordUserIds `json:"SubscribeRecordUserIds,omitnil,omitempty" name:"SubscribeRecordUserIds"`
}

func (r *StartRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "RoomId")
	delete(f, "RecordMode")
	delete(f, "SubscribeRecordUserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartRecordResponseParams struct {
	// 任务taskid。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartRecordResponse struct {
	*tchttp.BaseResponse
	Response *StartRecordResponseParams `json:"Response"`
}

func (r *StartRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatisticsItem struct {
	// 日期，格式为年-月-日，如2018-07-13
	StatDate *string `json:"StatDate,omitnil,omitempty" name:"StatDate"`

	// 统计值
	Data *uint64 `json:"Data,omitnil,omitempty" name:"Data"`
}

type StatusInfo struct {
	// 服务开关状态， 0-正常，1-关闭
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type StopAIConversationRequestParams struct {
	// 唯一标识任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopAIConversationRequest struct {
	*tchttp.BaseRequest
	
	// 唯一标识任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopAIConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAIConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopAIConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopAIConversationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopAIConversationResponse struct {
	*tchttp.BaseResponse
	Response *StopAIConversationResponseParams `json:"Response"`
}

func (r *StopAIConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAIConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRecordRequestParams struct {
	// 任务ID。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 应用ID。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`
}

type StopRecordRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 应用ID。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`
}

func (r *StopRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopRecordResponse struct {
	*tchttp.BaseResponse
	Response *StopRecordResponseParams `json:"Response"`
}

func (r *StopRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StreamTextStatisticsItem struct {
	// 统计值，单位：秒
	Data *float64 `json:"Data,omitnil,omitempty" name:"Data"`
}

type SubscribeRecordUserIds struct {
	// 订阅音频流黑名单，指定不订阅哪几个UserId的音频流，例如["1", "2", "3"], 代表不订阅UserId 1，2，3的音频流。默认不填订阅房间内所有音频流，订阅列表用户数不超过20。
	// 注意：只能同时设置UnSubscribeAudioUserIds、SubscribeAudioUserIds 其中1个参数
	UnSubscribeUserIds []*string `json:"UnSubscribeUserIds,omitnil,omitempty" name:"UnSubscribeUserIds"`

	// 订阅音频流白名单，指定订阅哪几个UserId的音频流，例如["1", "2", "3"], 代表订阅UserId 1，2，3的音频流。默认不填订阅房间内所有音频流，订阅列表用户数不超过20。
	// 注意：只能同时设置UnSubscribeAudioUserIds、SubscribeAudioUserIds 其中1个参数。
	SubscribeUserIds []*string `json:"SubscribeUserIds,omitnil,omitempty" name:"SubscribeUserIds"`
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type Task struct {
	// 数据的唯一ID
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 数据文件的url，为 urlencode 编码，流式则为拉流地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// gme实时语音房间ID，通过gme实时语音进行语音分析时输入
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// gme实时语音用户ID，通过gme实时语音进行语音分析时输入
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`
}

type TurnDetection struct {
	// TurnDetectionMode为3时生效，语义断句的灵敏程度
	// 
	// 
	// 功能简介：根据用户所说的话来判断其已完成发言来分割音频
	// 
	// 
	// 可选: "low" | "medium" | "high" | "auto"
	// 
	// 
	// auto 是默认值，与 medium 相同。
	// low 将让用户有足够的时间说话。
	// high 将尽快对音频进行分块。
	// 
	// 
	// 如果您希望模型在对话模式下更频繁地响应，可以将 SemanticEagerness 设置为 high
	// 如果您希望在用户停顿时，AI能够等待片刻，可以将 SemanticEagerness 设置为 low
	// 无论什么模式，最终都会分割送个大模型进行回复
	SemanticEagerness *string `json:"SemanticEagerness,omitnil,omitempty" name:"SemanticEagerness"`
}

// Predefined struct for user
type UpdateAIConversationRequestParams struct {
	// 唯一标识一个任务
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 不填写则不进行更新，机器人的欢迎语
	WelcomeMessage *string `json:"WelcomeMessage,omitnil,omitempty" name:"WelcomeMessage"`

	// 不填写则不进行更新。智能打断模式，0表示服务端自动打断，1表示服务端不打断，由端上发送打断信令进行打断
	InterruptMode *uint64 `json:"InterruptMode,omitnil,omitempty" name:"InterruptMode"`

	// 不填写则不进行更新。InterruptMode为0时使用，单位为毫秒，默认为500ms。表示服务端检测到持续InterruptSpeechDuration毫秒的人声则进行打断
	InterruptSpeechDuration *uint64 `json:"InterruptSpeechDuration,omitnil,omitempty" name:"InterruptSpeechDuration"`

	// 不填写则不进行更新，LLM配置，详情见StartAIConversation接口
	LLMConfig *string `json:"LLMConfig,omitnil,omitempty" name:"LLMConfig"`

	// 不填写则不进行更新，TTS配置，详情见StartAIConversation接口
	TTSConfig *string `json:"TTSConfig,omitnil,omitempty" name:"TTSConfig"`
}

type UpdateAIConversationRequest struct {
	*tchttp.BaseRequest
	
	// 唯一标识一个任务
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 不填写则不进行更新，机器人的欢迎语
	WelcomeMessage *string `json:"WelcomeMessage,omitnil,omitempty" name:"WelcomeMessage"`

	// 不填写则不进行更新。智能打断模式，0表示服务端自动打断，1表示服务端不打断，由端上发送打断信令进行打断
	InterruptMode *uint64 `json:"InterruptMode,omitnil,omitempty" name:"InterruptMode"`

	// 不填写则不进行更新。InterruptMode为0时使用，单位为毫秒，默认为500ms。表示服务端检测到持续InterruptSpeechDuration毫秒的人声则进行打断
	InterruptSpeechDuration *uint64 `json:"InterruptSpeechDuration,omitnil,omitempty" name:"InterruptSpeechDuration"`

	// 不填写则不进行更新，LLM配置，详情见StartAIConversation接口
	LLMConfig *string `json:"LLMConfig,omitnil,omitempty" name:"LLMConfig"`

	// 不填写则不进行更新，TTS配置，详情见StartAIConversation接口
	TTSConfig *string `json:"TTSConfig,omitnil,omitempty" name:"TTSConfig"`
}

func (r *UpdateAIConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAIConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "WelcomeMessage")
	delete(f, "InterruptMode")
	delete(f, "InterruptSpeechDuration")
	delete(f, "LLMConfig")
	delete(f, "TTSConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAIConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAIConversationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAIConversationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAIConversationResponseParams `json:"Response"`
}

func (r *UpdateAIConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAIConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateScanRoomsRequestParams struct {
	// 应用ID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 需要送检的所有房间号。多个房间号之间用","分隔，长度不超过1024字符。示例："0001,0002,0003"
	RoomIdString *string `json:"RoomIdString,omitnil,omitempty" name:"RoomIdString"`

	// 符合此正则表达式规则的房间号将被送检，最大不能超过10个。示例：^6.*（表示所有以6开头的房间号将被送检）
	RoomIdRegex []*string `json:"RoomIdRegex,omitnil,omitempty" name:"RoomIdRegex"`
}

type UpdateScanRoomsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 需要送检的所有房间号。多个房间号之间用","分隔，长度不超过1024字符。示例："0001,0002,0003"
	RoomIdString *string `json:"RoomIdString,omitnil,omitempty" name:"RoomIdString"`

	// 符合此正则表达式规则的房间号将被送检，最大不能超过10个。示例：^6.*（表示所有以6开头的房间号将被送检）
	RoomIdRegex []*string `json:"RoomIdRegex,omitnil,omitempty" name:"RoomIdRegex"`
}

func (r *UpdateScanRoomsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateScanRoomsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "RoomIdString")
	delete(f, "RoomIdRegex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateScanRoomsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateScanRoomsResponseParams struct {
	// 返回结果码
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateScanRoomsResponse struct {
	*tchttp.BaseResponse
	Response *UpdateScanRoomsResponseParams `json:"Response"`
}

func (r *UpdateScanRoomsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateScanRoomsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateScanUsersRequestParams struct {
	// 应用ID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 需要送检的所有用户号。多个用户号之间用","分隔，长度不超过1024字符。示例："0001,0002,0003"
	UserIdString *string `json:"UserIdString,omitnil,omitempty" name:"UserIdString"`

	// 符合此正则表达式规则的用户号将被送检，最大不能超过10个。示例：["^6.*"] 表示所有以6开头的用户号将被送检
	UserIdRegex []*string `json:"UserIdRegex,omitnil,omitempty" name:"UserIdRegex"`
}

type UpdateScanUsersRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 需要送检的所有用户号。多个用户号之间用","分隔，长度不超过1024字符。示例："0001,0002,0003"
	UserIdString *string `json:"UserIdString,omitnil,omitempty" name:"UserIdString"`

	// 符合此正则表达式规则的用户号将被送检，最大不能超过10个。示例：["^6.*"] 表示所有以6开头的用户号将被送检
	UserIdRegex []*string `json:"UserIdRegex,omitnil,omitempty" name:"UserIdRegex"`
}

func (r *UpdateScanUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateScanUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "UserIdString")
	delete(f, "UserIdRegex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateScanUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateScanUsersResponseParams struct {
	// 返回结果码
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateScanUsersResponse struct {
	*tchttp.BaseResponse
	Response *UpdateScanUsersResponseParams `json:"Response"`
}

func (r *UpdateScanUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateScanUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateVoicePrintRequestParams struct {
	// 声纹信息ID
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`

	// 毫秒时间戳
	ReqTimestamp *uint64 `json:"ReqTimestamp,omitnil,omitempty" name:"ReqTimestamp"`

	// 音频格式,目前只支持0,代表wav
	AudioFormat *uint64 `json:"AudioFormat,omitnil,omitempty" name:"AudioFormat"`

	// 整个wav音频文件的base64字符串,其中wav文件限定为16k采样率, 16bit位深, 单声道, 8到18秒音频时长,有效音频不小于6秒(不能有太多静音段),编码数据大小不超过2M
	Audio *string `json:"Audio,omitnil,omitempty" name:"Audio"`

	// 和声纹绑定的MetaInfo，长度最大不超过512
	AudioMetaInfo *string `json:"AudioMetaInfo,omitnil,omitempty" name:"AudioMetaInfo"`
}

type UpdateVoicePrintRequest struct {
	*tchttp.BaseRequest
	
	// 声纹信息ID
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`

	// 毫秒时间戳
	ReqTimestamp *uint64 `json:"ReqTimestamp,omitnil,omitempty" name:"ReqTimestamp"`

	// 音频格式,目前只支持0,代表wav
	AudioFormat *uint64 `json:"AudioFormat,omitnil,omitempty" name:"AudioFormat"`

	// 整个wav音频文件的base64字符串,其中wav文件限定为16k采样率, 16bit位深, 单声道, 8到18秒音频时长,有效音频不小于6秒(不能有太多静音段),编码数据大小不超过2M
	Audio *string `json:"Audio,omitnil,omitempty" name:"Audio"`

	// 和声纹绑定的MetaInfo，长度最大不超过512
	AudioMetaInfo *string `json:"AudioMetaInfo,omitnil,omitempty" name:"AudioMetaInfo"`
}

func (r *UpdateVoicePrintRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateVoicePrintRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VoicePrintId")
	delete(f, "ReqTimestamp")
	delete(f, "AudioFormat")
	delete(f, "Audio")
	delete(f, "AudioMetaInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateVoicePrintRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateVoicePrintResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateVoicePrintResponse struct {
	*tchttp.BaseResponse
	Response *UpdateVoicePrintResponseParams `json:"Response"`
}

func (r *UpdateVoicePrintResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateVoicePrintResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserMicStatus struct {
	// 开麦状态。1表示关闭麦克风，2表示打开麦克风。
	EnableMic *int64 `json:"EnableMic,omitnil,omitempty" name:"EnableMic"`

	// 客户端用于标识用户的Openid。（Uid、StrUid必须填一个，优先处理StrUid。）
	Uid *int64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 客户端用于标识字符串型用户的Openid。（Uid、StrUid必须填一个，优先处理StrUid。）
	StrUid *string `json:"StrUid,omitnil,omitempty" name:"StrUid"`
}

type VoiceFilterConf struct {
	// 语音过滤服务开关，取值：open/close
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 场景配置信息，如开关状态，回调地址。
	SceneInfos []*SceneInfo `json:"SceneInfos,omitnil,omitempty" name:"SceneInfos"`
}

type VoiceFilterStatisticsItem struct {
	// 语音过滤总时长，单位为min
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

type VoiceMessageConf struct {
	// 离线语音服务开关，取值：open/close
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 离线语音支持语种，取值： all-全部，cnen-中英文。默认为中英文
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`
}

type VoiceMessageStatisticsItem struct {
	// 离线语音DAU
	Dau *uint64 `json:"Dau,omitnil,omitempty" name:"Dau"`
}

type VoicePrint struct {
	// 默认为0，表示不启用声纹。1表示启用声纹，此时需要填写voiceprint id。
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// VoicePrint Mode为1时需要填写，目前仅支持填写一个声纹id
	IdList []*string `json:"IdList,omitnil,omitempty" name:"IdList"`
}

type VoicePrintInfo struct {
	// 声纹ID
	VoicePrintId *string `json:"VoicePrintId,omitnil,omitempty" name:"VoicePrintId"`

	// 应用id
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 和声纹绑定的MetaInfo
	VoicePrintMetaInfo *string `json:"VoicePrintMetaInfo,omitnil,omitempty" name:"VoicePrintMetaInfo"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 音频格式,当前只有0(代表wav)
	AudioFormat *uint64 `json:"AudioFormat,omitnil,omitempty" name:"AudioFormat"`

	// 音频名称
	AudioName *string `json:"AudioName,omitnil,omitempty" name:"AudioName"`

	// 请求毫秒时间戳
	ReqTimestamp *uint64 `json:"ReqTimestamp,omitnil,omitempty" name:"ReqTimestamp"`
}