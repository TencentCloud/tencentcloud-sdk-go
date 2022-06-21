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

package v20210514

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateAudioDepositRequestParams struct {
	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 数据Base64编码，大小不超过5M
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 带后缀的文件名称，如music.mp3
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

type CreateAudioDepositRequest struct {
	*tchttp.BaseRequest
	
	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 数据Base64编码，大小不超过5M
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 带后缀的文件名称，如music.mp3
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *CreateAudioDepositRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAudioDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceName")
	delete(f, "FileContent")
	delete(f, "FileName")
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "HashType")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAudioDepositRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAudioDepositResponseParams struct {
	// 业务ID 透传 长度最大不超过64
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 请求成功，返回存证编码,用于查询存证后续业务数据
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAudioDepositResponse struct {
	*tchttp.BaseResponse
	Response *CreateAudioDepositResponseParams `json:"Response"`
}

func (r *CreateAudioDepositResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAudioDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataDepositRequestParams struct {
	// 业务数据明文(json格式字符串)，最大256kb
	EvidenceInfo *string `json:"EvidenceInfo,omitempty" name:"EvidenceInfo"`

	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

type CreateDataDepositRequest struct {
	*tchttp.BaseRequest
	
	// 业务数据明文(json格式字符串)，最大256kb
	EvidenceInfo *string `json:"EvidenceInfo,omitempty" name:"EvidenceInfo"`

	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *CreateDataDepositRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceInfo")
	delete(f, "EvidenceName")
	delete(f, "BusinessId")
	delete(f, "HashType")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataDepositRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataDepositResponseParams struct {
	// 业务ID 透传 长度最大不超过64
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 请求成功，返回存证编码,用于查询存证后续业务数据
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDataDepositResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataDepositResponseParams `json:"Response"`
}

func (r *CreateDataDepositResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocDepositRequestParams struct {
	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 数据Base64编码，大小不超过5M
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 带后缀的文件名称，如 test.doc
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

type CreateDocDepositRequest struct {
	*tchttp.BaseRequest
	
	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 数据Base64编码，大小不超过5M
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 带后缀的文件名称，如 test.doc
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *CreateDocDepositRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceName")
	delete(f, "FileContent")
	delete(f, "FileName")
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "HashType")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDocDepositRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocDepositResponseParams struct {
	// 业务ID 透传 长度最大不超过64
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 请求成功，返回存证编码,用于查询存证后续业务数据
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDocDepositResponse struct {
	*tchttp.BaseResponse
	Response *CreateDocDepositResponseParams `json:"Response"`
}

func (r *CreateDocDepositResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHashDepositNoCertRequestParams struct {
	// 数据hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 该字段为透传字段，方便调用方做业务处理， 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 业务扩展信息
	EvidenceInfo *string `json:"EvidenceInfo,omitempty" name:"EvidenceInfo"`
}

type CreateHashDepositNoCertRequest struct {
	*tchttp.BaseRequest
	
	// 数据hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 该字段为透传字段，方便调用方做业务处理， 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 业务扩展信息
	EvidenceInfo *string `json:"EvidenceInfo,omitempty" name:"EvidenceInfo"`
}

func (r *CreateHashDepositNoCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHashDepositNoCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "EvidenceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHashDepositNoCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHashDepositNoCertResponseParams struct {
	// 透传字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateHashDepositNoCertResponse struct {
	*tchttp.BaseResponse
	Response *CreateHashDepositNoCertResponseParams `json:"Response"`
}

func (r *CreateHashDepositNoCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHashDepositNoCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHashDepositNoSealRequestParams struct {
	// 数据hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 该字段为透传字段，方便调用方做业务处理， 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 业务扩展信息
	EvidenceInfo *string `json:"EvidenceInfo,omitempty" name:"EvidenceInfo"`
}

type CreateHashDepositNoSealRequest struct {
	*tchttp.BaseRequest
	
	// 数据hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 该字段为透传字段，方便调用方做业务处理， 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 业务扩展信息
	EvidenceInfo *string `json:"EvidenceInfo,omitempty" name:"EvidenceInfo"`
}

func (r *CreateHashDepositNoSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHashDepositNoSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "EvidenceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHashDepositNoSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHashDepositNoSealResponseParams struct {
	// 透传字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateHashDepositNoSealResponse struct {
	*tchttp.BaseResponse
	Response *CreateHashDepositNoSealResponseParams `json:"Response"`
}

func (r *CreateHashDepositNoSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHashDepositNoSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHashDepositRequestParams struct {
	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 数据hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 该字段为透传字段，方便调用方做业务处理， 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

type CreateHashDepositRequest struct {
	*tchttp.BaseRequest
	
	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 数据hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 该字段为透传字段，方便调用方做业务处理， 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *CreateHashDepositRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHashDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceName")
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHashDepositRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHashDepositResponseParams struct {
	// 透传字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateHashDepositResponse struct {
	*tchttp.BaseResponse
	Response *CreateHashDepositResponseParams `json:"Response"`
}

func (r *CreateHashDepositResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHashDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageDepositRequestParams struct {
	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 数据Base64编码，大小不超过5M
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 带后缀的文件名称，如 test.png
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

type CreateImageDepositRequest struct {
	*tchttp.BaseRequest
	
	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 数据Base64编码，大小不超过5M
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 带后缀的文件名称，如 test.png
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *CreateImageDepositRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceName")
	delete(f, "FileContent")
	delete(f, "FileName")
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "HashType")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageDepositRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageDepositResponseParams struct {
	// 业务ID 透传 长度最大不超过64
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 请求成功，返回存证编码,用于查询存证后续业务数据
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateImageDepositResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageDepositResponseParams `json:"Response"`
}

func (r *CreateImageDepositResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoDepositRequestParams struct {
	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 数据Base64编码，大小不超过5M
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 带后缀的文件名称，如music.mkv
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

type CreateVideoDepositRequest struct {
	*tchttp.BaseRequest
	
	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 数据Base64编码，大小不超过5M
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 带后缀的文件名称，如music.mkv
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *CreateVideoDepositRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceName")
	delete(f, "FileContent")
	delete(f, "FileName")
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "HashType")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVideoDepositRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoDepositResponseParams struct {
	// 业务ID 透传 长度最大不超过64
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 请求成功，返回存证编码,用于查询存证后续业务数据
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVideoDepositResponse struct {
	*tchttp.BaseResponse
	Response *CreateVideoDepositResponseParams `json:"Response"`
}

func (r *CreateVideoDepositResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDepositCertRequestParams struct {
	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`
}

type GetDepositCertRequest struct {
	*tchttp.BaseRequest
	
	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`
}

func (r *GetDepositCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDepositCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDepositCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDepositCertResponseParams struct {
	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 存证证书文件临时链接
	EvidenceCert *string `json:"EvidenceCert,omitempty" name:"EvidenceCert"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDepositCertResponse struct {
	*tchttp.BaseResponse
	Response *GetDepositCertResponseParams `json:"Response"`
}

func (r *GetDepositCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDepositCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDepositFileRequestParams struct {
	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`
}

type GetDepositFileRequest struct {
	*tchttp.BaseRequest
	
	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`
}

func (r *GetDepositFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDepositFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDepositFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDepositFileResponseParams struct {
	// 存证编号
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 存证文件临时链接
	EvidenceFile *string `json:"EvidenceFile,omitempty" name:"EvidenceFile"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDepositFileResponse struct {
	*tchttp.BaseResponse
	Response *GetDepositFileResponseParams `json:"Response"`
}

func (r *GetDepositFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDepositFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDepositInfoRequestParams struct {
	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`
}

type GetDepositInfoRequest struct {
	*tchttp.BaseRequest
	
	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`
}

func (r *GetDepositInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDepositInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDepositInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDepositInfoResponseParams struct {
	// 存证编号
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 上链时间
	EvidenceTime *string `json:"EvidenceTime,omitempty" name:"EvidenceTime"`

	// 区块链交易哈希
	EvidenceTxHash *string `json:"EvidenceTxHash,omitempty" name:"EvidenceTxHash"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDepositInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetDepositInfoResponseParams `json:"Response"`
}

func (r *GetDepositInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDepositInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyEvidenceBlockChainTxHashRequestParams struct {
	// 区块链交易 hash，在“存证基本信息查询（GetDepositInfo）”接口中可以获取。
	EvidenceTxHash *string `json:"EvidenceTxHash,omitempty" name:"EvidenceTxHash"`
}

type VerifyEvidenceBlockChainTxHashRequest struct {
	*tchttp.BaseRequest
	
	// 区块链交易 hash，在“存证基本信息查询（GetDepositInfo）”接口中可以获取。
	EvidenceTxHash *string `json:"EvidenceTxHash,omitempty" name:"EvidenceTxHash"`
}

func (r *VerifyEvidenceBlockChainTxHashRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyEvidenceBlockChainTxHashRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceTxHash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyEvidenceBlockChainTxHashRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyEvidenceBlockChainTxHashResponseParams struct {
	// 核验结果，true为核验成功，fals为核验失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 存证时间，仅当核验结果为true时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	EvidenceTime *string `json:"EvidenceTime,omitempty" name:"EvidenceTime"`

	// 存证编码，仅当核验结果为true时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyEvidenceBlockChainTxHashResponse struct {
	*tchttp.BaseResponse
	Response *VerifyEvidenceBlockChainTxHashResponseParams `json:"Response"`
}

func (r *VerifyEvidenceBlockChainTxHashResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyEvidenceBlockChainTxHashResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyEvidenceHashRequestParams struct {
	// 存证内容hash，hash类型即为用户在存证时所用或所选的hash类型
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`
}

type VerifyEvidenceHashRequest struct {
	*tchttp.BaseRequest
	
	// 存证内容hash，hash类型即为用户在存证时所用或所选的hash类型
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`
}

func (r *VerifyEvidenceHashRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyEvidenceHashRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceHash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyEvidenceHashRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyEvidenceHashResponseParams struct {
	// 核验结果，true为核验成功，false为核验失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyEvidenceHashResponse struct {
	*tchttp.BaseResponse
	Response *VerifyEvidenceHashResponseParams `json:"Response"`
}

func (r *VerifyEvidenceHashResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyEvidenceHashResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}