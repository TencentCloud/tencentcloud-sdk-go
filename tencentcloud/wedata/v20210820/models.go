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

type AlarmInfo struct {
	// 关联任务id
	TaskIds *string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 告警类别；failure表示失败告警；overtime表示超时告警
	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`

	// 告警方式；SMS表示短信；Email表示邮件；HTTP 表示接口方式；Wechat表示微信方式
	AlarmWay *string `json:"AlarmWay,omitempty" name:"AlarmWay"`

	// 告警接收人，多个告警接收人以;分割
	AlarmRecipient *string `json:"AlarmRecipient,omitempty" name:"AlarmRecipient"`

	// 告警接收人id，多个告警接收人id以;分割
	AlarmRecipientId *string `json:"AlarmRecipientId,omitempty" name:"AlarmRecipientId"`

	// 预计运行的小时，取值范围0-23
	Hours *uint64 `json:"Hours,omitempty" name:"Hours"`

	// 预计运行分钟，取值范围0-59
	Minutes *uint64 `json:"Minutes,omitempty" name:"Minutes"`

	// 告警出发时机；1表示第一次运行失败；2表示所有重试完成后失败；
	TriggerType *uint64 `json:"TriggerType,omitempty" name:"TriggerType"`

	// 告警信息id
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// 告警状态设置；1表示可用；0表示不可用，默认可用
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type BatchDeleteTasksNewRequestParams struct {
	// 批量删除的任务TaskId
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// true : 删除后下游任务可正常运行
	// false：删除后下游任务不可运行
	DeleteMode *bool `json:"DeleteMode,omitempty" name:"DeleteMode"`

	// true：通知下游任务责任人
	// false:  不通知下游任务责任人
	EnableNotify *bool `json:"EnableNotify,omitempty" name:"EnableNotify"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchDeleteTasksNewRequest struct {
	*tchttp.BaseRequest
	
	// 批量删除的任务TaskId
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// true : 删除后下游任务可正常运行
	// false：删除后下游任务不可运行
	DeleteMode *bool `json:"DeleteMode,omitempty" name:"DeleteMode"`

	// true：通知下游任务责任人
	// false:  不通知下游任务责任人
	EnableNotify *bool `json:"EnableNotify,omitempty" name:"EnableNotify"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchDeleteTasksNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteTasksNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "DeleteMode")
	delete(f, "EnableNotify")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteTasksNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteTasksNewResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	Data *BatchOperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDeleteTasksNewResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteTasksNewResponseParams `json:"Response"`
}

func (r *BatchDeleteTasksNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteTasksNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyOwnersNewRequestParams struct {
	// 需要更新责任人的TaskId数组
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 需要更新的责任人
	Owners *string `json:"Owners,omitempty" name:"Owners"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchModifyOwnersNewRequest struct {
	*tchttp.BaseRequest
	
	// 需要更新责任人的TaskId数组
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 需要更新的责任人
	Owners *string `json:"Owners,omitempty" name:"Owners"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchModifyOwnersNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyOwnersNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "Owners")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyOwnersNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyOwnersNewResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	Data *BatchOperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchModifyOwnersNewResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyOwnersNewResponseParams `json:"Response"`
}

func (r *BatchModifyOwnersNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyOwnersNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchOperateResult struct {
	// 批量操作成功数
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 批量操作失败数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 批量操作的总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type BatchResult struct {
	// 正在运行的任务数
	Running *int64 `json:"Running,omitempty" name:"Running"`

	// 执行成功的任务数
	Success *int64 `json:"Success,omitempty" name:"Success"`

	// 执行失败的任务数
	Failed *int64 `json:"Failed,omitempty" name:"Failed"`

	// 总任务数
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type BatchReturn struct {
	// 执行结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 执行情况备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitempty" name:"ErrorDesc"`

	// 执行情况id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorId *string `json:"ErrorId,omitempty" name:"ErrorId"`
}

// Predefined struct for user
type BatchStopTasksNewRequestParams struct {
	// 批量停止任务的TaskId
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchStopTasksNewRequest struct {
	*tchttp.BaseRequest
	
	// 批量停止任务的TaskId
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchStopTasksNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStopTasksNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchStopTasksNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStopTasksNewResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	Data *BatchOperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchStopTasksNewResponse struct {
	*tchttp.BaseResponse
	Response *BatchStopTasksNewResponseParams `json:"Response"`
}

func (r *BatchStopTasksNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStopTasksNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CanvasInfo struct {
	// 画布任务信息
	TasksList []*TaskCanvasInfo `json:"TasksList,omitempty" name:"TasksList"`

	// 画布任务链接信息
	LinksList []*TaskLinkInfo `json:"LinksList,omitempty" name:"LinksList"`
}

type CommonContent struct {
	// 详情内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CommonId struct {
	// Id值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`
}

// Predefined struct for user
type CreateFolderRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`
}

type CreateFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`
}

func (r *CreateFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderName")
	delete(f, "ParentsFolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFolderResponseParams struct {
	// 文件夹Id，null则创建失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CommonId `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateFolderResponseParams `json:"Response"`
}

func (r *CreateFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 26离线同步，30Python，31PySpark，32DLC，33Impala，34Hive SQL，35Shell，36Spark SQL，39Spark，40CDW PG，92MapReduce
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 扩展属性
	TaskExt []*TaskExtInfo `json:"TaskExt,omitempty" name:"TaskExt"`
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 26离线同步，30Python，31PySpark，32DLC，33Impala，34Hive SQL，35Shell，36Spark SQL，39Spark，40CDW PG，92MapReduce
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 扩展属性
	TaskExt []*TaskExtInfo `json:"TaskExt,omitempty" name:"TaskExt"`
}

func (r *CreateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "TaskName")
	delete(f, "TaskType")
	delete(f, "TaskExt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskResponseParams struct {
	// 返回任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CommonId `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskResponseParams `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

type CreateWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

func (r *CreateWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowName")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowResponseParams struct {
	// 返回工作流Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CommonId `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkflowResponseParams `json:"Response"`
}

func (r *CreateWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFolderRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

type DeleteFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

func (r *DeleteFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFolderResponseParams struct {
	// true代表删除成功，false代表删除失败
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteFolderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFolderResponseParams `json:"Response"`
}

func (r *DeleteFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowNewRequestParams struct {
	// 工作流id
	WorkFlowId *string `json:"WorkFlowId,omitempty" name:"WorkFlowId"`

	// true : 删除后下游任务可正常运行
	// false：删除后下游任务不可运行
	DeleteMode *bool `json:"DeleteMode,omitempty" name:"DeleteMode"`

	// true：通知下游任务责任人
	// false:  不通知下游任务责任人
	EnableNotify *bool `json:"EnableNotify,omitempty" name:"EnableNotify"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DeleteWorkflowNewRequest struct {
	*tchttp.BaseRequest
	
	// 工作流id
	WorkFlowId *string `json:"WorkFlowId,omitempty" name:"WorkFlowId"`

	// true : 删除后下游任务可正常运行
	// false：删除后下游任务不可运行
	DeleteMode *bool `json:"DeleteMode,omitempty" name:"DeleteMode"`

	// true：通知下游任务责任人
	// false:  不通知下游任务责任人
	EnableNotify *bool `json:"EnableNotify,omitempty" name:"EnableNotify"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteWorkflowNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkFlowId")
	delete(f, "DeleteMode")
	delete(f, "EnableNotify")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkflowNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowNewResponseParams struct {
	// 返回删除结果
	Data *OperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteWorkflowNewResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkflowNewResponseParams `json:"Response"`
}

func (r *DeleteWorkflowNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DependencyConfig struct {
	// 仅五种周期运行依赖配置： HOUR,DAY,WEEK,MONTH,YEAR,CRONTAB,MINUTE
	DependConfType *string `json:"DependConfType,omitempty" name:"DependConfType"`

	// 依赖配置从属周期类型，CURRENT_HOUR，PREVIOUS_HOUR，CURRENT_DAY，PREVIOUS_DAY，PREVIOUS_WEEK，PREVIOUS_FRIDAY，PREVIOUS_WEEKEND，CURRENT_MONTH，PREVIOUS_MONTH，PREVIOUS_END_OF_MONTH
	//      * PREVIOUS_BEGIN_OF_MONTH，ALL_MONTH_OF_YEAR，ALL_DAY_OF_YEAR，CURRENT_YEAR，CURRENT，CURRENT_MINUTE，PREVIOUS_MINUTE_CYCLE，PREVIOUS_HOUR_CYCLE
	SubordinateCyclicType *string `json:"SubordinateCyclicType,omitempty" name:"SubordinateCyclicType"`

	// WAITING，等待（默认策略）EXECUTING:执行
	DependencyStrategy *string `json:"DependencyStrategy,omitempty" name:"DependencyStrategy"`

	// 父任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentTask *TaskInnerInfo `json:"ParentTask,omitempty" name:"ParentTask"`

	// 子任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SonTask *TaskInnerInfo `json:"SonTask,omitempty" name:"SonTask"`
}

// Predefined struct for user
type DescribeDependTasksNewRequestParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 上游/下游层级1-6级
	Deep *uint64 `json:"Deep,omitempty" name:"Deep"`

	// 1: 表示查询上游节点；0:表示查询下游节点；2：表示查询上游和下游节点
	Up *uint64 `json:"Up,omitempty" name:"Up"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`
}

type DescribeDependTasksNewRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 上游/下游层级1-6级
	Deep *uint64 `json:"Deep,omitempty" name:"Deep"`

	// 1: 表示查询上游节点；0:表示查询下游节点；2：表示查询上游和下游节点
	Up *uint64 `json:"Up,omitempty" name:"Up"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`
}

func (r *DescribeDependTasksNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDependTasksNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Deep")
	delete(f, "Up")
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDependTasksNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDependTasksNewResponseParams struct {
	// 画布任务和链接信息
	Data *CanvasInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDependTasksNewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDependTasksNewResponseParams `json:"Response"`
}

func (r *DescribeDependTasksNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDependTasksNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFolderListData struct {
	// 文件夹信息列表
	Items []*Folder `json:"Items,omitempty" name:"Items"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 页号
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

// Predefined struct for user
type DescribeFolderListRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`

	// 关键字
	KeyWords *string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeFolderListRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`

	// 关键字
	KeyWords *string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeFolderListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFolderListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentsFolderId")
	delete(f, "KeyWords")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFolderListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFolderListResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DescribeFolderListData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFolderListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFolderListResponseParams `json:"Response"`
}

func (r *DescribeFolderListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFolderListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFolderWorkflowListData struct {
	// 工作流信息列表
	Items []*Workflow `json:"Items,omitempty" name:"Items"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 页号
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

// Predefined struct for user
type DescribeFolderWorkflowListRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`

	// 关键字
	KeyWords *string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeFolderWorkflowListRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`

	// 关键字
	KeyWords *string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeFolderWorkflowListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFolderWorkflowListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentsFolderId")
	delete(f, "KeyWords")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFolderWorkflowListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFolderWorkflowListResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DescribeFolderWorkflowListData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFolderWorkflowListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFolderWorkflowListResponseParams `json:"Response"`
}

func (r *DescribeFolderWorkflowListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFolderWorkflowListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogsRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`
}

type DescribeInstanceLogsRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`
}

func (r *DescribeInstanceLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "CurRunDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogsResponseParams struct {
	// 返回日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*InstanceLog `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLogsResponseParams `json:"Response"`
}

func (r *DescribeInstanceLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

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

// Predefined struct for user
type DescribeTaskDetailRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务告警状态
	TaskAlarmStatus *int64 `json:"TaskAlarmStatus,omitempty" name:"TaskAlarmStatus"`
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务告警状态
	TaskAlarmStatus *int64 `json:"TaskAlarmStatus,omitempty" name:"TaskAlarmStatus"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "TaskAlarmStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailResponseParams struct {
	// 任务详情1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TaskInfoData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type DescribeTaskScriptRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeTaskScriptRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskScriptResponseParams struct {
	// 任务脚本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TaskScriptContent `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskScriptResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskScriptResponseParams `json:"Response"`
}

func (r *DescribeTaskScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksByPageRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeTasksByPageRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeTasksByPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksByPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksByPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksByPageResponseParams struct {
	// 无1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TaskInfoDataPage `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTasksByPageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksByPageResponseParams `json:"Response"`
}

func (r *DescribeTasksByPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksByPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Folder struct {
	// 文件ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 文件夹名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 所属项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type ForceSucInstancesRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例嵌套集合
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`
}

type ForceSucInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例嵌套集合
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`
}

func (r *ForceSucInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForceSucInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ForceSucInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ForceSucInstancesResponseParams struct {
	// 返回实例批量终止结果
	Data *OperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ForceSucInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ForceSucInstancesResponseParams `json:"Response"`
}

func (r *ForceSucInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForceSucInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeTasksByMultiWorkflowRequestParams struct {
	// 工作流Id集合
	WorkFlowIds []*string `json:"WorkFlowIds,omitempty" name:"WorkFlowIds"`
}

type FreezeTasksByMultiWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 工作流Id集合
	WorkFlowIds []*string `json:"WorkFlowIds,omitempty" name:"WorkFlowIds"`
}

func (r *FreezeTasksByMultiWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeTasksByMultiWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkFlowIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FreezeTasksByMultiWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeTasksByMultiWorkflowResponseParams struct {
	// 操作结果
	Data *OperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FreezeTasksByMultiWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *FreezeTasksByMultiWorkflowResponseParams `json:"Response"`
}

func (r *FreezeTasksByMultiWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeTasksByMultiWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeTasksRequestParams struct {
	// 任务列表
	Tasks []*SimpleTaskInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 任务操作是否消息通知下游任务责任人
	OperateIsInform *bool `json:"OperateIsInform,omitempty" name:"OperateIsInform"`
}

type FreezeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务列表
	Tasks []*SimpleTaskInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 任务操作是否消息通知下游任务责任人
	OperateIsInform *bool `json:"OperateIsInform,omitempty" name:"OperateIsInform"`
}

func (r *FreezeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tasks")
	delete(f, "OperateIsInform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FreezeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeTasksResponseParams struct {
	// 操作结果
	Data *OperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FreezeTasksResponse struct {
	*tchttp.BaseResponse
	Response *FreezeTasksResponseParams `json:"Response"`
}

func (r *FreezeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GeneralTaskParam struct {
	// 通用任务参数类型,例：SPARK_SQL
	Type *string `json:"Type,omitempty" name:"Type"`

	// 通用任务参数内容,直接作用于任务的参数。不同参数用;
	// 分割
	Value *string `json:"Value,omitempty" name:"Value"`
}

type InstanceInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`
}

type InstanceLog struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 尝试运行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tries *string `json:"Tries,omitempty" name:"Tries"`

	// 日志更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdate *string `json:"LastUpdate,omitempty" name:"LastUpdate"`

	// 日志所在节点
	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`

	// 文件名  含全路径
	OriginFileName *string `json:"OriginFileName,omitempty" name:"OriginFileName"`

	// 日志创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例日志类型, run: 运行; kill: 终止
	InstanceLogType *string `json:"InstanceLogType,omitempty" name:"InstanceLogType"`

	// 运行耗时
	CostTime *float64 `json:"CostTime,omitempty" name:"CostTime"`
}

type IntegrationNodeDetail struct {
	// 集成节点名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 集成节点类型
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 节点数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 节点描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// 节点配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config []*RecordField `json:"Config,omitempty" name:"Config"`

	// 节点扩展配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtConfig []*RecordField `json:"ExtConfig,omitempty" name:"ExtConfig"`

	// 节点schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schema []*IntegrationNodeSchema `json:"Schema,omitempty" name:"Schema"`

	// 节点映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeMapping *IntegrationNodeMapping `json:"NodeMapping,omitempty" name:"NodeMapping"`

	// owner uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

type IntegrationNodeMapping struct {
	// 源节点id
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// 目标节点id
	SinkId *string `json:"SinkId,omitempty" name:"SinkId"`

	// 源节点schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceSchema []*IntegrationNodeSchema `json:"SourceSchema,omitempty" name:"SourceSchema"`

	// 节点schema映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaMappings []*IntegrationNodeSchemaMapping `json:"SchemaMappings,omitempty" name:"SchemaMappings"`

	// 节点映射扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtConfig []*RecordField `json:"ExtConfig,omitempty" name:"ExtConfig"`
}

type IntegrationNodeSchema struct {
	// schema id
	Id *string `json:"Id,omitempty" name:"Id"`

	// schema名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// schema类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// schema值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// schema拓展属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*RecordField `json:"Properties,omitempty" name:"Properties"`
}

type IntegrationNodeSchemaMapping struct {
	// 源schema id
	SourceSchemaId *string `json:"SourceSchemaId,omitempty" name:"SourceSchemaId"`

	// 目标schema id
	SinkSchemaId *string `json:"SinkSchemaId,omitempty" name:"SinkSchemaId"`
}

// Predefined struct for user
type KillInstancesRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例嵌套集合
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`
}

type KillInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例嵌套集合
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`
}

func (r *KillInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillInstancesResponseParams struct {
	// 返回实例批量终止结果
	Data *OperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type KillInstancesResponse struct {
	*tchttp.BaseResponse
	Response *KillInstancesResponseParams `json:"Response"`
}

func (r *KillInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MakeUpTasksNewRequestParams struct {
	// 补录的当前任务的taskId数组
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 补录开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 补录结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 补录选项标识，1表示当前任务；2表示当前及下游任务；3表示下游任务
	MakeUpType *uint64 `json:"MakeUpType,omitempty" name:"MakeUpType"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// true: 检查父任务实例状态；false: 不检查父任务实例状态
	CheckParent *bool `json:"CheckParent,omitempty" name:"CheckParent"`
}

type MakeUpTasksNewRequest struct {
	*tchttp.BaseRequest
	
	// 补录的当前任务的taskId数组
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 补录开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 补录结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 补录选项标识，1表示当前任务；2表示当前及下游任务；3表示下游任务
	MakeUpType *uint64 `json:"MakeUpType,omitempty" name:"MakeUpType"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// true: 检查父任务实例状态；false: 不检查父任务实例状态
	CheckParent *bool `json:"CheckParent,omitempty" name:"CheckParent"`
}

func (r *MakeUpTasksNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MakeUpTasksNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MakeUpType")
	delete(f, "ProjectId")
	delete(f, "CheckParent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MakeUpTasksNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MakeUpTasksNewResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	Data *BatchOperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MakeUpTasksNewResponse struct {
	*tchttp.BaseResponse
	Response *MakeUpTasksNewResponseParams `json:"Response"`
}

func (r *MakeUpTasksNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MakeUpTasksNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MakeUpWorkflowNewRequestParams struct {
	// 工作流id
	WorkFlowId *string `json:"WorkFlowId,omitempty" name:"WorkFlowId"`

	// 补录开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 补录结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type MakeUpWorkflowNewRequest struct {
	*tchttp.BaseRequest
	
	// 工作流id
	WorkFlowId *string `json:"WorkFlowId,omitempty" name:"WorkFlowId"`

	// 补录开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 补录结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *MakeUpWorkflowNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MakeUpWorkflowNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkFlowId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MakeUpWorkflowNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MakeUpWorkflowNewResponseParams struct {
	// 返回补录成功或失败的任务数
	Data *BatchOperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MakeUpWorkflowNewResponse struct {
	*tchttp.BaseResponse
	Response *MakeUpWorkflowNewResponseParams `json:"Response"`
}

func (r *MakeUpWorkflowNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MakeUpWorkflowNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFolderRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 文件夹Id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`
}

type ModifyFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 文件夹Id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`
}

func (r *ModifyFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderName")
	delete(f, "FolderId")
	delete(f, "ParentsFolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFolderResponseParams struct {
	// true代表成功，false代表失败
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyFolderResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFolderResponseParams `json:"Response"`
}

func (r *ModifyFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskInfoRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 执行时间，单位分钟，天/周/月/年调度才有。比如天调度，每天的02:00点执行一次，delayTime就是120分钟
	DelayTime *int64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *int64 `json:"StartupTime,omitempty" name:"StartupTime"`

	// 自依赖类型  1:有序串行 一次一个 排队, 2: 无序串行 一次一个 不排队， 3:并行 一次多个
	SelfDepend *int64 `json:"SelfDepend,omitempty" name:"SelfDepend"`

	// 生效开始时间，格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 生效结束时间，格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// "周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitempty" name:"CycleType"`

	// 步长，间隔时间，最小1
	CycleStep *int64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// cron表达式  周期类型为crontab调度才需要
	CrontabExpression *string `json:"CrontabExpression,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" name:"ExecutionEndTime"`

	// 新的任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 失败重试间隔,单位分钟，创建任务的时候已经给了默认值
	RetryWait *int64 `json:"RetryWait,omitempty" name:"RetryWait"`

	// 失败重试次数，创建任务的时候已经给了默认值
	TryLimit *int64 `json:"TryLimit,omitempty" name:"TryLimit"`

	// 是否可重试，1代表可以重试
	Retriable *int64 `json:"Retriable,omitempty" name:"Retriable"`

	// 运行优先级，4高 5中 6低
	RunPriority *int64 `json:"RunPriority,omitempty" name:"RunPriority"`

	// 任务的扩展配置
	TaskExt []*TaskExtInfo `json:"TaskExt,omitempty" name:"TaskExt"`

	// 执行资源组id，需要去资源管理服务上创建调度资源组，并且绑定cvm机器
	ResourceGroup *string `json:"ResourceGroup,omitempty" name:"ResourceGroup"`

	// 资源池队列名称
	YarnQueue *string `json:"YarnQueue,omitempty" name:"YarnQueue"`

	// 资源组下具体执行机，any 表示可以跑在任意一台。
	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`

	// 责任人
	InCharge *string `json:"InCharge,omitempty" name:"InCharge"`

	// 任务备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// 任务参数
	TaskParamInfos []*ParamInfo `json:"TaskParamInfos,omitempty" name:"TaskParamInfos"`

	// 源数据源
	SourceServer *string `json:"SourceServer,omitempty" name:"SourceServer"`

	// 目标数据源
	TargetServer *string `json:"TargetServer,omitempty" name:"TargetServer"`

	// 是否支持工作流依赖 yes / no 默认 no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitempty" name:"DependencyWorkflow"`

	// 依赖配置
	DependencyConfigDTOs []*DependencyConfig `json:"DependencyConfigDTOs,omitempty" name:"DependencyConfigDTOs"`
}

type ModifyTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 执行时间，单位分钟，天/周/月/年调度才有。比如天调度，每天的02:00点执行一次，delayTime就是120分钟
	DelayTime *int64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *int64 `json:"StartupTime,omitempty" name:"StartupTime"`

	// 自依赖类型  1:有序串行 一次一个 排队, 2: 无序串行 一次一个 不排队， 3:并行 一次多个
	SelfDepend *int64 `json:"SelfDepend,omitempty" name:"SelfDepend"`

	// 生效开始时间，格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 生效结束时间，格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// "周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitempty" name:"CycleType"`

	// 步长，间隔时间，最小1
	CycleStep *int64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// cron表达式  周期类型为crontab调度才需要
	CrontabExpression *string `json:"CrontabExpression,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" name:"ExecutionEndTime"`

	// 新的任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 失败重试间隔,单位分钟，创建任务的时候已经给了默认值
	RetryWait *int64 `json:"RetryWait,omitempty" name:"RetryWait"`

	// 失败重试次数，创建任务的时候已经给了默认值
	TryLimit *int64 `json:"TryLimit,omitempty" name:"TryLimit"`

	// 是否可重试，1代表可以重试
	Retriable *int64 `json:"Retriable,omitempty" name:"Retriable"`

	// 运行优先级，4高 5中 6低
	RunPriority *int64 `json:"RunPriority,omitempty" name:"RunPriority"`

	// 任务的扩展配置
	TaskExt []*TaskExtInfo `json:"TaskExt,omitempty" name:"TaskExt"`

	// 执行资源组id，需要去资源管理服务上创建调度资源组，并且绑定cvm机器
	ResourceGroup *string `json:"ResourceGroup,omitempty" name:"ResourceGroup"`

	// 资源池队列名称
	YarnQueue *string `json:"YarnQueue,omitempty" name:"YarnQueue"`

	// 资源组下具体执行机，any 表示可以跑在任意一台。
	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`

	// 责任人
	InCharge *string `json:"InCharge,omitempty" name:"InCharge"`

	// 任务备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// 任务参数
	TaskParamInfos []*ParamInfo `json:"TaskParamInfos,omitempty" name:"TaskParamInfos"`

	// 源数据源
	SourceServer *string `json:"SourceServer,omitempty" name:"SourceServer"`

	// 目标数据源
	TargetServer *string `json:"TargetServer,omitempty" name:"TargetServer"`

	// 是否支持工作流依赖 yes / no 默认 no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitempty" name:"DependencyWorkflow"`

	// 依赖配置
	DependencyConfigDTOs []*DependencyConfig `json:"DependencyConfigDTOs,omitempty" name:"DependencyConfigDTOs"`
}

func (r *ModifyTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "DelayTime")
	delete(f, "StartupTime")
	delete(f, "SelfDepend")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskAction")
	delete(f, "CycleType")
	delete(f, "CycleStep")
	delete(f, "CrontabExpression")
	delete(f, "ExecutionStartTime")
	delete(f, "ExecutionEndTime")
	delete(f, "TaskName")
	delete(f, "RetryWait")
	delete(f, "TryLimit")
	delete(f, "Retriable")
	delete(f, "RunPriority")
	delete(f, "TaskExt")
	delete(f, "ResourceGroup")
	delete(f, "YarnQueue")
	delete(f, "BrokerIp")
	delete(f, "InCharge")
	delete(f, "Notes")
	delete(f, "TaskParamInfos")
	delete(f, "SourceServer")
	delete(f, "TargetServer")
	delete(f, "DependencyWorkflow")
	delete(f, "DependencyConfigDTOs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskInfoResponseParams struct {
	// 执行结果
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskInfoResponseParams `json:"Response"`
}

func (r *ModifyTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskLinksRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 父任务ID
	TaskFrom *string `json:"TaskFrom,omitempty" name:"TaskFrom"`

	// 子任务ID
	TaskTo *string `json:"TaskTo,omitempty" name:"TaskTo"`

	// 子任务工作流
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 父任务工作流
	RealFromWorkflowId *string `json:"RealFromWorkflowId,omitempty" name:"RealFromWorkflowId"`

	// 父子任务之间的依赖关系
	LinkDependencyType *string `json:"LinkDependencyType,omitempty" name:"LinkDependencyType"`
}

type ModifyTaskLinksRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 父任务ID
	TaskFrom *string `json:"TaskFrom,omitempty" name:"TaskFrom"`

	// 子任务ID
	TaskTo *string `json:"TaskTo,omitempty" name:"TaskTo"`

	// 子任务工作流
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 父任务工作流
	RealFromWorkflowId *string `json:"RealFromWorkflowId,omitempty" name:"RealFromWorkflowId"`

	// 父子任务之间的依赖关系
	LinkDependencyType *string `json:"LinkDependencyType,omitempty" name:"LinkDependencyType"`
}

func (r *ModifyTaskLinksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskLinksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskFrom")
	delete(f, "TaskTo")
	delete(f, "WorkflowId")
	delete(f, "RealFromWorkflowId")
	delete(f, "LinkDependencyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskLinksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskLinksResponseParams struct {
	// 成功或者失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTaskLinksResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskLinksResponseParams `json:"Response"`
}

func (r *ModifyTaskLinksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskLinksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskScriptRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 脚本内容 base64编码
	ScriptContent *string `json:"ScriptContent,omitempty" name:"ScriptContent"`

	// 集成任务脚本配置
	IntegrationNodeDetails []*IntegrationNodeDetail `json:"IntegrationNodeDetails,omitempty" name:"IntegrationNodeDetails"`
}

type ModifyTaskScriptRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 脚本内容 base64编码
	ScriptContent *string `json:"ScriptContent,omitempty" name:"ScriptContent"`

	// 集成任务脚本配置
	IntegrationNodeDetails []*IntegrationNodeDetail `json:"IntegrationNodeDetails,omitempty" name:"IntegrationNodeDetails"`
}

func (r *ModifyTaskScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "ScriptContent")
	delete(f, "IntegrationNodeDetails")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskScriptResponseParams struct {
	// 详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CommonContent `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTaskScriptResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskScriptResponseParams `json:"Response"`
}

func (r *ModifyTaskScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkflowInfoRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 责任人
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 责任人id
	OwnerId *string `json:"OwnerId,omitempty" name:"OwnerId"`

	// 备注
	WorkflowDesc *string `json:"WorkflowDesc,omitempty" name:"WorkflowDesc"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 工作流所属用户分组id  若有多个,分号隔开: a;b;c
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// 工作流所属用户分组名称  若有多个,分号隔开: a;b;c
	UserGroupName *string `json:"UserGroupName,omitempty" name:"UserGroupName"`

	// 工作流参数列表
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitempty" name:"WorkflowParams"`

	// 用于配置优化参数（线程、内存、CPU核数等），仅作用于Spark SQL节点。多个参数用英文分号分隔。
	GeneralTaskParams []*GeneralTaskParam `json:"GeneralTaskParams,omitempty" name:"GeneralTaskParams"`
}

type ModifyWorkflowInfoRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 责任人
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 责任人id
	OwnerId *string `json:"OwnerId,omitempty" name:"OwnerId"`

	// 备注
	WorkflowDesc *string `json:"WorkflowDesc,omitempty" name:"WorkflowDesc"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 工作流所属用户分组id  若有多个,分号隔开: a;b;c
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// 工作流所属用户分组名称  若有多个,分号隔开: a;b;c
	UserGroupName *string `json:"UserGroupName,omitempty" name:"UserGroupName"`

	// 工作流参数列表
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitempty" name:"WorkflowParams"`

	// 用于配置优化参数（线程、内存、CPU核数等），仅作用于Spark SQL节点。多个参数用英文分号分隔。
	GeneralTaskParams []*GeneralTaskParam `json:"GeneralTaskParams,omitempty" name:"GeneralTaskParams"`
}

func (r *ModifyWorkflowInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkflowInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "Owner")
	delete(f, "OwnerId")
	delete(f, "WorkflowDesc")
	delete(f, "WorkflowName")
	delete(f, "FolderId")
	delete(f, "UserGroupId")
	delete(f, "UserGroupName")
	delete(f, "WorkflowParams")
	delete(f, "GeneralTaskParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkflowInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkflowInfoResponseParams struct {
	// true代表成功，false代表失败
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWorkflowInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkflowInfoResponseParams `json:"Response"`
}

func (r *ModifyWorkflowInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkflowInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkflowScheduleRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 延迟时间，单位分钟
	DelayTime *int64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *int64 `json:"StartupTime,omitempty" name:"StartupTime"`

	// 自依赖类型  1:有序串行 一次一个 排队, 2: 无序串行 一次一个 不排队， 3:并行 一次多个
	SelfDepend *int64 `json:"SelfDepend,omitempty" name:"SelfDepend"`

	// "周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitempty" name:"CycleType"`

	// 步长，间隔时间，最小1
	CycleStep *int64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// 生效开始时间，格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 生效结束时间，格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// cron表达式  周期类型为crontab调度才需要
	CrontabExpression *string `json:"CrontabExpression,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" name:"ExecutionEndTime"`

	// 工作流依赖 ,yes 或者no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitempty" name:"DependencyWorkflow"`
}

type ModifyWorkflowScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 延迟时间，单位分钟
	DelayTime *int64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *int64 `json:"StartupTime,omitempty" name:"StartupTime"`

	// 自依赖类型  1:有序串行 一次一个 排队, 2: 无序串行 一次一个 不排队， 3:并行 一次多个
	SelfDepend *int64 `json:"SelfDepend,omitempty" name:"SelfDepend"`

	// "周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitempty" name:"CycleType"`

	// 步长，间隔时间，最小1
	CycleStep *int64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// 生效开始时间，格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 生效结束时间，格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// cron表达式  周期类型为crontab调度才需要
	CrontabExpression *string `json:"CrontabExpression,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" name:"ExecutionEndTime"`

	// 工作流依赖 ,yes 或者no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitempty" name:"DependencyWorkflow"`
}

func (r *ModifyWorkflowScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkflowScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "DelayTime")
	delete(f, "StartupTime")
	delete(f, "SelfDepend")
	delete(f, "CycleType")
	delete(f, "CycleStep")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskAction")
	delete(f, "CrontabExpression")
	delete(f, "ExecutionStartTime")
	delete(f, "ExecutionEndTime")
	delete(f, "DependencyWorkflow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkflowScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkflowScheduleResponseParams struct {
	// 执行结果
	Data *BatchResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWorkflowScheduleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkflowScheduleResponseParams `json:"Response"`
}

func (r *ModifyWorkflowScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkflowScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateResult struct {
	// 操作结果；true表示成功；false表示失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 错误编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorId *string `json:"ErrorId,omitempty" name:"ErrorId"`

	// 操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitempty" name:"ErrorDesc"`
}

type OrderField struct {
	// 排序字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排序方向：ASC|DESC
	Direction *string `json:"Direction,omitempty" name:"Direction"`
}

type ParamInfo struct {
	// 参数名
	ParamKey *string `json:"ParamKey,omitempty" name:"ParamKey"`

	// 参数值
	ParamValue *string `json:"ParamValue,omitempty" name:"ParamValue"`
}

type RecordField struct {
	// 字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段值
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type RegisterEventListenerRequestParams struct {
	// 关键字，如果是任务，则传任务Id
	Key *string `json:"Key,omitempty" name:"Key"`

	// 事件名称
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 事件类型，默认 REST_API
	Type *string `json:"Type,omitempty" name:"Type"`

	// 配置信息，比如最长等待时间1天配置json：{"maxWaitEventTime":1,"maxWaitEventTimeUnit":"DAYS"}
	Properties *string `json:"Properties,omitempty" name:"Properties"`
}

type RegisterEventListenerRequest struct {
	*tchttp.BaseRequest
	
	// 关键字，如果是任务，则传任务Id
	Key *string `json:"Key,omitempty" name:"Key"`

	// 事件名称
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 事件类型，默认 REST_API
	Type *string `json:"Type,omitempty" name:"Type"`

	// 配置信息，比如最长等待时间1天配置json：{"maxWaitEventTime":1,"maxWaitEventTimeUnit":"DAYS"}
	Properties *string `json:"Properties,omitempty" name:"Properties"`
}

func (r *RegisterEventListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterEventListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Key")
	delete(f, "EventName")
	delete(f, "ProjectId")
	delete(f, "Type")
	delete(f, "Properties")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterEventListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterEventListenerResponseParams struct {
	// 成功或者失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *BatchReturn `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterEventListenerResponse struct {
	*tchttp.BaseResponse
	Response *RegisterEventListenerResponseParams `json:"Response"`
}

func (r *RegisterEventListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterEventListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterEventRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 事件名称，支持英文、数字和下划线，最长20个字符, 不能以数字下划线开头。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 事件分割类型，周期类型: DAY，HOUR，MIN，SECOND
	EventSubType *string `json:"EventSubType,omitempty" name:"EventSubType"`

	// 广播：BROADCAST,单播：SINGLE
	EventBroadcastType *string `json:"EventBroadcastType,omitempty" name:"EventBroadcastType"`

	// 周期类型为天和小时为HOURS ，周期类型为分钟 ：MINUTES,周期类型为秒：SECONDS
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// TBDS 事件所属人
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 事件类型，默认值：TIME_SERIES
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 对应day： yyyyMMdd，对应HOUR：yyyyMMddHH，对应MIN：yyyyMMddHHmm，对应SECOND：yyyyMMddHHmmss
	DimensionFormat *string `json:"DimensionFormat,omitempty" name:"DimensionFormat"`

	// 存活时间
	TimeToLive *int64 `json:"TimeToLive,omitempty" name:"TimeToLive"`

	// 事件描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type RegisterEventRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 事件名称，支持英文、数字和下划线，最长20个字符, 不能以数字下划线开头。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 事件分割类型，周期类型: DAY，HOUR，MIN，SECOND
	EventSubType *string `json:"EventSubType,omitempty" name:"EventSubType"`

	// 广播：BROADCAST,单播：SINGLE
	EventBroadcastType *string `json:"EventBroadcastType,omitempty" name:"EventBroadcastType"`

	// 周期类型为天和小时为HOURS ，周期类型为分钟 ：MINUTES,周期类型为秒：SECONDS
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// TBDS 事件所属人
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 事件类型，默认值：TIME_SERIES
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 对应day： yyyyMMdd，对应HOUR：yyyyMMddHH，对应MIN：yyyyMMddHHmm，对应SECOND：yyyyMMddHHmmss
	DimensionFormat *string `json:"DimensionFormat,omitempty" name:"DimensionFormat"`

	// 存活时间
	TimeToLive *int64 `json:"TimeToLive,omitempty" name:"TimeToLive"`

	// 事件描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *RegisterEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "EventSubType")
	delete(f, "EventBroadcastType")
	delete(f, "TimeUnit")
	delete(f, "Owner")
	delete(f, "EventType")
	delete(f, "DimensionFormat")
	delete(f, "TimeToLive")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterEventResponseParams struct {
	// 成功或者失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *BatchReturn `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterEventResponse struct {
	*tchttp.BaseResponse
	Response *RegisterEventResponseParams `json:"Response"`
}

func (r *RegisterEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RerunInstancesRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例嵌套集合
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务
	CheckFather *bool `json:"CheckFather,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子
	RerunType *string `json:"RerunType,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖
	DependentWay *string `json:"DependentWay,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否
	SkipEventListening *bool `json:"SkipEventListening,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitempty" name:"SonInstanceType"`
}

type RerunInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例嵌套集合
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务
	CheckFather *bool `json:"CheckFather,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子
	RerunType *string `json:"RerunType,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖
	DependentWay *string `json:"DependentWay,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否
	SkipEventListening *bool `json:"SkipEventListening,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitempty" name:"SonInstanceType"`
}

func (r *RerunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Instances")
	delete(f, "CheckFather")
	delete(f, "RerunType")
	delete(f, "DependentWay")
	delete(f, "SkipEventListening")
	delete(f, "SonInstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RerunInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RerunInstancesResponseParams struct {
	// 返回实例批量终止结果
	Data *OperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RerunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RerunInstancesResponseParams `json:"Response"`
}

func (r *RerunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunTaskRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type RunTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RunTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunTaskResponseParams struct {
	// 运行成功或者失败
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunTaskResponse struct {
	*tchttp.BaseResponse
	Response *RunTaskResponseParams `json:"Response"`
}

func (r *RunTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTaskAlarmNewRequestParams struct {
	// 设置任务超时告警和失败告警信息
	AlarmInfoList []*AlarmInfo `json:"AlarmInfoList,omitempty" name:"AlarmInfoList"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type SetTaskAlarmNewRequest struct {
	*tchttp.BaseRequest
	
	// 设置任务超时告警和失败告警信息
	AlarmInfoList []*AlarmInfo `json:"AlarmInfoList,omitempty" name:"AlarmInfoList"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *SetTaskAlarmNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTaskAlarmNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmInfoList")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTaskAlarmNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTaskAlarmNewResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	Data *BatchOperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetTaskAlarmNewResponse struct {
	*tchttp.BaseResponse
	Response *SetTaskAlarmNewResponseParams `json:"Response"`
}

func (r *SetTaskAlarmNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTaskAlarmNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SimpleTaskInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

// Predefined struct for user
type SubmitTaskRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 版本备注
	VersionRemark *string `json:"VersionRemark,omitempty" name:"VersionRemark"`

	// 是否启动调度
	StartScheduling *bool `json:"StartScheduling,omitempty" name:"StartScheduling"`
}

type SubmitTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 版本备注
	VersionRemark *string `json:"VersionRemark,omitempty" name:"VersionRemark"`

	// 是否启动调度
	StartScheduling *bool `json:"StartScheduling,omitempty" name:"StartScheduling"`
}

func (r *SubmitTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "VersionRemark")
	delete(f, "StartScheduling")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTaskResponseParams struct {
	// 成功或者失败
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SubmitTaskResponse struct {
	*tchttp.BaseResponse
	Response *SubmitTaskResponseParams `json:"Response"`
}

func (r *SubmitTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitWorkflow struct {
	// 被提交的任务id集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 执行结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 执行情况备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitempty" name:"ErrorDesc"`

	// 执行情况id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorId *string `json:"ErrorId,omitempty" name:"ErrorId"`
}

// Predefined struct for user
type SubmitWorkflowRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 提交的版本备注
	VersionRemark *string `json:"VersionRemark,omitempty" name:"VersionRemark"`

	// 是否启动调度
	StartScheduling *bool `json:"StartScheduling,omitempty" name:"StartScheduling"`
}

type SubmitWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 提交的版本备注
	VersionRemark *string `json:"VersionRemark,omitempty" name:"VersionRemark"`

	// 是否启动调度
	StartScheduling *bool `json:"StartScheduling,omitempty" name:"StartScheduling"`
}

func (r *SubmitWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "VersionRemark")
	delete(f, "StartScheduling")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitWorkflowResponseParams struct {
	// 执行结果
	Data *SubmitWorkflow `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SubmitWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *SubmitWorkflowResponseParams `json:"Response"`
}

func (r *SubmitWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskCanvasInfo struct {
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

	// 任务状态，'Y','F','O','T','INVALID' 分别表示调度中、已停止、已暂停、停止中、已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务类型id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeId *uint64 `json:"TaskTypeId,omitempty" name:"TaskTypeId"`

	// 任务类型描述，其中任务类型id和任务类型描述的对应的关系为
	// 20	通用数据同步任务
	// 21	JDBC SQL
	// 22	Tbase
	// 25	数据ETL
	// 30	Python
	// 31	PySpark
	// 34	Hive SQL
	// 35	Shell
	// 36	Spark SQL
	// 37	HDFS到HBase
	// 38	SHELL
	// 39	Spark
	// 45	DATA_QUALITY
	// 55	THIVE到MYSQL
	// 56	THIVE到PG
	// 66	HDFS到PG
	// 67	HDFS到Oracle
	// 68	HDFS到MYSQL
	// 69	FTP到HDFS
	// 70	HIVE SQL
	// 72	HIVE到HDFS
	// 75	HDFS到HIVE
	// 81	PYTHONSQL脚本
	// 82	SPARKSCALA计算
	// 83	虫洞任务
	// 84	校验对账文件
	// 85	HDFS到THIVE
	// 86	TDW到HDFS
	// 87	HDFS到TDW
	// 88	校验对账文件
	// 91	FLINK任务
	// 92	MapReduce
	// 98	custom topology
	// 99	kafkatoHDFS
	// 100	kafkatoHbase
	// 101	MYSQL导入至HIVE(DX)
	// 104	MYSQL到HIVE
	// 105	HIVE到MYSQL
	// 106	SQL SERVER到HIVE
	// 107	HIVE到SQL SERVER
	// 108	ORACLE到HIVE
	// 109	HIVE到ORACLE
	// 111	HIVE到MYSQL(NEW)
	// 112	HIVE到PG
	// 113	HIVE到PHOENIX
	// 118	MYSQL到HDFS
	// 119	PG到HDFS
	// 120	ORACLE到HDFS
	// 121	数据质量
	// 10000	自定义业务
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

	// 最近提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstSubmitTime *string `json:"FirstSubmitTime,omitempty" name:"FirstSubmitTime"`

	// 首次运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstRunTime *string `json:"FirstRunTime,omitempty" name:"FirstRunTime"`

	// 调度计划展示描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleDesc *string `json:"ScheduleDesc,omitempty" name:"ScheduleDesc"`

	// 负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitempty" name:"InCharge"`

	// 调度周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleUnit *string `json:"CycleUnit,omitempty" name:"CycleUnit"`

	// 画布x轴坐标点
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftCoordinate *float64 `json:"LeftCoordinate,omitempty" name:"LeftCoordinate"`

	// 画布y轴坐标点
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopCoordinate *float64 `json:"TopCoordinate,omitempty" name:"TopCoordinate"`

	// 跨工作流虚拟任务标识；true标识跨工作流任务；false标识本工作流任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualFlag *bool `json:"VirtualFlag,omitempty" name:"VirtualFlag"`

	// 弹性周期配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// 延迟时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayTime *uint64 `json:"DelayTime,omitempty" name:"DelayTime"`
}

type TaskExtInfo struct {
	// 键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TaskInfoData struct {
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

	// 任务状态，'Y','F','O','T','INVALID' 分别表示调度中、已停止、已暂停、停止中、已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitempty" name:"InCharge"`

	// 跨工作流虚拟任务标识；true标识跨工作流任务；false标识本工作流任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualFlag *bool `json:"VirtualFlag,omitempty" name:"VirtualFlag"`

	// 延时实例生成时间(延时调度)，转换为分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayTime *uint64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// crontab表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabExpression *string `json:"CrontabExpression,omitempty" name:"CrontabExpression"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdate *string `json:"LastUpdate,omitempty" name:"LastUpdate"`

	// 生效日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 执行时间左闭区间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" name:"ExecutionEndTime"`

	// 周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *int64 `json:"CycleType,omitempty" name:"CycleType"`

	// 步长
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleStep *int64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// 延时执行时间（延时执行) 对应为 开始时间 状态为分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupTime *int64 `json:"StartupTime,omitempty" name:"StartupTime"`

	// 重试等待时间,单位分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryWait *int64 `json:"RetryWait,omitempty" name:"RetryWait"`

	// 是否可重试
	// 注意：此字段可能返回 null，表示取不到有效值。
	Retriable *int64 `json:"Retriable,omitempty" name:"Retriable"`

	// 调度扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// 运行次数限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	TryLimit *int64 `json:"TryLimit,omitempty" name:"TryLimit"`

	// 运行优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunPriority *int64 `json:"RunPriority,omitempty" name:"RunPriority"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 指定的运行节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`

	// 集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 最小数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinDateTime *string `json:"MinDateTime,omitempty" name:"MinDateTime"`

	// 最大数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDateTime *string `json:"MaxDateTime,omitempty" name:"MaxDateTime"`

	// 是否自身依赖 是1 否2 并行3
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfDepend *int64 `json:"SelfDepend,omitempty" name:"SelfDepend"`

	// 扩展属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskExt []*TaskExtInfo `json:"TaskExt,omitempty" name:"TaskExt"`

	// 任务备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// 队列
	// 注意：此字段可能返回 null，表示取不到有效值。
	YarnQueue *string `json:"YarnQueue,omitempty" name:"YarnQueue"`

	// 任务版本是否已提交
	// 注意：此字段可能返回 null，表示取不到有效值。
	Submit *bool `json:"Submit,omitempty" name:"Submit"`

	// 最新调度计划变更时间 仅生产态
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastSchedulerCommitTime *string `json:"LastSchedulerCommitTime,omitempty" name:"LastSchedulerCommitTime"`

	// 仅生产态存储于生产态序列化任务信息, 减少base CPU重复密集计算
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalizedJobStartTime *string `json:"NormalizedJobStartTime,omitempty" name:"NormalizedJobStartTime"`

	// 源数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceServer *string `json:"SourceServer,omitempty" name:"SourceServer"`

	// 创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creater *string `json:"Creater,omitempty" name:"Creater"`

	// 分支，依赖关系，and/or, 默认and
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyRel *string `json:"DependencyRel,omitempty" name:"DependencyRel"`

	// 是否支持工作流依赖 yes / no 默认 no
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyWorkflow *string `json:"DependencyWorkflow,omitempty" name:"DependencyWorkflow"`

	// 任务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*ParamInfo `json:"Params,omitempty" name:"Params"`

	// 最后修改的人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUser *string `json:"UpdateUser,omitempty" name:"UpdateUser"`

	// 最后修改的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 最后修改的人Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserId *string `json:"UpdateUserId,omitempty" name:"UpdateUserId"`

	// 调度计划
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDesc *string `json:"SchedulerDesc,omitempty" name:"SchedulerDesc"`

	// 资源组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroup *string `json:"ResourceGroup,omitempty" name:"ResourceGroup"`

	// 版本提交说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`

	// 真实工作流Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealWorkflowId *string `json:"RealWorkflowId,omitempty" name:"RealWorkflowId"`

	// 目标数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetServer *string `json:"TargetServer,omitempty" name:"TargetServer"`

	// 依赖配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyConfigs []*DependencyConfig `json:"DependencyConfigs,omitempty" name:"DependencyConfigs"`

	// 虚拟任务状态1
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualTaskStatus *string `json:"VirtualTaskStatus,omitempty" name:"VirtualTaskStatus"`

	// 虚拟任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualTaskId *string `json:"VirtualTaskId,omitempty" name:"VirtualTaskId"`
}

type TaskInfoDataPage struct {
	// 页号
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 任务集合信息
	Items []*TaskInfoData `json:"Items,omitempty" name:"Items"`

	// 总页数1
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type TaskInnerInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitempty" name:"CycleType"`

	// 虚拟任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualTaskId *string `json:"VirtualTaskId,omitempty" name:"VirtualTaskId"`

	// 虚拟任务标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualFlag *bool `json:"VirtualFlag,omitempty" name:"VirtualFlag"`

	// 真实任务工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealWorkflowId *string `json:"RealWorkflowId,omitempty" name:"RealWorkflowId"`
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

type TaskLinkInfo struct {
	// 下游任务id
	TaskTo *string `json:"TaskTo,omitempty" name:"TaskTo"`

	// 上游任务id
	TaskFrom *string `json:"TaskFrom,omitempty" name:"TaskFrom"`

	// 依赖边类型 1、“real_real”表示任务->任务；2、"virtual_real" 跨工作流任务->任务
	LinkType *string `json:"LinkType,omitempty" name:"LinkType"`

	// 依赖边id
	LinkId *string `json:"LinkId,omitempty" name:"LinkId"`
}

type TaskScriptContent struct {
	// 脚本内容 base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptContent *string `json:"ScriptContent,omitempty" name:"ScriptContent"`
}

// Predefined struct for user
type TriggerEventRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 案例名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间格式：如果选择触发时间：2022年6月21，则设置为20220621
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`
}

type TriggerEventRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 案例名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间格式：如果选择触发时间：2022年6月21，则设置为20220621
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *TriggerEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "Dimension")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TriggerEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TriggerEventResponseParams struct {
	// 成功或者失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *BatchReturn `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TriggerEventResponse struct {
	*tchttp.BaseResponse
	Response *TriggerEventResponseParams `json:"Response"`
}

func (r *TriggerEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Workflow struct {
	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 责任人Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerId *string `json:"OwnerId,omitempty" name:"OwnerId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目标识
	ProjectIdent *string `json:"ProjectIdent,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowDesc *string `json:"WorkflowDesc,omitempty" name:"WorkflowDesc"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 工作流所属用户分组id 若有多个,分号隔开: a;b;c
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// 工作流所属用户分组名称  若有多个,分号隔开: a;b;c
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroupName *string `json:"UserGroupName,omitempty" name:"UserGroupName"`
}