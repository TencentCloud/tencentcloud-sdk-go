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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

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

	// 人脸识别服务所用的算法模型版本。目前入参支持 “2.0”和“3.0“ 两个输入。  
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。  
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用最新版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection *uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

func (r *AnalyzeFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AnalyzeFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AnalyzeFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求的图片宽度。
		ImageWidth *uint64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

		// 请求的图片高度。
		ImageHeight *uint64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

		// 五官定位（人脸关键点）具体信息。
		FaceShapeSet []*FaceShape `json:"FaceShapeSet,omitempty" name:"FaceShapeSet" list`

		// 人脸识别所用的算法模型版本。
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AnalyzeFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AnalyzeFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Candidate struct {

	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人脸ID
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
	PersonGroupInfos []*PersonGroupInfo `json:"PersonGroupInfos,omitempty" name:"PersonGroupInfos" list`
}

type CheckSimilarPersonRequest struct {
	*tchttp.BaseRequest

	// 待整理的人员库列表。 
	// 人员库总人数不可超过200万，人员库个数不可超过10个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// 人员查重整理力度的控制。
	// 1：力度较高的档案整理，能够消除更多的重复身份，对应稍高的非重复身份误清除率；
	// 2：力度较低的档案整理，非重复身份的误清除率较低，对应稍低的重复身份消除率。
	UniquePersonControl *int64 `json:"UniquePersonControl,omitempty" name:"UniquePersonControl"`
}

func (r *CheckSimilarPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckSimilarPersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckSimilarPersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查重任务ID，用于查询、获取查重的进度和结果。
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckSimilarPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckSimilarPersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 人脸识别服务所用的算法模型版本。目前入参支持 “2.0”和“3.0“ 两个输入。 
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。 
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

func (r *CompareFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompareFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *CompareFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompareFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyPersonRequest struct {
	*tchttp.BaseRequest

	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 待加入的人员库列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`
}

func (r *CopyPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyPersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyPersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功加入的人员库数量
		SucGroupNum *uint64 `json:"SucGroupNum,omitempty" name:"SucGroupNum"`

		// 成功加入的人员库列表
		SucGroupIds []*string `json:"SucGroupIds,omitempty" name:"SucGroupIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyPersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFaceRequest struct {
	*tchttp.BaseRequest

	// 人员ID。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 人员人脸总数量不可超过5张。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Images []*string `json:"Images,omitempty" name:"Images" list`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 人员人脸总数量不可超过5张。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

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

func (r *CreateFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 加入成功的人脸数量
		SucFaceNum *uint64 `json:"SucFaceNum,omitempty" name:"SucFaceNum"`

		// 加入成功的人脸ID列表
		SucFaceIds []*string `json:"SucFaceIds,omitempty" name:"SucFaceIds" list`

		// 每张人脸图片添加结果，-1101 代表未检测到人脸，-1102 代表图片解码失败， 
	// -1601代表不符合图片质量控制要求, -1604 代表人脸相似度没有超过FaceMatchThreshold。 
	// 其他非 0 值代表算法服务异常。 
	// RetCode的顺序和入参中 Images 或 Urls 的顺序一致。
		RetCode []*int64 `json:"RetCode,omitempty" name:"RetCode" list`

		// 加入成功的人脸索引。索引顺序和入参中 Images 或 Urls 的顺序一致。 
	// 例， Urls 中 有 3 个 url，第二个 url 失败，则 SucIndexes 值为 [0,2] 。
		SucIndexes []*uint64 `json:"SucIndexes,omitempty" name:"SucIndexes" list`

		// 加入成功的人脸框位置。顺序和入参中 Images 或 Urls 的顺序一致。
		SucFaceRects []*FaceRect `json:"SucFaceRects,omitempty" name:"SucFaceRects" list`

		// 人脸识别所用的算法模型版本。
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions" list`

	// 人员库信息备注，[0，40]个字符。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 人脸识别服务所用的算法模型版本。目前入参支持 “2.0”和“3.0“ 两个输入。
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

func (r *CreateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人脸识别所用的算法模型版本。
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitempty" name:"PersonExDescriptionInfos" list`

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

func (r *CreatePersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *CreatePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFaceRequest struct {
	*tchttp.BaseRequest

	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 待删除的人脸ID列表
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds" list`
}

func (r *DeleteFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除成功的人脸数量
		SucDeletedNum *uint64 `json:"SucDeletedNum,omitempty" name:"SucDeletedNum"`

		// 删除成功的人脸ID列表
		SucFaceIds []*string `json:"SucFaceIds,omitempty" name:"SucFaceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeletePersonFromGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePersonFromGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePersonFromGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePersonFromGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeletePersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DetectFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求的图片宽度。
		ImageWidth *int64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

		// 请求的图片高度。
		ImageHeight *int64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

		// 人脸信息列表。包含人脸坐标信息、属性信息（若需要）、质量分信息（若需要）。
		FaceInfos []*FaceInfo `json:"FaceInfos,omitempty" name:"FaceInfos" list`

		// 人脸识别所用的算法模型版本。
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 人脸识别服务所用的算法模型版本。目前入参支持 “2.0”和“3.0“ 两个输入。  
	// 2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。 
	// 不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

func (r *DetectLiveFaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectLiveFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectLiveFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *DetectLiveFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectLiveFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EstimateCheckSimilarPersonCostTimeRequest struct {
	*tchttp.BaseRequest

	// 待整理的人员库列表。 
	// 人员库总人数不可超过200万，人员库个数不可超过10个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`
}

func (r *EstimateCheckSimilarPersonCostTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EstimateCheckSimilarPersonCostTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EstimateCheckSimilarPersonCostTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人员查重任务预估需要耗费时间。 单位为分钟。
		EstimatedTimeCost *uint64 `json:"EstimatedTimeCost,omitempty" name:"EstimatedTimeCost"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EstimateCheckSimilarPersonCostTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EstimateCheckSimilarPersonCostTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	FaceProfile []*Point `json:"FaceProfile,omitempty" name:"FaceProfile" list`

	// 描述左侧眼睛轮廓的 8 点。
	LeftEye []*Point `json:"LeftEye,omitempty" name:"LeftEye" list`

	// 描述右侧眼睛轮廓的 8 点。
	RightEye []*Point `json:"RightEye,omitempty" name:"RightEye" list`

	// 描述左侧眉毛轮廓的 8 点。
	LeftEyeBrow []*Point `json:"LeftEyeBrow,omitempty" name:"LeftEyeBrow" list`

	// 描述右侧眉毛轮廓的 8 点。
	RightEyeBrow []*Point `json:"RightEyeBrow,omitempty" name:"RightEyeBrow" list`

	// 描述嘴巴轮廓的 22 点。
	Mouth []*Point `json:"Mouth,omitempty" name:"Mouth" list`

	// 描述鼻子轮廓的 13 点。
	Nose []*Point `json:"Nose,omitempty" name:"Nose" list`

	// 左瞳孔轮廓的 1 个点。
	LeftPupil []*Point `json:"LeftPupil,omitempty" name:"LeftPupil" list`

	// 右瞳孔轮廓的 1 个点。
	RightPupil []*Point `json:"RightPupil,omitempty" name:"RightPupil" list`
}

type GetCheckSimilarPersonJobIdListRequest struct {
	*tchttp.BaseRequest

	// 起始序号，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetCheckSimilarPersonJobIdListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCheckSimilarPersonJobIdListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCheckSimilarPersonJobIdListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人员查重任务信息列表。
		JobIdInfos []*JobIdInfo `json:"JobIdInfos,omitempty" name:"JobIdInfos" list`

		// 查重任务总数量。
		JobIdNum *uint64 `json:"JobIdNum,omitempty" name:"JobIdNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCheckSimilarPersonJobIdListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCheckSimilarPersonJobIdListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *GetGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人员库名称
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// 人员库ID
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 人员库自定义描述字段
		GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions" list`

		// 人员库信息备注
		Tag *string `json:"Tag,omitempty" name:"Tag"`

		// 人脸识别所用的算法模型版本。
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// Group的创建时间和日期 CreationTimestamp。CreationTimestamp 的值是自 Unix 纪元时间到Group创建时间的毫秒数。
		CreationTimestamp *uint64 `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *GetGroupListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的人员库信息
		GroupInfos []*GroupInfo `json:"GroupInfos,omitempty" name:"GroupInfos" list`

		// 人员库总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *GetPersonBaseInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPersonBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人员名称
		PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

		// 人员性别
		Gender *int64 `json:"Gender,omitempty" name:"Gender"`

		// 包含的人脸 ID 列表
		FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPersonBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPersonBaseInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *GetPersonGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPersonGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 包含此人员的人员库及描述字段内容列表
		PersonGroupInfos []*PersonGroupInfo `json:"PersonGroupInfos,omitempty" name:"PersonGroupInfos" list`

		// 人员库总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`

		// 人脸识别服务所用的算法模型版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPersonGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPersonGroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *GetPersonListNumRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPersonListNumResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人员数量
		PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

		// 人脸数量
		FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPersonListNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPersonListNumResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *GetPersonListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPersonListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的人员信息
		PersonInfos []*PersonInfo `json:"PersonInfos,omitempty" name:"PersonInfos" list`

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
	} `json:"Response"`
}

func (r *GetPersonListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPersonListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSimilarPersonResultRequest struct {
	*tchttp.BaseRequest

	// 查重任务ID，用于查询、获取查重的进度和结果。
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *GetSimilarPersonResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSimilarPersonResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSimilarPersonResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查重任务完成进度。取值[0.0，100.0]。当且仅当值为100时，SimilarPersons才有意义。
		Progress *float64 `json:"Progress,omitempty" name:"Progress"`

		// 疑似同一人的人员信息文件临时下载链接， 有效时间为5分钟，结果文件实际保存90天。
	// 文件内容由 SimilarPerson 的数组组成。
		SimilarPersonsUrl *string `json:"SimilarPersonsUrl,omitempty" name:"SimilarPersonsUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSimilarPersonResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSimilarPersonResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GroupCandidate struct {

	// 人员库ID 。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 识别出的最相似候选人。
	Candidates []*Candidate `json:"Candidates,omitempty" name:"Candidates" list`
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
	GroupExDescriptions []*string `json:"GroupExDescriptions,omitempty" name:"GroupExDescriptions" list`

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

type JobIdInfo struct {

	// 查重任务ID，用于查询、获取查重的进度和结果。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 查重起始时间。 
	// StartTime的值是自 Unix 纪元时间到Group创建时间的毫秒数。 
	// Unix 纪元时间是 1970 年 1 月 1 日星期四，协调世界时 (UTC) 00:00:00。 
	// 有关更多信息，请参阅 Unix 时间。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查重任务是否已完成。0: 成功 1: 未完成 2: 失败
	JobStatus *int64 `json:"JobStatus,omitempty" name:"JobStatus"`
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest

	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员库名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 需要修改的人员库自定义描述字段，key-value
	GroupExDescriptionInfos []*GroupExDescriptionInfo `json:"GroupExDescriptionInfos,omitempty" name:"GroupExDescriptionInfos" list`

	// 人员库信息备注
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *ModifyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyPersonBaseInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPersonBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonBaseInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 人员库ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人员ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 需要修改的人员描述字段内容，key-value
	PersonExDescriptionInfos []*PersonExDescriptionInfo `json:"PersonExDescriptionInfos,omitempty" name:"PersonExDescriptionInfos" list`
}

func (r *ModifyPersonGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPersonGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonGroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	PersonExDescriptions []*string `json:"PersonExDescriptions,omitempty" name:"PersonExDescriptions" list`
}

type PersonInfo struct {

	// 人员名称
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 人员Id
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人员性别
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 人员描述字段内容
	PersonExDescriptions []*string `json:"PersonExDescriptions,omitempty" name:"PersonExDescriptions" list`

	// 包含的人脸照片列表
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds" list`

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
	Candidates []*Candidate `json:"Candidates,omitempty" name:"Candidates" list`

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
	GroupCandidates []*GroupCandidate `json:"GroupCandidates,omitempty" name:"GroupCandidates" list`

	// 检测出的人脸图片状态返回码。0 表示正常。 
	// -1601代表不符合图片质量控制要求，此时Candidate内容为空。
	RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`
}

type SearchFacesRequest struct {
	*tchttp.BaseRequest

	// 希望搜索的人员库列表，上限100个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

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

func (r *SearchFacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchFacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别结果。
		Results []*Result `json:"Results,omitempty" name:"Results" list`

		// 搜索的人员库中包含的人脸数。
		FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

		// 人脸识别所用的算法模型版本。
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchFacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchFacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchFacesReturnsByGroupRequest struct {
	*tchttp.BaseRequest

	// 希望搜索的人员库列表，上限60个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

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
	// 例，设MaxFaceNum为3，MaxPersonNum为5，则最多可能返回3*5=15个人员。
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

func (r *SearchFacesReturnsByGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchFacesReturnsByGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 搜索的人员库中包含的人脸数。
		FaceNum *uint64 `json:"FaceNum,omitempty" name:"FaceNum"`

		// 识别结果。
		ResultsReturnsByGroup []*ResultsReturnsByGroup `json:"ResultsReturnsByGroup,omitempty" name:"ResultsReturnsByGroup" list`

		// 人脸识别所用的算法模型版本。
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchFacesReturnsByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchFacesReturnsByGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchPersonsRequest struct {
	*tchttp.BaseRequest

	// 希望搜索的人员库列表，上限100个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

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

func (r *SearchPersonsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchPersonsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别结果。
		Results []*Result `json:"Results,omitempty" name:"Results" list`

		// 搜索的人员库中包含的人员数。若输入图片中所有人脸均不符合质量要求，则返回0。
		PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

		// 人脸识别所用的算法模型版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchPersonsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchPersonsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchPersonsReturnsByGroupRequest struct {
	*tchttp.BaseRequest

	// 希望搜索的人员库列表，上限60个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

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

func (r *SearchPersonsReturnsByGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchPersonsReturnsByGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 搜索的人员库中包含的人员数。若输入图片中所有人脸均不符合质量要求，则返回0。
		PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

		// 识别结果。
		ResultsReturnsByGroup []*ResultsReturnsByGroup `json:"ResultsReturnsByGroup,omitempty" name:"ResultsReturnsByGroup" list`

		// 人脸识别所用的算法模型版本。
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchPersonsReturnsByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchPersonsReturnsByGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *VerifyFaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyFaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 给定的人脸图片与 PersonId 对应人脸的相似度。若 PersonId 下有多张人脸（Face），返回相似度最大的分数。
	// 
	// 不同算法版本返回的相似度分数不同。
	// 若需要验证两张图片中人脸是否为同一人，3.0版本误识率千分之一对应分数为40分，误识率万分之一对应分数为50分，误识率十万分之一对应分数为60分。 一般超过50分则可认定为同一人。
	// 2.0版本误识率千分之一对应分数为70分，误识率万分之一对应分数为80分，误识率十万分之一对应分数为90分。 一般超过80分则可认定为同一人。
		Score *float64 `json:"Score,omitempty" name:"Score"`

		// 是否为同一人的判断。
		IsMatch *bool `json:"IsMatch,omitempty" name:"IsMatch"`

		// 人脸识别所用的算法模型版本。
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyFaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyFaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyPersonRequest struct {
	*tchttp.BaseRequest

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

	// 待验证的人员ID。人员ID具体信息请参考人员库管理相关接口。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

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

func (r *VerifyPersonRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyPersonResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 给定的人脸照片与 PersonId 对应的相似度。若 PersonId 下有多张人脸（Face），会融合多张人脸信息进行验证。
		Score *float64 `json:"Score,omitempty" name:"Score"`

		// 是否为同一人的判断。
		IsMatch *bool `json:"IsMatch,omitempty" name:"IsMatch"`

		// 人脸识别所用的算法模型版本。
		FaceModelVersion *string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyPersonResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
