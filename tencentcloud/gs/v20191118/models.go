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

package v20191118

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateSessionRequest struct {
	*tchttp.BaseRequest

	// 客户端session信息，从JSSDK请求中获得
	ClientSession *string `json:"ClientSession,omitempty" name:"ClientSession"`

	// 游戏用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 游戏ID
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 游戏区域，ap-guangzhou、ap-shanghai、ap-beijing等
	GameRegion *string `json:"GameRegion,omitempty" name:"GameRegion"`

	// 游戏参数
	GameParas *string `json:"GameParas,omitempty" name:"GameParas"`

	// 分辨率,，可设置为1080p或720p
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// 背景图url，格式为png或jpeg，宽高1920*1080
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 资源池编号，1表示正式，2表示测试
	SetNo *uint64 `json:"SetNo,omitempty" name:"SetNo"`

	// 单位Mbps，固定码率，后端不动态调整(MaxBitrate和MinBitrate将无效)
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 单位Mbps，动态调整最大码率
	MaxBitrate *uint64 `json:"MaxBitrate,omitempty" name:"MaxBitrate"`

	// 单位Mbps，动态调整最小码率
	MinBitrate *uint64 `json:"MinBitrate,omitempty" name:"MinBitrate"`

	// 帧率，可设置为30、45或60
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// 游戏用户IP，用于就近调度，例如125.127.178.228
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 优化项，便于客户灰度开启新的优化项，默认为0
	Optimization *uint64 `json:"Optimization,omitempty" name:"Optimization"`
}

func (r *CreateSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSessionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务端session信息，返回给JSSDK
		ServerSession *string `json:"ServerSession,omitempty" name:"ServerSession"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSessionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopGameRequest struct {
	*tchttp.BaseRequest

	// 游戏用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *StopGameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopGameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopGameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopGameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopGameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TrylockWorkerRequest struct {
	*tchttp.BaseRequest

	// 游戏用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 游戏ID
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 游戏区域，ap-guangzhou、ap-shanghai、ap-beijing等
	GameRegion *string `json:"GameRegion,omitempty" name:"GameRegion"`

	// 资源池编号，1表示共用，2表示测试
	SetNo *uint64 `json:"SetNo,omitempty" name:"SetNo"`

	// 游戏用户IP，用于就近调度，例如125.127.178.228
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *TrylockWorkerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TrylockWorkerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TrylockWorkerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TrylockWorkerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TrylockWorkerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
