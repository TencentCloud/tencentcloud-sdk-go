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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CRDLArg struct {
	// CPT ID
	CPTId *uint64 `json:"CPTId,omitnil" name:"CPTId"`

	// 签发者 did
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 过期时间
	ExpirationDate *string `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// 声明
	ClaimJson *string `json:"ClaimJson,omitnil" name:"ClaimJson"`

	// 凭证类型
	Type []*string `json:"Type,omitnil" name:"Type"`

	// 多方签名的用户did
	Parties []*string `json:"Parties,omitnil" name:"Parties"`
}

type ChainTransaction struct {
	// 交易哈希
	TransactionHash *string `json:"TransactionHash,omitnil" name:"TransactionHash"`
}

// Predefined struct for user
type CheckNewPurchaseRequestParams struct {

}

type CheckNewPurchaseRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CheckNewPurchaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckNewPurchaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckNewPurchaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckNewPurchaseResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckNewPurchaseResponse struct {
	*tchttp.BaseResponse
	Response *CheckNewPurchaseResponseParams `json:"Response"`
}

func (r *CheckNewPurchaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckNewPurchaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByHostRequestParams struct {
	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 自定义DID文档json属性
	CustomAttribute *string `json:"CustomAttribute,omitnil" name:"CustomAttribute"`
}

type CreateTDidByHostRequest struct {
	*tchttp.BaseRequest
	
	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 自定义DID文档json属性
	CustomAttribute *string `json:"CustomAttribute,omitnil" name:"CustomAttribute"`
}

func (r *CreateTDidByHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "CustomAttribute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByHostResponseParams struct {
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 链上交易信息
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTDidByHostResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidByHostResponseParams `json:"Response"`
}

func (r *CreateTDidByHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPubKeyRequestParams struct {
	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// pem格式的认证公钥
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// 自定义DID初始化属性json字符串
	CustomAttribute *string `json:"CustomAttribute,omitnil" name:"CustomAttribute"`

	// 0:did存在返回错误，1:did存在返回该did，默认:0
	IgnoreExisted *uint64 `json:"IgnoreExisted,omitnil" name:"IgnoreExisted"`
}

type CreateTDidByPubKeyRequest struct {
	*tchttp.BaseRequest
	
	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// pem格式的认证公钥
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// 自定义DID初始化属性json字符串
	CustomAttribute *string `json:"CustomAttribute,omitnil" name:"CustomAttribute"`

	// 0:did存在返回错误，1:did存在返回该did，默认:0
	IgnoreExisted *uint64 `json:"IgnoreExisted,omitnil" name:"IgnoreExisted"`
}

func (r *CreateTDidByPubKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPubKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "PublicKey")
	delete(f, "CustomAttribute")
	delete(f, "IgnoreExisted")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByPubKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPubKeyResponseParams struct {
	// did标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 链上交易信息
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTDidByPubKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidByPubKeyResponseParams `json:"Response"`
}

func (r *CreateTDidByPubKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPubKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CredentialState struct {
	// 凭证唯一id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 凭证状态（0：吊销；1：有效）
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 凭证颁发者Did
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// VC摘要，对应凭证Proof的vcDigest字段
	VCDigest *string `json:"VCDigest,omitnil" name:"VCDigest"`

	// 交易摘要，对应凭证Proof的txDigest字段 
	TXDigest *string `json:"TXDigest,omitnil" name:"TXDigest"`

	// 颁布凭证的UTC时间戳
	IssueTime *uint64 `json:"IssueTime,omitnil" name:"IssueTime"`

	// 凭证过期的UTC时间戳
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 凭证模板id
	CPTId *uint64 `json:"CPTId,omitnil" name:"CPTId"`

	// 凭证签名
	Signature *string `json:"Signature,omitnil" name:"Signature"`

	// 元数据摘要
	MetaDigest *string `json:"MetaDigest,omitnil" name:"MetaDigest"`
}

// Predefined struct for user
type DeactivateTDidRequestParams struct {
	// DID标识符
	Did *string `json:"Did,omitnil" name:"Did"`

	// 设置DID禁用状态的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`

	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 是否禁用
	Deactivated *string `json:"Deactivated,omitnil" name:"Deactivated"`
}

type DeactivateTDidRequest struct {
	*tchttp.BaseRequest
	
	// DID标识符
	Did *string `json:"Did,omitnil" name:"Did"`

	// 设置DID禁用状态的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`

	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 是否禁用
	Deactivated *string `json:"Deactivated,omitnil" name:"Deactivated"`
}

func (r *DeactivateTDidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeactivateTDidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "OperateCredential")
	delete(f, "DAPId")
	delete(f, "Deactivated")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeactivateTDidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeactivateTDidResponseParams struct {
	// 上链交易信息
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeactivateTDidResponse struct {
	*tchttp.BaseResponse
	Response *DeactivateTDidResponseParams `json:"Response"`
}

func (r *DeactivateTDidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeactivateTDidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialStateRequestParams struct {
	// 凭证唯一Id
	CredentialId *string `json:"CredentialId,omitnil" name:"CredentialId"`

	// 用户应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type GetCredentialStateRequest struct {
	*tchttp.BaseRequest
	
	// 凭证唯一Id
	CredentialId *string `json:"CredentialId,omitnil" name:"CredentialId"`

	// 用户应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *GetCredentialStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CredentialId")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCredentialStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialStateResponseParams struct {
	// 凭证状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CredentialState *CredentialState `json:"CredentialState,omitnil" name:"CredentialState"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetCredentialStateResponse struct {
	*tchttp.BaseResponse
	Response *GetCredentialStateResponseParams `json:"Response"`
}

func (r *GetCredentialStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidDocumentRequestParams struct {
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type GetTDidDocumentRequest struct {
	*tchttp.BaseRequest
	
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *GetTDidDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTDidDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidDocumentResponseParams struct {
	// DID文档内容
	Document *string `json:"Document,omitnil" name:"Document"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetTDidDocumentResponse struct {
	*tchttp.BaseResponse
	Response *GetTDidDocumentResponseParams `json:"Response"`
}

func (r *GetTDidDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IssueCredentialRequestParams struct {
	// 参数集合，详见示例
	CRDLArg *CRDLArg `json:"CRDLArg,omitnil" name:"CRDLArg"`

	// 是否未签名
	UnSigned *bool `json:"UnSigned,omitnil" name:"UnSigned"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type IssueCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 参数集合，详见示例
	CRDLArg *CRDLArg `json:"CRDLArg,omitnil" name:"CRDLArg"`

	// 是否未签名
	UnSigned *bool `json:"UnSigned,omitnil" name:"UnSigned"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *IssueCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IssueCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CRDLArg")
	delete(f, "UnSigned")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IssueCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IssueCredentialResponseParams struct {
	// 可验证凭证内容
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IssueCredentialResponse struct {
	*tchttp.BaseResponse
	Response *IssueCredentialResponseParams `json:"Response"`
}

func (r *IssueCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IssueCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCredentialStateRequestParams struct {
	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 更新VC状态的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`
}

type UpdateCredentialStateRequest struct {
	*tchttp.BaseRequest
	
	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 更新VC状态的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`
}

func (r *UpdateCredentialStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCredentialStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "OperateCredential")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCredentialStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCredentialStateResponseParams struct {
	// 更新是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateCredentialStateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCredentialStateResponseParams `json:"Response"`
}

func (r *UpdateCredentialStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCredentialStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyCredentialsRequestParams struct {
	// 0:仅验证签名，1:验证签名和链上状态，2, 仅验证链上状态，默认为0, 3.验证DID和凭证状态以及签名，4. 验证历史凭证有效性
	VerifyType *uint64 `json:"VerifyType,omitnil" name:"VerifyType"`

	// 凭证内容
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type VerifyCredentialsRequest struct {
	*tchttp.BaseRequest
	
	// 0:仅验证签名，1:验证签名和链上状态，2, 仅验证链上状态，默认为0, 3.验证DID和凭证状态以及签名，4. 验证历史凭证有效性
	VerifyType *uint64 `json:"VerifyType,omitnil" name:"VerifyType"`

	// 凭证内容
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *VerifyCredentialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VerifyType")
	delete(f, "CredentialData")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyCredentialsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyCredentialsResponseParams struct {
	// 是否验证成功
	Result *bool `json:"Result,omitnil" name:"Result"`

	// 验证返回码
	VerifyCode *uint64 `json:"VerifyCode,omitnil" name:"VerifyCode"`

	// 验证结果信息
	VerifyMessage *string `json:"VerifyMessage,omitnil" name:"VerifyMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VerifyCredentialsResponse struct {
	*tchttp.BaseResponse
	Response *VerifyCredentialsResponseParams `json:"Response"`
}

func (r *VerifyCredentialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}