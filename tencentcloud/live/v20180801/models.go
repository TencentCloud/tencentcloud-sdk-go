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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddCasterInputInfoRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台输入源详细信息。
	InputInfo *CasterInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`
}

type AddCasterInputInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台输入源详细信息。
	InputInfo *CasterInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`
}

func (r *AddCasterInputInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCasterInputInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "InputInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCasterInputInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCasterInputInfoResponseParams struct {
	// rtmp协议输入源播放地址。
	// 注：仅可作为预览使用，不可分发。
	InputPlayUrl *string `json:"InputPlayUrl,omitnil,omitempty" name:"InputPlayUrl"`

	// webrtc协议播放地址。
	// 注：
	// 1. 需配合使用腾讯云快直播播放SDK使用才可正常播放。
	// 2. 仅作为预览使用，不可分发。
	InputWebRTCPlayUrl *string `json:"InputWebRTCPlayUrl,omitnil,omitempty" name:"InputWebRTCPlayUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddCasterInputInfoResponse struct {
	*tchttp.BaseResponse
	Response *AddCasterInputInfoResponseParams `json:"Response"`
}

func (r *AddCasterInputInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCasterInputInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCasterLayoutInfoRequestParams struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台布局参数信息。
	LayoutInfo *CasterLayoutInfo `json:"LayoutInfo,omitnil,omitempty" name:"LayoutInfo"`
}

type AddCasterLayoutInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台布局参数信息。
	LayoutInfo *CasterLayoutInfo `json:"LayoutInfo,omitnil,omitempty" name:"LayoutInfo"`
}

func (r *AddCasterLayoutInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCasterLayoutInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "LayoutInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCasterLayoutInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCasterLayoutInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddCasterLayoutInfoResponse struct {
	*tchttp.BaseResponse
	Response *AddCasterLayoutInfoResponseParams `json:"Response"`
}

func (r *AddCasterLayoutInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCasterLayoutInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCasterMarkPicInfoRequestParams struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 图片水印详细参数。
	MarkPicInfo *CasterMarkPicInfo `json:"MarkPicInfo,omitnil,omitempty" name:"MarkPicInfo"`
}

type AddCasterMarkPicInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 图片水印详细参数。
	MarkPicInfo *CasterMarkPicInfo `json:"MarkPicInfo,omitnil,omitempty" name:"MarkPicInfo"`
}

func (r *AddCasterMarkPicInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCasterMarkPicInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "MarkPicInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCasterMarkPicInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCasterMarkPicInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddCasterMarkPicInfoResponse struct {
	*tchttp.BaseResponse
	Response *AddCasterMarkPicInfoResponseParams `json:"Response"`
}

func (r *AddCasterMarkPicInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCasterMarkPicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCasterMarkWordInfoRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 文本的详细配置。
	MarkWordInfo *CasterMarkWordInfo `json:"MarkWordInfo,omitnil,omitempty" name:"MarkWordInfo"`
}

type AddCasterMarkWordInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 文本的详细配置。
	MarkWordInfo *CasterMarkWordInfo `json:"MarkWordInfo,omitnil,omitempty" name:"MarkWordInfo"`
}

func (r *AddCasterMarkWordInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCasterMarkWordInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "MarkWordInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCasterMarkWordInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCasterMarkWordInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddCasterMarkWordInfoResponse struct {
	*tchttp.BaseResponse
	Response *AddCasterMarkWordInfoResponseParams `json:"Response"`
}

func (r *AddCasterMarkWordInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCasterMarkWordInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCasterOutputInfoRequestParams struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台推流参数信息。
	OutputInfo *CasterOutputInfo `json:"OutputInfo,omitnil,omitempty" name:"OutputInfo"`
}

type AddCasterOutputInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台推流参数信息。
	OutputInfo *CasterOutputInfo `json:"OutputInfo,omitnil,omitempty" name:"OutputInfo"`
}

func (r *AddCasterOutputInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCasterOutputInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "OutputInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCasterOutputInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCasterOutputInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddCasterOutputInfoResponse struct {
	*tchttp.BaseResponse
	Response *AddCasterOutputInfoResponseParams `json:"Response"`
}

func (r *AddCasterOutputInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCasterOutputInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDelayLiveStreamRequestParams struct {
	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 延播时间，单位：秒，上限：600秒。
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 延播设置的过期时间。UTC 格式，例如：2018-11-29T19:00:00Z。
	// 注意：
	// 1. 默认7天后过期，且最长支持7天内生效。
	// 2. 北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type AddDelayLiveStreamRequest struct {
	*tchttp.BaseRequest
	
	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 延播时间，单位：秒，上限：600秒。
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 延播设置的过期时间。UTC 格式，例如：2018-11-29T19:00:00Z。
	// 注意：
	// 1. 默认7天后过期，且最长支持7天内生效。
	// 2. 北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#I)。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
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

// Predefined struct for user
type AddDelayLiveStreamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddDelayLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *AddDelayLiveStreamResponseParams `json:"Response"`
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

// Predefined struct for user
type AddLiveDomainRequestParams struct {
	// 域名名称。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 域名类型，
	// 0：推流域名，
	// 1：播放域名。
	DomainType *uint64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`

	// 拉流域名类型：
	// 1：国内，
	// 2：全球，
	// 3：境外。
	// 默认值：1。
	PlayType *uint64 `json:"PlayType,omitnil,omitempty" name:"PlayType"`

	// 是否是慢直播：
	// 0： 普通直播，
	// 1 ：慢直播 。
	// 默认值： 0。
	IsDelayLive *int64 `json:"IsDelayLive,omitnil,omitempty" name:"IsDelayLive"`

	// 是否是小程序直播：
	// 0： 标准直播，
	// 1 ：小程序直播 。
	// 默认值： 0。
	IsMiniProgramLive *int64 `json:"IsMiniProgramLive,omitnil,omitempty" name:"IsMiniProgramLive"`

	// 域名归属校验类型。
	// 可取值（与 AuthenticateDomainOwner 接口的 VerifyType 参数一致。）：
	// dnsCheck ：立即验证配置 dns 的解析记录是否与待验证内容一致，成功则保存记录。
	// fileCheck ：立即验证 web 文件是否与待验证内容一致，成功则保存记录。
	// dbCheck :  检查是否已经验证成功过。
	// 若不传默认为 dbCheck 。
	VerifyOwnerType *string `json:"VerifyOwnerType,omitnil,omitempty" name:"VerifyOwnerType"`
}

type AddLiveDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名名称。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 域名类型，
	// 0：推流域名，
	// 1：播放域名。
	DomainType *uint64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`

	// 拉流域名类型：
	// 1：国内，
	// 2：全球，
	// 3：境外。
	// 默认值：1。
	PlayType *uint64 `json:"PlayType,omitnil,omitempty" name:"PlayType"`

	// 是否是慢直播：
	// 0： 普通直播，
	// 1 ：慢直播 。
	// 默认值： 0。
	IsDelayLive *int64 `json:"IsDelayLive,omitnil,omitempty" name:"IsDelayLive"`

	// 是否是小程序直播：
	// 0： 标准直播，
	// 1 ：小程序直播 。
	// 默认值： 0。
	IsMiniProgramLive *int64 `json:"IsMiniProgramLive,omitnil,omitempty" name:"IsMiniProgramLive"`

	// 域名归属校验类型。
	// 可取值（与 AuthenticateDomainOwner 接口的 VerifyType 参数一致。）：
	// dnsCheck ：立即验证配置 dns 的解析记录是否与待验证内容一致，成功则保存记录。
	// fileCheck ：立即验证 web 文件是否与待验证内容一致，成功则保存记录。
	// dbCheck :  检查是否已经验证成功过。
	// 若不传默认为 dbCheck 。
	VerifyOwnerType *string `json:"VerifyOwnerType,omitnil,omitempty" name:"VerifyOwnerType"`
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
	delete(f, "VerifyOwnerType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddLiveDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddLiveDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *AddLiveDomainResponseParams `json:"Response"`
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

// Predefined struct for user
type AddLiveWatermarkRequestParams struct {
	// 水印图片 URL。
	// URL中禁止包含的字符：
	//  ;(){}$>`#"\'|
	PictureUrl *string `json:"PictureUrl,omitnil,omitempty" name:"PictureUrl"`

	// 水印名称。
	// 最长30字节。
	WatermarkName *string `json:"WatermarkName,omitnil,omitempty" name:"WatermarkName"`

	// 显示位置，X轴偏移，单位是百分比，默认 0。
	XPosition *int64 `json:"XPosition,omitnil,omitempty" name:"XPosition"`

	// 显示位置，Y轴偏移，单位是百分比，默认 0。
	YPosition *int64 `json:"YPosition,omitnil,omitempty" name:"YPosition"`

	// 水印宽度，占直播原始画面宽度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印高度，占直播原始画面高度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始高度。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 背景水印宽度。默认宽度1920。
	BackgroundWidth *int64 `json:"BackgroundWidth,omitnil,omitempty" name:"BackgroundWidth"`

	// 背景水印高度。默认高度1080。
	BackgroundHeight *int64 `json:"BackgroundHeight,omitnil,omitempty" name:"BackgroundHeight"`
}

type AddLiveWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 水印图片 URL。
	// URL中禁止包含的字符：
	//  ;(){}$>`#"\'|
	PictureUrl *string `json:"PictureUrl,omitnil,omitempty" name:"PictureUrl"`

	// 水印名称。
	// 最长30字节。
	WatermarkName *string `json:"WatermarkName,omitnil,omitempty" name:"WatermarkName"`

	// 显示位置，X轴偏移，单位是百分比，默认 0。
	XPosition *int64 `json:"XPosition,omitnil,omitempty" name:"XPosition"`

	// 显示位置，Y轴偏移，单位是百分比，默认 0。
	YPosition *int64 `json:"YPosition,omitnil,omitempty" name:"YPosition"`

	// 水印宽度，占直播原始画面宽度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印高度，占直播原始画面高度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始高度。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 背景水印宽度。默认宽度1920。
	BackgroundWidth *int64 `json:"BackgroundWidth,omitnil,omitempty" name:"BackgroundWidth"`

	// 背景水印高度。默认高度1080。
	BackgroundHeight *int64 `json:"BackgroundHeight,omitnil,omitempty" name:"BackgroundHeight"`
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
	delete(f, "BackgroundWidth")
	delete(f, "BackgroundHeight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddLiveWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddLiveWatermarkResponseParams struct {
	// 水印ID。
	WatermarkId *uint64 `json:"WatermarkId,omitnil,omitempty" name:"WatermarkId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *AddLiveWatermarkResponseParams `json:"Response"`
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

// Predefined struct for user
type AuthenticateDomainOwnerRequestParams struct {
	// 要验证的域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 验证类型。可取值：
	// dnsCheck ：立即验证配置 dns 的解析记录是否与待验证内容一致，成功则保存记录。
	// fileCheck ：立即验证 web 文件是否与待验证内容一致，成功则保存记录。
	// dbCheck :  检查是否已经验证成功过。
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
}

type AuthenticateDomainOwnerRequest struct {
	*tchttp.BaseRequest
	
	// 要验证的域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 验证类型。可取值：
	// dnsCheck ：立即验证配置 dns 的解析记录是否与待验证内容一致，成功则保存记录。
	// fileCheck ：立即验证 web 文件是否与待验证内容一致，成功则保存记录。
	// dbCheck :  检查是否已经验证成功过。
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
}

func (r *AuthenticateDomainOwnerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuthenticateDomainOwnerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "VerifyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AuthenticateDomainOwnerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AuthenticateDomainOwnerResponseParams struct {
	// 验证内容。
	// VerifyType 传 dnsCheck 时，为要配的 TXT 记录值。
	// VerifyType 传 fileCheck 时，为文件内容。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 域名验证状态。
	// >=0 为已验证归属。
	// <0 未验证归属权。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// DomainName 对应的主域名。
	// 同一主域名下的所有域名只需成功验证一次，后续均无需再验证。
	MainDomain *string `json:"MainDomain,omitnil,omitempty" name:"MainDomain"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AuthenticateDomainOwnerResponse struct {
	*tchttp.BaseResponse
	Response *AuthenticateDomainOwnerResponseParams `json:"Response"`
}

func (r *AuthenticateDomainOwnerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuthenticateDomainOwnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupStreamDetailData struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	//  UTC 格式，例如：2018-06-29T19:00:00Z。
	// 注意：和北京时间相差8小时。
	PublishTime *string `json:"PublishTime,omitnil,omitempty" name:"PublishTime"`

	// 推流唯一标识。
	UpstreamSequence *string `json:"UpstreamSequence,omitnil,omitempty" name:"UpstreamSequence"`

	// 推流来源。示例：
	// 直推流；
	// 拉流转推(1234)；
	// 注意：拉流转推来源括号中为拉流转推的任务 
	//  ID。
	SourceFrom *string `json:"SourceFrom,omitnil,omitempty" name:"SourceFrom"`

	// 主备标识。
	// 当前流为主流：1，
	// 当前流为备流: 0。
	MasterFlag *int64 `json:"MasterFlag,omitnil,omitempty" name:"MasterFlag"`
}

type BackupStreamGroupInfo struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 主备流信息。
	BackupList []*BackupStreamDetailData `json:"BackupList,omitnil,omitempty" name:"BackupList"`

	// 是否对该流开启了择优调度。
	// 0 - 未开启。
	// 1 - 已开启。
	OptimalEnable *int64 `json:"OptimalEnable,omitnil,omitempty" name:"OptimalEnable"`

	// 域名分组的分组名称。
	HostGroupName *string `json:"HostGroupName,omitnil,omitempty" name:"HostGroupName"`
}

type BandwidthInfo struct {
	// 返回格式：
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 根据粒度会有不同程度的缩减。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 带宽。
	Bandwidth *float64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`
}

type BatchDomainOperateErrors struct {
	// 操作失败的域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// API3.0错误码。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// API3.0错误信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type BillAreaInfo struct {
	// 大区名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 国家或地区明细数据。
	Countrys []*BillCountryInfo `json:"Countrys,omitnil,omitempty" name:"Countrys"`
}

type BillCountryInfo struct {
	// 国家名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 带宽明细数据信息。
	BandInfoList []*BillDataInfo `json:"BandInfoList,omitnil,omitempty" name:"BandInfoList"`
}

type BillDataInfo struct {
	// 时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 带宽，单位是 Mbps。
	Bandwidth *float64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 流量，单位是 MB。
	Flux *float64 `json:"Flux,omitnil,omitempty" name:"Flux"`

	// 峰值时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	// 原始数据为5分钟粒度，如果查询小时和天粒度数据，则返回对应粒度内的带宽峰值时间点。
	PeakTime *string `json:"PeakTime,omitnil,omitempty" name:"PeakTime"`
}

type CallBackRuleInfo struct {
	// 规则创建时间。
	// 注：此字段为北京时间（UTC+8时区）。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 规则更新时间。
	// 注：此字段为北京时间（UTC+8时区）。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`
}

type CallBackTemplateInfo struct {
	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 开播回调 URL。
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitnil,omitempty" name:"StreamBeginNotifyUrl"`

	// 混流回调 URL。(参数已弃用)。
	StreamMixNotifyUrl *string `json:"StreamMixNotifyUrl,omitnil,omitempty" name:"StreamMixNotifyUrl"`

	// 断流回调 URL。
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitnil,omitempty" name:"StreamEndNotifyUrl"`

	// 录制文件回调 URL。
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitnil,omitempty" name:"RecordNotifyUrl"`

	// 截图回调 URL。
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitnil,omitempty" name:"SnapshotNotifyUrl"`

	// 鉴黄回调 URL。
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitnil,omitempty" name:"PornCensorshipNotifyUrl"`

	// 回调的鉴权 key。
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 推流异常回调 URL。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PushExceptionNotifyUrl *string `json:"PushExceptionNotifyUrl,omitnil,omitempty" name:"PushExceptionNotifyUrl"`

	// 音频审核回调 URL。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioAuditNotifyUrl *string `json:"AudioAuditNotifyUrl,omitnil,omitempty" name:"AudioAuditNotifyUrl"`

	// 录制异常回调 URL。
	RecordExceptionNotifyUrl *string `json:"RecordExceptionNotifyUrl,omitnil,omitempty" name:"RecordExceptionNotifyUrl"`

	// 录制异常回调级别，可选择：error: 错误。alarm: 告警。
	RecordExceptionLevels []*string `json:"RecordExceptionLevels,omitnil,omitempty" name:"RecordExceptionLevels"`
}

type CallbackEventInfo struct {
	// 事件时间。
	// 接口返回支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)
	// 2）yyyy-MM-dd HH:mm:ss：使用此格式时，默认代表北京时间。
	// 接口返回的时间格式和查询请求传入的时间格式一致。
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// 事件类型。
	EventType *uint64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 回调请求。
	Request *string `json:"Request,omitnil,omitempty" name:"Request"`

	// 回调响应。
	Response *string `json:"Response,omitnil,omitempty" name:"Response"`

	// 客户接口响应时间。
	// 接口返回支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)
	// 2）yyyy-MM-dd HH:mm:ss：使用此格式时，默认代表北京时间。
	// 接口返回的时间格式和查询请求传入的时间格式一致。
	ResponseTime *string `json:"ResponseTime,omitnil,omitempty" name:"ResponseTime"`

	// 回调结果。
	ResultCode *uint64 `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`

	// 流名称。
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`
}

// Predefined struct for user
type CancelCommonMixStreamRequestParams struct {
	// 混流会话（申请混流开始到取消混流结束）标识 ID。
	// 该值与CreateCommonMixStream中的MixStreamSessionId保持一致。
	MixStreamSessionId *string `json:"MixStreamSessionId,omitnil,omitempty" name:"MixStreamSessionId"`
}

type CancelCommonMixStreamRequest struct {
	*tchttp.BaseRequest
	
	// 混流会话（申请混流开始到取消混流结束）标识 ID。
	// 该值与CreateCommonMixStream中的MixStreamSessionId保持一致。
	MixStreamSessionId *string `json:"MixStreamSessionId,omitnil,omitempty" name:"MixStreamSessionId"`
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

// Predefined struct for user
type CancelCommonMixStreamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelCommonMixStreamResponse struct {
	*tchttp.BaseResponse
	Response *CancelCommonMixStreamResponseParams `json:"Response"`
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

type CasterBriefInfo struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台名称
	CasterName *string `json:"CasterName,omitnil,omitempty" name:"CasterName"`

	// 导播台的描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 开始计费时间，值为unix时间戳
	StartBillingTime *uint64 `json:"StartBillingTime,omitnil,omitempty" name:"StartBillingTime"`

	// 结束计费时间，值为unix时间戳
	StopBillingTime *uint64 `json:"StopBillingTime,omitnil,omitempty" name:"StopBillingTime"`

	// 创建时间，值为unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 导播台状态
	// 0：停止状态，无预监，无主监
	// 1：无预监，有主监
	// 2：有预监，无主监
	// 3：有预监，有主监
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 导播台的过期时间，值为-1或unix时间戳。
	// 当值为-1时，代表永不过期。
	// 当值为特定unix时间戳时，代表过期时间为对应的时间，导播台在该时间自动停止。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 计费字段，该字段暂无作用
	FeeType *int64 `json:"FeeType,omitnil,omitempty" name:"FeeType"`
}

type CasterDisplayInfo struct {
	// 布局Index。
	// 如果使用自定义布局，为自定义布局下标。
	// 如果使用单输入布局，如使用输入1，则LayoutIndexType=1， 且LayoutIndex=1，以此类推。
	LayoutIndex *uint64 `json:"LayoutIndex,omitnil,omitempty" name:"LayoutIndex"`

	// 使用的水印Index列表。
	// 注：当作为入参使用时，列表中的水印Index需要已经存在。
	MarkPicIndexList []*uint64 `json:"MarkPicIndexList,omitnil,omitempty" name:"MarkPicIndexList"`

	// 使用的文字水印Index列表。
	// 注：作为入参使用时，列表中的Index需要已经存在。
	MarkWordIndexList []*uint64 `json:"MarkWordIndexList,omitnil,omitempty" name:"MarkWordIndexList"`

	// 使用的转场类型。
	// 注：支持的转场类型可通过DescribeCasterTransitionTypes接口获取。
	TransitionType *string `json:"TransitionType,omitnil,omitempty" name:"TransitionType"`

	// 使用的音频输入Index列表。
	// 注：当该字段非空时，表示使用布局中对应的输入源的视频，AudioIndexList中对应的输入源的音频。且这些输入源需已存在。
	AudioIndexList []*uint64 `json:"AudioIndexList,omitnil,omitempty" name:"AudioIndexList"`

	// 作为入参时，表示使用点播输入源，单画面输入时，点播文件是否从头开始播放。
	// 默认为0。
	// 有效值，0,1。
	// 0代表不从头播放
	// 1代表从头播放
	InputStartType *int64 `json:"InputStartType,omitnil,omitempty" name:"InputStartType"`

	// LayoutIndex类型，
	// 默认值:0
	// 可选值[0,1]
	// 0:默认类型，代表普通布局
	// 1:单输入类型，代表单输入布局
	LayoutIndexType *int64 `json:"LayoutIndexType,omitnil,omitempty" name:"LayoutIndexType"`
}

type CasterInfo struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台名称
	CasterName *string `json:"CasterName,omitnil,omitempty" name:"CasterName"`

	// 导播台上一次启动pgm的时间，值为unix时间戳。
	StartLiveTime *uint64 `json:"StartLiveTime,omitnil,omitempty" name:"StartLiveTime"`

	// 导播台的描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 导播台创建时间，值为unix时间戳。
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 导播台状态 
	// 0：停止状态，无预监，无主监
	// 1：无预监，有主监
	// 2：有预监，无主监
	// 3：有预监，有主监
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 导播台的过期时间戳。值为-1或unix时间戳。 
	// 默认值为-1。 当值为-1时，表示该导播台永不过期。 
	// 当值为正常unix时间戳时，导播台将在该时间过期。 
	// 导播台过期后，预监与主监画面将自动停止，转推自动停止。 
	// 点播、直播url将停止转拉，推流url需自行停止推流。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 导播台延时播放时间，单位为秒。 
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 导播台主监输出的宽度，单位为像素。
	PgmWidth *uint64 `json:"PgmWidth,omitnil,omitempty" name:"PgmWidth"`

	// 导播台主监输出的高度，单位为像素。
	PgmHeight *uint64 `json:"PgmHeight,omitnil,omitempty" name:"PgmHeight"`

	// 导播台主监输出的帧率。
	PgmFps *uint64 `json:"PgmFps,omitnil,omitempty" name:"PgmFps"`

	// 导播台主监输出的码率，单位为kbps
	PgmBitRate *uint64 `json:"PgmBitRate,omitnil,omitempty" name:"PgmBitRate"`

	// 导播台主监输出的音频码率，单位为kbps。
	PgmAudioBitRate *uint64 `json:"PgmAudioBitRate,omitnil,omitempty" name:"PgmAudioBitRate"`

	// 导播台的计费类型。 
	// 0 通用型 1 播单型。
	// 注： 本参数暂无作用。
	FeeType *int64 `json:"FeeType,omitnil,omitempty" name:"FeeType"`

	// 录制模板id。
	RecordTemplateId *uint64 `json:"RecordTemplateId,omitnil,omitempty" name:"RecordTemplateId"`

	// 录制状态。 
	// 0：未录制 
	// 1：录制中
	RecordStatus *uint64 `json:"RecordStatus,omitnil,omitempty" name:"RecordStatus"`

	// 录制接口返回的taskid
	RecordTaskId *string `json:"RecordTaskId,omitnil,omitempty" name:"RecordTaskId"`
}

type CasterInputInfo struct {
	// 输入源Index。
	// 范围[1, 24]
	InputIndex *int64 `json:"InputIndex,omitnil,omitempty" name:"InputIndex"`

	// 输入源类型。
	// 范围[0,1,2,3,4]。
	// 0：推流地址。
	// 1：点播文件地址。
	// 2：直播拉流地址。
	// 3：图片地址。
	// 4：webrtc协议推流地址。
	InputType *int64 `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 输入源的源地址。
	// 最大允许长度512。
	// 当InputType为0（推流地址），2（直播拉流地址），3（图片地址）,4（webrtc推流地址）这几种类型时，URL需填入该字段。
	InputUrl *string `json:"InputUrl,omitnil,omitempty" name:"InputUrl"`

	// 输入源描述。
	// 最大允许长度256字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 点播地址列表。仅当input type为1（点播地址）时，将一个或多个点播地址，填入该字段。
	// 单个地址最大允许长度512字符。
	// 最多允许同时填入5个地址。
	// 注：此时需保持InputUrl字段为空。
	InputUrls []*string `json:"InputUrls,omitnil,omitempty" name:"InputUrls"`

	// 是否启用点播无限循环播放。
	// 注：当前该字段未生效，默认为True。
	LoopEnable *bool `json:"LoopEnable,omitnil,omitempty" name:"LoopEnable"`

	// 点播循环次数。
	// 允许值-1或正整数。
	// 当值为-1时，表示无限循环。
	// 当值为其他正整数时，表示循环对应次数。
	// 注：该字段暂未生效。
	LoopNumber *int64 `json:"LoopNumber,omitnil,omitempty" name:"LoopNumber"`

	// 是否启用拉取到导播台。
	// 注：该字段默认强制为true。
	PullPushEnable *bool `json:"PullPushEnable,omitnil,omitempty" name:"PullPushEnable"`

	// 输入源音量百分比。
	// 默认为100。表示音量为原始大小。
	// 允许值[0,200]。
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`
}

type CasterLayoutInfo struct {
	// 布局Index。
	LayoutIndex *int64 `json:"LayoutIndex,omitnil,omitempty" name:"LayoutIndex"`

	// 布局模板Id。
	// 有效值[1，20，21，31，32，41]
	// 当使用布局模版时，无需LayoutParams参数，导播台将使用模版布局参数。
	LayoutTemplateId *int64 `json:"LayoutTemplateId,omitnil,omitempty" name:"LayoutTemplateId"`

	// 布局绑定的输入列表。按布局LayerId从小到大，按顺序排列。
	// 已有两个画面的布局为例，输入1对应LayerId为1，输入2对应的LayerId为2，该字段应该填入"1|2"。
	InputIndexList *string `json:"InputIndexList,omitnil,omitempty" name:"InputIndexList"`

	// 详细的布局参数列表。
	LayoutParams []*CasterLayoutParam `json:"LayoutParams,omitnil,omitempty" name:"LayoutParams"`

	// 布局输出的宽度，单位为像素。
	// 默认为1280像素。
	// 注：该值仅在画中画布局，且未设置PgmWidth时生效。
	LayoutWidth *uint64 `json:"LayoutWidth,omitnil,omitempty" name:"LayoutWidth"`

	// 布局输出的高度，单位为像素。
	// 注：该参数仅在画中画布局，且未设置PgmHeight时生效。
	LayoutHeight *uint64 `json:"LayoutHeight,omitnil,omitempty" name:"LayoutHeight"`
}

type CasterLayoutParam struct {
	// 布局层ID。
	// 在画面最终渲染时，将按ID从小到大的顺序，由下至上渲染。
	LayerId *int64 `json:"LayerId,omitnil,omitempty" name:"LayerId"`

	// 布局层宽度。单位为百分比，范围[0.0,1.0]，表示该层在最终画面上所占的比例值。
	LayerWidth *float64 `json:"LayerWidth,omitnil,omitempty" name:"LayerWidth"`

	// 布局层高度。单位为百分比，范围[0.0,1.0]， 表示该层在最终画面上所占的比例值。
	LayerHeight *float64 `json:"LayerHeight,omitnil,omitempty" name:"LayerHeight"`

	// 布局层位置x坐标。单位为百分比，范围[0.0,1.0]，表示该层在最终画面上x坐标所占的比例值。
	LayerLocationX *float64 `json:"LayerLocationX,omitnil,omitempty" name:"LayerLocationX"`

	// 布局层位置Y坐标。单位为百分比，范围[0.0,1.0]， 表示该层在最终画面Y坐标上所占的比例值。
	LayerLocationY *float64 `json:"LayerLocationY,omitnil,omitempty" name:"LayerLocationY"`

	// 是否启用抠图。
	UsePortraitSegment *bool `json:"UsePortraitSegment,omitnil,omitempty" name:"UsePortraitSegment"`
}

type CasterMarkPicInfo struct {
	// 水印图片Index。
	MarkPicIndex *int64 `json:"MarkPicIndex,omitnil,omitempty" name:"MarkPicIndex"`

	// 注：该字段已废弃。
	MarkPicId *int64 `json:"MarkPicId,omitnil,omitempty" name:"MarkPicId"`

	// 水印图片在输出时的宽度。
	// 当该值为大于1的整数值时，单位为像素，允许范围[1,1920]。 
	// 当该值为小于1大于0的小数时，单位为百分比，表示水印在最终画面上所占的比例值。
	MarkPicWidth *float64 `json:"MarkPicWidth,omitnil,omitempty" name:"MarkPicWidth"`

	// 水印图片在输出时的高度。
	// 当该值为大于1的整数值时，单位为像素，允许范围[1,1080]。 
	// 当该值为小于1大于0的小数时，单位为百分比，表示水印在输出上所占的比例值。
	MarkPicHeight *float64 `json:"MarkPicHeight,omitnil,omitempty" name:"MarkPicHeight"`

	// 水印图片在输出时的X轴坐标。
	// 当该值为大于1的整数值时，单位为像素，允许范围[1,1920]。 
	// 当该值为小于1大于0的小数时，单位为百分比，表示水印在最终画面上x坐标所占的比例值。
	MarkPicLocationX *float64 `json:"MarkPicLocationX,omitnil,omitempty" name:"MarkPicLocationX"`

	// 水印图片在输出时的Y坐标。
	// 当该值为大于1的整数值时，单位为像素，允许范围[1,1080]。 
	// 当该值为小于1大于0的小数时，单位为百分比，表示水印在最终画面Y坐标上所占的比例值。
	MarkPicLocationY *float64 `json:"MarkPicLocationY,omitnil,omitempty" name:"MarkPicLocationY"`

	// 水印地址。
	// 最大长度256字符，且url需以jpg、jpeg、png、bmp、gif后缀结尾。
	MarkPicUrl *string `json:"MarkPicUrl,omitnil,omitempty" name:"MarkPicUrl"`

	// 水印描述。
	// 最大允许长度为256。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否启用了等比例缩放。
	// 注：该字段仅做状态保存，无实际效果。
	IsEqualProportion *bool `json:"IsEqualProportion,omitnil,omitempty" name:"IsEqualProportion"`
}

type CasterMarkWordInfo struct {
	// 文字水印Index。
	MarkWordIndex *int64 `json:"MarkWordIndex,omitnil,omitempty" name:"MarkWordIndex"`

	// 文字水印内容。
	MarkWordText *string `json:"MarkWordText,omitnil,omitempty" name:"MarkWordText"`

	// 文字水印的字号。
	// 范围[16, 60]
	MarkWordFontSize *uint64 `json:"MarkWordFontSize,omitnil,omitempty" name:"MarkWordFontSize"`

	// 文字水印的颜色，值为HEX颜色代码
	MarkWordFontColor *string `json:"MarkWordFontColor,omitnil,omitempty" name:"MarkWordFontColor"`

	// 文字水印的字体类型。
	// 范围[1,2]。
	// 1. 宋体
	// 2. 黑体
	MarkWordFontType *uint64 `json:"MarkWordFontType,omitnil,omitempty" name:"MarkWordFontType"`

	// 文字水印的x坐标位置，单位百分比。
	// 范围[0.0, 1.0]
	MarkWordLocationX *float64 `json:"MarkWordLocationX,omitnil,omitempty" name:"MarkWordLocationX"`

	// 文字水印的Y坐标位置，单位百分比。
	// 范围[0.0, 1.0]
	MarkWordLocationY *float64 `json:"MarkWordLocationY,omitnil,omitempty" name:"MarkWordLocationY"`

	// 是否开启文字跑马灯功能。
	// 默认为false。
	MarkWordRollEnable *bool `json:"MarkWordRollEnable,omitnil,omitempty" name:"MarkWordRollEnable"`

	// 跑马灯文字显示一遍的时间，单位为秒。
	// 默认为5s。
	// 范围[5, 600]。
	MarkWordRollOnceTime *uint64 `json:"MarkWordRollOnceTime,omitnil,omitempty" name:"MarkWordRollOnceTime"`

	// 跑马灯文字的方向。
	// 默认值为0。
	// 范围[0,1]。
	// 0 从左到右
	// 1 从右到左
	MarkWordRollDirection *int64 `json:"MarkWordRollDirection,omitnil,omitempty" name:"MarkWordRollDirection"`

	// 跑马灯文字显示的起始x坐标，单位百分比。
	// 范围[0.0, 1.0]
	MarkWordRollStartLocationX *float64 `json:"MarkWordRollStartLocationX,omitnil,omitempty" name:"MarkWordRollStartLocationX"`

	// 跑马灯文字显示的截止x坐标，单位百分比。
	// 范围[0.0, 1.0]
	MarkWordRollEndLocationX *float64 `json:"MarkWordRollEndLocationX,omitnil,omitempty" name:"MarkWordRollEndLocationX"`
}

type CasterOutputInfo struct {
	// 推流信息Index。
	// 当OutputType为1（表示推流到腾讯云直播）时，该值固定为0。
	// 范围[0,10]。
	OutputIndex *uint64 `json:"OutputIndex,omitnil,omitempty" name:"OutputIndex"`

	// rtmp协议推流地址。
	// 最大允许长度512字符。
	OutputUrl *string `json:"OutputUrl,omitnil,omitempty" name:"OutputUrl"`

	// 描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 推流到腾讯云直播源站时，使用的流ID。
	// 仅当OutputType为1时生效。
	// 最大允许128字符。
	OutputStreamId *string `json:"OutputStreamId,omitnil,omitempty" name:"OutputStreamId"`

	// 推流类型。
	// 范围[1,2]
	// 1. 推送到腾讯云直播源站。
	// 2. 推送到第三方源站。
	OutputType *uint64 `json:"OutputType,omitnil,omitempty" name:"OutputType"`

	// 推到腾讯云直播源站时，使用的域名。
	// 最大允许128字符，且域名需属于当前账号绑定的云直播推流域名。
	// 仅在OutputType为1时生效。
	OutputDomainName *string `json:"OutputDomainName,omitnil,omitempty" name:"OutputDomainName"`

	// 推到腾讯云直播源站时，使用的AppName。
	// 最大允许64字符。
	// 仅在OutputType为1时生效。
	OutputAppName *string `json:"OutputAppName,omitnil,omitempty" name:"OutputAppName"`

	// 推到腾讯云直播源站时需要添加的推流参数。
	// 最大允许长度256字符。
	// 仅在OutputType为1时生效。
	OutputParam *string `json:"OutputParam,omitnil,omitempty" name:"OutputParam"`
}

type CdnPlayStatData struct {
	// 时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 带宽，单位: Mbps。
	Bandwidth *float64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 流量，单位: MB。
	Flux *float64 `json:"Flux,omitnil,omitempty" name:"Flux"`

	// 新增请求数。
	Request *uint64 `json:"Request,omitnil,omitempty" name:"Request"`

	// 并发连接数。
	Online *uint64 `json:"Online,omitnil,omitempty" name:"Online"`
}

type CertInfo struct {
	// 证书 ID。
	CertId *int64 `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书名称。
	CertName *string `json:"CertName,omitnil,omitempty" name:"CertName"`

	// 描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间，UTC 格式。
	// 注：此字段为北京时间（UTC+8时区）。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 证书内容。
	HttpsCrt *string `json:"HttpsCrt,omitnil,omitempty" name:"HttpsCrt"`

	// 证书类型。
	// 0：用户添加证书，
	// 1：腾讯云托管证书。
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 证书过期时间，UTC 格式。
	// 注：此字段为北京时间（UTC+8时区）。
	CertExpireTime *string `json:"CertExpireTime,omitnil,omitempty" name:"CertExpireTime"`

	// 使用此证书的域名列表。
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
}

type ClientIpPlaySumInfo struct {
	// 客户端 IP，点分型。
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 客户端所在省份。
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 总流量。
	TotalFlux *float64 `json:"TotalFlux,omitnil,omitempty" name:"TotalFlux"`

	// 总请求数。
	TotalRequest *uint64 `json:"TotalRequest,omitnil,omitempty" name:"TotalRequest"`

	// 总失败请求数。
	TotalFailedRequest *uint64 `json:"TotalFailedRequest,omitnil,omitempty" name:"TotalFailedRequest"`

	// 客户端所在国家。
	CountryArea *string `json:"CountryArea,omitnil,omitempty" name:"CountryArea"`
}

type CommonMixControlParams struct {
	// 取值范围[0,1]。
	// 填1时，当参数中图层分辨率参数与视频实际分辨率不一致时，自动从视频中按图层设置的分辨率比例进行裁剪。
	UseMixCropCenter *int64 `json:"UseMixCropCenter,omitnil,omitempty" name:"UseMixCropCenter"`

	// 取值范围[0,1]
	// 填1时，当InputStreamList中个数为1时，且OutputParams.OutputStreamType为1时，不执行取消操作，执行拷贝流操作
	AllowCopy *int64 `json:"AllowCopy,omitnil,omitempty" name:"AllowCopy"`

	// 取值范围[0,1]
	// 填1时，透传原始流的sei
	PassInputSei *int64 `json:"PassInputSei,omitnil,omitempty" name:"PassInputSei"`

	// 取值范围[0,1]
	// 填1时，图片输入中的透明通道生效。
	UsePictureTransparent *int64 `json:"UsePictureTransparent,omitnil,omitempty" name:"UsePictureTransparent"`
}

type CommonMixCropParams struct {
	// 裁剪的宽度。取值范围[0，2000]。
	CropWidth *float64 `json:"CropWidth,omitnil,omitempty" name:"CropWidth"`

	// 裁剪的高度。取值范围[0，2000]。
	CropHeight *float64 `json:"CropHeight,omitnil,omitempty" name:"CropHeight"`

	// 裁剪的起始X坐标。取值范围[0，2000]。
	CropStartLocationX *float64 `json:"CropStartLocationX,omitnil,omitempty" name:"CropStartLocationX"`

	// 裁剪的起始Y坐标。取值范围[0，2000]。
	CropStartLocationY *float64 `json:"CropStartLocationY,omitnil,omitempty" name:"CropStartLocationY"`
}

type CommonMixInputParam struct {
	// 输入流名称。80字节以内，仅含字母、数字以及下划线的字符串。
	// 当LayoutParams.InputType=0(音视频)/4(纯音频)/5(纯视频)时，该值为需要混流的流名称。
	// 当LayoutParams.InputType=2(图片)/3(画布)时，该值仅用作标识输入，可用类似Canvas1、Pictrue1的名称。
	InputStreamName *string `json:"InputStreamName,omitnil,omitempty" name:"InputStreamName"`

	// 输入流布局参数。
	LayoutParams *CommonMixLayoutParams `json:"LayoutParams,omitnil,omitempty" name:"LayoutParams"`

	// 输入流裁剪参数。
	CropParams *CommonMixCropParams `json:"CropParams,omitnil,omitempty" name:"CropParams"`

	// 抠图参数。
	PortraitSegmentParams *MixPortraitSegmentParams `json:"PortraitSegmentParams,omitnil,omitempty" name:"PortraitSegmentParams"`
}

type CommonMixLayoutParams struct {
	// 输入图层。取值范围[1，16]。
	// 1)背景流（即大主播画面或画布）的 image_layer 填1。
	// 2)纯音频混流，该参数也需填。
	// 注意：不同输入，该值不可重复
	ImageLayer *int64 `json:"ImageLayer,omitnil,omitempty" name:"ImageLayer"`

	// 输入类型。取值范围[0，5]。
	// 不填默认为0。
	// 0表示输入流为音视频。
	// 2表示输入流为图片。
	// 3表示输入流为画布。 
	// 4表示输入流为音频。
	// 5表示输入流为纯视频。
	InputType *int64 `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 输入画面在输出时的高度。取值范围：
	// 像素：[0，2000]
	// 百分比：[0.01，0.99]
	// 不填默认为输入流的高度。
	// 使用百分比时，期望输出为（百分比 * 背景高）。
	ImageHeight *float64 `json:"ImageHeight,omitnil,omitempty" name:"ImageHeight"`

	// 输入画面在输出时的宽度。取值范围：
	// 像素：[0，2000]
	// 百分比：[0.01，0.99]
	// 不填默认为输入流的宽度。
	// 使用百分比时，期望输出为（百分比 * 背景宽）。
	ImageWidth *float64 `json:"ImageWidth,omitnil,omitempty" name:"ImageWidth"`

	// 输入在输出画面的X偏移。取值范围：
	// 像素：[0，2000]
	// 百分比：[0.01，0.99]
	// 不填默认为0。
	// 相对于大主播背景画面左上角的横向偏移。 
	// 使用百分比时，期望输出为（百分比 * 背景宽）。
	LocationX *float64 `json:"LocationX,omitnil,omitempty" name:"LocationX"`

	// 输入在输出画面的Y偏移。取值范围：
	// 像素：[0，2000]
	// 百分比：[0.01，0.99]
	// 不填默认为0。
	// 相对于大主播背景画面左上角的纵向偏移。 
	// 使用百分比时，期望输出为（百分比 * 背景宽）
	LocationY *float64 `json:"LocationY,omitnil,omitempty" name:"LocationY"`

	// 当InputType为3(画布)时，该值表示画布的颜色。
	// 常用的颜色有：
	// 红色：0xCC0033。
	// 黄色：0xCC9900。
	// 绿色：0xCCCC33。
	// 蓝色：0x99CCFF。
	// 黑色：0x000000。
	// 白色：0xFFFFFF。
	// 灰色：0x999999。
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`

	// 当InputType为2(图片)时，该值是水印ID。
	WatermarkId *int64 `json:"WatermarkId,omitnil,omitempty" name:"WatermarkId"`
}

type CommonMixOutputParams struct {
	// 输出流名称。
	OutputStreamName *string `json:"OutputStreamName,omitnil,omitempty" name:"OutputStreamName"`

	// 输出流类型，取值范围[0,1]。
	// 不填默认为0。
	// 当输出流为输入流 list 中的一条时，填写0。
	// 当期望生成的混流结果成为一条新流时，该值填为1。
	// 该值为1时，output_stream_id 不能出现在 input_stram_list 中，且直播后台中，不能存在相同 ID 的流。
	OutputStreamType *int64 `json:"OutputStreamType,omitnil,omitempty" name:"OutputStreamType"`

	// 输出流比特率。取值范围[1，10000]。
	// 不填的情况下，系统会自动判断。
	OutputStreamBitRate *int64 `json:"OutputStreamBitRate,omitnil,omitempty" name:"OutputStreamBitRate"`

	// 输出流GOP大小。取值范围[1,10]。
	// 不填的情况下，系统会自动判断。
	OutputStreamGop *int64 `json:"OutputStreamGop,omitnil,omitempty" name:"OutputStreamGop"`

	// 输出流帧率大小。取值范围[1,60]。
	// 不填的情况下，系统会自动判断。
	OutputStreamFrameRate *int64 `json:"OutputStreamFrameRate,omitnil,omitempty" name:"OutputStreamFrameRate"`

	// 输出流音频比特率。取值范围[1,500]
	// 不填的情况下，系统会自动判断。
	OutputAudioBitRate *int64 `json:"OutputAudioBitRate,omitnil,omitempty" name:"OutputAudioBitRate"`

	// 输出流音频采样率。取值范围[96000, 88200, 64000, 48000, 44100, 32000,24000, 22050, 16000, 12000, 11025, 8000]。
	// 不填的情况下，系统会自动判断。
	OutputAudioSampleRate *int64 `json:"OutputAudioSampleRate,omitnil,omitempty" name:"OutputAudioSampleRate"`

	// 输出流音频声道数。取值范围[1,2]。
	// 不填的情况下，系统会自动判断。
	OutputAudioChannels *int64 `json:"OutputAudioChannels,omitnil,omitempty" name:"OutputAudioChannels"`

	// 输出流中的sei信息。如果无特殊需要，不填。
	MixSei *string `json:"MixSei,omitnil,omitempty" name:"MixSei"`
}

type ConcurrentRecordStreamNum struct {
	// 时间点。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 路数。
	Num *uint64 `json:"Num,omitnil,omitempty" name:"Num"`
}

// Predefined struct for user
type CopyCasterRequestParams struct {
	// 源导播台的ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 复制产生的新导播台名称
	CasterName *string `json:"CasterName,omitnil,omitempty" name:"CasterName"`

	// 复制产生的导播台推送到云直播的流id
	// 注意：该流id不能与云直播中的流id重复
	OutputStreamId *string `json:"OutputStreamId,omitnil,omitempty" name:"OutputStreamId"`
}

type CopyCasterRequest struct {
	*tchttp.BaseRequest
	
	// 源导播台的ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 复制产生的新导播台名称
	CasterName *string `json:"CasterName,omitnil,omitempty" name:"CasterName"`

	// 复制产生的导播台推送到云直播的流id
	// 注意：该流id不能与云直播中的流id重复
	OutputStreamId *string `json:"OutputStreamId,omitnil,omitempty" name:"OutputStreamId"`
}

func (r *CopyCasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyCasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "CasterName")
	delete(f, "OutputStreamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyCasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyCasterResponseParams struct {
	// 复制生成的导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyCasterResponse struct {
	*tchttp.BaseResponse
	Response *CopyCasterResponseParams `json:"Response"`
}

func (r *CopyCasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyCasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCasterInputPushUrlRequestParams struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 请求生成推流地址的输入Index。
	// 允许范围[1,24]。
	InputIndex *uint64 `json:"InputIndex,omitnil,omitempty" name:"InputIndex"`

	// 生成推流地址协议。
	// 范围[rtmp,webrtc]。
	// 注：获取webrtc推流地址时，需配合腾讯云快直播推流sdk才可成功推流。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type CreateCasterInputPushUrlRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 请求生成推流地址的输入Index。
	// 允许范围[1,24]。
	InputIndex *uint64 `json:"InputIndex,omitnil,omitempty" name:"InputIndex"`

	// 生成推流地址协议。
	// 范围[rtmp,webrtc]。
	// 注：获取webrtc推流地址时，需配合腾讯云快直播推流sdk才可成功推流。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

func (r *CreateCasterInputPushUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCasterInputPushUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "InputIndex")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCasterInputPushUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCasterInputPushUrlResponseParams struct {
	// 生成可使用的推流地址。
	PushUrl *string `json:"PushUrl,omitnil,omitempty" name:"PushUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCasterInputPushUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateCasterInputPushUrlResponseParams `json:"Response"`
}

func (r *CreateCasterInputPushUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCasterInputPushUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCasterPgmFromPvwRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type CreateCasterPgmFromPvwRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *CreateCasterPgmFromPvwRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCasterPgmFromPvwRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCasterPgmFromPvwRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCasterPgmFromPvwResponseParams struct {
	// 主监任务的rtmp协议预览地址。 
	// 注：该地址仅供预览，不可分发。
	PgmPlayUrl *string `json:"PgmPlayUrl,omitnil,omitempty" name:"PgmPlayUrl"`

	// 注：该字段已废弃，请结合腾讯云直播播放地址生成策略生成cdn播放地址
	CdnPlayUrl *string `json:"CdnPlayUrl,omitnil,omitempty" name:"CdnPlayUrl"`

	// 主监任务在腾讯云直播侧的流ID。
	CdnStreamId *string `json:"CdnStreamId,omitnil,omitempty" name:"CdnStreamId"`

	// 主监任务的webrtc协议播放地址。 
	// 注：
	// 1.该预览地址仅作为预览，不可分发。
	// 2.webrtc播放地址需配合腾讯云快直播播放sdk使用。
	PgmWebRTCPlayUrl *string `json:"PgmWebRTCPlayUrl,omitnil,omitempty" name:"PgmWebRTCPlayUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCasterPgmFromPvwResponse struct {
	*tchttp.BaseResponse
	Response *CreateCasterPgmFromPvwResponseParams `json:"Response"`
}

func (r *CreateCasterPgmFromPvwResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCasterPgmFromPvwResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCasterPgmRequestParams struct {
	// 导播台展示信息参数。
	PgmDisplayInfo *CasterDisplayInfo `json:"PgmDisplayInfo,omitnil,omitempty" name:"PgmDisplayInfo"`

	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type CreateCasterPgmRequest struct {
	*tchttp.BaseRequest
	
	// 导播台展示信息参数。
	PgmDisplayInfo *CasterDisplayInfo `json:"PgmDisplayInfo,omitnil,omitempty" name:"PgmDisplayInfo"`

	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *CreateCasterPgmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCasterPgmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PgmDisplayInfo")
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCasterPgmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCasterPgmResponseParams struct {
	// 主监任务的rtmp协议预览地址。
	// 注：该地址仅供预览，不可分发。
	PgmPlayUrl *string `json:"PgmPlayUrl,omitnil,omitempty" name:"PgmPlayUrl"`

	// 注：该字段已废弃，请结合腾讯云直播播放地址生成策略生成cdn播放地址。
	CdnPlayUrl *string `json:"CdnPlayUrl,omitnil,omitempty" name:"CdnPlayUrl"`

	// 主监任务在腾讯云直播侧的流ID。
	CdnStreamId *string `json:"CdnStreamId,omitnil,omitempty" name:"CdnStreamId"`

	// 主监任务的webrtc协议播放地址。
	// 注：
	// 1. 该预览地址仅作为预览，不可分发。
	// 2. webrtc播放地址需配合腾讯云快直播播放sdk使用。
	PgmWebRTCPlayUrl *string `json:"PgmWebRTCPlayUrl,omitnil,omitempty" name:"PgmWebRTCPlayUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCasterPgmResponse struct {
	*tchttp.BaseResponse
	Response *CreateCasterPgmResponseParams `json:"Response"`
}

func (r *CreateCasterPgmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCasterPgmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCasterPvwRequestParams struct {
	// 导播台预监展示信息参数
	PvwDisplayInfo *CasterDisplayInfo `json:"PvwDisplayInfo,omitnil,omitempty" name:"PvwDisplayInfo"`

	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type CreateCasterPvwRequest struct {
	*tchttp.BaseRequest
	
	// 导播台预监展示信息参数
	PvwDisplayInfo *CasterDisplayInfo `json:"PvwDisplayInfo,omitnil,omitempty" name:"PvwDisplayInfo"`

	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *CreateCasterPvwRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCasterPvwRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PvwDisplayInfo")
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCasterPvwRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCasterPvwResponseParams struct {
	// 预监任务的画面rtmp协议预览地址。
	// 注：该预览地址仅供画面预览，不可分发。
	PvwPlayUrl *string `json:"PvwPlayUrl,omitnil,omitempty" name:"PvwPlayUrl"`

	// 预监任务的webrtc协议预览画面。
	// 注：
	// 1. 该预览地址仅供预览，不可分发。
	// 2. webrtc播放地址，需配合腾讯云快直播播放sdk使用
	PvwWebRTCPlayUrl *string `json:"PvwWebRTCPlayUrl,omitnil,omitempty" name:"PvwWebRTCPlayUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCasterPvwResponse struct {
	*tchttp.BaseResponse
	Response *CreateCasterPvwResponseParams `json:"Response"`
}

func (r *CreateCasterPvwResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCasterPvwResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCasterRequestParams struct {
	// 导播台名称
	CasterName *string `json:"CasterName,omitnil,omitempty" name:"CasterName"`

	// 导播台的描述
	// 最大允许长度256
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 导播台的过期时间戳。值为-1或unix时间戳。
	// 默认值为-1。
	// 当值为-1时，表示该导播台永不过期。
	// 当值为正常unix时间戳时，导播台将在该时间过期。
	// 导播台过期后，预监与主监画面将自动停止，转推自动停止。
	// 点播、直播url将停止转拉，推流url需自行停止推流。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 导播台延时播放时间，单位为秒。
	// 默认为0，最大支持300秒
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 导播台转场类型。
	// 默认为空。
	// 允许使用通过DescribeCasterTransitionTypes接口中查询到的转场类型。
	TransitionType *string `json:"TransitionType,omitnil,omitempty" name:"TransitionType"`

	// 导播台主监输出的宽度，单位为像素。
	// 默认为1280，最大允许4096。
	PgmWidth *uint64 `json:"PgmWidth,omitnil,omitempty" name:"PgmWidth"`

	// 导播台主监输出的高度，单位为像素。
	// 默认为720，最大允许2160。
	PgmHeight *uint64 `json:"PgmHeight,omitnil,omitempty" name:"PgmHeight"`

	// 导播台主监输出的帧率。
	// 默认为0，表示随源输出。
	// 最大支持60。
	PgmFps *uint64 `json:"PgmFps,omitnil,omitempty" name:"PgmFps"`

	// 导播台主监输出的码率，单位为kbps。
	// 默认为0，表示随源的码率输出。
	// 最大允许10000kbps。
	PgmBitRate *uint64 `json:"PgmBitRate,omitnil,omitempty" name:"PgmBitRate"`

	// 导播台的计费类型。
	// 0 通用型 
	// 1 播单型。
	// 注： 本参数暂无作用。
	FeeType *int64 `json:"FeeType,omitnil,omitempty" name:"FeeType"`

	// 导播台主监输出的音频码率，单位为kbps。
	// 可选项：[0, 128, 192, 256]
	// 默认值为0，表示随源的音频码率输出。
	PgmAudioBitRate *uint64 `json:"PgmAudioBitRate,omitnil,omitempty" name:"PgmAudioBitRate"`
}

type CreateCasterRequest struct {
	*tchttp.BaseRequest
	
	// 导播台名称
	CasterName *string `json:"CasterName,omitnil,omitempty" name:"CasterName"`

	// 导播台的描述
	// 最大允许长度256
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 导播台的过期时间戳。值为-1或unix时间戳。
	// 默认值为-1。
	// 当值为-1时，表示该导播台永不过期。
	// 当值为正常unix时间戳时，导播台将在该时间过期。
	// 导播台过期后，预监与主监画面将自动停止，转推自动停止。
	// 点播、直播url将停止转拉，推流url需自行停止推流。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 导播台延时播放时间，单位为秒。
	// 默认为0，最大支持300秒
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 导播台转场类型。
	// 默认为空。
	// 允许使用通过DescribeCasterTransitionTypes接口中查询到的转场类型。
	TransitionType *string `json:"TransitionType,omitnil,omitempty" name:"TransitionType"`

	// 导播台主监输出的宽度，单位为像素。
	// 默认为1280，最大允许4096。
	PgmWidth *uint64 `json:"PgmWidth,omitnil,omitempty" name:"PgmWidth"`

	// 导播台主监输出的高度，单位为像素。
	// 默认为720，最大允许2160。
	PgmHeight *uint64 `json:"PgmHeight,omitnil,omitempty" name:"PgmHeight"`

	// 导播台主监输出的帧率。
	// 默认为0，表示随源输出。
	// 最大支持60。
	PgmFps *uint64 `json:"PgmFps,omitnil,omitempty" name:"PgmFps"`

	// 导播台主监输出的码率，单位为kbps。
	// 默认为0，表示随源的码率输出。
	// 最大允许10000kbps。
	PgmBitRate *uint64 `json:"PgmBitRate,omitnil,omitempty" name:"PgmBitRate"`

	// 导播台的计费类型。
	// 0 通用型 
	// 1 播单型。
	// 注： 本参数暂无作用。
	FeeType *int64 `json:"FeeType,omitnil,omitempty" name:"FeeType"`

	// 导播台主监输出的音频码率，单位为kbps。
	// 可选项：[0, 128, 192, 256]
	// 默认值为0，表示随源的音频码率输出。
	PgmAudioBitRate *uint64 `json:"PgmAudioBitRate,omitnil,omitempty" name:"PgmAudioBitRate"`
}

func (r *CreateCasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterName")
	delete(f, "Description")
	delete(f, "ExpireTime")
	delete(f, "DelayTime")
	delete(f, "TransitionType")
	delete(f, "PgmWidth")
	delete(f, "PgmHeight")
	delete(f, "PgmFps")
	delete(f, "PgmBitRate")
	delete(f, "FeeType")
	delete(f, "PgmAudioBitRate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCasterResponseParams struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCasterResponse struct {
	*tchttp.BaseResponse
	Response *CreateCasterResponseParams `json:"Response"`
}

func (r *CreateCasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCommonMixStreamRequestParams struct {
	// 混流会话（申请混流开始到取消混流结束）标识 ID。80字节以内，仅含字母、数字以及下划线的字符串。
	MixStreamSessionId *string `json:"MixStreamSessionId,omitnil,omitempty" name:"MixStreamSessionId"`

	// 混流输入流列表。
	InputStreamList []*CommonMixInputParam `json:"InputStreamList,omitnil,omitempty" name:"InputStreamList"`

	// 混流输出流参数。
	OutputParams *CommonMixOutputParams `json:"OutputParams,omitnil,omitempty" name:"OutputParams"`

	// 输入模板 ID，若设置该参数，将按默认模板布局输出，无需填入自定义位置参数。
	// 不填默认为0。
	// 两输入源支持10，20，30，40，50。
	// 三输入源支持310，390，391。
	// 四输入源支持410。
	// 五输入源支持510，590。
	// 六输入源支持610。
	MixStreamTemplateId *int64 `json:"MixStreamTemplateId,omitnil,omitempty" name:"MixStreamTemplateId"`

	// 混流的特殊控制参数。如无特殊需求，无需填写。
	ControlParams *CommonMixControlParams `json:"ControlParams,omitnil,omitempty" name:"ControlParams"`
}

type CreateCommonMixStreamRequest struct {
	*tchttp.BaseRequest
	
	// 混流会话（申请混流开始到取消混流结束）标识 ID。80字节以内，仅含字母、数字以及下划线的字符串。
	MixStreamSessionId *string `json:"MixStreamSessionId,omitnil,omitempty" name:"MixStreamSessionId"`

	// 混流输入流列表。
	InputStreamList []*CommonMixInputParam `json:"InputStreamList,omitnil,omitempty" name:"InputStreamList"`

	// 混流输出流参数。
	OutputParams *CommonMixOutputParams `json:"OutputParams,omitnil,omitempty" name:"OutputParams"`

	// 输入模板 ID，若设置该参数，将按默认模板布局输出，无需填入自定义位置参数。
	// 不填默认为0。
	// 两输入源支持10，20，30，40，50。
	// 三输入源支持310，390，391。
	// 四输入源支持410。
	// 五输入源支持510，590。
	// 六输入源支持610。
	MixStreamTemplateId *int64 `json:"MixStreamTemplateId,omitnil,omitempty" name:"MixStreamTemplateId"`

	// 混流的特殊控制参数。如无特殊需求，无需填写。
	ControlParams *CommonMixControlParams `json:"ControlParams,omitnil,omitempty" name:"ControlParams"`
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

// Predefined struct for user
type CreateCommonMixStreamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCommonMixStreamResponse struct {
	*tchttp.BaseResponse
	Response *CreateCommonMixStreamResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLiveCallbackRuleRequestParams struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 模板ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type CreateLiveCallbackRuleRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 模板ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type CreateLiveCallbackRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveCallbackRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveCallbackRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLiveCallbackTemplateRequestParams struct {
	// 模板名称。
	// 长度上限：255字节。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	// 长度上限：1024字节。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 开播回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitnil,omitempty" name:"StreamBeginNotifyUrl"`

	// 断流回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitnil,omitempty" name:"StreamEndNotifyUrl"`

	// 录制文件回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitnil,omitempty" name:"RecordNotifyUrl"`

	// 录制状态回调 URL ，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	RecordStatusNotifyUrl *string `json:"RecordStatusNotifyUrl,omitnil,omitempty" name:"RecordStatusNotifyUrl"`

	// 截图回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitnil,omitempty" name:"SnapshotNotifyUrl"`

	// 鉴黄回调 URL ，
	// 相关协议文档：[事件消息通知](/document/product/267/32741)。
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitnil,omitempty" name:"PornCensorshipNotifyUrl"`

	// 回调 Key，回调 URL 公用，回调签名详见事件消息通知文档。
	// [事件消息通知](/document/product/267/32744)。
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 参数已弃用。
	StreamMixNotifyUrl *string `json:"StreamMixNotifyUrl,omitnil,omitempty" name:"StreamMixNotifyUrl"`

	// 推流异常回调 URL。
	PushExceptionNotifyUrl *string `json:"PushExceptionNotifyUrl,omitnil,omitempty" name:"PushExceptionNotifyUrl"`

	// 音频审核回调 URL。
	AudioAuditNotifyUrl *string `json:"AudioAuditNotifyUrl,omitnil,omitempty" name:"AudioAuditNotifyUrl"`

	// 录制异常回调 URL。
	RecordExceptionNotifyUrl *string `json:"RecordExceptionNotifyUrl,omitnil,omitempty" name:"RecordExceptionNotifyUrl"`

	// 录制异常回调级别，可选择：error、warning、info。
	RecordExceptionLevels []*string `json:"RecordExceptionLevels,omitnil,omitempty" name:"RecordExceptionLevels"`
}

type CreateLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称。
	// 长度上限：255字节。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	// 长度上限：1024字节。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 开播回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitnil,omitempty" name:"StreamBeginNotifyUrl"`

	// 断流回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitnil,omitempty" name:"StreamEndNotifyUrl"`

	// 录制文件回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitnil,omitempty" name:"RecordNotifyUrl"`

	// 录制状态回调 URL ，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	RecordStatusNotifyUrl *string `json:"RecordStatusNotifyUrl,omitnil,omitempty" name:"RecordStatusNotifyUrl"`

	// 截图回调 URL，
	// 相关协议文档：[事件消息通知](/document/product/267/32744)。
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitnil,omitempty" name:"SnapshotNotifyUrl"`

	// 鉴黄回调 URL ，
	// 相关协议文档：[事件消息通知](/document/product/267/32741)。
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitnil,omitempty" name:"PornCensorshipNotifyUrl"`

	// 回调 Key，回调 URL 公用，回调签名详见事件消息通知文档。
	// [事件消息通知](/document/product/267/32744)。
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 参数已弃用。
	StreamMixNotifyUrl *string `json:"StreamMixNotifyUrl,omitnil,omitempty" name:"StreamMixNotifyUrl"`

	// 推流异常回调 URL。
	PushExceptionNotifyUrl *string `json:"PushExceptionNotifyUrl,omitnil,omitempty" name:"PushExceptionNotifyUrl"`

	// 音频审核回调 URL。
	AudioAuditNotifyUrl *string `json:"AudioAuditNotifyUrl,omitnil,omitempty" name:"AudioAuditNotifyUrl"`

	// 录制异常回调 URL。
	RecordExceptionNotifyUrl *string `json:"RecordExceptionNotifyUrl,omitnil,omitempty" name:"RecordExceptionNotifyUrl"`

	// 录制异常回调级别，可选择：error、warning、info。
	RecordExceptionLevels []*string `json:"RecordExceptionLevels,omitnil,omitempty" name:"RecordExceptionLevels"`
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
	delete(f, "RecordStatusNotifyUrl")
	delete(f, "SnapshotNotifyUrl")
	delete(f, "PornCensorshipNotifyUrl")
	delete(f, "CallbackKey")
	delete(f, "StreamMixNotifyUrl")
	delete(f, "PushExceptionNotifyUrl")
	delete(f, "AudioAuditNotifyUrl")
	delete(f, "RecordExceptionNotifyUrl")
	delete(f, "RecordExceptionLevels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveCallbackTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLiveCallbackTemplateResponseParams struct {
	// 模板ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveCallbackTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLivePadRuleRequestParams struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 注：如果本参数设置为非空字符串，规则将只对此推流起作用。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type CreateLivePadRuleRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 注：如果本参数设置为非空字符串，规则将只对此推流起作用。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

func (r *CreateLivePadRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLivePadRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "TemplateId")
	delete(f, "AppName")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLivePadRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLivePadRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLivePadRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateLivePadRuleResponseParams `json:"Response"`
}

func (r *CreateLivePadRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLivePadRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLivePadTemplateRequestParams struct {
	// 模板名称。
	// 长度上限：255字节。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 垫片内容。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 描述信息。
	// 长度上限：1024字节。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 断流等待时间。
	// 取值范围：0-30000。
	// 单位：ms。
	WaitDuration *uint64 `json:"WaitDuration,omitnil,omitempty" name:"WaitDuration"`

	// 最大垫片时长。
	// 取值范围：0 - 正无穷。
	// 单位：ms。
	MaxDuration *uint64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// 垫片内容类型：
	// 1：图片，2：视频。
	// 默认值：1。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type CreateLivePadTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称。
	// 长度上限：255字节。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 垫片内容。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 描述信息。
	// 长度上限：1024字节。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 断流等待时间。
	// 取值范围：0-30000。
	// 单位：ms。
	WaitDuration *uint64 `json:"WaitDuration,omitnil,omitempty" name:"WaitDuration"`

	// 最大垫片时长。
	// 取值范围：0 - 正无穷。
	// 单位：ms。
	MaxDuration *uint64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// 垫片内容类型：
	// 1：图片，2：视频。
	// 默认值：1。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *CreateLivePadTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLivePadTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "Url")
	delete(f, "Description")
	delete(f, "WaitDuration")
	delete(f, "MaxDuration")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLivePadTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLivePadTemplateResponseParams struct {
	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLivePadTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateLivePadTemplateResponseParams `json:"Response"`
}

func (r *CreateLivePadTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLivePadTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLivePullStreamTaskRequestParams struct {
	// 拉流源的类型：
	// PullLivePushLive -直播，
	// PullVodPushLive -点播，
	// PullPicPushLive -图片。
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 拉流源 url 列表。
	// SourceType 为直播（PullLivePushLive）只可以填1个，
	// SourceType 为点播（PullVodPushLive）可以填多个，上限30个。
	// 当前支持的文件格式：flv，mp4，hls。
	// 当前支持的拉流协议：http，https，rtmp，rtmps，rtsp，srt。
	// 注意：
	// 1. 建议优先使用 flv 文件，对于 mp4 未交织好的文件轮播推流易产生卡顿以及因为频繁拉取导致源产生大量源出口带宽成本，可通过点播转码进行重新交织后再轮播，或提前创建任务并开启本地模式。
	// 2. 拒绝内网域名等攻击性拉流地址，如有使用，则做账号封禁处理。
	// 3. 源文件请保持时间戳正常交织递增，避免因源文件异常影响推流及播放。
	// 4. 视频编码格式仅支持: H264, H265。
	// 5. 音频编码格式仅支持: AAC。
	// 6. 点播源请使用小文件，尽量时长保持在1小时内，较大文件打开和续播耗时较久，耗时超过15秒会有无法正常转推风险。
	// 7. 避免使用低频存储的文件，该类文件因低频存储，拉取时容易出现慢速，影响拉转推质量。
	SourceUrls []*string `json:"SourceUrls,omitnil,omitempty" name:"SourceUrls"`

	// 推流域名。
	// 将拉取过来的流推到该域名。
	// 注意：如果目标地址为非云直播，且样式不同于云直播，请使用 ToUrl 传入完整推流地址，详细用法请参考 ToUrl 参数说明。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	// 将拉取过来的流推到该路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流名称。
	// 将拉取过来的流推到该流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 开始时间。
	// 使用 UTC 格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，注意：1. 结束时间必须大于开始时间；2. 结束时间必须大于当前时间；3. 结束时间 和 开始时间 间隔必须小于30天。使用 UTC 格式时间，例如：2019-01-08T10:00:00Z。注意：北京时间值为 UTC 时间值 + 8 小时。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务操作人备注。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 推流参数。
	// 推流时携带自定义参数。
	// 示例：
	// bak=1&test=2 。
	PushArgs *string `json:"PushArgs,omitnil,omitempty" name:"PushArgs"`

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
	CallbackEvents []*string `json:"CallbackEvents,omitnil,omitempty" name:"CallbackEvents"`

	// 点播拉流转推循环次数。默认：-1。
	// -1：无限循环，直到任务结束。
	// 0：不循环。
	// >0：具体循环次数。次数和时间以先结束的为准。
	// 注意：该配置仅对拉流源为点播时生效。
	VodLoopTimes *string `json:"VodLoopTimes,omitnil,omitempty" name:"VodLoopTimes"`

	// 点播更新SourceUrls后的播放方式：
	// ImmediateNewSource：立即播放新的拉流源内容；
	// ContinueBreakPoint：播放完当前正在播放的点播 url 后再使用新的拉流源播放。（旧拉流源未播放的点播 url 不会再播放）
	// 
	// 注意：该配置生效仅对变更前拉流源为点播时生效。
	VodRefreshType *string `json:"VodRefreshType,omitnil,omitempty" name:"VodRefreshType"`

	// 自定义回调地址。
	// 拉流转推任务相关事件会回调到该地址。
	// 回调事件使用方法请查看：
	// https://cloud.tencent.com/document/product/267/32744
	// https://cloud.tencent.com/document/product/267/56208
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 其他参数。
	// 示例: ignore_region  用于忽略传入地域, 内部按负载分配。
	ExtraCmd *string `json:"ExtraCmd,omitnil,omitempty" name:"ExtraCmd"`

	// 自定义任务 ID。
	// 注：
	// 1. 该自定义 ID 为可选参数，如果传入，请确保该账号下传入的 ID 唯一。
	// 2. 该自定义 ID 用于防止重复发起请求时产生重复任务。后面也可以用 SpecifyTaskId 来修改或删除任务。
	SpecifyTaskId *string `json:"SpecifyTaskId,omitnil,omitempty" name:"SpecifyTaskId"`

	// 任务描述，限制 512 字节。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 完整目标 URL 地址。
	// 用法注意：如果使用该参数来传完整目标地址，则 DomainName, AppName, StreamName 需要传入空字符串，任务将会使用该 ToUrl 参数指定的目标地址。
	// 
	// 使用该方式传入目标地址支持的协议有：
	// rtmp、rtmps、rtsp、rtp、srt。
	// 
	// 注意：签名时间需要超过任务结束时间，避免因签名过期造成任务失败。
	ToUrl *string `json:"ToUrl,omitnil,omitempty" name:"ToUrl"`

	// 指定播放文件索引。
	// 注意： 1. 从1开始，不大于SourceUrls中文件个数。
	// 2. 该偏移仅在首次轮播时有效。
	// 3. 提前创建的任务指定的偏移最长有效期为24小时，24小时后未开始的任务偏移失效。
	FileIndex *int64 `json:"FileIndex,omitnil,omitempty" name:"FileIndex"`

	// 指定播放文件偏移。
	// 注意：
	// 1. 单位：秒，配合FileIndex使用。
	OffsetTime *int64 `json:"OffsetTime,omitnil,omitempty" name:"OffsetTime"`

	// 备源的类型：
	// PullLivePushLive -直播，
	// PullVodPushLive -点播。
	// 注意：
	// 1. 仅当主源类型为直播源时，备源才会生效。
	// 2. 主直播源拉流中断时，自动使用备源进行拉流。
	// 3. 如果备源为点播文件时，则每次轮播完点播文件就检查主源是否恢复，如果主源恢复则自动切回到主源，否则继续拉备源。
	BackupSourceType *string `json:"BackupSourceType,omitnil,omitempty" name:"BackupSourceType"`

	// 备源 URL。
	// 只允许填一个备源 URL
	BackupSourceUrl *string `json:"BackupSourceUrl,omitnil,omitempty" name:"BackupSourceUrl"`

	// 水印信息列表。
	// 注意：
	// 1. 最多支持4个不同位置的水印。
	// 2. 水印图片 URL 请使用合法外网可访问地址。
	// 3. 支持的水印图片格式：png，jpg，gif 等。
	WatermarkList []*PullPushWatermarkInfo `json:"WatermarkList,omitnil,omitempty" name:"WatermarkList"`

	// 点播源是否启用本地推流模式，默认0，不启用。
	// 0 - 不启用。
	// 1 - 启用。
	// 注意：启用本地模式后，会将源列表中的 MP4 文件进行本地下载，优先使用本地已下载文件进行推流，提高点播源推流稳定性。使用本地下载文件推流时，会产生增值费用。
	VodLocalMode *int64 `json:"VodLocalMode,omitnil,omitempty" name:"VodLocalMode"`

	// 录制模板 ID。
	RecordTemplateId *string `json:"RecordTemplateId,omitnil,omitempty" name:"RecordTemplateId"`

	// 新的目标地址，用于任务同时推两路场景。
	BackupToUrl *string `json:"BackupToUrl,omitnil,omitempty" name:"BackupToUrl"`

	// 直播转码模板，使用云直播的转码功能进行转码后再转推出去。转码模板需在云直播控制台创建。
	TranscodeTemplateName *string `json:"TranscodeTemplateName,omitnil,omitempty" name:"TranscodeTemplateName"`
}

type CreateLivePullStreamTaskRequest struct {
	*tchttp.BaseRequest
	
	// 拉流源的类型：
	// PullLivePushLive -直播，
	// PullVodPushLive -点播，
	// PullPicPushLive -图片。
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 拉流源 url 列表。
	// SourceType 为直播（PullLivePushLive）只可以填1个，
	// SourceType 为点播（PullVodPushLive）可以填多个，上限30个。
	// 当前支持的文件格式：flv，mp4，hls。
	// 当前支持的拉流协议：http，https，rtmp，rtmps，rtsp，srt。
	// 注意：
	// 1. 建议优先使用 flv 文件，对于 mp4 未交织好的文件轮播推流易产生卡顿以及因为频繁拉取导致源产生大量源出口带宽成本，可通过点播转码进行重新交织后再轮播，或提前创建任务并开启本地模式。
	// 2. 拒绝内网域名等攻击性拉流地址，如有使用，则做账号封禁处理。
	// 3. 源文件请保持时间戳正常交织递增，避免因源文件异常影响推流及播放。
	// 4. 视频编码格式仅支持: H264, H265。
	// 5. 音频编码格式仅支持: AAC。
	// 6. 点播源请使用小文件，尽量时长保持在1小时内，较大文件打开和续播耗时较久，耗时超过15秒会有无法正常转推风险。
	// 7. 避免使用低频存储的文件，该类文件因低频存储，拉取时容易出现慢速，影响拉转推质量。
	SourceUrls []*string `json:"SourceUrls,omitnil,omitempty" name:"SourceUrls"`

	// 推流域名。
	// 将拉取过来的流推到该域名。
	// 注意：如果目标地址为非云直播，且样式不同于云直播，请使用 ToUrl 传入完整推流地址，详细用法请参考 ToUrl 参数说明。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	// 将拉取过来的流推到该路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流名称。
	// 将拉取过来的流推到该流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 开始时间。
	// 使用 UTC 格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，注意：1. 结束时间必须大于开始时间；2. 结束时间必须大于当前时间；3. 结束时间 和 开始时间 间隔必须小于30天。使用 UTC 格式时间，例如：2019-01-08T10:00:00Z。注意：北京时间值为 UTC 时间值 + 8 小时。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务操作人备注。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 推流参数。
	// 推流时携带自定义参数。
	// 示例：
	// bak=1&test=2 。
	PushArgs *string `json:"PushArgs,omitnil,omitempty" name:"PushArgs"`

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
	CallbackEvents []*string `json:"CallbackEvents,omitnil,omitempty" name:"CallbackEvents"`

	// 点播拉流转推循环次数。默认：-1。
	// -1：无限循环，直到任务结束。
	// 0：不循环。
	// >0：具体循环次数。次数和时间以先结束的为准。
	// 注意：该配置仅对拉流源为点播时生效。
	VodLoopTimes *string `json:"VodLoopTimes,omitnil,omitempty" name:"VodLoopTimes"`

	// 点播更新SourceUrls后的播放方式：
	// ImmediateNewSource：立即播放新的拉流源内容；
	// ContinueBreakPoint：播放完当前正在播放的点播 url 后再使用新的拉流源播放。（旧拉流源未播放的点播 url 不会再播放）
	// 
	// 注意：该配置生效仅对变更前拉流源为点播时生效。
	VodRefreshType *string `json:"VodRefreshType,omitnil,omitempty" name:"VodRefreshType"`

	// 自定义回调地址。
	// 拉流转推任务相关事件会回调到该地址。
	// 回调事件使用方法请查看：
	// https://cloud.tencent.com/document/product/267/32744
	// https://cloud.tencent.com/document/product/267/56208
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 其他参数。
	// 示例: ignore_region  用于忽略传入地域, 内部按负载分配。
	ExtraCmd *string `json:"ExtraCmd,omitnil,omitempty" name:"ExtraCmd"`

	// 自定义任务 ID。
	// 注：
	// 1. 该自定义 ID 为可选参数，如果传入，请确保该账号下传入的 ID 唯一。
	// 2. 该自定义 ID 用于防止重复发起请求时产生重复任务。后面也可以用 SpecifyTaskId 来修改或删除任务。
	SpecifyTaskId *string `json:"SpecifyTaskId,omitnil,omitempty" name:"SpecifyTaskId"`

	// 任务描述，限制 512 字节。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 完整目标 URL 地址。
	// 用法注意：如果使用该参数来传完整目标地址，则 DomainName, AppName, StreamName 需要传入空字符串，任务将会使用该 ToUrl 参数指定的目标地址。
	// 
	// 使用该方式传入目标地址支持的协议有：
	// rtmp、rtmps、rtsp、rtp、srt。
	// 
	// 注意：签名时间需要超过任务结束时间，避免因签名过期造成任务失败。
	ToUrl *string `json:"ToUrl,omitnil,omitempty" name:"ToUrl"`

	// 指定播放文件索引。
	// 注意： 1. 从1开始，不大于SourceUrls中文件个数。
	// 2. 该偏移仅在首次轮播时有效。
	// 3. 提前创建的任务指定的偏移最长有效期为24小时，24小时后未开始的任务偏移失效。
	FileIndex *int64 `json:"FileIndex,omitnil,omitempty" name:"FileIndex"`

	// 指定播放文件偏移。
	// 注意：
	// 1. 单位：秒，配合FileIndex使用。
	OffsetTime *int64 `json:"OffsetTime,omitnil,omitempty" name:"OffsetTime"`

	// 备源的类型：
	// PullLivePushLive -直播，
	// PullVodPushLive -点播。
	// 注意：
	// 1. 仅当主源类型为直播源时，备源才会生效。
	// 2. 主直播源拉流中断时，自动使用备源进行拉流。
	// 3. 如果备源为点播文件时，则每次轮播完点播文件就检查主源是否恢复，如果主源恢复则自动切回到主源，否则继续拉备源。
	BackupSourceType *string `json:"BackupSourceType,omitnil,omitempty" name:"BackupSourceType"`

	// 备源 URL。
	// 只允许填一个备源 URL
	BackupSourceUrl *string `json:"BackupSourceUrl,omitnil,omitempty" name:"BackupSourceUrl"`

	// 水印信息列表。
	// 注意：
	// 1. 最多支持4个不同位置的水印。
	// 2. 水印图片 URL 请使用合法外网可访问地址。
	// 3. 支持的水印图片格式：png，jpg，gif 等。
	WatermarkList []*PullPushWatermarkInfo `json:"WatermarkList,omitnil,omitempty" name:"WatermarkList"`

	// 点播源是否启用本地推流模式，默认0，不启用。
	// 0 - 不启用。
	// 1 - 启用。
	// 注意：启用本地模式后，会将源列表中的 MP4 文件进行本地下载，优先使用本地已下载文件进行推流，提高点播源推流稳定性。使用本地下载文件推流时，会产生增值费用。
	VodLocalMode *int64 `json:"VodLocalMode,omitnil,omitempty" name:"VodLocalMode"`

	// 录制模板 ID。
	RecordTemplateId *string `json:"RecordTemplateId,omitnil,omitempty" name:"RecordTemplateId"`

	// 新的目标地址，用于任务同时推两路场景。
	BackupToUrl *string `json:"BackupToUrl,omitnil,omitempty" name:"BackupToUrl"`

	// 直播转码模板，使用云直播的转码功能进行转码后再转推出去。转码模板需在云直播控制台创建。
	TranscodeTemplateName *string `json:"TranscodeTemplateName,omitnil,omitempty" name:"TranscodeTemplateName"`
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
	delete(f, "SpecifyTaskId")
	delete(f, "Comment")
	delete(f, "ToUrl")
	delete(f, "FileIndex")
	delete(f, "OffsetTime")
	delete(f, "BackupSourceType")
	delete(f, "BackupSourceUrl")
	delete(f, "WatermarkList")
	delete(f, "VodLocalMode")
	delete(f, "RecordTemplateId")
	delete(f, "BackupToUrl")
	delete(f, "TranscodeTemplateName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLivePullStreamTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLivePullStreamTaskResponseParams struct {
	// 任务 Id 。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLivePullStreamTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateLivePullStreamTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLiveRecordRequestParams struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流路径，与推流和播放地址中的 AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流域名。多域名推流必须设置。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 录制开始时间。中国标准时间，需要 URLEncode(rfc3986)。如 2017-01-01 10:10:01，编码为：2017-01-01+10%3a10%3a01。
	// 定时录制模式，必须设置该字段；实时视频录制模式，忽略该字段。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 录制结束时间。中国标准时间，需要 URLEncode(rfc3986)。如 2017-01-01 10:30:01，编码为：2017-01-01+10%3a30%3a01。
	// 定时录制模式，必须设置该字段；实时录制模式，为可选字段。如果通过Highlight参数，设置录制为实时视频录制模式，其设置的结束时间不应超过当前时间+30分钟，如果设置的结束时间超过当前时间+30分钟或者小于当前时间或者不设置该参数，则实际结束时间为当前时间+30分钟。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 录制类型。
	// “video” : 音视频录制【默认】。
	// “audio” : 纯音频录制。
	// 在定时录制模式或实时视频录制模式下，该参数均有效，不区分大小写。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 录制文件格式。其值为：
	// “flv”【默认】,“hls”,”mp4”,“aac”,”mp3”。
	// 在定时录制模式或实时视频录制模式下，该参数均有效，不区分大小写。
	FileFormat *string `json:"FileFormat,omitnil,omitempty" name:"FileFormat"`

	// 开启实时视频录制模式标志。
	// 0：不开启实时视频录制模式，即定时录制模式【默认】。见[示例一](#.E7.A4.BA.E4.BE.8B1-.E5.88.9B.E5.BB.BA.E5.AE.9A.E6.97.B6.E5.BD.95.E5.88.B6.E4.BB.BB.E5.8A.A1)。
	// 1：开启实时视频录制模式。见[示例二](#.E7.A4.BA.E4.BE.8B2-.E5.88.9B.E5.BB.BA.E5.AE.9E.E6.97.B6.E5.BD.95.E5.88.B6.E4.BB.BB.E5.8A.A1)。
	Highlight *int64 `json:"Highlight,omitnil,omitempty" name:"Highlight"`

	// 开启 A+B=C混流C流录制标志。
	// 0：不开启 A+B=C混流C流录制【默认】。
	// 1：开启 A+B=C混流C流录制。
	// 在定时录制模式或实时视频录制模式下，该参数均有效。
	MixStream *int64 `json:"MixStream,omitnil,omitempty" name:"MixStream"`

	// 录制流参数。当前支持以下参数：
	// record_interval - 录制分片时长，单位 秒，1800 - 7200。
	// storage_time - 录制文件存储时长，单位 秒。
	// eg. record_interval=3600&storage_time=2592000。
	// 注：参数需要url encode。
	// 在定时录制模式或实时视频录制模式下，该参数均有效。
	StreamParam *string `json:"StreamParam,omitnil,omitempty" name:"StreamParam"`
}

type CreateLiveRecordRequest struct {
	*tchttp.BaseRequest
	
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流路径，与推流和播放地址中的 AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流域名。多域名推流必须设置。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 录制开始时间。中国标准时间，需要 URLEncode(rfc3986)。如 2017-01-01 10:10:01，编码为：2017-01-01+10%3a10%3a01。
	// 定时录制模式，必须设置该字段；实时视频录制模式，忽略该字段。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 录制结束时间。中国标准时间，需要 URLEncode(rfc3986)。如 2017-01-01 10:30:01，编码为：2017-01-01+10%3a30%3a01。
	// 定时录制模式，必须设置该字段；实时录制模式，为可选字段。如果通过Highlight参数，设置录制为实时视频录制模式，其设置的结束时间不应超过当前时间+30分钟，如果设置的结束时间超过当前时间+30分钟或者小于当前时间或者不设置该参数，则实际结束时间为当前时间+30分钟。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 录制类型。
	// “video” : 音视频录制【默认】。
	// “audio” : 纯音频录制。
	// 在定时录制模式或实时视频录制模式下，该参数均有效，不区分大小写。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 录制文件格式。其值为：
	// “flv”【默认】,“hls”,”mp4”,“aac”,”mp3”。
	// 在定时录制模式或实时视频录制模式下，该参数均有效，不区分大小写。
	FileFormat *string `json:"FileFormat,omitnil,omitempty" name:"FileFormat"`

	// 开启实时视频录制模式标志。
	// 0：不开启实时视频录制模式，即定时录制模式【默认】。见[示例一](#.E7.A4.BA.E4.BE.8B1-.E5.88.9B.E5.BB.BA.E5.AE.9A.E6.97.B6.E5.BD.95.E5.88.B6.E4.BB.BB.E5.8A.A1)。
	// 1：开启实时视频录制模式。见[示例二](#.E7.A4.BA.E4.BE.8B2-.E5.88.9B.E5.BB.BA.E5.AE.9E.E6.97.B6.E5.BD.95.E5.88.B6.E4.BB.BB.E5.8A.A1)。
	Highlight *int64 `json:"Highlight,omitnil,omitempty" name:"Highlight"`

	// 开启 A+B=C混流C流录制标志。
	// 0：不开启 A+B=C混流C流录制【默认】。
	// 1：开启 A+B=C混流C流录制。
	// 在定时录制模式或实时视频录制模式下，该参数均有效。
	MixStream *int64 `json:"MixStream,omitnil,omitempty" name:"MixStream"`

	// 录制流参数。当前支持以下参数：
	// record_interval - 录制分片时长，单位 秒，1800 - 7200。
	// storage_time - 录制文件存储时长，单位 秒。
	// eg. record_interval=3600&storage_time=2592000。
	// 注：参数需要url encode。
	// 在定时录制模式或实时视频录制模式下，该参数均有效。
	StreamParam *string `json:"StreamParam,omitnil,omitempty" name:"StreamParam"`
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

// Predefined struct for user
type CreateLiveRecordResponseParams struct {
	// 任务 ID，全局唯一标识录制任务。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveRecordResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLiveRecordRuleRequestParams struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 注：如果本参数设置为非空字符串，规则将只对此推流起作用。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type CreateLiveRecordRuleRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 注：如果本参数设置为非空字符串，规则将只对此推流起作用。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
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

// Predefined struct for user
type CreateLiveRecordRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveRecordRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveRecordRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLiveRecordTemplateRequestParams struct {
	// 模板名。仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Flv录制参数，开启Flv录制时设置。
	FlvParam *RecordParam `json:"FlvParam,omitnil,omitempty" name:"FlvParam"`

	// Hls录制参数，开启hls录制时设置。
	HlsParam *RecordParam `json:"HlsParam,omitnil,omitempty" name:"HlsParam"`

	// Mp4录制参数，开启Mp4录制时设置。
	Mp4Param *RecordParam `json:"Mp4Param,omitnil,omitempty" name:"Mp4Param"`

	// Aac录制参数，开启Aac录制时设置。
	AacParam *RecordParam `json:"AacParam,omitnil,omitempty" name:"AacParam"`

	// 直播类型，默认 0。
	// 0：普通直播，
	// 1：慢直播。
	IsDelayLive *int64 `json:"IsDelayLive,omitnil,omitempty" name:"IsDelayLive"`

	// HLS专属录制参数。
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitnil,omitempty" name:"HlsSpecialParam"`

	// Mp3录制参数，开启Mp3录制时设置。
	Mp3Param *RecordParam `json:"Mp3Param,omitnil,omitempty" name:"Mp3Param"`

	// 是否存储至 cos，值为 1 时表示存储至 cos。
	CosStore *int64 `json:"CosStore,omitnil,omitempty" name:"CosStore"`

	// 是否去除水印，类型为慢直播时此参数无效。
	// 如果为false，则录制水印流或转码流；如果为true，则录制原始流。
	RemoveWatermark *bool `json:"RemoveWatermark,omitnil,omitempty" name:"RemoveWatermark"`

	// FLV 录制特殊参数。
	FlvSpecialParam *FlvSpecialParam `json:"FlvSpecialParam,omitnil,omitempty" name:"FlvSpecialParam"`
}

type CreateLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名。仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Flv录制参数，开启Flv录制时设置。
	FlvParam *RecordParam `json:"FlvParam,omitnil,omitempty" name:"FlvParam"`

	// Hls录制参数，开启hls录制时设置。
	HlsParam *RecordParam `json:"HlsParam,omitnil,omitempty" name:"HlsParam"`

	// Mp4录制参数，开启Mp4录制时设置。
	Mp4Param *RecordParam `json:"Mp4Param,omitnil,omitempty" name:"Mp4Param"`

	// Aac录制参数，开启Aac录制时设置。
	AacParam *RecordParam `json:"AacParam,omitnil,omitempty" name:"AacParam"`

	// 直播类型，默认 0。
	// 0：普通直播，
	// 1：慢直播。
	IsDelayLive *int64 `json:"IsDelayLive,omitnil,omitempty" name:"IsDelayLive"`

	// HLS专属录制参数。
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitnil,omitempty" name:"HlsSpecialParam"`

	// Mp3录制参数，开启Mp3录制时设置。
	Mp3Param *RecordParam `json:"Mp3Param,omitnil,omitempty" name:"Mp3Param"`

	// 是否存储至 cos，值为 1 时表示存储至 cos。
	CosStore *int64 `json:"CosStore,omitnil,omitempty" name:"CosStore"`

	// 是否去除水印，类型为慢直播时此参数无效。
	// 如果为false，则录制水印流或转码流；如果为true，则录制原始流。
	RemoveWatermark *bool `json:"RemoveWatermark,omitnil,omitempty" name:"RemoveWatermark"`

	// FLV 录制特殊参数。
	FlvSpecialParam *FlvSpecialParam `json:"FlvSpecialParam,omitnil,omitempty" name:"FlvSpecialParam"`
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
	delete(f, "CosStore")
	delete(f, "RemoveWatermark")
	delete(f, "FlvSpecialParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLiveRecordTemplateResponseParams struct {
	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveRecordTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLiveSnapshotRuleRequestParams struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 注：如果本参数设置为非空字符串，规则将只对此推流起作用。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type CreateLiveSnapshotRuleRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 注：如果本参数设置为非空字符串，规则将只对此推流起作用。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
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

// Predefined struct for user
type CreateLiveSnapshotRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveSnapshotRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveSnapshotRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLiveSnapshotTemplateRequestParams struct {
	// 模板名称。
	// 长度上限：255字节。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Cos 应用 ID。
	CosAppId *int64 `json:"CosAppId,omitnil,omitempty" name:"CosAppId"`

	// Cos Bucket名称。
	// 注：CosBucket参数值不能包含-[appid] 部分。
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// Cos地区。
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 描述信息。
	// 长度上限：1024字节。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 截图间隔，单位s，默认10s。
	// 范围： 2s ~ 300s。
	SnapshotInterval *int64 `json:"SnapshotInterval,omitnil,omitempty" name:"SnapshotInterval"`

	// 截图宽度。默认：0（原始宽）。
	// 范围：0-3000 。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度。默认：0（原始高）。
	// 范围：0-2000 。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 是否开启鉴黄，0：不开启，1：开启。默认：0。
	PornFlag *int64 `json:"PornFlag,omitnil,omitempty" name:"PornFlag"`

	// Cos Bucket文件夹前缀。
	// 如不传，实际按默认值
	// /{Year}-{Month}-{Day}
	// 生效
	CosPrefix *string `json:"CosPrefix,omitnil,omitempty" name:"CosPrefix"`

	// Cos 文件名称。
	// 如不传，实际按默认值
	// {StreamID}-screenshot-{Hour}-{Minute}-{Second}-{Width}x{Height}{Ext}
	// 生效
	CosFileName *string `json:"CosFileName,omitnil,omitempty" name:"CosFileName"`
}

type CreateLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称。
	// 长度上限：255字节。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Cos 应用 ID。
	CosAppId *int64 `json:"CosAppId,omitnil,omitempty" name:"CosAppId"`

	// Cos Bucket名称。
	// 注：CosBucket参数值不能包含-[appid] 部分。
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// Cos地区。
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 描述信息。
	// 长度上限：1024字节。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 截图间隔，单位s，默认10s。
	// 范围： 2s ~ 300s。
	SnapshotInterval *int64 `json:"SnapshotInterval,omitnil,omitempty" name:"SnapshotInterval"`

	// 截图宽度。默认：0（原始宽）。
	// 范围：0-3000 。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度。默认：0（原始高）。
	// 范围：0-2000 。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 是否开启鉴黄，0：不开启，1：开启。默认：0。
	PornFlag *int64 `json:"PornFlag,omitnil,omitempty" name:"PornFlag"`

	// Cos Bucket文件夹前缀。
	// 如不传，实际按默认值
	// /{Year}-{Month}-{Day}
	// 生效
	CosPrefix *string `json:"CosPrefix,omitnil,omitempty" name:"CosPrefix"`

	// Cos 文件名称。
	// 如不传，实际按默认值
	// {StreamID}-screenshot-{Hour}-{Minute}-{Second}-{Width}x{Height}{Ext}
	// 生效
	CosFileName *string `json:"CosFileName,omitnil,omitempty" name:"CosFileName"`
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

// Predefined struct for user
type CreateLiveSnapshotTemplateResponseParams struct {
	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveSnapshotTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLiveStreamMonitorRequestParams struct {
	// 监播任务的输出信息。
	OutputInfo *LiveStreamMonitorOutputInfo `json:"OutputInfo,omitnil,omitempty" name:"OutputInfo"`

	// 待监播的输入流信息列表。
	InputList []*LiveStreamMonitorInputInfo `json:"InputList,omitnil,omitempty" name:"InputList"`

	// 监播任务名称。字段长度小于128字节（一个汉字两个字节）。
	MonitorName *string `json:"MonitorName,omitnil,omitempty" name:"MonitorName"`

	// 监播事件通知策略。
	// 不填默认为没有任何通知。
	NotifyPolicy *LiveStreamMonitorNotifyPolicy `json:"NotifyPolicy,omitnil,omitempty" name:"NotifyPolicy"`

	// 智能语音识别语种设置：
	// 0 关闭 1 中文 2 英文 3 日文 4 韩文。
	AsrLanguage *uint64 `json:"AsrLanguage,omitnil,omitempty" name:"AsrLanguage"`

	// 智能文字识别语种设置：
	// 0 关闭 1 中、英文。
	OcrLanguage *uint64 `json:"OcrLanguage,omitnil,omitempty" name:"OcrLanguage"`

	// 智能语音识别的输入列表，若开启语音识别则必填。
	// （第1条输入流index为1）
	AiAsrInputIndexList []*uint64 `json:"AiAsrInputIndexList,omitnil,omitempty" name:"AiAsrInputIndexList"`

	// 智能文字识别的输入列表，若开启文字识别则必填。
	// （第1条输入流index为1）
	AiOcrInputIndexList []*uint64 `json:"AiOcrInputIndexList,omitnil,omitempty" name:"AiOcrInputIndexList"`

	// 是否开启断流检测。
	CheckStreamBroken *uint64 `json:"CheckStreamBroken,omitnil,omitempty" name:"CheckStreamBroken"`

	// 是否开启低帧率检测。
	CheckStreamLowFrameRate *uint64 `json:"CheckStreamLowFrameRate,omitnil,omitempty" name:"CheckStreamLowFrameRate"`

	// 是否存储监播事件到监播报告，以及是否允许查询监播报告。
	AllowMonitorReport *uint64 `json:"AllowMonitorReport,omitnil,omitempty" name:"AllowMonitorReport"`

	// 是否开启格式诊断。
	AiFormatDiagnose *uint64 `json:"AiFormatDiagnose,omitnil,omitempty" name:"AiFormatDiagnose"`

	// 是否开启内容质检。
	AiQualityControl *uint64 `json:"AiQualityControl,omitnil,omitempty" name:"AiQualityControl"`

	// 导播台监播对应的导播台场次id。
	CasterId *string `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 拉流转推监播任务对应的拉流转推场次id
	PullPushTaskId *string `json:"PullPushTaskId,omitnil,omitempty" name:"PullPushTaskId"`
}

type CreateLiveStreamMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 监播任务的输出信息。
	OutputInfo *LiveStreamMonitorOutputInfo `json:"OutputInfo,omitnil,omitempty" name:"OutputInfo"`

	// 待监播的输入流信息列表。
	InputList []*LiveStreamMonitorInputInfo `json:"InputList,omitnil,omitempty" name:"InputList"`

	// 监播任务名称。字段长度小于128字节（一个汉字两个字节）。
	MonitorName *string `json:"MonitorName,omitnil,omitempty" name:"MonitorName"`

	// 监播事件通知策略。
	// 不填默认为没有任何通知。
	NotifyPolicy *LiveStreamMonitorNotifyPolicy `json:"NotifyPolicy,omitnil,omitempty" name:"NotifyPolicy"`

	// 智能语音识别语种设置：
	// 0 关闭 1 中文 2 英文 3 日文 4 韩文。
	AsrLanguage *uint64 `json:"AsrLanguage,omitnil,omitempty" name:"AsrLanguage"`

	// 智能文字识别语种设置：
	// 0 关闭 1 中、英文。
	OcrLanguage *uint64 `json:"OcrLanguage,omitnil,omitempty" name:"OcrLanguage"`

	// 智能语音识别的输入列表，若开启语音识别则必填。
	// （第1条输入流index为1）
	AiAsrInputIndexList []*uint64 `json:"AiAsrInputIndexList,omitnil,omitempty" name:"AiAsrInputIndexList"`

	// 智能文字识别的输入列表，若开启文字识别则必填。
	// （第1条输入流index为1）
	AiOcrInputIndexList []*uint64 `json:"AiOcrInputIndexList,omitnil,omitempty" name:"AiOcrInputIndexList"`

	// 是否开启断流检测。
	CheckStreamBroken *uint64 `json:"CheckStreamBroken,omitnil,omitempty" name:"CheckStreamBroken"`

	// 是否开启低帧率检测。
	CheckStreamLowFrameRate *uint64 `json:"CheckStreamLowFrameRate,omitnil,omitempty" name:"CheckStreamLowFrameRate"`

	// 是否存储监播事件到监播报告，以及是否允许查询监播报告。
	AllowMonitorReport *uint64 `json:"AllowMonitorReport,omitnil,omitempty" name:"AllowMonitorReport"`

	// 是否开启格式诊断。
	AiFormatDiagnose *uint64 `json:"AiFormatDiagnose,omitnil,omitempty" name:"AiFormatDiagnose"`

	// 是否开启内容质检。
	AiQualityControl *uint64 `json:"AiQualityControl,omitnil,omitempty" name:"AiQualityControl"`

	// 导播台监播对应的导播台场次id。
	CasterId *string `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 拉流转推监播任务对应的拉流转推场次id
	PullPushTaskId *string `json:"PullPushTaskId,omitnil,omitempty" name:"PullPushTaskId"`
}

func (r *CreateLiveStreamMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveStreamMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OutputInfo")
	delete(f, "InputList")
	delete(f, "MonitorName")
	delete(f, "NotifyPolicy")
	delete(f, "AsrLanguage")
	delete(f, "OcrLanguage")
	delete(f, "AiAsrInputIndexList")
	delete(f, "AiOcrInputIndexList")
	delete(f, "CheckStreamBroken")
	delete(f, "CheckStreamLowFrameRate")
	delete(f, "AllowMonitorReport")
	delete(f, "AiFormatDiagnose")
	delete(f, "AiQualityControl")
	delete(f, "CasterId")
	delete(f, "PullPushTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveStreamMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLiveStreamMonitorResponseParams struct {
	// 监播任务ID。
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveStreamMonitorResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveStreamMonitorResponseParams `json:"Response"`
}

func (r *CreateLiveStreamMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveStreamMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLiveTimeShiftRuleRequestParams struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。注：如果需要对所有AppName生效，可填写“*”。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。注：如果本参数设置为非空字符串，规则将只对此推流起作用。如果需要对所有StreamName生效，可填写“*”。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type CreateLiveTimeShiftRuleRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。注：如果需要对所有AppName生效，可填写“*”。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。注：如果本参数设置为非空字符串，规则将只对此推流起作用。如果需要对所有StreamName生效，可填写“*”。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *CreateLiveTimeShiftRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveTimeShiftRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "StreamName")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveTimeShiftRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLiveTimeShiftRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveTimeShiftRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveTimeShiftRuleResponseParams `json:"Response"`
}

func (r *CreateLiveTimeShiftRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveTimeShiftRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLiveTimeShiftTemplateRequestParams struct {
	// 模板名称。
	// 长度上限：255字节。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 时移时长。
	// 单位：s。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 描述信息。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 地域。
	// Mainland：中国大陆。
	// Overseas：海外及港澳台地区。
	// 默认值：Mainland。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 分片时长。
	// 可取3-10。
	// 单位：s。
	// 默认值：5。
	ItemDuration *uint64 `json:"ItemDuration,omitnil,omitempty" name:"ItemDuration"`

	// 是否去除水印。
	// 传true则将录制原始流。
	// 默认值：false。
	RemoveWatermark *bool `json:"RemoveWatermark,omitnil,omitempty" name:"RemoveWatermark"`

	// 转码流id列表。
	// 此参数仅在 RemoveWatermark为false时生效。
	TranscodeTemplateIds []*int64 `json:"TranscodeTemplateIds,omitnil,omitempty" name:"TranscodeTemplateIds"`
}

type CreateLiveTimeShiftTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称。
	// 长度上限：255字节。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 时移时长。
	// 单位：s。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 描述信息。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 地域。
	// Mainland：中国大陆。
	// Overseas：海外及港澳台地区。
	// 默认值：Mainland。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 分片时长。
	// 可取3-10。
	// 单位：s。
	// 默认值：5。
	ItemDuration *uint64 `json:"ItemDuration,omitnil,omitempty" name:"ItemDuration"`

	// 是否去除水印。
	// 传true则将录制原始流。
	// 默认值：false。
	RemoveWatermark *bool `json:"RemoveWatermark,omitnil,omitempty" name:"RemoveWatermark"`

	// 转码流id列表。
	// 此参数仅在 RemoveWatermark为false时生效。
	TranscodeTemplateIds []*int64 `json:"TranscodeTemplateIds,omitnil,omitempty" name:"TranscodeTemplateIds"`
}

func (r *CreateLiveTimeShiftTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveTimeShiftTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "Duration")
	delete(f, "Description")
	delete(f, "Area")
	delete(f, "ItemDuration")
	delete(f, "RemoveWatermark")
	delete(f, "TranscodeTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveTimeShiftTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLiveTimeShiftTemplateResponseParams struct {
	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveTimeShiftTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveTimeShiftTemplateResponseParams `json:"Response"`
}

func (r *CreateLiveTimeShiftTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveTimeShiftTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLiveTranscodeRuleRequestParams struct {
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致。如果只绑定域名，则此处填""。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。已废弃该字段，请传入空字符串。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 指定已有的模板Id。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type CreateLiveTranscodeRuleRequest struct {
	*tchttp.BaseRequest
	
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致。如果只绑定域名，则此处填""。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。已废弃该字段，请传入空字符串。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 指定已有的模板Id。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type CreateLiveTranscodeRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveTranscodeRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveTranscodeRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLiveTranscodeTemplateRequestParams struct {
	// 模板名称，例： 900p 仅支持字母和数字的组合。
	// 长度限制：
	//   标准转码：1-10个字符
	//   极速高清转码：3-10个字符
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 视频码率。范围：0kbps - 8000kbps。
	// 0为保持原始码率。
	// 注: 转码模板有码率唯一要求，最终保存的码率可能与输入码率有所差别。
	VideoBitrate *int64 `json:"VideoBitrate,omitnil,omitempty" name:"VideoBitrate"`

	// 音频编码：aac，默认aac。
	// 注意：当前该参数未生效，待后续支持！
	Acodec *string `json:"Acodec,omitnil,omitempty" name:"Acodec"`

	// 音频码率，默认0。
	// 范围：0-500。
	AudioBitrate *int64 `json:"AudioBitrate,omitnil,omitempty" name:"AudioBitrate"`

	// 视频编码：h264/h265/origin，默认origin。
	// 
	// origin: 保持原始编码格式
	Vcodec *string `json:"Vcodec,omitnil,omitempty" name:"Vcodec"`

	// 模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否保留视频，0：否，1：是。默认1。
	NeedVideo *int64 `json:"NeedVideo,omitnil,omitempty" name:"NeedVideo"`

	// 宽，默认0。
	// 范围[0-3000]
	// 数值必须是2的倍数，0是原始宽度
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 是否保留音频，0：否，1：是。默认1。
	NeedAudio *int64 `json:"NeedAudio,omitnil,omitempty" name:"NeedAudio"`

	// 高，默认0。
	// 范围[0-3000]
	// 数值必须是2的倍数，0是原始高度。
	// 极速高清模板（AiTransCode = 1 的时候）必须传。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 帧率，默认0。
	// 范围0-60fps
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 关键帧间隔，单位：秒。
	// 默认原始的间隔
	// 范围2-6
	Gop *int64 `json:"Gop,omitnil,omitempty" name:"Gop"`

	// 旋转角度，默认0。
	// 可取值：0，90，180，270
	Rotate *int64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`

	// 编码质量：
	// baseline/main/high。默认baseline
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// 当设置的码率>原始码率时，是否以原始码率为准。
	// 0：否， 1：是
	// 默认 0。
	BitrateToOrig *int64 `json:"BitrateToOrig,omitnil,omitempty" name:"BitrateToOrig"`

	// 当设置的高度>原始高度时，是否以原始高度为准。
	// 0：否， 1：是
	// 默认 0。
	HeightToOrig *int64 `json:"HeightToOrig,omitnil,omitempty" name:"HeightToOrig"`

	// 当设置的帧率>原始帧率时，是否以原始帧率为准。
	// 0：否， 1：是
	// 默认 0。
	FpsToOrig *int64 `json:"FpsToOrig,omitnil,omitempty" name:"FpsToOrig"`

	// 是否是极速高清模板，0：否，1：是。默认0。
	AiTransCode *int64 `json:"AiTransCode,omitnil,omitempty" name:"AiTransCode"`

	// 极速高清视频码率压缩比。
	// 极速高清目标码率=VideoBitrate * (1-AdaptBitratePercent)
	// 
	// 取值范围：0.0到0.5
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitnil,omitempty" name:"AdaptBitratePercent"`

	// 是否以短边作为高度，0：否，1：是。默认0。
	ShortEdgeAsHeight *int64 `json:"ShortEdgeAsHeight,omitnil,omitempty" name:"ShortEdgeAsHeight"`

	// DRM 加密类型，可选值：fairplay、normalaes、widevine。
	// 不传递或者为空字符串，清空之前的DRM配置。
	DRMType *string `json:"DRMType,omitnil,omitempty" name:"DRMType"`

	// DRM 加密项，可选值：AUDIO、SD、HD、UHD1、UHD2，后四个为一组，同组中的内容只能选一个。
	// 不传递或者为空字符串，清空之前的DRM配置。
	DRMTracks *string `json:"DRMTracks,omitnil,omitempty" name:"DRMTracks"`
}

type CreateLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称，例： 900p 仅支持字母和数字的组合。
	// 长度限制：
	//   标准转码：1-10个字符
	//   极速高清转码：3-10个字符
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 视频码率。范围：0kbps - 8000kbps。
	// 0为保持原始码率。
	// 注: 转码模板有码率唯一要求，最终保存的码率可能与输入码率有所差别。
	VideoBitrate *int64 `json:"VideoBitrate,omitnil,omitempty" name:"VideoBitrate"`

	// 音频编码：aac，默认aac。
	// 注意：当前该参数未生效，待后续支持！
	Acodec *string `json:"Acodec,omitnil,omitempty" name:"Acodec"`

	// 音频码率，默认0。
	// 范围：0-500。
	AudioBitrate *int64 `json:"AudioBitrate,omitnil,omitempty" name:"AudioBitrate"`

	// 视频编码：h264/h265/origin，默认origin。
	// 
	// origin: 保持原始编码格式
	Vcodec *string `json:"Vcodec,omitnil,omitempty" name:"Vcodec"`

	// 模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否保留视频，0：否，1：是。默认1。
	NeedVideo *int64 `json:"NeedVideo,omitnil,omitempty" name:"NeedVideo"`

	// 宽，默认0。
	// 范围[0-3000]
	// 数值必须是2的倍数，0是原始宽度
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 是否保留音频，0：否，1：是。默认1。
	NeedAudio *int64 `json:"NeedAudio,omitnil,omitempty" name:"NeedAudio"`

	// 高，默认0。
	// 范围[0-3000]
	// 数值必须是2的倍数，0是原始高度。
	// 极速高清模板（AiTransCode = 1 的时候）必须传。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 帧率，默认0。
	// 范围0-60fps
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 关键帧间隔，单位：秒。
	// 默认原始的间隔
	// 范围2-6
	Gop *int64 `json:"Gop,omitnil,omitempty" name:"Gop"`

	// 旋转角度，默认0。
	// 可取值：0，90，180，270
	Rotate *int64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`

	// 编码质量：
	// baseline/main/high。默认baseline
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// 当设置的码率>原始码率时，是否以原始码率为准。
	// 0：否， 1：是
	// 默认 0。
	BitrateToOrig *int64 `json:"BitrateToOrig,omitnil,omitempty" name:"BitrateToOrig"`

	// 当设置的高度>原始高度时，是否以原始高度为准。
	// 0：否， 1：是
	// 默认 0。
	HeightToOrig *int64 `json:"HeightToOrig,omitnil,omitempty" name:"HeightToOrig"`

	// 当设置的帧率>原始帧率时，是否以原始帧率为准。
	// 0：否， 1：是
	// 默认 0。
	FpsToOrig *int64 `json:"FpsToOrig,omitnil,omitempty" name:"FpsToOrig"`

	// 是否是极速高清模板，0：否，1：是。默认0。
	AiTransCode *int64 `json:"AiTransCode,omitnil,omitempty" name:"AiTransCode"`

	// 极速高清视频码率压缩比。
	// 极速高清目标码率=VideoBitrate * (1-AdaptBitratePercent)
	// 
	// 取值范围：0.0到0.5
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitnil,omitempty" name:"AdaptBitratePercent"`

	// 是否以短边作为高度，0：否，1：是。默认0。
	ShortEdgeAsHeight *int64 `json:"ShortEdgeAsHeight,omitnil,omitempty" name:"ShortEdgeAsHeight"`

	// DRM 加密类型，可选值：fairplay、normalaes、widevine。
	// 不传递或者为空字符串，清空之前的DRM配置。
	DRMType *string `json:"DRMType,omitnil,omitempty" name:"DRMType"`

	// DRM 加密项，可选值：AUDIO、SD、HD、UHD1、UHD2，后四个为一组，同组中的内容只能选一个。
	// 不传递或者为空字符串，清空之前的DRM配置。
	DRMTracks *string `json:"DRMTracks,omitnil,omitempty" name:"DRMTracks"`
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
	delete(f, "DRMType")
	delete(f, "DRMTracks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLiveTranscodeTemplateResponseParams struct {
	// 模板Id。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveTranscodeTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLiveWatermarkRuleRequestParams struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 水印Id，即调用[AddLiveWatermark](/document/product/267/30154)接口返回的WatermarkId。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type CreateLiveWatermarkRuleRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 水印Id，即调用[AddLiveWatermark](/document/product/267/30154)接口返回的WatermarkId。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type CreateLiveWatermarkRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLiveWatermarkRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveWatermarkRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreatePullStreamConfigRequestParams struct {
	// 源 Url ，用于拉流的地址。目前可支持直播流及点播文件。
	// 注意：
	// 1. 多个点播url之间使用空格拼接。
	// 2. 目前上限支持10个url。
	// 3. 目前不支持https协议。
	// 4. 支持拉流文件格式：flv，rtmp，hls，mp4。
	FromUrl *string `json:"FromUrl,omitnil,omitempty" name:"FromUrl"`

	// 目的 Url ，用于推流的地址，目前限制该目标地址为腾讯域名。
	// 仅支持：rtmp 协议。
	ToUrl *string `json:"ToUrl,omitnil,omitempty" name:"ToUrl"`

	// 选择完成转拉推的服务所在区域:
	// 1-深圳，
	// 2-上海，
	// 3-天津，
	// 4-中国香港。
	AreaId *int64 `json:"AreaId,omitnil,omitempty" name:"AreaId"`

	// 选择完成转拉推服务使用的运营商网络：
	// 1-电信，
	// 2-移动，
	// 3-联通，
	// 4-其他。
	// 注：AreaId 为4的时候，IspId 只能为其他。
	IspId *int64 `json:"IspId,omitnil,omitempty" name:"IspId"`

	// 开始时间。
	// 使用 UTC 格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于七天。
	// 使用 UTC 格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type CreatePullStreamConfigRequest struct {
	*tchttp.BaseRequest
	
	// 源 Url ，用于拉流的地址。目前可支持直播流及点播文件。
	// 注意：
	// 1. 多个点播url之间使用空格拼接。
	// 2. 目前上限支持10个url。
	// 3. 目前不支持https协议。
	// 4. 支持拉流文件格式：flv，rtmp，hls，mp4。
	FromUrl *string `json:"FromUrl,omitnil,omitempty" name:"FromUrl"`

	// 目的 Url ，用于推流的地址，目前限制该目标地址为腾讯域名。
	// 仅支持：rtmp 协议。
	ToUrl *string `json:"ToUrl,omitnil,omitempty" name:"ToUrl"`

	// 选择完成转拉推的服务所在区域:
	// 1-深圳，
	// 2-上海，
	// 3-天津，
	// 4-中国香港。
	AreaId *int64 `json:"AreaId,omitnil,omitempty" name:"AreaId"`

	// 选择完成转拉推服务使用的运营商网络：
	// 1-电信，
	// 2-移动，
	// 3-联通，
	// 4-其他。
	// 注：AreaId 为4的时候，IspId 只能为其他。
	IspId *int64 `json:"IspId,omitnil,omitempty" name:"IspId"`

	// 开始时间。
	// 使用 UTC 格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于七天。
	// 使用 UTC 格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
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

// Predefined struct for user
type CreatePullStreamConfigResponseParams struct {
	// 配置成功后的 ID。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePullStreamConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreatePullStreamConfigResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRecordTaskRequestParams struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 录制任务结束时间，Unix时间戳。设置时间必须大于StartTime及当前时间，且小于当前时间+7天。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 录制任务开始时间，Unix时间戳。如果不填表示立即启动录制。EndTime - StartTime不能超过24小时。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 推流类型，默认0。取值：
	// 0-直播推流。
	// 1-合成流，即 A+B=C 类型混流。
	StreamType *uint64 `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// 录制模板ID，CreateLiveRecordTemplate 返回值。如果不填或者传入错误ID，则默认录制HLS格式、永久存储。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 扩展字段，暂无定义。默认为空。
	Extension *string `json:"Extension,omitnil,omitempty" name:"Extension"`
}

type CreateRecordTaskRequest struct {
	*tchttp.BaseRequest
	
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 录制任务结束时间，Unix时间戳。设置时间必须大于StartTime及当前时间，且小于当前时间+7天。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 录制任务开始时间，Unix时间戳。如果不填表示立即启动录制。EndTime - StartTime不能超过24小时。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 推流类型，默认0。取值：
	// 0-直播推流。
	// 1-合成流，即 A+B=C 类型混流。
	StreamType *uint64 `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// 录制模板ID，CreateLiveRecordTemplate 返回值。如果不填或者传入错误ID，则默认录制HLS格式、永久存储。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 扩展字段，暂无定义。默认为空。
	Extension *string `json:"Extension,omitnil,omitempty" name:"Extension"`
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

// Predefined struct for user
type CreateRecordTaskResponseParams struct {
	// 任务ID，全局唯一标识录制任务。返回TaskId字段说明录制任务创建成功。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRecordTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateRecordTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateScreenshotTaskRequestParams struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 截图任务结束时间，Unix时间戳。设置时间必须大于StartTime及当前时间，且小于当前时间+7天。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 截图模板ID，CreateLiveSnapshotTemplate 返回值。如果传入错误ID，则不拉起截图。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 截图任务开始时间，Unix时间戳。如果不填表示立即启动截图。EndTime - StartTime不能超过24小时。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 推流类型，默认0。取值：
	// 0-直播推流。
	// 1-合成流，即 A+B=C 类型混流。
	StreamType *uint64 `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// 扩展字段，暂无定义。默认为空。
	Extension *string `json:"Extension,omitnil,omitempty" name:"Extension"`
}

type CreateScreenshotTaskRequest struct {
	*tchttp.BaseRequest
	
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 截图任务结束时间，Unix时间戳。设置时间必须大于StartTime及当前时间，且小于当前时间+7天。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 截图模板ID，CreateLiveSnapshotTemplate 返回值。如果传入错误ID，则不拉起截图。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 截图任务开始时间，Unix时间戳。如果不填表示立即启动截图。EndTime - StartTime不能超过24小时。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 推流类型，默认0。取值：
	// 0-直播推流。
	// 1-合成流，即 A+B=C 类型混流。
	StreamType *uint64 `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// 扩展字段，暂无定义。默认为空。
	Extension *string `json:"Extension,omitnil,omitempty" name:"Extension"`
}

func (r *CreateScreenshotTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScreenshotTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StreamName")
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "EndTime")
	delete(f, "TemplateId")
	delete(f, "StartTime")
	delete(f, "StreamType")
	delete(f, "Extension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScreenshotTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScreenshotTaskResponseParams struct {
	// 任务ID，全局唯一标识截图任务。返回TaskId字段说明截图任务创建成功。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateScreenshotTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateScreenshotTaskResponseParams `json:"Response"`
}

func (r *CreateScreenshotTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScreenshotTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DayStreamPlayInfo struct {
	// 数据时间点，接口返回支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见ISO日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）yyyy-MM-dd HH:mm:ss：使用此格式时，默认代表北京时间。
	// 接口返回的时间格式和查询请求传入的时间格式一致。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 带宽（单位Mbps）。
	Bandwidth *float64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 流量 （单位MB）。
	Flux *float64 `json:"Flux,omitnil,omitempty" name:"Flux"`

	// 请求数。
	Request *uint64 `json:"Request,omitnil,omitempty" name:"Request"`

	// 在线人数。当前后端接口数据的统计方式与客户端侧不同，该数据仅供参考，具体数据请以业务客户端侧统计数据为准。
	Online *uint64 `json:"Online,omitnil,omitempty" name:"Online"`
}

type DelayInfo struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的 
	//  AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 延播时间，单位：秒。
	DelayInterval *uint64 `json:"DelayInterval,omitnil,omitempty" name:"DelayInterval"`

	// 创建时间，UTC 时间。
	// 注意：UTC时间和北京时间相差8小时。
	// 例如：2019-06-18T12:00:00Z（为北京时间 2019 年 6 月 18 日 20 点 0 分 0 秒）。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 过期时间，UTC 时间。
	// 注意：UTC时间和北京时间相差8小时。
	// 例如：2019-06-18T12:00:00Z（为北京时间 2019 年 6 月 18 日 20 点 0 分 0 秒）。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 当前状态:
	// -1：已过期。
	// 1： 生效中。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteCasterInputInfoRequestParams struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台输入Index。
	// 范围[0,20]
	// 注：该Index对应的输入源需存在。
	InputIndex *uint64 `json:"InputIndex,omitnil,omitempty" name:"InputIndex"`
}

type DeleteCasterInputInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台输入Index。
	// 范围[0,20]
	// 注：该Index对应的输入源需存在。
	InputIndex *uint64 `json:"InputIndex,omitnil,omitempty" name:"InputIndex"`
}

func (r *DeleteCasterInputInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCasterInputInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "InputIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCasterInputInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCasterInputInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCasterInputInfoResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCasterInputInfoResponseParams `json:"Response"`
}

func (r *DeleteCasterInputInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCasterInputInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCasterLayoutInfoRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 要删除的布局Index。
	// 注：待删除的Index对应的布局需存在。
	LayoutIndex *uint64 `json:"LayoutIndex,omitnil,omitempty" name:"LayoutIndex"`
}

type DeleteCasterLayoutInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 要删除的布局Index。
	// 注：待删除的Index对应的布局需存在。
	LayoutIndex *uint64 `json:"LayoutIndex,omitnil,omitempty" name:"LayoutIndex"`
}

func (r *DeleteCasterLayoutInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCasterLayoutInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "LayoutIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCasterLayoutInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCasterLayoutInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCasterLayoutInfoResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCasterLayoutInfoResponseParams `json:"Response"`
}

func (r *DeleteCasterLayoutInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCasterLayoutInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCasterMarkPicInfoRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 需要删除的水印Index。
	MarkPicIndex *int64 `json:"MarkPicIndex,omitnil,omitempty" name:"MarkPicIndex"`
}

type DeleteCasterMarkPicInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 需要删除的水印Index。
	MarkPicIndex *int64 `json:"MarkPicIndex,omitnil,omitempty" name:"MarkPicIndex"`
}

func (r *DeleteCasterMarkPicInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCasterMarkPicInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "MarkPicIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCasterMarkPicInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCasterMarkPicInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCasterMarkPicInfoResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCasterMarkPicInfoResponseParams `json:"Response"`
}

func (r *DeleteCasterMarkPicInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCasterMarkPicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCasterMarkWordInfoRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 需要删除的文字水印Index。
	MarkWordIndex *int64 `json:"MarkWordIndex,omitnil,omitempty" name:"MarkWordIndex"`
}

type DeleteCasterMarkWordInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 需要删除的文字水印Index。
	MarkWordIndex *int64 `json:"MarkWordIndex,omitnil,omitempty" name:"MarkWordIndex"`
}

func (r *DeleteCasterMarkWordInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCasterMarkWordInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "MarkWordIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCasterMarkWordInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCasterMarkWordInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCasterMarkWordInfoResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCasterMarkWordInfoResponseParams `json:"Response"`
}

func (r *DeleteCasterMarkWordInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCasterMarkWordInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCasterOutputInfoRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 待删除的推流信息Index。
	// 注：删除时，该Index对应的配置需要存在。
	OutputIndex *uint64 `json:"OutputIndex,omitnil,omitempty" name:"OutputIndex"`
}

type DeleteCasterOutputInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 待删除的推流信息Index。
	// 注：删除时，该Index对应的配置需要存在。
	OutputIndex *uint64 `json:"OutputIndex,omitnil,omitempty" name:"OutputIndex"`
}

func (r *DeleteCasterOutputInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCasterOutputInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "OutputIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCasterOutputInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCasterOutputInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCasterOutputInfoResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCasterOutputInfoResponseParams `json:"Response"`
}

func (r *DeleteCasterOutputInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCasterOutputInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCasterRequestParams struct {
	// 待删除的导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type DeleteCasterRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *DeleteCasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCasterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCasterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCasterResponseParams `json:"Response"`
}

func (r *DeleteCasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLiveCallbackRuleRequestParams struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`
}

type DeleteLiveCallbackRuleRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`
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

// Predefined struct for user
type DeleteLiveCallbackRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveCallbackRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveCallbackRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveCallbackTemplateRequestParams struct {
	// 模板 ID。
	// 1. 在创建回调模板接口 [CreateLiveCallbackTemplate](/document/product/267/32637) 调用的返回值中获取模板 ID。
	// 2. 可以从接口 [DescribeLiveCallbackTemplates](/document/product/267/32632) 查询已经创建的过的模板列表。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID。
	// 1. 在创建回调模板接口 [CreateLiveCallbackTemplate](/document/product/267/32637) 调用的返回值中获取模板 ID。
	// 2. 可以从接口 [DescribeLiveCallbackTemplates](/document/product/267/32632) 查询已经创建的过的模板列表。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type DeleteLiveCallbackTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveCallbackTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveDomainRequestParams struct {
	// 要删除的域名
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 类型。0-推流，1-播放
	DomainType *uint64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`
}

type DeleteLiveDomainRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的域名
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 类型。0-推流，1-播放
	DomainType *uint64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`
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

// Predefined struct for user
type DeleteLiveDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveDomainResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLivePadRuleRequestParams struct {
	// 推流域名。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 直播垫片模板id。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteLivePadRuleRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 直播垫片模板id。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteLivePadRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLivePadRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "StreamName")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLivePadRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLivePadRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLivePadRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLivePadRuleResponseParams `json:"Response"`
}

func (r *DeleteLivePadRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLivePadRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLivePadTemplateRequestParams struct {
	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteLivePadTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteLivePadTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLivePadTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLivePadTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLivePadTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLivePadTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLivePadTemplateResponseParams `json:"Response"`
}

func (r *DeleteLivePadTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLivePadTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLivePullStreamTaskRequestParams struct {
	// 任务 Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作人姓名。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定任务 ID。注意：用于删除使用自定义任务 ID 创建的任务。
	SpecifyTaskId *string `json:"SpecifyTaskId,omitnil,omitempty" name:"SpecifyTaskId"`
}

type DeleteLivePullStreamTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务 Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作人姓名。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定任务 ID。注意：用于删除使用自定义任务 ID 创建的任务。
	SpecifyTaskId *string `json:"SpecifyTaskId,omitnil,omitempty" name:"SpecifyTaskId"`
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
	delete(f, "SpecifyTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLivePullStreamTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLivePullStreamTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLivePullStreamTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLivePullStreamTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveRecordRequestParams struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 任务ID，由CreateLiveRecord接口返回。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteLiveRecordRequest struct {
	*tchttp.BaseRequest
	
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 任务ID，由CreateLiveRecord接口返回。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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

// Predefined struct for user
type DeleteLiveRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveRecordResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveRecordRuleRequestParams struct {
	// 推流域名。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type DeleteLiveRecordRuleRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 域名+AppName+StreamName唯一标识单个转码规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
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

// Predefined struct for user
type DeleteLiveRecordRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveRecordRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveRecordRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveRecordTemplateRequestParams struct {
	// DescribeRecordTemplates接口获取到的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// DescribeRecordTemplates接口获取到的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type DeleteLiveRecordTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveRecordTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveSnapshotRuleRequestParams struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type DeleteLiveSnapshotRuleRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
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

// Predefined struct for user
type DeleteLiveSnapshotRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveSnapshotRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveSnapshotRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveSnapshotTemplateRequestParams struct {
	// 模板 ID。
	// 1. 在创建截图模板接口 [CreateLiveSnapshotTemplate](/document/product/267/32624) 调用的返回值中获取。
	// 2. 可以从接口 [DescribeLiveSnapshotTemplates](/document/product/267/32619) 中查询已创建的截图模板列表。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID。
	// 1. 在创建截图模板接口 [CreateLiveSnapshotTemplate](/document/product/267/32624) 调用的返回值中获取。
	// 2. 可以从接口 [DescribeLiveSnapshotTemplates](/document/product/267/32619) 中查询已创建的截图模板列表。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type DeleteLiveSnapshotTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveSnapshotTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveStreamMonitorRequestParams struct {
	// 监播任务ID
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

type DeleteLiveStreamMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 监播任务ID
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

func (r *DeleteLiveStreamMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveStreamMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveStreamMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLiveStreamMonitorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveStreamMonitorResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveStreamMonitorResponseParams `json:"Response"`
}

func (r *DeleteLiveStreamMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveStreamMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLiveTimeShiftRuleRequestParams struct {
	// 推流域名。
	// 域名+AppName+StreamName唯一标识单个时移规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	// 域名+AppName+StreamName唯一标识单个时移规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 域名+AppName+StreamName唯一标识单个时移规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type DeleteLiveTimeShiftRuleRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	// 域名+AppName+StreamName唯一标识单个时移规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	// 域名+AppName+StreamName唯一标识单个时移规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 域名+AppName+StreamName唯一标识单个时移规则，如需删除需要强匹配，例如AppName为空也需要传空字符串进行强匹配。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

func (r *DeleteLiveTimeShiftRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveTimeShiftRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "AppName")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveTimeShiftRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLiveTimeShiftRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveTimeShiftRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveTimeShiftRuleResponseParams `json:"Response"`
}

func (r *DeleteLiveTimeShiftRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveTimeShiftRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLiveTimeShiftTemplateRequestParams struct {
	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteLiveTimeShiftTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteLiveTimeShiftTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveTimeShiftTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveTimeShiftTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLiveTimeShiftTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveTimeShiftTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveTimeShiftTemplateResponseParams `json:"Response"`
}

func (r *DeleteLiveTimeShiftTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveTimeShiftTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLiveTranscodeRuleRequestParams struct {
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 模板ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteLiveTranscodeRuleRequest struct {
	*tchttp.BaseRequest
	
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 模板ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type DeleteLiveTranscodeRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveTranscodeRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveTranscodeRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveTranscodeTemplateRequestParams struct {
	// 模板 ID。
	// 1. 在创建转码模板接口 [CreateLiveTranscodeTemplate](/document/product/267/32646) 调用的返回值中获取模板 ID。
	// 2. 可以从接口 [DescribeLiveTranscodeTemplates](/document/product/267/32641) 查询已经创建过的模板列表。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID。
	// 1. 在创建转码模板接口 [CreateLiveTranscodeTemplate](/document/product/267/32646) 调用的返回值中获取模板 ID。
	// 2. 可以从接口 [DescribeLiveTranscodeTemplates](/document/product/267/32641) 查询已经创建过的模板列表。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type DeleteLiveTranscodeTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveTranscodeTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveWatermarkRequestParams struct {
	// 水印 ID。
	// 在添加水印接口 [AddLiveWatermark](/document/product/267/30154) 调用返回值中获取水印 ID。
	// 或DescribeLiveWatermarks接口返回的水印ID。
	WatermarkId *int64 `json:"WatermarkId,omitnil,omitempty" name:"WatermarkId"`
}

type DeleteLiveWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 水印 ID。
	// 在添加水印接口 [AddLiveWatermark](/document/product/267/30154) 调用返回值中获取水印 ID。
	// 或DescribeLiveWatermarks接口返回的水印ID。
	WatermarkId *int64 `json:"WatermarkId,omitnil,omitempty" name:"WatermarkId"`
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

// Predefined struct for user
type DeleteLiveWatermarkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveWatermarkResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveWatermarkRuleRequestParams struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。与推流和播放地址中的 AppName 保持一致，默认为live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type DeleteLiveWatermarkRuleRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。与推流和播放地址中的 AppName 保持一致，默认为live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
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

// Predefined struct for user
type DeleteLiveWatermarkRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLiveWatermarkRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveWatermarkRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type DeletePullStreamConfigRequestParams struct {
	// 配置 ID。
	// 1. 在添加拉流配置接口 [CreatePullStreamConfig](/document/api/267/30159) 调用返回值中获取配置 ID。
	// 2. 可以从接口 [DescribePullStreamConfigs](/document/api/267/30158) 中查询已创建过的拉流配置列表。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

type DeletePullStreamConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置 ID。
	// 1. 在添加拉流配置接口 [CreatePullStreamConfig](/document/api/267/30159) 调用返回值中获取配置 ID。
	// 2. 可以从接口 [DescribePullStreamConfigs](/document/api/267/30158) 中查询已创建过的拉流配置列表。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
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

// Predefined struct for user
type DeletePullStreamConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePullStreamConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeletePullStreamConfigResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRecordTaskRequestParams struct {
	// 任务ID，CreateRecordTask返回。删除TaskId指定的录制任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteRecordTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，CreateRecordTask返回。删除TaskId指定的录制任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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

// Predefined struct for user
type DeleteRecordTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRecordTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteScreenshotTaskRequestParams struct {
	// 任务ID，CreateScreenshotTask返回。删除TaskId指定的截图任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteScreenshotTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，CreateScreenshotTask返回。删除TaskId指定的截图任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteScreenshotTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScreenshotTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScreenshotTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScreenshotTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteScreenshotTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScreenshotTaskResponseParams `json:"Response"`
}

func (r *DeleteScreenshotTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScreenshotTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllStreamPlayInfoListRequestParams struct {
	// 查询时间点，精确到分钟粒度，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO%20%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F,-ISO%20%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 支持最近1个月的数据查询，数据延迟为5分钟左右，如果要查询实时的数据，建议传递5分钟前的整点时间点。
	QueryTime *string `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// 播放域名列表，若不填，表示总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`
}

type DescribeAllStreamPlayInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 查询时间点，精确到分钟粒度，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO%20%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F,-ISO%20%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 支持最近1个月的数据查询，数据延迟为5分钟左右，如果要查询实时的数据，建议传递5分钟前的整点时间点。
	QueryTime *string `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// 播放域名列表，若不填，表示总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`
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

// Predefined struct for user
type DescribeAllStreamPlayInfoListResponseParams struct {
	// 查询时间点，回传的输入参数中的查询时间。
	QueryTime *string `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// 数据信息列表。
	DataInfoList []*MonitorStreamPlayInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllStreamPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllStreamPlayInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAreaBillBandwidthAndFluxListRequestParams struct {
	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS，起始和结束时间跨度不支持超过1天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 直播播放域名，若不填，表示总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`
}

type DescribeAreaBillBandwidthAndFluxListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS，起始和结束时间跨度不支持超过1天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 直播播放域名，若不填，表示总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`
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

// Predefined struct for user
type DescribeAreaBillBandwidthAndFluxListResponseParams struct {
	// 明细数据信息。
	DataInfoList []*BillAreaInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAreaBillBandwidthAndFluxListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAreaBillBandwidthAndFluxListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBackupStreamListRequestParams struct {
	// 流名称，用于精确查询。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type DescribeBackupStreamListRequest struct {
	*tchttp.BaseRequest
	
	// 流名称，用于精确查询。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

func (r *DescribeBackupStreamListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupStreamListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupStreamListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupStreamListResponseParams struct {
	// 主备流分组信息列表。
	StreamInfoList []*BackupStreamGroupInfo `json:"StreamInfoList,omitnil,omitempty" name:"StreamInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupStreamListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupStreamListResponseParams `json:"Response"`
}

func (r *DescribeBackupStreamListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupStreamListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillBandwidthAndFluxListRequestParams struct {
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 起始和结束时间跨度不支持超过31天。支持最近3年的数据查询
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 直播播放域名，若不填，表示总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 可选值：
	// Mainland：查询国内数据，
	// Oversea：则查询国外数据，
	// 默认：查询国内+国外的数据。
	// 注：LEB（快直播）只支持国内+国外数据查询。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 数据粒度，支持如下粒度：
	// 5：5分钟粒度，（跨度不支持超过1天），
	// 60：1小时粒度（跨度不支持超过一个月），
	// 1440：天粒度（跨度不支持超过一个月）。
	// 默认值：5。
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 服务名称，可选值包括LVB(标准直播)，LEB(快直播)，不填则查LVB+LEB总值。
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

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
	RegionNames []*string `json:"RegionNames,omitnil,omitempty" name:"RegionNames"`
}

type DescribeBillBandwidthAndFluxListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 起始和结束时间跨度不支持超过31天。支持最近3年的数据查询
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 直播播放域名，若不填，表示总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 可选值：
	// Mainland：查询国内数据，
	// Oversea：则查询国外数据，
	// 默认：查询国内+国外的数据。
	// 注：LEB（快直播）只支持国内+国外数据查询。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 数据粒度，支持如下粒度：
	// 5：5分钟粒度，（跨度不支持超过1天），
	// 60：1小时粒度（跨度不支持超过一个月），
	// 1440：天粒度（跨度不支持超过一个月）。
	// 默认值：5。
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 服务名称，可选值包括LVB(标准直播)，LEB(快直播)，不填则查LVB+LEB总值。
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

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
	RegionNames []*string `json:"RegionNames,omitnil,omitempty" name:"RegionNames"`
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

// Predefined struct for user
type DescribeBillBandwidthAndFluxListResponseParams struct {
	// 峰值带宽所在时间点，接口返回支持两种时间格式(与接口请求传递的时间格式一致)：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	PeakBandwidthTime *string `json:"PeakBandwidthTime,omitnil,omitempty" name:"PeakBandwidthTime"`

	// 峰值带宽，单位是Mbps。
	PeakBandwidth *float64 `json:"PeakBandwidth,omitnil,omitempty" name:"PeakBandwidth"`

	// 95峰值带宽所在时间点，接口返回支持两种时间格式(与接口请求传递的时间格式一致)：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	P95PeakBandwidthTime *string `json:"P95PeakBandwidthTime,omitnil,omitempty" name:"P95PeakBandwidthTime"`

	// 95峰值带宽，单位是Mbps。
	P95PeakBandwidth *float64 `json:"P95PeakBandwidth,omitnil,omitempty" name:"P95PeakBandwidth"`

	// 总流量，单位是MB。
	SumFlux *float64 `json:"SumFlux,omitnil,omitempty" name:"SumFlux"`

	// 明细数据信息。
	DataInfoList []*BillDataInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillBandwidthAndFluxListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillBandwidthAndFluxListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCallbackRecordsListRequestParams struct {
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 
	// 查询的起始和结束时间跨度不支持超过1天。仅支持查询最近14天的数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流名称，精确匹配。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 页码。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页条数。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 事件类型。
	// 0: "断流",
	// 1: "推流",
	// 100: "录制"
	// 200: "截图回调"。
	EventType *uint64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 回调结果。
	// 0为成功，其他为失败。
	ResultCode *uint64 `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`
}

type DescribeCallbackRecordsListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 
	// 查询的起始和结束时间跨度不支持超过1天。仅支持查询最近14天的数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流名称，精确匹配。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 页码。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页条数。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 事件类型。
	// 0: "断流",
	// 1: "推流",
	// 100: "录制"
	// 200: "截图回调"。
	EventType *uint64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 回调结果。
	// 0为成功，其他为失败。
	ResultCode *uint64 `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`
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

// Predefined struct for user
type DescribeCallbackRecordsListResponseParams struct {
	// 回调事件列表。
	DataInfoList []*CallbackEventInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 页码。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页条数。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总条数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 总页数。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCallbackRecordsListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallbackRecordsListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCasterDisplayInfoRequestParams struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type DescribeCasterDisplayInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *DescribeCasterDisplayInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterDisplayInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCasterDisplayInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterDisplayInfoResponseParams struct {
	// 导播台状态
	// 0：停止状态，无预监，无主监 
	// 1：无预监，有主监 
	// 2：有预监，无主监 
	// 3：有预监，有主监
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 预监使用的展示参数。
	PvwDisplayInfo *CasterDisplayInfo `json:"PvwDisplayInfo,omitnil,omitempty" name:"PvwDisplayInfo"`

	// 主监使用的展示参数。
	PgmDisplayInfo *CasterDisplayInfo `json:"PgmDisplayInfo,omitnil,omitempty" name:"PgmDisplayInfo"`

	// 启动直播的时间，值为unix时间戳。
	StartLiveTime *uint64 `json:"StartLiveTime,omitnil,omitempty" name:"StartLiveTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCasterDisplayInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCasterDisplayInfoResponseParams `json:"Response"`
}

func (r *DescribeCasterDisplayInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterDisplayInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterInputInfosRequestParams struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type DescribeCasterInputInfosRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *DescribeCasterInputInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterInputInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCasterInputInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterInputInfosResponseParams struct {
	// 导播台输入源信息列表。
	InputInfos []*CasterInputInfo `json:"InputInfos,omitnil,omitempty" name:"InputInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCasterInputInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCasterInputInfosResponseParams `json:"Response"`
}

func (r *DescribeCasterInputInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterInputInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterLayoutInfosRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type DescribeCasterLayoutInfosRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *DescribeCasterLayoutInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterLayoutInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCasterLayoutInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterLayoutInfosResponseParams struct {
	// 导播台的布局列表。
	LayoutInfos []*CasterLayoutInfo `json:"LayoutInfos,omitnil,omitempty" name:"LayoutInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCasterLayoutInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCasterLayoutInfosResponseParams `json:"Response"`
}

func (r *DescribeCasterLayoutInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterLayoutInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterListRequestParams struct {

}

type DescribeCasterListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCasterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCasterListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterListResponseParams struct {
	// 用户对应的导播台简要信息列表
	CasterList []*CasterBriefInfo `json:"CasterList,omitnil,omitempty" name:"CasterList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCasterListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCasterListResponseParams `json:"Response"`
}

func (r *DescribeCasterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterMarkPicInfosRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type DescribeCasterMarkPicInfosRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *DescribeCasterMarkPicInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterMarkPicInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCasterMarkPicInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterMarkPicInfosResponseParams struct {
	// 导播台的水印信息列表。
	MarkPicInfos []*CasterMarkPicInfo `json:"MarkPicInfos,omitnil,omitempty" name:"MarkPicInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCasterMarkPicInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCasterMarkPicInfosResponseParams `json:"Response"`
}

func (r *DescribeCasterMarkPicInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterMarkPicInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterMarkWordInfosRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type DescribeCasterMarkWordInfosRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *DescribeCasterMarkWordInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterMarkWordInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCasterMarkWordInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterMarkWordInfosResponseParams struct {
	// 导播台的文本信息列表。
	MarkWordInfos []*CasterMarkWordInfo `json:"MarkWordInfos,omitnil,omitempty" name:"MarkWordInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCasterMarkWordInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCasterMarkWordInfosResponseParams `json:"Response"`
}

func (r *DescribeCasterMarkWordInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterMarkWordInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterOutputInfosRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type DescribeCasterOutputInfosRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *DescribeCasterOutputInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterOutputInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCasterOutputInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterOutputInfosResponseParams struct {
	// 导播台的推流信息列表。
	OutputInfos []*CasterOutputInfo `json:"OutputInfos,omitnil,omitempty" name:"OutputInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCasterOutputInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCasterOutputInfosResponseParams `json:"Response"`
}

func (r *DescribeCasterOutputInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterOutputInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterPlayUrlRequestParams struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 请求播放url的类型。
	// 取值范围[1，2，3]
	// 1：获取输入源的播放url
	// 2：获取pvw的播放url
	// 3：获取pgm的播放url
	PlayUrlType *int64 `json:"PlayUrlType,omitnil,omitempty" name:"PlayUrlType"`

	// 仅在PlayUrlType为1时生效，此时该参数表示请求的输入源Index。
	// 注：对应的输入源必须存在。
	PlayUrlIndex *int64 `json:"PlayUrlIndex,omitnil,omitempty" name:"PlayUrlIndex"`
}

type DescribeCasterPlayUrlRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 请求播放url的类型。
	// 取值范围[1，2，3]
	// 1：获取输入源的播放url
	// 2：获取pvw的播放url
	// 3：获取pgm的播放url
	PlayUrlType *int64 `json:"PlayUrlType,omitnil,omitempty" name:"PlayUrlType"`

	// 仅在PlayUrlType为1时生效，此时该参数表示请求的输入源Index。
	// 注：对应的输入源必须存在。
	PlayUrlIndex *int64 `json:"PlayUrlIndex,omitnil,omitempty" name:"PlayUrlIndex"`
}

func (r *DescribeCasterPlayUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterPlayUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "PlayUrlType")
	delete(f, "PlayUrlIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCasterPlayUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterPlayUrlResponseParams struct {
	// 播放url。
	// 当导播台不存在预监或主监时，若请求预监或主监的播放地址，该字段为空。
	PlayUrl *string `json:"PlayUrl,omitnil,omitempty" name:"PlayUrl"`

	// webrtc协议播放地址。
	// 当导播台不存在预监或主监时，若请求预监或主监的webrtc播放地址，该字段为空。
	// 注：webrtc协议播放地址需配合腾讯云快直播播放sdk方可使用。
	WebRTCPlayUrl *string `json:"WebRTCPlayUrl,omitnil,omitempty" name:"WebRTCPlayUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCasterPlayUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCasterPlayUrlResponseParams `json:"Response"`
}

func (r *DescribeCasterPlayUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterPlayUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterRequestParams struct {
	// 需查询的导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type DescribeCasterRequest struct {
	*tchttp.BaseRequest
	
	// 需查询的导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *DescribeCasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterResponseParams struct {
	// 导播台信息
	CasterInfo *CasterInfo `json:"CasterInfo,omitnil,omitempty" name:"CasterInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCasterResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCasterResponseParams `json:"Response"`
}

func (r *DescribeCasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterTransitionTypesRequestParams struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type DescribeCasterTransitionTypesRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *DescribeCasterTransitionTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterTransitionTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCasterTransitionTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterTransitionTypesResponseParams struct {
	// 转场信息列表
	TransitionTypes []*TransitionTypeInfo `json:"TransitionTypes,omitnil,omitempty" name:"TransitionTypes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCasterTransitionTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCasterTransitionTypesResponseParams `json:"Response"`
}

func (r *DescribeCasterTransitionTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterTransitionTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterUserStatusRequestParams struct {

}

type DescribeCasterUserStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCasterUserStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterUserStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCasterUserStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCasterUserStatusResponseParams struct {
	// 0: 未开通导播台
	// 1:开通了导播台，且处于正常状态
	// 2:开通了导播台，但处于欠费状态
	// 3:开通了导播台，但处于封禁状态
	UserStatus *int64 `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCasterUserStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCasterUserStatusResponseParams `json:"Response"`
}

func (r *DescribeCasterUserStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCasterUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrentRecordStreamNumRequestParams struct {
	// 直播类型，SlowLive：慢直播。
	// NormalLive：普通直播。
	LiveType *string `json:"LiveType,omitnil,omitempty" name:"LiveType"`

	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 可以查询最近180天的数据。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 时间跨度最大支持31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 如果为空，查询所有地区数据；如果为“Mainland”，查询国内数据；如果为“Oversea”，则查询国外数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 推流域名列表，不填表示总体数据。
	PushDomains []*string `json:"PushDomains,omitnil,omitempty" name:"PushDomains"`
}

type DescribeConcurrentRecordStreamNumRequest struct {
	*tchttp.BaseRequest
	
	// 直播类型，SlowLive：慢直播。
	// NormalLive：普通直播。
	LiveType *string `json:"LiveType,omitnil,omitempty" name:"LiveType"`

	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 可以查询最近180天的数据。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 时间跨度最大支持31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 如果为空，查询所有地区数据；如果为“Mainland”，查询国内数据；如果为“Oversea”，则查询国外数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 推流域名列表，不填表示总体数据。
	PushDomains []*string `json:"PushDomains,omitnil,omitempty" name:"PushDomains"`
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

// Predefined struct for user
type DescribeConcurrentRecordStreamNumResponseParams struct {
	// 统计信息列表。
	DataInfoList []*ConcurrentRecordStreamNum `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrentRecordStreamNumResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrentRecordStreamNumResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDeliverBandwidthListRequestParams struct {
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 支持最近三个月的数据查询，时间跨度最大是1个月。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeDeliverBandwidthListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 支持最近三个月的数据查询，时间跨度最大是1个月。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
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

// Predefined struct for user
type DescribeDeliverBandwidthListResponseParams struct {
	// 转推计费带宽数据
	DataInfoList []*BandwidthInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeliverBandwidthListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeliverBandwidthListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDeliverLogDownListRequestParams struct {

}

type DescribeDeliverLogDownListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDeliverLogDownListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeliverLogDownListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeliverLogDownListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeliverLogDownListResponseParams struct {
	// 日志信息列表。
	LogInfoList []*PushLogInfo `json:"LogInfoList,omitnil,omitempty" name:"LogInfoList"`

	// 总条数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeliverLogDownListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeliverLogDownListResponseParams `json:"Response"`
}

func (r *DescribeDeliverLogDownListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeliverLogDownListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupProIspPlayInfoListRequestParams struct {
	// 起始时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 时间跨度在（0,3小时]，支持最近1个月数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 播放域名，默认为不填，表示求总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 省份列表，默认不填，则返回各省份的数据。
	ProvinceNames []*string `json:"ProvinceNames,omitnil,omitempty" name:"ProvinceNames"`

	// 运营商列表，默认不填，则返回整个运营商的数据。
	IspNames []*string `json:"IspNames,omitnil,omitempty" name:"IspNames"`

	// 国内还是国外，如果为空，查询所有地区数据；如果为“Mainland”，查询国内数据；如果为“Oversea”，则查询国外数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`
}

type DescribeGroupProIspPlayInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 时间跨度在（0,3小时]，支持最近1个月数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 播放域名，默认为不填，表示求总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 省份列表，默认不填，则返回各省份的数据。
	ProvinceNames []*string `json:"ProvinceNames,omitnil,omitempty" name:"ProvinceNames"`

	// 运营商列表，默认不填，则返回整个运营商的数据。
	IspNames []*string `json:"IspNames,omitnil,omitempty" name:"IspNames"`

	// 国内还是国外，如果为空，查询所有地区数据；如果为“Mainland”，查询国内数据；如果为“Oversea”，则查询国外数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`
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

// Predefined struct for user
type DescribeGroupProIspPlayInfoListResponseParams struct {
	// 数据内容。
	DataInfoList []*GroupProIspDataInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGroupProIspPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupProIspPlayInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeHttpStatusInfoListRequestParams struct {
	// 起始时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 播放域名列表。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`
}

type DescribeHttpStatusInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 播放域名列表。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`
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

// Predefined struct for user
type DescribeHttpStatusInfoListResponseParams struct {
	// 播放状态码列表。
	DataInfoList []*HttpStatusData `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHttpStatusInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHttpStatusInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveCallbackRulesRequestParams struct {

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

// Predefined struct for user
type DescribeLiveCallbackRulesResponseParams struct {
	// 规则信息列表。
	Rules []*CallBackRuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveCallbackRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveCallbackRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveCallbackTemplateRequestParams struct {
	// 模板 ID。
	// 1. 在创建回调模板接口 [CreateLiveCallbackTemplate](/document/product/267/32637) 调用的返回值中获取模板 ID。
	// 2. 可以从接口 [DescribeLiveCallbackTemplates](/document/product/267/32632) 查询已经创建的过的模板列表。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID。
	// 1. 在创建回调模板接口 [CreateLiveCallbackTemplate](/document/product/267/32637) 调用的返回值中获取模板 ID。
	// 2. 可以从接口 [DescribeLiveCallbackTemplates](/document/product/267/32632) 查询已经创建的过的模板列表。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type DescribeLiveCallbackTemplateResponseParams struct {
	// 回调模板信息。
	Template *CallBackTemplateInfo `json:"Template,omitnil,omitempty" name:"Template"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveCallbackTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveCallbackTemplatesRequestParams struct {

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

// Predefined struct for user
type DescribeLiveCallbackTemplatesResponseParams struct {
	// 模板信息列表。
	Templates []*CallBackTemplateInfo `json:"Templates,omitnil,omitempty" name:"Templates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveCallbackTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveCallbackTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveCertRequestParams struct {
	// DescribeLiveCerts接口获取到的证书Id。
	CertId *int64 `json:"CertId,omitnil,omitempty" name:"CertId"`
}

type DescribeLiveCertRequest struct {
	*tchttp.BaseRequest
	
	// DescribeLiveCerts接口获取到的证书Id。
	CertId *int64 `json:"CertId,omitnil,omitempty" name:"CertId"`
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

// Predefined struct for user
type DescribeLiveCertResponseParams struct {
	// 证书信息。
	CertInfo *CertInfo `json:"CertInfo,omitnil,omitempty" name:"CertInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveCertResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveCertResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveCertsRequestParams struct {

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

// Predefined struct for user
type DescribeLiveCertsResponseParams struct {
	// 证书信息列表。
	CertInfoSet []*CertInfo `json:"CertInfoSet,omitnil,omitempty" name:"CertInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveCertsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveCertsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveDelayInfoListRequestParams struct {

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

// Predefined struct for user
type DescribeLiveDelayInfoListResponseParams struct {
	// 延播信息列表。
	DelayInfoList []*DelayInfo `json:"DelayInfoList,omitnil,omitempty" name:"DelayInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveDelayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveDelayInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveDomainCertBindingsRequestParams struct {
	// 要搜索的域名字符串。
	DomainSearch *string `json:"DomainSearch,omitnil,omitempty" name:"DomainSearch"`

	// 记录行的位置，从0开始。默认0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 记录行的最大数目。默认50。
	// 若不传，则最多返回50条数据。
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 要查询的单个域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 可取值：
	// ExpireTimeAsc：证书过期时间升序。
	// ExpireTimeDesc：证书过期时间降序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

type DescribeLiveDomainCertBindingsRequest struct {
	*tchttp.BaseRequest
	
	// 要搜索的域名字符串。
	DomainSearch *string `json:"DomainSearch,omitnil,omitempty" name:"DomainSearch"`

	// 记录行的位置，从0开始。默认0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 记录行的最大数目。默认50。
	// 若不传，则最多返回50条数据。
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 要查询的单个域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 可取值：
	// ExpireTimeAsc：证书过期时间升序。
	// ExpireTimeDesc：证书过期时间降序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

func (r *DescribeLiveDomainCertBindingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDomainCertBindingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainSearch")
	delete(f, "Offset")
	delete(f, "Length")
	delete(f, "DomainName")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveDomainCertBindingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveDomainCertBindingsResponseParams struct {
	// 有绑定证书的域名信息数组。
	LiveDomainCertBindings []*LiveDomainCertBindings `json:"LiveDomainCertBindings,omitnil,omitempty" name:"LiveDomainCertBindings"`

	// 总的记录行数，便于分页。
	TotalNum *int64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveDomainCertBindingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveDomainCertBindingsResponseParams `json:"Response"`
}

func (r *DescribeLiveDomainCertBindingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveDomainCertBindingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveDomainCertRequestParams struct {
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
}

type DescribeLiveDomainCertRequest struct {
	*tchttp.BaseRequest
	
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
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

// Predefined struct for user
type DescribeLiveDomainCertResponseParams struct {
	// 证书信息。
	DomainCertInfo *DomainCertInfo `json:"DomainCertInfo,omitnil,omitempty" name:"DomainCertInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveDomainCertResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveDomainPlayInfoListRequestParams struct {
	// 播放域名列表。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`
}

type DescribeLiveDomainPlayInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 播放域名列表。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`
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

// Predefined struct for user
type DescribeLiveDomainPlayInfoListResponseParams struct {
	// 数据时间，格式为yyyy-mm-dd HH:MM:SS。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 实时总带宽。
	TotalBandwidth *float64 `json:"TotalBandwidth,omitnil,omitempty" name:"TotalBandwidth"`

	// 实时总流量。
	TotalFlux *float64 `json:"TotalFlux,omitnil,omitempty" name:"TotalFlux"`

	// 总请求数。
	TotalRequest *uint64 `json:"TotalRequest,omitnil,omitempty" name:"TotalRequest"`

	// 实时总连接数。
	TotalOnline *uint64 `json:"TotalOnline,omitnil,omitempty" name:"TotalOnline"`

	// 分域名的数据情况。
	DomainInfoList []*DomainInfoList `json:"DomainInfoList,omitnil,omitempty" name:"DomainInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveDomainPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveDomainPlayInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveDomainRefererRequestParams struct {
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
}

type DescribeLiveDomainRefererRequest struct {
	*tchttp.BaseRequest
	
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
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

// Predefined struct for user
type DescribeLiveDomainRefererResponseParams struct {
	// 域名 Referer 黑白名单配置。
	RefererAuthConfig *RefererAuthConfig `json:"RefererAuthConfig,omitnil,omitempty" name:"RefererAuthConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveDomainRefererResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveDomainRefererResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveDomainRequestParams struct {
	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
}

type DescribeLiveDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
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

// Predefined struct for user
type DescribeLiveDomainResponseParams struct {
	// 域名信息。
	DomainInfo *DomainInfo `json:"DomainInfo,omitnil,omitempty" name:"DomainInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveDomainResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveDomainsRequestParams struct {
	// 域名状态过滤。0-停用，1-启用。
	DomainStatus *uint64 `json:"DomainStatus,omitnil,omitempty" name:"DomainStatus"`

	// 域名类型过滤。0-推流，1-播放。
	DomainType *uint64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`

	// 分页大小，范围：10~100。默认10。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 取第几页，范围：1~100000。默认1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 0 普通直播 1慢直播 默认0。
	IsDelayLive *uint64 `json:"IsDelayLive,omitnil,omitempty" name:"IsDelayLive"`

	// 域名前缀。
	DomainPrefix *string `json:"DomainPrefix,omitnil,omitempty" name:"DomainPrefix"`

	// 播放区域，只在 DomainType=1 时该参数有意义。
	// 1: 国内。
	// 2: 全球。
	// 3: 海外。
	PlayType *uint64 `json:"PlayType,omitnil,omitempty" name:"PlayType"`
}

type DescribeLiveDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 域名状态过滤。0-停用，1-启用。
	DomainStatus *uint64 `json:"DomainStatus,omitnil,omitempty" name:"DomainStatus"`

	// 域名类型过滤。0-推流，1-播放。
	DomainType *uint64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`

	// 分页大小，范围：10~100。默认10。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 取第几页，范围：1~100000。默认1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 0 普通直播 1慢直播 默认0。
	IsDelayLive *uint64 `json:"IsDelayLive,omitnil,omitempty" name:"IsDelayLive"`

	// 域名前缀。
	DomainPrefix *string `json:"DomainPrefix,omitnil,omitempty" name:"DomainPrefix"`

	// 播放区域，只在 DomainType=1 时该参数有意义。
	// 1: 国内。
	// 2: 全球。
	// 3: 海外。
	PlayType *uint64 `json:"PlayType,omitnil,omitempty" name:"PlayType"`
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

// Predefined struct for user
type DescribeLiveDomainsResponseParams struct {
	// 总记录数。
	AllCount *uint64 `json:"AllCount,omitnil,omitempty" name:"AllCount"`

	// 域名详细信息列表。
	DomainList []*DomainInfo `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 可继续添加域名数量。
	CreateLimitCount *int64 `json:"CreateLimitCount,omitnil,omitempty" name:"CreateLimitCount"`

	// 启用的播放域名加速区域统计，数组元素分别为：中国大陆（境内），全球地区，国际/港澳台（境外）域名数量。
	PlayTypeCount []*int64 `json:"PlayTypeCount,omitnil,omitempty" name:"PlayTypeCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveDomainsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveEnhanceInfoListRequestParams struct {
	// 起始时间点，使用ISO格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见[ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 接口支持最近三个月的查询，开始时间和结束时间查询跨度不能超过三十天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，使用ISO格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见[ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 接口支持最近三个月的查询，开始时间和结束时间查询跨度不能超过三十天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询粒度，支持5，60分钟。
	Granularity *int64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 查询域名，如果不填则默认查全部的数据。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// 增强服务类型，如果不填则默认查全部的数据。
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分辨率，如果不填则默认查全部的数据。
	Resolution []*string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 帧率，如果不填则默认查全部的数据。
	Fps []*string `json:"Fps,omitnil,omitempty" name:"Fps"`
}

type DescribeLiveEnhanceInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，使用ISO格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见[ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 接口支持最近三个月的查询，开始时间和结束时间查询跨度不能超过三十天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，使用ISO格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见[ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 接口支持最近三个月的查询，开始时间和结束时间查询跨度不能超过三十天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询粒度，支持5，60分钟。
	Granularity *int64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 查询域名，如果不填则默认查全部的数据。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// 增强服务类型，如果不填则默认查全部的数据。
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分辨率，如果不填则默认查全部的数据。
	Resolution []*string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 帧率，如果不填则默认查全部的数据。
	Fps []*string `json:"Fps,omitnil,omitempty" name:"Fps"`
}

func (r *DescribeLiveEnhanceInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveEnhanceInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Granularity")
	delete(f, "DomainNames")
	delete(f, "Type")
	delete(f, "Resolution")
	delete(f, "Fps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveEnhanceInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveEnhanceInfoListResponseParams struct {
	// 直播增强统计信息列表。
	DataInfoList []*LiveEnhanceInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveEnhanceInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveEnhanceInfoListResponseParams `json:"Response"`
}

func (r *DescribeLiveEnhanceInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveEnhanceInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveForbidStreamListRequestParams struct {
	// 取得第几页，默认1。
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页大小，最大100。 
	// 取值：1~100之前的任意整数。
	// 默认值：10。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 按流名称查询。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type DescribeLiveForbidStreamListRequest struct {
	*tchttp.BaseRequest
	
	// 取得第几页，默认1。
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页大小，最大100。 
	// 取值：1~100之前的任意整数。
	// 默认值：10。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 按流名称查询。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
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

// Predefined struct for user
type DescribeLiveForbidStreamListResponseParams struct {
	// 符合条件的总个数。
	TotalNum *int64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 总页数。
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 分页的页码。
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页显示的条数。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 禁推流列表。
	ForbidStreamList []*ForbidStreamInfo `json:"ForbidStreamList,omitnil,omitempty" name:"ForbidStreamList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveForbidStreamListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveForbidStreamListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLivePackageInfoRequestParams struct {
	// 包类型，可选值：
	// 0：流量包
	// 1：转码包
	// 2: 连麦包。
	PackageType *int64 `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 排序规则:
	// 1. BuyTimeDesc： 最新购买的排在最前面
	// 2. BuyTimeAsc： 最老购买的排在最前面
	// 3. ExpireTimeDesc： 最后过期的排在最前面
	// 4. ExpireTimeAsc：最先过期的排在最前面。
	// 
	// 注意：
	// 1. PackageType 为 2（连麦包） 的时候，不支持 3、4 排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 取得第几页的数据，和 PageSize 同时传递才会生效。
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 分页大小，和 PageNum 同时传递才会生效。
	// 取值：10 ～ 100 之间的任意整数。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeLivePackageInfoRequest struct {
	*tchttp.BaseRequest
	
	// 包类型，可选值：
	// 0：流量包
	// 1：转码包
	// 2: 连麦包。
	PackageType *int64 `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 排序规则:
	// 1. BuyTimeDesc： 最新购买的排在最前面
	// 2. BuyTimeAsc： 最老购买的排在最前面
	// 3. ExpireTimeDesc： 最后过期的排在最前面
	// 4. ExpireTimeAsc：最先过期的排在最前面。
	// 
	// 注意：
	// 1. PackageType 为 2（连麦包） 的时候，不支持 3、4 排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 取得第几页的数据，和 PageSize 同时传递才会生效。
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 分页大小，和 PageNum 同时传递才会生效。
	// 取值：10 ～ 100 之间的任意整数。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
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

// Predefined struct for user
type DescribeLivePackageInfoResponseParams struct {
	// 套餐包信息。
	LivePackageInfoList []*LivePackageInfo `json:"LivePackageInfoList,omitnil,omitempty" name:"LivePackageInfoList"`

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
	PackageBillMode *int64 `json:"PackageBillMode,omitnil,omitempty" name:"PackageBillMode"`

	// 总页数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 数据总条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalNum *int64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 当前页数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 当前每页数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当请求参数 PackageType = 0 时生效，逗号分隔，从第一个到最后一个分别表示：
	// 标准直播，中国大陆（境内全地区）计费方式。
	// 标准直播，国际/港澳台（境外多地区）计费方式。
	// 快直播，中国大陆（境内全地区）计费方式。
	// 快直播，国际/港澳台（境外多地区）计费方式。
	FluxPackageBillMode *string `json:"FluxPackageBillMode,omitnil,omitempty" name:"FluxPackageBillMode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLivePackageInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLivePackageInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLivePadProcessorListRequestParams struct {
	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 您的推流域名。
	PushDomainName *string `json:"PushDomainName,omitnil,omitempty" name:"PushDomainName"`
}

type DescribeLivePadProcessorListRequest struct {
	*tchttp.BaseRequest
	
	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 您的推流域名。
	PushDomainName *string `json:"PushDomainName,omitnil,omitempty" name:"PushDomainName"`
}

func (r *DescribeLivePadProcessorListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePadProcessorListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "PushDomainName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLivePadProcessorListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLivePadProcessorListResponseParams struct {
	// 当前正在拉取垫片的流名称列表。
	StreamNameList []*string `json:"StreamNameList,omitnil,omitempty" name:"StreamNameList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLivePadProcessorListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLivePadProcessorListResponseParams `json:"Response"`
}

func (r *DescribeLivePadProcessorListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePadProcessorListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLivePadRulesRequestParams struct {

}

type DescribeLivePadRulesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLivePadRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePadRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLivePadRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLivePadRulesResponseParams struct {
	// 规则信息列表。
	Rules []*RuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLivePadRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLivePadRulesResponseParams `json:"Response"`
}

func (r *DescribeLivePadRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePadRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLivePadTemplateRequestParams struct {
	// 模板id。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeLivePadTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板id。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeLivePadTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePadTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLivePadTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLivePadTemplateResponseParams struct {
	// 直播垫片模板信息。
	Template *PadTemplate `json:"Template,omitnil,omitempty" name:"Template"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLivePadTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLivePadTemplateResponseParams `json:"Response"`
}

func (r *DescribeLivePadTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePadTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLivePadTemplatesRequestParams struct {

}

type DescribeLivePadTemplatesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLivePadTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePadTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLivePadTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLivePadTemplatesResponseParams struct {
	// 直播垫片模板信息。
	Templates []*PadTemplate `json:"Templates,omitnil,omitempty" name:"Templates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLivePadTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLivePadTemplatesResponseParams `json:"Response"`
}

func (r *DescribeLivePadTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePadTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLivePlayAuthKeyRequestParams struct {
	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
}

type DescribeLivePlayAuthKeyRequest struct {
	*tchttp.BaseRequest
	
	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
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

// Predefined struct for user
type DescribeLivePlayAuthKeyResponseParams struct {
	// 播放鉴权key信息。
	PlayAuthKeyInfo *PlayAuthKeyInfo `json:"PlayAuthKeyInfo,omitnil,omitempty" name:"PlayAuthKeyInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLivePlayAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLivePlayAuthKeyResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLivePullStreamTaskStatusRequestParams struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeLivePullStreamTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeLivePullStreamTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePullStreamTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLivePullStreamTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLivePullStreamTaskStatusResponseParams struct {
	// 任务状态信息。
	TaskStatusInfo *TaskStatusInfo `json:"TaskStatusInfo,omitnil,omitempty" name:"TaskStatusInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLivePullStreamTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLivePullStreamTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeLivePullStreamTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLivePullStreamTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLivePullStreamTasksRequestParams struct {
	// 任务 ID。 
	// 来源：调用 CreateLivePullStreamTask 接口时返回。
	// 不填默认查询所有任务，按更新时间倒序排序。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 取得第几页，默认值：1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 分页大小，默认值：10。
	// 取值范围：1~20 之前的任意整数。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 使用指定任务 ID 查询任务信息。
	// 注意：仅供使用指定 ID 创建的任务查询。
	SpecifyTaskId *string `json:"SpecifyTaskId,omitnil,omitempty" name:"SpecifyTaskId"`
}

type DescribeLivePullStreamTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID。 
	// 来源：调用 CreateLivePullStreamTask 接口时返回。
	// 不填默认查询所有任务，按更新时间倒序排序。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 取得第几页，默认值：1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 分页大小，默认值：10。
	// 取值范围：1~20 之前的任意整数。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 使用指定任务 ID 查询任务信息。
	// 注意：仅供使用指定 ID 创建的任务查询。
	SpecifyTaskId *string `json:"SpecifyTaskId,omitnil,omitempty" name:"SpecifyTaskId"`
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
	delete(f, "SpecifyTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLivePullStreamTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLivePullStreamTasksResponseParams struct {
	// 直播拉流任务信息列表。
	TaskInfos []*PullStreamTaskInfo `json:"TaskInfos,omitnil,omitempty" name:"TaskInfos"`

	// 分页的页码。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页大小。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 符合条件的总个数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 总页数。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 限制可创建的最大任务数。
	LimitTaskNum *uint64 `json:"LimitTaskNum,omitnil,omitempty" name:"LimitTaskNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLivePullStreamTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLivePullStreamTasksResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLivePushAuthKeyRequestParams struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
}

type DescribeLivePushAuthKeyRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
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

// Predefined struct for user
type DescribeLivePushAuthKeyResponseParams struct {
	// 推流鉴权key信息。
	PushAuthKeyInfo *PushAuthKeyInfo `json:"PushAuthKeyInfo,omitnil,omitempty" name:"PushAuthKeyInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLivePushAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLivePushAuthKeyResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveRecordRulesRequestParams struct {

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

// Predefined struct for user
type DescribeLiveRecordRulesResponseParams struct {
	// 规则列表。
	Rules []*RuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveRecordRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveRecordRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveRecordTemplateRequestParams struct {
	// [DescribeLiveRecordTemplates](/document/product/267/32609)接口获取到的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// [DescribeLiveRecordTemplates](/document/product/267/32609)接口获取到的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type DescribeLiveRecordTemplateResponseParams struct {
	// 录制模板信息。
	Template *RecordTemplateInfo `json:"Template,omitnil,omitempty" name:"Template"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveRecordTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveRecordTemplatesRequestParams struct {
	// 是否属于慢直播模板，默认：0。
	// 0： 标准直播。
	// 1：慢直播。
	IsDelayLive *int64 `json:"IsDelayLive,omitnil,omitempty" name:"IsDelayLive"`
}

type DescribeLiveRecordTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 是否属于慢直播模板，默认：0。
	// 0： 标准直播。
	// 1：慢直播。
	IsDelayLive *int64 `json:"IsDelayLive,omitnil,omitempty" name:"IsDelayLive"`
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

// Predefined struct for user
type DescribeLiveRecordTemplatesResponseParams struct {
	// 录制模板信息列表。
	Templates []*RecordTemplateInfo `json:"Templates,omitnil,omitempty" name:"Templates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveRecordTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveRecordTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveSnapshotRulesRequestParams struct {

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

// Predefined struct for user
type DescribeLiveSnapshotRulesResponseParams struct {
	// 规则列表。
	Rules []*RuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveSnapshotRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveSnapshotRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveSnapshotTemplateRequestParams struct {
	// 模板 ID。
	// 调用 [CreateLiveSnapshotTemplate](/document/product/267/32624) 时返回的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID。
	// 调用 [CreateLiveSnapshotTemplate](/document/product/267/32624) 时返回的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type DescribeLiveSnapshotTemplateResponseParams struct {
	// 截图模板信息。
	Template *SnapshotTemplateInfo `json:"Template,omitnil,omitempty" name:"Template"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveSnapshotTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveSnapshotTemplatesRequestParams struct {

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

// Predefined struct for user
type DescribeLiveSnapshotTemplatesResponseParams struct {
	// 截图模板列表。
	Templates []*SnapshotTemplateInfo `json:"Templates,omitnil,omitempty" name:"Templates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveSnapshotTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveSnapshotTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveStreamEventListRequestParams struct {
	// 起始时间。 UTC 格式，例如：2018-12-29T19:00:00Z。支持查询2个月内的历史记录。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。UTC 格式，例如：2018-12-29T20:00:00Z。不超过当前时间，且和起始时间相差不得超过1个月。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称，不支持通配符（*）查询，默认模糊匹配。
	// 可使用IsStrict字段改为精确查询。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 取得第几页。
	// 默认值：1。
	// 注： 目前只支持10000条内的查询。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 分页大小。
	// 最大值：100。
	// 取值范围：1~100 之间的任意整数。
	// 默认值：10。
	// 注： 目前只支持10000条内的查询。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 是否过滤，默认不过滤。
	// 0：不进行任何过滤。
	// 1：过滤掉开播失败的，只返回开播成功的。
	IsFilter *int64 `json:"IsFilter,omitnil,omitempty" name:"IsFilter"`

	// 是否精确查询，默认模糊匹配。
	// 0：模糊匹配。
	// 1：精确查询。
	// 注：使用StreamName时该参数生效。
	IsStrict *int64 `json:"IsStrict,omitnil,omitempty" name:"IsStrict"`

	// 是否按结束时间正序显示，默认逆序。
	// 0：逆序。
	// 1：正序。
	IsAsc *int64 `json:"IsAsc,omitnil,omitempty" name:"IsAsc"`
}

type DescribeLiveStreamEventListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间。 UTC 格式，例如：2018-12-29T19:00:00Z。支持查询2个月内的历史记录。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。UTC 格式，例如：2018-12-29T20:00:00Z。不超过当前时间，且和起始时间相差不得超过1个月。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称，不支持通配符（*）查询，默认模糊匹配。
	// 可使用IsStrict字段改为精确查询。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 取得第几页。
	// 默认值：1。
	// 注： 目前只支持10000条内的查询。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 分页大小。
	// 最大值：100。
	// 取值范围：1~100 之间的任意整数。
	// 默认值：10。
	// 注： 目前只支持10000条内的查询。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 是否过滤，默认不过滤。
	// 0：不进行任何过滤。
	// 1：过滤掉开播失败的，只返回开播成功的。
	IsFilter *int64 `json:"IsFilter,omitnil,omitempty" name:"IsFilter"`

	// 是否精确查询，默认模糊匹配。
	// 0：模糊匹配。
	// 1：精确查询。
	// 注：使用StreamName时该参数生效。
	IsStrict *int64 `json:"IsStrict,omitnil,omitempty" name:"IsStrict"`

	// 是否按结束时间正序显示，默认逆序。
	// 0：逆序。
	// 1：正序。
	IsAsc *int64 `json:"IsAsc,omitnil,omitempty" name:"IsAsc"`
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

// Predefined struct for user
type DescribeLiveStreamEventListResponseParams struct {
	// 推断流事件列表。
	EventList []*StreamEventInfo `json:"EventList,omitnil,omitempty" name:"EventList"`

	// 分页的页码。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页大小。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 符合条件的总个数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 总页数。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveStreamEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveStreamEventListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveStreamMonitorListRequestParams struct {
	// 查询列表时的起始偏移。
	Index *uint64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 本次查询的记录个数。最小值为1。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type DescribeLiveStreamMonitorListRequest struct {
	*tchttp.BaseRequest
	
	// 查询列表时的起始偏移。
	Index *uint64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 本次查询的记录个数。最小值为1。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

func (r *DescribeLiveStreamMonitorListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamMonitorListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Index")
	delete(f, "Count")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveStreamMonitorListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveStreamMonitorListResponseParams struct {
	// 账号下的直播流监播任务个数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 直播流监播任务列表
	LiveStreamMonitors []*LiveStreamMonitorInfo `json:"LiveStreamMonitors,omitnil,omitempty" name:"LiveStreamMonitors"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveStreamMonitorListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveStreamMonitorListResponseParams `json:"Response"`
}

func (r *DescribeLiveStreamMonitorListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamMonitorListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveStreamMonitorRequestParams struct {
	// 监播任务ID。
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

type DescribeLiveStreamMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 监播任务ID。
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

func (r *DescribeLiveStreamMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveStreamMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveStreamMonitorResponseParams struct {
	// 直播监播任务相关信息。
	LiveStreamMonitor *LiveStreamMonitorInfo `json:"LiveStreamMonitor,omitnil,omitempty" name:"LiveStreamMonitor"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveStreamMonitorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveStreamMonitorResponseParams `json:"Response"`
}

func (r *DescribeLiveStreamMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveStreamOnlineListRequestParams struct {
	// 推流域名。多域名用户需要填写 DomainName。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。多路径用户需要填写 AppName。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 取得第几页，默认1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页大小，最大100。 
	// 取值：10~100之间的任意整数。
	// 默认值：10。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 流名称，用于精确查询。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type DescribeLiveStreamOnlineListRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。多域名用户需要填写 DomainName。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。多路径用户需要填写 AppName。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 取得第几页，默认1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页大小，最大100。 
	// 取值：10~100之间的任意整数。
	// 默认值：10。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 流名称，用于精确查询。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
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

// Predefined struct for user
type DescribeLiveStreamOnlineListResponseParams struct {
	// 符合条件的总个数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 总页数。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 分页的页码。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页显示的条数。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 正在推送流的信息列表。
	OnlineInfo []*StreamOnlineInfo `json:"OnlineInfo,omitnil,omitempty" name:"OnlineInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveStreamOnlineListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveStreamOnlineListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveStreamPublishedListRequestParams struct {
	// 您的推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 结束时间。UTC 格式，例如：2016-06-30T19:00:00Z。不超过当前时间。注意：EndTime和StartTime相差不可超过1个月。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 起始时间。 UTC 格式，例如：2016-06-29T19:00:00Z。最长支持查询2个月内数据。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。不支持模糊匹配。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 取得第几页。
	// 默认值：1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 分页大小。
	// 最大值：100。
	// 取值范围：10~100 之前的任意整数。
	// 默认值：10。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 流名称，支持模糊匹配。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type DescribeLiveStreamPublishedListRequest struct {
	*tchttp.BaseRequest
	
	// 您的推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 结束时间。UTC 格式，例如：2016-06-30T19:00:00Z。不超过当前时间。注意：EndTime和StartTime相差不可超过1个月。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 起始时间。 UTC 格式，例如：2016-06-29T19:00:00Z。最长支持查询2个月内数据。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 推流路径，与推流和播放地址中的 AppName 保持一致，默认为 live。不支持模糊匹配。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 取得第几页。
	// 默认值：1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 分页大小。
	// 最大值：100。
	// 取值范围：10~100 之前的任意整数。
	// 默认值：10。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 流名称，支持模糊匹配。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
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

// Predefined struct for user
type DescribeLiveStreamPublishedListResponseParams struct {
	// 推流记录信息。
	PublishInfo []*StreamName `json:"PublishInfo,omitnil,omitempty" name:"PublishInfo"`

	// 分页的页码。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 符合条件的总个数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 总页数。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveStreamPublishedListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveStreamPublishedListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveStreamPushInfoListRequestParams struct {
	// 推流域名。
	PushDomain *string `json:"PushDomain,omitnil,omitempty" name:"PushDomain"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 页数，
	// 范围[1,10000]，
	// 默认值：1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，
	// 范围：[1,1000]，
	// 默认值： 200。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeLiveStreamPushInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	PushDomain *string `json:"PushDomain,omitnil,omitempty" name:"PushDomain"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 页数，
	// 范围[1,10000]，
	// 默认值：1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，
	// 范围：[1,1000]，
	// 默认值： 200。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
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

// Predefined struct for user
type DescribeLiveStreamPushInfoListResponseParams struct {
	// 直播流的统计信息列表。
	DataInfoList []*PushDataInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 所有在线流的总数量。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 总页数。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 当前数据所在页码。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页的在线流的个数。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveStreamPushInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveStreamPushInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveStreamStateRequestParams struct {
	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type DescribeLiveStreamStateRequest struct {
	*tchttp.BaseRequest
	
	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
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

// Predefined struct for user
type DescribeLiveStreamStateResponseParams struct {
	// 流状态，
	// active：活跃，
	// inactive：非活跃，
	// forbid：禁播。
	StreamState *string `json:"StreamState,omitnil,omitempty" name:"StreamState"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveStreamStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveStreamStateResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveTimeShiftBillInfoListRequestParams struct {
	// UTC开始时间，支持最近三个月的查询，查询时间最长跨度为一天。使用 UTC 格式时间，例如：2019-01-08T10:00:00Z。注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// UTC结束时间，支持最近三个月的查询，查询时间最长跨度为一天。使用 UTC 格式时间，例如：2019-01-08T10:00:00Z。注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 推流域名列表，若不传递此参数，则表示查询总体数据。
	PushDomains []*string `json:"PushDomains,omitnil,omitempty" name:"PushDomains"`
}

type DescribeLiveTimeShiftBillInfoListRequest struct {
	*tchttp.BaseRequest
	
	// UTC开始时间，支持最近三个月的查询，查询时间最长跨度为一天。使用 UTC 格式时间，例如：2019-01-08T10:00:00Z。注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// UTC结束时间，支持最近三个月的查询，查询时间最长跨度为一天。使用 UTC 格式时间，例如：2019-01-08T10:00:00Z。注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 推流域名列表，若不传递此参数，则表示查询总体数据。
	PushDomains []*string `json:"PushDomains,omitnil,omitempty" name:"PushDomains"`
}

func (r *DescribeLiveTimeShiftBillInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTimeShiftBillInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PushDomains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveTimeShiftBillInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveTimeShiftBillInfoListResponseParams struct {
	// 时移计费明细数据。
	DataInfoList []*TimeShiftBillData `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveTimeShiftBillInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveTimeShiftBillInfoListResponseParams `json:"Response"`
}

func (r *DescribeLiveTimeShiftBillInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTimeShiftBillInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveTimeShiftRulesRequestParams struct {

}

type DescribeLiveTimeShiftRulesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLiveTimeShiftRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTimeShiftRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveTimeShiftRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveTimeShiftRulesResponseParams struct {
	// 规则信息列表。
	Rules []*RuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveTimeShiftRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveTimeShiftRulesResponseParams `json:"Response"`
}

func (r *DescribeLiveTimeShiftRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTimeShiftRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveTimeShiftTemplatesRequestParams struct {

}

type DescribeLiveTimeShiftTemplatesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLiveTimeShiftTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTimeShiftTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveTimeShiftTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveTimeShiftTemplatesResponseParams struct {
	// 直播时移模板信息。
	Templates []*TimeShiftTemplate `json:"Templates,omitnil,omitempty" name:"Templates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveTimeShiftTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveTimeShiftTemplatesResponseParams `json:"Response"`
}

func (r *DescribeLiveTimeShiftTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTimeShiftTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveTimeShiftWriteSizeInfoListRequestParams struct {
	// 起始时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 支持最近六个月的查询，查询开始和结束时间跨度不支持超过31天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 支持最近六个月的查询，查询开始和结束时间跨度不支持超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 域名。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// 维度
	// Area地区、Domain 域名、StorageDays 时移天数
	Dimensions []*string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// 时移天数。
	StorageDays []*int64 `json:"StorageDays,omitnil,omitempty" name:"StorageDays"`

	// 时间跨度（分钟）
	// 默认5，可选 5、60或者1440。
	Granularity *int64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 区域
	// 可选Mainland、Oversea。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`
}

type DescribeLiveTimeShiftWriteSizeInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 支持最近六个月的查询，查询开始和结束时间跨度不支持超过31天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 支持最近六个月的查询，查询开始和结束时间跨度不支持超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 域名。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// 维度
	// Area地区、Domain 域名、StorageDays 时移天数
	Dimensions []*string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// 时移天数。
	StorageDays []*int64 `json:"StorageDays,omitnil,omitempty" name:"StorageDays"`

	// 时间跨度（分钟）
	// 默认5，可选 5、60或者1440。
	Granularity *int64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 区域
	// 可选Mainland、Oversea。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`
}

func (r *DescribeLiveTimeShiftWriteSizeInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTimeShiftWriteSizeInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DomainNames")
	delete(f, "Dimensions")
	delete(f, "StorageDays")
	delete(f, "Granularity")
	delete(f, "MainlandOrOversea")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveTimeShiftWriteSizeInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveTimeShiftWriteSizeInfoListResponseParams struct {
	// 直播时移写入量数据明细。
	DataInfoList []*TimeShiftWriteSizeData `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveTimeShiftWriteSizeInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveTimeShiftWriteSizeInfoListResponseParams `json:"Response"`
}

func (r *DescribeLiveTimeShiftWriteSizeInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveTimeShiftWriteSizeInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveTranscodeDetailInfoRequestParams struct {
	// 推流域名。
	PushDomain *string `json:"PushDomain,omitnil,omitempty" name:"PushDomain"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 查询时间，北京时间，
	// 格式：yyyymmdd。
	// 注意：支持查询近1个月内某天的详细数据，截止到昨天。
	DayTime *string `json:"DayTime,omitnil,omitempty" name:"DayTime"`

	// 页数，默认1，
	// 不超过100页。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，默认20，
	// 范围：[10,1000]。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 起始天时间，北京时间，
	// 格式：yyyymmdd。
	// 注意：支持查询近1个月内的详细数据。
	StartDayTime *string `json:"StartDayTime,omitnil,omitempty" name:"StartDayTime"`

	// 结束天时间，北京时间，
	// 格式：yyyymmdd。
	// 注意：支持查询近1个月内的详细数据，截止到昨天，注意DayTime 与（StartDayTime，EndDayTime）必须要传一个，如果都传，会以DayTime为准 。
	EndDayTime *string `json:"EndDayTime,omitnil,omitempty" name:"EndDayTime"`
}

type DescribeLiveTranscodeDetailInfoRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	PushDomain *string `json:"PushDomain,omitnil,omitempty" name:"PushDomain"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 查询时间，北京时间，
	// 格式：yyyymmdd。
	// 注意：支持查询近1个月内某天的详细数据，截止到昨天。
	DayTime *string `json:"DayTime,omitnil,omitempty" name:"DayTime"`

	// 页数，默认1，
	// 不超过100页。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，默认20，
	// 范围：[10,1000]。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 起始天时间，北京时间，
	// 格式：yyyymmdd。
	// 注意：支持查询近1个月内的详细数据。
	StartDayTime *string `json:"StartDayTime,omitnil,omitempty" name:"StartDayTime"`

	// 结束天时间，北京时间，
	// 格式：yyyymmdd。
	// 注意：支持查询近1个月内的详细数据，截止到昨天，注意DayTime 与（StartDayTime，EndDayTime）必须要传一个，如果都传，会以DayTime为准 。
	EndDayTime *string `json:"EndDayTime,omitnil,omitempty" name:"EndDayTime"`
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

// Predefined struct for user
type DescribeLiveTranscodeDetailInfoResponseParams struct {
	// 统计数据列表。
	DataInfoList []*TranscodeDetailInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 页码。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总个数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 总页数。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveTranscodeDetailInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveTranscodeDetailInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveTranscodeRulesRequestParams struct {
	// 要筛选的模板ID数组。
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// 要筛选的域名数组。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`
}

type DescribeLiveTranscodeRulesRequest struct {
	*tchttp.BaseRequest
	
	// 要筛选的模板ID数组。
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// 要筛选的域名数组。
	DomainNames []*string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`
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

// Predefined struct for user
type DescribeLiveTranscodeRulesResponseParams struct {
	// 转码规则列表。
	Rules []*RuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveTranscodeRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveTranscodeRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveTranscodeTemplateRequestParams struct {
	// 模板 ID。
	// 注意：在创建转码模板接口 [CreateLiveTranscodeTemplate](/document/product/267/32646) 调用的返回值中获取模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID。
	// 注意：在创建转码模板接口 [CreateLiveTranscodeTemplate](/document/product/267/32646) 调用的返回值中获取模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type DescribeLiveTranscodeTemplateResponseParams struct {
	// 模板信息。
	Template *TemplateInfo `json:"Template,omitnil,omitempty" name:"Template"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveTranscodeTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveTranscodeTemplatesRequestParams struct {
	// 转码模板类型，默认0。
	// 0：普通转码模板。
	// 1：自适应码率转码模板。
	TemplateType *int64 `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`
}

type DescribeLiveTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 转码模板类型，默认0。
	// 0：普通转码模板。
	// 1：自适应码率转码模板。
	TemplateType *int64 `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`
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
	delete(f, "TemplateType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveTranscodeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveTranscodeTemplatesResponseParams struct {
	// 转码模板列表。
	Templates []*TemplateInfo `json:"Templates,omitnil,omitempty" name:"Templates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveTranscodeTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveTranscodeTotalInfoRequestParams struct {
	// 开始时间，支持最近三个月的数据查询，起始时间和结束时间跨度不能超过三个月。接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，支持最近三个月的数据查询，起始时间和结束时间跨度不能超过三个月。接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 推流域名列表，若不填，表示查询所有域名总体数据。
	// 指定域名时返回1小时粒度数据。
	PushDomains []*string `json:"PushDomains,omitnil,omitempty" name:"PushDomains"`

	// 可选值：
	// Mainland：查询中国大陆（境内）数据，
	// Oversea：则查询国际/港澳台（境外）数据，
	// 默认：查询全球地区（境内+境外）的数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`
}

type DescribeLiveTranscodeTotalInfoRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，支持最近三个月的数据查询，起始时间和结束时间跨度不能超过三个月。接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，支持最近三个月的数据查询，起始时间和结束时间跨度不能超过三个月。接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 推流域名列表，若不填，表示查询所有域名总体数据。
	// 指定域名时返回1小时粒度数据。
	PushDomains []*string `json:"PushDomains,omitnil,omitempty" name:"PushDomains"`

	// 可选值：
	// Mainland：查询中国大陆（境内）数据，
	// Oversea：则查询国际/港澳台（境外）数据，
	// 默认：查询全球地区（境内+境外）的数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`
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

// Predefined struct for user
type DescribeLiveTranscodeTotalInfoResponseParams struct {
	// 统计数据列表。
	DataInfoList []*TranscodeTotalInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveTranscodeTotalInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveTranscodeTotalInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveWatermarkRequestParams struct {
	// DescribeLiveWatermarks接口返回的水印 ID。
	WatermarkId *uint64 `json:"WatermarkId,omitnil,omitempty" name:"WatermarkId"`
}

type DescribeLiveWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// DescribeLiveWatermarks接口返回的水印 ID。
	WatermarkId *uint64 `json:"WatermarkId,omitnil,omitempty" name:"WatermarkId"`
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

// Predefined struct for user
type DescribeLiveWatermarkResponseParams struct {
	// 水印信息。
	Watermark *WatermarkInfo `json:"Watermark,omitnil,omitempty" name:"Watermark"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveWatermarkResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveWatermarkRulesRequestParams struct {

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

// Predefined struct for user
type DescribeLiveWatermarkRulesResponseParams struct {
	// 水印规则列表。
	Rules []*RuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveWatermarkRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveWatermarkRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveWatermarksRequestParams struct {
	// 水印名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页的索引参数，从1开始
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 分页的大小参数，默认值500
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeLiveWatermarksRequest struct {
	*tchttp.BaseRequest
	
	// 水印名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页的索引参数，从1开始
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 分页的大小参数，默认值500
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
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
	delete(f, "Name")
	delete(f, "PageNo")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveWatermarksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveWatermarksResponseParams struct {
	// 水印总个数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 水印信息列表。
	WatermarkList []*WatermarkInfo `json:"WatermarkList,omitnil,omitempty" name:"WatermarkList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveWatermarksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveWatermarksResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveXP2PDetailInfoListRequestParams struct {
	// utc分钟粒度查询时间，查询某一分钟的用量数据，格式为：yyyy-mm-ddTHH:MM:00Z，参考https://cloud.tencent.com/document/product/266/11732#I，
	// 例如：北京时间2019-01-08 10:00:00，对应utc时间为：2019-01-08T10:00:00+08:00。
	// 
	// 支持最近六个月的查询。
	QueryTime *string `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// 类型数组，分直播live和点播vod，不传默认查全部。
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// 查询流数组，不传默认查所有流。
	StreamNames []*string `json:"StreamNames,omitnil,omitempty" name:"StreamNames"`

	// 查询维度，不传该参数则默认查询流维度的数据，传递该参数则只查对应维度的数据，和返回值的字段相关，目前支持AppId维度查询。
	Dimension []*string `json:"Dimension,omitnil,omitempty" name:"Dimension"`
}

type DescribeLiveXP2PDetailInfoListRequest struct {
	*tchttp.BaseRequest
	
	// utc分钟粒度查询时间，查询某一分钟的用量数据，格式为：yyyy-mm-ddTHH:MM:00Z，参考https://cloud.tencent.com/document/product/266/11732#I，
	// 例如：北京时间2019-01-08 10:00:00，对应utc时间为：2019-01-08T10:00:00+08:00。
	// 
	// 支持最近六个月的查询。
	QueryTime *string `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// 类型数组，分直播live和点播vod，不传默认查全部。
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// 查询流数组，不传默认查所有流。
	StreamNames []*string `json:"StreamNames,omitnil,omitempty" name:"StreamNames"`

	// 查询维度，不传该参数则默认查询流维度的数据，传递该参数则只查对应维度的数据，和返回值的字段相关，目前支持AppId维度查询。
	Dimension []*string `json:"Dimension,omitnil,omitempty" name:"Dimension"`
}

func (r *DescribeLiveXP2PDetailInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveXP2PDetailInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueryTime")
	delete(f, "Type")
	delete(f, "StreamNames")
	delete(f, "Dimension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveXP2PDetailInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLiveXP2PDetailInfoListResponseParams struct {
	// P2P流统计信息。
	DataInfoList []*XP2PDetailInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLiveXP2PDetailInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveXP2PDetailInfoListResponseParams `json:"Response"`
}

func (r *DescribeLiveXP2PDetailInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveXP2PDetailInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogDownloadListRequestParams struct {
	// 开始时间，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	// 注：此字段为北京时间（UTC+8时区）。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	// 注意：结束时间 - 开始时间 <=7天。
	// 注：此字段为北京时间（UTC+8时区）。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 域名列表。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 快直播还是标准直播，0：标准直播，1：快直播。默认为0。
	IsFastLive *int64 `json:"IsFastLive,omitnil,omitempty" name:"IsFastLive"`
}

type DescribeLogDownloadListRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	// 注：此字段为北京时间（UTC+8时区）。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	// 注意：结束时间 - 开始时间 <=7天。
	// 注：此字段为北京时间（UTC+8时区）。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 域名列表。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 快直播还是标准直播，0：标准直播，1：快直播。默认为0。
	IsFastLive *int64 `json:"IsFastLive,omitnil,omitempty" name:"IsFastLive"`
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
	delete(f, "IsFastLive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogDownloadListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogDownloadListResponseParams struct {
	// 日志信息列表。
	LogInfoList []*LogInfo `json:"LogInfoList,omitnil,omitempty" name:"LogInfoList"`

	// 总条数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogDownloadListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogDownloadListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeMonitorReportRequestParams struct {
	// 监播任务ID。
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

type DescribeMonitorReportRequest struct {
	*tchttp.BaseRequest
	
	// 监播任务ID。
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

func (r *DescribeMonitorReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMonitorReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorReportResponseParams struct {
	// 媒体处理结果信息。
	MPSResult *MPSResult `json:"MPSResult,omitnil,omitempty" name:"MPSResult"`

	// 媒体诊断结果信息。
	DiagnoseResult *DiagnoseResult `json:"DiagnoseResult,omitnil,omitempty" name:"DiagnoseResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMonitorReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMonitorReportResponseParams `json:"Response"`
}

func (r *DescribeMonitorReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlayErrorCodeDetailInfoListRequestParams struct {
	// 起始时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 注：EndTime 和 StartTime 只支持最近1天的数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询粒度：
	// 1-1分钟粒度。
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 是，可选值包括”4xx”,”5xx”，支持”4xx,5xx”等这种混合模式。
	StatType *string `json:"StatType,omitnil,omitempty" name:"StatType"`

	// 播放域名列表。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`
}

type DescribePlayErrorCodeDetailInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 注：EndTime 和 StartTime 只支持最近1天的数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询粒度：
	// 1-1分钟粒度。
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 是，可选值包括”4xx”,”5xx”，支持”4xx,5xx”等这种混合模式。
	StatType *string `json:"StatType,omitnil,omitempty" name:"StatType"`

	// 播放域名列表。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`
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

// Predefined struct for user
type DescribePlayErrorCodeDetailInfoListResponseParams struct {
	// 统计信息列表。
	HttpCodeList []*HttpCodeInfo `json:"HttpCodeList,omitnil,omitempty" name:"HttpCodeList"`

	// 统计类型。
	StatType *string `json:"StatType,omitnil,omitempty" name:"StatType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePlayErrorCodeDetailInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePlayErrorCodeDetailInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePlayErrorCodeSumInfoListRequestParams struct {
	// 起始时间点，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	// 注：EndTime 和 StartTime 只支持最近1天的数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 播放域名列表，不填表示总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 页数，范围[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，范围：[1,1000]，默认值是20。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 分组参数，可选值：CountryProIsp（默认值），Country（国家），默认是按照国家+省份+运营商来进行分组；目前国外的省份和运营商暂时无法识别。
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 输出字段使用的语言，可选值：Chinese（默认值），English，目前国家，省份和运营商支持多语言。
	OutLanguage *string `json:"OutLanguage,omitnil,omitempty" name:"OutLanguage"`
}

type DescribePlayErrorCodeSumInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，北京时间。
	// 格式：yyyy-mm-dd HH:MM:SS。
	// 注：EndTime 和 StartTime 只支持最近1天的数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 播放域名列表，不填表示总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 页数，范围[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，范围：[1,1000]，默认值是20。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 分组参数，可选值：CountryProIsp（默认值），Country（国家），默认是按照国家+省份+运营商来进行分组；目前国外的省份和运营商暂时无法识别。
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 输出字段使用的语言，可选值：Chinese（默认值），English，目前国家，省份和运营商支持多语言。
	OutLanguage *string `json:"OutLanguage,omitnil,omitempty" name:"OutLanguage"`
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

// Predefined struct for user
type DescribePlayErrorCodeSumInfoListResponseParams struct {
	// 分省份分运营商错误码为2或3或4或5开头的状态码数据信息。
	ProIspInfoList []*ProIspPlayCodeDataInfo `json:"ProIspInfoList,omitnil,omitempty" name:"ProIspInfoList"`

	// 所有状态码的加和的次数。
	TotalCodeAll *uint64 `json:"TotalCodeAll,omitnil,omitempty" name:"TotalCodeAll"`

	// 状态码为4开头的总次数。
	TotalCode4xx *uint64 `json:"TotalCode4xx,omitnil,omitempty" name:"TotalCode4xx"`

	// 状态码为5开头的总次数。
	TotalCode5xx *uint64 `json:"TotalCode5xx,omitnil,omitempty" name:"TotalCode5xx"`

	// 各状态码的总次数。
	TotalCodeList []*PlayCodeTotalInfo `json:"TotalCodeList,omitnil,omitempty" name:"TotalCodeList"`

	// 页号。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页大小。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 总记录数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 状态码为2开头的总次数。
	TotalCode2xx *uint64 `json:"TotalCode2xx,omitnil,omitempty" name:"TotalCode2xx"`

	// 状态码为3开头的总次数。
	TotalCode3xx *uint64 `json:"TotalCode3xx,omitnil,omitempty" name:"TotalCode3xx"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePlayErrorCodeSumInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePlayErrorCodeSumInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeProIspPlaySumInfoListRequestParams struct {
	// 起始时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 注：EndTime 和 StartTime 只支持最近1天的数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 统计的类型，可选值：”Province”(省份)，”Isp”(运营商)，“CountryOrArea”(国家或地区)。
	StatType *string `json:"StatType,omitnil,omitempty" name:"StatType"`

	// 播放域名列表，不填则为全部。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 页号，范围是[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，范围是[1,1000]，默认值是20。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 输出字段使用的语言，可选值：Chinese（默认值），English；目前国家，省份和运营商支持多语言。
	OutLanguage *string `json:"OutLanguage,omitnil,omitempty" name:"OutLanguage"`
}

type DescribeProIspPlaySumInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 注：EndTime 和 StartTime 只支持最近1天的数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 统计的类型，可选值：”Province”(省份)，”Isp”(运营商)，“CountryOrArea”(国家或地区)。
	StatType *string `json:"StatType,omitnil,omitempty" name:"StatType"`

	// 播放域名列表，不填则为全部。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 页号，范围是[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，范围是[1,1000]，默认值是20。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 输出字段使用的语言，可选值：Chinese（默认值），English；目前国家，省份和运营商支持多语言。
	OutLanguage *string `json:"OutLanguage,omitnil,omitempty" name:"OutLanguage"`
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

// Predefined struct for user
type DescribeProIspPlaySumInfoListResponseParams struct {
	// 总流量。
	TotalFlux *float64 `json:"TotalFlux,omitnil,omitempty" name:"TotalFlux"`

	// 总请求数。
	TotalRequest *uint64 `json:"TotalRequest,omitnil,omitempty" name:"TotalRequest"`

	// 统计的类型。
	StatType *string `json:"StatType,omitnil,omitempty" name:"StatType"`

	// 每页的记录数。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页号。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 总记录数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 总页数。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 省份，运营商，国家或地区汇总数据列表。
	DataInfoList []*ProIspPlaySumInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 下载速度，单位：MB/s，计算方式：总流量/总时长。
	AvgFluxPerSecond *float64 `json:"AvgFluxPerSecond,omitnil,omitempty" name:"AvgFluxPerSecond"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProIspPlaySumInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProIspPlaySumInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeProvinceIspPlayInfoListRequestParams struct {
	// 起始时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 注：EndTime 和 StartTime 只支持最近1天的数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 支持如下粒度：
	// 1：1分钟粒度（跨度不支持超过1天）
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 统计指标类型：
	// “Bandwidth”：带宽
	// “FluxPerSecond”：平均流量
	// “Flux”：流量
	// “Request”：请求数
	// “Online”：并发连接数
	StatType *string `json:"StatType,omitnil,omitempty" name:"StatType"`

	// 播放域名列表。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 要查询的省份（地区）英文名称列表，如 Beijing。
	ProvinceNames []*string `json:"ProvinceNames,omitnil,omitempty" name:"ProvinceNames"`

	// 要查询的运营商英文名称列表，如 China Mobile ，如果为空，查询所有运营商的数据。
	IspNames []*string `json:"IspNames,omitnil,omitempty" name:"IspNames"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// ip类型：
	// “Ipv6”：Ipv6数据
	// 如果为空，查询总的数据；
	IpType *string `json:"IpType,omitnil,omitempty" name:"IpType"`
}

type DescribeProvinceIspPlayInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 注：EndTime 和 StartTime 只支持最近1天的数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 支持如下粒度：
	// 1：1分钟粒度（跨度不支持超过1天）
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 统计指标类型：
	// “Bandwidth”：带宽
	// “FluxPerSecond”：平均流量
	// “Flux”：流量
	// “Request”：请求数
	// “Online”：并发连接数
	StatType *string `json:"StatType,omitnil,omitempty" name:"StatType"`

	// 播放域名列表。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 要查询的省份（地区）英文名称列表，如 Beijing。
	ProvinceNames []*string `json:"ProvinceNames,omitnil,omitempty" name:"ProvinceNames"`

	// 要查询的运营商英文名称列表，如 China Mobile ，如果为空，查询所有运营商的数据。
	IspNames []*string `json:"IspNames,omitnil,omitempty" name:"IspNames"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// ip类型：
	// “Ipv6”：Ipv6数据
	// 如果为空，查询总的数据；
	IpType *string `json:"IpType,omitnil,omitempty" name:"IpType"`
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

// Predefined struct for user
type DescribeProvinceIspPlayInfoListResponseParams struct {
	// 播放信息列表。
	DataInfoList []*PlayStatInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 统计的类型，和输入参数保持一致。
	StatType *string `json:"StatType,omitnil,omitempty" name:"StatType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProvinceIspPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProvinceIspPlayInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePullStreamConfigsRequestParams struct {
	// 配置 ID。
	// 获取途径：从 CreatePullStreamConfig 接口返回值获取。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

type DescribePullStreamConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 配置 ID。
	// 获取途径：从 CreatePullStreamConfig 接口返回值获取。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
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

// Predefined struct for user
type DescribePullStreamConfigsResponseParams struct {
	// 拉流配置。
	PullStreamConfigs []*PullStreamConfig `json:"PullStreamConfigs,omitnil,omitempty" name:"PullStreamConfigs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePullStreamConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePullStreamConfigsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePullTransformPushInfoRequestParams struct {
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 支持最近三年的查询，查询开始和结束时间跨度不支持超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 可选值：
	// Mainland：查询国内数据，
	// Oversea：则查询国外数据，
	// 默认：查询国内+国外的数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 可选值：  PullLivePushLive：拉流源类型为直播  PullVodPushLive：拉流源类型为点播  PullPicPushLive：拉流源类型为图片  默认：查询全部拉流源类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

type DescribePullTransformPushInfoRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 支持最近三年的查询，查询开始和结束时间跨度不支持超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 可选值：
	// Mainland：查询国内数据，
	// Oversea：则查询国外数据，
	// 默认：查询国内+国外的数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 可选值：  PullLivePushLive：拉流源类型为直播  PullVodPushLive：拉流源类型为点播  PullPicPushLive：拉流源类型为图片  默认：查询全部拉流源类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

func (r *DescribePullTransformPushInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePullTransformPushInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MainlandOrOversea")
	delete(f, "SourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePullTransformPushInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePullTransformPushInfoResponseParams struct {
	// 时长统计数据列表
	DataInfoList []*TaskDurationInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 拉流转推得总时长
	TotalDuration *uint64 `json:"TotalDuration,omitnil,omitempty" name:"TotalDuration"`

	// 拉流转推得总时长（秒）
	TotalDurationSecond *uint64 `json:"TotalDurationSecond,omitnil,omitempty" name:"TotalDurationSecond"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePullTransformPushInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePullTransformPushInfoResponseParams `json:"Response"`
}

func (r *DescribePullTransformPushInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePullTransformPushInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePushBandwidthAndFluxListRequestParams struct {
	// 查询开始时间点，精确到分钟粒度，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#52)。
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 支持最近两个月的查询，查询开始和结束时间跨度不支持超过31天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间点，精确到分钟粒度，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#52)。
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 支持最近两个月的查询，查询开始和结束时间跨度不支持超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 域名，可以填多个，若不填，表示总体数据。
	PushDomains []*string `json:"PushDomains,omitnil,omitempty" name:"PushDomains"`

	// 可选值：
	// Mainland：查询中国大陆（境内）数据，
	// Oversea：则查询国际/港澳台（境外）数据，
	// 不填则默认查询全球地区（境内+境外）的数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 数据粒度，支持如下粒度：
	// 5：5分钟粒度，（跨度不支持超过1天），
	// 60：1小时粒度（跨度不支持超过一个月），
	// 1440：天粒度（跨度不支持超过一个月）。
	// 默认值：5。
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

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
	RegionNames []*string `json:"RegionNames,omitnil,omitempty" name:"RegionNames"`

	// 国家，映射表参照如下文档：
	// https://cloud.tencent.com/document/product/267/34019。
	CountryNames []*string `json:"CountryNames,omitnil,omitempty" name:"CountryNames"`
}

type DescribePushBandwidthAndFluxListRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间点，精确到分钟粒度，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#52)。
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 支持最近两个月的查询，查询开始和结束时间跨度不支持超过31天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间点，精确到分钟粒度，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#52)。
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 支持最近两个月的查询，查询开始和结束时间跨度不支持超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 域名，可以填多个，若不填，表示总体数据。
	PushDomains []*string `json:"PushDomains,omitnil,omitempty" name:"PushDomains"`

	// 可选值：
	// Mainland：查询中国大陆（境内）数据，
	// Oversea：则查询国际/港澳台（境外）数据，
	// 不填则默认查询全球地区（境内+境外）的数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 数据粒度，支持如下粒度：
	// 5：5分钟粒度，（跨度不支持超过1天），
	// 60：1小时粒度（跨度不支持超过一个月），
	// 1440：天粒度（跨度不支持超过一个月）。
	// 默认值：5。
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

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
	RegionNames []*string `json:"RegionNames,omitnil,omitempty" name:"RegionNames"`

	// 国家，映射表参照如下文档：
	// https://cloud.tencent.com/document/product/267/34019。
	CountryNames []*string `json:"CountryNames,omitnil,omitempty" name:"CountryNames"`
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
	delete(f, "CountryNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePushBandwidthAndFluxListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePushBandwidthAndFluxListResponseParams struct {
	// 峰值带宽所在时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	PeakBandwidthTime *string `json:"PeakBandwidthTime,omitnil,omitempty" name:"PeakBandwidthTime"`

	// 峰值带宽，单位是 Mbps。
	PeakBandwidth *float64 `json:"PeakBandwidth,omitnil,omitempty" name:"PeakBandwidth"`

	// 95峰值带宽所在时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	P95PeakBandwidthTime *string `json:"P95PeakBandwidthTime,omitnil,omitempty" name:"P95PeakBandwidthTime"`

	// 95峰值带宽，单位是 Mbps。
	P95PeakBandwidth *float64 `json:"P95PeakBandwidth,omitnil,omitempty" name:"P95PeakBandwidth"`

	// 总流量，单位是 MB。
	SumFlux *float64 `json:"SumFlux,omitnil,omitempty" name:"SumFlux"`

	// 明细数据信息。
	DataInfoList []*BillDataInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePushBandwidthAndFluxListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePushBandwidthAndFluxListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRecordTaskRequestParams struct {
	// 查询任务开始时间，Unix 时间戳。设置时间不早于当前时间之前90天的时间，且查询时间跨度不超过一周。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询任务结束时间，Unix 时间戳。EndTime 必须大于 StartTime，设置时间不早于当前时间之前90天的时间，且查询时间跨度不超过一周。（注意：任务开始结束时间必须在查询时间范围内）。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`
}

type DescribeRecordTaskRequest struct {
	*tchttp.BaseRequest
	
	// 查询任务开始时间，Unix 时间戳。设置时间不早于当前时间之前90天的时间，且查询时间跨度不超过一周。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询任务结束时间，Unix 时间戳。EndTime 必须大于 StartTime，设置时间不早于当前时间之前90天的时间，且查询时间跨度不超过一周。（注意：任务开始结束时间必须在查询时间范围内）。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`
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

// Predefined struct for user
type DescribeRecordTaskResponseParams struct {
	// 翻页标识，当请求未返回所有数据，该字段表示下一条记录的 Token。当该字段为空，说明已无更多数据。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 录制任务列表。当该字段为空，说明已返回所有数据。
	TaskList []*RecordTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeScreenShotSheetNumListRequestParams struct {
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 支持最近三个月的查询，查询开始和结束时间跨度不支持超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 地域信息，可选值包括Mainland，Oversea，前者是查询中国大陆范围内的数据，后者是除中国大陆范围之外的数据，若不传该参数，则查询所有地区的数据。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 推流域名（支持查询2019年11 月1日之后的域名维度数据）。
	PushDomains []*string `json:"PushDomains,omitnil,omitempty" name:"PushDomains"`

	// 数据维度，数据延迟1个半小时，可选值包括：1、Minute（5分钟粒度，最大支持查询时间范围是31天），2、Day（天粒度，默认值，按照北京时间做跨天处理，最大支持查询时间范围是186天当天）。
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeScreenShotSheetNumListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 支持最近三个月的查询，查询开始和结束时间跨度不支持超过31天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 地域信息，可选值包括Mainland，Oversea，前者是查询中国大陆范围内的数据，后者是除中国大陆范围之外的数据，若不传该参数，则查询所有地区的数据。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 推流域名（支持查询2019年11 月1日之后的域名维度数据）。
	PushDomains []*string `json:"PushDomains,omitnil,omitempty" name:"PushDomains"`

	// 数据维度，数据延迟1个半小时，可选值包括：1、Minute（5分钟粒度，最大支持查询时间范围是31天），2、Day（天粒度，默认值，按照北京时间做跨天处理，最大支持查询时间范围是186天当天）。
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
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

// Predefined struct for user
type DescribeScreenShotSheetNumListResponseParams struct {
	// 数据信息列表。
	DataInfoList []*TimeValue `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScreenShotSheetNumListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScreenShotSheetNumListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeScreenshotTaskRequestParams struct {
	// 查询任务开始时间，Unix 时间戳。设置时间不早于当前时间之前90天的时间，且查询时间跨度不超过一周。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询任务结束时间，Unix 时间戳。EndTime 必须大于 StartTime，设置时间不早于当前时间之前90天的时间，且查询时间跨度不超过一周。（注意：任务开始结束时间必须在查询时间范围内）。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`
}

type DescribeScreenshotTaskRequest struct {
	*tchttp.BaseRequest
	
	// 查询任务开始时间，Unix 时间戳。设置时间不早于当前时间之前90天的时间，且查询时间跨度不超过一周。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询任务结束时间，Unix 时间戳。EndTime 必须大于 StartTime，设置时间不早于当前时间之前90天的时间，且查询时间跨度不超过一周。（注意：任务开始结束时间必须在查询时间范围内）。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`
}

func (r *DescribeScreenshotTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScreenshotTaskRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScreenshotTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScreenshotTaskResponseParams struct {
	// 翻页标识，当请求未返回所有数据，该字段表示下一条记录的 Token。当该字段为空，说明已无更多数据。
	ScrollToken *string `json:"ScrollToken,omitnil,omitempty" name:"ScrollToken"`

	// 截图任务列表。当该字段为空，说明已返回所有数据。
	TaskList []*ScreenshotTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScreenshotTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScreenshotTaskResponseParams `json:"Response"`
}

func (r *DescribeScreenshotTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScreenshotTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamDayPlayInfoListRequestParams struct {
	// 日期，格式：YYYY-mm-dd。
	// 第二天凌晨3点出昨天的数据，建议在这个时间点之后查询最新数据。支持最近3个月的数据查询。
	DayTime *string `json:"DayTime,omitnil,omitempty" name:"DayTime"`

	// 播放域名。
	PlayDomain *string `json:"PlayDomain,omitnil,omitempty" name:"PlayDomain"`

	// 页号，范围[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，范围[100,1000]，默认值是1000。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 可选值：
	// Mainland：查询国内数据，
	// Oversea：则查询国外数据，
	// 默认：查询国内+国外的数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 服务名称，可选值包括LVB(标准直播)，LEB(快直播)，不填则查LVB+LEB总值。
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

type DescribeStreamDayPlayInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 日期，格式：YYYY-mm-dd。
	// 第二天凌晨3点出昨天的数据，建议在这个时间点之后查询最新数据。支持最近3个月的数据查询。
	DayTime *string `json:"DayTime,omitnil,omitempty" name:"DayTime"`

	// 播放域名。
	PlayDomain *string `json:"PlayDomain,omitnil,omitempty" name:"PlayDomain"`

	// 页号，范围[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，范围[100,1000]，默认值是1000。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 可选值：
	// Mainland：查询国内数据，
	// Oversea：则查询国外数据，
	// 默认：查询国内+国外的数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 服务名称，可选值包括LVB(标准直播)，LEB(快直播)，不填则查LVB+LEB总值。
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
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

// Predefined struct for user
type DescribeStreamDayPlayInfoListResponseParams struct {
	// 播放数据信息列表。
	DataInfoList []*PlayDataInfoByStream `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 总数量。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 总页数。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 当前数据所处页码。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamDayPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamDayPlayInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeStreamPlayInfoListRequestParams struct {
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）yyyy-MM-dd HH:mm:ss：使用此格式时，默认代表北京时间。
	// 开始时间和结束时间的格式需要保持一致。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）yyyy-MM-dd HH:mm:ss：使用此格式时，默认代表北京时间。
	// 开始时间和结束时间的格式需要保持一致。结束时间和开始时间跨度不支持超过24小时，支持距当前时间一个月内的数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 播放域名，
	// 若不填，则为查询所有播放域名的在线流数据。
	PlayDomain *string `json:"PlayDomain,omitnil,omitempty" name:"PlayDomain"`

	// 流名称，精确匹配。
	// 若不填，则为查询总体播放数据。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流路径，与播放地址中的AppName保持一致，会精确匹配，在同时传递了StreamName时生效。
	// 若不填，则为查询总体播放数据。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 服务名称，可选值包括LVB(标准直播)，LEB(快直播)，不填则查LVB+LEB总值。
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

type DescribeStreamPlayInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）yyyy-MM-dd HH:mm:ss：使用此格式时，默认代表北京时间。
	// 开始时间和结束时间的格式需要保持一致。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）yyyy-MM-dd HH:mm:ss：使用此格式时，默认代表北京时间。
	// 开始时间和结束时间的格式需要保持一致。结束时间和开始时间跨度不支持超过24小时，支持距当前时间一个月内的数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 播放域名，
	// 若不填，则为查询所有播放域名的在线流数据。
	PlayDomain *string `json:"PlayDomain,omitnil,omitempty" name:"PlayDomain"`

	// 流名称，精确匹配。
	// 若不填，则为查询总体播放数据。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流路径，与播放地址中的AppName保持一致，会精确匹配，在同时传递了StreamName时生效。
	// 若不填，则为查询总体播放数据。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 服务名称，可选值包括LVB(标准直播)，LEB(快直播)，不填则查LVB+LEB总值。
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
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

// Predefined struct for user
type DescribeStreamPlayInfoListResponseParams struct {
	// 统计信息列表，时间粒度是1分钟。
	DataInfoList []*DayStreamPlayInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamPlayInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeStreamPushInfoListRequestParams struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 起始时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 支持查询最近7天数据，建议查询时间跨度在3小时之内。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 推流域名。
	PushDomain *string `json:"PushDomain,omitnil,omitempty" name:"PushDomain"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`
}

type DescribeStreamPushInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 起始时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 支持查询最近7天数据，建议查询时间跨度在3小时之内。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 推流域名。
	PushDomain *string `json:"PushDomain,omitnil,omitempty" name:"PushDomain"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`
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

// Predefined struct for user
type DescribeStreamPushInfoListResponseParams struct {
	// 返回的数据列表。
	DataInfoList []*PushQualityData `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamPushInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamPushInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTimeShiftRecordDetailRequestParams struct {
	// 推流域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 查询范围起始时间，Unix 时间戳。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询范围终止时间，Unix 时间戳。 
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 推流域名所属组，没有域名组或者域名组为空字符串可不填。
	DomainGroup *string `json:"DomainGroup,omitnil,omitempty" name:"DomainGroup"`

	// 转码模板ID，转码模板ID为0可不填。
	TransCodeId *uint64 `json:"TransCodeId,omitnil,omitempty" name:"TransCodeId"`
}

type DescribeTimeShiftRecordDetailRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 查询范围起始时间，Unix 时间戳。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询范围终止时间，Unix 时间戳。 
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 推流域名所属组，没有域名组或者域名组为空字符串可不填。
	DomainGroup *string `json:"DomainGroup,omitnil,omitempty" name:"DomainGroup"`

	// 转码模板ID，转码模板ID为0可不填。
	TransCodeId *uint64 `json:"TransCodeId,omitnil,omitempty" name:"TransCodeId"`
}

func (r *DescribeTimeShiftRecordDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimeShiftRecordDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "AppName")
	delete(f, "StreamName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DomainGroup")
	delete(f, "TransCodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimeShiftRecordDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimeShiftRecordDetailResponseParams struct {
	// 时移录制会话数组。
	RecordList []*TimeShiftRecord `json:"RecordList,omitnil,omitempty" name:"RecordList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTimeShiftRecordDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimeShiftRecordDetailResponseParams `json:"Response"`
}

func (r *DescribeTimeShiftRecordDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimeShiftRecordDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimeShiftStreamListRequestParams struct {
	// 查询范围起始时间，Unix 时间戳。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询范围结束时间，Unix 时间戳。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 推流域名所属域名组。
	DomainGroup *string `json:"DomainGroup,omitnil,omitempty" name:"DomainGroup"`

	// 用户指定要返回的最大结果数，取值范围[0,100]，不指定或者指定为0时，API 
	// 默认值为100。指定超过100时，API 强制使用100。指定值为负数时，接口返回错误。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 指定拉取的页码，不传时默认为1。
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`
}

type DescribeTimeShiftStreamListRequest struct {
	*tchttp.BaseRequest
	
	// 查询范围起始时间，Unix 时间戳。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询范围结束时间，Unix 时间戳。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 推流域名所属域名组。
	DomainGroup *string `json:"DomainGroup,omitnil,omitempty" name:"DomainGroup"`

	// 用户指定要返回的最大结果数，取值范围[0,100]，不指定或者指定为0时，API 
	// 默认值为100。指定超过100时，API 强制使用100。指定值为负数时，接口返回错误。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 指定拉取的页码，不传时默认为1。
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`
}

func (r *DescribeTimeShiftStreamListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimeShiftStreamListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "StreamName")
	delete(f, "Domain")
	delete(f, "DomainGroup")
	delete(f, "PageSize")
	delete(f, "PageNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimeShiftStreamListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimeShiftStreamListResponseParams struct {
	// 时间段内所有的数据量。
	TotalSize *int64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 流列表。
	StreamList []*TimeShiftStreamInfo `json:"StreamList,omitnil,omitempty" name:"StreamList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTimeShiftStreamListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimeShiftStreamListResponseParams `json:"Response"`
}

func (r *DescribeTimeShiftStreamListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimeShiftStreamListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopClientIpSumInfoListRequestParams struct {
	// 起始时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 时间跨度在[0,4小时]，支持最近1天数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 播放域名，默认为不填，表示求总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 页号，范围是[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，范围是[1,1000]，默认值是20。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序指标，可选值包括TotalRequest（默认值），FailedRequest,TotalFlux。
	OrderParam *string `json:"OrderParam,omitnil,omitempty" name:"OrderParam"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 输出字段使用的语言，可选值：Chinese（默认值），English；目前国家，省份和运营商支持多语言。
	OutLanguage *string `json:"OutLanguage,omitnil,omitempty" name:"OutLanguage"`
}

type DescribeTopClientIpSumInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#:~:text=I-,ISO,-%E6%97%A5%E6%9C%9F%E6%A0%BC%E5%BC%8F)。
	// 时间跨度在[0,4小时]，支持最近1天数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 播放域名，默认为不填，表示求总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 页号，范围是[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，范围是[1,1000]，默认值是20。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序指标，可选值包括TotalRequest（默认值），FailedRequest,TotalFlux。
	OrderParam *string `json:"OrderParam,omitnil,omitempty" name:"OrderParam"`

	// 地域，可选值：Mainland，Oversea，China，Foreign，Global（默认值）；如果为空，查询总的数据；如果为“Mainland”，查询中国大陆的数据；如果为“Oversea”，则查询中国大陆以外的数据；如果为China，查询中国的数据（包括港澳台）；如果为Foreign，查询国外的数据（不包括港澳台）。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 输出字段使用的语言，可选值：Chinese（默认值），English；目前国家，省份和运营商支持多语言。
	OutLanguage *string `json:"OutLanguage,omitnil,omitempty" name:"OutLanguage"`
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

// Predefined struct for user
type DescribeTopClientIpSumInfoListResponseParams struct {
	// 页号，范围是[1,1000]，默认值是1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，范围是[1,1000]，默认值是20。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序指标，可选值包括”TotalRequest”，”FailedRequest”,“TotalFlux”。
	OrderParam *string `json:"OrderParam,omitnil,omitempty" name:"OrderParam"`

	// 记录总数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 记录总页数。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 数据内容。
	DataInfoList []*ClientIpPlaySumInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopClientIpSumInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopClientIpSumInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTranscodeTaskNumRequestParams struct {
	// 起始时间，格式：yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，格式：yyyy-mm-dd HH:MM:SS。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 推流域名列表，不填表示总体数据。
	PushDomains []*string `json:"PushDomains,omitnil,omitempty" name:"PushDomains"`
}

type DescribeTranscodeTaskNumRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间，格式：yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，格式：yyyy-mm-dd HH:MM:SS。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 推流域名列表，不填表示总体数据。
	PushDomains []*string `json:"PushDomains,omitnil,omitempty" name:"PushDomains"`
}

func (r *DescribeTranscodeTaskNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeTaskNumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PushDomains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTranscodeTaskNumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeTaskNumResponseParams struct {
	// 任务数列表。
	DataInfoList []*TranscodeTaskNum `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTranscodeTaskNumResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTranscodeTaskNumResponseParams `json:"Response"`
}

func (r *DescribeTranscodeTaskNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeTaskNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUploadStreamNumsRequestParams struct {
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 起始和结束时间跨度不支持超过31天。支持最近31天的数据查询
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 直播域名，若不填，表示总体数据。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 数据粒度，支持如下粒度：
	// 5：5分钟粒度，（跨度不支持超过1天），
	// 1440：天粒度（跨度不支持超过一个月）。
	// 默认值：5。
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeUploadStreamNumsRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，接口查询支持两种时间格式：
	// 1）YYYY-MM-DDThh:mm:ssZ：UTC时间格式，详见IOS日期格式说明文档: https://cloud.tencent.com/document/product/266/11732#I
	// 2）YYYY-MM-DD hh:mm:ss：使用此格式时，默认代表北京时间。
	// 起始和结束时间跨度不支持超过31天。支持最近31天的数据查询
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 直播域名，若不填，表示总体数据。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 数据粒度，支持如下粒度：
	// 5：5分钟粒度，（跨度不支持超过1天），
	// 1440：天粒度（跨度不支持超过一个月）。
	// 默认值：5。
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`
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

// Predefined struct for user
type DescribeUploadStreamNumsResponseParams struct {
	// 明细数据信息
	DataInfoList []*ConcurrentRecordStreamNum `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUploadStreamNumsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUploadStreamNumsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeVisitTopSumInfoListRequestParams struct {
	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS
	// 时间跨度在(0,4小时]，支持最近1天数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 峰值指标，可选值包括”Domain”，”StreamId”。
	TopIndex *string `json:"TopIndex,omitnil,omitempty" name:"TopIndex"`

	// 播放域名，默认为不填，表示求总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 页号，
	// 范围是[1,1000]，
	// 默认值是1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，范围是[1,1000]，
	// 默认值是20。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序指标，可选值包括” AvgFluxPerSecond”，”TotalRequest”（默认）,“TotalFlux”。
	OrderParam *string `json:"OrderParam,omitnil,omitempty" name:"OrderParam"`
}

type DescribeVisitTopSumInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS
	// 时间跨度在(0,4小时]，支持最近1天数据查询。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 峰值指标，可选值包括”Domain”，”StreamId”。
	TopIndex *string `json:"TopIndex,omitnil,omitempty" name:"TopIndex"`

	// 播放域名，默认为不填，表示求总体数据。
	PlayDomains []*string `json:"PlayDomains,omitnil,omitempty" name:"PlayDomains"`

	// 页号，
	// 范围是[1,1000]，
	// 默认值是1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，范围是[1,1000]，
	// 默认值是20。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序指标，可选值包括” AvgFluxPerSecond”，”TotalRequest”（默认）,“TotalFlux”。
	OrderParam *string `json:"OrderParam,omitnil,omitempty" name:"OrderParam"`
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

// Predefined struct for user
type DescribeVisitTopSumInfoListResponseParams struct {
	// 页号，
	// 范围是[1,1000]，
	// 默认值是1。
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 每页个数，范围是[1,1000]，
	// 默认值是20。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 峰值指标，可选值包括”Domain”，”StreamId”。
	TopIndex *string `json:"TopIndex,omitnil,omitempty" name:"TopIndex"`

	// 排序指标，可选值包括” AvgFluxPerSecond”(按每秒平均流量排序)，”TotalRequest”（默认，按总请求数排序）,“TotalFlux”（按总流量排序）。
	OrderParam *string `json:"OrderParam,omitnil,omitempty" name:"OrderParam"`

	// 记录总数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 记录总页数。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 数据内容。
	DataInfoList []*PlaySumStatInfo `json:"DataInfoList,omitnil,omitempty" name:"DataInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVisitTopSumInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVisitTopSumInfoListResponseParams `json:"Response"`
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

type DiagnoseResult struct {
	// 断流信息
	StreamBrokenResults []*string `json:"StreamBrokenResults,omitnil,omitempty" name:"StreamBrokenResults"`

	// 低帧率信息
	LowFrameRateResults []*string `json:"LowFrameRateResults,omitnil,omitempty" name:"LowFrameRateResults"`

	// 流格式诊断信息
	StreamFormatResults []*string `json:"StreamFormatResults,omitnil,omitempty" name:"StreamFormatResults"`
}

type DomainCertInfo struct {
	// 证书Id。
	CertId *int64 `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书名称。
	CertName *string `json:"CertName,omitnil,omitempty" name:"CertName"`

	// 描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间，UTC格式。
	// 注：此字段为北京时间（UTC+8时区）。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 证书内容。
	HttpsCrt *string `json:"HttpsCrt,omitnil,omitempty" name:"HttpsCrt"`

	// 证书类型。
	// 0：用户添加证书，
	// 1：腾讯云托管证书。
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 证书过期时间，UTC格式。
	// 注：此字段为北京时间（UTC+8时区）。
	CertExpireTime *string `json:"CertExpireTime,omitnil,omitempty" name:"CertExpireTime"`

	// 使用此证书的域名名称。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 证书状态。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 证书本身标识的域名列表。
	// 比如: ["*.x.com"]
	CertDomains []*string `json:"CertDomains,omitnil,omitempty" name:"CertDomains"`

	// 腾讯云ssl的证书Id
	CloudCertId *string `json:"CloudCertId,omitnil,omitempty" name:"CloudCertId"`
}

type DomainDetailInfo struct {
	// 国内还是国外:
	// Mainland: 表示国内数据。
	// Oversea: 表示国外数据。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`

	// 带宽，单位: Mbps。
	Bandwidth *float64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 流量，单位: MB。
	Flux *float64 `json:"Flux,omitnil,omitempty" name:"Flux"`

	// 人数。
	Online *uint64 `json:"Online,omitnil,omitempty" name:"Online"`

	// 请求数。
	Request *uint64 `json:"Request,omitnil,omitempty" name:"Request"`
}

type DomainInfo struct {
	// 直播域名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 域名类型:
	// 0: 推流。
	// 1: 播放。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 域名状态:
	// 0: 停用。
	// 1: 启用。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 添加时间。
	// 注：此字段为北京时间（UTC+8时区）。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否有 CName 到固定规则域名:
	// 0: 否。
	// 1: 是。
	BCName *uint64 `json:"BCName,omitnil,omitempty" name:"BCName"`

	// cname 对应的域名。
	TargetDomain *string `json:"TargetDomain,omitnil,omitempty" name:"TargetDomain"`

	// 播放区域，只在 Type=1 时该参数有意义。
	// 1: 国内。
	// 2: 全球。
	// 3: 海外。
	PlayType *int64 `json:"PlayType,omitnil,omitempty" name:"PlayType"`

	// 是否慢直播:
	// 0: 普通直播。
	// 1: 慢直播。
	IsDelayLive *int64 `json:"IsDelayLive,omitnil,omitempty" name:"IsDelayLive"`

	// 当前客户使用的 cname 信息。
	CurrentCName *string `json:"CurrentCName,omitnil,omitempty" name:"CurrentCName"`

	// 失效参数，可忽略。
	RentTag *int64 `json:"RentTag,omitnil,omitempty" name:"RentTag"`

	// 失效参数，可忽略。
	// 注：此字段为北京时间（UTC+8时区）。
	RentExpireTime *string `json:"RentExpireTime,omitnil,omitempty" name:"RentExpireTime"`

	// 0: 标准直播。
	// 1: 小程序直播。
	IsMiniProgramLive *int64 `json:"IsMiniProgramLive,omitnil,omitempty" name:"IsMiniProgramLive"`
}

type DomainInfoList struct {
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 明细信息。
	DetailInfoList []*DomainDetailInfo `json:"DetailInfoList,omitnil,omitempty" name:"DetailInfoList"`
}

// Predefined struct for user
type DropLiveStreamRequestParams struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`
}

type DropLiveStreamRequest struct {
	*tchttp.BaseRequest
	
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`
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

// Predefined struct for user
type DropLiveStreamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DropLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *DropLiveStreamResponseParams `json:"Response"`
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

// Predefined struct for user
type EnableLiveDomainRequestParams struct {
	// 待启用的直播域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
}

type EnableLiveDomainRequest struct {
	*tchttp.BaseRequest
	
	// 待启用的直播域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
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

// Predefined struct for user
type EnableLiveDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *EnableLiveDomainResponseParams `json:"Response"`
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

// Predefined struct for user
type EnableOptimalSwitchingRequestParams struct {
	// 针对该流 ID 启用择优调度。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 启用开关，默认为启用。
	// 0 - 禁用。
	// 1 - 启用。
	EnableSwitch *int64 `json:"EnableSwitch,omitnil,omitempty" name:"EnableSwitch"`

	// 要启用自动择优的流所属的域名分组名称。
	HostGroupName *string `json:"HostGroupName,omitnil,omitempty" name:"HostGroupName"`
}

type EnableOptimalSwitchingRequest struct {
	*tchttp.BaseRequest
	
	// 针对该流 ID 启用择优调度。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 启用开关，默认为启用。
	// 0 - 禁用。
	// 1 - 启用。
	EnableSwitch *int64 `json:"EnableSwitch,omitnil,omitempty" name:"EnableSwitch"`

	// 要启用自动择优的流所属的域名分组名称。
	HostGroupName *string `json:"HostGroupName,omitnil,omitempty" name:"HostGroupName"`
}

func (r *EnableOptimalSwitchingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableOptimalSwitchingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StreamName")
	delete(f, "EnableSwitch")
	delete(f, "HostGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableOptimalSwitchingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableOptimalSwitchingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableOptimalSwitchingResponse struct {
	*tchttp.BaseResponse
	Response *EnableOptimalSwitchingResponseParams `json:"Response"`
}

func (r *EnableOptimalSwitchingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableOptimalSwitchingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlvSpecialParam struct {
	// 是否开启边录边传，仅flv格式有效。
	UploadInRecording *bool `json:"UploadInRecording,omitnil,omitempty" name:"UploadInRecording"`
}

// Predefined struct for user
type ForbidLiveDomainRequestParams struct {
	// 待停用的直播域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
}

type ForbidLiveDomainRequest struct {
	*tchttp.BaseRequest
	
	// 待停用的直播域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
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

// Predefined struct for user
type ForbidLiveDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ForbidLiveDomainResponse struct {
	*tchttp.BaseResponse
	Response *ForbidLiveDomainResponseParams `json:"Response"`
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

// Predefined struct for user
type ForbidLiveStreamRequestParams struct {
	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 恢复流的时间。UTC 格式，例如：2018-11-29T19:00:00Z。
	// 注意：
	// 1. 默认禁推7天，且最长支持禁推90天。
	// 2. 北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#I)。
	ResumeTime *string `json:"ResumeTime,omitnil,omitempty" name:"ResumeTime"`

	// 禁推原因。
	// 注明：请务必填写禁推原因，防止误操作。
	// 长度限制：2048字节。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type ForbidLiveStreamRequest struct {
	*tchttp.BaseRequest
	
	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 恢复流的时间。UTC 格式，例如：2018-11-29T19:00:00Z。
	// 注意：
	// 1. 默认禁推7天，且最长支持禁推90天。
	// 2. 北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#I)。
	ResumeTime *string `json:"ResumeTime,omitnil,omitempty" name:"ResumeTime"`

	// 禁推原因。
	// 注明：请务必填写禁推原因，防止误操作。
	// 长度限制：2048字节。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
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

// Predefined struct for user
type ForbidLiveStreamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ForbidLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *ForbidLiveStreamResponseParams `json:"Response"`
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
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 创建时间。
	// 注：此字段为北京时间（UTC+8时区）。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 禁推过期时间。
	// 注：此字段为北京时间（UTC+8时区）。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 推流路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
}

type GroupProIspDataInfo struct {
	// 省份。
	ProvinceName *string `json:"ProvinceName,omitnil,omitempty" name:"ProvinceName"`

	// 运营商。
	IspName *string `json:"IspName,omitnil,omitempty" name:"IspName"`

	// 分钟维度的明细数据。
	DetailInfoList []*CdnPlayStatData `json:"DetailInfoList,omitnil,omitempty" name:"DetailInfoList"`
}

type HlsSpecialParam struct {
	// HLS续流超时时间。
	// 取值范围[0，1800]。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowContinueDuration *uint64 `json:"FlowContinueDuration,omitnil,omitempty" name:"FlowContinueDuration"`
}

type HttpCodeInfo struct {
	// HTTP协议返回码。
	// 例："2xx", "3xx", "4xx", "5xx"。
	HttpCode *string `json:"HttpCode,omitnil,omitempty" name:"HttpCode"`

	// 统计信息，对于无数据的时间点，会补0。
	ValueList []*HttpCodeValue `json:"ValueList,omitnil,omitempty" name:"ValueList"`
}

type HttpCodeValue struct {
	// 时间，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 次数。
	Numbers *uint64 `json:"Numbers,omitnil,omitempty" name:"Numbers"`

	// 占比。
	Percentage *float64 `json:"Percentage,omitnil,omitempty" name:"Percentage"`
}

type HttpStatusData struct {
	// 数据时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 播放状态码详细信息。
	HttpStatusInfoList []*HttpStatusInfo `json:"HttpStatusInfoList,omitnil,omitempty" name:"HttpStatusInfoList"`
}

type HttpStatusInfo struct {
	// 播放HTTP状态码。
	HttpStatus *string `json:"HttpStatus,omitnil,omitempty" name:"HttpStatus"`

	// 个数。
	Num *uint64 `json:"Num,omitnil,omitempty" name:"Num"`
}

type LiveCertDomainInfo struct {
	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 是否启用域名的https规则。
	// 1：启用
	// 0：禁用
	// -1：保持不变
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type LiveDomainCertBindings struct {
	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 证书备注。与CertName同义。
	CertificateAlias *string `json:"CertificateAlias,omitnil,omitempty" name:"CertificateAlias"`

	// 证书类型。
	// 0：自有证书
	// 1：腾讯云ssl托管证书
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// https状态。
	// 1：已开启。
	// 0：已关闭。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 证书过期时间。
	// 注：此字段为北京时间（UTC+8时区）。
	CertExpireTime *string `json:"CertExpireTime,omitnil,omitempty" name:"CertExpireTime"`

	// 证书Id。
	CertId *int64 `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 腾讯云ssl的证书Id。
	CloudCertId *string `json:"CloudCertId,omitnil,omitempty" name:"CloudCertId"`

	// 规则最后更新时间。
	// 注：此字段为北京时间（UTC+8时区）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type LiveEnhanceInfo struct {
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 时间。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 计费时长，单位分钟。
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 帧率。
	Fps *string `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 分辨率。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 增强服务类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type LivePackageInfo struct {
	// 包 ID。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 总量。
	// 注意：当为流量包时单位为字节。
	// 当为转码包时单位为分钟。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 使用量。
	// 注意：当为流量包时单位为字节。
	// 当为转码包时单位为分钟。
	// 当为连麦包时单位为小时。
	Used *int64 `json:"Used,omitnil,omitempty" name:"Used"`

	// 剩余量。
	// 注意：当为流量包时单位为字节。
	// 当为转码包时单位为分钟。
	// 当为连麦包时单位为小时。
	Left *int64 `json:"Left,omitnil,omitempty" name:"Left"`

	// 购买时间。
	// 注：此字段为北京时间（UTC+8时区）。
	BuyTime *string `json:"BuyTime,omitnil,omitempty" name:"BuyTime"`

	// 过期时间。
	// 注：此字段为北京时间（UTC+8时区）。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 包类型，可选值:
	// 0: 流量包。
	// 1: 普通转码包。
	// 2: 极速高清包。
	// 3: 连麦包。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 包状态，可选值:
	// 0: 未使用。
	// 1: 使用中。
	// 2: 已过期。
	// 3: 已冻结。
	// 4: 已耗尽。
	// 5: 已退款
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否自动续购。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WillRenew *int64 `json:"WillRenew,omitnil,omitempty" name:"WillRenew"`

	// 续购状态。
	// 1 ：续购成功。
	// 0 ：尚未续购。
	// <0  : 续购失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewalResult *int64 `json:"RenewalResult,omitnil,omitempty" name:"RenewalResult"`
}

type LiveStreamMonitorInfo struct {
	// 监播任务ID。
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 监播任务名称。128字节以内。
	MonitorName *string `json:"MonitorName,omitnil,omitempty" name:"MonitorName"`

	// 监播任务输出信息。
	OutputInfo *LiveStreamMonitorOutputInfo `json:"OutputInfo,omitnil,omitempty" name:"OutputInfo"`

	// 待监播的输入流信息。
	InputList []*LiveStreamMonitorInputInfo `json:"InputList,omitnil,omitempty" name:"InputList"`

	// 监播任务状态。
	// 0： 代表空闲
	// 1： 代表监播中。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 上一次的启动时间，unix时间戳。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 上一次的停止时间，unix时间戳。
	StopTime *uint64 `json:"StopTime,omitnil,omitempty" name:"StopTime"`

	// 监播任务创建时间，unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 监播任务更新时间，unix时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 监播事件通知策略。
	NotifyPolicy *LiveStreamMonitorNotifyPolicy `json:"NotifyPolicy,omitnil,omitempty" name:"NotifyPolicy"`

	// 输出音频的输入Index列表。
	AudibleInputIndexList []*uint64 `json:"AudibleInputIndexList,omitnil,omitempty" name:"AudibleInputIndexList"`

	// 开启智能语音识别的输入Index列表
	AiAsrInputIndexList []*uint64 `json:"AiAsrInputIndexList,omitnil,omitempty" name:"AiAsrInputIndexList"`

	// 是否开启断流检测
	CheckStreamBroken *uint64 `json:"CheckStreamBroken,omitnil,omitempty" name:"CheckStreamBroken"`

	// 是否开启低帧率检测
	CheckStreamLowFrameRate *uint64 `json:"CheckStreamLowFrameRate,omitnil,omitempty" name:"CheckStreamLowFrameRate"`

	// 智能语音识别语种：
	// 0 关闭 1 中文 2 英文 3日文 4 韩文
	AsrLanguage *uint64 `json:"AsrLanguage,omitnil,omitempty" name:"AsrLanguage"`

	// 智能文字识别语种：
	// 0 关闭 1 中、英文
	OcrLanguage *uint64 `json:"OcrLanguage,omitnil,omitempty" name:"OcrLanguage"`

	// 开启智能文字识别的输入Index列表
	AiOcrInputIndexList []*uint64 `json:"AiOcrInputIndexList,omitnil,omitempty" name:"AiOcrInputIndexList"`

	// 是否存储监播事件到监播报告，以及是否允许查询监播报告
	AllowMonitorReport *uint64 `json:"AllowMonitorReport,omitnil,omitempty" name:"AllowMonitorReport"`

	// 是否开启格式诊断
	AiFormatDiagnose *uint64 `json:"AiFormatDiagnose,omitnil,omitempty" name:"AiFormatDiagnose"`

	// 是否开启内容质检。
	AiQualityControl *uint64 `json:"AiQualityControl,omitnil,omitempty" name:"AiQualityControl"`

	// 导播台监播对应的导播台场次id
	CasterId *string `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 拉流转推监播对应的拉流转推任务id
	PullPushTaskId *string `json:"PullPushTaskId,omitnil,omitempty" name:"PullPushTaskId"`
}

type LiveStreamMonitorInputInfo struct {
	// 待监播的输入流名称。256字节以内，只允许包含字母、数字、‘-’，‘_’，'.'字符。
	InputStreamName *string `json:"InputStreamName,omitnil,omitempty" name:"InputStreamName"`

	// 待监播的输入流推流域名。128字节以内，只允许填处于启用状态的推流域名。
	InputDomain *string `json:"InputDomain,omitnil,omitempty" name:"InputDomain"`

	// 待监播的输入流推流路径。32字节以内，只允许包含字母、数字、‘-’，‘_’，'.'字符。
	InputApp *string `json:"InputApp,omitnil,omitempty" name:"InputApp"`

	// 待监播的输入流推流url。一般场景下，无需该参数。
	InputUrl *string `json:"InputUrl,omitnil,omitempty" name:"InputUrl"`

	// 描述。256字节以内。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 导播台输入源索引（10000 pvw， 10001 pgm， 其余代表输入下标）
	CasterInputIndex *uint64 `json:"CasterInputIndex,omitnil,omitempty" name:"CasterInputIndex"`

	// 该输入源是否正在监播
	NeedMonitor *bool `json:"NeedMonitor,omitnil,omitempty" name:"NeedMonitor"`

	// 导播台pvw pgm的cdn流id
	CdnStreamId *string `json:"CdnStreamId,omitnil,omitempty" name:"CdnStreamId"`
}

type LiveStreamMonitorNotifyPolicy struct {
	// 通知策略类型：范围[0,1]
	// 0:代表不使用任何通知策略
	// 1:代表使用全局回调策略，所有事件通知到CallbackUrl。
	NotifyPolicyType *uint64 `json:"NotifyPolicyType,omitnil,omitempty" name:"NotifyPolicyType"`

	// 回调URL：长度[0,512]
	// 只支持http和https类型的url。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`
}

type LiveStreamMonitorOutputInfo struct {
	// 监播任务输出流宽度像素。范围[1,1920]。建议至少大于100像素。
	OutputStreamWidth *uint64 `json:"OutputStreamWidth,omitnil,omitempty" name:"OutputStreamWidth"`

	// 监播任务输出流长度像素。范围[1,1080]，建议至少大于100像素。
	OutputStreamHeight *uint64 `json:"OutputStreamHeight,omitnil,omitempty" name:"OutputStreamHeight"`

	// 监播任务输出流名称。
	// 不填时，系统会自动生成。
	// 256字节以内，只允许包含字母、数字、‘-’，‘_’，'.'字符。
	OutputStreamName *string `json:"OutputStreamName,omitnil,omitempty" name:"OutputStreamName"`

	// 监播任务播放域名。128字节以内，只允许填处于启用状态的播放域名。
	OutputDomain *string `json:"OutputDomain,omitnil,omitempty" name:"OutputDomain"`

	// 监播任务播放路径。32字节以内，只允许包含字母、数字、‘-’，‘_’，'.'字符。
	OutputApp *string `json:"OutputApp,omitnil,omitempty" name:"OutputApp"`
}

type LogInfo struct {
	// 日志名称。
	LogName *string `json:"LogName,omitnil,omitempty" name:"LogName"`

	// 日志 URL。
	LogUrl *string `json:"LogUrl,omitnil,omitempty" name:"LogUrl"`

	// 日志生成时间。
	// 注：此字段为北京时间（UTC+8时区）。
	LogTime *string `json:"LogTime,omitnil,omitempty" name:"LogTime"`

	// 文件大小。
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`
}

type MPSResult struct {
	// 智能语音识别结果。
	AiAsrResults []*string `json:"AiAsrResults,omitnil,omitempty" name:"AiAsrResults"`

	// 智能文字识别结果。
	AiOcrResults []*string `json:"AiOcrResults,omitnil,omitempty" name:"AiOcrResults"`

	// 内容质检结果。
	StreamQuaCtrlResults []*string `json:"StreamQuaCtrlResults,omitnil,omitempty" name:"StreamQuaCtrlResults"`
}

type MixPortraitSegmentParams struct {
	// 抠图背景颜色，
	// 常用的颜色有：
	// 红色：0xcc0033。
	// 黄色：0xcc9900。
	// 绿色：0xcccc33。
	// 蓝色：0x99CCFF。
	// 黑色：0x000000。
	// 白色：0xFFFFFF。
	// 灰色：0x999999。
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`
}

// Predefined struct for user
type ModifyCasterInputInfoRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 修改的导播台输入源信息
	InputInfo *CasterInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`
}

type ModifyCasterInputInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 修改的导播台输入源信息
	InputInfo *CasterInputInfo `json:"InputInfo,omitnil,omitempty" name:"InputInfo"`
}

func (r *ModifyCasterInputInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCasterInputInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "InputInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCasterInputInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCasterInputInfoResponseParams struct {
	// 修改输入源后的预览地址。
	// 注：该地址仅作为预览使用，不可分发。
	InputPlayUrl *string `json:"InputPlayUrl,omitnil,omitempty" name:"InputPlayUrl"`

	// 修改后的输入源webrtc预览地址。
	// 该地址需配合腾讯云快直播播放SDK使用。
	// 注：该地址仅做预览使用，不可分发。
	InputWebRTCPlayUrl *string `json:"InputWebRTCPlayUrl,omitnil,omitempty" name:"InputWebRTCPlayUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCasterInputInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCasterInputInfoResponseParams `json:"Response"`
}

func (r *ModifyCasterInputInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCasterInputInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCasterLayoutInfoRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台布局参数信息。
	LayoutInfo *CasterLayoutInfo `json:"LayoutInfo,omitnil,omitempty" name:"LayoutInfo"`
}

type ModifyCasterLayoutInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台布局参数信息。
	LayoutInfo *CasterLayoutInfo `json:"LayoutInfo,omitnil,omitempty" name:"LayoutInfo"`
}

func (r *ModifyCasterLayoutInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCasterLayoutInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "LayoutInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCasterLayoutInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCasterLayoutInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCasterLayoutInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCasterLayoutInfoResponseParams `json:"Response"`
}

func (r *ModifyCasterLayoutInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCasterLayoutInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCasterMarkPicInfoRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 新的水印信息。
	MarkPicInfo *CasterMarkPicInfo `json:"MarkPicInfo,omitnil,omitempty" name:"MarkPicInfo"`
}

type ModifyCasterMarkPicInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 新的水印信息。
	MarkPicInfo *CasterMarkPicInfo `json:"MarkPicInfo,omitnil,omitempty" name:"MarkPicInfo"`
}

func (r *ModifyCasterMarkPicInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCasterMarkPicInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "MarkPicInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCasterMarkPicInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCasterMarkPicInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCasterMarkPicInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCasterMarkPicInfoResponseParams `json:"Response"`
}

func (r *ModifyCasterMarkPicInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCasterMarkPicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCasterMarkWordInfoRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 修改的文本配置。
	MarkWordInfo *CasterMarkWordInfo `json:"MarkWordInfo,omitnil,omitempty" name:"MarkWordInfo"`
}

type ModifyCasterMarkWordInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 修改的文本配置。
	MarkWordInfo *CasterMarkWordInfo `json:"MarkWordInfo,omitnil,omitempty" name:"MarkWordInfo"`
}

func (r *ModifyCasterMarkWordInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCasterMarkWordInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "MarkWordInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCasterMarkWordInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCasterMarkWordInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCasterMarkWordInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCasterMarkWordInfoResponseParams `json:"Response"`
}

func (r *ModifyCasterMarkWordInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCasterMarkWordInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCasterOutputInfoRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台推流参数信息。
	OutputInfo *CasterOutputInfo `json:"OutputInfo,omitnil,omitempty" name:"OutputInfo"`
}

type ModifyCasterOutputInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台推流参数信息。
	OutputInfo *CasterOutputInfo `json:"OutputInfo,omitnil,omitempty" name:"OutputInfo"`
}

func (r *ModifyCasterOutputInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCasterOutputInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "OutputInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCasterOutputInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCasterOutputInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCasterOutputInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCasterOutputInfoResponseParams `json:"Response"`
}

func (r *ModifyCasterOutputInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCasterOutputInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCasterRequestParams struct {
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台名称
	CasterName *string `json:"CasterName,omitnil,omitempty" name:"CasterName"`

	// 导播台的描述
	// 最大允许长度256
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 录制模板id。
	// 默认为0。
	// 当使用直播录制功能时，可将使用的录制模版填入。
	// 该接口仅保存字段，不涉及任何录制功能。
	RecordTemplateId *uint64 `json:"RecordTemplateId,omitnil,omitempty" name:"RecordTemplateId"`

	// 录制状态，当调用录制接口后，可通过该字段保存录制状态。
	// 0：未录制 
	// 1：录制中
	// 该接口仅保存字段，不涉及任何录制处理。
	RecordStatus *uint64 `json:"RecordStatus,omitnil,omitempty" name:"RecordStatus"`

	// 导播台的过期时间戳。值为-1或unix时间戳。
	// 默认值为-1。 
	// 当值为-1时，表示该导播台永不过期。 
	// 当值为正常unix时间戳时，导播台将在该时间过期。
	// 导播台过期后，预监与主监画面将自动停止，转推自动停止。 
	// 点播、直播url将停止转拉，推流url需自行停止推流。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 导播台延时播放时间，单位为秒。 
	// 默认为0，最大支持300秒
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 导播台转场类型。 
	// 默认为空。 
	// 允许使用通过DescribeCasterTransitionTypes接口中查询到的转场类型。
	TransitionType *string `json:"TransitionType,omitnil,omitempty" name:"TransitionType"`

	// 导播台主监输出的宽度，单位为像素。 
	// 默认为1280，最大允许4096。
	PgmWidth *uint64 `json:"PgmWidth,omitnil,omitempty" name:"PgmWidth"`

	// 导播台主监输出的高度，单位为像素。 
	// 默认为720，最大允许2160。
	PgmHeight *uint64 `json:"PgmHeight,omitnil,omitempty" name:"PgmHeight"`

	// 导播台主监输出的帧率。 
	// 默认为0，表示随源输出。 最大支持60。
	PgmFps *uint64 `json:"PgmFps,omitnil,omitempty" name:"PgmFps"`

	// 导播台主监输出的码率，单位为kbps。 
	// 默认为0，表示随源的码率输出。 
	// 最大允许10000kbps。
	PgmBitRate *uint64 `json:"PgmBitRate,omitnil,omitempty" name:"PgmBitRate"`

	// 导播台的计费类型。 
	// 0 通用型 
	// 1 播单型。 
	// 注： 本参数暂无作用。
	FeeType *int64 `json:"FeeType,omitnil,omitempty" name:"FeeType"`

	// 录制接口返回的taskid
	// 注：该接口只做字段保存，不涉及录制操作。
	RecordTaskId *string `json:"RecordTaskId,omitnil,omitempty" name:"RecordTaskId"`

	// 导播台主监输出的音频码率，单位为kbps。 
	// 可选项：[0, 128, 192, 256] 
	// 默认值为0，表示随源的音频码率输出。
	PgmAudioBitRate *uint64 `json:"PgmAudioBitRate,omitnil,omitempty" name:"PgmAudioBitRate"`
}

type ModifyCasterRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`

	// 导播台名称
	CasterName *string `json:"CasterName,omitnil,omitempty" name:"CasterName"`

	// 导播台的描述
	// 最大允许长度256
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 录制模板id。
	// 默认为0。
	// 当使用直播录制功能时，可将使用的录制模版填入。
	// 该接口仅保存字段，不涉及任何录制功能。
	RecordTemplateId *uint64 `json:"RecordTemplateId,omitnil,omitempty" name:"RecordTemplateId"`

	// 录制状态，当调用录制接口后，可通过该字段保存录制状态。
	// 0：未录制 
	// 1：录制中
	// 该接口仅保存字段，不涉及任何录制处理。
	RecordStatus *uint64 `json:"RecordStatus,omitnil,omitempty" name:"RecordStatus"`

	// 导播台的过期时间戳。值为-1或unix时间戳。
	// 默认值为-1。 
	// 当值为-1时，表示该导播台永不过期。 
	// 当值为正常unix时间戳时，导播台将在该时间过期。
	// 导播台过期后，预监与主监画面将自动停止，转推自动停止。 
	// 点播、直播url将停止转拉，推流url需自行停止推流。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 导播台延时播放时间，单位为秒。 
	// 默认为0，最大支持300秒
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 导播台转场类型。 
	// 默认为空。 
	// 允许使用通过DescribeCasterTransitionTypes接口中查询到的转场类型。
	TransitionType *string `json:"TransitionType,omitnil,omitempty" name:"TransitionType"`

	// 导播台主监输出的宽度，单位为像素。 
	// 默认为1280，最大允许4096。
	PgmWidth *uint64 `json:"PgmWidth,omitnil,omitempty" name:"PgmWidth"`

	// 导播台主监输出的高度，单位为像素。 
	// 默认为720，最大允许2160。
	PgmHeight *uint64 `json:"PgmHeight,omitnil,omitempty" name:"PgmHeight"`

	// 导播台主监输出的帧率。 
	// 默认为0，表示随源输出。 最大支持60。
	PgmFps *uint64 `json:"PgmFps,omitnil,omitempty" name:"PgmFps"`

	// 导播台主监输出的码率，单位为kbps。 
	// 默认为0，表示随源的码率输出。 
	// 最大允许10000kbps。
	PgmBitRate *uint64 `json:"PgmBitRate,omitnil,omitempty" name:"PgmBitRate"`

	// 导播台的计费类型。 
	// 0 通用型 
	// 1 播单型。 
	// 注： 本参数暂无作用。
	FeeType *int64 `json:"FeeType,omitnil,omitempty" name:"FeeType"`

	// 录制接口返回的taskid
	// 注：该接口只做字段保存，不涉及录制操作。
	RecordTaskId *string `json:"RecordTaskId,omitnil,omitempty" name:"RecordTaskId"`

	// 导播台主监输出的音频码率，单位为kbps。 
	// 可选项：[0, 128, 192, 256] 
	// 默认值为0，表示随源的音频码率输出。
	PgmAudioBitRate *uint64 `json:"PgmAudioBitRate,omitnil,omitempty" name:"PgmAudioBitRate"`
}

func (r *ModifyCasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	delete(f, "CasterName")
	delete(f, "Description")
	delete(f, "RecordTemplateId")
	delete(f, "RecordStatus")
	delete(f, "ExpireTime")
	delete(f, "DelayTime")
	delete(f, "TransitionType")
	delete(f, "PgmWidth")
	delete(f, "PgmHeight")
	delete(f, "PgmFps")
	delete(f, "PgmBitRate")
	delete(f, "FeeType")
	delete(f, "RecordTaskId")
	delete(f, "PgmAudioBitRate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCasterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCasterResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCasterResponseParams `json:"Response"`
}

func (r *ModifyCasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLiveCallbackTemplateRequestParams struct {
	// DescribeLiveCallbackTemplates接口返回的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 开播回调 URL。
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitnil,omitempty" name:"StreamBeginNotifyUrl"`

	// 断流回调 URL。
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitnil,omitempty" name:"StreamEndNotifyUrl"`

	// 录制文件回调 URL。
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitnil,omitempty" name:"RecordNotifyUrl"`

	// 录制状态回调 URL 。
	RecordStatusNotifyUrl *string `json:"RecordStatusNotifyUrl,omitnil,omitempty" name:"RecordStatusNotifyUrl"`

	// 截图回调 URL。
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitnil,omitempty" name:"SnapshotNotifyUrl"`

	// 鉴黄回调 URL。
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitnil,omitempty" name:"PornCensorshipNotifyUrl"`

	// 回调 Key，回调 URL 公用，回调签名详见事件消息通知文档。
	// [事件消息通知](/document/product/267/32744)。
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 推流异常回调 URL。
	PushExceptionNotifyUrl *string `json:"PushExceptionNotifyUrl,omitnil,omitempty" name:"PushExceptionNotifyUrl"`

	// 音频审核回调 URL。
	AudioAuditNotifyUrl *string `json:"AudioAuditNotifyUrl,omitnil,omitempty" name:"AudioAuditNotifyUrl"`

	// 录制异常回调 URL。	
	RecordExceptionNotifyUrl *string `json:"RecordExceptionNotifyUrl,omitnil,omitempty" name:"RecordExceptionNotifyUrl"`

	// 录制异常回调级别，可选择：error、warning、info。
	RecordExceptionLevels []*string `json:"RecordExceptionLevels,omitnil,omitempty" name:"RecordExceptionLevels"`
}

type ModifyLiveCallbackTemplateRequest struct {
	*tchttp.BaseRequest
	
	// DescribeLiveCallbackTemplates接口返回的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 开播回调 URL。
	StreamBeginNotifyUrl *string `json:"StreamBeginNotifyUrl,omitnil,omitempty" name:"StreamBeginNotifyUrl"`

	// 断流回调 URL。
	StreamEndNotifyUrl *string `json:"StreamEndNotifyUrl,omitnil,omitempty" name:"StreamEndNotifyUrl"`

	// 录制文件回调 URL。
	RecordNotifyUrl *string `json:"RecordNotifyUrl,omitnil,omitempty" name:"RecordNotifyUrl"`

	// 录制状态回调 URL 。
	RecordStatusNotifyUrl *string `json:"RecordStatusNotifyUrl,omitnil,omitempty" name:"RecordStatusNotifyUrl"`

	// 截图回调 URL。
	SnapshotNotifyUrl *string `json:"SnapshotNotifyUrl,omitnil,omitempty" name:"SnapshotNotifyUrl"`

	// 鉴黄回调 URL。
	PornCensorshipNotifyUrl *string `json:"PornCensorshipNotifyUrl,omitnil,omitempty" name:"PornCensorshipNotifyUrl"`

	// 回调 Key，回调 URL 公用，回调签名详见事件消息通知文档。
	// [事件消息通知](/document/product/267/32744)。
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 推流异常回调 URL。
	PushExceptionNotifyUrl *string `json:"PushExceptionNotifyUrl,omitnil,omitempty" name:"PushExceptionNotifyUrl"`

	// 音频审核回调 URL。
	AudioAuditNotifyUrl *string `json:"AudioAuditNotifyUrl,omitnil,omitempty" name:"AudioAuditNotifyUrl"`

	// 录制异常回调 URL。	
	RecordExceptionNotifyUrl *string `json:"RecordExceptionNotifyUrl,omitnil,omitempty" name:"RecordExceptionNotifyUrl"`

	// 录制异常回调级别，可选择：error、warning、info。
	RecordExceptionLevels []*string `json:"RecordExceptionLevels,omitnil,omitempty" name:"RecordExceptionLevels"`
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
	delete(f, "RecordStatusNotifyUrl")
	delete(f, "SnapshotNotifyUrl")
	delete(f, "PornCensorshipNotifyUrl")
	delete(f, "CallbackKey")
	delete(f, "PushExceptionNotifyUrl")
	delete(f, "AudioAuditNotifyUrl")
	delete(f, "RecordExceptionNotifyUrl")
	delete(f, "RecordExceptionLevels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveCallbackTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLiveCallbackTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLiveCallbackTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLiveCallbackTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLiveDomainCertBindingsRequestParams struct {
	// 要绑定证书的播放域名/状态 信息列表。
	// 如果CloudCertId和证书公钥私钥对均不传，且域名列表已有绑定规则，只批量更新域名https规则的启用状态，并把未上传至腾讯云ssl的已有自有证书上传。
	DomainInfos []*LiveCertDomainInfo `json:"DomainInfos,omitnil,omitempty" name:"DomainInfos"`

	// 腾讯云ssl的证书Id。
	// 见 https://cloud.tencent.com/document/api/400/41665
	CloudCertId *string `json:"CloudCertId,omitnil,omitempty" name:"CloudCertId"`

	// 证书公钥。
	// CloudCertId和公钥私钥对二选一，若CloudCertId将会舍弃公钥和私钥参数，否则将自动先把公钥私钥对上传至ssl新建证书，并使用上传成功后返回的CloudCertId。
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// 证书私钥。
	// CloudCertId和公钥私钥对二选一，若传CloudCertId将会舍弃公钥和私钥参数，否则将自动先把公钥私钥对上传至ssl新建证书，并使用上传成功后返回的CloudCertId。
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// 上传至ssl证书中心的备注信息，只有新建证书时有效。传CloudCertId时会忽略。
	CertificateAlias *string `json:"CertificateAlias,omitnil,omitempty" name:"CertificateAlias"`
}

type ModifyLiveDomainCertBindingsRequest struct {
	*tchttp.BaseRequest
	
	// 要绑定证书的播放域名/状态 信息列表。
	// 如果CloudCertId和证书公钥私钥对均不传，且域名列表已有绑定规则，只批量更新域名https规则的启用状态，并把未上传至腾讯云ssl的已有自有证书上传。
	DomainInfos []*LiveCertDomainInfo `json:"DomainInfos,omitnil,omitempty" name:"DomainInfos"`

	// 腾讯云ssl的证书Id。
	// 见 https://cloud.tencent.com/document/api/400/41665
	CloudCertId *string `json:"CloudCertId,omitnil,omitempty" name:"CloudCertId"`

	// 证书公钥。
	// CloudCertId和公钥私钥对二选一，若CloudCertId将会舍弃公钥和私钥参数，否则将自动先把公钥私钥对上传至ssl新建证书，并使用上传成功后返回的CloudCertId。
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// 证书私钥。
	// CloudCertId和公钥私钥对二选一，若传CloudCertId将会舍弃公钥和私钥参数，否则将自动先把公钥私钥对上传至ssl新建证书，并使用上传成功后返回的CloudCertId。
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// 上传至ssl证书中心的备注信息，只有新建证书时有效。传CloudCertId时会忽略。
	CertificateAlias *string `json:"CertificateAlias,omitnil,omitempty" name:"CertificateAlias"`
}

func (r *ModifyLiveDomainCertBindingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveDomainCertBindingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainInfos")
	delete(f, "CloudCertId")
	delete(f, "CertificatePublicKey")
	delete(f, "CertificatePrivateKey")
	delete(f, "CertificateAlias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveDomainCertBindingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLiveDomainCertBindingsResponseParams struct {
	// DomainNames 入参中，与证书不匹配的域名列表，将会跳过处理。
	MismatchedDomainNames []*string `json:"MismatchedDomainNames,omitnil,omitempty" name:"MismatchedDomainNames"`

	// 操作失败的域名及错误码，错误信息，包括MismatchedDomainNames中的域名。
	Errors []*BatchDomainOperateErrors `json:"Errors,omitnil,omitempty" name:"Errors"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLiveDomainCertBindingsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLiveDomainCertBindingsResponseParams `json:"Response"`
}

func (r *ModifyLiveDomainCertBindingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveDomainCertBindingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLiveDomainRefererRequestParams struct {
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 是否开启当前域名的 Referer 黑白名单鉴权。
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 名单类型，0：黑名单，1：白名单。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否允许空 Referer，0：不允许，1：允许。
	AllowEmpty *int64 `json:"AllowEmpty,omitnil,omitempty" name:"AllowEmpty"`

	// Referer 名单列表，以;分隔。
	Rules *string `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type ModifyLiveDomainRefererRequest struct {
	*tchttp.BaseRequest
	
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 是否开启当前域名的 Referer 黑白名单鉴权。
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 名单类型，0：黑名单，1：白名单。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否允许空 Referer，0：不允许，1：允许。
	AllowEmpty *int64 `json:"AllowEmpty,omitnil,omitempty" name:"AllowEmpty"`

	// Referer 名单列表，以;分隔。
	Rules *string `json:"Rules,omitnil,omitempty" name:"Rules"`
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

// Predefined struct for user
type ModifyLiveDomainRefererResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLiveDomainRefererResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLiveDomainRefererResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLivePadTemplateRequestParams struct {
	// 模板id。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 垫片内容。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 断流等待时间。
	// 取值范围：0-30000。
	// 单位：ms。
	WaitDuration *uint64 `json:"WaitDuration,omitnil,omitempty" name:"WaitDuration"`

	// 最大垫片时长。
	// 取值范围：0 - 正无穷。
	// 单位：ms。
	MaxDuration *uint64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// 模板名称。
	// 长度上限：255字节。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	// 长度上限：1024字节。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 垫片内容类型： 1：图片，2：视频。 默认值：1。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type ModifyLivePadTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板id。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 垫片内容。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 断流等待时间。
	// 取值范围：0-30000。
	// 单位：ms。
	WaitDuration *uint64 `json:"WaitDuration,omitnil,omitempty" name:"WaitDuration"`

	// 最大垫片时长。
	// 取值范围：0 - 正无穷。
	// 单位：ms。
	MaxDuration *uint64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// 模板名称。
	// 长度上限：255字节。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	// 长度上限：1024字节。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 垫片内容类型： 1：图片，2：视频。 默认值：1。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *ModifyLivePadTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLivePadTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Url")
	delete(f, "WaitDuration")
	delete(f, "MaxDuration")
	delete(f, "TemplateName")
	delete(f, "Description")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLivePadTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLivePadTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLivePadTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLivePadTemplateResponseParams `json:"Response"`
}

func (r *ModifyLivePadTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLivePadTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLivePlayAuthKeyRequestParams struct {
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	// 不传表示不修改当前值。
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 鉴权key。
	// 不传表示不修改当前值。
	AuthKey *string `json:"AuthKey,omitnil,omitempty" name:"AuthKey"`

	// 有效时间，单位：秒。
	// 不传表示不修改当前值。
	AuthDelta *uint64 `json:"AuthDelta,omitnil,omitempty" name:"AuthDelta"`

	// 鉴权备用key。
	// 不传表示不修改当前值。
	AuthBackKey *string `json:"AuthBackKey,omitnil,omitempty" name:"AuthBackKey"`
}

type ModifyLivePlayAuthKeyRequest struct {
	*tchttp.BaseRequest
	
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	// 不传表示不修改当前值。
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 鉴权key。
	// 不传表示不修改当前值。
	AuthKey *string `json:"AuthKey,omitnil,omitempty" name:"AuthKey"`

	// 有效时间，单位：秒。
	// 不传表示不修改当前值。
	AuthDelta *uint64 `json:"AuthDelta,omitnil,omitempty" name:"AuthDelta"`

	// 鉴权备用key。
	// 不传表示不修改当前值。
	AuthBackKey *string `json:"AuthBackKey,omitnil,omitempty" name:"AuthBackKey"`
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

// Predefined struct for user
type ModifyLivePlayAuthKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLivePlayAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLivePlayAuthKeyResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLivePlayDomainRequestParams struct {
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 拉流域名类型。1-国内；2-全球；3-境外
	PlayType *int64 `json:"PlayType,omitnil,omitempty" name:"PlayType"`
}

type ModifyLivePlayDomainRequest struct {
	*tchttp.BaseRequest
	
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 拉流域名类型。1-国内；2-全球；3-境外
	PlayType *int64 `json:"PlayType,omitnil,omitempty" name:"PlayType"`
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

// Predefined struct for user
type ModifyLivePlayDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLivePlayDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLivePlayDomainResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLivePullStreamTaskRequestParams struct {
	// 任务Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作人姓名。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 拉流源url列表。
	// SourceType为直播（PullLivePushLive）只可以填1个，
	// SourceType为点播（PullVodPushLive）可以填多个，上限30个。
	SourceUrls []*string `json:"SourceUrls,omitnil,omitempty" name:"SourceUrls"`

	// 开始时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于30天。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 点播拉流转推循环次数。
	// -1：无限循环，直到任务结束。
	// 0：不循环。
	// >0：具体循环次数。次数和时间以先结束的为准。
	// 注意：拉流源为点播，该配置生效。
	VodLoopTimes *int64 `json:"VodLoopTimes,omitnil,omitempty" name:"VodLoopTimes"`

	// 点播更新SourceUrls后的播放方式：
	// ImmediateNewSource：立即从更新的拉流源开始播放；
	// ContinueBreakPoint：从上次断流url源的断点处继续，结束后再使用新的拉流源。
	// 注意：拉流源为点播，该配置生效。
	VodRefreshType *string `json:"VodRefreshType,omitnil,omitempty" name:"VodRefreshType"`

	// 任务状态：
	// enable - 启用，
	// pause - 暂停。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 选择需要回调的事件（不填则回调全部）：
	// TaskStart：任务启动回调，
	// TaskExit：任务停止回调，
	// VodSourceFileStart：从点播源文件开始拉流回调，
	// VodSourceFileFinish：从点播源文件拉流结束回调，
	// ResetTaskConfig：任务更新回调。
	CallbackEvents []*string `json:"CallbackEvents,omitnil,omitempty" name:"CallbackEvents"`

	// 自定义回调地址。
	// 相关事件会回调到该地址。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 指定播放文件索引。
	// 注意： 从1开始，不大于SourceUrls中文件个数。
	FileIndex *int64 `json:"FileIndex,omitnil,omitempty" name:"FileIndex"`

	// 指定播放文件偏移。
	// 注意：
	// 1. 单位：秒，配合FileIndex使用。
	OffsetTime *int64 `json:"OffsetTime,omitnil,omitempty" name:"OffsetTime"`

	// 指定任务 ID 修改任务。
	// 
	// 注意：该自定义任务 ID 只有在创建任务时指定了，才可在此处修改时使用。否则请使用系统返回的任务 ID。
	SpecifyTaskId *string `json:"SpecifyTaskId,omitnil,omitempty" name:"SpecifyTaskId"`

	// 目标 Url。
	// 换目标地址，会断流重推到新地址。
	ToUrl *string `json:"ToUrl,omitnil,omitempty" name:"ToUrl"`

	// 任务备注。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 备源的类型：
	// PullLivePushLive -直播，
	// PullVodPushLive -点播。
	// 注意：
	// 1. 仅当主源类型为直播源时，备源才会生效。
	// 2. 将该参数置为空，则可将任务去除备源信息。
	// 3. 主直播源拉流中断时，自动使用备源进行拉流。
	// 4. 如果备源为点播文件时，则每次轮播完点播文件就检查主源是否恢复，如果主源恢复则自动切回到主源，否则继续拉备源。
	BackupSourceType *string `json:"BackupSourceType,omitnil,omitempty" name:"BackupSourceType"`

	// 备源 URL。
	// 只允许填一个备源 URL
	BackupSourceUrl *string `json:"BackupSourceUrl,omitnil,omitempty" name:"BackupSourceUrl"`

	// 水印信息列表。
	// 注意：
	// 1. 最多支持4个不同位置的水印。
	// 2. 水印图片 URL 请使用合法外网可访问地址。
	// 3. 支持的水印图片格式：png，jpg等。
	// 4. 轮播任务修改水印后，轮播到下一个文件时新水印生效。
	// 5. 直播源任务修改水印后，水印立即生效。
	// 6. 清除水印时，需携带该水印列表参数，内容为空数组。
	// 7. 暂不支持动图水印。
	WatermarkList []*PullPushWatermarkInfo `json:"WatermarkList,omitnil,omitempty" name:"WatermarkList"`

	// 点播源是否启用本地推流模式，默认0，不启用。
	// 0 - 不启用。
	// 1 - 启用。
	// 注意：启用本地模式后，会将源列表中的 MP4 文件进行本地下载，优先使用本地已下载文件进行推流，提高点播源推流稳定性。使用本地下载文件推流时，会产生增值费用。
	VodLocalMode *int64 `json:"VodLocalMode,omitnil,omitempty" name:"VodLocalMode"`

	// 新的目标地址。传空值，则取消该地址的推流。
	// 传入新值，则替换原有地址。
	BackupToUrl *string `json:"BackupToUrl,omitnil,omitempty" name:"BackupToUrl"`

	// 点播垫片文件地址。注意：用于在主源拉不到时自动兜底到垫片文件，切到垫片文件后，每次播完垫片会尝试拉主源，如果主源恢复则自动切回主源。可根据需要的轮询检查时长来传入对应时长的垫片文件。
	BackupVodUrl *string `json:"BackupVodUrl,omitnil,omitempty" name:"BackupVodUrl"`
}

type ModifyLivePullStreamTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作人姓名。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 拉流源url列表。
	// SourceType为直播（PullLivePushLive）只可以填1个，
	// SourceType为点播（PullVodPushLive）可以填多个，上限30个。
	SourceUrls []*string `json:"SourceUrls,omitnil,omitempty" name:"SourceUrls"`

	// 开始时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于30天。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/267/38543#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 点播拉流转推循环次数。
	// -1：无限循环，直到任务结束。
	// 0：不循环。
	// >0：具体循环次数。次数和时间以先结束的为准。
	// 注意：拉流源为点播，该配置生效。
	VodLoopTimes *int64 `json:"VodLoopTimes,omitnil,omitempty" name:"VodLoopTimes"`

	// 点播更新SourceUrls后的播放方式：
	// ImmediateNewSource：立即从更新的拉流源开始播放；
	// ContinueBreakPoint：从上次断流url源的断点处继续，结束后再使用新的拉流源。
	// 注意：拉流源为点播，该配置生效。
	VodRefreshType *string `json:"VodRefreshType,omitnil,omitempty" name:"VodRefreshType"`

	// 任务状态：
	// enable - 启用，
	// pause - 暂停。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 选择需要回调的事件（不填则回调全部）：
	// TaskStart：任务启动回调，
	// TaskExit：任务停止回调，
	// VodSourceFileStart：从点播源文件开始拉流回调，
	// VodSourceFileFinish：从点播源文件拉流结束回调，
	// ResetTaskConfig：任务更新回调。
	CallbackEvents []*string `json:"CallbackEvents,omitnil,omitempty" name:"CallbackEvents"`

	// 自定义回调地址。
	// 相关事件会回调到该地址。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 指定播放文件索引。
	// 注意： 从1开始，不大于SourceUrls中文件个数。
	FileIndex *int64 `json:"FileIndex,omitnil,omitempty" name:"FileIndex"`

	// 指定播放文件偏移。
	// 注意：
	// 1. 单位：秒，配合FileIndex使用。
	OffsetTime *int64 `json:"OffsetTime,omitnil,omitempty" name:"OffsetTime"`

	// 指定任务 ID 修改任务。
	// 
	// 注意：该自定义任务 ID 只有在创建任务时指定了，才可在此处修改时使用。否则请使用系统返回的任务 ID。
	SpecifyTaskId *string `json:"SpecifyTaskId,omitnil,omitempty" name:"SpecifyTaskId"`

	// 目标 Url。
	// 换目标地址，会断流重推到新地址。
	ToUrl *string `json:"ToUrl,omitnil,omitempty" name:"ToUrl"`

	// 任务备注。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 备源的类型：
	// PullLivePushLive -直播，
	// PullVodPushLive -点播。
	// 注意：
	// 1. 仅当主源类型为直播源时，备源才会生效。
	// 2. 将该参数置为空，则可将任务去除备源信息。
	// 3. 主直播源拉流中断时，自动使用备源进行拉流。
	// 4. 如果备源为点播文件时，则每次轮播完点播文件就检查主源是否恢复，如果主源恢复则自动切回到主源，否则继续拉备源。
	BackupSourceType *string `json:"BackupSourceType,omitnil,omitempty" name:"BackupSourceType"`

	// 备源 URL。
	// 只允许填一个备源 URL
	BackupSourceUrl *string `json:"BackupSourceUrl,omitnil,omitempty" name:"BackupSourceUrl"`

	// 水印信息列表。
	// 注意：
	// 1. 最多支持4个不同位置的水印。
	// 2. 水印图片 URL 请使用合法外网可访问地址。
	// 3. 支持的水印图片格式：png，jpg等。
	// 4. 轮播任务修改水印后，轮播到下一个文件时新水印生效。
	// 5. 直播源任务修改水印后，水印立即生效。
	// 6. 清除水印时，需携带该水印列表参数，内容为空数组。
	// 7. 暂不支持动图水印。
	WatermarkList []*PullPushWatermarkInfo `json:"WatermarkList,omitnil,omitempty" name:"WatermarkList"`

	// 点播源是否启用本地推流模式，默认0，不启用。
	// 0 - 不启用。
	// 1 - 启用。
	// 注意：启用本地模式后，会将源列表中的 MP4 文件进行本地下载，优先使用本地已下载文件进行推流，提高点播源推流稳定性。使用本地下载文件推流时，会产生增值费用。
	VodLocalMode *int64 `json:"VodLocalMode,omitnil,omitempty" name:"VodLocalMode"`

	// 新的目标地址。传空值，则取消该地址的推流。
	// 传入新值，则替换原有地址。
	BackupToUrl *string `json:"BackupToUrl,omitnil,omitempty" name:"BackupToUrl"`

	// 点播垫片文件地址。注意：用于在主源拉不到时自动兜底到垫片文件，切到垫片文件后，每次播完垫片会尝试拉主源，如果主源恢复则自动切回主源。可根据需要的轮询检查时长来传入对应时长的垫片文件。
	BackupVodUrl *string `json:"BackupVodUrl,omitnil,omitempty" name:"BackupVodUrl"`
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
	delete(f, "SpecifyTaskId")
	delete(f, "ToUrl")
	delete(f, "Comment")
	delete(f, "BackupSourceType")
	delete(f, "BackupSourceUrl")
	delete(f, "WatermarkList")
	delete(f, "VodLocalMode")
	delete(f, "BackupToUrl")
	delete(f, "BackupVodUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLivePullStreamTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLivePullStreamTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLivePullStreamTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLivePullStreamTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLivePushAuthKeyRequestParams struct {
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	// 不传表示不修改当前值。
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 主鉴权key。
	// 不传表示不修改当前值。
	MasterAuthKey *string `json:"MasterAuthKey,omitnil,omitempty" name:"MasterAuthKey"`

	// 备鉴权key。
	// 不传表示不修改当前值。
	BackupAuthKey *string `json:"BackupAuthKey,omitnil,omitempty" name:"BackupAuthKey"`

	// 有效时间，单位：秒。
	AuthDelta *uint64 `json:"AuthDelta,omitnil,omitempty" name:"AuthDelta"`
}

type ModifyLivePushAuthKeyRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	// 不传表示不修改当前值。
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 主鉴权key。
	// 不传表示不修改当前值。
	MasterAuthKey *string `json:"MasterAuthKey,omitnil,omitempty" name:"MasterAuthKey"`

	// 备鉴权key。
	// 不传表示不修改当前值。
	BackupAuthKey *string `json:"BackupAuthKey,omitnil,omitempty" name:"BackupAuthKey"`

	// 有效时间，单位：秒。
	AuthDelta *uint64 `json:"AuthDelta,omitnil,omitempty" name:"AuthDelta"`
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

// Predefined struct for user
type ModifyLivePushAuthKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLivePushAuthKeyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLivePushAuthKeyResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLiveRecordTemplateRequestParams struct {
	// DescribeRecordTemplates接口获取到的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// FLV 录制参数，开启 FLV 录制时设置。
	FlvParam *RecordParam `json:"FlvParam,omitnil,omitempty" name:"FlvParam"`

	// HLS 录制参数，开启 HLS 录制时设置。
	HlsParam *RecordParam `json:"HlsParam,omitnil,omitempty" name:"HlsParam"`

	// MP4 录制参数，开启 MP4 录制时设置。
	Mp4Param *RecordParam `json:"Mp4Param,omitnil,omitempty" name:"Mp4Param"`

	// AAC 录制参数，开启 AAC 录制时设置。
	AacParam *RecordParam `json:"AacParam,omitnil,omitempty" name:"AacParam"`

	// HLS 录制定制参数。
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitnil,omitempty" name:"HlsSpecialParam"`

	// MP3 录制参数，开启 MP3 录制时设置。
	Mp3Param *RecordParam `json:"Mp3Param,omitnil,omitempty" name:"Mp3Param"`

	// 是否去除水印，类型为慢直播时此参数无效。
	RemoveWatermark *bool `json:"RemoveWatermark,omitnil,omitempty" name:"RemoveWatermark"`

	// FLV 录制定制参数。
	FlvSpecialParam *FlvSpecialParam `json:"FlvSpecialParam,omitnil,omitempty" name:"FlvSpecialParam"`
}

type ModifyLiveRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// DescribeRecordTemplates接口获取到的模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// FLV 录制参数，开启 FLV 录制时设置。
	FlvParam *RecordParam `json:"FlvParam,omitnil,omitempty" name:"FlvParam"`

	// HLS 录制参数，开启 HLS 录制时设置。
	HlsParam *RecordParam `json:"HlsParam,omitnil,omitempty" name:"HlsParam"`

	// MP4 录制参数，开启 MP4 录制时设置。
	Mp4Param *RecordParam `json:"Mp4Param,omitnil,omitempty" name:"Mp4Param"`

	// AAC 录制参数，开启 AAC 录制时设置。
	AacParam *RecordParam `json:"AacParam,omitnil,omitempty" name:"AacParam"`

	// HLS 录制定制参数。
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitnil,omitempty" name:"HlsSpecialParam"`

	// MP3 录制参数，开启 MP3 录制时设置。
	Mp3Param *RecordParam `json:"Mp3Param,omitnil,omitempty" name:"Mp3Param"`

	// 是否去除水印，类型为慢直播时此参数无效。
	RemoveWatermark *bool `json:"RemoveWatermark,omitnil,omitempty" name:"RemoveWatermark"`

	// FLV 录制定制参数。
	FlvSpecialParam *FlvSpecialParam `json:"FlvSpecialParam,omitnil,omitempty" name:"FlvSpecialParam"`
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
	delete(f, "RemoveWatermark")
	delete(f, "FlvSpecialParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLiveRecordTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLiveRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLiveRecordTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLiveSnapshotTemplateRequestParams struct {
	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Cos 应用 ID。
	// **注：此参数现在须必选。**
	CosAppId *int64 `json:"CosAppId,omitnil,omitempty" name:"CosAppId"`

	// Cos Bucket名称。
	// 注：CosBucket参数值不能包含-[appid] 部分。
	// **注：此参数现在须必选。**
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// Cos 地域。
	// **注：此参数现在须必选。**
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 模板名称。
	// 长度上限：255字节。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	// 长度上限：1024字节。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 截图间隔，单位s，默认10s。
	// 范围： 5s ~ 300s。
	SnapshotInterval *int64 `json:"SnapshotInterval,omitnil,omitempty" name:"SnapshotInterval"`

	// 截图宽度。默认：0（原始宽）。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度。默认：0（原始高）。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 是否开启鉴黄，默认 0 。
	// 0：不开启。
	// 1：开启。
	PornFlag *int64 `json:"PornFlag,omitnil,omitempty" name:"PornFlag"`

	// Cos Bucket文件夹前缀。
	CosPrefix *string `json:"CosPrefix,omitnil,omitempty" name:"CosPrefix"`

	// Cos 文件名称。
	CosFileName *string `json:"CosFileName,omitnil,omitempty" name:"CosFileName"`
}

type ModifyLiveSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Cos 应用 ID。
	// **注：此参数现在须必选。**
	CosAppId *int64 `json:"CosAppId,omitnil,omitempty" name:"CosAppId"`

	// Cos Bucket名称。
	// 注：CosBucket参数值不能包含-[appid] 部分。
	// **注：此参数现在须必选。**
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// Cos 地域。
	// **注：此参数现在须必选。**
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 模板名称。
	// 长度上限：255字节。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	// 长度上限：1024字节。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 截图间隔，单位s，默认10s。
	// 范围： 5s ~ 300s。
	SnapshotInterval *int64 `json:"SnapshotInterval,omitnil,omitempty" name:"SnapshotInterval"`

	// 截图宽度。默认：0（原始宽）。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度。默认：0（原始高）。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 是否开启鉴黄，默认 0 。
	// 0：不开启。
	// 1：开启。
	PornFlag *int64 `json:"PornFlag,omitnil,omitempty" name:"PornFlag"`

	// Cos Bucket文件夹前缀。
	CosPrefix *string `json:"CosPrefix,omitnil,omitempty" name:"CosPrefix"`

	// Cos 文件名称。
	CosFileName *string `json:"CosFileName,omitnil,omitempty" name:"CosFileName"`
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
	delete(f, "CosAppId")
	delete(f, "CosBucket")
	delete(f, "CosRegion")
	delete(f, "TemplateName")
	delete(f, "Description")
	delete(f, "SnapshotInterval")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "PornFlag")
	delete(f, "CosPrefix")
	delete(f, "CosFileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveSnapshotTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLiveSnapshotTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLiveSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLiveSnapshotTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLiveStreamMonitorRequestParams struct {
	// 监播任务ID。
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 监播任务的名称。长度128字节以内（一个汉字两个字节）。
	MonitorName *string `json:"MonitorName,omitnil,omitempty" name:"MonitorName"`

	// 监播任务输出信息。
	OutputInfo *LiveStreamMonitorOutputInfo `json:"OutputInfo,omitnil,omitempty" name:"OutputInfo"`

	// 待监播的输入流信息。
	InputList []*LiveStreamMonitorInputInfo `json:"InputList,omitnil,omitempty" name:"InputList"`

	// 监播事件通知策略。
	NotifyPolicy *LiveStreamMonitorNotifyPolicy `json:"NotifyPolicy,omitnil,omitempty" name:"NotifyPolicy"`

	// 智能语音识别语种：
	// 0 关闭 1 中文 2 英文 3 日文 4 韩文。
	AsrLanguage *uint64 `json:"AsrLanguage,omitnil,omitempty" name:"AsrLanguage"`

	// 智能文字识别语种：
	// 0 关闭 1 中、英文。
	OcrLanguage *uint64 `json:"OcrLanguage,omitnil,omitempty" name:"OcrLanguage"`

	// 语音识别输入流列表，1代表第一条输入流。
	AiAsrInputIndexList []*uint64 `json:"AiAsrInputIndexList,omitnil,omitempty" name:"AiAsrInputIndexList"`

	// 文字识别输入流列表，1代表第一条输入流。
	AiOcrInputIndexList []*uint64 `json:"AiOcrInputIndexList,omitnil,omitempty" name:"AiOcrInputIndexList"`

	// 是否开启断流检测。
	CheckStreamBroken *uint64 `json:"CheckStreamBroken,omitnil,omitempty" name:"CheckStreamBroken"`

	// 是否开启低帧率检测。
	CheckStreamLowFrameRate *uint64 `json:"CheckStreamLowFrameRate,omitnil,omitempty" name:"CheckStreamLowFrameRate"`

	// 是否存储监播事件到监播报告，以及是否允许查询监播报告。
	AllowMonitorReport *uint64 `json:"AllowMonitorReport,omitnil,omitempty" name:"AllowMonitorReport"`

	// 是否开启格式诊断。
	AiFormatDiagnose *uint64 `json:"AiFormatDiagnose,omitnil,omitempty" name:"AiFormatDiagnose"`

	// 是否开启内容质检。
	AiQualityControl *uint64 `json:"AiQualityControl,omitnil,omitempty" name:"AiQualityControl"`
}

type ModifyLiveStreamMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 监播任务ID。
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 监播任务的名称。长度128字节以内（一个汉字两个字节）。
	MonitorName *string `json:"MonitorName,omitnil,omitempty" name:"MonitorName"`

	// 监播任务输出信息。
	OutputInfo *LiveStreamMonitorOutputInfo `json:"OutputInfo,omitnil,omitempty" name:"OutputInfo"`

	// 待监播的输入流信息。
	InputList []*LiveStreamMonitorInputInfo `json:"InputList,omitnil,omitempty" name:"InputList"`

	// 监播事件通知策略。
	NotifyPolicy *LiveStreamMonitorNotifyPolicy `json:"NotifyPolicy,omitnil,omitempty" name:"NotifyPolicy"`

	// 智能语音识别语种：
	// 0 关闭 1 中文 2 英文 3 日文 4 韩文。
	AsrLanguage *uint64 `json:"AsrLanguage,omitnil,omitempty" name:"AsrLanguage"`

	// 智能文字识别语种：
	// 0 关闭 1 中、英文。
	OcrLanguage *uint64 `json:"OcrLanguage,omitnil,omitempty" name:"OcrLanguage"`

	// 语音识别输入流列表，1代表第一条输入流。
	AiAsrInputIndexList []*uint64 `json:"AiAsrInputIndexList,omitnil,omitempty" name:"AiAsrInputIndexList"`

	// 文字识别输入流列表，1代表第一条输入流。
	AiOcrInputIndexList []*uint64 `json:"AiOcrInputIndexList,omitnil,omitempty" name:"AiOcrInputIndexList"`

	// 是否开启断流检测。
	CheckStreamBroken *uint64 `json:"CheckStreamBroken,omitnil,omitempty" name:"CheckStreamBroken"`

	// 是否开启低帧率检测。
	CheckStreamLowFrameRate *uint64 `json:"CheckStreamLowFrameRate,omitnil,omitempty" name:"CheckStreamLowFrameRate"`

	// 是否存储监播事件到监播报告，以及是否允许查询监播报告。
	AllowMonitorReport *uint64 `json:"AllowMonitorReport,omitnil,omitempty" name:"AllowMonitorReport"`

	// 是否开启格式诊断。
	AiFormatDiagnose *uint64 `json:"AiFormatDiagnose,omitnil,omitempty" name:"AiFormatDiagnose"`

	// 是否开启内容质检。
	AiQualityControl *uint64 `json:"AiQualityControl,omitnil,omitempty" name:"AiQualityControl"`
}

func (r *ModifyLiveStreamMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveStreamMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorId")
	delete(f, "MonitorName")
	delete(f, "OutputInfo")
	delete(f, "InputList")
	delete(f, "NotifyPolicy")
	delete(f, "AsrLanguage")
	delete(f, "OcrLanguage")
	delete(f, "AiAsrInputIndexList")
	delete(f, "AiOcrInputIndexList")
	delete(f, "CheckStreamBroken")
	delete(f, "CheckStreamLowFrameRate")
	delete(f, "AllowMonitorReport")
	delete(f, "AiFormatDiagnose")
	delete(f, "AiQualityControl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveStreamMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLiveStreamMonitorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLiveStreamMonitorResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLiveStreamMonitorResponseParams `json:"Response"`
}

func (r *ModifyLiveStreamMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveStreamMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLiveTimeShiftTemplateRequestParams struct {
	// 时移模板id。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	// 长度上限：1024字节。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 时移时长。
	// 单位：s。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 分片时长。
	// 可取3-10。
	// 单位：s。
	// 默认值：5。
	ItemDuration *uint64 `json:"ItemDuration,omitnil,omitempty" name:"ItemDuration"`

	// 是否去除水印。
	// 传true则将录制原始流。
	// 默认值：false。
	RemoveWatermark *bool `json:"RemoveWatermark,omitnil,omitempty" name:"RemoveWatermark"`

	// 转码流id列表。
	// 此参数仅在 RemoveWatermark为false时生效。
	TranscodeTemplateIds []*int64 `json:"TranscodeTemplateIds,omitnil,omitempty" name:"TranscodeTemplateIds"`

	// 地域。
	// Mainland：中国大陆。
	// Overseas：海外及港澳台地区。
	// 默认值：Mainland。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type ModifyLiveTimeShiftTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 时移模板id。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	// 仅支持中文、英文、数字、_、-。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	// 长度上限：1024字节。
	// 仅支持中文、英文、数字、_、-。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 时移时长。
	// 单位：s。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 分片时长。
	// 可取3-10。
	// 单位：s。
	// 默认值：5。
	ItemDuration *uint64 `json:"ItemDuration,omitnil,omitempty" name:"ItemDuration"`

	// 是否去除水印。
	// 传true则将录制原始流。
	// 默认值：false。
	RemoveWatermark *bool `json:"RemoveWatermark,omitnil,omitempty" name:"RemoveWatermark"`

	// 转码流id列表。
	// 此参数仅在 RemoveWatermark为false时生效。
	TranscodeTemplateIds []*int64 `json:"TranscodeTemplateIds,omitnil,omitempty" name:"TranscodeTemplateIds"`

	// 地域。
	// Mainland：中国大陆。
	// Overseas：海外及港澳台地区。
	// 默认值：Mainland。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

func (r *ModifyLiveTimeShiftTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveTimeShiftTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateName")
	delete(f, "Description")
	delete(f, "Duration")
	delete(f, "ItemDuration")
	delete(f, "RemoveWatermark")
	delete(f, "TranscodeTemplateIds")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveTimeShiftTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLiveTimeShiftTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLiveTimeShiftTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLiveTimeShiftTemplateResponseParams `json:"Response"`
}

func (r *ModifyLiveTimeShiftTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveTimeShiftTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLiveTranscodeTemplateRequestParams struct {
	// 模板 Id。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 视频编码：h264/h265/origin，默认origin。
	// 
	// origin: 保持原始编码格式
	Vcodec *string `json:"Vcodec,omitnil,omitempty" name:"Vcodec"`

	// 音频编码：aac，默认aac。
	// 注意：当前该参数未生效，待后续支持！
	Acodec *string `json:"Acodec,omitnil,omitempty" name:"Acodec"`

	// 音频码率，默认0。
	// 范围：0-500。
	AudioBitrate *int64 `json:"AudioBitrate,omitnil,omitempty" name:"AudioBitrate"`

	// 模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 视频码率。范围：0kbps - 8000kbps。
	// 0为保持原始码率。
	// 注: 转码模板有码率唯一要求，最终保存的码率可能与输入码率有所差别。
	VideoBitrate *int64 `json:"VideoBitrate,omitnil,omitempty" name:"VideoBitrate"`

	// 宽。0-3000。
	// 数值必须是2的倍数，0是原始宽度
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 是否保留视频，0：否，1：是。默认1。
	NeedVideo *int64 `json:"NeedVideo,omitnil,omitempty" name:"NeedVideo"`

	// 是否保留音频，0：否，1：是。默认1。
	NeedAudio *int64 `json:"NeedAudio,omitnil,omitempty" name:"NeedAudio"`

	// 高。0-3000。
	// 数值必须是2的倍数，0是原始宽度
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 帧率，默认0。
	// 范围0-60
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 关键帧间隔，单位：秒。
	// 范围2-6
	Gop *int64 `json:"Gop,omitnil,omitempty" name:"Gop"`

	// 旋转角度，默认0。
	// 可取值：0，90，180，270
	Rotate *int64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`

	// 编码质量：
	// baseline/main/high。
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// 当设置的码率>原始码率时，是否以原始码率为准。
	// 0：否， 1：是
	// 默认 0。
	BitrateToOrig *int64 `json:"BitrateToOrig,omitnil,omitempty" name:"BitrateToOrig"`

	// 当设置的高度>原始高度时，是否以原始高度为准。
	// 0：否， 1：是
	// 默认 0。
	HeightToOrig *int64 `json:"HeightToOrig,omitnil,omitempty" name:"HeightToOrig"`

	// 当设置的帧率>原始帧率时，是否以原始帧率为准。
	// 0：否， 1：是
	// 默认 0。
	FpsToOrig *int64 `json:"FpsToOrig,omitnil,omitempty" name:"FpsToOrig"`

	// 极速高清视频码率压缩比。
	// 极速高清目标码率=VideoBitrate * (1-AdaptBitratePercent)
	// 
	// 取值范围：0.0到0.5
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitnil,omitempty" name:"AdaptBitratePercent"`

	// 是否以短边作为高度，0：否，1：是。默认0。
	ShortEdgeAsHeight *int64 `json:"ShortEdgeAsHeight,omitnil,omitempty" name:"ShortEdgeAsHeight"`

	// DRM 加密类型，可选值：fairplay、normalaes、widevine。
	// 不传递或者为空字符串，清空之前的DRM配置。
	DRMType *string `json:"DRMType,omitnil,omitempty" name:"DRMType"`

	// DRM 加密项，可选值：AUDIO、SD、HD、UHD1、UHD2，后四个为一组，同组中的内容只能选一个。
	// 不传递或者为空字符串，清空之前的DRM配置。
	DRMTracks *string `json:"DRMTracks,omitnil,omitempty" name:"DRMTracks"`
}

type ModifyLiveTranscodeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板 Id。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 视频编码：h264/h265/origin，默认origin。
	// 
	// origin: 保持原始编码格式
	Vcodec *string `json:"Vcodec,omitnil,omitempty" name:"Vcodec"`

	// 音频编码：aac，默认aac。
	// 注意：当前该参数未生效，待后续支持！
	Acodec *string `json:"Acodec,omitnil,omitempty" name:"Acodec"`

	// 音频码率，默认0。
	// 范围：0-500。
	AudioBitrate *int64 `json:"AudioBitrate,omitnil,omitempty" name:"AudioBitrate"`

	// 模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 视频码率。范围：0kbps - 8000kbps。
	// 0为保持原始码率。
	// 注: 转码模板有码率唯一要求，最终保存的码率可能与输入码率有所差别。
	VideoBitrate *int64 `json:"VideoBitrate,omitnil,omitempty" name:"VideoBitrate"`

	// 宽。0-3000。
	// 数值必须是2的倍数，0是原始宽度
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 是否保留视频，0：否，1：是。默认1。
	NeedVideo *int64 `json:"NeedVideo,omitnil,omitempty" name:"NeedVideo"`

	// 是否保留音频，0：否，1：是。默认1。
	NeedAudio *int64 `json:"NeedAudio,omitnil,omitempty" name:"NeedAudio"`

	// 高。0-3000。
	// 数值必须是2的倍数，0是原始宽度
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 帧率，默认0。
	// 范围0-60
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 关键帧间隔，单位：秒。
	// 范围2-6
	Gop *int64 `json:"Gop,omitnil,omitempty" name:"Gop"`

	// 旋转角度，默认0。
	// 可取值：0，90，180，270
	Rotate *int64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`

	// 编码质量：
	// baseline/main/high。
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// 当设置的码率>原始码率时，是否以原始码率为准。
	// 0：否， 1：是
	// 默认 0。
	BitrateToOrig *int64 `json:"BitrateToOrig,omitnil,omitempty" name:"BitrateToOrig"`

	// 当设置的高度>原始高度时，是否以原始高度为准。
	// 0：否， 1：是
	// 默认 0。
	HeightToOrig *int64 `json:"HeightToOrig,omitnil,omitempty" name:"HeightToOrig"`

	// 当设置的帧率>原始帧率时，是否以原始帧率为准。
	// 0：否， 1：是
	// 默认 0。
	FpsToOrig *int64 `json:"FpsToOrig,omitnil,omitempty" name:"FpsToOrig"`

	// 极速高清视频码率压缩比。
	// 极速高清目标码率=VideoBitrate * (1-AdaptBitratePercent)
	// 
	// 取值范围：0.0到0.5
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitnil,omitempty" name:"AdaptBitratePercent"`

	// 是否以短边作为高度，0：否，1：是。默认0。
	ShortEdgeAsHeight *int64 `json:"ShortEdgeAsHeight,omitnil,omitempty" name:"ShortEdgeAsHeight"`

	// DRM 加密类型，可选值：fairplay、normalaes、widevine。
	// 不传递或者为空字符串，清空之前的DRM配置。
	DRMType *string `json:"DRMType,omitnil,omitempty" name:"DRMType"`

	// DRM 加密项，可选值：AUDIO、SD、HD、UHD1、UHD2，后四个为一组，同组中的内容只能选一个。
	// 不传递或者为空字符串，清空之前的DRM配置。
	DRMTracks *string `json:"DRMTracks,omitnil,omitempty" name:"DRMTracks"`
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
	delete(f, "DRMType")
	delete(f, "DRMTracks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveTranscodeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLiveTranscodeTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLiveTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLiveTranscodeTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyPullStreamConfigRequestParams struct {
	// 配置 ID。
	// 获取来源：
	// 1. 创建拉流配置接口CreatePullStreamConfig返回的配置 ID。
	// 2. 通过查询接口DescribePullStreamConfigs获取配置 ID。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 源 URL，用于拉流的地址。目前可支持直播流及点播文件。
	// 注意：
	// 1. 多个点播 URL 之间使用空格拼接。
	// 2. 目前上限支持10个 URL。
	// 3. 支持拉流文件格式：FLV，RTMP，HLS，MP4。
	// 4. 使用标准三层样式，如：http://test.com/live/stream.flv。
	FromUrl *string `json:"FromUrl,omitnil,omitempty" name:"FromUrl"`

	// 目的 URL，用于推流的地址，目前限制该目标地址为腾讯域名。
	// 1. 仅支持 RTMP 协议。
	// 2. 使用标准三层样式，如：http://test.com/live/stream.flv。
	ToUrl *string `json:"ToUrl,omitnil,omitempty" name:"ToUrl"`

	// 区域 ID：
	// 1-深圳。
	// 2-上海。
	// 3-天津。
	// 4-中国香港。
	// 如有改动，需同时传入IspId。
	AreaId *int64 `json:"AreaId,omitnil,omitempty" name:"AreaId"`

	// 运营商 ID，
	// 1：电信。
	// 2：移动。
	// 3：联通。
	// 4：其他。
	// AreaId为4的时候，IspId只能为其他。如有改动，需同时传入AreaId。
	IspId *int64 `json:"IspId,omitnil,omitempty" name:"IspId"`

	// 开始时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于七天。
	// 
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ModifyPullStreamConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置 ID。
	// 获取来源：
	// 1. 创建拉流配置接口CreatePullStreamConfig返回的配置 ID。
	// 2. 通过查询接口DescribePullStreamConfigs获取配置 ID。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 源 URL，用于拉流的地址。目前可支持直播流及点播文件。
	// 注意：
	// 1. 多个点播 URL 之间使用空格拼接。
	// 2. 目前上限支持10个 URL。
	// 3. 支持拉流文件格式：FLV，RTMP，HLS，MP4。
	// 4. 使用标准三层样式，如：http://test.com/live/stream.flv。
	FromUrl *string `json:"FromUrl,omitnil,omitempty" name:"FromUrl"`

	// 目的 URL，用于推流的地址，目前限制该目标地址为腾讯域名。
	// 1. 仅支持 RTMP 协议。
	// 2. 使用标准三层样式，如：http://test.com/live/stream.flv。
	ToUrl *string `json:"ToUrl,omitnil,omitempty" name:"ToUrl"`

	// 区域 ID：
	// 1-深圳。
	// 2-上海。
	// 3-天津。
	// 4-中国香港。
	// 如有改动，需同时传入IspId。
	AreaId *int64 `json:"AreaId,omitnil,omitempty" name:"AreaId"`

	// 运营商 ID，
	// 1：电信。
	// 2：移动。
	// 3：联通。
	// 4：其他。
	// AreaId为4的时候，IspId只能为其他。如有改动，需同时传入AreaId。
	IspId *int64 `json:"IspId,omitnil,omitempty" name:"IspId"`

	// 开始时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于七天。
	// 
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
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

// Predefined struct for user
type ModifyPullStreamConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPullStreamConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPullStreamConfigResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyPullStreamStatusRequestParams struct {
	// 配置 ID 列表。
	ConfigIds []*string `json:"ConfigIds,omitnil,omitempty" name:"ConfigIds"`

	// 目标状态。0无效，2正在运行，4暂停。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyPullStreamStatusRequest struct {
	*tchttp.BaseRequest
	
	// 配置 ID 列表。
	ConfigIds []*string `json:"ConfigIds,omitnil,omitempty" name:"ConfigIds"`

	// 目标状态。0无效，2正在运行，4暂停。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
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

// Predefined struct for user
type ModifyPullStreamStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPullStreamStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPullStreamStatusResponseParams `json:"Response"`
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
	PlayDomain *string `json:"PlayDomain,omitnil,omitempty" name:"PlayDomain"`

	// 流id。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 播放码率，0表示原始码率。
	Rate *uint64 `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 播放协议，可选值包括 Unknown，Flv，Hls，Rtmp，Huyap2p。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 带宽，单位是Mbps。
	Bandwidth *float64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 在线人数，1分钟采样一个点，统计采样点的tcp链接数目。
	Online *uint64 `json:"Online,omitnil,omitempty" name:"Online"`

	// 请求数。
	Request *uint64 `json:"Request,omitnil,omitempty" name:"Request"`
}

type PadTemplate struct {
	// 模板id。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 垫片内容。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 模板创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板修改时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 断流等待时间。
	// 取值范围：0-30000。
	// 单位：ms。
	WaitDuration *uint64 `json:"WaitDuration,omitnil,omitempty" name:"WaitDuration"`

	// 最大垫片时长。
	// 取值范围：0 - 正无穷。
	// 单位：ms。
	MaxDuration *uint64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// 垫片内容类型： 1：图片，2：视频。 默认值：1。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type PlayAuthKeyInfo struct {
	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 是否启用:
	// 0: 关闭。
	// 1: 启用。
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 鉴权 Key。
	AuthKey *string `json:"AuthKey,omitnil,omitempty" name:"AuthKey"`

	// 有效时间，单位：秒。
	AuthDelta *uint64 `json:"AuthDelta,omitnil,omitempty" name:"AuthDelta"`

	// 鉴权 BackKey。
	AuthBackKey *string `json:"AuthBackKey,omitnil,omitempty" name:"AuthBackKey"`
}

type PlayCodeTotalInfo struct {
	// HTTP code，可选值包括:
	// 400，403，404，500，502，503，504。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 总次数。
	Num *uint64 `json:"Num,omitnil,omitempty" name:"Num"`
}

type PlayDataInfoByStream struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 总流量，单位: MB。
	TotalFlux *float64 `json:"TotalFlux,omitnil,omitempty" name:"TotalFlux"`
}

type PlayStatInfo struct {
	// 数据时间点。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 带宽/流量/请求数/并发连接数/下载速度的值，若没数据返回时该值为0。
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type PlaySumStatInfo struct {
	// 域名或流 ID。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 平均下载速度，
	// 单位: MB/s。
	// 计算公式: 每分钟的下载速度求平均值。
	AvgFluxPerSecond *float64 `json:"AvgFluxPerSecond,omitnil,omitempty" name:"AvgFluxPerSecond"`

	// 总流量，单位: MB。
	TotalFlux *float64 `json:"TotalFlux,omitnil,omitempty" name:"TotalFlux"`

	// 总请求数。
	TotalRequest *uint64 `json:"TotalRequest,omitnil,omitempty" name:"TotalRequest"`
}

type ProIspPlayCodeDataInfo struct {
	// 国家或地区。
	CountryAreaName *string `json:"CountryAreaName,omitnil,omitempty" name:"CountryAreaName"`

	// 省份。
	ProvinceName *string `json:"ProvinceName,omitnil,omitempty" name:"ProvinceName"`

	// 运营商。
	IspName *string `json:"IspName,omitnil,omitempty" name:"IspName"`

	// 错误码为2开头的次数。
	Code2xx *uint64 `json:"Code2xx,omitnil,omitempty" name:"Code2xx"`

	// 错误码为3开头的次数。
	Code3xx *uint64 `json:"Code3xx,omitnil,omitempty" name:"Code3xx"`

	// 错误码为4开头的次数。
	Code4xx *uint64 `json:"Code4xx,omitnil,omitempty" name:"Code4xx"`

	// 错误码为5开头的次数。
	Code5xx *uint64 `json:"Code5xx,omitnil,omitempty" name:"Code5xx"`
}

type ProIspPlaySumInfo struct {
	// 省份/运营商/国家或地区。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 总流量，单位: MB。
	TotalFlux *float64 `json:"TotalFlux,omitnil,omitempty" name:"TotalFlux"`

	// 总请求数。
	TotalRequest *uint64 `json:"TotalRequest,omitnil,omitempty" name:"TotalRequest"`

	// 平均下载流量，单位: MB/s。
	AvgFluxPerSecond *float64 `json:"AvgFluxPerSecond,omitnil,omitempty" name:"AvgFluxPerSecond"`
}

type PublishTime struct {
	// 推流时间。
	// UTC 格式，例如：2018-06-29T19:00:00Z。
	PublishTime *string `json:"PublishTime,omitnil,omitempty" name:"PublishTime"`
}

type PullPushWatermarkInfo struct {
	// 水印图片 URL。
	// URL中禁止包含的字符：
	// ;(){}$>`#"'|
	PictureUrl *string `json:"PictureUrl,omitnil,omitempty" name:"PictureUrl"`

	// 显示位置，X轴偏移，单位是百分比，默认 0。
	XPosition *int64 `json:"XPosition,omitnil,omitempty" name:"XPosition"`

	// 显示位置，Y轴偏移，单位是百分比，默认 0。
	YPosition *int64 `json:"YPosition,omitnil,omitempty" name:"YPosition"`

	// 水印宽度，占直播原始画面宽度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印高度，占直播原始画面高度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始高度。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 水印位置，默认 0。
	// 0：左上角。
	// 1：右上角。
	// 2：右下角。
	// 3：左下角。
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`
}

type PullStreamConfig struct {
	// 拉流配置 ID。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 源 URL。
	FromUrl *string `json:"FromUrl,omitnil,omitempty" name:"FromUrl"`

	// 目的 URL。
	ToUrl *string `json:"ToUrl,omitnil,omitempty" name:"ToUrl"`

	// 区域名。
	AreaName *string `json:"AreaName,omitnil,omitempty" name:"AreaName"`

	// 运营商名。
	IspName *string `json:"IspName,omitnil,omitempty" name:"IspName"`

	// 开始时间。
	// UTC格式时间，例如: 2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	// 
	// UTC格式时间，例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 状态:
	// 0: 无效。
	// 1: 初始状态。
	// 2: 正在运行。
	// 3: 拉起失败。
	// 4: 暂停。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type PullStreamTaskInfo struct {
	// 拉流任务Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 拉流源的类型：
	// PullLivePushLive -直播，
	// PullVodPushLive -点播，
	// PullPicPushLive -图片。
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 拉流源url列表。
	// SourceType为直播（PullLiveToLive）只可以填1个，
	// SourceType为点播（PullVodToLive）可以填多个，上限10个。
	SourceUrls []*string `json:"SourceUrls,omitnil,omitempty" name:"SourceUrls"`

	// 推流域名。
	// 将拉到的源推到该域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	// 将拉到的源推到该路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	// 将拉到的源推到该流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流参数。
	// 推流携带的自定义参数。
	PushArgs *string `json:"PushArgs,omitnil,omitempty" name:"PushArgs"`

	// 开始时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，注意：
	// 1. 结束时间必须大于开始时间；
	// 2. 结束时间和开始时间必须大于当前时间；
	// 3. 结束时间 和 开始时间 间隔必须小于七天。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务创建所在地域：
	// ap-beijing - 华北地区(北京)，
	// ap-shanghai -华东地区(上海)，
	// ap-guangzhou -华南地区(广州)，
	// ap-mumbai - 印度，
	// ap-hongkong - 香港，
	// eu-frankfurt - 德国，
	// ap-seoul - 韩国，
	// ap-bangkok - 泰国，
	// ap-singapore - 新加坡，
	// na-siliconvalley - 美西，
	// na-ashburn - 美东，
	// ap-tokyo - 日本。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 点播拉流转推循环次数。
	// -1：无限循环，直到任务结束。
	// 0：不循环。
	// >0：具体循环次数。次数和时间以先结束的为准。
	// 注意：拉流源为点播，该配置生效。
	VodLoopTimes *int64 `json:"VodLoopTimes,omitnil,omitempty" name:"VodLoopTimes"`

	// 点播更新SourceUrls后的播放方式：
	// ImmediateNewSource：立即从更新的拉流源开始播放；
	// ContinueBreakPoint：从上次断流url源的断点处继续，结束后再使用新的拉流源。
	// 
	// 注意：拉流源为点播，该配置生效。
	VodRefreshType *string `json:"VodRefreshType,omitnil,omitempty" name:"VodRefreshType"`

	// 任务创建时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务更新时间。
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建任务的操作者。
	CreateBy *string `json:"CreateBy,omitnil,omitempty" name:"CreateBy"`

	// 最后更新任务的操作者。
	UpdateBy *string `json:"UpdateBy,omitnil,omitempty" name:"UpdateBy"`

	// 回调地址。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 选择需要回调的事件：
	// TaskStart：任务启动回调，
	// TaskExit：任务停止回调，
	// VodSourceFileStart：从点播源文件开始拉流回调，
	// VodSourceFileFinish：从点播源文件拉流结束回调，
	// ResetTaskConfig：任务更新回调。
	CallbackEvents []*string `json:"CallbackEvents,omitnil,omitempty" name:"CallbackEvents"`

	// 注意：该信息暂不返回。
	// 最后一次回调信息。
	CallbackInfo *string `json:"CallbackInfo,omitnil,omitempty" name:"CallbackInfo"`

	// 注意：该信息暂不返回。
	// 错误信息。
	ErrorInfo *string `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 状态。
	// enable：生效中。
	// pause：暂停中。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 注意：该信息仅在查询单个任务时返回。
	// 任务最新拉流信息。
	// 包含：源 url，偏移时间，上报时间。
	RecentPullInfo *RecentPullInfo `json:"RecentPullInfo,omitnil,omitempty" name:"RecentPullInfo"`

	// 任务备注信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 备源类型：
	// PullLivePushLive -直播，
	// PullVodPushLive -点播。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupSourceType *string `json:"BackupSourceType,omitnil,omitempty" name:"BackupSourceType"`

	// 备源URL。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupSourceUrl *string `json:"BackupSourceUrl,omitnil,omitempty" name:"BackupSourceUrl"`

	// 水印信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkList []*PullPushWatermarkInfo `json:"WatermarkList,omitnil,omitempty" name:"WatermarkList"`

	// 点播源是否启用本地推流模式，默认0，不启用。
	// 0 - 不启用。
	// 1 - 启用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VodLocalMode *int64 `json:"VodLocalMode,omitnil,omitempty" name:"VodLocalMode"`

	// 录制模板 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordTemplateId *string `json:"RecordTemplateId,omitnil,omitempty" name:"RecordTemplateId"`

	// 新增的推流地址。用于单任务推两路场景。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupToUrl *string `json:"BackupToUrl,omitnil,omitempty" name:"BackupToUrl"`

	// 直播转码模板。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTemplateName *string `json:"TranscodeTemplateName,omitnil,omitempty" name:"TranscodeTemplateName"`
}

type PushAuthKeyInfo struct {
	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 主鉴权 Key。
	MasterAuthKey *string `json:"MasterAuthKey,omitnil,omitempty" name:"MasterAuthKey"`

	// 备鉴权 Key。
	BackupAuthKey *string `json:"BackupAuthKey,omitnil,omitempty" name:"BackupAuthKey"`

	// 有效时间，单位：秒。
	AuthDelta *uint64 `json:"AuthDelta,omitnil,omitempty" name:"AuthDelta"`
}

type PushDataInfo struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流客户端 IP。
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 接流服务器 IP。
	ServerIp *string `json:"ServerIp,omitnil,omitempty" name:"ServerIp"`

	// 推流视频帧率，单位: Hz。
	VideoFps *uint64 `json:"VideoFps,omitnil,omitempty" name:"VideoFps"`

	// 推流视频码率，单位: bps。
	VideoSpeed *uint64 `json:"VideoSpeed,omitnil,omitempty" name:"VideoSpeed"`

	// 推流音频帧率，单位: Hz。
	AudioFps *uint64 `json:"AudioFps,omitnil,omitempty" name:"AudioFps"`

	// 推流音频码率，单位: bps。
	AudioSpeed *uint64 `json:"AudioSpeed,omitnil,omitempty" name:"AudioSpeed"`

	// 推流域名。
	PushDomain *string `json:"PushDomain,omitnil,omitempty" name:"PushDomain"`

	// 推流开始时间。
	BeginPushTime *string `json:"BeginPushTime,omitnil,omitempty" name:"BeginPushTime"`

	// 音频编码格式，
	// 例："AAC"。
	Acodec *string `json:"Acodec,omitnil,omitempty" name:"Acodec"`

	// 视频编码格式，
	// 例："H264"。
	Vcodec *string `json:"Vcodec,omitnil,omitempty" name:"Vcodec"`

	// 分辨率。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 采样率。
	AsampleRate *uint64 `json:"AsampleRate,omitnil,omitempty" name:"AsampleRate"`

	// metadata 中的音频码率，单位: bps。
	MetaAudioSpeed *uint64 `json:"MetaAudioSpeed,omitnil,omitempty" name:"MetaAudioSpeed"`

	// metadata 中的视频码率，单位: bps。
	MetaVideoSpeed *uint64 `json:"MetaVideoSpeed,omitnil,omitempty" name:"MetaVideoSpeed"`

	// metadata 中的帧率。
	MetaFps *uint64 `json:"MetaFps,omitnil,omitempty" name:"MetaFps"`
}

type PushLogInfo struct {
	// 日志名称。
	LogName *string `json:"LogName,omitnil,omitempty" name:"LogName"`

	// 日志下载地址。
	LogUrl *string `json:"LogUrl,omitnil,omitempty" name:"LogUrl"`

	// 日志时间。UTC 格式，例如：2018-11-29T19:00:00Z。
	// 注意：
	// 1. 北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	LogTime *string `json:"LogTime,omitnil,omitempty" name:"LogTime"`

	// 文件大小，单位字节。
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`
}

type PushQualityData struct {
	// 数据时间，使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 推流域名。
	PushDomain *string `json:"PushDomain,omitnil,omitempty" name:"PushDomain"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流客户端 IP。
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 开始推流时间，格式: %Y-%m-%d %H:%M:%S.%ms，精确到毫秒级。
	BeginPushTime *string `json:"BeginPushTime,omitnil,omitempty" name:"BeginPushTime"`

	// 分辨率信息。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 视频编码格式。
	VCodec *string `json:"VCodec,omitnil,omitempty" name:"VCodec"`

	// 音频编码格式。
	ACodec *string `json:"ACodec,omitnil,omitempty" name:"ACodec"`

	// 推流序列号，用来唯一的标志一次推流。
	Sequence *string `json:"Sequence,omitnil,omitempty" name:"Sequence"`

	// 视频帧率。
	VideoFps *uint64 `json:"VideoFps,omitnil,omitempty" name:"VideoFps"`

	// 视频码率，单位: bps。
	VideoRate *uint64 `json:"VideoRate,omitnil,omitempty" name:"VideoRate"`

	// 音频帧率。
	AudioFps *uint64 `json:"AudioFps,omitnil,omitempty" name:"AudioFps"`

	// 音频码率，单位: bps。
	AudioRate *uint64 `json:"AudioRate,omitnil,omitempty" name:"AudioRate"`

	// 本地流逝时间，单位: ms，音视频流逝时间与本地流逝时间的差距越大表示推流质量越差，上行卡顿越严重。
	LocalTs *uint64 `json:"LocalTs,omitnil,omitempty" name:"LocalTs"`

	// 视频流逝时间，单位: ms。
	VideoTs *uint64 `json:"VideoTs,omitnil,omitempty" name:"VideoTs"`

	// 音频流逝时间，单位: ms。
	AudioTs *uint64 `json:"AudioTs,omitnil,omitempty" name:"AudioTs"`

	// metadata 中的视频码率，单位: kbps。
	MetaVideoRate *uint64 `json:"MetaVideoRate,omitnil,omitempty" name:"MetaVideoRate"`

	// metadata 中的音频码率，单位: kbps。
	MetaAudioRate *uint64 `json:"MetaAudioRate,omitnil,omitempty" name:"MetaAudioRate"`

	// metadata 中的帧率。
	MateFps *uint64 `json:"MateFps,omitnil,omitempty" name:"MateFps"`

	// 推流参数
	StreamParam *string `json:"StreamParam,omitnil,omitempty" name:"StreamParam"`

	// 带宽，单位Mbps。
	Bandwidth *float64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 流量，单位MB。
	Flux *float64 `json:"Flux,omitnil,omitempty" name:"Flux"`

	// 推流服务端 IP。
	ServerIp *string `json:"ServerIp,omitnil,omitempty" name:"ServerIp"`

	// 关键帧间隔 GOP ，单位：ms。
	GopSize *int64 `json:"GopSize,omitnil,omitempty" name:"GopSize"`
}

type RecentPullInfo struct {
	// 当前正在拉的文件地址。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 当前正在拉的文件偏移，单位：秒。
	OffsetTime *uint64 `json:"OffsetTime,omitnil,omitempty" name:"OffsetTime"`

	// 最新上报偏移信息时间。UTC格式。
	// 如：2020-07-23T03:20:39Z。
	// 注意：与北京时间相差八小时。
	ReportTime *string `json:"ReportTime,omitnil,omitempty" name:"ReportTime"`

	// 已经轮播的次数。
	LoopedTimes *int64 `json:"LoopedTimes,omitnil,omitempty" name:"LoopedTimes"`
}

type RecordParam struct {
	// 录制间隔。
	// 单位秒，默认：1800。
	// 取值范围：30-7200。
	// 此参数对 HLS 无效，当录制 HLS 时从推流到断流生成一个文件。
	RecordInterval *int64 `json:"RecordInterval,omitnil,omitempty" name:"RecordInterval"`

	// 录制存储时长。
	// 单位秒，取值范围： 0 - 1500天。
	// 0：表示永久存储。
	StorageTime *int64 `json:"StorageTime,omitnil,omitempty" name:"StorageTime"`

	// 是否开启当前格式录制，默认值为0，0：否， 1：是。
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 点播子应用 ID。
	VodSubAppId *int64 `json:"VodSubAppId,omitnil,omitempty" name:"VodSubAppId"`

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
	VodFileName *string `json:"VodFileName,omitnil,omitempty" name:"VodFileName"`

	// 任务流
	// 注意：此字段可能返回 null，表示取不到有效值。
	Procedure *string `json:"Procedure,omitnil,omitempty" name:"Procedure"`

	// 视频存储策略。
	// normal：标准存储。
	// cold：低频存储。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageMode *string `json:"StorageMode,omitnil,omitempty" name:"StorageMode"`

	// 点播应用分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassId *int64 `json:"ClassId,omitnil,omitempty" name:"ClassId"`

	// 存储至 cos 的 bucket 桶名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`

	// 存储至 cos 的 bucket 区域。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucketRegion *string `json:"CosBucketRegion,omitnil,omitempty" name:"CosBucketRegion"`

	// 存储至 cos 的 bucket 路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucketPath *string `json:"CosBucketPath,omitnil,omitempty" name:"CosBucketPath"`
}

type RecordTask struct {
	// 录制任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 任务开始时间，Unix时间戳。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间，Unix时间戳。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 录制模板ID。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 调用 StopRecordTask 停止任务时间，Unix时间戳。值为0表示未曾调用接口停止任务。
	Stopped *uint64 `json:"Stopped,omitnil,omitempty" name:"Stopped"`
}

type RecordTemplateInfo struct {
	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// FLV 录制参数。
	FlvParam *RecordParam `json:"FlvParam,omitnil,omitempty" name:"FlvParam"`

	// HLS 录制参数。
	HlsParam *RecordParam `json:"HlsParam,omitnil,omitempty" name:"HlsParam"`

	// MP4 录制参数。
	Mp4Param *RecordParam `json:"Mp4Param,omitnil,omitempty" name:"Mp4Param"`

	// AAC 录制参数。
	AacParam *RecordParam `json:"AacParam,omitnil,omitempty" name:"AacParam"`

	// 0：普通直播，
	// 1：慢直播。
	IsDelayLive *int64 `json:"IsDelayLive,omitnil,omitempty" name:"IsDelayLive"`

	// HLS 录制定制参数。
	HlsSpecialParam *HlsSpecialParam `json:"HlsSpecialParam,omitnil,omitempty" name:"HlsSpecialParam"`

	// MP3 录制参数。
	Mp3Param *RecordParam `json:"Mp3Param,omitnil,omitempty" name:"Mp3Param"`

	// 是否去除水印。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoveWatermark *bool `json:"RemoveWatermark,omitnil,omitempty" name:"RemoveWatermark"`

	// 是否存储至cos
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosStore *int64 `json:"CosStore,omitnil,omitempty" name:"CosStore"`

	// FLV 录制定制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlvSpecialParam *FlvSpecialParam `json:"FlvSpecialParam,omitnil,omitempty" name:"FlvSpecialParam"`
}

type RefererAuthConfig struct {
	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 是否启用，0：关闭，1：启用。
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 名单类型，0：黑名单，1：白名单。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否允许空Referer，0：不允许，1：允许。
	AllowEmpty *int64 `json:"AllowEmpty,omitnil,omitempty" name:"AllowEmpty"`

	// 名单列表，以分号(;)分隔。
	Rules *string `json:"Rules,omitnil,omitempty" name:"Rules"`
}

// Predefined struct for user
type ReleaseCasterRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type ReleaseCasterRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *ReleaseCasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseCasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseCasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseCasterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReleaseCasterResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseCasterResponseParams `json:"Response"`
}

func (r *ReleaseCasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseCasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartLivePullStreamTaskRequestParams struct {
	// 任务 Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作人备注名称。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type RestartLivePullStreamTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务 Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作人备注名称。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *RestartLivePullStreamTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartLivePullStreamTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartLivePullStreamTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartLivePullStreamTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartLivePullStreamTaskResponse struct {
	*tchttp.BaseResponse
	Response *RestartLivePullStreamTaskResponseParams `json:"Response"`
}

func (r *RestartLivePullStreamTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartLivePullStreamTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeDelayLiveStreamRequestParams struct {
	// 推流路径，与推流和播放地址中的AppName保持一致，默认为live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type ResumeDelayLiveStreamRequest struct {
	*tchttp.BaseRequest
	
	// 推流路径，与推流和播放地址中的AppName保持一致，默认为live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
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

// Predefined struct for user
type ResumeDelayLiveStreamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeDelayLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *ResumeDelayLiveStreamResponseParams `json:"Response"`
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

// Predefined struct for user
type ResumeLiveStreamRequestParams struct {
	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type ResumeLiveStreamRequest struct {
	*tchttp.BaseRequest
	
	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
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

// Predefined struct for user
type ResumeLiveStreamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *ResumeLiveStreamResponseParams `json:"Response"`
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
	// 注：此字段为北京时间（UTC+8时区）。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 规则更新时间。
	// 注：此字段为北京时间（UTC+8时区）。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type ScreenshotTask struct {
	// 截图任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 任务开始时间，Unix时间戳。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间，Unix时间戳。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 截图模板ID。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 调用 StopScreenshotTask 停止任务时间，Unix时间戳。值为0表示未曾调用接口停止任务。
	Stopped *uint64 `json:"Stopped,omitnil,omitempty" name:"Stopped"`
}

type SnapshotTemplateInfo struct {
	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 截图时间间隔，5-300秒。
	SnapshotInterval *int64 `json:"SnapshotInterval,omitnil,omitempty" name:"SnapshotInterval"`

	// 截图宽度，范围：0-3000。 
	// 0：原始宽度并适配原始比例。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 截图高度，范围：0-2000。
	// 0：原始高度并适配原始比例。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 是否开启鉴黄，0：不开启，1：开启。
	PornFlag *int64 `json:"PornFlag,omitnil,omitempty" name:"PornFlag"`

	// Cos 应用 ID。
	CosAppId *int64 `json:"CosAppId,omitnil,omitempty" name:"CosAppId"`

	// Cos Bucket名称。
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// Cos 地域。
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Cos Bucket文件夹前缀。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPrefix *string `json:"CosPrefix,omitnil,omitempty" name:"CosPrefix"`

	// Cos 文件名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosFileName *string `json:"CosFileName,omitnil,omitempty" name:"CosFileName"`
}

// Predefined struct for user
type StartLiveStreamMonitorRequestParams struct {
	// 监播ID。
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 监播画面声音InputIndex,支持多个输入声音。
	// 取值范围 InputIndex必须已经存在。
	// 不填默认无声音输出。
	AudibleInputIndexList []*uint64 `json:"AudibleInputIndexList,omitnil,omitempty" name:"AudibleInputIndexList"`
}

type StartLiveStreamMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 监播ID。
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 监播画面声音InputIndex,支持多个输入声音。
	// 取值范围 InputIndex必须已经存在。
	// 不填默认无声音输出。
	AudibleInputIndexList []*uint64 `json:"AudibleInputIndexList,omitnil,omitempty" name:"AudibleInputIndexList"`
}

func (r *StartLiveStreamMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartLiveStreamMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorId")
	delete(f, "AudibleInputIndexList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartLiveStreamMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartLiveStreamMonitorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartLiveStreamMonitorResponse struct {
	*tchttp.BaseResponse
	Response *StartLiveStreamMonitorResponseParams `json:"Response"`
}

func (r *StartLiveStreamMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartLiveStreamMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCasterPgmRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type StopCasterPgmRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *StopCasterPgmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCasterPgmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCasterPgmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCasterPgmResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopCasterPgmResponse struct {
	*tchttp.BaseResponse
	Response *StopCasterPgmResponseParams `json:"Response"`
}

func (r *StopCasterPgmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCasterPgmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCasterPvwRequestParams struct {
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

type StopCasterPvwRequest struct {
	*tchttp.BaseRequest
	
	// 导播台ID。
	CasterId *uint64 `json:"CasterId,omitnil,omitempty" name:"CasterId"`
}

func (r *StopCasterPvwRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCasterPvwRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCasterPvwRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCasterPvwResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopCasterPvwResponse struct {
	*tchttp.BaseResponse
	Response *StopCasterPvwResponseParams `json:"Response"`
}

func (r *StopCasterPvwResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCasterPvwResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopLivePadProcessorRequestParams struct {
	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 您的推流域名。
	PushDomainName *string `json:"PushDomainName,omitnil,omitempty" name:"PushDomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 操作人备注信息。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type StopLivePadProcessorRequest struct {
	*tchttp.BaseRequest
	
	// 推流路径，与推流和播放地址中的AppName保持一致，默认为 live。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 您的推流域名。
	PushDomainName *string `json:"PushDomainName,omitnil,omitempty" name:"PushDomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 操作人备注信息。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *StopLivePadProcessorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLivePadProcessorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "PushDomainName")
	delete(f, "StreamName")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopLivePadProcessorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopLivePadProcessorResponseParams struct {
	// 处理结果信息。
	ResultMessage *string `json:"ResultMessage,omitnil,omitempty" name:"ResultMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopLivePadProcessorResponse struct {
	*tchttp.BaseResponse
	Response *StopLivePadProcessorResponseParams `json:"Response"`
}

func (r *StopLivePadProcessorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLivePadProcessorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopLiveRecordRequestParams struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 任务ID，由CreateLiveRecord接口返回。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopLiveRecordRequest struct {
	*tchttp.BaseRequest
	
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 任务ID，由CreateLiveRecord接口返回。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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

// Predefined struct for user
type StopLiveRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopLiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *StopLiveRecordResponseParams `json:"Response"`
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

// Predefined struct for user
type StopLiveStreamMonitorRequestParams struct {
	// 监播ID
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

type StopLiveStreamMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 监播ID
	MonitorId *string `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`
}

func (r *StopLiveStreamMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLiveStreamMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopLiveStreamMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopLiveStreamMonitorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopLiveStreamMonitorResponse struct {
	*tchttp.BaseResponse
	Response *StopLiveStreamMonitorResponseParams `json:"Response"`
}

func (r *StopLiveStreamMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLiveStreamMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRecordTaskRequestParams struct {
	// 录制任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopRecordTaskRequest struct {
	*tchttp.BaseRequest
	
	// 录制任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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

// Predefined struct for user
type StopRecordTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopRecordTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopRecordTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type StopScreenshotTaskRequestParams struct {
	// 截图任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopScreenshotTaskRequest struct {
	*tchttp.BaseRequest
	
	// 截图任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopScreenshotTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopScreenshotTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopScreenshotTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopScreenshotTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopScreenshotTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopScreenshotTaskResponseParams `json:"Response"`
}

func (r *StopScreenshotTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopScreenshotTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StreamEventInfo struct {
	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流开始时间。
	// UTC 格式时间，例如：2019-01-07T12:00:00Z。
	StreamStartTime *string `json:"StreamStartTime,omitnil,omitempty" name:"StreamStartTime"`

	// 推流结束时间。
	// UTC 格式时间，例如：2019-01-07T15:00:00Z。
	StreamEndTime *string `json:"StreamEndTime,omitnil,omitempty" name:"StreamEndTime"`

	// 停止原因。
	StopReason *string `json:"StopReason,omitnil,omitempty" name:"StopReason"`

	// 推流持续时长，单位：秒。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 主播 IP。
	// 当客户端为内网推流时，展示为: - 。
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 分辨率。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`
}

type StreamName struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 推流开始时间。
	// UTC格式时间，例如：2019-01-07T12:00:00Z。
	StreamStartTime *string `json:"StreamStartTime,omitnil,omitempty" name:"StreamStartTime"`

	// 推流结束时间。
	// UTC格式时间，例如：2019-01-07T15:00:00Z。
	StreamEndTime *string `json:"StreamEndTime,omitnil,omitempty" name:"StreamEndTime"`

	// 停止原因。
	StopReason *string `json:"StopReason,omitnil,omitempty" name:"StopReason"`

	// 推流持续时长，单位：秒。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 主播 IP。
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 分辨率。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`
}

type StreamOnlineInfo struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 推流时间列表
	PublishTimeList []*PublishTime `json:"PublishTimeList,omitnil,omitempty" name:"PublishTimeList"`

	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 推流域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 流是否推送到延播。
	// 0 - 无延播，
	// 1 - 有延播。
	PushToDelay *int64 `json:"PushToDelay,omitnil,omitempty" name:"PushToDelay"`
}

// Predefined struct for user
type SwitchBackupStreamRequestParams struct {
	// 推流域名。
	PushDomainName *string `json:"PushDomainName,omitnil,omitempty" name:"PushDomainName"`

	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 查询接口获取到该流所有在推的上行 Sequence。指定要切到的目标上行 Sequence。
	UpstreamSequence *string `json:"UpstreamSequence,omitnil,omitempty" name:"UpstreamSequence"`
}

type SwitchBackupStreamRequest struct {
	*tchttp.BaseRequest
	
	// 推流域名。
	PushDomainName *string `json:"PushDomainName,omitnil,omitempty" name:"PushDomainName"`

	// 应用名称。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 查询接口获取到该流所有在推的上行 Sequence。指定要切到的目标上行 Sequence。
	UpstreamSequence *string `json:"UpstreamSequence,omitnil,omitempty" name:"UpstreamSequence"`
}

func (r *SwitchBackupStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchBackupStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PushDomainName")
	delete(f, "AppName")
	delete(f, "StreamName")
	delete(f, "UpstreamSequence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchBackupStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchBackupStreamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchBackupStreamResponse struct {
	*tchttp.BaseResponse
	Response *SwitchBackupStreamResponseParams `json:"Response"`
}

func (r *SwitchBackupStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchBackupStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskDurationInfo struct {
	// 拉流转推任务的时间点
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 拉流转推任务的时长，单位为分钟
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 拉流转推任务的时长，单位为秒
	DurationSecond *uint64 `json:"DurationSecond,omitnil,omitempty" name:"DurationSecond"`
}

type TaskStatusInfo struct {
	// 当前使用的源 URL。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 点播源任务的轮播次数。
	LoopedTimes *int64 `json:"LoopedTimes,omitnil,omitempty" name:"LoopedTimes"`

	// 点播源的播放偏移，单位：秒。
	OffsetTime *int64 `json:"OffsetTime,omitnil,omitempty" name:"OffsetTime"`

	// 最新心跳上报时间。UTC时间，例如：
	// 2022-02-11T10:00:00Z。
	// 注意：UTC时间与北京时间相差八小时。
	ReportTime *string `json:"ReportTime,omitnil,omitempty" name:"ReportTime"`

	// 实际运行状态：
	// active - 活跃，
	// inactive - 不活跃。
	RunStatus *string `json:"RunStatus,omitnil,omitempty" name:"RunStatus"`

	// 点播源的文件时长，单位：秒。
	FileDuration *int64 `json:"FileDuration,omitnil,omitempty" name:"FileDuration"`

	// 下一进度点播文件 URL。
	NextFileUrl *string `json:"NextFileUrl,omitnil,omitempty" name:"NextFileUrl"`
}

type TemplateInfo struct {
	// 视频编码：h264/h265/origin，默认h264。
	// 
	// origin: 保持原始编码格式
	Vcodec *string `json:"Vcodec,omitnil,omitempty" name:"Vcodec"`

	// 视频码率。范围：0kbps - 8000kbps。
	// 0为保持原始码率。
	// 注: 转码模板有码率唯一要求，最终保存的码率可能与输入码率有所差别。
	VideoBitrate *int64 `json:"VideoBitrate,omitnil,omitempty" name:"VideoBitrate"`

	// 音频编码：aac，默认aac。
	// 注意：当前该参数未生效，待后续支持！
	Acodec *string `json:"Acodec,omitnil,omitempty" name:"Acodec"`

	// 音频码率。取值范围：0kbps - 500kbps。
	// 默认0。
	AudioBitrate *int64 `json:"AudioBitrate,omitnil,omitempty" name:"AudioBitrate"`

	// 宽，默认0。
	// 范围[0-3000]
	// 数值必须是2的倍数，0是原始宽度
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 高，默认0。
	// 范围[0-3000]
	// 数值必须是2的倍数，0是原始宽度
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 帧率，默认0。
	// 范围0-60fps
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 关键帧间隔，单位：秒。
	// 默认原始的间隔
	// 范围2-6
	Gop *int64 `json:"Gop,omitnil,omitempty" name:"Gop"`

	// 旋转角度，默认0。
	// 可取值：0，90，180，270
	Rotate *int64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`

	// 编码质量：
	// baseline/main/high。默认baseline
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// 当设置的码率>原始码率时，是否以原始码率为准。
	// 0：否， 1：是
	// 默认 0。
	BitrateToOrig *int64 `json:"BitrateToOrig,omitnil,omitempty" name:"BitrateToOrig"`

	// 当设置的高度>原始高度时，是否以原始高度为准。
	// 0：否， 1：是
	// 默认 0。
	HeightToOrig *int64 `json:"HeightToOrig,omitnil,omitempty" name:"HeightToOrig"`

	// 当设置的帧率>原始帧率时，是否以原始帧率为准。
	// 0：否， 1：是
	// 默认 0。
	FpsToOrig *int64 `json:"FpsToOrig,omitnil,omitempty" name:"FpsToOrig"`

	// 是否保留视频。0：否，1：是。
	NeedVideo *int64 `json:"NeedVideo,omitnil,omitempty" name:"NeedVideo"`

	// 是否保留音频。0：否，1：是。
	NeedAudio *int64 `json:"NeedAudio,omitnil,omitempty" name:"NeedAudio"`

	// 模板 ID。
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否是极速高清模板，0：否，1：是。默认0。
	AiTransCode *int64 `json:"AiTransCode,omitnil,omitempty" name:"AiTransCode"`

	// 极速高清视频码率压缩比。
	// 极速高清目标码率=VideoBitrate * (1-AdaptBitratePercent)
	// 
	// 取值范围：0.0到0.5
	AdaptBitratePercent *float64 `json:"AdaptBitratePercent,omitnil,omitempty" name:"AdaptBitratePercent"`

	// 是否以短边作为高度，0：否，1：是。默认0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShortEdgeAsHeight *int64 `json:"ShortEdgeAsHeight,omitnil,omitempty" name:"ShortEdgeAsHeight"`

	// DRM 加密类型，可选值：fairplay、normalaes、widevine。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DRMType *string `json:"DRMType,omitnil,omitempty" name:"DRMType"`

	// DRM 加密项，多个用|分割，可选值：AUDIO、SD、HD、UHD1、UHD2，后四个为一组，同组中的内容只能选一个。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DRMTracks *string `json:"DRMTracks,omitnil,omitempty" name:"DRMTracks"`
}

type TimeShiftBillData struct {
	// 推流域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 时移文件时长，单位分钟。
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 时移配置天数，单位天。
	StoragePeriod *float64 `json:"StoragePeriod,omitnil,omitempty" name:"StoragePeriod"`

	// 时间点，格式: yyyy-mm-ddTHH:MM:SSZ。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 时移总时长，单位分钟。
	TotalDuration *float64 `json:"TotalDuration,omitnil,omitempty" name:"TotalDuration"`
}

type TimeShiftRecord struct {
	// 时移录制会话标识。
	Sid *string `json:"Sid,omitnil,omitempty" name:"Sid"`

	// 录制会话开始时间，Unix 时间戳。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 录制会话结束时间，Unix 时间戳。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type TimeShiftStreamInfo struct {
	// 推流域名所属组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainGroup *string `json:"DomainGroup,omitnil,omitempty" name:"DomainGroup"`

	// 推流域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 推流路径。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 流起始时间，Unix 时间戳。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止查询时流结束时间，Unix 时间戳。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 转码模板ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransCodeId *uint64 `json:"TransCodeId,omitnil,omitempty" name:"TransCodeId"`

	// 流类型，取值0为原始流，1为水印流，2为转码流。
	StreamType *int64 `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// 时移数据存储时长，单位秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 时移自适应码率子流列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeShiftSubStreams []*TimeShiftSubStream `json:"TimeShiftSubStreams,omitnil,omitempty" name:"TimeShiftSubStreams"`
}

type TimeShiftSubStream struct {
	// 时移自适应码率子流转码模板ID。
	TransCodeId *uint64 `json:"TransCodeId,omitnil,omitempty" name:"TransCodeId"`
}

type TimeShiftTemplate struct {
	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 时移时长。
	// 单位：秒。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 分片时长。
	// 可取3-10。
	// 单位：s。
	// 默认值：5。
	ItemDuration *uint64 `json:"ItemDuration,omitnil,omitempty" name:"ItemDuration"`

	// 模板id。
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 地域：
	// Mainland：中国大陆；
	// Overseas：海外及港澳台地区；
	// 默认值：Mainland。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 是否去除水印。
	// 为true则将录制原始流。
	// 默认值：false。
	RemoveWatermark *bool `json:"RemoveWatermark,omitnil,omitempty" name:"RemoveWatermark"`

	// 转码流id列表。
	// 此参数仅在 RemoveWatermark为false时生效。
	TranscodeTemplateIds []*uint64 `json:"TranscodeTemplateIds,omitnil,omitempty" name:"TranscodeTemplateIds"`
}

type TimeShiftWriteSizeData struct {
	// 区域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 时间，格式为：yyyy-mm-ddTHH:MM:SSZ。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 写入量（单位：字节）
	WriteSize *float64 `json:"WriteSize,omitnil,omitempty" name:"WriteSize"`

	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 时移天数。
	StorageDays *int64 `json:"StorageDays,omitnil,omitempty" name:"StorageDays"`
}

type TimeValue struct {
	// UTC 时间，时间格式：yyyy-mm-ddTHH:MM:SSZ。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 数值。
	Num *uint64 `json:"Num,omitnil,omitempty" name:"Num"`
}

type TranscodeDetailInfo struct {
	// 流名称。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// 开始时间（北京时间），格式：yyyy-mm-dd HH:MM。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（北京时间），格式：yyyy-mm-dd HH:MM。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 转码时长，单位：分钟。
	// 注意：因推流过程中可能有中断重推情况，此处时长为真实转码时长累加值，并非结束时间和开始时间的间隔。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 编码方式，带模块，
	// 示例：
	// liveprocessor_ACC:  音频转码，
	// liveprocessor_H264：直播转码-H264，
	// liveprocessor_H265： 直播转码-H265，
	// topspeed_H264：极速高清-H264，
	// topspeed_H265：极速高清-H265。
	ModuleCodec *string `json:"ModuleCodec,omitnil,omitempty" name:"ModuleCodec"`

	// 码率。
	Bitrate *uint64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// 类型，包含：转码（Transcode），混流（MixStream），水印（WaterMark），快直播（Webrtc）。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 推流域名。
	PushDomain *string `json:"PushDomain,omitnil,omitempty" name:"PushDomain"`

	// 分辨率。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 地域：
	// Mainland：国内。
	// Overseas：海外。
	MainlandOrOversea *string `json:"MainlandOrOversea,omitnil,omitempty" name:"MainlandOrOversea"`
}

type TranscodeTaskNum struct {
	// 时间点。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 码率。
	CodeRate *uint64 `json:"CodeRate,omitnil,omitempty" name:"CodeRate"`

	// 任务数。
	Num *uint64 `json:"Num,omitnil,omitempty" name:"Num"`
}

type TranscodeTotalInfo struct {
	// 时间点，
	// 使用UTC格式时间，
	// 例如：2019-01-08T10:00:00Z。
	// 注意：北京时间值为 UTC 时间值 + 8 小时，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#I)。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 转码时长，单位：分钟。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 编码方式，带模块，
	// 示例：
	// liveprocessor_H264 =》直播转码-H264，
	// liveprocessor_H265 =》 直播转码-H265，
	// topspeed_H264 =》极速高清-H264，
	// topspeed_H265 =》极速高清-H265。
	ModuleCodec *string `json:"ModuleCodec,omitnil,omitempty" name:"ModuleCodec"`

	// 分辨率，
	// 示例：540*480。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`
}

type TransitionTypeInfo struct {
	// 转场名称
	TransitionType *string `json:"TransitionType,omitnil,omitempty" name:"TransitionType"`

	// 素材url
	SourceUrl *string `json:"SourceUrl,omitnil,omitempty" name:"SourceUrl"`

	// 转场的下标，可用来排序，从1开始递增
	Index *uint64 `json:"Index,omitnil,omitempty" name:"Index"`
}

// Predefined struct for user
type UnBindLiveDomainCertRequestParams struct {
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 枚举值：
	// gray: 解绑灰度规则
	// formal(默认): 解绑正式规则
	// 
	// 不传则为formal
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type UnBindLiveDomainCertRequest struct {
	*tchttp.BaseRequest
	
	// 播放域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 枚举值：
	// gray: 解绑灰度规则
	// formal(默认): 解绑正式规则
	// 
	// 不传则为formal
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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

// Predefined struct for user
type UnBindLiveDomainCertResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnBindLiveDomainCertResponse struct {
	*tchttp.BaseResponse
	Response *UnBindLiveDomainCertResponseParams `json:"Response"`
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

// Predefined struct for user
type UpdateLiveWatermarkRequestParams struct {
	// 水印 ID。
	// 在添加水印接口 [AddLiveWatermark](/document/product/267/30154) 调用返回值中获取水印 ID。
	WatermarkId *int64 `json:"WatermarkId,omitnil,omitempty" name:"WatermarkId"`

	// 水印图片 URL。
	// URL中禁止包含的字符：
	//  ;(){}$>`#"\'|
	PictureUrl *string `json:"PictureUrl,omitnil,omitempty" name:"PictureUrl"`

	// 显示位置，X轴偏移，单位是百分比，默认 0。
	XPosition *int64 `json:"XPosition,omitnil,omitempty" name:"XPosition"`

	// 显示位置，Y轴偏移，单位是百分比，默认 0。
	YPosition *int64 `json:"YPosition,omitnil,omitempty" name:"YPosition"`

	// 水印名称。
	// 最长30字节。
	WatermarkName *string `json:"WatermarkName,omitnil,omitempty" name:"WatermarkName"`

	// 水印宽度，占直播原始画面宽度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印高度，占直播原始画面宽度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始高度。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 背景水印宽度。默认宽度1920。
	BackgroundWidth *int64 `json:"BackgroundWidth,omitnil,omitempty" name:"BackgroundWidth"`

	// 背景水印高度。默认高度1080。
	BackgroundHeight *int64 `json:"BackgroundHeight,omitnil,omitempty" name:"BackgroundHeight"`
}

type UpdateLiveWatermarkRequest struct {
	*tchttp.BaseRequest
	
	// 水印 ID。
	// 在添加水印接口 [AddLiveWatermark](/document/product/267/30154) 调用返回值中获取水印 ID。
	WatermarkId *int64 `json:"WatermarkId,omitnil,omitempty" name:"WatermarkId"`

	// 水印图片 URL。
	// URL中禁止包含的字符：
	//  ;(){}$>`#"\'|
	PictureUrl *string `json:"PictureUrl,omitnil,omitempty" name:"PictureUrl"`

	// 显示位置，X轴偏移，单位是百分比，默认 0。
	XPosition *int64 `json:"XPosition,omitnil,omitempty" name:"XPosition"`

	// 显示位置，Y轴偏移，单位是百分比，默认 0。
	YPosition *int64 `json:"YPosition,omitnil,omitempty" name:"YPosition"`

	// 水印名称。
	// 最长30字节。
	WatermarkName *string `json:"WatermarkName,omitnil,omitempty" name:"WatermarkName"`

	// 水印宽度，占直播原始画面宽度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印高度，占直播原始画面宽度百分比，建议高宽只设置一项，另外一项会自适应缩放，避免变形。默认原始高度。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 背景水印宽度。默认宽度1920。
	BackgroundWidth *int64 `json:"BackgroundWidth,omitnil,omitempty" name:"BackgroundWidth"`

	// 背景水印高度。默认高度1080。
	BackgroundHeight *int64 `json:"BackgroundHeight,omitnil,omitempty" name:"BackgroundHeight"`
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
	delete(f, "BackgroundWidth")
	delete(f, "BackgroundHeight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateLiveWatermarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateLiveWatermarkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateLiveWatermarkResponse struct {
	*tchttp.BaseResponse
	Response *UpdateLiveWatermarkResponseParams `json:"Response"`
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
	WatermarkId *int64 `json:"WatermarkId,omitnil,omitempty" name:"WatermarkId"`

	// 水印图片 URL。
	PictureUrl *string `json:"PictureUrl,omitnil,omitempty" name:"PictureUrl"`

	// 显示位置，X 轴偏移。
	XPosition *int64 `json:"XPosition,omitnil,omitempty" name:"XPosition"`

	// 显示位置，Y 轴偏移。
	YPosition *int64 `json:"YPosition,omitnil,omitempty" name:"YPosition"`

	// 水印名称。
	WatermarkName *string `json:"WatermarkName,omitnil,omitempty" name:"WatermarkName"`

	// 当前状态。0：未使用，1:使用中。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 添加时间。
	// 注：此字段为北京时间（UTC+8时区）。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 水印宽。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印高。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 背景水印宽。
	BackgroundWidth *int64 `json:"BackgroundWidth,omitnil,omitempty" name:"BackgroundWidth"`

	// 背景水印高。
	BackgroundHeight *int64 `json:"BackgroundHeight,omitnil,omitempty" name:"BackgroundHeight"`
}

type XP2PDetailInfo struct {
	// CDN流量。
	CdnBytes *int64 `json:"CdnBytes,omitnil,omitempty" name:"CdnBytes"`

	// P2P流量。
	P2pBytes *int64 `json:"P2pBytes,omitnil,omitempty" name:"P2pBytes"`

	// 卡播人数。
	StuckPeople *int64 `json:"StuckPeople,omitnil,omitempty" name:"StuckPeople"`

	// 卡播次数。
	StuckTimes *int64 `json:"StuckTimes,omitnil,omitempty" name:"StuckTimes"`

	// 在线人数。
	OnlinePeople *int64 `json:"OnlinePeople,omitnil,omitempty" name:"OnlinePeople"`

	// 起播请求次数
	Request *int64 `json:"Request,omitnil,omitempty" name:"Request"`

	// 起播成功次数
	RequestSuccess *int64 `json:"RequestSuccess,omitnil,omitempty" name:"RequestSuccess"`

	// 时间，一分钟粒度，utc格式：yyyy-mm-ddTHH:MM:SSZ，参考https://cloud.tencent.com/document/product/266/11732#I。。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 类型，分live和vod两种。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 流ID。
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// AppId。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}