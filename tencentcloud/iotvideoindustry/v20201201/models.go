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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
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

	// 设备接入协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type BindGroupDevicesRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 设备唯一标识列表
	DeviceList []*string `json:"DeviceList,omitempty" name:"DeviceList"`
}

func (r *BindGroupDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindGroupDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "DeviceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindGroupDevicesRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

func (r *ControlDevicePTZRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDevicePTZRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Command")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlDevicePTZRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDevicePTZResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ControlRecordStreamRequest struct {
	*tchttp.BaseRequest

	// 设备Id，设备的唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 流Id，流的唯一标识
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// |控制参数，CmdJson结构转义的json字符串。| Action  | string  |是|控制动作，play(用于暂停后恢复播放)、pause（暂停）、teardown(停止)、jump(拖动播放)
	// | Offset  | uint  |否|拖动播放时的时间偏移量（相对于起始时间）,单位：秒
	// 目前支持的command：
	// "Command": "{"Action":"PAUSE"}" 暂停
	// "Command": "{"Action":"PLAY"}" 暂停恢复
	// "Command": "{"Action":"PLAY","Offset":"15"}" 位置偏移，可以替代jump操作
	Command *string `json:"Command,omitempty" name:"Command"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

func (r *ControlRecordStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlRecordStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "StreamId")
	delete(f, "Command")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlRecordStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ControlRecordStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ControlRecordStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlRecordStreamResponse) FromJsonString(s string) error {
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceGroupRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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

	// 设备类型 2:国标IPC设备; 3:NVR设备
	DeviceType *int64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 设备需要绑定的分组ID，参数为空则默认绑定到根分组
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
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
	delete(f, "NickName")
	delete(f, "PassWord")
	delete(f, "DeviceType")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveChannelRequest struct {
	*tchttp.BaseRequest

	// 直播频道名称
	LiveChannelName *string `json:"LiveChannelName,omitempty" name:"LiveChannelName"`

	// 直播频道类型 1：固定直播；2：移动直播
	LiveChannelType *int64 `json:"LiveChannelType,omitempty" name:"LiveChannelType"`
}

func (r *CreateLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LiveChannelName")
	delete(f, "LiveChannelType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 直播频道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

		// 直播频道推流地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		PushStreamAddress *string `json:"PushStreamAddress,omitempty" name:"PushStreamAddress"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordPlanRequest struct {
	*tchttp.BaseRequest

	// 录制计划名
	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`

	// 计划类型 1：固定直播 2：移动直播
	PlanType *int64 `json:"PlanType,omitempty" name:"PlanType"`

	// 时间模板ID,固定直播时为必填
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 录制文件存储时长，单位天，默认30天
	RecordStorageTime *int64 `json:"RecordStorageTime,omitempty" name:"RecordStorageTime"`

	// 绑定的直播频道ID列表
	LiveChannelIds []*string `json:"LiveChannelIds,omitempty" name:"LiveChannelIds"`
}

func (r *CreateLiveRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanName")
	delete(f, "PlanType")
	delete(f, "TemplateId")
	delete(f, "RecordStorageTime")
	delete(f, "LiveChannelIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLiveRecordPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateLiveRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制计划名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLiveRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLiveRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMessageForwardRequest struct {
	*tchttp.BaseRequest

	// 区域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 区域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 实例ID
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// json数组， 转发类型 1: 告警 2:GPS
	MessageType *string `json:"MessageType,omitempty" name:"MessageType"`

	// kafka topic id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// kafka topic 名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *CreateMessageForwardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMessageForwardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegionId")
	delete(f, "RegionName")
	delete(f, "Instance")
	delete(f, "InstanceName")
	delete(f, "MessageType")
	delete(f, "TopicId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMessageForwardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMessageForwardResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		IntId *int64 `json:"IntId,omitempty" name:"IntId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMessageForwardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMessageForwardResponse) FromJsonString(s string) error {
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
	Devices []*DeviceItem `json:"Devices,omitempty" name:"Devices"`

	// 存储周期
	RecordStorageTime *int64 `json:"RecordStorageTime,omitempty" name:"RecordStorageTime"`
}

func (r *CreateRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "RecordStorageTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecordPlanRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSceneRequest struct {
	*tchttp.BaseRequest

	// 场景名称
	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`

	// 场景触发规则
	SceneTrigger *string `json:"SceneTrigger,omitempty" name:"SceneTrigger"`

	// 录制时长 (秒)
	RecordDuration *int64 `json:"RecordDuration,omitempty" name:"RecordDuration"`

	// 录像存储时长(天)
	StoreDuration *int64 `json:"StoreDuration,omitempty" name:"StoreDuration"`

	// 设备列表
	Devices []*DeviceItem `json:"Devices,omitempty" name:"Devices"`
}

func (r *CreateSceneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSceneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SceneName")
	delete(f, "SceneTrigger")
	delete(f, "RecordDuration")
	delete(f, "StoreDuration")
	delete(f, "Devices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSceneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSceneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 场景ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		IntId *int64 `json:"IntId,omitempty" name:"IntId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSceneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSceneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTimeTemplateRequest struct {
	*tchttp.BaseRequest

	// 时间模板名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否为每周全时录制（即7*24h录制），0：非全时录制，1；全时录制，默认0
	IsAllWeek *int64 `json:"IsAllWeek,omitempty" name:"IsAllWeek"`

	// 当IsAllWeek为0时必选，用于描述模板的各个时间片段
	TimeTemplateSpecs []*TimeTemplateSpec `json:"TimeTemplateSpecs,omitempty" name:"TimeTemplateSpecs"`
}

func (r *CreateTimeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTimeTemplateRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTimeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteChannelRequest struct {
	*tchttp.BaseRequest

	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

func (r *DeleteChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteChannelResponse) FromJsonString(s string) error {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceGroupRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveChannelRequest struct {
	*tchttp.BaseRequest

	// 直播频道ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`
}

func (r *DeleteLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LiveChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordPlanRequest struct {
	*tchttp.BaseRequest

	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
}

func (r *DeleteLiveRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveRecordPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除状态描述
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveVideoListRequest struct {
	*tchttp.BaseRequest

	// 视频ID 列表, 大小限制(100)
	IntIDs []*uint64 `json:"IntIDs,omitempty" name:"IntIDs"`
}

func (r *DeleteLiveVideoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveVideoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IntIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLiveVideoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLiveVideoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLiveVideoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLiveVideoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMessageForwardRequest struct {
	*tchttp.BaseRequest

	// 配置ID
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`
}

func (r *DeleteMessageForwardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMessageForwardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IntId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMessageForwardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMessageForwardResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMessageForwardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMessageForwardResponse) FromJsonString(s string) error {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordPlanRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSceneRequest struct {
	*tchttp.BaseRequest

	// 场景ID
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`
}

func (r *DeleteSceneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSceneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IntId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSceneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSceneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSceneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSceneResponse) FromJsonString(s string) error {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTimeTemplateRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVideoListRequest struct {
	*tchttp.BaseRequest

	// 视频ID列表长度限制100内
	InitIDs []*int64 `json:"InitIDs,omitempty" name:"InitIDs"`
}

func (r *DeleteVideoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVideoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InitIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVideoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVideoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVideoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVideoListResponse) FromJsonString(s string) error {
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
	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds"`

	// 设备类型过滤
	DeviceTypes []*int64 `json:"DeviceTypes,omitempty" name:"DeviceTypes"`
}

func (r *DescribeAllDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "DeviceTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllDeviceListRequest has unknown keys!", "")
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
		Devices []*AllDeviceInfo `json:"Devices,omitempty" name:"Devices"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindSceneDevicesRequest struct {
	*tchttp.BaseRequest

	// 场景ID
	SceneId *int64 `json:"SceneId,omitempty" name:"SceneId"`

	// 偏移值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 条数限制最大不能超过1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBindSceneDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindSceneDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SceneId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindSceneDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindSceneDevicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 设备列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		List []*DeviceItem `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindSceneDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindSceneDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChannelsByLiveRecordPlanRequest struct {
	*tchttp.BaseRequest

	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeChannelsByLiveRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelsByLiveRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChannelsByLiveRecordPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChannelsByLiveRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总个数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 通道详情数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		LiveChannels []*LiveChannelItem `json:"LiveChannels,omitempty" name:"LiveChannels"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeChannelsByLiveRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelsByLiveRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceGroupRequest struct {
	*tchttp.BaseRequest

	// 设备唯一标识列表
	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds"`
}

func (r *DescribeDeviceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备所在分组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		DevGroups []*DevGroupInfo `json:"DevGroups,omitempty" name:"DevGroups"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePassWordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicePassWordRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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

	// 通道唯一标识（接口升级字段为必填）
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

func (r *DescribeDeviceStreamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStreamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ExpireTime")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceStreamsRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupByIdRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupByPathRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupByPathRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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

	// 当Group是普通组的时候，支持根据deviceTypes筛选类型
	//  0: 普通摄像头1:  国标VMS设备 2: 国标IPC设备 3: 国标NVR设备  4: 国标NVR通道 5: 国标VMS通道 6: 国标IPC通道 9: 智能告警设备 10: 带有RTSP固定地址的设备
	DeviceTypes []*int64 `json:"DeviceTypes,omitempty" name:"DeviceTypes"`
}

func (r *DescribeGroupDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "DeviceTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupDevicesRequest has unknown keys!", "")
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
		DeviceList []*GroupDeviceItem `json:"DeviceList,omitempty" name:"DeviceList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsRequest struct {
	*tchttp.BaseRequest

	// 分组ID列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
}

func (r *DescribeGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分组详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Groups []*GroupInfo `json:"Groups,omitempty" name:"Groups"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPCChannelsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制，默认0
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道类型  0: 未知类型 1: 视频通道 2:  音频通道 3: 告警通道
	ChannelTypes []*uint64 `json:"ChannelTypes,omitempty" name:"ChannelTypes"`
}

func (r *DescribeIPCChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPCChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DeviceId")
	delete(f, "ChannelTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPCChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIPCChannelsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 通道详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		DeviceList []*GroupDeviceItem `json:"DeviceList,omitempty" name:"DeviceList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPCChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPCChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveChannelListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 直播频道类型，1：固定直播；2：移动直播
	LiveChannelType *int64 `json:"LiveChannelType,omitempty" name:"LiveChannelType"`

	// 直播录制计划ID, null: 直播录制计划为空
	RecordPlanId *string `json:"RecordPlanId,omitempty" name:"RecordPlanId"`

	// 频道名称 (支持模糊搜索)
	LiveChannelName *string `json:"LiveChannelName,omitempty" name:"LiveChannelName"`
}

func (r *DescribeLiveChannelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveChannelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "LiveChannelType")
	delete(f, "RecordPlanId")
	delete(f, "LiveChannelName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveChannelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveChannelListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 频道总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 频道信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		LiveChannels []*LiveChannelInfo `json:"LiveChannels,omitempty" name:"LiveChannels"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveChannelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveChannelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveChannelRequest struct {
	*tchttp.BaseRequest

	// 频道ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`
}

func (r *DescribeLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LiveChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 频道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

		// 频道名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		LiveChannelName *string `json:"LiveChannelName,omitempty" name:"LiveChannelName"`

		// 直播频道类型 1：固定直播；2：移动直播
	// 注意：此字段可能返回 null，表示取不到有效值。
		LiveChannelType *int64 `json:"LiveChannelType,omitempty" name:"LiveChannelType"`

		// 通道直播状态：1: 未推流，2: 推流中
	// 注意：此字段可能返回 null，表示取不到有效值。
		LiveStatus *int64 `json:"LiveStatus,omitempty" name:"LiveStatus"`

		// 推流地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		PushStreamAddress *string `json:"PushStreamAddress,omitempty" name:"PushStreamAddress"`

		// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreateTime []*string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		UpdateTime []*string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordPlanByIdRequest struct {
	*tchttp.BaseRequest

	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
}

func (r *DescribeLiveRecordPlanByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveRecordPlanByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveRecordPlanByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordPlanByIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计划名称
		PlanName *string `json:"PlanName,omitempty" name:"PlanName"`

		// 模板ID
		TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

		// 模板名称
		TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

		// 存储时间
		RecordStorageTime *int64 `json:"RecordStorageTime,omitempty" name:"RecordStorageTime"`

		// 计划类型
		PlanType *int64 `json:"PlanType,omitempty" name:"PlanType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordPlanByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveRecordPlanByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordPlanIdsRequest struct {
	*tchttp.BaseRequest

	// 时间模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeLiveRecordPlanIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveRecordPlanIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveRecordPlanIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveRecordPlanIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 计划数组
		Plans []*LiveRecordPlanItem `json:"Plans,omitempty" name:"Plans"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveRecordPlanIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveRecordPlanIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamRequest struct {
	*tchttp.BaseRequest

	// 频道ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

	// 过期时间
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *DescribeLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LiveChannelId")
	delete(f, "ExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 拉流地址，只有在推流情况下才有
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *StreamAddress `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveVideoListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 直播ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

	// 开始录制开始时间
	StartRecordTime *int64 `json:"StartRecordTime,omitempty" name:"StartRecordTime"`

	// 开始录制结束时间
	EndRecordTime *int64 `json:"EndRecordTime,omitempty" name:"EndRecordTime"`

	// 过期开始时间
	StartExpireTime *int64 `json:"StartExpireTime,omitempty" name:"StartExpireTime"`

	// 过期结束时间
	EndExpireTime *int64 `json:"EndExpireTime,omitempty" name:"EndExpireTime"`

	// 文件大小范围 Byte
	StartFileSize *int64 `json:"StartFileSize,omitempty" name:"StartFileSize"`

	// 文件大小范围 Byte
	EndFileSize *int64 `json:"EndFileSize,omitempty" name:"EndFileSize"`

	// 录制状态，5: 录制回写完
	IsRecording *int64 `json:"IsRecording,omitempty" name:"IsRecording"`
}

func (r *DescribeLiveVideoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveVideoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "LiveChannelId")
	delete(f, "StartRecordTime")
	delete(f, "EndRecordTime")
	delete(f, "StartExpireTime")
	delete(f, "EndExpireTime")
	delete(f, "StartFileSize")
	delete(f, "EndFileSize")
	delete(f, "IsRecording")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLiveVideoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveVideoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总的条数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 录制任务详情数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		RecordList []*LiveRecordItem `json:"RecordList,omitempty" name:"RecordList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveVideoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLiveVideoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMessageForwardRequest struct {
	*tchttp.BaseRequest

	// 配置ID
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`
}

func (r *DescribeMessageForwardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageForwardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IntId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMessageForwardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMessageForwardResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 区域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

		// 区域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

		// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		Instance *string `json:"Instance,omitempty" name:"Instance"`

		// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

		// 配置ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		IntId *int64 `json:"IntId,omitempty" name:"IntId"`

		// json数组， 转发类型 1: 告警 2:GPS
	// 注意：此字段可能返回 null，表示取不到有效值。
		MessageType *string `json:"MessageType,omitempty" name:"MessageType"`

		// kafka topic id
	// 注意：此字段可能返回 null，表示取不到有效值。
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// 配置创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 用户Uin信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Uin *string `json:"Uin,omitempty" name:"Uin"`

		// kafka topic 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMessageForwardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageForwardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMessageForwardsRequest struct {
	*tchttp.BaseRequest

	// 数量限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMessageForwardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageForwardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMessageForwardsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMessageForwardsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		List []*MessageForward `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMessageForwardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageForwardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordDatesByLiveRequest struct {
	*tchttp.BaseRequest

	// 直播频道ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

	// 分页值，本地录制时参数无效
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制值，本地录制时参数无效
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRecordDatesByLiveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordDatesByLiveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LiveChannelId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordDatesByLiveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordDatesByLiveResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制日期数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		Dates []*string `json:"Dates,omitempty" name:"Dates"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordDatesByLiveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordDatesByLiveResponse) FromJsonString(s string) error {
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

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

func (r *DescribeRecordStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordStreamRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSIPServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSIPServerRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSIPServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScenesRequest struct {
	*tchttp.BaseRequest

	// 条数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeScenesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScenesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScenesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScenesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 场景总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 场景列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		List []*SceneItem `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScenesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScenesResponse) FromJsonString(s string) error {
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
	// 5. X-P2P分享流量: P2PFluxTotal
	// 6. X-P2P峰值带宽: P2PPeakValue
	// 7. RTMP推流路数(直播推流): LivePushTotal
	StatisticField *string `json:"StatisticField,omitempty" name:"StatisticField"`
}

func (r *DescribeStatisticDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStatisticDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*StatisticItem `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatisticDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStatisticSummaryRequest has unknown keys!", "")
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

		// X-P2P分享流量。单位 Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
		P2PFluxTotal *float64 `json:"P2PFluxTotal,omitempty" name:"P2PFluxTotal"`

		// X-P2P峰值带宽。 单位bps
	// 注意：此字段可能返回 null，表示取不到有效值。
		P2PPeakValue *float64 `json:"P2PPeakValue,omitempty" name:"P2PPeakValue"`

		// RTMP推流路数 ( 直播推流)
	// 注意：此字段可能返回 null，表示取不到有效值。
		LivePushTotal *int64 `json:"LivePushTotal,omitempty" name:"LivePushTotal"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatisticSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子分组详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		GroupList []*GroupItem `json:"GroupList,omitempty" name:"GroupList"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionStatusRequest struct {
	*tchttp.BaseRequest

	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

func (r *DescribeSubscriptionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscriptionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubscriptionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备GB28181报警订阅状态 1：未开启订阅；2：已开启订阅
		AlarmStatus *int64 `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscriptionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoListRequest struct {
	*tchttp.BaseRequest

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 开始时间戳，秒级
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳，秒级
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 开始录制范围 开始
	StartRecordTime *int64 `json:"StartRecordTime,omitempty" name:"StartRecordTime"`

	// 开始录制范围 结束
	EndRecordTime *int64 `json:"EndRecordTime,omitempty" name:"EndRecordTime"`

	// 过期时间范围 开始
	StartExpireTime *int64 `json:"StartExpireTime,omitempty" name:"StartExpireTime"`

	// 过期时间范围 结束
	EndExpireTime *int64 `json:"EndExpireTime,omitempty" name:"EndExpireTime"`

	// 文件大小范围 开始 单位byte
	StartFileSize *int64 `json:"StartFileSize,omitempty" name:"StartFileSize"`

	// 文件大小范围 结束 单位byte
	EndFileSize *int64 `json:"EndFileSize,omitempty" name:"EndFileSize"`

	// 录制状态 99: 录制方已经回写状态 1: 开始录制了，等待回写 2: 已经到了时间模板的停止时间，在等待录制方回写
	IsRecording *int64 `json:"IsRecording,omitempty" name:"IsRecording"`

	// 通道ID默认必传
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 场景ID
	SceneId *int64 `json:"SceneId,omitempty" name:"SceneId"`

	// 告警ID
	WarnId *int64 `json:"WarnId,omitempty" name:"WarnId"`

	// 录制类型 1: 联动计划录制 2: 告警录制
	RecordType []*int64 `json:"RecordType,omitempty" name:"RecordType"`
}

func (r *DescribeVideoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DeviceId")
	delete(f, "StartRecordTime")
	delete(f, "EndRecordTime")
	delete(f, "StartExpireTime")
	delete(f, "EndExpireTime")
	delete(f, "StartFileSize")
	delete(f, "EndFileSize")
	delete(f, "IsRecording")
	delete(f, "ChannelId")
	delete(f, "PlanId")
	delete(f, "SceneId")
	delete(f, "WarnId")
	delete(f, "RecordType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoListRequest has unknown keys!", "")
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
		RecordList []*RecordTaskItem `json:"RecordList,omitempty" name:"RecordList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVideoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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

	// 通道唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

type GetRecordDatesByDevRequest struct {
	*tchttp.BaseRequest

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制量，默认200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 1: 云端录制 2: 本地录制
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *GetRecordDatesByDevRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRecordDatesByDevRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ChannelId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRecordDatesByDevRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRecordDatesByDevResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 含有录像文件的日期列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Dates []*string `json:"Dates,omitempty" name:"Dates"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRecordDatesByDevResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRecordPlanByDevRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRecordPlanByDevRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRecordPlanByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRecordPlanByIdRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRecordPlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRecordPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRecordPlansResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录制计划详情·列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Plans []*RecordPlanItem `json:"Plans,omitempty" name:"Plans"`

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

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTimeTemplateByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTimeTemplateByIdRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTimeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTimeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetTimeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 时间模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Templates []*TimeTemplateItem `json:"Templates,omitempty" name:"Templates"`

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

// FromJsonString It is highly **NOT** recommended to use this function
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

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 1: 云端录制 2: 本地录制
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *GetVideoListByConRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "ChannelId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetVideoListByConRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetVideoListByConResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录像详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		VideoList []*RecordTaskItem `json:"VideoList,omitempty" name:"VideoList"`

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

// FromJsonString It is highly **NOT** recommended to use this function
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

	// 设备接入协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
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

type LiveChannelInfo struct {

	// 频道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

	// 频道名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveChannelName *string `json:"LiveChannelName,omitempty" name:"LiveChannelName"`

	// 频道类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveChannelType *int64 `json:"LiveChannelType,omitempty" name:"LiveChannelType"`

	// 通道直播状态：1: 未推流，2: 推流中
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveStatus *int64 `json:"LiveStatus,omitempty" name:"LiveStatus"`

	// 推流地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PushStreamAddress *string `json:"PushStreamAddress,omitempty" name:"PushStreamAddress"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type LiveChannelItem struct {

	// 频道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 频道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`
}

type LiveRecordItem struct {

	// 录制文件自增ID
	IntID *int64 `json:"IntID,omitempty" name:"IntID"`

	// 直播频道ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

	// 过期时间
	ExpectDeleteTime *int64 `json:"ExpectDeleteTime,omitempty" name:"ExpectDeleteTime"`

	// 录制时长
	RecordTimeLen *int64 `json:"RecordTimeLen,omitempty" name:"RecordTimeLen"`

	// 文件大小
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 录制文件url
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// 录制计划ID
	RecordPlanId *string `json:"RecordPlanId,omitempty" name:"RecordPlanId"`

	// 录制开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 录制结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type LiveRecordPlanItem struct {

	// 计划ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 计划名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`
}

type MessageForward struct {

	// 配置ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`

	// 用户Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// json数组， 转发类型 1: 告警 2:GPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageType *string `json:"MessageType,omitempty" name:"MessageType"`

	// 区域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 区域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// kafka topic id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// topic 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type ModifyBindPlanLiveChannelRequest struct {
	*tchttp.BaseRequest

	// 直播录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 1: 绑定 2: 解绑
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 直播频道ID列表
	LiveChannelIds []*string `json:"LiveChannelIds,omitempty" name:"LiveChannelIds"`
}

func (r *ModifyBindPlanLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBindPlanLiveChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "Type")
	delete(f, "LiveChannelIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBindPlanLiveChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBindPlanLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBindPlanLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBindPlanLiveChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "NickName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceDataRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveChannelRequest struct {
	*tchttp.BaseRequest

	// 直播频道ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

	// 直播频道名
	LiveChannelName *string `json:"LiveChannelName,omitempty" name:"LiveChannelName"`
}

func (r *ModifyLiveChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LiveChannelId")
	delete(f, "LiveChannelName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveRecordPlanRequest struct {
	*tchttp.BaseRequest

	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 录制计划名
	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`

	// 时间模板ID，固定直播时为必填
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *ModifyLiveRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "PlanName")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveRecordPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveVideoRequest struct {
	*tchttp.BaseRequest

	// 视频ID 列表, 大小限制(100)
	IntIDs []*int64 `json:"IntIDs,omitempty" name:"IntIDs"`

	// 过期时间 秒 (-1: 为永不过期)
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *ModifyLiveVideoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveVideoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IntIDs")
	delete(f, "ExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLiveVideoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLiveVideoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLiveVideoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLiveVideoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMessageForwardRequest struct {
	*tchttp.BaseRequest

	// 配置ID
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`

	// json数组， 转发类型 1: 告警 2:GPS
	MessageType *string `json:"MessageType,omitempty" name:"MessageType"`
}

func (r *ModifyMessageForwardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMessageForwardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IntId")
	delete(f, "MessageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMessageForwardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMessageForwardResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMessageForwardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMessageForwardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscriptionStatusRequest struct {
	*tchttp.BaseRequest

	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 订阅状态 1：关闭订阅 2：开启订阅
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 订阅类型 Alarm:告警订阅 Catalog:目录订阅 MobilePosition:移动位置订阅
	SubscriptionItem *string `json:"SubscriptionItem,omitempty" name:"SubscriptionItem"`
}

func (r *ModifySubscriptionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscriptionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Status")
	delete(f, "SubscriptionItem")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubscriptionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubscriptionStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscriptionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscriptionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVideoInfoRequest struct {
	*tchttp.BaseRequest

	// 视频ID列表长度限制100内
	InitIDs []*int64 `json:"InitIDs,omitempty" name:"InitIDs"`

	// 过期时间 时间戳 -1: 永不过期 0: 无效值
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *ModifyVideoInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVideoInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InitIDs")
	delete(f, "ExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVideoInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVideoInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVideoInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVideoInfoResponse) FromJsonString(s string) error {
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
	Devices []*DeviceItem `json:"Devices,omitempty" name:"Devices"`
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

	// 场景ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneId *int64 `json:"SceneId,omitempty" name:"SceneId"`

	// 告警ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarnId *int64 `json:"WarnId,omitempty" name:"WarnId"`

	// 录制id，NVR下属设备有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
}

type SceneItem struct {

	// 场景ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`

	// 用户UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 场景名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`

	// 触发规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneTrigger *string `json:"SceneTrigger,omitempty" name:"SceneTrigger"`

	// 录制时长 秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordDuration *int64 `json:"RecordDuration,omitempty" name:"RecordDuration"`

	// 存储时长 天
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoreDuration *int64 `json:"StoreDuration,omitempty" name:"StoreDuration"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
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

type StreamAddress struct {

	// 流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// rtsp流地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RtspAddr *string `json:"RtspAddr,omitempty" name:"RtspAddr"`

	// rtmp流地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RtmpAddr *string `json:"RtmpAddr,omitempty" name:"RtmpAddr"`

	// hls流地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	HlsAddr *string `json:"HlsAddr,omitempty" name:"HlsAddr"`

	// flv流地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlvAddr *string `json:"FlvAddr,omitempty" name:"FlvAddr"`
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
	TimeTemplateSpecs []*TimeTemplateSpec `json:"TimeTemplateSpecs,omitempty" name:"TimeTemplateSpecs"`
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDeviceGroupRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDevicePassWordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PassWord")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDevicePassWordRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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
	Devices []*DeviceItem `json:"Devices,omitempty" name:"Devices"`

	// 是否更新绑定此录制计划的设备列表
	// 0 - 不更新
	// 1 - 更新，如果Devices参数为空则清空设备列表，Devices不为空则全量更新设备列表
	IsModifyDevices *int64 `json:"IsModifyDevices,omitempty" name:"IsModifyDevices"`
}

func (r *UpdateRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRecordPlanRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
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
	TimeTemplateSpecs []*TimeTemplateSpec `json:"TimeTemplateSpecs,omitempty" name:"TimeTemplateSpecs"`
}

func (r *UpdateTimeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTimeTemplateRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTimeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
