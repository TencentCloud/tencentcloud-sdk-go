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

package v20181213

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Aspect struct {
	// 维度名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 维度得分
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 维度分数占比
	Percentage *float64 `json:"Percentage,omitempty" name:"Percentage"`
}

type CompostionContext struct {
	// 作文内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 批改结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorrectData *CorrectData `json:"CorrectData,omitempty" name:"CorrectData"`

	// 任务 id，用于查询接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 图像识别唯一标识，一次识别一个 SessionId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type CorrectData struct {
	// 总得分
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 各项得分详情
	ScoreCat *ScoreCategory `json:"ScoreCat,omitempty" name:"ScoreCat"`

	// 综合评价
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 句子点评
	SentenceComments []*SentenceCom `json:"SentenceComments,omitempty" name:"SentenceComments"`
}

// Predefined struct for user
type CorrectMultiImageRequestParams struct {
	// 图片的url链接或base64数据。每张图片数据作为数组的一个元素，数组个数与图片个数保持一致。存放类别依据InputType而定，url与base64编码不能混合使用。
	Image []*string `json:"Image,omitempty" name:"Image"`

	// 输出图片类型，0 表示 Image 字段是图片所在的 url，1 表示 Image 字段是 base64 编码后的图像数据。
	InputType *int64 `json:"InputType,omitempty" name:"InputType"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数。
	EccAppid *string `json:"EccAppid,omitempty" name:"EccAppid"`

	// 图像识别唯一标识，一次识别一个 SessionId，使用识别功能时 SessionId 可用于使用文本批改接口，此时按图像批改价格收费；如使用文本批改接口时没有传入 SessionId，则需要收取文本批改的费用。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 服务类型，0：“多图像识别”，只返回识别结果；1：“多图像批改”，同时返回识别结果与批改结果。默认为 0。
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// 作文题目，可选参数
	Title *string `json:"Title,omitempty" name:"Title"`

	// 年级标准， 默认以 cet4 为标准，取值与意义如下：elementary 小学，grade7 grade8 grade9分别对应初一，初二，初三。 grade10 grade11 grade12 分别对应高一，高二，高三，以及 cet4 和 cet6 分别表示 英语4级和6级。
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 作文提纲，可选参数，作文的写作要求。
	Requirement *string `json:"Requirement,omitempty" name:"Requirement"`

	// 范文标题，可选参数，本接口可以依据提供的范文对作文进行评分。
	ModelTitle *string `json:"ModelTitle,omitempty" name:"ModelTitle"`

	// 范文内容，可选参数，同上，范文的正文部分。
	ModelContent *string `json:"ModelContent,omitempty" name:"ModelContent"`

	// 异步模式标识，0：同步模式，1：异步模式。默认为同步模式
	IsAsync *int64 `json:"IsAsync,omitempty" name:"IsAsync"`
}

type CorrectMultiImageRequest struct {
	*tchttp.BaseRequest
	
	// 图片的url链接或base64数据。每张图片数据作为数组的一个元素，数组个数与图片个数保持一致。存放类别依据InputType而定，url与base64编码不能混合使用。
	Image []*string `json:"Image,omitempty" name:"Image"`

	// 输出图片类型，0 表示 Image 字段是图片所在的 url，1 表示 Image 字段是 base64 编码后的图像数据。
	InputType *int64 `json:"InputType,omitempty" name:"InputType"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数。
	EccAppid *string `json:"EccAppid,omitempty" name:"EccAppid"`

	// 图像识别唯一标识，一次识别一个 SessionId，使用识别功能时 SessionId 可用于使用文本批改接口，此时按图像批改价格收费；如使用文本批改接口时没有传入 SessionId，则需要收取文本批改的费用。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 服务类型，0：“多图像识别”，只返回识别结果；1：“多图像批改”，同时返回识别结果与批改结果。默认为 0。
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// 作文题目，可选参数
	Title *string `json:"Title,omitempty" name:"Title"`

	// 年级标准， 默认以 cet4 为标准，取值与意义如下：elementary 小学，grade7 grade8 grade9分别对应初一，初二，初三。 grade10 grade11 grade12 分别对应高一，高二，高三，以及 cet4 和 cet6 分别表示 英语4级和6级。
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 作文提纲，可选参数，作文的写作要求。
	Requirement *string `json:"Requirement,omitempty" name:"Requirement"`

	// 范文标题，可选参数，本接口可以依据提供的范文对作文进行评分。
	ModelTitle *string `json:"ModelTitle,omitempty" name:"ModelTitle"`

	// 范文内容，可选参数，同上，范文的正文部分。
	ModelContent *string `json:"ModelContent,omitempty" name:"ModelContent"`

	// 异步模式标识，0：同步模式，1：异步模式。默认为同步模式
	IsAsync *int64 `json:"IsAsync,omitempty" name:"IsAsync"`
}

func (r *CorrectMultiImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CorrectMultiImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "InputType")
	delete(f, "EccAppid")
	delete(f, "SessionId")
	delete(f, "ServerType")
	delete(f, "Title")
	delete(f, "Grade")
	delete(f, "Requirement")
	delete(f, "ModelTitle")
	delete(f, "ModelContent")
	delete(f, "IsAsync")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CorrectMultiImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CorrectMultiImageResponseParams struct {
	// 接口返回数据
	Data *CompostionContext `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CorrectMultiImageResponse struct {
	*tchttp.BaseResponse
	Response *CorrectMultiImageResponseParams `json:"Response"`
}

func (r *CorrectMultiImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CorrectMultiImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskRequestParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数（暂时无需传入）。
	EccAppid *string `json:"EccAppid,omitempty" name:"EccAppid"`
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数（暂时无需传入）。
	EccAppid *string `json:"EccAppid,omitempty" name:"EccAppid"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "EccAppid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResponseParams struct {
	// 作文识别文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 整体的批改结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorrectData *CorrectData `json:"CorrectData,omitempty" name:"CorrectData"`

	// 任务状态，“Progressing”: 处理中（此时无结果返回）、“Finished”: 处理完成
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResponseParams `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ECCRequestParams struct {
	// 作文文本，必填
	Content *string `json:"Content,omitempty" name:"Content"`

	// 作文题目，可选参数
	Title *string `json:"Title,omitempty" name:"Title"`

	// 年级标准， 默认以cet4为标准，取值与意义如下：elementary 小学，grade7 grade8 grade9分别对应初一，初二，初三。 grade10 grade11 grade12 分别对应高一，高二，高三，以及cet4和cet6 分别表示 英语4级和6级。
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 作文提纲，可选参数，作文的写作要求。
	Requirement *string `json:"Requirement,omitempty" name:"Requirement"`

	// 范文标题，可选参数，本接口可以依据提供的范文对作文进行评分。
	ModelTitle *string `json:"ModelTitle,omitempty" name:"ModelTitle"`

	// 范文内容，可选参数，同上，范文的正文部分。
	ModelContent *string `json:"ModelContent,omitempty" name:"ModelContent"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数（暂时无需传入）。
	EccAppid *string `json:"EccAppid,omitempty" name:"EccAppid"`

	// 异步模式标识，0：同步模式，1：异步模式，默认为同步模式
	IsAsync *int64 `json:"IsAsync,omitempty" name:"IsAsync"`

	// 图像识别唯一标识，一次识别一个 SessionId。当传入此前识别接口使用过的 SessionId，则本次批改按图像批改价格收费；如使用了识别接口且本次没有传入 SessionId，则需要加取文本批改的费用；如果直接使用文本批改接口，则只收取文本批改的费用
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type ECCRequest struct {
	*tchttp.BaseRequest
	
	// 作文文本，必填
	Content *string `json:"Content,omitempty" name:"Content"`

	// 作文题目，可选参数
	Title *string `json:"Title,omitempty" name:"Title"`

	// 年级标准， 默认以cet4为标准，取值与意义如下：elementary 小学，grade7 grade8 grade9分别对应初一，初二，初三。 grade10 grade11 grade12 分别对应高一，高二，高三，以及cet4和cet6 分别表示 英语4级和6级。
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 作文提纲，可选参数，作文的写作要求。
	Requirement *string `json:"Requirement,omitempty" name:"Requirement"`

	// 范文标题，可选参数，本接口可以依据提供的范文对作文进行评分。
	ModelTitle *string `json:"ModelTitle,omitempty" name:"ModelTitle"`

	// 范文内容，可选参数，同上，范文的正文部分。
	ModelContent *string `json:"ModelContent,omitempty" name:"ModelContent"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数（暂时无需传入）。
	EccAppid *string `json:"EccAppid,omitempty" name:"EccAppid"`

	// 异步模式标识，0：同步模式，1：异步模式，默认为同步模式
	IsAsync *int64 `json:"IsAsync,omitempty" name:"IsAsync"`

	// 图像识别唯一标识，一次识别一个 SessionId。当传入此前识别接口使用过的 SessionId，则本次批改按图像批改价格收费；如使用了识别接口且本次没有传入 SessionId，则需要加取文本批改的费用；如果直接使用文本批改接口，则只收取文本批改的费用
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

func (r *ECCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ECCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "Title")
	delete(f, "Grade")
	delete(f, "Requirement")
	delete(f, "ModelTitle")
	delete(f, "ModelContent")
	delete(f, "EccAppid")
	delete(f, "IsAsync")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ECCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ECCResponseParams struct {
	// 整体的批改结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CorrectData `json:"Data,omitempty" name:"Data"`

	// 任务 id，用于查询接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ECCResponse struct {
	*tchttp.BaseResponse
	Response *ECCResponseParams `json:"Response"`
}

func (r *ECCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ECCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EHOCRRequestParams struct {
	// 图片所在的url或base64编码后的图像数据，依据InputType而定
	Image *string `json:"Image,omitempty" name:"Image"`

	// 输出图片类型，0 表示 Image 字段是图片所在的 url，1 表示 Image 字段是 base64 编码后的图像数据
	InputType *int64 `json:"InputType,omitempty" name:"InputType"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数（暂时无需传入）。
	EccAppid *string `json:"EccAppid,omitempty" name:"EccAppid"`

	// 图像识别唯一标识，一次识别一个 SessionId，使用识别功能时 SessionId 可用于使用文本批改接口，此时按图像批改价格收费；如使用文本批改接口时没有传入 SessionId，则需要收取文本批改的费用
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 服务类型，0：“图像识别”，只返回识别结果，1：“图像批改”，同时返回识别结果与批改结果。默认为 0
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// 作文题目，可选参数
	Title *string `json:"Title,omitempty" name:"Title"`

	// 年级标准， 默认以 cet4 为标准，取值与意义如下：elementary 小学，grade7 grade8 grade9分别对应初一，初二，初三。 grade10 grade11 grade12 分别对应高一，高二，高三，以及 cet4 和 cet6 分别表示 英语4级和6级。
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 作文提纲，可选参数，作文的写作要求。
	Requirement *string `json:"Requirement,omitempty" name:"Requirement"`

	// 范文标题，可选参数，本接口可以依据提供的范文对作文进行评分。
	ModelTitle *string `json:"ModelTitle,omitempty" name:"ModelTitle"`

	// 范文内容，可选参数，同上，范文的正文部分。
	ModelContent *string `json:"ModelContent,omitempty" name:"ModelContent"`

	// 异步模式标识，0：同步模式，1：异步模式。默认为同步模式
	IsAsync *int64 `json:"IsAsync,omitempty" name:"IsAsync"`
}

type EHOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片所在的url或base64编码后的图像数据，依据InputType而定
	Image *string `json:"Image,omitempty" name:"Image"`

	// 输出图片类型，0 表示 Image 字段是图片所在的 url，1 表示 Image 字段是 base64 编码后的图像数据
	InputType *int64 `json:"InputType,omitempty" name:"InputType"`

	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数（暂时无需传入）。
	EccAppid *string `json:"EccAppid,omitempty" name:"EccAppid"`

	// 图像识别唯一标识，一次识别一个 SessionId，使用识别功能时 SessionId 可用于使用文本批改接口，此时按图像批改价格收费；如使用文本批改接口时没有传入 SessionId，则需要收取文本批改的费用
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 服务类型，0：“图像识别”，只返回识别结果，1：“图像批改”，同时返回识别结果与批改结果。默认为 0
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// 作文题目，可选参数
	Title *string `json:"Title,omitempty" name:"Title"`

	// 年级标准， 默认以 cet4 为标准，取值与意义如下：elementary 小学，grade7 grade8 grade9分别对应初一，初二，初三。 grade10 grade11 grade12 分别对应高一，高二，高三，以及 cet4 和 cet6 分别表示 英语4级和6级。
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 作文提纲，可选参数，作文的写作要求。
	Requirement *string `json:"Requirement,omitempty" name:"Requirement"`

	// 范文标题，可选参数，本接口可以依据提供的范文对作文进行评分。
	ModelTitle *string `json:"ModelTitle,omitempty" name:"ModelTitle"`

	// 范文内容，可选参数，同上，范文的正文部分。
	ModelContent *string `json:"ModelContent,omitempty" name:"ModelContent"`

	// 异步模式标识，0：同步模式，1：异步模式。默认为同步模式
	IsAsync *int64 `json:"IsAsync,omitempty" name:"IsAsync"`
}

func (r *EHOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EHOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "InputType")
	delete(f, "EccAppid")
	delete(f, "SessionId")
	delete(f, "ServerType")
	delete(f, "Title")
	delete(f, "Grade")
	delete(f, "Requirement")
	delete(f, "ModelTitle")
	delete(f, "ModelContent")
	delete(f, "IsAsync")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EHOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EHOCRResponseParams struct {
	// 接口返回数据
	Data *CompostionContext `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EHOCRResponse struct {
	*tchttp.BaseResponse
	Response *EHOCRResponseParams `json:"Response"`
}

func (r *EHOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EHOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrorCoordinate struct {
	// 维度单词坐标
	Coordinate []*int64 `json:"Coordinate,omitempty" name:"Coordinate"`
}

type ScoreCategory struct {
	// 词汇维度
	Words *Aspect `json:"Words,omitempty" name:"Words"`

	// 句子维度
	Sentences *Aspect `json:"Sentences,omitempty" name:"Sentences"`

	// 篇章结构维度
	Structure *Aspect `json:"Structure,omitempty" name:"Structure"`

	// 内容维度
	Content *Aspect `json:"Content,omitempty" name:"Content"`

	// 维度得分
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 维度分数占比
	Percentage *float64 `json:"Percentage,omitempty" name:"Percentage"`
}

type SentenceCom struct {
	// 句子错误纠正信息
	Suggestions []*SentenceSuggest `json:"Suggestions,omitempty" name:"Suggestions"`

	// 句子信息
	Sentence *SentenceItem `json:"Sentence,omitempty" name:"Sentence"`
}

type SentenceItem struct {
	// 英语句子
	Sentence *string `json:"Sentence,omitempty" name:"Sentence"`

	// 段落id
	ParaID *int64 `json:"ParaID,omitempty" name:"ParaID"`

	// 句子id
	SentenceID *int64 `json:"SentenceID,omitempty" name:"SentenceID"`
}

type SentenceSuggest struct {
	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 错误类型
	ErrorType *string `json:"ErrorType,omitempty" name:"ErrorType"`

	// 原始单词
	Origin *string `json:"Origin,omitempty" name:"Origin"`

	// 替换成 的单词
	Replace *string `json:"Replace,omitempty" name:"Replace"`

	// 提示信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 维度单词位置，在句子的第几个到第几个单词之间
	ErrorPosition []*int64 `json:"ErrorPosition,omitempty" name:"ErrorPosition"`

	// 维度单词坐标，错误单词在图片中的坐标，只有传图片时正常返回，传文字时返回[ ]
	ErrorCoordinates []*ErrorCoordinate `json:"ErrorCoordinates,omitempty" name:"ErrorCoordinates"`
}