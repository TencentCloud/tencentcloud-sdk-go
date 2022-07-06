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

package v20210820

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeProjectRequestParams struct {
	// 项目id。一般使用项目Id来查询，与projectName必须存在一个。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 是否展示关联集群信息
	DescribeClusters *bool `json:"DescribeClusters,omitempty" name:"DescribeClusters"`

	// 是否展示关联执行组的信息，仅部分信息。
	DescribeExecutors *bool `json:"DescribeExecutors,omitempty" name:"DescribeExecutors"`

	// 默认不展示项目管理员信息
	DescribeAdminUsers *bool `json:"DescribeAdminUsers,omitempty" name:"DescribeAdminUsers"`

	// 默认不统计项目人员数量
	DescribeMemberCount *bool `json:"DescribeMemberCount,omitempty" name:"DescribeMemberCount"`

	// 默认不查询创建者的信息
	DescribeCreator *bool `json:"DescribeCreator,omitempty" name:"DescribeCreator"`

	// 项目名只在租户内唯一，一般用来转化为项目ID。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

type DescribeProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目id。一般使用项目Id来查询，与projectName必须存在一个。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 是否展示关联集群信息
	DescribeClusters *bool `json:"DescribeClusters,omitempty" name:"DescribeClusters"`

	// 是否展示关联执行组的信息，仅部分信息。
	DescribeExecutors *bool `json:"DescribeExecutors,omitempty" name:"DescribeExecutors"`

	// 默认不展示项目管理员信息
	DescribeAdminUsers *bool `json:"DescribeAdminUsers,omitempty" name:"DescribeAdminUsers"`

	// 默认不统计项目人员数量
	DescribeMemberCount *bool `json:"DescribeMemberCount,omitempty" name:"DescribeMemberCount"`

	// 默认不查询创建者的信息
	DescribeCreator *bool `json:"DescribeCreator,omitempty" name:"DescribeCreator"`

	// 项目名只在租户内唯一，一般用来转化为项目ID。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *DescribeProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DescribeClusters")
	delete(f, "DescribeExecutors")
	delete(f, "DescribeAdminUsers")
	delete(f, "DescribeMemberCount")
	delete(f, "DescribeCreator")
	delete(f, "ProjectName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectResponseParams `json:"Response"`
}

func (r *DescribeProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelatedInstancesRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 数据时间，格式yyyy-MM-dd HH:mm:ss
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 任务id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 距离当前任务的层级距离，-1表示取父节点，1表示子节点
	Depth *int64 `json:"Depth,omitempty" name:"Depth"`

	// 页号，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认为10，最大不超过200
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeRelatedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 数据时间，格式yyyy-MM-dd HH:mm:ss
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 任务id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 距离当前任务的层级距离，-1表示取父节点，1表示子节点
	Depth *int64 `json:"Depth,omitempty" name:"Depth"`

	// 页号，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认为10，最大不超过200
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeRelatedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelatedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CurRunDate")
	delete(f, "TaskId")
	delete(f, "Depth")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRelatedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelatedInstancesResponseParams struct {
	// 无
	Data *DescribeTaskInstancesData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRelatedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRelatedInstancesResponseParams `json:"Response"`
}

func (r *DescribeRelatedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelatedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskInstancesData struct {
	// 实例列表
	Items []*TaskInstanceInfo `json:"Items,omitempty" name:"Items"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 页号
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

// Predefined struct for user
type DescribeTaskInstancesRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 页号，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认为10，最大不超过200
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 工作流id列表
	WorkflowIdList []*string `json:"WorkflowIdList,omitempty" name:"WorkflowIdList"`

	// 工作流名称列表，支持模糊搜索
	WorkflowNameList []*string `json:"WorkflowNameList,omitempty" name:"WorkflowNameList"`

	// 起始数据时间，格式yyyy-MM-dd HH:mm:ss
	DateFrom *string `json:"DateFrom,omitempty" name:"DateFrom"`

	// 结束数据时间，格式yyyy-MM-dd HH:mm:ss
	DateTo *string `json:"DateTo,omitempty" name:"DateTo"`

	// 任务id列表
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 任务名称列表，支持模糊搜索
	TaskNameList []*string `json:"TaskNameList,omitempty" name:"TaskNameList"`

	// 责任人名称列表
	InChargeList []*string `json:"InChargeList,omitempty" name:"InChargeList"`

	// 任务类型码列表，26离线同步，30Python，31PySpark，32DLC，33Impala，34Hive SQL，35Shell，36Spark SQL，39Spark，40CDW PG，92MapReduce
	TaskTypeIdList []*int64 `json:"TaskTypeIdList,omitempty" name:"TaskTypeIdList"`

	// 实例状态列表，0等待事件，1等待上游，2等待运行，3运行中，4正在终止，5失败重试，6失败，7成功
	StateList []*string `json:"StateList,omitempty" name:"StateList"`

	// 周期类型列表，I分钟，H小时，D天，W周，M月，Y年，O一次性，C crontab
	TaskCycleUnitList []*string `json:"TaskCycleUnitList,omitempty" name:"TaskCycleUnitList"`

	// 实例类型，0补录实例，1周期实例，2非周期实例
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// 排序字段信息列表，ScheduleDateTime / CostTime / StartTime / EndTime
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`
}

type DescribeTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 页号，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认为10，最大不超过200
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 工作流id列表
	WorkflowIdList []*string `json:"WorkflowIdList,omitempty" name:"WorkflowIdList"`

	// 工作流名称列表，支持模糊搜索
	WorkflowNameList []*string `json:"WorkflowNameList,omitempty" name:"WorkflowNameList"`

	// 起始数据时间，格式yyyy-MM-dd HH:mm:ss
	DateFrom *string `json:"DateFrom,omitempty" name:"DateFrom"`

	// 结束数据时间，格式yyyy-MM-dd HH:mm:ss
	DateTo *string `json:"DateTo,omitempty" name:"DateTo"`

	// 任务id列表
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 任务名称列表，支持模糊搜索
	TaskNameList []*string `json:"TaskNameList,omitempty" name:"TaskNameList"`

	// 责任人名称列表
	InChargeList []*string `json:"InChargeList,omitempty" name:"InChargeList"`

	// 任务类型码列表，26离线同步，30Python，31PySpark，32DLC，33Impala，34Hive SQL，35Shell，36Spark SQL，39Spark，40CDW PG，92MapReduce
	TaskTypeIdList []*int64 `json:"TaskTypeIdList,omitempty" name:"TaskTypeIdList"`

	// 实例状态列表，0等待事件，1等待上游，2等待运行，3运行中，4正在终止，5失败重试，6失败，7成功
	StateList []*string `json:"StateList,omitempty" name:"StateList"`

	// 周期类型列表，I分钟，H小时，D天，W周，M月，Y年，O一次性，C crontab
	TaskCycleUnitList []*string `json:"TaskCycleUnitList,omitempty" name:"TaskCycleUnitList"`

	// 实例类型，0补录实例，1周期实例，2非周期实例
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// 排序字段信息列表，ScheduleDateTime / CostTime / StartTime / EndTime
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`
}

func (r *DescribeTaskInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "WorkflowIdList")
	delete(f, "WorkflowNameList")
	delete(f, "DateFrom")
	delete(f, "DateTo")
	delete(f, "TaskIdList")
	delete(f, "TaskNameList")
	delete(f, "InChargeList")
	delete(f, "TaskTypeIdList")
	delete(f, "StateList")
	delete(f, "TaskCycleUnitList")
	delete(f, "InstanceType")
	delete(f, "OrderFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInstancesResponseParams struct {
	// 无
	Data *DescribeTaskInstancesData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskInstancesResponseParams `json:"Response"`
}

func (r *DescribeTaskInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrderField struct {
	// 排序字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排序方向：ASC|DESC
	Direction *string `json:"Direction,omitempty" name:"Direction"`
}

type TaskInstanceInfo struct {
	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectIdent *string `json:"ProjectIdent,omitempty" name:"ProjectIdent"`

	// 实例状态，0等待事件，1等待上游，2等待运行，3运行中，4正在终止，5失败重试，6失败，7成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitempty" name:"State"`

	// 任务类型id，26离线同步，30Python，31PySpark，32DLC，33Impala，34Hive SQL，35Shell，36Spark SQL，39Spark，40CDW PG，92MapReduce
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeId *uint64 `json:"TaskTypeId,omitempty" name:"TaskTypeId"`

	// 任务类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeDesc *string `json:"TaskTypeDesc,omitempty" name:"TaskTypeDesc"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 调度计划展示描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDesc *string `json:"SchedulerDesc,omitempty" name:"SchedulerDesc"`

	// 负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitempty" name:"InCharge"`

	// 调度周期类型，I分钟，H小时，D天，W周，M月，Y年，O一次性，C crontab
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitempty" name:"CycleType"`

	// 实例开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 实例结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 实例类型，0补录实例，1周期实例，2非周期实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// 最大重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TryLimit *int64 `json:"TryLimit,omitempty" name:"TryLimit"`

	// 当前重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tries *int64 `json:"Tries,omitempty" name:"Tries"`

	// 计划调度时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDateTime *string `json:"SchedulerDateTime,omitempty" name:"SchedulerDateTime"`

	// 运行耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *string `json:"CostTime,omitempty" name:"CostTime"`
}