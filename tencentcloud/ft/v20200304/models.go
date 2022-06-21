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

package v20200304

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AgeInfo struct {
	// 变化到的人脸年龄 [10,80]。
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 人脸框位置。若不输入则选择 Image 或 Url 中面积最大的人脸。  
	// 您可以通过 [人脸检测与分析](https://cloud.tencent.com/document/api/867/32800)  接口获取人脸框位置信息。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`
}

// Predefined struct for user
type CancelFaceMorphJobRequestParams struct {
	// 人像渐变任务Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type CancelFaceMorphJobRequest struct {
	*tchttp.BaseRequest
	
	// 人像渐变任务Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *CancelFaceMorphJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelFaceMorphJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelFaceMorphJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelFaceMorphJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelFaceMorphJobResponse struct {
	*tchttp.BaseResponse
	Response *CancelFaceMorphJobResponseParams `json:"Response"`
}

func (r *CancelFaceMorphJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelFaceMorphJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeAgePicRequestParams struct {
	// 人脸变老变年轻信息。 
	// 您可以输入最多3个 AgeInfo 来实现给一张图中的最多3张人脸变老变年轻。
	AgeInfos []*AgeInfo `json:"AgeInfos,omitempty" name:"AgeInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

type ChangeAgePicRequest struct {
	*tchttp.BaseRequest
	
	// 人脸变老变年轻信息。 
	// 您可以输入最多3个 AgeInfo 来实现给一张图中的最多3张人脸变老变年轻。
	AgeInfos []*AgeInfo `json:"AgeInfos,omitempty" name:"AgeInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

func (r *ChangeAgePicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeAgePicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgeInfos")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "RspImgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeAgePicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeAgePicResponseParams struct {
	// RspImgType 为 base64 时，返回处理后的图片 base64 数据。默认返回base64
	ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChangeAgePicResponse struct {
	*tchttp.BaseResponse
	Response *ChangeAgePicResponseParams `json:"Response"`
}

func (r *ChangeAgePicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeAgePicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FaceCartoonPicRequestParams struct {
	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 关闭全图动漫化，传入true（不分大小写）即关闭全图动漫化。
	DisableGlobalEffect *string `json:"DisableGlobalEffect,omitempty" name:"DisableGlobalEffect"`
}

type FaceCartoonPicRequest struct {
	*tchttp.BaseRequest
	
	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 关闭全图动漫化，传入true（不分大小写）即关闭全图动漫化。
	DisableGlobalEffect *string `json:"DisableGlobalEffect,omitempty" name:"DisableGlobalEffect"`
}

func (r *FaceCartoonPicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FaceCartoonPicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "RspImgType")
	delete(f, "DisableGlobalEffect")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FaceCartoonPicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FaceCartoonPicResponseParams struct {
	// 结果图片Base64信息。
	ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。(默认为base64)
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FaceCartoonPicResponse struct {
	*tchttp.BaseResponse
	Response *FaceCartoonPicResponseParams `json:"Response"`
}

func (r *FaceCartoonPicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FaceCartoonPicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaceMorphOutput struct {
	// 人像渐变输出的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	MorphUrl *string `json:"MorphUrl,omitempty" name:"MorphUrl"`

	// 人像渐变输出的结果MD5，用于校验
	// 注意：此字段可能返回 null，表示取不到有效值。
	MorphMd5 *string `json:"MorphMd5,omitempty" name:"MorphMd5"`

	// 人像渐变输出的结果封面图base64字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverImage *string `json:"CoverImage,omitempty" name:"CoverImage"`
}

type FaceRect struct {
	// 人脸框左上角纵坐标。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸框左上角横坐标。
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸框宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 人脸框高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type GenderInfo struct {
	// 选择转换方向，0：男变女，1：女变男。
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 人脸框位置。若不输入则选择 Image 或 Url 中面积最大的人脸。  
	// 您可以通过 [人脸检测与分析](https://cloud.tencent.com/document/api/867/32800)  接口获取人脸框位置信息。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`
}

type GradientInfo struct {
	// 图片的展示时长，即单张图片静止不变的时间。GIF默认每张图片0.7s，视频默认每张图片0.5s。最大取值1s。
	Tempo *float64 `json:"Tempo,omitempty" name:"Tempo"`

	// 人像渐变的最长时间，即单张图片使用渐变特效的时间。 GIF默认值为0.5s，视频默值认为1s。最大取值1s。
	MorphTime *float64 `json:"MorphTime,omitempty" name:"MorphTime"`
}

// Predefined struct for user
type MorphFaceRequestParams struct {
	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 
	// 人员人脸总数量至少2张，不可超过5张。 
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Images []*string `json:"Images,omitempty" name:"Images"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 
	// Url、Image必须提供一个，如果都提供，只使用 Url。图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。 
	// 人员人脸总数量不可超过5张。 
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 人脸渐变参数。可调整每张图片的展示时长、人像渐变的最长时间
	GradientInfos []*GradientInfo `json:"GradientInfos,omitempty" name:"GradientInfos"`

	// 视频帧率，取值[1,25]。默认10
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 视频类型，取值0。目前仅支持MP4格式，默认为MP4格式
	OutputType *int64 `json:"OutputType,omitempty" name:"OutputType"`

	// 视频宽度，取值[128,1280]。默认值720
	OutputWidth *int64 `json:"OutputWidth,omitempty" name:"OutputWidth"`

	// 视频高度，取值[128,1280]。默认值1280
	OutputHeight *int64 `json:"OutputHeight,omitempty" name:"OutputHeight"`
}

type MorphFaceRequest struct {
	*tchttp.BaseRequest
	
	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 
	// 人员人脸总数量至少2张，不可超过5张。 
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Images []*string `json:"Images,omitempty" name:"Images"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 
	// Url、Image必须提供一个，如果都提供，只使用 Url。图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。 
	// 人员人脸总数量不可超过5张。 
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 人脸渐变参数。可调整每张图片的展示时长、人像渐变的最长时间
	GradientInfos []*GradientInfo `json:"GradientInfos,omitempty" name:"GradientInfos"`

	// 视频帧率，取值[1,25]。默认10
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 视频类型，取值0。目前仅支持MP4格式，默认为MP4格式
	OutputType *int64 `json:"OutputType,omitempty" name:"OutputType"`

	// 视频宽度，取值[128,1280]。默认值720
	OutputWidth *int64 `json:"OutputWidth,omitempty" name:"OutputWidth"`

	// 视频高度，取值[128,1280]。默认值1280
	OutputHeight *int64 `json:"OutputHeight,omitempty" name:"OutputHeight"`
}

func (r *MorphFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MorphFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Images")
	delete(f, "Urls")
	delete(f, "GradientInfos")
	delete(f, "Fps")
	delete(f, "OutputType")
	delete(f, "OutputWidth")
	delete(f, "OutputHeight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MorphFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MorphFaceResponseParams struct {
	// 人像渐变任务的Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 预估处理时间，粒度为秒
	EstimatedProcessTime *int64 `json:"EstimatedProcessTime,omitempty" name:"EstimatedProcessTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MorphFaceResponse struct {
	*tchttp.BaseResponse
	Response *MorphFaceResponseParams `json:"Response"`
}

func (r *MorphFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MorphFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFaceMorphJobRequestParams struct {
	// 人像渐变任务Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type QueryFaceMorphJobRequest struct {
	*tchttp.BaseRequest
	
	// 人像渐变任务Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *QueryFaceMorphJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFaceMorphJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFaceMorphJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFaceMorphJobResponseParams struct {
	// 当前任务状态：排队中、处理中、处理失败或者处理完成
	JobStatus *string `json:"JobStatus,omitempty" name:"JobStatus"`

	// 人像渐变输出的结果信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceMorphOutput *FaceMorphOutput `json:"FaceMorphOutput,omitempty" name:"FaceMorphOutput"`

	// 当前任务状态码：1：排队中、3: 处理中、5: 处理失败、7:处理完成
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobStatusCode *int64 `json:"JobStatusCode,omitempty" name:"JobStatusCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFaceMorphJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryFaceMorphJobResponseParams `json:"Response"`
}

func (r *QueryFaceMorphJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFaceMorphJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwapGenderPicRequestParams struct {
	// 人脸转化性别信息。 
	// 您可以输入最多3个 GenderInfo 来实现给一张图中的最多3张人脸转换性别。
	GenderInfos []*GenderInfo `json:"GenderInfos,omitempty" name:"GenderInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

type SwapGenderPicRequest struct {
	*tchttp.BaseRequest
	
	// 人脸转化性别信息。 
	// 您可以输入最多3个 GenderInfo 来实现给一张图中的最多3张人脸转换性别。
	GenderInfos []*GenderInfo `json:"GenderInfos,omitempty" name:"GenderInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

func (r *SwapGenderPicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwapGenderPicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GenderInfos")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "RspImgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwapGenderPicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwapGenderPicResponseParams struct {
	// RspImgType 为 base64 时，返回处理后的图片 base64 数据。默认返回base64
	ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwapGenderPicResponse struct {
	*tchttp.BaseResponse
	Response *SwapGenderPicResponseParams `json:"Response"`
}

func (r *SwapGenderPicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwapGenderPicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}