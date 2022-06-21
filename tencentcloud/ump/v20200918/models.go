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

package v20200918

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BunkZone struct {
	// 点位ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 点位名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 铺位编码
	BunkCodes *string `json:"BunkCodes,omitempty" name:"BunkCodes"`
}

type CameraConfig struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 楼层ID
	FloorId *int64 `json:"FloorId,omitempty" name:"FloorId"`

	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// 摄像头IP
	CameraIp *string `json:"CameraIp,omitempty" name:"CameraIp"`

	// 摄像头Mac
	CameraMac *string `json:"CameraMac,omitempty" name:"CameraMac"`

	// 摄像头类型:
	// 1: 码流机
	// 2: AI相机
	CameraType *int64 `json:"CameraType,omitempty" name:"CameraType"`

	// 摄像头功能:
	// 1: 人脸
	// 2: 人体
	CameraFeature *int64 `json:"CameraFeature,omitempty" name:"CameraFeature"`

	// 摄像头是否启用:
	// 0: 下线
	// 1: 启用
	CameraState *int64 `json:"CameraState,omitempty" name:"CameraState"`

	// 点位ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 点位类型:
	// 1: 场门
	// 3: 层门
	// 5: 特殊区域
	// 7: 门店
	// 8: 补位
	// 10: 开放式门店
	// 11: 品类区
	// 12: 公共区
	ZoneType *int64 `json:"ZoneType,omitempty" name:"ZoneType"`

	// 配置
	Config *Config `json:"Config,omitempty" name:"Config"`

	// 宽
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type CameraState struct {
	// 相机ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// 相机状态:
	// 10: 初始化
	// 11: 未知状态
	// 12: 网络异常
	// 13: 未授权
	// 14: 相机App异常
	// 15: 相机取流异常
	// 16: 状态正常
	State *uint64 `json:"State,omitempty" name:"State"`
}

type CameraZones struct {
	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// 摄像头名称
	CameraName *string `json:"CameraName,omitempty" name:"CameraName"`

	// 摄像头功能:
	// 1: 人脸
	// 2: 人体
	CameraFeature *int64 `json:"CameraFeature,omitempty" name:"CameraFeature"`

	// 摄像头IP
	CameraIp *string `json:"CameraIp,omitempty" name:"CameraIp"`

	// 摄像头状态:
	// 0: 异常 (不再使用)
	// 1: 正常 (不再使用)
	// 10: 初始化
	// 11: 未知状态 (因服务内部错误产生)
	// 12: 网络异常
	// 13: 未授权
	// 14: 相机App异常
	// 15: 相机取流异常
	// 16: 正常
	CameraState *int64 `json:"CameraState,omitempty" name:"CameraState"`

	// 点位列表
	Zones []*BunkZone `json:"Zones,omitempty" name:"Zones"`

	// 像素:
	// 130W(1280*960)
	// 200W(1920*1080)
	// 400W(2560*1440)
	Pixel *string `json:"Pixel,omitempty" name:"Pixel"`

	// 相机Rtsp地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RTSP *string `json:"RTSP,omitempty" name:"RTSP"`
}

type Config struct {
	// 摄像头厂商:
	// H: 海康
	// D: 大华
	// Y: 英飞拓
	// L: 联纵
	CameraProducer *string `json:"CameraProducer,omitempty" name:"CameraProducer"`

	// rtsp 地址
	RTSP *string `json:"RTSP,omitempty" name:"RTSP"`

	// 摄像头帧率
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 解码帧率
	DecodeFps *int64 `json:"DecodeFps,omitempty" name:"DecodeFps"`

	// 是否做客流计算:
	// 0: 否
	// 1: 是
	PassengerFlow *int64 `json:"PassengerFlow,omitempty" name:"PassengerFlow"`

	// 是否打开人脸曝光:
	// 0: 关闭
	// 1: 开启
	FaceExpose *int64 `json:"FaceExpose,omitempty" name:"FaceExpose"`

	// 门线标注
	MallArea []*Point `json:"MallArea,omitempty" name:"MallArea"`

	// 店门标注
	ShopArea []*Point `json:"ShopArea,omitempty" name:"ShopArea"`

	// 检测区标注
	TrackAreas []*Polygon `json:"TrackAreas,omitempty" name:"TrackAreas"`

	// 点位列表（品类区）
	Zones []*ZoneArea `json:"Zones,omitempty" name:"Zones"`
}

type CreateCameraAlertAlert struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 相机ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// 时间戳,ms,默认为告警请求到达时间
	CaptureTime *int64 `json:"CaptureTime,omitempty" name:"CaptureTime"`

	// 图片base64编码
	Image *string `json:"Image,omitempty" name:"Image"`

	// 移动告警
	MoveAlert *CreateCameraAlertsMoveAlert `json:"MoveAlert,omitempty" name:"MoveAlert"`

	// 遮挡告警
	CoverAlert *CreateCameraAlertsCoverAlert `json:"CoverAlert,omitempty" name:"CoverAlert"`
}

type CreateCameraAlertsCoverAlert struct {
	// 是否遮挡
	Cover *bool `json:"Cover,omitempty" name:"Cover"`

	// 是否移动置信度
	CoverConfidence *float64 `json:"CoverConfidence,omitempty" name:"CoverConfidence"`
}

type CreateCameraAlertsMoveAlert struct {
	// 是否移动
	Move *bool `json:"Move,omitempty" name:"Move"`

	// 是否移动置信度
	MoveConfidence *float64 `json:"MoveConfidence,omitempty" name:"MoveConfidence"`
}

// Predefined struct for user
type CreateCameraAlertsRequestParams struct {
	// 告警信息列表
	Alerts []*CreateCameraAlertAlert `json:"Alerts,omitempty" name:"Alerts"`
}

type CreateCameraAlertsRequest struct {
	*tchttp.BaseRequest
	
	// 告警信息列表
	Alerts []*CreateCameraAlertAlert `json:"Alerts,omitempty" name:"Alerts"`
}

func (r *CreateCameraAlertsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCameraAlertsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Alerts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCameraAlertsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCameraAlertsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCameraAlertsResponse struct {
	*tchttp.BaseResponse
	Response *CreateCameraAlertsResponseParams `json:"Response"`
}

func (r *CreateCameraAlertsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCameraAlertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCameraStateRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 场内所有相机的状态值
	CameraStates []*CameraState `json:"CameraStates,omitempty" name:"CameraStates"`
}

type CreateCameraStateRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 场内所有相机的状态值
	CameraStates []*CameraState `json:"CameraStates,omitempty" name:"CameraStates"`
}

func (r *CreateCameraStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCameraStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "MallId")
	delete(f, "CameraStates")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCameraStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCameraStateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCameraStateResponse struct {
	*tchttp.BaseResponse
	Response *CreateCameraStateResponseParams `json:"Response"`
}

func (r *CreateCameraStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCameraStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCaptureRequestParams struct {
	// 原始抓拍报文
	Data *string `json:"Data,omitempty" name:"Data"`
}

type CreateCaptureRequest struct {
	*tchttp.BaseRequest
	
	// 原始抓拍报文
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *CreateCaptureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCaptureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCaptureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCaptureResponseParams struct {
	// 原始应答报文
	// 注意：此字段可能返回 null，表示取不到有效值。
	RspData *string `json:"RspData,omitempty" name:"RspData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCaptureResponse struct {
	*tchttp.BaseResponse
	Response *CreateCaptureResponseParams `json:"Response"`
}

func (r *CreateCaptureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCaptureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiBizAlertRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 点位ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// 时间戳，毫秒
	CaptureTime *uint64 `json:"CaptureTime,omitempty" name:"CaptureTime"`

	// 状态: 
	// 1: 侵占
	// 2: 消失
	// 3: 即侵占又消失
	State *int64 `json:"State,omitempty" name:"State"`

	// 图片base64字符串
	Image *string `json:"Image,omitempty" name:"Image"`

	// 告警列表
	Warnings []*MultiBizWarning `json:"Warnings,omitempty" name:"Warnings"`
}

type CreateMultiBizAlertRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 点位ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// 时间戳，毫秒
	CaptureTime *uint64 `json:"CaptureTime,omitempty" name:"CaptureTime"`

	// 状态: 
	// 1: 侵占
	// 2: 消失
	// 3: 即侵占又消失
	State *int64 `json:"State,omitempty" name:"State"`

	// 图片base64字符串
	Image *string `json:"Image,omitempty" name:"Image"`

	// 告警列表
	Warnings []*MultiBizWarning `json:"Warnings,omitempty" name:"Warnings"`
}

func (r *CreateMultiBizAlertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiBizAlertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "MallId")
	delete(f, "ZoneId")
	delete(f, "CameraId")
	delete(f, "CaptureTime")
	delete(f, "State")
	delete(f, "Image")
	delete(f, "Warnings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMultiBizAlertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiBizAlertResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMultiBizAlertResponse struct {
	*tchttp.BaseResponse
	Response *CreateMultiBizAlertResponseParams `json:"Response"`
}

func (r *CreateMultiBizAlertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiBizAlertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProgramStateRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 进程监控信息列表
	ProgramStateItems []*ProgramStateItem `json:"ProgramStateItems,omitempty" name:"ProgramStateItems"`

	// 商场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`
}

type CreateProgramStateRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 进程监控信息列表
	ProgramStateItems []*ProgramStateItem `json:"ProgramStateItems,omitempty" name:"ProgramStateItems"`

	// 商场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`
}

func (r *CreateProgramStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProgramStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "ProgramStateItems")
	delete(f, "MallId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProgramStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProgramStateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProgramStateResponse struct {
	*tchttp.BaseResponse
	Response *CreateProgramStateResponseParams `json:"Response"`
}

func (r *CreateProgramStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProgramStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerStateRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 服务器监控信息列表
	ServerStateItems []*ServerStateItem `json:"ServerStateItems,omitempty" name:"ServerStateItems"`

	// 商场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 服务器监控信息上报时间戳，单位毫秒
	ReportTime *uint64 `json:"ReportTime,omitempty" name:"ReportTime"`
}

type CreateServerStateRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 服务器监控信息列表
	ServerStateItems []*ServerStateItem `json:"ServerStateItems,omitempty" name:"ServerStateItems"`

	// 商场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 服务器监控信息上报时间戳，单位毫秒
	ReportTime *uint64 `json:"ReportTime,omitempty" name:"ReportTime"`
}

func (r *CreateServerStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "ServerStateItems")
	delete(f, "MallId")
	delete(f, "ReportTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServerStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerStateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateServerStateResponse struct {
	*tchttp.BaseResponse
	Response *CreateServerStateResponseParams `json:"Response"`
}

func (r *CreateServerStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMultiBizAlertRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 点位ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// 消警动作:
	// 1: 误报
	// 2: 正报合规
	// 3: 正报不合规，整改完成
	ActionType *int64 `json:"ActionType,omitempty" name:"ActionType"`

	// 图片base64字符串
	Image *string `json:"Image,omitempty" name:"Image"`
}

type DeleteMultiBizAlertRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 点位ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// 消警动作:
	// 1: 误报
	// 2: 正报合规
	// 3: 正报不合规，整改完成
	ActionType *int64 `json:"ActionType,omitempty" name:"ActionType"`

	// 图片base64字符串
	Image *string `json:"Image,omitempty" name:"Image"`
}

func (r *DeleteMultiBizAlertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultiBizAlertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "MallId")
	delete(f, "ZoneId")
	delete(f, "CameraId")
	delete(f, "ActionType")
	delete(f, "Image")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMultiBizAlertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMultiBizAlertResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMultiBizAlertResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMultiBizAlertResponseParams `json:"Response"`
}

func (r *DeleteMultiBizAlertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultiBizAlertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DeleteTaskRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "MallId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTaskResponseParams `json:"Response"`
}

func (r *DeleteTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCamerasRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`
}

type DescribeCamerasRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`
}

func (r *DescribeCamerasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCamerasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "MallId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCamerasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCamerasResponseParams struct {
	// 摄像头列表
	Cameras []*CameraZones `json:"Cameras,omitempty" name:"Cameras"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCamerasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCamerasResponseParams `json:"Response"`
}

func (r *DescribeCamerasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCamerasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigRequestParams struct {
	// 会话ID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 摄像头签名
	CameraSign *string `json:"CameraSign,omitempty" name:"CameraSign"`

	// 摄像头app id
	CameraAppId *string `json:"CameraAppId,omitempty" name:"CameraAppId"`

	// 摄像头时间戳，毫秒
	CameraTimestamp *int64 `json:"CameraTimestamp,omitempty" name:"CameraTimestamp"`

	// MAC地址，字母大写
	ServerMac *string `json:"ServerMac,omitempty" name:"ServerMac"`

	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`
}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest
	
	// 会话ID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 摄像头签名
	CameraSign *string `json:"CameraSign,omitempty" name:"CameraSign"`

	// 摄像头app id
	CameraAppId *string `json:"CameraAppId,omitempty" name:"CameraAppId"`

	// 摄像头时间戳，毫秒
	CameraTimestamp *int64 `json:"CameraTimestamp,omitempty" name:"CameraTimestamp"`

	// MAC地址，字母大写
	ServerMac *string `json:"ServerMac,omitempty" name:"ServerMac"`

	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`
}

func (r *DescribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "CameraSign")
	delete(f, "CameraAppId")
	delete(f, "CameraTimestamp")
	delete(f, "ServerMac")
	delete(f, "GroupCode")
	delete(f, "MallId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigResponseParams struct {
	// 会话ID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 配置版本号
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// 摄像头列表
	Cameras []*CameraConfig `json:"Cameras,omitempty" name:"Cameras"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigResponseParams `json:"Response"`
}

func (r *DescribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`
}

type DescribeImageRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`
}

func (r *DescribeImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "MallId")
	delete(f, "CameraId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageResponseParams struct {
	// cos 临时 url，异步上传图片，client需要轮询
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageResponseParams `json:"Response"`
}

func (r *DescribeImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiBizBaseImageRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// 点位ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type DescribeMultiBizBaseImageRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// 点位ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeMultiBizBaseImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiBizBaseImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "MallId")
	delete(f, "CameraId")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMultiBizBaseImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMultiBizBaseImageResponseParams struct {
	// cos 临时 url
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMultiBizBaseImageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMultiBizBaseImageResponseParams `json:"Response"`
}

func (r *DescribeMultiBizBaseImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMultiBizBaseImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 任务类型:
	// 1: 底图拉取
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 任务类型:
	// 1: 底图拉取
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
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
	delete(f, "GroupCode")
	delete(f, "MallId")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 任务列表
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

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

// Predefined struct for user
type DescribeZonesRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "MallId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// 点位列表
	Zones []*ZoneConfig `json:"Zones,omitempty" name:"Zones"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskInfo struct {
	// 硬盘名字
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// 硬盘使用率
	Usage *float64 `json:"Usage,omitempty" name:"Usage"`
}

// Predefined struct for user
type ModifyMultiBizConfigRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 点位ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// 监控区域
	MonitoringAreas []*Polygon `json:"MonitoringAreas,omitempty" name:"MonitoringAreas"`
}

type ModifyMultiBizConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 点位ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// 监控区域
	MonitoringAreas []*Polygon `json:"MonitoringAreas,omitempty" name:"MonitoringAreas"`
}

func (r *ModifyMultiBizConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMultiBizConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "MallId")
	delete(f, "ZoneId")
	delete(f, "CameraId")
	delete(f, "MonitoringAreas")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMultiBizConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMultiBizConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMultiBizConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMultiBizConfigResponseParams `json:"Response"`
}

func (r *ModifyMultiBizConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMultiBizConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MultiBizWarning struct {
	// 编号
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 监控区域
	MonitoringArea []*Point `json:"MonitoringArea,omitempty" name:"MonitoringArea"`

	// 告警列表
	WarningInfos []*MultiBizWarningInfo `json:"WarningInfos,omitempty" name:"WarningInfos"`
}

type MultiBizWarningInfo struct {
	// 告警类型：
	// 0: 无变化
	// 1: 侵占
	// 2: 消失
	WarningType *int64 `json:"WarningType,omitempty" name:"WarningType"`

	// 告警侵占或消失面积
	WarningAreaSize *float64 `json:"WarningAreaSize,omitempty" name:"WarningAreaSize"`

	// 告警侵占或消失坐标
	WarningLocation *Point `json:"WarningLocation,omitempty" name:"WarningLocation"`

	// 告警侵占或消失轮廓
	WarningAreaContour []*Point `json:"WarningAreaContour,omitempty" name:"WarningAreaContour"`
}

type Point struct {
	// X坐标
	X *int64 `json:"X,omitempty" name:"X"`

	// Y坐标
	Y *int64 `json:"Y,omitempty" name:"Y"`
}

type Polygon struct {
	// 标注列表
	Points []*Point `json:"Points,omitempty" name:"Points"`
}

type ProgramStateItem struct {
	// 服务器IP
	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`

	// 进程名字
	ProgramName *string `json:"ProgramName,omitempty" name:"ProgramName"`

	// 在线个数
	OnlineCount *uint64 `json:"OnlineCount,omitempty" name:"OnlineCount"`

	// 离线个数
	OfflineCount *uint64 `json:"OfflineCount,omitempty" name:"OfflineCount"`

	// 上报状态:
	// 1: 正常上报
	// 2: 异常上报
	// 注：此处异常上报是指本次上报由于场内服务内部原因导致上报数据不可信等。此时离线个数重置为1，在线个数重置为0
	State *int64 `json:"State,omitempty" name:"State"`
}

// Predefined struct for user
type ReportServiceRegisterRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 服务上报当前的服务能力信息
	ServiceRegisterInfos []*ServiceRegisterInfo `json:"ServiceRegisterInfos,omitempty" name:"ServiceRegisterInfos"`

	// 服务内网Ip
	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`

	// 上报服务所在服务器的唯一ID
	ServerNodeId *string `json:"ServerNodeId,omitempty" name:"ServerNodeId"`

	// 上报时间戳, 单位毫秒
	ReportTime *int64 `json:"ReportTime,omitempty" name:"ReportTime"`
}

type ReportServiceRegisterRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 服务上报当前的服务能力信息
	ServiceRegisterInfos []*ServiceRegisterInfo `json:"ServiceRegisterInfos,omitempty" name:"ServiceRegisterInfos"`

	// 服务内网Ip
	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`

	// 上报服务所在服务器的唯一ID
	ServerNodeId *string `json:"ServerNodeId,omitempty" name:"ServerNodeId"`

	// 上报时间戳, 单位毫秒
	ReportTime *int64 `json:"ReportTime,omitempty" name:"ReportTime"`
}

func (r *ReportServiceRegisterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportServiceRegisterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "MallId")
	delete(f, "ServiceRegisterInfos")
	delete(f, "ServerIp")
	delete(f, "ServerNodeId")
	delete(f, "ReportTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportServiceRegisterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportServiceRegisterResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReportServiceRegisterResponse struct {
	*tchttp.BaseResponse
	Response *ReportServiceRegisterResponseParams `json:"Response"`
}

func (r *ReportServiceRegisterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportServiceRegisterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchImageRequestParams struct {
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 图片base64字符串
	Image *string `json:"Image,omitempty" name:"Image"`

	// 时间戳，毫秒
	ImageTime *uint64 `json:"ImageTime,omitempty" name:"ImageTime"`
}

type SearchImageRequest struct {
	*tchttp.BaseRequest
	
	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 图片base64字符串
	Image *string `json:"Image,omitempty" name:"Image"`

	// 时间戳，毫秒
	ImageTime *uint64 `json:"ImageTime,omitempty" name:"ImageTime"`
}

func (r *SearchImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupCode")
	delete(f, "MallId")
	delete(f, "Image")
	delete(f, "ImageTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchImageResponseParams struct {
	// face id
	FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

	// 搜索结果列表
	Results []*SearchResult `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchImageResponse struct {
	*tchttp.BaseResponse
	Response *SearchImageResponseParams `json:"Response"`
}

func (r *SearchImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchResult struct {
	// 图片base64数据
	Image *string `json:"Image,omitempty" name:"Image"`

	// 身份ID
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 相似度
	Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`
}

type ServerStateItem struct {
	// 服务器状态
	// 1: 在线
	// 2: 离线
	// 3: 重启
	ServerState *int64 `json:"ServerState,omitempty" name:"ServerState"`

	// 服务器IP
	ServerIp *string `json:"ServerIp,omitempty" name:"ServerIp"`

	// 硬盘监控信息列表
	DiskInfos []*DiskInfo `json:"DiskInfos,omitempty" name:"DiskInfos"`
}

type ServiceRegisterInfo struct {
	// 当前服务的回调地址
	CgiUrl *string `json:"CgiUrl,omitempty" name:"CgiUrl"`

	// 当前服务类型:
	// 1: 多经服务
	// 2: 相机误报警确认
	// 3: 底图更新
	ServiceType *uint64 `json:"ServiceType,omitempty" name:"ServiceType"`
}

type Task struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 集团编码
	GroupCode *string `json:"GroupCode,omitempty" name:"GroupCode"`

	// 广场ID
	MallId *uint64 `json:"MallId,omitempty" name:"MallId"`

	// 任务内容
	TaskContent *TaskContent `json:"TaskContent,omitempty" name:"TaskContent"`

	// 任务类型:
	// 1: 底图拉取
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
}

type TaskContent struct {
	// 摄像头ID
	CameraId *uint64 `json:"CameraId,omitempty" name:"CameraId"`

	// rtsp 地址
	RTSP *string `json:"RTSP,omitempty" name:"RTSP"`

	// 图片上传地址
	Url *string `json:"Url,omitempty" name:"Url"`
}

type ZoneArea struct {
	// 点位ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 店门标注
	ShopArea []*Point `json:"ShopArea,omitempty" name:"ShopArea"`
}

type ZoneConfig struct {
	// 点位ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 点位名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 点位类型:
	// 1: 场门
	// 3: 层门
	// 5: 特殊区域
	// 7: 门店
	// 8: 补位
	// 10: 开放式门店
	// 11: 品类区
	// 12: 公共区
	ZoneType *int64 `json:"ZoneType,omitempty" name:"ZoneType"`

	// 铺位编码
	BunkCodes *string `json:"BunkCodes,omitempty" name:"BunkCodes"`

	// 楼层名称
	FloorName *string `json:"FloorName,omitempty" name:"FloorName"`

	// 楼层ID
	FloorId *int64 `json:"FloorId,omitempty" name:"FloorId"`

	// 绑定数
	BindNum *int64 `json:"BindNum,omitempty" name:"BindNum"`

	// 调试数
	DebugNum *int64 `json:"DebugNum,omitempty" name:"DebugNum"`

	// 下发状态:
	// 1: 不可下发
	// 2: 可下发
	// 3: 已下发
	State *int64 `json:"State,omitempty" name:"State"`
}