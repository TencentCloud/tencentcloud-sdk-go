// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type ApplyChainMakerBatchUserCertRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 证书标识和证书请求文件，可参考TBaaS证书生成相关文档生成证书请求文件
	SignUserCsrList []*SignCertCsr `json:"SignUserCsrList,omitnil,omitempty" name:"SignUserCsrList"`
}

type ApplyChainMakerBatchUserCertRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 证书标识和证书请求文件，可参考TBaaS证书生成相关文档生成证书请求文件
	SignUserCsrList []*SignCertCsr `json:"SignUserCsrList,omitnil,omitempty" name:"SignUserCsrList"`
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
	SignUserCrtList []*string `json:"SignUserCrtList,omitnil,omitempty" name:"SignUserCrtList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：cert_apply_for_user
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 申请证书的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户证书标识，用于标识用户证书，要求由纯小写字母组成，长度小于10
	UserIdentity *string `json:"UserIdentity,omitnil,omitempty" name:"UserIdentity"`

	// 证书申请实体，使用腾讯云账号实名认证的名称
	Applicant *string `json:"Applicant,omitnil,omitempty" name:"Applicant"`

	// 证件号码。如果腾讯云账号对应的实名认证类型为企业认证，填入“0”；如果腾讯云账号对应的实名认证类型为个人认证，填入个人身份证号码
	IdentityNum *string `json:"IdentityNum,omitnil,omitempty" name:"IdentityNum"`

	// csr p10证书文件。需要用户根据文档生成证书的CSR文件
	CsrData *string `json:"CsrData,omitnil,omitempty" name:"CsrData"`

	// 证书备注信息
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`
}

type ApplyUserCertRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：cert_mng
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：cert_apply_for_user
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 申请证书的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户证书标识，用于标识用户证书，要求由纯小写字母组成，长度小于10
	UserIdentity *string `json:"UserIdentity,omitnil,omitempty" name:"UserIdentity"`

	// 证书申请实体，使用腾讯云账号实名认证的名称
	Applicant *string `json:"Applicant,omitnil,omitempty" name:"Applicant"`

	// 证件号码。如果腾讯云账号对应的实名认证类型为企业认证，填入“0”；如果腾讯云账号对应的实名认证类型为个人认证，填入个人身份证号码
	IdentityNum *string `json:"IdentityNum,omitnil,omitempty" name:"IdentityNum"`

	// csr p10证书文件。需要用户根据文档生成证书的CSR文件
	CsrData *string `json:"CsrData,omitnil,omitempty" name:"CsrData"`

	// 证书备注信息
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`
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
	CertId *uint64 `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书DN
	CertDn *string `json:"CertDn,omitnil,omitempty" name:"CertDn"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type Block struct {
	// 区块编号
	BlockNum *uint64 `json:"BlockNum,omitnil,omitempty" name:"BlockNum"`

	// 区块数据Hash数值
	DataHash *string `json:"DataHash,omitnil,omitempty" name:"DataHash"`

	// 区块ID，与区块编号一致
	BlockId *uint64 `json:"BlockId,omitnil,omitempty" name:"BlockId"`

	// 前一个区块Hash
	PreHash *string `json:"PreHash,omitnil,omitempty" name:"PreHash"`

	// 区块内的交易数量
	TxCount *uint64 `json:"TxCount,omitnil,omitempty" name:"TxCount"`
}

type ChainMakerContractResult struct {
	// 交易结果码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 交易结果码含义
	CodeMessage *string `json:"CodeMessage,omitnil,omitempty" name:"CodeMessage"`

	// 交易ID
	TxId *string `json:"TxId,omitnil,omitempty" name:"TxId"`

	// Gas使用量
	GasUsed *int64 `json:"GasUsed,omitnil,omitempty" name:"GasUsed"`

	// 合约返回消息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 合约函数返回，base64编码
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`
}

type ChainMakerTransactionResult struct {
	// 交易结果码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 交易结果码含义
	CodeMessage *string `json:"CodeMessage,omitnil,omitempty" name:"CodeMessage"`

	// 交易ID
	TxId *string `json:"TxId,omitnil,omitempty" name:"TxId"`

	// Gas使用量
	GasUsed *int64 `json:"GasUsed,omitnil,omitempty" name:"GasUsed"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`

	// 合约执行结果
	ContractEvent *string `json:"ContractEvent,omitnil,omitempty" name:"ContractEvent"`

	// 合约返回信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 交易时间，单位是秒
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

// Predefined struct for user
type DescribeFabricBlockRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 通道ID，可在通道列表或通道详情获取
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 区块高度，从0开始
	BlockHeight *int64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`
}

type DescribeFabricBlockRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 通道ID，可在通道列表或通道详情获取
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 区块高度，从0开始
	BlockHeight *int64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`
}

func (r *DescribeFabricBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFabricBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ChannelId")
	delete(f, "BlockHeight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFabricBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFabricBlockResponseParams struct {
	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`

	// 区块Hash
	BlockHash *string `json:"BlockHash,omitnil,omitempty" name:"BlockHash"`

	// 前置区块Hash
	PreBlockHash *string `json:"PreBlockHash,omitnil,omitempty" name:"PreBlockHash"`

	// 区块中交易数量
	TxCount *int64 `json:"TxCount,omitnil,omitempty" name:"TxCount"`

	// 区块中交易列表
	TransactionList []*Transaction `json:"TransactionList,omitnil,omitempty" name:"TransactionList"`

	// 创建时间戳
	CreateTimestamp *string `json:"CreateTimestamp,omitnil,omitempty" name:"CreateTimestamp"`

	// 提案组织
	ProposerOrg *string `json:"ProposerOrg,omitnil,omitempty" name:"ProposerOrg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFabricBlockResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFabricBlockResponseParams `json:"Response"`
}

func (r *DescribeFabricBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFabricBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFabricTransactionRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 通道ID，可在通道列表或通道详情获取
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 交易ID
	TxId *string `json:"TxId,omitnil,omitempty" name:"TxId"`
}

type DescribeFabricTransactionRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 通道ID，可在通道列表或通道详情获取
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 交易ID
	TxId *string `json:"TxId,omitnil,omitempty" name:"TxId"`
}

func (r *DescribeFabricTransactionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFabricTransactionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ChannelId")
	delete(f, "TxId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFabricTransactionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFabricTransactionResponseParams struct {
	// 交易ID
	TxId *string `json:"TxId,omitnil,omitempty" name:"TxId"`

	// 交易Hash
	TxHash *string `json:"TxHash,omitnil,omitempty" name:"TxHash"`

	// 交易状态
	TxStatus *string `json:"TxStatus,omitnil,omitempty" name:"TxStatus"`

	// 参与的组织列表
	JoinOrgList []*string `json:"JoinOrgList,omitnil,omitempty" name:"JoinOrgList"`

	// 交易发送者
	Sender *string `json:"Sender,omitnil,omitempty" name:"Sender"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`

	// 交易所属合约
	ChaincodeName *string `json:"ChaincodeName,omitnil,omitempty" name:"ChaincodeName"`

	// 交易数据，base64编码，解码后为json化的字符串
	TransactionData *string `json:"TransactionData,omitnil,omitempty" name:"TransactionData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFabricTransactionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFabricTransactionResponseParams `json:"Response"`
}

func (r *DescribeFabricTransactionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFabricTransactionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadUserCertRequestParams struct {
	// 模块名，固定字段：cert_mng
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：cert_download_for_user
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 证书ID，可以在证书详情页面获取
	CertId *uint64 `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书DN，可以在证书详情页面获取
	CertDn *string `json:"CertDn,omitnil,omitempty" name:"CertDn"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 下载证书的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type DownloadUserCertRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：cert_mng
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：cert_download_for_user
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 证书ID，可以在证书详情页面获取
	CertId *uint64 `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书DN，可以在证书详情页面获取
	CertDn *string `json:"CertDn,omitnil,omitempty" name:"CertDn"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 下载证书的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
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
	CertName *string `json:"CertName,omitnil,omitempty" name:"CertName"`

	// 证书内容
	//
	// Deprecated: CertCtx is deprecated.
	CertCtx *string `json:"CertCtx,omitnil,omitempty" name:"CertCtx"`

	// 证书内容
	Cert *string `json:"Cert,omitnil,omitempty" name:"Cert"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	EndorserGroupName *string `json:"EndorserGroupName,omitnil,omitempty" name:"EndorserGroupName"`

	// 背书节点列表
	EndorserPeerList []*string `json:"EndorserPeerList,omitnil,omitempty" name:"EndorserPeerList"`
}

// Predefined struct for user
type GetBlockListRequestParams struct {
	// 模块名称，固定字段：block
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名称，固定字段：block_list
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 通道ID，固定字段：0
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 组织ID，固定字段：0
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 需要查询的通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要获取的起始交易偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 需要获取的交易数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetBlockListRequest struct {
	*tchttp.BaseRequest
	
	// 模块名称，固定字段：block
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名称，固定字段：block_list
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 通道ID，固定字段：0
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 组织ID，固定字段：0
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 需要查询的通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要获取的起始交易偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 需要获取的交易数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 区块列表
	BlockList []*Block `json:"BlockList,omitnil,omitempty" name:"BlockList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：block_transaction_list_for_user
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 参与交易的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 区块ID，通过GetInvokeTx接口可以获取交易所在的区块ID
	BlockId *uint64 `json:"BlockId,omitnil,omitempty" name:"BlockId"`

	// 查询的交易列表起始偏移地址
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询的交易列表数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetBlockTransactionListForUserRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：block_transaction_list_for_user
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 参与交易的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 区块ID，通过GetInvokeTx接口可以获取交易所在的区块ID
	BlockId *uint64 `json:"BlockId,omitnil,omitempty" name:"BlockId"`

	// 查询的交易列表起始偏移地址
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询的交易列表数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 交易列表
	TransactionList []*TransactionItem `json:"TransactionList,omitnil,omitempty" name:"TransactionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type GetClusterSummaryRequestParams struct {
	// 模块名称，固定字段：cluster_mng
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名称，固定字段：cluster_summary
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 组织ID，固定字段：0
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type GetClusterSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 模块名称，固定字段：cluster_mng
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名称，固定字段：cluster_summary
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 组织ID，固定字段：0
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
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
	TotalChannelCount *uint64 `json:"TotalChannelCount,omitnil,omitempty" name:"TotalChannelCount"`

	// 当前组织创建的通道数量
	MyChannelCount *uint64 `json:"MyChannelCount,omitnil,omitempty" name:"MyChannelCount"`

	// 当前组织加入的通道数量
	JoinChannelCount *uint64 `json:"JoinChannelCount,omitnil,omitempty" name:"JoinChannelCount"`

	// 网络节点总数量
	TotalPeerCount *uint64 `json:"TotalPeerCount,omitnil,omitempty" name:"TotalPeerCount"`

	// 当前组织创建的节点数量
	MyPeerCount *uint64 `json:"MyPeerCount,omitnil,omitempty" name:"MyPeerCount"`

	// 其他组织创建的节点数量
	OrderCount *uint64 `json:"OrderCount,omitnil,omitempty" name:"OrderCount"`

	// 网络组织总数量
	TotalGroupCount *uint64 `json:"TotalGroupCount,omitnil,omitempty" name:"TotalGroupCount"`

	// 当前组织创建的组织数量
	MyGroupCount *uint64 `json:"MyGroupCount,omitnil,omitempty" name:"MyGroupCount"`

	// 网络智能合约总数量
	TotalChaincodeCount *uint64 `json:"TotalChaincodeCount,omitnil,omitempty" name:"TotalChaincodeCount"`

	// 最近7天发起的智能合约数量
	RecentChaincodeCount *uint64 `json:"RecentChaincodeCount,omitnil,omitempty" name:"RecentChaincodeCount"`

	// 当前组织发起的智能合约数量
	MyChaincodeCount *uint64 `json:"MyChaincodeCount,omitnil,omitempty" name:"MyChaincodeCount"`

	// 当前组织的证书总数量
	TotalCertCount *uint64 `json:"TotalCertCount,omitnil,omitempty" name:"TotalCertCount"`

	// 颁发给当前组织的证书数量
	TlsCertCount *uint64 `json:"TlsCertCount,omitnil,omitempty" name:"TlsCertCount"`

	// 网络背书节点证书数量
	PeerCertCount *uint64 `json:"PeerCertCount,omitnil,omitempty" name:"PeerCertCount"`

	// 当前组织业务证书数量
	ClientCertCount *uint64 `json:"ClientCertCount,omitnil,omitempty" name:"ClientCertCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：query_txid
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 执行该查询交易的节点名称，可以在通道详情中获取该通道上的节点名称及其所属组织名称
	PeerName *string `json:"PeerName,omitnil,omitempty" name:"PeerName"`

	// 执行该查询交易的节点所属组织名称，可以在通道详情中获取该通道上的节点名称及其所属组织名称
	PeerGroup *string `json:"PeerGroup,omitnil,omitempty" name:"PeerGroup"`

	// 交易ID
	TxId *string `json:"TxId,omitnil,omitempty" name:"TxId"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type GetInvokeTxRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：query_txid
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 执行该查询交易的节点名称，可以在通道详情中获取该通道上的节点名称及其所属组织名称
	PeerName *string `json:"PeerName,omitnil,omitempty" name:"PeerName"`

	// 执行该查询交易的节点所属组织名称，可以在通道详情中获取该通道上的节点名称及其所属组织名称
	PeerGroup *string `json:"PeerGroup,omitnil,omitempty" name:"PeerGroup"`

	// 交易ID
	TxId *string `json:"TxId,omitnil,omitempty" name:"TxId"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
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
	TxValidationCode *int64 `json:"TxValidationCode,omitnil,omitempty" name:"TxValidationCode"`

	// 交易执行消息
	TxValidationMsg *string `json:"TxValidationMsg,omitnil,omitempty" name:"TxValidationMsg"`

	// 交易所在区块ID
	BlockId *int64 `json:"BlockId,omitnil,omitempty" name:"BlockId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type GetLatestTransactionListRequestParams struct {
	// 模块名称，固定字段：transaction
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名称，固定字段：latest_transaction_list
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 组织ID，固定字段：0
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 通道ID，固定字段：0
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 获取的最新交易的区块数量，取值范围1~5
	LatestBlockNumber *uint64 `json:"LatestBlockNumber,omitnil,omitempty" name:"LatestBlockNumber"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 需要查询的通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要获取的起始交易偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 需要获取的交易数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetLatestTransactionListRequest struct {
	*tchttp.BaseRequest
	
	// 模块名称，固定字段：transaction
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名称，固定字段：latest_transaction_list
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 组织ID，固定字段：0
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 通道ID，固定字段：0
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 获取的最新交易的区块数量，取值范围1~5
	LatestBlockNumber *uint64 `json:"LatestBlockNumber,omitnil,omitempty" name:"LatestBlockNumber"`

	// 调用接口的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 需要查询的通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要获取的起始交易偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 需要获取的交易数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *GetLatestTransactionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLatestTransactionListRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLatestTransactionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLatestTransactionListResponseParams struct {
	// 交易总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 交易列表
	TransactionList []*TransactionItem `json:"TransactionList,omitnil,omitempty" name:"TransactionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetLatestTransactionListResponse struct {
	*tchttp.BaseResponse
	Response *GetLatestTransactionListResponseParams `json:"Response"`
}

func (r *GetLatestTransactionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLatestTransactionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTransactionDetailForUserRequestParams struct {
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：transaction_detail_for_user
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 参与交易的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 区块ID，通过GetInvokeTx接口可以获取交易所在的区块ID
	BlockId *uint64 `json:"BlockId,omitnil,omitempty" name:"BlockId"`

	// 交易ID，需要查询的详情的交易ID
	TransactionId *string `json:"TransactionId,omitnil,omitempty" name:"TransactionId"`
}

type GetTransactionDetailForUserRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：transaction_detail_for_user
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 参与交易的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 区块ID，通过GetInvokeTx接口可以获取交易所在的区块ID
	BlockId *uint64 `json:"BlockId,omitnil,omitempty" name:"BlockId"`

	// 交易ID，需要查询的详情的交易ID
	TransactionId *string `json:"TransactionId,omitnil,omitempty" name:"TransactionId"`
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
	TransactionId *string `json:"TransactionId,omitnil,omitempty" name:"TransactionId"`

	// 交易hash
	TransactionHash *string `json:"TransactionHash,omitnil,omitempty" name:"TransactionHash"`

	// 创建交易的组织名
	CreateOrgName *string `json:"CreateOrgName,omitnil,omitempty" name:"CreateOrgName"`

	// 交易类型（普通交易和配置交易）
	TransactionType *string `json:"TransactionType,omitnil,omitempty" name:"TransactionType"`

	// 交易状态
	TransactionStatus *string `json:"TransactionStatus,omitnil,omitempty" name:"TransactionStatus"`

	// 交易创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 交易数据
	TransactionData *string `json:"TransactionData,omitnil,omitempty" name:"TransactionData"`

	// 交易所在区块号
	BlockId *uint64 `json:"BlockId,omitnil,omitempty" name:"BlockId"`

	// 交易所在区块哈希
	BlockHash *string `json:"BlockHash,omitnil,omitempty" name:"BlockHash"`

	// 交易所在区块高度
	BlockHeight *uint64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`

	// 通道名称
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 交易所在合约名称
	ContractName *string `json:"ContractName,omitnil,omitempty" name:"ContractName"`

	// 背书组织列表
	EndorserOrgList []*EndorserGroup `json:"EndorserOrgList,omitnil,omitempty" name:"EndorserOrgList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type InvokeChainMakerContractRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitnil,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitnil,omitempty" name:"FuncParam"`

	// 是否异步执行，1为是，否则为0；如果异步执行，可使用返回值中的交易TxID查询执行结果
	AsyncFlag *int64 `json:"AsyncFlag,omitnil,omitempty" name:"AsyncFlag"`
}

type InvokeChainMakerContractRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitnil,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitnil,omitempty" name:"FuncParam"`

	// 是否异步执行，1为是，否则为0；如果异步执行，可使用返回值中的交易TxID查询执行结果
	AsyncFlag *int64 `json:"AsyncFlag,omitnil,omitempty" name:"AsyncFlag"`
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
	Result *ChainMakerContractResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitnil,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitnil,omitempty" name:"FuncParam"`

	// 是否异步执行，1为是，否则为0；如果异步执行，可使用返回值中的交易TxID查询执行结果
	AsyncFlag *int64 `json:"AsyncFlag,omitnil,omitempty" name:"AsyncFlag"`
}

type InvokeChainMakerDemoContractRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitnil,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitnil,omitempty" name:"FuncParam"`

	// 是否异步执行，1为是，否则为0；如果异步执行，可使用返回值中的交易TxID查询执行结果
	AsyncFlag *int64 `json:"AsyncFlag,omitnil,omitempty" name:"AsyncFlag"`
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
	Result *ChainMakerContractResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type InvokeFabricChaincodeRequestParams struct {
	// 网络ID，可在区块链网络详情获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 通道ID，可在通道列表或通道详情获取
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 合约名称，可在合约列表或合约详情获取
	ChaincodeName *string `json:"ChaincodeName,omitnil,omitempty" name:"ChaincodeName"`

	// 合约方法
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 合约方法入参
	FuncParam []*string `json:"FuncParam,omitnil,omitempty" name:"FuncParam"`

	// 是否异步执行，如果异步执行，可使用返回值中的交易TxID查询执行结果
	WithAsyncResult *bool `json:"WithAsyncResult,omitnil,omitempty" name:"WithAsyncResult"`
}

type InvokeFabricChaincodeRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 通道ID，可在通道列表或通道详情获取
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 合约名称，可在合约列表或合约详情获取
	ChaincodeName *string `json:"ChaincodeName,omitnil,omitempty" name:"ChaincodeName"`

	// 合约方法
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 合约方法入参
	FuncParam []*string `json:"FuncParam,omitnil,omitempty" name:"FuncParam"`

	// 是否异步执行，如果异步执行，可使用返回值中的交易TxID查询执行结果
	WithAsyncResult *bool `json:"WithAsyncResult,omitnil,omitempty" name:"WithAsyncResult"`
}

func (r *InvokeFabricChaincodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeFabricChaincodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ChannelId")
	delete(f, "ChaincodeName")
	delete(f, "FuncName")
	delete(f, "FuncParam")
	delete(f, "WithAsyncResult")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvokeFabricChaincodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeFabricChaincodeResponseParams struct {
	// 交易ID
	TxId *string `json:"TxId,omitnil,omitempty" name:"TxId"`

	// 交易状态
	TxStatus *string `json:"TxStatus,omitnil,omitempty" name:"TxStatus"`

	// 交易结果
	TxResult *string `json:"TxResult,omitnil,omitempty" name:"TxResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InvokeFabricChaincodeResponse struct {
	*tchttp.BaseResponse
	Response *InvokeFabricChaincodeResponseParams `json:"Response"`
}

func (r *InvokeFabricChaincodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeFabricChaincodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeRequestParams struct {
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：invoke
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务所属智能合约名称，可在智能合约详情或列表中获取
	ChaincodeName *string `json:"ChaincodeName,omitnil,omitempty" name:"ChaincodeName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 对该笔交易进行背书的节点列表（包括节点名称和节点所属组织名称，详见数据结构一节），可以在通道详情中获取该通道上的节点名称及其所属组织名称
	Peers []*PeerSet `json:"Peers,omitnil,omitempty" name:"Peers"`

	// 该笔交易需要调用的智能合约中的函数名称
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 被调用的函数参数列表，参数列表大小总和要求小于2M
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`

	// 同步调用标识，可选参数，值为0或者不传表示使用同步方法调用，调用后会等待交易执行后再返回执行结果；值为1时表示使用异步方式调用Invoke，执行后会立即返回交易对应的Txid，后续需要通过GetInvokeTx这个API查询该交易的执行结果。（对于逻辑较为简单的交易，可以使用同步模式；对于逻辑较为复杂的交易，建议使用异步模式，否则容易导致API因等待时间过长，返回等待超时）
	AsyncFlag *uint64 `json:"AsyncFlag,omitnil,omitempty" name:"AsyncFlag"`
}

type InvokeRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：invoke
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务所属智能合约名称，可在智能合约详情或列表中获取
	ChaincodeName *string `json:"ChaincodeName,omitnil,omitempty" name:"ChaincodeName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 对该笔交易进行背书的节点列表（包括节点名称和节点所属组织名称，详见数据结构一节），可以在通道详情中获取该通道上的节点名称及其所属组织名称
	Peers []*PeerSet `json:"Peers,omitnil,omitempty" name:"Peers"`

	// 该笔交易需要调用的智能合约中的函数名称
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 被调用的函数参数列表，参数列表大小总和要求小于2M
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`

	// 同步调用标识，可选参数，值为0或者不传表示使用同步方法调用，调用后会等待交易执行后再返回执行结果；值为1时表示使用异步方式调用Invoke，执行后会立即返回交易对应的Txid，后续需要通过GetInvokeTx这个API查询该交易的执行结果。（对于逻辑较为简单的交易，可以使用同步模式；对于逻辑较为复杂的交易，建议使用异步模式，否则容易导致API因等待时间过长，返回等待超时）
	AsyncFlag *uint64 `json:"AsyncFlag,omitnil,omitempty" name:"AsyncFlag"`
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
	//
	// Deprecated: Txid is deprecated.
	Txid *string `json:"Txid,omitnil,omitempty" name:"Txid"`

	// 交易执行结果
	Events *string `json:"Events,omitnil,omitempty" name:"Events"`

	// 交易ID
	TxId *string `json:"TxId,omitnil,omitempty" name:"TxId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type PeerSet struct {
	// 节点名称
	PeerName *string `json:"PeerName,omitnil,omitempty" name:"PeerName"`

	// 组织名称
	OrgName *string `json:"OrgName,omitnil,omitempty" name:"OrgName"`
}

// Predefined struct for user
type QueryChainMakerBlockTransactionRequestParams struct {
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`
}

type QueryChainMakerBlockTransactionRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`
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
	Result []*ChainMakerTransactionResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`

	// 交易数量
	TxCount *int64 `json:"TxCount,omitnil,omitempty" name:"TxCount"`

	// 区块时间戳，单位是秒
	BlockTimestamp *int64 `json:"BlockTimestamp,omitnil,omitempty" name:"BlockTimestamp"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitnil,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitnil,omitempty" name:"FuncParam"`
}

type QueryChainMakerContractRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitnil,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitnil,omitempty" name:"FuncParam"`
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
	Result *ChainMakerContractResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`
}

type QueryChainMakerDemoBlockTransactionRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`
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
	Result []*ChainMakerTransactionResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 区块高度
	BlockHeight *int64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`

	// 交易数量
	TxCount *int64 `json:"TxCount,omitnil,omitempty" name:"TxCount"`

	// 区块时间戳，单位是秒
	BlockTimestamp *int64 `json:"BlockTimestamp,omitnil,omitempty" name:"BlockTimestamp"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitnil,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitnil,omitempty" name:"FuncParam"`
}

type QueryChainMakerDemoContractRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 合约名称，可在合约管理中获取
	ContractName *string `json:"ContractName,omitnil,omitempty" name:"ContractName"`

	// 合约方法名
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 合约方法入参，json格式字符串，key/value都是string类型的map
	FuncParam *string `json:"FuncParam,omitnil,omitempty" name:"FuncParam"`
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
	Result *ChainMakerContractResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 交易ID，通过调用合约的返回值获取
	TxID *string `json:"TxID,omitnil,omitempty" name:"TxID"`
}

type QueryChainMakerDemoTransactionRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 交易ID，通过调用合约的返回值获取
	TxID *string `json:"TxID,omitnil,omitempty" name:"TxID"`
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
	Result *ChainMakerTransactionResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 交易ID，通过调用合约的返回值获取
	TxID *string `json:"TxID,omitnil,omitempty" name:"TxID"`
}

type QueryChainMakerTransactionRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务链ID，可在网络概览页获取
	ChainId *string `json:"ChainId,omitnil,omitempty" name:"ChainId"`

	// 交易ID，通过调用合约的返回值获取
	TxID *string `json:"TxID,omitnil,omitempty" name:"TxID"`
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
	Result *ChainMakerTransactionResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type QueryFabricChaincodeRequestParams struct {
	// 网络ID，可在区块链网络详情获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 通道ID，可在通道列表或通道详情获取
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 合约名称，可在合约列表或合约详情获取
	ChaincodeName *string `json:"ChaincodeName,omitnil,omitempty" name:"ChaincodeName"`

	// 合约方法
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 合约方法入参
	FuncParam []*string `json:"FuncParam,omitnil,omitempty" name:"FuncParam"`
}

type QueryFabricChaincodeRequest struct {
	*tchttp.BaseRequest
	
	// 网络ID，可在区块链网络详情获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 通道ID，可在通道列表或通道详情获取
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 合约名称，可在合约列表或合约详情获取
	ChaincodeName *string `json:"ChaincodeName,omitnil,omitempty" name:"ChaincodeName"`

	// 合约方法
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 合约方法入参
	FuncParam []*string `json:"FuncParam,omitnil,omitempty" name:"FuncParam"`
}

func (r *QueryFabricChaincodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFabricChaincodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ChannelId")
	delete(f, "ChaincodeName")
	delete(f, "FuncName")
	delete(f, "FuncParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFabricChaincodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFabricChaincodeResponseParams struct {
	// 交易ID
	TxId *string `json:"TxId,omitnil,omitempty" name:"TxId"`

	// 交易状态
	TxStatus *string `json:"TxStatus,omitnil,omitempty" name:"TxStatus"`

	// 交易结果
	TxResult *string `json:"TxResult,omitnil,omitempty" name:"TxResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryFabricChaincodeResponse struct {
	*tchttp.BaseResponse
	Response *QueryFabricChaincodeResponseParams `json:"Response"`
}

func (r *QueryFabricChaincodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFabricChaincodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryRequestParams struct {
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：query
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务所属智能合约名称，可在智能合约详情或列表中获取
	ChaincodeName *string `json:"ChaincodeName,omitnil,omitempty" name:"ChaincodeName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 执行该查询交易的节点列表（包括节点名称和节点所属组织名称，详见数据结构一节），可以在通道详情中获取该通道上的节点名称及其所属组织名称
	Peers []*PeerSet `json:"Peers,omitnil,omitempty" name:"Peers"`

	// 该笔交易查询需要调用的智能合约中的函数名称
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 被调用的函数参数列表
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`
}

type QueryRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，固定字段：transaction
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 操作名，固定字段：query
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 区块链网络ID，可在区块链网络详情或列表中获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 业务所属智能合约名称，可在智能合约详情或列表中获取
	ChaincodeName *string `json:"ChaincodeName,omitnil,omitempty" name:"ChaincodeName"`

	// 业务所属通道名称，可在通道详情或列表中获取
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 执行该查询交易的节点列表（包括节点名称和节点所属组织名称，详见数据结构一节），可以在通道详情中获取该通道上的节点名称及其所属组织名称
	Peers []*PeerSet `json:"Peers,omitnil,omitempty" name:"Peers"`

	// 该笔交易查询需要调用的智能合约中的函数名称
	FuncName *string `json:"FuncName,omitnil,omitempty" name:"FuncName"`

	// 调用合约的组织名称，可以在组织管理列表中获取当前组织的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 被调用的函数参数列表
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`
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
	Data []*string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertMark *string `json:"CertMark,omitnil,omitempty" name:"CertMark"`

	// 用户申请签名证书所需要的证书请求文件的base64编码
	SignCsrContent *string `json:"SignCsrContent,omitnil,omitempty" name:"SignCsrContent"`
}

// Predefined struct for user
type SrvInvokeRequestParams struct {
	// 服务类型，iss或者dam
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务接口，要调用的方法函数名
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 用户自定义json字符串
	Param *string `json:"Param,omitnil,omitempty" name:"Param"`
}

type SrvInvokeRequest struct {
	*tchttp.BaseRequest
	
	// 服务类型，iss或者dam
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务接口，要调用的方法函数名
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 用户自定义json字符串
	Param *string `json:"Param,omitnil,omitempty" name:"Param"`
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
	RetCode *int64 `json:"RetCode,omitnil,omitempty" name:"RetCode"`

	// 返回消息
	RetMsg *string `json:"RetMsg,omitnil,omitempty" name:"RetMsg"`

	// 返回数据
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type Transaction struct {
	// 交易ID
	TxId *string `json:"TxId,omitnil,omitempty" name:"TxId"`

	// 合约名称
	ChaincodeName *string `json:"ChaincodeName,omitnil,omitempty" name:"ChaincodeName"`

	// 交易发送者
	Sender *string `json:"Sender,omitnil,omitempty" name:"Sender"`

	// 交易创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 交易所在区块高度
	BlockHeight *int64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`

	// 交易在区块中的序号
	TxIndex *int64 `json:"TxIndex,omitnil,omitempty" name:"TxIndex"`
}

type TransactionItem struct {
	// 交易ID
	TransactionId *string `json:"TransactionId,omitnil,omitempty" name:"TransactionId"`

	// 交易hash
	TransactionHash *string `json:"TransactionHash,omitnil,omitempty" name:"TransactionHash"`

	// 创建交易的组织名
	CreateOrgName *string `json:"CreateOrgName,omitnil,omitempty" name:"CreateOrgName"`

	// 交易所在区块号
	BlockId *uint64 `json:"BlockId,omitnil,omitempty" name:"BlockId"`

	// 交易类型（普通交易和配置交易）
	TransactionType *string `json:"TransactionType,omitnil,omitempty" name:"TransactionType"`

	// 交易创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 交易所在区块高度
	BlockHeight *uint64 `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`

	// 交易状态
	TransactionStatus *string `json:"TransactionStatus,omitnil,omitempty" name:"TransactionStatus"`
}