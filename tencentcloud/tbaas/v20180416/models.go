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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("ApplyUserCertRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplyUserCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书ID
		CertId *uint64 `json:"CertId,omitempty" name:"CertId"`

		// 证书DN
		CertDn *string `json:"CertDn,omitempty" name:"CertDn"`

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

type BlockByNumberHandlerRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定字段：block
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：block_by_number
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 当前群组编号
	GroupPk *string `json:"GroupPk,omitempty" name:"GroupPk"`

	// 区块高度
	BlockNumber *int64 `json:"BlockNumber,omitempty" name:"BlockNumber"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BlockByNumberHandlerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "GroupPk")
	delete(f, "BlockNumber")
	if len(f) > 0 {
		return errors.New("BlockByNumberHandlerRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BlockByNumberHandlerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回区块json字符串
		BlockJson *string `json:"BlockJson,omitempty" name:"BlockJson"`

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
func (r *BlockByNumberHandlerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChannelDetailForUser struct {

	// 通道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 当前组织加入通道的节点列表
	PeerList []*PeerDetailForUser `json:"PeerList,omitempty" name:"PeerList" list`
}

type ClusterDetailForUser struct {

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 组织列表
	GroupList []*GroupDetailForUser `json:"GroupList,omitempty" name:"GroupList" list`

	// 网络名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("CreateChaincodeAndInstallForUserRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateChaincodeAndInstallForUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
func (r *CreateChaincodeAndInstallForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DeployDynamicBcosContractRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeployDynamicBcosContractResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 部署成功返回的合约地址
		ContractAddress *string `json:"ContractAddress,omitempty" name:"ContractAddress"`

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
func (r *DeployDynamicBcosContractResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployDynamicContractHandlerRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定字段：contract
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：deploy_dynamic_contract
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 群组编号
	GroupPk *string `json:"GroupPk,omitempty" name:"GroupPk"`

	// 合约名称
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 合约编译后的abi
	AbiInfo *string `json:"AbiInfo,omitempty" name:"AbiInfo"`

	// 合约编译后的binary
	ByteCodeBin *string `json:"ByteCodeBin,omitempty" name:"ByteCodeBin"`

	// 构造函数入参
	ConstructorParams []*string `json:"ConstructorParams,omitempty" name:"ConstructorParams" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployDynamicContractHandlerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "GroupPk")
	delete(f, "ContractName")
	delete(f, "AbiInfo")
	delete(f, "ByteCodeBin")
	delete(f, "ConstructorParams")
	if len(f) > 0 {
		return errors.New("DeployDynamicContractHandlerRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeployDynamicContractHandlerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 部署成功返回的合约地址
		ContractAddress *string `json:"ContractAddress,omitempty" name:"ContractAddress"`

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
func (r *DeployDynamicContractHandlerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DownloadUserCertRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DownloadUserCertResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书名称
		CertName *string `json:"CertName,omitempty" name:"CertName"`

		// 证书内容
		CertCtx *string `json:"CertCtx,omitempty" name:"CertCtx"`

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
func (r *DownloadUserCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EndorserGroup struct {

	// 背书组织名称
	EndorserGroupName *string `json:"EndorserGroupName,omitempty" name:"EndorserGroupName"`

	// 背书节点列表
	EndorserPeerList []*string `json:"EndorserPeerList,omitempty" name:"EndorserPeerList" list`
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetBcosBlockByNumberRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetBcosBlockByNumberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回区块json字符串
		BlockJson *string `json:"BlockJson,omitempty" name:"BlockJson"`

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
func (r *GetBcosBlockByNumberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetBcosBlockListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetBcosBlockListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回数据列表
		List []*BcosBlockObj `json:"List,omitempty" name:"List" list`

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
func (r *GetBcosBlockListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetBcosTransByHashRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetBcosTransByHashResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易信息json字符串
		TransactionJson *string `json:"TransactionJson,omitempty" name:"TransactionJson"`

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
func (r *GetBcosTransByHashResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetBcosTransListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetBcosTransListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回数据列表
		List []*BcosTransInfo `json:"List,omitempty" name:"List" list`

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
func (r *GetBcosTransListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBlockListHandlerRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定字段：block
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：get_block_list
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 记录偏移数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 当前群组编号
	GroupPk *string `json:"GroupPk,omitempty" name:"GroupPk"`

	// 区块哈希
	BlockHash *string `json:"BlockHash,omitempty" name:"BlockHash"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBlockListHandlerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "GroupPk")
	delete(f, "BlockHash")
	if len(f) > 0 {
		return errors.New("GetBlockListHandlerRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetBlockListHandlerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 当前群组编号
		GroupPk *string `json:"GroupPk,omitempty" name:"GroupPk"`

		// 返回数据列表
		List []*BcosBlockObj `json:"List,omitempty" name:"List" list`

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
func (r *GetBlockListHandlerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetBlockListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetBlockListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 区块数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 区块列表
		BlockList []*Block `json:"BlockList,omitempty" name:"BlockList" list`

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
func (r *GetBlockListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetBlockTransactionListForUserRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetBlockTransactionListForUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 交易列表
		TransactionList []*TransactionItem `json:"TransactionList,omitempty" name:"TransactionList" list`

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
func (r *GetBlockTransactionListForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetChaincodeCompileLogForUserRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeCompileLogForUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志总行数，上限2000条日志
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 日志列表
		CompileLogList []*LogDetailForUser `json:"CompileLogList,omitempty" name:"CompileLogList" list`

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
func (r *GetChaincodeCompileLogForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetChaincodeInitializeResultForUserRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeInitializeResultForUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例化结果：0，实例化中；1，实例化成功；2，实例化失败
		InitResult *uint64 `json:"InitResult,omitempty" name:"InitResult"`

		// 实例化信息
		InitMessage *string `json:"InitMessage,omitempty" name:"InitMessage"`

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
func (r *GetChaincodeInitializeResultForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetChaincodeLogForUserRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetChaincodeLogForUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回日志总行数，不会超过入参的RowNum
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 日志列表
		ChaincodeLogList []*LogDetailForUser `json:"ChaincodeLogList,omitempty" name:"ChaincodeLogList" list`

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
func (r *GetChaincodeLogForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetChannelListForUserRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetChannelListForUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通道总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 通道列表
		ChannelList []*ChannelDetailForUser `json:"ChannelList,omitempty" name:"ChannelList" list`

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
func (r *GetChannelListForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetClusterListForUserRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterListForUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 网络列表
		ClusterList []*ClusterDetailForUser `json:"ClusterList,omitempty" name:"ClusterList" list`

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
func (r *GetClusterListForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetClusterSummaryRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetClusterSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// 执行该查询交易的节点名称，可以在通道详情中获取该通道上的节点名称极其所属组织名称
	PeerName *string `json:"PeerName,omitempty" name:"PeerName"`

	// 执行该查询交易的节点所属组织名称，可以在通道详情中获取该通道上的节点名称极其所属组织名称
	PeerGroup *string `json:"PeerGroup,omitempty" name:"PeerGroup"`

	// 交易ID
	TxId *string `json:"TxId,omitempty" name:"TxId"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetInvokeTxRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetInvokeTxResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易执行状态码
		TxValidationCode *int64 `json:"TxValidationCode,omitempty" name:"TxValidationCode"`

		// 交易执行消息
		TxValidationMsg *string `json:"TxValidationMsg,omitempty" name:"TxValidationMsg"`

		// 交易所在区块ID
		BlockId *int64 `json:"BlockId,omitempty" name:"BlockId"`

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
func (r *GetInvokeTxResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetLatesdTransactionListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetLatesdTransactionListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 交易列表
		TransactionList []*TransactionItem `json:"TransactionList,omitempty" name:"TransactionList" list`

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
func (r *GetLatesdTransactionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetPeerLogForUserRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetPeerLogForUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回日志总行数，不会超过入参的RowNum
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 日志列表
		PeerLogList []*LogDetailForUser `json:"PeerLogList,omitempty" name:"PeerLogList" list`

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
func (r *GetPeerLogForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTransByHashHandlerRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：get_trans_by_hash
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 群组编号
	GroupPk *string `json:"GroupPk,omitempty" name:"GroupPk"`

	// 交易哈希
	TransHash *string `json:"TransHash,omitempty" name:"TransHash"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTransByHashHandlerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "GroupPk")
	delete(f, "TransHash")
	if len(f) > 0 {
		return errors.New("GetTransByHashHandlerRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetTransByHashHandlerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易信息json字符串
		TransactionJson *string `json:"TransactionJson,omitempty" name:"TransactionJson"`

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
func (r *GetTransByHashHandlerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTransListHandlerRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：get_trans_list
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 记录偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 群组编号
	GroupPk *string `json:"GroupPk,omitempty" name:"GroupPk"`

	// 交易哈希
	TransHash *string `json:"TransHash,omitempty" name:"TransHash"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTransListHandlerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "GroupPk")
	delete(f, "TransHash")
	if len(f) > 0 {
		return errors.New("GetTransListHandlerRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetTransListHandlerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 当前群组编号
		GroupPk *string `json:"GroupPk,omitempty" name:"GroupPk"`

		// 返回数据列表
		List []*BcosTransInfo `json:"List,omitempty" name:"List" list`

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
func (r *GetTransListHandlerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("GetTransactionDetailForUserRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetTransactionDetailForUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		EndorserOrgList []*EndorserGroup `json:"EndorserOrgList,omitempty" name:"EndorserOrgList" list`

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
func (r *GetTransactionDetailForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupDetailForUser struct {

	// 组织名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 组织MSP Identity
	GroupMSPId *string `json:"GroupMSPId,omitempty" name:"GroupMSPId"`
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
	Args []*string `json:"Args,omitempty" name:"Args" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("InitializeChaincodeForUserRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InitializeChaincodeForUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例化任务ID，用于查询实例化结果
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

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
func (r *InitializeChaincodeForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("InvokeBcosTransRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InvokeBcosTransResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易结果json字符串
		TransactionRsp *string `json:"TransactionRsp,omitempty" name:"TransactionRsp"`

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
func (r *InvokeBcosTransResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	Peers []*PeerSet `json:"Peers,omitempty" name:"Peers" list`

	// 该笔交易需要调用的智能合约中的函数名称
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 被调用的函数参数列表，参数列表大小总和要求小于2M
	Args []*string `json:"Args,omitempty" name:"Args" list`

	// 同步调用标识，可选参数，值为0或者不传表示使用同步方法调用，调用后会等待交易执行后再返回执行结果；值为1时表示使用异步方式调用Invoke，执行后会立即返回交易对应的Txid，后续需要通过GetInvokeTx这个API查询该交易的执行结果。（对于逻辑较为简单的交易，可以使用同步模式；对于逻辑较为复杂的交易，建议使用异步模式，否则容易导致API因等待时间过长，返回等待超时）
	AsyncFlag *uint64 `json:"AsyncFlag,omitempty" name:"AsyncFlag"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("InvokeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InvokeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易ID
		Txid *string `json:"Txid,omitempty" name:"Txid"`

		// 交易执行结果
		Events *string `json:"Events,omitempty" name:"Events"`

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
	Peers []*PeerSet `json:"Peers,omitempty" name:"Peers" list`

	// 该笔交易查询需要调用的智能合约中的函数名称
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 被调用的函数参数列表
	Args []*string `json:"Args,omitempty" name:"Args" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("QueryRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询结果数据
		Data []*string `json:"Data,omitempty" name:"Data" list`

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
func (r *QueryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendTransactionHandlerRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：send_transaction
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 群组编号
	GroupPk *string `json:"GroupPk,omitempty" name:"GroupPk"`

	// 合约编号
	ContractId *int64 `json:"ContractId,omitempty" name:"ContractId"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 合约方法入参
	FuncParam []*string `json:"FuncParam,omitempty" name:"FuncParam" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendTransactionHandlerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "GroupPk")
	delete(f, "ContractId")
	delete(f, "FuncName")
	delete(f, "FuncParam")
	if len(f) > 0 {
		return errors.New("SendTransactionHandlerRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendTransactionHandlerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易结果json字符串
		TransactionRsp *string `json:"TransactionRsp,omitempty" name:"TransactionRsp"`

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
func (r *SendTransactionHandlerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("SrvInvokeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SrvInvokeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回码
		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`

		// 返回消息
		RetMsg *string `json:"RetMsg,omitempty" name:"RetMsg"`

		// 返回数据
		Data *string `json:"Data,omitempty" name:"Data"`

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
func (r *SrvInvokeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransByDynamicContractHandlerRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，固定字段：trans_by_dynamic_contract
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 群组编号
	GroupPk *string `json:"GroupPk,omitempty" name:"GroupPk"`

	// 合约地址（合约部署成功，可得到合约地址）
	ContractAddress *string `json:"ContractAddress,omitempty" name:"ContractAddress"`

	// 合约名
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 合约编译后的abi
	AbiInfo *string `json:"AbiInfo,omitempty" name:"AbiInfo"`

	// 合约被调用方法名
	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`

	// 合约被调用方法的入参
	FuncParam []*string `json:"FuncParam,omitempty" name:"FuncParam" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransByDynamicContractHandlerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "GroupPk")
	delete(f, "ContractAddress")
	delete(f, "ContractName")
	delete(f, "AbiInfo")
	delete(f, "FuncName")
	delete(f, "FuncParam")
	if len(f) > 0 {
		return errors.New("TransByDynamicContractHandlerRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TransByDynamicContractHandlerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易结果json字符串
		TransactionRsp *string `json:"TransactionRsp,omitempty" name:"TransactionRsp"`

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
func (r *TransByDynamicContractHandlerResponse) FromJsonString(s string) error {
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
