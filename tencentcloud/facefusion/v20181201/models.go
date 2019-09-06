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

package v20181201

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type FaceFusionRequest struct {
	*tchttp.BaseRequest

	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 图片 base64 数据。请确保人脸为正脸，无旋转。若某些手机拍摄后人脸被旋转，请使用图片的 EXIF 信息对图片进行旋转处理；请勿在 base64 数据中包含头部，如“data:image/jpeg;base64,”。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 返回图像方式（url 或 base64) ，二选一。当前仅支持 url 方式，base64 方式后期开放。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 0表示不需要鉴黄，1表示需要鉴黄。2018年12月1号以前创建的活动默认值为0，其他情况默认值为1.
	PornDetect *int64 `json:"PornDetect,omitempty" name:"PornDetect"`

	// 0表示不需要鉴政，1表示需要鉴政。2018年12月1号以前创建的活动默认值为0，其他情况默认值为1。鉴政接口同时会对名人明星进行识别，您可以根据实际需要过滤。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitempty" name:"CelebrityIdentify"`
}

func (r *FaceFusionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FaceFusionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FaceFusionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// RspImgType 为 url 时，返回结果的 url， RspImgType 为 base64 时返回 base64 数据。当前仅支持 url 方式，base64 方式后期开放。
		Image *string `json:"Image,omitempty" name:"Image"`

		// 鉴黄鉴政结果
		ReviewResultSet []*FuseFaceReviewResult `json:"ReviewResultSet,omitempty" name:"ReviewResultSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FaceFusionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FaceFusionResponse) FromJsonString(s string) error {
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

type FuseFaceRequest struct {
	*tchttp.BaseRequest

	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为30天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 人脸图片和待被融合的素材模板图的人脸位置信息。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitempty" name:"MergeInfos" list`

	// 脸型融合比例，数值越高，融合后的脸型越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中脸型参数数值。
	FuseProfileDegree *int64 `json:"FuseProfileDegree,omitempty" name:"FuseProfileDegree"`

	// 五官融合比例，数值越高，融合后的五官越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中五官参数数值。
	FuseFaceDegree *int64 `json:"FuseFaceDegree,omitempty" name:"FuseFaceDegree"`
}

func (r *FuseFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FuseFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FuseFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// RspImgType 为 url 时，返回结果的 url， RspImgType 为 base64 时返回 base64 数据。
		FusedImage *string `json:"FusedImage,omitempty" name:"FusedImage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FuseFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FuseFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FuseFaceReviewDetail struct {

	// 鉴政使用字段, 为职业属性,其他审核结果对应上一级category
	Field *string `json:"Field,omitempty" name:"Field"`

	// 人员名称
	Label *string `json:"Label,omitempty" name:"Label"`

	// 对应识别label的置信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 此字段为保留字段，目前统一返回pass。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
}

type FuseFaceReviewResult struct {

	// 对应的类别名称 porn, politics, terror
	Category *string `json:"Category,omitempty" name:"Category"`

	// 对应子类别状态码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 对应子类别状态码信息描述
	CodeDescription *string `json:"CodeDescription,omitempty" name:"CodeDescription"`

	// 对应识别种类的置信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 此字段为保留字段，目前统一返回pass。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 审核详细内容
	DetailSet []*FuseFaceReviewDetail `json:"DetailSet,omitempty" name:"DetailSet" list`
}

type MergeInfo struct {

	// 输入图片base64
	Image *string `json:"Image,omitempty" name:"Image"`

	// 输入图片url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 上传的图片人脸位置信息（人脸框）
	InputImageFaceRect *FaceRect `json:"InputImageFaceRect,omitempty" name:"InputImageFaceRect"`

	// 控制台上传的素材人脸ID
	TemplateFaceID *string `json:"TemplateFaceID,omitempty" name:"TemplateFaceID"`
}
