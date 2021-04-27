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

package v20210303

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateAudioDepositRequest struct {
	*tchttp.BaseRequest

	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 对应数据Base64文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 数据Base64编码，大小不超过5M
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 资源访问链接 与fileContent必须二选一
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAudioDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceName")
	delete(f, "FileName")
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "FileContent")
	delete(f, "FileUrl")
	delete(f, "HashType")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return errors.New("CreateAudioDepositRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAudioDepositResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务ID 透传 长度最大不超过64
	// 注意：此字段可能返回 null，表示取不到有效值。
		BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

		// 请求成功，返回存证编码,用于查询存证后续业务数据
		EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAudioDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataDepositRequest struct {
	*tchttp.BaseRequest

	// 业务数据明文(json格式字符串)，最大256kb
	EvidenceInfo *string `json:"EvidenceInfo,omitempty" name:"EvidenceInfo"`

	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 数据hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceInfo")
	delete(f, "EvidenceName")
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "HashType")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return errors.New("CreateDataDepositRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataDepositResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务ID 透传 长度最大不超过64
	// 注意：此字段可能返回 null，表示取不到有效值。
		BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

		// 请求成功，返回存证编码,用于查询存证后续业务数据
		EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDocDepositRequest struct {
	*tchttp.BaseRequest

	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 对应数据Base64文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 数据Base64编码，大小不超过5M
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 资源访问链接 与fileContent必须二选一
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceName")
	delete(f, "FileName")
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "FileContent")
	delete(f, "FileUrl")
	delete(f, "HashType")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return errors.New("CreateDocDepositRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDocDepositResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务ID 透传 长度最大不超过64
	// 注意：此字段可能返回 null，表示取不到有效值。
		BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

		// 请求成功，返回存证编码,用于查询存证后续业务数据
		EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHashDepositNoCertRequest struct {
	*tchttp.BaseRequest

	// 数据hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 该字段为透传字段，方便调用方做业务处理， 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 业务扩展信息
	EvidenceInfo *string `json:"EvidenceInfo,omitempty" name:"EvidenceInfo"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHashDepositNoCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "HashType")
	delete(f, "EvidenceInfo")
	if len(f) > 0 {
		return errors.New("CreateHashDepositNoCertRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateHashDepositNoCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 透传字段
	// 注意：此字段可能返回 null，表示取不到有效值。
		BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

		// 存证编码
		EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

		// 上链时间
		EvidenceTime *string `json:"EvidenceTime,omitempty" name:"EvidenceTime"`

		// 区块链交易哈希
		EvidenceTxHash *string `json:"EvidenceTxHash,omitempty" name:"EvidenceTxHash"`

		// 区块高度
		BlockchainHeight *uint64 `json:"BlockchainHeight,omitempty" name:"BlockchainHeight"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHashDepositNoCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHashDepositNoSealRequest struct {
	*tchttp.BaseRequest

	// 数据hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 该字段为透传字段，方便调用方做业务处理， 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 业务扩展信息
	EvidenceInfo *string `json:"EvidenceInfo,omitempty" name:"EvidenceInfo"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHashDepositNoSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "HashType")
	delete(f, "EvidenceInfo")
	if len(f) > 0 {
		return errors.New("CreateHashDepositNoSealRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateHashDepositNoSealResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 透传字段
	// 注意：此字段可能返回 null，表示取不到有效值。
		BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

		// 存证编码
		EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

		// 上链时间
		EvidenceTime *string `json:"EvidenceTime,omitempty" name:"EvidenceTime"`

		// 区块链交易哈希
		EvidenceTxHash *string `json:"EvidenceTxHash,omitempty" name:"EvidenceTxHash"`

		// 区块高度
		BlockchainHeight *uint64 `json:"BlockchainHeight,omitempty" name:"BlockchainHeight"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHashDepositNoSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHashDepositRequest struct {
	*tchttp.BaseRequest

	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 数据hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 该字段为透传字段，方便调用方做业务处理， 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHashDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceName")
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "HashType")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return errors.New("CreateHashDepositRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateHashDepositResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 透传字段
	// 注意：此字段可能返回 null，表示取不到有效值。
		BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

		// 存证编码
		EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHashDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageDepositRequest struct {
	*tchttp.BaseRequest

	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 对应数据Base64文件名称 test.png
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 数据Base64编码，大小不超过5M
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 资源访问链接 与fileContent必须二选一
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceName")
	delete(f, "FileName")
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "FileContent")
	delete(f, "FileUrl")
	delete(f, "HashType")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return errors.New("CreateImageDepositRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateImageDepositResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务ID 透传 长度最大不超过64
	// 注意：此字段可能返回 null，表示取不到有效值。
		BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

		// 请求成功，返回存证编码,用于查询存证后续业务数据
		EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVideoDepositRequest struct {
	*tchttp.BaseRequest

	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 对应数据Base64文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件hash
	EvidenceHash *string `json:"EvidenceHash,omitempty" name:"EvidenceHash"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 数据Base64编码，大小不超过5M
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 资源访问链接 与fileContent必须二选一
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceName")
	delete(f, "FileName")
	delete(f, "EvidenceHash")
	delete(f, "BusinessId")
	delete(f, "FileContent")
	delete(f, "FileUrl")
	delete(f, "HashType")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return errors.New("CreateVideoDepositRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateVideoDepositResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务ID 透传 长度最大不超过64
	// 注意：此字段可能返回 null，表示取不到有效值。
		BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

		// 请求成功，返回存证编码,用于查询存证后续业务数据
		EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWebpageDepositRequest struct {
	*tchttp.BaseRequest

	// 存证名称(长度最大30)
	EvidenceName *string `json:"EvidenceName,omitempty" name:"EvidenceName"`

	// 网页链接
	EvidenceUrl *string `json:"EvidenceUrl,omitempty" name:"EvidenceUrl"`

	// 业务ID 透传 长度最大不超过64
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 算法类型 0 SM3, 1 SHA256, 2 SHA384 默认0
	HashType *uint64 `json:"HashType,omitempty" name:"HashType"`

	// 存证描述
	EvidenceDescription *string `json:"EvidenceDescription,omitempty" name:"EvidenceDescription"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebpageDepositRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceName")
	delete(f, "EvidenceUrl")
	delete(f, "BusinessId")
	delete(f, "HashType")
	delete(f, "EvidenceDescription")
	if len(f) > 0 {
		return errors.New("CreateWebpageDepositRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateWebpageDepositResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务ID 透传 长度最大不超过64
	// 注意：此字段可能返回 null，表示取不到有效值。
		BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

		// 请求成功，返回存证编码,用于查询存证后续业务数据
		EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebpageDepositResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDepositCertRequest struct {
	*tchttp.BaseRequest

	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDepositCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceId")
	if len(f) > 0 {
		return errors.New("GetDepositCertRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetDepositCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 存证编码
		EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

		// 存证证书文件临时链接
		EvidenceCert *string `json:"EvidenceCert,omitempty" name:"EvidenceCert"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDepositCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDepositFileRequest struct {
	*tchttp.BaseRequest

	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDepositFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceId")
	if len(f) > 0 {
		return errors.New("GetDepositFileRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetDepositFileResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 存证编号
		EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

		// 存证文件临时链接
		EvidenceFile *string `json:"EvidenceFile,omitempty" name:"EvidenceFile"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDepositFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDepositInfoRequest struct {
	*tchttp.BaseRequest

	// 存证编码
	EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDepositInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EvidenceId")
	if len(f) > 0 {
		return errors.New("GetDepositInfoRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetDepositInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 存证编号
		EvidenceId *string `json:"EvidenceId,omitempty" name:"EvidenceId"`

		// 上链时间
		EvidenceTime *string `json:"EvidenceTime,omitempty" name:"EvidenceTime"`

		// 区块链交易哈希
		EvidenceTxHash *string `json:"EvidenceTxHash,omitempty" name:"EvidenceTxHash"`

		// 区块高度
		BlockchainHeight *int64 `json:"BlockchainHeight,omitempty" name:"BlockchainHeight"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDepositInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
