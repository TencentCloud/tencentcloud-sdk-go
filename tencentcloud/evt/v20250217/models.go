// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20250217

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CompleteApprovalRequestParams struct {
	// <p>审批单号</p>
	ApprovalId *string `json:"ApprovalId,omitnil,omitempty" name:"ApprovalId"`

	// <p>审批节点</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>审批结果，1通过2拒绝</p>
	Result *uint64 `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>审批意见</p>
	Opinion *string `json:"Opinion,omitnil,omitempty" name:"Opinion"`

	// <p>审批人的身份认证Token，通过自定义角色体系回调接口分发</p><p>token信息</p>
	UserToken *string `json:"UserToken,omitnil,omitempty" name:"UserToken"`
}

type CompleteApprovalRequest struct {
	*tchttp.BaseRequest
	
	// <p>审批单号</p>
	ApprovalId *string `json:"ApprovalId,omitnil,omitempty" name:"ApprovalId"`

	// <p>审批节点</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>审批结果，1通过2拒绝</p>
	Result *uint64 `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>审批意见</p>
	Opinion *string `json:"Opinion,omitnil,omitempty" name:"Opinion"`

	// <p>审批人的身份认证Token，通过自定义角色体系回调接口分发</p><p>token信息</p>
	UserToken *string `json:"UserToken,omitnil,omitempty" name:"UserToken"`
}

func (r *CompleteApprovalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteApprovalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApprovalId")
	delete(f, "NodeId")
	delete(f, "Result")
	delete(f, "Opinion")
	delete(f, "UserToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompleteApprovalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteApprovalResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CompleteApprovalResponse struct {
	*tchttp.BaseResponse
	Response *CompleteApprovalResponseParams `json:"Response"`
}

func (r *CompleteApprovalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteApprovalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleUserRequestParams struct {
	// <p>自定义角色体系的ID</p><p>角色体系ID</p>
	RoleSystemId *int64 `json:"RoleSystemId,omitnil,omitempty" name:"RoleSystemId"`

	// <p>要添加的自定义用户ID，建议与腾讯云-子用户的用户名称保持一致</p><p>人员ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>自定义用户的名称</p><p>人员名称</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>是否启用当前用户</p>枚举值：<ul><li> 1： 启用</li><li> 2： 禁用</li></ul><p>是否启用</p>
	Enabled *uint64 `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>自定义用户的手机号</p><p>手机号</p>
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// <p>自定义用户的身份属性列表</p><p>属性列表</p>
	Attributes []*UserAttribute `json:"Attributes,omitnil,omitempty" name:"Attributes"`

	// <p>自定义用户与腾讯云-子用户关联，如果不传默认按照子用户名称进行匹配</p>
	TencentUin *uint64 `json:"TencentUin,omitnil,omitempty" name:"TencentUin"`
}

type CreateRoleUserRequest struct {
	*tchttp.BaseRequest
	
	// <p>自定义角色体系的ID</p><p>角色体系ID</p>
	RoleSystemId *int64 `json:"RoleSystemId,omitnil,omitempty" name:"RoleSystemId"`

	// <p>要添加的自定义用户ID，建议与腾讯云-子用户的用户名称保持一致</p><p>人员ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>自定义用户的名称</p><p>人员名称</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>是否启用当前用户</p>枚举值：<ul><li> 1： 启用</li><li> 2： 禁用</li></ul><p>是否启用</p>
	Enabled *uint64 `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>自定义用户的手机号</p><p>手机号</p>
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// <p>自定义用户的身份属性列表</p><p>属性列表</p>
	Attributes []*UserAttribute `json:"Attributes,omitnil,omitempty" name:"Attributes"`

	// <p>自定义用户与腾讯云-子用户关联，如果不传默认按照子用户名称进行匹配</p>
	TencentUin *uint64 `json:"TencentUin,omitnil,omitempty" name:"TencentUin"`
}

func (r *CreateRoleUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleSystemId")
	delete(f, "UserId")
	delete(f, "Username")
	delete(f, "Enabled")
	delete(f, "Phone")
	delete(f, "Attributes")
	delete(f, "TencentUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoleUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleUserResponseParams struct {
	// <p>人员ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRoleUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoleUserResponseParams `json:"Response"`
}

func (r *CreateRoleUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleUserRequestParams struct {
	// <p>自定义角色体系的ID</p>
	RoleSystemId *uint64 `json:"RoleSystemId,omitnil,omitempty" name:"RoleSystemId"`

	// <p>需要删除的自定义用户ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DeleteRoleUserRequest struct {
	*tchttp.BaseRequest
	
	// <p>自定义角色体系的ID</p>
	RoleSystemId *uint64 `json:"RoleSystemId,omitnil,omitempty" name:"RoleSystemId"`

	// <p>需要删除的自定义用户ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DeleteRoleUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleSystemId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoleUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRoleUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoleUserResponseParams `json:"Response"`
}

func (r *DeleteRoleUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutEventRequestParams struct {
	// <p>插件ID</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>需要推送的事件数据内容，格式为json，字段定义需要与事件中的定义一致</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>数据推送来源，会在生成的单据中展示数据来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>可以接受当前消息的Uin</p>
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

type PutEventRequest struct {
	*tchttp.BaseRequest
	
	// <p>插件ID</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>需要推送的事件数据内容，格式为json，字段定义需要与事件中的定义一致</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>数据推送来源，会在生成的单据中展示数据来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>可以接受当前消息的Uin</p>
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

func (r *PutEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "Data")
	delete(f, "Source")
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutEventResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PutEventResponse struct {
	*tchttp.BaseResponse
	Response *PutEventResponseParams `json:"Response"`
}

func (r *PutEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutMessageRequestParams struct {
	// <p>事件ID</p>
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// <p>需要推送的事件数据内容，格式为json，字段定义需要与事件中的定义一致</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>数据推送来源，会在生成的单据中展示数据来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

type PutMessageRequest struct {
	*tchttp.BaseRequest
	
	// <p>事件ID</p>
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// <p>需要推送的事件数据内容，格式为json，字段定义需要与事件中的定义一致</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>数据推送来源，会在生成的单据中展示数据来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

func (r *PutMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	delete(f, "Data")
	delete(f, "Source")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutMessageResponseParams struct {
	// <p>满足条件时生成的事件单id，不满足条件时为空</p>
	TicketId *string `json:"TicketId,omitnil,omitempty" name:"TicketId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PutMessageResponse struct {
	*tchttp.BaseResponse
	Response *PutMessageResponseParams `json:"Response"`
}

func (r *PutMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserAttribute struct {
	// <p>自定义角色体系中用户属性的ID</p><p>属性键名</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>自定义角色体系中的用户属性值，只支持传入对应用户属性支持的角色ID</p><p>属性值</p>
	Value []*int64 `json:"Value,omitnil,omitempty" name:"Value"`
}