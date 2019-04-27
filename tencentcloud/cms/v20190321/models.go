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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AudioModerationRequest struct {
	*tchttp.BaseRequest

	// 回调url
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 音频内容的base64
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 音频文件的MD5值
	FileMD5 *string `json:"FileMD5,omitempty" name:"FileMD5"`

	// 音频内容Url ，其中FileUrl和FileContent二选一
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
}

func (r *AudioModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AudioModerationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AudioModerationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务返回码 
	// 60001：成功请求回调任务
		BusinessCode *int64 `json:"BusinessCode,omitempty" name:"BusinessCode"`

		// 识别返回结果
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AudioModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AudioModerationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModerationOverviewRequest struct {
	*tchttp.BaseRequest

	// 日期，如2019-01-01， 查询该日期的概览数据
	Date *string `json:"Date,omitempty" name:"Date"`

	// 服务类型数组，可以动态配置，Text:文本，Image:图片，Audio:音频，Video:视频, 使用"ALL"表示所有类型, 不区分大小写，如 ["Text", "Image"]查询文本和图片服务的数据，["all"]查询所有服务的数据。
	ServiceTypes []*string `json:"ServiceTypes,omitempty" name:"ServiceTypes" list`

	// 渠道号数组，1:直播 2:点播 3:IM 4:GME，统计指定渠道组合的汇总数据，如[1,2]表示获取直播和点播两个渠道的汇总数据，不填[]为所有渠道汇总数据
	Channels []*uint64 `json:"Channels,omitempty" name:"Channels" list`
}

func (r *DescribeModerationOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModerationOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModerationOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 概览数据集合
		Results []*OverviewRecord `json:"Results,omitempty" name:"Results" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModerationOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModerationOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageData struct {

	// 是否恶意 0：正常 1：可疑
	EvilFlag *int64 `json:"EvilFlag,omitempty" name:"EvilFlag"`

	// 恶意类型
	// 100：正常 
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	// 21000：综合
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 图片违法详情
	IllegalDetect *ImageIllegalDetect `json:"IllegalDetect,omitempty" name:"IllegalDetect"`

	// 图片涉政详情
	PolityDetect *ImagePolityDetect `json:"PolityDetect,omitempty" name:"PolityDetect"`

	// 图片涉黄详情
	PornDetect *ImagePornDetect `json:"PornDetect,omitempty" name:"PornDetect"`

	// 图片相似度详情
	Similar *Similar `json:"Similar,omitempty" name:"Similar"`

	// 图片暴恐详情
	TerrorDetect *ImageTerrorDetect `json:"TerrorDetect,omitempty" name:"TerrorDetect"`
}

type ImageIllegalDetect struct {

	// 恶意类型
	// 100：正常 
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	// 21000：综合
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords" list`

	// 违法标签：返回违法特征中文描述，如赌桌，枪支
	Labels []*string `json:"Labels,omitempty" name:"Labels" list`

	// 违法分：分值范围 0-100，分数越高违法倾向越明显
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type ImageModerationRequest struct {
	*tchttp.BaseRequest

	// 文件内容 Base64,与FileUrl必须二填一
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 文件MD5值
	FileMD5 *string `json:"FileMD5,omitempty" name:"FileMD5"`

	// 文件地址
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
}

func (r *ImageModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageModerationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageModerationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别结果
		Data *ImageData `json:"Data,omitempty" name:"Data"`

		// 业务返回码
		BusinessCode *int64 `json:"BusinessCode,omitempty" name:"BusinessCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageModerationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImagePolityDetect struct {

	// 恶意类型
	// 100：正常 
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	// 21000：综合
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定  0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 命中的人脸名称
	FaceNames []*string `json:"FaceNames,omitempty" name:"FaceNames" list`

	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords" list`

	// 政治（人脸）分：分值范围 0-100，分数越高可疑程度越高
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type ImagePornDetect struct {

	// 恶意类型
	// 100：正常 
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	// 21000：综合
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords" list`

	// 色情标签：色情特征中文描述
	Labels []*string `json:"Labels,omitempty" name:"Labels" list`

	// 色情分：分值范围 0-100，分数越高色情倾向越明显
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type ImageTerrorDetect struct {

	// 恶意类型
	// 100：正常 
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	// 21000：综合
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords" list`

	// 暴恐标签：返回暴恐特征中文描述
	Labels []*string `json:"Labels,omitempty" name:"Labels" list`

	// 暴恐分：分值范围0--100，分数越高暴恐倾向越明显
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type OverviewRecord struct {

	// 调用恶意量
	EvilCount *uint64 `json:"EvilCount,omitempty" name:"EvilCount"`

	// Text表示文本，Image表示图片，Audio表示音频，Video表示视频
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 调用总量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 恶意量同比增长率
	Yoy *string `json:"Yoy,omitempty" name:"Yoy"`
}

type Similar struct {

	// 恶意类型
	// 100：正常 
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	// 21000：综合
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：未匹配到 1：恶意 2：白样本
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 返回的种子url
	SeedUrl *string `json:"SeedUrl,omitempty" name:"SeedUrl"`
}

type TextData struct {

	// 是否恶意 0：正常 1：可疑
	EvilFlag *int64 `json:"EvilFlag,omitempty" name:"EvilFlag"`

	// 恶意类型
	// 100：正常
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	// 21000：综合
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 命中的关键词
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords" list`
}

type TextModerationRequest struct {
	*tchttp.BaseRequest

	// 文本内容Base64编码
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *TextModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TextModerationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TextModerationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别结果
		Data *TextData `json:"Data,omitempty" name:"Data"`

		// 业务返回码
		BusinessCode *int64 `json:"BusinessCode,omitempty" name:"BusinessCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TextModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TextModerationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VideoModerationRequest struct {
	*tchttp.BaseRequest

	// 回调Url
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 视频文件MD5
	FileMD5 *string `json:"FileMD5,omitempty" name:"FileMD5"`

	// 视频内容base64
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 视频内容Url,其中FileUrl与FileContent二选一
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
}

func (r *VideoModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VideoModerationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VideoModerationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务返回码
	// 60001：成功请求回调任务
		BusinessCode *int64 `json:"BusinessCode,omitempty" name:"BusinessCode"`

		// 识别返回结果
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VideoModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VideoModerationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
