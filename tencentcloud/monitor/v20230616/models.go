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

package v20230616

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeAlarmNotifyHistoriesRequestParams struct {
	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 起始时间点，unix秒级时间戳
	QueryBaseTime *int64 `json:"QueryBaseTime,omitnil,omitempty" name:"QueryBaseTime"`

	// 从 QueryBaseTime 开始，需要查询往前多久的时间，单位秒
	QueryBeforeSeconds *int64 `json:"QueryBeforeSeconds,omitnil,omitempty" name:"QueryBeforeSeconds"`

	// 分页参数
	PageParams *PageByNoParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`

	// 当监控类型为 MT_QCE 时候需要填写，归属的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 当监控类型为 MT_QCE 时候需要填写， 告警策略类型
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 查询某个策略的通知历史
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DescribeAlarmNotifyHistoriesRequest struct {
	*tchttp.BaseRequest
	
	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 起始时间点，unix秒级时间戳
	QueryBaseTime *int64 `json:"QueryBaseTime,omitnil,omitempty" name:"QueryBaseTime"`

	// 从 QueryBaseTime 开始，需要查询往前多久的时间，单位秒
	QueryBeforeSeconds *int64 `json:"QueryBeforeSeconds,omitnil,omitempty" name:"QueryBeforeSeconds"`

	// 分页参数
	PageParams *PageByNoParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`

	// 当监控类型为 MT_QCE 时候需要填写，归属的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 当监控类型为 MT_QCE 时候需要填写， 告警策略类型
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 查询某个策略的通知历史
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmNotifyHistoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNotifyHistoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorType")
	delete(f, "QueryBaseTime")
	delete(f, "QueryBeforeSeconds")
	delete(f, "PageParams")
	delete(f, "Namespace")
	delete(f, "ModelName")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNotifyHistoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNotifyHistoriesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmNotifyHistoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmNotifyHistoriesResponseParams `json:"Response"`
}

func (r *DescribeAlarmNotifyHistoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNotifyHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PageByNoParams struct {
	// 每个分页的数量是多少
	// 注意：此字段可能返回 null，表示取不到有效值。
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// 第几个分页，从1开始
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNo *string `json:"PageNo,omitnil,omitempty" name:"PageNo"`
}