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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddDelayLiveStreamRequest struct {
	*tchttp.BaseRequest

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 延播时间，单位：秒，上限：600秒。
	DelayTime *uint64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// 延播设置的过期时间。UTC 格式，例如：2018-11-29T19:00:00Z。
	// 注意：
	// 1. 默认7天后过期，且最长支持7天内生效。
	// 2. 北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *AddDelayLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDelayLiveStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "DomainName")
	delete(f, "StreamName")
	delete(f, "DelayTime")
	delete(f, "ExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddDelayLiveStreamRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDelayLiveStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddLiveDomainRequest struct {
	*tchttp.BaseRequest

	// 域名名称。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 域名类型，
	// 0：推流域名，
	// 1：播放域名。
	DomainType *uint64 `json:"DomainType,omitempty" name:"DomainType"`

	// 拉流域名类型：
	// 1：国内，
	// 2：全球，
	// 3：境外。
	// 默认值：1。
	PlayType *uint64 `json:"PlayType,omitempty" name:"PlayType"`

	// 是否是慢直播：
	// 0： 普通直播，
	// 1 ：慢直播 。
	// 默认值： 0。
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// 是否是小程序直播：
	// 0： 标准直播，
	// 1 ：小程序直播 。
	// 默认值： 0。
	IsMiniProgramLive *int64 `json:"IsMiniProgramLive,omitempty" name:"IsMiniProgramLive"`
}

func (r *AddLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLiveDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "DomainType")
	delete(f, "PlayType")
	delete(f, "IsDelayLive")
	delete(f, "IsMiniProgramLive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddLiveDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLiveDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// 水印图片 URL。
	// URL中禁止包含的字符：
	//  ;(){}$>`#"\'|
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// 水印名称。
	// 最长16字节。
	WatermarkName *string `json:"WatermarkName,omitempty" name:"WatermarkName"`

	// 显示位置，X轴偏移，单位是百分比，默认 0。
	XPosition *int64 `json:"XPosition,omitempty" name:"XPosition"`

	// 显示位置，Y轴偏移，单位是百分比，默认 0。
	YPosition *int64 `json:"YPosition,omitempty" name:"YPosition"`

	// 水印宽度，占直播原始画面宽度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 水印高度，占直播原始画面高度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

func (r *AddLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLiveWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PictureUrl")
	delete(f, "WatermarkName")
	delete(f, "XPosition")
	delete(f, "YPosition")
	delete(f, "Width")
	delete(f, "Height")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddLiveWatermarkRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLiveWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BandwidthInfo struct {

	// 返回格式：
	// yyyy-mm-dd HH:MM:SS
	// 根据粒度会有不同程度的缩减。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 带宽。
	Bandwidth *float64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

type BillAreaInfo struct {

	// 大区名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 国家明细数据
	Countrys []*BillCountryInfo `json:"Countrys,omitempty" name:"Countrys"`
}

type BillCountryInfo struct {

	// 国家名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 带宽明细数据信息。
	BandInfoList []*BillDataInfo `json:"BandInfoList,omitempty" name:"BandInfoList"`
}

type BillDataInfo struct {

	// 时间点，格式: yyyy-mm-dd HH:MM:SS。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 带宽，单位是 Mbps。
	Bandwidth *float64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 流量，单位是 MB。
	Flux *float64 `json:"Flux,omitempty" name:"Flux"`

	// 峰值时间点，格式: yyyy-mm-dd HH:MM:SS，原始数据为5分钟粒度，如果查询小时和天粒度数据，则返回对应粒度内的带宽峰值时间点。
	PeakTime *string `json:"PeakTime,omitempty" name:"PeakTime"`
}

type BindLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// 证书Id。使用添加证书接口获取证书Id。
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// 播放域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// HTTPS开启状态，0： 关闭  1：打开。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *BindLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindLiveDomainCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertId")
	delete(f, "DomainName")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindLiveDomainCertRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindLiveDomainCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallBackRuleInfo struct {

	// 规则创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 规则更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

type CallBackTemplateInfo struct {

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 开播回调 URL。
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitempty" name:"StreamBeginNotifyUrl"`

	// 混流回调 URL。(参数已弃用)。
	StreamMixNotifyUrl *string `json:"StreamMixNotifyUrl,omitempty" name:"StreamMixNotifyUrl"`

	// 断流回调 URL。
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitempty" name:"StreamEndNotifyUrl"`

	// 录制回调 URL。
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitempty" name:"RecordNotifyUrl"`

	// 截图回调 URL。
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitempty" name:"SnapshotNotifyUrl"`

	// 鉴黄回调 URL。
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitempty" name:"PornCensorshipNotifyUrl"`

	// 回调的鉴权 key。
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

type CallbackEventInfo struct {

	// 事件时间
	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`

	// 事件类型
	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`

	// 回调请求
	Request *string `json:"Request,omitempty" name:"Request"`

	// 回调响应
	Response *string `json:"Response,omitempty" name:"Response"`

	// 客户接口响应时间
	ResponseTime *string `json:"ResponseTime,omitempty" name:"ResponseTime"`

	// 回调结果
	ResultCode *uint64 `json:"ResultCode,omitempty" name:"ResultCode"`

	// 流名称
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`
}

type CancelCommonMixStreamRequest struct {
	*tchttp.BaseRequest

	// 混流会话（申请混流开始到取消混流结束）标识 ID。
	// 该值与CreateCommonMixStream中的MixStreamSessionId保持一致。
	MixStreamSessionId *string `json:"MixStreamSessionId,omitempty" name:"MixStreamSessionId"`
}

func (r *CancelCommonMixStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelCommonMixStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MixStreamSessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelCommonMixStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelCommonMixStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelCommonMixStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelCommonMixStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CdnPlayStatData struct {

	// 时间点，格式: yyyy-mm-dd HH:MM:SS。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 带宽，单位: Mbps。
	Bandwidth *float64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 流量，单位: MB。
	Flux *float64 `json:"Flux,omitempty" name:"Flux"`

	// 新增请求数。
	Request *uint64 `json:"Request,omitempty" name:"Request"`

	// 并发连接数。
	Online *uint64 `json:"Online,omitempty" name:"Online"`
}

type CertInfo struct {

	// 证书 ID。
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`

	// 证书名称。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间，UTC 格式。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 证书内容。
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// 证书类型。
	// 0：用户添加证书，
	// 1：腾讯云托管证书。
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 证书过期时间，UTC 格式。
	CertExpireTime *string `json:"CertExpireTime,omitempty" name:"CertExpireTime"`

	// 使用此证书的域名列表。
	DomainList []*string `json:"DomainList,omitempty" name:"DomainList"`
}

type ClientIpPlaySumInfo struct {

	// 客户端 IP，点分型。
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 客户端所在省份。
	Province *string `json:"Province,omitempty" name:"Province"`

	// 总流量。
	TotalFlux *float64 `json:"TotalFlux,omitempty" name:"TotalFlux"`

	// 总请求数。
	TotalRequest *uint64 `json:"TotalRequest,omitempty" name:"TotalRequest"`

	// 总失败请求数。
	TotalFailedRequest *uint64 `json:"TotalFailedRequest,omitempty" name:"TotalFailedRequest"`

	// 客户端所在国家。
	CountryArea *string `json:"CountryArea,omitempty" name:"CountryArea"`
}

type CommonMixControlParams struct {

	// 取值范围[0,1]。
	// 填1时，当参数中图层分辨率参数与视频实际分辨率不一致时，自动从视频中按图层设置的分辨率比例进行裁剪。
	UseMixCropCenter *int64 `json:"UseMixCropCenter,omitempty" name:"UseMixCropCenter"`

	// 取值范围[0,1]
	// 填1时，当InputStreamList中个数为1时，且OutputParams.OutputStreamType为1时，不执行取消操作，执行拷贝流操作
	AllowCopy *int64 `json:"AllowCopy,omitempty" name:"AllowCopy"`

	// 取值范围[0,1]
	// 填1时，透传原始流的sei
	PassInputSei *int64 `json:"PassInputSei,omitempty" name:"PassInputSei"`
}

type CommonMixCropParams struct {

	// 裁剪的宽度。取值范围[0，2000]。
	CropWidth *float64 `json:"CropWidth,omitempty" name:"CropWidth"`

	// 裁剪的高度。取值范围[0，2000]。
	CropHeight *float64 `json:"CropHeight,omitempty" name:"CropHeight"`

	// 裁剪的起始X坐标。取值范围[0，2000]。
	CropStartLocationX *float64 `json:"CropStartLocationX,omitempty" name:"CropStartLocationX"`

	// 裁剪的起始Y坐标。取值范围[0，2000]。
	CropStartLocationY *float64 `json:"CropStartLocationY,omitempty" name:"CropStartLocationY"`
}

type CommonMixInputParam struct {

	// 输入流名称。80字节以内，仅含字母、数字以及下划线的字符串。
	// 当LayoutParams.InputType=0(音视频)/4(纯音频)/5(纯视频)时，该值为需要混流的流名称。
	// 当LayoutParams.InputType=2(图片)/3(画布)时，该值仅用作标识输入，可用类似Canvas1、Pictrue1的名称。
	InputStreamName *string `json:"InputStreamName,omitempty" name:"InputStreamName"`

	// 输入流布局参数。
	LayoutParams *CommonMixLayoutParams `json:"LayoutParams,omitempty" name:"LayoutParams"`

	// 输入流裁剪参数。
	CropParams *CommonMixCropParams `json:"CropParams,omitempty" name:"CropParams"`
}

type CommonMixLayoutParams struct {

	// 输入图层。取值范围[1，16]。
	// 1)背景流（即大主播画面或画布）的 image_layer 填1。
	// 2)纯音频混流，该参数也需填。
	// 注意：不同输入，该值不可重复
	ImageLayer *int64 `json:"ImageLayer,omitempty" name:"ImageLayer"`

	// 输入类型。取值范围[0，5]。
	// 不填默认为0。
	// 0表示输入流为音视频。
	// 2表示输入流为图片。
	// 3表示输入流为画布。 
	// 4表示输入流为音频。
	// 5表示输入流为纯视频。
	InputType *int64 `json:"InputType,omitempty" name:"InputType"`

	// 输入画面在输出时的高度。取值范围：
	// 像素：[0，2000]
	// 百分比：[0.01，0.99]
	// 不填默认为输入流的高度。
	// 使用百分比时，期望输出为（百分比 * 背景高）。
	ImageHeight *float64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

	// 输入画面在输出时的宽度。取值范围：
	// 像素：[0，2000]
	// 百分比：[0.01，0.99]
	// 不填默认为输入流的宽度。
	// 使用百分比时，期望输出为（百分比 * 背景宽）。
	ImageWidth *float64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

	// 输入在输出画面的X偏移。取值范围：
	// 像素：[0，2000]
	// 百分比：[0.01，0.99]
	// 不填默认为0。
	// 相对于大主播背景画面左上角的横向偏移。 
	// 使用百分比时，期望输出为（百分比 * 背景宽）。
	LocationX *float64 `json:"LocationX,omitempty" name:"LocationX"`

	// 输入在输出画面的Y偏移。取值范围：
	// 像素：[0，2000]
	// 百分比：[0.01，0.99]
	// 不填默认为0。
	// 相对于大主播背景画面左上角的纵向偏移。 
	// 使用百分比时，期望输出为（百分比 * 背景宽）
	LocationY *float64 `json:"LocationY,omitempty" name:"LocationY"`

	// 当InputType为3(画布)时，该值表示画布的颜色。
	// 常用的颜色有：
	// 红色：0xcc0033。
	// 黄色：0xcc9900。
	// 绿色：0xcccc33。
	// 蓝色：0x99CCFF。
	// 黑色：0x000000。
	// 白色：0xFFFFFF。
	// 灰色：0x999999。
	Color *string `json:"Color,omitempty" name:"Color"`

	// 当InputType为2(图片)时，该值是水印ID。
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`
}

type CommonMixOutputParams struct {

	// 输出流名称。
	OutputStreamName *string `json:"OutputStreamName,omitempty" name:"OutputStreamName"`

	// 输出流类型，取值范围[0,1]。
	// 不填默认为0。
	// 当输出流为输入流 list 中的一条时，填写0。
	// 当期望生成的混流结果成为一条新流时，该值填为1。
	// 该值为1时，output_stream_id 不能出现在 input_stram_list 中，且直播后台中，不能存在相同 ID 的流。
	OutputStreamType *int64 `json:"OutputStreamType,omitempty" name:"OutputStreamType"`

	// 输出流比特率。取值范围[1，50000]。
	// 不填的情况下，系统会自动判断。
	OutputStreamBitRate *int64 `json:"OutputStreamBitRate,omitempty" name:"OutputStreamBitRate"`

	// 输出流GOP大小。取值范围[1,10]。
	// 不填的情况下，系统会自动判断。
	OutputStreamGop *int64 `json:"OutputStreamGop,omitempty" name:"OutputStreamGop"`

	// 输出流帧率大小。取值范围[1,60]。
	// 不填的情况下，系统会自动判断。
	OutputStreamFrameRate *int64 `json:"OutputStreamFrameRate,omitempty" name:"OutputStreamFrameRate"`

	// 输出流音频比特率。取值范围[1,500]
	// 不填的情况下，系统会自动判断。
	OutputAudioBitRate *int64 `json:"OutputAudioBitRate,omitempty" name:"OutputAudioBitRate"`

	// 输出流音频采样率。取值范围[96000, 88200, 64000, 48000, 44100, 32000,24000, 22050, 16000, 12000, 11025, 8000]。
	// 不填的情况下，系统会自动判断。
	OutputAudioSampleRate *int64 `json:"OutputAudioSampleRate,omitempty" name:"OutputAudioSampleRate"`

	// 输出流音频声道数。取值范围[1,2]。
	// 不填的情况下，系统会自动判断。
	OutputAudioChannels *int64 `json:"OutputAudioChannels,omitempty" name:"OutputAudioChannels"`

	// 输出流中的sei信息。如果无特殊需要，不填。
	MixSei *string `json:"MixSei,omitempty" name:"MixSei"`
}

type ConcurrentRecordStreamNum struct {

	// 时间点。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 路数。
	Num *uint64 `json:"Num,omitempty" name:"Num"`
}

type CreateCommonMixStreamRequest struct {
	*tchttp.BaseRequest

	// 混流会话（申请混流开始到取消混流结束）标识 ID。80字节以内，仅含字母、数字以及下划线的字符串。
	MixStreamSessionId *string `json:"MixStreamSessionId,omitempty" name:"MixStreamSessionId"`

	// 混流输入流列表。
	InputStreamList []*CommonMixInputParam `json:"InputStreamList,omitempty" name:"InputStreamList"`

	// 混流输出流参数。
	OutputParams *CommonMixOutputParams `json:"OutputParams,omitempty" name:"OutputParams"`

	// 输入模板 ID，若设置该参数，将按默认模板布局输出，无需填入自定义位置参数。
	// 不填默认为0。
	// 两输入源支持10，20，30，40，50。
	// 三输入源支持310，390，391。
	// 四输入源支持410。
	// 五输入源支持510，590。
	// 六输入源支持610。
	MixStreamTemplateId *int64 `json:"MixStreamTemplateId,omitempty" name:"MixStreamTemplateId"`

	// 混流的特殊控制参数。如无特殊需求，无需填写。
	ControlParams *CommonMixControlParams `json:"ControlParams,omitempty" name:"ControlParams"`
}

func (r *CreateCommonMixStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCommonMixStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MixStreamSessionId")
	delete(f, "InputStreamList")
	delete(f, "OutputParams")
	delete(f, "MixStreamTemplateId")
	delete(f, "ControlParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCommonMixStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCommonMixStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCommonMixStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCommonMixStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 模板ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveCallbackRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveCallbackRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveCallbackRuleRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveCallbackRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名称。
	// 长度上限：255字节。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	// 长度上限：1024字节。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 开播回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitempty" name:"StreamBeginNotifyUrl"`

	// 断流回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitempty" name:"StreamEndNotifyUrl"`

	// 录制回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitempty" name:"RecordNotifyUrl"`

	// 截图回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitempty" name:"SnapshotNotifyUrl"`

	// 鉴黄回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32741)。
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitempty" name:"PornCensorshipNotifyUrl"`

	// 回调 Key，回调 URL 公用，回调签名详见事件消息通知文档。
	// [事件消息通知](/document/product/267/32744)。
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`

	// 参数已弃用。
	StreamMixNotifyUrl *string `json:"StreamMixNotifyUrl,omitempty" name:"StreamMixNotifyUrl"`
}

func (r *CreateLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveCallbackTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "Description")
	delete(f, "StreamBeginNotifyUrl")
	delete(f, "StreamEndNotifyUrl")
	delete(f, "RecordNotifyUrl")
	delete(f, "SnapshotNotifyUrl")
	delete(f, "PornCensorshipNotifyUrl")
	delete(f, "CallbackKey")
	delete(f, "StreamMixNotifyUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveCallbackTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板ID。
		TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveCallbackTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveCallbackTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveCertRequest struct {
	*tchttp.BaseRequest

	// 证书类型。0-用户添加证书；1-腾讯云托管证书。
	// 注意：当证书类型为0时，HttpsCrt和HttpsKey必选；
	// 当证书类型为1时，优先使用CloudCertId对应证书，若CloudCertId为空则使用HttpsCrt和HttpsKey。
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// 证书名称。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 证书内容，即公钥。
	HttpsCrt *string `json:"HttpsCrt,omitempty" name:"HttpsCrt"`

	// 私钥。
	HttpsKey *string `json:"HttpsKey,omitempty" name:"HttpsKey"`

	// 描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 腾讯云证书托管ID。
	CloudCertId *string `json:"CloudCertId,omitempty" name:"CloudCertId"`
}

func (r *CreateLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertType")
	delete(f, "CertName")
	delete(f, "HttpsCrt")
	delete(f, "HttpsKey")
	delete(f, "Description")
	delete(f, "CloudCertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveCertRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLivePullStreamTaskRequest struct {
	*tchttp.BaseRequest

	// 拉流源的类型：
	// PullLivePushLive -直播，
	// PullVodPushLive -点播。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 拉流源 url 列表。
	// SourceType 为直播（PullLivePushLive）只可以填1个，
	// SourceType 为点播（PullVodPushLive）可以填多个，上限30个。
	// 当前支持的文件格式：flv，mp4，hls。
	// 当前支持的拉流协议：http，https，rtmp。
	// 注意：
	// 1. 建议优先使用 flv 文件，对于 mp4 未交织好的文件轮播推流易产生卡顿，可通过点播转码进行重新交织后再轮播。
	// 2. 拒绝内网域名等攻击性拉流地址，如有使用，则做账号封禁处理。
	// 3. 源文件请保持时间戳正常交织递增，避免因源文件异常影响推流及播放。
	// 4. 视频编码格式仅支持: H264, H265。
	// 5. 音频编码格式仅支持: AAC。
	// 6. 点播源请使用小文件，尽量时长保持在1小时内，较大文件打开和续播耗时较久，耗时超过15秒会有无法正常转推风险。
	SourceUrls []*string `json:"SourceUrls,omitempty" name:"SourceUrls"`

	// 推流域名。
	// 将拉取过来的流推到该域名。
	// 注意：请使用已在云直播配置的推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	// 将拉取过来的流推到该路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 推流名称。
	// 将拉取过来的流推到该流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 开始时间。
	// 使用 UTC 格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于七天。
	// 使用 UTC 格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务操作人备注。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 推流参数。
	// 推流时携带自定义参数。
	// 示例：
	// bak=1&test=2 。
	PushArgs *string `json:"PushArgs,omitempty" name:"PushArgs"`

	// 选择需要回调的事件（不填则回调全部）：
	// TaskStart：任务启动回调，
	// TaskExit：任务停止回调，
	// VodSourceFileStart：从点播源文件开始拉流回调，
	// VodSourceFileFinish：从点播源文件拉流结束回调，
	// ResetTaskConfig：任务更新回调。
	// 
	// TaskAlarm: 用于告警事件通知，AlarmType 示例:
	// PullFileUnstable - 文件拉取不稳定，
	// PushStreamUnstable - 推流不稳定，
	// PullFileFailed - 文件拉取出错，
	// PushStreamFailed - 推流出现失败，
	// FileEndEarly - 文件提前结束。
	CallbackEvents []*string `json:"CallbackEvents,omitempty" name:"CallbackEvents"`

	// 点播拉流转推循环次数。默认：-1。
	// -1：无限循环，直到任务结束。
	// 0：不循环。
	// >0：具体循环次数。次数和时间以先结束的为准。
	// 注意：该配置仅对拉流源为点播时生效。
	VodLoopTimes *string `json:"VodLoopTimes,omitempty" name:"VodLoopTimes"`

	// 点播更新SourceUrls后的播放方式：
	// ImmediateNewSource：立即播放新的拉流源内容；
	// ContinueBreakPoint：播放完当前正在播放的点播 url 后再使用新的拉流源播放。（旧拉流源未播放的点播 url 不会再播放）
	// 
	// 注意：该配置生效仅对变更前拉流源为点播时生效。
	VodRefreshType *string `json:"VodRefreshType,omitempty" name:"VodRefreshType"`

	// 自定义回调地址。
	// 拉流转推任务相关事件会回调到该地址。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 其他参数。
	// 示例: ignore_region  用于忽略传入地域, 内部按负载分配。
	ExtraCmd *string `json:"ExtraCmd,omitempty" name:"ExtraCmd"`

	// 任务描述，限制 512 字节。
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *CreateLivePullStreamTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLivePullStreamTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceType")
	delete(f, "SourceUrls")
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "StreamName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Operator")
	delete(f, "PushArgs")
	delete(f, "CallbackEvents")
	delete(f, "VodLoopTimes")
	delete(f, "VodRefreshType")
	delete(f, "CallbackUrl")
	delete(f, "ExtraCmd")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLivePullStreamTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLivePullStreamTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 Id 。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLivePullStreamTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLivePullStreamTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordRequest struct {
	*tchttp.BaseRequest

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 推流路径，与推流和播放地址中的 AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 推流域名。多域名推流必须设置。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 录制开始时间。中国标准时间，需要 URLEncode(rfc3986)。如 2017-01-01 10:10:01，编码为：2017-01-01+10%3a10%3a01。
	// 定时录制模式，必须设置该字段；实时视频录制模式，忽略该字段。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 录制结束时间。中国标准时间，需要 URLEncode(rfc3986)。如 2017-01-01 10:30:01，编码为：2017-01-01+10%3a30%3a01。
	// 定时录制模式，必须设置该字段；实时录制模式，为可选字段。如果通过Highlight参数，设置录制为实时视频录制模式，其设置的结束时间不应超过当前时间+30分钟，如果设置的结束时间超过当前时间+30分钟或者小于当前时间或者不设置该参数，则实际结束时间为当前时间+30分钟。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 录制类型。
	// “video” : 音视频录制【默认】。
	// “audio” : 纯音频录制。
	// 在定时录制模式或实时视频录制模式下，该参数均有效，不区分大小写。
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 录制文件格式。其值为：
	// “flv”【默认】,“hls”,”mp4”,“aac”,”mp3”。
	// 在定时录制模式或实时视频录制模式下，该参数均有效，不区分大小写。
	FileFormat *string `json:"FileFormat,omitempty" name:"FileFormat"`

	// 开启实时视频录制模式标志。
	// 0：不开启实时视频录制模式，即定时录制模式【默认】。见[示例一](#.E7.A4.BA.E4.BE.8B1-.E5.88.9B.E5.BB.BA.E5.AE.9A.E6.97.B6.E5.BD.95.E5.88.B6.E4.BB.BB.E5.8A.A1)。
	// 1：开启实时视频录制模式。见[示例二](#.E7.A4.BA.E4.BE.8B2-.E5.88.9B.E5.BB.BA.E5.AE.9E.E6.97.B6.E5.BD.95.E5.88.B6.E4.BB.BB.E5.8A.A1)。
	Highlight *int64 `json:"Highlight,omitempty" name:"Highlight"`

	// 开启 A+B=C混流C流录制标志。
	// 0：不开启 A+B=C混流C流录制【默认】。
	// 1：开启 A+B=C混流C流录制。
	// 在定时录制模式或实时视频录制模式下，该参数均有效。
	MixStream *int64 `json:"MixStream,omitempty" name:"MixStream"`

	// 录制流参数。当前支持以下参数：
	// record_interval - 录制分片时长，单位 秒，1800 - 7200。
	// storage_time - 录制文件存储时长，单位 秒。
	// eg. record_interval=3600&storage_time=2592000。
	// 注：参数需要url encode。
	// 在定时录制模式或实时视频录制模式下，该参数均有效。
	StreamParam *string `json:"StreamParam,omitempty" name:"StreamParam"`
}

func (r *CreateLiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StreamName")
	delete(f, "AppName")
	delete(f, "DomainName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RecordType")
	delete(f, "FileFormat")
	delete(f, "Highlight")
	delete(f, "MixStream")
	delete(f, "StreamParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 ID，全局唯一标识录制任务。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	// 注：如果本参数设置为非空字符串，规则将只对此推流起作用。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *CreateLiveRecordRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveRecordRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "TemplateId")
	delete(f, "AppName")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveRecordRuleRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveRecordRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名。仅支持中文、英文、数字、_、-。
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

	// 直播类型，默认 0。
	// 0：普通直播，
	// 1：慢直播。
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// HLS专属录制参数。
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitempty" name:"HlsSpecialParam"`

	// Mp3录制参数，开启Mp3录制时设置。
	Mp3Param *RecordParam `json:"Mp3Param,omitempty" name:"Mp3Param"`
}

func (r *CreateLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "Description")
	delete(f, "FlvParam")
	delete(f, "HlsParam")
	delete(f, "Mp4Param")
	delete(f, "AacParam")
	delete(f, "IsDelayLive")
	delete(f, "HlsSpecialParam")
	delete(f, "Mp3Param")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveRecordTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	// 注：如果本参数设置为非空字符串，规则将只对此推流起作用。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *CreateLiveSnapshotRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveSnapshotRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "TemplateId")
	delete(f, "AppName")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveSnapshotRuleRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveSnapshotRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名称。
	// 长度上限：255字节。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Cos 应用 ID。
	CosAppId *int64 `json:"CosAppId,omitempty" name:"CosAppId"`

	// Cos Bucket名称。
	// 注：CosBucket参数值不能包含-[appid] 部分。
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// Cos地区。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 描述信息。
	// 长度上限：1024字节。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 截图间隔，单位s，默认10s。
	// 范围： 2s ~ 300s。
	SnapshotInterval *int64 `json:"SnapshotInterval,omitempty" name:"SnapshotInterval"`

	// 截图宽度。默认：0（原始宽）。
	// 范围：0-3000 。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 截图高度。默认：0（原始高）。
	// 范围：0-2000 。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 是否开启鉴黄，0：不开启，1：开启。默认：0。
	PornFlag *int64 `json:"PornFlag,omitempty" name:"PornFlag"`

	// Cos Bucket文件夹前缀。
	// 如不传，实际按默认值
	// /{Year}-{Month}-{Day}
	// 生效
	CosPrefix *string `json:"CosPrefix,omitempty" name:"CosPrefix"`

	// Cos 文件名称。
	// 如不传，实际按默认值
	// {StreamID}-screenshot-{Hour}-{Minute}-{Second}-{Width}x{Height}{Ext}
	// 生效
	CosFileName *string `json:"CosFileName,omitempty" name:"CosFileName"`
}

func (r *CreateLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "CosAppId")
	delete(f, "CosBucket")
	delete(f, "CosRegion")
	delete(f, "Description")
	delete(f, "SnapshotInterval")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "PornFlag")
	delete(f, "CosPrefix")
	delete(f, "CosFileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveSnapshotTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeRuleRequest struct {
	*tchttp.BaseRequest

	// 播放域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致。如果只绑定域名，则此处填空。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。如果只绑定域名或路径，则此处填空。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 指定已有的模板Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveTranscodeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveTranscodeRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "StreamName")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveTranscodeRuleRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveTranscodeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名称，例： 900p 仅支持字母和数字的组合。
	// 长度限制：
	//   标准转码：1-10个字符
	//   极速高清转码：3-10个字符
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 视频码率。范围：0kbps - 8000kbps。
	// 0为保持原始码率。
	// 注: 转码模板有码率唯一要求，最终保存的码率可能与输入码率有所差别。
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// 音频编码：aac，默认aac。
	// 注意：当前该参数未生效，待后续支持！
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// 音频码率，默认0。
	// 范围：0-500。
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// 视频编码：h264/h265/origin，默认origin。
	// 
	// origin: 保持原始编码格式
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// 模板描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否保留视频，0：否，1：是。默认1。
	NeedVideo *int64 `json:"NeedVideo,omitempty" name:"NeedVideo"`

	// 宽，默认0。
	// 范围[0-3000]
	// 数值必须是2的倍数，0是原始宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 是否保留音频，0：否，1：是。默认1。
	NeedAudio *int64 `json:"NeedAudio,omitempty" name:"NeedAudio"`

	// 高，默认0。
	// 范围[0-3000]
	// 数值必须是2的倍数，0是原始高度。
	// 极速高清模板（AiTransCode = 1 的时候）必须传。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 帧率，默认0。
	// 范围0-60fps
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 关键帧间隔，单位：秒。
	// 默认原始的间隔
	// 范围2-6
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// 旋转角度，默认0。
	// 可取值：0，90，180，270
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 编码质量：
	// baseline/main/high。默认baseline
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 当设置的码率>原始码率时，是否以原始码率为准。
	// 0：否， 1：是
	// 默认 0。
	BitrateToOrig *int64 `json:"BitrateToOrig,omitempty" name:"BitrateToOrig"`

	// 当设置的高度>原始高度时，是否以原始高度为准。
	// 0：否， 1：是
	// 默认 0。
	HeightToOrig *int64 `json:"HeightToOrig,omitempty" name:"HeightToOrig"`

	// 当设置的帧率>原始帧率时，是否以原始帧率为准。
	// 0：否， 1：是
	// 默认 0。
	FpsToOrig *int64 `json:"FpsToOrig,omitempty" name:"FpsToOrig"`

	// 是否是极速高清模板，0：否，1：是。默认0。
	AiTransCode *int64 `json:"AiTransCode,omitempty" name:"AiTransCode"`

	// 极速高清视频码率压缩比。
	// 极速高清目标码率=VideoBitrate * (1-AdaptBitratePercent)
	// 
	// 取值范围：0.0到0.5
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitempty" name:"AdaptBitratePercent"`

	// 是否以短边作为高度，0：否，1：是。默认0。
	ShortEdgeAsHeight *int64 `json:"ShortEdgeAsHeight,omitempty" name:"ShortEdgeAsHeight"`
}

func (r *CreateLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "VideoBitrate")
	delete(f, "Acodec")
	delete(f, "AudioBitrate")
	delete(f, "Vcodec")
	delete(f, "Description")
	delete(f, "NeedVideo")
	delete(f, "Width")
	delete(f, "NeedAudio")
	delete(f, "Height")
	delete(f, "Fps")
	delete(f, "Gop")
	delete(f, "Rotate")
	delete(f, "Profile")
	delete(f, "BitrateToOrig")
	delete(f, "HeightToOrig")
	delete(f, "FpsToOrig")
	delete(f, "AiTransCode")
	delete(f, "AdaptBitratePercent")
	delete(f, "ShortEdgeAsHeight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveTranscodeTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveTranscodeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveWatermarkRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 水印Id，即调用[AddLiveWatermark](/document/product/267/30154)接口返回的WatermarkId。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CreateLiveWatermarkRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveWatermarkRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "StreamName")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveWatermarkRuleRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveWatermarkRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePullStreamConfigRequest struct {
	*tchttp.BaseRequest

	// 源 Url ，用于拉流的地址。目前可支持直播流及点播文件。
	// 注意：
	// 1. 多个点播url之间使用空格拼接。
	// 2. 目前上限支持10个url。
	// 3. 支持拉流文件格式：flv，rtmp，hls，mp4。
	FromUrl *string `json:"FromUrl,omitempty" name:"FromUrl"`

	// 目的 Url ，用于推流的地址，目前限制该目标地址为腾讯域名。
	// 仅支持：rtmp 协议。
	ToUrl *string `json:"ToUrl,omitempty" name:"ToUrl"`

	// 选择完成转拉推的服务所在区域:
	// 1-深圳，
	// 2-上海，
	// 3-天津，
	// 4-中国香港。
	AreaId *int64 `json:"AreaId,omitempty" name:"AreaId"`

	// 选择完成转拉推服务使用的运营商网络：
	// 1-电信，
	// 2-移动，
	// 3-联通，
	// 4-其他。
	// 注：AreaId 为4的时候，IspId 只能为其他。
	IspId *int64 `json:"IspId,omitempty" name:"IspId"`

	// 开始时间。
	// 使用 UTC 格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于七天。
	// 使用 UTC 格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *CreatePullStreamConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePullStreamConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromUrl")
	delete(f, "ToUrl")
	delete(f, "AreaId")
	delete(f, "IspId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePullStreamConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePullStreamConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置成功后的 ID。
		ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePullStreamConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePullStreamConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordTaskRequest struct {
	*tchttp.BaseRequest

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 录制任务结束时间，Unix时间戳。设置时间必须大于StartTime及当前时间，且EndTime - StartTime不能超过24小时。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 录制任务开始时间，Unix时间戳。如果不填表示立即启动录制。StartTime不能超过当前时间+6天。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 推流类型，默认0。取值：
	// 0-直播推流。
	// 1-合成流，即 A+B=C 类型混流。
	StreamType *uint64 `json:"StreamType,omitempty" name:"StreamType"`

	// 录制模板ID，CreateLiveRecordTemplate 返回值。如果不填或者传入错误ID，则默认录制HLS格式、永久存储。
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 扩展字段，暂无定义。默认为空。
	Extension *string `json:"Extension,omitempty" name:"Extension"`
}

func (r *CreateRecordTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StreamName")
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "StreamType")
	delete(f, "TemplateId")
	delete(f, "Extension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecordTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID，全局唯一标识录制任务。返回TaskId字段说明录制任务创建成功。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRecordTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayStreamPlayInfo struct {

	// 数据时间点，格式：yyyy-mm-dd HH:MM:SS。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 带宽（单位Mbps）。
	Bandwidth *float64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 流量 （单位MB）。
	Flux *float64 `json:"Flux,omitempty" name:"Flux"`

	// 请求数。
	Request *uint64 `json:"Request,omitempty" name:"Request"`

	// 在线人数。
	Online *uint64 `json:"Online,omitempty" name:"Online"`
}

type DelayInfo struct {

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的 
	//  AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 延播时间，单位：秒。
	DelayInterval *uint64 `json:"DelayInterval,omitempty" name:"DelayInterval"`

	// 创建时间，UTC 时间。
	// 注意：UTC时间和北京时间相差8小时。
	// 例如：2019-06-18T12:00:00Z（为北京时间 2019 年 6 月 18 日 20 点 0 分 0 秒）。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 过期时间，UTC 时间。
	// 注意：UTC时间和北京时间相差8小时。
	// 例如：2019-06-18T12:00:00Z（为北京时间 2019 年 6 月 18 日 20 点 0 分 0 秒）。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 当前状态:
	// -1：已过期。
	// 1： 生效中。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DeleteLiveCallbackRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DeleteLiveCallbackRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveCallbackRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "AppName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveCallbackRuleRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveCallbackRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板 ID。
	// 1. 在创建回调模板接口 [CreateLiveCallbackTemplate](/document/product/267/32637) 调用的返回值中获取模板 ID。
	// 2. 可以从接口 [DescribeLiveCallbackTemplates](/document/product/267/32632) 查询已经创建的过的模板列表。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveCallbackTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveCallbackTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveCallbackTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveCertRequest struct {
	*tchttp.BaseRequest

	// DescribeLiveCerts接口获取到的证书Id。
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`
}

func (r *DeleteLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveCertRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveDomainRequest struct {
	*tchttp.BaseRequest

	// 要删除的域名
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 类型。0-推流，1-播放
	DomainType *uint64 `json:"DomainType,omitempty" name:"DomainType"`
}

func (r *DeleteLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "DomainType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLivePullStreamTaskRequest struct {
	*tchttp.BaseRequest

	// 任务 Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 操作人姓名。
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DeleteLivePullStreamTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLivePullStreamTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLivePullStreamTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLivePullStreamTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLivePullStreamTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLivePullStreamTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordRequest struct {
	*tchttp.BaseRequest

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 任务ID，由CreateLiveRecord接口返回。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteLiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StreamName")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveRecordRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DeleteLiveRecordRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveRecordRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveRecordRuleRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveRecordRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// DescribeRecordTemplates接口获取到的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveRecordTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DeleteLiveSnapshotRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveSnapshotRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveSnapshotRuleRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveSnapshotRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板 ID。
	// 1. 在创建截图模板接口 [CreateLiveSnapshotTemplate](/document/product/267/32624) 调用的返回值中获取。
	// 2. 可以从接口 [DescribeLiveSnapshotTemplates](/document/product/267/32619) 中查询已创建的截图模板列表。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveSnapshotTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeRuleRequest struct {
	*tchttp.BaseRequest

	// 播放域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 模板ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveTranscodeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveTranscodeRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "StreamName")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveTranscodeRuleRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveTranscodeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板 ID。
	// 1. 在创建转码模板接口 [CreateLiveTranscodeTemplate](/document/product/267/32646) 调用的返回值中获取模板 ID。
	// 2. 可以从接口 [DescribeLiveTranscodeTemplates](/document/product/267/32641) 查询已经创建的过的模板列表。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveTranscodeTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveTranscodeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// 水印 ID。
	// 在添加水印接口 [AddLiveWatermark](/document/product/267/30154) 调用返回值中获取水印 ID。
	// 或DescribeLiveWatermarks接口返回的水印ID。
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`
}

func (r *DeleteLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WatermarkId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveWatermarkRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveWatermarkRuleRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。与推流和播放地址中的 AppName 保持一致，默认为live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DeleteLiveWatermarkRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveWatermarkRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveWatermarkRuleRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveWatermarkRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePullStreamConfigRequest struct {
	*tchttp.BaseRequest

	// 配置 ID。
	// 1. 在添加拉流配置接口 [CreatePullStreamConfig](/document/api/267/30159) 调用返回值中获取配置 ID。
	// 2. 可以从接口 [DescribePullStreamConfigs](/document/api/267/30158) 中查询已创建过的拉流配置列表。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeletePullStreamConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePullStreamConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePullStreamConfigRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePullStreamConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID，CreateRecordTask返回。删除TaskId指定的录制任务。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteRecordTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRecordTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllStreamPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// 查询时间点，精确到分钟粒度，支持最近1个月的数据查询，数据延迟为5分钟左右，如果要查询实时的数据，建议传递5分钟前的时间点，格式为yyyy-mm-dd HH:MM:00。（只精确至分钟，秒数填00）。
	QueryTime *string `json:"QueryTime,omitempty" name:"QueryTime"`

	// 播放域名列表，若不填，表示总体数据。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`
}

func (r *DescribeAllStreamPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllStreamPlayInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueryTime")
	delete(f, "PlayDomains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllStreamPlayInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllStreamPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询时间点，回传的输入参数中的查询时间。
		QueryTime *string `json:"QueryTime,omitempty" name:"QueryTime"`

		// 数据信息列表。
		DataInfoList []*MonitorStreamPlayInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllStreamPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllStreamPlayInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAreaBillBandwidthAndFluxListRequest struct {
	*tchttp.BaseRequest

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS，起始和结束时间跨度不支持超过1天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 直播播放域名，若不填，表示总体数据。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`
}

func (r *DescribeAreaBillBandwidthAndFluxListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAreaBillBandwidthAndFluxListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PlayDomains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAreaBillBandwidthAndFluxListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAreaBillBandwidthAndFluxListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 明细数据信息。
		DataInfoList []*BillAreaInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAreaBillBandwidthAndFluxListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAreaBillBandwidthAndFluxListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillBandwidthAndFluxListRequest struct {
	*tchttp.BaseRequest

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS，起始和结束时间跨度不支持超过31天。支持最近3年的数据查询
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 直播播放域名，若不填，表示总体数据。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`

	// 可选值：
	// Mainland：查询国内数据，
	// Oversea：则查询国外数据，
	// 默认：查询国内+国外的数据。
	// 注：LEB（快直播）只支持国内+国外数据查询。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// 数据粒度，支持如下粒度：
	// 5：5分钟粒度，（跨度不支持超过1天），
	// 60：1小时粒度（跨度不支持超过一个月），
	// 1440：天粒度（跨度不支持超过一个月）。
	// 默认值：5。
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`

	// 服务名称，可选值包括LVB(标准直播)，LEB(快直播)，不填则查LVB+LEB总值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 大区，映射表如下：
	// China Mainland 中国大陆
	// Asia Pacific I 亚太一区
	// Asia Pacific II 亚太二区
	// Asia Pacific III 亚太三区
	// Europe 欧洲
	// North America 北美
	// South America 南美
	// Middle East 中东
	// Africa 非洲。
	RegionNames []*string `json:"RegionNames,omitempty" name:"RegionNames"`
}

func (r *DescribeBillBandwidthAndFluxListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillBandwidthAndFluxListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PlayDomains")
	delete(f, "MainlandOrOversea")
	delete(f, "Granularity")
	delete(f, "ServiceName")
	delete(f, "RegionNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillBandwidthAndFluxListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillBandwidthAndFluxListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 峰值带宽所在时间点，格式为yyyy-mm-dd HH:MM:SS。
		PeakBandwidthTime *string `json:"PeakBandwidthTime,omitempty" name:"PeakBandwidthTime"`

		// 峰值带宽，单位是Mbps。
		PeakBandwidth *float64 `json:"PeakBandwidth,omitempty" name:"PeakBandwidth"`

		// 95峰值带宽所在时间点，格式为yyyy-mm-dd HH:MM:SS。
		P95PeakBandwidthTime *string `json:"P95PeakBandwidthTime,omitempty" name:"P95PeakBandwidthTime"`

		// 95峰值带宽，单位是Mbps。
		P95PeakBandwidth *float64 `json:"P95PeakBandwidth,omitempty" name:"P95PeakBandwidth"`

		// 总流量，单位是MB。
		SumFlux *float64 `json:"SumFlux,omitempty" name:"SumFlux"`

		// 明细数据信息。
		DataInfoList []*BillDataInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillBandwidthAndFluxListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillBandwidthAndFluxListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCallbackRecordsListRequest struct {
	*tchttp.BaseRequest

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS，起始和结束时间跨度不支持超过1天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 流名称，精确匹配。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 页码。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 每页条数。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 事件类型。
	// 0: "断流",
	// 1: "推流",
	// 100: "录制"
	// 200: "截图回调"。
	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`

	// 回调结果。
	// 0为成功，其他为失败。
	ResultCode *uint64 `json:"ResultCode,omitempty" name:"ResultCode"`
}

func (r *DescribeCallbackRecordsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallbackRecordsListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "StreamName")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "EventType")
	delete(f, "ResultCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallbackRecordsListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCallbackRecordsListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 回调事件列表。
		DataInfoList []*CallbackEventInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 页码。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页条数。
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 总条数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCallbackRecordsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallbackRecordsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConcurrentRecordStreamNumRequest struct {
	*tchttp.BaseRequest

	// 直播类型，SlowLive：慢直播。
	// NormalLive：普通直播。
	LiveType *string `json:"LiveType,omitempty" name:"LiveType"`

	// 起始时间，格式：yyyy-mm-dd HH:MM:SS。
	// 可以查询最近180天的数据。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，格式：yyyy-mm-dd HH:MM:SS。
	// 时间跨度最大支持31天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 如果为空，查询所有地区数据；如果为“Mainland”，查询国内数据；如果为“Oversea”，则查询国外数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// 推流域名列表，不填表示总体数据。
	PushDomains []*string `json:"PushDomains,omitempty" name:"PushDomains"`
}

func (r *DescribeConcurrentRecordStreamNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrentRecordStreamNumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LiveType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MainlandOrOversea")
	delete(f, "PushDomains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrentRecordStreamNumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConcurrentRecordStreamNumResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计信息列表。
		DataInfoList []*ConcurrentRecordStreamNum `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConcurrentRecordStreamNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrentRecordStreamNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeliverBandwidthListRequest struct {
	*tchttp.BaseRequest

	// 起始时间，格式为%Y-%m-%d %H:%M:%S。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，格式为%Y-%m-%d %H:%M:%S，支持最近三个月的数据查询，时间跨度最大是1个月。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDeliverBandwidthListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeliverBandwidthListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeliverBandwidthListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeliverBandwidthListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 转推计费带宽数据
		DataInfoList []*BandwidthInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeliverBandwidthListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeliverBandwidthListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupProIspPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS
	// 时间跨度在（0,3小时]，支持最近1个月数据查询。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 播放域名，默认为不填，表示求总体数据。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`

	// 省份列表，默认不填，则返回各省份的数据。
	ProvinceNames []*string `json:"ProvinceNames,omitempty" name:"ProvinceNames"`

	// 运营商列表，默认不填，则返回整个运营商的数据。
	IspNames []*string `json:"IspNames,omitempty" name:"IspNames"`

	// 国内还是国外，如果为空，查询所有地区数据；如果为“Mainland”，查询国内数据；如果为“Oversea”，则查询国外数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`
}

func (r *DescribeGroupProIspPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupProIspPlayInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PlayDomains")
	delete(f, "ProvinceNames")
	delete(f, "IspNames")
	delete(f, "MainlandOrOversea")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupProIspPlayInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupProIspPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据内容。
		DataInfoList []*GroupProIspDataInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupProIspPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupProIspPlayInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHttpStatusInfoListRequest struct {
	*tchttp.BaseRequest

	// 起始时间，北京时间，
	// 格式：yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，北京时间，
	// 格式：yyyy-mm-dd HH:MM:SS。
	// 注：最大时间跨度支持1天，支持最近3个月的数据查询。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 播放域名列表。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`
}

func (r *DescribeHttpStatusInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHttpStatusInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PlayDomains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHttpStatusInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHttpStatusInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 播放状态码列表。
		DataInfoList []*HttpStatusData `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHttpStatusInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHttpStatusInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackRulesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveCallbackRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveCallbackRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveCallbackRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则信息列表。
		Rules []*CallBackRuleInfo `json:"Rules,omitempty" name:"Rules"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCallbackRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveCallbackRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板 ID。
	// 1. 在创建回调模板接口 [CreateLiveCallbackTemplate](/document/product/267/32637) 调用的返回值中获取模板 ID。
	// 2. 可以从接口 [DescribeLiveCallbackTemplates](/document/product/267/32632) 查询已经创建的过的模板列表。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveCallbackTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveCallbackTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveCallbackTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveCallbackTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCallbackTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板信息列表。
		Templates []*CallBackTemplateInfo `json:"Templates,omitempty" name:"Templates"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCallbackTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveCallbackTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertRequest struct {
	*tchttp.BaseRequest

	// DescribeLiveCerts接口获取到的证书Id。
	CertId *int64 `json:"CertId,omitempty" name:"CertId"`
}

func (r *DescribeLiveCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveCertRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveCertsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveCertsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveCertsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书信息列表。
		CertInfoSet []*CertInfo `json:"CertInfoSet,omitempty" name:"CertInfoSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveCertsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveCertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDelayInfoListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLiveDelayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDelayInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveDelayInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDelayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 延播信息列表。
		DelayInfoList []*DelayInfo `json:"DelayInfoList,omitempty" name:"DelayInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDelayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDelayInfoListResponse) FromJsonString(s string) error {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDomainCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveDomainCertRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDomainCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// 播放域名列表。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`
}

func (r *DescribeLiveDomainPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDomainPlayInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlayDomains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveDomainPlayInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据时间，格式为yyyy-mm-dd HH:MM:SS。
		Time *string `json:"Time,omitempty" name:"Time"`

		// 实时总带宽。
		TotalBandwidth *float64 `json:"TotalBandwidth,omitempty" name:"TotalBandwidth"`

		// 实时总流量。
		TotalFlux *float64 `json:"TotalFlux,omitempty" name:"TotalFlux"`

		// 总请求数。
		TotalRequest *uint64 `json:"TotalRequest,omitempty" name:"TotalRequest"`

		// 实时总连接数。
		TotalOnline *uint64 `json:"TotalOnline,omitempty" name:"TotalOnline"`

		// 分域名的数据情况。
		DomainInfoList []*DomainInfoList `json:"DomainInfoList,omitempty" name:"DomainInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDomainPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDomainPlayInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainRefererRequest struct {
	*tchttp.BaseRequest

	// 播放域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLiveDomainRefererRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDomainRefererRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveDomainRefererRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainRefererResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名 Referer 黑白名单配置。
		RefererAuthConfig *RefererAuthConfig `json:"RefererAuthConfig,omitempty" name:"RefererAuthConfig"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDomainRefererResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDomainRefererResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainRequest struct {
	*tchttp.BaseRequest

	// 域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *DescribeLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DomainInfo *DomainInfo `json:"DomainInfo,omitempty" name:"DomainInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainsRequest struct {
	*tchttp.BaseRequest

	// 域名状态过滤。0-停用，1-启用。
	DomainStatus *uint64 `json:"DomainStatus,omitempty" name:"DomainStatus"`

	// 域名类型过滤。0-推流，1-播放。
	DomainType *uint64 `json:"DomainType,omitempty" name:"DomainType"`

	// 分页大小，范围：10~100。默认10。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 取第几页，范围：1~100000。默认1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 0 普通直播 1慢直播 默认0。
	IsDelayLive *uint64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// 域名前缀。
	DomainPrefix *string `json:"DomainPrefix,omitempty" name:"DomainPrefix"`

	// 播放区域，只在 DomainType=1 时该参数有意义。
	// 1: 国内。
	// 2: 全球。
	// 3: 海外。
	PlayType *uint64 `json:"PlayType,omitempty" name:"PlayType"`
}

func (r *DescribeLiveDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainStatus")
	delete(f, "DomainType")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "IsDelayLive")
	delete(f, "DomainPrefix")
	delete(f, "PlayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数。
		AllCount *uint64 `json:"AllCount,omitempty" name:"AllCount"`

		// 域名详细信息列表。
		DomainList []*DomainInfo `json:"DomainList,omitempty" name:"DomainList"`

		// 可继续添加域名数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreateLimitCount *int64 `json:"CreateLimitCount,omitempty" name:"CreateLimitCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveForbidStreamListRequest struct {
	*tchttp.BaseRequest

	// 取得第几页，默认1。
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// 每页大小，最大100。 
	// 取值：1~100之前的任意整数。
	// 默认值：10。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 按流名称查询。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeLiveForbidStreamListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveForbidStreamListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveForbidStreamListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveForbidStreamListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的总个数。
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 总页数。
		TotalPage *int64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 分页的页码。
		PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页显示的条数。
		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

		// 禁推流列表。
		ForbidStreamList []*ForbidStreamInfo `json:"ForbidStreamList,omitempty" name:"ForbidStreamList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveForbidStreamListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveForbidStreamListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePackageInfoRequest struct {
	*tchttp.BaseRequest

	// 包类型，可选值：
	// 0：流量包
	// 1：转码包
	// 2: 连麦包。
	PackageType *int64 `json:"PackageType,omitempty" name:"PackageType"`

	// 排序规则:
	// 1. BuyTimeDesc： 最新购买的排在最前面
	// 2. BuyTimeAsc： 最老购买的排在最前面
	// 3. ExpireTimeDesc： 最后过期的排在最前面
	// 4. ExpireTimeAsc：最先过期的排在最前面。
	// 
	// 注意：
	// 1. PackageType 为 2（连麦包） 的时候，不支持 3、4 排序。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 取得第几页的数据，和 PageSize 同时传递才会生效。
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// 分页大小，和 PageNum 同时传递才会生效。
	// 取值：10 ～ 100 之间的任意整数。
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeLivePackageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePackageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageType")
	delete(f, "OrderBy")
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLivePackageInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePackageInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 套餐包信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		LivePackageInfoList []*LivePackageInfo `json:"LivePackageInfoList,omitempty" name:"LivePackageInfoList"`

		// 套餐包当前计费方式:
	// -1: 无计费方式或获取失败
	// 0: 无计费方式
	// 201: 月结带宽
	// 202: 月结流量
	// 203: 日结带宽
	// 204: 日结流量
	// 205: 日结时长
	// 206: 月结时长
	// 304: 日结流量。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PackageBillMode *int64 `json:"PackageBillMode,omitempty" name:"PackageBillMode"`

		// 总页数。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalPage *int64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 数据总条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 当前页数。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

		// 当前每页数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

		// 当请求参数 PackageType = 0 时生效，逗号分隔，从第一个到最后一个分别表示：
	// 标准直播，中国大陆（境内全地区）计费方式。
	// 标准直播，国际/港澳台（境外多地区）计费方式。
	// 快直播，中国大陆（境内全地区）计费方式。
	// 快直播，国际/港澳台（境外多地区）计费方式。
	// 注意：此字段可能返回 null，表示取不到有效值。
		FluxPackageBillMode *string `json:"FluxPackageBillMode,omitempty" name:"FluxPackageBillMode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLivePackageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePackageInfoResponse) FromJsonString(s string) error {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePlayAuthKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLivePlayAuthKeyRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePlayAuthKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePullStreamTasksRequest struct {
	*tchttp.BaseRequest

	// 任务 ID。 
	// 来源：调用 CreateLivePullStreamTask 接口时返回。
	// 不填默认查询所有任务，按更新时间倒序排序。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 取得第几页，默认值：1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 分页大小，默认值：10。
	// 取值范围：1~20 之前的任意整数。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeLivePullStreamTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePullStreamTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLivePullStreamTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLivePullStreamTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 直播拉流任务信息列表。
		TaskInfos []*PullStreamTaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos"`

		// 分页的页码。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页大小。
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 符合条件的总个数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 限制可创建的最大任务数。
		LimitTaskNum *uint64 `json:"LimitTaskNum,omitempty" name:"LimitTaskNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLivePullStreamTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePullStreamTasksResponse) FromJsonString(s string) error {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePushAuthKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLivePushAuthKeyRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveRecordRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveRecordRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则列表。
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveRecordRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// [DescribeLiveRecordTemplates](/document/product/267/32609)接口获取到的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveRecordTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplatesRequest struct {
	*tchttp.BaseRequest

	// 是否属于慢直播模板，默认：0。
	// 0： 标准直播。
	// 1：慢直播。
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`
}

func (r *DescribeLiveRecordTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveRecordTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsDelayLive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveRecordTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制模板信息列表。
		Templates []*RecordTemplateInfo `json:"Templates,omitempty" name:"Templates"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveSnapshotRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveSnapshotRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则列表。
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveSnapshotRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveSnapshotRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板 ID。
	// 调用 [CreateLiveSnapshotTemplate](/document/product/267/32624) 时返回的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveSnapshotTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveSnapshotTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveSnapshotTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveSnapshotTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 截图模板列表。
		Templates []*SnapshotTemplateInfo `json:"Templates,omitempty" name:"Templates"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveSnapshotTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveSnapshotTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamEventListRequest struct {
	*tchttp.BaseRequest

	// 起始时间。 
	// UTC 格式，例如：2018-12-29T19:00:00Z。
	// 支持查询60天内的历史记录。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	// UTC 格式，例如：2018-12-29T20:00:00Z。
	// 不超过当前时间，且和起始时间相差不得超过30天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 流名称，不支持通配符（*）查询，默认模糊匹配。
	// 可使用IsStrict字段改为精确查询。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 取得第几页。
	// 默认值：1。
	// 注： 目前只支持10000条内的查询。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 分页大小。
	// 最大值：100。
	// 取值范围：1~100 之间的任意整数。
	// 默认值：10。
	// 注： 目前只支持10000条内的查询。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 是否过滤，默认不过滤。
	// 0：不进行任何过滤。
	// 1：过滤掉开播失败的，只返回开播成功的。
	IsFilter *int64 `json:"IsFilter,omitempty" name:"IsFilter"`

	// 是否精确查询，默认模糊匹配。
	// 0：模糊匹配。
	// 1：精确查询。
	// 注：使用StreamName时该参数生效。
	IsStrict *int64 `json:"IsStrict,omitempty" name:"IsStrict"`

	// 是否按结束时间正序显示，默认逆序。
	// 0：逆序。
	// 1：正序。
	IsAsc *int64 `json:"IsAsc,omitempty" name:"IsAsc"`
}

func (r *DescribeLiveStreamEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppName")
	delete(f, "DomainName")
	delete(f, "StreamName")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "IsFilter")
	delete(f, "IsStrict")
	delete(f, "IsAsc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveStreamEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamEventListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 推断流事件列表。
		EventList []*StreamEventInfo `json:"EventList,omitempty" name:"EventList"`

		// 分页的页码。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页大小。
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 符合条件的总个数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineListRequest struct {
	*tchttp.BaseRequest

	// 推流域名。多域名用户需要填写 DomainName。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。多路径用户需要填写 AppName。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 取得第几页，默认1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 每页大小，最大100。 
	// 取值：10~100之间的任意整数。
	// 默认值：10。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 流名称，用于精确查询。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeLiveStreamOnlineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamOnlineListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveStreamOnlineListRequest has unknown keys!", "")
	}
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
		OnlineInfo []*StreamOnlineInfo `json:"OnlineInfo,omitempty" name:"OnlineInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamOnlineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamOnlineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPublishedListRequest struct {
	*tchttp.BaseRequest

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 结束时间。
	// UTC 格式，例如：2016-06-30T19:00:00Z。
	// 不超过当前时间。
	// 注意：EndTime和StartTime相差不可超过30天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 起始时间。 
	// UTC 格式，例如：2016-06-29T19:00:00Z。
	// 最长支持查询60天内数据。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。不支持模糊匹配。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 取得第几页。
	// 默认值：1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 分页大小。
	// 最大值：100。
	// 取值范围：10~100 之前的任意整数。
	// 默认值：10。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 流名称，支持模糊匹配。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeLiveStreamPublishedListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamPublishedListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "AppName")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveStreamPublishedListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPublishedListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 推流记录信息。
		PublishInfo []*StreamName `json:"PublishInfo,omitempty" name:"PublishInfo"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamPublishedListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPushInfoListRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	PushDomain *string `json:"PushDomain,omitempty" name:"PushDomain"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 页数，
	// 范围[1,10000]，
	// 默认值：1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 每页个数，
	// 范围：[1,1000]，
	// 默认值： 200。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeLiveStreamPushInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamPushInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PushDomain")
	delete(f, "AppName")
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveStreamPushInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPushInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 直播流的统计信息列表。
		DataInfoList []*PushDataInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 所有在线流的总数量。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 当前数据所在页码。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页的在线流的个数。
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamPushInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamPushInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamStateRequest struct {
	*tchttp.BaseRequest

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "DomainName")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveStreamStateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeDetailInfoRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	PushDomain *string `json:"PushDomain,omitempty" name:"PushDomain"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 查询时间，北京时间，
	// 格式：yyyymmdd。
	// 注意：支持查询近1个月内某天的详细数据，截止到昨天。
	DayTime *string `json:"DayTime,omitempty" name:"DayTime"`

	// 页数，默认1，
	// 不超过100页。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 每页个数，默认20，
	// 范围：[10,1000]。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 起始天时间，北京时间，
	// 格式：yyyymmdd。
	// 注意：支持查询近1个月内的详细数据。
	StartDayTime *string `json:"StartDayTime,omitempty" name:"StartDayTime"`

	// 结束天时间，北京时间，
	// 格式：yyyymmdd。
	// 注意：支持查询近1个月内的详细数据，截止到昨天，注意DayTime 与（StartDayTime，EndDayTime）必须要传一个，如果都传，会以DayTime为准 。
	EndDayTime *string `json:"EndDayTime,omitempty" name:"EndDayTime"`
}

func (r *DescribeLiveTranscodeDetailInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTranscodeDetailInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PushDomain")
	delete(f, "StreamName")
	delete(f, "DayTime")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "StartDayTime")
	delete(f, "EndDayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveTranscodeDetailInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeDetailInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计数据列表。
		DataInfoList []*TranscodeDetailInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 页码。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页个数。
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 总个数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeDetailInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTranscodeDetailInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeRulesRequest struct {
	*tchttp.BaseRequest

	// 要筛选的模板ID数组。
	TemplateIds []*int64 `json:"TemplateIds,omitempty" name:"TemplateIds"`

	// 要筛选的域名数组。
	DomainNames []*string `json:"DomainNames,omitempty" name:"DomainNames"`
}

func (r *DescribeLiveTranscodeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTranscodeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateIds")
	delete(f, "DomainNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveTranscodeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 转码规则列表。
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTranscodeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板 ID。
	// 注意：在创建转码模板接口 [CreateLiveTranscodeTemplate](/document/product/267/32646) 调用的返回值中获取模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveTranscodeTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTranscodeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveTranscodeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 转码模板列表。
		Templates []*TemplateInfo `json:"Templates,omitempty" name:"Templates"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTranscodeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTotalInfoRequest struct {
	*tchttp.BaseRequest

	// 开始时间，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 推流域名列表，若不填，表示查询所有域名总体数据。
	// 指定域名时返回1小时粒度数据。
	PushDomains []*string `json:"PushDomains,omitempty" name:"PushDomains"`

	// 可选值：
	// Mainland：查询中国大陆（境内）数据，
	// Oversea：则查询国际/港澳台（境外）数据，
	// 默认：查询全球地区（境内+境外）的数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`
}

func (r *DescribeLiveTranscodeTotalInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTranscodeTotalInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PushDomains")
	delete(f, "MainlandOrOversea")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveTranscodeTotalInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveTranscodeTotalInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DataInfoList []*TranscodeTotalInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveTranscodeTotalInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTranscodeTotalInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// DescribeLiveWatermarks接口返回的水印 ID。
	WatermarkId *uint64 `json:"WatermarkId,omitempty" name:"WatermarkId"`
}

func (r *DescribeLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WatermarkId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveWatermarkRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveWatermarkRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveWatermarkRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarkRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 水印规则列表。
		Rules []*RuleInfo `json:"Rules,omitempty" name:"Rules"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveWatermarkRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveWatermarksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveWatermarksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveWatermarksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 水印总个数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 水印信息列表。
		WatermarkList []*WatermarkInfo `json:"WatermarkList,omitempty" name:"WatermarkList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveWatermarksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveWatermarksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogDownloadListRequest struct {
	*tchttp.BaseRequest

	// 开始时间，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	// 注意：结束时间 - 开始时间 <=7天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 域名列表。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`
}

func (r *DescribeLogDownloadListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogDownloadListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PlayDomains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogDownloadListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogDownloadListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志信息列表。
		LogInfoList []*LogInfo `json:"LogInfoList,omitempty" name:"LogInfoList"`

		// 总条数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogDownloadListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogDownloadListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePlayErrorCodeDetailInfoListRequest struct {
	*tchttp.BaseRequest

	// 起始时间，北京时间，
	// 格式：yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，北京时间，
	// 格式：yyyy-mm-dd HH:MM:SS。
	// 注：EndTime 和 StartTime 只支持最近1天的数据查询。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询粒度：
	// 1-1分钟粒度。
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`

	// 是，可选值包括”4xx”,”5xx”，支持”4xx,5xx”等这种混合模式。
	StatType *string `json:"StatType,omitempty" name:"StatType"`

	// 播放域名列表。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`
}

func (r *DescribePlayErrorCodeDetailInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlayErrorCodeDetailInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Granularity")
	delete(f, "StatType")
	delete(f, "PlayDomains")
	delete(f, "MainlandOrOversea")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePlayErrorCodeDetailInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePlayErrorCodeDetailInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计信息列表。
		HttpCodeList []*HttpCodeInfo `json:"HttpCodeList,omitempty" name:"HttpCodeList"`

		// 统计类型。
		StatType *string `json:"StatType,omitempty" name:"StatType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePlayErrorCodeDetailInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlayErrorCodeDetailInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePlayErrorCodeSumInfoListRequest struct {
	*tchttp.BaseRequest

	// 起始时间点，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	// 注：EndTime 和 StartTime 只支持最近1天的数据查询。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 播放域名列表，不填表示总体数据。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`

	// 页数，范围[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 每页个数，范围：[1,1000]，默认值是20。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// 分组参数，可选值：CountryProIsp（默认值），Country（国家），默认是按照国家+省份+运营商来进行分组；目前国外的省份和运营商暂时无法识别。
	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`

	// 输出字段使用的语言，可选值：Chinese（默认值），English，目前国家，省份和运营商支持多语言。
	OutLanguage *string `json:"OutLanguage,omitempty" name:"OutLanguage"`
}

func (r *DescribePlayErrorCodeSumInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlayErrorCodeSumInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PlayDomains")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "MainlandOrOversea")
	delete(f, "GroupType")
	delete(f, "OutLanguage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePlayErrorCodeSumInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePlayErrorCodeSumInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分省份分运营商错误码为2或3或4或5开头的状态码数据信息。
		ProIspInfoList []*ProIspPlayCodeDataInfo `json:"ProIspInfoList,omitempty" name:"ProIspInfoList"`

		// 所有状态码的加和的次数。
		TotalCodeAll *uint64 `json:"TotalCodeAll,omitempty" name:"TotalCodeAll"`

		// 状态码为4开头的总次数。
		TotalCode4xx *uint64 `json:"TotalCode4xx,omitempty" name:"TotalCode4xx"`

		// 状态码为5开头的总次数。
		TotalCode5xx *uint64 `json:"TotalCode5xx,omitempty" name:"TotalCode5xx"`

		// 各状态码的总次数。
		TotalCodeList []*PlayCodeTotalInfo `json:"TotalCodeList,omitempty" name:"TotalCodeList"`

		// 页号。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页大小。
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 总记录数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 状态码为2开头的总次数。
		TotalCode2xx *uint64 `json:"TotalCode2xx,omitempty" name:"TotalCode2xx"`

		// 状态码为3开头的总次数。
		TotalCode3xx *uint64 `json:"TotalCode3xx,omitempty" name:"TotalCode3xx"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePlayErrorCodeSumInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlayErrorCodeSumInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProIspPlaySumInfoListRequest struct {
	*tchttp.BaseRequest

	// 起始时间，北京时间，
	// 格式：yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，北京时间，
	// 格式：yyyy-mm-dd HH:MM:SS。
	// 注：EndTime 和 StartTime 只支持最近1天的数据查询。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计的类型，可选值：”Province”(省份)，”Isp”(运营商)，“CountryOrArea”(国家或地区)。
	StatType *string `json:"StatType,omitempty" name:"StatType"`

	// 播放域名列表，不填则为全部。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`

	// 页号，范围是[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 每页个数，范围是[1,1000]，默认值是20。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// 输出字段使用的语言，可选值：Chinese（默认值），English；目前国家，省份和运营商支持多语言。
	OutLanguage *string `json:"OutLanguage,omitempty" name:"OutLanguage"`
}

func (r *DescribeProIspPlaySumInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProIspPlaySumInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "StatType")
	delete(f, "PlayDomains")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "MainlandOrOversea")
	delete(f, "OutLanguage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProIspPlaySumInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProIspPlaySumInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总流量。
		TotalFlux *float64 `json:"TotalFlux,omitempty" name:"TotalFlux"`

		// 总请求数。
		TotalRequest *uint64 `json:"TotalRequest,omitempty" name:"TotalRequest"`

		// 统计的类型。
		StatType *string `json:"StatType,omitempty" name:"StatType"`

		// 每页的记录数。
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 页号。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 总记录数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 省份，运营商，国家或地区汇总数据列表。
		DataInfoList []*ProIspPlaySumInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 下载速度，单位：MB/s，计算方式：总流量/总时长。
		AvgFluxPerSecond *float64 `json:"AvgFluxPerSecond,omitempty" name:"AvgFluxPerSecond"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProIspPlaySumInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProIspPlaySumInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProvinceIspPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// 起始时间点，当前使用北京时间，
	// 例：2019-02-21 10:00:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，当前使用北京时间，
	// 例：2019-02-21 12:00:00。
	// 注：EndTime 和 StartTime 只支持最近1天的数据查询。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 支持如下粒度：
	// 1：1分钟粒度（跨度不支持超过1天）
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`

	// 统计指标类型：
	// “Bandwidth”：带宽
	// “FluxPerSecond”：平均流量
	// “Flux”：流量
	// “Request”：请求数
	// “Online”：并发连接数
	StatType *string `json:"StatType,omitempty" name:"StatType"`

	// 播放域名列表。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`

	// 要查询的省份（地区）英文名称列表，如 Beijing。
	ProvinceNames []*string `json:"ProvinceNames,omitempty" name:"ProvinceNames"`

	// 要查询的运营商英文名称列表，如 China Mobile ，如果为空，查询所有运营商的数据。
	IspNames []*string `json:"IspNames,omitempty" name:"IspNames"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// ip类型：
	// “Ipv6”：Ipv6数据
	// 如果为空，查询总的数据；
	IpType *string `json:"IpType,omitempty" name:"IpType"`
}

func (r *DescribeProvinceIspPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProvinceIspPlayInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Granularity")
	delete(f, "StatType")
	delete(f, "PlayDomains")
	delete(f, "ProvinceNames")
	delete(f, "IspNames")
	delete(f, "MainlandOrOversea")
	delete(f, "IpType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProvinceIspPlayInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProvinceIspPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 播放信息列表。
		DataInfoList []*PlayStatInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 统计的类型，和输入参数保持一致。
		StatType *string `json:"StatType,omitempty" name:"StatType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProvinceIspPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProvinceIspPlayInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePullStreamConfigsRequest struct {
	*tchttp.BaseRequest

	// 配置 ID。
	// 获取途径：从 CreatePullStreamConfig 接口返回值获取。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribePullStreamConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePullStreamConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePullStreamConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePullStreamConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 拉流配置。
		PullStreamConfigs []*PullStreamConfig `json:"PullStreamConfigs,omitempty" name:"PullStreamConfigs"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePullStreamConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePullStreamConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePushBandwidthAndFluxListRequest struct {
	*tchttp.BaseRequest

	// 起始时间点，格式为 yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为 yyyy-mm-dd HH:MM:SS，起始和结束时间跨度不支持超过31天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 域名，可以填多个，若不填，表示总体数据。
	PushDomains []*string `json:"PushDomains,omitempty" name:"PushDomains"`

	// 可选值：
	// Mainland：查询中国大陆（境内）数据，
	// Oversea：则查询国际/港澳台（境外）数据，
	// 不填则默认查询全球地区（境内+境外）的数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// 数据粒度，支持如下粒度：
	// 5：5分钟粒度，（跨度不支持超过1天），
	// 60：1小时粒度（跨度不支持超过一个月），
	// 1440：天粒度（跨度不支持超过一个月）。
	// 默认值：5。
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`

	// 大区，映射表如下：
	// China Mainland 中国大陆
	// Asia Pacific I 亚太一区
	// Asia Pacific II 亚太二区
	// Asia Pacific III 亚太三区
	// Europe 欧洲
	// North America 北美
	// South America 南美
	// Middle East 中东
	// Africa 非洲。
	RegionNames []*string `json:"RegionNames,omitempty" name:"RegionNames"`
}

func (r *DescribePushBandwidthAndFluxListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePushBandwidthAndFluxListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PushDomains")
	delete(f, "MainlandOrOversea")
	delete(f, "Granularity")
	delete(f, "RegionNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePushBandwidthAndFluxListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePushBandwidthAndFluxListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 峰值带宽所在时间点，格式为 yyyy-mm-dd HH:MM:SS。
		PeakBandwidthTime *string `json:"PeakBandwidthTime,omitempty" name:"PeakBandwidthTime"`

		// 峰值带宽，单位是 Mbps。
		PeakBandwidth *float64 `json:"PeakBandwidth,omitempty" name:"PeakBandwidth"`

		// 95峰值带宽所在时间点，格式为 yyyy-mm-dd HH:MM:SS。
		P95PeakBandwidthTime *string `json:"P95PeakBandwidthTime,omitempty" name:"P95PeakBandwidthTime"`

		// 95峰值带宽，单位是 Mbps。
		P95PeakBandwidth *float64 `json:"P95PeakBandwidth,omitempty" name:"P95PeakBandwidth"`

		// 总流量，单位是 MB。
		SumFlux *float64 `json:"SumFlux,omitempty" name:"SumFlux"`

		// 明细数据信息。
		DataInfoList []*BillDataInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePushBandwidthAndFluxListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePushBandwidthAndFluxListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordTaskRequest struct {
	*tchttp.BaseRequest

	// 查询任务开始时间，Unix 时间戳。设置时间不早于当前时间之前90天的时间，且查询时间跨度不超过一周。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询任务结束时间，Unix 时间戳。EndTime 必须大于 StartTime，设置时间不早于当前时间之前90天的时间，且查询时间跨度不超过一周。（注意：任务开始结束时间必须在查询时间范围内）。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`
}

func (r *DescribeRecordTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "StreamName")
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "ScrollToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 翻页标识，当请求未返回所有数据，该字段表示下一条记录的 Token。当该字段为空，说明已无更多数据。
		ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

		// 录制任务列表。当该字段为空，说明已返回所有数据。
		TaskList []*RecordTask `json:"TaskList,omitempty" name:"TaskList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenShotSheetNumListRequest struct {
	*tchttp.BaseRequest

	// utc起始时间，格式为yyyy-mm-ddTHH:MM:SSZ
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// utc结束时间，格式为yyyy-mm-ddTHH:MM:SSZ，支持查询最近1年数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 地域信息，可选值包括Mainland，Oversea，前者是查询中国大陆范围内的数据，后者是除中国大陆范围之外的数据，若不传该参数，则查询所有地区的数据。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 推流域名（支持查询2019年11 月1日之后的域名维度数据）。
	PushDomains []*string `json:"PushDomains,omitempty" name:"PushDomains"`

	// 数据维度，数据延迟1个半小时，可选值包括：1、Minute（5分钟粒度，最大支持查询时间范围是31天），2、Day（天粒度，默认值，按照北京时间做跨天处理，最大支持查询时间范围是186天当天）。
	Granularity *string `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeScreenShotSheetNumListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScreenShotSheetNumListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Zone")
	delete(f, "PushDomains")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScreenShotSheetNumListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScreenShotSheetNumListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据信息列表。
		DataInfoList []*TimeValue `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScreenShotSheetNumListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScreenShotSheetNumListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamDayPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// 日期，格式：YYYY-mm-dd。
	// 第二天凌晨3点出昨天的数据，建议在这个时间点之后查询最新数据。支持最近3个月的数据查询。
	DayTime *string `json:"DayTime,omitempty" name:"DayTime"`

	// 播放域名。
	PlayDomain *string `json:"PlayDomain,omitempty" name:"PlayDomain"`

	// 页号，范围[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 每页个数，范围[100,1000]，默认值是1000。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 可选值：
	// Mainland：查询国内数据，
	// Oversea：则查询国外数据，
	// 默认：查询国内+国外的数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// 服务名称，可选值包括LVB(标准直播)，LEB(快直播)，不填则查LVB+LEB总值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DescribeStreamDayPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamDayPlayInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DayTime")
	delete(f, "PlayDomain")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "MainlandOrOversea")
	delete(f, "ServiceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamDayPlayInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamDayPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 播放数据信息列表。
		DataInfoList []*PlayDataInfoByStream `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 总数量。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 当前数据所处页码。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页个数。
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStreamDayPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamDayPlayInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// 开始时间，北京时间，格式为yyyy-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，北京时间，格式为yyyy-mm-dd HH:MM:SS，
	// 结束时间 和 开始时间跨度不支持超过24小时，支持距当前时间30天内的数据查询。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 播放域名，
	// 若不填，则为查询所有播放域名的在线流数据。
	PlayDomain *string `json:"PlayDomain,omitempty" name:"PlayDomain"`

	// 流名称，精确匹配。
	// 若不填，则为查询总体播放数据。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 推流路径，与播放地址中的AppName保持一致，会精确匹配，在同时传递了StreamName时生效。
	// 若不填，则为查询总体播放数据。
	// 注意：按AppName查询请先联系工单申请，开通后配置生效预计需要5个工作日左右，具体时间以最终回复为准。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 服务名称，可选值包括LVB(标准直播)，LEB(快直播)，不填则查LVB+LEB总值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DescribeStreamPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamPlayInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PlayDomain")
	delete(f, "StreamName")
	delete(f, "AppName")
	delete(f, "ServiceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamPlayInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计信息列表，时间粒度是1分钟。
		DataInfoList []*DayStreamPlayInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStreamPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamPlayInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamPushInfoListRequest struct {
	*tchttp.BaseRequest

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS，支持查询最近7天数据，建议查询时间跨度在3小时之内。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 推流域名。
	PushDomain *string `json:"PushDomain,omitempty" name:"PushDomain"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DescribeStreamPushInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamPushInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StreamName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PushDomain")
	delete(f, "AppName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamPushInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamPushInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据列表。
		DataInfoList []*PushQualityData `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStreamPushInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamPushInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopClientIpSumInfoListRequest struct {
	*tchttp.BaseRequest

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS
	// 时间跨度在[0,4小时]，支持最近1天数据查询。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 播放域名，默认为不填，表示求总体数据。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`

	// 页号，范围是[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 每页个数，范围是[1,1000]，默认值是20。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 排序指标，可选值包括TotalRequest（默认值），FailedRequest,TotalFlux。
	OrderParam *string `json:"OrderParam,omitempty" name:"OrderParam"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// 输出字段使用的语言，可选值：Chinese（默认值），English；目前国家，省份和运营商支持多语言。
	OutLanguage *string `json:"OutLanguage,omitempty" name:"OutLanguage"`
}

func (r *DescribeTopClientIpSumInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopClientIpSumInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PlayDomains")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "OrderParam")
	delete(f, "MainlandOrOversea")
	delete(f, "OutLanguage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopClientIpSumInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopClientIpSumInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 页号，范围是[1,1000]，默认值是1。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页个数，范围是[1,1000]，默认值是20。
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 排序指标，可选值包括”TotalRequest”，”FailedRequest”,“TotalFlux”。
		OrderParam *string `json:"OrderParam,omitempty" name:"OrderParam"`

		// 记录总数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 记录总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 数据内容。
		DataInfoList []*ClientIpPlaySumInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopClientIpSumInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopClientIpSumInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadStreamNumsRequest struct {
	*tchttp.BaseRequest

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS，起始和结束时间跨度不支持超过31天。支持最近31天的数据查询
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 直播域名，若不填，表示总体数据。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 数据粒度，支持如下粒度：
	// 5：5分钟粒度，（跨度不支持超过1天），
	// 1440：天粒度（跨度不支持超过一个月）。
	// 默认值：5。
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeUploadStreamNumsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadStreamNumsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Domains")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUploadStreamNumsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadStreamNumsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 明细数据信息
		DataInfoList []*ConcurrentRecordStreamNum `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUploadStreamNumsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadStreamNumsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVisitTopSumInfoListRequest struct {
	*tchttp.BaseRequest

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS
	// 时间跨度在(0,4小时]，支持最近1天数据查询。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 峰值指标，可选值包括”Domain”，”StreamId”。
	TopIndex *string `json:"TopIndex,omitempty" name:"TopIndex"`

	// 播放域名，默认为不填，表示求总体数据。
	PlayDomains []*string `json:"PlayDomains,omitempty" name:"PlayDomains"`

	// 页号，
	// 范围是[1,1000]，
	// 默认值是1。
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 每页个数，范围是[1,1000]，
	// 默认值是20。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 排序指标，可选值包括” AvgFluxPerSecond”，”TotalRequest”（默认）,“TotalFlux”。
	OrderParam *string `json:"OrderParam,omitempty" name:"OrderParam"`
}

func (r *DescribeVisitTopSumInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVisitTopSumInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TopIndex")
	delete(f, "PlayDomains")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "OrderParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVisitTopSumInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVisitTopSumInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 页号，
	// 范围是[1,1000]，
	// 默认值是1。
		PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

		// 每页个数，范围是[1,1000]，
	// 默认值是20。
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 峰值指标，可选值包括”Domain”，”StreamId”。
		TopIndex *string `json:"TopIndex,omitempty" name:"TopIndex"`

		// 排序指标，可选值包括” AvgFluxPerSecond”(按每秒平均流量排序)，”TotalRequest”（默认，按总请求数排序）,“TotalFlux”（按总流量排序）。
		OrderParam *string `json:"OrderParam,omitempty" name:"OrderParam"`

		// 记录总数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 记录总页数。
		TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// 数据内容。
		DataInfoList []*PlaySumStatInfo `json:"DataInfoList,omitempty" name:"DataInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVisitTopSumInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVisitTopSumInfoListResponse) FromJsonString(s string) error {
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
	// 0：用户添加证书，
	// 1：腾讯云托管证书。
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 证书过期时间，UTC格式。
	CertExpireTime *string `json:"CertExpireTime,omitempty" name:"CertExpireTime"`

	// 使用此证书的域名名称。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 证书状态。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 证书本身标识的域名列表。
	// 比如: ["*.x.com"]
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertDomains []*string `json:"CertDomains,omitempty" name:"CertDomains"`

	// 腾讯云ssl的证书Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudCertId *string `json:"CloudCertId,omitempty" name:"CloudCertId"`
}

type DomainDetailInfo struct {

	// 国内还是国外:
	// Mainland: 表示国内数据。
	// Oversea: 表示国外数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitempty" name:"MainlandOrOversea"`

	// 带宽，单位: Mbps。
	Bandwidth *float64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 流量，单位: MB。
	Flux *float64 `json:"Flux,omitempty" name:"Flux"`

	// 人数。
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// 请求数。
	Request *uint64 `json:"Request,omitempty" name:"Request"`
}

type DomainInfo struct {

	// 直播域名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 域名类型:
	// 0: 推流。
	// 1: 播放。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 域名状态:
	// 0: 停用。
	// 1: 启用。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 添加时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否有 CName 到固定规则域名:
	// 0: 否。
	// 1: 是。
	BCName *uint64 `json:"BCName,omitempty" name:"BCName"`

	// cname 对应的域名。
	TargetDomain *string `json:"TargetDomain,omitempty" name:"TargetDomain"`

	// 播放区域，只在 Type=1 时该参数有意义。
	// 1: 国内。
	// 2: 全球。
	// 3: 海外。
	PlayType *int64 `json:"PlayType,omitempty" name:"PlayType"`

	// 是否慢直播:
	// 0: 普通直播。
	// 1: 慢直播。
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// 当前客户使用的 cname 信息。
	CurrentCName *string `json:"CurrentCName,omitempty" name:"CurrentCName"`

	// 失效参数，可忽略。
	RentTag *int64 `json:"RentTag,omitempty" name:"RentTag"`

	// 失效参数，可忽略。
	RentExpireTime *string `json:"RentExpireTime,omitempty" name:"RentExpireTime"`

	// 0: 标准直播。
	// 1: 小程序直播。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMiniProgramLive *int64 `json:"IsMiniProgramLive,omitempty" name:"IsMiniProgramLive"`
}

type DomainInfoList struct {

	// 域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 明细信息。
	DetailInfoList []*DomainDetailInfo `json:"DetailInfoList,omitempty" name:"DetailInfoList"`
}

type DropLiveStreamRequest struct {
	*tchttp.BaseRequest

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DropLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropLiveStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StreamName")
	delete(f, "DomainName")
	delete(f, "AppName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DropLiveStreamRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropLiveStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableLiveDomainRequest struct {
	*tchttp.BaseRequest

	// 待启用的直播域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *EnableLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableLiveDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableLiveDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EnableLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableLiveDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveDomainRequest struct {
	*tchttp.BaseRequest

	// 待停用的直播域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *ForbidLiveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForbidLiveDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ForbidLiveDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForbidLiveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForbidLiveDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveStreamRequest struct {
	*tchttp.BaseRequest

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 恢复流的时间。UTC 格式，例如：2018-11-29T19:00:00Z。
	// 注意：
	// 1. 默认禁播7天，且最长支持禁播90天。
	// 2. 北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	ResumeTime *string `json:"ResumeTime,omitempty" name:"ResumeTime"`

	// 禁推原因。
	// 注明：请务必填写禁推原因，防止误操作。
	// 长度限制：2048字节。
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *ForbidLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForbidLiveStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "DomainName")
	delete(f, "StreamName")
	delete(f, "ResumeTime")
	delete(f, "Reason")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ForbidLiveStreamRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForbidLiveStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ForbidStreamInfo struct {

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 禁推过期时间。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type GroupProIspDataInfo struct {

	// 省份。
	ProvinceName *string `json:"ProvinceName,omitempty" name:"ProvinceName"`

	// 运营商。
	IspName *string `json:"IspName,omitempty" name:"IspName"`

	// 分钟维度的明细数据。
	DetailInfoList []*CdnPlayStatData `json:"DetailInfoList,omitempty" name:"DetailInfoList"`
}

type HlsSpecialParam struct {

	// HLS续流超时时间。
	// 取值范围[0，1800]。
	FlowContinueDuration *uint64 `json:"FlowContinueDuration,omitempty" name:"FlowContinueDuration"`
}

type HttpCodeInfo struct {

	// HTTP协议返回码。
	// 例："2xx", "3xx", "4xx", "5xx"。
	HttpCode *string `json:"HttpCode,omitempty" name:"HttpCode"`

	// 统计信息，对于无数据的时间点，会补0。
	ValueList []*HttpCodeValue `json:"ValueList,omitempty" name:"ValueList"`
}

type HttpCodeValue struct {

	// 时间，格式：yyyy-mm-dd HH:MM:SS。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 次数。
	Numbers *uint64 `json:"Numbers,omitempty" name:"Numbers"`

	// 占比。
	Percentage *float64 `json:"Percentage,omitempty" name:"Percentage"`
}

type HttpStatusData struct {

	// 数据时间点，
	// 格式：yyyy-mm-dd HH:MM:SS。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 播放状态码详细信息。
	HttpStatusInfoList []*HttpStatusInfo `json:"HttpStatusInfoList,omitempty" name:"HttpStatusInfoList"`
}

type HttpStatusInfo struct {

	// 播放HTTP状态码。
	HttpStatus *string `json:"HttpStatus,omitempty" name:"HttpStatus"`

	// 个数。
	Num *uint64 `json:"Num,omitempty" name:"Num"`
}

type LivePackageInfo struct {

	// 包 ID。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 总量。
	// 注意：当为流量包时单位为字节。
	// 当为转码包时单位为分钟。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 使用量。
	// 注意：当为流量包时单位为字节。
	// 当为转码包时单位为分钟。
	// 当为连麦包时单位为小时。
	Used *int64 `json:"Used,omitempty" name:"Used"`

	// 剩余量。
	// 注意：当为流量包时单位为字节。
	// 当为转码包时单位为分钟。
	// 当为连麦包时单位为小时。
	Left *int64 `json:"Left,omitempty" name:"Left"`

	// 购买时间。
	BuyTime *string `json:"BuyTime,omitempty" name:"BuyTime"`

	// 过期时间。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 包类型，可选值:
	// 0: 流量包。
	// 1: 普通转码包。
	// 2: 极速高清包。
	// 3: 连麦包。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 包状态，可选值:
	// 0: 未使用。
	// 1: 使用中。
	// 2: 已过期。
	// 3: 已冻结。
	// 4: 已耗尽。
	// 5: 已退款
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type LogInfo struct {

	// 日志名称。
	LogName *string `json:"LogName,omitempty" name:"LogName"`

	// 日志 URL。
	LogUrl *string `json:"LogUrl,omitempty" name:"LogUrl"`

	// 日志生成时间。
	LogTime *string `json:"LogTime,omitempty" name:"LogTime"`

	// 文件大小。
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`
}

type ModifyLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest

	// DescribeLiveCallbackTemplates接口返回的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 开播回调 URL。
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitempty" name:"StreamBeginNotifyUrl"`

	// 断流回调 URL。
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitempty" name:"StreamEndNotifyUrl"`

	// 录制回调 URL。
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitempty" name:"RecordNotifyUrl"`

	// 截图回调 URL。
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitempty" name:"SnapshotNotifyUrl"`

	// 鉴黄回调 URL。
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitempty" name:"PornCensorshipNotifyUrl"`

	// 回调 Key，回调 URL 公用，回调签名详见事件消息通知文档。
	// [事件消息通知](/document/product/267/32744)。
	CallbackKey *string `json:"CallbackKey,omitempty" name:"CallbackKey"`
}

func (r *ModifyLiveCallbackTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveCallbackTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateName")
	delete(f, "Description")
	delete(f, "StreamBeginNotifyUrl")
	delete(f, "StreamEndNotifyUrl")
	delete(f, "RecordNotifyUrl")
	delete(f, "SnapshotNotifyUrl")
	delete(f, "PornCensorshipNotifyUrl")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveCallbackTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertId")
	delete(f, "CertType")
	delete(f, "CertName")
	delete(f, "HttpsCrt")
	delete(f, "HttpsKey")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveCertRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveDomainCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "CertId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveDomainCertRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveDomainCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveDomainRefererRequest struct {
	*tchttp.BaseRequest

	// 播放域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 是否开启当前域名的 Referer 黑白名单鉴权。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 名单类型，0：黑名单，1：白名单。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 是否允许空 Referer，0：不允许，1：允许。
	AllowEmpty *int64 `json:"AllowEmpty,omitempty" name:"AllowEmpty"`

	// Referer 名单列表，以;分隔。
	Rules *string `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyLiveDomainRefererRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveDomainRefererRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "Enable")
	delete(f, "Type")
	delete(f, "AllowEmpty")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveDomainRefererRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveDomainRefererResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveDomainRefererResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveDomainRefererResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayAuthKeyRequest struct {
	*tchttp.BaseRequest

	// 播放域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	// 不传表示不修改当前值。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 鉴权key。
	// 不传表示不修改当前值。
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`

	// 有效时间，单位：秒。
	// 不传表示不修改当前值。
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`

	// 鉴权备用key。
	// 不传表示不修改当前值。
	AuthBackKey *string `json:"AuthBackKey,omitempty" name:"AuthBackKey"`
}

func (r *ModifyLivePlayAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLivePlayAuthKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "Enable")
	delete(f, "AuthKey")
	delete(f, "AuthDelta")
	delete(f, "AuthBackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLivePlayAuthKeyRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLivePlayAuthKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayDomainRequest struct {
	*tchttp.BaseRequest

	// 播放域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 拉流域名类型。1-国内；2-全球；3-境外
	PlayType *int64 `json:"PlayType,omitempty" name:"PlayType"`
}

func (r *ModifyLivePlayDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLivePlayDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "PlayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLivePlayDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePlayDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLivePlayDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLivePlayDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePullStreamTaskRequest struct {
	*tchttp.BaseRequest

	// 任务Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 操作人姓名。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 拉流源url列表。
	// SourceType为直播（PullLivePushLive）只可以填1个，
	// SourceType为点播（PullVodPushLive）可以填多个，上限30个。
	SourceUrls []*string `json:"SourceUrls,omitempty" name:"SourceUrls"`

	// 开始时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于七天。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 点播拉流转推循环次数。
	// -1：无限循环，直到任务结束。
	// 0：不循环。
	// >0：具体循环次数。次数和时间以先结束的为准。
	// 注意：拉流源为点播，该配置生效。
	VodLoopTimes *int64 `json:"VodLoopTimes,omitempty" name:"VodLoopTimes"`

	// 点播更新SourceUrls后的播放方式：
	// ImmediateNewSource：立即从更新的拉流源开始播放；
	// ContinueBreakPoint：从上次断流url源的断点处继续，结束后再使用新的拉流源。
	// 注意：拉流源为点播，该配置生效。
	VodRefreshType *string `json:"VodRefreshType,omitempty" name:"VodRefreshType"`

	// 任务状态：
	// enable - 启用，
	// pause - 暂停。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 选择需要回调的事件（不填则回调全部）：
	// TaskStart：任务启动回调，
	// TaskExit：任务停止回调，
	// VodSourceFileStart：从点播源文件开始拉流回调，
	// VodSourceFileFinish：从点播源文件拉流结束回调，
	// ResetTaskConfig：任务更新回调。
	CallbackEvents []*string `json:"CallbackEvents,omitempty" name:"CallbackEvents"`

	// 自定义回调地址。
	// 相关事件会回调到该地址。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 指定播放文件索引。
	// 注意： 从1开始，不大于SourceUrls中文件个数。
	FileIndex *int64 `json:"FileIndex,omitempty" name:"FileIndex"`

	// 指定播放文件偏移。
	// 注意：
	// 1. 单位：秒，配合FileIndex使用。
	OffsetTime *int64 `json:"OffsetTime,omitempty" name:"OffsetTime"`

	// 任务备注。
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *ModifyLivePullStreamTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLivePullStreamTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Operator")
	delete(f, "SourceUrls")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "VodLoopTimes")
	delete(f, "VodRefreshType")
	delete(f, "Status")
	delete(f, "CallbackEvents")
	delete(f, "CallbackUrl")
	delete(f, "FileIndex")
	delete(f, "OffsetTime")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLivePullStreamTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePullStreamTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLivePullStreamTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLivePullStreamTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLivePushAuthKeyRequest struct {
	*tchttp.BaseRequest

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	// 不传表示不修改当前值。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 主鉴权key。
	// 不传表示不修改当前值。
	MasterAuthKey *string `json:"MasterAuthKey,omitempty" name:"MasterAuthKey"`

	// 备鉴权key。
	// 不传表示不修改当前值。
	BackupAuthKey *string `json:"BackupAuthKey,omitempty" name:"BackupAuthKey"`

	// 有效时间，单位：秒。
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`
}

func (r *ModifyLivePushAuthKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLivePushAuthKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "Enable")
	delete(f, "MasterAuthKey")
	delete(f, "BackupAuthKey")
	delete(f, "AuthDelta")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLivePushAuthKeyRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLivePushAuthKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest

	// DescribeRecordTemplates接口获取到的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// FLV 录制参数，开启 FLV 录制时设置。
	FlvParam *RecordParam `json:"FlvParam,omitempty" name:"FlvParam"`

	// HLS 录制参数，开启 HLS 录制时设置。
	HlsParam *RecordParam `json:"HlsParam,omitempty" name:"HlsParam"`

	// MP4 录制参数，开启 MP4 录制时设置。
	Mp4Param *RecordParam `json:"Mp4Param,omitempty" name:"Mp4Param"`

	// AAC 录制参数，开启 AAC 录制时设置。
	AacParam *RecordParam `json:"AacParam,omitempty" name:"AacParam"`

	// HLS 录制定制参数。
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitempty" name:"HlsSpecialParam"`

	// MP3 录制参数，开启 MP3 录制时设置。
	Mp3Param *RecordParam `json:"Mp3Param,omitempty" name:"Mp3Param"`
}

func (r *ModifyLiveRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateName")
	delete(f, "Description")
	delete(f, "FlvParam")
	delete(f, "HlsParam")
	delete(f, "Mp4Param")
	delete(f, "AacParam")
	delete(f, "HlsSpecialParam")
	delete(f, "Mp3Param")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveRecordTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	// 长度上限：255字节。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	// 长度上限：1024字节。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 截图间隔，单位s，默认10s。
	// 范围： 5s ~ 300s。
	SnapshotInterval *int64 `json:"SnapshotInterval,omitempty" name:"SnapshotInterval"`

	// 截图宽度。默认：0（原始宽）。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 截图高度。默认：0（原始高）。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 是否开启鉴黄，默认 0 。
	// 0：不开启。
	// 1：开启。
	PornFlag *int64 `json:"PornFlag,omitempty" name:"PornFlag"`

	// Cos 应用 ID。
	CosAppId *int64 `json:"CosAppId,omitempty" name:"CosAppId"`

	// Cos Bucket名称。
	// 注：CosBucket参数值不能包含-[appid] 部分。
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// Cos 地域。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// Cos Bucket文件夹前缀。
	CosPrefix *string `json:"CosPrefix,omitempty" name:"CosPrefix"`

	// Cos 文件名称。
	CosFileName *string `json:"CosFileName,omitempty" name:"CosFileName"`
}

func (r *ModifyLiveSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateName")
	delete(f, "Description")
	delete(f, "SnapshotInterval")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "PornFlag")
	delete(f, "CosAppId")
	delete(f, "CosBucket")
	delete(f, "CosRegion")
	delete(f, "CosPrefix")
	delete(f, "CosFileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveSnapshotTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板 Id。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 视频编码：h264/h265/origin，默认origin。
	// 
	// origin: 保持原始编码格式
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// 音频编码：aac，默认aac。
	// 注意：当前该参数未生效，待后续支持！
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// 音频码率，默认0。
	// 范围：0-500。
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// 模板描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 视频码率。范围：0kbps - 8000kbps。
	// 0为保持原始码率。
	// 注: 转码模板有码率唯一要求，最终保存的码率可能与输入码率有所差别。
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// 宽。0-3000。
	// 数值必须是2的倍数，0是原始宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 是否保留视频，0：否，1：是。默认1。
	NeedVideo *int64 `json:"NeedVideo,omitempty" name:"NeedVideo"`

	// 是否保留音频，0：否，1：是。默认1。
	NeedAudio *int64 `json:"NeedAudio,omitempty" name:"NeedAudio"`

	// 高。0-3000。
	// 数值必须是2的倍数，0是原始宽度
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 帧率，默认0。
	// 范围0-60
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 关键帧间隔，单位：秒。
	// 范围2-6
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// 旋转角度，默认0。
	// 可取值：0，90，180，270
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 编码质量：
	// baseline/main/high。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 当设置的码率>原始码率时，是否以原始码率为准。
	// 0：否， 1：是
	// 默认 0。
	BitrateToOrig *int64 `json:"BitrateToOrig,omitempty" name:"BitrateToOrig"`

	// 当设置的高度>原始高度时，是否以原始高度为准。
	// 0：否， 1：是
	// 默认 0。
	HeightToOrig *int64 `json:"HeightToOrig,omitempty" name:"HeightToOrig"`

	// 当设置的帧率>原始帧率时，是否以原始帧率为准。
	// 0：否， 1：是
	// 默认 0。
	FpsToOrig *int64 `json:"FpsToOrig,omitempty" name:"FpsToOrig"`

	// 极速高清视频码率压缩比。
	// 极速高清目标码率=VideoBitrate * (1-AdaptBitratePercent)
	// 
	// 取值范围：0.0到0.5
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitempty" name:"AdaptBitratePercent"`

	// 是否以短边作为高度，0：否，1：是。默认0。
	ShortEdgeAsHeight *int64 `json:"ShortEdgeAsHeight,omitempty" name:"ShortEdgeAsHeight"`
}

func (r *ModifyLiveTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveTranscodeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Vcodec")
	delete(f, "Acodec")
	delete(f, "AudioBitrate")
	delete(f, "Description")
	delete(f, "VideoBitrate")
	delete(f, "Width")
	delete(f, "NeedVideo")
	delete(f, "NeedAudio")
	delete(f, "Height")
	delete(f, "Fps")
	delete(f, "Gop")
	delete(f, "Rotate")
	delete(f, "Profile")
	delete(f, "BitrateToOrig")
	delete(f, "HeightToOrig")
	delete(f, "FpsToOrig")
	delete(f, "AdaptBitratePercent")
	delete(f, "ShortEdgeAsHeight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveTranscodeTemplateRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveTranscodeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPullStreamConfigRequest struct {
	*tchttp.BaseRequest

	// 配置 ID。
	// 获取来源：
	// 1. 创建拉流配置接口CreatePullStreamConfig返回的配置 ID。
	// 2. 通过查询接口DescribePullStreamConfigs获取配置 ID。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 源 URL，用于拉流的地址。目前可支持直播流及点播文件。
	// 注意：
	// 1. 多个点播 URL 之间使用空格拼接。
	// 2. 目前上限支持10个 URL。
	// 3. 支持拉流文件格式：FLV，RTMP，HLS，MP4。
	// 4. 使用标准三层样式，如：http://test.com/live/stream.flv。
	FromUrl *string `json:"FromUrl,omitempty" name:"FromUrl"`

	// 目的 URL，用于推流的地址，目前限制该目标地址为腾讯域名。
	// 1. 仅支持 RTMP 协议。
	// 2. 使用标准三层样式，如：http://test.com/live/stream.flv。
	ToUrl *string `json:"ToUrl,omitempty" name:"ToUrl"`

	// 区域 ID：
	// 1-深圳。
	// 2-上海。
	// 3-天津。
	// 4-中国香港。
	// 如有改动，需同时传入IspId。
	AreaId *int64 `json:"AreaId,omitempty" name:"AreaId"`

	// 运营商 ID，
	// 1：电信。
	// 2：移动。
	// 3：联通。
	// 4：其他。
	// AreaId为4的时候，IspId只能为其他。如有改动，需同时传入AreaId。
	IspId *int64 `json:"IspId,omitempty" name:"IspId"`

	// 开始时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于七天。
	// 
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ModifyPullStreamConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPullStreamConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "FromUrl")
	delete(f, "ToUrl")
	delete(f, "AreaId")
	delete(f, "IspId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPullStreamConfigRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPullStreamConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPullStreamStatusRequest struct {
	*tchttp.BaseRequest

	// 配置 ID 列表。
	ConfigIds []*string `json:"ConfigIds,omitempty" name:"ConfigIds"`

	// 目标状态。0无效，2正在运行，4暂停。
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyPullStreamStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPullStreamStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigIds")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPullStreamStatusRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPullStreamStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorStreamPlayInfo struct {

	// 播放域名。
	PlayDomain *string `json:"PlayDomain,omitempty" name:"PlayDomain"`

	// 流id。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 播放码率，0表示原始码率。
	Rate *uint64 `json:"Rate,omitempty" name:"Rate"`

	// 播放协议，可选值包括 Unknown，Flv，Hls，Rtmp，Huyap2p。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 带宽，单位是Mbps。
	Bandwidth *float64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 在线人数，1分钟采样一个点，统计采样点的tcp链接数目。
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// 请求数。
	Request *uint64 `json:"Request,omitempty" name:"Request"`
}

type PlayAuthKeyInfo struct {

	// 域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 是否启用:
	// 0: 关闭。
	// 1: 启用。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 鉴权 Key。
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`

	// 有效时间，单位：秒。
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`

	// 鉴权 BackKey。
	AuthBackKey *string `json:"AuthBackKey,omitempty" name:"AuthBackKey"`
}

type PlayCodeTotalInfo struct {

	// HTTP code，可选值包括:
	// 400，403，404，500，502，503，504。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 总次数。
	Num *uint64 `json:"Num,omitempty" name:"Num"`
}

type PlayDataInfoByStream struct {

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 总流量，单位: MB。
	TotalFlux *float64 `json:"TotalFlux,omitempty" name:"TotalFlux"`
}

type PlayStatInfo struct {

	// 数据时间点。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 带宽/流量/请求数/并发连接数/下载速度的值，若没数据返回时该值为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type PlaySumStatInfo struct {

	// 域名或流 ID。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 平均下载速度，
	// 单位: MB/s。
	// 计算公式: 每分钟的下载速度求平均值。
	AvgFluxPerSecond *float64 `json:"AvgFluxPerSecond,omitempty" name:"AvgFluxPerSecond"`

	// 总流量，单位: MB。
	TotalFlux *float64 `json:"TotalFlux,omitempty" name:"TotalFlux"`

	// 总请求数。
	TotalRequest *uint64 `json:"TotalRequest,omitempty" name:"TotalRequest"`
}

type ProIspPlayCodeDataInfo struct {

	// 国家或地区。
	CountryAreaName *string `json:"CountryAreaName,omitempty" name:"CountryAreaName"`

	// 省份。
	ProvinceName *string `json:"ProvinceName,omitempty" name:"ProvinceName"`

	// 运营商。
	IspName *string `json:"IspName,omitempty" name:"IspName"`

	// 错误码为2开头的次数。
	Code2xx *uint64 `json:"Code2xx,omitempty" name:"Code2xx"`

	// 错误码为3开头的次数。
	Code3xx *uint64 `json:"Code3xx,omitempty" name:"Code3xx"`

	// 错误码为4开头的次数。
	Code4xx *uint64 `json:"Code4xx,omitempty" name:"Code4xx"`

	// 错误码为5开头的次数。
	Code5xx *uint64 `json:"Code5xx,omitempty" name:"Code5xx"`
}

type ProIspPlaySumInfo struct {

	// 省份/运营商/国家或地区。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 总流量，单位: MB。
	TotalFlux *float64 `json:"TotalFlux,omitempty" name:"TotalFlux"`

	// 总请求数。
	TotalRequest *uint64 `json:"TotalRequest,omitempty" name:"TotalRequest"`

	// 平均下载流量，单位: MB/s。
	AvgFluxPerSecond *float64 `json:"AvgFluxPerSecond,omitempty" name:"AvgFluxPerSecond"`
}

type PublishTime struct {

	// 推流时间。
	// UTC 格式，例如：2018-06-29T19:00:00Z。
	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
}

type PullStreamConfig struct {

	// 拉流配置 ID。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 源 URL。
	FromUrl *string `json:"FromUrl,omitempty" name:"FromUrl"`

	// 目的 URL。
	ToUrl *string `json:"ToUrl,omitempty" name:"ToUrl"`

	// 区域名。
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`

	// 运营商名。
	IspName *string `json:"IspName,omitempty" name:"IspName"`

	// 开始时间。
	// UTC格式时间，例如: 2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	// 
	// UTC格式时间，例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 状态:
	// 0: 无效。
	// 1: 初始状态。
	// 2: 正在运行。
	// 3: 拉起失败。
	// 4: 暂停。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type PullStreamTaskInfo struct {

	// 拉流任务Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 拉流源的类型：
	// PullLivePushLive -直播，
	// PullVodPushLive -点播。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 拉流源url列表。
	// SourceType为直播（PullLiveToLive）只可以填1个，
	// SourceType为点播（PullVodToLive）可以填多个，上限10个。
	SourceUrls []*string `json:"SourceUrls,omitempty" name:"SourceUrls"`

	// 推流域名。
	// 将拉到的源推到该域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	// 将拉到的源推到该路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	// 将拉到的源推到该流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 推流参数。
	// 推流携带的自定义参数。
	PushArgs *string `json:"PushArgs,omitempty" name:"PushArgs"`

	// 开始时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于七天。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 拉流源所在地域（请就近选取）：
	// ap-beijing - 华北地区(北京)，
	// ap-shanghai -华东地区(上海)，
	// ap-guangzhou -华南地区(广州)，
	// ap-mumbai - 印度。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 点播拉流转推循环次数。
	// -1：无限循环，直到任务结束。
	// 0：不循环。
	// >0：具体循环次数。次数和时间以先结束的为准。
	// 注意：拉流源为点播，该配置生效。
	VodLoopTimes *int64 `json:"VodLoopTimes,omitempty" name:"VodLoopTimes"`

	// 点播更新SourceUrls后的播放方式：
	// ImmediateNewSource：立即从更新的拉流源开始播放；
	// ContinueBreakPoint：从上次断流url源的断点处继续，结束后再使用新的拉流源。
	// 
	// 注意：拉流源为点播，该配置生效。
	VodRefreshType *string `json:"VodRefreshType,omitempty" name:"VodRefreshType"`

	// 任务创建时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务更新时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建任务的操作者。
	CreateBy *string `json:"CreateBy,omitempty" name:"CreateBy"`

	// 最后更新任务的操作者。
	UpdateBy *string `json:"UpdateBy,omitempty" name:"UpdateBy"`

	// 回调地址。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 选择需要回调的事件：
	// TaskStart：任务启动回调，
	// TaskExit：任务停止回调，
	// VodSourceFileStart：从点播源文件开始拉流回调，
	// VodSourceFileFinish：从点播源文件拉流结束回调，
	// ResetTaskConfig：任务更新回调。
	CallbackEvents []*string `json:"CallbackEvents,omitempty" name:"CallbackEvents"`

	// 注意：该信息暂不返回。
	// 最后一次回调信息。
	CallbackInfo *string `json:"CallbackInfo,omitempty" name:"CallbackInfo"`

	// 注意：该信息暂不返回。
	// 错误信息。
	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`

	// 状态。
	// enable：生效中。
	// pause：暂停中。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 注意：该信息仅在查询单个任务时返回。
	// 任务最新拉流信息。
	// 包含：源 url，偏移时间，上报时间。
	RecentPullInfo *RecentPullInfo `json:"RecentPullInfo,omitempty" name:"RecentPullInfo"`

	// 任务备注信息。
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type PushAuthKeyInfo struct {

	// 域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 主鉴权 Key。
	MasterAuthKey *string `json:"MasterAuthKey,omitempty" name:"MasterAuthKey"`

	// 备鉴权 Key。
	BackupAuthKey *string `json:"BackupAuthKey,omitempty" name:"BackupAuthKey"`

	// 有效时间，单位：秒。
	AuthDelta *uint64 `json:"AuthDelta,omitempty" name:"AuthDelta"`
}

type PushDataInfo struct {

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 推流客户端 IP。
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 接流服务器 IP。
	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`

	// 推流视频帧率，单位: Hz。
	VideoFps *uint64 `json:"VideoFps,omitempty" name:"VideoFps"`

	// 推流视频码率，单位: Kbps。
	VideoSpeed *uint64 `json:"VideoSpeed,omitempty" name:"VideoSpeed"`

	// 推流音频帧率，单位: Hz。
	AudioFps *uint64 `json:"AudioFps,omitempty" name:"AudioFps"`

	// 推流音频码率，单位: Kbps。
	AudioSpeed *uint64 `json:"AudioSpeed,omitempty" name:"AudioSpeed"`

	// 推流域名。
	PushDomain *string `json:"PushDomain,omitempty" name:"PushDomain"`

	// 推流开始时间。
	BeginPushTime *string `json:"BeginPushTime,omitempty" name:"BeginPushTime"`

	// 音频编码格式，
	// 例："AAC"。
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// 视频编码格式，
	// 例："H264"。
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// 分辨率。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// 采样率。
	AsampleRate *uint64 `json:"AsampleRate,omitempty" name:"AsampleRate"`

	// metadata 中的音频码率，单位: Kbps。
	MetaAudioSpeed *uint64 `json:"MetaAudioSpeed,omitempty" name:"MetaAudioSpeed"`

	// metadata 中的视频码率，单位: Kbps。
	MetaVideoSpeed *uint64 `json:"MetaVideoSpeed,omitempty" name:"MetaVideoSpeed"`

	// metadata 中的帧率。
	MetaFps *uint64 `json:"MetaFps,omitempty" name:"MetaFps"`
}

type PushQualityData struct {

	// 数据时间，格式: %Y-%m-%d %H:%M:%S.%ms，精确到毫秒级。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 推流域名。
	PushDomain *string `json:"PushDomain,omitempty" name:"PushDomain"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 推流客户端 IP。
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 开始推流时间，格式: %Y-%m-%d %H:%M:%S.%ms，精确到毫秒级。
	BeginPushTime *string `json:"BeginPushTime,omitempty" name:"BeginPushTime"`

	// 分辨率信息。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// 视频编码格式。
	VCodec *string `json:"VCodec,omitempty" name:"VCodec"`

	// 音频编码格式。
	ACodec *string `json:"ACodec,omitempty" name:"ACodec"`

	// 推流序列号，用来唯一的标志一次推流。
	Sequence *string `json:"Sequence,omitempty" name:"Sequence"`

	// 视频帧率。
	VideoFps *uint64 `json:"VideoFps,omitempty" name:"VideoFps"`

	// 视频码率，单位: bps。
	VideoRate *uint64 `json:"VideoRate,omitempty" name:"VideoRate"`

	// 音频帧率。
	AudioFps *uint64 `json:"AudioFps,omitempty" name:"AudioFps"`

	// 音频码率，单位: bps。
	AudioRate *uint64 `json:"AudioRate,omitempty" name:"AudioRate"`

	// 本地流逝时间，单位: ms，音视频流逝时间与本地流逝时间的差距越大表示推流质量越差，上行卡顿越严重。
	LocalTs *uint64 `json:"LocalTs,omitempty" name:"LocalTs"`

	// 视频流逝时间，单位: ms。
	VideoTs *uint64 `json:"VideoTs,omitempty" name:"VideoTs"`

	// 音频流逝时间，单位: ms。
	AudioTs *uint64 `json:"AudioTs,omitempty" name:"AudioTs"`

	// metadata 中的视频码率，单位: kbps。
	MetaVideoRate *uint64 `json:"MetaVideoRate,omitempty" name:"MetaVideoRate"`

	// metadata 中的音频码率，单位: kbps。
	MetaAudioRate *uint64 `json:"MetaAudioRate,omitempty" name:"MetaAudioRate"`

	// metadata 中的帧率。
	MateFps *uint64 `json:"MateFps,omitempty" name:"MateFps"`

	// 推流参数
	StreamParam *string `json:"StreamParam,omitempty" name:"StreamParam"`
}

type RecentPullInfo struct {

	// 当前正在拉的文件地址。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 当前正在拉的文件偏移，单位：秒。
	OffsetTime *uint64 `json:"OffsetTime,omitempty" name:"OffsetTime"`

	// 最新上报偏移信息时间。UTC格式。
	// 如：2020-07-23T03:20:39Z。
	// 注意：与北京时间相差八小时。
	ReportTime *string `json:"ReportTime,omitempty" name:"ReportTime"`

	// 已经轮播的次数。
	LoopedTimes *int64 `json:"LoopedTimes,omitempty" name:"LoopedTimes"`
}

type RecordParam struct {

	// 录制间隔。
	// 单位秒，默认：1800。
	// 取值范围：30-7200。
	// 此参数对 HLS 无效，当录制 HLS 时从推流到断流生成一个文件。
	RecordInterval *int64 `json:"RecordInterval,omitempty" name:"RecordInterval"`

	// 录制存储时长。
	// 单位秒，取值范围： 0 - 1500天。
	// 0：表示永久存储。
	StorageTime *int64 `json:"StorageTime,omitempty" name:"StorageTime"`

	// 是否开启当前格式录制，默认值为0，0：否， 1：是。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 点播子应用 ID。
	VodSubAppId *int64 `json:"VodSubAppId,omitempty" name:"VodSubAppId"`

	// 录制文件名。
	// 支持的特殊占位符有：
	// {StreamID}: 流ID
	// {StartYear}: 开始时间-年
	// {StartMonth}: 开始时间-月
	// {StartDay}: 开始时间-日
	// {StartHour}: 开始时间-小时
	// {StartMinute}: 开始时间-分钟
	// {StartSecond}: 开始时间-秒
	// {StartMillisecond}: 开始时间-毫秒
	// {EndYear}: 结束时间-年
	// {EndMonth}: 结束时间-月
	// {EndDay}: 结束时间-日
	// {EndHour}: 结束时间-小时
	// {EndMinute}: 结束时间-分钟
	// {EndSecond}: 结束时间-秒
	// {EndMillisecond}: 结束时间-毫秒
	// 
	// 若未设置默认录制文件名为{StreamID}_{StartYear}-{StartMonth}-{StartDay}-{StartHour}-{StartMinute}-{StartSecond}_{EndYear}-{EndMonth}-{EndDay}-{EndHour}-{EndMinute}-{EndSecond}
	VodFileName *string `json:"VodFileName,omitempty" name:"VodFileName"`

	// 任务流
	// 注意：此字段可能返回 null，表示取不到有效值。
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// 视频存储策略。
	// normal：标准存储。
	// cold：低频存储。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageMode *string `json:"StorageMode,omitempty" name:"StorageMode"`

	// 点播应用分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`
}

type RecordTask struct {

	// 录制任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 任务开始时间，Unix时间戳。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 任务结束时间，Unix时间戳。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 录制模板ID。
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 调用 StopRecordTask 停止任务时间，Unix时间戳。值为0表示未曾调用接口停止任务。
	Stopped *uint64 `json:"Stopped,omitempty" name:"Stopped"`
}

type RecordTemplateInfo struct {

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// FLV 录制参数。
	FlvParam *RecordParam `json:"FlvParam,omitempty" name:"FlvParam"`

	// HLS 录制参数。
	HlsParam *RecordParam `json:"HlsParam,omitempty" name:"HlsParam"`

	// MP4 录制参数。
	Mp4Param *RecordParam `json:"Mp4Param,omitempty" name:"Mp4Param"`

	// AAC 录制参数。
	AacParam *RecordParam `json:"AacParam,omitempty" name:"AacParam"`

	// 0：普通直播，
	// 1：慢直播。
	IsDelayLive *int64 `json:"IsDelayLive,omitempty" name:"IsDelayLive"`

	// HLS 录制定制参数
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitempty" name:"HlsSpecialParam"`

	// MP3 录制参数。
	Mp3Param *RecordParam `json:"Mp3Param,omitempty" name:"Mp3Param"`
}

type RefererAuthConfig struct {

	// 域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 名单类型，0：黑名单，1：白名单。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 是否允许空Referer，0：不允许，1：允许。
	AllowEmpty *int64 `json:"AllowEmpty,omitempty" name:"AllowEmpty"`

	// 名单列表，以分号(;)分隔。
	Rules *string `json:"Rules,omitempty" name:"Rules"`
}

type ResumeDelayLiveStreamRequest struct {
	*tchttp.BaseRequest

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *ResumeDelayLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeDelayLiveStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "DomainName")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeDelayLiveStreamRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeDelayLiveStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeLiveStreamRequest struct {
	*tchttp.BaseRequest

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *ResumeLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeLiveStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "DomainName")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeLiveStreamRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeLiveStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleInfo struct {

	// 规则创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 规则更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

type SnapshotTemplateInfo struct {

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 截图时间间隔，5-300秒。
	SnapshotInterval *int64 `json:"SnapshotInterval,omitempty" name:"SnapshotInterval"`

	// 截图宽度，范围：0-3000。 
	// 0：原始宽度并适配原始比例。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 截图高度，范围：0-2000。
	// 0：原始高度并适配原始比例。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 是否开启鉴黄，0：不开启，1：开启。
	PornFlag *int64 `json:"PornFlag,omitempty" name:"PornFlag"`

	// Cos 应用 ID。
	CosAppId *int64 `json:"CosAppId,omitempty" name:"CosAppId"`

	// Cos Bucket名称。
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// Cos 地域。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 模板描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// Cos Bucket文件夹前缀。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPrefix *string `json:"CosPrefix,omitempty" name:"CosPrefix"`

	// Cos 文件名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosFileName *string `json:"CosFileName,omitempty" name:"CosFileName"`
}

type StopLiveRecordRequest struct {
	*tchttp.BaseRequest

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 任务ID，由CreateLiveRecord接口返回。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopLiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLiveRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StreamName")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopLiveRecordRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLiveRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopRecordTaskRequest struct {
	*tchttp.BaseRequest

	// 录制任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopRecordTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRecordTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopRecordTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopRecordTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopRecordTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRecordTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StreamEventInfo struct {

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 推流开始时间。
	// UTC 格式时间，例如：2019-01-07T12:00:00Z。
	StreamStartTime *string `json:"StreamStartTime,omitempty" name:"StreamStartTime"`

	// 推流结束时间。
	// UTC 格式时间，例如：2019-01-07T15:00:00Z。
	StreamEndTime *string `json:"StreamEndTime,omitempty" name:"StreamEndTime"`

	// 停止原因。
	StopReason *string `json:"StopReason,omitempty" name:"StopReason"`

	// 推流持续时长，单位：秒。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 主播 IP。
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 分辨率。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`
}

type StreamName struct {

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 推流开始时间。
	// UTC格式时间，例如：2019-01-07T12:00:00Z。
	StreamStartTime *string `json:"StreamStartTime,omitempty" name:"StreamStartTime"`

	// 推流结束时间。
	// UTC格式时间，例如：2019-01-07T15:00:00Z。
	StreamEndTime *string `json:"StreamEndTime,omitempty" name:"StreamEndTime"`

	// 停止原因。
	StopReason *string `json:"StopReason,omitempty" name:"StopReason"`

	// 推流持续时长，单位：秒。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 主播 IP。
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 分辨率。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`
}

type StreamOnlineInfo struct {

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 推流时间列表
	PublishTimeList []*PublishTime `json:"PublishTimeList,omitempty" name:"PublishTimeList"`

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

type TemplateInfo struct {

	// 视频编码：h264/h265/origin，默认h264。
	// 
	// origin: 保持原始编码格式
	Vcodec *string `json:"Vcodec,omitempty" name:"Vcodec"`

	// 视频码率。范围：0kbps - 8000kbps。
	// 0为保持原始码率。
	// 注: 转码模板有码率唯一要求，最终保存的码率可能与输入码率有所差别。
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// 音频编码：aac，默认aac。
	// 注意：当前该参数未生效，待后续支持！
	Acodec *string `json:"Acodec,omitempty" name:"Acodec"`

	// 音频码率。取值范围：0kbps - 500kbps。
	// 默认0。
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// 宽，默认0。
	// 范围[0-3000]
	// 数值必须是2的倍数，0是原始宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高，默认0。
	// 范围[0-3000]
	// 数值必须是2的倍数，0是原始宽度
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 帧率，默认0。
	// 范围0-60fps
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 关键帧间隔，单位：秒。
	// 默认原始的间隔
	// 范围2-6
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// 旋转角度，默认0。
	// 可取值：0，90，180，270
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 编码质量：
	// baseline/main/high。默认baseline
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 当设置的码率>原始码率时，是否以原始码率为准。
	// 0：否， 1：是
	// 默认 0。
	BitrateToOrig *int64 `json:"BitrateToOrig,omitempty" name:"BitrateToOrig"`

	// 当设置的高度>原始高度时，是否以原始高度为准。
	// 0：否， 1：是
	// 默认 0。
	HeightToOrig *int64 `json:"HeightToOrig,omitempty" name:"HeightToOrig"`

	// 当设置的帧率>原始帧率时，是否以原始帧率为准。
	// 0：否， 1：是
	// 默认 0。
	FpsToOrig *int64 `json:"FpsToOrig,omitempty" name:"FpsToOrig"`

	// 是否保留视频。0：否，1：是。
	NeedVideo *int64 `json:"NeedVideo,omitempty" name:"NeedVideo"`

	// 是否保留音频。0：否，1：是。
	NeedAudio *int64 `json:"NeedAudio,omitempty" name:"NeedAudio"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模板描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否是极速高清模板，0：否，1：是。默认0。
	AiTransCode *int64 `json:"AiTransCode,omitempty" name:"AiTransCode"`

	// 极速高清视频码率压缩比。
	// 极速高清目标码率=VideoBitrate * (1-AdaptBitratePercent)
	// 
	// 取值范围：0.0到0.5
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitempty" name:"AdaptBitratePercent"`

	// 是否以短边作为高度，0：否，1：是。默认0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShortEdgeAsHeight *int64 `json:"ShortEdgeAsHeight,omitempty" name:"ShortEdgeAsHeight"`
}

type TimeValue struct {

	// UTC 时间，时间格式：yyyy-mm-ddTHH:MM:SSZ。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 数值。
	Num *uint64 `json:"Num,omitempty" name:"Num"`
}

type TranscodeDetailInfo struct {

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 开始时间（北京时间），格式：yyyy-mm-dd HH:MM。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（北京时间），格式：yyyy-mm-dd HH:MM。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 转码时长，单位：分钟。
	// 注意：因推流过程中可能有中断重推情况，此处时长为真实转码时长累加值，并非结束时间和开始时间的间隔。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 编码方式，带模块，
	// 示例：
	// liveprocessor_H264：直播转码-H264，
	// liveprocessor_H265： 直播转码-H265，
	// topspeed_H264：极速高清-H264，
	// topspeed_H265：极速高清-H265。
	ModuleCodec *string `json:"ModuleCodec,omitempty" name:"ModuleCodec"`

	// 码率。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 类型，包含：转码（Transcode），混流（MixStream），水印（WaterMark）。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 推流域名。
	PushDomain *string `json:"PushDomain,omitempty" name:"PushDomain"`

	// 分辨率。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`
}

type TranscodeTotalInfo struct {

	// 时间点，北京时间，
	// 示例：2019-03-01 00:00:00。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 转码时长，单位：分钟。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 编码方式，带模块，
	// 示例：
	// liveprocessor_H264 =》直播转码-H264，
	// liveprocessor_H265 =》 直播转码-H265，
	// topspeed_H264 =》极速高清-H264，
	// topspeed_H265 =》极速高清-H265。
	ModuleCodec *string `json:"ModuleCodec,omitempty" name:"ModuleCodec"`

	// 分辨率，
	// 示例：540*480。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`
}

type UnBindLiveDomainCertRequest struct {
	*tchttp.BaseRequest

	// 播放域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 枚举值：
	// gray: 解绑灰度规则
	// formal(默认): 解绑正式规则
	// 
	// 不传则为formal
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *UnBindLiveDomainCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindLiveDomainCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindLiveDomainCertRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindLiveDomainCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLiveWatermarkRequest struct {
	*tchttp.BaseRequest

	// 水印 ID。
	// 在添加水印接口 [AddLiveWatermark](/document/product/267/30154) 调用返回值中获取水印 ID。
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`

	// 水印图片 URL。
	// URL中禁止包含的字符：
	//  ;(){}$>`#"\'|
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// 显示位置，X轴偏移，单位是百分比，默认 0。
	XPosition *int64 `json:"XPosition,omitempty" name:"XPosition"`

	// 显示位置，Y轴偏移，单位是百分比，默认 0。
	YPosition *int64 `json:"YPosition,omitempty" name:"YPosition"`

	// 水印名称。
	// 最长16字节。
	WatermarkName *string `json:"WatermarkName,omitempty" name:"WatermarkName"`

	// 水印宽度，占直播原始画面宽度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 水印高度，占直播原始画面宽度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

func (r *UpdateLiveWatermarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateLiveWatermarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WatermarkId")
	delete(f, "PictureUrl")
	delete(f, "XPosition")
	delete(f, "YPosition")
	delete(f, "WatermarkName")
	delete(f, "Width")
	delete(f, "Height")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateLiveWatermarkRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateLiveWatermarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WatermarkInfo struct {

	// 水印 ID。
	WatermarkId *int64 `json:"WatermarkId,omitempty" name:"WatermarkId"`

	// 水印图片 URL。
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// 显示位置，X 轴偏移。
	XPosition *int64 `json:"XPosition,omitempty" name:"XPosition"`

	// 显示位置，Y 轴偏移。
	YPosition *int64 `json:"YPosition,omitempty" name:"YPosition"`

	// 水印名称。
	WatermarkName *string `json:"WatermarkName,omitempty" name:"WatermarkName"`

	// 当前状态。0：未使用，1:使用中。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 添加时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 水印宽。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 水印高。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}
