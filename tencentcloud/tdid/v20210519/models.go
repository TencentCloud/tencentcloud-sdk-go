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

type Authority struct {
	// 权威机构ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Did的ID
	DidId *uint64 `json:"DidId,omitempty" name:"DidId"`

	// DID具体信息
	Did *string `json:"Did,omitempty" name:"Did"`

	// 机构名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 权威认证 1:已认证，2:未认证
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// DID服务ID
	DidServiceId *uint64 `json:"DidServiceId,omitempty" name:"DidServiceId"`

	// 应用ID
	ContractAppId *uint64 `json:"ContractAppId,omitempty" name:"ContractAppId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 注册时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisterTime *string `json:"RegisterTime,omitempty" name:"RegisterTime"`

	// 认证时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecognizeTime *string `json:"RecognizeTime,omitempty" name:"RecognizeTime"`

	// 生成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 合约名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 链上标签
	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`
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
type CancelAuthorityIssuerRequestParams struct {
	// did具体信息
	Did *string `json:"Did,omitempty" name:"Did"`
}

type CancelAuthorityIssuerRequest struct {
	*tchttp.BaseRequest
	
	// did具体信息
	Did *string `json:"Did,omitempty" name:"Did"`
}

func (r *CancelAuthorityIssuerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAuthorityIssuerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelAuthorityIssuerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelAuthorityIssuerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelAuthorityIssuerResponse struct {
	*tchttp.BaseResponse
	Response *CancelAuthorityIssuerResponseParams `json:"Response"`
}

func (r *CancelAuthorityIssuerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAuthorityIssuerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// Predefined struct for user
type CheckDidDeployRequestParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type CheckDidDeployRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *CheckDidDeployRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDidDeployRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckDidDeployRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDidDeployResponseParams struct {
	// 服务信息
	Task *Task `json:"Task,omitempty" name:"Task"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckDidDeployResponse struct {
	*tchttp.BaseResponse
	Response *CheckDidDeployResponseParams `json:"Response"`
}

func (r *CheckDidDeployResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDidDeployResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsortiumItem struct {
	// 联盟id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 联盟名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type Contract struct {
	// 应用名
	ApplyName *string `json:"ApplyName,omitempty" name:"ApplyName"`

	// 合约状态 true:已启用 false:未启用
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 合约CNS地址
	Hash *string `json:"Hash,omitempty" name:"Hash"`

	// 合约CNS地址脱敏
	HashShow *string `json:"HashShow,omitempty" name:"HashShow"`

	// 部署机构DID
	WeId *string `json:"WeId,omitempty" name:"WeId"`

	// 部署机构名称
	DeployName *string `json:"DeployName,omitempty" name:"DeployName"`

	// 部署群组
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
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

type CptListData struct {
	// ID信息
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 模版名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 服务ID
	ServiceId *uint64 `json:"ServiceId,omitempty" name:"ServiceId"`

	// 合约应用ID
	ContractAppId *uint64 `json:"ContractAppId,omitempty" name:"ContractAppId"`

	// 凭证模板ID
	CptId *uint64 `json:"CptId,omitempty" name:"CptId"`

	// 模板类型，1: 系统模板，2: 用户模板，3:普通模板
	CptType *uint64 `json:"CptType,omitempty" name:"CptType"`

	// 凭证模版描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 凭证模板Json
	CptJson *string `json:"CptJson,omitempty" name:"CptJson"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建者DID
	CreatorDid *string `json:"CreatorDid,omitempty" name:"CreatorDid"`

	// 应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`
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
type CreateDidServiceRequestParams struct {
	// 联盟名称
	ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

	// 联盟ID
	ConsortiumId *int64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 群组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 机构名称
	AgencyName *string `json:"AgencyName,omitempty" name:"AgencyName"`

	// 应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type CreateDidServiceRequest struct {
	*tchttp.BaseRequest
	
	// 联盟名称
	ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

	// 联盟ID
	ConsortiumId *int64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 群组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 机构名称
	AgencyName *string `json:"AgencyName,omitempty" name:"AgencyName"`

	// 应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

func (r *CreateDidServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDidServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConsortiumName")
	delete(f, "ConsortiumId")
	delete(f, "GroupId")
	delete(f, "AgencyName")
	delete(f, "AppName")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDidServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDidServiceResponseParams struct {
	// 服务信息
	Task *Task `json:"Task,omitempty" name:"Task"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDidServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDidServiceResponseParams `json:"Response"`
}

func (r *CreateDidServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDidServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLabelRequestParams struct {
	// 标签名称
	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`

	// 网络Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

type CreateLabelRequest struct {
	*tchttp.BaseRequest
	
	// 标签名称
	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`

	// 网络Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LabelName")
	delete(f, "ClusterId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLabelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLabelResponse struct {
	*tchttp.BaseResponse
	Response *CreateLabelResponseParams `json:"Response"`
}

func (r *CreateLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLabelResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type DeployByNameRequestParams struct {
	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

type DeployByNameRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeployByNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployByNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationName")
	delete(f, "ClusterId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployByNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployByNameResponseParams struct {
	// 哈希值
	Hash *string `json:"Hash,omitempty" name:"Hash"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeployByNameResponse struct {
	*tchttp.BaseResponse
	Response *DeployByNameResponseParams `json:"Response"`
}

func (r *DeployByNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployByNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DidCluster struct {
	// 链ID
	ChainId *uint64 `json:"ChainId,omitempty" name:"ChainId"`

	// 链名称
	ChainName *string `json:"ChainName,omitempty" name:"ChainName"`

	// 组织数量
	AgencyCount *uint64 `json:"AgencyCount,omitempty" name:"AgencyCount"`

	// 联盟ID
	ConsortiumId *uint64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 网络状态
	ChainStatus *uint64 `json:"ChainStatus,omitempty" name:"ChainStatus"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 联盟名称
	ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

	// 组织ID
	AgencyId *uint64 `json:"AgencyId,omitempty" name:"AgencyId"`

	// 自动续费
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 网络节点总数
	TotalNetworkNode *uint64 `json:"TotalNetworkNode,omitempty" name:"TotalNetworkNode"`

	// 本机构节点数
	TotalCreateNode *uint64 `json:"TotalCreateNode,omitempty" name:"TotalCreateNode"`

	// 总群组数
	TotalGroups *uint64 `json:"TotalGroups,omitempty" name:"TotalGroups"`

	// DID总数
	DidCount *uint64 `json:"DidCount,omitempty" name:"DidCount"`
}

type DidData struct {
	// 服务ID
	ServiceId *uint64 `json:"ServiceId,omitempty" name:"ServiceId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// did号码
	Did *string `json:"Did,omitempty" name:"Did"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 权威机构认证状态 1未注册 2 未认证 3 已认证
	AuthorityState *int64 `json:"AuthorityState,omitempty" name:"AuthorityState"`

	// DID标签名称
	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`

	// DID创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 联盟名称
	AllianceName *string `json:"AllianceName,omitempty" name:"AllianceName"`

	// DID标签id
	LabelId *uint64 `json:"LabelId,omitempty" name:"LabelId"`
}

type DidServiceInfo struct {
	// DID服务索引
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 应用ID
	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`

	// 账号唯一标识
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 联盟id
	ConsortiumId *int64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 联盟名称
	ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 链ID
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 1为盟主，0为非盟主
	RoleType *uint64 `json:"RoleType,omitempty" name:"RoleType"`

	// 机构DID
	AgencyDid *string `json:"AgencyDid,omitempty" name:"AgencyDid"`

	// 机构名称
	CreateOrg *string `json:"CreateOrg,omitempty" name:"CreateOrg"`

	// 端点
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 生成时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 群组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

// Predefined struct for user
type DownCptRequestParams struct {
	// Cpt索引
	CptIndex *uint64 `json:"CptIndex,omitempty" name:"CptIndex"`
}

type DownCptRequest struct {
	*tchttp.BaseRequest
	
	// Cpt索引
	CptIndex *uint64 `json:"CptIndex,omitempty" name:"CptIndex"`
}

func (r *DownCptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownCptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CptIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownCptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownCptResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownCptResponse struct {
	*tchttp.BaseResponse
	Response *DownCptResponseParams `json:"Response"`
}

func (r *DownCptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownCptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableHashRequestParams struct {
	// 合约CNS地址
	Hash *string `json:"Hash,omitempty" name:"Hash"`
}

type EnableHashRequest struct {
	*tchttp.BaseRequest
	
	// 合约CNS地址
	Hash *string `json:"Hash,omitempty" name:"Hash"`
}

func (r *EnableHashRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableHashRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Hash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableHashRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableHashResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableHashResponse struct {
	*tchttp.BaseResponse
	Response *EnableHashResponseParams `json:"Response"`
}

func (r *EnableHashResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableHashResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
type GetAuthoritiesListRequestParams struct {
	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Did信息
	Did *string `json:"Did,omitempty" name:"Did"`

	// 权威认证 1:已认证，2:未认证
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type GetAuthoritiesListRequest struct {
	*tchttp.BaseRequest
	
	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Did信息
	Did *string `json:"Did,omitempty" name:"Did"`

	// 权威认证 1:已认证，2:未认证
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *GetAuthoritiesListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthoritiesListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Did")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAuthoritiesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAuthoritiesListResponseParams struct {
	// 数据集合
	ResultList []*Authority `json:"ResultList,omitempty" name:"ResultList"`

	// 总数
	AllCount *int64 `json:"AllCount,omitempty" name:"AllCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetAuthoritiesListResponse struct {
	*tchttp.BaseResponse
	Response *GetAuthoritiesListResponseParams `json:"Response"`
}

func (r *GetAuthoritiesListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthoritiesListResponse) FromJsonString(s string) error {
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
type GetCptListRequestParams struct {
	// 起始位置
	DisplayStart *uint64 `json:"DisplayStart,omitempty" name:"DisplayStart"`

	// 长度
	DisplayLength *uint64 `json:"DisplayLength,omitempty" name:"DisplayLength"`

	// 模板类型，0: 所有模板，1: 系统模板，2: 用户模板，3:普通模板
	CptType *uint64 `json:"CptType,omitempty" name:"CptType"`
}

type GetCptListRequest struct {
	*tchttp.BaseRequest
	
	// 起始位置
	DisplayStart *uint64 `json:"DisplayStart,omitempty" name:"DisplayStart"`

	// 长度
	DisplayLength *uint64 `json:"DisplayLength,omitempty" name:"DisplayLength"`

	// 模板类型，0: 所有模板，1: 系统模板，2: 用户模板，3:普通模板
	CptType *uint64 `json:"CptType,omitempty" name:"CptType"`
}

func (r *GetCptListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCptListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayStart")
	delete(f, "DisplayLength")
	delete(f, "CptType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCptListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCptListResponseParams struct {
	// cpt数据集合
	CptDataList []*CptListData `json:"CptDataList,omitempty" name:"CptDataList"`

	// 凭证模板总数
	AllCount *uint64 `json:"AllCount,omitempty" name:"AllCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetCptListResponse struct {
	*tchttp.BaseResponse
	Response *GetCptListResponseParams `json:"Response"`
}

func (r *GetCptListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCptListResponse) FromJsonString(s string) error {
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
type GetCredentialIssueRankRequestParams struct {
	// 开始时间（支持到天 2021-4-23）
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（支持到天 2021-4-23）
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type GetCredentialIssueRankRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间（支持到天 2021-4-23）
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（支持到天 2021-4-23）
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetCredentialIssueRankRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialIssueRankRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCredentialIssueRankRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialIssueRankResponseParams struct {
	// Rank集合
	RankIssueResult []*CptIssueRank `json:"RankIssueResult,omitempty" name:"RankIssueResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetCredentialIssueRankResponse struct {
	*tchttp.BaseResponse
	Response *GetCredentialIssueRankResponseParams `json:"Response"`
}

func (r *GetCredentialIssueRankResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialIssueRankResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialIssueTrendRequestParams struct {
	// 开始时间（支持到天 2021-4-23）
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（支持到天 2021-4-23）
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type GetCredentialIssueTrendRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间（支持到天 2021-4-23）
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（支持到天 2021-4-23）
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetCredentialIssueTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialIssueTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCredentialIssueTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialIssueTrendResponseParams struct {
	// Trend集合
	Trend []*Trend `json:"Trend,omitempty" name:"Trend"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetCredentialIssueTrendResponse struct {
	*tchttp.BaseResponse
	Response *GetCredentialIssueTrendResponseParams `json:"Response"`
}

func (r *GetCredentialIssueTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialIssueTrendResponse) FromJsonString(s string) error {
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
type GetDataPanelRequestParams struct {
	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type GetDataPanelRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetDataPanelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataPanelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDataPanelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataPanelResponseParams struct {
	// 区块链网络数量
	BlockNetworkCount *int64 `json:"BlockNetworkCount,omitempty" name:"BlockNetworkCount"`

	// 区块链网络名称
	BlockNetworkName *string `json:"BlockNetworkName,omitempty" name:"BlockNetworkName"`

	// 当前区块高度
	BlockHeight *int64 `json:"BlockHeight,omitempty" name:"BlockHeight"`

	// 区块链网络类型
	BlockNetworkType *int64 `json:"BlockNetworkType,omitempty" name:"BlockNetworkType"`

	// did数量
	DidCount *int64 `json:"DidCount,omitempty" name:"DidCount"`

	// 凭证模版数量
	CptCount *int64 `json:"CptCount,omitempty" name:"CptCount"`

	// 已认证权威机构数量
	CertificatedAuthCount *int64 `json:"CertificatedAuthCount,omitempty" name:"CertificatedAuthCount"`

	// 颁发凭证数量
	IssueCptCount *int64 `json:"IssueCptCount,omitempty" name:"IssueCptCount"`

	// 本周新增DID数量
	NewDidCount *int64 `json:"NewDidCount,omitempty" name:"NewDidCount"`

	// BCOS网络类型数量
	BcosCount *int64 `json:"BcosCount,omitempty" name:"BcosCount"`

	// Fabric网络类型数量
	FabricCount *int64 `json:"FabricCount,omitempty" name:"FabricCount"`

	// 长安链网络类型数量
	ChainMakerCount *int64 `json:"ChainMakerCount,omitempty" name:"ChainMakerCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDataPanelResponse struct {
	*tchttp.BaseResponse
	Response *GetDataPanelResponseParams `json:"Response"`
}

func (r *GetDataPanelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataPanelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeployInfoRequestParams struct {
	// 合约CNS地址
	Hash *string `json:"Hash,omitempty" name:"Hash"`
}

type GetDeployInfoRequest struct {
	*tchttp.BaseRequest
	
	// 合约CNS地址
	Hash *string `json:"Hash,omitempty" name:"Hash"`
}

func (r *GetDeployInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeployInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Hash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeployInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeployInfoResponseParams struct {
	// 合约CNS地址
	Hash *string `json:"Hash,omitempty" name:"Hash"`

	// 合约主群组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署机构DID
	DeployDid *string `json:"DeployDid,omitempty" name:"DeployDid"`

	// TDID SDK版本
	SdkVersion *string `json:"SdkVersion,omitempty" name:"SdkVersion"`

	// TDID 合约版本
	ContractVersion *string `json:"ContractVersion,omitempty" name:"ContractVersion"`

	// 区块链节点版本
	BlockVersion *string `json:"BlockVersion,omitempty" name:"BlockVersion"`

	// 区块链节点IP
	BlockIp *string `json:"BlockIp,omitempty" name:"BlockIp"`

	// DID合约地址
	DidAddress *string `json:"DidAddress,omitempty" name:"DidAddress"`

	// CPT合约地址
	CptAddress *string `json:"CptAddress,omitempty" name:"CptAddress"`

	// Authority Issuer地址
	AuthorityAddress *string `json:"AuthorityAddress,omitempty" name:"AuthorityAddress"`

	// Evidence合约地址
	EvidenceAddress *string `json:"EvidenceAddress,omitempty" name:"EvidenceAddress"`

	// Specific Issuer合约地址
	SpecificAddress *string `json:"SpecificAddress,omitempty" name:"SpecificAddress"`

	// 链ID
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDeployInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetDeployInfoResponseParams `json:"Response"`
}

func (r *GetDeployInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeployInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeployListRequestParams struct {
	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 起始位置
	DisplayStart *uint64 `json:"DisplayStart,omitempty" name:"DisplayStart"`

	// 长度
	DisplayLength *uint64 `json:"DisplayLength,omitempty" name:"DisplayLength"`
}

type GetDeployListRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 起始位置
	DisplayStart *uint64 `json:"DisplayStart,omitempty" name:"DisplayStart"`

	// 长度
	DisplayLength *uint64 `json:"DisplayLength,omitempty" name:"DisplayLength"`
}

func (r *GetDeployListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeployListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "DisplayStart")
	delete(f, "DisplayLength")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeployListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeployListResponseParams struct {
	// 合约总数
	AllCount *uint64 `json:"AllCount,omitempty" name:"AllCount"`

	// 合约部署列表
	Result []*Contract `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDeployListResponse struct {
	*tchttp.BaseResponse
	Response *GetDeployListResponseParams `json:"Response"`
}

func (r *GetDeployListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeployListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidClusterDetailRequestParams struct {
	// DID网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type GetDidClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// DID网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetDidClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidClusterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidClusterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidClusterDetailResponseParams struct {
	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织名称
	ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

	// 区块链组织名称
	ChainAgency *string `json:"ChainAgency,omitempty" name:"ChainAgency"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDidClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *GetDidClusterDetailResponseParams `json:"Response"`
}

func (r *GetDidClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidClusterListRequestParams struct {

}

type GetDidClusterListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetDidClusterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidClusterListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidClusterListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidClusterListResponseParams struct {
	// DID网络列表
	DidClusterList []*DidCluster `json:"DidClusterList,omitempty" name:"DidClusterList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDidClusterListResponse struct {
	*tchttp.BaseResponse
	Response *GetDidClusterListResponseParams `json:"Response"`
}

func (r *GetDidClusterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidDetailRequestParams struct {
	// DID号码的具体信息
	Did *string `json:"Did,omitempty" name:"Did"`
}

type GetDidDetailRequest struct {
	*tchttp.BaseRequest
	
	// DID号码的具体信息
	Did *string `json:"Did,omitempty" name:"Did"`
}

func (r *GetDidDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidDetailResponseParams struct {
	// DID名称
	Did *string `json:"Did,omitempty" name:"Did"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 公钥
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// 权威认证
	AuthorityState *int64 `json:"AuthorityState,omitempty" name:"AuthorityState"`

	// 联盟ID
	ConsortiumId *int64 `json:"ConsortiumId,omitempty" name:"ConsortiumId"`

	// 联盟名称
	ConsortiumName *string `json:"ConsortiumName,omitempty" name:"ConsortiumName"`

	// 群组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// bcos资源ID
	ResChainId *string `json:"ResChainId,omitempty" name:"ResChainId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDidDetailResponse struct {
	*tchttp.BaseResponse
	Response *GetDidDetailResponseParams `json:"Response"`
}

func (r *GetDidDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidDetailResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type GetDidListRequestParams struct {
	// 每页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// Did信息
	Did *string `json:"Did,omitempty" name:"Did"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

type GetDidListRequest struct {
	*tchttp.BaseRequest
	
	// 每页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// Did信息
	Did *string `json:"Did,omitempty" name:"Did"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetDidListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Did")
	delete(f, "ClusterId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidListResponseParams struct {
	// 数据列表
	DataList []*DidData `json:"DataList,omitempty" name:"DataList"`

	// 数据总条数
	AllCount *int64 `json:"AllCount,omitempty" name:"AllCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDidListResponse struct {
	*tchttp.BaseResponse
	Response *GetDidListResponseParams `json:"Response"`
}

func (r *GetDidListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidRegisterTrendRequestParams struct {
	// 开始时间（支持到天 2021-4-23）
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（支持到天 2021-4-23）
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type GetDidRegisterTrendRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间（支持到天 2021-4-23）
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（支持到天 2021-4-23）
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetDidRegisterTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidRegisterTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidRegisterTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidRegisterTrendResponseParams struct {
	// Trend集合
	Trend []*Trend `json:"Trend,omitempty" name:"Trend"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDidRegisterTrendResponse struct {
	*tchttp.BaseResponse
	Response *GetDidRegisterTrendResponseParams `json:"Response"`
}

func (r *GetDidRegisterTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidRegisterTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidServiceDetailRequestParams struct {
	// DID服务ID
	ServiceId *uint64 `json:"ServiceId,omitempty" name:"ServiceId"`
}

type GetDidServiceDetailRequest struct {
	*tchttp.BaseRequest
	
	// DID服务ID
	ServiceId *uint64 `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *GetDidServiceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidServiceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidServiceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidServiceDetailResponseParams struct {
	// did服务信息
	DidService *DidServiceInfo `json:"DidService,omitempty" name:"DidService"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDidServiceDetailResponse struct {
	*tchttp.BaseResponse
	Response *GetDidServiceDetailResponseParams `json:"Response"`
}

func (r *GetDidServiceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidServiceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidServiceListRequestParams struct {
	// 1: 以网络维度输出, 0: 以服务维度输出
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type GetDidServiceListRequest struct {
	*tchttp.BaseRequest
	
	// 1: 以网络维度输出, 0: 以服务维度输出
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *GetDidServiceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidServiceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidServiceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidServiceListResponseParams struct {
	// DID服务列表
	DidServiceList []*DidServiceInfo `json:"DidServiceList,omitempty" name:"DidServiceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDidServiceListResponse struct {
	*tchttp.BaseResponse
	Response *GetDidServiceListResponseParams `json:"Response"`
}

func (r *GetDidServiceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupListRequestParams struct {
	// 0为未部署DID服务的群组，1为已部署DID服务的群组
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type GetGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 0为未部署DID服务的群组，1为已部署DID服务的群组
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupListResponseParams struct {
	// 群组数据集合
	Result []*Group `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *GetGroupListResponseParams `json:"Response"`
}

func (r *GetGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLabelListRequestParams struct {
	// 每页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

type GetLabelListRequest struct {
	*tchttp.BaseRequest
	
	// 每页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetLabelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLabelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "ClusterId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLabelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLabelListResponseParams struct {
	// 数据集合
	Result []*Label `json:"Result,omitempty" name:"Result"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetLabelListResponse struct {
	*tchttp.BaseResponse
	Response *GetLabelListResponseParams `json:"Response"`
}

func (r *GetLabelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLabelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPolicyListRequestParams struct {
	// 起始位置
	DisplayStart *int64 `json:"DisplayStart,omitempty" name:"DisplayStart"`

	// 长度
	DisplayLength *int64 `json:"DisplayLength,omitempty" name:"DisplayLength"`
}

type GetPolicyListRequest struct {
	*tchttp.BaseRequest
	
	// 起始位置
	DisplayStart *int64 `json:"DisplayStart,omitempty" name:"DisplayStart"`

	// 长度
	DisplayLength *int64 `json:"DisplayLength,omitempty" name:"DisplayLength"`
}

func (r *GetPolicyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPolicyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayStart")
	delete(f, "DisplayLength")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPolicyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPolicyListResponseParams struct {
	// 策略Policy管理列表
	PolicyDataList []*Policy `json:"PolicyDataList,omitempty" name:"PolicyDataList"`

	// 总数
	AllCount *int64 `json:"AllCount,omitempty" name:"AllCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPolicyListResponse struct {
	*tchttp.BaseResponse
	Response *GetPolicyListResponseParams `json:"Response"`
}

func (r *GetPolicyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPublicKeyRequestParams struct {
	// did的具体号码
	Did *string `json:"Did,omitempty" name:"Did"`
}

type GetPublicKeyRequest struct {
	*tchttp.BaseRequest
	
	// did的具体号码
	Did *string `json:"Did,omitempty" name:"Did"`
}

func (r *GetPublicKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPublicKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPublicKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPublicKeyResponseParams struct {
	// DID的具体信息
	Did *string `json:"Did,omitempty" name:"Did"`

	// 公钥
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPublicKeyResponse struct {
	*tchttp.BaseResponse
	Response *GetPublicKeyResponseParams `json:"Response"`
}

func (r *GetPublicKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Group struct {
	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 所属机构节点数量
	NodeCountOfAgency *uint64 `json:"NodeCountOfAgency,omitempty" name:"NodeCountOfAgency"`

	// 群组描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参与角色，盟主或非盟主
	RoleType *uint64 `json:"RoleType,omitempty" name:"RoleType"`

	// 链id
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`
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

type Label struct {
	// 标签ID
	LabelId *uint64 `json:"LabelId,omitempty" name:"LabelId"`

	// 标签名称
	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`

	// did数量
	DidCount *int64 `json:"DidCount,omitempty" name:"DidCount"`

	// 创建者did
	Did *string `json:"Did,omitempty" name:"Did"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 群组ID
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

type Policy struct {
	// 披露策略索引
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 披露策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 服务ID
	ServiceId *uint64 `json:"ServiceId,omitempty" name:"ServiceId"`

	// 合约应用ID
	ContractAppId *uint64 `json:"ContractAppId,omitempty" name:"ContractAppId"`

	// 披露策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 凭证模板ID
	CptId *uint64 `json:"CptId,omitempty" name:"CptId"`

	// 策略Json
	PolicyJson *string `json:"PolicyJson,omitempty" name:"PolicyJson"`

	// 生成时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建者DID
	CreatorDid *string `json:"CreatorDid,omitempty" name:"CreatorDid"`

	// 应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 模板索引
	CptIndex *uint64 `json:"CptIndex,omitempty" name:"CptIndex"`
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
type QueryPolicyRequestParams struct {
	// policy索引
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`
}

type QueryPolicyRequest struct {
	*tchttp.BaseRequest
	
	// policy索引
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`
}

func (r *QueryPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryPolicyResponseParams struct {
	// 披露策略索引
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 披露策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 凭证模板ID
	CptId *uint64 `json:"CptId,omitempty" name:"CptId"`

	// 披露策略的具体信息
	PolicyData *string `json:"PolicyData,omitempty" name:"PolicyData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryPolicyResponse struct {
	*tchttp.BaseResponse
	Response *QueryPolicyResponseParams `json:"Response"`
}

func (r *QueryPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeAuthorityIssuerRequestParams struct {
	// did具体信息
	Did *string `json:"Did,omitempty" name:"Did"`
}

type RecognizeAuthorityIssuerRequest struct {
	*tchttp.BaseRequest
	
	// did具体信息
	Did *string `json:"Did,omitempty" name:"Did"`
}

func (r *RecognizeAuthorityIssuerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeAuthorityIssuerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeAuthorityIssuerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeAuthorityIssuerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RecognizeAuthorityIssuerResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeAuthorityIssuerResponseParams `json:"Response"`
}

func (r *RecognizeAuthorityIssuerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeAuthorityIssuerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterClaimPolicyRequestParams struct {
	// Cpt索引
	CptIndex *uint64 `json:"CptIndex,omitempty" name:"CptIndex"`

	// 披露策略
	Policy *string `json:"Policy,omitempty" name:"Policy"`
}

type RegisterClaimPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Cpt索引
	CptIndex *uint64 `json:"CptIndex,omitempty" name:"CptIndex"`

	// 披露策略
	Policy *string `json:"Policy,omitempty" name:"Policy"`
}

func (r *RegisterClaimPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterClaimPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CptIndex")
	delete(f, "Policy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterClaimPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterClaimPolicyResponseParams struct {
	// 披露策略索引
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 披露策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterClaimPolicyResponse struct {
	*tchttp.BaseResponse
	Response *RegisterClaimPolicyResponseParams `json:"Response"`
}

func (r *RegisterClaimPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterClaimPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
type RegisterIssuerRequestParams struct {
	// tdid
	Did *string `json:"Did,omitempty" name:"Did"`

	// 权威机构名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

type RegisterIssuerRequest struct {
	*tchttp.BaseRequest
	
	// tdid
	Did *string `json:"Did,omitempty" name:"Did"`

	// 权威机构名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *RegisterIssuerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterIssuerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterIssuerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterIssuerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterIssuerResponse struct {
	*tchttp.BaseResponse
	Response *RegisterIssuerResponseParams `json:"Response"`
}

func (r *RegisterIssuerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterIssuerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveHashRequestParams struct {
	// 合约CNS地址
	Hash *string `json:"Hash,omitempty" name:"Hash"`
}

type RemoveHashRequest struct {
	*tchttp.BaseRequest
	
	// 合约CNS地址
	Hash *string `json:"Hash,omitempty" name:"Hash"`
}

func (r *RemoveHashRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveHashRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Hash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveHashRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveHashResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveHashResponse struct {
	*tchttp.BaseResponse
	Response *RemoveHashResponseParams `json:"Response"`
}

func (r *RemoveHashResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveHashResponse) FromJsonString(s string) error {
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

type Task struct {
	// 任务ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 应用ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 服务ID
	ServiceId *uint64 `json:"ServiceId,omitempty" name:"ServiceId"`

	// 0: 部署中，1:部署成功，其他失败
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 错误提示
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 生成时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TransactionArg struct {
	// 凭证did
	InvokerTDid *string `json:"InvokerTDid,omitempty" name:"InvokerTDid"`
}

type Trend struct {
	// 时间点
	Time *string `json:"Time,omitempty" name:"Time"`

	// 数量
	Count *int64 `json:"Count,omitempty" name:"Count"`
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