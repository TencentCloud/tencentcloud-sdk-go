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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AppDeviceInfo struct {
	// 产品ID/设备名
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备别名
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// icon地址
	IconUrl *string `json:"IconUrl,omitempty" name:"IconUrl"`

	// 家庭ID
	FamilyId *string `json:"FamilyId,omitempty" name:"FamilyId"`

	// 房间ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 设备类型
	DeviceType *int64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type BatchProductionInfo struct {
	// 量产ID
	BatchProductionId *string `json:"BatchProductionId,omitempty" name:"BatchProductionId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 烧录方式
	BurnMethod *int64 `json:"BurnMethod,omitempty" name:"BurnMethod"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type BindDeviceInfo struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

// Predefined struct for user
type BindDevicesRequestParams struct {
	// 网关设备的产品ID。
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名。
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 被绑定设备的产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 被绑定的多个设备名。
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`
}

type BindDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 网关设备的产品ID。
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名。
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 被绑定设备的产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 被绑定的多个设备名。
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`
}

func (r *BindDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "GatewayDeviceName")
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDevicesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindDevicesResponse struct {
	*tchttp.BaseResponse
	Response *BindDevicesResponseParams `json:"Response"`
}

func (r *BindDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindProductInfo struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品所属项目ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 物模型类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataProtocol *int64 `json:"DataProtocol,omitempty" name:"DataProtocol"`

	// 产品分组模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 产品类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`

	// 连接类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevStatus *string `json:"DevStatus,omitempty" name:"DevStatus"`

	// 产品拥有者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductOwnerName *string `json:"ProductOwnerName,omitempty" name:"ProductOwnerName"`
}

// Predefined struct for user
type BindProductsRequestParams struct {
	// 网关产品ID。
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 待绑定的子产品ID数组。
	ProductIds []*string `json:"ProductIds,omitempty" name:"ProductIds"`
}

type BindProductsRequest struct {
	*tchttp.BaseRequest
	
	// 网关产品ID。
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 待绑定的子产品ID数组。
	ProductIds []*string `json:"ProductIds,omitempty" name:"ProductIds"`
}

func (r *BindProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "ProductIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindProductsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindProductsResponse struct {
	*tchttp.BaseResponse
	Response *BindProductsResponseParams `json:"Response"`
}

func (r *BindProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallDeviceActionAsyncRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitempty" name:"InputParams"`
}

type CallDeviceActionAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitempty" name:"InputParams"`
}

func (r *CallDeviceActionAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallDeviceActionAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ActionId")
	delete(f, "InputParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CallDeviceActionAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallDeviceActionAsyncResponseParams struct {
	// 调用Id
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 异步调用状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CallDeviceActionAsyncResponse struct {
	*tchttp.BaseResponse
	Response *CallDeviceActionAsyncResponseParams `json:"Response"`
}

func (r *CallDeviceActionAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallDeviceActionAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallDeviceActionSyncRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitempty" name:"InputParams"`
}

type CallDeviceActionSyncRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitempty" name:"InputParams"`
}

func (r *CallDeviceActionSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallDeviceActionSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ActionId")
	delete(f, "InputParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CallDeviceActionSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallDeviceActionSyncResponseParams struct {
	// 调用Id
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 输出参数，取值设备端上报$thing/up/action method为action_reply 的 response字段，物模型协议参考https://cloud.tencent.com/document/product/1081/34916#.E8.AE.BE.E5.A4.87.E8.A1.8C.E4.B8.BA.E8.B0.83.E7.94.A8
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputParams *string `json:"OutputParams,omitempty" name:"OutputParams"`

	// 返回状态，取值设备端上报$thing/up/action	method为action_reply 的 status字段，如果不包含status字段，则取默认值，空字符串，物模型协议参考https://cloud.tencent.com/document/product/1081/34916#.E8.AE.BE.E5.A4.87.E8.A1.8C.E4.B8.BA.E8.B0.83.E7.94.A8
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CallDeviceActionSyncResponse struct {
	*tchttp.BaseResponse
	Response *CallDeviceActionSyncResponseParams `json:"Response"`
}

func (r *CallDeviceActionSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallDeviceActionSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDeviceDataRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Data")
	delete(f, "Method")
	delete(f, "DeviceId")
	delete(f, "DataTimestamp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlDeviceDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDeviceDataResponseParams struct {
	// 返回信息
	Data *string `json:"Data,omitempty" name:"Data"`

	// JSON字符串， 返回下发控制的结果信息, 
	// Sent = 1 表示设备已经在线并且订阅了控制下发的mqtt topic.
	// pushResult 是表示发送结果，其中 0 表示成功， 23101 表示设备未在线或没有订阅相关的 MQTT Topic。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ControlDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *ControlDeviceDataResponseParams `json:"Response"`
}

func (r *ControlDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchProductionRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 烧录方式，0为直接烧录，1为动态注册。
	BurnMethod *int64 `json:"BurnMethod,omitempty" name:"BurnMethod"`

	// 生成方式，0为系统生成，1为文件上传。
	GenerationMethod *int64 `json:"GenerationMethod,omitempty" name:"GenerationMethod"`

	// 文件上传URL，用于文件上传时填写。
	UploadUrl *string `json:"UploadUrl,omitempty" name:"UploadUrl"`

	// 量产数量，用于系统生成时填写。
	BatchCnt *int64 `json:"BatchCnt,omitempty" name:"BatchCnt"`

	// 是否生成二维码,0为不生成，1为生成。
	GenerationQRCode *int64 `json:"GenerationQRCode,omitempty" name:"GenerationQRCode"`
}

type CreateBatchProductionRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 烧录方式，0为直接烧录，1为动态注册。
	BurnMethod *int64 `json:"BurnMethod,omitempty" name:"BurnMethod"`

	// 生成方式，0为系统生成，1为文件上传。
	GenerationMethod *int64 `json:"GenerationMethod,omitempty" name:"GenerationMethod"`

	// 文件上传URL，用于文件上传时填写。
	UploadUrl *string `json:"UploadUrl,omitempty" name:"UploadUrl"`

	// 量产数量，用于系统生成时填写。
	BatchCnt *int64 `json:"BatchCnt,omitempty" name:"BatchCnt"`

	// 是否生成二维码,0为不生成，1为生成。
	GenerationQRCode *int64 `json:"GenerationQRCode,omitempty" name:"GenerationQRCode"`
}

func (r *CreateBatchProductionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchProductionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ProductId")
	delete(f, "BurnMethod")
	delete(f, "GenerationMethod")
	delete(f, "UploadUrl")
	delete(f, "BatchCnt")
	delete(f, "GenerationQRCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchProductionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchProductionResponseParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 量产id
	BatchProductionId *string `json:"BatchProductionId,omitempty" name:"BatchProductionId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBatchProductionResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchProductionResponseParams `json:"Response"`
}

func (r *CreateBatchProductionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchProductionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceRequestParams struct {
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

	// 手动指定设备的PSK密钥
	DefinedPsk *string `json:"DefinedPsk,omitempty" name:"DefinedPsk"`
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

	// 手动指定设备的PSK密钥
	DefinedPsk *string `json:"DefinedPsk,omitempty" name:"DefinedPsk"`
}

func (r *CreateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "DevAddr")
	delete(f, "AppKey")
	delete(f, "DevEUI")
	delete(f, "AppSKey")
	delete(f, "NwkSKey")
	delete(f, "DefinedPsk")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceResponseParams struct {
	// 设备参数描述。
	Data *DeviceData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeviceResponseParams `json:"Response"`
}

func (r *CreateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFenceBindRequestParams struct {
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`

	// 围栏绑定的产品列表
	Items []*FenceBindProductItem `json:"Items,omitempty" name:"Items"`
}

type CreateFenceBindRequest struct {
	*tchttp.BaseRequest
	
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`

	// 围栏绑定的产品列表
	Items []*FenceBindProductItem `json:"Items,omitempty" name:"Items"`
}

func (r *CreateFenceBindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFenceBindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FenceId")
	delete(f, "Items")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFenceBindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFenceBindResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFenceBindResponse struct {
	*tchttp.BaseResponse
	Response *CreateFenceBindResponseParams `json:"Response"`
}

func (r *CreateFenceBindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFenceBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoRaFrequencyRequestParams struct {
	// 频点配置名称
	FreqName *string `json:"FreqName,omitempty" name:"FreqName"`

	// 数据上行信道
	ChannelsDataUp []*uint64 `json:"ChannelsDataUp,omitempty" name:"ChannelsDataUp"`

	// 数据下行RX1信道
	ChannelsDataRX1 []*uint64 `json:"ChannelsDataRX1,omitempty" name:"ChannelsDataRX1"`

	// 数据下行RX2信道
	ChannelsDataRX2 []*uint64 `json:"ChannelsDataRX2,omitempty" name:"ChannelsDataRX2"`

	// 入网上行信道
	ChannelsJoinUp []*uint64 `json:"ChannelsJoinUp,omitempty" name:"ChannelsJoinUp"`

	// 入网下行RX1信道
	ChannelsJoinRX1 []*uint64 `json:"ChannelsJoinRX1,omitempty" name:"ChannelsJoinRX1"`

	// 入网下行RX2信道
	ChannelsJoinRX2 []*uint64 `json:"ChannelsJoinRX2,omitempty" name:"ChannelsJoinRX2"`

	// 频点配置描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateLoRaFrequencyRequest struct {
	*tchttp.BaseRequest
	
	// 频点配置名称
	FreqName *string `json:"FreqName,omitempty" name:"FreqName"`

	// 数据上行信道
	ChannelsDataUp []*uint64 `json:"ChannelsDataUp,omitempty" name:"ChannelsDataUp"`

	// 数据下行RX1信道
	ChannelsDataRX1 []*uint64 `json:"ChannelsDataRX1,omitempty" name:"ChannelsDataRX1"`

	// 数据下行RX2信道
	ChannelsDataRX2 []*uint64 `json:"ChannelsDataRX2,omitempty" name:"ChannelsDataRX2"`

	// 入网上行信道
	ChannelsJoinUp []*uint64 `json:"ChannelsJoinUp,omitempty" name:"ChannelsJoinUp"`

	// 入网下行RX1信道
	ChannelsJoinRX1 []*uint64 `json:"ChannelsJoinRX1,omitempty" name:"ChannelsJoinRX1"`

	// 入网下行RX2信道
	ChannelsJoinRX2 []*uint64 `json:"ChannelsJoinRX2,omitempty" name:"ChannelsJoinRX2"`

	// 频点配置描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateLoRaFrequencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoRaFrequencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FreqName")
	delete(f, "ChannelsDataUp")
	delete(f, "ChannelsDataRX1")
	delete(f, "ChannelsDataRX2")
	delete(f, "ChannelsJoinUp")
	delete(f, "ChannelsJoinRX1")
	delete(f, "ChannelsJoinRX2")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoRaFrequencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoRaFrequencyResponseParams struct {
	// LoRa频点信息
	Data *LoRaFrequencyEntry `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLoRaFrequencyResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoRaFrequencyResponseParams `json:"Response"`
}

func (r *CreateLoRaFrequencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoRaFrequencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoRaGatewayRequestParams struct {
	// LoRa 网关Id
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// 网关名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 详情描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 位置坐标
	Location *LoRaGatewayLocation `json:"Location,omitempty" name:"Location"`

	// 位置信息
	Position *string `json:"Position,omitempty" name:"Position"`

	// 位置详情
	PositionDetails *string `json:"PositionDetails,omitempty" name:"PositionDetails"`

	// 是否公开
	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

	// 频点ID
	FrequencyId *string `json:"FrequencyId,omitempty" name:"FrequencyId"`
}

type CreateLoRaGatewayRequest struct {
	*tchttp.BaseRequest
	
	// LoRa 网关Id
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// 网关名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 详情描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 位置坐标
	Location *LoRaGatewayLocation `json:"Location,omitempty" name:"Location"`

	// 位置信息
	Position *string `json:"Position,omitempty" name:"Position"`

	// 位置详情
	PositionDetails *string `json:"PositionDetails,omitempty" name:"PositionDetails"`

	// 是否公开
	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

	// 频点ID
	FrequencyId *string `json:"FrequencyId,omitempty" name:"FrequencyId"`
}

func (r *CreateLoRaGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoRaGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Location")
	delete(f, "Position")
	delete(f, "PositionDetails")
	delete(f, "IsPublic")
	delete(f, "FrequencyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoRaGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoRaGatewayResponseParams struct {
	// LoRa 网关信息
	Gateway *LoRaGatewayItem `json:"Gateway,omitempty" name:"Gateway"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLoRaGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoRaGatewayResponseParams `json:"Response"`
}

func (r *CreateLoRaGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoRaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePositionFenceRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 围栏名称
	FenceName *string `json:"FenceName,omitempty" name:"FenceName"`

	// 围栏区域信息，采用 GeoJSON 格式
	FenceArea *string `json:"FenceArea,omitempty" name:"FenceArea"`

	// 围栏描述
	FenceDesc *string `json:"FenceDesc,omitempty" name:"FenceDesc"`
}

type CreatePositionFenceRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 围栏名称
	FenceName *string `json:"FenceName,omitempty" name:"FenceName"`

	// 围栏区域信息，采用 GeoJSON 格式
	FenceArea *string `json:"FenceArea,omitempty" name:"FenceArea"`

	// 围栏描述
	FenceDesc *string `json:"FenceDesc,omitempty" name:"FenceDesc"`
}

func (r *CreatePositionFenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePositionFenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "FenceName")
	delete(f, "FenceArea")
	delete(f, "FenceDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePositionFenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePositionFenceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePositionFenceResponse struct {
	*tchttp.BaseResponse
	Response *CreatePositionFenceResponseParams `json:"Response"`
}

func (r *CreatePositionFenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePositionFenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePositionSpaceRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 空间名称
	SpaceName *string `json:"SpaceName,omitempty" name:"SpaceName"`

	// 授权类型，0：只读 1：读写
	AuthorizeType *int64 `json:"AuthorizeType,omitempty" name:"AuthorizeType"`

	// 产品列表
	ProductIdList []*string `json:"ProductIdList,omitempty" name:"ProductIdList"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 缩略图
	Icon *string `json:"Icon,omitempty" name:"Icon"`
}

type CreatePositionSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 空间名称
	SpaceName *string `json:"SpaceName,omitempty" name:"SpaceName"`

	// 授权类型，0：只读 1：读写
	AuthorizeType *int64 `json:"AuthorizeType,omitempty" name:"AuthorizeType"`

	// 产品列表
	ProductIdList []*string `json:"ProductIdList,omitempty" name:"ProductIdList"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 缩略图
	Icon *string `json:"Icon,omitempty" name:"Icon"`
}

func (r *CreatePositionSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePositionSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SpaceName")
	delete(f, "AuthorizeType")
	delete(f, "ProductIdList")
	delete(f, "Description")
	delete(f, "Icon")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePositionSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePositionSpaceResponseParams struct {
	// 空间Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePositionSpaceResponse struct {
	*tchttp.BaseResponse
	Response *CreatePositionSpaceResponseParams `json:"Response"`
}

func (r *CreatePositionSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePositionSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectRequestParams struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`

	// 实例ID，不带实例ID，默认为公共实例
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`

	// 实例ID，不带实例ID，默认为公共实例
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectName")
	delete(f, "ProjectDesc")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// 返回信息
	Project *ProjectEntry `json:"Project,omitempty" name:"Project"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateProjectResponseParams `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStudioProductRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStudioProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductName")
	delete(f, "CategoryId")
	delete(f, "ProductType")
	delete(f, "EncryptionType")
	delete(f, "NetType")
	delete(f, "DataProtocol")
	delete(f, "ProductDesc")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStudioProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStudioProductResponseParams struct {
	// 产品描述
	Product *ProductEntry `json:"Product,omitempty" name:"Product"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *CreateStudioProductResponseParams `json:"Response"`
}

func (r *CreateStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStudioProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicPolicyRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Topic权限，1发布，2订阅，3订阅和发布
	Privilege *uint64 `json:"Privilege,omitempty" name:"Privilege"`
}

type CreateTopicPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Topic权限，1发布，2订阅，3订阅和发布
	Privilege *uint64 `json:"Privilege,omitempty" name:"Privilege"`
}

func (r *CreateTopicPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "TopicName")
	delete(f, "Privilege")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicPolicyResponseParams `json:"Response"`
}

func (r *CreateTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicRuleRequestParams struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则内容
	TopicRulePayload *TopicRulePayload `json:"TopicRulePayload,omitempty" name:"TopicRulePayload"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	delete(f, "TopicRulePayload")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicRuleResponseParams `json:"Response"`
}

func (r *CreateTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceRequestParams struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 是否删除绑定设备
	ForceDelete *bool `json:"ForceDelete,omitempty" name:"ForceDelete"`
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 是否删除绑定设备
	ForceDelete *bool `json:"ForceDelete,omitempty" name:"ForceDelete"`
}

func (r *DeleteDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ForceDelete")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceResponseParams struct {
	// 删除的结果代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultCode *string `json:"ResultCode,omitempty" name:"ResultCode"`

	// 删除的结果信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultMessage *string `json:"ResultMessage,omitempty" name:"ResultMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeviceResponseParams `json:"Response"`
}

func (r *DeleteDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDevicesRequestParams struct {
	// 多个设备标识
	DevicesItems []*DevicesItem `json:"DevicesItems,omitempty" name:"DevicesItems"`
}

type DeleteDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 多个设备标识
	DevicesItems []*DevicesItem `json:"DevicesItems,omitempty" name:"DevicesItems"`
}

func (r *DeleteDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DevicesItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDevicesResponseParams struct {
	// 删除的结果代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultCode *string `json:"ResultCode,omitempty" name:"ResultCode"`

	// 删除的结果信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultMessage *string `json:"ResultMessage,omitempty" name:"ResultMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDevicesResponseParams `json:"Response"`
}

func (r *DeleteDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFenceBindRequestParams struct {
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`

	// 围栏绑定的产品信息
	Items []*FenceBindProductItem `json:"Items,omitempty" name:"Items"`
}

type DeleteFenceBindRequest struct {
	*tchttp.BaseRequest
	
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`

	// 围栏绑定的产品信息
	Items []*FenceBindProductItem `json:"Items,omitempty" name:"Items"`
}

func (r *DeleteFenceBindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFenceBindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FenceId")
	delete(f, "Items")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFenceBindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFenceBindResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteFenceBindResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFenceBindResponseParams `json:"Response"`
}

func (r *DeleteFenceBindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFenceBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoRaFrequencyRequestParams struct {
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitempty" name:"FreqId"`
}

type DeleteLoRaFrequencyRequest struct {
	*tchttp.BaseRequest
	
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitempty" name:"FreqId"`
}

func (r *DeleteLoRaFrequencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoRaFrequencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FreqId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoRaFrequencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoRaFrequencyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLoRaFrequencyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoRaFrequencyResponseParams `json:"Response"`
}

func (r *DeleteLoRaFrequencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoRaFrequencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoRaGatewayRequestParams struct {
	// LoRa 网关 Id
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

type DeleteLoRaGatewayRequest struct {
	*tchttp.BaseRequest
	
	// LoRa 网关 Id
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`
}

func (r *DeleteLoRaGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoRaGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoRaGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoRaGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLoRaGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoRaGatewayResponseParams `json:"Response"`
}

func (r *DeleteLoRaGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoRaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePositionFenceRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`
}

type DeletePositionFenceRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`
}

func (r *DeletePositionFenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePositionFenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "FenceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePositionFenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePositionFenceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePositionFenceResponse struct {
	*tchttp.BaseResponse
	Response *DeletePositionFenceResponseParams `json:"Response"`
}

func (r *DeletePositionFenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePositionFenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePositionSpaceRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`
}

type DeletePositionSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`
}

func (r *DeletePositionSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePositionSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePositionSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePositionSpaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePositionSpaceResponse struct {
	*tchttp.BaseResponse
	Response *DeletePositionSpaceResponseParams `json:"Response"`
}

func (r *DeletePositionSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePositionSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteStudioProductRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStudioProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStudioProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStudioProductResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStudioProductResponseParams `json:"Response"`
}

func (r *DeleteStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStudioProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicPolicyRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type DeleteTopicPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DeleteTopicPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicPolicyResponseParams `json:"Response"`
}

func (r *DeleteTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicRuleRequestParams struct {
	// 规则名
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicRuleResponseParams `json:"Response"`
}

func (r *DeleteTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchProductionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 量产ID
	BatchProductionId *string `json:"BatchProductionId,omitempty" name:"BatchProductionId"`
}

type DescribeBatchProductionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 量产ID
	BatchProductionId *string `json:"BatchProductionId,omitempty" name:"BatchProductionId"`
}

func (r *DescribeBatchProductionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchProductionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "BatchProductionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchProductionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchProductionResponseParams struct {
	// 量产数量。
	BatchCnt *int64 `json:"BatchCnt,omitempty" name:"BatchCnt"`

	// 烧录方式。
	BurnMethod *int64 `json:"BurnMethod,omitempty" name:"BurnMethod"`

	// 创建时间。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 下载URL。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 生成方式。
	GenerationMethod *int64 `json:"GenerationMethod,omitempty" name:"GenerationMethod"`

	// 上传URL。
	UploadUrl *string `json:"UploadUrl,omitempty" name:"UploadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBatchProductionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchProductionResponseParams `json:"Response"`
}

func (r *DescribeBatchProductionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchProductionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindedProductsRequestParams struct {
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 是否跨账号绑定产品
	ProductSource *int64 `json:"ProductSource,omitempty" name:"ProductSource"`
}

type DescribeBindedProductsRequest struct {
	*tchttp.BaseRequest
	
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 是否跨账号绑定产品
	ProductSource *int64 `json:"ProductSource,omitempty" name:"ProductSource"`
}

func (r *DescribeBindedProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindedProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProductSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindedProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindedProductsResponseParams struct {
	// 当前分页的子产品数组
	Products []*BindProductInfo `json:"Products,omitempty" name:"Products"`

	// 绑定的子产品总数量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBindedProductsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindedProductsResponseParams `json:"Response"`
}

func (r *DescribeBindedProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindedProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceBindGatewayRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DescribeDeviceBindGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DescribeDeviceBindGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceBindGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceBindGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceBindGatewayResponseParams struct {
	// 网关产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 网关产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`

	// 设备对应产品所属的主账号名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayProductOwnerName *string `json:"GatewayProductOwnerName,omitempty" name:"GatewayProductOwnerName"`

	// 设备对应产品所属的主账号 UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayProductOwnerUin *string `json:"GatewayProductOwnerUin,omitempty" name:"GatewayProductOwnerUin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceBindGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceBindGatewayResponseParams `json:"Response"`
}

func (r *DescribeDeviceBindGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceBindGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceDataHistoryRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceDataHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinTime")
	delete(f, "MaxTime")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "FieldName")
	delete(f, "Limit")
	delete(f, "Context")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceDataHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceDataHistoryResponseParams struct {
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
	Results []*DeviceDataHistoryItem `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceDataHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceDataHistoryResponseParams `json:"Response"`
}

func (r *DescribeDeviceDataHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceDataHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceDataRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备ID，该字段有值将代替 ProductId/DeviceName
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceDataResponseParams struct {
	// 设备数据
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceDataResponseParams `json:"Response"`
}

func (r *DescribeDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceLocationSolveRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 定位解析类型，wifi或GNSSNavigation
	LocationType *string `json:"LocationType,omitempty" name:"LocationType"`

	// LoRaEdge卫星导航电文
	GNSSNavigation *string `json:"GNSSNavigation,omitempty" name:"GNSSNavigation"`

	// wifi信息
	WiFiInfo []*WifiInfo `json:"WiFiInfo,omitempty" name:"WiFiInfo"`
}

type DescribeDeviceLocationSolveRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 定位解析类型，wifi或GNSSNavigation
	LocationType *string `json:"LocationType,omitempty" name:"LocationType"`

	// LoRaEdge卫星导航电文
	GNSSNavigation *string `json:"GNSSNavigation,omitempty" name:"GNSSNavigation"`

	// wifi信息
	WiFiInfo []*WifiInfo `json:"WiFiInfo,omitempty" name:"WiFiInfo"`
}

func (r *DescribeDeviceLocationSolveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceLocationSolveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "LocationType")
	delete(f, "GNSSNavigation")
	delete(f, "WiFiInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceLocationSolveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceLocationSolveResponseParams struct {
	// 经度
	Longitude *float64 `json:"Longitude,omitempty" name:"Longitude"`

	// 纬度
	Latitude *float64 `json:"Latitude,omitempty" name:"Latitude"`

	// 类型
	LocationType *string `json:"LocationType,omitempty" name:"LocationType"`

	// 误差精度预估，单位为米
	// 注意：此字段可能返回 null，表示取不到有效值。
	Accuracy *float64 `json:"Accuracy,omitempty" name:"Accuracy"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceLocationSolveResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceLocationSolveResponseParams `json:"Response"`
}

func (r *DescribeDeviceLocationSolveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceLocationSolveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePositionListRequestParams struct {
	// 产品标识列表
	ProductIdList []*string `json:"ProductIdList,omitempty" name:"ProductIdList"`

	// 坐标类型
	CoordinateType *int64 `json:"CoordinateType,omitempty" name:"CoordinateType"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDevicePositionListRequest struct {
	*tchttp.BaseRequest
	
	// 产品标识列表
	ProductIdList []*string `json:"ProductIdList,omitempty" name:"ProductIdList"`

	// 坐标类型
	CoordinateType *int64 `json:"CoordinateType,omitempty" name:"CoordinateType"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDevicePositionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePositionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductIdList")
	delete(f, "CoordinateType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicePositionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePositionListResponseParams struct {
	// 产品设备位置信息列表
	Positions []*ProductDevicesPositionItem `json:"Positions,omitempty" name:"Positions"`

	// 产品设备位置信息的数目
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDevicePositionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicePositionListResponseParams `json:"Response"`
}

func (r *DescribeDevicePositionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePositionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备ID，该字段有值将代替 ProductId/DeviceName
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceResponseParams struct {
	// 设备信息
	Device *DeviceInfo `json:"Device,omitempty" name:"Device"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceResponseParams `json:"Response"`
}

func (r *DescribeDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFenceBindListRequestParams struct {
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeFenceBindListRequest struct {
	*tchttp.BaseRequest
	
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFenceBindListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFenceBindListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FenceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFenceBindListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFenceBindListResponseParams struct {
	// 围栏绑定的产品设备列表
	List []*FenceBindProductItem `json:"List,omitempty" name:"List"`

	// 围栏绑定的设备总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFenceBindListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFenceBindListResponseParams `json:"Response"`
}

func (r *DescribeFenceBindListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFenceBindListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFenceEventListRequestParams struct {
	// 围栏告警信息的查询起始时间，Unix时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 围栏告警信息的查询结束时间，Unix时间，单位为毫秒
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 告警对应的产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 告警对应的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DescribeFenceEventListRequest struct {
	*tchttp.BaseRequest
	
	// 围栏告警信息的查询起始时间，Unix时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 围栏告警信息的查询结束时间，Unix时间，单位为毫秒
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 告警对应的产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 告警对应的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DescribeFenceEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFenceEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "FenceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFenceEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFenceEventListResponseParams struct {
	// 围栏告警事件列表
	List []*FenceEventItem `json:"List,omitempty" name:"List"`

	// 围栏告警事件总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFenceEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFenceEventListResponseParams `json:"Response"`
}

func (r *DescribeFenceEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFenceEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareTaskRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeFirmwareTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeFirmwareTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirmwareTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareTaskResponseParams struct {
	// 固件任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 固件任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 固件任务创建时间，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 固件任务升级类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 固件任务升级模式。originalVersion（按版本号升级）、filename（提交文件升级）、devicenames（按设备名称升级）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradeMode *string `json:"UpgradeMode,omitempty" name:"UpgradeMode"`

	// 产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 原始固件版本号，在UpgradeMode是originalVersion升级模式下会返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalVersion *string `json:"OriginalVersion,omitempty" name:"OriginalVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFirmwareTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirmwareTaskResponseParams `json:"Response"`
}

func (r *DescribeFirmwareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayBindDevicesRequestParams struct {
	// 网关设备的产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 子产品的ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 分页的偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeGatewayBindDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 网关设备的产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 子产品的ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 分页的偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGatewayBindDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayBindDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "GatewayDeviceName")
	delete(f, "ProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayBindDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayBindDevicesResponseParams struct {
	// 子设备信息。
	Devices []*BindDeviceInfo `json:"Devices,omitempty" name:"Devices"`

	// 子设备总数。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 子设备所属的产品名。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGatewayBindDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayBindDevicesResponseParams `json:"Response"`
}

func (r *DescribeGatewayBindDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayBindDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewaySubDeviceListRequestParams struct {
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备名称
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeGatewaySubDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备名称
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGatewaySubDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewaySubDeviceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "GatewayDeviceName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewaySubDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewaySubDeviceListResponseParams struct {
	// 设备的总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 设备列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceList []*FamilySubDevice `json:"DeviceList,omitempty" name:"DeviceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGatewaySubDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewaySubDeviceListResponseParams `json:"Response"`
}

func (r *DescribeGatewaySubDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewaySubDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewaySubProductsRequestParams struct {
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 分页的偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 是否跨账号产品
	ProductSource *int64 `json:"ProductSource,omitempty" name:"ProductSource"`
}

type DescribeGatewaySubProductsRequest struct {
	*tchttp.BaseRequest
	
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 分页的偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 是否跨账号产品
	ProductSource *int64 `json:"ProductSource,omitempty" name:"ProductSource"`
}

func (r *DescribeGatewaySubProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewaySubProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProjectId")
	delete(f, "ProductSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewaySubProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewaySubProductsResponseParams struct {
	// 当前分页的可绑定或解绑的产品信息。
	Products []*BindProductInfo `json:"Products,omitempty" name:"Products"`

	// 可绑定或解绑的产品总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGatewaySubProductsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewaySubProductsResponseParams `json:"Response"`
}

func (r *DescribeGatewaySubProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewaySubProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoRaFrequencyRequestParams struct {
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitempty" name:"FreqId"`
}

type DescribeLoRaFrequencyRequest struct {
	*tchttp.BaseRequest
	
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitempty" name:"FreqId"`
}

func (r *DescribeLoRaFrequencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoRaFrequencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FreqId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoRaFrequencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoRaFrequencyResponseParams struct {
	// 返回详情项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *LoRaFrequencyEntry `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoRaFrequencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoRaFrequencyResponseParams `json:"Response"`
}

func (r *DescribeLoRaFrequencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoRaFrequencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelDefinitionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelDefinitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelDefinitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelDefinitionResponseParams struct {
	// 产品数据模板
	Model *ProductModelDefinition `json:"Model,omitempty" name:"Model"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeModelDefinitionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelDefinitionResponseParams `json:"Response"`
}

func (r *DescribeModelDefinitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelDefinitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePositionFenceListRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePositionFenceListRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePositionFenceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePositionFenceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePositionFenceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePositionFenceListResponseParams struct {
	// 围栏列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*PositionFenceInfo `json:"List,omitempty" name:"List"`

	// 围栏数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePositionFenceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePositionFenceListResponseParams `json:"Response"`
}

func (r *DescribePositionFenceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePositionFenceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectResponseParams struct {
	// 返回信息
	Project *ProjectEntryEx `json:"Project,omitempty" name:"Project"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectResponseParams `json:"Response"`
}

func (r *DescribeProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceFenceEventListRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 围栏告警信息的查询起始时间，Unix时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 围栏告警信息的查询结束时间，Unix时间，单位为毫秒
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSpaceFenceEventListRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 围栏告警信息的查询起始时间，Unix时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 围栏告警信息的查询结束时间，Unix时间，单位为毫秒
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSpaceFenceEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceFenceEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceFenceEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceFenceEventListResponseParams struct {
	// 围栏告警事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*FenceEventItem `json:"List,omitempty" name:"List"`

	// 围栏告警事件总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSpaceFenceEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceFenceEventListResponseParams `json:"Response"`
}

func (r *DescribeSpaceFenceEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceFenceEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStudioProductRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStudioProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStudioProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStudioProductResponseParams struct {
	// 产品详情
	Product *ProductEntry `json:"Product,omitempty" name:"Product"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStudioProductResponseParams `json:"Response"`
}

func (r *DescribeStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStudioProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicPolicyRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Topic名字
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type DescribeTopicPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Topic名字
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DescribeTopicPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicPolicyResponseParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Topic权限
	Privilege *uint64 `json:"Privilege,omitempty" name:"Privilege"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicPolicyResponseParams `json:"Response"`
}

func (r *DescribeTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicRuleRequestParams struct {
	// 规则名称。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type DescribeTopicRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *DescribeTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicRuleResponseParams struct {
	// 规则描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rule *TopicRule `json:"Rule,omitempty" name:"Rule"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicRuleResponseParams `json:"Response"`
}

func (r *DescribeTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicRuleResponse) FromJsonString(s string) error {
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

	// 创建人Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserId *int64 `json:"CreateUserId,omitempty" name:"CreateUserId"`

	// 创建人昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorNickName *string `json:"CreatorNickName,omitempty" name:"CreatorNickName"`

	// 启用/禁用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableState *int64 `json:"EnableState,omitempty" name:"EnableState"`

	// 产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 设备类型（设备、子设备、网关）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 是否是 lora 设备
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLora *bool `json:"IsLora,omitempty" name:"IsLora"`
}

type DevicePositionItem struct {
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 位置信息时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 设备经度信息
	Longitude *float64 `json:"Longitude,omitempty" name:"Longitude"`

	// 设备纬度信息
	Latitude *float64 `json:"Latitude,omitempty" name:"Latitude"`
}

type DeviceSignatureInfo struct {
	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备签名
	DeviceSignature *string `json:"DeviceSignature,omitempty" name:"DeviceSignature"`
}

type DeviceUser struct {
	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户角色 1所有者，0：其他分享者
	Role *int64 `json:"Role,omitempty" name:"Role"`
}

type DevicesItem struct {
	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

// Predefined struct for user
type DirectBindDeviceInFamilyRequestParams struct {
	// 小程序appid
	IotAppID *string `json:"IotAppID,omitempty" name:"IotAppID"`

	// 用户ID
	UserID *string `json:"UserID,omitempty" name:"UserID"`

	// 家庭ID
	FamilyId *string `json:"FamilyId,omitempty" name:"FamilyId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 房间ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

type DirectBindDeviceInFamilyRequest struct {
	*tchttp.BaseRequest
	
	// 小程序appid
	IotAppID *string `json:"IotAppID,omitempty" name:"IotAppID"`

	// 用户ID
	UserID *string `json:"UserID,omitempty" name:"UserID"`

	// 家庭ID
	FamilyId *string `json:"FamilyId,omitempty" name:"FamilyId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 房间ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DirectBindDeviceInFamilyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DirectBindDeviceInFamilyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IotAppID")
	delete(f, "UserID")
	delete(f, "FamilyId")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DirectBindDeviceInFamilyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DirectBindDeviceInFamilyResponseParams struct {
	// 返回设备信息
	AppDeviceInfo *AppDeviceInfo `json:"AppDeviceInfo,omitempty" name:"AppDeviceInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DirectBindDeviceInFamilyResponse struct {
	*tchttp.BaseResponse
	Response *DirectBindDeviceInFamilyResponseParams `json:"Response"`
}

func (r *DirectBindDeviceInFamilyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DirectBindDeviceInFamilyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableTopicRuleRequestParams struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableTopicRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *DisableTopicRuleResponseParams `json:"Response"`
}

func (r *DisableTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableTopicRuleRequestParams struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableTopicRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *EnableTopicRuleResponseParams `json:"Response"`
}

func (r *EnableTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type FamilySubDevice struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 设备别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasName *string `json:"AliasName,omitempty" name:"AliasName"`

	// 设备绑定的家庭ID
	FamilyId *string `json:"FamilyId,omitempty" name:"FamilyId"`

	// 设备所在的房间ID，默认"0"
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 图标
	// 注意：此字段可能返回 null，表示取不到有效值。
	IconUrl *string `json:"IconUrl,omitempty" name:"IconUrl"`

	// grid图标
	// 注意：此字段可能返回 null，表示取不到有效值。
	IconUrlGrid *string `json:"IconUrlGrid,omitempty" name:"IconUrlGrid"`

	// 设备绑定时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 设备更新时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type FenceAlarmPoint struct {
	// 围栏告警时间
	AlarmTime *int64 `json:"AlarmTime,omitempty" name:"AlarmTime"`

	// 围栏告警位置的经度
	Longitude *float64 `json:"Longitude,omitempty" name:"Longitude"`

	// 围栏告警位置的纬度
	Latitude *float64 `json:"Latitude,omitempty" name:"Latitude"`
}

type FenceBindDeviceItem struct {
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 告警条件(In，进围栏报警；Out，出围栏报警；InOrOut，进围栏或者出围栏均报警)
	AlertCondition *string `json:"AlertCondition,omitempty" name:"AlertCondition"`

	// 是否使能围栏(true，使能；false，禁用)
	FenceEnable *bool `json:"FenceEnable,omitempty" name:"FenceEnable"`

	// 告警处理方法
	Method *string `json:"Method,omitempty" name:"Method"`
}

type FenceBindProductItem struct {
	// 围栏绑定的设备信息
	Devices []*FenceBindDeviceItem `json:"Devices,omitempty" name:"Devices"`

	// 围栏绑定的产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type FenceEventItem struct {
	// 围栏事件的产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 围栏事件的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`

	// 围栏事件的告警类型（In，进围栏报警；Out，出围栏报警；InOrOut，进围栏或者出围栏均报警）
	AlertType *string `json:"AlertType,omitempty" name:"AlertType"`

	// 围栏事件的设备位置信息
	Data *FenceAlarmPoint `json:"Data,omitempty" name:"Data"`
}

type FirmwareInfo struct {
	// 固件版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 固件MD5值
	Md5sum *string `json:"Md5sum,omitempty" name:"Md5sum"`

	// 固件创建时间
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 固件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 固件描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件升级模块
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwType *string `json:"FwType,omitempty" name:"FwType"`

	// 创建者子 uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserId *int64 `json:"CreateUserId,omitempty" name:"CreateUserId"`

	// 创建者昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorNickName *string `json:"CreatorNickName,omitempty" name:"CreatorNickName"`
}

// Predefined struct for user
type GenSingleDeviceSignatureOfPublicRequestParams struct {
	// 设备所属的产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 需要绑定的设备
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备绑定签名的有效时间,以秒为单位。取值范围：0 < Expire <= 86400，Expire == -1（十年）
	Expire *int64 `json:"Expire,omitempty" name:"Expire"`
}

type GenSingleDeviceSignatureOfPublicRequest struct {
	*tchttp.BaseRequest
	
	// 设备所属的产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 需要绑定的设备
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备绑定签名的有效时间,以秒为单位。取值范围：0 < Expire <= 86400，Expire == -1（十年）
	Expire *int64 `json:"Expire,omitempty" name:"Expire"`
}

func (r *GenSingleDeviceSignatureOfPublicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenSingleDeviceSignatureOfPublicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Expire")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenSingleDeviceSignatureOfPublicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenSingleDeviceSignatureOfPublicResponseParams struct {
	// 设备签名
	DeviceSignature *DeviceSignatureInfo `json:"DeviceSignature,omitempty" name:"DeviceSignature"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenSingleDeviceSignatureOfPublicResponse struct {
	*tchttp.BaseResponse
	Response *GenSingleDeviceSignatureOfPublicResponseParams `json:"Response"`
}

func (r *GenSingleDeviceSignatureOfPublicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenSingleDeviceSignatureOfPublicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBatchProductionsListRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type GetBatchProductionsListRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetBatchProductionsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBatchProductionsListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetBatchProductionsListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBatchProductionsListResponseParams struct {
	// 返回详情信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchProductions []*BatchProductionInfo `json:"BatchProductions,omitempty" name:"BatchProductions"`

	// 返回数量。
	TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetBatchProductionsListResponse struct {
	*tchttp.BaseResponse
	Response *GetBatchProductionsListResponseParams `json:"Response"`
}

func (r *GetBatchProductionsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBatchProductionsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCOSURLRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 文件大小
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`
}

type GetCOSURLRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 文件大小
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`
}

func (r *GetCOSURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCOSURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	delete(f, "FileSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCOSURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCOSURLResponseParams struct {
	// 固件URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetCOSURLResponse struct {
	*tchttp.BaseResponse
	Response *GetCOSURLResponseParams `json:"Response"`
}

func (r *GetCOSURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCOSURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceListRequestParams struct {
	// 需要查看设备列表的产品ID, -1代表ProjectId来筛选
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 设备固件版本号，若不带此参数会返回所有固件版本的设备。传"None-FirmwareVersion"查询无版本号的设备
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 需要过滤的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 项目ID。产品 ID 为 -1 时，该参数必填
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type GetDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// 需要查看设备列表的产品ID, -1代表ProjectId来筛选
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 设备固件版本号，若不带此参数会返回所有固件版本的设备。传"None-FirmwareVersion"查询无版本号的设备
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 需要过滤的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 项目ID。产品 ID 为 -1 时，该参数必填
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *GetDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FirmwareVersion")
	delete(f, "DeviceName")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceListResponseParams struct {
	// 返回的设备列表, 注意列表设备的 DevicePsk 为空, 要获取设备的 DevicePsk 请使用 DescribeDevice
	// 注意：此字段可能返回 null，表示取不到有效值。
	Devices []*DeviceInfo `json:"Devices,omitempty" name:"Devices"`

	// 产品下的设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceListResponseParams `json:"Response"`
}

func (r *GetDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceLocationHistoryRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 查询起始时间，Unix时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，Unix时间，单位为毫秒
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 坐标类型
	CoordinateType *int64 `json:"CoordinateType,omitempty" name:"CoordinateType"`
}

type GetDeviceLocationHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 查询起始时间，Unix时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，Unix时间，单位为毫秒
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 坐标类型
	CoordinateType *int64 `json:"CoordinateType,omitempty" name:"CoordinateType"`
}

func (r *GetDeviceLocationHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceLocationHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "CoordinateType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceLocationHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceLocationHistoryResponseParams struct {
	// 历史位置列表
	Positions []*PositionItem `json:"Positions,omitempty" name:"Positions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDeviceLocationHistoryResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceLocationHistoryResponseParams `json:"Response"`
}

func (r *GetDeviceLocationHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceLocationHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFamilyDeviceUserListRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type GetFamilyDeviceUserListRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *GetFamilyDeviceUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFamilyDeviceUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFamilyDeviceUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFamilyDeviceUserListResponseParams struct {
	// 设备的用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserList []*DeviceUser `json:"UserList,omitempty" name:"UserList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetFamilyDeviceUserListResponse struct {
	*tchttp.BaseResponse
	Response *GetFamilyDeviceUserListResponseParams `json:"Response"`
}

func (r *GetFamilyDeviceUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFamilyDeviceUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGatewaySubDeviceListRequestParams struct {
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备名称
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type GetGatewaySubDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备名称
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetGatewaySubDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGatewaySubDeviceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "GatewayDeviceName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGatewaySubDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGatewaySubDeviceListResponseParams struct {
	// 设备的总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 设备列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceList *FamilySubDevice `json:"DeviceList,omitempty" name:"DeviceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetGatewaySubDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *GetGatewaySubDeviceListResponseParams `json:"Response"`
}

func (r *GetGatewaySubDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGatewaySubDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLoRaGatewayListRequestParams struct {
	// 是否是社区网关
	IsCommunity *bool `json:"IsCommunity,omitempty" name:"IsCommunity"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type GetLoRaGatewayListRequest struct {
	*tchttp.BaseRequest
	
	// 是否是社区网关
	IsCommunity *bool `json:"IsCommunity,omitempty" name:"IsCommunity"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetLoRaGatewayListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLoRaGatewayListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsCommunity")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLoRaGatewayListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLoRaGatewayListResponseParams struct {
	// 返回总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 返回详情项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gateways []*LoRaGatewayItem `json:"Gateways,omitempty" name:"Gateways"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetLoRaGatewayListResponse struct {
	*tchttp.BaseResponse
	Response *GetLoRaGatewayListResponseParams `json:"Response"`
}

func (r *GetLoRaGatewayListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLoRaGatewayListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPositionSpaceListRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type GetPositionSpaceListRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetPositionSpaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPositionSpaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPositionSpaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPositionSpaceListResponseParams struct {
	// 位置空间列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*PositionSpaceInfo `json:"List,omitempty" name:"List"`

	// 位置空间数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPositionSpaceListResponse struct {
	*tchttp.BaseResponse
	Response *GetPositionSpaceListResponseParams `json:"Response"`
}

func (r *GetPositionSpaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPositionSpaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProjectListRequestParams struct {
	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 个数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按项目ID搜索
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 按产品ID搜索
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 加载 ProductCount、DeviceCount、ApplicationCount，可选值：ProductCount、DeviceCount、ApplicationCount，可多选
	Includes []*string `json:"Includes,omitempty" name:"Includes"`

	// 按项目名称搜索
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

type GetProjectListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 个数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按项目ID搜索
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 按产品ID搜索
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 加载 ProductCount、DeviceCount、ApplicationCount，可选值：ProductCount、DeviceCount、ApplicationCount，可多选
	Includes []*string `json:"Includes,omitempty" name:"Includes"`

	// 按项目名称搜索
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *GetProjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	delete(f, "ProjectId")
	delete(f, "ProductId")
	delete(f, "Includes")
	delete(f, "ProjectName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProjectListResponseParams struct {
	// 项目列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Projects []*ProjectEntryEx `json:"Projects,omitempty" name:"Projects"`

	// 列表项个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetProjectListResponse struct {
	*tchttp.BaseResponse
	Response *GetProjectListResponseParams `json:"Response"`
}

func (r *GetProjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetStudioProductListRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品DevStatus
	DevStatus *string `json:"DevStatus,omitempty" name:"DevStatus"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type GetStudioProductListRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品DevStatus
	DevStatus *string `json:"DevStatus,omitempty" name:"DevStatus"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetStudioProductListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetStudioProductListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DevStatus")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetStudioProductListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetStudioProductListResponseParams struct {
	// 产品列表
	Products []*ProductEntry `json:"Products,omitempty" name:"Products"`

	// 产品数量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetStudioProductListResponse struct {
	*tchttp.BaseResponse
	Response *GetStudioProductListResponseParams `json:"Response"`
}

func (r *GetStudioProductListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetStudioProductListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTopicRuleListRequestParams struct {
	// 请求的页数
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// 分页的大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type GetTopicRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 请求的页数
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// 分页的大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *GetTopicRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTopicRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTopicRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTopicRuleListResponseParams struct {
	// 规则总数量
	TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

	// 规则列表
	Rules []*TopicRuleInfo `json:"Rules,omitempty" name:"Rules"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTopicRuleListResponse struct {
	*tchttp.BaseResponse
	Response *GetTopicRuleListResponseParams `json:"Response"`
}

func (r *GetTopicRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTopicRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEventHistoryRequestParams struct {
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

	// 单次获取的历史数据项目的最大数量, 缺省10
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitempty" name:"EventId"`
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

	// 单次获取的历史数据项目的最大数量, 缺省10
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *ListEventHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEventHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Context")
	delete(f, "Size")
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEventHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEventHistoryResponseParams struct {
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
	EventHistory []*EventHistoryItem `json:"EventHistory,omitempty" name:"EventHistory"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListEventHistoryResponse struct {
	*tchttp.BaseResponse
	Response *ListEventHistoryResponseParams `json:"Response"`
}

func (r *ListEventHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEventHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListFirmwaresRequestParams struct {
	// 获取的页数
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 分页的大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 搜索过滤条件
	Filters []*SearchKeyword `json:"Filters,omitempty" name:"Filters"`
}

type ListFirmwaresRequest struct {
	*tchttp.BaseRequest
	
	// 获取的页数
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 分页的大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 搜索过滤条件
	Filters []*SearchKeyword `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListFirmwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListFirmwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "ProductID")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListFirmwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListFirmwaresResponseParams struct {
	// 固件总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 固件列表
	Firmwares []*FirmwareInfo `json:"Firmwares,omitempty" name:"Firmwares"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListFirmwaresResponse struct {
	*tchttp.BaseResponse
	Response *ListFirmwaresResponseParams `json:"Response"`
}

func (r *ListFirmwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListFirmwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopicPolicyRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type ListTopicPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *ListTopicPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopicPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTopicPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopicPolicyResponseParams struct {
	// Topic列表
	Topics []*TopicItem `json:"Topics,omitempty" name:"Topics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ListTopicPolicyResponseParams `json:"Response"`
}

func (r *ListTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopicPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoRaFrequencyEntry struct {
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitempty" name:"FreqId"`

	// 频点名称
	FreqName *string `json:"FreqName,omitempty" name:"FreqName"`

	// 频点描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据上行信道
	ChannelsDataUp []*uint64 `json:"ChannelsDataUp,omitempty" name:"ChannelsDataUp"`

	// 数据下行信道RX1
	ChannelsDataRX1 []*uint64 `json:"ChannelsDataRX1,omitempty" name:"ChannelsDataRX1"`

	// 数据下行信道RX2
	ChannelsDataRX2 []*uint64 `json:"ChannelsDataRX2,omitempty" name:"ChannelsDataRX2"`

	// 入网上行信道
	ChannelsJoinUp []*uint64 `json:"ChannelsJoinUp,omitempty" name:"ChannelsJoinUp"`

	// 入网下行RX1信道
	ChannelsJoinRX1 []*uint64 `json:"ChannelsJoinRX1,omitempty" name:"ChannelsJoinRX1"`

	// 入网下行RX2信道
	ChannelsJoinRX2 []*uint64 `json:"ChannelsJoinRX2,omitempty" name:"ChannelsJoinRX2"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type LoRaGatewayItem struct {
	// LoRa 网关Id
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// 是否是公开网关
	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

	// 网关描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 网关名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 网关位置信息
	Position *string `json:"Position,omitempty" name:"Position"`

	// 网关位置详情
	PositionDetails *string `json:"PositionDetails,omitempty" name:"PositionDetails"`

	// LoRa 网关位置坐标
	Location *LoRaGatewayLocation `json:"Location,omitempty" name:"Location"`

	// 最后更新时间
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 最后上报时间
	LastSeenAt *string `json:"LastSeenAt,omitempty" name:"LastSeenAt"`

	// 频点ID
	FrequencyId *string `json:"FrequencyId,omitempty" name:"FrequencyId"`
}

type LoRaGatewayLocation struct {
	// 纬度
	Latitude *float64 `json:"Latitude,omitempty" name:"Latitude"`

	// 精度
	Longitude *float64 `json:"Longitude,omitempty" name:"Longitude"`

	// 准确度
	Accuracy *float64 `json:"Accuracy,omitempty" name:"Accuracy"`

	// 海拔
	Altitude *float64 `json:"Altitude,omitempty" name:"Altitude"`
}

// Predefined struct for user
type ModifyFenceBindRequestParams struct {
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`

	// 围栏绑定的产品列表
	Items []*FenceBindProductItem `json:"Items,omitempty" name:"Items"`
}

type ModifyFenceBindRequest struct {
	*tchttp.BaseRequest
	
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`

	// 围栏绑定的产品列表
	Items []*FenceBindProductItem `json:"Items,omitempty" name:"Items"`
}

func (r *ModifyFenceBindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFenceBindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FenceId")
	delete(f, "Items")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFenceBindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFenceBindResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyFenceBindResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFenceBindResponseParams `json:"Response"`
}

func (r *ModifyFenceBindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFenceBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoRaFrequencyRequestParams struct {
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitempty" name:"FreqId"`

	// 频点名称
	FreqName *string `json:"FreqName,omitempty" name:"FreqName"`

	// 频点描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据上行信道
	ChannelsDataUp []*uint64 `json:"ChannelsDataUp,omitempty" name:"ChannelsDataUp"`

	// 数据下行信道RX1
	ChannelsDataRX1 []*uint64 `json:"ChannelsDataRX1,omitempty" name:"ChannelsDataRX1"`

	// 数据下行信道RX2
	ChannelsDataRX2 []*uint64 `json:"ChannelsDataRX2,omitempty" name:"ChannelsDataRX2"`

	// 入网上行信道
	ChannelsJoinUp []*uint64 `json:"ChannelsJoinUp,omitempty" name:"ChannelsJoinUp"`

	// 入网下行信道RX1
	ChannelsJoinRX1 []*uint64 `json:"ChannelsJoinRX1,omitempty" name:"ChannelsJoinRX1"`

	// 入网下行信道RX2
	ChannelsJoinRX2 []*uint64 `json:"ChannelsJoinRX2,omitempty" name:"ChannelsJoinRX2"`
}

type ModifyLoRaFrequencyRequest struct {
	*tchttp.BaseRequest
	
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitempty" name:"FreqId"`

	// 频点名称
	FreqName *string `json:"FreqName,omitempty" name:"FreqName"`

	// 频点描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据上行信道
	ChannelsDataUp []*uint64 `json:"ChannelsDataUp,omitempty" name:"ChannelsDataUp"`

	// 数据下行信道RX1
	ChannelsDataRX1 []*uint64 `json:"ChannelsDataRX1,omitempty" name:"ChannelsDataRX1"`

	// 数据下行信道RX2
	ChannelsDataRX2 []*uint64 `json:"ChannelsDataRX2,omitempty" name:"ChannelsDataRX2"`

	// 入网上行信道
	ChannelsJoinUp []*uint64 `json:"ChannelsJoinUp,omitempty" name:"ChannelsJoinUp"`

	// 入网下行信道RX1
	ChannelsJoinRX1 []*uint64 `json:"ChannelsJoinRX1,omitempty" name:"ChannelsJoinRX1"`

	// 入网下行信道RX2
	ChannelsJoinRX2 []*uint64 `json:"ChannelsJoinRX2,omitempty" name:"ChannelsJoinRX2"`
}

func (r *ModifyLoRaFrequencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoRaFrequencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FreqId")
	delete(f, "FreqName")
	delete(f, "Description")
	delete(f, "ChannelsDataUp")
	delete(f, "ChannelsDataRX1")
	delete(f, "ChannelsDataRX2")
	delete(f, "ChannelsJoinUp")
	delete(f, "ChannelsJoinRX1")
	delete(f, "ChannelsJoinRX2")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoRaFrequencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoRaFrequencyResponseParams struct {
	// 频点信息
	Data *LoRaFrequencyEntry `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLoRaFrequencyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoRaFrequencyResponseParams `json:"Response"`
}

func (r *ModifyLoRaFrequencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoRaFrequencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoRaGatewayRequestParams struct {
	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// LoRa网关Id
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// LoRa网关位置坐标
	Location *LoRaGatewayLocation `json:"Location,omitempty" name:"Location"`

	// LoRa网关名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否公开可见
	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

	// 位置信息
	Position *string `json:"Position,omitempty" name:"Position"`

	// 位置详情
	PositionDetails *string `json:"PositionDetails,omitempty" name:"PositionDetails"`

	// 频点ID
	FrequencyId *string `json:"FrequencyId,omitempty" name:"FrequencyId"`
}

type ModifyLoRaGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// LoRa网关Id
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// LoRa网关位置坐标
	Location *LoRaGatewayLocation `json:"Location,omitempty" name:"Location"`

	// LoRa网关名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否公开可见
	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

	// 位置信息
	Position *string `json:"Position,omitempty" name:"Position"`

	// 位置详情
	PositionDetails *string `json:"PositionDetails,omitempty" name:"PositionDetails"`

	// 频点ID
	FrequencyId *string `json:"FrequencyId,omitempty" name:"FrequencyId"`
}

func (r *ModifyLoRaGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoRaGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Description")
	delete(f, "GatewayId")
	delete(f, "Location")
	delete(f, "Name")
	delete(f, "IsPublic")
	delete(f, "Position")
	delete(f, "PositionDetails")
	delete(f, "FrequencyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoRaGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoRaGatewayResponseParams struct {
	// 返回网关数据
	Gateway *LoRaGatewayItem `json:"Gateway,omitempty" name:"Gateway"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLoRaGatewayResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoRaGatewayResponseParams `json:"Response"`
}

func (r *ModifyLoRaGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoRaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelDefinitionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 数据模板定义
	ModelSchema *string `json:"ModelSchema,omitempty" name:"ModelSchema"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelDefinitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ModelSchema")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelDefinitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelDefinitionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModelDefinitionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelDefinitionResponseParams `json:"Response"`
}

func (r *ModifyModelDefinitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelDefinitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPositionFenceRequestParams struct {

}

type ModifyPositionFenceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyPositionFenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPositionFenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPositionFenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPositionFenceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPositionFenceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPositionFenceResponseParams `json:"Response"`
}

func (r *ModifyPositionFenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPositionFenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPositionSpaceRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 位置空间名称
	SpaceName *string `json:"SpaceName,omitempty" name:"SpaceName"`

	// 授权类型
	AuthorizeType *int64 `json:"AuthorizeType,omitempty" name:"AuthorizeType"`

	// 产品列表
	ProductIdList []*string `json:"ProductIdList,omitempty" name:"ProductIdList"`

	// 位置空间描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 缩略图
	Icon *string `json:"Icon,omitempty" name:"Icon"`
}

type ModifyPositionSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 位置空间名称
	SpaceName *string `json:"SpaceName,omitempty" name:"SpaceName"`

	// 授权类型
	AuthorizeType *int64 `json:"AuthorizeType,omitempty" name:"AuthorizeType"`

	// 产品列表
	ProductIdList []*string `json:"ProductIdList,omitempty" name:"ProductIdList"`

	// 位置空间描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 缩略图
	Icon *string `json:"Icon,omitempty" name:"Icon"`
}

func (r *ModifyPositionSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPositionSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "SpaceName")
	delete(f, "AuthorizeType")
	delete(f, "ProductIdList")
	delete(f, "Description")
	delete(f, "Icon")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPositionSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPositionSpaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPositionSpaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPositionSpaceResponseParams `json:"Response"`
}

func (r *ModifyPositionSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPositionSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ProjectName")
	delete(f, "ProjectDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// 项目详情
	Project *ProjectEntry `json:"Project,omitempty" name:"Project"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifySpacePropertyRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品属性
	Data *string `json:"Data,omitempty" name:"Data"`
}

type ModifySpacePropertyRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品属性
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *ModifySpacePropertyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpacePropertyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "ProductId")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySpacePropertyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpacePropertyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySpacePropertyResponse struct {
	*tchttp.BaseResponse
	Response *ModifySpacePropertyResponseParams `json:"Response"`
}

func (r *ModifySpacePropertyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpacePropertyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStudioProductRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStudioProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ProductName")
	delete(f, "ProductDesc")
	delete(f, "ModuleId")
	delete(f, "EnableProductScript")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStudioProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStudioProductResponseParams struct {
	// 产品描述
	Product *ProductEntry `json:"Product,omitempty" name:"Product"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStudioProductResponseParams `json:"Response"`
}

func (r *ModifyStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStudioProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicPolicyRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 更新前Topic名
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 更新后Topic名
	NewTopicName *string `json:"NewTopicName,omitempty" name:"NewTopicName"`

	// Topic权限
	Privilege *uint64 `json:"Privilege,omitempty" name:"Privilege"`
}

type ModifyTopicPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 更新前Topic名
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 更新后Topic名
	NewTopicName *string `json:"NewTopicName,omitempty" name:"NewTopicName"`

	// Topic权限
	Privilege *uint64 `json:"Privilege,omitempty" name:"Privilege"`
}

func (r *ModifyTopicPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "TopicName")
	delete(f, "NewTopicName")
	delete(f, "Privilege")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicPolicyResponseParams `json:"Response"`
}

func (r *ModifyTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicRuleRequestParams struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 替换的规则包体
	TopicRulePayload *TopicRulePayload `json:"TopicRulePayload,omitempty" name:"TopicRulePayload"`
}

type ModifyTopicRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 替换的规则包体
	TopicRulePayload *TopicRulePayload `json:"TopicRulePayload,omitempty" name:"TopicRulePayload"`
}

func (r *ModifyTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	delete(f, "TopicRulePayload")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicRuleResponseParams `json:"Response"`
}

func (r *ModifyTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PositionFenceInfo struct {
	// 围栏信息
	GeoFence *PositionFenceItem `json:"GeoFence,omitempty" name:"GeoFence"`

	// 围栏创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 围栏更新时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type PositionFenceItem struct {
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitempty" name:"FenceId"`

	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 围栏名称
	FenceName *string `json:"FenceName,omitempty" name:"FenceName"`

	// 围栏描述
	FenceDesc *string `json:"FenceDesc,omitempty" name:"FenceDesc"`

	// 围栏区域信息，采用 GeoJSON 格式
	FenceArea *string `json:"FenceArea,omitempty" name:"FenceArea"`
}

type PositionItem struct {
	// 位置点的时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 位置点的经度
	Longitude *float64 `json:"Longitude,omitempty" name:"Longitude"`

	// 位置点的纬度
	Latitude *float64 `json:"Latitude,omitempty" name:"Latitude"`

	// 位置点的定位类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocationType *string `json:"LocationType,omitempty" name:"LocationType"`

	// 位置点的精度预估，单位为米
	// 注意：此字段可能返回 null，表示取不到有效值。
	Accuracy *float64 `json:"Accuracy,omitempty" name:"Accuracy"`
}

type PositionSpaceInfo struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitempty" name:"SpaceId"`

	// 位置空间名称
	SpaceName *string `json:"SpaceName,omitempty" name:"SpaceName"`

	// 授权类型
	AuthorizeType *int64 `json:"AuthorizeType,omitempty" name:"AuthorizeType"`

	// 描述备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 产品列表
	ProductIdList []*string `json:"ProductIdList,omitempty" name:"ProductIdList"`

	// 缩略图
	Icon *string `json:"Icon,omitempty" name:"Icon"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 用户自定义地图缩放
	Zoom *uint64 `json:"Zoom,omitempty" name:"Zoom"`
}

type ProductDevicesPositionItem struct {
	// 设备位置列表
	Items []*DevicePositionItem `json:"Items,omitempty" name:"Items"`

	// 产品标识
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备位置数量
	Total *int64 `json:"Total,omitempty" name:"Total"`
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

	// 创建人 UinId
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserId *int64 `json:"CreateUserId,omitempty" name:"CreateUserId"`

	// 创建者昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorNickName *string `json:"CreatorNickName,omitempty" name:"CreatorNickName"`
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

	// 产品的连接类型的模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetTypeModel *string `json:"NetTypeModel,omitempty" name:"NetTypeModel"`
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

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationCount *uint64 `json:"ApplicationCount,omitempty" name:"ApplicationCount"`

	// 设备注册总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceCount *uint64 `json:"DeviceCount,omitempty" name:"DeviceCount"`
}

// Predefined struct for user
type PublishBroadcastMessageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 消息内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// 消息质量等级
	Qos *int64 `json:"Qos,omitempty" name:"Qos"`

	// ayload内容的编码格式，取值为base64或空。base64表示云端将收到的请求数据进行base64解码后下发到设备，空则直接将原始内容下发到设备
	PayloadEncoding *string `json:"PayloadEncoding,omitempty" name:"PayloadEncoding"`
}

type PublishBroadcastMessageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 消息内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// 消息质量等级
	Qos *int64 `json:"Qos,omitempty" name:"Qos"`

	// ayload内容的编码格式，取值为base64或空。base64表示云端将收到的请求数据进行base64解码后下发到设备，空则直接将原始内容下发到设备
	PayloadEncoding *string `json:"PayloadEncoding,omitempty" name:"PayloadEncoding"`
}

func (r *PublishBroadcastMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishBroadcastMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Payload")
	delete(f, "Qos")
	delete(f, "PayloadEncoding")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishBroadcastMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishBroadcastMessageResponseParams struct {
	// 广播消息任务Id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PublishBroadcastMessageResponse struct {
	*tchttp.BaseResponse
	Response *PublishBroadcastMessageResponseParams `json:"Response"`
}

func (r *PublishBroadcastMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishBroadcastMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishMessageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 消息发往的主题
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 云端下发到设备的控制报文
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// 消息服务质量等级，取值为0或1
	Qos *uint64 `json:"Qos,omitempty" name:"Qos"`

	// Payload的内容编码格式，取值为base64或空。base64表示云端将接收到的base64编码后的报文再转换成二进制报文下发至设备，为空表示不作转换，透传下发至设备
	PayloadEncoding *string `json:"PayloadEncoding,omitempty" name:"PayloadEncoding"`
}

type PublishMessageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 消息发往的主题
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 云端下发到设备的控制报文
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// 消息服务质量等级，取值为0或1
	Qos *uint64 `json:"Qos,omitempty" name:"Qos"`

	// Payload的内容编码格式，取值为base64或空。base64表示云端将接收到的base64编码后的报文再转换成二进制报文下发至设备，为空表示不作转换，透传下发至设备
	PayloadEncoding *string `json:"PayloadEncoding,omitempty" name:"PayloadEncoding"`
}

func (r *PublishMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Topic")
	delete(f, "Payload")
	delete(f, "Qos")
	delete(f, "PayloadEncoding")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishMessageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PublishMessageResponse struct {
	*tchttp.BaseResponse
	Response *PublishMessageResponseParams `json:"Response"`
}

func (r *PublishMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishRRPCMessageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 消息内容，utf8编码
	Payload *string `json:"Payload,omitempty" name:"Payload"`
}

type PublishRRPCMessageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 消息内容，utf8编码
	Payload *string `json:"Payload,omitempty" name:"Payload"`
}

func (r *PublishRRPCMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishRRPCMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Payload")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishRRPCMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishRRPCMessageResponseParams struct {
	// RRPC消息ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageId *int64 `json:"MessageId,omitempty" name:"MessageId"`

	// 设备回复的消息内容，采用base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayloadBase64 *string `json:"PayloadBase64,omitempty" name:"PayloadBase64"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PublishRRPCMessageResponse struct {
	*tchttp.BaseResponse
	Response *PublishRRPCMessageResponseParams `json:"Response"`
}

func (r *PublishRRPCMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishRRPCMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseStudioProductRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品DevStatus
	DevStatus *string `json:"DevStatus,omitempty" name:"DevStatus"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseStudioProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DevStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseStudioProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseStudioProductResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseStudioProductResponseParams `json:"Response"`
}

func (r *ReleaseStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseStudioProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchKeyword struct {
	// 搜索条件的Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 搜索条件的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type SearchPositionSpaceRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 位置空间名字
	SpaceName *string `json:"SpaceName,omitempty" name:"SpaceName"`

	// 偏移量，从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大获取数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type SearchPositionSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 位置空间名字
	SpaceName *string `json:"SpaceName,omitempty" name:"SpaceName"`

	// 偏移量，从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大获取数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *SearchPositionSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPositionSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SpaceName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchPositionSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchPositionSpaceResponseParams struct {
	// 位置空间列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*PositionSpaceInfo `json:"List,omitempty" name:"List"`

	// 符合条件的位置空间个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchPositionSpaceResponse struct {
	*tchttp.BaseResponse
	Response *SearchPositionSpaceResponseParams `json:"Response"`
}

func (r *SearchPositionSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPositionSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchStudioProductRequestParams struct {
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

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
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

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *SearchStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchStudioProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ProductName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "DevStatus")
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchStudioProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchStudioProductResponseParams struct {
	// 产品列表
	Products []*ProductEntry `json:"Products,omitempty" name:"Products"`

	// 产品数量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *SearchStudioProductResponseParams `json:"Response"`
}

func (r *SearchStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchStudioProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchTopicRuleRequestParams struct {
	// 规则名
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type SearchTopicRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

func (r *SearchTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchTopicRuleResponseParams struct {
	// 搜索到的规则总数
	TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

	// 规则信息列表
	Rules []*TopicRuleInfo `json:"Rules,omitempty" name:"Rules"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *SearchTopicRuleResponseParams `json:"Response"`
}

func (r *SearchTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopicItem struct {
	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Topic权限 , 1上报  2下发
	Privilege *uint64 `json:"Privilege,omitempty" name:"Privilege"`
}

type TopicRule struct {
	// 规则名称。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则的SQL语句，如： SELECT * FROM 'pid/dname/event'，然后对其进行base64编码，得：U0VMRUNUICogRlJPTSAncGlkL2RuYW1lL2V2ZW50Jw==
	Sql *string `json:"Sql,omitempty" name:"Sql"`

	// 规则描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 行为的JSON字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Actions *string `json:"Actions,omitempty" name:"Actions"`

	// 是否禁用规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleDisabled *bool `json:"RuleDisabled,omitempty" name:"RuleDisabled"`
}

type TopicRuleInfo struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	CreatedAt *int64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 规则是否禁用
	RuleDisabled *bool `json:"RuleDisabled,omitempty" name:"RuleDisabled"`
}

type TopicRulePayload struct {
	// 规则的SQL语句，如： SELECT * FROM 'pid/dname/event'，然后对其进行base64编码，得：U0VMRUNUICogRlJPTSAncGlkL2RuYW1lL2V2ZW50Jw==
	Sql *string `json:"Sql,omitempty" name:"Sql"`

	// 行为的JSON字符串，大部分种类举例如下：
	// [
	// {
	// "republish": {
	// "topic": "TEST/test"
	// }
	// },
	// {
	// "forward": {
	// "api": "http://test.com:8080"
	// }
	// },
	// {
	// "ckafka": {
	// "instance": {
	// "id": "ckafka-test",
	// "name": ""
	// },
	// "topic": {
	// "id": "topic-test",
	// "name": "test"
	// },
	// "region": "gz"
	// }
	// },
	// {
	// "cmqqueue": {
	// "queuename": "queue-test-TEST",
	// "region": "gz"
	// }
	// },
	// {
	// "mysql": {
	// "instanceid": "cdb-test",
	// "region": "gz",
	// "username": "test",
	// "userpwd": "*****",
	// "dbname": "d_mqtt",
	// "tablename": "t_test",
	// "fieldpairs": [
	// {
	// "field": "test",
	// "value": "test"
	// }
	// ],
	// "devicetype": "CUSTOM"
	// }
	// }
	// ]
	Actions *string `json:"Actions,omitempty" name:"Actions"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否禁用规则
	RuleDisabled *bool `json:"RuleDisabled,omitempty" name:"RuleDisabled"`
}

// Predefined struct for user
type UnbindDevicesRequestParams struct {
	// 网关设备的产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名列表
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`
}

type UnbindDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 网关设备的产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名列表
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`
}

func (r *UnbindDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "GatewayDeviceName")
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindDevicesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindDevicesResponse struct {
	*tchttp.BaseResponse
	Response *UnbindDevicesResponseParams `json:"Response"`
}

func (r *UnbindDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindProductsRequestParams struct {
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 待解绑的子产品ID数组
	ProductIds []*string `json:"ProductIds,omitempty" name:"ProductIds"`
}

type UnbindProductsRequest struct {
	*tchttp.BaseRequest
	
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 待解绑的子产品ID数组
	ProductIds []*string `json:"ProductIds,omitempty" name:"ProductIds"`
}

func (r *UnbindProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "ProductIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindProductsResponseParams struct {
	// 绑定了待解绑的LoRa产品下的设备的网关设备列表
	GatewayDeviceNames []*string `json:"GatewayDeviceNames,omitempty" name:"GatewayDeviceNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindProductsResponse struct {
	*tchttp.BaseResponse
	Response *UnbindProductsResponseParams `json:"Response"`
}

func (r *UnbindProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDevicesEnableStateRequestParams struct {
	// 多个设备标识
	DevicesItems []*DevicesItem `json:"DevicesItems,omitempty" name:"DevicesItems"`

	// 1：启用；0：禁用
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type UpdateDevicesEnableStateRequest struct {
	*tchttp.BaseRequest
	
	// 多个设备标识
	DevicesItems []*DevicesItem `json:"DevicesItems,omitempty" name:"DevicesItems"`

	// 1：启用；0：禁用
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateDevicesEnableStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDevicesEnableStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DevicesItems")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDevicesEnableStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDevicesEnableStateResponseParams struct {
	// 删除的结果代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultCode *string `json:"ResultCode,omitempty" name:"ResultCode"`

	// 删除的结果信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultMessage *string `json:"ResultMessage,omitempty" name:"ResultMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateDevicesEnableStateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDevicesEnableStateResponseParams `json:"Response"`
}

func (r *UpdateDevicesEnableStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDevicesEnableStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFirmwareRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 固件新的版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件原版本号
	FirmwareOriVersion *string `json:"FirmwareOriVersion,omitempty" name:"FirmwareOriVersion"`

	// 固件升级方式；0 静默升级 1 用户确认升级   不填默认静默升级
	UpgradeMethod *uint64 `json:"UpgradeMethod,omitempty" name:"UpgradeMethod"`
}

type UpdateFirmwareRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 固件新的版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件原版本号
	FirmwareOriVersion *string `json:"FirmwareOriVersion,omitempty" name:"FirmwareOriVersion"`

	// 固件升级方式；0 静默升级 1 用户确认升级   不填默认静默升级
	UpgradeMethod *uint64 `json:"UpgradeMethod,omitempty" name:"UpgradeMethod"`
}

func (r *UpdateFirmwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFirmwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "DeviceName")
	delete(f, "FirmwareVersion")
	delete(f, "FirmwareOriVersion")
	delete(f, "UpgradeMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateFirmwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFirmwareResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *UpdateFirmwareResponseParams `json:"Response"`
}

func (r *UpdateFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFirmwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFirmwareRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件的MD5值
	Md5sum *string `json:"Md5sum,omitempty" name:"Md5sum"`

	// 固件的大小
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 固件名称
	FirmwareName *string `json:"FirmwareName,omitempty" name:"FirmwareName"`

	// 固件描述
	FirmwareDescription *string `json:"FirmwareDescription,omitempty" name:"FirmwareDescription"`

	// 固件升级模块；可选值 mcu|moudule
	FwType *string `json:"FwType,omitempty" name:"FwType"`
}

type UploadFirmwareRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件的MD5值
	Md5sum *string `json:"Md5sum,omitempty" name:"Md5sum"`

	// 固件的大小
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 固件名称
	FirmwareName *string `json:"FirmwareName,omitempty" name:"FirmwareName"`

	// 固件描述
	FirmwareDescription *string `json:"FirmwareDescription,omitempty" name:"FirmwareDescription"`

	// 固件升级模块；可选值 mcu|moudule
	FwType *string `json:"FwType,omitempty" name:"FwType"`
}

func (r *UploadFirmwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFirmwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	delete(f, "Md5sum")
	delete(f, "FileSize")
	delete(f, "FirmwareName")
	delete(f, "FirmwareDescription")
	delete(f, "FwType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFirmwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFirmwareResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *UploadFirmwareResponseParams `json:"Response"`
}

func (r *UploadFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFirmwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WifiInfo struct {
	// mac地址
	MAC *string `json:"MAC,omitempty" name:"MAC"`

	// 信号强度
	RSSI *int64 `json:"RSSI,omitempty" name:"RSSI"`
}