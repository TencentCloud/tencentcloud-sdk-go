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

package v20201028

import (
    "encoding/json"
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AutomationAgentInfo struct {

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent 版本号。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 上次心跳时间
	LastHeartbeatTime *string `json:"LastHeartbeatTime,omitempty" name:"LastHeartbeatTime"`

	// Agent状态，取值范围：
	// <li> Online：在线
	// <li> Offline：离线
	AgentStatus *string `json:"AgentStatus,omitempty" name:"AgentStatus"`

	// Agent运行环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type Command struct {

	// 命令ID。
	CommandId *string `json:"CommandId,omitempty" name:"CommandId"`

	// 命令名称。
	CommandName *string `json:"CommandName,omitempty" name:"CommandName"`

	// 命令描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// Base64编码后的命令内容。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 命令类型。
	CommandType *string `json:"CommandType,omitempty" name:"CommandType"`

	// 命令执行路径。
	WorkingDirectory *string `json:"WorkingDirectory,omitempty" name:"WorkingDirectory"`

	// 命令超时时间。
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 命令创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 命令更新时间。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 是否启用自定义参数功能。
	EnableParameter *bool `json:"EnableParameter,omitempty" name:"EnableParameter"`

	// 自定义参数的默认取值。
	DefaultParameters *string `json:"DefaultParameters,omitempty" name:"DefaultParameters"`

	// 命令的结构化描述。公共命令有值，用户命令为空字符串。
	FormattedDescription *string `json:"FormattedDescription,omitempty" name:"FormattedDescription"`

	// 命令创建者。TAT 代表公共命令，USER 代表个人命令。
	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`

	// 命令关联的标签列表。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CommandDocument struct {

	// Base64 编码后的执行命令。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 命令类型。
	CommandType *string `json:"CommandType,omitempty" name:"CommandType"`

	// 超时时间。
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 执行路径。
	WorkingDirectory *string `json:"WorkingDirectory,omitempty" name:"WorkingDirectory"`
}

type CreateCommandRequest struct {
	*tchttp.BaseRequest

	// 命令名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	CommandName *string `json:"CommandName,omitempty" name:"CommandName"`

	// Base64编码后的命令内容，长度不可超过64KB。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 命令描述。不超过120字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 命令类型，目前仅支持取值：SHELL。默认：SHELL。
	CommandType *string `json:"CommandType,omitempty" name:"CommandType"`

	// 命令执行路径，默认：/root。
	WorkingDirectory *string `json:"WorkingDirectory,omitempty" name:"WorkingDirectory"`

	// 命令超时时间，默认60秒。取值范围[1, 86400]。
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 是否启用自定义参数功能。
	// 一旦创建，此值不提供修改。
	// 默认值：false。
	EnableParameter *bool `json:"EnableParameter,omitempty" name:"EnableParameter"`

	// 启用自定义参数功能时，自定义参数的默认取值。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果InvokeCommand时未提供参数取值，将使用这里的默认值进行替换。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	DefaultParameters *string `json:"DefaultParameters,omitempty" name:"DefaultParameters"`

	// 为命令关联的标签，列表长度不超过10。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandName")
	delete(f, "Content")
	delete(f, "Description")
	delete(f, "CommandType")
	delete(f, "WorkingDirectory")
	delete(f, "Timeout")
	delete(f, "EnableParameter")
	delete(f, "DefaultParameters")
	delete(f, "Tags")
	if len(f) > 0 {
		return errors.New("CreateCommandRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCommandResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命令ID。
		CommandId *string `json:"CommandId,omitempty" name:"CommandId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCommandRequest struct {
	*tchttp.BaseRequest

	// 待删除的命令ID。
	CommandId *string `json:"CommandId,omitempty" name:"CommandId"`
}

func (r *DeleteCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandId")
	if len(f) > 0 {
		return errors.New("DeleteCommandRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCommandResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutomationAgentStatusRequest struct {
	*tchttp.BaseRequest

	// 待查询的实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 过滤条件。<br> <li> agent-status - String - 是否必填：否 -（过滤条件）按照agent状态过滤，取值：Online 在线，Offline 离线。<br> <li> environment - String - 是否必填：否 -（过滤条件）按照agent运行环境查询，取值：Linux。<br> <li> instance-id - String - 是否必填：否 -（过滤条件）按照实例ID过滤。 <br>每次请求的 `Filters` 的上限为10， `Filter.Values` 的上限为5。参数不支持同时指定 `InstanceIds` 和 `Filters` 。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAutomationAgentStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutomationAgentStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeAutomationAgentStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutomationAgentStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Agent 信息列表。
		AutomationAgentSet []*AutomationAgentInfo `json:"AutomationAgentSet,omitempty" name:"AutomationAgentSet"`

		// 符合条件的 Agent 总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutomationAgentStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutomationAgentStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCommandsRequest struct {
	*tchttp.BaseRequest

	// 命令ID列表，每次请求的上限为100。参数不支持同时指定 `CommandIds` 和 `Filters` 。
	CommandIds []*string `json:"CommandIds,omitempty" name:"CommandIds"`

	// 过滤条件。
	// <li> command-id - String - 是否必填：否 -（过滤条件）按照命令ID过滤。
	// <li> command-name - String - 是否必填：否 -（过滤条件）按照命令名称过滤。
	// <li> created-by - String - 是否必填：否 -（过滤条件）按照命令创建者过滤，取值为 TAT 或 USER，TAT 代表公共命令，USER 代表由用户创建的命令。 
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例4</li>
	// 
	// 每次请求的 `Filters` 的上限为10， `Filter.Values` 的上限为5。参数不支持同时指定 `CommandIds` 和 `Filters` 。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCommandsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommandsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeCommandsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCommandsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的命令总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 命令详情列表。
		CommandSet []*Command `json:"CommandSet,omitempty" name:"CommandSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCommandsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommandsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvocationTasksRequest struct {
	*tchttp.BaseRequest

	// 执行任务ID列表，每次请求的上限为100。参数不支持同时指定 `InvocationTaskIds` 和 `Filters`。
	InvocationTaskIds []*string `json:"InvocationTaskIds,omitempty" name:"InvocationTaskIds"`

	// 过滤条件。<br> <li> invocation-id - String - 是否必填：否 -（过滤条件）按照执行活动ID过滤。<br> <li> invocation-task-id - String - 是否必填：否 -（过滤条件）按照执行任务ID过滤。<br> <li> instance-id - String - 是否必填：否 -（过滤条件）按照实例ID过滤。 <br>每次请求的 `Filters` 的上限为10， `Filter.Values` 的上限为5。参数不支持同时指定 `InvocationTaskIds` 和 `Filters` 。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否隐藏输出，取值范围：<br><li>True：隐藏输出 <br><li>False：不隐藏 <br>默认为 True。
	HideOutput *bool `json:"HideOutput,omitempty" name:"HideOutput"`
}

func (r *DescribeInvocationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvocationTaskIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "HideOutput")
	if len(f) > 0 {
		return errors.New("DescribeInvocationTasksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvocationTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的执行任务总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 执行任务列表。
		InvocationTaskSet []*InvocationTask `json:"InvocationTaskSet,omitempty" name:"InvocationTaskSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInvocationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvocationsRequest struct {
	*tchttp.BaseRequest

	// 执行活动ID列表，每次请求的上限为100。参数不支持同时指定 `InvocationIds` 和 `Filters`。
	InvocationIds []*string `json:"InvocationIds,omitempty" name:"InvocationIds"`

	// 过滤条件。<br> <li> invocation-id - String - 是否必填：否 -（过滤条件）按照执行活动ID过滤。<br> <li> command-id - String - 是否必填：否 -（过滤条件）按照命令ID过滤。 <br>每次请求的 `Filters` 的上限为10， `Filter.Values` 的上限为5。参数不支持同时指定 `InvocationIds` 和 `Filters` 。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInvocationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvocationIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeInvocationsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInvocationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的执行活动总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 执行活动列表。
		InvocationSet []*Invocation `json:"InvocationSet,omitempty" name:"InvocationSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInvocationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeRegionsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 地域数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 地域信息列表
		RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type Invocation struct {

	// 执行活动ID。
	InvocationId *string `json:"InvocationId,omitempty" name:"InvocationId"`

	// 命令ID。
	CommandId *string `json:"CommandId,omitempty" name:"CommandId"`

	// 执行任务状态。取值范围：
	// <li> PENDING：等待下发 
	// <li> RUNNING：命令运行中
	// <li> SUCCESS：命令成功
	// <li> FAILED：命令失败
	// <li> TIMEOUT：命令超时
	// <li> PARTIAL_FAILED：命令部分失败
	InvocationStatus *string `json:"InvocationStatus,omitempty" name:"InvocationStatus"`

	// 执行任务信息列表。
	InvocationTaskBasicInfoSet []*InvocationTaskBasicInfo `json:"InvocationTaskBasicInfoSet,omitempty" name:"InvocationTaskBasicInfoSet"`

	// 执行活动描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 执行活动开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 执行活动结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 执行活动创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 执行活动更新时间。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 自定义参数取值。
	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`

	// 自定义参数的默认取值。
	DefaultParameters *string `json:"DefaultParameters,omitempty" name:"DefaultParameters"`
}

type InvocationTask struct {

	// 执行活动ID。
	InvocationId *string `json:"InvocationId,omitempty" name:"InvocationId"`

	// 执行任务ID。
	InvocationTaskId *string `json:"InvocationTaskId,omitempty" name:"InvocationTaskId"`

	// 命令ID。
	CommandId *string `json:"CommandId,omitempty" name:"CommandId"`

	// 执行任务状态。取值范围：
	// <li> PENDING：等待下发 
	// <li> DELIVERING：下发中
	// <li> DELIVER_DELAYED：延时下发 
	// <li> DELIVER_FAILED：下发失败
	// <li> RUNNING：命令运行中
	// <li> SUCCESS：命令成功
	// <li> FAILED：命令失败
	// <li> TIMEOUT：命令超时
	// <li> TASK_TIMEOUT：执行任务超时
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 执行结果。
	TaskResult *TaskResult `json:"TaskResult,omitempty" name:"TaskResult"`

	// 执行任务开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 执行任务结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 更新时间。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 执行任务所执行的命令详情。
	CommandDocument *CommandDocument `json:"CommandDocument,omitempty" name:"CommandDocument"`
}

type InvocationTaskBasicInfo struct {

	// 执行任务ID。
	InvocationTaskId *string `json:"InvocationTaskId,omitempty" name:"InvocationTaskId"`

	// 执行任务状态。取值范围：
	// <li> PENDING：等待下发 
	// <li> DELIVERING：下发中
	// <li> DELIVER_DELAYED：延时下发 
	// <li> DELIVER_FAILED：下发失败
	// <li> RUNNING：命令运行中
	// <li> SUCCESS：命令成功
	// <li> FAILED：命令失败
	// <li> TIMEOUT：命令超时
	// <li> TASK_TIMEOUT：执行任务超时
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type InvokeCommandRequest struct {
	*tchttp.BaseRequest

	// 待触发的命令ID。
	CommandId *string `json:"CommandId,omitempty" name:"CommandId"`

	// 待执行命令的实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Command 的自定义参数。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果未提供该参数取值，将使用 Command 的 DefaultParameters 进行替换。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`
}

func (r *InvokeCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandId")
	delete(f, "InstanceIds")
	delete(f, "Parameters")
	if len(f) > 0 {
		return errors.New("InvokeCommandRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InvokeCommandResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 执行活动ID。
		InvocationId *string `json:"InvocationId,omitempty" name:"InvocationId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InvokeCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCommandRequest struct {
	*tchttp.BaseRequest

	// 命令ID。
	CommandId *string `json:"CommandId,omitempty" name:"CommandId"`

	// 命令名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	CommandName *string `json:"CommandName,omitempty" name:"CommandName"`

	// 命令描述。不超过120字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// Base64编码后的命令内容，长度不可超过64KB。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 命令类型，目前仅支持取值：SHELL。
	CommandType *string `json:"CommandType,omitempty" name:"CommandType"`

	// 命令执行路径，默认：`/root`。
	WorkingDirectory *string `json:"WorkingDirectory,omitempty" name:"WorkingDirectory"`

	// 命令超时时间，默认60秒。取值范围[1, 86400]。
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 启用自定义参数功能时，自定义参数的默认取值。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// 采取整体全覆盖式修改，即修改时必须提供所有新默认值。
	// 必须 Command 的 EnableParameter 为 true 时，才允许修改这个值。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	DefaultParameters *string `json:"DefaultParameters,omitempty" name:"DefaultParameters"`
}

func (r *ModifyCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommandId")
	delete(f, "CommandName")
	delete(f, "Description")
	delete(f, "Content")
	delete(f, "CommandType")
	delete(f, "WorkingDirectory")
	delete(f, "Timeout")
	delete(f, "DefaultParameters")
	if len(f) > 0 {
		return errors.New("ModifyCommandRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCommandResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreviewReplacedCommandContentRequest struct {
	*tchttp.BaseRequest

	// 本次预览采用的自定义参数。字段类型为 json encoded string，如：{\"varA\": \"222\"}。
	// key 为自定义参数名称，value 为该参数的取值。kv 均为字符串型。
	// 自定义参数最多 20 个。
	// 自定义参数名称需符合以下规范：字符数目上限 64，可选范围【a-zA-Z0-9-_】。
	// 如果将预览的 CommandId 设置过 DefaultParameters，本参数可以为空。
	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`

	// 要进行替换预览的命令，如果有设置过 DefaultParameters，会与 Parameters 进行叠加，后者覆盖前者。
	// CommandId 与 Content，必须且只能提供一个。
	CommandId *string `json:"CommandId,omitempty" name:"CommandId"`

	// 要预览的命令内容，经 Base64 编码，长度不可超过 64KB。
	// CommandId 与 Content，必须且只能提供一个。
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *PreviewReplacedCommandContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewReplacedCommandContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Parameters")
	delete(f, "CommandId")
	delete(f, "Content")
	if len(f) > 0 {
		return errors.New("PreviewReplacedCommandContentRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PreviewReplacedCommandContentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 自定义参数替换后的，经Base64编码的命令内容。
		ReplacedContent *string `json:"ReplacedContent,omitempty" name:"ReplacedContent"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreviewReplacedCommandContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewReplacedCommandContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 地域名称，例如，ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域描述，例如: 广州
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域是否可用状态，AVAILABLE 代表可用
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
}

type RunCommandRequest struct {
	*tchttp.BaseRequest

	// Base64编码后的命令内容，长度不可超过64KB。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 待执行命令的实例ID列表。 支持实例类型：
	// <li> CVM
	// <li> LIGHTHOUSE
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 命令名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	CommandName *string `json:"CommandName,omitempty" name:"CommandName"`

	// 命令描述。不超过120字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 命令类型，目前仅支持取值：SHELL。默认：SHELL。
	CommandType *string `json:"CommandType,omitempty" name:"CommandType"`

	// 命令执行路径，默认：/root。
	WorkingDirectory *string `json:"WorkingDirectory,omitempty" name:"WorkingDirectory"`

	// 命令超时时间，默认60秒。取值范围[1, 86400]。
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 是否保存命令，取值范围：
	// <li> True：保存
	// <li> False：不保存
	// 默认为 False。
	SaveCommand *bool `json:"SaveCommand,omitempty" name:"SaveCommand"`

	// 是否启用自定义参数功能。
	// 一旦创建，此值不提供修改。
	// 默认值：false。
	EnableParameter *bool `json:"EnableParameter,omitempty" name:"EnableParameter"`

	// 启用自定义参数功能时，自定义参数的默认取值。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果 Parameters 未提供，将使用这里的默认值进行替换。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	DefaultParameters *string `json:"DefaultParameters,omitempty" name:"DefaultParameters"`

	// Command 的自定义参数。字段类型为json encoded string。如：{\"varA\": \"222\"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果未提供该参数取值，将使用 DefaultParameters 进行替换。
	// 自定义参数最多20个。
	// 自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`

	// 如果保存命令，可为命令设置标签。列表长度不超过10。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *RunCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "InstanceIds")
	delete(f, "CommandName")
	delete(f, "Description")
	delete(f, "CommandType")
	delete(f, "WorkingDirectory")
	delete(f, "Timeout")
	delete(f, "SaveCommand")
	delete(f, "EnableParameter")
	delete(f, "DefaultParameters")
	delete(f, "Parameters")
	delete(f, "Tags")
	if len(f) > 0 {
		return errors.New("RunCommandRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RunCommandResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命令ID。
		CommandId *string `json:"CommandId,omitempty" name:"CommandId"`

		// 执行活动ID。
		InvocationId *string `json:"InvocationId,omitempty" name:"InvocationId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签键。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TaskResult struct {

	// 命令执行ExitCode。
	ExitCode *int64 `json:"ExitCode,omitempty" name:"ExitCode"`

	// Base64编码后的命令输出。最大长度24KB。
	Output *string `json:"Output,omitempty" name:"Output"`

	// 命令执行开始时间。
	ExecStartTime *string `json:"ExecStartTime,omitempty" name:"ExecStartTime"`

	// 命令执行结束时间。
	ExecEndTime *string `json:"ExecEndTime,omitempty" name:"ExecEndTime"`

	// 命令最终输出被截断的字节数。
	Dropped *uint64 `json:"Dropped,omitempty" name:"Dropped"`
}
