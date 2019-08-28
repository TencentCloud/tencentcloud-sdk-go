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

package v20190416

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Conditions struct {

	// 原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 具有相同原因的副本个数
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type CreateServiceConfigRequest struct {
	*tchttp.BaseRequest

	// 配置名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 运行环境
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 模型地址，支持cos路径，格式为 cos://bucket名-appid.cos.region名.myqcloud.com/模型文件夹路径。为模型文件的上一层文件夹地址。
	ModelUri *string `json:"ModelUri,omitempty" name:"ModelUri"`

	// 处理器配置, 单位为1/1000核；范围[100, 256000]
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存配置, 单位为1M；范围[100, 256000]
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// GPU算力配置，单位为1/100 tflops，范围 [0, 256000]
	TflopUnits *uint64 `json:"TflopUnits,omitempty" name:"TflopUnits"`

	// 显存配置, 单位为1M，范围 [0, 256000]
	GpuMemory *uint64 `json:"GpuMemory,omitempty" name:"GpuMemory"`
}

func (r *CreateServiceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务配置
		ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceRequest struct {
	*tchttp.BaseRequest

	// 扩缩容配置
	Scaler *Scaler `json:"Scaler,omitempty" name:"Scaler"`

	// 服务配置Id
	ServiceConfigId *int64 `json:"ServiceConfigId,omitempty" name:"ServiceConfigId"`

	// 服务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 扩缩容方式，支持AUTO, MANUAL，分别表示自动扩缩容和手动扩缩容
	ScaleMode *string `json:"ScaleMode,omitempty" name:"ScaleMode"`

	// 集群，不填则使用默认集群。
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`
}

func (r *CreateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务
		Service *Service `json:"Service,omitempty" name:"Service"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceConfigRequest struct {
	*tchttp.BaseRequest

	// 服务配置Id (deprecated)
	ServiceConfigId *int64 `json:"ServiceConfigId,omitempty" name:"ServiceConfigId"`

	// 服务配置名称
	ServiceConfigName *string `json:"ServiceConfigName,omitempty" name:"ServiceConfigName"`
}

func (r *DeleteServiceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceRequest struct {
	*tchttp.BaseRequest

	// 服务Id
	ServiceId *uint64 `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DeleteServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRuntimesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRuntimesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRuntimesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRuntimesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TIEMS支持的运行环境列表
		Runtimes []*Runtime `json:"Runtimes,omitempty" name:"Runtimes" list`

		// 用户对runtime对权限
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserAccess *uint64 `json:"UserAccess,omitempty" name:"UserAccess"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuntimesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRuntimesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceConfigsRequest struct {
	*tchttp.BaseRequest

	// 筛选选项，支持按照name等进行筛选
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME", "UPDATE_TIME", "NAME"
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeServiceConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务配置
		ServiceConfigs []*ServiceConfig `json:"ServiceConfigs,omitempty" name:"ServiceConfigs" list`

		// 服务配置总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServicesRequest struct {
	*tchttp.BaseRequest

	// 筛选选项，支持按照name等字段进行筛选
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME" "UPDATE_TIME"
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务列表
		Services []*Service `json:"Services,omitempty" name:"Services" list`

		// 服务总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 取值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type Option struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 取值
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type Runtime struct {

	// 运行环境名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 运行环境框架
	Framework *string `json:"Framework,omitempty" name:"Framework"`

	// 运行环境描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否为公开运行环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	Public *bool `json:"Public,omitempty" name:"Public"`

	// 是否打开健康检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckOn *bool `json:"HealthCheckOn,omitempty" name:"HealthCheckOn"`

	// 镜像地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Scaler struct {

	// 最大副本数
	MaxReplicas *uint64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`

	// 最小副本数
	MinReplicas *uint64 `json:"MinReplicas,omitempty" name:"MinReplicas"`

	// 起始副本数
	StartReplicas *uint64 `json:"StartReplicas,omitempty" name:"StartReplicas"`

	// 扩缩容指标，选择自动扩缩容时至少需要选择一个指标，支持CPU-UTIL、MEMORY-UTIL
	HpaMetrics []*Option `json:"HpaMetrics,omitempty" name:"HpaMetrics" list`
}

type Service struct {

	// 服务ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 运行集群
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 服务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 运行环境
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 模型地址
	ModelUri *string `json:"ModelUri,omitempty" name:"ModelUri"`

	// 处理器配置, 单位为1/1000核
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存配置, 单位为1M
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 处理器配置, 单位为1/100 tflops
	TflopUnits *uint64 `json:"TflopUnits,omitempty" name:"TflopUnits"`

	// 显存配置, 单位为1M
	GpuMemory *uint64 `json:"GpuMemory,omitempty" name:"GpuMemory"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 支持AUTO, MANUAL
	ScaleMode *string `json:"ScaleMode,omitempty" name:"ScaleMode"`

	// 弹性伸缩配置
	Scaler *Scaler `json:"Scaler,omitempty" name:"Scaler"`

	// 服务状态
	Status *ServiceStatus `json:"Status,omitempty" name:"Status"`

	// 服务地址
	ServingIp *string `json:"ServingIp,omitempty" name:"ServingIp"`

	// 访问密钥
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 服务配置Id
	ServiceConfigId *int64 `json:"ServiceConfigId,omitempty" name:"ServiceConfigId"`

	// 服务配置名
	ServiceConfigName *string `json:"ServiceConfigName,omitempty" name:"ServiceConfigName"`

	// 服务运行时长
	ServeSeconds *uint64 `json:"ServeSeconds,omitempty" name:"ServeSeconds"`

	// 配置版本
	ServiceConfigVersion *string `json:"ServiceConfigVersion,omitempty" name:"ServiceConfigVersion"`
}

type ServiceConfig struct {

	// Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模型地址
	ModelUri *string `json:"ModelUri,omitempty" name:"ModelUri"`

	// 处理器配置, 单位为1/1000核
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存配置, 单位为1M
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// GPU算力，单位为1/100 tflops
	// 注意：此字段可能返回 null，表示取不到有效值。
	TflopUnits *uint64 `json:"TflopUnits,omitempty" name:"TflopUnits"`

	// 显存配置, 单位为1M
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuMemory *uint64 `json:"GpuMemory,omitempty" name:"GpuMemory"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 运行环境
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 配置版本
	Version *string `json:"Version,omitempty" name:"Version"`
}

type ServiceStatus struct {

	// 预期副本数
	DesiredReplicas *uint64 `json:"DesiredReplicas,omitempty" name:"DesiredReplicas"`

	// 当前副本数
	CurrentReplicas *uint64 `json:"CurrentReplicas,omitempty" name:"CurrentReplicas"`

	// Normal：正常运行中；Abnormal：服务异常，例如容器启动失败等；Waiting：服务等待中，例如容器下载镜像过程等；Stopped：已停止 Stopping 停止中；Resuming：重启中；Updating：服务更新中
	Status *string `json:"Status,omitempty" name:"Status"`

	// 服务处于当前状态的原因集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*Conditions `json:"Conditions,omitempty" name:"Conditions" list`

	// 副本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas []*string `json:"Replicas,omitempty" name:"Replicas" list`
}

type UpdateServiceRequest struct {
	*tchttp.BaseRequest

	// 服务Id
	ServiceId *int64 `json:"ServiceId,omitempty" name:"ServiceId"`

	// 扩缩容配置
	Scaler *Scaler `json:"Scaler,omitempty" name:"Scaler"`

	// 服务配置Id
	ServiceConfigId *int64 `json:"ServiceConfigId,omitempty" name:"ServiceConfigId"`

	// 支持AUTO, MANUAL，分别表示自动扩缩容，手动扩缩容
	ScaleMode *string `json:"ScaleMode,omitempty" name:"ScaleMode"`

	// 支持STOP(停止) RESUME(重启)
	ServiceAction *string `json:"ServiceAction,omitempty" name:"ServiceAction"`
}

func (r *UpdateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务
		Service *Service `json:"Service,omitempty" name:"Service"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
