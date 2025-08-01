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

package v20250508

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type SearchProRequestParams struct {
	// 搜索词
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 返回结果类型，0-自然检索结果(默认)，1-垂类VR结果，2-混合结果（垂类VR结果+自然检索结果）
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 指定域名站内搜索（用于过滤自然检索结果）  注意：  mode=1模式下，参数无效 mode=0模式下对所有结果生效 mode=2模式下对输出的自然结果生效
	Site *string `json:"Site,omitnil,omitempty" name:"Site"`

	// 起始时间（用于过滤自然检索结果），精确到秒时间戳格式  注意：  mode=1模式下，参数无效 mode=0模式下对所有结果生效 mode=2模式下对输出的自然结果生效
	FromTime *int64 `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 结束时间（用于过滤自然检索结果），精确到秒时间戳格式  注意：  mode=1模式下，参数无效 mode=0模式下对所有结果生效 mode=2模式下对输出的自然结果生效
	ToTime *int64 `json:"ToTime,omitnil,omitempty" name:"ToTime"`
}

type SearchProRequest struct {
	*tchttp.BaseRequest
	
	// 搜索词
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 返回结果类型，0-自然检索结果(默认)，1-垂类VR结果，2-混合结果（垂类VR结果+自然检索结果）
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 指定域名站内搜索（用于过滤自然检索结果）  注意：  mode=1模式下，参数无效 mode=0模式下对所有结果生效 mode=2模式下对输出的自然结果生效
	Site *string `json:"Site,omitnil,omitempty" name:"Site"`

	// 起始时间（用于过滤自然检索结果），精确到秒时间戳格式  注意：  mode=1模式下，参数无效 mode=0模式下对所有结果生效 mode=2模式下对输出的自然结果生效
	FromTime *int64 `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 结束时间（用于过滤自然检索结果），精确到秒时间戳格式  注意：  mode=1模式下，参数无效 mode=0模式下对所有结果生效 mode=2模式下对输出的自然结果生效
	ToTime *int64 `json:"ToTime,omitnil,omitempty" name:"ToTime"`
}

func (r *SearchProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	delete(f, "Mode")
	delete(f, "Site")
	delete(f, "FromTime")
	delete(f, "ToTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchProResponseParams struct {
	// 原始查询语
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 搜索结果页面
	Pages []*string `json:"Pages,omitnil,omitempty" name:"Pages"`

	// 提示信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchProResponse struct {
	*tchttp.BaseResponse
	Response *SearchProResponseParams `json:"Response"`
}

func (r *SearchProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}