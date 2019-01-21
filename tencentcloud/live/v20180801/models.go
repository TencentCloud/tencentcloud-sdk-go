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

package v20180801

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddDelayLiveStreamRequest struct {
	*tchttp.BaseRequest

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 您的加速域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 延播时间，单位：秒，上限：600秒。
	DelayTime *uint64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// 延播设置的过期时间。UTC 格式，例如：2018-11-29T19:00:00Z。
	// 注意：默认7天后过期，且最长支持7天内生效。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *AddDelayLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddDelayLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddDelayLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddDelayLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddDelayLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// 水印图片url。
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// 水印名称。
	WatermarkName *string `json:"WatermarkName,omitempty" name:"WatermarkName"`

	// 显示位置,X轴偏移。
	XPosition *int64 `json:"XPosition,omitempty" name:"XPosition"`

	// 显示位置,Y轴偏移。
	YPosition *int64 `json:"YPosition,omitempty" name:"YPosition"`
}

func (r *AddLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddLiveWatermarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 水印ID。
		WatermarkId *uint64 `json:"WatermarkId,omitempty" name:"WatermarkId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddLiveWatermarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// 证书Id。
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// 播放域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 状态，0： 关闭  1：打开。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *BindLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindLiveDomainCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindLiveDomainCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindLiveDomainCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CallBackRuleInfo struct {

	// 规则创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 规则更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

type CallBackTemplateInfo struct {

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 开播回调URL。
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitempty" name:"StreamBeginNotifyUrl"`

	// 断流回调URL。
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitempty" name:"StreamEndNotifyUrl"`

	// 混流回调URL。
	StreamMixNotifyUrl *string `json:"StreamMixNotifyUrl,omitempty" name:"StreamMixNotifyUrl"`

	// 录制回调URL。
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitempty" name:"RecordNotifyUrl"`

	// 截图回调URL。
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitempty" name:"SnapshotNotifyUrl"`

	// 鉴黄回调URL。
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitempty" name:"PornCensorshipNotifyUrl"`
}

type CertInfo struct {

	// 证书Id。
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// 证书名称。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间，UTC格式。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 证书内容。
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// 证书类型。
	// 0：腾讯云托管证书
	// 1：用户添加证书。
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 证书过期时间，UTC格式。
	CertExpireTime *string `json:"CertExpireTime,omitempty" name:"CertExpireTime"`

	// 使用此证书的域名列表。
	DomainList []*string `json:"DomainList,omitempty" name:"DomainList" list`
}

type CreateLiveCallbackRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 模板ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveCallbackRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCallbackRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveCallbackRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCallbackRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名称。非空的字符串
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 开播回调URL。
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitempty" name:"StreamBeginNotifyUrl"`

	// 断流回调URL。
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitempty" name:"StreamEndNotifyUrl"`

	// 录制回调URL。
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitempty" name:"RecordNotifyUrl"`

	// 截图回调URL。
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitempty" name:"SnapshotNotifyUrl"`

	// 鉴黄回调URL。
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitempty" name:"PornCensorshipNotifyUrl"`
}

func (r *CreateLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCallbackTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板Id。
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCallbackTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCertRequest struct {
	*tchttp.BaseRequest

	// 证书类型。0-用户添加证书；1-腾讯云托管证书。
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// 证书内容，即公钥。
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// 私钥。
	HttpsKey *string `json:"HttpsKey,omitempty" name:"HttpsKey"`

	// 证书名称。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书ID
		CertId *int64 `json:"CertId,omitempty" name:"CertId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordRequest struct {
	*tchttp.BaseRequest

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 推流App名。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 推流域名。多域名推流必须设置。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 录制开始时间。非精彩视频录制，必须设置该字段。中国标准时间，需要URLEncode。如 2017-01-01 10:10:01，编码为：2017-01-01+10%3a10%3a01。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 录制结束时间。非精彩视频录制，必须设置该字段。中国标准时间，需要URLEncode。如 2017-01-01 10:30:01，编码为：2017-01-01+10%3a30%3a01。如果通过Highlight参数，设置录制为精彩视频录制，结束时间不应超过当前时间+30分钟，如果结束时间超过当前时间+30分钟或小于当前时间，则实际结束时间为当前时间+30分钟。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 录制类型。不区分大小写。
	// “video” : 音视频录制【默认】。
	// “audio” : 纯音频录制。
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 录制文件格式。不区分大小写。其值为：
	// “flv”,“hls”,”mp4”,“aac”,”mp3”，默认“flv”。
	FileFormat *string `json:"FileFormat,omitempty" name:"FileFormat"`

	// 开启精彩视频录制标志；0：不开启精彩视频录制【默认】；1：开启精彩视频录制。
	Highlight *int64 `json:"Highlight,omitempty" name:"Highlight"`

	// 开启A+B=C混流C流录制标志。0：不开启A+B=C混流C流录制【默认】；1：开启A+B=C混流C流录制。
	MixStream *int64 `json:"MixStream,omitempty" name:"MixStream"`

	// 录制流参数。当前支持以下参数：
	// record_interval - 录制分片时长，单位 秒，1800 - 7200
	// storage_time - 录制文件存储时长，单位 秒
	// eg. record_interval=3600&storage_time=7200
	// 注：参数需要url encode。
	StreamParam *string `json:"StreamParam,omitempty" name:"StreamParam"`
}

func (r *CreateLiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID，全局唯一标识录制任务。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveRecordRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveRecordRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名。非空的字符串
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// Flv录制参数，开启Flv录制时设置。
	FlvParam *RecordParam `json:"FlvParam,omitempty" name:"FlvParam"`

	// Hls录制参数，开启hls录制时设置。
	HlsParam *RecordParam `json:"HlsParam,omitempty" name:"HlsParam"`

	// Mp4录制参数，开启Mp4录制时设置。
	Mp4Param *RecordParam `json:"Mp4Param,omitempty" name:"Mp4Param"`

	// Aac录制参数，开启Aac录制时设置。
	AacParam *RecordParam `json:"AacParam,omitempty" name:"AacParam"`
}

func (r *CreateLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板Id。
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveRecordTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveSnapshotRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveSnapshotRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveSnapshotRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveSnapshotRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名称。非空的字符串
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Cos AppId。
	CosAppId *int64 `json:"CosAppId,omitempty" name:"CosAppId"`

	// Cos Bucket名称。
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// Cos地区。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 截图间隔，单位s，默认10s。
	SnapshotInterval *int64 `json:"SnapshotInterval,omitempty" name:"SnapshotInterval"`

	// 截图宽度。默认：0（原始高）
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 截图高度。默认：0（原始宽）
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 是否开启鉴黄，0：不开启，1：开启。默认：0.
	PornFlag *int64 `json:"PornFlag,omitempty" name:"PornFlag"`
}

func (r *CreateLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板Id。
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 指定已有的模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveTranscodeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveTranscodeRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveTranscodeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveTranscodeRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名称，例：900 900p 仅支持字母和数字的组合。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 视频码率。
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// 视频编码：
	// h264/h265。默认h264
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// 音频编码：
	// aac/mp3。默认原始音频格式
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// 音频码率：默认0。0-500
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// 模板描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 高，默认0。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 是否保留视频，0：否，1：是。默认1。
	NeedVideo *int64 `json:"NeedVideo,omitempty" name:"NeedVideo"`

	// 是否保留音频，0：否，1：是。默认1。
	NeedAudio *int64 `json:"NeedAudio,omitempty" name:"NeedAudio"`

	// 宽，默认0。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 帧率，默认0。
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 关键帧间隔，单位：秒。默认原始的间隔
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// 是否旋转，0：否，1：是。默认0。
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 编码质量：
	// baseline/main/high。默认baseline
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 是否不超过原始码率，0：否，1：是。默认0。
	BitrateToOrig *int64 `json:"BitrateToOrig,omitempty" name:"BitrateToOrig"`

	// 是否不超过原始高，0：否，1：是。默认0。
	HeightToOrig *int64 `json:"HeightToOrig,omitempty" name:"HeightToOrig"`

	// 是否不超过原始帧率，0：否，1：是。默认0。
	FpsToOrig *int64 `json:"FpsToOrig,omitempty" name:"FpsToOrig"`
}

func (r *CreateLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板Id。
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveWatermarkRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 水印Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveWatermarkRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveWatermarkRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLiveWatermarkRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveWatermarkRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLiveWatermarkRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePullStreamConfigRequest struct {
	*tchttp.BaseRequest

	// 源Url。
	FromUrl *string `json:"FromUrl,omitempty" name:"FromUrl"`

	// 目的Url，目前限制该目标地址为腾讯域名。
	ToUrl *string `json:"ToUrl,omitempty" name:"ToUrl"`

	// 区域id,1-深圳,2-上海，3-天津,4-香港。
	AreaId *int64 `json:"AreaId,omitempty" name:"AreaId"`

	// 运营商id,1-电信,2-移动,3-联通,4-其他,AreaId为4的时候,IspId只能为其他。
	IspId *int64 `json:"IspId,omitempty" name:"IspId"`

	// 开始时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于七天。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *CreatePullStreamConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePullStreamConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePullStreamConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置成功后的id。
		ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePullStreamConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePullStreamConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCallbackRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DeleteLiveCallbackRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCallbackRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCallbackRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveCallbackRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCallbackRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCallbackTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCallbackTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCertRequest struct {
	*tchttp.BaseRequest

	// 证书Id。
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`
}

func (r *DeleteLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordRequest struct {
	*tchttp.BaseRequest

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 任务ID，全局唯一标识录制任务。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteLiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，比如AppName为空也需要传空字符串进行强匹配
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，比如AppName为空也需要传空字符串进行强匹配
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，比如AppName为空也需要传空字符串进行强匹配
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DeleteLiveRecordRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveRecordRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveRecordTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DeleteLiveSnapshotRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveSnapshotRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveSnapshotRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveSnapshotRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。域名维度转码，域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，比如AppName为空也需要传空字符串进行强匹配
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。域名+AppName+StreamName+TemplateId唯一标识单个转码规则，如需删除需要强匹配，比如AppName为空也需要传空字符串进行强匹配
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。域名+AppName+StreamName+TemplateId唯一标识单个转码规则，如需删除需要强匹配，比如AppName为空也需要传空字符串进行强匹配
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 模板ID域名+AppName+StreamName+TemplateId唯一标识单个转码规则，如需删除需要强匹配，比如AppName为空也需要传空字符串进行强匹配
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveTranscodeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveTranscodeRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveTranscodeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveTranscodeRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// 水印ID。
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`
}

func (r *DeleteLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveWatermarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveWatermarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DeleteLiveWatermarkRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveWatermarkRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveWatermarkRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLiveWatermarkRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePullStreamConfigRequest struct {
	*tchttp.BaseRequest

	// 配置id。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeletePullStreamConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePullStreamConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePullStreamConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePullStreamConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePullStreamConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveCallbackRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则信息列表。
		Rules []*CallBackRuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCallbackRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 回调模板信息。
		Template *CallBackTemplateInfo `json:"Template,omitempty" name:"Template"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveCallbackTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板信息列表。
		Templates []*CallBackTemplateInfo `json:"Templates,omitempty" name:"Templates" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCallbackTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCallbackTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertRequest struct {
	*tchttp.BaseRequest

	// 证书Id。
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`
}

func (r *DescribeLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书信息。
		CertInfo *CertInfo `json:"CertInfo,omitempty" name:"CertInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveCertsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCertsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书信息列表。
		CertInfoSet []*CertInfo `json:"CertInfoSet,omitempty" name:"CertInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCertsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveCertsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// 播放域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书信息。
		DomainCertInfo *DomainCertInfo `json:"DomainCertInfo,omitempty" name:"DomainCertInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDomainCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveDomainCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePlayAuthKeyRequest struct {
	*tchttp.BaseRequest

	// 域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLivePlayAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLivePlayAuthKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePlayAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 播放鉴权key信息。
		PlayAuthKeyInfo *PlayAuthKeyInfo `json:"PlayAuthKeyInfo,omitempty" name:"PlayAuthKeyInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLivePlayAuthKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLivePlayAuthKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePushAuthKeyRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLivePushAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLivePushAuthKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePushAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 推流鉴权key信息。
		PushAuthKeyInfo *PushAuthKeyInfo `json:"PushAuthKeyInfo,omitempty" name:"PushAuthKeyInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLivePushAuthKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLivePushAuthKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveRecordRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则列表。
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制模板信息。
		Template *RecordTemplateInfo `json:"Template,omitempty" name:"Template"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveRecordTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制模板信息列表。
		Templates []*RecordTemplateInfo `json:"Templates,omitempty" name:"Templates" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveRecordTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveSnapshotRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则列表。
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveSnapshotRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 截图模板信息。
		Template *SnapshotTemplateInfo `json:"Template,omitempty" name:"Template"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveSnapshotTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 截图模板列表。
		Templates []*SnapshotTemplateInfo `json:"Templates,omitempty" name:"Templates" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveSnapshotTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveSnapshotTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineInfoRequest struct {
	*tchttp.BaseRequest

	// 取得第几页。
	// 默认值：1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 分页大小。
	// 最大值：100。
	// 取值范围：1~100 之前的任意整数。
	// 默认值：10。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 0:未开始推流 1:正在推流
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeLiveStreamOnlineInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamOnlineInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页的页码。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页大小。
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 符合条件的总个数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 流信息列表。
		StreamInfoList []*StreamInfo `json:"StreamInfoList,omitempty" name:"StreamInfoList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamOnlineInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamOnlineInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineListRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 取得第几页，默认1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 每页大小，最大100。 
	// 取值：1~100之前的任意整数。
	// 默认值：10。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 流名称，精确查询。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeLiveStreamOnlineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamOnlineListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的总个数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 分页的页码。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页显示的条数。
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 正在推送流的信息列表。
		OnlineInfo []*StreamOnlineInfo `json:"OnlineInfo,omitempty" name:"OnlineInfo" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamOnlineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamOnlineListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPublishedListRequest struct {
	*tchttp.BaseRequest

	// 您的域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 结束时间。
	// UTC 格式，例如：2016-06-30T19:00:00Z。
	// 不超过当前时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 起始时间。 
	// UTC 格式，例如：2016-06-29T19:00:00Z。
	// 和当前时间相隔不超过7天。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 直播流所属应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 取得第几页。
	// 默认值：1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 分页大小。
	// 最大值：100。
	// 取值范围：1~100 之前的任意整数。
	// 默认值：10。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeLiveStreamPublishedListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamPublishedListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPublishedListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 推流记录信息。
		PublishInfo []*StreamName `json:"PublishInfo,omitempty" name:"PublishInfo" list`

		// 分页的页码。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页大小
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 符合条件的总个数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamPublishedListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamPublishedListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamStateRequest struct {
	*tchttp.BaseRequest

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeLiveStreamStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamStateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流状态，
	// active：活跃，
	// inactive：非活跃，
	// forbid：禁播。
		StreamState *string `json:"StreamState,omitempty" name:"StreamState"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamStateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveTranscodeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 转码规则列表。
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板信息。
		Template *TemplateInfo `json:"Template,omitempty" name:"Template"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveTranscodeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 转码模板列表。
		Templates []*TemplateInfo `json:"Templates,omitempty" name:"Templates" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveTranscodeTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// 水印ID。
	WatermarkId *uint64 `json:"WatermarkId,omitempty" name:"WatermarkId"`
}

func (r *DescribeLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 水印信息。
		Watermark *WatermarkInfo `json:"Watermark,omitempty" name:"Watermark"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveWatermarkRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarkRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 水印规则列表。
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveWatermarkRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarkRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarksRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveWatermarksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 水印总个数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 水印信息列表。
		WatermarkList []*WatermarkInfo `json:"WatermarkList,omitempty" name:"WatermarkList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveWatermarksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveWatermarksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePullStreamConfigsRequest struct {
	*tchttp.BaseRequest

	// 配置id。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribePullStreamConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePullStreamConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePullStreamConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 拉流配置。
		PullStreamConfigs []*PullStreamConfig `json:"PullStreamConfigs,omitempty" name:"PullStreamConfigs" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePullStreamConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePullStreamConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DomainCertInfo struct {

	// 证书Id。
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// 证书名称。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间，UTC格式。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 证书内容。
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// 证书类型。
	// 0：腾讯云托管证书
	// 1：用户添加证书。
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 证书过期时间，UTC格式。
	CertExpireTime *string `json:"CertExpireTime,omitempty" name:"CertExpireTime"`

	// 使用此证书的域名名称。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 证书状态
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DropLiveStreamRequest struct {
	*tchttp.BaseRequest

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 您的加速域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DropLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DropLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DropLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DropLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DropLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveStreamRequest struct {
	*tchttp.BaseRequest

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 您的加速域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 恢复流的时间。UTC 格式，例如：2018-11-29T19:00:00Z。
	// 注意：默认禁播90天，且最长支持禁播90天。
	ResumeTime *string `json:"ResumeTime,omitempty" name:"ResumeTime"`
}

func (r *ForbidLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForbidLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 开播回调URL。
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitempty" name:"StreamBeginNotifyUrl"`

	// 断流回调URL。
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitempty" name:"StreamEndNotifyUrl"`

	// 录制回调URL。
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitempty" name:"RecordNotifyUrl"`

	// 截图回调URL。
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitempty" name:"SnapshotNotifyUrl"`

	// 鉴黄回调URL。
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitempty" name:"PornCensorshipNotifyUrl"`
}

func (r *ModifyLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveCallbackTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveCallbackTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveCertRequest struct {
	*tchttp.BaseRequest

	// 证书Id。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 证书类型。0-用户添加证书；1-腾讯云托管证书。
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// 证书名称。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 证书内容，即公钥。
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// 私钥。
	HttpsKey *string `json:"HttpsKey,omitempty" name:"HttpsKey"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// 播放域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 证书Id。
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// 状态，0：关闭  1：打开。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveDomainCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveDomainCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveDomainCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayAuthKeyRequest struct {
	*tchttp.BaseRequest

	// 域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 鉴权key。
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`

	// 有效时间，单位：秒。
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`

	// 鉴权backkey。
	AuthBackKey *string `json:"AuthBackKey,omitempty" name:"AuthBackKey"`
}

func (r *ModifyLivePlayAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePlayAuthKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLivePlayAuthKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePlayAuthKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePushAuthKeyRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 主鉴权key。
	MasterAuthKey *string `json:"MasterAuthKey,omitempty" name:"MasterAuthKey"`

	// 备鉴权key。
	BackupAuthKey *string `json:"BackupAuthKey,omitempty" name:"BackupAuthKey"`

	// 有效时间，单位：秒。
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`
}

func (r *ModifyLivePushAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePushAuthKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePushAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLivePushAuthKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLivePushAuthKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// Flv录制参数，开启Flv录制时设置。
	FlvParam *RecordParam `json:"FlvParam,omitempty" name:"FlvParam"`

	// Hls录制参数，开启hls录制时设置。
	HlsParam *RecordParam `json:"HlsParam,omitempty" name:"HlsParam"`

	// Mp4录制参数，开启Mp4录制时设置。
	Mp4Param *RecordParam `json:"Mp4Param,omitempty" name:"Mp4Param"`

	// Aac录制参数，开启Aac录制时设置。
	AacParam *RecordParam `json:"AacParam,omitempty" name:"AacParam"`
}

func (r *ModifyLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveRecordTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveRecordTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 截图时间间隔
	SnapshotInterval *int64 `json:"SnapshotInterval,omitempty" name:"SnapshotInterval"`

	// 截图宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 截图高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 是否开启鉴黄，0：不开启，1：开启。
	PornFlag *int64 `json:"PornFlag,omitempty" name:"PornFlag"`

	// Cos AppId。
	CosAppId *int64 `json:"CosAppId,omitempty" name:"CosAppId"`

	// Cos Bucket名称。
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// Cos 地域。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`
}

func (r *ModifyLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveSnapshotTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveSnapshotTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 视频编码：
	// h264/h265。
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// 音频编码：
	// aac/mp3。
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// 音频码率，默认0。0-500
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// 模板描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 视频码率。100-8000
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// 宽。0-3000
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 是否保留视频，0：否，1：是。默认1。
	NeedVideo *int64 `json:"NeedVideo,omitempty" name:"NeedVideo"`

	// 是否保留音频，0：否，1：是。默认1。
	NeedAudio *int64 `json:"NeedAudio,omitempty" name:"NeedAudio"`

	// 高。0-3000
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 帧率。0-200
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 关键帧间隔，单位：秒。0-50
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// 旋转角度。0 90 180 270
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 编码质量：
	// baseline/main/high。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 是否不超过原始码率。0：否，1：是。默认0。
	BitrateToOrig *int64 `json:"BitrateToOrig,omitempty" name:"BitrateToOrig"`

	// 是否不超过原始高。0：否，1：是。默认0。
	HeightToOrig *int64 `json:"HeightToOrig,omitempty" name:"HeightToOrig"`

	// 是否不超过原始帧率。0：否，1：是。默认0。
	FpsToOrig *int64 `json:"FpsToOrig,omitempty" name:"FpsToOrig"`
}

func (r *ModifyLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLiveTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPullStreamConfigRequest struct {
	*tchttp.BaseRequest

	// 配置id。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 源Url。
	FromUrl *string `json:"FromUrl,omitempty" name:"FromUrl"`

	// 目的Url。
	ToUrl *string `json:"ToUrl,omitempty" name:"ToUrl"`

	// 区域id,1-深圳,2-上海，3-天津,4-香港。如有改动，需同时传入IspId。
	AreaId *int64 `json:"AreaId,omitempty" name:"AreaId"`

	// 运营商id,1-电信,2-移动,3-联通,4-其他,AreaId为4的时候,IspId只能为其他。如有改动，需同时传入AreaId。
	IspId *int64 `json:"IspId,omitempty" name:"IspId"`

	// 开始时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于七天。
	// 
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ModifyPullStreamConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPullStreamConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPullStreamConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPullStreamConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPullStreamConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPullStreamStatusRequest struct {
	*tchttp.BaseRequest

	// 配置id列表。
	ConfigIds []*string `json:"ConfigIds,omitempty" name:"ConfigIds" list`

	// 目标状态。0无效，2正在运行，4暂停。
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyPullStreamStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPullStreamStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPullStreamStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPullStreamStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPullStreamStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PlayAuthKeyInfo struct {

	// 域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 鉴权key。
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`

	// 有效时间，单位：秒。
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`

	// 鉴权BackKey。
	AuthBackKey *string `json:"AuthBackKey,omitempty" name:"AuthBackKey"`
}

type PublishTime struct {

	// 推流时间
	// UTC 格式，例如：2018-06-29T19:00:00Z。
	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
}

type PullStreamConfig struct {

	// 拉流配置Id。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 源Url。
	FromUrl *string `json:"FromUrl,omitempty" name:"FromUrl"`

	// 目的Url。
	ToUrl *string `json:"ToUrl,omitempty" name:"ToUrl"`

	// 区域名。
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`

	// 运营商名。
	IspName *string `json:"IspName,omitempty" name:"IspName"`

	// 开始时间。
	// UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	// 
	// UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 0无效，1初始状态，2正在运行，3拉起失败，4暂停。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type PushAuthKeyInfo struct {

	// 域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 主鉴权key。
	MasterAuthKey *string `json:"MasterAuthKey,omitempty" name:"MasterAuthKey"`

	// 备鉴权key。
	BackupAuthKey *string `json:"BackupAuthKey,omitempty" name:"BackupAuthKey"`

	// 有效时间，单位：秒。
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`
}

type RecordParam struct {

	// 录制间隔。
	// 单位秒，默认值1800。
	// 取值范围:300-7200。
	// 此参数对 HLS 无效，当录制 HLS 时从推流到断流生成一个文件。
	RecordInterval *int64 `json:"RecordInterval,omitempty" name:"RecordInterval"`

	// 录制存储时长。
	// 单位秒，取值范围： 0-5184000。
	// 0表示永久存储。
	StorageTime *int64 `json:"StorageTime,omitempty" name:"StorageTime"`

	// 是否开启当前格式录制，0 否 1是。默认值0。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

type RecordTemplateInfo struct {

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// Flv录制参数。
	FlvParam *RecordParam `json:"FlvParam,omitempty" name:"FlvParam"`

	// Hls录制参数。
	HlsParam *RecordParam `json:"HlsParam,omitempty" name:"HlsParam"`

	// Mp4录制参数。
	Mp4Param *RecordParam `json:"Mp4Param,omitempty" name:"Mp4Param"`

	// Aac录制参数。
	AacParam *RecordParam `json:"AacParam,omitempty" name:"AacParam"`
}

type ResumeDelayLiveStreamRequest struct {
	*tchttp.BaseRequest

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 您的加速域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *ResumeDelayLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeDelayLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeDelayLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeDelayLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeDelayLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeLiveStreamRequest struct {
	*tchttp.BaseRequest

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 您的加速域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *ResumeLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RuleInfo struct {

	// 规则创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 规则更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

type SetLiveWatermarkStatusRequest struct {
	*tchttp.BaseRequest

	// 水印ID。
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`

	// 状态。0：停用，1:启用
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *SetLiveWatermarkStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetLiveWatermarkStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetLiveWatermarkStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetLiveWatermarkStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetLiveWatermarkStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SnapshotTemplateInfo struct {

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 截图时间间隔。5-300
	SnapshotInterval *int64 `json:"SnapshotInterval,omitempty" name:"SnapshotInterval"`

	// 截图宽度。0-3000 0原始宽度并适配原始比例
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 截图高度。0-2000 0原始高度并适配原始比例
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 是否开启鉴黄，0：不开启，1：开启。
	PornFlag *int64 `json:"PornFlag,omitempty" name:"PornFlag"`

	// Cos AppId。
	CosAppId *int64 `json:"CosAppId,omitempty" name:"CosAppId"`

	// Cos Bucket名称。
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// Cos 地域。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 模板描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type StopLiveRecordRequest struct {
	*tchttp.BaseRequest

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 任务ID，全局唯一标识录制任务。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopLiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopLiveRecordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopLiveRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopLiveRecordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StreamInfo struct {

	// 直播流所属应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 创建模式
	CreateMode *string `json:"CreateMode,omitempty" name:"CreateMode"`

	// 创建时间，如: 2018-07-13 14:48:23
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 流状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 流id
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// 流名称
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 水印id
	WaterMarkId *string `json:"WaterMarkId,omitempty" name:"WaterMarkId"`
}

type StreamName struct {

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

type StreamOnlineInfo struct {

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 推流时间列表
	PublishTimeList []*PublishTime `json:"PublishTimeList,omitempty" name:"PublishTimeList" list`

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

type TemplateInfo struct {

	// 视频编码：
	// h264/h265。
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// 视频码率。100-8000kbps
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// 音频编码：aac/mp3
	// aac/mp3。
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// 音频码率。0-500
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// 宽。0-3000
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高。0-3000
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 帧率。0-200
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 关键帧间隔，单位：秒。1-50
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// 旋转角度。0 90 180 270
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 编码质量：
	// baseline/main/high。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 是否不超过原始码率。0：否，1：是。
	BitrateToOrig *int64 `json:"BitrateToOrig,omitempty" name:"BitrateToOrig"`

	// 是否不超过原始高度。0：否，1：是。
	HeightToOrig *int64 `json:"HeightToOrig,omitempty" name:"HeightToOrig"`

	// 是否不超过原始帧率。0：否，1：是。
	FpsToOrig *int64 `json:"FpsToOrig,omitempty" name:"FpsToOrig"`

	// 是否保留视频。0：否，1：是。
	NeedVideo *int64 `json:"NeedVideo,omitempty" name:"NeedVideo"`

	// 是否保留音频。0：否，1：是。
	NeedAudio *int64 `json:"NeedAudio,omitempty" name:"NeedAudio"`

	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模板描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type UnBindLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// 播放域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *UnBindLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindLiveDomainCertRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindLiveDomainCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindLiveDomainCertResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// 水印ID。
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`

	// 水印图片url。
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// 显示位置，X轴偏移。
	XPosition *int64 `json:"XPosition,omitempty" name:"XPosition"`

	// 显示位置，Y轴偏移。
	YPosition *int64 `json:"YPosition,omitempty" name:"YPosition"`

	// 水印名称。
	WatermarkName *string `json:"WatermarkName,omitempty" name:"WatermarkName"`
}

func (r *UpdateLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateLiveWatermarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLiveWatermarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateLiveWatermarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WatermarkInfo struct {

	// 水印ID。
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`

	// 水印图片url。
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// 显示位置，X轴偏移。
	XPosition *int64 `json:"XPosition,omitempty" name:"XPosition"`

	// 显示位置，Y轴偏移。
	YPosition *int64 `json:"YPosition,omitempty" name:"YPosition"`

	// 水印名称。
	WatermarkName *string `json:"WatermarkName,omitempty" name:"WatermarkName"`

	// 当前状态。0：未使用，1:使用中。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 添加时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}
