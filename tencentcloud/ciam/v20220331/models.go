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

package v20220331

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateUserRequest struct {
	*tchttp.BaseRequest

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// 手机号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 昵称
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 地址
	Address *string `json:"Address,omitempty" name:"Address"`

	// 用户组ID
	UserGroup []*string `json:"UserGroup,omitempty" name:"UserGroup"`

	// 生日
	Birthdate *int64 `json:"Birthdate,omitempty" name:"Birthdate"`

	// 自定义属性
	CustomizationAttributes []*MemberMap `json:"CustomizationAttributes,omitempty" name:"CustomizationAttributes"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "PhoneNumber")
	delete(f, "Email")
	delete(f, "Password")
	delete(f, "UserName")
	delete(f, "Nickname")
	delete(f, "Address")
	delete(f, "UserGroup")
	delete(f, "Birthdate")
	delete(f, "CustomizationAttributes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的用户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		User *User `json:"User,omitempty" name:"User"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsersRequest struct {
	*tchttp.BaseRequest

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// 用户ID数组
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *DeleteUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserByIdRequest struct {
	*tchttp.BaseRequest

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeUserByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserByIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		User *User `json:"User,omitempty" name:"User"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// key值
	Key *string `json:"Key,omitempty" name:"Key"`

	// value值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 逻辑值
	Logic *bool `json:"Logic,omitempty" name:"Logic"`
}

type LinkAccountRequest struct {
	*tchttp.BaseRequest

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// 主用户ID
	PrimaryUserId *string `json:"PrimaryUserId,omitempty" name:"PrimaryUserId"`

	// 从用户ID
	SecondaryUserId *string `json:"SecondaryUserId,omitempty" name:"SecondaryUserId"`

	// 融合属性(PHONENUMBER,EMAIL)
	UserLinkedOnAttribute *string `json:"UserLinkedOnAttribute,omitempty" name:"UserLinkedOnAttribute"`
}

func (r *LinkAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LinkAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "PrimaryUserId")
	delete(f, "SecondaryUserId")
	delete(f, "UserLinkedOnAttribute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LinkAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type LinkAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LinkAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LinkAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUserByPropertyRequest struct {
	*tchttp.BaseRequest

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// 查询的属性（支持phoneNumber，email）
	PropertyCode *string `json:"PropertyCode,omitempty" name:"PropertyCode"`

	// 属性值
	PropertyValue *string `json:"PropertyValue,omitempty" name:"PropertyValue"`
}

func (r *ListUserByPropertyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserByPropertyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "PropertyCode")
	delete(f, "PropertyValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserByPropertyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListUserByPropertyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Users []*User `json:"Users,omitempty" name:"Users"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUserByPropertyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserByPropertyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUserRequest struct {
	*tchttp.BaseRequest

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// 分页数据
	Pageable *Pageable `json:"Pageable,omitempty" name:"Pageable"`

	// Key可选值为condition、userGroupId
	// 
	// <li> **condition** </li>	Values = 查询条件，用户ID，用户名称，手机或邮箱
	// <li> **userGroupId** </li>	Values = 用户组ID
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "Pageable")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 分页对象
	// 注意：此字段可能返回 null，表示取不到有效值。
		Pageable *Pageable `json:"Pageable,omitempty" name:"Pageable"`

		// 用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Content []*User `json:"Content,omitempty" name:"Content"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MemberMap struct {

	// 健
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type Pageable struct {

	// 每页数量
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 当前页码
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserStoreId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 重置后的用户密码
		Password *string `json:"Password,omitempty" name:"Password"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPasswordRequest struct {
	*tchttp.BaseRequest

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *SetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "UserId")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUserRequest struct {
	*tchttp.BaseRequest

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// 用户名称
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 手机号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 昵称
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 地址
	Address *string `json:"Address,omitempty" name:"Address"`

	// 用户组
	UserGroup []*string `json:"UserGroup,omitempty" name:"UserGroup"`

	// 生日
	Birthdate *int64 `json:"Birthdate,omitempty" name:"Birthdate"`

	// 自定义属性
	CustomizationAttributes []*MemberMap `json:"CustomizationAttributes,omitempty" name:"CustomizationAttributes"`
}

func (r *UpdateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserStoreId")
	delete(f, "UserName")
	delete(f, "PhoneNumber")
	delete(f, "Email")
	delete(f, "Nickname")
	delete(f, "Address")
	delete(f, "UserGroup")
	delete(f, "Birthdate")
	delete(f, "CustomizationAttributes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 更新之后的用户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		User *User `json:"User,omitempty" name:"User"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUserStatusRequest struct {
	*tchttp.BaseRequest

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// NORMAL（正常）,LOCK（锁定）,FREEZE（冻结）,请传英文大写字母
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateUserStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "UserId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUserStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateUserStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitempty" name:"Email"`

	// 上次登录时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastSignOn *int64 `json:"LastSignOn,omitempty" name:"LastSignOn"`

	// 创建时间
	CreatedDate *int64 `json:"CreatedDate,omitempty" name:"CreatedDate"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 用户来源
	UserDataSourceEnum *string `json:"UserDataSourceEnum,omitempty" name:"UserDataSourceEnum"`

	// 昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 生日
	// 注意：此字段可能返回 null，表示取不到有效值。
	Birthdate *int64 `json:"Birthdate,omitempty" name:"Birthdate"`

	// 用户组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroups []*string `json:"UserGroups,omitempty" name:"UserGroups"`

	// 上次修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifiedDate *int64 `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`

	// 自定义属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomAttributes []*MemberMap `json:"CustomAttributes,omitempty" name:"CustomAttributes"`

	// 身份证号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResidentIdentityCard *string `json:"ResidentIdentityCard,omitempty" name:"ResidentIdentityCard"`

	// QQ的OpenId
	// 注意：此字段可能返回 null，表示取不到有效值。
	QqOpenId *string `json:"QqOpenId,omitempty" name:"QqOpenId"`

	// QQ的UnionId
	// 注意：此字段可能返回 null，表示取不到有效值。
	QqUnionId *string `json:"QqUnionId,omitempty" name:"QqUnionId"`

	// 微信的WechatOpenId
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatOpenId *string `json:"WechatOpenId,omitempty" name:"WechatOpenId"`

	// 微信的WechatUnionId
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatUnionId *string `json:"WechatUnionId,omitempty" name:"WechatUnionId"`

	// 支付宝的AlipayUserId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlipayUserId *string `json:"AlipayUserId,omitempty" name:"AlipayUserId"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Locale *string `json:"Locale,omitempty" name:"Locale"`

	// 性别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gender *string `json:"Gender,omitempty" name:"Gender"`

	// 实名核验方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityVerificationMethod *string `json:"IdentityVerificationMethod,omitempty" name:"IdentityVerificationMethod"`

	// 是否已经实名核验
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityVerified *bool `json:"IdentityVerified,omitempty" name:"IdentityVerified"`

	// 工作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *string `json:"Job,omitempty" name:"Job"`

	// 国家
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nationality *string `json:"Nationality,omitempty" name:"Nationality"`

	// 是否主账号（进行过账号融合后，主账号为true，从账号为false）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Primary *bool `json:"Primary,omitempty" name:"Primary"`

	// 时区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 是否已经首次登录
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlreadyFirstLogin *bool `json:"AlreadyFirstLogin,omitempty" name:"AlreadyFirstLogin"`

	// 租户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`

	// 用户目录id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// 锁定类型（分为管理员锁定，和登录策略锁定）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockType *string `json:"LockType,omitempty" name:"LockType"`

	// 锁定时间点
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockTime *int64 `json:"LockTime,omitempty" name:"LockTime"`
}
