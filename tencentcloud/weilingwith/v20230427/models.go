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

package v20230427

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeApplicationListRequestParams struct {

}

type DescribeApplicationListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeApplicationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeApplicationListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationListResponseParams `json:"Response"`
}

func (r *DescribeApplicationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeApplicationTokenRequestParams struct {

}

type DescribeEdgeApplicationTokenRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEdgeApplicationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeApplicationTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeApplicationTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeApplicationTokenResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEdgeApplicationTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeApplicationTokenResponseParams `json:"Response"`
}

func (r *DescribeEdgeApplicationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeApplicationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInterfaceListRequestParams struct {

}

type DescribeInterfaceListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInterfaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInterfaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInterfaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInterfaceListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInterfaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInterfaceListResponseParams `json:"Response"`
}

func (r *DescribeInterfaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInterfaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceListRequestParams struct {

}

type DescribeWorkspaceListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWorkspaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkspaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceListResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceUserListRequestParams struct {
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil" name:"WorkspaceId"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil" name:"UpdateAt"`
}

type DescribeWorkspaceUserListRequest struct {
	*tchttp.BaseRequest
	
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil" name:"WorkspaceId"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil" name:"UpdateAt"`
}

func (r *DescribeWorkspaceUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "TenantId")
	delete(f, "UpdateAt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceUserListResponseParams struct {
	// 返回数据
	Result *SsoTeamUserResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkspaceUserListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceUserListResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SsoTeamUser struct {
	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 用户类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserType *string `json:"UserType,omitnil" name:"UserType"`

	// 所属租户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil" name:"Email"`

	// 电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 用户状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateAt *int64 `json:"CreateAt,omitnil" name:"CreateAt"`

	// 部门ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 部门名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentName *string `json:"DepartmentName,omitnil" name:"DepartmentName"`

	// 是否关联权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkFilter *int64 `json:"LinkFilter,omitnil" name:"LinkFilter"`
}

type SsoTeamUserResult struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 部门用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users []*SsoTeamUser `json:"Users,omitnil" name:"Users"`
}