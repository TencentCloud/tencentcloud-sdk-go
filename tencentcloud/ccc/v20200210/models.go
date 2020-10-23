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

type DescribeTelCdrRequest struct {
	*tchttp.BaseRequest

	// 实例 ID
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 起始时间戳，Unix 时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 时间戳
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 返回记录条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
