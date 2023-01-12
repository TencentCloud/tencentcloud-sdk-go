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

package v20180416

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ApplyChainMakerBatchUserCertRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 证书标识和证书请求文件，可参考TBaaS证书生成相关文档生成证书请求文件
	SignUserCsrList []*SignCertCsr `json:"SignUserCsrList,omitempty" name:"SignUserCsrList"`
}

type ApplyChainMakerBatchUserCertRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 证书标识和证书请求文件，可参考TBaaS证书生成相关文档生成证书请求文件
	SignUserCsrList []*SignCertCsr `json:"SignUserCsrList,omitempty" name:"SignUserCsrList"`
}

func (r *ApplyChainMakerBatchUserCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyChainMakerBatchUserCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SignUserCsrList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyChainMakerBatchUserCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyChainMakerBatchUserCertResponseParams struct {
	// 成功生成的用户证书的base64编码字符串列表，与SignUserCsrList一一对应
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignUserCrtList []*string `json:"SignUserCrtList,omitempty" name:"SignUserCrtList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyChainMakerBatchUserCertResponse struct {
	*tchttp.BaseResponse
	Response *ApplyChainMakerBatchUserCertResponseParams `json:"Response"`
}

func (r *ApplyChainMakerBatchUserCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyChainMakerBatchUserCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyUserCertRequestParams struct {
	// 模块名，固定字段：cert_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：cert_apply_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 申请证书的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 用户证书标识，用于标识用户证书，要求由纯小写字母组成，长度小于10
	UserIdentity *string `json:"UserIdentity,omitempty" name:"UserIdentity"`

	// 证书申请实体，使用腾讯云账号实名认证的名称
	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`

	// 证件号码。如果腾讯云账号对应的实名认证类型为企业认证，填入“0”；如果腾讯云账号对应的实名认证类型为个人认证，填入个人身份证号码
	IdentityNum *string `json:"IdentityNum,omitempty" name:"IdentityNum"`

	// csr p10证书文件。需要用户根据文档生成证书的CSR文件
	CsrData *string `json:"CsrData,omitempty" name:"CsrData"`

	// 证书备注信息
	Notes *string `json:"Notes,omitempty" name:"Notes"`
}

type ApplyUserCertRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：cert_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：cert_apply_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 申请证书的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 用户证书标识，用于标识用户证书，要求由纯小写字母组成，长度小于10
	UserIdentity *string `json:"UserIdentity,omitempty" name:"UserIdentity"`

	// 证书申请实体，使用腾讯云账号实名认证的名称
	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`

	// 证件号码。如果腾讯云账号对应的实名认证类型为企业认证，填入“0”；如果腾讯云账号对应的实名认证类型为个人认证，填入个人身份证号码
	IdentityNum *string `json:"IdentityNum,omitempty" name:"IdentityNum"`

	// csr p10证书文件。需要用户根据文档生成证书的CSR文件
	CsrData *string `json:"CsrData,omitempty" name:"CsrData"`

	// 证书备注信息
	Notes *string `json:"Notes,omitempty" name:"Notes"`
}

func (r *ApplyUserCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyUserCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	delete(f, "UserIdentity")
	delete(f, "Applicant")
	delete(f, "IdentityNum")
	delete(f, "CsrData")
	delete(f, "Notes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyUserCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyUserCertResponseParams struct {
	// 证书ID
	CertId *uint64 `json:"CertId,omitempty" name:"CertId"`

	// 证书DN
	CertDn *string `json:"CertDn,omitempty" name:"CertDn"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyUserCertResponse struct {
	*tchttp.BaseResponse
	Response *ApplyUserCertResponseParams `json:"Response"`
}

func (r *ApplyUserCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyUserCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BcosBlockObj struct {
	// 区块哈希
	BlockHash *string `json:"BlockHash,omitempty" name:"BlockHash"`

	// 区块高度
	BlockNumber *int64 `json:"BlockNumber,omitempty" name:"BlockNumber"`

	// 区块时间戳
	BlockTimestamp *string `json:"BlockTimestamp,omitempty" name:"BlockTimestamp"`

	// 打包节点ID
	Sealer *string `json:"Sealer,omitempty" name:"Sealer"`

	// 打包节点索引
	SealerIndex *int64 `json:"SealerIndex,omitempty" name:"SealerIndex"`

	// 记录保存时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 交易数量
	TransCount *int64 `json:"TransCount,omitempty" name:"TransCount"`

	// 记录修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type BcosTransInfo struct {
	// 所属区块高度
	BlockNumber *int64 `json:"BlockNumber,omitempty" name:"BlockNumber"`

	// 区块时间戳
	BlockTimestamp *string `json:"BlockTimestamp,omitempty" name:"BlockTimestamp"`

	// 交易哈希
	TransHash *string `json:"TransHash,omitempty" name:"TransHash"`

	// 交易发起者
	TransFrom *string `json:"TransFrom,omitempty" name:"TransFrom"`

	// 交易接收者
	TransTo *string `json:"TransTo,omitempty" name:"TransTo"`

	// 落库时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type Block struct {
	// 区块编号
	BlockNum *uint64 `json:"BlockNum,omitempty" name:"BlockNum"`

	// 区块数据Hash数值
	DataHash *string `json:"DataHash,omitempty" name:"DataHash"`

	// 区块ID，与区块编号一致
	BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

	// 前一个区块Hash
	PreHash *string `json:"PreHash,omitempty" name:"PreHash"`

	// 区块内的交易数量
	TxCount *uint64 `json:"TxCount,omitempty" name:"TxCount"`
}

type ChainMakerContractResult struct {
	// 交易结果码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 交易结果码含义
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeMessage *string `json:"CodeMessage,omitempty" name:"CodeMessage"`

	// 交易ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TxId *string `json:"TxId,omitempty" name:"TxId"`

	// Gas使用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GasUsed *int64 `json:"GasUsed,omitempty" name:"GasUsed"`

	// 合约返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 合约函数返回，base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`
}

type ChainMakerTransactionResult struct {
	// 交易结果码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 交易结果码含义
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeMessage *string `json:"CodeMessage,omitempty" name:"CodeMessage"`

	// 交易ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TxId *string `json:"TxId,omitempty" name:"TxId"`

	// Gas使用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GasUsed *int64 `json:"GasUsed,omitempty" name:"GasUsed"`

	// 区块高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockHeight *int64 `json:"BlockHeight,omitempty" name:"BlockHeight"`

	// 合约执行结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractEvent *string `json:"ContractEvent,omitempty" name:"ContractEvent"`

	// 合约返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 交易时间，单位是秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type ChannelDetailForUser struct {
	// 通道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 当前组织加入通道的节点列表
	PeerList []*PeerDetailForUser `json:"PeerList,omitempty" name:"PeerList"`
}

type ClusterDetailForUser struct {
	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织列表
	GroupList []*GroupDetailForUser `json:"GroupList,omitempty" name:"GroupList"`

	// 网络名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

// Predefined struct for user
type CreateChaincodeAndInstallForUserRequestParams struct {
	// 模块名，本接口取值：chaincode_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：chaincode_create_and_install_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 合约安装节点名称，可以在通道详情中获取该通道上的节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 智能合约名称，格式说明：以小写字母开头，由2-12位数字或小写字母组成
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 智能合约版本，格式说明：由1-12位数字、小写字母、特殊符号(“.”)组成，如v1.0
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 智能合约代码文件类型，支持类型：
	// 1. "go"：.go合约文件
	// 2. "gozip"：go合约工程zip包，要求压缩目录为代码根目录
	// 3. "javazip"：java合约工程zip包，要求压缩目录为代码根目录
	// 4. "nodezip"：nodejs合约工程zip包，要求压缩目录为代码根目录
	ChaincodeFileType *string `json:"ChaincodeFileType,omitempty" name:"ChaincodeFileType"`

	// 合约内容，合约文件或压缩包内容的base64编码，大小要求小于等于5M
	Chaincode *string `json:"Chaincode,omitempty" name:"Chaincode"`
}

type CreateChaincodeAndInstallForUserRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：chaincode_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：chaincode_create_and_install_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 合约安装节点名称，可以在通道详情中获取该通道上的节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 智能合约名称，格式说明：以小写字母开头，由2-12位数字或小写字母组成
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 智能合约版本，格式说明：由1-12位数字、小写字母、特殊符号(“.”)组成，如v1.0
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 智能合约代码文件类型，支持类型：
	// 1. "go"：.go合约文件
	// 2. "gozip"：go合约工程zip包，要求压缩目录为代码根目录
	// 3. "javazip"：java合约工程zip包，要求压缩目录为代码根目录
	// 4. "nodezip"：nodejs合约工程zip包，要求压缩目录为代码根目录
	ChaincodeFileType *string `json:"ChaincodeFileType,omitempty" name:"ChaincodeFileType"`

	// 合约内容，合约文件或压缩包内容的base64编码，大小要求小于等于5M
	Chaincode *string `json:"Chaincode,omitempty" name:"Chaincode"`
}

func (r *CreateChaincodeAndInstallForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChaincodeAndInstallForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	delete(f, "PeerName")
	delete(f, "ChaincodeName")
	delete(f, "ChaincodeVersion")
	delete(f, "ChaincodeFileType")
	delete(f, "Chaincode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChaincodeAndInstallForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChaincodeAndInstallForUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateChaincodeAndInstallForUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateChaincodeAndInstallForUserResponseParams `json:"Response"`
}

func (r *CreateChaincodeAndInstallForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChaincodeAndInstallForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployDynamicBcosContractRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组编号，可在群组列表中获取
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 合约编译后的ABI，可在合约详情获取
	AbiInfo *string `json:"AbiInfo,omitempty" name:"AbiInfo"`

	// 合约编译得到的字节码，hex编码，可在合约详情获取
	ByteCodeBin *string `json:"ByteCodeBin,omitempty" name:"ByteCodeBin"`

	// 签名用户编号，可在私钥管理页面获取
	SignUserId *string `json:"SignUserId,omitempty" name:"SignUserId"`

	// 构造函数入参，Json数组，多个参数以逗号分隔（参数为数组时同理），如：["str1",["arr1","arr2"]]
	ConstructorParams *string `json:"ConstructorParams,omitempty" name:"ConstructorParams"`
}

type DeployDynamicBcosContractRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组编号，可在群组列表中获取
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 合约编译后的ABI，可在合约详情获取
	AbiInfo *string `json:"AbiInfo,omitempty" name:"AbiInfo"`

	// 合约编译得到的字节码，hex编码，可在合约详情获取
	ByteCodeBin *string `json:"ByteCodeBin,omitempty" name:"ByteCodeBin"`

	// 签名用户编号，可在私钥管理页面获取
	SignUserId *string `json:"SignUserId,omitempty" name:"SignUserId"`

	// 构造函数入参，Json数组，多个参数以逗号分隔（参数为数组时同理），如：["str1",["arr1","arr2"]]
	ConstructorParams *string `json:"ConstructorParams,omitempty" name:"ConstructorParams"`
}

func (r *DeployDynamicBcosContractRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployDynamicBcosContractRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "AbiInfo")
	delete(f, "ByteCodeBin")
	delete(f, "SignUserId")
	delete(f, "ConstructorParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployDynamicBcosContractRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployDynamicBcosContractResponseParams struct {
	// 部署成功返回的合约地址
	ContractAddress *string `json:"ContractAddress,omitempty" name:"ContractAddress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeployDynamicBcosContractResponse struct {
	*tchttp.BaseResponse
	Response *DeployDynamicBcosContractResponseParams `json:"Response"`
}

func (r *DeployDynamicBcosContractResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployDynamicBcosContractResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadUserCertRequestParams struct {
	// 模块名，固定字段：cert_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：cert_download_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 证书ID，可以在证书详情页面获取
	CertId *uint64 `json:"CertId,omitempty" name:"CertId"`

	// 证书DN，可以在证书详情页面获取
	CertDn *string `json:"CertDn,omitempty" name:"CertDn"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 下载证书的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type DownloadUserCertRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：cert_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：cert_download_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 证书ID，可以在证书详情页面获取
	CertId *uint64 `json:"CertId,omitempty" name:"CertId"`

	// 证书DN，可以在证书详情页面获取
	CertDn *string `json:"CertDn,omitempty" name:"CertDn"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 下载证书的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

func (r *DownloadUserCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadUserCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "CertId")
	delete(f, "CertDn")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadUserCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadUserCertResponseParams struct {
	// 证书名称
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 证书内容
	CertCtx *string `json:"CertCtx,omitempty" name:"CertCtx"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadUserCertResponse struct {
	*tchttp.BaseResponse
	Response *DownloadUserCertResponseParams `json:"Response"`
}

func (r *DownloadUserCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadUserCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EndorserGroup struct {
	// 背书组织名称
	EndorserGroupName *string `json:"EndorserGroupName,omitempty" name:"EndorserGroupName"`

	// 背书节点列表
	EndorserPeerList []*string `json:"EndorserPeerList,omitempty" name:"EndorserPeerList"`
}

// Predefined struct for user
type GetBcosBlockByNumberRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组编号，可在群组列表中获取
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 区块高度，可以从InvokeBcosTrans接口的返回值中解析获取
	BlockNumber *int64 `json:"BlockNumber,omitempty" name:"BlockNumber"`
}

type GetBcosBlockByNumberRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组编号，可在群组列表中获取
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 区块高度，可以从InvokeBcosTrans接口的返回值中解析获取
	BlockNumber *int64 `json:"BlockNumber,omitempty" name:"BlockNumber"`
}

func (r *GetBcosBlockByNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBcosBlockByNumberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "BlockNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetBcosBlockByNumberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBcosBlockByNumberResponseParams struct {
	// 返回区块json字符串
	BlockJson *string `json:"BlockJson,omitempty" name:"BlockJson"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetBcosBlockByNumberResponse struct {
	*tchttp.BaseResponse
	Response *GetBcosBlockByNumberResponseParams `json:"Response"`
}

func (r *GetBcosBlockByNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBcosBlockByNumberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBcosBlockListRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组编号，可在群组列表中获取
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 当前页数，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页记录数，默认为10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 区块高度，可以从InvokeBcosTrans接口的返回值中解析获取
	BlockNumber *int64 `json:"BlockNumber,omitempty" name:"BlockNumber"`

	// 区块哈希，可以从InvokeBcosTrans接口的返回值中解析获取
	BlockHash *string `json:"BlockHash,omitempty" name:"BlockHash"`
}

type GetBcosBlockListRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组编号，可在群组列表中获取
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 当前页数，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页记录数，默认为10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 区块高度，可以从InvokeBcosTrans接口的返回值中解析获取
	BlockNumber *int64 `json:"BlockNumber,omitempty" name:"BlockNumber"`

	// 区块哈希，可以从InvokeBcosTrans接口的返回值中解析获取
	BlockHash *string `json:"BlockHash,omitempty" name:"BlockHash"`
}

func (r *GetBcosBlockListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBcosBlockListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "BlockNumber")
	delete(f, "BlockHash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetBcosBlockListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBcosBlockListResponseParams struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回数据列表
	List []*BcosBlockObj `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetBcosBlockListResponse struct {
	*tchttp.BaseResponse
	Response *GetBcosBlockListResponseParams `json:"Response"`
}

func (r *GetBcosBlockListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBcosBlockListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBcosTransByHashRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组编号，可在群组列表中获取
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 交易哈希值，可以从InvokeBcosTrans接口的返回值中解析获取
	TransHash *string `json:"TransHash,omitempty" name:"TransHash"`
}

type GetBcosTransByHashRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组编号，可在群组列表中获取
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 交易哈希值，可以从InvokeBcosTrans接口的返回值中解析获取
	TransHash *string `json:"TransHash,omitempty" name:"TransHash"`
}

func (r *GetBcosTransByHashRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBcosTransByHashRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "TransHash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetBcosTransByHashRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBcosTransByHashResponseParams struct {
	// 交易信息json字符串
	TransactionJson *string `json:"TransactionJson,omitempty" name:"TransactionJson"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetBcosTransByHashResponse struct {
	*tchttp.BaseResponse
	Response *GetBcosTransByHashResponseParams `json:"Response"`
}

func (r *GetBcosTransByHashResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBcosTransByHashResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBcosTransListRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组编号，可在群组列表中获取
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 当前页数，默认是1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页记录数，默认为10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 区块高度，可以从InvokeBcosTrans接口的返回值中解析获取
	BlockNumber *int64 `json:"BlockNumber,omitempty" name:"BlockNumber"`

	// 交易哈希，可以从InvokeBcosTrans接口的返回值中解析获取
	TransHash *string `json:"TransHash,omitempty" name:"TransHash"`
}

type GetBcosTransListRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组编号，可在群组列表中获取
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 当前页数，默认是1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页记录数，默认为10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 区块高度，可以从InvokeBcosTrans接口的返回值中解析获取
	BlockNumber *int64 `json:"BlockNumber,omitempty" name:"BlockNumber"`

	// 交易哈希，可以从InvokeBcosTrans接口的返回值中解析获取
	TransHash *string `json:"TransHash,omitempty" name:"TransHash"`
}

func (r *GetBcosTransListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBcosTransListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "BlockNumber")
	delete(f, "TransHash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetBcosTransListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBcosTransListResponseParams struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回数据列表
	List []*BcosTransInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetBcosTransListResponse struct {
	*tchttp.BaseResponse
	Response *GetBcosTransListResponseParams `json:"Response"`
}

func (r *GetBcosTransListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBcosTransListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBlockListRequestParams struct {
	// 模块名称，固定字段：block
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称，固定字段：block_list
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 通道ID，固定字段：0
	ChannelId *uint64 `json:"ChannelId,omitempty" name:"ChannelId"`

	// 组织ID，固定字段：0
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 需要查询的通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要获取的起始交易偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要获取的交易数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type GetBlockListRequest struct {
	*tchttp.BaseRequest
	
	// 模块名称，固定字段：block
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称，固定字段：block_list
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 通道ID，固定字段：0
	ChannelId *uint64 `json:"ChannelId,omitempty" name:"ChannelId"`

	// 组织ID，固定字段：0
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 需要查询的通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要获取的起始交易偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要获取的交易数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetBlockListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBlockListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ChannelId")
	delete(f, "GroupId")
	delete(f, "ChannelName")
	delete(f, "GroupName")
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetBlockListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBlockListResponseParams struct {
	// 区块数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 区块列表
	BlockList []*Block `json:"BlockList,omitempty" name:"BlockList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetBlockListResponse struct {
	*tchttp.BaseResponse
	Response *GetBlockListResponseParams `json:"Response"`
}

func (r *GetBlockListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBlockListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBlockTransactionListForUserRequestParams struct {
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：block_transaction_list_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 参与交易的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 区块ID，通过GetInvokeTx接口可以获取交易所在的区块ID
	BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

	// 查询的交易列表起始偏移地址
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的交易列表数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type GetBlockTransactionListForUserRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：block_transaction_list_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 参与交易的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 区块ID，通过GetInvokeTx接口可以获取交易所在的区块ID
	BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

	// 查询的交易列表起始偏移地址
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的交易列表数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetBlockTransactionListForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBlockTransactionListForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	delete(f, "ChannelName")
	delete(f, "BlockId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetBlockTransactionListForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBlockTransactionListForUserResponseParams struct {
	// 交易总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 交易列表
	TransactionList []*TransactionItem `json:"TransactionList,omitempty" name:"TransactionList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetBlockTransactionListForUserResponse struct {
	*tchttp.BaseResponse
	Response *GetBlockTransactionListForUserResponseParams `json:"Response"`
}

func (r *GetBlockTransactionListForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBlockTransactionListForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetChaincodeCompileLogForUserRequestParams struct {
	// 模块名，本接口取值：chaincode_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：chaincode_compile_log_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 调用合约的组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 业务所属智能合约名称
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属智能合约版本
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 合约安装节点名称，可以在通道详情中获取该通道上的节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 返回数据项数，本接口默认取值：10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回数据起始偏移，本接口默认取值：0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type GetChaincodeCompileLogForUserRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：chaincode_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：chaincode_compile_log_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 调用合约的组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 业务所属智能合约名称
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属智能合约版本
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 合约安装节点名称，可以在通道详情中获取该通道上的节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 返回数据项数，本接口默认取值：10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回数据起始偏移，本接口默认取值：0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetChaincodeCompileLogForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetChaincodeCompileLogForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	delete(f, "ChaincodeName")
	delete(f, "ChaincodeVersion")
	delete(f, "PeerName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetChaincodeCompileLogForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetChaincodeCompileLogForUserResponseParams struct {
	// 日志总行数，上限2000条日志
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 日志列表
	CompileLogList []*LogDetailForUser `json:"CompileLogList,omitempty" name:"CompileLogList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetChaincodeCompileLogForUserResponse struct {
	*tchttp.BaseResponse
	Response *GetChaincodeCompileLogForUserResponseParams `json:"Response"`
}

func (r *GetChaincodeCompileLogForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetChaincodeCompileLogForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetChaincodeInitializeResultForUserRequestParams struct {
	// 模块名，本接口取值：chaincode_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：chaincode_init_result_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 调用合约的组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 业务所属通道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 业务所属合约名称
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属智能合约版本
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 实例化任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type GetChaincodeInitializeResultForUserRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：chaincode_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：chaincode_init_result_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 调用合约的组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 业务所属通道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 业务所属合约名称
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属智能合约版本
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 实例化任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *GetChaincodeInitializeResultForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetChaincodeInitializeResultForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	delete(f, "ChannelName")
	delete(f, "ChaincodeName")
	delete(f, "ChaincodeVersion")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetChaincodeInitializeResultForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetChaincodeInitializeResultForUserResponseParams struct {
	// 实例化结果：0，实例化中；1，实例化成功；2，实例化失败
	InitResult *uint64 `json:"InitResult,omitempty" name:"InitResult"`

	// 实例化信息
	InitMessage *string `json:"InitMessage,omitempty" name:"InitMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetChaincodeInitializeResultForUserResponse struct {
	*tchttp.BaseResponse
	Response *GetChaincodeInitializeResultForUserResponseParams `json:"Response"`
}

func (r *GetChaincodeInitializeResultForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetChaincodeInitializeResultForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetChaincodeLogForUserRequestParams struct {
	// 模块名，本接口取值：chaincode_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：chaincode_log_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 调用合约的组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 业务所属智能合约名称
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属智能合约版本
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 合约安装节点名称，可以在通道详情中获取该通道上的节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 日志开始时间，如"2020-11-24 19:49:25"
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 返回日志行数的最大值，系统设定该参数最大为1000，且一行日志的最大字节数是500，即最大返回50万个字节数的日志数据
	RowNum *int64 `json:"RowNum,omitempty" name:"RowNum"`
}

type GetChaincodeLogForUserRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：chaincode_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：chaincode_log_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 调用合约的组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 业务所属智能合约名称
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属智能合约版本
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 合约安装节点名称，可以在通道详情中获取该通道上的节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 日志开始时间，如"2020-11-24 19:49:25"
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 返回日志行数的最大值，系统设定该参数最大为1000，且一行日志的最大字节数是500，即最大返回50万个字节数的日志数据
	RowNum *int64 `json:"RowNum,omitempty" name:"RowNum"`
}

func (r *GetChaincodeLogForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetChaincodeLogForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	delete(f, "ChaincodeName")
	delete(f, "ChaincodeVersion")
	delete(f, "PeerName")
	delete(f, "BeginTime")
	delete(f, "RowNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetChaincodeLogForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetChaincodeLogForUserResponseParams struct {
	// 返回日志总行数，不会超过入参的RowNum
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 日志列表
	ChaincodeLogList []*LogDetailForUser `json:"ChaincodeLogList,omitempty" name:"ChaincodeLogList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetChaincodeLogForUserResponse struct {
	*tchttp.BaseResponse
	Response *GetChaincodeLogForUserResponseParams `json:"Response"`
}

func (r *GetChaincodeLogForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetChaincodeLogForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetChannelListForUserRequestParams struct {
	// 模块名，本接口取值：channel_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：channel_list_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 返回数据项数，本接口默认取值：10，上限取值：20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回数据起始偏移，本接口默认取值：0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type GetChannelListForUserRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：channel_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：channel_list_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 返回数据项数，本接口默认取值：10，上限取值：20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回数据起始偏移，本接口默认取值：0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetChannelListForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetChannelListForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetChannelListForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetChannelListForUserResponseParams struct {
	// 通道总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 通道列表
	ChannelList []*ChannelDetailForUser `json:"ChannelList,omitempty" name:"ChannelList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetChannelListForUserResponse struct {
	*tchttp.BaseResponse
	Response *GetChannelListForUserResponseParams `json:"Response"`
}

func (r *GetChannelListForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetChannelListForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetClusterListForUserRequestParams struct {
	// 模块名，本接口取值：cluster_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：cluster_list_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 返回数据项数，本接口默认取值：10，上限取值：20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回数据起始偏移，本接口默认取值：0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type GetClusterListForUserRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：cluster_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：cluster_list_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 返回数据项数，本接口默认取值：10，上限取值：20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回数据起始偏移，本接口默认取值：0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetClusterListForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetClusterListForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetClusterListForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetClusterListForUserResponseParams struct {
	// 网络总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 网络列表
	ClusterList []*ClusterDetailForUser `json:"ClusterList,omitempty" name:"ClusterList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetClusterListForUserResponse struct {
	*tchttp.BaseResponse
	Response *GetClusterListForUserResponseParams `json:"Response"`
}

func (r *GetClusterListForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetClusterListForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetClusterSummaryRequestParams struct {
	// 模块名称，固定字段：cluster_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称，固定字段：cluster_summary
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织ID，固定字段：0
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type GetClusterSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 模块名称，固定字段：cluster_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称，固定字段：cluster_summary
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织ID，固定字段：0
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

func (r *GetClusterSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetClusterSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "GroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetClusterSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetClusterSummaryResponseParams struct {
	// 网络通道总数量
	TotalChannelCount *uint64 `json:"TotalChannelCount,omitempty" name:"TotalChannelCount"`

	// 当前组织创建的通道数量
	MyChannelCount *uint64 `json:"MyChannelCount,omitempty" name:"MyChannelCount"`

	// 当前组织加入的通道数量
	JoinChannelCount *uint64 `json:"JoinChannelCount,omitempty" name:"JoinChannelCount"`

	// 网络节点总数量
	TotalPeerCount *uint64 `json:"TotalPeerCount,omitempty" name:"TotalPeerCount"`

	// 当前组织创建的节点数量
	MyPeerCount *uint64 `json:"MyPeerCount,omitempty" name:"MyPeerCount"`

	// 其他组织创建的节点数量
	OrderCount *uint64 `json:"OrderCount,omitempty" name:"OrderCount"`

	// 网络组织总数量
	TotalGroupCount *uint64 `json:"TotalGroupCount,omitempty" name:"TotalGroupCount"`

	// 当前组织创建的组织数量
	MyGroupCount *uint64 `json:"MyGroupCount,omitempty" name:"MyGroupCount"`

	// 网络智能合约总数量
	TotalChaincodeCount *uint64 `json:"TotalChaincodeCount,omitempty" name:"TotalChaincodeCount"`

	// 最近7天发起的智能合约数量
	RecentChaincodeCount *uint64 `json:"RecentChaincodeCount,omitempty" name:"RecentChaincodeCount"`

	// 当前组织发起的智能合约数量
	MyChaincodeCount *uint64 `json:"MyChaincodeCount,omitempty" name:"MyChaincodeCount"`

	// 当前组织的证书总数量
	TotalCertCount *uint64 `json:"TotalCertCount,omitempty" name:"TotalCertCount"`

	// 颁发给当前组织的证书数量
	TlsCertCount *uint64 `json:"TlsCertCount,omitempty" name:"TlsCertCount"`

	// 网络背书节点证书数量
	PeerCertCount *uint64 `json:"PeerCertCount,omitempty" name:"PeerCertCount"`

	// 当前组织业务证书数量
	ClientCertCount *uint64 `json:"ClientCertCount,omitempty" name:"ClientCertCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetClusterSummaryResponse struct {
	*tchttp.BaseResponse
	Response *GetClusterSummaryResponseParams `json:"Response"`
}

func (r *GetClusterSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetClusterSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetInvokeTxRequestParams struct {
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：query_txid
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 执行该查询交易的节点名称，可以在通道详情中获取该通道上的节点名称及其所属组织名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 执行该查询交易的节点所属组织名称，可以在通道详情中获取该通道上的节点名称及其所属组织名称
	PeerGroup *string `json:"PeerGroup,omitempty" name:"PeerGroup"`

	// 交易ID
	TxId *string `json:"TxId,omitempty" name:"TxId"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type GetInvokeTxRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：query_txid
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 执行该查询交易的节点名称，可以在通道详情中获取该通道上的节点名称及其所属组织名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 执行该查询交易的节点所属组织名称，可以在通道详情中获取该通道上的节点名称及其所属组织名称
	PeerGroup *string `json:"PeerGroup,omitempty" name:"PeerGroup"`

	// 交易ID
	TxId *string `json:"TxId,omitempty" name:"TxId"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

func (r *GetInvokeTxRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetInvokeTxRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "ChannelName")
	delete(f, "PeerName")
	delete(f, "PeerGroup")
	delete(f, "TxId")
	delete(f, "GroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetInvokeTxRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetInvokeTxResponseParams struct {
	// 交易执行状态码
	TxValidationCode *int64 `json:"TxValidationCode,omitempty" name:"TxValidationCode"`

	// 交易执行消息
	TxValidationMsg *string `json:"TxValidationMsg,omitempty" name:"TxValidationMsg"`

	// 交易所在区块ID
	BlockId *int64 `json:"BlockId,omitempty" name:"BlockId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetInvokeTxResponse struct {
	*tchttp.BaseResponse
	Response *GetInvokeTxResponseParams `json:"Response"`
}

func (r *GetInvokeTxResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetInvokeTxResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLatesdTransactionListRequestParams struct {
	// 模块名称，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称，固定字段：latest_transaction_list
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID，固定字段：0
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID，固定字段：0
	ChannelId *uint64 `json:"ChannelId,omitempty" name:"ChannelId"`

	// 获取的最新交易的区块数量，取值范围1~5
	LatestBlockNumber *uint64 `json:"LatestBlockNumber,omitempty" name:"LatestBlockNumber"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 需要查询的通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要获取的起始交易偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要获取的交易数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type GetLatesdTransactionListRequest struct {
	*tchttp.BaseRequest
	
	// 模块名称，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名称，固定字段：latest_transaction_list
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 组织ID，固定字段：0
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 通道ID，固定字段：0
	ChannelId *uint64 `json:"ChannelId,omitempty" name:"ChannelId"`

	// 获取的最新交易的区块数量，取值范围1~5
	LatestBlockNumber *uint64 `json:"LatestBlockNumber,omitempty" name:"LatestBlockNumber"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 需要查询的通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要获取的起始交易偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要获取的交易数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetLatesdTransactionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLatesdTransactionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "GroupId")
	delete(f, "ChannelId")
	delete(f, "LatestBlockNumber")
	delete(f, "GroupName")
	delete(f, "ChannelName")
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLatesdTransactionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLatesdTransactionListResponseParams struct {
	// 交易总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 交易列表
	TransactionList []*TransactionItem `json:"TransactionList,omitempty" name:"TransactionList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetLatesdTransactionListResponse struct {
	*tchttp.BaseResponse
	Response *GetLatesdTransactionListResponseParams `json:"Response"`
}

func (r *GetLatesdTransactionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLatesdTransactionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPeerLogForUserRequestParams struct {
	// 模块名，本接口取值：peer_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：peer_log_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 调用合约的组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 日志开始时间，如"2020-11-24 19:49:25"
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 返回日志行数的最大值，系统设定该参数最大为1000，且一行日志的最大字节数是500，即最大返回50万个字节数的日志数据
	RowNum *int64 `json:"RowNum,omitempty" name:"RowNum"`
}

type GetPeerLogForUserRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：peer_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：peer_log_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 调用合约的组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 日志开始时间，如"2020-11-24 19:49:25"
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 返回日志行数的最大值，系统设定该参数最大为1000，且一行日志的最大字节数是500，即最大返回50万个字节数的日志数据
	RowNum *int64 `json:"RowNum,omitempty" name:"RowNum"`
}

func (r *GetPeerLogForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPeerLogForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	delete(f, "PeerName")
	delete(f, "BeginTime")
	delete(f, "RowNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPeerLogForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPeerLogForUserResponseParams struct {
	// 返回日志总行数，不会超过入参的RowNum
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 日志列表
	PeerLogList []*LogDetailForUser `json:"PeerLogList,omitempty" name:"PeerLogList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPeerLogForUserResponse struct {
	*tchttp.BaseResponse
	Response *GetPeerLogForUserResponseParams `json:"Response"`
}

func (r *GetPeerLogForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPeerLogForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTransactionDetailForUserRequestParams struct {
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：transaction_detail_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 参与交易的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 区块ID，通过GetInvokeTx接口可以获取交易所在的区块ID
	BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

	// 交易ID，需要查询的详情的交易ID
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`
}

type GetTransactionDetailForUserRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：transaction_detail_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 参与交易的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 区块ID，通过GetInvokeTx接口可以获取交易所在的区块ID
	BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

	// 交易ID，需要查询的详情的交易ID
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`
}

func (r *GetTransactionDetailForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTransactionDetailForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	delete(f, "ChannelName")
	delete(f, "BlockId")
	delete(f, "TransactionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTransactionDetailForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTransactionDetailForUserResponseParams struct {
	// 交易ID
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 交易hash
	TransactionHash *string `json:"TransactionHash,omitempty" name:"TransactionHash"`

	// 创建交易的组织名
	CreateOrgName *string `json:"CreateOrgName,omitempty" name:"CreateOrgName"`

	// 交易类型（普通交易和配置交易）
	TransactionType *string `json:"TransactionType,omitempty" name:"TransactionType"`

	// 交易状态
	TransactionStatus *string `json:"TransactionStatus,omitempty" name:"TransactionStatus"`

	// 交易创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 交易数据
	TransactionData *string `json:"TransactionData,omitempty" name:"TransactionData"`

	// 交易所在区块号
	BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

	// 交易所在区块哈希
	BlockHash *string `json:"BlockHash,omitempty" name:"BlockHash"`

	// 交易所在区块高度
	BlockHeight *uint64 `json:"BlockHeight,omitempty" name:"BlockHeight"`

	// 通道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 交易所在合约名称
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 背书组织列表
	EndorserOrgList []*EndorserGroup `json:"EndorserOrgList,omitempty" name:"EndorserOrgList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTransactionDetailForUserResponse struct {
	*tchttp.BaseResponse
	Response *GetTransactionDetailForUserResponseParams `json:"Response"`
}

func (r *GetTransactionDetailForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTransactionDetailForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupDetailForUser struct {
	// 组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 组织MSP Identity
	GroupMSPId *string `json:"GroupMSPId,omitempty" name:"GroupMSPId"`
}

// Predefined struct for user
type InitializeChaincodeForUserRequestParams struct {
	// 模块名，本接口取值：chaincode_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：chaincode_init_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 调用合约的组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 业务所属智能合约名称
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属智能合约版本
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 业务所属通道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 合约实例化节点名称，可以在通道详情中获取该通道上的节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 实例化的函数参数列表
	Args []*string `json:"Args,omitempty" name:"Args"`
}

type InitializeChaincodeForUserRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：chaincode_mng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：chaincode_init_for_user
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 调用合约的组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 业务所属智能合约名称
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属智能合约版本
	ChaincodeVersion *string `json:"ChaincodeVersion,omitempty" name:"ChaincodeVersion"`

	// 业务所属通道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 合约实例化节点名称，可以在通道详情中获取该通道上的节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 实例化的函数参数列表
	Args []*string `json:"Args,omitempty" name:"Args"`
}

func (r *InitializeChaincodeForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitializeChaincodeForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	delete(f, "ChaincodeName")
	delete(f, "ChaincodeVersion")
	delete(f, "ChannelName")
	delete(f, "PeerName")
	delete(f, "Args")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InitializeChaincodeForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InitializeChaincodeForUserResponseParams struct {
	// 实例化任务ID，用于查询实例化结果
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InitializeChaincodeForUserResponse struct {
	*tchttp.BaseResponse
	Response *InitializeChaincodeForUserResponseParams `json:"Response"`
}

func (r *InitializeChaincodeForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitializeChaincodeForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeBcosTransRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组编号，可在群组列表中获取
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 合约地址，可在合约详情获取
	ContractAddress *string `json:"ContractAddress,omitempty" name:"ContractAddress"`

	// 合约Abi的json数组格式的字符串，可在合约详情获取
	AbiInfo *string `json:"AbiInfo,omitempty" name:"AbiInfo"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 签名用户编号，可在私钥管理页面获取
	SignUserId *string `json:"SignUserId,omitempty" name:"SignUserId"`

	// 合约方法入参，json格式字符串
	FuncParam *string `json:"FuncParam,omitempty" name:"FuncParam"`
}

type InvokeBcosTransRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组编号，可在群组列表中获取
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 合约地址，可在合约详情获取
	ContractAddress *string `json:"ContractAddress,omitempty" name:"ContractAddress"`

	// 合约Abi的json数组格式的字符串，可在合约详情获取
	AbiInfo *string `json:"AbiInfo,omitempty" name:"AbiInfo"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 签名用户编号，可在私钥管理页面获取
	SignUserId *string `json:"SignUserId,omitempty" name:"SignUserId"`

	// 合约方法入参，json格式字符串
	FuncParam *string `json:"FuncParam,omitempty" name:"FuncParam"`
}

func (r *InvokeBcosTransRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeBcosTransRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "ContractAddress")
	delete(f, "AbiInfo")
	delete(f, "FuncName")
	delete(f, "SignUserId")
	delete(f, "FuncParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvokeBcosTransRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeBcosTransResponseParams struct {
	// 交易结果json字符串
	TransactionRsp *string `json:"TransactionRsp,omitempty" name:"TransactionRsp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InvokeBcosTransResponse struct {
	*tchttp.BaseResponse
	Response *InvokeBcosTransResponseParams `json:"Response"`
}

func (r *InvokeBcosTransResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeBcosTransResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeChainMakerContractRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitempty" name:"FuncParam"`

	// 是否异步执行，1为是，否则为0；如果异步执行，可使用返回值中的交易TxID查询执行结果
	AsyncFlag *int64 `json:"AsyncFlag,omitempty" name:"AsyncFlag"`
}

type InvokeChainMakerContractRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitempty" name:"FuncParam"`

	// 是否异步执行，1为是，否则为0；如果异步执行，可使用返回值中的交易TxID查询执行结果
	AsyncFlag *int64 `json:"AsyncFlag,omitempty" name:"AsyncFlag"`
}

func (r *InvokeChainMakerContractRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeChainMakerContractRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ChainId")
	delete(f, "ContractName")
	delete(f, "FuncName")
	delete(f, "FuncParam")
	delete(f, "AsyncFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvokeChainMakerContractRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeChainMakerContractResponseParams struct {
	// 交易结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ChainMakerContractResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InvokeChainMakerContractResponse struct {
	*tchttp.BaseResponse
	Response *InvokeChainMakerContractResponseParams `json:"Response"`
}

func (r *InvokeChainMakerContractResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeChainMakerContractResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeChainMakerDemoContractRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitempty" name:"FuncParam"`

	// 是否异步执行，1为是，否则为0；如果异步执行，可使用返回值中的交易TxID查询执行结果
	AsyncFlag *int64 `json:"AsyncFlag,omitempty" name:"AsyncFlag"`
}

type InvokeChainMakerDemoContractRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitempty" name:"FuncParam"`

	// 是否异步执行，1为是，否则为0；如果异步执行，可使用返回值中的交易TxID查询执行结果
	AsyncFlag *int64 `json:"AsyncFlag,omitempty" name:"AsyncFlag"`
}

func (r *InvokeChainMakerDemoContractRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeChainMakerDemoContractRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ChainId")
	delete(f, "ContractName")
	delete(f, "FuncName")
	delete(f, "FuncParam")
	delete(f, "AsyncFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvokeChainMakerDemoContractRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeChainMakerDemoContractResponseParams struct {
	// 交易结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ChainMakerContractResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InvokeChainMakerDemoContractResponse struct {
	*tchttp.BaseResponse
	Response *InvokeChainMakerDemoContractResponseParams `json:"Response"`
}

func (r *InvokeChainMakerDemoContractResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeChainMakerDemoContractResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeRequestParams struct {
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：invoke
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务所属智能合约名称，可在智能合约详情或列表中获取
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 对该笔交易进行背书的节点列表（包括节点名称和节点所属组织名称，详见数据结构一节），可以在通道详情中获取该通道上的节点名称及其所属组织名称
	Peers []*PeerSet `json:"Peers,omitempty" name:"Peers"`

	// 该笔交易需要调用的智能合约中的函数名称
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 被调用的函数参数列表，参数列表大小总和要求小于2M
	Args []*string `json:"Args,omitempty" name:"Args"`

	// 同步调用标识，可选参数，值为0或者不传表示使用同步方法调用，调用后会等待交易执行后再返回执行结果；值为1时表示使用异步方式调用Invoke，执行后会立即返回交易对应的Txid，后续需要通过GetInvokeTx这个API查询该交易的执行结果。（对于逻辑较为简单的交易，可以使用同步模式；对于逻辑较为复杂的交易，建议使用异步模式，否则容易导致API因等待时间过长，返回等待超时）
	AsyncFlag *uint64 `json:"AsyncFlag,omitempty" name:"AsyncFlag"`
}

type InvokeRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：invoke
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务所属智能合约名称，可在智能合约详情或列表中获取
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 对该笔交易进行背书的节点列表（包括节点名称和节点所属组织名称，详见数据结构一节），可以在通道详情中获取该通道上的节点名称及其所属组织名称
	Peers []*PeerSet `json:"Peers,omitempty" name:"Peers"`

	// 该笔交易需要调用的智能合约中的函数名称
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 被调用的函数参数列表，参数列表大小总和要求小于2M
	Args []*string `json:"Args,omitempty" name:"Args"`

	// 同步调用标识，可选参数，值为0或者不传表示使用同步方法调用，调用后会等待交易执行后再返回执行结果；值为1时表示使用异步方式调用Invoke，执行后会立即返回交易对应的Txid，后续需要通过GetInvokeTx这个API查询该交易的执行结果。（对于逻辑较为简单的交易，可以使用同步模式；对于逻辑较为复杂的交易，建议使用异步模式，否则容易导致API因等待时间过长，返回等待超时）
	AsyncFlag *uint64 `json:"AsyncFlag,omitempty" name:"AsyncFlag"`
}

func (r *InvokeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "ChaincodeName")
	delete(f, "ChannelName")
	delete(f, "Peers")
	delete(f, "FuncName")
	delete(f, "GroupName")
	delete(f, "Args")
	delete(f, "AsyncFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvokeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeResponseParams struct {
	// 交易ID
	Txid *string `json:"Txid,omitempty" name:"Txid"`

	// 交易执行结果
	Events *string `json:"Events,omitempty" name:"Events"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InvokeResponse struct {
	*tchttp.BaseResponse
	Response *InvokeResponseParams `json:"Response"`
}

func (r *InvokeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogDetailForUser struct {
	// 日志行号
	LineNumber *uint64 `json:"LineNumber,omitempty" name:"LineNumber"`

	// 日志详情
	LogMessage *string `json:"LogMessage,omitempty" name:"LogMessage"`
}

type PeerDetailForUser struct {
	// 节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`
}

type PeerSet struct {
	// 节点名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 组织名称
	OrgName *string `json:"OrgName,omitempty" name:"OrgName"`
}

// Predefined struct for user
type QueryChainMakerBlockTransactionRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitempty" name:"BlockHeight"`
}

type QueryChainMakerBlockTransactionRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitempty" name:"BlockHeight"`
}

func (r *QueryChainMakerBlockTransactionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChainMakerBlockTransactionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ChainId")
	delete(f, "BlockHeight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryChainMakerBlockTransactionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChainMakerBlockTransactionResponseParams struct {
	// 区块交易
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*ChainMakerTransactionResult `json:"Result,omitempty" name:"Result"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitempty" name:"BlockHeight"`

	// 交易数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TxCount *int64 `json:"TxCount,omitempty" name:"TxCount"`

	// 区块时间戳，单位是秒
	BlockTimestamp *int64 `json:"BlockTimestamp,omitempty" name:"BlockTimestamp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryChainMakerBlockTransactionResponse struct {
	*tchttp.BaseResponse
	Response *QueryChainMakerBlockTransactionResponseParams `json:"Response"`
}

func (r *QueryChainMakerBlockTransactionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChainMakerBlockTransactionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChainMakerContractRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitempty" name:"FuncParam"`
}

type QueryChainMakerContractRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitempty" name:"FuncParam"`
}

func (r *QueryChainMakerContractRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChainMakerContractRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ChainId")
	delete(f, "ContractName")
	delete(f, "FuncName")
	delete(f, "FuncParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryChainMakerContractRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChainMakerContractResponseParams struct {
	// 交易结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ChainMakerContractResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryChainMakerContractResponse struct {
	*tchttp.BaseResponse
	Response *QueryChainMakerContractResponseParams `json:"Response"`
}

func (r *QueryChainMakerContractResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChainMakerContractResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChainMakerDemoBlockTransactionRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitempty" name:"BlockHeight"`
}

type QueryChainMakerDemoBlockTransactionRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitempty" name:"BlockHeight"`
}

func (r *QueryChainMakerDemoBlockTransactionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChainMakerDemoBlockTransactionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ChainId")
	delete(f, "BlockHeight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryChainMakerDemoBlockTransactionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChainMakerDemoBlockTransactionResponseParams struct {
	// 区块交易
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*ChainMakerTransactionResult `json:"Result,omitempty" name:"Result"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitempty" name:"BlockHeight"`

	// 交易数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TxCount *int64 `json:"TxCount,omitempty" name:"TxCount"`

	// 区块时间戳，单位是秒
	BlockTimestamp *int64 `json:"BlockTimestamp,omitempty" name:"BlockTimestamp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryChainMakerDemoBlockTransactionResponse struct {
	*tchttp.BaseResponse
	Response *QueryChainMakerDemoBlockTransactionResponseParams `json:"Response"`
}

func (r *QueryChainMakerDemoBlockTransactionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChainMakerDemoBlockTransactionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChainMakerDemoContractRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitempty" name:"FuncParam"`
}

type QueryChainMakerDemoContractRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitempty" name:"FuncParam"`
}

func (r *QueryChainMakerDemoContractRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChainMakerDemoContractRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ChainId")
	delete(f, "ContractName")
	delete(f, "FuncName")
	delete(f, "FuncParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryChainMakerDemoContractRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChainMakerDemoContractResponseParams struct {
	// 交易结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ChainMakerContractResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryChainMakerDemoContractResponse struct {
	*tchttp.BaseResponse
	Response *QueryChainMakerDemoContractResponseParams `json:"Response"`
}

func (r *QueryChainMakerDemoContractResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChainMakerDemoContractResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChainMakerDemoTransactionRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 交易ID，通过调用合约的返回值获取
	TxID *string `json:"TxID,omitempty" name:"TxID"`
}

type QueryChainMakerDemoTransactionRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 交易ID，通过调用合约的返回值获取
	TxID *string `json:"TxID,omitempty" name:"TxID"`
}

func (r *QueryChainMakerDemoTransactionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChainMakerDemoTransactionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ChainId")
	delete(f, "TxID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryChainMakerDemoTransactionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChainMakerDemoTransactionResponseParams struct {
	// 交易结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ChainMakerTransactionResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryChainMakerDemoTransactionResponse struct {
	*tchttp.BaseResponse
	Response *QueryChainMakerDemoTransactionResponseParams `json:"Response"`
}

func (r *QueryChainMakerDemoTransactionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChainMakerDemoTransactionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChainMakerTransactionRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 交易ID，通过调用合约的返回值获取
	TxID *string `json:"TxID,omitempty" name:"TxID"`
}

type QueryChainMakerTransactionRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitempty" name:"ChainId"`

	// 交易ID，通过调用合约的返回值获取
	TxID *string `json:"TxID,omitempty" name:"TxID"`
}

func (r *QueryChainMakerTransactionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChainMakerTransactionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ChainId")
	delete(f, "TxID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryChainMakerTransactionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryChainMakerTransactionResponseParams struct {
	// 交易结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ChainMakerTransactionResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryChainMakerTransactionResponse struct {
	*tchttp.BaseResponse
	Response *QueryChainMakerTransactionResponseParams `json:"Response"`
}

func (r *QueryChainMakerTransactionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryChainMakerTransactionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryRequestParams struct {
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：query
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务所属智能合约名称，可在智能合约详情或列表中获取
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 执行该查询交易的节点列表（包括节点名称和节点所属组织名称，详见数据结构一节），可以在通道详情中获取该通道上的节点名称及其所属组织名称
	Peers []*PeerSet `json:"Peers,omitempty" name:"Peers"`

	// 该笔交易查询需要调用的智能合约中的函数名称
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 被调用的函数参数列表
	Args []*string `json:"Args,omitempty" name:"Args"`
}

type QueryRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：query
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 业务所属智能合约名称，可在智能合约详情或列表中获取
	ChaincodeName *string `json:"ChaincodeName,omitempty" name:"ChaincodeName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 执行该查询交易的节点列表（包括节点名称和节点所属组织名称，详见数据结构一节），可以在通道详情中获取该通道上的节点名称及其所属组织名称
	Peers []*PeerSet `json:"Peers,omitempty" name:"Peers"`

	// 该笔交易查询需要调用的智能合约中的函数名称
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 被调用的函数参数列表
	Args []*string `json:"Args,omitempty" name:"Args"`
}

func (r *QueryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ClusterId")
	delete(f, "ChaincodeName")
	delete(f, "ChannelName")
	delete(f, "Peers")
	delete(f, "FuncName")
	delete(f, "GroupName")
	delete(f, "Args")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryResponseParams struct {
	// 查询结果数据
	Data []*string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryResponse struct {
	*tchttp.BaseResponse
	Response *QueryResponseParams `json:"Response"`
}

func (r *QueryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SignCertCsr struct {
	// 用户签名证书的标识，会存在于用户申请的证书中
	CertMark *string `json:"CertMark,omitempty" name:"CertMark"`

	// 用户申请签名证书所需要的证书请求文件的base64编码
	SignCsrContent *string `json:"SignCsrContent,omitempty" name:"SignCsrContent"`
}

// Predefined struct for user
type SrvInvokeRequestParams struct {
	// 服务类型，iss或者dam
	Service *string `json:"Service,omitempty" name:"Service"`

	// 服务接口，要调用的方法函数名
	Method *string `json:"Method,omitempty" name:"Method"`

	// 用户自定义json字符串
	Param *string `json:"Param,omitempty" name:"Param"`
}

type SrvInvokeRequest struct {
	*tchttp.BaseRequest
	
	// 服务类型，iss或者dam
	Service *string `json:"Service,omitempty" name:"Service"`

	// 服务接口，要调用的方法函数名
	Method *string `json:"Method,omitempty" name:"Method"`

	// 用户自定义json字符串
	Param *string `json:"Param,omitempty" name:"Param"`
}

func (r *SrvInvokeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SrvInvokeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Service")
	delete(f, "Method")
	delete(f, "Param")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SrvInvokeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SrvInvokeResponseParams struct {
	// 返回码
	RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`

	// 返回消息
	RetMsg *string `json:"RetMsg,omitempty" name:"RetMsg"`

	// 返回数据
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SrvInvokeResponse struct {
	*tchttp.BaseResponse
	Response *SrvInvokeResponseParams `json:"Response"`
}

func (r *SrvInvokeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SrvInvokeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransactionItem struct {
	// 交易ID
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 交易hash
	TransactionHash *string `json:"TransactionHash,omitempty" name:"TransactionHash"`

	// 创建交易的组织名
	CreateOrgName *string `json:"CreateOrgName,omitempty" name:"CreateOrgName"`

	// 交易所在区块号
	BlockId *uint64 `json:"BlockId,omitempty" name:"BlockId"`

	// 交易类型（普通交易和配置交易）
	TransactionType *string `json:"TransactionType,omitempty" name:"TransactionType"`

	// 交易创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 交易所在区块高度
	BlockHeight *uint64 `json:"BlockHeight,omitempty" name:"BlockHeight"`

	// 交易状态
	TransactionStatus *string `json:"TransactionStatus,omitempty" name:"TransactionStatus"`
}