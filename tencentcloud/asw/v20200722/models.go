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

package v20200722

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateFlowServiceRequest struct {
	*tchttp.BaseRequest

	// 定义文本（JSON格式）
	Definition *string `json:"Definition,omitempty" name:"Definition"`

	// 状态机所属服务名
	FlowServiceName *string `json:"FlowServiceName,omitempty" name:"FlowServiceName"`

	// 是不是新的角色
	IsNewRole *bool `json:"IsNewRole,omitempty" name:"IsNewRole"`

	// 状态机类型（EXPRESS，STANDARD）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 状态机所属服务中文名
	FlowServiceChineseName *string `json:"FlowServiceChineseName,omitempty" name:"FlowServiceChineseName"`

	// 角色资源名
	RoleResource *string `json:"RoleResource,omitempty" name:"RoleResource"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateFlowServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFlowServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFlowServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态机所属服务资源
		FlowServiceResource *string `json:"FlowServiceResource,omitempty" name:"FlowServiceResource"`

		// 生成日期
		CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFlowServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFlowServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExecutionRequest struct {
	*tchttp.BaseRequest

	// 执行资源名
	ExecutionResourceName *string `json:"ExecutionResourceName,omitempty" name:"ExecutionResourceName"`
}

func (r *DescribeExecutionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExecutionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExecutionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 执行资源名
		ExecutionResourceName *string `json:"ExecutionResourceName,omitempty" name:"ExecutionResourceName"`

		// 资源名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 执行开始时间，毫秒
		StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

		// 执行结束时间，毫秒
		StopDate *string `json:"StopDate,omitempty" name:"StopDate"`

		// 状态机资源名
		StateMachineResourceName *string `json:"StateMachineResourceName,omitempty" name:"StateMachineResourceName"`

		// 执行状态。INIT，RUNNING，SUCCEED，FAILED，TERMINATED
		Status *string `json:"Status,omitempty" name:"Status"`

		// 执行的输入
	// 注意：此字段可能返回 null，表示取不到有效值。
		Input *string `json:"Input,omitempty" name:"Input"`

		// 执行的输出
	// 注意：此字段可能返回 null，表示取不到有效值。
		Output *string `json:"Output,omitempty" name:"Output"`

		// 启动执行时，状态机的定义
		ExecutionDefinition *string `json:"ExecutionDefinition,omitempty" name:"ExecutionDefinition"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExecutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExecutionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExecutionsRequest struct {
	*tchttp.BaseRequest

	// 状态机资源名
	StateMachineResourceName *string `json:"StateMachineResourceName,omitempty" name:"StateMachineResourceName"`

	// 页大小，最大100
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页序号，从1开始
	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 按状态过滤条件，INIT，RUNNING，SUCCEED，FAILED，TERMINATED
	FilterExecutionStatus *string `json:"FilterExecutionStatus,omitempty" name:"FilterExecutionStatus"`

	// 按执行名过滤条件
	FilterExecutionResourceName *string `json:"FilterExecutionResourceName,omitempty" name:"FilterExecutionResourceName"`
}

func (r *DescribeExecutionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExecutionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExecutionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExecutionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExecutionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowServiceDetailRequest struct {
	*tchttp.BaseRequest

	// 状态机所属服务资源名
	FlowServiceResource *string `json:"FlowServiceResource,omitempty" name:"FlowServiceResource"`
}

func (r *DescribeFlowServiceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowServiceDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowServiceDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态机所属服务名
		FlowServiceName *string `json:"FlowServiceName,omitempty" name:"FlowServiceName"`

		// 状态机状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 定义文本（JSON格式）
	// 注意：此字段可能返回 null，表示取不到有效值。
		Definition *string `json:"Definition,omitempty" name:"Definition"`

		// 角色资源名
	// 注意：此字段可能返回 null，表示取不到有效值。
		RoleResource *string `json:"RoleResource,omitempty" name:"RoleResource"`

		// 状态机的类型，可以为 （EXPRESS/STANDARD）
		Type *string `json:"Type,omitempty" name:"Type"`

		// 生成时间
		CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

		// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 状态机所属服务中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
		FlowServiceChineseName *string `json:"FlowServiceChineseName,omitempty" name:"FlowServiceChineseName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowServiceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowServiceDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyFlowServiceRequest struct {
	*tchttp.BaseRequest

	// 状态机资源名
	FlowServiceResource *string `json:"FlowServiceResource,omitempty" name:"FlowServiceResource"`

	// 定义JSON
	Definition *string `json:"Definition,omitempty" name:"Definition"`

	// 状态机所属服务名
	FlowServiceName *string `json:"FlowServiceName,omitempty" name:"FlowServiceName"`

	// 状态机所属服务中文名
	FlowServiceChineseName *string `json:"FlowServiceChineseName,omitempty" name:"FlowServiceChineseName"`

	// 是否是新角色
	IsNewRole *bool `json:"IsNewRole,omitempty" name:"IsNewRole"`

	// 状态机类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 角色资源名
	RoleResource *string `json:"RoleResource,omitempty" name:"RoleResource"`

	// 状态机备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyFlowServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFlowServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyFlowServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态机资源名
		FlowServiceResource *string `json:"FlowServiceResource,omitempty" name:"FlowServiceResource"`

		// 更新时间
		UpdateDate *string `json:"UpdateDate,omitempty" name:"UpdateDate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFlowServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFlowServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartExecutionRequest struct {
	*tchttp.BaseRequest

	// 状态机资源名
	StateMachineResourceName *string `json:"StateMachineResourceName,omitempty" name:"StateMachineResourceName"`

	// 输入参数
	Input *string `json:"Input,omitempty" name:"Input"`

	// 本次执行名。如果不填，系统会自动生成。如果填，应保证状态机下唯一
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *StartExecutionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartExecutionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartExecutionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 执行资源名
		ExecutionResourceName *string `json:"ExecutionResourceName,omitempty" name:"ExecutionResourceName"`

		// 执行开始时间
		StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartExecutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartExecutionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
