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

package v20180711

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AppStatisticsItem struct {

	// 实时语音统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealtimeSpeechStatisticsItem *RealTimeSpeechStatisticsItem `json:"RealtimeSpeechStatisticsItem,omitempty" name:"RealtimeSpeechStatisticsItem"`

	// 语音消息统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceMessageStatisticsItem *VoiceMessageStatisticsItem `json:"VoiceMessageStatisticsItem,omitempty" name:"VoiceMessageStatisticsItem"`

	// 语音过滤统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceFilterStatisticsItem *VoiceFilterStatisticsItem `json:"VoiceFilterStatisticsItem,omitempty" name:"VoiceFilterStatisticsItem"`

	// 统计时间
	Date *string `json:"Date,omitempty" name:"Date"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest

	// 应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 腾讯云项目ID，默认为0，表示默认项目
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 需要支持的引擎列表，默认全选。
	EngineList []*string `json:"EngineList,omitempty" name:"EngineList" list`

	// 服务区域列表，默认全选。
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList" list`

	// 实时语音服务配置数据
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitempty" name:"RealtimeSpeechConf"`

	// 语音消息及转文本服务配置数据
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitempty" name:"VoiceMessageConf"`

	// 语音分析服务配置数据
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitempty" name:"VoiceFilterConf"`

	// 需要添加的标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAppRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAppResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用ID，由后台自动生成。
		BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

		// 应用名称，透传输入参数的AppName
		AppName *string `json:"AppName,omitempty" name:"AppName"`

		// 项目ID，透传输入的ProjectId
		ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

		// 应用密钥，GME SDK初始化时使用
		SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

		// 服务创建时间戳
		CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

		// 实时语音服务配置数据
		RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitempty" name:"RealtimeSpeechConf"`

		// 语音消息及转文本服务配置数据
		VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitempty" name:"VoiceMessageConf"`

		// 语音分析服务配置数据
		VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitempty" name:"VoiceFilterConf"`
	} `json:"Response"`
}

func (r *CreateAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAppResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAppStatisticsRequest struct {
	*tchttp.BaseRequest

	// GME应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 数据开始时间，东八区时间，格式: 年-月-日，如: 2018-07-13
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 数据结束时间，东八区时间，格式: 年-月-日，如: 2018-07-13
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 要查询的服务列表，取值：RealTimeSpeech/VoiceMessage/VoiceFilter
	Services []*string `json:"Services,omitempty" name:"Services" list`
}

func (r *DescribeAppStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAppStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAppStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用用量统计数据
		AppStatistics []*AppStatisticsItem `json:"AppStatistics,omitempty" name:"AppStatistics" list`
	} `json:"Response"`
}

func (r *DescribeAppStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAppStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFilterResultListRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 开始时间，格式为 年-月-日，如: 2018-07-11
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束时间，格式为 年-月-日，如: 2018-07-11
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFilterResultListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFilterResultListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFilterResultListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 过滤结果总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 当前分页过滤结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*VoiceFilterInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFilterResultListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFilterResultListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFilterResultRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 文件ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`
}

func (r *DescribeFilterResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFilterResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFilterResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 过滤结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *VoiceFilterInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFilterResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFilterResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScanResult struct {

	// 业务返回码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 数据唯一 ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 检测完成的时间戳
	ScanFinishTime *uint64 `json:"ScanFinishTime,omitempty" name:"ScanFinishTime"`

	// 是否违规
	HitFlag *bool `json:"HitFlag,omitempty" name:"HitFlag"`

	// 是否为流
	Live *bool `json:"Live,omitempty" name:"Live"`

	// 业务返回描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 检测结果，Code 为 0 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanPiece []*ScanPiece `json:"ScanPiece,omitempty" name:"ScanPiece" list`

	// 提交检测的时间戳
	ScanStartTime *uint64 `json:"ScanStartTime,omitempty" name:"ScanStartTime"`

	// 语音检测场景，对应请求时的 Scene
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes" list`

	// 语音检测任务 ID，由后台分配
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 文件或接流地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 检测任务执行结果状态，分别为：
	// <li>Start: 任务开始</li>
	// <li>Success: 成功结束</li>
	// <li>Error: 异常</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeScanResultListRequest struct {
	*tchttp.BaseRequest

	// 应用 ID，登录[控制台](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 查询的任务 ID 列表，任务 ID 列表最多支持 100 个。
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList" list`

	// 任务返回结果数量，默认10，上限500。大文件任务忽略此参数，返回全量结果
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeScanResultListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScanResultListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScanResultListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 要查询的语音检测任务的结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*DescribeScanResult `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanResultListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScanResultListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAppStatusRequest struct {
	*tchttp.BaseRequest

	// 应用ID，创建应用后由后台生成并返回。
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 应用状态，取值：open/close
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAppStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAppStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAppStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// GME应用ID
		BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

		// 应用状态，取值：open/close
		Status *string `json:"Status,omitempty" name:"Status"`
	} `json:"Response"`
}

func (r *ModifyAppStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAppStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RealTimeSpeechStatisticsItem struct {

	// 大陆地区DAU
	MainLandDau *uint64 `json:"MainLandDau,omitempty" name:"MainLandDau"`

	// 大陆地区PCU
	MainLandPcu *uint64 `json:"MainLandPcu,omitempty" name:"MainLandPcu"`

	// 大陆地区总使用时长，单位为min
	MainLandDuration *uint64 `json:"MainLandDuration,omitempty" name:"MainLandDuration"`

	// 海外地区DAU
	OverseaDau *uint64 `json:"OverseaDau,omitempty" name:"OverseaDau"`

	// 海外地区PCU
	OverseaPcu *uint64 `json:"OverseaPcu,omitempty" name:"OverseaPcu"`

	// 海外地区总使用时长，单位为min
	OverseaDuration *uint64 `json:"OverseaDuration,omitempty" name:"OverseaDuration"`
}

type RealtimeSpeechConf struct {

	// 实时语音服务开关，取值：open/close
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实时语音音质类型，取值：high-高音质，ordinary-普通音质。默认高音质。普通音质仅白名单开放，如需要普通音质，请联系腾讯云商务。
	Quality *string `json:"Quality,omitempty" name:"Quality"`
}

type ScanDetail struct {

	// 违规场景，参照<a href="https://cloud.tencent.com/document/product/607/37622#Label_Value">Label</a>定义
	Label *string `json:"Label,omitempty" name:"Label"`

	// 该场景下概率[0.00,100.00],分值越大违规概率越高
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// 违规关键字
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 关键字在音频的开始时间，从0开始的偏移量，单位为毫秒
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 关键字在音频的结束时间，从0开始的偏移量,，单位为毫秒
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

type ScanPiece struct {

	// 流检测时返回，音频转存地址，保留30min
	// 注意：此字段可能返回 null，表示取不到有效值。
	DumpUrl *string `json:"DumpUrl,omitempty" name:"DumpUrl"`

	// 是否违规
	HitFlag *bool `json:"HitFlag,omitempty" name:"HitFlag"`

	// 违规主要类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainType *string `json:"MainType,omitempty" name:"MainType"`

	// 语音检测详情
	ScanDetail []*ScanDetail `json:"ScanDetail,omitempty" name:"ScanDetail" list`

	// gme实时语音房间ID，透传任务传入时的RoomId
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// gme实时语音用户ID，透传任务传入时的OpenId
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info *string `json:"Info,omitempty" name:"Info"`

	// 流检测时分片在流中的偏移时间，单位毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 流检测时分片时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 分片开始检测时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PieceStartTime *uint64 `json:"PieceStartTime,omitempty" name:"PieceStartTime"`
}

type ScanVoiceRequest struct {
	*tchttp.BaseRequest

	// 应用ID，登录[控制台 - 服务管理](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 语音检测场景，参数值目前要求为 default。 预留场景设置： 谩骂、色情、涉政、广告、暴恐、违禁等场景，<a href="#Label_Value">具体取值见上述 Label 说明。</a>
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes" list`

	// 是否为直播流。值为 false 时表示普通语音文件检测；为 true 时表示语音流检测。
	Live *bool `json:"Live,omitempty" name:"Live"`

	// 语音检测任务列表，列表最多支持100个检测任务。结构体中包含：
	// <li>DataId：数据的唯一ID</li>
	// <li>Url：数据文件的url，为 urlencode 编码，流式则为拉流地址</li>
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks" list`

	// 异步检测结果回调地址，具体见上述<a href="#Callback_Declare">回调相关说明</a>。（说明：该字段为空时，必须通过接口(查询语音检测结果)获取检测结果）。
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

func (r *ScanVoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ScanVoiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ScanVoiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 语音检测返回。Data 字段是 JSON 数组，每一个元素包含：<li>DataId： 请求中对应的 DataId。</li>
	// <li>TaskID ：该检测任务的 ID，用于轮询语音检测结果。</li>
		Data []*ScanVoiceResult `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanVoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ScanVoiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ScanVoiceResult struct {

	// 数据ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type Tag struct {

	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Task struct {

	// 数据的唯一ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 数据文件的url，为 urlencode 编码，流式则为拉流地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// gme实时语音房间ID，通过gme实时语音进行语音分析时输入
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// gme实时语音用户ID，通过gme实时语音进行语音分析时输入
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
}

type VoiceFilter struct {

	// 过滤类型，1：政治，2：色情，3：涉毒，4：谩骂
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 过滤命中关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Word *string `json:"Word,omitempty" name:"Word"`
}

type VoiceFilterConf struct {

	// 语音过滤服务开关，取值：open/close
	Status *string `json:"Status,omitempty" name:"Status"`
}

type VoiceFilterInfo struct {

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 文件ID，表示文件唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 数据创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 过滤结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*VoiceFilter `json:"Data,omitempty" name:"Data" list`
}

type VoiceFilterRequest struct {
	*tchttp.BaseRequest

	// 应用ID，登录[控制台](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 文件ID，表示文件唯一ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件url，urlencode编码，FileUrl和FileContent二选一
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 文件内容，base64编码，FileUrl和FileContent二选一
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 用户ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
}

func (r *VoiceFilterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VoiceFilterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VoiceFilterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VoiceFilterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VoiceFilterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VoiceFilterStatisticsItem struct {

	// 语音过滤总时长
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
}

type VoiceMessageConf struct {

	// 离线语音服务开关，取值：open/close
	Status *string `json:"Status,omitempty" name:"Status"`

	// 离线语音支持语种，取值： all-全部，cnen-中英文。默认为中英文
	Language *string `json:"Language,omitempty" name:"Language"`
}

type VoiceMessageStatisticsItem struct {

	// 离线语音DAU
	Dau *uint64 `json:"Dau,omitempty" name:"Dau"`
}
