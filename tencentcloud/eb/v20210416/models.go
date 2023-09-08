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

package v20210416

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type APIGWParams struct {
	// HTTPS
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// POST
	Method *string `json:"Method,omitnil" name:"Method"`
}

// Predefined struct for user
type CheckRuleRequestParams struct {
	// Event信息
	Event *string `json:"Event,omitnil" name:"Event"`

	// EventPattern信息
	EventPattern *string `json:"EventPattern,omitnil" name:"EventPattern"`
}

type CheckRuleRequest struct {
	*tchttp.BaseRequest
	
	// Event信息
	Event *string `json:"Event,omitnil" name:"Event"`

	// EventPattern信息
	EventPattern *string `json:"EventPattern,omitnil" name:"EventPattern"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Input *string `json:"Input,omitnil" name:"Input"`

	// 一个转换规则列表
	Transformations []*Transformation `json:"Transformations,omitnil" name:"Transformations"`
}

type CheckTransformationRequest struct {
	*tchttp.BaseRequest
	
	// 待处理的json字符串
	Input *string `json:"Input,omitnil" name:"Input"`

	// 一个转换规则列表
	Transformations []*Transformation `json:"Transformations,omitnil" name:"Transformations"`
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
	Output *string `json:"Output,omitnil" name:"Output"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// ckafka资源qcs六段式
	ResourceDescription *string `json:"ResourceDescription,omitnil" name:"ResourceDescription"`
}

type CkafkaParams struct {
	// kafka offset
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// ckafka  topic
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`
}

type CkafkaTargetParams struct {
	// 要投递到的ckafka topic
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// 重试策略
	RetryPolicy *RetryPolicy `json:"RetryPolicy,omitnil" name:"RetryPolicy"`
}

type Connection struct {
	// 状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 更新时间
	ModTime *string `json:"ModTime,omitnil" name:"ModTime"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil" name:"AddTime"`

	// 连接器ID
	ConnectionId *string `json:"ConnectionId,omitnil" name:"ConnectionId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 连接器描述
	ConnectionDescription *ConnectionDescription `json:"ConnectionDescription,omitnil" name:"ConnectionDescription"`

	// 连接器名称
	ConnectionName *string `json:"ConnectionName,omitnil" name:"ConnectionName"`

	// 类型
	Type *string `json:"Type,omitnil" name:"Type"`
}

type ConnectionBrief struct {
	// 连接器类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 连接器状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`
}

type ConnectionDescription struct {
	// 资源qcs六段式，更多参考 [资源六段式](https://cloud.tencent.com/document/product/598/10606)
	ResourceDescription *string `json:"ResourceDescription,omitnil" name:"ResourceDescription"`

	// apigw参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	APIGWParams *APIGWParams `json:"APIGWParams,omitnil" name:"APIGWParams"`

	// ckafka参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CkafkaParams *CkafkaParams `json:"CkafkaParams,omitnil" name:"CkafkaParams"`

	// data transfer service (DTS)参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DTSParams *DTSParams `json:"DTSParams,omitnil" name:"DTSParams"`
}

// Predefined struct for user
type CreateConnectionRequestParams struct {
	// 连接器描述
	ConnectionDescription *ConnectionDescription `json:"ConnectionDescription,omitnil" name:"ConnectionDescription"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 连接器名称
	ConnectionName *string `json:"ConnectionName,omitnil" name:"ConnectionName"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 类型
	Type *string `json:"Type,omitnil" name:"Type"`
}

type CreateConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 连接器描述
	ConnectionDescription *ConnectionDescription `json:"ConnectionDescription,omitnil" name:"ConnectionDescription"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 连接器名称
	ConnectionName *string `json:"ConnectionName,omitnil" name:"ConnectionName"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 类型
	Type *string `json:"Type,omitnil" name:"Type"`
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
	ConnectionId *string `json:"ConnectionId,omitnil" name:"ConnectionId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 事件集名称，只能包含字母、数字、下划线、连字符，以字母开头，以数字或字母结尾，2~60个字符
	EventBusName *string `json:"EventBusName,omitnil" name:"EventBusName"`

	// 事件集描述，不限字符类型，200字符描述以内
	Description *string `json:"Description,omitnil" name:"Description"`

	// EB存储时长
	SaveDays *int64 `json:"SaveDays,omitnil" name:"SaveDays"`

	// EB是否开启存储
	EnableStore *bool `json:"EnableStore,omitnil" name:"EnableStore"`
}

type CreateEventBusRequest struct {
	*tchttp.BaseRequest
	
	// 事件集名称，只能包含字母、数字、下划线、连字符，以字母开头，以数字或字母结尾，2~60个字符
	EventBusName *string `json:"EventBusName,omitnil" name:"EventBusName"`

	// 事件集描述，不限字符类型，200字符描述以内
	Description *string `json:"Description,omitnil" name:"Description"`

	// EB存储时长
	SaveDays *int64 `json:"SaveDays,omitnil" name:"SaveDays"`

	// EB是否开启存储
	EnableStore *bool `json:"EnableStore,omitnil" name:"EnableStore"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventPattern *string `json:"EventPattern,omitnil" name:"EventPattern"`

	// 事件集ID。
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件集名称，只能包含字母、数字、下划线、连字符，以字母开头，以数字或字母结尾，2~60个字符
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 使能开关。
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 事件集描述，不限字符类型，200字符描述以内
	Description *string `json:"Description,omitnil" name:"Description"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 参考：[事件模式](https://cloud.tencent.com/document/product/1359/56084)
	EventPattern *string `json:"EventPattern,omitnil" name:"EventPattern"`

	// 事件集ID。
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件集名称，只能包含字母、数字、下划线、连字符，以字母开头，以数字或字母结尾，2~60个字符
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 使能开关。
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 事件集描述，不限字符类型，200字符描述以内
	Description *string `json:"Description,omitnil" name:"Description"`
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
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 目标类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 目标描述
	TargetDescription *TargetDescription `json:"TargetDescription,omitnil" name:"TargetDescription"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`
}

type CreateTargetRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 目标类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 目标描述
	TargetDescription *TargetDescription `json:"TargetDescription,omitnil" name:"TargetDescription"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTargetResponseParams struct {
	// 目标ID
	TargetId *string `json:"TargetId,omitnil" name:"TargetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 规则id
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 一个转换规则列表，当前仅限定一个
	Transformations []*Transformation `json:"Transformations,omitnil" name:"Transformations"`
}

type CreateTransformationRequest struct {
	*tchttp.BaseRequest
	
	// 事件总线 id
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 规则id
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 一个转换规则列表，当前仅限定一个
	Transformations []*Transformation `json:"Transformations,omitnil" name:"Transformations"`
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
	TransformationId *string `json:"TransformationId,omitnil" name:"TransformationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

}

type DeadLetterConfig struct {
	// 支持dlq、丢弃、忽略错误继续传递三种模式, 分别对应: DLQ,DROP,IGNORE_ERROR
	DisposeMethod *string `json:"DisposeMethod,omitnil" name:"DisposeMethod"`

	// 设置了DLQ方式后,此选项必填. 错误消息会被投递到对应的kafka topic中
	// 注意：此字段可能返回 null，表示取不到有效值。
	CkafkaDeliveryParams *CkafkaDeliveryParams `json:"CkafkaDeliveryParams,omitnil" name:"CkafkaDeliveryParams"`
}

// Predefined struct for user
type DeleteConnectionRequestParams struct {
	// 连接器ID
	ConnectionId *string `json:"ConnectionId,omitnil" name:"ConnectionId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`
}

type DeleteConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 连接器ID
	ConnectionId *string `json:"ConnectionId,omitnil" name:"ConnectionId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`
}

type DeleteEventBusRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件目标ID
	TargetId *string `json:"TargetId,omitnil" name:"TargetId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`
}

type DeleteTargetRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件目标ID
	TargetId *string `json:"TargetId,omitnil" name:"TargetId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 转换器id
	TransformationId *string `json:"TransformationId,omitnil" name:"TransformationId"`
}

type DeleteTransformationRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 转换器id
	TransformationId *string `json:"TransformationId,omitnil" name:"TransformationId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 聚合字段
	GroupField *string `json:"GroupField,omitnil" name:"GroupField"`

	// 页数
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 每页数据大小
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 筛选条件
	Filter []*LogFilter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeLogTagValueRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 聚合字段
	GroupField *string `json:"GroupField,omitnil" name:"GroupField"`

	// 页数
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 每页数据大小
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 筛选条件
	Filter []*LogFilter `json:"Filter,omitnil" name:"Filter"`
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
	// 索引检索维度值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*string `json:"Results,omitnil" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	NetMode *string `json:"NetMode,omitnil" name:"NetMode"`

	// 索引前缀
	IndexPrefix *string `json:"IndexPrefix,omitnil" name:"IndexPrefix"`

	// es日志轮换粒度
	RotationInterval *string `json:"RotationInterval,omitnil" name:"RotationInterval"`

	// DTS事件配置
	OutputMode *string `json:"OutputMode,omitnil" name:"OutputMode"`

	// DTS索引配置
	IndexSuffixMode *string `json:"IndexSuffixMode,omitnil" name:"IndexSuffixMode"`

	// es模版类型
	IndexTemplateType *string `json:"IndexTemplateType,omitnil" name:"IndexTemplateType"`
}

type EtlFilter struct {
	// 语法Rule规则保持一致
	Filter *string `json:"Filter,omitnil" name:"Filter"`
}

type Event struct {
	// 事件源的信息,新产品上报必须符合EB的规范
	Source *string `json:"Source,omitnil" name:"Source"`

	// 事件数据，内容由创建事件的系统来控制，当前datacontenttype仅支持application/json;charset=utf-8，所以该字段是json字符串
	Data *string `json:"Data,omitnil" name:"Data"`

	// 事件类型，可自定义，选填。云服务默认写 COS:Created:PostObject，用“：”分割类型字段
	Type *string `json:"Type,omitnil" name:"Type"`

	// 事件来源详细描述，可自定义，选填。云服务默认为标准qcs资源表示语法：qcs::dts:ap-guangzhou:appid/uin:xxx
	Subject *string `json:"Subject,omitnil" name:"Subject"`

	// 事件发生的毫秒时间戳，
	// time.Now().UnixNano()/1e6
	Time *int64 `json:"Time,omitnil" name:"Time"`
}

type EventBus struct {
	// 更新时间
	ModTime *string `json:"ModTime,omitnil" name:"ModTime"`

	// 事件集描述，不限字符类型，200字符描述以内
	Description *string `json:"Description,omitnil" name:"Description"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil" name:"AddTime"`

	// 事件集名称，只能包含字母、数字、下划线、连字符，以字母开头，以数字或字母结尾，2~60个字符
	EventBusName *string `json:"EventBusName,omitnil" name:"EventBusName"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件集类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 计费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 连接器基础信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionBriefs []*ConnectionBrief `json:"ConnectionBriefs,omitnil" name:"ConnectionBriefs"`

	// 目标简要信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetBriefs []*TargetBrief `json:"TargetBriefs,omitnil" name:"TargetBriefs"`
}

type Extraction struct {
	// JsonPath, 不指定则使用默认值$.
	ExtractionInputPath *string `json:"ExtractionInputPath,omitnil" name:"ExtractionInputPath"`

	// 取值: TEXT/JSON
	Format *string `json:"Format,omitnil" name:"Format"`

	// 仅在Text需要传递
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextParams *TextParams `json:"TextParams,omitnil" name:"TextParams"`
}

type Filter struct {
	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 过滤键的名称。
	Name *string `json:"Name,omitnil" name:"Name"`
}

// Predefined struct for user
type GetEventBusRequestParams struct {
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`
}

type GetEventBusRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`
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
	ModTime *string `json:"ModTime,omitnil" name:"ModTime"`

	// 事件集描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 日志主题ID
	ClsTopicId *string `json:"ClsTopicId,omitnil" name:"ClsTopicId"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil" name:"AddTime"`

	// 日志集ID
	ClsLogsetId *string `json:"ClsLogsetId,omitnil" name:"ClsLogsetId"`

	// 事件集名称
	EventBusName *string `json:"EventBusName,omitnil" name:"EventBusName"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// （已废弃）事件集类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 计费模式
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// EB日志存储时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	SaveDays *int64 `json:"SaveDays,omitnil" name:"SaveDays"`

	// EB日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTopicId *string `json:"LogTopicId,omitnil" name:"LogTopicId"`

	// 是否开启存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableStore *bool `json:"EnableStore,omitnil" name:"EnableStore"`

	// 消息序列，是否有序
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkMode *string `json:"LinkMode,omitnil" name:"LinkMode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventType *string `json:"EventType,omitnil" name:"EventType"`
}

type GetPlatformEventTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 平台产品事件类型
	EventType *string `json:"EventType,omitnil" name:"EventType"`
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
	EventTemplate *string `json:"EventTemplate,omitnil" name:"EventTemplate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`
}

type GetRuleRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件规则id
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 事件规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 事件规则状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 事件规则描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 事件模式
	EventPattern *string `json:"EventPattern,omitnil" name:"EventPattern"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil" name:"AddTime"`

	// 更新时间
	ModTime *string `json:"ModTime,omitnil" name:"ModTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 转换器id
	TransformationId *string `json:"TransformationId,omitnil" name:"TransformationId"`
}

type GetTransformationRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 转换器id
	TransformationId *string `json:"TransformationId,omitnil" name:"TransformationId"`
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
	Transformations []*Transformation `json:"Transformations,omitnil" name:"Transformations"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 根据哪个字段进行返回结果排序，目前支持如下以下字段：AddTime, ModTime
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 以升序还是降序的方式返回结果，可选值 ASC 和 DESC
	Order *string `json:"Order,omitnil" name:"Order"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type ListConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 根据哪个字段进行返回结果排序，目前支持如下以下字段：AddTime, ModTime
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 以升序还是降序的方式返回结果，可选值 ASC 和 DESC
	Order *string `json:"Order,omitnil" name:"Order"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
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
	Connections []*Connection `json:"Connections,omitnil" name:"Connections"`

	// 连接器总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 根据哪个字段进行返回结果排序,支持以下字段：AddTime（创建时间）, ModTime（修改时间）
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 以升序还是降序的方式返回结果，可选值 ASC（升序） 和 DESC（降序）
	Order *string `json:"Order,omitnil" name:"Order"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 分页偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type ListEventBusesRequest struct {
	*tchttp.BaseRequest
	
	// 根据哪个字段进行返回结果排序,支持以下字段：AddTime（创建时间）, ModTime（修改时间）
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 以升序还是降序的方式返回结果，可选值 ASC（升序） 和 DESC（降序）
	Order *string `json:"Order,omitnil" name:"Order"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 分页偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
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
	EventBuses []*EventBus `json:"EventBuses,omitnil" name:"EventBuses"`

	// 事件集总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductType *string `json:"ProductType,omitnil" name:"ProductType"`
}

type ListPlatformEventNamesRequest struct {
	*tchttp.BaseRequest
	
	// 平台产品类型
	ProductType *string `json:"ProductType,omitnil" name:"ProductType"`
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
	EventNames []*PlatformEventDetail `json:"EventNames,omitnil" name:"EventNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProductType *string `json:"ProductType,omitnil" name:"ProductType"`
}

type ListPlatformEventPatternsRequest struct {
	*tchttp.BaseRequest
	
	// 平台产品类型
	ProductType *string `json:"ProductType,omitnil" name:"ProductType"`
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
	EventPatterns []*PlatformEventSummary `json:"EventPatterns,omitnil" name:"EventPatterns"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	PlatformProducts []*PlatformProduct `json:"PlatformProducts,omitnil" name:"PlatformProducts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 根据哪个字段进行返回结果排序,支持以下字段：AddTime（创建时间）, ModTime（修改时间）
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 以升序还是降序的方式返回结果，可选值 ASC（升序） 和 DESC（降序）
	Order *string `json:"Order,omitnil" name:"Order"`
}

type ListRulesRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 根据哪个字段进行返回结果排序,支持以下字段：AddTime（创建时间）, ModTime（修改时间）
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 以升序还是降序的方式返回结果，可选值 ASC（升序） 和 DESC（降序）
	Order *string `json:"Order,omitnil" name:"Order"`
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
	Rules []*Rule `json:"Rules,omitnil" name:"Rules"`

	// 事件规则总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 根据哪个字段进行返回结果排序,支持以下字段：AddTime（创建时间）, ModTime（修改时间）
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 以升序还是降序的方式返回结果，可选值 ASC（升序） 和 DESC（降序）
	Order *string `json:"Order,omitnil" name:"Order"`
}

type ListTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 根据哪个字段进行返回结果排序,支持以下字段：AddTime（创建时间）, ModTime（修改时间）
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 以升序还是降序的方式返回结果，可选值 ASC（升序） 和 DESC（降序）
	Order *string `json:"Order,omitnil" name:"Order"`
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
	delete(f, "OrderBy")
	delete(f, "RuleId")
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 目标信息
	Targets []*Target `json:"Targets,omitnil" name:"Targets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Key *string `json:"Key,omitnil" name:"Key"`

	// 运算符，全等 eq，不等 neq，相似 like，排除相似 not like,  小于 lt，小于且等于 lte，大于 gt，大于且等于 gte，在范围内 range，不在范围内 norange
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 过滤值,范围运算需要同时输入两个值，以英文逗号分隔
	Value *string `json:"Value,omitnil" name:"Value"`

	// 该层级filters逻辑关系，取值 "AND" 或 "OR"
	Type *string `json:"Type,omitnil" name:"Type"`

	// LogFilters数组
	Filters []*LogFilters `json:"Filters,omitnil" name:"Filters"`
}

type LogFilters struct {
	// 过滤字段名称
	Key *string `json:"Key,omitnil" name:"Key"`

	// 运算符, 全等 eq，不等 neq，相似 like，排除相似 not like,  小于 lt，小于且等于 lte，大于 gt，大于且等于 gte，在范围内 range，不在范围内 norange
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 过滤值，范围运算需要同时输入两个值，以英文逗号分隔
	Value *string `json:"Value,omitnil" name:"Value"`
}

type OutputStructParam struct {
	// 对应输出json中的key
	Key *string `json:"Key,omitnil" name:"Key"`

	// 可以填json-path也可以支持常量或者内置关键字date类型
	Value *string `json:"Value,omitnil" name:"Value"`

	// value的数据类型, 可选值: STRING, NUMBER,BOOLEAN,NULL,SYS_VARIABLE,JSONPATH
	ValueType *string `json:"ValueType,omitnil" name:"ValueType"`
}

type PlatformEventDetail struct {
	// 事件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventName *string `json:"EventName,omitnil" name:"EventName"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventType *string `json:"EventType,omitnil" name:"EventType"`
}

type PlatformEventSummary struct {
	// 平台事件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventName *string `json:"EventName,omitnil" name:"EventName"`

	// 平台事件匹配规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventPattern *string `json:"EventPattern,omitnil" name:"EventPattern"`
}

type PlatformProduct struct {
	// 平台产品名称
	ProductName *string `json:"ProductName,omitnil" name:"ProductName"`

	// 平台产品类型
	ProductType *string `json:"ProductType,omitnil" name:"ProductType"`
}

// Predefined struct for user
type PublishEventRequestParams struct {
	// 事件列表
	EventList []*Event `json:"EventList,omitnil" name:"EventList"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`
}

type PublishEventRequest struct {
	*tchttp.BaseRequest
	
	// 事件列表
	EventList []*Event `json:"EventList,omitnil" name:"EventList"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventList []*Event `json:"EventList,omitnil" name:"EventList"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`
}

type PutEventsRequest struct {
	*tchttp.BaseRequest
	
	// 事件列表
	EventList []*Event `json:"EventList,omitnil" name:"EventList"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RetryInterval *uint64 `json:"RetryInterval,omitnil" name:"RetryInterval"`

	// 最大重试次数
	MaxRetryAttempts *uint64 `json:"MaxRetryAttempts,omitnil" name:"MaxRetryAttempts"`
}

type Rule struct {
	// 状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 修改时间
	ModTime *string `json:"ModTime,omitnil" name:"ModTime"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil" name:"AddTime"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// Target 简要信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*TargetBrief `json:"Targets,omitnil" name:"Targets"`

	// rule设置的dlq规则. 可能为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterConfig *DeadLetterConfig `json:"DeadLetterConfig,omitnil" name:"DeadLetterConfig"`
}

type SCFParams struct {
	// 批量投递最长等待时间
	BatchTimeout *int64 `json:"BatchTimeout,omitnil" name:"BatchTimeout"`

	// 批量投递最大事件条数
	BatchEventCount *int64 `json:"BatchEventCount,omitnil" name:"BatchEventCount"`

	// 开启批量投递使能
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitnil" name:"EnableBatchDelivery"`
}

// Predefined struct for user
type SearchLogRequestParams struct {
	// 起始时间unix 毫秒时间戳
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间unix 毫秒时间戳
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 页码
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 每页数据大小
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 筛选条件
	Filter []*LogFilter `json:"Filter,omitnil" name:"Filter"`

	// 排序数组
	OrderFields []*string `json:"OrderFields,omitnil" name:"OrderFields"`

	// 排序方式，asc 从旧到新，desc 从新到旧
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`
}

type SearchLogRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间unix 毫秒时间戳
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间unix 毫秒时间戳
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 页码
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 每页数据大小
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 筛选条件
	Filter []*LogFilter `json:"Filter,omitnil" name:"Filter"`

	// 排序数组
	OrderFields []*string `json:"OrderFields,omitnil" name:"OrderFields"`

	// 排序方式，asc 从旧到新，desc 从新到旧
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`
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
	// 日志总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 每页日志条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 日志检索结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*SearchLogResult `json:"Results,omitnil" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Timestamp *string `json:"Timestamp,omitnil" name:"Timestamp"`

	// 日志内容详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 事件来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil" name:"Source"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 事件匹配规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleIds *string `json:"RuleIds,omitnil" name:"RuleIds"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subject *string `json:"Subject,omitnil" name:"Subject"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 事件状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`
}

type Target struct {
	// 目标类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 目标ID
	TargetId *string `json:"TargetId,omitnil" name:"TargetId"`

	// 目标描述
	TargetDescription *TargetDescription `json:"TargetDescription,omitnil" name:"TargetDescription"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 开启批量投递使能
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitnil" name:"EnableBatchDelivery"`

	// 批量投递最长等待时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchTimeout *int64 `json:"BatchTimeout,omitnil" name:"BatchTimeout"`

	// 批量投递最大事件条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchEventCount *int64 `json:"BatchEventCount,omitnil" name:"BatchEventCount"`
}

type TargetBrief struct {
	// 目标ID
	TargetId *string `json:"TargetId,omitnil" name:"TargetId"`

	// 目标类型
	Type *string `json:"Type,omitnil" name:"Type"`
}

type TargetDescription struct {
	// QCS资源六段式，更多参考 [资源六段式](https://cloud.tencent.com/document/product/598/10606)
	ResourceDescription *string `json:"ResourceDescription,omitnil" name:"ResourceDescription"`

	// 云函数参数
	SCFParams *SCFParams `json:"SCFParams,omitnil" name:"SCFParams"`

	// Ckafka参数
	CkafkaTargetParams *CkafkaTargetParams `json:"CkafkaTargetParams,omitnil" name:"CkafkaTargetParams"`

	// ElasticSearch参数
	ESTargetParams *ESTargetParams `json:"ESTargetParams,omitnil" name:"ESTargetParams"`
}

type TextParams struct {
	// 逗号、| 、制表符、空格、换行符、%、#，限制长度为 1。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Separator *string `json:"Separator,omitnil" name:"Separator"`

	// 填写正则表达式：长度128
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regex *string `json:"Regex,omitnil" name:"Regex"`
}

type Transform struct {
	// 描述如何数据转换
	OutputStructs []*OutputStructParam `json:"OutputStructs,omitnil" name:"OutputStructs"`
}

type Transformation struct {
	// 描述如何提取数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extraction *Extraction `json:"Extraction,omitnil" name:"Extraction"`

	// 描述如何过滤数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	EtlFilter *EtlFilter `json:"EtlFilter,omitnil" name:"EtlFilter"`

	// 描述如何数据转换
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transform *Transform `json:"Transform,omitnil" name:"Transform"`
}

// Predefined struct for user
type UpdateConnectionRequestParams struct {
	// 连接器ID
	ConnectionId *string `json:"ConnectionId,omitnil" name:"ConnectionId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 连接器名称
	ConnectionName *string `json:"ConnectionName,omitnil" name:"ConnectionName"`
}

type UpdateConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 连接器ID
	ConnectionId *string `json:"ConnectionId,omitnil" name:"ConnectionId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 使能开关
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 连接器名称
	ConnectionName *string `json:"ConnectionName,omitnil" name:"ConnectionName"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件集描述，不限字符类型，200字符描述以内
	Description *string `json:"Description,omitnil" name:"Description"`

	// 事件集名称，只能包含字母、数字、下划线、连字符，以字母开头，以数字或字母结尾，2~60个字符
	EventBusName *string `json:"EventBusName,omitnil" name:"EventBusName"`

	// EB日志存储时长
	SaveDays *int64 `json:"SaveDays,omitnil" name:"SaveDays"`

	// EB日志主题ID
	LogTopicId *string `json:"LogTopicId,omitnil" name:"LogTopicId"`

	// 是否开启存储
	EnableStore *bool `json:"EnableStore,omitnil" name:"EnableStore"`
}

type UpdateEventBusRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件集描述，不限字符类型，200字符描述以内
	Description *string `json:"Description,omitnil" name:"Description"`

	// 事件集名称，只能包含字母、数字、下划线、连字符，以字母开头，以数字或字母结尾，2~60个字符
	EventBusName *string `json:"EventBusName,omitnil" name:"EventBusName"`

	// EB日志存储时长
	SaveDays *int64 `json:"SaveDays,omitnil" name:"SaveDays"`

	// EB日志主题ID
	LogTopicId *string `json:"LogTopicId,omitnil" name:"LogTopicId"`

	// 是否开启存储
	EnableStore *bool `json:"EnableStore,omitnil" name:"EnableStore"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 使能开关。
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 规则描述，不限字符类型，200字符描述以内。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 参考：[事件模式](https://cloud.tencent.com/document/product/1359/56084)
	EventPattern *string `json:"EventPattern,omitnil" name:"EventPattern"`

	// 事件规则名称，只能包含字母、数字、下划线、连字符，以字母开头，以数字或字母结尾，2~60个字符
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`
}

type UpdateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 使能开关。
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 规则描述，不限字符类型，200字符描述以内。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 参考：[事件模式](https://cloud.tencent.com/document/product/1359/56084)
	EventPattern *string `json:"EventPattern,omitnil" name:"EventPattern"`

	// 事件规则名称，只能包含字母、数字、下划线、连字符，以字母开头，以数字或字母结尾，2~60个字符
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 事件目标ID
	TargetId *string `json:"TargetId,omitnil" name:"TargetId"`

	// 开启批量投递使能
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitnil" name:"EnableBatchDelivery"`

	// 批量投递最长等待时间
	BatchTimeout *int64 `json:"BatchTimeout,omitnil" name:"BatchTimeout"`

	// 批量投递最大事件条数
	BatchEventCount *int64 `json:"BatchEventCount,omitnil" name:"BatchEventCount"`
}

type UpdateTargetRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 事件规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 事件目标ID
	TargetId *string `json:"TargetId,omitnil" name:"TargetId"`

	// 开启批量投递使能
	EnableBatchDelivery *bool `json:"EnableBatchDelivery,omitnil" name:"EnableBatchDelivery"`

	// 批量投递最长等待时间
	BatchTimeout *int64 `json:"BatchTimeout,omitnil" name:"BatchTimeout"`

	// 批量投递最大事件条数
	BatchEventCount *int64 `json:"BatchEventCount,omitnil" name:"BatchEventCount"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 转换器id
	TransformationId *string `json:"TransformationId,omitnil" name:"TransformationId"`

	// 一个转换规则列表，当前仅限定一个
	Transformations []*Transformation `json:"Transformations,omitnil" name:"Transformations"`
}

type UpdateTransformationRequest struct {
	*tchttp.BaseRequest
	
	// 事件集ID
	EventBusId *string `json:"EventBusId,omitnil" name:"EventBusId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 转换器id
	TransformationId *string `json:"TransformationId,omitnil" name:"TransformationId"`

	// 一个转换规则列表，当前仅限定一个
	Transformations []*Transformation `json:"Transformations,omitnil" name:"Transformations"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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