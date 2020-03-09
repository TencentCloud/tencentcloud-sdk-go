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

package v20190313

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateSessionRequest struct {
	*tchttp.BaseRequest

	// 客户端session信息，从JSSDK请求中获得
	ClientSession *string `json:"ClientSession,omitempty" name:"ClientSession"`

	// 游戏ID
	GameId *string `json:"GameId,omitempty" name:"GameId"`

	// 游戏用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 游戏参数
	GameParas *string `json:"GameParas,omitempty" name:"GameParas"`

	// 游戏区域
	GameRegion *string `json:"GameRegion,omitempty" name:"GameRegion"`

	// 背景图url
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 分辨率
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`
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

type DayStreamPlayInfo struct {

	// 带宽（单位Mbps）。
	Bandwidth *float64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 流量 （单位MB）。
	Flux *float64 `json:"Flux,omitempty" name:"Flux"`

	// 在线人数。
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// 请求数。
	Request *uint64 `json:"Request,omitempty" name:"Request"`

	// 数据时间点，格式：yyyy-mm-dd HH:MM:SS。
	Time *string `json:"Time,omitempty" name:"Time"`
}

type DescribeStreamPlayInfoListRequest struct {
	*tchttp.BaseRequest

	// 结束时间，北京时间，格式：2019-04-28 10:36:00
	// 结束时间 和 开始时间  必须在同一天内。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 播放域名。
	PlayDomain *string `json:"PlayDomain,omitempty" name:"PlayDomain"`

	// 开始时间，北京时间，格式：2019-04-28 10:36:00
	// 当前时间 和 开始时间 间隔不超过30天。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 流名称，精确匹配。
	// 若不填，则为查询总体播放数据。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`
}

func (r *DescribeStreamPlayInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStreamPlayInfoListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamPlayInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计信息列表。
		DataInfoList []*DayStreamPlayInfo `json:"DataInfoList,omitempty" name:"DataInfoList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStreamPlayInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStreamPlayInfoListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWorkersRequest struct {
	*tchttp.BaseRequest
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

		// 各个区域的机器情况
		RegionDetail []*WorkerRegionInfo `json:"RegionDetail,omitempty" name:"RegionDetail" list`

		// 空闲机器总数量
		Idle *uint64 `json:"Idle,omitempty" name:"Idle"`

		// 区域个数
		RegionNum *uint64 `json:"RegionNum,omitempty" name:"RegionNum"`

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

type ForbidLiveStreamRequest struct {
	*tchttp.BaseRequest

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 您的推流域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty" name:"StreamName"`

	// 恢复流的时间。UTC 格式，例如：2018-11-29T19:00:00Z。
	// 注意：默认禁播90天，且最长支持禁播90天。
	ResumeTime *string `json:"ResumeTime,omitempty" name:"ResumeTime"`
}

func (r *ForbidLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForbidLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegisterIMRequest struct {
	*tchttp.BaseRequest

	// 用户昵称
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 用户唯一ID，建议采用用户小程序OpenID加盐形式
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户头像URL
	HeadImgUrl *string `json:"HeadImgUrl,omitempty" name:"HeadImgUrl"`

	// 用户身份，默认值：0，表示无特殊身份
	Level *int64 `json:"Level,omitempty" name:"Level"`
}

func (r *RegisterIMRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterIMRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegisterIMResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用来传递给插件的关键字段
		UserKey *string `json:"UserKey,omitempty" name:"UserKey"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterIMResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterIMResponse) FromJsonString(s string) error {
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

type WorkerRegionInfo struct {

	// 该区域空闲机器数量
	Idle *uint64 `json:"Idle,omitempty" name:"Idle"`

	// 区域
	Region *string `json:"Region,omitempty" name:"Region"`
}
