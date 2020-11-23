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

package v20200210

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateSDKLoginTokenRequest struct {
	*tchttp.BaseRequest

	// 应用ID。
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 坐席账号。
	SeatUserId *string `json:"SeatUserId,omitempty" name:"SeatUserId"`
}

func (r *CreateSDKLoginTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSDKLoginTokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSDKLoginTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SDK 登录 Token。
		Token *string `json:"Token,omitempty" name:"Token"`

		// 过期时间戳，Unix 时间戳。
		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

		// SDK 加载路径会随着 SDK 的发布而变动。
		SdkURL *string `json:"SdkURL,omitempty" name:"SdkURL"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSDKLoginTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSDKLoginTokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStaffRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 客服信息，个数不超过 10
	Staffs []*SeatUserInfo `json:"Staffs,omitempty" name:"Staffs" list`
}

func (r *CreateStaffRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStaffRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStaffResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStaffResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStaffResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeChatMessagesRequest struct {
	*tchttp.BaseRequest

	// 服务记录ID
	CdrId *string `json:"CdrId,omitempty" name:"CdrId"`

	// 实例ID
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 返回记录条数 最大为100默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回记录偏移 默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1为从早到晚，2为从晚到早，默认为2
	Order *int64 `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeChatMessagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeChatMessagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeChatMessagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 消息列表
		Messages []*MessageBody `json:"Messages,omitempty" name:"Messages" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeChatMessagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeChatMessagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIMCdrsRequest struct {
	*tchttp.BaseRequest

	// 起始时间
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 结束时间
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" name:"EndTimestamp"`

	// 实例ID
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 返回记录条数 最大为100默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回记录偏移 默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1为全媒体，2为文本客服，不填则查询全部
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeIMCdrsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIMCdrsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIMCdrsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 服务记录列表
		IMCdrs []*IMCdrInfo `json:"IMCdrs,omitempty" name:"IMCdrs" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIMCdrsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIMCdrsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTelCallInfoRequest struct {
	*tchttp.BaseRequest

	// 起始时间戳，Unix 时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 时间戳
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 应用ID列表，多个ID时，返回值为多个ID使用总和
	SdkAppIdList []*int64 `json:"SdkAppIdList,omitempty" name:"SdkAppIdList" list`
}

func (r *DescribeTelCallInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTelCallInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTelCallInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 电话呼出统计分钟数
		TelCallOutCount *int64 `json:"TelCallOutCount,omitempty" name:"TelCallOutCount"`

		// 电话呼入统计分钟数
		TelCallInCount *int64 `json:"TelCallInCount,omitempty" name:"TelCallInCount"`

		// 坐席使用统计个数
		SeatUsedCount *int64 `json:"SeatUsedCount,omitempty" name:"SeatUsedCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTelCallInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTelCallInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTelCdrRequest struct {
	*tchttp.BaseRequest

	// 起始时间戳，Unix 时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 时间戳
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 返回记录条数，上限 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 实例 ID
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用ID。
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *DescribeTelCdrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTelCdrRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTelCdrResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 话单记录总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 话单记录
		TelCdrs []*TelCdrInfo `json:"TelCdrs,omitempty" name:"TelCdrs" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTelCdrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTelCdrResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IMCdrInfo struct {

	// 服务记录ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 服务时长秒数
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 结束状态
	EndStatus *int64 `json:"EndStatus,omitempty" name:"EndStatus"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 服务类型 1为全媒体，2为文本客服
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 客服ID
	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`

	// 服务时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type Message struct {

	// 消息类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 消息内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

type MessageBody struct {

	// 消息时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 发消息的用户ID
	From *string `json:"From,omitempty" name:"From"`

	// 消息列表
	Messages []*Message `json:"Messages,omitempty" name:"Messages" list`
}

type SeatUserInfo struct {

	// 坐席名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 坐席邮箱
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 坐席电话号码
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 坐席昵称
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 坐席关联的技能组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupNameList []*string `json:"SkillGroupNameList,omitempty" name:"SkillGroupNameList" list`

	// 工号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`
}

type TelCdrInfo struct {

	// 主叫号码
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 被叫号码
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 呼叫发起时间戳，Unix 时间戳
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 呼入呼出方向 0 呼入 1 呼出
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// 通话时长
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 录音信息
	RecordURL *string `json:"RecordURL,omitempty" name:"RecordURL"`

	// 坐席信息
	SeatUser *SeatUserInfo `json:"SeatUser,omitempty" name:"SeatUser"`

	// 结束状态 0 未知 1 正常通话 2 未接通
	EndStatus *int64 `json:"EndStatus,omitempty" name:"EndStatus"`

	// 技能组
	SkillGroup *string `json:"SkillGroup,omitempty" name:"SkillGroup"`

	// 主叫归属地
	CallerLocation *string `json:"CallerLocation,omitempty" name:"CallerLocation"`
}
