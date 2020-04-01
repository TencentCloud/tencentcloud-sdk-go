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

	// 返回图像方式（url 或 base64) ，二选一。url有效期为30天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 历史遗留字段，无需填写。因为融合只需提取人脸特征，不需要鉴黄。
	PornDetect *int64 `json:"PornDetect,omitempty" name:"PornDetect"`

	// 0表示不需要鉴政，1表示需要鉴政。默认值为0。
	// 请注意，鉴政服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
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

		// RspImgType 为 url 时，返回结果的 url， RspImgType 为 base64 时返回 base64 数据。
		Image *string `json:"Image,omitempty" name:"Image"`

		// 鉴政结果
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

	// 用户人脸图片、素材模板图的人脸位置信息。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitempty" name:"MergeInfos" list`

	// 脸型融合比例，数值越高，融合后的脸型越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中脸型参数数值。（换脸版算法暂不支持此参数调整）
	FuseProfileDegree *int64 `json:"FuseProfileDegree,omitempty" name:"FuseProfileDegree"`

	// 五官融合比例，数值越高，融合后的五官越像素材人物。取值范围[0,100] 
	// 若此参数不填写，则使用人脸融合控制台中五官参数数值。（换脸版算法暂不支持此参数调整）
	FuseFaceDegree *int64 `json:"FuseFaceDegree,omitempty" name:"FuseFaceDegree"`

	// 0表示不需要鉴政，1表示需要鉴政。默认值为0。
	// 请注意，鉴政服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitempty" name:"CelebrityIdentify"`
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

		// 鉴政结果。该数组的顺序和请求中mergeinfo的顺序一致，一一对应
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReviewResultSet []*FuseFaceReviewResult `json:"ReviewResultSet,omitempty" name:"ReviewResultSet" list`

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

	// 保留字段
	Field *string `json:"Field,omitempty" name:"Field"`

	// 人员名称
	Label *string `json:"Label,omitempty" name:"Label"`

	// 对应识别label的置信度，分数越高意味涉政可能性越大。 
	// 0到70，Suggestion建议为PASS； 
	// 70到80，Suggestion建议为REVIEW； 
	// 80到100，Suggestion建议为BLOCK。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 识别场景的审核结论：  
	// PASS：正常 
	// REVIEW：疑似  
	// BLOCK：违规
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
}

type FuseFaceReviewResult struct {

	// 保留字段
	Category *string `json:"Category,omitempty" name:"Category"`

	// 状态码， 0为处理成功，其他值为处理失败
	Code *string `json:"Code,omitempty" name:"Code"`

	// 对应状态码信息描述
	CodeDescription *string `json:"CodeDescription,omitempty" name:"CodeDescription"`

	// 保留字段
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 保留字段
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
