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

package v20220817

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateRoomRequestParams struct {
	// 房间名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 预定的房间开始时间，unix时间戳。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// 最大连麦人数（不包括老师）。取值范围[0, 17)
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	// coteaching 双师
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// 老师ID。
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 进入房间时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（默认值）
	// 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// 禁止录制。可以有以下取值：
	// 0 不禁止录制（默认值）
	// 1 禁止录制
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// 助教Id列表。
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// 录制布局。
	RecordLayout *uint64 `json:"RecordLayout,omitempty" name:"RecordLayout"`
}

type CreateRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 预定的房间开始时间，unix时间戳。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// 最大连麦人数（不包括老师）。取值范围[0, 17)
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	// coteaching 双师
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// 老师ID。
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 进入房间时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（默认值）
	// 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// 禁止录制。可以有以下取值：
	// 0 不禁止录制（默认值）
	// 1 禁止录制
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// 助教Id列表。
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// 录制布局。
	RecordLayout *uint64 `json:"RecordLayout,omitempty" name:"RecordLayout"`
}

func (r *CreateRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "Resolution")
	delete(f, "MaxMicNumber")
	delete(f, "SubType")
	delete(f, "TeacherId")
	delete(f, "AutoMic")
	delete(f, "AudioQuality")
	delete(f, "DisableRecord")
	delete(f, "Assistants")
	delete(f, "RecordLayout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoomResponseParams struct {
	// 房间ID。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRoomResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoomResponseParams `json:"Response"`
}

func (r *CreateRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSupervisorRequestParams struct {

}

type CreateSupervisorRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateSupervisorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSupervisorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSupervisorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSupervisorResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSupervisorResponse struct {
	*tchttp.BaseResponse
	Response *CreateSupervisorResponseParams `json:"Response"`
}

func (r *CreateSupervisorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSupervisorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomRequestParams struct {
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

type DescribeRoomRequest struct {
	*tchttp.BaseRequest
	
	// 房间Id。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DescribeRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomResponseParams struct {
	// 房间名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 预定的房间开始时间，unix时间戳。
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 预定的房间结束时间，unix时间戳。
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 老师ID。
	TeacherId *string `json:"TeacherId,omitempty" name:"TeacherId"`

	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分辨率。可以有如下取值：
	// 1 标清
	// 2 高清
	// 3 全高清
	Resolution *uint64 `json:"Resolution,omitempty" name:"Resolution"`

	// 最大连麦人数（不包括老师）。取值范围[0, 17)
	MaxMicNumber *uint64 `json:"MaxMicNumber,omitempty" name:"MaxMicNumber"`

	// 进入房间时是否自动连麦。可以有以下取值：
	// 0 不自动连麦（默认值）
	// 1 自动连麦
	AutoMic *uint64 `json:"AutoMic,omitempty" name:"AutoMic"`

	// 高音质模式。可以有以下取值：
	// 0 不开启高音质（默认值）
	// 1 开启高音质
	AudioQuality *uint64 `json:"AudioQuality,omitempty" name:"AudioQuality"`

	// 房间子类型，可以有以下取值：
	// videodoc 文档+视频
	// video 纯视频
	// coteaching 双师
	SubType *string `json:"SubType,omitempty" name:"SubType"`

	// 禁止录制。可以有以下取值：
	// 0 不禁止录制（默认值）
	// 1 禁止录制
	DisableRecord *uint64 `json:"DisableRecord,omitempty" name:"DisableRecord"`

	// 助教Id列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Assistants []*string `json:"Assistants,omitempty" name:"Assistants"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRoomResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoomResponseParams `json:"Response"`
}

func (r *DescribeRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRequestParams struct {
	// 用户Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserResponseParams struct {
	// 应用Id。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户昵称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户头像Url。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserResponseParams `json:"Response"`
}

func (r *DescribeUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginOriginIdRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`
}

type LoginOriginIdRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`
}

func (r *LoginOriginIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginOriginIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "OriginId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LoginOriginIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginOriginIdResponseParams struct {
	// 用户Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 登录/注册成功后返回登录态token。有效期7天。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type LoginOriginIdResponse struct {
	*tchttp.BaseResponse
	Response *LoginOriginIdResponseParams `json:"Response"`
}

func (r *LoginOriginIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginOriginIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginUserRequestParams struct {
	// 注册获取的用户id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type LoginUserRequest struct {
	*tchttp.BaseRequest
	
	// 注册获取的用户id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *LoginUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LoginUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LoginUserResponseParams struct {
	// 用户Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 登录/注册成功后返回登录态token。有效期7天。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type LoginUserResponse struct {
	*tchttp.BaseResponse
	Response *LoginUserResponseParams `json:"Response"`
}

func (r *LoginUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LoginUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterUserRequestParams struct {
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 用户头像。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

type RegisterUserRequest struct {
	*tchttp.BaseRequest
	
	// 低代码互动课堂的SdkAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户在客户系统的Id，需要在同一应用下唯一。
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 用户头像。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

func (r *RegisterUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Name")
	delete(f, "OriginId")
	delete(f, "Avatar")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterUserResponseParams struct {
	// 用户Id。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 登录/注册成功后返回登录态token。有效期7天。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterUserResponse struct {
	*tchttp.BaseResponse
	Response *RegisterUserResponseParams `json:"Response"`
}

func (r *RegisterUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}