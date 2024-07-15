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

package v20221123

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CloudDedicatedZoneHostsInfo struct {
	// Host的唯一标识uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostUuid *string `json:"HostUuid,omitnil,omitempty" name:"HostUuid"`

	// 实例名称数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancesInfo []*string `json:"InstancesInfo,omitnil,omitempty" name:"InstancesInfo"`
}

type CloudDedicatedZoneResourceStatisticsInfo struct {
	// 资源统计项名称
	Item *string `json:"Item,omitnil,omitempty" name:"Item"`

	// 资源统计项单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 资源总量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 已用资源
	Usage *string `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 已用资源占比
	UsageRate *string `json:"UsageRate,omitnil,omitempty" name:"UsageRate"`

	// 剩余资源
	Remain *string `json:"Remain,omitnil,omitempty" name:"Remain"`

	// 剩余资源占比
	RemainRate *string `json:"RemainRate,omitnil,omitempty" name:"RemainRate"`

	// 本周一零点资源使用率
	ThisMondayUsageRate *string `json:"ThisMondayUsageRate,omitnil,omitempty" name:"ThisMondayUsageRate"`

	// 本周资源增长比例
	ThisMondayUsageGrowthRate *string `json:"ThisMondayUsageGrowthRate,omitnil,omitempty" name:"ThisMondayUsageGrowthRate"`

	// 上周资源增长比例
	LastMondayUsageGrowthRate *string `json:"LastMondayUsageGrowthRate,omitnil,omitempty" name:"LastMondayUsageGrowthRate"`
}

type CloudDedicatedZoneResourceSummaryInfo struct {
	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 子产品名称
	SubProductName *string `json:"SubProductName,omitnil,omitempty" name:"SubProductName"`

	// 资源统计详情
	Statistics []*CloudDedicatedZoneResourceStatisticsInfo `json:"Statistics,omitnil,omitempty" name:"Statistics"`
}

// Predefined struct for user
type DescribeCloudDedicatedZoneHostsRequestParams struct {
	// 专属可用区ID 
	CloudDedicatedZoneID *string `json:"CloudDedicatedZoneID,omitnil,omitempty" name:"CloudDedicatedZoneID"`

	// 一个或多个Host面的CVM实例信息。最大支持查询100台Host。
	HostUuids []*string `json:"HostUuids,omitnil,omitempty" name:"HostUuids"`

	// 查询一个实例或者多个实例所在的Host上面的CVM实例信息。最大支持查询100台实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。该参数仅与CloudDedicatedZoneID有关，传递了HostUuids和InstanceIds则会失效。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。该参数仅与CloudDedicatedZoneID有关，传递了HostUuids和InstanceIds则会失效。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCloudDedicatedZoneHostsRequest struct {
	*tchttp.BaseRequest
	
	// 专属可用区ID 
	CloudDedicatedZoneID *string `json:"CloudDedicatedZoneID,omitnil,omitempty" name:"CloudDedicatedZoneID"`

	// 一个或多个Host面的CVM实例信息。最大支持查询100台Host。
	HostUuids []*string `json:"HostUuids,omitnil,omitempty" name:"HostUuids"`

	// 查询一个实例或者多个实例所在的Host上面的CVM实例信息。最大支持查询100台实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。该参数仅与CloudDedicatedZoneID有关，传递了HostUuids和InstanceIds则会失效。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。该参数仅与CloudDedicatedZoneID有关，传递了HostUuids和InstanceIds则会失效。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeCloudDedicatedZoneHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudDedicatedZoneHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CloudDedicatedZoneID")
	delete(f, "HostUuids")
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudDedicatedZoneHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudDedicatedZoneHostsResponseParams struct {
	// 返回Host和Host上部署的实例信息之间的关系
	CloudDedicatedZoneHostsInfoSet []*CloudDedicatedZoneHostsInfo `json:"CloudDedicatedZoneHostsInfoSet,omitnil,omitempty" name:"CloudDedicatedZoneHostsInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudDedicatedZoneHostsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudDedicatedZoneHostsResponseParams `json:"Response"`
}

func (r *DescribeCloudDedicatedZoneHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudDedicatedZoneHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudDedicatedZoneResourceSummaryRequestParams struct {
	// 专属可用区唯一标识
	CdzId *string `json:"CdzId,omitnil,omitempty" name:"CdzId"`
}

type DescribeCloudDedicatedZoneResourceSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 专属可用区唯一标识
	CdzId *string `json:"CdzId,omitnil,omitempty" name:"CdzId"`
}

func (r *DescribeCloudDedicatedZoneResourceSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudDedicatedZoneResourceSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CdzId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudDedicatedZoneResourceSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudDedicatedZoneResourceSummaryResponseParams struct {
	// 资源水位详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceSummarySet []*CloudDedicatedZoneResourceSummaryInfo `json:"ResourceSummarySet,omitnil,omitempty" name:"ResourceSummarySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudDedicatedZoneResourceSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudDedicatedZoneResourceSummaryResponseParams `json:"Response"`
}

func (r *DescribeCloudDedicatedZoneResourceSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudDedicatedZoneResourceSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}