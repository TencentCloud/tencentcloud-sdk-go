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

package v20180709

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeResourceListRequest struct {
	*tchttp.BaseRequest

	// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包；bgpip表示高防IP；net表示高防IP专业版；shield表示棋牌）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 地域码搜索，可选，当不指定地域时空数组，当指定地域时，填地域码。例如：["gz", "sh"]
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList" list`

	// 线路搜索，可选，只有当获取高防IP资源列表是可以选填，取值为[1（BGP线路），2（南京电信），3（南京联通），99（第三方合作线路）]，当获取其他产品时请填空数组；
	Line []*uint64 `json:"Line,omitempty" name:"Line" list`

	// 资源ID搜索，可选，当不为空数组时表示获取指定资源的资源列表；
	IdList []*string `json:"IdList,omitempty" name:"IdList" list`

	// 资源名称搜索，可选，当不为空字符串时表示按名称搜索资源；
	Name *string `json:"Name,omitempty" name:"Name"`

	// IP搜索列表，可选，当不为空时表示安装IP搜索资源；
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`

	// 资源状态搜索列表，可选，取值为[0（运行中）, 1（清洗中）, 2（封堵中）]，当填空数组时不进行状态搜索；
	Status []*uint64 `json:"Status,omitempty" name:"Status" list`

	// 即将到期搜索；可选，取值为[0（不搜索），1（搜索即将到期的资源）]
	Expire *uint64 `json:"Expire,omitempty" name:"Expire"`

	// 排序字段，可选
	OderBy []*OrderBy `json:"OderBy,omitempty" name:"OderBy" list`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 高防IP专业版资源的CNAME，可选，只对高防IP专业版资源列表有效；
	CName *string `json:"CName,omitempty" name:"CName"`

	// 高防IP专业版资源的域名，可选，只对高防IP专业版资源列表有效；
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeResourceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 资源记录列表
		ServicePacks []*KeyValueRecord `json:"ServicePacks,omitempty" name:"ServicePacks" list`

		// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包；bgpip表示高防IP；net表示高防IP专业版；shield表示棋牌）
		Business *string `json:"Business,omitempty" name:"Business"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeResourceListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type KeyValue struct {

	// 字段名称
	Key *string `json:"Key,omitempty" name:"Key"`

	// 字段取值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type KeyValueRecord struct {

	// 一条记录的Key-Value数组
	Record []*KeyValue `json:"Record,omitempty" name:"Record" list`
}

type OrderBy struct {

	// 排序字段名称，取值[
	// bandwidth（带宽），
	// overloadCount（超峰值次数）
	// ]
	Field *string `json:"Field,omitempty" name:"Field"`

	// 升降序，取值为[asc（升序），（升序），desc（降序）， DESC（降序）]
	Order *string `json:"Order,omitempty" name:"Order"`
}
