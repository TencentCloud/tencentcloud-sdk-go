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

package v20220927

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeMaterialListRequestParams struct {
	// 活动Id
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 素材Id
	MaterialId *string `json:"MaterialId,omitnil,omitempty" name:"MaterialId"`

	// 每次拉取条数。
	// 每次拉取素材最多可支持拉取20条素材信息，如果需要拉取全部素材信息，可以分多次请求拉取全部素材信息。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeMaterialListRequest struct {
	*tchttp.BaseRequest
	
	// 活动Id
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 素材Id
	MaterialId *string `json:"MaterialId,omitnil,omitempty" name:"MaterialId"`

	// 每次拉取条数。
	// 每次拉取素材最多可支持拉取20条素材信息，如果需要拉取全部素材信息，可以分多次请求拉取全部素材信息。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeMaterialListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaterialListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActivityId")
	delete(f, "MaterialId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaterialListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaterialListResponseParams struct {
	// 素材列表数据
	MaterialInfos []*PublicMaterialInfos `json:"MaterialInfos,omitnil,omitempty" name:"MaterialInfos"`

	// 素材条数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMaterialListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaterialListResponseParams `json:"Response"`
}

func (r *DescribeMaterialListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaterialListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaceInfo struct {
	// 人脸框的横坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 人脸框的纵坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 人脸框的宽度
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 人脸框的高度
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
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
type FuseFaceRequestParams struct {
	// 活动 ID，请在<a href="https://console.cloud.tencent.com/facefusion" target="_blank">人脸融合控制台</a>查看。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 素材 ID，请在<a href="https://console.cloud.tencent.com/facefusion" target="_blank">人脸融合控制台</a>查看。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为7天。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`

	// 用户人脸图片、素材模板图的人脸位置信息。不能超过6个。
	// ●图片分辨率限制：图片中面部尺寸大于34 * 34；图片尺寸大于64 * 64，小于4096*4096（单边限制）。
	// ●图片大小限制：base64 编码后大小不可超过5M， url不超过10M。
	// ●支持图片格式：支持jpg或png
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil,omitempty" name:"MergeInfos"`

	// 脸型融合比例，数值越高，融合后的脸型越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中脸型参数数值。（换脸版算法暂不支持此参数调整）
	FuseProfileDegree *int64 `json:"FuseProfileDegree,omitnil,omitempty" name:"FuseProfileDegree"`

	// 五官融合比例，数值越高，融合后的五官越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中五官参数数值。（换脸版算法暂不支持此参数调整）
	FuseFaceDegree *int64 `json:"FuseFaceDegree,omitnil,omitempty" name:"FuseFaceDegree"`

	// 为融合结果图添加合成标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了人脸融合技术，是AI合成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在融合结果图右下角添加“本图片为AI合成图片”字样，您可根据自身需要替换为其他的Logo图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 融合参数。
	FuseParam *FuseParam `json:"FuseParam,omitnil,omitempty" name:"FuseParam"`
}

type FuseFaceRequest struct {
	*tchttp.BaseRequest
	
	// 活动 ID，请在<a href="https://console.cloud.tencent.com/facefusion" target="_blank">人脸融合控制台</a>查看。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 素材 ID，请在<a href="https://console.cloud.tencent.com/facefusion" target="_blank">人脸融合控制台</a>查看。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为7天。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`

	// 用户人脸图片、素材模板图的人脸位置信息。不能超过6个。
	// ●图片分辨率限制：图片中面部尺寸大于34 * 34；图片尺寸大于64 * 64，小于4096*4096（单边限制）。
	// ●图片大小限制：base64 编码后大小不可超过5M， url不超过10M。
	// ●支持图片格式：支持jpg或png
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil,omitempty" name:"MergeInfos"`

	// 脸型融合比例，数值越高，融合后的脸型越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中脸型参数数值。（换脸版算法暂不支持此参数调整）
	FuseProfileDegree *int64 `json:"FuseProfileDegree,omitnil,omitempty" name:"FuseProfileDegree"`

	// 五官融合比例，数值越高，融合后的五官越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中五官参数数值。（换脸版算法暂不支持此参数调整）
	FuseFaceDegree *int64 `json:"FuseFaceDegree,omitnil,omitempty" name:"FuseFaceDegree"`

	// 为融合结果图添加合成标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了人脸融合技术，是AI合成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在融合结果图右下角添加“本图片为AI合成图片”字样，您可根据自身需要替换为其他的Logo图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 融合参数。
	FuseParam *FuseParam `json:"FuseParam,omitnil,omitempty" name:"FuseParam"`
}

func (r *FuseFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FuseFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ModelId")
	delete(f, "RspImgType")
	delete(f, "MergeInfos")
	delete(f, "FuseProfileDegree")
	delete(f, "FuseFaceDegree")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "FuseParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FuseFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FuseFaceResponseParams struct {
	// RspImgType 为 url 时，返回结果的 url（有效期7天）， RspImgType 为 base64 时返回 base64 数据。
	FusedImage *string `json:"FusedImage,omitnil,omitempty" name:"FusedImage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FuseFaceResponse struct {
	*tchttp.BaseResponse
	Response *FuseFaceResponseParams `json:"Response"`
}

func (r *FuseFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FuseFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FuseFaceUltraRequestParams struct {
	// 返回融合结果图片方式（url 或 base64) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`

	// 用户人脸图片、素材模板图的人脸位置信息。主要用于素材模版中人脸以及用作融合的用户人脸相关信息，两种人脸都需要提供人脸图片，可选择提供人脸框位置，具体见MergeInfo说明 
	// 目前最多支持融合模板图片中的6张人脸。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil,omitempty" name:"MergeInfos"`

	// 素材模版图片的url地址。
	// ●base64 和 url 必须提供一个，如果都提供以 url 为准。
	// ●图片分辨率限制：图片中面部尺寸大于34 * 34；图片尺寸大于64 * 64，小于8000 * 8000（单边限制）。
	// ●图片大小限制：base64 编码后大小不可超过10M， url不超过20M。
	// ●图片格式：支持jpg或png
	ModelUrl *string `json:"ModelUrl,omitnil,omitempty" name:"ModelUrl"`

	// 素材模版图片base64数据。
	// ●base64 和 url 必须提供一个，如果都提供以 url 为准。
	// ●素材图片限制：图片中面部尺寸大于34 * 34；图片尺寸大于64 * 64，小于8000*8000（单边限制）。
	// ●图片大小限制：base64 编码后大小不可超过10M， url不超过20M。
	// ●支持图片格式：支持jpg或png
	ModelImage *string `json:"ModelImage,omitnil,omitempty" name:"ModelImage"`

	// 图片人脸融合（专业版）效果参数。可用于设置拉脸、人脸增强、磨皮、牙齿增强、妆容迁移等融合效果参数，生成理想的融合效果。不传默认使用接口推荐值。具体见FusionUltraParam说明
	FusionUltraParam *FusionUltraParam `json:"FusionUltraParam,omitnil,omitempty" name:"FusionUltraParam"`

	// 为融合结果图添加合成标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了人脸融合技术，是AI合成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在融合结果图右下角添加“本图片为AI合成图片”字样，您可根据自身需要替换为其他的Logo图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 融合模型类型参数：默认为1。
	// 图片人脸融合（专业版）针对不同场景，提供多种模型供选择。如您的产品是泛娱乐场景，推荐使用1；如您主要用于影像场景，推荐使用4、5。其他模型类型也可以结合您的产品使用场景进行选择，也许会有意想不到的效果
	// 1：默认泛娱乐场景，画面偏锐。
	// 2：影视级场景，画面偏自然。
	// 3：影视级场景，高分辨率，画面偏自然。
	// 4：影视级场景，高分辦率，高人脸相似度，画面偏自然，可用于证件照等场景。
	// 5：影视级场景，高分辨率，对闭眼和遮挡更友好。
	SwapModelType *int64 `json:"SwapModelType,omitnil,omitempty" name:"SwapModelType"`
}

type FuseFaceUltraRequest struct {
	*tchttp.BaseRequest
	
	// 返回融合结果图片方式（url 或 base64) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`

	// 用户人脸图片、素材模板图的人脸位置信息。主要用于素材模版中人脸以及用作融合的用户人脸相关信息，两种人脸都需要提供人脸图片，可选择提供人脸框位置，具体见MergeInfo说明 
	// 目前最多支持融合模板图片中的6张人脸。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitnil,omitempty" name:"MergeInfos"`

	// 素材模版图片的url地址。
	// ●base64 和 url 必须提供一个，如果都提供以 url 为准。
	// ●图片分辨率限制：图片中面部尺寸大于34 * 34；图片尺寸大于64 * 64，小于8000 * 8000（单边限制）。
	// ●图片大小限制：base64 编码后大小不可超过10M， url不超过20M。
	// ●图片格式：支持jpg或png
	ModelUrl *string `json:"ModelUrl,omitnil,omitempty" name:"ModelUrl"`

	// 素材模版图片base64数据。
	// ●base64 和 url 必须提供一个，如果都提供以 url 为准。
	// ●素材图片限制：图片中面部尺寸大于34 * 34；图片尺寸大于64 * 64，小于8000*8000（单边限制）。
	// ●图片大小限制：base64 编码后大小不可超过10M， url不超过20M。
	// ●支持图片格式：支持jpg或png
	ModelImage *string `json:"ModelImage,omitnil,omitempty" name:"ModelImage"`

	// 图片人脸融合（专业版）效果参数。可用于设置拉脸、人脸增强、磨皮、牙齿增强、妆容迁移等融合效果参数，生成理想的融合效果。不传默认使用接口推荐值。具体见FusionUltraParam说明
	FusionUltraParam *FusionUltraParam `json:"FusionUltraParam,omitnil,omitempty" name:"FusionUltraParam"`

	// 为融合结果图添加合成标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了人脸融合技术，是AI合成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在融合结果图右下角添加“本图片为AI合成图片”字样，您可根据自身需要替换为其他的Logo图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 融合模型类型参数：默认为1。
	// 图片人脸融合（专业版）针对不同场景，提供多种模型供选择。如您的产品是泛娱乐场景，推荐使用1；如您主要用于影像场景，推荐使用4、5。其他模型类型也可以结合您的产品使用场景进行选择，也许会有意想不到的效果
	// 1：默认泛娱乐场景，画面偏锐。
	// 2：影视级场景，画面偏自然。
	// 3：影视级场景，高分辨率，画面偏自然。
	// 4：影视级场景，高分辦率，高人脸相似度，画面偏自然，可用于证件照等场景。
	// 5：影视级场景，高分辨率，对闭眼和遮挡更友好。
	SwapModelType *int64 `json:"SwapModelType,omitnil,omitempty" name:"SwapModelType"`
}

func (r *FuseFaceUltraRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FuseFaceUltraRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RspImgType")
	delete(f, "MergeInfos")
	delete(f, "ModelUrl")
	delete(f, "ModelImage")
	delete(f, "FusionUltraParam")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "SwapModelType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FuseFaceUltraRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FuseFaceUltraResponseParams struct {
	// RspImgType 为 url 时，返回结果的 url， RspImgType 为 base64 时返回 base64 数据。url有效期为1天。
	FusedImage *string `json:"FusedImage,omitnil,omitempty" name:"FusedImage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FuseFaceUltraResponse struct {
	*tchttp.BaseResponse
	Response *FuseFaceUltraResponseParams `json:"Response"`
}

func (r *FuseFaceUltraResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FuseFaceUltraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FuseParam struct {
	// 图片编码参数
	ImageCodecParam *ImageCodecParam `json:"ImageCodecParam,omitnil,omitempty" name:"ImageCodecParam"`
}

type FusionUltraParam struct {
	// 拉脸强度。主要用于调整生成结果人脸脸型更像素材模板还是用户人脸。取值越大越像用户人脸。
	// 取值范围：0-1之间。默认取值0.7。
	// 
	// 该参数仅对SwapModelType（模型类型）取值1-5生效
	WarpRadio *float64 `json:"WarpRadio,omitnil,omitempty" name:"WarpRadio"`

	// 人脸增强强度。对整个人脸进行增强，增加清晰度，改善质量。当生成的人脸不够清晰，质感不够好的时候可以设置。取值越大增强强度越大。
	// 取值范围：0-1之间。默认取值0.5。
	// 
	// 该参数仅对SwapModelType（模型类型）取值1-5生效
	EnhanceRadio *float64 `json:"EnhanceRadio,omitnil,omitempty" name:"EnhanceRadio"`

	// 磨皮强度。当生成脸的图像面部显脏时，可进行设置。
	// 取值范围：0-1之间。默认取值0.5。
	// 
	// 该参数仅对SwapModelType（模型类型）取值1-5生效
	MpRadio *float64 `json:"MpRadio,omitnil,omitempty" name:"MpRadio"`

	// 人脸模糊开关（暂不支持）
	// 当生成人脸比较清晰时，将人脸模糊到接近模板的清晰度的程度
	// 
	// 该参数仅对SwapModelType（模型类型）取值1-5生效
	BlurRadio *float64 `json:"BlurRadio,omitnil,omitempty" name:"BlurRadio"`

	// 牙齿增强开关，默认取值为1
	// 牙齿增强，修复牙齿。当生成牙齿不好（如牙齿裂开）可以打开此开关
	// 0：牙齿增强关闭
	// 1：牙齿增强打开
	// 该参数仅对SwapModelType（模型类型）取值1-5生效
	TeethEnhanceRadio *float64 `json:"TeethEnhanceRadio,omitnil,omitempty" name:"TeethEnhanceRadio"`

	// 妆容迁移开关，默认取值为0。
	// 将素材模板的妆容迁移到融合结果上。即：如果希望妆容效果与模板图保持一致，可以打开此开关。
	// 0：妆容迁移关闭
	// 1：妆容迁移打开
	// 该参数仅对SwapModelType（模型类型）取值1-5生效
	MakeupTransferRadio *float64 `json:"MakeupTransferRadio,omitnil,omitempty" name:"MakeupTransferRadio"`
}

type ImageCodecParam struct {
	// 元数据是描述媒体文件的附加信息。通过添加自定义的元数据，可以将一些附加信息嵌入到文件中。这些信息可以用于版权、描述、标识等目的，并在后续的媒体处理或管理过程中使用。
	// 个数不能大于1。
	MetaData []*MetaData `json:"MetaData,omitnil,omitempty" name:"MetaData"`
}

type LogoParam struct {
	// 标识图片位于融合结果图中的坐标，将按照坐标对标识图片进行位置和大小的拉伸匹配。
	// Width、Height <= 2160。
	LogoRect *FaceRect `json:"LogoRect,omitnil,omitempty" name:"LogoRect"`

	// 标识图片Url地址
	// 
	// ●base64 和 url 必须提供一个，如果都提供以 url 为准。
	// ●支持图片格式：支持jpg或png
	// 专业版：base64 编码后大小不超过10M。
	// 非专业版：base64 编码后大小不超过5M。
	LogoUrl *string `json:"LogoUrl,omitnil,omitempty" name:"LogoUrl"`

	// 输入图片base64。
	// ●base64 和 url 必须提供一个，如果都提供以 url 为准。
	// ●支持图片格式：支持jpg或png
	// 专业版：base64 编码后大小不超过10M。
	// 非专业版：base64 编码后大小不超过5M。
	LogoImage *string `json:"LogoImage,omitnil,omitempty" name:"LogoImage"`
}

type MaterialFaces struct {
	// 人脸序号
	FaceId *string `json:"FaceId,omitnil,omitempty" name:"FaceId"`

	// 人脸框信息
	FaceInfo *FaceInfo `json:"FaceInfo,omitnil,omitempty" name:"FaceInfo"`
}

type MergeInfo struct {
	// 输入图片base64。
	// ●base64 和 url 必须提供一个，如果都提供以 url 为准。
	// ●素材图片限制：图片中面部尺寸大于34 * 34；图片尺寸大于64 * 64。（图片编码之后可能会大30%左右，建议合理控制图片大小）。
	// ●支持图片格式：支持jpg或png
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 输入图片url。
	// ●base64 和 url 必须提供一个，如果都提供以 url 为准。
	// ●素材图片限制：图片中面部尺寸大于34 * 34；图片尺寸大于64 * 64。（图片编码之后可能会大30%左右，建议合理控制图片大小）。
	// ●支持图片格式：支持jpg或png
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 上传的图片人脸位置信息（人脸框）
	// Width、Height >= 30。
	InputImageFaceRect *FaceRect `json:"InputImageFaceRect,omitnil,omitempty" name:"InputImageFaceRect"`

	// 素材人脸ID，不填默认取最大人脸。
	TemplateFaceID *string `json:"TemplateFaceID,omitnil,omitempty" name:"TemplateFaceID"`

	// 模板中人脸位置信息(人脸框)，不填默认取最大人脸。此字段仅适用于图片融合自定义模板素材场景。
	// Width、Height >= 30。
	TemplateFaceRect *FaceRect `json:"TemplateFaceRect,omitnil,omitempty" name:"TemplateFaceRect"`
}

type MetaData struct {
	// MetaData的Key，字符长度不能超过32
	MetaKey *string `json:"MetaKey,omitnil,omitempty" name:"MetaKey"`

	// MetaData的Value，字符长度不能超过256
	MetaValue *string `json:"MetaValue,omitnil,omitempty" name:"MetaValue"`
}

type PublicMaterialInfos struct {
	// 素材Id
	MaterialId *string `json:"MaterialId,omitnil,omitempty" name:"MaterialId"`

	// 素材状态
	// 
	// 字段取值: 
	// 0   审核中
	// 1   人工审核通过
	// 2   人工审核失败
	// 3   申诉中
	// 11  申诉成功
	// 12  申诉失败
	// 21  机器审核通过
	// 22  机器审核失败
	// 31  视频素材预处理成功，素材可用
	// 32  视频素材预处理失败
	// 33  角色不在视频中
	MaterialStatus *int64 `json:"MaterialStatus,omitnil,omitempty" name:"MaterialStatus"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 人脸信息
	MaterialFaceList []*MaterialFaces `json:"MaterialFaceList,omitnil,omitempty" name:"MaterialFaceList"`

	// 素材名
	MaterialName *string `json:"MaterialName,omitnil,omitempty" name:"MaterialName"`

	// 审核原因
	AuditResult *string `json:"AuditResult,omitnil,omitempty" name:"AuditResult"`
}