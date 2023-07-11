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

package v20190321

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CodeDetail struct {
	// 二维码文本的编码格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrCharset *string `json:"StrCharset,omitempty" name:"StrCharset"`

	// 二维码在图片中的位置，由边界点的坐标表示
	// 注意：此字段可能返回 null，表示取不到有效值。
	QrCodePosition []*CodePosition `json:"QrCodePosition,omitempty" name:"QrCodePosition"`

	// 二维码的文本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrQrCodeText *string `json:"StrQrCodeText,omitempty" name:"StrQrCodeText"`

	// 二维码的类型：1:ONED_BARCODE，2:QRCOD，3:WXCODE，4:PDF417，5:DATAMATRIX
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uint32QrCodeType *int64 `json:"Uint32QrCodeType,omitempty" name:"Uint32QrCodeType"`

	// 二维码文本的编码格式（已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeCharset *string `json:"CodeCharset,omitempty" name:"CodeCharset"`

	// 二维码在图片中的位置，由边界点的坐标表示（已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodePosition []*CodePosition `json:"CodePosition,omitempty" name:"CodePosition"`

	// 二维码的文本内容（已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeText *string `json:"CodeText,omitempty" name:"CodeText"`

	// 二维码的类型：1:ONED_BARCODE，2:QRCOD，3:WXCODE，4:PDF417，5:DATAMATRIX（已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeType *int64 `json:"CodeType,omitempty" name:"CodeType"`
}

type CodeDetect struct {
	// 检测是否成功，0：成功，-1：出错
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModerationCode *int64 `json:"ModerationCode,omitempty" name:"ModerationCode"`

	// 从图片中检测到的二维码，可能为多个
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModerationDetail []*CodeDetail `json:"ModerationDetail,omitempty" name:"ModerationDetail"`
}

type CodePosition struct {
	// 二维码边界点X轴坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	FloatX *float64 `json:"FloatX,omitempty" name:"FloatX"`

	// 二维码边界点Y轴坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	FloatY *float64 `json:"FloatY,omitempty" name:"FloatY"`
}

type Coordinate struct {
	// 宽度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 左上角纵坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cy *int64 `json:"Cy,omitempty" name:"Cy"`

	// 左上角横坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cx *int64 `json:"Cx,omitempty" name:"Cx"`

	// 高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

// Predefined struct for user
type CreateKeywordsSamplesRequestParams struct {
	// 关键词库信息：单次限制写入2000个，词库总容量不可超过10000个。
	UserKeywords []*UserKeyword `json:"UserKeywords,omitempty" name:"UserKeywords"`

	// 词库ID
	LibID *string `json:"LibID,omitempty" name:"LibID"`
}

type CreateKeywordsSamplesRequest struct {
	*tchttp.BaseRequest
	
	// 关键词库信息：单次限制写入2000个，词库总容量不可超过10000个。
	UserKeywords []*UserKeyword `json:"UserKeywords,omitempty" name:"UserKeywords"`

	// 词库ID
	LibID *string `json:"LibID,omitempty" name:"LibID"`
}

func (r *CreateKeywordsSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeywordsSamplesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserKeywords")
	delete(f, "LibID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKeywordsSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeywordsSamplesResponseParams struct {
	// 添加成功的关键词ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleIDs []*string `json:"SampleIDs,omitempty" name:"SampleIDs"`

	// 成功入库关键词列表
	SuccessInfos []*UserKeywordInfo `json:"SuccessInfos,omitempty" name:"SuccessInfos"`

	// 重复关键词列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DupInfos []*UserKeywordInfo `json:"DupInfos,omitempty" name:"DupInfos"`

	// 无效关键词列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvalidSamples []*InvalidSample `json:"InvalidSamples,omitempty" name:"InvalidSamples"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateKeywordsSamplesResponse struct {
	*tchttp.BaseResponse
	Response *CreateKeywordsSamplesResponseParams `json:"Response"`
}

func (r *CreateKeywordsSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeywordsSamplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomResult struct {
	// 命中的自定义关键词
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 自定义词库名称
	LibName *string `json:"LibName,omitempty" name:"LibName"`

	// 自定义库id
	LibId *string `json:"LibId,omitempty" name:"LibId"`

	// 命中的自定义关键词的类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

// Predefined struct for user
type DeleteLibSamplesRequestParams struct {
	// 关键词ID列表
	SampleIDs []*string `json:"SampleIDs,omitempty" name:"SampleIDs"`

	// 词库ID
	LibID *string `json:"LibID,omitempty" name:"LibID"`

	// 关键词内容列表
	SampleContents []*string `json:"SampleContents,omitempty" name:"SampleContents"`
}

type DeleteLibSamplesRequest struct {
	*tchttp.BaseRequest
	
	// 关键词ID列表
	SampleIDs []*string `json:"SampleIDs,omitempty" name:"SampleIDs"`

	// 词库ID
	LibID *string `json:"LibID,omitempty" name:"LibID"`

	// 关键词内容列表
	SampleContents []*string `json:"SampleContents,omitempty" name:"SampleContents"`
}

func (r *DeleteLibSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLibSamplesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SampleIDs")
	delete(f, "LibID")
	delete(f, "SampleContents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLibSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLibSamplesResponseParams struct {
	// 删除成功的数量
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 每个关键词删除的结果
	Details []*DeleteSampleDetails `json:"Details,omitempty" name:"Details"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLibSamplesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLibSamplesResponseParams `json:"Response"`
}

func (r *DeleteLibSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLibSamplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSampleDetails struct {
	// 关键词ID
	SampleID *string `json:"SampleID,omitempty" name:"SampleID"`

	// 关键词内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 是否删除成功
	Deleted *bool `json:"Deleted,omitempty" name:"Deleted"`

	// 错误信息
	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
}

// Predefined struct for user
type DescribeKeywordsLibsRequestParams struct {
	// 单页条数，最大为100条
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 条数偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤器(支持LibName模糊查询,CustomLibIDs词库id列表过滤)
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeKeywordsLibsRequest struct {
	*tchttp.BaseRequest
	
	// 单页条数，最大为100条
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 条数偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤器(支持LibName模糊查询,CustomLibIDs词库id列表过滤)
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeKeywordsLibsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeywordsLibsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeywordsLibsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeywordsLibsResponseParams struct {
	// 词库记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 词库详情
	Infos []*KeywordsLibInfo `json:"Infos,omitempty" name:"Infos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeKeywordsLibsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeywordsLibsResponseParams `json:"Response"`
}

func (r *DescribeKeywordsLibsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeywordsLibsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibSamplesRequestParams struct {
	// 单页条数，最大为100条
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 条数偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 词库ID
	LibID *string `json:"LibID,omitempty" name:"LibID"`

	// 词内容过滤
	Content *string `json:"Content,omitempty" name:"Content"`

	// 违规类型列表过滤
	EvilTypeList []*int64 `json:"EvilTypeList,omitempty" name:"EvilTypeList"`

	// 样本词ID列表过滤
	SampleIDs []*string `json:"SampleIDs,omitempty" name:"SampleIDs"`
}

type DescribeLibSamplesRequest struct {
	*tchttp.BaseRequest
	
	// 单页条数，最大为100条
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 条数偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 词库ID
	LibID *string `json:"LibID,omitempty" name:"LibID"`

	// 词内容过滤
	Content *string `json:"Content,omitempty" name:"Content"`

	// 违规类型列表过滤
	EvilTypeList []*int64 `json:"EvilTypeList,omitempty" name:"EvilTypeList"`

	// 样本词ID列表过滤
	SampleIDs []*string `json:"SampleIDs,omitempty" name:"SampleIDs"`
}

func (r *DescribeLibSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibSamplesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "LibID")
	delete(f, "Content")
	delete(f, "EvilTypeList")
	delete(f, "SampleIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibSamplesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibSamplesResponseParams struct {
	// 词记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 词详情
	Infos []*UserKeywordInfo `json:"Infos,omitempty" name:"Infos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLibSamplesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibSamplesResponseParams `json:"Response"`
}

func (r *DescribeLibSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibSamplesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailResult struct {
	// 该标签下命中的关键词
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 恶意类型
	// 100：正常
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂
	// 20105：广告引流 
	// 24001：暴恐
	EvilType *uint64 `json:"EvilType,omitempty" name:"EvilType"`

	// 该标签模型命中的分值
	Score *uint64 `json:"Score,omitempty" name:"Score"`

	// 恶意标签，Normal：正常，Polity：涉政，Porn：色情，Illegal：违法，Abuse：谩骂，Terror：暴恐，Ad：广告，Custom：自定义关键词
	EvilLabel *string `json:"EvilLabel,omitempty" name:"EvilLabel"`
}

type Device struct {
	// IOS设备，IDFV - Identifier For Vendor（应用开发商标识符）
	IDFV *string `json:"IDFV,omitempty" name:"IDFV"`

	// 设备指纹Token
	TokenId *string `json:"TokenId,omitempty" name:"TokenId"`

	// 用户IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// Mac地址
	Mac *string `json:"Mac,omitempty" name:"Mac"`

	// IOS设备，Identifier For Advertising（广告标识符）
	IDFA *string `json:"IDFA,omitempty" name:"IDFA"`

	// 设备指纹ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 设备序列号
	IMEI *string `json:"IMEI,omitempty" name:"IMEI"`
}

type Filters struct {
	// 查询字段
	Name *string `json:"Name,omitempty" name:"Name"`

	// 查询值
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ImageData struct {
	// 恶意类型
	// 100：正常 
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 20103：性感
	// 24001：暴恐
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 图片性感详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	HotDetect *ImageHotDetect `json:"HotDetect,omitempty" name:"HotDetect"`

	// 是否恶意 0：正常 1：可疑
	EvilFlag *int64 `json:"EvilFlag,omitempty" name:"EvilFlag"`

	// 图片二维码详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeDetect *CodeDetect `json:"CodeDetect,omitempty" name:"CodeDetect"`

	// 图片涉政详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolityDetect *ImagePolityDetect `json:"PolityDetect,omitempty" name:"PolityDetect"`

	// 图片违法详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	IllegalDetect *ImageIllegalDetect `json:"IllegalDetect,omitempty" name:"IllegalDetect"`

	// 图片涉黄详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornDetect *ImagePornDetect `json:"PornDetect,omitempty" name:"PornDetect"`

	// 图片暴恐详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerrorDetect *ImageTerrorDetect `json:"TerrorDetect,omitempty" name:"TerrorDetect"`

	// 图片OCR详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	OCRDetect *OCRDetect `json:"OCRDetect,omitempty" name:"OCRDetect"`

	// logo详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogoDetect *LogoDetail `json:"LogoDetect,omitempty" name:"LogoDetect"`

	// 图片相似度详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Similar *Similar `json:"Similar,omitempty" name:"Similar"`

	// 手机检测详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneDetect *PhoneDetect `json:"PhoneDetect,omitempty" name:"PhoneDetect"`
}

type ImageHotDetect struct {
	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 恶意类型
	// 100：正常
	// 20103：性感
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 性感标签：性感特征中文描述
	Labels []*string `json:"Labels,omitempty" name:"Labels"`

	// 性感分：分值范围 0-100，分数越高性感倾向越明显
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 处置判定 0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`
}

type ImageIllegalDetect struct {
	// 恶意类型
	// 100：正常 
	// 20006：涉毒违法
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 违法标签：返回违法特征中文描述，如赌桌，枪支
	Labels []*string `json:"Labels,omitempty" name:"Labels"`

	// 违法分：分值范围 0-100，分数越高违法倾向越明显
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

// Predefined struct for user
type ImageModerationRequestParams struct {
	// 文件地址
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 文件MD5值
	FileMD5 *string `json:"FileMD5,omitempty" name:"FileMD5"`

	// 文件内容 Base64,与FileUrl必须二填一
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
}

type ImageModerationRequest struct {
	*tchttp.BaseRequest
	
	// 文件地址
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 文件MD5值
	FileMD5 *string `json:"FileMD5,omitempty" name:"FileMD5"`

	// 文件内容 Base64,与FileUrl必须二填一
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
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
	delete(f, "FileUrl")
	delete(f, "FileMD5")
	delete(f, "FileContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageModerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageModerationResponseParams struct {
	// 业务返回码
	BusinessCode *int64 `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 识别结果
	Data *ImageData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type ImagePolityDetect struct {
	// 恶意类型
	// 100：正常 
	// 20001：政治
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定  0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 命中的人脸名称
	FaceNames []*string `json:"FaceNames,omitempty" name:"FaceNames"`

	// 命中的logo标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolityLogoDetail []*Logo `json:"PolityLogoDetail,omitempty" name:"PolityLogoDetail"`

	// 命中的政治物品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolityItems []*string `json:"PolityItems,omitempty" name:"PolityItems"`

	// 政治（人脸）分：分值范围 0-100，分数越高可疑程度越高
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`
}

type ImagePornDetect struct {
	// 恶意类型
	// 100：正常
	// 20002：色情
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 色情标签：色情特征中文描述
	Labels []*string `json:"Labels,omitempty" name:"Labels"`

	// 色情分：分值范围 0-100，分数越高色情倾向越明显
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type ImageTerrorDetect struct {
	// 关键词明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 恶意类型
	// 100：正常
	// 24001：暴恐
	// 注意：此字段可能返回 null，表示取不到有效值。
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 暴恐标签：返回暴恐特征中文描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*string `json:"Labels,omitempty" name:"Labels"`

	// 暴恐分：分值范围0--100，分数越高暴恐倾向越明显
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 处置判定 0：正常 1：可疑
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`
}

type InvalidSample struct {
	// 关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 无效代码:1-标签不存在;2-词过长;3-词类型不匹配;4-备注超长
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvalidCode *int64 `json:"InvalidCode,omitempty" name:"InvalidCode"`

	// 无效描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvalidMessage *string `json:"InvalidMessage,omitempty" name:"InvalidMessage"`
}

type KeywordsLibInfo struct {
	// 关键词库ID
	ID *string `json:"ID,omitempty" name:"ID"`

	// 关键词库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LibName *string `json:"LibName,omitempty" name:"LibName"`

	// 关键词库描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitempty" name:"Describe"`

	// 关键词库创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 审核建议(Review/Block)
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 匹配模式(ExactMatch/FuzzyMatch)
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// 关联策略BizType列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizTypes []*string `json:"BizTypes,omitempty" name:"BizTypes"`
}

type Logo struct {
	// logo图标置信度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// logo图标坐标信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RrectF *RrectF `json:"RrectF,omitempty" name:"RrectF"`

	// logo图标名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type LogoDetail struct {
	// 命中的Applogo详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppLogoDetail []*Logo `json:"AppLogoDetail,omitempty" name:"AppLogoDetail"`
}

type OCRDetect struct {
	// 识别到的详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item []*OCRItem `json:"Item,omitempty" name:"Item"`

	// 识别到的文本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextInfo *string `json:"TextInfo,omitempty" name:"TextInfo"`
}

type OCRItem struct {
	// 检测到的文本坐标信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextPosition *Coordinate `json:"TextPosition,omitempty" name:"TextPosition"`

	// 文本命中恶意违规类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 检测到的文本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextContent *string `json:"TextContent,omitempty" name:"TextContent"`

	// 文本涉嫌违规分值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rate *int64 `json:"Rate,omitempty" name:"Rate"`

	// 文本命中具体标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	EvilLabel *string `json:"EvilLabel,omitempty" name:"EvilLabel"`

	// 文本命中违规的关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`
}

type PhoneDetect struct {
	// 恶意类型
	// 100：正常
	// 21000：综合
	// 注意：此字段可能返回 null，表示取不到有效值。
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 特征中文描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*string `json:"Labels,omitempty" name:"Labels"`

	// 分值范围 0-100，分数越高倾向越明显
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *int64 `json:"Score,omitempty" name:"Score"`

	// 处置判定 0：正常 1：可疑
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`
}

type RiskDetails struct {
	// 预留字段，暂时不使用
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 预留字段，暂时不用
	Lable *string `json:"Lable,omitempty" name:"Lable"`

	// 风险类别，RiskAccount，RiskIP, RiskIMEI
	Label *string `json:"Label,omitempty" name:"Label"`

	// 风险等级，1:疑似，2：恶意
	Level *int64 `json:"Level,omitempty" name:"Level"`
}

type RrectF struct {
	// logo图标宽度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *float64 `json:"Width,omitempty" name:"Width"`

	// logo纵坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cy *float64 `json:"Cy,omitempty" name:"Cy"`

	// logo横坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cx *float64 `json:"Cx,omitempty" name:"Cx"`

	// logo图标中心旋转度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rotate *float64 `json:"Rotate,omitempty" name:"Rotate"`

	// logo图标高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *float64 `json:"Height,omitempty" name:"Height"`
}

type Similar struct {
	// 恶意类型
	// 100：正常 
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：未匹配到 1：恶意 2：白样本
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 返回的种子url
	// 注意：此字段可能返回 null，表示取不到有效值。
	SeedUrl *string `json:"SeedUrl,omitempty" name:"SeedUrl"`
}

type TextData struct {
	// 恶意类型
	// 100：正常
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂
	// 20105：广告引流 
	// 24001：暴恐
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 是否恶意 0：正常 1：可疑
	EvilFlag *int64 `json:"EvilFlag,omitempty" name:"EvilFlag"`

	// 和请求中的DataId一致，原样返回
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 输出的其他信息，不同客户内容不同
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 最终使用的BizType
	BizType *uint64 `json:"BizType,omitempty" name:"BizType"`

	// 消息类输出结果
	Res *TextOutputRes `json:"Res,omitempty" name:"Res"`

	// 账号风险检测结果
	RiskDetails []*RiskDetails `json:"RiskDetails,omitempty" name:"RiskDetails"`

	// 消息类ID信息
	ID *TextOutputID `json:"ID,omitempty" name:"ID"`

	// 命中的模型分值
	Score *uint64 `json:"Score,omitempty" name:"Score"`

	// 消息类公共相关参数
	Common *TextOutputComm `json:"Common,omitempty" name:"Common"`

	// 建议值,Block：打击,Review：待复审,Normal：正常
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 命中的关键词
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 返回的详细结果
	DetailResult []*DetailResult `json:"DetailResult,omitempty" name:"DetailResult"`

	// 返回的自定义词库结果
	CustomResult []*CustomResult `json:"CustomResult,omitempty" name:"CustomResult"`

	// 恶意标签，Normal：正常，Polity：涉政，Porn：色情，Illegal：违法，Abuse：谩骂，Terror：暴恐，Ad：广告，Custom：自定义关键词
	EvilLabel *string `json:"EvilLabel,omitempty" name:"EvilLabel"`
}

// Predefined struct for user
type TextModerationRequestParams struct {
	// 文本内容Base64编码。原文长度需小于15000字节，即5000个汉字以内。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 数据ID，英文字母、下划线、-组成，不超过64个字符
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 该字段用于标识业务场景。您可以在内容安全控制台创建对应的ID，配置不同的内容审核策略，通过接口调用，默认不填为0，后端使用默认策略
	BizType *uint64 `json:"BizType,omitempty" name:"BizType"`

	// 用户相关信息
	User *User `json:"User,omitempty" name:"User"`

	// 业务应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 设备相关信息
	Device *Device `json:"Device,omitempty" name:"Device"`
}

type TextModerationRequest struct {
	*tchttp.BaseRequest
	
	// 文本内容Base64编码。原文长度需小于15000字节，即5000个汉字以内。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 数据ID，英文字母、下划线、-组成，不超过64个字符
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 该字段用于标识业务场景。您可以在内容安全控制台创建对应的ID，配置不同的内容审核策略，通过接口调用，默认不填为0，后端使用默认策略
	BizType *uint64 `json:"BizType,omitempty" name:"BizType"`

	// 用户相关信息
	User *User `json:"User,omitempty" name:"User"`

	// 业务应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 设备相关信息
	Device *Device `json:"Device,omitempty" name:"Device"`
}

func (r *TextModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextModerationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "DataId")
	delete(f, "BizType")
	delete(f, "User")
	delete(f, "SdkAppId")
	delete(f, "Device")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextModerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextModerationResponseParams struct {
	// 业务返回码
	BusinessCode *int64 `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 识别结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TextData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextModerationResponse struct {
	*tchttp.BaseResponse
	Response *TextModerationResponseParams `json:"Response"`
}

func (r *TextModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextModerationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextOutputComm struct {
	// 接口唯一ID，旁路调用接口返回有该字段，标识唯一接口
	BUCtrlID *int64 `json:"BUCtrlID,omitempty" name:"BUCtrlID"`

	// 消息发送时间
	SendTime *int64 `json:"SendTime,omitempty" name:"SendTime"`

	// 接入业务的唯一ID
	AppID *int64 `json:"AppID,omitempty" name:"AppID"`

	// 请求字段里的Common.Uin
	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
}

type TextOutputID struct {
	// 接入业务的唯一ID
	MsgID *string `json:"MsgID,omitempty" name:"MsgID"`

	// 用户账号uin，对应请求协议里的Content.User.Uin。旁路结果有回带，串联结果无该字段
	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type TextOutputRes struct {
	// 操作人,信安处理人企业微信ID
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 恶意类型，广告（10001）， 政治（20001）， 色情（20002）， 社会事件（20004）， 暴力（20011）， 低俗（20012）， 违法犯罪（20006）， 欺诈（20008）， 版权（20013）， 谣言（20104）， 其他（21000）
	ResultType *int64 `json:"ResultType,omitempty" name:"ResultType"`

	// 恶意操作码，
	// 删除（1）， 通过（2）， 先审后发（100012）
	ResultCode *int64 `json:"ResultCode,omitempty" name:"ResultCode"`

	// 操作结果备注说明
	ResultMsg *string `json:"ResultMsg,omitempty" name:"ResultMsg"`
}

type User struct {
	// 用户等级，默认0 未知 1 低 2 中 3 高
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// 性别 默认0 未知 1 男性 2 女性
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 年龄 默认0 未知
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 用户账号ID，如填写，会根据账号历史恶意情况，判定消息有害结果，特别是有利于可疑恶意情况下的辅助判断。账号可以填写微信uin、QQ号、微信openid、QQopenid、字符串等。该字段和账号类别确定唯一账号。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 手机号
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 账号类别，"1-微信uin 2-QQ号 3-微信群uin 4-qq群号 5-微信openid 6-QQopenid 7-其它string"
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
}

type UserKeyword struct {
	// 关键词内容：最多40个字符，并且符合词类型的规则
	Content *string `json:"Content,omitempty" name:"Content"`

	// 关键词类型，取值范围为："Normal","Polity","Porn","Ad","Illegal","Abuse","Terror","Spam"
	Label *string `json:"Label,omitempty" name:"Label"`

	// 关键词备注：最多100个字符。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 词类型：Default,Pinyin,English,CompoundWord,ExclusionWord,AffixWord
	WordType *string `json:"WordType,omitempty" name:"WordType"`
}

type UserKeywordInfo struct {
	// 关键词条ID
	ID *string `json:"ID,omitempty" name:"ID"`

	// 关键词内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 关键词标签；取值范围为："Normal","Polity","Porn","Sexy","Ad","Illegal","Abuse","Terror","Spam","Moan"
	Label *string `json:"Label,omitempty" name:"Label"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 词类型：Default,Pinyin,English,CompoundWord,ExclusionWord,AffixWord
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordType *string `json:"WordType,omitempty" name:"WordType"`
}