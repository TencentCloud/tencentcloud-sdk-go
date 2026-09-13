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

package v20260715

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AdvancedDependencyConfig struct {
	// 逻辑运算符号OR / AND
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 任务运行条件规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*TaskRunConditionRule `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

type AdvancedParameter struct {
	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type AlarmBrief struct {
	// 告警 ID，创建时无需传入，由服务端生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 告警的监控对象类型，如工作流、任务等，当前支持 1. WORKFLOW 2. TASK
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmMonitorType *string `json:"AlarmMonitorType,omitnil,omitempty" name:"AlarmMonitorType"`

	// 告警组，最多 50 个
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmGroups []*AlarmGroup `json:"AlarmGroups,omitnil,omitempty" name:"AlarmGroups"`

	// 被跳过时免打扰，默认值 false
	// 注意：此字段可能返回 null，表示取不到有效值。
	DoNotDisturbWhenSkipped *bool `json:"DoNotDisturbWhenSkipped,omitnil,omitempty" name:"DoNotDisturbWhenSkipped"`

	// 被手动终止时免打扰，默认值 false
	// 注意：此字段可能返回 null，表示取不到有效值。
	DoNotDisturbWhenManuallyTerminated *bool `json:"DoNotDisturbWhenManuallyTerminated,omitnil,omitempty" name:"DoNotDisturbWhenManuallyTerminated"`

	// 最后一次重试前免打扰，默认值 false
	// 注意：此字段可能返回 null，表示取不到有效值。
	DoNotDisturbUntilTheLastRetry *bool `json:"DoNotDisturbUntilTheLastRetry,omitnil,omitempty" name:"DoNotDisturbUntilTheLastRetry"`
}

type AlarmGroup struct {
	// 通知渠道ID，可通过基础平台通知渠道相关接口获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 通知渠道名称，可以是用户组名称或邮箱地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 是否启用邮件渠道，默认值：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsEmailChannel *bool `json:"IsEmailChannel,omitnil,omitempty" name:"IsEmailChannel"`

	// 一组告警条件，有 启动，成功，失败和任务超时告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmConditions []*string `json:"AlarmConditions,omitnil,omitempty" name:"AlarmConditions"`

	// 通知渠道类型。取值：0 未指定，1 Email，2 Webhook，3 Teams，4 Slack
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelType *int64 `json:"ChannelType,omitnil,omitempty" name:"ChannelType"`
}

type AsyncActionRsp struct {
	// 多个操作项的结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionResults []*RunActionBrief `json:"ActionResults,omitnil,omitempty" name:"ActionResults"`
}

type AsyncOperation struct {
	// 是否异步执行；ZIP 解压创建时为 true
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAsync *bool `json:"IsAsync,omitnil,omitempty" name:"IsAsync"`

	// Workspace 持久化的异步作业 ID，用于查询作业进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 调用方生成的提交幂等与链路追踪标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationId *string `json:"OperationId,omitnil,omitempty" name:"OperationId"`

	// 异步作业状态：0-未指定，1-已受理，2-解压中，3-回调处理中，4-成功，5-部分失败，6-失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type CreateFileRequestParams struct {
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件名，含后缀，最长 255 字节。不能以 . 或 .. 开头/结尾，不能含空格与控制字符</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>父文件夹路径，以 / 开头、末尾不带 /，根目录传 /。来源：ListFiles 接口返回的 Path</p>
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>文件运行配置</p>
	FileConfig *FileConfig `json:"FileConfig,omitnil,omitempty" name:"FileConfig"`

	// <p>绑定的 BundleId。来源：ListBundles 接口返回的 BundleId</p>
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>绑定的 BundleInfo，JSON 字符串</p>
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>文件初始内容。不传则按FileType 生成默认内容</p>
	Storage *FileStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 是否将 Storage 中的 ZIP 文件异步解压创建，默认 false。true 时异步作业由 Workspace 负责全生命周期，响应仅通过 AsyncOperation 返回作业信息（FileId 为空）；作业进度查询由基础平台 WS 接口实现，不在本协议中定义。
	ExtractArchive *bool `json:"ExtractArchive,omitnil,omitempty" name:"ExtractArchive"`
}

type CreateFileRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件名，含后缀，最长 255 字节。不能以 . 或 .. 开头/结尾，不能含空格与控制字符</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>父文件夹路径，以 / 开头、末尾不带 /，根目录传 /。来源：ListFiles 接口返回的 Path</p>
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>文件运行配置</p>
	FileConfig *FileConfig `json:"FileConfig,omitnil,omitempty" name:"FileConfig"`

	// <p>绑定的 BundleId。来源：ListBundles 接口返回的 BundleId</p>
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>绑定的 BundleInfo，JSON 字符串</p>
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>文件初始内容。不传则按FileType 生成默认内容</p>
	Storage *FileStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 是否将 Storage 中的 ZIP 文件异步解压创建，默认 false。true 时异步作业由 Workspace 负责全生命周期，响应仅通过 AsyncOperation 返回作业信息（FileId 为空）；作业进度查询由基础平台 WS 接口实现，不在本协议中定义。
	ExtractArchive *bool `json:"ExtractArchive,omitnil,omitempty" name:"ExtractArchive"`
}

func (r *CreateFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "FileName")
	delete(f, "ParentFolderPath")
	delete(f, "FileType")
	delete(f, "FileConfig")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	delete(f, "Storage")
	delete(f, "ExtractArchive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileResponseParams struct {
	// <p>返回结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *FileInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateFileResponseParams `json:"Response"`
}

func (r *CreateFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流基本信息。必填，其中 WorkflowName 必填且工作空间内唯一</p>
	BaseInfo *WorkflowBaseInfo `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// <p>工作流调度配置</p>
	Trigger []*WorkflowTriggerConfiguration `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>工作流参数列表</p>
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>标签列表</p>
	LabelList []*LabelBrief `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// <p>工作流告警配置</p>
	Alarm *AlarmBrief `json:"Alarm,omitnil,omitempty" name:"Alarm"`

	// <p>监控指标配置。若告警条件中选择了监控告警，则本字段必填</p>
	MonitorMetric *MonitorMetricBrief `json:"MonitorMetric,omitnil,omitempty" name:"MonitorMetric"`

	// <p>工作流高级设置</p>
	AdvanceConfig *WorkflowAdvanceConfig `json:"AdvanceConfig,omitnil,omitempty" name:"AdvanceConfig"`

	// <p>工作流任务列表</p>
	TaskList []*WorkflowTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// <p>BundleId，可通过 Bundle 相关接口获取</p>
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>Bundle信息</p>
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>Git配置ID，可通过 Git 配置相关接口获取</p>
	GitConfigId *string `json:"GitConfigId,omitnil,omitempty" name:"GitConfigId"`

	// <p>Git分支信息</p>
	GitBranch *string `json:"GitBranch,omitnil,omitempty" name:"GitBranch"`
}

type CreateWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流基本信息。必填，其中 WorkflowName 必填且工作空间内唯一</p>
	BaseInfo *WorkflowBaseInfo `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// <p>工作流调度配置</p>
	Trigger []*WorkflowTriggerConfiguration `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>工作流参数列表</p>
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>标签列表</p>
	LabelList []*LabelBrief `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// <p>工作流告警配置</p>
	Alarm *AlarmBrief `json:"Alarm,omitnil,omitempty" name:"Alarm"`

	// <p>监控指标配置。若告警条件中选择了监控告警，则本字段必填</p>
	MonitorMetric *MonitorMetricBrief `json:"MonitorMetric,omitnil,omitempty" name:"MonitorMetric"`

	// <p>工作流高级设置</p>
	AdvanceConfig *WorkflowAdvanceConfig `json:"AdvanceConfig,omitnil,omitempty" name:"AdvanceConfig"`

	// <p>工作流任务列表</p>
	TaskList []*WorkflowTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// <p>BundleId，可通过 Bundle 相关接口获取</p>
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>Bundle信息</p>
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>Git配置ID，可通过 Git 配置相关接口获取</p>
	GitConfigId *string `json:"GitConfigId,omitnil,omitempty" name:"GitConfigId"`

	// <p>Git分支信息</p>
	GitBranch *string `json:"GitBranch,omitnil,omitempty" name:"GitBranch"`
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
	delete(f, "WorkspaceId")
	delete(f, "BaseInfo")
	delete(f, "Trigger")
	delete(f, "ParamList")
	delete(f, "LabelList")
	delete(f, "Alarm")
	delete(f, "MonitorMetric")
	delete(f, "AdvanceConfig")
	delete(f, "TaskList")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	delete(f, "GitConfigId")
	delete(f, "GitBranch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowResponseParams struct {
	// <p>创建工作流响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CreateWorkflowRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CreateWorkflowRsp struct {
	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

// Predefined struct for user
type DeleteFileRequestParams struct {
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID。来源：CreateFile / ListFiles / GetFile 接口返回的 FileId</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

type DeleteFileRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID。来源：CreateFile / ListFiles / GetFile 接口返回的 FileId</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

func (r *DeleteFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "FileId")
	delete(f, "FileType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileResponseParams struct {
	// <p>返回结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DeleteFileResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFileResponseParams `json:"Response"`
}

func (r *DeleteFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFileResult struct {
	// <p>被删除的文件 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>删除是否成功</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteWorkflowRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>待删除的工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DeleteWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>待删除的工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *DeleteWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowResponseParams struct {
	// <p>删除工作流响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DeleteWorkflowRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkflowResponseParams `json:"Response"`
}

func (r *DeleteWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkflowRsp struct {
	// 删除状态，true 表示成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type DependOnBrief struct {
	// 任务ID，可通过 ListWorkflowTasks 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

type FileConfig struct {
	// <p>高级运行参数，变量替换用，map-json String,String</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// <p>执行资源 ID。来源：ListComputeResources 接口返回的 ResourceId</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>默认 catalog</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultCatalog *string `json:"DefaultCatalog,omitnil,omitempty" name:"DefaultCatalog"`

	// <p>默认 schema</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultSchema *string `json:"DefaultSchema,omitnil,omitempty" name:"DefaultSchema"`

	// <p>高级配置，JSON 字符串</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceConfig *string `json:"AdvanceConfig,omitnil,omitempty" name:"AdvanceConfig"`

	// <p>扩展参数，JSON 字符串</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraParams *string `json:"ExtraParams,omitnil,omitempty" name:"ExtraParams"`

	// <p>Notebook 交互控件定义，JSON 字符串</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Widgets *string `json:"Widgets,omitnil,omitempty" name:"Widgets"`

	// <p>各单元格输出配置。仅 Get 出参返回，入参忽略</p>
	OutputConf []*FileOutputConf `json:"OutputConf,omitnil,omitempty" name:"OutputConf"`

	// <p>SQL脚本语法标记</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlSyntax *string `json:"SqlSyntax,omitnil,omitempty" name:"SqlSyntax"`

	// <p>平台集群id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type FileInfo struct {
	// <p>主账号 AppId</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>工作空间 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件名，含后缀</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>文件在工作空间中的完整路径，以 / 开头，如 /etl/daily/demo.ipynb</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// <p>文件运行配置</p>
	FileConfig *FileConfig `json:"FileConfig,omitnil,omitempty" name:"FileConfig"`

	// <p>绑定的 BundleId</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>绑定的 BundleInfo，JSON 字符串</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>文件状态。active=正常，deleted=已删除</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>文件负责人用户名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserName *string `json:"OwnerUserName,omitnil,omitempty" name:"OwnerUserName"`

	// <p>创建人子账号 Uin</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// <p>最近更新人子账号 Uin</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// <p>创建时间，毫秒级时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>最近更新时间，毫秒级时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>文件存储信息。仅当请求 IncludeContent=true 时返回内容</p>
	Storage *FileStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>当前调用方对该文件的权限点列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions *string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// <p>是否已发布</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseStatus *bool `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`

	// <p>资源模式。1=分布式，2=单节点</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceMode *int64 `json:"ResourceMode,omitnil,omitempty" name:"ResourceMode"`

	// ZIP 异步创建时透传 Workspace 作业信息；普通同步创建或其他复用该返回结构的接口不设置该字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncOperation *AsyncOperation `json:"AsyncOperation,omitnil,omitempty" name:"AsyncOperation"`
}

type FileOutputConf struct {
	// 单元格 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CellId *string `json:"CellId,omitnil,omitempty" name:"CellId"`

	// Dashboard 图表配置，JSON 字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	DashboardConf *string `json:"DashboardConf,omitnil,omitempty" name:"DashboardConf"`

	// 执行结果文件的预签名下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputPath *string `json:"OutputPath,omitnil,omitempty" name:"OutputPath"`
}

type FileStorage struct {
	// 存储类型
	StorageType *int64 `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 存储路径
	StoragePath *string `json:"StoragePath,omitnil,omitempty" name:"StoragePath"`

	// 文件内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type GetFileRequestParams struct {
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID。来源：CreateFile / ListFiles 接口返回的 FileId。与 FilePath 二选一</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>是否返回文件内容。true 时 Storage.Content 返回 base64 内容，默认 false</p>
	IncludeContent *bool `json:"IncludeContent,omitnil,omitempty" name:"IncludeContent"`

	// <p>文件版本 ID。来源：ListFileVersions 接口返回的 VersionId。不传则读取最新版本</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// <p>文件完整路径，以 / 开头，如 /etl/daily/demo.ipynb。与 FileId 二选一</p>
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

type GetFileRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID。来源：CreateFile / ListFiles 接口返回的 FileId。与 FilePath 二选一</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>是否返回文件内容。true 时 Storage.Content 返回 base64 内容，默认 false</p>
	IncludeContent *bool `json:"IncludeContent,omitnil,omitempty" name:"IncludeContent"`

	// <p>文件版本 ID。来源：ListFileVersions 接口返回的 VersionId。不传则读取最新版本</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// <p>文件完整路径，以 / 开头，如 /etl/daily/demo.ipynb。与 FileId 二选一</p>
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

func (r *GetFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "FileId")
	delete(f, "FileType")
	delete(f, "IncludeContent")
	delete(f, "VersionId")
	delete(f, "FilePath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFileResponseParams struct {
	// <p>返回结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *FileInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFileResponse struct {
	*tchttp.BaseResponse
	Response *GetFileResponseParams `json:"Response"`
}

func (r *GetFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWorkflowRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type GetWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *GetWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWorkflowResponseParams struct {
	// <p>获取工作流详细信息响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *GetWorkflowRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *GetWorkflowResponseParams `json:"Response"`
}

func (r *GetWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkflowRsp struct {
	// <p>工作空间ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流基本信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseInfo *WorkflowBaseInfoDetail `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// <p>工作流调度配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trigger []*WorkflowTriggerConfiguration `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>工作流参数列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>标签列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelList []*LabelBrief `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// <p>工作流告警配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alarm *AlarmBrief `json:"Alarm,omitnil,omitempty" name:"Alarm"`

	// <p>监控指标配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorMetric *MonitorMetricBrief `json:"MonitorMetric,omitnil,omitempty" name:"MonitorMetric"`

	// <p>工作流高级设置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceConfig *WorkflowAdvanceConfig `json:"AdvanceConfig,omitnil,omitempty" name:"AdvanceConfig"`

	// <p>工作流任务列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskList []*WorkflowTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// <p>工作流绑定的 Bundle唯一标识，未绑定时为空</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>Bundle信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>Git配置ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitConfigId *string `json:"GitConfigId,omitnil,omitempty" name:"GitConfigId"`

	// <p>Git分支信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitBranch *string `json:"GitBranch,omitnil,omitempty" name:"GitBranch"`
}

// Predefined struct for user
type GetWorkflowRunRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流运行ID，可通过 ListWorkflowRuns 获取。必填</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

type GetWorkflowRunRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流运行ID，可通过 ListWorkflowRuns 获取。必填</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

func (r *GetWorkflowRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowRunId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWorkflowRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWorkflowRunResponseParams struct {
	// <p>查询工作流运行详情响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *GetWorkflowRunRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWorkflowRunResponse struct {
	*tchttp.BaseResponse
	Response *GetWorkflowRunResponseParams `json:"Response"`
}

func (r *GetWorkflowRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkflowRunRsp struct {
	// 工作流运行信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRun *WorkflowRun `json:"WorkflowRun,omitnil,omitempty" name:"WorkflowRun"`
}

// Predefined struct for user
type GetWorkflowTaskRunRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>任务运行ID，可通过 ListWorkflowTaskRuns 获取。必填</p>
	WorkflowTaskRunId *string `json:"WorkflowTaskRunId,omitnil,omitempty" name:"WorkflowTaskRunId"`

	// <p>内嵌工作流任务运行列表选项（仅限 FOR_EACH 任务使用）。非必填</p>
	InnerWorkflowTaskRunListOption *InnerWorkflowTaskRunListOption `json:"InnerWorkflowTaskRunListOption,omitnil,omitempty" name:"InnerWorkflowTaskRunListOption"`
}

type GetWorkflowTaskRunRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>任务运行ID，可通过 ListWorkflowTaskRuns 获取。必填</p>
	WorkflowTaskRunId *string `json:"WorkflowTaskRunId,omitnil,omitempty" name:"WorkflowTaskRunId"`

	// <p>内嵌工作流任务运行列表选项（仅限 FOR_EACH 任务使用）。非必填</p>
	InnerWorkflowTaskRunListOption *InnerWorkflowTaskRunListOption `json:"InnerWorkflowTaskRunListOption,omitnil,omitempty" name:"InnerWorkflowTaskRunListOption"`
}

func (r *GetWorkflowTaskRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowTaskRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowTaskRunId")
	delete(f, "InnerWorkflowTaskRunListOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWorkflowTaskRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWorkflowTaskRunResponseParams struct {
	// <p>查询任务运行详情响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *GetWorkflowTaskRunRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWorkflowTaskRunResponse struct {
	*tchttp.BaseResponse
	Response *GetWorkflowTaskRunResponseParams `json:"Response"`
}

func (r *GetWorkflowTaskRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowTaskRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWorkflowTaskRunRsp struct {
	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务运行ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowTaskRunId *string `json:"WorkflowTaskRunId,omitnil,omitempty" name:"WorkflowTaskRunId"`

	// 运行状态。取值参考工作流任务运行状态枚举，如 Pending / Running / Succeeded / Failed / Killed
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunState *string `json:"RunState,omitnil,omitempty" name:"RunState"`

	// 工作空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流运行ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeName *string `json:"TaskTypeName,omitnil,omitempty" name:"TaskTypeName"`

	// 任务版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskVersionId *string `json:"TaskVersionId,omitnil,omitempty" name:"TaskVersionId"`

	// 触发类型 (参考SchedulerTriggerType枚举)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 所属资源组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeString *string `json:"ErrorCodeString,omitnil,omitempty" name:"ErrorCodeString"`

	// 运行用户UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserUin *string `json:"RunUserUin,omitnil,omitempty" name:"RunUserUin"`

	// 运行用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserName *string `json:"RunUserName,omitnil,omitempty" name:"RunUserName"`

	// 创建人UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 执行平台执行ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 创建时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 依赖任务完成时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependenceFinishedTime *string `json:"DependenceFinishedTime,omitnil,omitempty" name:"DependenceFinishedTime"`

	// 运行开始时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStartTime *string `json:"RunStartTime,omitnil,omitempty" name:"RunStartTime"`

	// 运行结束时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunEndTime *string `json:"RunEndTime,omitnil,omitempty" name:"RunEndTime"`

	// 运行时长，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunCostTime *string `json:"RunCostTime,omitnil,omitempty" name:"RunCostTime"`

	// 等待时长（依赖就绪到开始运行的等待耗时），单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaitTime *string `json:"WaitTime,omitnil,omitempty" name:"WaitTime"`

	// 下发执行平台时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssueTime *string `json:"IssueTime,omitnil,omitempty" name:"IssueTime"`

	// 时区
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// 依赖上游任务ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnList []*string `json:"DependOnList,omitnil,omitempty" name:"DependOnList"`

	// 运行参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunParams *string `json:"RunParams,omitnil,omitempty" name:"RunParams"`

	// 任务扩展信息，包含脚本路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeExtensions *string `json:"TaskTypeExtensions,omitnil,omitempty" name:"TaskTypeExtensions"`

	// 任务X坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftCoordinate *float64 `json:"LeftCoordinate,omitnil,omitempty" name:"LeftCoordinate"`

	// 任务Y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopCoordinate *float64 `json:"TopCoordinate,omitnil,omitempty" name:"TopCoordinate"`

	// 重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 重跑次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RerunTimes *int64 `json:"RerunTimes,omitnil,omitempty" name:"RerunTimes"`

	// 是否最新一次运行
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLatestRun *bool `json:"IsLatestRun,omitnil,omitempty" name:"IsLatestRun"`

	// 资源组信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupInfoList []*ResourceGroupInfo `json:"ResourceGroupInfoList,omitnil,omitempty" name:"ResourceGroupInfoList"`

	// 错误消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 运行结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunResult *string `json:"RunResult,omitnil,omitempty" name:"RunResult"`

	// 内嵌工作流任务运行详情（仅限 FOR_EACH 任务，其他任务类型不返回该字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerWorkflowTaskRun *InnerWorkflowTaskRun `json:"InnerWorkflowTaskRun,omitnil,omitempty" name:"InnerWorkflowTaskRun"`
}

type InnerWorkflowTaskBrief struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeName *string `json:"TaskTypeName,omitnil,omitempty" name:"TaskTypeName"`
}

type InnerWorkflowTaskRun struct {
	// 当前页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 迭代运行列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*InnerWorkflowTaskRunIteration `json:"Items,omitnil,omitempty" name:"Items"`

	// 迭代次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IterationCount *int64 `json:"IterationCount,omitnil,omitempty" name:"IterationCount"`

	// 失败次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureCount *int64 `json:"FailureCount,omitnil,omitempty" name:"FailureCount"`

	// 成功次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 内嵌工作流ID，可通过 ListWorkflows 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerWorkflowId *string `json:"InnerWorkflowId,omitnil,omitempty" name:"InnerWorkflowId"`

	// 内嵌任务ID，可通过 ListWorkflowTasks 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerTaskId *string `json:"InnerTaskId,omitnil,omitempty" name:"InnerTaskId"`

	// 内嵌任务运行状态数量统计（实例业务枚举键值对列表）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerTaskRunBizEnumInfos []*ScheduleBizEnumBrief `json:"InnerTaskRunBizEnumInfos,omitnil,omitempty" name:"InnerTaskRunBizEnumInfos"`
}

type InnerWorkflowTaskRunIteration struct {
	// <p>内嵌工作流运行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>迭代序号</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IterationIndex *string `json:"IterationIndex,omitnil,omitempty" name:"IterationIndex"`

	// <p>运行开始时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStartTime *string `json:"RunStartTime,omitnil,omitempty" name:"RunStartTime"`

	// <p>运行结束时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunEndTime *string `json:"RunEndTime,omitnil,omitempty" name:"RunEndTime"`

	// <p>运行状态（参考工作流运行状态枚举）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunState *string `json:"RunState,omitnil,omitempty" name:"RunState"`

	// <p>运行时长，单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunCostTime *string `json:"RunCostTime,omitnil,omitempty" name:"RunCostTime"`

	// <p>运行参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowParams *string `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// <p>错误码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeString *string `json:"ErrorCodeString,omitnil,omitempty" name:"ErrorCodeString"`

	// <p>内嵌工作流内部的任务运行</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerTaskRun *InnerWorkflowTaskRunIterationBrief `json:"InnerTaskRun,omitnil,omitempty" name:"InnerTaskRun"`
}

type InnerWorkflowTaskRunIterationBrief struct {
	// <p>任务运行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowTaskRunId *string `json:"WorkflowTaskRunId,omitnil,omitempty" name:"WorkflowTaskRunId"`

	// <p>迭代序号</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IterationIndex *string `json:"IterationIndex,omitnil,omitempty" name:"IterationIndex"`

	// <p>运行开始时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStartTime *string `json:"RunStartTime,omitnil,omitempty" name:"RunStartTime"`

	// <p>运行结束时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunEndTime *string `json:"RunEndTime,omitnil,omitempty" name:"RunEndTime"`

	// <p>运行状态</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunState *string `json:"RunState,omitnil,omitempty" name:"RunState"`

	// <p>运行时长，单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunCostTime *string `json:"RunCostTime,omitnil,omitempty" name:"RunCostTime"`

	// <p>运行参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskParams *string `json:"TaskParams,omitnil,omitempty" name:"TaskParams"`

	// <p>错误码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeString *string `json:"ErrorCodeString,omitnil,omitempty" name:"ErrorCodeString"`
}

type InnerWorkflowTaskRunListOption struct {
	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>迭代运行状态，精确匹配。非必填，多选（多个值之间为 OR 关系）。</p><p>可填 SUCCESS / FAILED 等，具体参考本接口出参 InnerWorkflowTaskRunIteration.RunState 字段返回值。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStates []*string `json:"RunStates,omitnil,omitempty" name:"RunStates"`
}

// Predefined struct for user
type KillWorkflowRunRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>待终止的工作流运行ID列表，可通过 ListWorkflowRuns 获取</p>
	WorkflowRunIds []*string `json:"WorkflowRunIds,omitnil,omitempty" name:"WorkflowRunIds"`

	// <p>是否终止该工作流下所有未进入终态的运行。非必填，默认 false</p>
	KillAllRuns *bool `json:"KillAllRuns,omitnil,omitempty" name:"KillAllRuns"`

	// <p>是否只终止处于等待中（Pending）状态的运行。非必填，默认 false</p>
	OnlyKillPendingRuns *bool `json:"OnlyKillPendingRuns,omitnil,omitempty" name:"OnlyKillPendingRuns"`
}

type KillWorkflowRunRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>待终止的工作流运行ID列表，可通过 ListWorkflowRuns 获取</p>
	WorkflowRunIds []*string `json:"WorkflowRunIds,omitnil,omitempty" name:"WorkflowRunIds"`

	// <p>是否终止该工作流下所有未进入终态的运行。非必填，默认 false</p>
	KillAllRuns *bool `json:"KillAllRuns,omitnil,omitempty" name:"KillAllRuns"`

	// <p>是否只终止处于等待中（Pending）状态的运行。非必填，默认 false</p>
	OnlyKillPendingRuns *bool `json:"OnlyKillPendingRuns,omitnil,omitempty" name:"OnlyKillPendingRuns"`
}

func (r *KillWorkflowRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillWorkflowRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	delete(f, "WorkflowRunIds")
	delete(f, "KillAllRuns")
	delete(f, "OnlyKillPendingRuns")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillWorkflowRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillWorkflowRunResponseParams struct {
	// <p>终止工作流的运行响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AsyncActionRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type KillWorkflowRunResponse struct {
	*tchttp.BaseResponse
	Response *KillWorkflowRunResponseParams `json:"Response"`
}

func (r *KillWorkflowRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillWorkflowRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelBrief struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelKey *string `json:"LabelKey,omitnil,omitempty" name:"LabelKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValue *string `json:"LabelValue,omitnil,omitempty" name:"LabelValue"`

	// 标签名称ID，可通过标签相关接口获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelKeyId *string `json:"LabelKeyId,omitnil,omitempty" name:"LabelKeyId"`

	// 标签值ID，可通过标签相关接口获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValueId *string `json:"LabelValueId,omitnil,omitempty" name:"LabelValueId"`
}

// Predefined struct for user
type ListWorkflowRunsRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>工作流ID，精确匹配。非必填，单值</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流名称关键字，对 WorkflowName 做模糊匹配。非必填，单值</p>
	WorkflowNameKeyword *string `json:"WorkflowNameKeyword,omitnil,omitempty" name:"WorkflowNameKeyword"`

	// <p>运行创建时间下界，范围匹配（CreateTime &gt;= 本值），单位：毫秒时间戳。<br>非必填，单值，对应出参 WorkflowRun.CreateTime</p>
	CreateStartTime *string `json:"CreateStartTime,omitnil,omitempty" name:"CreateStartTime"`

	// <p>运行创建时间上界，范围匹配（CreateTime &lt;= 本值），单位：毫秒时间戳。<br>非必填，单值，对应出参 WorkflowRun.CreateTime</p>
	CreateEndTime *string `json:"CreateEndTime,omitnil,omitempty" name:"CreateEndTime"`

	// <p>运行状态，精确匹配。非必填，多选（多个值之间为 OR 关系）。</p><p>可填 SUCCESS / FAILED 等，具体参考本接口出参 WorkflowRun.RunState 字段返回值。</p>
	RunStates []*string `json:"RunStates,omitnil,omitempty" name:"RunStates"`

	// <p>错误码，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	ErrorCodeStrings []*string `json:"ErrorCodeStrings,omitnil,omitempty" name:"ErrorCodeStrings"`

	// <p>运行人UIN，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	RunUserUins []*string `json:"RunUserUins,omitnil,omitempty" name:"RunUserUins"`

	// <p>标签名称ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelKeyIds []*string `json:"LabelKeyIds,omitnil,omitempty" name:"LabelKeyIds"`

	// <p>标签值ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelValueIds []*string `json:"LabelValueIds,omitnil,omitempty" name:"LabelValueIds"`

	// <p>排序条件，多个之间按数组顺序表示优先级。非必填，默认按 CreateTime Desc。<br>可排序字段白名单：CreateTime、EndTime、RunCostTime</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`
}

type ListWorkflowRunsRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>工作流ID，精确匹配。非必填，单值</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流名称关键字，对 WorkflowName 做模糊匹配。非必填，单值</p>
	WorkflowNameKeyword *string `json:"WorkflowNameKeyword,omitnil,omitempty" name:"WorkflowNameKeyword"`

	// <p>运行创建时间下界，范围匹配（CreateTime &gt;= 本值），单位：毫秒时间戳。<br>非必填，单值，对应出参 WorkflowRun.CreateTime</p>
	CreateStartTime *string `json:"CreateStartTime,omitnil,omitempty" name:"CreateStartTime"`

	// <p>运行创建时间上界，范围匹配（CreateTime &lt;= 本值），单位：毫秒时间戳。<br>非必填，单值，对应出参 WorkflowRun.CreateTime</p>
	CreateEndTime *string `json:"CreateEndTime,omitnil,omitempty" name:"CreateEndTime"`

	// <p>运行状态，精确匹配。非必填，多选（多个值之间为 OR 关系）。</p><p>可填 SUCCESS / FAILED 等，具体参考本接口出参 WorkflowRun.RunState 字段返回值。</p>
	RunStates []*string `json:"RunStates,omitnil,omitempty" name:"RunStates"`

	// <p>错误码，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	ErrorCodeStrings []*string `json:"ErrorCodeStrings,omitnil,omitempty" name:"ErrorCodeStrings"`

	// <p>运行人UIN，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	RunUserUins []*string `json:"RunUserUins,omitnil,omitempty" name:"RunUserUins"`

	// <p>标签名称ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelKeyIds []*string `json:"LabelKeyIds,omitnil,omitempty" name:"LabelKeyIds"`

	// <p>标签值ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelValueIds []*string `json:"LabelValueIds,omitnil,omitempty" name:"LabelValueIds"`

	// <p>排序条件，多个之间按数组顺序表示优先级。非必填，默认按 CreateTime Desc。<br>可排序字段白名单：CreateTime、EndTime、RunCostTime</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`
}

func (r *ListWorkflowRunsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "WorkflowId")
	delete(f, "WorkflowNameKeyword")
	delete(f, "CreateStartTime")
	delete(f, "CreateEndTime")
	delete(f, "RunStates")
	delete(f, "ErrorCodeStrings")
	delete(f, "RunUserUins")
	delete(f, "LabelKeyIds")
	delete(f, "LabelValueIds")
	delete(f, "OrderBys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkflowRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowRunsResponseParams struct {
	// <p>工作流运行列表响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ListWorkflowRunsRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkflowRunsResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkflowRunsResponseParams `json:"Response"`
}

func (r *ListWorkflowRunsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkflowRunsRsp struct {
	// 当前页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 工作流运行列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*WorkflowRun `json:"Items,omitnil,omitempty" name:"Items"`

	// 工作流运行状态数量统计。
	// 统计口径为当前筛选条件下的全量数据，不受 PageNumber / PageSize 影响
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizStateEnumInfos []*ScheduleBizEnumBrief `json:"BizStateEnumInfos,omitnil,omitempty" name:"BizStateEnumInfos"`

	// 工作流运行错误码数量统计。
	// 统计口径为当前筛选条件下的全量数据，不受 PageNumber / PageSize 影响
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizErrorCodeEnumInfos []*ScheduleBizEnumBrief `json:"BizErrorCodeEnumInfos,omitnil,omitempty" name:"BizErrorCodeEnumInfos"`
}

// Predefined struct for user
type ListWorkflowTaskRunsRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>任务ID，可通过 ListWorkflowTasks 获取。非必填，精确匹配。与 WorkflowRunId 至少传一个：仅传 TaskId 时查询该任务的全部运行历史。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>工作流运行ID，可通过 ListWorkflowRuns 获取。非必填，精确匹配。与 TaskId 至少传一个：仅传 WorkflowRunId 时查询该次工作流运行下的全部任务运行。</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListWorkflowTaskRunsRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>任务ID，可通过 ListWorkflowTasks 获取。非必填，精确匹配。与 WorkflowRunId 至少传一个：仅传 TaskId 时查询该任务的全部运行历史。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>工作流运行ID，可通过 ListWorkflowRuns 获取。非必填，精确匹配。与 TaskId 至少传一个：仅传 WorkflowRunId 时查询该次工作流运行下的全部任务运行。</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListWorkflowTaskRunsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowTaskRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "TaskId")
	delete(f, "WorkflowRunId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkflowTaskRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowTaskRunsResponseParams struct {
	// <p>查询工作流任务历史运行列表响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ListWorkflowTaskRunsRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkflowTaskRunsResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkflowTaskRunsResponseParams `json:"Response"`
}

func (r *ListWorkflowTaskRunsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowTaskRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkflowTaskRunsRsp struct {
	// 当前页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 任务运行历史列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*WorkflowTaskRun `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type ListWorkflowsRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>工作流名称关键字，对 WorkflowName 做模糊匹配。非必填，单值</p>
	WorkflowNameKeyword *string `json:"WorkflowNameKeyword,omitnil,omitempty" name:"WorkflowNameKeyword"`

	// <p>工作流名称，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	WorkflowNames []*string `json:"WorkflowNames,omitnil,omitempty" name:"WorkflowNames"`

	// <p>工作流ID，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// <p>工作流运行人UIN，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	RunUserUins []*string `json:"RunUserUins,omitnil,omitempty" name:"RunUserUins"`

	// <p>标签名称ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelKeyIds []*string `json:"LabelKeyIds,omitnil,omitempty" name:"LabelKeyIds"`

	// <p>标签值ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelValueIds []*string `json:"LabelValueIds,omitnil,omitempty" name:"LabelValueIds"`

	// <p>快速筛选类型。非必填，单值</p><p>对齐老云 API（wedata/2025-10-10）文档示例值：</p><ul><li>MY_FAVORITE：我收藏的</li><li>MY_OWNER：我负责的</li><li>MY_AUTHORITY：我有权限</li><li>WorkflowId：支持多个工作流ID筛选</li></ul><p>后端实现现状：当前仅 MY_FAVORITE 生效（设置 favoriteUserUin 过滤当前用户收藏），MY_OWNER / MY_AUTHORITY 暂未在 Service 层实现，传入会被忽略（按全量返回）。</p>
	QuickSelectionType *string `json:"QuickSelectionType,omitnil,omitempty" name:"QuickSelectionType"`

	// <p>排序条件，多个之间按数组顺序表示优先级。非必填。<br>可排序字段白名单：CreateTime</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`
}

type ListWorkflowsRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>分页页码，从 1 开始。非必填，默认 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小。非必填，默认 10，取值范围 [10, 200]</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>工作流名称关键字，对 WorkflowName 做模糊匹配。非必填，单值</p>
	WorkflowNameKeyword *string `json:"WorkflowNameKeyword,omitnil,omitempty" name:"WorkflowNameKeyword"`

	// <p>工作流名称，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	WorkflowNames []*string `json:"WorkflowNames,omitnil,omitempty" name:"WorkflowNames"`

	// <p>工作流ID，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// <p>工作流运行人UIN，精确匹配。非必填，多选（多个值之间为 OR 关系）</p>
	RunUserUins []*string `json:"RunUserUins,omitnil,omitempty" name:"RunUserUins"`

	// <p>标签名称ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelKeyIds []*string `json:"LabelKeyIds,omitnil,omitempty" name:"LabelKeyIds"`

	// <p>标签值ID，精确匹配，可通过标签相关接口获取。非必填，多选（多个值之间为 OR 关系）</p>
	LabelValueIds []*string `json:"LabelValueIds,omitnil,omitempty" name:"LabelValueIds"`

	// <p>快速筛选类型。非必填，单值</p><p>对齐老云 API（wedata/2025-10-10）文档示例值：</p><ul><li>MY_FAVORITE：我收藏的</li><li>MY_OWNER：我负责的</li><li>MY_AUTHORITY：我有权限</li><li>WorkflowId：支持多个工作流ID筛选</li></ul><p>后端实现现状：当前仅 MY_FAVORITE 生效（设置 favoriteUserUin 过滤当前用户收藏），MY_OWNER / MY_AUTHORITY 暂未在 Service 层实现，传入会被忽略（按全量返回）。</p>
	QuickSelectionType *string `json:"QuickSelectionType,omitnil,omitempty" name:"QuickSelectionType"`

	// <p>排序条件，多个之间按数组顺序表示优先级。非必填。<br>可排序字段白名单：CreateTime</p>
	OrderBys []*OrderBy `json:"OrderBys,omitnil,omitempty" name:"OrderBys"`
}

func (r *ListWorkflowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "WorkflowNameKeyword")
	delete(f, "WorkflowNames")
	delete(f, "WorkflowIds")
	delete(f, "RunUserUins")
	delete(f, "LabelKeyIds")
	delete(f, "LabelValueIds")
	delete(f, "QuickSelectionType")
	delete(f, "OrderBys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkflowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowsResponseParams struct {
	// <p>查询工作流列表响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ListWorkflowsRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkflowsResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkflowsResponseParams `json:"Response"`
}

func (r *ListWorkflowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkflowsRsp struct {
	// 当前页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 工作流列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*WorkflowBrief `json:"Items,omitnil,omitempty" name:"Items"`
}

type MonitorMetricBrief struct {
	// 监控指标 ID，创建时无需传入，由服务端生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorMetricId *string `json:"MonitorMetricId,omitnil,omitempty" name:"MonitorMetricId"`

	// 告警的监控对象类型，如工作流、任务等，当前支持 1. WORKFLOW 2. TASK
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmMonitorType *string `json:"AlarmMonitorType,omitnil,omitempty" name:"AlarmMonitorType"`

	// 监控指标列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*MonitorMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`
}

type MonitorMetricItem struct {
	// 监控指标类型,有三种类型：1. RUN_DURATION（运行时长）2. WAIT_DURATION（等待时长）3. COMPLETION_TIME（完成时间）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricType *string `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// 警告阈值，单位为毫秒级别，对于COMPLETION_TIME:从当日时间点00:00起算
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarningThreshold *string `json:"WarningThreshold,omitnil,omitempty" name:"WarningThreshold"`

	// 超时阈值，单位为毫秒级别，对于COMPLETION_TIME:从当日时间点00:00起算
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutThreshold *string `json:"TimeoutThreshold,omitnil,omitempty" name:"TimeoutThreshold"`
}

type OrderBy struct {
	// 排序方向，Asc（升序）或 Desc（降序），大小写不敏感
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 排序字段名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ParamInfo struct {
	// 参数ID，创建时无需传入，由服务端生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamId *string `json:"ParamId,omitnil,omitempty" name:"ParamId"`

	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

// Predefined struct for user
type RerunWorkflowRunRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流运行ID，可通过 ListWorkflowRuns 获取。必填</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>运行类型。必填。取值：1 普通运行，2 高级运行</p>
	RunType *int64 `json:"RunType,omitnil,omitempty" name:"RunType"`

	// <p>运行类型为高级运行时填写的自定义运行参数</p>
	AdvancedParams []*TaskSchedulingParameterBrief `json:"AdvancedParams,omitnil,omitempty" name:"AdvancedParams"`

	// <p>本次需要重跑指定的任务ID集合，可通过 ListWorkflowTasks 获取，不传默认重跑该工作流下所有任务</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

type RerunWorkflowRunRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流运行ID，可通过 ListWorkflowRuns 获取。必填</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>运行类型。必填。取值：1 普通运行，2 高级运行</p>
	RunType *int64 `json:"RunType,omitnil,omitempty" name:"RunType"`

	// <p>运行类型为高级运行时填写的自定义运行参数</p>
	AdvancedParams []*TaskSchedulingParameterBrief `json:"AdvancedParams,omitnil,omitempty" name:"AdvancedParams"`

	// <p>本次需要重跑指定的任务ID集合，可通过 ListWorkflowTasks 获取，不传默认重跑该工作流下所有任务</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

func (r *RerunWorkflowRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunWorkflowRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	delete(f, "WorkflowRunId")
	delete(f, "RunType")
	delete(f, "AdvancedParams")
	delete(f, "TaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RerunWorkflowRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RerunWorkflowRunResponseParams struct {
	// <p>重跑工作流响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AsyncActionRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RerunWorkflowRunResponse struct {
	*tchttp.BaseResponse
	Response *RerunWorkflowRunResponseParams `json:"Response"`
}

func (r *RerunWorkflowRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunWorkflowRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceGroupInfo struct {
	// 资源组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 资源组状态
	// COMPUTE_RESOURCE_STATUS_UNSPECIFIED 未指定
	// COMPUTE_RESOURCE_STATUS_PENDING_CREATE 待创建
	// COMPUTE_RESOURCE_STATUS_CREATING 创建中
	// COMPUTE_RESOURCE_STATUS_RUNNING 运行中
	// COMPUTE_RESOURCE_STATUS_STOPPED 已停止
	// COMPUTE_RESOURCE_STATUS_STOPPING 停止中
	// COMPUTE_RESOURCE_STATUS_STARTING 启动中
	// COMPUTE_RESOURCE_STATUS_UPDATING 更新中
	// COMPUTE_RESOURCE_STATUS_DELETING 删除中
	// COMPUTE_RESOURCE_STATUS_DELETED 已删除
	// COMPUTE_RESOURCE_STATUS_FAILED  失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupStatus *string `json:"ResourceGroupStatus,omitnil,omitempty" name:"ResourceGroupStatus"`
}

type RunActionBrief struct {
	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 操作动作ID，用于追踪具体的执行动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunActionId *string `json:"RunActionId,omitnil,omitempty" name:"RunActionId"`

	// 失败错误信息，操作失败时返回具体的错误描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 操作状态，true 表示成功，false 表示失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpStatus *bool `json:"OpStatus,omitnil,omitempty" name:"OpStatus"`

	// 工作流运行ID。重跑 / 终止场景返回被操作的运行ID；运行工作流场景为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

// Predefined struct for user
type RunWorkflowRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>运行类型。必填。取值：1 普通运行，2 高级运行</p>
	RunType *int64 `json:"RunType,omitnil,omitempty" name:"RunType"`

	// <p>运行类型为高级运行时填写的自定义运行参数</p>
	AdvancedParams []*TaskSchedulingParameterBrief `json:"AdvancedParams,omitnil,omitempty" name:"AdvancedParams"`

	// <p>本次需要运行指定的任务ID集合，可通过 ListWorkflowTasks 获取，不传默认运行该工作流下所有任务</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>幂等令牌。非必填，相同令牌的重复请求只会触发一次运行</p>
	IdempotencyToken *string `json:"IdempotencyToken,omitnil,omitempty" name:"IdempotencyToken"`
}

type RunWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>运行类型。必填。取值：1 普通运行，2 高级运行</p>
	RunType *int64 `json:"RunType,omitnil,omitempty" name:"RunType"`

	// <p>运行类型为高级运行时填写的自定义运行参数</p>
	AdvancedParams []*TaskSchedulingParameterBrief `json:"AdvancedParams,omitnil,omitempty" name:"AdvancedParams"`

	// <p>本次需要运行指定的任务ID集合，可通过 ListWorkflowTasks 获取，不传默认运行该工作流下所有任务</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>幂等令牌。非必填，相同令牌的重复请求只会触发一次运行</p>
	IdempotencyToken *string `json:"IdempotencyToken,omitnil,omitempty" name:"IdempotencyToken"`
}

func (r *RunWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	delete(f, "RunType")
	delete(f, "AdvancedParams")
	delete(f, "TaskIds")
	delete(f, "IdempotencyToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunWorkflowResponseParams struct {
	// <p>运行工作流响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AsyncActionRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *RunWorkflowResponseParams `json:"Response"`
}

func (r *RunWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScheduleBizEnumBrief struct {
	// 枚举标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelKey *string `json:"LabelKey,omitnil,omitempty" name:"LabelKey"`

	// 枚举标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValue *string `json:"LabelValue,omitnil,omitempty" name:"LabelValue"`

	// 枚举项统计数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type TaskRetryStrategy struct {
	// 最多重试次数，默认3
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil,omitempty" name:"MaxRetryTimes"`

	// 重试之间等待时间，默认5
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryBetweenWaitTime *int64 `json:"RetryBetweenWaitTime,omitnil,omitempty" name:"RetryBetweenWaitTime"`

	// 重试之间等待时间单位
	// 毫秒：MILLISECOND秒：SECOND分钟（默认）：MINUTE小时：HOUR
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryBetweenWaitTimeUnit *string `json:"RetryBetweenWaitTimeUnit,omitnil,omitempty" name:"RetryBetweenWaitTimeUnit"`

	// 任务运行失败时重试开关，默认为true
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRunFailureRetrySwitch *bool `json:"TaskRunFailureRetrySwitch,omitnil,omitempty" name:"TaskRunFailureRetrySwitch"`

	// 任务运行超时时重试开关，默认为false
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRunTimeoutRetrySwitch *bool `json:"TaskRunTimeoutRetrySwitch,omitnil,omitempty" name:"TaskRunTimeoutRetrySwitch"`
}

type TaskRunConditionRule struct {
	// 上游任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamTaskId *string `json:"UpstreamTaskId,omitnil,omitempty" name:"UpstreamTaskId"`

	// 上游任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamTaskName *string `json:"UpstreamTaskName,omitnil,omitempty" name:"UpstreamTaskName"`

	// 任务可运行条件
	// 支持的状态值： - SUCCESS: 成功 - FAILED: 失败 - UPSTREAM_FAILED: 上游失败 - EXCLUDED: 排除运行
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowedStates []*string `json:"AllowedStates,omitnil,omitempty" name:"AllowedStates"`
}

type TaskSchedulingParameterBrief struct {
	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type TaskType struct {
	// <p>任务类型：SQL：用于执行SQL查询和数据处理操作；DATA_INTEGRATION：用于离线数据接入操作；NOTEBOOK：用于运行Notebook脚本；RUN_WORKFLOW：用于执行嵌套工作流；PYTHON：用于运行Python脚本；RAY_JOB：用于运行Ray作业；DATA_QUALITY：用于数据质量监控；IF_ELSE：用于条件分支判断；FOR_EACH：用于循环遍历执行；</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeName *string `json:"TaskTypeName,omitnil,omitempty" name:"TaskTypeName"`

	// <p>Notebook 类型扩展信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Notebook *TaskTypeNotebookExt `json:"Notebook,omitnil,omitempty" name:"Notebook"`

	// <p>任务扩展属性列表，具体填写参考 ListWorkflowTaskTypeProperties 接口</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypePropertyList []*TaskTypeProperty `json:"TaskTypePropertyList,omitnil,omitempty" name:"TaskTypePropertyList"`

	// <p>运行时属性列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimePropertyList []*TaskTypeProperty `json:"RuntimePropertyList,omitnil,omitempty" name:"RuntimePropertyList"`
}

type TaskTypeNotebookExt struct {
	// 脚本来源。取值：SCRIPT_SOURCE_LOCAL（本地）/ SCRIPT_SOURCE_GIT（Git 仓库）/
	// SCRIPT_SOURCE_CFS（CFS 文件系统）/ SCRIPT_SOURCE_COS（COS 对象存储）/
	// SCRIPT_SOURCE_WORKSPACE（工作空间）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 前端显示使用，对执行平台无意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayPath *string `json:"DisplayPath,omitnil,omitempty" name:"DisplayPath"`

	// Notebook 相对路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotebookPath *string `json:"NotebookPath,omitnil,omitempty" name:"NotebookPath"`

	// Notebook 绝对路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotebookAbsolutePath *string `json:"NotebookAbsolutePath,omitnil,omitempty" name:"NotebookAbsolutePath"`
}

type TaskTypeProperty struct {
	// 属性名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PropertyKey *string `json:"PropertyKey,omitnil,omitempty" name:"PropertyKey"`

	// 属性值
	// 注意：此字段可能返回 null，表示取不到有效值。
	PropertyValue *string `json:"PropertyValue,omitnil,omitempty" name:"PropertyValue"`
}

// Predefined struct for user
type UnbindWorkflowBundleRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type UnbindWorkflowBundleRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *UnbindWorkflowBundleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindWorkflowBundleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindWorkflowBundleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindWorkflowBundleResponseParams struct {
	// <p>解绑工作流Bundle信息响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *UnbindWorkflowBundleRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindWorkflowBundleResponse struct {
	*tchttp.BaseResponse
	Response *UnbindWorkflowBundleResponseParams `json:"Response"`
}

func (r *UnbindWorkflowBundleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindWorkflowBundleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindWorkflowBundleRsp struct {
	// 操作状态，true 表示成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateFileRequestParams struct {
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID。来源：CreateFile / ListFiles / GetFile 接口返回的 FileId</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件运行配置。不传则不更新配置</p>
	FileConfig *FileConfig `json:"FileConfig,omitnil,omitempty" name:"FileConfig"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>绑定的 BundleId。来源：ListBundles 接口返回的 BundleId</p>
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>绑定的 BundleInfo，JSON 字符串</p>
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>文件内容。不传则不更新内容</p>
	Storage *FileStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>目标文件名，非空且与当前文件名不同时执行 rename 动作。长度不超过 SCRIPT_NAME_MAX_LENGTH，禁止以 . 或 .. 开头/结尾，禁止空格、双点、控制字符及 Linux 保留名（参考 docs/linux_filename_rules.md）。与 ExtensionType 一起校验后缀合法性</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>目标父目录路径，非空时执行 move 动作。根目录传 /；与 FileName 可同时出现，语义为「移动+重命名」。与 CreateFile 的 ParentFolderPath 保持一致</p>
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// <p>目标父目录的 FileType。取值：FOLDER、GIT_FOLDER。仅当 ParentFolderPath 非空时使用；缺省时按解析出的父目录实际类型处理</p>
	TargetFileType *string `json:"TargetFileType,omitnil,omitempty" name:"TargetFileType"`

	// <p>动作类型（必填，未来版本会强制校验）。取值：1 = UPDATE_CONTENT（仅更新 FileConfig / Storage / Bundle*，禁止传 FileName / ParentFolderPath / TargetFileType）；2 = RENAME（仅重命名，必须传 FileName，禁止传 ParentFolderPath / FileConfig / Storage / Bundle*）；3 = MOVE（仅移动，必须传 ParentFolderPath，禁止传 FileName / FileConfig / Storage / Bundle*）。参数互斥校验失败会返回 ParamIllegal 错误</p>
	UpdateAction *int64 `json:"UpdateAction,omitnil,omitempty" name:"UpdateAction"`
}

type UpdateFileRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间 ID。来源：ListWorkspaces 接口返回的 WorkspaceId</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>文件 ID。来源：CreateFile / ListFiles / GetFile 接口返回的 FileId</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件运行配置。不传则不更新配置</p>
	FileConfig *FileConfig `json:"FileConfig,omitnil,omitempty" name:"FileConfig"`

	// <p>文件类型。取值：FILE（普通文件/脚本）、NOTEBOOK_FILE（Notebook）、SQL_FILE（SQL文件）。对应 common/domain/entity.proto EntityType</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>绑定的 BundleId。来源：ListBundles 接口返回的 BundleId</p>
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>绑定的 BundleInfo，JSON 字符串</p>
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>文件内容。不传则不更新内容</p>
	Storage *FileStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// <p>目标文件名，非空且与当前文件名不同时执行 rename 动作。长度不超过 SCRIPT_NAME_MAX_LENGTH，禁止以 . 或 .. 开头/结尾，禁止空格、双点、控制字符及 Linux 保留名（参考 docs/linux_filename_rules.md）。与 ExtensionType 一起校验后缀合法性</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>目标父目录路径，非空时执行 move 动作。根目录传 /；与 FileName 可同时出现，语义为「移动+重命名」。与 CreateFile 的 ParentFolderPath 保持一致</p>
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// <p>目标父目录的 FileType。取值：FOLDER、GIT_FOLDER。仅当 ParentFolderPath 非空时使用；缺省时按解析出的父目录实际类型处理</p>
	TargetFileType *string `json:"TargetFileType,omitnil,omitempty" name:"TargetFileType"`

	// <p>动作类型（必填，未来版本会强制校验）。取值：1 = UPDATE_CONTENT（仅更新 FileConfig / Storage / Bundle*，禁止传 FileName / ParentFolderPath / TargetFileType）；2 = RENAME（仅重命名，必须传 FileName，禁止传 ParentFolderPath / FileConfig / Storage / Bundle*）；3 = MOVE（仅移动，必须传 ParentFolderPath，禁止传 FileName / FileConfig / Storage / Bundle*）。参数互斥校验失败会返回 ParamIllegal 错误</p>
	UpdateAction *int64 `json:"UpdateAction,omitnil,omitempty" name:"UpdateAction"`
}

func (r *UpdateFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "FileId")
	delete(f, "FileConfig")
	delete(f, "FileType")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	delete(f, "Storage")
	delete(f, "FileName")
	delete(f, "ParentFolderPath")
	delete(f, "TargetFileType")
	delete(f, "UpdateAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFileResponseParams struct {
	// <p>返回结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *FileInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateFileResponse struct {
	*tchttp.BaseResponse
	Response *UpdateFileResponseParams `json:"Response"`
}

func (r *UpdateFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkflowRequestParams struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>待更新的工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>需要清空的字段名列表，用于将指定字段重置为空</p>
	FieldToRemoveList []*string `json:"FieldToRemoveList,omitnil,omitempty" name:"FieldToRemoveList"`

	// <p>更新后的工作流配置，仅传入需要变更的部分即可</p>
	NewSetting *Workflow `json:"NewSetting,omitnil,omitempty" name:"NewSetting"`
}

type UpdateWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// <p>工作空间ID，可通过 ListWorkspaces 获取。必填</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>待更新的工作流ID，可通过 ListWorkflows 获取。必填</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>需要清空的字段名列表，用于将指定字段重置为空</p>
	FieldToRemoveList []*string `json:"FieldToRemoveList,omitnil,omitempty" name:"FieldToRemoveList"`

	// <p>更新后的工作流配置，仅传入需要变更的部分即可</p>
	NewSetting *Workflow `json:"NewSetting,omitnil,omitempty" name:"NewSetting"`
}

func (r *UpdateWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkspaceId")
	delete(f, "WorkflowId")
	delete(f, "FieldToRemoveList")
	delete(f, "NewSetting")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkflowResponseParams struct {
	// <p>更新工作流响应内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *UpdateWorkflowRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkflowResponseParams `json:"Response"`
}

func (r *UpdateWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkflowRsp struct {
	// 更新状态，true 表示成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type Workflow struct {
	// <p>工作空间ID，可通过 ListWorkspaces 获取</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流基本信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseInfo *WorkflowBaseInfo `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// <p>工作流调度配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trigger []*WorkflowTriggerConfiguration `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>工作流参数列表 参数名必填且只能包含数字、大小写字母、空格、.$@#!%^&amp;*()-_+=&gt;<!--'，最长128个字符--></p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>标签 标签名必填且只能包含数字、大小写字母、空格、.$@#!%^&amp;*()-_+=&gt;<!--'，最长128个字符--></p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelList []*LabelBrief `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// <p>工作流告警配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alarm *AlarmBrief `json:"Alarm,omitnil,omitempty" name:"Alarm"`

	// <p>监控指标配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorMetric *MonitorMetricBrief `json:"MonitorMetric,omitnil,omitempty" name:"MonitorMetric"`

	// <p>工作流高级设置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceConfig *WorkflowAdvanceConfig `json:"AdvanceConfig,omitnil,omitempty" name:"AdvanceConfig"`

	// <p>工作流任务列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskList []*WorkflowTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// <p>BundleId，可通过 Bundle 相关接口获取</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>Bundle信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>GIT配置ID，对应GetWorkspaceConfig接口中的ConfigKey</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitConfigId *string `json:"GitConfigId,omitnil,omitempty" name:"GitConfigId"`

	// <p>Git分支信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitBranch *string `json:"GitBranch,omitnil,omitempty" name:"GitBranch"`
}

type WorkflowAdvanceConfig struct {
	// 排队模式，ON（默认）, OFF
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueuingMode *string `json:"QueuingMode,omitnil,omitempty" name:"QueuingMode"`

	// 	
	// 默认值为1
	// 
	// QueuingMode为ON时，MaxConcurrentNum 设置才生效；只能输入大于0的整数，输入非法值自动转换为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConcurrentNum *int64 `json:"MaxConcurrentNum,omitnil,omitempty" name:"MaxConcurrentNum"`
}

type WorkflowBaseInfo struct {
	// 工作流名称，长度不超过 1024
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 工作流ID，创建时无需传入，由服务端生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流运行人UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserUin *string `json:"RunUserUin,omitnil,omitempty" name:"RunUserUin"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作流负责人用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserName *string `json:"OwnerUserName,omitnil,omitempty" name:"OwnerUserName"`

	// 创建人UIN。系统生成字段，入参传值不生效（服务端忽略且不报错）
	// 【已废弃】服务端忽略传入值，不报错。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`
}

type WorkflowBaseInfoDetail struct {
	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 创建人UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 工作流运行人UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserUin *string `json:"RunUserUin,omitnil,omitempty" name:"RunUserUin"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作流负责人用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserName *string `json:"OwnerUserName,omitnil,omitempty" name:"OwnerUserName"`

	// 工作流负责人UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserUin *string `json:"OwnerUserUin,omitnil,omitempty" name:"OwnerUserUin"`

	// 工作流负责人展示名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerDisplayName *string `json:"OwnerDisplayName,omitnil,omitempty" name:"OwnerDisplayName"`

	// 创建时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type WorkflowBrief struct {
	// <p>工作流名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// <p>工作流ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>创建人UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// <p>工作流负责人用户名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserName *string `json:"OwnerUserName,omitnil,omitempty" name:"OwnerUserName"`

	// <p>工作流负责人UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserUin *string `json:"OwnerUserUin,omitnil,omitempty" name:"OwnerUserUin"`

	// <p>工作流负责人展示名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerDisplayName *string `json:"OwnerDisplayName,omitnil,omitempty" name:"OwnerDisplayName"`

	// <p>创建时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>标签列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelList []*LabelBrief `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// <p>工作流调度配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trigger []*WorkflowTriggerConfiguration `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>工作流运行人UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserUin *string `json:"RunUserUin,omitnil,omitempty" name:"RunUserUin"`

	// <p>工作流运行人用户名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserName *string `json:"RunUserName,omitnil,omitempty" name:"RunUserName"`

	// <p>工作流任务节点列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskList []*WorkflowTaskNodeBrief `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// <p>工作流运行情况列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunList []*WorkflowRunBrief `json:"WorkflowRunList,omitnil,omitempty" name:"WorkflowRunList"`

	// <p>资源组信息列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupInfoList []*ResourceGroupInfo `json:"ResourceGroupInfoList,omitnil,omitempty" name:"ResourceGroupInfoList"`

	// <p>工作流权限信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`

	// <p>工作流绑定的 Bundle 唯一标识，未绑定时为空</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// <p>Bundle信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// <p>Git配置ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitConfigId *string `json:"GitConfigId,omitnil,omitempty" name:"GitConfigId"`

	// <p>Git分支信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GitBranch *string `json:"GitBranch,omitnil,omitempty" name:"GitBranch"`
}

type WorkflowRun struct {
	// 主账号ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流运行ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// 工作空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 触发方式，Scheduler、ManualTrigger、Event (参考SchedulerTriggerType)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 运行开始时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStartTime *string `json:"RunStartTime,omitnil,omitempty" name:"RunStartTime"`

	// pending 状态开始时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	PendingStartTime *string `json:"PendingStartTime,omitnil,omitempty" name:"PendingStartTime"`

	// queue 状态开始时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueStartTime *string `json:"QueueStartTime,omitnil,omitempty" name:"QueueStartTime"`

	// 运行结束时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunEndTime *string `json:"RunEndTime,omitnil,omitempty" name:"RunEndTime"`

	// 终态时间，运行进入终态时都有值，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 运行时长，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunCostTime *string `json:"RunCostTime,omitnil,omitempty" name:"RunCostTime"`

	// 并发排队花费时间，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueCostTime *string `json:"QueueCostTime,omitnil,omitempty" name:"QueueCostTime"`

	// 等待资源花费时间，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	PendingCostTime *string `json:"PendingCostTime,omitnil,omitempty" name:"PendingCostTime"`

	// 运行状态。取值参考工作流运行状态枚举，如 Pending / Running / Succeeded / Failed / Killed
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunState *string `json:"RunState,omitnil,omitempty" name:"RunState"`

	// 计算资源（任务的资源组ID集合）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitnil,omitempty" name:"ResourceGroupIds"`

	// 运行用户UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserUin *string `json:"RunUserUin,omitnil,omitempty" name:"RunUserUin"`

	// 运行用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserName *string `json:"RunUserName,omitnil,omitempty" name:"RunUserName"`

	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeString *string `json:"ErrorCodeString,omitnil,omitempty" name:"ErrorCodeString"`

	// 运行参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowParams *string `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// 工作流版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowVersionId *string `json:"WorkflowVersionId,omitnil,omitempty" name:"WorkflowVersionId"`

	// 当前工作流是否支持重跑
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportRerun *bool `json:"SupportRerun,omitnil,omitempty" name:"SupportRerun"`

	// 工作流运行创建时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 重跑次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RerunTimes *int64 `json:"RerunTimes,omitnil,omitempty" name:"RerunTimes"`

	// 运行的任务范围，任务ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelectedTaskIds []*string `json:"SelectedTaskIds,omitnil,omitempty" name:"SelectedTaskIds"`

	// 资源组信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupInfoList []*ResourceGroupInfo `json:"ResourceGroupInfoList,omitnil,omitempty" name:"ResourceGroupInfoList"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelList []*LabelBrief `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// 父工作流运行ID 【由嵌套工作流触发独有】
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentWorkflowRunId *string `json:"ParentWorkflowRunId,omitnil,omitempty" name:"ParentWorkflowRunId"`

	// 父工作流任务运行ID 【由嵌套工作流触发独有】
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentWorkflowTaskRunId *string `json:"ParentWorkflowTaskRunId,omitnil,omitempty" name:"ParentWorkflowTaskRunId"`

	// 父工作流任务运行名称 【由嵌套工作流触发独有】
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentWorkflowTaskRunName *string `json:"ParentWorkflowTaskRunName,omitnil,omitempty" name:"ParentWorkflowTaskRunName"`

	// 权限信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 工作流高级运行时用户填入的参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedParameters []*AdvancedParameter `json:"AdvancedParameters,omitnil,omitempty" name:"AdvancedParameters"`
}

type WorkflowRunBrief struct {
	// 工作流运行ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// 运行开始时间，单位：毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStartTime *string `json:"RunStartTime,omitnil,omitempty" name:"RunStartTime"`

	// 运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunState *string `json:"RunState,omitnil,omitempty" name:"RunState"`

	// 运行错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeString *string `json:"ErrorCodeString,omitnil,omitempty" name:"ErrorCodeString"`
}

type WorkflowTask struct {
	// 任务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// 任务依赖
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnList []*DependOnBrief `json:"DependOnList,omitnil,omitempty" name:"DependOnList"`

	// 任务ID，创建时无需传入，由服务端生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *TaskType `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 资源组ID，可通过资源组相关接口获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 任务告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alarm *AlarmBrief `json:"Alarm,omitnil,omitempty" name:"Alarm"`

	// 监控指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorMetric *MonitorMetricBrief `json:"MonitorMetric,omitnil,omitempty" name:"MonitorMetric"`

	// 任务重试策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRetryStrategy *TaskRetryStrategy `json:"TaskRetryStrategy,omitnil,omitempty" name:"TaskRetryStrategy"`

	// <p>任依赖运行条件</p><ul><li>ALL_SUCCESS: 全部成功：所有上游依赖任务均已执行并成功</li><li>ONE_SUCCESS: 至少一个成功：至少有一个上游依赖任务成功</li><li>NONE_FAILED: 目前没有失败：没有依赖任务失败，并且至少有一个依赖任务在运行中</li><li>ALL_DONE: 全部完成：所有上游依赖任务均已执行并完成（无论成功或失败</li><li>ONE_FAILED: 至少一个失败：至少有一个上游依赖任务失败</li><li>ALL_FAILED: 全部失败：所有上游依赖任务都失败</li><li>ALL_DONE_AT_LEAST_ONE_SUCCESS：上游全部完成至少一个成功: 所有上游依赖任务都达到终态时，进行依赖判断，至少有一个成功，则依赖判断成功，否则就是跳过运行</li><li>ALL_SKIPPED：上游全部完成，没有跳过运行: 所有上游依赖任务都达到终态时，进行依赖判断, 如果上游状态全部都是成功、失败、上游失败状态，则依赖判断成功，否则为跳过运行</li><li>ONE_DONE：至少一个完成：上游只要有一个完成了，就进行依赖判断，且依赖判断成功，否则还是等待上游</li><li>ALL_DONE_NONE_FAILED_AT_LEAST_ONE_SUCCESS：上游全部完成，没有失败，至少有一个成功: 所有上游依赖任务都达到终态时，进行依赖判断，上游没有一个失败且至少有一个成功的情况下，依赖判断成功，否则就是跳过运行</li><li>NONE_SKIPPED：上游全部完成，没有跳过运行: 所有上游依赖任务都达到终态时，进行依赖判断, 如果上游状态全部都是成功、失败、上游失败状态，则依赖判断成功，否则为跳过运行</li><li>ALL_DONE_AT_LEAST_ONE_FAILED：上游全部完成至少一个失败: 所有上游依赖任务都达到终态时，进行依赖判断，至少有一个失败，则依赖判断成功，否则就是跳过运行</li><li>ADVANCED:运行条件为高级模式时配置</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnRunCondition *string `json:"DependOnRunCondition,omitnil,omitempty" name:"DependOnRunCondition"`

	// 任务X坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftCoordinate *float64 `json:"LeftCoordinate,omitnil,omitempty" name:"LeftCoordinate"`

	// 任务Y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopCoordinate *float64 `json:"TopCoordinate,omitnil,omitempty" name:"TopCoordinate"`

	// <p>任务高级运行参数，当DependOnRunCondition为ADVANCED时配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedDependencyConfig *AdvancedDependencyConfig `json:"AdvancedDependencyConfig,omitnil,omitempty" name:"AdvancedDependencyConfig"`

	// <p>内嵌任务（FOR_EACH任务的子任务）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerTask *WorkflowTask `json:"InnerTask,omitnil,omitempty" name:"InnerTask"`

	// 创建时间，单位：毫秒时间戳。出参专用，系统生成，入参传值不生效
	// 【已废弃】服务端忽略传入值，不报错。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，单位：毫秒时间戳。出参专用，系统生成，入参传值不生效
	// 【已废弃】服务端忽略传入值，不报错。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建人UIN。出参专用，系统生成，入参传值不生效
	// 【已废弃】服务端忽略传入值，不报错。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`
}

type WorkflowTaskNodeBrief struct {
	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeName *string `json:"TaskTypeName,omitnil,omitempty" name:"TaskTypeName"`

	// 任务依赖列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnList []*DependOnBrief `json:"DependOnList,omitnil,omitempty" name:"DependOnList"`

	// 任务资源组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 任务资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 任务X坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftCoordinate *float64 `json:"LeftCoordinate,omitnil,omitempty" name:"LeftCoordinate"`

	// 任务Y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopCoordinate *float64 `json:"TopCoordinate,omitnil,omitempty" name:"TopCoordinate"`

	// 任务重试策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRetryStrategy *TaskRetryStrategy `json:"TaskRetryStrategy,omitnil,omitempty" name:"TaskRetryStrategy"`

	// 依赖运行条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnRunCondition *string `json:"DependOnRunCondition,omitnil,omitempty" name:"DependOnRunCondition"`

	// 高级依赖配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedDependencyConfig *AdvancedDependencyConfig `json:"AdvancedDependencyConfig,omitnil,omitempty" name:"AdvancedDependencyConfig"`

	// 内嵌工作流任务节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerTask *WorkflowTaskNodeBrief `json:"InnerTask,omitnil,omitempty" name:"InnerTask"`
}

type WorkflowTaskRun struct {
	// <p>任务名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>任务运行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowTaskRunId *string `json:"WorkflowTaskRunId,omitnil,omitempty" name:"WorkflowTaskRunId"`

	// <p>运行状态。取值参考工作流任务运行状态枚举，如 Pending / Running / Succeeded / Failed / Killed</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunState *string `json:"RunState,omitnil,omitempty" name:"RunState"`

	// <p>工作空间ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>工作流ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流运行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>任务ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务类型名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeName *string `json:"TaskTypeName,omitnil,omitempty" name:"TaskTypeName"`

	// <p>任务版本ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskVersionId *string `json:"TaskVersionId,omitnil,omitempty" name:"TaskVersionId"`

	// <p>触发类型 (参考SchedulerTriggerType枚举)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>所属资源组ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// <p>错误码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeString *string `json:"ErrorCodeString,omitnil,omitempty" name:"ErrorCodeString"`

	// <p>运行用户UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserUin *string `json:"RunUserUin,omitnil,omitempty" name:"RunUserUin"`

	// <p>运行用户名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunUserName *string `json:"RunUserName,omitnil,omitempty" name:"RunUserName"`

	// <p>创建人UIN</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// <p>执行平台执行ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// <p>创建时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>依赖任务完成时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependenceFinishedTime *string `json:"DependenceFinishedTime,omitnil,omitempty" name:"DependenceFinishedTime"`

	// <p>运行开始时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunStartTime *string `json:"RunStartTime,omitnil,omitempty" name:"RunStartTime"`

	// <p>运行结束时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunEndTime *string `json:"RunEndTime,omitnil,omitempty" name:"RunEndTime"`

	// <p>运行时长，单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunCostTime *string `json:"RunCostTime,omitnil,omitempty" name:"RunCostTime"`

	// <p>等待时长（依赖就绪到开始运行的等待耗时），单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaitTime *string `json:"WaitTime,omitnil,omitempty" name:"WaitTime"`

	// <p>下发执行平台时间，单位：毫秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssueTime *string `json:"IssueTime,omitnil,omitempty" name:"IssueTime"`

	// <p>时区</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// <p>依赖上游任务ID列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnList []*string `json:"DependOnList,omitnil,omitempty" name:"DependOnList"`

	// <p>运行参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunParams *string `json:"RunParams,omitnil,omitempty" name:"RunParams"`

	// <p>任务扩展信息，包含脚本路径</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeExtensions *string `json:"TaskTypeExtensions,omitnil,omitempty" name:"TaskTypeExtensions"`

	// <p>任务X坐标</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftCoordinate *float64 `json:"LeftCoordinate,omitnil,omitempty" name:"LeftCoordinate"`

	// <p>任务Y坐标</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopCoordinate *float64 `json:"TopCoordinate,omitnil,omitempty" name:"TopCoordinate"`

	// <p>重试次数，为 0 则表示首次运行</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryTimes *int64 `json:"RetryTimes,omitnil,omitempty" name:"RetryTimes"`

	// <p>工作流名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// <p>重跑次数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RerunTimes *int64 `json:"RerunTimes,omitnil,omitempty" name:"RerunTimes"`

	// <p>是否最新一次运行</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLatestRun *bool `json:"IsLatestRun,omitnil,omitempty" name:"IsLatestRun"`

	// <p>资源组信息列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupInfoList []*ResourceGroupInfo `json:"ResourceGroupInfoList,omitnil,omitempty" name:"ResourceGroupInfoList"`

	// <p>运行结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunResult *string `json:"RunResult,omitnil,omitempty" name:"RunResult"`

	// <p>依赖运行条件</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependOnRunCondition *string `json:"DependOnRunCondition,omitnil,omitempty" name:"DependOnRunCondition"`

	// <p>高级依赖配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedDependencyConfig *AdvancedDependencyConfig `json:"AdvancedDependencyConfig,omitnil,omitempty" name:"AdvancedDependencyConfig"`

	// <p>内嵌工作流任务信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerTask *InnerWorkflowTaskBrief `json:"InnerTask,omitnil,omitempty" name:"InnerTask"`
}

type WorkflowTriggerAdvancedConfiguration struct {
	// 任务重试模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRetryMode *string `json:"TaskRetryMode,omitnil,omitempty" name:"TaskRetryMode"`
}

type WorkflowTriggerConfiguration struct {
	// 调度配置ID，创建时无需传入，由服务端生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// 调度状态 启动：START，暂停：PAUSE
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerStatus *string `json:"SchedulerStatus,omitnil,omitempty" name:"SchedulerStatus"`

	// 触发方式，
	// - 定时触发：TIME_TRIGGER
	// - 持续运行：CONTINUE_RUN
	// 
	// 注意：
	// - TIME_TRIGGER 模式下，SchedulerStatus、SchedulerTimeZone、StartTime、EndTime、ConfigMode、CycleType、CrontabExpression 必填；
	// - CONTINUE_RUN 模式下，AdvancedConfig必填；
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerMode *string `json:"TriggerMode,omitnil,omitempty" name:"TriggerMode"`

	// 调度时区
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerTimeZone *string `json:"SchedulerTimeZone,omitnil,omitempty" name:"SchedulerTimeZone"`

	// 调度生效时间，单位：毫秒时间戳。必须小于 EndTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 调度结束时间，单位：毫秒时间戳。必须大于 StartTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 配置方式，常规：COMMON，CRON表达式：CRON_EXPRESSION
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigMode *string `json:"ConfigMode,omitnil,omitempty" name:"ConfigMode"`

	// 周期类型：支持的类型为 ONEOFF_CYCLE: 一次性 YEAR_CYCLE: 年 MONTH_CYCLE: 月 WEEK_CYCLE: 周 DAY_CYCLE: 天
	// HOUR_CYCLE: 小时 MINUTE_CYCLE: 分钟 CRONTAB_CYCLE: crontab表达式类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// cron表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// Json格式，对账使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 高级配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedConfig *WorkflowTriggerAdvancedConfiguration `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}