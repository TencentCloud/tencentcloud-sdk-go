// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20180321

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type BoundingBox struct {
	// 左上顶点x坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 左上顶点y坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 宽
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 高
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

// Predefined struct for user
type FileTranslateRequestParams struct {
	// 源语言，支持
	// zh：简体中文
	// zh-HK：繁体中文
	// zh-TW：繁体中文
	// zh-TR：繁体中文
	// en：英语
	// ar：阿拉伯语
	// de：德语
	// es：西班牙语
	// fr：法语
	// it：意大利语
	// ja：日语
	// pt：葡萄牙语
	// ru：俄语
	// ko：韩语
	// tr：土耳其语
	// vi：越南语
	// th：泰语
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// zh（简体中文）：en（英语）、ar（阿拉伯语）、de（德语）、es（西班牙语）、fr（法语）、it（意大利语）、ja（日语）、pt（葡萄牙语）、ru（俄语）、ko（韩语）、tr（土耳其语）、vi（越南语）、th（泰语）
	// zh-HK（繁体中文）：en（英语）、ar（阿拉伯语）、de（德语）、es（西班牙语）、fr（法语）、it（意大利语）、ja（日语）、pt（葡萄牙语）、ru（俄语）、ko（韩语）、tr（土耳其语）、vi（越南语）、th（泰语）
	// zh-TW（繁体中文）：en（英语）、ar（阿拉伯语）、de（德语）、es（西班牙语）、fr（法语）、it（意大利语）、ja（日语）、pt（葡萄牙语）、ru（俄语）、ko（韩语）、tr（土耳其语）、vi（越南语）、th（泰语）
	// zh-TR（繁体中文）:en（英语）、ar（阿拉伯语）、de（德语）、es（西班牙语）、fr（法语）、it（意大利语）、ja（日语）、pt（葡萄牙语）、ru（俄语）、ko（韩语）、tr（土耳其语）、vi（越南语）、th（泰语）
	// en（英语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、ar（阿拉伯语）、de（德语）、es（西班牙语）、fr（法语）、it（意大利语）、ja（日语）、pt（葡萄牙语）、ru（俄语）、ko（韩语）、tr（土耳其语）、vi（越南语）、th（泰语）
	// ar（阿拉伯语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// de（德语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// es（西班牙语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）、fr（法语）
	// fr（法语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）、es（西班牙语）
	// it（意大利语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// ja（日语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// pt（葡萄牙语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// ru（俄语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// ko（韩语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// tr（土耳其语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// vi（越南语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// th（泰语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 文档类型：可支持以下几种(pdf,docx,pptx,xlsx,txt,xml,html,markdown,properties)
	DocumentType *string `json:"DocumentType,omitnil,omitempty" name:"DocumentType"`

	// 数据来源，0：url，1：直接传文件编码后数据
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 需要翻译文件url，URL长度不能超过1000字符。文件限制如下：docx/xIsx/html/markdown文件不超过800万字符，doc/pdf/pptx文件不超过300页，txt/po文件不超过10MB，pdf/docx/pptx/xlsx不超过40MB
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 原始文档类型。该参数为高级参数，请留空，如需使用，请与工作人员确认后再使用。
	BasicDocumentType *string `json:"BasicDocumentType,omitnil,omitempty" name:"BasicDocumentType"`

	// 回调url，URL长度不能超过256字符。文件大于10MB或字符较多时，建议采用回调方式；回调时，所有内容会放入 Body 中，具体请参见[文件翻译回调说明](https://cloud.tencent.com/document/product/551/91138)。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 文件数据，当SourceType 值为1时必须填写，为0可不写。要base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。数据要小于5MB。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

type FileTranslateRequest struct {
	*tchttp.BaseRequest
	
	// 源语言，支持
	// zh：简体中文
	// zh-HK：繁体中文
	// zh-TW：繁体中文
	// zh-TR：繁体中文
	// en：英语
	// ar：阿拉伯语
	// de：德语
	// es：西班牙语
	// fr：法语
	// it：意大利语
	// ja：日语
	// pt：葡萄牙语
	// ru：俄语
	// ko：韩语
	// tr：土耳其语
	// vi：越南语
	// th：泰语
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// zh（简体中文）：en（英语）、ar（阿拉伯语）、de（德语）、es（西班牙语）、fr（法语）、it（意大利语）、ja（日语）、pt（葡萄牙语）、ru（俄语）、ko（韩语）、tr（土耳其语）、vi（越南语）、th（泰语）
	// zh-HK（繁体中文）：en（英语）、ar（阿拉伯语）、de（德语）、es（西班牙语）、fr（法语）、it（意大利语）、ja（日语）、pt（葡萄牙语）、ru（俄语）、ko（韩语）、tr（土耳其语）、vi（越南语）、th（泰语）
	// zh-TW（繁体中文）：en（英语）、ar（阿拉伯语）、de（德语）、es（西班牙语）、fr（法语）、it（意大利语）、ja（日语）、pt（葡萄牙语）、ru（俄语）、ko（韩语）、tr（土耳其语）、vi（越南语）、th（泰语）
	// zh-TR（繁体中文）:en（英语）、ar（阿拉伯语）、de（德语）、es（西班牙语）、fr（法语）、it（意大利语）、ja（日语）、pt（葡萄牙语）、ru（俄语）、ko（韩语）、tr（土耳其语）、vi（越南语）、th（泰语）
	// en（英语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、ar（阿拉伯语）、de（德语）、es（西班牙语）、fr（法语）、it（意大利语）、ja（日语）、pt（葡萄牙语）、ru（俄语）、ko（韩语）、tr（土耳其语）、vi（越南语）、th（泰语）
	// ar（阿拉伯语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// de（德语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// es（西班牙语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）、fr（法语）
	// fr（法语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）、es（西班牙语）
	// it（意大利语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// ja（日语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// pt（葡萄牙语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// ru（俄语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// ko（韩语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// tr（土耳其语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// vi（越南语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	// th（泰语）：zh（简体中文）、zh-HK（繁体中文）、zh-TW（繁体中文）、zh-TR（繁体中文）、en（英语）
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 文档类型：可支持以下几种(pdf,docx,pptx,xlsx,txt,xml,html,markdown,properties)
	DocumentType *string `json:"DocumentType,omitnil,omitempty" name:"DocumentType"`

	// 数据来源，0：url，1：直接传文件编码后数据
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 需要翻译文件url，URL长度不能超过1000字符。文件限制如下：docx/xIsx/html/markdown文件不超过800万字符，doc/pdf/pptx文件不超过300页，txt/po文件不超过10MB，pdf/docx/pptx/xlsx不超过40MB
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 原始文档类型。该参数为高级参数，请留空，如需使用，请与工作人员确认后再使用。
	BasicDocumentType *string `json:"BasicDocumentType,omitnil,omitempty" name:"BasicDocumentType"`

	// 回调url，URL长度不能超过256字符。文件大于10MB或字符较多时，建议采用回调方式；回调时，所有内容会放入 Body 中，具体请参见[文件翻译回调说明](https://cloud.tencent.com/document/product/551/91138)。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 文件数据，当SourceType 值为1时必须填写，为0可不写。要base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。数据要小于5MB。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *FileTranslateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FileTranslateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "DocumentType")
	delete(f, "SourceType")
	delete(f, "Url")
	delete(f, "BasicDocumentType")
	delete(f, "CallbackUrl")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FileTranslateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FileTranslateResponseParams struct {
	// 文件翻译的请求返回结果，包含结果查询需要的TaskId
	Data *Task `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FileTranslateResponse struct {
	*tchttp.BaseResponse
	Response *FileTranslateResponseParams `json:"Response"`
}

func (r *FileTranslateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FileTranslateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFileTranslateData struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态
	// 
	// - init：任务已初始化
	// - wait：任务等待执行
	// - success：任务执行成功
	// - fail：任务执行失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件数据，目标文件必须小于50M，否则请通过回调方式请求文件翻译接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileData *string `json:"FileData,omitnil,omitempty" name:"FileData"`

	// 错误提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 任务进度
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 本次翻译消耗的字符数	
	UsedAmount *int64 `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`
}

// Predefined struct for user
type GetFileTranslateRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetFileTranslateRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetFileTranslateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFileTranslateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFileTranslateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFileTranslateResponseParams struct {
	// 任务id
	Data *GetFileTranslateData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFileTranslateResponse struct {
	*tchttp.BaseResponse
	Response *GetFileTranslateResponseParams `json:"Response"`
}

func (r *GetFileTranslateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFileTranslateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageRecord struct {
	// 图片翻译结果
	Value []*ItemValue `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ImageTranslateLLMRequestParams struct {
	// 图片数据的Base64字符串，经Base64编码后不超过 9M，分辨率建议600*800以上，支持PNG、JPG、JPEG格式。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 目标语言，支持语言列表：
	// 
	// - 中文：zh
	// - 繁体（中国台湾）：zh-TW
	// - 繁体（中国香港）：zh-HK
	// - 英文：en
	// - 日语：ja
	// - 韩语：ko
	// - 泰语：th
	// - 越南语：vi
	// - 俄语：ru
	// - 德语：de
	// - 法语：fr
	// - 阿拉伯语：ar
	// - 西班牙语：es
	// - 意大利语：it
	// - 印度尼西亚语：id
	// - 马来西亚语：ms
	// - 葡萄牙语：pt
	// - 土耳其语：tr
	// - 
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 输入图 Url。 使用Url的时候，Data参数需要传入""。 图片限制：小于 10MB，分辨率建议600*800以上，格式支持 jpg、jpeg、png。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type ImageTranslateLLMRequest struct {
	*tchttp.BaseRequest
	
	// 图片数据的Base64字符串，经Base64编码后不超过 9M，分辨率建议600*800以上，支持PNG、JPG、JPEG格式。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 目标语言，支持语言列表：
	// 
	// - 中文：zh
	// - 繁体（中国台湾）：zh-TW
	// - 繁体（中国香港）：zh-HK
	// - 英文：en
	// - 日语：ja
	// - 韩语：ko
	// - 泰语：th
	// - 越南语：vi
	// - 俄语：ru
	// - 德语：de
	// - 法语：fr
	// - 阿拉伯语：ar
	// - 西班牙语：es
	// - 意大利语：it
	// - 印度尼西亚语：id
	// - 马来西亚语：ms
	// - 葡萄牙语：pt
	// - 土耳其语：tr
	// - 
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 输入图 Url。 使用Url的时候，Data参数需要传入""。 图片限制：小于 10MB，分辨率建议600*800以上，格式支持 jpg、jpeg、png。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

func (r *ImageTranslateLLMRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageTranslateLLMRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Target")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageTranslateLLMRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageTranslateLLMResponseParams struct {
	// 图片数据的Base64字符串，输出格式为JPG。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 原文本主要源语言。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标翻译语言。
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 图片中的全部原文本。
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// 图片中全部译文。
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`

	// 逆时针图片角度，取值范围为0-359
	Angle *float64 `json:"Angle,omitnil,omitempty" name:"Angle"`

	// 翻译详情信息
	TransDetails []*TransDetail `json:"TransDetails,omitnil,omitempty" name:"TransDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageTranslateLLMResponse struct {
	*tchttp.BaseResponse
	Response *ImageTranslateLLMResponseParams `json:"Response"`
}

func (r *ImageTranslateLLMResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageTranslateLLMResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageTranslateRequestParams struct {
	// 唯一id，返回时原样返回
	SessionUuid *string `json:"SessionUuid,omitnil,omitempty" name:"SessionUuid"`

	// doc:文档扫描
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 图片数据的Base64字符串，经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。图片中包含文字需要少于6000字符。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 源语言，支持语言列表：<li> auto：自动识别（识别为一种语言）</li> <li>zh：简体中文</li> <li>zh-TW：繁体中文</li> <li>en：英语</li> <li>ja：日语</li> <li>ko：韩语</li> <li>ru：俄语</li> <li>fr：法语</li> <li>de：德语</li> <li>it：意大利语</li> <li>es：西班牙语</li> <li>pt：葡萄牙语</li> <li>ms：马来西亚语</li> <li>th：泰语</li><li>vi：越南语</li>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下：
	// <li>zh（简体中文）：en（英语）、ja（日语）、ko（韩语）、ru（俄语）、fr（法语）、de（德语）、it（意大利语）、es（西班牙语）、pt（葡萄牙语）、ms（马来语）、th（泰语）、vi（越南语）</li>
	// <li>zh-TW（繁体中文）：en（英语）、ja（日语）、ko（韩语）、ru（俄语）、fr（法语）、de（德语）、it（意大利语）、es（西班牙语）、pt（葡萄牙语）、ms（马来语）、th（泰语）、vi（越南语）</li>
	// <li>en（英语）：zh（中文）、ja（日语）、ko（韩语）、ru（俄语）、fr（法语）、de（德语）、it（意大利语）、es（西班牙语）、pt（葡萄牙语）、ms（马来语）、th（泰语）、vi（越南语）</li>
	// <li>ja（日语）：zh（中文）、en（英语）、ko（韩语）</li>
	// <li>ko（韩语）：zh（中文）、en（英语）、ja（日语）</li>
	// <li>ru：俄语：zh（中文）、en（英语）</li>
	// <li>fr：法语：zh（中文）、en（英语）</li>
	// <li>de：德语：zh（中文）、en（英语）</li>
	// <li>it：意大利语：zh（中文）、en（英语）</li>
	// <li>es：西班牙语：zh（中文）、en（英语）</li>
	// <li>pt：葡萄牙语：zh（中文）、en（英语）</li>
	// <li>ms：马来西亚语：zh（中文）、en（英语）</li>
	// <li>th：泰语：zh（中文）、en（英语）</li>
	// <li>vi：越南语：zh（中文）、en（英语）</li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ImageTranslateRequest struct {
	*tchttp.BaseRequest
	
	// 唯一id，返回时原样返回
	SessionUuid *string `json:"SessionUuid,omitnil,omitempty" name:"SessionUuid"`

	// doc:文档扫描
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 图片数据的Base64字符串，经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。图片中包含文字需要少于6000字符。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 源语言，支持语言列表：<li> auto：自动识别（识别为一种语言）</li> <li>zh：简体中文</li> <li>zh-TW：繁体中文</li> <li>en：英语</li> <li>ja：日语</li> <li>ko：韩语</li> <li>ru：俄语</li> <li>fr：法语</li> <li>de：德语</li> <li>it：意大利语</li> <li>es：西班牙语</li> <li>pt：葡萄牙语</li> <li>ms：马来西亚语</li> <li>th：泰语</li><li>vi：越南语</li>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下：
	// <li>zh（简体中文）：en（英语）、ja（日语）、ko（韩语）、ru（俄语）、fr（法语）、de（德语）、it（意大利语）、es（西班牙语）、pt（葡萄牙语）、ms（马来语）、th（泰语）、vi（越南语）</li>
	// <li>zh-TW（繁体中文）：en（英语）、ja（日语）、ko（韩语）、ru（俄语）、fr（法语）、de（德语）、it（意大利语）、es（西班牙语）、pt（葡萄牙语）、ms（马来语）、th（泰语）、vi（越南语）</li>
	// <li>en（英语）：zh（中文）、ja（日语）、ko（韩语）、ru（俄语）、fr（法语）、de（德语）、it（意大利语）、es（西班牙语）、pt（葡萄牙语）、ms（马来语）、th（泰语）、vi（越南语）</li>
	// <li>ja（日语）：zh（中文）、en（英语）、ko（韩语）</li>
	// <li>ko（韩语）：zh（中文）、en（英语）、ja（日语）</li>
	// <li>ru：俄语：zh（中文）、en（英语）</li>
	// <li>fr：法语：zh（中文）、en（英语）</li>
	// <li>de：德语：zh（中文）、en（英语）</li>
	// <li>it：意大利语：zh（中文）、en（英语）</li>
	// <li>es：西班牙语：zh（中文）、en（英语）</li>
	// <li>pt：葡萄牙语：zh（中文）、en（英语）</li>
	// <li>ms：马来西亚语：zh（中文）、en（英语）</li>
	// <li>th：泰语：zh（中文）、en（英语）</li>
	// <li>vi：越南语：zh（中文）、en（英语）</li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *ImageTranslateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageTranslateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionUuid")
	delete(f, "Scene")
	delete(f, "Data")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageTranslateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageTranslateResponseParams struct {
	// 请求的SessionUuid返回
	SessionUuid *string `json:"SessionUuid,omitnil,omitempty" name:"SessionUuid"`

	// 源语言
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 图片翻译结果，翻译结果按识别的文本每一行独立翻译，后续会推出按段落划分并翻译的版本
	ImageRecord *ImageRecord `json:"ImageRecord,omitnil,omitempty" name:"ImageRecord"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageTranslateResponse struct {
	*tchttp.BaseResponse
	Response *ImageTranslateResponseParams `json:"Response"`
}

func (r *ImageTranslateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageTranslateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ItemValue struct {
	// 识别出的源文
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// 翻译后的译文
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`

	// X坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// Y坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 宽度
	W *int64 `json:"W,omitnil,omitempty" name:"W"`

	// 高度
	H *int64 `json:"H,omitnil,omitempty" name:"H"`
}

// Predefined struct for user
type LanguageDetectRequestParams struct {
	// 待识别的文本，文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败。单次请求的文本长度需要低于2000。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type LanguageDetectRequest struct {
	*tchttp.BaseRequest
	
	// 待识别的文本，文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败。单次请求的文本长度需要低于2000。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *LanguageDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LanguageDetectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LanguageDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LanguageDetectResponseParams struct {
	// 识别出的语言种类，参考语言列表
	// <li> zh : 中文 </li> <li> en : 英文 </li><li> jp : 日语 </li> <li> kr : 韩语 </li><li> de : 德语 </li><li> fr : 法语 </li><li> es : 西班牙文 </li> <li> it : 意大利文 </li><li> tr : 土耳其文 </li><li> ru : 俄文 </li><li> pt : 葡萄牙文 </li><li> vi : 越南文 </li><li> id : 印度尼西亚文 </li><li> ms : 马来西亚文 </li><li> th : 泰文 </li>
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LanguageDetectResponse struct {
	*tchttp.BaseResponse
	Response *LanguageDetectResponseParams `json:"Response"`
}

func (r *LanguageDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LanguageDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SpeechTranslateRequestParams struct {
	// 一段完整的语音对应一个SessionUuid
	SessionUuid *string `json:"SessionUuid,omitnil,omitempty" name:"SessionUuid"`

	// 音频中的语言类型，支持语言列表<li> zh : 中文 </li> <li> en : 英文 </li>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 翻译目标语言类型，支持的语言列表<li> zh : 中文 </li> <li> en : 英文 </li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// pcm : 146   speex : 16779154   mp3 : 83886080
	AudioFormat *int64 `json:"AudioFormat,omitnil,omitempty" name:"AudioFormat"`

	// 语音分片的序号，从0开始
	Seq *int64 `json:"Seq,omitnil,omitempty" name:"Seq"`

	// 是否最后一片语音分片，0-否，1-是
	IsEnd *int64 `json:"IsEnd,omitnil,omitempty" name:"IsEnd"`

	// 语音分片内容进行 Base64 编码后的字符串。音频内容需包含有效并可识别的文本信息。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 识别模式，该参数已废弃
	//
	// Deprecated: Mode is deprecated.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 该参数已废弃
	//
	// Deprecated: TransType is deprecated.
	TransType *int64 `json:"TransType,omitnil,omitempty" name:"TransType"`
}

type SpeechTranslateRequest struct {
	*tchttp.BaseRequest
	
	// 一段完整的语音对应一个SessionUuid
	SessionUuid *string `json:"SessionUuid,omitnil,omitempty" name:"SessionUuid"`

	// 音频中的语言类型，支持语言列表<li> zh : 中文 </li> <li> en : 英文 </li>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 翻译目标语言类型，支持的语言列表<li> zh : 中文 </li> <li> en : 英文 </li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// pcm : 146   speex : 16779154   mp3 : 83886080
	AudioFormat *int64 `json:"AudioFormat,omitnil,omitempty" name:"AudioFormat"`

	// 语音分片的序号，从0开始
	Seq *int64 `json:"Seq,omitnil,omitempty" name:"Seq"`

	// 是否最后一片语音分片，0-否，1-是
	IsEnd *int64 `json:"IsEnd,omitnil,omitempty" name:"IsEnd"`

	// 语音分片内容进行 Base64 编码后的字符串。音频内容需包含有效并可识别的文本信息。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 识别模式，该参数已废弃
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 该参数已废弃
	TransType *int64 `json:"TransType,omitnil,omitempty" name:"TransType"`
}

func (r *SpeechTranslateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SpeechTranslateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionUuid")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "AudioFormat")
	delete(f, "Seq")
	delete(f, "IsEnd")
	delete(f, "Data")
	delete(f, "ProjectId")
	delete(f, "Mode")
	delete(f, "TransType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SpeechTranslateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SpeechTranslateResponseParams struct {
	// 请求的SessionUuid直接返回
	SessionUuid *string `json:"SessionUuid,omitnil,omitempty" name:"SessionUuid"`

	// 语音识别状态 1-进行中 0-完成
	RecognizeStatus *int64 `json:"RecognizeStatus,omitnil,omitempty" name:"RecognizeStatus"`

	// 识别出的原文
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// 翻译出的译文
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`

	// 第几个语音分片
	Seq *int64 `json:"Seq,omitnil,omitempty" name:"Seq"`

	// 原语言
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 当请求的Mode参数填写bvad是，启动VadSeq。此时Seq会被设置为后台vad（静音检测）后的新序号，而VadSeq代表客户端原始Seq值
	VadSeq *int64 `json:"VadSeq,omitnil,omitempty" name:"VadSeq"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SpeechTranslateResponse struct {
	*tchttp.BaseResponse
	Response *SpeechTranslateResponseParams `json:"Response"`
}

func (r *SpeechTranslateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SpeechTranslateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Task struct {
	// 任务ID，可通过此ID在轮询接口获取识别状态与结果。注意：TaskId数据类型为字符串类型
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type TextTranslateBatchRequestParams struct {
	// 源语言，支持： 
	// auto：自动识别（识别为一种语言）
	// zh：简体中文
	// zh-TW：繁体中文
	// en：英语
	// ja：日语
	// ko：韩语
	// fr：法语
	// es：西班牙语
	// it：意大利语
	// de：德语
	// tr：土耳其语
	// ru：俄语
	// pt：葡萄牙语
	// vi：越南语
	// id：印尼语
	// th：泰语
	// ms：马来西亚语
	// ar：阿拉伯语
	// hi：印地语
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// 
	// <li> zh（简体中文）：en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）</li>
	// <li>zh-TW（繁体中文）：en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）</li>
	// <li>en（英语）：zh（中文）、zh-TW（繁体中文）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）、hi（印地语）</li>
	// <li>ja（日语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ko（韩语）</li>
	// <li>ko（韩语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ja（日语）</li>
	// <li>fr（法语）：zh（中文）、zh-TW（繁体中文）、en（英语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>es（西班牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>it（意大利语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>de（德语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>tr（土耳其语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>ru（俄语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、pt（葡萄牙语）</li>
	// <li>pt（葡萄牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）</li>
	// <li>vi（越南语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>id（印尼语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>th（泰语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>ms（马来语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>ar（阿拉伯语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>hi（印地语）：en（英语）</li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 待翻译的文本列表，批量接口可以以数组方式在一次请求中填写多个待翻译文本。文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败，请传入有效文本，html标记等非常规翻译文本可能会翻译失败。单次请求的文本长度总和需要低于6000字符。
	SourceTextList []*string `json:"SourceTextList,omitnil,omitempty" name:"SourceTextList"`

	// 需要使用的术语库列表，通过 [术语库操作指南](https://cloud.tencent.com/document/product/551/107926) 自行创建术语库获取。
	TermRepoIDList []*string `json:"TermRepoIDList,omitnil,omitempty" name:"TermRepoIDList"`

	// 需要使用的例句库列表，通过 [例句库操作指南](https://cloud.tencent.com/document/product/551/107927) 自行创建例句库获取。
	SentRepoIDList []*string `json:"SentRepoIDList,omitnil,omitempty" name:"SentRepoIDList"`
}

type TextTranslateBatchRequest struct {
	*tchttp.BaseRequest
	
	// 源语言，支持： 
	// auto：自动识别（识别为一种语言）
	// zh：简体中文
	// zh-TW：繁体中文
	// en：英语
	// ja：日语
	// ko：韩语
	// fr：法语
	// es：西班牙语
	// it：意大利语
	// de：德语
	// tr：土耳其语
	// ru：俄语
	// pt：葡萄牙语
	// vi：越南语
	// id：印尼语
	// th：泰语
	// ms：马来西亚语
	// ar：阿拉伯语
	// hi：印地语
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// 
	// <li> zh（简体中文）：en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）</li>
	// <li>zh-TW（繁体中文）：en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）</li>
	// <li>en（英语）：zh（中文）、zh-TW（繁体中文）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）、hi（印地语）</li>
	// <li>ja（日语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ko（韩语）</li>
	// <li>ko（韩语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ja（日语）</li>
	// <li>fr（法语）：zh（中文）、zh-TW（繁体中文）、en（英语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>es（西班牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>it（意大利语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>de（德语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>tr（土耳其语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>ru（俄语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、pt（葡萄牙语）</li>
	// <li>pt（葡萄牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）</li>
	// <li>vi（越南语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>id（印尼语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>th（泰语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>ms（马来语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>ar（阿拉伯语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>hi（印地语）：en（英语）</li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 待翻译的文本列表，批量接口可以以数组方式在一次请求中填写多个待翻译文本。文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败，请传入有效文本，html标记等非常规翻译文本可能会翻译失败。单次请求的文本长度总和需要低于6000字符。
	SourceTextList []*string `json:"SourceTextList,omitnil,omitempty" name:"SourceTextList"`

	// 需要使用的术语库列表，通过 [术语库操作指南](https://cloud.tencent.com/document/product/551/107926) 自行创建术语库获取。
	TermRepoIDList []*string `json:"TermRepoIDList,omitnil,omitempty" name:"TermRepoIDList"`

	// 需要使用的例句库列表，通过 [例句库操作指南](https://cloud.tencent.com/document/product/551/107927) 自行创建例句库获取。
	SentRepoIDList []*string `json:"SentRepoIDList,omitnil,omitempty" name:"SentRepoIDList"`
}

func (r *TextTranslateBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextTranslateBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "ProjectId")
	delete(f, "SourceTextList")
	delete(f, "TermRepoIDList")
	delete(f, "SentRepoIDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextTranslateBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextTranslateBatchResponseParams struct {
	// 源语言，详见入参Source
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，详见入参Target
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 翻译后的文本列表
	TargetTextList []*string `json:"TargetTextList,omitnil,omitempty" name:"TargetTextList"`

	// 本次翻译消耗的字符数
	UsedAmount *int64 `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TextTranslateBatchResponse struct {
	*tchttp.BaseResponse
	Response *TextTranslateBatchResponseParams `json:"Response"`
}

func (r *TextTranslateBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextTranslateBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextTranslateRequestParams struct {
	// 待翻译的文本，文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败，请传入有效文本，html标记等非常规翻译文本可能会翻译失败。单次请求的文本长度需要低于6000字符。
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// 源语言，支持：
	// auto：自动识别（识别为一种语言）
	// zh：简体中文
	// zh-TW：繁体中文
	// en：英语
	// ja：日语
	// ko：韩语
	// fr：法语
	// es：西班牙语
	// it：意大利语
	// de：德语
	// tr：土耳其语
	// ru：俄语
	// pt：葡萄牙语
	// vi：越南语
	// id：印尼语
	// th：泰语
	// ms：马来西亚语
	// ar：阿拉伯语
	// hi：印地语
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// 
	// <li> zh（简体中文）：zh-TW（繁体中文）、en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）</li>
	// <li>zh-TW（繁体中文）：zh（简体中文）、en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）</li>
	// <li>en（英语）：zh（中文）、zh-TW（繁体中文）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）、hi（印地语）</li>
	// <li>ja（日语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ko（韩语）</li>
	// <li>ko（韩语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ja（日语）</li>
	// <li>fr（法语）：zh（中文）、zh-TW（繁体中文）、en（英语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>es（西班牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>it（意大利语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>de（德语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>tr（土耳其语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>ru（俄语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、pt（葡萄牙语）</li>
	// <li>pt（葡萄牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）</li>
	// <li>vi（越南语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>id（印尼语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>th（泰语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>ms（马来语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>ar（阿拉伯语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>hi（印地语）：en（英语）</li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用来标记不希望被翻译的文本内容，如句子中的特殊符号、人名、地名等；每次请求只支持配置一个不被翻译的单词；仅支持配置人名、地名等名词，不要配置动词或短语，否则会影响翻译结果。
	UntranslatedText *string `json:"UntranslatedText,omitnil,omitempty" name:"UntranslatedText"`

	// 需要使用的术语库列表，通过 [术语库操作指南](https://cloud.tencent.com/document/product/551/107926) 自行创建术语库获取。
	TermRepoIDList []*string `json:"TermRepoIDList,omitnil,omitempty" name:"TermRepoIDList"`

	// 需要使用的例句库列表，通过 [例句库操作指南](https://cloud.tencent.com/document/product/551/107927) 自行创建例句库获取。
	SentRepoIDList []*string `json:"SentRepoIDList,omitnil,omitempty" name:"SentRepoIDList"`
}

type TextTranslateRequest struct {
	*tchttp.BaseRequest
	
	// 待翻译的文本，文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败，请传入有效文本，html标记等非常规翻译文本可能会翻译失败。单次请求的文本长度需要低于6000字符。
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// 源语言，支持：
	// auto：自动识别（识别为一种语言）
	// zh：简体中文
	// zh-TW：繁体中文
	// en：英语
	// ja：日语
	// ko：韩语
	// fr：法语
	// es：西班牙语
	// it：意大利语
	// de：德语
	// tr：土耳其语
	// ru：俄语
	// pt：葡萄牙语
	// vi：越南语
	// id：印尼语
	// th：泰语
	// ms：马来西亚语
	// ar：阿拉伯语
	// hi：印地语
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// 
	// <li> zh（简体中文）：zh-TW（繁体中文）、en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）</li>
	// <li>zh-TW（繁体中文）：zh（简体中文）、en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）</li>
	// <li>en（英语）：zh（中文）、zh-TW（繁体中文）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）、hi（印地语）</li>
	// <li>ja（日语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ko（韩语）</li>
	// <li>ko（韩语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ja（日语）</li>
	// <li>fr（法语）：zh（中文）、zh-TW（繁体中文）、en（英语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>es（西班牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>it（意大利语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>de（德语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>tr（土耳其语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>ru（俄语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、pt（葡萄牙语）</li>
	// <li>pt（葡萄牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）</li>
	// <li>vi（越南语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>id（印尼语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>th（泰语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>ms（马来语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>ar（阿拉伯语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>hi（印地语）：en（英语）</li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用来标记不希望被翻译的文本内容，如句子中的特殊符号、人名、地名等；每次请求只支持配置一个不被翻译的单词；仅支持配置人名、地名等名词，不要配置动词或短语，否则会影响翻译结果。
	UntranslatedText *string `json:"UntranslatedText,omitnil,omitempty" name:"UntranslatedText"`

	// 需要使用的术语库列表，通过 [术语库操作指南](https://cloud.tencent.com/document/product/551/107926) 自行创建术语库获取。
	TermRepoIDList []*string `json:"TermRepoIDList,omitnil,omitempty" name:"TermRepoIDList"`

	// 需要使用的例句库列表，通过 [例句库操作指南](https://cloud.tencent.com/document/product/551/107927) 自行创建例句库获取。
	SentRepoIDList []*string `json:"SentRepoIDList,omitnil,omitempty" name:"SentRepoIDList"`
}

func (r *TextTranslateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextTranslateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceText")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "ProjectId")
	delete(f, "UntranslatedText")
	delete(f, "TermRepoIDList")
	delete(f, "SentRepoIDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextTranslateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextTranslateResponseParams struct {
	// 翻译后的文本
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`

	// 源语言，详见入参Source
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，详见入参Target
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 本次翻译消耗的字符数
	UsedAmount *int64 `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TextTranslateResponse struct {
	*tchttp.BaseResponse
	Response *TextTranslateResponseParams `json:"Response"`
}

func (r *TextTranslateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextTranslateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransDetail struct {
	// 当前行的原文本
	SourceLineText *string `json:"SourceLineText,omitnil,omitempty" name:"SourceLineText"`

	// 当前行的译文
	TargetLineText *string `json:"TargetLineText,omitnil,omitempty" name:"TargetLineText"`

	// 段落文本框位置
	BoundingBox *BoundingBox `json:"BoundingBox,omitnil,omitempty" name:"BoundingBox"`

	// 行数
	LinesCount *int64 `json:"LinesCount,omitnil,omitempty" name:"LinesCount"`

	// 行高
	LineHeight *int64 `json:"LineHeight,omitnil,omitempty" name:"LineHeight"`

	// 正常段落spam_code字段为0；如果存在spam_code字段且值大于0（1: 命中垃圾检查；2: 命中安全策略；3: 其他。），则命中安全检查被过滤。
	SpamCode *int64 `json:"SpamCode,omitnil,omitempty" name:"SpamCode"`
}