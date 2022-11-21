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

package v20180301

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AnalyzeDenseLandmarksRequestParams struct {
	// 检测模式。0 为检测所有出现的人脸， 1 为检测面积最大的人脸。 
	// 默认为 0。 
	// 最多返回 5 张人脸的五官定位（人脸关键点）具体信息。
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。  
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。  
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。  
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 人脸识别服务所用的算法模型版本。本接口仅支持 “3.0“ 输入。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *int64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

type AnalyzeDenseLandmarksRequest struct {
	*tchttp.BaseRequest
	
	// 检测模式。0 为检测所有出现的人脸， 1 为检测面积最大的人脸。 
	// 默认为 0。 
	// 最多返回 5 张人脸的五官定位（人脸关键点）具体信息。
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。  
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。  
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。  
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 人脸识别服务所用的算法模型版本。本接口仅支持 “3.0“ 输入。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *int64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *AnalyzeDenseLandmarksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeDenseLandmarksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "FaceModelVersion")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AnalyzeDenseLandmarksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AnalyzeDenseLandmarksResponseParams struct {
	// 请求的图片宽度。
	ImageWidth *int64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

	// 请求的图片高度。
	ImageHeight *int64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

	// 稠密人脸关键点具体信息。
	DenseFaceShapeSet []*DenseFaceShape `json:"DenseFaceShapeSet,omitempty" name:"DenseFaceShapeSet"`

	// 人脸识别服务所用的算法模型版本。本接口仅支持 “3.0“ 输入。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AnalyzeDenseLandmarksResponse struct {
	*tchttp.BaseResponse
	Response *AnalyzeDenseLandmarksResponseParams `json:"Response"`
}

func (r *AnalyzeDenseLandmarksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeDenseLandmarksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AnalyzeFaceRequestParams struct {
	// 检测模式。0 为检测所有出现的人脸， 1 为检测面积最大的人脸。默认为 0。最多返回 10 张人脸的五官定位（人脸关键点）具体信息。
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 人脸识别服务所用的算法模型版本。
	// 
	// 目前入参支持 “2.0”和“3.0“ 两个输入。
	// 
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。
	// 
	// 2020年11月26日后开通服务的账号仅支持输入“3.0”。
	// 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

type AnalyzeFaceRequest struct {
	*tchttp.BaseRequest
	
	// 检测模式。0 为检测所有出现的人脸， 1 为检测面积最大的人脸。默认为 0。最多返回 10 张人脸的五官定位（人脸关键点）具体信息。
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 人脸识别服务所用的算法模型版本。
	// 
	// 目前入参支持 “2.0”和“3.0“ 两个输入。
	// 
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。
	// 
	// 2020年11月26日后开通服务的账号仅支持输入“3.0”。
	// 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *AnalyzeFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "FaceModelVersion")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AnalyzeFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AnalyzeFaceResponseParams struct {
	// 请求的图片宽度。
	ImageWidth *uint64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

	// 请求的图片高度。
	ImageHeight *uint64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

	// 五官定位（人脸关键点）具体信息。
	FaceShapeSet []*FaceShape `json:"FaceShapeSet,omitempty" name:"FaceShapeSet"`

	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AnalyzeFaceResponse struct {
	*tchttp.BaseResponse
	Response *AnalyzeFaceResponseParams `json:"Response"`
}

func (r *AnalyzeFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeItem struct {
	// 属性值
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Type识别概率值，【0,1】,代表判断正确的概率。
	Probability *float64 `json:"Probability,omitempty" name:"Probability"`
}

type Candidate struct {
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人脸ID，仅在SearchFaces/SearchFacesReturnsByGroup接口返回时有效。人员搜索类接口采用融合特征方式检索，该字段无意义
	FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

	// 候选者的匹配得分。 
	// 
	// 1万大小人脸底库下，误识率百分之一对应分数为70分，误识率千分之一对应分数为80分，误识率万分之一对应分数为90分；
	// 10万大小人脸底库下，误识率百分之一对应分数为80分，误识率千分之一对应分数为90分，误识率万分之一对应分数为100分；
	// 30万大小人脸底库下，误识率百分之一对应分数为85分，误识率千分之一对应分数为95分。
	// 
	// 一般80分左右可适用大部分场景，建议分数不要超过90分。您可以根据实际情况选择合适的分数。
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 人员名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 人员性别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 包含此人员的人员库及描述字段内容列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersonGroupInfos []*PersonGroupInfo `json:"PersonGroupInfos,omitempty" name:"PersonGroupInfos"`
}

// Predefined struct for user
type CompareFaceRequestParams struct {
	// A 图片 base64 数据，base64 编码后大小不可超过5M。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	ImageA *string `json:"ImageA,omitempty" name:"ImageA"`

	// B 图片 base64 数据，base64 编码后大小不可超过5M。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	ImageB *string `json:"ImageB,omitempty" name:"ImageB"`

	// A 图片的 Url ，对应图片 base64 编码后大小不可超过5M。
	// A 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	UrlA *string `json:"UrlA,omitempty" name:"UrlA"`

	// B 图片的 Url ，对应图片 base64 编码后大小不可超过5M。
	// B 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	UrlB *string `json:"UrlB,omitempty" name:"UrlB"`

	// 人脸识别服务所用的算法模型版本。
	// 
	// 目前入参支持 “2.0”和“3.0“ 两个输入。
	// 
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。
	// 
	// 2020年11月26日后开通服务的账号仅支持输入“3.0”。
	// 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

type CompareFaceRequest struct {
	*tchttp.BaseRequest
	
	// A 图片 base64 数据，base64 编码后大小不可超过5M。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	ImageA *string `json:"ImageA,omitempty" name:"ImageA"`

	// B 图片 base64 数据，base64 编码后大小不可超过5M。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	ImageB *string `json:"ImageB,omitempty" name:"ImageB"`

	// A 图片的 Url ，对应图片 base64 编码后大小不可超过5M。
	// A 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	UrlA *string `json:"UrlA,omitempty" name:"UrlA"`

	// B 图片的 Url ，对应图片 base64 编码后大小不可超过5M。
	// B 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	UrlB *string `json:"UrlB,omitempty" name:"UrlB"`

	// 人脸识别服务所用的算法模型版本。
	// 
	// 目前入参支持 “2.0”和“3.0“ 两个输入。
	// 
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。
	// 
	// 2020年11月26日后开通服务的账号仅支持输入“3.0”。
	// 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *CompareFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompareFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageA")
	delete(f, "ImageB")
	delete(f, "UrlA")
	delete(f, "UrlB")
	delete(f, "FaceModelVersion")
	delete(f, "QualityControl")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompareFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompareFaceResponseParams struct {
	// 两张图片中人脸的相似度分数。
	// 不同算法版本返回的相似度分数不同。 
	// 若需要验证两张图片中人脸是否为同一人，3.0版本误识率千分之一对应分数为40分，误识率万分之一对应分数为50分，误识率十万分之一对应分数为60分。  一般超过50分则可认定为同一人。 
	// 2.0版本误识率千分之一对应分数为70分，误识率万分之一对应分数为80分，误识率十万分之一对应分数为90分。 一般超过80分则可认定为同一人。 
	// 若需要验证两张图片中的人脸是否为同一人，建议使用人脸验证接口。
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CompareFaceResponse struct {
	*tchttp.BaseResponse
	Response *CompareFaceResponseParams `json:"Response"`
}

func (r *CompareFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompareFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyPersonRequestParams struct {
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 待加入的人员库列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
}

type CopyPersonRequest struct {
	*tchttp.BaseRequest
	
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 待加入的人员库列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
}

func (r *CopyPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyPersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyPersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyPersonResponseParams struct {
	// 成功加入的人员库数量
	SucGroupNum *uint64 `json:"SucGroupNum,omitempty" name:"SucGroupNum"`

	// 成功加入的人员库列表
	SucGroupIds []*string `json:"SucGroupIds,omitempty" name:"SucGroupIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CopyPersonResponse struct {
	*tchttp.BaseResponse
	Response *CopyPersonResponseParams `json:"Response"`
}

func (r *CopyPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyPersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFaceRequestParams struct {
	// 人员ID。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 人员人脸总数量不可超过5张。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Images []*string `json:"Images,omitempty" name:"Images"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 人员人脸总数量不可超过5张。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 只有和该人员已有的人脸相似度超过FaceMatchThreshold值的人脸，才能增加人脸成功。 
	// 默认值60分。取值范围[0,100] 。
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

type CreateFaceRequest struct {
	*tchttp.BaseRequest
	
	// 人员ID。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 人员人脸总数量不可超过5张。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Images []*string `json:"Images,omitempty" name:"Images"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 人员人脸总数量不可超过5张。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 只有和该人员已有的人脸相似度超过FaceMatchThreshold值的人脸，才能增加人脸成功。 
	// 默认值60分。取值范围[0,100] 。
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *CreateFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Images")
	delete(f, "Urls")
	delete(f, "FaceMatchThreshold")
	delete(f, "QualityControl")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFaceResponseParams struct {
	// 加入成功的人脸数量
	SucFaceNum *uint64 `json:"SucFaceNum,omitempty" name:"SucFaceNum"`

	// 加入成功的人脸ID列表
	SucFaceIds []*string `json:"SucFaceIds,omitempty" name:"SucFaceIds"`

	// 每张人脸图片添加结果，-1101 代表未检测到人脸，-1102 代表图片解码失败， 
	// -1601代表不符合图片质量控制要求, -1604 代表人脸相似度没有超过FaceMatchThreshold。 
	// 其他非 0 值代表算法服务异常。 
	// RetCode的顺序和入参中 Images 或 Urls 的顺序一致。
	RetCode []*int64 `json:"RetCode,omitempty" name:"RetCode"`

	// 加入成功的人脸索引。索引顺序和入参中 Images 或 Urls 的顺序一致。 
	// 例， Urls 中 有 3 个 url，第二个 url 失败，则 SucIndexes 值为 [0,2] 。
	SucIndexes []*uint64 `json:"SucIndexes,omitempty" name:"SucIndexes"`

	// 加入成功的人脸框位置。顺序和入参中 Images 或 Urls 的顺序一致。
	SucFaceRects []*FaceRect `json:"SucFaceRects,omitempty" name:"SucFaceRects"`

	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateFaceResponseParams `json:"Response"`
}

func (r *CreateFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupRequestParams struct {
	// 人员库名称，[1,60]个字符，可修改，不可重复。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 人员库 ID，不可修改，不可重复。支持英文、数字、-%@#&_，长度限制64B。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员库自定义描述字段，用于描述人员库中人员属性，该人员库下所有人员将拥有此描述字段。 
	// 最多可以创建5个。 
	// 每个自定义描述字段支持[1,30]个字符。 
	// 在同一人员库中自定义描述字段不可重复。 
	// 例： 设置某人员库“自定义描述字段”为["学号","工号","手机号"]， 
	// 则该人员库下所有人员将拥有名为“学号”、“工号”、“手机号”的描述字段， 
	// 可在对应人员描述字段中填写内容，登记该人员的学号、工号、手机号等信息。
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions"`

	// 人员库信息备注，[0，40]个字符。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 人脸识别服务所用的算法模型版本。
	// 
	// 目前入参支持 “2.0”和“3.0“ 两个输入。
	// 
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。
	// 
	// 2020年11月26日后开通服务的账号仅支持输入“3.0”。
	// 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest
	
	// 人员库名称，[1,60]个字符，可修改，不可重复。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 人员库 ID，不可修改，不可重复。支持英文、数字、-%@#&_，长度限制64B。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员库自定义描述字段，用于描述人员库中人员属性，该人员库下所有人员将拥有此描述字段。 
	// 最多可以创建5个。 
	// 每个自定义描述字段支持[1,30]个字符。 
	// 在同一人员库中自定义描述字段不可重复。 
	// 例： 设置某人员库“自定义描述字段”为["学号","工号","手机号"]， 
	// 则该人员库下所有人员将拥有名为“学号”、“工号”、“手机号”的描述字段， 
	// 可在对应人员描述字段中填写内容，登记该人员的学号、工号、手机号等信息。
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions"`

	// 人员库信息备注，[0，40]个字符。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 人脸识别服务所用的算法模型版本。
	// 
	// 目前入参支持 “2.0”和“3.0“ 两个输入。
	// 
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。
	// 
	// 2020年11月26日后开通服务的账号仅支持输入“3.0”。
	// 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

func (r *CreateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "GroupId")
	delete(f, "GroupExDescriptions")
	delete(f, "Tag")
	delete(f, "FaceModelVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupResponseParams struct {
	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateGroupResponseParams `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonRequestParams struct {
	// 待加入的人员库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员名称。[1，60]个字符，可修改，可重复。
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 人员ID，单个腾讯云账号下不可修改，不可重复。支持英文、数字、-%@#&_，长度限制64B。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 0代表未填写，1代表男性，2代表女性。
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 人员描述字段内容，key-value。[0，60]个字符，可修改，可重复。
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitempty" name:"PersonExDescriptionInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 此参数用于控制判断 Image 或 Url 中图片包含的人脸，是否在人员库中已有疑似的同一人。 
	// 如果判断为已有相同人在人员库中，则不会创建新的人员，返回疑似同一人的人员信息。 
	// 如果判断没有，则完成创建人员。 
	// 0: 不进行判断，无论是否有疑似同一人在库中均完成入库； 
	// 1:较低的同一人判断要求（百一误识别率）； 
	// 2: 一般的同一人判断要求（千一误识别率）； 
	// 3: 较高的同一人判断要求（万一误识别率）； 
	// 4: 很高的同一人判断要求（十万一误识别率）。 
	// 默认 0。  
	// 注： 要求越高，则疑似同一人的概率越小。不同要求对应的误识别率仅为参考值，您可以根据实际情况调整。
	UniquePersonControl *uint64 `json:"UniquePersonControl,omitempty" name:"UniquePersonControl"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

type CreatePersonRequest struct {
	*tchttp.BaseRequest
	
	// 待加入的人员库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员名称。[1，60]个字符，可修改，可重复。
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 人员ID，单个腾讯云账号下不可修改，不可重复。支持英文、数字、-%@#&_，长度限制64B。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 0代表未填写，1代表男性，2代表女性。
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 人员描述字段内容，key-value。[0，60]个字符，可修改，可重复。
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitempty" name:"PersonExDescriptionInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 此参数用于控制判断 Image 或 Url 中图片包含的人脸，是否在人员库中已有疑似的同一人。 
	// 如果判断为已有相同人在人员库中，则不会创建新的人员，返回疑似同一人的人员信息。 
	// 如果判断没有，则完成创建人员。 
	// 0: 不进行判断，无论是否有疑似同一人在库中均完成入库； 
	// 1:较低的同一人判断要求（百一误识别率）； 
	// 2: 一般的同一人判断要求（千一误识别率）； 
	// 3: 较高的同一人判断要求（万一误识别率）； 
	// 4: 很高的同一人判断要求（十万一误识别率）。 
	// 默认 0。  
	// 注： 要求越高，则疑似同一人的概率越小。不同要求对应的误识别率仅为参考值，您可以根据实际情况调整。
	UniquePersonControl *uint64 `json:"UniquePersonControl,omitempty" name:"UniquePersonControl"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *CreatePersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "PersonName")
	delete(f, "PersonId")
	delete(f, "Gender")
	delete(f, "PersonExDescriptionInfos")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "UniquePersonControl")
	delete(f, "QualityControl")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonResponseParams struct {
	// 人脸图片唯一标识。
	FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

	// 检测出的人脸框的位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

	// 疑似同一人的PersonId。 
	// 当 UniquePersonControl 参数不为0且人员库中有疑似的同一人，此参数才有意义。
	SimilarPersonId *string `json:"SimilarPersonId,omitempty" name:"SimilarPersonId"`

	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePersonResponse struct {
	*tchttp.BaseResponse
	Response *CreatePersonResponseParams `json:"Response"`
}

func (r *CreatePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFaceRequestParams struct {
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 待删除的人脸ID列表
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds"`
}

type DeleteFaceRequest struct {
	*tchttp.BaseRequest
	
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 待删除的人脸ID列表
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds"`
}

func (r *DeleteFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "FaceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFaceResponseParams struct {
	// 删除成功的人脸数量
	SucDeletedNum *uint64 `json:"SucDeletedNum,omitempty" name:"SucDeletedNum"`

	// 删除成功的人脸ID列表
	SucFaceIds []*string `json:"SucFaceIds,omitempty" name:"SucFaceIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteFaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFaceResponseParams `json:"Response"`
}

func (r *DeleteFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupRequestParams struct {
	// 人员库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// 人员库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupResponseParams `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonFromGroupRequestParams struct {
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DeletePersonFromGroupRequest struct {
	*tchttp.BaseRequest
	
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeletePersonFromGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonFromGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePersonFromGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonFromGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePersonFromGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeletePersonFromGroupResponseParams `json:"Response"`
}

func (r *DeletePersonFromGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonFromGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonRequestParams struct {
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

type DeletePersonRequest struct {
	*tchttp.BaseRequest
	
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

func (r *DeletePersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePersonResponse struct {
	*tchttp.BaseResponse
	Response *DeletePersonResponseParams `json:"Response"`
}

func (r *DeletePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DenseFaceShape struct {
	// 人脸框左上角横坐标。
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸框左上角纵坐标。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸框宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 人脸框高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 描述左侧眼睛轮廓的 XX 点。
	LeftEye []*Point `json:"LeftEye,omitempty" name:"LeftEye"`

	// 描述右侧眼睛轮廓的 XX 点。
	RightEye []*Point `json:"RightEye,omitempty" name:"RightEye"`

	// 描述左侧眉毛轮廓的 XX 点。
	LeftEyeBrow []*Point `json:"LeftEyeBrow,omitempty" name:"LeftEyeBrow"`

	// 描述右侧眉毛轮廓的 XX 点。
	RightEyeBrow []*Point `json:"RightEyeBrow,omitempty" name:"RightEyeBrow"`

	// 描述外嘴巴轮廓的 XX 点， 从左侧开始逆时针返回。
	MouthOutside []*Point `json:"MouthOutside,omitempty" name:"MouthOutside"`

	// 描述内嘴巴轮廓的 XX 点，从左侧开始逆时针返回。
	MouthInside []*Point `json:"MouthInside,omitempty" name:"MouthInside"`

	// 描述鼻子轮廓的 XX 点。
	Nose []*Point `json:"Nose,omitempty" name:"Nose"`

	// 左瞳孔轮廓的 XX 个点。
	LeftPupil []*Point `json:"LeftPupil,omitempty" name:"LeftPupil"`

	// 右瞳孔轮廓的 XX 个点。
	RightPupil []*Point `json:"RightPupil,omitempty" name:"RightPupil"`

	// 中轴线轮廓的 XX 个点。
	CentralAxis []*Point `json:"CentralAxis,omitempty" name:"CentralAxis"`

	// 下巴轮廓的 XX 个点。
	Chin []*Point `json:"Chin,omitempty" name:"Chin"`

	// 左眼袋的 XX 个点。
	LeftEyeBags []*Point `json:"LeftEyeBags,omitempty" name:"LeftEyeBags"`

	// 右眼袋的 XX 个点。
	RightEyeBags []*Point `json:"RightEyeBags,omitempty" name:"RightEyeBags"`

	// 额头的 XX 个点。
	Forehead []*Point `json:"Forehead,omitempty" name:"Forehead"`
}

// Predefined struct for user
type DetectFaceAttributesRequestParams struct {
	// 最多处理的人脸数目。 
	// 默认值为1（仅检测图片中面积最大的那张人脸），最大值为120。 
	// 此参数用于控制处理待检测图片中的人脸个数，值越小，处理速度越快。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。 
	// 对应图片 base64 编码后大小不可超过5M。
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 
	// Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 是否返回年龄、性别、情绪等属性。 
	// 合法值为（大小写不敏感）：None、Age、Beauty、Emotion、Eye、Eyebrow、 
	// Gender、Hair、Hat、Headpose、Mask、Mouth、Moustache、Nose、Shape、Skin、Smile。 
	// None为不需要返回。默认为 None。即FaceAttributesType属性为空时，各属性返回值为0。
	// 需要将属性组成一个用逗号分隔的字符串，属性之间的顺序没有要求。 
	// 关于各属性的详细描述，参见下文出参。 
	// 最多返回面积最大的 5 张人脸属性信息，超过 5 张人脸（第 6 张及以后的人脸）的 AttributesInfo 不具备参考意义。
	FaceAttributesType *string `json:"FaceAttributesType,omitempty" name:"FaceAttributesType"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`

	// 人脸识别服务所用的算法模型版本。本接口仅支持“3.0”输入
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

type DetectFaceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 最多处理的人脸数目。 
	// 默认值为1（仅检测图片中面积最大的那张人脸），最大值为120。 
	// 此参数用于控制处理待检测图片中的人脸个数，值越小，处理速度越快。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。 
	// 对应图片 base64 编码后大小不可超过5M。
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 
	// Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 是否返回年龄、性别、情绪等属性。 
	// 合法值为（大小写不敏感）：None、Age、Beauty、Emotion、Eye、Eyebrow、 
	// Gender、Hair、Hat、Headpose、Mask、Mouth、Moustache、Nose、Shape、Skin、Smile。 
	// None为不需要返回。默认为 None。即FaceAttributesType属性为空时，各属性返回值为0。
	// 需要将属性组成一个用逗号分隔的字符串，属性之间的顺序没有要求。 
	// 关于各属性的详细描述，参见下文出参。 
	// 最多返回面积最大的 5 张人脸属性信息，超过 5 张人脸（第 6 张及以后的人脸）的 AttributesInfo 不具备参考意义。
	FaceAttributesType *string `json:"FaceAttributesType,omitempty" name:"FaceAttributesType"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`

	// 人脸识别服务所用的算法模型版本。本接口仅支持“3.0”输入
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

func (r *DetectFaceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFaceAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaxFaceNum")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "FaceAttributesType")
	delete(f, "NeedRotateDetection")
	delete(f, "FaceModelVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectFaceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectFaceAttributesResponseParams struct {
	// 请求的图片宽度。
	ImageWidth *uint64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

	// 请求的图片高度。
	ImageHeight *uint64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

	// 人脸信息列表。
	FaceDetailInfos []*FaceDetailInfo `json:"FaceDetailInfos,omitempty" name:"FaceDetailInfos"`

	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectFaceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DetectFaceAttributesResponseParams `json:"Response"`
}

func (r *DetectFaceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFaceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectFaceRequestParams struct {
	// 最多处理的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为120。 
	// 此参数用于控制处理待检测图片中的人脸个数，值越小，处理速度越快。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。
	// 默认为34。建议不低于34。
	// 低于MinFaceSize值的人脸不会被检测。
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 是否需要返回人脸属性信息（FaceAttributesInfo）。0 为不需要返回，1 为需要返回。默认为 0。 
	// 非 1 值均视为不需要返回，此时 FaceAttributesInfo 不具备参考意义。  
	// 最多返回面积最大的 5 张人脸属性信息，超过 5 张人脸（第 6 张及以后的人脸）的 FaceAttributesInfo 不具备参考意义。  
	// 提取人脸属性信息较为耗时，如不需要人脸属性信息，建议关闭此项功能，加快人脸检测速度。
	NeedFaceAttributes *uint64 `json:"NeedFaceAttributes,omitempty" name:"NeedFaceAttributes"`

	// 是否开启质量检测。0 为关闭，1 为开启。默认为 0。 
	// 非 1 值均视为不进行质量检测。
	// 最多返回面积最大的 30 张人脸质量分信息，超过 30 张人脸（第 31 张及以后的人脸）的 FaceQualityInfo不具备参考意义。  
	// 建议：人脸入库操作建议开启此功能。
	NeedQualityDetection *uint64 `json:"NeedQualityDetection,omitempty" name:"NeedQualityDetection"`

	// 人脸识别服务所用的算法模型版本。目前入参支持 “2.0”和“3.0“ 两个输入。  
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

type DetectFaceRequest struct {
	*tchttp.BaseRequest
	
	// 最多处理的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为120。 
	// 此参数用于控制处理待检测图片中的人脸个数，值越小，处理速度越快。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。
	// 默认为34。建议不低于34。
	// 低于MinFaceSize值的人脸不会被检测。
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 是否需要返回人脸属性信息（FaceAttributesInfo）。0 为不需要返回，1 为需要返回。默认为 0。 
	// 非 1 值均视为不需要返回，此时 FaceAttributesInfo 不具备参考意义。  
	// 最多返回面积最大的 5 张人脸属性信息，超过 5 张人脸（第 6 张及以后的人脸）的 FaceAttributesInfo 不具备参考意义。  
	// 提取人脸属性信息较为耗时，如不需要人脸属性信息，建议关闭此项功能，加快人脸检测速度。
	NeedFaceAttributes *uint64 `json:"NeedFaceAttributes,omitempty" name:"NeedFaceAttributes"`

	// 是否开启质量检测。0 为关闭，1 为开启。默认为 0。 
	// 非 1 值均视为不进行质量检测。
	// 最多返回面积最大的 30 张人脸质量分信息，超过 30 张人脸（第 31 张及以后的人脸）的 FaceQualityInfo不具备参考意义。  
	// 建议：人脸入库操作建议开启此功能。
	NeedQualityDetection *uint64 `json:"NeedQualityDetection,omitempty" name:"NeedQualityDetection"`

	// 人脸识别服务所用的算法模型版本。目前入参支持 “2.0”和“3.0“ 两个输入。  
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *DetectFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaxFaceNum")
	delete(f, "MinFaceSize")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "NeedFaceAttributes")
	delete(f, "NeedQualityDetection")
	delete(f, "FaceModelVersion")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectFaceResponseParams struct {
	// 请求的图片宽度。
	ImageWidth *int64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

	// 请求的图片高度。
	ImageHeight *int64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

	// 人脸信息列表。包含人脸坐标信息、属性信息（若需要）、质量分信息（若需要）。
	FaceInfos []*FaceInfo `json:"FaceInfos,omitempty" name:"FaceInfos"`

	// 人脸识别服务所用的算法模型版本。
	// 
	// 目前入参支持 “2.0”和“3.0“ 两个输入。
	// 
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。
	// 
	// 2020年11月26日后开通服务的账号仅支持输入“3.0”。
	// 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectFaceResponse struct {
	*tchttp.BaseResponse
	Response *DetectFaceResponseParams `json:"Response"`
}

func (r *DetectFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectLiveFaceRequestParams struct {
	// 图片 base64 数据，base64 编码后大小不可超过5M（图片的宽高比请接近3:4，不符合宽高比的图片返回的分值不具备参考意义）。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。 
	// （图片的宽高比请接近 3:4，不符合宽高比的图片返回的分值不具备参考意义） 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 人脸识别服务所用的算法模型版本。
	// 
	// 目前入参支持 “2.0”和“3.0“ 两个输入。
	// 
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。
	// 
	// 2020年11月26日后开通服务的账号仅支持输入“3.0”。
	// 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

type DetectLiveFaceRequest struct {
	*tchttp.BaseRequest
	
	// 图片 base64 数据，base64 编码后大小不可超过5M（图片的宽高比请接近3:4，不符合宽高比的图片返回的分值不具备参考意义）。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。 
	// （图片的宽高比请接近 3:4，不符合宽高比的图片返回的分值不具备参考意义） 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 人脸识别服务所用的算法模型版本。
	// 
	// 目前入参支持 “2.0”和“3.0“ 两个输入。
	// 
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。
	// 
	// 2020年11月26日后开通服务的账号仅支持输入“3.0”。
	// 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

func (r *DetectLiveFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLiveFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "FaceModelVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectLiveFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectLiveFaceResponseParams struct {
	// 活体打分，取值范围 [0,100]，分数一般落于[80, 100]区间内，0分也为常见值。推荐相大于 87 时可判断为活体。可根据具体场景自行调整阈值。
	// 本字段当且仅当FaceModelVersion为2.0时才具备参考意义。
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 活体检测是否通过。
	// 本字段只有FaceModelVersion为3.0时才具备参考意义。
	IsLiveness *bool `json:"IsLiveness,omitempty" name:"IsLiveness"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectLiveFaceResponse struct {
	*tchttp.BaseResponse
	Response *DetectLiveFaceResponseParams `json:"Response"`
}

func (r *DetectLiveFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLiveFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Eye struct {
	// 识别是否佩戴眼镜。
	// AttributeItem对应的Type为 —— 0：无眼镜，1：普通眼镜，2：墨镜
	Glass *AttributeItem `json:"Glass,omitempty" name:"Glass"`

	// 识别眼睛的睁开、闭合状态。
	// AttributeItem对应的Type为 —— 0：睁开，1：闭眼
	EyeOpen *AttributeItem `json:"EyeOpen,omitempty" name:"EyeOpen"`

	// 识别是否双眼皮。
	// AttributeItem对应的Type为 —— 0：无，1：有。
	EyelidType *AttributeItem `json:"EyelidType,omitempty" name:"EyelidType"`

	// 眼睛大小。
	// AttributeItem对应的Type为 —— 0：小眼睛，1：普通眼睛，2：大眼睛。
	EyeSize *AttributeItem `json:"EyeSize,omitempty" name:"EyeSize"`
}

type Eyebrow struct {
	// 眉毛浓密。
	// AttributeItem对应的Type为 —— 0：淡眉，1：浓眉。
	EyebrowDensity *AttributeItem `json:"EyebrowDensity,omitempty" name:"EyebrowDensity"`

	// 眉毛弯曲。
	// AttributeItem对应的Type为 —— 0：不弯，1：弯眉。
	EyebrowCurve *AttributeItem `json:"EyebrowCurve,omitempty" name:"EyebrowCurve"`

	// 眉毛长短。
	// AttributeItem对应的Type为 —— 0：短眉毛，1：长眉毛。
	EyebrowLength *AttributeItem `json:"EyebrowLength,omitempty" name:"EyebrowLength"`
}

type FaceAttributesInfo struct {
	// 性别[0~49]为女性，[50，100]为男性，越接近0和100表示置信度越高。NeedFaceAttributes 不为 1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 年龄 [0~100]。NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 微笑[0(normal，正常)~50(smile，微笑)~100(laugh，大笑)]。NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Expression *int64 `json:"Expression,omitempty" name:"Expression"`

	// 是否有眼镜 [true,false]。NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Glass *bool `json:"Glass,omitempty" name:"Glass"`

	// 上下偏移[-30,30]，单位角度。NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。 
	// 建议：人脸入库选择[-10,10]的图片。
	Pitch *int64 `json:"Pitch,omitempty" name:"Pitch"`

	// 左右偏移[-30,30]，单位角度。 NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。 
	// 建议：人脸入库选择[-10,10]的图片。
	Yaw *int64 `json:"Yaw,omitempty" name:"Yaw"`

	// 平面旋转[-180,180]，单位角度。 NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。  
	// 建议：人脸入库选择[-20,20]的图片。
	Roll *int64 `json:"Roll,omitempty" name:"Roll"`

	// 魅力[0~100]。NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Beauty *int64 `json:"Beauty,omitempty" name:"Beauty"`

	// 是否有帽子 [true,false]。NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hat *bool `json:"Hat,omitempty" name:"Hat"`

	// 是否有口罩 [true,false]。NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mask *bool `json:"Mask,omitempty" name:"Mask"`

	// 头发信息，包含头发长度（length）、有无刘海（bang）、头发颜色（color）。NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hair *FaceHairAttributesInfo `json:"Hair,omitempty" name:"Hair"`

	// 双眼是否睁开 [true,false]。只要有超过一只眼睛闭眼，就返回false。 NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EyeOpen *bool `json:"EyeOpen,omitempty" name:"EyeOpen"`
}

type FaceDetailAttributesInfo struct {
	// 年龄 [0,65]，其中65代表“65岁及以上”。 
	// FaceAttributesType 不为含Age 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 美丑打分[0,100]。 
	// FaceAttributesType 不含 Beauty 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Beauty *int64 `json:"Beauty,omitempty" name:"Beauty"`

	// 情绪，可识别自然、高兴、惊讶、生气、悲伤、厌恶、害怕。 
	// AttributeItem对应的Type为 —— 0：自然，1：高兴，2：惊讶，3：生气，4：悲伤，5：厌恶，6：害怕
	// FaceAttributesType 不含Emotion 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Emotion *AttributeItem `json:"Emotion,omitempty" name:"Emotion"`

	// 眼睛相关信息，可识别是否戴眼镜、是否闭眼、是否双眼皮和眼睛大小。 
	// FaceAttributesType 不含Eye 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Eye *Eye `json:"Eye,omitempty" name:"Eye"`

	// 眉毛相关信息，可识别眉毛浓密、弯曲、长短信息。 
	// FaceAttributesType 不含Eyebrow 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Eyebrow *Eyebrow `json:"Eyebrow,omitempty" name:"Eyebrow"`

	// 性别信息。 
	// AttributeItem对应的Type为 —— 	0：男性，1：女性。
	// FaceAttributesType 不含Gender 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Gender *AttributeItem `json:"Gender,omitempty" name:"Gender"`

	// 头发信息，包含头发长度、有无刘海、头发颜色。 
	// FaceAttributesType 不含Hair 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Hair *Hair `json:"Hair,omitempty" name:"Hair"`

	// 帽子信息，可识别是否佩戴帽子、帽子款式、帽子颜色。 
	// FaceAttributesType 不含Hat 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Hat *Hat `json:"Hat,omitempty" name:"Hat"`

	// 姿态信息，包含人脸的上下偏移、左右偏移、平面旋转信息。 
	// FaceAttributesType 不含Headpose 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	HeadPose *HeadPose `json:"HeadPose,omitempty" name:"HeadPose"`

	// 口罩佩戴信息。 
	// AttributeItem对应的Type为 —— 0: 无口罩， 1: 有口罩不遮脸，2: 有口罩遮下巴，3: 有口罩遮嘴，4: 正确佩戴口罩。
	// FaceAttributesType 不含Mask 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Mask *AttributeItem `json:"Mask,omitempty" name:"Mask"`

	// 嘴巴信息，可识别是否张嘴、嘴唇厚度。 
	// FaceAttributesType 不含 Mouth 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Mouth *Mouth `json:"Mouth,omitempty" name:"Mouth"`

	// 胡子信息。
	// AttributeItem对应的Type为 —— 0：无胡子，1：有胡子。 
	// FaceAttributesType 不含 Moustache 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Moustache *AttributeItem `json:"Moustache,omitempty" name:"Moustache"`

	// 鼻子信息。 
	// AttributeItem对应的Type为 —— 0：朝天鼻，1：鹰钩鼻，2：普通，3：圆鼻头
	// FaceAttributesType 不含 Nose 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Nose *AttributeItem `json:"Nose,omitempty" name:"Nose"`

	// 脸型信息。 
	// AttributeItem对应的Type为 —— 0：方脸，1：三角脸，2：鹅蛋脸，3：心形脸，4：圆脸。
	// FaceAttributesType 不含 Shape 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Shape *AttributeItem `json:"Shape,omitempty" name:"Shape"`

	// 肤色信息。 
	// AttributeItem对应的Type为 —— 0：黄色皮肤，1：棕色皮肤，2：黑色皮肤，3：白色皮肤。
	// FaceAttributesType 不含 Skin 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Skin *AttributeItem `json:"Skin,omitempty" name:"Skin"`

	// 微笑程度，[0,100]。 
	// FaceAttributesType 不含 Smile 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Smile *int64 `json:"Smile,omitempty" name:"Smile"`
}

type FaceDetailInfo struct {
	// 检测出的人脸框位置。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

	// 人脸属性信息，根据 FaceAttributesType 输入的类型，返回年龄（Age）、颜值（Beauty） 
	// 情绪（Emotion）、眼睛信息（Eye）、眉毛（Eyebrow）、性别（Gender） 
	// 头发（Hair）、帽子（Hat）、姿态（Headpose）、口罩（Mask）、嘴巴（Mouse）、胡子（Moustache） 
	// 鼻子（Nose）、脸型（Shape）、肤色（Skin）、微笑（Smile）等人脸属性信息。  
	// 若 FaceAttributesType 没有输入相关类型，则FaceDetaiAttributesInfo返回的细项不具备参考意义。
	FaceDetailAttributesInfo *FaceDetailAttributesInfo `json:"FaceDetailAttributesInfo,omitempty" name:"FaceDetailAttributesInfo"`
}

type FaceHairAttributesInfo struct {
	// 0：光头，1：短发，2：中发，3：长发，4：绑发
	// 注意：此字段可能返回 null，表示取不到有效值。
	Length *int64 `json:"Length,omitempty" name:"Length"`

	// 0：有刘海，1：无刘海
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bang *int64 `json:"Bang,omitempty" name:"Bang"`

	// 0：黑色，1：金色，2：棕色，3：灰白色
	// 注意：此字段可能返回 null，表示取不到有效值。
	Color *int64 `json:"Color,omitempty" name:"Color"`
}

type FaceInfo struct {
	// 人脸框左上角横坐标。
	// 人脸框包含人脸五官位置并在此基础上进行一定的扩展，若人脸框超出图片范围，会导致坐标负值。 
	// 若需截取完整人脸，可以在完整分completess满足需求的情况下，将负值坐标取0。
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸框左上角纵坐标。 
	// 人脸框包含人脸五官位置并在此基础上进行一定的扩展，若人脸框超出图片范围，会导致坐标负值。 
	// 若需截取完整人脸，可以在完整分completess满足需求的情况下，将负值坐标取0。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸框宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 人脸框高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 人脸属性信息，包含性别( gender )、年龄( age )、表情( expression )、 
	// 魅力( beauty )、眼镜( glass )、口罩（mask）、头发（hair）和姿态 (pitch，roll，yaw )。只有当 NeedFaceAttributes 设为 1 时才返回有效信息。
	FaceAttributesInfo *FaceAttributesInfo `json:"FaceAttributesInfo,omitempty" name:"FaceAttributesInfo"`

	// 人脸质量信息，包含质量分（score）、模糊分（sharpness）、光照分（brightness）、遮挡分（completeness）。只有当NeedFaceDetection设为1时才返回有效信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceQualityInfo *FaceQualityInfo `json:"FaceQualityInfo,omitempty" name:"FaceQualityInfo"`
}

type FaceQualityCompleteness struct {
	// 眉毛的遮挡分数[0,100]，分数越高遮挡越少。 
	// 参考范围：[0,80]表示发生遮挡。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Eyebrow *int64 `json:"Eyebrow,omitempty" name:"Eyebrow"`

	// 眼睛的遮挡分数[0,100],分数越高遮挡越少。 
	// 参考范围：[0,80]表示发生遮挡。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Eye *int64 `json:"Eye,omitempty" name:"Eye"`

	// 鼻子的遮挡分数[0,100],分数越高遮挡越少。 
	// 参考范围：[0,60]表示发生遮挡。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nose *int64 `json:"Nose,omitempty" name:"Nose"`

	// 脸颊的遮挡分数[0,100],分数越高遮挡越少。 
	// 参考范围：[0,70]表示发生遮挡。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cheek *int64 `json:"Cheek,omitempty" name:"Cheek"`

	// 嘴巴的遮挡分数[0,100],分数越高遮挡越少。 
	// 参考范围：[0,50]表示发生遮挡。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mouth *int64 `json:"Mouth,omitempty" name:"Mouth"`

	// 下巴的遮挡分数[0,100],分数越高遮挡越少。 
	// 参考范围：[0,70]表示发生遮挡。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Chin *int64 `json:"Chin,omitempty" name:"Chin"`
}

type FaceQualityInfo struct {
	// 质量分: [0,100]，综合评价图像质量是否适合人脸识别，分数越高质量越好。 
	// 正常情况，只需要使用Score作为质量分总体的判断标准即可。Sharpness、Brightness、Completeness等细项分仅供参考。
	// 参考范围：[0,40]较差，[40,60] 一般，[60,80]较好，[80,100]很好。 
	// 建议：人脸入库选取70以上的图片。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 清晰分：[0,100]，评价图片清晰程度，分数越高越清晰。 
	// 参考范围：[0,40]特别模糊，[40,60]模糊，[60,80]一般，[80,100]清晰。 
	// 建议：人脸入库选取80以上的图片。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sharpness *int64 `json:"Sharpness,omitempty" name:"Sharpness"`

	// 光照分：[0,100]，评价图片光照程度，分数越高越亮。 
	// 参考范围： [0,30]偏暗，[30,70]光照正常，[70,100]偏亮。 
	// 建议：人脸入库选取[30,70]的图片。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Brightness *int64 `json:"Brightness,omitempty" name:"Brightness"`

	// 五官遮挡分，评价眉毛（Eyebrow）、眼睛（Eye）、鼻子（Nose）、脸颊（Cheek）、嘴巴（Mouth）、下巴（Chin）的被遮挡程度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Completeness *FaceQualityCompleteness `json:"Completeness,omitempty" name:"Completeness"`
}

type FaceRect struct {
	// 人脸框左上角横坐标。 
	// 人脸框包含人脸五官位置并在此基础上进行一定的扩展，若人脸框超出图片范围，会导致坐标负值。 
	// 若需截取完整人脸，可以在完整分completess满足需求的情况下，将负值坐标取0。
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸框左上角纵坐标。 
	// 人脸框包含人脸五官位置并在此基础上进行一定的扩展，若人脸框超出图片范围，会导致坐标负值。 
	// 若需截取完整人脸，可以在完整分completess满足需求的情况下，将负值坐标取0。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸宽度
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 人脸高度
	Height *uint64 `json:"Height,omitempty" name:"Height"`
}

type FaceShape struct {
	// 描述脸型轮廓的 21 点。
	FaceProfile []*Point `json:"FaceProfile,omitempty" name:"FaceProfile"`

	// 描述左侧眼睛轮廓的 8 点。
	LeftEye []*Point `json:"LeftEye,omitempty" name:"LeftEye"`

	// 描述右侧眼睛轮廓的 8 点。
	RightEye []*Point `json:"RightEye,omitempty" name:"RightEye"`

	// 描述左侧眉毛轮廓的 8 点。
	LeftEyeBrow []*Point `json:"LeftEyeBrow,omitempty" name:"LeftEyeBrow"`

	// 描述右侧眉毛轮廓的 8 点。
	RightEyeBrow []*Point `json:"RightEyeBrow,omitempty" name:"RightEyeBrow"`

	// 描述嘴巴轮廓的 22 点。
	Mouth []*Point `json:"Mouth,omitempty" name:"Mouth"`

	// 描述鼻子轮廓的 13 点。
	Nose []*Point `json:"Nose,omitempty" name:"Nose"`

	// 左瞳孔轮廓的 1 个点。
	LeftPupil []*Point `json:"LeftPupil,omitempty" name:"LeftPupil"`

	// 右瞳孔轮廓的 1 个点。
	RightPupil []*Point `json:"RightPupil,omitempty" name:"RightPupil"`
}

// Predefined struct for user
type GetGroupInfoRequestParams struct {
	// 人员库 ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type GetGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// 人员库 ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupInfoResponseParams struct {
	// 人员库名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员库自定义描述字段
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions"`

	// 人员库信息备注
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// Group的创建时间和日期 CreationTimestamp。CreationTimestamp 的值是自 Unix 纪元时间到Group创建时间的毫秒数。
	CreationTimestamp *uint64 `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetGroupInfoResponseParams `json:"Response"`
}

func (r *GetGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupListRequestParams struct {
	// 起始序号，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type GetGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 起始序号，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupListResponseParams struct {
	// 返回的人员库信息
	GroupInfos []*GroupInfo `json:"GroupInfos,omitempty" name:"GroupInfos"`

	// 人员库总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *GetGroupListResponseParams `json:"Response"`
}

func (r *GetGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonBaseInfoRequestParams struct {
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

type GetPersonBaseInfoRequest struct {
	*tchttp.BaseRequest
	
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`
}

func (r *GetPersonBaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonBaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPersonBaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonBaseInfoResponseParams struct {
	// 人员名称
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 人员性别
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 包含的人脸 ID 列表
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPersonBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetPersonBaseInfoResponseParams `json:"Response"`
}

func (r *GetPersonBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonGroupInfoRequestParams struct {
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 起始序号，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type GetPersonGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 起始序号，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetPersonGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPersonGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonGroupInfoResponseParams struct {
	// 包含此人员的人员库及描述字段内容列表
	PersonGroupInfos []*PersonGroupInfo `json:"PersonGroupInfos,omitempty" name:"PersonGroupInfos"`

	// 人员库总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`

	// 人脸识别服务所用的算法模型版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPersonGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetPersonGroupInfoResponseParams `json:"Response"`
}

func (r *GetPersonGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonListNumRequestParams struct {
	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type GetPersonListNumRequest struct {
	*tchttp.BaseRequest
	
	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetPersonListNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListNumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPersonListNumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonListNumResponseParams struct {
	// 人员数量
	PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

	// 人脸数量
	FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPersonListNumResponse struct {
	*tchttp.BaseResponse
	Response *GetPersonListNumResponseParams `json:"Response"`
}

func (r *GetPersonListNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonListRequestParams struct {
	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 起始序号，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type GetPersonListRequest struct {
	*tchttp.BaseRequest
	
	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 起始序号，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetPersonListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPersonListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonListResponseParams struct {
	// 返回的人员信息
	PersonInfos []*PersonInfo `json:"PersonInfos,omitempty" name:"PersonInfos"`

	// 该人员库的人员数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

	// 该人员库的人脸数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

	// 人脸识别所用的算法模型版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPersonListResponse struct {
	*tchttp.BaseResponse
	Response *GetPersonListResponseParams `json:"Response"`
}

func (r *GetPersonListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUpgradeGroupFaceModelVersionJobListRequestParams struct {
	// 起始序号，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type GetUpgradeGroupFaceModelVersionJobListRequest struct {
	*tchttp.BaseRequest
	
	// 起始序号，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetUpgradeGroupFaceModelVersionJobListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUpgradeGroupFaceModelVersionJobListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUpgradeGroupFaceModelVersionJobListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUpgradeGroupFaceModelVersionJobListResponseParams struct {
	// 人员库升级任务信息列表。
	JobInfos []*UpgradeJobInfo `json:"JobInfos,omitempty" name:"JobInfos"`

	// 升级任务总数量。
	JobNum *uint64 `json:"JobNum,omitempty" name:"JobNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetUpgradeGroupFaceModelVersionJobListResponse struct {
	*tchttp.BaseResponse
	Response *GetUpgradeGroupFaceModelVersionJobListResponseParams `json:"Response"`
}

func (r *GetUpgradeGroupFaceModelVersionJobListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUpgradeGroupFaceModelVersionJobListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUpgradeGroupFaceModelVersionResultRequestParams struct {
	// 升级任务ID，用于查询、获取人员库升级的进度和结果。
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type GetUpgradeGroupFaceModelVersionResultRequest struct {
	*tchttp.BaseRequest
	
	// 升级任务ID，用于查询、获取人员库升级的进度和结果。
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *GetUpgradeGroupFaceModelVersionResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUpgradeGroupFaceModelVersionResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUpgradeGroupFaceModelVersionResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUpgradeGroupFaceModelVersionResultResponseParams struct {
	// 人员升级任务预估结束时间。 StartTimestamp的值是自 Unix 纪元时间到人员查重任务预估结束的毫秒数。  
	// Unix 纪元时间是 1970 年 1 月 1 日星期四，协调世界时 (UTC) 00:00:00。 
	// 如果为0表示这个任务已经执行完毕。
	EndTimestamp *uint64 `json:"EndTimestamp,omitempty" name:"EndTimestamp"`

	// 升级任务完成进度。取值[0.0，100.0]。
	Progress *float64 `json:"Progress,omitempty" name:"Progress"`

	// 0表示升级中，1表示升级完毕，2表示回滚完毕。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 升级起始时间。 
	// StartTime的值是自 Unix 纪元时间到Group创建时间的毫秒数。 
	// Unix 纪元时间是 1970 年 1 月 1 日星期四，协调世界时 (UTC) 00:00:00。 
	// 有关更多信息，请参阅 Unix 时间。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 当前算法模型版本。
	FromFaceModelVersion *string `json:"FromFaceModelVersion,omitempty" name:"FromFaceModelVersion"`

	// 目标算法模型版本。
	ToFaceModelVersion *string `json:"ToFaceModelVersion,omitempty" name:"ToFaceModelVersion"`

	// 人员库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 无法升级的人脸Id信息，文件格式为json。半小时有效
	FailedFacesUrl *string `json:"FailedFacesUrl,omitempty" name:"FailedFacesUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetUpgradeGroupFaceModelVersionResultResponse struct {
	*tchttp.BaseResponse
	Response *GetUpgradeGroupFaceModelVersionResultResponseParams `json:"Response"`
}

func (r *GetUpgradeGroupFaceModelVersionResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUpgradeGroupFaceModelVersionResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupCandidate struct {
	// 人员库ID 。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 识别出的最相似候选人。
	Candidates []*Candidate `json:"Candidates,omitempty" name:"Candidates"`
}

type GroupExDescriptionInfo struct {
	// 人员库自定义描述字段Index，从0开始
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupExDescriptionIndex *uint64 `json:"GroupExDescriptionIndex,omitempty" name:"GroupExDescriptionIndex"`

	// 需要更新的人员库自定义描述字段内容
	GroupExDescription *string `json:"GroupExDescription,omitempty" name:"GroupExDescription"`
}

type GroupInfo struct {
	// 人员库名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员库自定义描述字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions"`

	// 人员库信息备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 人脸识别所用的算法模型版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// Group的创建时间和日期 CreationTimestamp。CreationTimestamp 的值是自 Unix 纪元时间到Group创建时间的毫秒数。 
	// Unix 纪元时间是 1970 年 1 月 1 日星期四，协调世界时 (UTC) 00:00:00。有关更多信息，请参阅 Unix 时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTimestamp *uint64 `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
}

type Hair struct {
	// 头发长度信息。
	// AttributeItem对应的Type为 —— 0：光头，1：短发，2：中发，3：长发，4：绑发。
	Length *AttributeItem `json:"Length,omitempty" name:"Length"`

	// 刘海信息。
	// AttributeItem对应的Type为 —— 0：无刘海，1：有刘海。
	Bang *AttributeItem `json:"Bang,omitempty" name:"Bang"`

	// 头发颜色信息。
	// AttributeItem对应的Type为 —— 0：黑色，1：金色，2：棕色，3：灰白色。
	Color *AttributeItem `json:"Color,omitempty" name:"Color"`
}

type Hat struct {
	// 帽子佩戴状态信息。
	// AttributeItem对应的Type为 —— 0：不戴帽子，1：普通帽子，2：头盔，3：保安帽。
	Style *AttributeItem `json:"Style,omitempty" name:"Style"`

	// 帽子颜色。
	// AttributeItem对应的Type为 —— 0：不戴帽子，1：红色系，2：黄色系，3：蓝色系，4：黑色系，5：灰白色系，6：混色系子。
	Color *AttributeItem `json:"Color,omitempty" name:"Color"`
}

type HeadPose struct {
	// 上下偏移[-30,30]。
	Pitch *int64 `json:"Pitch,omitempty" name:"Pitch"`

	// 左右偏移[-30,30]。
	Yaw *int64 `json:"Yaw,omitempty" name:"Yaw"`

	// 平面旋转[-180,180]。
	Roll *int64 `json:"Roll,omitempty" name:"Roll"`
}

// Predefined struct for user
type ModifyGroupRequestParams struct {
	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员库名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 需要修改的人员库自定义描述字段，key-value
	GroupExDescriptionInfos []*GroupExDescriptionInfo `json:"GroupExDescriptionInfos,omitempty" name:"GroupExDescriptionInfos"`

	// 人员库信息备注
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员库名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 需要修改的人员库自定义描述字段，key-value
	GroupExDescriptionInfos []*GroupExDescriptionInfo `json:"GroupExDescriptionInfos,omitempty" name:"GroupExDescriptionInfos"`

	// 人员库信息备注
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *ModifyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "GroupExDescriptionInfos")
	delete(f, "Tag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGroupResponseParams `json:"Response"`
}

func (r *ModifyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonBaseInfoRequestParams struct {
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 需要修改的人员名称
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 需要修改的人员性别
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`
}

type ModifyPersonBaseInfoRequest struct {
	*tchttp.BaseRequest
	
	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 需要修改的人员名称
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 需要修改的人员性别
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`
}

func (r *ModifyPersonBaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonBaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "PersonName")
	delete(f, "Gender")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPersonBaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonBaseInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPersonBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPersonBaseInfoResponseParams `json:"Response"`
}

func (r *ModifyPersonBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonGroupInfoRequestParams struct {
	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 需要修改的人员描述字段内容，key-value
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitempty" name:"PersonExDescriptionInfos"`
}

type ModifyPersonGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 需要修改的人员描述字段内容，key-value
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitempty" name:"PersonExDescriptionInfos"`
}

func (r *ModifyPersonGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "PersonId")
	delete(f, "PersonExDescriptionInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPersonGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonGroupInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPersonGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPersonGroupInfoResponseParams `json:"Response"`
}

func (r *ModifyPersonGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Mouth struct {
	// 是否张嘴信息。
	// AttributeItem对应的Type为 —— 0：不张嘴，1：张嘴。
	MouthOpen *AttributeItem `json:"MouthOpen,omitempty" name:"MouthOpen"`
}

type PersonExDescriptionInfo struct {
	// 人员描述字段Index，从0开始
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersonExDescriptionIndex *uint64 `json:"PersonExDescriptionIndex,omitempty" name:"PersonExDescriptionIndex"`

	// 需要更新的人员描述字段内容
	PersonExDescription *string `json:"PersonExDescription,omitempty" name:"PersonExDescription"`
}

type PersonGroupInfo struct {
	// 包含此人员的人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员描述字段内容
	PersonExDescriptions []*string `json:"PersonExDescriptions,omitempty" name:"PersonExDescriptions"`
}

type PersonInfo struct {
	// 人员名称
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 人员Id
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人员性别
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 人员描述字段内容
	PersonExDescriptions []*string `json:"PersonExDescriptions,omitempty" name:"PersonExDescriptions"`

	// 包含的人脸照片列表
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds"`

	// 人员的创建时间和日期 CreationTimestamp。CreationTimestamp 的值是自 Unix 纪元时间到Group创建时间的毫秒数。 
	// Unix 纪元时间是 1970 年 1 月 1 日星期四，协调世界时 (UTC) 00:00:00。有关更多信息，请参阅 Unix 时间。
	CreationTimestamp *uint64 `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
}

type Point struct {
	// x坐标
	X *int64 `json:"X,omitempty" name:"X"`

	// Y坐标
	Y *int64 `json:"Y,omitempty" name:"Y"`
}

type Result struct {
	// 识别出的最相似候选人
	Candidates []*Candidate `json:"Candidates,omitempty" name:"Candidates"`

	// 检测出的人脸框位置
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

	// 检测出的人脸图片状态返回码。0 表示正常。 
	// -1601代表不符合图片质量控制要求，此时Candidate内容为空。
	RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
}

type ResultsReturnsByGroup struct {
	// 检测出的人脸框位置。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

	// 识别结果。
	GroupCandidates []*GroupCandidate `json:"GroupCandidates,omitempty" name:"GroupCandidates"`

	// 检测出的人脸图片状态返回码。0 表示正常。 
	// -1601代表不符合图片质量控制要求，此时Candidate内容为空。
	RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
}

// Predefined struct for user
type RevertGroupFaceModelVersionRequestParams struct {
	// 需要回滚的升级任务ID。
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type RevertGroupFaceModelVersionRequest struct {
	*tchttp.BaseRequest
	
	// 需要回滚的升级任务ID。
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *RevertGroupFaceModelVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevertGroupFaceModelVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevertGroupFaceModelVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevertGroupFaceModelVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RevertGroupFaceModelVersionResponse struct {
	*tchttp.BaseResponse
	Response *RevertGroupFaceModelVersionResponseParams `json:"Response"`
}

func (r *RevertGroupFaceModelVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevertGroupFaceModelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchFacesRequestParams struct {
	// 希望搜索的人员库列表，上限100个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 最多识别的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为10。 
	// MaxFaceNum用于，当输入的待识别图片包含多张人脸时，设定要搜索的人脸的数量。 
	// 例：输入的Image或Url中的图片包含多张人脸，设MaxFaceNum=5，则会识别图片中面积最大的5张人脸。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。默认为34。低于34的人脸图片无法被识别。建议设置为80。
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 单张被识别的人脸返回的最相似人员数量。默认值为5，最大值为100。 
	// 例，设MaxFaceNum为1，MaxPersonNum为8，则返回Top8相似的人员信息。
	// 值越大，需要处理的时间越长。建议不要超过10。
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitempty" name:"MaxPersonNum"`

	// 是否返回人员具体信息。0 为关闭，1 为开启。默认为 0。其他非0非1值默认为0
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 出参Score中，只有超过FaceMatchThreshold值的结果才会返回。默认为0。
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

type SearchFacesRequest struct {
	*tchttp.BaseRequest
	
	// 希望搜索的人员库列表，上限100个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 最多识别的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为10。 
	// MaxFaceNum用于，当输入的待识别图片包含多张人脸时，设定要搜索的人脸的数量。 
	// 例：输入的Image或Url中的图片包含多张人脸，设MaxFaceNum=5，则会识别图片中面积最大的5张人脸。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。默认为34。低于34的人脸图片无法被识别。建议设置为80。
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 单张被识别的人脸返回的最相似人员数量。默认值为5，最大值为100。 
	// 例，设MaxFaceNum为1，MaxPersonNum为8，则返回Top8相似的人员信息。
	// 值越大，需要处理的时间越长。建议不要超过10。
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitempty" name:"MaxPersonNum"`

	// 是否返回人员具体信息。0 为关闭，1 为开启。默认为 0。其他非0非1值默认为0
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 出参Score中，只有超过FaceMatchThreshold值的结果才会返回。默认为0。
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *SearchFacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "MaxFaceNum")
	delete(f, "MinFaceSize")
	delete(f, "MaxPersonNum")
	delete(f, "NeedPersonInfo")
	delete(f, "QualityControl")
	delete(f, "FaceMatchThreshold")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchFacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchFacesResponseParams struct {
	// 识别结果。
	Results []*Result `json:"Results,omitempty" name:"Results"`

	// 搜索的人员库中包含的人脸数。
	FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchFacesResponse struct {
	*tchttp.BaseResponse
	Response *SearchFacesResponseParams `json:"Response"`
}

func (r *SearchFacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchFacesReturnsByGroupRequestParams struct {
	// 希望搜索的人员库列表，上限60个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 最多识别的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为10。
	// MaxFaceNum用于，当输入的待识别图片包含多张人脸时，设定要搜索的人脸的数量。
	// 例：输入的Image或Url中的图片包含多张人脸，设MaxFaceNum=5，则会识别图片中面积最大的5张人脸。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。默认为34。低于34将影响搜索精度。建议设置为80。
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 被检测到的人脸，对应最多返回的最相似人员数目。默认值为5，最大值为10。  
	// 例，设MaxFaceNum为3，MaxPersonNumPerGroup为5，GroupIds长度为3，则最多可能返回3*5*3=45个人员。
	MaxPersonNumPerGroup *uint64 `json:"MaxPersonNumPerGroup,omitempty" name:"MaxPersonNumPerGroup"`

	// 是否返回人员具体信息。0 为关闭，1 为开启。默认为 0。其他非0非1值默认为0
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 出参Score中，只有大于等于FaceMatchThreshold值的结果才会返回。
	// 默认为0。
	// 取值范围[0.0,100.0) 。
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

type SearchFacesReturnsByGroupRequest struct {
	*tchttp.BaseRequest
	
	// 希望搜索的人员库列表，上限60个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 最多识别的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为10。
	// MaxFaceNum用于，当输入的待识别图片包含多张人脸时，设定要搜索的人脸的数量。
	// 例：输入的Image或Url中的图片包含多张人脸，设MaxFaceNum=5，则会识别图片中面积最大的5张人脸。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。默认为34。低于34将影响搜索精度。建议设置为80。
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 被检测到的人脸，对应最多返回的最相似人员数目。默认值为5，最大值为10。  
	// 例，设MaxFaceNum为3，MaxPersonNumPerGroup为5，GroupIds长度为3，则最多可能返回3*5*3=45个人员。
	MaxPersonNumPerGroup *uint64 `json:"MaxPersonNumPerGroup,omitempty" name:"MaxPersonNumPerGroup"`

	// 是否返回人员具体信息。0 为关闭，1 为开启。默认为 0。其他非0非1值默认为0
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 出参Score中，只有大于等于FaceMatchThreshold值的结果才会返回。
	// 默认为0。
	// 取值范围[0.0,100.0) 。
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *SearchFacesReturnsByGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFacesReturnsByGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "MaxFaceNum")
	delete(f, "MinFaceSize")
	delete(f, "MaxPersonNumPerGroup")
	delete(f, "NeedPersonInfo")
	delete(f, "QualityControl")
	delete(f, "FaceMatchThreshold")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchFacesReturnsByGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchFacesReturnsByGroupResponseParams struct {
	// 搜索的人员库中包含的人脸数。
	FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

	// 识别结果。
	ResultsReturnsByGroup []*ResultsReturnsByGroup `json:"ResultsReturnsByGroup,omitempty" name:"ResultsReturnsByGroup"`

	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchFacesReturnsByGroupResponse struct {
	*tchttp.BaseResponse
	Response *SearchFacesReturnsByGroupResponseParams `json:"Response"`
}

func (r *SearchFacesReturnsByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFacesReturnsByGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchPersonsRequestParams struct {
	// 希望搜索的人员库列表，上限100个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 最多识别的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为10。
	// MaxFaceNum用于，当输入的待识别图片包含多张人脸时，设定要搜索的人脸的数量。
	// 例：输入的Image或Url中的图片包含多张人脸，设MaxFaceNum=5，则会识别图片中面积最大的5张人脸。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。默认为34。低于34将影响搜索精度。建议设置为80。
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 单张被识别的人脸返回的最相似人员数量。默认值为5，最大值为100。
	// 例，设MaxFaceNum为1，MaxPersonNum为8，则返回Top8相似的人员信息。
	// 值越大，需要处理的时间越长。建议不要超过10。
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitempty" name:"MaxPersonNum"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 出参Score中，只有大于等于FaceMatchThreshold值的结果才会返回。默认为0。取值范围[0.0,100.0) 。
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// 是否返回人员具体信息。0 为关闭，1 为开启。默认为 0。其他非0非1值默认为0
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

type SearchPersonsRequest struct {
	*tchttp.BaseRequest
	
	// 希望搜索的人员库列表，上限100个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 最多识别的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为10。
	// MaxFaceNum用于，当输入的待识别图片包含多张人脸时，设定要搜索的人脸的数量。
	// 例：输入的Image或Url中的图片包含多张人脸，设MaxFaceNum=5，则会识别图片中面积最大的5张人脸。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。默认为34。低于34将影响搜索精度。建议设置为80。
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 单张被识别的人脸返回的最相似人员数量。默认值为5，最大值为100。
	// 例，设MaxFaceNum为1，MaxPersonNum为8，则返回Top8相似的人员信息。
	// 值越大，需要处理的时间越长。建议不要超过10。
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitempty" name:"MaxPersonNum"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 出参Score中，只有大于等于FaceMatchThreshold值的结果才会返回。默认为0。取值范围[0.0,100.0) 。
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// 是否返回人员具体信息。0 为关闭，1 为开启。默认为 0。其他非0非1值默认为0
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *SearchPersonsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPersonsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "MaxFaceNum")
	delete(f, "MinFaceSize")
	delete(f, "MaxPersonNum")
	delete(f, "QualityControl")
	delete(f, "FaceMatchThreshold")
	delete(f, "NeedPersonInfo")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchPersonsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchPersonsResponseParams struct {
	// 识别结果。
	Results []*Result `json:"Results,omitempty" name:"Results"`

	// 搜索的人员库中包含的人员数。若输入图片中所有人脸均不符合质量要求，则返回0。
	PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

	// 人脸识别所用的算法模型版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchPersonsResponse struct {
	*tchttp.BaseResponse
	Response *SearchPersonsResponseParams `json:"Response"`
}

func (r *SearchPersonsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPersonsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchPersonsReturnsByGroupRequestParams struct {
	// 希望搜索的人员库列表，上限60个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 最多识别的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为10。
	// MaxFaceNum用于，当输入的待识别图片包含多张人脸时，设定要搜索的人脸的数量。
	// 例：输入的Image或Url中的图片包含多张人脸，设MaxFaceNum=5，则会识别图片中面积最大的5张人脸。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。默认为34。低于34将影响搜索精度。建议设置为80。
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 被检测到的人脸，对应最多返回的最相似人员数目。默认值为5，最大值为10。  
	// 例，设MaxFaceNum为3，MaxPersonNumPerGroup为5，GroupIds长度为3，则最多可能返回3*5*3=45个人员。
	MaxPersonNumPerGroup *uint64 `json:"MaxPersonNumPerGroup,omitempty" name:"MaxPersonNumPerGroup"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 出参Score中，只有超过FaceMatchThreshold值的结果才会返回。默认为0。
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// 是否返回人员具体信息。0 为关闭，1 为开启。默认为 0。其他非0非1值默认为0
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

type SearchPersonsReturnsByGroupRequest struct {
	*tchttp.BaseRequest
	
	// 希望搜索的人员库列表，上限60个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 最多识别的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为10。
	// MaxFaceNum用于，当输入的待识别图片包含多张人脸时，设定要搜索的人脸的数量。
	// 例：输入的Image或Url中的图片包含多张人脸，设MaxFaceNum=5，则会识别图片中面积最大的5张人脸。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。默认为34。低于34将影响搜索精度。建议设置为80。
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 被检测到的人脸，对应最多返回的最相似人员数目。默认值为5，最大值为10。  
	// 例，设MaxFaceNum为3，MaxPersonNumPerGroup为5，GroupIds长度为3，则最多可能返回3*5*3=45个人员。
	MaxPersonNumPerGroup *uint64 `json:"MaxPersonNumPerGroup,omitempty" name:"MaxPersonNumPerGroup"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 出参Score中，只有超过FaceMatchThreshold值的结果才会返回。默认为0。
	FaceMatchThreshold *float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// 是否返回人员具体信息。0 为关闭，1 为开启。默认为 0。其他非0非1值默认为0
	NeedPersonInfo *int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *SearchPersonsReturnsByGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPersonsReturnsByGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "MaxFaceNum")
	delete(f, "MinFaceSize")
	delete(f, "MaxPersonNumPerGroup")
	delete(f, "QualityControl")
	delete(f, "FaceMatchThreshold")
	delete(f, "NeedPersonInfo")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchPersonsReturnsByGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchPersonsReturnsByGroupResponseParams struct {
	// 搜索的人员库中包含的人员数。若输入图片中所有人脸均不符合质量要求，则返回0。
	PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

	// 识别结果。
	ResultsReturnsByGroup []*ResultsReturnsByGroup `json:"ResultsReturnsByGroup,omitempty" name:"ResultsReturnsByGroup"`

	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchPersonsReturnsByGroupResponse struct {
	*tchttp.BaseResponse
	Response *SearchPersonsReturnsByGroupResponseParams `json:"Response"`
}

func (r *SearchPersonsReturnsByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPersonsReturnsByGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeGroupFaceModelVersionRequestParams struct {
	// 需要升级的人员库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 需要升级至的算法模型版本。默认为最新版本。不可逆向升级
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

type UpgradeGroupFaceModelVersionRequest struct {
	*tchttp.BaseRequest
	
	// 需要升级的人员库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 需要升级至的算法模型版本。默认为最新版本。不可逆向升级
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

func (r *UpgradeGroupFaceModelVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeGroupFaceModelVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "FaceModelVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeGroupFaceModelVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeGroupFaceModelVersionResponseParams struct {
	// 升级任务ID，用于查询、获取升级的进度和结果。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeGroupFaceModelVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeGroupFaceModelVersionResponseParams `json:"Response"`
}

func (r *UpgradeGroupFaceModelVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeGroupFaceModelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeJobInfo struct {
	// 人员库升级任务ID，用于查询、获取升级的进度和结果。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 人员库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 当前算法模型版本。
	FromFaceModelVersion *string `json:"FromFaceModelVersion,omitempty" name:"FromFaceModelVersion"`

	// 目标算法模型版本。
	ToFaceModelVersion *string `json:"ToFaceModelVersion,omitempty" name:"ToFaceModelVersion"`

	// 升级起始时间。 
	// StartTime的值是自 Unix 纪元时间到Group创建时间的毫秒数。 
	// Unix 纪元时间是 1970 年 1 月 1 日星期四，协调世界时 (UTC) 00:00:00。 
	// 有关更多信息，请参阅 Unix 时间。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 0表示升级中，1表示升级完毕，2表示回滚完毕。
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type VerifyFaceRequestParams struct {
	// 待验证的人员ID。人员ID具体信息请参考人员库管理相关接口。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

type VerifyFaceRequest struct {
	*tchttp.BaseRequest
	
	// 待验证的人员ID。人员ID具体信息请参考人员库管理相关接口。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *VerifyFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "QualityControl")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyFaceResponseParams struct {
	// 给定的人脸图片与 PersonId 对应人脸的相似度。若 PersonId 下有多张人脸（Face），返回相似度最大的分数。
	// 
	// 不同算法版本返回的相似度分数不同。
	// 若需要验证两张图片中人脸是否为同一人，3.0版本误识率千分之一对应分数为40分，误识率万分之一对应分数为50分，误识率十万分之一对应分数为60分。 一般超过50分则可认定为同一人。
	// 2.0版本误识率千分之一对应分数为70分，误识率万分之一对应分数为80分，误识率十万分之一对应分数为90分。 一般超过80分则可认定为同一人。
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 是否为同一人判断，固定阈值分数为60分，若想更灵活地调整阈值可取Score参数返回进行判断
	IsMatch *bool `json:"IsMatch,omitempty" name:"IsMatch"`

	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyFaceResponse struct {
	*tchttp.BaseResponse
	Response *VerifyFaceResponseParams `json:"Response"`
}

func (r *VerifyFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyPersonRequestParams struct {
	// 待验证的人员ID。人员ID具体信息请参考人员库管理相关接口。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 图片 base64 数据。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

type VerifyPersonRequest struct {
	*tchttp.BaseRequest
	
	// 待验证的人员ID。人员ID具体信息请参考人员库管理相关接口。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 图片 base64 数据。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 图片质量控制。 
	// 0: 不进行控制； 
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况； 
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况； 
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 
	// 默认 0。 
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl *uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *VerifyPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "QualityControl")
	delete(f, "NeedRotateDetection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyPersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyPersonResponseParams struct {
	// 给定的人脸照片与 PersonId 对应的相似度。若 PersonId 下有多张人脸（Face），会融合多张人脸信息进行验证。
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 是否为同一人的判断。
	IsMatch *bool `json:"IsMatch,omitempty" name:"IsMatch"`

	// 人脸识别所用的算法模型版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyPersonResponse struct {
	*tchttp.BaseResponse
	Response *VerifyPersonResponseParams `json:"Response"`
}

func (r *VerifyPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}