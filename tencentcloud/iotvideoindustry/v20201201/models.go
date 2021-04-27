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

package v20201201

import (
    "encoding/json"
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AllDeviceInfo struct {

	// 设备唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 设备类型；2：IPC
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceType *int64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 设备状态；0：设备不在线；1：设备在线；2：设备隔离中；3：设备未注册
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 设备扩展属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInformation *string `json:"ExtraInformation,omitempty" name:"ExtraInformation"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 设备绑定分组路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupPath *string `json:"GroupPath,omitempty" name:"GroupPath"`

	// 设备编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceCode *string `json:"DeviceCode,omitempty" name:"DeviceCode"`

	// 是否存在录像,，0:不存在；1：存在
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRecord *int64 `json:"IsRecord,omitempty" name:"IsRecord"`

	// 该设备是否可录制
	// 注意：此字段可能返回 null，表示取不到有效值。
	Recordable *int64 `json:"Recordable,omitempty" name:"Recordable"`
}

type BindGroupDevicesRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 设备唯一标识列表
	DeviceList []*string `json:"DeviceList,omitempty" name:"DeviceList" list`
}

func (r *BindGroupDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindGroupDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "DeviceList")
	if len(f) > 0 {
		return errors.New("BindGroupDevicesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindGroupDevicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindGroupDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindGroupDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ControlDevicePTZRequest struct {
	*tchttp.BaseRequest

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// PTZ控制命令类型：
	// stop - 停止当前PTZ信令
	// left - 向左移动
	// right - 向右移动
	// up - 向上移动
	// down - 向下移动
	// leftUp - 左上移动
	// leftDown - 左下移动
	// rightUp - 右上移动
	// rightDown - 右下移动
	// zoomOut - 镜头缩小
	// zoomIn - 镜头放大
	// irisIn - 光圈缩小
	// irisOut - 光圈放大
	// focusIn - 焦距变近
	// focusOut - 焦距变远
	Command *string `json:"Command,omitempty" name:"Command"`
}

func (r *ControlDevicePTZRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDevicePTZRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Command")
	if len(f) > 0 {
		return errors.New("ControlDevicePTZRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ControlDevicePTZResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ControlDevicePTZResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDevicePTZResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceGroupRequest struct {
	*tchttp.BaseRequest

	// 分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 父分组ID
	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`

	// 分组描述
	GroupDescribe *string `json:"GroupDescribe,omitempty" name:"GroupDescribe"`
}

func (r *CreateDeviceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "ParentId")
	delete(f, "GroupDescribe")
	if len(f) > 0 {
		return errors.New("CreateDeviceGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 响应结果，“OK”为成功，其他为失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeviceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceRequest struct {
	*tchttp.BaseRequest

	// 设备名称
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 设备密码
	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`

	// 设备需要绑定的分组ID，参数为空则默认绑定到根分组
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NickName")
	delete(f, "PassWord")
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("CreateDeviceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备编码
	// 注意：此字段可能返回 null，表示取不到有效值。
		DeviceCode *string `json:"DeviceCode,omitempty" name:"DeviceCode"`

		// 设备唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
		DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

		// 设备虚拟组信息，仅在创建NVR/VMS时返回该值
	// 注意：此字段可能返回 null，表示取不到有效值。
		VirtualGroupId *string `json:"VirtualGroupId,omitempty" name:"VirtualGroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordPlanRequest struct {
	*tchttp.BaseRequest

	// 计划名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间模板ID
	TimeTemplateId *string `json:"TimeTemplateId,omitempty" name:"TimeTemplateId"`

	// 触发录制的事件类别 1:全部
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 该录制计划绑定的设备列表
	Devices []*DeviceItem `json:"Devices,omitempty" name:"Devices" list`
}

func (r *CreateRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "TimeTemplateId")
	delete(f, "EventId")
	delete(f, "Devices")
	if len(f) > 0 {
		return errors.New("CreateRecordPlanRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制计划ID
		PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTimeTemplateRequest struct {
	*tchttp.BaseRequest

	// 时间模板名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否为每周全时录制（即7*24h录制），0：非全时录制，1；全时录制，默认0
	IsAllWeek *int64 `json:"IsAllWeek,omitempty" name:"IsAllWeek"`

	// 当IsAllWeek为0时必选，用于描述模板的各个时间片段
	TimeTemplateSpecs []*TimeTemplateSpec `json:"TimeTemplateSpecs,omitempty" name:"TimeTemplateSpecs" list`
}

func (r *CreateTimeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTimeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "IsAllWeek")
	delete(f, "TimeTemplateSpecs")
	if len(f) > 0 {
		return errors.New("CreateTimeTemplateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTimeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 时间模板ID
		TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTimeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTimeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceGroupRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteDeviceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("DeleteDeviceGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 响应结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeviceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

func (r *DeleteDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return errors.New("DeleteDeviceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordPlanRequest struct {
	*tchttp.BaseRequest

	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
}

func (r *DeleteRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return errors.New("DeleteRecordPlanRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作结果，OK：成功，其他：失败
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTimeTemplateRequest struct {
	*tchttp.BaseRequest

	// 时间模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteTimeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return errors.New("DeleteTimeTemplateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTimeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作结果，OK：成功，其他：失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTimeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllDeviceListRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制，默认200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 设备名称，需要模糊匹配设备名称时为必填
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// DeviceId列表，需要精确查找设备时为必填
	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds" list`
}

func (r *DescribeAllDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllDeviceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NickName")
	delete(f, "DeviceIds")
	if len(f) > 0 {
		return errors.New("DescribeAllDeviceListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 设备详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Devices []*AllDeviceInfo `json:"Devices,omitempty" name:"Devices" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceGroupRequest struct {
	*tchttp.BaseRequest

	// 设备唯一标识列表
	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds" list`
}

func (r *DescribeDeviceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceIds")
	if len(f) > 0 {
		return errors.New("DescribeDeviceGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备所在分组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		DevGroups []*DevGroupInfo `json:"DevGroups,omitempty" name:"DevGroups" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicePassWordRequest struct {
	*tchttp.BaseRequest

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

func (r *DescribeDevicePassWordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePassWordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return errors.New("DescribeDevicePassWordRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicePassWordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备密码
		PassWord *string `json:"PassWord,omitempty" name:"PassWord"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDevicePassWordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePassWordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceStreamsData struct {

	// rtsp地址
	RtspAddr *string `json:"RtspAddr,omitempty" name:"RtspAddr"`

	// rtmp地址
	RtmpAddr *string `json:"RtmpAddr,omitempty" name:"RtmpAddr"`

	// hls地址
	HlsAddr *string `json:"HlsAddr,omitempty" name:"HlsAddr"`

	// flv地址
	FlvAddr *string `json:"FlvAddr,omitempty" name:"FlvAddr"`
}

type DescribeDeviceStreamsRequest struct {
	*tchttp.BaseRequest

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 流地址失效时间
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *DescribeDeviceStreamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStreamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ExpireTime")
	if len(f) > 0 {
		return errors.New("DescribeDeviceStreamsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceStreamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备实时流地址列表
		Data *DescribeDeviceStreamsData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceStreamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStreamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupByIdRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return errors.New("DescribeGroupByIdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupByIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分组信息详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		Group *GroupItem `json:"Group,omitempty" name:"Group"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupByPathRequest struct {
	*tchttp.BaseRequest

	// 分组路径，格式为/aaa(/bbb/ccc)
	GroupPath *string `json:"GroupPath,omitempty" name:"GroupPath"`
}

func (r *DescribeGroupByPathRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupByPathRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupPath")
	if len(f) > 0 {
		return errors.New("DescribeGroupByPathRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupByPathResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分组信息详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		Group *GroupItem `json:"Group,omitempty" name:"Group"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupByPathResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupByPathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupDevicesRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制值，默认200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 设备名称，根据设备名称模糊匹配时必填
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 过滤不可录制设备
	Recordable *int64 `json:"Recordable,omitempty" name:"Recordable"`
}

func (r *DescribeGroupDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NickName")
	delete(f, "Recordable")
	if len(f) > 0 {
		return errors.New("DescribeGroupDevicesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupDevicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分组绑定的设备数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 设备详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		DeviceList []*GroupDeviceItem `json:"DeviceList,omitempty" name:"DeviceList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsRequest struct {
	*tchttp.BaseRequest

	// 分组ID列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds" list`
}

func (r *DescribeGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	if len(f) > 0 {
		return errors.New("DescribeGroupsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分组详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Groups []*GroupInfo `json:"Groups,omitempty" name:"Groups" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordStreamData struct {

	// Rtsp地址
	RtspAddr *string `json:"RtspAddr,omitempty" name:"RtspAddr"`

	// Rtmp地址
	RtmpAddr *string `json:"RtmpAddr,omitempty" name:"RtmpAddr"`

	// Hls地址
	HlsAddr *string `json:"HlsAddr,omitempty" name:"HlsAddr"`

	// Flv地址
	FlvAddr *string `json:"FlvAddr,omitempty" name:"FlvAddr"`

	// 流Id
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`
}

type DescribeRecordStreamRequest struct {
	*tchttp.BaseRequest

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 流失效时间
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 录像文件Id
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 录像流开始时间，当录像文件Id为空时有效
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 录像流结束时间，当录像文件Id为空时有效
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeRecordStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ExpireTime")
	delete(f, "RecordId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return errors.New("DescribeRecordStreamRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果
		Data *DescribeRecordStreamData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSIPServerRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSIPServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSIPServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeSIPServerRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSIPServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SIP服务器相关配置项
		Data *ServerConfiguration `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSIPServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSIPServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticDetailsRequest struct {
	*tchttp.BaseRequest

	// 开始日期，格式【YYYY-MM-DD】
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式【YYYY-MM-DD】
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 统计项。取值范围：
	// 1.录制设备数：RecordingDevice
	// 2.非录制设备数：NonRecordingDevice
	// 3.观看流量总数：WatchFlux
	// 4.已用存储容量总数：StorageUsage
	StatisticField *string `json:"StatisticField,omitempty" name:"StatisticField"`
}

func (r *DescribeStatisticDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "StatisticField")
	if len(f) > 0 {
		return errors.New("DescribeStatisticDetailsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*StatisticItem `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatisticDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticSummaryRequest struct {
	*tchttp.BaseRequest

	// 指定日期。格式【YYYY-MM-DD】
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DescribeStatisticSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Date")
	if len(f) > 0 {
		return errors.New("DescribeStatisticSummaryRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		RecordingDevice *uint64 `json:"RecordingDevice,omitempty" name:"RecordingDevice"`

		// 非录制设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		NonRecordingDevice *uint64 `json:"NonRecordingDevice,omitempty" name:"NonRecordingDevice"`

		// 观看流量总数。为直播观看流量与点播观看流量之和。单位：GB
	// 注意：此字段可能返回 null，表示取不到有效值。
		WatchFlux *float64 `json:"WatchFlux,omitempty" name:"WatchFlux"`

		// 累计有效存储容量总数。单位：GB
	// 注意：此字段可能返回 null，表示取不到有效值。
		StorageUsage *float64 `json:"StorageUsage,omitempty" name:"StorageUsage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatisticSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubGroupsRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称，根据名称模糊匹配子分组时为必填
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数，默认200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否统计子分组下的设备数，0：统计，1：不统计
	OnlyGroup *int64 `json:"OnlyGroup,omitempty" name:"OnlyGroup"`
}

func (r *DescribeSubGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OnlyGroup")
	if len(f) > 0 {
		return errors.New("DescribeSubGroupsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子分组详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		GroupList []*GroupItem `json:"GroupList,omitempty" name:"GroupList" list`

		// 子分组总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoListRequest struct {
	*tchttp.BaseRequest

	// 开始时间戳，秒级
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳，秒级
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

func (r *DescribeVideoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return errors.New("DescribeVideoListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 已废弃
		VideoList *RecordTaskItem `json:"VideoList,omitempty" name:"VideoList"`

		// 录像详情列表
		RecordList []*RecordTaskItem `json:"RecordList,omitempty" name:"RecordList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVideoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DevGroupInfo struct {

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组路径
	GroupPath *string `json:"GroupPath,omitempty" name:"GroupPath"`

	// 父分组ID
	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`

	// 设备错误，仅在用户没权限或者设备已删除时返回具体结果
	Error *string `json:"Error,omitempty" name:"Error"`
}

type DeviceItem struct {

	// 设备唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

type GetRecordDatesByDevRequest struct {
	*tchttp.BaseRequest

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制量，默认200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetRecordDatesByDevRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRecordDatesByDevRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("GetRecordDatesByDevRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRecordDatesByDevResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 含有录像文件的日期列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Dates []*string `json:"Dates,omitempty" name:"Dates" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRecordDatesByDevResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRecordDatesByDevResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRecordPlanByDevRequest struct {
	*tchttp.BaseRequest

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

func (r *GetRecordPlanByDevRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRecordPlanByDevRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return errors.New("GetRecordPlanByDevRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRecordPlanByDevResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制计划详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		Plan *RecordPlanItem `json:"Plan,omitempty" name:"Plan"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRecordPlanByDevResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRecordPlanByDevResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRecordPlanByIdRequest struct {
	*tchttp.BaseRequest

	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
}

func (r *GetRecordPlanByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRecordPlanByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return errors.New("GetRecordPlanByIdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRecordPlanByIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制计划详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		Plan *RecordPlanItem `json:"Plan,omitempty" name:"Plan"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRecordPlanByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRecordPlanByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRecordPlansRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRecordPlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRecordPlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("GetRecordPlansRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRecordPlansResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制计划详情·列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Plans []*RecordPlanItem `json:"Plans,omitempty" name:"Plans" list`

		// 录制计划总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRecordPlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRecordPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTimeTemplateByIdRequest struct {
	*tchttp.BaseRequest

	// 时间模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *GetTimeTemplateByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTimeTemplateByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return errors.New("GetTimeTemplateByIdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetTimeTemplateByIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 时间模板详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		Template *TimeTemplateItem `json:"Template,omitempty" name:"Template"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTimeTemplateByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTimeTemplateByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTimeTemplatesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetTimeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTimeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("GetTimeTemplatesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetTimeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 时间模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Templates []*TimeTemplateItem `json:"Templates,omitempty" name:"Templates" list`

		// 时间模板总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTimeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTimeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetVideoListByConRequest struct {
	*tchttp.BaseRequest

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制量，默认200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 0：查询指定日期的录像；1：查询最近一天的录像；默认0
	LatestDay *int64 `json:"LatestDay,omitempty" name:"LatestDay"`

	// 指定某天。取值【YYYY-MM-DD】
	// 当LatestDay为空或为0时，本参数不允许为空。
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *GetVideoListByConRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVideoListByConRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "LatestDay")
	delete(f, "Date")
	if len(f) > 0 {
		return errors.New("GetVideoListByConRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetVideoListByConResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录像详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		VideoList []*RecordTaskItem `json:"VideoList,omitempty" name:"VideoList" list`

		// 录像总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetVideoListByConResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVideoListByConResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupDeviceItem struct {

	// 设备唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 设备状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInformation *string `json:"ExtraInformation,omitempty" name:"ExtraInformation"`

	// 设备类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceType *int64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// rtsp地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RTSPUrl *string `json:"RTSPUrl,omitempty" name:"RTSPUrl"`

	// 设备编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceCode *string `json:"DeviceCode,omitempty" name:"DeviceCode"`

	// 是否存在录像
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRecord *int64 `json:"IsRecord,omitempty" name:"IsRecord"`

	// 该设备是否可录制
	// 注意：此字段可能返回 null，表示取不到有效值。
	Recordable *int64 `json:"Recordable,omitempty" name:"Recordable"`
}

type GroupInfo struct {

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组类型
	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`

	// 分组路径
	GroupPath *string `json:"GroupPath,omitempty" name:"GroupPath"`

	// 父分组ID
	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`

	// 分组描述
	GroupDescribe *string `json:"GroupDescribe,omitempty" name:"GroupDescribe"`

	// 扩展信息
	ExtraInformation *string `json:"ExtraInformation,omitempty" name:"ExtraInformation"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 分组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupStatus *int64 `json:"GroupStatus,omitempty" name:"GroupStatus"`

	// 设备不存在时产生的错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitempty" name:"Error"`
}

type GroupItem struct {

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 父分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`

	// 分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupPath *string `json:"GroupPath,omitempty" name:"GroupPath"`

	// 分组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupDescribe *string `json:"GroupDescribe,omitempty" name:"GroupDescribe"`

	// 分组绑定设备数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceNum *int64 `json:"DeviceNum,omitempty" name:"DeviceNum"`

	// 子分组数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubGroupNum *int64 `json:"SubGroupNum,omitempty" name:"SubGroupNum"`

	// 分组附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInformation *string `json:"ExtraInformation,omitempty" name:"ExtraInformation"`

	// 分组类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 分组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupStatus *int64 `json:"GroupStatus,omitempty" name:"GroupStatus"`
}

type ModifyDeviceDataRequest struct {
	*tchttp.BaseRequest

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 设备名称
	NickName *string `json:"NickName,omitempty" name:"NickName"`
}

func (r *ModifyDeviceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "NickName")
	if len(f) > 0 {
		return errors.New("ModifyDeviceDataRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作结果,“OK”表示成功，其他表示失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordPlanItem struct {

	// 计划ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 计划名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeTemplateId *string `json:"TimeTemplateId,omitempty" name:"TimeTemplateId"`

	// 时间模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeTemplateName *string `json:"TimeTemplateName,omitempty" name:"TimeTemplateName"`

	// 录制类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 绑定的设备列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Devices []*DeviceItem `json:"Devices,omitempty" name:"Devices" list`
}

type RecordTaskItem struct {

	// 录像任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordTaskId *string `json:"RecordTaskId,omitempty" name:"RecordTaskId"`

	// 录制计划ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordPlanId *string `json:"RecordPlanId,omitempty" name:"RecordPlanId"`

	// 本录制片段开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 本录制片段结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 录制模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 本录制片段对应的录制文件URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// 本录制片段当前的录制状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordStatus *int64 `json:"RecordStatus,omitempty" name:"RecordStatus"`
}

type ServerConfiguration struct {

	// SIP服务器地址
	Host *string `json:"Host,omitempty" name:"Host"`

	// SIP服务器端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// SIP服务器编码
	Serial *string `json:"Serial,omitempty" name:"Serial"`

	// SIP服务器域
	Realm *string `json:"Realm,omitempty" name:"Realm"`
}

type StatisticItem struct {

	// 日期。格式【YYYY-MM-DD】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 统计数额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sum *float64 `json:"Sum,omitempty" name:"Sum"`
}

type TimeTemplateItem struct {

	// 时间模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否全时录制，即7*24小时录制
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllWeek *int64 `json:"IsAllWeek,omitempty" name:"IsAllWeek"`

	// 是否为自定义模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 时间片段详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeTemplateSpecs []*TimeTemplateSpec `json:"TimeTemplateSpecs,omitempty" name:"TimeTemplateSpecs" list`
}

type TimeTemplateSpec struct {

	// 一周中的周几
	// 注意：此字段可能返回 null，表示取不到有效值。
	DayofWeek *int64 `json:"DayofWeek,omitempty" name:"DayofWeek"`

	// 时间片段的开始时分。格式【HH:MM】
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 时间片段的结束时分。格式【HH:MM】
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type UpdateDeviceGroupRequest struct {
	*tchttp.BaseRequest

	// 分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组描述
	GroupDescribe *string `json:"GroupDescribe,omitempty" name:"GroupDescribe"`

	// 新父分组ID，用于修改分组路径
	NewParentId *string `json:"NewParentId,omitempty" name:"NewParentId"`
}

func (r *UpdateDeviceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "GroupId")
	delete(f, "GroupDescribe")
	delete(f, "NewParentId")
	if len(f) > 0 {
		return errors.New("UpdateDeviceGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDeviceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDevicePassWordRequest struct {
	*tchttp.BaseRequest

	// 设备密码
	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

func (r *UpdateDevicePassWordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDevicePassWordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PassWord")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return errors.New("UpdateDevicePassWordRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDevicePassWordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作结果，“OK”表示成功，其他表示失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDevicePassWordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDevicePassWordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecordPlanRequest struct {
	*tchttp.BaseRequest

	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 计划名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间模板ID
	TimeTemplateId *string `json:"TimeTemplateId,omitempty" name:"TimeTemplateId"`

	// 触发录制的事件 1：全部
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 录制设备列表
	Devices []*DeviceItem `json:"Devices,omitempty" name:"Devices" list`

	// 是否更新绑定此录制计划的设备列表
	// 0 - 不更新
	// 1 - 更新，如果Devices参数为空则清空设备列表，Devices不为空则全量更新设备列表
	IsModifyDevices *int64 `json:"IsModifyDevices,omitempty" name:"IsModifyDevices"`
}

func (r *UpdateRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "Name")
	delete(f, "TimeTemplateId")
	delete(f, "EventId")
	delete(f, "Devices")
	delete(f, "IsModifyDevices")
	if len(f) > 0 {
		return errors.New("UpdateRecordPlanRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作结果
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTimeTemplateRequest struct {
	*tchttp.BaseRequest

	// 时间模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 时间模板名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否全时录制，即7*24小时录制。
	// 0：非全时录制；1：全时录制。默认1
	IsAllWeek *int64 `json:"IsAllWeek,omitempty" name:"IsAllWeek"`

	// 录制时间片段
	TimeTemplateSpecs []*TimeTemplateSpec `json:"TimeTemplateSpecs,omitempty" name:"TimeTemplateSpecs" list`
}

func (r *UpdateTimeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTimeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Name")
	delete(f, "IsAllWeek")
	delete(f, "TimeTemplateSpecs")
	if len(f) > 0 {
		return errors.New("UpdateTimeTemplateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTimeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作结果，“OK”表示成功，其他表示失败。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTimeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTimeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
