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

package v20210416

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type APIGWParams struct {
	// HTTPS
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// POST
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`
}

// Predefined struct for user
type CheckRuleRequestParams struct {
	// Event信息
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// EventPattern信息
	EventPattern *string `json:"EventPattern,omitnil,omitempty" name:"EventPattern"`
}

type CheckRuleRequest struct {
	*tchttp.BaseRequest
	
	// Event信息
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// EventPattern信息
	EventPattern *string `json:"EventPattern,omitnil,omitempty" name:"EventPattern"`
}

func (r *CheckRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Event")
	delete(f, "EventPattern")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckRuleResponse struct {
	*tchttp.BaseResponse
	Response *CheckRuleResponseParams `json:"Response"`
}

func (r *CheckRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTransformationRequestParams struct {
	// 待处理的json字符串
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// 一个转换规则列表
	Transformations []*Transformation `json:"Transformations,omitnil,omitempty" name:"Transformations"`
}

type CheckTransformationRequest struct {
	*tchttp.BaseRequest
	
	// 待处理的json字符串
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// 一个转换规则列表
	Transformations []*Transformation `json:"Transformations,omitnil,omitempty" name:"Transformations"`
}

func (r *CheckTransformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTransformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Input")
	delete(f, "Transformations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckTransformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTransformationResponseParams struct {
	// 经过Transformations处理之后的数据
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckTransformationResponse struct {
	*tchttp.BaseResponse
	Response *CheckTransformationResponseParams `json:"Response"`
}

func (r *CheckTransformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTransformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CkafkaDeliveryParams struct {
	// ckafka topic name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// ckafka资源qcs六段式
	ResourceDescription *string `json:"ResourceDescription,omitnil,omitempty" name:"ResourceDescription"`
}

type CkafkaParams struct {
	// kafka offset
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// ckafka  topic
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type CkafkaTargetParams struct {
	// 要投递到的ckafka topic
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 重试策略
	RetryPolicy *RetryPolicy `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`
}

type Connection struct {
	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 更新时间
	ModTime *string `json:"ModTime,omitnil,omitempty" name:"ModTime"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 连接器ID
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 连接器描述
	ConnectionDescription *ConnectionDescription `json:"ConnectionDescription,omitnil,omitempty" name:"ConnectionDescription"`

	// 连接器名称
	ConnectionName *string `json:"ConnectionName,omitnil,omitempty" name:"ConnectionName"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ConnectionBrief struct {
	// 连接器类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 连接器状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ConnectionDescription struct {
	// 资源qcs六段式，更多参考 [资源六段式](https://cloud.tencent.com/document/product/598/10606)
	ResourceDescription *string `json:"ResourceDescription,omitnil,omitempty" name:"ResourceDescription"`

	// apigw参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	APIGWParams *APIGWParams `json:"APIGWParams,omitnil,omitempty" name:"APIGWParams"`

	// ckafka参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CkafkaParams *CkafkaParams `json:"CkafkaParams,omitnil,omitempty" name:"CkafkaParams"`

	// data transfer service (DTS)参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DTSParams *DTSParams `json:"DTSParams,omitnil,omitempty" name:"DTSParams"`

	// tdmq参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TDMQParams *TDMQParams `json:"TDMQParams,omitnil,omitempty" name:"TDMQParams"`
}

// Predefined struct for user
type CreateConnectionRequestParams struct {
	// 连接器描述
	ConnectionDescription *ConnectionDescription `json:"ConnectionDescription,omitnil,omitempty" name:"ConnectionDescription"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 连接器名称
	ConnectionName *string `json:"ConnectionName,omitnil,omitempty" name:"ConnectionName"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 连接器类型，目前支持以下类型:apigw/ckafka/dts/tdmq
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type CreateConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 连接器描述
	ConnectionDescription *ConnectionDescription `json:"ConnectionDescription,omitnil,omitempty" name:"ConnectionDescription"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 连接器名称
	ConnectionName *string `json:"ConnectionName,omitnil,omitempty" name:"ConnectionName"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 连接器类型，目前支持以下类型:apigw/ckafka/dts/tdmq
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *CreateConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConnectionDescription")
	delete(f, "EventBusId")
	delete(f, "ConnectionName")
	delete(f, "Description")
	delete(f, "Enable")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConnectionResponseParams struct {
	// 连接器ID
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConnectionResponse struct {
	*tchttp.BaseResponse
	Response *CreateConnectionResponseParams `json:"Response"`
}

func (r *CreateConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEventBusRequestParams struct {
	// 事件集名称，只能包含字母、中文、数字、下划线、连字符，以字母/中文开头，以数字、字母或中文结尾，2~60个字符
	EventBusName *string `json:"EventBusName,omitnil,omitempty" name:"EventBusName"`

	// 事件集描述，只能包含数字、中英文及常用标点符号，不超过200个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// EB存储时长
	SaveDays *int64 `json:"SaveDays,omitnil,omitempty" name:"SaveDays"`

	// EB是否开启存储
	EnableStore *bool `json:"EnableStore,omitnil,omitempty" name:"EnableStore"`
}

type CreateEventBusRequest struct {
	*tchttp.BaseRequest
	
	// 事件集名称，只能包含字母、中文、数字、下划线、连字符，以字母/中文开头，以数字、字母或中文结尾，2~60个字符
	EventBusName *string `json:"EventBusName,omitnil,omitempty" name:"EventBusName"`

	// 事件集描述，只能包含数字、中英文及常用标点符号，不超过200个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// EB存储时长
	SaveDays *int64 `json:"SaveDays,omitnil,omitempty" name:"SaveDays"`

	// EB是否开启存储
	EnableStore *bool `json:"EnableStore,omitnil,omitempty" name:"EnableStore"`
}

func (r *CreateEventBusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEventBusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusName")
	delete(f, "Description")
	delete(f, "SaveDays")
	delete(f, "EnableStore")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEventBusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEventBusResponseParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEventBusResponse struct {
	*tchttp.BaseResponse
	Response *CreateEventBusResponseParams `json:"Response"`
}

func (r *CreateEventBusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEventBusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleRequestParams struct {
	// 参考：[事件模式](https://cloud.tencent.com/document/product/1359/56084)
	EventPattern *string `json:"EventPattern,omitnil,omitempty" name:"EventPattern"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件规则名称，只能包含字母、中文、数字、下划线、连字符，以字母/中文开头，以数字、字母或中文结尾，2~60个字符
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 使能开关。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 事件规则描述，只能包含数字、中英文及常用标点符号，不超过200个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 参考：[事件模式](https://cloud.tencent.com/document/product/1359/56084)
	EventPattern *string `json:"EventPattern,omitnil,omitempty" name:"EventPattern"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件规则名称，只能包含字母、中文、数字、下划线、连字符，以字母/中文开头，以数字、字母或中文结尾，2~60个字符
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 使能开关。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 事件规则描述，只能包含数字、中英文及常用标点符号，不超过200个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventPattern")
	delete(f, "EventBusId")
	delete(f, "RuleName")
	delete(f, "Enable")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleResponseParams struct {
	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRuleResponseParams `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTargetRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 目标类型;取值范围:scf(云函数)/cls(日志服务)/amp(消息推送)/ckafka(消息推送)/es(大数据elastic-search)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 目标描述;scf类型示例:
	// {"ResourceDescription":"qcs::scf:ap-guangzhou:uin/2252646423:namespace/default/function/john-test-0326/$LATEST"};
	// cls类型示例:
	// {"ResourceDescription":"qcs::cls:ap-guangzhou:uin/12323442323:topic/7103f705-6c38-4b64-ac9d-428af0f2e732"}
	// ckafka类型示例:
	// {"ResourceDescription":"qcs::ckafka:ap-guangzhou:uin/1500000688:ckafkaId/uin/1500000688/ckafka-018q1nwj","CkafkaTargetParams":{"TopicName":"alert","RetryPolicy":{"RetryInterval":60,"MaxRetryAttempts":360}}}
	// amp类型-邮件/短信示例:
	// {"ResourceDescription":"qcs::eb-amp:ap-guangzhou:uin/100012505002:","AMPParams":{"NotificationTemplateId":10181,"Lang":"cn","NoticeReceivers":[{"UserType":"User","UserIds":["9424525"],"TimeWindow":{"From":"09:30:00","To":"23:30:00"},"Channels":["Email","SMS"]}]}}
	// es类型示例:
	// {"ResourceDescription":"qcs::es:ap-guangzhou:uin/1500000688:instance/es-7cplmhsd","ESTargetParams":{"EsVersion":"7.14.2","UserName":"elastic","Password":"xxxxx","NetMode":"privateLink","IndexPrefix":"auto-test","IndexSuffixMode":"default","RotationInterval":"none","IndexTemplateType":"","OutputMode":"default"}}
	TargetDescription *TargetDescription `json:"TargetDescription,omitnil,omitempty" name:"TargetDescription"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 批量投递最长等待时间
	BatchTimeout *int64 `json:"BatchTimeout,omitnil,omitempty" name:"BatchTimeout"`

	// 批量投递最大事件条数
	BatchEventCount *int64 `json:"BatchEventCount,omitnil,omitempty" name:"BatchEventCount"`

	// 开启批量投递使能
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitnil,omitempty" name:"EnableBatchDelivery"`
}

type CreateTargetRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 目标类型;取值范围:scf(云函数)/cls(日志服务)/amp(消息推送)/ckafka(消息推送)/es(大数据elastic-search)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 目标描述;scf类型示例:
	// {"ResourceDescription":"qcs::scf:ap-guangzhou:uin/2252646423:namespace/default/function/john-test-0326/$LATEST"};
	// cls类型示例:
	// {"ResourceDescription":"qcs::cls:ap-guangzhou:uin/12323442323:topic/7103f705-6c38-4b64-ac9d-428af0f2e732"}
	// ckafka类型示例:
	// {"ResourceDescription":"qcs::ckafka:ap-guangzhou:uin/1500000688:ckafkaId/uin/1500000688/ckafka-018q1nwj","CkafkaTargetParams":{"TopicName":"alert","RetryPolicy":{"RetryInterval":60,"MaxRetryAttempts":360}}}
	// amp类型-邮件/短信示例:
	// {"ResourceDescription":"qcs::eb-amp:ap-guangzhou:uin/100012505002:","AMPParams":{"NotificationTemplateId":10181,"Lang":"cn","NoticeReceivers":[{"UserType":"User","UserIds":["9424525"],"TimeWindow":{"From":"09:30:00","To":"23:30:00"},"Channels":["Email","SMS"]}]}}
	// es类型示例:
	// {"ResourceDescription":"qcs::es:ap-guangzhou:uin/1500000688:instance/es-7cplmhsd","ESTargetParams":{"EsVersion":"7.14.2","UserName":"elastic","Password":"xxxxx","NetMode":"privateLink","IndexPrefix":"auto-test","IndexSuffixMode":"default","RotationInterval":"none","IndexTemplateType":"","OutputMode":"default"}}
	TargetDescription *TargetDescription `json:"TargetDescription,omitnil,omitempty" name:"TargetDescription"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 批量投递最长等待时间
	BatchTimeout *int64 `json:"BatchTimeout,omitnil,omitempty" name:"BatchTimeout"`

	// 批量投递最大事件条数
	BatchEventCount *int64 `json:"BatchEventCount,omitnil,omitempty" name:"BatchEventCount"`

	// 开启批量投递使能
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitnil,omitempty" name:"EnableBatchDelivery"`
}

func (r *CreateTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "Type")
	delete(f, "TargetDescription")
	delete(f, "RuleId")
	delete(f, "BatchTimeout")
	delete(f, "BatchEventCount")
	delete(f, "EnableBatchDelivery")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTargetResponseParams struct {
	// 目标ID
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTargetResponse struct {
	*tchttp.BaseResponse
	Response *CreateTargetResponseParams `json:"Response"`
}

func (r *CreateTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTransformationRequestParams struct {
	// 事件总线 id
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 规则id
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 一个转换规则列表，当前仅限定一个;示例如下：[{"Extraction":{"ExtractionInputPath":"$.data.payload","Format":"JSON"},"EtlFilter":{"Filter":"{\"source\":\"ckafka.cloud.tencent\"}"},"Transform":{"OutputStructs":[{"Key":"op","Value":"$.op","ValueType":"JSONPATH"},{"Key":"table","Value":"$.source.table","ValueType":"JSONPATH"},{"Key":"id","Value":"$.after.id","ValueType":"JSONPATH"},{"Key":"app_id","Value":"$.after.app_id","ValueType":"JSONPATH"},{"Key":"spu_id","Value":"$.after.spu_id","ValueType":"JSONPATH"}]}}]
	Transformations []*Transformation `json:"Transformations,omitnil,omitempty" name:"Transformations"`
}

type CreateTransformationRequest struct {
	*tchttp.BaseRequest
	
	// 事件总线 id
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 规则id
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 一个转换规则列表，当前仅限定一个;示例如下：[{"Extraction":{"ExtractionInputPath":"$.data.payload","Format":"JSON"},"EtlFilter":{"Filter":"{\"source\":\"ckafka.cloud.tencent\"}"},"Transform":{"OutputStructs":[{"Key":"op","Value":"$.op","ValueType":"JSONPATH"},{"Key":"table","Value":"$.source.table","ValueType":"JSONPATH"},{"Key":"id","Value":"$.after.id","ValueType":"JSONPATH"},{"Key":"app_id","Value":"$.after.app_id","ValueType":"JSONPATH"},{"Key":"spu_id","Value":"$.after.spu_id","ValueType":"JSONPATH"}]}}]
	Transformations []*Transformation `json:"Transformations,omitnil,omitempty" name:"Transformations"`
}

func (r *CreateTransformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTransformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	delete(f, "Transformations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTransformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTransformationResponseParams struct {
	// 生成的转换器id
	TransformationId *string `json:"TransformationId,omitnil,omitempty" name:"TransformationId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTransformationResponse struct {
	*tchttp.BaseResponse
	Response *CreateTransformationResponseParams `json:"Response"`
}

func (r *CreateTransformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTransformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DTSParams struct {
	// Consumer Group Name
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// 账户名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type DeadLetterConfig struct {
	// 支持dlq、丢弃、忽略错误继续传递三种模式, 分别对应: DLQ,DROP,IGNORE_ERROR
	DisposeMethod *string `json:"DisposeMethod,omitnil,omitempty" name:"DisposeMethod"`

	// 设置了DLQ方式后,此选项必填. 错误消息会被投递到对应的kafka topic中
	// 注意：此字段可能返回 null，表示取不到有效值。
	CkafkaDeliveryParams *CkafkaDeliveryParams `json:"CkafkaDeliveryParams,omitnil,omitempty" name:"CkafkaDeliveryParams"`
}

// Predefined struct for user
type DeleteConnectionRequestParams struct {
	// 连接器ID
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`
}

type DeleteConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 连接器ID
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`
}

func (r *DeleteConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConnectionId")
	delete(f, "EventBusId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConnectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConnectionResponseParams `json:"Response"`
}

func (r *DeleteConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEventBusRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`
}

type DeleteEventBusRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`
}

func (r *DeleteEventBusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEventBusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEventBusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEventBusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEventBusResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEventBusResponseParams `json:"Response"`
}

func (r *DeleteEventBusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEventBusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRuleResponseParams `json:"Response"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTargetRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件目标ID
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteTargetRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件目标ID
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *DeleteTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "TargetId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTargetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTargetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTargetResponseParams `json:"Response"`
}

func (r *DeleteTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTransformationRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转换器id
	TransformationId *string `json:"TransformationId,omitnil,omitempty" name:"TransformationId"`
}

type DeleteTransformationRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转换器id
	TransformationId *string `json:"TransformationId,omitnil,omitempty" name:"TransformationId"`
}

func (r *DeleteTransformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTransformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	delete(f, "TransformationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTransformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTransformationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTransformationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTransformationResponseParams `json:"Response"`
}

func (r *DeleteTransformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTransformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogTagValueRequestParams struct {
	// 起始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 聚合字段,取值范围如下：Source(事件源),RuleIds(命中规则),Subject(实例ID),Region(地域)
	GroupField *string `json:"GroupField,omitnil,omitempty" name:"GroupField"`

	// 页数
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选条件
	Filter []*LogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeLogTagValueRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 聚合字段,取值范围如下：Source(事件源),RuleIds(命中规则),Subject(实例ID),Region(地域)
	GroupField *string `json:"GroupField,omitnil,omitempty" name:"GroupField"`

	// 页数
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选条件
	Filter []*LogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeLogTagValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogTagValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "EventBusId")
	delete(f, "GroupField")
	delete(f, "Page")
	delete(f, "Limit")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogTagValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogTagValueResponseParams struct {
	// 事件查询维度值结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*string `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogTagValueResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogTagValueResponseParams `json:"Response"`
}

func (r *DescribeLogTagValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogTagValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ESTargetParams struct {
	// 网络连接类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetMode *string `json:"NetMode,omitnil,omitempty" name:"NetMode"`

	// 索引前缀
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexPrefix *string `json:"IndexPrefix,omitnil,omitempty" name:"IndexPrefix"`

	// es日志轮换粒度
	// 注意：此字段可能返回 null，表示取不到有效值。
	RotationInterval *string `json:"RotationInterval,omitnil,omitempty" name:"RotationInterval"`

	// DTS事件配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputMode *string `json:"OutputMode,omitnil,omitempty" name:"OutputMode"`

	// DTS索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexSuffixMode *string `json:"IndexSuffixMode,omitnil,omitempty" name:"IndexSuffixMode"`

	// es模版类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexTemplateType *string `json:"IndexTemplateType,omitnil,omitempty" name:"IndexTemplateType"`
}

type EtlFilter struct {
	// 语法Rule规则保持一致
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type Event struct {
	// 事件源的信息,新产品上报必须符合EB的规范
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 事件数据，内容由创建事件的系统来控制，当前datacontenttype仅支持application/json;charset=utf-8，所以该字段是json字符串
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 事件类型，可自定义，选填。云服务默认写 COS:Created:PostObject，用“：”分割类型字段
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件来源详细描述，可自定义，选填。云服务默认为标准qcs资源表示语法：qcs::dts:ap-guangzhou:appid/uin:xxx
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// 事件发生的毫秒时间戳，
	// time.Now().UnixNano()/1e6
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 事件的地域信息，没有则默认是EB所在的地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 用于描述事件状态，非必须，默认是""
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 事件的唯一id，用户侧主动上传则需要保证风格一致
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type EventBus struct {
	// 更新时间
	ModTime *string `json:"ModTime,omitnil,omitempty" name:"ModTime"`

	// 事件集描述，不限字符类型，200字符描述以内
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 事件集名称，只能包含字母、数字、下划线、连字符，以字母开头，以数字或字母结尾，2~60个字符
	EventBusName *string `json:"EventBusName,omitnil,omitempty" name:"EventBusName"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件集类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 计费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 连接器基础信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionBriefs []*ConnectionBrief `json:"ConnectionBriefs,omitnil,omitempty" name:"ConnectionBriefs"`

	// 目标简要信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetBriefs []*TargetBrief `json:"TargetBriefs,omitnil,omitempty" name:"TargetBriefs"`
}

type Extraction struct {
	// JsonPath, 不指定则使用默认值$.
	ExtractionInputPath *string `json:"ExtractionInputPath,omitnil,omitempty" name:"ExtractionInputPath"`

	// 取值: TEXT/JSON
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 仅在Text需要传递
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextParams *TextParams `json:"TextParams,omitnil,omitempty" name:"TextParams"`
}

type Filter struct {
	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 过滤键的名称。EventBusName(事件集名称)/EventBusId(事件集Id)/Type(事件集类型:Cloud(云服务);Platform(平台型);Custom(自定义))/TagKey(标签键)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

// Predefined struct for user
type GetEventBusRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`
}

type GetEventBusRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`
}

func (r *GetEventBusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEventBusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetEventBusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEventBusResponseParams struct {
	// 更新时间
	ModTime *string `json:"ModTime,omitnil,omitempty" name:"ModTime"`

	// 事件集描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 日志主题ID
	ClsTopicId *string `json:"ClsTopicId,omitnil,omitempty" name:"ClsTopicId"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 日志集ID
	ClsLogsetId *string `json:"ClsLogsetId,omitnil,omitempty" name:"ClsLogsetId"`

	// 事件集名称
	EventBusName *string `json:"EventBusName,omitnil,omitempty" name:"EventBusName"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// （已废弃）事件集类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 计费模式
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// EB日志存储时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	SaveDays *int64 `json:"SaveDays,omitnil,omitempty" name:"SaveDays"`

	// EB日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`

	// 是否开启存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableStore *bool `json:"EnableStore,omitnil,omitempty" name:"EnableStore"`

	// 消息序列，是否有序
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkMode *string `json:"LinkMode,omitnil,omitempty" name:"LinkMode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetEventBusResponse struct {
	*tchttp.BaseResponse
	Response *GetEventBusResponseParams `json:"Response"`
}

func (r *GetEventBusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEventBusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPlatformEventTemplateRequestParams struct {
	// 平台产品事件类型
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`
}

type GetPlatformEventTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 平台产品事件类型
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`
}

func (r *GetPlatformEventTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPlatformEventTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPlatformEventTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPlatformEventTemplateResponseParams struct {
	// 平台产品事件模板
	EventTemplate *string `json:"EventTemplate,omitnil,omitempty" name:"EventTemplate"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetPlatformEventTemplateResponse struct {
	*tchttp.BaseResponse
	Response *GetPlatformEventTemplateResponseParams `json:"Response"`
}

func (r *GetPlatformEventTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPlatformEventTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRuleRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type GetRuleRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *GetRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRuleResponseParams struct {
	// 事件集id
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件规则id
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 事件规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 事件规则状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 事件规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 事件模式
	EventPattern *string `json:"EventPattern,omitnil,omitempty" name:"EventPattern"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 更新时间
	ModTime *string `json:"ModTime,omitnil,omitempty" name:"ModTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetRuleResponse struct {
	*tchttp.BaseResponse
	Response *GetRuleResponseParams `json:"Response"`
}

func (r *GetRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTransformationRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转换器id
	TransformationId *string `json:"TransformationId,omitnil,omitempty" name:"TransformationId"`
}

type GetTransformationRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转换器id
	TransformationId *string `json:"TransformationId,omitnil,omitempty" name:"TransformationId"`
}

func (r *GetTransformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTransformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	delete(f, "TransformationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTransformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTransformationResponseParams struct {
	// 转换规则列表
	Transformations []*Transformation `json:"Transformations,omitnil,omitempty" name:"Transformations"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTransformationResponse struct {
	*tchttp.BaseResponse
	Response *GetTransformationResponseParams `json:"Response"`
}

func (r *GetTransformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTransformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConnectionsRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 根据哪个字段进行返回结果排序，目前支持如下以下字段：AddTime, ModTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 以升序还是降序的方式返回结果，可选值 ASC 和 DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 根据哪个字段进行返回结果排序，目前支持如下以下字段：AddTime, ModTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 以升序还是降序的方式返回结果，可选值 ASC 和 DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConnectionsResponseParams struct {
	// 连接器信息
	Connections []*Connection `json:"Connections,omitnil,omitempty" name:"Connections"`

	// 连接器总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *ListConnectionsResponseParams `json:"Response"`
}

func (r *ListConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEventBusesRequestParams struct {
	// 根据哪个字段进行返回结果排序,支持以下字段：created_at（创建时间）, updated_at（修改时间）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 以升序还是降序的方式返回结果，可选值 ASC（升序） 和 DESC（降序）
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 过滤字段范围: EventBusName(事件集名称)/EventBusId(事件集Id)/Type(事件集类型:Cloud(云服务);Platform(平台型);Custom(自定义))/TagKey(标签键)。每次请求的Filters的上限为10，Filter.Values的上限为5。[{"Name":"Type","Values":["Cloud","Platform"]}]
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListEventBusesRequest struct {
	*tchttp.BaseRequest
	
	// 根据哪个字段进行返回结果排序,支持以下字段：created_at（创建时间）, updated_at（修改时间）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 以升序还是降序的方式返回结果，可选值 ASC（升序） 和 DESC（降序）
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 过滤字段范围: EventBusName(事件集名称)/EventBusId(事件集Id)/Type(事件集类型:Cloud(云服务);Platform(平台型);Custom(自定义))/TagKey(标签键)。每次请求的Filters的上限为10，Filter.Values的上限为5。[{"Name":"Type","Values":["Cloud","Platform"]}]
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListEventBusesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEventBusesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "Filters")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEventBusesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEventBusesResponseParams struct {
	// 事件集信息
	EventBuses []*EventBus `json:"EventBuses,omitnil,omitempty" name:"EventBuses"`

	// 事件集总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListEventBusesResponse struct {
	*tchttp.BaseResponse
	Response *ListEventBusesResponseParams `json:"Response"`
}

func (r *ListEventBusesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEventBusesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPlatformEventNamesRequestParams struct {
	// 平台产品类型
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`
}

type ListPlatformEventNamesRequest struct {
	*tchttp.BaseRequest
	
	// 平台产品类型
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`
}

func (r *ListPlatformEventNamesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPlatformEventNamesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListPlatformEventNamesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPlatformEventNamesResponseParams struct {
	// 平台产品列表
	EventNames []*PlatformEventDetail `json:"EventNames,omitnil,omitempty" name:"EventNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListPlatformEventNamesResponse struct {
	*tchttp.BaseResponse
	Response *ListPlatformEventNamesResponseParams `json:"Response"`
}

func (r *ListPlatformEventNamesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPlatformEventNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPlatformEventPatternsRequestParams struct {
	// 平台产品类型
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`
}

type ListPlatformEventPatternsRequest struct {
	*tchttp.BaseRequest
	
	// 平台产品类型
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`
}

func (r *ListPlatformEventPatternsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPlatformEventPatternsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListPlatformEventPatternsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPlatformEventPatternsResponseParams struct {
	// 平台产品事件匹配规则
	EventPatterns []*PlatformEventSummary `json:"EventPatterns,omitnil,omitempty" name:"EventPatterns"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListPlatformEventPatternsResponse struct {
	*tchttp.BaseResponse
	Response *ListPlatformEventPatternsResponseParams `json:"Response"`
}

func (r *ListPlatformEventPatternsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPlatformEventPatternsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPlatformProductsRequestParams struct {

}

type ListPlatformProductsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListPlatformProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPlatformProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListPlatformProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPlatformProductsResponseParams struct {
	// 平台产品列表
	PlatformProducts []*PlatformProduct `json:"PlatformProducts,omitnil,omitempty" name:"PlatformProducts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListPlatformProductsResponse struct {
	*tchttp.BaseResponse
	Response *ListPlatformProductsResponseParams `json:"Response"`
}

func (r *ListPlatformProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPlatformProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRulesRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 根据哪个字段进行返回结果排序,支持以下字段：AddTime（创建时间）, ModTime（修改时间）,name（规则名称）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 以升序还是降序的方式返回结果，可选值 ASC（升序） 和 DESC（降序）
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type ListRulesRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 根据哪个字段进行返回结果排序,支持以下字段：AddTime（创建时间）, ModTime（修改时间）,name（规则名称）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 以升序还是降序的方式返回结果，可选值 ASC（升序） 和 DESC（降序）
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *ListRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRulesResponseParams struct {
	// 事件规则信息
	Rules []*Rule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 事件规则总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRulesResponse struct {
	*tchttp.BaseResponse
	Response *ListRulesResponseParams `json:"Response"`
}

func (r *ListRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTargetsRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 根据哪个字段进行返回结果排序,支持以下字段：AddTime（创建时间）, ModTime（修改时间）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 以升序还是降序的方式返回结果，可选值 ASC（升序） 和 DESC（降序）
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type ListTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 根据哪个字段进行返回结果排序,支持以下字段：AddTime（创建时间）, ModTime（修改时间）
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 以升序还是降序的方式返回结果，可选值 ASC（升序） 和 DESC（降序）
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *ListTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTargetsResponseParams struct {
	// 目标总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 目标信息
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTargetsResponse struct {
	*tchttp.BaseResponse
	Response *ListTargetsResponseParams `json:"Response"`
}

func (r *ListTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogFilter struct {
	// 过滤字段名称
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 运算符，全等 eq，不等 neq，相似 like，排除相似 not like,  小于 lt，小于且等于 lte，大于 gt，大于且等于 gte，在范围内 range，不在范围内 norange
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 过滤值,范围运算需要同时输入两个值，以英文逗号分隔
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 该层级filters逻辑关系，取值 "AND" 或 "OR"
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// LogFilters数组
	Filters []*LogFilters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type LogFilters struct {
	// 过滤字段名称，取值范围如下:region(地域)，type(事件类型)，source(事件源)，status(事件状态)
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 运算符, 全等 eq，不等 neq，相似 like，排除相似 not like,  小于 lt，小于且等于 lte，大于 gt，大于且等于 gte，在范围内 range，不在范围内 norange
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 过滤值，范围运算需要同时输入两个值，以英文逗号分隔
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputStructParam struct {
	// 对应输出json中的key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 可以填json-path也可以支持常量或者内置关键字date类型
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// value的数据类型, 可选值: STRING, NUMBER,BOOLEAN,NULL,SYS_VARIABLE,JSONPATH
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`
}

type PlatformEventDetail struct {
	// 事件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`
}

type PlatformEventSummary struct {
	// 平台事件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 平台事件匹配规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventPattern *string `json:"EventPattern,omitnil,omitempty" name:"EventPattern"`
}

type PlatformProduct struct {
	// 平台产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 平台产品类型
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`
}

// Predefined struct for user
type PublishEventRequestParams struct {
	// 事件列表
	EventList []*Event `json:"EventList,omitnil,omitempty" name:"EventList"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`
}

type PublishEventRequest struct {
	*tchttp.BaseRequest
	
	// 事件列表
	EventList []*Event `json:"EventList,omitnil,omitempty" name:"EventList"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`
}

func (r *PublishEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventList")
	delete(f, "EventBusId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishEventResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PublishEventResponse struct {
	*tchttp.BaseResponse
	Response *PublishEventResponseParams `json:"Response"`
}

func (r *PublishEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutEventsRequestParams struct {
	// 事件列表
	EventList []*Event `json:"EventList,omitnil,omitempty" name:"EventList"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`
}

type PutEventsRequest struct {
	*tchttp.BaseRequest
	
	// 事件列表
	EventList []*Event `json:"EventList,omitnil,omitempty" name:"EventList"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`
}

func (r *PutEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventList")
	delete(f, "EventBusId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutEventsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PutEventsResponse struct {
	*tchttp.BaseResponse
	Response *PutEventsResponseParams `json:"Response"`
}

func (r *PutEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryPolicy struct {
	// 重试间隔 单位:秒
	RetryInterval *uint64 `json:"RetryInterval,omitnil,omitempty" name:"RetryInterval"`

	// 最大重试次数
	MaxRetryAttempts *uint64 `json:"MaxRetryAttempts,omitnil,omitempty" name:"MaxRetryAttempts"`
}

type Rule struct {
	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 修改时间
	ModTime *string `json:"ModTime,omitnil,omitempty" name:"ModTime"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Target 简要信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*TargetBrief `json:"Targets,omitnil,omitempty" name:"Targets"`

	// rule设置的dlq规则. 可能为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterConfig *DeadLetterConfig `json:"DeadLetterConfig,omitnil,omitempty" name:"DeadLetterConfig"`
}

type SCFParams struct {
	// 批量投递最长等待时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchTimeout *int64 `json:"BatchTimeout,omitnil,omitempty" name:"BatchTimeout"`

	// 批量投递最大事件条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchEventCount *int64 `json:"BatchEventCount,omitnil,omitempty" name:"BatchEventCount"`

	// 开启批量投递使能
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitnil,omitempty" name:"EnableBatchDelivery"`
}

// Predefined struct for user
type SearchLogRequestParams struct {
	// 起始时间unix 毫秒时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间unix 毫秒时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 页码
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 事件查询筛选条件；示例如下：[{"key":"host","operator":"eq","value":"106.53.106.243"},{"type":"AND","filters":[{"key":"region","operator":"like","value":"*guangzhou*"},{"key":"type","operator":"eq","value":"cvm:ErrorEvent:GuestReboot"}]},{"type":"OR","filters":[{"key":"field1","operator":"like","value":"*access*"},{"key":"field2","operator":"eq","value":"custorm"}]}]
	Filter []*LogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 事件查询结果排序，["timestamp","subject"]
	OrderFields []*string `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 排序方式，asc 从旧到新，desc 从新到旧
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

type SearchLogRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间unix 毫秒时间戳
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间unix 毫秒时间戳
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 页码
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数据大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 事件查询筛选条件；示例如下：[{"key":"host","operator":"eq","value":"106.53.106.243"},{"type":"AND","filters":[{"key":"region","operator":"like","value":"*guangzhou*"},{"key":"type","operator":"eq","value":"cvm:ErrorEvent:GuestReboot"}]},{"type":"OR","filters":[{"key":"field1","operator":"like","value":"*access*"},{"key":"field2","operator":"eq","value":"custorm"}]}]
	Filter []*LogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 事件查询结果排序，["timestamp","subject"]
	OrderFields []*string `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 排序方式，asc 从旧到新，desc 从新到旧
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

func (r *SearchLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "EventBusId")
	delete(f, "Page")
	delete(f, "Limit")
	delete(f, "Filter")
	delete(f, "OrderFields")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchLogResponseParams struct {
	// 事件总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 每页事件条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 事件查询结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*SearchLogResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchLogResponseParams `json:"Response"`
}

func (r *SearchLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogResult struct {
	// 单条日志上报时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 日志内容详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 事件来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件匹配规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleIds *string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 事件状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type TDMQParams struct {
	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群支撑网接入点
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterEndPoint *string `json:"ClusterEndPoint,omitnil,omitempty" name:"ClusterEndPoint"`
}

type Tag struct {
	// 标签名称
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Target struct {
	// 目标类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 目标ID
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 目标描述
	TargetDescription *TargetDescription `json:"TargetDescription,omitnil,omitempty" name:"TargetDescription"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 开启批量投递使能
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitnil,omitempty" name:"EnableBatchDelivery"`

	// 批量投递最长等待时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchTimeout *int64 `json:"BatchTimeout,omitnil,omitempty" name:"BatchTimeout"`

	// 批量投递最大事件条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchEventCount *int64 `json:"BatchEventCount,omitnil,omitempty" name:"BatchEventCount"`
}

type TargetBrief struct {
	// 目标ID
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 目标类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type TargetDescription struct {
	// QCS资源六段式，更多参考 [资源六段式](https://cloud.tencent.com/document/product/598/10606)；scf资源六段式示例[qcs::scf:ap-guangzhou:uin/123:namespace/test(函数命名空间)/function/test(函数名)/$LATEST(函数版本)] amp资源六段式示例[qcs::eb-amp:ap-guangzhou:uin/123:] ckafka资源六段式示例[qcs::ckafka:ap-guangzhou:uin/123:ckafkaId/uin/123/ckafka-123(ckafka实例Id)] cls资源六段式示例[qcs::cls:ap-guangzhou:uin/123:topic/122332442(topicId)] es资源六段式示例[qcs::es:ap-guangzhou:appid/123/uin/456:instance/es-7cplmhsd(es实例Id)]
	ResourceDescription *string `json:"ResourceDescription,omitnil,omitempty" name:"ResourceDescription"`

	// 云函数参数
	SCFParams *SCFParams `json:"SCFParams,omitnil,omitempty" name:"SCFParams"`

	// Ckafka参数
	CkafkaTargetParams *CkafkaTargetParams `json:"CkafkaTargetParams,omitnil,omitempty" name:"CkafkaTargetParams"`

	// ElasticSearch参数
	ESTargetParams *ESTargetParams `json:"ESTargetParams,omitnil,omitempty" name:"ESTargetParams"`
}

type TextParams struct {
	// 逗号、| 、制表符、空格、换行符、%、#，限制长度为 1。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Separator *string `json:"Separator,omitnil,omitempty" name:"Separator"`

	// 填写正则表达式：长度128
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`
}

type Transform struct {
	// 描述如何数据转换
	OutputStructs []*OutputStructParam `json:"OutputStructs,omitnil,omitempty" name:"OutputStructs"`
}

type Transformation struct {
	// 描述如何提取数据，{"ExtractionInputPath":"$.data.payload","Format":"JSON"}
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extraction *Extraction `json:"Extraction,omitnil,omitempty" name:"Extraction"`

	// 描述如何过滤数据;{"Filter":"{\"source\":\"ckafka.cloud.tencent\"}"}
	// 注意：此字段可能返回 null，表示取不到有效值。
	EtlFilter *EtlFilter `json:"EtlFilter,omitnil,omitempty" name:"EtlFilter"`

	// 描述如何数据转换;"OutputStructs":[{"Key":"op","Value":"$.op","ValueType":"JSONPATH"}]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transform *Transform `json:"Transform,omitnil,omitempty" name:"Transform"`
}

// Predefined struct for user
type UpdateConnectionRequestParams struct {
	// 连接器ID
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 连接器名称
	ConnectionName *string `json:"ConnectionName,omitnil,omitempty" name:"ConnectionName"`
}

type UpdateConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 连接器ID
	ConnectionId *string `json:"ConnectionId,omitnil,omitempty" name:"ConnectionId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 连接器名称
	ConnectionName *string `json:"ConnectionName,omitnil,omitempty" name:"ConnectionName"`
}

func (r *UpdateConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConnectionId")
	delete(f, "EventBusId")
	delete(f, "Enable")
	delete(f, "Description")
	delete(f, "ConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateConnectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateConnectionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateConnectionResponseParams `json:"Response"`
}

func (r *UpdateConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEventBusRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件集描述，不限字符类型，200字符描述以内
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 事件集名称，只能包含字母、数字、下划线、连字符，以字母开头，以数字或字母结尾，2~60个字符
	EventBusName *string `json:"EventBusName,omitnil,omitempty" name:"EventBusName"`

	// EB日志存储时长
	SaveDays *int64 `json:"SaveDays,omitnil,omitempty" name:"SaveDays"`

	// EB日志主题ID
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`

	// 是否开启存储
	EnableStore *bool `json:"EnableStore,omitnil,omitempty" name:"EnableStore"`
}

type UpdateEventBusRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件集描述，不限字符类型，200字符描述以内
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 事件集名称，只能包含字母、数字、下划线、连字符，以字母开头，以数字或字母结尾，2~60个字符
	EventBusName *string `json:"EventBusName,omitnil,omitempty" name:"EventBusName"`

	// EB日志存储时长
	SaveDays *int64 `json:"SaveDays,omitnil,omitempty" name:"SaveDays"`

	// EB日志主题ID
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`

	// 是否开启存储
	EnableStore *bool `json:"EnableStore,omitnil,omitempty" name:"EnableStore"`
}

func (r *UpdateEventBusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEventBusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "Description")
	delete(f, "EventBusName")
	delete(f, "SaveDays")
	delete(f, "LogTopicId")
	delete(f, "EnableStore")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEventBusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEventBusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateEventBusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEventBusResponseParams `json:"Response"`
}

func (r *UpdateEventBusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEventBusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRuleRequestParams struct {
	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 使能开关。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 事件规则描述，只能包含数字、中英文及常用标点符号，不超过200个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 参考：[事件模式](https://cloud.tencent.com/document/product/1359/56084)
	EventPattern *string `json:"EventPattern,omitnil,omitempty" name:"EventPattern"`

	// 事件规则名称，只能包含字母、中文、数字、下划线、连字符，以字母/中文开头，以数字、字母或中文结尾，2~60个字符
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type UpdateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 使能开关。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 事件规则描述，只能包含数字、中英文及常用标点符号，不超过200个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 参考：[事件模式](https://cloud.tencent.com/document/product/1359/56084)
	EventPattern *string `json:"EventPattern,omitnil,omitempty" name:"EventPattern"`

	// 事件规则名称，只能包含字母、中文、数字、下划线、连字符，以字母/中文开头，以数字、字母或中文结尾，2~60个字符
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

func (r *UpdateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "EventBusId")
	delete(f, "Enable")
	delete(f, "Description")
	delete(f, "EventPattern")
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRuleResponseParams `json:"Response"`
}

func (r *UpdateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTargetRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 事件目标ID
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 开启批量投递使能
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitnil,omitempty" name:"EnableBatchDelivery"`

	// 批量投递最长等待时间
	BatchTimeout *int64 `json:"BatchTimeout,omitnil,omitempty" name:"BatchTimeout"`

	// 批量投递最大事件条数
	BatchEventCount *int64 `json:"BatchEventCount,omitnil,omitempty" name:"BatchEventCount"`
}

type UpdateTargetRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 事件目标ID
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 开启批量投递使能
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitnil,omitempty" name:"EnableBatchDelivery"`

	// 批量投递最长等待时间
	BatchTimeout *int64 `json:"BatchTimeout,omitnil,omitempty" name:"BatchTimeout"`

	// 批量投递最大事件条数
	BatchEventCount *int64 `json:"BatchEventCount,omitnil,omitempty" name:"BatchEventCount"`
}

func (r *UpdateTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	delete(f, "TargetId")
	delete(f, "EnableBatchDelivery")
	delete(f, "BatchTimeout")
	delete(f, "BatchEventCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTargetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateTargetResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTargetResponseParams `json:"Response"`
}

func (r *UpdateTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTransformationRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转换器id
	TransformationId *string `json:"TransformationId,omitnil,omitempty" name:"TransformationId"`

	// 一个转换规则列表，当前仅限定一个
	Transformations []*Transformation `json:"Transformations,omitnil,omitempty" name:"Transformations"`
}

type UpdateTransformationRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil,omitempty" name:"EventBusId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转换器id
	TransformationId *string `json:"TransformationId,omitnil,omitempty" name:"TransformationId"`

	// 一个转换规则列表，当前仅限定一个
	Transformations []*Transformation `json:"Transformations,omitnil,omitempty" name:"Transformations"`
}

func (r *UpdateTransformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTransformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventBusId")
	delete(f, "RuleId")
	delete(f, "TransformationId")
	delete(f, "Transformations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTransformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTransformationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateTransformationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTransformationResponseParams `json:"Response"`
}

func (r *UpdateTransformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTransformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}