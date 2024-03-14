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

package v20210515

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AttrItem struct {
	// 字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 字段类型
	// text:文本类型, 
	// longtext:长文本类型, banner:单图片类型, image:多图片类型,
	// video:视频类型,
	// mp:小程序类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 只读
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 扫码展示
	Hidden *bool `json:"Hidden,omitnil,omitempty" name:"Hidden"`

	// 多个值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 类型标识
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 扩展字段
	Ext *string `json:"Ext,omitnil,omitempty" name:"Ext"`
}

// Predefined struct for user
type AuthorizedTransferRequestParams struct {
	// 业务加密入参。
	BusinessSecurityData *InputEncryptData `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type AuthorizedTransferRequest struct {
	*tchttp.BaseRequest
	
	// 业务加密入参。
	BusinessSecurityData *InputEncryptData `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *AuthorizedTransferRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuthorizedTransferRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AuthorizedTransferRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AuthorizedTransferResponseParams struct {
	// 业务出参。
	Data *OutputAuthorizedTransfer `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AuthorizedTransferResponse struct {
	*tchttp.BaseResponse
	Response *AuthorizedTransferResponseParams `json:"Response"`
}

func (r *AuthorizedTransferResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuthorizedTransferResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChainData struct {
	// 区块hash
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockHash *string `json:"BlockHash,omitnil,omitempty" name:"BlockHash"`

	// 区块高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockHeight *string `json:"BlockHeight,omitnil,omitempty" name:"BlockHeight"`

	// 区块时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockTime *string `json:"BlockTime,omitnil,omitempty" name:"BlockTime"`
}

type CodeBatch struct {
	// 批次号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 企业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *int64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次编码(未使用)
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchCode *string `json:"BatchCode,omitnil,omitempty" name:"BatchCode"`

	// 码数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeCnt *int64 `json:"CodeCnt,omitnil,omitempty" name:"CodeCnt"`

	// 所属商户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 批次类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchType *int64 `json:"BatchType,omitnil,omitempty" name:"BatchType"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 微信模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	MpTpl *string `json:"MpTpl,omitnil,omitempty" name:"MpTpl"`

	// 批次状态 0: 未激活 1: 已激活 -1: 已冻结
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 所属商户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantName *string `json:"MerchantName,omitnil,omitempty" name:"MerchantName"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 未使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Ext is deprecated.
	Ext *Ext `json:"Ext,omitnil,omitempty" name:"Ext"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TplName *string `json:"TplName,omitnil,omitempty" name:"TplName"`

	// 调度任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *Job `json:"Job,omitnil,omitempty" name:"Job"`

	// 生产日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductionDate *string `json:"ProductionDate,omitnil,omitempty" name:"ProductionDate"`

	// 有效期
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

	// 扩展属性
	Attrs []*AttrItem `json:"Attrs,omitnil,omitempty" name:"Attrs"`
}

type CodeItem struct {
	// 无
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

type CodePack struct {
	// 码id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`

	// 企业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *int64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 商户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 制码状态 init: 初始化, pending: 执行中, done: 完成, error: 失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 执行日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// 码数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Amount *int64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 防伪码长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeLength *int64 `json:"CodeLength,omitnil,omitempty" name:"CodeLength"`

	// 码类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeType *string `json:"CodeType,omitnil,omitempty" name:"CodeType"`

	// 是否暗码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cipher *int64 `json:"Cipher,omitnil,omitempty" name:"Cipher"`

	// [弃用] 文字码地址，通过另一个接口查
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextUrl *string `json:"TextUrl,omitnil,omitempty" name:"TextUrl"`

	// [弃用] 二维码地址，通过另一个接口查
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackUrl *string `json:"PackUrl,omitnil,omitempty" name:"PackUrl"`

	// 商户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantName *string `json:"MerchantName,omitnil,omitempty" name:"MerchantName"`

	// 码规则类型 0: 默认, 1: 自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 自定义码规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 码包类型 0: 普通码包 1: 层级码包
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackType *int64 `json:"PackType,omitnil,omitempty" name:"PackType"`

	// 生码层级
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackLevel *uint64 `json:"PackLevel,omitnil,omitempty" name:"PackLevel"`

	// 层级码配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackSpec []*PackSpec `json:"PackSpec,omitnil,omitempty" name:"PackSpec"`

	// 商品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 商品规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductSpecification *string `json:"ProductSpecification,omitnil,omitempty" name:"ProductSpecification"`

	// 商品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 码关系是否预关联
	// 0:否, 1:是
	RelateType *int64 `json:"RelateType,omitnil,omitempty" name:"RelateType"`
}

type CodePart struct {
	// 码段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 码段类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 码段内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 码段长度
	Length *uint64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 扩展字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ext *string `json:"Ext,omitnil,omitempty" name:"Ext"`
}

type CorpQuota struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 企业名称
	CorpName *string `json:"CorpName,omitnil,omitempty" name:"CorpName"`

	// 额度
	Quota *Quota `json:"Quota,omitnil,omitempty" name:"Quota"`

	// 额度使用量
	UsageQuota *UsageQuota `json:"UsageQuota,omitnil,omitempty" name:"UsageQuota"`
}

// Predefined struct for user
type CreateCodeBatchRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 批次类型 0:溯源 1:营销
	BatchType *uint64 `json:"BatchType,omitnil,omitempty" name:"BatchType"`

	// 批次ID，留空时系统自动生成
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 模板ID，或者活动ID
	MpTpl *string `json:"MpTpl,omitnil,omitempty" name:"MpTpl"`

	// 克隆批次ID，同时会复制溯源信息
	CloneId *string `json:"CloneId,omitnil,omitempty" name:"CloneId"`

	// 批次编号，业务字段不判断唯一性
	BatchCode *string `json:"BatchCode,omitnil,omitempty" name:"BatchCode"`

	// 有效期
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

	// 生产日期
	ProductionDate *string `json:"ProductionDate,omitnil,omitempty" name:"ProductionDate"`
}

type CreateCodeBatchRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 批次类型 0:溯源 1:营销
	BatchType *uint64 `json:"BatchType,omitnil,omitempty" name:"BatchType"`

	// 批次ID，留空时系统自动生成
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 模板ID，或者活动ID
	MpTpl *string `json:"MpTpl,omitnil,omitempty" name:"MpTpl"`

	// 克隆批次ID，同时会复制溯源信息
	CloneId *string `json:"CloneId,omitnil,omitempty" name:"CloneId"`

	// 批次编号，业务字段不判断唯一性
	BatchCode *string `json:"BatchCode,omitnil,omitempty" name:"BatchCode"`

	// 有效期
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

	// 生产日期
	ProductionDate *string `json:"ProductionDate,omitnil,omitempty" name:"ProductionDate"`
}

func (r *CreateCodeBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodeBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	delete(f, "MerchantId")
	delete(f, "ProductId")
	delete(f, "BatchType")
	delete(f, "BatchId")
	delete(f, "Remark")
	delete(f, "MpTpl")
	delete(f, "CloneId")
	delete(f, "BatchCode")
	delete(f, "ValidDate")
	delete(f, "ProductionDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCodeBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCodeBatchResponseParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCodeBatchResponse struct {
	*tchttp.BaseResponse
	Response *CreateCodeBatchResponseParams `json:"Response"`
}

func (r *CreateCodeBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodeBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCodePackRequestParams struct {
	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 码长度
	CodeLength *uint64 `json:"CodeLength,omitnil,omitempty" name:"CodeLength"`

	// 码类型 alphabet 字母, number 数字, mixin 混合
	CodeType *string `json:"CodeType,omitnil,omitempty" name:"CodeType"`

	// 生码数量 普通码包时必填
	Amount *int64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码包类型 0: 普通码包 1: 层级码包
	PackType *uint64 `json:"PackType,omitnil,omitempty" name:"PackType"`

	// 码包层级
	PackLevel *uint64 `json:"PackLevel,omitnil,omitempty" name:"PackLevel"`

	// 码包规格
	PackSpec []*PackSpec `json:"PackSpec,omitnil,omitempty" name:"PackSpec"`

	// 批次ID，如果传了生码后会同时绑定批次，并激活码
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 是否有流水码 0:无 1:有
	SerialType *uint64 `json:"SerialType,omitnil,omitempty" name:"SerialType"`

	// 关联产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 层级码时是否提前生成关联关系，默认为 1
	RelateType *int64 `json:"RelateType,omitnil,omitempty" name:"RelateType"`
}

type CreateCodePackRequest struct {
	*tchttp.BaseRequest
	
	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 码长度
	CodeLength *uint64 `json:"CodeLength,omitnil,omitempty" name:"CodeLength"`

	// 码类型 alphabet 字母, number 数字, mixin 混合
	CodeType *string `json:"CodeType,omitnil,omitempty" name:"CodeType"`

	// 生码数量 普通码包时必填
	Amount *int64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码包类型 0: 普通码包 1: 层级码包
	PackType *uint64 `json:"PackType,omitnil,omitempty" name:"PackType"`

	// 码包层级
	PackLevel *uint64 `json:"PackLevel,omitnil,omitempty" name:"PackLevel"`

	// 码包规格
	PackSpec []*PackSpec `json:"PackSpec,omitnil,omitempty" name:"PackSpec"`

	// 批次ID，如果传了生码后会同时绑定批次，并激活码
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 是否有流水码 0:无 1:有
	SerialType *uint64 `json:"SerialType,omitnil,omitempty" name:"SerialType"`

	// 关联产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 层级码时是否提前生成关联关系，默认为 1
	RelateType *int64 `json:"RelateType,omitnil,omitempty" name:"RelateType"`
}

func (r *CreateCodePackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodePackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "CodeLength")
	delete(f, "CodeType")
	delete(f, "Amount")
	delete(f, "CorpId")
	delete(f, "PackType")
	delete(f, "PackLevel")
	delete(f, "PackSpec")
	delete(f, "BatchId")
	delete(f, "SerialType")
	delete(f, "ProductId")
	delete(f, "RelateType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCodePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCodePackResponseParams struct {
	// 码包ID
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCodePackResponse struct {
	*tchttp.BaseResponse
	Response *CreateCodePackResponseParams `json:"Response"`
}

func (r *CreateCodePackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodePackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCorporationOrderRequestParams struct {
	// 企业名称
	CorpName *string `json:"CorpName,omitnil,omitempty" name:"CorpName"`

	// 所有者ID
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 溯源码额度
	CodeQuota *uint64 `json:"CodeQuota,omitnil,omitempty" name:"CodeQuota"`

	// 额度过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 金额
	Amount *int64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 联系人
	ContactPerson *string `json:"ContactPerson,omitnil,omitempty" name:"ContactPerson"`

	// 联系电话
	ContactNumber *string `json:"ContactNumber,omitnil,omitempty" name:"ContactNumber"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateCorporationOrderRequest struct {
	*tchttp.BaseRequest
	
	// 企业名称
	CorpName *string `json:"CorpName,omitnil,omitempty" name:"CorpName"`

	// 所有者ID
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 溯源码额度
	CodeQuota *uint64 `json:"CodeQuota,omitnil,omitempty" name:"CodeQuota"`

	// 额度过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 金额
	Amount *int64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 联系人
	ContactPerson *string `json:"ContactPerson,omitnil,omitempty" name:"ContactPerson"`

	// 联系电话
	ContactNumber *string `json:"ContactNumber,omitnil,omitempty" name:"ContactNumber"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateCorporationOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCorporationOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpName")
	delete(f, "Owner")
	delete(f, "CodeQuota")
	delete(f, "ExpireTime")
	delete(f, "Amount")
	delete(f, "CorpId")
	delete(f, "ContactPerson")
	delete(f, "ContactNumber")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCorporationOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCorporationOrderResponseParams struct {
	// 企业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCorporationOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateCorporationOrderResponseParams `json:"Response"`
}

func (r *CreateCorporationOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCorporationOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomPackRequestParams struct {
	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 生码数量, 普通码包时必填
	Amount *uint64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码包类型 0: 普通码包 1: 层级码包
	PackType *uint64 `json:"PackType,omitnil,omitempty" name:"PackType"`

	// 码包层级
	PackLevel *uint64 `json:"PackLevel,omitnil,omitempty" name:"PackLevel"`

	// 层级码包规则
	PackSpec []*PackSpec `json:"PackSpec,omitnil,omitempty" name:"PackSpec"`

	// 码规则ID,  和CodeParts二选一必填
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 码段配置，和CustomId二选一必填
	CodeParts []*CodePart `json:"CodeParts,omitnil,omitempty" name:"CodeParts"`

	// 批次ID，如果传了生码后会同时绑定批次，并激活码
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 是否有流水码 0:无 1:有
	SerialType *uint64 `json:"SerialType,omitnil,omitempty" name:"SerialType"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 是否预生成码关系
	// 0: 否, 1:是
	// 默认为1，仅对层级码有效
	RelateType *int64 `json:"RelateType,omitnil,omitempty" name:"RelateType"`
}

type CreateCustomPackRequest struct {
	*tchttp.BaseRequest
	
	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 生码数量, 普通码包时必填
	Amount *uint64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码包类型 0: 普通码包 1: 层级码包
	PackType *uint64 `json:"PackType,omitnil,omitempty" name:"PackType"`

	// 码包层级
	PackLevel *uint64 `json:"PackLevel,omitnil,omitempty" name:"PackLevel"`

	// 层级码包规则
	PackSpec []*PackSpec `json:"PackSpec,omitnil,omitempty" name:"PackSpec"`

	// 码规则ID,  和CodeParts二选一必填
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 码段配置，和CustomId二选一必填
	CodeParts []*CodePart `json:"CodeParts,omitnil,omitempty" name:"CodeParts"`

	// 批次ID，如果传了生码后会同时绑定批次，并激活码
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 是否有流水码 0:无 1:有
	SerialType *uint64 `json:"SerialType,omitnil,omitempty" name:"SerialType"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 是否预生成码关系
	// 0: 否, 1:是
	// 默认为1，仅对层级码有效
	RelateType *int64 `json:"RelateType,omitnil,omitempty" name:"RelateType"`
}

func (r *CreateCustomPackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomPackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "Amount")
	delete(f, "CorpId")
	delete(f, "PackType")
	delete(f, "PackLevel")
	delete(f, "PackSpec")
	delete(f, "CustomId")
	delete(f, "CodeParts")
	delete(f, "BatchId")
	delete(f, "SerialType")
	delete(f, "ProductId")
	delete(f, "RelateType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomPackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomPackResponseParams struct {
	// 码包ID
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomPackResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomPackResponseParams `json:"Response"`
}

func (r *CreateCustomPackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomPackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomRuleRequestParams struct {
	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 码长度
	CodeLength *uint64 `json:"CodeLength,omitnil,omitempty" name:"CodeLength"`

	// 码段配置
	CodeParts []*CodePart `json:"CodeParts,omitnil,omitempty" name:"CodeParts"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type CreateCustomRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 码长度
	CodeLength *uint64 `json:"CodeLength,omitnil,omitempty" name:"CodeLength"`

	// 码段配置
	CodeParts []*CodePart `json:"CodeParts,omitnil,omitempty" name:"CodeParts"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *CreateCustomRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "MerchantId")
	delete(f, "CodeLength")
	delete(f, "CodeParts")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomRuleResponseParams struct {
	// 码规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomRuleResponseParams `json:"Response"`
}

func (r *CreateCustomRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMerchantRequestParams struct {
	// 商户名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码包来源 0:自建, 1:第三发
	CodeType *int64 `json:"CodeType,omitnil,omitempty" name:"CodeType"`

	// 码包前缀地址 第三方码包时必填
	CodeUrl *string `json:"CodeUrl,omitnil,omitempty" name:"CodeUrl"`
}

type CreateMerchantRequest struct {
	*tchttp.BaseRequest
	
	// 商户名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码包来源 0:自建, 1:第三发
	CodeType *int64 `json:"CodeType,omitnil,omitempty" name:"CodeType"`

	// 码包前缀地址 第三方码包时必填
	CodeUrl *string `json:"CodeUrl,omitnil,omitempty" name:"CodeUrl"`
}

func (r *CreateMerchantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMerchantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "CorpId")
	delete(f, "CodeType")
	delete(f, "CodeUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMerchantResponseParams struct {
	// 商户标识码
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMerchantResponse struct {
	*tchttp.BaseResponse
	Response *CreateMerchantResponseParams `json:"Response"`
}

func (r *CreateMerchantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMerchantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProductRequestParams struct {
	// 商品名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 商户名称
	MerchantName *string `json:"MerchantName,omitnil,omitempty" name:"MerchantName"`

	// 商品规格
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 商品图片
	Logo []*string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 预留字段
	Ext *Ext `json:"Ext,omitnil,omitempty" name:"Ext"`
}

type CreateProductRequest struct {
	*tchttp.BaseRequest
	
	// 商品名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 商户名称
	MerchantName *string `json:"MerchantName,omitnil,omitempty" name:"MerchantName"`

	// 商品规格
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 商品图片
	Logo []*string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 预留字段
	Ext *Ext `json:"Ext,omitnil,omitempty" name:"Ext"`
}

func (r *CreateProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "MerchantId")
	delete(f, "Remark")
	delete(f, "MerchantName")
	delete(f, "Specification")
	delete(f, "Logo")
	delete(f, "CorpId")
	delete(f, "Ext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProductResponseParams struct {
	// 商品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProductResponse struct {
	*tchttp.BaseResponse
	Response *CreateProductResponseParams `json:"Response"`
}

func (r *CreateProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceChainRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 溯源ID
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`
}

type CreateTraceChainRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 溯源ID
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`
}

func (r *CreateTraceChainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTraceChainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	delete(f, "TraceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTraceChainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceChainResponseParams struct {
	// 溯源ID
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTraceChainResponse struct {
	*tchttp.BaseResponse
	Response *CreateTraceChainResponseParams `json:"Response"`
}

func (r *CreateTraceChainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTraceChainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceCodesAsyncRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 上传文件Key，仅支持 csv 或者 zip 类型
	FileKey *string `json:"FileKey,omitnil,omitempty" name:"FileKey"`
}

type CreateTraceCodesAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 上传文件Key，仅支持 csv 或者 zip 类型
	FileKey *string `json:"FileKey,omitnil,omitempty" name:"FileKey"`
}

func (r *CreateTraceCodesAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTraceCodesAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	delete(f, "BatchId")
	delete(f, "FileKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTraceCodesAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceCodesAsyncResponseParams struct {
	// 批次ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTraceCodesAsyncResponse struct {
	*tchttp.BaseResponse
	Response *CreateTraceCodesAsyncResponseParams `json:"Response"`
}

func (r *CreateTraceCodesAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTraceCodesAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceCodesRequestParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码
	Codes []*CodeItem `json:"Codes,omitnil,omitempty" name:"Codes"`

	// 码绑定激活策略，默认  0
	// 0: 传什么码就激活什么码
	// 1: 层级码 + 层级子码
	CodeType *uint64 `json:"CodeType,omitnil,omitempty" name:"CodeType"`

	// 错误检查类型，默认 0
	// 0: 没有新导入码时正常返回
	// 1: 没有新导入码时报错，并返回没有导入成功的原因
	CheckType *uint64 `json:"CheckType,omitnil,omitempty" name:"CheckType"`
}

type CreateTraceCodesRequest struct {
	*tchttp.BaseRequest
	
	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码
	Codes []*CodeItem `json:"Codes,omitnil,omitempty" name:"Codes"`

	// 码绑定激活策略，默认  0
	// 0: 传什么码就激活什么码
	// 1: 层级码 + 层级子码
	CodeType *uint64 `json:"CodeType,omitnil,omitempty" name:"CodeType"`

	// 错误检查类型，默认 0
	// 0: 没有新导入码时正常返回
	// 1: 没有新导入码时报错，并返回没有导入成功的原因
	CheckType *uint64 `json:"CheckType,omitnil,omitempty" name:"CheckType"`
}

func (r *CreateTraceCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTraceCodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	delete(f, "CorpId")
	delete(f, "Codes")
	delete(f, "CodeType")
	delete(f, "CheckType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTraceCodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceCodesResponseParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 导入成功码数量
	ActiveCnt *uint64 `json:"ActiveCnt,omitnil,omitempty" name:"ActiveCnt"`

	// 批次码数量
	CodeCnt *uint64 `json:"CodeCnt,omitnil,omitempty" name:"CodeCnt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTraceCodesResponse struct {
	*tchttp.BaseResponse
	Response *CreateTraceCodesResponseParams `json:"Response"`
}

func (r *CreateTraceCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTraceCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceDataRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 溯源阶段 0:商品 1:通用 2:生产溯源 3:销售溯源
	Phase *uint64 `json:"Phase,omitnil,omitempty" name:"Phase"`

	// 溯源阶段名称
	PhaseName *string `json:"PhaseName,omitnil,omitempty" name:"PhaseName"`

	// [无效] 上链状态
	ChainStatus *uint64 `json:"ChainStatus,omitnil,omitempty" name:"ChainStatus"`

	// [无效] 码类型 0: 批次, 1: 码, 2: 生产任务, 3: 物流信息
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// [无效] 溯源ID
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 溯源信息
	TraceItems []*TraceItem `json:"TraceItems,omitnil,omitempty" name:"TraceItems"`

	// 溯源状态 0: 无效, 1: 有效
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 环节数据
	PhaseData *PhaseData `json:"PhaseData,omitnil,omitempty" name:"PhaseData"`
}

type CreateTraceDataRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 溯源阶段 0:商品 1:通用 2:生产溯源 3:销售溯源
	Phase *uint64 `json:"Phase,omitnil,omitempty" name:"Phase"`

	// 溯源阶段名称
	PhaseName *string `json:"PhaseName,omitnil,omitempty" name:"PhaseName"`

	// [无效] 上链状态
	ChainStatus *uint64 `json:"ChainStatus,omitnil,omitempty" name:"ChainStatus"`

	// [无效] 码类型 0: 批次, 1: 码, 2: 生产任务, 3: 物流信息
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// [无效] 溯源ID
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 溯源信息
	TraceItems []*TraceItem `json:"TraceItems,omitnil,omitempty" name:"TraceItems"`

	// 溯源状态 0: 无效, 1: 有效
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 环节数据
	PhaseData *PhaseData `json:"PhaseData,omitnil,omitempty" name:"PhaseData"`
}

func (r *CreateTraceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTraceDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	delete(f, "BatchId")
	delete(f, "TaskId")
	delete(f, "Phase")
	delete(f, "PhaseName")
	delete(f, "ChainStatus")
	delete(f, "Type")
	delete(f, "TraceId")
	delete(f, "TraceItems")
	delete(f, "Status")
	delete(f, "PhaseData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTraceDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceDataResponseParams struct {
	// 溯源ID
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTraceDataResponse struct {
	*tchttp.BaseResponse
	Response *CreateTraceDataResponseParams `json:"Response"`
}

func (r *CreateTraceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTraceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomRule struct {
	// 码规则ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 码规则名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 企业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 商户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 码ID长度
	CodeLength *uint64 `json:"CodeLength,omitnil,omitempty" name:"CodeLength"`

	// 规则状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 码段配置
	CodeParts []*CodePart `json:"CodeParts,omitnil,omitempty" name:"CodeParts"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type DeleteCodeBatchRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

type DeleteCodeBatchRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

func (r *DeleteCodeBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCodeBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCodeBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCodeBatchResponseParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCodeBatchResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCodeBatchResponseParams `json:"Response"`
}

func (r *DeleteCodeBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCodeBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMerchantRequestParams struct {
	// 商户标识码
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DeleteMerchantRequest struct {
	*tchttp.BaseRequest
	
	// 商户标识码
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DeleteMerchantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMerchantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMerchantResponseParams struct {
	// 商户标识码
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMerchantResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMerchantResponseParams `json:"Response"`
}

func (r *DeleteMerchantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMerchantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProductRequestParams struct {
	// 商品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DeleteProductRequest struct {
	*tchttp.BaseRequest
	
	// 商品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DeleteProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProductResponseParams struct {
	// 商品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProductResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProductResponseParams `json:"Response"`
}

func (r *DeleteProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTraceDataRequestParams struct {
	// 溯源ID
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DeleteTraceDataRequest struct {
	*tchttp.BaseRequest
	
	// 溯源ID
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DeleteTraceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTraceDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TraceId")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTraceDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTraceDataResponseParams struct {
	// 溯源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTraceDataResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTraceDataResponseParams `json:"Response"`
}

func (r *DeleteTraceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTraceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentCorpsRequestParams struct {
	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 渠道ID
	AgentId *uint64 `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DescribeAgentCorpsRequest struct {
	*tchttp.BaseRequest
	
	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 渠道ID
	AgentId *uint64 `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DescribeAgentCorpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentCorpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "AgentId")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentCorpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentCorpsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentCorpsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentCorpsResponseParams `json:"Response"`
}

func (r *DescribeAgentCorpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentCorpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodeBatchByIdRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

type DescribeCodeBatchByIdRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

func (r *DescribeCodeBatchByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodeBatchByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCodeBatchByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodeBatchByIdResponseParams struct {
	// 批次
	CodeBatch *CodeBatch `json:"CodeBatch,omitnil,omitempty" name:"CodeBatch"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCodeBatchByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCodeBatchByIdResponseParams `json:"Response"`
}

func (r *DescribeCodeBatchByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodeBatchByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodeBatchesRequestParams struct {
	// 查询商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 查询商品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 批次类型 0:溯源 1:营销
	BatchType *string `json:"BatchType,omitnil,omitempty" name:"BatchType"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeCodeBatchesRequest struct {
	*tchttp.BaseRequest
	
	// 查询商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 查询商品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 批次类型 0:溯源 1:营销
	BatchType *string `json:"BatchType,omitnil,omitempty" name:"BatchType"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeCodeBatchesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodeBatchesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "ProductId")
	delete(f, "Keyword")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "BatchType")
	delete(f, "CorpId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCodeBatchesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodeBatchesResponseParams struct {
	// 批次列表
	CodeBatches []*CodeBatch `json:"CodeBatches,omitnil,omitempty" name:"CodeBatches"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCodeBatchesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCodeBatchesResponseParams `json:"Response"`
}

func (r *DescribeCodeBatchesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodeBatchesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodeBatchsRequestParams struct {
	// 查询商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 查询商品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 批次类型 0:溯源 1:营销
	BatchType *string `json:"BatchType,omitnil,omitempty" name:"BatchType"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeCodeBatchsRequest struct {
	*tchttp.BaseRequest
	
	// 查询商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 查询商品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 批次类型 0:溯源 1:营销
	BatchType *string `json:"BatchType,omitnil,omitempty" name:"BatchType"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeCodeBatchsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodeBatchsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "ProductId")
	delete(f, "Keyword")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "BatchType")
	delete(f, "CorpId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCodeBatchsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodeBatchsResponseParams struct {
	// 批次列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeBatchs []*CodeBatch `json:"CodeBatchs,omitnil,omitempty" name:"CodeBatchs"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCodeBatchsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCodeBatchsResponseParams `json:"Response"`
}

func (r *DescribeCodeBatchsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodeBatchsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodePackStatusRequestParams struct {
	// 码包ID
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DescribeCodePackStatusRequest struct {
	*tchttp.BaseRequest
	
	// 码包ID
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DescribeCodePackStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodePackStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackId")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCodePackStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodePackStatusResponseParams struct {
	// 码包状态 init: 初始化, pending: 执行中, done: 完成, error: 失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCodePackStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCodePackStatusResponseParams `json:"Response"`
}

func (r *DescribeCodePackStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodePackStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodePackUrlRequestParams struct {
	// 码包ID
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DescribeCodePackUrlRequest struct {
	*tchttp.BaseRequest
	
	// 码包ID
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DescribeCodePackUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodePackUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackId")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCodePackUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodePackUrlResponseParams struct {
	// 文字码包地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 图片码包地址，可能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImgUrl *string `json:"ImgUrl,omitnil,omitempty" name:"ImgUrl"`

	// 文字码包Key，用于上传导入
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileKey *string `json:"FileKey,omitnil,omitempty" name:"FileKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCodePackUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCodePackUrlResponseParams `json:"Response"`
}

func (r *DescribeCodePackUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodePackUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodePacksRequestParams struct {
	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 是否有流水码 0:无 1:有
	SerialType *uint64 `json:"SerialType,omitnil,omitempty" name:"SerialType"`

	// 资源类型 batch:批次, order_in 入库, order_out: 出入
	ResType *string `json:"ResType,omitnil,omitempty" name:"ResType"`

	// 资源ID ResType是 batch 时对应是批次ID, 是 order_in, order_out时，则是订单ID
	ResId *string `json:"ResId,omitnil,omitempty" name:"ResId"`
}

type DescribeCodePacksRequest struct {
	*tchttp.BaseRequest
	
	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 是否有流水码 0:无 1:有
	SerialType *uint64 `json:"SerialType,omitnil,omitempty" name:"SerialType"`

	// 资源类型 batch:批次, order_in 入库, order_out: 出入
	ResType *string `json:"ResType,omitnil,omitempty" name:"ResType"`

	// 资源ID ResType是 batch 时对应是批次ID, 是 order_in, order_out时，则是订单ID
	ResId *string `json:"ResId,omitnil,omitempty" name:"ResId"`
}

func (r *DescribeCodePacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodePacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Keyword")
	delete(f, "CorpId")
	delete(f, "SerialType")
	delete(f, "ResType")
	delete(f, "ResId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCodePacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodePacksResponseParams struct {
	// 码列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodePacks []*CodePack `json:"CodePacks,omitnil,omitempty" name:"CodePacks"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCodePacksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCodePacksResponseParams `json:"Response"`
}

func (r *DescribeCodePacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodePacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodesByPackRequestParams struct {
	// 码包ID
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DescribeCodesByPackRequest struct {
	*tchttp.BaseRequest
	
	// 码包ID
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DescribeCodesByPackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodesByPackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackId")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCodesByPackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodesByPackResponseParams struct {
	// 码列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codes []*CodeItem `json:"Codes,omitnil,omitempty" name:"Codes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCodesByPackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCodesByPackResponseParams `json:"Response"`
}

func (r *DescribeCodesByPackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodesByPackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCorpQuotasRequestParams struct {
	// 渠道商ID，不要传
	AgentId *uint64 `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 搜索企业ID
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeCorpQuotasRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商ID，不要传
	AgentId *uint64 `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 搜索企业ID
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeCorpQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCorpQuotasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCorpQuotasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCorpQuotasResponseParams struct {
	// 子企业额度使用情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpQuotas []*CorpQuota `json:"CorpQuotas,omitnil,omitempty" name:"CorpQuotas"`

	// 记录总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCorpQuotasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCorpQuotasResponseParams `json:"Response"`
}

func (r *DescribeCorpQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCorpQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomRuleByIdRequestParams struct {
	// 码规则ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DescribeCustomRuleByIdRequest struct {
	*tchttp.BaseRequest
	
	// 码规则ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DescribeCustomRuleByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomRuleByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomId")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomRuleByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomRuleByIdResponseParams struct {
	// 码规则信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomRule *CustomRule `json:"CustomRule,omitnil,omitempty" name:"CustomRule"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomRuleByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomRuleByIdResponseParams `json:"Response"`
}

func (r *DescribeCustomRuleByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomRuleByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomRulesRequestParams struct {
	// 搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 条数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码规则状态 0:未生效 1:已生效 -1:已失效
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`
}

type DescribeCustomRulesRequest struct {
	*tchttp.BaseRequest
	
	// 搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 条数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码规则状态 0:未生效 1:已生效 -1:已失效
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`
}

func (r *DescribeCustomRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Keyword")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "CorpId")
	delete(f, "Status")
	delete(f, "MerchantId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomRulesResponseParams struct {
	// 码规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomRules []*CustomRule `json:"CustomRules,omitnil,omitempty" name:"CustomRules"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomRulesResponseParams `json:"Response"`
}

func (r *DescribeCustomRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobFileUrlRequestParams struct {
	// 调度ID
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DescribeJobFileUrlRequest struct {
	*tchttp.BaseRequest
	
	// 调度ID
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DescribeJobFileUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobFileUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobFileUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobFileUrlResponseParams struct {
	// 码包地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobFileUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobFileUrlResponseParams `json:"Response"`
}

func (r *DescribeJobFileUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobFileUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMerchantByIdRequestParams struct {
	// 商户标识码
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DescribeMerchantByIdRequest struct {
	*tchttp.BaseRequest
	
	// 商户标识码
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DescribeMerchantByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMerchantByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMerchantByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMerchantByIdResponseParams struct {
	// 商户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Merchant *Merchant `json:"Merchant,omitnil,omitempty" name:"Merchant"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMerchantByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMerchantByIdResponseParams `json:"Response"`
}

func (r *DescribeMerchantByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMerchantByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMerchantsRequestParams struct {
	// 搜索商户名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 条数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码来源类型 0:自建, 1:第三方
	CodeType *int64 `json:"CodeType,omitnil,omitempty" name:"CodeType"`
}

type DescribeMerchantsRequest struct {
	*tchttp.BaseRequest
	
	// 搜索商户名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 条数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码来源类型 0:自建, 1:第三方
	CodeType *int64 `json:"CodeType,omitnil,omitempty" name:"CodeType"`
}

func (r *DescribeMerchantsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMerchantsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "CorpId")
	delete(f, "CodeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMerchantsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMerchantsResponseParams struct {
	// 商户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Merchants []*Merchant `json:"Merchants,omitnil,omitempty" name:"Merchants"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMerchantsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMerchantsResponseParams `json:"Response"`
}

func (r *DescribeMerchantsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMerchantsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlanQRCodeScanRecordsRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribePlanQRCodeScanRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribePlanQRCodeScanRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlanQRCodeScanRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageNo")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePlanQRCodeScanRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlanQRCodeScanRecordsResponseParams struct {
	// 返回码
	Ret *int64 `json:"Ret,omitnil,omitempty" name:"Ret"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数据
	Data []*PlanQRCodeRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePlanQRCodeScanRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePlanQRCodeScanRecordsResponseParams `json:"Response"`
}

func (r *DescribePlanQRCodeScanRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlanQRCodeScanRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlanQRCodesRequestParams struct {
	// 计划ID
	PlanId *int64 `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribePlanQRCodesRequest struct {
	*tchttp.BaseRequest
	
	// 计划ID
	PlanId *int64 `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribePlanQRCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlanQRCodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageNo")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePlanQRCodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePlanQRCodesResponseParams struct {
	// 返回码
	Ret *int64 `json:"Ret,omitnil,omitempty" name:"Ret"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数据
	Data []*PlanQRCode `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePlanQRCodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePlanQRCodesResponseParams `json:"Response"`
}

func (r *DescribePlanQRCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlanQRCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductByIdRequestParams struct {
	// 商品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DescribeProductByIdRequest struct {
	*tchttp.BaseRequest
	
	// 商品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DescribeProductByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductByIdResponseParams struct {
	// 商品信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Product *Product `json:"Product,omitnil,omitempty" name:"Product"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProductByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductByIdResponseParams `json:"Response"`
}

func (r *DescribeProductByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductsRequestParams struct {
	// 商品名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 条数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 商品ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 认证状态
	CertState *int64 `json:"CertState,omitnil,omitempty" name:"CertState"`
}

type DescribeProductsRequest struct {
	*tchttp.BaseRequest
	
	// 商品名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 条数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 商品ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 认证状态
	CertState *int64 `json:"CertState,omitnil,omitempty" name:"CertState"`
}

func (r *DescribeProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "MerchantId")
	delete(f, "CorpId")
	delete(f, "CertState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductsResponseParams struct {
	// 商品列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Products []*Product `json:"Products,omitnil,omitempty" name:"Products"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProductsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductsResponseParams `json:"Response"`
}

func (r *DescribeProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRawScanLogsRequestParams struct {
	// 企业ID, 默认为当前企业
	// 如果有渠道权限，可以传 0 会查渠道下所有的企业
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 分页数量，默认为 20，最大为 1000
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前分页，默认为 1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 从哪个日志后查询
	// 即: LogId > $AfterLogId
	AfterLogId *uint64 `json:"AfterLogId,omitnil,omitempty" name:"AfterLogId"`

	// 开始时间 >= StartTime
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间 < EndTime
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeRawScanLogsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID, 默认为当前企业
	// 如果有渠道权限，可以传 0 会查渠道下所有的企业
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 分页数量，默认为 20，最大为 1000
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前分页，默认为 1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 从哪个日志后查询
	// 即: LogId > $AfterLogId
	AfterLogId *uint64 `json:"AfterLogId,omitnil,omitempty" name:"AfterLogId"`

	// 开始时间 >= StartTime
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间 < EndTime
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeRawScanLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRawScanLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "AfterLogId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRawScanLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRawScanLogsResponseParams struct {
	// 原始扫码日志
	ScanLogs []*RawScanLog `json:"ScanLogs,omitnil,omitempty" name:"ScanLogs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRawScanLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRawScanLogsResponseParams `json:"Response"`
}

func (r *DescribeRawScanLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRawScanLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanLogsRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 分页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前分页
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 安心码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 小程序用户ID
	Openid *string `json:"Openid,omitnil,omitempty" name:"Openid"`
}

type DescribeScanLogsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 分页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前分页
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 安心码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 小程序用户ID
	Openid *string `json:"Openid,omitnil,omitempty" name:"Openid"`
}

func (r *DescribeScanLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Code")
	delete(f, "Openid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanLogsResponseParams struct {
	// 【弃用】
	//
	// Deprecated: Products is deprecated.
	Products []*ScanLog `json:"Products,omitnil,omitempty" name:"Products"`

	// 条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 扫描记录
	ScanLogs []*ScanLog `json:"ScanLogs,omitnil,omitempty" name:"ScanLogs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScanLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanLogsResponseParams `json:"Response"`
}

func (r *DescribeScanLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanStatsRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 分页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前分页
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 安心码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

type DescribeScanStatsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 分页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前分页
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 安心码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

func (r *DescribeScanStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "MerchantId")
	delete(f, "ProductId")
	delete(f, "BatchId")
	delete(f, "Code")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanStatsResponseParams struct {
	// 统计记录
	ScanStats []*ScanStat `json:"ScanStats,omitnil,omitempty" name:"ScanStats"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScanStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanStatsResponseParams `json:"Response"`
}

func (r *DescribeScanStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTmpTokenRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DescribeTmpTokenRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DescribeTmpTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTmpTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTmpTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTmpTokenResponseParams struct {
	// 临时token
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTmpTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTmpTokenResponseParams `json:"Response"`
}

func (r *DescribeTmpTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTmpTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTraceCodeByIdRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 二维码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

type DescribeTraceCodeByIdRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 二维码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

func (r *DescribeTraceCodeByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTraceCodeByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	delete(f, "Code")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTraceCodeByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTraceCodeByIdResponseParams struct {
	// 无
	TraceCode *TraceCode `json:"TraceCode,omitnil,omitempty" name:"TraceCode"`

	// 码路径，如level是2，则为 [1级, 2级]
	CodePath []*string `json:"CodePath,omitnil,omitempty" name:"CodePath"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTraceCodeByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTraceCodeByIdResponseParams `json:"Response"`
}

func (r *DescribeTraceCodeByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTraceCodeByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTraceCodesRequestParams struct {
	// 搜索关键字 码标识，或者批次ID
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 条数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页码
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 批次ID，弃用
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DescribeTraceCodesRequest struct {
	*tchttp.BaseRequest
	
	// 搜索关键字 码标识，或者批次ID
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 条数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页码
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 批次ID，弃用
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DescribeTraceCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTraceCodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Keyword")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "BatchId")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTraceCodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTraceCodesResponseParams struct {
	// 标识列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceCodes []*TraceCode `json:"TraceCodes,omitnil,omitempty" name:"TraceCodes"`

	// 条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTraceCodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTraceCodesResponseParams `json:"Response"`
}

func (r *DescribeTraceCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTraceCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTraceDataByIdRequestParams struct {
	// 溯源ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type DescribeTraceDataByIdRequest struct {
	*tchttp.BaseRequest
	
	// 溯源ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *DescribeTraceDataByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTraceDataByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTraceDataByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTraceDataByIdResponseParams struct {
	// 无
	TraceData *TraceData `json:"TraceData,omitnil,omitempty" name:"TraceData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTraceDataByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTraceDataByIdResponseParams `json:"Response"`
}

func (r *DescribeTraceDataByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTraceDataByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTraceDataListRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 任务ID 用于外部溯源
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 二维码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 溯源阶段 0:商品 1:通用 2:内部溯源 3:外部溯源
	Phase *uint64 `json:"Phase,omitnil,omitempty" name:"Phase"`

	// 数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeTraceDataListRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 任务ID 用于外部溯源
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 二维码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 溯源阶段 0:商品 1:通用 2:内部溯源 3:外部溯源
	Phase *uint64 `json:"Phase,omitnil,omitempty" name:"Phase"`

	// 数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeTraceDataListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTraceDataListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	delete(f, "BatchId")
	delete(f, "TaskId")
	delete(f, "PageNumber")
	delete(f, "Code")
	delete(f, "Phase")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTraceDataListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTraceDataListResponseParams struct {
	// 数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 无
	TraceDataList []*TraceData `json:"TraceDataList,omitnil,omitempty" name:"TraceDataList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTraceDataListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTraceDataListResponseParams `json:"Response"`
}

func (r *DescribeTraceDataListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTraceDataListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EffectFeedbackRequestParams struct {
	// 业务加密入参。
	BusinessSecurityData *InputEncryptData `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type EffectFeedbackRequest struct {
	*tchttp.BaseRequest
	
	// 业务加密入参。
	BusinessSecurityData *InputEncryptData `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *EffectFeedbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EffectFeedbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EffectFeedbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EffectFeedbackResponseParams struct {
	// 业务出参。
	Data *OutputAuthorizedTransfer `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EffectFeedbackResponse struct {
	*tchttp.BaseResponse
	Response *EffectFeedbackResponseParams `json:"Response"`
}

func (r *EffectFeedbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EffectFeedbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Ext struct {
	// 字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Value is deprecated.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type InputEncryptData struct {
	// 加密方式，0：AES加密；
	EncryptMethod *int64 `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`

	// 加密算法中的块处理模式，1：CBC模式； 目前只支持CBC模式
	EncryptMode *int64 `json:"EncryptMode,omitnil,omitempty" name:"EncryptMode"`

	// 填充模式，0：ZeroPadding；1：PKCS5Padding；2：
	// PKCS7Padding。
	PaddingType *int64 `json:"PaddingType,omitnil,omitempty" name:"PaddingType"`

	// 加密数据，将AuthorizedData结构体数组（数组最大长度不超过20）序列化成JSON字符串，对得到的字符串加密并填充到该字段。
	EncryptData *string `json:"EncryptData,omitnil,omitempty" name:"EncryptData"`

	// 用户是否授权，本接口取值：1，已授权。
	IsAuthorized *int64 `json:"IsAuthorized,omitnil,omitempty" name:"IsAuthorized"`
}

type Job struct {
	// 调度ID
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 执行状态 init:初始化, pending: 执行中, done: 执行成功, error: 执行失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type Merchant struct {
	// 商户标识码
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 企业id
	CorpId *int64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 商户名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 商户码规则
	CodeRule *string `json:"CodeRule,omitnil,omitempty" name:"CodeRule"`

	// 码来源类型 0: 安心平台 1: 第三方码
	CodeType *int64 `json:"CodeType,omitnil,omitempty" name:"CodeType"`

	// 第三方码域名前缀
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeUrl *string `json:"CodeUrl,omitnil,omitempty" name:"CodeUrl"`
}

// Predefined struct for user
type ModifyCodeBatchRequestParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 状态 0: 未激活 1: 已激活 -1: 已冻结
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 模板ID，或者活动ID
	MpTpl *string `json:"MpTpl,omitnil,omitempty" name:"MpTpl"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 商品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 批次编码，业务字段不判断唯一性
	BatchCode *string `json:"BatchCode,omitnil,omitempty" name:"BatchCode"`

	// 有效期
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

	// 生产日期
	ProductionDate *string `json:"ProductionDate,omitnil,omitempty" name:"ProductionDate"`
}

type ModifyCodeBatchRequest struct {
	*tchttp.BaseRequest
	
	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 状态 0: 未激活 1: 已激活 -1: 已冻结
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 模板ID，或者活动ID
	MpTpl *string `json:"MpTpl,omitnil,omitempty" name:"MpTpl"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 商品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 批次编码，业务字段不判断唯一性
	BatchCode *string `json:"BatchCode,omitnil,omitempty" name:"BatchCode"`

	// 有效期
	ValidDate *string `json:"ValidDate,omitnil,omitempty" name:"ValidDate"`

	// 生产日期
	ProductionDate *string `json:"ProductionDate,omitnil,omitempty" name:"ProductionDate"`
}

func (r *ModifyCodeBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCodeBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	delete(f, "CorpId")
	delete(f, "Status")
	delete(f, "MpTpl")
	delete(f, "MerchantId")
	delete(f, "ProductId")
	delete(f, "Remark")
	delete(f, "BatchCode")
	delete(f, "ValidDate")
	delete(f, "ProductionDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCodeBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCodeBatchResponseParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCodeBatchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCodeBatchResponseParams `json:"Response"`
}

func (r *ModifyCodeBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCodeBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomRuleRequestParams struct {
	// 码规则ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 码长度
	CodeLength *uint64 `json:"CodeLength,omitnil,omitempty" name:"CodeLength"`

	// 码段配置
	CodeParts []*CodePart `json:"CodeParts,omitnil,omitempty" name:"CodeParts"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type ModifyCustomRuleRequest struct {
	*tchttp.BaseRequest
	
	// 码规则ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 码长度
	CodeLength *uint64 `json:"CodeLength,omitnil,omitempty" name:"CodeLength"`

	// 码段配置
	CodeParts []*CodePart `json:"CodeParts,omitnil,omitempty" name:"CodeParts"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *ModifyCustomRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomId")
	delete(f, "Name")
	delete(f, "CodeLength")
	delete(f, "CodeParts")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomRuleResponseParams struct {
	// 码规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCustomRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomRuleResponseParams `json:"Response"`
}

func (r *ModifyCustomRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomRuleStatusRequestParams struct {
	// 码规则ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 码规则状态 0:未生效 1:已生效 -1:已失效
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type ModifyCustomRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 码规则ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 码规则状态 0:未生效 1:已生效 -1:已失效
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *ModifyCustomRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomId")
	delete(f, "Status")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomRuleStatusResponseParams struct {
	// 码规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCustomRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyCustomRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMerchantRequestParams struct {
	// 商户名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商户标识码
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码包来源 0:自建, 1:第三码包，暂不支持修改
	CodeType *int64 `json:"CodeType,omitnil,omitempty" name:"CodeType"`

	// 码包前缀地址 第三方码包时必填
	CodeUrl *string `json:"CodeUrl,omitnil,omitempty" name:"CodeUrl"`
}

type ModifyMerchantRequest struct {
	*tchttp.BaseRequest
	
	// 商户名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商户标识码
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码包来源 0:自建, 1:第三码包，暂不支持修改
	CodeType *int64 `json:"CodeType,omitnil,omitempty" name:"CodeType"`

	// 码包前缀地址 第三方码包时必填
	CodeUrl *string `json:"CodeUrl,omitnil,omitempty" name:"CodeUrl"`
}

func (r *ModifyMerchantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMerchantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "MerchantId")
	delete(f, "Remark")
	delete(f, "CorpId")
	delete(f, "CodeType")
	delete(f, "CodeUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMerchantResponseParams struct {
	// 商户标识码
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMerchantResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMerchantResponseParams `json:"Response"`
}

func (r *ModifyMerchantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMerchantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProductRequestParams struct {
	// 商品名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 商品规格
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 商品图片
	Logo []*string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 预留字段
	Ext *Ext `json:"Ext,omitnil,omitempty" name:"Ext"`
}

type ModifyProductRequest struct {
	*tchttp.BaseRequest
	
	// 商品名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 商品规格
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 商品图片
	Logo []*string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 预留字段
	Ext *Ext `json:"Ext,omitnil,omitempty" name:"Ext"`
}

func (r *ModifyProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ProductId")
	delete(f, "Remark")
	delete(f, "Specification")
	delete(f, "Logo")
	delete(f, "CorpId")
	delete(f, "Ext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProductResponseParams struct {
	// 商品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProductResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProductResponseParams `json:"Response"`
}

func (r *ModifyProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTraceCodeRequestParams struct {
	// 二维码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 状态 0: 冻结 1: 激活
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyTraceCodeRequest struct {
	*tchttp.BaseRequest
	
	// 二维码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 状态 0: 冻结 1: 激活
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyTraceCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTraceCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Code")
	delete(f, "CorpId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTraceCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTraceCodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTraceCodeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTraceCodeResponseParams `json:"Response"`
}

func (r *ModifyTraceCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTraceCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTraceCodeUnlinkRequestParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 溯源码列表
	Codes []*string `json:"Codes,omitnil,omitempty" name:"Codes"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type ModifyTraceCodeUnlinkRequest struct {
	*tchttp.BaseRequest
	
	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 溯源码列表
	Codes []*string `json:"Codes,omitnil,omitempty" name:"Codes"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *ModifyTraceCodeUnlinkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTraceCodeUnlinkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	delete(f, "Codes")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTraceCodeUnlinkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTraceCodeUnlinkResponseParams struct {
	// 成功解绑溯源码的数量
	UnlinkCnt *uint64 `json:"UnlinkCnt,omitnil,omitempty" name:"UnlinkCnt"`

	// 当前批次的码数量
	CodeCnt *uint64 `json:"CodeCnt,omitnil,omitempty" name:"CodeCnt"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTraceCodeUnlinkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTraceCodeUnlinkResponseParams `json:"Response"`
}

func (r *ModifyTraceCodeUnlinkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTraceCodeUnlinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTraceDataRanksRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 生产任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 溯源ID
	TraceIds []*string `json:"TraceIds,omitnil,omitempty" name:"TraceIds"`
}

type ModifyTraceDataRanksRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 生产任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 溯源ID
	TraceIds []*string `json:"TraceIds,omitnil,omitempty" name:"TraceIds"`
}

func (r *ModifyTraceDataRanksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTraceDataRanksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CorpId")
	delete(f, "BatchId")
	delete(f, "TaskId")
	delete(f, "TraceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTraceDataRanksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTraceDataRanksResponseParams struct {
	// 批次ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTraceDataRanksResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTraceDataRanksResponseParams `json:"Response"`
}

func (r *ModifyTraceDataRanksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTraceDataRanksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTraceDataRequestParams struct {
	// 溯源ID
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 生产溯源任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 溯源信息
	TraceItems []*TraceItem `json:"TraceItems,omitnil,omitempty" name:"TraceItems"`

	// 溯源阶段名称
	PhaseName *string `json:"PhaseName,omitnil,omitempty" name:"PhaseName"`

	// 环节数据
	PhaseData *PhaseData `json:"PhaseData,omitnil,omitempty" name:"PhaseData"`

	// 溯源状态 0: 无效, 1: 有效
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序
	Rank *uint64 `json:"Rank,omitnil,omitempty" name:"Rank"`

	// [无效] 类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// [无效] 溯源码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// [无效] 溯源阶段 0:商品 1:通用 2:生产溯源 3:销售溯源
	Phase *uint64 `json:"Phase,omitnil,omitempty" name:"Phase"`

	// [无效] 溯源时间
	TraceTime *string `json:"TraceTime,omitnil,omitempty" name:"TraceTime"`

	// [无效] 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// [无效] 上链状态
	ChainStatus *uint64 `json:"ChainStatus,omitnil,omitempty" name:"ChainStatus"`

	// [无效] 上链时间
	ChainTime *string `json:"ChainTime,omitnil,omitempty" name:"ChainTime"`

	// [无效] 上链数据
	ChainData *ChainData `json:"ChainData,omitnil,omitempty" name:"ChainData"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

type ModifyTraceDataRequest struct {
	*tchttp.BaseRequest
	
	// 溯源ID
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 生产溯源任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 溯源信息
	TraceItems []*TraceItem `json:"TraceItems,omitnil,omitempty" name:"TraceItems"`

	// 溯源阶段名称
	PhaseName *string `json:"PhaseName,omitnil,omitempty" name:"PhaseName"`

	// 环节数据
	PhaseData *PhaseData `json:"PhaseData,omitnil,omitempty" name:"PhaseData"`

	// 溯源状态 0: 无效, 1: 有效
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序
	Rank *uint64 `json:"Rank,omitnil,omitempty" name:"Rank"`

	// [无效] 类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// [无效] 溯源码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// [无效] 溯源阶段 0:商品 1:通用 2:生产溯源 3:销售溯源
	Phase *uint64 `json:"Phase,omitnil,omitempty" name:"Phase"`

	// [无效] 溯源时间
	TraceTime *string `json:"TraceTime,omitnil,omitempty" name:"TraceTime"`

	// [无效] 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// [无效] 上链状态
	ChainStatus *uint64 `json:"ChainStatus,omitnil,omitempty" name:"ChainStatus"`

	// [无效] 上链时间
	ChainTime *string `json:"ChainTime,omitnil,omitempty" name:"ChainTime"`

	// [无效] 上链数据
	ChainData *ChainData `json:"ChainData,omitnil,omitempty" name:"ChainData"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`
}

func (r *ModifyTraceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTraceDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TraceId")
	delete(f, "BatchId")
	delete(f, "TaskId")
	delete(f, "TraceItems")
	delete(f, "PhaseName")
	delete(f, "PhaseData")
	delete(f, "Status")
	delete(f, "Rank")
	delete(f, "Type")
	delete(f, "Code")
	delete(f, "Phase")
	delete(f, "TraceTime")
	delete(f, "CreateTime")
	delete(f, "ChainStatus")
	delete(f, "ChainTime")
	delete(f, "ChainData")
	delete(f, "CorpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTraceDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTraceDataResponseParams struct {
	// 溯源ID
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTraceDataResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTraceDataResponseParams `json:"Response"`
}

func (r *ModifyTraceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTraceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OutputAuthorizedTransfer struct {
	// 推送状态，0表示成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 错误信息描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type PackSpec struct {
	// 层级
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 比例
	Rate *uint64 `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 数量
	Amount *uint64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 码规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 码段配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeParts []*CodePart `json:"CodeParts,omitnil,omitempty" name:"CodeParts"`
}

type PhaseData struct {
	// 启用头
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadEnabled *bool `json:"HeadEnabled,omitnil,omitempty" name:"HeadEnabled"`

	// 标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadTitle *string `json:"HeadTitle,omitnil,omitempty" name:"HeadTitle"`

	// 标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 小程序AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 小程序AppPath
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppPath *string `json:"AppPath,omitnil,omitempty" name:"AppPath"`

	// 小程序名称AppName
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`
}

type PlanQRCode struct {
	// 二维码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 状态，0:未激活 1:已激活 2:已冻结
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type PlanQRCodeRecord struct {
	// 二维码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// OpenID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 扫码时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`

	// IP 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 国家
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitnil,omitempty" name:"City"`
}

type Product struct {
	// 商户标识码
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 商品名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商品id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 企业id
	CorpId *int64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 商品编号
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 商品规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 商品图片
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logo []*string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 预留字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ext *Ext `json:"Ext,omitnil,omitempty" name:"Ext"`

	// 商户名称
	MerchantName *string `json:"MerchantName,omitnil,omitempty" name:"MerchantName"`

	// 认证状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertState *int64 `json:"CertState,omitnil,omitempty" name:"CertState"`
}

type Quota struct {
	// 服务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 服务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 配额ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaId *uint64 `json:"QuotaId,omitnil,omitempty" name:"QuotaId"`

	// 企业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 开通服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Services []*string `json:"Services,omitnil,omitempty" name:"Services"`

	// 商户配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	FactoryQuota *int64 `json:"FactoryQuota,omitnil,omitempty" name:"FactoryQuota"`

	// 商品配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemQuota *int64 `json:"ItemQuota,omitnil,omitempty" name:"ItemQuota"`

	// 溯源码配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrackQuota *int64 `json:"TrackQuota,omitnil,omitempty" name:"TrackQuota"`

	// 销售码配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	SaleQuota *int64 `json:"SaleQuota,omitnil,omitempty" name:"SaleQuota"`

	// 上链配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChainQuota *int64 `json:"ChainQuota,omitnil,omitempty" name:"ChainQuota"`

	// 风控配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskQuota *int64 `json:"RiskQuota,omitnil,omitempty" name:"RiskQuota"`

	// AI文字数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AigcTextQuota *int64 `json:"AigcTextQuota,omitnil,omitempty" name:"AigcTextQuota"`

	// AI图片数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AigcImageQuota *int64 `json:"AigcImageQuota,omitnil,omitempty" name:"AigcImageQuota"`

	// 溯源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrackType *int64 `json:"TrackType,omitnil,omitempty" name:"TrackType"`

	// 开通版本 lite:轻量版, basic:基础版, standard:标准版
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 是否开启企业认证
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCertify *uint64 `json:"ProductCertify,omitnil,omitempty" name:"ProductCertify"`
}

type RawScanLog struct {
	// 日志ID
	LogId *int64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 微信小程序openid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Openid *string `json:"Openid,omitnil,omitempty" name:"Openid"`

	// 扫码时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 溯源码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 商户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 商品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 批次ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 地市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 区/县
	// 注意：此字段可能返回 null，表示取不到有效值。
	District *string `json:"District,omitnil,omitempty" name:"District"`
}

// Predefined struct for user
type ReportBatchCallbackStatusRequestParams struct {
	// 业务加密入参。
	BusinessSecurityData *InputEncryptData `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type ReportBatchCallbackStatusRequest struct {
	*tchttp.BaseRequest
	
	// 业务加密入参。
	BusinessSecurityData *InputEncryptData `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *ReportBatchCallbackStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportBatchCallbackStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportBatchCallbackStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportBatchCallbackStatusResponseParams struct {
	// 业务出参。
	Data *OutputAuthorizedTransfer `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReportBatchCallbackStatusResponse struct {
	*tchttp.BaseResponse
	Response *ReportBatchCallbackStatusResponseParams `json:"Response"`
}

func (r *ReportBatchCallbackStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportBatchCallbackStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanLog struct {
	// 行ID
	LogId *int64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 微信openid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Openid *string `json:"Openid,omitnil,omitempty" name:"Openid"`

	// 微信昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 企业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *int64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 商户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 商品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// ip地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 国家
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 县/区
	// 注意：此字段可能返回 null，表示取不到有效值。
	District *string `json:"District,omitnil,omitempty" name:"District"`

	// 微信 unionid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unionid *string `json:"Unionid,omitnil,omitempty" name:"Unionid"`

	// 首次扫码 0:否, 1:是
	// 注意：此字段可能返回 null，表示取不到有效值。
	First *int64 `json:"First,omitnil,omitempty" name:"First"`

	// 批次ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 扫码类型 0:无效扫码 1: 小程序扫码 2: 商家扫码
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 商户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantName *string `json:"MerchantName,omitnil,omitempty" name:"MerchantName"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`
}

type ScanStat struct {
	// 安心码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 扫码次数
	Pv *uint64 `json:"Pv,omitnil,omitempty" name:"Pv"`

	// 扫码人数
	Uv *uint64 `json:"Uv,omitnil,omitempty" name:"Uv"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 商户名称
	MerchantName *string `json:"MerchantName,omitnil,omitempty" name:"MerchantName"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`
}

type TraceCode struct {
	// 二维码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码包ID
	PackId *string `json:"PackId,omitnil,omitempty" name:"PackId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 所属商户ID
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 码状态 0: 冻结 1: 激活
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 商户名称
	MerchantName *string `json:"MerchantName,omitnil,omitempty" name:"MerchantName"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 渠道商ID
	AgentId *uint64 `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 码层级 0: 最小级, 1: 一级, 2: 二级
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`
}

type TraceData struct {
	// 溯源ID
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 码类型 0: 批次, 1: 码, 2: 生产任务
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 码值，跟码类型一一对应
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 排序，在Phase相同情况下，值越小排名靠前
	Rank *uint64 `json:"Rank,omitnil,omitempty" name:"Rank"`

	// 溯源阶段 0:商品 1:通用 2:生产溯源 3:销售溯源
	Phase *uint64 `json:"Phase,omitnil,omitempty" name:"Phase"`

	// 溯源环节名称
	PhaseName *string `json:"PhaseName,omitnil,omitempty" name:"PhaseName"`

	// 溯源时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceTime *string `json:"TraceTime,omitnil,omitempty" name:"TraceTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 上链状态 0: 未上链 1: 上链中 2: 已上链 -1: 异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChainStatus *uint64 `json:"ChainStatus,omitnil,omitempty" name:"ChainStatus"`

	// 上链时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChainTime *string `json:"ChainTime,omitnil,omitempty" name:"ChainTime"`

	// 上链数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChainData *ChainData `json:"ChainData,omitnil,omitempty" name:"ChainData"`

	// 溯源阶段配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhaseData *PhaseData `json:"PhaseData,omitnil,omitempty" name:"PhaseData"`

	// 溯源阶段状态 0: 无效, 1: 有效
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 无
	TraceItems []*TraceItem `json:"TraceItems,omitnil,omitempty" name:"TraceItems"`
}

type TraceItem struct {
	// 字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 字段类型
	// text:文本类型, 
	// longtext:长文本类型, banner:单图片类型, image:多图片类型,
	// video:视频类型,
	// mp:小程序类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 多个值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 只读
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 扫码展示
	Hidden *bool `json:"Hidden,omitnil,omitempty" name:"Hidden"`

	// 类型标识
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 扩展字段
	Ext *string `json:"Ext,omitnil,omitempty" name:"Ext"`

	// 额外属性
	Attrs []*TraceItem `json:"Attrs,omitnil,omitempty" name:"Attrs"`

	// 子页面，只读
	List []*TraceData `json:"List,omitnil,omitempty" name:"List"`
}

type UsageQuota struct {
	// 企业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *uint64 `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// 商户配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	FactoryCnt *int64 `json:"FactoryCnt,omitnil,omitempty" name:"FactoryCnt"`

	// 商品数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemCnt *int64 `json:"ItemCnt,omitnil,omitempty" name:"ItemCnt"`

	// 溯源码量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrackCnt *int64 `json:"TrackCnt,omitnil,omitempty" name:"TrackCnt"`

	// 营销码额度
	// 注意：此字段可能返回 null，表示取不到有效值。
	SaleCnt *int64 `json:"SaleCnt,omitnil,omitempty" name:"SaleCnt"`

	// 区块链上链次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChainCnt *int64 `json:"ChainCnt,omitnil,omitempty" name:"ChainCnt"`

	// 营销风控次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskCnt *int64 `json:"RiskCnt,omitnil,omitempty" name:"RiskCnt"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}