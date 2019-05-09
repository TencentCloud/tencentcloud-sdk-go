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

package v20180614

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Attribute struct {

	// 属性列表
	Tags []*DeviceTag `json:"Tags,omitempty" name:"Tags" list`
}

type BatchPublishMessage struct {

	// 消息发往的主题。为 Topic 权限中去除 ProductID 和 DeviceName 的部分，如 “event”
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 消息内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`
}

type BatchUpdateShadow struct {

	// 设备影子的期望状态，格式为 Json 对象序列化之后的字符串
	Desired *string `json:"Desired,omitempty" name:"Desired"`
}

type BrokerSubscribe struct {

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type CancelTaskRequest struct {
	*tchttp.BaseRequest

	// 任务 ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *CancelTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceRequest struct {
	*tchttp.BaseRequest

	// 产品 ID 。创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称。命名规则：[a-zA-Z0-9:_-]{1,48}。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备属性
	Attribute *Attribute `json:"Attribute,omitempty" name:"Attribute"`

	// 是否使用自定义PSK，默认不使用
	DefinedPsk *string `json:"DefinedPsk,omitempty" name:"DefinedPsk"`

	// 运营商类型，当产品是NB-IoT产品时，此字段必填。1表示中国电信，2表示中国移动，3表示中国联通
	Isp *uint64 `json:"Isp,omitempty" name:"Isp"`

	// IMEI，当产品是NB-IoT产品时，此字段必填
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// LoRa设备的DevEui，当创建LoRa时，此字段必填
	LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

	// LoRa设备的MoteType
	LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

	// 创建LoRa设备需要skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// LoRa设备的AppKey
	LoraAppKey *string `json:"LoraAppKey,omitempty" name:"LoraAppKey"`
}

func (r *CreateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备名称
		DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

		// 对称加密密钥，base64编码。采用对称加密时返回该参数
		DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

		// 设备证书，用于 TLS 建立链接时校验客户端身份。采用非对称加密时返回该参数
		DeviceCert *string `json:"DeviceCert,omitempty" name:"DeviceCert"`

		// 设备私钥，用于 TLS 建立链接时校验客户端身份，腾讯云后台不保存，请妥善保管。采用非对称加密时返回该参数
		DevicePrivateKey *string `json:"DevicePrivateKey,omitempty" name:"DevicePrivateKey"`

		// LoRa设备的DevEui，当设备是LoRa设备时，会返回该字段
		LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

		// LoRa设备的MoteType，当设备是LoRa设备时，会返回该字段
		LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

		// LoRa设备的AppKey，当设备是LoRa设备时，会返回该字段
		LoraAppKey *string `json:"LoraAppKey,omitempty" name:"LoraAppKey"`

		// LoRa设备的NwkKey，当设备是LoRa设备时，会返回该字段
		LoraNwkKey *string `json:"LoraNwkKey,omitempty" name:"LoraNwkKey"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLoraDeviceRequest struct {
	*tchttp.BaseRequest

	// 产品 ID ，创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备类型 ，目前支持A、B、C三种
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// LoRa应用UUID
	AppEui *string `json:"AppEui,omitempty" name:"AppEui"`

	// LoRa设备UUID
	DeviceEui *string `json:"DeviceEui,omitempty" name:"DeviceEui"`

	// LoRa应用密钥
	AppKey *string `json:"AppKey,omitempty" name:"AppKey"`

	// LoRa设备验证密钥
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`

	// 设备备注
	Memo *string `json:"Memo,omitempty" name:"Memo"`
}

func (r *CreateLoraDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLoraDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLoraDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// LoRa应用UUID
		AppEui *string `json:"AppEui,omitempty" name:"AppEui"`

		// LoRa设备UUID
		DeviceEui *string `json:"DeviceEui,omitempty" name:"DeviceEui"`

		// 设备类型,目前支持A、B、C三种
		ClassType *string `json:"ClassType,omitempty" name:"ClassType"`

		// 设备名称
		DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLoraDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLoraDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMultiDeviceRequest struct {
	*tchttp.BaseRequest

	// 产品 ID。创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 批量创建的设备名数组，单次最多创建 100 个设备。命名规则：[a-zA-Z0-9:_-]{1,48}
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames" list`
}

func (r *CreateMultiDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMultiDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMultiDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID，腾讯云生成全局唯一的任务 ID，有效期一个月，一个月之后任务失效。可以调用获取创建多设备任务状态接口获取该任务的执行状态，当状态为成功时，可以调用获取创建多设备任务结果接口获取该任务的结果
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMultiDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMultiDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProductRequest struct {
	*tchttp.BaseRequest

	// 产品名称，名称不能和已经存在的产品名称重复。命名规则：[a-zA-Z0-9:_-]{1,32}
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品属性
	ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`

	// 创建LoRa产品需要的Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *CreateProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品名称
		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

		// 产品 ID，腾讯云生成全局唯一 ID
		ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

		// 产品属性
		ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest

	// 任务类型，取值为 “UpdateShadow” 或者 “PublishMessage”
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 执行任务的产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 执行任务的设备名的正则表达式
	DeviceNameFilter *string `json:"DeviceNameFilter,omitempty" name:"DeviceNameFilter"`

	// 任务开始执行的时间。 取值为 Unix 时间戳，单位秒，且需大于等于当前时间时间戳，0为系统当前时间时间戳，即立即执行，最大为当前时间86400秒后，超过则取值为当前时间86400秒后
	ScheduleTimeInSeconds *uint64 `json:"ScheduleTimeInSeconds,omitempty" name:"ScheduleTimeInSeconds"`

	// 任务描述细节，描述见下 Task
	Tasks *Task `json:"Tasks,omitempty" name:"Tasks"`

	// 最长执行时间，单位秒，被调度后超过此时间仍未有结果则视为任务失败。取值为0-86400，默认为86400
	MaxExecutionTimeInSeconds *uint64 `json:"MaxExecutionTimeInSeconds,omitempty" name:"MaxExecutionTimeInSeconds"`
}

func (r *CreateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicPolicyRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Topic权限，1发布，2订阅，3订阅和发布
	Privilege *uint64 `json:"Privilege,omitempty" name:"Privilege"`

	// 代理订阅信息
	BrokerSubscribe *BrokerSubscribe `json:"BrokerSubscribe,omitempty" name:"BrokerSubscribe"`
}

func (r *CreateTopicPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRuleRequest struct {
	*tchttp.BaseRequest

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则内容
	TopicRulePayload *TopicRulePayload `json:"TopicRulePayload,omitempty" name:"TopicRulePayload"`
}

func (r *CreateTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest

	// 设备所属的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 需要删除的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 删除LoRa设备以及LoRa网关设备需要skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *DeleteDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLoraDeviceRequest struct {
	*tchttp.BaseRequest

	// 设备所属产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DeleteLoraDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLoraDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLoraDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoraDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLoraDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProductRequest struct {
	*tchttp.BaseRequest

	// 需要删除的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 删除LoRa产品需要skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *DeleteProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicRuleRequest struct {
	*tchttp.BaseRequest

	// 规则名
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *DeleteTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DescribeDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备名
		DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

		// 设备是否在线，0不在线，1在线
		Online *uint64 `json:"Online,omitempty" name:"Online"`

		// 设备登录时间
		LoginTime *uint64 `json:"LoginTime,omitempty" name:"LoginTime"`

		// 设备固件版本
		Version *string `json:"Version,omitempty" name:"Version"`

		// 设备最后更新时间
		LastUpdateTime *uint64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

		// 设备证书
		DeviceCert *string `json:"DeviceCert,omitempty" name:"DeviceCert"`

		// 设备密钥
		DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

		// 设备属性
		Tags []*DeviceTag `json:"Tags,omitempty" name:"Tags" list`

		// 设备类型
		DeviceType *uint64 `json:"DeviceType,omitempty" name:"DeviceType"`

		// IMEI
		Imei *string `json:"Imei,omitempty" name:"Imei"`

		// 运营商类型
		Isp *uint64 `json:"Isp,omitempty" name:"Isp"`

		// IP地址
		ConnIP *uint64 `json:"ConnIP,omitempty" name:"ConnIP"`

		// NB IoT运营商处的DeviceID
		NbiotDeviceID *string `json:"NbiotDeviceID,omitempty" name:"NbiotDeviceID"`

		// Lora设备的dev eui
		LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

		// Lora设备的mote type
		LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

		// 设备的sdk日志等级
	// 注意：此字段可能返回 null，表示取不到有效值。
		LogLevel *uint64 `json:"LogLevel,omitempty" name:"LogLevel"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceShadowRequest struct {
	*tchttp.BaseRequest

	// 产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称。命名规则：[a-zA-Z0-9:_-]{1,48}
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DescribeDeviceShadowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceShadowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceShadowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备影子数据
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceShadowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceShadowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest

	// 需要查看设备列表的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 设备固件版本号，若不带此参数会返回所有固件版本的设备
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDevicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 设备详细信息列表
		Devices []*DeviceInfo `json:"Devices,omitempty" name:"Devices" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDevicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoraDeviceRequest struct {
	*tchttp.BaseRequest

	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DescribeLoraDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoraDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLoraDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备名称
		DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

		// LoRa应用UUID
		AppEui *string `json:"AppEui,omitempty" name:"AppEui"`

		// LoRa设备UUID
		DeviceEui *string `json:"DeviceEui,omitempty" name:"DeviceEui"`

		// LoRa应用密钥
		AppKey *string `json:"AppKey,omitempty" name:"AppKey"`

		// 设备类型,目前支持A、B、C三种
		ClassType *string `json:"ClassType,omitempty" name:"ClassType"`

		// 设备所属产品id
		ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoraDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLoraDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMultiDevTaskRequest struct {
	*tchttp.BaseRequest

	// 任务 ID，由批量创建设备接口返回
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 产品 ID，创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeMultiDevTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMultiDevTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMultiDevTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 任务是否完成。0 代表任务未开始，1 代表任务正在执行，2 代表任务已完成
		TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMultiDevTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMultiDevTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMultiDevicesRequest struct {
	*tchttp.BaseRequest

	// 产品 ID，创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 任务 ID，由批量创建设备接口返回
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，每页返回的设备个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMultiDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMultiDevicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMultiDevicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 ID，由批量创建设备接口返回
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 设备详细信息列表
		DevicesInfo []*MultiDevicesInfo `json:"DevicesInfo,omitempty" name:"DevicesInfo" list`

		// 该任务创建设备的总数
		TotalDevNum *uint64 `json:"TotalDevNum,omitempty" name:"TotalDevNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMultiDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMultiDevicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductsRequest struct {
	*tchttp.BaseRequest

	// 分页偏移，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，当前页面中显示的最大数量，值范围 10-250。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 产品详细信息列表
		Products []*ProductInfo `json:"Products,omitempty" name:"Products" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务类型，目前取值为 “UpdateShadow” 或者 “PublishMessage”
		Type *string `json:"Type,omitempty" name:"Type"`

		// 任务 ID
		Id *string `json:"Id,omitempty" name:"Id"`

		// 产品 ID
		ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

		// 状态。1表示等待处理，2表示调度处理中，3表示已完成，4表示失败，5表示已取消
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 任务创建时间，Unix 时间戳
		CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

		// 最后任务更新时间，Unix 时间戳
		UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 任务完成时间，Unix 时间戳
		DoneTime *uint64 `json:"DoneTime,omitempty" name:"DoneTime"`

		// 被调度时间，Unix 时间戳
		ScheduleTime *uint64 `json:"ScheduleTime,omitempty" name:"ScheduleTime"`

		// 返回的错误码
		RetCode *uint64 `json:"RetCode,omitempty" name:"RetCode"`

		// 返回的错误信息
		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

		// 完成任务的设备比例
		Percent *uint64 `json:"Percent,omitempty" name:"Percent"`

		// 匹配到的需执行任务的设备数目
		AllDeviceCnt *uint64 `json:"AllDeviceCnt,omitempty" name:"AllDeviceCnt"`

		// 已完成任务的设备数目
		DoneDeviceCnt *uint64 `json:"DoneDeviceCnt,omitempty" name:"DoneDeviceCnt"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 分页偏移，从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 1-250
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户一个月内创建的任务总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 此页任务对象的数组，按创建时间排序
		Tasks []*TaskInfo `json:"Tasks,omitempty" name:"Tasks" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeviceInfo struct {

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备是否在线，0不在线，1在线
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// 设备登录时间
	LoginTime *uint64 `json:"LoginTime,omitempty" name:"LoginTime"`

	// 设备版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 设备证书，证书加密的设备返回
	DeviceCert *string `json:"DeviceCert,omitempty" name:"DeviceCert"`

	// 设备密钥，密钥加密的设备返回
	DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

	// 设备属性
	Tags []*DeviceTag `json:"Tags,omitempty" name:"Tags" list`

	// 设备类型
	DeviceType *uint64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// IMEI
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// 运营商类型
	Isp *uint64 `json:"Isp,omitempty" name:"Isp"`

	// NB IOT运营商处的DeviceID
	NbiotDeviceID *string `json:"NbiotDeviceID,omitempty" name:"NbiotDeviceID"`

	// IP地址
	ConnIP *uint64 `json:"ConnIP,omitempty" name:"ConnIP"`

	// 设备最后更新时间
	LastUpdateTime *uint64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// LoRa设备的dev eui
	LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

	// LoRa设备的Mote type
	LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

	// 首次上线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstOnlineTime *uint64 `json:"FirstOnlineTime,omitempty" name:"FirstOnlineTime"`

	// 最近下线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOfflineTime *uint64 `json:"LastOfflineTime,omitempty" name:"LastOfflineTime"`

	// 设备创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 设备日志级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogLevel *uint64 `json:"LogLevel,omitempty" name:"LogLevel"`
}

type DeviceTag struct {

	// 属性名称
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 属性值的类型，1 int，2 string
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 属性的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DisableTopicRuleRequest struct {
	*tchttp.BaseRequest

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *DisableTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableTopicRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableTopicRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableTopicRuleRequest struct {
	*tchttp.BaseRequest

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *EnableTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableTopicRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableTopicRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 过滤键的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type MultiDevicesInfo struct {

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 对称加密密钥，base64 编码，采用对称加密时返回该参数
	DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

	// 设备证书，采用非对称加密时返回该参数
	DeviceCert *string `json:"DeviceCert,omitempty" name:"DeviceCert"`

	// 设备私钥，采用非对称加密时返回该参数，腾讯云为用户缓存起来，其生命周期与任务生命周期一致
	DevicePrivateKey *string `json:"DevicePrivateKey,omitempty" name:"DevicePrivateKey"`

	// 错误码
	Result *uint64 `json:"Result,omitempty" name:"Result"`

	// 错误信息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type ProductInfo struct {

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品元数据
	ProductMetadata *ProductMetadata `json:"ProductMetadata,omitempty" name:"ProductMetadata"`

	// 产品属性
	ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`
}

type ProductMetadata struct {

	// 产品创建时间
	CreationDate *uint64 `json:"CreationDate,omitempty" name:"CreationDate"`
}

type ProductProperties struct {

	// 产品描述
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 加密类型，1表示证书认证，2表示签名认证。如不填写，默认值是1
	EncryptionType *string `json:"EncryptionType,omitempty" name:"EncryptionType"`

	// 产品所属区域，目前只支持广州（gz）
	Region *string `json:"Region,omitempty" name:"Region"`

	// 产品类型，各个类型值代表的节点-类型如下：
	// 0 普通产品，2 NB-IoT产品，4 LoRa产品，3 LoRa网关产品，5 普通网关产品   默认值是0
	ProductType *uint64 `json:"ProductType,omitempty" name:"ProductType"`

	// 数据格式，取值为json或者custom，默认值是json
	Format *string `json:"Format,omitempty" name:"Format"`

	// 产品所属平台，默认值是0
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// LoRa产品运营侧APPEUI，只有LoRa产品需要填写
	Appeui *string `json:"Appeui,omitempty" name:"Appeui"`

	// 产品绑定的物模型ID，-1表示不绑定
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 产品绑定的物模型名称
	ModelName *string `json:"ModelName,omitempty" name:"ModelName"`

	// 产品密钥，suite产品才会有
	ProductKey *string `json:"ProductKey,omitempty" name:"ProductKey"`

	// 动态注册类型 0-关闭, 1-预定义设备名 2-动态定义设备名
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// 动态注册产品秘钥
	ProductSecret *string `json:"ProductSecret,omitempty" name:"ProductSecret"`

	// RegisterType为2时，设备动态创建的限制数量
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`
}

type PublishAsDeviceRequest struct {
	*tchttp.BaseRequest

	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// LoRa 设备端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 消息内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`
}

func (r *PublishAsDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PublishAsDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PublishAsDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishAsDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PublishAsDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PublishMessageRequest struct {
	*tchttp.BaseRequest

	// 消息发往的主题。命名规则：${ProductId}/${DeviceName}/[a-zA-Z0-9:_-]{1,128}
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 消息内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 服务质量等级，取值为0或1
	Qos *uint64 `json:"Qos,omitempty" name:"Qos"`
}

func (r *PublishMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PublishMessageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PublishMessageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PublishMessageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PublishToDeviceRequest struct {
	*tchttp.BaseRequest

	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// LoRa 端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 消息内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`
}

func (r *PublishToDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PublishToDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PublishToDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishToDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PublishToDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceTopicRuleRequest struct {
	*tchttp.BaseRequest

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 替换的规则包体
	TopicRulePayload *TopicRulePayload `json:"TopicRulePayload,omitempty" name:"TopicRulePayload"`

	// 修改类型，0：其他，1：创建行为，2：更新行为，3：删除行为
	ModifyType *uint64 `json:"ModifyType,omitempty" name:"ModifyType"`

	// action增删改变更填对应topicRulePayload里面第几个action
	ActionIndex *uint64 `json:"ActionIndex,omitempty" name:"ActionIndex"`
}

func (r *ReplaceTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceTopicRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceTopicRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Task struct {

	// 批量更新影子任务的描述细节，当 taskType 取值为 “UpdateShadow” 时，此字段必填。描述见下 BatchUpdateShadow
	UpdateShadowTask *BatchUpdateShadow `json:"UpdateShadowTask,omitempty" name:"UpdateShadowTask"`

	// 批量下发消息任务的描述细节，当 taskType 取值为 “PublishMessage” 时，此字段必填。描述见下 BatchPublishMessage
	PublishMessageTask *BatchPublishMessage `json:"PublishMessageTask,omitempty" name:"PublishMessageTask"`
}

type TaskInfo struct {

	// 任务类型，目前取值为 “UpdateShadow” 或者 “PublishMessage”
	Type *string `json:"Type,omitempty" name:"Type"`

	// 任务 ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 状态。1表示等待处理，2表示调度处理中，3表示已完成，4表示失败，5表示已取消
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 任务创建时间，Unix 时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后任务更新时间，Unix 时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 返回的错误码
	RetCode *uint64 `json:"RetCode,omitempty" name:"RetCode"`

	// 返回的错误信息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type TopicRulePayload struct {

	// 规则的SQL语句，如： SELECT * FROM 'pid/dname/event'，然后对其进行base64编码，得：U0VMRUNUICogRlJPTSAncGlkL2RuYW1lL2V2ZW50Jw==
	Sql *string `json:"Sql,omitempty" name:"Sql"`

	// 行为的JSON字符串，大部分种类举例如下：
	// [{"republish":{"topic":"TEST/test"}},{"forward":{"api":"http://127.0.0.1:8080"}},{"ckafka":{"instance":{"id":"ckafka-test","name":""},"topic":{"id":"topic-test","name":"test"},"region":"gz"}},{"cmqqueue":{"queuename":"queue-test-TEST","region":"gz"}},{"mysql":{"instanceid":"cdb-test","region":"gz","username":"test","userpwd":"*****","dbname":"d_mqtt","tablename":"t_test","fieldpairs":[{"field":"test","value":"test"}],"devicetype":"CUSTOM"}}]
	Actions *string `json:"Actions,omitempty" name:"Actions"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否禁用规则
	RuleDisabled *bool `json:"RuleDisabled,omitempty" name:"RuleDisabled"`
}

type UpdateDeviceShadowRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 虚拟设备的状态，JSON字符串格式，由desired结构组成
	State *string `json:"State,omitempty" name:"State"`

	// 当前版本号，需要和后台的version保持一致，才能更新成功
	ShadowVersion *uint64 `json:"ShadowVersion,omitempty" name:"ShadowVersion"`
}

func (r *UpdateDeviceShadowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateDeviceShadowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateDeviceShadowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备影子数据，JSON字符串格式
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDeviceShadowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateDeviceShadowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateTopicPolicyRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 更新前Topic名
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 更新后Topic名
	NewTopicName *string `json:"NewTopicName,omitempty" name:"NewTopicName"`

	// Topic权限
	Privilege *uint64 `json:"Privilege,omitempty" name:"Privilege"`

	// 代理订阅信息
	BrokerSubscribe *BrokerSubscribe `json:"BrokerSubscribe,omitempty" name:"BrokerSubscribe"`
}

func (r *UpdateTopicPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateTopicPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateTopicPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
