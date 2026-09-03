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

	// <p>返回结果类型，不传默认为0。（<strong>仅标准版、尊享版支持该参数</strong>）</p><p>枚举值：</p><ul><li>0： 公开网页信源结果（自然结果）</li><li>1： 优质权威垂直信源结果（VR 卡）</li><li>2： 混合结果（自然结果+VR卡）</li></ul>
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>指定网址搜索/过滤。（<strong>仅标准版、尊享版、旗舰版支持该参数</strong>）</p><ul><li>指定网址搜索：需要查询某个特定网址的内容时，传入&quot;Site=qq.com&quot;，实现只搜索qq.com的结果；每次搜索仅支持指定一个域名。</li><li>指定网址过滤：需要排除某个特定网址的内容时，传入&quot;Site=exclude:qq.com|sohu.com&quot;，实现过滤qq.com和sohu.com的结果；每次搜索最多支持过滤五个域名。</li></ul><p>注意： 该参数与mode参数共同使用时，仅对公开网页信源结果（自然结果）生效，对优质权威垂直信源结果（VR卡）不生效。</p>
	Site *string `json:"Site,omitnil,omitempty" name:"Site"`

	// <p>控制返回结果条数，可取值：cnt=10/20/30/40/50。（<strong>仅尊享版、旗舰版支持该参数</strong>）</p><p>枚举值：</p><ul><li>10： 返回10条结果</li><li>20： 返回20条结果</li><li>30： 返回30条结果</li><li>40： 返回40条结果</li><li>50： 返回50条结果</li></ul>
	Cnt *uint64 `json:"Cnt,omitnil,omitempty" name:"Cnt"`

	// <p>垂直领域搜索。（<strong>仅尊享版、旗舰版支持该参数</strong>）</p><p>枚举值：</p><ul><li>gov： 政府</li><li>news： 新闻</li><li>acad： 学术</li><li>finance： 财经</li></ul>
	Industry *string `json:"Industry,omitnil,omitempty" name:"Industry"`

	// <p>搜索时效范围，以下五种入参形态不支持混合使用。（<strong>仅标准版、尊享版、旗舰版支持该参数</strong>）</p><ul><li><p>d[N]：最近N天，N取值1-30整数，N值为空时默认N=1。</p></li><li><p>m[N]：最近N月，N取值1-12整数，N值为空时默认N=1。</p></li><li><p>y[N]：最近N年，N取值1-5整数，N值为空时默认N=1。</p></li><li><p>yyyy-mm-dd：指定某一日。（不得早于1970-01-01，不得晚于请求当天日期）</p></li><li><p>yyyy-mm-dd,yyyy-mm-dd：从日期A至日期B，包含AB。（日期B不得晚于日期A；二者均不得早于1970-01-01，不得晚于请求当天日期）</p></li></ul><p>示例：2026.6.15分别传参d2/m2/y2进行搜索，则搜索结果的时间范围分别为“2026.6.15、2026.6.14”/“2026.6、2026.5”/“2026、2025”，以此类推。</p><p>枚举值：</p><ul><li>d7： 最近七天</li><li>m3： 最近三月</li><li>y2： 最近两年</li><li>2026-08-20： 2026-08-20当天</li><li>2026-08-20,2026-08-30： 2026-08-20至2026-08-30</li></ul>
	Freshness *string `json:"Freshness,omitnil,omitempty" name:"Freshness"`

	// <p>返回附件子链信息（<strong>仅旗舰版支持该参数</strong>）</p><p>附件子链信息包括&quot;子链标题&quot;和&quot;子链URL&quot;，单个doc最多返回10条子链信息。</p><ul><li>true：返回</li><li>false：不返回</li><li>未传参时默认不返回</li></ul>
	Deeplinks *bool `json:"Deeplinks,omitnil,omitempty" name:"Deeplinks"`
}

type SearchProRequest struct {
	*tchttp.BaseRequest
	
	// <p>搜索词</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>返回结果类型，不传默认为0。（<strong>仅标准版、尊享版支持该参数</strong>）</p><p>枚举值：</p><ul><li>0： 公开网页信源结果（自然结果）</li><li>1： 优质权威垂直信源结果（VR 卡）</li><li>2： 混合结果（自然结果+VR卡）</li></ul>
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>指定网址搜索/过滤。（<strong>仅标准版、尊享版、旗舰版支持该参数</strong>）</p><ul><li>指定网址搜索：需要查询某个特定网址的内容时，传入&quot;Site=qq.com&quot;，实现只搜索qq.com的结果；每次搜索仅支持指定一个域名。</li><li>指定网址过滤：需要排除某个特定网址的内容时，传入&quot;Site=exclude:qq.com|sohu.com&quot;，实现过滤qq.com和sohu.com的结果；每次搜索最多支持过滤五个域名。</li></ul><p>注意： 该参数与mode参数共同使用时，仅对公开网页信源结果（自然结果）生效，对优质权威垂直信源结果（VR卡）不生效。</p>
	Site *string `json:"Site,omitnil,omitempty" name:"Site"`

	// <p>控制返回结果条数，可取值：cnt=10/20/30/40/50。（<strong>仅尊享版、旗舰版支持该参数</strong>）</p><p>枚举值：</p><ul><li>10： 返回10条结果</li><li>20： 返回20条结果</li><li>30： 返回30条结果</li><li>40： 返回40条结果</li><li>50： 返回50条结果</li></ul>
	Cnt *uint64 `json:"Cnt,omitnil,omitempty" name:"Cnt"`

	// <p>垂直领域搜索。（<strong>仅尊享版、旗舰版支持该参数</strong>）</p><p>枚举值：</p><ul><li>gov： 政府</li><li>news： 新闻</li><li>acad： 学术</li><li>finance： 财经</li></ul>
	Industry *string `json:"Industry,omitnil,omitempty" name:"Industry"`

	// <p>搜索时效范围，以下五种入参形态不支持混合使用。（<strong>仅标准版、尊享版、旗舰版支持该参数</strong>）</p><ul><li><p>d[N]：最近N天，N取值1-30整数，N值为空时默认N=1。</p></li><li><p>m[N]：最近N月，N取值1-12整数，N值为空时默认N=1。</p></li><li><p>y[N]：最近N年，N取值1-5整数，N值为空时默认N=1。</p></li><li><p>yyyy-mm-dd：指定某一日。（不得早于1970-01-01，不得晚于请求当天日期）</p></li><li><p>yyyy-mm-dd,yyyy-mm-dd：从日期A至日期B，包含AB。（日期B不得晚于日期A；二者均不得早于1970-01-01，不得晚于请求当天日期）</p></li></ul><p>示例：2026.6.15分别传参d2/m2/y2进行搜索，则搜索结果的时间范围分别为“2026.6.15、2026.6.14”/“2026.6、2026.5”/“2026、2025”，以此类推。</p><p>枚举值：</p><ul><li>d7： 最近七天</li><li>m3： 最近三月</li><li>y2： 最近两年</li><li>2026-08-20： 2026-08-20当天</li><li>2026-08-20,2026-08-30： 2026-08-20至2026-08-30</li></ul>
	Freshness *string `json:"Freshness,omitnil,omitempty" name:"Freshness"`

	// <p>返回附件子链信息（<strong>仅旗舰版支持该参数</strong>）</p><p>附件子链信息包括&quot;子链标题&quot;和&quot;子链URL&quot;，单个doc最多返回10条子链信息。</p><ul><li>true：返回</li><li>false：不返回</li><li>未传参时默认不返回</li></ul>
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

	// <p>搜索结果页面详情，格式为json字符串。</p><ul><li><p>title：结果标题</p></li><li><p>date：内容发布时间</p></li><li><p>url：内容发布源url</p></li><li><p>passage：标准摘要</p></li><li><p>content：动态摘要（<strong>仅尊享版、旗舰版返回该字段</strong>）</p></li><li><p>site：网站名称，部分不知名站点结果可能为空</p></li><li><p>score：相关性得分，取值0～1，越靠近1表示越相关</p></li><li><p>authority_level：权威度得分，取值0～5，数值越大表示越权威（<strong>仅旗舰版返回该字段</strong>）</p></li><li><p>pics：图片列表，单个doc返回0～10条（<strong>仅标准版、尊享版、旗舰版返回该字段</strong>）</p><ul><li>caption：图片描述</li><li>origin_url：源图url地</li></ul></li><li><p>favicon：网站图标链接，部分不知名站点结果可能为空</p></li><li><p>deeplinks：附件子链信息，单个doc最多返回10条子链信息。（<strong>仅旗舰版返回该字段，通过Deeplinks入参控制</strong>）</p><ul><li>title：子链标题</li><li>url：子链地址</li></ul></li></ul>
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