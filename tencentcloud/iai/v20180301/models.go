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

	// 图片 base64 数据。支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`
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
	// 10万大小人脸库，若人脸均为类似抓拍照（人脸质量较差）， 
	// 误识率百分之一对应分数为70分，误识率千分之一对应分数为80分，误识率万分之一对应分数为90分； 
	// 若人脸均为类似自拍照（人脸质量较好）， 
	// 误识率百分之一对应分数为60分，误识率千分之一对应分数为70分，误识率万分之一对应分数为80分。 
	// 建议分数不要超过90分。您可以根据实际情况选择合适的分数。
	Score *float64 `json:"Score,omitempty" name:"Score"`
}

type CompareFaceRequest struct {
	*tchttp.BaseRequest

	// A 图片 base64 数据。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	ImageA *string `json:"ImageA,omitempty" name:"ImageA"`

	// B 图片 base64 数据。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	ImageB *string `json:"ImageB,omitempty" name:"ImageB"`

	// A 图片的 Url 。A 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	UrlA *string `json:"UrlA,omitempty" name:"UrlA"`

	// B 图片的 Url 。B 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	UrlB *string `json:"UrlB,omitempty" name:"UrlB"`
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
	// 若需要验证两张图片中人脸是否为同一人，则误识率千分之一对应分数为70分，误识率万分之一对应分数为80分，误识率十万分之一对应分数为90分。  
	// 一般超过80分则可认定为同一人。 
	// 若需要验证两张图片中的人脸是否为同一人，建议使用人脸验证接口。
		Score *float64 `json:"Score,omitempty" name:"Score"`

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

	// 图片 base64 数据。人员人脸总数量不可超过5张。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Images []*string `json:"Images,omitempty" name:"Images" list`

	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 人员人脸总数量不可超过5张。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
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

		// 每张人脸图片添加结果，-1101 代表未检测到人脸，-1102 代表图片解码失败，其他非 0 值代表算法服务异常。
		RetCode []*int64 `json:"RetCode,omitempty" name:"RetCode" list`

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

	// 图片 base64 数据。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`
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

	// 最多处理的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为30。 
	// 此参数用于控制处理待检测图片中的人脸个数，值越小，处理速度越快。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。默认为40。低于此尺寸的人脸不会被检测。
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 图片 base64 数据。支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
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
	// 建议：人脸入库操作建议开启此功能。
	NeedQualityDetection *uint64 `json:"NeedQualityDetection,omitempty" name:"NeedQualityDetection"`
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

		// 人脸信息列表。
		FaceInfos []*FaceInfo `json:"FaceInfos,omitempty" name:"FaceInfos" list`

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

	// 图片 base64 数据（图片的宽高比请接近3:4，不符合宽高比的图片返回的分值不具备参考意义）。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// （图片的宽高比请接近 3:4，不符合宽高比的图片返回的分值不具备参考意义） 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`
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
		Score *float64 `json:"Score,omitempty" name:"Score"`

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

type FaceAttributesInfo struct {

	// 性别 [0(female)~100(male)]。 NeedFaceAttributes 不为 1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 年龄 [0~100]。NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 微笑[0(normal)~50(smile)~100(laugh)]。NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Expression *int64 `json:"Expression,omitempty" name:"Expression"`

	// 是否有眼镜 [true,false]。NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。
	Glass *bool `json:"Glass,omitempty" name:"Glass"`

	// 上下偏移[-30,30]。NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。 
	// 建议：人脸入库选择[-10,10]的图片。
	Pitch *int64 `json:"Pitch,omitempty" name:"Pitch"`

	// 左右偏移[-30,30]。 NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。 
	// 建议：人脸入库选择[-10,10]的图片。
	Yaw *int64 `json:"Yaw,omitempty" name:"Yaw"`

	// 平面旋转[-180,180]。 NeedFaceAttributes 不为1 或检测超过 5 张人脸时，此参数仍返回，但不具备参考意义。  
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

	// 人脸框左上角纵坐标。 
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

	// 描述左测眼睛轮廓的 8 点。
	LeftEye []*Point `json:"LeftEye,omitempty" name:"LeftEye" list`

	// 描述右测眼睛轮廓的 8 点。
	RightEye []*Point `json:"RightEye,omitempty" name:"RightEye" list`

	// 描述左测眉毛轮廓的 8 点。
	LeftEyeBrow []*Point `json:"LeftEyeBrow,omitempty" name:"LeftEyeBrow" list`

	// 描述右测眉毛轮廓的 8 点。
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

		// 包含的人脸图片列表
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

type GroupExDescriptionInfo struct {

	// 人员库自定义描述字段Index，从0开始
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
}

type SearchFacesRequest struct {
	*tchttp.BaseRequest

	// 希望搜索的人员库列表，上限100个。
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`

	// 图片 base64 数据。支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 最多处理的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为10。 
	// MaxFaceNum用于，当待识别图片包含多张人脸时，要搜索的人脸数量。 
	// 当 MaxFaceNum 不为1时，设MaxFaceNum=M，则实际上是 M:N 的人脸搜索。
	MaxFaceNum *uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。默认为80。低于40将影响搜索精度。建议设置为80。
	MinFaceSize *uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 被检测到的人脸，对应最多返回的最相似人员数目。默认值为5，最大值为10。  
	// 例，设MaxFaceNum为3，MaxPersonNum为5，则最多可能返回3*5=15个人员。
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitempty" name:"MaxPersonNum"`
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

type VerifyFaceRequest struct {
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
		Score *float64 `json:"Score,omitempty" name:"Score"`

		// 是否为同一人的判断。
		IsMatch *bool `json:"IsMatch,omitempty" name:"IsMatch"`

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
