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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
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

	// 【已废弃】只在TrylockWorker时生效
	GameRegion *string `json:"GameRegion,omitempty" name:"GameRegion"`

	// 游戏参数
	GameParas *string `json:"GameParas,omitempty" name:"GameParas"`

	// 分辨率,，可设置为1080p或720p或1920x1080格式
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`

	// 背景图url，格式为png或jpeg，宽高1920*1080
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 【已废弃】
	SetNo *uint64 `json:"SetNo,omitempty" name:"SetNo"`

	// 单位Mbps，固定码率建议值，有一定浮动范围，后端不动态调整(MaxBitrate和MinBitrate将无效)
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 单位Mbps，动态调整最大码率建议值，会按实际情况调整
	MaxBitrate *uint64 `json:"MaxBitrate,omitempty" name:"MaxBitrate"`

	// 单位Mbps，动态调整最小码率建议值，会按实际情况调整
	MinBitrate *uint64 `json:"MinBitrate,omitempty" name:"MinBitrate"`

	// 帧率，可设置为30、45、60、90、120、144
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// 【已废弃】只在TrylockWorker时生效
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 【已废弃】优化项，便于客户灰度开启新的优化项，默认为0
	Optimization *uint64 `json:"Optimization,omitempty" name:"Optimization"`

	// 【互动云游】游戏主机用户ID
	HostUserId *string `json:"HostUserId,omitempty" name:"HostUserId"`

	// 【互动云游】角色；Player表示玩家；Viewer表示观察者
	Role *string `json:"Role,omitempty" name:"Role"`

	// 游戏相关参数
	GameContext *string `json:"GameContext,omitempty" name:"GameContext"`
}

func (r *CreateSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientSession")
	delete(f, "UserId")
	delete(f, "GameId")
	delete(f, "GameRegion")
	delete(f, "GameParas")
	delete(f, "Resolution")
	delete(f, "ImageUrl")
	delete(f, "SetNo")
	delete(f, "Bitrate")
	delete(f, "MaxBitrate")
	delete(f, "MinBitrate")
	delete(f, "Fps")
	delete(f, "UserIp")
	delete(f, "Optimization")
	delete(f, "HostUserId")
	delete(f, "Role")
	delete(f, "GameContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务端session信息，返回给JSSDK
		ServerSession *string `json:"ServerSession,omitempty" name:"ServerSession"`

		// 【已废弃】
		RoleNumber *string `json:"RoleNumber,omitempty" name:"RoleNumber"`

		// 【互动云游】角色；Player表示玩家；Viewer表示观察者
		Role *string `json:"Role,omitempty" name:"Role"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesCountRequest struct {
	*tchttp.BaseRequest

	// 游戏ID
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 实例分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeInstancesCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesCountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 客户的实例总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 客户的实例运行数
		Running *uint64 `json:"Running,omitempty" name:"Running"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaveGameArchiveRequest struct {
	*tchttp.BaseRequest

	// 游戏用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 游戏ID
	GameId *string `json:"GameId,omitempty" name:"GameId"`
}

func (r *SaveGameArchiveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveGameArchiveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "GameId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaveGameArchiveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SaveGameArchiveResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveGameArchiveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveGameArchiveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopGameRequest struct {
	*tchttp.BaseRequest

	// 游戏用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 【多人游戏】游戏主机用户ID
	HostUserId *string `json:"HostUserId,omitempty" name:"HostUserId"`
}

func (r *StopGameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopGameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "HostUserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopGameRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopGameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchGameArchiveRequest struct {
	*tchttp.BaseRequest

	// 游戏用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 游戏ID
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 游戏存档Url
	GameArchiveUrl *string `json:"GameArchiveUrl,omitempty" name:"GameArchiveUrl"`

	// 游戏相关参数
	GameContext *string `json:"GameContext,omitempty" name:"GameContext"`
}

func (r *SwitchGameArchiveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchGameArchiveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "GameId")
	delete(f, "GameArchiveUrl")
	delete(f, "GameContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchGameArchiveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SwitchGameArchiveResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchGameArchiveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchGameArchiveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrylockWorkerRequest struct {
	*tchttp.BaseRequest

	// 游戏用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 游戏ID
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 游戏区域，ap-guangzhou、ap-shanghai、ap-beijing等，如果不为空，优先按照该区域进行调度分配机器
	GameRegion *string `json:"GameRegion,omitempty" name:"GameRegion"`

	// 【废弃】资源池编号
	SetNo *uint64 `json:"SetNo,omitempty" name:"SetNo"`

	// 【必选】用户IP，用于就近调度，不填将严重影响用户体验
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *TrylockWorkerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrylockWorkerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "GameId")
	delete(f, "GameRegion")
	delete(f, "SetNo")
	delete(f, "UserIp")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TrylockWorkerRequest has unknown keys!", "")
	}
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrylockWorkerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
