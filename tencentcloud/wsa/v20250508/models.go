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

	// <p>搜索实效范围（仅旗舰版参数）</p><ul><li>d[N]：最近N天，N取值1-30整数。</li><li>m[N]：最近N月，N取值1-12整数。</li><li>y[N]：最近N年，N取值1-5整数。</li></ul><p>示例说明：</p><ul><li>d1/m1/y1：当天/当月/当年。<br>例如，2026.6.15分别传参d1/m1/y1进行搜索，则搜索结果的时间范围分别为“2026.6.15”/“2026.6”/“2026”，以此类推。</li><li>d/m/y：N值为空时，默认N=1，即等效入参d1/m1/y1。</li><li>未传参时，默认不生效。</li></ul><p>枚举值：</p><ul><li>d7： 最近七天</li><li>m3： 最近三月</li><li>y2： 最近两年</li><li>d： 当天</li><li>m： 当月</li><li>y： 当年</li></ul>
	Freshness *string `json:"Freshness,omitnil,omitempty" name:"Freshness"`

	// <p>返回子链信息（仅旗舰版参数）</p><p>子链信息包括&quot;子链标题&quot;和&quot;子链URL&quot;，单个doc最多返回10条子链信息。</p><ul><li>true：返回</li><li>false：不返回</li><li>未传参时默认不返回</li></ul>
	Deeplinks *bool `json:"Deeplinks,omitnil,omitempty" name:"Deeplinks"`
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

	// <p>搜索实效范围（仅旗舰版参数）</p><ul><li>d[N]：最近N天，N取值1-30整数。</li><li>m[N]：最近N月，N取值1-12整数。</li><li>y[N]：最近N年，N取值1-5整数。</li></ul><p>示例说明：</p><ul><li>d1/m1/y1：当天/当月/当年。<br>例如，2026.6.15分别传参d1/m1/y1进行搜索，则搜索结果的时间范围分别为“2026.6.15”/“2026.6”/“2026”，以此类推。</li><li>d/m/y：N值为空时，默认N=1，即等效入参d1/m1/y1。</li><li>未传参时，默认不生效。</li></ul><p>枚举值：</p><ul><li>d7： 最近七天</li><li>m3： 最近三月</li><li>y2： 最近两年</li><li>d： 当天</li><li>m： 当月</li><li>y： 当年</li></ul>
	Freshness *string `json:"Freshness,omitnil,omitempty" name:"Freshness"`

	// <p>返回子链信息（仅旗舰版参数）</p><p>子链信息包括&quot;子链标题&quot;和&quot;子链URL&quot;，单个doc最多返回10条子链信息。</p><ul><li>true：返回</li><li>false：不返回</li><li>未传参时默认不返回</li></ul>
	Deeplinks *bool `json:"Deeplinks,omitnil,omitempty" name:"Deeplinks"`
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
	delete(f, "Freshness")
	delete(f, "Deeplinks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchProResponseParams struct {
	// <p>原始查询语</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>搜索结果页面详情，格式为json字符串。</p><ul><li><p>title：结果标题</p></li><li><p>date：内容发布时间</p></li><li><p>url：内容发布源url</p></li><li><p>passage：标准摘要</p></li><li><p>content：动态摘要（仅尊享版、旗舰版返回该字段）</p></li><li><p>site：网站名称，部分不知名站点结果可能为空</p></li><li><p>score：相关性得分，取值0～1，越靠近1表示越相关</p></li><li><p>images：图片列表（旗舰版无该出参）</p></li><li><p>pics：图片列表，单个doc返回0～10条（仅旗舰版参数）</p><ul><li>caption：图片描述</li><li>origin_url：源图url地</li></ul></li><li><p>favicon：网站图标链接，部分不知名站点结果可能为空</p></li><li><p>deeplinks：子链信息，单个doc最多返回10条子链信息。（仅旗舰版参数，通过Deeplinks入参控制）</p><ul><li>title：子链标题</li><li>url：子链地址</li></ul></li></ul>
	Pages []*string `json:"Pages,omitnil,omitempty" name:"Pages"`

	// <p>用户版本：standard/premium/lite/flagship（标准/尊享/轻量/旗舰）</p>
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