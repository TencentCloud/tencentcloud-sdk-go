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

package v20210622

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AppOverviewScore struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>端</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>加权总分，全部缺失时省略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalScore *float64 `json:"TotalScore,omitnil,omitempty" name:"TotalScore"`

	// <p>计分项</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OverviewScoreItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type ApplicationItem struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用 Key</p>
	ProjectKey *string `json:"ProjectKey,omitnil,omitempty" name:"ProjectKey"`

	// <p>业务系统 ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>业务系统名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>应用名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>端类型</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>描述</p>
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// <p>是否排除已忽略错误</p>
	ExcludeIgnoredErrors *bool `json:"ExcludeIgnoredErrors,omitnil,omitempty" name:"ExcludeIgnoredErrors"`

	// <p>状态：2 运行 / 6 停止 / 7 删除</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>创建时间</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>该应用绑定的腾讯云 APM 实例 ID 列表。空数组表示未绑定。</p>
	ApmInstanceIds []*string `json:"ApmInstanceIds,omitnil,omitempty" name:"ApmInstanceIds"`
}

type BusinessSystemItem struct {
	// <p>业务系统 ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>描述</p>
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`

	// <p>状态：2 运行 / 6 停止 / 7 删除</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>状态： 2（隔离），1（使用中），3（已销毁）</p>
	ChargeStatus *int64 `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// <p>创建时间</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>更新时间</p>
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type CompareCondition struct {
	// App版本
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 筛选条件
	Filters *Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

// Predefined struct for user
type CreateFieldTransformRuleRequestParams struct {
	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>目标字段</p>
	TargetField *string `json:"TargetField,omitnil,omitempty" name:"TargetField"`

	// <p>匹配方式</p>
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// <p>匹配规则</p>
	MatchPattern *string `json:"MatchPattern,omitnil,omitempty" name:"MatchPattern"`

	// <p>转换类型</p>
	TransformAction *string `json:"TransformAction,omitnil,omitempty" name:"TransformAction"`

	// <p>规则名</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>替换值</p>
	ReplaceValue *string `json:"ReplaceValue,omitnil,omitempty" name:"ReplaceValue"`

	// <p>优先级</p>
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// <p>是否开启</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type CreateFieldTransformRuleRequest struct {
	*tchttp.BaseRequest
	
	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>目标字段</p>
	TargetField *string `json:"TargetField,omitnil,omitempty" name:"TargetField"`

	// <p>匹配方式</p>
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// <p>匹配规则</p>
	MatchPattern *string `json:"MatchPattern,omitnil,omitempty" name:"MatchPattern"`

	// <p>转换类型</p>
	TransformAction *string `json:"TransformAction,omitnil,omitempty" name:"TransformAction"`

	// <p>规则名</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>替换值</p>
	ReplaceValue *string `json:"ReplaceValue,omitnil,omitempty" name:"ReplaceValue"`

	// <p>优先级</p>
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// <p>是否开启</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *CreateFieldTransformRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFieldTransformRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TargetField")
	delete(f, "MatchType")
	delete(f, "MatchPattern")
	delete(f, "TransformAction")
	delete(f, "Name")
	delete(f, "ReplaceValue")
	delete(f, "Priority")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFieldTransformRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFieldTransformRuleResponseParams struct {
	// <p>结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误消息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFieldTransformRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateFieldTransformRuleResponseParams `json:"Response"`
}

func (r *CreateFieldTransformRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFieldTransformRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseFileRequestParams struct {
	// 项目 id
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 文件信息列表
	Files []*ReleaseFile `json:"Files,omitnil,omitempty" name:"Files"`
}

type CreateReleaseFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目 id
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 文件信息列表
	Files []*ReleaseFile `json:"Files,omitnil,omitempty" name:"Files"`
}

func (r *CreateReleaseFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "Files")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReleaseFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseFileResponseParams struct {
	// 调用结果
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReleaseFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateReleaseFileResponseParams `json:"Response"`
}

func (r *CreateReleaseFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStarProjectRequestParams struct {
	// 实例ID：taw-123
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type CreateStarProjectRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID：taw-123
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *CreateStarProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStarProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStarProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStarProjectResponseParams struct {
	// 接口返回信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStarProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateStarProjectResponseParams `json:"Response"`
}

func (r *CreateStarProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStarProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDEMApplicationRequestParams struct {
	// <p>业务系统 ID（必填）</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>应用名称（必填）</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>应用类型。取值：web、wx-miniprogram、uniapp、ios、android、electron、rn（React Native）。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>应用描述（可选）</p>
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// <p>APM 实例 ID</p>
	ApmInstanceIds []*string `json:"ApmInstanceIds,omitnil,omitempty" name:"ApmInstanceIds"`
}

type CreateTDEMApplicationRequest struct {
	*tchttp.BaseRequest
	
	// <p>业务系统 ID（必填）</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>应用名称（必填）</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>应用类型。取值：web、wx-miniprogram、uniapp、ios、android、electron、rn（React Native）。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>应用描述（可选）</p>
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// <p>APM 实例 ID</p>
	ApmInstanceIds []*string `json:"ApmInstanceIds,omitnil,omitempty" name:"ApmInstanceIds"`
}

func (r *CreateTDEMApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDEMApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Desc")
	delete(f, "ApmInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDEMApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDEMApplicationResponseParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用 Key</p>
	ProjectKey *string `json:"ProjectKey,omitnil,omitempty" name:"ProjectKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTDEMApplicationResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDEMApplicationResponseParams `json:"Response"`
}

func (r *CreateTDEMApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDEMApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDEMBusinessSystemRequestParams struct {
	// <p>业务系统名称（必填）</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>业务系统描述（可选）</p>
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`

	// <p>标签列表</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateTDEMBusinessSystemRequest struct {
	*tchttp.BaseRequest
	
	// <p>业务系统名称（必填）</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>业务系统描述（可选）</p>
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`

	// <p>标签列表</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateTDEMBusinessSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDEMBusinessSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "InstanceDesc")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDEMBusinessSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDEMBusinessSystemResponseParams struct {
	// <p>业务系统 ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>业务系统状态</p><p>状态：2 运行 / 6 停止 / 7 删除</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTDEMBusinessSystemResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDEMBusinessSystemResponseParams `json:"Response"`
}

func (r *CreateTDEMBusinessSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDEMBusinessSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDEMDialTaskRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>任务名称，最长 128 个字符</p>
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>探测类型。取值说明：icmp、tcp。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>执行时机。取值说明：realtime、after_request、after_error、after_slow、after_dns_slow。tcp 不能使用 after_dns_slow。</p>
	ExecutionTime *string `json:"ExecutionTime,omitnil,omitempty" name:"ExecutionTime"`

	// <p>单次探测超时，单位毫秒。范围 1000 到 20000。</p>
	TimeoutMs *int64 `json:"TimeoutMs,omitnil,omitempty" name:"TimeoutMs"`

	// <p>探测周期，单位毫秒，必须大于 0。</p>
	PeriodMs *int64 `json:"PeriodMs,omitnil,omitempty" name:"PeriodMs"`

	// <p>发包次数。范围 3 到 40。</p>
	PackageCount *int64 `json:"PackageCount,omitnil,omitempty" name:"PackageCount"`

	// <p>任务状态。取值说明：active、paused。不传时默认为 active。</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>监听域名。ExecutionTime 不是 realtime 时必填。</p>
	WatchDomain *string `json:"WatchDomain,omitnil,omitempty" name:"WatchDomain"`

	// <p>拨测目标。域名、IP 或 ldns。after_dns_slow 固定为 ldns，其它时机必填。</p>
	DialDomain *string `json:"DialDomain,omitnil,omitempty" name:"DialDomain"`

	// <p>TCP 端口，1 到 65535。Type 为 tcp 时必填。</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>实时拨测间隔，单位毫秒。ExecutionTime 为 realtime 时必填。取值说明：60000、300000、600000、900000、1800000、3600000。</p>
	FrequencyMs *int64 `json:"FrequencyMs,omitnil,omitempty" name:"FrequencyMs"`

	// <p>ICMP 包大小，单位字节。范围 32 到 1024。不传时默认 64。</p>
	PackageSize *int64 `json:"PackageSize,omitnil,omitempty" name:"PackageSize"`

	// <p>是否做路由追踪。取值说明：0 否，1 是。不传时为 0。</p>
	NeedMTR *int64 `json:"NeedMTR,omitnil,omitempty" name:"NeedMTR"`

	// <p>路由追踪超时，单位毫秒。NeedMTR 为 1 时必填。</p>
	MTRTimeoutMs *int64 `json:"MTRTimeoutMs,omitnil,omitempty" name:"MTRTimeoutMs"`

	// <p>路由追踪最大跳数。NeedMTR 为 1 时填写。</p>
	MTRMaxHop *int64 `json:"MTRMaxHop,omitnil,omitempty" name:"MTRMaxHop"`

	// <p>每一跳发包次数。范围 3 到 40。</p>
	MTRPingCount *int64 `json:"MTRPingCount,omitnil,omitempty" name:"MTRPingCount"`

	// <p>慢请求阈值，单位毫秒。after_slow、after_dns_slow 不传时默认 2000。</p>
	SlowThresholdMs *int64 `json:"SlowThresholdMs,omitnil,omitempty" name:"SlowThresholdMs"`

	// <p>错误类型过滤。after_error 时使用。</p>
	ErrorType *string `json:"ErrorType,omitnil,omitempty" name:"ErrorType"`

	// <p>允许启动的客户端地区，逗号分隔，名称必须与分析页客户端地区完全一致。不传或空字符串表示不限制。</p>
	Regions *string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// <p>允许启动的客户端运营商，逗号分隔，名称必须与分析页客户端运营商完全一致。不传或空字符串表示不限制。</p>
	Operators *string `json:"Operators,omitnil,omitempty" name:"Operators"`

	// <p>允许启动的用户 ID，逗号分隔。不传或空字符串表示不限制。</p>
	UserIds *string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// <p>允许启动的设备 ID，逗号分隔。不传或空字符串表示不限制。</p>
	DeviceIds *string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`
}

type CreateTDEMDialTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>任务名称，最长 128 个字符</p>
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>探测类型。取值说明：icmp、tcp。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>执行时机。取值说明：realtime、after_request、after_error、after_slow、after_dns_slow。tcp 不能使用 after_dns_slow。</p>
	ExecutionTime *string `json:"ExecutionTime,omitnil,omitempty" name:"ExecutionTime"`

	// <p>单次探测超时，单位毫秒。范围 1000 到 20000。</p>
	TimeoutMs *int64 `json:"TimeoutMs,omitnil,omitempty" name:"TimeoutMs"`

	// <p>探测周期，单位毫秒，必须大于 0。</p>
	PeriodMs *int64 `json:"PeriodMs,omitnil,omitempty" name:"PeriodMs"`

	// <p>发包次数。范围 3 到 40。</p>
	PackageCount *int64 `json:"PackageCount,omitnil,omitempty" name:"PackageCount"`

	// <p>任务状态。取值说明：active、paused。不传时默认为 active。</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>监听域名。ExecutionTime 不是 realtime 时必填。</p>
	WatchDomain *string `json:"WatchDomain,omitnil,omitempty" name:"WatchDomain"`

	// <p>拨测目标。域名、IP 或 ldns。after_dns_slow 固定为 ldns，其它时机必填。</p>
	DialDomain *string `json:"DialDomain,omitnil,omitempty" name:"DialDomain"`

	// <p>TCP 端口，1 到 65535。Type 为 tcp 时必填。</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>实时拨测间隔，单位毫秒。ExecutionTime 为 realtime 时必填。取值说明：60000、300000、600000、900000、1800000、3600000。</p>
	FrequencyMs *int64 `json:"FrequencyMs,omitnil,omitempty" name:"FrequencyMs"`

	// <p>ICMP 包大小，单位字节。范围 32 到 1024。不传时默认 64。</p>
	PackageSize *int64 `json:"PackageSize,omitnil,omitempty" name:"PackageSize"`

	// <p>是否做路由追踪。取值说明：0 否，1 是。不传时为 0。</p>
	NeedMTR *int64 `json:"NeedMTR,omitnil,omitempty" name:"NeedMTR"`

	// <p>路由追踪超时，单位毫秒。NeedMTR 为 1 时必填。</p>
	MTRTimeoutMs *int64 `json:"MTRTimeoutMs,omitnil,omitempty" name:"MTRTimeoutMs"`

	// <p>路由追踪最大跳数。NeedMTR 为 1 时填写。</p>
	MTRMaxHop *int64 `json:"MTRMaxHop,omitnil,omitempty" name:"MTRMaxHop"`

	// <p>每一跳发包次数。范围 3 到 40。</p>
	MTRPingCount *int64 `json:"MTRPingCount,omitnil,omitempty" name:"MTRPingCount"`

	// <p>慢请求阈值，单位毫秒。after_slow、after_dns_slow 不传时默认 2000。</p>
	SlowThresholdMs *int64 `json:"SlowThresholdMs,omitnil,omitempty" name:"SlowThresholdMs"`

	// <p>错误类型过滤。after_error 时使用。</p>
	ErrorType *string `json:"ErrorType,omitnil,omitempty" name:"ErrorType"`

	// <p>允许启动的客户端地区，逗号分隔，名称必须与分析页客户端地区完全一致。不传或空字符串表示不限制。</p>
	Regions *string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// <p>允许启动的客户端运营商，逗号分隔，名称必须与分析页客户端运营商完全一致。不传或空字符串表示不限制。</p>
	Operators *string `json:"Operators,omitnil,omitempty" name:"Operators"`

	// <p>允许启动的用户 ID，逗号分隔。不传或空字符串表示不限制。</p>
	UserIds *string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// <p>允许启动的设备 ID，逗号分隔。不传或空字符串表示不限制。</p>
	DeviceIds *string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`
}

func (r *CreateTDEMDialTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDEMDialTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "TaskName")
	delete(f, "Type")
	delete(f, "ExecutionTime")
	delete(f, "TimeoutMs")
	delete(f, "PeriodMs")
	delete(f, "PackageCount")
	delete(f, "Status")
	delete(f, "WatchDomain")
	delete(f, "DialDomain")
	delete(f, "Port")
	delete(f, "FrequencyMs")
	delete(f, "PackageSize")
	delete(f, "NeedMTR")
	delete(f, "MTRTimeoutMs")
	delete(f, "MTRMaxHop")
	delete(f, "MTRPingCount")
	delete(f, "SlowThresholdMs")
	delete(f, "ErrorType")
	delete(f, "Regions")
	delete(f, "Operators")
	delete(f, "UserIds")
	delete(f, "DeviceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDEMDialTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDEMDialTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTDEMDialTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDEMDialTaskResponseParams `json:"Response"`
}

func (r *CreateTDEMDialTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDEMDialTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDEMFieldTransformRuleRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>目标字段（必填）：page_url / normalized_page_url / page_title / referrer / resource_url / user_id / target_element / target_text</p>
	TargetField *string `json:"TargetField,omitnil,omitempty" name:"TargetField"`

	// <p>匹配方式（必填）：exact / prefix / contains / regex</p>
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// <p>匹配模式（必填）</p>
	MatchPattern *string `json:"MatchPattern,omitnil,omitempty" name:"MatchPattern"`

	// <p>替换动作（必填）：replace / regex_replace / mask</p>
	TransformAction *string `json:"TransformAction,omitnil,omitempty" name:"TransformAction"`

	// <p>规则名称（可选）</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>替换值（replace / regex_replace 使用）</p>
	ReplaceValue *string `json:"ReplaceValue,omitnil,omitempty" name:"ReplaceValue"`

	// <p>优先级（可选）</p>
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// <p>是否启用，默认 true</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type CreateTDEMFieldTransformRuleRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>目标字段（必填）：page_url / normalized_page_url / page_title / referrer / resource_url / user_id / target_element / target_text</p>
	TargetField *string `json:"TargetField,omitnil,omitempty" name:"TargetField"`

	// <p>匹配方式（必填）：exact / prefix / contains / regex</p>
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// <p>匹配模式（必填）</p>
	MatchPattern *string `json:"MatchPattern,omitnil,omitempty" name:"MatchPattern"`

	// <p>替换动作（必填）：replace / regex_replace / mask</p>
	TransformAction *string `json:"TransformAction,omitnil,omitempty" name:"TransformAction"`

	// <p>规则名称（可选）</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>替换值（replace / regex_replace 使用）</p>
	ReplaceValue *string `json:"ReplaceValue,omitnil,omitempty" name:"ReplaceValue"`

	// <p>优先级（可选）</p>
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// <p>是否启用，默认 true</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *CreateTDEMFieldTransformRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDEMFieldTransformRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "TargetField")
	delete(f, "MatchType")
	delete(f, "MatchPattern")
	delete(f, "TransformAction")
	delete(f, "Name")
	delete(f, "ReplaceValue")
	delete(f, "Priority")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDEMFieldTransformRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDEMFieldTransformRuleResponseParams struct {
	// <p>创建后的规则</p>
	Rule *FieldTransformRuleItem `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTDEMFieldTransformRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDEMFieldTransformRuleResponseParams `json:"Response"`
}

func (r *CreateTDEMFieldTransformRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDEMFieldTransformRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDEMFunnelRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗名</p>
	FunnelName *string `json:"FunnelName,omitnil,omitempty" name:"FunnelName"`

	// <p>漏斗步骤</p>
	Steps []*FunnelStepInput `json:"Steps,omitnil,omitempty" name:"Steps"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>漏斗状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>漏斗来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

type CreateTDEMFunnelRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗名</p>
	FunnelName *string `json:"FunnelName,omitnil,omitempty" name:"FunnelName"`

	// <p>漏斗步骤</p>
	Steps []*FunnelStepInput `json:"Steps,omitnil,omitempty" name:"Steps"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>漏斗状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>漏斗来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

func (r *CreateTDEMFunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDEMFunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FunnelName")
	delete(f, "Steps")
	delete(f, "Description")
	delete(f, "Status")
	delete(f, "Source")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDEMFunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDEMFunnelResponseParams struct {
	// <p>请求结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTDEMFunnelResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDEMFunnelResponseParams `json:"Response"`
}

func (r *CreateTDEMFunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDEMFunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWhitelistRequestParams struct {
	// 实例ID：taw-123
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 备注（暂未作字节数限制）
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// uin：业务方标识
	WhitelistUin *string `json:"WhitelistUin,omitnil,omitempty" name:"WhitelistUin"`

	// 业务方标识
	Aid *string `json:"Aid,omitnil,omitempty" name:"Aid"`
}

type CreateWhitelistRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID：taw-123
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 备注（暂未作字节数限制）
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// uin：业务方标识
	WhitelistUin *string `json:"WhitelistUin,omitnil,omitempty" name:"WhitelistUin"`

	// 业务方标识
	Aid *string `json:"Aid,omitnil,omitempty" name:"Aid"`
}

func (r *CreateWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "Remark")
	delete(f, "WhitelistUin")
	delete(f, "Aid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWhitelistResponseParams struct {
	// 消息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 白名单ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *CreateWhitelistResponseParams `json:"Response"`
}

func (r *CreateWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFieldTransformRuleRequestParams struct {
	// <p>规则id</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteFieldTransformRuleRequest struct {
	*tchttp.BaseRequest
	
	// <p>规则id</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DeleteFieldTransformRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFieldTransformRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFieldTransformRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFieldTransformRuleResponseParams struct {
	// <p>结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFieldTransformRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFieldTransformRuleResponseParams `json:"Response"`
}

func (r *DeleteFieldTransformRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFieldTransformRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceRequestParams struct {
	// 需要删除的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceResponseParams `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMiniProgramReplayPackRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>回放包id（前缀mpr_）</p>
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`
}

type DeleteMiniProgramReplayPackRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>回放包id（前缀mpr_）</p>
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`
}

func (r *DeleteMiniProgramReplayPackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMiniProgramReplayPackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "PackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMiniProgramReplayPackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMiniProgramReplayPackResponseParams struct {
	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// <p>请求结果（JSON字符串，包含 AppId/PackId/Deleted）</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMiniProgramReplayPackResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMiniProgramReplayPackResponseParams `json:"Response"`
}

func (r *DeleteMiniProgramReplayPackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMiniProgramReplayPackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectRequestParams struct {
	// 需要删除的项目 ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的项目 ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectResponseParams struct {
	// 操作信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectResponseParams `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReleaseFileRequestParams struct {
	// 文件 id
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteReleaseFileRequest struct {
	*tchttp.BaseRequest
	
	// 文件 id
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteReleaseFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReleaseFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReleaseFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReleaseFileResponseParams struct {
	// 接口请求返回字符串
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReleaseFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReleaseFileResponseParams `json:"Response"`
}

func (r *DeleteReleaseFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReleaseFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStarProjectRequestParams struct {
	// 实例ID：****-1792
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteStarProjectRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID：****-1792
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteStarProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStarProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStarProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStarProjectResponseParams struct {
	// 返回消息,请求成功才会返回，出现异常默认为null
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteStarProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStarProjectResponseParams `json:"Response"`
}

func (r *DeleteStarProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStarProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSymbolArtifactRequestParams struct {
	// <p>制品类型。取值：js_sourcemap、ios_dsym、macos_dsym、windows_pdb、android_mapping、android_native、system_symbol。</p>
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// <p>制品 ID。</p>
	ArtifactId *string `json:"ArtifactId,omitnil,omitempty" name:"ArtifactId"`

	// <p>应用 ID。删除应用级制品时必填，须与制品所属应用一致。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>兼容别名，未传 AppId 时使用。</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>兼容字段，请勿传入。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>兼容字段，请勿传入。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>系统名称。system_symbol 场景可用于归属校验。</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>系统版本。system_symbol 场景可用于归属校验。</p>
	OSVersion *string `json:"OSVersion,omitnil,omitempty" name:"OSVersion"`
}

type DeleteSymbolArtifactRequest struct {
	*tchttp.BaseRequest
	
	// <p>制品类型。取值：js_sourcemap、ios_dsym、macos_dsym、windows_pdb、android_mapping、android_native、system_symbol。</p>
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// <p>制品 ID。</p>
	ArtifactId *string `json:"ArtifactId,omitnil,omitempty" name:"ArtifactId"`

	// <p>应用 ID。删除应用级制品时必填，须与制品所属应用一致。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>兼容别名，未传 AppId 时使用。</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>兼容字段，请勿传入。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>兼容字段，请勿传入。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>系统名称。system_symbol 场景可用于归属校验。</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>系统版本。system_symbol 场景可用于归属校验。</p>
	OSVersion *string `json:"OSVersion,omitnil,omitempty" name:"OSVersion"`
}

func (r *DeleteSymbolArtifactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSymbolArtifactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Kind")
	delete(f, "ArtifactId")
	delete(f, "AppId")
	delete(f, "ProjectId")
	delete(f, "AppVersion")
	delete(f, "Platform")
	delete(f, "OS")
	delete(f, "OSVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSymbolArtifactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSymbolArtifactResponseParams struct {
	// <p>删除结果，JSON 格式字符串，包含 artifact_id、status 和 deleted_at。</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSymbolArtifactResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSymbolArtifactResponseParams `json:"Response"`
}

func (r *DeleteSymbolArtifactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSymbolArtifactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTDEMApplicationRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DeleteTDEMApplicationRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *DeleteTDEMApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTDEMApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTDEMApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTDEMApplicationResponseParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTDEMApplicationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTDEMApplicationResponseParams `json:"Response"`
}

func (r *DeleteTDEMApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTDEMApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTDEMBusinessSystemRequestParams struct {
	// <p>业务系统 ID（必填）</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteTDEMBusinessSystemRequest struct {
	*tchttp.BaseRequest
	
	// <p>业务系统 ID（必填）</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteTDEMBusinessSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTDEMBusinessSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTDEMBusinessSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTDEMBusinessSystemResponseParams struct {
	// <p>业务系统 ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTDEMBusinessSystemResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTDEMBusinessSystemResponseParams `json:"Response"`
}

func (r *DeleteTDEMBusinessSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTDEMBusinessSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTDEMDialTaskRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>任务 ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteTDEMDialTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>任务 ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteTDEMDialTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTDEMDialTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTDEMDialTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTDEMDialTaskResponseParams struct {
	// <p>已删除的任务 ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTDEMDialTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTDEMDialTaskResponseParams `json:"Response"`
}

func (r *DeleteTDEMDialTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTDEMDialTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTDEMFieldTransformRuleRequestParams struct {
	// <p>规则 ID（必填）</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DeleteTDEMFieldTransformRuleRequest struct {
	*tchttp.BaseRequest
	
	// <p>规则 ID（必填）</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *DeleteTDEMFieldTransformRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTDEMFieldTransformRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTDEMFieldTransformRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTDEMFieldTransformRuleResponseParams struct {
	// <p>规则 ID</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>是否已删除</p>
	Deleted *bool `json:"Deleted,omitnil,omitempty" name:"Deleted"`

	// <p>提示信息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTDEMFieldTransformRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTDEMFieldTransformRuleResponseParams `json:"Response"`
}

func (r *DeleteTDEMFieldTransformRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTDEMFieldTransformRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTDEMFunnelRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`
}

type DeleteTDEMFunnelRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`
}

func (r *DeleteTDEMFunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTDEMFunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FunnelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTDEMFunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTDEMFunnelResponseParams struct {
	// <p>请求结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTDEMFunnelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTDEMFunnelResponseParams `json:"Response"`
}

func (r *DeleteTDEMFunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTDEMFunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWhitelistRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 名单ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteWhitelistRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 名单ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWhitelistResponseParams struct {
	// 消息success
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWhitelistResponseParams `json:"Response"`
}

func (r *DeleteWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppDimensionMetricsRequestParams struct {
	// app 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 fields
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 业务上下文参数
	BusinessContext *string `json:"BusinessContext,omitnil,omitempty" name:"BusinessContext"`
}

type DescribeAppDimensionMetricsRequest struct {
	*tchttp.BaseRequest
	
	// app 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 fields
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 业务上下文参数
	BusinessContext *string `json:"BusinessContext,omitnil,omitempty" name:"BusinessContext"`
}

func (r *DescribeAppDimensionMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppDimensionMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "From")
	delete(f, "Fields")
	delete(f, "Filter")
	delete(f, "FilterSimple")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BusinessContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppDimensionMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppDimensionMetricsResponseParams struct {
	// 查询数据返回
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppDimensionMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppDimensionMetricsResponseParams `json:"Response"`
}

func (r *DescribeAppDimensionMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppDimensionMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppMetricsDataRequestParams struct {
	// app 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 field
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// group by 参数
	GroupByModifier *string `json:"GroupByModifier,omitnil,omitempty" name:"GroupByModifier"`
}

type DescribeAppMetricsDataRequest struct {
	*tchttp.BaseRequest
	
	// app 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 field
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// group by 参数
	GroupByModifier *string `json:"GroupByModifier,omitnil,omitempty" name:"GroupByModifier"`
}

func (r *DescribeAppMetricsDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppMetricsDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "From")
	delete(f, "Fields")
	delete(f, "Filter")
	delete(f, "FilterSimple")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "GroupByModifier")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppMetricsDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppMetricsDataResponseParams struct {
	// 查询数据返回
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppMetricsDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppMetricsDataResponseParams `json:"Response"`
}

func (r *DescribeAppMetricsDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppMetricsDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppSingleCaseDetailListRequestParams struct {
	// app 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 field
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAppSingleCaseDetailListRequest struct {
	*tchttp.BaseRequest
	
	// app 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 field
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeAppSingleCaseDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppSingleCaseDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "From")
	delete(f, "Fields")
	delete(f, "Filter")
	delete(f, "FilterSimple")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppSingleCaseDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppSingleCaseDetailListResponseParams struct {
	// 查询数据返回
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppSingleCaseDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppSingleCaseDetailListResponseParams `json:"Response"`
}

func (r *DescribeAppSingleCaseDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppSingleCaseDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppSingleCaseListRequestParams struct {
	// app 项目 ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 field
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAppSingleCaseListRequest struct {
	*tchttp.BaseRequest
	
	// app 项目 ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 查询的表名
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 查询指标 field
	Fields *string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 查询的过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询简单过滤条件
	FilterSimple *string `json:"FilterSimple,omitnil,omitempty" name:"FilterSimple"`

	// group by 条件
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// order by 条件
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// limit 参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset 参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeAppSingleCaseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppSingleCaseListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "From")
	delete(f, "Fields")
	delete(f, "Filter")
	delete(f, "FilterSimple")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppSingleCaseListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppSingleCaseListResponseParams struct {
	// 查询数据返回
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppSingleCaseListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppSingleCaseListResponseParams `json:"Response"`
}

func (r *DescribeAppSingleCaseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppSingleCaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationExitReportDetailRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>问题Id</p>
	ClientIdentify *string `json:"ClientIdentify,omitnil,omitempty" name:"ClientIdentify"`

	// <p>开始时间</p>
	StartEventTime *int64 `json:"StartEventTime,omitnil,omitempty" name:"StartEventTime"`

	// <p>结束时间</p>
	EndEventTime *int64 `json:"EndEventTime,omitnil,omitempty" name:"EndEventTime"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeApplicationExitReportDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>问题Id</p>
	ClientIdentify *string `json:"ClientIdentify,omitnil,omitempty" name:"ClientIdentify"`

	// <p>开始时间</p>
	StartEventTime *int64 `json:"StartEventTime,omitnil,omitempty" name:"StartEventTime"`

	// <p>结束时间</p>
	EndEventTime *int64 `json:"EndEventTime,omitnil,omitempty" name:"EndEventTime"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeApplicationExitReportDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationExitReportDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ParamToken")
	delete(f, "ClientIdentify")
	delete(f, "StartEventTime")
	delete(f, "EndEventTime")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationExitReportDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationExitReportDetailResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationExitReportDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationExitReportDetailResponseParams `json:"Response"`
}

func (r *DescribeApplicationExitReportDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationExitReportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationExitReportListRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeApplicationExitReportListRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeApplicationExitReportListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationExitReportListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ParamToken")
	delete(f, "FormListString")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationExitReportListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationExitReportListResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApplicationExitReportListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationExitReportListResponseParams `json:"Response"`
}

func (r *DescribeApplicationExitReportListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationExitReportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataBridgeUrlV2RequestParams struct {
	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// pagepv：性能视图，allcount：性能视图，falls：页面加载瀑布图，samp：首屏时间，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 是否海外
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境变量
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// url名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// http状态码
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitnil,omitempty" name:"Ret"`

	// 网络状态
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeDataBridgeUrlV2Request struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// pagepv：性能视图，allcount：性能视图，falls：页面加载瀑布图，samp：首屏时间，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 是否海外
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境变量
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// url名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// http状态码
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitnil,omitempty" name:"Ret"`

	// 网络状态
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *DescribeDataBridgeUrlV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataBridgeUrlV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "Level")
	delete(f, "Isp")
	delete(f, "Area")
	delete(f, "NetType")
	delete(f, "Platform")
	delete(f, "Device")
	delete(f, "VersionNum")
	delete(f, "ExtFirst")
	delete(f, "ExtSecond")
	delete(f, "ExtThird")
	delete(f, "IsAbroad")
	delete(f, "Browser")
	delete(f, "Os")
	delete(f, "Engine")
	delete(f, "Brand")
	delete(f, "From")
	delete(f, "CostType")
	delete(f, "Env")
	delete(f, "Name")
	delete(f, "Status")
	delete(f, "Ret")
	delete(f, "NetStatus")
	delete(f, "ExtFourth")
	delete(f, "ExtFifth")
	delete(f, "ExtSixth")
	delete(f, "ExtSeventh")
	delete(f, "ExtEighth")
	delete(f, "ExtNinth")
	delete(f, "ExtTenth")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataBridgeUrlV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataBridgeUrlV2ResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataBridgeUrlV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataBridgeUrlV2ResponseParams `json:"Response"`
}

func (r *DescribeDataBridgeUrlV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataBridgeUrlV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataCustomUrlRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top：资源top视图，allcount：性能视图，day：14天数据，condition：条件列表，pagepv：性能视图，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（使用 percentile2 计算分位数）；"avg" 表示均值（使用 avg 计算）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 自定义测速的key的值
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataCustomUrlRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top：资源top视图，allcount：性能视图，day：14天数据，condition：条件列表，pagepv：性能视图，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（使用 percentile2 计算分位数）；"avg" 表示均值（使用 avg 计算）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 自定义测速的key的值
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataCustomUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataCustomUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataCustomUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataCustomUrlResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataCustomUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataCustomUrlResponseParams `json:"Response"`
}

func (r *DescribeDataCustomUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataCustomUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataCustomUrlV2RequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top：资源top视图，allcount：性能视图，day：14天数据，condition：条件列表，pagepv：性能视图，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 自定义测速的key的值
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeDataCustomUrlV2Request struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top：资源top视图，allcount：性能视图，day：14天数据，condition：条件列表，pagepv：性能视图，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 自定义测速的key的值
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *DescribeDataCustomUrlV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataCustomUrlV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	delete(f, "ExtFourth")
	delete(f, "ExtFifth")
	delete(f, "ExtSixth")
	delete(f, "ExtSeventh")
	delete(f, "ExtEighth")
	delete(f, "ExtNinth")
	delete(f, "ExtTenth")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataCustomUrlV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataCustomUrlV2ResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataCustomUrlV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataCustomUrlV2ResponseParams `json:"Response"`
}

func (r *DescribeDataCustomUrlV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataCustomUrlV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEventUrlRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，ckuv：获取uv趋势，ckpv：获取pv趋势，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// netType | 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。也可通过 Type=condition（show tag values）查询当前数据集中的实际可选值。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 筛选条件
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataEventUrlRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，ckuv：获取uv趋势，ckpv：获取pv趋势，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// netType | 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。也可通过 Type=condition（show tag values）查询当前数据集中的实际可选值。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 筛选条件
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataEventUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEventUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Name")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEventUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEventUrlResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataEventUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEventUrlResponseParams `json:"Response"`
}

func (r *DescribeDataEventUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEventUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEventUrlV2RequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，ckuv：获取uv趋势，ckpv：获取pv趋势，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 筛选条件
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeDataEventUrlV2Request struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，ckuv：获取uv趋势，ckpv：获取pv趋势，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 筛选条件
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *DescribeDataEventUrlV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEventUrlV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Name")
	delete(f, "Env")
	delete(f, "ExtFourth")
	delete(f, "ExtFifth")
	delete(f, "ExtSixth")
	delete(f, "ExtSeventh")
	delete(f, "ExtEighth")
	delete(f, "ExtNinth")
	delete(f, "ExtTenth")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEventUrlV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEventUrlV2ResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataEventUrlV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataEventUrlV2ResponseParams `json:"Response"`
}

func (r *DescribeDataEventUrlV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEventUrlV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchProjectRequestParams struct {
	// 开始时间，示例值：1625454840
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间，示例值：1625454840
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级（1表示白名单日志，2表示一般日志，4表示错误日志，8表示Promise 错误，16表示Ajax 请求异常，32表示JS 加载异常，64表示图片加载异常，128表示css 加载异常，256表示console.error，512表示音视频资源异常，1024表示retcode 异常，2048表示aegis report，4096表示PV日志，8192表示自定义事件，16384表示小程序 页面不存在，32768表示websocket错误，65536表示js bridge错误）
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// HTTP 状态码（tag 值）：用于过滤字段 status；取值一般为 200/301/404/500 等，也可配合 4xx/5xx 统计逻辑使用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitnil,omitempty" name:"Ret"`
}

type DescribeDataFetchProjectRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，示例值：1625454840
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间，示例值：1625454840
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级（1表示白名单日志，2表示一般日志，4表示错误日志，8表示Promise 错误，16表示Ajax 请求异常，32表示JS 加载异常，64表示图片加载异常，128表示css 加载异常，256表示console.error，512表示音视频资源异常，1024表示retcode 异常，2048表示aegis report，4096表示PV日志，8192表示自定义事件，16384表示小程序 页面不存在，32768表示websocket错误，65536表示js bridge错误）
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// HTTP 状态码（tag 值）：用于过滤字段 status；取值一般为 200/301/404/500 等，也可配合 4xx/5xx 统计逻辑使用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitnil,omitempty" name:"Ret"`
}

func (r *DescribeDataFetchProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	delete(f, "Status")
	delete(f, "Ret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataFetchProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchProjectResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataFetchProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataFetchProjectResponseParams `json:"Response"`
}

func (r *DescribeDataFetchProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlInfoRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Type	是	String	无枚举值。此接口不使用 Type 做分支判断，SQL 固定 group by "url"，Type 字段传任何值不影响查询。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataFetchUrlInfoRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Type	是	String	无枚举值。此接口不使用 Type 做分支判断，SQL 固定 group by "url"，Type 字段传任何值不影响查询。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataFetchUrlInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataFetchUrlInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlInfoResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataFetchUrlInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataFetchUrlInfoResponseParams `json:"Response"`
}

func (r *DescribeDataFetchUrlInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，pagepv：pv视图，day：14天数据，count40x：40X视图，count50x：50X视图，count5xand4x：40∑50视图，top：资源top视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// HTTP 状态码（tag 值）：用于过滤字段 status；支持 4xx（status =~ /4[0-9]{2}/）和 5xx（status =~ /5[0-9]{2}/）等错误码统计。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitnil,omitempty" name:"Ret"`

	// 网络状态(tag 值):用于过滤/聚合字段 netStatus；枚举值：0(正常)、1(弱网)、2(断网)、3(其他)。
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`
}

type DescribeDataFetchUrlRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，pagepv：pv视图，day：14天数据，count40x：40X视图，count50x：50X视图，count5xand4x：40∑50视图，top：资源top视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// HTTP 状态码（tag 值）：用于过滤字段 status；支持 4xx（status =~ /4[0-9]{2}/）和 5xx（status =~ /5[0-9]{2}/）等错误码统计。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitnil,omitempty" name:"Ret"`

	// 网络状态(tag 值):用于过滤/聚合字段 netStatus；枚举值：0(正常)、1(弱网)、2(断网)、3(其他)。
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`
}

func (r *DescribeDataFetchUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	delete(f, "Status")
	delete(f, "Ret")
	delete(f, "NetStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataFetchUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataFetchUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataFetchUrlResponseParams `json:"Response"`
}

func (r *DescribeDataFetchUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlV2RequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，pagepv：pv视图，day：14天数据，count40x：40X视图，count50x：50X视图，count5xand4x：40∑50视图，top：资源top视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// httpcode响应码
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitnil,omitempty" name:"Ret"`

	// 网络状态
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeDataFetchUrlV2Request struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，pagepv：pv视图，day：14天数据，count40x：40X视图，count50x：50X视图，count5xand4x：40∑50视图，top：资源top视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// httpcode响应码
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// retcode
	Ret *string `json:"Ret,omitnil,omitempty" name:"Ret"`

	// 网络状态
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *DescribeDataFetchUrlV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	delete(f, "Status")
	delete(f, "Ret")
	delete(f, "NetStatus")
	delete(f, "ExtFourth")
	delete(f, "ExtFifth")
	delete(f, "ExtSixth")
	delete(f, "ExtSeventh")
	delete(f, "ExtEighth")
	delete(f, "ExtNinth")
	delete(f, "ExtTenth")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataFetchUrlV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFetchUrlV2ResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataFetchUrlV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataFetchUrlV2ResponseParams `json:"Response"`
}

func (r *DescribeDataFetchUrlV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFetchUrlV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlInfoRequestParams struct {
	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeDataLogUrlInfoRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeDataLogUrlInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataLogUrlInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlInfoResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataLogUrlInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataLogUrlInfoResponseParams `json:"Response"`
}

func (r *DescribeDataLogUrlInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlStatisticsRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// analysis：异常分析，compare：异常列表对比，allcount：性能视图，condition：条件列表，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境区分
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// js异常信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`
}

type DescribeDataLogUrlStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// analysis：异常分析，compare：异常列表对比，allcount：性能视图，condition：条件列表，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境区分
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// js异常信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`
}

func (r *DescribeDataLogUrlStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Env")
	delete(f, "ErrorMsg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataLogUrlStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlStatisticsResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataLogUrlStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataLogUrlStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDataLogUrlStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlStatisticsV2RequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// analysis：异常分析，compare：异常列表对比，allcount：性能视图，condition：条件列表，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境区分
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// js异常信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeDataLogUrlStatisticsV2Request struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// analysis：异常分析，compare：异常列表对比，allcount：性能视图，condition：条件列表，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境区分
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// js异常信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *DescribeDataLogUrlStatisticsV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlStatisticsV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Env")
	delete(f, "ErrorMsg")
	delete(f, "ExtFourth")
	delete(f, "ExtFifth")
	delete(f, "ExtSixth")
	delete(f, "ExtSeventh")
	delete(f, "ExtEighth")
	delete(f, "ExtNinth")
	delete(f, "ExtTenth")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataLogUrlStatisticsV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataLogUrlStatisticsV2ResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataLogUrlStatisticsV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataLogUrlStatisticsV2ResponseParams `json:"Response"`
}

func (r *DescribeDataLogUrlStatisticsV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlStatisticsV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPerformancePageRequestParams struct {
	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// pagepv：pv视图，allcount：性能视图，falls：页面加载瀑布图，samp：首屏时间，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境变量
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 网络状态(tag 值):用于过滤/聚合字段 netStatus；枚举值：0(正常)、1(弱网)、2(断网)、3(其他)。 
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`

	// 是否返回webvitals数据
	WebVitals *bool `json:"WebVitals,omitnil,omitempty" name:"WebVitals"`
}

type DescribeDataPerformancePageRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// pagepv：pv视图，allcount：性能视图，falls：页面加载瀑布图，samp：首屏时间，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境变量
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 网络状态(tag 值):用于过滤/聚合字段 netStatus；枚举值：0(正常)、1(弱网)、2(断网)、3(其他)。 
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`

	// 是否返回webvitals数据
	WebVitals *bool `json:"WebVitals,omitnil,omitempty" name:"WebVitals"`
}

func (r *DescribeDataPerformancePageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPerformancePageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "Level")
	delete(f, "Isp")
	delete(f, "Area")
	delete(f, "NetType")
	delete(f, "Platform")
	delete(f, "Device")
	delete(f, "VersionNum")
	delete(f, "ExtFirst")
	delete(f, "ExtSecond")
	delete(f, "ExtThird")
	delete(f, "IsAbroad")
	delete(f, "Browser")
	delete(f, "Os")
	delete(f, "Engine")
	delete(f, "Brand")
	delete(f, "From")
	delete(f, "CostType")
	delete(f, "Env")
	delete(f, "NetStatus")
	delete(f, "WebVitals")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPerformancePageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPerformancePageResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataPerformancePageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataPerformancePageResponseParams `json:"Response"`
}

func (r *DescribeDataPerformancePageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPerformancePageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPerformancePageV2RequestParams struct {
	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// pagepv：pv视图，allcount：性能视图，falls：页面加载瀑布图，samp：首屏时间，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境变量
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 网络状态
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`

	// 是否返回webvitals数据
	WebVitals *bool `json:"WebVitals,omitnil,omitempty" name:"WebVitals"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeDataPerformancePageV2Request struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// pagepv：pv视图，allcount：性能视图，falls：页面加载瀑布图，samp：首屏时间，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境变量
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 网络状态
	NetStatus *string `json:"NetStatus,omitnil,omitempty" name:"NetStatus"`

	// 是否返回webvitals数据
	WebVitals *bool `json:"WebVitals,omitnil,omitempty" name:"WebVitals"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *DescribeDataPerformancePageV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPerformancePageV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "Level")
	delete(f, "Isp")
	delete(f, "Area")
	delete(f, "NetType")
	delete(f, "Platform")
	delete(f, "Device")
	delete(f, "VersionNum")
	delete(f, "ExtFirst")
	delete(f, "ExtSecond")
	delete(f, "ExtThird")
	delete(f, "IsAbroad")
	delete(f, "Browser")
	delete(f, "Os")
	delete(f, "Engine")
	delete(f, "Brand")
	delete(f, "From")
	delete(f, "CostType")
	delete(f, "Env")
	delete(f, "NetStatus")
	delete(f, "WebVitals")
	delete(f, "ExtFourth")
	delete(f, "ExtFifth")
	delete(f, "ExtSixth")
	delete(f, "ExtSeventh")
	delete(f, "ExtEighth")
	delete(f, "ExtNinth")
	delete(f, "ExtTenth")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPerformancePageV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPerformancePageV2ResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataPerformancePageV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataPerformancePageV2ResponseParams `json:"Response"`
}

func (r *DescribeDataPerformancePageV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPerformancePageV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlInfoRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询类型（string）：当前后端实现固定按 from 聚合统计，未使用该字段（保留字段，传值不会影响结果）。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataPvUrlInfoRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询类型（string）：当前后端实现固定按 from 聚合统计，未使用该字段（保留字段，传值不会影响结果）。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataPvUrlInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPvUrlInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlInfoResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataPvUrlInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataPvUrlInfoResponseParams `json:"Response"`
}

func (r *DescribeDataPvUrlInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlStatisticsRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，vp：性能，ckuv：uv，ckpv：pv，condition：条件列表，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// group by 参数值枚举1:1m  2:5m  3:30m  4:1h 
	//  5:1d
	GroupByType *int64 `json:"GroupByType,omitnil,omitempty" name:"GroupByType"`

	// 1: 查询智研
	// 0: 走旧逻辑，已下线，勿使用
	IsNewData *int64 `json:"IsNewData,omitnil,omitempty" name:"IsNewData"`
}

type DescribeDataPvUrlStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，vp：性能，ckuv：uv，ckpv：pv，condition：条件列表，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// group by 参数值枚举1:1m  2:5m  3:30m  4:1h 
	//  5:1d
	GroupByType *int64 `json:"GroupByType,omitnil,omitempty" name:"GroupByType"`

	// 1: 查询智研
	// 0: 走旧逻辑，已下线，勿使用
	IsNewData *int64 `json:"IsNewData,omitnil,omitempty" name:"IsNewData"`
}

func (r *DescribeDataPvUrlStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Env")
	delete(f, "GroupByType")
	delete(f, "IsNewData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPvUrlStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlStatisticsResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataPvUrlStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataPvUrlStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDataPvUrlStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlStatisticsV2RequestParams struct {
	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>allcount：性能视图，day：14天数据，vp：性能，ckuv：uv，ckpv：pv，condition：条件列表，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>自定义2</p>
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// <p>浏览器引擎</p>
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// <p>运营商</p>
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// <p>来源页面， 可多选，用 “,” 隔开的 String</p>
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// <p>日志等级</p>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>品牌</p>
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// <p>地区</p>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// <p>版本</p>
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>自定义3</p>
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// <p>自定义1</p>
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// <p>网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>机型</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。</p>
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// <p>操作系统</p>
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// <p>浏览器</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>环境</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>group by 参数值枚举1:1m  2:5m  3:30m  4:1h<br> 5:1d</p>
	GroupByType *int64 `json:"GroupByType,omitnil,omitempty" name:"GroupByType"`

	// <p>1: 查询智研<br>0: 走旧逻辑，已下线，勿使用</p>
	IsNewData *int64 `json:"IsNewData,omitnil,omitempty" name:"IsNewData"`

	// <p>自定义4</p>
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// <p>自定义5</p>
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// <p>自定义6</p>
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// <p>自定义7</p>
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// <p>自定义8</p>
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// <p>自定义9</p>
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// <p>自定义10</p>
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// <p>时间段</p>
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeDataPvUrlStatisticsV2Request struct {
	*tchttp.BaseRequest
	
	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>allcount：性能视图，day：14天数据，vp：性能，ckuv：uv，ckpv：pv，condition：条件列表，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>自定义2</p>
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// <p>浏览器引擎</p>
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// <p>运营商</p>
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// <p>来源页面， 可多选，用 “,” 隔开的 String</p>
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// <p>日志等级</p>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>品牌</p>
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// <p>地区</p>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// <p>版本</p>
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>自定义3</p>
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// <p>自定义1</p>
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// <p>网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>机型</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。</p>
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// <p>操作系统</p>
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// <p>浏览器</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>环境</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>group by 参数值枚举1:1m  2:5m  3:30m  4:1h<br> 5:1d</p>
	GroupByType *int64 `json:"GroupByType,omitnil,omitempty" name:"GroupByType"`

	// <p>1: 查询智研<br>0: 走旧逻辑，已下线，勿使用</p>
	IsNewData *int64 `json:"IsNewData,omitnil,omitempty" name:"IsNewData"`

	// <p>自定义4</p>
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// <p>自定义5</p>
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// <p>自定义6</p>
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// <p>自定义7</p>
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// <p>自定义8</p>
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// <p>自定义9</p>
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// <p>自定义10</p>
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// <p>时间段</p>
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *DescribeDataPvUrlStatisticsV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlStatisticsV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Env")
	delete(f, "GroupByType")
	delete(f, "IsNewData")
	delete(f, "ExtFourth")
	delete(f, "ExtFifth")
	delete(f, "ExtSixth")
	delete(f, "ExtSeventh")
	delete(f, "ExtEighth")
	delete(f, "ExtNinth")
	delete(f, "ExtTenth")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPvUrlStatisticsV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataPvUrlStatisticsV2ResponseParams struct {
	// <p>返回值</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataPvUrlStatisticsV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataPvUrlStatisticsV2ResponseParams `json:"Response"`
}

func (r *DescribeDataPvUrlStatisticsV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlStatisticsV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataReportCountRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 上报类型（custom，event，log，miniProgramData，performance，pv，speed，webvitals）
	ReportType *string `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeDataReportCountRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 上报类型（custom，event，log，miniProgramData，performance，pv，speed，webvitals）
	ReportType *string `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

func (r *DescribeDataReportCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataReportCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ReportType")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataReportCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataReportCountResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataReportCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataReportCountResponseParams `json:"Response"`
}

func (r *DescribeDataReportCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataReportCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataReportCountV2RequestParams struct {
	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>上报类型（custom，event，log，miniProgramData，performance，pv，speed，webvitals）</p>
	ReportType *string `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// <p>实例ID</p>
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeDataReportCountV2Request struct {
	*tchttp.BaseRequest
	
	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>上报类型（custom，event，log，miniProgramData，performance，pv，speed，webvitals）</p>
	ReportType *string `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// <p>实例ID</p>
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

func (r *DescribeDataReportCountV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataReportCountV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ReportType")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataReportCountV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataReportCountV2ResponseParams struct {
	// <p>返回值</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataReportCountV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataReportCountV2ResponseParams `json:"Response"`
}

func (r *DescribeDataReportCountV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataReportCountV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataRequestParams struct {
	// 查询字符串
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询字符串
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataResponseParams `json:"Response"`
}

func (r *DescribeDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSetUrlStatisticsRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，data：小程序，component：小程序相关，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时/数据量口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 小程序包类型（tag 值）：用于过滤字段 type（请求参数名为 PackageType）；取值由上报数据决定，可通过 Type=condition（show tag values）获取可选值集合。
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`
}

type DescribeDataSetUrlStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，data：小程序，component：小程序相关，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时/数据量口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 小程序包类型（tag 值）：用于过滤字段 type（请求参数名为 PackageType）；取值由上报数据决定，可通过 Type=condition（show tag values）获取可选值集合。
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`
}

func (r *DescribeDataSetUrlStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSetUrlStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Env")
	delete(f, "PackageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataSetUrlStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSetUrlStatisticsResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataSetUrlStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataSetUrlStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDataSetUrlStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSetUrlStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSetUrlStatisticsV2RequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，data：小程序，component：小程序相关，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 获取package
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeDataSetUrlStatisticsV2Request struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，data：小程序，component：小程序相关，day：14天数据，nettype：网络/平台视图，performance：页面性能TOP视图，version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：ISP视图/地区视图/浏览器视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 获取package
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *DescribeDataSetUrlStatisticsV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSetUrlStatisticsV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Env")
	delete(f, "PackageType")
	delete(f, "ExtFourth")
	delete(f, "ExtFifth")
	delete(f, "ExtSixth")
	delete(f, "ExtSeventh")
	delete(f, "ExtEighth")
	delete(f, "ExtNinth")
	delete(f, "ExtTenth")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataSetUrlStatisticsV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSetUrlStatisticsV2ResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataSetUrlStatisticsV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataSetUrlStatisticsV2ResponseParams `json:"Response"`
}

func (r *DescribeDataSetUrlStatisticsV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSetUrlStatisticsV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticProjectRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级（1表示白名单日志，2表示一般日志，4表示错误日志，8表示Promise 错误，16表示Ajax 请求异常，32表示JS 加载异常，64表示图片加载异常，128表示css 加载异常，256表示console.error，512表示音视频资源异常，1024表示retcode 异常，2048表示aegis report，4096表示PV日志，8192表示自定义事件，16384表示小程序 页面不存在，32768表示websocket错误，65536表示js bridge错误）
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url []*string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataStaticProjectRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级（1表示白名单日志，2表示一般日志，4表示错误日志，8表示Promise 错误，16表示Ajax 请求异常，32表示JS 加载异常，64表示图片加载异常，128表示css 加载异常，256表示console.error，512表示音视频资源异常，1024表示retcode 异常，2048表示aegis report，4096表示PV日志，8192表示自定义事件，16384表示小程序 页面不存在，32768表示websocket错误，65536表示js bridge错误）
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url []*string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataStaticProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataStaticProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticProjectResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataStaticProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataStaticProjectResponseParams `json:"Response"`
}

func (r *DescribeDataStaticProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticProjectV2RequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级（1表示白名单日志，2表示一般日志，4表示错误日志，8表示Promise 错误，16表示Ajax 请求异常，32表示JS 加载异常，64表示图片加载异常，128表示css 加载异常，256表示console.error，512表示音视频资源异常，1024表示retcode 异常，2048表示aegis report，4096表示PV日志，8192表示自定义事件，16384表示小程序 页面不存在，32768表示websocket错误，65536表示js bridge错误）
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url []*string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeDataStaticProjectV2Request struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// allcount：性能视图，day：14天数据，condition：条件列表，area：请求速度分布，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图/ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级（1表示白名单日志，2表示一般日志，4表示错误日志，8表示Promise 错误，16表示Ajax 请求异常，32表示JS 加载异常，64表示图片加载异常，128表示css 加载异常，256表示console.error，512表示音视频资源异常，1024表示retcode 异常，2048表示aegis report，4096表示PV日志，8192表示自定义事件，16384表示小程序 页面不存在，32768表示websocket错误，65536表示js bridge错误）
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型（1,2,3,4,5,100），1表示WIFI, 2表示2G, 3表示3G, 4表示4G, 5表示5G, 6表示6G, 100表示未知。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url []*string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *DescribeDataStaticProjectV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticProjectV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	delete(f, "ExtFourth")
	delete(f, "ExtFifth")
	delete(f, "ExtSixth")
	delete(f, "ExtSeventh")
	delete(f, "ExtEighth")
	delete(f, "ExtNinth")
	delete(f, "ExtTenth")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataStaticProjectV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticProjectV2ResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataStaticProjectV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataStaticProjectV2ResponseParams `json:"Response"`
}

func (r *DescribeDataStaticProjectV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticProjectV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticResourceRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top：资源top视图，count40x：40X视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图//ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataStaticResourceRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top：资源top视图，count40x：40X视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图//ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataStaticResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataStaticResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticResourceResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataStaticResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataStaticResourceResponseParams `json:"Response"`
}

func (r *DescribeDataStaticResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticResourceV2RequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top：资源top视图，count40x：40X视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图//ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeDataStaticResourceV2Request struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top：资源top视图，count40x：40X视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图//ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *DescribeDataStaticResourceV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticResourceV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	delete(f, "ExtFourth")
	delete(f, "ExtFifth")
	delete(f, "ExtSixth")
	delete(f, "ExtSeventh")
	delete(f, "ExtEighth")
	delete(f, "ExtNinth")
	delete(f, "ExtTenth")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataStaticResourceV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticResourceV2ResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataStaticResourceV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataStaticResourceV2ResponseParams `json:"Response"`
}

func (r *DescribeDataStaticResourceV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticResourceV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticUrlRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// pagepv：性能视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图//ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataStaticUrlRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// pagepv：性能视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图//ext1视图等等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataStaticUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataStaticUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticUrlResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataStaticUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataStaticUrlResponseParams `json:"Response"`
}

func (r *DescribeDataStaticUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticUrlV2RequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// pagepv：性能视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图//ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeDataStaticUrlV2Request struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// pagepv：性能视图，nettype/version/platform/isp/region/device/browser/ext1/ext2/ext3/ret/status/from/url/env/：网络平台视图/Version视图/设备视图/ISP视图/地区视图/浏览器视图//ext1视图等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 来源
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 自定义4
	ExtFourth *string `json:"ExtFourth,omitnil,omitempty" name:"ExtFourth"`

	// 自定义5
	ExtFifth *string `json:"ExtFifth,omitnil,omitempty" name:"ExtFifth"`

	// 自定义6
	ExtSixth *string `json:"ExtSixth,omitnil,omitempty" name:"ExtSixth"`

	// 自定义7
	ExtSeventh *string `json:"ExtSeventh,omitnil,omitempty" name:"ExtSeventh"`

	// 自定义8
	ExtEighth *string `json:"ExtEighth,omitnil,omitempty" name:"ExtEighth"`

	// 自定义9
	ExtNinth *string `json:"ExtNinth,omitnil,omitempty" name:"ExtNinth"`

	// 自定义10
	ExtTenth *string `json:"ExtTenth,omitnil,omitempty" name:"ExtTenth"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *DescribeDataStaticUrlV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticUrlV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Url")
	delete(f, "Env")
	delete(f, "ExtFourth")
	delete(f, "ExtFifth")
	delete(f, "ExtSixth")
	delete(f, "ExtSeventh")
	delete(f, "ExtEighth")
	delete(f, "ExtNinth")
	delete(f, "ExtTenth")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataStaticUrlV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataStaticUrlV2ResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataStaticUrlV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataStaticUrlV2ResponseParams `json:"Response"`
}

func (r *DescribeDataStaticUrlV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataStaticUrlV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataWebVitalsPageRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 查询维度：from=按页面来源(from)聚合输出；其他值/空值=输出整体汇总。
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 按页面来源分组（group by from），返回每个页面的 LCP/FID/CLS/FCP
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。 
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// Web Vitals 口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。 
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDataWebVitalsPageRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 查询维度：from=按页面来源(from)聚合输出；其他值/空值=输出整体汇总。
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 按页面来源分组（group by from），返回每个页面的 LCP/FID/CLS/FCP
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型(tag 值):用于过滤/聚合字段 netType；枚举值：1(WiFi)、2(2G)、3(3G)、4(4G)、5(5G)、6(6G)、100(未知网络)。 
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// Web Vitals 口径："50"/"75"/"90"/"95"/"99"/"99.5" 分别表示 TP50/TP75/TP90/TP95/TP99/TP99.5（percentile2）；"avg" 表示均值（avg）。 
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeDataWebVitalsPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataWebVitalsPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Type")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataWebVitalsPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataWebVitalsPageResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataWebVitalsPageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataWebVitalsPageResponseParams `json:"Response"`
}

func (r *DescribeDataWebVitalsPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataWebVitalsPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataWebVitalsPageV2RequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 类型暂无
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type DescribeDataWebVitalsPageV2Request struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitnil,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 类型暂无
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitnil,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitnil,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// 显示是否海外,1表示海外，0表示非海外；默认值为空，查询所有。
	IsAbroad *string `json:"IsAbroad,omitnil,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// 耗时计算
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 环境
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 时间段
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *DescribeDataWebVitalsPageV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataWebVitalsPageV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Type")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "CostType")
	delete(f, "Env")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataWebVitalsPageV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataWebVitalsPageV2ResponseParams struct {
	// 返回值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataWebVitalsPageV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeDataWebVitalsPageV2ResponseParams `json:"Response"`
}

func (r *DescribeDataWebVitalsPageV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataWebVitalsPageV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeErrorRequestParams struct {
	// 日期
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeErrorRequest struct {
	*tchttp.BaseRequest
	
	// 日期
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeErrorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Date")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeErrorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeErrorResponseParams struct {
	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeErrorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeErrorResponseParams `json:"Response"`
}

func (r *DescribeErrorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExceptionDetailRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>消息唯一标识</p>
	ClientIdentify *string `json:"ClientIdentify,omitnil,omitempty" name:"ClientIdentify"`

	// <p>集群堆栈类型</p>
	ClusterStackType *int64 `json:"ClusterStackType,omitnil,omitempty" name:"ClusterStackType"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>问题类型</p>
	IssueType *int64 `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// <p>事件开始时间</p>
	StartEventTime *int64 `json:"StartEventTime,omitnil,omitempty" name:"StartEventTime"`

	// <p>事件结束时间</p>
	EndEventTime *int64 `json:"EndEventTime,omitnil,omitempty" name:"EndEventTime"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeExceptionDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>消息唯一标识</p>
	ClientIdentify *string `json:"ClientIdentify,omitnil,omitempty" name:"ClientIdentify"`

	// <p>集群堆栈类型</p>
	ClusterStackType *int64 `json:"ClusterStackType,omitnil,omitempty" name:"ClusterStackType"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>问题类型</p>
	IssueType *int64 `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// <p>事件开始时间</p>
	StartEventTime *int64 `json:"StartEventTime,omitnil,omitempty" name:"StartEventTime"`

	// <p>事件结束时间</p>
	EndEventTime *int64 `json:"EndEventTime,omitnil,omitempty" name:"EndEventTime"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeExceptionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExceptionDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ClientIdentify")
	delete(f, "ClusterStackType")
	delete(f, "Feature")
	delete(f, "IssueType")
	delete(f, "StartEventTime")
	delete(f, "EndEventTime")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExceptionDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExceptionDetailResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExceptionDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExceptionDetailResponseParams `json:"Response"`
}

func (r *DescribeExceptionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExceptionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExceptionReportListRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 form_list 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>问题类型</p>
	IssueType *int64 `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序方式</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>每页数目</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>拓展字段</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeExceptionReportListRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 form_list 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>问题类型</p>
	IssueType *int64 `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序方式</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>每页数目</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>拓展字段</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeExceptionReportListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExceptionReportListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "FormListString")
	delete(f, "ParamToken")
	delete(f, "IssueType")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "Feature")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExceptionReportListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExceptionReportListResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExceptionReportListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExceptionReportListResponseParams `json:"Response"`
}

func (r *DescribeExceptionReportListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExceptionReportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFOOMMallocProblemDetailRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>消息唯一标识</p>
	ClientIdentify *string `json:"ClientIdentify,omitnil,omitempty" name:"ClientIdentify"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>事件开始时间</p>
	StartEventTime *int64 `json:"StartEventTime,omitnil,omitempty" name:"StartEventTime"`

	// <p>事件结束时间</p>
	EndEventTime *int64 `json:"EndEventTime,omitnil,omitempty" name:"EndEventTime"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeFOOMMallocProblemDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>消息唯一标识</p>
	ClientIdentify *string `json:"ClientIdentify,omitnil,omitempty" name:"ClientIdentify"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>事件开始时间</p>
	StartEventTime *int64 `json:"StartEventTime,omitnil,omitempty" name:"StartEventTime"`

	// <p>事件结束时间</p>
	EndEventTime *int64 `json:"EndEventTime,omitnil,omitempty" name:"EndEventTime"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeFOOMMallocProblemDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFOOMMallocProblemDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ClientIdentify")
	delete(f, "Feature")
	delete(f, "StartEventTime")
	delete(f, "EndEventTime")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFOOMMallocProblemDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFOOMMallocProblemDetailResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFOOMMallocProblemDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFOOMMallocProblemDetailResponseParams `json:"Response"`
}

func (r *DescribeFOOMMallocProblemDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFOOMMallocProblemDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFOOMMallocProblemListRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeFOOMMallocProblemListRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeFOOMMallocProblemListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFOOMMallocProblemListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ParamToken")
	delete(f, "FormListString")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFOOMMallocProblemListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFOOMMallocProblemListResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFOOMMallocProblemListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFOOMMallocProblemListResponseParams `json:"Response"`
}

func (r *DescribeFOOMMallocProblemListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFOOMMallocProblemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFOOMMallocReportListRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeFOOMMallocReportListRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeFOOMMallocReportListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFOOMMallocReportListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "FormListString")
	delete(f, "ParamToken")
	delete(f, "Feature")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFOOMMallocReportListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFOOMMallocReportListResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFOOMMallocReportListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFOOMMallocReportListResponseParams `json:"Response"`
}

func (r *DescribeFOOMMallocReportListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFOOMMallocReportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFOOMProblemDetailRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>消息唯一标识</p>
	ClientIdentify *string `json:"ClientIdentify,omitnil,omitempty" name:"ClientIdentify"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>事件开始时间</p>
	StartEventTime *int64 `json:"StartEventTime,omitnil,omitempty" name:"StartEventTime"`

	// <p>事件结束时间</p>
	EndEventTime *int64 `json:"EndEventTime,omitnil,omitempty" name:"EndEventTime"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeFOOMProblemDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>消息唯一标识</p>
	ClientIdentify *string `json:"ClientIdentify,omitnil,omitempty" name:"ClientIdentify"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>事件开始时间</p>
	StartEventTime *int64 `json:"StartEventTime,omitnil,omitempty" name:"StartEventTime"`

	// <p>事件结束时间</p>
	EndEventTime *int64 `json:"EndEventTime,omitnil,omitempty" name:"EndEventTime"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeFOOMProblemDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFOOMProblemDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ClientIdentify")
	delete(f, "Feature")
	delete(f, "StartEventTime")
	delete(f, "EndEventTime")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFOOMProblemDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFOOMProblemDetailResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFOOMProblemDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFOOMProblemDetailResponseParams `json:"Response"`
}

func (r *DescribeFOOMProblemDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFOOMProblemDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFOOMProblemListRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeFOOMProblemListRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeFOOMProblemListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFOOMProblemListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ParamToken")
	delete(f, "FormListString")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFOOMProblemListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFOOMProblemListResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFOOMProblemListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFOOMProblemListResponseParams `json:"Response"`
}

func (r *DescribeFOOMProblemListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFOOMProblemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFOOMReportListRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeFOOMReportListRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeFOOMReportListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFOOMReportListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "FormListString")
	delete(f, "ParamToken")
	delete(f, "Feature")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFOOMReportListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFOOMReportListResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFOOMReportListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFOOMReportListResponseParams `json:"Response"`
}

func (r *DescribeFOOMReportListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFOOMReportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFieldTransformRulesRequestParams struct {
	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>偏移</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeFieldTransformRulesRequest struct {
	*tchttp.BaseRequest
	
	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>偏移</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeFieldTransformRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFieldTransformRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFieldTransformRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFieldTransformRulesResponseParams struct {
	// <p>结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFieldTransformRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFieldTransformRulesResponseParams `json:"Response"`
}

func (r *DescribeFieldTransformRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFieldTransformRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFieldValuesRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间</p><p>单位：毫秒</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p><p>单位：毫秒</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>字段名</p>
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// <p>数量限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>事件分类</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>可选。按前缀过滤字段去重结果。查询 user_id 或 device_id 且开启 StrictPiiPrefix 时必填，去除首尾空格后长度须不少于 2。</p>
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// <p>用于采样配置，是否对 user_id、device_id 强制前缀匹配。为 true 时必须同时传入合法 Prefix，否则返回 InvalidParameter，且不会返回无前缀的全量去重列表。为 false 或不传时，按字段返回热门去重值，行为与历史接口一致。</p><p>默认值：false</p>
	StrictPiiPrefix *bool `json:"StrictPiiPrefix,omitnil,omitempty" name:"StrictPiiPrefix"`

	// <p>查询口径。可选。只允许空或 event。Issue 标签分布传 event。为 event 时不能同时传 HasError、MinDuration、MaxDuration，也不能把 Category 设为 session。</p>
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>是否只看有错误的会话。可选。不传表示不过滤。Scope 为 event 时不能传。</p>
	HasError *bool `json:"HasError,omitnil,omitempty" name:"HasError"`

	// <p>会话最短时长，单位毫秒。可选。0 或不传表示不限制。Scope 为 event 时不能传。</p>
	MinDuration *int64 `json:"MinDuration,omitnil,omitempty" name:"MinDuration"`

	// <p>会话最长时长，单位毫秒。可选。0 或不传表示不限制。Scope 为 event 时不能传。</p>
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`
}

type DescribeFieldValuesRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间</p><p>单位：毫秒</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p><p>单位：毫秒</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>字段名</p>
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// <p>数量限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>事件分类</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>可选。按前缀过滤字段去重结果。查询 user_id 或 device_id 且开启 StrictPiiPrefix 时必填，去除首尾空格后长度须不少于 2。</p>
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// <p>用于采样配置，是否对 user_id、device_id 强制前缀匹配。为 true 时必须同时传入合法 Prefix，否则返回 InvalidParameter，且不会返回无前缀的全量去重列表。为 false 或不传时，按字段返回热门去重值，行为与历史接口一致。</p><p>默认值：false</p>
	StrictPiiPrefix *bool `json:"StrictPiiPrefix,omitnil,omitempty" name:"StrictPiiPrefix"`

	// <p>查询口径。可选。只允许空或 event。Issue 标签分布传 event。为 event 时不能同时传 HasError、MinDuration、MaxDuration，也不能把 Category 设为 session。</p>
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>是否只看有错误的会话。可选。不传表示不过滤。Scope 为 event 时不能传。</p>
	HasError *bool `json:"HasError,omitnil,omitempty" name:"HasError"`

	// <p>会话最短时长，单位毫秒。可选。0 或不传表示不限制。Scope 为 event 时不能传。</p>
	MinDuration *int64 `json:"MinDuration,omitnil,omitempty" name:"MinDuration"`

	// <p>会话最长时长，单位毫秒。可选。0 或不传表示不限制。Scope 为 event 时不能传。</p>
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`
}

func (r *DescribeFieldValuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFieldValuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "FieldName")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Category")
	delete(f, "Prefix")
	delete(f, "StrictPiiPrefix")
	delete(f, "Scope")
	delete(f, "HasError")
	delete(f, "MinDuration")
	delete(f, "MaxDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFieldValuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFieldValuesResponseParams struct {
	// <p>返回结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFieldValuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFieldValuesResponseParams `json:"Response"`
}

func (r *DescribeFieldValuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFieldValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileCertificateRequestParams struct {
	// 超时时间，不填默认是 5 分钟
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// bucket类型，不填默认1:web，2:app
	FileType *int64 `json:"FileType,omitnil,omitempty" name:"FileType"`

	// ProjectID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// reigon对应的id
	RegionID *int64 `json:"RegionID,omitnil,omitempty" name:"RegionID"`
}

type DescribeFileCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 超时时间，不填默认是 5 分钟
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// bucket类型，不填默认1:web，2:app
	FileType *int64 `json:"FileType,omitnil,omitempty" name:"FileType"`

	// ProjectID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// reigon对应的id
	RegionID *int64 `json:"RegionID,omitnil,omitempty" name:"RegionID"`
}

func (r *DescribeFileCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Timeout")
	delete(f, "FileType")
	delete(f, "ID")
	delete(f, "RegionID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileCertificateResponseParams struct {
	// 临时密钥key
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 临时密钥 id
	SecretID *string `json:"SecretID,omitnil,omitempty" name:"SecretID"`

	// 临时密钥临时 token
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// 开始时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 过期时间戳
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 桶地址
	BucketAddress *string `json:"BucketAddress,omitnil,omitempty" name:"BucketAddress"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileCertificateResponseParams `json:"Response"`
}

func (r *DescribeFileCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueAffectedUsersRequestParams struct {
	// <p>应用 ID（string，必填）。空 → InvalidParameter</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>问题 ID（string，必填）。空 → InvalidParameter</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>页码（整数，选填）。&lt; 1 钳制为 1；缺省 1</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>每页条数（整数，选填）。&lt; 1 或 &gt; 1000 回落默认 20；导出场景按 1000 一批拉全量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>应用版本（string，选填）。按 app_version 精确匹配；空表示不限版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>起始毫秒时间戳（整数，选填）。&lt;= 0 表示该侧不限</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束毫秒时间戳（整数，选填）。&lt;= 0 表示该侧不限；与 StartTime 同时给出且 StartTime &gt; EndTime → InvalidParameter</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>身份值关键字（string，选填）。不区分大小写的子串匹配，命中 if(user_id='', session_id, user_id)；空表示不限</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>排序字段（string，选填）。可选 last_seen / first_seen / event_count；缺省 last_seen；其它值 → InvalidParameter</p>
	OrderKey *string `json:"OrderKey,omitnil,omitempty" name:"OrderKey"`

	// <p>排序方向（string，选填）。可选 asc / desc；缺省 desc；其它值 → InvalidParameter</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

type DescribeIssueAffectedUsersRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（string，必填）。空 → InvalidParameter</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>问题 ID（string，必填）。空 → InvalidParameter</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>页码（整数，选填）。&lt; 1 钳制为 1；缺省 1</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>每页条数（整数，选填）。&lt; 1 或 &gt; 1000 回落默认 20；导出场景按 1000 一批拉全量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>应用版本（string，选填）。按 app_version 精确匹配；空表示不限版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>起始毫秒时间戳（整数，选填）。&lt;= 0 表示该侧不限</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束毫秒时间戳（整数，选填）。&lt;= 0 表示该侧不限；与 StartTime 同时给出且 StartTime &gt; EndTime → InvalidParameter</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>身份值关键字（string，选填）。不区分大小写的子串匹配，命中 if(user_id='', session_id, user_id)；空表示不限</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>排序字段（string，选填）。可选 last_seen / first_seen / event_count；缺省 last_seen；其它值 → InvalidParameter</p>
	OrderKey *string `json:"OrderKey,omitnil,omitempty" name:"OrderKey"`

	// <p>排序方向（string，选填）。可选 asc / desc；缺省 desc；其它值 → InvalidParameter</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

func (r *DescribeIssueAffectedUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueAffectedUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "IssueId")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "AppVersion")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Keyword")
	delete(f, "OrderKey")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIssueAffectedUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueAffectedUsersResponseParams struct {
	// <p>业务数据 JSON 字符串，反序列化后：{Users,Total,Page,PageSize}。 Users[] 含 UserId（身份值）、HasUin（是否来自 user_id）、EventCount、FirstSeen / LastSeen（毫秒时间戳）。</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息（成功时为空）</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIssueAffectedUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIssueAffectedUsersResponseParams `json:"Response"`
}

func (r *DescribeIssueAffectedUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueAffectedUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueBreadcrumbsRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Issue ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>事件 ID；为空时默认回退到最新事件</p>
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// <p>向前回溯秒数，默认 60，最大 300</p>
	Before *int64 `json:"Before,omitnil,omitempty" name:"Before"`

	// <p>返回条数上限，默认 50，最大 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeIssueBreadcrumbsRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Issue ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>事件 ID；为空时默认回退到最新事件</p>
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// <p>向前回溯秒数，默认 60，最大 300</p>
	Before *int64 `json:"Before,omitnil,omitempty" name:"Before"`

	// <p>返回条数上限，默认 50，最大 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeIssueBreadcrumbsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueBreadcrumbsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "IssueId")
	delete(f, "EventId")
	delete(f, "Before")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIssueBreadcrumbsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueBreadcrumbsResponseParams struct {
	// <p>业务结果 JSON 字符串；前端会再做一次 JSON 解析。Result 内业务 JSON 为 Data，内容是 session_id + items</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息；正常返回时为空</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIssueBreadcrumbsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIssueBreadcrumbsResponseParams `json:"Response"`
}

func (r *DescribeIssueBreadcrumbsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueBreadcrumbsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueCrashLogDownloadURLRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Issue ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>Issue 事件ID</p>
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`
}

type DescribeIssueCrashLogDownloadURLRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Issue ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>Issue 事件ID</p>
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`
}

func (r *DescribeIssueCrashLogDownloadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueCrashLogDownloadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "IssueId")
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIssueCrashLogDownloadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueCrashLogDownloadURLResponseParams struct {
	// <p>请求结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误时返回的错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIssueCrashLogDownloadURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIssueCrashLogDownloadURLResponseParams `json:"Response"`
}

func (r *DescribeIssueCrashLogDownloadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueCrashLogDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueDetailRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Issue ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>开始时间，Unix 毫秒时间戳。可选。与 EndTime 必须成对传入；都不传时按全保留期统计。</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒时间戳。可选。与 StartTime 必须成对传入。</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>环境。可选。为空表示全部环境。</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。可选。不传表示不附加这组筛选。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeIssueDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Issue ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>开始时间，Unix 毫秒时间戳。可选。与 EndTime 必须成对传入；都不传时按全保留期统计。</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒时间戳。可选。与 StartTime 必须成对传入。</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>环境。可选。为空表示全部环境。</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。可选。不传表示不附加这组筛选。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeIssueDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "IssueId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Env")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIssueDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueDetailResponseParams struct {
	// <p>业务结果 JSON 字符串；前端会再做一次 JSON 解析。Result 内业务 JSON 为 Data，内容是单个 Issue 详情</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息；正常返回时为空</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIssueDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIssueDetailResponseParams `json:"Response"`
}

func (r *DescribeIssueDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueEventDetailRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Issue ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>事件 ID</p>
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`
}

type DescribeIssueEventDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Issue ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>事件 ID</p>
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`
}

func (r *DescribeIssueEventDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueEventDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "IssueId")
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIssueEventDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueEventDetailResponseParams struct {
	// <p>业务结果 JSON 字符串；前端会再做一次 JSON 解析。Result 内业务 JSON 为 Data，内容是单个事件详情</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息；正常返回时为空</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIssueEventDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIssueEventDetailResponseParams `json:"Response"`
}

func (r *DescribeIssueEventDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueEventsRequestParams struct {
	// 应用 ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Issue ID
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// 开始时间，Unix 毫秒时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，Unix 毫秒时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 事件关键字过滤
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 扩展过滤器 JSON 字符串；白名单字段包括 app_version、country、province、city、device_type、browser、os、isp、referrer、user_type、platform、normalized_page_url、event_type、issue_id、environment、process_type
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 页码，默认 1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 分页大小，默认 20，范围 1~100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序方向，可选值：asc、desc
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeIssueEventsRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Issue ID
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// 开始时间，Unix 毫秒时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，Unix 毫秒时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 事件关键字过滤
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 扩展过滤器 JSON 字符串；白名单字段包括 app_version、country、province、city、device_type、browser、os、isp、referrer、user_type、platform、normalized_page_url、event_type、issue_id、environment、process_type
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 页码，默认 1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 分页大小，默认 20，范围 1~100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序方向，可选值：asc、desc
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *DescribeIssueEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "IssueId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Keyword")
	delete(f, "Filters")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIssueEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueEventsResponseParams struct {
	// 业务结果 JSON 字符串；前端会再做一次 JSON 解析。Result 内业务 JSON 为 Data，内容是 total + items
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 错误信息；正常返回时为空
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIssueEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIssueEventsResponseParams `json:"Response"`
}

func (r *DescribeIssueEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueTrendRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Issue ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>开始时间，Unix 毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒时间戳</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>趋势分桶，可选值：auto、1m、10m、1h、1d</p>
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// <p>环境。可选。为空表示全部环境。</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>关键字。可选。不传表示不过滤。</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。可选。不传表示不附加这组筛选。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeIssueTrendRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Issue ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>开始时间，Unix 毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒时间戳</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>趋势分桶，可选值：auto、1m、10m、1h、1d</p>
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// <p>环境。可选。为空表示全部环境。</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>关键字。可选。不传表示不过滤。</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。可选。不传表示不附加这组筛选。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeIssueTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "IssueId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Bucket")
	delete(f, "Env")
	delete(f, "Keyword")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIssueTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssueTrendResponseParams struct {
	// <p>业务结果 JSON 字符串；前端会再做一次 JSON 解析。Result 内业务 JSON 为 Data，内容是 bucket + points</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息；正常返回时为空</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIssueTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIssueTrendResponseParams `json:"Response"`
}

func (r *DescribeIssueTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssueTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssuesDistributionRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>分布维度是自定义维度时，填‘user_custom’</p>
	DimType *string `json:"DimType,omitnil,omitempty" name:"DimType"`

	// <p>维度，e.g. os_version, app_version, model等</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>数字类型字段的区间范围</p>
	Intervals []*int64 `json:"Intervals,omitnil,omitempty" name:"Intervals"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 form_list 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>问题Id</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>问题类型</p>
	IssueType *int64 `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// <p>限制返回的个数，默认返回所有值</p>
	ParamLimit *int64 `json:"ParamLimit,omitnil,omitempty" name:"ParamLimit"`

	// <p>键</p>
	MapKey *string `json:"MapKey,omitnil,omitempty" name:"MapKey"`

	// <p>名称</p>
	MapName *string `json:"MapName,omitnil,omitempty" name:"MapName"`

	// <p>指标类型</p>
	MetricType *int64 `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// <p>每页数目</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>用户自定义维度key</p>
	UserCustomKey *string `json:"UserCustomKey,omitnil,omitempty" name:"UserCustomKey"`

	// <p>拓展字段</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeIssuesDistributionRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>分布维度是自定义维度时，填‘user_custom’</p>
	DimType *string `json:"DimType,omitnil,omitempty" name:"DimType"`

	// <p>维度，e.g. os_version, app_version, model等</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>数字类型字段的区间范围</p>
	Intervals []*int64 `json:"Intervals,omitnil,omitempty" name:"Intervals"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 form_list 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>问题Id</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>问题类型</p>
	IssueType *int64 `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// <p>限制返回的个数，默认返回所有值</p>
	ParamLimit *int64 `json:"ParamLimit,omitnil,omitempty" name:"ParamLimit"`

	// <p>键</p>
	MapKey *string `json:"MapKey,omitnil,omitempty" name:"MapKey"`

	// <p>名称</p>
	MapName *string `json:"MapName,omitnil,omitempty" name:"MapName"`

	// <p>指标类型</p>
	MetricType *int64 `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// <p>每页数目</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>用户自定义维度key</p>
	UserCustomKey *string `json:"UserCustomKey,omitnil,omitempty" name:"UserCustomKey"`

	// <p>拓展字段</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeIssuesDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssuesDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "FormListString")
	delete(f, "DimType")
	delete(f, "Dimension")
	delete(f, "Intervals")
	delete(f, "ParamToken")
	delete(f, "IssueId")
	delete(f, "IssueType")
	delete(f, "ParamLimit")
	delete(f, "MapKey")
	delete(f, "MapName")
	delete(f, "MetricType")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "UserCustomKey")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIssuesDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssuesDistributionResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIssuesDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIssuesDistributionResponseParams `json:"Response"`
}

func (r *DescribeIssuesDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssuesDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssuesListRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormList *string `json:"FormList,omitnil,omitempty" name:"FormList"`

	// <p>接口调试专用，对比模式下条件A，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListA *string `json:"FormListA,omitnil,omitempty" name:"FormListA"`

	// <p>接口调试专用，对比模式下条件B，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListB *string `json:"FormListB,omitnil,omitempty" name:"FormListB"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 form_list 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>问题类型</p>
	IssueType *int64 `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序方式</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>每页数目</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>问题对比列表模式下，用于标识是按照sort_field字段的A值排序还是B值还是ratio值</p>
	SortABRatio *string `json:"SortABRatio,omitnil,omitempty" name:"SortABRatio"`

	// <p>模式：false:问题列表模式，true:对比列表模式</p>
	Compare *bool `json:"Compare,omitnil,omitempty" name:"Compare"`

	// <p>对比状态 0:所有 1:新增 2：遗留 3:已解决</p>
	CompareStatus *int64 `json:"CompareStatus,omitnil,omitempty" name:"CompareStatus"`

	// <p>拓展字段</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeIssuesListRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormList *string `json:"FormList,omitnil,omitempty" name:"FormList"`

	// <p>接口调试专用，对比模式下条件A，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListA *string `json:"FormListA,omitnil,omitempty" name:"FormListA"`

	// <p>接口调试专用，对比模式下条件B，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListB *string `json:"FormListB,omitnil,omitempty" name:"FormListB"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 form_list 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>问题类型</p>
	IssueType *int64 `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序方式</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>每页数目</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>问题对比列表模式下，用于标识是按照sort_field字段的A值排序还是B值还是ratio值</p>
	SortABRatio *string `json:"SortABRatio,omitnil,omitempty" name:"SortABRatio"`

	// <p>模式：false:问题列表模式，true:对比列表模式</p>
	Compare *bool `json:"Compare,omitnil,omitempty" name:"Compare"`

	// <p>对比状态 0:所有 1:新增 2：遗留 3:已解决</p>
	CompareStatus *int64 `json:"CompareStatus,omitnil,omitempty" name:"CompareStatus"`

	// <p>拓展字段</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeIssuesListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssuesListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "FormList")
	delete(f, "FormListA")
	delete(f, "FormListB")
	delete(f, "ParamToken")
	delete(f, "IssueType")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "SortABRatio")
	delete(f, "Compare")
	delete(f, "CompareStatus")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIssuesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssuesListResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIssuesListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIssuesListResponseParams `json:"Response"`
}

func (r *DescribeIssuesListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssuesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssuesRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，Unix 毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒时间戳</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>环境过滤，如 prod、test</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>Issue 类型过滤，可选值：native_crash、js_error、event_loop_block</p>
	IssueTypes []*string `json:"IssueTypes,omitnil,omitempty" name:"IssueTypes"`

	// <p>状态过滤，可选值：Unresolved、Resolved、Ignored、Regressed</p>
	Statuses []*string `json:"Statuses,omitnil,omitempty" name:"Statuses"`

	// <p>Issue 标题或错误关键字模糊搜索</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>扩展过滤器 JSON 字符串；白名单字段包括 app_version、country、province、city、device_type、browser、os、isp、referrer、user_type、platform、normalized_page_url、event_type、issue_id、environment、process_type</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>排序字段，可选值：last_seen、first_seen、event_count、affected_users</p>
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// <p>排序方向，可选值：asc、desc</p>
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// <p>页码，默认 1</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>分页大小，默认 20，范围 1~100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeIssuesRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，Unix 毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒时间戳</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>环境过滤，如 prod、test</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>Issue 类型过滤，可选值：native_crash、js_error、event_loop_block</p>
	IssueTypes []*string `json:"IssueTypes,omitnil,omitempty" name:"IssueTypes"`

	// <p>状态过滤，可选值：Unresolved、Resolved、Ignored、Regressed</p>
	Statuses []*string `json:"Statuses,omitnil,omitempty" name:"Statuses"`

	// <p>Issue 标题或错误关键字模糊搜索</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>扩展过滤器 JSON 字符串；白名单字段包括 app_version、country、province、city、device_type、browser、os、isp、referrer、user_type、platform、normalized_page_url、event_type、issue_id、environment、process_type</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>排序字段，可选值：last_seen、first_seen、event_count、affected_users</p>
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// <p>排序方向，可选值：asc、desc</p>
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// <p>页码，默认 1</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>分页大小，默认 20，范围 1~100</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeIssuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Env")
	delete(f, "IssueTypes")
	delete(f, "Statuses")
	delete(f, "Keyword")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "SortOrder")
	delete(f, "Page")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIssuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssuesResponseParams struct {
	// <p>业务结果 JSON 字符串；前端会再做一次 JSON 解析。Result 内业务 JSON 为 Total + Items</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息；正常返回时为空</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIssuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIssuesResponseParams `json:"Response"`
}

func (r *DescribeIssuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssuesStatisticsTrendRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 提供给前端使用，当填写本字段时，会覆盖 form_list 的值
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// 接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息
	FormList *string `json:"FormList,omitnil,omitempty" name:"FormList"`

	// 问题Id
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// 问题类型
	IssueType *int64 `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// 时间窗口
	TimeWindow *int64 `json:"TimeWindow,omitnil,omitempty" name:"TimeWindow"`

	// 累计值
	TotalSize *bool `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 无
	Stat *int64 `json:"Stat,omitnil,omitempty" name:"Stat"`

	// 指标类型
	MetricType *int64 `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// 拓展字段
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// 请求头
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`

	// 无
	TrendStat *int64 `json:"TrendStat,omitnil,omitempty" name:"TrendStat"`
}

type DescribeIssuesStatisticsTrendRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 提供给前端使用，当填写本字段时，会覆盖 form_list 的值
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// 接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息
	FormList *string `json:"FormList,omitnil,omitempty" name:"FormList"`

	// 问题Id
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// 问题类型
	IssueType *int64 `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// 时间窗口
	TimeWindow *int64 `json:"TimeWindow,omitnil,omitempty" name:"TimeWindow"`

	// 累计值
	TotalSize *bool `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 无
	Stat *int64 `json:"Stat,omitnil,omitempty" name:"Stat"`

	// 指标类型
	MetricType *int64 `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// 拓展字段
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// 请求头
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`

	// 无
	TrendStat *int64 `json:"TrendStat,omitnil,omitempty" name:"TrendStat"`
}

func (r *DescribeIssuesStatisticsTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssuesStatisticsTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ParamToken")
	delete(f, "FormList")
	delete(f, "IssueId")
	delete(f, "IssueType")
	delete(f, "TimeWindow")
	delete(f, "TotalSize")
	delete(f, "Stat")
	delete(f, "MetricType")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	delete(f, "TrendStat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIssuesStatisticsTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIssuesStatisticsTrendResponseParams struct {
	// 返回值
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 状态码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 消息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIssuesStatisticsTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIssuesStatisticsTrendResponseParams `json:"Response"`
}

func (r *DescribeIssuesStatisticsTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIssuesStatisticsTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLagANRProblemAccountDetailRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>消息唯一标识</p>
	ClientIdentify *string `json:"ClientIdentify,omitnil,omitempty" name:"ClientIdentify"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>事件开始时间</p>
	StartEventTime *int64 `json:"StartEventTime,omitnil,omitempty" name:"StartEventTime"`

	// <p>事件结束时间</p>
	EndEventTime *int64 `json:"EndEventTime,omitnil,omitempty" name:"EndEventTime"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeLagANRProblemAccountDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>消息唯一标识</p>
	ClientIdentify *string `json:"ClientIdentify,omitnil,omitempty" name:"ClientIdentify"`

	// <p>特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>事件开始时间</p>
	StartEventTime *int64 `json:"StartEventTime,omitnil,omitempty" name:"StartEventTime"`

	// <p>事件结束时间</p>
	EndEventTime *int64 `json:"EndEventTime,omitnil,omitempty" name:"EndEventTime"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeLagANRProblemAccountDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLagANRProblemAccountDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ClientIdentify")
	delete(f, "Feature")
	delete(f, "StartEventTime")
	delete(f, "EndEventTime")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLagANRProblemAccountDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLagANRProblemAccountDetailResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLagANRProblemAccountDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLagANRProblemAccountDetailResponseParams `json:"Response"`
}

func (r *DescribeLagANRProblemAccountDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLagANRProblemAccountDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLagANRProblemFeatureAccountsRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>问题特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeLagANRProblemFeatureAccountsRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>问题特征</p>
	Feature *string `json:"Feature,omitnil,omitempty" name:"Feature"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeLagANRProblemFeatureAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLagANRProblemFeatureAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "FormListString")
	delete(f, "ParamToken")
	delete(f, "Feature")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLagANRProblemFeatureAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLagANRProblemFeatureAccountsResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLagANRProblemFeatureAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLagANRProblemFeatureAccountsResponseParams `json:"Response"`
}

func (r *DescribeLagANRProblemFeatureAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLagANRProblemFeatureAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLagANRProblemListRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeLagANRProblemListRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>提供给前端使用，当填写本字段时，会覆盖 formlist 的值</p>
	ParamToken *string `json:"ParamToken,omitnil,omitempty" name:"ParamToken"`

	// <p>接口调试专用，当 token 为空时，以这里的 value 作为筛选表单信息</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>当前页码</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页展示最大数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序类型</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>拓展数据</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeLagANRProblemListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLagANRProblemListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ParamToken")
	delete(f, "FormListString")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLagANRProblemListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLagANRProblemListResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLagANRProblemListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLagANRProblemListResponseParams `json:"Response"`
}

func (r *DescribeLagANRProblemListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLagANRProblemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMiniProgramReplayPackRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>回放包hash（sha256前16位hex）</p>
	BundleHash *string `json:"BundleHash,omitnil,omitempty" name:"BundleHash"`
}

type DescribeMiniProgramReplayPackRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>回放包hash（sha256前16位hex）</p>
	BundleHash *string `json:"BundleHash,omitnil,omitempty" name:"BundleHash"`
}

func (r *DescribeMiniProgramReplayPackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMiniProgramReplayPackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "BundleHash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMiniProgramReplayPackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMiniProgramReplayPackResponseParams struct {
	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// <p>请求结果（JSON字符串，包含 PackId/AppId/BundleHash/Status/PageCount/ComponentCount/Pack）</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMiniProgramReplayPackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMiniProgramReplayPackResponseParams `json:"Response"`
}

func (r *DescribeMiniProgramReplayPackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMiniProgramReplayPackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMiniProgramReplayPacksRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>状态过滤（compiling/ready/failed，可选）</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>返回条数（默认20，最大100）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页游标（可选）</p>
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

type DescribeMiniProgramReplayPacksRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>状态过滤（compiling/ready/failed，可选）</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>返回条数（默认20，最大100）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页游标（可选）</p>
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

func (r *DescribeMiniProgramReplayPacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMiniProgramReplayPacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Status")
	delete(f, "Limit")
	delete(f, "Cursor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMiniProgramReplayPacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMiniProgramReplayPacksResponseParams struct {
	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// <p>请求结果（JSON字符串，包含 Items/NextCursor）</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMiniProgramReplayPacksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMiniProgramReplayPacksResponseParams `json:"Response"`
}

func (r *DescribeMiniProgramReplayPacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMiniProgramReplayPacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePerformanceMetricsRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>过滤条件</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>分位值</p>
	Quantile *float64 `json:"Quantile,omitnil,omitempty" name:"Quantile"`

	// <p>旅程筛选条件 JSON 字符串，格式与 Filters 相同：[{Key,Op,Values}]。可选。不传表示不过滤。</p>
	JourneyFilters *string `json:"JourneyFilters,omitnil,omitempty" name:"JourneyFilters"`

	// <p>是否只看包含挣扎的会话。可选。不传表示不过滤。</p>
	HasStruggle *bool `json:"HasStruggle,omitnil,omitempty" name:"HasStruggle"`

	// <p>旅程起点页面。可选。不传表示不过滤。</p>
	StartPage *string `json:"StartPage,omitnil,omitempty" name:"StartPage"`

	// <p>旅程终点页面。可选。不传表示不过滤。</p>
	EndPage *string `json:"EndPage,omitnil,omitempty" name:"EndPage"`
}

type DescribePerformanceMetricsRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>过滤条件</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>分位值</p>
	Quantile *float64 `json:"Quantile,omitnil,omitempty" name:"Quantile"`

	// <p>旅程筛选条件 JSON 字符串，格式与 Filters 相同：[{Key,Op,Values}]。可选。不传表示不过滤。</p>
	JourneyFilters *string `json:"JourneyFilters,omitnil,omitempty" name:"JourneyFilters"`

	// <p>是否只看包含挣扎的会话。可选。不传表示不过滤。</p>
	HasStruggle *bool `json:"HasStruggle,omitnil,omitempty" name:"HasStruggle"`

	// <p>旅程起点页面。可选。不传表示不过滤。</p>
	StartPage *string `json:"StartPage,omitnil,omitempty" name:"StartPage"`

	// <p>旅程终点页面。可选。不传表示不过滤。</p>
	EndPage *string `json:"EndPage,omitnil,omitempty" name:"EndPage"`
}

func (r *DescribePerformanceMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePerformanceMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Platform")
	delete(f, "Filters")
	delete(f, "Quantile")
	delete(f, "JourneyFilters")
	delete(f, "HasStruggle")
	delete(f, "StartPage")
	delete(f, "EndPage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePerformanceMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePerformanceMetricsResponseParams struct {
	// <p>请求结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePerformanceMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePerformanceMetricsResponseParams `json:"Response"`
}

func (r *DescribePerformanceMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePerformanceMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectLimitsRequestParams struct {
	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`
}

type DescribeProjectLimitsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`
}

func (r *DescribeProjectLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectLimitsResponseParams struct {
	// 上报率数组列表
	ProjectLimitSet []*ProjectLimit `json:"ProjectLimitSet,omitnil,omitempty" name:"ProjectLimitSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectLimitsResponseParams `json:"Response"`
}

func (r *DescribeProjectLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsRequestParams struct {
	// 分页每页数目，整型
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页页码，整型
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤参数；demo模式传{"Name": "IsDemo", "Values":["1"]}
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 该参数已废弃，demo模式请在Filters内注明
	//
	// Deprecated: IsDemo is deprecated.
	IsDemo *int64 `json:"IsDemo,omitnil,omitempty" name:"IsDemo"`
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest
	
	// 分页每页数目，整型
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页页码，整型
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤参数；demo模式传{"Name": "IsDemo", "Values":["1"]}
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 该参数已废弃，demo模式请在Filters内注明
	IsDemo *int64 `json:"IsDemo,omitnil,omitempty" name:"IsDemo"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "IsDemo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsResponseParams struct {
	// 列表总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 项目列表
	ProjectSet []*RumProject `json:"ProjectSet,omitnil,omitempty" name:"ProjectSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectsResponseParams `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePvListRequestParams struct {
	// ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 对PV指标的查询维度。获取day：d，   获取min则不填。
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`
}

type DescribePvListRequest struct {
	*tchttp.BaseRequest
	
	// ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 对PV指标的查询维度。获取day：d，   获取min则不填。
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`
}

func (r *DescribePvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "Dimension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePvListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePvListResponseParams struct {
	// pv列表
	ProjectPvSet []*RumPvInfo `json:"ProjectPvSet,omitnil,omitempty" name:"ProjectPvSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePvListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePvListResponseParams `json:"Response"`
}

func (r *DescribePvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseFileSignRequestParams struct {
	// 超时时间，不填默认是 5 分钟
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// bucket类型，不填默认1:web，2:app
	FileType *int64 `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 获取临时签名的bucket是境内还是境外（1表示境外，其它表示境内）
	Site *int64 `json:"Site,omitnil,omitempty" name:"Site"`

	// ProjectID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeReleaseFileSignRequest struct {
	*tchttp.BaseRequest
	
	// 超时时间，不填默认是 5 分钟
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// bucket类型，不填默认1:web，2:app
	FileType *int64 `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 获取临时签名的bucket是境内还是境外（1表示境外，其它表示境内）
	Site *int64 `json:"Site,omitnil,omitempty" name:"Site"`

	// ProjectID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeReleaseFileSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseFileSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Timeout")
	delete(f, "FileType")
	delete(f, "Site")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseFileSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseFileSignResponseParams struct {
	// 临时密钥key
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 临时密钥 id
	SecretID *string `json:"SecretID,omitnil,omitempty" name:"SecretID"`

	// 临时密钥临时 token
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// 开始时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 过期时间戳
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseFileSignResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseFileSignResponseParams `json:"Response"`
}

func (r *DescribeReleaseFileSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseFileSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseFilesRequestParams struct {
	// <p>项目 id</p>
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// <p>文件版本</p>
	FileVersion *string `json:"FileVersion,omitnil,omitempty" name:"FileVersion"`

	// <p>查询过滤条件（根据sourcemap的文件名模糊匹配）</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>false/不传=保留「最近 3 个月」约束（旧行为）；true=绕过时间窗口</p>
	IgnoreDefaultTimeRange *bool `json:"IgnoreDefaultTimeRange,omitnil,omitempty" name:"IgnoreDefaultTimeRange"`
}

type DescribeReleaseFilesRequest struct {
	*tchttp.BaseRequest
	
	// <p>项目 id</p>
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// <p>文件版本</p>
	FileVersion *string `json:"FileVersion,omitnil,omitempty" name:"FileVersion"`

	// <p>查询过滤条件（根据sourcemap的文件名模糊匹配）</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>false/不传=保留「最近 3 个月」约束（旧行为）；true=绕过时间窗口</p>
	IgnoreDefaultTimeRange *bool `json:"IgnoreDefaultTimeRange,omitnil,omitempty" name:"IgnoreDefaultTimeRange"`
}

func (r *DescribeReleaseFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "FileVersion")
	delete(f, "FileName")
	delete(f, "IgnoreDefaultTimeRange")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseFilesResponseParams struct {
	// <p>文件信息列表</p>
	Files []*ReleaseFile `json:"Files,omitnil,omitempty" name:"Files"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseFilesResponseParams `json:"Response"`
}

func (r *DescribeReleaseFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplayDataRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>回放id</p>
	ReplayId *string `json:"ReplayId,omitnil,omitempty" name:"ReplayId"`

	// <p>session id</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>segment id</p>
	SegmentId *int64 `json:"SegmentId,omitnil,omitempty" name:"SegmentId"`

	// <p>单个录制分片解压后的最大允许字节数</p><p>取值范围：[1, 8388608]</p><p>单位：字节（Byte）</p>
	MaxDecodedBytes *int64 `json:"MaxDecodedBytes,omitnil,omitempty" name:"MaxDecodedBytes"`
}

type DescribeReplayDataRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>回放id</p>
	ReplayId *string `json:"ReplayId,omitnil,omitempty" name:"ReplayId"`

	// <p>session id</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>segment id</p>
	SegmentId *int64 `json:"SegmentId,omitnil,omitempty" name:"SegmentId"`

	// <p>单个录制分片解压后的最大允许字节数</p><p>取值范围：[1, 8388608]</p><p>单位：字节（Byte）</p>
	MaxDecodedBytes *int64 `json:"MaxDecodedBytes,omitnil,omitempty" name:"MaxDecodedBytes"`
}

func (r *DescribeReplayDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplayDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "ReplayId")
	delete(f, "SessionId")
	delete(f, "SegmentId")
	delete(f, "MaxDecodedBytes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplayDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplayDataResponseParams struct {
	// <p>返回结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReplayDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplayDataResponseParams `json:"Response"`
}

func (r *DescribeReplayDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplayDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplayDetailsRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>会话id</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>回放id</p>
	ReplayId *string `json:"ReplayId,omitnil,omitempty" name:"ReplayId"`

	// <p>事件类别， 默认为all</p>
	Categories *string `json:"Categories,omitnil,omitempty" name:"Categories"`

	// <p>每页返回数量</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>游标分页，传上一次返回的 NextCursor 值。首次请求不传</p>
	Cursor *int64 `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// <p>起始偏移（毫秒），相对于录屏开始时间。用于只查某个时间段的事件</p>
	StartOffsetMs *int64 `json:"StartOffsetMs,omitnil,omitempty" name:"StartOffsetMs"`

	// <p>结束偏移（毫秒），必须 ≥ StartOffsetMs</p>
	EndOffsetMs *int64 `json:"EndOffsetMs,omitnil,omitempty" name:"EndOffsetMs"`
}

type DescribeReplayDetailsRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>会话id</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>回放id</p>
	ReplayId *string `json:"ReplayId,omitnil,omitempty" name:"ReplayId"`

	// <p>事件类别， 默认为all</p>
	Categories *string `json:"Categories,omitnil,omitempty" name:"Categories"`

	// <p>每页返回数量</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>游标分页，传上一次返回的 NextCursor 值。首次请求不传</p>
	Cursor *int64 `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// <p>起始偏移（毫秒），相对于录屏开始时间。用于只查某个时间段的事件</p>
	StartOffsetMs *int64 `json:"StartOffsetMs,omitnil,omitempty" name:"StartOffsetMs"`

	// <p>结束偏移（毫秒），必须 ≥ StartOffsetMs</p>
	EndOffsetMs *int64 `json:"EndOffsetMs,omitnil,omitempty" name:"EndOffsetMs"`
}

func (r *DescribeReplayDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplayDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "SessionId")
	delete(f, "ReplayId")
	delete(f, "Categories")
	delete(f, "Limit")
	delete(f, "Cursor")
	delete(f, "StartOffsetMs")
	delete(f, "EndOffsetMs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplayDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplayDetailsResponseParams struct {
	// <p>返回结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReplayDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplayDetailsResponseParams `json:"Response"`
}

func (r *DescribeReplayDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplayDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplayMetaRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>回放id</p>
	ReplayId *string `json:"ReplayId,omitnil,omitempty" name:"ReplayId"`

	// <p>会话id</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>是否仅查询指定录制分片的轻量索引。默认false。为true时必须通过IndexSegmentIds指定1～16个分片ID，返回IndexEntries，不返回录制正文。</p>
	IndexOnly *bool `json:"IndexOnly,omitnil,omitempty" name:"IndexOnly"`

	// <p>待查询索引的录制分片ID列表。仅在IndexOnly为true时使用，包含1～16个非负且不重复的分片ID。ProbeMissingIndex为true时仅允许指定一个分片ID。</p>
	IndexSegmentIds []*int64 `json:"IndexSegmentIds,omitnil,omitempty" name:"IndexSegmentIds"`

	// <p>是否在录制分片缺少索引时进行受限兼容探测。默认false。为true时必须同时设置IndexOnly为true，且IndexSegmentIds仅包含一个分片ID。</p>
	ProbeMissingIndex *bool `json:"ProbeMissingIndex,omitnil,omitempty" name:"ProbeMissingIndex"`
}

type DescribeReplayMetaRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>回放id</p>
	ReplayId *string `json:"ReplayId,omitnil,omitempty" name:"ReplayId"`

	// <p>会话id</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>是否仅查询指定录制分片的轻量索引。默认false。为true时必须通过IndexSegmentIds指定1～16个分片ID，返回IndexEntries，不返回录制正文。</p>
	IndexOnly *bool `json:"IndexOnly,omitnil,omitempty" name:"IndexOnly"`

	// <p>待查询索引的录制分片ID列表。仅在IndexOnly为true时使用，包含1～16个非负且不重复的分片ID。ProbeMissingIndex为true时仅允许指定一个分片ID。</p>
	IndexSegmentIds []*int64 `json:"IndexSegmentIds,omitnil,omitempty" name:"IndexSegmentIds"`

	// <p>是否在录制分片缺少索引时进行受限兼容探测。默认false。为true时必须同时设置IndexOnly为true，且IndexSegmentIds仅包含一个分片ID。</p>
	ProbeMissingIndex *bool `json:"ProbeMissingIndex,omitnil,omitempty" name:"ProbeMissingIndex"`
}

func (r *DescribeReplayMetaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplayMetaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "ReplayId")
	delete(f, "SessionId")
	delete(f, "IndexOnly")
	delete(f, "IndexSegmentIds")
	delete(f, "ProbeMissingIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplayMetaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplayMetaResponseParams struct {
	// <p>返回结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReplayMetaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplayMetaResponseParams `json:"Response"`
}

func (r *DescribeReplayMetaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplayMetaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumGroupLogRequestParams struct {
	// <p>排序方式  desc  asc（必填）</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>开始时间（必填）</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>单次查询返回的原始日志条数，最大值为100（必填）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>页数，第几页（注意：该字段已废弃！）</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>查询语句，参考控制台请求参数，语句长度最大为4096（必填）</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>结束时间（必填）</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID（必填）</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>聚合字段</p>
	GroupField *string `json:"GroupField,omitnil,omitempty" name:"GroupField"`
}

type DescribeRumGroupLogRequest struct {
	*tchttp.BaseRequest
	
	// <p>排序方式  desc  asc（必填）</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>开始时间（必填）</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>单次查询返回的原始日志条数，最大值为100（必填）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>页数，第几页（注意：该字段已废弃！）</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>查询语句，参考控制台请求参数，语句长度最大为4096（必填）</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>结束时间（必填）</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID（必填）</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>聚合字段</p>
	GroupField *string `json:"GroupField,omitnil,omitempty" name:"GroupField"`
}

func (r *DescribeRumGroupLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumGroupLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderBy")
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Page")
	delete(f, "Query")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "GroupField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumGroupLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumGroupLogResponseParams struct {
	// <p>返回字符串</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumGroupLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumGroupLogResponseParams `json:"Response"`
}

func (r *DescribeRumGroupLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumGroupLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumGroupLogV2RequestParams struct {
	// <p>排序方式 desc asc（必填）</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>查询日志开始时间（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>单次查询返回的原始日志条数，最大值为100（必填）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件，其中key与运算符(Operator)对应可选字段枚举值与示例查看文档：<a href="https://cloud.tencent.com/document/product/248/87223">日志查询相关</a></p>
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>查询日志结束时间（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID（必填）</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>聚合字段</p>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// <p>页数，第几页（注意：该字段已废弃！）</p>
	Last *int64 `json:"Last,omitnil,omitempty" name:"Last"`
}

type DescribeRumGroupLogV2Request struct {
	*tchttp.BaseRequest
	
	// <p>排序方式 desc asc（必填）</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>查询日志开始时间（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>单次查询返回的原始日志条数，最大值为100（必填）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件，其中key与运算符(Operator)对应可选字段枚举值与示例查看文档：<a href="https://cloud.tencent.com/document/product/248/87223">日志查询相关</a></p>
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>查询日志结束时间（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID（必填）</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>聚合字段</p>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// <p>页数，第几页（注意：该字段已废弃！）</p>
	Last *int64 `json:"Last,omitnil,omitempty" name:"Last"`
}

func (r *DescribeRumGroupLogV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumGroupLogV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderBy")
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Filter")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "Label")
	delete(f, "Last")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumGroupLogV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumGroupLogV2ResponseParams struct {
	// <p>日志聚合结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumGroupLogV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeRumGroupLogV2ResponseParams `json:"Response"`
}

func (r *DescribeRumGroupLogV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumGroupLogV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogDetailsV2RequestParams struct {
	// <p>排序方式 desc asc</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>查询日志开始时间（必填）格式为时间戳 毫秒</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>单次查询返回的原始日志条数，最大值为100（必填）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件，其中key与运算符(Operator)对应可选字段枚举值与示例查看文档：<a href="https://cloud.tencent.com/document/product/248/87223">日志查询相关</a></p>
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>查询日志结束时间（必填）格式为时间戳 毫秒</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID（必填）</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>上次查询的最后一个日志的时间戳</p>
	LastTime *int64 `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// <p>上次查询的最后一个日志的rowId</p>
	LastRowId *int64 `json:"LastRowId,omitnil,omitempty" name:"LastRowId"`
}

type DescribeRumLogDetailsV2Request struct {
	*tchttp.BaseRequest
	
	// <p>排序方式 desc asc</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>查询日志开始时间（必填）格式为时间戳 毫秒</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>单次查询返回的原始日志条数，最大值为100（必填）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件，其中key与运算符(Operator)对应可选字段枚举值与示例查看文档：<a href="https://cloud.tencent.com/document/product/248/87223">日志查询相关</a></p>
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>查询日志结束时间（必填）格式为时间戳 毫秒</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID（必填）</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>上次查询的最后一个日志的时间戳</p>
	LastTime *int64 `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// <p>上次查询的最后一个日志的rowId</p>
	LastRowId *int64 `json:"LastRowId,omitnil,omitempty" name:"LastRowId"`
}

func (r *DescribeRumLogDetailsV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogDetailsV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderBy")
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Filter")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "LastTime")
	delete(f, "LastRowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumLogDetailsV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogDetailsV2ResponseParams struct {
	// <p>日志明细</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumLogDetailsV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeRumLogDetailsV2ResponseParams `json:"Response"`
}

func (r *DescribeRumLogDetailsV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogDetailsV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogExportRequestParams struct {
	// 导出标识name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 开始时间（必填）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// field条件
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`
}

type DescribeRumLogExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出标识name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 开始时间（必填）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// field条件
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`
}

func (r *DescribeRumLogExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "StartTime")
	delete(f, "Query")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "Fields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumLogExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogExportResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumLogExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumLogExportResponseParams `json:"Response"`
}

func (r *DescribeRumLogExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogExportV2RequestParams struct {
	// <p>Name（string，必填）：导出任务名称标识，需要在使用API创建导出任务时填写。此字段为后续能力扩展预留，不在查询结果中返回。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>查询日志开始时间(必填)</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>过滤条件，其中key与运算符(Operator)对应可选字段枚举值与示例查看文档：<a href="https://cloud.tencent.com/document/product/248/87223">日志查询相关</a></p>
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>查询日志结束时间(必填)</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>注意：这个字段必选，不能为空！选择要导出哪些字段的参数，对应日志的列，可查看文档：<a href="https://cloud.tencent.com/document/product/248/87223">日志查询相关</a></p><p>枚举值：</p><ul><li>level： 日志类型</li><li>url： URL</li><li>ip： 用户终端IP</li><li>version： 版本</li></ul>
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`
}

type DescribeRumLogExportV2Request struct {
	*tchttp.BaseRequest
	
	// <p>Name（string，必填）：导出任务名称标识，需要在使用API创建导出任务时填写。此字段为后续能力扩展预留，不在查询结果中返回。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>查询日志开始时间(必填)</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>过滤条件，其中key与运算符(Operator)对应可选字段枚举值与示例查看文档：<a href="https://cloud.tencent.com/document/product/248/87223">日志查询相关</a></p>
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>查询日志结束时间(必填)</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>注意：这个字段必选，不能为空！选择要导出哪些字段的参数，对应日志的列，可查看文档：<a href="https://cloud.tencent.com/document/product/248/87223">日志查询相关</a></p><p>枚举值：</p><ul><li>level： 日志类型</li><li>url： URL</li><li>ip： 用户终端IP</li><li>version： 版本</li></ul>
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`
}

func (r *DescribeRumLogExportV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogExportV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "StartTime")
	delete(f, "Filter")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "Fields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumLogExportV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogExportV2ResponseParams struct {
	// <p>日志导出任务TaskId</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumLogExportV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeRumLogExportV2ResponseParams `json:"Response"`
}

func (r *DescribeRumLogExportV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogExportV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogExportsRequestParams struct {
	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数，第几页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeRumLogExportsRequest struct {
	*tchttp.BaseRequest
	
	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数，第几页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeRumLogExportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogExportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumLogExportsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogExportsResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumLogExportsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumLogExportsResponseParams `json:"Response"`
}

func (r *DescribeRumLogExportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogExportsV2RequestParams struct {
	// <p>每页返回任务数</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>当前页码</p>
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// <p>项目ID</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeRumLogExportsV2Request struct {
	*tchttp.BaseRequest
	
	// <p>每页返回任务数</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>当前页码</p>
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// <p>项目ID</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeRumLogExportsV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogExportsV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumLogExportsV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogExportsV2ResponseParams struct {
	// <p>日志导出列表</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumLogExportsV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeRumLogExportsV2ResponseParams `json:"Response"`
}

func (r *DescribeRumLogExportsV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogExportsV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogListRequestParams struct {
	// 排序方式  desc  asc（必填）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 开始时间（必填）格式为时间戳 毫秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 单次查询返回的原始日志条数，最大值为100（必填）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页数，第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）格式为时间戳 毫秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeRumLogListRequest struct {
	*tchttp.BaseRequest
	
	// 排序方式  desc  asc（必填）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 开始时间（必填）格式为时间戳 毫秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 单次查询返回的原始日志条数，最大值为100（必填）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页数，第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）格式为时间戳 毫秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeRumLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderBy")
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Page")
	delete(f, "Query")
	delete(f, "EndTime")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogListResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumLogListResponseParams `json:"Response"`
}

func (r *DescribeRumLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogTotalV2RequestParams struct {
	// <p>排序方式 desc asc</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>查询日志开始时间（必填）格式为时间戳 毫秒</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>单次查询返回的原始日志条数，最大值为100（必填）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件，其中key与运算符(Operator)对应可选字段枚举值与示例查看文档：<a href="https://cloud.tencent.com/document/product/248/87223">日志查询相关</a></p>
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>查询日志结束时间（必填）格式为时间戳 毫秒</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID（必填）</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>上次查询的最后一个日志的时间戳</p>
	LastTime *int64 `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// <p>上次查询的最后一个日志的rowId</p>
	LastRowId *int64 `json:"LastRowId,omitnil,omitempty" name:"LastRowId"`
}

type DescribeRumLogTotalV2Request struct {
	*tchttp.BaseRequest
	
	// <p>排序方式 desc asc</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>查询日志开始时间（必填）格式为时间戳 毫秒</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>单次查询返回的原始日志条数，最大值为100（必填）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件，其中key与运算符(Operator)对应可选字段枚举值与示例查看文档：<a href="https://cloud.tencent.com/document/product/248/87223">日志查询相关</a></p>
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>查询日志结束时间（必填）格式为时间戳 毫秒</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID（必填）</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>上次查询的最后一个日志的时间戳</p>
	LastTime *int64 `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// <p>上次查询的最后一个日志的rowId</p>
	LastRowId *int64 `json:"LastRowId,omitnil,omitempty" name:"LastRowId"`
}

func (r *DescribeRumLogTotalV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogTotalV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderBy")
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Filter")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "LastTime")
	delete(f, "LastRowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumLogTotalV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumLogTotalV2ResponseParams struct {
	// <p>日志总量</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumLogTotalV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeRumLogTotalV2ResponseParams `json:"Response"`
}

func (r *DescribeRumLogTotalV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumLogTotalV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumStatsLogListRequestParams struct {
	// 开始时间（必填）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 单次查询返回的原始日志条数，最大值为100（必填）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeRumStatsLogListRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间（必填）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 单次查询返回的原始日志条数，最大值为100（必填）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询语句，参考控制台请求参数，语句长度最大为4096（必填）
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 结束时间（必填）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目ID（必填）
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeRumStatsLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumStatsLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Query")
	delete(f, "EndTime")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumStatsLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumStatsLogListResponseParams struct {
	// 返回字符串
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumStatsLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRumStatsLogListResponseParams `json:"Response"`
}

func (r *DescribeRumStatsLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumStatsLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumStatsLogListV2RequestParams struct {
	// <p>查询日志开始时间（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>单次查询返回的原始日志条数，最大值为100（必填）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件，其中key与运算符(Operator)对应可选字段枚举值与示例查看文档：<a href="https://cloud.tencent.com/document/product/248/87223">日志查询相关</a></p>
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>查询日志结束时间（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID（必填）</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DescribeRumStatsLogListV2Request struct {
	*tchttp.BaseRequest
	
	// <p>查询日志开始时间（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>单次查询返回的原始日志条数，最大值为100（必填）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件，其中key与运算符(Operator)对应可选字段枚举值与示例查看文档：<a href="https://cloud.tencent.com/document/product/248/87223">日志查询相关</a></p>
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>查询日志结束时间（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目ID（必填）</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DescribeRumStatsLogListV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumStatsLogListV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Filter")
	delete(f, "EndTime")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRumStatsLogListV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRumStatsLogListV2ResponseParams struct {
	// <p>日志时间分布</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRumStatsLogListV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeRumStatsLogListV2ResponseParams `json:"Response"`
}

func (r *DescribeRumStatsLogListV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRumStatsLogListV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSampleRateRequestParams struct {
	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeSampleRateRequest struct {
	*tchttp.BaseRequest
	
	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeSampleRateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSampleRateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSampleRateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSampleRateResponseParams struct {
	// <p>结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSampleRateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSampleRateResponseParams `json:"Response"`
}

func (r *DescribeSampleRateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSampleRateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScoresRequestParams struct {
	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 该参数已废弃
	//
	// Deprecated: IsDemo is deprecated.
	IsDemo *int64 `json:"IsDemo,omitnil,omitempty" name:"IsDemo"`

	// 项目 ID 列表
	IDList []*int64 `json:"IDList,omitnil,omitempty" name:"IDList"`
}

type DescribeScoresRequest struct {
	*tchttp.BaseRequest
	
	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 项目ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 该参数已废弃
	IsDemo *int64 `json:"IsDemo,omitnil,omitempty" name:"IsDemo"`

	// 项目 ID 列表
	IDList []*int64 `json:"IDList,omitnil,omitempty" name:"IDList"`
}

func (r *DescribeScoresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScoresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "ID")
	delete(f, "IsDemo")
	delete(f, "IDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScoresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScoresResponseParams struct {
	// 数组
	ScoreSet []*ScoreInfo `json:"ScoreSet,omitnil,omitempty" name:"ScoreSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScoresResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScoresResponseParams `json:"Response"`
}

func (r *DescribeScoresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScoresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScoresV2RequestParams struct {
	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目 ID 列表</p>
	IDList []*int64 `json:"IDList,omitnil,omitempty" name:"IDList"`

	// <p>查询粒度，hour 或 day</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>环境</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeScoresV2Request struct {
	*tchttp.BaseRequest
	
	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>项目 ID 列表</p>
	IDList []*int64 `json:"IDList,omitnil,omitempty" name:"IDList"`

	// <p>查询粒度，hour 或 day</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>环境</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeScoresV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScoresV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "IDList")
	delete(f, "Type")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScoresV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScoresV2ResponseParams struct {
	// <p>项目得分数组</p>
	ScoreSet []*ScoreInfoV2 `json:"ScoreSet,omitnil,omitempty" name:"ScoreSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScoresV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeScoresV2ResponseParams `json:"Response"`
}

func (r *DescribeScoresV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScoresV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionListRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间</p><p>单位：毫秒</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p><p>单位：毫秒</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>session类型</p><p>枚举值：</p><ul><li>all： 所有</li><li>error： 错</li><li>slow： 慢</li></ul>
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// <p>页面号</p><p>默认值：1</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>每一页的数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// <p>排序方式，ASC或DESC</p>
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// <p>用户id</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>session id</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>设备类型</p>
	DeviceTypes []*string `json:"DeviceTypes,omitnil,omitempty" name:"DeviceTypes"`

	// <p>浏览器</p>
	Browsers []*string `json:"Browsers,omitnil,omitempty" name:"Browsers"`

	// <p>操作系统</p>
	OSList []*string `json:"OSList,omitnil,omitempty" name:"OSList"`

	// <p>平台</p>
	Platforms []*string `json:"Platforms,omitnil,omitempty" name:"Platforms"`

	// <p>应用版本</p>
	AppVersions []*string `json:"AppVersions,omitnil,omitempty" name:"AppVersions"`

	// <p>国家</p>
	Countries []*string `json:"Countries,omitnil,omitempty" name:"Countries"`

	// <p>是否有挣扎事件</p>
	HasStruggle *bool `json:"HasStruggle,omitnil,omitempty" name:"HasStruggle"`

	// <p>挣扎等级</p>
	StruggleRanks []*string `json:"StruggleRanks,omitnil,omitempty" name:"StruggleRanks"`

	// <p>挣扎事件类型</p>
	StruggleTypes []*string `json:"StruggleTypes,omitnil,omitempty" name:"StruggleTypes"`

	// <p>是否有错误</p>
	HasError *bool `json:"HasError,omitnil,omitempty" name:"HasError"`

	// <p>访问页面</p>
	VisitedPage *string `json:"VisitedPage,omitnil,omitempty" name:"VisitedPage"`

	// <p>包含事件</p>
	ContainsEvent *string `json:"ContainsEvent,omitnil,omitempty" name:"ContainsEvent"`

	// <p>最小持续时间</p>
	MinDuration *int64 `json:"MinDuration,omitnil,omitempty" name:"MinDuration"`

	// <p>最大持续时间</p>
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// <p>性能过滤</p>
	PerfFilter *string `json:"PerfFilter,omitnil,omitempty" name:"PerfFilter"`

	// <p>地区</p>
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// <p>环境</p>
	Environments []*string `json:"Environments,omitnil,omitempty" name:"Environments"`

	// <p>筛选</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>旅程起点页面。可选。不传表示不过滤。</p>
	StartPage *string `json:"StartPage,omitnil,omitempty" name:"StartPage"`

	// <p>旅程终点页面。可选。不传表示不过滤。</p>
	EndPage *string `json:"EndPage,omitnil,omitempty" name:"EndPage"`

	// <p>旅程筛选条件 JSON 字符串，格式与 Filters 相同：[{Key,Op,Values}]。可选。不传表示不过滤。</p>
	JourneyFilters *string `json:"JourneyFilters,omitnil,omitempty" name:"JourneyFilters"`
}

type DescribeSessionListRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间</p><p>单位：毫秒</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p><p>单位：毫秒</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>session类型</p><p>枚举值：</p><ul><li>all： 所有</li><li>error： 错</li><li>slow： 慢</li></ul>
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// <p>页面号</p><p>默认值：1</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>每一页的数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序字段</p>
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// <p>排序方式，ASC或DESC</p>
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// <p>用户id</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>session id</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>设备类型</p>
	DeviceTypes []*string `json:"DeviceTypes,omitnil,omitempty" name:"DeviceTypes"`

	// <p>浏览器</p>
	Browsers []*string `json:"Browsers,omitnil,omitempty" name:"Browsers"`

	// <p>操作系统</p>
	OSList []*string `json:"OSList,omitnil,omitempty" name:"OSList"`

	// <p>平台</p>
	Platforms []*string `json:"Platforms,omitnil,omitempty" name:"Platforms"`

	// <p>应用版本</p>
	AppVersions []*string `json:"AppVersions,omitnil,omitempty" name:"AppVersions"`

	// <p>国家</p>
	Countries []*string `json:"Countries,omitnil,omitempty" name:"Countries"`

	// <p>是否有挣扎事件</p>
	HasStruggle *bool `json:"HasStruggle,omitnil,omitempty" name:"HasStruggle"`

	// <p>挣扎等级</p>
	StruggleRanks []*string `json:"StruggleRanks,omitnil,omitempty" name:"StruggleRanks"`

	// <p>挣扎事件类型</p>
	StruggleTypes []*string `json:"StruggleTypes,omitnil,omitempty" name:"StruggleTypes"`

	// <p>是否有错误</p>
	HasError *bool `json:"HasError,omitnil,omitempty" name:"HasError"`

	// <p>访问页面</p>
	VisitedPage *string `json:"VisitedPage,omitnil,omitempty" name:"VisitedPage"`

	// <p>包含事件</p>
	ContainsEvent *string `json:"ContainsEvent,omitnil,omitempty" name:"ContainsEvent"`

	// <p>最小持续时间</p>
	MinDuration *int64 `json:"MinDuration,omitnil,omitempty" name:"MinDuration"`

	// <p>最大持续时间</p>
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// <p>性能过滤</p>
	PerfFilter *string `json:"PerfFilter,omitnil,omitempty" name:"PerfFilter"`

	// <p>地区</p>
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// <p>环境</p>
	Environments []*string `json:"Environments,omitnil,omitempty" name:"Environments"`

	// <p>筛选</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>旅程起点页面。可选。不传表示不过滤。</p>
	StartPage *string `json:"StartPage,omitnil,omitempty" name:"StartPage"`

	// <p>旅程终点页面。可选。不传表示不过滤。</p>
	EndPage *string `json:"EndPage,omitnil,omitempty" name:"EndPage"`

	// <p>旅程筛选条件 JSON 字符串，格式与 Filters 相同：[{Key,Op,Values}]。可选。不传表示不过滤。</p>
	JourneyFilters *string `json:"JourneyFilters,omitnil,omitempty" name:"JourneyFilters"`
}

func (r *DescribeSessionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SessionType")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "SortBy")
	delete(f, "SortOrder")
	delete(f, "UserId")
	delete(f, "SessionId")
	delete(f, "DeviceTypes")
	delete(f, "Browsers")
	delete(f, "OSList")
	delete(f, "Platforms")
	delete(f, "AppVersions")
	delete(f, "Countries")
	delete(f, "HasStruggle")
	delete(f, "StruggleRanks")
	delete(f, "StruggleTypes")
	delete(f, "HasError")
	delete(f, "VisitedPage")
	delete(f, "ContainsEvent")
	delete(f, "MinDuration")
	delete(f, "MaxDuration")
	delete(f, "PerfFilter")
	delete(f, "Regions")
	delete(f, "Environments")
	delete(f, "Filters")
	delete(f, "StartPage")
	delete(f, "EndPage")
	delete(f, "JourneyFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSessionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionListResponseParams struct {
	// <p>返回结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSessionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSessionListResponseParams `json:"Response"`
}

func (r *DescribeSessionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStruggleDistributionRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间戳</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>过滤条件</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>旅程筛选条件 JSON 字符串，格式与 Filters 相同：[{Key,Op,Values}]。可选。不传表示不过滤。</p>
	JourneyFilters *string `json:"JourneyFilters,omitnil,omitempty" name:"JourneyFilters"`

	// <p>是否只看包含挣扎的会话。可选。不传表示不过滤。</p>
	HasStruggle *bool `json:"HasStruggle,omitnil,omitempty" name:"HasStruggle"`

	// <p>旅程起点页面。可选。不传表示不过滤。</p>
	StartPage *string `json:"StartPage,omitnil,omitempty" name:"StartPage"`

	// <p>旅程终点页面。可选。不传表示不过滤。</p>
	EndPage *string `json:"EndPage,omitnil,omitempty" name:"EndPage"`
}

type DescribeStruggleDistributionRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间戳</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>过滤条件</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>旅程筛选条件 JSON 字符串，格式与 Filters 相同：[{Key,Op,Values}]。可选。不传表示不过滤。</p>
	JourneyFilters *string `json:"JourneyFilters,omitnil,omitempty" name:"JourneyFilters"`

	// <p>是否只看包含挣扎的会话。可选。不传表示不过滤。</p>
	HasStruggle *bool `json:"HasStruggle,omitnil,omitempty" name:"HasStruggle"`

	// <p>旅程起点页面。可选。不传表示不过滤。</p>
	StartPage *string `json:"StartPage,omitnil,omitempty" name:"StartPage"`

	// <p>旅程终点页面。可选。不传表示不过滤。</p>
	EndPage *string `json:"EndPage,omitnil,omitempty" name:"EndPage"`
}

func (r *DescribeStruggleDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStruggleDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Platform")
	delete(f, "Filters")
	delete(f, "JourneyFilters")
	delete(f, "HasStruggle")
	delete(f, "StartPage")
	delete(f, "EndPage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStruggleDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStruggleDistributionResponseParams struct {
	// <p>请求结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStruggleDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStruggleDistributionResponseParams `json:"Response"`
}

func (r *DescribeStruggleDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStruggleDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSymbolArtifactUploadCredentialRequestParams struct {
	// <p>制品类型。取值：js_sourcemap、ios_dsym、macos_dsym、windows_pdb、android_mapping、android_native、system_symbol。</p>
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// <p>上传文件名。</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>应用 ID。应用级制品必填；system_symbol 不传。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>兼容别名，未传 AppId 时使用。</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>保留字段，请勿传入。</p>
	Release *string `json:"Release,omitnil,omitempty" name:"Release"`

	// <p>保留字段，请勿传入。</p>
	Dist *string `json:"Dist,omitnil,omitempty" name:"Dist"`

	// <p>保留字段，请勿传入。</p>
	MinifiedUrl *string `json:"MinifiedUrl,omitnil,omitempty" name:"MinifiedUrl"`

	// <p>保留字段，请勿传入。</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>保留字段，请勿传入。</p>
	DebugId *string `json:"DebugId,omitnil,omitempty" name:"DebugId"`

	// <p>保留字段，请勿传入。</p>
	BuildId *string `json:"BuildId,omitnil,omitempty" name:"BuildId"`

	// <p>保留字段，请勿传入。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>保留字段，请勿传入。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>系统名称。</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>系统版本。</p>
	OSVersion *string `json:"OSVersion,omitnil,omitempty" name:"OSVersion"`

	// <p>保留字段，请勿传入。</p>
	StorePath *string `json:"StorePath,omitnil,omitempty" name:"StorePath"`

	// <p>文件内容类型，如 application/json。</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`
}

type DescribeSymbolArtifactUploadCredentialRequest struct {
	*tchttp.BaseRequest
	
	// <p>制品类型。取值：js_sourcemap、ios_dsym、macos_dsym、windows_pdb、android_mapping、android_native、system_symbol。</p>
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// <p>上传文件名。</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>应用 ID。应用级制品必填；system_symbol 不传。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>兼容别名，未传 AppId 时使用。</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>保留字段，请勿传入。</p>
	Release *string `json:"Release,omitnil,omitempty" name:"Release"`

	// <p>保留字段，请勿传入。</p>
	Dist *string `json:"Dist,omitnil,omitempty" name:"Dist"`

	// <p>保留字段，请勿传入。</p>
	MinifiedUrl *string `json:"MinifiedUrl,omitnil,omitempty" name:"MinifiedUrl"`

	// <p>保留字段，请勿传入。</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>保留字段，请勿传入。</p>
	DebugId *string `json:"DebugId,omitnil,omitempty" name:"DebugId"`

	// <p>保留字段，请勿传入。</p>
	BuildId *string `json:"BuildId,omitnil,omitempty" name:"BuildId"`

	// <p>保留字段，请勿传入。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>保留字段，请勿传入。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>系统名称。</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>系统版本。</p>
	OSVersion *string `json:"OSVersion,omitnil,omitempty" name:"OSVersion"`

	// <p>保留字段，请勿传入。</p>
	StorePath *string `json:"StorePath,omitnil,omitempty" name:"StorePath"`

	// <p>文件内容类型，如 application/json。</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`
}

func (r *DescribeSymbolArtifactUploadCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSymbolArtifactUploadCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Kind")
	delete(f, "FileName")
	delete(f, "AppId")
	delete(f, "ProjectId")
	delete(f, "Release")
	delete(f, "Dist")
	delete(f, "MinifiedUrl")
	delete(f, "FileType")
	delete(f, "DebugId")
	delete(f, "BuildId")
	delete(f, "AppVersion")
	delete(f, "Platform")
	delete(f, "OS")
	delete(f, "OSVersion")
	delete(f, "StorePath")
	delete(f, "ContentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSymbolArtifactUploadCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSymbolArtifactUploadCredentialResponseParams struct {
	// <p>上传临时凭证，JSON 格式字符串。</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSymbolArtifactUploadCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSymbolArtifactUploadCredentialResponseParams `json:"Response"`
}

func (r *DescribeSymbolArtifactUploadCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSymbolArtifactUploadCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSymbolArtifactsRequestParams struct {
	// <p>制品类型。取值：js_sourcemap、ios_dsym、macos_dsym、windows_pdb、android_mapping、android_native、system_symbol。native_symbol 仅查询历史数据。</p>
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// <p>应用 ID。js_sourcemap 必填，其他应用级制品用于过滤。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>兼容别名，未传 AppId 时使用。</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>JS 应用版本过滤条件。</p>
	Release *string `json:"Release,omitnil,omitempty" name:"Release"`

	// <p>JS 制品的历史 Dist 过滤条件。</p>
	Dist *string `json:"Dist,omitnil,omitempty" name:"Dist"`

	// <p>JS 制品文件类型过滤条件。</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>应用版本过滤条件。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>历史 native_symbol 的平台过滤条件。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>系统名称。system_symbol 场景过滤条件。</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>系统版本。system_symbol 场景过滤条件。</p>
	OSVersion *string `json:"OSVersion,omitnil,omitempty" name:"OSVersion"`

	// <p>状态过滤。可按 active、superseded、deleted、invalid 过滤。</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>分页大小。默认 50。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页游标。</p>
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

type DescribeSymbolArtifactsRequest struct {
	*tchttp.BaseRequest
	
	// <p>制品类型。取值：js_sourcemap、ios_dsym、macos_dsym、windows_pdb、android_mapping、android_native、system_symbol。native_symbol 仅查询历史数据。</p>
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// <p>应用 ID。js_sourcemap 必填，其他应用级制品用于过滤。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>兼容别名，未传 AppId 时使用。</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>JS 应用版本过滤条件。</p>
	Release *string `json:"Release,omitnil,omitempty" name:"Release"`

	// <p>JS 制品的历史 Dist 过滤条件。</p>
	Dist *string `json:"Dist,omitnil,omitempty" name:"Dist"`

	// <p>JS 制品文件类型过滤条件。</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>应用版本过滤条件。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>历史 native_symbol 的平台过滤条件。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>系统名称。system_symbol 场景过滤条件。</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>系统版本。system_symbol 场景过滤条件。</p>
	OSVersion *string `json:"OSVersion,omitnil,omitempty" name:"OSVersion"`

	// <p>状态过滤。可按 active、superseded、deleted、invalid 过滤。</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>分页大小。默认 50。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页游标。</p>
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

func (r *DescribeSymbolArtifactsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSymbolArtifactsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Kind")
	delete(f, "AppId")
	delete(f, "ProjectId")
	delete(f, "Release")
	delete(f, "Dist")
	delete(f, "FileType")
	delete(f, "AppVersion")
	delete(f, "Platform")
	delete(f, "OS")
	delete(f, "OSVersion")
	delete(f, "Status")
	delete(f, "Limit")
	delete(f, "Cursor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSymbolArtifactsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSymbolArtifactsResponseParams struct {
	// <p>制品列表，JSON 格式字符串，包含 artifacts、next_cursor 和 total。</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSymbolArtifactsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSymbolArtifactsResponseParams `json:"Response"`
}

func (r *DescribeSymbolArtifactsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSymbolArtifactsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMApplicationsRequestParams struct {
	// <p>分页偏移，默认 0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量，默认 20，上限 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>业务系统 ID（可选，缺省返回名下全部应用）</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>演示模式开关</p>
	DemoMode *bool `json:"DemoMode,omitnil,omitempty" name:"DemoMode"`
}

type DescribeTDEMApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// <p>分页偏移，默认 0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量，默认 20，上限 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>业务系统 ID（可选，缺省返回名下全部应用）</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>演示模式开关</p>
	DemoMode *bool `json:"DemoMode,omitnil,omitempty" name:"DemoMode"`
}

func (r *DescribeTDEMApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	delete(f, "DemoMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMApplicationsResponseParams struct {
	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>应用列表</p>
	ApplicationSet []*ApplicationItem `json:"ApplicationSet,omitnil,omitempty" name:"ApplicationSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMApplicationsResponseParams `json:"Response"`
}

func (r *DescribeTDEMApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMBusinessSystemsRequestParams struct {
	// <p>分页偏移，默认 0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量，默认 20，上限 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>演示模式开关</p>
	DemoMode *bool `json:"DemoMode,omitnil,omitempty" name:"DemoMode"`
}

type DescribeTDEMBusinessSystemsRequest struct {
	*tchttp.BaseRequest
	
	// <p>分页偏移，默认 0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量，默认 20，上限 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>演示模式开关</p>
	DemoMode *bool `json:"DemoMode,omitnil,omitempty" name:"DemoMode"`
}

func (r *DescribeTDEMBusinessSystemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMBusinessSystemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DemoMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMBusinessSystemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMBusinessSystemsResponseParams struct {
	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>业务系统列表</p>
	BusinessSystemSet []*BusinessSystemItem `json:"BusinessSystemSet,omitnil,omitempty" name:"BusinessSystemSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMBusinessSystemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMBusinessSystemsResponseParams `json:"Response"`
}

func (r *DescribeTDEMBusinessSystemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMBusinessSystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMDialDimensionsRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，毫秒时间戳，必须大于 StartTime</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>维度。取值说明：platform、app_version、device、net_type、operator、region、country、target_ip。</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>任务 ID。不传表示不过滤。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>客户端平台。不传表示不过滤。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本。不传表示不过滤。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备机型，对应 device_type。不传表示不过滤。</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>网络类型，对应 network_type。不传表示不过滤。</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>客户端运营商，对应 isp。不传表示不过滤。不是目标 IP 运营商。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>客户端地区，对应 province。云 API 公共参数已占用 Region，筛选入参名为 ClientRegion。不传表示不过滤。不是目标 IP 地区。</p>
	ClientRegion *string `json:"ClientRegion,omitnil,omitempty" name:"ClientRegion"`

	// <p>目标 IP。不传表示不过滤。</p>
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// <p>环境。可选。不传表示不过滤。</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

type DescribeTDEMDialDimensionsRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，毫秒时间戳，必须大于 StartTime</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>维度。取值说明：platform、app_version、device、net_type、operator、region、country、target_ip。</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>任务 ID。不传表示不过滤。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>客户端平台。不传表示不过滤。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本。不传表示不过滤。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备机型，对应 device_type。不传表示不过滤。</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>网络类型，对应 network_type。不传表示不过滤。</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>客户端运营商，对应 isp。不传表示不过滤。不是目标 IP 运营商。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>客户端地区，对应 province。云 API 公共参数已占用 Region，筛选入参名为 ClientRegion。不传表示不过滤。不是目标 IP 地区。</p>
	ClientRegion *string `json:"ClientRegion,omitnil,omitempty" name:"ClientRegion"`

	// <p>目标 IP。不传表示不过滤。</p>
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// <p>环境。可选。不传表示不过滤。</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

func (r *DescribeTDEMDialDimensionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMDialDimensionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Dimension")
	delete(f, "TaskId")
	delete(f, "Platform")
	delete(f, "AppVersion")
	delete(f, "Device")
	delete(f, "NetType")
	delete(f, "Operator")
	delete(f, "ClientRegion")
	delete(f, "TargetIp")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMDialDimensionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMDialDimensionsResponseParams struct {
	// <p>本次查询的维度</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>维度统计项</p>
	Items []*DialDimensionItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMDialDimensionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMDialDimensionsResponseParams `json:"Response"`
}

func (r *DescribeTDEMDialDimensionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMDialDimensionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMDialExecutionDetailRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，毫秒时间戳，必须大于 StartTime</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>任务 ID。不传表示不过滤。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>拨测批次 ID</p>
	TaskSessionId *string `json:"TaskSessionId,omitnil,omitempty" name:"TaskSessionId"`

	// <p>客户端平台。不传表示不过滤。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本。不传表示不过滤。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备机型，对应 device_type。不传表示不过滤。</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>网络类型，对应 network_type。不传表示不过滤。</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>客户端运营商，对应 isp。不传表示不过滤。不是目标 IP 运营商。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>客户端地区，对应 province。云 API 公共参数已占用 Region，筛选入参名为 ClientRegion。不传表示不过滤。不是目标 IP 地区。</p>
	ClientRegion *string `json:"ClientRegion,omitnil,omitempty" name:"ClientRegion"`

	// <p>目标 IP。不传表示不过滤。</p>
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// <p>环境。可选。不传表示不过滤。</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

type DescribeTDEMDialExecutionDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，毫秒时间戳，必须大于 StartTime</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>任务 ID。不传表示不过滤。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>拨测批次 ID</p>
	TaskSessionId *string `json:"TaskSessionId,omitnil,omitempty" name:"TaskSessionId"`

	// <p>客户端平台。不传表示不过滤。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本。不传表示不过滤。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备机型，对应 device_type。不传表示不过滤。</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>网络类型，对应 network_type。不传表示不过滤。</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>客户端运营商，对应 isp。不传表示不过滤。不是目标 IP 运营商。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>客户端地区，对应 province。云 API 公共参数已占用 Region，筛选入参名为 ClientRegion。不传表示不过滤。不是目标 IP 地区。</p>
	ClientRegion *string `json:"ClientRegion,omitnil,omitempty" name:"ClientRegion"`

	// <p>目标 IP。不传表示不过滤。</p>
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// <p>环境。可选。不传表示不过滤。</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

func (r *DescribeTDEMDialExecutionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMDialExecutionDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskId")
	delete(f, "TaskSessionId")
	delete(f, "Platform")
	delete(f, "AppVersion")
	delete(f, "Device")
	delete(f, "NetType")
	delete(f, "Operator")
	delete(f, "ClientRegion")
	delete(f, "TargetIp")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMDialExecutionDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMDialExecutionDetailResponseParams struct {
	// <p>任务 ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务名称</p>
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>拨测批次 ID</p>
	TaskSessionId *string `json:"TaskSessionId,omitnil,omitempty" name:"TaskSessionId"`

	// <p>批次开始时间，毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>所属用户会话 ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>本批次是否全部探测成功</p>
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`

	// <p>平均时延，单位毫秒</p>
	AvgDuration *float64 `json:"AvgDuration,omitnil,omitempty" name:"AvgDuration"`

	// <p>客户端平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备机型</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>网络类型</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>客户端运营商</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>客户端地区</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>本批次目标 IP</p>
	TargetIps []*string `json:"TargetIps,omitnil,omitempty" name:"TargetIps"`

	// <p>每个目标的探测结果</p>
	Probes []*DialProbe `json:"Probes,omitnil,omitempty" name:"Probes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMDialExecutionDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMDialExecutionDetailResponseParams `json:"Response"`
}

func (r *DescribeTDEMDialExecutionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMDialExecutionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMDialExecutionListRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，毫秒时间戳，必须大于 StartTime</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>任务 ID。不传表示不过滤。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>客户端平台。不传表示不过滤。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本。不传表示不过滤。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备机型，对应 device_type。不传表示不过滤。</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>网络类型，对应 network_type。不传表示不过滤。</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>客户端运营商，对应 isp。不传表示不过滤。不是目标 IP 运营商。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>客户端地区，对应 province。云 API 公共参数已占用 Region，筛选入参名为 ClientRegion。不传表示不过滤。不是目标 IP 地区。</p>
	ClientRegion *string `json:"ClientRegion,omitnil,omitempty" name:"ClientRegion"`

	// <p>目标 IP。不传表示不过滤。</p>
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// <p>返回条数。范围 1 到 100，不传时为 20。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量。不传时为 0。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>是否只看成功或失败批次。不传表示不过滤。</p>
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`

	// <p>环境。可选。不传表示不过滤。</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

type DescribeTDEMDialExecutionListRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，毫秒时间戳，必须大于 StartTime</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>任务 ID。不传表示不过滤。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>客户端平台。不传表示不过滤。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本。不传表示不过滤。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备机型，对应 device_type。不传表示不过滤。</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>网络类型，对应 network_type。不传表示不过滤。</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>客户端运营商，对应 isp。不传表示不过滤。不是目标 IP 运营商。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>客户端地区，对应 province。云 API 公共参数已占用 Region，筛选入参名为 ClientRegion。不传表示不过滤。不是目标 IP 地区。</p>
	ClientRegion *string `json:"ClientRegion,omitnil,omitempty" name:"ClientRegion"`

	// <p>目标 IP。不传表示不过滤。</p>
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// <p>返回条数。范围 1 到 100，不传时为 20。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量。不传时为 0。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>是否只看成功或失败批次。不传表示不过滤。</p>
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`

	// <p>环境。可选。不传表示不过滤。</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

func (r *DescribeTDEMDialExecutionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMDialExecutionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskId")
	delete(f, "Platform")
	delete(f, "AppVersion")
	delete(f, "Device")
	delete(f, "NetType")
	delete(f, "Operator")
	delete(f, "ClientRegion")
	delete(f, "TargetIp")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Success")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMDialExecutionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMDialExecutionListResponseParams struct {
	// <p>满足条件的批次数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>执行批次列表</p>
	ExecutionSet []*DialExecution `json:"ExecutionSet,omitnil,omitempty" name:"ExecutionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMDialExecutionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMDialExecutionListResponseParams `json:"Response"`
}

func (r *DescribeTDEMDialExecutionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMDialExecutionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMDialOverviewRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，毫秒时间戳，必须大于 StartTime</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>任务 ID。不传表示不过滤。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>客户端平台。不传表示不过滤。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本。不传表示不过滤。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备机型，对应 device_type。不传表示不过滤。</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>网络类型，对应 network_type。不传表示不过滤。</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>客户端运营商，对应 isp。不传表示不过滤。不是目标 IP 运营商。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>客户端地区，对应 province。云 API 公共参数已占用 Region，筛选入参名为 ClientRegion。不传表示不过滤。不是目标 IP 地区。</p>
	ClientRegion *string `json:"ClientRegion,omitnil,omitempty" name:"ClientRegion"`

	// <p>目标 IP。不传表示不过滤。</p>
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// <p>环境。可选。不传表示不过滤。</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

type DescribeTDEMDialOverviewRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，毫秒时间戳，必须大于 StartTime</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>任务 ID。不传表示不过滤。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>客户端平台。不传表示不过滤。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本。不传表示不过滤。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备机型，对应 device_type。不传表示不过滤。</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>网络类型，对应 network_type。不传表示不过滤。</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>客户端运营商，对应 isp。不传表示不过滤。不是目标 IP 运营商。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>客户端地区，对应 province。云 API 公共参数已占用 Region，筛选入参名为 ClientRegion。不传表示不过滤。不是目标 IP 地区。</p>
	ClientRegion *string `json:"ClientRegion,omitnil,omitempty" name:"ClientRegion"`

	// <p>目标 IP。不传表示不过滤。</p>
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// <p>环境。可选。不传表示不过滤。</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

func (r *DescribeTDEMDialOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMDialOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskId")
	delete(f, "Platform")
	delete(f, "AppVersion")
	delete(f, "Device")
	delete(f, "NetType")
	delete(f, "Operator")
	delete(f, "ClientRegion")
	delete(f, "TargetIp")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMDialOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMDialOverviewResponseParams struct {
	// <p>任务数</p>
	TaskCount *int64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// <p>执行批次数</p>
	ExecutionCount *int64 `json:"ExecutionCount,omitnil,omitempty" name:"ExecutionCount"`

	// <p>成功批次数</p>
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// <p>失败批次数</p>
	FailCount *int64 `json:"FailCount,omitnil,omitempty" name:"FailCount"`

	// <p>平均时延，单位毫秒</p>
	AvgDuration *float64 `json:"AvgDuration,omitnil,omitempty" name:"AvgDuration"`

	// <p>平均丢包率，0 到 100</p>
	AvgLoss *float64 `json:"AvgLoss,omitnil,omitempty" name:"AvgLoss"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMDialOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMDialOverviewResponseParams `json:"Response"`
}

func (r *DescribeTDEMDialOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMDialOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMDialTasksRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DescribeTDEMDialTasksRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *DescribeTDEMDialTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMDialTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMDialTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMDialTasksResponseParams struct {
	// <p>拨测任务列表</p>
	TaskSet []*DialTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMDialTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMDialTasksResponseParams `json:"Response"`
}

func (r *DescribeTDEMDialTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMDialTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMDialTrendRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，毫秒时间戳，必须大于 StartTime</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>任务 ID。不传表示不过滤。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>客户端平台。不传表示不过滤。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本。不传表示不过滤。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备机型，对应 device_type。不传表示不过滤。</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>网络类型，对应 network_type。不传表示不过滤。</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>客户端运营商，对应 isp。不传表示不过滤。不是目标 IP 运营商。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>客户端地区，对应 province。云 API 公共参数已占用 Region，筛选入参名为 ClientRegion。不传表示不过滤。不是目标 IP 地区。</p>
	ClientRegion *string `json:"ClientRegion,omitnil,omitempty" name:"ClientRegion"`

	// <p>目标 IP。不传表示不过滤。</p>
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// <p>时间粒度。取值说明：hour、day。不传时为 hour。</p>
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// <p>环境。可选。不传表示不过滤。</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

type DescribeTDEMDialTrendRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，毫秒时间戳，必须大于 StartTime</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>任务 ID。不传表示不过滤。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>客户端平台。不传表示不过滤。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本。不传表示不过滤。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备机型，对应 device_type。不传表示不过滤。</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>网络类型，对应 network_type。不传表示不过滤。</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>客户端运营商，对应 isp。不传表示不过滤。不是目标 IP 运营商。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>客户端地区，对应 province。云 API 公共参数已占用 Region，筛选入参名为 ClientRegion。不传表示不过滤。不是目标 IP 地区。</p>
	ClientRegion *string `json:"ClientRegion,omitnil,omitempty" name:"ClientRegion"`

	// <p>目标 IP。不传表示不过滤。</p>
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// <p>时间粒度。取值说明：hour、day。不传时为 hour。</p>
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// <p>环境。可选。不传表示不过滤。</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

func (r *DescribeTDEMDialTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMDialTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskId")
	delete(f, "Platform")
	delete(f, "AppVersion")
	delete(f, "Device")
	delete(f, "NetType")
	delete(f, "Operator")
	delete(f, "ClientRegion")
	delete(f, "TargetIp")
	delete(f, "Granularity")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMDialTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMDialTrendResponseParams struct {
	// <p>趋势点</p>
	Points []*DialTrendPoint `json:"Points,omitnil,omitempty" name:"Points"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMDialTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMDialTrendResponseParams `json:"Response"`
}

func (r *DescribeTDEMDialTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMDialTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMEventDetailRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>事件 ID（必填）</p>
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// <p>分区日期 YYYY-MM-DD（推荐）</p>
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

type DescribeTDEMEventDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>事件 ID（必填）</p>
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// <p>分区日期 YYYY-MM-DD（推荐）</p>
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

func (r *DescribeTDEMEventDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMEventDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "EventId")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMEventDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMEventDetailResponseParams struct {
	// <p>事件主表视图</p>
	Event *EventView `json:"Event,omitnil,omitempty" name:"Event"`

	// <p>符号化结果，未翻译时可省略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Symbolicated *SymbolicatedView `json:"Symbolicated,omitnil,omitempty" name:"Symbolicated"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMEventDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMEventDetailResponseParams `json:"Response"`
}

func (r *DescribeTDEMEventDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMEventRankRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>排行类型：error/api/resource/api_page/resource_page/api_status/api_status_top/resource_status/tag；exception_dimension 为异常维度同源统计，使用 ExceptionDistribution 返回结果。</p>
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// <p>旧 Kind 保持既有错误类型过滤规则；Kind=exception_dimension 标签模式必填，允许 js_error、promise_error、api_error、resource_error、blank_screen、log、native_crash、crash、no_response、lag_detected、app_launch。合法重复值去重，未知值拒绝；普通维度模式不传。</p>
	EventTypes []*string `json:"EventTypes,omitnil,omitempty" name:"EventTypes"`

	// <p>旧 Kind=tag 保持既有标签查询；Kind=exception_dimension 标签模式必填，与 GroupBy 二选一。使用既有标签名规范，长度 1–64，允许字母数字及 _-.；标签值保留原始内容。</p>
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// <p>返回条数，上限100。旧 Kind 缺省50；Kind=exception_dimension 缺省100，显式值为1–100，限制完整归一化维度行数，不裁剪分类向量。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>旧 Kind 可选粒度5m/1h/1d；Kind=exception_dimension 不支持，传非空值返回 InvalidParameter。</p>
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// <p>筛选条件 JSON 字符串，走 events 表白名单</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>旧 Kind 耗时口径 avg/p50/p75/p95/p99；Kind=exception_dimension 不支持，传非空值返回 InvalidParameter。</p>
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// <p>异常分布的聚合维度，仅 Kind=exception_dimension 的普通维度模式使用且必填。标签模式只传 TagKey，不传 GroupBy。</p><p>枚举值：</p><ul><li>device_type： 设备型号</li><li>os： 操作系统</li><li>browser： 浏览器</li><li>app_version： 应用版本</li><li>province： 省份</li><li>platform： 平台</li><li>network_type： 网络类型</li></ul>
	GroupBy *string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// <p>仅 Kind=exception_dimension 普通维度模式必填。允许 web、android、ios，须与应用平台及 CategoryKeys 目录一致。标签模式不传。</p>
	ProjectType *string `json:"ProjectType,omitnil,omitempty" name:"ProjectType"`

	// <p>仅 Kind=exception_dimension 普通维度模式必填，为页面实际有效的异常目录 key 非空去重集合。仅接受对应 ProjectType 的固定目录，包含项须同时选择父项；未知、重复、空目录及互斥参数组合返回 InvalidParameter。标签模式不传。</p>
	CategoryKeys []*string `json:"CategoryKeys,omitnil,omitempty" name:"CategoryKeys"`
}

type DescribeTDEMEventRankRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>排行类型：error/api/resource/api_page/resource_page/api_status/api_status_top/resource_status/tag；exception_dimension 为异常维度同源统计，使用 ExceptionDistribution 返回结果。</p>
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// <p>旧 Kind 保持既有错误类型过滤规则；Kind=exception_dimension 标签模式必填，允许 js_error、promise_error、api_error、resource_error、blank_screen、log、native_crash、crash、no_response、lag_detected、app_launch。合法重复值去重，未知值拒绝；普通维度模式不传。</p>
	EventTypes []*string `json:"EventTypes,omitnil,omitempty" name:"EventTypes"`

	// <p>旧 Kind=tag 保持既有标签查询；Kind=exception_dimension 标签模式必填，与 GroupBy 二选一。使用既有标签名规范，长度 1–64，允许字母数字及 _-.；标签值保留原始内容。</p>
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// <p>返回条数，上限100。旧 Kind 缺省50；Kind=exception_dimension 缺省100，显式值为1–100，限制完整归一化维度行数，不裁剪分类向量。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>旧 Kind 可选粒度5m/1h/1d；Kind=exception_dimension 不支持，传非空值返回 InvalidParameter。</p>
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// <p>筛选条件 JSON 字符串，走 events 表白名单</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>旧 Kind 耗时口径 avg/p50/p75/p95/p99；Kind=exception_dimension 不支持，传非空值返回 InvalidParameter。</p>
	CostType *string `json:"CostType,omitnil,omitempty" name:"CostType"`

	// <p>异常分布的聚合维度，仅 Kind=exception_dimension 的普通维度模式使用且必填。标签模式只传 TagKey，不传 GroupBy。</p><p>枚举值：</p><ul><li>device_type： 设备型号</li><li>os： 操作系统</li><li>browser： 浏览器</li><li>app_version： 应用版本</li><li>province： 省份</li><li>platform： 平台</li><li>network_type： 网络类型</li></ul>
	GroupBy *string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// <p>仅 Kind=exception_dimension 普通维度模式必填。允许 web、android、ios，须与应用平台及 CategoryKeys 目录一致。标签模式不传。</p>
	ProjectType *string `json:"ProjectType,omitnil,omitempty" name:"ProjectType"`

	// <p>仅 Kind=exception_dimension 普通维度模式必填，为页面实际有效的异常目录 key 非空去重集合。仅接受对应 ProjectType 的固定目录，包含项须同时选择父项；未知、重复、空目录及互斥参数组合返回 InvalidParameter。标签模式不传。</p>
	CategoryKeys []*string `json:"CategoryKeys,omitnil,omitempty" name:"CategoryKeys"`
}

func (r *DescribeTDEMEventRankRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMEventRankRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Kind")
	delete(f, "EventTypes")
	delete(f, "TagKey")
	delete(f, "Limit")
	delete(f, "Granularity")
	delete(f, "Filters")
	delete(f, "CostType")
	delete(f, "GroupBy")
	delete(f, "ProjectType")
	delete(f, "CategoryKeys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMEventRankRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMEventRankResponseParams struct {
	// <p>排行行</p>
	Items []*EventRankItem `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>仅 Kind=exception_dimension 返回。Rows 为完整维度行及固定分类向量；空数据返回 Rows=[]。父行 TotalCount 等于 IncludedInTotal=true 的主项 Count 之和，包含项仅展示且不重复计总。所有计数为0至9007199254740991的整数。旧 Kind 继续使用 Items。</p>
	ExceptionDistribution *ExceptionDistribution `json:"ExceptionDistribution,omitnil,omitempty" name:"ExceptionDistribution"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMEventRankResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMEventRankResponseParams `json:"Response"`
}

func (r *DescribeTDEMEventRankResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMEventRankResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMEventSearchRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>事件分类（必填，禁止 session）</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>筛选条件 JSON 字符串：[{Key,Op,Values}]，须在探索白名单</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>排序字段：timestamp（默认）/ duration</p>
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// <p>排序方向：asc / desc，默认 desc</p>
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// <p>页码，&gt;=1，默认 1</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>每页条数，1-100，默认 20</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>是否只看有错误的会话。可选。不传表示不过滤。Category 为 session 的汇总查询会使用该条件。</p>
	HasError *bool `json:"HasError,omitnil,omitempty" name:"HasError"`

	// <p>会话最短时长，单位毫秒。可选。0 或不传表示不限制。Category 为 session 的汇总查询会使用该条件。</p>
	MinDuration *int64 `json:"MinDuration,omitnil,omitempty" name:"MinDuration"`

	// <p>会话最长时长，单位毫秒。可选。0 或不传表示不限制。Category 为 session 的汇总查询会使用该条件。</p>
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`
}

type DescribeTDEMEventSearchRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>事件分类（必填，禁止 session）</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>筛选条件 JSON 字符串：[{Key,Op,Values}]，须在探索白名单</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>排序字段：timestamp（默认）/ duration</p>
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// <p>排序方向：asc / desc，默认 desc</p>
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// <p>页码，&gt;=1，默认 1</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>每页条数，1-100，默认 20</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>是否只看有错误的会话。可选。不传表示不过滤。Category 为 session 的汇总查询会使用该条件。</p>
	HasError *bool `json:"HasError,omitnil,omitempty" name:"HasError"`

	// <p>会话最短时长，单位毫秒。可选。0 或不传表示不限制。Category 为 session 的汇总查询会使用该条件。</p>
	MinDuration *int64 `json:"MinDuration,omitnil,omitempty" name:"MinDuration"`

	// <p>会话最长时长，单位毫秒。可选。0 或不传表示不限制。Category 为 session 的汇总查询会使用该条件。</p>
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`
}

func (r *DescribeTDEMEventSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMEventSearchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Category")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "SortOrder")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "HasError")
	delete(f, "MinDuration")
	delete(f, "MaxDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMEventSearchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMEventSearchResponseParams struct {
	// <p>事件列表</p>
	Events []*ExplorationEventItem `json:"Events,omitnil,omitempty" name:"Events"`

	// <p>总条数</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>页码</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>每页条数</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMEventSearchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMEventSearchResponseParams `json:"Response"`
}

func (r *DescribeTDEMEventSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMEventSearchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMEventSearchSummaryRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>事件分类（必填，可含 session）</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>筛选条件 JSON 字符串：[{Key,Op,Values}]</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>分桶秒数，正整数；缺省由服务端按窗口选取</p>
	IntervalSec *int64 `json:"IntervalSec,omitnil,omitempty" name:"IntervalSec"`

	// <p>是否只看有错误的会话。可选。不传表示不过滤。Category 为 session 的汇总查询会使用该条件。</p>
	HasError *bool `json:"HasError,omitnil,omitempty" name:"HasError"`

	// <p>会话最短时长，单位毫秒。可选。0 或不传表示不限制。Category 为 session 的汇总查询会使用该条件。</p>
	MinDuration *int64 `json:"MinDuration,omitnil,omitempty" name:"MinDuration"`

	// <p>会话最长时长，单位毫秒。可选。0 或不传表示不限制。Category 为 session 的汇总查询会使用该条件。</p>
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`
}

type DescribeTDEMEventSearchSummaryRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>事件分类（必填，可含 session）</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>筛选条件 JSON 字符串：[{Key,Op,Values}]</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>分桶秒数，正整数；缺省由服务端按窗口选取</p>
	IntervalSec *int64 `json:"IntervalSec,omitnil,omitempty" name:"IntervalSec"`

	// <p>是否只看有错误的会话。可选。不传表示不过滤。Category 为 session 的汇总查询会使用该条件。</p>
	HasError *bool `json:"HasError,omitnil,omitempty" name:"HasError"`

	// <p>会话最短时长，单位毫秒。可选。0 或不传表示不限制。Category 为 session 的汇总查询会使用该条件。</p>
	MinDuration *int64 `json:"MinDuration,omitnil,omitempty" name:"MinDuration"`

	// <p>会话最长时长，单位毫秒。可选。0 或不传表示不限制。Category 为 session 的汇总查询会使用该条件。</p>
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`
}

func (r *DescribeTDEMEventSearchSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMEventSearchSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Category")
	delete(f, "Filters")
	delete(f, "IntervalSec")
	delete(f, "HasError")
	delete(f, "MinDuration")
	delete(f, "MaxDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMEventSearchSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMEventSearchSummaryResponseParams struct {
	// <p>各分类窗口计数，Key 为分类 id</p>
	CategoryCounts []*IntKeyValue `json:"CategoryCounts,omitnil,omitempty" name:"CategoryCounts"`

	// <p>当前分类窗口条数</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>访问用户（设备）</p>
	Visitors *int64 `json:"Visitors,omitnil,omitempty" name:"Visitors"`

	// <p>业务用户</p>
	Users *int64 `json:"Users,omitnil,omitempty" name:"Users"`

	// <p>实际分桶秒数</p>
	IntervalSec *int64 `json:"IntervalSec,omitnil,omitempty" name:"IntervalSec"`

	// <p>分桶趋势</p>
	Series []*EventSearchSeriesPoint `json:"Series,omitnil,omitempty" name:"Series"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMEventSearchSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMEventSearchSummaryResponseParams `json:"Response"`
}

func (r *DescribeTDEMEventSearchSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMEventSearchSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFieldTransformRulesRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>分页偏移，默认 0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTDEMFieldTransformRulesRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>分页偏移，默认 0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTDEMFieldTransformRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFieldTransformRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMFieldTransformRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFieldTransformRulesResponseParams struct {
	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>规则列表</p>
	FieldTransformRuleSet []*FieldTransformRuleItem `json:"FieldTransformRuleSet,omitnil,omitempty" name:"FieldTransformRuleSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMFieldTransformRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMFieldTransformRulesResponseParams `json:"Response"`
}

func (r *DescribeTDEMFieldTransformRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFieldTransformRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFunnelDetailRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`
}

type DescribeTDEMFunnelDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`
}

func (r *DescribeTDEMFunnelDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFunnelDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FunnelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMFunnelDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFunnelDetailResponseParams struct {
	// <p>请求结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMFunnelDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMFunnelDetailResponseParams `json:"Response"`
}

func (r *DescribeTDEMFunnelDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFunnelDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFunnelDimensionsRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>维度</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>设备类型</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>浏览器</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>操作系统</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>页面来源</p>
	Referrer *string `json:"Referrer,omitnil,omitempty" name:"Referrer"`

	// <p>应用版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>环境</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

type DescribeTDEMFunnelDimensionsRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>维度</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>设备类型</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>浏览器</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>操作系统</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>页面来源</p>
	Referrer *string `json:"Referrer,omitnil,omitempty" name:"Referrer"`

	// <p>应用版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>环境</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

func (r *DescribeTDEMFunnelDimensionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFunnelDimensionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FunnelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Dimension")
	delete(f, "Platform")
	delete(f, "DeviceType")
	delete(f, "Browser")
	delete(f, "OS")
	delete(f, "Country")
	delete(f, "Referrer")
	delete(f, "AppVersion")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMFunnelDimensionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFunnelDimensionsResponseParams struct {
	// <p>结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMFunnelDimensionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMFunnelDimensionsResponseParams `json:"Response"`
}

func (r *DescribeTDEMFunnelDimensionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFunnelDimensionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFunnelListRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>关键字，用于过滤漏斗名</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>漏斗状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>页面号</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>页面数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeTDEMFunnelListRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>关键字，用于过滤漏斗名</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>漏斗状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>页面号</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>页面数量</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeTDEMFunnelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFunnelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Keyword")
	delete(f, "Status")
	delete(f, "Page")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMFunnelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFunnelListResponseParams struct {
	// <p>结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMFunnelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMFunnelListResponseParams `json:"Response"`
}

func (r *DescribeTDEMFunnelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFunnelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFunnelReportRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>设备类型</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>浏览器</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>操作系统</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>页面来源</p>
	Referrer *string `json:"Referrer,omitnil,omitempty" name:"Referrer"`

	// <p>应用版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>环境选择</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

type DescribeTDEMFunnelReportRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>设备类型</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>浏览器</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>操作系统</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>页面来源</p>
	Referrer *string `json:"Referrer,omitnil,omitempty" name:"Referrer"`

	// <p>应用版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>环境选择</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

func (r *DescribeTDEMFunnelReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFunnelReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FunnelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Platform")
	delete(f, "DeviceType")
	delete(f, "Browser")
	delete(f, "OS")
	delete(f, "Country")
	delete(f, "Referrer")
	delete(f, "AppVersion")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMFunnelReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFunnelReportResponseParams struct {
	// <p>结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMFunnelReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMFunnelReportResponseParams `json:"Response"`
}

func (r *DescribeTDEMFunnelReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFunnelReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFunnelStepSessionListRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>步骤索引</p>
	StepIndex *int64 `json:"StepIndex,omitnil,omitempty" name:"StepIndex"`

	// <p>漏斗步骤</p>
	Steps *string `json:"Steps,omitnil,omitempty" name:"Steps"`

	// <p>下钻类型</p>
	DrillType *string `json:"DrillType,omitnil,omitempty" name:"DrillType"`

	// <p>页面号</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>每页大小</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备类型</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>浏览器</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>操作系统</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>页面来源</p>
	Referrer *string `json:"Referrer,omitnil,omitempty" name:"Referrer"`

	// <p>环境</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

type DescribeTDEMFunnelStepSessionListRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>步骤索引</p>
	StepIndex *int64 `json:"StepIndex,omitnil,omitempty" name:"StepIndex"`

	// <p>漏斗步骤</p>
	Steps *string `json:"Steps,omitnil,omitempty" name:"Steps"`

	// <p>下钻类型</p>
	DrillType *string `json:"DrillType,omitnil,omitempty" name:"DrillType"`

	// <p>页面号</p>
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>每页大小</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备类型</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>浏览器</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>操作系统</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>页面来源</p>
	Referrer *string `json:"Referrer,omitnil,omitempty" name:"Referrer"`

	// <p>环境</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

func (r *DescribeTDEMFunnelStepSessionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFunnelStepSessionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FunnelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "StepIndex")
	delete(f, "Steps")
	delete(f, "DrillType")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "Platform")
	delete(f, "AppVersion")
	delete(f, "DeviceType")
	delete(f, "Browser")
	delete(f, "OS")
	delete(f, "Country")
	delete(f, "Referrer")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMFunnelStepSessionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFunnelStepSessionListResponseParams struct {
	// <p>请求结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMFunnelStepSessionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMFunnelStepSessionListResponseParams `json:"Response"`
}

func (r *DescribeTDEMFunnelStepSessionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFunnelStepSessionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFunnelTrendRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>间隔</p>
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备类型</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>浏览器</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>操作系统</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>页面来源</p>
	Referrer *string `json:"Referrer,omitnil,omitempty" name:"Referrer"`

	// <p>环境</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

type DescribeTDEMFunnelTrendRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>间隔</p>
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备类型</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>浏览器</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>操作系统</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>页面来源</p>
	Referrer *string `json:"Referrer,omitnil,omitempty" name:"Referrer"`

	// <p>环境</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

func (r *DescribeTDEMFunnelTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFunnelTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FunnelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Interval")
	delete(f, "Platform")
	delete(f, "AppVersion")
	delete(f, "DeviceType")
	delete(f, "Browser")
	delete(f, "OS")
	delete(f, "Country")
	delete(f, "Referrer")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMFunnelTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMFunnelTrendResponseParams struct {
	// <p>结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMFunnelTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMFunnelTrendResponseParams `json:"Response"`
}

func (r *DescribeTDEMFunnelTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMFunnelTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMMetricDistributionRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>顶层事件类型；单查必填，批量用 Queries 时可不填</p>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>子类型；未知返回 0</p>
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// <p>筛选条件 JSON 字符串：[{Key,Op,Values}]，走分析立方白名单</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>指标名列表，如 count / rate；单查必填</p>
	Metrics []*string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// <p>分位/均值必填，如 lcp / duration</p>
	MeasurementKey *string `json:"MeasurementKey,omitnil,omitempty" name:"MeasurementKey"`

	// <p>单维度（Distribution 用）</p>
	GroupBy *string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// <p>返回条数，Top 默认 10 上限 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>批量查询，上限 32</p>
	Queries []*MetricQueryItem `json:"Queries,omitnil,omitempty" name:"Queries"`
}

type DescribeTDEMMetricDistributionRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>顶层事件类型；单查必填，批量用 Queries 时可不填</p>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>子类型；未知返回 0</p>
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// <p>筛选条件 JSON 字符串：[{Key,Op,Values}]，走分析立方白名单</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>指标名列表，如 count / rate；单查必填</p>
	Metrics []*string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// <p>分位/均值必填，如 lcp / duration</p>
	MeasurementKey *string `json:"MeasurementKey,omitnil,omitempty" name:"MeasurementKey"`

	// <p>单维度（Distribution 用）</p>
	GroupBy *string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// <p>返回条数，Top 默认 10 上限 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>批量查询，上限 32</p>
	Queries []*MetricQueryItem `json:"Queries,omitnil,omitempty" name:"Queries"`
}

func (r *DescribeTDEMMetricDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMMetricDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "EventType")
	delete(f, "SubType")
	delete(f, "Filters")
	delete(f, "Metrics")
	delete(f, "MeasurementKey")
	delete(f, "GroupBy")
	delete(f, "Limit")
	delete(f, "Queries")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMMetricDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMMetricDistributionResponseParams struct {
	// <p>单查分布桶</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*MetricDistributionItem `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>单查合计，批量为空时省略</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>批量结果，Queries 非空时返回</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*MetricDistributionBatchItem `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMMetricDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMMetricDistributionResponseParams `json:"Response"`
}

func (r *DescribeTDEMMetricDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMMetricDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMMetricSummaryRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>顶层事件类型；单查必填，批量用 Queries 时可不填</p>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>子类型；未知返回 0</p>
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// <p>筛选条件 JSON 字符串：[{Key,Op,Values}]，走分析立方白名单</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>指标名列表，如 count / rate；单查必填</p>
	Metrics []*string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// <p>分位/均值必填，如 lcp / duration</p>
	MeasurementKey *string `json:"MeasurementKey,omitnil,omitempty" name:"MeasurementKey"`

	// <p>批量查询，上限 32</p>
	Queries []*MetricQueryItem `json:"Queries,omitnil,omitempty" name:"Queries"`
}

type DescribeTDEMMetricSummaryRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>顶层事件类型；单查必填，批量用 Queries 时可不填</p>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>子类型；未知返回 0</p>
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// <p>筛选条件 JSON 字符串：[{Key,Op,Values}]，走分析立方白名单</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>指标名列表，如 count / rate；单查必填</p>
	Metrics []*string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// <p>分位/均值必填，如 lcp / duration</p>
	MeasurementKey *string `json:"MeasurementKey,omitnil,omitempty" name:"MeasurementKey"`

	// <p>批量查询，上限 32</p>
	Queries []*MetricQueryItem `json:"Queries,omitnil,omitempty" name:"Queries"`
}

func (r *DescribeTDEMMetricSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMMetricSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "EventType")
	delete(f, "SubType")
	delete(f, "Filters")
	delete(f, "Metrics")
	delete(f, "MeasurementKey")
	delete(f, "Queries")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMMetricSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMMetricSummaryResponseParams struct {
	// <p>单查指标。缺测指标不出现，禁止 null</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*FloatNameValue `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// <p>批量结果，Queries 非空时返回</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*MetricSummaryBatchItem `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMMetricSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMMetricSummaryResponseParams `json:"Response"`
}

func (r *DescribeTDEMMetricSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMMetricSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMMetricTrendRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>粒度：5m（默认）/ 1h / 1d；跨度大时服务端可加粗</p>
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// <p>顶层事件类型；单查必填，批量用 Queries 时可不填</p>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>子类型；未知返回 0</p>
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// <p>筛选条件 JSON 字符串：[{Key,Op,Values}]，走分析立方白名单</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>指标名列表，如 count / rate；单查必填</p>
	Metrics []*string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// <p>分位/均值必填，如 lcp / duration</p>
	MeasurementKey *string `json:"MeasurementKey,omitnil,omitempty" name:"MeasurementKey"`

	// <p>批量查询，上限 32</p>
	Queries []*MetricQueryItem `json:"Queries,omitnil,omitempty" name:"Queries"`

	// <p>对比窗开始 Unix 毫秒（仅 Trend 批量，须与 CompareEndTime 成对）</p>
	CompareStartTime *int64 `json:"CompareStartTime,omitnil,omitempty" name:"CompareStartTime"`

	// <p>对比窗结束 Unix 毫秒（仅 Trend 批量）</p>
	CompareEndTime *int64 `json:"CompareEndTime,omitnil,omitempty" name:"CompareEndTime"`
}

type DescribeTDEMMetricTrendRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>粒度：5m（默认）/ 1h / 1d；跨度大时服务端可加粗</p>
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// <p>顶层事件类型；单查必填，批量用 Queries 时可不填</p>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>子类型；未知返回 0</p>
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// <p>筛选条件 JSON 字符串：[{Key,Op,Values}]，走分析立方白名单</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>指标名列表，如 count / rate；单查必填</p>
	Metrics []*string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// <p>分位/均值必填，如 lcp / duration</p>
	MeasurementKey *string `json:"MeasurementKey,omitnil,omitempty" name:"MeasurementKey"`

	// <p>批量查询，上限 32</p>
	Queries []*MetricQueryItem `json:"Queries,omitnil,omitempty" name:"Queries"`

	// <p>对比窗开始 Unix 毫秒（仅 Trend 批量，须与 CompareEndTime 成对）</p>
	CompareStartTime *int64 `json:"CompareStartTime,omitnil,omitempty" name:"CompareStartTime"`

	// <p>对比窗结束 Unix 毫秒（仅 Trend 批量）</p>
	CompareEndTime *int64 `json:"CompareEndTime,omitnil,omitempty" name:"CompareEndTime"`
}

func (r *DescribeTDEMMetricTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMMetricTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Granularity")
	delete(f, "EventType")
	delete(f, "SubType")
	delete(f, "Filters")
	delete(f, "Metrics")
	delete(f, "MeasurementKey")
	delete(f, "Queries")
	delete(f, "CompareStartTime")
	delete(f, "CompareEndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMMetricTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMMetricTrendResponseParams struct {
	// <p>实际生效粒度</p>
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// <p>单查时序</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Series []*MetricTrendPoint `json:"Series,omitnil,omitempty" name:"Series"`

	// <p>批量结果，Queries 非空时返回</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*MetricTrendBatchItem `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMMetricTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMMetricTrendResponseParams `json:"Response"`
}

func (r *DescribeTDEMMetricTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMMetricTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMOverviewScoresRequestParams struct {
	// <p>应用 ID 列表，1-200（必填）</p>
	AppIds []*string `json:"AppIds,omitnil,omitempty" name:"AppIds"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。可选。总览总分按环境筛选时传入。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTDEMOverviewScoresRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID 列表，1-200（必填）</p>
	AppIds []*string `json:"AppIds,omitnil,omitempty" name:"AppIds"`

	// <p>开始时间 Unix 毫秒（必填）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。可选。总览总分按环境筛选时传入。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTDEMOverviewScoresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMOverviewScoresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppIds")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMOverviewScoresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMOverviewScoresResponseParams struct {
	// <p>各应用评分</p>
	Scores []*AppOverviewScore `json:"Scores,omitnil,omitempty" name:"Scores"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMOverviewScoresResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMOverviewScoresResponseParams `json:"Response"`
}

func (r *DescribeTDEMOverviewScoresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMOverviewScoresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMRelatedEventsRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填，用于分区裁剪）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>模式：timeline / page / breadcrumb（必填）</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>会话 ID（必填）</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>页面 URL（Mode=page 必填）</p>
	PageUrl *string `json:"PageUrl,omitnil,omitempty" name:"PageUrl"`

	// <p>锚点 Unix 毫秒（Mode=page/breadcrumb 使用）</p>
	AnchorTimestamp *int64 `json:"AnchorTimestamp,omitnil,omitempty" name:"AnchorTimestamp"`

	// <p>偏移，&gt;=0，默认 0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>条数，1-200，默认 200</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>事件分类（可选）</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>关键字（可选）</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>视图 ID</p>
	ViewId *string `json:"ViewId,omitnil,omitempty" name:"ViewId"`
}

type DescribeTDEMRelatedEventsRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间 Unix 毫秒（必填，用于分区裁剪）</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间 Unix 毫秒（必填）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>模式：timeline / page / breadcrumb（必填）</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>会话 ID（必填）</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>页面 URL（Mode=page 必填）</p>
	PageUrl *string `json:"PageUrl,omitnil,omitempty" name:"PageUrl"`

	// <p>锚点 Unix 毫秒（Mode=page/breadcrumb 使用）</p>
	AnchorTimestamp *int64 `json:"AnchorTimestamp,omitnil,omitempty" name:"AnchorTimestamp"`

	// <p>偏移，&gt;=0，默认 0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>条数，1-200，默认 200</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>事件分类（可选）</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>关键字（可选）</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>视图 ID</p>
	ViewId *string `json:"ViewId,omitnil,omitempty" name:"ViewId"`
}

func (r *DescribeTDEMRelatedEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMRelatedEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Mode")
	delete(f, "SessionId")
	delete(f, "PageUrl")
	delete(f, "AnchorTimestamp")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Category")
	delete(f, "Keyword")
	delete(f, "ViewId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMRelatedEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMRelatedEventsResponseParams struct {
	// <p>关联事件，schema 与 Search 列表项相同</p>
	Events []*ExplorationEventItem `json:"Events,omitnil,omitempty" name:"Events"`

	// <p>偏移</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>条数</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>是否还有下一页</p>
	HasMore *bool `json:"HasMore,omitnil,omitempty" name:"HasMore"`

	// <p>各分类计数，Key 为分类 id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryCounts []*IntKeyValue `json:"CategoryCounts,omitnil,omitempty" name:"CategoryCounts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMRelatedEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMRelatedEventsResponseParams `json:"Response"`
}

func (r *DescribeTDEMRelatedEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMRelatedEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSSEDimensionsRequestParams struct {
	// <p>应用 ID。这里是业务应用标识，不是腾讯云账号 AppId。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，Unix 毫秒。</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒。须通过现有时间范围校验。</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>指标。取值说明：token_throughput（token 吞吐率，token/秒）、first_message_delay（首次消息时延，毫秒）、error_rate（错误率，已乘 100 的百分比）、stall_rate（卡死率，已乘 100 的百分比）、fluency（流畅度，已乘 100 的百分比）。</p>
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// <p>多维分析维度。取值说明：region、app_version、platform、os、network_type、server_ip、isp、model_version。region 由 country 与 province 拼接，不属于趋势维度。</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。Op 只允许 in。Key 只允许 model_version、app_version、platform、os、region、is_error、is_stall、network_type、server_ip、isp、user_id。is_error 与 is_stall 的值只允许 true 或 false。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTDEMSSEDimensionsRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID。这里是业务应用标识，不是腾讯云账号 AppId。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，Unix 毫秒。</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒。须通过现有时间范围校验。</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>指标。取值说明：token_throughput（token 吞吐率，token/秒）、first_message_delay（首次消息时延，毫秒）、error_rate（错误率，已乘 100 的百分比）、stall_rate（卡死率，已乘 100 的百分比）、fluency（流畅度，已乘 100 的百分比）。</p>
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// <p>多维分析维度。取值说明：region、app_version、platform、os、network_type、server_ip、isp、model_version。region 由 country 与 province 拼接，不属于趋势维度。</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。Op 只允许 in。Key 只允许 model_version、app_version、platform、os、region、is_error、is_stall、network_type、server_ip、isp、user_id。is_error 与 is_stall 的值只允许 true 或 false。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTDEMSSEDimensionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSSEDimensionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Dimension")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMSSEDimensionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSSEDimensionsResponseParams struct {
	// <p>按当前指标值降序的维度行，最多 200 行。没有对话时为空数组。</p>
	Rows []*TDEMSSEDimensionRow `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMSSEDimensionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMSSEDimensionsResponseParams `json:"Response"`
}

func (r *DescribeTDEMSSEDimensionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSSEDimensionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSSEDistributionRequestParams struct {
	// <p>应用 ID。这里是业务应用标识，不是腾讯云账号 AppId。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，Unix 毫秒。</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒。须通过现有时间范围校验。</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>指标。只允许 token_throughput、first_message_delay、fluency。error_rate 与 stall_rate 返回 InvalidParameter。</p>
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。Op 只允许 in。Key 只允许 model_version、app_version、platform、os、region、is_error、is_stall、network_type、server_ip、isp、user_id。is_error 与 is_stall 的值只允许 true 或 false。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTDEMSSEDistributionRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID。这里是业务应用标识，不是腾讯云账号 AppId。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，Unix 毫秒。</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒。须通过现有时间范围校验。</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>指标。只允许 token_throughput、first_message_delay、fluency。error_rate 与 stall_rate 返回 InvalidParameter。</p>
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。Op 只允许 in。Key 只允许 model_version、app_version、platform、os、region、is_error、is_stall、network_type、server_ip、isp、user_id。is_error 与 is_stall 的值只允许 true 或 false。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTDEMSSEDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSSEDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMSSEDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSSEDistributionResponseParams struct {
	// <p>分布档。没有样本时为空数组。最后一档 Overflow 为 true 且省略 End。</p>
	Bins []*TDEMSSEDistributionBin `json:"Bins,omitnil,omitempty" name:"Bins"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMSSEDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMSSEDistributionResponseParams `json:"Response"`
}

func (r *DescribeTDEMSSEDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSSEDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSSEPercentileRequestParams struct {
	// <p>应用 ID。这里是业务应用标识，不是腾讯云账号 AppId。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，Unix 毫秒。</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒。须通过现有时间范围校验。</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>指标。只允许 token_throughput、first_message_delay、fluency。error_rate 与 stall_rate 返回 InvalidParameter。</p>
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。Op 只允许 in。Key 只允许 model_version、app_version、platform、os、region、is_error、is_stall、network_type、server_ip、isp、user_id。is_error 与 is_stall 的值只允许 true 或 false。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTDEMSSEPercentileRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID。这里是业务应用标识，不是腾讯云账号 AppId。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，Unix 毫秒。</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒。须通过现有时间范围校验。</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>指标。只允许 token_throughput、first_message_delay、fluency。error_rate 与 stall_rate 返回 InvalidParameter。</p>
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。Op 只允许 in。Key 只允许 model_version、app_version、platform、os、region、is_error、is_stall、network_type、server_ip、isp、user_id。is_error 与 is_stall 的值只允许 true 或 false。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTDEMSSEPercentileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSSEPercentileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMSSEPercentileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSSEPercentileResponseParams struct {
	// <p>时间桶宽度，单位秒，与趋势查询使用同一套分桶。</p>
	IntervalSec *int64 `json:"IntervalSec,omitnil,omitempty" name:"IntervalSec"`

	// <p>整个时间窗的算术平均。没有有效样本时为 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Average *float64 `json:"Average,omitnil,omitempty" name:"Average"`

	// <p>整个时间窗的 50 分位。没有有效样本时为 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	P50 *float64 `json:"P50,omitnil,omitempty" name:"P50"`

	// <p>整个时间窗的 75 分位。没有有效样本时为 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	P75 *float64 `json:"P75,omitnil,omitempty" name:"P75"`

	// <p>整个时间窗的 95 分位。没有有效样本时为 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	P95 *float64 `json:"P95,omitnil,omitempty" name:"P95"`

	// <p>整个时间窗的 99 分位。没有有效样本时为 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	P99 *float64 `json:"P99,omitnil,omitempty" name:"P99"`

	// <p>各时间桶的平均值和分位。没有样本的桶不返回，全部没有样本时为空数组。</p>
	Points []*TDEMSSEPercentilePoint `json:"Points,omitnil,omitempty" name:"Points"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMSSEPercentileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMSSEPercentileResponseParams `json:"Response"`
}

func (r *DescribeTDEMSSEPercentileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSSEPercentileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSSESummaryRequestParams struct {
	// <p>应用 ID。这里是业务应用标识，不是腾讯云账号 AppId。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，Unix 毫秒。</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒。须通过现有时间范围校验。</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。Op 只允许 in。Key 只允许 model_version、app_version、platform、os、region、is_error、is_stall、network_type、server_ip、isp、user_id。is_error 与 is_stall 的值只允许 true 或 false。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTDEMSSESummaryRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID。这里是业务应用标识，不是腾讯云账号 AppId。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，Unix 毫秒。</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒。须通过现有时间范围校验。</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。Op 只允许 in。Key 只允许 model_version、app_version、platform、os、region、is_error、is_stall、network_type、server_ip、isp、user_id。is_error 与 is_stall 的值只允许 true 或 false。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTDEMSSESummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSSESummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMSSESummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSSESummaryResponseParams struct {
	// <p>对话次数。一条大模型对话对应一条 sse_call。</p>
	ConversationCount *int64 `json:"ConversationCount,omitnil,omitempty" name:"ConversationCount"`

	// <p>token 总数。对话次数为 0 时为 0。</p>
	TotalTokenCount *int64 `json:"TotalTokenCount,omitnil,omitempty" name:"TotalTokenCount"`

	// <p>token 吞吐率，单位 token/秒。对话次数为 0 时为 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenThroughput *float64 `json:"TokenThroughput,omitnil,omitempty" name:"TokenThroughput"`

	// <p>首次消息时延，单位毫秒。对话次数为 0 时为 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstMessageDelay *float64 `json:"FirstMessageDelay,omitnil,omitempty" name:"FirstMessageDelay"`

	// <p>错误率，已乘 100。对话次数为 0 时为 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorRate *float64 `json:"ErrorRate,omitnil,omitempty" name:"ErrorRate"`

	// <p>卡死率，已乘 100。对话次数为 0 时为 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StallRate *float64 `json:"StallRate,omitnil,omitempty" name:"StallRate"`

	// <p>流畅度，已乘 100。对话次数为 0 时为 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fluency *float64 `json:"Fluency,omitnil,omitempty" name:"Fluency"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMSSESummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMSSESummaryResponseParams `json:"Response"`
}

func (r *DescribeTDEMSSESummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSSESummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSSETrendRequestParams struct {
	// <p>应用 ID。这里是业务应用标识，不是腾讯云账号 AppId。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，Unix 毫秒。</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒。须通过现有时间范围校验。</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>指标。取值说明：token_throughput（token 吞吐率，token/秒）、first_message_delay（首次消息时延，毫秒）、error_rate（错误率，已乘 100 的百分比）、stall_rate（卡死率，已乘 100 的百分比）、fluency（流畅度，已乘 100 的百分比）。</p>
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// <p>趋势维度。取值说明：overall、model_version、network_type、server_ip、isp。overall 返回 Points，其他维度返回 Series。</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。Op 只允许 in。Key 只允许 model_version、app_version、platform、os、region、is_error、is_stall、network_type、server_ip、isp、user_id。is_error 与 is_stall 的值只允许 true 或 false。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>头部维度数量，只允许 7、10、20。Dimension 为 overall 时忽略。</p>
	TopN *int64 `json:"TopN,omitnil,omitempty" name:"TopN"`
}

type DescribeTDEMSSETrendRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID。这里是业务应用标识，不是腾讯云账号 AppId。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间，Unix 毫秒。</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，Unix 毫秒。须通过现有时间范围校验。</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>指标。取值说明：token_throughput（token 吞吐率，token/秒）、first_message_delay（首次消息时延，毫秒）、error_rate（错误率，已乘 100 的百分比）、stall_rate（卡死率，已乘 100 的百分比）、fluency（流畅度，已乘 100 的百分比）。</p>
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// <p>趋势维度。取值说明：overall、model_version、network_type、server_ip、isp。overall 返回 Points，其他维度返回 Series。</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>筛选条件 JSON 字符串，格式为 [{Key,Op,Values}]。Op 只允许 in。Key 只允许 model_version、app_version、platform、os、region、is_error、is_stall、network_type、server_ip、isp、user_id。is_error 与 is_stall 的值只允许 true 或 false。</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>头部维度数量，只允许 7、10、20。Dimension 为 overall 时忽略。</p>
	TopN *int64 `json:"TopN,omitnil,omitempty" name:"TopN"`
}

func (r *DescribeTDEMSSETrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSSETrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Dimension")
	delete(f, "Filters")
	delete(f, "TopN")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMSSETrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSSETrendResponseParams struct {
	// <p>时间桶宽度，单位秒。</p>
	IntervalSec *int64 `json:"IntervalSec,omitnil,omitempty" name:"IntervalSec"`

	// <p>Dimension 为 overall 时的时间桶。没有对话时为空数组。</p>
	Points []*TDEMSSETrendPoint `json:"Points,omitnil,omitempty" name:"Points"`

	// <p>非 overall 时按指标从高到低取 TopN 后的序列。没有对话时为空数组。</p>
	Series []*TDEMSSETrendSeries `json:"Series,omitnil,omitempty" name:"Series"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMSSETrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMSSETrendResponseParams `json:"Response"`
}

func (r *DescribeTDEMSSETrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSSETrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSampleConfigRequestParams struct {
	// <p>项目ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DescribeTDEMSampleConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>项目ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *DescribeTDEMSampleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSampleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMSampleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSampleConfigResponseParams struct {
	// <p>业务数据 JSON 字符串，反序列化后：{ProjectId,Event,Replay,SnapshotId,UpdatedAt}。Event/Replay 结构见 ModifyTDEMSampleConfig 入参说明。无行时 SnapshotId=0、UpdatedAt 为空。</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMSampleConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMSampleConfigResponseParams `json:"Response"`
}

func (r *DescribeTDEMSampleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSampleRateRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DescribeTDEMSampleRateRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *DescribeTDEMSampleRateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSampleRateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMSampleRateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSampleRateResponseParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>采集采样率 0-100</p>
	SampleRate *float64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// <p>会话回放采样率 0-100；未配置时为 1</p>
	ReplaySampleRate *float64 `json:"ReplaySampleRate,omitnil,omitempty" name:"ReplaySampleRate"`

	// <p>创建时间</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>更新时间</p>
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMSampleRateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMSampleRateResponseParams `json:"Response"`
}

func (r *DescribeTDEMSampleRateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSampleRateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSessionReplayTimelineRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>会话id</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>数量限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>游标</p>
	Cursor *int64 `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// <p>时间线分页的事件ID游标，与时间游标Cursor共同定位续页起点，用于同一时间戳下的稳定分页。</p>
	CursorEventId *string `json:"CursorEventId,omitnil,omitempty" name:"CursorEventId"`
}

type DescribeTDEMSessionReplayTimelineRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>会话id</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>数量限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>游标</p>
	Cursor *int64 `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// <p>时间线分页的事件ID游标，与时间游标Cursor共同定位续页起点，用于同一时间戳下的稳定分页。</p>
	CursorEventId *string `json:"CursorEventId,omitnil,omitempty" name:"CursorEventId"`
}

func (r *DescribeTDEMSessionReplayTimelineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSessionReplayTimelineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "SessionId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Cursor")
	delete(f, "CursorEventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMSessionReplayTimelineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMSessionReplayTimelineResponseParams struct {
	// <p>请求结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMSessionReplayTimelineResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMSessionReplayTimelineResponseParams `json:"Response"`
}

func (r *DescribeTDEMSessionReplayTimelineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMSessionReplayTimelineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMUsageDetailRequestParams struct {
	// <p>开始时间，YYYY-MM-DD HH:00:00（小时对齐，必填）</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，YYYY-MM-DD HH:00:00（小时对齐，必填）</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>可选关键字，匹配业务系统名称、应用名称或应用 ID</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>分页偏移，默认 0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量，默认 20，上限 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>演示模式开关</p>
	DemoMode *bool `json:"DemoMode,omitnil,omitempty" name:"DemoMode"`
}

type DescribeTDEMUsageDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>开始时间，YYYY-MM-DD HH:00:00（小时对齐，必填）</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，YYYY-MM-DD HH:00:00（小时对齐，必填）</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>可选关键字，匹配业务系统名称、应用名称或应用 ID</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>分页偏移，默认 0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量，默认 20，上限 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>演示模式开关</p>
	DemoMode *bool `json:"DemoMode,omitnil,omitempty" name:"DemoMode"`
}

func (r *DescribeTDEMUsageDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMUsageDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Keyword")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DemoMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMUsageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMUsageDetailResponseParams struct {
	// <p>总行数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>当前结果集上报合计</p>
	ReportVolume *int64 `json:"ReportVolume,omitnil,omitempty" name:"ReportVolume"`

	// <p>明细</p>
	DetailSet []*UsageDetailRow `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMUsageDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMUsageDetailResponseParams `json:"Response"`
}

func (r *DescribeTDEMUsageDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMUsageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMUsageSummaryRequestParams struct {
	// <p>开始时间，YYYY-MM-DD HH:00:00（小时对齐，必填）</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，YYYY-MM-DD HH:00:00（小时对齐，必填）</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>演示模式开关</p>
	DemoMode *bool `json:"DemoMode,omitnil,omitempty" name:"DemoMode"`
}

type DescribeTDEMUsageSummaryRequest struct {
	*tchttp.BaseRequest
	
	// <p>开始时间，YYYY-MM-DD HH:00:00（小时对齐，必填）</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，YYYY-MM-DD HH:00:00（小时对齐，必填）</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>演示模式开关</p>
	DemoMode *bool `json:"DemoMode,omitnil,omitempty" name:"DemoMode"`
}

func (r *DescribeTDEMUsageSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMUsageSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DemoMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTDEMUsageSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTDEMUsageSummaryResponseParams struct {
	// <p>上报总量</p>
	ReportVolume *int64 `json:"ReportVolume,omitnil,omitempty" name:"ReportVolume"`

	// <p>窗口内有量应用数</p>
	ResourceCount *int64 `json:"ResourceCount,omitnil,omitempty" name:"ResourceCount"`

	// <p>hour 或 day</p>
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// <p>趋势</p>
	ByDate []*UsageTimePoint `json:"ByDate,omitnil,omitempty" name:"ByDate"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTDEMUsageSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTDEMUsageSummaryResponseParams `json:"Response"`
}

func (r *DescribeTDEMUsageSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTDEMUsageSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTawAreasRequestParams struct {
	// 片区Id
	AreaIds []*int64 `json:"AreaIds,omitnil,omitempty" name:"AreaIds"`

	// 片区Key
	AreaKeys []*string `json:"AreaKeys,omitnil,omitempty" name:"AreaKeys"`

	// 分页Limit，默认根据AreaKeys和AreaStatuses参数查询所有。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 片区状态(1=有效，2=无效)
	AreaStatuses []*int64 `json:"AreaStatuses,omitnil,omitempty" name:"AreaStatuses"`

	// 分页Offset，默认根据AreaKeys和AreaStatuses参数查询所有。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTawAreasRequest struct {
	*tchttp.BaseRequest
	
	// 片区Id
	AreaIds []*int64 `json:"AreaIds,omitnil,omitempty" name:"AreaIds"`

	// 片区Key
	AreaKeys []*string `json:"AreaKeys,omitnil,omitempty" name:"AreaKeys"`

	// 分页Limit，默认根据AreaKeys和AreaStatuses参数查询所有。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 片区状态(1=有效，2=无效)
	AreaStatuses []*int64 `json:"AreaStatuses,omitnil,omitempty" name:"AreaStatuses"`

	// 分页Offset，默认根据AreaKeys和AreaStatuses参数查询所有。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTawAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTawAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AreaIds")
	delete(f, "AreaKeys")
	delete(f, "Limit")
	delete(f, "AreaStatuses")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTawAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTawAreasResponseParams struct {
	// 片区总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 片区列表
	AreaSet []*RumAreaInfo `json:"AreaSet,omitnil,omitempty" name:"AreaSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTawAreasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTawAreasResponseParams `json:"Response"`
}

func (r *DescribeTawAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTawAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTawInstancesRequestParams struct {
	// 计费状态
	ChargeStatuses []*int64 `json:"ChargeStatuses,omitnil,omitempty" name:"ChargeStatuses"`

	// 计费类型
	ChargeTypes []*int64 `json:"ChargeTypes,omitnil,omitempty" name:"ChargeTypes"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 片区Id
	AreaIds []*int64 `json:"AreaIds,omitnil,omitempty" name:"AreaIds"`

	// 实例状态(1=创建中，2=运行中，3=异常，4=重启中，5=停止中，6=已停止，7=销毁中，8=已销毁), 该参数已废弃，请在Filters内注明
	//
	// Deprecated: InstanceStatuses is deprecated.
	InstanceStatuses []*int64 `json:"InstanceStatuses,omitnil,omitempty" name:"InstanceStatuses"`

	// 实例Id, 该参数已废弃，请在Filters内注明
	//
	// Deprecated: InstanceIds is deprecated.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 过滤参数；demo模式传{"Name": "IsDemo", "Values":["1"]}
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 该参数已废弃，demo模式请在Filters内注明
	//
	// Deprecated: IsDemo is deprecated.
	IsDemo *int64 `json:"IsDemo,omitnil,omitempty" name:"IsDemo"`
}

type DescribeTawInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 计费状态
	ChargeStatuses []*int64 `json:"ChargeStatuses,omitnil,omitempty" name:"ChargeStatuses"`

	// 计费类型
	ChargeTypes []*int64 `json:"ChargeTypes,omitnil,omitempty" name:"ChargeTypes"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 片区Id
	AreaIds []*int64 `json:"AreaIds,omitnil,omitempty" name:"AreaIds"`

	// 实例状态(1=创建中，2=运行中，3=异常，4=重启中，5=停止中，6=已停止，7=销毁中，8=已销毁), 该参数已废弃，请在Filters内注明
	InstanceStatuses []*int64 `json:"InstanceStatuses,omitnil,omitempty" name:"InstanceStatuses"`

	// 实例Id, 该参数已废弃，请在Filters内注明
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 过滤参数；demo模式传{"Name": "IsDemo", "Values":["1"]}
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 该参数已废弃，demo模式请在Filters内注明
	IsDemo *int64 `json:"IsDemo,omitnil,omitempty" name:"IsDemo"`
}

func (r *DescribeTawInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTawInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChargeStatuses")
	delete(f, "ChargeTypes")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "AreaIds")
	delete(f, "InstanceStatuses")
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "IsDemo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTawInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTawInstancesResponseParams struct {
	// 实例列表
	InstanceSet []*RumInstanceInfo `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTawInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTawInstancesResponseParams `json:"Response"`
}

func (r *DescribeTawInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTawInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenRequestParams struct {
	// <p>筛选条件</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>仅对比模式下填写，筛选条件A</p>
	FormListAString *string `json:"FormListAString,omitnil,omitempty" name:"FormListAString"`

	// <p>仅对比模式下填写，筛选条件B</p>
	FormListBString *string `json:"FormListBString,omitnil,omitempty" name:"FormListBString"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`

	// <p>拓展字段</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`
}

type DescribeTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>筛选条件</p>
	FormListString *string `json:"FormListString,omitnil,omitempty" name:"FormListString"`

	// <p>仅对比模式下填写，筛选条件A</p>
	FormListAString *string `json:"FormListAString,omitnil,omitempty" name:"FormListAString"`

	// <p>仅对比模式下填写，筛选条件B</p>
	FormListBString *string `json:"FormListBString,omitnil,omitempty" name:"FormListBString"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`

	// <p>拓展字段</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`
}

func (r *DescribeTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FormListString")
	delete(f, "FormListAString")
	delete(f, "FormListBString")
	delete(f, "RequestHeader")
	delete(f, "ExtraData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenResponseParams `json:"Response"`
}

func (r *DescribeTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopIssuesRequestParams struct {
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>需要对比的查询条件，没有则不填</p>
	Compare *CompareCondition `json:"Compare,omitnil,omitempty" name:"Compare"`

	// <p>查询条件</p>
	Condition *CompareCondition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// <p>问题类型</p>
	IssueType *int64 `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序方式</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>topN</p>
	TopNum *int64 `json:"TopNum,omitnil,omitempty" name:"TopNum"`

	// <p>拓展字段</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

type DescribeTopIssuesRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品Id</p>
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// <p>需要对比的查询条件，没有则不填</p>
	Compare *CompareCondition `json:"Compare,omitnil,omitempty" name:"Compare"`

	// <p>查询条件</p>
	Condition *CompareCondition `json:"Condition,omitnil,omitempty" name:"Condition"`

	// <p>问题类型</p>
	IssueType *int64 `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// <p>排序字段</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>排序方式</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>topN</p>
	TopNum *int64 `json:"TopNum,omitnil,omitempty" name:"TopNum"`

	// <p>拓展字段</p>
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// <p>请求头</p>
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`
}

func (r *DescribeTopIssuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopIssuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Compare")
	delete(f, "Condition")
	delete(f, "IssueType")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "TopNum")
	delete(f, "ExtraData")
	delete(f, "RequestHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopIssuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopIssuesResponseParams struct {
	// <p>返回值</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>状态码</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopIssuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopIssuesResponseParams `json:"Response"`
}

func (r *DescribeTopIssuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopIssuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserJourneySanKeyRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>过滤条件</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>每一步骤的数量</p>
	TopPagesPerStep *int64 `json:"TopPagesPerStep,omitnil,omitempty" name:"TopPagesPerStep"`

	// <p>每一个节点最小的会话数量</p>
	MinSessionThreshold *int64 `json:"MinSessionThreshold,omitnil,omitempty" name:"MinSessionThreshold"`

	// <p>开始页面</p>
	StartPage *string `json:"StartPage,omitnil,omitempty" name:"StartPage"`

	// <p>结束页面</p>
	EndPage *string `json:"EndPage,omitnil,omitempty" name:"EndPage"`

	// <p>是否有挣扎事件</p>
	HasStruggle *bool `json:"HasStruggle,omitnil,omitempty" name:"HasStruggle"`
}

type DescribeUserJourneySanKeyRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>过滤条件</p>
	Filters *string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>每一步骤的数量</p>
	TopPagesPerStep *int64 `json:"TopPagesPerStep,omitnil,omitempty" name:"TopPagesPerStep"`

	// <p>每一个节点最小的会话数量</p>
	MinSessionThreshold *int64 `json:"MinSessionThreshold,omitnil,omitempty" name:"MinSessionThreshold"`

	// <p>开始页面</p>
	StartPage *string `json:"StartPage,omitnil,omitempty" name:"StartPage"`

	// <p>结束页面</p>
	EndPage *string `json:"EndPage,omitnil,omitempty" name:"EndPage"`

	// <p>是否有挣扎事件</p>
	HasStruggle *bool `json:"HasStruggle,omitnil,omitempty" name:"HasStruggle"`
}

func (r *DescribeUserJourneySanKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserJourneySanKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Platform")
	delete(f, "Filters")
	delete(f, "TopPagesPerStep")
	delete(f, "MinSessionThreshold")
	delete(f, "StartPage")
	delete(f, "EndPage")
	delete(f, "HasStruggle")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserJourneySanKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserJourneySanKeyResponseParams struct {
	// <p>请求结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserJourneySanKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserJourneySanKeyResponseParams `json:"Response"`
}

func (r *DescribeUserJourneySanKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserJourneySanKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUvListRequestParams struct {
	// ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取day：d，   min:m
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`
}

type DescribeUvListRequest struct {
	*tchttp.BaseRequest
	
	// ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 获取day：d，   min:m
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`
}

func (r *DescribeUvListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUvListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "Dimension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUvListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUvListResponseParams struct {
	// uv列表
	ProjectUvSet []*RumUvInfo `json:"ProjectUvSet,omitnil,omitempty" name:"ProjectUvSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUvListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUvListResponseParams `json:"Response"`
}

func (r *DescribeUvListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUvListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhitelistsRequestParams struct {
	// 实例instance-ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeWhitelistsRequest struct {
	*tchttp.BaseRequest
	
	// 实例instance-ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

func (r *DescribeWhitelistsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhitelistsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhitelistsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhitelistsResponseParams struct {
	// 白名单列表
	WhitelistSet []*Whitelist `json:"WhitelistSet,omitnil,omitempty" name:"WhitelistSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWhitelistsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhitelistsResponseParams `json:"Response"`
}

func (r *DescribeWhitelistsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhitelistsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DialDimensionItem struct {
	// <p>维度值。空值在控制台展示为未知。</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// <p>执行次数。target_ip 按探测条数，其它维度按批次。</p>
	ExecutionCount *int64 `json:"ExecutionCount,omitnil,omitempty" name:"ExecutionCount"`

	// <p>成功次数</p>
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// <p>失败次数</p>
	FailCount *int64 `json:"FailCount,omitnil,omitempty" name:"FailCount"`

	// <p>平均时延，单位毫秒</p>
	AvgDuration *float64 `json:"AvgDuration,omitnil,omitempty" name:"AvgDuration"`

	// <p>平均丢包率，0 到 100</p>
	AvgLoss *float64 `json:"AvgLoss,omitnil,omitempty" name:"AvgLoss"`
}

type DialExecution struct {
	// <p>任务 ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务名称。任务已删除时仍返回当时名称，找不到则为空。</p>
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>拨测批次 ID</p>
	TaskSessionId *string `json:"TaskSessionId,omitnil,omitempty" name:"TaskSessionId"`

	// <p>批次开始时间，毫秒时间戳</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>客户端平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备机型</p>
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>网络类型</p>
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// <p>客户端运营商。不是目标 IP 运营商。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>客户端地区。不是目标 IP 地区。</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>本批次探测到的目标 IP，顺序与 TargetIpGeo 一致</p>
	TargetIps []*string `json:"TargetIps,omitnil,omitempty" name:"TargetIps"`

	// <p>每个目标 IP 的地区和运营商</p>
	TargetIpGeo []*DialTargetIpGeo `json:"TargetIpGeo,omitnil,omitempty" name:"TargetIpGeo"`

	// <p>本批次是否全部探测成功</p>
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`

	// <p>本批次是否未实际探测</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotProbed *bool `json:"NotProbed,omitnil,omitempty" name:"NotProbed"`

	// <p>本批次成功目标的平均时延，单位毫秒</p>
	AvgDuration *float64 `json:"AvgDuration,omitnil,omitempty" name:"AvgDuration"`

	// <p>本批次是否包含路由追踪</p>
	HasMTR *bool `json:"HasMTR,omitnil,omitempty" name:"HasMTR"`

	// <p>所属用户会话 ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DialProbe struct {
	// <p>探测目标 IP</p>
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// <p>本次探测是否成功</p>
	PingSuccess *bool `json:"PingSuccess,omitnil,omitempty" name:"PingSuccess"`

	// <p>发送包数</p>
	Sent *float64 `json:"Sent,omitnil,omitempty" name:"Sent"`

	// <p>接收包数</p>
	Received *float64 `json:"Received,omitnil,omitempty" name:"Received"`

	// <p>丢包率，0 到 100</p>
	Loss *float64 `json:"Loss,omitnil,omitempty" name:"Loss"`

	// <p>最小时延，单位毫秒</p>
	Min *float64 `json:"Min,omitnil,omitempty" name:"Min"`

	// <p>最大时延，单位毫秒</p>
	Max *float64 `json:"Max,omitnil,omitempty" name:"Max"`

	// <p>平均时延，单位毫秒</p>
	Avg *float64 `json:"Avg,omitnil,omitempty" name:"Avg"`

	// <p>探测协议</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>探测目标域名或地址</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>未探测原因</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkipReason *string `json:"SkipReason,omitnil,omitempty" name:"SkipReason"`

	// <p>目标 IP 所在地区。与客户端地区无关。无法解析时为未知。</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>目标 IP 所属运营商。与客户端运营商无关。无法解析时为未知。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type DialTargetIpGeo struct {
	// <p>目标 IP，与 TargetIps 中的地址一致</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>目标 IP 所在地区。无法解析时为未知。</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>目标 IP 所属运营商。无法解析时为未知。</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type DialTask struct {
	// <p>任务 ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务名称</p>
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>任务状态。取值说明：active、paused；已删除任务不返回。</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>探测类型。取值说明：icmp、tcp。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>执行时机。取值说明：realtime、after_request、after_error、after_slow、after_dns_slow。</p>
	ExecutionTime *string `json:"ExecutionTime,omitnil,omitempty" name:"ExecutionTime"`

	// <p>监听域名。非 realtime 时有值。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatchDomain *string `json:"WatchDomain,omitnil,omitempty" name:"WatchDomain"`

	// <p>拨测目标。域名、IP 或 ldns。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DialDomain *string `json:"DialDomain,omitnil,omitempty" name:"DialDomain"`

	// <p>TCP 端口。仅 tcp 任务返回。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>实时拨测间隔，单位毫秒。仅 realtime 返回。取值说明：60000、300000、600000、900000、1800000、3600000。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrequencyMs *int64 `json:"FrequencyMs,omitnil,omitempty" name:"FrequencyMs"`

	// <p>单次探测超时，单位毫秒。范围 1000 到 20000。</p>
	TimeoutMs *int64 `json:"TimeoutMs,omitnil,omitempty" name:"TimeoutMs"`

	// <p>探测周期，单位毫秒。</p>
	PeriodMs *int64 `json:"PeriodMs,omitnil,omitempty" name:"PeriodMs"`

	// <p>发包次数。范围 3 到 40。</p>
	PackageCount *int64 `json:"PackageCount,omitnil,omitempty" name:"PackageCount"`

	// <p>ICMP 包大小，单位字节。范围 32 到 1024，默认 64。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageSize *int64 `json:"PackageSize,omitnil,omitempty" name:"PackageSize"`

	// <p>是否做路由追踪。取值说明：0 否，1 是。</p>
	NeedMTR *int64 `json:"NeedMTR,omitnil,omitempty" name:"NeedMTR"`

	// <p>路由追踪超时，单位毫秒。NeedMTR 为 1 时有值。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MTRTimeoutMs *int64 `json:"MTRTimeoutMs,omitnil,omitempty" name:"MTRTimeoutMs"`

	// <p>路由追踪最大跳数。NeedMTR 为 1 时有值。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MTRMaxHop *int64 `json:"MTRMaxHop,omitnil,omitempty" name:"MTRMaxHop"`

	// <p>每一跳发包次数。范围 3 到 40。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MTRPingCount *int64 `json:"MTRPingCount,omitnil,omitempty" name:"MTRPingCount"`

	// <p>慢请求阈值，单位毫秒。after_slow、after_dns_slow 时使用，默认 2000。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlowThresholdMs *int64 `json:"SlowThresholdMs,omitnil,omitempty" name:"SlowThresholdMs"`

	// <p>错误类型过滤。after_error 时使用。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorType *string `json:"ErrorType,omitnil,omitempty" name:"ErrorType"`

	// <p>允许启动的客户端地区，逗号分隔，名称与分析页客户端地区一致。空表示不限制。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regions *string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// <p>允许启动的客户端运营商，逗号分隔，名称与分析页客户端运营商一致。空表示不限制。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operators *string `json:"Operators,omitnil,omitempty" name:"Operators"`

	// <p>允许启动的用户 ID，逗号分隔。空表示不限制。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserIds *string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// <p>允许启动的设备 ID，逗号分隔。空表示不限制。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceIds *string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`

	// <p>创建人</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`

	// <p>创建时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>更新时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type DialTrendPoint struct {
	// <p>时间桶，格式 yyyy-MM-dd HH:mm:ss</p>
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// <p>执行批次数</p>
	ExecutionCount *int64 `json:"ExecutionCount,omitnil,omitempty" name:"ExecutionCount"`

	// <p>成功批次数</p>
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// <p>失败批次数</p>
	FailCount *int64 `json:"FailCount,omitnil,omitempty" name:"FailCount"`

	// <p>平均时延，单位毫秒</p>
	AvgDuration *float64 `json:"AvgDuration,omitnil,omitempty" name:"AvgDuration"`
}

type EventFrame struct {
	// <p>绝对路径</p>
	AbsPath *string `json:"AbsPath,omitnil,omitempty" name:"AbsPath"`

	// <p>列号</p>
	ColNo *int64 `json:"ColNo,omitnil,omitempty" name:"ColNo"`

	// <p>文件名</p>
	Filename *string `json:"Filename,omitnil,omitempty" name:"Filename"`

	// <p>函数名</p>
	Function *string `json:"Function,omitnil,omitempty" name:"Function"`

	// <p>行号</p>
	LineNo *int64 `json:"LineNo,omitnil,omitempty" name:"LineNo"`

	// <p>是否应用内 0/1</p>
	InApp *int64 `json:"InApp,omitnil,omitempty" name:"InApp"`

	// <p>模块</p>
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// <p>包名</p>
	Package *string `json:"Package,omitnil,omitempty" name:"Package"`

	// <p>栈层级</p>
	StackLevel *int64 `json:"StackLevel,omitnil,omitempty" name:"StackLevel"`
}

type EventRankItem struct {
	// <p>错误指纹</p>
	Fingerprint *string `json:"Fingerprint,omitnil,omitempty" name:"Fingerprint"`

	// <p>事件类型</p>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>崩溃类型</p>
	CrashType *string `json:"CrashType,omitnil,omitempty" name:"CrashType"`

	// <p>错误类型</p>
	ErrorType *string `json:"ErrorType,omitnil,omitempty" name:"ErrorType"`

	// <p>错误值</p>
	ErrorValue *string `json:"ErrorValue,omitnil,omitempty" name:"ErrorValue"`

	// <p>问题单 ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>问题单状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>API 或资源 URL</p>
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// <p>次数</p>
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// <p>影响用户数</p>
	AffectedUser *int64 `json:"AffectedUser,omitnil,omitempty" name:"AffectedUser"`

	// <p>失败次数</p>
	Fail *int64 `json:"Fail,omitnil,omitempty" name:"Fail"`

	// <p>耗时</p>
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>连接时长</p>
	ConnectTime *float64 `json:"ConnectTime,omitnil,omitempty" name:"ConnectTime"`

	// <p>DNS 时长</p>
	DomainLookupTime *float64 `json:"DomainLookupTime,omitnil,omitempty" name:"DomainLookupTime"`

	// <p>时序</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Series []*EventRankPoint `json:"Series,omitnil,omitempty" name:"Series"`
}

type EventRankPoint struct {
	// <p>时间戳 Unix 毫秒</p>
	Ts *int64 `json:"Ts,omitnil,omitempty" name:"Ts"`

	// <p>次数</p>
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// <p>失败次数</p>
	Fail *int64 `json:"Fail,omitnil,omitempty" name:"Fail"`

	// <p>耗时</p>
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

type EventSearchSeriesPoint struct {
	// <p>时间戳 Unix 毫秒</p>
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// <p>次数</p>
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type EventStack struct {
	// <p>异常类型</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>异常信息</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// <p>mechanism 类型</p>
	MechanismType *string `json:"MechanismType,omitnil,omitempty" name:"MechanismType"`

	// <p>是否已处理 0/1</p>
	MechanismHandled *int64 `json:"MechanismHandled,omitnil,omitempty" name:"MechanismHandled"`
}

type EventView struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>事件 ID</p>
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// <p>视图 ID</p>
	ViewId *string `json:"ViewId,omitnil,omitempty" name:"ViewId"`

	// <p>事件类型</p>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>事件分类</p>
	EventCategory *string `json:"EventCategory,omitnil,omitempty" name:"EventCategory"`

	// <p>事件名</p>
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// <p>时间 RFC3339</p>
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// <p>分区日期 YYYY-MM-DD</p>
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>SDK 版本</p>
	SdkVersion *string `json:"SdkVersion,omitnil,omitempty" name:"SdkVersion"`

	// <p>问题单 ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>错误指纹</p>
	Fingerprint *string `json:"Fingerprint,omitnil,omitempty" name:"Fingerprint"`

	// <p>是否已符号化 0/1</p>
	IsSourceMapped *int64 `json:"IsSourceMapped,omitnil,omitempty" name:"IsSourceMapped"`

	// <p>异常栈</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptionStacks []*EventStack `json:"ExceptionStacks,omitnil,omitempty" name:"ExceptionStacks"`

	// <p>原始堆栈帧</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptionFrames []*EventFrame `json:"ExceptionFrames,omitnil,omitempty" name:"ExceptionFrames"`

	// <p>事件属性</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventProperties []*StringKeyValue `json:"EventProperties,omitnil,omitempty" name:"EventProperties"`

	// <p>会话 ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>业务用户 ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>设备 ID</p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>页面 URL</p>
	PageUrl *string `json:"PageUrl,omitnil,omitempty" name:"PageUrl"`

	// <p>归一化页面 URL</p>
	NormalizedPageUrl *string `json:"NormalizedPageUrl,omitnil,omitempty" name:"NormalizedPageUrl"`

	// <p>页面标题</p>
	PageTitle *string `json:"PageTitle,omitnil,omitempty" name:"PageTitle"`

	// <p>资源 URL</p>
	ResourceUrl *string `json:"ResourceUrl,omitnil,omitempty" name:"ResourceUrl"`

	// <p>资源类型</p>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>目标元素</p>
	TargetElement *string `json:"TargetElement,omitnil,omitempty" name:"TargetElement"`

	// <p>目标文本</p>
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`

	// <p>日志级别</p>
	LogLevel *string `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// <p>IP</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>省份</p>
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// <p>城市</p>
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// <p>运营商</p>
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// <p>网络类型</p>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// <p>设备类型</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>操作系统</p>
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// <p>浏览器</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>环境</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`

	// <p>测量值</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Measurements []*FloatKeyValue `json:"Measurements,omitnil,omitempty" name:"Measurements"`

	// <p>标签</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*StringKeyValue `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ExceptionBucket struct {
	// <p>固定目录分类 key，与请求平台、CategoryKeys 或 EventTypes 派生目录一致。</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>精确事件行数，0至9007199254740991的整数，0必须保留。</p>
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// <p>是否为计入 TotalCount 的互斥主项。包含项为false，不重复计总。</p>
	IncludedInTotal *bool `json:"IncludedInTotal,omitnil,omitempty" name:"IncludedInTotal"`

	// <p>包含项所属主项的 Key；主项不返回该字段。</p>
	ParentKey *string `json:"ParentKey,omitnil,omitempty" name:"ParentKey"`
}

type ExceptionDimensionRow struct {
	// <p>由目录域、维度或TagKey及归一化值确定的稳定、不透明唯一标识。</p>
	RowKey *string `json:"RowKey,omitnil,omitempty" name:"RowKey"`

	// <p>维度显示值；标签保留原始值，普通province/platform/network_type按既有显示别名归一化。</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// <p>本行全部 IncludedInTotal=true 主项的 Count 之和，安全非负整数。</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>完整固定分类向量，含0计数及适用包含项；不按数量裁剪。</p>
	Items []*ExceptionBucket `json:"Items,omitnil,omitempty" name:"Items"`
}

type ExceptionDistribution struct {
	// <p>按 TotalCount 降序、归一化值升序返回的完整维度行，行数不超过Limit；无数据返回空数组。</p>
	Rows []*ExceptionDimensionRow `json:"Rows,omitnil,omitempty" name:"Rows"`
}

type ExplorationEventItem struct {
	// <p>事件 ID</p>
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// <p>事件类型</p>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>事件名</p>
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// <p>事件分类</p>
	EventCategory *string `json:"EventCategory,omitnil,omitempty" name:"EventCategory"`

	// <p>时间戳 Unix 毫秒</p>
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// <p>会话 ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>视图 ID</p>
	ViewId *string `json:"ViewId,omitnil,omitempty" name:"ViewId"`

	// <p>业务用户 ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>设备 ID</p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>页面 URL</p>
	PageUrl *string `json:"PageUrl,omitnil,omitempty" name:"PageUrl"`

	// <p>归一化页面 URL</p>
	NormalizedPageUrl *string `json:"NormalizedPageUrl,omitnil,omitempty" name:"NormalizedPageUrl"`

	// <p>页面标题</p>
	PageTitle *string `json:"PageTitle,omitnil,omitempty" name:"PageTitle"`

	// <p>平台</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>操作系统</p>
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// <p>浏览器</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>设备类型</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>省份</p>
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// <p>城市</p>
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// <p>运营商</p>
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// <p>网络类型</p>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// <p>IP</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>资源 URL</p>
	ResourceUrl *string `json:"ResourceUrl,omitnil,omitempty" name:"ResourceUrl"`

	// <p>资源类型</p>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>目标元素</p>
	TargetElement *string `json:"TargetElement,omitnil,omitempty" name:"TargetElement"`

	// <p>目标文本</p>
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`

	// <p>日志级别</p>
	LogLevel *string `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// <p>问题单 ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>是否已符号化 0/1</p>
	IsSourceMapped *int64 `json:"IsSourceMapped,omitnil,omitempty" name:"IsSourceMapped"`

	// <p>测量值</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Measurements []*FloatKeyValue `json:"Measurements,omitnil,omitempty" name:"Measurements"`

	// <p>事件属性</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventProperties []*StringKeyValue `json:"EventProperties,omitnil,omitempty" name:"EventProperties"`

	// <p>异常栈摘要</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptionStacks []*ExplorationExceptionStack `json:"ExceptionStacks,omitnil,omitempty" name:"ExceptionStacks"`

	// <p>事件标签列表。APM 全链路时包含 Key=trace_id，Value 为 32 位 hex TraceId（不是整段 traceparent）。无标签时省略或空数组。</p>
	Tags []*StringKeyValue `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ExplorationExceptionStack struct {
	// <p>异常类型</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>异常信息</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// <p>mechanism 类型</p>
	MechanismType *string `json:"MechanismType,omitnil,omitempty" name:"MechanismType"`

	// <p>是否已处理 0/1</p>
	MechanismHandled *int64 `json:"MechanismHandled,omitnil,omitempty" name:"MechanismHandled"`
}

type FieldTransformRuleItem struct {
	// <p>规则 ID</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>规则名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>目标字段</p>
	TargetField *string `json:"TargetField,omitnil,omitempty" name:"TargetField"`

	// <p>匹配方式：exact / prefix / contains / regex</p>
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// <p>匹配模式</p>
	MatchPattern *string `json:"MatchPattern,omitnil,omitempty" name:"MatchPattern"`

	// <p>替换动作：replace / regex_replace / mask</p>
	TransformAction *string `json:"TransformAction,omitnil,omitempty" name:"TransformAction"`

	// <p>替换值</p>
	ReplaceValue *string `json:"ReplaceValue,omitnil,omitempty" name:"ReplaceValue"`

	// <p>优先级</p>
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// <p>是否启用</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>创建人</p>
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`

	// <p>更新人</p>
	UpdatedBy *string `json:"UpdatedBy,omitnil,omitempty" name:"UpdatedBy"`

	// <p>创建时间</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>更新时间</p>
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type Filter struct {
	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 过滤键的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type Filters struct {
	// 是否反转
	IsReversed *bool `json:"IsReversed,omitnil,omitempty" name:"IsReversed"`

	// 键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 运算符
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FloatKeyValue struct {
	// <p>键</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>值</p>
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type FloatNameValue struct {
	// <p>指标名，如 Count / Rate / P95</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>指标值</p>
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type FunnelStepInput struct {
	// <p>应用id</p>
	StepName *string `json:"StepName,omitnil,omitempty" name:"StepName"`

	// <p>页面特征</p>
	PagePattern *string `json:"PagePattern,omitnil,omitempty" name:"PagePattern"`

	// <p>页面匹配模式</p>
	PageMatchMode *string `json:"PageMatchMode,omitnil,omitempty" name:"PageMatchMode"`

	// <p>控件Selector</p>
	Selector *string `json:"Selector,omitnil,omitempty" name:"Selector"`

	// <p>事件类型</p>
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// <p>可视化名称，如控件文本、自定义事件名</p>
	VisualName *string `json:"VisualName,omitnil,omitempty" name:"VisualName"`
}

type IntKeyValue struct {
	// <p>键，如 error / session</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>整型值</p>
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type Kafka struct {
	// 1：开启
	// 0：关闭
	EnableKafka *uint64 `json:"EnableKafka,omitnil,omitempty" name:"EnableKafka"`

	// host地址
	KafkaHost *string `json:"KafkaHost,omitnil,omitempty" name:"KafkaHost"`

	// topic
	KafkaTopic *string `json:"KafkaTopic,omitnil,omitempty" name:"KafkaTopic"`

	// 版本
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// username
	SaslUserName *string `json:"SaslUserName,omitnil,omitempty" name:"SaslUserName"`

	// password
	SaslPassword *string `json:"SaslPassword,omitnil,omitempty" name:"SaslPassword"`

	// ssl
	SaslMechanism *string `json:"SaslMechanism,omitnil,omitempty" name:"SaslMechanism"`

	// 默认算子id为0新增算子
	// 一旦算子新增成功会返回正确的算子id值
	SinkId *uint64 `json:"SinkId,omitnil,omitempty" name:"SinkId"`
}

type MetricDistributionBatchItem struct {
	// <p>条目 Key</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>分布桶</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*MetricDistributionItem `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>合计</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type MetricDistributionItem struct {
	// <p>维度值</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// <p>次数</p>
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// <p>影响用户数</p>
	AffectedUser *int64 `json:"AffectedUser,omitnil,omitempty" name:"AffectedUser"`

	// <p>率，分母为 0 时省略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rate *float64 `json:"Rate,omitnil,omitempty" name:"Rate"`

	// <p>影响用户率，缺省省略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AffectedUserRate *float64 `json:"AffectedUserRate,omitnil,omitempty" name:"AffectedUserRate"`

	// <p>均值，缺省省略</p>
	Avg *float64 `json:"Avg,omitnil,omitempty" name:"Avg"`

	// <p>P50，缺省省略</p>
	P50 *float64 `json:"P50,omitnil,omitempty" name:"P50"`

	// <p>P75，缺省省略</p>
	P75 *float64 `json:"P75,omitnil,omitempty" name:"P75"`

	// <p>P95，缺省省略</p>
	P95 *float64 `json:"P95,omitnil,omitempty" name:"P95"`

	// <p>P99，缺省省略</p>
	P99 *float64 `json:"P99,omitnil,omitempty" name:"P99"`
}

type MetricQueryItem struct {
	// <p>事件类型</p>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>指标名列表，如 count / rate</p>
	Metrics []*string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// <p>条目 Key，缺省为 EventType/SubType/MeasurementKey</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>子类型</p>
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// <p>分位/均值测量键</p>
	MeasurementKey *string `json:"MeasurementKey,omitnil,omitempty" name:"MeasurementKey"`
}

type MetricSummaryBatchItem struct {
	// <p>条目 Key</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>指标列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*FloatNameValue `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// <p>可能尚未采集，false 时省略</p>
	MaybeUncollected *bool `json:"MaybeUncollected,omitnil,omitempty" name:"MaybeUncollected"`
}

type MetricTrendBatchItem struct {
	// <p>条目 Key</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>主窗时序</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Series []*MetricTrendPoint `json:"Series,omitnil,omitempty" name:"Series"`

	// <p>对比窗时序</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareSeries []*MetricTrendPoint `json:"CompareSeries,omitnil,omitempty" name:"CompareSeries"`

	// <p>可能尚未采集，false 时省略</p>
	MaybeUncollected *bool `json:"MaybeUncollected,omitnil,omitempty" name:"MaybeUncollected"`
}

type MetricTrendPoint struct {
	// <p>时间戳 Unix 毫秒</p>
	Ts *int64 `json:"Ts,omitnil,omitempty" name:"Ts"`

	// <p>该桶指标</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*FloatNameValue `json:"Metrics,omitnil,omitempty" name:"Metrics"`
}

// Predefined struct for user
type ModifyFieldTransformRuleRequestParams struct {
	// <p>规则id</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>规则名</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>目标字段</p>
	TargetField *string `json:"TargetField,omitnil,omitempty" name:"TargetField"`

	// <p>匹配方式</p>
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// <p>匹配规则</p>
	MatchPattern *string `json:"MatchPattern,omitnil,omitempty" name:"MatchPattern"`

	// <p>转换方式</p>
	TransformAction *string `json:"TransformAction,omitnil,omitempty" name:"TransformAction"`

	// <p>替换值</p>
	ReplaceValue *string `json:"ReplaceValue,omitnil,omitempty" name:"ReplaceValue"`

	// <p>规则优先级</p>
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// <p>是否开启</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ModifyFieldTransformRuleRequest struct {
	*tchttp.BaseRequest
	
	// <p>规则id</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>规则名</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>目标字段</p>
	TargetField *string `json:"TargetField,omitnil,omitempty" name:"TargetField"`

	// <p>匹配方式</p>
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// <p>匹配规则</p>
	MatchPattern *string `json:"MatchPattern,omitnil,omitempty" name:"MatchPattern"`

	// <p>转换方式</p>
	TransformAction *string `json:"TransformAction,omitnil,omitempty" name:"TransformAction"`

	// <p>替换值</p>
	ReplaceValue *string `json:"ReplaceValue,omitnil,omitempty" name:"ReplaceValue"`

	// <p>规则优先级</p>
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// <p>是否开启</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *ModifyFieldTransformRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFieldTransformRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "TargetField")
	delete(f, "MatchType")
	delete(f, "MatchPattern")
	delete(f, "TransformAction")
	delete(f, "ReplaceValue")
	delete(f, "Priority")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFieldTransformRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFieldTransformRuleResponseParams struct {
	// <p>结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFieldTransformRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFieldTransformRuleResponseParams `json:"Response"`
}

func (r *ModifyFieldTransformRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFieldTransformRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// 要修改的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新的实例名称(长度最大不超过255)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 新的实例描述(长度最大不超过1024)
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 要修改的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新的实例名称(长度最大不超过255)
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 新的实例描述(长度最大不超过1024)
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "InstanceDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIssueStatusRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Issue ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>状态动作，可选值：resolve、ignore、reopen</p>
	IssueAction *string `json:"IssueAction,omitnil,omitempty" name:"IssueAction"`

	// <p>解决版本；当 IssueAction=resolve 时可传，不传则默认使用 latest_release</p>
	ResolvedRelease *string `json:"ResolvedRelease,omitnil,omitempty" name:"ResolvedRelease"`
}

type ModifyIssueStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Issue ID</p>
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// <p>状态动作，可选值：resolve、ignore、reopen</p>
	IssueAction *string `json:"IssueAction,omitnil,omitempty" name:"IssueAction"`

	// <p>解决版本；当 IssueAction=resolve 时可传，不传则默认使用 latest_release</p>
	ResolvedRelease *string `json:"ResolvedRelease,omitnil,omitempty" name:"ResolvedRelease"`
}

func (r *ModifyIssueStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIssueStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "IssueId")
	delete(f, "IssueAction")
	delete(f, "ResolvedRelease")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIssueStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIssueStatusResponseParams struct {
	// <p>业务结果 JSON 字符串；前端会再做一次 JSON 解析。Result 内业务 JSON 为 Data，内容是状态流转结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息；正常返回时为空</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIssueStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIssueStatusResponseParams `json:"Response"`
}

func (r *ModifyIssueStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIssueStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMiniProgramReplayPackRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>回放包id（前缀mpr_）</p>
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`

	// <p>应用版本（可选）</p>
	Release *string `json:"Release,omitnil,omitempty" name:"Release"`
}

type ModifyMiniProgramReplayPackRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>回放包id（前缀mpr_）</p>
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`

	// <p>应用版本（可选）</p>
	Release *string `json:"Release,omitnil,omitempty" name:"Release"`
}

func (r *ModifyMiniProgramReplayPackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMiniProgramReplayPackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "PackId")
	delete(f, "Release")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMiniProgramReplayPackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMiniProgramReplayPackResponseParams struct {
	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// <p>请求结果（JSON字符串，包含 PackId/AppId/BundleHash/Release/Status）</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMiniProgramReplayPackResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMiniProgramReplayPackResponseParams `json:"Response"`
}

func (r *ModifyMiniProgramReplayPackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMiniProgramReplayPackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectLimitRequestParams struct {
	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 取值为[log speed performance webvitals pv event custom miniProgramData]其中之一
	ProjectInterface *string `json:"ProjectInterface,omitnil,omitempty" name:"ProjectInterface"`

	// 上报比例   10代表10%
	ReportRate *int64 `json:"ReportRate,omitnil,omitempty" name:"ReportRate"`

	// 上报类型 1：比例  2：上报量
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 主键ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type ModifyProjectLimitRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// 取值为[log speed performance webvitals pv event custom miniProgramData]其中之一
	ProjectInterface *string `json:"ProjectInterface,omitnil,omitempty" name:"ProjectInterface"`

	// 上报比例   10代表10%
	ReportRate *int64 `json:"ReportRate,omitnil,omitempty" name:"ReportRate"`

	// 上报类型 1：比例  2：上报量
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 主键ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *ModifyProjectLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectID")
	delete(f, "ProjectInterface")
	delete(f, "ReportRate")
	delete(f, "ReportType")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectLimitResponseParams struct {
	// 返回信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProjectLimitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectLimitResponseParams `json:"Response"`
}

func (r *ModifyProjectLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectRequestParams struct {
	// 项目 id
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 应用名称(可选，不为空且最长为 200字符)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目网页地址(可选，最长为 256)
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 项目仓库地址(可选，最长为 256)
	Repo *string `json:"Repo,omitnil,omitempty" name:"Repo"`

	// 项目需要转移到的实例 id(可选)
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目采样率(可选)
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 是否开启聚类(可选)
	EnableURLGroup *uint64 `json:"EnableURLGroup,omitnil,omitempty" name:"EnableURLGroup"`

	// 项目类型(可接受值为 "web", "mp", "android", "ios", "node", "hippy", "weex", "viola", "rn")
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 应用描述(可选，最长为 1000字符)
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 启动kafka配置
	EnableKafka *uint64 `json:"EnableKafka,omitnil,omitempty" name:"EnableKafka"`

	// kafka_host
	KafkaHost *string `json:"KafkaHost,omitnil,omitempty" name:"KafkaHost"`

	// topic
	KafkaTopic *string `json:"KafkaTopic,omitnil,omitempty" name:"KafkaTopic"`

	// kafka_version
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// kafka_username
	SaslUserName *string `json:"SaslUserName,omitnil,omitempty" name:"SaslUserName"`

	// kafka_pwd
	SaslPassword *string `json:"SaslPassword,omitnil,omitempty" name:"SaslPassword"`

	// SaslMechanism
	SaslMechanism *string `json:"SaslMechanism,omitnil,omitempty" name:"SaslMechanism"`

	// sink_id，日知汇算子id
	SinkId *uint64 `json:"SinkId,omitnil,omitempty" name:"SinkId"`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目 id
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 应用名称(可选，不为空且最长为 200字符)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目网页地址(可选，最长为 256)
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 项目仓库地址(可选，最长为 256)
	Repo *string `json:"Repo,omitnil,omitempty" name:"Repo"`

	// 项目需要转移到的实例 id(可选)
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目采样率(可选)
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 是否开启聚类(可选)
	EnableURLGroup *uint64 `json:"EnableURLGroup,omitnil,omitempty" name:"EnableURLGroup"`

	// 项目类型(可接受值为 "web", "mp", "android", "ios", "node", "hippy", "weex", "viola", "rn")
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 应用描述(可选，最长为 1000字符)
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 启动kafka配置
	EnableKafka *uint64 `json:"EnableKafka,omitnil,omitempty" name:"EnableKafka"`

	// kafka_host
	KafkaHost *string `json:"KafkaHost,omitnil,omitempty" name:"KafkaHost"`

	// topic
	KafkaTopic *string `json:"KafkaTopic,omitnil,omitempty" name:"KafkaTopic"`

	// kafka_version
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`

	// kafka_username
	SaslUserName *string `json:"SaslUserName,omitnil,omitempty" name:"SaslUserName"`

	// kafka_pwd
	SaslPassword *string `json:"SaslPassword,omitnil,omitempty" name:"SaslPassword"`

	// SaslMechanism
	SaslMechanism *string `json:"SaslMechanism,omitnil,omitempty" name:"SaslMechanism"`

	// sink_id，日知汇算子id
	SinkId *uint64 `json:"SinkId,omitnil,omitempty" name:"SinkId"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "Name")
	delete(f, "URL")
	delete(f, "Repo")
	delete(f, "InstanceID")
	delete(f, "Rate")
	delete(f, "EnableURLGroup")
	delete(f, "Type")
	delete(f, "Desc")
	delete(f, "EnableKafka")
	delete(f, "KafkaHost")
	delete(f, "KafkaTopic")
	delete(f, "KafkaVersion")
	delete(f, "SaslUserName")
	delete(f, "SaslPassword")
	delete(f, "SaslMechanism")
	delete(f, "SinkId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// 操作信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 项目id
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectResponseParams `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySampleRateRequestParams struct {
	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>会话采样率</p>
	SampleRate *float64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// <p>录制采样率</p>
	ReplaySampleRate *float64 `json:"ReplaySampleRate,omitnil,omitempty" name:"ReplaySampleRate"`
}

type ModifySampleRateRequest struct {
	*tchttp.BaseRequest
	
	// <p>项目id</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>会话采样率</p>
	SampleRate *float64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// <p>录制采样率</p>
	ReplaySampleRate *float64 `json:"ReplaySampleRate,omitnil,omitempty" name:"ReplaySampleRate"`
}

func (r *ModifySampleRateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySampleRateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SampleRate")
	delete(f, "ReplaySampleRate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySampleRateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySampleRateResponseParams struct {
	// <p>返回结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySampleRateResponse struct {
	*tchttp.BaseResponse
	Response *ModifySampleRateResponseParams `json:"Response"`
}

func (r *ModifySampleRateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySampleRateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMApplicationRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用名称（可选）</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>应用描述（可选）</p>
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// <p>是否排除已忽略错误</p>
	ExcludeIgnoredErrors *bool `json:"ExcludeIgnoredErrors,omitnil,omitempty" name:"ExcludeIgnoredErrors"`

	// <p>APM 实例 ID 列表</p>
	ApmInstanceIds []*string `json:"ApmInstanceIds,omitnil,omitempty" name:"ApmInstanceIds"`
}

type ModifyTDEMApplicationRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用名称（可选）</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>应用描述（可选）</p>
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// <p>是否排除已忽略错误</p>
	ExcludeIgnoredErrors *bool `json:"ExcludeIgnoredErrors,omitnil,omitempty" name:"ExcludeIgnoredErrors"`

	// <p>APM 实例 ID 列表</p>
	ApmInstanceIds []*string `json:"ApmInstanceIds,omitnil,omitempty" name:"ApmInstanceIds"`
}

func (r *ModifyTDEMApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Name")
	delete(f, "Desc")
	delete(f, "ExcludeIgnoredErrors")
	delete(f, "ApmInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTDEMApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMApplicationResponseParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTDEMApplicationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTDEMApplicationResponseParams `json:"Response"`
}

func (r *ModifyTDEMApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMBusinessSystemRequestParams struct {
	// <p>业务系统 ID（必填）</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>业务系统名称（可选）</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>业务系统描述（可选）</p>
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`
}

type ModifyTDEMBusinessSystemRequest struct {
	*tchttp.BaseRequest
	
	// <p>业务系统 ID（必填）</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>业务系统名称（可选）</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>业务系统描述（可选）</p>
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`
}

func (r *ModifyTDEMBusinessSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMBusinessSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "InstanceDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTDEMBusinessSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMBusinessSystemResponseParams struct {
	// <p>业务系统 ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTDEMBusinessSystemResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTDEMBusinessSystemResponseParams `json:"Response"`
}

func (r *ModifyTDEMBusinessSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMBusinessSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMDialTaskRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>任务 ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务名称。不传表示不修改。</p>
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>任务状态。取值说明：active、paused。不传表示不修改。不能改为 deleted。</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>探测类型。创建后不可修改。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>执行时机。创建后不可修改。</p>
	ExecutionTime *string `json:"ExecutionTime,omitnil,omitempty" name:"ExecutionTime"`

	// <p>监听域名。不传表示不修改。</p>
	WatchDomain *string `json:"WatchDomain,omitnil,omitempty" name:"WatchDomain"`

	// <p>拨测目标。不传表示不修改。</p>
	DialDomain *string `json:"DialDomain,omitnil,omitempty" name:"DialDomain"`

	// <p>TCP 端口。不传表示不修改。</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>实时拨测间隔，单位毫秒。不传表示不修改。</p>
	FrequencyMs *int64 `json:"FrequencyMs,omitnil,omitempty" name:"FrequencyMs"`

	// <p>单次探测超时，单位毫秒。不传表示不修改。</p>
	TimeoutMs *int64 `json:"TimeoutMs,omitnil,omitempty" name:"TimeoutMs"`

	// <p>探测周期，单位毫秒。不传表示不修改。</p>
	PeriodMs *int64 `json:"PeriodMs,omitnil,omitempty" name:"PeriodMs"`

	// <p>发包次数。不传表示不修改。</p>
	PackageCount *int64 `json:"PackageCount,omitnil,omitempty" name:"PackageCount"`

	// <p>ICMP 包大小，单位字节。不传表示不修改。</p>
	PackageSize *int64 `json:"PackageSize,omitnil,omitempty" name:"PackageSize"`

	// <p>是否做路由追踪。取值说明：0 否，1 是。不传表示不修改。</p>
	NeedMTR *int64 `json:"NeedMTR,omitnil,omitempty" name:"NeedMTR"`

	// <p>路由追踪超时，单位毫秒。不传表示不修改。</p>
	MTRTimeoutMs *int64 `json:"MTRTimeoutMs,omitnil,omitempty" name:"MTRTimeoutMs"`

	// <p>路由追踪最大跳数。不传表示不修改。</p>
	MTRMaxHop *int64 `json:"MTRMaxHop,omitnil,omitempty" name:"MTRMaxHop"`

	// <p>每一跳发包次数。不传表示不修改。</p>
	MTRPingCount *int64 `json:"MTRPingCount,omitnil,omitempty" name:"MTRPingCount"`

	// <p>慢请求阈值，单位毫秒。不传表示不修改。</p>
	SlowThresholdMs *int64 `json:"SlowThresholdMs,omitnil,omitempty" name:"SlowThresholdMs"`

	// <p>错误类型过滤。不传表示不修改。</p>
	ErrorType *string `json:"ErrorType,omitnil,omitempty" name:"ErrorType"`

	// <p>允许启动的客户端地区，逗号分隔，名称必须与分析页客户端地区完全一致。不传保留原值，空字符串清空。空名单表示不限制。</p>
	Regions *string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// <p>允许启动的客户端运营商，逗号分隔，名称必须与分析页客户端运营商完全一致。不传保留原值，空字符串清空。空名单表示不限制。</p>
	Operators *string `json:"Operators,omitnil,omitempty" name:"Operators"`

	// <p>允许启动的用户 ID，逗号分隔。不传保留原值，空字符串清空。空名单表示不限制。</p>
	UserIds *string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// <p>允许启动的设备 ID，逗号分隔。不传保留原值，空字符串清空。空名单表示不限制。</p>
	DeviceIds *string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`
}

type ModifyTDEMDialTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>任务 ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务名称。不传表示不修改。</p>
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>任务状态。取值说明：active、paused。不传表示不修改。不能改为 deleted。</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>探测类型。创建后不可修改。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>执行时机。创建后不可修改。</p>
	ExecutionTime *string `json:"ExecutionTime,omitnil,omitempty" name:"ExecutionTime"`

	// <p>监听域名。不传表示不修改。</p>
	WatchDomain *string `json:"WatchDomain,omitnil,omitempty" name:"WatchDomain"`

	// <p>拨测目标。不传表示不修改。</p>
	DialDomain *string `json:"DialDomain,omitnil,omitempty" name:"DialDomain"`

	// <p>TCP 端口。不传表示不修改。</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>实时拨测间隔，单位毫秒。不传表示不修改。</p>
	FrequencyMs *int64 `json:"FrequencyMs,omitnil,omitempty" name:"FrequencyMs"`

	// <p>单次探测超时，单位毫秒。不传表示不修改。</p>
	TimeoutMs *int64 `json:"TimeoutMs,omitnil,omitempty" name:"TimeoutMs"`

	// <p>探测周期，单位毫秒。不传表示不修改。</p>
	PeriodMs *int64 `json:"PeriodMs,omitnil,omitempty" name:"PeriodMs"`

	// <p>发包次数。不传表示不修改。</p>
	PackageCount *int64 `json:"PackageCount,omitnil,omitempty" name:"PackageCount"`

	// <p>ICMP 包大小，单位字节。不传表示不修改。</p>
	PackageSize *int64 `json:"PackageSize,omitnil,omitempty" name:"PackageSize"`

	// <p>是否做路由追踪。取值说明：0 否，1 是。不传表示不修改。</p>
	NeedMTR *int64 `json:"NeedMTR,omitnil,omitempty" name:"NeedMTR"`

	// <p>路由追踪超时，单位毫秒。不传表示不修改。</p>
	MTRTimeoutMs *int64 `json:"MTRTimeoutMs,omitnil,omitempty" name:"MTRTimeoutMs"`

	// <p>路由追踪最大跳数。不传表示不修改。</p>
	MTRMaxHop *int64 `json:"MTRMaxHop,omitnil,omitempty" name:"MTRMaxHop"`

	// <p>每一跳发包次数。不传表示不修改。</p>
	MTRPingCount *int64 `json:"MTRPingCount,omitnil,omitempty" name:"MTRPingCount"`

	// <p>慢请求阈值，单位毫秒。不传表示不修改。</p>
	SlowThresholdMs *int64 `json:"SlowThresholdMs,omitnil,omitempty" name:"SlowThresholdMs"`

	// <p>错误类型过滤。不传表示不修改。</p>
	ErrorType *string `json:"ErrorType,omitnil,omitempty" name:"ErrorType"`

	// <p>允许启动的客户端地区，逗号分隔，名称必须与分析页客户端地区完全一致。不传保留原值，空字符串清空。空名单表示不限制。</p>
	Regions *string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// <p>允许启动的客户端运营商，逗号分隔，名称必须与分析页客户端运营商完全一致。不传保留原值，空字符串清空。空名单表示不限制。</p>
	Operators *string `json:"Operators,omitnil,omitempty" name:"Operators"`

	// <p>允许启动的用户 ID，逗号分隔。不传保留原值，空字符串清空。空名单表示不限制。</p>
	UserIds *string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// <p>允许启动的设备 ID，逗号分隔。不传保留原值，空字符串清空。空名单表示不限制。</p>
	DeviceIds *string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`
}

func (r *ModifyTDEMDialTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMDialTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "TaskId")
	delete(f, "TaskName")
	delete(f, "Status")
	delete(f, "Type")
	delete(f, "ExecutionTime")
	delete(f, "WatchDomain")
	delete(f, "DialDomain")
	delete(f, "Port")
	delete(f, "FrequencyMs")
	delete(f, "TimeoutMs")
	delete(f, "PeriodMs")
	delete(f, "PackageCount")
	delete(f, "PackageSize")
	delete(f, "NeedMTR")
	delete(f, "MTRTimeoutMs")
	delete(f, "MTRMaxHop")
	delete(f, "MTRPingCount")
	delete(f, "SlowThresholdMs")
	delete(f, "ErrorType")
	delete(f, "Regions")
	delete(f, "Operators")
	delete(f, "UserIds")
	delete(f, "DeviceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTDEMDialTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMDialTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTDEMDialTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTDEMDialTaskResponseParams `json:"Response"`
}

func (r *ModifyTDEMDialTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMDialTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMFieldTransformRuleRequestParams struct {
	// <p>规则 ID（必填）</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>目标字段（必填）</p>
	TargetField *string `json:"TargetField,omitnil,omitempty" name:"TargetField"`

	// <p>匹配方式（必填）：exact / prefix / contains / regex</p>
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// <p>匹配模式（必填）</p>
	MatchPattern *string `json:"MatchPattern,omitnil,omitempty" name:"MatchPattern"`

	// <p>替换动作（必填）：replace / regex_replace / mask</p>
	TransformAction *string `json:"TransformAction,omitnil,omitempty" name:"TransformAction"`

	// <p>规则名称（可选）</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>替换值</p>
	ReplaceValue *string `json:"ReplaceValue,omitnil,omitempty" name:"ReplaceValue"`

	// <p>优先级</p>
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// <p>是否启用</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ModifyTDEMFieldTransformRuleRequest struct {
	*tchttp.BaseRequest
	
	// <p>规则 ID（必填）</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>目标字段（必填）</p>
	TargetField *string `json:"TargetField,omitnil,omitempty" name:"TargetField"`

	// <p>匹配方式（必填）：exact / prefix / contains / regex</p>
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// <p>匹配模式（必填）</p>
	MatchPattern *string `json:"MatchPattern,omitnil,omitempty" name:"MatchPattern"`

	// <p>替换动作（必填）：replace / regex_replace / mask</p>
	TransformAction *string `json:"TransformAction,omitnil,omitempty" name:"TransformAction"`

	// <p>规则名称（可选）</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>替换值</p>
	ReplaceValue *string `json:"ReplaceValue,omitnil,omitempty" name:"ReplaceValue"`

	// <p>优先级</p>
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// <p>是否启用</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *ModifyTDEMFieldTransformRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMFieldTransformRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "AppId")
	delete(f, "TargetField")
	delete(f, "MatchType")
	delete(f, "MatchPattern")
	delete(f, "TransformAction")
	delete(f, "Name")
	delete(f, "ReplaceValue")
	delete(f, "Priority")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTDEMFieldTransformRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMFieldTransformRuleResponseParams struct {
	// <p>修改后的规则</p>
	Rule *FieldTransformRuleItem `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTDEMFieldTransformRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTDEMFieldTransformRuleResponseParams `json:"Response"`
}

func (r *ModifyTDEMFieldTransformRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMFieldTransformRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMFunnelRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`

	// <p>漏斗名</p>
	FunnelName *string `json:"FunnelName,omitnil,omitempty" name:"FunnelName"`

	// <p>漏斗步骤</p>
	Steps []*FunnelStepInput `json:"Steps,omitnil,omitempty" name:"Steps"`

	// <p>漏斗描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>漏斗状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>漏斗来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

type ModifyTDEMFunnelRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>漏斗id</p>
	FunnelId *string `json:"FunnelId,omitnil,omitempty" name:"FunnelId"`

	// <p>漏斗名</p>
	FunnelName *string `json:"FunnelName,omitnil,omitempty" name:"FunnelName"`

	// <p>漏斗步骤</p>
	Steps []*FunnelStepInput `json:"Steps,omitnil,omitempty" name:"Steps"`

	// <p>漏斗描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>漏斗状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>漏斗来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

func (r *ModifyTDEMFunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMFunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FunnelId")
	delete(f, "FunnelName")
	delete(f, "Steps")
	delete(f, "Description")
	delete(f, "Status")
	delete(f, "Source")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTDEMFunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMFunnelResponseParams struct {
	// <p>请求结果</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTDEMFunnelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTDEMFunnelResponseParams `json:"Response"`
}

func (r *ModifyTDEMFunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMFunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMSampleConfigRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>事件采样配置 JSON 字符串。反序列化后：{"Enabled":bool,"Types":{&lt;event_type&gt;:{"SampleRate":0-100最多两位小数,"Rules":[{"Name":string,"Enabled":bool,"SampleRate":0-100,"Conditions":[{"Field":string,"Operator":string,"Value":string或string[]}]}]}}}。Types 可只含改过的 type；未出现的 type 服务端视为 100% + 空规则。Types key 闭集：js_error,promise_error,resource_error,api_error,blank_screen,no_response,page_load,web_vitals,lag_detected,memory_usage,memory_oom,resource_load,api_call,sse_call,click,page_view,tab_hidden,tab_visible,form_focus,form_blur,form_change,form_submit,form_submit_result,struggle,custom,log（不含 rage_click/lcp）。Operator 白名单：equals,not_equals,contains,not_contains,in,not_in,starts_with,ends_with。规则 Conditions 至少 1 条；in/not_in 的 Value 为字符串数组。</p>
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// <p>回放采样配置 JSON 字符串。反序列化后：{"Enabled":bool,"ExperienceSampleRate":0-100最多两位小数,"ErrorTraceback":bool,"ErrorSampleRate":0-100最多两位小数}。控制台/远程字段名为 ExperienceSampleRate（非 SessionSampleRate）。</p>
	Replay *string `json:"Replay,omitnil,omitempty" name:"Replay"`
}

type ModifyTDEMSampleConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>事件采样配置 JSON 字符串。反序列化后：{"Enabled":bool,"Types":{&lt;event_type&gt;:{"SampleRate":0-100最多两位小数,"Rules":[{"Name":string,"Enabled":bool,"SampleRate":0-100,"Conditions":[{"Field":string,"Operator":string,"Value":string或string[]}]}]}}}。Types 可只含改过的 type；未出现的 type 服务端视为 100% + 空规则。Types key 闭集：js_error,promise_error,resource_error,api_error,blank_screen,no_response,page_load,web_vitals,lag_detected,memory_usage,memory_oom,resource_load,api_call,sse_call,click,page_view,tab_hidden,tab_visible,form_focus,form_blur,form_change,form_submit,form_submit_result,struggle,custom,log（不含 rage_click/lcp）。Operator 白名单：equals,not_equals,contains,not_contains,in,not_in,starts_with,ends_with。规则 Conditions 至少 1 条；in/not_in 的 Value 为字符串数组。</p>
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// <p>回放采样配置 JSON 字符串。反序列化后：{"Enabled":bool,"ExperienceSampleRate":0-100最多两位小数,"ErrorTraceback":bool,"ErrorSampleRate":0-100最多两位小数}。控制台/远程字段名为 ExperienceSampleRate（非 SessionSampleRate）。</p>
	Replay *string `json:"Replay,omitnil,omitempty" name:"Replay"`
}

func (r *ModifyTDEMSampleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMSampleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Event")
	delete(f, "Replay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTDEMSampleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMSampleConfigResponseParams struct {
	// <p>业务数据 JSON 字符串，反序列化后：{ProjectId,Event,Replay,SnapshotId,UpdatedAt}。SnapshotId 为本次写入的快照主键。</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTDEMSampleConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTDEMSampleConfigResponseParams `json:"Response"`
}

func (r *ModifyTDEMSampleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMSampleRateRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>采集采样率 0-100（必填）</p>
	SampleRate *float64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// <p>会话回放采样率 0-100（可选）</p>
	ReplaySampleRate *float64 `json:"ReplaySampleRate,omitnil,omitempty" name:"ReplaySampleRate"`
}

type ModifyTDEMSampleRateRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>采集采样率 0-100（必填）</p>
	SampleRate *float64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// <p>会话回放采样率 0-100（可选）</p>
	ReplaySampleRate *float64 `json:"ReplaySampleRate,omitnil,omitempty" name:"ReplaySampleRate"`
}

func (r *ModifyTDEMSampleRateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMSampleRateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "SampleRate")
	delete(f, "ReplaySampleRate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTDEMSampleRateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTDEMSampleRateResponseParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>采集采样率 0-100</p>
	SampleRate *float64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// <p>会话回放采样率 0-100</p>
	ReplaySampleRate *float64 `json:"ReplaySampleRate,omitnil,omitempty" name:"ReplaySampleRate"`

	// <p>更新时间</p>
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTDEMSampleRateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTDEMSampleRateResponseParams `json:"Response"`
}

func (r *ModifyTDEMSampleRateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTDEMSampleRateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverviewScoreItem struct {
	// <p>计分项，如 crash / anr / lag</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>单项 0-100，缺失省略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>原始值（率为小数、耗时为 ms）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type PreviewTDEMFunnelReportRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间戳（毫秒）</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间戳（毫秒）</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>漏斗步骤配置JSON字符串（2~20步），格式：[{"StepName":"首页","PagePattern":"/","PageMatchMode":"exact","Selector":"","ActionType":"page_view","VisualName":""}]</p>
	Steps *string `json:"Steps,omitnil,omitempty" name:"Steps"`

	// <p>平台筛选</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本筛选</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备类型筛选</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>浏览器筛选</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>操作系统筛选</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>国家/地区筛选</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>来源筛选</p>
	Referrer *string `json:"Referrer,omitnil,omitempty" name:"Referrer"`

	// <p>环境选择</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

type PreviewTDEMFunnelReportRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>开始时间戳（毫秒）</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间戳（毫秒）</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>漏斗步骤配置JSON字符串（2~20步），格式：[{"StepName":"首页","PagePattern":"/","PageMatchMode":"exact","Selector":"","ActionType":"page_view","VisualName":""}]</p>
	Steps *string `json:"Steps,omitnil,omitempty" name:"Steps"`

	// <p>平台筛选</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>应用版本筛选</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>设备类型筛选</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>浏览器筛选</p>
	Browser *string `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>操作系统筛选</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>国家/地区筛选</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>来源筛选</p>
	Referrer *string `json:"Referrer,omitnil,omitempty" name:"Referrer"`

	// <p>环境选择</p>
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

func (r *PreviewTDEMFunnelReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewTDEMFunnelReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Steps")
	delete(f, "Platform")
	delete(f, "AppVersion")
	delete(f, "DeviceType")
	delete(f, "Browser")
	delete(f, "OS")
	delete(f, "Country")
	delete(f, "Referrer")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PreviewTDEMFunnelReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PreviewTDEMFunnelReportResponseParams struct {
	// <p>请求结果（JSON字符串，包含 FunnelId/FunnelName/Summary/Steps/Chart）</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PreviewTDEMFunnelReportResponse struct {
	*tchttp.BaseResponse
	Response *PreviewTDEMFunnelReportResponseParams `json:"Response"`
}

func (r *PreviewTDEMFunnelReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewTDEMFunnelReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProjectLimit struct {
	// 接口
	ProjectInterface *string `json:"ProjectInterface,omitnil,omitempty" name:"ProjectInterface"`

	// 上报率
	ReportRate *int64 `json:"ReportRate,omitnil,omitempty" name:"ReportRate"`

	// 上报类型 1：上报率  2：上报量限制
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 主键ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`
}

type ReleaseFile struct {
	// <p>文件版本</p>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// <p>文件唯一 key</p>
	FileKey *string `json:"FileKey,omitnil,omitempty" name:"FileKey"`

	// <p>文件名</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>文件哈希值</p>
	FileHash *string `json:"FileHash,omitnil,omitempty" name:"FileHash"`

	// <p>文件 id</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>创建时间</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`
}

// Predefined struct for user
type ResumeInstanceRequestParams struct {
	// 需要恢复的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改是否包括白名单
	IsModifyAll *bool `json:"IsModifyAll,omitnil,omitempty" name:"IsModifyAll"`
}

type ResumeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 需要恢复的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改是否包括白名单
	IsModifyAll *bool `json:"IsModifyAll,omitnil,omitempty" name:"IsModifyAll"`
}

func (r *ResumeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IsModifyAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ResumeInstanceResponseParams `json:"Response"`
}

func (r *ResumeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeProjectRequestParams struct {
	// 项目 id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ResumeProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目 id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *ResumeProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeProjectResponse struct {
	*tchttp.BaseResponse
	Response *ResumeProjectResponseParams `json:"Response"`
}

func (r *ResumeProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeTDEMApplicationRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type ResumeTDEMApplicationRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *ResumeTDEMApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeTDEMApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeTDEMApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeTDEMApplicationResponseParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>恢复后为 2</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeTDEMApplicationResponse struct {
	*tchttp.BaseResponse
	Response *ResumeTDEMApplicationResponseParams `json:"Response"`
}

func (r *ResumeTDEMApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeTDEMApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RumAreaInfo struct {
	// 片区Id
	AreaId *int64 `json:"AreaId,omitnil,omitempty" name:"AreaId"`

	// 片区状态(1=有效，2=无效)
	AreaStatus *int64 `json:"AreaStatus,omitnil,omitempty" name:"AreaStatus"`

	// 片区名称
	AreaName *string `json:"AreaName,omitnil,omitempty" name:"AreaName"`

	// 片区Key
	AreaKey *string `json:"AreaKey,omitnil,omitempty" name:"AreaKey"`

	// 地域码表 id
	AreaRegionID *string `json:"AreaRegionID,omitnil,omitempty" name:"AreaRegionID"`

	// 地域码表 code 如 ap-xxx（xxx 为地域词）
	AreaRegionCode *string `json:"AreaRegionCode,omitnil,omitempty" name:"AreaRegionCode"`

	// 地域缩写
	AreaAbbr *string `json:"AreaAbbr,omitnil,omitempty" name:"AreaAbbr"`
}

type RumInstanceInfo struct {
	// 实例状态(1=创建中，2=运行中，3=异常，4=重启中，5=停止中，6=已停止，7=已删除)
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 片区Id
	AreaId *int64 `json:"AreaId,omitnil,omitempty" name:"AreaId"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群Id
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例描述
	InstanceDesc *string `json:"InstanceDesc,omitnil,omitempty" name:"InstanceDesc"`

	// 计费状态(1=使用中，2=已过期，3=已销毁，4=分配中，5=分配失败)
	ChargeStatus *int64 `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// 计费类型(1=免费版，2=预付费，3=后付费)
	ChargeType *int64 `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 数据保留时间(天)
	DataRetentionDays *int64 `json:"DataRetentionDays,omitnil,omitempty" name:"DataRetentionDays"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 实例类型 1:原web相关类型 2:app端类型
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type RumProject struct {
	// 项目名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 创建者 id
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 实例 id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 项目类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 项目仓库地址
	Repo *string `json:"Repo,omitnil,omitempty" name:"Repo"`

	// 项目网址地址
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 项目采样频率
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 项目唯一key（长度 12 位）
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 是否开启url聚类
	EnableURLGroup *int64 `json:"EnableURLGroup,omitnil,omitempty" name:"EnableURLGroup"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 项目 ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 实例 key
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// 项目描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 是否星标  1:是 0:否
	IsStar *int64 `json:"IsStar,omitnil,omitempty" name:"IsStar"`

	// 项目状态(1 创建中，2 运行中，3 异常，4 重启中，5 停止中，6 已停止， 7 销毁中，8 已销毁)
	ProjectStatus *int64 `json:"ProjectStatus,omitnil,omitempty" name:"ProjectStatus"`

	// 日志接入点，用户忽略。
	AccessPoint *string `json:"AccessPoint,omitnil,omitempty" name:"AccessPoint"`

	// kafka旁路配置信息
	Kafka *Kafka `json:"Kafka,omitnil,omitempty" name:"Kafka"`
}

type RumPvInfo struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// pv访问量
	Pv *string `json:"Pv,omitnil,omitempty" name:"Pv"`

	// 时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type RumUvInfo struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// uv访问量
	Uv *string `json:"Uv,omitnil,omitempty" name:"Uv"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ScoreInfo struct {
	// duration
	StaticDuration *string `json:"StaticDuration,omitnil,omitempty" name:"StaticDuration"`

	// pv
	PagePv *string `json:"PagePv,omitnil,omitempty" name:"PagePv"`

	// 失败
	ApiFail *string `json:"ApiFail,omitnil,omitempty" name:"ApiFail"`

	// 请求
	ApiNum *string `json:"ApiNum,omitnil,omitempty" name:"ApiNum"`

	// fail
	StaticFail *string `json:"StaticFail,omitnil,omitempty" name:"StaticFail"`

	// 项目id
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// uv
	PageUv *string `json:"PageUv,omitnil,omitempty" name:"PageUv"`

	// 请求次数
	ApiDuration *string `json:"ApiDuration,omitnil,omitempty" name:"ApiDuration"`

	// 项目总分
	Score *string `json:"Score,omitnil,omitempty" name:"Score"`

	// error
	PageError *string `json:"PageError,omitnil,omitempty" name:"PageError"`

	// num
	StaticNum *string `json:"StaticNum,omitnil,omitempty" name:"StaticNum"`

	// num
	RecordNum *int64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// Duration
	PageDuration *string `json:"PageDuration,omitnil,omitempty" name:"PageDuration"`

	// 时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 页面性能评分
	PagePerformanceScore *string `json:"PagePerformanceScore,omitnil,omitempty" name:"PagePerformanceScore"`

	// js错误评分
	JsErrorScore *string `json:"JsErrorScore,omitnil,omitempty" name:"JsErrorScore"`

	// API性能评分
	ApiPerformanceScore *string `json:"ApiPerformanceScore,omitnil,omitempty" name:"ApiPerformanceScore"`

	// API可用性评分
	ApiAvaliableScore *string `json:"ApiAvaliableScore,omitnil,omitempty" name:"ApiAvaliableScore"`

	// 静态资源性能评分
	StaticPerformanceScore *string `json:"StaticPerformanceScore,omitnil,omitempty" name:"StaticPerformanceScore"`

	// 静态资源可用性评分
	StaticAvaliableScore *string `json:"StaticAvaliableScore,omitnil,omitempty" name:"StaticAvaliableScore"`
}

type ScoreInfoV2 struct {
	// <p>项目id</p>
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// <p>项目总分</p>
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>API性能评分</p>
	ApiPerformanceScore *float64 `json:"ApiPerformanceScore,omitnil,omitempty" name:"ApiPerformanceScore"`

	// <p>API可用性评分</p>
	ApiAvailableScore *float64 `json:"ApiAvailableScore,omitnil,omitempty" name:"ApiAvailableScore"`

	// <p>API调用总数</p>
	ApiNum *int64 `json:"ApiNum,omitnil,omitempty" name:"ApiNum"`

	// <p>API失败次数</p>
	ApiFail *int64 `json:"ApiFail,omitnil,omitempty" name:"ApiFail"`

	// <p>API平均持续时间</p>
	ApiDuration *float64 `json:"ApiDuration,omitnil,omitempty" name:"ApiDuration"`

	// <p>页面性能评分</p>
	PagePerformanceScore *float64 `json:"PagePerformanceScore,omitnil,omitempty" name:"PagePerformanceScore"`

	// <p>页面浏览量</p>
	PagePv *int64 `json:"PagePv,omitnil,omitempty" name:"PagePv"`

	// <p>独立访客数</p>
	PageUv *int64 `json:"PageUv,omitnil,omitempty" name:"PageUv"`

	// <p>页面错误数</p>
	PageError *int64 `json:"PageError,omitnil,omitempty" name:"PageError"`

	// <p>首屏时间</p>
	PageDuration *float64 `json:"PageDuration,omitnil,omitempty" name:"PageDuration"`

	// <p>平均 LCP</p>
	PageLCP *float64 `json:"PageLCP,omitnil,omitempty" name:"PageLCP"`

	// <p>平均 FID</p>
	PageFID *float64 `json:"PageFID,omitnil,omitempty" name:"PageFID"`

	// <p>平均 CLS</p>
	PageCLS *float64 `json:"PageCLS,omitnil,omitempty" name:"PageCLS"`

	// <p>平均 FCP</p>
	PageFCP *float64 `json:"PageFCP,omitnil,omitempty" name:"PageFCP"`

	// <p>平均 INP</p>
	PageINP *float64 `json:"PageINP,omitnil,omitempty" name:"PageINP"`

	// <p>JavaScript错误评分</p>
	JsErrorScore *float64 `json:"JsErrorScore,omitnil,omitempty" name:"JsErrorScore"`

	// <p>静态资源可用性评分</p>
	StaticAvailableScore *float64 `json:"StaticAvailableScore,omitnil,omitempty" name:"StaticAvailableScore"`

	// <p>静态资源性能评分</p>
	StaticPerformanceScore *float64 `json:"StaticPerformanceScore,omitnil,omitempty" name:"StaticPerformanceScore"`

	// <p>静态资源请求总数</p>
	StaticNum *int64 `json:"StaticNum,omitnil,omitempty" name:"StaticNum"`

	// <p>静态资源加载失败数</p>
	StaticFail *int64 `json:"StaticFail,omitnil,omitempty" name:"StaticFail"`

	// <p>静态资源加载时间</p>
	StaticDuration *float64 `json:"StaticDuration,omitnil,omitempty" name:"StaticDuration"`

	// <p>忽略的配置项</p>
	Exclusion *string `json:"Exclusion,omitnil,omitempty" name:"Exclusion"`
}

// Predefined struct for user
type StopInstanceRequestParams struct {
	// 需要停止的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改是否包括白名单
	IsModifyAll *bool `json:"IsModifyAll,omitnil,omitempty" name:"IsModifyAll"`
}

type StopInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 需要停止的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 修改是否包括白名单
	IsModifyAll *bool `json:"IsModifyAll,omitnil,omitempty" name:"IsModifyAll"`
}

func (r *StopInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IsModifyAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StopInstanceResponseParams `json:"Response"`
}

func (r *StopInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopProjectRequestParams struct {
	// 项目 id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type StopProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目 id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *StopProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopProjectResponse struct {
	*tchttp.BaseResponse
	Response *StopProjectResponseParams `json:"Response"`
}

func (r *StopProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTDEMApplicationRequestParams struct {
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type StopTDEMApplicationRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID（必填）</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *StopTDEMApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTDEMApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopTDEMApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTDEMApplicationResponseParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>停止后为 6</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopTDEMApplicationResponse struct {
	*tchttp.BaseResponse
	Response *StopTDEMApplicationResponseParams `json:"Response"`
}

func (r *StopTDEMApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTDEMApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StringKeyValue struct {
	// <p>键</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>值</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type SymbolicatedView struct {
	// <p>符号化版本</p>
	SymbolicationVersion *int64 `json:"SymbolicationVersion,omitnil,omitempty" name:"SymbolicationVersion"`

	// <p>符号化时间</p>
	SymbolicatedAt *string `json:"SymbolicatedAt,omitnil,omitempty" name:"SymbolicatedAt"`

	// <p>状态：success / missing / skipped</p>
	SymbolicationStatus *string `json:"SymbolicationStatus,omitnil,omitempty" name:"SymbolicationStatus"`

	// <p>跳过原因</p>
	SymbolicationSkipReason *string `json:"SymbolicationSkipReason,omitnil,omitempty" name:"SymbolicationSkipReason"`

	// <p>来源</p>
	SymbolicationSource *string `json:"SymbolicationSource,omitnil,omitempty" name:"SymbolicationSource"`

	// <p>release</p>
	SymbolicationRelease *string `json:"SymbolicationRelease,omitnil,omitempty" name:"SymbolicationRelease"`

	// <p>dist</p>
	SymbolicationDist *string `json:"SymbolicationDist,omitnil,omitempty" name:"SymbolicationDist"`

	// <p>耗时毫秒</p>
	SymbolicationDurationMs *int64 `json:"SymbolicationDurationMs,omitnil,omitempty" name:"SymbolicationDurationMs"`

	// <p>Symbolicator 请求 ID</p>
	SymbolicatorRequestId *string `json:"SymbolicatorRequestId,omitnil,omitempty" name:"SymbolicatorRequestId"`

	// <p>翻译后堆栈帧</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Frames []*EventFrame `json:"Frames,omitnil,omitempty" name:"Frames"`
}

type TDEMSSEDimensionRow struct {
	// <p>维度取值。空值保留为空字符串。</p>
	DimensionValue *string `json:"DimensionValue,omitnil,omitempty" name:"DimensionValue"`

	// <p>该取值的对话次数。</p>
	ConversationCount *int64 `json:"ConversationCount,omitnil,omitempty" name:"ConversationCount"`

	// <p>该取值上当前指标的值。没有可计算样本时为 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type TDEMSSEDistributionBin struct {
	// <p>档位左端，左闭。</p>
	Start *float64 `json:"Start,omitnil,omitempty" name:"Start"`

	// <p>档位右端。溢出档省略。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	End *float64 `json:"End,omitnil,omitempty" name:"End"`

	// <p>是否为溢出档。溢出档为 true，且不返回 End。</p>
	Overflow *bool `json:"Overflow,omitnil,omitempty" name:"Overflow"`

	// <p>落入该档的对话次数。</p>
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// <p>该档占有效样本的百分比，已乘 100。</p>
	Rate *float64 `json:"Rate,omitnil,omitempty" name:"Rate"`
}

type TDEMSSEPercentilePoint struct {
	// <p>时间桶起点，Unix 毫秒。</p>
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// <p>该桶算术平均值。没有样本的桶不返回。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Average *float64 `json:"Average,omitnil,omitempty" name:"Average"`

	// <p>50 分位。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	P50 *float64 `json:"P50,omitnil,omitempty" name:"P50"`

	// <p>75 分位。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	P75 *float64 `json:"P75,omitnil,omitempty" name:"P75"`

	// <p>95 分位。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	P95 *float64 `json:"P95,omitnil,omitempty" name:"P95"`

	// <p>99 分位。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	P99 *float64 `json:"P99,omitnil,omitempty" name:"P99"`
}

type TDEMSSETrendPoint struct {
	// <p>时间桶起点，Unix 毫秒。</p>
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// <p>该桶内的对话次数。</p>
	ConversationCount *int64 `json:"ConversationCount,omitnil,omitempty" name:"ConversationCount"`

	// <p>该桶内当前指标的值。桶内没有可计算样本时为 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type TDEMSSETrendSeries struct {
	// <p>维度取值。空值保留为空字符串。</p>
	DimensionValue *string `json:"DimensionValue,omitnil,omitempty" name:"DimensionValue"`

	// <p>该取值在各个时间桶上的点。</p>
	Points []*TDEMSSETrendPoint `json:"Points,omitnil,omitempty" name:"Points"`
}

type Tag struct {
	// 标签key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type UploadMiniProgramSourceRequestParams struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>源码zip的Base64，请求包上限10240KB</p>
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// <p>应用版本（可选）</p>
	Release *string `json:"Release,omitnil,omitempty" name:"Release"`

	// <p>原始文件名（可选）</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

type UploadMiniProgramSourceRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>源码zip的Base64，请求包上限10240KB</p>
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// <p>应用版本（可选）</p>
	Release *string `json:"Release,omitnil,omitempty" name:"Release"`

	// <p>原始文件名（可选）</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

func (r *UploadMiniProgramSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadMiniProgramSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FileBase64")
	delete(f, "Release")
	delete(f, "FileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadMiniProgramSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadMiniProgramSourceResponseParams struct {
	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// <p>请求结果（JSON字符串，包含 PackId/AppId/Release/BundleHash/Status/PageCount/ComponentCount/FileSize/Idempotent）</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadMiniProgramSourceResponse struct {
	*tchttp.BaseResponse
	Response *UploadMiniProgramSourceResponseParams `json:"Response"`
}

func (r *UploadMiniProgramSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadMiniProgramSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadSymbolArtifactRequestParams struct {
	// <p>制品类型。取值：js_sourcemap、ios_dsym、macos_dsym、windows_pdb、android_mapping、android_native、system_symbol。</p>
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// <p>对象大小（字节）。</p>
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// <p>对象内容的 SHA256 十六进制摘要，长度固定为 64 个字符。</p>
	Sha256 *string `json:"Sha256,omitnil,omitempty" name:"Sha256"`

	// <p>已上传对象的腾讯云 COS URL。与 FileUrl 二选一至少传一个；优先使用 COSUrl。必须属于当前配置的 Bucket 和 Region。</p>
	COSUrl *string `json:"COSUrl,omitnil,omitempty" name:"COSUrl"`

	// <p>应用 ID。应用级制品必填；system_symbol 不传。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>兼容别名，未传 AppId 时使用。</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>JS 应用版本。未传 AppVersion 时使用。</p>
	Release *string `json:"Release,omitnil,omitempty" name:"Release"`

	// <p>兼容字段，不参与制品匹配。</p>
	Dist *string `json:"Dist,omitnil,omitempty" name:"Dist"`

	// <p>压缩文件 URL。Kind=js_sourcemap 且 FileType 非 source_bundle 时必填。</p>
	MinifiedUrl *string `json:"MinifiedUrl,omitnil,omitempty" name:"MinifiedUrl"`

	// <p>JS 制品文件类型。可选值说明：sourcemap、minified、source_bundle；为空时默认 sourcemap。</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>调试 ID。除 js_sourcemap 外均必填。</p>
	DebugId *string `json:"DebugId,omitnil,omitempty" name:"DebugId"`

	// <p>构建 ID。ios_dsym、android_mapping、android_native 必填。</p>
	BuildId *string `json:"BuildId,omitnil,omitempty" name:"BuildId"`

	// <p>应用版本。js_sourcemap 必填，未传时使用 Release。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>兼容字段，请勿传入。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>系统名称。system_symbol 可填写。</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>系统版本。system_symbol 可填写。</p>
	OSVersion *string `json:"OSVersion,omitnil,omitempty" name:"OSVersion"`

	// <p>兼容字段，请勿传入。</p>
	StorePath *string `json:"StorePath,omitnil,omitempty" name:"StorePath"`

	// <p>文件名。原生和系统符号必填；ios_dsym 使用模块名，android_native 使用 abi/文件名。</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>文件内容类型，如 application/json。</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>COSUrl 的兼容别名；当未传 COSUrl 时使用。与 COSUrl 二选一至少传一个。</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`
}

type UploadSymbolArtifactRequest struct {
	*tchttp.BaseRequest
	
	// <p>制品类型。取值：js_sourcemap、ios_dsym、macos_dsym、windows_pdb、android_mapping、android_native、system_symbol。</p>
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// <p>对象大小（字节）。</p>
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// <p>对象内容的 SHA256 十六进制摘要，长度固定为 64 个字符。</p>
	Sha256 *string `json:"Sha256,omitnil,omitempty" name:"Sha256"`

	// <p>已上传对象的腾讯云 COS URL。与 FileUrl 二选一至少传一个；优先使用 COSUrl。必须属于当前配置的 Bucket 和 Region。</p>
	COSUrl *string `json:"COSUrl,omitnil,omitempty" name:"COSUrl"`

	// <p>应用 ID。应用级制品必填；system_symbol 不传。</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>兼容别名，未传 AppId 时使用。</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>JS 应用版本。未传 AppVersion 时使用。</p>
	Release *string `json:"Release,omitnil,omitempty" name:"Release"`

	// <p>兼容字段，不参与制品匹配。</p>
	Dist *string `json:"Dist,omitnil,omitempty" name:"Dist"`

	// <p>压缩文件 URL。Kind=js_sourcemap 且 FileType 非 source_bundle 时必填。</p>
	MinifiedUrl *string `json:"MinifiedUrl,omitnil,omitempty" name:"MinifiedUrl"`

	// <p>JS 制品文件类型。可选值说明：sourcemap、minified、source_bundle；为空时默认 sourcemap。</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>调试 ID。除 js_sourcemap 外均必填。</p>
	DebugId *string `json:"DebugId,omitnil,omitempty" name:"DebugId"`

	// <p>构建 ID。ios_dsym、android_mapping、android_native 必填。</p>
	BuildId *string `json:"BuildId,omitnil,omitempty" name:"BuildId"`

	// <p>应用版本。js_sourcemap 必填，未传时使用 Release。</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>兼容字段，请勿传入。</p>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>系统名称。system_symbol 可填写。</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>系统版本。system_symbol 可填写。</p>
	OSVersion *string `json:"OSVersion,omitnil,omitempty" name:"OSVersion"`

	// <p>兼容字段，请勿传入。</p>
	StorePath *string `json:"StorePath,omitnil,omitempty" name:"StorePath"`

	// <p>文件名。原生和系统符号必填；ios_dsym 使用模块名，android_native 使用 abi/文件名。</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>文件内容类型，如 application/json。</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>COSUrl 的兼容别名；当未传 COSUrl 时使用。与 COSUrl 二选一至少传一个。</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`
}

func (r *UploadSymbolArtifactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadSymbolArtifactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Kind")
	delete(f, "FileSize")
	delete(f, "Sha256")
	delete(f, "COSUrl")
	delete(f, "AppId")
	delete(f, "ProjectId")
	delete(f, "Release")
	delete(f, "Dist")
	delete(f, "MinifiedUrl")
	delete(f, "FileType")
	delete(f, "DebugId")
	delete(f, "BuildId")
	delete(f, "AppVersion")
	delete(f, "Platform")
	delete(f, "OS")
	delete(f, "OSVersion")
	delete(f, "StorePath")
	delete(f, "FileName")
	delete(f, "ContentType")
	delete(f, "FileUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadSymbolArtifactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadSymbolArtifactResponseParams struct {
	// <p>登记的制品信息，JSON 格式字符串。</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadSymbolArtifactResponse struct {
	*tchttp.BaseResponse
	Response *UploadSymbolArtifactResponseParams `json:"Response"`
}

func (r *UploadSymbolArtifactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadSymbolArtifactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsageDetailRow struct {
	// <p>业务系统 ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>业务系统名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>业务系统状态</p>
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用名称</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>端类型</p>
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// <p>应用状态</p>
	AppStatus *int64 `json:"AppStatus,omitnil,omitempty" name:"AppStatus"`

	// <p>上报量</p>
	ReportVolume *int64 `json:"ReportVolume,omitnil,omitempty" name:"ReportVolume"`

	// <p>占当前结果集比例</p>
	Percent *float64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// <p>行趋势</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ByDate []*UsageTimePoint `json:"ByDate,omitnil,omitempty" name:"ByDate"`
}

type UsageTimePoint struct {
	// <p>时间 YYYY-MM-DD HH:00:00</p>
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// <p>上报量</p>
	ReportVolume *int64 `json:"ReportVolume,omitnil,omitempty" name:"ReportVolume"`
}

type Whitelist struct {
	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 截止时间
	Ttl *string `json:"Ttl,omitnil,omitempty" name:"Ttl"`

	// 白名单自增ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 业务唯一标识
	WhitelistUin *string `json:"WhitelistUin,omitnil,omitempty" name:"WhitelistUin"`

	// 创建者ID
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// aid标识
	Aid *string `json:"Aid,omitnil,omitempty" name:"Aid"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}