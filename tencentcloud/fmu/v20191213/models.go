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

package v20191213

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type BeautifyPicRequestParams struct {
	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。 
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 美白程度，取值范围[0,100]。0不美白，100代表最高程度。默认值30。
	Whitening *uint64 `json:"Whitening,omitempty" name:"Whitening"`

	// 磨皮程度，取值范围[0,100]。0不磨皮，100代表最高程度。默认值10。
	Smoothing *uint64 `json:"Smoothing,omitempty" name:"Smoothing"`

	// 瘦脸程度，取值范围[0,100]。0不瘦脸，100代表最高程度。默认值70。
	FaceLifting *uint64 `json:"FaceLifting,omitempty" name:"FaceLifting"`

	// 大眼程度，取值范围[0,100]。0不大眼，100代表最高程度。默认值70。
	EyeEnlarging *uint64 `json:"EyeEnlarging,omitempty" name:"EyeEnlarging"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

type BeautifyPicRequest struct {
	*tchttp.BaseRequest
	
	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。 
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 美白程度，取值范围[0,100]。0不美白，100代表最高程度。默认值30。
	Whitening *uint64 `json:"Whitening,omitempty" name:"Whitening"`

	// 磨皮程度，取值范围[0,100]。0不磨皮，100代表最高程度。默认值10。
	Smoothing *uint64 `json:"Smoothing,omitempty" name:"Smoothing"`

	// 瘦脸程度，取值范围[0,100]。0不瘦脸，100代表最高程度。默认值70。
	FaceLifting *uint64 `json:"FaceLifting,omitempty" name:"FaceLifting"`

	// 大眼程度，取值范围[0,100]。0不大眼，100代表最高程度。默认值70。
	EyeEnlarging *uint64 `json:"EyeEnlarging,omitempty" name:"EyeEnlarging"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

func (r *BeautifyPicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BeautifyPicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "Whitening")
	delete(f, "Smoothing")
	delete(f, "FaceLifting")
	delete(f, "EyeEnlarging")
	delete(f, "RspImgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BeautifyPicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BeautifyPicResponseParams struct {
	// RspImgType 为 base64 时，返回处理后的图片 base64 数据。默认返回base64
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BeautifyPicResponse struct {
	*tchttp.BaseResponse
	Response *BeautifyPicResponseParams `json:"Response"`
}

func (r *BeautifyPicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BeautifyPicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BeautifyVideoOutput struct {
	// 视频美颜输出的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// 视频美颜输出的视频MD5，用于校验
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoMD5 *string `json:"VideoMD5,omitempty" name:"VideoMD5"`

	// 美颜输出的视频封面图base64字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverImage *string `json:"CoverImage,omitempty" name:"CoverImage"`

	// 视频宽度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 视频高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 每秒传输帧数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fps *float64 `json:"Fps,omitempty" name:"Fps"`

	// 视频播放时长，单位为秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	DurationInSec *float64 `json:"DurationInSec,omitempty" name:"DurationInSec"`
}

// Predefined struct for user
type BeautifyVideoRequestParams struct {
	// 视频url地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 美颜参数 - 美白、平滑、大眼和瘦脸。参数值范围[0, 100]。参数值为0，则不做美颜。参数默认值为0。目前默认取数组第一个元素是对所有人脸美颜。
	BeautyParam []*BeautyParam `json:"BeautyParam,omitempty" name:"BeautyParam"`

	// 目前只支持mp4
	OutputVideoType *string `json:"OutputVideoType,omitempty" name:"OutputVideoType"`
}

type BeautifyVideoRequest struct {
	*tchttp.BaseRequest
	
	// 视频url地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 美颜参数 - 美白、平滑、大眼和瘦脸。参数值范围[0, 100]。参数值为0，则不做美颜。参数默认值为0。目前默认取数组第一个元素是对所有人脸美颜。
	BeautyParam []*BeautyParam `json:"BeautyParam,omitempty" name:"BeautyParam"`

	// 目前只支持mp4
	OutputVideoType *string `json:"OutputVideoType,omitempty" name:"OutputVideoType"`
}

func (r *BeautifyVideoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BeautifyVideoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	delete(f, "BeautyParam")
	delete(f, "OutputVideoType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BeautifyVideoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BeautifyVideoResponseParams struct {
	// 视频美颜任务的Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 预估处理时间，粒度为秒
	EstimatedProcessTime *int64 `json:"EstimatedProcessTime,omitempty" name:"EstimatedProcessTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BeautifyVideoResponse struct {
	*tchttp.BaseResponse
	Response *BeautifyVideoResponseParams `json:"Response"`
}

func (r *BeautifyVideoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BeautifyVideoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BeautyParam struct {
	// 美白程度，取值范围[0,100]。0不美白，100代表最高程度。默认值30。
	WhitenLevel *uint64 `json:"WhitenLevel,omitempty" name:"WhitenLevel"`

	// 磨皮程度，取值范围[0,100]。0不磨皮，100代表最高程度。默认值30。
	SmoothingLevel *uint64 `json:"SmoothingLevel,omitempty" name:"SmoothingLevel"`

	// 大眼程度，取值范围[0,100]。0不大眼，100代表最高程度。默认值70。
	EyeEnlargeLevel *uint64 `json:"EyeEnlargeLevel,omitempty" name:"EyeEnlargeLevel"`

	// 瘦脸程度，取值范围[0,100]。0不瘦脸，100代表最高程度。默认值70。
	FaceShrinkLevel *uint64 `json:"FaceShrinkLevel,omitempty" name:"FaceShrinkLevel"`
}

// Predefined struct for user
type CancelBeautifyVideoJobRequestParams struct {
	// 美颜视频的Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type CancelBeautifyVideoJobRequest struct {
	*tchttp.BaseRequest
	
	// 美颜视频的Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *CancelBeautifyVideoJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelBeautifyVideoJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelBeautifyVideoJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelBeautifyVideoJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelBeautifyVideoJobResponse struct {
	*tchttp.BaseResponse
	Response *CancelBeautifyVideoJobResponseParams `json:"Response"`
}

func (r *CancelBeautifyVideoJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelBeautifyVideoJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelRequestParams struct {
	// 图片base64数据，用于试唇色，要求必须是LUT 格式的cube文件转换成512*512的PNG图片。查看 [LUT文件的使用说明](https://cloud.tencent.com/document/product/1172/41701)。了解 [cube文件转png图片小工具](http://yyb.gtimg.com/aiplat/static/qcloud-cube-to-png.html)。
	LUTFile *string `json:"LUTFile,omitempty" name:"LUTFile"`

	// 文件描述信息，可用于备注。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateModelRequest struct {
	*tchttp.BaseRequest
	
	// 图片base64数据，用于试唇色，要求必须是LUT 格式的cube文件转换成512*512的PNG图片。查看 [LUT文件的使用说明](https://cloud.tencent.com/document/product/1172/41701)。了解 [cube文件转png图片小工具](http://yyb.gtimg.com/aiplat/static/qcloud-cube-to-png.html)。
	LUTFile *string `json:"LUTFile,omitempty" name:"LUTFile"`

	// 文件描述信息，可用于备注。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LUTFile")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelResponseParams struct {
	// 唇色素材ID。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateModelResponse struct {
	*tchttp.BaseResponse
	Response *CreateModelResponseParams `json:"Response"`
}

func (r *CreateModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelRequestParams struct {
	// 素材ID。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`
}

type DeleteModelRequest struct {
	*tchttp.BaseRequest
	
	// 素材ID。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`
}

func (r *DeleteModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteModelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModelResponseParams `json:"Response"`
}

func (r *DeleteModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaceRect struct {
	// 人脸框左上角横坐标。
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸框左上角纵坐标。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸框宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 人脸框高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

// Predefined struct for user
type GetModelListRequestParams struct {
	// 起始序号，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type GetModelListRequest struct {
	*tchttp.BaseRequest
	
	// 起始序号，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetModelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetModelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetModelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetModelListResponseParams struct {
	// 唇色素材总数量。
	ModelIdNum *int64 `json:"ModelIdNum,omitempty" name:"ModelIdNum"`

	// 素材数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInfos []*ModelInfo `json:"ModelInfos,omitempty" name:"ModelInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetModelListResponse struct {
	*tchttp.BaseResponse
	Response *GetModelListResponseParams `json:"Response"`
}

func (r *GetModelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetModelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LipColorInfo struct {
	// 使用RGBA模型试唇色。
	RGBA *RGBAInfo `json:"RGBA,omitempty" name:"RGBA"`

	// 使用已注册的 LUT 文件试唇色。  
	// ModelId 和 RGBA 两个参数只需提供一个，若都提供只使用 ModelId。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 人脸框位置。若不输入则选择 Image 或 Url 中面积最大的人脸。  
	// 您可以通过 [人脸检测与分析](https://cloud.tencent.com/document/api/867/32800)  接口获取人脸框位置信息。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

	// 涂妆浓淡[0,100]。建议取值50。本参数仅控制ModelId对应的涂妆浓淡。
	ModelAlpha *int64 `json:"ModelAlpha,omitempty" name:"ModelAlpha"`
}

type ModelInfo struct {
	// 唇色素材ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 唇色素材 url 。 LUT 文件 url 5分钟有效。
	LUTFileUrl *string `json:"LUTFileUrl,omitempty" name:"LUTFileUrl"`

	// 文件描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`
}

// Predefined struct for user
type QueryBeautifyVideoJobRequestParams struct {
	// 视频美颜Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type QueryBeautifyVideoJobRequest struct {
	*tchttp.BaseRequest
	
	// 视频美颜Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *QueryBeautifyVideoJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBeautifyVideoJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryBeautifyVideoJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryBeautifyVideoJobResponseParams struct {
	// 当前任务状态：排队中、处理中、处理失败或者处理完成
	JobStatus *string `json:"JobStatus,omitempty" name:"JobStatus"`

	// 视频美颜输出的结果信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeautifyVideoOutput *BeautifyVideoOutput `json:"BeautifyVideoOutput,omitempty" name:"BeautifyVideoOutput"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryBeautifyVideoJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryBeautifyVideoJobResponseParams `json:"Response"`
}

func (r *QueryBeautifyVideoJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBeautifyVideoJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RGBAInfo struct {
	// R通道数值。[0,255]。
	R *int64 `json:"R,omitempty" name:"R"`

	// G通道数值。[0,255]。
	G *int64 `json:"G,omitempty" name:"G"`

	// B通道数值。[0,255]。
	B *int64 `json:"B,omitempty" name:"B"`

	// A通道数值。[0,100]。建议取值50。
	A *int64 `json:"A,omitempty" name:"A"`
}

// Predefined struct for user
type StyleImageProRequestParams struct {
	// 滤镜类型，取值如下： 
	// 1.白茶；2 白皙；3.初夏；4.东京；5.告白；6.暖阳；7.蔷薇；8.清澄；9.清透；10.甜薄荷；11.默认；12.心动；13.哑灰；14.樱桃布丁；15.自然；16.清逸；17.黑白；18.水果；19.爱情；20.冬日；21.相片；22.夏日；23.香氛；24.魅惑；25.悸动；26.沙滩；27.街拍；28.甜美；29.初吻；30.午后；31.活力；32.朦胧；33.悦动；34.时尚；35.气泡；36.柠檬；37.棉花糖；38.小溪；39.丽人；40.咖啡；41.嫩芽；42.热情；43.渐暖；44.早餐；45.白茶；46.白嫩；47.圣代；48.森林；49.冲浪；50.奶咖；51.清澈；52.微风；53.日落；54.水光；55.日系；56.星光；57.阳光；58.落叶；59.生机；60.甜心；61.清逸；62.春意；63.罗马；64.青涩；65.清风；66.暖心；67.海水；68.神秘；69.旧调1；70.旧调2；71.雪顶；72.日光；73.浮云；74.流彩；75.胶片；76.回味；77.奶酪；78.蝴蝶。
	FilterType *int64 `json:"FilterType,omitempty" name:"FilterType"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。  
	// 支持PNG、JPG、JPEG、BMP 等图片格式，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 滤镜效果，取值[0,100]，0表示无效果，100表示满滤镜效果。默认值为80。
	FilterDegree *int64 `json:"FilterDegree,omitempty" name:"FilterDegree"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

type StyleImageProRequest struct {
	*tchttp.BaseRequest
	
	// 滤镜类型，取值如下： 
	// 1.白茶；2 白皙；3.初夏；4.东京；5.告白；6.暖阳；7.蔷薇；8.清澄；9.清透；10.甜薄荷；11.默认；12.心动；13.哑灰；14.樱桃布丁；15.自然；16.清逸；17.黑白；18.水果；19.爱情；20.冬日；21.相片；22.夏日；23.香氛；24.魅惑；25.悸动；26.沙滩；27.街拍；28.甜美；29.初吻；30.午后；31.活力；32.朦胧；33.悦动；34.时尚；35.气泡；36.柠檬；37.棉花糖；38.小溪；39.丽人；40.咖啡；41.嫩芽；42.热情；43.渐暖；44.早餐；45.白茶；46.白嫩；47.圣代；48.森林；49.冲浪；50.奶咖；51.清澈；52.微风；53.日落；54.水光；55.日系；56.星光；57.阳光；58.落叶；59.生机；60.甜心；61.清逸；62.春意；63.罗马；64.青涩；65.清风；66.暖心；67.海水；68.神秘；69.旧调1；70.旧调2；71.雪顶；72.日光；73.浮云；74.流彩；75.胶片；76.回味；77.奶酪；78.蝴蝶。
	FilterType *int64 `json:"FilterType,omitempty" name:"FilterType"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。  
	// 支持PNG、JPG、JPEG、BMP 等图片格式，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 滤镜效果，取值[0,100]，0表示无效果，100表示满滤镜效果。默认值为80。
	FilterDegree *int64 `json:"FilterDegree,omitempty" name:"FilterDegree"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

func (r *StyleImageProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StyleImageProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FilterType")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "FilterDegree")
	delete(f, "RspImgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StyleImageProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StyleImageProResponseParams struct {
	// RspImgType 为 base64 时，返回处理后的图片 base64 数据。默认返回base64
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StyleImageProResponse struct {
	*tchttp.BaseResponse
	Response *StyleImageProResponseParams `json:"Response"`
}

func (r *StyleImageProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StyleImageProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StyleImageRequestParams struct {
	// 滤镜类型，取值如下： 
	// 1.白茶；2 白皙；3.初夏；4.东京；5.告白；6.暖阳；7.蔷薇；8.清澄；9.清透；10.甜薄荷；11.默认；12.心动；13.哑灰；14.樱桃布丁；15.自然；16.清逸；17.黑白；18.水果；19.爱情；20.冬日；21.相片；22.夏日；23.香氛；24.魅惑；25.悸动；26.沙滩；27.街拍；28.甜美；29.初吻；30.午后。
	FilterType *int64 `json:"FilterType,omitempty" name:"FilterType"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。  
	// 支持PNG、JPG、JPEG、BMP 等图片格式，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 滤镜效果，取值[0,100]，0表示无效果，100表示满滤镜效果。默认值为80。
	FilterDegree *int64 `json:"FilterDegree,omitempty" name:"FilterDegree"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

type StyleImageRequest struct {
	*tchttp.BaseRequest
	
	// 滤镜类型，取值如下： 
	// 1.白茶；2 白皙；3.初夏；4.东京；5.告白；6.暖阳；7.蔷薇；8.清澄；9.清透；10.甜薄荷；11.默认；12.心动；13.哑灰；14.樱桃布丁；15.自然；16.清逸；17.黑白；18.水果；19.爱情；20.冬日；21.相片；22.夏日；23.香氛；24.魅惑；25.悸动；26.沙滩；27.街拍；28.甜美；29.初吻；30.午后。
	FilterType *int64 `json:"FilterType,omitempty" name:"FilterType"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。  
	// 支持PNG、JPG、JPEG、BMP 等图片格式，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 滤镜效果，取值[0,100]，0表示无效果，100表示满滤镜效果。默认值为80。
	FilterDegree *int64 `json:"FilterDegree,omitempty" name:"FilterDegree"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

func (r *StyleImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StyleImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FilterType")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "FilterDegree")
	delete(f, "RspImgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StyleImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StyleImageResponseParams struct {
	// RspImgType 为 base64 时，返回处理后的图片 base64 数据。默认返回base64
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StyleImageResponse struct {
	*tchttp.BaseResponse
	Response *StyleImageResponseParams `json:"Response"`
}

func (r *StyleImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StyleImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TryLipstickPicRequestParams struct {
	// 唇色信息。 
	// 您可以输入最多3个 LipColorInfo 来实现给一张图中的最多3张人脸试唇色。
	LipColorInfos []*LipColorInfo `json:"LipColorInfos,omitempty" name:"LipColorInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过6M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过6M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

type TryLipstickPicRequest struct {
	*tchttp.BaseRequest
	
	// 唇色信息。 
	// 您可以输入最多3个 LipColorInfo 来实现给一张图中的最多3张人脸试唇色。
	LipColorInfos []*LipColorInfo `json:"LipColorInfos,omitempty" name:"LipColorInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过6M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过6M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

func (r *TryLipstickPicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TryLipstickPicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LipColorInfos")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "RspImgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TryLipstickPicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TryLipstickPicResponseParams struct {
	// RspImgType 为 base64 时，返回处理后的图片 base64 数据。默认返回base64
	ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TryLipstickPicResponse struct {
	*tchttp.BaseResponse
	Response *TryLipstickPicResponseParams `json:"Response"`
}

func (r *TryLipstickPicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TryLipstickPicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}