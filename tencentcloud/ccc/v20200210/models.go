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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSDKLoginTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SeatUserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSDKLoginTokenRequest has unknown keys!", "")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSDKLoginTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStaffRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 客服信息，个数不超过 10
	Staffs []*SeatUserInfo `json:"Staffs,omitempty" name:"Staffs"`
}

func (r *CreateStaffRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaffRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Staffs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStaffRequest has unknown keys!", "")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChatMessagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CdrId")
	delete(f, "InstanceId")
	delete(f, "SdkAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChatMessagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChatMessagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 消息列表
		Messages []*MessageBody `json:"Messages,omitempty" name:"Messages"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeChatMessagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIMCdrsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimestamp")
	delete(f, "EndTimestamp")
	delete(f, "InstanceId")
	delete(f, "SdkAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIMCdrsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIMCdrsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 服务记录列表
		IMCdrs []*IMCdrInfo `json:"IMCdrs,omitempty" name:"IMCdrs"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIMCdrsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIMCdrsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePSTNActiveSessionListRequest struct {
	*tchttp.BaseRequest

	// 应用 ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 数据偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回的数据条数，最大 25
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePSTNActiveSessionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePSTNActiveSessionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePSTNActiveSessionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePSTNActiveSessionListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表总条数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 列表内容
		Sessions []*PSTNSessionInfo `json:"Sessions,omitempty" name:"Sessions"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePSTNActiveSessionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePSTNActiveSessionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSeatUserListRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSeatUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSeatUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSeatUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSeatUserListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 此实例的坐席用户总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 坐席用户信息列表
		SeatUsers []*SeatUserInfo `json:"SeatUsers,omitempty" name:"SeatUsers"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSeatUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSeatUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSkillGroupInfoListRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 技能组ID，查询单个技能组时使用
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 查询修改时间大于等于ModifiedTime的技能组时使用
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

func (r *DescribeSkillGroupInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillGroupInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "SkillGroupId")
	delete(f, "ModifiedTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillGroupInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSkillGroupInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 技能组总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 技能组信息列表
		SkillGroupList []*SkillGroupInfoItem `json:"SkillGroupList,omitempty" name:"SkillGroupList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSkillGroupInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillGroupInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStaffInfoListRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 坐席账号，查询单个坐席时使用
	StaffMail *string `json:"StaffMail,omitempty" name:"StaffMail"`

	// 查询修改时间大于等于ModifiedTime的坐席时使用
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

func (r *DescribeStaffInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "StaffMail")
	delete(f, "ModifiedTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStaffInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStaffInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 坐席用户总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 坐席用户信息列表
		StaffList []*StaffInfo `json:"StaffList,omitempty" name:"StaffList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStaffInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTelCallInfoRequest struct {
	*tchttp.BaseRequest

	// 起始时间戳，Unix 时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 时间戳，查询时间范围最大为90天
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 应用ID列表，多个ID时，返回值为多个ID使用总和
	SdkAppIdList []*int64 `json:"SdkAppIdList,omitempty" name:"SdkAppIdList"`
}

func (r *DescribeTelCallInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCallInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeStamp")
	delete(f, "EndTimeStamp")
	delete(f, "SdkAppIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelCallInfoRequest has unknown keys!", "")
	}
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCallInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTelCdrRequest struct {
	*tchttp.BaseRequest

	// 起始时间戳，Unix 时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 时间戳
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 返回数据条数，上限（deprecated）
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移（deprecated）
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 实例 ID（deprecated）
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用 ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

func (r *DescribeTelCdrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCdrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeStamp")
	delete(f, "EndTimeStamp")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceId")
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelCdrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTelCdrResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 话单记录总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 话单记录
		TelCdrs []*TelCdrInfo `json:"TelCdrs,omitempty" name:"TelCdrs"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTelCdrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCdrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTelSessionRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 会话ID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

func (r *DescribeTelSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTelSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 会话信息
		Session *PSTNSession `json:"Session,omitempty" name:"Session"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTelSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelSessionResponse) FromJsonString(s string) error {
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

type IVRKeyPressedElement struct {

	// 按键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 按键关联的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`
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
	Messages []*Message `json:"Messages,omitempty" name:"Messages"`
}

type PSTNSession struct {

	// 会话 ID
	SessionID *string `json:"SessionID,omitempty" name:"SessionID"`

	// 会话临时房间 ID
	RoomID *string `json:"RoomID,omitempty" name:"RoomID"`

	// 主叫
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 被叫
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 开始时间，Unix 时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 振铃时间，Unix 时间戳
	RingTimestamp *int64 `json:"RingTimestamp,omitempty" name:"RingTimestamp"`

	// 接听时间，Unix 时间戳
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitempty" name:"AcceptTimestamp"`

	// 坐席邮箱
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 坐席工号
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`

	// 会话状态
	// ringing 振铃中
	// seatJoining  等待坐席接听
	// inProgress 进行中
	// finished 已完成
	SessionStatus *string `json:"SessionStatus,omitempty" name:"SessionStatus"`

	// 会话呼叫方向， 0 呼入 | 1 - 呼出
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// 转外线使用的号码（转外线主叫）
	OutBoundCaller *string `json:"OutBoundCaller,omitempty" name:"OutBoundCaller"`

	// 转外线被叫
	OutBoundCallee *string `json:"OutBoundCallee,omitempty" name:"OutBoundCallee"`
}

type PSTNSessionInfo struct {

	// 会话 ID
	SessionID *string `json:"SessionID,omitempty" name:"SessionID"`

	// 会话临时房间 ID
	RoomID *string `json:"RoomID,omitempty" name:"RoomID"`

	// 主叫
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 被叫
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 开始时间，Unix 时间戳
	StartTimestamp *string `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 接听时间，Unix 时间戳
	AcceptTimestamp *string `json:"AcceptTimestamp,omitempty" name:"AcceptTimestamp"`

	// 坐席邮箱
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 坐席工号
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`

	// 坐席状态 inProgress 进行中
	SessionStatus *string `json:"SessionStatus,omitempty" name:"SessionStatus"`

	// 会话呼叫方向， 0 呼入 | 1 - 呼出
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// 振铃时间，Unix 时间戳
	RingTimestamp *int64 `json:"RingTimestamp,omitempty" name:"RingTimestamp"`
}

type SeatUserInfo struct {

	// 坐席名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 坐席邮箱
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 坐席电话号码（带0086前缀）
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 坐席昵称
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 坐席关联的技能组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupNameList []*string `json:"SkillGroupNameList,omitempty" name:"SkillGroupNameList"`

	// 工号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`
}

type ServeParticipant struct {

	// 坐席邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 坐席电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 振铃时间戳，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	RingTimestamp *int64 `json:"RingTimestamp,omitempty" name:"RingTimestamp"`

	// 接听时间戳，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitempty" name:"AcceptTimestamp"`

	// 结束时间戳，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndedTimestamp *int64 `json:"EndedTimestamp,omitempty" name:"EndedTimestamp"`

	// 录音 ID，能够索引到坐席侧的录音
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 参与者类型，"staffSeat", "outboundSeat", "staffPhoneSeat"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 转接来源坐席信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferFrom *string `json:"TransferFrom,omitempty" name:"TransferFrom"`

	// 转接去向坐席信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferTo *string `json:"TransferTo,omitempty" name:"TransferTo"`

	// 转接去向参与者类型，取值与 Type 一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferToType *string `json:"TransferToType,omitempty" name:"TransferToType"`

	// 技能组 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 结束状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndStatusString *string `json:"EndStatusString,omitempty" name:"EndStatusString"`

	// 录音 URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordURL *string `json:"RecordURL,omitempty" name:"RecordURL"`

	// 参与者序号，从 0 开始
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sequence *int64 `json:"Sequence,omitempty" name:"Sequence"`

	// 开始时间戳，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 技能组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`
}

type SkillGroupInfoItem struct {

	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`

	// 类型：IM、TEL、ALL（全媒体）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 会话分配策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoutePolicy *string `json:"RoutePolicy,omitempty" name:"RoutePolicy"`

	// 会话分配是否优先上次服务坐席
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsingLastSeat *int64 `json:"UsingLastSeat,omitempty" name:"UsingLastSeat"`

	// 单客服最大并发数（电话类型默认1）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" name:"MaxConcurrency"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTimestamp *int64 `json:"LastModifyTimestamp,omitempty" name:"LastModifyTimestamp"`
}

type SkillGroupItem struct {

	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`

	// 优先级
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 类型：IM、TEL、ALL（全媒体）
	Type *string `json:"Type,omitempty" name:"Type"`
}

type StaffInfo struct {

	// 坐席名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 坐席邮箱
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 坐席电话号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 坐席昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 坐席工号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`

	// 所属技能组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupList []*SkillGroupItem `json:"SkillGroupList,omitempty" name:"SkillGroupList"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTimestamp *int64 `json:"LastModifyTimestamp,omitempty" name:"LastModifyTimestamp"`
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

	// 结束状态
	// 0	错误
	// 1	正常结束
	// 2	未接通
	// 17	坐席未接
	// 100	黑名单
	// 101	坐席转接
	// 102	IVR 期间用户放弃
	// 103	会话排队期间用户放弃
	// 104	会话振铃期间用户放弃
	// 105	无坐席在线
	// 106	非工作时间
	// 107	IVR后直接结束
	// 201	未知状态
	// 202	未接听
	// 203	拒接挂断
	// 204	关机
	// 205	空号
	// 206	通话中
	// 207	欠费
	// 208	运营商线路异常
	// 209	主叫取消
	// 210	不在服务区
	EndStatus *int64 `json:"EndStatus,omitempty" name:"EndStatus"`

	// 技能组名称
	SkillGroup *string `json:"SkillGroup,omitempty" name:"SkillGroup"`

	// 主叫归属地
	CallerLocation *string `json:"CallerLocation,omitempty" name:"CallerLocation"`

	// IVR 阶段耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	IVRDuration *int64 `json:"IVRDuration,omitempty" name:"IVRDuration"`

	// 振铃时间戳，UNIX 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	RingTimestamp *int64 `json:"RingTimestamp,omitempty" name:"RingTimestamp"`

	// 接听时间戳，UNIX 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitempty" name:"AcceptTimestamp"`

	// 结束时间戳，UNIX 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndedTimestamp *int64 `json:"EndedTimestamp,omitempty" name:"EndedTimestamp"`

	// IVR 按键信息 ，e.g. ["1","2","3"]
	// 注意：此字段可能返回 null，表示取不到有效值。
	IVRKeyPressed []*string `json:"IVRKeyPressed,omitempty" name:"IVRKeyPressed"`

	// 挂机方 seat 坐席 user 用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	HungUpSide *string `json:"HungUpSide,omitempty" name:"HungUpSide"`

	// 服务参与者列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServeParticipants []*ServeParticipant `json:"ServeParticipants,omitempty" name:"ServeParticipants"`

	// 技能组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// error                   错误
	// ok                       正常结束
	// unconnected      未接通
	// seatGiveUp         坐席未接
	// blackList             黑名单
	// seatForward       坐席转接
	// ivrGiveUp           IVR 期间用户放弃
	// waitingGiveUp   会话排队期间用户放弃
	// ringingGiveUp   会话振铃期间用户放弃
	// noSeatOnline     无坐席在线
	// notWorkTime     非工作时间
	// ivrEnd                 IVR后直接结束
	// unknown            未知状态
	// notAnswer          未接听
	// userReject          拒接挂断
	// powerOff            关机
	// numberNotExist  空号
	// busy                    通话中
	// outOfCredit        欠费
	// operatorError     运营商线路异常
	// callerCancel        主叫取消
	// notInService       不在服务区
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndStatusString *string `json:"EndStatusString,omitempty" name:"EndStatusString"`

	// 会话开始时间戳，UNIX 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 进入排队时间，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueuedTimestamp *int64 `json:"QueuedTimestamp,omitempty" name:"QueuedTimestamp"`

	// 后置IVR按键信息（e.g. [{"Key":"1","Label":"非常满意"}]）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostIVRKeyPressed []*IVRKeyPressedElement `json:"PostIVRKeyPressed,omitempty" name:"PostIVRKeyPressed"`

	// 排队技能组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueuedSkillGroupId *int64 `json:"QueuedSkillGroupId,omitempty" name:"QueuedSkillGroupId"`

	// 会话 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}
