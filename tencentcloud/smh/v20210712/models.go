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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateLibraryRequestParams struct {
	// 媒体库名称，最多 50 个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 存储桶全名，新建后不可更改
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 存储桶所在地域，新建后不可更改
	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`

	// 媒体库配置项，部分参数新建后不可更改
	LibraryExtension *LibraryExtension `json:"LibraryExtension,omitempty" name:"LibraryExtension"`

	// 备注，最多 250 个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type CreateLibraryRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库名称，最多 50 个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 存储桶全名，新建后不可更改
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 存储桶所在地域，新建后不可更改
	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`

	// 媒体库配置项，部分参数新建后不可更改
	LibraryExtension *LibraryExtension `json:"LibraryExtension,omitempty" name:"LibraryExtension"`

	// 备注，最多 250 个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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
	delete(f, "BucketName")
	delete(f, "BucketRegion")
	delete(f, "LibraryExtension")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLibraryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLibraryResponseParams struct {
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteLibraryRequestParams struct {
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`
}

type DeleteLibraryRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeLibrariesRequestParams struct {
	// 按照一个或者多个媒体库 ID 查询，每次请求的上限为 100 个。
	LibraryIds []*string `json:"LibraryIds,omitempty" name:"LibraryIds"`

	// 页码，整型，配合 PageSize 使用，默认值为 1。
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页数目，整型，配合 PageNumber 使用，默认值为 20，最大值为 100。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeLibrariesRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个媒体库 ID 查询，每次请求的上限为 100 个。
	LibraryIds []*string `json:"LibraryIds,omitempty" name:"LibraryIds"`

	// 页码，整型，配合 PageSize 使用，默认值为 1。
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页数目，整型，配合 PageNumber 使用，默认值为 20，最大值为 100。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
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
	List []*Library `json:"List,omitempty" name:"List"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`
}

type DescribeLibrarySecretRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`
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
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 查询到的媒体库密钥
	LibrarySecret *string `json:"LibrarySecret,omitempty" name:"LibrarySecret"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SuperAdminAccount *bool `json:"SuperAdminAccount,omitempty" name:"SuperAdminAccount"`

	// 按照一个或者多个实例 ID 查询，每次请求的上限为 100 个。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 页码，整型，配合 PageSize 使用，默认值为 1。
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页数目，整型，配合 PageNumber 使用，默认值为 20，最大值为 100。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 对指定列进行排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 续费管理筛选类型
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 超级管理管理员账号是否绑定了手机号
	BindPhone *bool `json:"BindPhone,omitempty" name:"BindPhone"`
}

type DescribeOfficialInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 是否查询实例绑定的超级管理员账号，默认值为 false。
	SuperAdminAccount *bool `json:"SuperAdminAccount,omitempty" name:"SuperAdminAccount"`

	// 按照一个或者多个实例 ID 查询，每次请求的上限为 100 个。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 页码，整型，配合 PageSize 使用，默认值为 1。
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页数目，整型，配合 PageNumber 使用，默认值为 20，最大值为 100。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 对指定列进行排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 续费管理筛选类型
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 超级管理管理员账号是否绑定了手机号
	BindPhone *bool `json:"BindPhone,omitempty" name:"BindPhone"`
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
	List []*Instance `json:"List,omitempty" name:"List"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Quantity *uint64 `json:"Quantity,omitempty" name:"Quantity"`

	// 已经使用的总存储量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。
	Storage *string `json:"Storage,omitempty" name:"Storage"`

	// 已经分配和使用的总用户数
	UserCount *uint64 `json:"UserCount,omitempty" name:"UserCount"`

	// 本月外网下行流量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。
	InternetTraffic *string `json:"InternetTraffic,omitempty" name:"InternetTraffic"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 页码，整型，配合 PageSize 使用，默认值为 1。
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页数目，整型，配合 PageNumber 使用，默认值为 20，最大值为 100。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 对指定列进行排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 来源类型筛选
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type DescribeTrafficPackagesRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个资源 ID 查询，每次请求的上限为 100 个。
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 页码，整型，配合 PageSize 使用，默认值为 1。
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页数目，整型，配合 PageNumber 使用，默认值为 20，最大值为 100。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 对指定列进行排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 来源类型筛选
	Type *uint64 `json:"Type,omitempty" name:"Type"`
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
	List []*TrafficPackage `json:"List,omitempty" name:"List"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Instance struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 专属域名。如果实例无专属域名，则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 生效时间
	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`

	// 过期时间。如果为按量计费或永久有效实例，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 用户数量。如果为按量计费或不限制用户数量实例，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserLimit *uint64 `json:"UserLimit,omitempty" name:"UserLimit"`

	// 存储容量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。如果为按量计费或不限制存储容量实例，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLimit *string `json:"StorageLimit,omitempty" name:"StorageLimit"`

	// 存储容量，单位为 GB。如果为按量计费或不限制存储容量实例，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLimitGB *uint64 `json:"StorageLimitGB,omitempty" name:"StorageLimitGB"`

	// 是否过期隔离
	Isolated *bool `json:"Isolated,omitempty" name:"Isolated"`

	// 续费标识。0：手动续费；1：自动续费；2：到期不续。
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 超级管理员账号，如果未选择查询实例绑定的超级管理员账号或当前实例未绑定超级管理员账号，则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuperAdminAccount *string `json:"SuperAdminAccount,omitempty" name:"SuperAdminAccount"`
}

type Library struct {
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 媒体库友好名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 媒体库绑定的 COS 存储桶
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 媒体库绑定的 COS 存储桶所在的地域
	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`

	// 媒体库创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 媒体库配置项
	LibraryExtension *LibraryExtension `json:"LibraryExtension,omitempty" name:"LibraryExtension"`

	// 媒体库用量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。
	Size *string `json:"Size,omitempty" name:"Size"`

	// 媒体库目录数，由于数字类型精度限制，该字段为 String 类型。
	DirNum *string `json:"DirNum,omitempty" name:"DirNum"`

	// 媒体库文件数，由于数字类型精度限制，该字段为 String 类型。
	FileNum *string `json:"FileNum,omitempty" name:"FileNum"`
}

type LibraryExtension struct {
	// true 为文件类型媒体库，可存储任何类型文件；false 为媒体类型媒体库，仅可存储照片和视频类型文件。默认为 false。在媒体库创建后不能修改。
	IsFileLibrary *bool `json:"IsFileLibrary,omitempty" name:"IsFileLibrary"`

	// true 为多租户空间媒体库，可创建多个租户空间；false 为单租户空间媒体库，不能创建租户空间，仅可使用默认的单一租户空间。默认为 false。在媒体库创建后不能修改。
	IsMultiSpace *bool `json:"IsMultiSpace,omitempty" name:"IsMultiSpace"`

	// 保存至 COS 对象存储的文件的存储类型，仅支持 STANDARD、STANDARD_IA、INTELLIGENT_TIERING、MAZ_STANDARD、MAZ_STANDARD_IA 和 MAZ_INTELLIGENT_TIERING，默认为 STANDARD，当使用多 AZ 存储桶时将自动使用 MAZ_ 开头的用于多 AZ 的存储类型，否则自动使用非 MAZ_ 开头的用于非多 AZ 的存储类型。当指定智能分层存储 INTELLIGENT_TIERING 或 MAZ_INTELLIGENT_TIERING 时，需要事先为存储桶开启智能分层存储，否则将返回失败。在媒体库创建后不能修改。
	CosStorageClass *string `json:"CosStorageClass,omitempty" name:"CosStorageClass"`

	// 是否开启回收站功能。默认为 false。
	UseRecycleBin *bool `json:"UseRecycleBin,omitempty" name:"UseRecycleBin"`

	// 当开启回收站时，自动删除回收站项目的天数，不能超过 1095（3 年），指定为 0 则不自动删除，默认为 0。当未开启回收站时，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRemoveRecycledDays *uint64 `json:"AutoRemoveRecycledDays,omitempty" name:"AutoRemoveRecycledDays"`

	// 是否启用文件路径搜索功能。默认为 false。
	EnableSearch *bool `json:"EnableSearch,omitempty" name:"EnableSearch"`

	// 设置媒体库或租户空间配额且配额小于已使用存储量时，是否拒绝设置请求。默认为 false。
	DenyOnQuotaLessThanUsage *bool `json:"DenyOnQuotaLessThanUsage,omitempty" name:"DenyOnQuotaLessThanUsage"`

	// 是否开启历史版本。默认为 false。
	EnableFileHistory *bool `json:"EnableFileHistory,omitempty" name:"EnableFileHistory"`

	// 当开启历史版本时，指定单个文件保留的历史版本的数量上限，不能超过 999，指定为 0 则不限制，默认为 0。当未开启历史版本时，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileHistoryCount *uint64 `json:"FileHistoryCount,omitempty" name:"FileHistoryCount"`

	// 当开启历史版本时，指定历史版本保留的最长天数，不能超过 999，指定为 0 则不限制，默认为 0。当未开启历史版本时，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileHistoryExpireDay *uint64 `json:"FileHistoryExpireDay,omitempty" name:"FileHistoryExpireDay"`

	// 目录或文件名的最长长度，不能超过 255，默认为 255。修改该参数不会影响存量目录或文件名，即如果将该字段的值改小，已经存在的长度超过目标值的目录或文件名不会发生变化。
	MaxDirFileNameLength *uint64 `json:"MaxDirFileNameLength,omitempty" name:"MaxDirFileNameLength"`

	// 是否开启公有读，开启后读操作无需使用访问令牌，默认为 false。仅单租户空间媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPublicRead *bool `json:"IsPublicRead,omitempty" name:"IsPublicRead"`

	// 媒体类型媒体库是否开启多相簿，开启后可创建一级目录（即相簿）且媒体文件只能保存在各相簿中，否则不能创建相簿且媒体文件只能保存在根目录。默认为 false。仅单租户空间媒体类型媒体库支持该属性，否则该属性为 null。在媒体库创建后不能修改。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMultiAlbum *bool `json:"IsMultiAlbum,omitempty" name:"IsMultiAlbum"`

	// 媒体类型媒体库是否允许上传照片，默认为 true。仅单租户空间媒体类型媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowPhoto *bool `json:"AllowPhoto,omitempty" name:"AllowPhoto"`

	// 当媒体类型媒体库允许上传照片时，指定允许的扩展名，默认为空数组，此时将根据文件扩展名对应的 MIME 类型自动判断。仅单租户空间媒体类型媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowPhotoExtName []*string `json:"AllowPhotoExtName,omitempty" name:"AllowPhotoExtName"`

	// 媒体类型媒体库是否允许上传视频，默认为 true。仅单租户空间媒体类型媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowVideo *bool `json:"AllowVideo,omitempty" name:"AllowVideo"`

	// 当媒体类型媒体库允许上传视频时，指定允许的扩展名，默认为空数组，此时将根据文件扩展名对应的 MIME 类型自动判断。仅单租户空间媒体类型媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowVideoExtName []*string `json:"AllowVideoExtName,omitempty" name:"AllowVideoExtName"`

	// 指定文件类型媒体库允许的文件扩展名，默认为空数组，此时允许上传所有类型文件。仅单租户空间文件类型媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowFileExtName []*string `json:"AllowFileExtName,omitempty" name:"AllowFileExtName"`

	// 照片上传时是否进行敏感内容鉴定，默认为 false。仅单租户空间媒体库支持该属性，否则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecognizeSensitiveContent *bool `json:"RecognizeSensitiveContent,omitempty" name:"RecognizeSensitiveContent"`
}

// Predefined struct for user
type ModifyLibraryRequestParams struct {
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 媒体库名称，最多 50 个字符。如不传则不修改。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注，最多 250 个字符。如不传则不修改。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 媒体库配置项，部分参数在新建后不可更改，且仅修改传入的参数。如不传该参数则不修改任何配置项。
	LibraryExtension *LibraryExtension `json:"LibraryExtension,omitempty" name:"LibraryExtension"`
}

type ModifyLibraryRequest struct {
	*tchttp.BaseRequest
	
	// 媒体库 ID
	LibraryId *string `json:"LibraryId,omitempty" name:"LibraryId"`

	// 媒体库名称，最多 50 个字符。如不传则不修改。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注，最多 250 个字符。如不传则不修改。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 媒体库配置项，部分参数在新建后不可更改，且仅修改传入的参数。如不传该参数则不修改任何配置项。
	LibraryExtension *LibraryExtension `json:"LibraryExtension,omitempty" name:"LibraryExtension"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type SendSmsCodeRequestParams struct {
	// 验证码目的，当前支持换绑超级管理员账号， BindSuperAdmin；体验版企业升级，ChannelUpdateVerify等
	Purpose *string `json:"Purpose,omitempty" name:"Purpose"`

	// 将作为超级管理员账号的手机号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 官方云盘实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 将作为超级管理员账号的手机号码的国家代码。默认为 +86。
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
}

type SendSmsCodeRequest struct {
	*tchttp.BaseRequest
	
	// 验证码目的，当前支持换绑超级管理员账号， BindSuperAdmin；体验版企业升级，ChannelUpdateVerify等
	Purpose *string `json:"Purpose,omitempty" name:"Purpose"`

	// 将作为超级管理员账号的手机号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 官方云盘实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 将作为超级管理员账号的手机号码的国家代码。默认为 +86。
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 专属域名。如果实例无专属域名，则该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 流量资源包来源类型，0 为付费购买，1 为赠送。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 总流量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。
	Size *string `json:"Size,omitempty" name:"Size"`

	// 总流量，单位为 GB
	SizeGB *uint64 `json:"SizeGB,omitempty" name:"SizeGB"`

	// 剩余流量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。
	Remain *string `json:"Remain,omitempty" name:"Remain"`

	// 已使用流量，单位为 Bytes，由于数字类型精度限制，该字段为 String 类型。
	Used *string `json:"Used,omitempty" name:"Used"`

	// 已使用百分比，由于数字类型精度限制，该字段为 String 类型。
	UsedPercentage *string `json:"UsedPercentage,omitempty" name:"UsedPercentage"`

	// 生效时间，即流量资源包的订购时间
	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`

	// 过期时间，即所抵扣的实例的过期时间。如果流量资源包所抵扣的实例为按量计费或永久有效实例，该属性为 null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type VerifySmsCodeRequestParams struct {
	// 验证码目的，当前支持换绑超级管理员账号，BindSuperAdmin；体验版企业升级验证ChannelUpdateVerify，等
	Purpose *string `json:"Purpose,omitempty" name:"Purpose"`

	// 将作为超级管理员账号的手机号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 短信验证码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 官方云盘实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 将作为超级管理员账号的手机号码的国家代码。默认为 +86。
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
}

type VerifySmsCodeRequest struct {
	*tchttp.BaseRequest
	
	// 验证码目的，当前支持换绑超级管理员账号，BindSuperAdmin；体验版企业升级验证ChannelUpdateVerify，等
	Purpose *string `json:"Purpose,omitempty" name:"Purpose"`

	// 将作为超级管理员账号的手机号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 短信验证码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 官方云盘实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 将作为超级管理员账号的手机号码的国家代码。默认为 +86。
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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