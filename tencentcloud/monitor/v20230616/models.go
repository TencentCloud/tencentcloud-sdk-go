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

package v20230616

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type DataPoint struct {
	// 实例对象维度组合
	Dimensions []*Dimension `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// 时间戳数组，表示那些时间点有数据，缺失的时间戳，没有数据点，可以理解为掉点了
	Timestamps []*float64 `json:"Timestamps,omitnil,omitempty" name:"Timestamps"`

	// 监控值数组，该数组和Timestamps一一对应
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`
}

type Dimension struct {
	// 实例维度名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例维度值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Expr struct {
	// 算子名称
	Function *string `json:"Function,omitnil,omitempty" name:"Function"`

	// 算子入参值
	N *float64 `json:"N,omitnil,omitempty" name:"N"`
}

// Predefined struct for user
type GetMonitorDataRequestParams struct {
	// 命名空间，如QCE/CVM。各个云产品的详细命名空间说明请参阅各个产品[监控指标](https://cloud.tencent.com/document/product/248/6140)文档
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 指标名称，如CPUUsage，仅支持单指标拉取。各个云产品的详细指标说明请参阅各个产品[监控指标](https://cloud.tencent.com/document/product/248/6140)文档，对应的指标英文名即为MetricName
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 实例对象的维度组合，格式为key-value键值对形式的集合。不同类型的实例字段完全不同，如CVM为[{"Name":"InstanceId","Value":"ins-j0hk02zo"}]，Ckafka为[{"Name":"instanceId","Value":"ckafka-l49k54dd"}]，COS为[{"Name":"appid","Value":"1258344699"},{"Name":"bucket","Value":"rig-1258344699"}]。各个云产品的维度请参阅各个产品[监控指标](https://cloud.tencent.com/document/product/248/6140)文档，对应的维度列即为维度组合的key，value为key对应的值。单请求最多支持批量拉取10个实例的监控数据。
	Instances []*Instance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 监控统计周期，如60。默认为取值为300，单位为s。每个指标支持的统计周期不一定相同，各个云产品支持的统计周期请参阅各个产品[监控指标](https://cloud.tencent.com/document/product/248/6140)文档，对应的统计周期列即为支持的统计周期。单请求的数据点数限制为1440个。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 起始时间，如2018-09-22T19:51:23+08:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如2018-09-22T20:51:23+08:00，默认为当前时间。 EndTime不能小于StartTime
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 计算算子
	Expr *Expr `json:"Expr,omitnil,omitempty" name:"Expr"`
}

type GetMonitorDataRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间，如QCE/CVM。各个云产品的详细命名空间说明请参阅各个产品[监控指标](https://cloud.tencent.com/document/product/248/6140)文档
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 指标名称，如CPUUsage，仅支持单指标拉取。各个云产品的详细指标说明请参阅各个产品[监控指标](https://cloud.tencent.com/document/product/248/6140)文档，对应的指标英文名即为MetricName
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 实例对象的维度组合，格式为key-value键值对形式的集合。不同类型的实例字段完全不同，如CVM为[{"Name":"InstanceId","Value":"ins-j0hk02zo"}]，Ckafka为[{"Name":"instanceId","Value":"ckafka-l49k54dd"}]，COS为[{"Name":"appid","Value":"1258344699"},{"Name":"bucket","Value":"rig-1258344699"}]。各个云产品的维度请参阅各个产品[监控指标](https://cloud.tencent.com/document/product/248/6140)文档，对应的维度列即为维度组合的key，value为key对应的值。单请求最多支持批量拉取10个实例的监控数据。
	Instances []*Instance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 监控统计周期，如60。默认为取值为300，单位为s。每个指标支持的统计周期不一定相同，各个云产品支持的统计周期请参阅各个产品[监控指标](https://cloud.tencent.com/document/product/248/6140)文档，对应的统计周期列即为支持的统计周期。单请求的数据点数限制为1440个。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 起始时间，如2018-09-22T19:51:23+08:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如2018-09-22T20:51:23+08:00，默认为当前时间。 EndTime不能小于StartTime
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 计算算子
	Expr *Expr `json:"Expr,omitnil,omitempty" name:"Expr"`
}

func (r *GetMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMonitorDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Namespace")
	delete(f, "MetricName")
	delete(f, "Instances")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Expr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMonitorDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMonitorDataResponseParams struct {
	// 统计周期
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 指标名
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 数据点数组
	DataPoints []*DataPoint `json:"DataPoints,omitnil,omitempty" name:"DataPoints"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 返回信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *GetMonitorDataResponseParams `json:"Response"`
}

func (r *GetMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {
	// 实例的维度组合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions []*Dimension `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`
}