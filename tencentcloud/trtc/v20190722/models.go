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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AbnormalEvent struct {
	// 异常事件ID，具体值查看附录：异常体验ID映射表：https://cloud.tencent.com/document/product/647/44916
	AbnormalEventId *uint64 `json:"AbnormalEventId,omitnil,omitempty" name:"AbnormalEventId"`

	// 远端用户ID,""：表示异常事件不是由远端用户产生
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerId *string `json:"PeerId,omitnil,omitempty" name:"PeerId"`
}

type AbnormalExperience struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 异常体验ID
	ExperienceId *uint64 `json:"ExperienceId,omitnil,omitempty" name:"ExperienceId"`

	// 字符串房间号
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 异常事件数组
	AbnormalEventList []*AbnormalEvent `json:"AbnormalEventList,omitnil,omitempty" name:"AbnormalEventList"`

	// 异常事件的上报时间
	EventTime *uint64 `json:"EventTime,omitnil,omitempty" name:"EventTime"`
}

type AgentConfig struct {
	// 机器人的UserId，用于进房发起任务。【注意】这个UserId不能与当前房间内的主播观众[UserId](https://cloud.tencent.com/document/product/647/46351#userid)重复。如果一个房间发起多个任务时，机器人的UserId也不能相互重复，否则会中断前一个任务。需要保证机器人UserId在房间内唯一。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 机器人UserId对应的校验签名，即UserId和UserSig相当于机器人进房的登录密码，具体计算方法请参考TRTC计算[UserSig](https://cloud.tencent.com/document/product/647/45910#UserSig)的方案。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`

	// 机器人拉流的UserId, 填写后，机器人会拉取该UserId的流进行实时处理
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetUserId *string `json:"TargetUserId,omitnil,omitempty" name:"TargetUserId"`

	// 房间内超过MaxIdleTime 没有推流，后台自动关闭任务，默认值是60s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxIdleTime *uint64 `json:"MaxIdleTime,omitnil,omitempty" name:"MaxIdleTime"`

	// 机器人的欢迎语
	// 注意：此字段可能返回 null，表示取不到有效值。
	WelcomeMessage *string `json:"WelcomeMessage,omitnil,omitempty" name:"WelcomeMessage"`

	// 智能打断模式，默认为0，0表示服务端自动打断，1表示服务端不打断，由端上发送打断信令进行打断
	// 注意：此字段可能返回 null，表示取不到有效值。
	InterruptMode *uint64 `json:"InterruptMode,omitnil,omitempty" name:"InterruptMode"`

	// InterruptMode为0时使用，单位为毫秒，默认为500ms。表示服务端检测到持续InterruptSpeechDuration毫秒的人声则进行打断。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InterruptSpeechDuration *uint64 `json:"InterruptSpeechDuration,omitnil,omitempty" name:"InterruptSpeechDuration"`
}

type AgentParams struct {
	// 转推服务在TRTC房间使用的[UserId](https://cloud.tencent.com/document/product/647/46351#userid)，注意这个userId不能与其他TRTC或者转推服务等已经使用的UserId重复，建议可以把房间ID作为userId的标识的一部分。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 转推服务加入TRTC房间的用户签名，当前 UserId 对应的验证签名，相当于登录密码，具体计算方法请参考TRTC计算[UserSig](https://cloud.tencent.com/document/product/647/45910#UserSig)的方案。
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`

	// 所有参与混流转推的主播持续离开TRTC房间或切换成观众超过MaxIdleTime的时长，自动停止转推，单位：秒。默认值为 30 秒，该值需大于等于 5秒，且小于等于 86400秒(24小时)。
	MaxIdleTime *uint64 `json:"MaxIdleTime,omitnil,omitempty" name:"MaxIdleTime"`
}

type AudioEncode struct {
	// 输出流音频采样率。取值为[48000, 44100, 32000, 24000, 16000, 8000]，单位是Hz。
	SampleRate *uint64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 输出流音频声道数，取值范围[1,2]，1表示混流输出音频为单声道，2表示混流输出音频为双声道。
	Channel *uint64 `json:"Channel,omitnil,omitempty" name:"Channel"`

	// 输出流音频码率。取值范围[8,500]，单位为kbps。
	BitRate *uint64 `json:"BitRate,omitnil,omitempty" name:"BitRate"`

	// 输出流音频编码类型，取值范围[0, 1, 2]，0为LC-AAC，1为HE-AAC，2为HE-AACv2。默认值为0。当音频编码设置为HE-AACv2时，只支持输出流音频声道数为双声道。HE-AAC和HE-AACv2支持的输出流音频采样率范围为[48000, 44100, 32000, 24000, 16000]。
	Codec *uint64 `json:"Codec,omitnil,omitempty" name:"Codec"`
}

type AudioEncodeParams struct {
	// 音频采样率，取值为[48000, 44100]，单位是Hz。
	SampleRate *uint64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 音频声道数，取值范围[1,2]，1表示音频为单声道，2表示音频为双声道。
	Channel *uint64 `json:"Channel,omitnil,omitempty" name:"Channel"`

	// 音频码率，取值范围[8,500]，单位为kbps。
	BitRate *uint64 `json:"BitRate,omitnil,omitempty" name:"BitRate"`

	// 音量，取值范围[0,300]。默认100，表示原始音量；0表示静音。
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`
}

type AudioParams struct {
	// 音频采样率枚举值:(注意1 代表48000HZ, 2 代表44100HZ, 3 代表16000HZ)
	// 1：48000Hz（默认）;
	// 2：44100Hz
	// 3：16000Hz。
	SampleRate *uint64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 声道数枚举值:
	// 1：单声道;
	// 2：双声道（默认）。
	Channel *uint64 `json:"Channel,omitnil,omitempty" name:"Channel"`

	// 音频码率: 取值范围[32000, 128000] ，单位bps，默认64000bps。
	BitRate *uint64 `json:"BitRate,omitnil,omitempty" name:"BitRate"`
}

type CloudStorage struct {
	// 腾讯云对象存储COS以及第三方云存储账号信息
	// 0：腾讯云对象存储 COS
	// 1：AWS
	// 【注意】目前第三方云存储仅支持AWS，更多第三方云存储陆续支持中
	// 示例值：0
	Vendor *uint64 `json:"Vendor,omitnil,omitempty" name:"Vendor"`

	// 腾讯云对象存储的[地域信息]（https://cloud.tencent.com/document/product/436/6224#.E5.9C.B0.E5.9F.9F）。
	// 示例值：cn-shanghai-1
	// 
	// AWS S3[地域信息]（https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-regions）
	// 示例值：ap-southeast-3	
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 云存储桶名称。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 云存储的access_key账号信息。
	// 若存储至腾讯云对象存储COS，请前往https://console.cloud.tencent.com/cam/capi 查看或创建，对应链接中密钥字段的SecretId值。
	// 示例值：test-accesskey
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 云存储的secret_key账号信息。
	// 若存储至腾讯云对象存储COS，请前往https://console.cloud.tencent.com/cam/capi 查看或创建，对应链接中密钥字段的SecretKey值。
	// 示例值：test-secretkey
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 云存储bucket 的指定位置，由字符串数组组成。合法的字符串范围az,AZ,0~9,'_'和'-'，举个例子，录制文件xxx.m3u8在 ["prefix1", "prefix2"]作用下，会变成prefix1/prefix2/TaskId/xxx.m3u8。
	// 示例值：["prefix1", "prefix2"]
	FileNamePrefix []*string `json:"FileNamePrefix,omitnil,omitempty" name:"FileNamePrefix"`
}

type CloudVod struct {
	// 腾讯云点播相关参数。
	TencentVod *TencentVod `json:"TencentVod,omitnil,omitempty" name:"TencentVod"`
}

// Predefined struct for user
type ControlAIConversationRequestParams struct {
	// 任务唯一标识
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 控制命令，目前支持命令如下：
	// 
	// - ServerPushText，服务端发送文本给AI机器人，AI机器人会播报该文本
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 服务端发送播报文本命令，当Command为ServerPushText时必填
	ServerPushText *ServerPushText `json:"ServerPushText,omitnil,omitempty" name:"ServerPushText"`
}

type ControlAIConversationRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一标识
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 控制命令，目前支持命令如下：
	// 
	// - ServerPushText，服务端发送文本给AI机器人，AI机器人会播报该文本
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 服务端发送播报文本命令，当Command为ServerPushText时必填
	ServerPushText *ServerPushText `json:"ServerPushText,omitnil,omitempty" name:"ServerPushText"`
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
type CreateCloudRecordingRequestParams struct {
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和录制的房间所对应的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// TRTC的[RoomId](https://cloud.tencent.com/document/product/647/46351#roomid)，录制的TRTC房间所对应的RoomId。
	// 注：房间号类型默认为整型，若房间号类型为字符串，请通过RoomIdType指定。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 录制机器人的UserId，用于进房发起录制任务。
	// 【*注意】这个UserId不能与当前房间内的主播观众[UserId](https://cloud.tencent.com/document/product/647/46351#userid)重复。如果一个房间发起多个录制任务时，机器人的userid也不能相互重复，否则会中断前一个录制任务。建议可以把房间ID作为UserId的标识的一部分，即录制机器人UserId在房间内唯一。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 录制机器人UserId对应的校验签名，即UserId和UserSig相当于录制机器人进房的登录密码，具体计算方法请参考TRTC计算[UserSig](https://cloud.tencent.com/document/product/647/45910#UserSig)的方案。
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`

	// 云端录制控制参数。
	RecordParams *RecordParams `json:"RecordParams,omitnil,omitempty" name:"RecordParams"`

	// 云端录制文件上传到云存储的参数（不支持同时设置云点播VOD和对象存储COS）
	StorageParams *StorageParams `json:"StorageParams,omitnil,omitempty" name:"StorageParams"`

	// TRTC房间号的类型。
	// 【*注意】必须和录制的房间所对应的RoomId类型相同:
	// 0: 字符串类型的RoomId
	// 1: 32位整型的RoomId（默认）
	RoomIdType *uint64 `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`

	// 合流的转码参数，录制模式为合流的时候可以设置。
	MixTranscodeParams *MixTranscodeParams `json:"MixTranscodeParams,omitnil,omitempty" name:"MixTranscodeParams"`

	// 合流的布局参数，录制模式为合流的时候可以设置。
	MixLayoutParams *MixLayoutParams `json:"MixLayoutParams,omitnil,omitempty" name:"MixLayoutParams"`

	// 接口可以调用的时效性，从成功开启录制并获得任务ID后开始计算，超时后无法调用查询、更新和停止等接口，但是录制任务不会停止。 参数的单位是小时，默认72小时（3天），最大可设置720小时（30天），最小设置6小时。举例说明：如果不设置该参数，那么开始录制成功后，查询、更新和停止录制的调用时效为72个小时。
	ResourceExpiredHour *uint64 `json:"ResourceExpiredHour,omitnil,omitempty" name:"ResourceExpiredHour"`

	// TRTC房间权限加密串，只有在TRTC控制台启用了高级权限控制的时候需要携带，在TRTC控制台如果开启高级权限控制后，TRTC 的后台服务系统会校验一个叫做 [PrivateMapKey] 的“权限票据”，权限票据中包含了一个加密后的 RoomId 和一个加密后的“权限位列表”。由于 PrivateMapKey 中包含 RoomId，所以只提供了 UserSig 没有提供 PrivateMapKey 时，并不能进入指定的房间。
	PrivateMapKey *string `json:"PrivateMapKey,omitnil,omitempty" name:"PrivateMapKey"`
}

type CreateCloudRecordingRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和录制的房间所对应的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// TRTC的[RoomId](https://cloud.tencent.com/document/product/647/46351#roomid)，录制的TRTC房间所对应的RoomId。
	// 注：房间号类型默认为整型，若房间号类型为字符串，请通过RoomIdType指定。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 录制机器人的UserId，用于进房发起录制任务。
	// 【*注意】这个UserId不能与当前房间内的主播观众[UserId](https://cloud.tencent.com/document/product/647/46351#userid)重复。如果一个房间发起多个录制任务时，机器人的userid也不能相互重复，否则会中断前一个录制任务。建议可以把房间ID作为UserId的标识的一部分，即录制机器人UserId在房间内唯一。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 录制机器人UserId对应的校验签名，即UserId和UserSig相当于录制机器人进房的登录密码，具体计算方法请参考TRTC计算[UserSig](https://cloud.tencent.com/document/product/647/45910#UserSig)的方案。
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`

	// 云端录制控制参数。
	RecordParams *RecordParams `json:"RecordParams,omitnil,omitempty" name:"RecordParams"`

	// 云端录制文件上传到云存储的参数（不支持同时设置云点播VOD和对象存储COS）
	StorageParams *StorageParams `json:"StorageParams,omitnil,omitempty" name:"StorageParams"`

	// TRTC房间号的类型。
	// 【*注意】必须和录制的房间所对应的RoomId类型相同:
	// 0: 字符串类型的RoomId
	// 1: 32位整型的RoomId（默认）
	RoomIdType *uint64 `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`

	// 合流的转码参数，录制模式为合流的时候可以设置。
	MixTranscodeParams *MixTranscodeParams `json:"MixTranscodeParams,omitnil,omitempty" name:"MixTranscodeParams"`

	// 合流的布局参数，录制模式为合流的时候可以设置。
	MixLayoutParams *MixLayoutParams `json:"MixLayoutParams,omitnil,omitempty" name:"MixLayoutParams"`

	// 接口可以调用的时效性，从成功开启录制并获得任务ID后开始计算，超时后无法调用查询、更新和停止等接口，但是录制任务不会停止。 参数的单位是小时，默认72小时（3天），最大可设置720小时（30天），最小设置6小时。举例说明：如果不设置该参数，那么开始录制成功后，查询、更新和停止录制的调用时效为72个小时。
	ResourceExpiredHour *uint64 `json:"ResourceExpiredHour,omitnil,omitempty" name:"ResourceExpiredHour"`

	// TRTC房间权限加密串，只有在TRTC控制台启用了高级权限控制的时候需要携带，在TRTC控制台如果开启高级权限控制后，TRTC 的后台服务系统会校验一个叫做 [PrivateMapKey] 的“权限票据”，权限票据中包含了一个加密后的 RoomId 和一个加密后的“权限位列表”。由于 PrivateMapKey 中包含 RoomId，所以只提供了 UserSig 没有提供 PrivateMapKey 时，并不能进入指定的房间。
	PrivateMapKey *string `json:"PrivateMapKey,omitnil,omitempty" name:"PrivateMapKey"`
}

func (r *CreateCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserId")
	delete(f, "UserSig")
	delete(f, "RecordParams")
	delete(f, "StorageParams")
	delete(f, "RoomIdType")
	delete(f, "MixTranscodeParams")
	delete(f, "MixLayoutParams")
	delete(f, "ResourceExpiredHour")
	delete(f, "PrivateMapKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudRecordingResponseParams struct {
	// 云录制服务分配的任务 ID。任务 ID 是对一次录制生命周期过程的唯一标识，结束录制时会失去意义。任务 ID需要业务保存下来，作为下次针对这个录制任务操作的参数。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudRecordingResponseParams `json:"Response"`
}

func (r *CreateCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePictureRequestParams struct {
	// 应用id
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 图片内容经base64编码后的string格式,最大长度为2M
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 图片后缀名
	Suffix *string `json:"Suffix,omitnil,omitempty" name:"Suffix"`

	// 图片长度
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 图片宽度
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 显示位置x轴方向
	XPosition *uint64 `json:"XPosition,omitnil,omitempty" name:"XPosition"`

	// 显示位置y轴方向
	YPosition *uint64 `json:"YPosition,omitnil,omitempty" name:"YPosition"`
}

type CreatePictureRequest struct {
	*tchttp.BaseRequest
	
	// 应用id
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 图片内容经base64编码后的string格式,最大长度为2M
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 图片后缀名
	Suffix *string `json:"Suffix,omitnil,omitempty" name:"Suffix"`

	// 图片长度
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 图片宽度
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 显示位置x轴方向
	XPosition *uint64 `json:"XPosition,omitnil,omitempty" name:"XPosition"`

	// 显示位置y轴方向
	YPosition *uint64 `json:"YPosition,omitnil,omitempty" name:"YPosition"`
}

func (r *CreatePictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Content")
	delete(f, "Suffix")
	delete(f, "Height")
	delete(f, "Width")
	delete(f, "XPosition")
	delete(f, "YPosition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePictureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePictureResponseParams struct {
	// 图片id
	PictureId *uint64 `json:"PictureId,omitnil,omitempty" name:"PictureId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePictureResponse struct {
	*tchttp.BaseResponse
	Response *CreatePictureResponseParams `json:"Response"`
}

func (r *CreatePictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudRecordingRequestParams struct {
	// TRTC的SDKAppId，和录制的房间所对应的SDKAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 录制任务的唯一Id，在启动录制成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteCloudRecordingRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId，和录制的房间所对应的SDKAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 录制任务的唯一Id，在启动录制成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudRecordingResponseParams struct {
	// 云录制服务分配的任务 ID。任务 ID 是对一次录制生命周期过程的唯一标识，结束录制时会失去意义。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudRecordingResponseParams `json:"Response"`
}

func (r *DeleteCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePictureRequestParams struct {
	// 图片id
	PictureId *uint64 `json:"PictureId,omitnil,omitempty" name:"PictureId"`

	// 应用id
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DeletePictureRequest struct {
	*tchttp.BaseRequest
	
	// 图片id
	PictureId *uint64 `json:"PictureId,omitnil,omitempty" name:"PictureId"`

	// 应用id
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DeletePictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PictureId")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePictureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePictureResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePictureResponse struct {
	*tchttp.BaseResponse
	Response *DeletePictureResponseParams `json:"Response"`
}

func (r *DeletePictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIConversationRequestParams struct {
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和开启转录任务的房间使用的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 唯一标识一次任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 开启任务时填写的SessionId，如果没写则不返回。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeAIConversationRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和开启转录任务的房间使用的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 唯一标识一次任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 开启任务时填写的SessionId，如果没写则不返回。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
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
	delete(f, "SessionId")
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

	// 任务的唯一标识，在启动任务时生成
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
type DescribeAITranscriptionRequestParams struct {
	// 查询任务状态，不使用时传入空字符串。
	// 有两种查询方式：
	// 1、只填写TaskId，这种方式使用TaskId来查询任务
	// 2、TaskId为空字符串，填写SdkAppId和SessionId，这种方式不需要使用TaskId查询任务
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// TRTC的SdkAppId，和SessionId配合使用。
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 开启转录任务时传入的SessionId，和SdkAppId配合使用。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeAITranscriptionRequest struct {
	*tchttp.BaseRequest
	
	// 查询任务状态，不使用时传入空字符串。
	// 有两种查询方式：
	// 1、只填写TaskId，这种方式使用TaskId来查询任务
	// 2、TaskId为空字符串，填写SdkAppId和SessionId，这种方式不需要使用TaskId查询任务
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// TRTC的SdkAppId，和SessionId配合使用。
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 开启转录任务时传入的SessionId，和SdkAppId配合使用。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeAITranscriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAITranscriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAITranscriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAITranscriptionResponseParams struct {
	// 任务开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 转录任务状态。
	// 有4个值：
	// 1、Idle表示任务未开始
	// 2、Preparing表示任务准备中
	// 3、InProgress表示任务正在运行
	// 4、Stopped表示任务已停止，正在清理资源中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一标识一次任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 开启转录任务时填写的SessionId，如果没写则不返回。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAITranscriptionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAITranscriptionResponseParams `json:"Response"`
}

func (r *DescribeAITranscriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAITranscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallDetailInfoRequestParams struct {
	// 通话 ID（唯一标识一次通话）： SdkAppId_RoomId（房间号）_ CreateTime（房间创建时间，unix时间戳，单位为s）例：1400xxxxxx_218695_1590065777。通过 DescribeRoomInfo（查询历史房间列表）接口获取（[查询历史房间列表](https://cloud.tencent.com/document/product/647/44050)）。
	CommId *string `json:"CommId,omitnil,omitempty" name:"CommId"`

	// 查询开始时间，本地unix时间戳，单位为秒（如：1590065777），
	// 注意：支持查询14天内的数据。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳，单位为秒（如：1590065877）
	// 注意：查询起止时间需小于1小时，超过则返回null，即与StartTime间隔时间不超过1小时。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户SdkAppId（如：1400xxxxxx）。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需查询的用户数组，默认不填返回6个用户。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 需查询的指标，不填则只返回用户列表，填all则返回所有指标。
	// appCpu：APP CPU使用率；
	// sysCpu：系统 CPU使用率；
	// aBit：上/下行音频码率；单位：bps
	// aBlock：音频卡顿时长；单位：ms
	// bigvBit：上/下行视频码率；单位：bps
	// bigvCapFps：视频采集帧率；
	// bigvEncFps：视频发送帧率；
	// bigvDecFps：渲染帧率；
	// bigvBlock：视频卡顿时长；单位：ms
	// aLoss：上/下行音频丢包率；
	// bigvLoss：上/下行视频丢包率；
	// bigvWidth：上/下行分辨率宽；
	// bigvHeight：上/下行分辨率高；
	// aCapEnergy：音频采集能量；
	// aPlayEnergy：音频播放能量；
	// rtt：SDK到云端的往返延时；单位: ms
	DataType []*string `json:"DataType,omitnil,omitempty" name:"DataType"`

	// 当前页数，默认为0，
	// 注意：PageNumber和PageSize 其中一个不填均默认返回6条数据。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页个数，默认为6，
	// 范围：[1，100]
	// 注意：DataType不为null，UserIds长度不能超过6，PageSize最大值不超过6；
	// DataType 为null，UserIds长度不超过100，PageSize最大不超过100。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeCallDetailInfoRequest struct {
	*tchttp.BaseRequest
	
	// 通话 ID（唯一标识一次通话）： SdkAppId_RoomId（房间号）_ CreateTime（房间创建时间，unix时间戳，单位为s）例：1400xxxxxx_218695_1590065777。通过 DescribeRoomInfo（查询历史房间列表）接口获取（[查询历史房间列表](https://cloud.tencent.com/document/product/647/44050)）。
	CommId *string `json:"CommId,omitnil,omitempty" name:"CommId"`

	// 查询开始时间，本地unix时间戳，单位为秒（如：1590065777），
	// 注意：支持查询14天内的数据。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳，单位为秒（如：1590065877）
	// 注意：查询起止时间需小于1小时，超过则返回null，即与StartTime间隔时间不超过1小时。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户SdkAppId（如：1400xxxxxx）。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需查询的用户数组，默认不填返回6个用户。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 需查询的指标，不填则只返回用户列表，填all则返回所有指标。
	// appCpu：APP CPU使用率；
	// sysCpu：系统 CPU使用率；
	// aBit：上/下行音频码率；单位：bps
	// aBlock：音频卡顿时长；单位：ms
	// bigvBit：上/下行视频码率；单位：bps
	// bigvCapFps：视频采集帧率；
	// bigvEncFps：视频发送帧率；
	// bigvDecFps：渲染帧率；
	// bigvBlock：视频卡顿时长；单位：ms
	// aLoss：上/下行音频丢包率；
	// bigvLoss：上/下行视频丢包率；
	// bigvWidth：上/下行分辨率宽；
	// bigvHeight：上/下行分辨率高；
	// aCapEnergy：音频采集能量；
	// aPlayEnergy：音频播放能量；
	// rtt：SDK到云端的往返延时；单位: ms
	DataType []*string `json:"DataType,omitnil,omitempty" name:"DataType"`

	// 当前页数，默认为0，
	// 注意：PageNumber和PageSize 其中一个不填均默认返回6条数据。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页个数，默认为6，
	// 范围：[1，100]
	// 注意：DataType不为null，UserIds长度不能超过6，PageSize最大值不超过6；
	// DataType 为null，UserIds长度不超过100，PageSize最大不超过100。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeCallDetailInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallDetailInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "UserIds")
	delete(f, "DataType")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallDetailInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallDetailInfoResponseParams struct {
	// 返回的用户总条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 用户信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserList []*UserInformation `json:"UserList,omitnil,omitempty" name:"UserList"`

	// 质量数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*QualityData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCallDetailInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallDetailInfoResponseParams `json:"Response"`
}

func (r *DescribeCallDetailInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallDetailInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRecordingRequestParams struct {
	// TRTC的SDKAppId，和录制的房间所对应的SDKAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 录制任务的唯一Id，在启动录制成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeCloudRecordingRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId，和录制的房间所对应的SDKAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 录制任务的唯一Id，在启动录制成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRecordingResponseParams struct {
	// 录制任务的唯一Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 云端录制任务的状态信息。
	// Idle：表示当前录制任务空闲中
	// InProgress：表示当前录制任务正在进行中。
	// Exited：表示当前录制任务正在退出的过程中。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 录制文件信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageFileList []*StorageFile `json:"StorageFileList,omitnil,omitempty" name:"StorageFileList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudRecordingResponseParams `json:"Response"`
}

func (r *DescribeCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMixTranscodingUsageRequestParams struct {
	// 查询开始时间，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// TRTC的SdkAppId，和房间所对应的SdkAppId相同。如果没有这个参数，返回用户下全部实时音视频应用的汇总。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeMixTranscodingUsageRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// TRTC的SdkAppId，和房间所对应的SdkAppId相同。如果没有这个参数，返回用户下全部实时音视频应用的汇总。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeMixTranscodingUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMixTranscodingUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMixTranscodingUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMixTranscodingUsageResponseParams struct {
	// 用量类型，与UsageValue中各个位置的值对应。
	UsageKey []*string `json:"UsageKey,omitnil,omitempty" name:"UsageKey"`

	// 各个时间点用量明细。
	UsageList []*TrtcUsage `json:"UsageList,omitnil,omitempty" name:"UsageList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMixTranscodingUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMixTranscodingUsageResponseParams `json:"Response"`
}

func (r *DescribeMixTranscodingUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMixTranscodingUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePictureRequestParams struct {
	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 图片ID，不填时返回该应用下所有图片
	PictureId *uint64 `json:"PictureId,omitnil,omitempty" name:"PictureId"`

	// 每页数量，不填时默认为10
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码，不填时默认为1
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`
}

type DescribePictureRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 图片ID，不填时返回该应用下所有图片
	PictureId *uint64 `json:"PictureId,omitnil,omitempty" name:"PictureId"`

	// 每页数量，不填时默认为10
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码，不填时默认为1
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`
}

func (r *DescribePictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PictureId")
	delete(f, "PageSize")
	delete(f, "PageNo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePictureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePictureResponseParams struct {
	// 返回的图片记录数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 图片信息列表
	PictureInfo []*PictureInfo `json:"PictureInfo,omitnil,omitempty" name:"PictureInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePictureResponse struct {
	*tchttp.BaseResponse
	Response *DescribePictureResponseParams `json:"Response"`
}

func (r *DescribePictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordStatisticRequestParams struct {
	// 查询开始日期，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束日期，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用ID，可不传。传应用ID时返回的是该应用的用量，不传时返回多个应用的合计值。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeRecordStatisticRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始日期，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束日期，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用ID，可不传。传应用ID时返回的是该应用的用量，不传时返回多个应用的合计值。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeRecordStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordStatisticResponseParams struct {
	// 应用的用量信息数组。
	SdkAppIdUsages []*SdkAppIdRecordUsage `json:"SdkAppIdUsages,omitnil,omitempty" name:"SdkAppIdUsages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordStatisticResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordStatisticResponseParams `json:"Response"`
}

func (r *DescribeRecordStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordingUsageRequestParams struct {
	// 查询开始时间，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询单流录制或合流录制，值为"single"或"multi"。
	MixType *string `json:"MixType,omitnil,omitempty" name:"MixType"`

	// TRTC的SdkAppId，和房间所对应的SdkAppId相同。如果没有这个参数，返回用户下全部实时音视频应用的汇总。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeRecordingUsageRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询单流录制或合流录制，值为"single"或"multi"。
	MixType *string `json:"MixType,omitnil,omitempty" name:"MixType"`

	// TRTC的SdkAppId，和房间所对应的SdkAppId相同。如果没有这个参数，返回用户下全部实时音视频应用的汇总。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeRecordingUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordingUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MixType")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordingUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordingUsageResponseParams struct {
	// 用量类型，与UsageValue中各个位置的值对应。
	UsageKey []*string `json:"UsageKey,omitnil,omitempty" name:"UsageKey"`

	// 各个时间点用量明细。
	UsageList []*TrtcUsage `json:"UsageList,omitnil,omitempty" name:"UsageList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordingUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordingUsageResponseParams `json:"Response"`
}

func (r *DescribeRecordingUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordingUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelayUsageRequestParams struct {
	// 查询开始时间，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// TRTC的SdkAppId，和房间所对应的SdkAppId相同。如果没有这个参数，返回用户下全部实时音视频应用的汇总。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeRelayUsageRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// TRTC的SdkAppId，和房间所对应的SdkAppId相同。如果没有这个参数，返回用户下全部实时音视频应用的汇总。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeRelayUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelayUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRelayUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelayUsageResponseParams struct {
	// 用量类型，与UsageValue中各个位置的值对应。
	UsageKey []*string `json:"UsageKey,omitnil,omitempty" name:"UsageKey"`

	// 各个时间点用量明细。
	UsageList []*TrtcUsage `json:"UsageList,omitnil,omitempty" name:"UsageList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRelayUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRelayUsageResponseParams `json:"Response"`
}

func (r *DescribeRelayUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelayUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomInfoRequestParams struct {
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，本地unix时间戳，单位为秒（如：1590065777）注意：最大支持查询14天内的数据
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳，单位为秒（如：1590065877）
	// 注意：与StartTime间隔时间不超过24小时。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 房间号（如：223)
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 当前页数，默认为0，
	// 注意：PageNumber和PageSize 其中一个不填均默认返回10条数据。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页个数，默认为10，
	// 范围：[1，100]
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeRoomInfoRequest struct {
	*tchttp.BaseRequest
	
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，本地unix时间戳，单位为秒（如：1590065777）注意：最大支持查询14天内的数据
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳，单位为秒（如：1590065877）
	// 注意：与StartTime间隔时间不超过24小时。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 房间号（如：223)
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 当前页数，默认为0，
	// 注意：PageNumber和PageSize 其中一个不填均默认返回10条数据。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页个数，默认为10，
	// 范围：[1，100]
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
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
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomInfoResponseParams struct {
	// 返回当页数据总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 房间信息列表
	RoomList []*RoomState `json:"RoomList,omitnil,omitempty" name:"RoomList"`

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

// Predefined struct for user
type DescribeScaleInfoRequestParams struct {
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，本地unix时间戳，单位为秒（如：1590065777）
	// 注意：支持查询14天内的数据。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳，单位为秒（如：1590065877），建议与StartTime间隔时间超过24小时。
	// 注意：按天统计，结束时间大于前一天，否则查询数据为空（如：需查询20号数据，结束时间需晚于20号0点）。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeScaleInfoRequest struct {
	*tchttp.BaseRequest
	
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，本地unix时间戳，单位为秒（如：1590065777）
	// 注意：支持查询14天内的数据。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳，单位为秒（如：1590065877），建议与StartTime间隔时间超过24小时。
	// 注意：按天统计，结束时间大于前一天，否则查询数据为空（如：需查询20号数据，结束时间需晚于20号0点）。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeScaleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScaleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScaleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScaleInfoResponseParams struct {
	// 返回的数据条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 返回的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleList []*ScaleInfomation `json:"ScaleList,omitnil,omitempty" name:"ScaleList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScaleInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScaleInfoResponseParams `json:"Response"`
}

func (r *DescribeScaleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScaleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamIngestRequestParams struct {
	// TRTC的SDKAppId，和任务的房间所对应的SDKAppId相同
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务的唯一Id，在启动任务成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeStreamIngestRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId，和任务的房间所对应的SDKAppId相同
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务的唯一Id，在启动任务成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeStreamIngestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamIngestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamIngestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamIngestResponseParams struct {
	// 任务的状态信息。
	// InProgress：表示当前任务正在进行中。
	// NotExist：表示当前任务不存在。
	// 示例值：InProgress
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamIngestResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamIngestResponseParams `json:"Response"`
}

func (r *DescribeStreamIngestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamIngestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCMarketQualityDataRequestParams struct {
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，格式为YYYY-MM-DD。（查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 返回数据的粒度，支持设为以下值：
	// d：按天。此时返回查询时间范围内 UTC 时间为零点的数据。
	// h：按小时。此时返回查询时间范围内 UTC 时间为整小时的数据。
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

type DescribeTRTCMarketQualityDataRequest struct {
	*tchttp.BaseRequest
	
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，格式为YYYY-MM-DD。（查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 返回数据的粒度，支持设为以下值：
	// d：按天。此时返回查询时间范围内 UTC 时间为零点的数据。
	// h：按小时。此时返回查询时间范围内 UTC 时间为整小时的数据。
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *DescribeTRTCMarketQualityDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCMarketQualityDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTRTCMarketQualityDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCMarketQualityDataResponseParams struct {
	// TRTC监控数据出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TRTCDataResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTRTCMarketQualityDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTRTCMarketQualityDataResponseParams `json:"Response"`
}

func (r *DescribeTRTCMarketQualityDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCMarketQualityDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCMarketQualityMetricDataRequestParams struct {
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，格式为YYYY-MM-DD。（查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 返回数据的粒度，支持设为以下值：
	// d：按天。此时返回查询时间范围内 UTC 时间为零点的数据。
	// h：按小时。此时返回查询时间范围内 UTC 时间为整小时的数据。
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

type DescribeTRTCMarketQualityMetricDataRequest struct {
	*tchttp.BaseRequest
	
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，格式为YYYY-MM-DD。（查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 返回数据的粒度，支持设为以下值：
	// d：按天。此时返回查询时间范围内 UTC 时间为零点的数据。
	// h：按小时。此时返回查询时间范围内 UTC 时间为整小时的数据。
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *DescribeTRTCMarketQualityMetricDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCMarketQualityMetricDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTRTCMarketQualityMetricDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCMarketQualityMetricDataResponseParams struct {
	// TRTC监控数据出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TRTCDataResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTRTCMarketQualityMetricDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTRTCMarketQualityMetricDataResponseParams `json:"Response"`
}

func (r *DescribeTRTCMarketQualityMetricDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCMarketQualityMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCMarketScaleDataRequestParams struct {
	// 用户SdkAppId
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，格式为YYYY-MM-DD。（查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 返回数据的粒度，支持设为以下值：
	// d：按天。此时返回查询时间范围内 UTC 时间为零点的数据。
	// h：按小时。此时返回查询时间范围内 UTC 时间为整小时的数据。
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

type DescribeTRTCMarketScaleDataRequest struct {
	*tchttp.BaseRequest
	
	// 用户SdkAppId
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，格式为YYYY-MM-DD。（查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 返回数据的粒度，支持设为以下值：
	// d：按天。此时返回查询时间范围内 UTC 时间为零点的数据。
	// h：按小时。此时返回查询时间范围内 UTC 时间为整小时的数据。
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *DescribeTRTCMarketScaleDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCMarketScaleDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTRTCMarketScaleDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCMarketScaleDataResponseParams struct {
	// TRTC监控数据出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TRTCDataResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTRTCMarketScaleDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTRTCMarketScaleDataResponseParams `json:"Response"`
}

func (r *DescribeTRTCMarketScaleDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCMarketScaleDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCMarketScaleMetricDataRequestParams struct {
	// 用户SdkAppId
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，格式为YYYY-MM-DD。（查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 返回数据的粒度，支持设为以下值：
	// d：按天。此时返回查询时间范围内 UTC 时间为零点的数据。
	// h：按小时。此时返回查询时间范围内 UTC 时间为整小时的数据。
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

type DescribeTRTCMarketScaleMetricDataRequest struct {
	*tchttp.BaseRequest
	
	// 用户SdkAppId
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，格式为YYYY-MM-DD。（查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 返回数据的粒度，支持设为以下值：
	// d：按天。此时返回查询时间范围内 UTC 时间为零点的数据。
	// h：按小时。此时返回查询时间范围内 UTC 时间为整小时的数据。
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *DescribeTRTCMarketScaleMetricDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCMarketScaleMetricDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTRTCMarketScaleMetricDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCMarketScaleMetricDataResponseParams struct {
	// TRTC监控数据出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TRTCDataResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTRTCMarketScaleMetricDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTRTCMarketScaleMetricDataResponseParams `json:"Response"`
}

func (r *DescribeTRTCMarketScaleMetricDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCMarketScaleMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCRealTimeQualityDataRequestParams struct {
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 开始时间，unix时间戳，单位：秒（查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间戳，单位：秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 房间ID
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeTRTCRealTimeQualityDataRequest struct {
	*tchttp.BaseRequest
	
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 开始时间，unix时间戳，单位：秒（查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间戳，单位：秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 房间ID
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DescribeTRTCRealTimeQualityDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCRealTimeQualityDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTRTCRealTimeQualityDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCRealTimeQualityDataResponseParams struct {
	// TRTC监控数据出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TRTCDataResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTRTCRealTimeQualityDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTRTCRealTimeQualityDataResponseParams `json:"Response"`
}

func (r *DescribeTRTCRealTimeQualityDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCRealTimeQualityDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCRealTimeQualityMetricDataRequestParams struct {
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 开始时间，unix时间戳，单位：秒（查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间戳，单位：秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 房间ID
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeTRTCRealTimeQualityMetricDataRequest struct {
	*tchttp.BaseRequest
	
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 开始时间，unix时间戳，单位：秒（查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间戳，单位：秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 房间ID
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DescribeTRTCRealTimeQualityMetricDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCRealTimeQualityMetricDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTRTCRealTimeQualityMetricDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCRealTimeQualityMetricDataResponseParams struct {
	// TRTC监控数据出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TRTCDataResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTRTCRealTimeQualityMetricDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTRTCRealTimeQualityMetricDataResponseParams `json:"Response"`
}

func (r *DescribeTRTCRealTimeQualityMetricDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCRealTimeQualityMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCRealTimeScaleDataRequestParams struct {
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 开始时间，unix时间戳，单位：秒（查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间戳，单位：秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 房间ID
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeTRTCRealTimeScaleDataRequest struct {
	*tchttp.BaseRequest
	
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 开始时间，unix时间戳，单位：秒（查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间戳，单位：秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 房间ID
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DescribeTRTCRealTimeScaleDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCRealTimeScaleDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTRTCRealTimeScaleDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCRealTimeScaleDataResponseParams struct {
	// TRTC监控数据出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TRTCDataResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTRTCRealTimeScaleDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTRTCRealTimeScaleDataResponseParams `json:"Response"`
}

func (r *DescribeTRTCRealTimeScaleDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCRealTimeScaleDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCRealTimeScaleMetricDataRequestParams struct {
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 开始时间，unix时间戳，单位：秒（查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间戳，单位：秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 房间ID
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeTRTCRealTimeScaleMetricDataRequest struct {
	*tchttp.BaseRequest
	
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 开始时间，unix时间戳，单位：秒（查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间戳，单位：秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 房间ID
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DescribeTRTCRealTimeScaleMetricDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCRealTimeScaleMetricDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTRTCRealTimeScaleMetricDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCRealTimeScaleMetricDataResponseParams struct {
	// TRTC监控数据出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TRTCDataResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTRTCRealTimeScaleMetricDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTRTCRealTimeScaleMetricDataResponseParams `json:"Response"`
}

func (r *DescribeTRTCRealTimeScaleMetricDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCRealTimeScaleMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrtcMcuTranscodeTimeRequestParams struct {
	// 查询开始时间，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用ID，可不传。传应用ID时返回的是该应用的用量，不传时返回多个应用的合计值。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeTrtcMcuTranscodeTimeRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用ID，可不传。传应用ID时返回的是该应用的用量，不传时返回多个应用的合计值。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeTrtcMcuTranscodeTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcMcuTranscodeTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrtcMcuTranscodeTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrtcMcuTranscodeTimeResponseParams struct {
	// 应用的用量信息数组。
	Usages []*OneSdkAppIdTranscodeTimeUsagesInfo `json:"Usages,omitnil,omitempty" name:"Usages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrtcMcuTranscodeTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrtcMcuTranscodeTimeResponseParams `json:"Response"`
}

func (r *DescribeTrtcMcuTranscodeTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcMcuTranscodeTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrtcRoomUsageRequestParams struct {
	// TRTC的SdkAppId，和房间所对应的SdkAppId相同。
	SdkAppid *uint64 `json:"SdkAppid,omitnil,omitempty" name:"SdkAppid"`

	// 查询开始时间，格式为YYYY-MM-DD HH:MM，精确到分钟级。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD HH:MM，单次查询不超过24h。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeTrtcRoomUsageRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SdkAppId，和房间所对应的SdkAppId相同。
	SdkAppid *uint64 `json:"SdkAppid,omitnil,omitempty" name:"SdkAppid"`

	// 查询开始时间，格式为YYYY-MM-DD HH:MM，精确到分钟级。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD HH:MM，单次查询不超过24h。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeTrtcRoomUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcRoomUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppid")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrtcRoomUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrtcRoomUsageResponseParams struct {
	// 房间维度用量数据，csv文件格式，单位：秒。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrtcRoomUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrtcRoomUsageResponseParams `json:"Response"`
}

func (r *DescribeTrtcRoomUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcRoomUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrtcUsageRequestParams struct {
	// 查询开始时间，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// TRTC的SdkAppId，和房间所对应的SdkAppId相同。如果没有这个参数，返回用户下全部实时音视频应用的汇总。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeTrtcUsageRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// TRTC的SdkAppId，和房间所对应的SdkAppId相同。如果没有这个参数，返回用户下全部实时音视频应用的汇总。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeTrtcUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrtcUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrtcUsageResponseParams struct {
	// 用量类型，与UsageValue中各个位置的值对应。
	UsageKey []*string `json:"UsageKey,omitnil,omitempty" name:"UsageKey"`

	// 各个时间点用量明细，单位:分钟
	UsageList []*TrtcUsage `json:"UsageList,omitnil,omitempty" name:"UsageList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrtcUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrtcUsageResponseParams `json:"Response"`
}

func (r *DescribeTrtcUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnusualEventRequestParams struct {
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，本地unix时间戳，单位为秒（如：1590065777）
	// 注意：支持查询14天内的数据
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳，单位为秒（如：1590065877）注意：与StartTime间隔时间不超过1小时。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 房间号，查询房间内任意20条以内异常体验事件
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DescribeUnusualEventRequest struct {
	*tchttp.BaseRequest
	
	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 查询开始时间，本地unix时间戳，单位为秒（如：1590065777）
	// 注意：支持查询14天内的数据
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳，单位为秒（如：1590065877）注意：与StartTime间隔时间不超过1小时。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 房间号，查询房间内任意20条以内异常体验事件
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DescribeUnusualEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnusualEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnusualEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnusualEventResponseParams struct {
	// 返回的数据总条数
	// 范围：[0，20]
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 异常体验列表
	AbnormalExperienceList []*AbnormalExperience `json:"AbnormalExperienceList,omitnil,omitempty" name:"AbnormalExperienceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUnusualEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnusualEventResponseParams `json:"Response"`
}

func (r *DescribeUnusualEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnusualEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserEventRequestParams struct {
	// 通话 ID（唯一标识一次通话）： SdkAppId_RoomId（房间号）_ CreateTime（房间创建时间，unix时间戳，单位为s）例：1400xxxxxx_218695_1590065777。通过 DescribeRoomInfo（查询历史房间列表）接口获取（[查询历史房间列表](https://cloud.tencent.com/document/product/647/44050)）。
	CommId *string `json:"CommId,omitnil,omitempty" name:"CommId"`

	// 查询开始时间，本地unix时间戳，单位为秒（如：1590065777）
	// 注意：支持查询14天内的数据
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳，单位为秒（如：1590065877）
	// 注意：查询时间大于房间结束时间，以房间结束时间为准。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户UserId
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 房间号（如：223）
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeUserEventRequest struct {
	*tchttp.BaseRequest
	
	// 通话 ID（唯一标识一次通话）： SdkAppId_RoomId（房间号）_ CreateTime（房间创建时间，unix时间戳，单位为s）例：1400xxxxxx_218695_1590065777。通过 DescribeRoomInfo（查询历史房间列表）接口获取（[查询历史房间列表](https://cloud.tencent.com/document/product/647/44050)）。
	CommId *string `json:"CommId,omitnil,omitempty" name:"CommId"`

	// 查询开始时间，本地unix时间戳，单位为秒（如：1590065777）
	// 注意：支持查询14天内的数据
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳，单位为秒（如：1590065877）
	// 注意：查询时间大于房间结束时间，以房间结束时间为准。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户UserId
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 房间号（如：223）
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeUserEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserId")
	delete(f, "RoomId")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserEventResponseParams struct {
	// 返回的事件列表，若没有数据，会返回空数组。
	Data []*EventList `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserEventResponseParams `json:"Response"`
}

func (r *DescribeUserEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInfoRequestParams struct {
	// 通话 ID（唯一标识一次通话）： SdkAppId_RoomId（房间号）_ CreateTime（房间创建时间，unix时间戳，单位为s）例：1400xxxxxx_218695_1590065777。通过 DescribeRoomInfo（查询历史房间列表）接口获取（[查询历史房间列表](https://cloud.tencent.com/document/product/647/44050)）。
	CommId *string `json:"CommId,omitnil,omitempty" name:"CommId"`

	// 查询开始时间，本地unix时间戳，单位为秒（如：1590065777）注意：最大支持查询14天内的数据
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳，单位为秒（如：1590065877）
	// 注意：与StartTime间隔时间不超过4小时。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需查询的用户数组，不填默认返回6个用户
	// 范围：[1，100]。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 当前页数，默认为0，
	// 注意：PageNumber和PageSize 其中一个不填均默认返回6条数据。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页个数，默认为6，
	// 范围：[1，100]。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// 通话 ID（唯一标识一次通话）： SdkAppId_RoomId（房间号）_ CreateTime（房间创建时间，unix时间戳，单位为s）例：1400xxxxxx_218695_1590065777。通过 DescribeRoomInfo（查询历史房间列表）接口获取（[查询历史房间列表](https://cloud.tencent.com/document/product/647/44050)）。
	CommId *string `json:"CommId,omitnil,omitempty" name:"CommId"`

	// 查询开始时间，本地unix时间戳，单位为秒（如：1590065777）注意：最大支持查询14天内的数据
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳，单位为秒（如：1590065877）
	// 注意：与StartTime间隔时间不超过4小时。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户SdkAppId（如：1400xxxxxx）
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需查询的用户数组，不填默认返回6个用户
	// 范围：[1，100]。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 当前页数，默认为0，
	// 注意：PageNumber和PageSize 其中一个不填均默认返回6条数据。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页个数，默认为6，
	// 范围：[1，100]。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "UserIds")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInfoResponseParams struct {
	// 返回的用户总条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 用户信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserList []*UserInformation `json:"UserList,omitnil,omitempty" name:"UserList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserInfoResponseParams `json:"Response"`
}

func (r *DescribeUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebRecordRequestParams struct {
	// 开始页面录制时返回的任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 发起页面录制时传递的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 发起录制时传递的RecordId, 传入此值时需要传递SdkAppId
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`
}

type DescribeWebRecordRequest struct {
	*tchttp.BaseRequest
	
	// 开始页面录制时返回的任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 发起页面录制时传递的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 发起录制时传递的RecordId, 传入此值时需要传递SdkAppId
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`
}

func (r *DescribeWebRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "SdkAppId")
	delete(f, "RecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebRecordResponseParams struct {
	// 1: 正在录制中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 在使用RecordId查询时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 在使用TaskId查询时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWebRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebRecordResponseParams `json:"Response"`
}

func (r *DescribeWebRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismissRoomByStrRoomIdRequestParams struct {
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 字符串类型房间号。
	// 本接口仅支持解散字符串类型房间号，如需解散数字类型房间号，请使用：DismissRoom
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DismissRoomByStrRoomIdRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 字符串类型房间号。
	// 本接口仅支持解散字符串类型房间号，如需解散数字类型房间号，请使用：DismissRoom
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DismissRoomByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DismissRoomByStrRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismissRoomByStrRoomIdResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DismissRoomByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *DismissRoomByStrRoomIdResponseParams `json:"Response"`
}

func (r *DismissRoomByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismissRoomRequestParams struct {
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 数字房间号。本接口仅支持解散数字类型房间号，如需解散字符串类型房间号，请使用DismissRoomByStrRoomId。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DismissRoomRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 数字房间号。本接口仅支持解散数字类型房间号，如需解散字符串类型房间号，请使用DismissRoomByStrRoomId。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DismissRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DismissRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismissRoomResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DismissRoomResponse struct {
	*tchttp.BaseResponse
	Response *DismissRoomResponseParams `json:"Response"`
}

func (r *DismissRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EncodeParams struct {
	// 混流-输出流音频采样率。取值为[48000, 44100, 32000, 24000, 16000, 8000]，单位是Hz。混流任务发起过程中，为了保持CDN链接的稳定，不要修改音频参数（codec、采样率、码率、声道数）。
	AudioSampleRate *uint64 `json:"AudioSampleRate,omitnil,omitempty" name:"AudioSampleRate"`

	// 混流-输出流音频码率。取值范围[8,500]，单位为kbps。混流任务发起过程中，为了保持CDN链接的稳定，不要修改音频参数（codec、采样率、码率、声道数）。
	AudioBitrate *uint64 `json:"AudioBitrate,omitnil,omitempty" name:"AudioBitrate"`

	// 混流-输出流音频声道数，取值范围[1,2]，1表示混流输出音频为单声道，2表示混流输出音频为双声道。混流任务发起过程中，为了保持CDN链接的稳定，不要修改音频参数（codec、采样率、码率、声道数）。
	AudioChannels *uint64 `json:"AudioChannels,omitnil,omitempty" name:"AudioChannels"`

	// 混流-输出流宽，音视频输出时必填。取值范围[0,1920]，单位为像素值。
	VideoWidth *uint64 `json:"VideoWidth,omitnil,omitempty" name:"VideoWidth"`

	// 混流-输出流高，音视频输出时必填。取值范围[0,1080]，单位为像素值。
	VideoHeight *uint64 `json:"VideoHeight,omitnil,omitempty" name:"VideoHeight"`

	// 混流-输出流码率，音视频输出时必填。取值范围[1,10000]，单位为kbps。
	VideoBitrate *uint64 `json:"VideoBitrate,omitnil,omitempty" name:"VideoBitrate"`

	// 混流-输出流帧率，音视频输出时必填。取值范围[1,60]，表示混流的输出帧率可选范围为1到60fps。
	VideoFramerate *uint64 `json:"VideoFramerate,omitnil,omitempty" name:"VideoFramerate"`

	// 混流-输出流gop，音视频输出时必填。取值范围[1,5]，单位为秒。
	VideoGop *uint64 `json:"VideoGop,omitnil,omitempty" name:"VideoGop"`

	// 混流-输出流背景色，取值是十进制整数。常用的颜色有：
	// 红色：0xff0000，对应的十进制整数是16724736。
	// 黄色：0xffff00。对应的十进制整数是16776960。
	// 绿色：0x33cc00。对应的十进制整数是3394560。
	// 蓝色：0x0066ff。对应的十进制整数是26367。
	// 黑色：0x000000。对应的十进制整数是0。
	// 白色：0xFFFFFF。对应的十进制整数是16777215。
	// 灰色：0x999999。对应的十进制整数是10066329。
	BackgroundColor *uint64 `json:"BackgroundColor,omitnil,omitempty" name:"BackgroundColor"`

	// 混流-输出流背景图片，取值为实时音视频控制台上传的图片ID。
	BackgroundImageId *uint64 `json:"BackgroundImageId,omitnil,omitempty" name:"BackgroundImageId"`

	// 混流-输出流音频编码类型，取值范围[0,1, 2]，0为LC-AAC，1为HE-AAC，2为HE-AACv2。默认值为0。当音频编码设置为HE-AACv2时，只支持输出流音频声道数为双声道。HE-AAC和HE-AACv2支持的输出流音频采样率范围为[48000, 44100, 32000, 24000, 16000]。混流任务发起过程中，为了保持CDN链接的稳定，不要修改音频参数（codec、采样率、码率、声道数）。
	AudioCodec *uint64 `json:"AudioCodec,omitnil,omitempty" name:"AudioCodec"`

	// 混流-输出流背景图片URL地址，支持png、jpg、jpeg、bmp格式，暂不支持透明通道。URL链接长度限制为512字节。BackgroundImageUrl和BackgroundImageId参数都填时，以BackgroundImageUrl为准。图片大小限制不超过2MB。
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitnil,omitempty" name:"BackgroundImageUrl"`
}

type EventList struct {
	// 数据内容
	Content []*EventMessage `json:"Content,omitnil,omitempty" name:"Content"`

	// 发送端的userId
	PeerId *string `json:"PeerId,omitnil,omitempty" name:"PeerId"`
}

type EventMessage struct {
	// 视频流类型：
	// 0：与视频无关的事件；
	// 2：视频为大画面；
	// 3：视频为小画面；
	// 7：视频为旁路画面；
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件上报的时间戳，unix时间（1589891188801ms)
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 事件Id：分为sdk的事件和webrtc的事件，详情见：附录/事件 ID 映射表：https://cloud.tencent.com/document/product/647/44916
	EventId *uint64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 事件的第一个参数，如视频分辨率宽
	ParamOne *int64 `json:"ParamOne,omitnil,omitempty" name:"ParamOne"`

	// 事件的第二个参数，如视频分辨率高
	ParamTwo *int64 `json:"ParamTwo,omitnil,omitempty" name:"ParamTwo"`
}

type LayoutParams struct {
	// 混流布局模板ID，0为悬浮模板(默认);1为九宫格模板;2为屏幕分享模板;3为画中画模板;4为自定义模板。
	Template *uint64 `json:"Template,omitnil,omitempty" name:"Template"`

	// 屏幕分享模板、悬浮模板、画中画模板中有效，代表大画面对应的用户ID。
	MainVideoUserId *string `json:"MainVideoUserId,omitnil,omitempty" name:"MainVideoUserId"`

	// 屏幕分享模板、悬浮模板、画中画模板中有效，代表大画面对应的流类型，0为摄像头，1为屏幕分享。左侧大画面为web用户时此值填0。
	MainVideoStreamType *uint64 `json:"MainVideoStreamType,omitnil,omitempty" name:"MainVideoStreamType"`

	// 画中画模板中有效，代表小画面的布局参数。
	SmallVideoLayoutParams *SmallVideoLayoutParams `json:"SmallVideoLayoutParams,omitnil,omitempty" name:"SmallVideoLayoutParams"`

	// 屏幕分享模板有效。设置为1时代表大画面居右，小画面居左布局。默认为0。
	MainVideoRightAlign *uint64 `json:"MainVideoRightAlign,omitnil,omitempty" name:"MainVideoRightAlign"`

	// 指定混视频的用户ID列表。设置此参数后，输出流混合此参数中包含用户的音视频，以及其他用户的纯音频。悬浮模板、九宫格、屏幕分享模板有效，最多可设置16个用户。
	MixVideoUids []*string `json:"MixVideoUids,omitnil,omitempty" name:"MixVideoUids"`

	// 自定义模板中有效，指定用户视频在混合画面中的位置。
	PresetLayoutConfig []*PresetLayoutConfig `json:"PresetLayoutConfig,omitnil,omitempty" name:"PresetLayoutConfig"`

	// 自定义模板中有效，设置为1时代表启用占位图功能，0时代表不启用占位图功能，默认为0。启用占位图功能时，在预设位置的用户没有上行视频时可显示对应的占位图。
	PlaceHolderMode *uint64 `json:"PlaceHolderMode,omitnil,omitempty" name:"PlaceHolderMode"`

	// 悬浮模板、九宫格、屏幕分享模板生效，用于控制纯音频上行是否占用画面布局位置。设置为0是代表后台默认处理方式，悬浮小画面占布局位置，九宫格画面占布局位置、屏幕分享小画面不占布局位置；设置为1时代表纯音频上行占布局位置；设置为2时代表纯音频上行不占布局位置。默认为0。
	PureAudioHoldPlaceMode *uint64 `json:"PureAudioHoldPlaceMode,omitnil,omitempty" name:"PureAudioHoldPlaceMode"`

	// 水印参数。
	WaterMarkParams *WaterMarkParams `json:"WaterMarkParams,omitnil,omitempty" name:"WaterMarkParams"`

	// 屏幕分享模板、悬浮模板、九宫格模板、画中画模版有效，画面在输出时的显示模式：0为裁剪，1为缩放，2为缩放并显示黑底，不填采用后台的默认渲染方式（屏幕分享大画面为缩放，其他为裁剪）。若此参数不生效，请提交工单寻求帮助。
	RenderMode *uint64 `json:"RenderMode,omitnil,omitempty" name:"RenderMode"`
}

type MaxVideoUser struct {
	// 用户媒体流参数。
	UserMediaStream *UserMediaStream `json:"UserMediaStream,omitnil,omitempty" name:"UserMediaStream"`
}

type McuAudioParams struct {
	// 音频编码参数。
	AudioEncode *AudioEncode `json:"AudioEncode,omitnil,omitempty" name:"AudioEncode"`

	// 音频用户白名单，start时，为空或不填表示混所有主播音频，填具体值表示混指定主播音频；update时，不填表示不更新，为空表示更新为混所有主播音频，填具体值表示更新为混指定主播音频。
	// 使用黑白名单时，黑白名单必须同时填写。都不填写时表示不更新。同一个用户同时在黑白名单时，以黑名单为主。
	SubscribeAudioList []*McuUserInfoParams `json:"SubscribeAudioList,omitnil,omitempty" name:"SubscribeAudioList"`

	// 音频用户黑名单，为空或不填表示无黑名单，填具体值表示不混指定主播音频。update时，不填表示不更新，为空表示更新为清空黑名单，填具体值表示更新为不混指定主播音频。
	// 使用黑白名单时，黑白名单必须同时填写。都不填写时表示不更新。同一个用户同时在黑白名单时，以黑名单为主。
	UnSubscribeAudioList []*McuUserInfoParams `json:"UnSubscribeAudioList,omitnil,omitempty" name:"UnSubscribeAudioList"`
}

type McuBackgroundCustomRender struct {
	// 自定义渲染画面的宽度，单位为像素值，需大于0，且不能超过子布局的宽。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 自定义渲染画面的高度，单位为像素值，需大于0，且不能超过子布局的高。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 自定义渲染画面的圆角半径，单位为像素值，不能超过渲染画面Width和Height最小值的一半，不指定默认为0，表示直角。
	Radius *uint64 `json:"Radius,omitnil,omitempty" name:"Radius"`
}

type McuCustomCrop struct {
	// 自定义裁剪起始位置的X偏移，单位为像素值，大于等于0。
	LocationX *uint64 `json:"LocationX,omitnil,omitempty" name:"LocationX"`

	// 自定义裁剪起始位置的Y偏移，单位为像素值，大于等于0。
	LocationY *uint64 `json:"LocationY,omitnil,omitempty" name:"LocationY"`

	// 自定义裁剪画面的宽度，单位为像素值，大于0，且LocationX+Width不超过10000
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 自定义裁剪画面的高度，单位为像素值，大于0，且LocationY+Height不超过10000
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type McuFeedBackRoomParams struct {
	// 回推房间的RoomId。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 房间类型，必须和回推房间所对应的RoomId类型相同，0为整形房间号，1为字符串房间号。
	RoomIdType *uint64 `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`

	// 回推房间使用的UserId(https://cloud.tencent.com/document/product/647/46351#userid)，注意这个userId不能与其他TRTC或者转推服务等已经使用的UserId重复，建议可以把房间ID作为userId的标识的一部分。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 回推房间UserId对应的用户签名，相当于登录密码，具体计算方法请参考TRTC计算[UserSig](https://cloud.tencent.com/document/product/647/45910#UserSig)的方案。
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`
}

type McuLayout struct {
	// 用户媒体流参数。不填时腾讯云后台按照上行主播的进房顺序自动填充。
	UserMediaStream *UserMediaStream `json:"UserMediaStream,omitnil,omitempty" name:"UserMediaStream"`

	// 子画面在输出时的宽度，单位为像素值，不填默认为0。
	ImageWidth *uint64 `json:"ImageWidth,omitnil,omitempty" name:"ImageWidth"`

	// 子画面在输出时的高度，单位为像素值，不填默认为0。
	ImageHeight *uint64 `json:"ImageHeight,omitnil,omitempty" name:"ImageHeight"`

	// 子画面在输出时的X偏移，单位为像素值，LocationX与ImageWidth之和不能超过混流输出的总宽度，不填默认为0。
	LocationX *uint64 `json:"LocationX,omitnil,omitempty" name:"LocationX"`

	// 子画面在输出时的Y偏移，单位为像素值，LocationY与ImageHeight之和不能超过混流输出的总高度，不填默认为0。
	LocationY *uint64 `json:"LocationY,omitnil,omitempty" name:"LocationY"`

	// 子画面在输出时的层级，不填默认为0。
	ZOrder *uint64 `json:"ZOrder,omitnil,omitempty" name:"ZOrder"`

	// 子画面在输出时的显示模式：0为裁剪，1为缩放并显示背景，2为缩放并显示黑底。不填默认为0。
	RenderMode *uint64 `json:"RenderMode,omitnil,omitempty" name:"RenderMode"`

	// 【此参数配置无效，暂不支持】子画面的背景颜色，常用的颜色有：
	// 红色：0xcc0033。
	// 黄色：0xcc9900。
	// 绿色：0xcccc33。
	// 蓝色：0x99CCFF。
	// 黑色：0x000000。
	// 白色：0xFFFFFF。
	// 灰色：0x999999。
	BackGroundColor *string `json:"BackGroundColor,omitnil,omitempty" name:"BackGroundColor"`

	// 子画面的背景图url，填写该参数，当用户关闭摄像头或未进入TRTC房间时，会在布局位置填充为指定图片。若指定图片与布局位置尺寸比例不一致，则会对图片进行拉伸处理，优先级高于BackGroundColor。支持png、jpg、jpeg、bmp、gif、webm格式。图片大小限制不超过5MB。
	// 注：您需要确保图片链接的可访问性，后台单次下载超时时间为10秒，最多重试3次，若最终图片下载失败，子画面的背景图将不会生效。
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitnil,omitempty" name:"BackgroundImageUrl"`

	// 客户自定义裁剪，针对原始输入流裁剪
	CustomCrop *McuCustomCrop `json:"CustomCrop,omitnil,omitempty" name:"CustomCrop"`

	// 子背景图在输出时的显示模式：0为裁剪，1为缩放并显示背景，2为缩放并显示黑底，3为变比例伸缩，4为自定义渲染。不填默认为3。
	BackgroundRenderMode *uint64 `json:"BackgroundRenderMode,omitnil,omitempty" name:"BackgroundRenderMode"`

	// 子画面的透明模版url，指向一张包含透明通道的模板图片。填写该参数，后台混流时会提取该模板图片的透明通道，将其缩放作为目标画面的透明通道，再和其他画面进行混合。您可以通过透明模版实现目标画面的半透明效果和任意形状裁剪（如圆角、星形、心形等）。 支持png格式。图片大小限制不超过5MB。
	// 注：1，模板图片宽高比应接近目标画面宽高比，以避免缩放适配目标画面时出现模板效果变形；2，透明模版只有RenderMode为0（裁剪）时才生效；3，您需要确保图片链接的可访问性，后台单次下载超时时间为10秒，最多重试3次，若最终图片下载失败，透明模版将不会生效。
	TransparentUrl *string `json:"TransparentUrl,omitnil,omitempty" name:"TransparentUrl"`

	// 子背景图的自定义渲染参数，当BackgroundRenderMode为4时必须配置。
	BackgroundCustomRender *McuBackgroundCustomRender `json:"BackgroundCustomRender,omitnil,omitempty" name:"BackgroundCustomRender"`

	// 子背景色生效模式，默认值为0表示均不生效。
	// bit0:占位图缩放是否生效。
	// bit1:上行流缩放是否生效。
	// 您可以将相应bit位置1启动生效，例如：
	// 0(00)表示子背景色不生效。
	// 1(01)表示子背景色只在占位图缩放时生效。
	// 2(10)表示子背景色只在上行流缩放时生效。
	// 3(11)表示子背景色在占位图缩放和上行流缩放时均生效。
	BackGroundColorMode *uint64 `json:"BackGroundColorMode,omitnil,omitempty" name:"BackGroundColorMode"`
}

type McuLayoutParams struct {
	// 布局模式：动态布局（1：悬浮布局（默认），2：屏幕分享布局，3：九宫格布局），静态布局（4：自定义布局）。
	MixLayoutMode *uint64 `json:"MixLayoutMode,omitnil,omitempty" name:"MixLayoutMode"`

	// 纯音频上行是否占布局位置，只在动态布局中有效。0表示纯音频不占布局位置，1表示纯音频占布局位置，不填默认为0。
	PureAudioHoldPlaceMode *uint64 `json:"PureAudioHoldPlaceMode,omitnil,omitempty" name:"PureAudioHoldPlaceMode"`

	// 自定义模板中有效，指定用户视频在混合画面中的位置。
	MixLayoutList []*McuLayout `json:"MixLayoutList,omitnil,omitempty" name:"MixLayoutList"`

	// 指定动态布局中悬浮布局和屏幕分享布局的大画面信息，只在悬浮布局和屏幕分享布局有效。
	MaxVideoUser *MaxVideoUser `json:"MaxVideoUser,omitnil,omitempty" name:"MaxVideoUser"`

	// 屏幕分享模板、悬浮模板、九宫格模版有效，画面在输出时的显示模式：0为裁剪，1为缩放，2为缩放并显示黑底
	RenderMode *uint64 `json:"RenderMode,omitnil,omitempty" name:"RenderMode"`
}

type McuLayoutVolume struct {
	// AppData的内容，会被写入自定义SEI中的app_data字段，长度需小于4096。
	AppData *string `json:"AppData,omitnil,omitempty" name:"AppData"`

	// SEI消息的payload_type，默认值100，取值范围100-254（244除外，244为我们内部自定义的时间戳SEI）
	PayloadType *uint64 `json:"PayloadType,omitnil,omitempty" name:"PayloadType"`

	// SEI发送间隔，单位毫秒，默认值为1000。
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 取值范围[0,1]，填1：发送关键帧时会确保带SEI；填0：发送关键帧时不确保带SEI。默认值为0。
	FollowIdr *uint64 `json:"FollowIdr,omitnil,omitempty" name:"FollowIdr"`
}

type McuPassThrough struct {
	// 透传SEI的payload内容。
	PayloadContent *string `json:"PayloadContent,omitnil,omitempty" name:"PayloadContent"`

	// SEI消息的payload_type，取值范围5、100-254（244除外，244为我们内部自定义的时间戳SEI）。
	PayloadType *uint64 `json:"PayloadType,omitnil,omitempty" name:"PayloadType"`

	// PayloadType为5，PayloadUuid必须填写。PayloadType不是5时，不需要填写，填写会被后台忽略。该值必须是32长度的十六进制。
	PayloadUuid *string `json:"PayloadUuid,omitnil,omitempty" name:"PayloadUuid"`

	// SEI发送间隔，单位毫秒，默认值为1000。
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 取值范围[0,1]，填1：发送关键帧时会确保带SEI；填0：发送关键帧时不确保带SEI。默认值为0。
	FollowIdr *uint64 `json:"FollowIdr,omitnil,omitempty" name:"FollowIdr"`
}

type McuPublishCdnParam struct {
	// CDN转推URL，只支持rtmp链接。
	// 注：若更新转推时，URL有任何变化，都会断流重推。
	PublishCdnUrl *string `json:"PublishCdnUrl,omitnil,omitempty" name:"PublishCdnUrl"`

	// 是否是腾讯云CDN，0为转推非腾讯云CDN，1为转推腾讯CDN，不携带该参数默认为1。注意：1，为避免误产生转推费用，该参数建议明确填写，转推非腾讯云CDN时会产生转推费用，详情参见接口文档说明；2，国内站默认只支持转推腾讯云CDN，如您有转推第三方CDN需求，请联系腾讯云技术支持。
	IsTencentCdn *uint64 `json:"IsTencentCdn,omitnil,omitempty" name:"IsTencentCdn"`
}

type McuSeiParams struct {
	// 音量布局SEI
	LayoutVolume *McuLayoutVolume `json:"LayoutVolume,omitnil,omitempty" name:"LayoutVolume"`

	// 透传SEI
	PassThrough *McuPassThrough `json:"PassThrough,omitnil,omitempty" name:"PassThrough"`
}

type McuUserInfoParams struct {
	// 用户参数。
	UserInfo *MixUserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 混音的音量调整：取值范围是0到100，100为原始上行音量，不填默认为100，值越小则音量越低。
	// 注：该参数只在音量白名单下配置生效，其他场景配置无效。
	SoundLevel *uint64 `json:"SoundLevel,omitnil,omitempty" name:"SoundLevel"`
}

type McuVideoParams struct {
	// 输出流视频编码参数。
	VideoEncode *VideoEncode `json:"VideoEncode,omitnil,omitempty" name:"VideoEncode"`

	// 混流布局参数。
	LayoutParams *McuLayoutParams `json:"LayoutParams,omitnil,omitempty" name:"LayoutParams"`

	// 整个画布背景颜色，常用的颜色有：
	// 红色：0xcc0033。
	// 黄色：0xcc9900。
	// 绿色：0xcccc33。
	// 蓝色：0x99CCFF。
	// 黑色：0x000000。
	// 白色：0xFFFFFF。
	// 灰色：0x999999。
	BackGroundColor *string `json:"BackGroundColor,omitnil,omitempty" name:"BackGroundColor"`

	// 整个画布的背景图url，优先级高于BackGroundColor。支持png、jpg、jpeg格式。图片大小限制不超过5MB。
	// 注：您需要确保图片链接的可访问性，后台单次下载超时时间为10秒，最多重试3次，若最终图片下载失败，背景图将不会生效。
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitnil,omitempty" name:"BackgroundImageUrl"`

	// 混流布局的水印参数。
	WaterMarkList []*McuWaterMarkParams `json:"WaterMarkList,omitnil,omitempty" name:"WaterMarkList"`

	// 背景图在输出时的显示模式：0为裁剪，1为缩放并显示黑底，2为变比例伸缩。后台默认为变比例伸缩。
	BackgroundRenderMode *uint64 `json:"BackgroundRenderMode,omitnil,omitempty" name:"BackgroundRenderMode"`
}

type McuWaterMarkImage struct {
	// 水印图片URL地址，支持png、jpg、jpeg格式。图片大小限制不超过5MB。
	// 注：您需要确保图片链接的可访问性，后台单次下载超时时间为10秒，最多重试3次，若最终图片下载失败，水印图片将不会生效。
	WaterMarkUrl *string `json:"WaterMarkUrl,omitnil,omitempty" name:"WaterMarkUrl"`

	// 水印在输出时的宽。单位为像素值。
	WaterMarkWidth *uint64 `json:"WaterMarkWidth,omitnil,omitempty" name:"WaterMarkWidth"`

	// 水印在输出时的高。单位为像素值。
	WaterMarkHeight *uint64 `json:"WaterMarkHeight,omitnil,omitempty" name:"WaterMarkHeight"`

	// 水印在输出时的X偏移。单位为像素值。
	LocationX *uint64 `json:"LocationX,omitnil,omitempty" name:"LocationX"`

	// 水印在输出时的Y偏移。单位为像素值。
	LocationY *uint64 `json:"LocationY,omitnil,omitempty" name:"LocationY"`

	// 水印在输出时的层级，不填默认为0。
	ZOrder *uint64 `json:"ZOrder,omitnil,omitempty" name:"ZOrder"`

	// 动态水印类型，默认为0。0:关闭；1:随机位置，每秒变动一次；2:边界扫描反弹，每帧变动一次。
	DynamicPosType *uint64 `json:"DynamicPosType,omitnil,omitempty" name:"DynamicPosType"`
}

type McuWaterMarkParams struct {
	// 水印类型，0为图片（默认），1为文字。
	WaterMarkType *uint64 `json:"WaterMarkType,omitnil,omitempty" name:"WaterMarkType"`

	// 图片水印参数。WaterMarkType为0指定。
	WaterMarkImage *McuWaterMarkImage `json:"WaterMarkImage,omitnil,omitempty" name:"WaterMarkImage"`

	// 文字水印参数。WaterMarkType为1指定。
	WaterMarkText *McuWaterMarkText `json:"WaterMarkText,omitnil,omitempty" name:"WaterMarkText"`
}

type McuWaterMarkText struct {
	// 文字水印内容。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 水印在输出时的宽。单位为像素值。
	WaterMarkWidth *uint64 `json:"WaterMarkWidth,omitnil,omitempty" name:"WaterMarkWidth"`

	// 水印在输出时的高。单位为像素值。
	WaterMarkHeight *uint64 `json:"WaterMarkHeight,omitnil,omitempty" name:"WaterMarkHeight"`

	// 水印在输出时的X偏移。单位为像素值。
	LocationX *uint64 `json:"LocationX,omitnil,omitempty" name:"LocationX"`

	// 水印在输出时的Y偏移。单位为像素值。
	LocationY *uint64 `json:"LocationY,omitnil,omitempty" name:"LocationY"`

	// 字体大小
	FontSize *uint64 `json:"FontSize,omitnil,omitempty" name:"FontSize"`

	// 字体颜色，默认为白色。常用的颜色有： 红色：0xcc0033。 黄色：0xcc9900。 绿色：0xcccc33。 蓝色：0x99CCFF。 黑色：0x000000。 白色：0xFFFFFF。 灰色：0x999999。	
	FontColor *string `json:"FontColor,omitnil,omitempty" name:"FontColor"`

	// 字体背景色，不配置默认为透明。常用的颜色有： 红色：0xcc0033。 黄色：0xcc9900。 绿色：0xcccc33。 蓝色：0x99CCFF。 黑色：0x000000。 白色：0xFFFFFF。 灰色：0x999999。	
	BackGroundColor *string `json:"BackGroundColor,omitnil,omitempty" name:"BackGroundColor"`

	// 动态水印类型，默认为0。0:关闭；1:随机位置，每秒变动一次；2:边界扫描反弹，每帧变动一次。
	DynamicPosType *uint64 `json:"DynamicPosType,omitnil,omitempty" name:"DynamicPosType"`

	// 水印在输出时的层级，不填默认为0。
	ZOrder *uint64 `json:"ZOrder,omitnil,omitempty" name:"ZOrder"`

	// 水印字体，不填默认为Tencent。支持设置以下值： Tencent （默认） SourceHanSans
	Font *string `json:"Font,omitnil,omitempty" name:"Font"`
}

type MixLayout struct {
	// 画布上该画面左上角的 y 轴坐标，取值范围 [0, 1920]，不能超过画布的高。
	Top *uint64 `json:"Top,omitnil,omitempty" name:"Top"`

	// 画布上该画面左上角的 x 轴坐标，取值范围 [0, 1920]，不能超过画布的宽。
	Left *uint64 `json:"Left,omitnil,omitempty" name:"Left"`

	// 画布上该画面宽度的相对值，取值范围 [0, 1920]，与Left相加不应超过画布的宽。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 画布上该画面高度的相对值，取值范围 [0, 1920]，与Top相加不应超过画布的高。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 字符串内容为待显示在该画面的主播对应的UserId，如果不指定，会按照主播加入房间的顺序匹配。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 画布的透明度值，取值范围[0, 255]。0表示不透明，255表示全透明。默认值为0。
	Alpha *uint64 `json:"Alpha,omitnil,omitempty" name:"Alpha"`

	// 0 ：拉伸模式，这个模式下整个视频内容会全部显示，并填满子画面，在源视频和目的视频宽高比不一致的时候，画面不会缺少内容，但是画面可能产生形变；
	// 
	// 1 ：剪裁模式（默认），这个模式下会严格按照目的视频的宽高比对源视频剪裁之后再拉伸，并填满子画面画布，在源视频和目的视频宽高比不一致的时候，画面保持不变形，但是会被剪裁；
	// 
	// 2 ：填黑模式，这个模式下会严格保持源视频的宽高比进行等比缩放，在源视频和目的视频宽高比不一致的时候，画面的上下侧边缘或者左右侧边缘会露出子画面画布的背景；
	// 
	// 3 ：智能拉伸模式，这个模式类似剪裁模式，区别是在源视频和目的视频宽高比不一致的时候，限制了最大剪裁比例为画面的宽度或者高度的20%；
	RenderMode *uint64 `json:"RenderMode,omitnil,omitempty" name:"RenderMode"`

	// 对应订阅流的主辅路标识：
	// 0：主流（默认）；
	// 1：辅流；
	MediaId *uint64 `json:"MediaId,omitnil,omitempty" name:"MediaId"`

	// 该画布的图层顺序, 这个值越小表示图层越靠后。默认值为0。
	ImageLayer *uint64 `json:"ImageLayer,omitnil,omitempty" name:"ImageLayer"`

	// 图片的url地址， 只支持jpg, png, jpeg，大小限制不超过5M。注意，url必须携带格式后缀，url内只支持特定的字符串, 范围是a-z A-Z 0-9 '-', '.', '_', '~', ':', '/', '?', '#', '[', ']' '@', '!', '&', '(', ')', '*', '+', ',', '%', '='
	SubBackgroundImage *string `json:"SubBackgroundImage,omitnil,omitempty" name:"SubBackgroundImage"`
}

type MixLayoutParams struct {
	// 布局模式:
	// 1：悬浮布局；
	// 2：屏幕分享布局；
	// 3：九宫格布局；
	// 4：自定义布局；
	// 
	// 悬浮布局：默认第一个进入房间的主播（也可以指定一个主播）的视频画面会铺满整个屏幕。其他主播的视频画面从左下角开始依次按照进房顺序水平排列，显示为小画面，小画面悬浮于大画面之上。当画面数量小于等于17个时，每行4个（4 x 4排列）。当画面数量大于17个时，重新布局小画面为每行5个（5 x 5）排列。最多支持25个画面，如果用户只发送音频，仍然会占用画面位置。
	// 
	// 屏幕分享布局：指定一个主播在屏幕左侧的大画面位置（如果不指定，那么大画面位置为背景色），其他主播自上而下依次垂直排列于右侧。当画面数量少于17个的时候，右侧每列最多8人，最多占据两列。当画面数量多于17个的时候，超过17个画面的主播从左下角开始依次水平排列。最多支持25个画面，如果主播只发送音频，仍然会占用画面位置。
	// 
	// 九宫格布局：根据主播的数量自动调整每个画面的大小，每个主播的画面大小一致，最多支持25个画面。
	// 
	// 自定义布局：根据需要在MixLayoutList内定制每个主播画面的布局。
	MixLayoutMode *uint64 `json:"MixLayoutMode,omitnil,omitempty" name:"MixLayoutMode"`

	// 如果MixLayoutMode 选择为4自定义布局模式的话，设置此参数为每个主播所对应的布局画面的详细信息，最大不超过25个。
	MixLayoutList []*MixLayout `json:"MixLayoutList,omitnil,omitempty" name:"MixLayoutList"`

	// 录制背景颜色，RGB的颜色表的16进制表示，每个颜色通过8bit长度标识，默认为黑色。比如橙色对应的RGB为 R:255 G:165 B:0, 那么对应的字符串描述为#FFA500，格式规范：‘#‘开头，后面跟固定RGB的颜色值
	BackGroundColor *string `json:"BackGroundColor,omitnil,omitempty" name:"BackGroundColor"`

	// 在布局模式为1：悬浮布局和 2：屏幕分享布局时，设定为显示大视频画面的UserId。不填的话：悬浮布局默认是第一个进房间的主播，屏幕分享布局默认是背景色
	MaxResolutionUserId *string `json:"MaxResolutionUserId,omitnil,omitempty" name:"MaxResolutionUserId"`

	// 主辅路标识，
	// 0：主流（默认）；
	// 1：辅流（屏幕分享）；
	// 这个位置的MediaId代表的是对应MaxResolutionUserId的主辅路，MixLayoutList内代表的是自定义用户的主辅路。
	MediaId *uint64 `json:"MediaId,omitnil,omitempty" name:"MediaId"`

	// 图片的url地址，只支持jpg, png, jpeg，大小限制不超过5M。注意，url必须携带格式后缀，url内只支持特定的字符串, 范围是a-z A-Z 0-9 '-', '.', '_', '~', ':', '/', '?', '#', '[', ']' '@', '!', '&', '(', ')', '*', '+', ',', '%', '='
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitnil,omitempty" name:"BackgroundImageUrl"`

	// 设置为1时代表启用占位图功能，0时代表不启用占位图功能，默认为0。启用占位图功能时，在预设位置的用户没有上行视频时可显示对应的占位图。
	PlaceHolderMode *uint64 `json:"PlaceHolderMode,omitnil,omitempty" name:"PlaceHolderMode"`

	// 背景画面宽高比不一致的时候处理方案，与MixLayoufList定义的RenderMode一致。
	BackgroundImageRenderMode *uint64 `json:"BackgroundImageRenderMode,omitnil,omitempty" name:"BackgroundImageRenderMode"`

	// 子画面占位图url地址，只支持jpg, png, jpeg，大小限制不超过5M。注意，url必须携带格式后缀，url内只支持特定的字符串, 范围是a-z A-Z 0-9 '-', '.', '_', '~', ':', '/', '?', '#', '[', ']' '@', '!', '&', '(', ')', '*', '+', ',', '%', '='
	DefaultSubBackgroundImage *string `json:"DefaultSubBackgroundImage,omitnil,omitempty" name:"DefaultSubBackgroundImage"`

	// 水印布局参数， 最多支持25个。
	WaterMarkList []*WaterMark `json:"WaterMarkList,omitnil,omitempty" name:"WaterMarkList"`

	// 模板布局下，背景画面宽高比不一致的时候处理方案。自定义布局不生效，与MixLayoufList定义的RenderMode一致。
	RenderMode *uint64 `json:"RenderMode,omitnil,omitempty" name:"RenderMode"`

	// 屏幕分享模板有效。设置为1时代表大画面居右，小画面居左布局。默认为0。
	MaxResolutionUserAlign *uint64 `json:"MaxResolutionUserAlign,omitnil,omitempty" name:"MaxResolutionUserAlign"`
}

type MixTranscodeParams struct {
	// 录制视频转码参数，注意如果设置了这个参数，那么里面的字段都是必填的，没有默认值，如果不填这个参数，那么取值为默认值。
	VideoParams *VideoParams `json:"VideoParams,omitnil,omitempty" name:"VideoParams"`

	// 录制音频转码参数，注意如果设置了这个参数，那么里面的字段都是必填的，没有默认值，如果不填这个参数，那么取值为默认值。
	AudioParams *AudioParams `json:"AudioParams,omitnil,omitempty" name:"AudioParams"`
}

type MixUserInfo struct {
	// 用户ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 动态布局时房间信息必须和主房间信息保持一致，自定义布局时房间信息必须和MixLayoutList中对应用户的房间信息保持一致，不填时默认与主房间信息一致。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 房间号类型，0为整型房间号，1为字符串房间号。
	RoomIdType *uint64 `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`
}

// Predefined struct for user
type ModifyCloudRecordingRequestParams struct {
	// TRTC的SDKAppId，和录制的房间所对应的SDKAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 录制任务的唯一Id，在启动录制成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 需要更新的混流的布局参数。
	MixLayoutParams *MixLayoutParams `json:"MixLayoutParams,omitnil,omitempty" name:"MixLayoutParams"`

	// 指定订阅流白名单或者黑名单。
	SubscribeStreamUserIds *SubscribeStreamUserIds `json:"SubscribeStreamUserIds,omitnil,omitempty" name:"SubscribeStreamUserIds"`
}

type ModifyCloudRecordingRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId，和录制的房间所对应的SDKAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 录制任务的唯一Id，在启动录制成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 需要更新的混流的布局参数。
	MixLayoutParams *MixLayoutParams `json:"MixLayoutParams,omitnil,omitempty" name:"MixLayoutParams"`

	// 指定订阅流白名单或者黑名单。
	SubscribeStreamUserIds *SubscribeStreamUserIds `json:"SubscribeStreamUserIds,omitnil,omitempty" name:"SubscribeStreamUserIds"`
}

func (r *ModifyCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	delete(f, "MixLayoutParams")
	delete(f, "SubscribeStreamUserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudRecordingResponseParams struct {
	// 云录制服务分配的任务 ID。任务 ID 是对一次录制生命周期过程的唯一标识，结束录制时会失去意义。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudRecordingResponseParams `json:"Response"`
}

func (r *ModifyCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPictureRequestParams struct {
	// 图片id
	PictureId *uint64 `json:"PictureId,omitnil,omitempty" name:"PictureId"`

	// 应用id
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 图片长度
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 图片宽度
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 显示位置x轴方向
	XPosition *uint64 `json:"XPosition,omitnil,omitempty" name:"XPosition"`

	// 显示位置y轴方向
	YPosition *uint64 `json:"YPosition,omitnil,omitempty" name:"YPosition"`
}

type ModifyPictureRequest struct {
	*tchttp.BaseRequest
	
	// 图片id
	PictureId *uint64 `json:"PictureId,omitnil,omitempty" name:"PictureId"`

	// 应用id
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 图片长度
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 图片宽度
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 显示位置x轴方向
	XPosition *uint64 `json:"XPosition,omitnil,omitempty" name:"XPosition"`

	// 显示位置y轴方向
	YPosition *uint64 `json:"YPosition,omitnil,omitempty" name:"YPosition"`
}

func (r *ModifyPictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PictureId")
	delete(f, "SdkAppId")
	delete(f, "Height")
	delete(f, "Width")
	delete(f, "XPosition")
	delete(f, "YPosition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPictureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPictureResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPictureResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPictureResponseParams `json:"Response"`
}

func (r *ModifyPictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OneSdkAppIdTranscodeTimeUsagesInfo struct {
	// 旁路转码时长查询结果数组
	SdkAppIdTranscodeTimeUsages []*SdkAppIdTrtcMcuTranscodeTimeUsage `json:"SdkAppIdTranscodeTimeUsages,omitnil,omitempty" name:"SdkAppIdTranscodeTimeUsages"`

	// 查询记录数量
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 所查询的应用ID，可能值为:1-应用的应用ID，2-total，显示为total则表示查询的是所有应用的用量合计值。
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type OutputParams struct {
	// 直播流 ID，由用户自定义设置，该流 ID 不能与用户旁路的流 ID 相同，限制64字节。
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// 取值范围[0,1]， 填0：直播流为音视频(默认); 填1：直播流为纯音频
	PureAudioStream *uint64 `json:"PureAudioStream,omitnil,omitempty" name:"PureAudioStream"`

	// 自定义录制文件名称前缀。请先在实时音视频控制台开通录制功能，https://cloud.tencent.com/document/product/647/50768。
	// 【注意】该方式仅对旧版云端录制功能的应用生效，新版云端录制功能的应用请用接口CreateCloudRecording发起录制。新、旧云端录制类型判断方式请见：https://cloud.tencent.com/document/product/647/50768#record
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 取值范围[0,1]，填0无实际含义; 填1：指定录制文件格式为mp3。此参数不建议使用，建议在实时音视频控制台配置纯音频录制模板。
	RecordAudioOnly *uint64 `json:"RecordAudioOnly,omitnil,omitempty" name:"RecordAudioOnly"`
}

type PictureInfo struct {
	// 图片长度
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 图片宽度
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 显示位置x轴方向
	XPosition *uint64 `json:"XPosition,omitnil,omitempty" name:"XPosition"`

	// 显示位置y轴方向
	YPosition *uint64 `json:"YPosition,omitnil,omitempty" name:"YPosition"`

	// 应用id
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 图片id
	PictureId *uint64 `json:"PictureId,omitnil,omitempty" name:"PictureId"`
}

type PresetLayoutConfig struct {
	// 指定显示在该画面上的用户ID。如果不指定用户ID，会按照用户加入房间的顺序自动匹配PresetLayoutConfig中的画面设置。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 当该画面指定用户时，代表用户的流类型。0为摄像头，1为屏幕分享。小画面为web用户时此值填0。
	StreamType *uint64 `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// 该画面在输出时的宽度，单位为像素值，不填默认为0。
	ImageWidth *uint64 `json:"ImageWidth,omitnil,omitempty" name:"ImageWidth"`

	// 该画面在输出时的高度，单位为像素值，不填默认为0。
	ImageHeight *uint64 `json:"ImageHeight,omitnil,omitempty" name:"ImageHeight"`

	// 该画面在输出时的X偏移，单位为像素值，LocationX与ImageWidth之和不能超过混流输出的总宽度，不填默认为0。
	LocationX *uint64 `json:"LocationX,omitnil,omitempty" name:"LocationX"`

	// 该画面在输出时的Y偏移，单位为像素值，LocationY与ImageHeight之和不能超过混流输出的总高度，不填默认为0。
	LocationY *uint64 `json:"LocationY,omitnil,omitempty" name:"LocationY"`

	// 该画面在输出时的层级，不填默认为0。
	ZOrder *uint64 `json:"ZOrder,omitnil,omitempty" name:"ZOrder"`

	// 该画面在输出时的显示模式：0为裁剪，1为缩放，2为缩放并显示黑底。不填默认为0。
	RenderMode *uint64 `json:"RenderMode,omitnil,omitempty" name:"RenderMode"`

	// 该当前位置用户混入的流类型：0为混入音视频，1为只混入视频，2为只混入音频。默认为0，建议配合指定用户ID使用。
	MixInputType *uint64 `json:"MixInputType,omitnil,omitempty" name:"MixInputType"`

	// 占位图ID。启用占位图功能时，在当前位置的用户没有上行视频时显示占位图。占位图大小不能超过2M，在实时音视频控制台上传并生成，https://cloud.tencent.com/document/product/647/50769
	PlaceImageId *uint64 `json:"PlaceImageId,omitnil,omitempty" name:"PlaceImageId"`
}

type PublishCdnParams struct {
	// 腾讯云直播BizId。
	BizId *uint64 `json:"BizId,omitnil,omitempty" name:"BizId"`

	// 第三方CDN转推的目的地址，同时只支持转推一个第三方CDN地址。
	PublishCdnUrls []*string `json:"PublishCdnUrls,omitnil,omitempty" name:"PublishCdnUrls"`
}

type QualityData struct {
	// 数据内容
	Content []*TimeValue `json:"Content,omitnil,omitempty" name:"Content"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 对端Id,为空时表示上行数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerId *string `json:"PeerId,omitnil,omitempty" name:"PeerId"`

	// 数据类型
	DataType *string `json:"DataType,omitnil,omitempty" name:"DataType"`
}

type RecognizeConfig struct {
	// 语音转文字支持识别的语言，默认是"zh" 中文
	// 目前全量支持的语言如下，等号左面是语言英文名，右面是Language字段需要填写的值，该值遵循[ISO639](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes)：
	// 可通过购买「语音转文本时长包」解锁或领取包月套餐体验版解锁此功能。
	// 
	// 语音转文本支持语言类型如下：
	// - Chinese = "zh" # 中文
	// - Chinese_TW = "zh-TW" # 中国台湾
	// - English = "en" # 英语
	// - Chinese_YUE = "zh-yue" # 中国粤语
	// - Chinese_DIALECT = "zh-dialect" # 中国方言
	// - English = "en" # 英语
	// - Vietnamese = "vi" # 越南语
	// - Japanese = "ja" # 日语
	// - Korean = "ko" # 汉语
	// - Indonesia = "id" # 印度尼西亚语
	// - Thai = "th" # 泰语
	// - Portuguese = "pt" # 葡萄牙语
	// - Turkish = "tr" # 土耳其语
	// - Arabic = "ar" # 阿拉伯语
	// - Spanish = "es" # 西班牙语
	// - Hindi = "hi" # 印地语
	// - French = "fr" # 法语
	// - Malay = "ms" # 马来语
	// - Filipino = "fil" # 菲律宾语
	// - German = "de" # 德语
	// - Italian = "it" # 意大利语
	// - Russian = "ru" # 俄语
	// 
	// 注意：
	// 如果缺少满足您需求的语言，请联系我们技术人员。
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// 发起模糊识别额外可能替代语言类型,最多填写3种语言类型。
	// 注：Language指定为"zh-dialect" # 中国方言 时，不支持模糊识别，该字段无效
	AlternativeLanguage []*string `json:"AlternativeLanguage,omitnil,omitempty" name:"AlternativeLanguage"`

	// 目前已不支持
	//
	// Deprecated: Model is deprecated.
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 目前已不支持
	//
	// Deprecated: TranslationLanguage is deprecated.
	TranslationLanguage *string `json:"TranslationLanguage,omitnil,omitempty" name:"TranslationLanguage"`
}

type RecordParams struct {
	// 录制模式：
	// 1：单流录制，分别录制房间的订阅UserId的音频和视频，将录制文件上传至云存储；
	// 2：合流录制，将房间内订阅UserId的音视频混录成一个音视频文件，将录制文件上传至云存储；
	RecordMode *uint64 `json:"RecordMode,omitnil,omitempty" name:"RecordMode"`

	// 房间内持续没有用户（主播）上行推流的状态超过MaxIdleTime的时长，自动停止录制，单位：秒。默认值为 30 秒，该值需大于等于 5秒，且小于等于 86400秒(24小时)。
	MaxIdleTime *uint64 `json:"MaxIdleTime,omitnil,omitempty" name:"MaxIdleTime"`

	// 录制的媒体流类型：
	// 0：录制音频+视频流（默认）;
	// 1：仅录制音频流；
	// 2：仅录制视频流，
	StreamType *uint64 `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// 指定订阅流白名单或者黑名单。
	SubscribeStreamUserIds *SubscribeStreamUserIds `json:"SubscribeStreamUserIds,omitnil,omitempty" name:"SubscribeStreamUserIds"`

	// 输出文件的格式（存储至COS等第三方存储时有效）。0：(默认)输出文件为hls格式。1：输出文件格式为hls+mp4。2：输出文件格式为hls+aac 。3：输出文件格式为mp4。4：输出文件格式为aac。
	// 
	// 存储到云点播VOD时此参数无效，存储到VOD时请通过TencentVod（https://cloud.tencent.com/document/api/647/44055#TencentVod）内的MediaType设置。
	OutputFormat *uint64 `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// 单流录制模式下，用户的音视频是否合并，0：单流音视频不合并（默认）。1：单流音视频合并成一个ts。合流录制此参数无需设置，默认音视频合并。
	AvMerge *uint64 `json:"AvMerge,omitnil,omitempty" name:"AvMerge"`

	// 如果是aac或者mp4文件格式，超过长度限制后，系统会自动拆分视频文件。单位：分钟。默认为1440min（24h），取值范围为1-1440。【单文件限制最大为2G，满足文件大小 >2G 或录制时长度 > 24h任意一个条件，文件都会自动切分】
	// Hls 格式录制此参数不生效。
	MaxMediaFileDuration *uint64 `json:"MaxMediaFileDuration,omitnil,omitempty" name:"MaxMediaFileDuration"`

	// 指定录制主辅流，0：主流+辅流（默认）；1:主流；2:辅流。
	MediaId *uint64 `json:"MediaId,omitnil,omitempty" name:"MediaId"`

	// 上行视频停止时，录制的补帧类型，0：补最后一帧 1：补黑帧
	FillType *uint64 `json:"FillType,omitnil,omitempty" name:"FillType"`
}

type RecordUsage struct {
	// 本组数据对应的时间点，格式如:2020-09-07或2020-09-07 00:05:05。
	TimeKey *string `json:"TimeKey,omitnil,omitempty" name:"TimeKey"`

	// 视频时长-标清SD，单位：秒。
	Class1VideoTime *uint64 `json:"Class1VideoTime,omitnil,omitempty" name:"Class1VideoTime"`

	// 视频时长-高清HD，单位：秒。
	Class2VideoTime *uint64 `json:"Class2VideoTime,omitnil,omitempty" name:"Class2VideoTime"`

	// 视频时长-超清HD，单位：秒。
	Class3VideoTime *uint64 `json:"Class3VideoTime,omitnil,omitempty" name:"Class3VideoTime"`

	// 语音时长，单位：秒。
	AudioTime *uint64 `json:"AudioTime,omitnil,omitempty" name:"AudioTime"`
}

// Predefined struct for user
type RemoveUserByStrRoomIdRequestParams struct {
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 要移出的用户列表，最多10个。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

type RemoveUserByStrRoomIdRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 要移出的用户列表，最多10个。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

func (r *RemoveUserByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveUserByStrRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserByStrRoomIdResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveUserByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *RemoveUserByStrRoomIdResponseParams `json:"Response"`
}

func (r *RemoveUserByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserRequestParams struct {
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 要移出的用户列表，最多10个。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

type RemoveUserRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 要移出的用户列表，最多10个。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

func (r *RemoveUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveUserResponse struct {
	*tchttp.BaseResponse
	Response *RemoveUserResponseParams `json:"Response"`
}

func (r *RemoveUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoomState struct {
	// 通话ID（唯一标识一次通话）
	CommId *string `json:"CommId,omitnil,omitempty" name:"CommId"`

	// 房间号
	RoomString *string `json:"RoomString,omitnil,omitempty" name:"RoomString"`

	// 房间创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 房间销毁时间
	DestroyTime *uint64 `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`

	// 房间是否已经结束
	IsFinished *bool `json:"IsFinished,omitnil,omitempty" name:"IsFinished"`

	// 房间创建者Id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type RowValues struct {
	// 数据值
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowValue []*int64 `json:"RowValue,omitnil,omitempty" name:"RowValue"`
}

type STTConfig struct {
	// 语音识别支持的语言，默认是"zh" 中文
	// 目前全量支持的语言如下，等号左面是语言英文名，右面是Language字段需要填写的值，该值遵循[ISO639](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes)：
	// 1.     Chinese = "zh" # 中文
	// 2.     Chinese_TW = "zh-TW" # 中国台湾
	// 3.     Chinese_DIALECT = "zh-dialect" # 中国方言
	// 4.     English = "en" # 英语
	// 5.     Vietnamese = "vi" # 越南语
	// 6.     Japanese = "ja" # 日语
	// 7.     Korean = "ko" # 汉语
	// 8.     Indonesia = "id" # 印度尼西亚语
	// 9.     Thai = "th" # 泰语
	// 10.     Portuguese = "pt" # 葡萄牙语
	// 11.     Turkish = "tr" # 土耳其语
	// 12.     Arabic = "ar" # 阿拉伯语
	// 13.     Spanish = "es" # 西班牙语
	// 14.     Hindi = "hi" # 印地语
	// 15.     French = "fr" # 法语
	// 16.     Malay = "ms" # 马来语
	// 17.     Filipino = "fil" # 菲律宾语
	// 18.     German = "de" # 德语
	// 19.     Italian = "it" # 意大利语
	// 20.     Russian = "ru" # 俄语
	// 
	// 注意：
	// 如果缺少满足您需求的语言，请联系我们技术人员。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// 发起模糊识别额外可能替代语言类型,最多填写3种语言类型, 
	// 注：Language指定为"zh-dialect" # 中国方言 时，不支持模糊识别，该字段无效
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlternativeLanguage []*string `json:"AlternativeLanguage,omitnil,omitempty" name:"AlternativeLanguage"`
}

type ScaleInfomation struct {
	// 每天开始的时间
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 房间人数，用户重复进入同一个房间为1次
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserNumber *uint64 `json:"UserNumber,omitnil,omitempty" name:"UserNumber"`

	// 房间人次，用户每次进入房间为一次
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserCount *uint64 `json:"UserCount,omitnil,omitempty" name:"UserCount"`

	// sdkappid下一天内的房间数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomNumbers *uint64 `json:"RoomNumbers,omitnil,omitempty" name:"RoomNumbers"`
}

type SdkAppIdRecordUsage struct {
	// SdkAppId的值。
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 统计的时间点数据。
	Usages []*RecordUsage `json:"Usages,omitnil,omitempty" name:"Usages"`
}

type SdkAppIdTrtcMcuTranscodeTimeUsage struct {
	// 本组数据对应的时间点，格式如：2020-09-07或2020-09-07 00:05:05。
	TimeKey *string `json:"TimeKey,omitnil,omitempty" name:"TimeKey"`

	// 语音时长，单位：秒。
	AudioTime *uint64 `json:"AudioTime,omitnil,omitempty" name:"AudioTime"`

	// 视频时长-标清SD，单位：秒。
	VideoTimeSd *uint64 `json:"VideoTimeSd,omitnil,omitempty" name:"VideoTimeSd"`

	// 视频时长-高清HD，单位：秒。
	VideoTimeHd *uint64 `json:"VideoTimeHd,omitnil,omitempty" name:"VideoTimeHd"`

	// 视频时长-全高清FHD，单位：秒。
	VideoTimeFhd *uint64 `json:"VideoTimeFhd,omitnil,omitempty" name:"VideoTimeFhd"`

	// 带宽，单位：Mbps。
	Flux *float64 `json:"Flux,omitnil,omitempty" name:"Flux"`
}

type SeriesInfo struct {
	// 数据列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*string `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 数据值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*int64 `json:"Values,omitnil,omitempty" name:"Values"`
}

type SeriesInfos struct {
	// 数据列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*string `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 数据值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*RowValues `json:"Values,omitnil,omitempty" name:"Values"`
}

type ServerPushText struct {
	// 服务端推送播报文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 是否允许该文本打断机器人说话
	Interrupt *bool `json:"Interrupt,omitnil,omitempty" name:"Interrupt"`

	// 播报完文本后，是否自动关闭对话任务
	StopAfterPlay *bool `json:"StopAfterPlay,omitnil,omitempty" name:"StopAfterPlay"`
}

type SingleSubscribeParams struct {
	// 用户媒体流参数。
	UserMediaStream *UserMediaStream `json:"UserMediaStream,omitnil,omitempty" name:"UserMediaStream"`
}

type SmallVideoLayoutParams struct {
	// 代表小画面对应的用户ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 代表小画面对应的流类型，0为摄像头，1为屏幕分享。小画面为web用户时此值填0。
	StreamType *uint64 `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// 小画面在输出时的宽度，单位为像素值，不填默认为0。
	ImageWidth *uint64 `json:"ImageWidth,omitnil,omitempty" name:"ImageWidth"`

	// 小画面在输出时的高度，单位为像素值，不填默认为0。
	ImageHeight *uint64 `json:"ImageHeight,omitnil,omitempty" name:"ImageHeight"`

	// 小画面在输出时的X偏移，单位为像素值，LocationX与ImageWidth之和不能超过混流输出的总宽度，不填默认为0。
	LocationX *uint64 `json:"LocationX,omitnil,omitempty" name:"LocationX"`

	// 小画面在输出时的Y偏移，单位为像素值，LocationY与ImageHeight之和不能超过混流输出的总高度，不填默认为0。
	LocationY *uint64 `json:"LocationY,omitnil,omitempty" name:"LocationY"`
}

// Predefined struct for user
type StartAIConversationRequestParams struct {
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和开启转录任务的房间使用的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// TRTC的[RoomId](https://cloud.tencent.com/document/product/647/46351#roomid)，表示开启对话任务的房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 机器人参数
	AgentConfig *AgentConfig `json:"AgentConfig,omitnil,omitempty" name:"AgentConfig"`

	// 调用方传入的唯一Id，可用于客户侧防止重复发起任务以及可以通过该字段查询任务状态。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// TRTC房间号的类型，0代表数字房间号，1代表字符串房间号。不填默认是数字房间号。
	RoomIdType *uint64 `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`

	// 语音识别配置。
	STTConfig *STTConfig `json:"STTConfig,omitnil,omitempty" name:"STTConfig"`

	// LLM配置。需符合openai规范，为JSON字符串，示例如下：
	// <pre> { <br> &emsp;  "LLMType": “大模型类型"，  // String 必填，如："openai" <br> &emsp;  "Model": "您的模型名称", // String 必填，指定使用的模型<br>    "APIKey": "您的LLM API密钥", // String 必填 <br> &emsp;  "APIUrl": "https://api.xxx.com/chat/completions", // String 必填，LLM API访问的URL<br> &emsp;  "Streaming": true // Boolean 非必填，指定是否使用流式传输<br> &emsp;} </pre>
	LLMConfig *string `json:"LLMConfig,omitnil,omitempty" name:"LLMConfig"`

	// TTS配置，为JSON字符串，腾讯云TTS示例如下：
	//  <pre>{ <br> &emsp; "AppId": 您的应用ID, // Integer 必填<br> &emsp; "TTSType": "TTS类型", // String TTS类型, 固定为"tencent"<br> &emsp; "SecretId": "您的密钥ID", // String 必填<br> &emsp; "SecretKey":  "您的密钥Key", // String 必填<br> &emsp; "VoiceType": 101001, // Integer  必填，音色 ID，包括标准音色与精品音色，精品音色拟真度更高，价格不同于标准音色，请参见<a href="https://cloud.tencent.com/document/product/1073/34112">语音合成计费概述</a>。完整的音色 ID 列表请参见<a href="https://cloud.tencent.com/document/product/1073/92668#55924b56-1a73-4663-a7a1-a8dd82d6e823">语音合成音色列表</a>。<br> &emsp; "Speed": 1.25, // Integer 非必填，语速，范围：[-2，6]，分别对应不同语速： -2: 代表0.6倍 -1: 代表0.8倍 0: 代表1.0倍（默认） 1: 代表1.2倍 2: 代表1.5倍  6: 代表2.5倍  如果需要更细化的语速，可以保留小数点后 2 位，例如0.5/1.25/2.81等。 参数值与实际语速转换，可参考 <a href="https://sdk-1300466766.cos.ap-shanghai.myqcloud.com/sample/speed_sample.tar.gz">语速转换</a><br> &emsp; "Volume": 5, // Integer 非必填，音量大小，范围：[0，10]，分别对应11个等级的音量，默认值为0，代表正常音量。<br> &emsp; "PrimaryLanguage": "zh-CN" // String 非必填，主要语言<br> &emsp;}</pre>
	TTSConfig *string `json:"TTSConfig,omitnil,omitempty" name:"TTSConfig"`
}

type StartAIConversationRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和开启转录任务的房间使用的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// TRTC的[RoomId](https://cloud.tencent.com/document/product/647/46351#roomid)，表示开启对话任务的房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 机器人参数
	AgentConfig *AgentConfig `json:"AgentConfig,omitnil,omitempty" name:"AgentConfig"`

	// 调用方传入的唯一Id，可用于客户侧防止重复发起任务以及可以通过该字段查询任务状态。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// TRTC房间号的类型，0代表数字房间号，1代表字符串房间号。不填默认是数字房间号。
	RoomIdType *uint64 `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`

	// 语音识别配置。
	STTConfig *STTConfig `json:"STTConfig,omitnil,omitempty" name:"STTConfig"`

	// LLM配置。需符合openai规范，为JSON字符串，示例如下：
	// <pre> { <br> &emsp;  "LLMType": “大模型类型"，  // String 必填，如："openai" <br> &emsp;  "Model": "您的模型名称", // String 必填，指定使用的模型<br>    "APIKey": "您的LLM API密钥", // String 必填 <br> &emsp;  "APIUrl": "https://api.xxx.com/chat/completions", // String 必填，LLM API访问的URL<br> &emsp;  "Streaming": true // Boolean 非必填，指定是否使用流式传输<br> &emsp;} </pre>
	LLMConfig *string `json:"LLMConfig,omitnil,omitempty" name:"LLMConfig"`

	// TTS配置，为JSON字符串，腾讯云TTS示例如下：
	//  <pre>{ <br> &emsp; "AppId": 您的应用ID, // Integer 必填<br> &emsp; "TTSType": "TTS类型", // String TTS类型, 固定为"tencent"<br> &emsp; "SecretId": "您的密钥ID", // String 必填<br> &emsp; "SecretKey":  "您的密钥Key", // String 必填<br> &emsp; "VoiceType": 101001, // Integer  必填，音色 ID，包括标准音色与精品音色，精品音色拟真度更高，价格不同于标准音色，请参见<a href="https://cloud.tencent.com/document/product/1073/34112">语音合成计费概述</a>。完整的音色 ID 列表请参见<a href="https://cloud.tencent.com/document/product/1073/92668#55924b56-1a73-4663-a7a1-a8dd82d6e823">语音合成音色列表</a>。<br> &emsp; "Speed": 1.25, // Integer 非必填，语速，范围：[-2，6]，分别对应不同语速： -2: 代表0.6倍 -1: 代表0.8倍 0: 代表1.0倍（默认） 1: 代表1.2倍 2: 代表1.5倍  6: 代表2.5倍  如果需要更细化的语速，可以保留小数点后 2 位，例如0.5/1.25/2.81等。 参数值与实际语速转换，可参考 <a href="https://sdk-1300466766.cos.ap-shanghai.myqcloud.com/sample/speed_sample.tar.gz">语速转换</a><br> &emsp; "Volume": 5, // Integer 非必填，音量大小，范围：[0，10]，分别对应11个等级的音量，默认值为0，代表正常音量。<br> &emsp; "PrimaryLanguage": "zh-CN" // String 非必填，主要语言<br> &emsp;}</pre>
	TTSConfig *string `json:"TTSConfig,omitnil,omitempty" name:"TTSConfig"`
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
	delete(f, "SessionId")
	delete(f, "RoomIdType")
	delete(f, "STTConfig")
	delete(f, "LLMConfig")
	delete(f, "TTSConfig")
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
type StartAITranscriptionRequestParams struct {
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和开启转录任务的房间使用的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// TRTC的[RoomId](https://cloud.tencent.com/document/product/647/46351#roomid)，表示开启转录任务的房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 转录机器人的参数。
	TranscriptionParams *TranscriptionParams `json:"TranscriptionParams,omitnil,omitempty" name:"TranscriptionParams"`

	// 调用方传入的唯一Id，服务端用来去重。
	// 注意：
	// 如果传入该参数，服务端优先使用该参数来去重。
	// 如果不传该参数，服务端的去重策略如下：
	// - 如果TranscriptionMode字段是0，则一个房间只能开启一个任务
	// - 如果TranscriptionMode字段是1，则一个TargetUserId只能开启一个任务
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// TRTC房间号的类型，0代表数字房间号，1代表字符串房间号。不填默认是数字房间号。
	RoomIdType *uint64 `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`

	// 语音识别配置。
	RecognizeConfig *RecognizeConfig `json:"RecognizeConfig,omitnil,omitempty" name:"RecognizeConfig"`
}

type StartAITranscriptionRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和开启转录任务的房间使用的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// TRTC的[RoomId](https://cloud.tencent.com/document/product/647/46351#roomid)，表示开启转录任务的房间号。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 转录机器人的参数。
	TranscriptionParams *TranscriptionParams `json:"TranscriptionParams,omitnil,omitempty" name:"TranscriptionParams"`

	// 调用方传入的唯一Id，服务端用来去重。
	// 注意：
	// 如果传入该参数，服务端优先使用该参数来去重。
	// 如果不传该参数，服务端的去重策略如下：
	// - 如果TranscriptionMode字段是0，则一个房间只能开启一个任务
	// - 如果TranscriptionMode字段是1，则一个TargetUserId只能开启一个任务
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// TRTC房间号的类型，0代表数字房间号，1代表字符串房间号。不填默认是数字房间号。
	RoomIdType *uint64 `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`

	// 语音识别配置。
	RecognizeConfig *RecognizeConfig `json:"RecognizeConfig,omitnil,omitempty" name:"RecognizeConfig"`
}

func (r *StartAITranscriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAITranscriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "TranscriptionParams")
	delete(f, "SessionId")
	delete(f, "RoomIdType")
	delete(f, "RecognizeConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartAITranscriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartAITranscriptionResponseParams struct {
	// 用于唯一标识转录任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartAITranscriptionResponse struct {
	*tchttp.BaseResponse
	Response *StartAITranscriptionResponseParams `json:"Response"`
}

func (r *StartAITranscriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAITranscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMCUMixTranscodeByStrRoomIdRequestParams struct {
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 字符串房间号。
	StrRoomId *string `json:"StrRoomId,omitnil,omitempty" name:"StrRoomId"`

	// 混流输出控制参数。
	OutputParams *OutputParams `json:"OutputParams,omitnil,omitempty" name:"OutputParams"`

	// 混流输出编码参数。
	EncodeParams *EncodeParams `json:"EncodeParams,omitnil,omitempty" name:"EncodeParams"`

	// 混流输出布局参数。
	LayoutParams *LayoutParams `json:"LayoutParams,omitnil,omitempty" name:"LayoutParams"`

	// 第三方CDN转推参数。如需转推至腾讯云云直播，此参数无需填写，会默认转推
	PublishCdnParams *PublishCdnParams `json:"PublishCdnParams,omitnil,omitempty" name:"PublishCdnParams"`
}

type StartMCUMixTranscodeByStrRoomIdRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 字符串房间号。
	StrRoomId *string `json:"StrRoomId,omitnil,omitempty" name:"StrRoomId"`

	// 混流输出控制参数。
	OutputParams *OutputParams `json:"OutputParams,omitnil,omitempty" name:"OutputParams"`

	// 混流输出编码参数。
	EncodeParams *EncodeParams `json:"EncodeParams,omitnil,omitempty" name:"EncodeParams"`

	// 混流输出布局参数。
	LayoutParams *LayoutParams `json:"LayoutParams,omitnil,omitempty" name:"LayoutParams"`

	// 第三方CDN转推参数。如需转推至腾讯云云直播，此参数无需填写，会默认转推
	PublishCdnParams *PublishCdnParams `json:"PublishCdnParams,omitnil,omitempty" name:"PublishCdnParams"`
}

func (r *StartMCUMixTranscodeByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMCUMixTranscodeByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StrRoomId")
	delete(f, "OutputParams")
	delete(f, "EncodeParams")
	delete(f, "LayoutParams")
	delete(f, "PublishCdnParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMCUMixTranscodeByStrRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMCUMixTranscodeByStrRoomIdResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartMCUMixTranscodeByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *StartMCUMixTranscodeByStrRoomIdResponseParams `json:"Response"`
}

func (r *StartMCUMixTranscodeByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMCUMixTranscodeByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMCUMixTranscodeRequestParams struct {
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 混流输出控制参数。
	OutputParams *OutputParams `json:"OutputParams,omitnil,omitempty" name:"OutputParams"`

	// 混流输出编码参数。
	EncodeParams *EncodeParams `json:"EncodeParams,omitnil,omitempty" name:"EncodeParams"`

	// 混流输出布局参数。
	LayoutParams *LayoutParams `json:"LayoutParams,omitnil,omitempty" name:"LayoutParams"`

	// 第三方CDN转推参数。如需转推至腾讯云云直播，此参数无需填写，会默认转推
	PublishCdnParams *PublishCdnParams `json:"PublishCdnParams,omitnil,omitempty" name:"PublishCdnParams"`
}

type StartMCUMixTranscodeRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 混流输出控制参数。
	OutputParams *OutputParams `json:"OutputParams,omitnil,omitempty" name:"OutputParams"`

	// 混流输出编码参数。
	EncodeParams *EncodeParams `json:"EncodeParams,omitnil,omitempty" name:"EncodeParams"`

	// 混流输出布局参数。
	LayoutParams *LayoutParams `json:"LayoutParams,omitnil,omitempty" name:"LayoutParams"`

	// 第三方CDN转推参数。如需转推至腾讯云云直播，此参数无需填写，会默认转推
	PublishCdnParams *PublishCdnParams `json:"PublishCdnParams,omitnil,omitempty" name:"PublishCdnParams"`
}

func (r *StartMCUMixTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMCUMixTranscodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "OutputParams")
	delete(f, "EncodeParams")
	delete(f, "LayoutParams")
	delete(f, "PublishCdnParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMCUMixTranscodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMCUMixTranscodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartMCUMixTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *StartMCUMixTranscodeResponseParams `json:"Response"`
}

func (r *StartMCUMixTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMCUMixTranscodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishCdnStreamRequestParams struct {
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和转推的房间所对应的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 主房间信息RoomId，转推的TRTC房间所对应的RoomId。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 主房间信息RoomType，必须和转推的房间所对应的RoomId类型相同，0为整型房间号，1为字符串房间号。
	RoomIdType *uint64 `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`

	// 转推服务加入TRTC房间的机器人参数。
	AgentParams *AgentParams `json:"AgentParams,omitnil,omitempty" name:"AgentParams"`

	// 是否转码，0表示无需转码，1表示需要转码。是否收取转码费是由WithTranscoding参数决定的，WithTranscoding为0，表示旁路转推，不会收取转码费用，WithTranscoding为1，表示混流转推，会收取转码费用。
	WithTranscoding *uint64 `json:"WithTranscoding,omitnil,omitempty" name:"WithTranscoding"`

	// 转推流的音频编码参数。由于音频是必转码的（不会收取转码费用），所以启动任务的时候，必须填写。
	AudioParams *McuAudioParams `json:"AudioParams,omitnil,omitempty" name:"AudioParams"`

	// 转推流的视频编码参数，不填表示纯音频转推。
	VideoParams *McuVideoParams `json:"VideoParams,omitnil,omitempty" name:"VideoParams"`

	// 需要单流旁路转推的用户上行参数，单流旁路转推时，WithTranscoding需要设置为0。
	SingleSubscribeParams *SingleSubscribeParams `json:"SingleSubscribeParams,omitnil,omitempty" name:"SingleSubscribeParams"`

	// 转推的CDN参数。和回推房间参数必须要有一个。
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitnil,omitempty" name:"PublishCdnParams"`

	// 混流SEI参数
	SeiParams *McuSeiParams `json:"SeiParams,omitnil,omitempty" name:"SeiParams"`

	// 回推房间信息，和转推CDN参数必须要有一个。注：回推房间需使用10.4及以上SDK版本，如您有需求，请联系腾讯云技术支持。
	FeedBackRoomParams []*McuFeedBackRoomParams `json:"FeedBackRoomParams,omitnil,omitempty" name:"FeedBackRoomParams"`
}

type StartPublishCdnStreamRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和转推的房间所对应的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 主房间信息RoomId，转推的TRTC房间所对应的RoomId。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 主房间信息RoomType，必须和转推的房间所对应的RoomId类型相同，0为整型房间号，1为字符串房间号。
	RoomIdType *uint64 `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`

	// 转推服务加入TRTC房间的机器人参数。
	AgentParams *AgentParams `json:"AgentParams,omitnil,omitempty" name:"AgentParams"`

	// 是否转码，0表示无需转码，1表示需要转码。是否收取转码费是由WithTranscoding参数决定的，WithTranscoding为0，表示旁路转推，不会收取转码费用，WithTranscoding为1，表示混流转推，会收取转码费用。
	WithTranscoding *uint64 `json:"WithTranscoding,omitnil,omitempty" name:"WithTranscoding"`

	// 转推流的音频编码参数。由于音频是必转码的（不会收取转码费用），所以启动任务的时候，必须填写。
	AudioParams *McuAudioParams `json:"AudioParams,omitnil,omitempty" name:"AudioParams"`

	// 转推流的视频编码参数，不填表示纯音频转推。
	VideoParams *McuVideoParams `json:"VideoParams,omitnil,omitempty" name:"VideoParams"`

	// 需要单流旁路转推的用户上行参数，单流旁路转推时，WithTranscoding需要设置为0。
	SingleSubscribeParams *SingleSubscribeParams `json:"SingleSubscribeParams,omitnil,omitempty" name:"SingleSubscribeParams"`

	// 转推的CDN参数。和回推房间参数必须要有一个。
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitnil,omitempty" name:"PublishCdnParams"`

	// 混流SEI参数
	SeiParams *McuSeiParams `json:"SeiParams,omitnil,omitempty" name:"SeiParams"`

	// 回推房间信息，和转推CDN参数必须要有一个。注：回推房间需使用10.4及以上SDK版本，如您有需求，请联系腾讯云技术支持。
	FeedBackRoomParams []*McuFeedBackRoomParams `json:"FeedBackRoomParams,omitnil,omitempty" name:"FeedBackRoomParams"`
}

func (r *StartPublishCdnStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishCdnStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "RoomIdType")
	delete(f, "AgentParams")
	delete(f, "WithTranscoding")
	delete(f, "AudioParams")
	delete(f, "VideoParams")
	delete(f, "SingleSubscribeParams")
	delete(f, "PublishCdnParams")
	delete(f, "SeiParams")
	delete(f, "FeedBackRoomParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartPublishCdnStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishCdnStreamResponseParams struct {
	// 用于唯一标识转推任务，由腾讯云服务端生成，后续更新和停止请求都需要携带TaskiD参数。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartPublishCdnStreamResponse struct {
	*tchttp.BaseResponse
	Response *StartPublishCdnStreamResponseParams `json:"Response"`
}

func (r *StartPublishCdnStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishCdnStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartStreamIngestRequestParams struct {
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和TRTC的房间所对应的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// TRTC的[RoomId](https://cloud.tencent.com/document/product/647/46351#roomid)，录制的TRTC房间所对应的RoomId。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// TRTC房间号的类型。
	// 【*注意】必须和录制的房间所对应的RoomId类型相同:
	// 0: 字符串类型的RoomId
	// 1: 32位整型的RoomId（默认）
	RoomIdType *uint64 `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`

	// 输入在线媒体流机器人的UserId，用于进房发起拉流转推任务。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 输入在线媒体流机器人UserId对应的校验签名，即UserId和UserSig相当于机器人进房的登录密码，具体计算方法请参考TRTC计算[UserSig](https://cloud.tencent.com/document/product/647/45910#UserSig)的方案。
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`

	// 源流URL【必填】。如果是视频流，分辨率请保持不变。
	StreamUrl *string `json:"StreamUrl,omitnil,omitempty" name:"StreamUrl"`

	// TRTC房间权限加密串，只有在TRTC控制台启用了高级权限控制的时候需要携带，在TRTC控制台如果开启高级权限控制后，TRTC 的后台服务系统会校验一个叫做 [PrivateMapKey] 的“权限票据”，权限票据中包含了一个加密后的 RoomId 和一个加密后的“权限位列表”。由于 PrivateMapKey 中包含 RoomId，所以只提供了 UserSig 没有提供 PrivateMapKey 时，并不能进入指定的房间。
	PrivateMapKey *string `json:"PrivateMapKey,omitnil,omitempty" name:"PrivateMapKey"`

	// 【本字段已废弃】视频编码参数。可选，如果不填，保持原始流的参数。
	//
	// Deprecated: VideoEncodeParams is deprecated.
	VideoEncodeParams *VideoEncodeParams `json:"VideoEncodeParams,omitnil,omitempty" name:"VideoEncodeParams"`

	// 【本字段已废弃】音频编码参数。可选，如果不填，保持原始流的参数。
	//
	// Deprecated: AudioEncodeParams is deprecated.
	AudioEncodeParams *AudioEncodeParams `json:"AudioEncodeParams,omitnil,omitempty" name:"AudioEncodeParams"`

	// 【本字段已废弃，请使用 StreamUrl 字段】源流URL，支持一个地址。
	//
	// Deprecated: SourceUrl is deprecated.
	SourceUrl []*string `json:"SourceUrl,omitnil,omitempty" name:"SourceUrl"`

	// 指定视频从某个秒时间戳播放
	SeekSecond *int64 `json:"SeekSecond,omitnil,omitempty" name:"SeekSecond"`

	// 开启自动旁路推流，请确认控制台已经开启该功能。
	AutoPush *bool `json:"AutoPush,omitnil,omitempty" name:"AutoPush"`

	// 循环播放次数, 取值范围[-1, 1000],  默认1次。
	//  - 0 无效值
	//  - -1 循环播放, 需要主动调用停止接口或设置MaxDuration
	RepeatNum *int64 `json:"RepeatNum,omitnil,omitempty" name:"RepeatNum"`

	// 循环播放最大时长,仅支持RepeatNum设置-1时生效，取值范围[1, 10080]，单位分钟。
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`
}

type StartStreamIngestRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和TRTC的房间所对应的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// TRTC的[RoomId](https://cloud.tencent.com/document/product/647/46351#roomid)，录制的TRTC房间所对应的RoomId。
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// TRTC房间号的类型。
	// 【*注意】必须和录制的房间所对应的RoomId类型相同:
	// 0: 字符串类型的RoomId
	// 1: 32位整型的RoomId（默认）
	RoomIdType *uint64 `json:"RoomIdType,omitnil,omitempty" name:"RoomIdType"`

	// 输入在线媒体流机器人的UserId，用于进房发起拉流转推任务。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 输入在线媒体流机器人UserId对应的校验签名，即UserId和UserSig相当于机器人进房的登录密码，具体计算方法请参考TRTC计算[UserSig](https://cloud.tencent.com/document/product/647/45910#UserSig)的方案。
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`

	// 源流URL【必填】。如果是视频流，分辨率请保持不变。
	StreamUrl *string `json:"StreamUrl,omitnil,omitempty" name:"StreamUrl"`

	// TRTC房间权限加密串，只有在TRTC控制台启用了高级权限控制的时候需要携带，在TRTC控制台如果开启高级权限控制后，TRTC 的后台服务系统会校验一个叫做 [PrivateMapKey] 的“权限票据”，权限票据中包含了一个加密后的 RoomId 和一个加密后的“权限位列表”。由于 PrivateMapKey 中包含 RoomId，所以只提供了 UserSig 没有提供 PrivateMapKey 时，并不能进入指定的房间。
	PrivateMapKey *string `json:"PrivateMapKey,omitnil,omitempty" name:"PrivateMapKey"`

	// 【本字段已废弃】视频编码参数。可选，如果不填，保持原始流的参数。
	VideoEncodeParams *VideoEncodeParams `json:"VideoEncodeParams,omitnil,omitempty" name:"VideoEncodeParams"`

	// 【本字段已废弃】音频编码参数。可选，如果不填，保持原始流的参数。
	AudioEncodeParams *AudioEncodeParams `json:"AudioEncodeParams,omitnil,omitempty" name:"AudioEncodeParams"`

	// 【本字段已废弃，请使用 StreamUrl 字段】源流URL，支持一个地址。
	SourceUrl []*string `json:"SourceUrl,omitnil,omitempty" name:"SourceUrl"`

	// 指定视频从某个秒时间戳播放
	SeekSecond *int64 `json:"SeekSecond,omitnil,omitempty" name:"SeekSecond"`

	// 开启自动旁路推流，请确认控制台已经开启该功能。
	AutoPush *bool `json:"AutoPush,omitnil,omitempty" name:"AutoPush"`

	// 循环播放次数, 取值范围[-1, 1000],  默认1次。
	//  - 0 无效值
	//  - -1 循环播放, 需要主动调用停止接口或设置MaxDuration
	RepeatNum *int64 `json:"RepeatNum,omitnil,omitempty" name:"RepeatNum"`

	// 循环播放最大时长,仅支持RepeatNum设置-1时生效，取值范围[1, 10080]，单位分钟。
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`
}

func (r *StartStreamIngestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStreamIngestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "RoomIdType")
	delete(f, "UserId")
	delete(f, "UserSig")
	delete(f, "StreamUrl")
	delete(f, "PrivateMapKey")
	delete(f, "VideoEncodeParams")
	delete(f, "AudioEncodeParams")
	delete(f, "SourceUrl")
	delete(f, "SeekSecond")
	delete(f, "AutoPush")
	delete(f, "RepeatNum")
	delete(f, "MaxDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartStreamIngestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartStreamIngestResponseParams struct {
	// 输入在线媒体流的任务 ID。任务 ID 是对一次输入在线媒体流生命周期过程的唯一标识，结束任务时会失去意义。任务 ID 需要业务保存下来，作为下次针对这个任务操作的参数。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartStreamIngestResponse struct {
	*tchttp.BaseResponse
	Response *StartStreamIngestResponseParams `json:"Response"`
}

func (r *StartStreamIngestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStreamIngestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartWebRecordRequestParams struct {
	// 需要录制的网页URL
	RecordUrl *string `json:"RecordUrl,omitnil,omitempty" name:"RecordUrl"`

	// 录制最大时长限制， 单位 s, 合法取值范围[1800, 36000], 默认 36000s(10 小时)
	MaxDurationLimit *uint64 `json:"MaxDurationLimit,omitnil,omitempty" name:"MaxDurationLimit"`

	// 【必填】云存储相关的参数，目前支持腾讯云对象存储以及腾讯云云点播VOD，不支持第三方云存储；输出文件的存储格式仅支持hls或mp4
	StorageParams *StorageParams `json:"StorageParams,omitnil,omitempty" name:"StorageParams"`

	// 页面录制视频参数
	WebRecordVideoParams *WebRecordVideoParams `json:"WebRecordVideoParams,omitnil,omitempty" name:"WebRecordVideoParams"`

	// 【必填】TRTC的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 当对重复任务敏感时，请关注此值： 为了避免任务在短时间内重复发起，导致任务重复
	// 传入录制RecordId来标识此次任务， 小于32字节，若携带RecordId发起两次以上的开始录制请求，任务只会启动一个，第二个报错FailedOperation.TaskExist。注意StartWebRecord调用失败时而非FailedOperation.TaskExist错误，请更换RecordId重新发起。
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 若您想要推流到CDN，可以使用PublishCdnParams.N参数设置，支持最多同时推流到10个CDN地址。若转推地址是腾讯云CDN时，请将IsTencentCdn明确设置为1
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitnil,omitempty" name:"PublishCdnParams"`
}

type StartWebRecordRequest struct {
	*tchttp.BaseRequest
	
	// 需要录制的网页URL
	RecordUrl *string `json:"RecordUrl,omitnil,omitempty" name:"RecordUrl"`

	// 录制最大时长限制， 单位 s, 合法取值范围[1800, 36000], 默认 36000s(10 小时)
	MaxDurationLimit *uint64 `json:"MaxDurationLimit,omitnil,omitempty" name:"MaxDurationLimit"`

	// 【必填】云存储相关的参数，目前支持腾讯云对象存储以及腾讯云云点播VOD，不支持第三方云存储；输出文件的存储格式仅支持hls或mp4
	StorageParams *StorageParams `json:"StorageParams,omitnil,omitempty" name:"StorageParams"`

	// 页面录制视频参数
	WebRecordVideoParams *WebRecordVideoParams `json:"WebRecordVideoParams,omitnil,omitempty" name:"WebRecordVideoParams"`

	// 【必填】TRTC的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 当对重复任务敏感时，请关注此值： 为了避免任务在短时间内重复发起，导致任务重复
	// 传入录制RecordId来标识此次任务， 小于32字节，若携带RecordId发起两次以上的开始录制请求，任务只会启动一个，第二个报错FailedOperation.TaskExist。注意StartWebRecord调用失败时而非FailedOperation.TaskExist错误，请更换RecordId重新发起。
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 若您想要推流到CDN，可以使用PublishCdnParams.N参数设置，支持最多同时推流到10个CDN地址。若转推地址是腾讯云CDN时，请将IsTencentCdn明确设置为1
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitnil,omitempty" name:"PublishCdnParams"`
}

func (r *StartWebRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartWebRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RecordUrl")
	delete(f, "MaxDurationLimit")
	delete(f, "StorageParams")
	delete(f, "WebRecordVideoParams")
	delete(f, "SdkAppId")
	delete(f, "RecordId")
	delete(f, "PublishCdnParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartWebRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartWebRecordResponseParams struct {
	// 录制任务的唯一Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartWebRecordResponse struct {
	*tchttp.BaseResponse
	Response *StartWebRecordResponseParams `json:"Response"`
}

func (r *StartWebRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartWebRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
type StopAITranscriptionRequestParams struct {
	// 唯一标识转录任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopAITranscriptionRequest struct {
	*tchttp.BaseRequest
	
	// 唯一标识转录任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopAITranscriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAITranscriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopAITranscriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopAITranscriptionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopAITranscriptionResponse struct {
	*tchttp.BaseResponse
	Response *StopAITranscriptionResponseParams `json:"Response"`
}

func (r *StopAITranscriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAITranscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMCUMixTranscodeByStrRoomIdRequestParams struct {
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 字符串房间号。
	StrRoomId *string `json:"StrRoomId,omitnil,omitempty" name:"StrRoomId"`
}

type StopMCUMixTranscodeByStrRoomIdRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 字符串房间号。
	StrRoomId *string `json:"StrRoomId,omitnil,omitempty" name:"StrRoomId"`
}

func (r *StopMCUMixTranscodeByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMCUMixTranscodeByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StrRoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMCUMixTranscodeByStrRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMCUMixTranscodeByStrRoomIdResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopMCUMixTranscodeByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *StopMCUMixTranscodeByStrRoomIdResponseParams `json:"Response"`
}

func (r *StopMCUMixTranscodeByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMCUMixTranscodeByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMCUMixTranscodeRequestParams struct {
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type StopMCUMixTranscodeRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *StopMCUMixTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMCUMixTranscodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMCUMixTranscodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMCUMixTranscodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopMCUMixTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *StopMCUMixTranscodeResponseParams `json:"Response"`
}

func (r *StopMCUMixTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMCUMixTranscodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishCdnStreamRequestParams struct {
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和转推的房间所对应的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 唯一标识转推任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopPublishCdnStreamRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和转推的房间所对应的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 唯一标识转推任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopPublishCdnStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishCdnStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopPublishCdnStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishCdnStreamResponseParams struct {
	// 转推任务唯一的String Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopPublishCdnStreamResponse struct {
	*tchttp.BaseResponse
	Response *StopPublishCdnStreamResponseParams `json:"Response"`
}

func (r *StopPublishCdnStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishCdnStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopStreamIngestRequestParams struct {
	// TRTC的SDKAppId，和任务的房间所对应的SDKAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务的唯一Id，在启动任务成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopStreamIngestRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId，和任务的房间所对应的SDKAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务的唯一Id，在启动任务成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopStreamIngestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopStreamIngestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopStreamIngestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopStreamIngestResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopStreamIngestResponse struct {
	*tchttp.BaseResponse
	Response *StopStreamIngestResponseParams `json:"Response"`
}

func (r *StopStreamIngestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopStreamIngestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWebRecordRequestParams struct {
	// 需要停止的任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopWebRecordRequest struct {
	*tchttp.BaseRequest
	
	// 需要停止的任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopWebRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWebRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopWebRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWebRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopWebRecordResponse struct {
	*tchttp.BaseResponse
	Response *StopWebRecordResponseParams `json:"Response"`
}

func (r *StopWebRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWebRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageFile struct {
	// 录制文件对应的UserId，如果是混流的话的这里返回的是空串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 录制索引文件名。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 录制文件流信息。
	// video：视频录制文件
	// audio：音频录制文件
	// audio_video：音视频录制文件
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrackType *string `json:"TrackType,omitnil,omitempty" name:"TrackType"`

	// 录制文件开始Unix时间戳。
	BeginTimeStamp *uint64 `json:"BeginTimeStamp,omitnil,omitempty" name:"BeginTimeStamp"`
}

type StorageParams struct {
	// 腾讯云对象存储COS以及第三方云存储的账号信息
	CloudStorage *CloudStorage `json:"CloudStorage,omitnil,omitempty" name:"CloudStorage"`

	// 腾讯云云点播Vod的存储信息
	CloudVod *CloudVod `json:"CloudVod,omitnil,omitempty" name:"CloudVod"`
}

type SubscribeStreamUserIds struct {
	// 订阅音频流白名单，指定订阅哪几个UserId的音频流，例如["1", "2", "3"], 代表订阅UserId 1，2，3的音频流；["1.*$"], 代表订阅UserId前缀为1的音频流。默认不填订阅房间内所有的音频流，订阅列表用户数不超过32。
	SubscribeAudioUserIds []*string `json:"SubscribeAudioUserIds,omitnil,omitempty" name:"SubscribeAudioUserIds"`

	// 订阅音频流黑名单，指定不订阅哪几个UserId的音频流，例如["1", "2", "3"], 代表不订阅UserId 1，2，3的音频流；["1.*$"], 代表不订阅UserId前缀为1的音频流。默认不填订阅房间内所有音频流，订阅列表用户数不超过32。
	UnSubscribeAudioUserIds []*string `json:"UnSubscribeAudioUserIds,omitnil,omitempty" name:"UnSubscribeAudioUserIds"`

	// 订阅视频流白名单，指定订阅哪几个UserId的视频流，例如["1", "2", "3"], 代表订阅UserId  1，2，3的视频流；["1.*$"], 代表订阅UserId前缀为1的视频流。默认不填订阅房间内所有视频流，订阅列表用户数不超过32。
	SubscribeVideoUserIds []*string `json:"SubscribeVideoUserIds,omitnil,omitempty" name:"SubscribeVideoUserIds"`

	// 订阅视频流黑名单，指定不订阅哪几个UserId的视频流，例如["1", "2", "3"], 代表不订阅UserId  1，2，3的视频流；["1.*$"], 代表不订阅UserId前缀为1的视频流。默认不填订阅房间内所有视频流，订阅列表用户数不超过32。
	UnSubscribeVideoUserIds []*string `json:"UnSubscribeVideoUserIds,omitnil,omitempty" name:"UnSubscribeVideoUserIds"`
}

// Predefined struct for user
type SummarizeTranscriptionRequestParams struct {

}

type SummarizeTranscriptionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *SummarizeTranscriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SummarizeTranscriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SummarizeTranscriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SummarizeTranscriptionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SummarizeTranscriptionResponse struct {
	*tchttp.BaseResponse
	Response *SummarizeTranscriptionResponseParams `json:"Response"`
}

func (r *SummarizeTranscriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SummarizeTranscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TRTCDataResp struct {
	// StatementID值，监控仪表盘下固定为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatementID *int64 `json:"StatementID,omitnil,omitempty" name:"StatementID"`

	// 查询结果数据，以Columns-Values形式返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Series []*SeriesInfo `json:"Series,omitnil,omitempty" name:"Series"`

	// Total值，监控仪表盘功能下固定为1。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type TRTCDataResult struct {
	// StatementID值，监控仪表盘下固定为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatementID *int64 `json:"StatementID,omitnil,omitempty" name:"StatementID"`

	// 查询结果数据，以Columns-Values形式返回。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Series []*SeriesInfos `json:"Series,omitnil,omitempty" name:"Series"`

	// Total值，监控仪表盘功能下固定为1。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type TencentVod struct {
	// 媒体后续任务处理操作，即完成媒体上传后，可自动发起任务流操作。参数值为任务流模板名，云点播支持 创建任务流模板 并为模板命名。
	Procedure *string `json:"Procedure,omitnil,omitempty" name:"Procedure"`

	// 媒体文件过期时间，为当前时间的绝对过期时间；保存一天，就填"86400"，永久保存就填"0"，默认永久保存。
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 指定上传园区，仅适用于对上传地域有特殊需求的用户。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 分类ID，用于对媒体进行分类管理，可通过 创建分类 接口，创建分类，获得分类 ID。
	// 默认值：0，表示其他分类。
	ClassId *uint64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 点播 子应用 ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 任务流上下文，任务完成回调时透传。
	SessionContext *string `json:"SessionContext,omitnil,omitempty" name:"SessionContext"`

	// 上传上下文，上传完成回调时透传。
	SourceContext *string `json:"SourceContext,omitnil,omitempty" name:"SourceContext"`

	// 上传到vod平台的录制文件格式类型，0：mp4(默认), 1: hls, 2:aac(StreamType=1纯音频录制时有效),
	// 3: hls+mp4, 4: hls+aac(StreamType=1纯音频录制时有效)。
	MediaType *uint64 `json:"MediaType,omitnil,omitempty" name:"MediaType"`

	// 仅支持API录制上传vod，该参数表示用户可以自定义录制文件名前缀，【限制长度为64字节，只允许包含大小写英文字母（a-zA-Z）、数字（0-9）及下划线和连词符】。前缀与自动生成的录制文件名之间用`__UserDefine_u_` 分开。
	UserDefineRecordId *string `json:"UserDefineRecordId,omitnil,omitempty" name:"UserDefineRecordId"`
}

type TimeValue struct {
	// 时间，unix时间戳（1590065877s)
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 当前时间返回参数取值，如（bigvCapFps在1590065877取值为0，则Value：0 ）
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type TranscriptionParams struct {
	// 转录机器人的UserId，用于进房发起转录任务。【注意】这个UserId不能与当前房间内的主播观众[UserId](https://cloud.tencent.com/document/product/647/46351#userid)重复。如果一个房间发起多个转录任务时，机器人的userid也不能相互重复，否则会中断前一个任务。需要保证转录机器人UserId在房间内唯一。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 转录机器人UserId对应的校验签名，即UserId和UserSig相当于转录机器人进房的登录密码，具体计算方法请参考TRTC计算[UserSig](https://cloud.tencent.com/document/product/647/45910#UserSig)的方案。
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`

	// IM[管理员账户](
	// https://cloud.tencent.com/document/product/269/31999#app-.E7.AE.A1.E7.90.86.E5.91.98)，如果填写，后台下发消息会使用IM通道，而不是TRTC自定义消息。
	//
	// Deprecated: IMAdminUserId is deprecated.
	IMAdminUserId *string `json:"IMAdminUserId,omitnil,omitempty" name:"IMAdminUserId"`

	// IM管理员账户生成的签名，用于向特定群组发送消息。如果填写，后台下发消息会使用IM通道，而不是TRTC自定义消息。必须和IM管理员的UserId一起填写。
	//
	// Deprecated: IMAdminUserSig is deprecated.
	IMAdminUserSig *string `json:"IMAdminUserSig,omitnil,omitempty" name:"IMAdminUserSig"`

	// 房间内推流用户全部退出后超过MaxIdleTime秒，后台自动关闭转录任务，默认值是60s。
	MaxIdleTime *uint64 `json:"MaxIdleTime,omitnil,omitempty" name:"MaxIdleTime"`

	// 1表示机器人只订阅单个人的流，0表示机器人订阅整个房间的流，如果不填默认订阅整个房间的流。
	TranscriptionMode *uint64 `json:"TranscriptionMode,omitnil,omitempty" name:"TranscriptionMode"`

	// TranscriptionMode为1时必填，机器人只会拉该userid的流，忽略房间里其他用户。
	TargetUserId *string `json:"TargetUserId,omitnil,omitempty" name:"TargetUserId"`
}

type TrtcUsage struct {
	// 时间点，格式为YYYY-MM-DD HH:mm:ss。多天查询时，HH:mm:ss为00:00:00。
	TimeKey *string `json:"TimeKey,omitnil,omitempty" name:"TimeKey"`

	// 时间点时间戳
	TimeStampKey *uint64 `json:"TimeStampKey,omitnil,omitempty" name:"TimeStampKey"`

	// 用量数组。每个数值含义与UsageKey对应。单位:分钟。
	UsageValue []*float64 `json:"UsageValue,omitnil,omitempty" name:"UsageValue"`
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
type UpdatePublishCdnStreamRequestParams struct {
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和转推的房间所对应的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 唯一标识转推任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 客户保证同一个任务，每次更新请求中的SequenceNumber递增，防止请求乱序。
	SequenceNumber *uint64 `json:"SequenceNumber,omitnil,omitempty" name:"SequenceNumber"`

	// 是否转码，0表示无需转码，1表示需要转码。
	WithTranscoding *uint64 `json:"WithTranscoding,omitnil,omitempty" name:"WithTranscoding"`

	// 更新相关参数，只支持更新参与混音的主播列表参数，不支持更新Codec、采样率、码率和声道数。不填表示不更新此参数。
	AudioParams *McuAudioParams `json:"AudioParams,omitnil,omitempty" name:"AudioParams"`

	// 更新视频相关参数，转码时支持更新除编码类型之外的编码参数，视频布局参数，背景图片和背景颜色参数，水印参数。不填表示不更新此参数。
	VideoParams *McuVideoParams `json:"VideoParams,omitnil,omitempty" name:"VideoParams"`

	// 更新单流转推的用户上行参数，仅在非转码时有效。不填表示不更新此参数。
	SingleSubscribeParams *SingleSubscribeParams `json:"SingleSubscribeParams,omitnil,omitempty" name:"SingleSubscribeParams"`

	// 更新转推的CDN参数。不填表示不更新此参数。
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitnil,omitempty" name:"PublishCdnParams"`

	// 混流SEI参数
	SeiParams *McuSeiParams `json:"SeiParams,omitnil,omitempty" name:"SeiParams"`

	// 回推房间信息
	FeedBackRoomParams []*McuFeedBackRoomParams `json:"FeedBackRoomParams,omitnil,omitempty" name:"FeedBackRoomParams"`
}

type UpdatePublishCdnStreamRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的[SdkAppId](https://cloud.tencent.com/document/product/647/46351#sdkappid)，和转推的房间所对应的SdkAppId相同。
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 唯一标识转推任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 客户保证同一个任务，每次更新请求中的SequenceNumber递增，防止请求乱序。
	SequenceNumber *uint64 `json:"SequenceNumber,omitnil,omitempty" name:"SequenceNumber"`

	// 是否转码，0表示无需转码，1表示需要转码。
	WithTranscoding *uint64 `json:"WithTranscoding,omitnil,omitempty" name:"WithTranscoding"`

	// 更新相关参数，只支持更新参与混音的主播列表参数，不支持更新Codec、采样率、码率和声道数。不填表示不更新此参数。
	AudioParams *McuAudioParams `json:"AudioParams,omitnil,omitempty" name:"AudioParams"`

	// 更新视频相关参数，转码时支持更新除编码类型之外的编码参数，视频布局参数，背景图片和背景颜色参数，水印参数。不填表示不更新此参数。
	VideoParams *McuVideoParams `json:"VideoParams,omitnil,omitempty" name:"VideoParams"`

	// 更新单流转推的用户上行参数，仅在非转码时有效。不填表示不更新此参数。
	SingleSubscribeParams *SingleSubscribeParams `json:"SingleSubscribeParams,omitnil,omitempty" name:"SingleSubscribeParams"`

	// 更新转推的CDN参数。不填表示不更新此参数。
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitnil,omitempty" name:"PublishCdnParams"`

	// 混流SEI参数
	SeiParams *McuSeiParams `json:"SeiParams,omitnil,omitempty" name:"SeiParams"`

	// 回推房间信息
	FeedBackRoomParams []*McuFeedBackRoomParams `json:"FeedBackRoomParams,omitnil,omitempty" name:"FeedBackRoomParams"`
}

func (r *UpdatePublishCdnStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePublishCdnStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	delete(f, "SequenceNumber")
	delete(f, "WithTranscoding")
	delete(f, "AudioParams")
	delete(f, "VideoParams")
	delete(f, "SingleSubscribeParams")
	delete(f, "PublishCdnParams")
	delete(f, "SeiParams")
	delete(f, "FeedBackRoomParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePublishCdnStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePublishCdnStreamResponseParams struct {
	// 转推任务唯一的String Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdatePublishCdnStreamResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePublishCdnStreamResponseParams `json:"Response"`
}

func (r *UpdatePublishCdnStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePublishCdnStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateStreamIngestRequestParams struct {
	// TRTC的SDKAppId，和任务的房间所对应的SDKAppId相同
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务的唯一Id，在启动任务成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 源流URL【必填】。
	StreamUrl *string `json:"StreamUrl,omitnil,omitempty" name:"StreamUrl"`
}

type UpdateStreamIngestRequest struct {
	*tchttp.BaseRequest
	
	// TRTC的SDKAppId，和任务的房间所对应的SDKAppId相同
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务的唯一Id，在启动任务成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 源流URL【必填】。
	StreamUrl *string `json:"StreamUrl,omitnil,omitempty" name:"StreamUrl"`
}

func (r *UpdateStreamIngestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateStreamIngestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	delete(f, "StreamUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateStreamIngestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateStreamIngestResponseParams struct {
	// 任务的状态信息。InProgress：表示当前任务正在进行中。NotExist：表示当前任务不存在。示例值：InProgress
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateStreamIngestResponse struct {
	*tchttp.BaseResponse
	Response *UpdateStreamIngestResponseParams `json:"Response"`
}

func (r *UpdateStreamIngestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateStreamIngestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInformation struct {
	// 房间号
	RoomStr *string `json:"RoomStr,omitnil,omitempty" name:"RoomStr"`

	// 用户Id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户进房时间
	JoinTs *uint64 `json:"JoinTs,omitnil,omitempty" name:"JoinTs"`

	// 用户退房时间，用户没有退房则返回当前时间
	LeaveTs *uint64 `json:"LeaveTs,omitnil,omitempty" name:"LeaveTs"`

	// 终端类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Sdk版本号
	SdkVersion *string `json:"SdkVersion,omitnil,omitempty" name:"SdkVersion"`

	// 客户端IP地址
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 判断用户是否已经离开房间
	Finished *bool `json:"Finished,omitnil,omitempty" name:"Finished"`
}

type UserMediaStream struct {
	// TRTC用户参数。
	UserInfo *MixUserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 主辅路流类型，0为摄像头，1为屏幕分享，不填默认为0。
	StreamType *uint64 `json:"StreamType,omitnil,omitempty" name:"StreamType"`
}

type VideoEncode struct {
	// 输出流宽，音视频输出时必填。取值范围[0,1920]，单位为像素值。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 输出流高，音视频输出时必填。取值范围[0,1080]，单位为像素值。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 输出流帧率，音视频输出时必填。取值范围[1,60]，表示混流的输出帧率可选范围为1到60fps。
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 输出流码率，音视频输出时必填。取值范围[1,10000]，单位为kbps。
	BitRate *uint64 `json:"BitRate,omitnil,omitempty" name:"BitRate"`

	// 输出流gop，音视频输出时必填。取值范围[1,5]，单位为秒。
	Gop *uint64 `json:"Gop,omitnil,omitempty" name:"Gop"`
}

type VideoEncodeParams struct {
	// 宽。取值范围[0,1920]，单位为像素值。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 高。取值范围[0,1080]，单位为像素值。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 帧率。取值范围[1,60]，表示帧率可选范围为1到60fps。
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 码率。取值范围[1,10000]，单位为kbps。
	BitRate *uint64 `json:"BitRate,omitnil,omitempty" name:"BitRate"`

	// gop。取值范围[1,2]，单位为秒。
	Gop *uint64 `json:"Gop,omitnil,omitempty" name:"Gop"`
}

type VideoParams struct {
	// 视频的宽度值，单位为像素，默认值360。不能超过1920，与height的乘积不能超过1920*1080。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频的高度值，单位为像素，默认值640。不能超过1920，与width的乘积不能超过1920*1080。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 视频的帧率，范围[1, 60]，默认15。
	Fps *uint64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 视频的码率,单位是bps，范围[64000, 8192000]，默认550000bps。
	BitRate *uint64 `json:"BitRate,omitnil,omitempty" name:"BitRate"`

	// 视频关键帧时间间隔，单位秒，默认值10秒。
	Gop *uint64 `json:"Gop,omitnil,omitempty" name:"Gop"`
}

type WaterMark struct {
	// 水印类型，0为图片（默认），1为文字，2为时间戳。
	WaterMarkType *uint64 `json:"WaterMarkType,omitnil,omitempty" name:"WaterMarkType"`

	// 水印为图片时的参数列表，水印为图片时校验必填。
	WaterMarkImage *WaterMarkImage `json:"WaterMarkImage,omitnil,omitempty" name:"WaterMarkImage"`

	// 水印为文字时的参数列表，水印为文字时校验必填。
	WaterMarkChar *WaterMarkChar `json:"WaterMarkChar,omitnil,omitempty" name:"WaterMarkChar"`

	// 水印为时间戳时的参数列表，水印为时间戳时校验必填。
	WaterMarkTimestamp *WaterMarkTimestamp `json:"WaterMarkTimestamp,omitnil,omitempty" name:"WaterMarkTimestamp"`
}

type WaterMarkChar struct {
	// 文字水印的起始坐标Y值，从左上角开始
	Top *uint64 `json:"Top,omitnil,omitempty" name:"Top"`

	// 文字水印的起始坐标X值，从左上角开始
	Left *uint64 `json:"Left,omitnil,omitempty" name:"Left"`

	// 文字水印的宽度，单位像素值
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 文字水印的高度，单位像素值
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 水印文字的内容
	Chars *string `json:"Chars,omitnil,omitempty" name:"Chars"`

	// 水印文字的大小，单位像素，默认14
	FontSize *uint64 `json:"FontSize,omitnil,omitempty" name:"FontSize"`

	// 水印文字的颜色，默认白色
	FontColor *string `json:"FontColor,omitnil,omitempty" name:"FontColor"`

	// 水印文字的背景色，为空代表背景透明，默认为空
	BackGroundColor *string `json:"BackGroundColor,omitnil,omitempty" name:"BackGroundColor"`

	// 文字水印的字体，支持设置以下值：
	// 1. Tencent （默认）
	// 2. SourceHanSans
	Font *string `json:"Font,omitnil,omitempty" name:"Font"`
}

type WaterMarkImage struct {
	// 下载的url地址， 只支持jpg, png, jpeg，大小限制不超过5M。注意，url必须携带格式后缀，url内只支持特定的字符串, 范围是a-z A-Z 0-9 '-', '.', '_', '~', ':', '/', '?', '#', '[', ']' '@', '!', '&', '(', ')', '*', '+', ',', '%', '='
	WaterMarkUrl *string `json:"WaterMarkUrl,omitnil,omitempty" name:"WaterMarkUrl"`

	// 画布上该画面左上角的 y 轴坐标，取值范围 [0, 2560]，不能超过画布的高。
	Top *uint64 `json:"Top,omitnil,omitempty" name:"Top"`

	// 画布上该画面左上角的 x 轴坐标，取值范围 [0, 2560]，不能超过画布的宽。
	Left *uint64 `json:"Left,omitnil,omitempty" name:"Left"`

	// 画布上该画面宽度的相对值，取值范围 [0, 2560]，与Left相加不应超过画布的宽。
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 画布上该画面高度的相对值，取值范围 [0, 2560]，与Top相加不应超过画布的高。
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type WaterMarkParams struct {
	// 混流-水印图片ID。取值为实时音视频控制台上传的图片ID。
	WaterMarkId *uint64 `json:"WaterMarkId,omitnil,omitempty" name:"WaterMarkId"`

	// 混流-水印宽。单位为像素值。水印宽+X偏移不能超过整个画布宽。
	WaterMarkWidth *uint64 `json:"WaterMarkWidth,omitnil,omitempty" name:"WaterMarkWidth"`

	// 混流-水印高。单位为像素值。水印高+Y偏移不能超过整个画布高。
	WaterMarkHeight *uint64 `json:"WaterMarkHeight,omitnil,omitempty" name:"WaterMarkHeight"`

	// 水印在输出时的X偏移。单位为像素值。水印宽+X偏移不能超过整个画布宽。
	LocationX *uint64 `json:"LocationX,omitnil,omitempty" name:"LocationX"`

	// 水印在输出时的Y偏移。单位为像素值。水印高+Y偏移不能超过整个画布高。
	LocationY *uint64 `json:"LocationY,omitnil,omitempty" name:"LocationY"`

	// 混流-水印图片URL地址，支持png、jpg、jpeg、bmp格式，暂不支持透明通道。URL链接长度限制为512字节。WaterMarkUrl和WaterMarkId参数都填时，以WaterMarkUrl为准。图片大小限制不超过2MB。
	WaterMarkUrl *string `json:"WaterMarkUrl,omitnil,omitempty" name:"WaterMarkUrl"`
}

type WaterMarkTimestamp struct {
	// 时间戳的位置，取值范围0-6，分别代表上左，上右，下左，下右，上居中，下居中，居中
	Pos *uint64 `json:"Pos,omitnil,omitempty" name:"Pos"`

	// 显示时间戳的时区，默认东八区
	TimeZone *uint64 `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// 文字水印的字体，支持设置以下值：
	// 1. Tencent （默认）
	// 2. SourceHanSans
	Font *string `json:"Font,omitnil,omitempty" name:"Font"`
}

type WebRecordVideoParams struct {
	// 录制画面宽度，默认为1280，取值范围[0, 1920]
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 录制画面高度，默认为720，取值范围[0, 1080]
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 指定输出格式，可选hls,mp4。存储到云点播VOD时此参数无效，存储到VOD时请通过TencentVod（https://cloud.tencent.com/document/api/647/44055#TencentVod）内的MediaType设置。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 如果是aac或者mp4文件格式，超过长度限制后，系统会自动拆分视频文件。单位：分钟。默认为1440min（24h），取值范围为1-1440。【单文件限制最大为2G，满足文件大小 >2G 或录制时长度 > 24h任意一个条件，文件都会自动切分】
	// Hls 格式录制此参数不生效。
	// 示例值：1440
	MaxMediaFileDuration *int64 `json:"MaxMediaFileDuration,omitnil,omitempty" name:"MaxMediaFileDuration"`
}