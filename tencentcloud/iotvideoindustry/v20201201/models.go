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

type AbnormalEvents struct {
	// 对应查询日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info []*AbnormalEventsInfo `json:"Info,omitempty" name:"Info"`
}

type AbnormalEventsInfo struct {
	// 类型值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *int64 `json:"Key,omitempty" name:"Key"`

	// 类型总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

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

// Predefined struct for user
type BindGroupDevicesRequestParams struct {
	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 设备唯一标识列表
	DeviceList []*string `json:"DeviceList,omitempty" name:"DeviceList"`
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

// Predefined struct for user
type BindGroupDevicesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindGroupDevicesResponse struct {
	*tchttp.BaseResponse
	Response *BindGroupDevicesResponseParams `json:"Response"`
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

type ChannelDetail struct {
	// 通道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 通道类型 0：未知；1：视频通道；2：音频通道；3：告警通道
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelType *int64 `json:"ChannelType,omitempty" name:"ChannelType"`

	// 20位国标通道编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelCode *string `json:"ChannelCode,omitempty" name:"ChannelCode"`

	// 通道扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInformation *string `json:"ExtraInformation,omitempty" name:"ExtraInformation"`

	// 通道在线状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 通道是否存在录像标识 0：无录像；1：有录像
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRecord *int64 `json:"IsRecord,omitempty" name:"IsRecord"`

	// 通道所属设备唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道所属虚拟组织的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessGroupId *string `json:"BusinessGroupId,omitempty" name:"BusinessGroupId"`
}

type ChannelItem struct {
	// 设备唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

// Predefined struct for user
type ControlChannelLocalRecordRequestParams struct {
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 流Id，流的唯一标识
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// 控制参数，转义的json字符串
	// 
	// 目前支持的command：
	// "Command": "{"Action":"PAUSE"}" 暂停
	// "Command": "{"Action":"PLAY"}" 暂停恢复
	// "Command": "{"Action":"PLAY","Offset":"15"}" 基于文件起始时间点的位置偏移，单位秒
	Command *string `json:"Command,omitempty" name:"Command"`
}

type ControlChannelLocalRecordRequest struct {
	*tchttp.BaseRequest
	
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 流Id，流的唯一标识
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// 控制参数，转义的json字符串
	// 
	// 目前支持的command：
	// "Command": "{"Action":"PAUSE"}" 暂停
	// "Command": "{"Action":"PLAY"}" 暂停恢复
	// "Command": "{"Action":"PLAY","Offset":"15"}" 基于文件起始时间点的位置偏移，单位秒
	Command *string `json:"Command,omitempty" name:"Command"`
}

func (r *ControlChannelLocalRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlChannelLocalRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ChannelId")
	delete(f, "StreamId")
	delete(f, "Command")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlChannelLocalRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlChannelLocalRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ControlChannelLocalRecordResponse struct {
	*tchttp.BaseResponse
	Response *ControlChannelLocalRecordResponseParams `json:"Response"`
}

func (r *ControlChannelLocalRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlChannelLocalRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlChannelPTZRequestParams struct {
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

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

type ControlChannelPTZRequest struct {
	*tchttp.BaseRequest
	
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

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

func (r *ControlChannelPTZRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlChannelPTZRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ChannelId")
	delete(f, "Command")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlChannelPTZRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlChannelPTZResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ControlChannelPTZResponse struct {
	*tchttp.BaseResponse
	Response *ControlChannelPTZResponseParams `json:"Response"`
}

func (r *ControlChannelPTZResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlChannelPTZResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDevicePTZRequestParams struct {
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

// Predefined struct for user
type ControlDevicePTZResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ControlDevicePTZResponse struct {
	*tchttp.BaseResponse
	Response *ControlDevicePTZResponseParams `json:"Response"`
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

// Predefined struct for user
type ControlHomePositionRequestParams struct {
	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 看守位使能 0-停用看守位 1-启用看守位
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 预置位编码 范围1-8，启用看守位时必填
	PresetId *int64 `json:"PresetId,omitempty" name:"PresetId"`

	// 看守位自动归位时间， 启用看守位时必填
	ResetTime *int64 `json:"ResetTime,omitempty" name:"ResetTime"`
}

type ControlHomePositionRequest struct {
	*tchttp.BaseRequest
	
	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 看守位使能 0-停用看守位 1-启用看守位
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 预置位编码 范围1-8，启用看守位时必填
	PresetId *int64 `json:"PresetId,omitempty" name:"PresetId"`

	// 看守位自动归位时间， 启用看守位时必填
	ResetTime *int64 `json:"ResetTime,omitempty" name:"ResetTime"`
}

func (r *ControlHomePositionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlHomePositionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "DeviceId")
	delete(f, "Enable")
	delete(f, "PresetId")
	delete(f, "ResetTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlHomePositionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlHomePositionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ControlHomePositionResponse struct {
	*tchttp.BaseResponse
	Response *ControlHomePositionResponseParams `json:"Response"`
}

func (r *ControlHomePositionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlHomePositionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlPresetRequestParams struct {
	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 控制命令：
	// Set-设置当前位置为预置位
	// Del-删除指定的预置位
	// Call-调用指定的预置位
	Command *string `json:"Command,omitempty" name:"Command"`

	// 预置位编码 范围1-8
	PresetId *int64 `json:"PresetId,omitempty" name:"PresetId"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

type ControlPresetRequest struct {
	*tchttp.BaseRequest
	
	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 控制命令：
	// Set-设置当前位置为预置位
	// Del-删除指定的预置位
	// Call-调用指定的预置位
	Command *string `json:"Command,omitempty" name:"Command"`

	// 预置位编码 范围1-8
	PresetId *int64 `json:"PresetId,omitempty" name:"PresetId"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

func (r *ControlPresetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlPresetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "Command")
	delete(f, "PresetId")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlPresetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlPresetResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ControlPresetResponse struct {
	*tchttp.BaseResponse
	Response *ControlPresetResponseParams `json:"Response"`
}

func (r *ControlPresetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlPresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlRecordStreamRequestParams struct {
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

// Predefined struct for user
type ControlRecordStreamResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ControlRecordStreamResponse struct {
	*tchttp.BaseResponse
	Response *ControlRecordStreamResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateDeviceGroupRequestParams struct {
	// 分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 父分组ID
	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`

	// 分组描述
	GroupDescribe *string `json:"GroupDescribe,omitempty" name:"GroupDescribe"`
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

// Predefined struct for user
type CreateDeviceGroupResponseParams struct {
	// 响应结果，“OK”为成功，其他为失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeviceGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateDeviceRequestParams struct {
	// 设备名称
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 设备密码
	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`

	// 设备类型，1：国标VMS设备(公有云不支持此类型)，2：国标IPC设备，3：国标NVR设备，9：智能告警设备(公有云不支持此类型)
	DeviceType *int64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 设备需要绑定的分组ID，参数为空则默认绑定到根分组
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type CreateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备名称
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 设备密码
	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`

	// 设备类型，1：国标VMS设备(公有云不支持此类型)，2：国标IPC设备，3：国标NVR设备，9：智能告警设备(公有云不支持此类型)
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

// Predefined struct for user
type CreateDeviceResponseParams struct {
	// 设备编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceCode *string `json:"DeviceCode,omitempty" name:"DeviceCode"`

	// 设备唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 设备虚拟组信息，仅在创建NVR时返回该值
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualGroupId *string `json:"VirtualGroupId,omitempty" name:"VirtualGroupId"`

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
type CreateLiveChannelRequestParams struct {
	// 直播频道名称
	LiveChannelName *string `json:"LiveChannelName,omitempty" name:"LiveChannelName"`

	// 直播频道类型 1：固定直播；2：移动直播
	LiveChannelType *int64 `json:"LiveChannelType,omitempty" name:"LiveChannelType"`
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

// Predefined struct for user
type CreateLiveChannelResponseParams struct {
	// 直播频道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

	// 直播频道推流地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PushStreamAddress *string `json:"PushStreamAddress,omitempty" name:"PushStreamAddress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveChannelResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateLiveRecordPlanRequestParams struct {
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

// Predefined struct for user
type CreateLiveRecordPlanResponseParams struct {
	// 录制计划名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLiveRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateLiveRecordPlanResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateMessageForwardRequestParams struct {
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

// Predefined struct for user
type CreateMessageForwardResponseParams struct {
	// 配置ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMessageForwardResponse struct {
	*tchttp.BaseResponse
	Response *CreateMessageForwardResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRecordPlanRequestParams struct {
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

// Predefined struct for user
type CreateRecordPlanResponseParams struct {
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateRecordPlanResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRecordingPlanRequestParams struct {
	// 计划名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间模板ID
	TimeTemplateId *string `json:"TimeTemplateId,omitempty" name:"TimeTemplateId"`

	// 该录制计划绑定的通道列表
	Channels []*ChannelItem `json:"Channels,omitempty" name:"Channels"`

	// 存储周期(天)；默认存储30天
	RecordStorageTime *int64 `json:"RecordStorageTime,omitempty" name:"RecordStorageTime"`
}

type CreateRecordingPlanRequest struct {
	*tchttp.BaseRequest
	
	// 计划名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间模板ID
	TimeTemplateId *string `json:"TimeTemplateId,omitempty" name:"TimeTemplateId"`

	// 该录制计划绑定的通道列表
	Channels []*ChannelItem `json:"Channels,omitempty" name:"Channels"`

	// 存储周期(天)；默认存储30天
	RecordStorageTime *int64 `json:"RecordStorageTime,omitempty" name:"RecordStorageTime"`
}

func (r *CreateRecordingPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordingPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "TimeTemplateId")
	delete(f, "Channels")
	delete(f, "RecordStorageTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecordingPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecordingPlanResponseParams struct {
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRecordingPlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateRecordingPlanResponseParams `json:"Response"`
}

func (r *CreateRecordingPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordingPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSceneRequestParams struct {
	// 场景名称
	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`

	// 场景触发规则
	SceneTrigger *string `json:"SceneTrigger,omitempty" name:"SceneTrigger"`

	// 录制时长 (秒)
	RecordDuration *int64 `json:"RecordDuration,omitempty" name:"RecordDuration"`

	// 录像存储时长(天)
	StoreDuration *int64 `json:"StoreDuration,omitempty" name:"StoreDuration"`

	// 设备列表(不推荐使用)
	Devices []*DeviceItem `json:"Devices,omitempty" name:"Devices"`

	// 通道列表
	Channels []*ChannelItem `json:"Channels,omitempty" name:"Channels"`
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

	// 设备列表(不推荐使用)
	Devices []*DeviceItem `json:"Devices,omitempty" name:"Devices"`

	// 通道列表
	Channels []*ChannelItem `json:"Channels,omitempty" name:"Channels"`
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
	delete(f, "Channels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSceneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSceneResponseParams struct {
	// 场景ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSceneResponse struct {
	*tchttp.BaseResponse
	Response *CreateSceneResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateTimeTemplateRequestParams struct {
	// 时间模板名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否为每周全时录制（即7*24h录制），0：非全时录制，1；全时录制，默认0
	IsAllWeek *int64 `json:"IsAllWeek,omitempty" name:"IsAllWeek"`

	// 当IsAllWeek为0时必选，用于描述模板的各个时间片段
	TimeTemplateSpecs []*TimeTemplateSpec `json:"TimeTemplateSpecs,omitempty" name:"TimeTemplateSpecs"`
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

// Predefined struct for user
type CreateTimeTemplateResponseParams struct {
	// 时间模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTimeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateTimeTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteChannelRequestParams struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
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

// Predefined struct for user
type DeleteChannelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteChannelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteChannelResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteDeviceGroupRequestParams struct {
	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
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

// Predefined struct for user
type DeleteDeviceGroupResponseParams struct {
	// 响应结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeviceGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteDeviceRequestParams struct {
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
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

// Predefined struct for user
type DeleteDeviceResponseParams struct {
	// 操作结果 OK-成功； 其他-失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

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
type DeleteLiveChannelRequestParams struct {
	// 直播频道ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`
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

// Predefined struct for user
type DeleteLiveChannelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveChannelResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveRecordPlanRequestParams struct {
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
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

// Predefined struct for user
type DeleteLiveRecordPlanResponseParams struct {
	// 删除状态描述
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLiveRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveRecordPlanResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLiveVideoListRequestParams struct {
	// 视频ID 列表, 大小限制(100)
	IntIDs []*uint64 `json:"IntIDs,omitempty" name:"IntIDs"`
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

// Predefined struct for user
type DeleteLiveVideoListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLiveVideoListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLiveVideoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteMessageForwardRequestParams struct {
	// 配置ID
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`
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

// Predefined struct for user
type DeleteMessageForwardResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMessageForwardResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMessageForwardResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRecordPlanRequestParams struct {
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
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

// Predefined struct for user
type DeleteRecordPlanResponseParams struct {
	// 操作结果，OK：成功，其他：失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordPlanResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRecordingPlanRequestParams struct {
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
}

type DeleteRecordingPlanRequest struct {
	*tchttp.BaseRequest
	
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
}

func (r *DeleteRecordingPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordingPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordingPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordingPlanResponseParams struct {
	// 操作结果，OK：成功，其他：失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRecordingPlanResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordingPlanResponseParams `json:"Response"`
}

func (r *DeleteRecordingPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordingPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSceneRequestParams struct {
	// 场景ID
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`
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

// Predefined struct for user
type DeleteSceneResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSceneResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSceneResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteTimeTemplateRequestParams struct {
	// 时间模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type DeleteTimeTemplateResponseParams struct {
	// 操作结果，OK：成功，其他：失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTimeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTimeTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteVideoListRequestParams struct {
	// 视频ID列表长度限制100内
	InitIDs []*int64 `json:"InitIDs,omitempty" name:"InitIDs"`
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

// Predefined struct for user
type DeleteVideoListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteVideoListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVideoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteWarningRequestParams struct {
	// 告警ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 告警索引
	Index *string `json:"Index,omitempty" name:"Index"`
}

type DeleteWarningRequest struct {
	*tchttp.BaseRequest
	
	// 告警ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 告警索引
	Index *string `json:"Index,omitempty" name:"Index"`
}

func (r *DeleteWarningRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWarningRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Index")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWarningRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWarningResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteWarningResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWarningResponseParams `json:"Response"`
}

func (r *DeleteWarningResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWarningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalEventsRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeAbnormalEventsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAbnormalEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAbnormalEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalEventsResponseParams struct {
	// 异动事件走势列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*AbnormalEvents `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAbnormalEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAbnormalEventsResponseParams `json:"Response"`
}

func (r *DescribeAbnormalEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllDeviceListRequestParams struct {
	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制，默认200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 设备名称，需要模糊匹配设备名称时为必填
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// DeviceId列表，需要精确查找设备时为必填
	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds"`

	// 设备类型过滤，设备类型，1：国标VMS设备(公有云不支持此类型)，2：国标IPC设备，3：国标NVR设备，9：智能告警设备(公有云不支持此类型)
	DeviceTypes []*int64 `json:"DeviceTypes,omitempty" name:"DeviceTypes"`
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

	// 设备类型过滤，设备类型，1：国标VMS设备(公有云不支持此类型)，2：国标IPC设备，3：国标NVR设备，9：智能告警设备(公有云不支持此类型)
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

// Predefined struct for user
type DescribeAllDeviceListResponseParams struct {
	// 设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 设备详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Devices []*AllDeviceInfo `json:"Devices,omitempty" name:"Devices"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAllDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllDeviceListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBindSceneChannelsRequestParams struct {
	// 条数限制最大不能超过1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 场景ID
	SceneId *int64 `json:"SceneId,omitempty" name:"SceneId"`

	// 偏移值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeBindSceneChannelsRequest struct {
	*tchttp.BaseRequest
	
	// 条数限制最大不能超过1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 场景ID
	SceneId *int64 `json:"SceneId,omitempty" name:"SceneId"`

	// 偏移值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBindSceneChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindSceneChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "SceneId")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindSceneChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindSceneChannelsResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 通道列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*ChannelItem `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBindSceneChannelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindSceneChannelsResponseParams `json:"Response"`
}

func (r *DescribeBindSceneChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindSceneChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindSceneDevicesRequestParams struct {
	// 场景ID
	SceneId *int64 `json:"SceneId,omitempty" name:"SceneId"`

	// 偏移值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 条数限制最大不能超过1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeBindSceneDevicesResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 设备列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DeviceItem `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBindSceneDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindSceneDevicesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeChannelLiveStreamURLRequestParams struct {
	// 设备唯一标识，必填参数
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识（接口升级字段为必填），必填参数
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

type DescribeChannelLiveStreamURLRequest struct {
	*tchttp.BaseRequest
	
	// 设备唯一标识，必填参数
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识（接口升级字段为必填），必填参数
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

func (r *DescribeChannelLiveStreamURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelLiveStreamURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChannelLiveStreamURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelLiveStreamURLResponseParams struct {
	// 设备实时流地址列表
	Data *DescribeDeviceStreamsData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeChannelLiveStreamURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChannelLiveStreamURLResponseParams `json:"Response"`
}

func (r *DescribeChannelLiveStreamURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelLiveStreamURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelLocalRecordURLRequestParams struct {
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 录像文件Id，通过获取本地录像返回
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// UNIX 时间戳，30天内，URL失效时间，如180s无人观看此流则URL提前失效
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 录像文件推送的开始时间，需要在RecordId参数起始时间内，可以通过此参数控制回放流起始点
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 录像文件推送的结束时间，需要在RecordId参数起始时间内，可以通过此参数控制回放流起始点
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeChannelLocalRecordURLRequest struct {
	*tchttp.BaseRequest
	
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 录像文件Id，通过获取本地录像返回
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// UNIX 时间戳，30天内，URL失效时间，如180s无人观看此流则URL提前失效
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 录像文件推送的开始时间，需要在RecordId参数起始时间内，可以通过此参数控制回放流起始点
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 录像文件推送的结束时间，需要在RecordId参数起始时间内，可以通过此参数控制回放流起始点
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeChannelLocalRecordURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelLocalRecordURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ChannelId")
	delete(f, "RecordId")
	delete(f, "ExpireTime")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChannelLocalRecordURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelLocalRecordURLResponseParams struct {
	// 结果
	Data *DescribeRecordStreamData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeChannelLocalRecordURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChannelLocalRecordURLResponseParams `json:"Response"`
}

func (r *DescribeChannelLocalRecordURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelLocalRecordURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelStreamURLRequestParams struct {
	// 设备唯一标识，必填参数
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 流地址失效时间，固定值填写0，其他参数无效，必填参数
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 通道唯一标识（接口升级字段为必填），必填参数
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

type DescribeChannelStreamURLRequest struct {
	*tchttp.BaseRequest
	
	// 设备唯一标识，必填参数
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 流地址失效时间，固定值填写0，其他参数无效，必填参数
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 通道唯一标识（接口升级字段为必填），必填参数
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

func (r *DescribeChannelStreamURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelStreamURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ExpireTime")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChannelStreamURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelStreamURLResponseParams struct {
	// 设备实时流地址列表
	Data *DescribeDeviceStreamsData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeChannelStreamURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChannelStreamURLResponseParams `json:"Response"`
}

func (r *DescribeChannelStreamURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelStreamURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelsByLiveRecordPlanRequestParams struct {
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeChannelsByLiveRecordPlanResponseParams struct {
	// 总个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 通道详情数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveChannels []*LiveChannelItem `json:"LiveChannels,omitempty" name:"LiveChannels"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeChannelsByLiveRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChannelsByLiveRecordPlanResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeChannelsRequestParams struct {
	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 限制，默认0
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 通道类型  0: 未知类型 1: 视频通道 2:  音频通道 3: 告警通道
	ChannelTypes []*uint64 `json:"ChannelTypes,omitempty" name:"ChannelTypes"`

	// 录制计划ID， 当为"null"值时未绑定录制计划
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 告警联动场景ID， 当为 -1 值时未绑定场景
	SceneId *int64 `json:"SceneId,omitempty" name:"SceneId"`
}

type DescribeChannelsRequest struct {
	*tchttp.BaseRequest
	
	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 限制，默认0
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 通道类型  0: 未知类型 1: 视频通道 2:  音频通道 3: 告警通道
	ChannelTypes []*uint64 `json:"ChannelTypes,omitempty" name:"ChannelTypes"`

	// 录制计划ID， 当为"null"值时未绑定录制计划
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 告警联动场景ID， 当为 -1 值时未绑定场景
	SceneId *int64 `json:"SceneId,omitempty" name:"SceneId"`
}

func (r *DescribeChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ChannelTypes")
	delete(f, "PlanId")
	delete(f, "SceneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelsResponseParams struct {
	// 通道总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 通道详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Channels []*ChannelDetail `json:"Channels,omitempty" name:"Channels"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeChannelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChannelsResponseParams `json:"Response"`
}

func (r *DescribeChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurrentDeviceDataRequestParams struct {

}

type DescribeCurrentDeviceDataRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCurrentDeviceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurrentDeviceDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCurrentDeviceDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurrentDeviceDataResponseParams struct {
	// 通道数
	Channels *int64 `json:"Channels,omitempty" name:"Channels"`

	// 设备数
	Devices *int64 `json:"Devices,omitempty" name:"Devices"`

	// 在线通道数
	OnlineChannels *int64 `json:"OnlineChannels,omitempty" name:"OnlineChannels"`

	// 在线设备数
	OnlineDevices *int64 `json:"OnlineDevices,omitempty" name:"OnlineDevices"`

	// 正在录制通道数
	RecordingChannels *int64 `json:"RecordingChannels,omitempty" name:"RecordingChannels"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCurrentDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCurrentDeviceDataResponseParams `json:"Response"`
}

func (r *DescribeCurrentDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurrentDeviceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceEventRequestParams struct {
	// 开始时间，秒级时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，秒级时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 默认为全部 事件类型 1:注册 2:心跳 4:录制异常 5:播放异常 6:流中断
	EventTypes []*int64 `json:"EventTypes,omitempty" name:"EventTypes"`

	// 偏移值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// limit限制值
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDeviceEventRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，秒级时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，秒级时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 默认为全部 事件类型 1:注册 2:心跳 4:录制异常 5:播放异常 6:流中断
	EventTypes []*int64 `json:"EventTypes,omitempty" name:"EventTypes"`

	// 偏移值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// limit限制值
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDeviceEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DeviceId")
	delete(f, "EventTypes")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceEventResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Events []*Events `json:"Events,omitempty" name:"Events"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceEventResponseParams `json:"Response"`
}

func (r *DescribeDeviceEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceGroupRequestParams struct {
	// 设备唯一标识列表
	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds"`
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

// Predefined struct for user
type DescribeDeviceGroupResponseParams struct {
	// 设备所在分组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevGroups []*DevGroupInfo `json:"DevGroups,omitempty" name:"DevGroups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDeviceListRequestParams struct {
	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制，默认200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 设备名前缀
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 设备类型，1：国标VMS设备(公有云不支持此类型)，2：国标IPC设备，3：国标NVR设备，9：智能告警设备(公有云不支持此类型)
	DeviceTypes []*int64 `json:"DeviceTypes,omitempty" name:"DeviceTypes"`
}

type DescribeDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制，默认200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 设备名前缀
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 设备类型，1：国标VMS设备(公有云不支持此类型)，2：国标IPC设备，3：国标NVR设备，9：智能告警设备(公有云不支持此类型)
	DeviceTypes []*int64 `json:"DeviceTypes,omitempty" name:"DeviceTypes"`
}

func (r *DescribeDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NickName")
	delete(f, "DeviceTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceListResponseParams struct {
	// 设备总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 设备详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Devices []*AllDeviceInfo `json:"Devices,omitempty" name:"Devices"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceListResponseParams `json:"Response"`
}

func (r *DescribeDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceMonitorDataRequestParams struct {
	// 开始时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 类型 支持 OnlineChannels/OnlineDevices/RecordingChannels
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间粒度 目前只支持 1h
	TimesSpec *string `json:"TimesSpec,omitempty" name:"TimesSpec"`
}

type DescribeDeviceMonitorDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 类型 支持 OnlineChannels/OnlineDevices/RecordingChannels
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时间粒度 目前只支持 1h
	TimesSpec *string `json:"TimesSpec,omitempty" name:"TimesSpec"`
}

func (r *DescribeDeviceMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceMonitorDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "TimesSpec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceMonitorDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceMonitorDataResponseParams struct {
	// 查询设备统计monitor信息列表
	Data []*DeviceMonitorValue `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceMonitorDataResponseParams `json:"Response"`
}

func (r *DescribeDeviceMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePassWordRequestParams struct {
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
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

// Predefined struct for user
type DescribeDevicePassWordResponseParams struct {
	// 设备密码
	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDevicePassWordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicePassWordResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDeviceRequestParams struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

type DescribeDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID
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
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceResponseParams struct {
	// 设备详情信息
	Device *AllDeviceInfo `json:"Device,omitempty" name:"Device"`

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

// Predefined struct for user
type DescribeDeviceStreamsRequestParams struct {
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 流地址失效时间
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 通道唯一标识（接口升级字段为必填）
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
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

// Predefined struct for user
type DescribeDeviceStreamsResponseParams struct {
	// 设备实时流地址列表
	Data *DescribeDeviceStreamsData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceStreamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceStreamsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeGroupByIdRequestParams struct {
	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
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

// Predefined struct for user
type DescribeGroupByIdResponseParams struct {
	// 分组信息详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Group *GroupItem `json:"Group,omitempty" name:"Group"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupByIdResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeGroupByPathRequestParams struct {
	// 分组路径，格式为/aaa(/bbb/ccc)
	GroupPath *string `json:"GroupPath,omitempty" name:"GroupPath"`
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

// Predefined struct for user
type DescribeGroupByPathResponseParams struct {
	// 分组信息详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Group *GroupItem `json:"Group,omitempty" name:"Group"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupByPathResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupByPathResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeGroupDevicesRequestParams struct {
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

	// 当Group是普通组的时候，支持根据DeviceTypes筛选类型，
	//  设备类型，1：国标VMS设备(公有云不支持此类型)，2：国标IPC设备，3：国标NVR设备，9：智能告警设备(公有云不支持此类型)
	DeviceTypes []*int64 `json:"DeviceTypes,omitempty" name:"DeviceTypes"`
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

	// 当Group是普通组的时候，支持根据DeviceTypes筛选类型，
	//  设备类型，1：国标VMS设备(公有云不支持此类型)，2：国标IPC设备，3：国标NVR设备，9：智能告警设备(公有云不支持此类型)
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

// Predefined struct for user
type DescribeGroupDevicesResponseParams struct {
	// 分组绑定的设备数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 设备详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceList []*GroupDeviceItem `json:"DeviceList,omitempty" name:"DeviceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupDevicesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeGroupsRequestParams struct {
	// 分组ID列表
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
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

// Predefined struct for user
type DescribeGroupsResponseParams struct {
	// 分组详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*GroupInfo `json:"Groups,omitempty" name:"Groups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeIPCChannelsRequestParams struct {
	// 偏移量，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制，默认0
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道类型  0: 未知类型 1: 视频通道 2:  音频通道 3: 告警通道
	ChannelTypes []*uint64 `json:"ChannelTypes,omitempty" name:"ChannelTypes"`
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

// Predefined struct for user
type DescribeIPCChannelsResponseParams struct {
	// 通道总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 通道详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceList []*GroupDeviceItem `json:"DeviceList,omitempty" name:"DeviceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIPCChannelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPCChannelsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveChannelListRequestParams struct {
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

// Predefined struct for user
type DescribeLiveChannelListResponseParams struct {
	// 频道总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 频道信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveChannels []*LiveChannelInfo `json:"LiveChannels,omitempty" name:"LiveChannels"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLiveChannelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveChannelListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveChannelRequestParams struct {
	// 频道ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`
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

// Predefined struct for user
type DescribeLiveChannelResponseParams struct {
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
}

type DescribeLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveChannelResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveRecordPlanByIdRequestParams struct {
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
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

// Predefined struct for user
type DescribeLiveRecordPlanByIdResponseParams struct {
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
}

type DescribeLiveRecordPlanByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveRecordPlanByIdResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveRecordPlanIdsRequestParams struct {
	// 时间模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeLiveRecordPlanIdsResponseParams struct {
	// 总个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 计划数组
	Plans []*LiveRecordPlanItem `json:"Plans,omitempty" name:"Plans"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLiveRecordPlanIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveRecordPlanIdsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveStreamRequestParams struct {
	// 频道ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

	// 过期时间 秒级unix时间戳
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type DescribeLiveStreamRequest struct {
	*tchttp.BaseRequest
	
	// 频道ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

	// 过期时间 秒级unix时间戳
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

// Predefined struct for user
type DescribeLiveStreamResponseParams struct {
	// 拉流地址，只有在推流情况下才有
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *StreamAddress `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveStreamResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeLiveVideoListRequestParams struct {
	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 直播频道ID
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

type DescribeLiveVideoListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 直播频道ID
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

// Predefined struct for user
type DescribeLiveVideoListResponseParams struct {
	// 总的条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 录制任务详情数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordList []*LiveRecordItem `json:"RecordList,omitempty" name:"RecordList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLiveVideoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLiveVideoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeMessageForwardRequestParams struct {
	// 配置ID
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`
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

// Predefined struct for user
type DescribeMessageForwardResponseParams struct {
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
}

type DescribeMessageForwardResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMessageForwardResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeMessageForwardsRequestParams struct {
	// 数量限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

// Predefined struct for user
type DescribeMessageForwardsResponseParams struct {
	// 配置总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*MessageForward `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMessageForwardsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMessageForwardsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeMonitorDataByDateRequestParams struct {
	// 开始时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳 最多显示30天数据
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeMonitorDataByDateRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳 最多显示30天数据
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeMonitorDataByDateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorDataByDateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMonitorDataByDateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorDataByDateResponseParams struct {
	// 统计数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*RecordStatistic `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMonitorDataByDateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMonitorDataByDateResponseParams `json:"Response"`
}

func (r *DescribeMonitorDataByDateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorDataByDateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePresetListRequestParams struct {
	// 视频通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

type DescribePresetListRequest struct {
	*tchttp.BaseRequest
	
	// 视频通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

func (r *DescribePresetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePresetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePresetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePresetListResponseParams struct {
	// 预置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*PresetItem `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePresetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePresetListResponseParams `json:"Response"`
}

func (r *DescribePresetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePresetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordDatesByChannelRequestParams struct {
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 1: 云端录制 2: 本地录制
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 限制量，默认200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeRecordDatesByChannelRequest struct {
	*tchttp.BaseRequest
	
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 1: 云端录制 2: 本地录制
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 限制量，默认200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeRecordDatesByChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordDatesByChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ChannelId")
	delete(f, "Type")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordDatesByChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordDatesByChannelResponseParams struct {
	// 含有录像文件的日期列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dates []*string `json:"Dates,omitempty" name:"Dates"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRecordDatesByChannelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordDatesByChannelResponseParams `json:"Response"`
}

func (r *DescribeRecordDatesByChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordDatesByChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordDatesByLiveRequestParams struct {
	// 直播频道ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

	// 分页值，本地录制时参数无效
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制值，本地录制时参数无效
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeRecordDatesByLiveResponseParams struct {
	// 录制日期数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dates []*string `json:"Dates,omitempty" name:"Dates"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRecordDatesByLiveResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordDatesByLiveResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRecordStreamRequestParams struct {
	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 流失效时间，UNIX时间戳，30天内
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 录像文件ID
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 录像流开始时间，当录像文件ID为空时有效，UNIX时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 录像流结束时间，当录像文件iD为空时有效，UNIX时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 通道唯一标识（此接口升级为必填字段）
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`
}

type DescribeRecordStreamRequest struct {
	*tchttp.BaseRequest
	
	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 流失效时间，UNIX时间戳，30天内
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 录像文件ID
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 录像流开始时间，当录像文件ID为空时有效，UNIX时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 录像流结束时间，当录像文件iD为空时有效，UNIX时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 通道唯一标识（此接口升级为必填字段）
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

// Predefined struct for user
type DescribeRecordStreamResponseParams struct {
	// 结果
	Data *DescribeRecordStreamData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRecordStreamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordStreamResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRecordingPlanByIdRequestParams struct {
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
}

type DescribeRecordingPlanByIdRequest struct {
	*tchttp.BaseRequest
	
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
}

func (r *DescribeRecordingPlanByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordingPlanByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordingPlanByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordingPlanByIdResponseParams struct {
	// 录制计划详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Plan *RecordPlanDetail `json:"Plan,omitempty" name:"Plan"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRecordingPlanByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordingPlanByIdResponseParams `json:"Response"`
}

func (r *DescribeRecordingPlanByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordingPlanByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordingPlansRequestParams struct {

}

type DescribeRecordingPlansRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRecordingPlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordingPlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordingPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordingPlansResponseParams struct {
	// 录制计划详情·列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Plans []*RecordPlanDetail `json:"Plans,omitempty" name:"Plans"`

	// 录制计划总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRecordingPlansResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordingPlansResponseParams `json:"Response"`
}

func (r *DescribeRecordingPlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordingPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSIPServerRequestParams struct {

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

// Predefined struct for user
type DescribeSIPServerResponseParams struct {
	// SIP服务器相关配置项
	Data *ServerConfiguration `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSIPServerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSIPServerResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSceneRequestParams struct {
	// 场景ID
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`
}

type DescribeSceneRequest struct {
	*tchttp.BaseRequest
	
	// 场景ID
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`
}

func (r *DescribeSceneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSceneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IntId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSceneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSceneResponseParams struct {
	// 场景ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`

	// 录制时长(秒)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordDuration *int64 `json:"RecordDuration,omitempty" name:"RecordDuration"`

	// 场景名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`

	// 场景触发规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneTrigger *string `json:"SceneTrigger,omitempty" name:"SceneTrigger"`

	// 存储时长 (天)
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoreDuration *int64 `json:"StoreDuration,omitempty" name:"StoreDuration"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 用户Uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSceneResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSceneResponseParams `json:"Response"`
}

func (r *DescribeSceneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSceneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScenesRequestParams struct {
	// 条数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

// Predefined struct for user
type DescribeScenesResponseParams struct {
	// 场景总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 场景列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*SceneItem `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeScenesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScenesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeStatisticDetailsRequestParams struct {
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

// Predefined struct for user
type DescribeStatisticDetailsResponseParams struct {
	// 统计详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*StatisticItem `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStatisticDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStatisticDetailsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeStatisticSummaryRequestParams struct {
	// 指定日期。格式【YYYY-MM-DD】
	Date *string `json:"Date,omitempty" name:"Date"`
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

// Predefined struct for user
type DescribeStatisticSummaryResponseParams struct {
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
}

type DescribeStatisticSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStatisticSummaryResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSubGroupsRequestParams struct {
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

// Predefined struct for user
type DescribeSubGroupsResponseParams struct {
	// 子分组详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupList []*GroupItem `json:"GroupList,omitempty" name:"GroupList"`

	// 子分组总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubGroupsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSubscriptionStatusRequestParams struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
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

// Predefined struct for user
type DescribeSubscriptionStatusResponseParams struct {
	// 设备GB28181报警订阅状态 1：未开启订阅；2：已开启订阅
	AlarmStatus *int64 `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubscriptionStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubscriptionStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeVideoListByChannelRequestParams struct {
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 1: 云端录制 2: 本地录制
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 指定某天。取值【YYYY-MM-DD】
	// 为空时默认查询最近一天的记录
	Date *string `json:"Date,omitempty" name:"Date"`

	// 限制量，默认2000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeVideoListByChannelRequest struct {
	*tchttp.BaseRequest
	
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道唯一标识
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 1: 云端录制 2: 本地录制
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 指定某天。取值【YYYY-MM-DD】
	// 为空时默认查询最近一天的记录
	Date *string `json:"Date,omitempty" name:"Date"`

	// 限制量，默认2000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVideoListByChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoListByChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ChannelId")
	delete(f, "Type")
	delete(f, "Date")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoListByChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoListByChannelResponseParams struct {
	// 录像详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoList []*RecordTaskItem `json:"VideoList,omitempty" name:"VideoList"`

	// 录像总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVideoListByChannelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoListByChannelResponseParams `json:"Response"`
}

func (r *DescribeVideoListByChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoListByChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoListRequestParams struct {
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

// Predefined struct for user
type DescribeVideoListResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 已废弃
	VideoList *RecordTaskItem `json:"VideoList,omitempty" name:"VideoList"`

	// 录像详情列表
	RecordList []*RecordTaskItem `json:"RecordList,omitempty" name:"RecordList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVideoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeWarnModRequestParams struct {

}

type DescribeWarnModRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWarnModRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWarnModRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWarnModRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWarnModResponseParams struct {
	// 告警类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWarnModResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWarnModResponseParams `json:"Response"`
}

func (r *DescribeWarnModResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWarnModResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWarningsRequestParams struct {
	// 1:创建时间倒序 2：创建时间升序 3：level倒序 4：leve升序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 可选设备id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 如果不传则查询所有，取值参见配置
	WarnLevelArray []*int64 `json:"WarnLevelArray,omitempty" name:"WarnLevelArray"`

	// 如果不传则查询所有，取值参见配置
	WarnModeArray []*int64 `json:"WarnModeArray,omitempty" name:"WarnModeArray"`

	// 不传认为是0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 不传认为是20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 形似：2021-05-21 00:00:00 .取值在当前日前30天内，不传默认是当前日前30天日期
	DateBegin *string `json:"DateBegin,omitempty" name:"DateBegin"`

	// 形似：2021-05-21 23:59:59 .取值在当前日前30天内，不传默认是当前日前30天日期
	DateEnd *string `json:"DateEnd,omitempty" name:"DateEnd"`
}

type DescribeWarningsRequest struct {
	*tchttp.BaseRequest
	
	// 1:创建时间倒序 2：创建时间升序 3：level倒序 4：leve升序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 可选设备id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 如果不传则查询所有，取值参见配置
	WarnLevelArray []*int64 `json:"WarnLevelArray,omitempty" name:"WarnLevelArray"`

	// 如果不传则查询所有，取值参见配置
	WarnModeArray []*int64 `json:"WarnModeArray,omitempty" name:"WarnModeArray"`

	// 不传认为是0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 不传认为是20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 形似：2021-05-21 00:00:00 .取值在当前日前30天内，不传默认是当前日前30天日期
	DateBegin *string `json:"DateBegin,omitempty" name:"DateBegin"`

	// 形似：2021-05-21 23:59:59 .取值在当前日前30天内，不传默认是当前日前30天日期
	DateEnd *string `json:"DateEnd,omitempty" name:"DateEnd"`
}

func (r *DescribeWarningsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWarningsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderType")
	delete(f, "DeviceId")
	delete(f, "WarnLevelArray")
	delete(f, "WarnModeArray")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DateBegin")
	delete(f, "DateEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWarningsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWarningsResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 告警列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*WarningsData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWarningsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWarningsResponseParams `json:"Response"`
}

func (r *DescribeWarningsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWarningsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeXP2PDataRequestParams struct {
	// P2P应用ID
	P2PAppId *string `json:"P2PAppId,omitempty" name:"P2PAppId"`

	// 查询开始时间，时间戳秒
	From *int64 `json:"From,omitempty" name:"From"`

	// 查询结束时间，时间戳秒
	To *int64 `json:"To,omitempty" name:"To"`

	// P2P通路ID
	P2PChannelId *string `json:"P2PChannelId,omitempty" name:"P2PChannelId"`
}

type DescribeXP2PDataRequest struct {
	*tchttp.BaseRequest
	
	// P2P应用ID
	P2PAppId *string `json:"P2PAppId,omitempty" name:"P2PAppId"`

	// 查询开始时间，时间戳秒
	From *int64 `json:"From,omitempty" name:"From"`

	// 查询结束时间，时间戳秒
	To *int64 `json:"To,omitempty" name:"To"`

	// P2P通路ID
	P2PChannelId *string `json:"P2PChannelId,omitempty" name:"P2PChannelId"`
}

func (r *DescribeXP2PDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeXP2PDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "P2PAppId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "P2PChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeXP2PDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeXP2PDataResponseParams struct {
	// [log_time,cdn_bytes , p2p_bytes, online_people, stuck_times, stuck_people,request,request_success,request_fail,play_fail]
	// [时间戳,cdn流量(字节) , p2p流量(字节), 在线人数, 卡播次数, 卡播人数,起播请求次数,起播成功次数,起播失败次数,播放失败次数, pcdn cdn流量（字节), pcdn路由流量(字节), 上传流量(字节)]
	// [1481016480, 46118502414, 75144943171, 61691, 3853, 0,0,0,0,0, 0, 0, 0]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeXP2PDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeXP2PDataResponseParams `json:"Response"`
}

func (r *DescribeXP2PDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeXP2PDataResponse) FromJsonString(s string) error {
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

type DeviceMonitorValue struct {
	// 统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitempty" name:"Value"`

	// 统计时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *int64 `json:"Time,omitempty" name:"Time"`
}

type Events struct {
	// 开始时间，秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventTime *int64 `json:"EventTime,omitempty" name:"EventTime"`

	// 事件类型 1:注册 2:心跳 4:录制异常 5:播放异常 6:流中断
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventType *int64 `json:"EventType,omitempty" name:"EventType"`

	// 事件描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventDesc *string `json:"EventDesc,omitempty" name:"EventDesc"`

	// 设备类型
	DeviceType *int64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 设备地址
	DeviceAddress *string `json:"DeviceAddress,omitempty" name:"DeviceAddress"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 通道Id
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 事件日志
	EventLog *string `json:"EventLog,omitempty" name:"EventLog"`

	// 设备备注名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

// Predefined struct for user
type GetRecordDatesByDevRequestParams struct {
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 限制量，默认200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 通道唯一标识，对于NVR设备，多通道IPC设备，设备编码与通道编码不一致的IPC设备，此字段为必填
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 1: 云端录制 2: 本地录制
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type GetRecordDatesByDevRequest struct {
	*tchttp.BaseRequest
	
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 限制量，默认200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 通道唯一标识，对于NVR设备，多通道IPC设备，设备编码与通道编码不一致的IPC设备，此字段为必填
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
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ChannelId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRecordDatesByDevRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRecordDatesByDevResponseParams struct {
	// 含有录像文件的日期列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dates []*string `json:"Dates,omitempty" name:"Dates"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRecordDatesByDevResponse struct {
	*tchttp.BaseResponse
	Response *GetRecordDatesByDevResponseParams `json:"Response"`
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

// Predefined struct for user
type GetRecordPlanByDevRequestParams struct {
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
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

// Predefined struct for user
type GetRecordPlanByDevResponseParams struct {
	// 录制计划详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Plan *RecordPlanItem `json:"Plan,omitempty" name:"Plan"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRecordPlanByDevResponse struct {
	*tchttp.BaseResponse
	Response *GetRecordPlanByDevResponseParams `json:"Response"`
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

// Predefined struct for user
type GetRecordPlanByIdRequestParams struct {
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`
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

// Predefined struct for user
type GetRecordPlanByIdResponseParams struct {
	// 录制计划详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Plan *RecordPlanItem `json:"Plan,omitempty" name:"Plan"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRecordPlanByIdResponse struct {
	*tchttp.BaseResponse
	Response *GetRecordPlanByIdResponseParams `json:"Response"`
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

// Predefined struct for user
type GetRecordPlansRequestParams struct {

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

// Predefined struct for user
type GetRecordPlansResponseParams struct {
	// 录制计划详情·列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Plans []*RecordPlanItem `json:"Plans,omitempty" name:"Plans"`

	// 录制计划总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRecordPlansResponse struct {
	*tchttp.BaseResponse
	Response *GetRecordPlansResponseParams `json:"Response"`
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

// Predefined struct for user
type GetTimeTemplateByIdRequestParams struct {
	// 时间模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type GetTimeTemplateByIdResponseParams struct {
	// 时间模板详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Template *TimeTemplateItem `json:"Template,omitempty" name:"Template"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTimeTemplateByIdResponse struct {
	*tchttp.BaseResponse
	Response *GetTimeTemplateByIdResponseParams `json:"Response"`
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

// Predefined struct for user
type GetTimeTemplatesRequestParams struct {

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

// Predefined struct for user
type GetTimeTemplatesResponseParams struct {
	// 时间模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Templates []*TimeTemplateItem `json:"Templates,omitempty" name:"Templates"`

	// 时间模板总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTimeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *GetTimeTemplatesResponseParams `json:"Response"`
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

// Predefined struct for user
type GetVideoListByConRequestParams struct {
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 通道唯一标识，对于NVR设备，多通道IPC设备，设备编码与通道编码不一致的IPC设备，此字段为必填
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 0：查询指定日期的录像；1：查询最近一天的录像；默认0
	LatestDay *int64 `json:"LatestDay,omitempty" name:"LatestDay"`

	// 指定某天。取值【YYYY-MM-DD】
	// 为空时默认查询最近一天的记录
	Date *string `json:"Date,omitempty" name:"Date"`

	// 1: 云端录制 2: 本地录制
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type GetVideoListByConRequest struct {
	*tchttp.BaseRequest
	
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 通道唯一标识，对于NVR设备，多通道IPC设备，设备编码与通道编码不一致的IPC设备，此字段为必填
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 0：查询指定日期的录像；1：查询最近一天的录像；默认0
	LatestDay *int64 `json:"LatestDay,omitempty" name:"LatestDay"`

	// 指定某天。取值【YYYY-MM-DD】
	// 为空时默认查询最近一天的记录
	Date *string `json:"Date,omitempty" name:"Date"`

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
	delete(f, "ChannelId")
	delete(f, "LatestDay")
	delete(f, "Date")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetVideoListByConRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetVideoListByConResponseParams struct {
	// 录像详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoList []*RecordTaskItem `json:"VideoList,omitempty" name:"VideoList"`

	// 录像总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetVideoListByConResponse struct {
	*tchttp.BaseResponse
	Response *GetVideoListByConResponseParams `json:"Response"`
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

	// 设备创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 设备通道总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelNum *int64 `json:"ChannelNum,omitempty" name:"ChannelNum"`

	// 设备视频通道总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoChannelNum *int64 `json:"VideoChannelNum,omitempty" name:"VideoChannelNum"`
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

// Predefined struct for user
type ModifyBindPlanLiveChannelRequestParams struct {
	// 直播录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 1: 绑定 2: 解绑
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 直播频道ID列表
	LiveChannelIds []*string `json:"LiveChannelIds,omitempty" name:"LiveChannelIds"`
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

// Predefined struct for user
type ModifyBindPlanLiveChannelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBindPlanLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBindPlanLiveChannelResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyBindRecordingPlanRequestParams struct {
	// 操作类型： 1-绑定设备 ；2-解绑设备
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 录制通道列表
	Channels []*ChannelItem `json:"Channels,omitempty" name:"Channels"`
}

type ModifyBindRecordingPlanRequest struct {
	*tchttp.BaseRequest
	
	// 操作类型： 1-绑定设备 ；2-解绑设备
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 录制通道列表
	Channels []*ChannelItem `json:"Channels,omitempty" name:"Channels"`
}

func (r *ModifyBindRecordingPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBindRecordingPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "PlanId")
	delete(f, "Channels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBindRecordingPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBindRecordingPlanResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBindRecordingPlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBindRecordingPlanResponseParams `json:"Response"`
}

func (r *ModifyBindRecordingPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBindRecordingPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBindSceneChannelsRequestParams struct {
	// 场景ID
	SceneId *int64 `json:"SceneId,omitempty" name:"SceneId"`

	// 1: 绑定 2: 解绑
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 通道列表
	Channels []*ChannelItem `json:"Channels,omitempty" name:"Channels"`
}

type ModifyBindSceneChannelsRequest struct {
	*tchttp.BaseRequest
	
	// 场景ID
	SceneId *int64 `json:"SceneId,omitempty" name:"SceneId"`

	// 1: 绑定 2: 解绑
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 通道列表
	Channels []*ChannelItem `json:"Channels,omitempty" name:"Channels"`
}

func (r *ModifyBindSceneChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBindSceneChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SceneId")
	delete(f, "Type")
	delete(f, "Channels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBindSceneChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBindSceneChannelsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBindSceneChannelsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBindSceneChannelsResponseParams `json:"Response"`
}

func (r *ModifyBindSceneChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBindSceneChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBindSceneDeviceRequestParams struct {
	// 场景ID
	SceneId *int64 `json:"SceneId,omitempty" name:"SceneId"`

	// 1: 绑定 2: 解绑
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 设备列表
	Devices []*DeviceItem `json:"Devices,omitempty" name:"Devices"`
}

type ModifyBindSceneDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 场景ID
	SceneId *int64 `json:"SceneId,omitempty" name:"SceneId"`

	// 1: 绑定 2: 解绑
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 设备列表
	Devices []*DeviceItem `json:"Devices,omitempty" name:"Devices"`
}

func (r *ModifyBindSceneDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBindSceneDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SceneId")
	delete(f, "Type")
	delete(f, "Devices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBindSceneDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBindSceneDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBindSceneDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBindSceneDeviceResponseParams `json:"Response"`
}

func (r *ModifyBindSceneDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBindSceneDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceDataRequestParams struct {
	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 设备名称
	NickName *string `json:"NickName,omitempty" name:"NickName"`
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

// Predefined struct for user
type ModifyDeviceDataResponseParams struct {
	// 操作结果,“OK”表示成功，其他表示失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceDataResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLiveChannelRequestParams struct {
	// 直播频道ID
	LiveChannelId *string `json:"LiveChannelId,omitempty" name:"LiveChannelId"`

	// 直播频道名
	LiveChannelName *string `json:"LiveChannelName,omitempty" name:"LiveChannelName"`
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

// Predefined struct for user
type ModifyLiveChannelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLiveChannelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLiveChannelResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLiveRecordPlanRequestParams struct {
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 录制计划名
	PlanName *string `json:"PlanName,omitempty" name:"PlanName"`

	// 时间模板ID，固定直播时为必填
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
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

// Predefined struct for user
type ModifyLiveRecordPlanResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLiveRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLiveRecordPlanResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyLiveVideoRequestParams struct {
	// 视频ID 列表, 大小限制(100)
	IntIDs []*int64 `json:"IntIDs,omitempty" name:"IntIDs"`

	// 过期时间 秒 (-1: 为永不过期)
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
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

// Predefined struct for user
type ModifyLiveVideoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLiveVideoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLiveVideoResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyMessageForwardRequestParams struct {
	// 配置ID
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`

	// json数组， 转发类型 1: 告警 2:GPS
	MessageType *string `json:"MessageType,omitempty" name:"MessageType"`
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

// Predefined struct for user
type ModifyMessageForwardResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMessageForwardResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMessageForwardResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyPresetRequestParams struct {
	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 预置位编码 范围1-8
	PresetId *int64 `json:"PresetId,omitempty" name:"PresetId"`

	// 预制位名称
	PresetName *string `json:"PresetName,omitempty" name:"PresetName"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

type ModifyPresetRequest struct {
	*tchttp.BaseRequest
	
	// 通道ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 预置位编码 范围1-8
	PresetId *int64 `json:"PresetId,omitempty" name:"PresetId"`

	// 预制位名称
	PresetName *string `json:"PresetName,omitempty" name:"PresetName"`

	// 设备Id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

func (r *ModifyPresetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPresetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "PresetId")
	delete(f, "PresetName")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPresetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPresetResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPresetResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPresetResponseParams `json:"Response"`
}

func (r *ModifyPresetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordingPlanRequestParams struct {
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 计划名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间模板ID
	TimeTemplateId *string `json:"TimeTemplateId,omitempty" name:"TimeTemplateId"`
}

type ModifyRecordingPlanRequest struct {
	*tchttp.BaseRequest
	
	// 录制计划ID
	PlanId *string `json:"PlanId,omitempty" name:"PlanId"`

	// 计划名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间模板ID
	TimeTemplateId *string `json:"TimeTemplateId,omitempty" name:"TimeTemplateId"`
}

func (r *ModifyRecordingPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordingPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "Name")
	delete(f, "TimeTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordingPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordingPlanResponseParams struct {
	// 操作结果
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRecordingPlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRecordingPlanResponseParams `json:"Response"`
}

func (r *ModifyRecordingPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordingPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySceneRequestParams struct {
	// 场景ID
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`

	// 场景名称
	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`

	// 触发条件
	SceneTrigger *string `json:"SceneTrigger,omitempty" name:"SceneTrigger"`

	// 录制时长(秒)
	RecordDuration *int64 `json:"RecordDuration,omitempty" name:"RecordDuration"`
}

type ModifySceneRequest struct {
	*tchttp.BaseRequest
	
	// 场景ID
	IntId *int64 `json:"IntId,omitempty" name:"IntId"`

	// 场景名称
	SceneName *string `json:"SceneName,omitempty" name:"SceneName"`

	// 触发条件
	SceneTrigger *string `json:"SceneTrigger,omitempty" name:"SceneTrigger"`

	// 录制时长(秒)
	RecordDuration *int64 `json:"RecordDuration,omitempty" name:"RecordDuration"`
}

func (r *ModifySceneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySceneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IntId")
	delete(f, "SceneName")
	delete(f, "SceneTrigger")
	delete(f, "RecordDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySceneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySceneResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySceneResponse struct {
	*tchttp.BaseResponse
	Response *ModifySceneResponseParams `json:"Response"`
}

func (r *ModifySceneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySceneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscriptionStatusRequestParams struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 订阅状态 1：关闭订阅 2：开启订阅
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 订阅类型 Alarm:告警订阅 Catalog:目录订阅 MobilePosition:移动位置订阅
	SubscriptionItem *string `json:"SubscriptionItem,omitempty" name:"SubscriptionItem"`
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

// Predefined struct for user
type ModifySubscriptionStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubscriptionStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubscriptionStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyVideoInfoRequestParams struct {
	// 视频ID列表长度限制100内
	InitIDs []*int64 `json:"InitIDs,omitempty" name:"InitIDs"`

	// 过期时间 时间戳 -1: 永不过期 0: 无效值
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
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

// Predefined struct for user
type ModifyVideoInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVideoInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVideoInfoResponseParams `json:"Response"`
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

type PresetItem struct {
	// 预置位ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetId *int64 `json:"PresetId,omitempty" name:"PresetId"`

	// 预置位名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetName *string `json:"PresetName,omitempty" name:"PresetName"`

	// 预置位状态 0:未设置预置位 1:已设置预置位 2:已设置预置位&看守位
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 预置位启用时的自动归位时间
	ResetTime *int64 `json:"ResetTime,omitempty" name:"ResetTime"`
}

type RecordPlanDetail struct {
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

	// 绑定的通道列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Channels []*ChannelItem `json:"Channels,omitempty" name:"Channels"`

	// 存储周期（天）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordStorageTime *int64 `json:"RecordStorageTime,omitempty" name:"RecordStorageTime"`
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

type RecordStatistic struct {
	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *RecordStatisticValue `json:"Value,omitempty" name:"Value"`
}

type RecordStatisticValue struct {
	// 期望执行时间 秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpectTimeLen *int64 `json:"ExpectTimeLen,omitempty" name:"ExpectTimeLen"`

	// 实际执行时间 秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordTimeLen *int64 `json:"RecordTimeLen,omitempty" name:"RecordTimeLen"`

	// 存储大小 G
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *float64 `json:"FileSize,omitempty" name:"FileSize"`
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

// Predefined struct for user
type ResetWarningRequestParams struct {
	// 告警ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Es中告警ID
	Index *string `json:"Index,omitempty" name:"Index"`
}

type ResetWarningRequest struct {
	*tchttp.BaseRequest
	
	// 告警ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// Es中告警ID
	Index *string `json:"Index,omitempty" name:"Index"`
}

func (r *ResetWarningRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetWarningRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Index")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetWarningRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetWarningResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetWarningResponse struct {
	*tchttp.BaseResponse
	Response *ResetWarningResponseParams `json:"Response"`
}

func (r *ResetWarningResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetWarningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// 是否全时录制，即7*24小时录制 0-否 1-是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllWeek *int64 `json:"IsAllWeek,omitempty" name:"IsAllWeek"`

	// 是否为自定义模板 0-否 1-是
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

// Predefined struct for user
type UpdateDeviceGroupRequestParams struct {
	// 分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组描述
	GroupDescribe *string `json:"GroupDescribe,omitempty" name:"GroupDescribe"`

	// 新父分组ID，用于修改分组路径
	NewParentId *string `json:"NewParentId,omitempty" name:"NewParentId"`
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

// Predefined struct for user
type UpdateDeviceGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDeviceGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type UpdateDevicePassWordRequestParams struct {
	// 设备密码
	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`

	// 设备唯一标识
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
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

// Predefined struct for user
type UpdateDevicePassWordResponseParams struct {
	// 操作结果，“OK”表示成功，其他表示失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateDevicePassWordResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDevicePassWordResponseParams `json:"Response"`
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

// Predefined struct for user
type UpdateRecordPlanRequestParams struct {
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

// Predefined struct for user
type UpdateRecordPlanResponseParams struct {
	// 操作结果
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRecordPlanResponseParams `json:"Response"`
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

// Predefined struct for user
type UpdateTimeTemplateRequestParams struct {
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

// Predefined struct for user
type UpdateTimeTemplateResponseParams struct {
	// 操作结果，“OK”表示成功，其他表示失败。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateTimeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTimeTemplateResponseParams `json:"Response"`
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

type WarningsData struct {
	// 唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 设备ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 告警通道
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarnChannel *string `json:"WarnChannel,omitempty" name:"WarnChannel"`

	// 告警级别 1: "一级警情", 2: "二级警情", 3: "三级警情", 4: "四级警情",
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarnLevel *int64 `json:"WarnLevel,omitempty" name:"WarnLevel"`

	// 告警级别名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarnLevelName *string `json:"WarnLevelName,omitempty" name:"WarnLevelName"`

	// 告警方式 2 设备报警 5 视频报警 6 设备故障报警
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarnMode *int64 `json:"WarnMode,omitempty" name:"WarnMode"`

	// 告警方式名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarnModeName *string `json:"WarnModeName,omitempty" name:"WarnModeName"`

	// 告警类型  2: {
	// 			Name: "设备报警",
	// 			WarnType: map[int]string{
	// 				1: "视频丢失报警",
	// 				2: "设备防拆报警",
	// 				3: "存储设备磁盘满报警",
	// 				4: "设备高温报警",
	// 				5: "设备低温报警",
	// 			},
	// 		},
	// 		5: {
	// 			Name: "视频报警",
	// 			WarnType: map[int]string{
	// 				1:  "人工视频报警",
	// 				2:  "运动目标检测报警",
	// 				3:  "遗留物检测报警",
	// 				4:  "物体移除检测报警",
	// 				5:  "绊线检测报警",
	// 				6:  "入侵检测报警",
	// 				7:  "逆行检测报警",
	// 				8:  "徘徊检测报警",
	// 				9:  "流量统计报警",
	// 				10: "密度检测报警",
	// 				11: "视频异常检测报警",
	// 				12: "快速移动报警",
	// 			},
	// 		},
	// 		6: {
	// 			Name: "设备故障报警",
	// 			WarnType: map[int]string{
	// 				1: "存储设备磁盘故障报警",
	// 				2: "存储设备风扇故障报警",
	// 			},
	// 		}
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarnType *int64 `json:"WarnType,omitempty" name:"WarnType"`

	// 是否删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Del *int64 `json:"Del,omitempty" name:"Del"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}