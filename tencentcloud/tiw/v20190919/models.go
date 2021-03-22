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

package v20190919

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Canvas struct {

	// 混流画布宽高配置
	LayoutParams *LayoutParams `json:"LayoutParams,omitempty" name:"LayoutParams"`

	// 背景颜色，默认为黑色，格式为RGB格式，如红色为"#FF0000"
	BackgroundColor *string `json:"BackgroundColor,omitempty" name:"BackgroundColor"`
}

type Concat struct {

	// 是否开启拼接功能
	// 在开启了视频拼接功能的情况下，实时录制服务会把同一个用户因为暂停导致的多段视频拼接成一个视频
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 视频拼接时使用的垫片图片下载地址，不填默认用全黑的图片进行视频垫片
	Image *string `json:"Image,omitempty" name:"Image"`
}

type CreateTranscodeRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需要进行转码文件地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 是否为静态PPT，默认为False；
	// 如果IsStaticPPT为False，后缀名为.ppt或.pptx的文档会动态转码成HTML5页面，其他格式的文档会静态转码成图片；如果IsStaticPPT为True，所有格式的文档会静态转码成图片；
	IsStaticPPT *bool `json:"IsStaticPPT,omitempty" name:"IsStaticPPT"`

	// 转码后文档的最小分辨率，不传、传空字符串或分辨率格式错误则使用文档原分辨率
	// 
	// 注意分辨率宽高中间为英文字母"xyz"的"x"
	MinResolution *string `json:"MinResolution,omitempty" name:"MinResolution"`

	// 动态PPT转码可以为文件生成该分辨率的缩略图，不传、传空字符串或分辨率格式错误则不生成缩略图，分辨率格式同MinResolution
	ThumbnailResolution *string `json:"ThumbnailResolution,omitempty" name:"ThumbnailResolution"`

	// 转码文件压缩格式，不传、传空字符串或不是指定的格式则不生成压缩文件，目前支持如下压缩格式：
	// 
	// zip： 生成`.zip`压缩包
	// tar.gz： 生成`.tar.gz`压缩包
	CompressFileType *string `json:"CompressFileType,omitempty" name:"CompressFileType"`

	// 内部参数
	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`
}

func (r *CreateTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTranscodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文档转码任务的唯一标识Id，用于查询该任务的进度以及转码结果
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTranscodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVideoGenerationTaskRequest struct {
	*tchttp.BaseRequest

	// 录制任务的TaskId
	OnlineRecordTaskId *string `json:"OnlineRecordTaskId,omitempty" name:"OnlineRecordTaskId"`

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 视频生成的白板参数，例如白板宽高等。
	// 
	// 此参数与开始录制接口提供的Whiteboard参数互斥，在本接口与开始录制接口都提供了Whiteboard参数时，优先使用本接口指定的Whiteboard参数进行视频生成，否则使用开始录制接口提供的Whiteboard参数进行视频生成。
	Whiteboard *Whiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`

	// 视频拼接参数
	// 
	// 此参数与开始录制接口提供的Concat参数互斥，在本接口与开始录制接口都提供了Concat参数时，优先使用本接口指定的Concat参数进行视频拼接，否则使用开始录制接口提供的Concat参数进行视频拼接。
	Concat *Concat `json:"Concat,omitempty" name:"Concat"`

	// 视频生成混流参数
	// 
	// 此参数与开始录制接口提供的MixStream参数互斥，在本接口与开始录制接口都提供了MixStream参数时，优先使用本接口指定的MixStream参数进行视频混流，否则使用开始录制接口提供的MixStream参数进行视频拼混流。
	MixStream *MixStream `json:"MixStream,omitempty" name:"MixStream"`

	// 视频生成控制参数，用于更精细地指定需要生成哪些流，某一路流是否禁用音频，是否只录制小画面等
	// 
	// 此参数与开始录制接口提供的RecordControl参数互斥，在本接口与开始录制接口都提供了RecordControl参数时，优先使用本接口指定的RecordControl参数进行视频生成控制，否则使用开始录制接口提供的RecordControl参数进行视频拼生成控制。
	RecordControl *RecordControl `json:"RecordControl,omitempty" name:"RecordControl"`

	// 内部参数
	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`
}

func (r *CreateVideoGenerationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVideoGenerationTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVideoGenerationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 视频生成的任务Id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVideoGenerationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVideoGenerationTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomLayout struct {

	// 混流画布参数
	Canvas *Canvas `json:"Canvas,omitempty" name:"Canvas"`

	// 流布局参数，每路流的布局不能超出画布区域
	InputStreamList []*StreamLayout `json:"InputStreamList,omitempty" name:"InputStreamList" list`
}

type DescribeOnlineRecordCallbackRequest struct {
	*tchttp.BaseRequest

	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeOnlineRecordCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOnlineRecordCallbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOnlineRecordCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实时录制事件回调地址，如果未设置回调地址，该字段为空字符串
		Callback *string `json:"Callback,omitempty" name:"Callback"`

		// 实时录制回调鉴权密钥
		CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOnlineRecordCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOnlineRecordCallbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOnlineRecordRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 实时录制任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOnlineRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制结束原因，
	// - AUTO: 房间内长时间没有音视频上行及白板操作导致自动停止录制
	// - USER_CALL: 主动调用了停止录制接口
	// - EXCEPTION: 录制异常结束
		FinishReason *string `json:"FinishReason,omitempty" name:"FinishReason"`

		// 需要查询结果的录制任务Id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 录制任务状态
	// - PREPARED: 表示录制正在准备中（进房/启动录制服务等操作）
	// - RECORDING: 表示录制已开始
	// - PAUSED: 表示录制已暂停
	// - STOPPED: 表示录制已停止，正在处理并上传视频
	// - FINISHED: 表示视频处理并上传完成，成功生成录制结果
		Status *string `json:"Status,omitempty" name:"Status"`

		// 房间号
		RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

		// 白板的群组 Id
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 录制用户Id
		RecordUserId *string `json:"RecordUserId,omitempty" name:"RecordUserId"`

		// 实际开始录制时间，Unix 时间戳，单位秒
		RecordStartTime *int64 `json:"RecordStartTime,omitempty" name:"RecordStartTime"`

		// 实际停止录制时间，Unix 时间戳，单位秒
		RecordStopTime *int64 `json:"RecordStopTime,omitempty" name:"RecordStopTime"`

		// 回放视频总时长（单位：毫秒）
		TotalTime *int64 `json:"TotalTime,omitempty" name:"TotalTime"`

		// 录制过程中出现异常的次数
		ExceptionCnt *int64 `json:"ExceptionCnt,omitempty" name:"ExceptionCnt"`

		// 拼接视频中被忽略的时间段，只有开启视频拼接功能的时候，这个参数才是有效的
		OmittedDurations []*OmittedDuration `json:"OmittedDurations,omitempty" name:"OmittedDurations" list`

		// 录制视频列表
		VideoInfos []*VideoInfo `json:"VideoInfos,omitempty" name:"VideoInfos" list`

		// 回放URL，需配合信令播放器使用。此字段仅适用于`视频生成模式`
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReplayUrl *string `json:"ReplayUrl,omitempty" name:"ReplayUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOnlineRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTranscodeCallbackRequest struct {
	*tchttp.BaseRequest

	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeTranscodeCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTranscodeCallbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTranscodeCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文档转码回调地址
		Callback *string `json:"Callback,omitempty" name:"Callback"`

		// 文档转码回调鉴权密钥
		CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTranscodeCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTranscodeCallbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTranscodeRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 文档转码任务的唯一标识Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTranscodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文档的总页数
		Pages *int64 `json:"Pages,omitempty" name:"Pages"`

		// 转码的当前进度,取值范围为0~100
		Progress *int64 `json:"Progress,omitempty" name:"Progress"`

		// 文档的分辨率
		Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

		// 转码完成后结果的URL
	// 动态转码：PPT转动态H5的链接
	// 静态转码：文档每一页的图片URL前缀，比如，该URL前缀为`http://example.com/g0jb42ps49vtebjshilb/`，那么文档第1页的图片URL为
	// `http://example.com/g0jb42ps49vtebjshilb/1.jpg`，其它页以此类推
		ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

		// 任务的当前状态
	// - QUEUED: 正在排队等待转换
	// - PROCESSING: 转换中
	// - FINISHED: 转换完成
		Status *string `json:"Status,omitempty" name:"Status"`

		// 转码任务的唯一标识Id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 文档的文件名
		Title *string `json:"Title,omitempty" name:"Title"`

		// 缩略图URL前缀，比如，该URL前缀为`http://example.com/g0jb42ps49vtebjshilb/ `，那么动态PPT第1页的缩略图URL为
	// `http://example.com/g0jb42ps49vtebjshilb/1.jpg`，其它页以此类推
	// 
	// 如果发起文档转码请求参数中带了ThumbnailResolution参数，并且转码类型为动态转码，该参数不为空，其余情况该参数为空字符串
		ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" name:"ThumbnailUrl"`

		// 动态转码缩略图生成分辨率
		ThumbnailResolution *string `json:"ThumbnailResolution,omitempty" name:"ThumbnailResolution"`

		// 转码压缩文件下载的URL，如果发起文档转码请求参数中`CompressFileType`为空或者不是支持的压缩格式，该参数为空字符串
		CompressFileUrl *string `json:"CompressFileUrl,omitempty" name:"CompressFileUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTranscodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoGenerationTaskCallbackRequest struct {
	*tchttp.BaseRequest

	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeVideoGenerationTaskCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVideoGenerationTaskCallbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoGenerationTaskCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制视频生成回调地址
		Callback *string `json:"Callback,omitempty" name:"Callback"`

		// 录制视频生成回调鉴权密钥
		CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVideoGenerationTaskCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVideoGenerationTaskCallbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoGenerationTaskRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 录制视频生成的任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeVideoGenerationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVideoGenerationTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoGenerationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务对应的群组Id
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 任务对应的房间号
		RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

		// 任务的Id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 已废弃
		Progress *int64 `json:"Progress,omitempty" name:"Progress"`

		// 录制视频生成任务状态
	// - QUEUED: 正在排队
	// - PROCESSING: 正在生成视频
	// - FINISHED: 生成视频结束（成功完成或失败结束，可以通过错误码和错误信息进一步判断）
		Status *string `json:"Status,omitempty" name:"Status"`

		// 回放视频总时长,单位：毫秒
		TotalTime *int64 `json:"TotalTime,omitempty" name:"TotalTime"`

		// 已废弃，请使用`VideoInfoList`参数
		VideoInfos *VideoInfo `json:"VideoInfos,omitempty" name:"VideoInfos"`

		// 录制视频生成视频列表
		VideoInfoList []*VideoInfo `json:"VideoInfoList,omitempty" name:"VideoInfoList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVideoGenerationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVideoGenerationTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteboardPushCallbackRequest struct {
	*tchttp.BaseRequest

	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeWhiteboardPushCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWhiteboardPushCallbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteboardPushCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 白板推流事件回调地址，如果未设置回调地址，该字段为空字符串
		Callback *string `json:"Callback,omitempty" name:"Callback"`

		// 白板推流回调鉴权密钥
		CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWhiteboardPushCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWhiteboardPushCallbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteboardPushRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 白板推流任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeWhiteboardPushRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWhiteboardPushRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteboardPushResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 推流结束原因，
	// - AUTO: 房间内长时间没有音视频上行及白板操作导致自动停止推流
	// - USER_CALL: 主动调用了停止推流接口
	// - EXCEPTION: 推流异常结束
		FinishReason *string `json:"FinishReason,omitempty" name:"FinishReason"`

		// 需要查询结果的白板推流任务Id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 推流任务状态
	// - PREPARED: 表示推流正在准备中（进房/启动推流服务等操作）
	// - PUSHING: 表示推流已开始
	// - STOPPED: 表示推流已停止
		Status *string `json:"Status,omitempty" name:"Status"`

		// 房间号
		RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

		// 白板的群组 Id
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 推流用户Id
		PushUserId *string `json:"PushUserId,omitempty" name:"PushUserId"`

		// 实际开始推流时间，Unix 时间戳，单位秒
		PushStartTime *int64 `json:"PushStartTime,omitempty" name:"PushStartTime"`

		// 实际停止推流时间，Unix 时间戳，单位秒
		PushStopTime *int64 `json:"PushStopTime,omitempty" name:"PushStopTime"`

		// 推流过程中出现异常的次数
		ExceptionCnt *int64 `json:"ExceptionCnt,omitempty" name:"ExceptionCnt"`

		// 白板推流首帧对应的IM时间戳，可用于录制回放时IM聊天消息与白板推流视频进行同步对时。
		IMSyncTime *int64 `json:"IMSyncTime,omitempty" name:"IMSyncTime"`

		// 备份推流任务结果信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Backup *string `json:"Backup,omitempty" name:"Backup"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWhiteboardPushResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWhiteboardPushResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LayoutParams struct {

	// 流画面宽，取值范围[2,3000]
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 流画面高，取值范围[2,3000]
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 当前画面左上角顶点相对于Canvas左上角顶点的x轴偏移量，默认为0，取值范围[0,3000]
	X *int64 `json:"X,omitempty" name:"X"`

	// 当前画面左上角顶点相对于Canvas左上角顶点的y轴偏移量，默认为0， 取值范围[0,3000]
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 画面z轴位置，默认为0
	// z轴确定了重叠画面的遮盖顺序，z轴值大的画面处于顶层
	ZOrder *int64 `json:"ZOrder,omitempty" name:"ZOrder"`
}

type MixStream struct {

	// 是否开启混流
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 是否禁用音频混流
	DisableAudio *bool `json:"DisableAudio,omitempty" name:"DisableAudio"`

	// 内置混流布局模板ID, 取值 [1, 2], 区别见内置混流布局模板样式示例说明
	// 在没有填Custom字段时候，ModelId是必填的
	ModelId *int64 `json:"ModelId,omitempty" name:"ModelId"`

	// 老师用户ID
	// 此字段只有在ModelId填了的情况下生效
	// 填写TeacherId的效果是把指定为TeacherId的用户视频流显示在内置模板的第一个小画面中
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 自定义混流布局参数
	// 当此字段存在时，ModelId 及 TeacherId 字段将被忽略
	Custom *CustomLayout `json:"Custom,omitempty" name:"Custom"`
}

type OmittedDuration struct {

	// 录制暂停时间戳对应的视频播放时间(单位: 毫秒)
	VideoTime *int64 `json:"VideoTime,omitempty" name:"VideoTime"`

	// 录制暂停时间戳(单位: 毫秒)
	PauseTime *int64 `json:"PauseTime,omitempty" name:"PauseTime"`

	// 录制恢复时间戳(单位: 毫秒)
	ResumeTime *int64 `json:"ResumeTime,omitempty" name:"ResumeTime"`
}

type PauseOnlineRecordRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 实时录制任务 Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *PauseOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PauseOnlineRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PauseOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PauseOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PauseOnlineRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecordControl struct {

	// 设置是否开启录制控制参数，只有设置为true的时候，录制控制参数才生效。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 设置是否禁用录制的全局控制参数。一般与`StreamControls`参数配合使用。
	// 
	// true - 所有流都不录制。
	// false - 所有流都录制。默认为false。
	// 
	// 这里的设置对所有流都生效，如果同时在 `StreamControls` 列表中针对指定流设置了控制参数，则优先采用`StreamControls`中设置的控制参数。
	DisableRecord *bool `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// 设置是否禁用所有流的音频录制的全局控制参数。一般与`StreamControls`参数配合使用。
	// 
	// true - 所有流的录制都不对音频进行录制。
	// false - 所有流的录制都需要对音频进行录制。默认为false。
	// 
	// 这里的设置对所有流都生效，如果同时在 `StreamControls` 列表中针对指定流设置了控制参数，则优先采用`StreamControls`中设置的控制参数。
	DisableAudio *bool `json:"DisableAudio,omitempty" name:"DisableAudio"`

	// 设置是否所有流都只录制小画面的全局控制参数。一般与`StreamControls`参数配合使用。
	// 
	// true - 所有流都只录制小画面。设置为true时，请确保上行端在推流的时候同时上行了小画面，否则录制视频可能是黑屏。
	// false - 所有流都录制大画面，默认为false。
	// 
	// 这里的设置对所有流都生效，如果同时在 `StreamControls` 列表中针对指定流设置了控制参数，则优先采用`StreamControls`中设置的控制参数。
	PullSmallVideo *bool `json:"PullSmallVideo,omitempty" name:"PullSmallVideo"`

	// 针对具体流指定控制参数，如果列表为空，则所有流采用全局配置的控制参数进行录制。列表不为空，则列表中指定的流将优先按此列表指定的控制参数进行录制。
	StreamControls []*StreamControl `json:"StreamControls,omitempty" name:"StreamControls" list`
}

type ResumeOnlineRecordRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 恢复录制的实时录制任务 Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ResumeOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeOnlineRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeOnlineRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetOnlineRecordCallbackKeyRequest struct {
	*tchttp.BaseRequest

	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 设置实时录制回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥。回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *SetOnlineRecordCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetOnlineRecordCallbackKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetOnlineRecordCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetOnlineRecordCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetOnlineRecordCallbackKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetOnlineRecordCallbackRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 实时录制任务结果回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持 http或https协议，即回调地址以http://或https://开头。回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40258
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

func (r *SetOnlineRecordCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetOnlineRecordCallbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetOnlineRecordCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetOnlineRecordCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetOnlineRecordCallbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetTranscodeCallbackKeyRequest struct {
	*tchttp.BaseRequest

	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 设置文档转码回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *SetTranscodeCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetTranscodeCallbackKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetTranscodeCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetTranscodeCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetTranscodeCallbackKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetTranscodeCallbackRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 文档转码进度回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持http或https协议，即回调地址以http://或https://开头。
	// 回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40260
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

func (r *SetTranscodeCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetTranscodeCallbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetTranscodeCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetTranscodeCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetTranscodeCallbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetVideoGenerationTaskCallbackKeyRequest struct {
	*tchttp.BaseRequest

	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 设置视频生成回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *SetVideoGenerationTaskCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetVideoGenerationTaskCallbackKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetVideoGenerationTaskCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetVideoGenerationTaskCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetVideoGenerationTaskCallbackKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetVideoGenerationTaskCallbackRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 课后录制任务结果回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持 http或https协议，即回调地址以http://或https://开头
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

func (r *SetVideoGenerationTaskCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetVideoGenerationTaskCallbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetVideoGenerationTaskCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetVideoGenerationTaskCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetVideoGenerationTaskCallbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetWhiteboardPushCallbackKeyRequest struct {
	*tchttp.BaseRequest

	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 设置白板推流回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥。回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *SetWhiteboardPushCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetWhiteboardPushCallbackKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetWhiteboardPushCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetWhiteboardPushCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetWhiteboardPushCallbackKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetWhiteboardPushCallbackRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 白板推流任务结果回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持 http或https协议，即回调地址以http://或https://开头。回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

func (r *SetWhiteboardPushCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetWhiteboardPushCallbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetWhiteboardPushCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetWhiteboardPushCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetWhiteboardPushCallbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartOnlineRecordRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需要录制的房间号，取值范围: (1, 4294967295)
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// 用于录制服务进房的用户ID，最大长度不能大于60个字节，格式为`tic_record_user_${RoomId}_${Random}`，其中 `${RoomId} `与录制房间号对应，`${Random}`为一个随机字符串。
	// 该ID必须是一个单独的未在SDK中使用的ID，录制服务使用这个用户ID进入房间进行音视频与白板录制，若该ID和SDK中使用的ID重复，会导致SDK和录制服务互踢，影响正常录制。
	RecordUserId *string `json:"RecordUserId,omitempty" name:"RecordUserId"`

	// 与RecordUserId对应的签名
	RecordUserSig *string `json:"RecordUserSig,omitempty" name:"RecordUserSig"`

	// （已废弃，设置无效）白板的 IM 群组 Id，默认同房间号
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 录制视频拼接参数
	Concat *Concat `json:"Concat,omitempty" name:"Concat"`

	// 录制白板参数，例如白板宽高等
	Whiteboard *Whiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`

	// 录制混流参数
	// 特别说明：
	// 1. 混流功能需要根据额外开通， 请联系腾讯云互动白板客服人员
	// 2. 使用混流功能，必须提供 Extras 参数，且 Extras 参数中必须包含 "MIX_STREAM"
	MixStream *MixStream `json:"MixStream,omitempty" name:"MixStream"`

	// 使用到的高级功能列表
	// 可以选值列表：
	// MIX_STREAM - 混流功能
	Extras []*string `json:"Extras,omitempty" name:"Extras" list`

	// 是否需要在结果回调中返回各路流的纯音频录制文件，文件格式为mp3
	AudioFileNeeded *bool `json:"AudioFileNeeded,omitempty" name:"AudioFileNeeded"`

	// 录制控制参数，用于更精细地指定需要录制哪些流，某一路流是否禁用音频，是否只录制小画面等
	RecordControl *RecordControl `json:"RecordControl,omitempty" name:"RecordControl"`

	// 录制模式
	// 
	// REALTIME_MODE - 实时录制模式（默认）
	// VIDEO_GENERATION_MODE - 视频生成模式（内测中，需邮件申请开通）
	RecordMode *string `json:"RecordMode,omitempty" name:"RecordMode"`

	// 聊天群组ID，此字段仅适用于`视频生成模式`
	// 
	// 在`视频生成模式`下，默认会记录白板群组内的非白板信令消息，如果指定了`ChatGroupId`，则会记录指定群ID的聊天消息。
	ChatGroupId *string `json:"ChatGroupId,omitempty" name:"ChatGroupId"`

	// 自动停止录制超时时间，单位秒，取值范围[300, 86400], 默认值为300秒。
	// 
	// 当超过设定时间房间内没有音视频上行且没有白板操作的时候，录制服务会自动停止当前录制任务。
	AutoStopTimeout *int64 `json:"AutoStopTimeout,omitempty" name:"AutoStopTimeout"`

	// 内部参数，可忽略
	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`
}

func (r *StartOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartOnlineRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制任务Id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartOnlineRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartWhiteboardPushRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需要推流白板的房间号，取值范围: (1, 4294967295)
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// 用于白板推流服务进房进行推流的用户ID，最大长度不能大于60个字节，该ID必须是一个单独的未在SDK中使用的ID，白板推流服务使用这个用户ID进入房间进行白板音视频推流，若该ID和SDK中使用的ID重复，会导致SDK和白板推流服务互踢，影响正常推流。
	PushUserId *string `json:"PushUserId,omitempty" name:"PushUserId"`

	// 与PushUserId对应的签名
	PushUserSig *string `json:"PushUserSig,omitempty" name:"PushUserSig"`

	// 白板参数，例如白板宽高、背景颜色等
	Whiteboard *Whiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`

	// 自动停止推流超时时间，单位秒，取值范围[300, 259200], 默认值为1800秒。
	// 
	// 当白板超过设定时间没有操作的时候，白板推流服务会自动停止白板推流。
	AutoStopTimeout *int64 `json:"AutoStopTimeout,omitempty" name:"AutoStopTimeout"`

	// 对主白板推流任务进行操作时，是否同时同步操作备份任务
	AutoManageBackup *bool `json:"AutoManageBackup,omitempty" name:"AutoManageBackup"`

	// 备份白板推流相关参数。
	// 
	// 指定了备份参数的情况下，白板推流服务会在房间内新增一路白板画面视频流，即同一个房间内会有两路白板画面推流。
	Backup *WhiteboardPushBackupParam `json:"Backup,omitempty" name:"Backup"`

	// TRTC高级权限控制参数，如果在实时音视频开启了高级权限控制功能，必须提供PrivateMapKey才能保证正常推流。
	PrivateMapKey *string `json:"PrivateMapKey,omitempty" name:"PrivateMapKey"`

	// 白板推流视频帧率，取值范围[0, 30]，默认20fps
	VideoFPS *int64 `json:"VideoFPS,omitempty" name:"VideoFPS"`

	// 白板推流码率， 取值范围[0, 2000]，默认1200kbps。
	// 
	// 这里的码率设置是一个参考值，实际推流的时候使用的是动态码率，所以真实码率不会固定为指定值，会在指定值附近波动。
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// 在实时音视频云端录制模式选择为 `指定用户录制` 模式的时候是否自动录制白板推流。
	// 
	// 默认在实时音视频的云端录制模式选择为 `指定用户录制` 模式的情况下，不会自动进行白板推流录制，如果希望进行白板推流录制，请将此参数设置为true。
	// 
	// 如果实时音视频的云端录制模式选择为 `全局自动录制` 模式，可忽略此参数。
	AutoRecord *bool `json:"AutoRecord,omitempty" name:"AutoRecord"`

	// 指定白板推流录制的RecordID，指定的RecordID会用于填充实时音视频云端录制完成后的回调消息中的 "userdefinerecordid" 字段内容，便于您更方便的识别录制回调，以及在点播媒体资源管理中查找相应的录制视频文件。
	// 
	// 限制长度为64字节，只允许包含大小写英文字母（a-zA-Z）、数字（0-9）及下划线和连词符。
	// 
	// 此字段设置后，不管`AutoRecord`字段取值如何，都将自动进行白板推流录制。
	// 
	// 默认RecordId生成规则如下：
	// urlencode(SdkAppID_RoomID_PushUserID)
	// 
	// 例如：
	// SdkAppID = 12345678，RoomID = 12345，PushUserID = push_user_1
	// 那么：RecordId = 12345678_12345_push_user_1
	UserDefinedRecordId *string `json:"UserDefinedRecordId,omitempty" name:"UserDefinedRecordId"`

	// 在实时音视频旁路推流模式选择为`指定用户旁路`模式的时候，是否自动旁路白板推流。
	// 
	// 默认在实时音视频的旁路推流模式选择为 `指定用户旁路` 模式的情况下，不会自动旁路白板推流，如果希望旁路白板推流，请将此参数设置为true。
	// 
	// 如果实时音视频的旁路推流模式选择为 `全局自动旁路` 模式，可忽略此参数。
	AutoPublish *bool `json:"AutoPublish,omitempty" name:"AutoPublish"`

	// 指定实时音视频在旁路白板推流时的StreamID，设置之后，您就可以在腾讯云直播 CDN 上通过标准直播方案（FLV或HLS）播放该用户的音视频流。
	// 
	// 限制长度为64字节，只允许包含大小写英文字母（a-zA-Z）、数字（0-9）及下划线和连词符。
	// 
	// 此字段设置后，不管`AutoPublish`字段取值如何，都将自动旁路白板推流。
	// 
	// 默认StreamID生成规则如下：
	// urlencode(SdkAppID_RoomID_PushUserID_main)
	// 
	// 例如：
	// SdkAppID = 12345678，RoomID = 12345，PushUserID = push_user_1
	// 那么：StreamID = 12345678_12345_push_user_1_main
	UserDefinedStreamId *string `json:"UserDefinedStreamId,omitempty" name:"UserDefinedStreamId"`

	// 内部参数，不需要关注此参数
	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`
}

func (r *StartWhiteboardPushRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartWhiteboardPushRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartWhiteboardPushResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 推流任务Id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 备份任务结果参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Backup *string `json:"Backup,omitempty" name:"Backup"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartWhiteboardPushResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartWhiteboardPushResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopOnlineRecordRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需要停止录制的任务 Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopOnlineRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopOnlineRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopWhiteboardPushRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需要停止的白板推流任务 Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopWhiteboardPushRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopWhiteboardPushRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopWhiteboardPushResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 备份任务相关参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Backup *string `json:"Backup,omitempty" name:"Backup"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopWhiteboardPushResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopWhiteboardPushResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StreamControl struct {

	// 视频流ID
	// 视频流ID的取值含义如下：
	// 1. tic_record_user - 表示白板视频流
	// 2. tic_substream - 表示辅路视频流
	// 3. 特定用户ID - 表示指定用户的视频流
	// 
	// 在实际录制过程中，视频流ID的匹配规则为前缀匹配，只要真实流ID的前缀与指定的流ID一致就认为匹配成功。
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// 设置是否对此路流开启录制。
	// 
	// true - 表示不对这路流进行录制，录制结果将不包含这路流的视频。
	// false - 表示需要对这路流进行录制，录制结果会包含这路流的视频。
	// 
	// 默认为 false。
	DisableRecord *bool `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// 设置是否禁用这路流的音频录制。
	// 
	// true - 表示不对这路流的音频进行录制，录制结果里这路流的视频将会没有声音。
	// false - 录制视频会保留音频，如果设置为true，则录制视频会丢弃这路流的音频。
	// 
	// 默认为 false。
	DisableAudio *bool `json:"DisableAudio,omitempty" name:"DisableAudio"`

	// 设置当前流录制视频是否只录制小画面。
	// 
	// true - 录制小画面。设置为true时，请确保上行端同时上行了小画面，否则录制视频可能是黑屏。
	// false - 录制大画面。
	// 
	// 默认为 false。
	PullSmallVideo *bool `json:"PullSmallVideo,omitempty" name:"PullSmallVideo"`
}

type StreamLayout struct {

	// 流布局配置参数
	LayoutParams *LayoutParams `json:"LayoutParams,omitempty" name:"LayoutParams"`

	// 视频流ID
	// 流ID的取值含义如下：
	// 1. tic_record_user - 表示当前画面用于显示白板视频流
	// 2. tic_substream - 表示当前画面用于显示辅路视频流
	// 3. 特定用户ID - 表示当前画面用于显示指定用户的视频流
	// 4. 不填 - 表示当前画面用于备选，当有新的视频流加入时，会从这些备选的空位中选择一个没有被占用的位置来显示新的视频流画面
	InputStreamId *string `json:"InputStreamId,omitempty" name:"InputStreamId"`

	// 背景颜色，默认为黑色，格式为RGB格式，如红色为"#FF0000"
	BackgroundColor *string `json:"BackgroundColor,omitempty" name:"BackgroundColor"`

	// 视频画面填充模式。
	// 
	// 0 - 自适应模式，对视频画面进行等比例缩放，在指定区域内显示完整的画面。此模式可能存在黑边。
	// 1 - 全屏模式，对视频画面进行等比例缩放，让画面填充满整个指定区域。此模式不会存在黑边，但会将超出区域的那一部分画面裁剪掉。
	FillMode *int64 `json:"FillMode,omitempty" name:"FillMode"`
}

type VideoInfo struct {

	// 视频开始播放的时间（单位：毫秒）
	VideoPlayTime *int64 `json:"VideoPlayTime,omitempty" name:"VideoPlayTime"`

	// 视频大小（字节）
	VideoSize *int64 `json:"VideoSize,omitempty" name:"VideoSize"`

	// 视频格式
	VideoFormat *string `json:"VideoFormat,omitempty" name:"VideoFormat"`

	// 视频播放时长（单位：毫秒）
	VideoDuration *int64 `json:"VideoDuration,omitempty" name:"VideoDuration"`

	// 视频文件URL
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// 视频文件Id
	VideoId *string `json:"VideoId,omitempty" name:"VideoId"`

	// 视频流类型 
	// - 0：摄像头视频 
	// - 1：屏幕分享视频
	// - 2：白板视频 
	// - 3：混流视频
	// - 4：纯音频（mp3)
	VideoType *int64 `json:"VideoType,omitempty" name:"VideoType"`

	// 摄像头/屏幕分享视频所属用户的 Id（白板视频为空、混流视频tic_mixstream_房间号_混流布局类型、辅路视频tic_substream_用户Id）
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 视频分辨率的宽
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 视频分辨率的高
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type Whiteboard struct {

	// 实时录制结果里白板视频宽，默认为1280
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 实时录制结果里白板视频高，默认为960
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 白板初始化参数，透传到白板 SDK
	InitParam *string `json:"InitParam,omitempty" name:"InitParam"`
}

type WhiteboardPushBackupParam struct {

	// 用于白板推流服务进房的用户ID，
	// 该ID必须是一个单独的未在SDK中使用的ID，白板推流服务将使用这个用户ID进入房间进行白板推流，若该ID和SDK中使用的ID重复，会导致SDK和录制服务互踢，影响正常推流。
	PushUserId *string `json:"PushUserId,omitempty" name:"PushUserId"`

	// 与PushUserId对应的签名
	PushUserSig *string `json:"PushUserSig,omitempty" name:"PushUserSig"`
}
