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
type CreateTaskFromTemplateRequestParams struct {
	// 从经验库中查询到的经验模板ID
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 演练的配置参数
	TaskConfig *TaskConfig `json:"TaskConfig,omitempty" name:"TaskConfig"`
}

type CreateTaskFromTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 从经验库中查询到的经验模板ID
	TemplateId *uint64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 演练的配置参数
	TaskConfig *TaskConfig `json:"TaskConfig,omitempty" name:"TaskConfig"`
}

func (r *CreateTaskFromTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFromTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TaskConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskFromTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFromTemplateResponseParams struct {
	// 创建成功的演练ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTaskFromTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskFromTemplateResponseParams `json:"Response"`
}

func (r *CreateTaskFromTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFromTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskRequestParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DeleteTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTaskResponseParams `json:"Response"`
}

func (r *DeleteTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicy struct {
	// 保护策略ID列表
	TaskPolicyIdList []*string `json:"TaskPolicyIdList,omitempty" name:"TaskPolicyIdList"`

	// 保护策略状态
	TaskPolicyStatus *string `json:"TaskPolicyStatus,omitempty" name:"TaskPolicyStatus"`

	// 策略规则
	TaskPolicyRule *string `json:"TaskPolicyRule,omitempty" name:"TaskPolicyRule"`

	// 护栏策略生效处理策略 1:顺序执行，2:暂停
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskPolicyDealType *int64 `json:"TaskPolicyDealType,omitempty" name:"TaskPolicyDealType"`
}

// Predefined struct for user
type DescribeTaskExecuteLogsRequestParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 返回的内容行数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 日志起始的行数。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeTaskExecuteLogsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 返回的内容行数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 日志起始的行数。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTaskExecuteLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskExecuteLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskExecuteLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskExecuteLogsResponseParams struct {
	// 日志数据
	LogMessage []*string `json:"LogMessage,omitempty" name:"LogMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskExecuteLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskExecuteLogsResponseParams `json:"Response"`
}

func (r *DescribeTaskExecuteLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskExecuteLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskListRequestParams struct {
	// 分页Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 演练名称
	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`

	// 标签键
	TaskTag []*string `json:"TaskTag,omitempty" name:"TaskTag"`

	// 状态
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 开始时间，固定格式%Y-%m-%d %H:%M:%S
	TaskStartTime *string `json:"TaskStartTime,omitempty" name:"TaskStartTime"`

	// 结束时间，固定格式%Y-%m-%d %H:%M:%S
	TaskEndTime *string `json:"TaskEndTime,omitempty" name:"TaskEndTime"`

	// 标签对
	Tags []*TagWithDescribe `json:"Tags,omitempty" name:"Tags"`
}

type DescribeTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 分页Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 演练名称
	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`

	// 标签键
	TaskTag []*string `json:"TaskTag,omitempty" name:"TaskTag"`

	// 状态
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 开始时间，固定格式%Y-%m-%d %H:%M:%S
	TaskStartTime *string `json:"TaskStartTime,omitempty" name:"TaskStartTime"`

	// 结束时间，固定格式%Y-%m-%d %H:%M:%S
	TaskEndTime *string `json:"TaskEndTime,omitempty" name:"TaskEndTime"`

	// 标签对
	Tags []*TagWithDescribe `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "TaskTitle")
	delete(f, "TaskTag")
	delete(f, "TaskStatus")
	delete(f, "TaskStartTime")
	delete(f, "TaskEndTime")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskListResponseParams struct {
	// 无
	TaskList []*TaskListItem `json:"TaskList,omitempty" name:"TaskList"`

	// 列表数量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskListResponseParams `json:"Response"`
}

func (r *DescribeTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskRequestParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResponseParams struct {
	// 任务信息
	Task *Task `json:"Task,omitempty" name:"Task"`

	// 任务对应的演练报告信息，null表示未导出报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportInfo *TaskReportInfo `json:"ReportInfo,omitempty" name:"ReportInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResponseParams `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateListRequestParams struct {
	// 分页Limit, 最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 演练名称
	Title *string `json:"Title,omitempty" name:"Title"`

	// 标签键
	Tag []*string `json:"Tag,omitempty" name:"Tag"`

	// 状态，1---使用中， 2---停用
	IsUsed *int64 `json:"IsUsed,omitempty" name:"IsUsed"`

	// 标签对
	Tags []*TagWithDescribe `json:"Tags,omitempty" name:"Tags"`
}

type DescribeTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// 分页Limit, 最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 演练名称
	Title *string `json:"Title,omitempty" name:"Title"`

	// 标签键
	Tag []*string `json:"Tag,omitempty" name:"Tag"`

	// 状态，1---使用中， 2---停用
	IsUsed *int64 `json:"IsUsed,omitempty" name:"IsUsed"`

	// 标签对
	Tags []*TagWithDescribe `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Title")
	delete(f, "Tag")
	delete(f, "IsUsed")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateListResponseParams struct {
	// 经验库列表
	TemplateList []*TemplateListItem `json:"TemplateList,omitempty" name:"TemplateList"`

	// 列表数量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTemplateListResponseParams `json:"Response"`
}

func (r *DescribeTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateRequestParams struct {
	// 经验库ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 经验库ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateResponseParams struct {
	// 经验库详情
	Template *Template `json:"Template,omitempty" name:"Template"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTemplateResponseParams `json:"Response"`
}

func (r *DescribeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteTaskInstanceRequestParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务动作ID
	TaskActionId *uint64 `json:"TaskActionId,omitempty" name:"TaskActionId"`

	// 任务动作实例ID
	TaskInstanceIds []*uint64 `json:"TaskInstanceIds,omitempty" name:"TaskInstanceIds"`

	// 是否操作整个任务
	IsOperateAll *bool `json:"IsOperateAll,omitempty" name:"IsOperateAll"`

	// 操作类型：（1--启动   2--执行  3--跳过   5--重试）
	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`

	// 动作组ID
	TaskGroupId *uint64 `json:"TaskGroupId,omitempty" name:"TaskGroupId"`
}

type ExecuteTaskInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务动作ID
	TaskActionId *uint64 `json:"TaskActionId,omitempty" name:"TaskActionId"`

	// 任务动作实例ID
	TaskInstanceIds []*uint64 `json:"TaskInstanceIds,omitempty" name:"TaskInstanceIds"`

	// 是否操作整个任务
	IsOperateAll *bool `json:"IsOperateAll,omitempty" name:"IsOperateAll"`

	// 操作类型：（1--启动   2--执行  3--跳过   5--重试）
	ActionType *uint64 `json:"ActionType,omitempty" name:"ActionType"`

	// 动作组ID
	TaskGroupId *uint64 `json:"TaskGroupId,omitempty" name:"TaskGroupId"`
}

func (r *ExecuteTaskInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteTaskInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TaskActionId")
	delete(f, "TaskInstanceIds")
	delete(f, "IsOperateAll")
	delete(f, "ActionType")
	delete(f, "TaskGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteTaskInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteTaskInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExecuteTaskInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteTaskInstanceResponseParams `json:"Response"`
}

func (r *ExecuteTaskInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteTaskInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteTaskRequestParams struct {
	// 需要执行的任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type ExecuteTaskRequest struct {
	*tchttp.BaseRequest
	
	// 需要执行的任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ExecuteTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExecuteTaskResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteTaskResponseParams `json:"Response"`
}

func (r *ExecuteTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskRunStatusRequestParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态, 1001--未开始 1002--进行中（执行）1003--进行中（暂停）1004--执行结束
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 执行结果是否符合预期（当前扭转状态为执行结束时，需要必传此字段）
	IsExpect *bool `json:"IsExpect,omitempty" name:"IsExpect"`

	// 演习结论（当演习状态转变为执行结束时，需要填写此字段）
	Summary *string `json:"Summary,omitempty" name:"Summary"`
}

type ModifyTaskRunStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态, 1001--未开始 1002--进行中（执行）1003--进行中（暂停）1004--执行结束
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 执行结果是否符合预期（当前扭转状态为执行结束时，需要必传此字段）
	IsExpect *bool `json:"IsExpect,omitempty" name:"IsExpect"`

	// 演习结论（当演习状态转变为执行结束时，需要填写此字段）
	Summary *string `json:"Summary,omitempty" name:"Summary"`
}

func (r *ModifyTaskRunStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskRunStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Status")
	delete(f, "IsExpect")
	delete(f, "Summary")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskRunStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskRunStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTaskRunStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskRunStatusResponseParams `json:"Response"`
}

func (r *ModifyTaskRunStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskRunStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagWithCreate struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagWithDescribe struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Task struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务标题
	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`

	// 任务描述
	TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`

	// 自定义标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTag *string `json:"TaskTag,omitempty" name:"TaskTag"`

	// 任务状态，1001--未开始  1002--进行中（执行）1003--进行中（暂停）1004--执行结束
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 任务结束状态，表明任务以何种状态结束: 0 -- 尚未结束，1 -- 成功，2-- 失败，3--终止
	TaskStatusType *int64 `json:"TaskStatusType,omitempty" name:"TaskStatusType"`

	// 保护策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskProtectStrategy *string `json:"TaskProtectStrategy,omitempty" name:"TaskProtectStrategy"`

	// 任务创建时间
	TaskCreateTime *string `json:"TaskCreateTime,omitempty" name:"TaskCreateTime"`

	// 任务更新时间
	TaskUpdateTime *string `json:"TaskUpdateTime,omitempty" name:"TaskUpdateTime"`

	// 任务动作组
	TaskGroups []*TaskGroup `json:"TaskGroups,omitempty" name:"TaskGroups"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStartTime *string `json:"TaskStartTime,omitempty" name:"TaskStartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskEndTime *string `json:"TaskEndTime,omitempty" name:"TaskEndTime"`

	// 是否符合预期。1：符合预期，2：不符合预期
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskExpect *int64 `json:"TaskExpect,omitempty" name:"TaskExpect"`

	// 演习记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskSummary *string `json:"TaskSummary,omitempty" name:"TaskSummary"`

	// 任务模式。1:手工执行，2:自动执行
	TaskMode *int64 `json:"TaskMode,omitempty" name:"TaskMode"`

	// 自动暂停时长。单位分钟
	TaskPauseDuration *int64 `json:"TaskPauseDuration,omitempty" name:"TaskPauseDuration"`

	// 演练创建者Uin
	TaskOwnerUin *string `json:"TaskOwnerUin,omitempty" name:"TaskOwnerUin"`

	// 地域ID
	TaskRegionId *int64 `json:"TaskRegionId,omitempty" name:"TaskRegionId"`

	// 监控指标列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskMonitors []*TaskMonitor `json:"TaskMonitors,omitempty" name:"TaskMonitors"`

	// 保护策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskPolicy *DescribePolicy `json:"TaskPolicy,omitempty" name:"TaskPolicy"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagWithDescribe `json:"Tags,omitempty" name:"Tags"`
}

type TaskConfig struct {
	// 动作组配置，需要保证配置个数和经验中的动作组个数一致
	TaskGroupsConfig []*TaskGroupConfig `json:"TaskGroupsConfig,omitempty" name:"TaskGroupsConfig"`

	// 更改后的演练名称，不填则默认取经验名称
	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`

	// 更改后的演练描述，不填则默认取经验描述
	TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`

	// 演练执行模式：1----手工执行/ 2 ---自动执行，不填则默认取经验执行模式
	TaskMode *uint64 `json:"TaskMode,omitempty" name:"TaskMode"`

	// 演练自动暂停时间，单位分钟, 不填则默认取经验自动暂停时间
	TaskPauseDuration *uint64 `json:"TaskPauseDuration,omitempty" name:"TaskPauseDuration"`

	// 演练标签信息，不填则默认取经验标签
	Tags []*TagWithCreate `json:"Tags,omitempty" name:"Tags"`
}

type TaskGroup struct {
	// 任务动作ID
	TaskGroupId *int64 `json:"TaskGroupId,omitempty" name:"TaskGroupId"`

	// 分组标题
	TaskGroupTitle *string `json:"TaskGroupTitle,omitempty" name:"TaskGroupTitle"`

	// 分组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupDescription *string `json:"TaskGroupDescription,omitempty" name:"TaskGroupDescription"`

	// 任务分组顺序
	TaskGroupOrder *int64 `json:"TaskGroupOrder,omitempty" name:"TaskGroupOrder"`

	// 对象类型ID
	ObjectTypeId *int64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`

	// 任务分组创建时间
	TaskGroupCreateTime *string `json:"TaskGroupCreateTime,omitempty" name:"TaskGroupCreateTime"`

	// 任务分组更新时间
	TaskGroupUpdateTime *string `json:"TaskGroupUpdateTime,omitempty" name:"TaskGroupUpdateTime"`

	// 动作分组动作列表
	TaskGroupActions []*TaskGroupAction `json:"TaskGroupActions,omitempty" name:"TaskGroupActions"`

	// 实例列表
	TaskGroupInstanceList []*string `json:"TaskGroupInstanceList,omitempty" name:"TaskGroupInstanceList"`

	// 执行模式。1 --- 顺序执行，2 --- 阶段执行
	TaskGroupMode *int64 `json:"TaskGroupMode,omitempty" name:"TaskGroupMode"`
}

type TaskGroupAction struct {
	// 任务分组动作ID
	TaskGroupActionId *int64 `json:"TaskGroupActionId,omitempty" name:"TaskGroupActionId"`

	// 任务分组动作实例列表
	TaskGroupInstances []*TaskGroupInstance `json:"TaskGroupInstances,omitempty" name:"TaskGroupInstances"`

	// 动作ID
	ActionId *int64 `json:"ActionId,omitempty" name:"ActionId"`

	// 分组动作顺序
	TaskGroupActionOrder *int64 `json:"TaskGroupActionOrder,omitempty" name:"TaskGroupActionOrder"`

	// 分组动作通用配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupActionGeneralConfiguration *string `json:"TaskGroupActionGeneralConfiguration,omitempty" name:"TaskGroupActionGeneralConfiguration"`

	// 分组动作自定义配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupActionCustomConfiguration *string `json:"TaskGroupActionCustomConfiguration,omitempty" name:"TaskGroupActionCustomConfiguration"`

	// 分组动作状态
	TaskGroupActionStatus *int64 `json:"TaskGroupActionStatus,omitempty" name:"TaskGroupActionStatus"`

	// 动作分组创建时间
	TaskGroupActionCreateTime *string `json:"TaskGroupActionCreateTime,omitempty" name:"TaskGroupActionCreateTime"`

	// 动作分组更新时间
	TaskGroupActionUpdateTime *string `json:"TaskGroupActionUpdateTime,omitempty" name:"TaskGroupActionUpdateTime"`

	// 动作名称
	ActionTitle *string `json:"ActionTitle,omitempty" name:"ActionTitle"`

	// 状态类型: 0 -- 无状态，1 -- 成功，2-- 失败，3--终止，4--跳过
	TaskGroupActionStatusType *int64 `json:"TaskGroupActionStatusType,omitempty" name:"TaskGroupActionStatusType"`

	// RandomId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupActionRandomId *int64 `json:"TaskGroupActionRandomId,omitempty" name:"TaskGroupActionRandomId"`

	// RecoverId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupActionRecoverId *int64 `json:"TaskGroupActionRecoverId,omitempty" name:"TaskGroupActionRecoverId"`

	// ExecuteId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupActionExecuteId *int64 `json:"TaskGroupActionExecuteId,omitempty" name:"TaskGroupActionExecuteId"`

	// 调用api类型，0:tat, 1:云api
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionApiType *int64 `json:"ActionApiType,omitempty" name:"ActionApiType"`

	// 1:故障，2:恢复
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionAttribute *int64 `json:"ActionAttribute,omitempty" name:"ActionAttribute"`

	// 动作类型：平台、自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 是否可重试
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsExecuteRedo *bool `json:"IsExecuteRedo,omitempty" name:"IsExecuteRedo"`

	// 动作风险级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionRisk *string `json:"ActionRisk,omitempty" name:"ActionRisk"`

	// 动作运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupActionExecuteTime *int64 `json:"TaskGroupActionExecuteTime,omitempty" name:"TaskGroupActionExecuteTime"`
}

type TaskGroupActionConfig struct {
	// 该动作在动作组中的顺序，从1开始，不填或填错将匹配不到经验中要修改参数的动作
	TaskGroupActionOrder *uint64 `json:"TaskGroupActionOrder,omitempty" name:"TaskGroupActionOrder"`

	// 动作通用参数，需要json序列化传入，可以从查询经验详情接口获取，不填默认使用经验中动作参数
	TaskGroupActionGeneralConfiguration *string `json:"TaskGroupActionGeneralConfiguration,omitempty" name:"TaskGroupActionGeneralConfiguration"`

	// 动作自定义参数，需要json序列化传入，可以从查询经验详情接口获取，不填默认使用经验中动作参数
	TaskGroupActionCustomConfiguration *string `json:"TaskGroupActionCustomConfiguration,omitempty" name:"TaskGroupActionCustomConfiguration"`
}

type TaskGroupConfig struct {
	// 动作组所关联的实例对象
	TaskGroupInstances []*string `json:"TaskGroupInstances,omitempty" name:"TaskGroupInstances"`

	// 动作组标题，不填默认取经验中的动作组名称
	TaskGroupTitle *string `json:"TaskGroupTitle,omitempty" name:"TaskGroupTitle"`

	// 动作组描述，不填默认取经验中的动作组描述
	TaskGroupDescription *string `json:"TaskGroupDescription,omitempty" name:"TaskGroupDescription"`

	// 动作执行模式。1 --- 顺序执行，2 --- 阶段执行, 不填默认取经验中的动作组执行模式
	TaskGroupMode *uint64 `json:"TaskGroupMode,omitempty" name:"TaskGroupMode"`

	// 动作组中的动作参数，不填默认使用经验中的动作参数，配置时可以只指定想要修改参数的动作
	TaskGroupActionsConfig []*TaskGroupActionConfig `json:"TaskGroupActionsConfig,omitempty" name:"TaskGroupActionsConfig"`
}

type TaskGroupInstance struct {
	// 实例ID
	TaskGroupInstanceId *int64 `json:"TaskGroupInstanceId,omitempty" name:"TaskGroupInstanceId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupInstanceObjectId *string `json:"TaskGroupInstanceObjectId,omitempty" name:"TaskGroupInstanceObjectId"`

	// 实例动作执行状态
	TaskGroupInstanceStatus *int64 `json:"TaskGroupInstanceStatus,omitempty" name:"TaskGroupInstanceStatus"`

	// 实例动作执行日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupInstanceExecuteLog *string `json:"TaskGroupInstanceExecuteLog,omitempty" name:"TaskGroupInstanceExecuteLog"`

	// 实例创建时间
	TaskGroupInstanceCreateTime *string `json:"TaskGroupInstanceCreateTime,omitempty" name:"TaskGroupInstanceCreateTime"`

	// 实例更新时间
	TaskGroupInstanceUpdateTime *string `json:"TaskGroupInstanceUpdateTime,omitempty" name:"TaskGroupInstanceUpdateTime"`

	// 状态类型: 0 -- 无状态，1 -- 成功，2-- 失败，3--终止，4--跳过
	TaskGroupInstanceStatusType *int64 `json:"TaskGroupInstanceStatusType,omitempty" name:"TaskGroupInstanceStatusType"`

	// 执行开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupInstanceStartTime *string `json:"TaskGroupInstanceStartTime,omitempty" name:"TaskGroupInstanceStartTime"`

	// 执行结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupInstanceEndTime *string `json:"TaskGroupInstanceEndTime,omitempty" name:"TaskGroupInstanceEndTime"`

	// 实例是否可重试
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupInstanceIsRedo *bool `json:"TaskGroupInstanceIsRedo,omitempty" name:"TaskGroupInstanceIsRedo"`

	// 动作实例执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupInstanceExecuteTime *int64 `json:"TaskGroupInstanceExecuteTime,omitempty" name:"TaskGroupInstanceExecuteTime"`
}

type TaskListItem struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务标题
	TaskTitle *string `json:"TaskTitle,omitempty" name:"TaskTitle"`

	// 任务描述
	TaskDescription *string `json:"TaskDescription,omitempty" name:"TaskDescription"`

	// 任务标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTag *string `json:"TaskTag,omitempty" name:"TaskTag"`

	// 任务状态(1001 -- 未开始   1002 -- 进行中  1003 -- 暂停中   1004 -- 任务结束)
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 任务创建时间
	TaskCreateTime *string `json:"TaskCreateTime,omitempty" name:"TaskCreateTime"`

	// 任务更新时间
	TaskUpdateTime *string `json:"TaskUpdateTime,omitempty" name:"TaskUpdateTime"`

	// 0--未开始，1--进行中，2--已完成
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskPreCheckStatus *int64 `json:"TaskPreCheckStatus,omitempty" name:"TaskPreCheckStatus"`

	// 环境检查是否通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskPreCheckSuccess *bool `json:"TaskPreCheckSuccess,omitempty" name:"TaskPreCheckSuccess"`
}

type TaskMonitor struct {
	// 监控指标ID
	TaskMonitorId *int64 `json:"TaskMonitorId,omitempty" name:"TaskMonitorId"`

	// 监控指标对象类型ID
	TaskMonitorObjectTypeId *int64 `json:"TaskMonitorObjectTypeId,omitempty" name:"TaskMonitorObjectTypeId"`

	// 指标名称
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 实例ID列表
	InstancesIds []*string `json:"InstancesIds,omitempty" name:"InstancesIds"`

	// 中文指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricChineseName *string `json:"MetricChineseName,omitempty" name:"MetricChineseName"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type TaskReportInfo struct {
	// 0--未开始，1--正在导出，2--导出成功，3--导出失败
	Stage *int64 `json:"Stage,omitempty" name:"Stage"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 有效期截止时间
	ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`

	// 是否有效
	Expired *bool `json:"Expired,omitempty" name:"Expired"`

	// 演练报告cos文件地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosUrl *string `json:"CosUrl,omitempty" name:"CosUrl"`

	// 演练报告导出日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Log *string `json:"Log,omitempty" name:"Log"`
}

type Template struct {
	// 经验库ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 经验库标题
	TemplateTitle *string `json:"TemplateTitle,omitempty" name:"TemplateTitle"`

	// 经验库描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 自定义标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateTag *string `json:"TemplateTag,omitempty" name:"TemplateTag"`

	// 使用状态。1 ---- 使用中，2 --- 停用
	TemplateIsUsed *int64 `json:"TemplateIsUsed,omitempty" name:"TemplateIsUsed"`

	// 经验库创建时间
	TemplateCreateTime *string `json:"TemplateCreateTime,omitempty" name:"TemplateCreateTime"`

	// 经验库更新时间
	TemplateUpdateTime *string `json:"TemplateUpdateTime,omitempty" name:"TemplateUpdateTime"`

	// 经验库模式。1:手工执行，2:自动执行
	TemplateMode *int64 `json:"TemplateMode,omitempty" name:"TemplateMode"`

	// 自动暂停时长。单位分钟
	TemplatePauseDuration *int64 `json:"TemplatePauseDuration,omitempty" name:"TemplatePauseDuration"`

	// 演练创建者Uin
	TemplateOwnerUin *string `json:"TemplateOwnerUin,omitempty" name:"TemplateOwnerUin"`

	// 地域ID
	TemplateRegionId *int64 `json:"TemplateRegionId,omitempty" name:"TemplateRegionId"`

	// 动作组
	TemplateGroups []*TemplateGroup `json:"TemplateGroups,omitempty" name:"TemplateGroups"`

	// 监控指标
	TemplateMonitors []*TemplateMonitor `json:"TemplateMonitors,omitempty" name:"TemplateMonitors"`

	// 护栏监控
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplatePolicy *TemplatePolicy `json:"TemplatePolicy,omitempty" name:"TemplatePolicy"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagWithDescribe `json:"Tags,omitempty" name:"Tags"`
}

type TemplateGroup struct {
	// 经验库动作ID
	TemplateGroupId *int64 `json:"TemplateGroupId,omitempty" name:"TemplateGroupId"`

	// 经验库动作分组动作列表
	TemplateGroupActions []*TemplateGroupAction `json:"TemplateGroupActions,omitempty" name:"TemplateGroupActions"`

	// 分组标题
	Title *string `json:"Title,omitempty" name:"Title"`

	// 分组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 分组顺序
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 执行模式。1 --- 顺序执行，2 --- 阶段执行
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 对象类型ID
	ObjectTypeId *int64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`

	// 分组创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 分组更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TemplateGroupAction struct {
	// 经验库分组动作ID
	TemplateGroupActionId *int64 `json:"TemplateGroupActionId,omitempty" name:"TemplateGroupActionId"`

	// 动作ID
	ActionId *int64 `json:"ActionId,omitempty" name:"ActionId"`

	// 分组动作顺序
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 分组动作通用配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	GeneralConfiguration *string `json:"GeneralConfiguration,omitempty" name:"GeneralConfiguration"`

	// 分组动作自定义配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomConfiguration *string `json:"CustomConfiguration,omitempty" name:"CustomConfiguration"`

	// 动作分组创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 动作分组更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 动作名称
	ActionTitle *string `json:"ActionTitle,omitempty" name:"ActionTitle"`

	// 自身随机id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RandomId *int64 `json:"RandomId,omitempty" name:"RandomId"`

	// 恢复动作id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecoverId *int64 `json:"RecoverId,omitempty" name:"RecoverId"`

	// 执行动作id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecuteId *int64 `json:"ExecuteId,omitempty" name:"ExecuteId"`

	// 调用api类型，0:tat, 1:云api
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionApiType *int64 `json:"ActionApiType,omitempty" name:"ActionApiType"`

	// 1:故障，2:恢复
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionAttribute *int64 `json:"ActionAttribute,omitempty" name:"ActionAttribute"`

	// 动作类型：平台和自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

type TemplateListItem struct {
	// 经验库ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 经验库标题
	TemplateTitle *string `json:"TemplateTitle,omitempty" name:"TemplateTitle"`

	// 经验库描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 经验库标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateTag *string `json:"TemplateTag,omitempty" name:"TemplateTag"`

	// 经验库状态。1 -- 使用中，2 -- 停用
	TemplateIsUsed *int64 `json:"TemplateIsUsed,omitempty" name:"TemplateIsUsed"`

	// 经验库创建时间
	TemplateCreateTime *string `json:"TemplateCreateTime,omitempty" name:"TemplateCreateTime"`

	// 经验库更新时间
	TemplateUpdateTime *string `json:"TemplateUpdateTime,omitempty" name:"TemplateUpdateTime"`

	// 经验库关联的任务数量
	TemplateUsedNum *int64 `json:"TemplateUsedNum,omitempty" name:"TemplateUsedNum"`
}

type TemplateMonitor struct {
	// 监控指标ID
	MonitorId *int64 `json:"MonitorId,omitempty" name:"MonitorId"`

	// 监控指标对象类型ID
	ObjectTypeId *int64 `json:"ObjectTypeId,omitempty" name:"ObjectTypeId"`

	// 指标名称
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 中文指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricChineseName *string `json:"MetricChineseName,omitempty" name:"MetricChineseName"`
}

type TemplatePolicy struct {
	// 保护策略ID列表
	TemplatePolicyIdList []*string `json:"TemplatePolicyIdList,omitempty" name:"TemplatePolicyIdList"`

	// 策略规则
	TemplatePolicyRule *string `json:"TemplatePolicyRule,omitempty" name:"TemplatePolicyRule"`

	// 护栏策略生效处理策略 1:顺序执行，2:暂停
	TemplatePolicyDealType *int64 `json:"TemplatePolicyDealType,omitempty" name:"TemplatePolicyDealType"`
}