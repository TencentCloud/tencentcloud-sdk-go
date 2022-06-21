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

package v20201117

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ApplyStackRequestParams struct {
	// 资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 待执行apply事件的版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

type ApplyStackRequest struct {
	*tchttp.BaseRequest
	
	// 资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 待执行apply事件的版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *ApplyStackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyStackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StackId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyStackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyStackResponseParams struct {
	// 执行的事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyStackResponse struct {
	*tchttp.BaseResponse
	Response *ApplyStackResponseParams `json:"Response"`
}

func (r *ApplyStackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyStackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStackRequestParams struct {
	// 资源栈名称，不得超过60个字符
	StackName *string `json:"StackName,omitempty" name:"StackName"`

	// 资源栈所在地域
	StackRegion *string `json:"StackRegion,omitempty" name:"StackRegion"`

	// HCL模板URL，⽬前仅限 COS URL, ⽂件为zip压缩格式
	TemplateUrl *string `json:"TemplateUrl,omitempty" name:"TemplateUrl"`

	// 资源栈描述，不得超过200个字符
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateStackRequest struct {
	*tchttp.BaseRequest
	
	// 资源栈名称，不得超过60个字符
	StackName *string `json:"StackName,omitempty" name:"StackName"`

	// 资源栈所在地域
	StackRegion *string `json:"StackRegion,omitempty" name:"StackRegion"`

	// HCL模板URL，⽬前仅限 COS URL, ⽂件为zip压缩格式
	TemplateUrl *string `json:"TemplateUrl,omitempty" name:"TemplateUrl"`

	// 资源栈描述，不得超过200个字符
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateStackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StackName")
	delete(f, "StackRegion")
	delete(f, "TemplateUrl")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStackResponseParams struct {
	// 创建得到的资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 资源栈版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateStackResponse struct {
	*tchttp.BaseResponse
	Response *CreateStackResponseParams `json:"Response"`
}

func (r *CreateStackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStackVersionRequestParams struct {
	// 待增加版本的资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 模板 URL，⽬前仅限 COS URL, ⽂件为zip压缩格式
	TemplateUrl *string `json:"TemplateUrl,omitempty" name:"TemplateUrl"`

	// 版本名称，不得超过60个字符
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 版本描述，不得超过200个字符
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateStackVersionRequest struct {
	*tchttp.BaseRequest
	
	// 待增加版本的资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 模板 URL，⽬前仅限 COS URL, ⽂件为zip压缩格式
	TemplateUrl *string `json:"TemplateUrl,omitempty" name:"TemplateUrl"`

	// 版本名称，不得超过60个字符
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 版本描述，不得超过200个字符
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateStackVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStackVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StackId")
	delete(f, "TemplateUrl")
	delete(f, "VersionName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStackVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStackVersionResponseParams struct {
	// 新创建的版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateStackVersionResponse struct {
	*tchttp.BaseResponse
	Response *CreateStackVersionResponseParams `json:"Response"`
}

func (r *CreateStackVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStackVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStackRequestParams struct {
	// 待删除的资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`
}

type DeleteStackRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`
}

func (r *DeleteStackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStackResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteStackResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStackResponseParams `json:"Response"`
}

func (r *DeleteStackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStackVersionRequestParams struct {
	// 待删除的版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

type DeleteStackVersionRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *DeleteStackVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStackVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStackVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStackVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteStackVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStackVersionResponseParams `json:"Response"`
}

func (r *DeleteStackVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStackVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStackEventRequestParams struct {
	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type DescribeStackEventRequest struct {
	*tchttp.BaseRequest
	
	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeStackEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStackEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStackEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStackEventResponseParams struct {
	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 事件类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 事件状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 状态信息
	EventMessage *string `json:"EventMessage,omitempty" name:"EventMessage"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 控制台输出文本
	ConsoleLog *string `json:"ConsoleLog,omitempty" name:"ConsoleLog"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStackEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStackEventResponseParams `json:"Response"`
}

func (r *DescribeStackEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStackEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStackEventsRequestParams struct {
	// 按照⼀个或者多个事件ID查询
	EventIds []*string `json:"EventIds,omitempty" name:"EventIds"`

	// <li>**VersionId**</li>
	// 按照【**版本ID**】过滤，VersionId形如 `ver-kg8hn58h`
	// 类型：string
	// 
	// <li>**StackId**</li>
	// 按照【**资源栈ID**】过滤，StackId形如 `stk-hz5vn3te`
	// 类型：string
	// 
	// <li>**Type**</li>
	// 按照【**事件类型**】过滤，Type 形如 plan, apply, destroy
	// 类型：string
	// 
	// <li>**Status**</li>
	// 按照【**事件状态**】过滤，Status形如 queueing, running, success, failed
	// 类型：string
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeStackEventsRequest struct {
	*tchttp.BaseRequest
	
	// 按照⼀个或者多个事件ID查询
	EventIds []*string `json:"EventIds,omitempty" name:"EventIds"`

	// <li>**VersionId**</li>
	// 按照【**版本ID**】过滤，VersionId形如 `ver-kg8hn58h`
	// 类型：string
	// 
	// <li>**StackId**</li>
	// 按照【**资源栈ID**】过滤，StackId形如 `stk-hz5vn3te`
	// 类型：string
	// 
	// <li>**Type**</li>
	// 按照【**事件类型**】过滤，Type 形如 plan, apply, destroy
	// 类型：string
	// 
	// <li>**Status**</li>
	// 按照【**事件状态**】过滤，Status形如 queueing, running, success, failed
	// 类型：string
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeStackEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStackEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStackEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStackEventsResponseParams struct {
	// 符合条件的事件数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 事件详细信息列表
	Events []*EventInfo `json:"Events,omitempty" name:"Events"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStackEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStackEventsResponseParams `json:"Response"`
}

func (r *DescribeStackEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStackEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStackVersionsRequestParams struct {
	// 按照⼀个或者多个版本ID查询
	VersionIds []*string `json:"VersionIds,omitempty" name:"VersionIds"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// <li>**Name**</li>
	// 按照【**版本名称**】进行过滤
	// 类型：string
	// 
	// <li>**Status**</li>
	// 按照【**版本状态**】过滤，形如`VERSION_EDITING`，`PLAN_IN_PROGRESS`等
	// 类型：string
	// 
	// <li>**StackId**</li>
	// 按照版本所属的【**资源栈ID**】进行过滤，形如`stk-xxxxxx`
	// 类型：string
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeStackVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 按照⼀个或者多个版本ID查询
	VersionIds []*string `json:"VersionIds,omitempty" name:"VersionIds"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// <li>**Name**</li>
	// 按照【**版本名称**】进行过滤
	// 类型：string
	// 
	// <li>**Status**</li>
	// 按照【**版本状态**】过滤，形如`VERSION_EDITING`，`PLAN_IN_PROGRESS`等
	// 类型：string
	// 
	// <li>**StackId**</li>
	// 按照版本所属的【**资源栈ID**】进行过滤，形如`stk-xxxxxx`
	// 类型：string
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeStackVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStackVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VersionIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStackVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStackVersionsResponseParams struct {
	// 符合条件的版本数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 版本详细信息列表
	Versions []*VersionInfo `json:"Versions,omitempty" name:"Versions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStackVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStackVersionsResponseParams `json:"Response"`
}

func (r *DescribeStackVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStackVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStacksRequestParams struct {
	// 按照⼀个或者多个资源栈ID查询
	StackIds []*string `json:"StackIds,omitempty" name:"StackIds"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeStacksRequest struct {
	*tchttp.BaseRequest
	
	// 按照⼀个或者多个资源栈ID查询
	StackIds []*string `json:"StackIds,omitempty" name:"StackIds"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeStacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StackIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStacksResponseParams struct {
	// 符合条件的资源栈数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 资源栈详细信息列表
	Stacks []*StackInfo `json:"Stacks,omitempty" name:"Stacks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStacksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStacksResponseParams `json:"Response"`
}

func (r *DescribeStacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyStackRequestParams struct {
	// 资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 待执行destroy事件的版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

type DestroyStackRequest struct {
	*tchttp.BaseRequest
	
	// 资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 待执行destroy事件的版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *DestroyStackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyStackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StackId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyStackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyStackResponseParams struct {
	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroyStackResponse struct {
	*tchttp.BaseResponse
	Response *DestroyStackResponseParams `json:"Response"`
}

func (r *DestroyStackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyStackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventInfo struct {
	// 事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 事件类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 版本状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 状态信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Filter struct {
	// 条件名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 匹配的值，可以有多个
	Values []*string `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type PlanStackRequestParams struct {
	// 资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 待执行plan事件的版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

type PlanStackRequest struct {
	*tchttp.BaseRequest
	
	// 资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 待执行plan事件的版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *PlanStackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PlanStackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StackId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PlanStackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PlanStackResponseParams struct {
	// 执行的事件ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PlanStackResponse struct {
	*tchttp.BaseResponse
	Response *PlanStackResponseParams `json:"Response"`
}

func (r *PlanStackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PlanStackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StackInfo struct {
	// 资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 资源栈名称
	StackName *string `json:"StackName,omitempty" name:"StackName"`

	// 资源栈描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 所处地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 资源栈状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type UpdateStackRequestParams struct {
	// 待更新的资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 资源栈名称，不得超过60个字符
	StackName *string `json:"StackName,omitempty" name:"StackName"`

	// 资源栈描述，不得超过200个字符
	Description *string `json:"Description,omitempty" name:"Description"`
}

type UpdateStackRequest struct {
	*tchttp.BaseRequest
	
	// 待更新的资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 资源栈名称，不得超过60个字符
	StackName *string `json:"StackName,omitempty" name:"StackName"`

	// 资源栈描述，不得超过200个字符
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateStackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateStackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StackId")
	delete(f, "StackName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateStackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateStackResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateStackResponse struct {
	*tchttp.BaseResponse
	Response *UpdateStackResponseParams `json:"Response"`
}

func (r *UpdateStackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateStackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateStackVersionRequestParams struct {
	// 待更新的版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 模板 URL，⽬前仅限 COS URL, ⽂件为zip压缩格式
	TemplateUrl *string `json:"TemplateUrl,omitempty" name:"TemplateUrl"`

	// 版本名称，不得超过60个字符
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 版本描述，不得超过200个字符
	Description *string `json:"Description,omitempty" name:"Description"`
}

type UpdateStackVersionRequest struct {
	*tchttp.BaseRequest
	
	// 待更新的版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 模板 URL，⽬前仅限 COS URL, ⽂件为zip压缩格式
	TemplateUrl *string `json:"TemplateUrl,omitempty" name:"TemplateUrl"`

	// 版本名称，不得超过60个字符
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 版本描述，不得超过200个字符
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateStackVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateStackVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VersionId")
	delete(f, "TemplateUrl")
	delete(f, "VersionName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateStackVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateStackVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateStackVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateStackVersionResponseParams `json:"Response"`
}

func (r *UpdateStackVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateStackVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VersionInfo struct {
	// 版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 版本描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 资源栈ID
	StackId *string `json:"StackId,omitempty" name:"StackId"`

	// 版本状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}