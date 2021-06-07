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
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AbnormalEvent struct {

	// 异常事件ID，具体值查看附录：异常体验ID映射表：https://cloud.tencent.com/document/product/647/44916
	AbnormalEventId *uint64 `json:"AbnormalEventId,omitempty" name:"AbnormalEventId"`

	// 远端用户ID,""：表示异常事件不是由远端用户产生
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`
}

type AbnormalExperience struct {

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 异常体验ID
	ExperienceId *uint64 `json:"ExperienceId,omitempty" name:"ExperienceId"`

	// 字符串房间号
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 异常事件数组
	AbnormalEventList []*AbnormalEvent `json:"AbnormalEventList,omitempty" name:"AbnormalEventList"`

	// 异常事件的上报时间
	EventTime *uint64 `json:"EventTime,omitempty" name:"EventTime"`
}

type CreatePictureRequest struct {
	*tchttp.BaseRequest

	// 应用id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 图片内容经base64编码后的string格式
	Content *string `json:"Content,omitempty" name:"Content"`

	// 图片后缀名
	Suffix *string `json:"Suffix,omitempty" name:"Suffix"`

	// 图片长度
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 图片宽度
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 显示位置x轴方向
	XPosition *uint64 `json:"XPosition,omitempty" name:"XPosition"`

	// 显示位置y轴方向
	YPosition *uint64 `json:"YPosition,omitempty" name:"YPosition"`
}

func (r *CreatePictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Content")
	delete(f, "Suffix")
	delete(f, "Height")
	delete(f, "Width")
	delete(f, "XPosition")
	delete(f, "YPosition")
	if len(f) > 0 {
		return errors.New("CreatePictureRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePictureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图片id
		PictureId *uint64 `json:"PictureId,omitempty" name:"PictureId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTroubleInfoRequest struct {
	*tchttp.BaseRequest

	// 应用的ID
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 老师用户ID
	TeacherUserId *string `json:"TeacherUserId,omitempty" name:"TeacherUserId"`

	// 学生用户ID
	StudentUserId *string `json:"StudentUserId,omitempty" name:"StudentUserId"`

	// 体验异常端（老师或学生）的用户 ID。
	TroubleUserId *string `json:"TroubleUserId,omitempty" name:"TroubleUserId"`

	// 异常类型。
	// 1. 仅视频异常
	// 2. 仅声音异常
	// 3. 音视频都异常
	// 5. 进房异常
	// 4. 切课
	// 6. 求助
	// 7. 问题反馈
	// 8. 投诉
	TroubleType *uint64 `json:"TroubleType,omitempty" name:"TroubleType"`

	// 异常发生的UNIX 时间戳，单位为秒。
	TroubleTime *uint64 `json:"TroubleTime,omitempty" name:"TroubleTime"`

	// 异常详情
	TroubleMsg *string `json:"TroubleMsg,omitempty" name:"TroubleMsg"`
}

func (r *CreateTroubleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTroubleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "TeacherUserId")
	delete(f, "StudentUserId")
	delete(f, "TroubleUserId")
	delete(f, "TroubleType")
	delete(f, "TroubleTime")
	delete(f, "TroubleMsg")
	if len(f) > 0 {
		return errors.New("CreateTroubleInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTroubleInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTroubleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTroubleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePictureRequest struct {
	*tchttp.BaseRequest

	// 图片id
	PictureId *uint64 `json:"PictureId,omitempty" name:"PictureId"`

	// 应用id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DeletePictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PictureId")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return errors.New("DeletePictureRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePictureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalEventRequest struct {
	*tchttp.BaseRequest

	// 用户SDKAppID，查询SDKAppID下任意20条异常体验事件（可能不同房间）
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 房间号，查询房间内任意20条以内异常体验事件
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DescribeAbnormalEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	if len(f) > 0 {
		return errors.New("DescribeAbnormalEventRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据总条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 异常体验列表
		AbnormalExperienceList []*AbnormalExperience `json:"AbnormalExperienceList,omitempty" name:"AbnormalExperienceList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCallDetailRequest struct {
	*tchttp.BaseRequest

	// 通话 ID（唯一标识一次通话）： sdkappid_roomgString（房间号_createTime（房间创建时间，unix时间戳，单位为s）例：1400353843_218695_1590065777。通过 DescribeRoomInformation（查询房间列表）接口获取（链接：https://cloud.tencent.com/document/product/647/44050）
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// 查询开始时间，14天内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户SDKAppID（1400188366）
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需查询的用户数组，不填默认返回6个用户,最多可填6个用户
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// 需查询的指标，不填则只返回用户列表，填all则返回所有指标。
	// appCpu：APP CPU使用率；
	// sysCpu：系统 CPU使用率；
	// aBit：上/下行音频码率；
	// aBlock：音频卡顿时长；
	// bigvBit：上/下行视频码率；
	// bigvCapFps：视频采集帧率；
	// bigvEncFps：视频发送帧率；
	// bigvDecFps：渲染帧率；
	// bigvBlock：视频卡顿时长；
	// aLoss：上/下行音频丢包；
	// bigvLoss：上/下行视频丢包；
	// bigvWidth：上/下行分辨率宽；
	// bigvHeight：上/下行分辨率高
	DataType []*string `json:"DataType,omitempty" name:"DataType"`

	// 设置分页index，从0开始（PageNumber和PageSize 其中一个不填均默认返回6条数据）
	PageNumber *string `json:"PageNumber,omitempty" name:"PageNumber"`

	// 设置分页大小（PageNumber和PageSize 其中一个不填均默认返回6条数据,DataType，UserIds不为null，PageSize最大不超过6，DataType，UserIds为null，PageSize最大不超过100）
	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeCallDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "UserIds")
	delete(f, "DataType")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return errors.New("DescribeCallDetailRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCallDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的用户总条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 用户信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserList []*UserInformation `json:"UserList,omitempty" name:"UserList"`

		// 质量数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*QualityData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCallDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDetailEventRequest struct {
	*tchttp.BaseRequest

	// 通话 ID（唯一标识一次通话）： sdkappid_roomgString（房间号_createTime（房间创建时间，unix时间戳，单位s）。通过 DescribeRoomInformation（查询房间列表）接口获取。（链接：https://cloud.tencent.com/document/product/647/44050）
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// 查询开始时间，14天内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户id
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 房间号
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DescribeDetailEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetailEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return errors.New("DescribeDetailEventRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDetailEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的事件列表
		Data []*EventList `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDetailEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetailEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryScaleRequest struct {
	*tchttp.BaseRequest

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询开始时间，5天内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeHistoryScaleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryScaleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return errors.New("DescribeHistoryScaleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryScaleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 返回的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		ScaleList []*ScaleInfomation `json:"ScaleList,omitempty" name:"ScaleList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistoryScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryScaleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePictureRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 图片ID，不填时返回该应用下所有图片
	PictureId *uint64 `json:"PictureId,omitempty" name:"PictureId"`

	// 每页数量，不填时默认为10
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页码，不填时默认为1
	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`
}

func (r *DescribePictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PictureId")
	delete(f, "PageSize")
	delete(f, "PageNo")
	if len(f) > 0 {
		return errors.New("DescribePictureRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePictureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的图片记录数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 图片信息列表
		PictureInfo []*PictureInfo `json:"PictureInfo,omitempty" name:"PictureInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeNetworkRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间，24小时内，本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需查询的数据类型
	// sendLossRateRaw：上行丢包率
	// recvLossRateRaw：下行丢包率
	DataType []*string `json:"DataType,omitempty" name:"DataType"`
}

func (r *DescribeRealtimeNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "DataType")
	if len(f) > 0 {
		return errors.New("DescribeRealtimeNetworkRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeNetworkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询返回的数据
		Data []*RealtimeData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealtimeNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeQualityRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间，24小时内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询的数据类型
	// enterTotalSuccPercent：进房成功率
	// fistFreamInSecRate：首帧秒开率
	// blockPercent：视频卡顿率
	// audioBlockPercent：音频卡顿率
	DataType []*string `json:"DataType,omitempty" name:"DataType"`
}

func (r *DescribeRealtimeQualityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeQualityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "DataType")
	if len(f) > 0 {
		return errors.New("DescribeRealtimeQualityRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeQualityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据类型
		Data []*RealtimeData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealtimeQualityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeQualityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeScaleRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间，24小时内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询的数据类型
	// UserNum：通话人数；
	// RoomNum：房间数
	DataType []*string `json:"DataType,omitempty" name:"DataType"`
}

func (r *DescribeRealtimeScaleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeScaleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "DataType")
	if len(f) > 0 {
		return errors.New("DescribeRealtimeScaleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealtimeScaleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的数据数组
		Data []*RealtimeData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealtimeScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeScaleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordStatisticRequest struct {
	*tchttp.BaseRequest

	// 查询开始日期，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束日期，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 应用ID，可不传。传应用ID时返回的是该应用的用量，不传时返回多个应用的合计值。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeRecordStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return errors.New("DescribeRecordStatisticRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordStatisticResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用的用量信息数组。
		SdkAppIdUsages []*SdkAppIdRecordUsage `json:"SdkAppIdUsages,omitempty" name:"SdkAppIdUsages"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoomInformationRequest struct {
	*tchttp.BaseRequest

	// 用户sdkappid
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 查询开始时间，14天内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 字符串房间号
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 分页index，从0开始（PageNumber和PageSize 其中一个不填均默认返回10条数据）
	PageNumber *string `json:"PageNumber,omitempty" name:"PageNumber"`

	// 分页大小（PageNumber和PageSize 其中一个不填均默认返回10条数据,最大不超过100）
	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeRoomInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return errors.New("DescribeRoomInformationRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoomInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回当页数据总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 房间信息列表
		RoomList []*RoomState `json:"RoomList,omitempty" name:"RoomList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoomInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrtcInteractiveTimeRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 应用ID，可不传。传应用ID时返回的是该应用的用量，不传时返回所有应用的合计值。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeTrtcInteractiveTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcInteractiveTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return errors.New("DescribeTrtcInteractiveTimeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrtcInteractiveTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用的用量信息数组。
		Usages []*OneSdkAppIdUsagesInfo `json:"Usages,omitempty" name:"Usages"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrtcInteractiveTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcInteractiveTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrtcMcuTranscodeTimeRequest struct {
	*tchttp.BaseRequest

	// 查询开始时间，格式为YYYY-MM-DD。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，格式为YYYY-MM-DD。
	// 单次查询统计区间最多不能超过31天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 应用ID，可不传。传应用ID时返回的是该应用的用量，不传时返回多个应用的合计值。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeTrtcMcuTranscodeTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcMcuTranscodeTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return errors.New("DescribeTrtcMcuTranscodeTimeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrtcMcuTranscodeTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用的用量信息数组。
		Usages []*OneSdkAppIdTranscodeTimeUsagesInfo `json:"Usages,omitempty" name:"Usages"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrtcMcuTranscodeTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcMcuTranscodeTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInformationRequest struct {
	*tchttp.BaseRequest

	// 通话 ID（唯一标识一次通话）： sdkappid_roomgString（房间号_createTime（房间创建时间，unix时间戳，单位为s）例：1400353843_218695_1590065777。通过 DescribeRoomInformation（查询房间列表）接口获取（链接：https://cloud.tencent.com/document/product/647/44050）
	CommId *string `json:"CommId,omitempty" name:"CommId"`

	// 查询开始时间，14天内。本地unix时间戳（1588031999s）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，本地unix时间戳（1588031999s）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 用户SDKAppID（1400188366）
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 需查询的用户数组，不填默认返回6个用户,最多可填6个用户
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// 设置分页index，从0开始（PageNumber和PageSize 其中一个不填均默认返回6条数据）
	PageNumber *string `json:"PageNumber,omitempty" name:"PageNumber"`

	// 设置分页大小（PageNumber和PageSize 其中一个不填均默认返回6条数据,PageSize最大不超过100）
	PageSize *string `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeUserInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "UserIds")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return errors.New("DescribeUserInformationRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的用户总条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 用户信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserList []*UserInformation `json:"UserList,omitempty" name:"UserList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DismissRoomByStrRoomIdRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DismissRoomByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return errors.New("DismissRoomByStrRoomIdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DismissRoomByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DismissRoomByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomByStrRoomIdResponse) FromJsonString(s string) error {
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return errors.New("DismissRoomRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EncodeParams struct {

	// 混流-输出流音频采样率。取值为[48000, 44100, 32000, 24000, 16000, 8000]，单位是Hz。
	AudioSampleRate *uint64 `json:"AudioSampleRate,omitempty" name:"AudioSampleRate"`

	// 混流-输出流音频码率。取值范围[8,500]，单位为kbps。
	AudioBitrate *uint64 `json:"AudioBitrate,omitempty" name:"AudioBitrate"`

	// 混流-输出流音频声道数，取值范围[1,2]，1表示混流输出音频为单声道，2表示混流输出音频为双声道。
	AudioChannels *uint64 `json:"AudioChannels,omitempty" name:"AudioChannels"`

	// 混流-输出流宽，音视频输出时必填。取值范围[0,1920]，单位为像素值。
	VideoWidth *uint64 `json:"VideoWidth,omitempty" name:"VideoWidth"`

	// 混流-输出流高，音视频输出时必填。取值范围[0,1080]，单位为像素值。
	VideoHeight *uint64 `json:"VideoHeight,omitempty" name:"VideoHeight"`

	// 混流-输出流码率，音视频输出时必填。取值范围[1,10000]，单位为kbps。
	VideoBitrate *uint64 `json:"VideoBitrate,omitempty" name:"VideoBitrate"`

	// 混流-输出流帧率，音视频输出时必填。取值范围[1,60]，表示混流的输出帧率可选范围为1到60fps。
	VideoFramerate *uint64 `json:"VideoFramerate,omitempty" name:"VideoFramerate"`

	// 混流-输出流gop，音视频输出时必填。取值范围[1,5]，单位为秒。
	VideoGop *uint64 `json:"VideoGop,omitempty" name:"VideoGop"`

	// 混流-输出流背景色，取值是十进制整数。常用的颜色有：
	// 红色：0xff0000，对应的十进制整数是16724736。
	// 黄色：0xffff00。对应的十进制整数是16776960。
	// 绿色：0x33cc00。对应的十进制整数是3394560。
	// 蓝色：0x0066ff。对应的十进制整数是26367。
	// 黑色：0x000000。对应的十进制整数是0。
	// 白色：0xFFFFFF。对应的十进制整数是16777215。
	// 灰色：0x999999。对应的十进制整数是10066329。
	BackgroundColor *uint64 `json:"BackgroundColor,omitempty" name:"BackgroundColor"`

	// 混流-输出流背景图片，取值为实时音视频控制台上传的图片ID。
	BackgroundImageId *uint64 `json:"BackgroundImageId,omitempty" name:"BackgroundImageId"`

	// 混流-输出流音频编码类型，取值范围[0,1, 2]，0为LC-AAC，1为HE-AAC，2为HE-AACv2。默认值为0。当音频编码设置为HE-AACv2时，只支持输出流音频声道数为双声道。HE-AAC和HE-AACv2支持的输出流音频采样率范围为[48000, 44100, 32000, 24000, 16000]
	AudioCodec *uint64 `json:"AudioCodec,omitempty" name:"AudioCodec"`
}

type EventList struct {

	// 数据内容
	Content []*EventMessage `json:"Content,omitempty" name:"Content"`

	// 发送端的userId
	PeerId *string `json:"PeerId,omitempty" name:"PeerId"`
}

type EventMessage struct {

	// 视频流类型：
	// 0：与视频无关的事件；
	// 2：视频为大画面；
	// 3：视频为小画面；
	// 7：视频为旁路画面；
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 事件上报的时间戳，unix时间（1589891188801ms)
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 事件Id：分为sdk的事件和webrtc的事件，详情见：附录/事件 ID 映射表：https://cloud.tencent.com/document/product/647/44916
	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`

	// 事件的第一个参数，如视频分辨率宽
	ParamOne *int64 `json:"ParamOne,omitempty" name:"ParamOne"`

	// 事件的第二个参数，如视频分辨率高
	ParamTwo *int64 `json:"ParamTwo,omitempty" name:"ParamTwo"`
}

type LayoutParams struct {

	// 混流布局模板ID，0为悬浮模板(默认);1为九宫格模板;2为屏幕分享模板;3为画中画模板;4为自定义模板。
	Template *uint64 `json:"Template,omitempty" name:"Template"`

	// 屏幕分享模板、悬浮模板、画中画模板中有效，代表大画面对应的用户ID。
	MainVideoUserId *string `json:"MainVideoUserId,omitempty" name:"MainVideoUserId"`

	// 屏幕分享模板、悬浮模板、画中画模板中有效，代表大画面对应的流类型，0为摄像头，1为屏幕分享。左侧大画面为web用户时此值填0。
	MainVideoStreamType *uint64 `json:"MainVideoStreamType,omitempty" name:"MainVideoStreamType"`

	// 画中画模板中有效，代表小画面的布局参数。
	SmallVideoLayoutParams *SmallVideoLayoutParams `json:"SmallVideoLayoutParams,omitempty" name:"SmallVideoLayoutParams"`

	// 屏幕分享模板有效。设置为1时代表大画面居右，小画面居左布局。默认为0。
	MainVideoRightAlign *uint64 `json:"MainVideoRightAlign,omitempty" name:"MainVideoRightAlign"`

	// 悬浮模板、九宫格、屏幕分享模板有效。设置此参数后，输出流混合此参数中包含用户的音视频，以及其他用户的纯音频。最多可设置16个用户。
	MixVideoUids []*string `json:"MixVideoUids,omitempty" name:"MixVideoUids"`

	// 自定义模板中有效，指定用户视频在混合画面中的位置。
	PresetLayoutConfig []*PresetLayoutConfig `json:"PresetLayoutConfig,omitempty" name:"PresetLayoutConfig"`

	// 自定义模板中有效，设置为1时代表启用占位图功能，0时代表不启用占位图功能，默认为0。启用占位图功能时，在预设位置的用户没有上行视频时可显示对应的占位图。
	PlaceHolderMode *uint64 `json:"PlaceHolderMode,omitempty" name:"PlaceHolderMode"`

	// 悬浮模板、九宫格、屏幕分享模板生效，用于控制纯音频上行是否占用画面布局位置。设置为0是代表后台默认处理方式，悬浮小画面占布局位置，九宫格画面占布局位置、屏幕分享小画面不占布局位置；设置为1时代表纯音频上行占布局位置；设置为2时代表纯音频上行不占布局位置。默认为0。
	PureAudioHoldPlaceMode *uint64 `json:"PureAudioHoldPlaceMode,omitempty" name:"PureAudioHoldPlaceMode"`

	// 水印参数。
	WaterMarkParams *WaterMarkParams `json:"WaterMarkParams,omitempty" name:"WaterMarkParams"`
}

type ModifyPictureRequest struct {
	*tchttp.BaseRequest

	// 图片id
	PictureId *uint64 `json:"PictureId,omitempty" name:"PictureId"`

	// 应用id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 图片长度
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 图片宽度
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 显示位置x轴方向
	XPosition *uint64 `json:"XPosition,omitempty" name:"XPosition"`

	// 显示位置y轴方向
	YPosition *uint64 `json:"YPosition,omitempty" name:"YPosition"`
}

func (r *ModifyPictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PictureId")
	delete(f, "SdkAppId")
	delete(f, "Height")
	delete(f, "Width")
	delete(f, "XPosition")
	delete(f, "YPosition")
	if len(f) > 0 {
		return errors.New("ModifyPictureRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPictureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OneSdkAppIdTranscodeTimeUsagesInfo struct {

	// 旁路转码时长查询结果数组
	SdkAppIdTranscodeTimeUsages []*SdkAppIdTrtcMcuTranscodeTimeUsage `json:"SdkAppIdTranscodeTimeUsages,omitempty" name:"SdkAppIdTranscodeTimeUsages"`

	// 查询记录数量
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// 所查询的应用ID，可能值为:1-应用的应用ID，2-total，显示为total则表示查询的是所有应用的用量合计值。
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type OneSdkAppIdUsagesInfo struct {

	// 该 SdkAppId 对应的用量记录数长度
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// 用量数组
	SdkAppIdTrtcTimeUsages []*SdkAppIdTrtcUsage `json:"SdkAppIdTrtcTimeUsages,omitempty" name:"SdkAppIdTrtcTimeUsages"`

	// 应用ID
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type OutputParams struct {

	// 直播流 ID，由用户自定义设置，该流 ID 不能与用户旁路的流 ID 相同。
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// 取值范围[0,1]， 填0：直播流为音视频(默认); 填1：直播流为纯音频
	PureAudioStream *uint64 `json:"PureAudioStream,omitempty" name:"PureAudioStream"`

	// 自定义录制文件名称前缀。请先在实时音视频控制台开通录制功能，https://cloud.tencent.com/document/product/647/50768
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 取值范围[0,1]，填0无实际含义; 填1：指定录制文件格式为mp3。此参数不建议使用，建议在实时音视频控制台配置纯音频录制模板。
	RecordAudioOnly *uint64 `json:"RecordAudioOnly,omitempty" name:"RecordAudioOnly"`
}

type PictureInfo struct {

	// 图片长度
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 图片宽度
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 显示位置x轴方向
	XPosition *uint64 `json:"XPosition,omitempty" name:"XPosition"`

	// 显示位置y轴方向
	YPosition *uint64 `json:"YPosition,omitempty" name:"YPosition"`

	// 应用id
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 图片id
	PictureId *uint64 `json:"PictureId,omitempty" name:"PictureId"`
}

type PresetLayoutConfig struct {

	// 指定显示在该画面上的用户ID。如果不指定用户ID，会按照用户加入房间的顺序自动匹配PresetLayoutConfig中的画面设置。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 当该画面指定用户时，代表用户的流类型。0为摄像头，1为屏幕分享。小画面为web用户时此值填0。
	StreamType *uint64 `json:"StreamType,omitempty" name:"StreamType"`

	// 该画面在输出时的宽度，单位为像素值，不填默认为0。
	ImageWidth *uint64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

	// 该画面在输出时的高度，单位为像素值，不填默认为0。
	ImageHeight *uint64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

	// 该画面在输出时的X偏移，单位为像素值，LocationX与ImageWidth之和不能超过混流输出的总宽度，不填默认为0。
	LocationX *uint64 `json:"LocationX,omitempty" name:"LocationX"`

	// 该画面在输出时的Y偏移，单位为像素值，LocationY与ImageHeight之和不能超过混流输出的总高度，不填默认为0。
	LocationY *uint64 `json:"LocationY,omitempty" name:"LocationY"`

	// 该画面在输出时的层级，不填默认为0。
	ZOrder *uint64 `json:"ZOrder,omitempty" name:"ZOrder"`

	// 该画面在输出时的显示模式：0为裁剪，1为缩放，2为缩放并显示黑底。不填默认为0。
	RenderMode *uint64 `json:"RenderMode,omitempty" name:"RenderMode"`

	// 该当前位置用户混入的流类型：0为混入音视频，1为只混入视频，2为只混入音频。默认为0，建议配合指定用户ID使用。
	MixInputType *uint64 `json:"MixInputType,omitempty" name:"MixInputType"`

	// 占位图ID。启用占位图功能时，在当前位置的用户没有上行视频时显示占位图。占位图在实时音视频控制台上传并生成，https://cloud.tencent.com/document/product/647/50769
	PlaceImageId *uint64 `json:"PlaceImageId,omitempty" name:"PlaceImageId"`
}

type PublishCdnParams struct {

	// 腾讯云直播BizId。
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 第三方CDN转推的目的地址，同时只支持转推一个第三方CDN地址。
	PublishCdnUrls []*string `json:"PublishCdnUrls,omitempty" name:"PublishCdnUrls"`
}

type QualityData struct {

	// 数据内容
	Content []*TimeValue `json:"Content,omitempty" name:"Content"`

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
	Content []*TimeValue `json:"Content,omitempty" name:"Content"`

	// 数据类型字段
	DataType *string `json:"DataType,omitempty" name:"DataType"`
}

type RecordUsage struct {

	// 本组数据对应的时间点，格式如:2020-09-07或2020-09-07 00:05:05。
	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`

	// 视频时长-标清SD，单位：秒。
	Class1VideoTime *uint64 `json:"Class1VideoTime,omitempty" name:"Class1VideoTime"`

	// 视频时长-高清HD，单位：秒。
	Class2VideoTime *uint64 `json:"Class2VideoTime,omitempty" name:"Class2VideoTime"`

	// 视频时长-超清HD，单位：秒。
	Class3VideoTime *uint64 `json:"Class3VideoTime,omitempty" name:"Class3VideoTime"`

	// 语音时长，单位：秒。
	AudioTime *uint64 `json:"AudioTime,omitempty" name:"AudioTime"`
}

type RemoveUserByStrRoomIdRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 要移出的用户列表，最多10个。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *RemoveUserByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserIds")
	if len(f) > 0 {
		return errors.New("RemoveUserByStrRoomIdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RemoveUserByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveUserByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveUserRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 要移出的用户列表，最多10个。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *RemoveUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserIds")
	if len(f) > 0 {
		return errors.New("RemoveUserRequest has unknown keys!")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 房间人数，用户重复进入同一个房间为1次
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserNumber *uint64 `json:"UserNumber,omitempty" name:"UserNumber"`

	// 房间人次，用户每次进入房间为一次
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserCount *uint64 `json:"UserCount,omitempty" name:"UserCount"`

	// sdkappid下一天内的房间数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomNumbers *uint64 `json:"RoomNumbers,omitempty" name:"RoomNumbers"`
}

type SdkAppIdRecordUsage struct {

	// SdkAppId的值。
	SdkAppId *string `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 统计的时间点数据。
	Usages []*RecordUsage `json:"Usages,omitempty" name:"Usages"`
}

type SdkAppIdTrtcMcuTranscodeTimeUsage struct {

	// 本组数据对应的时间点，格式如：2020-09-07或2020-09-07 00:05:05。
	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`

	// 语音时长，单位：秒。
	AudioTime *uint64 `json:"AudioTime,omitempty" name:"AudioTime"`

	// 视频时长-标清SD，单位：秒。
	VideoTimeSd *uint64 `json:"VideoTimeSd,omitempty" name:"VideoTimeSd"`

	// 视频时长-高清HD，单位：秒。
	VideoTimeHd *uint64 `json:"VideoTimeHd,omitempty" name:"VideoTimeHd"`

	// 视频时长-全高清FHD，单位：秒。
	VideoTimeFhd *uint64 `json:"VideoTimeFhd,omitempty" name:"VideoTimeFhd"`
}

type SdkAppIdTrtcUsage struct {

	// 本组数据对应的时间点，格式如：2020-09-07或2020-09-07 00:05:05。
	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`

	// 语音时长，单位：秒。
	AudioTime *uint64 `json:"AudioTime,omitempty" name:"AudioTime"`

	// 音视频时长，单位：秒。
	// 2019年10月11日前注册，没有变更为 [新计费模式](https://cloud.tencent.com/document/product/647/17157) 的用户才会返回此值。
	AudioVideoTime *uint64 `json:"AudioVideoTime,omitempty" name:"AudioVideoTime"`

	// 视频时长-标清SD，单位：秒。
	VideoTimeSd *uint64 `json:"VideoTimeSd,omitempty" name:"VideoTimeSd"`

	// 视频时长-高清HD，单位：秒。
	VideoTimeHd *uint64 `json:"VideoTimeHd,omitempty" name:"VideoTimeHd"`

	// 视频时长-超清HD，单位：秒。
	VideoTimeHdp *uint64 `json:"VideoTimeHdp,omitempty" name:"VideoTimeHdp"`
}

type SmallVideoLayoutParams struct {

	// 代表小画面对应的用户ID。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 代表小画面对应的流类型，0为摄像头，1为屏幕分享。小画面为web用户时此值填0。
	StreamType *uint64 `json:"StreamType,omitempty" name:"StreamType"`

	// 小画面在输出时的宽度，单位为像素值，不填默认为0。
	ImageWidth *uint64 `json:"ImageWidth,omitempty" name:"ImageWidth"`

	// 小画面在输出时的高度，单位为像素值，不填默认为0。
	ImageHeight *uint64 `json:"ImageHeight,omitempty" name:"ImageHeight"`

	// 小画面在输出时的X偏移，单位为像素值，LocationX与ImageWidth之和不能超过混流输出的总宽度，不填默认为0。
	LocationX *uint64 `json:"LocationX,omitempty" name:"LocationX"`

	// 小画面在输出时的Y偏移，单位为像素值，LocationY与ImageHeight之和不能超过混流输出的总高度，不填默认为0。
	LocationY *uint64 `json:"LocationY,omitempty" name:"LocationY"`
}

type StartMCUMixTranscodeByStrRoomIdRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 字符串房间号。
	StrRoomId *string `json:"StrRoomId,omitempty" name:"StrRoomId"`

	// 混流输出控制参数。
	OutputParams *OutputParams `json:"OutputParams,omitempty" name:"OutputParams"`

	// 混流输出编码参数。
	EncodeParams *EncodeParams `json:"EncodeParams,omitempty" name:"EncodeParams"`

	// 混流输出布局参数。
	LayoutParams *LayoutParams `json:"LayoutParams,omitempty" name:"LayoutParams"`

	// 第三方CDN转推参数。
	PublishCdnParams *PublishCdnParams `json:"PublishCdnParams,omitempty" name:"PublishCdnParams"`
}

func (r *StartMCUMixTranscodeByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMCUMixTranscodeByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StrRoomId")
	delete(f, "OutputParams")
	delete(f, "EncodeParams")
	delete(f, "LayoutParams")
	delete(f, "PublishCdnParams")
	if len(f) > 0 {
		return errors.New("StartMCUMixTranscodeByStrRoomIdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartMCUMixTranscodeByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMCUMixTranscodeByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMCUMixTranscodeByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartMCUMixTranscodeRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 混流输出控制参数。
	OutputParams *OutputParams `json:"OutputParams,omitempty" name:"OutputParams"`

	// 混流输出编码参数。
	EncodeParams *EncodeParams `json:"EncodeParams,omitempty" name:"EncodeParams"`

	// 混流输出布局参数。
	LayoutParams *LayoutParams `json:"LayoutParams,omitempty" name:"LayoutParams"`

	// 第三方CDN转推参数。
	PublishCdnParams *PublishCdnParams `json:"PublishCdnParams,omitempty" name:"PublishCdnParams"`
}

func (r *StartMCUMixTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMCUMixTranscodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "OutputParams")
	delete(f, "EncodeParams")
	delete(f, "LayoutParams")
	delete(f, "PublishCdnParams")
	if len(f) > 0 {
		return errors.New("StartMCUMixTranscodeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartMCUMixTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMCUMixTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMCUMixTranscodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopMCUMixTranscodeByStrRoomIdRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 字符串房间号。
	StrRoomId *string `json:"StrRoomId,omitempty" name:"StrRoomId"`
}

func (r *StopMCUMixTranscodeByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMCUMixTranscodeByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StrRoomId")
	if len(f) > 0 {
		return errors.New("StopMCUMixTranscodeByStrRoomIdRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopMCUMixTranscodeByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMCUMixTranscodeByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMCUMixTranscodeByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopMCUMixTranscodeRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *StopMCUMixTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMCUMixTranscodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return errors.New("StopMCUMixTranscodeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopMCUMixTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMCUMixTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMCUMixTranscodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TimeValue struct {

	// 时间，unix时间戳（1590065877s)
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 当前时间返回参数取值，如（bigvCapFps在1590065877取值为0，则Value：0 ）
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type UserInformation struct {

	// 房间号
	RoomStr *string `json:"RoomStr,omitempty" name:"RoomStr"`

	// 用户Id
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户进房时间
	JoinTs *uint64 `json:"JoinTs,omitempty" name:"JoinTs"`

	// 用户退房时间，用户没有退房则返回当前时间
	LeaveTs *uint64 `json:"LeaveTs,omitempty" name:"LeaveTs"`

	// 终端类型
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Sdk版本号
	SdkVersion *string `json:"SdkVersion,omitempty" name:"SdkVersion"`

	// 客户端IP地址
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 判断用户是否已经离开房间
	Finished *bool `json:"Finished,omitempty" name:"Finished"`
}

type WaterMarkParams struct {

	// 混流-水印图片ID。取值为实时音视频控制台上传的图片ID。
	WaterMarkId *uint64 `json:"WaterMarkId,omitempty" name:"WaterMarkId"`

	// 混流-水印宽。单位为像素值。
	WaterMarkWidth *uint64 `json:"WaterMarkWidth,omitempty" name:"WaterMarkWidth"`

	// 混流-水印高。单位为像素值。
	WaterMarkHeight *uint64 `json:"WaterMarkHeight,omitempty" name:"WaterMarkHeight"`

	// 水印在输出时的X偏移。单位为像素值。
	LocationX *uint64 `json:"LocationX,omitempty" name:"LocationX"`

	// 水印在输出时的Y偏移。单位为像素值。
	LocationY *uint64 `json:"LocationY,omitempty" name:"LocationY"`
}
