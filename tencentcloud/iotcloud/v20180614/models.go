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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Attribute struct {
	// 属性列表
	Tags []*DeviceTag `json:"Tags,omitempty" name:"Tags"`
}

type BatchPublishMessage struct {
	// 消息发往的主题。为 Topic 权限中去除 ProductID 和 DeviceName 的部分，如 “event”
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 消息内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`
}

// Predefined struct for user
type BatchUpdateFirmwareRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件新版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件原版本号，根据文件列表升级固件不需要填写此参数
	FirmwareOriVersion *string `json:"FirmwareOriVersion,omitempty" name:"FirmwareOriVersion"`

	// 升级方式，0 静默升级  1 用户确认升级。 不填默认为静默升级方式
	UpgradeMethod *uint64 `json:"UpgradeMethod,omitempty" name:"UpgradeMethod"`

	// 设备列表文件名称，根据文件列表升级固件需要填写此参数
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 设备列表的文件md5值
	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`

	// 设备列表的文件大小值
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 需要升级的设备名称列表
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 固件升级任务，默认超时时间。 最小取值60秒，最大为3600秒
	TimeoutInterval *uint64 `json:"TimeoutInterval,omitempty" name:"TimeoutInterval"`
}

type BatchUpdateFirmwareRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件新版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件原版本号，根据文件列表升级固件不需要填写此参数
	FirmwareOriVersion *string `json:"FirmwareOriVersion,omitempty" name:"FirmwareOriVersion"`

	// 升级方式，0 静默升级  1 用户确认升级。 不填默认为静默升级方式
	UpgradeMethod *uint64 `json:"UpgradeMethod,omitempty" name:"UpgradeMethod"`

	// 设备列表文件名称，根据文件列表升级固件需要填写此参数
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 设备列表的文件md5值
	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`

	// 设备列表的文件大小值
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 需要升级的设备名称列表
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 固件升级任务，默认超时时间。 最小取值60秒，最大为3600秒
	TimeoutInterval *uint64 `json:"TimeoutInterval,omitempty" name:"TimeoutInterval"`
}

func (r *BatchUpdateFirmwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchUpdateFirmwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	delete(f, "FirmwareOriVersion")
	delete(f, "UpgradeMethod")
	delete(f, "FileName")
	delete(f, "FileMd5")
	delete(f, "FileSize")
	delete(f, "DeviceNames")
	delete(f, "TimeoutInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchUpdateFirmwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchUpdateFirmwareResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchUpdateFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *BatchUpdateFirmwareResponseParams `json:"Response"`
}

func (r *BatchUpdateFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchUpdateFirmwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchUpdateShadow struct {
	// 设备影子的期望状态，格式为 Json 对象序列化之后的字符串
	Desired *string `json:"Desired,omitempty" name:"Desired"`
}

// Predefined struct for user
type BindDevicesRequestParams struct {
	// 网关设备的产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 被绑定设备的产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 被绑定的多个设备名
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 中兴CLAA设备的绑定需要skey，普通的设备不需要
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

type BindDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 网关设备的产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 被绑定设备的产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 被绑定的多个设备名
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 中兴CLAA设备的绑定需要skey，普通的设备不需要
	Skey *string `json:"Skey,omitempty" name:"Skey"`
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
	delete(f, "Skey")
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

type BrokerSubscribe struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type CLSLogItem struct {
	// 日志内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 设备名称
	Devicename *string `json:"Devicename,omitempty" name:"Devicename"`

	// 产品ID
	Productid *string `json:"Productid,omitempty" name:"Productid"`

	// 请求ID
	Requestid *string `json:"Requestid,omitempty" name:"Requestid"`

	// 结果
	Result *string `json:"Result,omitempty" name:"Result"`

	// 模块
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 日志时间
	Time *string `json:"Time,omitempty" name:"Time"`

	// 腾讯云账号
	Userid *string `json:"Userid,omitempty" name:"Userid"`
}

// Predefined struct for user
type CancelDeviceFirmwareTaskRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件升级任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type CancelDeviceFirmwareTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件升级任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *CancelDeviceFirmwareTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDeviceFirmwareTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "DeviceName")
	delete(f, "FirmwareVersion")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelDeviceFirmwareTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelDeviceFirmwareTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelDeviceFirmwareTaskResponse struct {
	*tchttp.BaseResponse
	Response *CancelDeviceFirmwareTaskResponseParams `json:"Response"`
}

func (r *CancelDeviceFirmwareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDeviceFirmwareTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelTaskRequestParams struct {
	// 任务 ID
	Id *string `json:"Id,omitempty" name:"Id"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelTaskResponse struct {
	*tchttp.BaseResponse
	Response *CancelTaskResponseParams `json:"Response"`
}

func (r *CancelTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Attribute")
	delete(f, "DefinedPsk")
	delete(f, "Isp")
	delete(f, "Imei")
	delete(f, "LoraDevEui")
	delete(f, "LoraMoteType")
	delete(f, "Skey")
	delete(f, "LoraAppKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceResponseParams struct {
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
type CreateLoraDeviceRequestParams struct {
	// 产品Id ，创建产品时腾讯云为用户分配全局唯一的Id
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

type CreateLoraDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id ，创建产品时腾讯云为用户分配全局唯一的Id
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoraDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "DeviceType")
	delete(f, "AppEui")
	delete(f, "DeviceEui")
	delete(f, "AppKey")
	delete(f, "AuthKey")
	delete(f, "Memo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoraDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoraDeviceResponseParams struct {
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
}

type CreateLoraDeviceResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoraDeviceResponseParams `json:"Response"`
}

func (r *CreateLoraDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoraDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiDeviceRequestParams struct {
	// 产品 ID。创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 批量创建的设备名数组，单次最多创建 100 个设备。命名规则：[a-zA-Z0-9:_-]{1,48}
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`
}

type CreateMultiDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID。创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 批量创建的设备名数组，单次最多创建 100 个设备。命名规则：[a-zA-Z0-9:_-]{1,48}
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`
}

func (r *CreateMultiDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMultiDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiDeviceResponseParams struct {
	// 任务ID，腾讯云生成全局唯一的任务 ID，有效期一个月，一个月之后任务失效。可以调用获取创建多设备任务状态接口获取该任务的执行状态，当状态为成功时，可以调用获取创建多设备任务结果接口获取该任务的结果
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMultiDeviceResponse struct {
	*tchttp.BaseResponse
	Response *CreateMultiDeviceResponseParams `json:"Response"`
}

func (r *CreateMultiDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiDevicesTaskRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 参数类型 cosfile-文件上传 random-随机创建
	ParametersType *string `json:"ParametersType,omitempty" name:"ParametersType"`

	// 文件上传类型时文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件上传类型时文件大小
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 随机创建时设备创建个数
	BatchCount *uint64 `json:"BatchCount,omitempty" name:"BatchCount"`

	// 文件上传类型时文件md5值
	Hash *string `json:"Hash,omitempty" name:"Hash"`
}

type CreateMultiDevicesTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 参数类型 cosfile-文件上传 random-随机创建
	ParametersType *string `json:"ParametersType,omitempty" name:"ParametersType"`

	// 文件上传类型时文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件上传类型时文件大小
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 随机创建时设备创建个数
	BatchCount *uint64 `json:"BatchCount,omitempty" name:"BatchCount"`

	// 文件上传类型时文件md5值
	Hash *string `json:"Hash,omitempty" name:"Hash"`
}

func (r *CreateMultiDevicesTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiDevicesTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ParametersType")
	delete(f, "FileName")
	delete(f, "FileSize")
	delete(f, "BatchCount")
	delete(f, "Hash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMultiDevicesTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiDevicesTaskResponseParams struct {
	// 任务ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMultiDevicesTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateMultiDevicesTaskResponseParams `json:"Response"`
}

func (r *CreateMultiDevicesTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiDevicesTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProductRequestParams struct {
	// 产品名称，名称不能和已经存在的产品名称重复。命名规则：[a-zA-Z0-9:_-]{1,32}
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品属性
	ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`

	// 创建CLAA产品时，需要Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

type CreateProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品名称，名称不能和已经存在的产品名称重复。命名规则：[a-zA-Z0-9:_-]{1,32}
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品属性
	ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`

	// 创建CLAA产品时，需要Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *CreateProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductName")
	delete(f, "ProductProperties")
	delete(f, "Skey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProductResponseParams struct {
	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品 ID，腾讯云生成全局唯一 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品属性
	ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProductResponse struct {
	*tchttp.BaseResponse
	Response *CreateProductResponseParams `json:"Response"`
}

func (r *CreateProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFileUrlRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type CreateTaskFileUrlRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *CreateTaskFileUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFileUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskFileUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFileUrlResponseParams struct {
	// 任务文件上传链接
	Url *string `json:"Url,omitempty" name:"Url"`

	// 任务文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTaskFileUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskFileUrlResponseParams `json:"Response"`
}

func (r *CreateTaskFileUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFileUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "ProductId")
	delete(f, "DeviceNameFilter")
	delete(f, "ScheduleTimeInSeconds")
	delete(f, "Tasks")
	delete(f, "MaxExecutionTimeInSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskResponseParams struct {
	// 创建的任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskResponseParams `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicPolicyRequestParams struct {
	// 产品自身ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Topic权限，1发布，2订阅，3订阅和发布
	Privilege *uint64 `json:"Privilege,omitempty" name:"Privilege"`

	// 代理订阅信息，网关产品为绑定的子产品创建topic时需要填写，内容为子产品的ID和设备信息。
	BrokerSubscribe *BrokerSubscribe `json:"BrokerSubscribe,omitempty" name:"BrokerSubscribe"`
}

type CreateTopicPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 产品自身ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// Topic名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Topic权限，1发布，2订阅，3订阅和发布
	Privilege *uint64 `json:"Privilege,omitempty" name:"Privilege"`

	// 代理订阅信息，网关产品为绑定的子产品创建topic时需要填写，内容为子产品的ID和设备信息。
	BrokerSubscribe *BrokerSubscribe `json:"BrokerSubscribe,omitempty" name:"BrokerSubscribe"`
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
	delete(f, "ProductID")
	delete(f, "TopicName")
	delete(f, "Privilege")
	delete(f, "BrokerSubscribe")
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
	// 设备所属的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 需要删除的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 删除LoRa设备以及LoRa网关设备需要skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Skey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceResourceRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DeleteDeviceResourceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DeleteDeviceResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "Name")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceResourceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDeviceResourceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeviceResourceResponseParams `json:"Response"`
}

func (r *DeleteDeviceResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceResponseParams struct {
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
type DeleteLoraDeviceRequestParams struct {
	// 设备所属产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DeleteLoraDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备所属产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DeleteLoraDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoraDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoraDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoraDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLoraDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoraDeviceResponseParams `json:"Response"`
}

func (r *DeleteLoraDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoraDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProductRequestParams struct {
	// 需要删除的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 删除LoRa产品需要skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Skey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProductResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteProductResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProductResponseParams `json:"Response"`
}

func (r *DeleteProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProductResponse) FromJsonString(s string) error {
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
type DescribeAllDevicesRequestParams struct {
	// 查询偏移量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询设备数量。最大支持250个
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAllDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 查询偏移量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询设备数量。最大支持250个
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAllDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllDevicesResponseParams struct {
	// 设备总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 查询的设备列表信息。
	Devices []*DeviceProperty `json:"Devices,omitempty" name:"Devices"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAllDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllDevicesResponseParams `json:"Response"`
}

func (r *DescribeAllDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceClientKeyRequestParams struct {
	// 所属产品的Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DescribeDeviceClientKeyRequest struct {
	*tchttp.BaseRequest
	
	// 所属产品的Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DescribeDeviceClientKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceClientKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceClientKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceClientKeyResponseParams struct {
	// 设备的私钥
	ClientKey *string `json:"ClientKey,omitempty" name:"ClientKey"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceClientKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceClientKeyResponseParams `json:"Response"`
}

func (r *DescribeDeviceClientKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceClientKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceResourceRequestParams struct {
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 具体的设备资源名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeDeviceResourceRequest struct {
	*tchttp.BaseRequest
	
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 具体的设备资源名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeDeviceResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceName")
	delete(f, "ProductID")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceResourceResponseParams struct {
	// 设备资源详情
	Result *DeviceResourceInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceResourceResponseParams `json:"Response"`
}

func (r *DescribeDeviceResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceResourcesRequestParams struct {
	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-250
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 需要过滤的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 资源搜索开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 资源搜索结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeDeviceResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-250
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 需要过滤的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 资源搜索开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 资源搜索结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDeviceResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProductID")
	delete(f, "DeviceName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceResourcesResponseParams struct {
	// 资源总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 资源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*DeviceResourceInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceResourcesResponseParams `json:"Response"`
}

func (r *DescribeDeviceResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceResponseParams struct {
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
	Tags []*DeviceTag `json:"Tags,omitempty" name:"Tags"`

	// 设备类型
	DeviceType *uint64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 国际移动设备识别码 IMEI
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

	// 首次上线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstOnlineTime *uint64 `json:"FirstOnlineTime,omitempty" name:"FirstOnlineTime"`

	// 最近下线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOfflineTime *uint64 `json:"LastOfflineTime,omitempty" name:"LastOfflineTime"`

	// 设备创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 设备证书获取状态，0 未获取过设备密钥, 1 已获取过设备密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertState *uint64 `json:"CertState,omitempty" name:"CertState"`

	// 设备启用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`

	// 设备标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*DeviceLabel `json:"Labels,omitempty" name:"Labels"`

	// MQTT客户端IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`

	// 设备固件更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirmwareUpdateTime *uint64 `json:"FirmwareUpdateTime,omitempty" name:"FirmwareUpdateTime"`

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
type DescribeDeviceShadowRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称。命名规则：[a-zA-Z0-9:_-]{1,48}
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceShadowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceShadowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceShadowResponseParams struct {
	// 设备影子数据
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceShadowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceShadowResponseParams `json:"Response"`
}

func (r *DescribeDeviceShadowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceShadowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesRequestParams struct {
	// 需要查看设备列表的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-250
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 设备固件版本号，若不带此参数会返回所有固件版本的设备。传"None-FirmwareVersion"查询无版本号的设备
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 需要过滤的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备是否启用，0禁用状态1启用状态，默认不区分
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查看设备列表的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-250
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 设备固件版本号，若不带此参数会返回所有固件版本的设备。传"None-FirmwareVersion"查询无版本号的设备
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 需要过滤的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备是否启用，0禁用状态1启用状态，默认不区分
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FirmwareVersion")
	delete(f, "DeviceName")
	delete(f, "EnableState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesResponseParams struct {
	// 设备总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 设备详细信息列表
	Devices []*DeviceInfo `json:"Devices,omitempty" name:"Devices"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicesResponseParams `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`
}

type DescribeFirmwareRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`
}

func (r *DescribeFirmwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirmwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareResponseParams struct {
	// 固件版本号
	Version *string `json:"Version,omitempty" name:"Version"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 固件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 固件描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 固件Md5值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5sum *string `json:"Md5sum,omitempty" name:"Md5sum"`

	// 固件上传的秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Createtime *uint64 `json:"Createtime,omitempty" name:"Createtime"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 固件类型。选项：mcu、module
	FwType *string `json:"FwType,omitempty" name:"FwType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirmwareResponseParams `json:"Response"`
}

func (r *DescribeFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareTaskDevicesRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 筛选条件
	Filters []*SearchKeyword `json:"Filters,omitempty" name:"Filters"`

	// 查询偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeFirmwareTaskDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 筛选条件
	Filters []*SearchKeyword `json:"Filters,omitempty" name:"Filters"`

	// 查询偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFirmwareTaskDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareTaskDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirmwareTaskDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareTaskDevicesResponseParams struct {
	// 固件升级任务的设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 固件升级任务的设备列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Devices []*DeviceUpdateStatus `json:"Devices,omitempty" name:"Devices"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFirmwareTaskDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirmwareTaskDevicesResponseParams `json:"Response"`
}

func (r *DescribeFirmwareTaskDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareTaskDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareTaskDistributionRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件升级任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeFirmwareTaskDistributionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件升级任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeFirmwareTaskDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareTaskDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirmwareTaskDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareTaskDistributionResponseParams struct {
	// 固件升级任务状态分布信息
	StatusInfos []*StatusStatistic `json:"StatusInfos,omitempty" name:"StatusInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFirmwareTaskDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirmwareTaskDistributionResponseParams `json:"Response"`
}

func (r *DescribeFirmwareTaskDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareTaskDistributionResponse) FromJsonString(s string) error {
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

	// 固件任务创建时间，单位:秒
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
type DescribeFirmwareTaskStatisticsRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`
}

type DescribeFirmwareTaskStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`
}

func (r *DescribeFirmwareTaskStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareTaskStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirmwareTaskStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareTaskStatisticsResponseParams struct {
	// 升级成功的设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessTotal *uint64 `json:"SuccessTotal,omitempty" name:"SuccessTotal"`

	// 升级失败的设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureTotal *uint64 `json:"FailureTotal,omitempty" name:"FailureTotal"`

	// 正在升级的设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradingTotal *uint64 `json:"UpgradingTotal,omitempty" name:"UpgradingTotal"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFirmwareTaskStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirmwareTaskStatisticsResponseParams `json:"Response"`
}

func (r *DescribeFirmwareTaskStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareTaskStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareTasksRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 查询偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回查询结果条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索过滤条件
	Filters []*SearchKeyword `json:"Filters,omitempty" name:"Filters"`
}

type DescribeFirmwareTasksRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 查询偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回查询结果条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索过滤条件
	Filters []*SearchKeyword `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeFirmwareTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirmwareTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareTasksResponseParams struct {
	// 固件升级任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskInfos []*FirmwareTaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos"`

	// 固件升级任务总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFirmwareTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirmwareTasksResponseParams `json:"Response"`
}

func (r *DescribeFirmwareTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoraDeviceRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DescribeLoraDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DescribeLoraDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoraDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoraDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoraDeviceResponseParams struct {
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
}

type DescribeLoraDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoraDeviceResponseParams `json:"Response"`
}

func (r *DescribeLoraDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoraDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiDevTaskRequestParams struct {
	// 任务 ID，由批量创建设备接口返回
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 产品 ID，创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiDevTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMultiDevTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiDevTaskResponseParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务是否完成。0 代表任务未开始，1 代表任务正在执行，2 代表任务已完成
	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMultiDevTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMultiDevTaskResponseParams `json:"Response"`
}

func (r *DescribeMultiDevTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiDevTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiDevicesRequestParams struct {
	// 产品 ID，创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 任务 ID，由批量创建设备接口返回
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，每页返回的设备个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "TaskId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMultiDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiDevicesResponseParams struct {
	// 任务 ID，由批量创建设备接口返回
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 设备详细信息列表
	DevicesInfo []*MultiDevicesInfo `json:"DevicesInfo,omitempty" name:"DevicesInfo"`

	// 该任务创建设备的总数
	TotalDevNum *uint64 `json:"TotalDevNum,omitempty" name:"TotalDevNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMultiDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMultiDevicesResponseParams `json:"Response"`
}

func (r *DescribeMultiDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribeProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductResourceRequestParams struct {
	// 需要查看资源列表的产品 ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 需要过滤的资源名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeProductResourceRequest struct {
	*tchttp.BaseRequest
	
	// 需要查看资源列表的产品 ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 需要过滤的资源名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeProductResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductResourceResponseParams struct {
	// 资源详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ProductResourceInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductResourceResponseParams `json:"Response"`
}

func (r *DescribeProductResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductResourcesRequestParams struct {
	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-250
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 需要查看资源列表的产品 ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 需要过滤的资源名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeProductResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-250
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 需要查看资源列表的产品 ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 需要过滤的资源名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeProductResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProductID")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductResourcesResponseParams struct {
	// 资源总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 资源详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*ProductResourceInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductResourcesResponseParams `json:"Response"`
}

func (r *DescribeProductResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductResponseParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品元数据
	ProductMetadata *ProductMetadata `json:"ProductMetadata,omitempty" name:"ProductMetadata"`

	// 产品属性
	ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductResponseParams `json:"Response"`
}

func (r *DescribeProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductTaskRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeProductTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeProductTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductTaskResponseParams struct {
	// 产品任务详细信息
	TaskInfo *ProductTaskInfo `json:"TaskInfo,omitempty" name:"TaskInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductTaskResponseParams `json:"Response"`
}

func (r *DescribeProductTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductTasksRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品级别任务列表偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 产品级别任务列表拉取个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeProductTasksRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品级别任务列表偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 产品级别任务列表拉取个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProductTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductTasksResponseParams struct {
	// 符合条件的任务总个数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务详细信息列表
	TaskInfos []*ProductTaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductTasksResponseParams `json:"Response"`
}

func (r *DescribeProductTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductsRequestParams struct {
	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，当前页面中显示的最大数量，值范围 10-250。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeProductsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，当前页面中显示的最大数量，值范围 10-250。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductsResponseParams struct {
	// 产品总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 产品详细信息列表
	Products []*ProductInfo `json:"Products,omitempty" name:"Products"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductsResponseParams `json:"Response"`
}

func (r *DescribeProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePushResourceTaskStatisticsRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribePushResourceTaskStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribePushResourceTaskStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePushResourceTaskStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePushResourceTaskStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePushResourceTaskStatisticsResponseParams struct {
	// 推送成功的设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessTotal *uint64 `json:"SuccessTotal,omitempty" name:"SuccessTotal"`

	// 推送失败的设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureTotal *uint64 `json:"FailureTotal,omitempty" name:"FailureTotal"`

	// 正在推送的设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradingTotal *uint64 `json:"UpgradingTotal,omitempty" name:"UpgradingTotal"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePushResourceTaskStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePushResourceTaskStatisticsResponseParams `json:"Response"`
}

func (r *DescribePushResourceTaskStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePushResourceTaskStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTasksRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 查询偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回查询结果条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索过滤条件
	Filters []*SearchKeyword `json:"Filters,omitempty" name:"Filters"`
}

type DescribeResourceTasksRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 查询偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回查询结果条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索过滤条件
	Filters []*SearchKeyword `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeResourceTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTasksResponseParams struct {
	// 资源任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskInfos []*FirmwareTaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos"`

	// 资源任务总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceTasksResponseParams `json:"Response"`
}

func (r *DescribeResourceTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskRequestParams struct {
	// 任务ID
	Id *string `json:"Id,omitempty" name:"Id"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResponseParams struct {
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
type DescribeTasksRequestParams struct {
	// 偏移量，从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 1-250
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 1-250
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 用户一个月内创建的任务总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 此页任务对象的数组，按创建时间排序
	Tasks []*TaskInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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
	Tags []*DeviceTag `json:"Tags,omitempty" name:"Tags"`

	// 设备类型
	DeviceType *uint64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 国际移动设备识别码 IMEI
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

	// 设备证书获取状态, 1 已获取过设备密钥，0 未获取过设备密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertState *uint64 `json:"CertState,omitempty" name:"CertState"`

	// 设备可用状态，0禁用，1启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`

	// 设备标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*DeviceLabel `json:"Labels,omitempty" name:"Labels"`

	// MQTT客户端IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`

	// ota最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirmwareUpdateTime *uint64 `json:"FirmwareUpdateTime,omitempty" name:"FirmwareUpdateTime"`
}

type DeviceLabel struct {
	// 标签标识
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeviceProperty struct {
	// 产品ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 设备名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备资源ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DeviceResourceInfo struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 产品名
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资源文件md5
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 资源文件大小
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 资源更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备资源上传状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 设备资源上传百分比
	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`
}

type DeviceTag struct {
	// 属性名称
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 属性值的类型，1 int，2 string
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 属性的值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 属性描述名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeviceUpdateStatus struct {
	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 最后处理时间
	LastProcessTime *uint64 `json:"LastProcessTime,omitempty" name:"LastProcessTime"`

	// 状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 错误消息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 返回码
	Retcode *int64 `json:"Retcode,omitempty" name:"Retcode"`

	// 目标更新版本
	DstVersion *string `json:"DstVersion,omitempty" name:"DstVersion"`

	// 下载中状态时的下载进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`

	// 原版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriVersion *string `json:"OriVersion,omitempty" name:"OriVersion"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
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
type DownloadDeviceResourceRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DownloadDeviceResourceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DownloadDeviceResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadDeviceResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "Name")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadDeviceResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadDeviceResourceResponseParams struct {
	// 设备资源的cos链接
	Url *string `json:"Url,omitempty" name:"Url"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadDeviceResourceResponse struct {
	*tchttp.BaseResponse
	Response *DownloadDeviceResourceResponseParams `json:"Response"`
}

func (r *DownloadDeviceResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadDeviceResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditFirmwareRequestParams struct {
	// 产品ID。
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号。
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件名称。
	FirmwareName *string `json:"FirmwareName,omitempty" name:"FirmwareName"`

	// 固件描述
	FirmwareDescription *string `json:"FirmwareDescription,omitempty" name:"FirmwareDescription"`

	// 固件类型：选填 mcu、moudule。默认：mcu
	FwType *string `json:"FwType,omitempty" name:"FwType"`
}

type EditFirmwareRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID。
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号。
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件名称。
	FirmwareName *string `json:"FirmwareName,omitempty" name:"FirmwareName"`

	// 固件描述
	FirmwareDescription *string `json:"FirmwareDescription,omitempty" name:"FirmwareDescription"`

	// 固件类型：选填 mcu、moudule。默认：mcu
	FwType *string `json:"FwType,omitempty" name:"FwType"`
}

func (r *EditFirmwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditFirmwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	delete(f, "FirmwareName")
	delete(f, "FirmwareDescription")
	delete(f, "FwType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditFirmwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditFirmwareResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EditFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *EditFirmwareResponseParams `json:"Response"`
}

func (r *EditFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditFirmwareResponse) FromJsonString(s string) error {
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

type FirmwareTaskInfo struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 任务创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type GetCOSURLRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`
}

type GetCOSURLRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`
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
type GetUserResourceInfoRequestParams struct {

}

type GetUserResourceInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetUserResourceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserResourceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUserResourceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserResourceInfoResponseParams struct {
	// 已使用的资源字节数
	UsedSize *uint64 `json:"UsedSize,omitempty" name:"UsedSize"`

	// 可以使用资源的总大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetUserResourceInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetUserResourceInfoResponseParams `json:"Response"`
}

func (r *GetUserResourceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserResourceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListLogPayloadRequestParams struct {
	// 日志开始时间，毫秒级时间戳
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 日志结束时间，毫秒级时间戳
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 查询关键字，可以同时支持键值查询和文本查询，例如，查询某key的值为value，并且包含某word的日志，该参数为：key:value word。键值或文本可以包含多个，以空格隔开。其中可以索引的key比如：RequestID、ProductID、DeviceName等。
	// 一个典型的查询示例：ProductID:ABCDE12345 DeviceName:test publish
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`

	// 日志检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志最大条数
	MaxNum *int64 `json:"MaxNum,omitempty" name:"MaxNum"`
}

type ListLogPayloadRequest struct {
	*tchttp.BaseRequest
	
	// 日志开始时间，毫秒级时间戳
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 日志结束时间，毫秒级时间戳
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 查询关键字，可以同时支持键值查询和文本查询，例如，查询某key的值为value，并且包含某word的日志，该参数为：key:value word。键值或文本可以包含多个，以空格隔开。其中可以索引的key比如：RequestID、ProductID、DeviceName等。
	// 一个典型的查询示例：ProductID:ABCDE12345 DeviceName:test publish
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`

	// 日志检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志最大条数
	MaxNum *int64 `json:"MaxNum,omitempty" name:"MaxNum"`
}

func (r *ListLogPayloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLogPayloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinTime")
	delete(f, "MaxTime")
	delete(f, "Keywords")
	delete(f, "Context")
	delete(f, "MaxNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListLogPayloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListLogPayloadResponseParams struct {
	// 日志上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 是否还有日志，如有仍有日志，下次查询的请求带上当前请求返回的Context
	Listover *bool `json:"Listover,omitempty" name:"Listover"`

	// 日志列表
	Results []*PayloadLogItem `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListLogPayloadResponse struct {
	*tchttp.BaseResponse
	Response *ListLogPayloadResponseParams `json:"Response"`
}

func (r *ListLogPayloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLogPayloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListLogRequestParams struct {
	// 日志开始时间，毫秒级时间戳
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 日志结束时间，毫秒级时间戳
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 查询关键字，可以同时支持键值查询和文本查询，例如，查询某key的值为value，并且包含某word的日志，该参数为：key:value word。键值或文本可以包含多个，以空格隔开。其中可以索引的key包括：requestid、productid、devicename、scene、content。
	// 一个典型的查询示例：productid:ABCDE12345 devicename:test scene:SHADOW content:Device%20connect publish
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`

	// 日志检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 查询条数
	MaxNum *int64 `json:"MaxNum,omitempty" name:"MaxNum"`
}

type ListLogRequest struct {
	*tchttp.BaseRequest
	
	// 日志开始时间，毫秒级时间戳
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 日志结束时间，毫秒级时间戳
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 查询关键字，可以同时支持键值查询和文本查询，例如，查询某key的值为value，并且包含某word的日志，该参数为：key:value word。键值或文本可以包含多个，以空格隔开。其中可以索引的key包括：requestid、productid、devicename、scene、content。
	// 一个典型的查询示例：productid:ABCDE12345 devicename:test scene:SHADOW content:Device%20connect publish
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`

	// 日志检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 查询条数
	MaxNum *int64 `json:"MaxNum,omitempty" name:"MaxNum"`
}

func (r *ListLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinTime")
	delete(f, "MaxTime")
	delete(f, "Keywords")
	delete(f, "Context")
	delete(f, "MaxNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListLogResponseParams struct {
	// 日志上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 是否还有日志，如有仍有日志，下次查询的请求带上当前请求返回的Context
	Listover *bool `json:"Listover,omitempty" name:"Listover"`

	// 日志列表
	Results []*CLSLogItem `json:"Results,omitempty" name:"Results"`

	// 日志总条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListLogResponse struct {
	*tchttp.BaseResponse
	Response *ListLogResponseParams `json:"Response"`
}

func (r *ListLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSDKLogRequestParams struct {
	// 日志开始时间
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 日志结束时间
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 查询关键字，可以同时支持键值查询和文本查询，
	// 例如，查询某key的值为value，并且包含某word的日志，该参数为：key:value word。
	// 键值或文本可以包含多个，以空格隔开。
	// 其中可以索引的key包括：productid、devicename、loglevel
	// 一个典型的查询示例：productid:7JK1G72JNE devicename:name publish loglevel:WARN一个典型的查询示例：productid:ABCDE12345 devicename:test scene:SHADOW publish
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`

	// 日志检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 查询条数
	MaxNum *uint64 `json:"MaxNum,omitempty" name:"MaxNum"`
}

type ListSDKLogRequest struct {
	*tchttp.BaseRequest
	
	// 日志开始时间
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 日志结束时间
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 查询关键字，可以同时支持键值查询和文本查询，
	// 例如，查询某key的值为value，并且包含某word的日志，该参数为：key:value word。
	// 键值或文本可以包含多个，以空格隔开。
	// 其中可以索引的key包括：productid、devicename、loglevel
	// 一个典型的查询示例：productid:7JK1G72JNE devicename:name publish loglevel:WARN一个典型的查询示例：productid:ABCDE12345 devicename:test scene:SHADOW publish
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`

	// 日志检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 查询条数
	MaxNum *uint64 `json:"MaxNum,omitempty" name:"MaxNum"`
}

func (r *ListSDKLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSDKLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinTime")
	delete(f, "MaxTime")
	delete(f, "Keywords")
	delete(f, "Context")
	delete(f, "MaxNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSDKLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSDKLogResponseParams struct {
	// 日志检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 是否还有日志，如有仍有日志，下次查询的请求带上当前请求返回的Context
	Listover *bool `json:"Listover,omitempty" name:"Listover"`

	// 日志列表
	Results []*SDKLogItem `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListSDKLogResponse struct {
	*tchttp.BaseResponse
	Response *ListSDKLogResponseParams `json:"Response"`
}

func (r *ListSDKLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSDKLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type PayloadLogItem struct {
	// 账号id
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 产品id
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 来源类型
	SrcType *string `json:"SrcType,omitempty" name:"SrcType"`

	// 来源名称
	SrcName *string `json:"SrcName,omitempty" name:"SrcName"`

	// 消息topic
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 内容格式类型
	PayloadFmtType *string `json:"PayloadFmtType,omitempty" name:"PayloadFmtType"`

	// 内容信息
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// 请求ID
	RequestID *string `json:"RequestID,omitempty" name:"RequestID"`

	// 日期时间
	DateTime *string `json:"DateTime,omitempty" name:"DateTime"`
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

type ProductResourceInfo struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 产品名
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 资源名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资源文件md5
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 资源文件大小
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 资源文件描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 资源创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ProductTaskInfo struct {
	// 任务ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 任务类型 0-批量创建设备类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 任务状态 0-创建中 1-待执行 2-执行中 3-执行失败 4-子任务部分失败 5-执行成功
	State *uint64 `json:"State,omitempty" name:"State"`

	// 任务参数类型 cosfile-文件输入 random-随机生成
	ParametersType *string `json:"ParametersType,omitempty" name:"ParametersType"`

	// 任务参数
	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`

	// 任务执行结果类型 cosfile-文件输出 errmsg-错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultType *string `json:"ResultType,omitempty" name:"ResultType"`

	// 任务执行结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 子任务总个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchCount *uint64 `json:"BatchCount,omitempty" name:"BatchCount"`

	// 子任务已执行个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchOffset *uint64 `json:"BatchOffset,omitempty" name:"BatchOffset"`

	// 任务创建时间
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务更新时间
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 任务完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompleteTime *uint64 `json:"CompleteTime,omitempty" name:"CompleteTime"`
}

// Predefined struct for user
type PublishAsDeviceRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// LoRa 设备端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 消息内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`
}

type PublishAsDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishAsDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Port")
	delete(f, "Payload")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishAsDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishAsDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PublishAsDeviceResponse struct {
	*tchttp.BaseResponse
	Response *PublishAsDeviceResponseParams `json:"Response"`
}

func (r *PublishAsDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishAsDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishBroadcastMessageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 消息内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// 消息质量等级
	Qos *int64 `json:"Qos,omitempty" name:"Qos"`

	// Payload内容的编码格式，取值为base64或空。base64表示云端将收到的请求数据进行base64解码后下发到设备，空则直接将原始内容下发到设备
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

	// Payload内容的编码格式，取值为base64或空。base64表示云端将收到的请求数据进行base64解码后下发到设备，空则直接将原始内容下发到设备
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
	// 广播消息任务ID
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

	// Payload内容的编码格式，取值为base64或空。base64表示云端将收到的请求数据进行base64解码后下发到设备，空则直接将原始内容下发到设备
	PayloadEncoding *string `json:"PayloadEncoding,omitempty" name:"PayloadEncoding"`
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

	// Payload内容的编码格式，取值为base64或空。base64表示云端将收到的请求数据进行base64解码后下发到设备，空则直接将原始内容下发到设备
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
	delete(f, "Topic")
	delete(f, "Payload")
	delete(f, "ProductId")
	delete(f, "DeviceName")
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
	MessageId *int64 `json:"MessageId,omitempty" name:"MessageId"`

	// 设备回复的消息内容，采用base64编码
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
type PublishToDeviceRequestParams struct {
	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// LoRa 端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 消息内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishToDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Port")
	delete(f, "Payload")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishToDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishToDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PublishToDeviceResponse struct {
	*tchttp.BaseResponse
	Response *PublishToDeviceResponseParams `json:"Response"`
}

func (r *PublishToDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishToDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceTopicRuleRequestParams struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 替换的规则包体
	TopicRulePayload *TopicRulePayload `json:"TopicRulePayload,omitempty" name:"TopicRulePayload"`

	// 修改类型，0：其他，1：创建行为，2：更新行为，3：删除行为
	ModifyType *uint64 `json:"ModifyType,omitempty" name:"ModifyType"`

	// action增删改变更填对应topicRulePayload里面第几个action
	ActionIndex *uint64 `json:"ActionIndex,omitempty" name:"ActionIndex"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	delete(f, "TopicRulePayload")
	delete(f, "ModifyType")
	delete(f, "ActionIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceTopicRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReplaceTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceTopicRuleResponseParams `json:"Response"`
}

func (r *ReplaceTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetDeviceResult struct {
	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 是否成功
	Success *bool `json:"Success,omitempty" name:"Success"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

// Predefined struct for user
type ResetDeviceStateRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`
}

type ResetDeviceStateRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`
}

func (r *ResetDeviceStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDeviceStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetDeviceStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetDeviceStateResponseParams struct {
	// 批量重置设备成功数
	SuccessCount *uint64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 批量重置设备结果
	ResetDeviceResults []*ResetDeviceResult `json:"ResetDeviceResults,omitempty" name:"ResetDeviceResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetDeviceStateResponse struct {
	*tchttp.BaseResponse
	Response *ResetDeviceStateResponseParams `json:"Response"`
}

func (r *ResetDeviceStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDeviceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDeviceFirmwareTaskRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件升级任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type RetryDeviceFirmwareTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件升级任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RetryDeviceFirmwareTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDeviceFirmwareTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "DeviceName")
	delete(f, "FirmwareVersion")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryDeviceFirmwareTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDeviceFirmwareTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RetryDeviceFirmwareTaskResponse struct {
	*tchttp.BaseResponse
	Response *RetryDeviceFirmwareTaskResponseParams `json:"Response"`
}

func (r *RetryDeviceFirmwareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDeviceFirmwareTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SDKLogItem struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 日志等级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 日志时间
	DateTime *string `json:"DateTime,omitempty" name:"DateTime"`

	// 日志内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

type SearchKeyword struct {
	// 搜索条件的Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 搜索条件的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type SetProductsForbiddenStatusRequestParams struct {
	// 要设置禁用状态的产品列表
	ProductID []*string `json:"ProductID,omitempty" name:"ProductID"`

	// 0启用，1禁用
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type SetProductsForbiddenStatusRequest struct {
	*tchttp.BaseRequest
	
	// 要设置禁用状态的产品列表
	ProductID []*string `json:"ProductID,omitempty" name:"ProductID"`

	// 0启用，1禁用
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SetProductsForbiddenStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetProductsForbiddenStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetProductsForbiddenStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetProductsForbiddenStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetProductsForbiddenStatusResponse struct {
	*tchttp.BaseResponse
	Response *SetProductsForbiddenStatusResponseParams `json:"Response"`
}

func (r *SetProductsForbiddenStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetProductsForbiddenStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatusStatistic struct {
	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 统计总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`
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
	// [
	//     {
	//         "republish": {
	//             "topic": "TEST/test"
	//         }
	//     },
	//     {
	//         "forward": {
	//             "api": "http://127.0.0.1:8080",
	//             "token":"xxx"
	//         }
	//     },
	//     {
	//         "ckafka": {
	//             "instance": {
	//                 "id": "ckafka-test",
	//                 "name": ""
	//             },
	//             "topic": {
	//                 "id": "topic-test",
	//                 "name": "test"
	//             },
	//             "region": "gz"
	//         }
	//     },
	//     {
	//         "cmqqueue": {
	//             "queuename": "queue-test-TEST",
	//             "region": "gz"
	//         }
	//     },
	//     {
	//         "mysql": {
	//             "instanceid": "cdb-test",
	//             "region": "gz",
	//             "username": "test",
	//             "userpwd": "*****",
	//             "dbname": "d_mqtt",
	//             "tablename": "t_test",
	//             "fieldpairs": [
	//                 {
	//                     "field": "test",
	//                     "value": "test"
	//                 }
	//             ],
	//             "devicetype": "CUSTOM"
	//         }
	//     }
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

	// 多个设备名
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 中兴CLAA设备的解绑需要Skey，普通设备不需要
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

type UnbindDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 网关设备的产品ID
	GatewayProductId *string `json:"GatewayProductId,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名
	GatewayDeviceName *string `json:"GatewayDeviceName,omitempty" name:"GatewayDeviceName"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 多个设备名
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 中兴CLAA设备的解绑需要Skey，普通设备不需要
	Skey *string `json:"Skey,omitempty" name:"Skey"`
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
	delete(f, "Skey")
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
type UpdateDeviceAvailableStateRequestParams struct {
	// 设备所属产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 要设置的设备状态，1为启用，0为禁用
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`
}

type UpdateDeviceAvailableStateRequest struct {
	*tchttp.BaseRequest
	
	// 设备所属产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 要设置的设备状态，1为启用，0为禁用
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`
}

func (r *UpdateDeviceAvailableStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceAvailableStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "EnableState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDeviceAvailableStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceAvailableStateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateDeviceAvailableStateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDeviceAvailableStateResponseParams `json:"Response"`
}

func (r *UpdateDeviceAvailableStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceAvailableStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceShadowRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 虚拟设备的状态，JSON字符串格式，由desired结构组成
	State *string `json:"State,omitempty" name:"State"`

	// 当前版本号，需要和后台的version保持一致，才能更新成功
	ShadowVersion *uint64 `json:"ShadowVersion,omitempty" name:"ShadowVersion"`

	// 下发delta消息的topic前缀，可选类型: "$shadow","$template"。不填写默认"$shadow"。
	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
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

	// 下发delta消息的topic前缀，可选类型: "$shadow","$template"。不填写默认"$shadow"。
	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
}

func (r *UpdateDeviceShadowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceShadowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "State")
	delete(f, "ShadowVersion")
	delete(f, "Prefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDeviceShadowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceShadowResponseParams struct {
	// 设备影子数据，JSON字符串格式
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateDeviceShadowResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDeviceShadowResponseParams `json:"Response"`
}

func (r *UpdateDeviceShadowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceShadowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDevicesEnableStateRequestParams struct {
	// 设备所属产品id
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 设备名称集合
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 要设置的设备状态，1为启用，0为禁用
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type UpdateDevicesEnableStateRequest struct {
	*tchttp.BaseRequest
	
	// 设备所属产品id
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 设备名称集合
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 要设置的设备状态，1为启用，0为禁用
	Status *uint64 `json:"Status,omitempty" name:"Status"`
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
	delete(f, "ProductID")
	delete(f, "DeviceNames")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDevicesEnableStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDevicesEnableStateResponseParams struct {
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
type UpdateProductDynamicRegisterRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 动态注册类型，0-关闭 1-预创建设备 2-自动创建设备
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// 动态注册设备上限
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`
}

type UpdateProductDynamicRegisterRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 动态注册类型，0-关闭 1-预创建设备 2-自动创建设备
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// 动态注册设备上限
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`
}

func (r *UpdateProductDynamicRegisterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProductDynamicRegisterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "RegisterType")
	delete(f, "RegisterLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateProductDynamicRegisterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProductDynamicRegisterResponseParams struct {
	// 动态注册类型，0-关闭 1-预创建设备 2-自动创建设备
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// 动态注册产品密钥
	ProductSecret *string `json:"ProductSecret,omitempty" name:"ProductSecret"`

	// 动态注册设备上限
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateProductDynamicRegisterResponse struct {
	*tchttp.BaseResponse
	Response *UpdateProductDynamicRegisterResponseParams `json:"Response"`
}

func (r *UpdateProductDynamicRegisterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProductDynamicRegisterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTopicPolicyRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTopicPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "TopicName")
	delete(f, "NewTopicName")
	delete(f, "Privilege")
	delete(f, "BrokerSubscribe")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTopicPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTopicPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTopicPolicyResponseParams `json:"Response"`
}

func (r *UpdateTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTopicPolicyResponse) FromJsonString(s string) error {
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

	// 固件类型：选填 mcu、moudule。默认：mcu
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

	// 固件类型：选填 mcu、moudule。默认：mcu
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