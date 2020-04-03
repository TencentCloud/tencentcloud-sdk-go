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

package v20190423

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CallDeviceActionAsyncRequest struct {
	*tchttp.BaseRequest

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 动作Id
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitempty" name:"InputParams"`
}

func (r *CallDeviceActionAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CallDeviceActionAsyncRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CallDeviceActionAsyncResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 调用Id
		ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

		// 异步调用状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CallDeviceActionAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CallDeviceActionAsyncResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CallDeviceActionSyncRequest struct {
	*tchttp.BaseRequest

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 动作Id
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitempty" name:"InputParams"`
}

func (r *CallDeviceActionSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CallDeviceActionSyncRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CallDeviceActionSyncResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 调用Id
		ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

		// 输出参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		OutputParams *string `json:"OutputParams,omitempty" name:"OutputParams"`

		// 返回状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CallDeviceActionSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CallDeviceActionSyncResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ControlDeviceDataRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 属性数据, JSON格式字符串, 注意字段需要在物模型属性里定义
	Data *string `json:"Data,omitempty" name:"Data"`

	// 请求类型 , 不填该参数或者 desired 表示下发属性给设备,  reported 表示模拟设备上报属性
	Method *string `json:"Method,omitempty" name:"Method"`

	// 设备ID，该字段有值将代替 ProductId/DeviceName , 通常情况不需要填写
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 上报数据UNIX时间戳(毫秒), 仅对Method:reported有效
	DataTimestamp *int64 `json:"DataTimestamp,omitempty" name:"DataTimestamp"`
}

func (r *ControlDeviceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ControlDeviceDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ControlDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *string `json:"Data,omitempty" name:"Data"`

		// JSON字符串， 返回下发控制的结果信息, 
	// Sent = 1 表示设备已经在线并且订阅了控制下发的mqtt topic
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ControlDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ControlDeviceDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceRequest struct {
	*tchttp.BaseRequest

	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称。命名规则：[a-zA-Z0-9:_-]{1,48}。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// LoRaWAN 设备地址
	DevAddr *string `json:"DevAddr,omitempty" name:"DevAddr"`

	// LoRaWAN 应用密钥
	AppKey *string `json:"AppKey,omitempty" name:"AppKey"`

	// LoRaWAN 设备唯一标识
	DevEUI *string `json:"DevEUI,omitempty" name:"DevEUI"`

	// LoRaWAN 应用会话密钥
	AppSKey *string `json:"AppSKey,omitempty" name:"AppSKey"`

	// LoRaWAN 网络会话密钥
	NwkSKey *string `json:"NwkSKey,omitempty" name:"NwkSKey"`
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

		// 设备参数描述。
		Data *DeviceData `json:"Data,omitempty" name:"Data"`

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

type CreateProjectRequest struct {
	*tchttp.BaseRequest

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Project *ProjectEntry `json:"Project,omitempty" name:"Project"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStudioProductRequest struct {
	*tchttp.BaseRequest

	// 产品名称，名称不能和已经存在的产品名称重复。命名规则：[a-zA-Z0-9:_-]{1,32}
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品分组模板ID , ( 自定义模板填写1 , 控制台调用会使用预置的其他ID)
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 产品类型 填写 ( 0 普通产品 )
	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`

	// 加密类型 加密类型，1表示证书认证，2表示签名认证。
	EncryptionType *string `json:"EncryptionType,omitempty" name:"EncryptionType"`

	// 连接类型 可以填写 wifi cellular else
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 数据协议 (1 使用物模型 2 为自定义)
	DataProtocol *int64 `json:"DataProtocol,omitempty" name:"DataProtocol"`

	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitempty" name:"ProductDesc"`

	// 产品的项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStudioProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品描述
		Product *ProductEntry `json:"Product,omitempty" name:"Product"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStudioProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest

	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
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

type DeleteProjectRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteStudioProductRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DeleteStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteStudioProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteStudioProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceDataHistoryRequest struct {
	*tchttp.BaseRequest

	// 区间开始时间（Unix 时间戳，毫秒级）
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 区间结束时间（Unix 时间戳，毫秒级）
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 属性字段名称，对应数据模板中功能属性的标识符
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// 返回条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`
}

func (r *DescribeDeviceDataHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceDataHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceDataHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 属性字段名称，对应数据模板中功能属性的标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
		FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

		// 数据是否已全部返回，true 表示数据全部返回，false 表示还有数据待返回，可将 Context 作为入参，继续查询返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Listover *bool `json:"Listover,omitempty" name:"Listover"`

		// 检索上下文，当 ListOver 为false时，可以用此上下文，继续读取后续数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Context *string `json:"Context,omitempty" name:"Context"`

		// 历史数据结果数组，返回对应时间点及取值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Results []*DeviceDataHistoryItem `json:"Results,omitempty" name:"Results" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceDataHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceDataHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceDataRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备ID，该字段有值将代替 ProductId/DeviceName
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

func (r *DescribeDeviceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备数据
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备ID，该字段有值将代替 ProductId/DeviceName
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
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

		// 设备信息
		Device *DeviceInfo `json:"Device,omitempty" name:"Device"`

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

type DescribeModelDefinitionRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeModelDefinitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModelDefinitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModelDefinitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品数据模板
		Model *ProductModelDefinition `json:"Model,omitempty" name:"Model"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModelDefinitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModelDefinitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Project *ProjectEntryEx `json:"Project,omitempty" name:"Project"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStudioProductRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStudioProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品详情
		Product *ProductEntry `json:"Product,omitempty" name:"Product"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStudioProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeviceData struct {

	// 设备证书，用于 TLS 建立链接时校验客户端身份。采用非对称加密时返回该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceCert *string `json:"DeviceCert,omitempty" name:"DeviceCert"`

	// 设备名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备私钥，用于 TLS 建立链接时校验客户端身份，腾讯云后台不保存，请妥善保管。采用非对称加密时返回该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevicePrivateKey *string `json:"DevicePrivateKey,omitempty" name:"DevicePrivateKey"`

	// 对称加密密钥，base64编码。采用对称加密时返回该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`
}

type DeviceDataHistoryItem struct {

	// 时间点，毫秒时间戳
	Time *string `json:"Time,omitempty" name:"Time"`

	// 字段取值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeviceInfo struct {

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 0: 离线, 1: 在线, 2: 获取失败, 3 未激活
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 设备密钥，密钥加密的设备返回
	DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

	// 首次上线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstOnlineTime *int64 `json:"FirstOnlineTime,omitempty" name:"FirstOnlineTime"`

	// 最后一次上线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoginTime *int64 `json:"LoginTime,omitempty" name:"LoginTime"`

	// 设备创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 设备固件版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 设备证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceCert *string `json:"DeviceCert,omitempty" name:"DeviceCert"`

	// 日志级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogLevel *int64 `json:"LogLevel,omitempty" name:"LogLevel"`

	// LoRaWAN 设备地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevAddr *string `json:"DevAddr,omitempty" name:"DevAddr"`

	// LoRaWAN 应用密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppKey *string `json:"AppKey,omitempty" name:"AppKey"`

	// LoRaWAN 设备唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevEUI *string `json:"DevEUI,omitempty" name:"DevEUI"`

	// LoRaWAN 应用会话密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppSKey *string `json:"AppSKey,omitempty" name:"AppSKey"`

	// LoRaWAN 网络会话密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	NwkSKey *string `json:"NwkSKey,omitempty" name:"NwkSKey"`
}

type EventHistoryItem struct {

	// 事件的时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeStamp *int64 `json:"TimeStamp,omitempty" name:"TimeStamp"`

	// 事件的产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 事件的设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 事件的标识符ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 事件的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 事件的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`
}

type GetDeviceListRequest struct {
	*tchttp.BaseRequest

	// 需要查看设备列表的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDeviceListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的设备列表, 注意列表设备的 DevicePsk 为空, 要获取设备的 DevicePsk 请使用 DescribeDevice
	// 注意：此字段可能返回 null，表示取不到有效值。
		Devices []*DeviceInfo `json:"Devices,omitempty" name:"Devices" list`

		// 产品下的设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDeviceListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetProjectListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 个数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetProjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProjectListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetProjectListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Projects []*ProjectEntryEx `json:"Projects,omitempty" name:"Projects" list`

		// 列表项个数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProjectListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetStudioProductListRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品DevStatus
	DevStatus *string `json:"DevStatus,omitempty" name:"DevStatus"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetStudioProductListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetStudioProductListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetStudioProductListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品列表
		Products []*ProductEntry `json:"Products,omitempty" name:"Products" list`

		// 产品数量
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStudioProductListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetStudioProductListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListEventHistoryRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 搜索的事件类型：alert 表示告警，fault 表示故障，info 表示信息，为空则表示查询上述所有类型事件
	Type *string `json:"Type,omitempty" name:"Type"`

	// 起始时间（Unix 时间戳，秒级）, 为0 表示 当前时间 - 24h
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（Unix 时间戳，秒级）, 为0 表示当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 搜索上下文, 用作查询游标
	Context *string `json:"Context,omitempty" name:"Context"`

	// 单次获取的历史数据项目的最大数量
	Size *int64 `json:"Size,omitempty" name:"Size"`
}

func (r *ListEventHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListEventHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListEventHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 搜索上下文, 用作查询游标
	// 注意：此字段可能返回 null，表示取不到有效值。
		Context *string `json:"Context,omitempty" name:"Context"`

		// 搜索结果数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 搜索结果是否已经结束
	// 注意：此字段可能返回 null，表示取不到有效值。
		Listover *bool `json:"Listover,omitempty" name:"Listover"`

		// 搜集结果集
	// 注意：此字段可能返回 null，表示取不到有效值。
		EventHistory []*EventHistoryItem `json:"EventHistory,omitempty" name:"EventHistory" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEventHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListEventHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModelDefinitionRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 数据模板定义
	ModelSchema *string `json:"ModelSchema,omitempty" name:"ModelSchema"`
}

func (r *ModifyModelDefinitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModelDefinitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModelDefinitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModelDefinitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModelDefinitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目详情
		Project *ProjectEntry `json:"Project,omitempty" name:"Project"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyStudioProductRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitempty" name:"ProductDesc"`

	// 模型ID
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`

	// 是否打开二进制转Json功能, 取值为字符串 true/false
	EnableProductScript *string `json:"EnableProductScript,omitempty" name:"EnableProductScript"`
}

func (r *ModifyStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyStudioProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品描述
		Product *ProductEntry `json:"Product,omitempty" name:"Product"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyStudioProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProductEntry struct {

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品分组模板ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 加密类型
	EncryptionType *string `json:"EncryptionType,omitempty" name:"EncryptionType"`

	// 连接类型
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 数据协议
	DataProtocol *int64 `json:"DataProtocol,omitempty" name:"DataProtocol"`

	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitempty" name:"ProductDesc"`

	// 状态
	DevStatus *string `json:"DevStatus,omitempty" name:"DevStatus"`

	// 创建时间
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 产品类型
	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品ModuleId
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`

	// 是否使用脚本进行二进制转json功能 可以取值 true / false
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableProductScript *string `json:"EnableProductScript,omitempty" name:"EnableProductScript"`
}

type ProductModelDefinition struct {

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 模型定义
	ModelDefine *string `json:"ModelDefine,omitempty" name:"ModelDefine"`

	// 更新时间，秒级时间戳
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间，秒级时间戳
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 产品所属分类的模型快照（产品创建时刻的）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryModel *string `json:"CategoryModel,omitempty" name:"CategoryModel"`
}

type ProjectEntry struct {

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`

	// 创建时间，unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间，unix时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ProjectEntryEx struct {

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`

	// 项目创建时间，unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目更新时间，unix时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 产品数量
	ProductCount *uint64 `json:"ProductCount,omitempty" name:"ProductCount"`

	// NativeApp数量
	NativeAppCount *uint64 `json:"NativeAppCount,omitempty" name:"NativeAppCount"`

	// WebApp数量
	WebAppCount *uint64 `json:"WebAppCount,omitempty" name:"WebAppCount"`
}

type ReleaseStudioProductRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品DevStatus
	DevStatus *string `json:"DevStatus,omitempty" name:"DevStatus"`
}

func (r *ReleaseStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseStudioProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseStudioProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStudioProductRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 列表Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 列表Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 产品Status
	DevStatus *string `json:"DevStatus,omitempty" name:"DevStatus"`
}

func (r *SearchStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchStudioProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品列表
		Products []*ProductEntry `json:"Products,omitempty" name:"Products" list`

		// 产品数量
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchStudioProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
