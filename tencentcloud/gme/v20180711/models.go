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

package v20180711

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeFilterResultListRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 开始时间，格式为 年-月-日，如: 2018-07-11
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束时间，格式为 年-月-日，如: 2018-07-11
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFilterResultListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFilterResultListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFilterResultListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 过滤结果总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 当前分页过滤结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*VoiceFilterInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFilterResultListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFilterResultListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFilterResultRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 文件ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`
}

func (r *DescribeFilterResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFilterResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFilterResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 过滤结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *VoiceFilterInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFilterResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFilterResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VoiceFilter struct {

	// 过滤类型，1：政治，2：色情，3：涉毒，4：谩骂
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 过滤命中关键词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Word *string `json:"Word,omitempty" name:"Word"`
}

type VoiceFilterInfo struct {

	// 应用id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 文件id，表示文件唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 数据创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 过滤结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*VoiceFilter `json:"Data,omitempty" name:"Data" list`
}

type VoiceFilterRequest struct {
	*tchttp.BaseRequest

	// 应用ID，登录[控制台](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 文件ID，表示文件唯一ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件url，urlencode编码，FileUrl和FileContent二选一
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 文件内容，base64编码，FileUrl和FileContent二选一
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 用户ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
}

func (r *VoiceFilterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VoiceFilterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VoiceFilterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VoiceFilterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VoiceFilterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
