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

package v20200324

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BodyDetectResult struct {

	// 检测出的人体置信度。 
	// 误识率百分之十对应的阈值是0.14；误识率百分之五对应的阈值是0.32；误识率百分之二对应的阈值是0.62；误识率百分之一对应的阈值是0.81。 
	// 通常情况建议使用阈值0.32，可适用大多数情况。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 图中检测出来的人体框
	BodyRect *BodyRect `json:"BodyRect,omitempty" name:"BodyRect"`
}

type BodyRect struct {

	// 人体框左上角横坐标。
	X *uint64 `json:"X,omitempty" name:"X"`

	// 人体框左上角纵坐标。
	Y *uint64 `json:"Y,omitempty" name:"Y"`

	// 人体宽度。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 人体高度。
	Height *uint64 `json:"Height,omitempty" name:"Height"`
}

type Candidate struct {

	// 人员ID。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人体轨迹ID。
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 候选者的匹配得分。 
	// 十万人体库下，误识率百分之五对应的分数为70分；误识率百分之二对应的分数为80分；误识率百分之一对应的分数为90分。
	//  
	// 二十万人体库下，误识率百分之五对应的分数为80分；误识率百分之二对应的分数为90分；误识率百分之一对应的分数为95分。
	//  
	// 通常情况建议使用分数80分（保召回）。若希望获得较高精度，建议使用分数90分（保准确）。
	Score *float64 `json:"Score,omitempty" name:"Score"`
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// 人体库名称，[1,60]个字符，可修改，不可重复。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 人体库 ID，不可修改，不可重复。支持英文、数字、-%@#&_，长度限制64B。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人体库信息备注，[0，40]个字符。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 人体识别所用的算法模型版本。 
	// 目前入参仅支持 “1.0”1个输入。 默认为"1.0"。  
	// 不同算法模型版本对应的人体识别算法不同，新版本的整体效果会优于旧版本，后续我们将推出更新版本。
	BodyModelVersion *string `json:"BodyModelVersion,omitempty" name:"BodyModelVersion"`
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

	// 人员ID，单个腾讯云账号下不可修改，不可重复。 
	// 支持英文、数字、-%@#&_，，长度限制64B。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人体轨迹信息。
	Trace *Trace `json:"Trace,omitempty" name:"Trace"`
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

		// 人员轨迹唯一标识。
		TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

		// 人体识别所用的算法模型版本。
		BodyModelVersion *string `json:"BodyModelVersion,omitempty" name:"BodyModelVersion"`

		// 输入的人体轨迹图片中的合法性校验结果。
	// 只有为0时结果才有意义。
	// -1001: 输入图片不合法。-1002: 输入图片不能构成轨迹。
		InputRetCode *int64 `json:"InputRetCode,omitempty" name:"InputRetCode"`

		// 输入的人体轨迹图片中的合法性校验结果详情。 
	// -1101:图片无效，-1102:url不合法。-1103:图片过大。-1104:图片下载失败。-1105:图片解码失败。-1109:图片分辨率过高。-2023:轨迹中有非同人图片。-2024: 轨迹提取失败。-2025: 人体检测失败。
	// RetCode 的顺序和入参中Images 或 Urls 的顺序一致。
		InputRetCodeDetails []*int64 `json:"InputRetCodeDetails,omitempty" name:"InputRetCodeDetails" list`

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

type CreateTraceRequest struct {
	*tchttp.BaseRequest

	// 人员ID。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人体轨迹信息。
	Trace *Trace `json:"Trace,omitempty" name:"Trace"`
}

func (r *CreateTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTraceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTraceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人员轨迹唯一标识。
		TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

		// 人体识别所用的算法模型版本。
		BodyModelVersion *string `json:"BodyModelVersion,omitempty" name:"BodyModelVersion"`

		// 输入的人体轨迹图片中的合法性校验结果。
	// 只有为0时结果才有意义。
	// -1001: 输入图片不合法。-1002: 输入图片不能构成轨迹。
		InputRetCode *int64 `json:"InputRetCode,omitempty" name:"InputRetCode"`

		// 输入的人体轨迹图片中的合法性校验结果详情。 
	// -1101:图片无效，-1102:url不合法。-1103:图片过大。-1104:图片下载失败。-1105:图片解码失败。-1109:图片分辨率过高。-2023:轨迹中有非同人图片。-2024: 轨迹提取失败。-2025: 人体检测失败。
		InputRetCodeDetails []*int64 `json:"InputRetCodeDetails,omitempty" name:"InputRetCodeDetails" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTraceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// 人体库ID。
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

type DeletePersonRequest struct {
	*tchttp.BaseRequest

	// 人员ID。
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

type DetectBodyRequest struct {
	*tchttp.BaseRequest

	// 人体图片 Base64 数据。
	// 图片 base64 编码后大小不可超过5M。
	// 图片分辨率不得超过 2048*2048。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 人体图片 Url 。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片 base64 编码后大小不可超过5M。 
	// 图片分辨率不得超过 2048*2048。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 最多检测的人体数目，默认值为1（仅检测图片中面积最大的那个人体）； 最大值10 ，检测图片中面积最大的10个人体。
	MaxBodyNum *uint64 `json:"MaxBodyNum,omitempty" name:"MaxBodyNum"`
}

func (r *DetectBodyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectBodyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectBodyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图中检测出来的人体框。
		BodyDetectResults []*BodyDetectResult `json:"BodyDetectResults,omitempty" name:"BodyDetectResults" list`

		// 人体识别所用的算法模型版本。
		BodyModelVersion *string `json:"BodyModelVersion,omitempty" name:"BodyModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectBodyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectBodyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupListRequest struct {
	*tchttp.BaseRequest

	// 起始序号，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000。
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

		// 返回的人体库信息。
		GroupInfos []*GroupInfo `json:"GroupInfos,omitempty" name:"GroupInfos" list`

		// 人体库总数量。
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

type GetPersonListRequest struct {
	*tchttp.BaseRequest

	// 人体库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 起始序号，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000。
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

		// 返回的人员信息。
		PersonInfos []*PersonInfo `json:"PersonInfos,omitempty" name:"PersonInfos" list`

		// 该人体库的人员数量。
		PersonNum *uint64 `json:"PersonNum,omitempty" name:"PersonNum"`

		// 人体识别所用的算法模型版本。
		BodyModelVersion *string `json:"BodyModelVersion,omitempty" name:"BodyModelVersion"`

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

type GroupInfo struct {

	// 人体库名称。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 人体库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人体库信息备注。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 人体识别所用的算法模型版本。
	BodyModelVersion *string `json:"BodyModelVersion,omitempty" name:"BodyModelVersion"`

	// Group的创建时间和日期 CreationTimestamp。CreationTimestamp 的值是自 Unix 纪元时间到Group创建时间的毫秒数。  
	// Unix 纪元时间是 1970 年 1 月 1 日星期四，协调世界时 (UTC) 。
	CreationTimestamp *uint64 `json:"CreationTimestamp,omitempty" name:"CreationTimestamp"`
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest

	// 人体库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人体库名称。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 人体库信息备注。
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

type ModifyPersonInfoRequest struct {
	*tchttp.BaseRequest

	// 人员ID。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人员名称。
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`
}

func (r *ModifyPersonInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPersonInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PersonInfo struct {

	// 人员名称。
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 人员ID。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 包含的人体轨迹图片信息列表。
	TraceInfos []*TraceInfo `json:"TraceInfos,omitempty" name:"TraceInfos" list`
}

type SearchTraceRequest struct {
	*tchttp.BaseRequest

	// 希望搜索的人体库ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 人体轨迹信息。
	Trace *Trace `json:"Trace,omitempty" name:"Trace"`

	// 单张被识别的人体轨迹返回的最相似人员数量。
	// 默认值为5，最大值为100。
	//  例，设MaxPersonNum为8，则返回Top8相似的人员信息。 值越大，需要处理的时间越长。建议不要超过10。
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitempty" name:"MaxPersonNum"`

	// 出参Score中，只有超过TraceMatchThreshold值的结果才会返回。
	// 默认为0。范围[0, 100.0]。
	TraceMatchThreshold *float64 `json:"TraceMatchThreshold,omitempty" name:"TraceMatchThreshold"`
}

func (r *SearchTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchTraceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchTraceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别出的最相似候选人。
		Candidates []*Candidate `json:"Candidates,omitempty" name:"Candidates" list`

		// 输入的人体轨迹图片中的合法性校验结果。
	// 只有为0时结果才有意义。
	// -1001: 输入图片不合法。-1002: 输入图片不能构成轨迹。
		InputRetCode *int64 `json:"InputRetCode,omitempty" name:"InputRetCode"`

		// 输入的人体轨迹图片中的合法性校验结果详情。 
	// -1101:图片无效，-1102:url不合法。-1103:图片过大。-1104:图片下载失败。-1105:图片解码失败。-1109:图片分辨率过高。-2023:轨迹中有非同人图片。-2024: 轨迹提取失败。-2025: 人体检测失败。
		InputRetCodeDetails []*int64 `json:"InputRetCodeDetails,omitempty" name:"InputRetCodeDetails" list`

		// 人体识别所用的算法模型版本。
		BodyModelVersion *string `json:"BodyModelVersion,omitempty" name:"BodyModelVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchTraceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SegmentPortraitPicRequest struct {
	*tchttp.BaseRequest

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 图片分辨率须小于2000*2000。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片分辨率须小于2000*2000 ，图片 base64 编码后大小不可超过5M。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *SegmentPortraitPicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SegmentPortraitPicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SegmentPortraitPicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 处理后的图片 base64 数据，透明背景图
		ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

		// 一个通过 Base64 编码的文件，解码后文件由 Float 型浮点数组成。这些浮点数代表原图从左上角开始的每一行的每一个像素点，每一个浮点数的值是原图相应像素点位于人体轮廓内的置信度（0-1）转化的灰度值（0-255）
		ResultMask *string `json:"ResultMask,omitempty" name:"ResultMask"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SegmentPortraitPicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SegmentPortraitPicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Trace struct {

	// 人体轨迹图片 Base64 数组。 
	// 数组长度最小为1最大为5。 
	// 单个图片 base64 编码后大小不可超过2M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Images []*string `json:"Images,omitempty" name:"Images" list`

	// 人体轨迹图片 Url 数组。 
	// 数组长度最小为1最大为5。 
	// 单个图片 base64 编码后大小不可超过2M。 
	// Urls、Images必须提供一个，如果都提供，只使用 Urls。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

	// 若输入的Images 和 Urls 是已经裁剪后的人体小图，则可以忽略本参数。 
	// 若否，或图片中包含多个人体，则需要通过本参数来指定图片中的人体框。 
	// 顺序对应 Images 或 Urls 中的顺序。  
	// 当不输入本参数时，我们将认为输入图片已是经过裁剪后的人体小图，不会进行人体检测而直接进行特征提取处理。
	BodyRects []*BodyRect `json:"BodyRects,omitempty" name:"BodyRects" list`
}

type TraceInfo struct {

	// 人体轨迹ID。
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 包含的人体轨迹图片Id列表。
	BodyIds []*string `json:"BodyIds,omitempty" name:"BodyIds" list`
}
