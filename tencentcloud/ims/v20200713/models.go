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

package v20200713

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Device struct {
	// 发表消息设备IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Mac地址
	Mac *string `json:"Mac,omitnil,omitempty" name:"Mac"`

	// 设备指纹Token
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// 设备指纹ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备序列号
	IMEI *string `json:"IMEI,omitnil,omitempty" name:"IMEI"`

	// IOS设备，Identifier For Advertising（广告标识符）
	IDFA *string `json:"IDFA,omitnil,omitempty" name:"IDFA"`

	// IOS设备，IDFV - Identifier For Vendor（应用开发商标识符）
	IDFV *string `json:"IDFV,omitnil,omitempty" name:"IDFV"`

	// IP地址类型 0 代表ipv4 1 代表ipv6
	IpType *uint64 `json:"IpType,omitnil,omitempty" name:"IpType"`
}

// Predefined struct for user
type ImageModerationRequestParams struct {
	// 该字段用于标识业务场景。您可以在内容安全控制台创建对应的ID，配置不同的内容审核策略，通过接口调用，默认不填为0，后端使用默认策略。 -- 该字段暂未开放。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 数据ID，可以由英文字母、数字、下划线、-、@#组成，不超过64个字符
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 数据Base64编码，图片检测接口为图片文件内容，大小不能超过5M
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// 图片资源访问链接，__与FileContent参数必须二选一输入__ 。由于网络安全策略，送审带重定向的链接，可能引起下载失败，请尽量避免，比如Http返回302状态码的链接，可能导致接口返回ResourceUnavailable.ImageDownloadError
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 截帧频率，GIF图/长图检测专用，默认值为0，表示只会检测GIF图/长图的第一帧
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// GIF图/长图检测专用，代表均匀最大截帧数量，默认值为1（即只取GIF第一张，或长图不做切分处理（可能会造成处理超时））。
	MaxFrames *int64 `json:"MaxFrames,omitnil,omitempty" name:"MaxFrames"`

	// 账号相关信息字段，填入后可识别违规风险账号。
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// 设备相关信息字段，填入后可识别违规风险设备。
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`
}

type ImageModerationRequest struct {
	*tchttp.BaseRequest
	
	// 该字段用于标识业务场景。您可以在内容安全控制台创建对应的ID，配置不同的内容审核策略，通过接口调用，默认不填为0，后端使用默认策略。 -- 该字段暂未开放。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 数据ID，可以由英文字母、数字、下划线、-、@#组成，不超过64个字符
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 数据Base64编码，图片检测接口为图片文件内容，大小不能超过5M
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// 图片资源访问链接，__与FileContent参数必须二选一输入__ 。由于网络安全策略，送审带重定向的链接，可能引起下载失败，请尽量避免，比如Http返回302状态码的链接，可能导致接口返回ResourceUnavailable.ImageDownloadError
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 截帧频率，GIF图/长图检测专用，默认值为0，表示只会检测GIF图/长图的第一帧
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// GIF图/长图检测专用，代表均匀最大截帧数量，默认值为1（即只取GIF第一张，或长图不做切分处理（可能会造成处理超时））。
	MaxFrames *int64 `json:"MaxFrames,omitnil,omitempty" name:"MaxFrames"`

	// 账号相关信息字段，填入后可识别违规风险账号。
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// 设备相关信息字段，填入后可识别违规风险设备。
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`
}

func (r *ImageModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageModerationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizType")
	delete(f, "DataId")
	delete(f, "FileContent")
	delete(f, "FileUrl")
	delete(f, "Interval")
	delete(f, "MaxFrames")
	delete(f, "User")
	delete(f, "Device")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageModerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageModerationResponseParams struct {
	// 数据是否属于恶意类型。
	// 0：正常，1：可疑；
	HitFlag *int64 `json:"HitFlag,omitnil,omitempty" name:"HitFlag"`

	// 建议您拿到判断结果后的执行操作。
	// 建议值，Block：建议屏蔽，Review：建议复审，Pass：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 恶意标签，Normal：正常，Porn：色情，Abuse：谩骂，Ad：广告，Custom：自定义图片。
	// 以及令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 子标签名称，如色情--性行为；当未命中子标签时，返回空字符串；
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 机器判断当前分类的置信度，取值范围：0.00~100.00。分数越高，表示越有可能属于当前分类。
	// （如：色情 99.99，则该样本属于色情的置信度非常高。）
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 智能模型的识别结果，包括涉黄、广告等令人反感、不安全或不适宜的内容类型识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelResults []*LabelResult `json:"LabelResults,omitnil,omitempty" name:"LabelResults"`

	// 物体检测模型的审核结果，包括实体、广告台标/二维码等物体坐标信息与内容审核信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectResults []*ObjectResult `json:"ObjectResults,omitnil,omitempty" name:"ObjectResults"`

	// OCR识别后的文本识别结果，包括文本所处图片的OCR坐标信息以及图片文本的识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrResults []*OcrResult `json:"OcrResults,omitnil,omitempty" name:"OcrResults"`

	// 基于图片风险库识别的结果。
	// 风险库包括不安全黑库与正常白库的结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibResults []*LibResult `json:"LibResults,omitnil,omitempty" name:"LibResults"`

	// 请求参数中的DataId。
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 您在入参时所填入的Biztype参数。 -- 该字段暂未开放。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 扩展字段，用于特定信息返回，不同客户/Biztype下返回信息不同。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 该字段用于返回仅识别图片元素的模型结果；包括：场景模型命中的标签、置信度和位置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecognitionResults []*RecognitionResult `json:"RecognitionResults,omitnil,omitempty" name:"RecognitionResults"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageModerationResponse struct {
	*tchttp.BaseResponse
	Response *ImageModerationResponseParams `json:"Response"`
}

func (r *ImageModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageModerationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelDetailItem struct {
	// 序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 子标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 子标签分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`
}

type LabelResult struct {
	// 场景识别结果
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 建议您拿到判断结果后的执行操作。
	// 建议值，Block：建议屏蔽，Review：建议复审，Pass：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 恶意标签，Normal：正常，Porn：色情，Abuse：谩骂，Ad：广告，Custom：自定义图片。
	// 以及令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 子标签检测结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该标签模型命中的分值
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 分类模型命中子标签结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*LabelDetailItem `json:"Details,omitnil,omitempty" name:"Details"`
}

type LibDetail struct {
	// 序号
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 仅当Label为Custom自定义关键词时有效，表示自定义库id
	LibId *string `json:"LibId,omitnil,omitempty" name:"LibId"`

	// 仅当Label为Custom自定义关键词时有效，表示自定义库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// 图片ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 恶意标签，Normal：正常，Porn：色情，Abuse：谩骂，Ad：广告，Custom：自定义词库。
	// 以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 自定义标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 命中的模型分值
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`
}

type LibResult struct {
	// 场景识别结果
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 建议您拿到判断结果后的执行操作。
	// 建议值，Block：建议屏蔽，Review：建议复审，Pass：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 恶意标签，Normal：正常，Porn：色情，Abuse：谩骂，Ad：广告，Custom：自定义词库。
	// 以及令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 子标签检测结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该标签模型命中的分值
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 黑白库结果明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*LibDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type Location struct {
	// 左上角横坐标
	X *float64 `json:"X,omitnil,omitempty" name:"X"`

	// 左上角纵坐标
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 宽度
	Width *float64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 高度
	Height *float64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 检测框的旋转角度
	Rotate *float64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`
}

type ObjectDetail struct {
	// 序号
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 标签名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签值，
	// 当标签为二维码时，表示URL地址，如Name为QrCode时，Value为"http//abc.com/aaa"
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 分数
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 检测框坐标
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`

	// 二级标签名称
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 图库或人脸库id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 图或人脸id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`
}

type ObjectResult struct {
	// 场景识别结果
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 建议您拿到判断结果后的执行操作。
	// 建议值，Block：建议屏蔽，Review：建议复审，Pass：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 恶意标签，Normal：正常，Porn：色情，Abuse：谩骂，Ad：广告，Custom：自定义图片。
	// 以及令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 子标签检测结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该标签模型命中的分值
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 实体名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 实体检测结果明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*ObjectDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type OcrResult struct {
	// 场景识别结果
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 建议您拿到判断结果后的执行操作。
	// 建议值，Block：建议屏蔽，Review：建议复审，Pass：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 恶意标签，Normal：正常，Porn：色情，Abuse：谩骂，Ad：广告，Custom：自定义词库。
	// 以及令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 子标签检测结果
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该标签模型命中的分值
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// ocr结果详情
	Details []*OcrTextDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// ocr识别出的文本结果
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 是否命中结果，0 未命中 1命中
	HitFlag *uint64 `json:"HitFlag,omitnil,omitempty" name:"HitFlag"`
}

type OcrTextDetail struct {
	// OCR文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 恶意标签，Normal：正常，Porn：色情，Abuse：谩骂，Ad：广告，Custom：自定义词库。
	// 以及令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 仅当Label为Custom自定义关键词时有效，表示自定义库id
	LibId *string `json:"LibId,omitnil,omitempty" name:"LibId"`

	// 仅当Label为Custom自定义关键词时有效，表示自定义库名称
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// 该标签下命中的关键词
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 该标签模型命中的分值
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// OCR位置
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`

	// OCR文本识别置信度
	Rate *uint64 `json:"Rate,omitnil,omitempty" name:"Rate"`

	// OCR文本命中的二级标签
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`
}

type RecognitionResult struct {
	// 当前可能的取值：Scene（图片场景模型）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// Label对应模型下的识别标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*RecognitionTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type RecognitionTag struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 置信分：0～100，数值越大表示置信度越高
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 标签位置信息，若模型无位置信息，则可能为零值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`
}

type User struct {
	// 业务用户ID 如填写，会根据账号历史恶意情况，判定消息有害结果，特别是有利于可疑恶意情况下的辅助判断。账号可以填写微信uin、QQ号、微信openid、QQopenid、字符串等。该字段和账号类别确定唯一账号。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 业务用户ID类型 "1-微信uin 2-QQ号 3-微信群uin 4-qq群号 5-微信openid 6-QQopenid 7-其它string"
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 性别 默认0 未知 1 男性 2 女性
	Gender *uint64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// 年龄 默认0 未知
	Age *uint64 `json:"Age,omitnil,omitempty" name:"Age"`

	// 用户等级，默认0 未知 1 低 2 中 3 高
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 用户简介，长度不超过5000字
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 用户头像图片链接
	HeadUrl *string `json:"HeadUrl,omitnil,omitempty" name:"HeadUrl"`
}