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

package v20190529

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AssessQualityRequestParams struct {
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果。 
	// • 长宽比：长边：短边<5。
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过Base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要Base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type AssessQualityRequest struct {
	*tchttp.BaseRequest
	
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果。 
	// • 长宽比：长边：短边<5。
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过Base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要Base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *AssessQualityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessQualityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssessQualityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssessQualityResponseParams struct {
	// 取值为TRUE或FALSE，TRUE为长图，FALSE为正常图，长图定义为长宽比大于等于3或小于等于1/3的图片。
	LongImage *bool `json:"LongImage,omitempty" name:"LongImage"`

	// 取值为TRUE或FALSE，TRUE为黑白图，FALSE为否。黑白图即灰度图，指红绿蓝三个通道都是以灰度色阶显示的图片，并非视觉上的“黑白图片”。
	BlackAndWhite *bool `json:"BlackAndWhite,omitempty" name:"BlackAndWhite"`

	// 取值为TRUE或FALSE，TRUE为小图，FALSE为否, 小图定义为最长边小于179像素的图片。当一张图片被判断为小图时，不建议做推荐和展示，其他字段统一输出为0或FALSE。
	SmallImage *bool `json:"SmallImage,omitempty" name:"SmallImage"`

	// 取值为TRUE或FALSE，TRUE为大图，FALSE为否，定义为最短边大于1000像素的图片
	BigImage *bool `json:"BigImage,omitempty" name:"BigImage"`

	// 取值为TRUE或FALSE，TRUE为纯色图或纯文字图，即没有内容或只有简单内容的图片，FALSE为正常图片。
	PureImage *bool `json:"PureImage,omitempty" name:"PureImage"`

	// 综合评分。图像清晰度的得分，对图片的噪声、曝光、模糊、压缩等因素进行综合评估，取值为[0, 100]，值越大，越清晰。一般大于50为较清晰图片，标准可以自行把握。
	ClarityScore *int64 `json:"ClarityScore,omitempty" name:"ClarityScore"`

	// 综合评分。图像美观度得分， 从构图、色彩等多个艺术性维度评价图片，取值为[0, 100]，值越大，越美观。一般大于50为较美观图片，标准可以自行把握。
	AestheticScore *int64 `json:"AestheticScore,omitempty" name:"AestheticScore"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssessQualityResponse struct {
	*tchttp.BaseResponse
	Response *AssessQualityResponseParams `json:"Response"`
}

func (r *AssessQualityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessQualityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Attribute struct {
	// 属性
	Type *string `json:"Type,omitempty" name:"Type"`

	// 属性详情
	Details *string `json:"Details,omitempty" name:"Details"`
}

type AttributesForBody struct {
	// 人体框。当不开启人体检测时，内部参数默认为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rect *ImageRect `json:"Rect,omitempty" name:"Rect"`

	// 人体检测置信度。取值0-1之间，当不开启人体检测开关时默认为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectConfidence *float64 `json:"DetectConfidence,omitempty" name:"DetectConfidence"`

	// 属性信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attributes []*BodyAttributes `json:"Attributes,omitempty" name:"Attributes"`
}

type BodyAttributes struct {
	// 属性值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 置信度，取值0-1之间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 属性名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type Box struct {
	// 图像主体区域。
	Rect *ImageRect `json:"Rect,omitempty" name:"Rect"`

	// 置信度。
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 主体区域类目ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
}

type CarPlateContent struct {
	// 车牌信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Plate *string `json:"Plate,omitempty" name:"Plate"`

	// 车牌颜色。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Color *string `json:"Color,omitempty" name:"Color"`

	// 车牌类型，包含：0普通蓝牌，1双层黄牌，2单层黄牌，3新能源车牌，4使馆车牌，5领馆车牌，6澳门车牌，7香港车牌，8警用车牌，9教练车牌，10武警车牌，11军用车牌   -2遮挡污损模糊车牌/异常   其他无牌
	// 注意：
	// 此字段可能返回 null，表示取不到有效值。
	// 此字段结果遮挡污损模糊车牌/异常：包含PlateStatus参数的“遮挡污损模糊车牌”，针对车牌异常，建议参考此字段，更全面
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 车牌在图片中的坐标信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlateLocation []*Coord `json:"PlateLocation,omitempty" name:"PlateLocation"`

	// 判断车牌是否遮挡：“遮挡污损模糊车牌”和"正常车牌"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlateStatus *string `json:"PlateStatus,omitempty" name:"PlateStatus"`

	// 车牌遮挡的置信度，0-100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlateStatusConfidence *int64 `json:"PlateStatusConfidence,omitempty" name:"PlateStatusConfidence"`

	// 车牌角度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlateAngle *float64 `json:"PlateAngle,omitempty" name:"PlateAngle"`
}

type CarTagItem struct {
	// 车系
	Serial *string `json:"Serial,omitempty" name:"Serial"`

	// 车辆品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 车辆类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 车辆颜色
	Color *string `json:"Color,omitempty" name:"Color"`

	// 车系置信度，0-100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 年份，没识别出年份的时候返回0
	Year *int64 `json:"Year,omitempty" name:"Year"`

	// 车辆在图片中的坐标信息
	CarLocation []*Coord `json:"CarLocation,omitempty" name:"CarLocation"`

	// 车牌信息，仅车辆识别（增强版）支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlateContent *CarPlateContent `json:"PlateContent,omitempty" name:"PlateContent"`

	// 车牌信息置信度，0-100，仅车辆识别（增强版）支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlateConfidence *int64 `json:"PlateConfidence,omitempty" name:"PlateConfidence"`

	// 车辆类型置信度，0-100，仅车辆识别（增强版）支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeConfidence *int64 `json:"TypeConfidence,omitempty" name:"TypeConfidence"`

	// 车辆颜色置信度，0-100，仅车辆识别（增强版）支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColorConfidence *int64 `json:"ColorConfidence,omitempty" name:"ColorConfidence"`

	// 车辆朝向，仅车辆识别（增强版）支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	Orientation *string `json:"Orientation,omitempty" name:"Orientation"`

	// 车辆朝向置信度，0-100，仅车辆识别（增强版）支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrientationConfidence *int64 `json:"OrientationConfidence,omitempty" name:"OrientationConfidence"`
}

type ColorInfo struct {
	// RGB颜色值（16进制），例如：291A18。
	Color *string `json:"Color,omitempty" name:"Color"`

	// 当前颜色标签所占比例。
	Percentage *float64 `json:"Percentage,omitempty" name:"Percentage"`

	// 颜色标签。蜜柚色，米驼色等。
	Label *string `json:"Label,omitempty" name:"Label"`
}

type Coord struct {
	// 横坐标x
	X *int64 `json:"X,omitempty" name:"X"`

	// 纵坐标y
	Y *int64 `json:"Y,omitempty" name:"Y"`
}

// Predefined struct for user
type CreateGroupRequestParams struct {
	// 图库ID，不可重复，仅支持字母、数字和下划线。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 图库名称描述。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 图片库可容纳的最大图片特征条数，一张图片对应一条图片特征数据，不支持修改。
	// 单个图片库容量最大可达亿级，达到容量限制后继续创建图片将会报错。
	// 注意，包月计费下支持绑定的最小库容量为500万。
	MaxCapacity *uint64 `json:"MaxCapacity,omitempty" name:"MaxCapacity"`

	// 图库简介。
	Brief *string `json:"Brief,omitempty" name:"Brief"`

	// 访问限制默认为10qps，如需扩容请联系[在线客服](https://cloud.tencent.com/online-service)申请。
	MaxQps *uint64 `json:"MaxQps,omitempty" name:"MaxQps"`

	// 图库类型，用于决定图像搜索的服务类型和算法版本，默认为4。
	// GroupType不支持修改，若不确定适用的服务类型，建议先对不同类型分别小规模测试后再开始正式使用。
	// 参数取值：
	// 4：通用图像搜索1.0版。
	// 7：商品图像搜索2.0升级版。
	// 5：商品图像搜索1.0版。
	// 6：图案花纹搜索1.0版。
	// 1 - 3：通用图像搜索旧版，不推荐使用。
	GroupType *uint64 `json:"GroupType,omitempty" name:"GroupType"`
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest
	
	// 图库ID，不可重复，仅支持字母、数字和下划线。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 图库名称描述。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 图片库可容纳的最大图片特征条数，一张图片对应一条图片特征数据，不支持修改。
	// 单个图片库容量最大可达亿级，达到容量限制后继续创建图片将会报错。
	// 注意，包月计费下支持绑定的最小库容量为500万。
	MaxCapacity *uint64 `json:"MaxCapacity,omitempty" name:"MaxCapacity"`

	// 图库简介。
	Brief *string `json:"Brief,omitempty" name:"Brief"`

	// 访问限制默认为10qps，如需扩容请联系[在线客服](https://cloud.tencent.com/online-service)申请。
	MaxQps *uint64 `json:"MaxQps,omitempty" name:"MaxQps"`

	// 图库类型，用于决定图像搜索的服务类型和算法版本，默认为4。
	// GroupType不支持修改，若不确定适用的服务类型，建议先对不同类型分别小规模测试后再开始正式使用。
	// 参数取值：
	// 4：通用图像搜索1.0版。
	// 7：商品图像搜索2.0升级版。
	// 5：商品图像搜索1.0版。
	// 6：图案花纹搜索1.0版。
	// 1 - 3：通用图像搜索旧版，不推荐使用。
	GroupType *uint64 `json:"GroupType,omitempty" name:"GroupType"`
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
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "MaxCapacity")
	delete(f, "Brief")
	delete(f, "MaxQps")
	delete(f, "GroupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupResponseParams struct {
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
type CreateImageRequestParams struct {
	// 图库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 物品ID，最多支持64个字符。 
	// 一个物品ID可以包含多张图片，若EntityId已存在，则对其追加图片。同一个EntityId，最大支持10张图。
	EntityId *string `json:"EntityId,omitempty" name:"EntityId"`

	// 图片名称，最多支持64个字符， 
	// PicName唯一确定一张图片，具有唯一性。
	PicName *string `json:"PicName,omitempty" name:"PicName"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。  
	// ImageUrl和ImageBase64必须提供一个，如果都提供，只使用ImageUrl。
	// 图片限制：
	// • 图片格式：支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// • 图片大小：对应图片 base64 编码后大小不可超过5M。图片分辨率不超过4096\*4096。
	// • 如果在商品图像搜索中开启主体识别，分辨率不超过2000\*2000，图片长宽比小于10。
	// 建议：
	// • 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的Url速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片自定义备注内容，最多支持4096个字符，查询时原样带回。
	CustomContent *string `json:"CustomContent,omitempty" name:"CustomContent"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 图片限制：
	// • 图片格式：支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// • 图片大小：base64 编码后大小不可超过5M。图片分辨率不超过4096\*4096。
	// • 如果在商品图像搜索中开启主体识别，分辨率不超过2000\*2000，图片长宽比小于10。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片自定义标签，最多不超过10个，格式为JSON。
	Tags *string `json:"Tags,omitempty" name:"Tags"`

	// 是否需要启用主体识别，默认为**TRUE**。
	// • 为**TRUE**时，启用主体识别，返回主体信息。若没有指定**ImageRect**，自动提取最大面积主体创建图片并进行主体识别。主体识别结果可在**Response**中获取。
	// • 为**FALSE**时，不启用主体识别，不返回主体信息。若没有指定**ImageRect**，以整张图创建图片。
	// **<font color=#1E90FF>注意：仅服务类型为商品图像搜索时才生效。</font>**
	EnableDetect *bool `json:"EnableDetect,omitempty" name:"EnableDetect"`

	// 图像类目ID。
	// 若设置类目ID，提取以下类目的主体创建图片。
	// 类目取值说明：
	// 0：上衣。
	// 1：裙装。
	// 2：下装。
	// 3：包。
	// 4：鞋。
	// 5：配饰。
	// **<font color=#1E90FF>注意：仅服务类型为商品图像搜索时才生效。</font>**
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 图像主体区域。
	// 若设置主体区域，提取指定的区域创建图片。
	ImageRect *Rect `json:"ImageRect,omitempty" name:"ImageRect"`
}

type CreateImageRequest struct {
	*tchttp.BaseRequest
	
	// 图库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 物品ID，最多支持64个字符。 
	// 一个物品ID可以包含多张图片，若EntityId已存在，则对其追加图片。同一个EntityId，最大支持10张图。
	EntityId *string `json:"EntityId,omitempty" name:"EntityId"`

	// 图片名称，最多支持64个字符， 
	// PicName唯一确定一张图片，具有唯一性。
	PicName *string `json:"PicName,omitempty" name:"PicName"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。  
	// ImageUrl和ImageBase64必须提供一个，如果都提供，只使用ImageUrl。
	// 图片限制：
	// • 图片格式：支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// • 图片大小：对应图片 base64 编码后大小不可超过5M。图片分辨率不超过4096\*4096。
	// • 如果在商品图像搜索中开启主体识别，分辨率不超过2000\*2000，图片长宽比小于10。
	// 建议：
	// • 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的Url速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片自定义备注内容，最多支持4096个字符，查询时原样带回。
	CustomContent *string `json:"CustomContent,omitempty" name:"CustomContent"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 图片限制：
	// • 图片格式：支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// • 图片大小：base64 编码后大小不可超过5M。图片分辨率不超过4096\*4096。
	// • 如果在商品图像搜索中开启主体识别，分辨率不超过2000\*2000，图片长宽比小于10。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片自定义标签，最多不超过10个，格式为JSON。
	Tags *string `json:"Tags,omitempty" name:"Tags"`

	// 是否需要启用主体识别，默认为**TRUE**。
	// • 为**TRUE**时，启用主体识别，返回主体信息。若没有指定**ImageRect**，自动提取最大面积主体创建图片并进行主体识别。主体识别结果可在**Response**中获取。
	// • 为**FALSE**时，不启用主体识别，不返回主体信息。若没有指定**ImageRect**，以整张图创建图片。
	// **<font color=#1E90FF>注意：仅服务类型为商品图像搜索时才生效。</font>**
	EnableDetect *bool `json:"EnableDetect,omitempty" name:"EnableDetect"`

	// 图像类目ID。
	// 若设置类目ID，提取以下类目的主体创建图片。
	// 类目取值说明：
	// 0：上衣。
	// 1：裙装。
	// 2：下装。
	// 3：包。
	// 4：鞋。
	// 5：配饰。
	// **<font color=#1E90FF>注意：仅服务类型为商品图像搜索时才生效。</font>**
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 图像主体区域。
	// 若设置主体区域，提取指定的区域创建图片。
	ImageRect *Rect `json:"ImageRect,omitempty" name:"ImageRect"`
}

func (r *CreateImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "EntityId")
	delete(f, "PicName")
	delete(f, "ImageUrl")
	delete(f, "CustomContent")
	delete(f, "ImageBase64")
	delete(f, "Tags")
	delete(f, "EnableDetect")
	delete(f, "CategoryId")
	delete(f, "ImageRect")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageResponseParams struct {
	// 输入图的主体信息。
	// 若启用主体识别且在请求中指定了类目ID或主体区域，以指定的主体为准。若启用主体识别且没有指定，以最大面积主体为准。
	// **<font color=#1E90FF>注意：仅服务类型为商品图像搜索时才生效。</font>**
	// 注意：此字段可能返回 null，表示取不到有效值。
	Object *ObjectInfo `json:"Object,omitempty" name:"Object"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateImageResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageResponseParams `json:"Response"`
}

func (r *CreateImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CropImageRequestParams struct {
	// 需要裁剪区域的宽度，与Height共同组成所需裁剪的图片宽高比例。
	// 输入数字请大于0、小于图片宽度的像素值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 需要裁剪区域的高度，与Width共同组成所需裁剪的图片宽高比例。
	// 输入数字请大于0、小于图片高度的像素值。
	// 宽高比例（Width : Height）会简化为最简分数，即如果Width输入10、Height输入20，会简化为1：2。
	// Width : Height建议取值在[1, 2.5]之间，超过这个范围可能会影响效果。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果。
	// • 长宽比：长边：短边<5。 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过Base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// 注意：图片需要Base64编码，并且要去掉编码头部。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type CropImageRequest struct {
	*tchttp.BaseRequest
	
	// 需要裁剪区域的宽度，与Height共同组成所需裁剪的图片宽高比例。
	// 输入数字请大于0、小于图片宽度的像素值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 需要裁剪区域的高度，与Width共同组成所需裁剪的图片宽高比例。
	// 输入数字请大于0、小于图片高度的像素值。
	// 宽高比例（Width : Height）会简化为最简分数，即如果Width输入10、Height输入20，会简化为1：2。
	// Width : Height建议取值在[1, 2.5]之间，超过这个范围可能会影响效果。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果。
	// • 长宽比：长边：短边<5。 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过Base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// 注意：图片需要Base64编码，并且要去掉编码头部。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *CropImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CropImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CropImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CropImageResponseParams struct {
	// 裁剪区域左上角X坐标值
	X *int64 `json:"X,omitempty" name:"X"`

	// 裁剪区域左上角Y坐标值
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 裁剪区域的宽度，单位为像素
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 裁剪区域的高度，单位为像素
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 原图宽度，单位为像素
	OriginalWidth *int64 `json:"OriginalWidth,omitempty" name:"OriginalWidth"`

	// 原图高度，单位为像素
	OriginalHeight *int64 `json:"OriginalHeight,omitempty" name:"OriginalHeight"`

	// 0：抠图正常；
	// 1：原图过长，指原图的高度是宽度的1.8倍以上；
	// 2：原图过宽，指原图的宽度是高度的1.8倍以上；
	// 3：抠图区域过长，指抠图的高度是主体备选框高度的1.6倍以上；
	// 4：抠图区域过宽，指当没有检测到人脸时，抠图区域宽度是检测出的原图主体区域宽度的1.6倍以上；
	// 5：纯色图，指裁剪区域视觉较为单一、缺乏主体部分 ；
	// 6：宽高比异常，指Width : Height取值超出[1, 2.5]的范围；
	// 
	// 以上是辅助决策的参考建议，可以根据业务需求选择采纳或忽视。
	CropResult *int64 `json:"CropResult,omitempty" name:"CropResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CropImageResponse struct {
	*tchttp.BaseResponse
	Response *CropImageResponseParams `json:"Response"`
}

func (r *CropImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CropImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImagesRequestParams struct {
	// 图库名称。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 物品ID。
	EntityId *string `json:"EntityId,omitempty" name:"EntityId"`

	// 图片名称，如果不指定本参数，则删除EntityId下所有的图片；否则删除指定的图。
	PicName *string `json:"PicName,omitempty" name:"PicName"`
}

type DeleteImagesRequest struct {
	*tchttp.BaseRequest
	
	// 图库名称。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 物品ID。
	EntityId *string `json:"EntityId,omitempty" name:"EntityId"`

	// 图片名称，如果不指定本参数，则删除EntityId下所有的图片；否则删除指定的图。
	PicName *string `json:"PicName,omitempty" name:"PicName"`
}

func (r *DeleteImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "EntityId")
	delete(f, "PicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImagesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImagesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImagesResponseParams `json:"Response"`
}

func (r *DeleteImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupsRequestParams struct {
	// 起始序号，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 图库ID，如果不为空，则返回指定库信息。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DescribeGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 起始序号，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 图库ID，如果不为空，则返回指定库信息。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupsResponseParams struct {
	// 图库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*GroupInfo `json:"Groups,omitempty" name:"Groups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupsResponseParams `json:"Response"`
}

func (r *DescribeGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesRequestParams struct {
	// 图库名称。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 物品ID。
	EntityId *string `json:"EntityId,omitempty" name:"EntityId"`

	// 图片名称。
	PicName *string `json:"PicName,omitempty" name:"PicName"`
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
	
	// 图库名称。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 物品ID。
	EntityId *string `json:"EntityId,omitempty" name:"EntityId"`

	// 图片名称。
	PicName *string `json:"PicName,omitempty" name:"PicName"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "EntityId")
	delete(f, "PicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesResponseParams struct {
	// 图库名称。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 物品ID。
	EntityId *string `json:"EntityId,omitempty" name:"EntityId"`

	// 图片信息。
	ImageInfos []*ImageInfo `json:"ImageInfos,omitempty" name:"ImageInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagesResponseParams `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectChefDressRequestParams struct {
	// 图片的 Url 。
	// ImageUrl和ImageBase64必须提供一个，同时存在时优先使用ImageUrl字段。
	// 图片限制：
	// • 图片格式：支持PNG、JPG、JPEG、不支持 GIF 图片。
	// • 图片大小：对应图片 base64 编码后大小不可超过5M。图片分辨率不超过 3840 x 2160pixel。
	// 建议：
	// • 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。与ImageUrl同时存在时优先使用ImageUrl字段。
	// 注意：图片需要base64编码，并且要去掉编码头部。
	// 支持的图片格式：PNG、JPG、JPEG、暂不支持GIF格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过5M。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 人体检测模型开关，“true”为开启，“false”为关闭
	// 默认为开启，开启后可先对图片中的人体进行检测之后再进行属性识别
	EnableDetect *bool `json:"EnableDetect,omitempty" name:"EnableDetect"`

	// 人体优选开关，“true”为开启，“false”为关闭
	// 开启后自动对检测质量低的人体进行优选过滤，有助于提高属性识别的准确率。
	// 默认为开启，仅在人体检测开关开启时可配置，人体检测模型关闭时人体优选也关闭
	// 人体优选开启时，检测到的人体分辨率不超过1920*1080 pixel
	EnablePreferred *bool `json:"EnablePreferred,omitempty" name:"EnablePreferred"`
}

type DetectChefDressRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Url 。
	// ImageUrl和ImageBase64必须提供一个，同时存在时优先使用ImageUrl字段。
	// 图片限制：
	// • 图片格式：支持PNG、JPG、JPEG、不支持 GIF 图片。
	// • 图片大小：对应图片 base64 编码后大小不可超过5M。图片分辨率不超过 3840 x 2160pixel。
	// 建议：
	// • 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。与ImageUrl同时存在时优先使用ImageUrl字段。
	// 注意：图片需要base64编码，并且要去掉编码头部。
	// 支持的图片格式：PNG、JPG、JPEG、暂不支持GIF格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过5M。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 人体检测模型开关，“true”为开启，“false”为关闭
	// 默认为开启，开启后可先对图片中的人体进行检测之后再进行属性识别
	EnableDetect *bool `json:"EnableDetect,omitempty" name:"EnableDetect"`

	// 人体优选开关，“true”为开启，“false”为关闭
	// 开启后自动对检测质量低的人体进行优选过滤，有助于提高属性识别的准确率。
	// 默认为开启，仅在人体检测开关开启时可配置，人体检测模型关闭时人体优选也关闭
	// 人体优选开启时，检测到的人体分辨率不超过1920*1080 pixel
	EnablePreferred *bool `json:"EnablePreferred,omitempty" name:"EnablePreferred"`
}

func (r *DetectChefDressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectChefDressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "EnableDetect")
	delete(f, "EnablePreferred")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectChefDressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectChefDressResponseParams struct {
	// 识别到的人体属性信息。单个人体属性信息包括人体检测置信度，属性信息，人体检测框。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bodies []*AttributesForBody `json:"Bodies,omitempty" name:"Bodies"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectChefDressResponse struct {
	*tchttp.BaseResponse
	Response *DetectChefDressResponseParams `json:"Response"`
}

func (r *DetectChefDressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectChefDressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectDisgustRequestParams struct {
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type DetectDisgustRequest struct {
	*tchttp.BaseRequest
	
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *DetectDisgustRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectDisgustRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectDisgustRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectDisgustResponseParams struct {
	// 对于图片中包含恶心内容的置信度，取值[0,1]，一般超过0.5则表明可能是恶心图片。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 与图像内容最相似的恶心内容的类别，包含腐烂、密集、畸形、血腥、蛇、虫子、牙齿等。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectDisgustResponse struct {
	*tchttp.BaseResponse
	Response *DetectDisgustResponseParams `json:"Response"`
}

func (r *DetectDisgustResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectDisgustResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectEnvelopeRequestParams struct {
	// 图片的URL地址。图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 图片大小的限制为4M，图片像素的限制为4k。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。与ImageUrl同时存在时优先使用ImageUrl字段。 
	// 图片大小的限制为4M，图片像素的限制为4k。
	// **注意：图片需要base64编码，并且要去掉编码头部。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type DetectEnvelopeRequest struct {
	*tchttp.BaseRequest
	
	// 图片的URL地址。图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 图片大小的限制为4M，图片像素的限制为4k。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。与ImageUrl同时存在时优先使用ImageUrl字段。 
	// 图片大小的限制为4M，图片像素的限制为4k。
	// **注意：图片需要base64编码，并且要去掉编码头部。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *DetectEnvelopeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectEnvelopeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectEnvelopeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectEnvelopeResponseParams struct {
	// 一级标签结果数组。识别是否文件封。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstTags []*ImageTag `json:"FirstTags,omitempty" name:"FirstTags"`

	// 二级标签结果数组。识别文件封正反面。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondTags []*ImageTag `json:"SecondTags,omitempty" name:"SecondTags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectEnvelopeResponse struct {
	*tchttp.BaseResponse
	Response *DetectEnvelopeResponseParams `json:"Response"`
}

func (r *DetectEnvelopeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectEnvelopeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectLabelBetaRequestParams struct {
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 本次调用支持的识别场景，可选值如下：
	// WEB，针对网络图片优化;
	// CAMERA，针对手机摄像头拍摄图片优化;
	// ALBUM，针对手机相册、网盘产品优化;
	// NEWS，针对新闻、资讯、广电等行业优化；
	// NONECAM，非实拍图；
	// LOCATION，主体位置识别；
	// 如果不传此参数，则默认为WEB。
	// 
	// 支持多场景（Scenes）一起检测。例如，使用 Scenes=["WEB", "CAMERA"]，即对一张图片使用两个模型同时检测，输出两套识别结果。
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`
}

type DetectLabelBetaRequest struct {
	*tchttp.BaseRequest
	
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 本次调用支持的识别场景，可选值如下：
	// WEB，针对网络图片优化;
	// CAMERA，针对手机摄像头拍摄图片优化;
	// ALBUM，针对手机相册、网盘产品优化;
	// NEWS，针对新闻、资讯、广电等行业优化；
	// NONECAM，非实拍图；
	// LOCATION，主体位置识别；
	// 如果不传此参数，则默认为WEB。
	// 
	// 支持多场景（Scenes）一起检测。例如，使用 Scenes=["WEB", "CAMERA"]，即对一张图片使用两个模型同时检测，输出两套识别结果。
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`
}

func (r *DetectLabelBetaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLabelBetaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "Scenes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectLabelBetaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectLabelBetaResponseParams struct {
	// Web网络版标签结果数组。如未选择WEB场景，则为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*DetectLabelItem `json:"Labels,omitempty" name:"Labels"`

	// Camera摄像头版标签结果数组。如未选择CAMERA场景，则为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CameraLabels []*DetectLabelItem `json:"CameraLabels,omitempty" name:"CameraLabels"`

	// Album相册版标签结果数组。如未选择ALBUM场景，则为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbumLabels []*DetectLabelItem `json:"AlbumLabels,omitempty" name:"AlbumLabels"`

	// News新闻版标签结果数组。如未选择NEWS场景，则为空。
	// 新闻版目前为测试阶段，暂不提供每个标签的一级、二级分类信息的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewsLabels []*DetectLabelItem `json:"NewsLabels,omitempty" name:"NewsLabels"`

	// 非实拍标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoneCamLabels []*DetectLabelItem `json:"NoneCamLabels,omitempty" name:"NoneCamLabels"`

	// 识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocationLabels []*Product `json:"LocationLabels,omitempty" name:"LocationLabels"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectLabelBetaResponse struct {
	*tchttp.BaseResponse
	Response *DetectLabelBetaResponseParams `json:"Response"`
}

func (r *DetectLabelBetaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLabelBetaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectLabelItem struct {
	// 图片中的物体名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 算法对于Name的置信度，0-100之间，值越高，表示对于Name越确定。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 标签的一级分类
	FirstCategory *string `json:"FirstCategory,omitempty" name:"FirstCategory"`

	// 标签的二级分类
	SecondCategory *string `json:"SecondCategory,omitempty" name:"SecondCategory"`
}

// Predefined struct for user
type DetectLabelProRequestParams struct {
	// 图片 URL 地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG、BMP。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边:短边<5； 
	// • 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片 Base64 编码数据。
	// 与ImageUrl同时存在时优先使用ImageUrl字段。
	// 图片限制：
	// • 图片格式：PNG、JPG、JPEG、BMP。 
	// • 图片大小：经Base64编码后不超过4M。
	// **<font color=#1E90FF>注意：图片需要Base64编码，并且要去掉编码头部。</font>**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type DetectLabelProRequest struct {
	*tchttp.BaseRequest
	
	// 图片 URL 地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG、BMP。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边:短边<5； 
	// • 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片 Base64 编码数据。
	// 与ImageUrl同时存在时优先使用ImageUrl字段。
	// 图片限制：
	// • 图片格式：PNG、JPG、JPEG、BMP。 
	// • 图片大小：经Base64编码后不超过4M。
	// **<font color=#1E90FF>注意：图片需要Base64编码，并且要去掉编码头部。</font>**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *DetectLabelProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLabelProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectLabelProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectLabelProResponseParams struct {
	// 返回标签数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*DetectLabelItem `json:"Labels,omitempty" name:"Labels"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectLabelProResponse struct {
	*tchttp.BaseResponse
	Response *DetectLabelProResponseParams `json:"Response"`
}

func (r *DetectLabelProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLabelProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectLabelRequestParams struct {
	// 图片 Base64 编码数据。
	// 与ImageUrl同时存在时优先使用ImageUrl字段。
	// 图片限制：
	// • 图片格式：PNG、JPG、JPEG、BMP。 
	// • 图片大小：经Base64编码后不超过4M。
	// **<font color=#1E90FF>注意：图片需要Base64编码，并且要去掉编码头部。</font>**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片 URL 地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG、BMP。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边:短边<5； 
	// • 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 本次调用支持的识别场景，可选值如下：
	// • WEB，针对网络图片优化;
	// • CAMERA，针对手机摄像头拍摄图片优化;
	// • ALBUM，针对手机相册、网盘产品优化;
	// • NEWS，针对新闻、资讯、广电等行业优化；
	// 如果不传此参数，则默认为WEB。
	// 
	// 支持多场景（Scenes）一起检测。例如，使用 Scenes=["WEB", "CAMERA"]，即对一张图片使用两个模型同时检测，输出两套识别结果。
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`
}

type DetectLabelRequest struct {
	*tchttp.BaseRequest
	
	// 图片 Base64 编码数据。
	// 与ImageUrl同时存在时优先使用ImageUrl字段。
	// 图片限制：
	// • 图片格式：PNG、JPG、JPEG、BMP。 
	// • 图片大小：经Base64编码后不超过4M。
	// **<font color=#1E90FF>注意：图片需要Base64编码，并且要去掉编码头部。</font>**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片 URL 地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG、BMP。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边:短边<5； 
	// • 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 本次调用支持的识别场景，可选值如下：
	// • WEB，针对网络图片优化;
	// • CAMERA，针对手机摄像头拍摄图片优化;
	// • ALBUM，针对手机相册、网盘产品优化;
	// • NEWS，针对新闻、资讯、广电等行业优化；
	// 如果不传此参数，则默认为WEB。
	// 
	// 支持多场景（Scenes）一起检测。例如，使用 Scenes=["WEB", "CAMERA"]，即对一张图片使用两个模型同时检测，输出两套识别结果。
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`
}

func (r *DetectLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Scenes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectLabelResponseParams struct {
	// Web网络版标签结果数组。如未选择WEB场景，则为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*DetectLabelItem `json:"Labels,omitempty" name:"Labels"`

	// Camera摄像头版标签结果数组。如未选择CAMERA场景，则为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CameraLabels []*DetectLabelItem `json:"CameraLabels,omitempty" name:"CameraLabels"`

	// Album相册版标签结果数组。如未选择ALBUM场景，则为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbumLabels []*DetectLabelItem `json:"AlbumLabels,omitempty" name:"AlbumLabels"`

	// News新闻版标签结果数组。如未选择NEWS场景，则为空。
	// 新闻版目前为测试阶段，暂不提供每个标签的一级、二级分类信息的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewsLabels []*DetectLabelItem `json:"NewsLabels,omitempty" name:"NewsLabels"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectLabelResponse struct {
	*tchttp.BaseResponse
	Response *DetectLabelResponseParams `json:"Response"`
}

func (r *DetectLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectMisbehaviorRequestParams struct {
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type DetectMisbehaviorRequest struct {
	*tchttp.BaseRequest
	
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *DetectMisbehaviorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectMisbehaviorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectMisbehaviorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectMisbehaviorResponseParams struct {
	// 对于图片中包含不良行为的置信度，取值[0,1]，一般超过0.5则表明可能包含不良行为内容；
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 图像中最可能包含的不良行为类别，包括赌博、打架斗殴、吸毒等。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectMisbehaviorResponse struct {
	*tchttp.BaseResponse
	Response *DetectMisbehaviorResponseParams `json:"Response"`
}

func (r *DetectMisbehaviorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectMisbehaviorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectPetRequestParams struct {
	// 图片的URL地址。图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 图片大小的限制为4M，图片像素的限制为4k。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。与ImageUrl同时存在时优先使用ImageUrl字段。 
	// 图片大小的限制为4M，图片像素的限制为4k。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type DetectPetRequest struct {
	*tchttp.BaseRequest
	
	// 图片的URL地址。图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 图片大小的限制为4M，图片像素的限制为4k。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。与ImageUrl同时存在时优先使用ImageUrl字段。 
	// 图片大小的限制为4M，图片像素的限制为4k。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *DetectPetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectPetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectPetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectPetResponseParams struct {
	// 识别出图片中的宠物信息列表。
	Pets []*Pet `json:"Pets,omitempty" name:"Pets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectPetResponse struct {
	*tchttp.BaseResponse
	Response *DetectPetResponseParams `json:"Response"`
}

func (r *DetectPetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectPetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectProductBetaRequestParams struct {
	// 图片限制：内测版仅支持jpg、jpeg，图片大小不超过1M，分辨率在25万到100万之间。 
	// 建议先对图片进行压缩，以便提升处理速度。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过1M，分辨率在25万到100万之间。 
	// 与ImageUrl同时存在时优先使用ImageUrl字段。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 是否需要百科信息 1：是，0: 否，默认是0
	NeedLemma *int64 `json:"NeedLemma,omitempty" name:"NeedLemma"`
}

type DetectProductBetaRequest struct {
	*tchttp.BaseRequest
	
	// 图片限制：内测版仅支持jpg、jpeg，图片大小不超过1M，分辨率在25万到100万之间。 
	// 建议先对图片进行压缩，以便提升处理速度。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过1M，分辨率在25万到100万之间。 
	// 与ImageUrl同时存在时优先使用ImageUrl字段。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 是否需要百科信息 1：是，0: 否，默认是0
	NeedLemma *int64 `json:"NeedLemma,omitempty" name:"NeedLemma"`
}

func (r *DetectProductBetaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectProductBetaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "NeedLemma")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectProductBetaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectProductBetaResponseParams struct {
	// 检测到的图片中的商品位置和品类预测。 
	// 当图片中存在多个商品时，输出多组坐标，按照__显著性__排序（综合考虑面积、是否在中心、检测算法置信度）。 
	// 最多可以输出__3组__检测结果。
	RegionDetected []*RegionDetected `json:"RegionDetected,omitempty" name:"RegionDetected"`

	// 图像识别出的商品的详细信息。 
	// 当图像中检测到多个物品时，会对显著性最高的进行识别。
	ProductInfo *ProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo"`

	// 相似商品信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductInfoList []*ProductInfo `json:"ProductInfoList,omitempty" name:"ProductInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectProductBetaResponse struct {
	*tchttp.BaseResponse
	Response *DetectProductBetaResponseParams `json:"Response"`
}

func (r *DetectProductBetaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectProductBetaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectProductRequestParams struct {
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type DetectProductRequest struct {
	*tchttp.BaseRequest
	
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *DetectProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectProductResponseParams struct {
	// 商品识别结果数组
	Products []*Product `json:"Products,omitempty" name:"Products"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectProductResponse struct {
	*tchttp.BaseResponse
	Response *DetectProductResponseParams `json:"Response"`
}

func (r *DetectProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectSecurityRequestParams struct {
	// 图片的 Url 。
	// ImageUrl和ImageBase64必须提供一个，同时存在时优先使用ImageUrl字段。
	// 图片限制：
	// • 图片格式：支持PNG、JPG、JPEG、不支持 GIF 图片。
	// • 图片大小：对应图片 base64 编码后大小不可超过5M。图片分辨率不超过3840 x 2160 pixel。
	// 建议：
	// • 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。
	// 最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// 注意：图片需要base64编码，并且要去掉编码头部。
	// 支持的图片格式：PNG、JPG、JPEG、暂不支持GIF格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过5M。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 人体检测模型开关，“true”为开启，“false”为关闭
	// 开启后可先对图片中的人体进行检测之后再进行属性识别，默认为开启
	EnableDetect *bool `json:"EnableDetect,omitempty" name:"EnableDetect"`

	// 人体优选开关，“true”为开启，“false”为关闭
	// 开启后自动对检测质量低的人体进行优选过滤，有助于提高属性识别的准确率。
	// 默认为开启，仅在人体检测开关开启时可配置，人体检测模型关闭时人体优选也关闭
	// 如开启人体优选，检测到的人体分辨率需不大于1920*1080 pixel
	EnablePreferred *bool `json:"EnablePreferred,omitempty" name:"EnablePreferred"`
}

type DetectSecurityRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Url 。
	// ImageUrl和ImageBase64必须提供一个，同时存在时优先使用ImageUrl字段。
	// 图片限制：
	// • 图片格式：支持PNG、JPG、JPEG、不支持 GIF 图片。
	// • 图片大小：对应图片 base64 编码后大小不可超过5M。图片分辨率不超过3840 x 2160 pixel。
	// 建议：
	// • 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。
	// 最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// 注意：图片需要base64编码，并且要去掉编码头部。
	// 支持的图片格式：PNG、JPG、JPEG、暂不支持GIF格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过5M。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 人体检测模型开关，“true”为开启，“false”为关闭
	// 开启后可先对图片中的人体进行检测之后再进行属性识别，默认为开启
	EnableDetect *bool `json:"EnableDetect,omitempty" name:"EnableDetect"`

	// 人体优选开关，“true”为开启，“false”为关闭
	// 开启后自动对检测质量低的人体进行优选过滤，有助于提高属性识别的准确率。
	// 默认为开启，仅在人体检测开关开启时可配置，人体检测模型关闭时人体优选也关闭
	// 如开启人体优选，检测到的人体分辨率需不大于1920*1080 pixel
	EnablePreferred *bool `json:"EnablePreferred,omitempty" name:"EnablePreferred"`
}

func (r *DetectSecurityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectSecurityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "EnableDetect")
	delete(f, "EnablePreferred")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectSecurityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectSecurityResponseParams struct {
	// 识别到的人体属性信息。单个人体属性信息包括人体检测置信度，属性信息，人体检测框。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bodies []*AttributesForBody `json:"Bodies,omitempty" name:"Bodies"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectSecurityResponse struct {
	*tchttp.BaseResponse
	Response *DetectSecurityResponseParams `json:"Response"`
}

func (r *DetectSecurityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectSecurityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnhanceImageRequestParams struct {
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，最大不超过250万像素，否则影响识别效果。 
	// • 长宽比：长边：短边<5。 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。图片经过Base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// 注意：图片需要Base64编码，并且要去掉编码头部。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type EnhanceImageRequest struct {
	*tchttp.BaseRequest
	
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，最大不超过250万像素，否则影响识别效果。 
	// • 长宽比：长边：短边<5。 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。图片经过Base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// 注意：图片需要Base64编码，并且要去掉编码头部。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *EnhanceImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnhanceImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnhanceImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnhanceImageResponseParams struct {
	// 增强后图片的base64编码。
	EnhancedImage *string `json:"EnhancedImage,omitempty" name:"EnhancedImage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnhanceImageResponse struct {
	*tchttp.BaseResponse
	Response *EnhanceImageResponseParams `json:"Response"`
}

func (r *EnhanceImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnhanceImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupInfo struct {
	// 图库Id。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 图库名称。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 图库简介。
	Brief *string `json:"Brief,omitempty" name:"Brief"`

	// 图库容量。
	MaxCapacity *uint64 `json:"MaxCapacity,omitempty" name:"MaxCapacity"`

	// 该库的访问限频 。
	MaxQps *uint64 `json:"MaxQps,omitempty" name:"MaxQps"`

	// 图库类型，对应不同服务类型，默认为1。建议手动调整为4～6，1～3为历史版本，不推荐。
	// 参数值：
	// 4：在自建图库中搜索相同原图，可支持裁剪、翻转、调色、加水印后的图片搜索，适用于图片版权保护、原图查询等场景。
	// 5：在自建图库中搜索相同或相似的商品图片，适用于商品分类、检索、推荐等电商场景。
	// 6：在自建图片库中搜索与输入图片高度相似的图片，适用于相似图案、logo、纹理等图像元素的搜索。
	GroupType *uint64 `json:"GroupType,omitempty" name:"GroupType"`

	// 图库图片数量。
	PicCount *uint64 `json:"PicCount,omitempty" name:"PicCount"`

	// 图库创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 图库更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ImageInfo struct {
	// 图片名称。
	EntityId *string `json:"EntityId,omitempty" name:"EntityId"`

	// 用户自定义的内容。
	CustomContent *string `json:"CustomContent,omitempty" name:"CustomContent"`

	// 图片自定义标签，JSON格式。
	Tags *string `json:"Tags,omitempty" name:"Tags"`

	// 图片名称。
	PicName *string `json:"PicName,omitempty" name:"PicName"`

	// 相似度。
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type ImageRect struct {
	// 左上角横坐标。
	X *int64 `json:"X,omitempty" name:"X"`

	// 左上角纵坐标。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type ImageTag struct {
	// 标签内容。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 置信度范围在0-100之间。值越高，表示目标为相应结果的可能性越高。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type LemmaInfo struct {
	// 词条
	// 注意：此字段可能返回 null，表示取不到有效值。
	LemmaTitle *string `json:"LemmaTitle,omitempty" name:"LemmaTitle"`

	// 词条描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	LemmaAbstract *string `json:"LemmaAbstract,omitempty" name:"LemmaAbstract"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

type Location struct {
	// 位置矩形框的左上角横坐标
	XMin *int64 `json:"XMin,omitempty" name:"XMin"`

	// 位置矩形框的左上角纵坐标
	YMin *int64 `json:"YMin,omitempty" name:"YMin"`

	// 位置矩形框的右下角横坐标
	XMax *int64 `json:"XMax,omitempty" name:"XMax"`

	// 位置矩形框的右下角纵坐标
	YMax *int64 `json:"YMax,omitempty" name:"YMax"`
}

type ObjectInfo struct {
	// 图像主体区域。
	Box *Box `json:"Box,omitempty" name:"Box"`

	// 主体类别ID。
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 整张图颜色信息。
	Colors []*ColorInfo `json:"Colors,omitempty" name:"Colors"`

	// 属性信息。
	Attributes []*Attribute `json:"Attributes,omitempty" name:"Attributes"`

	// 图像的所有主体区域，置信度，以及主体区域类别ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllBox []*Box `json:"AllBox,omitempty" name:"AllBox"`
}

type Pet struct {
	// 识别出的宠物类型（猫或者狗，暂不支持识别猫狗品种）。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别服务给识别目标打出的置信度，范围在0-100之间。值越高，表示目标为相应结果的可能性越高。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 识别目标在图片中的坐标。
	Location *Rect `json:"Location,omitempty" name:"Location"`
}

type Product struct {
	// 图片中商品的三级分类识别结果，选取所有三级分类中的置信度最大者
	Name *string `json:"Name,omitempty" name:"Name"`

	// 三级商品分类对应的一级分类和二级分类，两级之间用“-”（中划线）隔开，例如商品名称是“硬盘”，那么Parents输出为“电脑、办公-电脑配件”
	Parents *string `json:"Parents,omitempty" name:"Parents"`

	// 算法对于Name的置信度，0-100之间，值越高，表示对于Name越确定
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 商品坐标X轴的最小值
	XMin *int64 `json:"XMin,omitempty" name:"XMin"`

	// 商品坐标Y轴的最小值
	YMin *int64 `json:"YMin,omitempty" name:"YMin"`

	// 商品坐标X轴的最大值
	XMax *int64 `json:"XMax,omitempty" name:"XMax"`

	// 商品坐标Y轴的最大值
	YMax *int64 `json:"YMax,omitempty" name:"YMax"`
}

type ProductInfo struct {
	// 1表示找到同款商品，以下字段为同款商品信息； 
	// 0表示未找到同款商品， 具体商品信息为空（参考价格、名称、品牌等），仅提供商品类目和参考图片（商品库中找到的最相似图片，供参考）。  
	// 是否找到同款的判断依据为Score分值，分值越大则同款的可能性越大。
	FindSKU *int64 `json:"FindSKU,omitempty" name:"FindSKU"`

	// 本商品在图片中的坐标，表示为矩形框的四个顶点坐标。
	Location *Location `json:"Location,omitempty" name:"Location"`

	// 商品名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 商品品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 参考价格，综合多个信息源，仅供参考。
	Price *string `json:"Price,omitempty" name:"Price"`

	// 识别结果的商品类目。 
	// 包含：鞋、图书音像、箱包、美妆个护、服饰、家电数码、玩具乐器、食品饮料、珠宝、家居家装、药品、酒水、绿植园艺、其他商品、非商品等。 
	// 当类别为“非商品”时，除Location、Score和本字段之外的商品信息为空。
	ProductCategory *string `json:"ProductCategory,omitempty" name:"ProductCategory"`

	// 输入图片中的主体物品和输出结果的相似度。分值越大，输出结果与输入图片是同款的可能性越高。
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 搜索到的商品配图URL。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 百科词条列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LemmaInfoList []*LemmaInfo `json:"LemmaInfoList,omitempty" name:"LemmaInfoList"`
}

// Predefined struct for user
type RecognizeCarProRequestParams struct {
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	// 支持的图片格式：PNG、JPG、JPEG、BMP，暂不支持GIF格式。支持的图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type RecognizeCarProRequest struct {
	*tchttp.BaseRequest
	
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	// 支持的图片格式：PNG、JPG、JPEG、BMP，暂不支持GIF格式。支持的图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *RecognizeCarProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeCarProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeCarProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeCarProResponseParams struct {
	// 汽车的四个矩形顶点坐标，如果图片中存在多辆车，则输出最大车辆的坐标。
	CarCoords []*Coord `json:"CarCoords,omitempty" name:"CarCoords"`

	// 车辆属性识别的结果数组，如果识别到多辆车，则会输出每辆车的top1结果。
	// 注意：置信度是指车牌信息置信度。
	CarTags []*CarTagItem `json:"CarTags,omitempty" name:"CarTags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RecognizeCarProResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeCarProResponseParams `json:"Response"`
}

func (r *RecognizeCarProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeCarProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeCarRequestParams struct {
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	// 支持的图片格式：PNG、JPG、JPEG、BMP，暂不支持GIF格式。支持的图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type RecognizeCarRequest struct {
	*tchttp.BaseRequest
	
	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	// 支持的图片格式：PNG、JPG、JPEG、BMP，暂不支持GIF格式。支持的图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *RecognizeCarRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeCarRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeCarRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeCarResponseParams struct {
	// 汽车的四个矩形顶点坐标，如果图片中存在多辆车，则输出最大车辆的坐标。
	CarCoords []*Coord `json:"CarCoords,omitempty" name:"CarCoords"`

	// 车辆属性识别的结果数组，如果识别到多辆车，则会输出每辆车的top1结果。
	CarTags []*CarTagItem `json:"CarTags,omitempty" name:"CarTags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RecognizeCarResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeCarResponseParams `json:"Response"`
}

func (r *RecognizeCarResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeCarResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Rect struct {
	// x轴坐标
	X *int64 `json:"X,omitempty" name:"X"`

	// y轴坐标
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// (x,y)坐标距离长度
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// (x,y)坐标距离高度
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type RegionDetected struct {
	// 商品的品类预测结果。 
	// 包含：鞋、图书音像、箱包、美妆个护、服饰、家电数码、玩具乐器、食品饮料、珠宝、家居家装、药品、酒水、绿植园艺、其他商品、非商品等。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 商品品类预测的置信度
	CategoryScore *float64 `json:"CategoryScore,omitempty" name:"CategoryScore"`

	// 检测到的主体在图片中的坐标，表示为矩形框的四个顶点坐标
	Location *Location `json:"Location,omitempty" name:"Location"`
}

// Predefined struct for user
type SearchImageRequestParams struct {
	// 图库名称。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 图片的 Url 。
	// ImageUrl和ImageBase64必须提供一个，如果都提供，只使用ImageUrl。
	// 图片限制：
	// • 图片格式：支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// • 图片大小：对应图片 base64 编码后大小不可超过5M。图片分辨率不超过4096\*4096。
	// • 如果在商品图像搜索中开启主体识别，分辨率不超过2000\*2000，图片长宽比小于10。
	// 建议：
	// • 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的Url速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片 base64 数据。
	// ImageUrl和ImageBase64必须提供一个，如果都提供，只使用ImageUrl。
	// 图片限制：
	// • 图片格式：支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// • 图片大小：base64 编码后大小不可超过5M。图片分辨率不超过4096\*4096。
	// • 如果在商品图像搜索中开启主体识别，分辨率不超过2000\*2000，图片长宽比小于10。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 返回结果的数量，默认值为10，最大值为100。
	// 按照相似度分数由高到低排序。
	// **<font color=#1E90FF>服务类型为图案花纹搜索时Limit = 1，最多只能返回1个结果。</font>**
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果的起始序号，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 匹配阈值。
	// 只有图片相似度分数超过匹配阈值的结果才会返回。
	// 当MatchThreshold为0（默认值）时，各服务类型将按照以下默认的匹配阈值进行结果过滤：
	// • 通用图像搜索1.0版：50。
	// • 商品图像搜索2.0升级版：45。
	// • 商品图像搜索1.0版：28。
	// • 图案花纹搜索1.0版：56。
	// 建议：
	// 可以手动调整MatchThreshold值来控制输出结果的范围。如果发现无检索结果，可能是因为图片相似度较低导致检索结果被匹配阈值过滤，建议调整为较低的阈值后再次尝试检索。
	MatchThreshold *int64 `json:"MatchThreshold,omitempty" name:"MatchThreshold"`

	// 标签过滤条件。
	// 针对创建图片时提交的Tags信息进行条件过滤。支持>、>=、 <、 <=、=，!=，多个条件之间支持AND和OR进行连接。
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 图像主体区域。
	// 若设置主体区域，提取指定的区域进行检索。
	ImageRect *ImageRect `json:"ImageRect,omitempty" name:"ImageRect"`

	// 是否需要启用主体识别，默认为**TRUE** 。
	// • 为**TRUE**时，启用主体识别，返回主体信息。若没有指定**ImageRect**，自动提取最大面积主体进行检索并进行主体识别。主体识别结果可在**Response中**获取。
	// • 为**FALSE**时，不启用主体识别，不返回主体信息。若没有指定**ImageRect**，以整张图检索图片。
	// **<font color=#1E90FF>注意：仅服务类型为商品图像搜索时才生效。</font>**
	EnableDetect *bool `json:"EnableDetect,omitempty" name:"EnableDetect"`

	// 图像类目ID。
	// 若设置类目ID，提取以下类目的主体进行检索。
	// 类目取值说明：
	// 0：上衣。
	// 1：裙装。
	// 2：下装。
	// 3：包。
	// 4：鞋。
	// 5：配饰。
	// **<font color=#1E90FF>注意：仅服务类型为商品图像搜索时才生效。</font>**
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
}

type SearchImageRequest struct {
	*tchttp.BaseRequest
	
	// 图库名称。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 图片的 Url 。
	// ImageUrl和ImageBase64必须提供一个，如果都提供，只使用ImageUrl。
	// 图片限制：
	// • 图片格式：支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// • 图片大小：对应图片 base64 编码后大小不可超过5M。图片分辨率不超过4096\*4096。
	// • 如果在商品图像搜索中开启主体识别，分辨率不超过2000\*2000，图片长宽比小于10。
	// 建议：
	// • 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的Url速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片 base64 数据。
	// ImageUrl和ImageBase64必须提供一个，如果都提供，只使用ImageUrl。
	// 图片限制：
	// • 图片格式：支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// • 图片大小：base64 编码后大小不可超过5M。图片分辨率不超过4096\*4096。
	// • 如果在商品图像搜索中开启主体识别，分辨率不超过2000\*2000，图片长宽比小于10。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 返回结果的数量，默认值为10，最大值为100。
	// 按照相似度分数由高到低排序。
	// **<font color=#1E90FF>服务类型为图案花纹搜索时Limit = 1，最多只能返回1个结果。</font>**
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果的起始序号，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 匹配阈值。
	// 只有图片相似度分数超过匹配阈值的结果才会返回。
	// 当MatchThreshold为0（默认值）时，各服务类型将按照以下默认的匹配阈值进行结果过滤：
	// • 通用图像搜索1.0版：50。
	// • 商品图像搜索2.0升级版：45。
	// • 商品图像搜索1.0版：28。
	// • 图案花纹搜索1.0版：56。
	// 建议：
	// 可以手动调整MatchThreshold值来控制输出结果的范围。如果发现无检索结果，可能是因为图片相似度较低导致检索结果被匹配阈值过滤，建议调整为较低的阈值后再次尝试检索。
	MatchThreshold *int64 `json:"MatchThreshold,omitempty" name:"MatchThreshold"`

	// 标签过滤条件。
	// 针对创建图片时提交的Tags信息进行条件过滤。支持>、>=、 <、 <=、=，!=，多个条件之间支持AND和OR进行连接。
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 图像主体区域。
	// 若设置主体区域，提取指定的区域进行检索。
	ImageRect *ImageRect `json:"ImageRect,omitempty" name:"ImageRect"`

	// 是否需要启用主体识别，默认为**TRUE** 。
	// • 为**TRUE**时，启用主体识别，返回主体信息。若没有指定**ImageRect**，自动提取最大面积主体进行检索并进行主体识别。主体识别结果可在**Response中**获取。
	// • 为**FALSE**时，不启用主体识别，不返回主体信息。若没有指定**ImageRect**，以整张图检索图片。
	// **<font color=#1E90FF>注意：仅服务类型为商品图像搜索时才生效。</font>**
	EnableDetect *bool `json:"EnableDetect,omitempty" name:"EnableDetect"`

	// 图像类目ID。
	// 若设置类目ID，提取以下类目的主体进行检索。
	// 类目取值说明：
	// 0：上衣。
	// 1：裙装。
	// 2：下装。
	// 3：包。
	// 4：鞋。
	// 5：配饰。
	// **<font color=#1E90FF>注意：仅服务类型为商品图像搜索时才生效。</font>**
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
}

func (r *SearchImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "MatchThreshold")
	delete(f, "Filter")
	delete(f, "ImageRect")
	delete(f, "EnableDetect")
	delete(f, "CategoryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchImageResponseParams struct {
	// 返回结果数量。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 图片信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageInfos []*ImageInfo `json:"ImageInfos,omitempty" name:"ImageInfos"`

	// 输入图的主体信息。
	// 若启用主体识别且在请求中指定了类目ID或主体区域，以指定的主体为准。若启用主体识别且没有指定，以最大面积主体为准。
	// **<font color=#1E90FF>注意：仅服务类型为商品图像搜索时才生效。</font>**
	// 注意：此字段可能返回 null，表示取不到有效值。
	Object *ObjectInfo `json:"Object,omitempty" name:"Object"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchImageResponse struct {
	*tchttp.BaseResponse
	Response *SearchImageResponseParams `json:"Response"`
}

func (r *SearchImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateImageRequestParams struct {
	// 图库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 物品ID，最多支持64个字符。
	EntityId *string `json:"EntityId,omitempty" name:"EntityId"`

	// 图片名称，最多支持64个字符。
	PicName *string `json:"PicName,omitempty" name:"PicName"`

	// 新的自定义标签，最多不超过10个，格式为JSON。
	Tags *string `json:"Tags,omitempty" name:"Tags"`
}

type UpdateImageRequest struct {
	*tchttp.BaseRequest
	
	// 图库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 物品ID，最多支持64个字符。
	EntityId *string `json:"EntityId,omitempty" name:"EntityId"`

	// 图片名称，最多支持64个字符。
	PicName *string `json:"PicName,omitempty" name:"PicName"`

	// 新的自定义标签，最多不超过10个，格式为JSON。
	Tags *string `json:"Tags,omitempty" name:"Tags"`
}

func (r *UpdateImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "EntityId")
	delete(f, "PicName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateImageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateImageResponse struct {
	*tchttp.BaseResponse
	Response *UpdateImageResponseParams `json:"Response"`
}

func (r *UpdateImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}