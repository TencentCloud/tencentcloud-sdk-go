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

type Config struct {

	// Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模型地址
	ModelUri *string `json:"ModelUri,omitempty" name:"ModelUri"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 运行环境
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 配置版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 配置描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateJobRequest struct {
	*tchttp.BaseRequest

	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 同时处理任务的 Worker 个数
	WorkerCount *uint64 `json:"WorkerCount,omitempty" name:"WorkerCount"`

	// 使用的配置 Id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 处理器配置, 单位为1/1000核；范围[100, 256000]
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存配置, 单位为1M；范围[100, 256000]
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 运行集群
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 预测输入
	PredictInput *PredictInput `json:"PredictInput,omitempty" name:"PredictInput"`

	// 任务描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 使用的资源组 Id，默认使用共享资源组
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`

	// GPU算力配置，单位为1/1000 卡，范围 [0, 256000]
	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`

	// 显存配置, 单位为1M，范围 [0, 256000]
	GpuMemory *uint64 `json:"GpuMemory,omitempty" name:"GpuMemory"`

	// GPU类型
	GpuType *string `json:"GpuType,omitempty" name:"GpuType"`
}

func (r *CreateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务
		Job *Job `json:"Job,omitempty" name:"Job"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRuntimeRequest struct {
	*tchttp.BaseRequest

	// 全局唯一的运行环境名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 运行环境镜像地址
	Image *string `json:"Image,omitempty" name:"Image"`

	// 运行环境框架
	Framework *string `json:"Framework,omitempty" name:"Framework"`

	// 运行环境描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否支持健康检查，默认为False
	HealthCheckOn *bool `json:"HealthCheckOn,omitempty" name:"HealthCheckOn"`
}

func (r *CreateRuntimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRuntimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRuntimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 运行环境
		Runtime *Runtime `json:"Runtime,omitempty" name:"Runtime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRuntimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRuntimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceConfigRequest struct {
	*tchttp.BaseRequest

	// 配置名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 运行环境
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 模型地址，支持cos路径，格式为 cos://bucket名-appid.cos.region名.myqcloud.com/模型文件夹路径。为模型文件的上一层文件夹地址。
	ModelUri *string `json:"ModelUri,omitempty" name:"ModelUri"`

	// 配置描述
	Description *string `json:"Description,omitempty" name:"Description"`
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
		ServiceConfig *Config `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

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
	ServiceConfigId *string `json:"ServiceConfigId,omitempty" name:"ServiceConfigId"`

	// 服务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 扩缩容方式，支持AUTO, MANUAL，分别表示自动扩缩容和手动扩缩容
	ScaleMode *string `json:"ScaleMode,omitempty" name:"ScaleMode"`

	// 处理器配置, 单位为1/1000核；范围[100, 256000]
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存配置, 单位为1M；范围[100, 256000]
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 集群，不填则使用默认集群
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 部署要使用的资源组Id，默认为共享资源组
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`

	// 默认为空，表示不需要鉴权，TOKEN 表示选择 Token 鉴权方式
	Authentication *string `json:"Authentication,omitempty" name:"Authentication"`

	// GPU算力配置，单位为1/1000 卡，范围 [0, 256000]
	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`

	// 显存配置, 单位为1M，范围 [0, 256000]
	GpuMemory *uint64 `json:"GpuMemory,omitempty" name:"GpuMemory"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`

	// GPU类型
	GpuType *string `json:"GpuType,omitempty" name:"GpuType"`
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
		Service *ModelService `json:"Service,omitempty" name:"Service"`

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

type DeleteJobRequest struct {
	*tchttp.BaseRequest

	// 任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DeleteJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRuntimeRequest struct {
	*tchttp.BaseRequest

	// 要删除的Runtime名
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`
}

func (r *DeleteRuntimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRuntimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRuntimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRuntimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRuntimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceConfigRequest struct {
	*tchttp.BaseRequest

	// 服务配置Id
	ServiceConfigId *string `json:"ServiceConfigId,omitempty" name:"ServiceConfigId"`

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
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
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

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 筛选选项
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CREATE_TIME", "UPDATE_TIME", "NAME"
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 要查询的资源组 ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资源组下节点总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 资源组下节点列表
		Instances []*Instance `json:"Instances,omitempty" name:"Instances" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
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

	// 是否按照配置名分页
	PageByName *bool `json:"PageByName,omitempty" name:"PageByName"`
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
		ServiceConfigs []*Config `json:"ServiceConfigs,omitempty" name:"ServiceConfigs" list`

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
		Services []*ModelService `json:"Services,omitempty" name:"Services" list`

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

type Expose struct {

	// 暴露方式，支持 EXTERNAL（外网暴露），VPC （VPC内网打通）
	ExposeType *string `json:"ExposeType,omitempty" name:"ExposeType"`

	// 暴露Ip。暴露方式为 EXTERNAL 为外网 Ip，暴露方式为 VPC 时为指定 Vpc 下的Vip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 暴露方式为 VPC 时，打通的私有网络Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`

	// 暴露方式为 VPC 时，打通的子网Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnSubnetId *string `json:"UnSubnetId,omitempty" name:"UnSubnetId"`
}

type Filter struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 取值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type Instance struct {

	// 节点 ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 节点所在地区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 节点类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 节点充值类型
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Cpu 核数
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Gpu 核数
	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`

	// 节点状态
	State *string `json:"State,omitempty" name:"State"`

	// 节点故障信息
	AbnormalReason *string `json:"AbnormalReason,omitempty" name:"AbnormalReason"`

	// 创建时间
	Created *string `json:"Created,omitempty" name:"Created"`

	// 更新时间
	Updated *string `json:"Updated,omitempty" name:"Updated"`

	// 到期时间
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// 所属资源组 ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`

	// 自动续费标签
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 节点所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 当前 Cpu 申请使用量
	CpuRequested *uint64 `json:"CpuRequested,omitempty" name:"CpuRequested"`

	// 当前 Memory 申请使用量
	MemoryRequested *uint64 `json:"MemoryRequested,omitempty" name:"MemoryRequested"`

	// 当前 Gpu 申请使用量
	GpuRequested *uint64 `json:"GpuRequested,omitempty" name:"GpuRequested"`
}

type Job struct {

	// 任务 Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 集群名
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// Region 名
	Region *string `json:"Region,omitempty" name:"Region"`

	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// Worker 使用的运行环境
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 配置 Id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 预测输入
	PredictInput *PredictInput `json:"PredictInput,omitempty" name:"PredictInput"`

	// 任务状态
	Status *JobStatus `json:"Status,omitempty" name:"Status"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务取消时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CancelTime *string `json:"CancelTime,omitempty" name:"CancelTime"`

	// 任务使用资源组 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`

	// 处理器配置, 单位为1/1000核；范围[100, 256000]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存配置, 单位为1M；范围[100, 256000]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// GPU算力配置，单位为1/1000 卡，范围 [0, 256000]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`

	// 显存配置, 单位为1M，范围 [0, 256000]
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuMemory *uint64 `json:"GpuMemory,omitempty" name:"GpuMemory"`

	// 任务使用资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" name:"ResourceGroupName"`

	// GPU类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuType *string `json:"GpuType,omitempty" name:"GpuType"`

	// 配置名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

type JobStatus struct {

	// 任务状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误时为错误描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 预期Worker数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DesiredWorkers *uint64 `json:"DesiredWorkers,omitempty" name:"DesiredWorkers"`

	// 当前Worker数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentWorkers *uint64 `json:"CurrentWorkers,omitempty" name:"CurrentWorkers"`

	// 副本名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas []*string `json:"Replicas,omitempty" name:"Replicas" list`
}

type ModelService struct {

	// 服务ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 运行集群
	// 注意：此字段可能返回 null，表示取不到有效值。
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

	// GPU 配置, 单位为1/1000 卡
	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`

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

	// 访问密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessToken *string `json:"AccessToken,omitempty" name:"AccessToken"`

	// 服务配置Id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 服务配置名
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 服务运行时长
	ServeSeconds *uint64 `json:"ServeSeconds,omitempty" name:"ServeSeconds"`

	// 配置版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 服务使用资源组 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`

	// 暴露方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exposes []*Expose `json:"Exposes,omitempty" name:"Exposes" list`

	// Region 名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 服务使用资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" name:"ResourceGroupName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// GPU类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuType *string `json:"GpuType,omitempty" name:"GpuType"`
}

type Option struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 取值
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type PredictInput struct {

	// 输入路径，支持 cos 格式路径文件夹或文件
	InputPath *string `json:"InputPath,omitempty" name:"InputPath"`

	// 输出路径，支持 cos 格式路径
	OutputPath *string `json:"OutputPath,omitempty" name:"OutputPath"`

	// 输入数据格式，目前支持：JSON
	InputDataFormat *string `json:"InputDataFormat,omitempty" name:"InputDataFormat"`

	// 输出数据格式，目前支持：JSON
	OutputDataFormat *string `json:"OutputDataFormat,omitempty" name:"OutputDataFormat"`

	// 预测批大小，默认为 64
	BatchSize *uint64 `json:"BatchSize,omitempty" name:"BatchSize"`

	// 模型签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignatureName *string `json:"SignatureName,omitempty" name:"SignatureName"`
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

	// 运行状态对额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type UpdateJobRequest struct {
	*tchttp.BaseRequest

	// 任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 任务更新动作，支持：Cancel
	JobAction *string `json:"JobAction,omitempty" name:"JobAction"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务
	// 注意：此字段可能返回 null，表示取不到有效值。
		Job *Job `json:"Job,omitempty" name:"Job"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceRequest struct {
	*tchttp.BaseRequest

	// 服务Id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 扩缩容配置
	Scaler *Scaler `json:"Scaler,omitempty" name:"Scaler"`

	// 服务配置Id
	ServiceConfigId *string `json:"ServiceConfigId,omitempty" name:"ServiceConfigId"`

	// 支持AUTO, MANUAL，分别表示自动扩缩容，手动扩缩容
	ScaleMode *string `json:"ScaleMode,omitempty" name:"ScaleMode"`

	// 支持STOP(停止) RESUME(重启)
	ServiceAction *string `json:"ServiceAction,omitempty" name:"ServiceAction"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`

	// GPU卡类型
	GpuType *string `json:"GpuType,omitempty" name:"GpuType"`
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
		Service *ModelService `json:"Service,omitempty" name:"Service"`

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
