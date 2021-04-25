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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

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
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames" list`
}

func (r *BatchUpdateFirmwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchUpdateFirmwareRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchUpdateFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchUpdateFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchUpdateFirmwareResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *CancelDeviceFirmwareTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelDeviceFirmwareTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelDeviceFirmwareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelDeviceFirmwareTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckForwardAuthRequest struct {
	*tchttp.BaseRequest

	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 队列类型
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`
}

func (r *CheckForwardAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckForwardAuthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckForwardAuthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 腾讯云账号
		Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

		// 结果
		Result *uint64 `json:"Result,omitempty" name:"Result"`

		// 产品ID
		Productid *string `json:"Productid,omitempty" name:"Productid"`

		// 错误消息
		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

		// 队列类型
		QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckForwardAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	TimeList []*CloudStorageTimeInfo `json:"TimeList,omitempty" name:"TimeList" list`

	// 播放地址
	VideoURL *string `json:"VideoURL,omitempty" name:"VideoURL"`
}

type CloudStorageTimeInfo struct {

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
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

func (r *CreateBatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批次ID
		BatchId *uint64 `json:"BatchId,omitempty" name:"BatchId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
}

func (r *CreateCloudStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCloudStorageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCloudStorageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 队列类型
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

func (r *CreateForwardRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *CreateForwardRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateForwardRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProductRequest struct {
	*tchttp.BaseRequest

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品设备类型
	DeviceType *uint64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 产品有效期
	ProductVaildYears *uint64 `json:"ProductVaildYears,omitempty" name:"ProductVaildYears"`

	// 设备功能码
	Features []*string `json:"Features,omitempty" name:"Features" list`

	// 设备操作系统
	ChipOs *string `json:"ChipOs,omitempty" name:"ChipOs"`

	// 芯片厂商id
	ChipManufactureId *string `json:"ChipManufactureId,omitempty" name:"ChipManufactureId"`

	// 芯片id
	ChipId *string `json:"ChipId,omitempty" name:"ChipId"`

	// 产品描述信息
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 认证方式。2 PSK
	EncryptionType *uint64 `json:"EncryptionType,omitempty" name:"EncryptionType"`
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

		// 产品详情
		Data *VideoProduct `json:"Data,omitempty" name:"Data"`

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

type CreateTaskFileUrlRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *CreateTaskFileUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTaskFileUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTaskFileUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务文件上传链接
		Url *string `json:"Url,omitempty" name:"Url"`

		// 任务文件名
		FileName *string `json:"FileName,omitempty" name:"FileName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTaskFileUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTaskFileUrlResponse) FromJsonString(s string) error {
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

func (r *DeleteFirmwareRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFirmwareResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteForwardRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 腾讯云账号
		Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

		// 队列名称
		QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

		// 产品ID
		ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

		// 删除结果
		Result *uint64 `json:"Result,omitempty" name:"Result"`

		// 错误消息
		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteForwardRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteForwardRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type DescribeBatchRequest struct {
	*tchttp.BaseRequest

	// 批次ID
	BatchId *uint64 `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *DescribeBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批次详情
		Data *VideoBatch `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeBatchsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批次数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 批次列表详情
		Data []*VideoBatch `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBatchsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBatchsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeCategoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCategoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Category详情
		Data *ProductTemplate `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCategoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudStorageDateRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DescribeCloudStorageDateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudStorageDateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudStorageDateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云存日期数组，["2021-01-05","2021-01-06"]
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudStorageDateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudStorageDateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 单次获取的历史数据项目的最大数量, 缺省10
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeCloudStorageEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudStorageEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudStorageEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云存事件列表
		Events []*CloudStorageEvent `json:"Events,omitempty" name:"Events" list`

		// 请求上下文, 用作查询游标
		Context *string `json:"Context,omitempty" name:"Context"`

		// 拉取结果是否已经结束
		Listover *bool `json:"Listover,omitempty" name:"Listover"`

		// 拉取结果数量
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 视频播放URL
		VideoURL *string `json:"VideoURL,omitempty" name:"VideoURL"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudStorageEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudStorageEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudStorageRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DescribeCloudStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudStorageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云存开启状态，1为开启，0为未开启或已过期
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 云存类型，1为全时云存，2为事件云存
		Type *uint64 `json:"Type,omitempty" name:"Type"`

		// 云存套餐过期时间
		ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

		// 云存回看时长
		ShiftDuration *uint64 `json:"ShiftDuration,omitempty" name:"ShiftDuration"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudStorageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeCloudStorageThumbnailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudStorageThumbnailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 缩略图访问地址
		ThumbnailURL *string `json:"ThumbnailURL,omitempty" name:"ThumbnailURL"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudStorageThumbnailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudStorageThumbnailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudStorageTimeRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 云存日期，例如"2020-01-05"
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DescribeCloudStorageTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudStorageTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudStorageTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接口返回数据
		Data *CloudStorageTimeData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudStorageTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudStorageTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 查询条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 游标，标识查询位置。
	Context *string `json:"Context,omitempty" name:"Context"`
}

func (r *DescribeDeviceActionHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceActionHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceActionHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		TotalCounts *uint64 `json:"TotalCounts,omitempty" name:"TotalCounts"`

		// 动作历史
	// 注意：此字段可能返回 null，表示取不到有效值。
		ActionHistories []*ActionHistory `json:"ActionHistories,omitempty" name:"ActionHistories" list`

		// 用于标识查询结果的上下文，翻页用。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Context *string `json:"Context,omitempty" name:"Context"`

		// 搜索结果是否已经结束。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Listover *bool `json:"Listover,omitempty" name:"Listover"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceActionHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceActionHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceCommLogRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 结束时间
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 返回条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 类型：shadow 下行，device 上行
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeDeviceCommLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceCommLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceCommLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否已全部返回，true 表示数据全部返回，false 表示还有数据待返回，可将 Context 作为入参，继续查询返回结果。
		Listover *bool `json:"Listover,omitempty" name:"Listover"`

		// 检索上下文，当 ListOver 为false时，可以用此上下文，继续读取后续数据
		Context *string `json:"Context,omitempty" name:"Context"`

		// 日志数据结果数组，返回对应时间点及取值。
		Results []*DeviceCommLogItem `json:"Results,omitempty" name:"Results" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceCommLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceCommLogResponse) FromJsonString(s string) error {
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
	Limit []*uint64 `json:"Limit,omitempty" name:"Limit" list`

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

func (r *DescribeDeviceEventHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceEventHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		EventHistory []*EventHistoryItem `json:"EventHistory,omitempty" name:"EventHistory" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceEventHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceEventHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

func (r *DescribeFirmwareRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *DescribeFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFirmwareResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFirmwareTaskDevicesRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 筛选条件
	Filters []*SearchKeyword `json:"Filters,omitempty" name:"Filters" list`

	// 查询偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFirmwareTaskDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFirmwareTaskDevicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFirmwareTaskDevicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 固件升级任务的设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 固件升级任务的设备列表
		Devices []*DeviceUpdateStatus `json:"Devices,omitempty" name:"Devices" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFirmwareTaskDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFirmwareTaskDevicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeFirmwareTaskDistributionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFirmwareTaskDistributionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 固件升级任务状态分布信息
		StatusInfos []*StatusStatistic `json:"StatusInfos,omitempty" name:"StatusInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFirmwareTaskDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFirmwareTaskDistributionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeFirmwareTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFirmwareTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *DescribeFirmwareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFirmwareTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeFirmwareTaskStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFirmwareTaskStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *DescribeFirmwareTaskStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFirmwareTaskStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	Filters []*SearchKeyword `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeFirmwareTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFirmwareTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFirmwareTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 固件升级任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskInfos []*FirmwareTaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos" list`

		// 固件升级任务总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFirmwareTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFirmwareTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeForwardRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 腾讯云账号
		Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

		// 队列名称
		QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

		// 产品ID
		ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

		// 消息类型
		MsgType *uint64 `json:"MsgType,omitempty" name:"MsgType"`

		// 结果
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
	} `json:"Response"`
}

func (r *DescribeForwardRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeForwardRuleResponse) FromJsonString(s string) error {
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

type DescribeProductRequest struct {
	*tchttp.BaseRequest

	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品详情
		Data *VideoProduct `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeProductsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 产品详情列表
		Data []*VideoProduct `json:"Data,omitempty" name:"Data" list`

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

	// 设备启用状态
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`

	// 设备过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
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

func (r *EditFirmwareRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

type GetAllFirmwareVersionRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductID *string `json:"ProductID,omitempty" name:"ProductID"`
}

func (r *GetAllFirmwareVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAllFirmwareVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetAllFirmwareVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 无
		Version []*string `json:"Version,omitempty" name:"Version" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAllFirmwareVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetAllFirmwareVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *GetFirmwareURLRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFirmwareURLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 固件URL
		Url *string `json:"Url,omitempty" name:"Url"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFirmwareURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFirmwareURLResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ImportModelDefinitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportModelDefinitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportModelDefinitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportModelDefinitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	Filters []*SearchKeyword `json:"Filters,omitempty" name:"Filters" list`
}

func (r *ListFirmwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListFirmwaresRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListFirmwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 固件总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 固件列表
		Firmwares []*FirmwareInfo `json:"Firmwares,omitempty" name:"Firmwares" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListFirmwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListFirmwaresResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 队列类型
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

func (r *ModifyForwardRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 腾讯云账号
		Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

		// 产品ID
		ProductID *string `json:"ProductID,omitempty" name:"ProductID"`

		// 结果
		Result *uint64 `json:"Result,omitempty" name:"Result"`

		// 错误信息
		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

		// 队列类型
		QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyForwardRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyForwardRuleResponse) FromJsonString(s string) error {
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

type ModifyProductRequest struct {
	*tchttp.BaseRequest

	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 修改的产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 修改的产品描述
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`
}

func (r *ModifyProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

type ResetCloudStorageRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *ResetCloudStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetCloudStorageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetCloudStorageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *RetryDeviceFirmwareTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RetryDeviceFirmwareTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryDeviceFirmwareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RetryDeviceFirmwareTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchKeyword struct {

	// 搜索条件的Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 搜索条件的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SetForwardAuthRequest struct {
	*tchttp.BaseRequest

	// 控制台Skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// 消息队列类型
	QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`
}

func (r *SetForwardAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetForwardAuthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetForwardAuthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 腾讯云账号
		Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

		// 结果
		Result *uint64 `json:"Result,omitempty" name:"Result"`

		// 角色名
		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

		// 角色ID
		RoleID *uint64 `json:"RoleID,omitempty" name:"RoleID"`

		// 消息队列类型
		QueueType *uint64 `json:"QueueType,omitempty" name:"QueueType"`

		// 错误消息
		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetForwardAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

func (r *TransferCloudStorageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TransferCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransferCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransferCloudStorageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *UploadFirmwareRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	Features []*string `json:"Features,omitempty" name:"Features" list`

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
}
