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
	// 
	// 静态转码这个参数不起作用
	ThumbnailResolution *string `json:"ThumbnailResolution,omitempty" name:"ThumbnailResolution"`

	// 转码文件压缩格式，不传、传空字符串或不是指定的格式则不生成压缩文件，目前支持如下压缩格式：
	// 
	// zip： 生成`.zip`压缩包
	// tar.gz： 生成`.tar.gz`压缩包
	CompressFileType *string `json:"CompressFileType,omitempty" name:"CompressFileType"`
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

type SetOnlineRecordCallbackRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 实时录制任务结果回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持 http或https协议，即回调地址以http://或https://开头
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

type SetTranscodeCallbackRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 文档转码进度回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持http或https协议，即回调地址以http://或https://开头
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

type StartOnlineRecordRequest struct {
	*tchttp.BaseRequest

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需要录制的房间号
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// 用于实时录制服务进房的用户Id，格式为`tic_record_user_${RoomId}_${Random}`，其中 `${RoomId}` 与录制房间号对应，`${Random}`为一个随机字符串。
	// 实时录制服务会使用这个用户Id进房进行录制房间内的音视频与白板，为了防止进房冲突，请保证此 用户Id不重复
	RecordUserId *string `json:"RecordUserId,omitempty" name:"RecordUserId"`

	// 与RecordUserId对应的签名
	RecordUserSig *string `json:"RecordUserSig,omitempty" name:"RecordUserSig"`

	// 白板的 IM 群组 Id，默认同房间号
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 实时录制视频拼接参数
	Concat *Concat `json:"Concat,omitempty" name:"Concat"`

	// 实时录制白板参数，例如白板宽高等
	Whiteboard *Whiteboard `json:"Whiteboard,omitempty" name:"Whiteboard"`

	// 实时录制混流参数
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

		// 实时录制的任务Id
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
}

type Whiteboard struct {

	// 实时录制结果里白板视频宽，默认为1280
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 实时录制结果里白板视频高，默认为960
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 白板初始化参数，透传到白板 SDK
	InitParam *string `json:"InitParam,omitempty" name:"InitParam"`
}
