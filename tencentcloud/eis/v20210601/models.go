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

package v20210601

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AbstractRuntimeMC struct {
	// 环境id
	RuntimeId *int64 `json:"RuntimeId,omitempty" name:"RuntimeId"`

	// 环境名称，用户输入，同一uin内唯一
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 环境类型：0: sandbox, 1:shared, 2:private
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 环境所在地域，tianjin，beijiing，guangzhou等
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 环境所在地域，tianjin，beijiing，guangzhou等（同Zone）
	Area *string `json:"Area,omitempty" name:"Area"`

	// 环境应用listener地址后缀
	Addr *string `json:"Addr,omitempty" name:"Addr"`

	// 环境状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 环境过期时间
	ExpiredAt *int64 `json:"ExpiredAt,omitempty" name:"ExpiredAt"`

	// 环境运行类型：0:运行时类型、1:api类型
	RuntimeClass *int64 `json:"RuntimeClass,omitempty" name:"RuntimeClass"`

	// 是否已在当前环境发布
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deployed *bool `json:"Deployed,omitempty" name:"Deployed"`
}

// Predefined struct for user
type GetRuntimeMCRequestParams struct {
	// 环境id
	RuntimeId *int64 `json:"RuntimeId,omitempty" name:"RuntimeId"`

	// 环境地域
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 环境运行类型：0:运行时类型、1:api类型
	RuntimeClass *int64 `json:"RuntimeClass,omitempty" name:"RuntimeClass"`
}

type GetRuntimeMCRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	RuntimeId *int64 `json:"RuntimeId,omitempty" name:"RuntimeId"`

	// 环境地域
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 环境运行类型：0:运行时类型、1:api类型
	RuntimeClass *int64 `json:"RuntimeClass,omitempty" name:"RuntimeClass"`
}

func (r *GetRuntimeMCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRuntimeMCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuntimeId")
	delete(f, "Zone")
	delete(f, "RuntimeClass")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRuntimeMCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRuntimeMCResponseParams struct {
	// 运行时详情
	Runtime *RuntimeMC `json:"Runtime,omitempty" name:"Runtime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRuntimeMCResponse struct {
	*tchttp.BaseResponse
	Response *GetRuntimeMCResponseParams `json:"Response"`
}

func (r *GetRuntimeMCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRuntimeMCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRuntimeResourceMonitorMetricMCRequestParams struct {
	// 运行时id
	RuntimeId *int64 `json:"RuntimeId,omitempty" name:"RuntimeId"`

	// 起始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 指标类型：0:CPU, 1:MemUsageBytes, 2:K8sWorkloadNetworkReceiveBytesBw, 3:K8sWorkloadNetworkTransmitBytesBw
	MetricType *int64 `json:"MetricType,omitempty" name:"MetricType"`

	// 是否返回百分比数值，仅支持CPU，Memory
	RateType *bool `json:"RateType,omitempty" name:"RateType"`

	// 采样粒度：60(s), 300(s), 3600(s), 86400(s)
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// 环境运行类型：0:运行时类型、1:api类型
	RuntimeClass *int64 `json:"RuntimeClass,omitempty" name:"RuntimeClass"`
}

type GetRuntimeResourceMonitorMetricMCRequest struct {
	*tchttp.BaseRequest
	
	// 运行时id
	RuntimeId *int64 `json:"RuntimeId,omitempty" name:"RuntimeId"`

	// 起始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 指标类型：0:CPU, 1:MemUsageBytes, 2:K8sWorkloadNetworkReceiveBytesBw, 3:K8sWorkloadNetworkTransmitBytesBw
	MetricType *int64 `json:"MetricType,omitempty" name:"MetricType"`

	// 是否返回百分比数值，仅支持CPU，Memory
	RateType *bool `json:"RateType,omitempty" name:"RateType"`

	// 采样粒度：60(s), 300(s), 3600(s), 86400(s)
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// 环境运行类型：0:运行时类型、1:api类型
	RuntimeClass *int64 `json:"RuntimeClass,omitempty" name:"RuntimeClass"`
}

func (r *GetRuntimeResourceMonitorMetricMCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRuntimeResourceMonitorMetricMCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuntimeId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricType")
	delete(f, "RateType")
	delete(f, "Interval")
	delete(f, "RuntimeClass")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRuntimeResourceMonitorMetricMCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRuntimeResourceMonitorMetricMCResponseParams struct {
	// 指标名称，K8sWorkloadCpuCoreUsed，K8sWorkloadMemUsageBytes，K8sWorkloadNetworkReceiveBytesBw，K8sWorkloadNetworkTransmitBytesBw
	MetricType *string `json:"MetricType,omitempty" name:"MetricType"`

	// metric数值列表
	Values []*MetricValueMC `json:"Values,omitempty" name:"Values"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRuntimeResourceMonitorMetricMCResponse struct {
	*tchttp.BaseResponse
	Response *GetRuntimeResourceMonitorMetricMCResponseParams `json:"Response"`
}

func (r *GetRuntimeResourceMonitorMetricMCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRuntimeResourceMonitorMetricMCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDeployableRuntimesMCRequestParams struct {

}

type ListDeployableRuntimesMCRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListDeployableRuntimesMCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDeployableRuntimesMCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDeployableRuntimesMCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDeployableRuntimesMCResponseParams struct {
	// 运行时列表
	Runtimes []*AbstractRuntimeMC `json:"Runtimes,omitempty" name:"Runtimes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListDeployableRuntimesMCResponse struct {
	*tchttp.BaseResponse
	Response *ListDeployableRuntimesMCResponseParams `json:"Response"`
}

func (r *ListDeployableRuntimesMCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDeployableRuntimesMCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRuntimeDeployedInstancesMCRequestParams struct {
	// 运行时id
	RuntimeId *int64 `json:"RuntimeId,omitempty" name:"RuntimeId"`

	// 最大请求数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 请求偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序类型：1:创建时间排序, 2:更新时间排序（默认）
	SortType *int64 `json:"SortType,omitempty" name:"SortType"`

	// 排序方式：asc，desc（默认）
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 运行时地域
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 1:3.0版本新控制台传1；否则传0
	ApiVersion *int64 `json:"ApiVersion,omitempty" name:"ApiVersion"`

	// -1:不按项目筛选，获取所有
	// >=0: 按项目id筛选
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// -2: 不按状态筛选，获取所有
	// 0: 运行中
	// 2: 已停止
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ListRuntimeDeployedInstancesMCRequest struct {
	*tchttp.BaseRequest
	
	// 运行时id
	RuntimeId *int64 `json:"RuntimeId,omitempty" name:"RuntimeId"`

	// 最大请求数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 请求偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序类型：1:创建时间排序, 2:更新时间排序（默认）
	SortType *int64 `json:"SortType,omitempty" name:"SortType"`

	// 排序方式：asc，desc（默认）
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 运行时地域
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 1:3.0版本新控制台传1；否则传0
	ApiVersion *int64 `json:"ApiVersion,omitempty" name:"ApiVersion"`

	// -1:不按项目筛选，获取所有
	// >=0: 按项目id筛选
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// -2: 不按状态筛选，获取所有
	// 0: 运行中
	// 2: 已停止
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ListRuntimeDeployedInstancesMCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRuntimeDeployedInstancesMCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuntimeId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SortType")
	delete(f, "Sort")
	delete(f, "Zone")
	delete(f, "ApiVersion")
	delete(f, "GroupId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRuntimeDeployedInstancesMCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRuntimeDeployedInstancesMCResponseParams struct {
	// 运行时所部属的应用实例列表
	Instances []*RuntimeDeployedInstanceMC `json:"Instances,omitempty" name:"Instances"`

	// 满足条件的记录总数，用于分页器
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListRuntimeDeployedInstancesMCResponse struct {
	*tchttp.BaseResponse
	Response *ListRuntimeDeployedInstancesMCResponseParams `json:"Response"`
}

func (r *ListRuntimeDeployedInstancesMCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRuntimeDeployedInstancesMCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRuntimesMCRequestParams struct {
	// 环境运行类型：0:运行时类型、1:api类型
	RuntimeClass *int64 `json:"RuntimeClass,omitempty" name:"RuntimeClass"`
}

type ListRuntimesMCRequest struct {
	*tchttp.BaseRequest
	
	// 环境运行类型：0:运行时类型、1:api类型
	RuntimeClass *int64 `json:"RuntimeClass,omitempty" name:"RuntimeClass"`
}

func (r *ListRuntimesMCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRuntimesMCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuntimeClass")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRuntimesMCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRuntimesMCResponseParams struct {
	// 运行时列表
	Runtimes []*RuntimeMC `json:"Runtimes,omitempty" name:"Runtimes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListRuntimesMCResponse struct {
	*tchttp.BaseResponse
	Response *ListRuntimesMCResponseParams `json:"Response"`
}

func (r *ListRuntimesMCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRuntimesMCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricValueMC struct {
	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 对应的value值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Val *string `json:"Val,omitempty" name:"Val"`
}

type RuntimeDeployedInstanceMC struct {
	// 项目id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 项目名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 应用id
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 应用名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 应用实例id
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用实例版本
	InstanceVersion *int64 `json:"InstanceVersion,omitempty" name:"InstanceVersion"`

	// 应用实例创建时间
	InstanceCreatedAt *int64 `json:"InstanceCreatedAt,omitempty" name:"InstanceCreatedAt"`

	// 应用实例部署状态. 0:running, 1:deleting
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 应用实例部署创建时间
	CreatedAt *int64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 应用实例部署更新时间
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 应用类型：0:NormalApp普通应用 1:TemplateApp模板应用 2:LightApp轻应用 3:MicroConnTemplate微连接模板 4:MicroConnApp微连接应用
	ProjectType *int64 `json:"ProjectType,omitempty" name:"ProjectType"`

	// 应用版本：0:旧版 1:3.0新控制台
	ProjectVersion *int64 `json:"ProjectVersion,omitempty" name:"ProjectVersion"`
}

type RuntimeExtensionMC struct {
	// 扩展组件类型：0:cdc
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 部署规格vcore数
	Size *float64 `json:"Size,omitempty" name:"Size"`

	// 副本数
	Replica *int64 `json:"Replica,omitempty" name:"Replica"`

	// 扩展组件名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 状态 1:未启用 2:已启用
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreatedAt *int64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 修改时间
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type RuntimeMC struct {
	// 环境id
	RuntimeId *int64 `json:"RuntimeId,omitempty" name:"RuntimeId"`

	// 主账号uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 环境名称，用户输入，同一uin内唯一
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 环境所在地域，tianjin，beijiing，guangzhou等
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 环境类型：0: sandbox, 1:shared, 2:private 3: trial
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 运行时状态：1:running, 2:deleting, 3:creating, 4:scaling, 5:unavailable, 6:deleted, 7:errored
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 环境创建时间
	CreatedAt *int64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 环境更新时间
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 环境资源配置，worker总配额，0:0vCore0G, 1:1vCore2G, 2:2vCore4G, 4:4vCore8G, 8:8vCore16G, 12:12vCore24G, 16:16vCore32G, 100:unlimited
	WorkerSize *int64 `json:"WorkerSize,omitempty" name:"WorkerSize"`

	// 环境资源配置，worker副本数
	WorkerReplica *int64 `json:"WorkerReplica,omitempty" name:"WorkerReplica"`

	// 正在运行的应用实例数量
	RunningInstanceCount *int64 `json:"RunningInstanceCount,omitempty" name:"RunningInstanceCount"`

	// 已使用cpu核数
	CpuUsed *float64 `json:"CpuUsed,omitempty" name:"CpuUsed"`

	// cpu核数上限
	CpuLimit *float64 `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 已使用内存 MB
	MemoryUsed *float64 `json:"MemoryUsed,omitempty" name:"MemoryUsed"`

	// 内存上限 MB
	MemoryLimit *float64 `json:"MemoryLimit,omitempty" name:"MemoryLimit"`

	// 环境过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredAt *int64 `json:"ExpiredAt,omitempty" name:"ExpiredAt"`

	// 收费类型：0:缺省，1:自助下单页购买(支持续费/升配等操作)，2:代销下单页购买
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *int64 `json:"ChargeType,omitempty" name:"ChargeType"`

	// 资源限制类型：0:无限制，1:有限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceLimitType *int64 `json:"ResourceLimitType,omitempty" name:"ResourceLimitType"`

	// 是否开启自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewal *bool `json:"AutoRenewal,omitempty" name:"AutoRenewal"`

	// 扩展组件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkerExtensions []*RuntimeExtensionMC `json:"WorkerExtensions,omitempty" name:"WorkerExtensions"`

	// 环境类型：0: sandbox, 1:shared, 2:private 3: trial
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeType *int64 `json:"RuntimeType,omitempty" name:"RuntimeType"`

	// 环境运行类型：0:运行时类型、1:api类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeClass *int64 `json:"RuntimeClass,omitempty" name:"RuntimeClass"`

	// 已使用出带宽 Mbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthOutUsed *float64 `json:"BandwidthOutUsed,omitempty" name:"BandwidthOutUsed"`

	// 出带宽上限 Mbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthOutLimit *float64 `json:"BandwidthOutLimit,omitempty" name:"BandwidthOutLimit"`
}