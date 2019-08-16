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

type DissolveRoomRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`
}

func (r *DissolveRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DissolveRoomRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DissolveRoomResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DissolveRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DissolveRoomResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type KickOutUserRequest struct {
	*tchttp.BaseRequest

	// TRTC的SDKAppId。
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号。
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 要踢的用户列表，最多10个。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds" list`
}

func (r *KickOutUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *KickOutUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type KickOutUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *KickOutUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *KickOutUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
