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

package v20200324

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateSegmentationTaskRequestParams struct {
	// 需要分割的视频URL，可外网访问。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 背景图片URL。 
	// 可以将视频背景替换为输入的图片。 
	// 如果不输入背景图片，则输出人像区域mask。
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitnil,omitempty" name:"BackgroundImageUrl"`

	// 预留字段，后期用于展示更多识别信息。
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

type CreateSegmentationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 需要分割的视频URL，可外网访问。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 背景图片URL。 
	// 可以将视频背景替换为输入的图片。 
	// 如果不输入背景图片，则输出人像区域mask。
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitnil,omitempty" name:"BackgroundImageUrl"`

	// 预留字段，后期用于展示更多识别信息。
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *CreateSegmentationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSegmentationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VideoUrl")
	delete(f, "BackgroundImageUrl")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSegmentationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSegmentationTaskResponseParams struct {
	// 任务标识ID,可以用与追溯任务状态，查看任务结果
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 预估处理时间，单位为秒
	EstimatedProcessingTime *float64 `json:"EstimatedProcessingTime,omitnil,omitempty" name:"EstimatedProcessingTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSegmentationTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateSegmentationTaskResponseParams `json:"Response"`
}

func (r *CreateSegmentationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSegmentationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSegmentationTaskRequestParams struct {
	// 在提交分割任务成功时返回的任务标识ID。
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

type DescribeSegmentationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 在提交分割任务成功时返回的任务标识ID。
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

func (r *DescribeSegmentationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSegmentationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSegmentationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSegmentationTaskResponseParams struct {
	// 当前任务状态：
	// QUEUING 排队中
	// PROCESSING 处理中
	// FINISHED 处理完成
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 分割后视频URL, 存储于腾讯云COS
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 分割后视频MD5，用于校验
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultVideoMD5 *string `json:"ResultVideoMD5,omitnil,omitempty" name:"ResultVideoMD5"`

	// 视频基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoBasicInformation *VideoBasicInformation `json:"VideoBasicInformation,omitnil,omitempty" name:"VideoBasicInformation"`

	// 分割任务错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSegmentationTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSegmentationTaskResponseParams `json:"Response"`
}

func (r *DescribeSegmentationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSegmentationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageRect struct {
	// 左上角横坐标。
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 左上角纵坐标。
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 人体宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 人体高度。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 分割选项名称。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`
}

// Predefined struct for user
type SegmentCustomizedPortraitPicRequestParams struct {
	// 此参数为分割选项，请根据需要选择自己所想从图片中分割的部分。注意所有选项均为非必选，如未选择则值默认为false, 但是必须要保证多于一个选项的描述为true。
	SegmentationOptions *SegmentationOptions `json:"SegmentationOptions,omitnil,omitempty" name:"SegmentationOptions"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 图片分辨率须小于2000*2000。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片的 Url 。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片分辨率须小于2000*2000 ，图片 base64 编码后大小不可超过5M。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type SegmentCustomizedPortraitPicRequest struct {
	*tchttp.BaseRequest
	
	// 此参数为分割选项，请根据需要选择自己所想从图片中分割的部分。注意所有选项均为非必选，如未选择则值默认为false, 但是必须要保证多于一个选项的描述为true。
	SegmentationOptions *SegmentationOptions `json:"SegmentationOptions,omitnil,omitempty" name:"SegmentationOptions"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 图片分辨率须小于2000*2000。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片的 Url 。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片分辨率须小于2000*2000 ，图片 base64 编码后大小不可超过5M。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

func (r *SegmentCustomizedPortraitPicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SegmentCustomizedPortraitPicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SegmentationOptions")
	delete(f, "Image")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SegmentCustomizedPortraitPicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SegmentCustomizedPortraitPicResponseParams struct {
	// 根据指定标签分割输出的透明背景人像图片的 base64 数据。
	PortraitImage *string `json:"PortraitImage,omitnil,omitempty" name:"PortraitImage"`

	// 指定标签处理后的Mask。一个通过 Base64 编码的文件，解码后文件由 Float 型浮点数组成。这些浮点数代表原图从左上角开始的每一行的每一个像素点，每一个浮点数的值是原图相应像素点位于人体轮廓内的置信度（0-1）转化的灰度值（0-255）
	MaskImage *string `json:"MaskImage,omitnil,omitempty" name:"MaskImage"`

	// 坐标信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageRects []*ImageRect `json:"ImageRects,omitnil,omitempty" name:"ImageRects"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SegmentCustomizedPortraitPicResponse struct {
	*tchttp.BaseResponse
	Response *SegmentCustomizedPortraitPicResponseParams `json:"Response"`
}

func (r *SegmentCustomizedPortraitPicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SegmentCustomizedPortraitPicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SegmentPortraitPicRequestParams struct {
	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 图片分辨率须小于2000*2000。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片的 Url 。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片分辨率须小于2000*2000 ，图片 base64 编码后大小不可超过5M。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 返回图像方式（base64 或 Url ) ，二选一。url有效期为30分钟。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`

	// 适用场景类型。
	// 
	// 取值：GEN/GS。GEN为通用场景模式；GS为绿幕场景模式，针对绿幕场景下的人像分割效果更好。
	// 两种模式选择一种传入，默认为GEN。
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`
}

type SegmentPortraitPicRequest struct {
	*tchttp.BaseRequest
	
	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 图片分辨率须小于2000*2000。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片的 Url 。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片分辨率须小于2000*2000 ，图片 base64 编码后大小不可超过5M。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 返回图像方式（base64 或 Url ) ，二选一。url有效期为30分钟。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`

	// 适用场景类型。
	// 
	// 取值：GEN/GS。GEN为通用场景模式；GS为绿幕场景模式，针对绿幕场景下的人像分割效果更好。
	// 两种模式选择一种传入，默认为GEN。
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`
}

func (r *SegmentPortraitPicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SegmentPortraitPicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "RspImgType")
	delete(f, "Scene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SegmentPortraitPicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SegmentPortraitPicResponseParams struct {
	// 处理后的图片 base64 数据，透明背景图。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultImage *string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// 一个通过 base64 编码的文件，解码后文件由 Float 型浮点数组成。这些浮点数代表原图从左上角开始的每一行的每一个像素点，每一个浮点数的值是原图相应像素点位于人体轮廓内的置信度（0-1）转化的灰度值（0-255）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultMask *string `json:"ResultMask,omitnil,omitempty" name:"ResultMask"`

	// 图片是否存在前景。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasForeground *bool `json:"HasForeground,omitnil,omitempty" name:"HasForeground"`

	// 支持将处理过的图片 base64 数据，透明背景图以Url的形式返回值，Url有效期为30分钟。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultImageUrl *string `json:"ResultImageUrl,omitnil,omitempty" name:"ResultImageUrl"`

	// 一个通过 base64 编码的文件，解码后文件由 Float 型浮点数组成。支持以Url形式的返回值；Url有效期为30分钟。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultMaskUrl *string `json:"ResultMaskUrl,omitnil,omitempty" name:"ResultMaskUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SegmentPortraitPicResponse struct {
	*tchttp.BaseResponse
	Response *SegmentPortraitPicResponseParams `json:"Response"`
}

func (r *SegmentPortraitPicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SegmentPortraitPicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SegmentationOptions struct {
	// 分割选项-背景
	Background *bool `json:"Background,omitnil,omitempty" name:"Background"`

	// 分割选项-头发
	Hair *bool `json:"Hair,omitnil,omitempty" name:"Hair"`

	// 分割选项-左眉
	LeftEyebrow *bool `json:"LeftEyebrow,omitnil,omitempty" name:"LeftEyebrow"`

	// 分割选项-右眉
	RightEyebrow *bool `json:"RightEyebrow,omitnil,omitempty" name:"RightEyebrow"`

	// 分割选项-左眼
	LeftEye *bool `json:"LeftEye,omitnil,omitempty" name:"LeftEye"`

	// 分割选项-右眼
	RightEye *bool `json:"RightEye,omitnil,omitempty" name:"RightEye"`

	// 分割选项-鼻子
	Nose *bool `json:"Nose,omitnil,omitempty" name:"Nose"`

	// 分割选项-上唇
	UpperLip *bool `json:"UpperLip,omitnil,omitempty" name:"UpperLip"`

	// 分割选项-下唇
	LowerLip *bool `json:"LowerLip,omitnil,omitempty" name:"LowerLip"`

	// 分割选项-牙齿
	Tooth *bool `json:"Tooth,omitnil,omitempty" name:"Tooth"`

	// 分割选项-口腔（不包含牙齿）
	Mouth *bool `json:"Mouth,omitnil,omitempty" name:"Mouth"`

	// 分割选项-左耳
	LeftEar *bool `json:"LeftEar,omitnil,omitempty" name:"LeftEar"`

	// 分割选项-右耳
	RightEar *bool `json:"RightEar,omitnil,omitempty" name:"RightEar"`

	// 分割选项-面部(不包含眼、耳、口、鼻等五官及头发。)
	Face *bool `json:"Face,omitnil,omitempty" name:"Face"`

	// 复合分割选项-头部(包含所有的头部元素，相关装饰除外)
	Head *bool `json:"Head,omitnil,omitempty" name:"Head"`

	// 分割选项-身体（包含脖子）
	Body *bool `json:"Body,omitnil,omitempty" name:"Body"`

	// 分割选项-帽子
	Hat *bool `json:"Hat,omitnil,omitempty" name:"Hat"`

	// 分割选项-头饰
	Headdress *bool `json:"Headdress,omitnil,omitempty" name:"Headdress"`

	// 分割选项-耳环
	Earrings *bool `json:"Earrings,omitnil,omitempty" name:"Earrings"`

	// 分割选项-项链
	Necklace *bool `json:"Necklace,omitnil,omitempty" name:"Necklace"`

	// 分割选项-随身物品（ 例如伞、包、手机等。 ）
	Belongings *bool `json:"Belongings,omitnil,omitempty" name:"Belongings"`
}

// Predefined struct for user
type TerminateSegmentationTaskRequestParams struct {
	// 在提交分割任务成功时返回的任务标识ID。
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

type TerminateSegmentationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 在提交分割任务成功时返回的任务标识ID。
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

func (r *TerminateSegmentationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateSegmentationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateSegmentationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateSegmentationTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateSegmentationTaskResponse struct {
	*tchttp.BaseResponse
	Response *TerminateSegmentationTaskResponseParams `json:"Response"`
}

func (r *TerminateSegmentationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateSegmentationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VideoBasicInformation struct {
	// 视频宽度
	FrameWidth *int64 `json:"FrameWidth,omitnil,omitempty" name:"FrameWidth"`

	// 视频高度
	FrameHeight *int64 `json:"FrameHeight,omitnil,omitempty" name:"FrameHeight"`

	// 视频帧速率(FPS)
	FramesPerSecond *int64 `json:"FramesPerSecond,omitnil,omitempty" name:"FramesPerSecond"`

	// 视频时长
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 视频帧数
	TotalFrames *int64 `json:"TotalFrames,omitnil,omitempty" name:"TotalFrames"`
}