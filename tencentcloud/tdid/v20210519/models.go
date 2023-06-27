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

package v20210519

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddLabelRequestParams struct {
	// 标签ID
	LabelId *uint64 `json:"LabelId,omitempty" name:"LabelId"`

	// tdid
	Did *string `json:"Did,omitempty" name:"Did"`
}

type AddLabelRequest struct {
	*tchttp.BaseRequest
	
	// 标签ID
	LabelId *uint64 `json:"LabelId,omitempty" name:"LabelId"`

	// tdid
	Did *string `json:"Did,omitempty" name:"Did"`
}

func (r *AddLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LabelId")
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddLabelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddLabelResponse struct {
	*tchttp.BaseResponse
	Response *AddLabelResponseParams `json:"Response"`
}

func (r *AddLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BcosClusterItem struct {
	// 网络索引id
	ChainId *uint64 `json:"ChainId,omitempty" name:"ChainId"`

	// 网络名称
	ChainName *string `json:"ChainName,omitempty" name:"ChainName"`

	// 机构数量
	AgencyCount *uint64 `json:"AgencyCount,omitempty" name:"AgencyCount"`

	// 联盟id
	ConsortiumId *uint64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 网络状态
	ChainStatus *uint64 `json:"ChainStatus,omitempty" name:"ChainStatus"`

	// 资源 id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织名称
	ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

	// 机构id
	AgencyId *uint64 `json:"AgencyId,omitempty" name:"AgencyId"`

	// 续费状态
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 网络模式
	TotalNetworkNode *uint64 `json:"TotalNetworkNode,omitempty" name:"TotalNetworkNode"`

	// 创建节点数
	TotalCreateNode *uint64 `json:"TotalCreateNode,omitempty" name:"TotalCreateNode"`

	// 总群组数量
	TotalGroups *uint64 `json:"TotalGroups,omitempty" name:"TotalGroups"`
}

// Predefined struct for user
type CheckChainRequestParams struct {
	// 群组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// did服务机构名称
	AgencyName *string `json:"AgencyName,omitempty" name:"AgencyName"`
}

type CheckChainRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// did服务机构名称
	AgencyName *string `json:"AgencyName,omitempty" name:"AgencyName"`
}

func (r *CheckChainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckChainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ClusterId")
	delete(f, "AgencyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckChainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckChainResponseParams struct {
	// 1为盟主，0为非盟主
	RoleType *int64 `json:"RoleType,omitempty" name:"RoleType"`

	// 链ID
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckChainResponse struct {
	*tchttp.BaseResponse
	Response *CheckChainResponseParams `json:"Response"`
}

func (r *CheckChainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckChainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsortiumItem struct {
	// 联盟id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 联盟名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CptIssueRank struct {
	// 模板名称
	CptName *string `json:"CptName,omitempty" name:"CptName"`

	// 名次
	Rank *int64 `json:"Rank,omitempty" name:"Rank"`

	// 颁发量
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 应用名称
	ApplyName *string `json:"ApplyName,omitempty" name:"ApplyName"`

	// 应用ID
	ApplyId *uint64 `json:"ApplyId,omitempty" name:"ApplyId"`
}

// Predefined struct for user
type CreateCredentialRequestParams struct {
	// 参数集合，详见示例
	FunctionArg *FunctionArg `json:"FunctionArg,omitempty" name:"FunctionArg"`

	// 参数集合，详见示例
	TransactionArg *TransactionArg `json:"TransactionArg,omitempty" name:"TransactionArg"`

	// 版本
	VersionCredential *string `json:"VersionCredential,omitempty" name:"VersionCredential"`

	// 是否未签名
	UnSigned *bool `json:"UnSigned,omitempty" name:"UnSigned"`
}

type CreateCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 参数集合，详见示例
	FunctionArg *FunctionArg `json:"FunctionArg,omitempty" name:"FunctionArg"`

	// 参数集合，详见示例
	TransactionArg *TransactionArg `json:"TransactionArg,omitempty" name:"TransactionArg"`

	// 版本
	VersionCredential *string `json:"VersionCredential,omitempty" name:"VersionCredential"`

	// 是否未签名
	UnSigned *bool `json:"UnSigned,omitempty" name:"UnSigned"`
}

func (r *CreateCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionArg")
	delete(f, "TransactionArg")
	delete(f, "VersionCredential")
	delete(f, "UnSigned")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCredentialResponseParams struct {
	// Credential的具体信息
	CredentialData *string `json:"CredentialData,omitempty" name:"CredentialData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateCredentialResponseParams `json:"Response"`
}

func (r *CreateCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSelectiveCredentialRequestParams struct {
	// 参数集合
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitempty" name:"FunctionArg"`

	// 批露策略id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type CreateSelectiveCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 参数集合
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitempty" name:"FunctionArg"`

	// 批露策略id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *CreateSelectiveCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSelectiveCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionArg")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSelectiveCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSelectiveCredentialResponseParams struct {
	// 凭证字符串
	CredentialData *string `json:"CredentialData,omitempty" name:"CredentialData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSelectiveCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateSelectiveCredentialResponseParams `json:"Response"`
}

func (r *CreateSelectiveCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSelectiveCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPrivateKeyRequestParams struct {
	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 私钥
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

type CreateTDidByPrivateKeyRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 私钥
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

func (r *CreateTDidByPrivateKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPrivateKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "PrivateKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByPrivateKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPrivateKeyResponseParams struct {
	// did的具体信息
	Did *string `json:"Did,omitempty" name:"Did"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTDidByPrivateKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidByPrivateKeyResponseParams `json:"Response"`
}

func (r *CreateTDidByPrivateKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPrivateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPublicKeyRequestParams struct {
	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 身份公钥
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// 加密公钥
	EncryptPubKey *string `json:"EncryptPubKey,omitempty" name:"EncryptPubKey"`
}

type CreateTDidByPublicKeyRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 身份公钥
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// 加密公钥
	EncryptPubKey *string `json:"EncryptPubKey,omitempty" name:"EncryptPubKey"`
}

func (r *CreateTDidByPublicKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPublicKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "PublicKey")
	delete(f, "EncryptPubKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByPublicKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPublicKeyResponseParams struct {
	// did具体信息
	Did *string `json:"Did,omitempty" name:"Did"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTDidByPublicKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidByPublicKeyResponseParams `json:"Response"`
}

func (r *CreateTDidByPublicKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidRequestParams struct {
	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署机构为1，否则为0
	Relegation *uint64 `json:"Relegation,omitempty" name:"Relegation"`
}

type CreateTDidRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署机构为1，否则为0
	Relegation *uint64 `json:"Relegation,omitempty" name:"Relegation"`
}

func (r *CreateTDidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ClusterId")
	delete(f, "Relegation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidResponseParams struct {
	// TDID
	Did *string `json:"Did,omitempty" name:"Did"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTDidResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidResponseParams `json:"Response"`
}

func (r *CreateTDidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CredentialStatus struct {
	// 凭证唯一id
	CredentialId *string `json:"CredentialId,omitempty" name:"CredentialId"`

	// 凭证状态（0：吊销；1：有效）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 凭证颁发者Did
	Issuer *string `json:"Issuer,omitempty" name:"Issuer"`

	// 凭证摘要
	// 注意：此字段可能返回 null，表示取不到有效值。
	Digest *string `json:"Digest,omitempty" name:"Digest"`

	// 凭证签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Signature *string `json:"Signature,omitempty" name:"Signature"`

	// 更新时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeStamp *uint64 `json:"TimeStamp,omitempty" name:"TimeStamp"`
}

type FunctionArg struct {
	// CPT ID
	CptId *uint64 `json:"CptId,omitempty" name:"CptId"`

	// 签发者 did
	Issuer *string `json:"Issuer,omitempty" name:"Issuer"`

	// 过期时间
	ExpirationDate *string `json:"ExpirationDate,omitempty" name:"ExpirationDate"`

	// 声明
	ClaimJson *string `json:"ClaimJson,omitempty" name:"ClaimJson"`
}

// Predefined struct for user
type GetAgencyTDidRequestParams struct {
	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type GetAgencyTDidRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetAgencyTDidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAgencyTDidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAgencyTDidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAgencyTDidResponseParams struct {
	// 固定前缀
	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`

	// did详情
	Identity []*Identity `json:"Identity,omitempty" name:"Identity"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetAgencyTDidResponse struct {
	*tchttp.BaseResponse
	Response *GetAgencyTDidResponseParams `json:"Response"`
}

func (r *GetAgencyTDidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAgencyTDidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAuthorityIssuerRequestParams struct {
	// tdid
	Did *string `json:"Did,omitempty" name:"Did"`
}

type GetAuthorityIssuerRequest struct {
	*tchttp.BaseRequest
	
	// tdid
	Did *string `json:"Did,omitempty" name:"Did"`
}

func (r *GetAuthorityIssuerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthorityIssuerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAuthorityIssuerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAuthorityIssuerResponseParams struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 区块链网络id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 区块链群组id
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 权威机构did
	Did *string `json:"Did,omitempty" name:"Did"`

	// 机构备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 注册时间
	RegisterTime *string `json:"RegisterTime,omitempty" name:"RegisterTime"`

	// 认证时间
	RecognizeTime *string `json:"RecognizeTime,omitempty" name:"RecognizeTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetAuthorityIssuerResponse struct {
	*tchttp.BaseResponse
	Response *GetAuthorityIssuerResponseParams `json:"Response"`
}

func (r *GetAuthorityIssuerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthorityIssuerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetConsortiumClusterListRequestParams struct {
	// 联盟id
	ConsortiumId *uint64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`
}

type GetConsortiumClusterListRequest struct {
	*tchttp.BaseRequest
	
	// 联盟id
	ConsortiumId *uint64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`
}

func (r *GetConsortiumClusterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetConsortiumClusterListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConsortiumId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetConsortiumClusterListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetConsortiumClusterListResponseParams struct {
	// 网络列表
	ClusterList []*BcosClusterItem `json:"ClusterList,omitempty" name:"ClusterList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetConsortiumClusterListResponse struct {
	*tchttp.BaseResponse
	Response *GetConsortiumClusterListResponseParams `json:"Response"`
}

func (r *GetConsortiumClusterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetConsortiumClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetConsortiumListRequestParams struct {

}

type GetConsortiumListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetConsortiumListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetConsortiumListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetConsortiumListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetConsortiumListResponseParams struct {
	// 联盟列表
	ConsortiumList []*ConsortiumItem `json:"ConsortiumList,omitempty" name:"ConsortiumList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetConsortiumListResponse struct {
	*tchttp.BaseResponse
	Response *GetConsortiumListResponseParams `json:"Response"`
}

func (r *GetConsortiumListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetConsortiumListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCptInfoRequestParams struct {
	// Cpt索引
	CptIndex *uint64 `json:"CptIndex,omitempty" name:"CptIndex"`
}

type GetCptInfoRequest struct {
	*tchttp.BaseRequest
	
	// Cpt索引
	CptIndex *uint64 `json:"CptIndex,omitempty" name:"CptIndex"`
}

func (r *GetCptInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCptInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CptIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCptInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCptInfoResponseParams struct {
	// CptJsonData的具体信息
	CptJsonData *string `json:"CptJsonData,omitempty" name:"CptJsonData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetCptInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetCptInfoResponseParams `json:"Response"`
}

func (r *GetCptInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCptInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialCptRankRequestParams struct {
	// 开始时间（支持到天 2021-4-23）
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（支持到天 2021-4-23）
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type GetCredentialCptRankRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间（支持到天 2021-4-23）
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（支持到天 2021-4-23）
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetCredentialCptRankRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialCptRankRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCredentialCptRankRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialCptRankResponseParams struct {
	// Rank集合
	RankIssueResult []*CptIssueRank `json:"RankIssueResult,omitempty" name:"RankIssueResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetCredentialCptRankResponse struct {
	*tchttp.BaseResponse
	Response *GetCredentialCptRankResponseParams `json:"Response"`
}

func (r *GetCredentialCptRankResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialCptRankResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialStatusRequestParams struct {
	// 凭证id
	CredentialId *string `json:"CredentialId,omitempty" name:"CredentialId"`
}

type GetCredentialStatusRequest struct {
	*tchttp.BaseRequest
	
	// 凭证id
	CredentialId *string `json:"CredentialId,omitempty" name:"CredentialId"`
}

func (r *GetCredentialStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CredentialId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCredentialStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialStatusResponseParams struct {
	// 凭证状态信息
	CredentialStatus *CredentialStatus `json:"CredentialStatus,omitempty" name:"CredentialStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetCredentialStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetCredentialStatusResponseParams `json:"Response"`
}

func (r *GetCredentialStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidDocumentRequestParams struct {
	// tdid
	Did *string `json:"Did,omitempty" name:"Did"`
}

type GetDidDocumentRequest struct {
	*tchttp.BaseRequest
	
	// tdid
	Did *string `json:"Did,omitempty" name:"Did"`
}

func (r *GetDidDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidDocumentResponseParams struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// DID文档
	Document *string `json:"Document,omitempty" name:"Document"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDidDocumentResponse struct {
	*tchttp.BaseResponse
	Response *GetDidDocumentResponseParams `json:"Response"`
}

func (r *GetDidDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Identity struct {
	// 账户标识符
	AccountIdentifier *string `json:"AccountIdentifier,omitempty" name:"AccountIdentifier"`

	// 链ID
	ChainID *string `json:"ChainID,omitempty" name:"ChainID"`

	// 完整tdid
	Did *string `json:"Did,omitempty" name:"Did"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 群组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type Proof struct {
	// 创建时间
	Created *int64 `json:"Created,omitempty" name:"Created"`

	// 创建着did
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// salt值
	SaltJson *string `json:"SaltJson,omitempty" name:"SaltJson"`

	// 签名
	SignatureValue *string `json:"SignatureValue,omitempty" name:"SignatureValue"`

	// type类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

// Predefined struct for user
type RegisterCptRequestParams struct {
	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// CptJson的具体信息
	CptJson *string `json:"CptJson,omitempty" name:"CptJson"`

	// cptId 不填默认自增
	CptId *uint64 `json:"CptId,omitempty" name:"CptId"`
}

type RegisterCptRequest struct {
	*tchttp.BaseRequest
	
	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// CptJson的具体信息
	CptJson *string `json:"CptJson,omitempty" name:"CptJson"`

	// cptId 不填默认自增
	CptId *uint64 `json:"CptId,omitempty" name:"CptId"`
}

func (r *RegisterCptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterCptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ClusterId")
	delete(f, "CptJson")
	delete(f, "CptId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterCptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterCptResponseParams struct {
	// 凭证模板索引
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 凭证模板id
	CptId *uint64 `json:"CptId,omitempty" name:"CptId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterCptResponse struct {
	*tchttp.BaseResponse
	Response *RegisterCptResponseParams `json:"Response"`
}

func (r *RegisterCptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterCptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCredentialStatusRequestParams struct {
	// 凭证状态
	CredentialStatus *CredentialStatus `json:"CredentialStatus,omitempty" name:"CredentialStatus"`
}

type SetCredentialStatusRequest struct {
	*tchttp.BaseRequest
	
	// 凭证状态
	CredentialStatus *CredentialStatus `json:"CredentialStatus,omitempty" name:"CredentialStatus"`
}

func (r *SetCredentialStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCredentialStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CredentialStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetCredentialStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCredentialStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetCredentialStatusResponse struct {
	*tchttp.BaseResponse
	Response *SetCredentialStatusResponseParams `json:"Response"`
}

func (r *SetCredentialStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCredentialStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransactionArg struct {
	// 凭证did
	InvokerTDid *string `json:"InvokerTDid,omitempty" name:"InvokerTDid"`
}

// Predefined struct for user
type VerifyCredentialRequestParams struct {
	// 参数集合
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitempty" name:"FunctionArg"`
}

type VerifyCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 参数集合
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitempty" name:"FunctionArg"`
}

func (r *VerifyCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionArg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyCredentialResponseParams struct {
	// 是否验证成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 验证返回码
	VerifyCode *uint64 `json:"VerifyCode,omitempty" name:"VerifyCode"`

	// 验证消息
	VerifyMessage *string `json:"VerifyMessage,omitempty" name:"VerifyMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyCredentialResponse struct {
	*tchttp.BaseResponse
	Response *VerifyCredentialResponseParams `json:"Response"`
}

func (r *VerifyCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyFunctionArg struct {
	// CPT ID
	CptId *uint64 `json:"CptId,omitempty" name:"CptId"`

	// issuer did
	Issuer *string `json:"Issuer,omitempty" name:"Issuer"`

	// 过期时间
	ExpirationDate *int64 `json:"ExpirationDate,omitempty" name:"ExpirationDate"`

	// 声明
	ClaimJson *string `json:"ClaimJson,omitempty" name:"ClaimJson"`

	// 颁发时间
	IssuanceDate *int64 `json:"IssuanceDate,omitempty" name:"IssuanceDate"`

	// context值
	Context *string `json:"Context,omitempty" name:"Context"`

	// id值
	Id *string `json:"Id,omitempty" name:"Id"`

	// 签名值
	Proof *Proof `json:"Proof,omitempty" name:"Proof"`

	// type值
	Type []*string `json:"Type,omitempty" name:"Type"`
}