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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ChainData struct {
	// 区块hash
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockHash *string `json:"BlockHash,omitempty" name:"BlockHash"`

	// 区块高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockHeight *string `json:"BlockHeight,omitempty" name:"BlockHeight"`

	// 区块时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockTime *string `json:"BlockTime,omitempty" name:"BlockTime"`
}

type CodeBatch struct {
	// 批次号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 企业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *int64 `json:"CorpId,omitempty" name:"CorpId"`

	// 码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchCode *string `json:"BatchCode,omitempty" name:"BatchCode"`

	// 码数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeCnt *int64 `json:"CodeCnt,omitempty" name:"CodeCnt"`

	// 所属商户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 批次类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchType *int64 `json:"BatchType,omitempty" name:"BatchType"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 微信模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	MpTpl *string `json:"MpTpl,omitempty" name:"MpTpl"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 所属商户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 未使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ext *Ext `json:"Ext,omitempty" name:"Ext"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TplName *string `json:"TplName,omitempty" name:"TplName"`
}

type CodeItem struct {
	// 无
	Code *string `json:"Code,omitempty" name:"Code"`
}

// Predefined struct for user
type CreateCodeBatchRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 批次类型 0:溯源 1:营销
	BatchType *uint64 `json:"BatchType,omitempty" name:"BatchType"`

	// 批次ID，系统自动生成
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 活动ID
	MpTpl *string `json:"MpTpl,omitempty" name:"MpTpl"`
}

type CreateCodeBatchRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 批次类型 0:溯源 1:营销
	BatchType *uint64 `json:"BatchType,omitempty" name:"BatchType"`

	// 批次ID，系统自动生成
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 活动ID
	MpTpl *string `json:"MpTpl,omitempty" name:"MpTpl"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCodeBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCodeBatchResponseParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 码长度
	CodeLength *uint64 `json:"CodeLength,omitempty" name:"CodeLength"`

	// 码类型 alphabet 字母, number 数字, mixin 混合
	CodeType *string `json:"CodeType,omitempty" name:"CodeType"`

	// 生码数量 普通码包时必填
	Amount *int64 `json:"Amount,omitempty" name:"Amount"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 码包类型 0: 普通码包 1: 层级码包
	PackType *uint64 `json:"PackType,omitempty" name:"PackType"`

	// 码包层级
	PackLevel *uint64 `json:"PackLevel,omitempty" name:"PackLevel"`

	// 码包规格
	PackSpec []*PackSpec `json:"PackSpec,omitempty" name:"PackSpec"`
}

type CreateCodePackRequest struct {
	*tchttp.BaseRequest
	
	// 商户ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 码长度
	CodeLength *uint64 `json:"CodeLength,omitempty" name:"CodeLength"`

	// 码类型 alphabet 字母, number 数字, mixin 混合
	CodeType *string `json:"CodeType,omitempty" name:"CodeType"`

	// 生码数量 普通码包时必填
	Amount *int64 `json:"Amount,omitempty" name:"Amount"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 码包类型 0: 普通码包 1: 层级码包
	PackType *uint64 `json:"PackType,omitempty" name:"PackType"`

	// 码包层级
	PackLevel *uint64 `json:"PackLevel,omitempty" name:"PackLevel"`

	// 码包规格
	PackSpec []*PackSpec `json:"PackSpec,omitempty" name:"PackSpec"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCodePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCodePackResponseParams struct {
	// 码包ID
	PackId *string `json:"PackId,omitempty" name:"PackId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateMerchantRequestParams struct {
	// 商户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type CreateMerchantRequest struct {
	*tchttp.BaseRequest
	
	// 商户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMerchantResponseParams struct {
	// 商户标识码
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 商户名称
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// 商品规格
	Specification *string `json:"Specification,omitempty" name:"Specification"`

	// 商品图片
	Logo []*string `json:"Logo,omitempty" name:"Logo"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 预留字段
	Ext *Ext `json:"Ext,omitempty" name:"Ext"`
}

type CreateProductRequest struct {
	*tchttp.BaseRequest
	
	// 商品名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 商户名称
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// 商品规格
	Specification *string `json:"Specification,omitempty" name:"Specification"`

	// 商品图片
	Logo []*string `json:"Logo,omitempty" name:"Logo"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 预留字段
	Ext *Ext `json:"Ext,omitempty" name:"Ext"`
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
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 溯源ID
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`
}

type CreateTraceChainRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 溯源ID
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`
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
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateTraceCodesRequestParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 码
	Codes []*CodeItem `json:"Codes,omitempty" name:"Codes"`
}

type CreateTraceCodesRequest struct {
	*tchttp.BaseRequest
	
	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 码
	Codes []*CodeItem `json:"Codes,omitempty" name:"Codes"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTraceCodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceCodesResponseParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 导入成功码数量
	ActiveCnt *uint64 `json:"ActiveCnt,omitempty" name:"ActiveCnt"`

	// 批次码数量
	CodeCnt *uint64 `json:"CodeCnt,omitempty" name:"CodeCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 溯源阶段 0:商品 1:通用 2:内部溯源 3:外部溯源
	Phase *uint64 `json:"Phase,omitempty" name:"Phase"`

	// 溯源阶段名称
	PhaseName *string `json:"PhaseName,omitempty" name:"PhaseName"`

	// [无效] 上链状态
	ChainStatus *uint64 `json:"ChainStatus,omitempty" name:"ChainStatus"`

	// [无效] 码类型 0: 批次, 1: 码, 2: 生产任务, 3: 物流信息
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// [无效] 溯源ID
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 溯源信息
	TraceItems []*TraceItem `json:"TraceItems,omitempty" name:"TraceItems"`
}

type CreateTraceDataRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 溯源阶段 0:商品 1:通用 2:内部溯源 3:外部溯源
	Phase *uint64 `json:"Phase,omitempty" name:"Phase"`

	// 溯源阶段名称
	PhaseName *string `json:"PhaseName,omitempty" name:"PhaseName"`

	// [无效] 上链状态
	ChainStatus *uint64 `json:"ChainStatus,omitempty" name:"ChainStatus"`

	// [无效] 码类型 0: 批次, 1: 码, 2: 生产任务, 3: 物流信息
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// [无效] 溯源ID
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 溯源信息
	TraceItems []*TraceItem `json:"TraceItems,omitempty" name:"TraceItems"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTraceDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceDataResponseParams struct {
	// 溯源ID
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DeleteCodeBatchRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

type DeleteCodeBatchRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
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
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type DeleteMerchantRequest struct {
	*tchttp.BaseRequest
	
	// 商户标识码
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type DeleteProductRequest struct {
	*tchttp.BaseRequest
	
	// 商品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type DeleteTraceDataRequest struct {
	*tchttp.BaseRequest
	
	// 溯源ID
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeCodeBatchByIdRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

type DescribeCodeBatchByIdRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
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
	CodeBatch *CodeBatch `json:"CodeBatch,omitempty" name:"CodeBatch"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeCodeBatchsRequestParams struct {
	// 查询商户ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 查询商品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 条数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页数
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 批次类型 0:溯源 1:营销
	BatchType *string `json:"BatchType,omitempty" name:"BatchType"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type DescribeCodeBatchsRequest struct {
	*tchttp.BaseRequest
	
	// 查询商户ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 查询商品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 条数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页数
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 批次类型 0:溯源 1:营销
	BatchType *string `json:"BatchType,omitempty" name:"BatchType"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCodeBatchsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodeBatchsResponseParams struct {
	// 批次列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeBatchs []*CodeBatch `json:"CodeBatchs,omitempty" name:"CodeBatchs"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeCodePacksRequestParams struct {
	// 每页数量
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type DescribeCodePacksRequest struct {
	*tchttp.BaseRequest
	
	// 每页数量
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 查询关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCodePacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodePacksResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PackId *string `json:"PackId,omitempty" name:"PackId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type DescribeCodesByPackRequest struct {
	*tchttp.BaseRequest
	
	// 码包ID
	PackId *string `json:"PackId,omitempty" name:"PackId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	Codes []*CodeItem `json:"Codes,omitempty" name:"Codes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeMerchantByIdRequestParams struct {
	// 商户标识码
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type DescribeMerchantByIdRequest struct {
	*tchttp.BaseRequest
	
	// 商户标识码
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	Merchant *Merchant `json:"Merchant,omitempty" name:"Merchant"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 条数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type DescribeMerchantsRequest struct {
	*tchttp.BaseRequest
	
	// 搜索商户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 条数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMerchantsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMerchantsResponseParams struct {
	// 商户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Merchants []*Merchant `json:"Merchants,omitempty" name:"Merchants"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeProductByIdRequestParams struct {
	// 商品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type DescribeProductByIdRequest struct {
	*tchttp.BaseRequest
	
	// 商品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	Product *Product `json:"Product,omitempty" name:"Product"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 条数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 商品ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type DescribeProductsRequest struct {
	*tchttp.BaseRequest
	
	// 商品名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 条数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 商品ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductsResponseParams struct {
	// 商品列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Products []*Product `json:"Products,omitempty" name:"Products"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeTraceCodeByIdRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 二维码
	Code *string `json:"Code,omitempty" name:"Code"`
}

type DescribeTraceCodeByIdRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 二维码
	Code *string `json:"Code,omitempty" name:"Code"`
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
	TraceCode *TraceCode `json:"TraceCode,omitempty" name:"TraceCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 条数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页码
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 批次ID，弃用
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type DescribeTraceCodesRequest struct {
	*tchttp.BaseRequest
	
	// 搜索关键字 码标识，或者批次ID
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 条数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页码
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 批次ID，弃用
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	TraceCodes []*TraceCode `json:"TraceCodes,omitempty" name:"TraceCodes"`

	// 条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeTraceDataListRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 任务ID 用于外部溯源
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 二维码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 溯源阶段 0:商品 1:通用 2:内部溯源 3:外部溯源
	Phase *uint64 `json:"Phase,omitempty" name:"Phase"`

	// 数量
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeTraceDataListRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 任务ID 用于外部溯源
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 二维码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 溯源阶段 0:商品 1:通用 2:内部溯源 3:外部溯源
	Phase *uint64 `json:"Phase,omitempty" name:"Phase"`

	// 数量
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 无
	TraceDataList []*TraceData `json:"TraceDataList,omitempty" name:"TraceDataList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Ext struct {

}

type Merchant struct {
	// 商户标识码
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 企业id
	CorpId *int64 `json:"CorpId,omitempty" name:"CorpId"`

	// 商户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 商户码规则
	CodeRule *string `json:"CodeRule,omitempty" name:"CodeRule"`
}

// Predefined struct for user
type ModifyCodeBatchRequestParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 状态 0: 未激活 1: 已激活 -1: 已冻结
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 模板ID
	MpTpl *string `json:"MpTpl,omitempty" name:"MpTpl"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 商品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyCodeBatchRequest struct {
	*tchttp.BaseRequest
	
	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 状态 0: 未激活 1: 已激活 -1: 已冻结
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 模板ID
	MpTpl *string `json:"MpTpl,omitempty" name:"MpTpl"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 商品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCodeBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCodeBatchResponseParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyMerchantRequestParams struct {
	// 商户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 商户标识码
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
}

type ModifyMerchantRequest struct {
	*tchttp.BaseRequest
	
	// 商户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 商户标识码
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMerchantResponseParams struct {
	// 商户标识码
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 商品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 商品规格
	Specification *string `json:"Specification,omitempty" name:"Specification"`

	// 商品图片
	Logo []*string `json:"Logo,omitempty" name:"Logo"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 预留字段
	Ext *Ext `json:"Ext,omitempty" name:"Ext"`
}

type ModifyProductRequest struct {
	*tchttp.BaseRequest
	
	// 商品名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 商品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 商品规格
	Specification *string `json:"Specification,omitempty" name:"Specification"`

	// 商品图片
	Logo []*string `json:"Logo,omitempty" name:"Logo"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 预留字段
	Ext *Ext `json:"Ext,omitempty" name:"Ext"`
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
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Code *string `json:"Code,omitempty" name:"Code"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 状态 0: 冻结 1: 激活
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type ModifyTraceCodeRequest struct {
	*tchttp.BaseRequest
	
	// 二维码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 状态 0: 冻结 1: 激活
	Status *uint64 `json:"Status,omitempty" name:"Status"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyTraceDataRanksRequestParams struct {
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 生产任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 溯源ID
	TraceIds []*string `json:"TraceIds,omitempty" name:"TraceIds"`
}

type ModifyTraceDataRanksRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 生产任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 溯源ID
	TraceIds []*string `json:"TraceIds,omitempty" name:"TraceIds"`
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
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 生产溯源任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 溯源信息
	TraceItems []*TraceItem `json:"TraceItems,omitempty" name:"TraceItems"`

	// 溯源阶段名称
	PhaseName *string `json:"PhaseName,omitempty" name:"PhaseName"`

	// [无效] 类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// [无效] 溯源码
	Code *string `json:"Code,omitempty" name:"Code"`

	// [无效] 排序
	Rank *uint64 `json:"Rank,omitempty" name:"Rank"`

	// [无效] 溯源阶段 0:商品 1:通用 2:物流
	Phase *uint64 `json:"Phase,omitempty" name:"Phase"`

	// [无效] 溯源时间
	TraceTime *string `json:"TraceTime,omitempty" name:"TraceTime"`

	// [无效] 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// [无效] 上链状态
	ChainStatus *uint64 `json:"ChainStatus,omitempty" name:"ChainStatus"`

	// [无效] 上链时间
	ChainTime *string `json:"ChainTime,omitempty" name:"ChainTime"`

	// [无效] 上链数据
	ChainData *ChainData `json:"ChainData,omitempty" name:"ChainData"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// [无效] 溯源状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type ModifyTraceDataRequest struct {
	*tchttp.BaseRequest
	
	// 溯源ID
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 生产溯源任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 溯源信息
	TraceItems []*TraceItem `json:"TraceItems,omitempty" name:"TraceItems"`

	// 溯源阶段名称
	PhaseName *string `json:"PhaseName,omitempty" name:"PhaseName"`

	// [无效] 类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// [无效] 溯源码
	Code *string `json:"Code,omitempty" name:"Code"`

	// [无效] 排序
	Rank *uint64 `json:"Rank,omitempty" name:"Rank"`

	// [无效] 溯源阶段 0:商品 1:通用 2:物流
	Phase *uint64 `json:"Phase,omitempty" name:"Phase"`

	// [无效] 溯源时间
	TraceTime *string `json:"TraceTime,omitempty" name:"TraceTime"`

	// [无效] 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// [无效] 上链状态
	ChainStatus *uint64 `json:"ChainStatus,omitempty" name:"ChainStatus"`

	// [无效] 上链时间
	ChainTime *string `json:"ChainTime,omitempty" name:"ChainTime"`

	// [无效] 上链数据
	ChainData *ChainData `json:"ChainData,omitempty" name:"ChainData"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// [无效] 溯源状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`
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
	delete(f, "Type")
	delete(f, "Code")
	delete(f, "Rank")
	delete(f, "Phase")
	delete(f, "TraceTime")
	delete(f, "CreateTime")
	delete(f, "ChainStatus")
	delete(f, "ChainTime")
	delete(f, "ChainData")
	delete(f, "CorpId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTraceDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTraceDataResponseParams struct {
	// 溯源ID
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type PackSpec struct {
	// 层级
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 比例
	Rate *uint64 `json:"Rate,omitempty" name:"Rate"`

	// 数量
	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`
}

type Product struct {
	// 商品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 企业id
	CorpId *int64 `json:"CorpId,omitempty" name:"CorpId"`

	// 商户标识码
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 商品编号
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 商品名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 商品规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Specification *string `json:"Specification,omitempty" name:"Specification"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 商品图片
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logo []*string `json:"Logo,omitempty" name:"Logo"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 预留字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ext *Ext `json:"Ext,omitempty" name:"Ext"`

	// 商户名称
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`
}

type TraceCode struct {
	// 码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 企业ID
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 包ID
	PackId *string `json:"PackId,omitempty" name:"PackId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 所属商户ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 商户名称
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type TraceData struct {
	// 溯源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 企业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *uint64 `json:"CorpId,omitempty" name:"CorpId"`

	// 0
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rank *uint64 `json:"Rank,omitempty" name:"Rank"`

	// 溯源阶段 0:商品 1:通用 2:物流
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phase *uint64 `json:"Phase,omitempty" name:"Phase"`

	// 环节名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhaseName *string `json:"PhaseName,omitempty" name:"PhaseName"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceTime *string `json:"TraceTime,omitempty" name:"TraceTime"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceItems []*TraceItem `json:"TraceItems,omitempty" name:"TraceItems"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 上链状态 0: 未上链 1: 上链中 2: 已上链 -1: 异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChainStatus *uint64 `json:"ChainStatus,omitempty" name:"ChainStatus"`

	// 上链时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChainTime *string `json:"ChainTime,omitempty" name:"ChainTime"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChainData *ChainData `json:"ChainData,omitempty" name:"ChainData"`
}

type TraceItem struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 单个值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 只读
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// 扫码展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hidden *bool `json:"Hidden,omitempty" name:"Hidden"`

	// 多个值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}