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

package v20180321

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type FileTranslateRequestParams struct {
	// 源语言，支持
	// zh:简体中文
	// zh-HK：繁体中文
	// zh-TW : 繁体中文
	// zh-TR:  繁体中文
	// en ：英语
	// ar：阿拉伯语
	// de：德语
	// es：西班牙语
	// fr：法语
	// it：意大利语
	// ja：日语
	// pt：葡萄牙语
	// ru：俄语
	// ko：韩语
	// km：高棉语
	// lo：老挝语
	Source *string `json:"Source,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// zh（简体中文）： en （英语）、 ar (阿拉伯语）、 de （德语）、  es（西班牙语） 、fr（法语）、  it（意大利语） 、 ja （日语）、 pt （葡萄牙语）、 ru（俄语）、  ko（韩语）、 km（高棉语）、   lo（老挝语）
	// zh-HK（繁体中文） ：en （英语）、 ar (阿拉伯语）、 de （德语）、  es（西班牙语） 、fr（法语）、  it（意大利语） 、 ja （日语）、 pt （葡萄牙语）、 ru（俄语）、  ko（韩语）、 km（高棉语）、   lo（老挝语）
	// zh-TW（繁体中文）：en （英语）、 ar (阿拉伯语）、 de （德语）、  es（西班牙语） 、fr（法语）、  it（意大利语） 、 ja （日语）、 pt （葡萄牙语）、 ru（俄语）、  ko（韩语）、 km（高棉语）、   lo（老挝语）
	// zh-TR 繁体中文 : en （英语）、 ar (阿拉伯语）、 de （德语）、  es（西班牙语） 、fr（法语）、  it（意大利语） 、 ja （日语）、 pt （葡萄牙语）、 ru（俄语）、  ko（韩语）、 km（高棉语）、   lo（老挝语）
	// en （英语） ：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、 zh-TR(繁体中文）、 ar (阿拉伯语）、 de （德语）、  es（西班牙语） 、fr（法语）、  it（意大利语） 、 ja （日语）、 pt （葡萄牙语）、 ru（俄语）、  ko（韩语）、 km（高棉语）、   lo（老挝语）
	// ar（阿拉伯语） ：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// de（德语 ）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// es（西班牙语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// fr（法语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// it（意大利语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// ja（日语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// pt（葡萄牙语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// ru（俄语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// ko（韩语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// km（高棉语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// lo（老挝语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	Target *string `json:"Target,omitempty" name:"Target"`

	// 文档类型：可支持以下几种(pdf,docx,pptx,xlsx,txt,xml,html,markdown,properties)
	DocumentType *string `json:"DocumentType,omitempty" name:"DocumentType"`

	// 数据来源，0：url，1：直接传文件编码后数据
	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`

	// 需要翻译文件url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 原始文档类型
	BasicDocumentType *string `json:"BasicDocumentType,omitempty" name:"BasicDocumentType"`

	// 回调url
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 文件数据，当SourceType 值为1时必须填写，为0可不写。要base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。数据要小于5MB。
	Data *string `json:"Data,omitempty" name:"Data"`
}

type FileTranslateRequest struct {
	*tchttp.BaseRequest
	
	// 源语言，支持
	// zh:简体中文
	// zh-HK：繁体中文
	// zh-TW : 繁体中文
	// zh-TR:  繁体中文
	// en ：英语
	// ar：阿拉伯语
	// de：德语
	// es：西班牙语
	// fr：法语
	// it：意大利语
	// ja：日语
	// pt：葡萄牙语
	// ru：俄语
	// ko：韩语
	// km：高棉语
	// lo：老挝语
	Source *string `json:"Source,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// zh（简体中文）： en （英语）、 ar (阿拉伯语）、 de （德语）、  es（西班牙语） 、fr（法语）、  it（意大利语） 、 ja （日语）、 pt （葡萄牙语）、 ru（俄语）、  ko（韩语）、 km（高棉语）、   lo（老挝语）
	// zh-HK（繁体中文） ：en （英语）、 ar (阿拉伯语）、 de （德语）、  es（西班牙语） 、fr（法语）、  it（意大利语） 、 ja （日语）、 pt （葡萄牙语）、 ru（俄语）、  ko（韩语）、 km（高棉语）、   lo（老挝语）
	// zh-TW（繁体中文）：en （英语）、 ar (阿拉伯语）、 de （德语）、  es（西班牙语） 、fr（法语）、  it（意大利语） 、 ja （日语）、 pt （葡萄牙语）、 ru（俄语）、  ko（韩语）、 km（高棉语）、   lo（老挝语）
	// zh-TR 繁体中文 : en （英语）、 ar (阿拉伯语）、 de （德语）、  es（西班牙语） 、fr（法语）、  it（意大利语） 、 ja （日语）、 pt （葡萄牙语）、 ru（俄语）、  ko（韩语）、 km（高棉语）、   lo（老挝语）
	// en （英语） ：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、 zh-TR(繁体中文）、 ar (阿拉伯语）、 de （德语）、  es（西班牙语） 、fr（法语）、  it（意大利语） 、 ja （日语）、 pt （葡萄牙语）、 ru（俄语）、  ko（韩语）、 km（高棉语）、   lo（老挝语）
	// ar（阿拉伯语） ：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// de（德语 ）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// es（西班牙语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// fr（法语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// it（意大利语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// ja（日语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// pt（葡萄牙语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// ru（俄语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// ko（韩语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// km（高棉语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	// lo（老挝语）：zh（简体中文）、zh-HK（繁体中文）、 zh-TW（繁体中文)、zh-TR(繁体中文）
	Target *string `json:"Target,omitempty" name:"Target"`

	// 文档类型：可支持以下几种(pdf,docx,pptx,xlsx,txt,xml,html,markdown,properties)
	DocumentType *string `json:"DocumentType,omitempty" name:"DocumentType"`

	// 数据来源，0：url，1：直接传文件编码后数据
	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`

	// 需要翻译文件url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 原始文档类型
	BasicDocumentType *string `json:"BasicDocumentType,omitempty" name:"BasicDocumentType"`

	// 回调url
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 文件数据，当SourceType 值为1时必须填写，为0可不写。要base64编码(采用python语言时注意读取文件应该为string而不是byte，以byte格式读取后要decode()。编码后的数据不可带有回车换行符)。数据要小于5MB。
	Data *string `json:"Data,omitempty" name:"Data"`
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
	Data *Task `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 文件数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileData *string `json:"FileData,omitempty" name:"FileData"`

	// 错误提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 翻译进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
}

// Predefined struct for user
type GetFileTranslateRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type GetFileTranslateRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
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
	Data *GetFileTranslateData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Value []*ItemValue `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type ImageTranslateRequestParams struct {
	// 唯一id，返回时原样返回
	SessionUuid *string `json:"SessionUuid,omitempty" name:"SessionUuid"`

	// doc:文档扫描
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 图片数据的Base64字符串，图片大小上限为4M，建议对源图片进行一定程度压缩
	Data *string `json:"Data,omitempty" name:"Data"`

	// 源语言，支持语言列表：<li> auto：自动识别（识别为一种语言）</li> <li>zh：简体中文</li> <li>zh-TW：繁体中文</li> <li>en：英语</li> <li>ja：日语</li> <li>ko：韩语</li> <li>ru：俄语</li> <li>fr：法语</li> <li>de：德语</li> <li>it：意大利语</li> <li>es：西班牙语</li> <li>pt：葡萄牙语</li> <li>ms：马来西亚语</li> <li>th：泰语</li><li>vi：越南语</li>
	Source *string `json:"Source,omitempty" name:"Source"`

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
	Target *string `json:"Target,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type ImageTranslateRequest struct {
	*tchttp.BaseRequest
	
	// 唯一id，返回时原样返回
	SessionUuid *string `json:"SessionUuid,omitempty" name:"SessionUuid"`

	// doc:文档扫描
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 图片数据的Base64字符串，图片大小上限为4M，建议对源图片进行一定程度压缩
	Data *string `json:"Data,omitempty" name:"Data"`

	// 源语言，支持语言列表：<li> auto：自动识别（识别为一种语言）</li> <li>zh：简体中文</li> <li>zh-TW：繁体中文</li> <li>en：英语</li> <li>ja：日语</li> <li>ko：韩语</li> <li>ru：俄语</li> <li>fr：法语</li> <li>de：德语</li> <li>it：意大利语</li> <li>es：西班牙语</li> <li>pt：葡萄牙语</li> <li>ms：马来西亚语</li> <li>th：泰语</li><li>vi：越南语</li>
	Source *string `json:"Source,omitempty" name:"Source"`

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
	Target *string `json:"Target,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
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
	SessionUuid *string `json:"SessionUuid,omitempty" name:"SessionUuid"`

	// 源语言
	Source *string `json:"Source,omitempty" name:"Source"`

	// 目标语言
	Target *string `json:"Target,omitempty" name:"Target"`

	// 图片翻译结果，翻译结果按识别的文本每一行独立翻译，后续会推出按段落划分并翻译的版本
	ImageRecord *ImageRecord `json:"ImageRecord,omitempty" name:"ImageRecord"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SourceText *string `json:"SourceText,omitempty" name:"SourceText"`

	// 翻译后的译文
	TargetText *string `json:"TargetText,omitempty" name:"TargetText"`

	// X坐标
	X *int64 `json:"X,omitempty" name:"X"`

	// Y坐标
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 宽度
	W *int64 `json:"W,omitempty" name:"W"`

	// 高度
	H *int64 `json:"H,omitempty" name:"H"`
}

// Predefined struct for user
type LanguageDetectRequestParams struct {
	// 待识别的文本，文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败。单次请求的文本长度需要低于2000。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type LanguageDetectRequest struct {
	*tchttp.BaseRequest
	
	// 待识别的文本，文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败。单次请求的文本长度需要低于2000。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
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
	Lang *string `json:"Lang,omitempty" name:"Lang"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SessionUuid *string `json:"SessionUuid,omitempty" name:"SessionUuid"`

	// 音频中的语言类型，支持语言列表<li> zh : 中文 </li> <li> en : 英文 </li>
	Source *string `json:"Source,omitempty" name:"Source"`

	// 翻译目标语言类型，支持的语言列表<li> zh : 中文 </li> <li> en : 英文 </li>
	Target *string `json:"Target,omitempty" name:"Target"`

	// pcm : 146   speex : 16779154   mp3 : 83886080
	AudioFormat *int64 `json:"AudioFormat,omitempty" name:"AudioFormat"`

	// 语音分片的序号，从0开始
	Seq *int64 `json:"Seq,omitempty" name:"Seq"`

	// 是否最后一片语音分片，0-否，1-是
	IsEnd *int64 `json:"IsEnd,omitempty" name:"IsEnd"`

	// 语音分片内容进行 Base64 编码后的字符串。音频内容需包含有效并可识别的文本信息。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 识别模式，该参数已废弃
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 该参数已废弃
	TransType *int64 `json:"TransType,omitempty" name:"TransType"`
}

type SpeechTranslateRequest struct {
	*tchttp.BaseRequest
	
	// 一段完整的语音对应一个SessionUuid
	SessionUuid *string `json:"SessionUuid,omitempty" name:"SessionUuid"`

	// 音频中的语言类型，支持语言列表<li> zh : 中文 </li> <li> en : 英文 </li>
	Source *string `json:"Source,omitempty" name:"Source"`

	// 翻译目标语言类型，支持的语言列表<li> zh : 中文 </li> <li> en : 英文 </li>
	Target *string `json:"Target,omitempty" name:"Target"`

	// pcm : 146   speex : 16779154   mp3 : 83886080
	AudioFormat *int64 `json:"AudioFormat,omitempty" name:"AudioFormat"`

	// 语音分片的序号，从0开始
	Seq *int64 `json:"Seq,omitempty" name:"Seq"`

	// 是否最后一片语音分片，0-否，1-是
	IsEnd *int64 `json:"IsEnd,omitempty" name:"IsEnd"`

	// 语音分片内容进行 Base64 编码后的字符串。音频内容需包含有效并可识别的文本信息。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 识别模式，该参数已废弃
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 该参数已废弃
	TransType *int64 `json:"TransType,omitempty" name:"TransType"`
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
	SessionUuid *string `json:"SessionUuid,omitempty" name:"SessionUuid"`

	// 语音识别状态 1-进行中 0-完成
	RecognizeStatus *int64 `json:"RecognizeStatus,omitempty" name:"RecognizeStatus"`

	// 识别出的原文
	SourceText *string `json:"SourceText,omitempty" name:"SourceText"`

	// 翻译出的译文
	TargetText *string `json:"TargetText,omitempty" name:"TargetText"`

	// 第几个语音分片
	Seq *int64 `json:"Seq,omitempty" name:"Seq"`

	// 原语言
	Source *string `json:"Source,omitempty" name:"Source"`

	// 目标语言
	Target *string `json:"Target,omitempty" name:"Target"`

	// 当请求的Mode参数填写bvad是，启动VadSeq。此时Seq会被设置为后台vad（静音检测）后的新序号，而VadSeq代表客户端原始Seq值
	VadSeq *int64 `json:"VadSeq,omitempty" name:"VadSeq"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
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
	Source *string `json:"Source,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// 
	// <li> zh（简体中文）：en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）</li>
	// <li>zh-TW（繁体中文）：en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）</li>
	// <li>en（英语）：zh（中文）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）、hi（印地语）</li>
	// <li>ja（日语）：zh（中文）、en（英语）、ko（韩语）</li>
	// <li>ko（韩语）：zh（中文）、en（英语）、ja（日语）</li>
	// <li>fr（法语）：zh（中文）、en（英语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>es（西班牙语）：zh（中文）、en（英语）、fr（法语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>it（意大利语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>de（德语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>tr（土耳其语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>ru（俄语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、pt（葡萄牙语）</li>
	// <li>pt（葡萄牙语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）</li>
	// <li>vi（越南语）：zh（中文）、en（英语）</li>
	// <li>id（印尼语）：zh（中文）、en（英语）</li>
	// <li>th（泰语）：zh（中文）、en（英语）</li>
	// <li>ms（马来语）：zh（中文）、en（英语）</li>
	// <li>ar（阿拉伯语）：en（英语）</li>
	// <li>hi（印地语）：en（英语）</li>
	Target *string `json:"Target,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 待翻译的文本列表，批量接口可以以数组方式在一次请求中填写多个待翻译文本。文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败，请传入有效文本，html标记等非常规翻译文本可能会翻译失败。单次请求的文本长度总和需要低于2000字符。
	SourceTextList []*string `json:"SourceTextList,omitempty" name:"SourceTextList"`
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
	Source *string `json:"Source,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// 
	// <li> zh（简体中文）：en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）</li>
	// <li>zh-TW（繁体中文）：en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）</li>
	// <li>en（英语）：zh（中文）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）、hi（印地语）</li>
	// <li>ja（日语）：zh（中文）、en（英语）、ko（韩语）</li>
	// <li>ko（韩语）：zh（中文）、en（英语）、ja（日语）</li>
	// <li>fr（法语）：zh（中文）、en（英语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>es（西班牙语）：zh（中文）、en（英语）、fr（法语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>it（意大利语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>de（德语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>tr（土耳其语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>ru（俄语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、pt（葡萄牙语）</li>
	// <li>pt（葡萄牙语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）</li>
	// <li>vi（越南语）：zh（中文）、en（英语）</li>
	// <li>id（印尼语）：zh（中文）、en（英语）</li>
	// <li>th（泰语）：zh（中文）、en（英语）</li>
	// <li>ms（马来语）：zh（中文）、en（英语）</li>
	// <li>ar（阿拉伯语）：en（英语）</li>
	// <li>hi（印地语）：en（英语）</li>
	Target *string `json:"Target,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 待翻译的文本列表，批量接口可以以数组方式在一次请求中填写多个待翻译文本。文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败，请传入有效文本，html标记等非常规翻译文本可能会翻译失败。单次请求的文本长度总和需要低于2000字符。
	SourceTextList []*string `json:"SourceTextList,omitempty" name:"SourceTextList"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextTranslateBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextTranslateBatchResponseParams struct {
	// 源语言，详见入参Target
	Source *string `json:"Source,omitempty" name:"Source"`

	// 目标语言，详见入参Target
	Target *string `json:"Target,omitempty" name:"Target"`

	// 翻译后的文本列表
	TargetTextList []*string `json:"TargetTextList,omitempty" name:"TargetTextList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 待翻译的文本，文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败，请传入有效文本，html标记等非常规翻译文本可能会翻译失败。单次请求的文本长度需要低于2000字符。
	SourceText *string `json:"SourceText,omitempty" name:"SourceText"`

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
	Source *string `json:"Source,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// 
	// <li> zh（简体中文）：en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）</li>
	// <li>zh-TW（繁体中文）：en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）</li>
	// <li>en（英语）：zh（中文）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）、hi（印地语）</li>
	// <li>ja（日语）：zh（中文）、en（英语）、ko（韩语）</li>
	// <li>ko（韩语）：zh（中文）、en（英语）、ja（日语）</li>
	// <li>fr（法语）：zh（中文）、en（英语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>es（西班牙语）：zh（中文）、en（英语）、fr（法语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>it（意大利语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>de（德语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>tr（土耳其语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>ru（俄语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、pt（葡萄牙语）</li>
	// <li>pt（葡萄牙语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）</li>
	// <li>vi（越南语）：zh（中文）、en（英语）</li>
	// <li>id（印尼语）：zh（中文）、en（英语）</li>
	// <li>th（泰语）：zh（中文）、en（英语）</li>
	// <li>ms（马来语）：zh（中文）、en（英语）</li>
	// <li>ar（阿拉伯语）：en（英语）</li>
	// <li>hi（印地语）：en（英语）</li>
	Target *string `json:"Target,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 用来标记不希望被翻译的文本内容，如句子中的特殊符号、人名、地名等；每次请求只支持配置一个不被翻译的单词；仅支持配置人名、地名等名词，不要配置动词或短语，否则会影响翻译结果。
	UntranslatedText *string `json:"UntranslatedText,omitempty" name:"UntranslatedText"`
}

type TextTranslateRequest struct {
	*tchttp.BaseRequest
	
	// 待翻译的文本，文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败，请传入有效文本，html标记等非常规翻译文本可能会翻译失败。单次请求的文本长度需要低于2000字符。
	SourceText *string `json:"SourceText,omitempty" name:"SourceText"`

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
	Source *string `json:"Source,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// 
	// <li> zh（简体中文）：en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）</li>
	// <li>zh-TW（繁体中文）：en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）</li>
	// <li>en（英语）：zh（中文）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）、hi（印地语）</li>
	// <li>ja（日语）：zh（中文）、en（英语）、ko（韩语）</li>
	// <li>ko（韩语）：zh（中文）、en（英语）、ja（日语）</li>
	// <li>fr（法语）：zh（中文）、en（英语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>es（西班牙语）：zh（中文）、en（英语）、fr（法语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>it（意大利语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>de（德语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>tr（土耳其语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>ru（俄语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、pt（葡萄牙语）</li>
	// <li>pt（葡萄牙语）：zh（中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）</li>
	// <li>vi（越南语）：zh（中文）、en（英语）</li>
	// <li>id（印尼语）：zh（中文）、en（英语）</li>
	// <li>th（泰语）：zh（中文）、en（英语）</li>
	// <li>ms（马来语）：zh（中文）、en（英语）</li>
	// <li>ar（阿拉伯语）：en（英语）</li>
	// <li>hi（印地语）：en（英语）</li>
	Target *string `json:"Target,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 用来标记不希望被翻译的文本内容，如句子中的特殊符号、人名、地名等；每次请求只支持配置一个不被翻译的单词；仅支持配置人名、地名等名词，不要配置动词或短语，否则会影响翻译结果。
	UntranslatedText *string `json:"UntranslatedText,omitempty" name:"UntranslatedText"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextTranslateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextTranslateResponseParams struct {
	// 翻译后的文本
	TargetText *string `json:"TargetText,omitempty" name:"TargetText"`

	// 源语言，详见入参Target
	Source *string `json:"Source,omitempty" name:"Source"`

	// 目标语言，详见入参Target
	Target *string `json:"Target,omitempty" name:"Target"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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