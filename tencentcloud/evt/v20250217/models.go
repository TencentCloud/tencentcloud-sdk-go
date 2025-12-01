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
type CreateRoleUserRequestParams struct {
	// 角色体系ID
	RoleSystemId *int64 `json:"RoleSystemId,omitnil,omitempty" name:"RoleSystemId"`

	// 人员ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 人员名称
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 是否启用
	Enabled *uint64 `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 属性列表
	Attributes []*UserAttribute `json:"Attributes,omitnil,omitempty" name:"Attributes"`
}

type CreateRoleUserRequest struct {
	*tchttp.BaseRequest
	
	// 角色体系ID
	RoleSystemId *int64 `json:"RoleSystemId,omitnil,omitempty" name:"RoleSystemId"`

	// 人员ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 人员名称
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 是否启用
	Enabled *uint64 `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 属性列表
	Attributes []*UserAttribute `json:"Attributes,omitnil,omitempty" name:"Attributes"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoleUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleUserResponseParams struct {
	// 人员ID
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

type UserAttribute struct {
	// 属性键名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 属性值
	Value []*int64 `json:"Value,omitnil,omitempty" name:"Value"`
}