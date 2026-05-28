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
	// <p>搜索词</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>返回结果类型，0-自然检索结果(默认)，1-多模态VR结果，2-混合结果（多模态VR结果+自然检索结果）</p>
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>指定域名站内搜索（用于过滤自然检索结果）<br>注意： mode=1模式下，参数无效；mode=0模式下，对所有结果生效；mode=2模式下，对输出的自然结果生效</p>
	Site *string `json:"Site,omitnil,omitempty" name:"Site"`

	// <p>起始时间（用于过滤自然检索结果），精确到秒时间戳格式<br>注意： mode=1模式下，参数无效；mode=0模式下，对所有结果生效；mode=2模式下，对输出的自然结果生效</p>
	FromTime *int64 `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// <p>结束时间（用于过滤自然检索结果），精确到秒时间戳格式<br>注意：mode=1模式下，参数无效；mode=0模式下，对所有结果生效；mode=2模式下，对输出的自然结果生效</p>
	ToTime *int64 `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// <p>cnt=10/20/30/40/50，最多可支持返回50条搜索结果，<strong>仅限尊享版使用</strong></p>
	Cnt *uint64 `json:"Cnt,omitnil,omitempty" name:"Cnt"`

	// <p>Industry=gov/news/acad/finance，对应党政机关、权威媒体、学术（英文）、金融，<strong>仅限尊享版使用</strong></p>
	Industry *string `json:"Industry,omitnil,omitempty" name:"Industry"`
}

type SearchProRequest struct {
	*tchttp.BaseRequest
	
	// <p>搜索词</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>返回结果类型，0-自然检索结果(默认)，1-多模态VR结果，2-混合结果（多模态VR结果+自然检索结果）</p>
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>指定域名站内搜索（用于过滤自然检索结果）<br>注意： mode=1模式下，参数无效；mode=0模式下，对所有结果生效；mode=2模式下，对输出的自然结果生效</p>
	Site *string `json:"Site,omitnil,omitempty" name:"Site"`

	// <p>起始时间（用于过滤自然检索结果），精确到秒时间戳格式<br>注意： mode=1模式下，参数无效；mode=0模式下，对所有结果生效；mode=2模式下，对输出的自然结果生效</p>
	FromTime *int64 `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// <p>结束时间（用于过滤自然检索结果），精确到秒时间戳格式<br>注意：mode=1模式下，参数无效；mode=0模式下，对所有结果生效；mode=2模式下，对输出的自然结果生效</p>
	ToTime *int64 `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// <p>cnt=10/20/30/40/50，最多可支持返回50条搜索结果，<strong>仅限尊享版使用</strong></p>
	Cnt *uint64 `json:"Cnt,omitnil,omitempty" name:"Cnt"`

	// <p>Industry=gov/news/acad/finance，对应党政机关、权威媒体、学术（英文）、金融，<strong>仅限尊享版使用</strong></p>
	Industry *string `json:"Industry,omitnil,omitempty" name:"Industry"`
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
	delete(f, "Cnt")
	delete(f, "Industry")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchProResponseParams struct {
	// <p>原始查询语</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>搜索结果页面详情，格式为json字符串。<br>title：结果标题<br>date：内容发布时间<br>url：内容发布源url<br>passage：标准摘要<br>content：动态摘要 （尊享版字段）<br>site：网站名称，部分不知名站点结果可能为空<br>score：相关性得分，取值0～1，越靠近1表示越相关<br>images：图片列表<br>favicon：网站图标链接，部分不知名站点结果可能为空</p>
	Pages []*string `json:"Pages,omitnil,omitempty" name:"Pages"`

	// <p>用户版本：standard/premium/lite/flagship</p>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// <p>提示信息</p>
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