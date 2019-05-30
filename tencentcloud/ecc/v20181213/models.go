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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Aspect struct {

	// 项目 名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 该项得分
	Score *float64 `json:"Score,omitempty" name:"Score"`
}

type CompostionContext struct {

	// 作文内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CorrectData struct {

	// 总得分
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 各项得分详情
	ScoreCat *ScoreCategory `json:"ScoreCat,omitempty" name:"ScoreCat"`

	// 综合评价
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 句子点评
	SentenceComments []*SentenceCom `json:"SentenceComments,omitempty" name:"SentenceComments" list`
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
	Outline *string `json:"Outline,omitempty" name:"Outline"`

	// 范文标题，可选参数，本接口可以依据提供的范文对作文进行评分。
	ModelSubject *string `json:"ModelSubject,omitempty" name:"ModelSubject"`

	// 范文内容，可选参数，同上，范文的正文部分。
	ModelContent *string `json:"ModelContent,omitempty" name:"ModelContent"`
}

func (r *ECCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ECCRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ECCResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 整体的批改结果
		Data *CorrectData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ECCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ECCResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EHOCRRequest struct {
	*tchttp.BaseRequest

	// 图片所在的url或base64编码后的图像数据，依据InputType而定
	Image *string `json:"Image,omitempty" name:"Image"`

	// 输出图片类型，0表示Image字段是图片所在的url，1表示Image字段是base64编码后的图像数据
	InputType *int64 `json:"InputType,omitempty" name:"InputType"`
}

func (r *EHOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EHOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EHOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别后的作文内容
		Data *CompostionContext `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EHOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EHOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ScoreCategory struct {

	// 词汇项
	Words *Aspect `json:"Words,omitempty" name:"Words"`

	// 句子项
	Sentences *Aspect `json:"Sentences,omitempty" name:"Sentences"`

	// 篇章结构
	Structure *Aspect `json:"Structure,omitempty" name:"Structure"`

	// 内容
	Content *Aspect `json:"Content,omitempty" name:"Content"`
}

type SentenceCom struct {

	// 点评内容
	Suggestions []*SentenceSuggest `json:"Suggestions,omitempty" name:"Suggestions" list`

	// 点评的句子信息
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
}
