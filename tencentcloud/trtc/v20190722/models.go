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

package v20190722

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AbnormalEvent struct {

	// 异常事件ID，具体值查看附录：异常体验ID映射表：https://cloud.tencent.com/document/product/647/44916
	AbnormalEventId *uint64 `json:"AbnormalEventId,omitempty" name:"AbnormalEventId"`

	// 远端用户ID,""：表示异常事件不是由远端用户产生
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`
}

type AbnormalExperience struct {

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 异常体验ID
	ExperienceId *uint64 `json:"ExperienceId,omitempty" name:"ExperienceId"`

	// 字符串房间号
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 异常事件数组
	AbnormalEventList []*AbnormalEvent `json:"AbnormalEventList,omitempty" name:"AbnormalEventList" list`

	// 异常事件的上报时间
	EventTime *uint64 `json:"EventTime,omitempty" name:"EventTime"`
}

type CreateTroubleInfoRequest struct {
	*tchttp.BaseRequest

	// 应用的ID
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 老师用户ID
	TeacherUserId *string `json:"TeacherUserId,omitempty" name:"TeacherUserId"`

	// 学生用户ID
	StudentUserId *string `json:"StudentUserId,omitempty" name:"StudentUserId"`

	// 体验异常端（老师或学生）的用户 ID。
	TroubleUserId *string `json:"TroubleUserId,omitempty" name:"TroubleUserId"`

	// 异常类型。
	// 1. 仅视频异常
	// 2. 仅声音异常
	// 3. 音视频都异常
	// 5. 进房异常
	// 4. 切课
	// 6. 求助
	// 7. 问题反馈
	// 8. 投诉
	TroubleType *uint64 `json:"TroubleType,omitempty" name:"TroubleType"`

	// 异常发生的UNIX 时间戳，单位为秒。
	TroubleTime *uint64 `json:"TroubleTime,omitempty" name:"TroubleTime"`

	// 异常详情
	TroubleMsg *string `json:"TroubleMsg,omitempty" name:"TroubleMsg"`
}

func (r *CreateTroubleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTroubleInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTroubleInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTroubleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTroubleInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalEventRequest struct {
	*tchttp.BaseRequest

	// 用户SDKAppID，查询SDKAppID下任意20条异常体验事件（可能不同房间）
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 房间号，查询房间内任意20条以内异常体验事件
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DescribeAbnormalEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAbnormalEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据总条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 异常体验列表
		AbnormalExperienceList []*AbnormalExperience `json:"AbnormalExperienceList,omitempty" name:"AbnormalExperienceList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAbnormalEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCallDetailRequest struct {
	*tchttp.BaseRequest

	// 通话 ID（唯一标识一次通话）： sdkappid_roomgString（房间号_createTime（房间创建时间，unix时间戳，单位为s）。通过 DescribeRoomInformation（查询房间列表）接口获取。
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// 查询开始时间，5天内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需查询的用户数组，不填默认返回6个用户,最多可填6个用户
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds" list`

	// 需查询的指标，不填则只返回用户列表，填all则返回所有指标。
	// appCpu：APP CPU使用率；
	// sysCpu：系统 CPU使用率；
	// aBit：上/下行音频码率；
	// aBlock：音频卡顿时长；
	// bigvBit：上/下行视频码率；
	// bigvCapFps：视频采集帧率；
	// bigvEncFps：视频发送帧率；
	// bigvDecFps：渲染帧率；
	// bigvBlock：视频卡顿时长；
	// aLoss：上/下行音频丢包；
	// bigvLoss：上/下行视频丢包；
	// bigvWidth：上/下行分辨率宽；
	// bigvHeight：上/下行分辨率高
	DataType []*string `json:"DataType,omitempty" name:"DataType" list`
}

func (r *DescribeCallDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCallDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCallDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的用户总条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 用户信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserList []*UserInformation `json:"UserList,omitempty" name:"UserList" list`

		// 质量数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*QualityData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCallDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCallDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDetailEventRequest struct {
	*tchttp.BaseRequest

	// 通话 ID（唯一标识一次通话）： sdkappid_roomgString（房间号_createTime（房间创建时间，unix时间戳，单位s）。通过 DescribeRoomInformation（查询房间列表）接口获取。
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// 查询开始时间，5天内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户id
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 房间号
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DescribeDetailEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDetailEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDetailEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的事件列表
		Data []*EventList `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDetailEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDetailEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryScaleRequest struct {
	*tchttp.BaseRequest

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询开始时间，5天内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeHistoryScaleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHistoryScaleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryScaleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 返回的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		ScaleList []*ScaleInfomation `json:"ScaleList,omitempty" name:"ScaleList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistoryScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHistoryScaleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeNetworkRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间，24小时内，，本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需查询的数据类型
	// sendLossRateRaw：上行丢包率；
	// recvLossRateRaw：下行丢包率
	DataType []*string `json:"DataType,omitempty" name:"DataType" list`
}

func (r *DescribeRealtimeNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeNetworkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeNetworkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询返回的数据
		Data []*RealtimeData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealtimeNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeNetworkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeQualityRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间，24小时内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询的数据类型
	// enterTotalSuccPercent：进房成功率
	// fistFreamInSecRate：首帧秒开率
	// blockPercent：视频卡顿率
	// audioBlockPercent：音频卡顿率
	DataType []*string `json:"DataType,omitempty" name:"DataType" list`
}

func (r *DescribeRealtimeQualityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeQualityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeQualityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据类型
		Data []*RealtimeData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealtimeQualityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeQualityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeScaleRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间，24小时内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询的数据类型
	// UserNum：通话人数；
	// RoomNum：房间数
	DataType []*string `json:"DataType,omitempty" name:"DataType" list`
}

func (r *DescribeRealtimeScaleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeScaleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeScaleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据数组
		Data []*RealtimeData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealtimeScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeScaleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoomInformationRequest struct {
	*tchttp.BaseRequest

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询开始时间，5天内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 数字房间号
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 分页index（不填默认只返回10个）
	PageNumber *string `json:"PageNumber,omitempty" name:"PageNumber"`

	// 分页大小（不填默认返回10个,最多不超过100条）
	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeRoomInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoomInformationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoomInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据总条数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 房间信息列表
		RoomList []*RoomState `json:"RoomList,omitempty" name:"RoomList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoomInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoomInformationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DismissRoomRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DismissRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DismissRoomRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DismissRoomResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DismissRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DismissRoomResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EncodeParams struct {

	// 混流-输出流音频采样率。取值为[96000, 88200, 64000, 48000, 44100, 32000,24000, 22050, 16000, 12000, 11025, 8000]。
	AudioSampleRate *uint64 `json:"AudioSampleRate,omitempty" name:"AudioSampleRate"`

	// 混流-输出流音频码率。取值范围[8,500]，单位为Kbps。
	AudioBitrate *uint64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// 混流-输出流音频声道数，取值范围[1,2]。
	AudioChannels *uint64 `json:"AudioChannels,omitempty" name:"AudioChannels"`

	// 混流-输出流宽，音视频输出时必填。取值范围[0,1920]，单位为像素值。
	VideoWidth *uint64 `json:"VideoWidth,omitempty" name:"VideoWidth"`

	// 混流-输出流高，音视频输出时必填。取值范围[0,1080]，单位为像素值。
	VideoHeight *uint64 `json:"VideoHeight,omitempty" name:"VideoHeight"`

	// 混流-输出流码率，音视频输出时必填。取值范围[1,10000]，单位为Kbps。
	VideoBitrate *uint64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// 混流-输出流帧率，音视频输出时必填。取值为[6,12,15,24,30,48,60]，不在上述帧率值内系统会自动调整。
	VideoFramerate *uint64 `json:"VideoFramerate,omitempty" name:"VideoFramerate"`

	// 混流-输出流gop，音视频输出时必填。取值范围[1,5]，单位为秒。
	VideoGop *uint64 `json:"VideoGop,omitempty" name:"VideoGop"`

	// 混流-输出流背景色。
	BackgroundColor *uint64 `json:"BackgroundColor,omitempty" name:"BackgroundColor"`
}

type EventList struct {

	// 数据内容
	Content []*EventMessage `json:"Content,omitempty" name:"Content" list`

	// 发送端的userId
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`
}

type EventMessage struct {

	// 视频流类型：
	// 0：与视频无关的事件；
	// 2：视频为大画面；
	// 3：视频为小画面；
	// 7：视频为旁路画面；
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 事件上报的时间戳，unix时间（1589891188801ms)
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 事件Id：分为sdk的事件和webrtc的事件，详情见：附录/事件 ID 映射表：https://cloud.tencent.com/document/product/647/44916
	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`

	// 事件的第一个参数，如视频分辨率宽
	ParamOne *int64 `json:"ParamOne,omitempty" name:"ParamOne"`

	// 事件的第二个参数，如视频分辨率高
	ParamTwo *int64 `json:"ParamTwo,omitempty" name:"ParamTwo"`
}

type LayoutParams struct {

	// 混流布局模板ID，0为悬浮模板(默认);1为九宫格模板;2为屏幕分享模板
	Template *uint64 `json:"Template,omitempty" name:"Template"`

	// 屏幕分享模板中有效，代表左侧大画面对应的用户ID
	MainVideoUserId *string `json:"MainVideoUserId,omitempty" name:"MainVideoUserId"`

	// 屏幕分享模板中有效，代表左侧大画面对应的流类型，0为摄像头，1为屏幕分享。左侧大画面为web用户时此值填0
	MainVideoStreamType *uint64 `json:"MainVideoStreamType,omitempty" name:"MainVideoStreamType"`
}

type OutputParams struct {

	// 直播流ID，该流ID不能与用户旁路的流ID相同。
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// 取值范围[0,1]， 填0：直播流为音视频(默认); 填1：直播流为纯音频
	PureAudioStream *uint64 `json:"PureAudioStream,omitempty" name:"PureAudioStream"`

	// 自定义录制文件名
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 取值范围[0,1]，填1：指定录制文件格式为mp3
	RecordAudioOnly *uint64 `json:"RecordAudioOnly,omitempty" name:"RecordAudioOnly"`
}

type QualityData struct {

	// 数据内容
	Content []*TimeValue `json:"Content,omitempty" name:"Content" list`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 对端Id,为空时表示上行数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`

	// 数据类型
	DataType *string `json:"DataType,omitempty" name:"DataType"`
}

type RealtimeData struct {

	// 返回的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*TimeValue `json:"Content,omitempty" name:"Content" list`

	// 数据类型字段
	DataType *string `json:"DataType,omitempty" name:"DataType"`
}

type RemoveUserRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 要移出的用户列表，最多10个。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds" list`
}

func (r *RemoveUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RoomState struct {

	// 通话ID（唯一标识一次通话）
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// 房间号
	RoomString *string `json:"RoomString,omitempty" name:"RoomString"`

	// 房间创建时间
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 房间销毁时间
	DestroyTime *uint64 `json:"DestroyTime,omitempty" name:"DestroyTime"`

	// 房间是否已经结束
	IsFinished *bool `json:"IsFinished,omitempty" name:"IsFinished"`

	// 房间创建者Id
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type ScaleInfomation struct {

	// 每天开始的时间
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 房间人数，用户重复进入同一个房间为1次
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserNumber *uint64 `json:"UserNumber,omitempty" name:"UserNumber"`

	// 房间人次，用户每次进入房间为一次
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserCount *uint64 `json:"UserCount,omitempty" name:"UserCount"`

	// sdkappid下一天内的房间数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomNumbers *uint64 `json:"RoomNumbers,omitempty" name:"RoomNumbers"`
}

type StartMCUMixTranscodeRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 混流输出控制参数。
	OutputParams *OutputParams `json:"OutputParams,omitempty" name:"OutputParams"`

	// 混流输出编码参数。
	EncodeParams *EncodeParams `json:"EncodeParams,omitempty" name:"EncodeParams"`

	// 混流输出布局参数。
	LayoutParams *LayoutParams `json:"LayoutParams,omitempty" name:"LayoutParams"`
}

func (r *StartMCUMixTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMCUMixTranscodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartMCUMixTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMCUMixTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartMCUMixTranscodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopMCUMixTranscodeRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *StopMCUMixTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopMCUMixTranscodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopMCUMixTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMCUMixTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopMCUMixTranscodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TimeValue struct {

	// 时间
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 当前时间取值，unix时间戳
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type UserInformation struct {

	// 房间号
	RoomStr *string `json:"RoomStr,omitempty" name:"RoomStr"`

	// 用户Id
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户进房时间
	JoinTs *uint64 `json:"JoinTs,omitempty" name:"JoinTs"`

	// 用户退房时间
	LeaveTs *uint64 `json:"LeaveTs,omitempty" name:"LeaveTs"`

	// 终端类型
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Sdk版本号
	SdkVersion *string `json:"SdkVersion,omitempty" name:"SdkVersion"`

	// 客户端IP地址
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 判断用户是否已经离开房间
	Finished *bool `json:"Finished,omitempty" name:"Finished"`
}
