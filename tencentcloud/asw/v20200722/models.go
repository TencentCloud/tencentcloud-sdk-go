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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateFlowServiceRequestParams struct {
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

	// 角色资源名, 比如: qcs::cam::uin/20103392:roleName/SomeRoleForYourStateMachine
	RoleResource *string `json:"RoleResource,omitempty" name:"RoleResource"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否开启CLS日志投递功能
	EnableCLS *bool `json:"EnableCLS,omitempty" name:"EnableCLS"`

	// 该状态机的默认输入
	Input *string `json:"Input,omitempty" name:"Input"`
}

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

	// 角色资源名, 比如: qcs::cam::uin/20103392:roleName/SomeRoleForYourStateMachine
	RoleResource *string `json:"RoleResource,omitempty" name:"RoleResource"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否开启CLS日志投递功能
	EnableCLS *bool `json:"EnableCLS,omitempty" name:"EnableCLS"`

	// 该状态机的默认输入
	Input *string `json:"Input,omitempty" name:"Input"`
}

func (r *CreateFlowServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Definition")
	delete(f, "FlowServiceName")
	delete(f, "IsNewRole")
	delete(f, "Type")
	delete(f, "FlowServiceChineseName")
	delete(f, "RoleResource")
	delete(f, "Description")
	delete(f, "EnableCLS")
	delete(f, "Input")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowServiceResponseParams struct {
	// 状态机所属服务资源
	FlowServiceResource *string `json:"FlowServiceResource,omitempty" name:"FlowServiceResource"`

	// 生成日期
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFlowServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowServiceResponseParams `json:"Response"`
}

func (r *CreateFlowServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExecutionHistoryRequestParams struct {
	// 执行资源名
	ExecutionResourceName *string `json:"ExecutionResourceName,omitempty" name:"ExecutionResourceName"`
}

type DescribeExecutionHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 执行资源名
	ExecutionResourceName *string `json:"ExecutionResourceName,omitempty" name:"ExecutionResourceName"`
}

func (r *DescribeExecutionHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExecutionHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExecutionResourceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExecutionHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExecutionHistoryResponseParams struct {
	// 执行的事件列表
	Events []*ExecutionEvent `json:"Events,omitempty" name:"Events"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExecutionHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExecutionHistoryResponseParams `json:"Response"`
}

func (r *DescribeExecutionHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExecutionHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExecutionRequestParams struct {
	// 执行资源名
	ExecutionResourceName *string `json:"ExecutionResourceName,omitempty" name:"ExecutionResourceName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExecutionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExecutionResourceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExecutionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExecutionResponseParams struct {
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
}

type DescribeExecutionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExecutionResponseParams `json:"Response"`
}

func (r *DescribeExecutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExecutionsRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExecutionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StateMachineResourceName")
	delete(f, "PageSize")
	delete(f, "PageIndex")
	delete(f, "FilterExecutionStatus")
	delete(f, "FilterExecutionResourceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExecutionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExecutionsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExecutionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExecutionsResponseParams `json:"Response"`
}

func (r *DescribeExecutionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExecutionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowServiceDetailRequestParams struct {
	// 状态机所属服务资源名
	FlowServiceResource *string `json:"FlowServiceResource,omitempty" name:"FlowServiceResource"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowServiceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowServiceResource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowServiceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowServiceDetailResponseParams struct {
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

	// 是否开启日志CLS服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableCLS *bool `json:"EnableCLS,omitempty" name:"EnableCLS"`

	// CLS日志查看地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CLSUrl *string `json:"CLSUrl,omitempty" name:"CLSUrl"`

	// 工作流提示输入
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowInput *string `json:"FlowInput,omitempty" name:"FlowInput"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowServiceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowServiceDetailResponseParams `json:"Response"`
}

func (r *DescribeFlowServiceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowServiceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowServicesRequestParams struct {
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filter.Values的上限为5。参数名字仅支持FlowServiceName， Status, Type三种情况
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeFlowServicesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filter.Values的上限为5。参数名字仅支持FlowServiceName， Status, Type三种情况
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeFlowServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowServicesResponseParams struct {
	// 用户的状态机列表
	FlowServiceSet []*StateMachine `json:"FlowServiceSet,omitempty" name:"FlowServiceSet"`

	// 用户的状态机总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowServicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowServicesResponseParams `json:"Response"`
}

func (r *DescribeFlowServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecutionEvent struct {
	// 执行资源名
	ExecutionResourceName *string `json:"ExecutionResourceName,omitempty" name:"ExecutionResourceName"`

	// 自增序号
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 事件类型
	EventCategory *string `json:"EventCategory,omitempty" name:"EventCategory"`

	// 步骤节点名称
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// 该步骤引用的资源名
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 该事件发生时间，毫秒
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 事件内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exception *string `json:"Exception,omitempty" name:"Exception"`
}

type Filter struct {
	// 过滤器名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤器值的数组
	Values []*string `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type ModifyFlowServiceRequestParams struct {
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

	// 是否允许日志投递
	EnableCLS *bool `json:"EnableCLS,omitempty" name:"EnableCLS"`
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

	// 是否允许日志投递
	EnableCLS *bool `json:"EnableCLS,omitempty" name:"EnableCLS"`
}

func (r *ModifyFlowServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFlowServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowServiceResource")
	delete(f, "Definition")
	delete(f, "FlowServiceName")
	delete(f, "FlowServiceChineseName")
	delete(f, "IsNewRole")
	delete(f, "Type")
	delete(f, "RoleResource")
	delete(f, "Description")
	delete(f, "EnableCLS")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFlowServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFlowServiceResponseParams struct {
	// 状态机资源名
	FlowServiceResource *string `json:"FlowServiceResource,omitempty" name:"FlowServiceResource"`

	// 更新时间
	UpdateDate *string `json:"UpdateDate,omitempty" name:"UpdateDate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyFlowServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFlowServiceResponseParams `json:"Response"`
}

func (r *ModifyFlowServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFlowServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartExecutionRequestParams struct {
	// 状态机资源名
	StateMachineResourceName *string `json:"StateMachineResourceName,omitempty" name:"StateMachineResourceName"`

	// 输入参数，内容为JsonObject，长度不大于524288字符。
	Input *string `json:"Input,omitempty" name:"Input"`

	// 本次执行名。如果不填，系统会自动生成。如果填，应保证状态机下唯一
	Name *string `json:"Name,omitempty" name:"Name"`
}

type StartExecutionRequest struct {
	*tchttp.BaseRequest
	
	// 状态机资源名
	StateMachineResourceName *string `json:"StateMachineResourceName,omitempty" name:"StateMachineResourceName"`

	// 输入参数，内容为JsonObject，长度不大于524288字符。
	Input *string `json:"Input,omitempty" name:"Input"`

	// 本次执行名。如果不填，系统会自动生成。如果填，应保证状态机下唯一
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *StartExecutionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartExecutionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StateMachineResourceName")
	delete(f, "Input")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartExecutionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartExecutionResponseParams struct {
	// 执行资源名
	ExecutionResourceName *string `json:"ExecutionResourceName,omitempty" name:"ExecutionResourceName"`

	// 执行开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartExecutionResponse struct {
	*tchttp.BaseResponse
	Response *StartExecutionResponseParams `json:"Response"`
}

func (r *StartExecutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StateMachine struct {
	// 状态机资源
	FlowServiceResource *string `json:"FlowServiceResource,omitempty" name:"FlowServiceResource"`

	// 状态机类型。EXPRESS，STANDARD
	Type *string `json:"Type,omitempty" name:"Type"`

	// 状态机名称
	FlowServiceName *string `json:"FlowServiceName,omitempty" name:"FlowServiceName"`

	// 状态机中文名
	FlowServiceChineseName *string `json:"FlowServiceChineseName,omitempty" name:"FlowServiceChineseName"`

	// 创建时间。timestamp
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 修改时间。timestamp
	ModifyDate *string `json:"ModifyDate,omitempty" name:"ModifyDate"`

	// 状态机状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建者的subAccountUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 修改者的subAccountUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`

	// 状态机id
	FlowServiceId *string `json:"FlowServiceId,omitempty" name:"FlowServiceId"`

	// 模板id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

// Predefined struct for user
type StopExecutionRequestParams struct {
	// 执行名称
	ExecutionQrn *string `json:"ExecutionQrn,omitempty" name:"ExecutionQrn"`
}

type StopExecutionRequest struct {
	*tchttp.BaseRequest
	
	// 执行名称
	ExecutionQrn *string `json:"ExecutionQrn,omitempty" name:"ExecutionQrn"`
}

func (r *StopExecutionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopExecutionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExecutionQrn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopExecutionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopExecutionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopExecutionResponse struct {
	*tchttp.BaseResponse
	Response *StopExecutionResponseParams `json:"Response"`
}

func (r *StopExecutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}