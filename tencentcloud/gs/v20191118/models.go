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

	// 游戏区域
	GameRegion *string `json:"GameRegion,omitempty" name:"GameRegion"`

	// 游戏参数
	GameParas *string `json:"GameParas,omitempty" name:"GameParas"`

	// 分辨率
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// 背景图url
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 资源池编号
	SetNo *uint64 `json:"SetNo,omitempty" name:"SetNo"`
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

type DescribeWorkersRequest struct {
	*tchttp.BaseRequest

	// 资源池编号，值为2的幂，1表示共用，2表示测试
	SetNo *uint64 `json:"SetNo,omitempty" name:"SetNo"`
}

func (r *DescribeWorkersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWorkersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWorkersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 空闲机器总数量
		Idle *uint64 `json:"Idle,omitempty" name:"Idle"`

		// 区域个数
		RegionNum *uint64 `json:"RegionNum,omitempty" name:"RegionNum"`

		// 各个区域的机器情况
		RegionDetail []*WorkerRegionInfo `json:"RegionDetail,omitempty" name:"RegionDetail" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWorkersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWorkersResponse) FromJsonString(s string) error {
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

type WorkerRegionInfo struct {

	// 区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 该区域空闲机器数量
	Idle *uint64 `json:"Idle,omitempty" name:"Idle"`
}
