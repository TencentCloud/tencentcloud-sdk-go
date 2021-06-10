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

package v20190923

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateSecretRequest struct {
	*tchttp.BaseRequest

	// 凭据名称，同一region内不可重复，最长128字节，使用字母、数字或者 - _ 的组合，第一个字符必须为字母或者数字。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 凭据版本，查询凭据信息时需要根据SecretName 和 VersionId进行查询，最长64 字节，使用字母、数字或者 - _ . 的组合并且以字母或数字开头。
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 描述信息，用于详细描述用途等，最大支持2048字节。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 指定对凭据进行加密的KMS CMK。如果为空则表示使用Secrets Manager为您默认创建的CMK进行加密。您也可以指定在同region 下自行创建的KMS CMK进行加密。
	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`

	// 二进制凭据信息base64编码后的明文。SecretBinary 和 SecretString 必须且只能设置一个，最大支持4096字节。
	SecretBinary *string `json:"SecretBinary,omitempty" name:"SecretBinary"`

	// 文本类型凭据信息明文（不需要进行base64编码）。SecretBinary 和 SecretString 必须且只能设置一个，，最大支持4096字节。
	SecretString *string `json:"SecretString,omitempty" name:"SecretString"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "VersionId")
	delete(f, "Description")
	delete(f, "KmsKeyId")
	delete(f, "SecretBinary")
	delete(f, "SecretString")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新创建的凭据名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// 新创建的凭据版本。
		VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

		// 标签操作的返回码. 0: 成功；1: 内部错误；2: 业务处理错误
	// 注意：此字段可能返回 null，表示取不到有效值。
		TagCode *uint64 `json:"TagCode,omitempty" name:"TagCode"`

		// 标签操作的返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		TagMsg *string `json:"TagMsg,omitempty" name:"TagMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecretRequest struct {
	*tchttp.BaseRequest

	// 指定需要删除的凭据名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 指定计划删除日期，单位（天），0（默认）表示立即删除， 1-30 表示预留的天数，超出该日期之后彻底删除。
	RecoveryWindowInDays *uint64 `json:"RecoveryWindowInDays,omitempty" name:"RecoveryWindowInDays"`
}

func (r *DeleteSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "RecoveryWindowInDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指定删除的凭据名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// 凭据删除的日期，unix时间戳。
		DeleteTime *int64 `json:"DeleteTime,omitempty" name:"DeleteTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecretVersionRequest struct {
	*tchttp.BaseRequest

	// 指定凭据名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 指定该名称下需要删除的凭据的版本号。
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *DeleteSecretVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecretVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecretVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 凭据名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// 凭据版本号。
		VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecretVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecretRequest struct {
	*tchttp.BaseRequest

	// 指定需要获取凭据详细信息的凭据名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *DescribeSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 凭据名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// 凭据描述信息。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 用于加密的KMS CMK ID。
		KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`

		// 创建者UIN。
		CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

		// 凭据状态：Enabled、Disabled、PendingDelete
		Status *string `json:"Status,omitempty" name:"Status"`

		// 删除日期，uinx 时间戳，非计划删除状态的凭据为0。
		DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`

		// 创建日期。
		CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableSecretRequest struct {
	*tchttp.BaseRequest

	// 指定停用的凭据名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *DisableSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisableSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 停用的凭据名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableSecretRequest struct {
	*tchttp.BaseRequest

	// 指定启用凭据的名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *EnableSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EnableSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 启用的凭据名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// region列表。
		Regions []*string `json:"Regions,omitempty" name:"Regions"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSecretValueRequest struct {
	*tchttp.BaseRequest

	// 指定凭据的名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 指定对应凭据的版本号。
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *GetSecretValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSecretValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSecretValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetSecretValueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 凭据的名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// 该凭据对应的版本号。
		VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

		// 在创建凭据(CreateSecret)时，如果指定的是二进制数据，则该字段为返回结果，并且使用base64进行编码，应用方需要进行base64解码后获取原始数据。SecretBinary和SecretString只有一个不为空。
		SecretBinary *string `json:"SecretBinary,omitempty" name:"SecretBinary"`

		// 在创建凭据(CreateSecret)时，如果指定的是普通文本数据，则该字段为返回结果。SecretBinary和SecretString只有一个不为空。
		SecretString *string `json:"SecretString,omitempty" name:"SecretString"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSecretValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSecretValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *GetServiceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetServiceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetServiceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true表示服务已开通，false 表示服务尚未开通。
		ServiceEnabled *bool `json:"ServiceEnabled,omitempty" name:"ServiceEnabled"`

		// 服务不可用类型： 0-未购买，1-正常， 2-欠费停服， 3-资源释放。
		InvalidType *int64 `json:"InvalidType,omitempty" name:"InvalidType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSecretVersionIdsRequest struct {
	*tchttp.BaseRequest

	// 凭据名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *ListSecretVersionIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSecretVersionIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSecretVersionIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListSecretVersionIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 凭据名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// VersionId列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Versions []*VersionInfo `json:"Versions,omitempty" name:"Versions"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSecretVersionIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSecretVersionIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSecretsRequest struct {
	*tchttp.BaseRequest

	// 查询列表的起始位置，以0开始，不设置默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单次查询返回的最大数量，0或不设置则使用默认值 20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 根据创建时间的排序方式，0或者不设置则使用降序排序， 1 表示升序排序。
	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`

	// 根据凭据状态进行过滤，默认为0表示查询全部，1 表示查询Enabed 凭据列表，2表示查询Disabled 凭据列表， 3 表示查询PendingDelete 凭据列表。
	State *uint64 `json:"State,omitempty" name:"State"`

	// 根据凭据名称进行过滤，为空表示不过滤。
	SearchSecretName *string `json:"SearchSecretName,omitempty" name:"SearchSecretName"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

func (r *ListSecretsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSecretsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderType")
	delete(f, "State")
	delete(f, "SearchSecretName")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSecretsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListSecretsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 根据State和SearchSecretName 筛选的凭据总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回凭据信息列表。
		SecretMetadatas []*SecretMetadata `json:"SecretMetadatas,omitempty" name:"SecretMetadatas"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSecretsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSecretsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutSecretValueRequest struct {
	*tchttp.BaseRequest

	// 指定需要增加版本的凭据名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 指定新增加的版本号，最长64 字节，使用字母、数字或者 - _ . 的组合并且以字母或数字开头。
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 二进制凭据信息，使用base64编码。SecretBinary 和 SecretString 必须且只能设置一个。
	SecretBinary *string `json:"SecretBinary,omitempty" name:"SecretBinary"`

	// 文本类型凭据信息明文（不需要进行base64编码），SecretBinary 和 SecretString 必须且只能设置一个。
	SecretString *string `json:"SecretString,omitempty" name:"SecretString"`
}

func (r *PutSecretValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutSecretValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "VersionId")
	delete(f, "SecretBinary")
	delete(f, "SecretString")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutSecretValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PutSecretValueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 凭据名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// 新增加的版本号。
		VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutSecretValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutSecretValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreSecretRequest struct {
	*tchttp.BaseRequest

	// 指定需要恢复的凭据名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *RestoreSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestoreSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 凭据名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestoreSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecretMetadata struct {

	// 凭据名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 凭据的描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用于加密凭据的KMS KeyId。
	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`

	// 创建者UIN。
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 凭据状态：Enabled、Disabled、PendingDelete
	Status *string `json:"Status,omitempty" name:"Status"`

	// 凭据删除日期，对于status为PendingDelete 的有效，unix时间戳。
	DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`

	// 凭据创建时间，unix时间戳。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 用于加密凭据的KMS CMK类型，DEFAULT 表示SecretsManager 创建的默认密钥， CUSTOMER 表示用户指定的密钥。
	KmsKeyType *string `json:"KmsKeyType,omitempty" name:"KmsKeyType"`
}

type Tag struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagFilter struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type UpdateDescriptionRequest struct {
	*tchttp.BaseRequest

	// 指定需要更新描述信息的凭据名。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 新的描述信息，最大长度2048个字节。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 凭据名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSecretRequest struct {
	*tchttp.BaseRequest

	// 指定需要更新凭据内容的名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 指定需要更新凭据内容的版本号。
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 新的凭据内容为二进制的场景使用该字段，并使用base64进行编码。SecretBinary 和 SecretString 只能一个不为空。
	SecretBinary *string `json:"SecretBinary,omitempty" name:"SecretBinary"`

	// 新的凭据内容为文本的场景使用该字段，不需要base64编码。SecretBinary 和 SecretString 只能一个不为空。
	SecretString *string `json:"SecretString,omitempty" name:"SecretString"`
}

func (r *UpdateSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "VersionId")
	delete(f, "SecretBinary")
	delete(f, "SecretString")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 凭据名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// 凭据版本号。
		VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VersionInfo struct {

	// 版本号。
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 创建时间，unix时间戳。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}
