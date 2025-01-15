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

package v20230306

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeEventsRequestParams struct {
	// 事件的发生日期
	EventDate *string `json:"EventDate,omitnil,omitempty" name:"EventDate"`

	// 1. 不指定产品列表时将查询所有产品。
	// 2. 产品ID示例：cvm、lb、cdb、cdn、crs
	ProductIds []*string `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`

	// 1. 不指定地域列表时将查询所有地域。
	// 2. 查询非区域性产品事件时，地域ID指定为：non-regional
	// 3. 其他地域ID取值请参考：https://cloud.tencent.com/document/api/213/15692
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`
}

type DescribeEventsRequest struct {
	*tchttp.BaseRequest
	
	// 事件的发生日期
	EventDate *string `json:"EventDate,omitnil,omitempty" name:"EventDate"`

	// 1. 不指定产品列表时将查询所有产品。
	// 2. 产品ID示例：cvm、lb、cdb、cdn、crs
	ProductIds []*string `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`

	// 1. 不指定地域列表时将查询所有地域。
	// 2. 查询非区域性产品事件时，地域ID指定为：non-regional
	// 3. 其他地域ID取值请参考：https://cloud.tencent.com/document/api/213/15692
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`
}

func (r *DescribeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventDate")
	delete(f, "ProductIds")
	delete(f, "RegionIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsResponseParams struct {
	// 事件详情列表
	Data *ProductEventList `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventsResponseParams `json:"Response"`
}

func (r *DescribeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventDetail struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 地域ID，非区域性地域返回non-regional
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 事件开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事件结束时间，当事件正在发生还未结束时，结束时间返回空
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 事件当前状态：提示、异常、正常
	CurrentStatus *string `json:"CurrentStatus,omitnil,omitempty" name:"CurrentStatus"`
}

type ProductEventList struct {
	// 事件详情列表
	EventList []*EventDetail `json:"EventList,omitnil,omitempty" name:"EventList"`
}