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

package v20180801

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddDelayLiveStreamRequest struct {
	*tchttp.BaseRequest
	// 应用名称。
	AppName *string `json:"AppName" name:"AppName"`
	// 您的加速域名。
	DomainName *string `json:"DomainName" name:"DomainName"`
	// 流名称。
	StreamName *string `json:"StreamName" name:"StreamName"`
	// 延播时间，单位：秒，上限：600秒。
	DelayTime *uint64 `json:"DelayTime" name:"DelayTime"`
}

func (r *AddDelayLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddDelayLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddDelayLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddDelayLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddDelayLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineInfoRequest struct {
	*tchttp.BaseRequest
	// 取得第几页。
	// 默认值：1
	PageNum *uint64 `json:"PageNum" name:"PageNum"`
	// 分页大小。
	// 
	// 最大值：100。
	// 取值范围：1~100 之前的任意整数。
	// 默认值：10
	PageSize *uint64 `json:"PageSize" name:"PageSize"`
	// 0:未开始推流 1:正在推流 2:服务出错 3:已关闭。
	Status *int64 `json:"Status" name:"Status"`
	// 流名称。
	StreamName *string `json:"StreamName" name:"StreamName"`
}

func (r *DescribeLiveStreamOnlineInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamOnlineInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 分页的页码。
		PageNum *uint64 `json:"PageNum" name:"PageNum"`
		// 每页大小
		PageSize *uint64 `json:"PageSize" name:"PageSize"`
		// 符合条件的总个数。
		TotalNum *uint64 `json:"TotalNum" name:"TotalNum"`
		// 总页数。
		TotalPage *uint64 `json:"TotalPage" name:"TotalPage"`
		// 流信息列表
		StreamInfoList []*StreamInfo `json:"StreamInfoList" name:"StreamInfoList" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamOnlineInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamOnlineInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineListRequest struct {
	*tchttp.BaseRequest
	// 您的加速域名。
	DomainName *string `json:"DomainName" name:"DomainName"`
	// 应用名称。
	AppName *string `json:"AppName" name:"AppName"`
	// 取得第几页，默认1。
	PageNum *uint64 `json:"PageNum" name:"PageNum"`
	// 每页大小，最大100。 
	// 取值：1~100之前的任意整数。
	// 默认值：10
	PageSize *uint64 `json:"PageSize" name:"PageSize"`
}

func (r *DescribeLiveStreamOnlineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamOnlineListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamOnlineListResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的总个数。
		TotalNum *uint64 `json:"TotalNum" name:"TotalNum"`
		// 总页数。
		TotalPage *uint64 `json:"TotalPage" name:"TotalPage"`
		// 分页的页码。
		PageNum *uint64 `json:"PageNum" name:"PageNum"`
		// 每页显示的条数。
		PageSize *uint64 `json:"PageSize" name:"PageSize"`
		// 正在推送流的信息列表
		OnlineInfo []*StreamOnlineInfo `json:"OnlineInfo" name:"OnlineInfo" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamOnlineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamOnlineListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPublishedListRequest struct {
	*tchttp.BaseRequest
	// 您的域名。
	DomainName *string `json:"DomainName" name:"DomainName"`
	// 结束时间。
	// UTC 格式，例如：2016-06-30T19:00:00Z。
	// EndTime 和 StartTime 之间的间隔不能超过 30 天。
	EndTime *string `json:"EndTime" name:"EndTime"`
	// 起始时间。 
	// UTC 格式，例如：2016-06-29T19:00:00Z。
	StartTime *string `json:"StartTime" name:"StartTime"`
	// 直播流所属应用名称。
	AppName *string `json:"AppName" name:"AppName"`
	// 取得第几页。
	// 默认值：1
	PageNum *uint64 `json:"PageNum" name:"PageNum"`
	// 分页大小。
	// 
	// 最大值：100。
	// 取值范围：1~100 之前的任意整数。
	// 默认值：10
	PageSize *uint64 `json:"PageSize" name:"PageSize"`
}

func (r *DescribeLiveStreamPublishedListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamPublishedListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamPublishedListResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 推流记录信息。
		PublishInfo []*StreamName `json:"PublishInfo" name:"PublishInfo" list`
		// 分页的页码。
		PageNum *uint64 `json:"PageNum" name:"PageNum"`
		// 每页大小
		PageSize *uint64 `json:"PageSize" name:"PageSize"`
		// 符合条件的总个数。
		TotalNum *uint64 `json:"TotalNum" name:"TotalNum"`
		// 总页数。
		TotalPage *uint64 `json:"TotalPage" name:"TotalPage"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamPublishedListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamPublishedListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamStateRequest struct {
	*tchttp.BaseRequest
	// 应用名称。
	AppName *string `json:"AppName" name:"AppName"`
	// 您的加速域名。
	DomainName *string `json:"DomainName" name:"DomainName"`
	// 流名称。
	StreamName *string `json:"StreamName" name:"StreamName"`
}

func (r *DescribeLiveStreamStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamStateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLiveStreamStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 流状态
		StreamState *string `json:"StreamState" name:"StreamState"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLiveStreamStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLiveStreamStateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DropLiveStreamRequest struct {
	*tchttp.BaseRequest
	// 流名称。
	StreamName *string `json:"StreamName" name:"StreamName"`
	// 您的加速域名。
	DomainName *string `json:"DomainName" name:"DomainName"`
	// 应用名称。
	AppName *string `json:"AppName" name:"AppName"`
}

func (r *DropLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DropLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DropLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DropLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DropLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ForbidLiveStreamRequest struct {
	*tchttp.BaseRequest
	// 应用名称。
	AppName *string `json:"AppName" name:"AppName"`
	// 您的加速域名。
	DomainName *string `json:"DomainName" name:"DomainName"`
	// 流名称。
	StreamName *string `json:"StreamName" name:"StreamName"`
	// 恢复流的时间。UTC 格式，例如：2018-11-29T19:00:00Z。
	// 
	// UTC 时间，格式：2018-08-08T17:37:00Z。
	ResumeTime *string `json:"ResumeTime" name:"ResumeTime"`
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
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForbidLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ForbidLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PublishTime struct {
	// 推流时间
	// UTC 格式，例如：2018-06-29T19:00:00Z。
	PublishTime *string `json:"PublishTime" name:"PublishTime"`
}

type ResumeDelayLiveStreamRequest struct {
	*tchttp.BaseRequest
	// 应用名称。
	AppName *string `json:"AppName" name:"AppName"`
	// 您的加速域名。
	DomainName *string `json:"DomainName" name:"DomainName"`
	// 流名称。
	StreamName *string `json:"StreamName" name:"StreamName"`
}

func (r *ResumeDelayLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeDelayLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeDelayLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeDelayLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeDelayLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeLiveStreamRequest struct {
	*tchttp.BaseRequest
	// 应用名称。
	AppName *string `json:"AppName" name:"AppName"`
	// 您的加速域名。
	DomainName *string `json:"DomainName" name:"DomainName"`
	// 流名称。
	StreamName *string `json:"StreamName" name:"StreamName"`
}

func (r *ResumeLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeLiveStreamRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResumeLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResumeLiveStreamResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StreamInfo struct {
	// 直播流所属应用名称
	AppName *string `json:"AppName" name:"AppName"`
	// 创建模式
	CreateMode *string `json:"CreateMode" name:"CreateMode"`
	// 创建时间，如: 2018-07-13 14:48:23
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 流状态
	Status *int64 `json:"Status" name:"Status"`
	// 流id
	StreamId *string `json:"StreamId" name:"StreamId"`
	// 流名称
	StreamName *string `json:"StreamName" name:"StreamName"`
	// 水印id
	WaterMarkId *string `json:"WaterMarkId" name:"WaterMarkId"`
}

type StreamName struct {
	// 流名称。
	StreamName *string `json:"StreamName" name:"StreamName"`
}

type StreamOnlineInfo struct {
	// 流名称。
	StreamName *string `json:"StreamName" name:"StreamName"`
	// 推流时间列表
	PublishTimeList []*PublishTime `json:"PublishTimeList" name:"PublishTimeList" list`
}
