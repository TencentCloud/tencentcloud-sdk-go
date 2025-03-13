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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type BeautifyPicRequestParams struct {
	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。 
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 美白程度，取值范围[0,100]。0不美白，100代表最高程度。默认值30。
	Whitening *uint64 `json:"Whitening,omitnil,omitempty" name:"Whitening"`

	// 磨皮程度，取值范围[0,100]。0不磨皮，100代表最高程度。默认值10。
	Smoothing *uint64 `json:"Smoothing,omitnil,omitempty" name:"Smoothing"`

	// 瘦脸程度，取值范围[0,100]。0不瘦脸，100代表最高程度。默认值70。
	FaceLifting *uint64 `json:"FaceLifting,omitnil,omitempty" name:"FaceLifting"`

	// 大眼程度，取值范围[0,100]。0不大眼，100代表最高程度。默认值70。
	EyeEnlarging *uint64 `json:"EyeEnlarging,omitnil,omitempty" name:"EyeEnlarging"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
}

type BeautifyPicRequest struct {
	*tchttp.BaseRequest
	
	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。 
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 美白程度，取值范围[0,100]。0不美白，100代表最高程度。默认值30。
	Whitening *uint64 `json:"Whitening,omitnil,omitempty" name:"Whitening"`

	// 磨皮程度，取值范围[0,100]。0不磨皮，100代表最高程度。默认值10。
	Smoothing *uint64 `json:"Smoothing,omitnil,omitempty" name:"Smoothing"`

	// 瘦脸程度，取值范围[0,100]。0不瘦脸，100代表最高程度。默认值70。
	FaceLifting *uint64 `json:"FaceLifting,omitnil,omitempty" name:"FaceLifting"`

	// 大眼程度，取值范围[0,100]。0不大眼，100代表最高程度。默认值70。
	EyeEnlarging *uint64 `json:"EyeEnlarging,omitnil,omitempty" name:"EyeEnlarging"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
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
	ResultImage *string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultUrl *string `json:"ResultUrl,omitnil,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateModelRequestParams struct {
	// 图片base64数据，用于试唇色，要求必须是LUT 格式的cube文件转换成512*512的PNG图片。查看 [LUT文件的使用说明](https://cloud.tencent.com/document/product/1172/41701)。了解 [cube文件转png图片小工具](http://yyb.gtimg.com/aiplat/static/qcloud-cube-to-png.html)。
	LUTFile *string `json:"LUTFile,omitnil,omitempty" name:"LUTFile"`

	// 文件描述信息，可用于备注。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateModelRequest struct {
	*tchttp.BaseRequest
	
	// 图片base64数据，用于试唇色，要求必须是LUT 格式的cube文件转换成512*512的PNG图片。查看 [LUT文件的使用说明](https://cloud.tencent.com/document/product/1172/41701)。了解 [cube文件转png图片小工具](http://yyb.gtimg.com/aiplat/static/qcloud-cube-to-png.html)。
	LUTFile *string `json:"LUTFile,omitnil,omitempty" name:"LUTFile"`

	// 文件描述信息，可用于备注。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

type DeleteModelRequest struct {
	*tchttp.BaseRequest
	
	// 素材ID。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 人脸框左上角纵坐标。
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 人脸框宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 人脸框高度。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

// Predefined struct for user
type GetModelListRequestParams struct {
	// 起始序号，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetModelListRequest struct {
	*tchttp.BaseRequest
	
	// 起始序号，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	ModelIdNum *int64 `json:"ModelIdNum,omitnil,omitempty" name:"ModelIdNum"`

	// 素材数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInfos []*ModelInfo `json:"ModelInfos,omitnil,omitempty" name:"ModelInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RGBA *RGBAInfo `json:"RGBA,omitnil,omitempty" name:"RGBA"`

	// 使用已注册的 LUT 文件试唇色。  
	// ModelId 和 RGBA 两个参数只需提供一个，若都提供只使用 ModelId。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 人脸框位置。若不输入则选择 Image 或 Url 中面积最大的人脸。  
	// 您可以通过 [人脸检测与分析](https://cloud.tencent.com/document/api/867/32800)  接口获取人脸框位置信息。
	FaceRect *FaceRect `json:"FaceRect,omitnil,omitempty" name:"FaceRect"`

	// 涂妆浓淡[0,100]。建议取值50。本参数仅控制ModelId对应的涂妆浓淡。
	ModelAlpha *int64 `json:"ModelAlpha,omitnil,omitempty" name:"ModelAlpha"`
}

type ModelInfo struct {
	// 唇色素材ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 唇色素材 url 。 LUT 文件 url 5分钟有效。
	LUTFileUrl *string `json:"LUTFileUrl,omitnil,omitempty" name:"LUTFileUrl"`

	// 文件描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type RGBAInfo struct {
	// R通道数值。[0,255]。
	R *int64 `json:"R,omitnil,omitempty" name:"R"`

	// G通道数值。[0,255]。
	G *int64 `json:"G,omitnil,omitempty" name:"G"`

	// B通道数值。[0,255]。
	B *int64 `json:"B,omitnil,omitempty" name:"B"`

	// A通道数值。[0,100]。建议取值50。
	A *int64 `json:"A,omitnil,omitempty" name:"A"`
}

// Predefined struct for user
type StyleImageProRequestParams struct {
	// 滤镜类型，取值如下： 
	// 1.白茶1；2 白皙；3.初夏；4.东京；5.告白；6.暖阳；7.蔷薇；8.清澄；9.清透；10.甜薄荷；11.默认；12.心动；13.哑灰；14.樱桃布丁；15.自然；16.清逸1；17.黑白；18.水果；19.爱情；20.冬日；21.相片；22.夏日；23.香氛；24.魅惑；25.悸动；26.沙滩；27.街拍；28.甜美；29.初吻；30.午后；31.活力；32.朦胧；33.悦动；34.时尚；35.气泡；36.柠檬；37.棉花糖；38.小溪；39.丽人；40.咖啡；41.嫩芽；42.热情；43.渐暖；44.早餐；45.白茶2；46.白嫩；47.圣代；48.森林；49.冲浪；50.奶咖；51.清澈；52.微风；53.日落；54.水光；55.日系；56.星光；57.阳光；58.落叶；59.生机；60.甜心；61.清逸2；62.春意；63.罗马；64.青涩；65.清风；66.暖心；67.海水；68.神秘；69.旧调1；70.旧调2；71.雪顶；72.日光；73.浮云；74.流彩；75.胶片；76.回味；77.奶酪；78.蝴蝶。
	FilterType *int64 `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。  
	// 支持PNG、JPG、JPEG、BMP 等图片格式，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 滤镜效果，取值[0,100]，0表示无效果，100表示满滤镜效果。默认值为80。
	FilterDegree *int64 `json:"FilterDegree,omitnil,omitempty" name:"FilterDegree"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。默认为base64。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
}

type StyleImageProRequest struct {
	*tchttp.BaseRequest
	
	// 滤镜类型，取值如下： 
	// 1.白茶1；2 白皙；3.初夏；4.东京；5.告白；6.暖阳；7.蔷薇；8.清澄；9.清透；10.甜薄荷；11.默认；12.心动；13.哑灰；14.樱桃布丁；15.自然；16.清逸1；17.黑白；18.水果；19.爱情；20.冬日；21.相片；22.夏日；23.香氛；24.魅惑；25.悸动；26.沙滩；27.街拍；28.甜美；29.初吻；30.午后；31.活力；32.朦胧；33.悦动；34.时尚；35.气泡；36.柠檬；37.棉花糖；38.小溪；39.丽人；40.咖啡；41.嫩芽；42.热情；43.渐暖；44.早餐；45.白茶2；46.白嫩；47.圣代；48.森林；49.冲浪；50.奶咖；51.清澈；52.微风；53.日落；54.水光；55.日系；56.星光；57.阳光；58.落叶；59.生机；60.甜心；61.清逸2；62.春意；63.罗马；64.青涩；65.清风；66.暖心；67.海水；68.神秘；69.旧调1；70.旧调2；71.雪顶；72.日光；73.浮云；74.流彩；75.胶片；76.回味；77.奶酪；78.蝴蝶。
	FilterType *int64 `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。  
	// 支持PNG、JPG、JPEG、BMP 等图片格式，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 滤镜效果，取值[0,100]，0表示无效果，100表示满滤镜效果。默认值为80。
	FilterDegree *int64 `json:"FilterDegree,omitnil,omitempty" name:"FilterDegree"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。默认为base64。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
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
	ResultImage *string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultUrl *string `json:"ResultUrl,omitnil,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	FilterType *int64 `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。  
	// 支持PNG、JPG、JPEG、BMP 等图片格式，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 滤镜效果，取值[0,100]，0表示无效果，100表示满滤镜效果。默认值为80。
	FilterDegree *int64 `json:"FilterDegree,omitnil,omitempty" name:"FilterDegree"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。默认值为base64。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
}

type StyleImageRequest struct {
	*tchttp.BaseRequest
	
	// 滤镜类型，取值如下： 
	// 1.白茶；2 白皙；3.初夏；4.东京；5.告白；6.暖阳；7.蔷薇；8.清澄；9.清透；10.甜薄荷；11.默认；12.心动；13.哑灰；14.樱桃布丁；15.自然；16.清逸；17.黑白；18.水果；19.爱情；20.冬日；21.相片；22.夏日；23.香氛；24.魅惑；25.悸动；26.沙滩；27.街拍；28.甜美；29.初吻；30.午后。
	FilterType *int64 `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。  
	// 支持PNG、JPG、JPEG、BMP 等图片格式，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 滤镜效果，取值[0,100]，0表示无效果，100表示满滤镜效果。默认值为80。
	FilterDegree *int64 `json:"FilterDegree,omitnil,omitempty" name:"FilterDegree"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。默认值为base64。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
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
	ResultImage *string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultUrl *string `json:"ResultUrl,omitnil,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	LipColorInfos []*LipColorInfo `json:"LipColorInfos,omitnil,omitempty" name:"LipColorInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过6M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过6M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
}

type TryLipstickPicRequest struct {
	*tchttp.BaseRequest
	
	// 唇色信息。 
	// 您可以输入最多3个 LipColorInfo 来实现给一张图中的最多3张人脸试唇色。
	LipColorInfos []*LipColorInfo `json:"LipColorInfos,omitnil,omitempty" name:"LipColorInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过6M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过6M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 暂不支持带有alpha透明通道的图片。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
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
	ResultImage *string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。
	ResultUrl *string `json:"ResultUrl,omitnil,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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