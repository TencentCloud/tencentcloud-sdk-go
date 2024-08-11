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

package v20210712

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateLibraryRequestParams struct {
	// 媒体库名称，最多 50 个字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注，最多 250 个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 存储桶全名，新建后不可更改。当前版本不再支持指定存储桶。
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 存储桶所在地域，新建后不可更改。当前版本不再支持指定存储桶所在地域。
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 媒体库配置项，部分参数新建后不可更改
	LibraryExtension *LibraryExtension `json:"LibraryExtension,omitnil,omitempty" name:"LibraryExtension"`
}

type CreateLibraryRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库名称，最多 50 个字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注，最多 250 个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 存储桶全名，新建后不可更改。当前版本不再支持指定存储桶。
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 存储桶所在地域，新建后不可更改。当前版本不再支持指定存储桶所在地域。
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 媒体库配置项，部分参数新建后不可更改
	LibraryExtension *LibraryExtension `json:"LibraryExtension,omitnil,omitempty" name:"LibraryExtension"`
}

func (r *CreateLibraryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLibraryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "BucketName")
	delete(f, "BucketRegion")
	delete(f, "LibraryExtension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLibraryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLibraryResponseParams struct {
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLibraryResponse struct {
	*tchttp.BaseResponse
	Response *CreateLibraryResponseParams `json:"Response"`
}

func (r *CreateLibraryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserLifecycleRequestParams struct {
	// 媒体库 ID。
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 用于唯一查找用户的过滤器。
	Filter *UserFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 隔离时间，当时间超过该时间点后，指定用户将无法登录，但他的账号信息、文件资源会被保留，可以通过再次调用本接口更新隔离时间，恢复登录。如不指定，则代表不设置隔离时间，且当前用户已经设置的隔离时间会被删除。
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// 销毁时间，当时间超过该时间点后，指定用户的资源将被销毁且无法通过再次调用此接口更新时间。如果同时指定了 IsolateTime 则不能早于 IsolateTime 指定的时间。如不指定，则代表不设置销毁时间，且当前用户已经设置的销毁时间会被删除。
	DestroyTime *string `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`
}

type CreateUserLifecycleRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库 ID。
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 用于唯一查找用户的过滤器。
	Filter *UserFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 隔离时间，当时间超过该时间点后，指定用户将无法登录，但他的账号信息、文件资源会被保留，可以通过再次调用本接口更新隔离时间，恢复登录。如不指定，则代表不设置隔离时间，且当前用户已经设置的隔离时间会被删除。
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// 销毁时间，当时间超过该时间点后，指定用户的资源将被销毁且无法通过再次调用此接口更新时间。如果同时指定了 IsolateTime 则不能早于 IsolateTime 指定的时间。如不指定，则代表不设置销毁时间，且当前用户已经设置的销毁时间会被删除。
	DestroyTime *string `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`
}

func (r *CreateUserLifecycleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserLifecycleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LibraryId")
	delete(f, "Filter")
	delete(f, "IsolateTime")
	delete(f, "DestroyTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserLifecycleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserLifecycleResponseParams struct {
	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 设置的隔离时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// 设置的销毁时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestroyTime *string `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserLifecycleResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserLifecycleResponseParams `json:"Response"`
}

func (r *CreateUserLifecycleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserLifecycleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// 媒体库 ID。
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 用户角色，当只支持 user。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 是否启用。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 手机号国家码，不传默认为 null，此时无法使用该登录方式进行登录。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号码，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有手机号重复则报错。CountryCode 和 PhoneNumber 必须同时传入或同时不传入。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有邮箱重复则报错。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 账号，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有账号重复则报错。只能使用大小写字母、数字、中划线、下划线、小数点，长度不超过 50 个字符。
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 密码的 base64 形式，不传默认为 null，此时无法使用该登录方式进行登录。AccountName 和 AccountPassword 必须同时传入或同时不传入。
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// 第三方账号 ID，用于关联第三方账号体系，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有第三方账号重复则报错。只能使用大小写字母、数字、中划线、下划线、小数点，长度不超过 200 个字符。
	AccountUserId *string `json:"AccountUserId,omitnil,omitempty" name:"AccountUserId"`

	// 备注。不超过 255 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 昵称。不超过 100 个字符。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 用户头像地址。不超过 255 个字符。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 自定义信息。不超过 255 个字符。
	Customize *string `json:"Customize,omitnil,omitempty" name:"Customize"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库 ID。
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 用户角色，当只支持 user。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 是否启用。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 手机号国家码，不传默认为 null，此时无法使用该登录方式进行登录。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号码，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有手机号重复则报错。CountryCode 和 PhoneNumber 必须同时传入或同时不传入。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有邮箱重复则报错。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 账号，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有账号重复则报错。只能使用大小写字母、数字、中划线、下划线、小数点，长度不超过 50 个字符。
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 密码的 base64 形式，不传默认为 null，此时无法使用该登录方式进行登录。AccountName 和 AccountPassword 必须同时传入或同时不传入。
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// 第三方账号 ID，用于关联第三方账号体系，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有第三方账号重复则报错。只能使用大小写字母、数字、中划线、下划线、小数点，长度不超过 200 个字符。
	AccountUserId *string `json:"AccountUserId,omitnil,omitempty" name:"AccountUserId"`

	// 备注。不超过 255 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 昵称。不超过 100 个字符。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 用户头像地址。不超过 255 个字符。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 自定义信息。不超过 255 个字符。
	Customize *string `json:"Customize,omitnil,omitempty" name:"Customize"`
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
	delete(f, "LibraryId")
	delete(f, "Role")
	delete(f, "Enabled")
	delete(f, "CountryCode")
	delete(f, "PhoneNumber")
	delete(f, "Email")
	delete(f, "AccountName")
	delete(f, "AccountPassword")
	delete(f, "AccountUserId")
	delete(f, "Comment")
	delete(f, "Nickname")
	delete(f, "Avatar")
	delete(f, "Customize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// 用户所在的媒体库 ID。
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户创建时间。
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 用户角色.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 是否启用。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 手机号国家码，如未指定则为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号码，如未指定则为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱，如未指定则为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 账号，如未指定则为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 第三方账号 ID，用于关联第三方账号体系，如未指定则为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountUserId *string `json:"AccountUserId,omitnil,omitempty" name:"AccountUserId"`

	// 备注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 昵称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 用户头像地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 自定义信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Customize *string `json:"Customize,omitnil,omitempty" name:"Customize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteLibraryRequestParams struct {
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`
}

type DeleteLibraryRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`
}

func (r *DeleteLibraryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLibraryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LibraryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLibraryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLibraryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLibraryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLibraryResponseParams `json:"Response"`
}

func (r *DeleteLibraryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRequestParams struct {
	// 媒体库 ID。
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 用于唯一查找用户的过滤器数组，数组之间为 **或** 的关系，即满足任意一个过滤器的用户，都将被删除，单次传入的过滤器最多为 100 个。
	Filters []*UserFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库 ID。
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 用于唯一查找用户的过滤器数组，数组之间为 **或** 的关系，即满足任意一个过滤器的用户，都将被删除，单次传入的过滤器最多为 100 个。
	Filters []*UserFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LibraryId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserResponseParams `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibrariesRequestParams struct {
	// 按照一个或者多个媒体库 ID 查询，每次请求的上限为 100 个。
	LibraryIds []*string `json:"LibraryIds,omitnil,omitempty" name:"LibraryIds"`

	// 页码，整型，配合 PageSize 使用，默认值为 1。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目，整型，配合 PageNumber 使用，默认值为 20，最大值为 100。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeLibrariesRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个媒体库 ID 查询，每次请求的上限为 100 个。
	LibraryIds []*string `json:"LibraryIds,omitnil,omitempty" name:"LibraryIds"`

	// 页码，整型，配合 PageSize 使用，默认值为 1。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目，整型，配合 PageNumber 使用，默认值为 20，最大值为 100。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeLibrariesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibrariesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LibraryIds")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibrariesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibrariesResponseParams struct {
	// 媒体库列表
	List []*Library `json:"List,omitnil,omitempty" name:"List"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibrariesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibrariesResponseParams `json:"Response"`
}

func (r *DescribeLibrariesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibrariesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibrarySecretRequestParams struct {
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`
}

type DescribeLibrarySecretRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`
}

func (r *DescribeLibrarySecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibrarySecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LibraryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLibrarySecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLibrarySecretResponseParams struct {
	// 查询的媒体库 ID
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 查询到的媒体库密钥
	LibrarySecret *string `json:"LibrarySecret,omitnil,omitempty" name:"LibrarySecret"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLibrarySecretResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLibrarySecretResponseParams `json:"Response"`
}

func (r *DescribeLibrarySecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLibrarySecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfficialInstancesRequestParams struct {
	// 是否查询实例绑定的超级管理员账号，默认值为 false。
	SuperAdminAccount *bool `json:"SuperAdminAccount,omitnil,omitempty" name:"SuperAdminAccount"`

	// 按照一个或者多个实例 ID 查询，每次请求的上限为 100 个。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 页码，整型，配合 PageSize 使用，默认值为 1。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目，整型，配合 PageNumber 使用，默认值为 20，最大值为 100。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 对指定列进行排序
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 续费管理筛选类型
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 超级管理管理员账号是否绑定了手机号
	BindPhone *bool `json:"BindPhone,omitnil,omitempty" name:"BindPhone"`
}

type DescribeOfficialInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 是否查询实例绑定的超级管理员账号，默认值为 false。
	SuperAdminAccount *bool `json:"SuperAdminAccount,omitnil,omitempty" name:"SuperAdminAccount"`

	// 按照一个或者多个实例 ID 查询，每次请求的上限为 100 个。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 页码，整型，配合 PageSize 使用，默认值为 1。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目，整型，配合 PageNumber 使用，默认值为 20，最大值为 100。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 对指定列进行排序
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 续费管理筛选类型
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 超级管理管理员账号是否绑定了手机号
	BindPhone *bool `json:"BindPhone,omitnil,omitempty" name:"BindPhone"`
}

func (r *DescribeOfficialInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfficialInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SuperAdminAccount")
	delete(f, "InstanceIds")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "AutoRenew")
	delete(f, "BindPhone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOfficialInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfficialInstancesResponseParams struct {
	// 实例列表
	List []*Instance `json:"List,omitnil,omitempty" name:"List"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOfficialInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOfficialInstancesResponseParams `json:"Response"`
}

func (r *DescribeOfficialInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfficialInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfficialOverviewRequestParams struct {

}

type DescribeOfficialOverviewRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeOfficialOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfficialOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOfficialOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfficialOverviewResponseParams struct {
	// 云盘实例数量
	Quantity *uint64 `json:"Quantity,omitnil,omitempty" name:"Quantity"`

	// 已经使用的总存储量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。
	Storage *string `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 已经分配和使用的总用户数
	UserCount *uint64 `json:"UserCount,omitnil,omitempty" name:"UserCount"`

	// 本月外网下行流量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。
	InternetTraffic *string `json:"InternetTraffic,omitnil,omitempty" name:"InternetTraffic"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOfficialOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOfficialOverviewResponseParams `json:"Response"`
}

func (r *DescribeOfficialOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfficialOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficPackagesRequestParams struct {
	// 按照一个或者多个资源 ID 查询，每次请求的上限为 100 个。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 页码，整型，配合 PageSize 使用，默认值为 1。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目，整型，配合 PageNumber 使用，默认值为 20，最大值为 100。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 对指定列进行排序
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 来源类型筛选
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeTrafficPackagesRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个资源 ID 查询，每次请求的上限为 100 个。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 页码，整型，配合 PageSize 使用，默认值为 1。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目，整型，配合 PageNumber 使用，默认值为 20，最大值为 100。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 对指定列进行排序
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 来源类型筛选
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeTrafficPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceIds")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrafficPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficPackagesResponseParams struct {
	// 流量包列表
	List []*TrafficPackage `json:"List,omitnil,omitempty" name:"List"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrafficPackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrafficPackagesResponseParams `json:"Response"`
}

func (r *DescribeTrafficPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserLifecycleRequestParams struct {
	// 媒体库 ID。
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 用于唯一查找用户的过滤器。
	Filter *UserFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeUserLifecycleRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库 ID。
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 用于唯一查找用户的过滤器。
	Filter *UserFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeUserLifecycleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserLifecycleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LibraryId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserLifecycleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserLifecycleResponseParams struct {
	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 设置的隔离时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// 设置的销毁时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestroyTime *string `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserLifecycleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserLifecycleResponseParams `json:"Response"`
}

func (r *DescribeUserLifecycleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserLifecycleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 专属域名。如果实例无专属域名，则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 生效时间
	EffectiveTime *string `json:"EffectiveTime,omitnil,omitempty" name:"EffectiveTime"`

	// 过期时间。如果为按量计费或永久有效实例，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 用户数量。如果为按量计费或不限制用户数量实例，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserLimit *uint64 `json:"UserLimit,omitnil,omitempty" name:"UserLimit"`

	// 存储容量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。如果为按量计费或不限制存储容量实例，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLimit *string `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 存储容量，单位为 GB。如果为按量计费或不限制存储容量实例，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLimitGB *uint64 `json:"StorageLimitGB,omitnil,omitempty" name:"StorageLimitGB"`

	// 是否过期隔离
	Isolated *bool `json:"Isolated,omitnil,omitempty" name:"Isolated"`

	// 续费标识。0：手动续费；1：自动续费；2：到期不续。
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 超级管理员账号，如果未选择查询实例绑定的超级管理员账号或当前实例未绑定超级管理员账号，则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuperAdminAccount *string `json:"SuperAdminAccount,omitnil,omitempty" name:"SuperAdminAccount"`
}

type Library struct {
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 媒体库友好名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 媒体库绑定的 COS 存储桶
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 媒体库绑定的 COS 存储桶所在的地域
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 媒体库创建时间
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 媒体库配置项
	LibraryExtension *LibraryExtension `json:"LibraryExtension,omitnil,omitempty" name:"LibraryExtension"`

	// 媒体库用量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 媒体库目录数，由于数字类型精度限制，该字段为 String 类型。
	DirNum *string `json:"DirNum,omitnil,omitempty" name:"DirNum"`

	// 媒体库文件数，由于数字类型精度限制，该字段为 String 类型。
	FileNum *string `json:"FileNum,omitnil,omitempty" name:"FileNum"`
}

type LibraryExtension struct {
	// true 为文件类型媒体库，可存储任何类型文件；false 为媒体类型媒体库，仅可存储照片和视频类型文件。默认为 false。在媒体库创建后不能修改。
	IsFileLibrary *bool `json:"IsFileLibrary,omitnil,omitempty" name:"IsFileLibrary"`

	// true 为多租户空间媒体库，可创建多个租户空间；false 为单租户空间媒体库，不能创建租户空间，仅可使用默认的单一租户空间。默认为 false。在媒体库创建后不能修改。
	IsMultiSpace *bool `json:"IsMultiSpace,omitnil,omitempty" name:"IsMultiSpace"`

	// 保存至 COS 对象存储的文件的存储类型，仅支持 STANDARD、STANDARD_IA、INTELLIGENT_TIERING、MAZ_STANDARD、MAZ_STANDARD_IA 和 MAZ_INTELLIGENT_TIERING，默认为 STANDARD，当使用多 AZ 存储桶时将自动使用 MAZ_ 开头的用于多 AZ 的存储类型，否则自动使用非 MAZ_ 开头的用于非多 AZ 的存储类型。当指定智能分层存储 INTELLIGENT_TIERING 或 MAZ_INTELLIGENT_TIERING 时，需要事先为存储桶开启智能分层存储，否则将返回失败。在媒体库创建后不能修改。
	CosStorageClass *string `json:"CosStorageClass,omitnil,omitempty" name:"CosStorageClass"`

	// 是否开启回收站功能。默认为 false。
	UseRecycleBin *bool `json:"UseRecycleBin,omitnil,omitempty" name:"UseRecycleBin"`

	// 当开启回收站时，自动删除回收站项目的天数，不能超过 1095（3 年），指定为 0 则不自动删除，默认为 0。当未开启回收站时，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRemoveRecycledDays *uint64 `json:"AutoRemoveRecycledDays,omitnil,omitempty" name:"AutoRemoveRecycledDays"`

	// 是否启用文件路径搜索功能。默认为 false。
	EnableSearch *bool `json:"EnableSearch,omitnil,omitempty" name:"EnableSearch"`

	// 设置媒体库或租户空间配额且配额小于已使用存储量时，是否拒绝设置请求。默认为 false。
	DenyOnQuotaLessThanUsage *bool `json:"DenyOnQuotaLessThanUsage,omitnil,omitempty" name:"DenyOnQuotaLessThanUsage"`

	// 是否开启历史版本。默认为 false。
	EnableFileHistory *bool `json:"EnableFileHistory,omitnil,omitempty" name:"EnableFileHistory"`

	// 当开启历史版本时，指定单个文件保留的历史版本的数量上限，不能超过 999，指定为 0 则不限制，默认为 0。当未开启历史版本时，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileHistoryCount *uint64 `json:"FileHistoryCount,omitnil,omitempty" name:"FileHistoryCount"`

	// 当开启历史版本时，指定历史版本保留的最长天数，不能超过 999，指定为 0 则不限制，默认为 0。当未开启历史版本时，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileHistoryExpireDay *uint64 `json:"FileHistoryExpireDay,omitnil,omitempty" name:"FileHistoryExpireDay"`

	// 目录或文件名的最长长度，不能超过 255，默认为 255。修改该参数不会影响存量目录或文件名，即如果将该字段的值改小，已经存在的长度超过目标值的目录或文件名不会发生变化。
	MaxDirFileNameLength *uint64 `json:"MaxDirFileNameLength,omitnil,omitempty" name:"MaxDirFileNameLength"`

	// 是否开启公有读，开启后读操作无需使用访问令牌，默认为 false。仅单租户空间媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPublicRead *bool `json:"IsPublicRead,omitnil,omitempty" name:"IsPublicRead"`

	// 媒体类型媒体库是否开启多相簿，开启后可创建一级目录（即相簿）且媒体文件只能保存在各相簿中，否则不能创建相簿且媒体文件只能保存在根目录。默认为 false。仅单租户空间媒体类型媒体库支持该属性，否则该属性为 null。在媒体库创建后不能修改。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMultiAlbum *bool `json:"IsMultiAlbum,omitnil,omitempty" name:"IsMultiAlbum"`

	// 媒体类型媒体库是否允许上传照片，默认为 true。仅单租户空间媒体类型媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowPhoto *bool `json:"AllowPhoto,omitnil,omitempty" name:"AllowPhoto"`

	// 当媒体类型媒体库允许上传照片时，指定允许的扩展名，默认为空数组，此时将根据文件扩展名对应的 MIME 类型自动判断。仅单租户空间媒体类型媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowPhotoExtName []*string `json:"AllowPhotoExtName,omitnil,omitempty" name:"AllowPhotoExtName"`

	// 媒体类型媒体库是否允许上传视频，默认为 true。仅单租户空间媒体类型媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowVideo *bool `json:"AllowVideo,omitnil,omitempty" name:"AllowVideo"`

	// 当媒体类型媒体库允许上传视频时，指定允许的扩展名，默认为空数组，此时将根据文件扩展名对应的 MIME 类型自动判断。仅单租户空间媒体类型媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowVideoExtName []*string `json:"AllowVideoExtName,omitnil,omitempty" name:"AllowVideoExtName"`

	// 指定文件类型媒体库允许的文件扩展名，默认为空数组，此时允许上传所有类型文件。仅单租户空间文件类型媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowFileExtName []*string `json:"AllowFileExtName,omitnil,omitempty" name:"AllowFileExtName"`

	// 照片上传时是否进行敏感内容鉴定，默认为 false。仅单租户空间媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecognizeSensitiveContent *bool `json:"RecognizeSensitiveContent,omitnil,omitempty" name:"RecognizeSensitiveContent"`
}

// Predefined struct for user
type ModifyLibraryRequestParams struct {
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 媒体库名称，最多 50 个字符。如不传则不修改。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注，最多 250 个字符。如不传则不修改。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 媒体库配置项，部分参数在新建后不可更改，且仅修改传入的参数。如不传该参数则不修改任何配置项。
	LibraryExtension *LibraryExtension `json:"LibraryExtension,omitnil,omitempty" name:"LibraryExtension"`
}

type ModifyLibraryRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 媒体库名称，最多 50 个字符。如不传则不修改。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注，最多 250 个字符。如不传则不修改。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 媒体库配置项，部分参数在新建后不可更改，且仅修改传入的参数。如不传该参数则不修改任何配置项。
	LibraryExtension *LibraryExtension `json:"LibraryExtension,omitnil,omitempty" name:"LibraryExtension"`
}

func (r *ModifyLibraryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LibraryId")
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "LibraryExtension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLibraryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLibraryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLibraryResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLibraryResponseParams `json:"Response"`
}

func (r *ModifyLibraryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLibraryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRequestParams struct {
	// 媒体库 ID。
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 用于唯一查找用户的过滤器。
	Filter *UserFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 用户角色，当只支持 user。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 是否启用。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 手机号国家码，不传默认为 null，此时无法使用该登录方式进行登录。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号码，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有手机号重复则报错。CountryCode 和 PhoneNumber 必须同时传入或同时不传入。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有邮箱重复则报错。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 账号，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有账号重复则报错。只能使用大小写字母、数字、中划线、下划线、小数点，长度不超过 50 个字符。
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 密码的 base64 形式，不传默认为 null，此时无法使用该登录方式进行登录。AccountName 和 AccountPassword 必须同时传入或同时不传入。
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// 第三方账号 ID，用于关联第三方账号体系，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有第三方账号重复则报错。只能使用大小写字母、数字、中划线、下划线、小数点，长度不超过 200 个字符。
	AccountUserId *string `json:"AccountUserId,omitnil,omitempty" name:"AccountUserId"`

	// 备注。不超过 255 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 昵称。不超过 100 个字符。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 用户头像地址。不超过 255 个字符。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 自定义信息。不超过 255 个字符。
	Customize *string `json:"Customize,omitnil,omitempty" name:"Customize"`
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库 ID。
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 用于唯一查找用户的过滤器。
	Filter *UserFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 用户角色，当只支持 user。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 是否启用。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 手机号国家码，不传默认为 null，此时无法使用该登录方式进行登录。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号码，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有手机号重复则报错。CountryCode 和 PhoneNumber 必须同时传入或同时不传入。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有邮箱重复则报错。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 账号，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有账号重复则报错。只能使用大小写字母、数字、中划线、下划线、小数点，长度不超过 50 个字符。
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 密码的 base64 形式，不传默认为 null，此时无法使用该登录方式进行登录。AccountName 和 AccountPassword 必须同时传入或同时不传入。
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// 第三方账号 ID，用于关联第三方账号体系，不传默认为 null，此时无法使用该登录方式进行登录。如果与同一媒体库内已有第三方账号重复则报错。只能使用大小写字母、数字、中划线、下划线、小数点，长度不超过 200 个字符。
	AccountUserId *string `json:"AccountUserId,omitnil,omitempty" name:"AccountUserId"`

	// 备注。不超过 255 个字符。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 昵称。不超过 100 个字符。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 用户头像地址。不超过 255 个字符。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 自定义信息。不超过 255 个字符。
	Customize *string `json:"Customize,omitnil,omitempty" name:"Customize"`
}

func (r *ModifyUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LibraryId")
	delete(f, "Filter")
	delete(f, "Role")
	delete(f, "Enabled")
	delete(f, "CountryCode")
	delete(f, "PhoneNumber")
	delete(f, "Email")
	delete(f, "AccountName")
	delete(f, "AccountPassword")
	delete(f, "AccountUserId")
	delete(f, "Comment")
	delete(f, "Nickname")
	delete(f, "Avatar")
	delete(f, "Customize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserResponseParams struct {
	// 用户所在的媒体库 ID。
	LibraryId *string `json:"LibraryId,omitnil,omitempty" name:"LibraryId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户创建时间。
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 用户角色.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 是否启用。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 手机号国家码，如未指定则为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号码，如未指定则为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱，如未指定则为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 账号，如未指定则为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 第三方账号 ID，用于关联第三方账号体系，如未指定则为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountUserId *string `json:"AccountUserId,omitnil,omitempty" name:"AccountUserId"`

	// 备注。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 昵称。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 用户头像地址。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 自定义信息。
	Customize *string `json:"Customize,omitnil,omitempty" name:"Customize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserResponseParams `json:"Response"`
}

func (r *ModifyUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendSmsCodeRequestParams struct {
	// 验证码目的，当前支持换绑超级管理员账号， BindSuperAdmin；体验版企业升级，ChannelUpdateVerify等
	Purpose *string `json:"Purpose,omitnil,omitempty" name:"Purpose"`

	// 将作为超级管理员账号的手机号码
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 官方云盘实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 将作为超级管理员账号的手机号码的国家代码。默认为 +86。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`
}

type SendSmsCodeRequest struct {
	*tchttp.BaseRequest
	
	// 验证码目的，当前支持换绑超级管理员账号， BindSuperAdmin；体验版企业升级，ChannelUpdateVerify等
	Purpose *string `json:"Purpose,omitnil,omitempty" name:"Purpose"`

	// 将作为超级管理员账号的手机号码
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 官方云盘实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 将作为超级管理员账号的手机号码的国家代码。默认为 +86。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`
}

func (r *SendSmsCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendSmsCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Purpose")
	delete(f, "PhoneNumber")
	delete(f, "InstanceId")
	delete(f, "CountryCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendSmsCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendSmsCodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendSmsCodeResponse struct {
	*tchttp.BaseResponse
	Response *SendSmsCodeResponseParams `json:"Response"`
}

func (r *SendSmsCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendSmsCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrafficPackage struct {
	// 流量资源包所抵扣的实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 专属域名。如果实例无专属域名，则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 流量资源包来源类型，0 为付费购买，1 为赠送。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 总流量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 总流量，单位为 GB
	SizeGB *uint64 `json:"SizeGB,omitnil,omitempty" name:"SizeGB"`

	// 剩余流量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。
	Remain *string `json:"Remain,omitnil,omitempty" name:"Remain"`

	// 已使用流量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。
	Used *string `json:"Used,omitnil,omitempty" name:"Used"`

	// 已使用百分比，由于数字类型精度限制，该字段为 String 类型。
	UsedPercentage *string `json:"UsedPercentage,omitnil,omitempty" name:"UsedPercentage"`

	// 生效时间，即流量资源包的订购时间
	EffectiveTime *string `json:"EffectiveTime,omitnil,omitempty" name:"EffectiveTime"`

	// 过期时间，即所抵扣的实例的过期时间。如果流量资源包所抵扣的实例为按量计费或永久有效实例，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type UserFilter struct {
	// 过滤类型，当前支持：UserId、PhoneNumber、Email、AccountName、AccountUserId。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 过滤值，只支持完全匹配，不支持模糊搜索。针对不同的 Key，Value 的取值如下：
	// UserId: user12345678abcde
	// PhoneNumber: +86-13800000000（格式为：{CountryCode}-{PhoneNumber}）
	// Email: admin@mail.foobar.com
	// AccountName: account_name
	// AccountUserId: x53mYVqykfPqTCqekbNwwa4aXk4
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type VerifySmsCodeRequestParams struct {
	// 验证码目的，当前支持换绑超级管理员账号，BindSuperAdmin；体验版企业升级验证ChannelUpdateVerify，等
	Purpose *string `json:"Purpose,omitnil,omitempty" name:"Purpose"`

	// 将作为超级管理员账号的手机号码
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 短信验证码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 官方云盘实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 将作为超级管理员账号的手机号码的国家代码。默认为 +86。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`
}

type VerifySmsCodeRequest struct {
	*tchttp.BaseRequest
	
	// 验证码目的，当前支持换绑超级管理员账号，BindSuperAdmin；体验版企业升级验证ChannelUpdateVerify，等
	Purpose *string `json:"Purpose,omitnil,omitempty" name:"Purpose"`

	// 将作为超级管理员账号的手机号码
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 短信验证码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 官方云盘实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 将作为超级管理员账号的手机号码的国家代码。默认为 +86。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`
}

func (r *VerifySmsCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifySmsCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Purpose")
	delete(f, "PhoneNumber")
	delete(f, "Code")
	delete(f, "InstanceId")
	delete(f, "CountryCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifySmsCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifySmsCodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifySmsCodeResponse struct {
	*tchttp.BaseResponse
	Response *VerifySmsCodeResponseParams `json:"Response"`
}

func (r *VerifySmsCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifySmsCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}