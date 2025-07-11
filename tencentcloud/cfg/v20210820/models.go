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

package v20210820

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ActionFieldConfigDetail struct {
	// 组件类型
	// 可选项如下：
	// input  文本框
	// textarea  多行文本框
	// number  数值输入框
	// select   选择器
	// cascader  级联选择器
	// radio  单选
	// time   时间选择
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 组件label
	Lable *string `json:"Lable,omitnil,omitempty" name:"Lable"`

	// 组件唯一标识， 传回后端时的key
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 默认值
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 支持配置项如下,可根据需要选择配置项，不需要配置是设置空{}：
	// 
	// {
	// 
	//   placeholder: string (占位符)
	// 
	//   tooltip: string (提示信息)
	// 
	//   reg: RegExp (对输入内容格式进行正则校验的规则)
	// 
	//   max: number (对于输入框，限制最大输入字符数，对于数值输入框，设置上限)
	// 
	//   min: number (对于数值输入框，设置下限)
	// 
	//   step: number (设置数值输入框的步长，默认为1)
	// 
	//   format: string (时间选择的格式，如YYYY-MM-DD表示年月日, YYYY-MM-DD HH:mm:ss 表示时分秒)
	// 
	//   separator:  string[] (多行输入框的分隔符，不传或者为空时表示不分隔，直接返回用户输入的文本字符串)
	// 
	//   multiple: boolean (是否多选,对选择器和级联选择器有效)
	// 
	//   options:  选择器的选项【支持以下两种形式】
	// 
	// 直接给定选项数组  { value: string; label: string }[]
	// 通过调接口获取选项                                                                                                                                       { api: string(接口地址),                                                                                                                                       params: string[] (接口参数,对应于参数配置的field，前端根据field对应的所有组件的输入值作为参数查询数据， 为空时在组件加载时直接请求数据)                                                                                                    }
	// }
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 是否必填 (0 -- 否   1-- 是)
	Required *uint64 `json:"Required,omitnil,omitempty" name:"Required"`

	// compute配置依赖的其他field满足的条件时通过校验（如：三个表单项中必须至少有一个填写了）
	// 
	// [fieldName,
	// 
	// { config:  此项保留，等待后面具体场景细化  }
	// 
	// ]
	Validate *string `json:"Validate,omitnil,omitempty" name:"Validate"`

	// 是否可见
	Visible *string `json:"Visible,omitnil,omitempty" name:"Visible"`
}

type ActionFieldConfigResult struct {
	// 动作ID
	ActionId *uint64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 动作名称
	ActionName *string `json:"ActionName,omitnil,omitempty" name:"ActionName"`

	// 动作对应的栏位配置详情
	ConfigDetail []*ActionFieldConfigDetail `json:"ConfigDetail,omitnil,omitempty" name:"ConfigDetail"`
}

type ActionFilter struct {
	// 关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 搜索内容值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type ActionLibraryListResult struct {
	// 动作名称
	ActionName *string `json:"ActionName,omitnil,omitempty" name:"ActionName"`

	// 动作描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 动作类型。范围：["平台","自定义"]
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 创建人
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 动作风险描述
	RiskDesc *string `json:"RiskDesc,omitnil,omitempty" name:"RiskDesc"`

	// 动作ID
	ActionId *uint64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 动作属性（ 1：故障  2：恢复）
	AttributeId *uint64 `json:"AttributeId,omitnil,omitempty" name:"AttributeId"`

	// 关联的动作ID
	RelationActionId *uint64 `json:"RelationActionId,omitnil,omitempty" name:"RelationActionId"`

	// 操作命令
	ActionCommand *string `json:"ActionCommand,omitnil,omitempty" name:"ActionCommand"`

	// 动作类型（0 -- tat   1 -- 云API）
	ActionCommandType *uint64 `json:"ActionCommandType,omitnil,omitempty" name:"ActionCommandType"`

	// 自定义动作的参数，json string
	ActionContent *string `json:"ActionContent,omitnil,omitempty" name:"ActionContent"`

	// 二级分类
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 动作描述
	ActionDetail *string `json:"ActionDetail,omitnil,omitempty" name:"ActionDetail"`

	// 是否允许当前账号使用
	IsAllowed *bool `json:"IsAllowed,omitnil,omitempty" name:"IsAllowed"`

	// 最佳实践案例的链接地址
	ActionBestCase *string `json:"ActionBestCase,omitnil,omitempty" name:"ActionBestCase"`

	// 对象类型
	ObjectType *string `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// 监控指标ID列表
	MetricIdList []*uint64 `json:"MetricIdList,omitnil,omitempty" name:"MetricIdList"`

	// 是否是新动作
	IsNewAction *bool `json:"IsNewAction,omitnil,omitempty" name:"IsNewAction"`

	// 对象类型ID
	ObjectTypeId *int64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`
}

type ApmServiceInfo struct {
	// 业务ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 应用名称
	ServiceNameList []*string `json:"ServiceNameList,omitnil,omitempty" name:"ServiceNameList"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`
}

// Predefined struct for user
type CreateTaskFromActionRequestParams struct {
	// 动作ID，可从动作列表接口DescribeActionLibraryList获取
	TaskActionId *uint64 `json:"TaskActionId,omitnil,omitempty" name:"TaskActionId"`

	// 参与演练的实例ID
	TaskInstances []*string `json:"TaskInstances,omitnil,omitempty" name:"TaskInstances"`

	// 演练名称，不填则默认取动作名称
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// 演练描述，不填则默认取动作描述
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// 动作通用参数，需要json序列化传入，可以从动作详情接口DescribeActionFieldConfigList获取，不填默认使用动作默认参数
	TaskActionGeneralConfiguration *string `json:"TaskActionGeneralConfiguration,omitnil,omitempty" name:"TaskActionGeneralConfiguration"`

	// 动作自定义参数，需要json序列化传入，可以从动作详情接口DescribeActionFieldConfigList获取，不填默认使用动作默认参数，注意：必填参数，是没有默认值的 ，务必保证传入有效值
	TaskActionCustomConfiguration *string `json:"TaskActionCustomConfiguration,omitnil,omitempty" name:"TaskActionCustomConfiguration"`

	// 演练自动暂停时间，单位分钟, 不填则默认为60
	TaskPauseDuration *uint64 `json:"TaskPauseDuration,omitnil,omitempty" name:"TaskPauseDuration"`
}

type CreateTaskFromActionRequest struct {
	*tchttp.BaseRequest
	
	// 动作ID，可从动作列表接口DescribeActionLibraryList获取
	TaskActionId *uint64 `json:"TaskActionId,omitnil,omitempty" name:"TaskActionId"`

	// 参与演练的实例ID
	TaskInstances []*string `json:"TaskInstances,omitnil,omitempty" name:"TaskInstances"`

	// 演练名称，不填则默认取动作名称
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// 演练描述，不填则默认取动作描述
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// 动作通用参数，需要json序列化传入，可以从动作详情接口DescribeActionFieldConfigList获取，不填默认使用动作默认参数
	TaskActionGeneralConfiguration *string `json:"TaskActionGeneralConfiguration,omitnil,omitempty" name:"TaskActionGeneralConfiguration"`

	// 动作自定义参数，需要json序列化传入，可以从动作详情接口DescribeActionFieldConfigList获取，不填默认使用动作默认参数，注意：必填参数，是没有默认值的 ，务必保证传入有效值
	TaskActionCustomConfiguration *string `json:"TaskActionCustomConfiguration,omitnil,omitempty" name:"TaskActionCustomConfiguration"`

	// 演练自动暂停时间，单位分钟, 不填则默认为60
	TaskPauseDuration *uint64 `json:"TaskPauseDuration,omitnil,omitempty" name:"TaskPauseDuration"`
}

func (r *CreateTaskFromActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFromActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskActionId")
	delete(f, "TaskInstances")
	delete(f, "TaskTitle")
	delete(f, "TaskDescription")
	delete(f, "TaskActionGeneralConfiguration")
	delete(f, "TaskActionCustomConfiguration")
	delete(f, "TaskPauseDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskFromActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFromActionResponseParams struct {
	// 创建成功的演练ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTaskFromActionResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskFromActionResponseParams `json:"Response"`
}

func (r *CreateTaskFromActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFromActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFromMultiActionRequestParams struct {
	// 参与演练的实例ID
	// CVM ins-xxx
	// MySQL cdb-xxx
	// CLB lb-xxx
	// Redis crs-xxx
	// NAT网关 nat-xxx
	// Redis crs-xxx
	// 专线-独享专用通道 dcx-xxx
	// 标准集群普通节点 {"ClusterId":"cls-xxx","InstanceId":"ins-xxx","LanIP":"1.1.1.1"}
	// 标准集群Pod {"ClusterId":"cls-xxx","PodName":"podname","NodeName":"1.1.1.1","NameSpace":"ns","Workload":"workload"}
	// TDSQL-MySQL(InnoDB) tdsqlshard-xxx
	// TDSQL-C cynosdbmysql-xxx
	// VPC子网 subnet-xxxx
	// CKafka ckafka-xxx
	// MariaDB tdsql-xxxx
	// PostgreSQL postgres-xxx
	// 云原生网关 gateway-xxx
	// 标准集群超级节点 {"ClusterId":"cls-xxx","InstanceId":"eklet-xxx","LanIP":"1.1.1.1,"NodePoolId":"np-xxx"}
	// Serverless集群超级节点 {"ClusterId":"cls-xxxx","InstanceId":"eklet-xxxx","LanIP":"1.1.1.1"}
	// Elasticsearch集群 es-xxxx
	// RabbitMQ amqp-xxxx
	TaskInstances []*string `json:"TaskInstances,omitnil,omitempty" name:"TaskInstances"`

	// 演练名称，不填则默认取动作名称
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// 演练描述，不填则默认取动作描述
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// 演练自动暂停时间，单位分钟, 不填则默认为60
	TaskPauseDuration *uint64 `json:"TaskPauseDuration,omitnil,omitempty" name:"TaskPauseDuration"`

	// 演练动作组配置
	TaskAction []*TaskGroupForAction `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`
}

type CreateTaskFromMultiActionRequest struct {
	*tchttp.BaseRequest
	
	// 参与演练的实例ID
	// CVM ins-xxx
	// MySQL cdb-xxx
	// CLB lb-xxx
	// Redis crs-xxx
	// NAT网关 nat-xxx
	// Redis crs-xxx
	// 专线-独享专用通道 dcx-xxx
	// 标准集群普通节点 {"ClusterId":"cls-xxx","InstanceId":"ins-xxx","LanIP":"1.1.1.1"}
	// 标准集群Pod {"ClusterId":"cls-xxx","PodName":"podname","NodeName":"1.1.1.1","NameSpace":"ns","Workload":"workload"}
	// TDSQL-MySQL(InnoDB) tdsqlshard-xxx
	// TDSQL-C cynosdbmysql-xxx
	// VPC子网 subnet-xxxx
	// CKafka ckafka-xxx
	// MariaDB tdsql-xxxx
	// PostgreSQL postgres-xxx
	// 云原生网关 gateway-xxx
	// 标准集群超级节点 {"ClusterId":"cls-xxx","InstanceId":"eklet-xxx","LanIP":"1.1.1.1,"NodePoolId":"np-xxx"}
	// Serverless集群超级节点 {"ClusterId":"cls-xxxx","InstanceId":"eklet-xxxx","LanIP":"1.1.1.1"}
	// Elasticsearch集群 es-xxxx
	// RabbitMQ amqp-xxxx
	TaskInstances []*string `json:"TaskInstances,omitnil,omitempty" name:"TaskInstances"`

	// 演练名称，不填则默认取动作名称
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// 演练描述，不填则默认取动作描述
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// 演练自动暂停时间，单位分钟, 不填则默认为60
	TaskPauseDuration *uint64 `json:"TaskPauseDuration,omitnil,omitempty" name:"TaskPauseDuration"`

	// 演练动作组配置
	TaskAction []*TaskGroupForAction `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`
}

func (r *CreateTaskFromMultiActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFromMultiActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskInstances")
	delete(f, "TaskTitle")
	delete(f, "TaskDescription")
	delete(f, "TaskPauseDuration")
	delete(f, "TaskAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskFromMultiActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFromMultiActionResponseParams struct {
	// 创建成功的演练ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTaskFromMultiActionResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskFromMultiActionResponseParams `json:"Response"`
}

func (r *CreateTaskFromMultiActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFromMultiActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFromTemplateRequestParams struct {
	// 从经验库中查询到的经验模板ID
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 演练的配置参数
	TaskConfig *TaskConfig `json:"TaskConfig,omitnil,omitempty" name:"TaskConfig"`
}

type CreateTaskFromTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 从经验库中查询到的经验模板ID
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 演练的配置参数
	TaskConfig *TaskConfig `json:"TaskConfig,omitnil,omitempty" name:"TaskConfig"`
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
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeActionFieldConfigListRequestParams struct {
	// 动作ID列表
	ActionIds []*uint64 `json:"ActionIds,omitnil,omitempty" name:"ActionIds"`

	// 对象类型ID
	ObjectTypeId *uint64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`
}

type DescribeActionFieldConfigListRequest struct {
	*tchttp.BaseRequest
	
	// 动作ID列表
	ActionIds []*uint64 `json:"ActionIds,omitnil,omitempty" name:"ActionIds"`

	// 对象类型ID
	ObjectTypeId *uint64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`
}

func (r *DescribeActionFieldConfigListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionFieldConfigListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActionIds")
	delete(f, "ObjectTypeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeActionFieldConfigListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActionFieldConfigListResponseParams struct {
	// 通用栏位配置列表
	Common []*ActionFieldConfigResult `json:"Common,omitnil,omitempty" name:"Common"`

	// 动作栏位配置列表
	Results []*ActionFieldConfigResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 资源下线信息
	ResourceOffline []*ResourceOffline `json:"ResourceOffline,omitnil,omitempty" name:"ResourceOffline"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeActionFieldConfigListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeActionFieldConfigListResponseParams `json:"Response"`
}

func (r *DescribeActionFieldConfigListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionFieldConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActionLibraryListRequestParams struct {
	// 0-100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 默认值0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 对象类型ID
	ObjectType *uint64 `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// Keyword取值{"动作名称": "a_title", "描述": "a_desc", "动作类型": "a_type", "创建时间": "a_create_time", "二级分类": "a_resource_type"}
	Filters []*ActionFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 动作分类，1表示故障动作，2表示恢复动作
	Attribute []*int64 `json:"Attribute,omitnil,omitempty" name:"Attribute"`

	// 筛选项 -动作ID
	ActionIds []*uint64 `json:"ActionIds,omitnil,omitempty" name:"ActionIds"`
}

type DescribeActionLibraryListRequest struct {
	*tchttp.BaseRequest
	
	// 0-100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 默认值0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 对象类型ID
	ObjectType *uint64 `json:"ObjectType,omitnil,omitempty" name:"ObjectType"`

	// Keyword取值{"动作名称": "a_title", "描述": "a_desc", "动作类型": "a_type", "创建时间": "a_create_time", "二级分类": "a_resource_type"}
	Filters []*ActionFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 动作分类，1表示故障动作，2表示恢复动作
	Attribute []*int64 `json:"Attribute,omitnil,omitempty" name:"Attribute"`

	// 筛选项 -动作ID
	ActionIds []*uint64 `json:"ActionIds,omitnil,omitempty" name:"ActionIds"`
}

func (r *DescribeActionLibraryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionLibraryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ObjectType")
	delete(f, "Filters")
	delete(f, "Attribute")
	delete(f, "ActionIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeActionLibraryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActionLibraryListResponseParams struct {
	// 查询结果列表
	Results []*ActionLibraryListResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 符合记录条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeActionLibraryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeActionLibraryListResponseParams `json:"Response"`
}

func (r *DescribeActionLibraryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionLibraryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeObjectTypeListRequestParams struct {
	// 所支持的对象
	// 0：全平台产品
	// 1：平台接入的对象
	// 2：应用所支持的部分对象
	SupportType *int64 `json:"SupportType,omitnil,omitempty" name:"SupportType"`
}

type DescribeObjectTypeListRequest struct {
	*tchttp.BaseRequest
	
	// 所支持的对象
	// 0：全平台产品
	// 1：平台接入的对象
	// 2：应用所支持的部分对象
	SupportType *int64 `json:"SupportType,omitnil,omitempty" name:"SupportType"`
}

func (r *DescribeObjectTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeObjectTypeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SupportType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeObjectTypeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeObjectTypeListResponseParams struct {
	// 对象类型列表
	ObjectTypeList []*ObjectType `json:"ObjectTypeList,omitnil,omitempty" name:"ObjectTypeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeObjectTypeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeObjectTypeListResponseParams `json:"Response"`
}

func (r *DescribeObjectTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeObjectTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicy struct {
	// 保护策略ID列表
	TaskPolicyIdList []*string `json:"TaskPolicyIdList,omitnil,omitempty" name:"TaskPolicyIdList"`

	// 保护策略状态
	TaskPolicyStatus *string `json:"TaskPolicyStatus,omitnil,omitempty" name:"TaskPolicyStatus"`

	// 策略规则
	TaskPolicyRule *string `json:"TaskPolicyRule,omitnil,omitempty" name:"TaskPolicyRule"`

	// 护栏策略生效处理策略 1:顺序执行，2:暂停
	TaskPolicyDealType *int64 `json:"TaskPolicyDealType,omitnil,omitempty" name:"TaskPolicyDealType"`
}

// Predefined struct for user
type DescribeTaskExecuteLogsRequestParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 返回的内容行数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 日志起始的行数。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTaskExecuteLogsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 返回的内容行数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 日志起始的行数。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	LogMessage []*string `json:"LogMessage,omitnil,omitempty" name:"LogMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 演练名称
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// 标签键
	TaskTag []*string `json:"TaskTag,omitnil,omitempty" name:"TaskTag"`

	// 任务状态(1001 -- 未开始 1002 -- 进行中 1003 -- 暂停中 1004 -- 任务结束)
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 开始时间，固定格式%Y-%m-%d %H:%M:%S
	TaskStartTime *string `json:"TaskStartTime,omitnil,omitempty" name:"TaskStartTime"`

	// 结束时间，固定格式%Y-%m-%d %H:%M:%S
	TaskEndTime *string `json:"TaskEndTime,omitnil,omitempty" name:"TaskEndTime"`

	// 更新时间，固定格式%Y-%m-%d %H:%M:%S
	TaskUpdateTime *string `json:"TaskUpdateTime,omitnil,omitempty" name:"TaskUpdateTime"`

	// 标签对
	Tags []*TagWithDescribe `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 筛选条件
	Filters []*ActionFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 演练ID
	TaskId []*uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 关联应用ID筛选
	ApplicationId []*string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 关联应用筛选
	ApplicationName []*string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// 任务状态筛选--支持多选 任务状态(1001 -- 未开始 1002 -- 进行中 1003 -- 暂停中 1004 -- 任务结束)
	TaskStatusList []*uint64 `json:"TaskStatusList,omitnil,omitempty" name:"TaskStatusList"`

	// 架构ID
	ArchId *string `json:"ArchId,omitnil,omitempty" name:"ArchId"`

	// 架构名称
	ArchName *string `json:"ArchName,omitnil,omitempty" name:"ArchName"`
}

type DescribeTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 演练名称
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// 标签键
	TaskTag []*string `json:"TaskTag,omitnil,omitempty" name:"TaskTag"`

	// 任务状态(1001 -- 未开始 1002 -- 进行中 1003 -- 暂停中 1004 -- 任务结束)
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 开始时间，固定格式%Y-%m-%d %H:%M:%S
	TaskStartTime *string `json:"TaskStartTime,omitnil,omitempty" name:"TaskStartTime"`

	// 结束时间，固定格式%Y-%m-%d %H:%M:%S
	TaskEndTime *string `json:"TaskEndTime,omitnil,omitempty" name:"TaskEndTime"`

	// 更新时间，固定格式%Y-%m-%d %H:%M:%S
	TaskUpdateTime *string `json:"TaskUpdateTime,omitnil,omitempty" name:"TaskUpdateTime"`

	// 标签对
	Tags []*TagWithDescribe `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 筛选条件
	Filters []*ActionFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 演练ID
	TaskId []*uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 关联应用ID筛选
	ApplicationId []*string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 关联应用筛选
	ApplicationName []*string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// 任务状态筛选--支持多选 任务状态(1001 -- 未开始 1002 -- 进行中 1003 -- 暂停中 1004 -- 任务结束)
	TaskStatusList []*uint64 `json:"TaskStatusList,omitnil,omitempty" name:"TaskStatusList"`

	// 架构ID
	ArchId *string `json:"ArchId,omitnil,omitempty" name:"ArchId"`

	// 架构名称
	ArchName *string `json:"ArchName,omitnil,omitempty" name:"ArchName"`
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
	delete(f, "TaskUpdateTime")
	delete(f, "Tags")
	delete(f, "Filters")
	delete(f, "TaskId")
	delete(f, "ApplicationId")
	delete(f, "ApplicationName")
	delete(f, "TaskStatusList")
	delete(f, "ArchId")
	delete(f, "ArchName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskListResponseParams struct {
	// 无
	TaskList []*TaskListItem `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// 列表数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeTaskPolicyTriggerLogRequestParams struct {
	// 演练ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 页码
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 页数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeTaskPolicyTriggerLogRequest struct {
	*tchttp.BaseRequest
	
	// 演练ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 页码
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 页数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeTaskPolicyTriggerLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskPolicyTriggerLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Page")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskPolicyTriggerLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskPolicyTriggerLogResponseParams struct {
	// 触发日志
	TriggerLogs []*PolicyTriggerLog `json:"TriggerLogs,omitnil,omitempty" name:"TriggerLogs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskPolicyTriggerLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskPolicyTriggerLogResponseParams `json:"Response"`
}

func (r *DescribeTaskPolicyTriggerLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskPolicyTriggerLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskRequestParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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
	Task *Task `json:"Task,omitnil,omitempty" name:"Task"`

	// 任务对应的演练报告信息，null表示未导出报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportInfo *TaskReportInfo `json:"ReportInfo,omitnil,omitempty" name:"ReportInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 演练名称
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 标签键
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 状态，1---使用中， 2---停用
	IsUsed *int64 `json:"IsUsed,omitnil,omitempty" name:"IsUsed"`

	// 标签对
	Tags []*TagWithDescribe `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 经验来源 0-自建 1-专家推荐
	TemplateSource *int64 `json:"TemplateSource,omitnil,omitempty" name:"TemplateSource"`

	// 经验ID
	TemplateIdList []*int64 `json:"TemplateIdList,omitnil,omitempty" name:"TemplateIdList"`

	// 过滤参数
	Filters []*ActionFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// 分页Limit, 最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 演练名称
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 标签键
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 状态，1---使用中， 2---停用
	IsUsed *int64 `json:"IsUsed,omitnil,omitempty" name:"IsUsed"`

	// 标签对
	Tags []*TagWithDescribe `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 经验来源 0-自建 1-专家推荐
	TemplateSource *int64 `json:"TemplateSource,omitnil,omitempty" name:"TemplateSource"`

	// 经验ID
	TemplateIdList []*int64 `json:"TemplateIdList,omitnil,omitempty" name:"TemplateIdList"`

	// 过滤参数
	Filters []*ActionFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "TemplateSource")
	delete(f, "TemplateIdList")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateListResponseParams struct {
	// 经验库列表
	TemplateList []*TemplateListItem `json:"TemplateList,omitnil,omitempty" name:"TemplateList"`

	// 列表数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 经验库ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
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
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务动作ID
	TaskActionId *uint64 `json:"TaskActionId,omitnil,omitempty" name:"TaskActionId"`

	// 任务动作实例ID
	TaskInstanceIds []*uint64 `json:"TaskInstanceIds,omitnil,omitempty" name:"TaskInstanceIds"`

	// 是否操作整个任务
	IsOperateAll *bool `json:"IsOperateAll,omitnil,omitempty" name:"IsOperateAll"`

	// 操作类型：（1--启动   2--执行  3--跳过   5--重试）
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 动作组ID
	TaskGroupId *uint64 `json:"TaskGroupId,omitnil,omitempty" name:"TaskGroupId"`
}

type ExecuteTaskInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务动作ID
	TaskActionId *uint64 `json:"TaskActionId,omitnil,omitempty" name:"TaskActionId"`

	// 任务动作实例ID
	TaskInstanceIds []*uint64 `json:"TaskInstanceIds,omitnil,omitempty" name:"TaskInstanceIds"`

	// 是否操作整个任务
	IsOperateAll *bool `json:"IsOperateAll,omitnil,omitempty" name:"IsOperateAll"`

	// 操作类型：（1--启动   2--执行  3--跳过   5--重试）
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 动作组ID
	TaskGroupId *uint64 `json:"TaskGroupId,omitnil,omitempty" name:"TaskGroupId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type ExecuteTaskRequest struct {
	*tchttp.BaseRequest
	
	// 需要执行的任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态, 1001--未开始 1002--进行中（执行）1003--进行中（暂停）1004--执行结束
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 执行结果是否符合预期（当前扭转状态为执行结束时，需要必传此字段）
	IsExpect *bool `json:"IsExpect,omitnil,omitempty" name:"IsExpect"`

	// 演习结论（当演习状态转变为执行结束时，需要填写此字段）
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 问题以及改进
	Issue *string `json:"Issue,omitnil,omitempty" name:"Issue"`

	// 演练记录
	Record *string `json:"Record,omitnil,omitempty" name:"Record"`
}

type ModifyTaskRunStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态, 1001--未开始 1002--进行中（执行）1003--进行中（暂停）1004--执行结束
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 执行结果是否符合预期（当前扭转状态为执行结束时，需要必传此字段）
	IsExpect *bool `json:"IsExpect,omitnil,omitempty" name:"IsExpect"`

	// 演习结论（当演习状态转变为执行结束时，需要填写此字段）
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 问题以及改进
	Issue *string `json:"Issue,omitnil,omitempty" name:"Issue"`

	// 演练记录
	Record *string `json:"Record,omitnil,omitempty" name:"Record"`
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
	delete(f, "Issue")
	delete(f, "Record")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskRunStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskRunStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type ObjectType struct {
	// 对象类型ID
	ObjectTypeId *int64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`

	// 对象类型名称
	ObjectTypeTitle *string `json:"ObjectTypeTitle,omitnil,omitempty" name:"ObjectTypeTitle"`

	// 对象类型第一级
	ObjectTypeLevelOne *string `json:"ObjectTypeLevelOne,omitnil,omitempty" name:"ObjectTypeLevelOne"`

	// 对象类型参数
	ObjectTypeParams *ObjectTypeConfig `json:"ObjectTypeParams,omitnil,omitempty" name:"ObjectTypeParams"`

	// tke接口json解析规则，null不需要解析
	ObjectTypeJsonParse *ObjectTypeJsonParse `json:"ObjectTypeJsonParse,omitnil,omitempty" name:"ObjectTypeJsonParse"`

	// 是否包含新动作
	ObjectHasNewAction *bool `json:"ObjectHasNewAction,omitnil,omitempty" name:"ObjectHasNewAction"`

	// 对应在平台架构图中的资源类型名称
	ObjectPlatformName *string `json:"ObjectPlatformName,omitnil,omitempty" name:"ObjectPlatformName"`

	// 1：平台支持的对象 2：应用支持的部分对象
	ObjectSupportType *int64 `json:"ObjectSupportType,omitnil,omitempty" name:"ObjectSupportType"`

	// 1.接入层 2.逻辑层 3. 数据层
	ArchLayer *int64 `json:"ArchLayer,omitnil,omitempty" name:"ArchLayer"`

	// 是否支持演练生图
	IsArchSvg *bool `json:"IsArchSvg,omitnil,omitempty" name:"IsArchSvg"`
}

type ObjectTypeConfig struct {
	// 主键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 对象类型配置字段列表
	Fields []*ObjectTypeConfigFields `json:"Fields,omitnil,omitempty" name:"Fields"`
}

type ObjectTypeConfigFields struct {
	// instanceId
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 实例id
	Header *string `json:"Header,omitnil,omitempty" name:"Header"`

	// 字段值是否需要转译，当不需要转译时，此字段返回null
	Transfer *string `json:"Transfer,omitnil,omitempty" name:"Transfer"`

	// tke的pod字段信息解析
	JsonParse *string `json:"JsonParse,omitnil,omitempty" name:"JsonParse"`

	// 字段类型 0:str 1:list
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type ObjectTypeJsonParse struct {
	// 命名空间
	NameSpace *string `json:"NameSpace,omitnil,omitempty" name:"NameSpace"`

	// 工作负载名称
	WorkloadName *string `json:"WorkloadName,omitnil,omitempty" name:"WorkloadName"`

	// 节点IP
	LanIP *string `json:"LanIP,omitnil,omitempty" name:"LanIP"`

	// 节点ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type PolicyTriggerLog struct {
	// 演练ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 类型，0--触发，1--恢复
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 触发时间
	CreatTime *string `json:"CreatTime,omitnil,omitempty" name:"CreatTime"`
}

type ResourceOffline struct {
	// 资源ID
	ResourceId *int64 `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源下线时间
	ResourceDeleteTime *string `json:"ResourceDeleteTime,omitnil,omitempty" name:"ResourceDeleteTime"`

	// 资源下线提示
	ResourceDeleteMessage *string `json:"ResourceDeleteMessage,omitnil,omitempty" name:"ResourceDeleteMessage"`
}

type TagWithCreate struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagWithDescribe struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type Task struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务标题
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// 任务描述
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// 自定义标签
	TaskTag *string `json:"TaskTag,omitnil,omitempty" name:"TaskTag"`

	// 任务状态，1001--未开始  1002--进行中（执行）1003--进行中（暂停）1004--执行结束
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 任务结束状态，表明任务以何种状态结束: 0 -- 尚未结束，1 -- 成功，2-- 失败，3--终止
	TaskStatusType *int64 `json:"TaskStatusType,omitnil,omitempty" name:"TaskStatusType"`

	// 保护策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskProtectStrategy *string `json:"TaskProtectStrategy,omitnil,omitempty" name:"TaskProtectStrategy"`

	// 任务创建时间
	TaskCreateTime *string `json:"TaskCreateTime,omitnil,omitempty" name:"TaskCreateTime"`

	// 任务更新时间
	TaskUpdateTime *string `json:"TaskUpdateTime,omitnil,omitempty" name:"TaskUpdateTime"`

	// 任务动作组
	TaskGroups []*TaskGroup `json:"TaskGroups,omitnil,omitempty" name:"TaskGroups"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStartTime *string `json:"TaskStartTime,omitnil,omitempty" name:"TaskStartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskEndTime *string `json:"TaskEndTime,omitnil,omitempty" name:"TaskEndTime"`

	// 是否符合预期。1：符合预期，2：不符合预期
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskExpect *int64 `json:"TaskExpect,omitnil,omitempty" name:"TaskExpect"`

	// 演习记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskSummary *string `json:"TaskSummary,omitnil,omitempty" name:"TaskSummary"`

	// 任务模式。1:手工执行，2:自动执行
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// 自动暂停时长。单位分钟
	TaskPauseDuration *int64 `json:"TaskPauseDuration,omitnil,omitempty" name:"TaskPauseDuration"`

	// 演练创建者Uin
	TaskOwnerUin *string `json:"TaskOwnerUin,omitnil,omitempty" name:"TaskOwnerUin"`

	// 地域ID
	TaskRegionId *int64 `json:"TaskRegionId,omitnil,omitempty" name:"TaskRegionId"`

	// 监控指标列表
	TaskMonitors []*TaskMonitor `json:"TaskMonitors,omitnil,omitempty" name:"TaskMonitors"`

	// 保护策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskPolicy *DescribePolicy `json:"TaskPolicy,omitnil,omitempty" name:"TaskPolicy"`

	// 标签列表
	Tags []*TagWithDescribe `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 关联的演练计划ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskPlanId *int64 `json:"TaskPlanId,omitnil,omitempty" name:"TaskPlanId"`

	// 关联的演练计划名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskPlanTitle *string `json:"TaskPlanTitle,omitnil,omitempty" name:"TaskPlanTitle"`

	// 关联的应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 关联的应用名称
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// 关联的告警指标
	AlarmPolicy []*string `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`

	// 关联的APM服务
	ApmServiceList []*ApmServiceInfo `json:"ApmServiceList,omitnil,omitempty" name:"ApmServiceList"`

	// 关联的隐患验证项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyId *uint64 `json:"VerifyId,omitnil,omitempty" name:"VerifyId"`

	// 护栏处理方式，1--顺序回滚，2--演练暂停
	PolicyDealType *int64 `json:"PolicyDealType,omitnil,omitempty" name:"PolicyDealType"`

	// 计划开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskPlanStartTime *string `json:"TaskPlanStartTime,omitnil,omitempty" name:"TaskPlanStartTime"`

	// 计划结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskPlanEndTime *string `json:"TaskPlanEndTime,omitnil,omitempty" name:"TaskPlanEndTime"`

	// 人员组织
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskOrg []*TaskOrg `json:"TaskOrg,omitnil,omitempty" name:"TaskOrg"`

	// 问题和改进
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIssue *string `json:"TaskIssue,omitnil,omitempty" name:"TaskIssue"`

	// region信息
	TaskRegionName *string `json:"TaskRegionName,omitnil,omitempty" name:"TaskRegionName"`

	// 架构ID
	TaskArchId *string `json:"TaskArchId,omitnil,omitempty" name:"TaskArchId"`

	// 演练场景
	TaskScenario []*TaskTarget `json:"TaskScenario,omitnil,omitempty" name:"TaskScenario"`

	// 演练目的
	TaskPurpose []*TaskTarget `json:"TaskPurpose,omitnil,omitempty" name:"TaskPurpose"`
}

type TaskConfig struct {
	// 动作组配置，需要保证配置个数和经验中的动作组个数一致
	TaskGroupsConfig []*TaskGroupConfig `json:"TaskGroupsConfig,omitnil,omitempty" name:"TaskGroupsConfig"`

	// 更改后的演练名称，不填则默认取经验名称
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// 更改后的演练描述，不填则默认取经验描述
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// 演练执行模式：1----手工执行/ 2 ---自动执行，不填则默认取经验执行模式
	TaskMode *uint64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// 演练自动暂停时间，单位分钟, 不填则默认取经验自动暂停时间
	TaskPauseDuration *uint64 `json:"TaskPauseDuration,omitnil,omitempty" name:"TaskPauseDuration"`

	// 演练标签信息，不填则默认取经验标签
	Tags []*TagWithCreate `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 护栏处理方式，1--顺序回滚，2--演练暂停
	PolicyDealType *int64 `json:"PolicyDealType,omitnil,omitempty" name:"PolicyDealType"`
}

type TaskGroup struct {
	// 任务动作ID
	TaskGroupId *int64 `json:"TaskGroupId,omitnil,omitempty" name:"TaskGroupId"`

	// 分组标题
	TaskGroupTitle *string `json:"TaskGroupTitle,omitnil,omitempty" name:"TaskGroupTitle"`

	// 分组描述
	TaskGroupDescription *string `json:"TaskGroupDescription,omitnil,omitempty" name:"TaskGroupDescription"`

	// 任务分组顺序
	TaskGroupOrder *int64 `json:"TaskGroupOrder,omitnil,omitempty" name:"TaskGroupOrder"`

	// 对象类型ID
	ObjectTypeId *int64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`

	// 任务分组创建时间
	TaskGroupCreateTime *string `json:"TaskGroupCreateTime,omitnil,omitempty" name:"TaskGroupCreateTime"`

	// 任务分组更新时间
	TaskGroupUpdateTime *string `json:"TaskGroupUpdateTime,omitnil,omitempty" name:"TaskGroupUpdateTime"`

	// 动作分组动作列表
	TaskGroupActions []*TaskGroupAction `json:"TaskGroupActions,omitnil,omitempty" name:"TaskGroupActions"`

	// 实例列表
	TaskGroupInstanceList []*string `json:"TaskGroupInstanceList,omitnil,omitempty" name:"TaskGroupInstanceList"`

	// 执行模式。1 --- 顺序执行，2 --- 阶段执行
	TaskGroupMode *int64 `json:"TaskGroupMode,omitnil,omitempty" name:"TaskGroupMode"`

	// 不参演的实例列表
	TaskGroupDiscardInstanceList []*string `json:"TaskGroupDiscardInstanceList,omitnil,omitempty" name:"TaskGroupDiscardInstanceList"`

	// 参演实例列表
	TaskGroupSelectedInstanceList []*string `json:"TaskGroupSelectedInstanceList,omitnil,omitempty" name:"TaskGroupSelectedInstanceList"`

	// 机器选取规则
	TaskGroupInstancesExecuteRule []*TaskGroupInstancesExecuteRules `json:"TaskGroupInstancesExecuteRule,omitnil,omitempty" name:"TaskGroupInstancesExecuteRule"`
}

type TaskGroupAction struct {
	// 任务分组动作ID
	TaskGroupActionId *int64 `json:"TaskGroupActionId,omitnil,omitempty" name:"TaskGroupActionId"`

	// 任务分组动作实例列表
	TaskGroupInstances []*TaskGroupInstance `json:"TaskGroupInstances,omitnil,omitempty" name:"TaskGroupInstances"`

	// 动作ID
	ActionId *int64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 分组动作顺序
	TaskGroupActionOrder *int64 `json:"TaskGroupActionOrder,omitnil,omitempty" name:"TaskGroupActionOrder"`

	// 分组动作通用配置
	TaskGroupActionGeneralConfiguration *string `json:"TaskGroupActionGeneralConfiguration,omitnil,omitempty" name:"TaskGroupActionGeneralConfiguration"`

	// 分组动作自定义配置
	TaskGroupActionCustomConfiguration *string `json:"TaskGroupActionCustomConfiguration,omitnil,omitempty" name:"TaskGroupActionCustomConfiguration"`

	// 分组动作状态
	TaskGroupActionStatus *int64 `json:"TaskGroupActionStatus,omitnil,omitempty" name:"TaskGroupActionStatus"`

	// 动作分组创建时间
	TaskGroupActionCreateTime *string `json:"TaskGroupActionCreateTime,omitnil,omitempty" name:"TaskGroupActionCreateTime"`

	// 动作分组更新时间
	TaskGroupActionUpdateTime *string `json:"TaskGroupActionUpdateTime,omitnil,omitempty" name:"TaskGroupActionUpdateTime"`

	// 动作名称
	ActionTitle *string `json:"ActionTitle,omitnil,omitempty" name:"ActionTitle"`

	// 状态类型: 0 -- 无状态，1 -- 成功，2-- 失败，3--终止，4--跳过
	TaskGroupActionStatusType *int64 `json:"TaskGroupActionStatusType,omitnil,omitempty" name:"TaskGroupActionStatusType"`

	// RandomId
	TaskGroupActionRandomId *int64 `json:"TaskGroupActionRandomId,omitnil,omitempty" name:"TaskGroupActionRandomId"`

	// RecoverId
	TaskGroupActionRecoverId *int64 `json:"TaskGroupActionRecoverId,omitnil,omitempty" name:"TaskGroupActionRecoverId"`

	// ExecuteId
	TaskGroupActionExecuteId *int64 `json:"TaskGroupActionExecuteId,omitnil,omitempty" name:"TaskGroupActionExecuteId"`

	// 调用api类型，0:tat, 1:云api
	ActionApiType *int64 `json:"ActionApiType,omitnil,omitempty" name:"ActionApiType"`

	// 1:故障，2:恢复
	ActionAttribute *int64 `json:"ActionAttribute,omitnil,omitempty" name:"ActionAttribute"`

	// 动作类型：平台、自定义
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 是否可重试
	IsExecuteRedo *bool `json:"IsExecuteRedo,omitnil,omitempty" name:"IsExecuteRedo"`

	// 动作风险级别
	ActionRisk *string `json:"ActionRisk,omitnil,omitempty" name:"ActionRisk"`

	// 动作运行时间
	TaskGroupActionExecuteTime *int64 `json:"TaskGroupActionExecuteTime,omitnil,omitempty" name:"TaskGroupActionExecuteTime"`

	// 动作开始执行时间
	TaskGroupActionStartTime *string `json:"TaskGroupActionStartTime,omitnil,omitempty" name:"TaskGroupActionStartTime"`
}

type TaskGroupActionConfig struct {
	// 该动作在动作组中的顺序，从1开始，不填或填错将匹配不到经验中要修改参数的动作
	TaskGroupActionOrder *uint64 `json:"TaskGroupActionOrder,omitnil,omitempty" name:"TaskGroupActionOrder"`

	// 动作通用参数，需要json序列化传入，可以从查询经验详情接口获取，不填默认使用经验中动作参数
	TaskGroupActionGeneralConfiguration *string `json:"TaskGroupActionGeneralConfiguration,omitnil,omitempty" name:"TaskGroupActionGeneralConfiguration"`

	// 动作自定义参数，需要json序列化传入，可以从查询经验详情接口获取，不填默认使用经验中动作参数
	TaskGroupActionCustomConfiguration *string `json:"TaskGroupActionCustomConfiguration,omitnil,omitempty" name:"TaskGroupActionCustomConfiguration"`
}

type TaskGroupConfig struct {
	// 动作组所关联的实例对象
	// CVM ins-xxx
	// MySQL cdb-xxx
	// CLB lb-xxx
	// Redis crs-xxx
	// NAT网关 nat-xxx
	// 专线-独享专用通道 dcx-xxx
	// 标准集群普通节点 {"ClusterId":"cls-xxx","InstanceId":"ins-xxx","LanIP":"1.1.1.1"}
	// 标准集群Pod {"ClusterId":"cls-xxx","PodName":"podname","NodeName":"1.1.1.1","NameSpace":"ns","Workload":"workload"}
	// TDSQL-MySQL(InnoDB) tdsqlshard-xxx
	// TDSQL-C cynosdbmysql-xxx
	// VPC子网 subnet-xxxx
	// CKafka ckafka-xxx
	// MariaDB tdsql-xxxx
	// PostgreSQL postgres-xxx
	// 云原生网关 gateway-xxx
	// 标准集群超级节点 {"ClusterId":"cls-xxx","InstanceId":"eklet-xxx","LanIP":"1.1.1.1,"NodePoolId":"np-xxx"}
	// Serverless集群超级节点 {"ClusterId":"cls-xxxx","InstanceId":"eklet-xxxx","LanIP":"1.1.1.1"}
	// Elasticsearch集群 es-xxxx
	// RabbitMQ amqp-xxxx
	TaskGroupInstances []*string `json:"TaskGroupInstances,omitnil,omitempty" name:"TaskGroupInstances"`

	// 动作组标题，不填默认取经验中的动作组名称
	TaskGroupTitle *string `json:"TaskGroupTitle,omitnil,omitempty" name:"TaskGroupTitle"`

	// 动作组描述，不填默认取经验中的动作组描述
	TaskGroupDescription *string `json:"TaskGroupDescription,omitnil,omitempty" name:"TaskGroupDescription"`

	// 动作执行模式。1 --- 顺序执行，2 --- 阶段执行, 不填默认取经验中的动作组执行模式
	TaskGroupMode *uint64 `json:"TaskGroupMode,omitnil,omitempty" name:"TaskGroupMode"`

	// 动作组中的动作参数，不填默认使用经验中的动作参数，配置时可以只指定想要修改参数的动作
	TaskGroupActionsConfig []*TaskGroupActionConfig `json:"TaskGroupActionsConfig,omitnil,omitempty" name:"TaskGroupActionsConfig"`
}

type TaskGroupForAction struct {
	// 动作ID
	TaskActionId *int64 `json:"TaskActionId,omitnil,omitempty" name:"TaskActionId"`

	// {"ActionTimeout":1800}
	TaskActionGeneralConfiguration *string `json:"TaskActionGeneralConfiguration,omitnil,omitempty" name:"TaskActionGeneralConfiguration"`

	// {"ip": "0.0.0.0"}
	TaskActionCustomConfiguration *string `json:"TaskActionCustomConfiguration,omitnil,omitempty" name:"TaskActionCustomConfiguration"`
}

type TaskGroupInstance struct {
	// 实例ID
	TaskGroupInstanceId *int64 `json:"TaskGroupInstanceId,omitnil,omitempty" name:"TaskGroupInstanceId"`

	// 实例ID
	TaskGroupInstanceObjectId *string `json:"TaskGroupInstanceObjectId,omitnil,omitempty" name:"TaskGroupInstanceObjectId"`

	// 实例动作执行状态
	TaskGroupInstanceStatus *int64 `json:"TaskGroupInstanceStatus,omitnil,omitempty" name:"TaskGroupInstanceStatus"`

	// 实例创建时间
	TaskGroupInstanceCreateTime *string `json:"TaskGroupInstanceCreateTime,omitnil,omitempty" name:"TaskGroupInstanceCreateTime"`

	// 实例更新时间
	TaskGroupInstanceUpdateTime *string `json:"TaskGroupInstanceUpdateTime,omitnil,omitempty" name:"TaskGroupInstanceUpdateTime"`

	// 状态类型: 0 -- 无状态，1 -- 成功，2-- 失败，3--终止，4--跳过
	TaskGroupInstanceStatusType *int64 `json:"TaskGroupInstanceStatusType,omitnil,omitempty" name:"TaskGroupInstanceStatusType"`

	// 执行开始时间
	TaskGroupInstanceStartTime *string `json:"TaskGroupInstanceStartTime,omitnil,omitempty" name:"TaskGroupInstanceStartTime"`

	// 执行结束时间
	TaskGroupInstanceEndTime *string `json:"TaskGroupInstanceEndTime,omitnil,omitempty" name:"TaskGroupInstanceEndTime"`

	// 实例动作执行日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: TaskGroupInstanceExecuteLog is deprecated.
	TaskGroupInstanceExecuteLog *string `json:"TaskGroupInstanceExecuteLog,omitnil,omitempty" name:"TaskGroupInstanceExecuteLog"`

	// 实例是否可重试
	TaskGroupInstanceIsRedo *bool `json:"TaskGroupInstanceIsRedo,omitnil,omitempty" name:"TaskGroupInstanceIsRedo"`

	// 动作实例执行时间
	TaskGroupInstanceExecuteTime *int64 `json:"TaskGroupInstanceExecuteTime,omitnil,omitempty" name:"TaskGroupInstanceExecuteTime"`
}

type TaskGroupInstancesExecuteRules struct {
	// 实例选取模式
	TaskGroupInstancesExecuteMode *int64 `json:"TaskGroupInstancesExecuteMode,omitnil,omitempty" name:"TaskGroupInstancesExecuteMode"`

	// 按比例选取模式下选取比例
	TaskGroupInstancesExecutePercent *int64 `json:"TaskGroupInstancesExecutePercent,omitnil,omitempty" name:"TaskGroupInstancesExecutePercent"`

	// 按数量选取模式下选取数量
	TaskGroupInstancesExecuteNum *int64 `json:"TaskGroupInstancesExecuteNum,omitnil,omitempty" name:"TaskGroupInstancesExecuteNum"`
}

type TaskListItem struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务标题
	TaskTitle *string `json:"TaskTitle,omitnil,omitempty" name:"TaskTitle"`

	// 任务描述
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// 任务标签
	TaskTag *string `json:"TaskTag,omitnil,omitempty" name:"TaskTag"`

	// 任务状态(1001 -- 未开始   1002 -- 进行中  1003 -- 暂停中   1004 -- 任务结束)
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 任务创建时间
	TaskCreateTime *string `json:"TaskCreateTime,omitnil,omitempty" name:"TaskCreateTime"`

	// 任务更新时间
	TaskUpdateTime *string `json:"TaskUpdateTime,omitnil,omitempty" name:"TaskUpdateTime"`

	// 0--未开始，1--进行中，2--已完成
	TaskPreCheckStatus *int64 `json:"TaskPreCheckStatus,omitnil,omitempty" name:"TaskPreCheckStatus"`

	// 环境检查是否通过
	TaskPreCheckSuccess *bool `json:"TaskPreCheckSuccess,omitnil,omitempty" name:"TaskPreCheckSuccess"`

	// 演练是否符合预期 1-符合预期 2-不符合预期
	TaskExpect *int64 `json:"TaskExpect,omitnil,omitempty" name:"TaskExpect"`

	// 关联应用ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 关联应用名称
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// 验证项ID
	VerifyId *uint64 `json:"VerifyId,omitnil,omitempty" name:"VerifyId"`

	// 状态类型: 0 -- 无状态，1 -- 成功，2-- 失败，3--终止
	TaskStatusType *uint64 `json:"TaskStatusType,omitnil,omitempty" name:"TaskStatusType"`

	// 架构ID
	ArchId *string `json:"ArchId,omitnil,omitempty" name:"ArchId"`

	// 架构名称
	ArchName *string `json:"ArchName,omitnil,omitempty" name:"ArchName"`
}

type TaskMonitor struct {
	// 演练监控指标ID
	TaskMonitorId *int64 `json:"TaskMonitorId,omitnil,omitempty" name:"TaskMonitorId"`

	// 监控指标ID
	MetricId *uint64 `json:"MetricId,omitnil,omitempty" name:"MetricId"`

	// 监控指标对象类型ID
	TaskMonitorObjectTypeId *int64 `json:"TaskMonitorObjectTypeId,omitnil,omitempty" name:"TaskMonitorObjectTypeId"`

	// 指标名称
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 实例ID列表
	InstancesIds []*string `json:"InstancesIds,omitnil,omitempty" name:"InstancesIds"`

	// 中文指标
	MetricChineseName *string `json:"MetricChineseName,omitnil,omitempty" name:"MetricChineseName"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`
}

type TaskOrg struct {
	// 演练角色
	TaskRole *string `json:"TaskRole,omitnil,omitempty" name:"TaskRole"`

	// 负责人
	TaskOperator *string `json:"TaskOperator,omitnil,omitempty" name:"TaskOperator"`
}

type TaskReportInfo struct {
	// 0--未开始，1--正在导出，2--导出成功，3--导出失败
	Stage *int64 `json:"Stage,omitnil,omitempty" name:"Stage"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 有效期截止时间
	ExpirationTime *string `json:"ExpirationTime,omitnil,omitempty" name:"ExpirationTime"`

	// 是否有效
	Expired *bool `json:"Expired,omitnil,omitempty" name:"Expired"`

	// 演练报告cos文件地址
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 演练报告导出日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`

	// 0--未开始，1--正在归档，2--归档成功，3--归档失败
	ArchiveStage *int64 `json:"ArchiveStage,omitnil,omitempty" name:"ArchiveStage"`

	// 归档时间
	ArchiveTime *string `json:"ArchiveTime,omitnil,omitempty" name:"ArchiveTime"`

	// 归档ID
	ArchiveUuid *string `json:"ArchiveUuid,omitnil,omitempty" name:"ArchiveUuid"`
}

type TaskTarget struct {
	// 目标标签ID
	TargetId *int64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 目标描述
	TargetDesc *string `json:"TargetDesc,omitnil,omitempty" name:"TargetDesc"`

	// 1:演练场景
	// 2:演练目标
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 1:平台 2:用户个人
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`
}

type Template struct {
	// 经验库ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 经验库标题
	TemplateTitle *string `json:"TemplateTitle,omitnil,omitempty" name:"TemplateTitle"`

	// 经验库描述
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// 自定义标签
	TemplateTag *string `json:"TemplateTag,omitnil,omitempty" name:"TemplateTag"`

	// 使用状态。1 ---- 使用中，2 --- 停用
	TemplateIsUsed *int64 `json:"TemplateIsUsed,omitnil,omitempty" name:"TemplateIsUsed"`

	// 经验库创建时间
	TemplateCreateTime *string `json:"TemplateCreateTime,omitnil,omitempty" name:"TemplateCreateTime"`

	// 经验库更新时间
	TemplateUpdateTime *string `json:"TemplateUpdateTime,omitnil,omitempty" name:"TemplateUpdateTime"`

	// 经验库模式。1:手工执行，2:自动执行
	TemplateMode *int64 `json:"TemplateMode,omitnil,omitempty" name:"TemplateMode"`

	// 自动暂停时长。单位分钟
	TemplatePauseDuration *int64 `json:"TemplatePauseDuration,omitnil,omitempty" name:"TemplatePauseDuration"`

	// 演练创建者Uin
	TemplateOwnerUin *string `json:"TemplateOwnerUin,omitnil,omitempty" name:"TemplateOwnerUin"`

	// 地域ID
	TemplateRegionId *int64 `json:"TemplateRegionId,omitnil,omitempty" name:"TemplateRegionId"`

	// 动作组
	TemplateGroups []*TemplateGroup `json:"TemplateGroups,omitnil,omitempty" name:"TemplateGroups"`

	// 监控指标
	TemplateMonitors []*TemplateMonitor `json:"TemplateMonitors,omitnil,omitempty" name:"TemplateMonitors"`

	// 护栏监控
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplatePolicy *TemplatePolicy `json:"TemplatePolicy,omitnil,omitempty" name:"TemplatePolicy"`

	// 标签列表
	Tags []*TagWithDescribe `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 经验来源 0-自建 1-专家推荐
	TemplateSource *int64 `json:"TemplateSource,omitnil,omitempty" name:"TemplateSource"`

	// apm应用信息
	ApmServiceList []*ApmServiceInfo `json:"ApmServiceList,omitnil,omitempty" name:"ApmServiceList"`

	// 告警指标
	AlarmPolicy []*string `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`

	// 护栏处理方式，1--顺序回滚，2--演练暂停
	PolicyDealType *int64 `json:"PolicyDealType,omitnil,omitempty" name:"PolicyDealType"`

	// 演练场景
	TemplateScenario []*TaskTarget `json:"TemplateScenario,omitnil,omitempty" name:"TemplateScenario"`

	// 演练目的
	TemplatePurpose []*TaskTarget `json:"TemplatePurpose,omitnil,omitempty" name:"TemplatePurpose"`
}

type TemplateGroup struct {
	// 经验库动作ID
	TemplateGroupId *int64 `json:"TemplateGroupId,omitnil,omitempty" name:"TemplateGroupId"`

	// 经验库动作分组动作列表
	TemplateGroupActions []*TemplateGroupAction `json:"TemplateGroupActions,omitnil,omitempty" name:"TemplateGroupActions"`

	// 分组标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 分组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 分组顺序
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 执行模式。1 --- 顺序执行，2 --- 阶段执行
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 对象类型ID
	ObjectTypeId *int64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`

	// 分组创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 分组更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type TemplateGroupAction struct {
	// 经验库分组动作ID
	TemplateGroupActionId *int64 `json:"TemplateGroupActionId,omitnil,omitempty" name:"TemplateGroupActionId"`

	// 动作ID
	ActionId *int64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 分组动作顺序
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 分组动作通用配置
	GeneralConfiguration *string `json:"GeneralConfiguration,omitnil,omitempty" name:"GeneralConfiguration"`

	// 分组动作自定义配置
	CustomConfiguration *string `json:"CustomConfiguration,omitnil,omitempty" name:"CustomConfiguration"`

	// 动作分组创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 动作分组更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 动作名称
	ActionTitle *string `json:"ActionTitle,omitnil,omitempty" name:"ActionTitle"`

	// 自身随机id
	RandomId *int64 `json:"RandomId,omitnil,omitempty" name:"RandomId"`

	// 恢复动作id
	RecoverId *int64 `json:"RecoverId,omitnil,omitempty" name:"RecoverId"`

	// 执行动作id
	ExecuteId *int64 `json:"ExecuteId,omitnil,omitempty" name:"ExecuteId"`

	// 调用api类型，0:tat, 1:云api
	ActionApiType *int64 `json:"ActionApiType,omitnil,omitempty" name:"ActionApiType"`

	// 1:故障，2:恢复
	ActionAttribute *int64 `json:"ActionAttribute,omitnil,omitempty" name:"ActionAttribute"`

	// 动作类型：平台和自定义
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 动作风险等级，1:低风险 2:中风险 3:高风险
	ActionRisk *string `json:"ActionRisk,omitnil,omitempty" name:"ActionRisk"`

	// 故障表现
	FailurePerformance *string `json:"FailurePerformance,omitnil,omitempty" name:"FailurePerformance"`
}

type TemplateListItem struct {
	// 经验库ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 经验库标题
	TemplateTitle *string `json:"TemplateTitle,omitnil,omitempty" name:"TemplateTitle"`

	// 经验库描述
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// 经验库标签
	TemplateTag *string `json:"TemplateTag,omitnil,omitempty" name:"TemplateTag"`

	// 经验库状态。1 -- 使用中，2 -- 停用
	TemplateIsUsed *int64 `json:"TemplateIsUsed,omitnil,omitempty" name:"TemplateIsUsed"`

	// 经验库创建时间
	TemplateCreateTime *string `json:"TemplateCreateTime,omitnil,omitempty" name:"TemplateCreateTime"`

	// 经验库更新时间
	TemplateUpdateTime *string `json:"TemplateUpdateTime,omitnil,omitempty" name:"TemplateUpdateTime"`

	// 经验库关联的任务数量
	TemplateUsedNum *int64 `json:"TemplateUsedNum,omitnil,omitempty" name:"TemplateUsedNum"`

	// 经验库来源 0-自建经验 1-专家推荐
	TemplateSource *int64 `json:"TemplateSource,omitnil,omitempty" name:"TemplateSource"`
}

type TemplateMonitor struct {
	// pk
	MonitorId *int64 `json:"MonitorId,omitnil,omitempty" name:"MonitorId"`

	// 监控指标ID
	MetricId *int64 `json:"MetricId,omitnil,omitempty" name:"MetricId"`

	// 监控指标对象类型ID
	ObjectTypeId *int64 `json:"ObjectTypeId,omitnil,omitempty" name:"ObjectTypeId"`

	// 指标名称
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 中文指标
	MetricChineseName *string `json:"MetricChineseName,omitnil,omitempty" name:"MetricChineseName"`
}

type TemplatePolicy struct {
	// 保护策略ID列表
	TemplatePolicyIdList []*string `json:"TemplatePolicyIdList,omitnil,omitempty" name:"TemplatePolicyIdList"`

	// 策略规则
	TemplatePolicyRule *string `json:"TemplatePolicyRule,omitnil,omitempty" name:"TemplatePolicyRule"`

	// 护栏策略生效处理策略 1:顺序执行，2:暂停
	TemplatePolicyDealType *int64 `json:"TemplatePolicyDealType,omitnil,omitempty" name:"TemplatePolicyDealType"`
}

// Predefined struct for user
type TriggerPolicyRequestParams struct {
	// 混沌演练ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 触发内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 触发类型，0--触发；1--恢复
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`
}

type TriggerPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 混沌演练ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 触发内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 触发类型，0--触发；1--恢复
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`
}

func (r *TriggerPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Name")
	delete(f, "Content")
	delete(f, "TriggerType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TriggerPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TriggerPolicyResponseParams struct {
	// 演练ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 是否触发成功
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TriggerPolicyResponse struct {
	*tchttp.BaseResponse
	Response *TriggerPolicyResponseParams `json:"Response"`
}

func (r *TriggerPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}