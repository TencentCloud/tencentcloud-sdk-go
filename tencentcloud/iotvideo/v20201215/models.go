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

package v20201215

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AIModelApplication struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 申请状态：1-已申请；2-已取消；3-已拒绝；4-已通过
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type AIModelInfo struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 申请状态：1-已申请；2-已取消；3-已拒绝；4-已通过
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 可调用数量
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 已调用数量
	Used *uint64 `json:"Used,omitempty" name:"Used"`

	// 申请时间
	ApplyTime *uint64 `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// 审批通过时间
	ApprovalTime *uint64 `json:"ApprovalTime,omitempty" name:"ApprovalTime"`
}

type AIModelUsageInfo struct {
	// 开通时间
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 资源总量
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 已使用资源数量
	Used *uint64 `json:"Used,omitempty" name:"Used"`
}

type ActionHistory struct {
	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 动作Id
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 动作名称
	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`

	// 请求时间
	ReqTime *uint64 `json:"ReqTime,omitempty" name:"ReqTime"`

	// 响应时间
	RspTime *uint64 `json:"RspTime,omitempty" name:"RspTime"`

	// 输入参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputParams *string `json:"InputParams,omitempty" name:"InputParams"`

	// 输出参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputParams *string `json:"OutputParams,omitempty" name:"OutputParams"`

	// 调用方式
	Calling *string `json:"Calling,omitempty" name:"Calling"`

	// 调用Id
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 调用状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type ApplyAIModelRequestParams struct {
	// AI模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type ApplyAIModelRequest struct {
	*tchttp.BaseRequest
	
	// AI模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *ApplyAIModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyAIModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyAIModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyAIModelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyAIModelResponse struct {
	*tchttp.BaseResponse
	Response *ApplyAIModelResponseParams `json:"Response"`
}

func (r *ApplyAIModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyAIModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BalanceTransaction struct {
	// 账户类型：1-设备接入 2-云存。
	AccountType *uint64 `json:"AccountType,omitempty" name:"AccountType"`

	// 账户变更类型：Rechareg-充值；CreateOrder-新购。
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 流水ID。
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 变更金额，单位：分（人民币）。
	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`

	// 变更后账户余额，单位：分（人民币）。
	Balance *uint64 `json:"Balance,omitempty" name:"Balance"`

	// 变更时间。
	OperationTime *int64 `json:"OperationTime,omitempty" name:"OperationTime"`
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

// Predefined struct for user
type BindCloudStorageUserRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type BindCloudStorageUserRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *BindCloudStorageUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudStorageUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindCloudStorageUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindCloudStorageUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindCloudStorageUserResponse struct {
	*tchttp.BaseResponse
	Response *BindCloudStorageUserResponseParams `json:"Response"`
}

func (r *BindCloudStorageUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudStorageUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelAIModelApplicationRequestParams struct {
	// AI模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type CancelAIModelApplicationRequest struct {
	*tchttp.BaseRequest
	
	// AI模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *CancelAIModelApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAIModelApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelAIModelApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelAIModelApplicationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelAIModelApplicationResponse struct {
	*tchttp.BaseResponse
	Response *CancelAIModelApplicationResponseParams `json:"Response"`
}

func (r *CancelAIModelApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAIModelApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
type CheckForwardAuthRequestParams struct {
	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 队列类型 0.CMQ  1.Ckafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`
}

type CheckForwardAuthRequest struct {
	*tchttp.BaseRequest
	
	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 队列类型 0.CMQ  1.Ckafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`
}

func (r *CheckForwardAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckForwardAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Skey")
	delete(f, "QueueType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckForwardAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckForwardAuthResponseParams struct {
	// 腾讯云账号
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 结果
	Result *uint64 `json:"Result,omitempty" name:"Result"`

	// 产品ID
	Productid *string `json:"Productid,omitempty" name:"Productid"`

	// 错误消息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 队列类型 0.CMQ  1.Ckafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckForwardAuthResponse struct {
	*tchttp.BaseResponse
	Response *CheckForwardAuthResponseParams `json:"Response"`
}

func (r *CheckForwardAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckForwardAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudStorageEvent struct {
	// 事件起始时间（Unix 时间戳，秒级
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 事件结束时间（Unix 时间戳，秒级
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 事件缩略图
	Thumbnail *string `json:"Thumbnail,omitempty" name:"Thumbnail"`

	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type CloudStorageTimeData struct {
	// 云存时间轴信息列表
	TimeList []*CloudStorageTimeInfo `json:"TimeList,omitempty" name:"TimeList"`

	// 播放地址
	VideoURL *string `json:"VideoURL,omitempty" name:"VideoURL"`
}

type CloudStorageTimeInfo struct {
	// 开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

type CloudStorageUserInfo struct {
	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
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
	// Sent = 1 表示设备已经在线并且订阅了控制下发的mqtt topic
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
type CreateAIDetectionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// AI模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 图片上传的开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 图片上传的结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

type CreateAIDetectionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// AI模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 图片上传的开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 图片上传的结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *CreateAIDetectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIDetectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ModelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIDetectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIDetectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAIDetectionResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIDetectionResponseParams `json:"Response"`
}

func (r *CreateAIDetectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIDetectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 批次创建的设备数量
	DevNum *uint64 `json:"DevNum,omitempty" name:"DevNum"`

	// 批次创建的设备前缀。不超过24个字符
	DevPre *string `json:"DevPre,omitempty" name:"DevPre"`
}

type CreateBatchRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 批次创建的设备数量
	DevNum *uint64 `json:"DevNum,omitempty" name:"DevNum"`

	// 批次创建的设备前缀。不超过24个字符
	DevPre *string `json:"DevPre,omitempty" name:"DevPre"`
}

func (r *CreateBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DevNum")
	delete(f, "DevPre")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchResponseParams struct {
	// 批次ID
	BatchId *uint64 `json:"BatchId,omitempty" name:"BatchId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBatchResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchResponseParams `json:"Response"`
}

func (r *CreateBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCOSCredentialsRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type CreateCOSCredentialsRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *CreateCOSCredentialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCOSCredentialsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCOSCredentialsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCOSCredentialsResponseParams struct {
	// COS存储桶名称
	StorageBucket *string `json:"StorageBucket,omitempty" name:"StorageBucket"`

	// COS存储桶区域
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// COS存储桶路径
	StoragePath *string `json:"StoragePath,omitempty" name:"StoragePath"`

	// COS上传用的SecretID
	SecretID *string `json:"SecretID,omitempty" name:"SecretID"`

	// COS上传用的SecretKey
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// COS上传用的Token
	Token *string `json:"Token,omitempty" name:"Token"`

	// 密钥信息过期时间
	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCOSCredentialsResponse struct {
	*tchttp.BaseResponse
	Response *CreateCOSCredentialsResponseParams `json:"Response"`
}

func (r *CreateCOSCredentialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCOSCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudStorageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 云存套餐ID：
	// yc1m3d ： 全时3天存储月套餐。
	// yc1m7d ： 全时7天存储月套餐。
	// yc1m30d ：全时30天存储月套餐。
	// yc1y3d ：全时3天存储年套餐。
	// yc1y7d ：全时7天存储年套餐。
	// yc1y30d ：全时30天存储年套餐。
	// ye1m3d ：事件3天存储月套餐。
	// ye1m7d ：事件7天存储月套餐。
	// ye1m30d ：事件30天存储月套餐 。
	// ye1y3d ：事件3天存储年套餐。
	// ye1y7d ：事件7天存储年套餐。
	// ye1y30d ：事件30天存储年套餐。
	// yc1w7d : 全时7天存储周套餐。
	// ye1w7d : 事件7天存储周套餐。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 如果当前设备已开启云存套餐，Override=1会使用新套餐覆盖原有套餐。不传此参数则默认为0。
	Override *uint64 `json:"Override,omitempty" name:"Override"`
}

type CreateCloudStorageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 云存套餐ID：
	// yc1m3d ： 全时3天存储月套餐。
	// yc1m7d ： 全时7天存储月套餐。
	// yc1m30d ：全时30天存储月套餐。
	// yc1y3d ：全时3天存储年套餐。
	// yc1y7d ：全时7天存储年套餐。
	// yc1y30d ：全时30天存储年套餐。
	// ye1m3d ：事件3天存储月套餐。
	// ye1m7d ：事件7天存储月套餐。
	// ye1m30d ：事件30天存储月套餐 。
	// ye1y3d ：事件3天存储年套餐。
	// ye1y7d ：事件7天存储年套餐。
	// ye1y30d ：事件30天存储年套餐。
	// yc1w7d : 全时7天存储周套餐。
	// ye1w7d : 事件7天存储周套餐。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 如果当前设备已开启云存套餐，Override=1会使用新套餐覆盖原有套餐。不传此参数则默认为0。
	Override *uint64 `json:"Override,omitempty" name:"Override"`
}

func (r *CreateCloudStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "PackageId")
	delete(f, "Override")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudStorageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudStorageResponseParams `json:"Response"`
}

func (r *CreateCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataForwardRequestParams struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 转发地址。如果有鉴权Token，则需要自行传入，例如 [{\"forward\":{\"api\":\"http://123.207.117.108:1080/sub.php\",\"token\":\"testtoken\"}}]
	ForwardAddr *string `json:"ForwardAddr,omitempty" name:"ForwardAddr"`

	// 1-数据信息转发 2-设备上下线状态转发 3-数据信息转发&设备上下线状态转发
	DataChose *int64 `json:"DataChose,omitempty" name:"DataChose"`
}

type CreateDataForwardRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 转发地址。如果有鉴权Token，则需要自行传入，例如 [{\"forward\":{\"api\":\"http://123.207.117.108:1080/sub.php\",\"token\":\"testtoken\"}}]
	ForwardAddr *string `json:"ForwardAddr,omitempty" name:"ForwardAddr"`

	// 1-数据信息转发 2-设备上下线状态转发 3-数据信息转发&设备上下线状态转发
	DataChose *int64 `json:"DataChose,omitempty" name:"DataChose"`
}

func (r *CreateDataForwardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataForwardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ForwardAddr")
	delete(f, "DataChose")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataForwardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataForwardResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDataForwardResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataForwardResponseParams `json:"Response"`
}

func (r *CreateDataForwardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataForwardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateForwardRuleRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 消息类型
	MsgType *uint64 `json:"MsgType,omitempty" name:"MsgType"`

	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 队列区域
	QueueRegion *string `json:"QueueRegion,omitempty" name:"QueueRegion"`

	// 队列类型 0.CMQ  1.Ckafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 临时密钥
	Consecretid *string `json:"Consecretid,omitempty" name:"Consecretid"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 队列或主题ID
	QueueID *string `json:"QueueID,omitempty" name:"QueueID"`

	// 队列或主题名称
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type CreateForwardRuleRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 消息类型
	MsgType *uint64 `json:"MsgType,omitempty" name:"MsgType"`

	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 队列区域
	QueueRegion *string `json:"QueueRegion,omitempty" name:"QueueRegion"`

	// 队列类型 0.CMQ  1.Ckafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 临时密钥
	Consecretid *string `json:"Consecretid,omitempty" name:"Consecretid"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 队列或主题ID
	QueueID *string `json:"QueueID,omitempty" name:"QueueID"`

	// 队列或主题名称
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *CreateForwardRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateForwardRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "MsgType")
	delete(f, "Skey")
	delete(f, "QueueRegion")
	delete(f, "QueueType")
	delete(f, "Consecretid")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "QueueID")
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateForwardRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateForwardRuleResponseParams struct {
	// 腾讯云账号
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 队列名
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 消息类型
	MsgType *uint64 `json:"MsgType,omitempty" name:"MsgType"`

	// 结果
	Result *uint64 `json:"Result,omitempty" name:"Result"`

	// 角色名称
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 角色ID
	RoleID *uint64 `json:"RoleID,omitempty" name:"RoleID"`

	// 队列区
	QueueRegion *string `json:"QueueRegion,omitempty" name:"QueueRegion"`

	// 消息队列的类型。 0：CMQ，1：CKafaka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 实例id， 目前只有Ckafaka会用到
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称，目前只有Ckafaka会用到
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 错误消息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateForwardRuleResponseParams `json:"Response"`
}

func (r *CreateForwardRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateForwardRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProductRequestParams struct {
	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品设备类型 1.普通设备 2.NVR设备
	DeviceType *uint64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 产品有效期
	ProductVaildYears *uint64 `json:"ProductVaildYears,omitempty" name:"ProductVaildYears"`

	// 设备功能码 ypsxth音频双向通话 spdxth视频单向通话
	Features []*string `json:"Features,omitempty" name:"Features"`

	// 设备操作系统，通用设备填default
	ChipOs *string `json:"ChipOs,omitempty" name:"ChipOs"`

	// 芯片厂商id，通用设备填default
	ChipManufactureId *string `json:"ChipManufactureId,omitempty" name:"ChipManufactureId"`

	// 芯片id，通用设备填default
	ChipId *string `json:"ChipId,omitempty" name:"ChipId"`

	// 产品描述信息
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 认证方式 只支持取值为2 psk认证
	EncryptionType *uint64 `json:"EncryptionType,omitempty" name:"EncryptionType"`

	// 连接类型，wifi表示WIFI连接，cellular表示4G连接
	NetType *string `json:"NetType,omitempty" name:"NetType"`
}

type CreateProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品设备类型 1.普通设备 2.NVR设备
	DeviceType *uint64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 产品有效期
	ProductVaildYears *uint64 `json:"ProductVaildYears,omitempty" name:"ProductVaildYears"`

	// 设备功能码 ypsxth音频双向通话 spdxth视频单向通话
	Features []*string `json:"Features,omitempty" name:"Features"`

	// 设备操作系统，通用设备填default
	ChipOs *string `json:"ChipOs,omitempty" name:"ChipOs"`

	// 芯片厂商id，通用设备填default
	ChipManufactureId *string `json:"ChipManufactureId,omitempty" name:"ChipManufactureId"`

	// 芯片id，通用设备填default
	ChipId *string `json:"ChipId,omitempty" name:"ChipId"`

	// 产品描述信息
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 认证方式 只支持取值为2 psk认证
	EncryptionType *uint64 `json:"EncryptionType,omitempty" name:"EncryptionType"`

	// 连接类型，wifi表示WIFI连接，cellular表示4G连接
	NetType *string `json:"NetType,omitempty" name:"NetType"`
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
	delete(f, "DeviceType")
	delete(f, "ProductVaildYears")
	delete(f, "Features")
	delete(f, "ChipOs")
	delete(f, "ChipManufactureId")
	delete(f, "ChipId")
	delete(f, "ProductDescription")
	delete(f, "EncryptionType")
	delete(f, "NetType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProductResponseParams struct {
	// 产品详情
	Data *VideoProduct `json:"Data,omitempty" name:"Data"`

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

type DataForward struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 转发地址。
	ForwardAddr *string `json:"ForwardAddr,omitempty" name:"ForwardAddr"`

	// 转发状态。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 1-数据信息转发 2-设备上下线状态转发 3-数据信息转发&设备上下线状态转发
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataChose *int64 `json:"DataChose,omitempty" name:"DataChose"`
}

// Predefined struct for user
type DeleteDeviceRequestParams struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceRequest has unknown keys!", "")
	}
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
type DeleteFirmwareRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`
}

type DeleteFirmwareRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`
}

func (r *DeleteFirmwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFirmwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFirmwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFirmwareResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFirmwareResponseParams `json:"Response"`
}

func (r *DeleteFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFirmwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteForwardRuleRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 队列类型
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 队列名称
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type DeleteForwardRuleRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 队列类型
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 队列名称
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *DeleteForwardRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteForwardRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "Skey")
	delete(f, "QueueType")
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteForwardRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteForwardRuleResponseParams struct {
	// 腾讯云账号
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 队列名称
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 删除结果 0成功 其他不成功
	Result *uint64 `json:"Result,omitempty" name:"Result"`

	// 错误消息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteForwardRuleResponseParams `json:"Response"`
}

func (r *DeleteForwardRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteForwardRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProductRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DeleteProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
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
type DescribeAIModelApplicationsRequestParams struct {
	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 分页的大小，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribeAIModelApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 分页的大小，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeAIModelApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIModelApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIModelApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIModelApplicationsResponseParams struct {
	// 申请记录数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 申请记录数组
	Applications []*AIModelApplication `json:"Applications,omitempty" name:"Applications"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAIModelApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIModelApplicationsResponseParams `json:"Response"`
}

func (r *DescribeAIModelApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIModelApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIModelChannelRequestParams struct {
	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribeAIModelChannelRequest struct {
	*tchttp.BaseRequest
	
	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeAIModelChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIModelChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIModelChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIModelChannelResponseParams struct {
	// 推送类型。ckafka：消息队列；forward：http/https推送
	Type *string `json:"Type,omitempty" name:"Type"`

	// 第三方推送地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardAddress *string `json:"ForwardAddress,omitempty" name:"ForwardAddress"`

	// 第三方推送密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardKey *string `json:"ForwardKey,omitempty" name:"ForwardKey"`

	// ckafka地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	CKafkaRegion *string `json:"CKafkaRegion,omitempty" name:"CKafkaRegion"`

	// ckafka实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	CKafkaInstance *string `json:"CKafkaInstance,omitempty" name:"CKafkaInstance"`

	// ckafka订阅主题
	// 注意：此字段可能返回 null，表示取不到有效值。
	CKafkaTopic *string `json:"CKafkaTopic,omitempty" name:"CKafkaTopic"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAIModelChannelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIModelChannelResponseParams `json:"Response"`
}

func (r *DescribeAIModelChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIModelChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIModelUsageRequestParams struct {
	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 偏移量，从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAIModelUsageRequest struct {
	*tchttp.BaseRequest
	
	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 偏移量，从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAIModelUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIModelUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "ProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIModelUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIModelUsageResponseParams struct {
	// AI模型资源包总量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// AI模型资源包信息数组
	UsageInfo []*AIModelUsageInfo `json:"UsageInfo,omitempty" name:"UsageInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAIModelUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIModelUsageResponseParams `json:"Response"`
}

func (r *DescribeAIModelUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIModelUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIModelsRequestParams struct {
	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 申请状态：1-已申请；2-已取消；3-已拒绝；4-已通过
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAIModelsRequest struct {
	*tchttp.BaseRequest
	
	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 申请状态：1-已申请；2-已取消；3-已拒绝；4-已通过
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAIModelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIModelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIModelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIModelsResponseParams struct {
	// AI模型数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// AI模型信息数组
	Models []*AIModelInfo `json:"Models,omitempty" name:"Models"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAIModelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIModelsResponseParams `json:"Response"`
}

func (r *DescribeAIModelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBalanceRequestParams struct {
	// 账户类型：1-设备接入；2-云存。
	AccountType *uint64 `json:"AccountType,omitempty" name:"AccountType"`
}

type DescribeBalanceRequest struct {
	*tchttp.BaseRequest
	
	// 账户类型：1-设备接入；2-云存。
	AccountType *uint64 `json:"AccountType,omitempty" name:"AccountType"`
}

func (r *DescribeBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBalanceResponseParams struct {
	// 账户余额，单位：分（人民币）。
	Balance *uint64 `json:"Balance,omitempty" name:"Balance"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBalanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBalanceResponseParams `json:"Response"`
}

func (r *DescribeBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBalanceTransactionsRequestParams struct {
	// 账户类型：1-设备接入；2-云存。
	AccountType *uint64 `json:"AccountType,omitempty" name:"AccountType"`

	// 分页游标开始，默认为0开始拉取第一条。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页每页数量。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 流水类型：All-全部类型；Recharge-充值；CreateOrder-新购。默认为All
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

type DescribeBalanceTransactionsRequest struct {
	*tchttp.BaseRequest
	
	// 账户类型：1-设备接入；2-云存。
	AccountType *uint64 `json:"AccountType,omitempty" name:"AccountType"`

	// 分页游标开始，默认为0开始拉取第一条。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页每页数量。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 流水类型：All-全部类型；Recharge-充值；CreateOrder-新购。默认为All
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *DescribeBalanceTransactionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBalanceTransactionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Operation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBalanceTransactionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBalanceTransactionsResponseParams struct {
	// 账户流水总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 账户流水详情数组。
	Transactions []*BalanceTransaction `json:"Transactions,omitempty" name:"Transactions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBalanceTransactionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBalanceTransactionsResponseParams `json:"Response"`
}

func (r *DescribeBalanceTransactionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBalanceTransactionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchRequestParams struct {
	// 批次ID
	BatchId *uint64 `json:"BatchId,omitempty" name:"BatchId"`
}

type DescribeBatchRequest struct {
	*tchttp.BaseRequest
	
	// 批次ID
	BatchId *uint64 `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *DescribeBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchResponseParams struct {
	// 批次详情
	Data *VideoBatch `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBatchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchResponseParams `json:"Response"`
}

func (r *DescribeBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchsRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 分页的大小，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeBatchsRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 分页的大小，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBatchsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchsResponseParams struct {
	// 批次数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 批次列表详情
	Data []*VideoBatch `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBatchsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchsResponseParams `json:"Response"`
}

func (r *DescribeBatchsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCategoryRequestParams struct {
	// Category ID。
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type DescribeCategoryRequest struct {
	*tchttp.BaseRequest
	
	// Category ID。
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCategoryResponseParams struct {
	// Category详情
	Data *ProductTemplate `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCategoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCategoryResponseParams `json:"Response"`
}

func (r *DescribeCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageDateRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeCloudStorageDateRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeCloudStorageDateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageDateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageDateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageDateResponseParams struct {
	// 云存日期数组，["2021-01-05","2021-01-06"]
	Data []*string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudStorageDateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageDateResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageDateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageDateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageEventsRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 起始时间（Unix 时间戳，秒级）, 为0 表示 当前时间 - 24h
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（Unix 时间戳，秒级）, 为0 表示当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 请求上下文, 用作查询游标
	Context *string `json:"Context,omitempty" name:"Context"`

	// 查询数据项目的最大数量, 默认为10。假设传Size=10，返回的实际事件数量为N，则 5 <= N <= 10。
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 通道ID 非NVR设备则不填 NVR设备则必填 默认为无
	ChannelId *uint64 `json:"ChannelId,omitempty" name:"ChannelId"`
}

type DescribeCloudStorageEventsRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 起始时间（Unix 时间戳，秒级）, 为0 表示 当前时间 - 24h
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（Unix 时间戳，秒级）, 为0 表示当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 请求上下文, 用作查询游标
	Context *string `json:"Context,omitempty" name:"Context"`

	// 查询数据项目的最大数量, 默认为10。假设传Size=10，返回的实际事件数量为N，则 5 <= N <= 10。
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 通道ID 非NVR设备则不填 NVR设备则必填 默认为无
	ChannelId *uint64 `json:"ChannelId,omitempty" name:"ChannelId"`
}

func (r *DescribeCloudStorageEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Context")
	delete(f, "Size")
	delete(f, "EventId")
	delete(f, "UserId")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageEventsResponseParams struct {
	// 云存事件列表
	Events []*CloudStorageEvent `json:"Events,omitempty" name:"Events"`

	// 请求上下文, 用作查询游标
	Context *string `json:"Context,omitempty" name:"Context"`

	// 拉取结果是否已经结束
	Listover *bool `json:"Listover,omitempty" name:"Listover"`

	// 内部结果数量，并不等同于事件总数。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 视频播放URL
	VideoURL *string `json:"VideoURL,omitempty" name:"VideoURL"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudStorageEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageEventsResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 云存用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeCloudStorageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 云存用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeCloudStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageResponseParams struct {
	// 云存开启状态，1为开启，0为未开启或已过期
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 云存类型，1为全时云存，2为事件云存
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 云存套餐过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 云存回看时长
	ShiftDuration *uint64 `json:"ShiftDuration,omitempty" name:"ShiftDuration"`

	// 云存用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageThumbnailRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 缩略图文件名
	Thumbnail *string `json:"Thumbnail,omitempty" name:"Thumbnail"`
}

type DescribeCloudStorageThumbnailRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 缩略图文件名
	Thumbnail *string `json:"Thumbnail,omitempty" name:"Thumbnail"`
}

func (r *DescribeCloudStorageThumbnailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageThumbnailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Thumbnail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageThumbnailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageThumbnailResponseParams struct {
	// 缩略图访问地址
	ThumbnailURL *string `json:"ThumbnailURL,omitempty" name:"ThumbnailURL"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudStorageThumbnailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageThumbnailResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageThumbnailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageThumbnailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageTimeRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 云存日期，例如"2020-01-05"
	Date *string `json:"Date,omitempty" name:"Date"`

	// 开始时间，unix时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，unix时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeCloudStorageTimeRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 云存日期，例如"2020-01-05"
	Date *string `json:"Date,omitempty" name:"Date"`

	// 开始时间，unix时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，unix时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeCloudStorageTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Date")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageTimeResponseParams struct {
	// 接口返回数据
	Data *CloudStorageTimeData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudStorageTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageTimeResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageUsersRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 分页拉取数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页拉取偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeCloudStorageUsersRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 分页拉取数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页拉取偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeCloudStorageUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageUsersResponseParams struct {
	// 用户总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 用户信息
	Users []*CloudStorageUserInfo `json:"Users,omitempty" name:"Users"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudStorageUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageUsersResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataForwardListRequestParams struct {
	// 产品ID列表
	ProductIds *string `json:"ProductIds,omitempty" name:"ProductIds"`
}

type DescribeDataForwardListRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID列表
	ProductIds *string `json:"ProductIds,omitempty" name:"ProductIds"`
}

func (r *DescribeDataForwardListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataForwardListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataForwardListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataForwardListResponseParams struct {
	// 数据转发列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataForwardList []*DataForward `json:"DataForwardList,omitempty" name:"DataForwardList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataForwardListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataForwardListResponseParams `json:"Response"`
}

func (r *DescribeDataForwardListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataForwardListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceActionHistoryRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 开始范围的 unix 毫秒时间戳
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 结束范围的 unix 毫秒时间戳
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 动作Id
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 查询条数 默认为0 最大不超过500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 游标，标识查询位置。
	Context *string `json:"Context,omitempty" name:"Context"`
}

type DescribeDeviceActionHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 开始范围的 unix 毫秒时间戳
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 结束范围的 unix 毫秒时间戳
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 动作Id
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 查询条数 默认为0 最大不超过500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 游标，标识查询位置。
	Context *string `json:"Context,omitempty" name:"Context"`
}

func (r *DescribeDeviceActionHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceActionHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "MinTime")
	delete(f, "MaxTime")
	delete(f, "ActionId")
	delete(f, "Limit")
	delete(f, "Context")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceActionHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceActionHistoryResponseParams struct {
	// 总条数
	TotalCounts *uint64 `json:"TotalCounts,omitempty" name:"TotalCounts"`

	// 动作历史
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionHistories []*ActionHistory `json:"ActionHistories,omitempty" name:"ActionHistories"`

	// 用于标识查询结果的上下文，翻页用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitempty" name:"Context"`

	// 搜索结果是否已经结束。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Listover *bool `json:"Listover,omitempty" name:"Listover"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceActionHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceActionHistoryResponseParams `json:"Response"`
}

func (r *DescribeDeviceActionHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceActionHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceCommLogRequestParams struct {
	// 开始时间 13位时间戳 单位毫秒
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 结束时间 13位时间戳 单位毫秒
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 返回条数 默认为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 类型：shadow 下行，device 上行 默认为空则全部查询
	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeDeviceCommLogRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间 13位时间戳 单位毫秒
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 结束时间 13位时间戳 单位毫秒
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 返回条数 默认为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 类型：shadow 下行，device 上行 默认为空则全部查询
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeDeviceCommLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceCommLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinTime")
	delete(f, "MaxTime")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceCommLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceCommLogResponseParams struct {
	// 数据是否已全部返回，true 表示数据全部返回，false 表示还有数据待返回，可将 Context 作为入参，继续查询返回结果。
	Listover *bool `json:"Listover,omitempty" name:"Listover"`

	// 检索上下文，当 ListOver 为false时，可以用此上下文，继续读取后续数据
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志数据结果数组，返回对应时间点及取值。
	Results []*DeviceCommLogItem `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceCommLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceCommLogResponseParams `json:"Response"`
}

func (r *DescribeDeviceCommLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceCommLogResponse) FromJsonString(s string) error {
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
	Limit []*uint64 `json:"Limit,omitempty" name:"Limit"`

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
	Limit []*uint64 `json:"Limit,omitempty" name:"Limit"`

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
}

type DescribeDeviceDataRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
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
type DescribeDeviceEventHistoryRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 搜索的事件类型：alert 表示告警，fault 表示故障，info 表示信息，为空则表示查询上述所有类型事件
	Type *string `json:"Type,omitempty" name:"Type"`

	// 起始时间（Unix 时间戳，秒级）, 为0 表示 当前时间 - 24h
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（Unix 时间戳，秒级）, 为0 表示当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 搜索上下文, 用作查询游标
	Context *string `json:"Context,omitempty" name:"Context"`

	// 单次获取的历史数据项目的最大数量, 缺省10
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type DescribeDeviceEventHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 搜索的事件类型：alert 表示告警，fault 表示故障，info 表示信息，为空则表示查询上述所有类型事件
	Type *string `json:"Type,omitempty" name:"Type"`

	// 起始时间（Unix 时间戳，秒级）, 为0 表示 当前时间 - 24h
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（Unix 时间戳，秒级）, 为0 表示当前时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 搜索上下文, 用作查询游标
	Context *string `json:"Context,omitempty" name:"Context"`

	// 单次获取的历史数据项目的最大数量, 缺省10
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeDeviceEventHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceEventHistoryRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceEventHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceEventHistoryResponseParams struct {
	// 搜索上下文, 用作查询游标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitempty" name:"Context"`

	// 搜索结果数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 搜索结果是否已经结束
	// 注意：此字段可能返回 null，表示取不到有效值。
	Listover *bool `json:"Listover,omitempty" name:"Listover"`

	// 搜集结果集
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventHistory []*EventHistoryItem `json:"EventHistory,omitempty" name:"EventHistory"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceEventHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceEventHistoryResponseParams `json:"Response"`
}

func (r *DescribeDeviceEventHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceEventHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DescribeDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

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
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceResponseParams struct {
	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备是否在线，0不在线，1在线，2获取失败，3未激活
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// 设备最后上线时间
	LoginTime *uint64 `json:"LoginTime,omitempty" name:"LoginTime"`

	// 设备密钥
	DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

	// 设备启用状态
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`

	// 设备过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 设备的sdk日志等级，0：关闭，1：错误，2：告警，3：信息，4：调试
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogLevel *uint64 `json:"LogLevel,omitempty" name:"LogLevel"`

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
type DescribeDeviceStatusLogRequestParams struct {
	// 开始时间（毫秒）
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 结束时间（毫秒）
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 返回条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`
}

type DescribeDeviceStatusLogRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间（毫秒）
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 结束时间（毫秒）
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 返回条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`
}

func (r *DescribeDeviceStatusLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStatusLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinTime")
	delete(f, "MaxTime")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Limit")
	delete(f, "Context")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceStatusLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceStatusLogResponseParams struct {
	// 数据是否已全部返回，true 表示数据全部返回，false 表示还有数据待返回，可将 Context 作为入参，继续查询返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Listover *bool `json:"Listover,omitempty" name:"Listover"`

	// 检索上下文，当 ListOver 为false时，可以用此上下文，继续读取后续数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志数据结果数组，返回对应时间点及取值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*DeviceStatusLogItem `json:"Results,omitempty" name:"Results"`

	// 日志数据结果总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceStatusLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceStatusLogResponseParams `json:"Response"`
}

func (r *DescribeDeviceStatusLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStatusLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesRequestParams struct {
	// 需要查看设备列表的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 需要过滤的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查看设备列表的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 需要过滤的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
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
	delete(f, "DeviceName")
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

	// 查询偏移量 默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的数量 默认为50
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

	// 查询偏移量 默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的数量 默认为50
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
type DescribeForwardRuleRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 队列类型，0：CMQ，1：Ckafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 临时密钥
	Consecretid *string `json:"Consecretid,omitempty" name:"Consecretid"`
}

type DescribeForwardRuleRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 队列类型，0：CMQ，1：Ckafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 临时密钥
	Consecretid *string `json:"Consecretid,omitempty" name:"Consecretid"`
}

func (r *DescribeForwardRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForwardRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "Skey")
	delete(f, "QueueType")
	delete(f, "Consecretid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeForwardRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForwardRuleResponseParams struct {
	// 腾讯云账号
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 队列名称
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 消息类型 1设备上报信息 2设备状态变化通知 3为全选
	MsgType *uint64 `json:"MsgType,omitempty" name:"MsgType"`

	// 结果 2表示禁用 其他为成功
	Result *uint64 `json:"Result,omitempty" name:"Result"`

	// 角色名
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 角色ID
	RoleID *uint64 `json:"RoleID,omitempty" name:"RoleID"`

	// 队列区域
	QueueRegion *string `json:"QueueRegion,omitempty" name:"QueueRegion"`

	// 队列类型，0：CMQ，1：Ckafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 实例id， 目前只有Ckafaka会用到
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称，目前只有Ckafaka会用到
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 错误消息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeForwardRuleResponseParams `json:"Response"`
}

func (r *DescribeForwardRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForwardRuleResponse) FromJsonString(s string) error {
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
type DescribeProductDynamicRegisterRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribeProductDynamicRegisterRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeProductDynamicRegisterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductDynamicRegisterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductDynamicRegisterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductDynamicRegisterResponseParams struct {
	// 动态注册类型，0-关闭 1-预创建设备 2-自动创建设备
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// 动态注册产品密钥
	ProductSecret *string `json:"ProductSecret,omitempty" name:"ProductSecret"`

	// 动态注册设备上限
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductDynamicRegisterResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductDynamicRegisterResponseParams `json:"Response"`
}

func (r *DescribeProductDynamicRegisterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductDynamicRegisterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductRequestParams struct {
	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribeProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品id
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
type DescribeProductResponseParams struct {
	// 产品详情
	Data *VideoProduct `json:"Data,omitempty" name:"Data"`

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
type DescribeProductsRequestParams struct {
	// 分页的大小，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeProductsRequest struct {
	*tchttp.BaseRequest
	
	// 分页的大小，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
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
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductsResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 产品详情列表
	Data []*VideoProduct `json:"Data,omitempty" name:"Data"`

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
type DescribeSDKLogRequestParams struct {
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

type DescribeSDKLogRequest struct {
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

func (r *DescribeSDKLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSDKLogRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSDKLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSDKLogResponseParams struct {
	// 日志检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 是否还有日志，如有仍有日志，下次查询的请求带上当前请求返回的Context
	Listover *bool `json:"Listover,omitempty" name:"Listover"`

	// 日志列表
	Results []*SDKLogItem `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSDKLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSDKLogResponseParams `json:"Response"`
}

func (r *DescribeSDKLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSDKLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceCommLogItem struct {
	// 时间
	Time *string `json:"Time,omitempty" name:"Time"`

	// 日志类型，device 设备上行，shadow 服务端下行。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 通讯数据。
	Data *string `json:"Data,omitempty" name:"Data"`
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

	// 设备是否在线，0不在线，1在线，2获取失败，3未激活
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// 设备最后上线时间
	LoginTime *uint64 `json:"LoginTime,omitempty" name:"LoginTime"`

	// 设备密钥
	DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

	// 设备启用状态 0为停用 1为可用
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`

	// 设备过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type DeviceStatusLogItem struct {
	// 时间
	Time *string `json:"Time,omitempty" name:"Time"`

	// 状态类型： Online 上线，Offline 下线
	Type *string `json:"Type,omitempty" name:"Type"`

	// 日志信息
	Data *string `json:"Data,omitempty" name:"Data"`
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
type EditFirmwareRequestParams struct {
	// 产品ID。
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号。
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件名称。
	FirmwareName *string `json:"FirmwareName,omitempty" name:"FirmwareName"`

	// 固件描述。
	FirmwareDescription *string `json:"FirmwareDescription,omitempty" name:"FirmwareDescription"`
}

type EditFirmwareRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID。
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本号。
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 固件名称。
	FirmwareName *string `json:"FirmwareName,omitempty" name:"FirmwareName"`

	// 固件描述。
	FirmwareDescription *string `json:"FirmwareDescription,omitempty" name:"FirmwareDescription"`
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

type FirmwareInfo struct {
	// 固件版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 固件MD5值
	Md5sum *string `json:"Md5sum,omitempty" name:"Md5sum"`

	// 固件创建时间
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 固件名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 固件描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
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
type GenerateSignedVideoURLRequestParams struct {
	// 视频播放原始URL地址
	VideoURL *string `json:"VideoURL,omitempty" name:"VideoURL"`

	// 播放链接过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type GenerateSignedVideoURLRequest struct {
	*tchttp.BaseRequest
	
	// 视频播放原始URL地址
	VideoURL *string `json:"VideoURL,omitempty" name:"VideoURL"`

	// 播放链接过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *GenerateSignedVideoURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateSignedVideoURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VideoURL")
	delete(f, "ExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateSignedVideoURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateSignedVideoURLResponseParams struct {
	// 视频防盗链播放URL
	SignedVideoURL *string `json:"SignedVideoURL,omitempty" name:"SignedVideoURL"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateSignedVideoURLResponse struct {
	*tchttp.BaseResponse
	Response *GenerateSignedVideoURLResponseParams `json:"Response"`
}

func (r *GenerateSignedVideoURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateSignedVideoURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAllFirmwareVersionRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
}

type GetAllFirmwareVersionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
}

func (r *GetAllFirmwareVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAllFirmwareVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAllFirmwareVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAllFirmwareVersionResponseParams struct {
	// 固件可用版本列表
	Version []*string `json:"Version,omitempty" name:"Version"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetAllFirmwareVersionResponse struct {
	*tchttp.BaseResponse
	Response *GetAllFirmwareVersionResponseParams `json:"Response"`
}

func (r *GetAllFirmwareVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAllFirmwareVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFirmwareURLRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`
}

type GetFirmwareURLRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`
}

func (r *GetFirmwareURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFirmwareURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFirmwareURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFirmwareURLResponseParams struct {
	// 固件URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetFirmwareURLResponse struct {
	*tchttp.BaseResponse
	Response *GetFirmwareURLResponseParams `json:"Response"`
}

func (r *GetFirmwareURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFirmwareURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportModelDefinitionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 数据模板定义
	ModelSchema *string `json:"ModelSchema,omitempty" name:"ModelSchema"`
}

type ImportModelDefinitionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 数据模板定义
	ModelSchema *string `json:"ModelSchema,omitempty" name:"ModelSchema"`
}

func (r *ImportModelDefinitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportModelDefinitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ModelSchema")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportModelDefinitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportModelDefinitionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportModelDefinitionResponse struct {
	*tchttp.BaseResponse
	Response *ImportModelDefinitionResponseParams `json:"Response"`
}

func (r *ImportModelDefinitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportModelDefinitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InheritCloudStorageUserRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 原始用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 目标用户ID
	ToUserId *string `json:"ToUserId,omitempty" name:"ToUserId"`
}

type InheritCloudStorageUserRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 原始用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 目标用户ID
	ToUserId *string `json:"ToUserId,omitempty" name:"ToUserId"`
}

func (r *InheritCloudStorageUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InheritCloudStorageUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "UserId")
	delete(f, "ToUserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InheritCloudStorageUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InheritCloudStorageUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InheritCloudStorageUserResponse struct {
	*tchttp.BaseResponse
	Response *InheritCloudStorageUserResponseParams `json:"Response"`
}

func (r *InheritCloudStorageUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InheritCloudStorageUserResponse) FromJsonString(s string) error {
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
type ModifyDataForwardRequestParams struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 转发地址。如果有鉴权Token，则需要自行传入，例如 [{\"forward\":{\"api\":\"http://123.207.117.108:1080/sub.php\",\"token\":\"testtoken\"}}]
	ForwardAddr *string `json:"ForwardAddr,omitempty" name:"ForwardAddr"`

	// 1-数据信息转发 2-设备上下线状态转发 3-数据信息转发&设备上下线状态转发
	DataChose *int64 `json:"DataChose,omitempty" name:"DataChose"`
}

type ModifyDataForwardRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 转发地址。如果有鉴权Token，则需要自行传入，例如 [{\"forward\":{\"api\":\"http://123.207.117.108:1080/sub.php\",\"token\":\"testtoken\"}}]
	ForwardAddr *string `json:"ForwardAddr,omitempty" name:"ForwardAddr"`

	// 1-数据信息转发 2-设备上下线状态转发 3-数据信息转发&设备上下线状态转发
	DataChose *int64 `json:"DataChose,omitempty" name:"DataChose"`
}

func (r *ModifyDataForwardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataForwardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ForwardAddr")
	delete(f, "DataChose")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDataForwardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataForwardResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDataForwardResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDataForwardResponseParams `json:"Response"`
}

func (r *ModifyDataForwardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataForwardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataForwardStatusRequestParams struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 转发状态，1启用，0禁用。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ModifyDataForwardStatusRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 转发状态，1启用，0禁用。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDataForwardStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataForwardStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDataForwardStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataForwardStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDataForwardStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDataForwardStatusResponseParams `json:"Response"`
}

func (r *ModifyDataForwardStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataForwardStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceLogLevelRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 日志级别，0：关闭，1：错误，2：告警，3：信息，4：调试
	LogLevel *uint64 `json:"LogLevel,omitempty" name:"LogLevel"`
}

type ModifyDeviceLogLevelRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 日志级别，0：关闭，1：错误，2：告警，3：信息，4：调试
	LogLevel *uint64 `json:"LogLevel,omitempty" name:"LogLevel"`
}

func (r *ModifyDeviceLogLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceLogLevelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "LogLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceLogLevelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceLogLevelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDeviceLogLevelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceLogLevelResponseParams `json:"Response"`
}

func (r *ModifyDeviceLogLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceLogLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceRequestParams struct {
	// 设备所属产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 要设置的设备状态，1为启用，0为禁用
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`
}

type ModifyDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备所属产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 要设置的设备状态，1为启用，0为禁用
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`
}

func (r *ModifyDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "EnableState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceResponseParams `json:"Response"`
}

func (r *ModifyDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyForwardRuleRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 消息类型
	MsgType *uint64 `json:"MsgType,omitempty" name:"MsgType"`

	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 队列区域
	QueueRegion *string `json:"QueueRegion,omitempty" name:"QueueRegion"`

	// 队列类型 0.CMQ 1.CKafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 临时密钥
	Consecretid *string `json:"Consecretid,omitempty" name:"Consecretid"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 队列或主题ID
	QueueID *string `json:"QueueID,omitempty" name:"QueueID"`

	// 队列或主题名称
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type ModifyForwardRuleRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 消息类型
	MsgType *uint64 `json:"MsgType,omitempty" name:"MsgType"`

	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 队列区域
	QueueRegion *string `json:"QueueRegion,omitempty" name:"QueueRegion"`

	// 队列类型 0.CMQ 1.CKafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 临时密钥
	Consecretid *string `json:"Consecretid,omitempty" name:"Consecretid"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 队列或主题ID
	QueueID *string `json:"QueueID,omitempty" name:"QueueID"`

	// 队列或主题名称
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *ModifyForwardRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyForwardRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "MsgType")
	delete(f, "Skey")
	delete(f, "QueueRegion")
	delete(f, "QueueType")
	delete(f, "Consecretid")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "QueueID")
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyForwardRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyForwardRuleResponseParams struct {
	// 腾讯云账号
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 结果
	Result *uint64 `json:"Result,omitempty" name:"Result"`

	// 错误信息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 队列类型 0.CMQ 1.CKafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyForwardRuleResponseParams `json:"Response"`
}

func (r *ModifyForwardRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyForwardRuleResponse) FromJsonString(s string) error {
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
type ModifyProductDynamicRegisterRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 动态注册类型，0-关闭 1-预创建设备 2-自动创建设备
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// 动态注册设备上限
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`
}

type ModifyProductDynamicRegisterRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 动态注册类型，0-关闭 1-预创建设备 2-自动创建设备
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// 动态注册设备上限
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`
}

func (r *ModifyProductDynamicRegisterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProductDynamicRegisterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "RegisterType")
	delete(f, "RegisterLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProductDynamicRegisterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProductDynamicRegisterResponseParams struct {
	// 动态注册类型，0-关闭 1-预创建设备 2-自动创建设备
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// 动态注册产品密钥
	ProductSecret *string `json:"ProductSecret,omitempty" name:"ProductSecret"`

	// 动态注册设备上限
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyProductDynamicRegisterResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProductDynamicRegisterResponseParams `json:"Response"`
}

func (r *ModifyProductDynamicRegisterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProductDynamicRegisterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProductRequestParams struct {
	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 修改的产品名称 （支持中文、英文、数字、下划线组合，最多不超过20个字符）
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 修改的产品描述 （最多不超过128个字符）
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`
}

type ModifyProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 修改的产品名称 （支持中文、英文、数字、下划线组合，最多不超过20个字符）
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 修改的产品描述 （最多不超过128个字符）
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`
}

func (r *ModifyProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ProductName")
	delete(f, "ProductDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProductResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyProductResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProductResponseParams `json:"Response"`
}

func (r *ModifyProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type ProductTemplate struct {
	// 实体ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 分类字段
	CategoryKey *string `json:"CategoryKey,omitempty" name:"CategoryKey"`

	// 分类名称
	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`

	// 上层实体ID
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// 物模型
	ModelTemplate *string `json:"ModelTemplate,omitempty" name:"ModelTemplate"`

	// 排列顺序
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListOrder *int64 `json:"ListOrder,omitempty" name:"ListOrder"`

	// 分类图标地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IconUrl *string `json:"IconUrl,omitempty" name:"IconUrl"`

	// 九宫格图片地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IconUrlGrid *string `json:"IconUrlGrid,omitempty" name:"IconUrlGrid"`
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
type ReportAliveDeviceRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type ReportAliveDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *ReportAliveDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportAliveDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportAliveDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportAliveDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReportAliveDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ReportAliveDeviceResponseParams `json:"Response"`
}

func (r *ReportAliveDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportAliveDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetCloudStorageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 通道ID 非NVR设备则不填 NVR设备则必填 默认为无
	ChannelId *uint64 `json:"ChannelId,omitempty" name:"ChannelId"`

	// 云存用户Id，为空则为默认云存空间。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type ResetCloudStorageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 通道ID 非NVR设备则不填 NVR设备则必填 默认为无
	ChannelId *uint64 `json:"ChannelId,omitempty" name:"ChannelId"`

	// 云存用户Id，为空则为默认云存空间。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *ResetCloudStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetCloudStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ChannelId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetCloudStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetCloudStorageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *ResetCloudStorageResponseParams `json:"Response"`
}

func (r *ResetCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetCloudStorageResponse) FromJsonString(s string) error {
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
type SetForwardAuthRequestParams struct {
	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 消息队列类型  0.CMQ 1.CKafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`
}

type SetForwardAuthRequest struct {
	*tchttp.BaseRequest
	
	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 消息队列类型  0.CMQ 1.CKafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`
}

func (r *SetForwardAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetForwardAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Skey")
	delete(f, "QueueType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetForwardAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetForwardAuthResponseParams struct {
	// 腾讯云账号
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 结果
	Result *uint64 `json:"Result,omitempty" name:"Result"`

	// 角色名
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 角色ID
	RoleID *uint64 `json:"RoleID,omitempty" name:"RoleID"`

	// 消息队列类型  0.CMQ 1.CKafka
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

	// 错误消息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetForwardAuthResponse struct {
	*tchttp.BaseResponse
	Response *SetForwardAuthResponseParams `json:"Response"`
}

func (r *SetForwardAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetForwardAuthResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type TransferCloudStorageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 已开通云存的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 未开通云存的设备名称
	ToDeviceName *string `json:"ToDeviceName,omitempty" name:"ToDeviceName"`
}

type TransferCloudStorageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 已开通云存的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 未开通云存的设备名称
	ToDeviceName *string `json:"ToDeviceName,omitempty" name:"ToDeviceName"`
}

func (r *TransferCloudStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferCloudStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ToDeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransferCloudStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferCloudStorageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TransferCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *TransferCloudStorageResponseParams `json:"Response"`
}

func (r *TransferCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferCloudStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAIModelChannelRequestParams struct {
	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 推送类型。ckafka：消息队列；forward：http/https推送
	Type *string `json:"Type,omitempty" name:"Type"`

	// 第三方推送地址
	ForwardAddress *string `json:"ForwardAddress,omitempty" name:"ForwardAddress"`

	// 第三方推送密钥，不填写则腾讯云自动生成。
	ForwardKey *string `json:"ForwardKey,omitempty" name:"ForwardKey"`

	// ckafka地域
	CKafkaRegion *string `json:"CKafkaRegion,omitempty" name:"CKafkaRegion"`

	// ckafka实例
	CKafkaInstance *string `json:"CKafkaInstance,omitempty" name:"CKafkaInstance"`

	// ckafka订阅主题
	CKafkaTopic *string `json:"CKafkaTopic,omitempty" name:"CKafkaTopic"`
}

type UpdateAIModelChannelRequest struct {
	*tchttp.BaseRequest
	
	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 推送类型。ckafka：消息队列；forward：http/https推送
	Type *string `json:"Type,omitempty" name:"Type"`

	// 第三方推送地址
	ForwardAddress *string `json:"ForwardAddress,omitempty" name:"ForwardAddress"`

	// 第三方推送密钥，不填写则腾讯云自动生成。
	ForwardKey *string `json:"ForwardKey,omitempty" name:"ForwardKey"`

	// ckafka地域
	CKafkaRegion *string `json:"CKafkaRegion,omitempty" name:"CKafkaRegion"`

	// ckafka实例
	CKafkaInstance *string `json:"CKafkaInstance,omitempty" name:"CKafkaInstance"`

	// ckafka订阅主题
	CKafkaTopic *string `json:"CKafkaTopic,omitempty" name:"CKafkaTopic"`
}

func (r *UpdateAIModelChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAIModelChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "ProductId")
	delete(f, "Type")
	delete(f, "ForwardAddress")
	delete(f, "ForwardKey")
	delete(f, "CKafkaRegion")
	delete(f, "CKafkaInstance")
	delete(f, "CKafkaTopic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAIModelChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAIModelChannelResponseParams struct {
	// 第三方推送密钥，如果选择自动生成则会返回此字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardKey *string `json:"ForwardKey,omitempty" name:"ForwardKey"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateAIModelChannelResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAIModelChannelResponseParams `json:"Response"`
}

func (r *UpdateAIModelChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAIModelChannelResponse) FromJsonString(s string) error {
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

type VideoBatch struct {
	// 批次ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 状态：1：待创建设备 2：创建中 3：已完成
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 设备前缀
	DevPre *string `json:"DevPre,omitempty" name:"DevPre"`

	// 设备数量
	DevNum *uint64 `json:"DevNum,omitempty" name:"DevNum"`

	// 已创建设备数量
	DevNumCreated *uint64 `json:"DevNumCreated,omitempty" name:"DevNumCreated"`

	// 批次下载地址
	BatchURL *string `json:"BatchURL,omitempty" name:"BatchURL"`

	// 创建时间。unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间。unix时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type VideoProduct struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品设备类型（普通设备)	1.普通设备
	DeviceType *uint64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 认证方式：2：PSK
	EncryptionType *uint64 `json:"EncryptionType,omitempty" name:"EncryptionType"`

	// 设备功能码
	Features []*string `json:"Features,omitempty" name:"Features"`

	// 操作系统
	ChipOs *string `json:"ChipOs,omitempty" name:"ChipOs"`

	// 芯片厂商id
	ChipManufactureId *string `json:"ChipManufactureId,omitempty" name:"ChipManufactureId"`

	// 芯片id
	ChipId *string `json:"ChipId,omitempty" name:"ChipId"`

	// 产品描述信息
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 创建时间unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间unix时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 连接类型，wifi表示WIFI连接，cellular表示4G连接
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetType *string `json:"NetType,omitempty" name:"NetType"`
}

// Predefined struct for user
type WakeUpDeviceRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

type WakeUpDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *WakeUpDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WakeUpDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "WakeUpDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type WakeUpDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type WakeUpDeviceResponse struct {
	*tchttp.BaseResponse
	Response *WakeUpDeviceResponseParams `json:"Response"`
}

func (r *WakeUpDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WakeUpDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}