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

type CreateProductSecretRequest struct {
	*tchttp.BaseRequest

	// 凭据名称，同一region内不可重复，最长128字节，使用字母、数字或者 - _ 的组合，第一个字符必须为字母或者数字。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 用户账号名前缀，由用户自行指定，长度限定在8个字符以内，
	// 可选字符集包括：
	// 数字字符：[0, 9]，
	// 小写字符：[a, z]，
	// 大写字符：[A, Z]，
	// 特殊字符(全英文符号)：下划线(_)，
	// 前缀必须以大写或小写字母开头。
	UserNamePrefix *string `json:"UserNamePrefix,omitempty" name:"UserNamePrefix"`

	// 凭据所绑定的云产品名称，如Mysql，可以通过DescribeSupportedProducts接口获取所支持的云产品名称。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 云产品实例ID。
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 账号的域名，IP形式，支持填入%。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 将凭据与云产品实例绑定时，需要授予的权限列表。
	PrivilegesList []*ProductPrivilegeUnit `json:"PrivilegesList,omitempty" name:"PrivilegesList"`

	// 描述信息，用于详细描述用途等，最大支持2048字节。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 指定对凭据进行加密的KMS CMK。
	// 如果为空则表示使用Secrets Manager为您默认创建的CMK进行加密。
	// 您也可以指定在同region 下自行创建的KMS CMK进行加密。
	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`

	// 标签列表。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 用户自定义的开始轮转时间，格式：2006-01-02 15:04:05。
	// 当EnableRotation为True时，此参数必填。
	RotationBeginTime *string `json:"RotationBeginTime,omitempty" name:"RotationBeginTime"`

	// 是否开启轮转
	// True -- 开启
	// False -- 不开启
	// 如果不指定，默认为False。
	EnableRotation *bool `json:"EnableRotation,omitempty" name:"EnableRotation"`

	// 轮转周期，以天为单位，默认为1天。
	RotationFrequency *int64 `json:"RotationFrequency,omitempty" name:"RotationFrequency"`
}

func (r *CreateProductSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProductSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "UserNamePrefix")
	delete(f, "ProductName")
	delete(f, "InstanceID")
	delete(f, "Domains")
	delete(f, "PrivilegesList")
	delete(f, "Description")
	delete(f, "KmsKeyId")
	delete(f, "Tags")
	delete(f, "RotationBeginTime")
	delete(f, "EnableRotation")
	delete(f, "RotationFrequency")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProductSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProductSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的凭据名称。
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// 标签操作的返回码. 0: 成功；1: 内部错误；2: 业务处理错误。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TagCode *uint64 `json:"TagCode,omitempty" name:"TagCode"`

		// 标签操作的返回信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TagMsg *string `json:"TagMsg,omitempty" name:"TagMsg"`

		// 创建云产品凭据异步任务ID号。
		FlowID *int64 `json:"FlowID,omitempty" name:"FlowID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProductSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProductSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

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

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest

	// 异步任务ID号。
	FlowID *int64 `json:"FlowID,omitempty" name:"FlowID"`
}

func (r *DescribeAsyncRequestInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncRequestInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0:处理中，1:处理成功，2:处理失败
		TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

		// 任务描述信息。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRotationDetailRequest struct {
	*tchttp.BaseRequest

	// 指定需要获取凭据轮转详细信息的凭据名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *DescribeRotationDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRotationDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRotationDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRotationDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 否允许轮转，True表示开启轮转，False表示禁止轮转。
		EnableRotation *bool `json:"EnableRotation,omitempty" name:"EnableRotation"`

		// 轮转的频率，以天为单位，默认为1天。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Frequency *int64 `json:"Frequency,omitempty" name:"Frequency"`

		// 最近一次轮转的时间，显式可见的时间字符串，格式 2006-01-02 15:04:05。
	// 注意：此字段可能返回 null，表示取不到有效值。
		LatestRotateTime *string `json:"LatestRotateTime,omitempty" name:"LatestRotateTime"`

		// 下一次开始轮转的时间，显式可见的时间字符串，格式 2006-01-02 15:04:05。
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextRotateBeginTime *string `json:"NextRotateBeginTime,omitempty" name:"NextRotateBeginTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRotationDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRotationDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRotationHistoryRequest struct {
	*tchttp.BaseRequest

	// 指定需要获取凭据轮转历史的凭据名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *DescribeRotationHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRotationHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRotationHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRotationHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 版本号列表。
		VersionIDs []*string `json:"VersionIDs,omitempty" name:"VersionIDs"`

		// 版本号个数，可以给用户展示的版本号个数上限为10个。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRotationHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRotationHistoryResponse) FromJsonString(s string) error {
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

		// 凭据状态：Enabled、Disabled、PendingDelete, Creating, Failed。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 删除日期，uinx 时间戳，非计划删除状态的凭据为0。
		DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`

		// 创建日期。
		CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

		// 0 --  用户自定义凭据类型；1 -- 云产品凭据类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
		SecretType *int64 `json:"SecretType,omitempty" name:"SecretType"`

		// 云产品名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

		// 云产品实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ResourceID *string `json:"ResourceID,omitempty" name:"ResourceID"`

		// 是否开启轮转：True -- 开启轮转；False -- 禁止轮转。
	// 注意：此字段可能返回 null，表示取不到有效值。
		RotationStatus *bool `json:"RotationStatus,omitempty" name:"RotationStatus"`

		// 轮转周期，默认以天为单位。
	// 注意：此字段可能返回 null，表示取不到有效值。
		RotationFrequency *int64 `json:"RotationFrequency,omitempty" name:"RotationFrequency"`

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

type DescribeSupportedProductsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSupportedProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupportedProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportedProductsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 支持的产品列表。
		Products []*string `json:"Products,omitempty" name:"Products"`

		// 支持的产品个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSupportedProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedProductsResponse) FromJsonString(s string) error {
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
	// 对于云产品凭据如Mysql凭据，通过指定凭据名称和历史版本号来获取历史轮转凭据的明文信息，如果要获取当前正在使用的凭据版本的明文，需要将版本号指定为：SSM_Current。
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

		// 在创建凭据(CreateSecret)时，如果指定的是二进制数据，则该字段为返回结果，并且使用base64进行编码，应用方需要进行base64解码后获取原始数据。
	// SecretBinary和SecretString只有一个不为空。
		SecretBinary *string `json:"SecretBinary,omitempty" name:"SecretBinary"`

		// 在创建凭据(CreateSecret)时，如果指定的是普通文本数据，则该字段为返回结果。
	// SecretBinary和SecretString只有一个不为空。
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

	// 根据凭据状态进行过滤。
	// 默认为0表示查询全部。
	// 1 --  表示查询Enabled 凭据列表。
	// 2 --  表示查询Disabled 凭据列表。
	// 3 --  表示查询PendingDelete 凭据列表。
	// 4 --  表示PendingCreate。
	// 5 --  表示CreateFailed。
	// 其中状态PendingCreate和CreateFailed只有在SecretType为云产品凭据时生效
	State *uint64 `json:"State,omitempty" name:"State"`

	// 根据凭据名称进行过滤，为空表示不过滤。
	SearchSecretName *string `json:"SearchSecretName,omitempty" name:"SearchSecretName"`

	// 标签过滤条件。
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 0  -- 表示用户自定义凭据，默认为0。
	// 1  -- 表示用户云产品凭据。
	// 这个参数只能在云产品凭据(1)和用户自定义凭据(0)中二选一。
	SecretType *uint64 `json:"SecretType,omitempty" name:"SecretType"`
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
	delete(f, "SecretType")
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

type ProductPrivilegeUnit struct {

	// 权限名称，当前可选：
	// GlobalPrivileges
	// DatabasePrivileges
	// TablePrivileges
	// ColumnPrivileges
	// 
	// 当权限为DatabasePrivileges时，必须通过参数Database指定数据库名；
	// 
	// 当权限为TablePrivileges时，必须通过参数Database和TableName指定数据库名以及数据库中的表名；
	// 
	// 当权限为ColumnPrivileges时，必须通过参数Database、TableName和CoulmnName指定数据库、数据库中的表名以及表中的列名。
	PrivilegeName *string `json:"PrivilegeName,omitempty" name:"PrivilegeName"`

	// 权限列表。
	// 对于Mysql产品来说，可选权限值为：
	// 
	// 1. GlobalPrivileges 中权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "PROCESS", "DROP","REFERENCES","INDEX","ALTER","SHOW DATABASES","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	// 注意，不传该参数表示清除该权限。
	// 
	// 2. DatabasePrivileges 权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "DROP","REFERENCES","INDEX","ALTER","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	// 注意，不传该参数表示清除该权限。
	// 
	// 3. TablePrivileges 权限的可选值为：权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "DROP","REFERENCES","INDEX","ALTER","CREATE VIEW","SHOW VIEW", "TRIGGER"。
	// 注意，不传该参数表示清除该权限。
	// 
	// 4. ColumnPrivileges 权限的可选值为："SELECT","INSERT","UPDATE","REFERENCES"。
	// 注意，不传该参数表示清除该权限。
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges"`

	// 仅当PrivilegeName为DatabasePrivileges时这个值才有效。
	Database *string `json:"Database,omitempty" name:"Database"`

	// 仅当PrivilegeName为TablePrivileges时这个值才有效，并且此时需要填充Database显式指明所在的数据库实例。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 仅当PrivilegeName为ColumnPrivileges时这个值才生效，并且此时必须填充：
	// Database - 显式指明所在的数据库实例。
	// TableName - 显式指明所在表
	ColumnName *string `json:"ColumnName,omitempty" name:"ColumnName"`
}

type PutSecretValueRequest struct {
	*tchttp.BaseRequest

	// 指定需要增加版本的凭据名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 指定新增加的版本号，最长64 字节，使用字母、数字或者 - _ . 的组合并且以字母或数字开头。
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 二进制凭据信息，使用base64编码。
	// SecretBinary 和 SecretString 必须且只能设置一个。
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

type RotateProductSecretRequest struct {
	*tchttp.BaseRequest

	// 需要轮转的凭据名。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *RotateProductSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RotateProductSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RotateProductSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RotateProductSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 轮转异步任务ID号。
		FlowID *int64 `json:"FlowID,omitempty" name:"FlowID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RotateProductSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RotateProductSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecretMetadata struct {

	// 凭据名称
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 凭据的描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用于加密凭据的KMS KeyId
	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`

	// 创建者UIN
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 凭据状态：Enabled、Disabled、PendingDelete、Creating、Failed
	Status *string `json:"Status,omitempty" name:"Status"`

	// 凭据删除日期，对于status为PendingDelete 的有效，unix时间戳
	DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`

	// 凭据创建时间，unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 用于加密凭据的KMS CMK类型，DEFAULT 表示SecretsManager 创建的默认密钥， CUSTOMER 表示用户指定的密钥
	KmsKeyType *string `json:"KmsKeyType,omitempty" name:"KmsKeyType"`

	// 1:--开启轮转；0--禁止轮转
	// 注意：此字段可能返回 null，表示取不到有效值。
	RotationStatus *int64 `json:"RotationStatus,omitempty" name:"RotationStatus"`

	// 下一次轮转开始时间，uinx 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextRotationTime *uint64 `json:"NextRotationTime,omitempty" name:"NextRotationTime"`

	// 0 -- 用户自定义凭据；1 -- 云产品凭据
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretType *int64 `json:"SecretType,omitempty" name:"SecretType"`

	// 云产品名称，仅在SecretType为1，即凭据类型为云产品凭据时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
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

type UpdateRotationStatusRequest struct {
	*tchttp.BaseRequest

	// 云产品凭据名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 是否开启轮转。
	// True -- 开启轮转；
	// False -- 禁止轮转。
	EnableRotation *bool `json:"EnableRotation,omitempty" name:"EnableRotation"`

	// 轮转周期，以天为单位，最小为30天，最大为365天。
	Frequency *int64 `json:"Frequency,omitempty" name:"Frequency"`

	// 用户设置的期望开始轮转时间，格式为：2006-01-02 15:04:05。
	// 当EnableRotation为True时，如果不填RotationBeginTime，则默认填充为当前时间。
	RotationBeginTime *string `json:"RotationBeginTime,omitempty" name:"RotationBeginTime"`
}

func (r *UpdateRotationStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRotationStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "EnableRotation")
	delete(f, "Frequency")
	delete(f, "RotationBeginTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRotationStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRotationStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRotationStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRotationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSecretRequest struct {
	*tchttp.BaseRequest

	// 指定需要更新凭据内容的名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 指定需要更新凭据内容的版本号。
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 新的凭据内容为二进制的场景使用该字段，并使用base64进行编码。
	// SecretBinary 和 SecretString 只能一个不为空。
	SecretBinary *string `json:"SecretBinary,omitempty" name:"SecretBinary"`

	// 新的凭据内容为文本的场景使用该字段，不需要base64编码SecretBinary 和 SecretString 只能一个不为空。
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
