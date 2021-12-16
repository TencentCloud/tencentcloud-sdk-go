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

package v20201106

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest

	// 待查询产品的名称，例如cvm、vpc
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 地域数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 地域列表信息
		RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest

	// 待查询产品的名称，例如cvm、vpc
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可用区数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可用区列表信息。
		ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 地域名称，例如，ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域描述，例如，华南地区(广州)
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域是否可用状态
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
}

type ZoneInfo struct {

	// 可用区名称，例如，ap-guangzhou-3
	// 全网可用区名称如下：
	// <li> ap-chongqing-1 </li>
	// <li> ap-seoul-1 </li>
	// <li> ap-seoul-2 </li>
	// <li> ap-chengdu-1 </li>
	// <li> ap-chengdu-2 </li>
	// <li> ap-hongkong-1 </li>
	// <li> ap-hongkong-2 </li>
	// <li> ap-shenzhen-fsi-1 </li>
	// <li> ap-shenzhen-fsi-2 </li>
	// <li> ap-shenzhen-fsi-3 </li>
	// <li> ap-guangzhou-1（售罄）</li>
	// <li> ap-guangzhou-2（售罄）</li>
	// <li> ap-guangzhou-3 </li>
	// <li> ap-guangzhou-4 </li>
	// <li> ap-guangzhou-6 </li>
	// <li> ap-tokyo-1 </li>
	// <li> ap-singapore-1 </li>
	// <li> ap-singapore-2 </li>
	// <li> ap-shanghai-fsi-1 </li>
	// <li> ap-shanghai-fsi-2 </li>
	// <li> ap-shanghai-fsi-3 </li>
	// <li> ap-bangkok-1 </li>
	// <li> ap-shanghai-1（售罄） </li>
	// <li> ap-shanghai-2 </li>
	// <li> ap-shanghai-3 </li>
	// <li> ap-shanghai-4 </li>
	// <li> ap-shanghai-5 </li>
	// <li> ap-mumbai-1 </li>
	// <li> ap-mumbai-2 </li>
	// <li> eu-moscow-1 </li>
	// <li> ap-beijing-1 </li>
	// <li> ap-beijing-2 </li>
	// <li> ap-beijing-3 </li>
	// <li> ap-beijing-4 </li>
	// <li> ap-beijing-5 </li>
	// <li> na-siliconvalley-1 </li>
	// <li> na-siliconvalley-2 </li>
	// <li> eu-frankfurt-1 </li>
	// <li> na-toronto-1 </li>
	// <li> na-ashburn-1 </li>
	// <li> na-ashburn-2 </li>
	// <li> ap-nanjing-1 </li>
	// <li> ap-nanjing-2 </li>
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区描述，例如，广州三区
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 可用区ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区状态，包含AVAILABLE和UNAVAILABLE。AVAILABLE代表可用，UNAVAILABLE代表不可用。
	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}
