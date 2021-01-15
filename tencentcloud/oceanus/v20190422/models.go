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

package v20190422

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateJobConfigRequest struct {
	*tchttp.BaseRequest

	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 主类
	EntrypointClass *string `json:"EntrypointClass,omitempty" name:"EntrypointClass"`

	// 主类入参
	ProgramArgs *string `json:"ProgramArgs,omitempty" name:"ProgramArgs"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 资源引用数组
	ResourceRefs []*ResourceRef `json:"ResourceRefs,omitempty" name:"ResourceRefs" list`

	// 作业默认并行度
	DefaultParallelism *uint64 `json:"DefaultParallelism,omitempty" name:"DefaultParallelism"`

	// 系统参数
	Properties []*Property `json:"Properties,omitempty" name:"Properties" list`

	// 1: 作业配置达到上限之后，自动删除可删除的最早版本
	AutoDelete *int64 `json:"AutoDelete,omitempty" name:"AutoDelete"`

	// 作业使用的 COS 存储桶名
	COSBucket *string `json:"COSBucket,omitempty" name:"COSBucket"`
}

func (r *CreateJobConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateJobConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateJobConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 作业配置版本号
		Version *uint64 `json:"Version,omitempty" name:"Version"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateJobConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateJobConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateJobRequest struct {
	*tchttp.BaseRequest

	// 作业名称，允许输入长度小于等于50个字符的中文、英文、数字、-（横线）、_（下划线）、.（点），且符号必须半角字符。注意作业名不能和现有作业同名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 作业的类型，1 表示 SQL 作业，2 表示 JAR 作业
	JobType *int64 `json:"JobType,omitempty" name:"JobType"`

	// 集群的类型，1 表示共享集群，2 表示独享集群
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// 当 ClusterType=2 时，必选，用来指定该作业提交的独享集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 设置每 CU 的内存规格，单位为 GB，支持 2、4、8、16（需申请开通白名单后使用）。默认为 4，即 1 CU 对应 4 GB 的运行内存
	CuMem *uint64 `json:"CuMem,omitempty" name:"CuMem"`

	// 作业的备注信息，可以随意设置
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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

		// 作业Id
		JobId *string `json:"JobId,omitempty" name:"JobId"`

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

type CreateResourceConfigRequest struct {
	*tchttp.BaseRequest

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 位置信息
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitempty" name:"ResourceLoc"`

	// 资源描述信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 1： 资源版本达到上限，自动删除最早可删除的版本
	AutoDelete *int64 `json:"AutoDelete,omitempty" name:"AutoDelete"`
}

func (r *CreateResourceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateResourceConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateResourceConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资源版本ID
		Version *int64 `json:"Version,omitempty" name:"Version"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateResourceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateResourceConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateResourceRequest struct {
	*tchttp.BaseRequest

	// 资源位置
	ResourceLoc *ResourceLoc `json:"ResourceLoc,omitempty" name:"ResourceLoc"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资源类型，占时只支持jar，填1
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源描述
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 资源版本描述
	ResourceConfigRemark *string `json:"ResourceConfigRemark,omitempty" name:"ResourceConfigRemark"`
}

func (r *CreateResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资源ID
		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

		// 资源版本
		Version *int64 `json:"Version,omitempty" name:"Version"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTableConfigRequest struct {
	*tchttp.BaseRequest

	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 调试作业ID
	DebugId *int64 `json:"DebugId,omitempty" name:"DebugId"`

	// 表名
	TableName *string `json:"TableName,omitempty" name:"TableName"`
}

func (r *DeleteTableConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTableConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTableConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTableConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTableConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeJobsRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个作业ID查询。作业ID形如：cql-11112222，每次请求的作业上限为100。参数不支持同时指定JobIds和Filters。
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds" list`

	// 过滤条件，详见作业过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定JobIds和Filters。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeJobsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 作业总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 作业列表
		JobSet []*JobV1 `json:"JobSet,omitempty" name:"JobSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeJobsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemResourcesRequest struct {
	*tchttp.BaseRequest

	// 需要查询的资源ID数组
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 条数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询资源配置列表， 如果不填写，返回该ResourceId下所有作业配置列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeSystemResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSystemResourcesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSystemResourcesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资源详细信息集合
		ResourceSet []*SystemResourceItem `json:"ResourceSet,omitempty" name:"ResourceSet" list`

		// 总数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSystemResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSystemResourcesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 要过滤的字段
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type JobV1 struct {

	// 作业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 用户AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 用户UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 创建者UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 作业名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 作业类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobType *int64 `json:"JobType,omitempty" name:"JobType"`

	// 作业状态，1：未初始化，2：未发布，3：操作中，4：运行中，5：停止，6：暂停，-1：故障
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 作业创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 作业启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 作业停止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StopTime *string `json:"StopTime,omitempty" name:"StopTime"`

	// 作业更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 作业累计运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalRunMillis *int64 `json:"TotalRunMillis,omitempty" name:"TotalRunMillis"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 操作错误提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOpResult *string `json:"LastOpResult,omitempty" name:"LastOpResult"`

	// 集群名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 最新配置版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestJobConfigVersion *int64 `json:"LatestJobConfigVersion,omitempty" name:"LatestJobConfigVersion"`

	// 已发布的配置版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishedJobConfigVersion *int64 `json:"PublishedJobConfigVersion,omitempty" name:"PublishedJobConfigVersion"`

	// 运行的CU数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningCuNum *int64 `json:"RunningCuNum,omitempty" name:"RunningCuNum"`

	// 作业内存规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	CuMem *int64 `json:"CuMem,omitempty" name:"CuMem"`

	// 作业状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 运行状态时表示单次运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentRunMillis *int64 `json:"CurrentRunMillis,omitempty" name:"CurrentRunMillis"`

	// 作业所在的集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 作业管理WEB UI 入口
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebUIUrl *string `json:"WebUIUrl,omitempty" name:"WebUIUrl"`

	// 作业所在集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerType *int64 `json:"SchedulerType,omitempty" name:"SchedulerType"`
}

type Property struct {

	// 系统配置的Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 系统配置的Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ResourceLoc struct {

	// 资源位置的存储类型，目前只支持COS
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`

	// 描述资源位置的json
	Param *ResourceLocParam `json:"Param,omitempty" name:"Param"`
}

type ResourceLocParam struct {

	// 资源bucket
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 资源路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 资源所在地域，如果不填，则使用Resource的Region
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type ResourceRef struct {

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源版本ID，-1表示使用最新版本
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// 引用资源类型，例如主资源设置为1，代表main class所在的jar包
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type RunJobDescription struct {

	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 运行类型，1：启动，2：恢复
	RunType *int64 `json:"RunType,omitempty" name:"RunType"`

	// SQL类型作业启动参数：指定数据源消费起始时间点
	StartMode *string `json:"StartMode,omitempty" name:"StartMode"`

	// 已发布上线的作业配置版本
	JobConfigVersion *uint64 `json:"JobConfigVersion,omitempty" name:"JobConfigVersion"`
}

type RunJobsRequest struct {
	*tchttp.BaseRequest

	// 批量启动作业的描述信息
	RunJobDescriptions []*RunJobDescription `json:"RunJobDescriptions,omitempty" name:"RunJobDescriptions" list`
}

func (r *RunJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunJobsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunJobsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopJobDescription struct {

	// 作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 停止类型，1 停止 2 暂停
	StopType *int64 `json:"StopType,omitempty" name:"StopType"`
}

type StopJobsRequest struct {
	*tchttp.BaseRequest

	// 批量停止作业的描述信息
	StopJobDescriptions []*StopJobDescription `json:"StopJobDescriptions,omitempty" name:"StopJobDescriptions" list`
}

func (r *StopJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopJobsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopJobsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemResourceItem struct {

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资源类型
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 资源所属地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 资源的最新版本
	LatestResourceConfigVersion *int64 `json:"LatestResourceConfigVersion,omitempty" name:"LatestResourceConfigVersion"`
}
