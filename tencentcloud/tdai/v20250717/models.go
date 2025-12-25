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

package v20250717

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Agent struct {
	// 智能体Id
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 智能体名称
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// 智能体类型
	AgentInternalName *string `json:"AgentInternalName,omitnil,omitempty" name:"AgentInternalName"`

	// 智能体状态
	AgentStatus *string `json:"AgentStatus,omitnil,omitempty" name:"AgentStatus"`

	// 默认版本
	DefaultVersion *string `json:"DefaultVersion,omitnil,omitempty" name:"DefaultVersion"`

	// 智能体模式
	AgentType *string `json:"AgentType,omitnil,omitempty" name:"AgentType"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建者
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新者
	Updater *string `json:"Updater,omitnil,omitempty" name:"Updater"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AgentDutyTask struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务开始运行时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 任务状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 对外展示的Extra信息
	ResultExtraKey []*string `json:"ResultExtraKey,omitnil,omitempty" name:"ResultExtraKey"`

	// 业务的额外敏感信息
	Extra []*ExtraInfo `json:"Extra,omitnil,omitempty" name:"Extra"`
}

type AgentInstance struct {
	// 智能体实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 智能体实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 智能体ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 智能体名称
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// 智能体类型
	AgentInternalName *string `json:"AgentInternalName,omitnil,omitempty" name:"AgentInternalName"`

	// 智能体服务模式
	AgentType *string `json:"AgentType,omitnil,omitempty" name:"AgentType"`

	// 智能体版本
	AgentVersion *string `json:"AgentVersion,omitnil,omitempty" name:"AgentVersion"`

	// 智能体实例状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 智能体实例参数列表
	Parameters []*Parameter `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 资源绑定Tag列表
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ChatBrief struct {
	// 主账号Id
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子账号uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 智能体实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话ID
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 会话标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 会话状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 最后一条流式ID
	LastStreamingId *string `json:"LastStreamingId,omitnil,omitempty" name:"LastStreamingId"`

	// 最后一条流式TokenID
	LastStreamingTokenId *int64 `json:"LastStreamingTokenId,omitnil,omitempty" name:"LastStreamingTokenId"`
}

type ChatDetail struct {
	// 角色
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 用户消息
	UserMessage *string `json:"UserMessage,omitnil,omitempty" name:"UserMessage"`

	// 助手消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssistantMessage []*CreateChatCompletionRes `json:"AssistantMessage,omitnil,omitempty" name:"AssistantMessage"`
}

type CodeRepo struct {
	// 代码仓库地址
	RepoUrl *string `json:"RepoUrl,omitnil,omitempty" name:"RepoUrl"`

	// 分支名
	Branch *string `json:"Branch,omitnil,omitempty" name:"Branch"`

	// Commit信息
	GitCommitPipelines []*string `json:"GitCommitPipelines,omitnil,omitempty" name:"GitCommitPipelines"`

	// 数据库ORM类型
	GitORMType *string `json:"GitORMType,omitnil,omitempty" name:"GitORMType"`

	// 代码编写语言
	GitCodeLanguage *string `json:"GitCodeLanguage,omitnil,omitempty" name:"GitCodeLanguage"`
}

// Predefined struct for user
type ContinueAgentWorkRequestParams struct {
	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent任务类型
	AgentTaskType *string `json:"AgentTaskType,omitnil,omitempty" name:"AgentTaskType"`
}

type ContinueAgentWorkRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent任务类型
	AgentTaskType *string `json:"AgentTaskType,omitnil,omitempty" name:"AgentTaskType"`
}

func (r *ContinueAgentWorkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContinueAgentWorkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentTaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ContinueAgentWorkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ContinueAgentWorkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ContinueAgentWorkResponse struct {
	*tchttp.BaseResponse
	Response *ContinueAgentWorkResponseParams `json:"Response"`
}

func (r *ContinueAgentWorkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContinueAgentWorkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentInstanceRequestParams struct {
	// 智能体ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 智能体版本
	AgentVersion *string `json:"AgentVersion,omitnil,omitempty" name:"AgentVersion"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 智能体实例的参数列表
	Parameters []*Parameter `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// 资源的标签信息
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateAgentInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 智能体ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 智能体版本
	AgentVersion *string `json:"AgentVersion,omitnil,omitempty" name:"AgentVersion"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 智能体实例的参数列表
	Parameters []*Parameter `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// 资源的标签信息
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateAgentInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "AgentVersion")
	delete(f, "InstanceName")
	delete(f, "Parameters")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentInstanceResponseParams struct {
	// 智能体实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 智能体实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgentInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentInstanceResponseParams `json:"Response"`
}

func (r *CreateAgentInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChatCompletionRequestParams struct {
	// 是否隐藏
	IsHidden *bool `json:"IsHidden,omitnil,omitempty" name:"IsHidden"`

	// 是否隐藏会话
	IsChatHidden *bool `json:"IsChatHidden,omitnil,omitempty" name:"IsChatHidden"`
}

type CreateChatCompletionRequest struct {
	*tchttp.BaseRequest
	
	// 是否隐藏
	IsHidden *bool `json:"IsHidden,omitnil,omitempty" name:"IsHidden"`

	// 是否隐藏会话
	IsChatHidden *bool `json:"IsChatHidden,omitnil,omitempty" name:"IsChatHidden"`
}

func (r *CreateChatCompletionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChatCompletionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsHidden")
	delete(f, "IsChatHidden")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChatCompletionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateChatCompletionRes struct {
	// 枚举值，返回的数据类型
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 消息时间戳
	Created *int64 `json:"Created,omitnil,omitempty" name:"Created"`

	// 输出模型
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 用户标识
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账户标识
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 当前账户标识
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Request ID
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// 当前会话ID
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 当前流ID
	StreamingId *string `json:"StreamingId,omitnil,omitempty" name:"StreamingId"`

	// 当前任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 消息的数据详情
	Choices []*UploadChoice `json:"Choices,omitnil,omitempty" name:"Choices"`
}

// Predefined struct for user
type CreateChatCompletionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateChatCompletionResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *CreateChatCompletionResponseParams `json:"Response"`
}

func (r *CreateChatCompletionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChatCompletionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDutyTaskDetailRequestParams struct {
	// 业务参数列表
	Parameters []*Parameter `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

type DescribeAgentDutyTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 业务参数列表
	Parameters []*Parameter `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

func (r *DescribeAgentDutyTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDutyTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Parameters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentDutyTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDutyTaskDetailResponseParams struct {
	// 任务详细信息
	AgentDutyTask *AgentDutyTask `json:"AgentDutyTask,omitnil,omitempty" name:"AgentDutyTask"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentDutyTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentDutyTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeAgentDutyTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDutyTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDutyTasksRequestParams struct {
	// agent实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话ID
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 查询开始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列表查询数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 任务启动时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务类型
	AgentTaskType *string `json:"AgentTaskType,omitnil,omitempty" name:"AgentTaskType"`

	// 业务参数
	Parameters []*Parameter `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

type DescribeAgentDutyTasksRequest struct {
	*tchttp.BaseRequest
	
	// agent实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话ID
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 查询开始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列表查询数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 任务启动时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务类型
	AgentTaskType *string `json:"AgentTaskType,omitnil,omitempty" name:"AgentTaskType"`

	// 业务参数
	Parameters []*Parameter `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

func (r *DescribeAgentDutyTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDutyTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ChatId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AgentTaskType")
	delete(f, "Parameters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentDutyTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDutyTasksResponseParams struct {
	// 查询结果总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务详细信息
	DutyTasks []*AgentDutyTask `json:"DutyTasks,omitnil,omitempty" name:"DutyTasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentDutyTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentDutyTasksResponseParams `json:"Response"`
}

func (r *DescribeAgentDutyTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDutyTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentInstanceRequestParams struct {
	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAgentInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeAgentInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentInstanceResponseParams struct {
	// 智能体实例详情
	AgentInstance *AgentInstance `json:"AgentInstance,omitnil,omitempty" name:"AgentInstance"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentInstanceResponseParams `json:"Response"`
}

func (r *DescribeAgentInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentInstancesRequestParams struct {
	// 查询开始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列表查询数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名，为空时查询所有，如果填写则会根据InstanceName筛选
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 智能体ID，为空时查询所有，如果填写则会根据AgentId筛选
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 智能体名称，为空时查询所有，如果填写则会根据AgentName筛选
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// 智能体类型，为空时查询所有，如果填写则会根据AgentName筛选
	AgentInternalName *string `json:"AgentInternalName,omitnil,omitempty" name:"AgentInternalName"`

	// 智能体实例状态，为空时查询所有，如果填写则会根据Status筛选
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 标签过滤信息
	TagFilter []*TagFilter `json:"TagFilter,omitnil,omitempty" name:"TagFilter"`
}

type DescribeAgentInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列表查询数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名，为空时查询所有，如果填写则会根据InstanceName筛选
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 智能体ID，为空时查询所有，如果填写则会根据AgentId筛选
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 智能体名称，为空时查询所有，如果填写则会根据AgentName筛选
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// 智能体类型，为空时查询所有，如果填写则会根据AgentName筛选
	AgentInternalName *string `json:"AgentInternalName,omitnil,omitempty" name:"AgentInternalName"`

	// 智能体实例状态，为空时查询所有，如果填写则会根据Status筛选
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 标签过滤信息
	TagFilter []*TagFilter `json:"TagFilter,omitnil,omitempty" name:"TagFilter"`
}

func (r *DescribeAgentInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "AgentId")
	delete(f, "AgentName")
	delete(f, "AgentInternalName")
	delete(f, "Status")
	delete(f, "TagFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentInstancesResponseParams struct {
	// 查询结果总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 智能体实例列表
	Items []*AgentInstance `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentInstancesResponseParams `json:"Response"`
}

func (r *DescribeAgentInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentsRequestParams struct {
	// 查询开始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列表查询数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 智能体ID，为空时查询所有，如果填写则会根据AgentId筛选
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 智能体名称，为空时查询所有，如果填写则会根据AgentName筛选
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// 智能体类型，为空时查询所有，如果填写则会根据AgentName筛选
	AgentInternalName *string `json:"AgentInternalName,omitnil,omitempty" name:"AgentInternalName"`

	// 智能体状态，为空时查询所有，如果填写则会根据AgentStatus筛选
	AgentStatus *string `json:"AgentStatus,omitnil,omitempty" name:"AgentStatus"`
}

type DescribeAgentsRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列表查询数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 智能体ID，为空时查询所有，如果填写则会根据AgentId筛选
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 智能体名称，为空时查询所有，如果填写则会根据AgentName筛选
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// 智能体类型，为空时查询所有，如果填写则会根据AgentName筛选
	AgentInternalName *string `json:"AgentInternalName,omitnil,omitempty" name:"AgentInternalName"`

	// 智能体状态，为空时查询所有，如果填写则会根据AgentStatus筛选
	AgentStatus *string `json:"AgentStatus,omitnil,omitempty" name:"AgentStatus"`
}

func (r *DescribeAgentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AgentId")
	delete(f, "AgentName")
	delete(f, "AgentInternalName")
	delete(f, "AgentStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentsResponseParams struct {
	// 查询结果总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 智能体列表
	Items []*Agent `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentsResponseParams `json:"Response"`
}

func (r *DescribeAgentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChatDetailRequestParams struct {
	// 智能体ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话Id
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 流ID
	StreamingId *string `json:"StreamingId,omitnil,omitempty" name:"StreamingId"`

	// 开始拉取的流式TokenID。0表示从该流最早的TokenID开始获取
	BeginStreamingTokenId *int64 `json:"BeginStreamingTokenId,omitnil,omitempty" name:"BeginStreamingTokenId"`

	// 单次获取的token数量，默认2000
	TokenLimit *int64 `json:"TokenLimit,omitnil,omitempty" name:"TokenLimit"`
}

type DescribeChatDetailRequest struct {
	*tchttp.BaseRequest
	
	// 智能体ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话Id
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 流ID
	StreamingId *string `json:"StreamingId,omitnil,omitempty" name:"StreamingId"`

	// 开始拉取的流式TokenID。0表示从该流最早的TokenID开始获取
	BeginStreamingTokenId *int64 `json:"BeginStreamingTokenId,omitnil,omitempty" name:"BeginStreamingTokenId"`

	// 单次获取的token数量，默认2000
	TokenLimit *int64 `json:"TokenLimit,omitnil,omitempty" name:"TokenLimit"`
}

func (r *DescribeChatDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChatDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ChatId")
	delete(f, "StreamingId")
	delete(f, "BeginStreamingTokenId")
	delete(f, "TokenLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChatDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChatDetailResponseParams struct {
	// 主账号ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账号Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子账号Uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 智能体实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话ID
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 最后一条流式TokenID
	LastStreamingTokenId *int64 `json:"LastStreamingTokenId,omitnil,omitempty" name:"LastStreamingTokenId"`

	// Streaming数量
	StreamingCount *int64 `json:"StreamingCount,omitnil,omitempty" name:"StreamingCount"`

	// 对话流列表
	Streamings []*ChatDetail `json:"Streamings,omitnil,omitempty" name:"Streamings"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChatDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChatDetailResponseParams `json:"Response"`
}

func (r *DescribeChatDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChatDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChatsRequestParams struct {
	// 智能体ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询开始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列表查询数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeChatsRequest struct {
	*tchttp.BaseRequest
	
	// 智能体ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询开始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列表查询数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeChatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChatsResponseParams struct {
	// 查询结果总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 对话流列表
	Chats []*ChatBrief `json:"Chats,omitnil,omitempty" name:"Chats"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChatsResponseParams `json:"Response"`
}

func (r *DescribeChatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportUrlRequestParams struct {

}

type DescribeReportUrlRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeReportUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReportUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportUrlResponseParams struct {
	// 下载地址
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReportUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReportUrlResponseParams `json:"Response"`
}

func (r *DescribeReportUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtraInfo struct {
	// 出参额外信息的Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 额外信息描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// ExtraInfo的值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 值的数据结构类型
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`
}

type InstanceInfos struct {
	// 数据库地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 数据库实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

// Predefined struct for user
type IsolateAgentInstanceRequestParams struct {
	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type IsolateAgentInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *IsolateAgentInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateAgentInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateAgentInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateAgentInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateAgentInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateAgentInstanceResponseParams `json:"Response"`
}

func (r *IsolateAgentInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateAgentInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentInstanceParametersRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务类型, 可选，默认default
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 更新的参数列表
	Parameters *Parameter `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// 风险SQL智能体参数列表
	SqlAgentParameter *SqlAgentParameter `json:"SqlAgentParameter,omitnil,omitempty" name:"SqlAgentParameter"`
}

type ModifyAgentInstanceParametersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务类型, 可选，默认default
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 更新的参数列表
	Parameters *Parameter `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// 风险SQL智能体参数列表
	SqlAgentParameter *SqlAgentParameter `json:"SqlAgentParameter,omitnil,omitempty" name:"SqlAgentParameter"`
}

func (r *ModifyAgentInstanceParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentInstanceParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TaskType")
	delete(f, "Parameters")
	delete(f, "SqlAgentParameter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAgentInstanceParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentInstanceParametersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAgentInstanceParametersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAgentInstanceParametersResponseParams `json:"Response"`
}

func (r *ModifyAgentInstanceParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentInstanceParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyChatTitleRequestParams struct {
	// 智能体ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话Id
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

type ModifyChatTitleRequest struct {
	*tchttp.BaseRequest
	
	// 智能体ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话Id
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

func (r *ModifyChatTitleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyChatTitleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ChatId")
	delete(f, "Title")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyChatTitleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyChatTitleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyChatTitleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyChatTitleResponseParams `json:"Response"`
}

func (r *ModifyChatTitleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyChatTitleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Parameter struct {
	// 参数键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 枚举值，可取值包括：string(字符串), int(整型), float(浮点型), bool(布尔型), struct(结构体), array(数组), 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`
}

// Predefined struct for user
type PauseAgentWorkRequestParams struct {
	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent任务类型
	AgentTaskType *string `json:"AgentTaskType,omitnil,omitempty" name:"AgentTaskType"`
}

type PauseAgentWorkRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Agent任务类型
	AgentTaskType *string `json:"AgentTaskType,omitnil,omitempty" name:"AgentTaskType"`
}

func (r *PauseAgentWorkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseAgentWorkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentTaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseAgentWorkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseAgentWorkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseAgentWorkResponse struct {
	*tchttp.BaseResponse
	Response *PauseAgentWorkResponseParams `json:"Response"`
}

func (r *PauseAgentWorkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseAgentWorkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverAgentInstanceRequestParams struct {
	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type RecoverAgentInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *RecoverAgentInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverAgentInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverAgentInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverAgentInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecoverAgentInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RecoverAgentInstanceResponseParams `json:"Response"`
}

func (r *RecoverAgentInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverAgentInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveChatRequestParams struct {
	// 智能体ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话Id
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`
}

type RemoveChatRequest struct {
	*tchttp.BaseRequest
	
	// 智能体ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话Id
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`
}

func (r *RemoveChatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveChatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ChatId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveChatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveChatResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveChatResponse struct {
	*tchttp.BaseResponse
	Response *RemoveChatResponseParams `json:"Response"`
}

func (r *RemoveChatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveChatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SqlAgentParameter struct {
	// 数据库实例信息列表
	InstanceInfos []*InstanceInfos `json:"InstanceInfos,omitnil,omitempty" name:"InstanceInfos"`

	// 代码仓库信息
	CodeRepo *CodeRepo `json:"CodeRepo,omitnil,omitempty" name:"CodeRepo"`
}

// Predefined struct for user
type StartAgentTaskRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置Token
	InstanceToken *string `json:"InstanceToken,omitnil,omitempty" name:"InstanceToken"`
}

type StartAgentTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置Token
	InstanceToken *string `json:"InstanceToken,omitnil,omitempty" name:"InstanceToken"`
}

func (r *StartAgentTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAgentTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartAgentTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartAgentTaskResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartAgentTaskResponse struct {
	*tchttp.BaseResponse
	Response *StartAgentTaskResponseParams `json:"Response"`
}

func (r *StartAgentTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAgentTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagFilter struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagItem struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type TerminateAgentInstanceRequestParams struct {
	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type TerminateAgentInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，为空时查询所有，如果填写则会根据InstanceId筛选
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *TerminateAgentInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateAgentInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateAgentInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateAgentInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateAgentInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateAgentInstanceResponseParams `json:"Response"`
}

func (r *TerminateAgentInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateAgentInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadChoice struct {
	// 消息索引
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 当前消息步骤
	StepNo *int64 `json:"StepNo,omitnil,omitempty" name:"StepNo"`

	// 当前步骤
	CurrentStep *string `json:"CurrentStep,omitnil,omitempty" name:"CurrentStep"`

	// 增量信息
	Delta *UploadDelta `json:"Delta,omitnil,omitempty" name:"Delta"`

	// 结束原因
	FinishReason *string `json:"FinishReason,omitnil,omitempty" name:"FinishReason"`

	// 错误信息，FinishReason为error时有效
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

type UploadDelta struct {
	// 步骤摘要
	StepBrief *string `json:"StepBrief,omitnil,omitempty" name:"StepBrief"`

	// 步骤详情
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}