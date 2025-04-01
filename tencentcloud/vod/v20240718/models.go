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

package v20240718

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateIncrementalMigrationStrategyRequestParams struct {
	// <b>点播[专业版](/document/product/266/115396)[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 策略生效的存储桶 ID。
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// 增量迁移策略名称，名称长度不超过100个字符，允许的字符为：`中文、英文、0-9、_、-`。
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 源站类型。取值有：
	// <li>HTTP：HTTP 源。</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 增量迁移 HTTP 回源源站配置，当 OriginType 取值 `HTTP` 时必填。
	HttpOriginConfig *IncrementalMigrationHttpOriginConfig `json:"HttpOriginConfig,omitnil,omitempty" name:"HttpOriginConfig"`
}

type CreateIncrementalMigrationStrategyRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[专业版](/document/product/266/115396)[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 策略生效的存储桶 ID。
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// 增量迁移策略名称，名称长度不超过100个字符，允许的字符为：`中文、英文、0-9、_、-`。
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 源站类型。取值有：
	// <li>HTTP：HTTP 源。</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 增量迁移 HTTP 回源源站配置，当 OriginType 取值 `HTTP` 时必填。
	HttpOriginConfig *IncrementalMigrationHttpOriginConfig `json:"HttpOriginConfig,omitnil,omitempty" name:"HttpOriginConfig"`
}

func (r *CreateIncrementalMigrationStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIncrementalMigrationStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "BucketId")
	delete(f, "StrategyName")
	delete(f, "OriginType")
	delete(f, "HttpOriginConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIncrementalMigrationStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIncrementalMigrationStrategyResponseParams struct {
	// 增量迁移策略 ID。
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIncrementalMigrationStrategyResponse struct {
	*tchttp.BaseResponse
	Response *CreateIncrementalMigrationStrategyResponseParams `json:"Response"`
}

func (r *CreateIncrementalMigrationStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIncrementalMigrationStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageCredentialsRequestParams struct {
	// <b>点播专业版[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 按照下方语法组装好策略后，先序列化为字符串，再做 URL Encode，结果作为 Policy 字段入参。服务端会对该字段做 URL Decode，并按解析后的策略授予临时访问凭证权限，请按规范传入参数。
	// 注意： 
	// 1.策略语法参照[访问管理策略](/document/product/598/10603)。
	// 2.策略中不能包含 principal 元素。
	// 3.策略的 action 元素仅支持：<li>name/vod:PutObject;</li><li>name/vod:ListParts;</li><li>name/vod:PostObject;</li><li>name/vod:InitiateMultipartUpload;</li><li>name/vod:UploadPart;</li><li>name/vod:CompleteMultipartUpload;</li><li>name/vod:AbortMultipartUpload;</li><li>name/vod:ListMultipartUploads;</li>4.策略的 resource 元素填写格式为：`qcs::vod:[存储地域]:uid/[账号AppID]:prefix//[点播应用ID]/[存储桶ID]/[存储路径]`，其中存储地域、账号 AppID、点播应用 ID、存储桶 ID 和存储路径要按需填写，其他内容不允许改动，例：`qcs:ap-chongqing:vod::uid/1231456789:prefix//1234567890/2ceds3ew323w3mu/file_path`。
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 指定临时证书的有效期，单位：秒。
	// 默认 1800 秒，最大 129600 秒。
	DurationSeconds *uint64 `json:"DurationSeconds,omitnil,omitempty" name:"DurationSeconds"`
}

type CreateStorageCredentialsRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播专业版[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 按照下方语法组装好策略后，先序列化为字符串，再做 URL Encode，结果作为 Policy 字段入参。服务端会对该字段做 URL Decode，并按解析后的策略授予临时访问凭证权限，请按规范传入参数。
	// 注意： 
	// 1.策略语法参照[访问管理策略](/document/product/598/10603)。
	// 2.策略中不能包含 principal 元素。
	// 3.策略的 action 元素仅支持：<li>name/vod:PutObject;</li><li>name/vod:ListParts;</li><li>name/vod:PostObject;</li><li>name/vod:InitiateMultipartUpload;</li><li>name/vod:UploadPart;</li><li>name/vod:CompleteMultipartUpload;</li><li>name/vod:AbortMultipartUpload;</li><li>name/vod:ListMultipartUploads;</li>4.策略的 resource 元素填写格式为：`qcs::vod:[存储地域]:uid/[账号AppID]:prefix//[点播应用ID]/[存储桶ID]/[存储路径]`，其中存储地域、账号 AppID、点播应用 ID、存储桶 ID 和存储路径要按需填写，其他内容不允许改动，例：`qcs:ap-chongqing:vod::uid/1231456789:prefix//1234567890/2ceds3ew323w3mu/file_path`。
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 指定临时证书的有效期，单位：秒。
	// 默认 1800 秒，最大 129600 秒。
	DurationSeconds *uint64 `json:"DurationSeconds,omitnil,omitempty" name:"DurationSeconds"`
}

func (r *CreateStorageCredentialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageCredentialsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Policy")
	delete(f, "DurationSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStorageCredentialsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageCredentialsResponseParams struct {
	// 临时访问凭证。
	Credentials *Credentials `json:"Credentials,omitnil,omitempty" name:"Credentials"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStorageCredentialsResponse struct {
	*tchttp.BaseResponse
	Response *CreateStorageCredentialsResponseParams `json:"Response"`
}

func (r *CreateStorageCredentialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageRequestParams struct {
	// <b>点播专业版[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 存储地域，必须是系统支持地域。
	// 通过 [DescribeStorageRegions](https://cloud.tencent.com/document/product/266/72480) 接口可以查询到所有存储地域及已经开通存储桶的地域。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 存储名称。
	// <li>仅支持小写英文字母、数字、中划线 “-” 及其组合；</li>
	// <li>存储命名不能以 “-” 开头或结尾；</li>
	// <li>存储命名最大长度为 64 字符。</li>
	StorageName *string `json:"StorageName,omitnil,omitempty" name:"StorageName"`
}

type CreateStorageRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播专业版[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 存储地域，必须是系统支持地域。
	// 通过 [DescribeStorageRegions](https://cloud.tencent.com/document/product/266/72480) 接口可以查询到所有存储地域及已经开通存储桶的地域。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 存储名称。
	// <li>仅支持小写英文字母、数字、中划线 “-” 及其组合；</li>
	// <li>存储命名不能以 “-” 开头或结尾；</li>
	// <li>存储命名最大长度为 64 字符。</li>
	StorageName *string `json:"StorageName,omitnil,omitempty" name:"StorageName"`
}

func (r *CreateStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "StorageRegion")
	delete(f, "StorageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStorageResponseParams struct {
	// 存储桶 ID。
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStorageResponse struct {
	*tchttp.BaseResponse
	Response *CreateStorageResponseParams `json:"Response"`
}

func (r *CreateStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {
	// 访问凭证 ID。
	AccessKeyId *string `json:"AccessKeyId,omitnil,omitempty" name:"AccessKeyId"`

	// 访问凭证 Key。
	SecretAccessKey *string `json:"SecretAccessKey,omitnil,omitempty" name:"SecretAccessKey"`

	// 访问凭证 Token，长度和绑定的策略有关，最长不超过 4096 字节。
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// 访问凭证的过期时间。
	Expiration *string `json:"Expiration,omitnil,omitempty" name:"Expiration"`
}

// Predefined struct for user
type DeleteIncrementalMigrationStrategyRequestParams struct {
	// <b>点播[专业版](/document/product/266/115396)[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 策略生效的存储桶 ID。
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// 增量迁移策略 ID。
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

type DeleteIncrementalMigrationStrategyRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[专业版](/document/product/266/115396)[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 策略生效的存储桶 ID。
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// 增量迁移策略 ID。
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

func (r *DeleteIncrementalMigrationStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIncrementalMigrationStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "BucketId")
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIncrementalMigrationStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIncrementalMigrationStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIncrementalMigrationStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIncrementalMigrationStrategyResponseParams `json:"Response"`
}

func (r *DeleteIncrementalMigrationStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIncrementalMigrationStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIncrementalMigrationStrategyInfosRequestParams struct {
	// <b>点播[专业版](/document/product/266/115396)[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 过滤条件，Filters.Values 的上限为 `20`；若 Filters 长度为 `0` 则查询时无过滤条件限制。 详细的过滤条件如下： <li>BucketId<br>   按照【<strong>存储桶 ID</strong>】进行过滤<br>   类型：String<br>   必选：否<br></li><li>StrategyId<br>   按照【<strong>策略 ID</strong>】进行过滤。<br>   类型：String<br>   必选：否</li> 
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回结果的排序。 SortBy.Field 取值有：<li>UpdateTime：创建时间。</li>若不填，SortBy.Field 默认值为 `UpdateTime`，SortBy.Order 默认值为 `Desc`。
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 分页返回的起始偏移量，默认值为 `0`。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值为 `20`，最大值为 `100`。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeIncrementalMigrationStrategyInfosRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[专业版](/document/product/266/115396)[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 过滤条件，Filters.Values 的上限为 `20`；若 Filters 长度为 `0` 则查询时无过滤条件限制。 详细的过滤条件如下： <li>BucketId<br>   按照【<strong>存储桶 ID</strong>】进行过滤<br>   类型：String<br>   必选：否<br></li><li>StrategyId<br>   按照【<strong>策略 ID</strong>】进行过滤。<br>   类型：String<br>   必选：否</li> 
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回结果的排序。 SortBy.Field 取值有：<li>UpdateTime：创建时间。</li>若不填，SortBy.Field 默认值为 `UpdateTime`，SortBy.Order 默认值为 `Desc`。
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 分页返回的起始偏移量，默认值为 `0`。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值为 `20`，最大值为 `100`。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeIncrementalMigrationStrategyInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIncrementalMigrationStrategyInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIncrementalMigrationStrategyInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIncrementalMigrationStrategyInfosResponseParams struct {
	// 总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 策略信息集合。
	StrategyInfoSet []*IncrementalMigrationStrategyInfo `json:"StrategyInfoSet,omitnil,omitempty" name:"StrategyInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIncrementalMigrationStrategyInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIncrementalMigrationStrategyInfosResponseParams `json:"Response"`
}

func (r *DescribeIncrementalMigrationStrategyInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIncrementalMigrationStrategyInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageRequestParams struct {
	// <b>点播专业版[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 过滤条件，Filters.Values 的上限为 20；若 Filters 长度为 0 则分页查询子应用 SubAppId 下的存储信息。 详细的过滤条件如下：
	// <li>BucketId<br>   按照【<strong>存储桶 ID</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	// <li>StorageName<br>   按照【<strong>存储名称</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回结果的排序。 SortBy.Field 取值有：
	// <li>CreateTime：创建时间。</li>若不填，SortBy.Field 默认值为 CreateTime，SortBy.Order 默认值为 Asc。
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 分页返回的起始偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值为 20，最大值为 1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeStorageRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播专业版[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 过滤条件，Filters.Values 的上限为 20；若 Filters 长度为 0 则分页查询子应用 SubAppId 下的存储信息。 详细的过滤条件如下：
	// <li>BucketId<br>   按照【<strong>存储桶 ID</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	// <li>StorageName<br>   按照【<strong>存储名称</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回结果的排序。 SortBy.Field 取值有：
	// <li>CreateTime：创建时间。</li>若不填，SortBy.Field 默认值为 CreateTime，SortBy.Order 默认值为 Asc。
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 分页返回的起始偏移量，默认值为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页返回的记录条数，默认值为 20，最大值为 1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageResponseParams struct {
	// 符合条件的存储数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合条件的存储信息列表。
	StorageInfoSet []*StorageInfo `json:"StorageInfoSet,omitnil,omitempty" name:"StorageInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStorageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStorageResponseParams `json:"Response"`
}

func (r *DescribeStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type IncrementalMigrationHttpEndpointInfo struct {
	// 地址信息，支持域名或 IP 地址。
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 备份地址信息。
	StandbyEndpointSet []*string `json:"StandbyEndpointSet,omitnil,omitempty" name:"StandbyEndpointSet"`
}

type IncrementalMigrationHttpHeader struct {
	// Header 键。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Header 值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type IncrementalMigrationHttpHeaderInfo struct {
	// Http Header 透传模式。取值有：
	// <li>FOLLOW_ALL：透传所有头部信息；</li>
	// <li>FOLLOW_PART：透传部分头部信息；</li>
	// <li>IGNORE_PART：忽略部分头部信息。</li>参数必填。
	HeaderFollowMode *string `json:"HeaderFollowMode,omitnil,omitempty" name:"HeaderFollowMode"`

	// 需透传 Header Key 集合，仅当 HeaderFollowMode 取值 `FOLLOW_PART` 时需要填充。
	FollowHttpHeaderKeySet []*string `json:"FollowHttpHeaderKeySet,omitnil,omitempty" name:"FollowHttpHeaderKeySet"`

	// 新增 Header 键值对集合。
	NewHttpHeaderSet []*IncrementalMigrationHttpHeader `json:"NewHttpHeaderSet,omitnil,omitempty" name:"NewHttpHeaderSet"`
}

type IncrementalMigrationHttpOriginCondition struct {
	// 触发回源条件的 HTTP Code。若不填充，默认取值 `404`。
	HttpStatusCode *uint64 `json:"HttpStatusCode,omitnil,omitempty" name:"HttpStatusCode"`

	// 触发回源条件的对象键前缀。
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

type IncrementalMigrationHttpOriginConfig struct {
	// 回源源站信息。
	OriginInfo *IncrementalMigrationHttpOriginInfo `json:"OriginInfo,omitnil,omitempty" name:"OriginInfo"`

	// 回源参数。
	OriginParameter *IncrementalMigrationHttpOriginParameter `json:"OriginParameter,omitnil,omitempty" name:"OriginParameter"`

	// 回源模式。取值有：
	// <li>SYNC：同步回源；</li>
	// <li>ASYNC：异步回源。</li>若不填，默认取 `SYNC` 同步回源。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 回源条件。
	OriginCondition *IncrementalMigrationHttpOriginCondition `json:"OriginCondition,omitnil,omitempty" name:"OriginCondition"`
}

type IncrementalMigrationHttpOriginInfo struct {
	// 增量迁移源站地址信息。
	EndpointInfo *IncrementalMigrationHttpEndpointInfo `json:"EndpointInfo,omitnil,omitempty" name:"EndpointInfo"`

	// 增量迁移源站文件信息。
	FileInfo *IncrementalMigrationOriginFileInfo `json:"FileInfo,omitnil,omitempty" name:"FileInfo"`
}

type IncrementalMigrationHttpOriginParameter struct {
	// HTTP 头部透传信息。
	HttpHeaderInfo *IncrementalMigrationHttpHeaderInfo `json:"HttpHeaderInfo,omitnil,omitempty" name:"HttpHeaderInfo"`

	// 回源协议。取值有：
	// <li>HTTP：强制 HTTP；</li>
	// <li>HTTPS：强制 HTTPS；</li>
	// <li>FOLLOW：跟随请求协议。</li>若不填，默认取值 `FOLLOW`。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 请求参数透传模式。取值有：
	// <li>FOLLOW：全部透传；</li>
	// <li>IGNORE：忽略，全部不透传。</li> 默认取值 `FOLLOW`。
	QueryStringFollowMode *string `json:"QueryStringFollowMode,omitnil,omitempty" name:"QueryStringFollowMode"`

	// 重定向的 HTTP Code，目前仅支持 `301`，`302` 和 `307`。默认取值 `302`。
	HttpRedirectCode *uint64 `json:"HttpRedirectCode,omitnil,omitempty" name:"HttpRedirectCode"`

	// 源站重定向跟随模式。取值有：
	// <li>FOLLOW：跟随源站重定向；</li>
	// <li>IGNORE：忽略源站重定向。</li> 默认取值 `FOLLOW` 跟随源站重定向，即源站返回 `3xx` 时，会默认跟随至对应源站拉取数据。
	OriginRedirectionFollowMode *string `json:"OriginRedirectionFollowMode,omitnil,omitempty" name:"OriginRedirectionFollowMode"`
}

type IncrementalMigrationOriginFileInfo struct {
	// 文件前缀配置。
	PrefixConfig *IncrementalMigrationOriginPrefixConfig `json:"PrefixConfig,omitnil,omitempty" name:"PrefixConfig"`

	// 文件后缀配置。
	SuffixConfig *IncrementalMigrationOriginSuffixConfig `json:"SuffixConfig,omitnil,omitempty" name:"SuffixConfig"`

	// 固定文件配置。
	FixedFileConfig *IncrementalMigrationOriginFixedFileConfig `json:"FixedFileConfig,omitnil,omitempty" name:"FixedFileConfig"`
}

type IncrementalMigrationOriginFixedFileConfig struct {
	// 固定文件路径；如填充 `example/test.png`，则回源地址为： `http(s)://<回源域名>/example/test.png`。
	FixedFilePath *string `json:"FixedFilePath,omitnil,omitempty" name:"FixedFilePath"`
}

type IncrementalMigrationOriginPrefixConfig struct {
	// 源站地址前缀，如填充 `test/`，则回源地址为 `http(s)://<回源域名>/test/<文件名>`。
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

type IncrementalMigrationOriginSuffixConfig struct {
	// 文件后缀；如填充 `.ts` ，则回源地址为：`http(s)://<回源域名>/<文件名>.ts`。
	Suffix *string `json:"Suffix,omitnil,omitempty" name:"Suffix"`
}

type IncrementalMigrationStrategyInfo struct {
	// 策略 ID。
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 策略名称。
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// <b>策略生效的点播专业版[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 策略生效的存储桶 ID。
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// 源站类型。取值有：<li>HTTP：HTTP 源。</li>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// 回源源站配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpOriginConfig *IncrementalMigrationHttpOriginConfig `json:"HttpOriginConfig,omitnil,omitempty" name:"HttpOriginConfig"`
}

// Predefined struct for user
type ModifyIncrementalMigrationStrategyRequestParams struct {
	// <b>点播[专业版](/document/product/266/115396)[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 策略生效的存储桶 ID。
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// 增量迁移策略 ID。
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 策略名称。若不填充或填充空字符串，则不修改。
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 源站类型。取值有：<li>HTTP：HTTP 源。</li>若不填或填充空字符串，则不修改。
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// HTTP 回源源站配置，若不填则默认不修改。
	HttpOriginConfig *IncrementalMigrationHttpOriginConfig `json:"HttpOriginConfig,omitnil,omitempty" name:"HttpOriginConfig"`
}

type ModifyIncrementalMigrationStrategyRequest struct {
	*tchttp.BaseRequest
	
	// <b>点播[专业版](/document/product/266/115396)[应用](/document/product/266/14574) ID。</b>
	SubAppId *uint64 `json:"SubAppId,omitnil,omitempty" name:"SubAppId"`

	// 策略生效的存储桶 ID。
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// 增量迁移策略 ID。
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 策略名称。若不填充或填充空字符串，则不修改。
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 源站类型。取值有：<li>HTTP：HTTP 源。</li>若不填或填充空字符串，则不修改。
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// HTTP 回源源站配置，若不填则默认不修改。
	HttpOriginConfig *IncrementalMigrationHttpOriginConfig `json:"HttpOriginConfig,omitnil,omitempty" name:"HttpOriginConfig"`
}

func (r *ModifyIncrementalMigrationStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIncrementalMigrationStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubAppId")
	delete(f, "BucketId")
	delete(f, "StrategyId")
	delete(f, "StrategyName")
	delete(f, "OriginType")
	delete(f, "HttpOriginConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIncrementalMigrationStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIncrementalMigrationStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIncrementalMigrationStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIncrementalMigrationStrategyResponseParams `json:"Response"`
}

func (r *ModifyIncrementalMigrationStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIncrementalMigrationStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SortBy struct {
	// 排序字段。
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 排序方式，可选值有：
	// <li>Asc: 升序；</li>
	// <li>Desc: 降序。</li>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type StorageInfo struct {
	// 存储桶 ID。
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// 存储名称。
	StorageName *string `json:"StorageName,omitnil,omitempty" name:"StorageName"`

	// 存储所在区域。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 存储公网源站访问域名的状态，取值有：
	// <li>ONLINE：已生效；</li>
	// <li>DEPLOYING： 部署中。</li>
	InternetAccessDomainStatus *string `json:"InternetAccessDomainStatus,omitnil,omitempty" name:"InternetAccessDomainStatus"`

	// 存储公网源站访问域名。
	InternetAccessDomain *string `json:"InternetAccessDomain,omitnil,omitempty" name:"InternetAccessDomain"`

	// 存储的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}