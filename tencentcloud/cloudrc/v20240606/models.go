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

package v20240606

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeResourceRequestParams struct {
	// <p>资源类型</p>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>地域编码</p>
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// <p>资源ID</p>
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>视图ID</p>
	ViewId *string `json:"ViewId,omitnil,omitempty" name:"ViewId"`
}

type DescribeResourceRequest struct {
	*tchttp.BaseRequest
	
	// <p>资源类型</p>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>地域编码</p>
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// <p>资源ID</p>
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>视图ID</p>
	ViewId *string `json:"ViewId,omitnil,omitempty" name:"ViewId"`
}

func (r *DescribeResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	delete(f, "RegionCode")
	delete(f, "ResourceId")
	delete(f, "ViewId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceResponseParams struct {
	// <p>资源ID</p>
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>资源别名</p>
	ResourceAlias *string `json:"ResourceAlias,omitnil,omitempty" name:"ResourceAlias"`

	// <p>uin</p>
	Uin *int64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>资源类型</p>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>地域编码</p>
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// <p>可用区编码</p>
	ZoneCode *string `json:"ZoneCode,omitnil,omitempty" name:"ZoneCode"`

	// <p>付费类型</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>资源创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>资源过期时间</p>
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// <p>内网IP</p>
	PrivateIpAddress []*string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// <p>外网IP</p>
	PublicIpAddress []*string `json:"PublicIpAddress,omitnil,omitempty" name:"PublicIpAddress"`

	// <p>资源属性</p>
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`

	// <p>标签信息</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceResponseParams `json:"Response"`
}

func (r *DescribeResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtendedFilter struct {
	// <p>过滤条件键</p><p>枚举值：</p><ul><li>ResourceType： 资源类型</li><li>ResourceId： 资源ID</li><li>ResourceAlias： 资源名称</li><li>PayMode： 计费模式</li><li>RegionCode： 地域编码</li><li>ZoneCode： 可用区编码</li><li>PublicIpAddress： 外网IP</li><li>PrivateIpAddress： 内网IP</li><li>VpcId： VPC ID</li><li>SubnetId： 子网ID</li><li>Tag： 标签</li></ul>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>过滤条件值</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// <p>匹配方式</p><p>枚举值：</p><ul><li>Equals： 等于</li><li>NotEquals： 不等于</li><li>Contains： 包含</li><li>NotContains： 不包含</li><li>Exists： 存在</li><li>NotExists： 不存在</li></ul>
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`
}

type ResourceSummary struct {
	// <p>资源ID</p>
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>资源别名</p>
	ResourceAlias *string `json:"ResourceAlias,omitnil,omitempty" name:"ResourceAlias"`

	// <p>uin</p>
	Uin *int64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>资源类型</p>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>地域编码</p>
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// <p>可用区编码</p>
	ZoneCode *string `json:"ZoneCode,omitnil,omitempty" name:"ZoneCode"`

	// <p>付费类型，包括后付费(0)、预付费(1)、预留实例(2)</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>资源创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>资源过期时间</p>
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// <p>内网IP</p>
	PrivateIpAddress []*string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// <p>外网IP</p>
	PublicIpAddress []*string `json:"PublicIpAddress,omitnil,omitempty" name:"PublicIpAddress"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

// Predefined struct for user
type SearchResourcesRequestParams struct {
	// <p>视图ID</p>
	ViewId *string `json:"ViewId,omitnil,omitempty" name:"ViewId"`

	// <p>每页返回的最大记录数</p>
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>分页Token，首次查询不传</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// <p>过滤条件列表</p>
	Filters []*ExtendedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>排序条件</p><p>枚举值：</p><ul><li>CreateTime： 表示按资源创建时间排序</li><li>ExpireTime： 表示按资源到期时间排序</li><li>IpAddress： 表示按资源IP地址排序</li></ul>
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// <p>排序顺序</p><p>枚举值：</p><ul><li>Asc： 升序</li><li>Desc： 降序</li></ul><p>默认值：Asc</p>
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type SearchResourcesRequest struct {
	*tchttp.BaseRequest
	
	// <p>视图ID</p>
	ViewId *string `json:"ViewId,omitnil,omitempty" name:"ViewId"`

	// <p>每页返回的最大记录数</p>
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>分页Token，首次查询不传</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// <p>过滤条件列表</p>
	Filters []*ExtendedFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>排序条件</p><p>枚举值：</p><ul><li>CreateTime： 表示按资源创建时间排序</li><li>ExpireTime： 表示按资源到期时间排序</li><li>IpAddress： 表示按资源IP地址排序</li></ul>
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// <p>排序顺序</p><p>枚举值：</p><ul><li>Asc： 升序</li><li>Desc： 降序</li></ul><p>默认值：Asc</p>
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *SearchResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ViewId")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchResourcesResponseParams struct {
	// <p>下一页Token，为空时表示无更多数据</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// <p>资源列表</p>
	Resources []*ResourceSummary `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchResourcesResponse struct {
	*tchttp.BaseResponse
	Response *SearchResourcesResponseParams `json:"Response"`
}

func (r *SearchResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}