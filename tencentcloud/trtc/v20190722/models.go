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

package v20190722

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeCallDetailRequest struct {
	*tchttp.BaseRequest

	// 通话ID
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// 查询开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需查询的用户数组，不填默认返回6个用户
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds" list`

	// 需查询的指标，不填则只返回用户列表，填all则返回所有指标。
	// appCpu：APP CPU使用率；
	// sysCpu：系统 CPU使用率；
	// aBit：上/下行音频码率；
	// aBlock：音频卡顿时长；
	// vBit：上/下行视频码率；
	// vCapFps：视频采集帧率；
	// vEncFps：视频发送帧率；
	// vDecFps：渲染帧率；
	// vBlock：视频卡顿时长；
	// aLoss：上/下行音频丢包；
	// vLoss：上/下行视频丢包；
	// vWidth：上/下行分辨率宽；
	// vHeight：上/下行分辨率高
	DataType []*string `json:"DataType,omitempty" name:"DataType" list`
}

func (r *DescribeCallDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCallDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCallDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的用户总条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 用户信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserList []*UserInformation `json:"UserList,omitempty" name:"UserList" list`

		// 质量数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*QualityData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCallDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCallDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryScaleRequest struct {
	*tchttp.BaseRequest

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeHistoryScaleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHistoryScaleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryScaleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 返回的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		ScaleList []*ScaleInfomation `json:"ScaleList,omitempty" name:"ScaleList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistoryScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHistoryScaleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeNetworkRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需查询的数据类型
	// sendLossRateRaw：上行丢包率；
	// recvLossRateRaw：下行丢包率
	DataType []*string `json:"DataType,omitempty" name:"DataType" list`
}

func (r *DescribeRealtimeNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeNetworkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeNetworkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询返回的数据
		Data []*RealtimeData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealtimeNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeNetworkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeQualityRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询的数据类型
	// enterTotalSuccPercent：进房成功率
	// fistFreamInSecRate：首帧秒开率
	// blockPercent：视频卡顿率
	// audioBlockPercent：音频卡顿率
	DataType []*string `json:"DataType,omitempty" name:"DataType" list`
}

func (r *DescribeRealtimeQualityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeQualityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeQualityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据类型
		Data []*RealtimeData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealtimeQualityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeQualityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeScaleRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询的数据类型
	// UserNum：通话人数；
	// RoomNum：房间数
	DataType []*string `json:"DataType,omitempty" name:"DataType" list`
}

func (r *DescribeRealtimeScaleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeScaleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeScaleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据数组
		Data []*RealtimeData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealtimeScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRealtimeScaleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoomInformationRequest struct {
	*tchttp.BaseRequest

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 数字房间号
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 分页index（不填默认只返回10个）
	PageNumber *string `json:"PageNumber,omitempty" name:"PageNumber"`

	// 分页大小（不填默认返回10个,最多不超过100条）
	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeRoomInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoomInformationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoomInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据总条数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 房间信息列表
		RoomList []*RoomState `json:"RoomList,omitempty" name:"RoomList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoomInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoomInformationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DismissRoomRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DismissRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DismissRoomRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DismissRoomResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DismissRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DismissRoomResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QualityData struct {

	// 数据内容
	Content []*TimeValue `json:"Content,omitempty" name:"Content" list`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 对端Id,为空时表示上行数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`

	// 数据类型
	DataType *string `json:"DataType,omitempty" name:"DataType"`
}

type RealtimeData struct {

	// 返回的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*TimeValue `json:"Content,omitempty" name:"Content" list`

	// 数据类型字段
	DataType *string `json:"DataType,omitempty" name:"DataType"`
}

type RemoveUserRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 要移出的用户列表，最多10个。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds" list`
}

func (r *RemoveUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RoomState struct {

	// 通话ID（唯一标识一次通话）
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// 房间号
	RoomString *string `json:"RoomString,omitempty" name:"RoomString"`

	// 房间创建时间
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 房间销毁时间
	DestroyTime *uint64 `json:"DestroyTime,omitempty" name:"DestroyTime"`

	// 房间是否已经结束
	IsFinished *bool `json:"IsFinished,omitempty" name:"IsFinished"`

	// 房间创建者Id
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type ScaleInfomation struct {

	// 每天开始的时间
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 房间人数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserNumber *uint64 `json:"UserNumber,omitempty" name:"UserNumber"`

	// 房间人次
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserCount *uint64 `json:"UserCount,omitempty" name:"UserCount"`

	// 房间数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomNumbers *uint64 `json:"RoomNumbers,omitempty" name:"RoomNumbers"`
}

type TimeValue struct {

	// 时间
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 当前时间取值
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type UserInformation struct {

	// 房间号
	RoomStr *string `json:"RoomStr,omitempty" name:"RoomStr"`

	// 用户Id
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户进房时间
	JoinTs *uint64 `json:"JoinTs,omitempty" name:"JoinTs"`

	// 用户退房时间
	LeaveTs *uint64 `json:"LeaveTs,omitempty" name:"LeaveTs"`

	// 终端类型
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Sdk版本号
	SdkVersion *string `json:"SdkVersion,omitempty" name:"SdkVersion"`

	// 客户端IP地址
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
}
